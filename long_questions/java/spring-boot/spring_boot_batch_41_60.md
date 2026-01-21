## 🔹 Section 3: REST API with Spring Boot (41–60)

### Question 41: How to create a REST API using Spring Boot?

**Answer:**
1.  Add `spring-boot-starter-web` dependency.
2.  Create a class annotated with `@RestController`.
3.  Define methods with `@GetMapping`, `@PostMapping` mapped to URL paths.
4.  Return objects (POJOs), which are automatically serialized to JSON by Jackson.

---

### Question 42: What are `@RestController` and `@Controller`?

**Answer:**
*   **`@Controller`**: Standard Spring MVC controller. Used for returning Views (JSP/Thymeleaf). Returns a String (view name).
*   **`@RestController`**: Convenience annotation = `@Controller` + `@ResponseBody`. Returns Data (JSON/XML) directly to the response body.

---

### Question 43: How to handle HTTP methods (GET, POST, PUT, DELETE) in Spring Boot?

**Answer:**
Use specialized annotations:
*   `@GetMapping("/users")` -> READ
*   `@PostMapping("/users")` -> CREATE
*   `@PutMapping("/users/{id}")` -> UPDATE (Replace)
*   `@DeleteMapping("/users/{id}")` -> DELETE
*   `@PatchMapping("/users/{id}")` -> UPDATE (Partial)

---

### Question 44: What is the purpose of `@RequestMapping`, `@GetMapping`, etc.?

**Answer:**
*   **`@RequestMapping`**: Generic mapping. Can handle any method (`method = RequestMethod.GET`). Often used at Class Level to define base URL prefix.
*   **`@GetMapping`**: Shortcut for RequestMapping with GET method.

---

### Question 45: What is the use of `@PathVariable` and `@RequestParam`?

**Answer:**
*   **`@PathVariable`**: Extracts values from the URI path.
    *   URL: `/users/101` -> `@GetMapping("/users/{id}") func(@PathVariable id)`
*   **`@RequestParam`**: Extracts values from Query Parameters.
    *   URL: `/users?role=admin` -> `func(@RequestParam role)`

---

### Question 46: How to handle form-data and JSON in Spring Boot controllers?

**Answer:**
*   **JSON:** Use `@RequestBody` to bind the request body to a Java Object.
*   **Form Data:** Use `@ModelAttribute` or simply method arguments matching the form field names for `application/x-www-form-urlencoded`.

---

### Question 47: How to return a custom HTTP status in Spring Boot?

**Answer:**
1.  **ResponseEntity:** `return ResponseEntity.status(HttpStatus.CREATED).body(data);`
2.  **@ResponseStatus:** Annotate an Exception with `@ResponseStatus(HttpStatus.NOT_FOUND)` or the method itself.

---

### Question 48: What is `ResponseEntity` and how to use it?

**Answer:**
It represents the entire HTTP response: Status Code, Headers, and Body.
It gives you full control over the response.
```java
return ResponseEntity.ok()
        .header("Custom-Header", "value")
        .body(myObject);
```

---

### Question 49: How to handle exceptions globally using `@ControllerAdvice`?

**Answer:**
Create a class annotated with `@RestControllerAdvice`.
Define methods with `@ExceptionHandler(MyException.class)`.
This centralized error handling logic makes controllers cleaner.

```java
@ExceptionHandler(UserNotFoundException.class)
public ResponseEntity<String> handleUserNotFound(Exception ex) {
    return ResponseEntity.status(404).body(ex.getMessage());
}
```

---

### Question 50: What is the difference between synchronous and asynchronous API?

**Answer:**
*   **Sync:** The server thread blocks until the process finishes. Limits scalability if operations are slow (I/O).
*   **Async:** The server thread delegates work and returns immediately (release thread). Used with `CompletableFuture` or WebFlux (`Mono`/`Flux`) for high throughput.

---

### Question 51: What is CORS and how do you enable it in Spring Boot?

**Answer:**
Cross-Origin Resource Sharing. Browser security blocks requests from different domains (e.g., React on port 3000 calling Spring on 8080).
**Enable:**
1.  **Global:** `WebMvcConfigurer.addCorsMappings()`.
2.  **Local:** `@CrossOrigin(origins = "http://localhost:3000")` on controller method/class.

---

### Question 52: How to log incoming requests and responses in a REST controller?

**Answer:**
1.  **Filter:** Implement `CommonsRequestLoggingFilter` or write a custom `OncePerRequestFilter`.
2.  **Interceptor:** Use `HandlerInterceptor` (`preHandle`, `afterCompletion`).
3.  **AOP:** Use AspectJ to log around controller execution.

---

### Question 53: What is content negotiation in Spring Boot?

**Answer:**
The ability to serve different representations of the same resource (JSON, XML, PDF) based on the `Accept` header sent by the client.
Spring Boot supports this via `HttpMessageConverters`. Just add `jackson-dataformat-xml` dependency to support XML.

---

### Question 54: How to paginate and sort results in REST API?

**Answer:**
Use Spring Data `Pageable`.
Controller accepts `Pageable` object (handled by `PageableHandlerMethodArgumentResolver`).
URL: `/users?page=0&size=10&sort=name,asc`.
Repo: `repo.findAll(pageable)`.

---

### Question 55: What are DTOs and how are they used in Spring Boot?

**Answer:**
**Data Transfer Objects**.
Plain Java objects used to transfer data between subsystems (Controller <-> Service).
Decouples the internal Database Entity (JPA) from the API contract exposed to the user. Prevents leaking sensitive fields (password).

---

### Question 56: How to map entities to DTOs? (ModelMapper/MapStruct)

**Answer:**
*   **Manual:** Setters/Getters.
*   **ModelMapper:** Runtime reflection (Slower). `mapper.map(entity, Dto.class)`.
*   **MapStruct:** Compile-time code generation (Fastest). Defines an interface, MapStruct generates the impl.

---

### Question 57: How do you document a Spring Boot REST API using Swagger/OpenAPI?

**Answer:**
Add dependency: `springdoc-openapi-starter-webmvc-ui`.
Run the app.
Access UI at `/swagger-ui/index.html`.
It automatically scans controllers and generates interactive docs.

---

### Question 58: How to secure REST endpoints in Spring Boot?

**Answer:**
Use **Spring Security**.
Config: `SecurityFilterChain`.
`http.authorizeHttpRequests().requestMatchers("/admin/**").hasRole("ADMIN")`.
Use JWT or Basic Auth for stateless APIs.

---

### Question 59: How to test REST APIs using `MockMvc` or `RestTemplate`?

**Answer:**
*   **MockMvc:** Unit testing. Mocks the Servlet Container. Fast. Verifies HTTP status, body, headers without starting server.
*   **TestRestTemplate:** Integration testing. Starts actual server. Makes real HTTP calls.

---

### Question 60: What is the difference between `WebClient` and `RestTemplate`?

**Answer:**
*   **RestTemplate:** Blocking, Synchronous client. (Maintenance mode).
*   **WebClient:** Non-blocking, Asynchronous (Reactive). Part of Spring WebFlux. Can be used in sync code (`.block()`). The modern standard.

---
