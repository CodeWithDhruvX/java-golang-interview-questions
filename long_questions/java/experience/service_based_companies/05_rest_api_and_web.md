# 🌐 05 — REST API & Web Development
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Spring MVC request mapping annotations
- Request/response handling (body, params, headers, path variables)
- Validation with `@Valid` / Bean Validation
- Exception handling with `@ControllerAdvice`
- HTTP status codes
- JSON serialization with Jackson
- REST best practices

---

## ❓ Most Asked Questions

### Q1. How do you build a REST API in Spring Boot?

```java
@RestController
@RequestMapping("/api/v1/products")
public class ProductController {

    private final ProductService productService;

    public ProductController(ProductService productService) {
        this.productService = productService;
    }

    // GET all
    @GetMapping
    public ResponseEntity<List<ProductDTO>> getAll(
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "20") int size) {
        return ResponseEntity.ok(productService.findAll(page, size));
    }

    // GET by ID
    @GetMapping("/{id}")
    public ResponseEntity<ProductDTO> getById(@PathVariable Long id) {
        return ResponseEntity.ok(productService.findById(id));
    }

    // POST — create
    @PostMapping
    public ResponseEntity<ProductDTO> create(@Valid @RequestBody ProductDTO dto) {
        ProductDTO created = productService.create(dto);
        URI location = URI.create("/api/v1/products/" + created.getId());
        return ResponseEntity.created(location).body(created);  // 201 Created
    }

    // PUT — full update
    @PutMapping("/{id}")
    public ResponseEntity<ProductDTO> update(
            @PathVariable Long id,
            @Valid @RequestBody ProductDTO dto) {
        return ResponseEntity.ok(productService.update(id, dto));
    }

    // DELETE
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> delete(@PathVariable Long id) {
        productService.delete(id);
        return ResponseEntity.noContent().build();  // 204 No Content
    }
}
```

---

### 🎯 How to Explain in Interview

"Building a REST API in Spring Boot is straightforward with the right annotations. I start with @RestController which combines @Controller and @ResponseBody, so my methods return JSON automatically. @RequestMapping defines the base path for all endpoints, and then I use specific mapping annotations like @GetMapping, @PostMapping, @PutMapping, and @DeleteMapping for each HTTP method. For parameters, @PathVariable extracts values from the URL path, @RequestParam handles query parameters, and @RequestBody binds JSON to my DTO objects. I always wrap responses in ResponseEntity to have full control over HTTP status codes and headers. For example, I return 201 Created with a Location header for POST requests, and 204 No Content for successful DELETEs."

---

### Q2. How do you handle request validation?

```java
// DTO with Bean Validation annotations
public class ProductDTO {
    private Long id;

    @NotBlank(message = "Name cannot be blank")
    @Size(min = 2, max = 100, message = "Name must be 2-100 characters")
    private String name;

    @NotNull(message = "Price is required")
    @DecimalMin(value = "0.01", message = "Price must be positive")
    @Digits(integer = 10, fraction = 2)
    private BigDecimal price;

    @Min(value = 0, message = "Stock cannot be negative")
    private int stock;

    @Email(message = "Invalid supplier email")
    private String supplierEmail;

    @NotNull
    @Future(message = "Expiry must be in the future")
    private LocalDate expiryDate;
    
    // getters + setters
}

// Controller — trigger validation with @Valid
@PostMapping
public ResponseEntity<ProductDTO> create(@Valid @RequestBody ProductDTO dto,
                                          BindingResult result) {
    // BindingResult (optional) — if present, Spring won't auto-throw
    if (result.hasErrors()) {
        // handle manually
    }
    return ResponseEntity.ok(productService.create(dto));
}
// Without BindingResult — Spring throws MethodArgumentNotValidException automatically
```

---

### 🎯 How to Explain in Interview

"Request validation in Spring Boot is handled through Bean Validation annotations. I decorate my DTO fields with annotations like @NotBlank, @Size, @NotNull, @Min, and @Email to define validation rules. When I add @Valid to my controller method parameter, Spring automatically validates the incoming JSON against these rules. If validation fails, Spring throws MethodArgumentNotValidException by default. I can catch this globally with @ControllerAdvice to return consistent error responses. The beauty is that validation is declarative - I just declare the rules and Spring enforces them. I can also create custom validation annotations for complex business rules. This keeps my controller code clean and validation logic reusable across different endpoints."

---

### Q3. How do you handle exceptions globally?

