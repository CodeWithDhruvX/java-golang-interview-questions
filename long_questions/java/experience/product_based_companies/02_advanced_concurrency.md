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

### 🎯 How to Explain in Interview

"The Java Memory Model defines when writes by one thread become visible to other threads. Without these rules, the CPU and compiler can reorder operations for performance, causing subtle bugs. For example, without volatile, one thread might see a stale cached value or observe writes in a different order. The JMM establishes happens-before relationships - if action A happens-before action B, then A's writes are guaranteed to be visible to B. I use volatile to establish these relationships, along with synchronized, thread operations, and atomic operations. Volatile is crucial because it prevents both caching and reordering. Understanding JMM is essential for writing correct concurrent code, especially in high-performance systems where visibility and ordering matter."

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

### 🎯 How to Explain in Interview

"ReentrantLock is more flexible than synchronized blocks. While synchronized is simple, ReentrantLock gives me tryLock with timeout for non-blocking attempts, interruptible lock acquisition, multiple Condition variables for fine-grained waiting, and fair mode where the longest-waiting thread gets the lock first. I use it when I need these advanced features - like implementing a bounded buffer with separate notFull and notEmpty conditions. The key is always unlocking in a finally block to avoid deadlocks. For simple cases, synchronized is still fine, but ReentrantLock shines in complex concurrent scenarios where I need more control over locking behavior."

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

### 🎯 How to Explain in Interview

"ReentrantReadWriteLock is perfect for read-heavy workloads where multiple threads can read simultaneously but writers need exclusive access. I use separate read and write locks - multiple readers can acquire the read lock concurrently, but any writer blocks all readers and other writers. This dramatically improves performance for scenarios like cached data that's read often but updated rarely. For even better performance, Java 8 introduced StampedLock with optimistic reads - I can try to read without acquiring any lock, then validate that no write occurred. If a write happened, I retry with a proper read lock. This lock-free approach is faster when writes are rare, making it ideal for high-frequency read scenarios."

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

### 🎯 How to Explain in Interview

"Semaphore is perfect for controlling access to limited resources like database connections or thread pools. I initialize it with a permit count representing the resource capacity. Threads acquire permits before accessing resources and release them when done. If no permits are available, threads block until one becomes available. I can also use tryAcquire with timeout for non-blocking attempts. This pattern is much cleaner than manual resource counting and provides fair ordering if needed. Semaphores are ideal for scenarios where I need to limit concurrent access to external resources or implement resource pooling with proper blocking semantics."

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

### 🎯 How to Explain in Interview

"ForkJoinPool is designed for divide-and-conquer algorithms that can be split into smaller subtasks. It implements a work-stealing algorithm where idle threads steal tasks from busy ones, maximizing CPU utilization. I extend RecursiveTask for tasks that return values or RecursiveAction for void tasks. The key is splitting work until it reaches a threshold, then computing directly. Parallel streams use the common ForkJoinPool automatically, making it easy to parallelize operations. For custom parallelism, I can create my own ForkJoinPool. This approach is perfect for CPU-intensive work like processing large arrays or recursive algorithms, but not for I/O-bound tasks where virtual threads are better."

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

### 🎯 How to Explain in Interview

"CompletableFuture is incredibly powerful for composing asynchronous operations. I can chain operations with different executors - use I/O threads for database calls and CPU threads for computations. Java 9 added timeout support with orTimeout and completeOnTimeout for graceful degradation. The anyOf method is perfect for redundancy scenarios where I want the first response from multiple services. I choose between thenApply for transformations, thenAccept for consuming results, and thenRun for final actions. This chaining approach replaces complex callback hell with clean, readable asynchronous pipelines that handle errors naturally through the exceptionally method."

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

### 🎯 How to Explain in Interview

"Virtual threads are a game-changer for I/O-bound applications. Unlike platform threads that map 1:1 to OS threads and are expensive, virtual threads are lightweight and managed by the JVM. I can create millions of virtual threads without issues. The key benefit is that blocking I/O operations don't block the underlying OS thread - the JVM unmounts the virtual thread and mounts another one. This makes thread-per-request programming models viable again. I use Executors.newVirtualThreadPerTaskExecutor() for I/O-heavy workloads. Spring Boot 3.2+ even supports virtual threads globally. The pitfall is that synchronized blocks can pin virtual threads to carrier threads, so I replace them with ReentrantLock. Virtual threads are perfect for microservices, web servers, and any I/O-bound scenario."

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

---

### 🎯 How to Explain in Interview

"Thread-safe singleton has several approaches. Double-checked locking uses volatile to ensure visibility and ordering - I check twice, once without locking and once with synchronized. This avoids synchronization overhead after initialization. A cleaner approach is the initialization-on-demand holder idiom - the inner class is loaded only when getInstance() is called, leveraging the JVM's class loading guarantees. The simplest and most robust is the enum singleton - it's inherently thread-safe, serialization-safe, and reflection-safe. For most cases, I prefer the enum approach for its simplicity and safety. Double-checked locking is useful when I need lazy initialization with performance requirements, but it's easy to get wrong without proper volatile usage."

---
