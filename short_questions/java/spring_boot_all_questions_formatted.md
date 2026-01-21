# Spring Boot Interview Questions & Answers (Summary Version)

> **Quick reference guide with concise explanations for Spring Boot interview questions 1-200**

---

## 🔹 Basics of Spring Boot (Questions 1-20)

**Q1: What is Spring Boot and how is it different from Spring?**
Spring Boot is an extension of the Spring Framework that simplifies setup and development. It provides "opinionated" defaults, auto-configuration, and an embedded server, whereas Spring requires manual configuration and deployment to an external server.

**Q2: What are the main features of Spring Boot?**
Auto-configuration, Starter dependencies, Embedded servers (Tomcat/Jetty), Actuator (metrics/health checks), CLI, and externalized configuration.

**Q3: What is the purpose of the `@SpringBootApplication` annotation?**
It is a convenience annotation that combines `@Configuration`, `@EnableAutoConfiguration`, and `@ComponentScan`. It marks the main class of the application.

**Q4: How does Spring Boot auto-configuration work?**
It scans the classpath for dependencies and automatically configures beans. For example, if `H2` is on the classpath, it configures an in-memory database.

**Q5: What is `spring-boot-starter`? Name a few starters.**
A set of convenient dependency descriptors. Examples: `spring-boot-starter-web` (REST), `spring-boot-starter-data-jpa` (DB), `spring-boot-starter-test` (Testing), `spring-boot-starter-security`.

**Q6: What is the use of `application.properties` or `application.yml`?**
Used to externalize configuration (DB URLs, ports, logging levels). `yml` allows hierarchical structure, while `properties` uses key-value pairs.

**Q7: How do you create a Spring Boot application from scratch?**
Use Spring Initializr (start.spring.io), Spring Tool Suite (STS), or CLI. Select dependencies, generate the project, and import into IDE.

**Q8: How does Spring Boot reduce boilerplate code?**
By using `starters` (simplifies build config) and `auto-configuration` (removes manual bean definitions and XML config).

**Q9: What is embedded Tomcat in Spring Boot?**
Spring Boot embeds a web server (Tomcat by default) inside the application JAR, allowing it to run as a standalone Java application (`java -jar`) without a separate server installation.

**Q10: What are Spring Initializr and its advantages?**
A web tool to bootstrap Spring Boot projects. It generates the project structure and build file (Maven/Gradle) with selected dependencies, saving setup time.

**Q11: Can you run Spring Boot without a web server?**
Yes. You can build non-web applications (e.g., batch jobs, CLI tools) by disabling the web environment (`spring.main.web-application-type=none`) or not including the web starter.

**Q12: How do you package a Spring Boot application (JAR vs WAR)?**
Default is JAR (embedded server). WAR is used for deploying to an external Servlet container (like WildFly or external Tomcat).

**Q13: What is the default port of Spring Boot web application?**
Port `8080`.

**Q14: How do you change the default port in Spring Boot?**
Set `server.port=8081` in `application.properties` or implement `WebServerFactoryCustomizer`.

**Q15: What is a CommandLineRunner in Spring Boot?**
functional interface with a `run()` method. Beans implementing this are executed once after the application context is loaded.

**Q16: What are actuators in Spring Boot?**
Production-ready features to monitor and manage the app. Exposes endpoints like `/health`, `/metrics`, `/info`, `/mappings`.

**Q17: How do you enable and use Spring Boot Actuator endpoints?**
Add `spring-boot-starter-actuator`. Enable endpoints via `management.endpoints.web.exposure.include=*` in properties.

**Q18: What is DevTools in Spring Boot and how does it help?**
A developer tool module that enables features like automatic restart (hot reload) when code changes and live reload of browser resources.

**Q19: How does Spring Boot support externalized configuration?**
It reads properties from multiple sources in a specific order: Command line args, Java System properties, OS env vars, `application.properties` (profile-specific then default).

**Q20: How can you run Spring Boot in different environments (dev, test, prod)?**
Use profiles (e.g., `application-dev.properties`). Activate via `-Dspring.profiles.active=dev` or env var `SPRING_PROFILES_ACTIVE`.

---

## 🔹 Annotations & Dependency Injection (Questions 21-40)

**Q21: What is the use of `@Component`, `@Service`, and `@Repository`?**
`@Component`: Generic stereotype. `@Service`: Business logic. `@Repository`: Persistence layer (enables exception translation). All are candidates for auto-detection.

**Q22: How does Spring Boot handle dependency injection?**
It uses the Spring IoC container. Dependencies are injected via Constructor (recommended), Setter, or Field (`@Autowired`).

**Q23: What is the difference between `@Autowired` and `@Qualifier`?**
`@Autowired`: Injects bean by type. `@Qualifier`: Used with `@Autowired` to specify a bean by name when multiple beans of the same type exist.

**Q24: What is the difference between `@Value` and `@ConfigurationProperties`?**
`@Value`: Injects single property value. Loosely typed.
`@ConfigurationProperties`: Binds a group of properties to a structured POJO. Type-safe and validated.

**Q25: How do you create custom annotations in Spring Boot?**
Define an interface with `@interface`. You can meta-annotate it with existing Spring annotations (e.g., `@Service`, `@Transactional`) to compose behavior.

**Q26: What is the use of `@Bean` and how is it different from `@Component`?**
`@Bean`: Used in `@Configuration` classes to explicitly declare a bean (method-level). Good for 3rd party classes.
`@Component`: Class-level annotation for auto-scanning.

**Q27: What is the `@Configuration` annotation?**
Indicates a class declares one or more `@Bean` methods and may be processed by the Spring container to generate bean definitions.

**Q28: Explain the use of `@EnableAutoConfiguration`.**
Tells Spring Boot to start adding beans based on classpath settings, other beans, and various property settings.

**Q29: What is the use of `@Profile` in Spring Boot?**
Restricts bean creation to specific profiles (environments). Example: `@Profile("dev")` bean is only created when `dev` profile is active.

**Q30: How to handle circular dependencies in Spring Boot?**
Use `@Lazy` injection (break the cycle by creating a proxy), Setter injection (instead of constructor), or refactor code.

**Q31: Can we use constructor injection in Spring Boot?**
Yes, it's the preferred method. It ensures immutability and makes testing easier. `@Autowired` is optional if there's only one constructor.

**Q32: What is a `@Scope` annotation and its types?**
Defines bean lifecycle. `singleton` (default, one per container), `prototype` (new per request), `request`, `session`, `application`.

**Q33: What are singleton and prototype beans?**
Singleton: One instance shared across the app. Prototype: New instance created every time the bean is requested.

**Q34: What is lazy initialization and how do you enable it?**
Beans are created only when requested, not at startup. Enable per bean with `@Lazy` or globally in properties (`spring.main.lazy-initialization=true`).

**Q35: How to define conditional beans? (`@ConditionalOnProperty`)**
Use annotations like `@ConditionalOnProperty`, `@ConditionalOnClass`, `@ConditionalOnMissingBean` to control creation based on config or presence of classes.

**Q36: What is a proxy bean and when is it created?**
A wrapper around the actual bean to add behavior (transaction management, AOP, lazy loading). Created when using AOP or CGLIB interfaces.

**Q37: What are factory methods in Spring beans?**
Methods annotated with `@Bean` inside a `@Configuration` class. They produce the bean instance managed by the container.

**Q38: What is the difference between `ApplicationContext` and `BeanFactory`?**
`BeanFactory`: Basic container, lazy loading.
`ApplicationContext`: Extends `BeanFactory`, adds AOP, event propagation, message resources, eager loading.

**Q39: How does Spring Boot manage bean lifecycle?**
Instantiate -> Populate properties -> `BeanNameAware`/`BeanFactoryAware` -> `PostProcessBeforeInitialization` -> `@PostConstruct` -> `PostProcessAfterInitialization` -> Use -> `@PreDestroy` -> Destroy.

**Q40: What is dependency injection and inversion of control in Spring Boot?**
IoC: Framework controls program flow. DI: Pattern to implement IoC where dependencies are provided to objects rather than created by them.

---

## 🔹 REST API with Spring Boot (Questions 41-60)

**Q41: How to create a REST API using Spring Boot?**
Use `spring-boot-starter-web`. Create a class with `@RestController`. Define endpoint methods with mapping annotations (`@GetMapping`).

**Q42: What are `@RestController` and `@Controller`?**
`@Controller`: Returns Views (JSP/Thymeleaf).
`@RestController`: Combines `@Controller` + `@ResponseBody`. Returns data (JSON/XML) directly.

**Q43: How to handle HTTP methods (GET, POST, PUT, DELETE) in Spring Boot?**
Use specific annotations: `@GetMapping`, `@PostMapping`, `@PutMapping`, `@DeleteMapping` mapping to controller methods.

**Q44: What is the purpose of `@RequestMapping`, `@GetMapping`, etc.?**
`@RequestMapping`: General mapping (can handle all methods). Helpers like `@GetMapping` are shortcuts for specific HTTP methods.

**Q45: What is the use of `@PathVariable` and `@RequestParam`?**
`@PathVariable`: Extract values from URI path (`/users/{id}`).
`@RequestParam`: Extract query parameters (`/users?id=1`).

**Q46: How to handle form-data and JSON in Spring Boot controllers?**
JSON: Use `@RequestBody` to map to POJO.
Form-data: Use `@RequestParam` or `@ModelAttribute` to bind form fields.

**Q47: How to return a custom HTTP status in Spring Boot?**
Return `ResponseEntity` object: `return ResponseEntity.status(HttpStatus.CREATED).body(data);` or use `@ResponseStatus` on exception/method.

**Q48: What is ResponseEntity and how to use it?**
A wrapper for the entire HTTP response: status code, headers, and body. Gives full control over the response.

**Q49: How to handle exceptions globally using `@ControllerAdvice`?**
Create a class with `@ControllerAdvice`. Define methods with `@ExceptionHandler` to catch specific exceptions and return custom error responses.

**Q50: What is the difference between synchronous and asynchronous API?**
Sync: Client waits for response. Thread blocked.
Async: Client may get "Accepted" response immediately, processing happens in background (`CompletableFuture`, `@Async`, WebFlux).

**Q51: What is CORS and how do you enable it in Spring Boot?**
Cross-Origin Resource Sharing. Enable via `@CrossOrigin` on controller/method or globally via `WebMvcConfigurer.addCorsMappings`.

**Q52: How to log incoming requests and responses in a REST controller?**
Use a Filter (implement `Filter` or `OncePerRequestFilter`) or HandlerInterceptor (`preHandle`, `afterCompletion`) to log details.

**Q53: What is content negotiation in Spring Boot?**
Serving different data formats (JSON, XML) based on `Accept` header. Configured via `ContentNegotiationConfigurer` and dependencies (Jackson XML).

**Q54: How to paginate and sort results in REST API?**
Use `Pageable` interface in Controller method. Pass `page`, `size`, `sort` params. Spring Data repositories support `Pageable` directly.

**Q55: What are DTOs and how are they used in Spring Boot?**
Data Transfer Objects. POJOs used to transport data between processes. Decouples internal Entity structure from API contract.

**Q56: How to map entities to DTOs? (ModelMapper/MapStruct)**
Manual set/get, or libraries like MapStruct (compile-time, fast) or ModelMapper (runtime). MapStruct is preferred for performance.

**Q57: How do you document a Spring Boot REST API using Swagger/OpenAPI?**
Add `springdoc-openapi-starter-webmvc-ui`. Access UI at `/swagger-ui.html`. Use annotations like `@Operation`, `@ApiResponse` for details.

**Q58: How to secure REST endpoints in Spring Boot?**
Use Spring Security. Configure `SecurityFilterChain`. Authenticate via Basic Auth, JWT, or OAuth2. Authorize via `@PreAuthorize`.

**Q59: How to test REST APIs using MockMvc or RestTemplate?**
`MockMvc`: Tests controller without starting server (Unit/Slice test).
`RestTemplate`/`TestRestTemplate`: Tests running server (Integration test).

**Q60: What is the difference between WebClient and RestTemplate?**
`RestTemplate`: Blocking, synchronous HTTP client (Maintenance mode).
`WebClient`: Non-blocking, reactive HTTP client (Spring WebFlux), supports sync and async.

---

## 🔹 JPA & Database (Questions 61-80)

**Q61: What is Spring Data JPA?**
Abstraction over JPA (Hibernate). Reduces boilerplate by providing repository interfaces (`JpaRepository`) with built-in CRUD and query methods.

**Q62: How do you define an entity in Spring Boot?**
Annotate a class with `@Entity`, specify primary key with `@Id`. Map fields to columns.

**Q63: What is the use of `@Entity`, `@Table`, and `@Id`?**
`@Entity`: Marks class as JPA entity. `@Table`: Customizes table name. `@Id`: Marks primary key field.

**Q64: Difference between CrudRepository, JpaRepository, and PagingAndSortingRepository?**
`CrudRepository`: Basic CRUD.
`PagingAndSortingRepository`: Adds pagination/sorting.
`JpaRepository`: Extends both + JPA specific methods (flush, batch delete).

**Q65: How do you write custom queries using `@Query` annotation?**
Use JPQL (`SELECT u FROM User u`) or Native SQL (`nativeQuery = true`) directly on repository methods.

**Q66: How does Spring Boot handle transactions?**
Uses `@Transactional`. Managed by TransactionManager. Wraps method in transaction; commit on success, rollback on RuntimeException.

**Q67: What is the use of `@Transactional` annotation?**
Defines the scope of a database transaction. configuring propagation, isolation, timeout, and rollback rules.

**Q68: How to enable lazy and eager loading in Spring Boot JPA?**
FetchType in relationships: `@OneToMany(fetch = FetchType.LAZY)` (default) or `EAGER`. Lazy loads only when accessed.

**Q69: How to use `application.properties` to configure the database?**
Set `spring.datasource.url`, `spring.datasource.username`, `spring.datasource.password`, `spring.jpa.hibernate.ddl-auto`.

**Q70: How to switch between in-memory DB (H2) and MySQL in Spring Boot?**
Change dependencies (remove h2, add mysql-connector) and update `spring.datasource.*` properties. Profiles can manage this switch (dev=h2, prod=mysql).

**Q71: What is the role of EntityManager?**
Interface to interact with persistence context. Used behind the scenes by Spring Data, or manually for complex dynamic queries.

**Q72: What is optimistic vs pessimistic locking in JPA?**
Optimistic: Uses version column (`@Version`). Checks on save.
Pessimistic: Locks DB row (`SELECT ... FOR UPDATE`). Waits for lock.

**Q73: What are native queries in Spring Boot JPA?**
SQL queries specific to the underlying DB (MySQL, Postgres). Defined with `@Query(value="...", nativeQuery=true)`.

**Q74: How to implement OneToOne, OneToMany, ManyToOne relationships?**
Use annotations: `@OneToOne`, `@OneToMany` (list/set), `@ManyToOne`. Define `mappedBy` on the owning side to avoid duplicate foreign keys.

**Q75: How to fetch child entities with parent using joins in JPA?**
Use JPQL `JOIN FETCH` (e.g., `SELECT p FROM Parent p JOIN FETCH p.children`) to load related entities in a single query (prevents N+1).

