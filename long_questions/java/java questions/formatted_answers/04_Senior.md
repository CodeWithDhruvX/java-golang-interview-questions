# Senior Level Java Interview Questions

## From 24 Mixed Concepts Patterns DB Testing
# 24. Mixed Concepts (Advanced SQL, Design Patterns, JDBC, Testing)

**Q: SQL: Handle Nulls and Unions**
> "**Handling Nulls**: Use `COALESCE(column_name, 'Default Value')` (standard SQL) or `IFNULL()` (MySQL). It checks if the value is null and creates a fallback.
>
> **Union vs Union All**:
> *   `UNION` combines results from two queries and **removes duplicates**. It's slower because it has to sort/check to find those duplicates.
> *   `UNION ALL` just blind appends everything. It's much faster. Use it if you know your data sets are unique or if you *want* duplicates."

**Indepth:**
> **Performance**: `UNION ALL` is significantly faster than `UNION` because it involves simple concatenation, whereas `UNION` requires sorting/hashing the entire result set to identify duplicates. Always prefer `UNION ALL` unless you specifically need deduplication.


---

**Q: SQL: Self Join and Ranking**
> "**Self Join**: Joining a table to itself.
> Useful for hierarchy (Managers/Employees) or finding pairs (Finding 2 employees with the same salary: `JOIN on e1.salary = e2.salary AND e1.id <> e2.id`).
>
> **Ranking**:
> Use Window Functions: `RANK()`, `DENSE_RANK()`, or `ROW_NUMBER()`.
> *   `ROW_NUMBER()`: Unique ID for every row (1, 2, 3, 4).
> *   `RANK()`: Skips numbers for ties (1, 1, 3).
> *   `DENSE_RANK()`: No skipping (1, 1, 2). Usually the best one for 'Nth highest' questions."

**Indepth:**
> **Optimization**: Window functions like `RANK()` are generally much more performant than using self-joins or correlated subqueries for ranking problems because the database engine can optimize the window operation in a single pass.


---

**Q: Factory vs Builder Pattern**
> "**Factory Pattern**: You have a method that creates objects for you. You don't say `new Dog()`, you say `AnimalFactory.create("Dog")`. It hides the complex logic of *instantiation*.
>
> **Builder Pattern**: Used for complex objects with many parameters (some optional). Instead of a constructor with 10 arguments (`new Pizza("Thin", null, "Pepperoni", true...)`), you chain methods:
> `new PizzaBuilder().crust("Thin").topping("Pepperoni").build()`. It's readable and flexible."

**Indepth:**
> **Fluency**: The Builder pattern often uses a static inner class `Builder` to ensure thread safety during construction and immutability of the final object. It solves the "Telescoping Constructor" anti-pattern.


---

**Q: Observer and Strategy Patterns**
> "**Observer**: Push notifications. An object (Subject) has a list of 'Subscribers'. When something changes, it loops through the list and calls `notify()`. Used in Event Listeners and Chat apps.
>
> **Strategy**: Pluggable behavior. You define a family of algorithms (SortFast, SortSlow, SortReverse) and make them interchangeable. You can pass the specific 'Strategy' into the method at runtime using an interface."

**Indepth:**
> **Decoupling**: Strategy allows you to change the guts of an object (how it does something) without changing the object itself. Observer allows you to react to changes without polling. Both heavily rely on Interfaces.


---

**Q: JDBC Steps (The 5 Steps)**
> "1.  **Load Driver**: `Class.forName(...)` (Though modern JDBC drivers auto-load).
> 2.  **Get Connection**: `DriverManager.getConnection(url, user, pass)`.
> 3.  **Create Statement**: `con.prepareStatement(sql)`.
> 4.  **Execute**: `stmt.executeQuery()` (for SELECT) or `stmt.executeUpdate()` (for INSERT/UPDATE).
> 5.  **Close**: Always close `ResultSet`, `Statement`, and `Connection` in a `try-with-resources` block to avoid leaks."

**Indepth:**
> **Modern JDBC**: In production, nobody writes raw JDBC like this anymore. We use Connection Pools (HikariCP) to reuse connections and Frameworks (Spring JDBC, Hibernate) to handle the boilerplate and resource cleanup.


---

**Q: PreparedStatement vs Statement**
> "Always use **PreparedStatement**.
> 1.  **Security**: It prevents SQL Injection by escaping inputs automatically.
> 2.  **Performance**: The database can compile and cache the query plan effectively because the structure (`SELECT * FROM users WHERE id = ?`) doesn't change, only the parameter changes."

**Indepth:**
> **Caching**: The DB parses the SQL template `SELECT * FROM users WHERE id = ?` once and caches the execution plan. If you run it 1000 times with different IDs, it reuses that plan. With `Statement`, every query string is different (`id=1`, `id=2`), forcing re-compilation.


---

**Q: ExecutorService & CompletableFuture**
> "**ExecutorService**: The manager for threads. You give it tasks (`Runnable` or `Callable`), and it assigns them to a pool of worker threads.
> *   `Callable`: Like Runnable, but it can return a result (`Future`) and throw Exceptions.
>
> **CompletableFuture**: The modern, non-blocking way to do async.
> It lets you chain tasks: 'Do task A, **then** do task B, **then** handle the error'. It keeps your main thread free."

**Indepth:**
> **Composition**: The real power of `CompletableFuture` is composition. `thenCompose()` allows you to chain dependent async operations (result of A is input to B), while `thenCombine()` runs independent operations in parallel and merges results.


---

