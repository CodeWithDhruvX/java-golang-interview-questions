# 🌱 04 — Spring Boot Fundamentals
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Spring Boot auto-configuration
- Dependency Injection (DI) and IoC container
- Bean lifecycle and scopes
- Application properties / YAML
- Spring Boot Actuator
- AOP basics
- Profiles

---

## ❓ Most Asked Questions

### Q1. What is Spring Boot and how is it different from Spring?

| Feature | Spring Framework | Spring Boot |
|---------|-----------------|-------------|
| Configuration | Extensive XML / Java config | Auto-configuration — convention over config |
| Server | Must configure external server | Embedded Tomcat/Jetty/Undertow built-in |
| Dependency management | Manual version management | Starter POMs handle versions |
| Bootstrapping | Lots of boilerplate | `@SpringBootApplication` + `main()` |
| Production readiness | Manual setup | Actuator built-in |

```java
// Minimal Spring Boot application
@SpringBootApplication  // = @Configuration + @EnableAutoConfiguration + @ComponentScan
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```

> Spring Boot **doesn't replace** Spring — it wraps it with sensible defaults and eliminates boilerplate.

---

### 🎯 How to Explain in Interview

"Spring Boot is essentially Spring on steroids - it takes all the power of the Spring framework and wraps it in a layer of convenience. The main difference is in the philosophy: Spring requires you to configure everything manually with XML or Java config, while Spring Boot follows convention over configuration with auto-configuration. For example, in traditional Spring, I'd need to configure a datasource bean, transaction manager, and web server manually. With Spring Boot, I just add the right starter dependency and it automatically configures everything for me. Spring Boot also comes with an embedded Tomcat server, so I can run my application with just a main method. It doesn't replace Spring - it makes Spring much easier to use for rapid development."

---

### Q2. What is Dependency Injection (DI)?

```java
// Without DI — tightly coupled
public class OrderService {
    private PaymentService paymentService = new PaymentServiceImpl();  // ❌ hard dependency
}

// With DI — loosely coupled (Spring manages creation)
@Service
public class OrderService {
    private final PaymentService paymentService;  // injected by Spring

    // Constructor injection (PREFERRED — immutable, testable)
    @Autowired  // optional when only one constructor
    public OrderService(PaymentService paymentService) {
        this.paymentService = paymentService;
    }
}

// Setter injection (use when dependency is optional)
@Service
public class ReportService {
    private NotificationService notificationService;

    @Autowired(required = false)
    public void setNotificationService(NotificationService svc) {
        this.notificationService = svc;
    }
}
```

---

### Q3. What are Spring Bean scopes?

```java
// Singleton — one instance per ApplicationContext (DEFAULT)
@Component
@Scope("singleton")  // default, can omit
public class ConfigService { }

// Prototype — new instance every time bean is requested
@Component
@Scope("prototype")
public class RequestProcessor { }

// Web-aware scopes (only in web applications)
@Component
@Scope(value = WebApplicationContext.SCOPE_REQUEST, proxyMode = ScopedProxyMode.TARGET_CLASS)
public class RequestContext { }   // new instance per HTTP request

@Component
@Scope(value = WebApplicationContext.SCOPE_SESSION, proxyMode = ScopedProxyMode.TARGET_CLASS)
public class UserSession { }     // new instance per HTTP session
```

| Scope | Instances | Lifecycle |
|-------|-----------|-----------|
| `singleton` | 1 per context | Lives as long as the context |
| `prototype` | 1 per injection/request | Spring creates, you manage destruction |
| `request` | 1 per HTTP request | Request lifecycle |
| `session` | 1 per HTTP session | Session lifecycle |

---

### 🎯 How to Explain in Interview

"Spring Bean scopes define how long a bean lives and how many instances exist. The default is singleton - Spring creates one instance for the entire application context. This is perfect for stateless services like controllers or business logic. Prototype creates a new instance every time it's requested - great for stateful objects like user-specific data. Then there are web-aware scopes: request creates a new bean for each HTTP request, and session creates one per user session. The key thing to remember is that singleton beans are shared, so they must be thread-safe, while prototype beans give each requester their own instance. Choosing the right scope is crucial for both performance and correctness."

---

### Q4. Explain `application.properties` vs `application.yml`

```properties
# application.properties
server.port=8080
spring.datasource.url=jdbc:mysql://localhost:3306/mydb
spring.datasource.username=root
spring.datasource.password=secret
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true
logging.level.com.example=DEBUG
```

