## 🔹 Section 3: REST API Internals (141-160)

### Question 141: How does Spring Boot handle parameter binding in REST endpoints?

**Answer:**
It uses `HandlerMethodArgumentResolver` strategies.
- Looks at annotation (`@RequestParam`, `@PathVariable`).
- Uses reflection to convert String (from URL) to Target Type (Int, Date, Enum).
- Uses `WebDataBinder` for type conversion.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot handle parameter binding in REST endpoints?
**Your Response:** "Spring Boot uses a sophisticated parameter binding system based on `HandlerMethodArgumentResolver` strategies. It examines the annotations on method parameters like `@RequestParam` or `@PathVariable`, then uses reflection to convert string values from URLs to the target Java types like integers, dates, or enums. The `WebDataBinder` handles the actual type conversion and formatting. This system is highly extensible - I can create custom argument resolvers for complex binding scenarios. The beauty is that this all happens automatically, so I can simply declare method parameters with the right annotations and Spring Boot handles the conversion and binding behind the scenes."

---

### Question 142: What is the difference between `HttpMessageConverter` and `ObjectMapper`?

**Answer:**
- **ObjectMapper (Jackson):** The low-level library that actually converts Java Object <-> JSON.
- **HttpMessageConverter:** The Spring abstraction that selects *which* library to use (Jackson, Gson) based on Content-Type header. It delegates the work to ObjectMapper.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `HttpMessageConverter` and `ObjectMapper`?
**Your Response:** "`ObjectMapper` is the low-level Jackson library that actually converts Java objects to JSON and vice versa. `HttpMessageConverter` is Spring's higher-level abstraction that selects which library to use based on the Content-Type header. When a request comes in with `application/json`, Spring chooses the appropriate `HttpMessageConverter` which then delegates the actual JSON conversion to `ObjectMapper`. This separation allows Spring to support multiple message formats - JSON, XML, CSV - by using different converters while keeping the same programming model. `HttpMessageConverter` is the selector, `ObjectMapper` is the worker."

---

### Question 143: How do you control serialization and deserialization of JSON in Spring Boot?

**Answer:**
1.  **Annotations:** `@JsonProperty("name")`, `@JsonIgnore`.
2.  **Config:** `spring.jackson.date-format`, `spring.jackson.property-naming-strategy`.
3.  **Customizer:** Register a `Jackson2ObjectMapperBuilderCustomizer` bean to tweak the global Mapper.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you control serialization and deserialization of JSON in Spring Boot?
**Your Response:** "I have multiple levels of control over JSON serialization. At the field level, I use annotations like `@JsonProperty('name')` to customize field names or `@JsonIgnore` to exclude fields. For global configuration, I use Spring Boot properties like `spring.jackson.date-format` or `spring.jackson.property-naming-strategy`. For fine-grained control, I can register a `Jackson2ObjectMapperBuilderCustomizer` bean to customize the global ObjectMapper - this lets me configure custom serializers, date formats, or naming strategies programmatically. This layered approach gives me complete control over how my objects are serialized to JSON."

---

### Question 144: How do you return paginated responses with metadata?

**Answer:**
Return a `Page<T>` object (from Spring Data).
Spring Boot automatically serializes this into JSON containing:
- `content`: [List of items]
- `pageable`: { pageNumber, pageSize }
- `totalPages`, `totalElements`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you return paginated responses with metadata?
**Your Response:** "Spring Data makes pagination incredibly simple. I return a `Page<T>` object from my repository, and Spring Boot automatically serializes it to JSON with rich metadata. The response includes the actual data in a `content` array, pagination details in `pageable` with page number and size, and overall information like `totalPages` and `totalElements`. This gives clients everything they need to implement pagination controls - current page, total pages, total records, and the actual data. The best part is that this happens automatically - I just need to accept a `Pageable` parameter in my controller and return the Page object."

---

### Question 145: What is the difference between `@ModelAttribute` and `@RequestBody`?

