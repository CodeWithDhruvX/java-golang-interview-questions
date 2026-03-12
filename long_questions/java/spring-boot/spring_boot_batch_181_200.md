## 🔹 Section 5: Spring Security Deep Dive (181-200)

### Question 181: How do you configure method-level security in Spring Boot?

**Answer:**
1.  Add `@EnableMethodSecurity`.
2.  Use annotations on Service/Controller methods:
    `@PreAuthorize("hasRole('ADMIN') and #id == authentication.id")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure method-level security in Spring Boot?
**Your Response:** "I configure method-level security by adding `@EnableMethodSecurity` to my configuration class and then using security annotations directly on my service or controller methods. The most powerful annotation is `@PreAuthorize` which allows me to write complex expressions like `@PreAuthorize('hasRole('ADMIN') and #id == authentication.id')` where I can check roles and even access method parameters. This gives me fine-grained control over who can execute specific methods. Spring Security evaluates these expressions before the method executes, throwing an access denied exception if the security check fails."

---

### Question 182: What is the difference between `@Secured`, `@PreAuthorize`, and `@RolesAllowed`?

**Answer:**
- **`@Secured`:** Core Spring (Legacy). Simple role check (`"ROLE_ADMIN"`). No SpEL.
- **`@RolesAllowed`:** JSR-250 (Java Standard). Similar to Secured.
- **`@PreAuthorize`:** Spring EL enabled. Powerful. Can access method arguments and return objects.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@Secured`, `@PreAuthorize`, and `@RolesAllowed`?
**Your Response:** "These annotations serve different purposes and come from different standards. `@Secured` is the original Spring Security annotation - it's simple but limited to basic role checks like `'ROLE_ADMIN'` and doesn't support SpEL expressions. `@RolesAllowed` is the JSR-250 standard annotation, similar to `@Secured` but part of the Java EE standard. `@PreAuthorize` is the most powerful - it supports Spring Expression Language, allowing me to access method parameters, authentication details, and write complex security logic. I always prefer `@PreAuthorize` for its flexibility and expressiveness."

---

### Question 183: How do you customize user details service in Spring Security?

**Answer:**
Implement `UserDetailsService` interface.
Override `loadUserByUsername(String username)`.
Fetch user from DB/LDAP.
Return `UserDetails` object (containing password and authorities).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you customize user details service in Spring Security?
**Your Response:** "I customize authentication by implementing the `UserDetailsService` interface and overriding the `loadUserByUsername(String username)` method. In this method, I fetch the user from my database, LDAP, or any other authentication source, and return a `UserDetails` object containing the user's password and authorities. Spring Security uses this service during authentication to validate credentials. This approach gives me complete control over how users are loaded and what authorities they have, allowing me to integrate with any user storage system while maintaining Spring Security's authentication framework."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement role hierarchy in Spring Boot Security?
**Your Response:** "I implement role hierarchy by creating a `RoleHierarchy` bean that defines the relationship between roles. For example, I create a bean that sets the hierarchy as `'ROLE_ADMIN > ROLE_USER'`, meaning anyone with the Admin role automatically gets all User role permissions. This eliminates the need to assign both roles to users and makes the security model more maintainable. Spring Security automatically resolves these hierarchical relationships during authorization checks, so I only need to check for the most specific role in my security annotations."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to enable OAuth2 login in Spring Boot?
**Your Response:** "I enable OAuth2 login by adding the `spring-boot-starter-oauth2-client` dependency and configuring the OAuth2 client properties. I set up the client registration with provider details like Google, including client-id and client-secret. Then in my security configuration, I call `.oauth2Login(withDefaults())` to enable OAuth2 login. Spring Boot handles the entire OAuth2 flow automatically - redirecting users to the provider, handling the callback, and establishing the authentication. This makes it incredibly easy to add social login or enterprise SSO to my applications."

---

### Question 186: How does Spring Security filter chain work in Boot apps?

