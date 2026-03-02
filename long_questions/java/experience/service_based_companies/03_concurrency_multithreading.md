# ⚡ 03 — Concurrency & Multithreading
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Thread creation (`Thread` class, `Runnable`, `Callable`)
- `ExecutorService` and thread pools
- `synchronized`, `wait`, `notify`
- `volatile` keyword
- `java.util.concurrent` utilities
- Deadlock, race condition, thread safety

---

## ❓ Most Asked Questions

### Q1. How do you create a thread in Java?

```java
// Method 1: Extend Thread class
class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("Thread: " + Thread.currentThread().getName());
    }
}
new MyThread().start();

// Method 2: Implement Runnable (preferred — doesn't lock inheritance)
Runnable task = () -> System.out.println("Runnable task running");
Thread t = new Thread(task, "WorkerThread");
t.start();

// Method 3: Callable (returns result + can throw checked exceptions)
Callable<Integer> callable = () -> {
    Thread.sleep(100);
    return 42;
};
FutureTask<Integer> future = new FutureTask<>(callable);
new Thread(future).start();
int result = future.get();  // blocks until done — returns 42
```

> **Key:** `start()` creates a new thread; calling `run()` directly executes in the current thread.

---

### Q2. What is `ExecutorService` and why use it?

```java
import java.util.concurrent.*;

// Fixed thread pool — reuses N threads
ExecutorService pool = Executors.newFixedThreadPool(4);

// Submit tasks
pool.submit(() -> System.out.println("Task 1"));

Callable<String> task = () -> "Result from task";
Future<String> future = pool.submit(task);
String result = future.get(5, TimeUnit.SECONDS);  // with timeout

// Other pool types
ExecutorService cached  = Executors.newCachedThreadPool();   // unbounded, reuses idle
ExecutorService single  = Executors.newSingleThreadExecutor();  // ordered tasks
ScheduledExecutorService sched = Executors.newScheduledThreadPool(2);

// Scheduled execution
sched.scheduleAtFixedRate(() -> System.out.println("Heartbeat"),
    0, 5, TimeUnit.SECONDS);  // initial delay 0, repeat every 5s

// Graceful shutdown (ALWAYS do this)
pool.shutdown();                           // stop accepting new tasks
boolean done = pool.awaitTermination(30, TimeUnit.SECONDS);
if (!done) pool.shutdownNow();             // force stop if needed
```

---

### Q3. What is `synchronized` and when to use it?

```java
public class SafeCounter {
    private int count = 0;

    // synchronized method — lock on 'this' object
    public synchronized void increment() {
        count++;  // read-modify-write is now atomic
    }

    // synchronized block — finer control over lock scope
    public void incrementBlock() {
        // do non-critical work...
        synchronized (this) {
            count++;    // only this section is locked
        }
        // do more non-critical work...
    }

    // static synchronized — lock on the Class object
    private static int instances = 0;
    public static synchronized void trackInstance() {
        instances++;
    }
}

// Without synchronization — race condition:
// Thread A reads count = 5
// Thread B reads count = 5
// Thread A writes count = 6
// Thread B writes count = 6  ← incremented only once instead of twice!
```

---

### Q4. What is `volatile` keyword?

```java
// Problem: Without volatile, threads may cache variable in CPU registers
class Task implements Runnable {
    private boolean running = true;  // NOT volatile — may be cached

    public void stop() { running = false; }

    @Override
    public void run() {
        while (running) {    // may never see 'running = false' due to caching!
            // do work
        }
    }
}

// Solution: volatile — guarantees visibility across threads
class SafeTask implements Runnable {
    private volatile boolean running = true;  // always reads from main memory

    public void stop() { running = false; }   // immediately visible to all threads

    @Override
    public void run() {
        while (running) {   // always sees latest value
            // do work
        }
    }
}
```

> **volatile** guarantees **visibility** but NOT **atomicity**.  
> For `i++` (read-modify-write), use `AtomicInteger` or `synchronized`.

---

### Q5. What are `wait()`, `notify()`, and `notifyAll()`?

```java
class Buffer {
    private final Queue<Integer> queue = new LinkedList<>();
    private final int MAX = 5;

    // Producer
    public synchronized void produce(int item) throws InterruptedException {
        while (queue.size() == MAX) {
            wait();  // release lock and WAIT until notified
        }
        queue.add(item);
        System.out.println("Produced: " + item);
        notifyAll();  // wake up ALL waiting threads
    }

    // Consumer
    public synchronized int consume() throws InterruptedException {
        while (queue.isEmpty()) {
            wait();  // release lock and wait
        }
        int item = queue.poll();
        System.out.println("Consumed: " + item);
        notifyAll();
        return item;
    }
}
```

> - `wait()` — releases lock, goes to WAITING state  
> - `notify()` — wakes one arbitrary waiting thread  
> - `notifyAll()` — wakes all waiting threads (prefer this to avoid missed signals)

---

