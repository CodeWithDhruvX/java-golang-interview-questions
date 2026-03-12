# Spring Security Interview Questions

## Core Concepts

### 1. What is Spring Security?
Spring Security is a powerful and highly customizable authentication and access-control framework. It is the de-facto standard for securing Spring-based applications. It focuses on providing both authentication and authorization to Java applications.

**Explanation:** Spring Security provides a comprehensive security solution that integrates seamlessly with Spring applications through filters and dependency injection, offering declarative security configuration and protection against common vulnerabilities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Security?
**Your Response:** Spring Security is the standard security framework for Spring applications that provides both authentication and authorization capabilities. It works through a filter chain that intercepts all requests and applies security rules. What makes it powerful is its flexibility - I can secure everything from simple web apps to complex microservices using the same framework. It handles everything from basic form login to OAuth2 and JWT, and protects against common attacks like CSRF out of the box.

### 2. What are the main features of Spring Security?
- **Authentication:** Verifying the identity of a user (who they claim to be).
- **Authorization:** Verifying if the authenticated user has permission to access a resource.
- **Protection against common attacks:** CSRF (Cross-Site Request Forgery), Session Fixation, Clickjacking.
- **Servlet API integration.**
- **Optional integration with Spring Web MVC.**

**Explanation:** Spring Security provides a complete security ecosystem, from basic auth mechanisms to advanced protection against sophisticated attacks, all while maintaining clean integration with Spring's programming model.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the main features of Spring Security?
**Your Response:** Spring Security offers comprehensive security features. It handles authentication - verifying who users are through various methods like passwords, tokens, or OAuth2. Then it provides authorization to control what authenticated users can access. Importantly, it includes built-in protection against common attacks like CSRF and session fixation. It integrates cleanly with the Servlet API and Spring MVC, making it easy to secure Spring applications without major architectural changes.

### 3. Explain the Spring Security Filter Chain/Architecture.
Spring Security's infrastructure is based on Servlet Filters.
1.  **DelegatingFilterProxy:** A standard Servlet Filter that delegates execution to a Spring Bean (FilterChainProxy).
2.  **FilterChainProxy:** Manages the security filter chain. It decides which security filters should be applied to a request.
3.  **SecurityFilterChain:** Contains a list of standard filters (e.g., `UsernamePasswordAuthenticationFilter`, `BasicAuthenticationFilter`, `ExceptionTranslationFilter`, `FilterSecurityInterceptor`).

**Explanation:** The filter chain architecture allows Spring Security to apply security policies consistently across all requests while maintaining flexibility in which filters are applied based on request characteristics.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain the Spring Security filter chain architecture?
**Your Response:** Spring Security uses a sophisticated filter chain architecture. It starts with DelegatingFilterProxy, which is a standard servlet filter that bridges the servlet container and Spring application. This delegates to FilterChainProxy, which manages multiple security filter chains. Each SecurityFilterChain contains specific security filters like authentication filters, authorization filters, and exception handlers. This design allows different security configurations for different URL patterns while maintaining a clean separation of concerns.

### 4. What is the difference between Authentication and Authorization?
- **Authentication (AuthN):** "Who are you?" Validating credentials (username/password, tokens, biometrics).
- **Authorization (AuthZ):** "What are you allowed to do?" Checking permissions and roles after successful authentication.

**Explanation:** Authentication establishes identity, while authorization defines access rights. This two-step process ensures that only verified users can access resources, and only within their permitted scope.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between authentication and authorization?
**Your Response:** Authentication and authorization are two distinct but related security concepts. Authentication is about verifying who someone is - like checking a username and password or validating a JWT token. It answers the question 'Who are you?'. Authorization happens after authentication and determines what that person is allowed to do - answering 'What can you access?'. For example, I authenticate a user with their credentials, then authorize them to access only their own data based on their roles and permissions.

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