**Answer:**
It's a chain of Servlet Filters (DelegatingFilterProxy -> FilterChainProxy).
Order matters.
Includes: `SecurityContextPersistenceFilter` -> `UsernamePasswordAuthenticationFilter` -> `ExceptionTranslationFilter` -> `FilterSecurityInterceptor`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Security filter chain work in Boot apps?
**Your Response:** "Spring Security uses a chain of servlet filters, starting with `DelegatingFilterProxy` which delegates to `FilterChainProxy`. The order of filters is crucial - it starts with `SecurityContextPersistenceFilter` to load and save the security context, followed by authentication filters like `UsernamePasswordAuthenticationFilter`, then `ExceptionTranslationFilter` for handling security exceptions, and finally `FilterSecurityInterceptor` for authorization decisions. Each filter has a specific responsibility, and the chain processes requests in sequence. This modular approach makes Spring Security highly customizable and extensible."

---

### Question 187: What is the difference between `AuthenticationManager` and `SecurityContext`?

**Answer:**
- **AuthenticationManager:** The API that fails or validates credentials. Using `authenticate()` method.
- **SecurityContext:** Stores the currently authenticated user (`Principal`). Accessible via `SecurityContextHolder`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `AuthenticationManager` and `SecurityContext`?
**Your Response:** "`AuthenticationManager` and `SecurityContext` serve different purposes in the authentication process. `AuthenticationManager` is the API that validates credentials - it takes authentication requests and either returns an authenticated Authentication object or throws an exception. `SecurityContext` stores the currently authenticated user's Authentication object after successful authentication. I access the SecurityContext via `SecurityContextHolder` to get the current user anywhere in my application. So AuthenticationManager handles the authentication process, while SecurityContext holds the result."

---

### Question 188: How do you secure actuator endpoints in Spring Boot?

**Answer:**
They are just HTTP endpoints. Secure them in `SecurityFilterChain`.
`http.requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole("ADMIN")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure actuator endpoints in Spring Boot?
**Your Response:** "I secure actuator endpoints using Spring Security's `SecurityFilterChain` configuration. I use `http.requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole('ADMIN')` to restrict access to all actuator endpoints to users with the ADMIN role. I can also secure specific endpoints individually using `EndpointRequest.to(HealthEndpoint.class)` for fine-grained control. Since actuator endpoints expose sensitive application information, I always secure them in production, typically restricting them to administrative users only. This approach integrates seamlessly with my overall security configuration."

---

### Question 189: How to implement custom JWT token generation and validation?

**Answer:**
(See Q83). Use library `jjwt` or `nimbus-jose-jwt`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement custom JWT token generation and validation?
**Your Response:** "I implement custom JWT handling using libraries like `jjwt` or `nimbus-jose-jwt`. For token generation, I create a JWT with claims like username, authorities, and expiration time, then sign it with a secret key. For validation, I parse the token, verify the signature, check expiration, and extract the claims to create an Authentication object. I typically implement this in a custom filter that intercepts requests, validates the token from the Authorization header, and sets the security context. This gives me complete control over the JWT format and validation logic while integrating with Spring Security."

---

### Question 190: How do you secure REST endpoints using token-based authentication?

**Answer:**
Add a Custom Filter *before* `UsernamePasswordAuthenticationFilter`.
Read token -> Validate -> Set Context.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure REST endpoints using token-based authentication?
**Your Response:** "I secure REST endpoints with token-based authentication by adding a custom filter before the `UsernamePasswordAuthenticationFilter`. My filter reads the token from the Authorization header, validates it using my JWT library, and if valid, creates an Authentication object and sets it in the SecurityContext. This approach makes every subsequent request authenticated without requiring session state. The filter handles both token validation and context setting, making the authentication transparent to my controllers. This is the standard approach for securing stateless REST APIs with JWT tokens."

---

### Question 191: How do you secure WebSockets in a Spring Boot application?

**Answer:**
Extend `AbstractSecurityWebSocketMessageBrokerConfigurer`.
Override `configureInbound` methods.
`messages.simpDestMatchers("/user/**").authenticated()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure WebSockets in a Spring Boot application?
**YourResponse:** "I secure WebSockets by extending `AbstractSecurityWebSocketMessageBrokerConfigurer` and overriding the `configureInbound` methods. I use `messages.simpDestMatchers('/user/**').authenticated()` to require authentication for specific destination patterns. This allows me to apply security rules to WebSocket messages based on their destinations, similar to how I secure HTTP endpoints. I can also authorize based on message types or user roles. This integration ensures that WebSocket connections respect the same security rules as the rest of my application, preventing unauthorized access to real-time features."

