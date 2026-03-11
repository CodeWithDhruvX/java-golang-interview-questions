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

### 🎯 How to Explain in Interview

"In Java, I have three main ways to create threads. The first is extending the Thread class, but I rarely use this because it prevents me from extending other classes. The second, and most common way, is implementing Runnable - this gives me flexibility since I can still extend another class. The third way is using Callable when I need a result back or want to throw checked exceptions. The key difference is that Runnable's run() method doesn't return anything, while Callable's call() method returns a Future that I can use to get the result. I always call start() to create a new thread - if I call run() directly, it just executes in the current thread, which defeats the purpose of multithreading."

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

### 🎯 How to Explain in Interview

"ExecutorService is the modern way to handle threading in Java - it's like having a professional thread manager instead of manually creating threads. The big advantage is thread reuse - creating threads is expensive, so ExecutorService maintains a pool of worker threads that can execute multiple tasks. I can choose from different pool types: fixed thread pools for predictable concurrency, cached pools for bursty workloads, or single-thread executors for sequential tasks. The best part is the graceful shutdown - I can stop accepting new tasks and wait for existing ones to finish. This prevents resource leaks and makes my applications more robust. ExecutorService also gives me Future objects that let me track task progress and get results."

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

---

### 🎯 How to Explain in Interview

"Synchronization is Java's built-in mechanism for preventing race conditions. When multiple threads access shared data, I need to ensure that operations are atomic. The synchronized keyword does this by allowing only one thread to execute a critical section at a time. I can synchronize an entire method or just a specific block of code for finer control. There's also static synchronization, which locks the Class object instead of the instance. The key thing to remember is that synchronization solves both visibility and atomicity problems - it ensures that all threads see the latest data and that compound operations like increment are atomic. But it comes with a performance cost, so I only synchronize what's absolutely necessary."

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

### 🎯 How to Explain in Interview

"The volatile keyword is all about visibility in multithreaded environments. Without volatile, threads might cache variables in their CPU registers, so changes made by one thread might not be visible to others immediately. This is why you might see a thread that never stops even after you set a flag to false. Volatile guarantees that every read goes to main memory and every write goes to main memory, so all threads see the latest value. But here's the catch - volatile only guarantees visibility, not atomicity. For compound operations like increment, I still need synchronization or atomic classes. I use volatile for flags or status variables where one thread writes and others read, but for complex operations, I reach for AtomicInteger or synchronized blocks."

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

### 🎯 How to Explain in Interview

"wait(), notify(), and notifyAll() are Java's low-level mechanisms for thread coordination. They work like a handoff system - wait() makes a thread release its lock and go to sleep until someone wakes it up. notify() wakes up one random waiting thread, while notifyAll() wakes up all of them. The key thing is that these must be called within synchronized blocks, and wait() automatically releases the lock while waiting. I use these for producer-consumer scenarios where threads need to coordinate. The classic pattern is checking a condition in a while loop, calling wait() if the condition isn't met, and calling notifyAll() when the condition changes. I prefer notifyAll() over notify() to avoid missed signals."

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

### 🎯 How to Explain in Interview

"Atomic classes are Java's lock-free way to handle thread-safe operations. They use a technique called Compare-And-Swap (CAS) under the hood, which is much faster than synchronization in low-contention scenarios. AtomicInteger gives me atomic operations like incrementAndGet() and compareAndSet() without any locks. The beauty is that CAS operations are atomic at the hardware level - they check if the value is what I expect, and if so, update it. This happens in a single CPU instruction. For high-contention scenarios, LongAdder can be even better than AtomicLong because it uses multiple counters to reduce contention. I use atomic classes when I need simple atomic operations on single variables."

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

### 🎯 How to Explain in Interview

"Deadlock is one of the most frustrating concurrency problems - it's when threads are stuck waiting for each other forever. The classic scenario is Thread A holding Lock 1 and waiting for Lock 2, while Thread B holds Lock 2 and waits for Lock 1. Neither can progress, so they're deadlocked. The key to preventing deadlock is consistent lock ordering - if all threads acquire locks in the same order, deadlock can't happen. Another approach is using tryLock with timeouts, so threads don't wait forever. I also try to keep lock scopes as small as possible and avoid nested locks when possible. Deadlock detection tools can help identify these issues, but prevention through good design is always better."

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

### 🎯 How to Explain in Interview

"CountDownLatch and CyclicBarrier are both synchronization utilities, but they serve different purposes. CountDownLatch is a one-way barrier - threads count down until it reaches zero, then waiting threads are released. It's perfect for scenarios like waiting for multiple workers to complete initialization. Once the count reaches zero, the latch can't be reused. CyclicBarrier is different - it's reusable and designed for scenarios where multiple threads need to wait for each other at specific points. Think of it as a meeting point where all threads gather before proceeding to the next phase. The barrier resets automatically after all threads pass through, so it can be used repeatedly for different phases of work."

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

### 🎯 How to Explain in Interview

"CompletableFuture is Java's modern approach to asynchronous programming - it's like a promise that a value will be available in the future. Unlike traditional Future, CompletableFuture lets me chain operations elegantly with methods like thenApply() for transformations, thenAccept() for consuming results, and exceptionally() for handling errors. I can combine multiple futures, run them in parallel, and compose complex asynchronous workflows without blocking. The beauty is that I can write asynchronous code that reads like sequential code. It's much more powerful than basic Future because it supports callbacks, composition, and non-blocking completion. For modern concurrent applications, CompletableFuture is the way to go."

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

---

### 🎯 How to Explain in Interview

"Thread-safe collections are essential when multiple threads need to share data structures. ConcurrentHashMap is the workhorse - it uses sophisticated locking to allow concurrent reads and writes without the performance hit of full synchronization. CopyOnWriteArrayList is specialized for scenarios with many reads and few writes - it creates a new copy of the array on every write, so readers never block. BlockingQueue is perfect for producer-consumer patterns - it handles the blocking and waking up logic automatically. For existing collections, Collections.synchronizedWrapper() provides a quick thread-safe version, but it uses coarse-grained locking, so performance can suffer under high contention. The key is choosing the right collection for the specific access pattern."
