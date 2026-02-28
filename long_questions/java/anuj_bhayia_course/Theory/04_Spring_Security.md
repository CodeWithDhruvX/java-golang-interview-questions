# Spring Security - Interview Questions and Answers

## 1. What is Spring Security and what are its core concepts?
**Answer:**
Spring Security is a powerful and highly customizable authentication and access-control (authorization) framework for Java applications. It is the de-facto standard for securing Spring-based applications.

**Core Concepts:**
1. **Authentication (Who are you?):** The process of verifying the identity of a user, device, or system. Common forms include username/password, OTP, OAuth2, and biometrics.
2. **Authorization (What are you allowed to do?):** Also known as Access Control. Once authenticated, deciding whether the principal (user) has permission to access a specific resource or execute a specific method based on their roles or authorities.
3. **Principal:** The currently authenticated user or system interacting with the application.
4. **GrantedAuthority:** A permission or role assigned to the principal (e.g., `ROLE_ADMIN`, `READ_PRIVILEGE`).
5. **SecurityContext:** The central interface where Spring Security stores the details of the currently authenticated principal (the `Authentication` object). It is usually stored in a `ThreadLocal`, making the security context available anywhere in the same thread.

## 2. Explain the core architecture/flow of Spring Security when a request comes in.
**Answer:**
Spring Security operates primarily on a chain of Servlet Filters called the **DelegatingFilterProxy** and the **FilterChainProxy**.

**The Request Flow:**
1. **DelegatingFilterProxy:** This is a standard Servlet filter registered with the application server (e.g., Tomcat). Its job is to intercept the incoming HTTP request and delegate the work to a Spring Bean named `springSecurityFilterChain`.
2. **FilterChainProxy:** This bean manages a list of Security Filter Chains. It determines which configured `SecurityFilterChain` should be applied to the current request URL.
3. **Security Filters:** The request passes through a sequence of specialized filters (e.g., `CorsFilter`, `CsrfFilter`, `UsernamePasswordAuthenticationFilter`, `BasicAuthenticationFilter`, `BearerTokenAuthenticationFilter`).
4. **Authentication Filter:** If the request hits an authentication filter (like `UsernamePasswordAuthenticationFilter` during a login attempt), the filter extracts the credentials from the request and creates an unauthenticated `Authentication` object.
5. **AuthenticationManager & AuthenticationProvider:** The filter passes the token to the `AuthenticationManager` (the central interface for authentication). The manager iterates through configured `AuthenticationProvider`s (like `DaoAuthenticationProvider` for database lookups).
6. **UserDetailsService:** The `DaoAuthenticationProvider` calls the `UserDetailsService` (usually your custom implementation connecting to a database) to fetch the user's encoded password and authorities by username.
7. **PasswordEncoder:** The provider uses a `PasswordEncoder` to verify if the raw password provided in the request matches the encoded password from the database.
8. **SecurityContextHolder:** If successful, a fully populated, authenticated `Authentication` object is created and stored in the `SecurityContextHolder`.
9. **AuthorizationFilter:** Finally, the request reaches the `AuthorizationFilter` (which replaces the older `FilterSecurityInterceptor`). It checks the `SecurityContext` against the required authorization rules for the endpoint.
10. **Controller:** If authorized, the request proceeds to the Spring MVC Controller.

## 3. How do you implement Role-Based Access Control (RBAC) in Spring Boot?
**Answer:**
RBAC ensures users can only access endpoints or execute methods they are permitted to based on their assigned roles.

**1. Configuration Level (URL-based security):**
In your `SecurityFilterChain` bean configuration (Spring Security 6+ style), you define authorization rules for specific request matchers:
```java
@Bean
public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http
        .authorizeHttpRequests(authorize -> authorize
            .requestMatchers("/public/**", "/login", "/register").permitAll()
            .requestMatchers("/admin/**").hasRole("ADMIN")
            .requestMatchers("/api/users/**").hasAnyRole("ADMIN", "USER")
            .anyRequest().authenticated()
        )
        // ... configure login, session, csrf
    return http.build();
}
```

**2. Method Level Security (Granular control):**
You enable it by adding `@EnableMethodSecurity` to a configuration class. Then you can use annotations directly on service methods or controllers:
- `@PreAuthorize("hasRole('ADMIN')")`: Evaluates the expression *before* entering the method.
- `@PostAuthorize("returnObject.owner == authentication.name")`: Evaluates the expression *after* the method executes (can check the returned object).
- `@Secured("ROLE_ADMIN")`: Older, less flexible annotation. Prefer `@PreAuthorize`.

