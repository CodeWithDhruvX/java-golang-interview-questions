# 30. Spring Boot (Data Access & Security)

**Q: Database Views in JPA**
> "JPA doesn't officially distinguish between a Table and a View.
> You map a View exactly like a Table: using `@Entity`.
>
> **The Trick**: Since Views are read-only, you should verify that you don't accidentally try to save data to it. Use `@Immutable` (Hibernate annotation) on the entity to prevent writes."

**Indepth:**
> **Performance**: Be careful. If the View logic is complex (joins, aggregations), querying it might be slow. The View code runs inside the DB engine, so specialized indexes on the underlying tables are crucial.


---

**Q: Pagination with Slice vs Page**
> "Both are used for pagination, but `Page` is more expensive.
>
> *   **Page**: Returns the data chunk **plus** the total count of pages/rows. This requires an extra `COUNT(*)` query, which can be slow on huge tables.
> *   **Slice**: Returns the data chunk and just a flag `hasNext()`. It doesn't know the total size. Use this for 'Infinite Scroll' features where you don't care if there are 100 or 1000 pages left, just 'is there more?'."

**Indepth:**
> **Mobile Apps**: `Slice` is perfect for mobile "Infinite Scroll". Calculating total pages (`COUNT(*)`) is expensive and often unnecessary for a Twitter-like feed where you just want the next 10 items.


---

**Q: Custom Repository Implementation**
> "Sometimes the standard `save/findAll` isn't enough, and `@Query` is too messy.
>
> 1.  Create an interface `MyRepoCustom` with your method: `complexSearch()`.
> 2.  Create a class `MyRepoImpl` implementing it. **CRITICAL**: The name must end in `Impl` (by default).
> 3.  Inject `EntityManager` inside `Impl` and write full custom logic (Criteria API or Native SQL).
> 4.  Have your main interface extend `JpaRepository` AND `MyRepoCustom`. Spring merges them automatically."

**Indepth:**
> **Composition**: This pattern (Composition over Inheritance) allows you to keep the clean `findBy` methods of `JpaRepository` while injecting completely arbitrary code execution into the same bean.


---

**Q: @Secured vs @PreAuthorize**
> "**@Secured**: Older, simple, but limited. `@Secured("ROLE_ADMIN")`.
>
> "**@PreAuthorize**: The modern standard. It supports SpEL (Spring Expression Language), allowing widely powerful logic:
> `@PreAuthorize("hasRole('ADMIN') or #param.name == authentication.name")`.
> Always use `@PreAuthorize` unless you're on a very old legacy system."

**Indepth:**
> **SecurityContext**: SpEL checks happen against the `SecurityContext`. You can even access method arguments: `@PreAuthorize("#order.owner == authentication.name")`. This is fine-grained "Instance Level Security".


---

**Q: OAuth2 Login Flow**
> "In Spring Boot, it's almost zero-config.
> 1.  Add `spring-boot-starter-oauth2-client`.
> 2.  In `application.yml`, register your provider (Google/GitHub) with `clientId` and `clientSecret`.
>
> Spring automatically configures the login page (`/login`), handles the redirect to Google, catches the callback code, exchanges it for a Token, and logs the user in."

**Indepth:**
> **CommonAuth2**: Note that "Social Login" (Google) and "Enterprise SSO" (Okta/Active Directory) both use the standardized OAuth2/OIDC protocol. Spring handles them identically.


---

**Q: Filter Chain (Spring Security)**
> "Spring Security is essentially a giant chain of Servlet Filters.
> Request -> [LogoutFilter] -> [UsernamePasswordAuthFilter] -> [BasicAuthFilter] -> ... -> [MyCustomFilter] -> Controller.
>
> If any filter throws an exception or denies access, the request stops there. It never reaches your controller. You customize security by inserting your own filters into this chain."

**Indepth:**
> **DelegatingFilterProxy**: The bridge between the Servlet container (Tomcat) and Spring's ApplicationContext is the `DelegatingFilterProxy`. It delegates standard servlet requests to Spring beans (the SecurityFilterChain).


