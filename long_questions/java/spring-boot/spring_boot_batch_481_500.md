## 🔹 Section 5: Security & Authorization (481-500)

### Question 481: How do you use `@PreAuthorize` and `@PostAuthorize` annotations?

**Answer:**
(See Q181).
- `Pre`: Checks before method entry.
- `Post`: Checks after method execution (can inspect `returnObject`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `@PreAuthorize` and `@PostAuthorize` annotations?
**Your Response:** "I use `@PreAuthorize` to check permissions before method execution and `@PostAuthorize` to check after method execution. The pre-check is perfect for preventing unauthorized access, while the post-check can inspect the return object to implement fine-grained security. For example, `@PreAuthorize('hasRole('ADMIN')')` checks before execution, while `@PostAuthorize('returnObject.owner == authentication.name')` checks ownership after the method returns. This combination provides comprehensive security control at the method level."

---

### Question 482: How do you enable method-level security globally in Spring Boot?

**Answer:**
(See Q181). `@EnableMethodSecurity`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enable method-level security globally in Spring Boot?
**Your Response:** "I enable method-level security globally by adding `@EnableMethodSecurity` to my configuration class. This annotation activates support for security annotations like `@PreAuthorize`, `@PostAuthorize`, and `@Secured` on any bean method in the application. I can also configure pre-post annotations and JSR-250 support through parameters. This global approach ensures that all my security annotations are processed consistently across the entire application without needing additional configuration."

---

### Question 483: What is the purpose of `SecurityContextPersistenceFilter`?

**Answer:**
It loads the `SecurityContext` from the `SecurityContextRepository` (Session) at the start of request.
And saves it back at end of request.
Ensures Authentication persists across requests in session-based apps.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `SecurityContextPersistenceFilter`?
**Your Response:** "`SecurityContextPersistenceFilter` manages the SecurityContext throughout the request lifecycle. At the start of each request, it loads the SecurityContext from the SecurityContextRepository (typically the session). At the end of the request, it saves any changes back to the repository. This ensures that authentication information persists across multiple requests in session-based applications. The filter is essential for maintaining the user's authentication state throughout their session with the application."

---

### Question 484: How do you define custom authentication providers?

**Answer:**
Implement `AuthenticationProvider.authenticate()`.
Register it as a Bean.
Spring Security will call it during `authManager.authenticate()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define custom authentication providers?
**Your Response:** "I define custom authentication providers by implementing the `AuthenticationProvider` interface and its `authenticate()` method. I register this implementation as a bean, and Spring Security automatically calls it during the authentication process. This allows me to implement custom authentication logic beyond the standard username/password approach, such as integrating with external authentication systems, implementing custom password validation, or handling multi-factor authentication. The provider interface gives me full control over how authentication is performed."

---

### Question 485: How do you implement LDAP authentication in Spring Boot?

**Answer:**
Add `spring-boot-starter-data-ldap`.
Config:
`auth.ldapAuthentication().userDnPatterns("uid={0},ou=people").contextSource(contextSource())`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement LDAP authentication in Spring Boot?
**Your Response:** "I implement LDAP authentication by adding the `spring-boot-starter-data-ldap` dependency and configuring the LDAP authentication in my security configuration. I use `auth.ldapAuthentication().userDnPatterns('uid={0},ou=people').contextSource(contextSource())` to define how users are found in the LDAP directory. Spring Security handles the LDAP connection, authentication, and user details loading automatically. This approach allows me to integrate with corporate LDAP directories for centralized user management."

---

### Question 486: How do you use `BCryptPasswordEncoder` in a login system?

**Answer:**
Bean definition:
`@Bean PasswordEncoder passwordEncoder() { return new BCryptPasswordEncoder(); }`.
On Register: `repo.save(user.setPassword(encoder.encode(rawPassword)))`.
On Login: Spring Security uses it automatically to match.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `BCryptPasswordEncoder` in a login system?
**Your Response:** "I use `BCryptPasswordEncoder` by defining it as a bean with `@Bean PasswordEncoder passwordEncoder() { return new BCryptPasswordEncoder(); }`. During user registration, I encode the password with `encoder.encode(rawPassword)` before storing it. During login, Spring Security automatically uses the same encoder to compare the entered password with the stored hash. BCrypt provides strong, one-way hashing with salt, making passwords secure against rainbow table attacks. This is the recommended approach for password security in Spring applications."

---

### Question 487: How do you secure REST endpoints using JWT in Spring Boot?

**Answer:**
(See Q83).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure REST endpoints using JWT in Spring Boot?
**Your Response:** "I secure REST endpoints with JWT by implementing a custom authentication filter that extracts JWT tokens from the Authorization header. The filter validates the token signature and extracts user claims, which are used to create an Authentication object. I configure Spring Security to use this filter and disable session management. This stateless approach allows any service that can validate the JWT to authenticate users, making it perfect for microservice architectures and distributed systems."

---

### Question 488: What is CSRF and how do you handle it in APIs?

**Answer:**
(See Q84).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CSRF and how do you handle it in APIs?
**Your Response:** "CSRF (Cross-Site Request Forgery) is an attack that forces authenticated users to execute unwanted actions on web applications. For REST APIs that use stateless authentication like JWT, I typically disable CSRF protection since the API doesn't rely on browser sessions or cookies. For traditional web applications with session-based authentication, I enable CSRF protection which Spring Security provides automatically. The key is to disable CSRF for stateless APIs while keeping it enabled for stateful web applications to prevent CSRF attacks."

---

### Question 489: How do you build a custom login page with Spring Boot Security?

**Answer:**
`http.formLogin().loginPage("/login").permitAll()`.
Create a Controller for `/login` returning a Thymeleaf view.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a custom login page with Spring Boot Security?
**Your Response:** "I build a custom login page by configuring Spring Security with `http.formLogin().loginPage('/login').permitAll()` and creating a controller that returns a Thymeleaf view for the `/login` endpoint. The login form should POST to the default `/login` processing URL with username and password parameters. Spring Security handles the authentication processing automatically. This approach gives me full control over the login page design while leveraging Spring Security's robust authentication backend."

---

### Question 490: How can you invalidate sessions after password reset?

**Answer:**
Iterate `SessionRegistry`.
Find sessions for principal.
Call `sessionInformation.expireNow()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you invalidate sessions after password reset?
**Your Response:** "I invalidate sessions after password reset by iterating through the `SessionRegistry` to find all sessions for a specific principal. For each session found, I call `sessionInformation.expireNow()` to immediately invalidate it. This ensures that when a user resets their password, all existing sessions are invalidated, forcing them to log in again with the new password. This security measure prevents unauthorized access from existing sessions after a password change."

---

### Question 491: How do you use `@PreAuthorize` and `@PostAuthorize` annotations?

**Answer:**
(Duplicate of 481).

---

### Question 492: How do you enable method-level security globally in Spring Boot?

**Answer:**
(Duplicate of 482).

---

### Question 493: What is the purpose of `SecurityContextPersistenceFilter`?

**Answer:**
(Duplicate of 483).

---

### Question 494: How do you define custom authentication providers?

**Answer:**
(Duplicate of 484).

---

### Question 495: How do you implement LDAP authentication in Spring Boot?

**Answer:**
(Duplicate of 485).

---

### Question 496: How do you use `BCryptPasswordEncoder` in a login system?

**Answer:**
(Duplicate of 486).

---

### Question 497: How do you secure REST endpoints using JWT in Spring Boot?

**Answer:**
(Duplicate of 487).

---

### Question 498: What is CSRF and how do you handle it in APIs?

**Answer:**
(Duplicate of 488).

---

### Question 499: How do you build a custom login page with Spring Boot Security?

**Answer:**
(Duplicate of 489).

---

### Question 500: How can you invalidate sessions after password reset?

**Answer:**
(Duplicate of 490).

---
