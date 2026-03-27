# Virtusa Java Interview – Part 3: Scenario-Based & Advanced Topics (Q&A with Code)

---

## Q1. You implemented parallel processing using CompletableFuture — but threads are blocking. What's your fix?

**Spoken Answer:**
"The most common cause is calling `get()` or `join()` inside an async chain — that blocks a thread from the ForkJoinPool, starving other tasks.

The fix is to use **non-blocking composition**: `thenApply()`, `thenCompose()`, `thenCombine()`. I also supply a **dedicated ExecutorService** for IO-bound tasks so they don't eat up CPU pool threads. I only call `get()` once, at the very boundary where I need the final result."

```java
import java.util.concurrent.*;

public class CompletableFutureDemo {

    static final ExecutorService IO_POOL  = Executors.newFixedThreadPool(10);
    static final ExecutorService CPU_POOL = Executors.newFixedThreadPool(
            Runtime.getRuntime().availableProcessors());

    // ❌ Bad: get() inside async chain blocks ForkJoinPool thread
    static void badApproach() throws Exception {
        CompletableFuture<String> f = CompletableFuture.supplyAsync(() -> "data");
        String result = f.get(); // ❌ blocks calling thread
    }

    // ✅ Good: non-blocking chain
    static CompletableFuture<String> fetchUser(int id) {
        return CompletableFuture.supplyAsync(() -> "User:" + id, IO_POOL);
    }

    static CompletableFuture<String> fetchOrders(String user) {
        return CompletableFuture.supplyAsync(() -> "Orders-for-" + user, IO_POOL);
    }

    public static void main(String[] args) throws Exception {

        // ✅ thenCompose: sequential async (flatMap equivalent)
        CompletableFuture<String> result = fetchUser(42)
                .thenComposeAsync(user -> fetchOrders(user), IO_POOL)
                .thenApplyAsync(String::toUpperCase, CPU_POOL)
                .exceptionally(ex -> "Error: " + ex.getMessage());

        System.out.println(result.get(5, TimeUnit.SECONDS)); // only get() here at the end

        // ✅ thenCombine: two parallel tasks merged
        CompletableFuture<String> userCf   = fetchUser(1);
        CompletableFuture<String> orderCf  = fetchOrders("User:1");
        String combined = userCf.thenCombine(orderCf, (u, o) -> u + " | " + o)
                               .get(5, TimeUnit.SECONDS);
        System.out.println("Combined: " + combined);

        // ✅ allOf: wait for multiple parallel futures
        var f1 = fetchUser(1); var f2 = fetchUser(2); var f3 = fetchUser(3);
        CompletableFuture.allOf(f1, f2, f3)
                .thenRun(() -> {
                    try { System.out.println("All: " + f1.get() + ", " + f2.get() + ", " + f3.get()); }
                    catch (Exception e) { e.printStackTrace(); }
                }).get();

        IO_POOL.shutdown(); CPU_POOL.shutdown();
    }
}
```

---

## Q2. How do you safely shutdown a thread pool in production without losing tasks?

**Spoken Answer:**
"I call `shutdown()` — not `shutdownNow()`. `shutdown()` stops accepting new tasks but lets queued tasks finish. Then I call `awaitTermination()` with a generous timeout. If tasks are still running after the timeout, I call `shutdownNow()` to interrupt them. In Spring Boot I use `@PreDestroy` to hook this into the application shutdown lifecycle."

```java
import java.util.concurrent.*;

public class SafeShutdown {

    private final ExecutorService executor = Executors.newFixedThreadPool(10);

    // @PreDestroy in Spring Boot — called on application shutdown
    public void shutdown() throws InterruptedException {
        System.out.println("Shutting down — draining queued tasks...");
        executor.shutdown(); // no new tasks accepted; queued tasks continue

        if (!executor.awaitTermination(60, TimeUnit.SECONDS)) {
            System.err.println("Tasks still running after 60s — forcing shutdown...");
            executor.shutdownNow(); // interrupt running tasks

            if (!executor.awaitTermination(10, TimeUnit.SECONDS)) {
                System.err.println("ExecutorService did not terminate cleanly");
            }
        } else {
            System.out.println("All tasks completed gracefully ✅");
        }
    }

    public static void main(String[] args) throws Exception {
        SafeShutdown demo = new SafeShutdown();

        // Submit 20 tasks
        for (int i = 0; i < 20; i++) {
            final int id = i;
            demo.executor.submit(() -> {
                try {
                    Thread.sleep(200);
                    System.out.println("Task " + id + " done");
                } catch (InterruptedException e) {
                    System.out.println("Task " + id + " interrupted");
                    Thread.currentThread().interrupt();
                }
            });
        }

        // JVM shutdown hook as backup
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            try { demo.shutdown(); } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }));

        demo.shutdown();
    }
}
```

