# ⚡ 02 — Advanced Concurrency in Java
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Java Memory Model (JMM)
- `ReentrantLock` and `ReentrantReadWriteLock`
- `Semaphore`, `Phaser`, `Exchanger`
- `ForkJoinPool` and parallel streams
- Lock-free data structures (CAS, AtomicXxx)
- `CompletableFuture` advanced chaining
- Virtual threads (Java 21 Project Loom)

---

## ❓ Most Asked Questions

### Q1. What is the Java Memory Model (JMM)?

```java
// JMM defines: when writes by one thread are visible to other threads

// Problem: Without rules, CPU/compiler reorders operations for performance
class Example {
    boolean ready = false;   // NOT volatile — may be cached
    int value = 0;

    // Thread 1
    void writer() {
        value = 42;          // write 1
        ready = true;        // write 2 — CPU may reorder these!
    }

    // Thread 2
    void reader() {
        while (!ready) {}    // may spin forever (stale cached value)
        System.out.println(value);  // may print 0! (reordered write)
    }
}

// Solution: volatile establishes happens-before
class FixedExample {
    volatile boolean ready = false;  // volatile
    int value = 0;

    void writer() {
        value = 42;          // happens-before volatile write
        ready = true;        // volatile write — publishes everything above
    }

    void reader() {
        while (!ready) {}    // volatile read — sees all writes before ready=true
        System.out.println(value);  // guaranteed 42
    }
}
// happens-before relationships also established by:
// synchronized, Thread.start(), Thread.join(), AtomicXxx operations
```

---

### Q2. What is `ReentrantLock` and when to prefer it over `synchronized`?

```java
import java.util.concurrent.locks.*;

public class BoundedBuffer<T> {
    private final Queue<T> buffer = new LinkedList<>();
    private final int maxSize;
    private final ReentrantLock lock = new ReentrantLock(true); // fair lock
    private final Condition notFull  = lock.newCondition();
    private final Condition notEmpty = lock.newCondition();

    public BoundedBuffer(int maxSize) { this.maxSize = maxSize; }

    public void put(T item) throws InterruptedException {
        lock.lock();
        try {
            while (buffer.size() == maxSize) notFull.await();  // release lock + wait
            buffer.add(item);
            notEmpty.signal();   // notify a waiting consumer
        } finally {
            lock.unlock();   // always unlock in finally!
        }
    }

    public T take() throws InterruptedException {
        lock.lock();
        try {
            while (buffer.isEmpty()) notEmpty.await();
            T item = buffer.poll();
            notFull.signal();
            return item;
        } finally {
            lock.unlock();
        }
    }
}

// ReentrantLock advantages over synchronized:
// - tryLock(timeout) — non-blocking attempt
// - interruptible lock acquisition
// - multiple Condition variables (vs single wait/notify)
// - fair mode — longest-waiting thread gets lock first
// - lockInterruptibly() — allows interruption while waiting
```

---

### Q3. What is `ReentrantReadWriteLock`?

```java
// Multiple readers can read simultaneously; writers need exclusive access
public class CachedData<T> {
    private volatile T data;
    private final ReentrantReadWriteLock rwLock = new ReentrantReadWriteLock();
    private final ReadWriteLock.ReadLock  readLock  = rwLock.readLock();
    private final ReadWriteLock.WriteLock writeLock = rwLock.writeLock();

    public T read() {
        readLock.lock();
        try {
            return data;   // multiple threads can read concurrently
        } finally {
            readLock.unlock();
        }
    }

    public void write(T newData) {
        writeLock.lock();
        try {
            data = newData;   // exclusive — blocks all readers and other writers
        } finally {
            writeLock.unlock();
        }
    }
}

// StampedLock (Java 8+) — even faster, supports optimistic reads
StampedLock sl = new StampedLock();
long stamp = sl.tryOptimisticRead();   // no lock acquired
double x = this.x, y = this.y;        // read speculatively
if (!sl.validate(stamp)) {            // was there a write? re-read with lock
    stamp = sl.readLock();
    try { x = this.x; y = this.y; } finally { sl.unlockRead(stamp); }
}
```

---

### Q4. What is `Semaphore`?

```java
// Controls access to a limited resource pool
public class ConnectionPool {
    private final Semaphore semaphore;
    private final Queue<Connection> pool;

    public ConnectionPool(int poolSize) {
        semaphore = new Semaphore(poolSize, true);  // fair
        pool = new ConcurrentLinkedQueue<>();
        for (int i = 0; i < poolSize; i++) pool.offer(createConnection());
    }

    public Connection acquire() throws InterruptedException {
        semaphore.acquire();   // blocks if no permits
        return pool.poll();
    }

    public void release(Connection conn) {
        pool.offer(conn);
        semaphore.release();   // return permit — unblocks a waiting thread
    }

    public Connection tryAcquire(long timeout, TimeUnit unit)
            throws InterruptedException {
        if (semaphore.tryAcquire(timeout, unit)) {
            return pool.poll();
        }
        throw new TimeoutException("No connection available");
    }
}
```

---

### Q5. What is `ForkJoinPool` and parallel streams?