**Q: Unit Testing (JUnit Basics)**
> "A Unit Test checks a small, isolated piece of logic.
>
> **Key Annotations**:
> *   `@Test`: Marks a method as a test.
> *   `@BeforeEach`: Runs before *every* test (good for resetting data).
> *   `@BeforeAll`: Runs once before *anything* (good for expensive setup like DB connections).
>
> Use `Assertions` to verify: `assertEquals(expected, actual)`.
> Mock dependencies (using Mockito) so you are testing *only* your class, not the database or network."

**Indepth:**
> **Isolation**: A unit test should never touch the file system, network, or database. If it does, it's an Integration Test. Unit tests must be fast (milliseconds) and deterministic. Use Mocking to fake external dependencies.


## From 30 Spring Boot Data Security
# 30. Spring Boot (Data Access & Security)

**Q: Database Views in JPA**
> "JPA doesn't officially distinguish between a Table and a View.
> You map a View exactly like a Table: using `@Entity`.
>
> **The Trick**: Since Views are read-only, you should verify that you don't accidentally try to save data to it. Use `@Immutable` (Hibernate annotation) on the entity to prevent writes."

**Indepth:**
> **Performance**: Be careful. If the View logic is complex (joins, aggregations), querying it might be slow. The View code runs inside the DB engine, so specialized indexes on the underlying tables are crucial.


---

**Q: Pagination with Slice vs Page**
> "Both are used for pagination, but `Page` is more expensive.
>
> *   **Page**: Returns the data chunk **plus** the total count of pages/rows. This requires an extra `COUNT(*)` query, which can be slow on huge tables.
> *   **Slice**: Returns the data chunk and just a flag `hasNext()`. It doesn't know the total size. Use this for 'Infinite Scroll' features where you don't care if there are 100 or 1000 pages left, just 'is there more?'."

**Indepth:**
> **Mobile Apps**: `Slice` is perfect for mobile "Infinite Scroll". Calculating total pages (`COUNT(*)`) is expensive and often unnecessary for a Twitter-like feed where you just want the next 10 items.


---

**Q: Custom Repository Implementation**
> "Sometimes the standard `save/findAll` isn't enough, and `@Query` is too messy.
>
> 1.  Create an interface `MyRepoCustom` with your method: `complexSearch()`.
> 2.  Create a class `MyRepoImpl` implementing it. **CRITICAL**: The name must end in `Impl` (by default).
> 3.  Inject `EntityManager` inside `Impl` and write full custom logic (Criteria API or Native SQL).
> 4.  Have your main interface extend `JpaRepository` AND `MyRepoCustom`. Spring merges them automatically."

**Indepth:**
> **Composition**: This pattern (Composition over Inheritance) allows you to keep the clean `findBy` methods of `JpaRepository` while injecting completely arbitrary code execution into the same bean.


---

**Q: @Secured vs @PreAuthorize**
> "**@Secured**: Older, simple, but limited. `@Secured("ROLE_ADMIN")`.
>
> "**@PreAuthorize**: The modern standard. It supports SpEL (Spring Expression Language), allowing widely powerful logic:
> `@PreAuthorize("hasRole('ADMIN') or #param.name == authentication.name")`.
> Always use `@PreAuthorize` unless you're on a very old legacy system."

**Indepth:**
> **SecurityContext**: SpEL checks happen against the `SecurityContext`. You can even access method arguments: `@PreAuthorize("#order.owner == authentication.name")`. This is fine-grained "Instance Level Security".


---

**Q: OAuth2 Login Flow**
> "In Spring Boot, it's almost zero-config.
> 1.  Add `spring-boot-starter-oauth2-client`.
> 2.  In `application.yml`, register your provider (Google/GitHub) with `clientId` and `clientSecret`.
>
> Spring automatically configures the login page (`/login`), handles the redirect to Google, catches the callback code, exchanges it for a Token, and logs the user in."

**Indepth:**
> **CommonAuth2**: Note that "Social Login" (Google) and "Enterprise SSO" (Okta/Active Directory) both use the standardized OAuth2/OIDC protocol. Spring handles them identically.


---

**Q: Filter Chain (Spring Security)**
> "Spring Security is essentially a giant chain of Servlet Filters.
> Request -> [LogoutFilter] -> [UsernamePasswordAuthFilter] -> [BasicAuthFilter] -> ... -> [MyCustomFilter] -> Controller.
>
> If any filter throws an exception or denies access, the request stops there. It never reaches your controller. You customize security by inserting your own filters into this chain."

**Indepth:**
> **DelegatingFilterProxy**: The bridge between the Servlet container (Tomcat) and Spring's ApplicationContext is the `DelegatingFilterProxy`. It delegates standard servlet requests to Spring beans (the SecurityFilterChain).


---

**Q: Securing Actuator Endpoints**
> "Actuator endpoints (like `/heapdump` or `/env`) expose sensitive data.
>
> 1.  **Exclude by default**: `management.endpoints.web.exposure.include=health,info` (Only safe ones).
> 2.  **Secure via Security Config**:
>     `http.requestMatchers("/actuator/**").hasRole("ADMIN")`
>     Never leave Actuator open to the public Internet."

**Indepth:**
> **Information Leakage**: A `heapdump` endpoint exposed publicly allows attackers to download your entire memory, extracting passwords, API keys, and customer data. This is a catastrophic vulnerability.


---