---

## Q3. How do two threads end up in a deadlock (Thread-1 has A, wants B; Thread-2 has B, wants A)? How do you handle it?

**Spoken Answer:**
"Deadlock happens when two threads each hold a lock the other needs, and both wait forever.

The solutions are:
1. **Lock ordering** — always acquire locks in the same order (always A then B). If every thread follows the same sequence, circular waiting is impossible.
2. **tryLock with timeout** — using `ReentrantLock.tryLock()`, if a thread can't get the second lock within a timeout it releases the first and retries later.
3. **Avoid nested locks** — redesign so you don't need to hold two locks simultaneously."

```java
import java.util.concurrent.*;
import java.util.concurrent.locks.*;

public class DeadlockFix {

    static final Object A = new Object();
    static final Object B = new Object();
    static final ReentrantLock lockA = new ReentrantLock();
    static final ReentrantLock lockB = new ReentrantLock();

    // ─── ❌ Deadlock: inconsistent order ───
    static Thread badThread1() {
        return new Thread(() -> {
            synchronized (A) {
                sleep(50);
                synchronized (B) { System.out.println("T1 done (bad — won't reach)"); }
            }
        });
    }
    static Thread badThread2() {
        return new Thread(() -> {
            synchronized (B) { // opposite order → deadlock
                sleep(50);
                synchronized (A) { System.out.println("T2 done (bad — won't reach)"); }
            }
        });
    }

    // ─── ✅ Fix 1: Consistent lock ordering (always A → B) ───
    static Thread goodThread(String name) {
        return new Thread(() -> {
            synchronized (A) {          // same order for both threads
                synchronized (B) {
                    System.out.println(name + " done ✅");
                }
            }
        }, name);
    }

    // ─── ✅ Fix 2: tryLock with timeout — release and retry ───
    static void tryLockApproach(String name) throws InterruptedException {
        while (true) {
            boolean gotA = false, gotB = false;
            try {
                gotA = lockA.tryLock(100, TimeUnit.MILLISECONDS);
                gotB = lockB.tryLock(100, TimeUnit.MILLISECONDS);
                if (gotA && gotB) {
                    System.out.println(name + " got both locks via tryLock ✅");
                    return;
                }
            } finally {
                if (gotA) lockA.unlock();
                if (gotB) lockB.unlock();
            }
            Thread.sleep(10 + (long)(Math.random() * 10)); // randomized backoff
        }
    }

    static void sleep(long ms) { try { Thread.sleep(ms); } catch (InterruptedException e) {} }

    public static void main(String[] args) throws InterruptedException {
        // Fix 1: consistent ordering
        Thread t1 = goodThread("T1"); Thread t2 = goodThread("T2");
        t1.start(); t2.start(); t1.join(); t2.join();

        // Fix 2: tryLock
        Thread r1 = new Thread(() -> { try { tryLockApproach("R1"); } catch (Exception e) {} });
        Thread r2 = new Thread(() -> { try { tryLockApproach("R2"); } catch (Exception e) {} });
        r1.start(); r2.start(); r1.join(); r2.join();
    }
}
```

---

## Q4. How do you implement a Circuit Breaker? When do you use the retry mechanism?

**Spoken Answer:**
"A Circuit Breaker wraps external calls and monitors failure rate. When failures exceed a threshold (e.g. 50%), the circuit 'opens' and all calls fail fast without hitting the service — preventing cascading failures. After a cooldown period it enters 'half-open' and allows one test call. If that succeeds, the circuit closes.

I use Resilience4j in Spring Boot. For retries, I use them for **transient failures** — network timeouts, momentary 503s. I configure exponential backoff: 500ms, 1s, 2s. I do NOT retry on 4xx errors — those are client errors, retrying won't help."

