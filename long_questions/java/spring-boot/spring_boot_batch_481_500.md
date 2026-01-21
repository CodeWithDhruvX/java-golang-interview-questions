## 🔹 Section 5: Security & Authorization (481-500)

### Question 481: How do you use `@PreAuthorize` and `@PostAuthorize` annotations?

**Answer:**
(See Q181).
- `Pre`: Checks before method entry.
- `Post`: Checks after method execution (can inspect `returnObject`).

---

### Question 482: How do you enable method-level security globally in Spring Boot?

**Answer:**
(See Q181). `@EnableMethodSecurity`.

---

### Question 483: What is the purpose of `SecurityContextPersistenceFilter`?

**Answer:**
It loads the `SecurityContext` from the `SecurityContextRepository` (Session) at the start of request.
And saves it back at end of request.
Ensures Authentication persists across requests in session-based apps.

---

### Question 484: How do you define custom authentication providers?

**Answer:**
Implement `AuthenticationProvider.authenticate()`.
Register it as a Bean.
Spring Security will call it during `authManager.authenticate()`.

---

### Question 485: How do you implement LDAP authentication in Spring Boot?

**Answer:**
Add `spring-boot-starter-data-ldap`.
Config:
`auth.ldapAuthentication().userDnPatterns("uid={0},ou=people").contextSource(contextSource())`.

---

### Question 486: How do you use `BCryptPasswordEncoder` in a login system?

**Answer:**
Bean definition:
`@Bean PasswordEncoder passwordEncoder() { return new BCryptPasswordEncoder(); }`.
On Register: `repo.save(user.setPassword(encoder.encode(rawPassword)))`.
On Login: Spring Security uses it automatically to match.

---

### Question 487: How do you secure REST endpoints using JWT in Spring Boot?

**Answer:**
(See Q83).

---

### Question 488: What is CSRF and how do you handle it in APIs?

**Answer:**
(See Q84).

---

### Question 489: How do you build a custom login page with Spring Boot Security?

**Answer:**
`http.formLogin().loginPage("/login").permitAll()`.
Create a Controller for `/login` returning a Thymeleaf view.

---

### Question 490: How can you invalidate sessions after password reset?

**Answer:**
Iterate `SessionRegistry`.
Find sessions for principal.
Call `sessionInformation.expireNow()`.

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