**Q: JWT (JSON Web Token) Implementation**
> "Spring Security doesn't have a default 'Generate JWT' button. You implement it:
>
> 1.  **Login**: User sends User/Pass. You verify it.
> 2.  **Generate**: detailed library (jjwt/nimbus) to sign a Token containing claims (User ID, Role, Expiry). Return it to client.
> 3.  **Validate**: Add a `JwtFilter` that runs before every request. It parses the header `Authorization: Bearer <token>`, verifies the signature, and sets the `SecurityContext` if valid."

**Indepth:**
> **Statelessness**: JWTs make the server stateless. You don't need a session store (Redis). The token itself contains the user data. The trade-off is revocation: you can't easily ban a user until their token expires.


---

**Q: CORS (Cross-Origin Resource Sharing)**
> "If your Frontend (React on port 3000) tries to call Backend (Boot on port 8080), browsers block it.
>
> Global Fix in Spring:
> ```java
> @Bean
> WebMvcConfigurer corsConfigurer() {
>     return new WebMvcConfigurer() {
>         public void addCorsMappings(Registry registry) {
>             registry.addMapping("/**").allowedOrigins("http://localhost:3000");
>         }
>     };
> }
> ```
> In Spring Security, you must also enable `http.cors()` in the security filter chain."

**Indepth:**
> **Pre-flight**: The browser sends an `OPTIONS` request first (Pre-flight) to check if the cross-origin call is allowed. If your server doesn't handle `OPTIONS`, the real request never happens.


---

**Q: CSRF (Cross-Site Request Forgery)**
> "CSRF attacks trick a user's browser into sending a request (like 'Transfer Money') while they are logged in.
>
> **Stateful Apps (Session)**: ENABLE CSRF. Spring does this by default. It expects a CSRF token with every POST/PUT.
>
> **Stateless Apps (JWT/REST)**: DISABLE CSRF. Since you don't rely on cookies for auth, the browser can't be tricked in the same way. `http.csrf().disable()`."

**Indepth:**
> **SameSite**: Modern browsers execute strict `SameSite` cookie policies, which partially mitigates CSRF. However, explicit token validation is still the defense-in-depth standard for session-based apps.


---

**Q: Method Level Security**
> "To secure specific service methods (not just URLs):
> 1.  Add `@EnableMethodSecurity` to your main config.
> 2.  Annotate methods:
>     `@PreAuthorize("hasAuthority('WRITE_PRIVILEGE')")`
>     `public void deleteDatabase() { ... }`
>
> This is critical for 'Defense in Depth'. Even if a hacker bypasses the URL check, the service layer stops them."

**Indepth:**
> **AOP**: Method security is implemented using AOP (Aspect Oriented Programming) proxies. The check happens *before* the method body executes. If auth fails, an `AccessDeniedException` is thrown.


---

**Q: API Key Authentication**
> "For machine-to-machine communication where you don't need a user login.
>
> 1.  Create a `OncePerRequestFilter`.
> 2.  Check for a header: `X-API-KEY`.
> 3.  Compare it against a stored value (DB or Properties).
> 4.  If valid, manually create an `Authentication` object (like `ApiKeyAuthenticationToken`) and push it to the context."

**Indepth:**
> **M2M**: API Keys are often long-lived. For better security, rotate them periodically. In high-security systems, use mTLS (Mutual TLS) where the client presents a certificate instead of a header string.


## From 35 Spring Boot REST CLI MongoDB
# 35. Advanced REST, CLI & MongoDB

**Q: Rate Limiting per IP**
> "You need a 'Bucket' for every IP address.
> Using `Bucket4j`:
> 1.  Create a Map: `Map<String, Bucket> cache`. Key is the IP.
> 2.  In a Filter, extract IP: `request.getRemoteAddr()`.
> 3.  Get or Create the bucket for that IP.
> 4.  Call `bucket.tryConsume(1)`. If false, return 429."

**Indepth:**
> **Distributed**: Distributed Rate Limiting. In K8s with 10 pods, local memory rate limiting allows 10x the traffic. You must use a centralized store (Redis/Hazelcast) with Lua scripts to ensure atomic token consumption across the cluster.


---

**Q: @ControllerAdvice vs @ExceptionHandler**
> "**@ExceptionHandler** handles exceptions for **one specific Controller**.
>
> "**@ControllerAdvice** is global. It wraps **all Controllers**.
> Always use `@ControllerAdvice` so you have a single, central place for error handling logic (like converting `UserNotFoundException` to a 404 JSON response)."

**Indepth:**
> **Response Body**: `ResponseBodyAdvice`. You can also implement `ResponseBodyAdvice` in a `@ControllerAdvice` class to intercept and modify the *return body* of every controller (e.g., wrapping every response in a standardized `{ "data": ..., "status": "success" }` envelope).


---

**Q: Native SQL in Spring Data JPA**
> "Sometimes HQL/JPQL is too restrictive.
> Use `value` and `nativeQuery = true`.
>
> ```java
> @Query(value = "SELECT * FROM users WHERE email = ?1", nativeQuery = true)
> User findByEmail(String email);
> ```
> Use this sparingly because it ties you to a specific database (Postgres/MySQL) and breaks portability."

**Indepth:**
> **Projections**: You don't have to return the Entity. You can return an Interface (`public interface UserSummary { String getName(); }`). Spring Data JPA's Native Query mapping is smart enough to map the result set columns to the interface getters.


---

**Q: Pagination (Slice vs Page)**
> "**Page<T>** executes two queries:
> 1.  Select the data (`LIMIT 10`).
> 2.  Count total rows (`COUNT(*)`).
>
> "**Slice<T>** executes only **one** query (`LIMIT 11`).
> If it gets 11 rows, it knows there is a 'Next Page'. It doesn't calculate the total pages. This is much faster for large datasets where you only need 'Load More' buttons."

