# Java Intermediate — Concurrency Patterns

> **Topics:** Thread lifecycle, synchronized, volatile, java.util.concurrent, Executor framework, CountDownLatch, CyclicBarrier, Semaphore, ReentrantLock, BlockingQueue, ConcurrentHashMap, CompletableFuture, AtomicXxx

---

## 📋 Reading Progress

- [ ] **Section 1:** Thread Basics & synchronized (Q1–Q15)
- [ ] **Section 2:** java.util.concurrent Locking (Q16–Q28)
- [ ] **Section 3:** Executor Framework & Thread Pools (Q29–Q40)
- [ ] **Section 4:** Concurrent Collections (Q41–Q50)
- [ ] **Section 5:** CompletableFuture (Q51–Q60)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Thread Basics & synchronized (Q1–Q15)

### 1. Race Condition Without Synchronization
**Q: What is the bug?**
```java
public class Main {
    static int counter = 0;

    static void increment() {
        counter++; // not atomic: read + modify + write
    }

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        t1.start(); t2.start();
        t1.join(); t2.join();
        System.out.println(counter); // not reliably 20000!
    }
}
```
**A:** Some value less than 20000 (unpredictable). **Data race** — `counter++` is not atomic. Multiple threads interleave read-modify-write operations.

---

### 2. synchronized Method — Correct Fix
**Q: What is the output?**
```java
public class Main {
    static int counter = 0;

    static synchronized void increment() { counter++; }

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        t1.start(); t2.start();
        t1.join(); t2.join();
        System.out.println(counter);
    }
}
```
**A:** `20000`. `synchronized` on a static method locks on the class object — only one thread can execute it at a time.

---

### 3. synchronized Block — More Granular
**Q: What is the output?**
```java
public class Main {
    static int counter = 0;
    static final Object lock = new Object();

    public static void main(String[] args) throws InterruptedException {
        Runnable task = () -> {
            for (int i = 0; i < 10000; i++) {
                synchronized (lock) {
                    counter++;
                }
            }
        };
        Thread t1 = new Thread(task), t2 = new Thread(task);
        t1.start(); t2.start();
        t1.join(); t2.join();
        System.out.println(counter);
    }
}
```
**A:** `20000`. `synchronized` blocks are more granular — only the critical section is locked, not the entire method.

---

### 4. volatile — Visibility Guarantee
**Q: What does volatile do?**
```java
public class Main {
    static volatile boolean running = true;

    public static void main(String[] args) throws InterruptedException {
        Thread worker = new Thread(() -> {
            while (running) { /* busy wait */ }
            System.out.println("stopped");
        });
        worker.start();
        Thread.sleep(100);
        running = false; // visible to worker thread immediately due to volatile
        worker.join();
    }
}
```
**A:** `stopped`. Without `volatile`, the worker thread might cache `running` in its CPU register and never see the update from the main thread. `volatile` ensures reads/writes go directly to main memory.

---

### 5. volatile Does NOT Provide Atomicity
**Q: What is the bug?**
```java
public class Main {
    static volatile int counter = 0;

    public static void main(String[] args) throws InterruptedException {
        Runnable inc = () -> { for (int i = 0; i < 10000; i++) counter++; };
        Thread t1 = new Thread(inc), t2 = new Thread(inc);
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(counter);
    }
}
```
**A:** Some value less than 20000. `volatile` guarantees visibility but NOT atomicity. `counter++` is still a 3-step read-modify-write that threads can interleave. Use `AtomicInteger` or `synchronized`.

---

### 6. AtomicInteger — Lock-Free Thread Safety
**Q: What is the output?**
```java
import java.util.concurrent.atomic.*;
public class Main {
    static AtomicInteger counter = new AtomicInteger(0);

    public static void main(String[] args) throws InterruptedException {
        Runnable inc = () -> { for (int i = 0; i < 10000; i++) counter.incrementAndGet(); };
        Thread t1 = new Thread(inc), t2 = new Thread(inc);
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(counter.get());
    }
}
```
**A:** `20000`. `AtomicInteger.incrementAndGet()` is a lock-free atomic operation backed by CAS (Compare-And-Swap) CPU instruction.

---

### 7. wait() and notify() — Producer-Consumer
**Q: What is the output?**
```java
public class Main {
    static final Object lock = new Object();
    static boolean ready = false;

    public static void main(String[] args) throws InterruptedException {
        Thread producer = new Thread(() -> {
            synchronized (lock) {
                ready = true;
                lock.notifyAll();
                System.out.println("produced");
            }
        });
        Thread consumer = new Thread(() -> {
            synchronized (lock) {
                while (!ready) {
                    try { lock.wait(); } catch (InterruptedException e) { Thread.currentThread().interrupt(); }
                }
                System.out.println("consumed");
            }
        });
        consumer.start();
        Thread.sleep(100);
        producer.start();
        consumer.join(); producer.join();
    }
}
```
**A:**
```
produced
consumed
```
`wait()` releases the lock and suspends. `notifyAll()` wakes all waiting threads. Always wrap `wait()` in a `while` loop to handle spurious wakeups.

