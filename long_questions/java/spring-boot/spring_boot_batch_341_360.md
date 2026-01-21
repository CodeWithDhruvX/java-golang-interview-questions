## 🔹 Section 5: Spring Security Deep Dive (341-350)

### Question 341: How do you implement login throttling in Spring Security?

**Answer:**
**AuthenticationFailureHandler**.
Count failures in Cache/DB.
If count > 5, throw `LockedException`.
Wait for N minutes before allowing check again.

---

### Question 342: What is the difference between stateless JWT authentication and session-based authentication?

**Answer:**
(Duplicate of Q195).

---

### Question 343: How do you secure static resources like CSS or JS files?

**Answer:**
`http.authorizeHttpRequests().requestMatchers("/static/**", "/css/**", "/js/**").permitAll()`.
Ensure these are public.

---

### Question 344: How to protect APIs using OAuth2 with custom token validators?

**Answer:**
Customize `JwtDecoder`.
`NimbusJwtDecoder.withJwkSetUri(uri).jwtProcessor(myProcessor).build()`.
Add custom validation logic (e.g., check `aud` claim or blacklist check).

---

### Question 345: What is the use of `SecurityContextHolder` in Spring Security?

**Answer:**
ThreadLocal storage for `Authentication` object.
`SecurityContextHolder.getContext().getAuthentication()`.
How you access "Who is logged in" anywhere in code.

---

### Question 346: How do you configure `SecurityFilterChain` with multiple chains?

**Answer:**
Define multiple Beans of type `SecurityFilterChain` with different `@Order`.
Chain 1 (Order 1): `requestMatchers("/api/**")` -> Stateless (JWT).
Chain 2 (Order 2): `requestMatchers("/admin/**")` -> Stateful (Form Login).

---

### Question 347: What are antMatchers vs mvcMatchers in Spring Security configuration?

**Answer:**
- **AntMatchers:** Matches URL patterns (`/path/*`).
- **MvcMatchers:** Considers Spring MVC mapping rules (Handles optional trailing slashes). **Preferred** for Spring MVC apps as it is safer.
Note: In Spring Security 6, both are replaced by `requestMatchers`.

---

### Question 348: How do you secure actuator endpoints conditionally?

**Answer:**
(Duplicate of Q188).

---

### Question 349: How do you implement permission-based access (fine-grained auth)?

**Answer:**
Custom Permission Evaluator.
`@PreAuthorize("hasPermission(#id, 'Project', 'read')")`.
Implement `PermissionEvaluator` interface to check DB if user can read Project ID.

---

### Question 350: How to log user authentication/authorization events in Spring Security?

**Answer:**
Publish `AuthenticationSuccessEvent` or `AuthorizationFailureEvent`.
Create `ApplicationListener<AuthenticationSuccessEvent>` to log logins.

## 🔹 Section 6: Spring Boot CLI, DevTools & Utilities (351-360)

### Question 351: What is Spring Boot CLI, and when should you use it?

**Answer:**
Command Line Interface tool.
Allows running Groovy scripts (`spring run app.groovy`).
Good for rapid prototyping or scripts. Not commonly used for Production enterprise apps.

---

### Question 352: How do you install and run Groovy scripts with Spring Boot CLI?

**Answer:**
Install GVM (SDKMan). `sdk install springboot`.
Write `app.groovy`:
```groovy
@RestController
class App {
    @GetMapping("/") def home() { "Hello" }
}
```
Run: `spring run app.groovy`.

---

### Question 353: What are live reload features provided by Spring Boot DevTools?

**Answer:**
(See Q18). Hot restart + LiveReload browser trigger.

---

### Question 354: How do you exclude DevTools from production environments?

**Answer:**
By default, if you package as JAR, DevTools is **disabled** automatically.
Only active when running from IDE or `mvn spring-boot:run`.

---

### Question 355: How do you auto-open the browser on app startup using DevTools?

**Answer:**
Not a native feature of DevTools (it reloads browser, doesn't open it).
You can write a `EventListener` for `ApplicationReadyEvent` that runs `Runtime.exec("open http://localhost:8080")`.

---

### Question 356: How do you use Spring Initializr CLI for generating projects?

**Answer:**
`spring init --dependencies=web,data-jpa my-project.zip`.
`unzip my-project.zip`.

---

### Question 357: How can you create custom starters for internal libraries?

**Answer:**
1.  Create project `acme-spring-boot-starter`.
2.  Include `acme-spring-boot-autoconfigure`.
3.  Write configuration and register in `spring.factories`.

---

### Question 358: What is the purpose of `.spring-boot-devtools.properties`?

**Answer:**
Global configuration for DevTools on your machine (e.g., inside `~/.spring-boot-devtools.properties`).
Can set `spring.devtools.restart.trigger-file` globally.

---

### Question 359: How do you define global logging formats via CLI or YAML?

**Answer:**
`logging.pattern.console=%d %-5level %logger : %msg%n`.

---

### Question 360: How to quickly prototype apps using CLI and embedded templates?

**Answer:**
Use `@Grab` in Groovy script to pull dependencies.
`@Grab("spring-boot-starter-thymeleaf")`.
Return "hello" (Template name).

---