**Explanation:** The component-based configuration provides better type safety and composition compared to the older inheritance-based approach, making security configuration more maintainable and testable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure Spring Security in Spring Boot?
**Your Response:** In modern Spring Boot applications, I configure Spring Security by creating a SecurityFilterChain bean. I use @Configuration and @EnableWebSecurity annotations, then define a filterChain method that takes HttpSecurity as a parameter. In this method, I can configure CSRF protection, define authorization rules for different URL patterns, set up authentication methods like form login or HTTP basic, and build the security filter chain. This approach is more flexible than the older WebSecurityConfigurerAdapter method.

### 6. What is the `SecurityContext` and `SecurityContextHolder`?
- **SecurityContext:** Contains the `Authentication` object of the currently authenticated user.
- **SecurityContextHolder:** A helper class to access the `SecurityContext`. By default, it stores the context in a `ThreadLocal`, making it available throughout the thread execution.

**Explanation:** This pattern provides thread-safe access to authentication information throughout the application, enabling security decisions at any layer without passing authentication objects explicitly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is SecurityContext and SecurityContextHolder?
**Your Response:** SecurityContext holds the Authentication object for the current user, while SecurityContextHolder is the utility class that provides access to this context. By default, Spring Security stores the SecurityContext in a ThreadLocal, which means it's available throughout the current request thread. This allows me to access user authentication information anywhere in my application - from controllers to services to repositories - without having to pass it around manually. I can get the current user's details anytime using SecurityContextHolder.getContext().getAuthentication().

### 7. What is the `UserDetailsService` interface?
It is a core interface used to retrieve user data. It has one method: `loadUserByUsername(String username)`. Can be implemented to load users from a database, LDAP, or memory.
```java
public interface UserDetailsService {
    UserDetails loadUserByUsername(String username) throws UsernameNotFoundException;
}
```

**Explanation:** UserDetailsService serves as the bridge between authentication mechanisms and user data sources, enabling flexible user management while maintaining a consistent interface for Spring Security.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the UserDetailsService interface?
**Your Response:** UserDetailsService is Spring Security's core interface for loading user-specific data. It has just one method - loadUserByUsername - that takes a username and returns a UserDetails object. I implement this interface to connect Spring Security to my user data source, whether it's a database, LDAP directory, or even in-memory storage. When a user tries to authenticate, Spring Security calls this method to load their details including password, roles, and account status. This makes authentication flexible - I can switch from database authentication to LDAP without changing my security configuration.

## Advanced Security (OAuth2, JWT, Method Security)

### 8. Explain OAuth2 flow simply.
OAuth2 is an authorization framework.
1.  **User** wants to access a resource (e.g., Google Photos) via a **Client** app.
2.  **Client** redirects User to **Authorization Server** (Google).
3.  **User** logs in and grants permission.
4.  **Authorization Server** issues an **Access Token** to the Client.
5.  **Client** uses the token to access the **Resource Server** (Google Photos API).

**Explanation:** OAuth2 separates authentication from authorization, allowing secure delegated access without sharing credentials, making it ideal for third-party integrations and microservices.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain OAuth2 flow simply?
**Your Response:** OAuth2 is an authorization framework that lets applications access resources on behalf of users without sharing passwords. The flow works like this: when a user wants to use my app to access their Google Photos, my app redirects them to Google's authorization server. The user logs into Google and grants permission to my app. Google then gives my app an access token, which I can use to make API calls to Google Photos on the user's behalf. The user never shares their Google password with my app, making it much more secure than traditional credential sharing.

### 9. What is JWT and how does it work with Spring Security?
**JWT (JSON Web Token)** is a compact, URL-safe means of representing claims to be transferred between two parties.
- **Implementation:**
    1.  User logs in.
    2.  Server validates credentials and generates a signed JWT.
    3.  Server sends JWT to client.
    4.  Client sends JWT in `Authorization: Bearer <token>` header for subsequent requests.
    5.  Server validates the signature of the JWT on every request (stateless).