```java
// Maven: resilience4j-spring-boot3

import io.github.resilience4j.circuitbreaker.annotation.CircuitBreaker;
import io.github.resilience4j.retry.annotation.Retry;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

// application.yml:
/*
resilience4j:
  circuitbreaker:
    instances:
      paymentService:
        slidingWindowSize: 10               # evaluate last 10 calls
        failureRateThreshold: 50            # open circuit if 50% fail
        waitDurationInOpenState: 10000      # stay open for 10 seconds
        permittedNumberOfCallsInHalfOpenState: 3  # allow 3 test calls
        automaticTransitionFromOpenToHalfOpenEnabled: true
  retry:
    instances:
      paymentService:
        maxAttempts: 3
        waitDuration: 500ms
        enableExponentialBackoff: true
        exponentialBackoffMultiplier: 2     # 500ms, 1s, 2s
        retryExceptions:
          - java.io.IOException
          - java.net.SocketTimeoutException
        ignoreExceptions:
          - com.example.BadRequestException   # don't retry 4xx
*/

@Service
class PaymentService {

    private final RestTemplate restTemplate;

    PaymentService(RestTemplate restTemplate) { this.restTemplate = restTemplate; }

    @CircuitBreaker(name = "paymentService", fallbackMethod = "paymentFallback")
    @Retry(name = "paymentService")
    public String processPayment(String orderId, double amount) {
        return restTemplate.postForObject(
            "https://external-payment-api.com/pay",
            new java.util.HashMap<>(java.util.Map.of("orderId", orderId, "amount", amount)),
            String.class
        );
    }

    // Called when circuit is OPEN or all retries exhausted
    public String paymentFallback(String orderId, double amount, Exception ex) {
        // Save to DB queue for async retry when service recovers
        System.out.println("Circuit open — queuing payment for order: " + orderId);
        return "Payment queued. Will process shortly.";
    }
}
```

---

## Q5. What was the purpose of Kafka in your project? What is Kafka and where is it used?

**Spoken Answer:**
"Kafka is a distributed event streaming platform. In my project I used it to **decouple microservices** — when an order is placed, the Order service publishes an event to a Kafka topic. The Notification service, Inventory service, and Analytics service all consume that event independently without the Order service needing to know about them.

Benefits: high throughput, fault tolerance via replication, each consumer group reads at its own pace, and events are retained so we can replay them."

```java
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.*;

// application.yml:
/*
spring:
  kafka:
    bootstrap-servers: localhost:9092
    producer:
      value-serializer: org.springframework.kafka.support.serializer.JsonSerializer
    consumer:
      auto-offset-reset: earliest
      value-deserializer: org.springframework.kafka.support.serializer.JsonDeserializer
      properties.spring.json.trusted.packages: "*"
*/

record OrderEvent(String orderId, String product, double amount) {}

// ─── Producer: Order Service ───
@Service
class OrderProducer {
    private final KafkaTemplate<String, OrderEvent> kafka;

    OrderProducer(KafkaTemplate<String, OrderEvent> kafka) { this.kafka = kafka; }

    public void publishOrder(OrderEvent event) {
        kafka.send("order-events", event.orderId(), event);
        System.out.println("Published: " + event.orderId());
    }
}

// ─── Consumer 1: Notification Service (independent consumer group) ───
@Component
class NotificationConsumer {
    @KafkaListener(topics = "order-events", groupId = "notification-group")
    public void handle(OrderEvent event) {
        System.out.println("Sending email for order: " + event.orderId());
    }
}

// ─── Consumer 2: Inventory Service (independent consumer group) ───
@Component
class InventoryConsumer {
    @KafkaListener(topics = "order-events", groupId = "inventory-group")
    public void handle(OrderEvent event) {
        System.out.println("Reducing stock for: " + event.product());
    }
}

// ─── Fix for consumer stopping after partition rebalance ───
@Component
class ResilientConsumer {
    @KafkaListener(topics = "order-events", groupId = "resilient-group")
    public void handle(OrderEvent event,
                       org.springframework.kafka.support.Acknowledgment ack) {
        try {
            process(event);
            ack.acknowledge(); // manual commit only on success
        } catch (Exception e) {
            System.err.println("Processing failed — will retry after rebalance: " + event.orderId());
            // do NOT ack — Kafka will re-deliver after rebalance
        }
    }
    void process(OrderEvent e) { System.out.println("Processed: " + e); }
}
```

