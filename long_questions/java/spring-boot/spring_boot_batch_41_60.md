## 🔹 Section 3: REST API with Spring Boot (41–60)

### Question 41: How to create a REST API using Spring Boot?

**Answer:**
1.  Add `spring-boot-starter-web` dependency.
2.  Create a class annotated with `@RestController`.
3.  Define methods with `@GetMapping`, `@PostMapping` mapped to URL paths.
4.  Return objects (POJOs), which are automatically serialized to JSON by Jackson.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to create a REST API using Spring Boot?
**Your Response:** "Creating a REST API in Spring Boot is straightforward. First, I add the `spring-boot-starter-web` dependency which brings in Spring MVC and an embedded Tomcat server. Then I create a controller class annotated with `@RestController` instead of the traditional `@Controller`. I define endpoints using annotations like `@GetMapping` and `@PostMapping` mapped to specific URL paths. The beauty of Spring Boot is that I can simply return POJOs from my controller methods, and Jackson automatically serializes them to JSON. No manual JSON conversion needed - Spring handles all the content negotiation and serialization automatically."

---

### Question 42: What are `@RestController` and `@Controller`?

**Answer:**
*   **`@Controller`**: Standard Spring MVC controller. Used for returning Views (JSP/Thymeleaf). Returns a String (view name).
*   **`@RestController`**: Convenience annotation = `@Controller` + `@ResponseBody`. Returns Data (JSON/XML) directly to the response body.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are `@RestController` and `@Controller`?
**Your Response:** "The key difference is their purpose. `@Controller` is the traditional Spring MVC annotation used for web applications that return views like JSP or Thymeleaf pages - it typically returns a String representing the view name. `@RestController` is a convenience annotation that combines `@Controller` and `@ResponseBody`, specifically designed for REST APIs. When I use `@RestController`, every method automatically adds `@ResponseBody`, meaning the return value is written directly to the HTTP response body as JSON or XML, rather than being treated as a view name. This makes it perfect for building RESTful web services."

---

### Question 43: How to handle HTTP methods (GET, POST, PUT, DELETE) in Spring Boot?

**Answer:**
Use specialized annotations:
*   `@GetMapping("/users")` -> READ
*   `@PostMapping("/users")` -> CREATE
*   `@PutMapping("/users/{id}")` -> UPDATE (Replace)
*   `@DeleteMapping("/users/{id}")` -> DELETE
*   `@PatchMapping("/users/{id}")` -> UPDATE (Partial)

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle HTTP methods (GET, POST, PUT, DELETE) in Spring Boot?
**Your Response:** "Spring Boot provides specialized annotations for each HTTP method that make building REST APIs intuitive. I use `@GetMapping` for reading data, like fetching all users or a specific user. `@PostMapping` is for creating new resources. `@PutMapping` handles full updates where the client sends the complete resource representation. `@DeleteMapping` is for removing resources. And `@PatchMapping` is for partial updates when the client only wants to modify certain fields. These annotations are more readable and type-safe compared to the generic `@RequestMapping` with method attributes."

---

### Question 44: What is the purpose of `@RequestMapping`, `@GetMapping`, etc.?

**Answer:**
*   **`@RequestMapping`**: Generic mapping. Can handle any method (`method = RequestMethod.GET`). Often used at Class Level to define base URL prefix.
*   **`@GetMapping`**: Shortcut for RequestMapping with GET method.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `@RequestMapping`, `@GetMapping`, etc.?
**Your Response:** "`@RequestMapping` is the generic mapping annotation that can handle any HTTP method when specified with the method attribute. I often use it at the class level to define a base URL prefix for all endpoints in that controller. The specialized annotations like `@GetMapping`, `@PostMapping`, etc., are essentially shortcuts - they're more concise and readable. For example, `@GetMapping("/users")` is equivalent to `@RequestMapping(value="/users", method=RequestMethod.GET)`. Using the specific method annotations makes my code more expressive and easier to understand at a glance."

---

### Question 45: What is the use of `@PathVariable` and `@RequestParam`?

**Answer:**
*   **`@PathVariable`**: Extracts values from the URI path.
    *   URL: `/users/101` -> `@GetMapping("/users/{id}") func(@PathVariable id)`