---

### 8. Deadlock — Classic Example
**Q: What happens?**
```java
public class Main {
    static final Object lock1 = new Object();
    static final Object lock2 = new Object();

    public static void main(String[] args) {
        Thread t1 = new Thread(() -> {
            synchronized (lock1) {
                try { Thread.sleep(100); } catch (InterruptedException e) {}
                synchronized (lock2) { System.out.println("t1 done"); }
            }
        });
        Thread t2 = new Thread(() -> {
            synchronized (lock2) {
                try { Thread.sleep(100); } catch (InterruptedException e) {}
                synchronized (lock1) { System.out.println("t2 done"); }
            }
        });
        t1.start(); t2.start();
    }
}
```
**A:** **Deadlock** — `t1` holds `lock1` and waits for `lock2`. `t2` holds `lock2` and waits for `lock1`. Neither can proceed. **Fix:** Always acquire locks in the same order.

---

### 9. Livelock — Threads Keep Responding to Each Other
**Q: What is the concept?**
```java
// Livelock: threads are not blocked but can't progress
// Example: two people meet in a corridor and keep stepping aside in the same direction
// Unlike deadlock (threads are blocked), livelock threads are running but making no progress.
// Common fix: randomized backoff, or ordered lock acquisition.
public class Main {
    public static void main(String[] args) {
        System.out.println("Livelock = threads active but no progress");
        System.out.println("Deadlock = threads blocked waiting for each other");
        System.out.println("Starvation = a thread never gets CPU time");
    }
}
```
**A:** Prints the descriptions. Understanding the difference between these three thread hazards is critical for interviews.

---

### 10. Thread.join() — Wait for Completion
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> {
            try { Thread.sleep(200); } catch (InterruptedException e) {}
            System.out.println("thread done");
        });
        t.start();
        t.join(); // wait for t to finish
        System.out.println("main done");
    }
}
```
**A:**
```
thread done
main done
```
Without `join()`, `main done` might print before `thread done`.

---

### 11. Thread Lifecycle
**Q: What are the valid states of a Thread?**
```java
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> {
            try { Thread.sleep(1000); } catch (InterruptedException e) {}
        });
        System.out.println(t.getState()); // NEW
        t.start();
        System.out.println(t.getState()); // RUNNABLE or TIMED_WAITING
        t.join();
        System.out.println(t.getState()); // TERMINATED
    }
}
```
**A:**
```
NEW
RUNNABLE (or TIMED_WAITING)
TERMINATED
```
States: NEW → RUNNABLE → BLOCKED/WAITING/TIMED_WAITING → TERMINATED.

---

### 12. Daemon Thread
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread daemon = new Thread(() -> {
            while (true) {
                try { Thread.sleep(100); System.out.println("daemon running"); }
                catch (InterruptedException e) { break; }
            }
        });
        daemon.setDaemon(true); // must be set BEFORE start()
        daemon.start();
        Thread.sleep(250);
        System.out.println("main exiting");
        // JVM exits when all non-daemon threads finish — daemon is killed
    }
}
```
**A:** `daemon running` (twice), then `main exiting`. The JVM exits when only daemon threads remain. Useful for background tasks (GC, monitoring).

---

### 13. Thread.interrupt() and isInterrupted()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> {
            while (!Thread.currentThread().isInterrupted()) {
                // do work
            }
            System.out.println("interrupted: " + Thread.currentThread().isInterrupted());
        });
        t.start();
        Thread.sleep(100);
        t.interrupt();
        t.join();
    }
}
```
**A:** `interrupted: true`. `interrupt()` sets the interrupted flag. Check it with `isInterrupted()` for clean shutdown. Blocked methods (`sleep`, `wait`) throw `InterruptedException` when the flag is set.

---

### 14. synchronized on Different Objects — Not Mutually Exclusive
**Q: What is the bug?**
```java
public class Main {
    static class Counter {
        private int count = 0;
        void increment() {
            synchronized (new Object()) { // BUG: different object each time!
                count++;
            }
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Counter c = new Counter();
        Thread t1 = new Thread(() -> { for (int i = 0; i < 10000; i++) c.increment(); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 10000; i++) c.increment(); });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(c.count);
    }
}
```
**A:** Some value less than 20000. Each call creates a new lock object — threads never contend for the same lock. **Fix:** `synchronized (this)` or a dedicated shared lock object.

---

### 15. Thread-Safe Singleton — Double-Checked Locking
**Q: Is this correct?**
```java
public class Main {
    static class Singleton {
        private static volatile Singleton instance;
        private Singleton() {}
        static Singleton getInstance() {
            if (instance == null) {                    // first check (no lock)
                synchronized (Singleton.class) {
                    if (instance == null) {            // second check (with lock)
                        instance = new Singleton();
                    }
                }
            }
            return instance;
        }
    }
    public static void main(String[] args) {
        System.out.println(Singleton.getInstance() == Singleton.getInstance());
    }
}
```
**A:** `true`. Double-checked locking is correct **only with `volatile`** (Java 5+). `volatile` prevents instruction reordering that would expose a partially-initialized singleton.

---

## Section 2: java.util.concurrent Locking (Q16–Q28)

### 16. ReentrantLock — Explicit Locking
**Q: What is the output?**
```java
import java.util.concurrent.locks.*;
public class Main {
    static ReentrantLock lock = new ReentrantLock();
    static int counter = 0;

