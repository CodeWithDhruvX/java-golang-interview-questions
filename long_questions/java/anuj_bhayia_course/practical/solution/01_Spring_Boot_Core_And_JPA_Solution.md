# Solution: Spring Boot Core, Web, and Data JPA

## 1. REST API Setup: Library Management System

**Solution:**

```java
// Entity
import jakarta.persistence.*;
import jakarta.validation.constraints.*;
import lombok.Data;

@Entity
@Data
public class Book {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(unique = true, nullable = false, length = 13)
    private String isbn;

    @NotBlank
    private String title;

    private String author;

    @Min(0)
    private int availableCopies;
}

// DTO
@Data
public class BookDto {
    @Size(min = 13, max = 13, message = "ISBN must be exactly 13 characters")
    private String isbn;

    @NotBlank(message = "Title cannot be blank")
    private String title;

    private String author;

    @Min(value = 0, message = "Available copies cannot be negative")
    private int availableCopies;
}

// Repository
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface BookRepository extends JpaRepository<Book, Long> {
}

// Service
import org.springframework.stereotype.Service;

@Service
public class BookService {

    private final BookRepository bookRepository;

    public BookService(BookRepository bookRepository) {
        this.bookRepository = bookRepository;
    }

    public Book saveBook(BookDto bookDto) {
        Book book = new Book();
        book.setIsbn(bookDto.getIsbn());
        book.setTitle(bookDto.getTitle());
        book.setAuthor(bookDto.getAuthor());
        book.setAvailableCopies(bookDto.getAvailableCopies());
        
        return bookRepository.save(book);
    }

    public Book getBookById(Long id) {
        return bookRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Book with ID " + id + " not found."));
    }
}

// Controller
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import jakarta.validation.Valid;

@RestController
@RequestMapping("/books")
public class BookController {

    private final BookService bookService;

    public BookController(BookService bookService) {
        this.bookService = bookService;
    }

    @PostMapping
    public ResponseEntity<Book> createBook(@Valid @RequestBody BookDto bookDto) {
        Book savedBook = bookService.saveBook(bookDto);
        return new ResponseEntity<>(savedBook, HttpStatus.CREATED);
    }

    @GetMapping("/{id}")
    public ResponseEntity<Book> getBook(@PathVariable Long id) {
        Book book = bookService.getBookById(id);
        return ResponseEntity.ok(book);
    }
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }
}
```

---

## 2. Global Exception Handling

**Solution:**

```java
// 1. Custom Exception
public class ResourceNotFoundException extends RuntimeException {
    public ResourceNotFoundException(String message) {
        super(message);
    }
}

// 2. Error Response DTO
@Data
@AllArgsConstructor
public class ErrorResponse {
    private LocalDateTime timestamp;
    private int status;
    private String error;
    private String message;
}

// 3. Global Exception Handler
@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleResourceNotFound(ResourceNotFoundException ex) {
        ErrorResponse error = new ErrorResponse(
                LocalDateTime.now(),
                HttpStatus.NOT_FOUND.value(),
                HttpStatus.NOT_FOUND.getReasonPhrase(),
                ex.getMessage()
        );
        return new ResponseEntity<>(error, HttpStatus.NOT_FOUND);
    }
    
    // Bonus: Handle Validation Errors from @Valid
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, String>> handleValidationExceptions(MethodArgumentNotValidException ex) {
        Map<String, String> errors = new HashMap<>();
        ex.getBindingResult().getFieldErrors().forEach(error -> 
            errors.put(error.getField(), error.getDefaultMessage()));
        return new ResponseEntity<>(errors, HttpStatus.BAD_REQUEST);
    }
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }
}
```

---

## 3. Spring Data JPA: Relational Mappings & Custom Queries

**Solution:**