**Q76: What is CascadeType in JPA?**
Defines how state changes (persist, merge, remove) propagate from parent to child. e.g., `CascadeType.ALL` deletes children if parent is deleted.

**Q77: How do you handle schema generation in Spring Boot?**
Use `spring.jpa.hibernate.ddl-auto`: `create` (drop/create), `update` (alter schema), `validate` (check match), `none` (do nothing - production).

**Q78: What is Flyway or Liquibase and how do you use it with Spring Boot?**
Database migration tools. Manage schema versioning via scripts (`V1__init.sql`). Spring Boot auto-detects and runs them on startup.

**Q79: How to log SQL queries in Spring Boot?**
`spring.jpa.show-sql=true` (basic). For values, logging level of `org.hibernate.SQL` to `DEBUG` and `org.hibernate.type.descriptor.sql` to `TRACE`.

**Q80: How to test a Spring Boot repository with @DataJpaTest?**
Annotation for slice testing JPA components. Configures in-memory DB, scans `@Entity` classes and Repositories. Transactional (rolls back after test).

---

## 🔹 Security, Testing, and Deployment (Questions 81-100)

**Q81: How to secure a Spring Boot application with Spring Security?**
Add `spring-boot-starter-security`. Define `SecurityFilterChain` bean to configure auth rules (httpBasic, formLogin, antMatchers).

**Q82: How to configure Basic Authentication in Spring Boot?**
In `SecurityFilterChain`, call `http.httpBasic()`. Default behavior if starter is present without config.

**Q83: How to implement JWT-based authentication in Spring Boot?**
Disable default Session (`STATELESS`). Add Filter to validate JWT on requests. Add Endpoint to generate JWT on login.

**Q84: What is CSRF and how do you disable/enable it?**
Cross-Site Request Forgery. Enabled by default. Disable for stateless REST APIs via `http.csrf().disable()` (deprecated syntax in Boot 3) or `http.csrf(AbstractHttpConfigurer::disable)`.

**Q85: How to test Spring Security configurations?**
Use `@WebMvcTest` with secure controllers. Use `RequestPostProcessor` (userId/roles) or `@WithMockUser` annotation.

**Q86: How to use `@WithMockUser` in unit tests?**
Annotate test method: `@WithMockUser(username="admin", roles={"ADMIN"})`. Simulates authenticated context.

**Q87: What are unit tests, integration tests, and end-to-end tests in Spring Boot?**
Unit: Mock dependencies (Mockito). Fast.
Integration: Test interaction (DB, API). Slower.
E2E: Full system (Selenium/Cypress). Real user flow.

**Q88: What is `@SpringBootTest`?**
Bootstraps full application context. Used for integration tests. Can start real web environment (`webEnvironment=RANDOM_PORT`).

**Q89: How to mock beans using Mockito and MockBean?**
Mockito: Pure unit tests (`Mockito.mock()`).
`@MockBean`: Spring Boot annotation. Replaces bean in ApplicationContext with a Mock.

**Q90: How to test REST APIs with MockMvc or WebTestClient?**
`MockMvc`: `mockMvc.perform(get("/api")).andExpect(status().isOk())`.
`WebTestClient`: Reactive client, can also test standard MVC endpoints.

**Q91: What is the purpose of `@DataJpaTest` and `@WebMvcTest`?**
Slice tests. Load only relevant beans (JPA or Web layer) for faster execution than `@SpringBootTest`.

**Q92: How to deploy a Spring Boot app on a cloud (AWS/GCP/Azure)?**
Containerize (Docker), push to Registry (ECR/GCR). Deploy to orchestration service (EKS/ECS, Cloud Run, App Service) or PaaS (Elastic Beanstalk).

**Q93: How do you containerize a Spring Boot application using Docker?**
Create Dockerfile: `FROM openjdk`, `COPY jar`, `ENTRYPOINT ["java", "-jar", "app.jar"]`. Build image. Or use `mvn spring-boot:build-image` (Buildpacks).

**Q94: What is Spring Boot Admin?**
Writing a UI to manage and monitor Spring Boot applications. Clients register with Server. Displays Actuator data (logs, heap, threads).

**Q95: How to monitor and health-check a Spring Boot app in production?**
Use Actuator (`/health`). Integrate with monitoring tools (Prometheus, Grafana, Datadog).

**Q96: How to enable logging and tracing in Spring Boot?**
SLF4J/Logback default. config via `logback-spring.xml` or properties. Trace via Micrometer Tracing (formerly Sleuth) + Zipkin/Jaeger.

**Q97: How to handle application startup failures gracefully?**
Implement `FailureAnalyzer`. Catch startup exceptions and provide human-readable error messages.

**Q98: How to perform graceful shutdown in Spring Boot?**
Enable `server.shutdown=graceful`. Boot waits for active requests to complete (within timeout) before stopping server.

**Q99: How does Spring Boot integrate with message brokers (Kafka, RabbitMQ)?**
Starters (`spring-boot-starter-amqp`, `spring-kafka`). Auto-configures connection factories, templates (`KafkaTemplate`), and listeners (`@KafkaListener`).

**Q100: What is reactive programming in Spring Boot and how does it differ from traditional MVC?**
Non-blocking, event-driven (Reactor). Uses `Flux`/`Mono` instead of blocking collections. Scales better with high concurrency/low threads.

---

## 🔹 Core Spring Boot Concepts (Questions 101–120)

**Q101: How does Spring Boot handle backward compatibility across versions?**
Deprecation cycles. Release notes. `spring-boot-properties-migrator` helps identify removed/renamed properties on startup.

**Q102: What happens internally when a Spring Boot application starts?**
Starts listeners -> Prepares Environment -> Creates ApplicationContext -> Refreshes Context (loads beans) -> Calls Runners.

**Q103: What are banner files in Spring Boot, and how do you customize them?**
ASCII art shown at startup. Place `banner.txt` in `src/main/resources`. customizable with colors and version variables.

**Q104: What is lazy initialization in Spring Boot 2.2+?**
Configurable feature (`spring.main.lazy-initialization=true`) to create all beans lazily. Reduces startup time / memory.

**Q105: How do you disable specific auto-configurations in Spring Boot?**
`@SpringBootApplication(exclude = {DataSourceAutoConfiguration.class})` or via property `spring.autoconfigure.exclude`.

**Q106: How do conditional annotations like `@ConditionalOnClass` work?**
They check conditions (class present, bean missing, property set) before registering a bean definition. Core of auto-configuration.

**Q107: What is the difference between `spring.main.web-application-type=reactive` vs `servlet`?**
`servlet`: Classic MVC (Tomcat/Jetty).
`reactive`: WebFlux (Netty).
`none`: No web server.

**Q108: How to exclude a Spring Boot starter dependency?**
In Maven/Gradle: `<exclusion>` tag inside the dependency declaration. commonly used to exclude default logging (Logback) for Log4j2.

**Q109: Can you create a multi-module Spring Boot project?**
Yes. Parent POM manages dependencies. Modules (core, web, api) have their own POMs. Packaging builds composite application.

**Q110: How does Spring Boot handle memory management?**
Running on JVM. Dependent on Heap settings (`-Xmx`). Spring Boot itself has overhead for ClassLoading and Metaspace generally requires attention.

**Q111: How to set JVM arguments in Spring Boot applications?**
Pass them before the `-jar` command: `java -Xmx1024m -jar app.jar`.

**Q112: What is the role of the `META-INF/spring.factories` file?**
Before Boot 3: Used to register AutoConfiguration classes, EnvironmentPostProcessors, Listeners. (In Boot 3+, use `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`).

**Q113: How does Spring Boot support reactive programming?**
Via Spring WebFlux starter. Uses Project Reactor (`Flux`, `Mono`) and Netty server by default.

**Q114: How to enable debugging logs for specific packages?**
Property: `logging.level.com.example.mypackage=DEBUG`.

**Q115: What are some common pitfalls when using Spring Boot in production?**
Using default passwords, exposing Actuator publicly, ignoring connection pool tuning, not setting memory limits.

**Q116: What is the Spring Boot DevTools restart strategy?**
Uses two classloaders. Base (libs) and Restart (app code). Only reloads Restart classloader on change. Fast.

**Q117: What is the difference between `@Import` and `@ComponentScan`?**
`@ComponentScan`: Finds beans by package scanning (implicit).
`@Import`: Explicitly registers specific configuration classes or components.

**Q118: How to create a custom auto-configuration module?**
Create `@Configuration` class, add Conditional annotations, register in `spring.factories` (or imports file). Wrap in a starter.

**Q119: How does Spring Boot detect dependencies for conditional configuration?**
Uses `@ConditionalOnClass` checking if bytecode for a class exists on the runtime classpath.

**Q120: What is the role of `spring.factories` vs `spring.components`?**
`factories`: Framework extension points (AutoConfig).
`components`: Indexer file to speed up ComponentScan (avoiding full classpath scan).

---

## 🔹 Advanced Configuration & Profiles (Questions 121–140)

**Q121: What is the order of precedence for property sources in Spring Boot?**
DevTools > @TestPropertySource > CLI args > Inline JSON > System Props > OS Env > RandomValue > Profile config > Application config > Default.

**Q122: How do you use the `EnvironmentPostProcessor` interface?**
To customize the `Environment` before application context starts. Useful for loading custom config files or decrypting properties.

**Q123: What is a PropertySource and how do you load custom ones?**
Abstraction of key-value source. Load via `@PropertySource` annotation or `EnvironmentPostProcessor` for non-standard formats/locations.

**Q124: How can you externalize secrets (like passwords) securely in Spring Boot?**
Vault, AWS Secrets Manager, Config Server, or encrypted properties (Jasypt). Avoid plain text in git.

**Q125: How to implement hierarchical configuration using YAML in Spring Boot?**
YAML supports maps/lists naturally. Spring binds these to nested POJO structures via `@ConfigurationProperties`.

**Q126: How do you resolve configuration conflicts in multi-environment setups?**
Most specific profile wins. passed profile config overrides default `application.yml`.

**Q127: What is relaxed binding in Spring Boot?**
Flexible property naming. `my.property-name`, `my.propertyName`, `MY_PROPERTY_NAME` all map to `my.propertyName` in Java.

**Q128: How do you override application properties programmatically?**
`SpringApplication.setDefaultProperties(Map)` or implementation of `EnvironmentPostProcessor`.

**Q129: Can you use SpEL in Spring Boot configuration files?**
Not directly in `application.properties`/`yml`. But inside `@Value("#{...}")` in code, yes.

**Q130: How to enable configuration reloading without restarting the app?**
Spring Cloud Config (`@RefreshScope`), or DevTools (local).

---

## 🔹 REST API Internals (Questions 141–160)

**Q141: How does Spring Boot handle parameter binding in REST endpoints?**
Uses `HandlerMethodArgumentResolver`. Binds URI vars, query params, headers, and bodies to method arguments using reflection and converters.

**Q142: What is the difference between `HttpMessageConverter` and `ObjectMapper`?**
`ObjectMapper`: Jackson library class for JSON<->Object logic.
`HttpMessageConverter`: Spring interface using ObjectMapper to write HTTP responses/read requests.

**Q143: How do you control serialization and deserialization of JSON in Spring Boot?**
Customize Jackson `ObjectMapper` (bean), or use `@JsonInclude`, `@JsonProperty` on DTO fields.

**Q144: How do you return paginated responses with metadata?**
Return `Page<T>` object (Spring Data). Serializes into JSON with `content` list and `pageable` metadata (totalElements, totalPages).

**Q145: What is the difference between `@ModelAttribute` and `@RequestBody`?**
`@RequestBody`: Raw body content (JSON/XML).
`@ModelAttribute`: Form data (application/x-www-form-urlencoded) or URI params bound to object.

**Q146: How to validate request parameters in Spring Boot REST API?**
Use `@Validated` on Controller class and `@Min`, `@Max`, `@NotNull` on method arguments directly.

**Q147: How does Spring Boot support HATEOAS APIs?**
`spring-boot-starter-hateoas`. Use `RepresentationModel`, `EntityModel`, and `WebMvcLinkBuilder` to add relations links to responses.

**Q148: How do you apply global validation error handling?**
Catch `MethodArgumentNotValidException` in `@ControllerAdvice`. Format binding errors into custom error response.

**Q149: What is the use of `BindingResult` in controller methods?**
Argument following `@Valid` object. Holds validation errors. If present, Exception is NOT thrown; controller must handle errors manually.

**Q150: How do you apply rate limiting to Spring Boot REST endpoints?**
Bucket4j, Resilience4j, or Redis-based implementation (Lua script). Interceptor checks remaining tokens.

---

## 🔹 JPA, Queries & Data Layer (Questions 161–180)

**Q161: How to handle null-safe queries in Spring Data JPA?**
Java `Optional` return format or using Specification API / QueryDSL which handle null checks dynamically.

**Q162: What is Specification API and how is it used in Spring Boot?**
Programmatic query construction. Interface `Specification<T>`. Allows combining predicates (`AND` / `OR`) for dynamic filtering.

**Q163: How to implement soft deletes using JPA?**
Add `deleted` column. Use `@SQLDelete` (custom delete command) and `@Where` (filter active only) annotations on Entity.

**Q164: How do you audit entity changes in Spring Boot?**
`@EnableJpaAuditing`. Annotate entity fields with `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`, `@LastModifiedBy`.

**Q165: What are entity graphs and how do you use them in JPA?**
Performance optimization. Defines which attributes to fetch eagerly. Use `@EntityGraph` on Repository method used to override default fetch types.

**Q166: How do you use stored procedures in Spring Boot JPA?**
`@Procedure` annotation on repository method matching procedure name, or `@NamedStoredProcedureQuery` on Entity.

**Q167: What is the role of `JpaSpecificationExecutor`?**
Repository interface to be extended. Enables using Specifications methods (`findAll(Specification spec)`) in repository.

**Q168: How to manage complex joins using Criteria API?**
Type-safe way to build queries. Use `CriteriaBuilder` from `EntityManager`. Verbose but dynamic.

**Q169: What is the default transaction propagation behavior in Spring Boot?**
`REQUIRED`. Join existing transaction if open, otherwise create new.

**Q170: How can you implement multi-tenancy with Spring Boot JPA?**
Database per tenant (via `AbstractRoutingDataSource`) or Schema per tenant, or Discriminator Column (Hibernate filter).

---

## 🔹 Spring Security Deep Dive (Questions 181–200)

**Q181: How do you configure method-level security in Spring Boot?**
`@EnableGlobalMethodSecurity(prePostEnabled = true)` (Boot 2) or `@EnableMethodSecurity` (Boot 3). Use `@PreAuthorize("hasRole('ADMIN')")`.

**Q182: What is the difference between `@Secured`, `@PreAuthorize`, and `@RolesAllowed`?**
`@Secured`: Legacy, limited. `@RolesAllowed`: JSR-250 standard. `@PreAuthorize`: SpEL supported, powerful expression evaluation.

**Q183: How do you customize user details service in Spring Security?**
Implement `UserDetailsService` interface. Override `loadUserByUsername` to fetch user/roles from DB. Expose as Bean.

**Q184: How do you implement role hierarchy in Spring Boot Security?**
Define `RoleHierarchy` bean. "ADMIN > USER". Security expressions will respect inheritance.

