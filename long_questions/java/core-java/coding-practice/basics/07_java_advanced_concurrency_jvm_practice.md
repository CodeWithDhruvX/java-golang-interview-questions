# 🚀 Java Advanced Concurrency & JVM Practice
Contains runnable code for advanced topics: Async Programming, Explicit Locks, JVM Internals, and Modern Threading.

## Question 1: How to use `CompletableFuture` for asynchronous programming?

### Answer
Chaining tasks without blocking main thread. `supplyAsync`, `thenApply`, `thenAccept`.

### Runnable Code
```java
package advanced;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

public class CompletableFutureDemo {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        System.out.println("Main: " + Thread.currentThread().getName());
        
        CompletableFuture<Void> future = CompletableFuture.supplyAsync(() -> {
            // Task 1: Fetch Data
            System.out.println("Fetching: " + Thread.currentThread().getName());
            try { Thread.sleep(500); } catch (Exception e) {}
            return "Data";
        }).thenApply(data -> {
            // Task 2: Process Data
            System.out.println("Processing: " + Thread.currentThread().getName());
            return "Processed " + data;
        }).thenAccept(result -> {
            // Task 3: Consume Result
            System.out.println("Result: " + result);
        });
        
        future.get(); // Block for demo purposes only
    }
}
```

---

## Question 2: Difference between `synchronized` and `ReentrantLock`?

### Answer
`ReentrantLock` offers fairness, `tryLock` (timeout), and ability to interrupt waiting threads.

### Runnable Code
```java
package advanced;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.concurrent.TimeUnit;

public class LockDemo {
    private final Lock lock = new ReentrantLock();

    public void safeMethod() {
        // tryLock avoids deadlock by not waiting forever
        try {
            if (lock.tryLock(100, TimeUnit.MILLISECONDS)) {
                try {
                    System.out.println("Lock acquired by " + Thread.currentThread().getName());
                    Thread.sleep(50);
                } finally {
                    lock.unlock();
                }
            } else {
                System.out.println("Could not acquire lock: " + Thread.currentThread().getName());
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }

    public static void main(String[] args) {
        LockDemo demo = new LockDemo();
        Runnable r = demo::safeMethod;
        new Thread(r).start();
        new Thread(r).start();
    }
}
```

---

## Question 3: What is `ThreadLocal` and when to use it?

### Answer
Per-thread storage. Useful for passing context (User ID, Transaction ID) without method arguments. **Warning:** Must remove to avoid memory leaks in thread pools.

### Runnable Code
```java
package advanced;

public class ThreadLocalDemo {
    // Each thread gets its own independent SimpleDateFormat or value
    private static final ThreadLocal<Integer> userContext = ThreadLocal.withInitial(() -> 0);

    public static void main(String[] args) {
        Runnable r = () -> {
            int id = (int) (Math.random() * 100);
            userContext.set(id);
            System.out.println(Thread.currentThread().getName() + " set ID: " + userContext.get());
            
            try { Thread.sleep(50); } catch (Exception e) {}
            
            System.out.println(Thread.currentThread().getName() + " get ID: " + userContext.get());
            userContext.remove(); // Importance cleanup
        };

        new Thread(r, "T1").start();
        new Thread(r, "T2").start();
    }
}
```

---

## Question 4: How does a `ClassLoader` work? (Delegation Model)

### Answer
Bootstrap -> Platform (Ext) -> App (System). Delegates request to parent first.

### Runnable Code
```java
package advanced;

public class ClassLoaderDemo {
    public static void main(String[] args) {
        // App ClassLoader
        ClassLoader appLoader = ClassLoaderDemo.class.getClassLoader();
        System.out.println("App Loader: " + appLoader);
        
        // Platform Loader (JDK classes)
        ClassLoader platformLoader = appLoader.getParent();
        System.out.println("Parent (Platform): " + platformLoader);
        
        // Bootstrap Loader (Core Java) - represented as null in some implementations
        ClassLoader bootstrapLoader = platformLoader.getParent();
        System.out.println("Grandparent (Bootstrap): " + bootstrapLoader);
        
        // Loading a class manually
        try {
            Class<?> c = appLoader.loadClass("java.util.List");
            System.out.println("Loaded: " + c.getName());
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
    }
}
```

---

## Question 5: What are Virtual Threads (Java 21+)?

### Answer
Lighweight threads managed by JVM (not OS). High throughput for I/O bound apps.

### Runnable Code
```java
package advanced;

import java.util.concurrent.Executors;

public class VirtualThreadDemo {
    public static void main(String[] args) {
        // Only works on JDK 19+ (Preview) or JDK 21+
        try {
            // 1. Create primitive
            Thread.ofVirtual().start(() -> System.out.println("Virtual Thread Running"));
            
            // 2. Executor for Virtual Threads
            try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
                for (int i = 0; i < 5; i++) {
                    int index = i;
                    executor.submit(() -> {
                        System.out.println("Task " + index + " on " + Thread.currentThread());
                    });
                }
            } // Executor auto-closes and waits for tasks
            
        } catch (NoSuchMethodError e) {
            System.out.println("This JDK does not support Virtual Threads. Please run on JDK 21+.");
        }
    }
}
```

---

## Question 6: What is the `ReadWriteLock`?

### Answer
Allow multiple readers, single writer. Improves performance for read-heavy data.

### Runnable Code
```java
package advanced;

import java.util.concurrent.locks.*;
import java.util.HashMap;
import java.util.Map;

public class ReadWriteLockDemo {
    private final Map<String, String> map = new HashMap<>();
    private final ReadWriteLock rwLock = new ReentrantReadWriteLock();
    private final Lock readLock = rwLock.readLock();
    private final Lock writeLock = rwLock.writeLock();

    public void put(String key, String value) {
        writeLock.lock();
        try {
            System.out.println("Writing " + key);
            try { Thread.sleep(100); } catch (Exception e) {}
            map.put(key, value);
        } finally {
            writeLock.unlock();
        }
    }

    public String get(String key) {
        readLock.lock();
        try {
            System.out.println("Reading " + key);
            return map.get(key);
        } finally {
            readLock.unlock();
        }
    }

    public static void main(String[] args) {
        ReadWriteLockDemo cache = new ReadWriteLockDemo();
        
        new Thread(() -> cache.put("key", "val")).start();
        new Thread(() -> cache.get("key")).start(); // Can run parallel with other readers if implemented
        new Thread(() -> cache.get("key")).start();
    }
}
```
