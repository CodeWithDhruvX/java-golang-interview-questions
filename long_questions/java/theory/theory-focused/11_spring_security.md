# Spring Security - Interview Answers

> 🎯 **Focus:** Security is scary. These answers show you understand the flow and modern practices like JWT/OAuth2.

### 1. Authentication vs Authorization?
"**Authentication (AuthN)** is 'Who are you?'. It validates credentials (username/password).
**Authorization (AuthZ)** is 'What are you allowed to do?'. It validates permissions (Roles, Scopes).

In Spring Security, the `AuthenticationManager` handles the 'Who', and the `SecurityFilterChain` (or AccessDecisionManager) handles the 'What'."

---

### 2. How does the Spring Security Filter Chain work?
"Spring Security is essentially a chain of Servlet Filters.
When a request comes in, it passes through filters like:
1. `JwtAuthenticationFilter` (Custom): Checks for a token.
2. `UsernamePasswordAuthenticationFilter`: Checks for login form data.
3. `ExceptionTranslationFilter`: Handles access denied errors.
4. `FilterSecurityInterceptor`: The final gatekeeper that decides if the request can proceed to the Controller.

Customizing security usually means inserting my own filter into this chain."

---

### 3. State(ful) vs Stateless Authentication?
"**Stateful** uses Sessions (JSESSIONID). The server remembers the user in memory. It's simple but harder to scale horizontally because you need Sticky Sessions or Session Replication.

**Stateless** uses Tokens (JWT). The server remembers nothing. The client holds the token. Every request is independent. This is the standard for Microservices because any server instance can validate the token."

---

### 4. What is a JWT (JSON Web Token)?
"It’s a digitally signed JSON object. It has three parts:
1. **Header**: Algorithm type.
2. **Payload (Claims)**: Data like User ID, Role, Expiration.
3. **Signature**: Cryptographic proof that the token hasn't been tampered with.

Since it's signed, the server trusts it without checking the database. That's why it's so fast."

---

### 5. What is OAuth2?
"It’s an authorization framework. It allows a user to grant a third-party app access to their data without sharing their password.
Think 'Log in with Google'.
Google (Provider) gives my App (Client) a token that represents the User.

I use `spring-boot-starter-oauth2-client` to implement this. It handles the redirect to Google, the callback code exchange, and fetching the user profile automatically."

---

### 6. Why disable CSRF?
"Cross-Site Request Forgery (CSRF) protection is enabled by default. It expects a rigorous token for every POST request.

This is critical for Session-based apps (monoliths with server-side templates).
But for **Stateless REST APIs** (using JWT), we usually disable it (`.csrf().disable()`).
Since we don't use cookies for auth, the browser doesn't automatically send credentials, so the CSRF attack vector doesn't exist."

---

### 7. Method Level Security (`@PreAuthorize`)?
"Instead of securing URLs in the config, I prefer securing methods directly.
I enable it with `@EnableGlobalMethodSecurity`.
Then I can use:
`@PreAuthorize("hasRole('ADMIN')")`
or even simpler logic:
`@PreAuthorize("#userId == authentication.principal.id")`

This is powerful because it secures the logic regardless of whether it's called from a Controller or an internal service."

---

### 8. `UserDetailsService`?
"It’s the core interface for loading user data. It has one method: `loadUserByUsername()`.
Spring Security calls this when a user tries to log in.
My job is to implement this interface, query my database (UserRepository), and return a `UserDetails` object containing the password and authorities (roles).
Spring then takes that object and checks if the password matches."

---

### 9. Password Encoding (`BCrypt`)?
"We never store plain text passwords. Never.
Spring provides `PasswordEncoder` interface. The standard implementation is `BCryptPasswordEncoder`.

It hashes the password with a random Salt.
When a user logs in, we don't decrypt the stored password (we can't). Instead, we hash the input password and compare the two hashes."

---

### 10. CORS (Cross-Origin Resource Sharing)?
"Browsers block AJAX requests from one domain to another (e.g., React on localhost:3000 to Spring on localhost:8080) for safety.
To allow it, the backend must send specific headers like `Access-Control-Allow-Origin`.

In Spring, I configure a `CorsConfigurationSource` bean or use `@CrossOrigin` on the controller.
Common pitfall: If you use Spring Security, you must enable CORS in the Security Config too, otherwise the Security Filter rejects the pre-flight Options request."
