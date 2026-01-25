# Java Multithreading & Concurrency Complete Guide

## 1. Creating Threads
Three main ways to create a thread in Java.

### A. Extending Thread Class
```java
class MyThread extends Thread {
    public void run() {
        System.out.println("Thread running");
    }
}
// Usage: new MyThread().start();
```

### B. Implementing Runnable Interface (Preferred)
Better because you can extend another class.
```java
class MyTask implements Runnable {
    public void run() {
        System.out.println("Task running");
    }
}
// Usage: new Thread(new MyTask()).start();
```

### C. Implementing Callable (Returing value)
Used with `ExecutorService`.
```java
Callable<String> task = () -> "Result";
Future<String> future = executor.submit(task);
String result = future.get(); // Blocks until result is ready
```

## 2. Executor Framework (Java 5+)
Stop manually creating `new Thread()`. Use Thread Pools.

### Types of Pools
1.  **FixedThreadPool(n)**: Reuses `n` threads. Good for server apps.
2.  **CachedThreadPool()**: Creates new threads as needed, kills idle ones. Good for many short-lived tasks.
3.  **SingleThreadExecutor()**: Ensures tasks are executed sequentially.
4.  **ScheduledThreadPool()**: For periodic tasks.

```java
ExecutorService executor = Executors.newFixedThreadPool(10);
executor.submit(() -> System.out.println("Task 1"));
executor.shutdown();
```

## 3. CompletableFuture (Java 8+)
Asynchronous programming (Promises).

```java
CompletableFuture.supplyAsync(() -> "Hello")
    .thenApply(s -> s + " World")
    .thenAccept(System.out::println);
// Non-blocking!
```

## 4. Virtual Threads (Java 21+)
Project Loom. Lightweight threads managed by JVM, not OS.
*   **Goal**: Create millions of threads.
*   **Performance**: High throughput for I/O bound applications.
```java
Thread.startVirtualThread(() -> {
    System.out.println("I am a virtual thread");
});
```

## 5. Synchronization & Locks

### synchronized keyword
Implicit lock (Monitor).
```java
public synchronized void increment() { count++; }
```

### ReentrantLock
Explicit lock. More flexible (tryLock, lockInterruptibly).
```java
Lock lock = new ReentrantLock();
lock.lock();
try {
    // critical section
} finally {
    lock.unlock();
}
```

### ReadWriteLock
Allows multiple readers, single writer.
```java
ReadWriteLock rwLock = new ReentrantReadWriteLock();
// rwLock.readLock().lock();
// rwLock.writeLock().lock();
```

## 6. Interview Questions
1.  **Difference between `submit()` and `execute()`?**
    *   `execute(Runnable)`: Fire and forget. No return value.
    *   `submit(Callable)`: Returns a `Future` to handle correct/exception.
2.  **What is a Daemon Thread?**
    *   Low priority thread (GC is a daemon). JVM exits when only Daemon threads are left.
3.  **Difference between `wait()` and `sleep()`?**
    *   `wait()`: Releases the lock. Must be in synchronized block.
    *   `sleep()`: Keeps the lock. Just pauses execution.
