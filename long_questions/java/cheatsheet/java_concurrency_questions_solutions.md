# Java Concurrency Questions & Solutions

This guide provides idiomatic solutions to common Java concurrency interview questions.

---

## 1. Producer-Consumer using BlockingQueue
**Question:** One thread produces numbers, multiple worker threads process them, and the main thread collects results.

```java
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class ProducerConsumer {
    public static void main(String[] args) throws InterruptedException {
        BlockingQueue<Integer> queue = new LinkedBlockingQueue<>(5);

        // Producer
        Thread producer = new Thread(() -> {
            try {
                for (int i = 1; i <= 10; i++) {
                    System.out.println("Produced: " + i);
                    queue.put(i);
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });

        // Consumers
        Runnable consumerTask = () -> {
            try {
                while (true) {
                    int item = queue.take(); // Blocks if empty
                    System.out.println(Thread.currentThread().getName() + " Consumed: " + item);
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        };

        Thread c1 = new Thread(consumerTask, "Consumer-1");
        Thread c2 = new Thread(consumerTask, "Consumer-2");

        producer.start();
        c1.start();
        c2.start();

        producer.join();
        Thread.sleep(100); // Give consumers time to finish queue
        c1.interrupt();
        c2.interrupt();
    }
}
```

## 2. Fixed-Size Worker Pool
**Question:** Implement a fixed-size worker pool to process N jobs concurrently.

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class WorkerPool {
    public static void main(String[] args) throws InterruptedException {
        int numWorkers = 3;
        int numJobs = 10;
        
        // Create pool of 3 threads
        ExecutorService executor = Executors.newFixedThreadPool(numWorkers);

        // Submit 10 jobs
        for (int i = 1; i <= numJobs; i++) {
            final int jobId = i;
            executor.submit(() -> {
                System.out.println(Thread.currentThread().getName() + " processing job: " + jobId);
                try {
                    Thread.sleep(500); // Simulate work
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            });
        }

        // Wait for workers to finish
        executor.shutdown();
        executor.awaitTermination(1, TimeUnit.MINUTES);
        System.out.println("All jobs finished.");
    }
}
```

## 3. Limiting Concurrency (Semaphore)
**Question:** Limit the number of concurrent threads accessing a shared resource.

## The Core Logic
The magic happens with the Semaphore:
1. **Capacity:** `Semaphore semaphore = new Semaphore(3);` Creates a room that can only holp 3 permits at a time.
2. **Acquire:** `semaphore.acquire();` Blocks if 3 permits are currently taken.
3. **Release:** `semaphore.release();` Releases the permit, making room for another waiting thread.

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;

public class SemaphoreExample {
    public static void main(String[] args) {
        int maxConcurrent = 3;
        Semaphore semaphore = new Semaphore(maxConcurrent);
        ExecutorService executor = Executors.newFixedThreadPool(10);

        for (int i = 1; i <= 10; i++) {
            final int taskId = i;
            executor.submit(() -> {
                try {
                    semaphore.acquire(); // Acquire permit
                    System.out.println("Task " + taskId + " accessed resource.");
                    Thread.sleep(500);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                } finally {
                    semaphore.release(); // Release permit
                }
            });
        }

        executor.shutdown();
    }
}
```

## 4. Sequencing Threads (CountDownLatch)
**Question:** Ensure a specific order of execution, for example, waiting for 3 services to start before the main app starts.

```java
import java.util.concurrent.CountDownLatch;

public class SequencingExample {
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch latch = new CountDownLatch(3);

        for (int i = 1; i <= 3; i++) {
            final int serviceId = i;
            new Thread(() -> {
                System.out.println("Service " + serviceId + " starting...");
                try {
                    Thread.sleep((long) (Math.random() * 1000));
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
                System.out.println("Service " + serviceId + " started.");
                latch.countDown(); // Decrease wait count
            }).start();
        }

        latch.await(); // Main thread waits until count drops to 0
        System.out.println("All services started. Main app running.");
    }
}
```

## 5. Graceful Shutdown
**Question:** Gracefully shut down an ExecutorService when a timeout occurs or application closes.

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class GracefulShutdown {
    public static void main(String[] args) {
        ExecutorService executor = Executors.newFixedThreadPool(2);

        executor.submit(() -> {
            try {
                Thread.sleep(5000);
                System.out.println("Task finished.");
            } catch (InterruptedException e) {
                System.out.println("Task interrupted.");
            }
        });

        shutdownAndAwaitTermination(executor);
    }

    static void shutdownAndAwaitTermination(ExecutorService pool) {
        pool.shutdown(); // Disable new tasks from being submitted
        try {
            // Wait a while for existing tasks to terminate
            if (!pool.awaitTermination(2, TimeUnit.SECONDS)) {
                pool.shutdownNow(); // Cancel currently executing tasks
                // Wait a while for tasks to respond to being cancelled
                if (!pool.awaitTermination(2, TimeUnit.SECONDS))
                    System.err.println("Pool did not terminate");
            }
        } catch (InterruptedException ie) {
            // (Re-)Cancel if current thread also interrupted
            pool.shutdownNow();
            // Preserve interrupt status
            Thread.currentThread().interrupt();
        }
    }
}
```

## 6. Stop on First Error (CompletableFuture.anyOf)
**Question:** Stop all tasks when any worker returns an error or a result first.

```java
import java.util.concurrent.CompletableFuture;

public class StopOnFirstError {
    public static void main(String[] args) {
        CompletableFuture<String> task1 = CompletableFuture.supplyAsync(() -> {
            try { Thread.sleep(1000); } catch (InterruptedException e) {}
            return "Task 1 Done";
        });

        CompletableFuture<String> task2 = CompletableFuture.supplyAsync(() -> {
            try { Thread.sleep(100); } catch (InterruptedException e) {}
            throw new RuntimeException("Task 2 Failed");
        });

        CompletableFuture<String> task3 = CompletableFuture.supplyAsync(() -> {
            try { Thread.sleep(500); } catch (InterruptedException e) {}
            return "Task 3 Done";
        });

        // Continue as soon as one starts or fails
        CompletableFuture<Object> firstCompleted = CompletableFuture.anyOf(task1, task2, task3);

        firstCompleted.whenComplete((result, exception) -> {
            if (exception != null) {
                System.out.println("Stopped due to error: " + exception.getMessage());
            } else {
                System.out.println("First result: " + result);
            }
        });

        firstCompleted.join();
    }
}
```

## 7. Fix Deadlock
**Question:** Fix a program that deadlocks due to incorrect locking sequence.

```java
public class DeadlockFix {
    private static final Object lock1 = new Object();
    private static final Object lock2 = new Object();

    public static void main(String[] args) {
        // BROKEN (Deadlock due to different lock order):
        // Thread 1 locks lock1 then lock2
        // Thread 2 locks lock2 then lock1

        // FIXED (Always acquire locks in the same order):
        Thread t1 = new Thread(() -> {
            synchronized (lock1) {
                System.out.println("Thread 1: Holding lock 1...");
                try { Thread.sleep(10); } catch (InterruptedException e) {}
                synchronized (lock2) {
                    System.out.println("Thread 1: Holding lock 1 & 2...");
                }
            }
        });

        Thread t2 = new Thread(() -> {
            synchronized (lock1) { // Changed order from lock2->lock1 to lock1->lock2
                System.out.println("Thread 2: Holding lock 1...");
                try { Thread.sleep(10); } catch (InterruptedException e) {}
                synchronized (lock2) {
                    System.out.println("Thread 2: Holding lock 1 & 2...");
                }
            }
        });

        t1.start();
        t2.start();
    }
}
```

## 8. Thread-Safe Counter
**Question:** Implement a thread-safe counter efficiently.

```java
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class ThreadSafeCounter {
    // Using AtomicInteger for lock-free thread safety
    private AtomicInteger count = new AtomicInteger(0);

    public void increment() {
        count.incrementAndGet();
    }

    public int getValue() {
        return count.get();
    }

    public static void main(String[] args) throws InterruptedException {
        ThreadSafeCounter counter = new ThreadSafeCounter();
        ExecutorService executor = Executors.newFixedThreadPool(10);

        for (int i = 0; i < 1000; i++) {
            executor.submit(counter::increment);
        }

        executor.shutdown();
        executor.awaitTermination(10, TimeUnit.SECONDS);
        System.out.println("Final Count: " + counter.getValue());
    }
}
```

## 9. ReadWriteLock (Many Readers)
**Question:** Handle a many-readers, few-writers scenario safely.

```java
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class SafeMap {
    private final Map<String, Integer> data = new HashMap<>();
    private final ReadWriteLock rwLock = new ReentrantReadWriteLock();

    public int read(String key) {
        rwLock.readLock().lock(); // Multiple readers allowed
        try {
            return data.getOrDefault(key, 0);
        } finally {
            rwLock.readLock().unlock();
        }
    }

    public void write(String key, int val) {
        rwLock.writeLock().lock(); // Exclusive lock
        try {
            data.put(key, val);
        } finally {
            rwLock.writeLock().unlock();
        }
    }

    public static void main(String[] args) {
        SafeMap map = new SafeMap();
        map.write("foo", 1);

        for (int i = 0; i < 5; i++) {
            new Thread(() -> System.out.println(map.read("foo"))).start();
        }
    }
}
```

## 10. Ping Pong (Alternating Turns)
**Question:** Coordinate two threads to alternately print “ping” and “pong”.

```java
public class PingPong {
    private boolean isPingTurn = true;

    public synchronized void printPing() {
        while (!isPingTurn) {
            try { wait(); } catch (InterruptedException e) {}
        }
        System.out.println("ping");
        isPingTurn = false;
        notify();
    }

    public synchronized void printPong() {
        while (isPingTurn) {
            try { wait(); } catch (InterruptedException e) {}
        }
        System.out.println("pong");
        isPingTurn = true;
        notify();
    }

    public static void main(String[] args) {
        PingPong game = new PingPong();

        Thread pingThread = new Thread(() -> {
            for (int i = 0; i < 5; i++) game.printPing();
        });

        Thread pongThread = new Thread(() -> {
            for (int i = 0; i < 5; i++) game.printPong();
        });

        pingThread.start();
        pongThread.start();
    }
}
```

## 11. Pipeline Pattern (CompletableFuture)
**Question:** Build a pipeline where each stage runs asynchronously.

```java
import java.util.concurrent.CompletableFuture;

public class PipelineExample {
    public static void main(String[] args) {
        CompletableFuture.supplyAsync(() -> {
            System.out.println("Stage 1: Generating data in " + Thread.currentThread().getName());
            return 5;
        }).thenApplyAsync(num -> {
            System.out.println("Stage 2: Squaring in " + Thread.currentThread().getName());
            return num * num;
        }).thenAcceptAsync(result -> {
            System.out.println("Stage 3: Printing result in " + Thread.currentThread().getName() + " -> " + result);
        }).join(); // Wait for completion
    }
}
```

## 12. Context Cancellation Timeout (CompletableFuture CompleteOnTimeout)
**Question:** Cancel a task / provide fallback if it exceeds a timeout.

```java
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;

public class TimeoutFallback {
    public static void main(String[] args) {
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
            try {
                Thread.sleep(2000); // Takes 2 seconds
            } catch (InterruptedException e) {}
            return "Success Result";
        });

        // Java 9+: completeOnTimeout
        String result = future.completeOnTimeout("Timeout Fallback Result", 1, TimeUnit.SECONDS).join();
        System.out.println("Result: " + result);

        // Alternatively if throwing exception:
        // future.orTimeout(1, TimeUnit.SECONDS);
    }
}
```

## 13. Print 1 to N sequentially using 3 Threads
**Question:** Print numbers from 1 to 10 in order using multiple threads.

```java
public class PrintInOrder {
    private int counter = 1;
    private final int limit = 10;
    private final int numThreads = 3;

