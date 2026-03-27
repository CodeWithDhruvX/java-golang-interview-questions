# 🧵 Java Concurrency & Multithreading (Rapid-Fire)

> 🔑 **Master Keyword:** **"TELF-SDA"** → Threads, Executors, Locks, Futures, Synchronized, Deadlock, Atomic

---

## 🧵 Section 1: Thread Basics

### Q1: Three Ways to Create a Thread?
🔑 **Keyword: "TRC"** → Thread-Runnable-Callable

```java
// 1. Extend Thread (avoid — ties you to Thread class)
class MyThread extends Thread {
    public void run() { System.out.println("Thread Running"); }
}
new MyThread().start();

// 2. Implement Runnable (PREFERRED — can extend another class)
class MyTask implements Runnable {
    public void run() { System.out.println("Task Running"); }
}
new Thread(new MyTask()).start();

// 3. Implement Callable (returns result + can throw checked exception)
Callable<String> task = () -> "Result";
Future<String> future = executor.submit(task);
String result = future.get(); // blocks until done
```

---

### Q2: `Thread.start()` vs `Thread.run()`?
🔑 **Keyword: "SNR"** → Start=New-thread, Run=same-thread

- `start()` → creates a **new thread**, then calls `run()` in it
- `run()` → executes in the **current thread** (just a regular method call, no new thread!)

---

### Q3: Thread States (Lifecycle)?
🔑 **Keyword: "NRWBTD"** → New→Runnable→Waiting/Blocked/Timed-waiting→Dead

```
NEW → start() → RUNNABLE → running...
RUNNABLE → wait()/sleep()/blocking-IO → WAITING/TIMED_WAITING/BLOCKED
Any → completes or exception → TERMINATED (DEAD)
```

---

### Q4: `wait()` vs `sleep()`?
🔑 **Keyword: "WRL-SK"** → Wait=Releases-Lock, Sleep=Keeps-lock

| Feature | `wait()` | `sleep()` |
|---|---|---|
| Lock | **Releases** lock | **Keeps** lock |
| Location | Inside `synchronized` block | Anywhere |
| Wakeup | `notify()` / `notifyAll()` | Timer expires |
| Class | `Object` | `Thread` |

---

### Q5: Daemon Thread?
🔑 **Keyword: "DJVM"** → Daemon=JVM-exits-when-only-daemons-left

```java
Thread t = new Thread(() -> { /* background task */ });
t.setDaemon(true); // must set BEFORE start()
t.start();
```
JVM exits when only daemon threads remain. GC is a daemon thread. Set before `start()`!

---

## ⚙️ Section 2: Executor Framework

### Q6: Why Executor Framework? (Java 5+)
🔑 **Keyword: "RP-Pool"** → Reuse+Pool instead of new Thread

Stop creating `new Thread()` manually. Use Thread Pools:

```java
ExecutorService executor = Executors.newFixedThreadPool(10);
executor.submit(() -> System.out.println("Task 1")); // returns Future
executor.execute(() -> System.out.println("Task 2")); // returns void
executor.shutdown(); // graceful shutdown
```

---

### Q7: Types of Thread Pools?
🔑 **Keyword: "FCSS"** → Fixed/Cached/Single/Scheduled

| Pool | Description | Best for |
|---|---|---|
| `newFixedThreadPool(n)` | Fixed n threads, queue extra tasks | Server apps, known load |
| `newCachedThreadPool()` | Creates new threads as needed, kills idle | Many short-lived tasks |
| `newSingleThreadExecutor()` | 1 thread, sequential | Sequential processing |
| `newScheduledThreadPool(n)` | Periodic/delayed tasks | Cron jobs |

---

### Q8: `submit()` vs `execute()`?
🔑 **Keyword: "SFE"** → Submit=Future, Execute=fire-and-forget

```java
// execute — fire and forget, returns void
executor.execute(() -> System.out.println("done"));

// submit — returns Future to check result or exception
Future<String> f = executor.submit(() -> "result");
String val = f.get(); // blocks waiting for result
```

---

### Q9: `Future` limitations → `CompletableFuture`?
🔑 **Keyword: "CFP"** → CompletableFuture=Promise-like-Chaining