**Indepth:**
> **Keyset Pagination**: Keyset Pagination (Seek Method) is faster than Offset Pagination (`LIMIT 10 OFFSET 1000000`) for deep scrolling because it uses the index (`WHERE id > last_seen_id LIMIT 10`) instead of scanning and discarding rows.


---

**Q: Login Throttling (Brute Force Protection)**
> "Spring Security doesn't do this out of the box. You have to implement it.
>
> On `AuthenticationFailureEvent`:
> 1.  Increment a counter in Redis/DB for that Username/IP.
> 2.  If counter > 5, lock the account for 15 minutes.
>
> On `AuthenticationSuccessEvent`:
> 1.  Reset the counter."

**Indepth:**
> **Soft Lock**: Simply blocking IP is risky (CGNAT shares IPs). A better "Soft Lock" strategy is to require a ReCaptcha challenge after 3 failed attempts instead of a hard lockout.


---

**Q: Stateless vs Stateful Authentication**
> "**Stateful (Session)**: Server keeps a `SessionID` in memory map. Client sends `JSESSIONID` cookie.
> *   Pros: Easy logout (just delete session).
> *   Cons: Hard to scale horizontally (need Sticky Sessions or Redis Session Store).
>
> "**Stateless (JWT)**: Server keeps **nothing**. Client sends a signed `Token`.
> *   Pros: Instantly scalable. Server doesn't care which node handles the request.
> *   Cons: Hard to logout (cannot invalidate a token until it expires)."

**Indepth:**
> **Size**: JWTs grow linearly with claims. If you put too much data (permissions, user profile) in the token, you hit HTTP Header size limits (usually 8KB). Keep JWTs small (just UserID + Roles).


---

**Q: Live Reload (DevTools)**
> "DevTools runs a tiny LiveReload Server in your app.
> You install the **LiveReload Browser Extension**.
>
> When you compile your Java code, Spring restarts (fast).
> When you edit HTML/CSS, DevTools triggers the browser extension to refresh the page automatically. It saves you from hitting F5 a thousand times."

**Indepth:**
> **State**: `LiveReload` doesn't work well with shared state in static variables (because the classloader resets, but system classloader generic statics might not). It also consumes more memory in Dev mode.


---

**Q: Spring Boot CLI**
> "It's a command-line tool that lets you run Groovy scripts as Spring Boot apps.
>
> File `app.groovy`:
> ```groovy
> @RestController
> class App {
>     @GetMapping("/")
>     def home() { "Hello" }
> }
> ```
> Run: `spring run app.groovy`.
> It automatically imports dependencies and starts Tomcat. Great for quick prototyping."

**Indepth:**
> **POCs**: It's rarely used for production apps. It's primarily for quick Proof of Concepts (POCs) or scripting server-side tasks where you need the power of the Spring ecosystem without the boilerplate of Maven/Gradle.


---

**Q: MongoDB Query Methods**
> "Spring Data MongoDB works just like JPA.
>
> `interface UserRepo extends MongoRepository<User, String>`
>
> You can write:
> `List<User> findByLastNameAndAgeGreaterThan(String name, int age);`
>
> Spring automatically translates this into a MongoDB JSON query:
> `{ "lastName": name, "age": { "$gt": age } }`."

**Indepth:**
> **JSON Query**: You can write raw JSON queries: `@Query("{ 'age' : { $gt: ?0 } }")`. This gives you access to specific Mongo operators (`$elemMatch`, `$regex`) that method naming conventions can't express effortlessly.


## From 37 Spring Boot Internals Testing
# 37. Spring Boot Internals & Testing Strategies

**Q: Spring Boot Startup Process (Internals)**
> "When you call `SpringApplication.run()`:
> 1.  It starts a **StopWatch** to track time.
> 2.  It prepares the **Environment** (reads properties, profiles).
> 3.  It prints the **Banner**.
> 4.  It starts the **IoC Container** (ApplicationContext).
> 5.  It triggers **Auto-Configuration** (scanning classpath).
> 6.  It calls any `CommandLineRunners`.
> It’s a carefully choreographed sequence of events."

**Indepth:**
> **Ready Event**: `ApplicationReadyEvent`. The last step. If you want to engage the user (send an email "Server Started"), listen for this event. It means everything is fully up and running.


---

**Q: Auto-Configuration Mechanism**
> "How does Spring know to configure H2?
> It uses `@Conditional` annotations.
>
> It looks at `spring.factories` (or `imports` file in Boot 3).
> It finds `H2AutoConfiguration`.
> It checks: `@ConditionalOnClass(H2.class)`.
> If the H2 JAR is on the classpath, the condition passes, and the bean is created. If not, it's ignored."

**Indepth:**
> **Boot 3**: The "Starter" pattern relies on the `META-INF` files. In Boot 2.7+, `spring.factories` is deprecated for auto-config; use `META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports`.


---

**Q: Disabling Specific Auto-Configuration**
> "Sometimes Spring tries to be too smart. For example, configuring a DataSource when you don't even want one yet.
>
> You can exclude it:
> `@SpringBootApplication(exclude = { DataSourceAutoConfiguration.class })`
>
> This tells Boot: 'I know you see the driver, but don't touch it. I'll handle it manually'."

**Indepth:**
> **Properties**: You can also exclude via properties: `spring.autoconfigure.exclude=org.spring...DataSourceAutoConfiguration`. This is useful for "testing" profiles or when debugging weird conflicts.


---