```java
// 1. Entity Mappings
@Entity
@Data
public class LibraryMember {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    private String name;

    // LibraryMember is the non-owning side of the relationship
    @OneToMany(mappedBy = "member", cascade = CascadeType.ALL)
    private List<Book> borrowedBooks = new ArrayList<>();
}

@Entity
@Data
public class Book {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String title;
    private String author;
    private int availableCopies;

    // Book is the owning side because it has the foreign key column "member_id"
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "member_id")
    private LibraryMember member;
}

// 2 & 3 & 4. Repository Interface
public interface BookRepository extends JpaRepository<Book, Long> {
    
    // 3. Derived Query
    List<Book> findByAuthorAndAvailableCopiesGreaterThan(String author, int copies);
}

public interface LibraryMemberRepository extends JpaRepository<LibraryMember, Long> {

    // 4. Custom JPQL Query
    @Query("SELECT DISTINCT m FROM LibraryMember m JOIN m.borrowedBooks b")
    List<LibraryMember> findMembersWithBorrowedBooks();
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }
}
```

---

## 4. Transaction Management & Propagation (`REQUIRES_NEW`)

**Solution:**

```java
@Service
public class AuditLogService {

    @Autowired
    private AuditRepository auditRepository;

    // REQUIRES_NEW suspends the outer transaction and creates a new one
    // so this always commits, even if the outer transaction rolls back.
    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void logAttempt(String message) {
        AuditLog log = new AuditLog(message, LocalDateTime.now());
        auditRepository.save(log);
    }
}

@Service
public class OrderService {

    @Autowired
    private InventoryService inventoryService;
    @Autowired
    private AuditLogService auditLogService;

    // REQUIRED is the default. If an exception triggers rollback here,
    // the audit log will still persist because it ran in its own isolated transaction.
    @Transactional(propagation = Propagation.REQUIRED)
    public void checkout(Long itemId, int quantity) {
        
        // 1. Log the attempt (executes in completely separate transaction)
        auditLogService.logAttempt("User attempted checkout for item " + itemId);
        
        // 2. Process order (if this throws OutOfStockException, this specific transaction rolls back)
        inventoryService.deductStock(itemId, quantity);
        
        // 3. Place order logic...
    }
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.transaction.annotation.EnableTransactionManagement;

@SpringBootApplication
@EnableTransactionManagement
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }
}
```

---

## 5. Spring Boot Actuator & Custom Health Indicator

**Solution:**

```xml
<!-- In pom.xml -->
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
```

```java
import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

@Component
public class StripeApiHealthIndicator implements HealthIndicator {

    private final RestTemplate restTemplate;

    public StripeApiHealthIndicator(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    @Override
    public Health health() {
        try {
            // Ping Stripe API
            String response = restTemplate.getForObject("https://api.stripe.com/healthcheck", String.class);
            
            if (response != null) {
                return Health.up().withDetail("Stripe API", "Reachable and functioning").build();
            }
        } catch (Exception ex) {
            return Health.down().withDetail("Stripe API", "Unreachable").withException(ex).build();
        }
        return Health.down().withDetail("Stripe API", "Unknown Error").build();
    }
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;

@SpringBootApplication
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
```

---

## 6. Consuming External REST APIs (RestTemplate)

**Solution:**

```java
// 1. DTO Structure based on JSON response
@Data
public class WeatherResponse {
    private Location location;
    private Current current;

    @Data
    public static class Location {
        private String name;
        private String country;
    }

    @Data
    public static class Current {
        private double temp_c;
        private String condition;
    }
}

// 2 & 3 & 4. Service Implementation
@Service
public class WeatherService {

    private final RestTemplate restTemplate;
    private final String API_KEY = "your_api_key_here";

    public WeatherService(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    public String getCurrentTemperature(String city) {
        String url = "https://api.weatherapi.com/v1/current.json?key=" + API_KEY + "&q=" + city;

        try {
            WeatherResponse response = restTemplate.getForObject(url, WeatherResponse.class);
            
            if (response != null && response.getCurrent() != null) {
                return "The temperature in " + response.getLocation().getName() + 
                       " is " + response.getCurrent().getTemp_c() + "°C.";
            }
            return "Unable to parse temperature data.";
            
        } catch (HttpClientErrorException.NotFound ex) {
            // Handle HTTP 404 from external API gracefully
            return "City not found: " + city;
        } catch (Exception ex) {
            return "Error calling weather service: " + ex.getMessage();
        }
    }
}

// Main Application
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;

@SpringBootApplication
public class PracticalInterviewApplication {
    public static void main(String[] args) {
        SpringApplication.run(PracticalInterviewApplication.class, args);
    }

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
```
