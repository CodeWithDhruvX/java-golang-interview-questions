## 🔹 Section 3: REST API Internals (141-160)

### Question 141: How does Spring Boot handle parameter binding in REST endpoints?

**Answer:**
It uses `HandlerMethodArgumentResolver` strategies.
- Looks at annotation (`@RequestParam`, `@PathVariable`).
- Uses reflection to convert String (from URL) to Target Type (Int, Date, Enum).
- Uses `WebDataBinder` for type conversion.

---

### Question 142: What is the difference between `HttpMessageConverter` and `ObjectMapper`?

**Answer:**
- **ObjectMapper (Jackson):** The low-level library that actually converts Java Object <-> JSON.
- **HttpMessageConverter:** The Spring abstraction that selects *which* library to use (Jackson, Gson) based on Content-Type header. It delegates the work to ObjectMapper.

---

### Question 143: How do you control serialization and deserialization of JSON in Spring Boot?

**Answer:**
1.  **Annotations:** `@JsonProperty("name")`, `@JsonIgnore`.
2.  **Config:** `spring.jackson.date-format`, `spring.jackson.property-naming-strategy`.
3.  **Customizer:** Register a `Jackson2ObjectMapperBuilderCustomizer` bean to tweak the global Mapper.

---

### Question 144: How do you return paginated responses with metadata?

**Answer:**
Return a `Page<T>` object (from Spring Data).
Spring Boot automatically serializes this into JSON containing:
- `content`: [List of items]
- `pageable`: { pageNumber, pageSize }
- `totalPages`, `totalElements`.

---

### Question 145: What is the difference between `@ModelAttribute` and `@RequestBody`?

**Answer:**
- **`@RequestBody`:** Consumes the HTTP Body (usually JSON/XML). Uses MessageConverters.
- **`@ModelAttribute`:** Binds Query Parameters or Form Data (`application/x-www-form-urlencoded`) to a Java Bean. Used in MVC Views usually.

---

### Question 146: How to validate request parameters in Spring Boot REST API?

**Answer:**
1.  Add `spring-boot-starter-validation`.
2.  Annotate Controller Class with `@Validated`.
3.  Annotate Param: `getItems(@RequestParam @Min(1) int count)`.
Throws `ConstraintViolationException` if invalid.

---

### Question 147: How does Spring Boot support HATEOAS APIs?

**Answer:**
Dependency: `spring-boot-starter-hateoas`.
Wraps response (DTO) in `EntityModel<T>`.
Adds links: `model.add(linkTo(methodOn(Controller.class).get(id)).withSelfRel())`.
Response includes `_links` section (Hypermedia).

---

### Question 148: How do you apply global validation error handling?

**Answer:**
Use `@ControllerAdvice`.
Catch `MethodArgumentNotValidException` (triggered by `@Valid` on `@RequestBody`).
Extract FieldErrors and return a clean JSON error map.

---

### Question 149: What is the use of `BindingResult` in controller methods?

**Answer:**
Used immediately after `@Valid` argument.
`func(@Valid @RequestBody User user, BindingResult result)`.
If validation fails, Spring does NOT throw exception. Instead, it populates `result`. You can check `result.hasErrors()` and handle logic manually.

---

### Question 150: How do you apply rate limiting to Spring Boot REST endpoints?

**Answer:**
Spring Boot doesn't have built-in Rate Limiting.
1.  **Bucket4j:** Library providing Token Bucket algorithm.
2.  **Resilience4j:** RateLimiter module.
3.  **Redis:** Implement custom Lua script or use API Gateway (Zuul/Spring Cloud Gateway).

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

---

### Question 154: How do you implement role-based access control at controller level?

**Answer:**
Enable Global Method Security: `@EnableMethodSecurity`.
Annotate Method: `@PreAuthorize("hasRole('ADMIN')")`.
If User lacks role, Spring Security throws `AccessDeniedException`.

---

### Question 155: How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?

**Answer:**
Return `StreamingResponseBody`.
Writes directly to the OutputStream of the Servlet Response asynchronously.
Avoids loading the whole file into RAM.
Set `Content-Type: application/pdf`.

---

### Question 156: How do you implement ETag support in REST responses?

**Answer:**
Bean: `ShallowEtagHeaderFilter`.
It intercepts the response, hashes the body (MD5), and sets `ETag` header.
On next request, checks `If-None-Match`. If matches, returns 304 Not Modified (empty body). Saves Bandwidth, not CPU.

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

---

### Question 158: What are the options to handle request timeout in Spring controllers?

**Answer:**
1.  **Callable<T>:** Spring MVC manages a separate thread.
2.  **DeferredResult<T>:** For specialized async.
3.  **WebFlux:** Native timeout support.
4.  **Properties:** `spring.mvc.async.request-timeout`.

---

### Question 159: How do you use filters to intercept and modify incoming requests?

**Answer:**
Implement `javax.servlet.Filter`.
Annotate with `@Component` (Auto-registered for `/*`).
Override `doFilter(req, res, chain)`.
Or use `FilterRegistrationBean` to control URL mapping patterns.

---

### Question 160: How does Spring Boot support OpenAPI/Swagger documentation?

**Answer:**
(See Q57). `springdoc-openapi`.

---
