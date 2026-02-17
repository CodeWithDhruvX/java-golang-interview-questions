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