```java
// Custom exception
public class ResourceNotFoundException extends RuntimeException {
    public ResourceNotFoundException(String message) { super(message); }
}

// Error response DTO
public record ErrorResponse(int status, String message, LocalDateTime timestamp) {}

// Global exception handler
@RestControllerAdvice   // = @ControllerAdvice + @ResponseBody
public class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleNotFound(ResourceNotFoundException ex) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
            .body(new ErrorResponse(404, ex.getMessage(), LocalDateTime.now()));
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, String>> handleValidation(
            MethodArgumentNotValidException ex) {
        Map<String, String> errors = new HashMap<>();
        ex.getBindingResult().getFieldErrors()
            .forEach(err -> errors.put(err.getField(), err.getDefaultMessage()));
        return ResponseEntity.badRequest().body(errors);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGeneral(Exception ex) {
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
            .body(new ErrorResponse(500, "Internal Server Error", LocalDateTime.now()));
    }
}
```

---

### 🎯 How to Explain in Interview

"Global exception handling with @RestControllerAdvice is essential for clean REST APIs. Instead of scattering try-catch blocks throughout my controllers, I create a centralized exception handler. @RestControllerAdvice combines @ControllerAdvice and @ResponseBody, so it handles exceptions across all controllers and returns JSON responses. I define methods with @ExceptionHandler for specific exception types - like ResourceNotFoundException returning 404, or MethodArgumentNotValidException for validation errors returning 400. I also include a catch-all handler for generic exceptions returning 500. This approach ensures consistent error responses across my entire API and makes the code much cleaner. I can create custom exception classes for different business scenarios and handle them all in one place."

---

### Q4. What are common HTTP status codes?

| Code | Name | When to Use |
|------|------|-------------|
| 200 | OK | Successful GET, PUT, PATCH |
| 201 | Created | Successful POST (resource created) |
| 204 | No Content | Successful DELETE |
| 400 | Bad Request | Validation errors, malformed request |
| 401 | Unauthorized | Not authenticated |
| 403 | Forbidden | Authenticated but no permission |
| 404 | Not Found | Resource doesn't exist |
| 409 | Conflict | Duplicate resource, concurrency conflict |
| 422 | Unprocessable Entity | Semantic validation failure |
| 429 | Too Many Requests | Rate limiting |
| 500 | Internal Server Error | Unexpected server error |
| 503 | Service Unavailable | Server down / maintenance |

---

### 🎯 How to Explain in Interview

"HTTP status codes are the language of REST APIs - they tell clients what happened with their requests. I use 200 for successful operations like GET or PUT, 201 when I create a new resource with POST, and 204 for successful DELETEs where there's nothing to return. For client errors, 400 is for bad requests like validation failures, 401 when someone isn't authenticated, 403 when they're authenticated but don't have permission, and 404 when a resource doesn't exist. I use 409 for conflicts like duplicate creations, 422 for semantic validation failures, and 429 for rate limiting. For server errors, 500 is for unexpected problems, and 503 when the service is temporarily unavailable. Using the right status codes makes my API predictable and easy to work with."

---

### Q5. How do you customize JSON serialization with Jackson?

```java
// Jackson annotations
public class UserDTO {

    @JsonProperty("user_id")     // different JSON key
    private Long id;

    @JsonIgnore                  // exclude from JSON
    private String password;

    @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")  // date format
    private LocalDateTime createdAt;

    @JsonInclude(JsonInclude.Include.NON_NULL)     // skip null fields
    private String middleName;

    @JsonSerialize(using = PriceSerializer.class)  // custom serializer
    private BigDecimal price;
}

// Custom serializer
public class PriceSerializer extends JsonSerializer<BigDecimal> {
    @Override
    public void serialize(BigDecimal value, JsonGenerator gen, SerializerProvider sp)
            throws IOException {
        gen.writeString("$" + value.setScale(2, RoundingMode.HALF_UP));
    }
}

// Global ObjectMapper configuration
@Configuration
public class JacksonConfig {
    @Bean
    public ObjectMapper objectMapper() {
        return new ObjectMapper()
            .registerModule(new JavaTimeModule())                           // Java 8 dates
            .disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS)       // ISO-8601 dates
            .setSerializationInclusion(JsonInclude.Include.NON_NULL)       // skip nulls globally
            .configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
    }
}
```

---

### 🎯 How to Explain in Interview

"Jackson JSON serialization is highly customizable in Spring Boot. I use annotations like @JsonProperty to change JSON field names, @JsonIgnore to exclude sensitive data like passwords, and @JsonFormat to control date formatting. @JsonInclude lets me skip null fields to keep responses clean. For complex formatting, I can create custom serializers by extending JsonSerializer - perfect for formatting prices as currency strings. I also configure the global ObjectMapper in a @Configuration class to set defaults like ISO-8601 date formatting, skipping nulls globally, and being lenient about unknown JSON properties. This level of control ensures my API responses are exactly how I want them, whether that means clean, minimal JSON or rich, formatted data."

---

### Q6. How do you implement pagination and sorting?

