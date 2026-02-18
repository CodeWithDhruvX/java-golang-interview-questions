# 32. Spring Core, Monitoring & WebFlux

**Q: @Component vs @Service vs @Repository**
> "Technically, they are all the same. `@Service` and `@Repository` are just aliases for `@Component`.
>
> However, we use them for **Semantics** and **Exception Translation**:
> *   `@Repository`: Tells Spring 'This class talks to a DB'. It also automatically translates low-level SQL exceptions into cleaner Spring DataAccessExceptions.
> *   `@Service`: Tells devs 'Business Logic lives here'.
> *   `@Component`: Generic utility classes."

**Indepth:**
> **AOP**: `@Service` and `@Repository` can be targeted by Aspect-Oriented Programming pointcuts. For example, `@Transactional` attributes usually apply to the Service layer, while Exception Translation (translating SQL errors to Spring ones) only happens on `@Repository`.


---

**Q: @Autowired vs @Qualifier**
> "**@Autowired** by default looks for a bean by **Type**.
> If you have an interface `PaymentService` and two implementations (`CreditCardService`, `PayPalService`), Spring throws `NoUniqueBeanDefinitionException`.
>
> **@Qualifier** tells Spring to look by **Name**.
> `@Autowired @Qualifier("payPalService")` resolves the ambiguity."

**Indepth:**
> **Primary**: Alternatively, you can simplify injection by using `@Primary` on one of the implementations. This tells Spring "if there's ambiguity and no qualifier is present, use this bean by default."


---

**Q: @Value vs @ConfigurationProperties**
> "**@Value** is for injecting single values.
> `@Value("${app.timeout}") int timeout;`
> Good for quick, one-off properties.
>
> "**@ConfigurationProperties** is for grouping properties.
> It maps a hierarchical structure (`server.port`, `server.address`) to a Java POJO.
> It is type-safe, validates fields, and supports loose binding (kebab-case to camelCase). Always prefer this for complex configs."

**Indepth:**
> **Relaxed Binding**: `@ConfigurationProperties` supports relaxed binding. `my.app-name` in properties matches `myAppName`, `my_app_name`, and `my.app.name` in Java. `@Value` requires exact string matches.


---

**Q: Constructor Injection**
> "Stop using Field Injection (`@Autowired private Repo repo`). It makes testing hard (you have to use Reflection to set the repo).
>
> **Constructor Injection** is the standard.
> ```java
> private final Repo repo;
> public Service(Repo repo) { this.repo = repo; }
> ```
> It forces dependencies to be provided, ensures immutability (`final`), and makes Unit Testing trivial (just pass a mock in the constructor)."

**Indepth:**
> **Circular Dependencies**: Constructor injection prevents circular dependencies at compile-time/start-time (Bean A needs B, B needs A). Spring throws `BeanCurrentlyInCreationException` immediately, forcing you to refactor your bad design.


---

**Q: Bean Scopes (Singleton vs Prototype)**
> "**Singleton** (Default): Spring creates **one** instance of the bean per container. Shared by everyone. Stateless services should be Singletons.
>
> "**Prototype**: Spring creates a **new** instance every time you ask for it (`context.getBean()`). Stateful beans (like a 'ShoppingCart' or 'UserSession') might use this, though `SessionScope` is usually better for web apps."

**Indepth:**
> **Proxy**: If you inject a Prototype bean into a Singleton bean, the Prototype is created *only once* (when the Singleton is created). To get a new Prototype every time, you must use `ObjectFactory<MyPrototype>` or `Lookup` method injection.


---

**Q: Actuator Health Endpoint**
> "`/actuator/health` provides a status check.
> By default, it just says `{"status": "UP"}`.
>
> If you enable details (`management.endpoint.health.show-details=always`), it checks:
> *   Disk Space
> *   Database Connection
> *   Message Broker Connectivity
> If **any** of these down, the overall status becomes `DOWN` (503 Service Unavailable)."

**Indepth:**
> **Custom**: You can write your own `HealthIndicator`. For example, checking if a critical 3rd party API is reachable. Implement the interface and return `Health.up()` or `Health.down().withDetail("reason", "timeout")`.


---

**Q: Micrometer & Prometheus**
> "**Micrometer** is like SLF4J but for metrics.
> It's a facade. You write code against Micrometer (`Counter.increment()`), and it translates that to whatever backend you use (Prometheus, Datadog, NewRelic).
>
> **Prometheus** scrapes these metrics. You expose `/actuator/prometheus`, and Prometheus comes and 'pulls' the data every 15 seconds."

**Indepth:**
> **Dimensionality**: Micrometer supports tags (dimensions). Instead of just `http_requests_total`, you track `http_requests_total{method="GET", status="200"}`. This allows powerful querying like "Show me only 500 errors on POST requests".


---

**Q: Application Monitoring (Memory/CPU)**
> "Use the `/actuator/metrics` endpoint.
> *   `jvm.memory.used`
> *   `system.cpu.usage`
> *   `hikaricp.connections.active`
>
> You don't usually read the JSON manually. You connect Grafana to visualize 'CPU Spikes' or 'Memory Leaks' over time."

**Indepth:**
> **Alerting**: Monitoring is useless without alerting. Set up Prometheus/Grafana alerts for "High Memory Usage (> 85%)" or "High Error Rate (> 1%)". Don't wait for users to complain.


---

**Q: Spring WebFlux vs Spring MVC**
> "**Spring MVC**: Thread-per-request.
> 1 Request = 1 Thread. If the thread waits for DB, it sits idle (Blocked).
> Good for standard CRUD apps.
>
> "**Spring WebFlux**: Event-Loop based (like Node.js).
> Small number of threads handle thousands of concurrent requests. If waiting for DB, the thread goes to work on another request.
> Returns **Mono** (0-1 item) or **Flux** (0-N items).
> Good for High-Scale Streaming apps."

**Indepth:**
> **Backpressure**: WebFlux supports **Backpressure**. If the client (Consumer) is slow, it tells the Server (Producer) to slow down ("I can only handle 5 items right now"). Spring MVC just overwhelms the client.


---

**Q: Mono vs Flux**
> "In Reactive Programming, we don't return `List<User>`.
>
> *   **Mono<T>**: A wrapper for zero or one item. 'I promise to give you a User (or error) in the future'.
> *   **Flux<T>**: A wrapper for zero to N items. 'I will stream Users to you as they arrive'.
>
> You subscribe to them to get the data."

**Indepth:**
> **Cold vs Hot**: `Flux` is "Cold" by default. Nothing happens until you subscribe. If you have a DB call in a Flux but nobody subscribes, the DB call never executes.

