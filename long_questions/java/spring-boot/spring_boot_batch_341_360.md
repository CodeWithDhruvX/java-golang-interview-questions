## 🔹 Section 5: Spring Security Deep Dive (341-350)

### Question 341: How do you implement login throttling in Spring Security?

**Answer:**
**AuthenticationFailureHandler**.
Count failures in Cache/DB.
If count > 5, throw `LockedException`.
Wait for N minutes before allowing check again.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement login throttling in Spring Security?
**Your Response:** "I implement login throttling using a custom `AuthenticationFailureHandler`. I count failed login attempts in a cache or database, and when the count exceeds a threshold like 5 attempts, I throw a `LockedException`. I implement a wait period - after too many failures, the user must wait N minutes before attempting again. This prevents brute force attacks while allowing legitimate users to regain access after a timeout. I store the attempt counts with expiration to automatically reset after the wait period."

---

### Question 342: What is the difference between stateless JWT authentication and session-based authentication?

**Answer:**
(Duplicate of Q195).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between stateless JWT authentication and session-based authentication?
**Your Response:** "Stateless JWT authentication stores nothing on the server - the JWT token contains all user information and is cryptographically signed. Session-based authentication stores session IDs on the server and sends cookies to clients. JWT is better for distributed systems and APIs since any server can validate the token. Sessions are simpler for traditional web apps but require session sharing for horizontal scaling. I choose JWT for stateless APIs and microservices, and sessions for monolithic web applications."

---

### Question 343: How do you secure static resources like CSS or JS files?

**Answer:**
`http.authorizeHttpRequests().requestMatchers("/static/**", "/css/**", "/js/**").permitAll()`.
Ensure these are public.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure static resources like CSS or JS files?
**Your Response:** "I secure static resources using Spring Security's `authorizeHttpRequests()` configuration. I use `requestMatchers('/static/**', '/css/**', '/js/**').permitAll()` to ensure these resources are publicly accessible. Static resources don't contain sensitive information and need to be available for the UI to function properly. I typically secure all API endpoints while keeping static resources public, as they're meant to be served directly to browsers without authentication."

---

### Question 344: How to protect APIs using OAuth2 with custom token validators?

**Answer:**
Customize `JwtDecoder`.
`NimbusJwtDecoder.withJwkSetUri(uri).jwtProcessor(myProcessor).build()`.
Add custom validation logic (e.g., check `aud` claim or blacklist check).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to protect APIs using OAuth2 with custom token validators?
**Your Response:** "I protect APIs with custom OAuth2 token validation by customizing the `JwtDecoder`. I use `NimbusJwtDecoder.withJwkSetUri(uri).jwtProcessor(myProcessor).build()` to create a decoder with custom validation logic. I can add checks for the `aud` claim, implement token blacklisting, or validate custom claims. This approach gives me fine-grained control over token validation beyond the standard JWT signature and expiration checks, allowing me to implement organization-specific security requirements."

---

### Question 345: What is the use of `SecurityContextHolder` in Spring Security?

**Answer:**
ThreadLocal storage for `Authentication` object.
`SecurityContextHolder.getContext().getAuthentication()`.
How you access "Who is logged in" anywhere in code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `SecurityContextHolder` in Spring Security?
**Your Response:** "`SecurityContextHolder` provides ThreadLocal storage for the current `Authentication` object. I use `SecurityContextHolder.getContext().getAuthentication()` anywhere in my code to access the currently authenticated user. This is how I can check 'who is logged in' throughout the application. The context is automatically populated by Spring Security filters and cleared at the end of each request. This ThreadLocal approach makes authentication information readily available without passing it through method parameters."

---

### Question 346: How do you configure `SecurityFilterChain` with multiple chains?

**Answer:**
Define multiple Beans of type `SecurityFilterChain` with different `@Order`.
Chain 1 (Order 1): `requestMatchers("/api/**")` -> Stateless (JWT).
Chain 2 (Order 2): `requestMatchers("/admin/**")` -> Stateful (Form Login).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure `SecurityFilterChain` with multiple chains?
**Your Response:** "I configure multiple security filter chains by defining multiple beans of type `SecurityFilterChain` with different `@Order` annotations. For example, Chain 1 with order 1 handles `/api/**` requests with stateless JWT authentication, while Chain 2 with order 2 handles `/admin/**` requests with stateful form login. The order determines which chain processes which requests first. This allows me to apply different security strategies to different parts of the application within the same Spring Boot application."

---

### Question 347: What are antMatchers vs mvcMatchers in Spring Security configuration?