    public synchronized void print(int threadId) {
        while (counter <= limit) {
            if (counter % numThreads != threadId) {
                try {
                    wait();
                } catch (InterruptedException e) {}
            } else {
                if (counter <= limit) {
                    System.out.println("Thread " + threadId + " printed: " + counter);
                    counter++;
                    notifyAll();
                }
            }
        }
    }

    public static void main(String[] args) {
        PrintInOrder printer = new PrintInOrder();

        Thread t1 = new Thread(() -> printer.print(1));
        Thread t2 = new Thread(() -> printer.print(2));
        Thread t3 = new Thread(() -> printer.print(0)); // 3 % 3 == 0

        t1.start();
        t2.start();
        t3.start();
    }
}
```

## 14. Synchronizing Threads with CyclicBarrier
**Question:** Make multiple threads wait for each other to reach a common barrier point.

```java
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CyclicBarrier;

public class BarrierExample {
    public static void main(String[] args) {
        int numWorkers = 3;
        CyclicBarrier barrier = new CyclicBarrier(numWorkers, () -> {
            System.out.println("All parties arrived at the barrier. Let's proceed!");
        });

        for (int i = 1; i <= numWorkers; i++) {
            new Thread(() -> {
                try {
                    System.out.println(Thread.currentThread().getName() + " is working...");
                    Thread.sleep((long) (Math.random() * 1000));
                    System.out.println(Thread.currentThread().getName() + " reached barrier.");
                    barrier.await();
                    System.out.println(Thread.currentThread().getName() + " continued execution.");
                } catch (InterruptedException | BrokenBarrierException e) {
                    e.printStackTrace();
                }
            }, "Worker-" + i).start();
        }
    }
}
```

## 15. ThreadLocal for Per-Thread Data
**Question:** Create a variable that has a distinct value for each thread.

```java
public class ThreadLocalExample {
    // Each thread accessing this ThreadLocal will have its own independent copy
    private static final ThreadLocal<Integer> threadId = ThreadLocal.withInitial(() -> 0);

