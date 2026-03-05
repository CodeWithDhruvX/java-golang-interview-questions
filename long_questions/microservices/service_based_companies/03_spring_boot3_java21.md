# 🌱 Microservices — Spring Boot 3 & Java 21 Features

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** TCS (digital roles), Infosys (Topaz/digital), Wipro (digital), Cognizant (digital), Capgemini — any company running modernized Java stacks

> **Why this matters:** Spring Boot 3 (released Nov 2022) and Java 21 (LTS, released Sep 2023) are now the standard. Service companies have begun migrating. Interviewers now ask "Are you familiar with Spring Boot 3?" in mid-level interviews.

---

## Q1. What are the major changes in Spring Boot 3 vs Spring Boot 2?

**"Spring Boot 3 is a significant major release with several breaking and enhancing changes:**

**1. Java 17 as Minimum Baseline (Java 21 recommended)**
- Spring Boot 2.x supported Java 8+
- Spring Boot 3.x requires Java 17+ as a minimum
- Full support for Java 21 features (Virtual Threads, Records, Sealed Classes)

**2. Jakarta EE 10 (not javax.* anymore — this breaks old code!)**
- The biggest source of migration pain
- All `javax.*` packages renamed to `jakarta.*`
```java
// Spring Boot 2 — javax
import javax.persistence.Entity;
import javax.validation.constraints.NotNull;
import javax.servlet.http.HttpServletRequest;

// Spring Boot 3 — jakarta
import jakarta.persistence.Entity;
import jakarta.validation.constraints.NotNull;
import jakarta.servlet.http.HttpServletRequest;
```

**3. Native Image Support (GraalVM)**
- Spring Boot 3 is designed for GraalVM AOT compilation
- Compile Spring Boot apps to native binaries: startup in < 100ms, 10× lower memory
```bash
./mvnw -Pnative native:compile
./target/my-app  # Native binary — no JVM needed!
```

**4. Improved Observability (Micrometer Observations API)**
- New `@Observed` annotation auto-creates metrics + traces + spans
```java
@Observed(name = "order.create", contextualName = "creating-order")
public Order createOrder(CreateOrderRequest req) {
    // Micrometer automatically creates:
    // - Timer metric: order.create.duration
    // - Trace span: creating-order (sent to Zipkin/Jaeger via OpenTelemetry)
    return orderRepository.save(new Order(req));
}
```

**5. `@HttpExchange` — declarative HTTP clients (replaces Feign in stdlib)**
- No more external Feign dependency needed
```java
// Declare the client interface
@HttpExchange("http://payment-service")
public interface PaymentClient {
    @PostExchange("/api/payments")
    PaymentResponse charge(@RequestBody PaymentRequest request);

    @GetExchange("/api/payments/{id}")
    PaymentResponse getPayment(@PathVariable String id);
}

// Register it as a bean
@Bean
public PaymentClient paymentClient(WebClient.Builder builder) {
    WebClient client = builder.baseUrl("http://payment-service").build();
    return HttpServiceProxyFactory.builderFor(WebClientAdapter.create(client))
        .build()
        .createClient(PaymentClient.class);
}
```

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Infosys (Topaz team), TCS (iON digital), Wipro — hiring managers ask "What's new in Spring Boot 3?" to gauge how up-to-date the candidate is.

---

## Q2. What are Java Virtual Threads (Project Loom)? How do they change microservices?

**"Virtual Threads (Java 21) are the biggest concurrency change in Java's history. They fundamentally change how you write high-throughput server code.**

**The old model — Platform Threads:**
- Each request gets an OS thread
- OS thread = ~1MB stack + OS scheduler overhead
- A typical 8-CPU server handles ~500–1000 concurrent threads before performance degrades
- Blocking I/O (DB queries, HTTP calls) = thread is parked but consuming OS resources

**Virtual Threads:**
- Lightweight threads managed by the JVM, not the OS
- 1MB OS thread can multiplex MILLIONS of virtual threads
- Blocking I/O in a virtual thread = JVM parks it without blocking the OS thread
- Cost per virtual thread: ~a few KB (vs ~1MB for OS thread)

