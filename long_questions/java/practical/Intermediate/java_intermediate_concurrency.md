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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug in this code?
**Your Response:** This code has a classic race condition. The `counter++` operation looks atomic but it's actually three separate operations: read the current value, increment it, and write it back. When two threads execute this simultaneously, they might both read the same value (say 100), both increment to 101, and both write 101 back. So instead of getting 20000, we get some unpredictable lower number. The fix is to make this operation atomic using synchronization or atomic variables.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output now?
**Your Response:** Now we get exactly 20000. By adding the `synchronized` keyword to the static method, we're ensuring that only one thread can execute the increment method at a time. The lock is on the class object itself since it's a static method. This eliminates the race condition we saw before. The threads will take turns executing the increment method, so every increment is properly applied.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output here and how is this different from the previous approach?
**Your Response:** We still get 20000, but this approach is more efficient. Instead of locking the entire method, we're only locking the critical section where we modify the counter. This means other threads can execute the non-critical parts of the code concurrently. Using a synchronized block with a separate lock object gives us better performance because the lock is held for a shorter time.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does volatile do here?
**Your Response:** The volatile keyword ensures visibility of changes across threads. Without it, the worker thread might cache the `running` variable in its CPU cache and never see when the main thread sets it to false. With volatile, every read and write goes directly to main memory, so when the main thread sets `running` to false, the worker thread immediately sees the change and exits the loop. This is about visibility, not atomicity - volatile doesn't solve race conditions for compound operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the bug here even with volatile?
**Your Response:** Great question! Even though we made the counter volatile, we still have a race condition. Volatile only ensures visibility - that other threads will see the latest value. But it doesn't make the `counter++` operation atomic. It's still three separate operations: read, increment, write. Two threads can still read the same value, both increment, and both write back the same result. To fix this, we need either synchronized blocks or atomic classes like AtomicInteger.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and how does AtomicInteger work?
**Your Response:** We get exactly 20000. AtomicInteger uses lock-free atomic operations backed by CPU-level Compare-And-Swap instructions. Unlike synchronized blocks which use locks and can cause thread contention, AtomicInteger uses hardware-level atomic operations. The incrementAndGet method ensures that the read-increment-write operation happens atomically without any locks. This is much more efficient under high contention than synchronized blocks for simple atomic operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and how does wait/notify work?
**Your Response:** The output is "produced" then "consumed". This is a classic producer-consumer pattern using wait() and notifyAll(). The consumer thread starts first and calls wait(), which releases the lock and puts the thread to sleep. When the producer sets ready to true and calls notifyAll(), it wakes up the consumer. The consumer then resumes execution. The key points are that wait() releases the lock while the thread waits, and we always use a while loop around wait() to handle spurious wakeups - threads can sometimes wake up without being notified.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens in this code and how do you fix it?
**Your Response:** This is a classic deadlock scenario. Thread t1 acquires lock1 and then tries to get lock2, while thread t2 acquires lock2 and then tries to get lock1. Each thread is holding one lock that the other needs, so neither can proceed - they're stuck forever. The fix is simple: always acquire locks in the same consistent order. If both threads acquire lock1 first, then lock2, the deadlock is eliminated. This is why lock ordering discipline is crucial in multi-threaded programming.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain the difference between deadlock, livelock, and starvation?
**Your Response:** Absolutely! Deadlock is when threads are blocked waiting for each other - like two cars at a four-way stop, each waiting for the other to go. Livelock is more subtle - threads are actively running but keep responding to each other in a way that prevents progress, like two people in a hallway who keep stepping aside in the same direction. Starvation is when a thread never gets CPU time or access to a resource because other threads always get priority. The key difference: deadlock threads are blocked, livelock threads are active but stuck, starvation threads are perpetually denied access.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and what does join() do?
**Your Response:** The output is "thread done" then "main done". The join() method makes the main thread wait until the worker thread completes. Without join(), the main thread would continue immediately and we might see "main done" before "thread done" because the threads run concurrently. Join() is essentially a synchronization point - it blocks the calling thread until the target thread finishes execution. This is commonly used when you need to wait for background work to complete before proceeding.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the thread states and how does the lifecycle work?
**Your Response:** A thread goes through several states during its lifecycle. It starts as NEW when created but not started. After calling start(), it becomes RUNNABLE, meaning it's ready to run or currently running. While executing, it can enter BLOCKED (waiting for a lock), WAITING (indefinitely waiting with wait()), or TIMED_WAITING (sleeping or timed wait). Finally, when it finishes execution, it becomes TERMINATED. Understanding this lifecycle is crucial for debugging concurrency issues and managing thread behavior properly.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and what are daemon threads used for?
**Your Response:** The output shows "daemon running" twice followed by "main exiting". Daemon threads are background threads that don't prevent the JVM from shutting down. When the main thread finishes, the JVM will exit even if daemon threads are still running. This is different from regular user threads, which would keep the JVM alive. Daemon threads are perfect for background tasks like garbage collection, monitoring, or cleanup services that shouldn't block application shutdown.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does interrupt() do and how does it work?
**Your Response:** The interrupt() mechanism is Java's cooperative way to stop threads. When we call interrupt(), it sets an interrupted flag on the thread. The thread can check this flag using isInterrupted() and decide to stop gracefully. If the thread is blocked in methods like sleep() or wait(), it will immediately wake up and throw an InterruptedException. This allows threads to clean up resources before terminating. It's much better than the deprecated stop() method which would abruptly kill the thread.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the bug in this synchronization code?
**Your Response:** This is a subtle but critical bug. The code is creating a new lock object every time the increment method is called. So when thread A calls increment(), it creates one lock object and acquires it. When thread B calls increment(), it creates a completely different lock object and acquires that one. Since they're locking on different objects, there's no mutual exclusion - they can both execute the critical section simultaneously. The fix is to use a shared lock object - either synchronize on 'this' or use a final static lock object that all threads share.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is this double-checked locking implementation correct?
**Your Response:** Yes, this implementation is correct, but only because of the volatile keyword. The double-checked locking pattern avoids synchronization after the singleton is initialized, which improves performance. The first check without a lock lets most calls return quickly. Only when the instance is null do we acquire the lock and check again. The volatile is crucial - without it, the JVM could reorder instructions and another thread might see a partially constructed object. Since Java 5, volatile guarantees proper happens-before relationships, making this pattern thread-safe.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output and when would you use ReentrantLock over synchronized?
**Your Response:** The output is 20000, showing that ReentrantLock provides proper mutual exclusion. ReentrantLock is more flexible than synchronized - it offers features like tryLock() for non-blocking attempts, timed lock attempts, interruptible lock acquisition, and fair locking policies. It's particularly useful when you need to handle lock contention more gracefully, implement timeout-based locking, or build complex synchronization patterns. However, it requires manual unlock in a finally block, whereas synchronized handles this automatically.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does tryLock() do and when would you use it?
**Your Response:** tryLock() is a non-blocking attempt to acquire a lock. It immediately returns true if the lock is available, or false if another thread holds it. This is incredibly useful for avoiding deadlocks and implementing timeout-based operations. For example, when acquiring multiple locks, you can use tryLock() with timeouts to ensure you don't block indefinitely. It's also great for building responsive applications where you don't want threads to block for long periods.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does reentrant mean and why is it important?
**Your Response:** Reentrant means a thread can acquire the same lock multiple times without deadlocking itself. The lock maintains a hold count - each lock() call increments it, each unlock() decrements it, and the lock is only released when the count reaches zero. This is crucial because it allows a thread to call methods that also acquire the same lock. Without reentrancy, a synchronized method calling another synchronized method on the same object would deadlock. Both synchronized blocks and ReentrantLock are reentrant in Java.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use ReadWriteLock?
**Your Response:** ReadWriteLock is perfect for read-heavy workloads where data is read much more frequently than written. It allows multiple threads to read simultaneously, which dramatically improves throughput compared to a regular exclusive lock. The trade-off is that writes are more expensive - they need exclusive access and must wait for all readers to finish. This pattern is common in caching scenarios, configuration management, or any scenario where you have many readers and few writers. The performance improvement can be substantial when reads dominate.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is optimistic reading in StampedLock?
**Your Response:** Optimistic reading is a lock-free approach that's even faster than ReadWriteLock for read-heavy scenarios. It assumes there will be no writes while reading - you take an optimistic stamp, read the data, then validate the stamp. If validation fails, it means a write occurred, so you fall back to a regular read lock. This approach eliminates lock contention entirely when writes are rare, providing excellent performance for mostly-read scenarios. It's particularly useful in high-frequency trading, real-time analytics, or cache systems.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CountDownLatch and when would you use it?
**Your Response:** CountDownLatch is a synchronization aid that lets one or more threads wait until a set of operations being performed in other threads completes. It's initialized with a count, and each thread calls countDown() to decrement it. The await() method blocks until the count reaches zero. This is perfect for scenarios where you need to wait for multiple parallel tasks to complete before proceeding, like waiting for all services to initialize, or for all data processing jobs to finish. It's a one-time use synchronization point - once the count reaches zero, it can't be reset.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between CountDownLatch and CyclicBarrier?
**Your Response:** While both coordinate multiple threads, they serve different purposes. CountDownLatch is one-way - threads count down and waiting threads proceed once. CyclicBarrier is for multi-phase synchronization - threads wait at a barrier, execute a barrier action, then continue, and can reuse the barrier for multiple phases. Think of CountDownLatch as waiting for workers to complete, while CyclicBarrier is like getting all team members to meet at checkpoints during a race. CyclicBarrier is perfect for iterative algorithms that need synchronization between iterations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Semaphore used for?
**Your Response:** Semaphore controls access to a limited resource by maintaining a set of permits. Threads acquire permits before accessing the resource and release them when done. This is perfect for rate limiting, connection pooling, or any scenario where you want to limit concurrent access. Unlike locks which are binary (one thread at a time), semaphores can allow multiple threads up to the permit limit. It's also commonly used to implement producer-consumer patterns where producers add permits and consumers consume them.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does Exchanger do and when would you use it?
**Your Response:** Exchanger is a synchronization point where two threads can exchange objects. Both threads call exchange() and block until the other thread arrives, then they swap their objects atomically. This is perfect for scenarios like genetic algorithms where two organisms need to exchange genetic material, or in pipeline processing where stages need to swap data buffers. It's a specialized tool for two-way data exchange - for more participants, you'd use other concurrent utilities.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Phaser and how is it different from CyclicBarrier?
**Your Response:** Phaser is a more flexible synchronization barrier that can handle dynamic numbers of parties. Unlike CyclicBarrier which has a fixed number of parties, Phaser allows threads to register and deregister dynamically. It supports multiple phases like CyclicBarrier but also provides methods like arriveAndAwaitAdvance() and arriveAndDeregister(). Phaser is great for parallel algorithms where the number of worker threads might change during execution, or when you need complex multi-phase synchronization patterns. It's essentially the evolution of CyclicBarrier with more capabilities.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is LockSupport and why is it important?
**Your Response:** LockSupport provides low-level thread blocking primitives that are the foundation for most Java concurrency utilities. The park() method blocks a thread, and unpark() unblocks it. Unlike wait()/notify(), these operations don't require holding a lock and don't suffer from spurious wakeups. LockSupport is what powers AbstractQueuedSynchronizer, which in turn implements ReentrantLock, Semaphore, CountDownLatch, and most other java.util.concurrent classes. It's a building block for creating custom synchronization constructs when the high-level utilities don't fit your needs.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is ThreadLocal and when would you use it?
**Your Response:** ThreadLocal provides thread-scoped variables - each thread gets its own independent copy of the variable. This is perfect for data that shouldn't be shared between threads, like user authentication context, database connections per thread, or request-specific data in web applications. The main advantage is that you don't need synchronization since each thread only accesses its own copy. However, be careful with thread pools - ThreadLocal values can persist across requests and cause memory leaks if not properly cleaned up.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the risk with ThreadLocal in thread pools?
**Your Response:** ThreadLocal can cause serious memory leaks when used with thread pools. The issue is that thread pool threads are reused across multiple tasks, so ThreadLocal values set by one task persist and are visible to subsequent tasks. If you're storing large objects in ThreadLocal and not cleaning them up, each thread in the pool will accumulate memory that never gets garbage collected. The fix is to always call threadLocal.remove() in a finally block to clean up the value after each task completes. This is a common source of memory leaks in web applications using thread pools for request handling.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is ExecutorService better than creating raw threads?
**Your Response:** ExecutorService is much better than raw threads because it manages the entire thread lifecycle for you. Instead of manually creating and managing threads, you just submit tasks. The pool handles thread reuse, which avoids the overhead of creating new threads for each task. It also provides proper shutdown semantics and prevents resource leaks. Raw threads are error-prone - developers often forget to handle exceptions properly or don't shut down threads cleanly. ExecutorService gives you better control over resource usage and is the standard way to handle concurrent execution in production applications.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do Future and Callable work together?
**Your Response:** Future represents the result of an asynchronous computation. When you submit a Callable to an ExecutorService, you get a Future immediately. The Callable is like a Runnable that can return a value and throw exceptions. The Future.get() method blocks until the computation is complete, then returns the result. This is perfect for fire-and-forget tasks where you need the result later. You can also check if the task is done with isDone(), cancel it with cancel(), or get with timeout to avoid blocking indefinitely.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle timeouts with Future?
**Your Response:** Future.get() with timeout is essential for building responsive applications. Instead of blocking indefinitely, you specify a timeout, and if the task doesn't complete in time, it throws a TimeoutException. This prevents your application from hanging on slow or stuck tasks. When you get a timeout, you should cancel the task with future.cancel(true) to interrupt the underlying thread and free up resources. This pattern is crucial for web applications, microservices, or any system where you need to enforce response time guarantees.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between invokeAll and invokeAny?
**Your Response:** invokeAll waits for all tasks to complete and returns a list of Future objects, while invokeAny returns the result of the first task that completes successfully and cancels the remaining tasks. invokeAny is perfect for redundant operations like trying multiple servers or algorithms - you get the fastest successful result. invokeAll is for when you need all results to proceed. Both methods block, but invokeAny can return much faster if you have multiple redundant services and only need one successful response.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the different types of thread pools and when would you use each?
**Your Response:** Each thread pool type serves different needs. Fixed thread pools have a bounded number of threads with an unbounded queue - great for predictable resource usage and controlling concurrency. Cached pools grow on demand and kill idle threads after 60 seconds - good for bursty workloads but can cause OOM under sustained high load. Single thread executor ensures sequential execution - perfect for when you need ordering guarantees. Scheduled executor handles delayed and periodic tasks, replacing the old Timer class with better thread management. The choice depends on your workload characteristics and resource constraints.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between scheduleAtFixedRate and scheduleWithFixedDelay?
**Your Response:** scheduleAtFixedRate runs at fixed intervals regardless of how long the task takes - if a task takes longer than the interval, it can run concurrently or queue up. scheduleWithFixedDelay waits for the task to complete, then waits the delay period before scheduling the next one. Use scheduleAtFixedRate for things like clock ticks or polling where you want consistent timing. Use scheduleWithFixedDelay for tasks where you want to ensure they don't overlap, like processing batches or cleanup operations. Both are much better than the old Timer class which has limitations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does ForkJoinPool work and when would you use it?
**Your Response:** ForkJoinPool is designed for divide-and-conquer algorithms using work-stealing. Each worker thread has its own task queue, and when it runs out of work, it can steal tasks from other threads' queues. This maximizes CPU utilization for recursive problems. The key methods are fork() to schedule a task asynchronously and join() to wait for its result. ForkJoinPool is what powers Java's parallel streams behind the scenes. It's perfect for recursive algorithms like quicksort, merge sort, or processing large data sets that can be split into independent subproblems.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when a thread pool's queue is full?
**Your Response:** When a thread pool can't accept more tasks, it applies its rejection policy. AbortPolicy throws RejectedExecutionException, which is the default. CallerRunsPolicy runs the task in the caller's thread, providing backpressure. DiscardPolicy silently drops the task, which is dangerous but sometimes acceptable for non-critical work. DiscardOldestPolicy drops the oldest queued task to make room. The choice depends on your requirements - CallerRunsPolicy is often best for providing natural throttling, while AbortPolicy ensures you don't silently lose work.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does CompletableFuture chaining work?
**Your Response:** CompletableFuture allows you to chain asynchronous operations in a fluent API. supplyAsync() starts the computation, then methods like thenApply() transform the result when it's ready. Each step runs asynchronously after the previous one completes. This creates a pipeline of operations that execute without blocking threads. The beauty is that you can compose complex async workflows that are both readable and efficient. CompletableFuture is much more powerful than Future - it supports composition, exception handling, and non-blocking completion.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between thenApply and thenCompose?
**Your Response:** thenApply transforms the result of a CompletableFuture to a new value, while thenCompose is used when the transformation itself returns another CompletableFuture. thenCompose is like flatMap - it flattens nested CompletableFutures so you don't end up with CompletableFuture<CompletableFuture<T>>. Use thenApply for synchronous transformations, and thenCompose when you need to chain asynchronous operations that depend on each other. This distinction is crucial for building clean async workflows without nested callbacks.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do allOf and anyOf work with CompletableFuture?
**Your Response:** allOf waits for all provided CompletableFutures to complete, which is perfect for parallel execution when you need all results before proceeding. anyOf returns when any of the futures completes, giving you the result of the fastest one. This is useful for redundancy - trying multiple services and using the first response. Both methods return a CompletableFuture that completes when their condition is met. allOf is like a barrier for async operations, while anyOf is like a race condition where the winner's result is used.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle exceptions in CompletableFuture?
**Your Response:** CompletableFuture provides several ways to handle exceptions. exceptionally() is like a catch block - it only executes when there's an exception and lets you provide a fallback value. handle() is more like finally - it gets called regardless of success or failure, receiving both the result and any exception. You can also use whenComplete() for side effects that need to happen regardless of outcome. This exception handling is much cleaner than try-catch blocks with Future.get(), making async error handling composable and readable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does ConcurrentHashMap achieve thread safety?
**Your Response:** ConcurrentHashMap uses lock striping instead of a single lock, which allows multiple threads to read and write simultaneously without blocking each other. It divides the map into segments and locks only the segment being modified. Methods like merge(), computeIfAbsent(), and putIfAbsent() are atomic operations. This design provides much better scalability than synchronized HashMap or Hashtable. In Java 8+, they improved it further with finer-grained locking at the bucket level. It's the go-to choice for concurrent maps in production.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the atomic operations in ConcurrentHashMap?
**Your Response:** ConcurrentHashMap provides several atomic compound operations that are crucial for thread-safe programming. putIfAbsent() atomically inserts a value only if the key doesn't exist. computeIfAbsent() is even more powerful - it atomically computes a value if the key is absent, perfect for lazy initialization patterns. merge() atomically combines existing values with new ones. These operations eliminate the race conditions you'd have with check-then-act sequences using regular maps. They're essential for building concurrent data structures and caching systems.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use CopyOnWriteArrayList?
**Your Response:** CopyOnWriteArrayList is perfect for scenarios with many reads and few writes. It creates a new copy of the underlying array every time you modify it, which means iterators never see ConcurrentModificationException and don't need synchronization. This makes it ideal for configuration data, listener lists, or any scenario where you iterate frequently but modify rarely. The trade-off is expensive writes - every add or operation copies the entire array. But if you have thousands of readers and occasional updates, it's much faster than synchronized ArrayList.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do BlockingQueues work?
**Your Response:** BlockingQueue provides thread-safe coordination between producers and consumers. The put() method blocks when the queue is full, preventing producers from overwhelming consumers. The take() method blocks when empty, preventing consumers from spinning when there's no work. This natural backpressure makes it perfect for producer-consumer patterns. LinkedBlockingQueue can be bounded or unbounded, while ArrayBlockingQueue is always bounded. BlockingQueues eliminate all the complexity of manual wait/notify coordination for common producer-consumer scenarios.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this producer-consumer pattern work?
**Your Response:** This demonstrates the classic producer-consumer pattern using BlockingQueue. The producer calls put() which blocks when the queue is full, naturally throttling production. The consumer calls take() which blocks when empty, preventing busy waiting. This coordination happens automatically without any explicit synchronization - the BlockingQueue handles all the wait/notify logic internally. When the queue fills up, the producer waits; when it empties, the consumer waits. This pattern is fundamental to building concurrent systems and is much simpler than manual coordination.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between ConcurrentLinkedQueue and BlockingQueue?
**Your Response:** ConcurrentLinkedQueue is a lock-free, non-blocking queue that uses Compare-And-Swap operations under the hood. Unlike BlockingQueue, its operations never block - offer() always returns immediately, and poll() returns null if empty instead of waiting. This makes it faster when you don't need the blocking behavior, but you have to handle empty queues yourself. It's perfect for high-throughput scenarios where you can poll in a loop or use it in patterns that don't require waiting. BlockingQueue is better when you want threads to wait naturally for work.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does AtomicReference work?
**Your Response:** AtomicReference provides atomic operations on object references using Compare-And-Swap. The compareAndSet() method atomically updates the reference only if the current value matches what you expect. This prevents race conditions in check-then-act sequences. In the example, the first CAS succeeds because the current value is 'initial', but the second fails because the value was already changed to 'updated'. AtomicReference is perfect for building lock-free data structures, managing state transitions, or implementing optimistic locking patterns.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does accumulateAndGet do?
**Your Response:** accumulateAndGet is an atomic operation that applies a function to the current value and a new value, updating the atomically. In this case, it atomically adds 5 to the current value of 10, resulting in 15. This is more powerful than simple increment operations because you can apply any function, like multiplication, taking min/max, or custom business logic. It's part of Java 8's enhanced atomic classes that support functional operations, making atomic code more expressive and readable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use LongAdder over AtomicLong?
**Your Response:** LongAdder is designed for high-contention scenarios where many threads are incrementing a counter. Unlike AtomicLong which uses a single CAS operation that threads compete for, LongAdder distributes the increments across multiple cells, one per thread initially. This dramatically reduces contention under high load. When you need the final value, it sums all the cells. Use LongAdder for statistics like request counts in high-traffic servers. Use AtomicLong when you need the exact current value frequently or when contention is low.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use ConcurrentSkipListMap?
**Your Response:** ConcurrentSkipListMap is the thread-safe version of TreeMap that maintains sorted order while supporting concurrent operations. It uses a skip list data structure which provides O(log n) performance for most operations. Unlike ConcurrentHashMap which doesn't maintain ordering, ConcurrentSkipListMap keeps keys sorted and provides navigation methods like firstKey(), floorKey(), and subMap(). It's perfect for scenarios where you need both thread safety and ordered access, like leaderboards, time-series data, or range queries. The trade-off is slightly higher memory usage than HashMap-based structures.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does runAsync work and where does it execute?
**Your Response:** runAsync executes a Runnable task asynchronously and returns a CompletableFuture. By default, it uses the ForkJoinPool.commonPool(), which is a shared thread pool for the entire JVM. This is efficient for most async operations but you can provide a custom Executor if you need specific thread management. runAsync is perfect for fire-and-forget operations where you don't need a return value. For operations that return values, use supplyAsync instead. The common pool automatically scales based on available processors, making it ideal for CPU-bound tasks.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between thenAccept and thenRun?
**Your Response:** thenAccept consumes the result of the previous stage without producing a new result - it's like a consumer that performs side effects. thenRun doesn't even look at the previous result - it just runs an action when the previous stage completes. Use thenAccept when you need to process the result, like saving to a database or sending a notification. Use thenRun for cleanup or finalization tasks that need to happen regardless of the result. Both return CompletableFuture<Void> since they don't produce new values for the chain.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does thenCombine work?
**Your Response:** thenCombine is perfect for combining results from independent asynchronous operations. It waits for both futures to complete, then applies a function to combine their results. This is different from thenCompose which is for dependent operations. Use thenCombine when you can run operations in parallel and need both results, like fetching user data and order details simultaneously. The combining function receives both results and produces a new value. This pattern is essential for building efficient async workflows that maximize parallelism.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between whenComplete and exceptionally?
**Your Response:** whenComplete is like a finally block - it gets called regardless of whether the computation succeeded or failed. It receives both the result and any exception, but can't change the outcome. exceptionally is like a catch block - it only executes when there's an exception and can provide a recovery value. Use whenComplete for logging, cleanup, or side effects that need to happen every time. Use exceptionally for error recovery and fallback values. You can chain them to get both behaviors: whenComplete for logging, then exceptionally for recovery.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does timeout handling work in CompletableFuture?
**Your Response:** orTimeout is a Java 9+ feature that adds automatic timeout capability to CompletableFuture. If the computation doesn't complete within the specified time, the future completes exceptionally with TimeoutException. This is much cleaner than manually implementing timeout logic with separate threads. You can combine it with exceptionally to provide fallback values when timeouts occur. This is essential for building resilient systems that need to enforce SLAs and prevent operations from hanging indefinitely.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this CompletableFuture chain demonstrate?
**Your Response:** This chain shows a real-world async workflow pattern. We start by fetching a user, then use thenCompose to fetch roles based on that user (dependent operation), then transform the roles array into a formatted string, then consume the result by printing it. This demonstrates how to build complex async pipelines that are both readable and efficient. Each step runs after the previous one completes, and the entire chain is non-blocking. This pattern is common in microservices where you need to orchestrate multiple async calls sequentially.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you get results from allOf?
**Your Response:** allOf itself returns a CompletableFuture<Void> that completes when all input futures complete, but it doesn't directly give you the results. The common pattern is to call allOf, then in a thenApply, iterate through the original futures and call join() on each to collect their results. This works because by the time thenApply runs, all futures are guaranteed to be complete, so join() won't block. This pattern is perfect for aggregating results from parallel operations, like calculating totals or building composite objects from multiple service calls.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the key advantage of CompletableFuture over Future?
**Your Response:** The key advantage is that CompletableFuture enables non-blocking asynchronous programming through callbacks and composition. With traditional Future, you can only check if it's done or block with get(). CompletableFuture allows you to attach callbacks that execute when the result is available, without blocking any threads. This enables building reactive, event-driven systems that scale better under load. You can compose operations, handle exceptions gracefully, and create complex async pipelines - all things that are difficult or impossible with the blocking nature of Future.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does exception propagation work in CompletableFuture chains?
**Your Response:** When an exception occurs in a CompletableFuture chain, it skips all subsequent transformation stages like thenApply and jumps directly to the next exception handler like exceptionally or handle. This makes error handling predictable - you don't need to check for exceptions at each step. The exception propagates through the chain until it finds a handler. This is similar to how exceptions work in regular synchronous code, but with the benefit of asynchronous execution. You can place exception handlers strategically in your chain to recover from specific types of errors.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What pattern does this CompletableFuture example demonstrate?
**Your Response:** This demonstrates a reactive-style data processing pattern using CompletableFuture. We start by fetching products asynchronously, then transform the results using stream operations to filter expensive items and extract names. The beauty is that this entire pipeline is non-blocking - each stage executes when the previous one completes. This pattern combines the power of CompletableFuture for async operations with Java Streams for data transformation. It's perfect for building responsive applications that need to process data from multiple sources without blocking threads.