    public static void main(String[] args) {
        Runnable task = () -> {
            int id = (int) (Math.random() * 100);
            threadId.set(id);
            try { Thread.sleep(100); } catch (InterruptedException e) {}
            System.out.println(Thread.currentThread().getName() + " ID: " + threadId.get());
        };

        new Thread(task, "Thread-1").start();
        new Thread(task, "Thread-2").start();
    }
}
```

## 16. Custom Blocking Queue using wait/notify
**Question:** Implement your own bounded BlockingQueue without using `java.util.concurrent`.

```java
import java.util.LinkedList;
import java.util.Queue;

public class CustomBlockingQueue<T> {
    private final Queue<T> queue = new LinkedList<>();
    private final int capacity;

    public CustomBlockingQueue(int capacity) {
        this.capacity = capacity;
    }

    public synchronized void put(T item) throws InterruptedException {
        while (queue.size() == capacity) {
            wait();
        }
        queue.add(item);
        notifyAll();
    }

    public synchronized T take() throws InterruptedException {
        while (queue.isEmpty()) {
            wait();
        }
        T item = queue.poll();
        notifyAll();
        return item;
    }

    public static void main(String[] args) throws InterruptedException {
        CustomBlockingQueue<Integer> q = new CustomBlockingQueue<>(2);
        
        new Thread(() -> {
            try {
                q.put(1); q.put(2);
                System.out.println("Queue full. Waiting to put 3...");
                q.put(3);
                System.out.println("Put 3");
            } catch (InterruptedException e) {}
        }).start();

        Thread.sleep(1000);
        System.out.println("Taking: " + q.take());
    }
}
```

## 17. ForkJoinPool Data Parallelism
**Question:** Split a large computational task into smaller tasks using ForkJoinPool.

```java
import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.RecursiveTask;