**Answer:**
- **`@RequestBody`:** Consumes the HTTP Body (usually JSON/XML). Uses MessageConverters.
- **`@ModelAttribute`:** Binds Query Parameters or Form Data (`application/x-www-form-urlencoded`) to a Java Bean. Used in MVC Views usually.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `@ModelAttribute` and `@RequestBody`?
**Your Response:** "These annotations handle different parts of HTTP requests. `@RequestBody` consumes the entire HTTP request body, typically JSON or XML, and uses MessageConverters to deserialize it into a Java object. `@ModelAttribute` binds query parameters or form data from `application/x-www-form-urlencoded` requests to a Java bean. I use `@RequestBody` for REST APIs that accept JSON payloads, while `@ModelAttribute` is more common in traditional MVC applications that handle HTML form submissions. The key difference is that `@RequestBody` works with the request body, while `@ModelAttribute` works with form parameters."

---

### Question 146: How to validate request parameters in Spring Boot REST API?

**Answer:**
1.  Add `spring-boot-starter-validation`.
2.  Annotate Controller Class with `@Validated`.
3.  Annotate Param: `getItems(@RequestParam @Min(1) int count)`.
Throws `ConstraintViolationException` if invalid.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to validate request parameters in Spring Boot REST API?
**Your Response:** "To validate request parameters, I add the `spring-boot-starter-validation` dependency and annotate my controller class with `@Validated`. Then I can apply validation annotations directly to parameters, like `getItems(@RequestParam @Min(1) int count)`. If validation fails, Spring throws a `ConstraintViolationException`. I typically handle this with a global exception handler using `@ControllerAdvice` to return clean error responses. This approach gives me declarative validation with automatic error handling, keeping my controller logic clean while ensuring data integrity."

---

### Question 147: How does Spring Boot support HATEOAS APIs?

**Answer:**
Dependency: `spring-boot-starter-hateoas`.
Wraps response (DTO) in `EntityModel<T>`.
Adds links: `model.add(linkTo(methodOn(Controller.class).get(id)).withSelfRel())`.
Response includes `_links` section (Hypermedia).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot support HATEOAS APIs?
**Your Response:** "Spring Boot supports HATEOAS through the `spring-boot-starter-hateoas` dependency. I wrap my response DTOs in `EntityModel<T>` and add hypermedia links using methods like `linkTo(methodOn(Controller.class).get(id)).withSelfRel()`. The response automatically includes an `_links` section with related resource URLs. This makes my API more discoverable - clients can navigate the API by following links rather than hardcoding URLs. HATEOAS transforms my REST API from a simple data service into a true hypermedia system where the API itself guides clients to related resources and actions."

---

### Question 148: How do you apply global validation error handling?

**Answer:**
Use `@ControllerAdvice`.
Catch `MethodArgumentNotValidException` (triggered by `@Valid` on `@RequestBody`).
Extract FieldErrors and return a clean JSON error map.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you apply global validation error handling?
**Your Response:** "I implement global validation error handling using `@ControllerAdvice`. I create exception handler methods that catch `MethodArgumentNotValidException`, which is triggered when `@Valid` validation fails on `@RequestBody` parameters. In the handler, I extract the field errors and build a clean, structured JSON error response that clients can easily understand and display. This approach centralizes error handling across all controllers, ensuring consistent error responses and keeping my controller methods clean - they don't need try-catch blocks for validation errors."

---

### Question 149: What is the use of `BindingResult` in controller methods?

**Answer:**
Used immediately after `@Valid` argument.
`func(@Valid @RequestBody User user, BindingResult result)`.
If validation fails, Spring does NOT throw exception. Instead, it populates `result`. You can check `result.hasErrors()` and handle logic manually.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `BindingResult` in controller methods?
**Your Response:** "`BindingResult` gives me manual control over validation error handling. When I place it immediately after a `@Valid` annotated parameter like `func(@Valid @RequestBody User user, BindingResult result)`, Spring doesn't throw an exception if validation fails. Instead, it populates the `BindingResult` with all validation errors. I can then check `result.hasErrors()` and handle the validation logic manually - maybe returning different error responses or performing additional validation. This is useful when I need custom validation error handling beyond the default exception-based approach."