```yaml
# application.yml — hierarchical, less repetition
server:
  port: 8080

spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb
    username: root
    password: secret
  jpa:
    hibernate:
      ddl-auto: update
    show-sql: true

logging:
  level:
    com.example: DEBUG
```

```java
// Reading custom properties
@Component
public class AppConfig {
    @Value("${app.name:MyApp}")     // with default value "MyApp"
    private String appName;

    @Value("${app.max-connections:10}")
    private int maxConnections;
}

// Better: Type-safe config with @ConfigurationProperties
@Configuration
@ConfigurationProperties(prefix = "app")
public class AppProperties {
    private String name;
    private int maxConnections;
    // getters and setters required
}
```

---

### 🎯 How to Explain in Interview

"Spring Boot gives me two main ways to handle configuration: properties and YAML. Properties files use simple key-value pairs, while YAML files are hierarchical and more readable for complex configurations. YAML is great because it reduces repetition - instead of writing spring.datasource.url, spring.datasource.username, I can nest them under spring.datasource. For reading configuration values, I can use @Value for individual properties, but the better approach is @ConfigurationProperties which gives me type-safe configuration binding. This way I get all the benefits of strong typing, validation, and IDE autocompletion. Plus, Spring Boot automatically loads different config files based on active profiles, so I can have dev, test, and prod configurations."

---

### Q5. What are Spring stereotypes (`@Component`, `@Service`, `@Repository`, `@Controller`)?

```java
// @Component — generic Spring-managed bean
@Component
public class UtilityHelper { }

// @Service — business logic layer (semantic — same as @Component)
@Service
public class UserService {
    public User findById(Long id) { return userRepository.findById(id).orElseThrow(); }
}

// @Repository — data access layer
// Translates persistence exceptions to Spring's DataAccessException hierarchy
@Repository
public class UserRepository extends JpaRepository<User, Long> { }

// @Controller — handles HTTP requests (returns views)
@Controller
public class HomeController {
    @GetMapping("/")
    public String home(Model model) {
        model.addAttribute("greeting", "Hello!");
        return "home";  // Thymeleaf template name
    }
}

// @RestController = @Controller + @ResponseBody (returns JSON/data directly)
@RestController
@RequestMapping("/api/users")
public class UserController {
    @GetMapping("/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        return ResponseEntity.ok(userService.findById(id));
    }
}
```

---

### Q6. What is a Spring Profile?

```yaml
# application.yml — defaults
app:
  name: MyApp

---
# application-dev.yml — auto-loaded when profile = dev
spring:
  datasource:
    url: jdbc:h2:mem:testdb    # in-memory DB for dev

---
# application-prod.yml — auto-loaded when profile = prod
spring:
  datasource:
    url: jdbc:mysql://prod-server:3306/mydb
```

```java
// Activate profile via env var or JVM arg
// java -jar app.jar --spring.profiles.active=prod
// Or: SPRING_PROFILES_ACTIVE=dev

// Profile-specific beans
@Configuration
public class DataSourceConfig {
    @Bean
    @Profile("dev")
    public DataSource h2DataSource() { return new EmbeddedDatabaseBuilder().build(); }

    @Bean
    @Profile("prod")
    public DataSource mysqlDataSource() {
        HikariDataSource ds = new HikariDataSource();
        ds.setJdbcUrl("jdbc:mysql://prod-server:3306/mydb");  // ✅ fixed
        return ds;
    }
}
```

---

### 🎯 How to Explain in Interview

"Spring Profiles are my way of dealing with different environments - dev, test, prod. They let me define different configurations for different environments and activate the right one at runtime. For example, in dev I might use an in-memory H2 database, while in prod I use a PostgreSQL cluster. I can create profile-specific configuration files like application-dev.yml and application-prod.yml, and Spring Boot automatically picks the right one based on the active profile. I can also create profile-specific beans using @Profile annotation. This makes environment-specific configuration clean and maintainable - no more commenting out database configurations or changing URLs manually."

---

### Q7. What is Spring AOP?

