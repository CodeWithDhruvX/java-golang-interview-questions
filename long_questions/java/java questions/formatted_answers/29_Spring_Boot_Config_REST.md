# 29. Spring Boot (Configuration, REST & JPA)

**Q: @ConfigurationPropertiesScan vs @EnableConfigurationProperties**
> "In the old days, you had to manually list every single config class: `@EnableConfigurationProperties(MyConfig.class, OtherConfig.class)`.
>
> Now, just add `@ConfigurationPropertiesScan` to your main class. Spring scans your package, finds any class annotated with `@ConfigurationProperties`, and registers it automatically. It makes the main class much cleaner."

**Indepth:**
> **Immutable**: A modern best practice is to use **Java Records** or Constructor Binding with `@ConfigurationProperties`. This makes your config objects immutable (final fields), which prevents accidental modification at runtime.


---

**Q: Validating Configuration Properties**
> "You can put JSR-303 annotations directly on your **Properties Class**.
>
> ```java
> @ConfigurationProperties(prefix = \"app\")
> @Validated
> public class AppProps {
>     @NotNull
>     private String name;
>
>     @Min(10)
>     private int timeout;
> }
> ```
> If the user sets `app.timeout=5`, the application **will fail to start** with a validation error. This is a fail-fast mechanism."

**Indepth:**
> **Nested Validation**: To validate nested objects (e.g., `app.database.url`), you must annotate the nested field with `@Valid`. Without it, the validator inspects the parent but skips the child object fields.


---

**Q: Profiles Groups**
> "You can group profiles so you don't have to list them all effectively.
> In `application.properties`:
> `spring.profiles.group.prod = db-prod, security-prod, cloud-prod`
>
> Now, you just start with `--spring.profiles.active=prod` and it automatically activates all three sub-profiles."

**Indepth:**
> **Activation**: You can also activate profiles conditionally based on OS, JDK version, or presence of other profiles (`!prod`) using the newer `spring.config.activate.on-profile` syntax in multi-document YAML files.


---

**Q: Config File Merging & Precedence**
> "Spring Boot merges properties files.
> A value in `application-prod.yml` **overrides** value in `application.yml`.
>
> **Precedence Order (Simplest to Strongest)**:
> 1.  `application.properties` (inside jar)
> 2.  `application.properties` (outside jar)
> 3.  Environment Variables (`SERVER_PORT=8080`)
> 4.  Command Line Arguments (`--server.port=9000`) - **Strongest**."

**Indepth:**
> **Test Properties**: Note that `@SpringBootTest(properties = "foo=bar")` or `@TestPropertySource` overrides almost everything else, designed specifically for integration testing isolation.


---

**Q: @RequestBody vs @ModelAttribute**
> "**@RequestBody** is for JSON/XML.
> It uses an `HttpMessageConverter` (like Jackson) to parse the raw body of the request into an object.
>
> "**@ModelAttribute** is for Form Data (`application/x-www-form-urlencoded`).
> It takes query parameters (`?name=John`) or form fields and binds them to a Java object setters. Used mostly in MVC web pages, not REST APIs."

**Indepth:**
> **Deserialization**: Jackson (the default JSON library) uses getters/setters or direct field access (via reflection) to populate the `@RequestBody` object. It requires a default constructor unless configured with a custom module (like `jackson-module-parameter-names`).


---

**Q: Global Exception Handling (@ControllerAdvice)**
> "Don't write try-catch blocks in every controller.
>
> creates a class annotated with `@ControllerAdvice`.
> Inside, define methods annotated with `@ExceptionHandler(MyException.class)`.
>
> When *any* controller throws that exception, this central handler catches it and returns a standard JSON error response."

**Indepth:**
> **Hierarchy**: You can refine the scope. `@ControllerAdvice` applies globally. Accessing a `@ExceptionHandler` method inside a *specific* Controller class applies only to that controller. This allows for granular error handling strategies.


---

**Q: Rate Limiting in Spring Boot**
> "Spring Boot doesn't have a built-in Rate Limiter.
> You usually use a library like **Bucket4j** or **Resilience4j**.
>
> You define a 'Bucket' (e.g., 10 tokens per minute).
> In an Interceptor or Filter, you check `bucket.tryConsume(1)`. If false, you return `429 Too Many Requests`."