**Code change — it's minimal:**
```java
// Spring Boot 3.2+ — Enable virtual threads globally
# application.yml
spring:
  threads:
    virtual:
      enabled: true  # That's it! Spring will use virtual threads for all requests

# OR manually for specific cases:
Thread virtualThread = Thread.ofVirtual()
    .name("payment-processor")
    .start(() -> {
        // This blocking DB call does NOT block an OS thread
        Payment payment = paymentRepository.findById(id); // Blocks, but cheap!
        processPayment(payment);
    });
```

**Impact on microservices:**
```
Traditional (Platform Threads):
100 concurrent requests → 100 OS threads → ~100MB RAM → OS scheduler under load

Virtual Threads (Java 21):
100,000 concurrent requests → 100,000 virtual threads → ~a few hundred MB → JVM scheduler
→ Same server handles 1000× more concurrent I/O bound requests
```

**When NOT to use virtual threads:**
- CPU-intensive tasks (virtual threads provide no benefit — CPU is the bottleneck)
- Tasks using `synchronized` blocks with blocking I/O (pin to OS thread, reducing benefit)
  → Use `ReentrantLock` instead of `synchronized` in virtual-thread-heavy code"

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Infosys digital, TCS (modern project interviews), Wipro (digital teams). Even if the company hasn't adopted it yet, asking about it shows you're forward-looking — interviewers love that.

#### Indepth
**Virtual Threads vs Reactive (WebFlux/Project Reactor):**
| | Virtual Threads | Reactive (WebFlux) |
|--|-----------------|-------------------|
| Programming model | Imperative (normal Java) | Functional/reactive (complex) |
| Debugging | Normal stack traces | Complex, async stack traces |
| Ecosystem compatibility | All Java libraries work | Must use reactive-compatible libs |
| Performance for I/O bound | Excellent | Excellent |
| Learning curve | Near zero | Very high |
| **Verdict** | **Preferred for new code** | Legacy reactive systems |

---

## Q3. What is `@RestClient` / `@HttpExchange` in Spring Boot 3? Write an example.

**"Spring Boot 3.2 introduced `RestClient` — a modern, fluent, synchronous HTTP client that replaces `RestTemplate` without requiring reactive programming."**

```java
// Method 1: RestClient (synchronous, replaces RestTemplate)
@Service
public class PaymentService {

    private final RestClient restClient;

    public PaymentService(RestClient.Builder builder) {
        this.restClient = builder
            .baseUrl("http://payment-service")
            .defaultHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
            .build();
    }

    public PaymentResponse charge(PaymentRequest request) {
        return restClient.post()
            .uri("/api/payments")
            .body(request)
            .retrieve()
            .onStatus(HttpStatusCode::is4xxClientError, (req, res) -> {
                throw new PaymentException("Payment rejected: " + res.getStatusCode());
            })
            .body(PaymentResponse.class); // Synchronous — returns result directly
    }

    public PaymentResponse getPayment(String paymentId) {
        return restClient.get()
            .uri("/api/payments/{id}", paymentId)
            .retrieve()
            .body(PaymentResponse.class);
    }
}

// Method 2: @HttpExchange (declarative, like Feign — built into Spring 6)
@HttpExchange(url = "http://inventory-service", accept = "application/json")
public interface InventoryClient {

    @GetExchange("/api/inventory/{productId}")
    InventoryStatus getInventory(@PathVariable String productId);

    @PostExchange("/api/inventory/reserve")
    ReservationResponse reserve(@RequestBody ReservationRequest request);

    @DeleteExchange("/api/inventory/reservations/{reservationId}")
    void cancelReservation(@PathVariable String reservationId);
}

// Bean configuration
@Bean
public InventoryClient inventoryClient(RestClient.Builder builder) {
    RestClient restClient = builder.baseUrl("http://inventory-service").build();
    RestClientAdapter adapter = RestClientAdapter.create(restClient);
    HttpServiceProxyFactory factory = HttpServiceProxyFactory.builderFor(adapter).build();
    return factory.createClient(InventoryClient.class);
}
```