**Explanation:** JWTs are self-contained tokens that include user information and claims, digitally signed to prevent tampering, enabling stateless authentication in distributed systems.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is JWT and how does it work with Spring Security?
**Your Response:** JWT or JSON Web Token is a compact, self-contained token that securely carries information between parties. In Spring Security, I use JWTs for stateless authentication. When a user logs in, I validate their credentials and generate a signed JWT containing their user information and claims. The client stores this token and sends it in the Authorization header for subsequent requests. On each request, Spring Security validates the JWT signature to ensure it hasn't been tampered with, then extracts the user information. This approach is great for microservices because it eliminates the need for session state.

### 10. How do you secure specific methods (Method Security)?
Enable it with `@EnableMethodSecurity`.
Then use annotations:
- `@PreAuthorize("hasRole('ADMIN')")`: Checks before method execution.
- `@PostAuthorize`: Checks after execution.
- `@Secured("ROLE_USER")`: Legacy annotation.

**Explanation:** Method security provides fine-grained access control at the business logic level, allowing security rules to be applied directly to service methods rather than just endpoints.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure specific methods in Spring Security?
**Your Response:** I secure specific methods using Spring Security's method-level security. First I enable it with @EnableMethodSecurity in my configuration class. Then I can use annotations like @PreAuthorize to check conditions before method execution, @PostAuthorize to check after execution, or @Secured for simple role checks. The most powerful is @PreAuthorize which supports SpEL expressions - I can check roles, evaluate method parameters, or combine conditions. This lets me implement business-level security rules directly in my service layer, like only allowing users to delete their own posts.

### 11. What is CSRF and why do we disable it for REST APIs?
**CSRF (Cross-Site Request Forgery)** tricks a user into executing unwanted actions on a web application where they are authenticated.
- **Browser-based apps:** Essential to enable CSRF protection.
- **REST APIs (Stateless):** Since REST APIs usually use tokens (JWT) instead of cookies for auth, and are stateless, CSRF protection is typically disabled because there is no session to hijack in the traditional sense.

**Explanation:** CSRF protection relies on session-based authentication using cookies, while stateless REST APIs with token-based authentication are inherently protected against CSRF attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CSRF and why do we disable it for REST APIs?
**Your Response:** CSRF or Cross-Site Request Forgery is an attack where a malicious website tricks a user's browser into making unwanted requests to a site where the user is authenticated. CSRF protection works by including a token in forms and checking it on the server. However, for REST APIs that use stateless authentication like JWT tokens instead of session cookies, CSRF attacks aren't really possible because there's no session cookie to hijack. That's why I typically disable CSRF protection for stateless REST APIs - it adds overhead without providing real security benefits in that context.

### 12. What is `AuthenticationProvider`?
It is responsible for authentication logic. It validates the `Authentication` object. You can implement custom providers for different auth mechanisms (e.g., OTP, Biometric, Legacy Systems).

**Explanation:** AuthenticationProvider serves as the strategy interface for different authentication mechanisms, allowing Spring Security to support multiple authentication methods simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AuthenticationProvider?
**Your Response:** AuthenticationProvider is the interface that handles the actual authentication logic in Spring Security. It takes an Authentication object, validates the credentials, and returns a fully populated Authentication object if successful. What makes it powerful is that I can implement multiple providers to support different authentication methods - like one for username/password, another for OTP tokens, and a third for biometric authentication. Spring Security tries each provider until one succeeds, giving me great flexibility in how users can authenticate to my application.

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

**Explanation:** Password encoding uses one-way hashing algorithms with salts to protect passwords even if the database is compromised, following defense-in-depth principles.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle password encoding in Spring Security?
**Your Response:** I never store passwords in plain text - that's a critical security mistake. Instead, I use Spring Security's PasswordEncoder interface with a strong hashing algorithm like BCrypt. BCrypt is the industry standard because it automatically handles salting and has a configurable work factor that makes it resistant to brute force attacks. I define it as a bean and Spring Security automatically uses it to encode passwords when users register and verify them during login. This ensures that even if my database is compromised, attackers can't recover the original passwords.

## Expert Level Topics

### 14. How do you add a custom filter to the security chain?
You can execute a custom filter before or after a standard filter using `addFilterBefore()` or `addFilterAfter()` in the `SecurityFilterChain` configuration.
```java
http.addFilterBefore(new JwtAuthenticationFilter(), UsernamePasswordAuthenticationFilter.class);
```

