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