    static void increment() {
        lock.lock();
        try { counter++; }
        finally { lock.unlock(); } // always unlock in finally!
    }

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 10000; i++) increment(); });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(counter);
    }
}
```
**A:** `20000`. `ReentrantLock` gives more flexibility than `synchronized`: tryLock, timed lock, interruptible lock, fair locking.

---

### 17. ReentrantLock.tryLock() — Non-Blocking
**Q: What is the output?**
```java
import java.util.concurrent.locks.*;
public class Main {
    static ReentrantLock lock = new ReentrantLock();

    public static void main(String[] args) throws InterruptedException {
        lock.lock();
        Thread t = new Thread(() -> {
            boolean acquired = lock.tryLock(); // non-blocking
            System.out.println("acquired: " + acquired);
            if (acquired) lock.unlock();
        });
        t.start(); t.join();
        lock.unlock();
    }
}
```
**A:** `acquired: false`. `tryLock()` returns immediately — `false` if lock is held, `true` if acquired.

---

### 18. ReentrantLock — Reentrant Property
**Q: What is the output?**
```java
import java.util.concurrent.locks.*;
public class Main {
    static ReentrantLock lock = new ReentrantLock();

    static void outer() {
        lock.lock();
        try {
            System.out.println("outer: holdCount=" + lock.getHoldCount());
            inner();
        } finally { lock.unlock(); }
    }
    static void inner() {
        lock.lock();
        try { System.out.println("inner: holdCount=" + lock.getHoldCount()); }
        finally { lock.unlock(); }
    }

    public static void main(String[] args) { outer(); }
}
```
**A:**
```
outer: holdCount=1
inner: holdCount=2
```
`ReentrantLock` can be acquired multiple times by the same thread. It maintains a hold count.

---

### 19. ReadWriteLock — Concurrent Reads
**Q: What is the output?**
```java
import java.util.concurrent.locks.*;
public class Main {
    static ReadWriteLock rwLock = new ReentrantReadWriteLock();
    static int data = 100;

    static int read() {
        rwLock.readLock().lock();
        try { return data; }
        finally { rwLock.readLock().unlock(); }
    }

    static void write(int val) {
        rwLock.writeLock().lock();
        try { data = val; }
        finally { rwLock.writeLock().unlock(); }
    }

    public static void main(String[] args) {
        System.out.println(read());
        write(200);
        System.out.println(read());
    }
}
```
**A:**
```
100
200
```
`ReadWriteLock` allows multiple concurrent readers OR one exclusive writer. Read-heavy workloads are significantly faster than a plain `Lock`.

---

### 20. StampedLock — Optimistic Reading (Java 8+)
**Q: What is the pattern?**
```java
import java.util.concurrent.locks.*;
public class Main {
    static StampedLock lock = new StampedLock();
    static int data = 0;

    static int optimisticRead() {
        long stamp = lock.tryOptimisticRead();
        int value = data;
        if (!lock.validate(stamp)) {    // was there a write?
            stamp = lock.readLock();    // fall back to regular read lock
            try { value = data; }
            finally { lock.unlockRead(stamp); }
        }
        return value;
    }

    public static void main(String[] args) { System.out.println(optimisticRead()); }
}
```
**A:** `0`. `StampedLock.tryOptimisticRead()` is even faster than `ReadWriteLock` — it doesn't block, but you must validate the stamp afterwards.

---

### 21. CountDownLatch — Wait for Multiple Tasks
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch latch = new CountDownLatch(3);

        for (int i = 0; i < 3; i++) {
            int id = i;
            new Thread(() -> {
                System.out.println("task " + id + " done");
                latch.countDown();
            }).start();
        }

        latch.await(); // wait until count reaches zero
        System.out.println("all tasks done");
    }
}
```
**A:** Prints `task 0 done`, `task 1 done`, `task 2 done` (in some order), then `all tasks done`. `CountDownLatch` is a one-shot gate — can't be reused.

---