## 4. Why is Password Encoding crucial, and how do you implement it in Spring Security?
**Answer:**
**Why crucial:** Storing plain-text passwords in a database is a massive security vulnerability. If the database is compromised, all user accounts are immediately exposed. Passwords must be hashed using a strong, slow cryptographic function.

**Implementation:**
Spring Security provides the `PasswordEncoder` interface.

1. **`BCryptPasswordEncoder` (The Standard):** This is the most common and recommended implementation. BCrypt is a hashing algorithm that incorporates a "salt" (random data added before hashing to protect against rainbow table attacks) and a "work factor" (making the hashing process intentionally slow to resist brute-force attacks).
2. **Configuration:** You define it as a bean in your security configuration:
    ```java
    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
    ```
3. **Usage - Registration:** When a new user registers, you encode their password before saving the entity:
    ```java
    user.setPassword(passwordEncoder.encode(rawPassword));
    userRepository.save(user);
    ```
4. **Usage - Login:** Spring Security's `DaoAuthenticationProvider` uses the `PasswordEncoder` bean automatically to verify (`passwordEncoder.matches(raw, encoded)`) during the login process.

## 5. What is a JWT (JSON Web Token), and how is it used for stateless authentication in Microservices?
**Answer:**
**JWT:** A JSON Web Token is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed (using a secret via HMAC or a public/private key pair like RSA).

**Structure:** A JWT consists of three parts separated by dots (`.`): `Header.Payload.Signature`.
- **Header:** Contains the algorithm used for the signature (e.g., HS256).
- **Payload:** Contains the claims (the actual data, like user ID, roles, expiration time).
- **Signature:** Used to verify that the sender of the JWT is who it says it is and to ensure that the message wasn't changed along the way.

**Usage in Stateless Auth (Microservices/REST):**
1. **Login:** User sends credentials to an Auth Service.
2. **Token Generation:** The Auth Service validates credentials, generates a JWT containing the user's roles and ID in the payload, signs it with a secret key, and returns the token to the client.
3. **Subsequent Requests:** The client stores the token (e.g., in localStorage or an HttpOnly cookie) and includes it in the `Authorization: Bearer <token>` header of every subsequent request to resource servers (microservices).
4. **Stateless Verification:** The resource server intercepts the request using a custom Spring Security filter (`JwtAuthenticationFilter`). It extracts the token, verifies the signature using the shared secret (or public key), parses the claims, and sets the Spring `SecurityContext`.
5. **No Sessions:** The server does not need to store the user's session state in memory or a database (like Redis) because all necessary information to authorize the user is contained within the token itself. This makes scaling microservices extremely easy.

## 6. How do you implement OAuth2 with Google in Spring Boot?
**Answer:**
Spring Security makes integrating third-party OAuth2 login (like Google, GitHub, Facebook) relatively straightforward using the `spring-boot-starter-oauth2-client` dependency.

**Steps:**
1. **Google Cloud Console:** Register a new application in the Google Developer Console to obtain a `Client ID` and `Client Secret`. Configure authorized Redirect URIs (e.g., `http://localhost:8080/login/oauth2/code/google`).
2. **Application Properties:** Configure the application with the obtained credentials:
    ```properties
    spring.security.oauth2.client.registration.google.client-id=YOUR_CLIENT_ID
    spring.security.oauth2.client.registration.google.client-secret=YOUR_CLIENT_SECRET
    # Spring Boot auto-configures the provider URLs for common providers like Google
    ```
3. **Security Configuration:** Enable OAuth2 login in your `SecurityFilterChain`:
    ```java
    http
        .authorizeHttpRequests(authorize -> authorize
            .anyRequest().authenticated()
        )
        .oauth2Login(Customizer.withDefaults()); // Or customize the success handler
    ```
4. **Flow:**
    - User accesses a protected resource.
    - Spring redirects them to the Google Authorization Server.
    - User logs in to Google and grants consent.
    - Google redirects back to the Spring app with an authorization code.
    - Spring Security automatically exchanges this code for an Access Token and fetches the user's profile details using the `OAuth2UserService`.
    - An `OAuth2AuthenticationToken` is placed in the `SecurityContext`.

## 7. What is Cross-Site Request Forgery (CSRF), and how does Spring Security handle it?
**Answer:**
**CSRF Attack:** CSRF is an attack that forces an end user to execute unwanted actions on a web application in which they are currently authenticated. If a user is logged into their bank, a malicious website could embed a hidden form or an image tag that sends a state-changing POST request (like transferring money) to the bank's servers. Because the user's browser automatically includes their session cookies, the bank server processes the request, believing the user intended it.