---

## Q6. How do you process a 10GB CSV file without running out of memory?

**Spoken Answer:**
"The key is **never loading the whole file into memory**. Java's `Files.lines()` returns a lazy stream — it reads one line at a time. I pair that with `BufferedReader` with a tuned buffer size.

For Spring Batch production scenarios, I use chunk-oriented processing: read 1000 rows, process, write to DB — all in one transaction chunk. It's restartable and parallel-partition-capable."

```java
import java.io.*;
import java.nio.file.*;
import java.util.stream.*;

public class LargeFileProcessing {

    // ✅ Files.lines — lazy stream, O(1) memory per line
    static void streamCSV(String filePath) throws IOException {
        long[] count = {0};
        double[] total = {0};

        try (Stream<String> lines = Files.lines(Path.of(filePath))) {
            lines.skip(1) // skip header
                 .map(line -> line.split(","))
                 .filter(parts -> parts.length >= 3)
                 .forEach(parts -> {
                     try {
                         double salary = Double.parseDouble(parts[2].trim());
                         if (salary > 55_000) { count[0]++; total[0] += salary; }
                     } catch (NumberFormatException ignored) {}
                 });
        }
        System.out.printf("Found: %d, Avg salary: %.0f%n", count[0],
                count[0] > 0 ? total[0] / count[0] : 0);
    }

    // ✅ BufferedReader — explicit 64KB buffer, total control
    static void bufferedReader(String path) throws IOException {
        try (BufferedReader reader = new BufferedReader(
                new FileReader(path), 64 * 1024)) { // 64KB buffer
            reader.readLine(); // skip header
            String line;
            while ((line = reader.readLine()) != null) {
                process(line); // one line at a time — never accumulate
            }
        }
    }

    static void process(String line) {
        String[] parts = line.split(",");
        System.out.println("Processing: " + parts[0]);
    }

    // Spring Batch (conceptual):
    /*
    @Bean Step processFileStep(JobRepository jr, PlatformTransactionManager tm) {
        return new StepBuilder("processFile", jr)
            .<InputRow, OutputRow>chunk(1000, tm)   // 1000 rows per transaction
            .reader(flatFileItemReader())             // reads sequentially from file
            .processor(rowProcessor())               // filter/transform
            .writer(jdbcBatchItemWriter())           // batch INSERT to DB
            .build();
        // Benefits: checkpointing, restartability, parallel partitioning
    }
    */

    public static void main(String[] args) throws IOException {
        Path tmp = Files.createTempFile("demo", ".csv");
        Files.write(tmp, "name,age,salary\nAlice,30,60000\nBob,25,45000\n".getBytes());
        streamCSV(tmp.toString());
        bufferedReader(tmp.toString());
        Files.delete(tmp);
    }
}
```

---

## Q7. Your POST request returns 200 OK but data is not saved. How do you debug this in a layered architecture?

**Spoken Answer:**
"This is a classic debugging scenario. I check these in order:

1. Is `@Transactional` on the service method? Without it, Spring doesn't manage the transaction.
2. **Self-invocation bug** — if the `@Transactional` method is called from another method in the same class, Spring's proxy is bypassed and the transaction never starts.
3. **Swallowed exception** — an empty catch block swallows the error, method returns normally, but the transaction was rolled back silently.
4. Wrong `rollbackFor` — by default Spring rolls back on `RuntimeException`. Checked exceptions pass through without rollback unless you add `rollbackFor = Exception.class`."

