# Java Intermediate — Spring Boot Practical Snippets

> **Topics:** Dependency Injection (`@Component`, `@Autowired`, `@Bean`), REST API (`@RestController`, `ResponseEntity`), AOP (`@Aspect`, `@Around`), JPA (`@Entity`, `@Transactional`), Spring Security, Spring Testing (`@WebMvcTest`, `MockMvc`), Exception Handling (`@ControllerAdvice`), Configuration (`@ConfigurationProperties`)

---

## 📋 Reading Progress

- [ ] **Section 1:** Dependency Injection & Bean Lifecycle (Q1–Q15)
- [ ] **Section 2:** REST API — Building Endpoints (Q16–Q28)
- [ ] **Section 3:** AOP — Aspects & Pointcuts (Q29–Q36)
- [ ] **Section 4:** JPA & Transactions (Q37–Q47)
- [ ] **Section 5:** Spring Testing & Exception Handling (Q48–Q65)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Dependency Injection & Bean Lifecycle (Q1–Q15)

### 1. Constructor vs Field Injection
**Q: Which is preferred and why?**
```java
import org.springframework.stereotype.*;
import org.springframework.beans.factory.annotation.*;

// AVOID: field injection
@Service
class FieldService {
    @Autowired private UserRepository repo; // not testable without Spring context
}

// PREFER: constructor injection
@Service
class ConstructorService {
    private final UserRepository repo;

    // @Autowired is optional if only one constructor (Spring 4.3+)
    ConstructorService(UserRepository repo) {
        this.repo = repo;
    }
}
```
**A:** Constructor injection is preferred because: (1) fields can be `final` (immutable), (2) no Spring context needed in unit tests — just call `new ConstructorService(mockRepo)`, (3) circular dependencies are detected at startup.

---

### 2. @Component vs @Service vs @Repository vs @Controller
**Q: What is the functional difference?**
```java
@Component   // generic Spring-managed bean
@Service     // business logic layer (semantic alias of @Component)
@Repository  // data access layer — also translates SQLExceptions to Spring DataAccessException
@Controller  // MVC controller — processes web requests
@RestController // = @Controller + @ResponseBody
```
**A:** Functionally, `@Service`, `@Component`, and `@Controller` are identical — they all register a singleton bean. `@Repository` additionally wraps persistence exceptions. Use the right one semantically; Spring scans all of them.

---

### 3. @Bean in @Configuration
**Q: What is the output / how many instances are created?**
```java
import org.springframework.context.annotation.*;

@Configuration
class AppConfig {
    @Bean
    public MyService myService() {
        System.out.println("creating MyService");
        return new MyService();
    }
}

@SpringBootApplication
class Main {
    public static void main(String[] args) {
        var ctx = SpringApplication.run(Main.class, args);
        ctx.getBean(MyService.class);
        ctx.getBean(MyService.class); // second call
    }
}
```
**A:** `creating MyService` is printed **once**. Spring's `@Configuration` uses CGLIB to intercept `@Bean` method calls — returning the same singleton each time. Without `@Configuration` (using `@Component` instead), each call would create a new instance.

---

### 4. @Scope — Singleton vs Prototype
**Q: What is the output?**
```java
import org.springframework.context.annotation.*;
import org.springframework.stereotype.*;

@Component
@Scope("prototype")
class PrototypeBean { }

@Service
class ConsumerService {
    @Autowired ApplicationContext ctx;
    void check() {
        PrototypeBean a = ctx.getBean(PrototypeBean.class);
        PrototypeBean b = ctx.getBean(PrototypeBean.class);
        System.out.println(a == b);
    }
}
```
**A:** `false`. Prototype scope creates a **new instance** per `getBean()` call. The default scope is `singleton` — one instance per `ApplicationContext`.

---

### 5. @Primary and @Qualifier
**Q: What happens without @Qualifier?**
```java
import org.springframework.context.annotation.*;
import org.springframework.stereotype.*;
import org.springframework.beans.factory.annotation.*;

interface PaymentService { void pay(); }

@Service @Primary
class StripeService   implements PaymentService { public void pay() { System.out.println("stripe"); } }

@Service
class PaypalService   implements PaymentService { public void pay() { System.out.println("paypal"); } }

@Service
class OrderService {
    @Autowired PaymentService ps; // injects StripeService due to @Primary

    @Autowired @Qualifier("paypalService")
    PaymentService ps2; // explicit qualifier
}
```
**A:** Without qualifiers, `@Primary` wins. With `@Qualifier("beanName")`, the explicitly named bean is injected. Disambiguation order: `@Qualifier` > `@Primary` > type match.

---

### 6. @Lazy — Deferred Initialization
**Q: What is the output?**
```java
import org.springframework.context.annotation.*;
import org.springframework.stereotype.*;

@Component @Lazy
class HeavyService {
    HeavyService() { System.out.println("HeavyService created"); }
    void serve() { System.out.println("serving"); }
}

@SpringBootApplication
class Main {
    public static void main(String[] args) {
        var ctx = SpringApplication.run(Main.class, args);
        System.out.println("context started");
        ctx.getBean(HeavyService.class).serve(); // created here
    }
}
```
**A:**
```
context started
HeavyService created
serving
```
`@Lazy` delays bean creation until first use. Without it, Spring creates all singletons eagerly at startup.

---