```java
// Using Spring Data's Pageable
@GetMapping
public ResponseEntity<Page<ProductDTO>> getProducts(
        @RequestParam(defaultValue = "0") int page,
        @RequestParam(defaultValue = "20") int size,
        @RequestParam(defaultValue = "name") String sortBy,
        @RequestParam(defaultValue = "asc") String direction) {

    Sort sort = direction.equalsIgnoreCase("desc")
        ? Sort.by(sortBy).descending()
        : Sort.by(sortBy).ascending();

    Pageable pageable = PageRequest.of(page, size, sort);
    Page<Product> products = productRepository.findAll(pageable);

    return ResponseEntity.ok(products.map(productMapper::toDTO));
}

// Response includes: content[], totalElements, totalPages, size, number, last, first
// GET /api/products?page=0&size=10&sortBy=price&direction=desc
```

---

### 🎯 How to Explain in Interview

"Pagination and sorting are crucial for APIs that return large datasets. Spring Data makes this elegant with the Pageable interface. I accept page, size, sortBy, and direction parameters, then create a PageRequest object that encapsulates all this information. The beauty is that Spring Data repositories return a Page object which contains not just the data, but also metadata like total elements, total pages, and whether it's the first or last page. This makes it easy for clients to implement pagination controls. For sorting, I use the Sort class which supports multiple sort criteria and ascending/descending order. The response includes everything needed for pagination UI - current page, total pages, and navigation links."

---

### Q7. How do you implement request/response logging?

```java
// Using HandlerInterceptor
@Component
public class LoggingInterceptor implements HandlerInterceptor {

    private static final Logger log = LoggerFactory.getLogger(LoggingInterceptor.class);

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response,
                             Object handler) {
        log.info("[{}] {} — IP: {}", request.getMethod(),
            request.getRequestURI(), request.getRemoteAddr());
        request.setAttribute("startTime", System.currentTimeMillis());
        return true;  // continue processing
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response,
                                Object handler, Exception ex) {
        long start = (Long) request.getAttribute("startTime");
        log.info("Completed [{}] {} — Status: {} in {}ms",
            request.getMethod(), request.getRequestURI(),
            response.getStatus(), System.currentTimeMillis() - start);
    }
}

@Configuration
public class WebConfig implements WebMvcConfigurer {
    private final LoggingInterceptor loggingInterceptor;
    
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(loggingInterceptor).addPathPatterns("/api/**");
    }
}
```

---

### 🎯 How to Explain in Interview

"Request/response logging is essential for monitoring and debugging REST APIs. I implement this using Spring's HandlerInterceptor mechanism. In the preHandle method, I log the incoming request method, URI, and client IP, and store the start time. In afterCompletion, I calculate the total duration and log the response status. This gives me a complete picture of each request's journey through my application. The interceptor can be selectively applied to specific URL patterns using addPathPatterns, so I can focus on API endpoints while ignoring static resources. This approach is much cleaner than adding logging to every controller method and provides consistent logging across the entire API."

---

### Q8. What are REST API best practices?

```text
1. Use nouns for resources:      GET /users, NOT GET /getUsers
2. Use HTTP verbs semantically:  POST=create, PUT=replace, PATCH=partial, DELETE=remove
3. Version your API:             /api/v1/users, /api/v2/users
4. Return proper HTTP codes:     201 for create, 204 for delete, 404 for not found
5. Use plural for resources:     /products, /orders (not /product, /order)
6. Filter, sort, paginate:       GET /products?category=shoes&page=0&size=20&sort=price
7. Use DTOs, not entities:       Never expose JPA entities directly
8. Consistent error format:      Always return {status, message, timestamp, errors[]}
9. HATEOAS (optional):           Include links to related resources
10. Secure with Spring Security: JWT or OAuth2 for authentication/authorization
```

```java
// HATEOAS example (Spring HATEOAS)
@GetMapping("/{id}")
public EntityModel<ProductDTO> getById(@PathVariable Long id) {
    ProductDTO dto = productService.findById(id);
    return EntityModel.of(dto,
        linkTo(methodOn(ProductController.class).getById(id)).withSelfRel(),
        linkTo(methodOn(ProductController.class).getAll(0, 20)).withRel("products")
    );
}
// Response includes: {"id":1, "name":"...", "_links":{"self":{"href":"/api/products/1"}}}
```

---

### 🎯 How to Explain in Interview

"REST API best practices are about creating APIs that are intuitive, consistent, and maintainable. I use nouns for resource names like /users instead of verbs like /getUsers. I use HTTP verbs semantically - POST for creating, PUT for replacing, PATCH for partial updates, and DELETE for removing. I always version my APIs with /api/v1/ to allow for future changes without breaking existing clients. I return appropriate HTTP status codes and use plural nouns for collections. I implement filtering, sorting, and pagination to handle large datasets efficiently. I never expose JPA entities directly - always use DTOs. I maintain consistent error response formats and secure the API with Spring Security using JWT or OAuth2. These practices make APIs that developers love to work with."