**Q: Custom Banner**
> "It's a fun feature.
> Just drop a `banner.txt` file in `src/main/resources`.
> You can use ASCII art generators. You can even use placeholders like `${spring-boot.version}` inside the text file to print the version number dynamically on startup."

**Indepth:**
> **CI/CD**: You can turn it off (`spring.main.banner-mode=off`) to speed up startup logs slightly and reduce noise in CI (Continuous Integration) pipelines.


---

**Q: Logging Configuration (Logback)**
> "Spring Boot uses **Logback** by default.
> You don't need a config file for simple changes.
> `logging.level.org.springframework=DEBUG` in `application.properties` is enough.
>
> For complex stuff (like rotating files daily, producing JSON logs for Splunk), create a `logback-spring.xml` in resources. Spring will pick it up automatically."

**Indepth:**
> **Profiles**: `<springProfile name="prod">`. Inside `logback-spring.xml`, you can nest config blocks. "If profile is dev, print to Console. If profile is prod, print to File and Logstash appender."


---

**Q: REST Controller vs Controller**
> "**@Controller** is for standard Spring MVC. It typically returns a String (view name like `index.jsp`).
>
> "**@RestController** is a convenience annotation. It combines `@Controller` and `@ResponseBody`.
> It tells Spring: 'Whatever I return from methods, write it directly to the HTTP Response body as JSON'. You don't need to annotate every method with `@ResponseBody`."

**Indepth:**
> **Under the Hood**: `@ResponseBody` works by using `HttpMessageConverters`. If the return type is String, it uses `StringHttpMessageConverter`. If Object, `MappingJackson2HttpMessageConverter`.


---

**Q: Content Negotiation**
> "One URL, multiple formats.
> The Client asks: `Accept: application/json` -> Server returns JSON.
> The Client asks: `Accept: application/xml` -> Server returns XML.
>
> You don't change code. You just add the `jackson-dataformat-xml` dependency, and Spring automatically supports both formats based on the header."

**Indepth:**
> **Parameters**: `format=json`. You can also configure Spring to look at a query param `?format=xml` instead of the Header. This is easier for testing in browsers.


---

**Q: Testing: @MockBean**
> "When unit testing a Service, you don't want to hit the real Database.
>
> `@MockBean` is used in Spring tests to replace a real bean with a Mockito mock.
>
> ```java
> @MockBean
> private UserRepository userRepo;
>
> // In test
> given(userRepo.findById(1)).willReturn(Optional.of(mockUser));
> ```
> This bypasses the actual database logic entirely."

**Indepth:**
> **Context Caching**: Spring Test caches the context. If you use `@MockBean`, it modifies the context (swaps a bean). This "dirties" the context, forcing Spring to reload a fresh context for the next test class. Too many `@MockBean`s slow down your test suite massively.


---

**Q: Testing: @DataJpaTest**
> "This is a 'Slice Test'.
> It creates a Spring Context, but **only** loads JPA components (Entities, Repositories). It does **not** load Controllers or Services.
>
> It also automatically configures an in-memory DB (H2).
> It's perfect for testing if your custom JPQL queries actually work without starting the whole heavy application."

**Indepth:**
> **Rollback**: By default, every test method is `@Transactional` and rolls back at the end. Your data is clean for the next test. If you want to see data in the DB for debugging, use `@Commit`.


---

**Q: Profile-Specific Testing**
> "You should never run tests against your Production database config.
>
> Annotate your test class:
> `@ActiveProfiles("test")`
>
> Then create `application-test.yml` with H2 or TestContainer settings. Spring will switch to this config during the test run, keeping your data safe."

**Indepth:**
> **Inline**: `@TestPropertySource(properties = "app.feature=false")` has higher precedence than `application-test.yml`. It's great for overriding one specific flag for one specific test case.


## From 38 Spring Boot Deployment Security JPA
# 38. Spring Boot Deployment, Security & JPA

**Q: Spring Boot Profiles**
> "Profiles let you separate configuration for different environments.
> You don't want to connect to `localhost:3306` in Production.
>
> You define `application-dev.yml` and `application-prod.yml`.
> You activate one at runtime: `-Dspring.profiles.active=prod`.
> Spring automatically loads the correct file and ignores the others."

**Indepth:**
> **Multi-Doc**: `---`. You can put multiple profiles in a single `application.yml` separated by three dashes. This is useful for keeping "default" vs "dev" configs in one file while still separating them logically.


---

**Q: Graceful Shutdown**
> "In the past, if you killed the app, active requests just failed immediately.
>
> Now, you enable `server.shutdown=graceful`.
> When you send a *SIGTERM* (kill signal), Spring Boot stops accepting *new* requests but waits (e.g., 30s) for *active* requests to finish processing before actually shutting down. It prevents user errors during deployments."

**Indepth:**
> **Kubernetes**: `preStop` hook. In K8s, when a pod is terminated, it stops receiving traffic. But there's a race condition where the Load Balancer might still send a request. A graceful shutdown + a small `sleep` in the preStop hook ensures zero-downtime deployments.


---

**Q: Method Level Security (@PreAuthorize)**
> "Instead of securing URLs in a config file (which can get messy), you secure the Java methods directly.
>
> `@PreAuthorize("hasRole('ADMIN')")`
> `public void deleteUser(String id) { ... }`
>
> This is cleaner and safer. Even if you forget to secure the URL, the service method itself is protected."

**Indepth:**
> **PostAuthorize**: Less common but powerful. It runs *after* the method. You can check the *return value*. `@PostAuthorize("returnObject.owner == authentication.name")`. Use this to ensure a user can only read their *own* data.


---