### 7. @PostConstruct and @PreDestroy
**Q: What is the lifecycle order?**
```java
import jakarta.annotation.*;
import org.springframework.stereotype.*;

@Component
class LifecycleBean {
    LifecycleBean() { System.out.println("1. constructor"); }

    @PostConstruct
    void init() { System.out.println("2. @PostConstruct"); }

    @PreDestroy
    void cleanup() { System.out.println("3. @PreDestroy"); }
}
```
**A:** On startup: `1. constructor` → `2. @PostConstruct`. On context close: `3. @PreDestroy`. Use `@PostConstruct` for initialization after dependency injection, not in the constructor (dependencies aren't yet injected there).

---

### 8. ApplicationContext vs BeanFactory
**Q: What is the difference?**
```java
// BeanFactory: basic DI container, lazy loading
// ApplicationContext: extends BeanFactory + adds:
//   - Event publication (@EventListener)
//   - MessageSource (i18n)
//   - AOP auto-proxy
//   - @PostConstruct / @PreDestroy support
//   - Eager singleton initialization

@SpringBootApplication
class Main {
    public static void main(String[] args) {
        ApplicationContext ctx = SpringApplication.run(Main.class, args);
        System.out.println(ctx.getBeanDefinitionCount() + " beans registered");
    }
}
```
**A:** Spring Boot auto-registers dozens of beans. Always use `ApplicationContext`. `BeanFactory` is only relevant for embedded/constrained environments.

---

### 9. @Value — Injecting Properties
**Q: What is the output?**
```java
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

// application.properties: app.name=MyApp   app.maxRetries=3
@Component
class ConfigReader {
    @Value("${app.name}")
    private String appName;

    @Value("${app.maxRetries:5}") // default = 5 if property missing
    private int maxRetries;

    @Value("#{2 * T(Math).PI}")  // Spring Expression Language
    private double twoPi;

    void print() {
        System.out.printf("%s, retries=%d, 2π=%.4f%n", appName, maxRetries, twoPi);
    }
}
```
**A:** `MyApp, retries=3, 2π=6.2832`. `${}` reads from properties; `#{}` is Spring EL. The `:default` syntax provides fallback values.

---

### 10. @ConfigurationProperties — Type-Safe Config
**Q: What is the advantage over @Value?**
```java
import org.springframework.boot.context.properties.*;
import org.springframework.stereotype.*;

// application.yml:
// db:
//   host: localhost
//   port: 5432
//   pool-size: 10

@Component
@ConfigurationProperties(prefix = "db")
class DbProperties {
    private String host;
    private int port;
    private int poolSize;
    // getters and setters...
}
```
**A:** `@ConfigurationProperties` binds a whole group of properties to a POJO with type safety, IDE completion, and validation (`@Validated` + `@NotNull`). Prefer over scattered `@Value` annotations for structured configuration.

---

### 11. Circular Dependency
**Q: What happens at startup?**
```java
@Service class A {
    @Autowired B b; // A depends on B
}
@Service class B {
    @Autowired A a; // B depends on A — CIRCULAR!
}
```
**A:** With constructor injection → `BeanCurrentlyInCreationException` at startup (good — fail fast). With field injection → Spring may resolve it via proxy injection but still warns. **Fix:** break the cycle by extracting shared logic to a third bean, or use `@Lazy` on one injection point.

---

### 12. @EventListener — Application Events
**Q: What is the output?**
```java
import org.springframework.context.*;
import org.springframework.stereotype.*;

@Component
class OrderService {
    @Autowired ApplicationEventPublisher publisher;
    void place(String item) {
        publisher.publishEvent(new OrderPlacedEvent(this, item));
    }
}

class OrderPlacedEvent extends ApplicationEvent {
    String item;
    OrderPlacedEvent(Object source, String item) { super(source); this.item = item; }
}

@Component
class EmailService {
    @EventListener
    void onOrder(OrderPlacedEvent e) {
        System.out.println("Email sent for: " + e.item);
    }
}
```
**A:** `Email sent for: <item>`. Events decouple publishers from subscribers — components don't need direct references. By default synchronous in the same thread; use `@Async` for async dispatch.

---

### 13. BeanDefinitionRegistryPostProcessor — Programmatic Bean Registration
**Q: When would you use this?**
```java
import org.springframework.beans.factory.support.*;
import org.springframework.context.annotation.*;

@Component
class DynamicRegistrar implements BeanDefinitionRegistryPostProcessor {
    public void postProcessBeanDefinitionRegistry(BeanDefinitionRegistry registry) {
        RootBeanDefinition def = new RootBeanDefinition(MyDynamicService.class);
        registry.registerBeanDefinition("myDynamicService", def);
        System.out.println("Dynamic bean registered");
    }
    public void postProcessBeanFactory(ConfigurableListableBeanFactory f) {}
}
```
**A:** Useful for frameworks that register beans dynamically at startup (e.g., JPA repositories, MyBatis mappers). Spring Data uses this internally to register repository interfaces as beans.

---

### 14. @Profile — Environment-Specific Beans
**Q: Which bean is created?**
```java
import org.springframework.context.annotation.*;
import org.springframework.stereotype.*;

@Service @Profile("dev")
class MockPaymentService implements PaymentService {
    public void pay() { System.out.println("mock payment"); }
}

@Service @Profile("prod")
class RealPaymentService implements PaymentService {
    public void pay() { System.out.println("real payment"); }
}
// Active profile set via: spring.profiles.active=dev
```
**A:** Only `MockPaymentService` is created in `dev`. Use `@Profile("!prod")` to activate unless prod. Profiles are set via `spring.profiles.active` property or `SPRING_PROFILES_ACTIVE` env var.

---

### 15. Bean Initialization Order — @DependsOn
**Q: Is initialization order guaranteed?**
```java
import org.springframework.context.annotation.*;
import org.springframework.stereotype.*;

@Component @DependsOn("databaseInit")
class UserRepository {
    UserRepository() { System.out.println("UserRepository created"); }
}

@Component("databaseInit")
class DatabaseInitializer {
    DatabaseInitializer() { System.out.println("Database initialized"); }
}
```
**A:**
```
Database initialized
UserRepository created
```
Spring does not guarantee bean initialization order unless you use `@DependsOn`. Always use it when a bean must be created after another (e.g., schema init before repository).

---

## Section 2: REST API — Building Endpoints (Q16–Q28)

### 16. @RestController — Basic GET Endpoint
**Q: What does this return?**
```java
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/users")
class UserController {

    @GetMapping("/{id}")
    public User getUser(@PathVariable Long id) {
        return new User(id, "Alice"); // auto-serialized to JSON
    }
}
```
**A:** Returns `{"id":1,"name":"Alice"}`. `@RestController` = `@Controller` + `@ResponseBody` — no need to annotate each method. Jackson serializes the returned object to JSON automatically.

---

### 17. ResponseEntity — Full HTTP Control
**Q: What HTTP response does this return?**
```java
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/items")
class ItemController {

    @PostMapping
    public ResponseEntity<Item> create(@RequestBody Item item) {
        Item saved = save(item); // service call
        return ResponseEntity
            .created(URI.create("/items/" + saved.getId()))
            .body(saved);
    }
}
```
**A:** HTTP `201 Created` with `Location: /items/{id}` header and the saved `Item` as JSON body. `ResponseEntity` gives full control over status code, headers, and body.

---

### 18. @RequestBody — JSON Deserialization
**Q: What happens if Content-Type is missing or wrong?**
```java
@PostMapping("/orders")
public Order createOrder(@RequestBody Order order) {
    // Spring uses Jackson to deserialize request body
    return orderService.create(order);
}
```
**A:** If `Content-Type: application/json` header is missing → `415 Unsupported Media Type`. If body is malformed JSON → `400 Bad Request` with `HttpMessageNotReadableException`.

---

### 19. @RequestParam vs @PathVariable
**Q: What are the differences?**
```java
@GetMapping("/products/{id}")
public Product get(
    @PathVariable Long id,              // /products/42
    @RequestParam(defaultValue = "USD") String currency // ?currency=EUR
) {
    return productService.get(id, currency);
}
```
**A:** `@PathVariable` extracts from the URI template. `@RequestParam` extracts from query string. `@PathVariable` required by default; `@RequestParam` can have `required=false` or `defaultValue`.

---

### 20. @RequestHeader and @CookieValue
**Q: How do you access HTTP headers and cookies?**
```java
@GetMapping("/profile")
public String profile(
    @RequestHeader("Authorization") String authHeader,
    @RequestHeader(value = "X-Custom", required = false) String custom,
    @CookieValue(value = "sessionId", defaultValue = "anon") String sessionId
) {
    return "auth=" + authHeader + " session=" + sessionId;
}
```
**A:** `@RequestHeader` extracts named HTTP headers. `@CookieValue` reads cookies. Both support `required` and `defaultValue`.

---

### 21. ResponseEntity with Error Responses
**Q: What is the output for a 404?**
```java
@GetMapping("/users/{id}")
public ResponseEntity<Object> getUser(@PathVariable Long id) {
    return userService.findById(id)
        .map(u -> ResponseEntity.ok((Object) u))
        .orElse(ResponseEntity
            .status(HttpStatus.NOT_FOUND)
            .body(Map.of("error", "User not found", "id", id)));
}
```
**A:** If user exists: `200 OK` with User JSON. If not: `404 Not Found` with `{"error":"User not found","id":99}`.

---

### 22. @RequestMapping — HTTP Method Matching
**Q: What HTTP methods does this match?**
```java
@RestController
class VerbController {
    @GetMapping("/items")    public List<Item> list()                    { return List.of(); }
    @PostMapping("/items")   public Item       create(@RequestBody Item i) { return i; }
    @PutMapping("/items/{id}")  public Item    update(@PathVariable Long id, @RequestBody Item i) { return i; }
    @PatchMapping("/items/{id}") public Item   patch(@PathVariable Long id, @RequestBody Map<String, Object> changes) { return null; }
    @DeleteMapping("/items/{id}") public void  delete(@PathVariable Long id) {}
}
```
**A:** Each annotation maps to the corresponding HTTP method. `@GetMapping` = `@RequestMapping(method = RequestMethod.GET)`. Sending a POST to a GET-only endpoint → `405 Method Not Allowed`.

---

### 23. @Valid — Bean Validation
**Q: What happens if validation fails?**
```java
import jakarta.validation.constraints.*;
import jakarta.validation.*;

record CreateUserRequest(
    @NotBlank String username,
    @Email String email,
    @Min(18) int age
) {}

@PostMapping("/users")
public User create(@Valid @RequestBody CreateUserRequest req) {
    return userService.create(req);
}
```
**A:** If `username` is blank → `400 Bad Request` with `MethodArgumentNotValidException`. Spring auto-creates the error response. Use `@ControllerAdvice` to customize the error format.

---

### 24. @ResponseStatus — Declarative Status Codes
**Q: What does this return?**
```java
@ResponseStatus(HttpStatus.NO_CONTENT)
@DeleteMapping("/users/{id}")
public void delete(@PathVariable Long id) {
    userService.delete(id);
}

@ResponseStatus(HttpStatus.CREATED)
@PostMapping("/users")
public User create(@RequestBody User u) {
    return userService.save(u);
}
```
**A:** DELETE → `204 No Content`. POST → `201 Created`. `@ResponseStatus` is a simpler alternative to returning `ResponseEntity` when you only need to set the status code.

---

### 25. Async REST with CompletableFuture
**Q: What is the benefit of returning CompletableFuture?**
```java
import java.util.concurrent.*;
import org.springframework.scheduling.annotation.*;

@EnableAsync
@RestController
class AsyncController {

    @GetMapping("/slow")
    public CompletableFuture<String> slow() {
        return CompletableFuture.supplyAsync(() -> {
            Thread.sleep_unchecked(2000); // simulate slow I/O
            return "result";
        });
    }
}
```
**A:** The Tomcat/Jetty worker thread is freed while the async computation runs — better throughput under load. Spring automatically writes the response when the `CompletableFuture` completes.

---

### 26. @CrossOrigin — CORS Configuration
**Q: What does this allow?**
```java
@CrossOrigin(origins = "https://myfrontend.com", methods = {GET, POST})
@RestController
@RequestMapping("/api")
class ApiController {
    @GetMapping("/data")
    public String data() { return "ok"; }
}
```
**A:** Allows cross-origin requests from `https://myfrontend.com` for GET and POST. Without `@CrossOrigin`, browsers block cross-origin requests. Global CORS can be configured via `WebMvcConfigurer.addCorsMappings()`.

---

### 27. Content Negotiation — JSON vs XML
**Q: How does Spring select the response format?**
```java
@GetMapping(value = "/user", produces = { MediaType.APPLICATION_JSON_VALUE, MediaType.APPLICATION_XML_VALUE })
public User getUser() {
    return new User(1L, "Alice");
}
```
**A:** Spring checks the `Accept` header: `Accept: application/json` → JSON; `Accept: application/xml` → XML (requires Jackson Dataformat XML on classpath). If no match → `406 Not Acceptable`.

---

### 28. Multipart File Upload
**Q: How does Spring handle file uploads?**
```java
@PostMapping(value = "/upload", consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
public ResponseEntity<String> upload(
    @RequestParam("file") MultipartFile file,
    @RequestParam("description") String desc) {

    System.out.println("name: "    + file.getOriginalFilename());
    System.out.println("size: "    + file.getSize());
    System.out.println("type: "    + file.getContentType());
    // file.getInputStream() to read, file.transferTo(path) to save
    return ResponseEntity.ok("uploaded: " + desc);
}
```
**A:** Spring binds `MultipartFile` automatically. Max upload size is configured via `spring.servlet.multipart.max-file-size`. Use `@RequestPart` for mixed multipart (file + JSON).

---

## Section 3: AOP — Aspects & Pointcuts (Q29–Q36)

### 29. @Before Advice — Method Entry Logging
**Q: What is the output?**
```java
import org.aspectj.lang.*;
import org.aspectj.lang.annotation.*;
import org.springframework.stereotype.*;

@Aspect @Component
class LoggingAspect {
    @Before("execution(* com.example.service.*.*(..))")
    public void logBefore(JoinPoint jp) {
        System.out.println("Calling: " + jp.getSignature().getName());
    }
}

// Calling userService.getUser(1L) prints:
// "Calling: getUser"
// then getUser executes
```
**A:** `@Before` runs before the matched method. The pointcut `execution(* com.example.service.*.*(..))` matches any method in any class in the `service` package.

---

### 30. @Around — Most Powerful Advice
**Q: What is the output?**
```java
@Aspect @Component
class TimingAspect {
    @Around("@annotation(Timed)")
    public Object time(ProceedingJoinPoint pjp) throws Throwable {
        long start = System.currentTimeMillis();
        Object result = pjp.proceed(); // call the actual method
        long elapsed = System.currentTimeMillis() - start;
        System.out.println(pjp.getSignature().getName() + " took " + elapsed + "ms");
        return result;
    }
}

@Target(METHOD) @Retention(RUNTIME)
@interface Timed {}

@Service class OrderService {
    @Timed
    public Order process(Long id) { /* ... */ }
}
```
**A:** Prints e.g. `process took 42ms`. `@Around` can modify args, return value, or suppress exceptions. `pjp.proceed()` must be called to execute the original method.

---

### 31. @AfterReturning and @AfterThrowing
**Q: When does each fire?**
```java
@Aspect @Component
class AuditAspect {
    @AfterReturning(pointcut = "execution(* save(..))", returning = "result")
    public void afterSave(Object result) {
        System.out.println("saved: " + result);
    }

    @AfterThrowing(pointcut = "execution(* delete(..))", throwing = "ex")
    public void afterFail(Exception ex) {
        System.out.println("delete failed: " + ex.getMessage());
    }
}
```
**A:** `@AfterReturning` fires only if method returns normally (binding the return value). `@AfterThrowing` fires only if it throws. `@After` fires in both cases (like `finally`).

---

### 32. Pointcut Expressions — Common Patterns
**Q: Match each pattern to its meaning**
```java
execution(* com.app.service.*.*(..))         // any method in service package
execution(public * *(..))                    // any public method
execution(* *Service.get*(Long))             // methods named get* taking Long in classes ending in Service
@annotation(org.springframework.transaction.annotation.Transactional) // has @Transactional
within(com.app.web.*)                        // any joinpoint within web package
args(Long, ..)                               // first arg is Long
bean(userService)                            // only on userService bean
```
**A:** Pointcut expressions are parsed by AspectJ. Combine with `&&`, `||`, `!`. Reuse pointcuts with `@Pointcut` method annotations.

---

### 33. AOP Self-Invocation — The Gotcha
**Q: Does the aspect fire?**
```java
@Service
class OrderService {
    public void processAll() {
        for (Long id : getIds()) {
            process(id); // self-invocation — bypasses AOP proxy!
        }
    }

    @Transactional // WON'T fire new transaction per call due to self-invocation
    public void process(Long id) { /* ... */ }
}
```
**A:** **No.** Spring AOP works via proxy — when `processAll()` calls `this.process()`, it bypasses the proxy. Fix: inject `self` reference (`@Autowired OrderService self`), use `ApplicationContext.getBean()`, or use AspectJ weaving.

---

### 34. JDK Dynamic Proxy vs CGLIB
**Q: When does Spring use each?**
```java
interface UserService { User getUser(Long id); }

@Service
class UserServiceImpl implements UserService {
    public User getUser(Long id) { return new User(id); }
}

// Spring AOP proxy type:
// - If bean implements interface(s) → JDK dynamic proxy (implements same interfaces)
// - If bean has no interfaces → CGLIB subclass proxy
// Force CGLIB: @EnableAspectJAutoProxy(proxyTargetClass = true)
```
**A:** JDK proxy wraps the interface; CGLIB subclasses the concrete class. Implication: `@Transactional` on non-interface methods only works with CGLIB. Spring Boot defaults to CGLIB since 2.0.

---

### 35. @Transactional as AOP
**Q: Does the transaction start before or after the method body?**
```java
@Service
class AccountService {
    @Transactional
    public void transfer(Long from, Long to, BigDecimal amount) {
        // transaction starts here (via @Around proxy)
        debit(from, amount);
        credit(to, amount);
        // transaction commits here (or rolls back on unchecked exception)
    }
}
```
**A:** `@Transactional` is implemented as an `@Around` advice. The proxy starts a transaction, calls `proceed()` (your method), then commits or rolls back. This means exceptions thrown and caught inside the method **do not** roll back.

---

### 36. Aspect Ordering — @Order
**Q: Which aspect runs first?**
```java
@Aspect @Component @Order(1)
class SecurityAspect {
    @Before("execution(* *(..))") void check() { System.out.println("security"); }
}

@Aspect @Component @Order(2)
class LoggingAspect {
    @Before("execution(* *(..))") void log() { System.out.println("logging"); }
}
```
**A:**
```
security
logging
```
Lower `@Order` value = higher priority = runs first (outermost in `@Around`). Use `Ordered.HIGHEST_PRECEDENCE` and `Ordered.LOWEST_PRECEDENCE` for clarity.

---

## Section 4: JPA & Transactions (Q37–Q47)

### 37. @Entity — Basic Mapping
**Q: What is required for a JPA entity?**
```java
import jakarta.persistence.*;

@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, unique = true)
    private String email;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    // No-arg constructor required by JPA spec
    protected User() {}
    public User(String email) { this.email = email; }
}
```
**A:** JPA requires: `@Entity`, `@Id`, and a **no-arg constructor** (can be protected). `@Column` customizes the column mapping. Without `@Table`, the table name defaults to the class name.

---

### 38. @OneToMany — Lazy vs Eager Loading
**Q: What is the default fetch type and what is the N+1 problem?**
```java
@Entity
class Order {
    @Id @GeneratedValue Long id;

    @OneToMany(mappedBy = "order", fetch = FetchType.LAZY) // default
    List<OrderItem> items; // NOT loaded until accessed
}

// N+1 problem:
List<Order> orders = orderRepo.findAll(); // 1 query
for (Order o : orders) {
    int count = o.getItems().size(); // N queries (one per order)!
}

// Fix: @Query("SELECT o FROM Order o JOIN FETCH o.items")
```
**A:** `@OneToMany` defaults to `LAZY`. Accessing `items` outside a transaction throws `LazyInitializationException`. Eagerly fetching all associations causes N+1. Fix with `JOIN FETCH` or `@EntityGraph`.

---

### 39. @Transactional — Propagation
**Q: Does serviceB run in serviceA's transaction?**
```java
@Service
class ServiceA {
    @Autowired ServiceB b;

    @Transactional
    public void doA() {
        // transaction T1 starts
        b.doB(); // what happens here?
    }
}

@Service
class ServiceB {
    @Transactional // REQUIRED (default) — joins existing transaction
    public void doB() {
        // runs in T1 (same transaction)
    }

    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void doBNew() {
        // suspends T1, starts T2 — independent transaction
    }
}
```
**A:** `REQUIRED` (default) joins the existing transaction. `REQUIRES_NEW` suspends it and creates a new one. `NESTED` creates a savepoint within the existing transaction.

---

### 40. @Transactional — Rollback Rules
**Q: Does this roll back?**
```java
@Transactional
public void process() {
    save(entity);
    try {
        riskyOp(); // throws IOException (checked)
    } catch (IOException e) {
        log.warn("handled"); // swallowed!
    }
}

@Transactional(rollbackFor = Exception.class)
public void strictProcess() throws Exception {
    save(entity);
    riskyOp(); // IOException propagates → rollback
}
```
**A:** Default `@Transactional` only rolls back on **unchecked** exceptions (`RuntimeException` + `Error`). Checked exceptions that propagate → **no rollback** by default. Use `rollbackFor = Exception.class` to roll back on checked exceptions too.

---

### 41. @Version — Optimistic Locking
**Q: What happens on concurrent updates?**
```java
@Entity
class Product {
    @Id Long id;
    String name;
    int stock;

    @Version
    int version; // JPA manages this field
}

// Thread 1: load product (version=1), save (version becomes 2)
// Thread 2: load same product (version=1), save → OptimisticLockException!
// because DB row now has version=2

// Fix: catch OptimisticLockException and retry
```
**A:** `@Version` implements optimistic locking with a version counter. JPA adds `WHERE version = 1` to the UPDATE — if another thread already updated, the WHERE clause matches 0 rows → `OptimisticLockException`.

---

### 42. Spring Data JPA — Query Methods
**Q: Does this compile and work?**
```java
import org.springframework.data.jpa.repository.*;

interface UserRepository extends JpaRepository<User, Long> {
    List<User> findByEmailAndActive(String email, boolean active);
    Optional<User> findFirstByOrderByCreatedAtDesc();
    long countByActiveTrue();
    @Query("SELECT u FROM User u WHERE u.age > :minAge")
    List<User> findOlderThan(@Param("minAge") int age);
}
```
**A:** Yes — Spring Data derives SQL from method names at startup. `findByEmailAndActive` → `WHERE email = ? AND active = ?`. `@Query` allows custom JPQL. These work at runtime with no implementation code needed.

---

### 43. @Modifying — Bulk Update/Delete
**Q: What is missing?**
```java
interface UserRepository extends JpaRepository<User, Long> {
    @Query("UPDATE User u SET u.active = false WHERE u.lastLoginDate < :cutoff")
    // @Modifying // ← REQUIRED for UPDATE/DELETE queries
    int deactivateInactive(@Param("cutoff") LocalDate cutoff);
}
```
**A:** `@Modifying` is missing. Without it, Spring throws `InvalidDataAccessApiUsageException` because it treats the query as a SELECT. `@Modifying(clearAutomatically = true)` clears the first-level cache after the update.

---

### 44. EntityManager — First-Level Cache
**Q: Are both queries the same instance?**
```java
@Transactional
public void check() {
    User u1 = em.find(User.class, 1L); // DB query
    User u2 = em.find(User.class, 1L); // cache hit — no DB query
    System.out.println(u1 == u2); // true!
}
```
**A:** `true`. The JPA first-level cache (persistence context) is scoped to the `EntityManager`/transaction. The second `find` returns the same in-memory object — no SQL issued. After the transaction ends, the context is closed.

---

### 45. flush() vs commit()
**Q: Does this execute SQL immediately?**
```java
@Transactional
public void example() {
    User u = new User("alice@example.com");
    em.persist(u);                    // schedules INSERT
    System.out.println(u.getId());    // may be null (before flush)
    em.flush();                       // executes INSERT SQL NOW (but no commit)
    System.out.println(u.getId());    // now has DB-generated ID
    // commit happens at end of @Transactional method
}
```
**A:** `flush()` sends pending SQL to the database within the current transaction. `commit()` makes it permanent. Useful when you need the generated ID before the transaction ends.

---

### 46. Cascade Types
**Q: What cascades to child when parent is persisted/deleted?**
```java
@Entity
class Post {
    @Id @GeneratedValue Long id;

    @OneToMany(mappedBy = "post", cascade = CascadeType.ALL, orphanRemoval = true)
    List<Comment> comments = new ArrayList<>();
}

// em.persist(post) → also persists all comments (CASCADE.PERSIST)
// em.remove(post)  → also removes all comments (CASCADE.REMOVE)
// removing comment from list → deletes comment (orphanRemoval = true)
```
**A:** `CascadeType.ALL` = propagate all operations (PERSIST, MERGE, REMOVE, REFRESH, DETACH). `orphanRemoval = true` deletes child entities when removed from the parent's collection.

---

### 47. @NamedQuery at Class Level
**Q: What is the advantage?**
```java
@Entity
@NamedQuery(name = "User.findActive",
            query = "SELECT u FROM User u WHERE u.active = true")
class User {
    @Id Long id;
    boolean active;
}

// Usage:
List<User> users = em.createNamedQuery("User.findActive", User.class).getResultList();
```
**A:** Named queries are validated at startup (not at runtime) — syntax errors fail fast. They're also cached by the JPA provider. Prefer over string-inline queries in `createQuery()`.

---

## Section 5: Spring Testing & Exception Handling (Q48–Q65)

### 48. @SpringBootTest vs @WebMvcTest
**Q: What is the difference?**
```java
@SpringBootTest       // loads full ApplicationContext — integration test
class FullTest { ... }

@WebMvcTest(UserController.class) // loads only MVC layer — faster
class WebLayerTest {
    @MockBean UserService userService; // mock the service layer
    @Autowired MockMvc mockMvc;
}
```
**A:** `@SpringBootTest` starts the complete context (slow, for integration tests). `@WebMvcTest` loads only controllers, filters, and MVC config — service/repo beans must be `@MockBean`. Use `@WebMvcTest` for controller unit tests.

---

### 49. MockMvc — Testing a GET Endpoint
**Q: What does this test verify?**
```java
@WebMvcTest(UserController.class)
class UserControllerTest {
    @Autowired MockMvc mockMvc;
    @MockBean UserService userService;

    @Test
    void getUser_returnsUser() throws Exception {
        when(userService.findById(1L)).thenReturn(new User(1L, "Alice"));

        mockMvc.perform(get("/api/users/1")
                .accept(MediaType.APPLICATION_JSON))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.name").value("Alice"));
    }
}
```
**A:** Verifies status is `200`, and response JSON has `name = "Alice"`. `MockMvc` dispatches requests through the full MVC pipeline without starting a real HTTP server.

---

### 50. MockMvc — Testing POST with Body
**Q: What does this verify?**
```java
@Test
void createUser_returns201() throws Exception {
    User req = new User(null, "Bob");
    User saved = new User(2L, "Bob");
    when(userService.create(any())).thenReturn(saved);

    mockMvc.perform(post("/api/users")
            .contentType(MediaType.APPLICATION_JSON)
            .content(objectMapper.writeValueAsString(req)))
        .andExpect(status().isCreated())
        .andExpect(header().string("Location", containsString("/users/2")));
}
```
**A:** Verifies the POST returns `201 Created` with a `Location` header pointing to the new resource. Always test error responses too (400, 404, 409).

---

### 51. @DataJpaTest — Repository Tests
**Q: What does @DataJpaTest configure?**
```java
@DataJpaTest // configures H2 in-memory DB, JPA, repositories
class UserRepositoryTest {
    @Autowired UserRepository repo;

    @Test
    void saveAndFind() {
        User saved = repo.save(new User("test@email.com"));
        Optional<User> found = repo.findById(saved.getId());
        assertThat(found).isPresent();
        assertThat(found.get().getEmail()).isEqualTo("test@email.com");
    }
}
```
**A:** `@DataJpaTest` configures an in-memory H2 database, JPA, and Spring Data repositories — fast, isolated from the full context. Each test runs in a transaction that's rolled back after the test.

---

### 52. @MockBean vs @Mock
**Q: What is the difference?**
```java
// @MockBean: adds a Mockito mock to the Spring ApplicationContext (Spring test)
@WebMvcTest(OrderController.class)
class OrderControllerTest {
    @MockBean OrderService orderService; // replaces existing Spring bean with mock
}

// @Mock: plain Mockito mock, no Spring context
class OrderServiceTest {
    @Mock  OrderRepository repo;
    @InjectMocks OrderService service; // injects @Mock into service
}
```
**A:** `@MockBean` is for Spring tests (replaces a bean in context). `@Mock` is for pure unit tests (no Spring context, faster). Prefer `@Mock` + `@InjectMocks` for service/repo unit tests.

---

### 53. @ControllerAdvice — Global Exception Handling
**Q: What is the output when UserNotFoundException is thrown?**
```java
@RestControllerAdvice
class GlobalExceptionHandler {

    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<Map<String, String>> handleNotFound(UserNotFoundException ex) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
            .body(Map.of("error", ex.getMessage()));
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, List<String>>> handleValidation(MethodArgumentNotValidException ex) {
        List<String> errors = ex.getBindingResult().getFieldErrors().stream()
            .map(fe -> fe.getField() + ": " + fe.getDefaultMessage())
            .toList();
        return ResponseEntity.badRequest().body(Map.of("errors", errors));
    }
}
```
**A:** `UserNotFoundException` → `404 Not Found` `{"error":"..."}`. Validation failure → `400 Bad Request` with field error list. `@RestControllerAdvice` applies globally to all controllers.

---

### 54. ProblemDetail — RFC 9457 (Spring 6+)
**Q: What is the standard error format?**
```java
@ExceptionHandler(UserNotFoundException.class)
public ProblemDetail handleNotFound(UserNotFoundException ex) {
    ProblemDetail pd = ProblemDetail.forStatusAndDetail(HttpStatus.NOT_FOUND, ex.getMessage());
    pd.setTitle("User Not Found");
    pd.setProperty("userId", ex.getUserId());
    return pd;
}
// Response: {"type":"about:blank","title":"User Not Found","status":404,
//            "detail":"User 42 not found","userId":42}
```
**A:** `ProblemDetail` implements RFC 9457 — a standard machine-readable format for HTTP error responses. Spring 6/Boot 3 supports it natively. Enable with `spring.mvc.problemdetails.enabled=true`.

---

### 55. @Transactional in Tests — Rollback Behavior
**Q: Is the DB mutated after the test?**
```java
@DataJpaTest // or @SpringBootTest with @Transactional
class RepoTest {
    @Autowired UserRepository repo;

    @Test
    @Transactional // each test is rolled back automatically
    void insertAndCheck() {
        repo.save(new User("test@test.com"));
        assertThat(repo.count()).isEqualTo(1);
    }
    // DB is empty after this test — transaction rolled back
}
```
**A:** `@DataJpaTest` tests are transactional and rolled back by default. `@SpringBootTest` tests are NOT transactional by default — add `@Transactional` explicitly or use `@Sql(scripts = "cleanup.sql", executionPhase = AFTER_TEST_METHOD)`.

---

### 56. TestRestTemplate — Full HTTP Tests
**Q: When do you use TestRestTemplate over MockMvc?**
```java
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class IntegrationTest {
    @Autowired TestRestTemplate restTemplate;

    @Test
    void health() {
        ResponseEntity<String> res = restTemplate.getForEntity("/actuator/health", String.class);
        assertThat(res.getStatusCode()).isEqualTo(HttpStatus.OK);
    }
}
```
**A:** `TestRestTemplate` starts a real embedded server and sends actual HTTP requests — tests the full stack including servlet filters, error handling, CORS, etc. Slower but more realistic than `MockMvc`.

---

### 57. Spring Actuator — Health and Metrics
**Q: What does /actuator/health return?**
```yaml
# application.yml
management:
  endpoints:
    web:
      exposure:
        include: health, info, metrics, env
  endpoint:
    health:
      show-details: always
```
**A:** Returns `{"status":"UP","components":{"db":{"status":"UP"},"diskSpace":{...}}}`. Actuator exposes runtime info. Used by Kubernetes readiness/liveness probes. Secure with Spring Security in production.

---

### 58. @Cacheable — Response Caching
**Q: How many DB calls are made?**
```java
@Service
class ProductService {
    @Cacheable(value = "products", key = "#id")
    public Product getProduct(Long id) {
        System.out.println("DB call for: " + id);
        return repo.findById(id).orElseThrow();
    }

    @CacheEvict(value = "products", key = "#product.id")
    public Product update(Product product) { return repo.save(product); }
}

// productService.getProduct(1L); // "DB call for: 1" printed
// productService.getProduct(1L); // cache hit — nothing printed
// productService.getProduct(2L); // "DB call for: 2" printed
```
**A:** 2 DB calls (one for `id=1`, one for `id=2`). The second call for `id=1` is served from cache. `@CacheEvict` removes stale entries. Backed by Caffeine, Redis, EhCache etc.

---

### 59. Spring Security — @PreAuthorize
**Q: What happens without the ADMIN role?**
```java
@Service
class AdminService {
    @PreAuthorize("hasRole('ADMIN')")
    public void deleteAll() {
        System.out.println("deleting all");
    }

    @PreAuthorize("hasRole('USER') and #userId == authentication.principal.id")
    public User getProfile(Long userId) { return userRepo.findById(userId).orElseThrow(); }
}
```
**A:** Without `ADMIN` role → `AccessDeniedException` → `403 Forbidden`. `@PreAuthorize` uses Spring EL. The second example additionally checks the authenticated user can only access their own profile.

---

### 60. Security Filter Chain — Custom
**Q: What does this configure?**
```java
@Configuration
@EnableWebSecurity
class SecurityConfig {
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        return http
            .csrf(csrf -> csrf.disable())
            .sessionManagement(sm -> sm.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/api/public/**").permitAll()
                .requestMatchers("/api/admin/**").hasRole("ADMIN")
                .anyRequest().authenticated())
            .addFilterBefore(jwtFilter, UsernamePasswordAuthenticationFilter.class)
            .build();
    }
}
```
**A:** Configures stateless JWT authentication: CSRF disabled (stateless), no session, public endpoints open, admin endpoints restricted, all others need auth, JWT filter added before username/password filter.

---

### 61. @Scheduled — Background Tasks
**Q: What does this do?**
```java
@Component
@EnableScheduling
class ScheduledTasks {
    @Scheduled(fixedRate = 60_000)            // every 60 seconds
    public void polling() { System.out.println("polling"); }

    @Scheduled(cron = "0 0 2 * * *")         // every day at 2 AM
    public void dailyCleanup() { System.out.println("cleanup"); }

    @Scheduled(fixedDelay = 5_000, initialDelay = 10_000)
    public void delayed() {} // starts 10s after startup, then 5s after each completion
}
```
**A:** All three patterns are valid. `fixedRate` = period between invocation starts. `fixedDelay` = period between invocation end and next start. `cron` = full cron expression (seconds, minutes, hours, day, month, weekday).

---

### 62. ApplicationRunner and CommandLineRunner
**Q: What is the output order?**
```java
@Component @Order(1)
class First implements CommandLineRunner {
    public void run(String... args) { System.out.println("first"); }
}

@Component @Order(2)
class Second implements ApplicationRunner {
    public void run(ApplicationArguments args) { System.out.println("second"); }
}
```
**A:**
```
first
second
```
Both run after the Spring context starts (just before the app is ready). `CommandLineRunner` receives raw `String[]`; `ApplicationRunner` receives parsed `ApplicationArguments`. `@Order` controls execution order.

---

### 63. @ConditionalOnProperty — Feature Flags
**Q: When is FeatureBean created?**
```java
@Component
@ConditionalOnProperty(name = "feature.experimental", havingValue = "true", matchIfMissing = false)
class ExperimentalFeature {
    ExperimentalFeature() { System.out.println("experimental feature ON"); }
}
// Only created if application.properties has: feature.experimental=true
```
**A:** Conditional bean creation based on a property. `matchIfMissing=false` means if the property is absent, the bean is NOT created. Useful for feature flags, A/B testing, enabling optional integrations.

---

### 64. RestTemplate vs WebClient
**Q: What is the difference?**
```java
// RestTemplate — synchronous (Spring 5 deprecated for new code)
ResponseEntity<User> res = restTemplate.getForEntity("http://service/users/1", User.class);

// WebClient — reactive, non-blocking (preferred)
User user = webClient.get()
    .uri("/users/{id}", 1L)
    .retrieve()
    .bodyToMono(User.class)
    .block(); // block only if in a non-reactive context

// WebClient chaining — preferred in reactive controllers:
Mono<User> userMono = webClient.get()
    .uri("/users/{id}", 1L)
    .retrieve()
    .bodyToMono(User.class);
```
**A:** `RestTemplate` is synchronous and blocking. `WebClient` is reactive and non-blocking — better throughput. Spring Boot 3 recommends `WebClient` for all new code. For tests, use `MockWebServer` (OkHttp) to mock the HTTP server.

---

### 65. @Async — Non-Blocking Methods
**Q: What is the output?**
```java
@SpringBootApplication
@EnableAsync
class App {
    @Async
    public CompletableFuture<String> asyncWork() {
        System.out.println("async thread: " + Thread.currentThread().getName());
        return CompletableFuture.completedFuture("done");
    }

    public static void main(String[] args) throws Exception {
        ApplicationContext ctx = SpringApplication.run(App.class, args);
        App a = ctx.getBean(App.class);
        System.out.println("main thread: " + Thread.currentThread().getName());
        CompletableFuture<String> f = a.asyncWork();
        System.out.println("result: " + f.get());
    }
}
```
**A:**
```
main thread: main
async thread: task-1        (different thread from Spring's TaskExecutor)
result: done
```
`@Async` runs the method in a separate thread. **Must be called via the Spring proxy** — self-invocation bypasses it (same gotcha as AOP). Always return `CompletableFuture`, `Future`, or `void`.

---

> 🔖 **Last read:** <!-- update here -->
