# Specialized Topic 2: Multithreading Deep Dive

**Goal**: Move beyond `Thread.start()` and learn modern concurrency tools (`ExecutorService`, `CompletableFuture`).

## 1. ExecutorService (Thread Pools)
**Why?** Creating threads is expensive. Thread pools reuse threads.

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

class Task implements Runnable {
    String name;
    Task(String name) { this.name = name; }
    public void run() {
        System.out.println("Thread: " + Thread.currentThread().getName() + " executing " + name);
        try { Thread.sleep(500); } catch (Exception e) {}
    }   
}

public class ThreadPoolDemo {
    public static void main(String[] args) {
        // Create a pool of 3 threads
        ExecutorService pool = Executors.newFixedThreadPool(3);

        // Submit 5 tasks
        for (int i = 1; i <= 5; i++) {
            pool.execute(new Task("Task " + i));
        }

        // Shutdown pool
        pool.shutdown();
    }
}
```

## 2. Callable & Future (Returning Values)
**Why?** `Runnable` returns void. `Callable` returns a result.

```java
import java.util.concurrent.*;

class SumTask implements Callable<Integer> {
    int num;
    SumTask(int num) { this.num = num; }
    
    public Integer call() throws Exception {
        int sum = 0;
        for (int i = 1; i <= num; i++) sum += i;
        Thread.sleep(1000); // Simulate delay
        return sum;
    }
}

public class FutureDemo {
    public static void main(String[] args) throws Exception {
        ExecutorService pool = Executors.newCachedThreadPool();
        
        Future<Integer> future = pool.submit(new SumTask(10));
        
        System.out.println("Calculating...");
        
        // get() blocks until the result is available
        Integer result = future.get(); 
        
        System.out.println("Sum of 10 numbers: " + result);
        pool.shutdown();
    }
}
```

## 3. CompletableFuture (Async Programming)
**Why?** Non-blocking chains of actions (Java 8+).

```java
import java.util.concurrent.CompletableFuture;