*   **`@RequestParam`**: Extracts values from Query Parameters.
    *   URL: `/users?role=admin` -> `func(@RequestParam role)`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of `@PathVariable` and `@RequestParam`?
**Your Response:** "These annotations extract different parts of the URL. `@PathVariable` extracts values from the URI path itself - perfect for resource identifiers. For example, in `/users/101`, I can extract the `101` using `@PathVariable` with a path template like `/users/{id}`. `@RequestParam` extracts query parameters from the URL after the question mark - useful for filtering, sorting, or pagination. For instance, in `/users?role=admin`, I can extract the role parameter using `@RequestParam`. Path variables are for required resource identification, while request parameters are optional filtering criteria."

---

### Question 46: How to handle form-data and JSON in Spring Boot controllers?
### Question 46: which annotation maps a json request body directly into java object?
**Answer:**
*   **JSON:** Use `@RequestBody` to bind the request body to a Java Object.
*   **Form Data:** Use `@ModelAttribute` or simply method arguments matching the form field names for `application/x-www-form-urlencoded`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle form-data and JSON in Spring Boot controllers?
**Your Response:** "For different content types, I use different annotations. When clients send JSON data, I use `@RequestBody` which automatically deserializes the JSON request body into a Java object using Jackson. For traditional form submissions with `application/x-www-form-urlencoded` content type, I can use `@ModelAttribute` to bind form fields to an object, or simply use method parameters that match the form field names - Spring automatically maps them. The key is matching the annotation to the content type: `@RequestBody` for JSON payloads and `@ModelAttribute` or simple parameters for form data."

---

### Question 47: How to return a custom HTTP status in Spring Boot?

**Answer:**
1.  **ResponseEntity:** `return ResponseEntity.status(HttpStatus.CREATED).body(data);`
2.  **@ResponseStatus:** Annotate an Exception with `@ResponseStatus(HttpStatus.NOT_FOUND)` or the method itself.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to return a custom HTTP status in Spring Boot?
**Your Response:** "I have two main ways to control HTTP status codes. The most flexible is using `ResponseEntity` which gives me complete control over the response - I can set the status, headers, and body all in one fluent call like `ResponseEntity.status(HttpStatus.CREATED).body(data)`. For simpler cases where I just need to set a status code, I can use the `@ResponseStatus` annotation either on the method itself or on an exception class. When I annotate an exception with `@ResponseStatus(HttpStatus.NOT_FOUND)`, Spring automatically returns that status code when the exception is thrown. ResponseEntity is more powerful, while @ResponseStatus is concise for straightforward cases."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `ResponseEntity` and how to use it?
**Your Response:** "`ResponseEntity` represents the entire HTTP response - it gives me complete control over the status code, headers, and body. Instead of just returning data, I can return a `ResponseEntity` and build the response using its fluent API. For example, I can use `ResponseEntity.ok()` for a 200 status, add custom headers with `.header()`, and set the body with `.body()`. I can also create responses with different statuses like `ResponseEntity.created()` for 201 or `ResponseEntity.notFound()` for 404. This is particularly useful when I need to set custom headers or return specific status codes based on business logic."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle exceptions globally using `@ControllerAdvice`?
**Your Response:** "`@ControllerAdvice` (or `@RestControllerAdvice` for REST APIs) allows me to create centralized exception handling across all controllers. Instead of adding try-catch blocks in every controller method, I create a separate class annotated with `@RestControllerAdvice` and define methods with `@ExceptionHandler` for specific exception types. When any controller throws that exception, Spring automatically routes it to my handler method. This makes my controllers much cleaner since they focus on business logic, and I get consistent error responses across the entire application. It's the standard way to implement global error handling in Spring Boot."

---

### Question 50: What is the difference between synchronous and asynchronous API?

**Answer:**
*   **Sync:** The server thread blocks until the process finishes. Limits scalability if operations are slow (I/O).
*   **Async:** The server thread delegates work and returns immediately (release thread). Used with `CompletableFuture` or WebFlux (`Mono`/`Flux`) for high throughput.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between synchronous and asynchronous API?
**Your Response:** "The key difference is how threads are handled. In synchronous APIs, the server thread blocks until the entire request processing completes, which limits scalability when dealing with slow operations like database queries or external API calls. In asynchronous APIs, the server thread delegates the work to background workers and returns immediately, freeing up the thread to handle other requests. I implement async using `CompletableFuture` in traditional Spring MVC, or better yet, use Spring WebFlux with reactive types `Mono` and `Flux` for true non-blocking I/O. This allows handling many more concurrent connections with fewer threads, significantly improving throughput for I/O-bound operations."

