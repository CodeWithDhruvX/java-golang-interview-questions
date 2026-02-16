# Complete Spring Boot Developer Cheat Sheet

This comprehensive guide covers Spring Boot from **core keywords** to **advanced internals**, designed for interview preparation and architectural understanding.

---

## 1. Core Spring Boot Annotations & Setup

### 🔹 `@SpringBootApplication`
- **Definition**: Main entry point. Combines `@Configuration`, `@EnableAutoConfiguration`, and `@ComponentScan`.
- **Dependency**: `spring-boot-starter`
- **Syntax**:
```java
@SpringBootApplication
public class App {
    public static void main(String[] args) {
        SpringApplication.run(App.class, args);
    }
}
```

### 🔹 Stereotype Annotations (Component Scanning)
| Annotation | Layer | Description |
|---|---|---|
| `@Component` | Generic | Base annotation for any Spring-managed component. |
| `@Service` | Service | Business logic layer. |
| `@Repository` | Persistence | Data access layer (enables exception translation). |
| `@Controller` | Web | MVC Controller (returns views). |
| `@RestController` | Web | REST Controller (`@Controller` + `@ResponseBody`). Returns JSON/XML. |

### 🔹 Dependency Injection (DI)
- **`@Autowired`**: Injects a bean (Field, Setter, or Constructor).
  - *Best Practice*: Use **Constructor Injection** for immutability and testing.
- **`@Qualifier("beanName")`**: Resolves ambiguity when multiple beans of same type exist.
- **`@Primary`**: Marks a bean as the default candidate for injection.

### 🔹 Configuration
- **`@Configuration`**: Marks a class as a source of bean definitions.
- **`@Bean`**: Manually declares a bean method.
- **`@Value("${prop.key}")`**: Injects property values from `application.properties`.
- **`@ConfigurationProperties(prefix="app")`**: Binds a group of properties to a POJO (Type-safe configuration).

---

## 2. Auto-Configuration Internals (Advanced)

