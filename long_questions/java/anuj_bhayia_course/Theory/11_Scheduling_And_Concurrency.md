# Scheduling and Java Executor Framework - Interview Questions and Answers

## 1. How do you implement task scheduling (Cron Jobs) in a Spring Boot application?
**Answer:**
Spring provides built-in support for executing tasks dynamically at specific times or intervals using the `@Scheduled` annotation.

**Implementation Steps:**
1. **Enable Scheduling:** Add the `@EnableScheduling` annotation to your main application class or a configuration class. This tells Spring's background task executor to search for methods annotated with `@Scheduled`.
2. **Annotate Methods:** Create a Spring Bean (e.g., a `@Service` or `@Component`) and annotate a method with `@Scheduled`.
   *The method must return `void` and must not accept any arguments.*

**Configuration Options for `@Scheduled`:**
- **Fixed Rate:** Runs the task at a fixed interval, starting from the *start* of the previous execution.
  ```java
  @Scheduled(fixedRate = 5000) // Runs every 5 seconds
  public void runTask() { ... }
  ```
- **Fixed Delay:** Runs the task at a fixed interval, starting from the *completion* of the previous execution. Useful if the task duration varies and you don't want them overlapping.
  ```java
  @Scheduled(fixedDelay = 5000) // Waits 5s after the last run finished
  public void runTask() { ... }
  ```
- **Cron Expression:** Provides fine-grained scheduling (Seconds, Minutes, Hours, Day of Month, Month, Day of Week).
  ```java
  @Scheduled(cron = "0 0 12 * * ?") // Runs every day at 12:00 PM
  public void runCronTask() { ... }
  ```

## 2. If a `@Scheduled` task takes 10 seconds to run, but its `fixedRate` is set to 5 seconds, what happens by default in Spring Boot?
**Answer:**
By default, Spring Boot's internal task scheduler uses a **single-threaded pool** (`ThreadPoolTaskScheduler` with a pool size of 1).

**What happens:**
Even though the `fixedRate` dictates the task should fire every 5 seconds, because there is only one thread, **the executions will not overlap**.
1. At T=0s, Execution A starts.
2. At T=5s, Execution B is scheduled to start, but the single thread is still busy with Execution A. Execution B is blocked/queued.
3. At T=10s, Execution A finishes. The thread immediately picks up Execution B.

**How to fix overlapping needs:**
If you truly need them to run concurrently (e.g., overlapping is safe and necessary for throughput), you must:
1. Annotate the method with `@Async` as well as `@Scheduled`.
2. Configure a thread pool with more than one thread.

## 3. Explain the Java Executor Framework. Why is it preferred over manually creating `Thread` objects?
**Answer:**
The **Java Executor Framework** (introduced in java.util.concurrent in Java 5) is a robust mechanism for managing the execution of asynchronous tasks. It abstracts away the low-level details of thread creation, lifecycle management, and queuing.

**Why it's preferred over `new Thread(Runnable).start()`:**
1. **Thread Lifecycle Overhead:** Creating and destroying an OS-level thread is extremely expensive in terms of CPU and memory. The Executor framework uses **Thread Pools**, which reuse existing threads for new tasks, drastically reducing this overhead.
2. **Resource Management:** If you spawn a new thread for every incoming web request without limits, the JVM will quickly run out of memory (OutOfMemoryError) and crash. Executors allow you to define maximum limits (e.g., maximum 50 concurrent threads) and queue excess incoming tasks.
3. **Return Values and Exceptions:** Standard Runnables cannot return a value or throw checked exceptions easily. The Executor framework introduces `Callable<T>`, which returns a `Future<T>`, allowing the main thread to retrieve the result of the background task once it completes or cleanly catch exceptions.
4. **Convenience:** It provides built-in mechanisms for shutting down existing threads gracefully and waiting for tasks to finish (`shutdown()`, `awaitTermination()`).

## 4. What is a `ThreadPoolTaskExecutor` in Spring, and how do you configure it for `@Async` methods?
**Answer:**
While the standard Java `ExecutorService` works well, Spring provides its own abstraction, `ThreadPoolTaskExecutor`, which exposes Java's `ThreadPoolExecutor` as a Spring Bean with easy configuration properties.

**Usage with `@Async`:**
When you annotate a method with `@Async`, Spring creates a proxy. When the main thread calls that method, it returns immediately, and Spring submits the actual method execution to a background `TaskExecutor`.

**Configuration:**
If you don't configure a custom executor, Spring Boot auto-configures a default one (`SimpleAsyncTaskExecutor`), which creates a *new thread for every task* (defeating the purpose of thread pooling). You must define a custom one for production:

```java
@Configuration
@EnableAsync // Required to enable @Async processing
public class AsyncConfig {

    @Bean(name = "taskExecutor")
    public Executor taskExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(5);   // Minimum threads always kept alive
        executor.setMaxPoolSize(20);  // Maximum threads allowed if queue is full
        executor.setQueueCapacity(100); // How many tasks can wait before creating new threads up to Max
        executor.setThreadNamePrefix("AsyncThread-");
        executor.initialize();
        return executor;
    }
}
```

## 5. What are `Runnable` vs `Callable`? What is a `Future` or `CompletableFuture`?
**Answer:**

**1. `Runnable` vs `Callable` (The Tasks):**
- **`Runnable`:** An interface with a single `void run()` method. It executes a task but cannot return a result to the caller and cannot throw checked exceptions.
- **`Callable<T>`:** An interface with a single `T call() throws Exception` method. It is designed to be executed by an Executor, can return a result of type `T`, and can throw checked exceptions.

**2. `Future` vs `CompletableFuture` (The Results):**
- **`Future<T>`:** When you submit a `Callable` to an Executor, it immediately returns a `Future`. A `Future` represents the pending result of the asynchronous computation. The caller can use `future.get()` to retrieve the result, but this call is **blocking** (the main thread freezes until the background task finishes).
- **`CompletableFuture<T>`:** Introduced in Java 8, it is an advanced implementation of `Future`. Instead of blocking with `.get()`, it allows you to chain non-blocking callbacks. You can say: "When this task completes, then pass its result to Method B, and if it fails, trigger Method C."

**Example in Spring (`@Async`):**
```java
@Async("taskExecutor")
public CompletableFuture<User> fetchUserAsync(Long id) {
    // heavy database call
    User user = userRepository.findById(id);
    return CompletableFuture.completedFuture(user);
}

// In the caller service:
userService.fetchUserAsync(1L).thenAccept(user -> {
    System.out.println("User fetched in background: " + user.getName());
});
// Main thread continues execution immediately without blocking.
```