**Q: Custom Authentication Provider**
> "If you aren't using standard Database or LDAP auth—maybe you verify users against a Legacy Mainframe or a 3rd Party API.
>
> You implement `AuthenticationProvider`.
> *   `authenticate()`: You write the logic to call the Mainframe.
> *   If success, return a valid `UsernamePasswordAuthenticationToken`.
> *   Spring Security plugs this specific logic into its general login flow."

**Indepth:**
> **Supports**: `supports()`. The `AuthenticationManager` loops through all providers. Your custom provider must implement `supports(Class<?> authentication)` to tell Spring "I know how to handle this specific type of token/login".


---

**Q: JPA N+1 Problem**
> "This is the most common performance killer.
> You fetch 10 Departments. Then for *each* Department, you fetch its Employees.
> *   1 Query for Departments.
> *   10 Queries for Employees.
> Total 11 queries.
>
> **Fix**: Use `JOIN FETCH` in JPQL: `SELECT d FROM Department d JOIN FETCH d.employees`. This fetches everything in **one** single query."

**Indepth:**
> **Statistics**: `Hibernate Statistics`. Enable `spring.jpa.properties.hibernate.generate_statistics=true` in tests. It prints "Session Metrics" at the end, showing exactly how many JDBC statements were executed. Ideally, it should be 1, not N+1.


---

**Q: Optimistic vs Pessimistic Locking**
> "**Optimistic**: You assume conflicts are rare. You add a `@Version` column. If two people save at the same time, the second one fails with `OptimisticLockException`.
>
> "**Pessimistic**: You assume conflicts are frequent. You lock the database row (`SELECT ... FOR UPDATE`). No one else can even *read* it until you are done. It's safer but slower."

**Indepth:**
> **Deadlocks**: Pessimistic Locking can cause deadlocks if two transactions lock resources in different orders. Always order your locks consistently (e.g., sort by ID before locking) to prevent this.


---

**Q: Cascade Types**
> "Cascading means: 'What happens to the Child when I do something to the Parent?'
>
> *   `CascadeType.PERSIST`: If I save the Parent, save the Children too.
> *   `CascadeType.REMOVE`: If I delete the Parent, delete all Children (orphans).
> *   `CascadeType.ALL`: Do everything."

**Indepth:**
> **Orphans**: `orphanRemoval=true`. This is different from Cascade DELETE. If you remove a child from the parent's list (`parent.getChildren().remove(child)`), `orphanRemoval` deletes the child from the DB. Cascade DELETE only works if you delete the *Parent*.


---

**Q: Flyway / Liquibase**
> "Never let Hibernate auto-generate your Production schema (`ddl-auto=update`). It's dangerous.
>
> Use **Flyway**.
> You write SQL scripts: `V1__init.sql`, `V2__add_column.sql`.
> On startup, Flyway checks the DB version. If it's at V1, it runs V2.
> This ensures your Database structure is always in sync with your Java code."

**Indepth:**
> **Baseline**: `baselineOnMigrate`. If you introduce Flyway to an existing populated database, you need to "baseline" it. This tells Flyway "Assume the DB is already at version 1, start running scripts from V2".


---

**Q: Logging SQL**
> "To see what Hibernate is doing:
> `spring.jpa.show-sql=true` (Basic).
>
> To see the *values* inside the `?`:
> `logging.level.org.hibernate.type.descriptor.sql=TRACE`.
>
> This is vital for debugging N+1 problems or wrong parameter bindings."

**Indepth:**
> **P6Spy**: `P6Spy` is a library that wraps the JDBC driver. It logs the *actual* SQL sent to the DB with all `?` replaced by real values, plus execution time. It is much better than Hibernate's built-in logging.


---

**Q: @Entity vs @Table**
> "**@Entity**: Makes it a JPA Object. The class name becomes the Table name by default.
>
> "**@Table**: Allows you to customize the DB table details.
> `@Table(name = "tbl_users", indexes = @Index(...))`
> You use `@Table` when the DB name doesn't match your Java class name."

**Indepth:**
> **Constraints**: Unique Constraints. You can define multi-column unique constraints inside `@Table(uniqueConstraints = @UniqueConstraint(columnNames = {"firstName", "lastName"}))`.


## From 39 Spring Boot Security Testing JPA Revision
# 39. Spring Boot (Security, Testing & JPA Deep Dive)

**Q: Basic Auth vs JWT**
> "**Basic Auth**: The browser sends the username/password with *every single request*. It’s simple but risky (if not using HTTPS) and requires the server to validate it every time.
>
> "**JWT (JSON Web Token)**: The user logs in *once*. The server gives them a signed 'Badge' (Token).
> For future requests, they just show the Badge. The server sees the signature is valid and lets them in. No database check needed."

**Indepth:**
> **Revocation**: The main downside of JWT is that you cannot "force logout" a user instantly properly without a blacklist (Redis). Basic Auth / Sessions are easier to invalidate server-side (just delete the session).


---

**Q: @WithMockUser (Testing Security)**
> "Testing guarded endpoints is annoying because you have to 'login' in your test code.
>
> `@WithMockUser(username="admin", roles={"ADMIN"})`
>
> This annotation automatically creates a fake 'Logged In' context for that test method. You don't need to actually send an Authorization header. It just assumes the user is already authenticated."

**Indepth:**
> **Custom User**: `@WithUserDetails` is better if your app relies on custom fields in your `User` object (like `tenantId`). It actually loads the user from your Test `UserDetailsService` rather than creating a fake mock.


---

