# 🌱 Spring Boot & Spring MVC (Rapid-Fire)

> 🔑 **Master Keyword:** **"SADIE"** → Spring-Autoconfigure, Dependency-Injection, Embedded-server, IoC-container, Everything configured

---

## 🚀 Section 1: Spring Boot Basics

### Q1: Spring Boot vs Spring Framework?
🔑 **Keyword: "OADJ"** → Opinionated-Auto-Deploy-JAR

| Feature | Spring Framework | Spring Boot |
|---|---|---|
| Config | Manual XML / Java | Auto-configured |
| Server | External (Tomcat WAR) | Embedded (JAR) |
| Setup time | Long | Seconds (Spring Initializr) |
| Boilerplate | High | Minimal |

> Spring Boot = Spring + Opinionated defaults + Embedded server + Auto-configuration

---

### Q2: `@SpringBootApplication` — What does it do?
🔑 **Keyword: "CEC"** → Configuration+EnableAutoConfig+ComponentScan

```java
@SpringBootApplication  // = @Configuration + @EnableAutoConfiguration + @ComponentScan
public class MyApp {
    public static void main(String[] args) {
        SpringApplication.run(MyApp.class, args);
    }
}
```

| Part | Role |
|---|---|
| `@Configuration` | This class is a bean source |
| `@EnableAutoConfiguration` | Auto-configure based on classpath |
| `@ComponentScan` | Scan for components in package |

---

### Q3: How does Auto-configuration work?
🔑 **Keyword: "CCA"** → Classpath→ConditionalAnnotation→AutoConfig

1. `@EnableAutoConfiguration` reads `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`
2. Each auto-config class uses conditions: `@ConditionalOnClass`, `@ConditionalOnMissingBean`
3. If conditions match → beans are registered automatically

Example: H2 on classpath + no custom `DataSource` → Spring Boot creates in-memory H2 `DataSource`.

---

### Q4: What are Spring Boot Starters?
🔑 **Keyword: "SBB"** → Starter=Bundle-of-Blessed-dependencies

| Starter | Includes |
|---|---|
| `spring-boot-starter-web` | Spring MVC + Jackson + Tomcat |
| `spring-boot-starter-data-jpa` | Spring Data JPA + Hibernate |
| `spring-boot-starter-security` | Spring Security |
| `spring-boot-starter-test` | JUnit + Mockito + AssertJ |
| `spring-boot-starter-actuator` | Metrics + Health + Monitoring |

---

### Q5: Configuration Properties — Priority Order?
🔑 **Keyword: "CESO"** → Command-line > Env-vars > Specific-profile > base-properties

From highest to lowest priority:
1. **Command Line Args** `--server.port=9090`
2. **OS Environment Variables** `SERVER_PORT=9090`
3. **Profile-specific** `application-prod.properties`
4. **Base** `application.properties`

---

### Q6: Spring Profiles?
🔑 **Keyword: "ADPT"** → Activate Different Properties per Target

```properties
# application.properties
spring.profiles.active=dev

# application-dev.properties  
spring.datasource.url=jdbc:h2:mem:devdb

# application-prod.properties
spring.datasource.url=jdbc:mysql://prod-server/mydb
```

Runtime: `java -jar app.jar --spring.profiles.active=prod`

Bean-level: `@Profile("dev")` — activates only for dev profile.

---

### Q7: `CommandLineRunner` vs `ApplicationRunner`?
🔑 **Keyword: "CLR-ATR"** → CLR=String-args, ApplicationRunner=ApplicationArguments

```java
// CommandLineRunner — raw String args
@Bean
public CommandLineRunner init(UserRepository repo) {
    return args -> repo.save(new User("Admin"));
}

// ApplicationRunner — structured args with named options
@Bean
public ApplicationRunner runner() {
    return args -> {
        String name = args.getOptionValues("name").get(0);
    };
}
```
Runs after Spring context loads. Used for: data seeding, startup tasks, warmup.

---

### Q8: Spring Boot Actuator?
🔑 **Keyword: "HMEI"** → Health/Metrics/Env/Info endpoints

```properties
management.endpoints.web.exposure.include=*  # expose all
```

| Endpoint | Purpose |
|---|---|
| `/actuator/health` | App health (UP/DOWN) |
| `/actuator/info` | App info |
| `/actuator/metrics` | Performance metrics |
| `/actuator/env` | Environment + properties |
| `/actuator/beans` | All Spring beans |
| `/actuator/httptrace` | HTTP request trace |

---

### Q9: JAR vs WAR packaging?
🔑 **Keyword: "JS-WO"** → JAR=Self-contained, WAR=Old-containers

| JAR (Default) | WAR |
|---|---|
| Executable, self-contained | Deploy to external server |
| Embedded Tomcat | External Tomcat/Wildfly |
| `java -jar app.jar` | Need extends `SpringBootServletInitializer` |
| Best for Microservices / Docker | Legacy enterprise environments |