**Q185: How to enable OAuth2 login in Spring Boot?**
`spring-boot-starter-oauth2-client`. Configure provider (Google/GitHub) details in `application.yml` (`client-id`, `client-secret`).

**Q186: How does Spring Security filter chain work in Boot apps?**
DelegatingFilterProxy -> FilterChainProxy -> List of SecurityFilters (Auth, CSRF, ExceptionTranslation, Authorization).

**Q187: What is the difference between `AuthenticationManager` and `SecurityContext`?**
`AuthManager`: Processes authentication request.
`SecurityContext`: Holds currently authenticated user details (Principal) for the session/request.

**Q188: How do you secure actuator endpoints in Spring Boot?**
Configure security rules in `SecurityFilterChain` to require ADMIN role for `/actuator/**` paths.

**Q189: How to implement custom JWT token generation and validation?**
Use JJWT or Nimbus library. Generate token on login. Create `OncePerRequestFilter` to parse header, validate sig, set Context.

**Q190: How do you secure REST endpoints using token-based authentication?**
Stateless config. Disable session. Validate token in filter. 401 if invalid.

**Q191: How to handle null-safe queries in Spring Data JPA?**
(Duplicate of Q161)

**Q192: What is Specification API and how is it used in Spring Boot?**
(Duplicate of Q162)

**Q193: How to implement soft deletes using JPA?**
(Duplicate of Q163)

**Q194: How do you audit entity changes in Spring Boot?**
(Duplicate of Q164)

**Q195: What are entity graphs and how do you use them in JPA?**
(Duplicate of Q165)

**Q196: How do you use stored procedures in Spring Boot JPA?**
(Duplicate of Q166)

**Q197: What is the role of `JpaSpecificationExecutor`?**
(Duplicate of Q167)

**Q198: How to manage complex joins using Criteria API?**
(Duplicate of Q168)

**Q199: What is the default transaction propagation behavior in Spring Boot?**
(Duplicate of Q169)

**Q200: How can you implement multi-tenancy with Spring Boot JPA?**
(Duplicate of Q170)

---

## 🔹 Testing Advanced (Questions 201-210)

**Q201: How to write parameterized tests in Spring Boot?**
Use JUnit 5 `@ParameterizedTest` with `@ValueSource`, `@CsvSource`, or `@MethodSource` to run the same test with different inputs.

**Q202: How do you test database interactions using TestContainers?**
Spin up real Docker containers (Postgres/MySQL) during tests using the `Testcontainers` library. cleaner and more reliable than H2.

**Q203: What is the purpose of `@Transactional` in Spring Boot tests?**
Rolls back the transaction after each test method, ensuring a clean state for the next test.

**Q204: How do you mock service layers in a controller test?**
Use `@MockBean` to inject a Mockito mock into the Spring ApplicationContext, replacing the real service bean.

**Q205: What are embedded databases and how are they used for testing?**
In-memory DBs like H2, HSQLDB. Fast, zero-config. Spring Boot can auto-configure them for `@DataJpaTest` to replace the real DB.

**Q206: How do you test exception handling in controllers?**
Use `MockMvc` to perform requests that trigger exceptions and assert on the standard error response or status code (e.g., `status().isBadRequest()`).

**Q207: How do you test configuration properties in Spring Boot?**
Use `@TestPropertySource` or `@SpringBootTest(properties = "...")` to inject test-specific values. Or direct binding test using `Binder`.

**Q208: What is the difference between `@MockBean` and `@SpyBean`?**
`@MockBean`: Complete mock (no real behavior).
`@SpyBean`: Wraps real bean. Calls real methods unless specific ones are mocked/stubbed.

**Q209: How to isolate integration tests using test profiles?**
Create `application-test.yml`. Annotate test class with `@ActiveProfiles("test")` to load that specific configuration.

**Q210: How to use Postman/Newman for Spring Boot API testing?**
Export Postman Collection. Run via Newman (CLI) in CI/CD pipeline to verify running API endpoints (E2E testing).

---

## 🔹 Actuator, Metrics & Monitoring (Questions 211-220)

**Q211: What is the `/actuator/health` endpoint and how to customize it?**
Shows app health (UP/DOWN). Customize by implementing `HealthIndicator` interface to check specific dependencies (DB, Kafka).

**Q212: How do you secure actuator endpoints?**
Use Spring Security. Restrict access to specific roles (ADMIN) or IP ranges. `requestMatchers(EndpointRequest.toAnyEndpoint()).hasRole("ADMIN")`.

**Q213: How do you integrate Prometheus and Grafana with Spring Boot?**
Add `micrometer-registry-prometheus`. Exposes `/actuator/prometheus`. Configure Prometheus to scrape this endpoint. Visualize in Grafana.

**Q214: What is Micrometer in Spring Boot?**
Vendor-neutral metrics facade. Allows instrumenting code once and exporting to different monitoring systems (Prometheus, Datadog, NewRelic).

**Q215: How do you define custom metrics using Micrometer?**
Inject `MeterRegistry`. Use `Counter`, `Gauge`, `Timer`, or `DistributionSummary` to record custom events.

**Q216: How to add tags and dimensions to metrics?**
Use `.tag("key", "value")` builder methods on metrics. Useful for slicing data (e.g., error counts by "region").

**Q217: How to monitor application memory and CPU usage using Actuator?**
Enabled by default in `/actuator/metrics`. Query `/actuator/metrics/jvm.memory.used` or `process.cpu.usage`.

**Q218: What is the difference between `@Timed`, `@Metered`, and `@Gauge` annotations?**
`@Timed`: Measures execution time and frequency.
`@Counted` (Metered): Counts invocations.
`@Gauge`: Returns instantaneous value (e.g., list size).

**Q219: How to enable distributed tracing using Spring Cloud Sleuth?**
(Now Micrometer Tracing in Boot 3). Adds Trace ID / Span ID to logs. Exports traces to Zipkin/Jaeger.

**Q220: How to monitor thread pool metrics in Spring Boot?**
Actuator exposes `executor` metrics if `ThreadPoolTaskExecutor` is used. Check `executor.active`, `executor.queued`.

---

## 🔹 Reactive, Kafka & Messaging (Questions 221-230)

**Q221: What is Spring WebFlux and how does it differ from Spring MVC?**
Reactive web framework. Non-blocking, event-loop based (Netty). MVC is blocking, thread-per-request (Tomcat).

**Q222: What are Mono and Flux in WebFlux?**
`Mono<T>`: 0 or 1 item.
`Flux<T>`: 0 to N items (stream).
Core Publishers in Project Reactor.

**Q223: How do you implement reactive REST endpoints in Spring Boot?**
Return `Mono` or `Flux` from `@RestController` methods. Spring handles subscription and non-blocking response.

**Q224: How to handle backpressure in WebFlux?**
Reactive Streams specification. Subscribers signal how much data they can handle (`Subscription.request(n)`). Operators like `limitRate` help.

**Q225: What is Server-Sent Events and how to implement it in Spring Boot?**
One-way streaming from server to client. Return `Flux<ServerSentEvent<T>>` or produce `text/event-stream`.

**Q226: How to connect Spring Boot with Kafka for message publishing?**
`spring-kafka` starter. Use `KafkaTemplate.send(topic, data)`. Configure bootstrap servers in yaml.

**Q227: How to consume Kafka messages in Spring Boot?**
Annotate method with `@KafkaListener(topics = "my-topic", groupId = "my-group")`.

**Q228: How to handle Kafka error handling and retries?**
Configure `DefaultErrorHandler` with `FixedBackOff` or `DeadLetterPublishingRecoverer` in the container factory.

**Q229: How to test Kafka consumers in Spring Boot?**
Use `EmbeddedKafka` (`@EmbeddedKafka`). asserts messages received or sent to topics in an in-memory broker.

**Q230: How does Spring Boot support RabbitMQ integration?**
`spring-boot-starter-amqp`. Uses `RabbitTemplate` for sending and `@RabbitListener` for receiving.

---

## 🔹 Deployment, Docker & CI/CD (Questions 231-240)

**Q231: How do you build a Spring Boot Docker image?**
`docker build` with Dockerfile (COPY jar). Or Cloud Native Buildpacks via `mvn spring-boot:build-image` (no Dockerfile needed).

**Q232: What is the benefit of layering JARs in Spring Boot?**
Separates dependencies, resources, and app code into layers in the JAR. Docker can cache stable dependency layers, speeding up builds.

**Q233: How do you use Jib for Spring Boot container builds?**
Google's Maven/Gradle plugin. Builds Docker image directly to registry without local Docker daemon. Fast, layered.

**Q234: How to optimize Docker image size for Spring Boot apps?**
Use lightweight base image (`eclipse-temurin:17-jre-alpine`). Layered JARs. Multi-stage builds (build in Maven image, copy jar to JRE image).

**Q235: How to deploy Spring Boot to Kubernetes?**
Create Deployment and Service YAMLs. Define replicas, image, ports. Apply with `kubectl`.

**Q236: How to configure Spring Boot with environment variables in Docker?**
Map Docker ENV vars to Spring properties. `ENV SPRING_DATASOURCE_URL` in Dockerfile maps to `spring.datasource.url`.

**Q237: What is the role of `entrypoint` in Docker for Spring Boot?**
Defines the command to run the app. `ENTRYPOINT ["java", "-jar", "/app.jar"]`. Allows passing additional args.

**Q238: How do you implement rolling updates for Spring Boot services?**
Kubernetes handles this. Update image tag in Deployment. K8s replaces pods one by one ensuring zero downtime (if readiness probes enabled).

**Q239: How to set up a CI/CD pipeline for Spring Boot in GitHub Actions?**
Workflow file `.github/workflows/main.yml`. Steps: Checkout, Set up JDK, Run Tests (`mvn test`), Build Image, Push to Registry, Deploy.

**Q240: How do you handle configuration secrets in containerized Spring Boot deployments?**
K8s Secrets, Docker Secrets, or Vault. Inject as Env Vars or mount as files. Never bake into image.

---

## 🔹 Spring Boot Internals & Bootstrapping (Questions 241-250)

**Q241: What is the role of `SpringApplicationBuilder` in Spring Boot?**
Fluent API to build `SpringApplication`. Allows chaining methods to configure hierarchy, listeners, profiles before running.

**Q242: How to conditionally load beans based on environment variables?**
`@ConditionalOnProperty(name = "my.env", havingValue = "true")`. Or `@ConditionalOnExpression`.

**Q243: What is the role of `spring-boot-loader` in executable JARs?**
Library inside the JAR that knows how to load classes from nested JARs (dependencies). Handles the `Main-Class` execution.

**Q244: How to load external JARs at runtime in Spring Boot?**
`PropertiesLauncher` (loader). Or configure `loader.path` to point to external lib folder.

**Q245: How do you implement custom logging initialization in Spring Boot?**
Add `logback-spring.xml` in classpath. Spring Boot detects it and configures logging before ApplicationContext starts.

**Q246: What is a non-web Spring Boot application, and how do you build one?**
Disable web starter. Implement `CommandLineRunner`. Set `WebApplicationType.NONE`. Useful for cron jobs.

**Q247: How does Spring Boot optimize startup performance?**
Classpath indexer, Lazy initialization, identifying beans via `spring.factories` (avoiding scan), AOT (Boot 3).

**Q248: What is the difference between `CommandLineRunner` and `ApplicationRunner`?**
`CommandLineRunner`: Receives raw `String... args`.
`ApplicationRunner`: Receives `ApplicationArguments` object (parsed args with options/values).

**Q249: How to extend Spring Boot's application lifecycle?**
Implement `SpringApplicationRunListener`. Hook into events like `starting`, `environmentPrepared`, `contextLoaded`.

**Q250: What is `SpringBootExceptionReporter` used for?**
Callback interface used to support custom reporting of startup errors (e.g., printing failure analysis).

---

## 🔹 Configuration & Profiles (Questions 251-260)

**Q251: How to inject nested configuration properties using `@ConfigurationProperties`?**
Create inner static classes in the Properties POJO. Spring maps nested YAML structure `app.security.oauth` to the inner class fields.

**Q252: How do you configure multiple datasources in Spring Boot?**
Define multiple `DataSource` beans. Mark one `@Primary`. Create separate `EntityManagerFactory` and `TransactionManager` for each.

**Q253: What is the role of `@ConfigurationPropertiesScan`?**
Scans for classes annotated with `@ConfigurationProperties` and registers them as beans (avoids manual `@EnableConfigurationProperties`).

**Q254: How do you validate fields in a `@ConfigurationProperties` class?**
Add `@Validated` on the class and JSR-303 annotations (`@NotNull`, `@Min`) on fields. Startup fails if invalid.

**Q255: What are custom configuration namespaces in Spring Boot?**
Refers to defining your own prefixes in `application.yml` (e.g. `myapp.feature`).

**Q256: How to use `spring.profiles.group` for grouped profile activation?**
Define alias: `spring.profiles.group.prod=proddb,prodmq,prodmetrics`. Activating `prod` activates the group members.

**Q257: How do you create a fail-fast configuration mechanism?**
Validation on `@ConfigurationProperties`. Or `Assert.notNull` in `@PostConstruct` of a config bean.

**Q258: How to define system-wide default properties in Spring Boot?**
`SpringApplication.setDefaultProperties()`. Or `application-default.properties`.

**Q259: How do you merge configuration files in Spring Boot?**
Spring Boot automatically merges `application.yml` from different locations (classpath, file system). Also `spring.config.import` allows importing others.

**Q260: What is profile-specific YAML merging behavior?**
Profile-specific properties override default ones. List properties are replaced, not merged, unless configured otherwise.

---

## 🔹 RESTful API (Advanced) (Questions 261-270)

**Q261: How does Spring Boot support OpenAPI/Swagger documentation?**
`springdoc-openapi`. Scans controllers. Generates JSON at `/v3/api-docs` and UI at `/swagger-ui.html`.

**Q262: How do you set global CORS configuration for REST APIs?**
`WebMvcConfigurer` bean. Override `addCorsMappings`. `registry.addMapping("/**").allowedOrigins("*")`.

**Q263: How do you handle matrix variables in Spring Boot controllers?**
Enable `removeSemicolonContent=false` in `PathMatchConfigurer`. Use `@MatrixVariable` in controller method.

**Q264: How do you throttle API requests in Spring Boot?**
(Duplicate of Q150). Use Rate Limiting libraries (Bucket4j).

**Q265: What is the use of `@RequestPart` in REST controllers?**
Used for `multipart/form-data`. Can bind a part to JSON/POJO (metadata) and another to `MultipartFile` (binary).

**Q266: How to handle partial updates (PATCH) in Spring Boot?**
Use `@PatchMapping`. Manually check non-null fields in DTO and update entity. Or use JsonPatch (complex).

**Q267: What’s the difference between `@RequestParam` and `@PathVariable`?**
(Duplicate of Q45). PathVariable is part of URL hierarchy, RequestParam is query string.

**Q268: How do you enable HTTP/2 support in Spring Boot?**
`server.http2.enabled=true`. Requires SSL (HTTPS) usually.

**Q269: How to handle large file uploads in Spring Boot?**
Streaming. `servlet.multipart.max-file-size`. Use `InputStream` to read instead of loading into memory.