**Explanation:** Custom filters allow extending Spring Security's functionality for specific requirements like JWT validation, API key authentication, or custom logging without modifying the core framework.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add a custom filter to the security chain?
**Your Response:** I add custom filters to the Spring Security chain using the addFilterBefore or addFilterAfter methods in my SecurityFilterChain configuration. For example, if I'm implementing JWT authentication, I'll add a custom JWT authentication filter before the UsernamePasswordAuthenticationFilter. This gives me precise control over where my custom logic fits in the security processing pipeline. I can also use addFilterAt to replace a standard filter entirely. This approach lets me extend Spring Security's capabilities for specific requirements like custom authentication mechanisms or additional logging.

### 15. What is Role Hierarchy?
Spring Security supports role hierarchy, where a higher role automatically includes the authorities of a lower role.
- Example: `ROLE_ADMIN > ROLE_USER`. If a user has `ROLE_ADMIN`, they automatically pass checks for `ROLE_USER`.
- Configured via `RoleHierarchy` bean.

**Explanation:** Role hierarchy reduces the complexity of authorization rules by establishing inheritance relationships between roles, making security configurations more maintainable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Role Hierarchy in Spring Security?
**Your Response:** Role hierarchy allows me to define inheritance relationships between roles in Spring Security. For example, I can configure ROLE_ADMIN to be higher than ROLE_USER, which means any user with ROLE_ADMIN automatically gets all the permissions of ROLE_USER. This eliminates the need to assign both roles to admin users and simplifies my authorization checks. I configure it using a RoleHierarchy bean where I define the hierarchy like 'ROLE_ADMIN > ROLE_USER > ROLE_GUEST'. This makes my security configuration much cleaner and easier to maintain as the application grows.

### 16. @Secured vs @PreAuthorize vs @RolesAllowed
- **`@Secured`:** Spring specific, legacy. Supports simple role checks.
- **`@PreAuthorize`:** Spring specific, powerful (SpEL support). Can check arguments, combine conditions (`hasRole('A') or hasRole('B')`).
- **`@RolesAllowed`:** JSR-250 standard (Java standard). Similar to `@Secured`.

**Explanation:** Each annotation serves different use cases - @Secured for simple role checks, @PreAuthorize for complex conditional logic, and @RolesAllowed for standards compliance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @Secured, @PreAuthorize, and @RolesAllowed?
**Your Response:** These annotations provide different levels of method security. @Secured is the older Spring-specific annotation that only supports simple role checks. @PreAuthorize is much more powerful - it supports Spring Expression Language, so I can check method parameters, combine multiple conditions, and write complex authorization logic. @RolesAllowed is the JSR-250 standard annotation, which makes it portable across different Java security frameworks. I generally prefer @PreAuthorize for its flexibility, but might use @RolesAllowed if I need standards compliance or @Secured for very simple role-based checks.

### 17. Detailed difference between CORS and CSRF?
- **CORS (Cross-Origin Resource Sharing):** A browser security feature that restricts cross-origin HTTP requests. It protects the *client* (data leakage). You configure it to *allow* other domains to access your API.
- **CSRF:** An attack vector. It protects the *server* from unauthorized state-changing actions. You configure protections to *block* malicious requests.

**Explanation:** CORS is a browser security mechanism that servers configure to allow cross-origin requests, while CSRF is an attack type that servers protect against. They address different security concerns in opposite directions.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the detailed difference between CORS and CSRF?
**Your Response:** CORS and CSRF address completely different security issues. CORS is a browser security feature that prevents web pages from making requests to different domains - it protects the client from malicious websites trying to steal data. I configure CORS to explicitly allow trusted domains to access my API. CSRF is an attack where a malicious website tricks a user's browser into making unwanted requests to a site where they're authenticated - it protects the server from unauthorized actions. I configure CSRF protection to block these malicious requests. So CORS is about allowing legitimate cross-origin requests, while CSRF is about blocking malicious ones.