```java
// Future — blocking, can't chain
Future<String> f = executor.submit(() -> "Hello");
f.get(); // BLOCKS!

// CompletableFuture — non-blocking, chainable (Java 8+)
CompletableFuture.supplyAsync(() -> "Hello")          // async supply
    .thenApply(s -> s + " World")                     // transform
    .thenAccept(System.out::println)                  // consume
    .exceptionally(e -> { System.out.println("Error"); return null; });
```

---

### Q10: Virtual Threads (Java 21 — Project Loom)?
🔑 **Keyword: "VJML"** → Virtual=JVM-managed-Lightweight, millions possible

```java
// Create millions of virtual threads!
Thread.startVirtualThread(() -> {
    System.out.println("I am a virtual thread");
});

// Or with executor
ExecutorService vt = Executors.newVirtualThreadPerTaskExecutor();
vt.submit(() -> callSomeBlockingAPI());
```
- Goal: 1 virtual thread per request. JVM maps many virtual threads to few OS threads
- Perfect for: **I/O bound** applications (HTTP calls, DB queries)

---

## 🔒 Section 3: Synchronization & Locks

### Q11: `synchronized` keyword?
🔑 **Keyword: "MOIM"** → Monitor-Object-Implicit-Mutual-exclusion

```java
// Method-level lock (lock = this)
public synchronized void increment() { count++; }

// Block-level lock (explicit object)
public void increment() {
    synchronized(this) { count++; }
}

// Static sync (lock = Class object)
public static synchronized void method() { ... }
```

---

### Q12: `ReentrantLock` vs `synchronized`?
🔑 **Keyword: "RTIF"** → ReentrantLock=TryLock+Interruptible+Fair

| Feature | `synchronized` | `ReentrantLock` |
|---|---|---|
| Explicit unlock | ❌ | ✅ (must call `unlock()`) |
| Try lock | ❌ | ✅ `tryLock()` |
| Interruptible | ❌ | ✅ `lockInterruptibly()` |
| Fairness | ❌ | ✅ (optional) |

```java
Lock lock = new ReentrantLock();
lock.lock();
try { /* critical section */ }
finally { lock.unlock(); } // ALWAYS in finally!
```

---

### Q13: `ReadWriteLock`?
🔑 **Keyword: "MR-SW"** → Multiple-Readers, Single-Writer

```java
ReadWriteLock rwLock = new ReentrantReadWriteLock();

// Multiple readers can hold simultaneously
rwLock.readLock().lock();

// Only one writer at a time (exclusive)
rwLock.writeLock().lock();
```

---

### Q14: Deadlock — What causes it?
🔑 **Keyword: "MHCN-4"** → Mutual-exclusion, Hold+wait, Circular-wait, No-preemption

4 conditions for deadlock (all 4 must hold):
1. **Mutual Exclusion** — resource held by only 1 thread
2. **Hold and Wait** — thread holds one lock, waits for another
3. **No Preemption** — can't forcibly remove lock from thread
4. **Circular Wait** — T1 waits for T2, T2 waits for T1

**Prevention:** Always acquire locks in **same order** across all threads.

```java
// Deadlock example:
Thread1: lock(A) then lock(B)
Thread2: lock(B) then lock(A) // ← circular wait!
```

---

### Q15: Livelock vs Starvation?
🔑 **Keyword: "LCB-SN"** → Livelock=Courtesy-but-no-progress, Starvation=Never-gets-chance

- **Livelock:** Threads keep responding to each other but make no progress (polite deadlock)
- **Starvation:** A thread never gets CPU time because others always have priority

---

## ⚛️ Section 4: Atomic Operations & Visibility

### Q16: `volatile` for concurrency?
🔑 **Keyword: "VMR"** → Visibility+Memory+no-Reorder

- Ensures **visibility** — writes are immediately visible to all threads (main memory, not cache)
- Prevents **instruction reordering**
- ❌ Does NOT guarantee **atomicity** (read-modify-write is NOT atomic)

```java
volatile boolean running = true;  // thread-safe flag

// BUT:
count++; // NOT atomic even with volatile!
// count++ = read → increment → write (3 steps, can interleave)
```