```java
// ForkJoinPool — splits work recursively (divide and conquer)
public class SumTask extends RecursiveTask<Long> {
    private final long[] arr;
    private final int start, end;
    private static final int THRESHOLD = 1000;

    public SumTask(long[] arr, int start, int end) {
        this.arr = arr; this.start = start; this.end = end;
    }

    @Override
    protected Long compute() {
        if (end - start <= THRESHOLD) {
            long sum = 0;
            for (int i = start; i < end; i++) sum += arr[i];
            return sum;
        }
        int mid = (start + end) / 2;
        SumTask left  = new SumTask(arr, start, mid);
        SumTask right = new SumTask(arr, mid, end);
        left.fork();                          // run left in pool asynchronously
        long rightResult = right.compute();   // run right in current thread
        long leftResult  = left.join();       // wait for left
        return leftResult + rightResult;
    }
}

ForkJoinPool pool = new ForkJoinPool(4);  // 4 worker threads
long result = pool.invoke(new SumTask(arr, 0, arr.length));

// Parallel streams use the common ForkJoinPool
long sum = LongStream.rangeClosed(1, 1_000_000)
    .parallel()           // splits stream across ForkJoinPool.commonPool()
    .sum();

// Custom parallelism level for a single stream
ForkJoinPool customPool = new ForkJoinPool(8);
long result2 = customPool.submit(() ->
    list.parallelStream().mapToLong(Long::parseLong).sum()
).get();
```

---

### Q6. Advanced `CompletableFuture` patterns

```java
// Compose 3 async calls with different executors
ExecutorService io = Executors.newCachedThreadPool();
ExecutorService cpu = Executors.newFixedThreadPool(4);

CompletableFuture<UserProfile> profile =
    CompletableFuture.supplyAsync(() -> fetchUser(userId), io)       // DB call
        .thenComposeAsync(user -> fetchOrders(user.getId()), io)     // another DB call
        .thenApplyAsync(orders -> enrichProfile(orders), cpu);       // CPU work

// Timeout (Java 9+)
profile.orTimeout(3, TimeUnit.SECONDS)               // completes exceptionally on timeout
       .exceptionally(ex -> UserProfile.empty());   // fallback

// completeOnTimeout (Java 9+)
profile.completeOnTimeout(UserProfile.empty(), 3, TimeUnit.SECONDS);  // fallback value

// anyOf — first one to complete wins
CompletableFuture<Object> fastest = CompletableFuture.anyOf(
    CompletableFuture.supplyAsync(() -> callPrimaryDB()),
    CompletableFuture.supplyAsync(() -> callReplicaDB())
);

// thenAccept vs thenApply vs thenRun
profile.thenApply(p -> p.getName())    // transform result — returns new CF
       .thenAccept(name -> log(name))  // consume result — returns CF<Void>
       .thenRun(() -> cleanup());      // ignores result — returns CF<Void>
```

---

### Q7. What are Virtual Threads (Java 21 — Project Loom)?

```java
// Platform threads — OS threads, expensive (~1MB stack, context switch cost)
// Virtual threads — cheap, JVM-managed, millions can coexist

// Create virtual thread
Thread vt = Thread.ofVirtual().start(() -> {
    // blocking I/O here does NOT block the OS thread!
    String result = httpClient.get("https://api.example.com");
    System.out.println(result);
});

// Virtual thread executor — ideal for I/O-bound workloads
try (ExecutorService exec = Executors.newVirtualThreadPerTaskExecutor()) {
    List<Future<String>> futures = new ArrayList<>();
    for (int i = 0; i < 10_000; i++) {
        final int id = i;
        futures.add(exec.submit(() -> fetchFromDB(id)));  // 10K virtual threads, no problem!
    }
    for (Future<String> f : futures) System.out.println(f.get());
}

// Spring Boot 3.2+ — enable virtual threads globally
// application.yml: spring.threads.virtual.enabled: true
// All @Async, Tomcat request threads become virtual threads automatically

// ⚠️ Pitfall: synchronized blocks pin virtual threads to carrier thread
// Replace synchronized with ReentrantLock when used with virtual threads
```

---

### Q8. Implement a Thread-safe Singleton

```java
// Double-checked locking (Java 5+ — safe with volatile)
public class Config {
    private static volatile Config instance;  // volatile — visibility + ordering

    private Config() { /* load config */ }

    public static Config getInstance() {
        if (instance == null) {         // first check (no locking)
            synchronized (Config.class) {
                if (instance == null) { // second check (with locking)
                    instance = new Config();
                }
            }
        }
        return instance;
    }
}

// Better: Initialization-on-demand holder idiom (lazy, thread-safe, no synchronization overhead)
public class ConfigHolder {
    private ConfigHolder() {}

    private static class Holder {
        static final ConfigHolder INSTANCE = new ConfigHolder();  // class loaded on first access
    }

    public static ConfigHolder getInstance() { return Holder.INSTANCE; }
}

// Simplest: Enum singleton (serialization-safe, reflection-safe)
public enum AppConfig {
    INSTANCE;

    private final String dbUrl = "jdbc:mysql://...";
    public String getDbUrl() { return dbUrl; }
}
AppConfig.INSTANCE.getDbUrl();
```