```java
import org.springframework.transaction.annotation.Transactional;
import org.springframework.stereotype.*;

@Service
class DebugService {

    private final EmployeeRepository repo;
    DebugService(EmployeeRepository repo) { this.repo = repo; }

    // ❌ Bug 1: self-invocation — proxy bypassed, @Transactional has NO effect
    public void saveWrapper(Employee emp) {
        saveEmployee(emp); // 'this.saveEmployee()' — not through Spring proxy!
    }

    @Transactional
    public void saveEmployee(Employee emp) {
        repo.save(emp); // transaction only starts when called FROM OUTSIDE the class
    }

    // ❌ Bug 2: swallowed exception — 200 returned but data rolled back
    @Transactional
    public void saveSilentFail(Employee emp) {
        try {
            repo.save(emp);
            if (emp.name() == null) throw new RuntimeException("validation failed");
        } catch (Exception e) {
            // empty catch → exception swallowed → @Transactional rolls back quietly
            System.out.println("Suppressed: " + e.getMessage()); // ❌
        }
    }

    // ❌ Bug 3: checked exception does NOT trigger rollback by default
    @Transactional
    public void saveWithChecked(Employee emp) throws Exception {
        repo.save(emp);
        throw new Exception("checked"); // NOT rolled back by default!
    }

    // ✅ Correct: let RuntimeException propagate out of @Transactional method
    @Transactional
    public void saveCorrectly(Employee emp) {
        repo.save(emp); // if exception here, propagate up — transaction rolls back
    }

    // ✅ Correct: rollback on checked exceptions too
    @Transactional(rollbackFor = Exception.class)
    public void saveWithCheckedFixed(Employee emp) throws Exception {
        repo.save(emp);
        throw new Exception("now rolls back"); // ✅
    }
}
```

---

## Q8. How do you implement IP whitelisting for sensitive APIs in Spring Boot?

**Spoken Answer:**
"I create a `OncePerRequestFilter` that checks the remote IP against a whitelist before passing the request to the controller. For requests behind a load balancer or reverse proxy, the real client IP is in the `X-Forwarded-For` header. I externalise the whitelist to `application.properties` so it's configurable per environment."

```java
import jakarta.servlet.*;
import jakarta.servlet.http.*;
import org.springframework.web.filter.OncePerRequestFilter;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import java.io.IOException;
import java.util.Set;

@Component
class IpWhitelistFilter extends OncePerRequestFilter {

    // In production: load from @ConfigurationProperties / application.yml
    private static final Set<String> ALLOWED = Set.of("127.0.0.1", "10.0.0.5", "192.168.1.100");
    private static final Set<String> PROTECTED = Set.of("/admin", "/actuator", "/internal");

    @Override
    protected void doFilterInternal(HttpServletRequest req,
                                    HttpServletResponse res,
                                    FilterChain chain) throws ServletException, IOException {
        String uri = req.getRequestURI();
        boolean isProtected = PROTECTED.stream().anyMatch(uri::startsWith);

        if (isProtected && !ALLOWED.contains(getClientIp(req))) {
            res.setStatus(HttpStatus.FORBIDDEN.value());
            res.getWriter().write("{\"error\": \"Access denied: IP not whitelisted\"}");
            return;
        }
        chain.doFilter(req, res);
    }

    private String getClientIp(HttpServletRequest req) {
        // X-Forwarded-For is set by load balancers / reverse proxies
        String forwarded = req.getHeader("X-Forwarded-For");
        if (forwarded != null && !forwarded.isBlank()) {
            return forwarded.split(",")[0].trim(); // first = original client IP
        }
        return req.getRemoteAddr();
    }
}
```

---

## Q9. What are output-based questions you should know?

### Output Questions with Explanations

