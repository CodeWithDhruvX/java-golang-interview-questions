# 🚀 05 — Performance Optimization
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- JMH benchmarking
- Profiling with async-profiler / JFR / VisualVM
- Connection pool tuning (HikariCP)
- Reducing allocations (object reuse, pooling)
- String optimization
- Lazy initialization, caching strategies

---

## ❓ Most Asked Questions

### Q1. How do you benchmark Java code with JMH?

```java
// JMH — Java Microbenchmark Harness (the correct way to benchmark)
@BenchmarkMode(Mode.AverageTime)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@State(Scope.Thread)       // new state per thread
@Warmup(iterations = 5, time = 1)    // warm up JIT
@Measurement(iterations = 10, time = 1)
@Fork(2)                   // run in 2 separate JVM processes to avoid JIT variance
public class StringBenchmark {

    private static final int SIZE = 10_000;
    private String[] words;

    @Setup
    public void setup() {
        words = IntStream.range(0, SIZE)
                         .mapToObj(i -> "word" + i)
                         .toArray(String[]::new);
    }

    @Benchmark
    public String concatenationPlus() {
        String result = "";
        for (String w : words) result += w;  // ❌ O(n²) allocations
        return result;
    }

    @Benchmark
    public String stringBuilder() {
        StringBuilder sb = new StringBuilder(SIZE * 8);
        for (String w : words) sb.append(w);  // ✅ amortized O(n)
        return sb.toString();
    }

    @Benchmark
    public String joining() {
        return String.join("", words);  // ✅ even cleaner
    }
}
// Run: java -jar benchmarks.jar StringBenchmark -f 2 -wi 5 -i 10
```

---

### Q2. How do you profile a Java application?

```bash
# Java Flight Recorder (JFR) — low-overhead, built-in Java 11+
# Start recording
jcmd <pid> JFR.start duration=60s filename=/tmp/app.jfr

# Dump on OOM
java -XX:StartFlightRecording=duration=60s,filename=app.jfr \
     -XX:FlightRecorderOptions=stackdepth=128 MyApp

# Analyze with JDK Mission Control (JMC) — GUI
# Or: async-profiler for CPU + allocation + lock profiling
./profiler.sh -d 30 -e cpu -f /tmp/cpu.html <pid>    # CPU flame graph
./profiler.sh -d 30 -e alloc -f /tmp/alloc.html <pid> # Allocation profile

# VisualVM — GUI profiler (attach to live process)
# Heap dumps, CPU snapshots, thread monitoring
```

```java
// Programmatic JFR event
public class DatabaseQueryEvent extends jdk.jfr.Event {
    @Label("Query") String sql;
    @Label("Duration (ms)") long durationMs;
}

// Usage
DatabaseQueryEvent event = new DatabaseQueryEvent();
event.begin();
try {
    result = executeQuery(sql);
} finally {
    event.sql = sql;
    event.durationMs = System.currentTimeMillis() - start;
    event.commit();   // sends to JFR
}
```

---

### Q3. How do you tune HikariCP connection pool?

```yaml
spring:
  datasource:
    hikari:
      minimum-idle: 5           # keep 5 connections ready
      maximum-pool-size: 20     # max connections to DB
      idle-timeout: 600000      # remove idle connections after 10 min
      max-lifetime: 1800000     # recycle connections every 30 min
      connection-timeout: 30000 # wait 30s before "no connection" error
      leak-detection-threshold: 60000  # warn if connection held > 60s
      pool-name: OrderDB-Pool
```

```java
// Formula for pool size:
// pool_size = (core_count * 2) + effective_spindle_count
// For most web apps: max-pool-size = (CPU cores * 2) connections

// Anti-patterns:
// 1. Pool too large — too many DB connections, DB gets overwhelmed
// 2. Pool too small — connection waiting → timeout → errors
// 3. Holding connections too long (large transactions, external calls within transaction)

// Check pool metrics via Actuator
// GET /actuator/metrics/hikaricp.connections.active
// GET /actuator/metrics/hikaricp.connections.pending
// GET /actuator/metrics/hikaricp.connections.max
```

---

### Q4. How do you reduce object allocations?