**Answer:**
- **AntMatchers:** Matches URL patterns (`/path/*`).
- **MvcMatchers:** Considers Spring MVC mapping rules (Handles optional trailing slashes). **Preferred** for Spring MVC apps as it is safer.
Note: In Spring Security 6, both are replaced by `requestMatchers`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are antMatchers vs mvcMatchers in Spring Security configuration?
**Your Response:** "AntMatchers match URL patterns using Ant-style patterns like `/path/*`, while MvcMatchers consider Spring MVC mapping rules and handle optional trailing slashes. MvcMatchers are safer for Spring MVC apps because they understand how Spring MVC maps URLs. In Spring Security 6, both have been replaced by `requestMatchers` which provides the best of both approaches. I use `requestMatchers` in modern applications as it's the current recommended approach and provides more accurate URL matching."

---

### Question 348: How do you secure actuator endpoints conditionally?

**Answer:**
(Duplicate of Q188).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure actuator endpoints conditionally?
**Your Response:** "I secure actuator endpoints conditionally using Spring Security's `requestMatchers` with conditions based on profiles or environment variables. For example, I might secure all endpoints in production with `requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole('ADMIN')` but leave them open in development profiles. I can also secure specific endpoints individually based on their sensitivity. This conditional security ensures that sensitive management endpoints are protected in production while remaining accessible during development."

---

### Question 349: How do you implement permission-based access (fine-grained auth)?

**Answer:**
Custom Permission Evaluator.
`@PreAuthorize("hasPermission(#id, 'Project', 'read')")`.
Implement `PermissionEvaluator` interface to check DB if user can read Project ID.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement permission-based access (fine-grained auth)?
**Your Response:** "I implement fine-grained permission-based access using a custom `PermissionEvaluator`. I use annotations like `@PreAuthorize('hasPermission(#id, 'Project', 'read')')` where I implement the `PermissionEvaluator` interface to check the database if the user has permission to read the specific project. This approach allows me to implement complex business rules for authorization that go beyond simple role checks. I can check ownership, organizational hierarchy, or any custom permission logic in the evaluator."

---

### Question 350: How to log user authentication/authorization events in Spring Security?

**Answer:**
Publish `AuthenticationSuccessEvent` or `AuthorizationFailureEvent`.
Create `ApplicationListener<AuthenticationSuccessEvent>` to log logins.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to log user authentication/authorization events in Spring Security?
**Your Response:** "I log authentication and authorization events by publishing and listening to Spring Security events. I create an `ApplicationListener<AuthenticationSuccessEvent>` to log successful logins, and `ApplicationListener<AuthenticationFailureEvent>` for failed attempts. For authorization failures, I listen to `AuthorizationFailureEvent`. These events provide rich context about the authentication attempt, including username, IP address, and failure reasons. This audit trail is essential for security monitoring and compliance requirements."

## 🔹 Section 6: Spring Boot CLI, DevTools & Utilities (351-360)

### Question 351: What is Spring Boot CLI, and when should you use it?

**Answer:**
Command Line Interface tool.
Allows running Groovy scripts (`spring run app.groovy`).
Good for rapid prototyping or scripts. Not commonly used for Production enterprise apps.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Boot CLI, and when should you use it?
**Your Response:** "Spring Boot CLI is a command-line tool for running Groovy scripts with Spring Boot auto-configuration. I use it for rapid prototyping or quick scripts where I don't want to set up a full project structure. I can write a simple Groovy file with `@RestController` and run it directly with `spring run app.groovy`. While it's great for experimentation and quick demos, I typically don't use it for production enterprise applications where I prefer the structure and type safety of Java projects."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you install and run Groovy scripts with Spring Boot CLI?
**Your Response:** "I install Spring Boot CLI using SDKMAN with `sdk install springboot`. Then I write Groovy scripts with Spring annotations like `@RestController` and `@GetMapping`. For example, a simple web service can be written in a single Groovy file. I run it with `spring run app.groovy`, and Spring Boot CLI automatically detects the annotations and starts an embedded server. This approach is incredibly fast for creating simple services or prototypes without the overhead of a full project setup."

---

### Question 353: What are live reload features provided by Spring Boot DevTools?

**Answer:**
(See Q18). Hot restart + LiveReload browser trigger.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are live reload features provided by Spring Boot DevTools?
**Your Response:** "Spring Boot DevTools provides two main live reload features. Hot restart automatically restarts the application when I change class files, which is much faster than a full restart. LiveReload automatically refreshes the browser when I change static resources like templates or CSS. These features dramatically speed up development by eliminating the need to manually restart the application and refresh the browser after each change. DevTools detects changes and applies them automatically, making the development experience much more productive."

---

### Question 354: How do you exclude DevTools from production environments?