---

### Question 150: How do you apply rate limiting to Spring Boot REST endpoints?

**Answer:**
Spring Boot doesn't have built-in Rate Limiting.
1.  **Bucket4j:** Library providing Token Bucket algorithm.
2.  **Resilience4j:** RateLimiter module.
3.  **Redis:** Implement custom Lua script or use API Gateway (Zuul/Spring Cloud Gateway).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you apply rate limiting to Spring Boot REST endpoints?
**Your Response:** "Spring Boot doesn't have built-in rate limiting, but I have several excellent options. Bucket4j provides a token bucket algorithm implementation that's easy to integrate. Resilience4j offers a RateLimiter module as part of its resilience patterns. For distributed systems, I can use Redis with custom Lua scripts or implement rate limiting at the API gateway level using Spring Cloud Gateway. Each approach has different trade-offs - Bucket4j is great for single-instance rate limiting, Redis-based solutions work across multiple instances, and gateway-level rate limiting provides centralized control. I choose based on my deployment architecture and requirements."

---

### Question 151: How does Spring Boot handle parameter binding in REST endpoints?

**Answer:**
(Duplicate of 141). See 141.

---

### Question 152: What is the difference between `@ControllerAdvice` and `@ExceptionHandler`?

**Answer:**
- **`@ExceptionHandler`:** Handles exceptions for a *single* controller (unless in Base Controller).
- **`@ControllerAdvice`:** Global. Contains `@ExceptionHandler` methods that apply to *all* controllers.

---

### Question 153: What is the role of `@CrossOrigin` in Spring Boot?

**Answer:**
(See Q51). Enables CORS for specific processing methods/controllers via Annotation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of `@CrossOrigin` in Spring Boot?
**Your Response:** "`@CrossOrigin` is the annotation-based approach to enable CORS for specific controllers or methods. While I can configure CORS globally through `WebMvcConfigurer`, `@CrossOrigin` lets me enable CORS for individual endpoints with fine-grained control. I can specify exactly which origins are allowed, what HTTP methods are permitted, and whether credentials are supported. This is useful when different endpoints have different CORS policies - some might be public while others need stricter access controls. The annotation approach gives me method-level control over cross-origin resource sharing."

---

### Question 154: How do you implement role-based access control at controller level?

**Answer:**
Enable Global Method Security: `@EnableMethodSecurity`.
Annotate Method: `@PreAuthorize("hasRole('ADMIN')")`.
If User lacks role, Spring Security throws `AccessDeniedException`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement role-based access control at controller level?
**Your Response:** "I implement role-based access control using Spring Security's method-level security. First, I enable global method security with `@EnableMethodSecurity`. Then I annotate individual methods with `@PreAuthorize('hasRole('ADMIN')')` to restrict access based on user roles. If a user lacks the required role, Spring Security throws an `AccessDeniedException` which I can handle globally. This approach provides fine-grained security control at the method level, allowing me to specify exactly who can access each endpoint based on their roles and permissions. It's more precise than URL-based security and integrates seamlessly with Spring Security's authentication system."

---

### Question 155: How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?

**Answer:**
Return `StreamingResponseBody`.
Writes directly to the OutputStream of the Servlet Response asynchronously.
Avoids loading the whole file into RAM.
Set `Content-Type: application/pdf`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?
**Your Response:** "For large files like PDFs or ZIP archives, I use `StreamingResponseBody` to stream content directly to the client without loading the entire file into memory. I implement the `writeTo` method to write directly to the response's OutputStream asynchronously. This approach is memory-efficient and allows serving large files without causing heap issues. I set the appropriate `Content-Type` header like `application/pdf` so browsers know how to handle the file. This streaming approach is essential for applications that serve large files or have memory constraints."

---