public class ForkJoinSum extends RecursiveTask<Long> {
    private final long[] array;
    private final int start, end;
    private static final int THRESHOLD = 1000;

    public ForkJoinSum(long[] array, int start, int end) {
        this.array = array;
        this.start = start;
        this.end = end;
    }

    @Override
    protected Long compute() {
        if (end - start <= THRESHOLD) {
            long sum = 0;
            for (int i = start; i < end; i++) {
                sum += array[i];
            }
            return sum;
        } else {
            int mid = (start + end) / 2;
            ForkJoinSum leftTask = new ForkJoinSum(array, start, mid);
            ForkJoinSum rightTask = new ForkJoinSum(array, mid, end);
            
            leftTask.fork(); // Async execute left
            long rightResult = rightTask.compute(); // Sync execute right
            long leftResult = leftTask.join(); // Wait for left
            
            return leftResult + rightResult;
        }
    }

    public static void main(String[] args) {
        long[] array = new long[10000];
        for (int i = 0; i < array.length; i++) array[i] = i;

        ForkJoinPool pool = new ForkJoinPool();
        long result = pool.invoke(new ForkJoinSum(array, 0, array.length));
        System.out.println("Sum: " + result);
    }
}
```

## 18. Exchanger (Swap data between 2 threads)
**Question:** Two threads wait for each other to arrive at a swap point to exchange data.

```java
import java.util.concurrent.Exchanger;