```java
// 1. Object pooling — reuse expensive objects
// Apache Commons Pool
GenericObjectPool<Connection> pool = new GenericObjectPool<>(factory, config);
Connection conn = pool.borrowObject();
try { /* use */ } finally { pool.returnObject(conn); }

// 2. ThreadLocal for per-thread reuse (avoid creating per-request)
private static final ThreadLocal<SimpleDateFormat> DATE_FORMAT =
    ThreadLocal.withInitial(() -> new SimpleDateFormat("yyyy-MM-dd"));

// Usage
String formatted = DATE_FORMAT.get().format(new Date());

// 3. StringBuilder pre-sizing — avoid repeated resizing
// BAD:
StringBuilder sb = new StringBuilder();   // default 16 chars
// GOOD:
StringBuilder sb2 = new StringBuilder(256);  // pre-size if approximate size is known

// 4. Use primitives and primitive arrays where possible
int[] primitiveArr = new int[1000];    // 4 KB
Integer[] boxedArr  = new Integer[1000]; // ~20 KB (header + pointer per element)

// 5. Avoid autoboxing in hot paths
// BAD:
Map<String, Integer> map = new HashMap<>();
for (int i = 0; i < 1_000_000; i++) map.put("key" + i, i);  // 1M Integer allocations!

// BETTER: Use Eclipse Collections or Trove for int-keyed maps
// Or consider whether a Map is the right structure

// 6. String deduplication (Java 8u20+, G1 only)
// -XX:+UseStringDeduplication — JVM eliminates duplicate String char arrays
```

---

### Q5. What are lazy initialization patterns?

```java
// Double-checked locking (DCL) — thread-safe lazy init
public class LazyService {
    private volatile ExpensiveResource resource;

    public ExpensiveResource getResource() {
        if (resource == null) {              // fast path — no lock
            synchronized (this) {
                if (resource == null) {      // slow path — with lock
                    resource = initExpensiveResource();
                }
            }
        }
        return resource;
    }
}

// Holder idiom — cleaner, zero synchronization overhead
public class ServiceLocator {
    private ServiceLocator() {}

    private static class Holder {
        static final ServiceLocator INSTANCE = new ServiceLocator();
    }

    public static ServiceLocator getInstance() { return Holder.INSTANCE; }
}

// Spring @Lazy — delay bean initialization until first use
@Service
@Lazy   // not created at startup — only when first autowired/requested
public class HeavyReportingService {
    public HeavyReportingService() {
        System.out.println("Initialized (only when needed)");
    }
}

// Programmatic lazy with Supplier
private Supplier<ExpensiveReport> reportSupplier =
    () -> generateReport();  // called only when needed

// Memoize (compute once, return same result)
private Supplier<Config> memoized = Suppliers.memoize(() -> loadConfig());  // Guava
```

---

### Q6. How do you optimize DB-intensive workloads?

```java
// 1. Fetch only what you need — projections instead of full entities
public interface UserSummary {
    Long getId();
    String getFirstName();
    String getEmail();
}

@Query("SELECT u.id AS id, u.firstName AS firstName, u.email AS email FROM User u")
List<UserSummary> findAllSummaries();  // much smaller than fetching all User fields

// 2. Batch inserts — reduce roundtrips
@Modifying
@Query("INSERT INTO audit_log (user_id, action, ts) VALUES (:userId, :action, :ts)")
void logBatch(@Param("userId") Long userId, @Param("action") String action,
              @Param("ts") Instant ts);

// Spring Data: configure batch size
// spring.jpa.properties.hibernate.jdbc.batch_size=50
// spring.jpa.properties.hibernate.order_inserts=true

// 3. Read replicas — route reads to replica, writes to primary
@DataSource("replica")
@Transactional(readOnly = true)
public List<ProductDTO> getAllProducts() { ... }

// 4. Async DB operations — don't block request thread on slow queries
@Async("asyncPool")
public CompletableFuture<List<Report>> generateReports(DateRange range) {
    return CompletableFuture.completedFuture(reportRepository.findByRange(range));
}

// 5. Pagination — never load entire table
Page<Product> page = productRepository.findAll(PageRequest.of(0, 50));
// Or: Cursor-based for large datasets (stable, no "skip N" problem)
List<Product> products = productRepository.findByIdGreaterThan(lastId, PageRequest.of(0, 50));
```