```java
public class OutputQuestions {
    public static void main(String[] args) {

        // Q: System.out.println(true && false || true);
        // && has higher precedence than ||
        // → (true && false) || true → false || true → true
        System.out.println(true && false || true);  // true

        // Q: i++ vs ++i when i = 0
        int i = 0;
        System.out.println(i++); // 0  ← prints THEN increments
        System.out.println(i);   // 1  ← i is now 1
        i = 0;
        System.out.println(++i); // 1  ← increments THEN prints

        // Q: ++ch where ch = 'A'
        char ch = 'A'; // ASCII 65
        System.out.println(++ch); // B  (65+1 = 66 = 'B')

        // Q: 1+2+"Java"+(3+4)
        // Left to right: 1+2=3 → 3+"Java"="3Java" → "3Java"+(7)="3Java7"
        System.out.println(1 + 2 + "Java" + (3 + 4)); // 3Java7

        // Q: 5+5+"5"+5+5
        // 5+5=10 → 10+"5"="105" → "105"+5="1055" → "1055"+5="10555"
        System.out.println(5 + 5 + "5" + 5 + 5);      // 10555

        // Q: 10 == 10.0
        // int 10 promoted to double 10.0 for comparison
        System.out.println(10 == 10.0); // true

        // Q: Integer a=128, b=128 — a==b?
        // IntegerCache only caches -128 to 127. 128 is a new heap object each time.
        Integer a = 128, b = 128;
        System.out.println(a == b);      // false ← different heap objects
        System.out.println(a.equals(b)); // true  ← same value → always use equals!
        Integer x = 127, y = 127;
        System.out.println(x == y); // true ← from cache

        // Q: Can you mutate a final List?
        final java.util.List<String> list = new java.util.ArrayList<>();
        list.add("yes");             // ✅ contents are mutable
        // list = new ArrayList<>(); // ❌ COMPILE ERROR — reference is final

        // Q: ArrayStoreException
        Object[] objs = new String[3]; // compiles — covariant array type
        objs[0] = "OK";                // fine
        try {
            objs[1] = 42;              // ❌ ArrayStoreException at runtime!
        } catch (ArrayStoreException e) {
            System.out.println("Can't store Integer into String[]");
        }

        // Q: Can interfaces have logic? Yes — since Java 8
        // default methods and static methods in interfaces are allowed
    }
}
```

---

## Q10. How do you implement Spring Scheduler? How do you handle high-volume REST APIs (10K req/s with spikes)?

**Q: Spring Scheduler**

**Spoken Answer:**
"I annotate the main class with `@EnableScheduling` and annotate any method with `@Scheduled`. For cron-based or fixed-rate tasks, this is sufficient. For jobs that might overlap or run long, I configure a `ThreadPoolTaskScheduler` so tasks run on a pool, not the single default thread."

```java
import org.springframework.scheduling.annotation.*;
import org.springframework.stereotype.*;

@Component
class ScheduledJobs {

    // Fixed rate: every 30 seconds regardless of execution time
    @Scheduled(fixedRate = 30_000)
    public void syncInventory() {
        System.out.println("Syncing inventory at: " + java.time.LocalTime.now());
    }

    // Fixed delay: 10s AFTER previous execution completes
    @Scheduled(fixedDelay = 10_000)
    public void cleanTempFiles() {
        System.out.println("Cleaning temp files...");
    }

    // Cron expression: every day at 2:30 AM
    @Scheduled(cron = "0 30 2 * * *")
    public void dailyReport() {
        System.out.println("Generating daily report...");
    }
}

// application.yml — configure thread pool so jobs don't block each other:
/*
spring:
  task:
    scheduling:
      pool:
        size: 5
*/
```

**Q: Optimizing REST API for 10K requests/second with spike response times**

**Spoken Answer:**
"I approach this in layers:
1. **Caching** — add Redis for frequently-read data with `@Cacheable`. Eliminates DB hits.
2. **Circuit Breaker** — Resilience4j to fail fast when dependencies are slow.
3. **Connection pooling** — tune HikariCP pool size (`maximumPoolSize`, `connectionTimeout`).
4. **Async processing** — offload non-critical work to a thread pool or Kafka.
5. **Rate limiting** — Bucket4j or API gateway to protect the service from overload."

```java
import org.springframework.cache.annotation.*;
import org.springframework.stereotype.*;

@Service
@CacheConfig(cacheNames = "employees")
class EmployeeCachedService {

    // @Cacheable: result cached in Redis after first call
    @Cacheable(key = "#id")
    public String getEmployee(Long id) {
        simulateDbCall();
        return "Employee-" + id;
    }

    // @CacheEvict: remove from cache on update
    @CacheEvict(key = "#id")
    public void updateEmployee(Long id, String name) {
        System.out.println("Updated and cache evicted for: " + id);
    }

    void simulateDbCall() { System.out.println("DB hit (only on first call per id)"); }
}

// application.yml:
/*
spring:
  cache:
    type: redis
  data:
    redis:
      host: localhost
      port: 6379

  datasource:
    hikari:
      maximum-pool-size: 30
      minimum-idle: 10
      connection-timeout: 3000
      idle-timeout: 600000
*/
```