public class ExchangerExample {
    public static void main(String[] args) {
        Exchanger<String> exchanger = new Exchanger<>();

        new Thread(() -> {
            try {
                String myData = "Data from Thread 1";
                System.out.println("Thread 1 waiting to exchange...");
                String received = exchanger.exchange(myData);
                System.out.println("Thread 1 received: " + received);
            } catch (InterruptedException e) {}
        }).start();

        new Thread(() -> {
            try {
                Thread.sleep(1000); // Simulate work
                String myData = "Data from Thread 2";
                System.out.println("Thread 2 exchanging...");
                String received = exchanger.exchange(myData);
                System.out.println("Thread 2 received: " + received);
            } catch (InterruptedException e) {}
        }).start();
    }
}
```

## 19. ScheduledExecutorService (Periodic Tasks)
**Question:** Implement a rate limiter or periodic job runner.

```java
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

public class RateLimiterExample {
    public static void main(String[] args) {
        // Pool of 1 thread for scheduling
        ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

        Runnable task = () -> System.out.println("Task executed at: " + System.currentTimeMillis());

        // Runs with an initial delay of 0, and repeats every 2 seconds
        scheduler.scheduleAtFixedRate(task, 0, 2, TimeUnit.SECONDS);

        // Stop scheduler after 7 seconds
        scheduler.schedule(() -> {
            System.out.println("Shutting down scheduler...");
            scheduler.shutdown();
        }, 7, TimeUnit.SECONDS);
    }
}
```

## 20. Lock with Condition (Alternative to wait/notify)
**Question:** Refactor a wait/notify block to use ReentrantLock and Condition.

```java
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class ConditionExample {
    private final Lock lock = new ReentrantLock();
    private final Condition condition = lock.newCondition();
    private boolean ready = false;

    public void waitForReady() throws InterruptedException {
        lock.lock();
        try {
            while (!ready) {
                condition.await(); // wait equivalent
            }
            System.out.println("Proceeding because it's ready!");
        } finally {
            lock.unlock();
        }
    }

    public void makeReady() {
        lock.lock();
        try {
            ready = true;
            condition.signalAll(); // notifyAll equivalent
        } finally {
            lock.unlock();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        ConditionExample example = new ConditionExample();

        new Thread(() -> {
            try { example.waitForReady(); } catch (InterruptedException e) {}
        }).start();

        Thread.sleep(1000);
        example.makeReady();
    }
}
```