---

### Q10: Embedded Tomcat — How does it work?
🔑 **Keyword: "EPT"** → Embedded=Programmatic-Tomcat-Startup

- Tomcat included as library dependency (not installed separately)
- Spring Boot programmatically starts Tomcat on `main()` call
- Creates Servlet context, deploys your app inside
- Runs on port 8080 by default

---

### Q11: DevTools Features?
🔑 **Keyword: "ALC"** → Auto-restart, LiveReload, Cache-disabled

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-devtools</artifactId>
    <scope>runtime</scope>
</dependency>
```

- **Auto-Restart:** Detects classpath changes → fast restart
- **LiveReload:** Triggers browser refresh on static resource change
- **Property Defaults:** Disables template caching for dev

---

## 🌐 Section 2: Spring MVC & REST

### Q12: Spring MVC Request Flow?
🔑 **Keyword: "DHCVR"** → DispatcherServlet→Handler-mapping→Controller→ViewResolver→Response

```
HTTP Request
    → DispatcherServlet (front controller)
        → HandlerMapping (find controller method)
            → HandlerAdapter → Controller.method()
                → returns ModelAndView / ResponseBody
                    → ViewResolver (if view)
                        → Render → HTTP Response
```

---

### Q13: `@Controller` vs `@RestController`?
🔑 **Keyword: "VJ"** → View-return vs JSON-return

```java
@Controller
public class WebController {
    @GetMapping("/home")
    public String home() { return "home"; } // returns view name
}

@RestController  // = @Controller + @ResponseBody on every method
public class ApiController {
    @GetMapping("/api/users")
    public List<User> users() { return userService.getAll(); } // returns JSON
}
```

---

### Q14: `@PathVariable` vs `@RequestParam`?
🔑 **Keyword: "PU-RQ"** → PathVariable=URL-path, RequestParam=?query

```java
// PathVariable: /users/123
@GetMapping("/users/{id}")
public User getUser(@PathVariable Long id) { ... }

// RequestParam: /users?name=alice&active=true
@GetMapping("/users")
public List<User> search(@RequestParam String name,
                         @RequestParam(required = false) Boolean active) { ... }
```

---

### Q15: `@RequestBody` vs `@ResponseBody`?
🔑 **Keyword: "RB-IN-OUT"** → RequestBody=deserialize-in, ResponseBody=serialize-out

```java
@PostMapping("/users")
public @ResponseBody User createUser(@RequestBody User user) {
    // @RequestBody: JSON → User object (deserialization)
    // @ResponseBody: User → JSON response (serialization)
    return userService.save(user);
}
```

---

### Q16: Key Spring MVC Annotations?
🔑 **Keyword: "MGPPD"** → Mapping/Get/Post/Put/Delete shortcuts

| Annotation | Purpose |
|---|---|
| `@RestController` | All methods return JSON |
| `@GetMapping` | HTTP GET |
| `@PostMapping` | HTTP POST |
| `@PutMapping` | HTTP PUT |
| `@DeleteMapping` | HTTP DELETE |
| `@RequestBody` | Deserialize request body |
| `@ResponseBody` | Serialize to response |
| `@PathVariable` | Capture from URL path |
| `@RequestParam` | Capture from query string |
| `@RequestHeader` | Read HTTP headers |
| `@ExceptionHandler` | Handle exceptions in controller |

---

### Q17: Exception Handling in Spring MVC?
🔑 **Keyword: "EGA"** → ExceptionHandler+GlobalControllerAdvice

```java
// Local — handles only in this controller
@ExceptionHandler(UserNotFoundException.class)
public ResponseEntity<String> handleNotFound(UserNotFoundException e) {
    return ResponseEntity.status(404).body(e.getMessage());
}

// Global — handles across all controllers
@RestControllerAdvice  // = @ControllerAdvice + @ResponseBody
public class GlobalExceptionHandler {
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleAll(Exception e) {
        return ResponseEntity.status(500).body(new ErrorResponse(e.getMessage()));
    }
}
```

---

### Q18: Filters vs Interceptors?
🔑 **Keyword: "FS-SI"** → Filter=Servlet/before-Spring, Interceptor=Spring-context

| Feature | Filter | Interceptor |
|---|---|---|
| Level | Servlet container | Spring MVC |
| Runs before | DispatcherServlet | Controller method |
| Access to Spring | ❌ | ✅ |
| Interface | `javax.servlet.Filter` | `HandlerInterceptor` |
| Best for | Encoding, security, CORS | Logging, auth checks, timing |

```java
// Filter — web.xml level
@Component
public class LogFilter implements Filter {
    public void doFilter(ServletRequest req, ServletResponse res, FilterChain chain) { ... }
}