---

### Question 192: How to configure custom CORS in Spring Security?

**Answer:**
If Spring Security is present, it MUST handle CORS (before MVC).
Bean `CorsConfigurationSource`.
Apply `.cors(cors -> cors.configurationSource(source))` in Filter Chain.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to configure custom CORS in Spring Security?
**Your Response:** "When Spring Security is present, it must handle CORS before MVC processes the request. I configure CORS by creating a `CorsConfigurationSource` bean that defines my CORS policies - allowed origins, methods, headers, and credentials. Then I apply this configuration in my security filter chain using `.cors(cors -> cors.configurationSource(source))`. This ensures that Spring Security handles CORS preflight requests and applies the correct CORS headers. This approach is essential when Spring Security is in the chain, as it prevents CORS issues that would otherwise occur due to security filtering."

---

### Question 193: What is CSRF protection and how do you enable/disable it in Spring Boot?

**Answer:**
(See Q84).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CSRF protection and how do you enable/disable it in Spring Boot?
**Your Response:** "CSRF protection prevents Cross-Site Request Forgery attacks by requiring a security token for state-changing requests. Spring Security enables CSRF protection by default for session-based authentication. I can disable it using `.csrf().disable()` in my security configuration, but I only do this for stateless APIs that use JWT tokens, where CSRF isn't a concern. For traditional web applications with form submissions, I keep CSRF protection enabled and ensure my forms include the CSRF token. The protection works by comparing the token from the request with one stored in the user's session."

---

### Question 194: How to create a custom login success handler?