**Q270: How to add versioning to REST APIs?**
(Duplicate of Q111). URL, Header, Parameter.

---

## 🔹 Data Handling & Persistence (Questions 271-280)

**Q271: How do you implement pessimistic locking in JPA with Spring Boot?**
`@Lock(LockModeType.PESSIMISTIC_WRITE)` on Repository method.

**Q272: How to use entity listeners in JPA with Spring Boot?**
`@EntityListeners(MyListener.class)`. Listener class contains `@PrePersist`, `@PostUpdate` methods.

**Q273: How to perform batch inserts or updates with Spring Boot JPA?**
`spring.jpa.properties.hibernate.jdbc.batch_size=50`. Use `saveAll()`. Ensure Identity generation strategy doesn't disable batching.

**Q274: How do you track entity lifecycle events?**
(Duplicate of Q272/Q164). Auditing or EntityListeners.

**Q275: How do you implement DTO conversion in Spring Boot JPA?**
Query Projections (Interfaces), Constructor Expression within JPQL, or Mapper libraries.

**Q276: What is the difference between fetch joins and entity graphs?**
Both solve N+1. Fetch Join is hardcoded in Query. Entity Graph is declarative and can be reused/overlayed on multiple queries.

**Q277: How to customize naming strategies in Spring Boot JPA?**
`spring.jpa.hibernate.naming.physical-strategy`. Implement `PhysicalNamingStrategy` to convert camelCase to snake_case etc.

**Q278: What is the `hibernate.ddl-auto` property and its options?**
(Duplicate of Q77). create, update, validate, none.

**Q279: How to prevent N+1 queries in Spring Boot?**
(Duplicate of Q65/Q75). Join Fetch, Entity Graphs.

**Q280: How to log SQL queries with parameter values?**
(Duplicate of Q79). Set logging level for `org.hibernate.type`.

---

## 🔹 Spring Security Advanced (Questions 281-290)

**Q281: How do you secure WebSockets in a Spring Boot application?**
`@EnableWebSocketSecurity`. Extend `AbstractSecurityWebSocketMessageBrokerConfigurer`. Secure destination patterns.

**Q282: How to configure custom CORS in Spring Security?**
Pass a `CorsConfigurationSource` bean to `http.cors()`.

**Q283: What is CSRF protection and how do you enable/disable it in Spring Boot?**
(Duplicate of Q84).

**Q284: How to create a custom login success handler?**
Implement `AuthenticationSuccessHandler`. redirect to specific page or return JSON. Config in `formLogin().successHandler()`.

**Q285: What is the difference between stateless and stateful authentication?**
Stateful: Server keeps session (JSessionID). Stateless: Server checks token (JWT) every request, no server state.

**Q286: How do you implement MFA (Multi-Factor Authentication) in Spring Boot?**
Multi-step auth. Step 1: Password -> Generate TOTP/Email code -> Partial Auth. Step 2: Validate code -> Full Auth.

**Q287: How do you secure method invocations in service layers?**
(Duplicate of Q181). `@PreAuthorize`.

**Q288: How to implement API key-based authentication in Spring Boot?**
Custom Filter checking a header (`X-API-KEY`). Validate against DB/Config. Set Authentication in Context.

**Q289: What are the common pitfalls when using Spring Security?**
Circular dependencies (UserDetailsService vs PasswordEncoder), disabling CSRF incorrectly, Order of filter chain, mixing Session/Stateless strategies.

**Q290: How do you use OAuth2 scopes in Spring Boot security?**
In Resource Server: `hasAuthority('SCOPE_read')`. Scopes are mapped to GrantedAuthorities.

---

## 🔹 Testing, Mocking & Quality (Questions 291-300)

**Q291: What is the role of `@TestConfiguration` in Spring Boot?**
Defines additional beans or overrides for tests only. Not picked up by component scan automatically; import via `@Import`.

**Q292: How to run integration tests with a specific profile?**
`@ActiveProfiles("integration")`.

**Q293: How do you use `MockMvc` vs `WebTestClient`?**
(Duplicate of Q90).

**Q294: What is the difference between `@WebMvcTest` and `@SpringBootTest`?**
(Duplicate of Q137). Slice vs Full context.

**Q295: How do you mock configuration properties in tests?**
(Duplicate of Q207).

**Q296: How to write integration tests using H2 in-memory DB?**
Default behavior of `@DataJpaTest`. Or add H2 dependency and `@AutoConfigureTestDatabase`.

**Q297: How to test scheduled jobs in Spring Boot?**
Waitility library (wait for condition). Or manually trigger method annotated with `@Scheduled` if public.

**Q298: How do you test Spring Events in Spring Boot?**
`@RecordApplicationEvents`. Or mock the EventPublisher. Or capture events with a TestListener.

**Q299: How do you write parameterized integration tests?**
(Duplicate of Q201, applied to integration).

**Q300: What are TestContainers and how do you use them with Spring Boot?**
(Duplicate of Q202).

---

## 🔹 Core Concepts & Annotations (Questions 301–320)

**Q301: What is the use of `@ConditionalOnMissingBean` in Spring Boot?**
Registers a bean only if a bean of that type or name is NOT already contained in the BeanFactory. Crucial for library authors to provide defaults that users can override.

**Q302: What does `@SpringBootConfiguration` do under the hood?**
It is a specialized form of `@Configuration` that indicates the class provides configuration for a Spring Boot application. It allows finding the main config for tests.

**Q303: How does Spring Boot handle circular dependency issues?**
By default, it fails on startup. You can break it using `@Lazy` injection, Setter injection, or by enabling `spring.main.allow-circular-references=true` (not recommended).

**Q304: What is `@EnableAutoConfiguration`, and how does it work with `@SpringBootApplication`?**
It triggers the auto-configuration logic. It checks the classpath and registers beans that are likely needed (e.g., DataSource if H2 is present). `@SpringBootApplication` includes it.

**Q305: How do you use `@Lazy` initialization in Spring Boot?**
Annotate a `@Bean` or `@Component` with `@Lazy`. The bean is created only when first requested, not at startup. Global config: `spring.main.lazy-initialization=true`.

**Q306: What is the difference between `@Import` and `@ComponentScan`?**
`@ComponentScan`: Scans packages for stereotypes (`@Component`).
`@Import`: Explicitly loads Configuration classes or Components, even if they aren't in the scan path.

**Q307: How can you override auto-configured beans in Spring Boot?**
Define a bean with the same name/type in your own configuration. Since user config runs before auto-config, `@ConditionalOnMissingBean` in auto-config sees your bean and backs off.

**Q308: What is the purpose of `@DependsOn` annotation?**
Forces a specific initialization order. Bean A annotated with `@DependsOn("B")` will wait for Bean B to be created.

**Q309: What are factory beans in Spring Boot and when to use them?**
Beans implementing `FactoryBean<T>`. Used to create complex objects (like connection pools) where the instantiation logic is non-trivial.

**Q310: What are some common bean scope types in Spring Boot?**
`singleton` (default), `prototype`, `request`, `session`, `application`, `websocket`.

**Q311: How do you bind lists and maps from YAML to `@ConfigurationProperties`?**
(See Q125/251). Define `List<String>` or `Map<String, String>` fields in the POJO. Spring automatically maps structure.

**Q312: What is the order of precedence in Spring Boot property resolution?**
(Duplicate of Q121). Command line > Typesafe Config > System Props > OS Env > Config Files.

**Q313: How do you inject system environment variables into your Spring Boot config?**
Use `${ENV_VAR_NAME}` in `application.properties` or `@Value("${ENV_VAR_NAME}")`. Spring maps OS variables (uppercase/underscores) loosely to properties.

**Q314: How do you provide a fallback configuration in Spring Boot?**
Use `@ConditionalOnMissingBean` or profiles. Default properties can be set in `SpringApplication.setDefaultProperties`.

**Q315: How can you encrypt configuration properties in Spring Boot?**
Use Jasypt (`jasypt-spring-boot-starter`). Values are stored as `ENC(encryptedstring)`. Decrypted at runtime.

**Q316: What is the role of `PropertySourcesPlaceholderConfigurer`?**
It resolves `${...}` placeholders in bean definitions against the current Spring Environment and local properties.

**Q317: How do you resolve placeholders in Spring Boot programmatically?**
Inject `Environment` and call `env.resolvePlaceholders("${property}")`.

**Q318: How do you externalize secrets in Spring Boot in a secure way?**
(Duplicate of Q124). Vault, Secrets Manager, etc.

**Q319: How do you define immutable configuration beans?**
Use `@ConstructorBinding` (Boot 2.2+) on the `@ConfigurationProperties` class and make fields `final`.

**Q320: How can you log all loaded configuration properties during startup?**
Enable debug logging or use `ContextRefreshedEvent` listener to print `Environment` sources. Actuator `/env` endpoint also shows them.

---

## 🔹 Advanced Configuration & Property Handling (Questions 321–340)
*(Note: Some overlap with previous section)*

**Q321: How do you bind lists and maps from YAML to `@ConfigurationProperties`?**
(Duplicate of Q311).

**Q322: What is the order of precedence in Spring Boot property resolution?**
(Duplicate of Q121).

**Q323: How do you inject system environment variables into your Spring Boot config?**
(Duplicate of Q313).

**Q324: How do you provide a fallback configuration in Spring Boot?**
(Duplicate of Q314).

**Q325: How can you encrypt configuration properties in Spring Boot?**
(Duplicate of Q315).

**Q326: What is the role of `PropertySourcesPlaceholderConfigurer`?**
(Duplicate of Q316).

**Q327: How do you resolve placeholders in Spring Boot programmatically?**
(Duplicate of Q317).

**Q328: How do you externalize secrets in Spring Boot in a secure way?**
(Duplicate of Q318).

**Q329: How do you define immutable configuration beans?**
(Duplicate of Q319).

**Q330: How can you log all loaded configuration properties during startup?**
(Duplicate of Q320).