**Answer:**
By default, if you package as JAR, DevTools is **disabled** automatically.
Only active when running from IDE or `mvn spring-boot:run`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you exclude DevTools from production environments?
**Your Response:** "DevTools is automatically excluded from production environments by default. When I package the application as a JAR, DevTools is disabled automatically. It's only active when running from an IDE or using `mvn spring-boot:run`. This automatic exclusion ensures that development tools don't accidentally end up in production builds, where they could be a security risk or performance issue. I don't need any special configuration to exclude DevTools - Spring Boot handles this automatically."

---

### Question 355: How do you auto-open the browser on app startup using DevTools?

**Answer:**
Not a native feature of DevTools (it reloads browser, doesn't open it).
You can write a `EventListener` for `ApplicationReadyEvent` that runs `Runtime.exec("open http://localhost:8080")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you auto-open the browser on app startup using DevTools?
**Your Response:** "DevTools doesn't natively auto-open the browser - it only reloads the browser when the application restarts. To auto-open the browser, I write a custom `EventListener` for `ApplicationReadyEvent` that executes `Runtime.exec('open http://localhost:8080')` on macOS or similar commands for other operating systems. This event fires when the application is fully started, so the browser opens only after the application is ready to serve requests. This is a nice convenience for development environments."

---

### Question 356: How do you use Spring Initializr CLI for generating projects?

**Answer:**
`spring init --dependencies=web,data-jpa my-project.zip`.
`unzip my-project.zip`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Spring Initializr CLI for generating projects?
**Your Response:** "I use the Spring Initializr CLI to generate projects from the command line with `spring init --dependencies=web,data-jpa my-project.zip`. This creates a complete Spring Boot project structure with the specified dependencies. After unzipping the project, I have a ready-to-run application with Maven or Gradle build files, main class, and test structure. This approach is perfect for quickly scaffolding new projects or for scripting project creation in automated workflows."

---

### Question 357: How can you create custom starters for internal libraries?

**Answer:**
1.  Create project `acme-spring-boot-starter`.
2.  Include `acme-spring-boot-autoconfigure`.
3.  Write configuration and register in `spring.factories`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you create custom starters for internal libraries?
**Your Response:** "I create custom starters by following a specific structure. I create a project named `acme-spring-boot-starter` that includes the auto-configuration module `acme-spring-boot-autoconfigure`. I write the configuration classes and register them in `META-INF/spring.factories` so Spring Boot can discover them. The starter project provides the dependencies, while the auto-configuration module provides the actual configuration logic. This approach allows teams to create reusable, convention-based configuration for internal libraries just like Spring Boot's official starters."

---

### Question 358: What is the purpose of `.spring-boot-devtools.properties`?

**Answer:**
Global configuration for DevTools on your machine (e.g., inside `~/.spring-boot-devtools.properties`).
Can set `spring.devtools.restart.trigger-file` globally.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `.spring-boot-devtools.properties`?
**Your Response:** "`.spring-boot-devtools.properties` provides global configuration for DevTools on my machine. I place this file in my home directory (`~/.spring-boot-devtools.properties`) to set DevTools settings that apply to all Spring Boot projects I work on. For example, I can set `spring.devtools.restart.trigger-file` globally to specify a file that triggers restarts when changed. This is useful for maintaining consistent DevTools behavior across multiple projects without having to configure each project individually."

---

### Question 359: How do you define global logging formats via CLI or YAML?

**Answer:**
`logging.pattern.console=%d %-5level %logger : %msg%n`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define global logging formats via CLI or YAML?
**Your Response:** "I define global logging formats using properties like `logging.pattern.console=%d %-5level %logger : %msg%n`. This pattern controls how log messages are formatted in the console, including timestamp, log level, logger name, and message. I can also set `logging.pattern.file` for file logging. These patterns use Logback's conversion specifiers to control the exact format. Consistent logging formats are important for log aggregation and analysis tools, so I typically standardize these patterns across all my applications."

---

### Question 360: How to quickly prototype apps using CLI and embedded templates?

**Answer:**
Use `@Grab` in Groovy script to pull dependencies.
`@Grab("spring-boot-starter-thymeleaf")`.
Return "hello" (Template name).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to quickly prototype apps using CLI and embedded templates?
**Your Response:** "I quickly prototype apps using Spring Boot CLI with `@Grab` annotations to pull dependencies. For example, `@Grab('spring-boot-starter-thymeleaf')` automatically adds Thymeleaf support. I can create a simple Groovy script with a controller that returns 'hello', and Spring Boot CLI automatically resolves this to a Thymeleaf template. This approach lets me prototype complete web applications in a single file without setting up a full project structure, making it perfect for proof-of-concept work or rapid experimentation."

---