**Answer:**
Implement `AuthenticationSuccessHandler`.
Override `onAuthenticationSuccess`.
Logic: Redirect user to `/dashboard` if Admin, or `/home` if User.
Config: `.formLogin().successHandler(myHandler)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to create a custom login success handler?
**Your Response:** "I create custom login success handling by implementing the `AuthenticationSuccessHandler` interface and overriding the `onAuthenticationSuccess` method. In this method, I implement logic like redirecting admins to `/dashboard` and regular users to `/home` based on their roles. I can also log login attempts, set custom cookies, or perform other post-authentication actions. Then I configure Spring Security to use my handler with `.formLogin().successHandler(myHandler)`. This gives me complete control over what happens after successful authentication, allowing me to implement custom user flows."

---

### Question 195: What is the difference between stateless and stateful authentication?

**Answer:**
- **Stateful (Session):** Server stores Session ID in memory/Redis. Client sends Cookie. Harder to scale horizontally.
- **Stateless (JWT):** Server stores nothing. Token contains all info signature. Easy to scale.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between stateless and stateful authentication?
**Your Response:** "Stateful authentication uses server-side sessions - the server stores the session ID and associates it with user data, while the client only sends a session cookie. This makes horizontal scaling harder because sessions need to be shared across servers. Stateless authentication with JWT stores nothing on the server - the token contains all necessary information and is cryptographically signed. This makes scaling much easier since any server can validate the token. I choose stateful for traditional web apps and stateless for REST APIs and microservices where scalability is important."

---

### Question 196: How do you implement MFA (Multi-Factor Authentication) in Spring Boot?

**Answer:**
After 1st Login step (Password), return a temporary "PRE_AUTH" token.
User must submit OTP.
Validate OTP.
Issue Final JWT.
Usually handled by Identity Providers (Keycloak/Okta) rather than raw Spring Security code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement MFA (Multi-Factor Authentication) in Spring Boot?
**Your Response:** "I implement MFA in a multi-step process. After the first login step with password, I return a temporary 'PRE_AUTH' token instead of the final JWT. The user must then submit their OTP through a separate endpoint. I validate the OTP and if correct, issue the final JWT. This approach requires careful state management and token handling. In practice, I often use identity providers like Keycloak or Okta for MFA rather than implementing it from scratch, as they provide robust MFA solutions with TOTP, SMS, and push notifications out of the box."

---

### Question 197: How do you secure method invocations in service layers?

**Answer:**
(See Q181). `@PreAuthorize`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure method invocations in service layers?
**Your Response:** "I secure method invocations in service layers using method-level security annotations, primarily `@PreAuthorize`. I apply these annotations directly to my service methods to enforce security rules based on user roles, permissions, or method parameters. For example, `@PreAuthorize('hasRole('ADMIN')')` restricts method access to administrators. This approach provides fine-grained security at the business logic level, ensuring that even if someone bypasses controller security, they still can't execute unauthorized service operations. It's a crucial defense-in-depth strategy for securing application logic."

---

### Question 198: How to implement API key-based authentication in Spring Boot?

**Answer:**
Custom Filter checking header `X-API-KEY`.
Lookup Key in DB.
Use `PreAuthenticatedAuthenticationToken`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to implement API key-based authentication in Spring Boot?
**Your Response:** "I implement API key authentication by creating a custom filter that checks for the `X-API-KEY` header. I look up the key in my database to validate it and retrieve the associated user or service account. If valid, I create a `PreAuthenticatedAuthenticationToken` with the API key's authorities and set it in the SecurityContext. This approach is perfect for service-to-service communication or for providing programmatic access to my APIs. Unlike username/password authentication, API keys are long-lived and don't expire, making them ideal for automated systems."

---

### Question 199: What are the common pitfalls when using Spring Security?

**Answer:**
1.  Disabling CSRF for Session-based apps (opens vulnerability).
2.  Exposing Actuator publicly.
3.  Improper Matcher order (Broadest rules first instead of Specific ones).
4.  Leaking detailed error messages on Auth failure.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the common pitfalls when using Spring Security?
**Your Response:** "I've seen several common Spring Security pitfalls. First, developers often disable CSRF protection for session-based apps, which opens serious vulnerabilities. Second, exposing actuator endpoints publicly without proper security restrictions. Third, using incorrect matcher order - putting broad rules before specific ones, which prevents specific rules from ever matching. Fourth, leaking detailed error messages on authentication failures that can help attackers. I always recommend keeping CSRF enabled for web apps, securing all actuator endpoints, ordering matchers from specific to general, and using generic error messages."

---

### Question 200: How do you use OAuth2 scopes in Spring Boot security?

**Answer:**
Scopes control access rights (e.g., `read:email`).
In Resource Server: `.jwt().jwtAuthenticationConverter(...)`.
Convert scopes "LOCALE" to Authorities "SCOPE_LOCALE".
Then use `@PreAuthorize("hasAuthority('SCOPE_read')")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use OAuth2 scopes in Spring Boot security?
**Your Response:** "OAuth2 scopes control access rights to specific resources, like `read:email` or `write:posts`. In a resource server, I configure scope handling by customizing the JWT authentication converter to transform scope claims into authorities. I convert scopes like 'read' into authorities with the prefix 'SCOPE_read'. Then I can use `@PreAuthorize('hasAuthority('SCOPE_read')')` to secure endpoints based on scopes. This approach allows fine-grained access control where different clients can have different permissions based on the scopes they were granted during OAuth2 authorization."

---