---

### Question 51: What is CORS and how do you enable it in Spring Boot?

**Answer:**
Cross-Origin Resource Sharing. Browser security blocks requests from different domains (e.g., React on port 3000 calling Spring on 8080).
**Enable:**
1.  **Global:** `WebMvcConfigurer.addCorsMappings()`.
2.  **Local:** `@CrossOrigin(origins = "http://localhost:3000")` on controller method/class.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CORS and how do you enable it in Spring Boot?
**Your Response:** "CORS stands for Cross-Origin Resource Sharing, which is a browser security mechanism that blocks requests from different domains. For example, if my React app runs on port 3000 and tries to call my Spring Boot API on port 8080, the browser will block it by default. To enable CORS, I have two main approaches. For global configuration, I implement `WebMvcConfigurer` and override the `addCorsMappings` method to define CORS policies for the entire application. For more specific control, I can use the `@CrossOrigin` annotation on individual controllers or methods, specifying exactly which origins are allowed. This gives me fine-grained control over which domains can access my API."

---

### Question 52: How to log incoming requests and responses in a REST controller?

**Answer:**
1.  **Filter:** Implement `CommonsRequestLoggingFilter` or write a custom `OncePerRequestFilter`.
2.  **Interceptor:** Use `HandlerInterceptor` (`preHandle`, `afterCompletion`).
3.  **AOP:** Use AspectJ to log around controller execution.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to log incoming requests and responses in a REST controller?
**Your Response:** "I have several approaches for logging HTTP requests and responses. I can implement a custom filter by extending `OncePerRequestFilter` which gives me access to the request and response objects before and after controller execution. Another option is using Spring's `HandlerInterceptor` with `preHandle` and `afterCompletion` methods. For a more elegant solution, I can use AOP with AspectJ to create advice around controller methods. Each approach has different benefits - filters are great for low-level HTTP concerns, interceptors are good for cross-cutting controller concerns, and AOP provides clean separation of logging logic from business code."

---

### Question 53: What is content negotiation in Spring Boot?

**Answer:**
The ability to serve different representations of the same resource (JSON, XML, PDF) based on the `Accept` header sent by the client.
Spring Boot supports this via `HttpMessageConverters`. Just add `jackson-dataformat-xml` dependency to support XML.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is content negotiation in Spring Boot?
**Your Response:** "Content negotiation is Spring Boot's ability to serve different representations of the same resource based on what the client requests via the `Accept` header. For example, the same endpoint can return JSON, XML, or even PDF depending on the client's preferences. Spring Boot handles this automatically through `HttpMessageConverters`. By default, it supports JSON, but if I add a dependency like `jackson-dataformat-xml`, it can automatically serialize responses as XML when clients request it with `Accept: application/xml`. This means I can write one controller method and serve multiple client types without changing the code - Spring handles the content negotiation automatically."

---

### Question 54: How to paginate and sort results in REST API?