### 🔹 How It Works
1. **Bootstrap**: `SpringApplication.run()` starts the context.
2. **Discovery**: Reads `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.
3. **Loading**: Loads configuration classes.
4. **Filtering**: Applies **Conditional Annotations** to decide which beans to register.

### 🔹 Key Conditional Annotations
- **`@ConditionalOnClass`**: Loads only if a specific class is on the classpath.
- **`@ConditionalOnMissingBean`**: Loads only if the user hasn't defined their own bean of this type (prevents duplicates).
- **`@ConditionalOnProperty`**: Loads based on a configuration property value.
- **`@ConditionalOnWebApplication`**: Loads only if running in a web context.

---

## 3. Bean Lifecycle & Container Internals

### 🔹 Lifecycle Flow
1. **Instantiation**: Constructor called.
2. **Populate Properties**: DI happens.
3. **`Aware` Interfaces**: ApplicationContextAware, BeanNameAware.
4. **`BeanPostProcessor` (Before)**: Custom modification.
5. **Initialization**: `@PostConstruct`, `InitializingBean.afterPropertiesSet()`.
6. **`BeanPostProcessor` (After)**: AOP proxies applied here.
7. **Ready for Use**.
8. **Destruction**: `@PreDestroy`, `DisposableBean`.

### 🔹 Bean Scope
- **`singleton`** (Default): One instance per context.
- **`prototype`**: New instance per request.
- **`request` / `session`**: Web-specific scopes.

---

## 4. REST & Web (Spring MVC)

### 🔹 Annotations
- **`@RequestMapping`**: Base path mapping.
- **`@GetMapping`, `@PostMapping`, `@PutMapping`, `@DeleteMapping`**: Method-specific shortcuts.
- **`@RequestBody`**: Deserializes JSON to Java Object.
- **`@PathVariable`**: variable from URI path (`/users/{id}`).
- **`@RequestParam`**: Query parameter (`/users?id=1`).
- **`@ResponseStatus`**: Sets HTTP status code usually on Exceptions.

### 🔹 Request Flow (DispatcherServlet)
1. **Request** hits `DispatcherServlet`.
2. **`HandlerMapping`** finds the correct Controller method.
3. **`HandlerAdapter`** invokes the method.
4. **Controller** processes request and returns response.
5. **`ViewResolver`** (MVC) or MessageConverter (REST) formats output.

---

## 5. Data Access (Spring Data JPA)

### 🔹 Core JPA Annotations
- **`@Entity`**: Marks class as DB table.
- **`@Id`, `@GeneratedValue`**: Primary Key.
- **`@Table(name="...")`**: Custom table name.
- **`@Column`**: Column details.
- **`@OneToMany`, `@ManyToOne`**: Relationships.
- **`@Transactional`**: Check Transaction section.

### 🔹 Repository Interfaces
- **`JpaRepository`**: Standard CRUD + Paging + Sorting.
- **`@Query`**: Custom JPQL or Native SQL queries.

### 🔹 Fetch Strategies
- **EAGER**: Loads related entities immediately (Default for `@ManyToOne`).
- **LAZY**: Loads on demand (Default for `@OneToMany`). *Use LAZY to specific N+1 problems.*

### 🔹 The N+1 Problem
- **Cause**: Fetching a list of entities (1 query), then iterating to access a lazy relation (N queries).
- **Fix**: Use `JOIN FETCH` has JPQL or `@EntityGraph`.

---

## 6. Transaction Management

### 🔹 `@Transactional`
- Works via **AOP Proxies**.
- **Important**: Self-invocation (calling `@Transactional` method from same class) **ignores** transaction because proxy is bypassed.

### 🔹 Propagation Types
- **`REQUIRED`** (Default): Joins existing transaction or creates new one.
- **`REQUIRES_NEW`**: Suspends existing, creates connection.
- **`SUPPORTS`**: Runs in transaction if exists, else non-transactional.
- **`MANDATORY`**: Throws exception if no active transaction.

### 🔹 Isolation Levels
- **`READ_COMMITTED`**: Prevents dirty reads.
- **`REPEATABLE_READ`**: Prevents dirty & non-repeatable reads.
- **`SERIALIZABLE`**: Strict, prevents phantom reads.

---

## 7. Spring Security Internals

### 🔹 Security Filter Chain by default
Request → `DelegatingFilterProxy` → `FilterChainProxy` → Security Filters.

### 🔹 Core Components
- **`AuthenticationManager`**: Coordinates authentication.
- **`UserDetailsService`**: Loads user data from DB.
- **`PasswordEncoder`**: Hashes passwords (BCrypt).
- **`SecurityContextHolder`**: Stores details of authenticated user (ThreadLocal).

### 🔹 JWT Flow
1. **Filter** intercepts request.
2. Extracts Token from Header.
3. Validates Token.
4. Creates `Authentication` object.
5. Sets it in `SecurityContextHolder`.

---

## 8. Aspect-Oriented Programming (AOP)

### 🔹 Concepts
- **Aspect**: The cross-cutting concern (Logging, Security).
- **Advice**: Action taken (`@Before`, `@After`, `@Around`).
- **Pointcut**: Where to apply logic (Expression).

### 🔹 Proxy Mechanism
- **JDK Dynamic Proxy**: Used if class implements Interface.
- **CGLIB**: Used if class does not implement Interface (Class-based proxying).

---

## 9. Testing

- **`@SpringBootTest`**: Loads complete application context (Integration Test).
- **`@WebMvcTest`**: Loads only Web layer (Controllers), mocks Service layer.
- **`@DataJpaTest`**: Loads only JPA components, uses embedded DB.
- **`@MockBean`**: Mocks a Spring Bean in the context.

---

## 10. Microservices & Cloud

- **Spring Cloud Config**: Centralized configuration.
- **Eureka / Consul**: Service Discovery.
- **Spring Cloud Gateway**: API Gateway.
- **Resilience4j**: Circuit Breaking (`@CircuitBreaker`), Rate Limiting.
- **OpenFeign**: Declarative REST Client.

---

## 11. Actuator & Observability

- **Dependency**: `spring-boot-starter-actuator`
- **Endpoints**:
  - `/actuator/health`: App health status.
  - `/actuator/metrics`: JVM, DB, HTTP metrics.
  - `/actuator/info`: App info.
- **Metrics**: Micrometer (bridge to Prometheus/Grafana).

---

## 12. Common Interview Questions

1. **Difference between `@Component`, `@Service`, `@Repository`?**
   - Use `@Repository` for DB exception translation. `@Service` documents business logic. `@Component` is generic.
2. **What is Auto-Configuration?**
   - Spring Boot automatically configures beans based on classpath dependencies using `@Conditional...`.
3. **How does `@Transactional` work?**
   - Uses AOP proxy to open/commit/rollback JDBC connection.
4. **Why is Constructor Injection better?**
   - Ensures required dependencies are present (immutability), easier to test (no reflection needed).
5. **Difference between Filter and Interceptor?**
---

## 13. Microservices & Cloud Patterns

### 🔹 Core Components
- **Config Server**: Centralized configuration management.
- **Service Discovery (Eureka/Consul)**: Dynamic service registration.
- **API Gateway (Spring Cloud Gateway)**: Routing, filtering, security.
- **Circuit Breaker (Resilience4j)**: Fault tolerance (`@CircuitBreaker`).
- **Distributed Tracing (Sleuth/Zipkin)**: Request tracking across services.

### 🔹 OpenFeign
Client-side load balancing and declarative REST client.
```java
@FeignClient(name = "user-service")
public interface UserClient {
    @GetMapping("/users/{id}")
    User getUser(@PathVariable("id") Long id);
}
```

---

## 14. Messaging (Kafka / RabbitMQ)

### 🔹 Kafka
- **Dependency**: `spring-kafka`
- **Listener**:
```java
@KafkaListener(topics = "orders", groupId = "group_id")
public void consume(String message) {
    // process message
}
```

### 🔹 RabbitMQ
- **Dependency**: `spring-boot-starter-amqp`
- **Listener**:
```java
@RabbitListener(queues = "ordersQueue")
public void consume(String message) {
    // process message
}
```

---

## 15. Production Checklist & Optimization

### 🔹 Performance
1. **Connection Pooling**: Use HikariCP (Default). Tune `maximum-pool-size`.
2. **Lazy Loading**: Avoid EAGER fetching in JPA.
3. **Caching**: Enable `@EnableCaching` with Redis/Caffeine.
4. **Compression**: Enable GZIP (`server.compression.enabled=true`).

### 🔹 Security
1. **HTTPS**: Enforce TLS.
2. **Secrets**: Never commit secrets. Use Vault or Environment Variables.
3. **Actuator**: Secure sensitive endpoints (`/env`, `/heapdump`).

### 🔹 Deployment
- **Docker**: Use multi-stage builds.
- **Health Checks**: Configure Liveness & Readiness probes for K8s.
- **Graceful Shutdown**: `server.shutdown=graceful`.

---

## 16. Common "Senior" Interview Questions

1. **How to handle distributed transactions?**
   - Saga Pattern (Choreography/Orchestration), Two-Phase Commit (2PC - avoided due to locking).
2. **Difference between `@Controller` and `@RestController` in terms of flow?**
   - `@Controller` returns View name resolved by ViewResolver. `@RestController` writes directly to response body using MessageConverters.
3. **Explain Spring Boot Starters.**
   - Curated set of dependencies (POMs) that auto-configure the app (e.g., `starter-web` brings Tomcat, Jackson, MVC).
4. **What is the difference between `@Mock` and `@MockBean`?**
   - `@Mock`: Mockito native (Unit test). `@MockBean`: Spring Boot specific, replaces bean in ApplicationContext (Integration test).
5. **How to solve circular dependencies?**
   - Use `@Lazy`, Setter Injection, or refactor architecture (best).