### 22. CyclicBarrier — Phases / All Meet Here
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        CyclicBarrier barrier = new CyclicBarrier(3, () -> System.out.println("--- all at barrier ---"));
        for (int i = 0; i < 3; i++) {
            int id = i;
            new Thread(() -> {
                System.out.println("thread " + id + " waiting");
                try { barrier.await(); }
                catch (Exception e) {}
                System.out.println("thread " + id + " past barrier");
            }).start();
        }
        Thread.sleep(500);
    }
}
```
**A:** Each thread prints "waiting", then `--- all at barrier ---`, then each "past barrier". `CyclicBarrier` can be reused (unlike `CountDownLatch`).

---

### 23. Semaphore — Rate Limiting
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static Semaphore sem = new Semaphore(2); // max 2 concurrent

    public static void main(String[] args) throws InterruptedException {
        for (int i = 0; i < 5; i++) {
            int id = i;
            new Thread(() -> {
                try {
                    sem.acquire();
                    System.out.println("in: " + id + " | permits=" + sem.availablePermits());
                    Thread.sleep(100);
                } catch (InterruptedException e) {}
                finally {
                    sem.release();
                    System.out.println("out: " + id);
                }
            }).start();
        }
        Thread.sleep(600);
    }
}
```
**A:** At most 2 threads run concurrently. `acquire()` blocks when no permits are available. `release()` returns a permit.

---

### 24. Exchanger — Swap Data Between Threads
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static Exchanger<String> exchanger = new Exchanger<>();

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> {
            try {
                String got = exchanger.exchange("from t1");
                System.out.println("t1 received: " + got);
            } catch (InterruptedException e) {}
        });
        Thread t2 = new Thread(() -> {
            try {
                String got = exchanger.exchange("from t2");
                System.out.println("t2 received: " + got);
            } catch (InterruptedException e) {}
        });
        t1.start(); t2.start();
        t1.join(); t2.join();
    }
}
```
**A:**
```
t1 received: from t2
t2 received: from t1
```
`Exchanger` synchronizes two threads and allows them to swap objects atomically.

---

### 25. Phaser — Advanced Synchronization (Java 7+)
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        Phaser phaser = new Phaser(3); // 3 parties
        for (int i = 0; i < 3; i++) {
            int id = i;
            new Thread(() -> {
                System.out.println("Phase 0: " + id);
                phaser.arriveAndAwaitAdvance(); // wait for all
                System.out.println("Phase 1: " + id);
                phaser.arriveAndDeregister(); // done
            }).start();
        }
    }
}
```
**A:** All threads print "Phase 0" (order varies), then all print "Phase 1". `Phaser` is a reusable, flexible replacement for `CyclicBarrier` + `CountDownLatch`.

---

### 26. LockSupport.park and unpark
**Q: What is the output?**
```java
import java.util.concurrent.locks.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> {
            System.out.println("parking...");
            LockSupport.park();
            System.out.println("unparked!");
        });
        t.start();
        Thread.sleep(100);
        LockSupport.unpark(t); // wake up t
        t.join();
    }
}
```
**A:**
```
parking...
unparked!
```
`LockSupport` is the building block for `AbstractQueuedSynchronizer` (the foundation of `ReentrantLock`, `Semaphore` etc.).

---

### 27. ThreadLocal — Per-Thread State
**Q: What is the output?**
```java
public class Main {
    static ThreadLocal<Integer> local = ThreadLocal.withInitial(() -> 0);

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> {
            local.set(100);
            System.out.println("t1: " + local.get());
        });
        Thread t2 = new Thread(() -> {
            local.set(200);
            System.out.println("t2: " + local.get());
        });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println("main: " + local.get()); // main thread's copy
    }
}
```
**A:**
```
t1: 100
t2: 200
main: 0
```
Each thread has its own `ThreadLocal` value. No synchronization needed. Used for per-request data (e.g., database connections, user sessions).

---

### 28. ThreadLocal Memory Leak in Thread Pools
**Q: What is the risk?**
```java
import java.util.concurrent.*;
public class Main {
    static ThreadLocal<byte[]> cache = new ThreadLocal<>();

    public static void main(String[] args) {
        ExecutorService pool = Executors.newFixedThreadPool(4);
        for (int i = 0; i < 100; i++) {
            pool.submit(() -> {
                cache.set(new byte[1024 * 1024]); // 1MB per thread
                // BUG: never calling cache.remove()!
            });
        }
        pool.shutdown();
        System.out.println("done — but thread pool threads still hold 1MB each!");
        // FIX: always call cache.remove() in a finally block
    }
}
```
**A:** **Memory leak.** Thread pool threads are reused — their `ThreadLocal` values persist. Always call `threadLocal.remove()` in a `finally` block when using `ThreadLocal` with thread pools.

---

## Section 3: Executor Framework & Thread Pools (Q29–Q40)