```java
// Aspect-Oriented Programming — cross-cutting concerns (logging, security, transactions)

@Aspect
@Component
public class LoggingAspect {

    // Pointcut — which methods to intercept
    @Pointcut("execution(* com.example.service.*.*(..))")
    public void serviceLayer() {}

    // Before advice — runs before method execution
    @Before("serviceLayer()")
    public void logBefore(JoinPoint jp) {
        System.out.println("Calling: " + jp.getSignature().getName());
    }

    // After returning — runs after successful execution
    @AfterReturning(pointcut = "serviceLayer()", returning = "result")
    public void logAfter(JoinPoint jp, Object result) {
        System.out.println("Returned: " + result);
    }

    // Around — wraps the method (most powerful)
    @Around("serviceLayer()")
    public Object timeMethod(ProceedingJoinPoint pjp) throws Throwable {
        long start = System.currentTimeMillis();
        Object result = pjp.proceed();         // call the actual method
        long elapsed = System.currentTimeMillis() - start;
        System.out.println(pjp.getSignature() + " took " + elapsed + "ms");
        return result;
    }
}
```

---

### Q8. What is Spring Boot Actuator?

```yaml
# Enable all actuator endpoints
management:
  endpoints:
    web:
      exposure:
        include: "*"
  endpoint:
    health:
      show-details: always
```

```java
// Key built-in endpoints
// GET /actuator/health       — application health (UP/DOWN)
// GET /actuator/info         — application info
// GET /actuator/metrics      — JVM metrics, HTTP stats
// GET /actuator/env          — environment properties
// GET /actuator/beans        — all Spring beans
// GET /actuator/mappings     — all @RequestMapping paths
// POST /actuator/loggers     — change log level at runtime

// Custom health indicator
@Component
public class DatabaseHealthIndicator implements HealthIndicator {
    @Override
    public Health health() {
        boolean dbUp = checkDatabaseConnection();
        if (dbUp) return Health.up().withDetail("db", "PostgreSQL").build();
        return Health.down().withDetail("error", "Cannot connect to DB").build();
    }
}

// Custom metric
@Component
public class OrderMetrics {
    private final MeterRegistry registry;
    private final Counter orderCounter;

    public OrderMetrics(MeterRegistry registry) {
        this.registry = registry;
        this.orderCounter = Counter.builder("orders.created")
            .description("Number of orders created")
            .register(registry);
    }

    public void orderCreated() { orderCounter.increment(); }
}
```

---

### 🎯 How to Explain in Interview

"Spring Boot Actuator is like having a doctor's office for my application - it gives me health checks and diagnostics out of the box. It provides endpoints like /actuator/health to check if my app is running, /actuator/metrics for performance data, and /actuator/info for application information. The best part is that I can create custom health indicators to check things like database connectivity or external service availability. I can also add custom metrics to track business-specific KPIs. Actuator turns my application into a production-ready service with monitoring capabilities, which is essential for DevOps and microservices. It's the difference between having an application that just runs and one that I can actually monitor and manage in production."

---

### Q9. What is `@Transactional`?

```java
@Service
public class TransferService {

    @Transactional  // auto creates/commits/rolls-back DB transaction
    public void transfer(Long fromId, Long toId, BigDecimal amount) {
        Account from = accountRepo.findById(fromId).orElseThrow();
        Account to   = accountRepo.findById(toId).orElseThrow();

        from.debit(amount);   // if this succeeds but next fails →
        to.credit(amount);    // rollback BOTH (consistency guaranteed)

        accountRepo.save(from);
        accountRepo.save(to);
    }

    // Fine-grained control
    @Transactional(
        rollbackFor    = {InsufficientFundsException.class},  // rollback on checked exception
        noRollbackFor  = {AuditLoggingException.class},       // don't rollback for this
        propagation    = Propagation.REQUIRED,                // join existing or create new (default)
        isolation      = Isolation.READ_COMMITTED,            // prevent dirty reads
        timeout        = 30,                                  // seconds
        readOnly       = true                                 // optimization hint for reads
    )
    public Account getAccount(Long id) { return accountRepo.findById(id).orElseThrow(); }
}
```

> **Key:** `@Transactional` works via Spring AOP proxy — it **won't work** if called from within the same class (self-invocation problem).

---

### 🎯 How to Explain in Interview

"@Transactional is Spring's way of handling database transactions automatically. Instead of manually managing transaction boundaries with begin, commit, and rollback, I just annotate my method and Spring handles everything. If the method completes successfully, Spring commits the transaction. If any exception occurs, it rolls back automatically. This is crucial for data consistency - in a money transfer, either both the debit and credit succeed, or neither does. I can also fine-tune the behavior with attributes like rollbackFor to specify which exceptions should trigger rollback, or propagation to control how transactions interact with existing ones. The key thing to remember is that @Transactional works through AOP proxies, so it won't work if I call a transactional method from within the same class."