**Spring Security Handling:**
By default, Spring Security enables CSRF protection for any request that could potentially change state (POST, PUT, DELETE, PATCH). It ignores idempotent requests (GET, HEAD, OPTIONS).

**How it works (Synchronizer Token Pattern):**
1. When a user authenticates, Spring Security generates a unique, cryptographically strong, unpredictable CSRF Token.
2. This token is stored in the user's HTTP session.
3. Every state-changing form rendered by the application must include this token as a hidden field (e.g., `<input type="hidden" name="${_csrf.parameterName}" value="${_csrf.token}"/>`). Thymeleaf automatically injects this.
4. When the form is submitted, Spring Security's `CsrfFilter` intercepts the request. It compares the token received in the request against the one stored in the session.
5. If they match, the request proceeds. If they are missing or mismatched, Spring rejects the request with a 403 Forbidden error.

*Note on REST APIs:* If you are building a pure, stateless REST API that uses JWTs or Bearer tokens in the `Authorization` header instead of session cookies, CSRF attacks are generally not possible. In this specific scenario, you explicitly disable CSRF protection: `http.csrf(csrf -> csrf.disable());`.

## 8. What is `UserDetailsService` and `UserDetails` in Spring Security?
**Answer:**
These are foundational interfaces used by Spring Security for database-backed authentication.

- **`UserDetails`:** This interface represents the core user information inside Spring Security. It encapsulates details like the username, password (encoded), granted authorities (roles), and account status (is enabled, is locked, are credentials expired). You usually implement this interface on your custom `User` entity, or create a wrapper class around your entity.
## 9. How do you create Custom Request Filters in Spring Security?
**Answer:**
Often, you need to intercept a request *before* it reaches the controller to perform custom logic (e.g., extracting a specialized API key, logging request details, or modifying the response).

**Implementation:**
1. **Create the Filter:** Extend the `OncePerRequestFilter` class. This guarantees that your filter's `doFilterInternal` method is executed exactly once per HTTP request.
2. **Implement Logic:**
    ```java
    @Component
    public class ApiKeyFilter extends OncePerRequestFilter {
        @Override
        protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) 
                throws ServletException, IOException {
            
            String apiKey = request.getHeader("X-API-KEY");
            if (apiKey != null && isValid(apiKey)) {
                // Key is valid, create an Authentication token and put it in the Context
                UsernamePasswordAuthenticationToken auth = new UsernamePasswordAuthenticationToken("apiUser", null, getAuthorities());
                SecurityContextHolder.getContext().setAuthentication(auth);
            }
            // CRITICAL: Always continue the chain!
            filterChain.doFilter(request, response);
        }
    }
    ```
3. **Register the Filter:** In your `SecurityFilterChain` definition, you explicitly insert your custom filter into the chain, usually specifying its position relative to standard filters:
    ```java
    http.addFilterBefore(new ApiKeyFilter(), UsernamePasswordAuthenticationFilter.class);
    ```

## 10. What is XSS (Cross-Site Scripting), and how does Spring Security help mitigate it?
**Answer:**
**XSS Attack:** Occurs when an attacker injects malicious executable scripts (usually JavaScript) into the HTML pages viewed by other users. For example, if a user submits a blog comment containing `<script>stealCookies();</script>` and the server doesn't sanitize it, anyone viewing that comment will execute the script in their browser.

**How Spring Security helps:**
Unlike CSRF, Spring Security cannot prevent XSS completely on its own because XSS prevention heavily relies on how data is rendered in the View Layer (Frontend/Templates). However, it provides strong defense-in-depth mechanisms via HTTP Response Headers.

**Security Headers (Enabled by Default in Spring Security):**
1. **Content-Security-Policy (CSP):** The most powerful defense. You configure Spring Security to send a `Content-Security-Policy` header. This tells the browser exactly which domains are allowed to execute scripts, load images, or fetch data. Inline scripts can be completely disabled.
   `http.headers().contentSecurityPolicy("script-src 'self' https://trustedscripts.example.com; object-src 'none'");`
2. **X-XSS-Protection:** (Legacy) Instructs older browsers to stop rendering the page if they detect a reflected XSS attack.
3. **X-Content-Type-Options:** Sends `nosniff`. Prevents the browser from trying to "guess" the MIME type of a file based on its content, preventing script execution from mislabeled image or text files.
4. **Strict-Transport-Security (HSTS):** Forces to use HTTPS.
*(Note: True XSS prevention requires input validation on the server and contextual output encoding/escaping in the UI layer, like Thymeleaf or React automatically providing).*