### 29. ExecutorService vs raw Thread
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        ExecutorService pool = Executors.newFixedThreadPool(2);

        for (int i = 0; i < 4; i++) {
            int id = i;
            pool.submit(() -> System.out.println("task " + id + " on " + Thread.currentThread().getName()));
        }
        pool.shutdown();
        pool.awaitTermination(1, TimeUnit.SECONDS);
    }
}
```
**A:** Prints 4 task lines, max 2 running simultaneously. Prefer `ExecutorService` over raw threads — manages thread lifecycle, reuse, and shutdown.

---

### 30. Future and Callable — Get Result
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        ExecutorService pool = Executors.newSingleThreadExecutor();
        Future<Integer> future = pool.submit(() -> {
            Thread.sleep(200);
            return 42;
        });
        System.out.println("submitted");
        int result = future.get(); // blocks until done
        System.out.println("result: " + result);
        pool.shutdown();
    }
}
```
**A:**
```
submitted
result: 42
```

---

### 31. Future.get() Timeout
**Q: What happens?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        ExecutorService pool = Executors.newSingleThreadExecutor();
        Future<Integer> future = pool.submit(() -> { Thread.sleep(2000); return 1; });
        try {
            future.get(500, TimeUnit.MILLISECONDS); // timeout!
        } catch (TimeoutException e) {
            System.out.println("Timed out!");
            future.cancel(true); // interrupt the task
        }
        pool.shutdown();
    }
}
```
**A:** `Timed out!`. `get(timeout)` throws `TimeoutException` if the task doesn't complete within the timeout.

---

### 32. invokeAll and invokeAny
**Q: What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        ExecutorService pool = Executors.newFixedThreadPool(3);
        List<Callable<String>> tasks = List.of(
            () -> { Thread.sleep(300); return "A"; },
            () -> { Thread.sleep(100); return "B"; },
            () -> { Thread.sleep(200); return "C"; }
        );
        // invokeAny: returns the result of the FIRST completed task
        String first = pool.invokeAny(tasks);
        System.out.println("First: " + first);
        pool.shutdown();
    }
}
```
**A:** `First: B`. `invokeAny` returns the first successful result and cancels the rest.

---

### 33. Thread Pool Types
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        // Fixed pool: bounded threads, unbounded queue
        ExecutorService fixed  = Executors.newFixedThreadPool(4);

        // Cached pool: grows on demand, idle threads die after 60s — dangerous for high load
        ExecutorService cached = Executors.newCachedThreadPool();

        // Single thread: tasks execute sequentially
        ExecutorService single = Executors.newSingleThreadExecutor();

        // Scheduled: for delayed/periodic tasks
        ScheduledExecutorService scheduled = Executors.newScheduledThreadPool(2);

        System.out.println("Thread pools created");
        fixed.shutdown(); cached.shutdown(); single.shutdown(); scheduled.shutdown();
    }
}
```
**A:** `Thread pools created`. Know the trade-offs: fixed = predictable resource usage; cached = auto-scaling but can OOM; scheduled = replaces `Timer`.

---

### 34. ScheduledExecutorService — Periodic Task
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);
        scheduler.scheduleAtFixedRate(
            () -> System.out.println("tick at " + System.currentTimeMillis() / 1000),
            0, 100, TimeUnit.MILLISECONDS
        );
        Thread.sleep(350);
        scheduler.shutdown();
    }
}
```
**A:** Prints "tick" approximately 4 times (at 0ms, 100ms, 200ms, 300ms). `scheduleAtFixedRate` runs exactly at fixed intervals regardless of task duration.

---

### 35. ForkJoinPool — Divide and Conquer
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static class SumTask extends RecursiveTask<Long> {
        int[] arr; int lo, hi;
        SumTask(int[] arr, int lo, int hi) { this.arr = arr; this.lo = lo; this.hi = hi; }

        @Override
        protected Long compute() {
            if (hi - lo <= 10) {
                long sum = 0; for (int i = lo; i < hi; i++) sum += arr[i]; return sum;
            }
            int mid = (lo + hi) / 2;
            SumTask left = new SumTask(arr, lo, mid);
            SumTask right = new SumTask(arr, mid, hi);
            left.fork();
            return right.compute() + left.join();
        }
    }
    public static void main(String[] args) {
        int[] arr = new int[100];
        for (int i = 0; i < 100; i++) arr[i] = i + 1;
        ForkJoinPool pool = new ForkJoinPool();
        System.out.println(pool.invoke(new SumTask(arr, 0, 100)));
    }
}
```
**A:** `5050`. ForkJoinPool uses work-stealing for divide-and-conquer parallelism. Used internally by parallel streams.

---

### 36. Executor — Reject Policy
**Q: What happens when the queue is full?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        ThreadPoolExecutor pool = new ThreadPoolExecutor(
            1, 1, 0L, TimeUnit.MILLISECONDS,
            new ArrayBlockingQueue<>(1), // queue of size 1
            new ThreadPoolExecutor.AbortPolicy() // throws on overflow
        );
        pool.submit(() -> { try { Thread.sleep(500); } catch (InterruptedException e) {} });
        pool.submit(() -> {}); // fills the queue
        try {
            pool.submit(() -> {}); // rejected! queue full, max threads reached
        } catch (RejectedExecutionException e) {
            System.out.println("Rejected!");
        }
        pool.shutdown();
    }
}
```
**A:** `Rejected!`. Know the 4 rejection policies: `AbortPolicy` (throw), `CallerRunsPolicy` (run in caller), `DiscardPolicy` (silently drop), `DiscardOldestPolicy` (drop oldest).