**(Wait, the unique questions list had some section header overlap issues in numbering, let's fix Q331-Q340 based on unique topics)**

**Q331: How do you use the `EnvironmentPostProcessor` interface?**
(Duplicate of Q122).

**Q332: What is a PropertySource and how do you load custom ones?**
(Duplicate of Q123).

**Q333: How can you externalize secrets (like passwords) securely in Spring Boot?**
(Duplicate of Q124).

**Q334: How to implement hierarchical configuration using YAML in Spring Boot?**
(Duplicate of Q125).

**Q335: How do you resolve configuration conflicts in multi-environment setups?**
(Duplicate of Q126).

**Q336: What is relaxed binding in Spring Boot?**
(Duplicate of Q127).

**Q337: How do you override application properties programmatically?**
(Duplicate of Q128).

**Q338: Can you use SpEL in Spring Boot configuration files?**
(Duplicate of Q129).

**Q339: How to enable configuration reloading without restarting the app?**
(Duplicate of Q130).

**Q340: What is profile-specific YAML merging behavior?**
(Duplicate of Q260).

---

## 🔹 REST APIs & Web Layer (Questions 341–360)

**Q341: How do you implement request rate limiting per IP in Spring Boot?**
(Duplicate of Q150). Use Bucket4j/Resilience4j filters keyed by `request.getRemoteAddr()`.

**Q342: What is the difference between `@ControllerAdvice` and `@ExceptionHandler`?**
`@ExceptionHandler` works locally in a Controller. `@ControllerAdvice` makes it global for all Controllers.

**Q343: What is the role of `@CrossOrigin` in Spring Boot?**
Enables CORS for that specific controller/method. Adds `Access-Control-Allow-Origin` headers.

**Q344: How do you implement role-based access control at controller level?**
Use `@PreAuthorize("hasRole('ADMIN')")`.

**Q345: How do you deal with file streaming (PDF, ZIP) in Spring Boot REST?**
Return `StreamingResponseBody` or `ResponseEntity<Resource>`. Set content-type header (`application/pdf`).

**Q346: How do you implement ETag support in REST responses?**
Use `ShallowEtagHeaderFilter`. It hashes the response body and sends ETag. Client sends `If-None-Match`.

**Q347: How to apply request/response compression in Spring Boot?**
`server.compression.enabled=true`. `server.compression.mime-types=text/html,application/json`.

**Q348: What are the options to handle request timeout in Spring controllers?**
`spring.mvc.async.request-timeout`. Or use `Callable` / `DeferredResult` / `WebFlux`.

**Q349: How do you enable HATEOAS in Spring Boot?**
(Duplicate of Q62/147).

**Q350: How do you use filters to intercept and modify incoming requests?**
Implement `Filter` or `OncePerRequestFilter`. Register as `@Component` or `FilterRegistrationBean`.

**Q351: How does Spring Boot handle parameter binding in REST endpoints?**
(Duplicate of Q141).

**Q352: What is the difference between `HttpMessageConverter` and `ObjectMapper`?**
(Duplicate of Q142).

**Q353: How do you control serialization and deserialization of JSON in Spring Boot?**
(Duplicate of Q143).

**Q354: How do you return paginated responses with metadata?**
(Duplicate of Q144).

**Q355: What is the difference between `@ModelAttribute` and `@RequestBody`?**
(Duplicate of Q145).

**Q356: How to validate request parameters in Spring Boot REST API?**
(Duplicate of Q146).

**Q357: How does Spring Boot support HATEOAS APIs?**
(Duplicate of Q147).

**Q358: How do you apply global validation error handling?**
(Duplicate of Q148).

**Q359: What is the use of `BindingResult` in controller methods?**
(Duplicate of Q149).

**Q360: How do you apply rate limiting to Spring Boot REST endpoints?**
(Duplicate of Q150).

---

## 🔹 Data & JPA Advanced (Questions 361–380)

**Q361: How do you use `@Query` with native SQL in Spring Data JPA?**
`@Query(value = "SELECT * FROM users", nativeQuery = true)`.

**Q362: What is the difference between `EntityManager.merge()` and `save()`?**
`save()` (Spring Data): Checks ID. If null/new, calls persist. If exists, calls merge.
`merge()`: Updates a detached entity.

**Q363: How do you implement cascading deletes in JPA?**
`cascade = CascadeType.REMOVE` on `@OneToMany`. Or `orphanRemoval = true`.

**Q364: How do you use database views in Spring Boot JPA?**
Map the view to an `@Entity` marked with `@Immutable`. Treat it like a readonly table.

**Q365: How do you map stored procedures using Spring Boot JPA?**
(Duplicate of Q166).

**Q366: How do you implement pagination using `Slice` and `Page`?**
`Page` returns total count (extra count query). `Slice` only returns data + "hasNext" flag (faster, no count query).

**Q367: How do you create and use custom repository implementations in Spring Data?**
Create interface `MyRepoCustom` and class `MyRepoImpl`. Spring Data automatically merges it with the standard JPA repository proxy.

**Q368: What’s the role of `JpaSpecificationExecutor` in Spring Data JPA?**
(Duplicate of Q167).

**Q369: How do you use `@SqlResultSetMapping` for native queries?**
Define mapping in Entity. Use name in `@NamedNativeQuery` or `createNativeQuery`. Maps SQL result to DTO/Entity.

**Q370: How do you use `@Converter` to transform custom types in JPA entities?**
Implement `AttributeConverter<X, Y>`. Annotate with `@Converter(autoApply=true)`. Converts Java Object <-> DB Column.

**Q371: How to handle null-safe queries in Spring Data JPA?**
(Duplicate of Q161).

**Q372: What is Specification API and how is it used in Spring Boot?**
(Duplicate of Q162).

**Q373: How to implement soft deletes using JPA?**
(Duplicate of Q163).

**Q374: How do you audit entity changes in Spring Boot?**
(Duplicate of Q164).

**Q375: What are entity graphs and how do you use them in JPA?**
(Duplicate of Q165).

**Q376: How do you use stored procedures in Spring Boot JPA?**
(Duplicate of Q166).

**Q377: What is the role of `JpaSpecificationExecutor`?**
(Duplicate of Q167).

**Q378: How to manage complex joins using Criteria API?**
(Duplicate of Q168).

**Q379: What is the default transaction propagation behavior in Spring Boot?**
(Duplicate of Q169).

**Q380: How can you implement multi-tenancy with Spring Boot JPA?**
(Duplicate of Q170).

---

## 🔹 Spring Security Deep Dive (Questions 381–400)

**Q381: How do you implement login throttling in Spring Security?**
Listen to `AuthenticationFailureBadCredentialsEvent`. Increment counter in Cache/DB for IP/User. Block after N attempts.

**Q382: What is the difference between stateless JWT authentication and session-based authentication?**
(Duplicate of Q285).

**Q383: How do you secure static resources like CSS or JS files?**
`http.authorizeHttpRequests().requestMatchers("/css/**", "/js/**").permitAll()`.

**Q384: How to protect APIs using OAuth2 with custom token validators?**
Configure `JwtDecoder`. Add custom validators (e.g. check audience, issuer, custom claims).

**Q385: What is the use of `SecurityContextHolder` in Spring Security?**
ThreadLocal storage for `SecurityContext`. Access current user via `SecurityContextHolder.getContext().getAuthentication()`.

**Q386: How do you configure `SecurityFilterChain` with multiple chains?**
Define multiple `SecurityFilterChain` beans with `@Order`. Use `securityMatcher` to define which chain applies to which URL pattern.

**Q387: What are antMatchers vs mvcMatchers in Spring Security configuration?**
`antMatchers`: Matches URL path.
`mvcMatchers`: Matches Spring MVC mappings (handles trailing slashes, etc.). More secure for MVC apps.

**Q388: How do you secure actuator endpoints conditionally?**
(Duplicate of Q188/212).

**Q389: How do you implement permission-based access (fine-grained auth)?**
Custom PermissionEvaluator. `@PreAuthorize("hasPermission(#id, 'Target', 'READ')")`.

**Q390: How to log user authentication/authorization events in Spring Security?**
Publish `AuthenticationSuccessEvent` / `AuthorizationFailureEvent`. Create Listeners to log them.

**Q391: How do you configure method-level security in Spring Boot?**
(Duplicate of Q181).

**Q392: What is the difference between `@Secured`, `@PreAuthorize`, and `@RolesAllowed`?**
(Duplicate of Q182).

**Q393: How do you customize user details service in Spring Security?**
(Duplicate of Q183).

**Q394: How do you implement role hierarchy in Spring Boot Security?**
(Duplicate of Q184).

**Q395: How to enable OAuth2 login in Spring Boot?**
(Duplicate of Q185).

**Q396: How does Spring Security filter chain work in Boot apps?**
(Duplicate of Q186).

**Q397: What is the difference between `AuthenticationManager` and `SecurityContext`?**
(Duplicate of Q187).

**Q398: How do you secure actuator endpoints in Spring Boot?**
(Duplicate of Q188).

**Q399: How to implement custom JWT token generation and validation?**
(Duplicate of Q189).

**Q400: How do you secure REST endpoints using token-based authentication?**
(Duplicate of Q190).

---

## 🔹 Spring Boot CLI, DevTools & Utilities (Questions 401–420)

**Q401: What is Spring Boot CLI, and when should you use it?**
Command line tool to run Groovy scripts as Spring Boot apps. Good for rapid prototyping.

**Q402: How do you install and run Groovy scripts with Spring Boot CLI?**
Install CLI (SDKMAN/Homebrew). Run `spring run app.groovy`. CLI grabs dependencies (`@Grab`) automatically.

**Q403: What are live reload features provided by Spring Boot DevTools?**
Starts a LiveReload server. Browser extensions trigger page refresh when resources change.

**Q404: How do you exclude DevTools from production environments?**
It is automatically disabled when running as a fully packaged JAR/WAR. Maven scope `<optional>true</optional>`.

**Q405: How do you auto-open the browser on app startup using DevTools?**
Not a native feature, but can be scripted or done via `EventListener` on `ApplicationReadyEvent`.

**Q406: How do you use Spring Initializr CLI for generating projects?**
`spring init --dependencies=web,data-jpa my-project`.

**Q407: How can you create custom starters for internal libraries?**
Create project with AutoConfiguration class as per `spring.factories` (or imports). Naming convention: `my-lib-spring-boot-starter`.

**Q408: What is the purpose of `.spring-boot-devtools.properties`?**
Global configuration for DevTools (e.g., global trigger file for restart). Located in User Home.

**Q409: How do you define global logging formats via CLI or YAML?**
`logging.pattern.console=%d{HH:mm:ss} - %msg%n`.

**Q410: How to quickly prototype apps using CLI and embedded templates?**
CLI supports Thymeleaf/Groovy templates. Just put files in templates folder and run.

**Q411: What are the internals of Spring Boot’s startup process?**
(Duplicate of Q102).

**Q412: How does Spring Boot reduce boilerplate configuration?**
(Duplicate of Q8).

**Q413: What is a starter dependency and how does it work?**
(Duplicate of Q5).

**Q414: How does Spring Boot decide which auto-configurations to apply?**
(Duplicate of Q4/106).

**Q415: What is the significance of `spring.factories` in Spring Boot?**
(Duplicate of Q112).

**Q416: How do you programmatically disable specific auto-configurations?**
(Duplicate of Q105).

**Q417: How do you create a custom banner for your Spring Boot application?**
(Duplicate of Q103).

**Q418: What is the role of `@EnableConfigurationProperties`?**
Enables support for `@ConfigurationProperties` beans.

**Q419: How is SpringApplication.run() internally implemented?**
(Duplicate of Q102).

**Q420: What is the purpose of `SpringBootExceptionReporter`?**
(Duplicate of Q250).

---

## 🔹 Spring Boot with NoSQL (Questions 421–440)

**Q421: How do you configure MongoDB with Spring Boot?**
`spring-boot-starter-data-mongodb`. Config properties `spring.data.mongodb.uri`.

**Q422: How does Spring Data MongoDB support query methods?**
Same as JPA. `findByLastname(String lastname)`. Generates MongoDB JSON queries.

**Q423: How do you define compound indexes in MongoDB using Spring annotations?**
`@CompoundIndexes({ @CompoundIndex(def = "{'name':1, 'age':-1}") })` on document class.

**Q424: What is the role of `ReactiveMongoRepository`?**
Non-blocking repository based on Reactor (`Mono`/`Flux`). Used with WebFlux.

**Q425: How do you work with Cassandra in Spring Boot?**
`spring-boot-starter-data-cassandra`. Define `@Table` entities. `CassandraRepository`.

**Q426: How do you implement auditing in MongoDB?**
`@EnableMongoAuditing`. `@CreatedDate`, `@LastModifiedDate` on Document fields.

**Q427: How to use Redis as a primary data store with Spring Boot?**
`spring-boot-starter-data-redis`. Use `RedisTemplate` or `RedisRepository` (`@RedisHash`).

**Q428: How do you handle key expiration in Redis?**
`@RedisHash(timeToLive = 60)`. Or `redisTemplate.expire(key, timeout, unit)`.

**Q429: How do you store objects in Redis using Spring Boot?**
Serialize to JSON/Binary. Configure `RedisTemplate` with `Jackson2JsonRedisSerializer`.

**Q430: How do you implement caching using MongoDB or Redis?**
`@EnableCaching`. `@Cacheable`. Configure RedisCacheManager. MongoDB usually not used for temporary cache as much as Redis.

**Q431: How do you enable and customize access logs in Spring Boot?**
`server.tomcat.accesslog.enabled=true`.

**Q432: How do you configure JSON logging format in Spring Boot?**
Use logging framework plugins (Logstash encoder for Logback). `logging.config=classpath:logback-json.xml`.

**Q433: How can you implement centralized logging with ELK stack in Spring Boot?**
Output logs to file/TCP (Logstash). Logstash forwards to Elasticsearch. View in Kibana.

**Q434: How do you create custom metrics using Micrometer?**
(Duplicate of Q215).

**Q435: How do you expose custom health indicators in Spring Boot?**
(Duplicate of Q211).

**Q436: What are some built-in health checks provided by Spring Boot Actuator?**
DiskSpace, DataSource, Ping, Cassandra, RabbitMQ, Kafka.

**Q437: How do you monitor thread usage in a running Spring Boot app?**
(Duplicate of Q220).

**Q438: What is the use of `/actuator/metrics` endpoint?**
Lists available metrics names.

**Q439: How do you push metrics to a time-series database from Spring Boot?**
Add Micrometer registry dependency (e.g. `micrometer-registry-influx`). Auto-configured.

**Q440: How do you configure log levels dynamically at runtime?**
POST to `/actuator/loggers/{package}`.

---

## 🔹 Integration Patterns (Questions 441–460)

**Q441: What is Spring Integration DSL?**
Fluent Java API to define integration flows (Channels, Filters, Transformers).

**Q442: How do you build a file polling system using Spring Integration?**
`IntegrationFlows.from(Files.inboundAdapter(dir)).transform(...).handle(...)`.

**Q443: How do you integrate with SOAP services using Spring Boot?**
`spring-boot-starter-web-services`. Use `WebServiceTemplate`. Generate classes from WSDL (JAXB).

**Q444: How do you implement FTP file transfers with Spring Integration?**
Use `Ftp.inboundAdapter` or `Ftp.outboundAdapter`.

**Q445: How to trigger workflows using Spring Events?**
Publish event: `publisher.publishEvent(new MyEvent())`. Listener: `@EventListener` triggers flow.

**Q446: What is an `IntegrationFlow` in Spring Integration?**
Defines the pipeline. Message Source -> Channels/Handlers -> Sink.

**Q447: What are channels in Spring Integration, and how are they used?**
Pipes connecting components. `DirectChannel` (Point-to-Point), `PublishSubscribeChannel` (Broadcast), `QueueChannel` (Buffered).

**Q448: What is a gateway in Spring Integration?**
Interface that hides the messaging system from application code. API call -> Gateway -> Message Channel.

**Q449: How do you implement message transformation pipelines?**
`.transform(String::toUpperCase)` or `.transform(Transformers.toJson())`.

**Q450: How do you integrate with Apache Camel from Spring Boot?**
`camel-spring-boot-starter`. Define `RouteBuilder` beans (`from("direct:start").to("log:info")`).

**Q451: What is the difference between `@SpringBootTest` and `@WebMvcTest`?**
(Duplicate of Q137).

**Q452: How do you test REST controllers in Spring Boot?**
(Duplicate of Q136).

**Q453: How do you mock services using `@MockBean`?**
(Duplicate of Q89/204).

**Q454: What is the use of `TestRestTemplate` in Spring Boot?**
(Duplicate of Q59).

**Q455: How do you test JPA repositories effectively?**
(Duplicate of Q80).

**Q456: How do you use in-memory databases for unit tests?**
(Duplicate of Q206).

**Q457: What is `@DataJpaTest` used for?**
(Duplicate of Q91).

**Q458: How do you perform integration testing with embedded containers?**
(Duplicate of Q202).

**Q459: What is the purpose of `@TestConfiguration`?**
(Duplicate of Q291).

**Q460: How do you test exception scenarios in Spring Boot?**
(Duplicate of Q206).

---

## 🔹 Microservices & Spring Cloud (Questions 461–480)

**Q461: What is Spring Cloud Sleuth and how does it enhance traceability?**
(Duplicate of Q219). (Deprecated in 3.0, replaced by Micrometer Tracing).

**Q462: How does Spring Cloud Config work internally?**
Server maps Git/SVN repo to property environment. Client fetches config on startup via HTTP.

**Q463: What is `spring.cloud.bootstrap.enabled=true` used for?**
Enables legacy bootstrap context (bootstrap.yml) in newer Spring Cloud versions.

**Q464: How do you use Feign clients in Spring Boot?**
`@EnableFeignClients`. Interface with `@FeignClient`. Define methods mapping to HTTP endpoints. Declarative REST client.

**Q465: What is Spring Cloud Bus and how does it work with Kafka or RabbitMQ?**
Propagates state changes (like config refresh) to all instances via message broker.

**Q466: What is service discovery, and how does Eureka support it?**
Dynamic registry of service instances (IP/Port). Clients look up "user-service" to get actual address.

**Q467: What are ribbon clients, and how are they configured?**
Client-side load balancer (Maintenance mode -> Spring Cloud LoadBalancer). Distributes requests across discovered instances.

**Q468: How do you implement distributed tracing with Zipkin and Sleuth?**
Sleuth adds IDs. Zipkin server collects and visualizes trace data.

**Q469: How do you configure load balancing across microservices?**
`@LoadBalanced` on RestTemplate. Or Feign (built-in). Uses Round Robin by default.

**Q470: What is the purpose of Circuit Breakers in Spring Boot microservices?**
(Duplicate of Q120). Resilience4j.

**Q471: What are Spring Boot profiles and how do they work?**
(Duplicate of Q29).

**Q472: How do you activate multiple profiles simultaneously?**
`-Dspring.profiles.active=dev,db-local`.

**Q473: How do you deploy a Spring Boot app as a WAR file?**
Extend `SpringBootServletInitializer`. Change packing to `war`.

**Q474: What is the difference between embedded Tomcat and external Tomcat deployment?**
Embedded: Runs in JAR process. External: App is library in container process.

**Q475: How do you deploy Spring Boot apps on Heroku?**
Git push to Heroku remote. Procfile: `web: java -jar app.jar`.

**Q476: What is `spring.profiles.active` and where can you define it?**
(Duplicate of Q20).

**Q477: How do you implement zero-downtime deployment for Spring Boot apps?**
(Duplicate of Q238). Rolling updates.

**Q478: How do you enable graceful shutdown in Spring Boot?**
(Duplicate of Q98).

**Q479: How do you deploy Spring Boot apps to AWS Lambda?**
Use Spring Cloud Function + AWS Adapter. Function acts as handler.

**Q480: What is the role of `application.properties` vs `application.yml`?**
(Duplicate of Q6).

---

## 🔹 Advanced DevOps & CI/CD (Questions 481–500)

**Q481: How do you build Spring Boot apps as OCI-compliant containers using buildpacks?**
(Duplicate of Q231).

**Q482: How do you configure a multistage Dockerfile for a Spring Boot app?**
(Duplicate of Q234).

**Q483: How do you push Spring Boot metrics to Prometheus?**
(Duplicate of Q213).

**Q484: How to automate release versioning using Maven or Gradle plugins?**
Maven Release Plugin, or Semantic Release tools.

**Q485: What is the role of `spring-boot-maven-plugin` and `spring-boot-gradle-plugin`?**
Repackages JAR to be executable (adds manifest, layout). Allows `mvn spring-boot:run`.

**Q486: How to detect and fail builds on deprecated Spring APIs?**
Compiler flag `-Xlint:deprecation`. `failOnWarning` in build tool.

**Q487: How to dynamically reload configuration in Kubernetes with Spring Boot?**
Mount ConfigMap as volume. Watcher sidecar updates usage, or Spring Cloud Kubernetes Reload feature.

**Q488: What is the use of Spring Boot Admin in DevOps pipelines?**
Visual dashboard for health/status. Quick check if deployment is up.

**Q489: How do you expose Prometheus-compatible metrics with Micrometer?**
(Duplicate of Q213).

**Q490: How do you automate blue-green deployment strategies with Spring Boot?**
Router/Load Balancer switches traffic between Green (New) and Blue (Old) environments.

**Q491: How do you use `@PreAuthorize` and `@PostAuthorize` annotations?**
(Duplicate of Q181/391).

**Q492: How do you enable method-level security globally in Spring Boot?**
(Duplicate of Q181).

**Q493: What is the purpose of `SecurityContextPersistenceFilter`?**
Restores SecurityContext from Session before request, saves it back after request.

**Q494: How do you define custom authentication providers?**
Implement `AuthenticationProvider`. Register in `AuthenticationManagerBuilder`.

**Q495: How do you implement LDAP authentication in Spring Boot?**
`spring-boot-starter-data-ldap`. Configure `auth.ldapAuthentication()`.

**Q496: How do you use `BCryptPasswordEncoder` in a login system?**
Encode password on signup. Match raw password with hash on login. `passwordEncoder.matches()`.

**Q497: How do you secure REST endpoints using JWT in Spring Boot?**
(Duplicate of Q190).

**Q498: What is CSRF and how do you handle it in APIs?**
(Duplicate of Q84).

**Q499: How do you build a custom login page with Spring Boot Security?**
`http.formLogin().loginPage("/login")`. Create Controller for `/login`.

**Q500: How can you invalidate sessions after password reset?**
Find sessions by principal (SessionRegistry). Expire them manually.

---

## 🔹 Spring Boot Internals & Configuration (Questions 501–520)

**Q501: What are the key phases of Spring Boot application lifecycle?**
Starting -> Environment Prepared -> Context Prepared -> Context Loaded -> Refreshed -> Started -> Ready.

**Q502: How do `ApplicationContextInitializer` and `ApplicationListener` fit into Spring Boot startup?**
`Initializer`: callback to customize ConfigurableApplicationContext before refersh.
`Listener`: responds to events (Started, Failed) during lifecycle.

**Q503: What is the purpose of `META-INF/spring.factories`?**
(Duplicate of Q112).

**Q504: What is the use of `SpringApplicationBuilder`?**
(Duplicate of Q241).

**Q505: How can you fail fast on misconfigured Spring Boot applications?**
(Duplicate of Q257).

**Q506: How do you override default beans in Spring Boot auto-configuration?**
Define bean with same name (if `allow-bean-definition-overriding=true`). Or rely on auto-config using `@ConditionalOnMissingBean` (your bean wins).

**Q507: What is lazy initialization and how do you enable it in Spring Boot?**
(Duplicate of Q34/104).

**Q508: How does Spring Boot manage configuration precedence across sources?**
(Duplicate of Q121).

**Q509: What is `ApplicationPidFileWriter` used for?**
Listener that writes the application PID to a file (`application.pid`). Useful for shutdown scripts.

**Q510: How do you detect circular dependency issues in Spring Boot?**
Spring throws `BeanCurrentlyInCreationException` on startup.

**Q511: What is Spring Boot’s support for RabbitMQ and Kafka?**
(Duplicate of Q99).

**Q512: How do you send and consume messages using Kafka in Spring Boot?**
(Duplicate of Q226/227).

**Q513: What is `@KafkaListener` and how do you use it?**
(Duplicate of Q227).

**Q514: How do you use `@EventListener` for in-app events?**
Annotate method with `@EventListener`. It listens for ApplicationEvents published within the same JVM.

**Q515: How do you make Spring events asynchronous?**
Annotate listener with `@Async`. Requires `@EnableAsync`.

**Q516: How do you retry failed Kafka messages in Spring Boot?**
(Duplicate of Q228).

**Q517: How do you configure message acknowledgment in RabbitMQ?**
`spring.rabbitmq.listener.simple.acknowledge-mode=manual`. Call `channel.basicAck`.

**Q518: What is the difference between durable and non-durable queues in RabbitMQ?**
Durable: Persists to disk, survives broker restart. Non-durable: In-memory only.

**Q519: How do you use message converters in Spring Boot?**
Define `MessageConverter` bean (e.g. `Jackson2JsonMessageConverter`). Spring AMQP/Kafka uses it for serialization.

**Q520: How do you handle dead letter queues?**
Configure DLQ argumentation (x-dead-letter-exchange) on queue declaration. Reject messages to route them there.

---

## 🔹 Data Handling (JPA, JDBC, NoSQL) (Questions 521–540)

**Q521: What are the benefits of using Spring Data JPA in Spring Boot?**
(Duplicate of Q61).

**Q522: How do you create custom repository methods in Spring Boot?**
(Duplicate of Q367).

**Q523: What is the purpose of `@Query` annotation in Spring Data JPA?**
(Duplicate of Q65).

**Q524: How do you prevent N+1 query problem in JPA with Spring Boot?**
(Duplicate of Q65).

**Q525: How do you use `EntityManager` in Spring Boot for custom queries?**
Inject `EntityManager`. Use `em.createQuery()` or `CriteriaBuilder`.

**Q526: How do you paginate and sort results in Spring Boot JPA?**
(Duplicate of Q54).

**Q527: What is the use of `@Modifying` annotation in Spring Boot?**
Used on `@Query` methods that change data (UPDATE/DELETE). Tells JPA to clear cache/execute update.

**Q528: How do you map a native SQL query to a DTO?**
(Duplicate of Q275/369).

**Q529: How do you audit entities using Spring Boot and JPA?**
(Duplicate of Q164).

**Q530: How do you use MongoDB with Spring Boot?**
(Duplicate of Q421).

**Q531: How do you implement REST API versioning in Spring Boot?**
(Duplicate of Q111).

**Q532: What are different strategies for API versioning?**
(Duplicate of Q111).

**Q533: How do you generate OpenAPI 3.0 documentation in Spring Boot?**
(Duplicate of Q261).

**Q534: How do you use Swagger annotations in Spring Boot REST controllers?**
`@Operation(summary = "Get user")`. `@ApiResponse(responseCode = "200")`.

**Q535: How can you restrict Swagger to only non-prod environments?**
Enable based on profile. `@Profile("!prod")` on OpenAPI config bean.

**Q536: How do you expose grouped APIs with Springdoc?**
`GroupedOpenApi` bean. Filter by paths (`pathsToMatch("/v1/**")`).

**Q537: How do you add security schemes in Swagger docs?**
`@SecurityScheme(type = SecuritySchemeType.HTTP, scheme = "bearer")`.

**Q538: How do you enable file upload in API documentation?**
Define `MediaType.MULTIPART_FORM_DATA_VALUE`. SpringDoc detects `MultipartFile`.

**Q539: How do you customize Swagger UI look and feel?**
Use properties `springdoc.swagger-ui.*`. or provide custom CSS.

**Q540: How do you document response examples in Swagger?**
`@Content(examples = @ExampleObject(...))`.

---

## 🔹 Reactive Programming (Questions 541–560)

**Q541: What is the difference between `Mono` and `Flux` in Spring WebFlux?**
(Duplicate of Q222).

**Q542: What is the role of `ReactiveCrudRepository`?**
CRUD interface for Reactive streams. Returns Mono/Flux types.

**Q543: How does backpressure work in Spring WebFlux?**
(Duplicate of Q224).

**Q544: What is the use of `RouterFunction` and `HandlerFunction`?**
Functional style alternative to `@Controller`. `RouterFunction` maps URL to `HandlerFunction` (logic).

**Q545: How do you create non-blocking REST APIs using Spring Boot?**
(Duplicate of Q223).

**Q546: How does Spring Boot integrate with Project Reactor?**
It is the default reactive library. Spring WebFlux uses Reactor types (Mono/Flux) in core interfaces.

**Q547: What is the role of `WebClient` in Spring Boot?**
(Duplicate of Q60).

**Q548: How do you handle exceptions reactively in WebFlux?**
Operators: `onErrorResume`, `onErrorReturn`. Global: `WebExceptionHandler`.

**Q549: What is the use of `@EnableWebFlux` annotation?**
Enables manual configuration of WebFlux. typically not needed in Boot (auto-configured).

**Q550: How is WebFlux different from Spring MVC in terms of thread model?**
MVC: One thread per request (blocking). WebFlux: Event loop, few threads shared (non-blocking).

**Q551: How do you write custom condition annotations?**
Implement `Condition`. Create annotation `@Conditional(MyCondition.class)`.

**Q552: What is the use of `ApplicationRunner` vs `CommandLineRunner`?**
(Duplicate of Q248).

**Q553: How do you build a plugin architecture using Spring Boot?**
Scan for Interface implementations. Logic iterates and executes plugins. Service Provider Interface (SPI).

**Q554: How can you register beans programmatically at runtime?**
`GenericApplicationContext.registerBean(...)`.

**Q555: How do you define hierarchical configuration in Spring Boot?**
(Duplicate of Q125).

**Q556: How do you implement request correlation IDs?**
Filter/Interceptor. Check header `X-Request-ID`. If missing, generate. Put in MDC (Logging).

**Q557: What are context refresh events and how are they triggered?**
`ContextRefreshedEvent`. Triggered when ApplicationContext is initialized or refreshed (all beans loaded).

**Q558: How do you listen to shutdown hooks in Spring Boot?**
`@PreDestroy`. or `ContextClosedEvent`.

**Q559: How do you write a custom Spring Boot starter?**
(Duplicate of Q407).

**Q560: How do you include third-party auto-configurations?**
Import via `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.

---

## 🔹 Validation & Exception Handling (Questions 561–580)

**Q561: How do you validate request bodies in Spring Boot using annotations?**
(Duplicate of Q146).

**Q562: What is the difference between `@Valid` and `@Validated`?**
`@Valid`: Standard Java. `@Validated`: Spring variant, supports validation Groups.

**Q563: How do you create a global exception handler in Spring Boot?**
(Duplicate of Q49).

**Q564: What is the purpose of `ResponseEntityExceptionHandler`?**
Base class for `@ControllerAdvice` providing default handling for standard Spring MVC exceptions (405, 400).

**Q565: How do you define custom error responses in Spring Boot REST APIs?**
Create POJO `ErrorResponse`. Return it in `ResponseEntity`.

**Q566: What is a `ConstraintViolationException` and how do you handle it?**
Thrown by JPA/Hibernate validation. Handle in `@ExceptionHandler`.

**Q567: How do you handle 404 errors in Spring Boot?**
`spring.mvc.throw-exception-if-no-handler-found=true`. Catch `NoHandlerFoundException`.

**Q568: How do you customize default error pages in Spring Boot?**
Add `error.html` in templates, or implement `ErrorController`.

**Q569: How do you write unit tests for custom exception handlers?**
`MockMvc` tests triggering the exception. Verify JSON response structure.

**Q570: How do you map validation errors to structured JSON responses?**
Extract `FieldErrors` from `BindingResult`. Map to list of `{field, message}`.

**Q571: How do you implement OAuth 2.0 login with GitHub or Google in Spring Boot?**
(Duplicate of Q185).

**Q572: What is `AuthorizationServerConfigurerAdapter` and how do you use it?**
(Deprecated in Spring Security 5/6). Used to build Auth Server (issuing tokens). Replaced by Spring Authorization Server project.

**Q573: What are scopes in OAuth2 and how do you configure them in Spring Boot?**
Permissions requested (read, write). Configured in client registration properties (`scope: openid,email`).

**Q574: What is the use of `ResourceServerConfigurerAdapter`?**
(Deprecated). Configures app as Resource Server (validates tokens). Now `http.oauth2ResourceServer()`.

**Q575: How do you configure refresh tokens in Spring Boot Security?**
Auth Server feature. Client uses `refresh_token` grant type to get new access token.

**Q576: How do you secure an endpoint with multiple roles?**
`@PreAuthorize("hasAnyRole('ADMIN', 'MANAGER')")`.

**Q577: What is `SecurityFilterChain` in Spring Boot 3?**
Bean creating the security filter chain. Replaces `WebSecurityConfigurerAdapter`.

**Q578: How do you set up custom CORS configurations in Spring Boot Security?**
(Duplicate of Q282).

**Q579: How do you add custom claims in JWT token?**
In Auth Server: `TokenEnhancer` (Legacy) or `OAuth2TokenCustomizer` (SAS). Add to claims map.

**Q580: What is token introspection and how is it handled in Spring Boot?**
Resource server calls Auth Server to check token validity (Opaque tokens). `http.oauth2ResourceServer().opaqueToken()`.

---

## 🔹 Advanced Spring Boot Features (Questions 601-620)
*(Note: Section headings in questions list were reset here, keeping unique numbering)*

**Q601: What is the difference between `@SpringBootApplication` and `@Configuration` + `@EnableAutoConfiguration` + `@ComponentScan`?**
No difference. `@SpringBootApplication` is a meta-annotation comprising the other three.

**Q602: How does Spring Boot handle circular references between beans?**
(Duplicate of Q303).

**Q603: How can you override default error attributes in Spring Boot?**
Extend `DefaultErrorAttributes`. Override `getErrorAttributes()`.

**Q604: What is the use of `ApplicationContextAware` in Spring Boot?**
Interface. Bean implementing it gets reference to `ApplicationContext`. Useful to look up other beans programmatically.

**Q605: What are functional beans in Spring Boot 3 and how are they different?**
Registered via `ApplicationContextInitializer` using lambdas. No reflection, faster startup.

**Q606: How does Spring Boot implement lazy bean initialization at the class level?**
(Duplicate of Q305).

**Q607: What is the purpose of `@ConditionalOnExpression`?**
Condition based on SpEL expression. `@ConditionalOnExpression("${my.prop} == true")`.

**Q608: How does `@AutoConfigureAfter` affect configuration ordering?**
Hints that this AutoConfiguration should run *after* specific other ones. Ensures dependencies (like DataSource) exist first.

**Q609: What’s the role of `SpringFactoriesLoader` in Spring Boot?**
Internal utility to load factory implementations from `META-INF/spring.factories`.

**Q610: How can you write a custom environment post-processor?**
Implement `EnvironmentPostProcessor`. Register in `spring.factories`. Modifies Environment (properties) early in startup.

**Q611: How do you create a Docker image of a Spring Boot application?**
(Duplicate of Q231).

**Q612: What is the use of `spring-boot-maven-plugin`?**
(Duplicate of Q485).

**Q613: How do you configure Spring Boot for Kubernetes deployment?**
(Duplicate of Q235).

**Q614: How do you externalize configuration for Dockerized Spring Boot apps?**
(Duplicate of Q236).

**Q615: How do you create multi-stage Docker builds for Spring Boot apps?**
(Duplicate of Q234).

**Q616: What is the purpose of `BOOT-INF/classes` in Spring Boot jar?**
Layout in Executable Jar. Application classes go here; dependencies go in `BOOT-INF/lib`. Loader reads from here.

**Q617: How do you run Spring Boot in different environments using profiles in CI/CD?**
(Duplicate of Q417). Pass profile as Environment Variable in pipeline.

**Q618: How do you enable health checks for Dockerized Spring Boot apps?**
(Duplicate of Q95).

**Q619: How do you build a Spring Boot executable jar with Gradle?**
`./gradlew bootJar`.

**Q620: How do you deploy a Spring Boot app on Google Cloud Run?**
Build container. Push to GCR. `gcloud run deploy --image gcr.io/...`. Stateless, scales to zero.

---

## 🔹 Spring Boot REST APIs (Questions 621–640)

**Q621: How do you handle file uploads with Spring Boot REST API?**
(Duplicate of Q265). `MultipartFile` parameter.

**Q622: How do you serve static files in a Spring Boot web application?**
Place in `src/main/resources/static` or `public`. Served at root `/` by default.

**Q623: How do you customize `HttpMessageConverter` in Spring Boot?**
Bean of type `HttpMessageConverter`. Or override `configureMessageConverters` in `WebMvcConfigurer`.

**Q624: What is the difference between `@RequestParam`, `@PathVariable`, and `@RequestBody`?**
(Duplicate of Q45/145/267). Query param, URL path segment, Body content.

**Q625: How do you create global CORS configurations for REST APIs?**
(Duplicate of Q262).

**Q626: How can you expose API errors with problem details using Spring Boot?**
Enable RFC 7807 support (Boot 3): `spring.mvc.problemdetails.enabled=true`. Returns standard JSON error format.

**Q627: What are the best practices for designing RESTful endpoints in Spring Boot?**
Nouns for resources, HTTP verbs for actions, plural naming, proper status codes, versioning.

**Q628: How do you implement rate limiting in a Spring Boot REST API?**
(Duplicate of Q150).

**Q629: How do you set default response headers globally in a Spring Boot app?**
Use a Filter. `response.setHeader("X-My-Header", "Val")`.

**Q630: How do you validate a collection of objects in a REST POST request?**
Wrap list in a wrapper DTO `@Valid WrapperDto`. Or use `@Validated` on Controller and `@Valid` on `List<Obj>` param.

**Q631: How do you serve a React/Angular frontend with Spring Boot backend?**
Bundle frontend build (`index.html`, js) into `static` folder. Or run separately and proxy (CORS).

**Q632: How do you handle route refresh issues when serving SPAs from Spring Boot?**
Forward 404s to `index.html` so frontend router handles the path.

**Q633: What is the purpose of `ResourceHandlerRegistry` in Spring Boot?**
Configures static resource serving locations and cache headers.

**Q634: How do you enable Gzip compression for static content in Spring Boot?**
(Duplicate of Q347).

**Q635: How do you proxy API requests to Spring Boot from a frontend dev server?**
Configure proxy in `package.json` (React) or `proxy.conf.json` (Angular) to point to `localhost:8080`.

**Q636: What’s the best practice to deploy a full-stack Spring Boot + JS framework app?**
Separate artifacts (CDN for frontend, API for backend) is best for scale. Bundled (Monolith) is best for simplicity.

**Q637: How can you bundle a frontend build into a Spring Boot JAR?**
Build frontend (npm build). Copy output to `src/main/resources/static` via Maven/Gradle plugin before package step.

**Q638: How do you prevent browser caching for frontend served via Spring Boot?**
Set `Cache-Control: no-cache` headers via `ResourceHandlerRegistry` or Filter, commonly for `index.html`.

**Q639: How do you secure frontend assets from unauthorized users?**
Apply Security Filters. Allow public access to `/index.html`, `/assets/**`. Require auth for `/api/**`.

**Q640: How do you handle 404s in SPAs routed through Spring Boot?**
(Duplicate of Q632).

---

## 🔹 Asynchronous & Scheduling (Questions 641–680)
*(Adjusted numbering to align with flow)*

**Q661: How do you enable and use `@Async` in Spring Boot?**
`@EnableAsync` on config. `@Async` on method. Method runs in separate thread.

**Q662: How can you define a custom executor for `@Async` methods?**
Define `Executor` bean. Refer to it: `@Async("myExecutor")`.

**Q663: What is the default behavior of exception handling in asynchronous methods?**
Void return: Exception lost (logged). Future return: Exception captured in Future. Implement `AsyncUncaughtExceptionHandler` for global handling.

**Q664: How do you schedule tasks with `@Scheduled`?**
`@EnableScheduling`. `@Scheduled(fixedRate = 1000)` on method.

**Q665: How do you run scheduled jobs only in a specific environment or profile?**
`@Profile("prod")` on the component containing the scheduled method.

**Q666: How do you use `CronTrigger` with custom cron expressions?**
`@Scheduled(cron = "0 0 * * * *")`. Or configure `TaskScheduler` manually with `new CronTrigger(...)`.

**Q667: How do you retry a failed async method in Spring Boot?**
Use `@Retryable` (Spring Retry). `@Async` alone doesn't retry.

**Q668: How do you prevent overlapping of scheduled jobs in Spring Boot?**
Default scheduler is single-threaded (won't overlap unless async). To run parallel, configure pool. To prevent across instances: ShedLock.

**Q669: How do you ensure thread safety in scheduled tasks?**
Tasks should be stateless. If accessing shared state, use synchronization or atomic variables.

**Q670: How do you use `TaskScheduler` manually in Spring Boot?**
Inject `TaskScheduler`. Call `.schedule(runnable, instant/trigger)`.

**Q681: How does Spring Boot fit in a microservices architecture?**
Provides independent, executable services. Integrates with Spring Cloud for distributed system patterns.

**Q682: How do you implement service discovery in Spring Boot without Eureka?**
Kubernetes Service Discovery, Consul, or Zookeeper.

**Q683: How do you use OpenFeign with Spring Boot?**
(Duplicate of Q464).

**Q684: What is Spring Cloud Gateway and how does it integrate with Spring Boot?**
API Gateway based on WebFlux. Routes traffic to microservices. Handles cross-cutting concerns (Auth, Rate Limit).

**Q685: How do you propagate headers across microservices in Spring Boot?**
Use Sleuth (auto-propagates trace headers). Or `FeignRequestInterceptor` to manually copy headers.

**Q686: How do you implement distributed tracing with Spring Boot?**
(Duplicate of Q219/468).

**Q687: How do you secure inter-service communication in Spring Boot?**
mTLS, or JWT passed in Authorization header. Service-to-service auth (Client Credentials flow).

**Q688: How do you manage configuration across microservices in Spring Boot?**
Spring Cloud Config Server. Centralized git repo.

**Q689: What is Hystrix and how is it used in Spring Boot?**
(Deprecated). Circuit Breaker library. Replaced by Resilience4j.

**Q690: How do you handle request context propagation in microservices?**
Async context doesn't propagate automatically. Use `DelegatingSecurityContextExecutor` or MDCAware decorators.

**Q701: How do you implement two-factor authentication in Spring Boot?**
(Duplicate of Q286).

**Q702: What is the use of `SecurityContextHolder`?**
(Duplicate of Q385).

**Q703: How do you restrict access based on client IP addresses?**
`requestMatchers(...).hasIpAddress("192.168...")`.

**Q704: How do you implement session timeout and auto-logout?**
`server.servlet.session.timeout=30m`. Frontend detects 401/redirect and cleans up.

**Q705: What is the difference between stateless and stateful sessions in Spring Security?**
(Duplicate of Q285).

**Q706: How do you configure multiple authentication mechanisms (e.g., form + token)?**
Configure multiple filters in the chain. E.g. `UsernamePasswordAuthenticationFilter` and `JwtFilter`.

**Q707: How do you integrate Keycloak with Spring Boot?**
`Keycloak-spring-boot-starter` (deprecated) or standard `spring-boot-starter-oauth2-client` + `oauth2-resource-server`. Point to Keycloak realm.

**Q708: How do you secure WebSocket connections in Spring Boot?**
(Duplicate of Q281).

**Q709: What are security implications of exposing actuator endpoints?**
Information disclosure (Heap dumps, env vars). RCE vulnerability in older versions if Jolokia exposed. Always secure them.

**Q710: How do you protect against session fixation in Spring Boot?**
Enabled by default. `http.sessionManagement().sessionFixation().migrateSession()`. Changes SessionID on login.

---

## 🔹 Spring Boot Advanced Configuration (Questions 721–740)
*(Adjusted numbering for unique sequence)*

**Q721: What is the role of `@ImportResource` in Spring Boot?**
Loads legacy XML configuration files (`applicationContext.xml`) into the Spring Boot context.

**Q722: How can you log all application properties at runtime?**
(Duplicate of Q320/891).

**Q723: What are meta-annotations in Spring Boot, and how are they useful?**
Annotations annotated with other annotations. Useful for creating composed, semantic annotations (e.g. `@RestController` is `@Controller` + `@ResponseBody`).

**Q724: How can you provide default values for configuration properties?**
In Java: `@Value("${prop:default}")`. In ConfigProperties: assign default value to field.

**Q725: How does Spring Boot resolve `@Value("${}")` placeholders?**
(Duplicate of Q316). `PropertySourcesPlaceholderConfigurer`.

**Q726: What are relaxed binding rules in Spring Boot configuration?**
(Duplicate of Q127).

**Q727: How do you add additional property sources at runtime?**
(Duplicate of Q123).

**Q728: How can you use system properties to override configuration in Spring Boot?**
`-Dproperty.name=value`. Higher precedence than config files.

**Q729: How do you encrypt and decrypt sensitive configuration values?**
(Duplicate of Q315).

**Q730: What is the role of `@Conditional` annotations in Spring Boot?**
(Duplicate of Q35/106). Base for all auto-configuration.

**Q731: How can you define a composite primary key in Spring Boot JPA?**
`@EmbeddedId` or `@IdClass`.

**Q732: How do you implement soft deletes using JPA and Spring Boot?**
(Duplicate of Q163).

**Q733: What is the difference between `fetch = FetchType.LAZY` and `EAGER`?**
(Duplicate of Q68).

**Q734: How do you prevent infinite recursion in bi-directional entity relationships?**
Use `@JsonManagedReference` (parent) and `@JsonBackReference` (child). Or `@JsonIgnore` on one side of relation.

**Q735: How do you dynamically filter entities using Spring JPA Specifications?**
(Duplicate of Q162).

**Q736: What is the use of `@EntityGraph` in Spring Data JPA?**
(Duplicate of Q165).

**Q737: How can you use query by example (QBE) in Spring Boot JPA?**
`Example.of(probe)`. `repo.findAll(example)`. Finds records matching non-null fields of the probe object.

**Q738: How do you call a stored procedure using Spring Data JPA?**
(Duplicate of Q166).

**Q739: What is the difference between `CrudRepository` and `JpaRepository`?**
(Duplicate of Q64).

**Q740: How do you use projections in Spring Boot JPA for custom DTOs?**
Define interface with getter methods matching entity properties. `findAllProjectedBy(...)`.

---

## 🔹 Testing in Spring Boot (Questions 741–760)

**Q741: How do you write integration tests for a Spring Boot REST controller?**
(Duplicate of Q90).

**Q742: What is the purpose of `@SpringBootTest`?**
(Duplicate of Q88).

**Q743: How do you use `@MockBean` and `@SpyBean`?**
(Duplicate of Q208).

**Q744: What’s the difference between `@WebMvcTest` and `@DataJpaTest`?**
(Duplicate of Q91).

**Q745: How do you load a custom application.properties for test cases?**
(Duplicate of Q209).

**Q746: How can you mock an external API call in Spring Boot tests?**
Use `WireMock`. Or mock the RestTemplate/WebClient bean.

**Q747: How do you test scheduled jobs in Spring Boot?**
(Duplicate of Q297).

**Q748: How do you run tests on different Spring profiles?**
(Duplicate of Q209/292).

**Q749: How do you disable security during tests?**
`@AutoConfigureMockMvc(addFilters = false)`. Or use `@WithMockUser`.

**Q750: What is the use of `TestEntityManager`?**
Helper in `@DataJpaTest`. Allows persisting/finding entities in tests without using the Repository.

---

## 🔹 Spring Boot & Observability (Questions 761–800)

**Q761: How do you enable tracing with Spring Boot?**
(Duplicate of Q219).

**Q762: What are metrics and how do you expose them in Spring Boot?**
(Duplicate of Q217).

**Q763: What’s the use of Micrometer in Spring Boot?**
(Duplicate of Q214).

**Q764: How do you export metrics to Prometheus?**
(Duplicate of Q213).

**Q765: How do you integrate Spring Boot with Grafana dashboards?**
(Duplicate of Q213). Grafana reads from Prometheus.

**Q766: What is the difference between counters and gauges in metrics?**
Counter: Monotonically increasing (requests count). Gauge: Fluctuating value (current memory usage).

**Q767: How can you monitor JVM metrics using Spring Boot?**
(Duplicate of Q217). Built-in.

**Q768: How do you customize health indicators in Spring Boot?**
(Duplicate of Q211).

**Q769: What’s the role of `HealthContributor` interface?**
Base interface for `HealthIndicator`. Used to create composite health checks.

**Q770: How do you monitor custom business KPIs in Spring Boot?**
Register custom Micrometer metrics (`Counter`, `Gauge`) in your business service code.

**Q771: How do you enable lazy loading globally for JPA?**
`spring.jpa.properties.hibernate.enable_lazy_load_no_trans=true` (Anti-pattern, but exists). Correct way is OpenSessionInView (default) or DTOs.

**Q772: What are the best practices for caching in Spring Boot?**
Cache immutable read-heavy data. Define TTL. Handle cache invalidation. Use `@Cacheable`.

**Q773: How do you implement second-level cache in Spring Boot JPA?**
Enable Shared Cache in Hibernate properties. Use provider like Ehcache/Redis. Annotate Entity `@Cacheable`.

**Q774: What is the use of `@EnableCaching`?**
Enables Spring's annotation-driven cache management capability.

**Q775: How can you use Redis as a cache provider in Spring Boot?**
(Duplicate of Q430).

**Q776: How do you profile memory usage in a Spring Boot app?**
Heap dumps, Java Flight Recorder (JFR), or Actuator Metrics.

**Q777: How do you handle large file uploads efficiently?**
(Duplicate of Q269).

**Q778: How can you compress REST API responses?**
(Duplicate of Q347).

**Q779: How do you minimize startup time in large Spring Boot applications?**
Lazy Init, remove unused starters, Spring Indexer, CDS (Class Data Sharing).

**Q780: How do you implement custom caching strategies?**
Implement `CacheManager` / `Cache` interfaces manually.

---

## 🔹 Spring Boot with External Systems & Tools (Questions 801–820)
*(Numbering adjusted due to source overlap)*

**Q801: How do you use Spring Boot with RabbitMQ?**
(Duplicate of Q230).

**Q802: How do you publish and consume Kafka messages using Spring Boot?**
(Duplicate of Q226).

**Q803: How do you configure and use ActiveMQ in Spring Boot?**
`spring-boot-starter-activemq`. `JmsTemplate` to send. `@JmsListener` to receive.

**Q804: What is the difference between `@KafkaListener` and `@RabbitListener`?**
Conceptually same. Specific to underlying transport protocol (Kafka vs AMQP).

**Q805: How do you handle dead-letter queues in Spring Boot?**
(Duplicate of Q520).

**Q806: How do you use Spring Boot with Elasticsearch?**
`spring-boot-starter-data-elasticsearch`. `ElasticsearchRepository`.

**Q807: How do you configure Spring Boot to use an external identity provider?**
(Duplicate of Q185).

**Q808: How do you secure external API calls using OAuth2 client credentials?**
`WebClient` configured with `ServletOAuth2AuthorizedClientExchangeFilterFunction`. Auto-fetches tokens.

**Q809: How do you consume a SOAP web service in Spring Boot?**
(Duplicate of Q443).

**Q810: How do you expose a SOAP endpoint from Spring Boot?**
`spring-boot-starter-web-services`. `MessageDispatcherServlet`. `@Endpoint` annotation.

---

## 🔹 Spring Boot with DevOps & CI/CD (Questions 821–840)

**Q821: How do you build and package a Spring Boot application using Maven?**
`mvn clean package`.

**Q822: How do you build and package a Spring Boot application using Gradle?**
`./gradlew build`.

**Q823: How do you generate a Docker image from a Spring Boot application?**
(Duplicate of Q231).

**Q824: What is the use of the `spring-boot-maven-plugin`?**
(Duplicate of Q485).

**Q825: How do you set up a multi-stage Dockerfile for Spring Boot?**
(Duplicate of Q234).

**Q826: How do you externalize configuration in Docker containers for Spring Boot?**
(Duplicate of Q236).

**Q827: How do you integrate Spring Boot with Jenkins pipelines?**
Standard CI steps: Checkout -> Build (mvn/gradle) -> Test -> Docker Build -> Push.

**Q828: How do you deploy Spring Boot applications on Kubernetes?**
(Duplicate of Q235).

**Q829: What is the role of `ConfigMap` and `Secrets` in Spring Boot on K8s?**
Mapped to Volume or Env Vars. Spring Boot reads them as property sources.

**Q830: How do you perform blue-green deployments with Spring Boot?**
(Duplicate of Q490).

**Q831: What is the benefit of using `actuator/health` endpoint in a CI/CD pipeline?**
Used as smoke test target after deployment to verify app started successfully.

**Q832: How do you monitor a Spring Boot app deployed via Docker?**
Container monitoring tools (cAdvisor, Docker stats) + internal Actuator metrics.

**Q833: How do you configure rolling updates in Kubernetes for Spring Boot?**
(Duplicate of Q238).

**Q834: How do you manage multiple Spring Boot services in a CI/CD pipeline?**
Matrix builds, or separate pipeline per service/repo.

**Q835: How do you version REST APIs in Spring Boot and handle backward compatibility?**
(Duplicate of Q111).

**Q836: What strategies can be used to gracefully shut down Spring Boot services?**
(Duplicate of Q98).

**Q837: How do you handle logging and log aggregation in Spring Boot microservices?**
(Duplicate of Q433).

**Q838: What are liveness and readiness probes in Kubernetes for Spring Boot apps?**
Exposed via Actuator. `/actuator/health/liveness`, `/actuator/health/readiness`.

**Q839: How do you inject secrets into Spring Boot apps in cloud environments?**
(Duplicate of Q124/829).

**Q840: How do you enable zero-downtime deployments for Spring Boot?**
(Duplicate of Q238).

---

## 🔹 Spring Boot 3.x & Major Updates (Questions 841–860)

**Q841: What are the key differences between Spring Boot 2 and 3?**
Baseline Java 17. Jakarta EE 9/10 (javax -> jakarta). Native Image support (GraalVM). Observability improvements.

**Q842: How does Jakarta EE namespace impact Spring Boot 3 apps?**
Imports must change from `javax.*` to `jakarta.*` (Servlet, JPA, Validation).

**Q843: What is native image support in Spring Boot 3 using GraalVM?**
First-class support. AOT processing generates GraalVM-compatible sources to build standalone native executable.

**Q844: How do you compile a native executable with Spring Boot 3?**
`mvn -Pnative native:compile`.

**Q845: What are the limitations of GraalVM native images in Spring Boot?**
Slower build time. No dynamic class loading. Reflection requires configuration/hints.

**Q846: How does reflection configuration differ in GraalVM builds?**
Need `remote-reflection-configuration.json` usually. Spring AOT attempts to generate hints automatically.

**Q847: What is `@NamedNativeQuery` and how is it used in JPA with Spring Boot 3?**
Same as before, but mapped to `jakarta.persistence` API.

**Q848: How do you handle `jakarta.persistence` migration in Spring Boot 3?**
Update dependency to Hibernate 6. Change imports.

**Q849: What is the role of AOT (Ahead of Time) compilation in Spring Boot 3?**
Optimizes application context computation at build time. Generates code to start context faster in native mode.

**Q850: How do you enable experimental features in Spring Boot 3?**
Usually via config flags or snapshot versions. Native support is now stable.

**Q851: How does `@RestClientTest` change in Spring Boot 3?**
Updated to support `RestClient` (new fluent client) if available.

**Q852: How do you migrate a legacy app from Spring Boot 2.x to 3.x?**
Upgrade to 2.7 first. Use `spring-boot-properties-migrator`. Upgrade JDK to 17. Run migration recipes (OpenRewrite).

**Q853: What is the replacement for `javax.validation` in Spring Boot 3?**
`jakarta.validation`.

**Q854: What is `@ContextConfiguration` and how is it affected by Spring Boot 3?**
Standard testing annotation. Unchanged functionality, but interacts with AOT generated context.

**Q855: How does Spring Boot 3 improve startup performance?**
AOT optimizations (even on JVM), Spring Framework 6 baseline updates.

**Q856: What is `NativeHint` in Spring Boot native images?**
Annotation to provide hints to GraalVM about reflection/resources needed (if AOT misses them).

**Q857: How do you use `Functional Bean Registration` in Spring Boot 3?**
(Duplicate of Q605).

**Q858: What are core annotations deprecated in Spring Boot 3?**
Some configuration properties. `WebSecurityConfigurerAdapter` (removed).

**Q859: How do you test native executables of Spring Boot?**
`mvn -Pnative native:test`. Runs standard tests against the native binary.

**Q860: What tools are used to analyze native image memory in Spring Boot 3?**
GraalVM Native Image tools (dashboard).

---

## 🔹 Real-World Use Cases (Questions 861–880)

**Q861: How do you implement audit logging using Spring Boot and JPA?**
(Duplicate of Q164).

**Q862: How do you manage multi-language support (i18n) in Spring Boot?**
`MessageSource` bean. Resource bundles (`messages.properties`, `messages_fr.properties`).

**Q863: How do you enforce business rules using Spring Expression Language (SpEL)?**
Use inside `@PreAuthorize`, `@Value`, or manually via `ExpressionParser`.

**Q864: How do you implement custom HTTP status codes in responses?**
(Duplicate of Q47).

**Q865: How do you handle pagination and sorting in Spring Boot REST APIs?**
(Duplicate of Q54).

**Q866: How do you expose a CSV/Excel export endpoint in Spring Boot?**
Set `Content-Disposition`. Write to `HttpServletResponse.getOutputStream()` using Apache POI (Excel) or CSV printer.

**Q867: How do you implement a file preview service in Spring Boot?**
Return content with inline disposition and correct MIME type (image/png, application/pdf).

**Q868: How can you stream large data sets from DB to client efficiently?**
Use `Stream<Entity>` in Transaction. Write to `ResponseBodyEmitter` or `StreamingResponseBody`.

**Q869: How do you implement retry logic with exponential backoff?**
`@Retryable(backoff = @Backoff(delay = 1000, multiplier = 2))`.

**Q870: How do you implement API quota and usage limit per user?**
Rate limiting (Bucket4j) keyed by User ID.

**Q871: How do you use `@ControllerAdvice` to globally format all exceptions?**
(Duplicate of Q49).

**Q872: How do you write modular Spring Boot applications using modules?**
(Duplicate of Q109).

**Q873: How do you implement read/write DB splitting in Spring Boot?**
`AbstractRoutingDataSource`. Determine target (Primary/Replica) based on transaction read-only flag.

**Q874: How do you run dynamic SQL queries based on user input safely?**
(Duplicate of Q65/162). Criteria API or Specifications.

**Q875: How do you write a Spring Boot CLI application?**
(Duplicate of Q11).

**Q876: How do you use `MultiPartResolver` for multiple file uploads?**
Standard `MultipartFile[]` array in Controller.

**Q877: How do you build an event notification system in Spring Boot?**
Internal: Spring Events. External: Websockets (STOMP) or Push Notifications (Firebase).

**Q878: How do you encrypt and store files securely in Spring Boot?**
Encrypt stream (AES) before writing to disk/S3. Store key in Vault.

**Q879: How do you build a reactive system using Spring Boot WebFlux?**
(Duplicate of Q221).

**Q880: How do you implement real-time chat using WebSocket in Spring Boot?**
`@EnableWebSocketMessageBroker`. Use STOMP over WebSocket. `@MessageMapping` to handle messages.

---

## 🔹 Cloud Providers & Troubleshooting (Questions 881–900)

**Q881: How do you deploy a Spring Boot app to AWS Elastic Beanstalk?**
Build JAR. Upload to EB Console or use EB CLI `eb deploy`.

**Q882: How do you configure Spring Boot with AWS RDS?**
Standard JDBC URL pointing to RDS endpoint.

**Q883: How do you integrate AWS S3 with Spring Boot?**
AWS SDK (v2). Create `S3Client`. `putObject`, `getObject`.

**Q884: How do you use AWS Secrets Manager with Spring Boot?**
`spring-cloud-starter-aws-secrets-manager-config`. Auto-loads secrets into Environment properties.

**Q885: How do you implement an SQS listener in Spring Boot?**
`@SqsListener` (Spring Cloud AWS).

**Q886: How do you send push notifications using Firebase in Spring Boot?**
Firebase Admin SDK. `FirebaseMessaging.getInstance().send(message)`.

**Q887: How do you integrate Azure Key Vault with Spring Boot?**
`spring-cloud-azure-starter-keyvault`.

**Q888: How do you use GCP Pub/Sub in Spring Boot?**
`spring-cloud-gcp-starter-pubsub`. `PubSubTemplate`.

**Q889: How do you use AWS Lambda as a Spring Boot function?**
(Duplicate of Q479).

**Q890: How do you access metadata from EC2 inside Spring Boot?**
Query `169.254.169.254`. Or use Spring Cloud AWS context.

**Q891: How do you debug circular dependency errors?**
Analyze stack trace `BeanCurrentlyInCreationException`. Refactor to setter injection or `@Lazy`.

**Q892: How do you resolve "port already in use" errors in Spring Boot?**
Identify process (`lsof -i :8080`). Kill it. or change `server.port`.

**Q893: How do you analyze thread dumps of Spring Boot apps?**
Capture via Actuator (`/actuator/threaddump`) or `jstack`. Look for BLOCKED threads.

**Q894: How do you detect and fix memory leaks in Spring Boot?**
(Duplicate of Q34).

**Q895: How do you handle long GC pause issues in production?**
Analyze GC logs. Tune Heap Sizing. Switch GC (G1GC, ZGC).

**Q896: How do you enable debug logging for only specific packages?**
(Duplicate of Q114).

**Q897: How do you diagnose high CPU usage from Spring Boot?**
Profile (AsyncProfiler). Find hotspot methods.

**Q898: How do you measure and improve database connection pool usage?**
Metrics (`hicari.pool.*`). Tune `maximum-pool-size`.

**Q899: How do you capture startup performance metrics?**
Startup Actuator endpoint (`/actuator/startup`). Requires `BufferingApplicationStartup`.

**Q900: How do you restart a Spring Boot app programmatically?**
`Restarter.getInstance().restart()` (DevTools only) or `context.close()` and create new `SpringApplication`. Cloud platforms restart containers externally.

---

## 🔹 Security, Messaging & Advanced API (Questions 901–1000)

**Q901: How do you configure multiple security filters in Spring Boot?**
(Duplicate of Q386).

**Q902: What is the difference between `SecurityFilterChain` and `WebSecurityConfigurerAdapter`?**
Adapter is legacy (inheritance). Component-based configuration (Beans) is the new standard.

**Q903: How do you configure CORS in Spring Boot Security?**
(Duplicate of Q282).

**Q904: How do you handle CSRF tokens in a stateless REST API?**
Usually disable CSRF for stateless. If needed, send token in cookie/header.

**Q905: How do you implement multi-role authentication in Spring Boot?**
Users have List of GrantedAuthorities.

**Q906: How do you secure WebSocket connections with Spring Security?**
(Duplicate of Q281).

**Q907: How do you customize login and logout URLs in Spring Security?**
`http.formLogin().loginPage("/my-login").logoutUrl("/my-logout")`.

**Q908: How do you store session tokens in Redis for Spring Boot apps?**
`spring-session-data-redis`. Replaces standard HttpSession with Redis-backed session.

**Q909: What is the difference between role-based and permission-based security?**
Role: Broad group (`ADMIN`). Permission: Granular action (`READ_USER`).

**Q910: How do you dynamically restrict access to APIs at runtime?**
Data-driven security. Check DB for permissions in Filter/Guard.

**Q911: How do you implement HATEOAS in a Spring Boot REST API?**
(Duplicate of Q147).

**Q912: How do you cache GET responses in a Spring Boot REST API?**
(Duplicate of Q772).

**Q913: How do you validate nested JSON objects in Spring Boot?**
(Duplicate of Q630).

**Q914: What are the best practices for REST error handling in Spring Boot?**
(Duplicate of Q49/148).

**Q915: How do you add Swagger 3.0 (OpenAPI) to your project?**
(Duplicate of Q57/261).

**Q916: How do you version REST APIs using headers in Spring Boot?**
(Duplicate of Q111).

**Q917: How do you support both XML and JSON responses in Spring Boot?**
(Duplicate of Q53).

**Q918: How do you create a generic exception handling framework in Spring Boot?**
(Duplicate of Q49).

**Q919: How do you throttle specific endpoints per IP/user?**
(Duplicate of Q150).

**Q920: How do you handle multipart form data in Spring Boot REST API?**
(Duplicate of Q46/265).

**Q930: How do you integrate RabbitMQ with Spring Boot?**
(Duplicate of Q230).

**Q940: How do you use Kafka with Spring Boot?**
(Duplicate of Q226).

**Q950: How do you serve a React or Angular SPA from a Spring Boot backend?**
(Duplicate of Q631).

**Q960: How do you configure CORS globally for frontend apps?**
(Duplicate of Q51).

**Q970: How do you protect static assets from unauthorized access?**
(Duplicate of Q639).

**Q980: How do you implement feature flags in Spring Boot applications?**
(Duplicate of Q153).

**Q990: How do you use Spring Boot with Apache Camel?**
(Duplicate of Q450).

**Q1000: How do you modularize a large Spring Boot monolith into smaller components?**
(Duplicate of Q109/872).