### Question 156: How do you implement ETag support in REST responses?

**Answer:**
Bean: `ShallowEtagHeaderFilter`.
It intercepts the response, hashes the body (MD5), and sets `ETag` header.
On next request, checks `If-None-Match`. If matches, returns 304 Not Modified (empty body). Saves Bandwidth, not CPU.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement ETag support in REST responses?
**Your Response:** "I implement ETag support using the `ShallowEtagHeaderFilter` bean. This filter automatically intercepts responses, hashes the response body using MD5, and sets the `ETag` header. On subsequent requests, it checks the `If-None-Match` header - if the ETag matches, it returns a 304 Not Modified response with an empty body. This saves bandwidth because the full response isn't sent when the content hasn't changed. While it uses some CPU to calculate the hash, it's very effective for APIs that serve relatively static content or for client-side caching strategies."

---

### Question 157: How to apply request/response compression in Spring Boot?

**Answer:**
Properties:
```properties
server.compression.enabled=true
server.compression.mime-types=text/html,application/json
server.compression.min-response-size=1024
```
Uses GZIP.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to apply request/response compression in Spring Boot?
**Your Response:** "I enable compression through simple configuration properties. I set `server.compression.enabled=true` and specify which MIME types to compress like `text/html,application/json`. I can also set a minimum response size threshold with `server.compression.min-response-size=1024` to avoid compressing very small responses. Spring Boot uses GZIP compression automatically for responses that meet the criteria. This significantly reduces response sizes and improves performance, especially for APIs that return large JSON payloads or HTML content. The configuration is straightforward but has a big impact on network performance."

---

### Question 158: What are the options to handle request timeout in Spring controllers?

**Answer:**
1.  **Callable<T>:** Spring MVC manages a separate thread.
2.  **DeferredResult<T>:** For specialized async.
3.  **WebFlux:** Native timeout support.
4.  **Properties:** `spring.mvc.async.request-timeout`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the options to handle request timeout in Spring controllers?
**Your Response:** "I have several approaches for handling request timeouts in Spring controllers. For traditional MVC, I can use `Callable<T>` which Spring executes on a separate thread, or `DeferredResult<T>` for more specialized async scenarios. For reactive applications, WebFlux has native timeout support. I can also configure global timeout settings with properties like `spring.mvc.async.request-timeout`. Each approach serves different needs - `Callable` is simple async processing, `DeferredResult` gives me manual control over when the response is ready, and reactive timeouts work naturally with the reactive programming model."

---

### Question 159: How do you use filters to intercept and modify incoming requests?

**Answer:**
Implement `javax.servlet.Filter`.
Annotate with `@Component` (Auto-registered for `/*`).
Override `doFilter(req, res, chain)`.
Or use `FilterRegistrationBean` to control URL mapping patterns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use filters to intercept and modify incoming requests?
**Your Response:** "I implement filters by creating a class that implements `javax.servlet.Filter` and annotating it with `@Component` for automatic registration. I override the `doFilter` method to intercept requests and responses. For more control over URL patterns, I use `FilterRegistrationBean` to specify which URLs the filter should apply to. Filters are perfect for cross-cutting concerns like logging, authentication, or request modification. They run before the controller and can modify both the incoming request and outgoing response, making them ideal for implementing security, logging, or request/response transformations."

---

### Question 160: How does Spring Boot support OpenAPI/Swagger documentation?

**Answer:**
(See Q57). `springdoc-openapi`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot support OpenAPI/Swagger documentation?
**Your Response:** "Spring Boot supports OpenAPI documentation through the `springdoc-openapi` library. I add the dependency and Spring Boot automatically generates interactive API documentation based on my controllers and models. The documentation is available at `/swagger-ui/index.html` where I can explore and test the API endpoints. Spring Boot automatically detects REST controllers, request/response models, and annotations to build comprehensive documentation. This makes it easy to maintain up-to-date API documentation that's always in sync with the actual code, which is invaluable for both development teams and API consumers."

---