---

### 37. CompletableFuture Basics
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<String> cf = CompletableFuture.supplyAsync(() -> {
            return "hello";
        }).thenApply(String::toUpperCase)
          .thenApply(s -> s + "!");

        System.out.println(cf.get());
    }
}
```
**A:** `HELLO!`. `CompletableFuture` chains async steps. Each `thenApply` transforms the result.

---

### 38. CompletableFuture — thenCompose (flatMap)
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static CompletableFuture<String> fetchUser(int id) {
        return CompletableFuture.supplyAsync(() -> "User" + id);
    }
    static CompletableFuture<String> fetchOrders(String user) {
        return CompletableFuture.supplyAsync(() -> user + "'s orders: [O1, O2]");
    }
    public static void main(String[] args) throws Exception {
        String result = fetchUser(42)
                .thenCompose(Main::fetchOrders) // flatMap — chain dependent futures
                .get();
        System.out.println(result);
    }
}
```
**A:** `User42's orders: [O1, O2]`. `thenCompose` avoids `CompletableFuture<CompletableFuture<T>>` nesting.

---

### 39. CompletableFuture.allOf and anyOf
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<String> f1 = CompletableFuture.supplyAsync(() -> "A");
        CompletableFuture<String> f2 = CompletableFuture.supplyAsync(() -> "B");
        CompletableFuture<String> f3 = CompletableFuture.supplyAsync(() -> "C");

        CompletableFuture.allOf(f1, f2, f3).join(); // wait for all
        System.out.println(f1.get() + f2.get() + f3.get());

        // anyOf: returns result of first completed
        CompletableFuture<Object> first = CompletableFuture.anyOf(f1, f2, f3);
        System.out.println(first.get() != null);
    }
}
```
**A:**
```
ABC
true
```

---

### 40. CompletableFuture — exceptionally and handle
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<Integer> result = CompletableFuture.<Integer>supplyAsync(() -> {
                    throw new RuntimeException("oops");
                })
                .exceptionally(ex -> {
                    System.out.println("Error: " + ex.getMessage());
                    return -1; // default value
                });
        System.out.println(result.get());
    }
}
```
**A:**
```
Error: oops
-1
```
`exceptionally` handles errors and provides a fallback. `handle` is called in both success and failure cases.

---

## Section 4: Concurrent Collections (Q41–Q50)

### 41. ConcurrentHashMap — Thread-Safe Map
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
        Thread t1 = new Thread(() -> { for (int i = 0; i < 1000; i++) map.merge("key", 1, Integer::sum); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 1000; i++) map.merge("key", 1, Integer::sum); });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(map.get("key"));
    }
}
```
**A:** `2000`. `ConcurrentHashMap.merge()` is atomic. `ConcurrentHashMap` uses lock striping (not a single lock) for high concurrency.

---

### 42. ConcurrentHashMap Does NOT Lock on putIfAbsent
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
        // Both putIfAbsent and computeIfAbsent are atomic
        map.putIfAbsent("x", 1);
        map.putIfAbsent("x", 2); // not inserted — key exists
        System.out.println(map.get("x"));

        // computeIfAbsent is preferred for lazy-init patterns
        map.computeIfAbsent("list", k -> 42);
        System.out.println(map.get("list"));
    }
}
```
**A:**
```
1
42
```

---

### 43. CopyOnWriteArrayList — Safe Iteration
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        CopyOnWriteArrayList<Integer> list = new CopyOnWriteArrayList<>(new Integer[]{1, 2, 3});
        Thread writer = new Thread(() -> { list.add(4); list.add(5); });
        writer.start();
        // Iteration is on a snapshot — safe even while writer modifies
        for (int x : list) System.out.print(x + " ");
        writer.join();
        System.out.println("\nfinal: " + list);
    }
}
```
**A:** `1 2 3` (snapshot), then `final: [1, 2, 3, 4, 5]`. `CopyOnWriteArrayList` copies the array on every mutation — expensive for writes but safe for iteration.

---

### 44. LinkedBlockingQueue — Bounded Buffer
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        BlockingQueue<String> queue = new LinkedBlockingQueue<>(2);
        queue.put("task1");
        queue.put("task2");
        // queue.put("task3"); // would BLOCK — queue is full

        System.out.println(queue.take()); // removes task1
        System.out.println(queue.size());
    }
}
```
**A:**
```
task1
1
```

