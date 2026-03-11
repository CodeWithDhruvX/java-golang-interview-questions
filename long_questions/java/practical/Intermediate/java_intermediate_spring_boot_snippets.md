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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which injection approach do you prefer and why?
**Your Response:** I strongly prefer constructor injection. First, it allows me to make the dependencies final, which makes the object immutable and thread-safe. Second, it makes unit testing much easier - I can just call `new ConstructorService(mockRepo)` without needing to start up the Spring context. Third, Spring will detect circular dependencies at startup if there are any. Field injection might look cleaner but it has these significant drawbacks.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the functional difference between these annotations?
**Your Response:** Functionally, they all do the same thing - they tell Spring to create a singleton bean. The difference is semantic and behavioral. `@Component` is the generic one. `@Service` indicates business logic layer. `@Controller` is for web controllers. `@Repository` is special because it also translates persistence exceptions into Spring's DataAccessException hierarchy. Using the right annotation makes the code more readable and self-documenting, and in the case of @Repository, you get automatic exception translation.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How many times will "creating MyService" be printed?
**Your Response:** It will only be printed once. This is a key feature of Spring's @Configuration classes. Spring uses CGLIB to create a subclass of the configuration class that intercepts calls to @Bean methods. So even though we call myService() twice, Spring intercepts the second call and returns the same singleton instance from the first call. If we had used @Component instead of @Configuration, each call would create a new instance, which is usually not what we want for beans.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and how does prototype scope work?
**Your Response:** The output is false because prototype scope creates a new instance every time we request the bean. Unlike the default singleton scope where Spring creates one instance and reuses it, prototype scope gives us a fresh instance each time we call getBean(). This is useful for stateful objects where each consumer needs its own instance. However, be careful - Spring doesn't manage the complete lifecycle of prototype beans after creation, so you're responsible for cleanup.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring resolve ambiguity when multiple beans implement the same interface?
**Your Response:** Spring has a clear precedence for resolving bean ambiguity. The most specific approach is @Qualifier, which explicitly tells Spring which bean to inject by name. If there's no qualifier, Spring looks for a bean marked with @Primary, which serves as the default choice. If neither qualifier nor primary is specified, Spring throws an exception because it can't decide. The priority order is: @Qualifier first, then @Primary, then just type matching. This gives you fine-grained control over dependency injection when you have multiple implementations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @Lazy do and when would you use it?
**Your Response:** @Lazy defers bean initialization until it's actually needed, rather than creating all singleton beans at startup. This can significantly improve application startup time, especially for beans that are expensive to create or might not be used in every execution path. It's particularly useful for optional features, background processing components, or beans that depend on external systems that might not be available during startup. The trade-off is that any initialization issues will only be discovered when the bean is first accessed.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the lifecycle order of these methods?
**Your Response:** The lifecycle order is: first the constructor runs, then @PostConstruct executes after dependency injection is complete, and finally @PreDestroy runs when the application context shuts down. The key point is that @PostConstruct is the right place for initialization logic because at that point, all dependencies have been injected. If you put initialization in the constructor, the injected dependencies might still be null. @PreDestroy is perfect for cleanup tasks like closing connections or releasing resources.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between ApplicationContext and BeanFactory?
**Your Response:** ApplicationContext is the full-featured Spring container that extends BeanFactory with additional enterprise features. While BeanFactory provides basic dependency injection with lazy loading, ApplicationContext adds event publication, internationalization support, AOP proxying, lifecycle callbacks like @PostConstruct, and eager bean initialization. In Spring Boot applications, you should always use ApplicationContext since it provides the complete Spring experience. BeanFactory is mainly relevant for resource-constrained environments like mobile apps where you might want the minimal DI container.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the different ways to inject values with @Value?
**Your Response:** @Value supports multiple syntaxes for different needs. The ${} syntax reads values from properties files and supports default values with the colon notation. The #{} syntax enables Spring Expression Language, allowing you to evaluate expressions, call methods, or perform calculations. You can even combine them and access system properties, environment variables, or call static methods. This flexibility makes @Value powerful for injecting configuration values, though for complex configuration structures, @ConfigurationProperties is often a better choice.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use @ConfigurationProperties vs @Value?
**Your Response:** @ConfigurationProperties is ideal for structured configuration with multiple related properties. It binds all properties with a common prefix to a typed POJO, giving you type safety, IDE auto-completion, and validation support. This is much cleaner than scattering @Value annotations everywhere. @Value is better for individual properties or when you need Spring Expression Language. For example, use @ConfigurationProperties for database connection settings, and @Value for feature flags or simple configuration values.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring handle circular dependencies?
**Your Response:** Spring handles circular dependencies differently based on the injection type. With constructor injection, Spring fails fast at startup with a BeanCurrentlyInCreationException, which is good because it prevents subtle runtime issues. With field injection, Spring might resolve it using proxy objects, but it still logs a warning. The best solutions are to break the circular dependency by extracting shared logic into a third bean, or use @Lazy on one of the dependencies to defer its initialization. Circular dependencies are often a design smell indicating the need for refactoring.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do Spring's application events work?
**Your Response:** Spring's event system provides a publish-subscribe pattern that decouples components. Publishers use ApplicationEventPublisher to broadcast events, while subscribers use @EventListener to receive them. This eliminates direct dependencies between components. By default, events are processed synchronously in the same thread, but you can add @Async for asynchronous processing. This is perfect for cross-cutting concerns like audit logging, notifications, or reacting to domain events without tight coupling between the business logic and the side effects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use BeanDefinitionRegistryPostProcessor?
**Your Response:** I would use BeanDefinitionRegistryPostProcessor when I need to register beans programmatically at startup, typically when building frameworks or libraries. This interface allows me to modify the bean definition registry before regular bean initialization happens. For example, Spring Data uses this to automatically create repository implementations for interface definitions. It's particularly useful when you don't know the exact beans at compile time or need to dynamically create beans based on configuration or scanning results.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do Spring profiles work and when would you use them?
**Your Response:** Spring profiles allow me to conditionally register beans based on the active environment. When I annotate a bean with @Profile("dev"), Spring only creates that bean when the "dev" profile is active. This is perfect for environment-specific configurations like using mock services in development, real implementations in production, or different database connections for testing. I can activate profiles via the spring.profiles.active property, environment variables, or programmatically. I can also use negative expressions like @Profile("!prod") to create beans for all environments except production.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can you control bean initialization order in Spring?
**Your Response:** Spring doesn't guarantee initialization order by default, which can cause issues when one bean depends on another being fully initialized first. I use @DependsOn to explicitly declare dependencies between beans. For example, if my UserRepository needs the database to be initialized first, I'd annotate it with @DependsOn("databaseInit"). This ensures Spring creates the DatabaseInitializer bean before the UserRepository. This is crucial for setup sequences like schema initialization, cache warming, or when beans need to access resources that other beans prepare during startup.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @RestController do and how does it work?
**Your Response:** @RestController is a convenient annotation that combines @Controller and @ResponseBody. When I use it, Spring automatically serializes the return value of each method to JSON or XML and writes it directly to the HTTP response body. I don't need to annotate each method with @ResponseBody individually. Under the hood, Spring uses Jackson (or another HTTP message converter) to convert my Java objects to JSON. This makes building REST APIs much cleaner - I just return POJOs and Spring handles the serialization automatically.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use ResponseEntity vs just returning a POJO?
**Your Response:** I use ResponseEntity when I need full control over the HTTP response. While returning a POJO works for simple cases, ResponseEntity lets me set custom status codes, headers, and control the response body. For example, when creating a resource, I can return a 201 Created status with a Location header pointing to the new resource. I can also set custom headers like Cache-Control or handle error responses with appropriate status codes. ResponseEntity is essential for building RESTful APIs that follow HTTP semantics properly.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does @RequestBody work and what are common error scenarios?
**Your Response:** @RequestBody tells Spring to deserialize the HTTP request body into a Java object using Jackson. Spring checks the Content-Type header to determine how to parse the body. If the Content-Type header is missing or doesn't match what Spring expects, it returns a 415 Unsupported Media Type error. If the JSON is malformed or doesn't match the object structure, Spring throws a HttpMessageNotReadableException and returns a 400 Bad Request. I always make sure clients send the correct Content-Type header and valid JSON to avoid these errors.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @PathVariable and @RequestParam?
**Your Response:** @PathVariable extracts values from the URL path itself, like the ID in /products/42. It's used for required path segments. @RequestParam extracts values from the query string after the question mark, like ?currency=EUR. The key difference is that @PathVariable is part of the URL structure and required by default, while @RequestParam is optional and can have default values. I use @PathVariable for resource identification and @RequestParam for filtering, pagination, or optional parameters.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you access HTTP headers and cookies in Spring controllers?
**Your Response:** I use @RequestHeader to extract HTTP headers and @CookieValue to read cookie values. Both annotations work similarly - they bind the header or cookie value to method parameters. I can make them optional with required=false and provide default values. For example, I'd use @RequestHeader("Authorization") to get the auth token for API security, or @CookieValue("sessionId") to track user sessions. Spring automatically handles the extraction and type conversion, making it much cleaner than manually parsing the HttpServletRequest.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you handle error responses in a REST API?
**Your Response:** I use ResponseEntity to return appropriate HTTP status codes and error messages. For example, when a user isn't found, I return a 404 status with a descriptive error body. I can use the Optional.map() pattern to elegantly handle both success and error cases. For success, I return ResponseEntity.ok(user), and for not found, I return ResponseEntity.status(HttpStatus.NOT_FOUND) with an error object. This approach gives clients clear feedback about what went wrong and maintains consistent error handling across the API.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do the different mapping annotations work in Spring?
**Your Response:** Spring provides specific annotations for each HTTP method: @GetMapping for GET, @PostMapping for POST, @PutMapping for PUT, @PatchMapping for PATCH, and @DeleteMapping for DELETE. These are just shortcuts for @RequestMapping with the method parameter specified. Using the specific annotations makes the code more readable and clearly shows what HTTP operation each endpoint handles. If a client sends the wrong HTTP method to an endpoint, Spring automatically returns a 405 Method Not Allowed error.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does validation work with @Valid in Spring?
**Your Response:** When I annotate a parameter with @Valid, Spring automatically runs bean validation on the object using annotations like @NotBlank, @Email, and @Min. If validation fails, Spring throws a MethodArgumentNotValidException and returns a 400 Bad Request response by default. The response contains detailed validation errors, which is helpful for API clients. I can customize the error response format using @ControllerAdvice and @ExceptionHandler methods to create consistent error responses across my entire application.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use @ResponseStatus vs ResponseEntity?
**Your Response:** I use @ResponseStatus when I only need to set a fixed HTTP status code for an endpoint, making the code cleaner. For example, DELETE operations typically return 204 No Content, and POST operations return 201 Created. @ResponseStatus is declarative and simpler than wrapping the return value in ResponseEntity just to set the status. However, if I need to control headers or the response body dynamically, I'd use ResponseEntity instead. @ResponseStatus is perfect for straightforward cases where the status is always the same.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why would you return CompletableFuture from a REST endpoint?
**Your Response:** I return CompletableFuture when I need to handle long-running operations without blocking server threads. This is crucial for scalability because it frees up Tomcat worker threads while the async computation runs in the background. For example, if I'm calling an external API or doing heavy computation, the thread is released back to the pool and can handle other requests. When the CompletableFuture completes, Spring automatically sends the response. This approach significantly improves throughput under load, especially for I/O-bound operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle CORS in Spring Boot applications?
**Your Response:** I configure CORS using @CrossOrigin on controllers or globally via WebMvcConfigurer. The @CrossOrigin annotation lets me specify which origins are allowed, what HTTP methods are permitted, and other CORS settings. Without proper CORS configuration, browsers block cross-origin requests for security reasons. I can also configure CORS globally using addCorsMappings() method, which is useful when I want consistent CORS settings across all endpoints. For production, I typically restrict origins to specific domains rather than using wildcards.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does content negotiation work in Spring?
**Your Response:** Spring supports content negotiation through the Accept header. When a client makes a request, it can specify what response format it prefers using the Accept header. I can configure my endpoints to support multiple formats using the produces attribute. Spring then matches the client's preference against the available formats and automatically serializes the response accordingly. For JSON, it uses Jackson by default, and for XML, I need to add Jackson Dataformat XML to the classpath. If the client requests a format that's not supported, Spring returns a 406 Not Acceptable error.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle file uploads in Spring Boot?
**Your Response:** I handle file uploads using MultipartFile parameters with the @RequestParam annotation. Spring automatically parses multipart form data and binds the file to the MultipartFile object, which gives me access to the original filename, content type, and file contents. I can read the file using getInputStream() or save it using transferTo(). For mixed uploads containing both files and JSON data, I use @RequestPart. I also configure upload limits via spring.servlet.multipart.max-file-size property to prevent oversized uploads from affecting server performance.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does AOP work in Spring and what are pointcuts?
**Your Response:** Spring AOP allows me to add cross-cutting concerns like logging, security, or transactions without modifying business logic code. The @Before advice runs before matched methods execute. Pointcuts define which methods should be intercepted - for example, execution(* com.example.service.*.*(..)) matches all methods in all classes in the service package. This enables me to add consistent logging across all service methods. Spring AOP uses proxy-based weaving, so it only works with Spring-managed beans and method calls that go through the proxy.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is @Around advice and when would you use it?
**Your Response:** @Around is the most powerful type of AOP advice because it wraps the entire method execution. I can run code before and after the method, modify method arguments, change the return value, or even prevent the method from executing. The key is calling pjp.proceed() to execute the original method. I use @Around for performance monitoring, caching, or retry logic. For example, I can time method execution by recording the start time, calling proceed(), then calculating and logging the elapsed time. This gives me complete control over the method execution flow.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the different types of AOP advice and when do they execute?
**Your Response:** Spring AOP provides several advice types that execute at different points. @AfterReturning runs only when a method completes successfully and I can bind the return value. @AfterThrowing executes only when a method throws an exception, letting me access the exception. @After runs in both cases, similar to a finally block in Java. I use @AfterReturning for post-processing successful operations like audit logging, @AfterThrowing for error handling and logging, and @After for cleanup tasks that need to run regardless of success or failure.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do pointcut expressions work in Spring AOP?
**Your Response:** Pointcut expressions define which methods should be intercepted by aspects. They use a declarative syntax where I can match based on method names, packages, annotations, parameters, and more. For example, execution(* com.app.service.*.*(..)) matches all methods in the service package. I can combine expressions with logical operators like &&, ||, and ! to create complex matching rules. I also use @Pointcut annotations to create reusable pointcut definitions that I can reference across multiple advice methods, making the code more maintainable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the AOP self-invocation problem and how do you solve it?
**Your Response:** The self-invocation problem occurs when a method within the same class calls another method that has AOP advice. Since Spring AOP uses proxies, internal method calls bypass the proxy, so the advice doesn't execute. For example, if processAll() calls this.process(), the @Transactional on process() won't work. I solve this by either injecting a self-reference, getting the bean from ApplicationContext, or using AspectJ compile-time weaving. The self-injection approach is cleanest - I just @Autowired OrderService self and call self.process() instead of this.process().

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between JDK dynamic proxy and CGLIB?
**Your Response:** Spring AOP uses two types of proxies depending on whether the bean implements interfaces. If a bean implements interfaces, Spring uses JDK dynamic proxy which creates a proxy that implements the same interfaces. If the bean doesn't implement interfaces, Spring uses CGLIB to create a subclass proxy. The key difference is that with JDK proxy, only interface methods can be intercepted, while CGLIB can proxy all methods including concrete class methods. Since Spring Boot 2.0, CGLIB is the default, which means @Transactional works on private and package-private methods too.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does @Transactional actually work under the hood?
**Your Response:** @Transactional is implemented as AOP advice - specifically an @Around advice. The Spring proxy intercepts method calls, starts a transaction before executing the method, calls the actual method via proceed(), then commits or rolls back the transaction based on the outcome. A crucial point is that if I catch and handle an exception within the method, the transaction won't roll back because the proxy doesn't see the exception. For rollback to occur, the exception must propagate from the method to the proxy. This is why I either let unchecked exceptions bubble up or use TransactionAspectSupport.currentTransactionStatus().setRollbackOnly().

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you control the order of multiple aspects in Spring?
**Your Response:** I control aspect execution order using the @Order annotation. Lower numeric values have higher priority and execute first. For example, @Order(1) runs before @Order(2). This is crucial when aspects have dependencies - like security should run before logging. In @Around advice, the highest priority aspect wraps the others, so it executes first on entry and last on exit. I can also use Ordered.HIGHEST_PRECEDENCE and Ordered.LOWEST_PRECEDENCE constants for better readability instead of magic numbers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the requirements for a JPA entity?
**Your Response:** A JPA entity needs three key things: the @Entity annotation to mark it as a persistent class, an @Id field to identify the primary key, and a no-argument constructor which JPA uses for instantiation. The constructor can be protected to prevent direct instantiation. I use @Table to customize the table name, and @Column to control column properties like nullability and uniqueness. Without these annotations, JPA uses sensible defaults - the class name becomes the table name, and field names become column names.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does transaction propagation work in Spring?
**Your Response:** Transaction propagation controls how transactions behave when one transactional method calls another. REQUIRED is the default - it joins any existing transaction or creates a new one if none exists. REQUIRES_NEW always suspends any current transaction and starts a new one, useful for independent operations that need their own commit. NESTED creates a savepoint within the existing transaction, allowing partial rollbacks. I choose the propagation type based on business requirements - like logging might use REQUIRES_NEW to ensure audit records are saved even if the main transaction rolls back.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When does Spring rollback transactions?
**Your Response:** By default, Spring only rolls back transactions for unchecked exceptions - RuntimeExceptions and Errors. Checked exceptions don't trigger rollback even if they propagate. This is a common gotcha - if I catch a checked exception and don't rethrow it, the transaction won't rollback. I can override this behavior using rollbackFor = Exception.class to rollback on all exceptions. The key is understanding that Spring's default rollback behavior follows the principle that runtime exceptions typically indicate programming errors that should rollback, while checked exceptions might be recoverable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does optimistic locking work in JPA?
**Your Response:** Optimistic locking prevents lost updates using a version field that JPA manages automatically. When I annotate a field with @Version, JPA increments it on each update and includes it in the WHERE clause. If two threads try to update the same entity simultaneously, the first one succeeds and increments the version, but the second one's UPDATE fails because the version no longer matches, throwing an OptimisticLockException. I handle this by catching the exception and retrying the operation. This approach is more scalable than pessimistic locking because it doesn't hold database locks during the transaction.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do Spring Data JPA query methods work?
**Your Response:** Spring Data JPA automatically generates SQL queries from method names at startup. When I define a method like findByEmailAndActive, Spring parses the name and creates the corresponding JPQL query with WHERE email = ? AND active = ?. This eliminates the need to write boilerplate repository implementations. I can also use @Query for custom JPQL when the method name gets too complex. The beauty is that these are type-safe and validated at startup, so I catch query errors early rather than at runtime.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the purpose of @Modifying in Spring Data JPA?
**Your Response:** @Modifying is required for UPDATE and DELETE queries in Spring Data JPA. Without it, Spring assumes the query is a SELECT and tries to return results, which causes an exception. When I annotate a query method with @Modifying, Spring knows it's a bulk operation that modifies data. I can also use clearAutomatically = true to clear the persistence context after the update, which prevents stale entity state. This annotation is essential for any bulk update or delete operations in repository methods.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the JPA first-level cache?
**Your Response:** The JPA first-level cache, also called the persistence context, is automatically managed within a transaction. When I load an entity using EntityManager.find(), JPA caches it in memory. If I request the same entity again in the same transaction, JPA returns the cached instance instead of hitting the database. This ensures entity identity consistency - the same entity instance is always returned for the same database row. The cache is cleared when the transaction ends. This is different from the second-level cache which is optional and spans multiple transactions.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between flush() and commit() in JPA?
**Your Response:** flush() executes the pending SQL statements against the database but doesn't commit the transaction, while commit() makes the changes permanent. I use flush() when I need the database-generated values like IDs before the transaction ends. For example, if I persist an entity and immediately need its ID for subsequent operations, I'll call flush() to get the ID. The changes are still within the transaction and can be rolled back until commit() is called at the end of the @Transactional method.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do JPA cascades work?
**Your Response:** JPA cascades allow me to propagate entity operations from parent to child entities automatically. CascadeType.ALL means all operations like persist, merge, and remove cascade to related entities. I also use orphanRemoval = true to automatically delete child entities when they're removed from the parent's collection. This simplifies object graph management - when I persist a Post with Comments, all comments are automatically persisted. The key is understanding that cascades reduce boilerplate code but I need to be careful about unintended side effects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the advantages of named queries in JPA?
**Your Response:** Named queries have several advantages over inline queries. First, they're validated at application startup, so I catch syntax errors early rather than at runtime. Second, they're cached by the JPA provider for better performance. Third, they centralize query definitions with the entity classes, making the code more organized and maintainable. I define them using @NamedQuery annotations on the entity class or in orm.xml files. While inline queries in createQuery() work for simple cases, named queries are better for complex, reusable queries that are part of the entity's contract.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @SpringBootTest and @WebMvcTest?
**Your Response:** @SpringBootTest loads the entire Spring application context, making it ideal for integration tests that need all components like databases, message queues, and the full stack. It's slower but comprehensive. @WebMvcTest is more focused - it only loads the web layer including controllers, filters, and MVC configuration, while mocking out service and repository layers with @MockBean. This makes controller tests much faster and more isolated. I use @WebMvcTest for unit testing my REST endpoints and @SpringBootTest for end-to-end integration testing.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test REST controllers with MockMvc?
**Your Response:** I use MockMvc to test my REST controllers without starting a real HTTP server. MockMvc simulates HTTP requests and goes through the complete Spring MVC pipeline including controllers, filters, and interceptors. I can verify the response status, headers, and JSON content using fluent assertions. For example, I mock the service layer with @MockBean, then perform a GET request and expect status 200 with specific JSON values. This approach is much faster than integration tests while still providing comprehensive coverage of the web layer.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test POST endpoints with request bodies?
**Your Response:** For testing POST endpoints, I use MockMvc to send JSON request bodies and verify the complete response. I create the request object, serialize it to JSON using ObjectMapper, then perform the POST request with the correct content type. I verify both the success response - like 201 Created with a Location header for resource creation - and error scenarios like 400 for validation failures. It's important to test both happy path and error cases to ensure the API behaves correctly. I also test that the service layer is called with the correct data using Mockito verify.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does @DataJpaTest do?
**Your Response:** @DataJpaTest is a specialized test annotation that configures only the JPA layer for testing. It sets up an in-memory H2 database, configures JPA, and creates Spring Data repository beans - but doesn't load the full application context. This makes repository tests much faster and more focused. Each test method runs in a transaction that's automatically rolled back afterward, ensuring test isolation. I use @DataJpaTest when I want to test my repository methods and JPA mappings without the overhead of loading the entire Spring context.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @MockBean and @Mock?
**Your Response:** @MockBean and @Mock serve different testing purposes. @MockBean is used in Spring integration tests to replace a real bean in the Spring application context with a Mockito mock. This allows me to test how my application behaves with mocked dependencies. @Mock is for pure unit tests without Spring context - it creates a standalone Mockito mock that I inject manually using @InjectMocks. @Mock is faster because it doesn't start Spring, but @MockBean is necessary when testing Spring-specific features like controllers or aspects. I use @Mock for service layer unit tests and @MockBean for web layer integration tests.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle exceptions globally in Spring Boot?
**Your Response:** I use @RestControllerAdvice to implement global exception handling across all controllers. This annotation allows me to define centralized exception handlers using @ExceptionHandler for specific exception types. For example, I can handle UserNotFoundException by returning a 404 status with a descriptive error message, and validation failures by returning 400 with detailed field errors. This approach ensures consistent error responses across my entire API and eliminates the need for try-catch blocks in individual controller methods. It's the Spring way of implementing cross-cutting error handling.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is ProblemDetail in Spring Boot?
**Your Response:** ProblemDetail is Spring's implementation of RFC 9457, which standardizes HTTP error response formats. Instead of returning custom error objects, I can use ProblemDetail to create consistent, machine-readable error responses that follow web standards. It includes standard fields like type, title, status, and detail, plus I can add custom properties. This makes my API more interoperable and easier for clients to parse. I enable it with spring.mvc.problemdetails.enabled=true and use it in my @ExceptionHandler methods to return standardized error responses.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle database state in Spring tests?
**Your Response:** Spring provides different transaction management for different test types. @DataJpaTest tests are automatically transactional and roll back by default, keeping the database clean between tests. However, @SpringBootTest tests are not transactional by default, so database changes persist. I either add @Transactional to @SpringBootTest tests or use @Sql annotations with cleanup scripts. The automatic rollback in @DataJpaTest is convenient because it ensures test isolation without manual cleanup, but I need to be aware that @SpringBootTest behaves differently.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use TestRestTemplate over MockMvc?
**Your Response:** I use TestRestTemplate for full integration tests when I need to test the complete application stack including servlet filters, security, CORS, and actual HTTP processing. Unlike MockMvc which tests the MVC layer in isolation, TestRestTemplate starts a real embedded server and makes actual HTTP requests. This is slower but more realistic, catching issues that MockMvc might miss like configuration problems or filter chain issues. I use TestRestTemplate for end-to-end testing and MockMvc for faster, focused controller testing. Both approaches complement each other in a comprehensive testing strategy.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring Boot Actuator and how do you use it?
**Your Response:** Spring Boot Actuator provides production-ready features like health checks, metrics, and application information. I configure it in application.yml to expose specific endpoints like /actuator/health, /actuator/info, and /actuator/metrics. The health endpoint returns the application status and component health, which Kubernetes uses for readiness and liveness probes. I can also expose environment variables and configuration for debugging. In production, I secure these endpoints with Spring Security and only expose what's necessary. Actuator is essential for monitoring and managing Spring Boot applications in production.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does caching work with @Cacheable in Spring?
**Your Response:** Spring's caching abstraction allows me to cache method results easily. When I annotate a method with @Cacheable, Spring checks if a result exists in the cache before executing the method. If found, it returns the cached value; if not, it executes the method and caches the result. I use @CacheEvict to remove stale entries when data changes, and @CachePut to always update the cache. Spring supports multiple cache providers like Caffeine for in-memory caching or Redis for distributed caching. This significantly improves performance for frequently accessed data that doesn't change often.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement method-level security in Spring?
**Your Response:** I use @PreAuthorize to implement method-level security using Spring Expression Language. This annotation allows me to define complex access rules directly on methods. For example, hasRole('ADMIN') restricts access to users with admin role, and I can combine conditions using AND/OR operators. I can also access method parameters and authentication details, like checking if a user can only access their own data by comparing the userId parameter with authentication.principal.id. This approach provides fine-grained security that's more maintainable than scattered security checks in the code.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you configure Spring Security for JWT authentication?
**Your Response:** I configure Spring Security using a SecurityFilterChain bean where I define the security rules. For JWT stateless authentication, I disable CSRF since we're not using sessions, set the session creation policy to STATELESS, and configure authorization rules. I permit all requests to public endpoints like /api/public/**, restrict admin endpoints to ADMIN role, and require authentication for all other requests. I also add a custom JWT filter before the UsernamePasswordAuthenticationFilter to validate tokens on each request. This setup provides secure, stateless authentication suitable for REST APIs.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you schedule background tasks in Spring Boot?
**Your Response:** I use @Scheduled along with @EnableScheduling to create background tasks. Spring provides different scheduling options - fixedRate runs the task at regular intervals regardless of execution time, fixedDelay waits for the specified delay after each completion, and cron allows complex scheduling using cron expressions. I can also specify an initialDelay to start tasks after application startup. These tasks run in a default thread pool, but I can customize the executor for more control. This is perfect for periodic maintenance tasks, data cleanup, or polling external systems.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use CommandLineRunner or ApplicationRunner?
**Your Response:** I use CommandLineRunner and ApplicationRunner to execute code after Spring Boot starts up but before the application receives requests. This is perfect for initialization tasks like loading reference data, setting up caches, or running migrations. CommandLineRunner receives raw command line arguments as String[], while ApplicationRunner provides parsed ApplicationArguments with options and parameter values. I use @Order to control execution order when I have multiple runners. These are essential for application bootstrap tasks that need to run once during startup.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement feature flags in Spring Boot?
**Your Response:** I implement feature flags using @ConditionalOnProperty, which conditionally creates beans based on configuration properties. This allows me to enable or disable features without code changes. I can set havingValue to match specific property values, and matchIfMissing to define the default behavior when the property isn't present. This is perfect for experimental features, A/B testing, or optional integrations that should only be active in certain environments. By combining with @Configuration classes, I can conditionally load entire feature modules based on a single property.

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
**A:** `RestTemplate` is synchronous and blocking. `WebClient` is reactive and non-blocking — better throughput. Spring Boot 3 recommends `WebClient` for all new code. For tests, use `MockWebServer` (OkHttp) to mock of HTTP server.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between RestTemplate and WebClient?
**Your Response:** RestTemplate is the traditional synchronous, blocking HTTP client, while WebClient is the modern reactive, non-blocking alternative. RestTemplate blocks the calling thread until the response arrives, which limits scalability. WebClient uses reactive streams and doesn't block threads, allowing much better throughput under load. Spring Boot 3 recommends WebClient for all new code. I use WebClient when I need high concurrency or when working in a reactive stack, and RestTemplate only for legacy code or simple blocking calls. For testing WebClient, I use MockWebServer from OkHttp.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does @Async work in Spring and what are the gotchas?
**Your Response:** @Async allows me to execute methods asynchronously in a separate thread pool. When I annotate a method with @Async, Spring intercepts the call and runs it in a background thread from its TaskExecutor. The key gotcha is the same as AOP - self-invocation doesn't work because it bypasses the Spring proxy. I must call the async method from another bean or inject a self-reference. The method should return CompletableFuture, Future, or void. I enable this feature with @EnableAsync and can customize the thread pool for better control over execution. This is perfect for fire-and-forget operations or parallel processing.

---

> 🔖 **Last read:** <!-- update here -->