---

**Q: Securing Actuator Endpoints**
> "Actuator endpoints (like `/heapdump` or `/env`) expose sensitive data.
>
> 1.  **Exclude by default**: `management.endpoints.web.exposure.include=health,info` (Only safe ones).
> 2.  **Secure via Security Config**:
>     `http.requestMatchers("/actuator/**").hasRole("ADMIN")`
>     Never leave Actuator open to the public Internet."

**Indepth:**
> **Information Leakage**: A `heapdump` endpoint exposed publicly allows attackers to download your entire memory, extracting passwords, API keys, and customer data. This is a catastrophic vulnerability.


---

**Q: JWT (JSON Web Token) Implementation**
> "Spring Security doesn't have a default 'Generate JWT' button. You implement it:
>
> 1.  **Login**: User sends User/Pass. You verify it.
> 2.  **Generate**: detailed library (jjwt/nimbus) to sign a Token containing claims (User ID, Role, Expiry). Return it to client.
> 3.  **Validate**: Add a `JwtFilter` that runs before every request. It parses the header `Authorization: Bearer <token>`, verifies the signature, and sets the `SecurityContext` if valid."

**Indepth:**
> **Statelessness**: JWTs make the server stateless. You don't need a session store (Redis). The token itself contains the user data. The trade-off is revocation: you can't easily ban a user until their token expires.


---

**Q: CORS (Cross-Origin Resource Sharing)**
> "If your Frontend (React on port 3000) tries to call Backend (Boot on port 8080), browsers block it.
>
> Global Fix in Spring:
> ```java
> @Bean
> WebMvcConfigurer corsConfigurer() {
>     return new WebMvcConfigurer() {
>         public void addCorsMappings(Registry registry) {
>             registry.addMapping("/**").allowedOrigins("http://localhost:3000");
>         }
>     };
> }
> ```
> In Spring Security, you must also enable `http.cors()` in the security filter chain."

**Indepth:**
> **Pre-flight**: The browser sends an `OPTIONS` request first (Pre-flight) to check if the cross-origin call is allowed. If your server doesn't handle `OPTIONS`, the real request never happens.


---

**Q: CSRF (Cross-Site Request Forgery)**
> "CSRF attacks trick a user's browser into sending a request (like 'Transfer Money') while they are logged in.
>
> **Stateful Apps (Session)**: ENABLE CSRF. Spring does this by default. It expects a CSRF token with every POST/PUT.
>
> **Stateless Apps (JWT/REST)**: DISABLE CSRF. Since you don't rely on cookies for auth, the browser can't be tricked in the same way. `http.csrf().disable()`."

**Indepth:**
> **SameSite**: Modern browsers execute strict `SameSite` cookie policies, which partially mitigates CSRF. However, explicit token validation is still the defense-in-depth standard for session-based apps.


---

**Q: Method Level Security**
> "To secure specific service methods (not just URLs):
> 1.  Add `@EnableMethodSecurity` to your main config.
> 2.  Annotate methods:
>     `@PreAuthorize("hasAuthority('WRITE_PRIVILEGE')")`
>     `public void deleteDatabase() { ... }`
>
> This is critical for 'Defense in Depth'. Even if a hacker bypasses the URL check, the service layer stops them."

**Indepth:**
> **AOP**: Method security is implemented using AOP (Aspect Oriented Programming) proxies. The check happens *before* the method body executes. If auth fails, an `AccessDeniedException` is thrown.


---

**Q: API Key Authentication**
> "For machine-to-machine communication where you don't need a user login.
>
> 1.  Create a `OncePerRequestFilter`.
> 2.  Check for a header: `X-API-KEY`.
> 3.  Compare it against a stored value (DB or Properties).
> 4.  If valid, manually create an `Authentication` object (like `ApiKeyAuthenticationToken`) and push it to the context."

**Indepth:**
> **M2M**: API Keys are often long-lived. For better security, rotate them periodically. In high-security systems, use mTLS (Mutual TLS) where the client presents a certificate instead of a header string.