**Q: Unit vs Integration vs E2E Tests**
> "**Unit Test**: Testing **one class** in isolation. Mock everything else. Fast (milliseconds).
>
> "**Integration Test**: Testing how **two things talk** (e.g., Service + DB, or Controller + Service). Slower (seconds).
>
> "**E2E (End-to-End)**: Testing the **entire flow**. Spin up the whole app, hit the API, check the DB. Slowest (minutes)."

**Indepth:**
> **Testcontainers**: For Integration/E2E tests, *always* use Testcontainers. Using H2 for integration tests is dangerous because H2 behaves differently than Postgres (syntax, locking, constraints).


---

**Q: @SpringBootTest**
> "This is the 'Big Gun' of testing.
> It starts your **entire** application. It finds your main class, reads `application.properties`, connects to the (test) database, and initializes all beans.
>
> Use this when you want to be 100% sure the application wires together correctly, but use it sparingly because it's slow."

**Indepth:**
> **Random Port**: `RANDOM_PORT`. Always use `webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT`. This allows running tests in parallel without port conflicts.


---

**Q: Spring Boot Admin**
> "Managing 50 microservices is a nightmare.
> **Spring Boot Admin** is a UI Dashboard.
>
> All your microservices (Clients) register with the Admin Server.
> The UI shows you a traffic light view: 'Service A is DOWN', 'Service B is running out of memory'. You can also view logs and change log levels directly from the browser."

**Indepth:**
> **Security**: The Admin server itself must be secured (login), otherwise anyone can view your env vars and heap dumps. The Client must also authenticate to register itself.


---

**Q: Entity Lifecycle States (JPA)**
> "An Entity can be in 4 states:
> 1.  **Transient**: Just created (`new User()`). Hibernate doesn't know about it.
> 2.  **Persistent**: Currently managed by Hibernate (e.g., just saved or fetched). Changes are auto-saved.
> 3.  **Detached**: Session is closed. Java object exists, but changes won't be saved to DB.
> 4.  **Removed**: Scheduled for deletion."

**Indepth:**
> **Merge**: `merge()` vs `persist()`. `persist` takes a transient instance. `merge` takes a detached instance and copies its state to a managed instance. It returns the *new* managed instance; it doesn't make the passed object managed.


---

**Q: Dirty Checking**
> "In Hibernate, you rarely call `save()` for updates.
> You just fetch an object, change a field (`user.setName("New Name")`), and walk away.
>
> When the Transaction commits, Hibernate compares the object with its original snapshot. If it sees a change, it **automatically** generates specific `UPDATE` SQL. This is Dirty Checking."

**Indepth:**
> **Optimization**: If you mark a transaction `@Transactional(readOnly = true)`, Spring disables dirty checking for performance. Hibernate won't snapshot the objects, saving memory and CPU.


---

**Q: L1 vs L2 Cache**
> "**L1 Cache (Session)**: Enabled by default. If you ask for User #1 twice in the *same* transaction, Hibernate returns the same object instantly without hitting the DB twice.
>
> "**L2 Cache (SessionFactory)**: Shared across *all* users. Requires a provider (EhCache/Redis). If User A fetches 'Country List', User B gets it from the cache. rarely used in Microservices (we prefer Redis at the app layer)."

**Indepth:**
> **Stale Data**: L2 Cache is tricky in distributed systems. If Instance A updates a user, Instance B's cache might be stale. You need a distributed cache invalidation mechanism (like Redis Pub/Sub or Hibernate Clustered Cache).


---

**Q: JPA Projections**
> "Don't fetch the whole Entity if you only need the name.
>
> Define an Interface:
> ```java
> interface UserNameOnly {
>     String getFirstName();
>     String getLastName();
> }
> ```
> `List<UserNameOnly> findByAge(int age);`
>
> Spring generates an efficient SQL query: `SELECT first_name, last_name FROM...` instead of `SELECT *`."

**Indepth:**
> **Class-based**: You can also use a Java Record/Class with a constructor matching the fields. `Select new com.example.UserDto(u.name, u.age)...`. This is type-safe and cleaner than Interfaces.


---

**Q: Auditing (@CreatedDate)**
> "Stop manually setting `createdAt` and `updatedAt`.
> 1.  Add `@EnableJpaAuditing`.
> 2.  Add `@EntityListeners(AuditingEntityListener.class)` to your entity.
> 3.  Annotate fields: `@CreatedDate`, `@LastModifiedDate`, `@CreatedBy`.
>
> Spring automatically fills these timestamps whenever you save or update."

**Indepth:**
> **AuditorAware**: How does Spring know *who* logged in? You implement `AuditorAware<String>`. It fetches the current user from `SecurityContextHolder` and returns the username to be stamped on the record.


## From 40 Spring MVC Security WebFlux Revision
# 40. Spring MVC, Security & WebFlux (Final Revision)

**Q: DispatcherServlet (Front Controller Pattern)**
> "Imagine a hotel reception. You don't walk directly to the chef to order food. You go to the front desk.
>
> **DispatcherServlet** is that Front Desk.
> Every single HTTP request (`/login`, `/users`, `/home`) hits this servlet first.
> It checks the URL, looks up the 'Handler Mapping' to find the right Controller, and delegates the work. You never interact with it directly, but it runs the show."

**Indepth:**
> **Context Hierarchy**: `WebApplicationContext`. The DispatcherServlet creates its own child context (containing Controllers, ViewResolvers) which inherits from the Root WebApplicationContext (containing Services, Repositories). This separation allows multiple DispatcherServlets to share common beans.


---

