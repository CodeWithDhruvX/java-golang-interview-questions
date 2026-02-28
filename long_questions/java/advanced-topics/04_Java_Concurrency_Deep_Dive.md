# Java Concurrency & Multithreading Deep Dive

Product-based companies (especially those handling massive throughput like Amazon, Uber, Atlassian) require a very deep understanding of the Java Memory Model, threading primitives, and the `java.util.concurrent` package.

---

## 🧠 1. Core Concurrency Concepts & Must-Knows

### 1. Threads vs. Processes
*   **Process:** An executing instance of an application with its own memory space (heap). Heavyweight.
*   **Thread:** A path of execution within a process that shares the same heap but has its own stack and program counter. Lightweight. Context switching between threads is faster.

### 2. Thread Lifecycle States
1.  **NEW:** Created but `start()` not invoked.
2.  **RUNNABLE:** Ready to run and waiting for CPU allocation or currently executing.
3.  **BLOCKED:** Waiting for a monitor lock to enter/re-enter a `synchronized` block/method.
4.  **WAITING:** Waiting indefinitely for another thread to perform a specific action (e.g., `Object.wait()`, `Thread.join()`).
5.  **TIMED_WAITING:** Waiting for another thread for a specified time (e.g., `Thread.sleep(ms)`, `Object.wait(ms)`).
6.  **TERMINATED:** Exited run method either normally or exceptionally.

### 3. The Java Memory Model (JMM)
JMM decides how threads interact with memory (Heap vs. CPU Caches).
*   **Visibility Problem:** A thread caches a variable (e.g., a flag) in its CPU L1/L2 cache and doesn't see updates made by another thread writing to the main memory.
    *   *Solution:* Use `volatile`. This guarantees visibility. Reads and writes bypass the CPU cache and go straight to the main memory (RAM).
*   **Atomicity Problem:** Operations like `count++` are actually 3 steps (Read, Increment, Write). Multiple threads can interleave these steps, leading to lost updates.
    *   *Solution:* Synchronization (Locks) or `AtomicInteger` (uses Compare-And-Swap/CAS).

### 4. Locks & Synchronization
*   **Intrinsic Locks (`synchronized`):** Every object in Java has a built-in monitor lock. It is reentrant. You can lock on the method (`this` instance) or the class (`Class.class` for static). Downsides: no timeout mechanism, cannot be interrupted.
*   **Explicit Locks (`ReentrantLock`, `ReentrantReadWriteLock`):** Part of `java.util.concurrent.locks`. They offer fairness (longest waiting thread gets lock first), ability to interrupt, and `tryLock(timeout)` mechanisms to avoid infinite blocking. A `ReadWriteLock` allows multiple concurrent readers but only one exclusive writer.

### 5. Advanced Primitives (`java.util.concurrent`)
You must be able to explain the exact difference between these:
*   **`CountDownLatch`:** initialized with a count. Threads waiting on it are blocked until other threads call `countDown()` enough times to hit 0. *Cannot be reset.* (Use case: Wait for 3 worker threads to finish DB queries before aggregating the result).
*   **`CyclicBarrier`:** initialized with a count. Threads call `await()` and wait until N threads reach the barrier. Once N threads arrive, they all proceed simultaneously. *Can be reset/reused.* (Use case: Parallelizing a massive calculation in phases).
*   **`Semaphore`:** Controls access to a shared resource using permits. Initialize with N permits. Threads call `acquire()` and `release()`. Useful for implementing **Rate Limiters** or limiting connections to a slow external API.
*   **`CompletableFuture`:** The modern standard for asynchronous programming in Java (since Java 8). Allows chaining, combining (`thenCombine`), and exceptionally handling non-blocking callbacks without writing "Callback Hell".

---

## 🔥 2. Top Frequently Asked Concurrency Questions

### Question 1: How exactly does `ConcurrentHashMap` work? Why is it better than `Hashtable` or `Collections.synchronizedMap`?
**Answer:**
`Hashtable` and `Collections.synchronizedMap` lock the *entire table* during every read/write. This causes immense contention in high-traffic applications.
`ConcurrentHashMap` in Java 8+ uses **Node/Bucket level locking** (instead of segment locking in Java 7).
*   **Reads (get):** Completely lock-free. Variables inside the node are `volatile`.
*   **Writes (put):** It locks only the specific array index (bucket) it is modifying using a `synchronized` block on the head node of that bucket. Other threads can write to different buckets concurrently.
*   **CAS Operations:** Uses `Compare-And-Swap` (CAS) to safely insert the very first node into an empty bucket without even acquiring a lock.

