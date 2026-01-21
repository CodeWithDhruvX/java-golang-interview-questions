## 🔹 Section 5: Spring Security Deep Dive (181-200)

### Question 181: How do you configure method-level security in Spring Boot?

**Answer:**
1.  Add `@EnableMethodSecurity`.
2.  Use annotations on Service/Controller methods:
    `@PreAuthorize("hasRole('ADMIN') and #id == authentication.id")`.

---

### Question 182: What is the difference between `@Secured`, `@PreAuthorize`, and `@RolesAllowed`?

**Answer:**
- **`@Secured`:** Core Spring (Legacy). Simple role check (`"ROLE_ADMIN"`). No SpEL.
- **`@RolesAllowed`:** JSR-250 (Java Standard). Similar to Secured.
- **`@PreAuthorize`:** Spring EL enabled. Powerful. Can access method arguments and return objects.

---

### Question 183: How do you customize user details service in Spring Security?

**Answer:**
Implement `UserDetailsService` interface.
Override `loadUserByUsername(String username)`.
Fetch user from DB/LDAP.
Return `UserDetails` object (containing password and authorities).

---

### Question 184: How do you implement role hierarchy in Spring Boot Security?

**Answer:**
Bean `RoleHierarchy`.
```java
@Bean
public RoleHierarchy roleHierarchy() {
    RoleHierarchyImpl r = new RoleHierarchyImpl();
    r.setHierarchy("ROLE_ADMIN > ROLE_USER");
    return r;
}
```
If you have Admin role, you automatically have User role permissions.

---

### Question 185: How to enable OAuth2 login in Spring Boot?

**Answer:**
Dependency: `spring-boot-starter-oauth2-client`.
Props:
```yaml
spring:
  security:
    oauth2:
      client:
        registration:
          google:
            client-id: ...
            client-secret: ...
```
Security Config: `.oauth2Login(withDefaults())`.

---

### Question 186: How does Spring Security filter chain work in Boot apps?

**Answer:**
It's a chain of Servlet Filters (DelegatingFilterProxy -> FilterChainProxy).
Order matters.
Includes: `SecurityContextPersistenceFilter` -> `UsernamePasswordAuthenticationFilter` -> `ExceptionTranslationFilter` -> `FilterSecurityInterceptor`.

---

### Question 187: What is the difference between `AuthenticationManager` and `SecurityContext`?

**Answer:**
- **AuthenticationManager:** The API that fails or validates credentials. Using `authenticate()` method.
- **SecurityContext:** Stores the currently authenticated user (`Principal`). Accessible via `SecurityContextHolder`.

---

### Question 188: How do you secure actuator endpoints in Spring Boot?

**Answer:**
They are just HTTP endpoints. Secure them in `SecurityFilterChain`.
`http.requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole("ADMIN")`.

---

### Question 189: How to implement custom JWT token generation and validation?

**Answer:**
(See Q83). Use library `jjwt` or `nimbus-jose-jwt`.

---

### Question 190: How do you secure REST endpoints using token-based authentication?

**Answer:**
Add a Custom Filter *before* `UsernamePasswordAuthenticationFilter`.
Read token -> Validate -> Set Context.

---

### Question 191: How do you secure WebSockets in a Spring Boot application?

**Answer:**
Extend `AbstractSecurityWebSocketMessageBrokerConfigurer`.
Override `configureInbound` methods.
`messages.simpDestMatchers("/user/**").authenticated()`.

---

### Question 192: How to configure custom CORS in Spring Security?

**Answer:**
If Spring Security is present, it MUST handle CORS (before MVC).
Bean `CorsConfigurationSource`.
Apply `.cors(cors -> cors.configurationSource(source))` in Filter Chain.

---

### Question 193: What is CSRF protection and how do you enable/disable it in Spring Boot?

**Answer:**
(See Q84).

---

### Question 194: How to create a custom login success handler?

**Answer:**
Implement `AuthenticationSuccessHandler`.
Override `onAuthenticationSuccess`.
Logic: Redirect user to `/dashboard` if Admin, or `/home` if User.
Config: `.formLogin().successHandler(myHandler)`.

---

### Question 195: What is the difference between stateless and stateful authentication?

**Answer:**
- **Stateful (Session):** Server stores Session ID in memory/Redis. Client sends Cookie. Harder to scale horizontally.
- **Stateless (JWT):** Server stores nothing. Token contains all info signature. Easy to scale.

---

### Question 196: How do you implement MFA (Multi-Factor Authentication) in Spring Boot?

**Answer:**
After 1st Login step (Password), return a temporary "PRE_AUTH" token.
User must submit OTP.
Validate OTP.
Issue Final JWT.
Usually handled by Identity Providers (Keycloak/Okta) rather than raw Spring Security code.

---

### Question 197: How do you secure method invocations in service layers?

**Answer:**
(See Q181). `@PreAuthorize`.

---

### Question 198: How to implement API key-based authentication in Spring Boot?

**Answer:**
Custom Filter checking header `X-API-KEY`.
Lookup Key in DB.
Use `PreAuthenticatedAuthenticationToken`.

---

### Question 199: What are the common pitfalls when using Spring Security?

**Answer:**
1.  Disabling CSRF for Session-based apps (opens vulnerability).
2.  Exposing Actuator publicly.
3.  Improper Matcher order (Broadest rules first instead of Specific ones).
4.  Leaking detailed error messages on Auth failure.

---

### Question 200: How do you use OAuth2 scopes in Spring Boot security?

**Answer:**
Scopes control access rights (e.g., `read:email`).
In Resource Server: `.jwt().jwtAuthenticationConverter(...)`.
Convert scopes "LOCALE" to Authorities "SCOPE_LOCALE".
Then use `@PreAuthorize("hasAuthority('SCOPE_read')")`.

---
