# ⚡ Java Concurrency Practice
Contains runnable code examples for missing concurrency questions (based on Short Questions Q19-Q25).

## Question 1: Difference between `Thread` and `Runnable`?

### Answer
Thread (Class) vs Runnable (Interface).

### Runnable Code
```java
package concurrency;

// 1. Extend Thread
class MyThread extends Thread {
    public void run() { System.out.println("Thread Class Running"); }
}

// 2. Implement Runnable
class MyTask implements Runnable {
    public void run() { System.out.println("Runnable Task Running"); }
}

public class ThreadVsRunnable {
    public static void main(String[] args) {
        new MyThread().start();
        
        new Thread(new MyTask()).start();
        
        // Lambda (Modern Runnable)
        new Thread(() -> System.out.println("Lambda Running")).start();
    }
}
```

---

## Question 2: What is a deadlock? How do you create one?

### Answer
Circular dependency of locks.

### Runnable Code
```java
package concurrency;

public class DeadlockDemo {
    static final Object lock1 = new Object();
    static final Object lock2 = new Object();

    public static void main(String[] args) {
        Thread t1 = new Thread(() -> {
            synchronized (lock1) {
                System.out.println("Thread 1: Locked 1");
                try { Thread.sleep(50); } catch (Exception e) {}
                synchronized (lock2) {
                    System.out.println("Thread 1: Locked 2");
                }
            }
        });

        Thread t2 = new Thread(() -> {
            synchronized (lock2) {
                System.out.println("Thread 2: Locked 2");
                try { Thread.sleep(50); } catch (Exception e) {}
                synchronized (lock1) { // Waiting for execution of lock1
                    System.out.println("Thread 2: Locked 1");
                }
            }
        });

        t1.start();
        t2.start();
        // Result: Program hangs indefinitely
    }
}
```

---

## Question 3: What is `ExecutorService`?

### Answer
Thread Pool management.

### Runnable Code
```java
package concurrency;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ExecutorDemo {
    public static void main(String[] args) {
        // Fixed pool of 2 threads
        ExecutorService pool = Executors.newFixedThreadPool(2);
        
        for (int i = 0; i < 5; i++) {
            final int id = i;
            pool.submit(() -> {
                System.out.println("Task " + id + " running on " + Thread.currentThread().getName());
            });
        }
        
        pool.shutdown();
    }
}
```

---

## Question 4: Difference between `Callable` and `Runnable`?

### Answer
`Callable` returns value + throws Exception.

### Runnable Code
```java
package concurrency;

import java.util.concurrent.*;

public class CallableDemo {
    public static void main(String[] args) throws Exception {
        ExecutorService pool = Executors.newCachedThreadPool();
        
        // Callable returning String
        Callable<String> task = () -> {
            Thread.sleep(100);
            return "Task Complete";
        };
        
        Future<String> future = pool.submit(task);
        
        System.out.println("Waiting...");
        String result = future.get(); // Blocks until done
        System.out.println("Result: " + result);
        
        pool.shutdown();
    }
}
```

---

## Question 5: What are atomic classes (`AtomicInteger`)?

### Answer
CAS (Lock-free) thread safety.

### Runnable Code
```java
package concurrency;

import java.util.concurrent.atomic.AtomicInteger;

public class AtomicDemo {
    static int unsafeCount = 0;
    static AtomicInteger safeCount = new AtomicInteger(0);
    
    public static void main(String[] args) throws Exception {
        Runnable r = () -> {
            for(int i=0; i<1000; i++) {
                unsafeCount++;
                safeCount.incrementAndGet();
            }
        };
        
        Thread t1 = new Thread(r);
        Thread t2 = new Thread(r);
        t1.start(); t2.start();
        t1.join(); t2.join();
        
        System.out.println("Unsafe: " + unsafeCount); // < 2000 (likely)
        System.out.println("Safe:   " + safeCount);   // 2000 (always)
    }
}
```

---

## Question 6: Thread-safe Singleton (Double-Checked Locking).

### Answer
Standard interview implementation.

### Runnable Code
```java
package concurrency;

class DCLSingleton {
    private static volatile DCLSingleton instance;
    
    private DCLSingleton() {}
    
    public static DCLSingleton getInstance() {
        if (instance == null) { // 1st Check
            synchronized (DCLSingleton.class) {
                if (instance == null) { // 2nd Check
                    instance = new DCLSingleton();
                }
            }
        }
        return instance;
    }
}

public class SingletonDCL {
    public static void main(String[] args) {
        DCLSingleton s = DCLSingleton.getInstance();
    }
}
```