---

### 45. Producer-Consumer with BlockingQueue
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static BlockingQueue<Integer> queue = new ArrayBlockingQueue<>(3);

    public static void main(String[] args) throws InterruptedException {
        Thread producer = new Thread(() -> {
            for (int i = 1; i <= 5; i++) {
                try { queue.put(i); System.out.println("produced: " + i); }
                catch (InterruptedException e) { Thread.currentThread().interrupt(); }
            }
        });
        Thread consumer = new Thread(() -> {
            for (int i = 0; i < 5; i++) {
                try { int v = queue.take(); System.out.println("consumed: " + v); }
                catch (InterruptedException e) { Thread.currentThread().interrupt(); }
            }
        });
        producer.start(); consumer.start();
        producer.join(); consumer.join();
    }
}
```
**A:** Produces and consumes 1–5, naturally flow-controlling when queue is full/empty. Classic producer-consumer pattern with `BlockingQueue`.

---

### 46. ConcurrentLinkedQueue — Lock-Free
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        ConcurrentLinkedQueue<Integer> clq = new ConcurrentLinkedQueue<>();
        Thread t1 = new Thread(() -> { for (int i = 0; i < 5; i++) clq.offer(i); });
        Thread t2 = new Thread(() -> { for (int i = 10; i < 15; i++) clq.offer(i); });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(clq.size());
    }
}
```
**A:** `10`. `ConcurrentLinkedQueue` is a lock-free, non-blocking queue using CAS. Does not block, so `poll()` returns `null` if empty (unlike `blocking` queues).

---

### 47. AtomicReference — CAS on Objects
**Q: What is the output?**
```java
import java.util.concurrent.atomic.*;
public class Main {
    public static void main(String[] args) {
        AtomicReference<String> ref = new AtomicReference<>("initial");

        boolean swapped = ref.compareAndSet("initial", "updated"); // CAS
        System.out.println(swapped + " → " + ref.get());

        boolean swapped2 = ref.compareAndSet("initial", "again"); // old value mismatch
        System.out.println(swapped2 + " → " + ref.get());
    }
}
```
**A:**
```
true → updated
false → updated
```

---

### 48. AtomicLong.accumulateAndGet
**Q: What is the output?**
```java
import java.util.concurrent.atomic.*;
public class Main {
    public static void main(String[] args) {
        AtomicLong value = new AtomicLong(10);
        long result = value.accumulateAndGet(5, Long::sum); // atomic: 10 + 5
        System.out.println(result);
        System.out.println(value.get());
    }
}
```
**A:**
```
15
15
```

---

### 49. LongAdder vs AtomicLong — High Contention Counter
**Q: What is the output?**
```java
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        LongAdder adder = new LongAdder();
        Thread t1 = new Thread(() -> { for (int i = 0; i < 100000; i++) adder.increment(); });
        Thread t2 = new Thread(() -> { for (int i = 0; i < 100000; i++) adder.increment(); });
        t1.start(); t2.start(); t1.join(); t2.join();
        System.out.println(adder.sum());
    }
}
```
**A:** `200000`. `LongAdder` is faster than `AtomicLong` under high contention — it distributes counts across cells per thread, reducing CAS conflicts.

---

### 50. ConcurrentSkipListMap — Sorted + Concurrent
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        ConcurrentSkipListMap<Integer, String> map = new ConcurrentSkipListMap<>();
        map.put(3, "three"); map.put(1, "one"); map.put(4, "four"); map.put(2, "two");
        System.out.println(map); // sorted
        System.out.println(map.firstKey());
        System.out.println(map.floorKey(3));
    }
}
```
**A:**
```
{1=one, 2=two, 3=three, 4=four}
1
3
```
`ConcurrentSkipListMap` is a thread-safe `TreeMap` alternative with O(log n) concurrent operations.

---

## Section 5: CompletableFuture (Q51–Q60)

### 51. CompletableFuture.runAsync
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<Void> cf = CompletableFuture.runAsync(() -> {
            System.out.println("async: " + Thread.currentThread().getName());
        });
        cf.join();
        System.out.println("main: " + Thread.currentThread().getName());
    }
}
```
**A:**
```
async: ForkJoinPool.commonPool-worker-1
main: main
```
`runAsync` runs on ForkJoinPool's common pool by default. Use `supplyAsync(task, customExecutor)` for custom thread pools.

---

### 52. CompletableFuture — thenAccept and thenRun
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture.supplyAsync(() -> "result")
                .thenAccept(s -> System.out.println("consumed: " + s)) // has input, no output
                .thenRun(() -> System.out.println("done!"))              // no input, no output
                .join();
    }
}
```
**A:**
```
consumed: result
done!
```

---

### 53. CompletableFuture.thenCombine — Independent Futures
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<Integer> price    = CompletableFuture.supplyAsync(() -> 100);
        CompletableFuture<Integer> discount = CompletableFuture.supplyAsync(() -> 20);
        int finalPrice = price.thenCombine(discount, (p, d) -> p - d).get();
        System.out.println(finalPrice);
    }
}
```
**A:** `80`. `thenCombine` waits for two independent futures and combines their results.