### Q6. What are `java.util.concurrent` atomic classes?

```java
import java.util.concurrent.atomic.*;

// AtomicInteger — thread-safe without synchronization
AtomicInteger counter = new AtomicInteger(0);
counter.incrementAndGet();    // atomic increment — returns new value
counter.getAndIncrement();    // atomic increment — returns old value
counter.compareAndSet(5, 10); // CAS: if value==5, set to 10, return true

// AtomicReference — CAS for objects
AtomicReference<String> ref = new AtomicReference<>("initial");
ref.compareAndSet("initial", "updated");

// LongAdder — better performance than AtomicLong under high contention
LongAdder adder = new LongAdder();
adder.increment();
adder.add(5);
long total = adder.sum();
```

---

### Q7. What is a deadlock and how to prevent it?

```java
// DEADLOCK SCENARIO:
Object lock1 = new Object();
Object lock2 = new Object();

Thread t1 = new Thread(() -> {
    synchronized (lock1) {
        sleep(50);
        synchronized (lock2) { /* ... */ }  // waiting for lock2
    }
});

Thread t2 = new Thread(() -> {
    synchronized (lock2) {
        sleep(50);
        synchronized (lock1) { /* ... */ }  // waiting for lock1
    }
});
// t1 holds lock1, waits for lock2
// t2 holds lock2, waits for lock1  → DEADLOCK!

// PREVENTION — Always acquire locks in same order
Thread t1Fixed = new Thread(() -> {
    synchronized (lock1) {
        synchronized (lock2) { /* ... */ }  // consistent order
    }
});

Thread t2Fixed = new Thread(() -> {
    synchronized (lock1) {  // same order as t1
        synchronized (lock2) { /* ... */ }
    }
});

// Or use tryLock with timeout (non-blocking)
ReentrantLock rl = new ReentrantLock();
if (rl.tryLock(1, TimeUnit.SECONDS)) {
    try { /* ... */ } finally { rl.unlock(); }
}
```

---

### Q8. What is `CountDownLatch` and `CyclicBarrier`?

```java
// CountDownLatch — wait until N events happen (one-shot)
CountDownLatch latch = new CountDownLatch(3);

// 3 workers count down
ExecutorService pool = Executors.newFixedThreadPool(3);
for (int i = 0; i < 3; i++) {
    final int id = i;
    pool.submit(() -> {
        doWork(id);
        latch.countDown();  // count: 3 → 2 → 1 → 0
    });
}
latch.await();  // main thread blocks until count reaches 0
System.out.println("All workers done!");

// CyclicBarrier — N threads wait for each other at a point (reusable)
CyclicBarrier barrier = new CyclicBarrier(3, () -> System.out.println("All at barrier!"));

Runnable worker = () -> {
    phase1Work();
    barrier.await();      // wait for ALL 3 threads to finish phase 1
    phase2Work();
    barrier.await();      // wait again for phase 2
};
```

---

### Q9. What is `CompletableFuture`?

```java
// Asynchronous, non-blocking computation
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
    // runs in ForkJoinPool
    return fetchDataFromDB();
});

// Chain transformations
CompletableFuture<Integer> result = future
    .thenApply(data -> data.length())       // sync transform
    .thenApplyAsync(len -> len * 2)         // async transform
    .exceptionally(ex -> -1);               // fallback on error

// Combine two futures
CompletableFuture<String> f1 = CompletableFuture.supplyAsync(() -> "Hello");
CompletableFuture<String> f2 = CompletableFuture.supplyAsync(() -> "World");

CompletableFuture<String> combined = f1.thenCombine(f2,
    (s1, s2) -> s1 + " " + s2);
System.out.println(combined.get());  // "Hello World"

// Wait for all
CompletableFuture.allOf(f1, f2).thenRun(() -> System.out.println("Both done"));
```

---

### Q10. What are thread-safe collections?

```java
// ConcurrentHashMap — thread-safe map, segment-level locking (Java 8: CAS)
ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
map.put("a", 1);
map.computeIfAbsent("b", k -> expensiveCompute(k));  // atomic operation

// CopyOnWriteArrayList — thread-safe list, copies on every write (rare writes)
CopyOnWriteArrayList<String> list = new CopyOnWriteArrayList<>();
list.add("element");  // creates new backing array — safe for concurrent readers

// BlockingQueue — producer-consumer queues
BlockingQueue<String> queue = new LinkedBlockingQueue<>(100);
queue.put("item");          // blocks if full
String item = queue.take(); // blocks if empty

ArrayBlockingQueue<Integer> bounded = new ArrayBlockingQueue<>(50);
bounded.offer(1, 1, TimeUnit.SECONDS);  // try for 1 second

// Collections.synchronizedXxx — wraps existing collection (coarse lock)
List<String> syncList = Collections.synchronizedList(new ArrayList<>());
// NOTE: iteration still needs external synchronization
synchronized (syncList) {
    for (String s : syncList) { /* ... */ }
}
```