**Answer:**
Use Spring Data `Pageable`.
Controller accepts `Pageable` object (handled by `PageableHandlerMethodArgumentResolver`).
URL: `/users?page=0&size=10&sort=name,asc`.
Repo: `repo.findAll(pageable)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to paginate and sort results in REST API?
**Your Response:** "Spring Data makes pagination and sorting incredibly easy. I simply accept a `Pageable` parameter in my controller method, and Spring automatically resolves URL parameters like `page=0&size=10&sort=name,asc` into a `Pageable` object. I pass this to my repository's `findAll(pageable)` method, which returns a `Page` object containing the data slice plus metadata like total elements and total pages. The client gets back not just the data but also pagination information they can use to build navigation controls. This approach is clean, type-safe, and requires minimal code - Spring handles all the SQL generation and pagination logic behind the scenes."

---

### Question 55: What are DTOs and how are they used in Spring Boot?

**Answer:**
**Data Transfer Objects**.
Plain Java objects used to transfer data between subsystems (Controller <-> Service).
Decouples the internal Database Entity (JPA) from the API contract exposed to the user. Prevents leaking sensitive fields (password).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are DTOs and how are they used in Spring Boot?
**Your Response:** "DTOs, or Data Transfer Objects, are plain Java objects used to transfer data between different layers of my application, particularly between the controller and service layers. The key benefit is that they decouple my internal database entities from the API contract I expose to clients. This prevents leaking sensitive information like passwords or internal fields that clients shouldn't see. DTOs also give me control over the exact shape of my API responses - I can include only the fields clients need and format them appropriately. This separation makes my API more stable and secure, as I can change my internal database schema without breaking the public API."

---

### Question 56: How to map entities to DTOs? (ModelMapper/MapStruct)

**Answer:**
*   **Manual:** Setters/Getters.
*   **ModelMapper:** Runtime reflection (Slower). `mapper.map(entity, Dto.class)`.
*   **MapStruct:** Compile-time code generation (Fastest). Defines an interface, MapStruct generates the impl.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to map entities to DTOs? (ModelMapper/MapStruct)
**Your Response:** "I have several approaches for mapping entities to DTOs. The manual approach using setters and getters gives me full control but requires boilerplate code. ModelMapper uses reflection at runtime to automatically map fields with similar names - it's convenient but slower due to the reflection overhead. My preferred approach is MapStruct, which generates mapping code at compile time. I define an interface with mapping methods, and MapStruct generates the implementation during compilation. This gives me the performance of manual mapping with the convenience of automatic mapping. It's type-safe, fast, and catches mapping errors during compilation rather than at runtime."

---

### Question 57: How do you document a Spring Boot REST API using Swagger/OpenAPI?

**Answer:**
Add dependency: `springdoc-openapi-starter-webmvc-ui`.
Run the app.
Access UI at `/swagger-ui/index.html`.
It automatically scans controllers and generates interactive docs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you document a Spring Boot REST API using Swagger/OpenAPI?
**Your Response:** "Documenting REST APIs with Spring Boot is straightforward using OpenAPI. I add the `springdoc-openapi-starter-webmvc-ui` dependency, and Spring Boot automatically generates interactive API documentation. The library scans my controllers and method signatures to create comprehensive documentation that's available at `/swagger-ui/index.html. I can enhance the documentation by adding annotations like `@Tag` for controller descriptions and `@Operation` for endpoint details. The best part is that clients can actually try out the API directly from the Swagger UI, which makes it incredibly useful for both documentation and testing."

---

### Question 58: How to secure REST endpoints in Spring Boot?

**Answer:**
Use **Spring Security**.
Config: `SecurityFilterChain`.
`http.authorizeHttpRequests().requestMatchers("/admin/**").hasRole("ADMIN")`.
Use JWT or Basic Auth for stateless APIs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to secure REST endpoints in Spring Boot?
**Your Response:** "I secure REST endpoints using Spring Security, which provides comprehensive security features. I configure a `SecurityFilterChain` bean where I define security rules using `authorizeHttpRequests()`. For example, I can restrict admin endpoints with `.requestMatchers('/admin/**').hasRole('ADMIN')` while allowing public access to certain endpoints. For authentication, I use JWT tokens for stateless APIs where the client includes a bearer token in each request, or Basic Authentication for simpler scenarios. Spring Security handles the authentication and authorization automatically, integrating seamlessly with my REST controllers."

---

### Question 59: How to test REST APIs using `MockMvc` or `RestTemplate`?

**Answer:**
*   **MockMvc:** Unit testing. Mocks the Servlet Container. Fast. Verifies HTTP status, body, headers without starting server.
*   **TestRestTemplate:** Integration testing. Starts actual server. Makes real HTTP calls.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test REST APIs using `MockMvc` or `RestTemplate`?
**Your Response:** "I use different tools for different testing scenarios. `MockMvc` is perfect for unit testing - it mocks the servlet container so tests run fast without starting a real server. I can verify HTTP status codes, response bodies, and headers in isolation. For integration testing, I use `TestRestTemplate` which starts the actual server and makes real HTTP calls, testing the complete request-response cycle. MockMvc tests are faster and good for testing controller logic in isolation, while TestRestTemplate tests are slower but catch integration issues that MockMvc might miss, like configuration problems or filter chain issues."

---

### Question 60: What is the difference between `WebClient` and `RestTemplate`?