**Indepth:**
> **Distributed**: If you run multiple instances of your API (microservices), a local in-memory rate limiter wont work (users can hit instance A then instance B). You need a distributed store like Redis (using Redisson) to count tokens globally.


---

**Q: ETag & Caching**
> "ETag is a hash of the response content.
> 1.  Server sends response with Header `ETag: "12345"`.
> 2.  Client requests again, sending Header `If-None-Match: "12345"`.
> 3.  Server checks: 'Has data changed? No? ok.'
> 4.  Server returns `304 Not Modified` (Empty Body).
>
> Enable it in Spring: `spring.web.resources.cache.use-last-modified=true` or use `ShallowEtagHeaderFilter`."

**Indepth:**
> **Weak vs Strong**: A "Strong ETag" means the content is byte-for-byte identical. A "Weak ETag" (`W/"123"`) means the content is semantically equivalent (maybe different formatting). Spring usually generates Weak ETags.


---

**Q: Request/Response Compression**
> "You can turn on GZIP compression with just properties:
> `server.compression.enabled=true`
> `server.compression.mime-types=text/html,application/json`
> `server.compression.min-response-size=1024`
>
> It reduces bandwidth significantly but increases CPU usage slightly."

**Indepth:**
> **Breach Attack**: Be careful with compression if you serve secrets (CSRF tokens) in the same response as user-controlled data. The BREACH attack allows attackers to guess secrets by observing compressed sizes. Disable compression for sensitive endpoints.


---

**Q: Interceptors vs Filters**
> "**Filters** (Servlet Filter) run **outside** Spring context. Good for security, logging, compression. They check the raw request.
>
> **Interceptors** (HandlerInterceptor) run **inside** Spring MVC. They have access to the *Handler* (the controller method). Good for logic like 'Is this specific user allowed to call this specific method?'."

**Indepth:**
> **Order**: Filters trigger *before* Interceptors. The chain is: Request -> Filter Chain -> DispatcherServlet -> Interceptor (preHandle) -> Controller -> Interceptor (postHandle) -> View -> Filter (response).


---

**Q: Specification API (JPA)**
> "When you have dynamic search filters (e.g., User can filter by Name OR Age OR City), writing plain methods is hard (`findByNameAndAgeAndCity...`).
>
> **Specifications** let you build queries programmatically:
> `Spec s = Spec.where(hasName("John")).and(hasAge(25));`
> `repo.findAll(s);`
> It generates the exact WHERE clause needed."

**Indepth:**
> **Criteria API**: Under the hood, Specifications use the JPA Criteria API. It is type-safe but verbose. Specifications provide a cleaner, fluent wrapper around it.


---

**Q: Entity Graph (N+1 Problem)**
> "The N+1 problem happens when you fetch a List of Orders, and then for *each* Order, Hibernate runs a separate query to fetch the Customer.
>
> **EntityGraph** fixes this eagerly.
> `@EntityGraph(attributePaths = {"customer"})`
> `List<Order> findAll();`
>
> This tells JPA: 'When you fetch Orders, do a **LEFT JOIN FETCH** on Customer strictly in one query'."

**Indepth:**
> **Projections**: Alternatively, use "Interface-based Projections" (fetching only specific columns into a DTO). It avoids the N+1 problem entirely because the query selects exactly what you need, nothing more.


---

**Q: Soft Deletes**
> "Never actually delete data (`DELETE FROM`). Business usually wants to keep history.
>
> 1.  Add a column `deleted = false`.
> 2.  Use Hibernate annotation:
>     `@SQLDelete(sql = "UPDATE user SET deleted = true WHERE id = ?")`
>     `@Where(clause = "deleted = false")`
>
> Now, `repo.delete(user)` runs an UPDATE, and `repo.findAll()` only returns active users automatically."

**Indepth:**
> **Caveat**: Hard deletes are sometimes necessary for GDPR (Right to be Forgotten). If you use Soft Deletes, you must have a separate process to permanently purge data or anonymize it upon request.