**Key difference RestClient vs RestTemplate vs WebClient:**
| | RestTemplate | WebClient | RestClient |
|--|------------|-----------|----------|
| Spring version | Spring 3+ | Spring 5 (WebFlux) | Spring 6 / Boot 3.2 |
| Blocking? | Yes | No (reactive) | Yes |
| API style | Verbose | Functional/reactive | Fluent + type-safe |
| Status | Deprecated (maintenance mode) | Supported but complex | ✅ Recommended for new code |

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS, Infosys — if your resume says "Spring Boot 3", expect this question.

---

## Q4. What are Java Records? How are they used in Spring Boot microservices?

**"Records (Java 16+, standard in 17+) are immutable data classes with auto-generated constructor, getters, `equals()`, `hashCode()`, and `toString()`. They eliminate boilerplate for DTOs."**

```java
// OLD WAY (Spring Boot 2 + Java 8) — lots of boilerplate
public class CreateOrderRequest {
    private final String userId;
    private final List<OrderItem> items;
    private final BigDecimal discount;

    public CreateOrderRequest(String userId, List<OrderItem> items, BigDecimal discount) {
        this.userId = userId;
        this.items = items;
        this.discount = discount;
    }
    // + getters, equals, hashCode, toString — 30+ lines of boilerplate
}

// NEW WAY (Spring Boot 3 + Java 17/21) — Records
public record CreateOrderRequest(
    @NotNull String userId,
    @NotEmpty List<OrderItem> items,
    @DecimalMin("0") BigDecimal discount
) {}

// Used in controllers — works perfectly with Spring's JSON deserialization
@RestController
@RequestMapping("/api/orders")
public class OrderController {

    @PostMapping
    public ResponseEntity<OrderResponse> createOrder(
            @Valid @RequestBody CreateOrderRequest request) {
        // request.userId(), request.items(), request.discount() — auto-generated
        return ResponseEntity.ok(orderService.createOrder(request));
    }
}

// Records for API responses too
public record OrderResponse(
    String orderId,
    String status,
    BigDecimal total,
    LocalDateTime createdAt
) {}
```

**Records with JPA (caveat):** Records cannot be JPA entities (JPA requires mutable state + no-arg constructor). Use Records for DTOs only; use `@Entity` classes for database entities.

---

## Q5. What is GraalVM Native Image? When would you use it for microservices?

**"GraalVM Native Image compiles a Java application to a platform-specific native binary AOT (Ahead of Time) instead of running through the JVM. The result: dramatically faster startup and lower memory.**

**Performance comparison (Spring Boot 3 order-service):**
| | JVM (HotSpot) | GraalVM Native |
|--|--------------|---------------|
| Startup time | 3–6 seconds | 50–200 ms |
| Memory at idle | 200–400 MB | 30–80 MB |
| Peak throughput | Better (JIT optimization) | Slightly lower |
| Build time | Fast (30s) | Slow (3–10 mins) |
| Docker image size | ~200 MB | ~50 MB |

**Build a native image:**
```xml
<!-- pom.xml -->
<build>
  <plugins>
    <plugin>
      <groupId>org.graalvm.buildtools</groupId>
      <artifactId>native-maven-plugin</artifactId>
    </plugin>
  </plugins>
</build>
```
```bash
mvn -Pnative native:compile
./target/order-service # Starts in 80ms!

# Or with Docker (no GraalVM needed on dev machine)
mvn spring-boot:build-image -Pnative
```

**When to use Native Image:**
- **Serverless (AWS Lambda, Google Cloud Functions):** Cold starts go from 8 seconds to 80ms — huge cost and UX win
- **Short-lived batch jobs / CLIs:** No wasted time waiting for JVM warmup
- **Very memory-constrained environments:** IoT, edge computing
- **Many sidecar services:** If you have 200 microservices, cutting per-pod memory from 300MB to 60MB saves significant infrastructure cost

**When NOT to use Native Image:**
- Long-running services with heavy computational load (JIT-compiled JVM eventually beats native throughput)
- When you use libraries that rely on Java reflection heavily and don't support GraalVM hints (check compatibility list)
- When build times in CI/CD are already a bottleneck (native build takes 3–10 minutes)"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Infosys (Topaz), TCS (modern digital projects). Shows you're ahead of the curve. GraalVM Native is becoming industry standard for Lambda-deployed microservices.