**Answer:**
*   **RestTemplate:** Blocking, Synchronous client. (Maintenance mode).
*   **WebClient:** Non-blocking, Asynchronous (Reactive). Part of Spring WebFlux. Can be used in sync code (`.block()`). The modern standard.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `WebClient` and `RestTemplate`?
**Your Response:** "The key difference is their programming model. `RestTemplate` is the traditional blocking, synchronous client - each HTTP call blocks the thread until the response arrives. It's now in maintenance mode. `WebClient` is the modern, non-blocking reactive client from Spring WebFlux. It can handle many concurrent requests with fewer threads, making it much more scalable. Even though it's reactive, I can use it in synchronous code by calling `.block()` when needed. For new applications, I always choose `WebClient` because it's more efficient, supports modern reactive patterns, and is the recommended approach going forward."

---

### Question 61: Scenario: PNC Bank Microservice - Custom Exception Handling for Business Rule Violations

**Answer:**
1. **Create Custom Exception Hierarchy:** Base `BankingException` with specific `UnauthorizedTransactionException`
2. **Service Layer Validation:** Manual exception throwing when business rules are violated
3. **Global Exception Handler:** `@RestControllerAdvice` with `@ExceptionHandler` for consistent error responses
4. **HTTP Status Mapping:** Return appropriate status codes (403 Forbidden, 400 Bad Request)

### How to Explain in Interview (Spoken style format)
**Interviewer:** You're building a microservice for PNC where a specific business rule is violated (e.g., an unauthorized transaction attempt). Explain how you would implement a Custom Manual Exception and ensure it is caught to return a meaningful error message and status code to the client.

**Your Response:** "For this PNC microservice scenario, I would implement a comprehensive custom exception handling strategy. First, I'd create a domain-specific custom exception called `UnauthorizedTransactionException` that extends a base `BankingException` class. This exception would include important context like the transaction amount, account number, and the specific business rule that was violated.

In my service layer, I would validate the business rules before processing any transaction. For example, I'd check if the account has sufficient funds, if the transaction amount exceeds daily limits, or if there are any regulatory restrictions. When a rule is violated, I would manually throw my custom exception with a descriptive message and error code.

To ensure consistent error handling across the entire microservice, I would implement a global exception handler using `@RestControllerAdvice`. This global handler would catch my `UnauthorizedTransactionException` specifically and map it to an appropriate HTTP status code - typically `403 Forbidden` for authorization issues or `400 Bad Request` for business rule violations. The handler would return a structured error response with the error code, message, and additional context that helps the client understand what went wrong.

The key benefits are that my controllers stay clean and focused on business logic, I get consistent error responses across all endpoints, and I can easily add logging and monitoring for these specific business rule violations."

**Code Implementation:**
```java
// Base exception
@RestControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(UnauthorizedTransactionException.class)
    public ResponseEntity<ErrorResponse> handleUnauthorizedTransaction(
            UnauthorizedTransactionException ex) {
        ErrorResponse error = new ErrorResponse(
            "UNAUTHORIZED_TRANSACTION", 
            ex.getMessage(),
            LocalDateTime.now()
        );
        return ResponseEntity
            .status(HttpStatus.FORBIDDEN)
            .body(error);
    }
    
    @ExceptionHandler(InsufficientFundsException.class)
    public ResponseEntity<ErrorResponse> handleInsufficientFunds(
            InsufficientFundsException ex) {
        return ResponseEntity
            .status(HttpStatus.BAD_REQUEST)
            .body(new ErrorResponse("INSUFFICIENT_FUNDS", ex.getMessage()));
    }
}

// Service layer
@Service
public class TransactionService {
    public void processTransaction(TransactionRequest request) 
            throws UnauthorizedTransactionException {
        if (!isAuthorized(request.getAccountNumber(), request.getAmount())) {
            throw new UnauthorizedTransactionException(
                request.getAccountNumber(), 
                request.getAmount(),
                "Daily transaction limit exceeded"
            );
        }
    }
}
```

**Follow-up:** For different types of violations, I'd create a hierarchy: `InsufficientFundsException`, `DailyLimitExceededException`, `RegulatoryViolationException`, each with specific `@ExceptionHandler` methods returning appropriate HTTP status codes.

---