---

### Q17: Atomic classes?
🔑 **Keyword: "ACas"** → Atomic=CompareAndSwap hardware instruction

```java
AtomicInteger count = new AtomicInteger(0);
count.incrementAndGet(); // atomic, lock-free
count.compareAndSet(5, 10); // set to 10 only if current = 5
```
`AtomicInteger`, `AtomicLong`, `AtomicBoolean`, `AtomicReference` — from `java.util.concurrent.atomic`.

---

### Q18: Thread-safe Collections?
🔑 **Keyword: "CCOL"** → ConcurrentHashMap, CopyOnWrite, offer/poll

| Thread-Safe | Equivalent |
|---|---|
| `ConcurrentHashMap` | `HashMap` |
| `CopyOnWriteArrayList` | `ArrayList` |
| `CopyOnWriteArraySet` | `HashSet` |
| `LinkedBlockingQueue` | `LinkedList` (queue) |
| `ConcurrentLinkedQueue` | LinkedList (non-blocking) |

---

### Q19: `CountDownLatch` vs `CyclicBarrier`?
🔑 **Keyword: "CDO-CRR"** → CountDown=One-time, Cyclic=Reusable

| Feature | CountDownLatch | CyclicBarrier |
|---|---|---|
| Reusable | ❌ | ✅ |
| Use-case | Wait for N tasks to complete | N threads wait for each other |
| Count | Counts down to 0 | Counts up to barrier |

```java
// CountDownLatch
CountDownLatch latch = new CountDownLatch(3);
// workers: latch.countDown()
// main: latch.await() // waits for 3 countDowns

// CyclicBarrier
CyclicBarrier barrier = new CyclicBarrier(3, () -> System.out.println("All ready!"));
// each thread: barrier.await() — waits until all 3 reach barrier
```

---

### Q20: `Semaphore`?
🔑 **Keyword: "SP-Permits"** → Semaphore=Parking-Permits (limit concurrent access)

```java
Semaphore semaphore = new Semaphore(3); // 3 permits = 3 concurrent threads max

// Thread:
semaphore.acquire(); // get permit (blocks if 0 available)
try { /* work */ }
finally { semaphore.release(); } // return permit
```

---

## 🔀 Section 5: Inter-Thread Communication

### Q21: `wait()` / `notify()` / `notifyAll()`?
🔑 **Keyword: "WNN-Monitor"** → Must be inside synchronized block

```java
synchronized(obj) {
    while (!condition) {
        obj.wait(); // release lock and wait
    }
    // do work
}

synchronized(obj) {
    condition = true;
    obj.notifyAll(); // wake all waiting threads
}
```
> Always use **`while`** loop (not `if`) to check condition after `wait()` (spurious wakeups)

---

### Q22: Producer-Consumer Pattern?
🔑 **Keyword: "BQPC"** → BlockingQueue=simplest-Producer-Consumer

```java
BlockingQueue<String> queue = new LinkedBlockingQueue<>(10);

// Producer
Thread producer = new Thread(() -> {
    for (int i = 0; i < 100; i++) queue.put("item-" + i);
});

// Consumer
Thread consumer = new Thread(() -> {
    while (true) {
        String item = queue.take(); // blocks if empty
        process(item);
    }
});
```

---

## 📊 Section 6: Java Memory Model (JMM)

### Q23: What is Java Memory Model?
🔑 **Keyword: "HMM-VSO"** → Heap-MethodStack-Memory, Visibility-Sync-Order

- **Heap:** Objects, shared between all threads
- **Stack (per-thread):** Local variables, method calls — thread-private
- **JMM** defines how threads see each other's variable writes (happens-before relationship)

---

### Q24: Happens-Before Guarantee?
🔑 **Keyword: "WMSV"** → Write-happens-before-Subsequent-reads

Conditions that establish happens-before:
- `synchronized` block release → next lock acquisition
- `volatile` write → subsequent `volatile` read
- Thread `start()` → any action in that thread
- `join()` → code after join sees all thread's writes

---

*End of File — Concurrency & Multithreading*
