# Spring Security Interview Questions

## Core Concepts

### 1. What is Spring Security?
Spring Security is a powerful and highly customizable authentication and access-control framework. It is the de-facto standard for securing Spring-based applications. It focuses on providing both authentication and authorization to Java applications.

### 2. What are the main features of Spring Security?
- **Authentication:** Verifying the identity of a user (who they claim to be).
- **Authorization:** Verifying if the authenticated user has permission to access a resource.
- **Protection against common attacks:** CSRF (Cross-Site Request Forgery), Session Fixation, Clickjacking.
- **Servlet API integration.**
- **Optional integration with Spring Web MVC.**

### 3. Explain the Spring Security Filter Chain/Architecture.
Spring Security's infrastructure is based on Servlet Filters.
1.  **DelegatingFilterProxy:** A standard Servlet Filter that delegates execution to a Spring Bean (FilterChainProxy).
2.  **FilterChainProxy:** Manages the security filter chain. It decides which security filters should be applied to a request.
3.  **SecurityFilterChain:** Contains a list of standard filters (e.g., `UsernamePasswordAuthenticationFilter`, `BasicAuthenticationFilter`, `ExceptionTranslationFilter`, `FilterSecurityInterceptor`).

### 4. What is the difference between Authentication and Authorization?
- **Authentication (AuthN):** "Who are you?" Validating credentials (username/password, tokens, biometrics).
- **Authorization (AuthZ):** "What are you allowed to do?" Checking permissions and roles after successful authentication.

## Configuration & Filters

### 5. How do you configure Spring Security in a Spring Boot application?
In modern Spring Boot (Security 5.7+), we define a `SecurityFilterChain` bean instead of extending `WebSecurityConfigurerAdapter`.
```java
@Configuration
@EnableWebSecurity
public class SecurityConfig {
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .csrf(AbstractHttpConfigurer::disable)
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/public/**").permitAll()
                .anyRequest().authenticated()
            )
            .httpBasic(Customizer.withDefaults()); // or .formLogin()
        return http.build();
    }
}
```

### 6. What is the `SecurityContext` and `SecurityContextHolder`?
- **SecurityContext:** Contains the `Authentication` object of the currently authenticated user.
- **SecurityContextHolder:** A helper class to access the `SecurityContext`. By default, it stores the context in a `ThreadLocal`, making it available throughout the thread execution.

### 7. What is the `UserDetailsService` interface?
It is a core interface used to retrieve user data. It has one method: `loadUserByUsername(String username)`. Can be implemented to load users from a database, LDAP, or memory.
```java
public interface UserDetailsService {
    UserDetails loadUserByUsername(String username) throws UsernameNotFoundException;
}
```

## Advanced Security (OAuth2, JWT, Method Security)

### 8. Explain OAuth2 flow simply.
OAuth2 is an authorization framework.
1.  **User** wants to access a resource (e.g., Google Photos) via a **Client** app.
2.  **Client** redirects User to **Authorization Server** (Google).
3.  **User** logs in and grants permission.
4.  **Authorization Server** issues an **Access Token** to the Client.
5.  **Client** uses the token to access the **Resource Server** (Google Photos API).

### 9. What is JWT and how does it work with Spring Security?
**JWT (JSON Web Token)** is a compact, URL-safe means of representing claims to be transferred between two parties.
- **Implementation:**
    1.  User logs in.
    2.  Server validates credentials and generates a signed JWT.
    3.  Server sends JWT to client.
    4.  Client sends JWT in `Authorization: Bearer <token>` header for subsequent requests.
    5.  Server validates the signature of the JWT on every request (stateless).

### 10. How do you secure specific methods (Method Security)?
Enable it with `@EnableMethodSecurity`.
Then use annotations:
- `@PreAuthorize("hasRole('ADMIN')")`: Checks before method execution.
- `@PostAuthorize`: Checks after execution.
- `@Secured("ROLE_USER")`: Legacy annotation.

### 11. What is CSRF and why do we disable it for REST APIs?
**CSRF (Cross-Site Request Forgery)** tricks a user into executing unwanted actions on a web application where they are authenticated.
- **Browser-based apps:** Essential to enable CSRF protection.
- **REST APIs (Stateless):** Since REST APIs usually use tokens (JWT) instead of cookies for auth, and are stateless, CSRF protection is typically disabled because there is no session to hijack in the traditional sense.

### 12. What is `AuthenticationProvider`?
It is responsible for authentication logic. It validates the `Authentication` object. You can implement custom providers for different auth mechanisms (e.g., OTP, Biometric, Legacy Systems).

### 13. How to handle password encoding?
Never store passwords in plain text. Use `PasswordEncoder`.
- `BCryptPasswordEncoder` (Standard industry practice)
- `Pbkdf2PasswordEncoder`
- `SCryptPasswordEncoder`

```java
@Bean
public PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder();
}
```

## Expert Level Topics

### 14. How do you add a custom filter to the security chain?
You can execute a custom filter before or after a standard filter using `addFilterBefore()` or `addFilterAfter()` in the `SecurityFilterChain` configuration.
```java
http.addFilterBefore(new JwtAuthenticationFilter(), UsernamePasswordAuthenticationFilter.class);
```

### 15. What is Role Hierarchy?
Spring Security supports role hierarchy, where a higher role automatically includes the authorities of a lower role.
- Example: `ROLE_ADMIN > ROLE_USER`. If a user has `ROLE_ADMIN`, they automatically pass checks for `ROLE_USER`.
- Configured via `RoleHierarchy` bean.

### 16. @Secured vs @PreAuthorize vs @RolesAllowed
- **`@Secured`:** Spring specific, legacy. Supports simple role checks.
- **`@PreAuthorize`:** Spring specific, powerful (SpEL support). Can check arguments, combine conditions (`hasRole('A') or hasRole('B')`).
- **`@RolesAllowed`:** JSR-250 standard (Java standard). Similar to `@Secured`.

### 17. Detailed difference between CORS and CSRF?
- **CORS (Cross-Origin Resource Sharing):** A browser security feature that restricts cross-origin HTTP requests. It protects the *client* (data leakage). You configure it to *allow* other domains to access your API.
- **CSRF:** An attack vector. It protects the *server* from unauthorized state-changing actions. You configure protections to *block* malicious requests.