// Interceptor — Spring MVC level
@Component
public class AuthInterceptor implements HandlerInterceptor {
    public boolean preHandle(HttpServletRequest req, HttpServletResponse res, Object handler) { ... }
    public void postHandle(...) { ... }
    public void afterCompletion(...) { ... }
}
```

---

### Q19: Input Validation?
🔑 **Keyword: "VABS"** → @Valid+@Annotations+BindingResult+@Size/@NotNull

```java
// DTO
public class UserRequest {
    @NotNull
    @Size(min = 2, max = 50)
    private String name;

    @Email
    @NotEmpty
    private String email;

    @Min(18)
    private int age;
}

// Controller
@PostMapping("/users")
public ResponseEntity<?> create(@Valid @RequestBody UserRequest req) {
    // @Valid triggers validation, throws MethodArgumentNotValidException on failure
    return ResponseEntity.ok(userService.save(req));
}

// Handle in @ControllerAdvice
@ExceptionHandler(MethodArgumentNotValidException.class)
public ResponseEntity<?> handleValidation(MethodArgumentNotValidException e) { ... }
```

---

### Q20: Content Negotiation?
🔑 **Keyword: "ACH"** → Accept-Header-determines-format

- Same endpoint returns JSON or XML based on client `Accept` header
- Spring uses `ContentNegotiatingViewResolver`
- `Accept: application/json` → JSON
- `Accept: application/xml` → XML (need Jackson XML module)

---

## 💉 Section 3: Spring Core — DI & IoC

### Q21: IoC vs Dependency Injection?
🔑 **Keyword: "IoC-invert, DI-inject"** → IoC=control moved to framework, DI=inject dependencies

- **IoC (Inversion of Control):** You don't create objects; the framework does
- **DI (Dependency Injection):** Framework injects dependencies when creating objects

```java
// Without DI (you control creation)
UserService service = new UserService(new UserRepository(new JdbcTemplate(...)));

// With DI (Spring controls creation and injection)
@Service
public class UserService {
    @Autowired  // Spring injects this!
    private UserRepository userRepository;
}
```

---

### Q22: Types of Dependency Injection?
🔑 **Keyword: "CFS"** → Constructor/Field/Setter injection

```java
// 1. Constructor Injection (RECOMMENDED — immutable, testable)
@Service
public class UserService {
    private final UserRepository repo;
    public UserService(UserRepository repo) { this.repo = repo; }
}

// 2. Field Injection (concise but hard to test)
@Service
public class UserService {
    @Autowired
    private UserRepository repo;
}

// 3. Setter Injection (optional dependencies)
@Service
public class UserService {
    private UserRepository repo;
    @Autowired
    public void setRepo(UserRepository repo) { this.repo = repo; }
}
```

---

### Q23: `@Component` vs `@Service` vs `@Repository` vs `@Controller`?
🔑 **Keyword: "CSRC-roles"** → Component=generic, Service=business, Repository=data, Controller=web

All are specializations of `@Component` (all auto-detected by ComponentScan):

| Annotation | Layer | Extra behavior |
|---|---|---|
| `@Component` | Generic | Just a Spring bean |
| `@Service` | Business logic | No extra, semantic clarity |
| `@Repository` | Data access | Translates DB exceptions to Spring exceptions |
| `@Controller` | Web layer | Handles HTTP requests |

---

### Q24: `@Autowired` vs `@Qualifier` vs `@Primary`?
🔑 **Keyword: "AQP"** → Autowired=find-by-type, Qualifier=specify-name, Primary=default-bean

```java
// Multiple implementations of PaymentService
@Component @Primary  // default when no qualifier specified
public class PayPalService implements PaymentService { }

@Component
public class StripeService implements PaymentService { }

// Usage
@Autowired  // injects PayPalService (primary)
private PaymentService paymentService;

@Autowired
@Qualifier("stripeService")  // explicitly specify which one
private PaymentService stripePayment;
```

---

### Q25: Bean Scopes?
🔑 **Keyword: "SPRRS"** → Singleton/Prototype/Request/Response/Session

| Scope | Description |
|---|---|
| `singleton` (default) | One instance per ApplicationContext |
| `prototype` | New instance every time requested |
| `request` | One per HTTP request (web) |
| `session` | One per HTTP session (web) |

```java
@Bean
@Scope("prototype")
public ExpensiveObject expensive() { return new ExpensiveObject(); }
```

---

### Q26: `@Bean` vs `@Component`?
🔑 **Keyword: "BC-TM"** → @Bean=third-party-manual, @Component=your-class-auto

- `@Component` → on **YOUR class** → auto-detected via component scan
- `@Bean` → on a **method** in `@Configuration` class → you manually construct the bean (useful for third-party classes)

```java
// @Component usage
@Component
public class MyService { }

// @Bean usage — for external library you can't annotate
@Configuration
public class Config {
    @Bean
    public DataSource dataSource() {
        return new HikariDataSource(hikariConfig);
    }
}
```

---

*End of File — Spring Boot & Spring MVC*