**Q: @PathVariable vs @RequestParam**
> "**@PathVariable**: It's part of the **Identity** of the resource.
> `/users/123` -> 123 identifies a specific user.
>
> "**@RequestParam**: It's for **Filtering** or Sorting.
> `/users?country=US&sort=age` -> You are looking at the users resource, but filtering it.
> Use PathVariable for IDs, RequestParam for options."

**Indepth:**
> **Encoding**: URL Encoding. Path variables are part of the URL structure and must be URL-encoded if they contain special characters. Request params are standard query strings. Spring decodes them automatically, but clients must send them correctly.


---

**Q: Spring Security Filter Chain**
> "Security doesn't happen in the Controller. It happens at the door.
>
> Spring Security is a chain of 10-15 filters that sit *before* the DispatcherServlet.
> *   **JwtAuthenticationFilter**: 'Do you have a token?'
> *   **UsernamePasswordFilter**: 'Are you logging in?'
> *   **AuthorizationFilter**: 'Are you allowed to see this?'
>
> If you pass all filters, you get to the Controller. If one fails, you get thrown out (401/403)."

**Indepth:**
> **Debugging**: `logging.level.org.springframework.security=DEBUG`. This is the single most useful config for debugging 403s. It prints the execution of every filter in the chain so you can see exactly which one denied the request and why (e.g., "CsrfFilter denied request").


---

**Q: Authentication vs Authorization**
> "**Authentication (Who are you?)**: 'I am John.' Checks username/password.
>
> "**Authorization (What can you do?)**: 'John is an Admin.' Checks roles and permissions.
>
> You must authenticate *before* you can be authorized."

**Indepth:**
> **HTTP Codes**: 401 Unauthorized actually means "Unauthenticated" (I don't know who you are). 403 Forbidden means "Unauthorized" (I know who you are, but you can't do this). The naming is confusing but standard.


---

**Q: OAuth2 Flow (Simple Explanation)**
> "Think of 'Login with Google'.
> 1.  You click the button.
> 2.  You are redirected to Google's server.
> 3.  You sign in there. Google asks: 'Allow this app to see your email?'
> 4.  You say Yes. Google sends a 'Code' back to your App.
> 5.  Your App talks to Google silently: 'Here is the Code, give me an Access Token.'
> 6.  Now your app has a Token to fetch your email from Google."

**Indepth:**
> **PKCE**: Proof Key for Code Exchange. In modern mobile/SPA apps, the "Code" can be intercepted. PKCE adds a cryptographic hash/verifier to simpler flows to prevent authorization code injection attacks.


---

**Q: CSRF (Cross-Site Request Forgery)**
> "If you log into your bank, and then visit `evil.com`.
> `evil.com` tries to send a hidden form POST to `bank.com/transfer`.
> Since your browser automatically sends the Bank Cookies, the Bank thinks *you* did it.
>
> **The Fix**: The Bank expects a secret `CSRF-Token` in the form. `evil.com` doesn't know this token, so the request fails.
> We disable this for REST APIs because we use Headers (Authorization), not Cookies."

**Indepth:**
> **Safe Methods**: GET, HEAD, OPTIONS are considered "Safe" (Read-only). CSRF only protects unsafe methods (POST, PUT, DELETE). Browsers execute Safe methods without restrictions, so never change state (DB writes) in a GET request.


---

**Q: Reactive Programming (Backpressure)**
> "In traditional systems, if the Producer sends data too fast, the Consumer crashes (Out of Memory).
>
> **Backpressure** is the Consumer saying: 'I am overwhelmed! Stop sending!' or 'Send me only 5 items'.
> It allows the system to handle massive load gracefully without crashing."

**Indepth:**
> **Strategies**: `onBackpressureBuffer/Drop`. What if the consumer *cannot* keep up? You can choose to Buffer the extra items (risking OOM), Drop them (data loss), or Error out. Reactive streams force you to decide this failure mode upfront.


---

**Q: WebFlux vs MVC (Threading)**
> "**MVC**: One Thread per Request. Blocking. If you have 200 threads and 201 concurrent users, the last one waits.
>
> "**WebFlux**: Event Loop. Non-Blocking. One thread handles many requests. When a request waits for DB, the thread serves someone else. It scales to 10,000+ concurrent connections with very little hardware."

**Indepth:**
> **Context Switching**: WebFlux reduces context switching because threads don't block. However, cpu-bound tasks (heavy calculation) can freeze the Event Loop, stopping *all* requests. You must offload CPU-heavy work to a separate standard thread pool.


---

**Q: R2DBC**
> "JDBC is blocking. If you use JDBC in WebFlux, you kill the performance benefits.
>
> **R2DBC** (Reactive Relational Database Connectivity) is the new driver standard for SQL databases (Postgres, MySQL) that is fully non-blocking/reactive. It's strictly required for true WebFlux apps talking to SQL."

**Indepth:**
> **Maturity**: R2DBC is newer than JDBC. It lacks some mature features like robust caching, complex mapping, or stored procedure support compared to Hibernate/JPA.


---

**Q: Mono vs Flux**
> "They are the 'futures' of Reactive Java (Project Reactor).
> *   **Mono**: A promise for 0 or 1 result. (e.g., `findById`).
> *   **Flux**: A stream of 0 to N results. (e.g., `findAll` or a live stock ticker).
> You chain operators on them (`.map`, `.filter`) to define the pipeline."

**Indepth:**
> **Multicasting**: `share()`. By default, if two people subscribe to a Flux, the source (e.g., DB query) is executed *twice*. Using `.share()` or `.publish()` multicasts the result to all subscribers, executing the source only once.