### Question 2: What is the difference between `sleep()`, `wait()`, and `yield()`?
**Answer:**
*   `Thread.sleep(ms)`: Pauses execution for a specified time. It **does NOT release the lock** it holds.
*   `Object.wait()`: Can only be called from within a `synchronized` block. It **releases the lock**, allowing other threads to acquire it, and waits until another thread calls `notify()`/`notifyAll()`.
*   `Thread.yield()`: A hint to the thread scheduler that the current thread is willing to yield its current use of the CPU. The scheduler can ignore this. It does NOT release any locks.

### Question 3: Write code to create a Deadlock and then explain how to resolve it.
**Answer:**
A deadlock happens when Thread A holds Lock 1 and waits for Lock 2, while Thread B holds Lock 2 and waits for Lock 1.

*The Code (Causing Deadlock):*
```java
public class DeadlockDemo {
    private final Object lock1 = new Object();
    private final Object lock2 = new Object();

    public void methodA() {
        synchronized (lock1) {
            System.out.println("Thread 1: Holding lock 1...");
            try { Thread.sleep(10); } catch (InterruptedException e) {}
            System.out.println("Thread 1: Waiting for lock 2...");
            synchronized (lock2) {
                System.out.println("Thread 1: Holding lock 1 & 2...");
            }
        }
    }

    public void methodB() {
        synchronized (lock2) {
            System.out.println("Thread 2: Holding lock 2...");
            try { Thread.sleep(10); } catch (InterruptedException e) {}
            System.out.println("Thread 2: Waiting for lock 1...");
            synchronized (lock1) {
                System.out.println("Thread 2: Holding lock 1 & 2...");
            }
        }
    }
}
```
*How to Resolve/Prevent it:*
1.  **Lock Ordering:** Ensure all threads acquire locks in the *exact same order* (e.g., both methodA and methodB must acquire lock1 first, then lock2). This breaks the circular wait condition.
2.  **Lock Timeout:** Use `ReentrantLock.tryLock(timeout)`. If it fails to acquire the second lock within a specific time, it backs off and releases the first lock.

### Question 4: Executor Framework Deep Dive: `FixedThreadPool` vs `CachedThreadPool`. When would you use which?
**Answer:**
Both are implementations of `ExecutorService` created via the `Executors` factory class.
*   **`Executors.newFixedThreadPool(n)`:**
    *   Has a core pool size of `n` and max pool size of `n`.
    *   Uses an **unbounded queue** (`LinkedBlockingQueue`). If all `n` threads are busy, new tasks are queued indefinitely.
    *   *Risk:* Can cause `OutOfMemoryError` if tasks pile up faster than they can be processed. Best for CPU-intensive tasks with predictable load.
*   **`Executors.newCachedThreadPool()`:**
    *   Core pool size is 0, Max pool size is `Integer.MAX_VALUE`.
    *   Uses a **synchronous queue** (a queue with size 0). If a task arrives and all existing threads are busy, it immediately spawns a *new* thread. Idle threads are terminated after 60 seconds.
    *   *Risk:* Can crash the system by creating thousands of threads if a massive spike of tasks occurs. Best for many short-lived, asynchronous tasks where execution time is very small.

### Question 5: Implement a classic Producer-Consumer using `wait()` and `notify()`.
**(Often a live coding round question)**

**Answer:**
Consider using `BlockingQueue` in production, but interviewers want to see the underlying mechanics.
```java
import java.util.LinkedList;
import java.util.Queue;

public class ProducerConsumer {
    private final Queue<Integer> buffer = new LinkedList<>();
    private final int CAPACITY = 5;

    public void produce(int value) throws InterruptedException {
        synchronized (this) {
            while (buffer.size() == CAPACITY) { // ALWAYS loop, never 'if' to prevent spurious wakeups
                wait(); // Releases the lock
            }
            buffer.add(value);
            System.out.println("Produced: " + value);
            notifyAll(); // Wake up consumers
        }
    }

    public void consume() throws InterruptedException {
        synchronized (this) {
            while (buffer.isEmpty()) {
                wait(); // Releases the lock
            }
            int value = buffer.poll();
            System.out.println("Consumed: " + value);
            notifyAll(); // Wake up producers
        }
    }
}
```

### Question 6: What is `ThreadLocal`? Have you ever used it?
**Answer:**
`ThreadLocal` provides thread-local variables. Each thread accessing the variable via `.get()` or `.set()` has its own independent, isolated copy of the variable.
*   **Use Cases:** Simple ways to maintain state for a single request in a web server without passing context objects down every method signature (e.g., storing the correlation ID or User Session Info during an HTTP request in Spring Boot).
*   **Memory Leaks:** If you use Thread Pools (like Tomcat does), a thread is reused. If you don't call `threadLocal.remove()` at the end of the request, the old request's data bleeds into the new request, leading to memory leaks and security issues.