public class CompletableDemo {
    public static void main(String[] args) {
        CompletableFuture.supplyAsync(() -> {
            try { Thread.sleep(1000); } catch (InterruptedException e) {}
            return "Hello";
        }).thenApply(msg -> {
            return msg + " World";
        }).thenAccept(finalMsg -> {
            System.out.println("Result: " + finalMsg);
        }).join(); // Wait for completion
    }
}
```

## Interview Questions
1.  **Runnable vs Callable?**
    *   Runnable: return `void`, cannot throw checked exception.
    *   Callable: returns `V`, can throw Exception.
2.  **execute() vs submit()?**
    *   `execute()`: For Runnable (fire and forget).
    *   `submit()`: For Callable/Runnable, returns `Future`.
3.  **Why use Thread Pool?**
    *   Low overhead (reuse threads), control resource consumption (limit max threads).

---

## 📋 Comprehensive Interview Questions

### **Thread Pool & ExecutorService Questions**

**Q1: Explain different types of thread pools available in Executors class.**
**A**: "Executors provides several thread pool types: `newFixedThreadPool()` creates a fixed number of threads and uses an unbounded queue - good for controlled concurrency. `newCachedThreadPool()` creates threads as needed and reuses idle ones - great for bursty workloads but can create unlimited threads. `newSingleThreadExecutor()` uses a single thread with an unbounded queue - ensures sequential execution. `newScheduledThreadPool()` supports delayed and periodic task execution. `newWorkStealingPool()` (Java 8+) uses fork-join pool for work-stealing. I choose based on workload characteristics and resource constraints."

**Q2: What is the difference between execute() and submit() in ExecutorService?**
**A**: "`execute()` is for fire-and-forget tasks that return void - it takes a Runnable and doesn't return anything. If the task throws an exception, it's handled by the uncaught exception handler. `submit()` is more versatile - it can take both Runnable and Callable, and returns a Future object. The Future allows me to check if the task is complete, wait for completion, get the result, or cancel the task. `submit()` also gives me access to exceptions through Future.get()."

**Q3: How does ThreadPoolExecutor work internally?**
**A**: "ThreadPoolExecutor uses a blocking queue to manage tasks. When I submit a task, if the pool has idle threads, the task runs immediately. If all threads are busy and the queue isn't full, the task waits in the queue. If the queue is full, it creates new threads up to the maximum pool size. If the queue is full and we're at max threads, it uses the rejection policy. The core pool size is the minimum threads that stay alive, while maximum threads can grow and shrink based on demand. This design balances throughput with resource usage."

### **Callable & Future Questions**

**Q4: Explain the difference between Runnable and Callable in detail.**
**A**: "Runnable is the original interface for tasks that don't return values - its run() method returns void and can't throw checked exceptions. Callable was introduced in Java 5 for tasks that need to return results - its call() method returns a generic type and can throw checked exceptions. Runnable is simpler for fire-and-forget tasks, while Callable is essential when I need results or need to handle exceptions explicitly. Future represents the result of an asynchronous computation and works with both."

**Q5: How do you handle exceptions in tasks submitted to ExecutorService?**
**A**: "For Runnable tasks, uncaught exceptions go to the thread's default exception handler. For Callable tasks, exceptions are wrapped in ExecutionException and thrown when I call Future.get(). I can handle them by catching ExecutionException and checking the cause. For better exception handling, I can use a custom ThreadFactory that sets uncaught exception handlers, or wrap tasks in error-handling code. In production, I'd log exceptions properly and implement retry logic for transient failures."

**Q6: What is FutureTask and how does it work?**
**A**: "FutureTask is the concrete implementation of Future that also implements Runnable. It acts as a bridge between Runnable and Callable - I can wrap a Callable in FutureTask and submit it to an ExecutorService that only accepts Runnable. FutureTask maintains the state of the computation (pending, running, completed, cancelled) and handles synchronization between the thread that computes the result and threads waiting for it. It's useful when I need to cancel tasks or check their status."

### **CompletableFuture & Async Programming Questions**

**Q7: Explain CompletableFuture and its advantages over Future.**
**A**: "CompletableFuture is a major improvement over Future - it allows me to chain asynchronous operations without blocking. While Future requires me to block with get(), CompletableFuture provides non-blocking methods like thenAccept(), thenApply(), and thenCompose(). It supports exception handling with exceptionally() and handle(), allows combining multiple futures, and can be completed programmatically. It represents the future of asynchronous programming in Java, making code more readable and efficient."

**Q8: What are the different ways to create CompletableFuture?**
**A**: "I can create CompletableFuture in several ways: `supplyAsync()` for tasks that return values, `runAsync()` for tasks that don't return values, `completedFuture()` for already-known results, or by creating an empty CompletableFuture and completing it later. I can also use `allOf()` to combine multiple futures or `anyOf()` to wait for the first completion. The async versions run on the common ForkJoinPool by default, but I can specify my own Executor."

**Q9: How do you handle exceptions in CompletableFuture chains?**
**A**: "CompletableFuture provides several exception handling methods: `exceptionally()` handles exceptions and can provide fallback values, `whenComplete()` gets both result and exception regardless of success/failure, and `handle()` gets both and can transform either case. I can chain these operations to create robust error handling. The exception propagation follows the chain - if any stage fails, subsequent stages are skipped unless they're exception handlers."

### **Synchronization & Concurrency Questions**

**Q10: What is the difference between synchronized block and ReentrantLock?**
**A**: "Synchronized blocks are built-in Java synchronization - they're simpler but less flexible. ReentrantLock is more powerful - it supports timed waits, interruptible locks, fair ordering, and multiple condition variables. synchronized automatically releases locks on exceptions, while with ReentrantLock I need explicit unlock in finally blocks. ReentrantLock also provides better performance under contention and supports tryLock() to avoid blocking. For simple cases, synchronized is fine; for complex scenarios, ReentrantLock offers more control."

**Q11: Explain the Java Memory Model and its importance in concurrency.**
**A**: "The Java Memory Model defines how threads interact through memory and what behaviors are guaranteed. Without proper memory model guarantees, threads could see stale data or operations could appear out of order. The model defines happens-before relationships that ensure visibility of changes between threads. Keywords like volatile, synchronized, and final establish these relationships. Understanding JMM is crucial for writing correct concurrent code - it prevents subtle bugs like seeing partially constructed objects or reordering issues."

**Q12: What is the difference between volatile and synchronized?**
**A**: "volatile ensures visibility of changes between threads - reads and writes go directly to main memory. It also prevents certain reordering but doesn't provide atomicity for compound actions. synchronized provides both visibility and mutual exclusion - only one thread can execute a synchronized block at a time, and changes are flushed to main memory. volatile is lighter weight and suitable for simple flags or status variables, while synchronized is needed for compound operations or critical sections requiring atomicity."

### **Advanced Concurrency Questions**

**Q13: Explain the Fork/Join Framework and when to use it.**
**A**: "The Fork/Join Framework is designed for divide-and-conquer algorithms that can be broken into smaller subtasks. It uses a work-stealing pool where idle threads steal tasks from busy threads' queues. I extend RecursiveTask for tasks that return values or RecursiveAction for void tasks. The framework is ideal for recursive problems like parallel sorting, tree traversal, or matrix operations. It's more efficient than general thread pools for CPU-intensive, recursively splittable workloads."

**Q14: What are atomic variables and how do they work?**
**A**: "Atomic variables like AtomicInteger, AtomicLong, and AtomicReference provide lock-free thread-safe operations on single variables. They use CAS (Compare-And-Swap) operations at the hardware level to ensure atomicity without locks. They provide methods like getAndSet(), getAndIncrement(), and compareAndSet(). They're more efficient than synchronized for simple operations and avoid thread blocking. However, they only work for single-variable operations - for compound actions, I still need synchronization."

**Q15: How do you implement producer-consumer pattern in Java?**
**A**: "I can implement producer-consumer using several approaches: 1) BlockingQueue is the simplest - producers put() items, consumers take() items, and the queue handles all synchronization. 2) wait()/notify() with shared buffer - more complex but educational. 3) Semaphore for controlling access to resources. 4) Disruptor pattern for ultra-high performance. BlockingQueue is usually the best choice - it's efficient, well-tested, and handles edge cases like buffer full/empty conditions properly."

### **Performance & Optimization Questions**

**Q16: How do you choose the optimal thread pool size?**
**A**: "For CPU-intensive tasks, I typically use `Runtime.getRuntime().availableProcessors()` threads - one per core. For I/O-intensive tasks, I use more threads since threads spend time waiting for I/O. A common formula is `threads = cores * (1 + wait_time/compute_time)`. I also consider memory usage, database connection limits, and other resource constraints. I monitor the application and adjust based on actual performance metrics rather than theoretical calculations."

**Q17: What are common performance pitfalls in concurrent programming?**
**A**: "Common pitfalls include: 1) Excessive synchronization leading to contention, 2) Creating too many threads causing context switching overhead, 3) Lock contention where multiple threads compete for the same locks, 4) False sharing where threads modify different variables on the same cache line, 5) Memory leaks from thread-local storage or thread pools, 6) Deadlock from inconsistent lock ordering. I profile the application to identify bottlenecks and use appropriate synchronization strategies."

**Q18: How do you debug concurrent programs?**
**A**: "Debugging concurrent programs requires special techniques: 1) Use logging with thread information to trace execution flow, 2) Use thread dumps and stack traces to see what threads are doing, 3) Use tools like JVisualVM or YourKit to monitor thread states and contention, 4) Use assertions and invariants to detect race conditions, 5) Reproduce issues deterministically by controlling thread scheduling, 6) Use formal verification tools for critical sections. Systematic testing with various timing scenarios is essential."

### **Real-World Scenarios Questions**

**Q19: How would you implement a rate limiter using concurrency utilities?**
**A**: "I'd implement a rate limiter using Semaphore or atomic counters. For a simple rate limiter, I could use a Semaphore with permits equal to the rate limit, releasing permits periodically. For more sophisticated rate limiting, I'd use a sliding window approach with atomic operations and time-based calculations. I could also use Guava's RateLimiter which implements token bucket algorithm. The key is ensuring thread-safe rate tracking while minimizing contention."

**Q20: How do you handle timeout in concurrent operations?**
**A**: "For timeouts, I use several approaches: 1) Future.get(timeout) with explicit timeout values, 2) ExecutorService.awaitTermination() for shutdown timeouts, 3) Lock.tryLock(timeout) for lock acquisition timeouts, 4) CompletableFuture.orTimeout() or completeOnTimeout() for async operations. I always implement timeout handling to prevent threads from blocking indefinitely. I also consider timeout policies - whether to fail fast, retry, or use fallback values."

**Q21: What is the difference between parallel streams and thread pools?**
**A**: "Parallel streams use the common ForkJoinPool by default, which is shared across the application. Thread pools I create are dedicated to specific tasks. Parallel streams are good for data-parallel operations on collections, while thread pools are better for task-parallel operations. Parallel streams have limitations - they can't be easily controlled or monitored, and they use a shared pool which can interfere with other operations. For production applications, I often prefer dedicated thread pools for better control."

**Q22: How do you implement graceful shutdown of a multi-threaded application?**
**A**: "For graceful shutdown, I: 1) Use shutdown hooks with Runtime.getRuntime().addShutdownHook(), 2) Call ExecutorService.shutdown() to stop accepting new tasks, 3) Call awaitTermination() with timeout to wait for running tasks, 4) Force shutdown with shutdownNow() if timeout expires, 5) Interrupt threads that respond to interruption, 6) Clean up resources in finally blocks, 7) Log shutdown progress. This ensures all in-progress work completes cleanly while preventing new work from starting."
