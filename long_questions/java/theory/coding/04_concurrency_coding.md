# Coding: Concurrency - Interview Answers

> 🎯 **Focus:** Multithreading code is tricky. Focus on safety (locks) and modern tools (`java.util.concurrent`).

### 1. Implement Producer-Consumer
"I won't use `wait()` and `notify()` because they are low-level and error-prone.
Instead, I'll use a `BlockingQueue`. It handles all the thread safety and waiting logic internally.
The Producer `put()`s (blocking if full), and the Consumer `take()`s (blocking if empty)."

```java
BlockingQueue<Integer> queue = new ArrayBlockingQueue<>(10);

// Producer
Runnable producer = () -> {
    try {
        while (true) queue.put(new Random().nextInt(100));
    } catch (InterruptedException e) { Thread.currentThread().interrupt(); }
};

// Consumer
Runnable consumer = () -> {
    try {
        while (true) System.out.println("Consumed: " + queue.take());
    } catch (InterruptedException e) { Thread.currentThread().interrupt(); }
};
```

---

### 2. Create a Deadlock
"A deadlock happens when Thread A holds Lock 1 and wants Lock 2, while Thread B holds Lock 2 and wants Lock 1.
I can demonstrate this by nesting synchronized blocks in reverse order."

```java
Object lock1 = new Object();
Object lock2 = new Object();

Thread t1 = new Thread(() -> {
    synchronized(lock1) {
        try { Thread.sleep(100); } catch (Exception e) {}
        synchronized(lock2) { System.out.println("Thread 1 won"); }
    }
});

Thread t2 = new Thread(() -> {
    synchronized(lock2) {
        try { Thread.sleep(100); } catch (Exception e) {}
        synchronized(lock1) { System.out.println("Thread 2 won"); }
    }
});
// t1 waits for lock2, t2 waits for lock1 -> Forever.
```

---

### 3. Thread-Safe Singleton (Double-Checked Locking)
"I use the `volatile` keyword to ensure visibility.
I check if the instance is null *twice*. Once without locking (for performance), and once inside the synchronized block (for safety). This ensures we only sync the very first time."

```java
public class Singleton {
    private static volatile Singleton instance;
    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

---

### 4. Print Even and Odd numbers using 2 Threads
"I need a shared lock and a shared counter.
Each thread checks the parity of the counter. If it's my turn, print and notify. If not, wait."

```java
class Printer {
    int count = 1;
    int limit = 20;

    public synchronized void printOdd() {
        while (count <= limit) {
            if (count % 2 == 0) { try { wait(); } catch(Exception e){} }
            System.out.println("Odd: " + count++);
            notify();
        }
    }

    public synchronized void printEven() {
        while (count <= limit) {
            if (count % 2 != 0) { try { wait(); } catch(Exception e){} }
            System.out.println("Even: " + count++);
            notify();
        }
    }
}
```

---

### 5. Using CompletableFuture (Modern Async)
"Instead of manually creating threads, I use `CompletableFuture`. It allows chaining tasks.
Here, I fetch user data asynchronously, then process it when it arrives, without blocking the main thread."

```java
CompletableFuture.supplyAsync(() -> fetchUserFromDB(1))
    .thenApply(user -> user.getEmail())
    .thenAccept(email -> sendWelcomeEmail(email))
    .join(); // or .get()
```
