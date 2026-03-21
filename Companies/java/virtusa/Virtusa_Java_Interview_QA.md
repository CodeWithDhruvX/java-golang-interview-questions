# Virtusa Java Developer Interview Questions & Answers

## Table of Contents
1. [Core Java Concepts](#core-java-concepts)
2. [Collections Framework](#collections-framework)
3. [Java 8+ Features](#java-8-features)
4. [Spring Framework](#spring-framework)
5. [REST API Development](#rest-api-development)
6. [Database & SQL](#database--sql)
7. [Multithreading & Concurrency](#multithreading--concurrency)
8. [Microservices Architecture](#microservices-architecture)
9. [DevOps & Cloud](#devops--cloud)
10. [Coding Problems](#coding-problems)

---

## Core Java Concepts

### Q1: What is immutability in Java and why is String immutable?

**Answer:** Immutability means that once an object is created, its state cannot be changed after construction. In Java, String is immutable for several important reasons:

1. **Security**: Strings are used for sensitive data like passwords, file paths, and network connections. Immutability prevents accidental or malicious modifications.

2. **Thread Safety**: Immutable objects are inherently thread-safe since their state cannot change, eliminating the need for synchronization in multi-threaded environments.

3. **Caching**: The JVM can cache String literals in the string pool, saving memory when the same string is used multiple times.

4. **HashCode Stability**: The hashCode of an immutable object remains constant, making it ideal for use as keys in HashMap.

**Example:**
```java
String str1 = "Hello";
String str2 = str1.concat(" World"); // Creates new String object
System.out.println(str1); // Still "Hello" - original unchanged
System.out.println(str2); // "Hello World" - new object
```

### Q2: Explain generics in Java and their benefits.

**Answer:** Generics provide a way to create classes, interfaces, and methods that can work with different data types while maintaining type safety at compile time.

**Key Benefits:**
1. **Type Safety**: Compile-time checking prevents runtime ClassCastException
2. **Code Reusability**: Write generic code that works with multiple types
3. **Elimination of Casting**: No need for explicit type casting
4. **Better IDE Support**: Enhanced code completion and refactoring

**Example:**
```java
// Before generics (requires casting)
List list = new ArrayList();
list.add("Hello");
String str = (String) list.get(0); // Casting required

// With generics (type-safe)
List<String> list = new ArrayList<>();
list.add("Hello");
String str = list.get(0); // No casting needed
```

### Q3: What are the SOLID principles in object-oriented design?

**Answer:** SOLID is an acronym for five design principles:

1. **S - Single Responsibility Principle**: A class should have only one reason to change.
   ```java
   // Good: Separate concerns
   class UserService { void createUser(); }
   class EmailService { void sendEmail(); }
   ```

2. **O - Open/Closed Principle**: Software entities should be open for extension but closed for modification.
   ```java
   interface PaymentProcessor { void processPayment(); }
   class CreditCardProcessor implements PaymentProcessor { /* implementation */ }
   class PayPalProcessor implements PaymentProcessor { /* implementation */ }
   ```

3. **L - Liskov Substitution Principle**: Subtypes must be substitutable for their base types.
   ```java
   // If Rectangle extends Square, LSP is violated
   // Square and Rectangle have different behaviors
   ```

4. **I - Interface Segregation Principle**: Clients shouldn't be forced to depend on interfaces they don't use.
   ```java
   // Good: Specific interfaces
   interface Printable { void print(); }
   interface Readable { String read(); }
   ```

5. **D - Dependency Inversion Principle**: High-level modules shouldn't depend on low-level modules.
   ```java
   // Good: Depend on abstractions
   class OrderService {
       private PaymentProcessor processor; // Interface, not concrete class
   }
   ```

### Q4: What is ConcurrentModificationException and how to avoid it?

**Answer:** ConcurrentModificationException occurs when a collection is modified while being iterated, typically using a non-concurrent iterator.

**Common Causes:**
- Removing elements from a collection while iterating with enhanced for-loop
- Adding elements during iteration
- Multiple threads modifying the same collection simultaneously

**Solutions:**

1. **Use Iterator's remove method:**
```java
List<String> list = new ArrayList<>();
Iterator<String> iterator = list.iterator();
while (iterator.hasNext()) {
    String item = iterator.next();
    if (item.equals("remove")) {
        iterator.remove(); // Safe removal
    }
}
```

2. **Use removeIf method (Java 8+):**
```java
list.removeIf(item -> item.equals("remove"));
```

3. **Use concurrent collections:**
```java
ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
// Thread-safe iteration and modification
```

---

## Collections Framework

### Q5: Difference between ArrayList vs LinkedList

**Answer:**

| Aspect | ArrayList | LinkedList |
|--------|-----------|------------|
| **Internal Structure** | Dynamic array | Doubly-linked list |
| **Performance** | O(1) for random access | O(n) for random access |
| **Add/Remove** | O(n) (requires shifting) | O(1) for ends, O(n) for middle |
| **Memory** | Less overhead per element | More overhead (prev/next pointers) |
| **Use Case** | Frequent random access | Frequent insertions/deletions |

**When to use:**
- **ArrayList**: When you need fast random access and fewer modifications
- **LinkedList**: When you need frequent insertions/deletions at ends

### Q6: Difference between HashMap and ConcurrentHashMap

**Answer:**

| Feature | HashMap | ConcurrentHashMap |
|---------|---------|---------------------|
| **Thread Safety** | Not thread-safe | Thread-safe |
| **Null Keys/Values** | Allows one null key, multiple null values | Doesn't allow null keys/values |
| **Performance** | Faster in single-threaded | Slightly slower due to synchronization |
| **Iterator** | Fail-fast | Weakly consistent |

**ConcurrentHashMap Internal Working:**
```java
// Uses segment-based locking (Java 7) or CAS + synchronized (Java 8+)
// Java 8+ implementation:
ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
map.put("key", 1); // Uses CAS operations
map.compute("key", (k, v) -> v == null ? 1 : v + 1); // Atomic operations
```

---

## Java 8+ Features

### Q7: Explain Optional.of(), Optional.ofNullable(), Optional.empty()

**Answer:** Optional is a container object that may or may not contain a non-null value, designed to eliminate NullPointerException.

1. **Optional.of(value)**: Creates Optional with non-null value
   ```java
   String name = "John";
   Optional<String> opt = Optional.of(name); // Throws NPE if name is null
   ```

2. **Optional.ofNullable(value)**: Creates Optional that may contain null
   ```java
   String name = null;
   Optional<String> opt = Optional.ofNullable(name); // Creates empty Optional
   ```

3. **Optional.empty()**: Creates empty Optional
   ```java
   Optional<String> empty = Optional.empty();
   ```

**Best Practices:**
```java
// Good: Use Optional as return type from methods
public Optional<User> findUserById(int id) {
    User user = repository.findById(id);
    return Optional.ofNullable(user);
}

// Good: Chain operations
String userName = findUserById(123)
    .map(User::getName)
    .orElse("Unknown");

// Avoid: Using Optional for fields
```

### Q8: What is the difference between Lambda expressions and Functional Interfaces?

**Answer:**

**Functional Interface:**
- An interface with exactly one abstract method
- Annotated with @FunctionalInterface (optional but recommended)
- Examples: Runnable, Comparator, Function, Predicate

```java
@FunctionalInterface
interface Calculator {
    int calculate(int a, int b);
}
```

**Lambda Expression:**
- Anonymous implementation of functional interface
- Concise syntax for writing functional interfaces
- Enables functional programming in Java

```java
// Traditional approach
Calculator add = new Calculator() {
    @Override
    public int calculate(int a, int b) {
        return a + b;
    }
};

// Lambda expression
Calculator add = (a, b) -> a + b;

// Method reference
Calculator add = Integer::sum;
```

**Key Differences:**
- Functional Interface is the target type
- Lambda is the implementation
- Lambda cannot exist without a functional interface target

---

## Spring Framework

### Q9: Difference between Spring Framework vs Spring Boot

**Answer:**

| Aspect | Spring Framework | Spring Boot |
|--------|------------------|-------------|
| **Purpose** | Provides comprehensive programming and configuration model | Makes it easy to create stand-alone, production-ready applications |
| **Configuration** | Requires XML or Java configuration | Auto-configuration, minimal configuration |
| **Dependencies** | Manual dependency management | Starter dependencies with managed versions |
| **Embedded Server** | Requires external server setup | Built-in embedded servers (Tomcat, Jetty) |
| **Development Speed** | Slower setup | Rapid development |

**Spring Boot Advantages:**
```java
// Spring Boot - Minimal configuration
@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}

@RestController
public class UserController {
    @Autowired
    private UserService userService;
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable Long id) {
        return userService.findById(id);
    }
}
```

### Q10: Difference between @Component vs @Service vs @Repository

**Answer:** All are stereotype annotations that mark classes as Spring beans, but with different semantic meanings:

**@Component:**
- Generic stereotype for any Spring-managed component
- Base annotation for other stereotypes
```java
@Component
class EmailValidator {
    public boolean isValid(String email) { /* ... */ }
}
```

**@Service:**
- Specialized for service layer components
- Indicates business logic
- Good for transaction management
```java
@Service
public class UserService {
    @Transactional
    public User createUser(UserDto dto) { /* business logic */ }
}
```

**@Repository:**
- Specialized for data access layer
- Enables exception translation (PersistenceException → DataAccessException)
- Indicates DAO pattern
```java
@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    // Spring Data JPA methods
}
```

### Q11: Difference between @Controller and @RestController

**Answer:**

**@Controller:**
- Traditional Spring MVC controller
- Returns ModelAndView or String (view name)
- Requires @ResponseBody for REST responses
```java
@Controller
public class WebController {
    @GetMapping("/home")
    public String home(Model model) {
        model.addAttribute("message", "Welcome!");
        return "home"; // Returns view name
    }
    
    @GetMapping("/api/data")
    @ResponseBody // Required for REST
    public String getData() {
        return "data";
    }
}
```

**@RestController:**
- Combination of @Controller + @ResponseBody
- All methods return JSON/XML by default
- Designed for REST APIs
```java
@RestController
public class ApiController {
    @GetMapping("/api/users")
    public List<User> getUsers() {
        return userService.findAll(); // Automatically serialized to JSON
    }
}
```

---

## REST API Development

### Q12: Difference between @RequestParam vs @PathVariable

**Answer:**

**@RequestParam:**
- Extracts query parameters from URL
- Used for optional parameters, filtering, pagination
```java
// GET /api/users?page=1&size=10&sort=name
@GetMapping("/users")
public Page<User> getUsers(
    @RequestParam(defaultValue = "0") int page,
    @RequestParam(defaultValue = "10") int size,
    @RequestParam(required = false) String sort) {
    return userService.findAll(page, size, sort);
}
```

**@PathVariable:**
- Extracts values from URL path
- Used for resource identification
```java
// GET /api/users/123/orders/456
@GetMapping("/users/{userId}/orders/{orderId}")
public Order getOrder(
    @PathVariable Long userId,
    @PathVariable Long orderId) {
    return orderService.findByUserIdAndOrderId(userId, orderId);
}
```

**Best Practices:**
- Use @PathVariable for required resource identifiers
- Use @RequestParam for optional filtering and pagination
- Combine both for complex queries

### Q13: How do you handle global exceptions in Spring Boot?

**Answer:** Use @ControllerAdvice and @ExceptionHandler for centralized exception handling:

```java
@ControllerAdvice
public class GlobalExceptionHandler {
    
    // Handle specific exceptions
    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleResourceNotFound(
            ResourceNotFoundException ex) {
        ErrorResponse error = new ErrorResponse(
            "NOT_FOUND",
            ex.getMessage(),
            LocalDateTime.now()
        );
        return new ResponseEntity<>(error, HttpStatus.NOT_FOUND);
    }
    
    // Handle validation exceptions
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleValidationExceptions(
            MethodArgumentNotValidException ex) {
        Map<String, String> errors = new HashMap<>();
        ex.getBindingResult().getFieldErrors().forEach(error ->
            errors.put(error.getField(), error.getDefaultMessage())
        );
        
        ErrorResponse error = new ErrorResponse(
            "VALIDATION_ERROR",
            "Invalid input parameters",
            errors
        );
        return new ResponseEntity<>(error, HttpStatus.BAD_REQUEST);
    }
    
    // Handle all other exceptions
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleAllExceptions(Exception ex) {
        ErrorResponse error = new ErrorResponse(
            "INTERNAL_SERVER_ERROR",
            "An unexpected error occurred",
            LocalDateTime.now()
        );
        return new ResponseEntity<>(error, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
```

### Q14: How do you secure a Spring Boot application?

**Answer:** Implement multiple layers of security:

**1. Spring Security Configuration:**
```java
@Configuration
@EnableWebSecurity
public class SecurityConfig {
    
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .csrf().disable()
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/api/public/**").permitAll()
                .requestMatchers("/api/admin/**").hasRole("ADMIN")
                .anyRequest().authenticated()
            )
            .sessionManagement(session -> session
                .sessionCreationPolicy(SessionCreationPolicy.STATELESS)
            )
            .addFilterBefore(jwtAuthenticationFilter, 
                UsernamePasswordAuthenticationFilter.class);
        
        return http.build();
    }
}
```

**2. JWT Authentication:**
```java
@Component
public class JwtAuthenticationFilter extends OncePerRequestFilter {
    
    @Override
    protected void doFilterInternal(HttpServletRequest request, 
                                   HttpServletResponse response, 
                                   FilterChain filterChain) {
        String token = extractToken(request);
        if (token != null && jwtUtil.validateToken(token)) {
            UsernamePasswordAuthenticationToken authentication = 
                new UsernamePasswordAuthenticationToken(
                    jwtUtil.getUsername(token), null, 
                    jwtUtil.getAuthorities(token)
                );
            SecurityContextHolder.getContext().setAuthentication(authentication);
        }
        filterChain.doFilter(request, response);
    }
}
```

**3. Input Validation:**
```java
public class UserDto {
    @NotBlank(message = "Name is required")
    @Size(min = 2, max = 50)
    private String name;
    
    @Email(message = "Invalid email format")
    private String email;
    
    @Pattern(regexp = "^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$",
             message = "Password must be at least 8 characters with letters and numbers")
    private String password;
}
```

---

## Database & SQL

### Q15: What is Normalization in SQL?

**Answer:** Normalization is the process of organizing database tables to reduce data redundancy and improve data integrity.

**Normal Forms:**

1. **First Normal Form (1NF):** Each cell contains atomic values
   ```sql
   -- Before 1NF
   CREATE TABLE Orders (
       order_id INT,
       customer_name VARCHAR(100),
       products VARCHAR(500) -- "Laptop,Mouse,Keyboard"
   );
   
   -- After 1NF
   CREATE TABLE Orders (
       order_id INT,
       customer_name VARCHAR(100)
   );
   CREATE TABLE Order_Items (
       order_id INT,
       product_name VARCHAR(100)
       -- Each product in separate row
   );
   ```

2. **Second Normal Form (2NF):** No partial dependencies (only applies to composite keys)
3. **Third Normal Form (3NF):** No transitive dependencies

**Benefits:**
- Eliminates data redundancy
- Ensures data consistency
- Improves database organization
- Facilitates easier maintenance

### Q16: Difference between DELETE, TRUNCATE, and DROP in SQL

**Answer:**

| Command | Purpose | WHERE Clause | Transaction | Speed | Identity Reset |
|---------|---------|--------------|-------------|-------|----------------|
| **DELETE** | Removes specific rows | Yes | Can be rolled back | Slower | No |
| **TRUNCATE** | Removes all rows | No | Cannot be rolled back | Faster | Yes |
| **DROP** | Removes table entirely | N/A | Cannot be rolled back | Fastest | N/A |

**Examples:**
```sql
-- DELETE: Removes specific rows, can use WHERE
DELETE FROM users WHERE status = 'inactive';
-- Can be rolled back: ROLLBACK;

-- TRUNCATE: Removes all rows, cannot use WHERE
TRUNCATE TABLE users;
-- Cannot be rolled back in most databases

-- DROP: Removes entire table structure
DROP TABLE users;
-- Cannot be rolled back
```

---

## Multithreading & Concurrency

### Q17: Why are immutable objects thread-safe?

**Answer:** Immutable objects are thread-safe because:

1. **No State Changes**: Once created, their state cannot be modified
2. **No Synchronization Needed**: Multiple threads can access simultaneously without conflicts
3. **Safe Publication**: Can be safely shared between threads without additional synchronization

**Example:**
```java
// Immutable class
public final class ImmutablePerson {
    private final String name;
    private final int age;
    private final List<String> hobbies; // Defensive copy needed
    
    public ImmutablePerson(String name, int age, List<String> hobbies) {
        this.name = name;
        this.age = age;
        this.hobbies = Collections.unmodifiableList(new ArrayList<>(hobbies));
    }
    
    public String getName() { return name; }
    public int getAge() { return age; }
    public List<String> getHobbies() { return hobbies; }
    
    // No setter methods
}

// Thread-safe usage
public class SharedResource {
    private final ImmutablePerson person = new ImmutablePerson("John", 30, 
        Arrays.asList("Reading", "Swimming"));
    
    // Multiple threads can safely access person
    public void processPerson() {
        String name = person.getName(); // Thread-safe
        List<String> hobbies = person.getHobbies(); // Thread-safe
    }
}
```

### Q18: How do you safely shutdown a thread pool in production?

**Answer:** Use proper shutdown sequence to ensure graceful termination:

```java
public class ThreadPoolManager {
    private ExecutorService executorService;
    
    public void initialize() {
        executorService = Executors.newFixedThreadPool(10);
    }
    
    public void shutdown() {
        // 1. Stop accepting new tasks
        executorService.shutdown();
        
        try {
            // 2. Wait for existing tasks to complete
            if (!executorService.awaitTermination(60, TimeUnit.SECONDS)) {
                // 3. Force shutdown if tasks don't complete
                executorService.shutdownNow();
                
                // 4. Wait for tasks to respond to cancellation
                if (!executorService.awaitTermination(60, TimeUnit.SECONDS)) {
                    System.err.println("Thread pool did not terminate");
                }
            }
        } catch (InterruptedException e) {
            // 5. Re-interrupt if thread was interrupted
            executorService.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }
}
```

**Best Practices:**
- Use shutdown() for graceful termination
- Use shutdownNow() for immediate termination
- Handle InterruptedException properly
- Monitor task completion
- Consider using try-with-resources for AutoCloseable executors

---

## Microservices Architecture

### Q19: Why do we use microservices architecture?

**Answer:** Microservices architecture offers several advantages:

**Key Benefits:**
1. **Scalability**: Independent scaling of services based on demand
2. **Technology Diversity**: Different services can use different technologies
3. **Fault Isolation**: Failure in one service doesn't affect others
4. **Team Autonomy**: Teams can develop and deploy independently
5. **Faster Deployment**: Smaller codebase enables quicker releases

**Example Scenario:**
```java
// E-commerce Microservices
// 1. User Service
@RestController
@RequestMapping("/api/users")
public class UserController {
    // Handles user management
}

// 2. Product Service  
@RestController
@RequestMapping("/api/products")
public class ProductController {
    // Handles product catalog
}

// 3. Order Service
@RestController
@RequestMapping("/api/orders")
public class OrderController {
    // Handles order processing
}

// 4. Payment Service
@RestController
@RequestMapping("/api/payments")
public class PaymentController {
    // Handles payment processing
}
```

### Q20: If multiple microservices are communicating and one service goes down, how do you handle it?

**Answer:** Implement several resilience patterns:

**1. Circuit Breaker Pattern:**
```java
@Service
public class OrderService {
    
    @CircuitBreaker(name = "paymentService", fallbackMethod = "fallbackPayment")
    public PaymentResponse processPayment(PaymentRequest request) {
        return paymentClient.processPayment(request);
    }
    
    public PaymentResponse fallbackPayment(PaymentRequest request, Exception ex) {
        // Fallback logic when payment service is down
        return PaymentResponse.builder()
            .status("PENDING")
            .message("Payment service temporarily unavailable")
            .build();
    }
}
```

**2. Retry Pattern:**
```java
@Retryable(value = {ServiceUnavailableException.class}, 
           maxAttempts = 3, 
           backoff = @Backoff(delay = 1000))
public OrderResponse createOrder(OrderRequest request) {
    return orderClient.createOrder(request);
}
```

**3. Timeout Configuration:**
```yaml
# application.yml
resilience4j:
  circuitbreaker:
    instances:
      paymentService:
        failureRateThreshold: 50
        waitDurationInOpenState: 30s
        slidingWindowSize: 10
        minimumNumberOfCalls: 5
  timeout:
    instances:
      paymentService:
        timeoutDuration: 3s
```

**4. Health Monitoring:**
```java
@Component
public class ServiceHealthIndicator implements HealthIndicator {
    
    @Override
    public Health health() {
        // Check external service health
        if (isExternalServiceHealthy()) {
            return Health.up().build();
        }
        return Health.down().withDetail("error", "Service unavailable").build();
    }
}
```

---

## DevOps & Cloud

### Q21: How would you design a highly available 3-tier architecture across multiple AZs?

**Answer:** Design with redundancy and fault tolerance:

**Architecture Components:**

1. **Load Balancer Layer:**
   - Application Load Balancer (ALB) with multi-AZ deployment
   - Health checks and automatic failover
   - SSL termination

2. **Web/Application Layer:**
   - Auto Scaling Groups across multiple Availability Zones
   - EC2 instances with AMI-based deployments
   - Session state in external cache (Redis/ElastiCache)

3. **Database Layer:**
   - Multi-AZ RDS deployment
   - Read replicas for scaling reads
   - Automated backups and point-in-time recovery

**Example Infrastructure:**
```yaml
# CloudFormation template example
Resources:
  WebAppLoadBalancer:
    Type: AWS::ElasticLoadBalancingV2::LoadBalancer
    Properties:
      Scheme: internet-facing
      Type: application
      Subnets:
        - !Ref PublicSubnetAZ1
        - !Ref PublicSubnetAZ2
        - !Ref PublicSubnetAZ3
      
  WebAppASG:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      VPCZoneIdentifier:
        - !Ref PrivateSubnetAZ1
        - !Ref PrivateSubnetAZ2
        - !Ref PrivateSubnetAZ3
      MinSize: 3
      MaxSize: 10
      DesiredCapacity: 5
      
  Database:
    Type: AWS::RDS::DBInstance
    Properties:
      MultiAZ: true
      AllocatedStorage: 100
      DBInstanceClass: db.t3.medium
      Engine: postgres
      EngineVersion: "13.7"
```

### Q22: How would you implement cross-account access securely?

**Answer:** Use AWS IAM roles and resource-based policies:

**1. IAM Role for Cross-Account Access:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::ACCOUNT-B:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "unique-external-id"
        }
      }
    }
  ]
}
```

**2. Resource-Based Policy:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::ACCOUNT-B:role/CrossAccountRole"
      },
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource": "arn:aws:s3:::shared-bucket/*"
    }
  ]
}
```

**3. Assume Role Implementation:**
```java
@Service
public class CrossAccountService {
    
    @Value("${aws.cross-account.role-arn}")
    private String crossAccountRoleArn;
    
    public String assumeCrossAccountRole() {
        AWSSecurityTokenService sts = AWSSecurityTokenServiceClientBuilder.defaultClient();
        
        AssumeRoleRequest assumeRoleRequest = new AssumeRoleRequest()
            .withRoleArn(crossAccountRoleArn)
            .withRoleSessionName("cross-account-session")
            .withDurationSeconds(3600);
            
        AssumeRoleResult assumeRoleResult = sts.assumeRole(assumeRoleRequest);
        return assumeRoleResult.getCredentials().getAccessKeyId();
    }
}
```

---

## Coding Problems

### Q23: Write a program to convert a List into a Map in Java

**Answer:**
```java
import java.util.*;
import java.util.stream.Collectors;

public class ListToMapConverter {
    
    public static void main(String[] args) {
        List<User> users = Arrays.asList(
            new User(1, "John", "john@email.com"),
            new User(2, "Jane", "jane@email.com"),
            new User(3, "Bob", "bob@email.com")
        );
        
        // Method 1: Using Collectors.toMap()
        Map<Integer, User> userMap = users.stream()
            .collect(Collectors.toMap(
                User::getId,        // Key mapper
                user -> user,       // Value mapper
                (existing, replacement) -> existing // Merge function for duplicates
            ));
        
        // Method 2: Using groupingBy
        Map<String, List<User>> usersByEmailDomain = users.stream()
            .collect(Collectors.groupingBy(
                user -> user.getEmail().split("@")[1]
            ));
        
        // Method 3: Traditional approach
        Map<Integer, String> idToNameMap = new HashMap<>();
        for (User user : users) {
            idToNameMap.put(user.getId(), user.getName());
        }
        
        System.out.println("User Map: " + userMap);
        System.out.println("Email Domain Map: " + usersByEmailDomain);
        System.out.println("ID to Name Map: " + idToNameMap);
    }
}

class User {
    private int id;
    private String name;
    private String email;
    
    public User(int id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }
    
    // Getters
    public int getId() { return id; }
    public String getName() { return name; }
    public String getEmail() { return email; }
    
    @Override
    public String toString() {
        return "User{id=" + id + ", name='" + name + "'}";
    }
}
```

### Q24: Write code to find the maximum sum of a subarray for a given array (Kadane's Algorithm)

**Answer:**
```java
public class MaximumSubarraySum {
    
    public static void main(String[] args) {
        int[] arr = {-2, -3, 4, -1, -2, 1, 5, -3};
        
        int result = maxSubarraySum(arr);
        System.out.println("Maximum subarray sum: " + result);
        
        int[] resultWithIndices = maxSubarraySumWithIndices(arr);
        System.out.println("Maximum sum: " + resultWithIndices[0] + 
                          ", Start index: " + resultWithIndices[1] + 
                          ", End index: " + resultWithIndices[2]);
    }
    
    // Kadane's Algorithm - O(n) time, O(1) space
    public static int maxSubarraySum(int[] arr) {
        if (arr == null || arr.length == 0) {
            return 0;
        }
        
        int maxSoFar = arr[0];
        int maxEndingHere = arr[0];
        
        for (int i = 1; i < arr.length; i++) {
            maxEndingHere = Math.max(arr[i], maxEndingHere + arr[i]);
            maxSoFar = Math.max(maxSoFar, maxEndingHere);
        }
        
        return maxSoFar;
    }
    
    // Kadane's Algorithm with indices
    public static int[] maxSubarraySumWithIndices(int[] arr) {
        if (arr == null || arr.length == 0) {
            return new int[]{0, -1, -1};
        }
        
        int maxSoFar = arr[0];
        int maxEndingHere = arr[0];
        int start = 0;
        int end = 0;
        int tempStart = 0;
        
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > maxEndingHere + arr[i]) {
                maxEndingHere = arr[i];
                tempStart = i;
            } else {
                maxEndingHere = maxEndingHere + arr[i];
            }
            
            if (maxEndingHere > maxSoFar) {
                maxSoFar = maxEndingHere;
                start = tempStart;
                end = i;
            }
        }
        
        return new int[]{maxSoFar, start, end};
    }
    
    // Using Streams (Java 8+)
    public static int maxSubarraySumUsingStreams(int[] arr) {
        return Arrays.stream(arr)
            .boxed()
            .reduce(0, (max, current) -> {
                // This approach is less efficient for Kadane's algorithm
                // but demonstrates stream usage
                return Math.max(max, current);
            });
    }
}
```

### Q25: Print employee names with salary > 55,000 using Streams

**Answer:**
```java
import java.util.*;
import java.util.stream.Collectors;

public class EmployeeStreamExample {
    
    public static void main(String[] args) {
        List<Employee> employees = Arrays.asList(
            new Employee(1, "John Doe", 60000, "IT"),
            new Employee(2, "Jane Smith", 55000, "HR"),
            new Employee(3, "Bob Johnson", 70000, "Finance"),
            new Employee(4, "Alice Brown", 52000, "IT"),
            new Employee(5, "Charlie Wilson", 80000, "Sales")
        );
        
        // Method 1: Filter and collect names
        List<String> highEarners = employees.stream()
            .filter(emp -> emp.getSalary() > 55000)
            .map(Employee::getName)
            .collect(Collectors.toList());
        
        System.out.println("Employees with salary > 55,000: " + highEarners);
        
        // Method 2: Print directly
        System.out.println("\nDirect printing:");
        employees.stream()
            .filter(emp -> emp.getSalary() > 55000)
            .map(Employee::getName)
            .forEach(System.out::println);
        
        // Method 3: Group by department with salary filter
        Map<String, List<String>> highEarnersByDept = employees.stream()
            .filter(emp -> emp.getSalary() > 55000)
            .collect(Collectors.groupingBy(
                Employee::getDepartment,
                Collectors.mapping(Employee::getName, Collectors.toList())
            ));
        
        System.out.println("\nHigh earners by department: " + highEarnersByDept);
        
        // Method 4: Sort by salary descending
        List<String> sortedHighEarners = employees.stream()
            .filter(emp -> emp.getSalary() > 55000)
            .sorted(Comparator.comparing(Employee::getSalary).reversed())
            .map(Employee::getName)
            .collect(Collectors.toList());
        
        System.out.println("\nHigh earners sorted by salary: " + sortedHighEarners);
    }
}

class Employee {
    private int id;
    private String name;
    private double salary;
    private String department;
    
    public Employee(int id, String name, double salary, String department) {
        this.id = id;
        this.name = name;
        this.salary = salary;
        this.department = department;
    }
    
    // Getters
    public int getId() { return id; }
    public String getName() { return name; }
    public double getSalary() { return salary; }
    public String getDepartment() { return department; }
    
    @Override
    public String toString() {
        return String.format("%s ($%.2f)", name, salary);
    }
}
```

### Q26: Longest Common Subsequence (LCS) Problem

**Answer:**
```java
public class LongestCommonSubsequence {
    
    public static void main(String[] args) {
        String str1 = "ABC";
        String str2 = "ACD";
        
        LCSResult result = findLCS(str1, str2);
        System.out.println("String 1: " + str1);
        System.out.println("String 2: " + str2);
        System.out.println("LCS: " + result.lcs);
        System.out.println("Length: " + result.length);
    }
    
    // Dynamic Programming approach - O(m*n) time, O(m*n) space
    public static LCSResult findLCS(String str1, String str2) {
        int m = str1.length();
        int n = str2.length();
        
        // Create DP table
        int[][] dp = new int[m + 1][n + 1];
        
        // Fill the DP table
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (str1.charAt(i - 1) == str2.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                }
            }
        }
        
        // Backtrack to find the LCS string
        StringBuilder lcs = new StringBuilder();
        int i = m, j = n;
        
        while (i > 0 && j > 0) {
            if (str1.charAt(i - 1) == str2.charAt(j - 1)) {
                lcs.append(str1.charAt(i - 1));
                i--;
                j--;
            } else if (dp[i - 1][j] > dp[i][j - 1]) {
                i--;
            } else {
                j--;
            }
        }
        
        return new LCSResult(lcs.reverse().toString(), dp[m][n]);
    }
    
    // Recursive approach with memoization
    public static String findLCSRecursive(String str1, String str2, int m, int n, 
                                         Map<String, String> memo) {
        String key = m + "," + n;
        
        if (memo.containsKey(key)) {
            return memo.get(key);
        }
        
        if (m == 0 || n == 0) {
            return "";
        }
        
        if (str1.charAt(m - 1) == str2.charAt(n - 1)) {
            String result = findLCSRecursive(str1, str2, m - 1, n - 1, memo) + 
                           str1.charAt(m - 1);
            memo.put(key, result);
            return result;
        } else {
            String result1 = findLCSRecursive(str1, str2, m - 1, n, memo);
            String result2 = findLCSRecursive(str1, str2, m, n - 1, memo);
            String finalResult = result1.length() > result2.length() ? result1 : result2;
            memo.put(key, finalResult);
            return finalResult;
        }
    }
    
    static class LCSResult {
        String lcs;
        int length;
        
        public LCSResult(String lcs, int length) {
            this.lcs = lcs;
            this.length = length;
        }
    }
}
```

---

## Level-Based Interview Questions

### Level 1 Questions

#### Q28: Explain about your project and tech stack you are using

**Answer:** When answering this question, structure your response as follows:

**Project Overview:**
"I'm currently working on [Project Name], which is a [brief description of what the system does]. It's a [web/mobile/enterprise] application that serves [target users] by solving [specific problem]."

**Tech Stack:**
- **Backend**: Java 11/17, Spring Boot 2.7/3.x, Spring Security, Spring Data JPA
- **Database**: PostgreSQL/MySQL with Redis for caching
- **Message Queue**: Apache Kafka for event-driven architecture
- **Cloud**: AWS (EC2, S3, RDS, Lambda) or Azure
- **CI/CD**: Jenkins/GitHub Actions with Docker and Kubernetes
- **Monitoring**: Prometheus, Grafana, ELK stack

**Key Features:**
- User authentication and authorization
- Real-time data processing
- RESTful APIs with versioning
- Microservices architecture with service discovery

#### Q29: What are your roles and responsibilities in your team?

**Answer:** Structure your response to show ownership and impact:

**Primary Responsibilities:**
- **Development**: Design and implement RESTful APIs using Spring Boot
- **Code Quality**: Write unit tests with JUnit 5 and integration tests
- **Database Design**: Optimize SQL queries and design database schemas
- **Code Reviews**: Review team members' code for best practices
- **Technical Documentation**: Create and maintain API documentation

**Collaboration:**
- Work with product managers to understand requirements
- Collaborate with frontend developers for API integration
- Participate in agile ceremonies (daily standups, sprint planning)
- Mentor junior developers

**Achievements:**
- Reduced API response time by 40% through caching strategies
- Implemented automated testing reducing production bugs by 60%
- Led migration from monolith to microservices architecture

#### Q30: Method overriding example, what should be the access specifier of the overriding method?

**Answer:** In method overriding, the access specifier of the overriding method can be the same or more accessible than the parent class method, but never more restrictive.

**Access Modifier Rules:**
- **private** → Cannot be overridden (not visible)
- **default (package-private)** → Can be overridden as default, protected, or public
- **protected** → Can be overridden as protected or public
- **public** → Must remain public

**Example:**
```java
class Parent {
    protected void display() {
        System.out.println("Parent display");
    }
    
    public void show() {
        System.out.println("Parent show");
    }
}

class Child extends Parent {
    // Valid: same or more accessible
    @Override
    public void display() { // Changed from protected to public
        System.out.println("Child display");
    }
    
    // Valid: same accessibility
    @Override
    public void show() {
        System.out.println("Child show");
    }
    
    // Invalid: Cannot reduce accessibility
    // @Override
    // protected void show() { } // Compile error
}
```

#### Q31: What are access specifiers in java? Name a few you are familiar with

**Answer:** Java has four access specifiers that control the visibility of classes, methods, and fields:

**1. private:**
- Accessible only within the same class
- Most restrictive access level
```java
class MyClass {
    private int privateField = 10;
    private void privateMethod() { }
}
```

**2. default (package-private):**
- Accessible within the same package only
- No explicit keyword needed
```java
class MyClass {
    int defaultField = 20; // Package-private
    void defaultMethod() { }
}
```

**3. protected:**
- Accessible within the same package and subclasses
- Used for inheritance scenarios
```java
class Parent {
    protected int protectedField = 30;
    protected void protectedMethod() { }
}
```

**4. public:**
- Accessible from anywhere
- Most permissive access level
```java
public class MyClass {
    public int publicField = 40;
    public void publicMethod() { }
}
```

**Visibility Table:**
| Modifier | Same Class | Same Package | Subclass | World |
|----------|-------------|---------------|-----------|-------|
| private  | ✓           | ✗             | ✗         | ✗     |
| default  | ✓           | ✓             | ✗         | ✗     |
| protected| ✓           | ✓             | ✓         | ✗     |
| public   | ✓           | ✓             | ✓         | ✓     |

#### Q32: Example of what happens when you insert different objects with the same data into HashSet?

**Answer:** When you insert different objects with the same data into HashSet, the behavior depends on whether the objects have proper `equals()` and `hashCode()` implementations.

**Scenario 1: Without overriding equals() and hashCode()**
```java
class Person {
    String name;
    int age;
    
    Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
}

public class Main {
    public static void main(String[] args) {
        Set<Person> set = new HashSet<>();
        Person p1 = new Person("John", 25);
        Person p2 = new Person("John", 25);
        
        set.add(p1);
        set.add(p2);
        
        System.out.println(set.size()); // Output: 2
        // Both objects are added because they have different memory addresses
    }
}
```

**Scenario 2: With proper equals() and hashCode()**
```java
class Person {
    String name;
    int age;
    
    Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
    
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Person person = (Person) o;
        return age == person.age && Objects.equals(name, person.name);
    }
    
    @Override
    public int hashCode() {
        return Objects.hash(name, age);
    }
}

public class Main {
    public static void main(String[] args) {
        Set<Person> set = new HashSet<>();
        Person p1 = new Person("John", 25);
        Person p2 = new Person("John", 25);
        
        set.add(p1);
        set.add(p2);
        
        System.out.println(set.size()); // Output: 1
        // Only one object is added because they are considered equal
    }
}
```

**Key Points:**
- HashSet uses `hashCode()` to determine the bucket
- If hash codes are equal, it uses `equals()` to check for equality
- Always override both methods together
- Follow the contract: equal objects must have equal hash codes

#### Q33: Write a query to fetch the student who has scored top third highest marks

**Answer:** There are multiple approaches to find the third highest marks:

**Approach 1: Using LIMIT and OFFSET (MySQL/PostgreSQL)**
```sql
SELECT * FROM students 
ORDER BY marks DESC 
LIMIT 1 OFFSET 2;
```

**Approach 2: Using Subquery (Database Independent)**
```sql
SELECT * FROM students 
WHERE marks = (
    SELECT DISTINCT marks 
    FROM students 
    ORDER BY marks DESC 
    LIMIT 1 OFFSET 2
);
```

**Approach 3: Using Window Functions (Modern SQL)**
```sql
WITH ranked_students AS (
    SELECT *, 
           DENSE_RANK() OVER (ORDER BY marks DESC) as rank_num
    FROM students
)
SELECT * FROM ranked_students 
WHERE rank_num = 3;
```

**Approach 4: Handle Ties (Multiple students with same marks)**
```sql
SELECT * FROM students 
WHERE marks = (
    SELECT DISTINCT marks 
    FROM students 
    ORDER BY marks DESC 
    LIMIT 1 OFFSET 2
)
ORDER BY student_name;
```

**Sample Data and Expected Output:**
```sql
-- Sample table
CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    marks INT
);

INSERT INTO students VALUES 
(1, 'Alice', 95),
(2, 'Bob', 87),
(3, 'Charlie', 92),
(4, 'David', 87),
(5, 'Eve', 78);

-- Query result would show students with 87 marks (Bob and David)
```

#### Q34: How to generalize a common response object for all end points in the project?

**Answer:** Create a standardized response wrapper using Spring Boot:

**1. Create Common Response Class:**
```java
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ApiResponse<T> {
    private boolean success;
    private String message;
    private T data;
    private String timestamp;
    private String path;
    
    public static <T> ApiResponse<T> success(T data) {
        return ApiResponse.<T>builder()
            .success(true)
            .data(data)
            .timestamp(LocalDateTime.now().toString())
            .build();
    }
    
    public static <T> ApiResponse<T> success(T data, String message) {
        return ApiResponse.<T>builder()
            .success(true)
            .message(message)
            .data(data)
            .timestamp(LocalDateTime.now().toString())
            .build();
    }
    
    public static <T> ApiResponse<T> error(String message) {
        return ApiResponse.<T>builder()
            .success(false)
            .message(message)
            .timestamp(LocalDateTime.now().toString())
            .build();
    }
}
```

**2. Create Response Controller Advice:**
```java
@RestControllerAdvice
public class ResponseAdvice implements ResponseBodyAdvice<Object> {
    
    @Override
    public boolean supports(MethodParameter returnType, 
                           Class<? extends HttpMessageConverter<?>> converterType) {
        // Don't wrap ApiResponse objects
        return !returnType.getParameterType().equals(ApiResponse.class);
    }
    
    @Override
    public Object beforeBodyWrite(Object body, MethodParameter returnType,
                                 MediaType selectedContentType,
                                 Class<? extends HttpMessageConverter<?>> selectedConverterType,
                                 ServerHttpRequest request, ServerHttpResponse response) {
        
        // Handle null responses
        if (body == null) {
            return ApiResponse.success(null);
        }
        
        // Handle string responses separately
        if (body instanceof String) {
            return ApiResponse.success(body);
        }
        
        return ApiResponse.success(body);
    }
}
```

**3. Usage in Controllers:**
```java
@RestController
@RequestMapping("/api/users")
public class UserController {
    
    @GetMapping("/{id}")
    public User getUser(@PathVariable Long id) {
        return userService.findById(id);
        // Automatically wrapped in ApiResponse
    }
    
    @PostMapping
    public ApiResponse<User> createUser(@RequestBody UserDto dto) {
        User user = userService.create(dto);
        return ApiResponse.success(user, "User created successfully");
        // Custom message
    }
}
```

**4. Error Response Integration:**
```java
@RestControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ApiResponse<Void>> handleNotFound(
            ResourceNotFoundException ex) {
        ApiResponse<Void> response = ApiResponse.error(ex.getMessage());
        return ResponseEntity.status(HttpStatus.NOT_FOUND).body(response);
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse<Void>> handleGeneric(Exception ex) {
        ApiResponse<Void> response = ApiResponse.error("Internal server error");
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(response);
    }
}
```

#### Q35: How can you create spring beans of two different implementations of a single interface without any issues?

**Answer:** Use @Qualifier and @Primary annotations to resolve ambiguity:

**1. Define Interface and Implementations:**
```java
public interface NotificationService {
    void sendNotification(String message);
}

@Service
public class EmailNotificationService implements NotificationService {
    @Override
    public void sendNotification(String message) {
        System.out.println("Email: " + message);
    }
}

@Service
public class SmsNotificationService implements NotificationService {
    @Override
    public void sendNotification(String message) {
        System.out.println("SMS: " + message);
    }
}
```

**2. Using @Qualifier:**
```java
@Service
public class NotificationManager {
    
    private final NotificationService emailService;
    private final NotificationService smsService;
    
    public NotificationManager(
            @Qualifier("emailNotificationService") NotificationService emailService,
            @Qualifier("smsNotificationService") NotificationService smsService) {
        this.emailService = emailService;
        this.smsService = smsService;
    }
    
    public void sendEmail(String message) {
        emailService.sendNotification(message);
    }
    
    public void sendSms(String message) {
        smsService.sendNotification(message);
    }
}
```

**3. Using Custom Qualifiers:**
```java
@Service
@Qualifier("email")
public class EmailNotificationService implements NotificationService {
    // Implementation
}

@Service
@Qualifier("sms")
public class SmsNotificationService implements NotificationService {
    // Implementation
}

@Service
public class NotificationManager {
    
    @Autowired
    @Qualifier("email")
    private NotificationService emailService;
    
    @Autowired
    @Qualifier("sms")
    private NotificationService smsService;
}
```

**4. Using @Primary:**
```java
@Service
@Primary
public class EmailNotificationService implements NotificationService {
    // This will be the default implementation
}

@Service
public class SmsNotificationService implements NotificationService {
    // Secondary implementation
}

@Service
public class NotificationManager {
    
    @Autowired
    private NotificationService defaultService; // EmailNotificationService
    
    @Autowired
    @Qualifier("smsNotificationService")
    private NotificationService smsService;
}
```

**5. Using @Bean Configuration:**
```java
@Configuration
public class NotificationConfig {
    
    @Bean
    @Primary
    public NotificationService emailNotificationService() {
        return new EmailNotificationService();
    }
    
    @Bean
    public NotificationService smsNotificationService() {
        return new SmsNotificationService();
    }
}
```

---

### Level 2 Questions

#### Q36: Which version of java are you using?

**Answer:** "I'm currently working with Java 17 in our production environment, which is an LTS (Long-Term Support) version. We chose Java 17 for its stability and long-term support until 2029. 

**Key Features of Java 17 we're using:**
- **Sealed Classes**: For better domain modeling
- **Pattern Matching for instanceof**: Cleaner type checking
- **Text Blocks**: For multi-line strings in SQL queries
- **Records**: For immutable data carriers
- **Switch Expressions**: More concise conditional logic

**Migration Benefits:**
- Improved performance compared to Java 8
- Better garbage collection with ZGC
- Enhanced security features
- Modern language features for cleaner code

**Development Setup:**
- Local development: Java 17
- CI/CD: OpenJDK 17
- Production: Amazon Corretto 17"

#### Q37: What are the functional interfaces and how do you use them?

**Answer:** Functional interfaces are interfaces with exactly one abstract method, enabling lambda expressions and method references.

**Built-in Functional Interfaces:**

**1. Predicate<T>**: Boolean-valued function
```java
Predicate<String> isNotEmpty = str -> !str.isEmpty();
List<String> nonEmpty = list.stream()
    .filter(isNotEmpty)
    .collect(Collectors.toList());
```

**2. Function<T, R>**: Transform input to output
```java
Function<String, Integer> stringLength = String::length;
List<Integer> lengths = names.stream()
    .map(stringLength)
    .collect(Collectors.toList());
```

**3. Consumer<T>**: Accept input, return void
```java
Consumer<String> printer = System.out::println;
names.forEach(printer);
```

**4. Supplier<T>**: Provide values
```java
Supplier<Double> randomValue = Math::random;
Double value = randomValue.get();
```

**5. UnaryOperator<T>**: Same input and output type
```java
UnaryOperator<String> toUpperCase = String::toUpperCase;
List<String> upper = names.stream()
    .map(toUpperCase)
    .collect(Collectors.toList());
```

**Custom Functional Interface:**
```java
@FunctionalInterface
interface Validator<T> {
    boolean validate(T t);
    
    // Default method allowed
    default Validator<T> and(Validator<T> other) {
        return t -> this.validate(t) && other.validate(t);
    }
}

// Usage
Validator<String> notNull = str -> str != null;
Validator<String> notEmpty = str -> !str.isEmpty();
Validator<String> emailValidator = str -> str.contains("@");

Validator<String> completeValidation = notNull.and(notEmpty).and(emailValidator);
boolean isValid = completeValidation.validate("test@example.com");
```

#### Q38: How do you read huge amounts of data from a file with limited CPU and memory?

**Answer:** Use streaming and buffered approaches to handle large files efficiently:

**1. BufferedReader for Line-by-Line Processing:**
```java
public void processLargeFile(String filePath) throws IOException {
    try (BufferedReader reader = Files.newBufferedReader(Paths.get(filePath))) {
        String line;
        while ((line = reader.readLine()) != null) {
            // Process one line at a time
            processLine(line);
        }
    }
}
```

**2. Stream API with Files.lines():**
```java
public void processWithStream(String filePath) throws IOException {
    try (Stream<String> lines = Files.lines(Paths.get(filePath))) {
        lines.parallel()
            .filter(line -> !line.trim().isEmpty())
            .map(this::parseLine)
            .forEach(this::processRecord);
    }
}
```

**3. Memory-Efficient CSV Processing:**
```java
public void processLargeCSV(String filePath) throws IOException {
    try (CSVReader reader = new CSVReader(new FileReader(filePath))) {
        String[] nextLine;
        while ((nextLine = reader.readNext()) != null) {
            // Process each row
            processCSVRow(nextLine);
        }
    }
}
```

**4. Batch Processing for Database Operations:**
```java
public void batchInsertFromFile(String filePath, int batchSize) throws IOException {
    List<Record> batch = new ArrayList<>(batchSize);
    
    try (BufferedReader reader = Files.newBufferedReader(Paths.get(filePath))) {
        String line;
        while ((line = reader.readLine()) != null) {
            Record record = parseLine(line);
            batch.add(record);
            
            if (batch.size() >= batchSize) {
                repository.saveAll(batch);
                batch.clear();
            }
        }
        
        // Save remaining records
        if (!batch.isEmpty()) {
            repository.saveAll(batch);
        }
    }
}
```

**5. Memory-Mapped Files for Very Large Files:**
```java
public void processWithMemoryMapping(String filePath) throws IOException {
    try (RandomAccessFile file = new RandomAccessFile(filePath, "r");
         FileChannel channel = file.getChannel()) {
        
        MappedByteBuffer buffer = channel.map(
            FileChannel.MapMode.READ_ONLY, 0, channel.size()
        );
        
        // Process in chunks
        int chunkSize = 8192; // 8KB chunks
        for (int position = 0; position < buffer.limit(); position += chunkSize) {
            int end = Math.min(position + chunkSize, buffer.limit());
            ByteBuffer chunk = buffer.slice(position, end - position);
            processChunk(chunk);
        }
    }
}
```

**Best Practices:**
- Use try-with-resources for automatic resource management
- Process data in streams rather than loading entire file
- Implement batch processing for database operations
- Consider parallel streams for CPU-intensive processing
- Monitor memory usage and adjust chunk sizes accordingly

#### Q39: How do you assign a task to a thread? You can use thread frameworks

**Answer:** There are several ways to assign tasks to threads in Java:

**1. Basic Thread Creation:**
```java
public class BasicThreadExample {
    public static void main(String[] args) {
        Thread thread = new Thread(() -> {
            System.out.println("Task running in: " + Thread.currentThread().getName());
        });
        thread.start();
    }
}
```

**2. Using ExecutorService (Recommended):**
```java
public class ExecutorServiceExample {
    private final ExecutorService executor = Executors.newFixedThreadPool(10);
    
    public void submitTask() {
        Future<String> future = executor.submit(() -> {
            // Task implementation
            Thread.sleep(1000);
            return "Task completed";
        });
        
        // Get result later
        try {
            String result = future.get();
            System.out.println(result);
        } catch (InterruptedException | ExecutionException e) {
            Thread.currentThread().interrupt();
        }
    }
    
    public void shutdown() {
        executor.shutdown();
    }
}
```

**3. Using CompletableFuture (Modern Approach):**
```java
public class CompletableFutureExample {
    
    public void asyncTask() {
        CompletableFuture.supplyAsync(() -> {
            // Long-running task
            return performExpensiveOperation();
        })
        .thenApply(result -> result.toUpperCase())
        .thenAccept(result -> System.out.println("Result: " + result))
        .exceptionally(throwable -> {
            System.err.println("Error: " + throwable.getMessage());
            return null;
        });
    }
    
    public void parallelTasks() {
        CompletableFuture<String> task1 = CompletableFuture.supplyAsync(() -> {
            return performTask1();
        });
        
        CompletableFuture<String> task2 = CompletableFuture.supplyAsync(() -> {
            return performTask2();
        });
        
        CompletableFuture<Void> allTasks = CompletableFuture.allOf(task1, task2);
        allTasks.thenRun(() -> {
            System.out.println("All tasks completed");
        });
    }
}
```

**4. Using Spring's @Async:**
```java
@Configuration
@EnableAsync
public class AsyncConfig {
    
    @Bean
    public TaskExecutor taskExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(5);
        executor.setMaxPoolSize(10);
        executor.setQueueCapacity(100);
        executor.setThreadNamePrefix("Async-");
        executor.initialize();
        return executor;
    }
}

@Service
public class AsyncService {
    
    @Async
    public CompletableFuture<String> asyncMethod() {
        // Async task implementation
        return CompletableFuture.completedFuture("Async result");
    }
}
```

**5. Custom Thread Pool with ThreadPoolExecutor:**
```java
public class CustomThreadPoolExample {
    private final ThreadPoolExecutor executor;
    
    public CustomThreadPoolExample() {
        this.executor = new ThreadPoolExecutor(
            5, // core pool size
            10, // max pool size
            60, // keep alive time
            TimeUnit.SECONDS,
            new LinkedBlockingQueue<>(100),
            new ThreadFactoryBuilder().setNameFormat("worker-%d").build(),
            new ThreadPoolExecutor.CallerRunsPolicy() // rejection policy
        );
    }
    
    public void submitTask(Runnable task) {
        executor.submit(task);
    }
    
    public void shutdown() {
        executor.shutdown();
        try {
            if (!executor.awaitTermination(60, TimeUnit.SECONDS)) {
                executor.shutdownNow();
            }
        } catch (InterruptedException e) {
            executor.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }
}
```

**Best Practices:**
- Prefer ExecutorService over manual thread creation
- Use CompletableFuture for modern async programming
- Configure appropriate thread pool sizes
- Handle exceptions properly in async tasks
- Always shutdown thread pools gracefully
- Consider using Spring's @Async for Spring applications

---

## Additional Advanced Topics

### Q27: What is Kafka and where is it used?

**Answer:** Apache Kafka is a distributed streaming platform designed for high-throughput, fault-tolerant, and scalable real-time data streaming.

**Key Features:**
- **Publish-Subscribe Model**: Producers write to topics, consumers read from topics
- **Distributed Architecture**: Multiple brokers for scalability and fault tolerance
- **Durability**: Messages persisted to disk with configurable retention
- **High Performance**: Handles millions of messages per second

**Use Cases:**
1. **Event Streaming**: Real-time event processing
2. **Log Aggregation**: Centralized logging from multiple services
3. **Message Queue**: Decoupling microservices
4. **Data Pipeline**: ETL processes and data integration

**Spring Boot Integration:**
```java
@Configuration
public class KafkaConfig {
    
    @Bean
    public ProducerFactory<String, String> producerFactory() {
        Map<String, Object> configProps = new HashMap<>();
        configProps.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
        configProps.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        configProps.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        return new DefaultKafkaProducerFactory<>(configProps);
    }
    
    @Bean
    public ConsumerFactory<String, String> consumerFactory() {
        Map<String, Object> configProps = new HashMap<>();
        configProps.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
        configProps.put(ConsumerConfig.GROUP_ID_CONFIG, "my-group");
        configProps.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        configProps.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        return new DefaultKafkaConsumerFactory<>(configProps);
    }
}

@Service
public class KafkaProducerService {
    
    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;
    
    public void sendMessage(String topic, String message) {
        kafkaTemplate.send(topic, message);
    }
}

@Service
public class KafkaConsumerService {
    
    @KafkaListener(topics = "user-events", groupId = "user-group")
    public void consumeMessage(String message) {
        System.out.println("Received message: " + message);
        // Process the message
    }
}
```

---

## Interview Tips for Virtusa

### Technical Round Preparation:
1. **Focus on Practical Scenarios**: Virtusa emphasizes real-world problem-solving
2. **Spring Boot Expertise**: Strong knowledge of Spring ecosystem is crucial
3. **Java 8+ Features**: Be comfortable with streams, lambdas, and optional
4. **Database Skills**: SQL optimization and transaction management
5. **System Design**: Understanding of microservices and distributed systems

### Scenario-Based Questions:
- Always think about production scenarios
- Consider edge cases and error handling
- Discuss performance implications
- Mention security considerations
- Talk about monitoring and observability

### Coding Round Tips:
- Write clean, readable code
- Handle edge cases
- Use appropriate data structures
- Consider time and space complexity
- Test your solutions with sample inputs

---

## Additional Missing Questions & Answers

### Q25: What is the purpose of annotations in Spring Boot / Java?

**Answer:** Annotations provide metadata about code and enable declarative programming. In Java and Spring Boot, they serve several purposes:

**In Java:**
- **Compile-time checking**: `@Override`, `@Deprecated`
- **Runtime processing**: `@Entity`, `@RestController`
- **Documentation**: `@param`, `@return`

**In Spring Boot:**
- **Dependency Injection**: `@Autowired`, `@Component`, `@Service`, `@Repository`
- **Configuration**: `@Configuration`, `@Bean`, `@Value`
- **Web Mapping**: `@RestController`, `@GetMapping`, `@PostMapping`
- **Validation**: `@Valid`, `@NotNull`, `@Size`
- **Transaction Management**: `@Transactional`

**Example:**
```java
@RestController
public class UserController {
    
    @Autowired
    private UserService userService;
    
    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable @Min(1) Long id) {
        User user = userService.findById(id);
        return ResponseEntity.ok(user);
    }
}
```

### Q26: What is blue/green deployment?

**Answer:** Blue/Green deployment is a deployment strategy that minimizes downtime by running two identical production environments:

**Key Concepts:**
- **Blue Environment**: Current production version
- **Green Environment**: New version with updates
- **Router**: Switches traffic between environments

**Deployment Process:**
1. Deploy new version to Green environment
2. Run smoke tests and integration tests
3. Switch traffic from Blue to Green
4. Monitor for issues
5. Keep Blue as rollback backup

**AWS Implementation:**
```yaml
# Using AWS CodeDeploy with Blue/Green
Resources:
  Application:
    Type: AWS::CodeDeploy::Application
    
  DeploymentGroup:
    Type: AWS::CodeDeploy::DeploymentGroup
    Properties:
      ApplicationName: !Ref Application
      DeploymentConfigName: CodeDeployDefault.AllAtOnce
      DeploymentStyle:
        DeploymentType: BLUE_GREEN
        DeploymentOption: WITH_TRAFFIC_CONTROL
      AutoRollbackConfiguration:
        Enabled: true
        Events:
          - DEPLOYMENT_FAILURE
          - DEPLOYMENT_STOP_ON_ALARM
          - DEPLOYMENT_STOP_ON_REQUEST
```

### Q27: What is SSL?

**Answer:** SSL (Secure Sockets Layer) is a cryptographic protocol that provides secure communication over computer networks. SSL has been succeeded by TLS (Transport Layer Security), but the term SSL is still commonly used.

**Key Features:**
- **Encryption**: Data is encrypted during transmission
- **Authentication**: Server identity verification
- **Integrity**: Data integrity verification

**SSL/TLS Handshake Process:**
1. Client sends Hello message
2. Server responds with certificate
3. Client verifies certificate
4. Key exchange for symmetric encryption
5. Secure communication begins

**Spring Boot SSL Configuration:**
```yaml
server:
  ssl:
    enabled: true
    key-store: classpath:keystore.p12
    key-store-password: password
    key-store-type: PKCS12
    key-alias: tomcat
    trust-store: classpath:truststore.p12
    trust-store-password: password
```

### Q28: HTTP vs HTTPS?

**Answer:**

| Aspect | HTTP | HTTPS |
|--------|------|-------|
| **Security** | Unencrypted, plain text | Encrypted with SSL/TLS |
| **Port** | Port 80 | Port 443 |
| **Certificate** | Not required | SSL certificate required |
| **Performance** | Faster | Slightly slower due to encryption |
| **SEO** | Lower ranking | Better SEO ranking |
| **Data Integrity** | Vulnerable to tampering | Protected against tampering |

**Key Differences:**
- **HTTP**: Data travels in plain text, vulnerable to eavesdropping
- **HTTPS**: Data encrypted, secure authentication and integrity

**Spring Boot HTTPS Redirect:**
```java
@Configuration
public class HttpsConfig {
    
    @Bean
    public ServletWebServerFactory servletContainer() {
        TomcatServletWebServerFactory tomcat = new TomcatServletWebServerFactory() {
            @Override
            protected void postProcessContext(Context context) {
                SecurityConstraint securityConstraint = new SecurityConstraint();
                securityConstraint.setUserConstraint("CONFIDENTIAL");
                SecurityCollection collection = new SecurityCollection();
                collection.addPattern("/*");
                securityConstraint.addCollection(collection);
                context.addConstraint(securityConstraint);
            }
        };
        tomcat.addAdditionalTomcatConnectors(redirectConnector());
        return tomcat;
    }
    
    private Connector redirectConnector() {
        Connector connector = new Connector("org.apache.coyote.http11.Http11NioProtocol");
        connector.setScheme("http");
        connector.setPort(8080);
        connector.setSecure(false);
        connector.setRedirectPort(8443);
        return connector;
    }
}
```

### Q29: How did you implement circuit breaking?

**Answer:** Circuit breaking prevents cascading failures by stopping requests to failing services:

**Implementation using Resilience4j:**
```java
@Service
public class PaymentService {
    
    @Autowired
    private PaymentClient paymentClient;
    
    @CircuitBreaker(name = "paymentService", fallbackMethod = "fallbackPayment")
    public PaymentResponse processPayment(PaymentRequest request) {
        return paymentClient.processPayment(request);
    }
    
    public PaymentResponse fallbackPayment(PaymentRequest request, Exception ex) {
        log.error("Payment service failed, using fallback: {}", ex.getMessage());
        return PaymentResponse.builder()
            .status("PENDING")
            .message("Payment service temporarily unavailable")
            .retryAfter(Duration.ofMinutes(5))
            .build();
    }
}
```

**Configuration:**
```yaml
resilience4j:
  circuitbreaker:
    instances:
      paymentService:
        failureRateThreshold: 50
        waitDurationInOpenState: 30s
        slidingWindowSize: 10
        minimumNumberOfCalls: 5
        permittedNumberOfCallsInHalfOpenState: 3
        automaticTransitionFromOpenToHalfOpenEnabled: true
```

### Q30: When did you use the retry mechanism?

**Answer:** Retry mechanism is used for transient failures:

**Implementation:**
```java
@Service
public class EmailService {
    
    @Retryable(
        value = {SMTPException.class, ConnectException.class},
        maxAttempts = 3,
        backoff = @Backoff(delay = 1000, multiplier = 2)
    )
    public void sendEmail(EmailRequest request) {
        emailClient.send(request);
    }
    
    @Recover
    public void recoverEmail(EmailRequest request, Exception ex) {
        log.error("Failed to send email after 3 attempts: {}", ex.getMessage());
        // Add to dead letter queue for manual retry
        deadLetterQueue.add(request);
    }
}
```

### Q31: What was the purpose of using Kafka in your project?

**Answer:** Kafka was used for several key purposes:

**1. Event-Driven Architecture:**
```java
// User registration event
@KafkaListener(topics = "user-events", groupId = "user-service")
public void handleUserRegistration(UserEvent event) {
    // Process user registration
    userService.createUser(event);
    
    // Publish to other services
    kafkaTemplate.send("notification-events", 
        new NotificationEvent("welcome", event.getUserId()));
}
```

**2. Asynchronous Processing:**
- Order processing without blocking user requests
- Data synchronization between microservices
- Real-time analytics and monitoring

**3. Decoupling Services:**
```java
@Service
public class OrderService {
    
    public Order createOrder(OrderRequest request) {
        Order order = orderRepository.save(request);
        
        // Async processing
        OrderEvent event = new OrderEvent(order.getId(), "CREATED");
        kafkaTemplate.send("order-events", event);
        
        return order;
    }
}
```

### Q32: What was the reason for choosing SingleStore database in your project?

**Answer:** SingleStore was chosen for its high-performance capabilities:

**Key Benefits:**
- **Real-time Analytics**: Fast aggregations on large datasets
- **Hybrid Architecture**: Combines OLTP and OLAP workloads
- **Memory-Optimized**: In-memory processing for speed
- **Distributed**: Horizontal scaling capabilities

**Use Cases:**
```java
// Fast aggregations for real-time dashboards
@Repository
public interface AnalyticsRepository extends JpaRepository<Analytics, Long> {
    
    @Query("SELECT category, COUNT(*), AVG(price) FROM Product " +
           "WHERE created_at > :date GROUP BY category")
    List<Object[]> getRealTimeAnalytics(@Param("date") LocalDateTime date);
}
```

### Q33: How will you fix the issues when multiple instances are failing to serve the incoming requests?

**Answer:** Multi-step approach to handle service failures:

**1. Auto-scaling:**
```yaml
# Kubernetes HPA
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: app-deployment
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

**2. Load Balancing:**
```java
@Configuration
public class LoadBalancerConfig {
    
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
```

**3. Circuit Breaker:**
```java
@CircuitBreaker(name = "service", fallbackMethod = "fallback")
public Response callService(Request request) {
    return serviceClient.call(request);
}
```

**4. Health Monitoring:**
```java
@Component
public class HealthMonitor {
    
    @Scheduled(fixedRate = 30000)
    public void checkServiceHealth() {
        List<ServiceInstance> instances = discoveryClient.getInstances("service");
        instances.forEach(this::checkInstance);
    }
}
```

---

## Advanced Scenario-Based Questions & Answers

### Q34: Why does an endpoint work with `@RequestMapping(method=POST)` but fail with `@PostMapping`?

**Answer:** This issue typically occurs due to missing or conflicting import statements:

**Root Cause:**
```java
// Wrong import - causes @PostMapping to not work
import org.springframework.web.bind.annotation.RequestMapping;

// Correct import for @PostMapping
import org.springframework.web.bind.annotation.PostMapping;
```

**Solution:**
```java
@RestController
public class UserController {
    
    // Both work when proper imports are used
    @RequestMapping(method = RequestMethod.POST, value = "/users")
    public ResponseEntity<User> createUserWithRequestMapping(@RequestBody User user) {
        return ResponseEntity.ok(userService.save(user));
    }
    
    @PostMapping("/users")
    public ResponseEntity<User> createUserWithPostMapping(@RequestBody User user) {
        return ResponseEntity.ok(userService.save(user));
    }
}
```

**Key Points:**
- Ensure `@PostMapping` is imported from `org.springframework.web.bind.annotation.PostMapping`
- `@RequestMapping` is more generic, `@PostMapping` is a specialized annotation
- Both annotations function identically when properly imported

### Q35: Why doesn't `RestTemplate` retry on a socket timeout even after configuring retries?

**Answer:** Socket timeouts are not retryable by default in Spring Retry:

**Root Cause:**
```java
// Socket timeout is not considered a transient exception
@Retryable(value = {ConnectException.class}, maxAttempts = 3)
public ResponseEntity<String> callExternalService() {
    // This won't retry on SocketTimeoutException
    return restTemplate.getForEntity(url, String.class);
}
```

**Solution:**
```java
@Configuration
public class RetryConfig {
    
    @Bean
    public RestTemplate restTemplate() {
        RestTemplate restTemplate = new RestTemplate();
        
        // Configure retry with custom retry policy
        HttpComponentsClientHttpRequestFactory factory = 
            new HttpComponentsClientHttpRequestFactory();
        factory.setConnectTimeout(5000);
        factory.setReadTimeout(5000);
        
        restTemplate.setRequestFactory(factory);
        return restTemplate;
    }
}

@Service
public class ExternalService {
    
    @Retryable(
        value = {
            ConnectException.class, 
            SocketTimeoutException.class,
            ResourceAccessException.class
        },
        maxAttempts = 3,
        backoff = @Backoff(delay = 1000, multiplier = 2)
    )
    public ResponseEntity<String> callExternalService() {
        return restTemplate.getForEntity(url, String.class);
    }
}
```

**Alternative - Using RetryTemplate:**
```java
@Service
public class ExternalService {
    
    @Autowired
    private RetryTemplate retryTemplate;
    
    public ResponseEntity<String> callExternalService() {
        return retryTemplate.execute(context -> {
            return restTemplate.getForEntity(url, String.class);
        });
    }
    
    @Bean
    public RetryTemplate retryTemplate() {
        RetryTemplate retryTemplate = new RetryTemplate();
        
        FixedBackOffPolicy backOffPolicy = new FixedBackOffPolicy();
        backOffPolicy.setBackOffPeriod(1000);
        
        SimpleRetryPolicy retryPolicy = new SimpleRetryPolicy();
        retryPolicy.setMaxAttempts(3);
        
        retryTemplate.setBackOffPolicy(backOffPolicy);
        retryTemplate.setRetryPolicy(retryPolicy);
        
        return retryTemplate;
    }
}
```

### Q36: Why does `Actuator/health` sometimes report a `DOWN` status for a database or service when everything appears fine?

**Answer:** Health check indicators may have timeout or configuration issues:

**Common Causes:**
1. **Connection Pool Exhaustion**
2. **Network Latency**
3. **Query Timeout**
4. **Authentication Issues**

**Solution:**
```java
@Configuration
public class HealthConfig {
    
    @Bean
    public HealthIndicator customHealthIndicator() {
        return new AbstractHealthIndicator() {
            @Override
            protected void doHealthCheck(Health.Builder builder) throws Exception {
                try {
                    // Custom health check logic
                    boolean isHealthy = checkServiceHealth();
                    
                    if (isHealthy) {
                        builder.up()
                               .withDetail("status", "All services operational")
                               .withDetail("timestamp", Instant.now());
                    } else {
                        builder.down()
                               .withDetail("error", "Service unavailable")
                               .withDetail("timestamp", Instant.now());
                    }
                } catch (Exception e) {
                    builder.down()
                           .withDetail("error", e.getMessage())
                           .withDetail("timestamp", Instant.now());
                }
            }
        };
    }
    
    private boolean checkServiceHealth() {
        try {
            // Quick health check with timeout
            return restTemplate.getForObject(
                "http://service/health", 
                String.class
            ).contains("UP");
        } catch (Exception e) {
            return false;
        }
    }
}
```

**Database Health Check Configuration:**
```yaml
# application.yml
management:
  health:
    db:
      enabled: true
    defaults:
      enabled: true
  endpoint:
    health:
      show-details: always
      show-components: always
  metrics:
    export:
      prometheus:
        enabled: true

spring:
  datasource:
    hikari:
      maximum-pool-size: 20
      minimum-idle: 5
      connection-timeout: 30000
      idle-timeout: 600000
      max-lifetime: 1800000
```

### Q37: Why might a Spring Kafka consumer stop consuming messages after a partition rebalance?

**Answer:** Common causes include consumer group issues and offset management:

**Root Causes:**
1. **Consumer Group Mismatch**
2. **Offset Commit Issues**
3. **Serialization Problems**
4. **Listener Container Configuration**

**Solution:**
```java
@Configuration
@EnableKafka
public class KafkaConfig {
    
    @Bean
    public ConsumerFactory<String, String> consumerFactory() {
        Map<String, Object> props = new HashMap<>();
        props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
        props.put(ConsumerConfig.GROUP_ID_CONFIG, "user-group");
        props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        props.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, false); // Manual offset management
        props.put(ConsumerConfig.MAX_POLL_RECORDS_CONFIG, 100);
        props.put(ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG, 30000);
        props.put(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG, 10000);
        
        return new DefaultKafkaConsumerFactory<>(props);
    }
    
    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, String> kafkaListenerContainerFactory() {
        ConcurrentKafkaListenerContainerFactory<String, String> factory = 
            new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(consumerFactory());
        factory.getContainerProperties().setAckMode(AckMode.MANUAL_IMMEDIATE);
        factory.getContainerProperties().setPollTimeout(3000);
        factory.setConcurrency(3);
        return factory;
    }
}

@Service
public class KafkaConsumerService {
    
    @KafkaListener(topics = "user-events", groupId = "user-group")
    public void consumeMessage(
            @Payload String message,
            @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
            @Header(KafkaHeaders.RECEIVED_PARTITION_ID) int partition,
            @Header(KafkaHeaders.OFFSET) long offset,
            Acknowledgment acknowledgment) {
        
        try {
            // Process message
            processMessage(message);
            
            // Manually acknowledge
            acknowledgment.acknowledge();
            
            log.info("Processed message from topic: {}, partition: {}, offset: {}", 
                    topic, partition, offset);
                    
        } catch (Exception e) {
            log.error("Error processing message: {}", e.getMessage());
            // Don't acknowledge - message will be redelivered
        }
    }
}
```

### Q38: Why does a scheduled job work under normal load but miss executions during heavy traffic?

**Answer:** Thread pool exhaustion and scheduling conflicts:

**Root Cause:**
```java
// Default single-threaded scheduler
@Scheduled(fixedRate = 5000)
public void processScheduledTask() {
    // Long-running task blocks scheduler thread
    heavyProcessing();
}
```

**Solution:**
```java
@Configuration
@EnableScheduling
public class SchedulerConfig {
    
    @Bean
    public TaskScheduler taskScheduler() {
        ThreadPoolTaskScheduler scheduler = new ThreadPoolTaskScheduler();
        scheduler.setPoolSize(10); // Multiple threads
        scheduler.setThreadNamePrefix("scheduled-task-");
        scheduler.setWaitForTasksToCompleteOnShutdown(true);
        scheduler.setAwaitTerminationSeconds(60);
        scheduler.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        return scheduler;
    }
}

@Service
public class ScheduledService {
    
    @Scheduled(fixedRate = 5000)
    @Async("taskExecutor")
    public CompletableFuture<Void> processScheduledTask() {
        try {
            // Process asynchronously
            heavyProcessing();
            log.info("Scheduled task completed");
        } catch (Exception e) {
            log.error("Scheduled task failed: {}", e.getMessage());
        }
        return CompletableFuture.completedFuture(null);
    }
    
    @Bean
    public TaskExecutor taskExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(5);
        executor.setMaxPoolSize(10);
        executor.setQueueCapacity(100);
        executor.setThreadNamePrefix("async-task-");
        executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        executor.initialize();
        return executor;
    }
}
```

### Q39: How would you optimize a REST API that processes 10K requests/second but has response times spiking over 3 seconds?

**Answer:** Multi-layered optimization approach:

**1. Caching Strategy:**
```java
@Service
public class UserService {
    
    @Cacheable(value = "users", key = "#id", unless = "#result == null")
    public User findById(Long id) {
        return userRepository.findById(id).orElse(null);
    }
    
    @CacheEvict(value = "users", key = "#user.id")
    public User updateUser(User user) {
        return userRepository.save(user);
    }
}

@Configuration
@EnableCaching
public class CacheConfig {
    
    @Bean
    public CacheManager cacheManager() {
        RedisCacheManager.Builder builder = RedisCacheManager
            .RedisCacheManagerBuilder
            .fromConnectionFactory(redisConnectionFactory())
            .cacheDefaults(cacheConfiguration());
        return builder.build();
    }
    
    private RedisCacheConfiguration cacheConfiguration() {
        return RedisCacheConfiguration.defaultCacheConfig()
            .entryTtl(Duration.ofMinutes(10))
            .disableCachingNullValues()
            .serializeKeysWith(RedisSerializationContext.SerializationPair
                .fromSerializer(new StringRedisSerializer()))
            .serializeValuesWith(RedisSerializationContext.SerializationPair
                .fromSerializer(new GenericJackson2JsonRedisSerializer()));
    }
}
```

**2. Database Optimization:**
```java
@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    
    @QueryHints(value = @QueryHint(name = "org.hibernate.fetchSize", value = "100"))
    @Query("SELECT u FROM User u WHERE u.status = :status")
    Stream<User> findByStatusStream(@Param("status") String status);
    
    // Batch operations
    @Modifying
    @Query("UPDATE User u SET u.lastLogin = :timestamp WHERE u.id IN :ids")
    int updateLastLogin(@Param("timestamp") LocalDateTime timestamp, 
                        @Param("ids") List<Long> ids);
}
```

**3. Async Processing:**
```java
@RestController
public class UserController {
    
    @Autowired
    private UserService userService;
    
    @Autowired
    private TaskExecutor taskExecutor;
    
    @GetMapping("/users/{id}")
    public CompletableFuture<ResponseEntity<User>> getUser(@PathVariable Long id) {
        return CompletableFuture
            .supplyAsync(() -> userService.findById(id), taskExecutor)
            .thenApply(user -> ResponseEntity.ok(user))
            .exceptionally(throwable -> ResponseEntity.status(500).build());
    }
}
```

**4. Connection Pooling:**
```yaml
spring:
  datasource:
    hikari:
      maximum-pool-size: 50
      minimum-idle: 10
      connection-timeout: 20000
      idle-timeout: 300000
      max-lifetime: 1200000
      leak-detection-threshold: 60000

  redis:
    lettuce:
      pool:
        max-active: 50
        max-idle: 10
        min-idle: 5
        max-wait: 5000ms
```

### Q40: How do you handle a dependency on an unreliable third-party API that frequently experiences timeouts or 502 errors?

**Answer:** Implement resilience patterns:

**1. Circuit Breaker with Fallback:**
```java
@Service
public class ExternalApiService {
    
    @Autowired
    private RestTemplate restTemplate;
    
    @CircuitBreaker(name = "externalApi", fallbackMethod = "fallbackResponse")
    @Retryable(value = {ResourceAccessException.class}, maxAttempts = 2)
    public ExternalResponse callExternalApi(Request request) {
        try {
            ResponseEntity<ExternalResponse> response = restTemplate.postForEntity(
                "https://external-api.com/endpoint", 
                request, 
                ExternalResponse.class
            );
            
            if (response.getStatusCode().is2xxSuccessful()) {
                return response.getBody();
            } else {
                throw new ExternalApiException("API returned status: " + response.getStatusCode());
            }
        } catch (ResourceAccessException e) {
            throw new ExternalApiException("Connection timeout: " + e.getMessage());
        }
    }
    
    public ExternalResponse fallbackResponse(Request request, Exception ex) {
        log.warn("External API failed, using fallback: {}", ex.getMessage());
        
        // Return cached or default response
        return ExternalResponse.builder()
            .status("FALLBACK")
            .message("Service temporarily unavailable")
            .timestamp(Instant.now())
            .build();
    }
}
```

**2. Configuration:**
```yaml
resilience4j:
  circuitbreaker:
    instances:
      externalApi:
        failureRateThreshold: 50
        waitDurationInOpenState: 60s
        slidingWindowSize: 10
        minimumNumberOfCalls: 5
        permittedNumberOfCallsInHalfOpenState: 3
        automaticTransitionFromOpenToHalfOpenEnabled: true
  retry:
    instances:
      externalApi:
        maxAttempts: 3
        waitDuration: 1s
        retryExceptions:
          - org.springframework.web.client.ResourceAccessException
          - org.springframework.web.client.HttpServerErrorException
  timelimiter:
    instances:
      externalApi:
        timeoutDuration: 3s
```

**3. Timeout Configuration:**
```java
@Configuration
public class RestTemplateConfig {
    
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate() {
        RestTemplate restTemplate = new RestTemplate();
        
        HttpComponentsClientHttpRequestFactory factory = 
            new HttpComponentsClientHttpRequestFactory();
        factory.setConnectTimeout(2000); // 2 seconds connect timeout
        factory.setReadTimeout(3000);    // 3 seconds read timeout
        
        restTemplate.setRequestFactory(factory);
        
        // Add retry interceptor
        restTemplate.setInterceptors(List.of(new RetryInterceptor()));
        
        return restTemplate;
    }
}

@Component
public class RetryInterceptor implements ClientHttpRequestInterceptor {
    
    @Override
    public ClientHttpResponse intercept(
            HttpRequest request, 
            byte[] body, 
            ClientHttpRequestExecution execution) throws IOException {
        
        int retryCount = 0;
        int maxRetries = 2;
        
        while (retryCount <= maxRetries) {
            try {
                return execution.execute(request, body);
            } catch (IOException e) {
                if (retryCount == maxRetries) {
                    throw e;
                }
                retryCount++;
                try {
                    Thread.sleep(1000 * retryCount); // Exponential backoff
                } catch (InterruptedException ie) {
                    Thread.currentThread().interrupt();
                    throw e;
                }
            }
        }
        return execution.execute(request, body);
    }
}
```

---

**This comprehensive guide covers all the major topics asked in Virtusa Java Developer interviews. Practice these concepts and scenarios to excel in your interview!**