---

### 54. CompletableFuture — whenComplete (always called)
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<String> cf = CompletableFuture.<String>supplyAsync(() -> {
                    throw new RuntimeException("error!");
                })
                .whenComplete((result, ex) -> {
                    if (ex != null) System.out.println("Exception: " + ex.getCause().getMessage());
                    else System.out.println("Result: " + result);
                });
        try { cf.join(); } catch (CompletionException e) {}
    }
}
```
**A:** `Exception: error!`. `whenComplete` is called regardless of success or failure — like a `finally` block.

---

### 55. CompletableFuture — Timeout (Java 9+)
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<String> cf = CompletableFuture.supplyAsync(() -> {
                    try { Thread.sleep(2000); } catch (InterruptedException e) {}
                    return "slow result";
                })
                .orTimeout(500, TimeUnit.MILLISECONDS) // Java 9+
                .exceptionally(ex -> "timeout fallback");
        System.out.println(cf.get());
    }
}
```
**A:** `timeout fallback`. `orTimeout` (Java 9+) automatically completes the future exceptionally if it doesn't finish in time.

---

### 56. CompletableFuture Chain — Real Pattern
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    static CompletableFuture<String> fetchUser(int id) {
        return CompletableFuture.supplyAsync(() -> "Alice");
    }
    static CompletableFuture<String[]> fetchRoles(String user) {
        return CompletableFuture.supplyAsync(() -> new String[]{"ADMIN", "USER"});
    }
    public static void main(String[] args) throws Exception {
        fetchUser(1)
            .thenCompose(Main::fetchRoles)
            .thenApply(roles -> String.join(", ", roles))
            .thenAccept(System.out::println)
            .join();
    }
}
```
**A:** `ADMIN, USER`

---

### 57. CompletableFuture.allOf — Wait for All with Results
**Q: What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) throws Exception {
        List<CompletableFuture<Integer>> futures = List.of(
            CompletableFuture.supplyAsync(() -> 1),
            CompletableFuture.supplyAsync(() -> 2),
            CompletableFuture.supplyAsync(() -> 3)
        );
        int sum = CompletableFuture.allOf(futures.toArray(new CompletableFuture[0]))
                .thenApply(v -> futures.stream().mapToInt(CompletableFuture::join).sum())
                .get();
        System.out.println(sum);
    }
}
```
**A:** `6`. `allOf` returns `Void` — a common pattern is to collect results from individual futures after `allOf` completes.

---

### 58. CompletableFuture vs Future — Non-Blocking
**Q: What is the key difference?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        // Old way: Future — blocks on get()
        ExecutorService pool = Executors.newSingleThreadExecutor();
        Future<Integer> blockingFuture = pool.submit(() -> 42);
        int x = blockingFuture.get(); // BLOCKS current thread

        // New way: CompletableFuture — non-blocking callback
        CompletableFuture<Integer> cf = CompletableFuture.supplyAsync(() -> 42);
        cf.thenAccept(v -> System.out.println("Got: " + v)); // callback, no blocking
        cf.join(); // wait for demo purposes
        pool.shutdown();
    }
}
```
**A:** `Got: 42`. `CompletableFuture` enables non-blocking pipelines via callbacks. `Future.get()` blocks the calling thread.

---

### 59. Exception Propagation in CompletableFuture Chain
**Q: What is the output?**
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) throws Exception {
        CompletableFuture<String> cf = CompletableFuture.supplyAsync(() -> "step1")
                .thenApply(s -> { throw new RuntimeException("step2 failed"); })
                .thenApply(s -> s + " step3") // skipped due to exception
                .exceptionally(ex -> "recovered: " + ex.getCause().getMessage());
        System.out.println(cf.get());
    }
}
```
**A:** `recovered: step2 failed`. Exceptions skip subsequent `thenApply` stages and jump directly to `exceptionally`.

---

### 60. Reactive-Style Aggregation Pattern
**Q: What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.stream.*;
public class Main {
    record Product(String name, double price) {}

    static CompletableFuture<List<Product>> fetchProducts() {
        return CompletableFuture.supplyAsync(() -> List.of(
            new Product("Apple", 1.5), new Product("Banana", 0.5), new Product("Cherry", 2.5)));
    }

    public static void main(String[] args) throws Exception {
        fetchProducts()
            .thenApply(products -> products.stream()
                .filter(p -> p.price() > 1.0)
                .map(Product::name)
                .collect(Collectors.joining(", ")))
            .thenAccept(names -> System.out.println("Expensive: " + names))
            .join();
    }
}
```
**A:** `Expensive: Apple, Cherry`
