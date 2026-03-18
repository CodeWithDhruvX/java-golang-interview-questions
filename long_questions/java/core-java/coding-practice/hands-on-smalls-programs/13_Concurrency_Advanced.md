# Advanced Concurrency: Practical Programs

**Goal**: Master advanced concurrency concepts including synchronized blocks, locks, concurrent collections, and producer-consumer patterns.

## 1. Synchronized Methods and Blocks

### Thread Safety with Synchronization

```java
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.concurrent.locks.*;

// Bank account with synchronization
class BankAccount {
    private double balance;
    private final Object lock = new Object();
    private final AtomicInteger transactionCount = new AtomicInteger(0);
    
    public BankAccount(double initialBalance) {
        this.balance = initialBalance;
    }
    
    // Synchronized method
    public synchronized void deposit(double amount) {
        if (amount <= 0) {
            throw new IllegalArgumentException("Amount must be positive");
        }
        
        double oldBalance = balance;
        balance += amount;
        transactionCount.incrementAndGet();
        
        System.out.printf("%s deposited $%.2f. Balance: $%.2f -> $%.2f%n",
                         Thread.currentThread().getName(), amount, oldBalance, balance);
    }
    
    // Synchronized method
    public synchronized void withdraw(double amount) {
        if (amount <= 0) {
            throw new IllegalArgumentException("Amount must be positive");
        }
        
        if (balance < amount) {
            throw new IllegalArgumentException("Insufficient funds");
        }
        
        double oldBalance = balance;
        balance -= amount;
        transactionCount.incrementAndGet();
        
        System.out.printf("%s withdrew $%.2f. Balance: $%.2f -> $%.2f%n",
                         Thread.currentThread().getName(), amount, oldBalance, balance);
    }
    
    // Synchronized block
    public void transfer(BankAccount targetAccount, double amount) {
        synchronized (this) {
            synchronized (targetAccount) {
                // Avoid deadlock by always acquiring locks in the same order
                if (this.hashCode() > targetAccount.hashCode()) {
                    synchronized (targetAccount) {
                        synchronized (this) {
                            performTransfer(targetAccount, amount);
                        }
                    }
                } else {
                    performTransfer(targetAccount, amount);
                }
            }
        }
    }
    
    private void performTransfer(BankAccount targetAccount, double amount) {
        if (amount <= 0) {
            throw new IllegalArgumentException("Amount must be positive");
        }
        
        if (balance < amount) {
            throw new IllegalArgumentException("Insufficient funds for transfer");
        }
        
        this.balance -= amount;
        targetAccount.balance += amount;
        transactionCount.incrementAndGet();
        
        System.out.printf("%s transferred $%.2f to account %s. New balance: $%.2f%n",
                         Thread.currentThread().getName(), amount, 
                         targetAccount.hashCode(), this.balance);
    }
    
    // Synchronized getter
    public synchronized double getBalance() {
        return balance;
    }
    
    public int getTransactionCount() {
        return transactionCount.get();
    }
    
    // Demonstrate synchronized block for complex operations
    public synchronized void complexOperation() {
        synchronized (lock) {
            System.out.println(Thread.currentThread().getName() + " performing complex operation");
            
            // Simulate complex calculation
            double interest = balance * 0.05;
            balance += interest;
            
            System.out.printf("%s calculated interest: $%.2f. New balance: $%.2f%n",
                             Thread.currentThread().getName(), interest, balance);
        }
    }
}

// Task for testing bank account operations
class BankTask implements Callable<Integer> {
    private final BankAccount account;
    private final String operation;
    private final double amount;
    private final int iterations;
    
    public BankTask(BankAccount account, String operation, double amount, int iterations) {
        this.account = account;
        this.operation = operation;
        this.amount = amount;
        this.iterations = iterations;
    }
    
    @Override
    public Integer call() {
        for (int i = 0; i < iterations; i++) {
            try {
                switch (operation.toLowerCase()) {
                    case "deposit":
                        account.deposit(amount);
                        break;
                    case "withdraw":
                        account.withdraw(amount);
                        break;
                    case "complex":
                        account.complexOperation();
                        break;
                }
                
                // Add some delay to increase chance of race conditions
                Thread.sleep(1);
                
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
        return iterations;
    }
}

public class SynchronizationDemo {
    public static void main(String[] args) throws InterruptedException {
        System.out.println("=== Synchronized Methods and Blocks Demo ===");
        
        // Test basic synchronization
        testBasicSynchronization();
        
        // Test synchronized transfer
        testSynchronizedTransfer();
        
        // Test deadlock prevention
        testDeadlockPrevention();
        
        // Test synchronized blocks
        testSynchronizedBlocks();
    }
    
    private static void testBasicSynchronization() throws InterruptedException {
        System.out.println("\n--- Basic Synchronization Test ---");
        
        BankAccount account = new BankAccount(1000.0);
        ExecutorService executor = Executors.newFixedThreadPool(10);
        
        // Create concurrent tasks
        List<Future<Integer>> futures = new ArrayList<>();
        
        for (int i = 0; i < 5; i++) {
            futures.add(executor.submit(new BankTask(account, "deposit", 100.0, 10)));
        }
        
        for (int i = 0; i < 3; i++) {
            futures.add(executor.submit(new BankTask(account, "withdraw", 50.0, 10)));
        }
        
        // Wait for all tasks to complete
        int totalOperations = 0;
        for (Future<Integer> future : futures) {
            try {
                totalOperations += future.get();
            } catch (ExecutionException e) {
                System.err.println("Task failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.printf("Final balance: $%.2f%n", account.getBalance());
        System.out.printf("Expected balance: $%.2f%n", 1000.0 + (5 * 10 * 100.0) - (3 * 10 * 50.0));
        System.out.printf("Transaction count: %d%n", account.getTransactionCount());
        System.out.printf("Total operations: %d%n", totalOperations);
    }
    
    private static void testSynchronizedTransfer() throws InterruptedException {
        System.out.println("\n--- Synchronized Transfer Test ---");
        
        BankAccount account1 = new BankAccount(1000.0);
        BankAccount account2 = new BankAccount(500.0);
        
        ExecutorService executor = Executors.newFixedThreadPool(5);
        
        // Create transfer tasks
        List<Future<Integer>> futures = new ArrayList<>();
        
        for (int i = 0; i < 10; i++) {
            final int taskId = i;
            Future<Integer> future = executor.submit(() -> {
                if (taskId % 2 == 0) {
                    account1.transfer(account2, 50.0);
                } else {
                    account2.transfer(account1, 30.0);
                }
                return 1;
            });
            futures.add(future);
        }
        
        // Wait for completion
        for (Future<Integer> future : futures) {
            try {
                future.get();
            } catch (ExecutionException e) {
                System.err.println("Transfer failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.printf("Account 1 balance: $%.2f%n", account1.getBalance());
        System.out.printf("Account 2 balance: $%.2f%n", account2.getBalance());
        System.out.printf("Total balance: $%.2f%n", account1.getBalance() + account2.getBalance());
    }
    
    private static void testDeadlockPrevention() throws InterruptedException {
        System.out.println("\n--- Deadlock Prevention Test ---");
        
        BankAccount account1 = new BankAccount(1000.0);
        BankAccount account2 = new BankAccount(1000.0);
        
        ExecutorService executor = Executors.newFixedThreadPool(2);
        
        // Create tasks that could potentially deadlock
        Future<?> future1 = executor.submit(() -> {
            for (int i = 0; i < 100; i++) {
                account1.transfer(account2, 10.0);
            }
        });
        
        Future<?> future2 = executor.submit(() -> {
            for (int i = 0; i < 100; i++) {
                account2.transfer(account1, 5.0);
            }
        });
        
        try {
            future1.get();
            future2.get();
            System.out.println("No deadlock occurred!");
        } catch (ExecutionException e) {
            System.err.println("Error during transfer: " + e.getMessage());
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.printf("Account 1 balance: $%.2f%n", account1.getBalance());
        System.out.printf("Account 2 balance: $%.2f%n", account2.getBalance());
    }
    
    private static void testSynchronizedBlocks() throws InterruptedException {
        System.out.println("\n--- Synchronized Blocks Test ---");
        
        BankAccount account = new BankAccount(1000.0);
        ExecutorService executor = Executors.newFixedThreadPool(5);
        
        List<Future<Integer>> futures = new ArrayList<>();
        
        for (int i = 0; i < 5; i++) {
            futures.add(executor.submit(new BankTask(account, "complex", 0, 20)));
        }
        
        // Wait for completion
        for (Future<Integer> future : futures) {
            try {
                future.get();
            } catch (ExecutionException e) {
                System.err.println("Complex operation failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.printf("Final balance after complex operations: $%.2f%n", account.getBalance());
    }
}
```

## 2. Advanced Lock Mechanisms

### ReentrantLock and ReadWriteLock

```java
import java.util.concurrent.*;
import java.util.concurrent.locks.*;
import java.util.concurrent.atomic.*;

// Resource manager with different lock types
class ResourceManager {
    private final ReentrantLock reentrantLock = new ReentrantLock();
    private final ReentrantReadWriteLock readWriteLock = new ReentrantReadWriteLock();
    private final Lock readLock = readWriteLock.readLock();
    private final Lock writeLock = readWriteLock.writeLock();
    
    private final Map<String, String> data = new ConcurrentHashMap<>();
    private final AtomicInteger readCount = new AtomicInteger(0);
    private final AtomicInteger writeCount = new AtomicInteger(0);
    
    // Using ReentrantLock
    public void performReentrantLockOperation(String operation) {
        reentrantLock.lock();
        try {
            System.out.printf("%s acquired reentrant lock for operation: %s%n",
                             Thread.currentThread().getName(), operation);
            
            // Simulate work
            Thread.sleep(100);
            
            // Check if lock is held by current thread
            if (reentrantLock.isHeldByCurrentThread()) {
                System.out.printf("%s is holding the lock (queue length: %d)%n",
                                 Thread.currentThread().getName(), 
                                 reentrantLock.getQueueLength());
            }
            
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        } finally {
            reentrantLock.unlock();
            System.out.printf("%s released reentrant lock%n", Thread.currentThread().getName());
        }
    }
    
    // Try lock with timeout
    public boolean tryLockWithTimeout(String operation, long timeout, TimeUnit unit) {
        try {
            if (reentrantLock.tryLock(timeout, unit)) {
                try {
                    System.out.printf("%s acquired lock with timeout for: %s%n",
                                     Thread.currentThread().getName(), operation);
                    Thread.sleep(50);
                    return true;
                } finally {
                    reentrantLock.unlock();
                }
            } else {
                System.out.printf("%s failed to acquire lock within timeout for: %s%n",
                                 Thread.currentThread().getName(), operation);
                return false;
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            return false;
        }
    }
    
    // Read operation with ReadWriteLock
    public String readData(String key) {
        readLock.lock();
        try {
            readCount.incrementAndGet();
            System.out.printf("%s reading data for key: %s (active readers: %d)%n",
                             Thread.currentThread().getName(), key, readCount.get());
            
            // Simulate read operation
            Thread.sleep(50);
            
            return data.get(key);
            
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            return null;
        } finally {
            readLock.unlock();
            readCount.decrementAndGet();
        }
    }
    
    // Write operation with ReadWriteLock
    public void writeData(String key, String value) {
        writeLock.lock();
        try {
            writeCount.incrementAndGet();
            System.out.printf("%s writing data for key: %s (active writers: %d)%n",
                             Thread.currentThread().getName(), key, writeCount.get());
            
            // Simulate write operation
            Thread.sleep(100);
            
            data.put(key, value);
            
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        } finally {
            writeLock.unlock();
            writeCount.decrementAndGet();
        }
    }
    
    // Demonstrate lock fairness
    public void fairLockDemo(String operation) {
        ReentrantLock fairLock = new ReentrantLock(true); // Fair lock
        fairLock.lock();
        try {
            System.out.printf("%s acquired fair lock for: %s%n",
                             Thread.currentThread().getName(), operation);
            Thread.sleep(50);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        } finally {
            fairLock.unlock();
        }
    }
    
    // Get statistics
    public void printStatistics() {
        System.out.printf("Data size: %d, Current readers: %d, Current writers: %d%n",
                         data.size(), readCount.get(), writeCount.get());
    }
}

// Task for testing different lock types
class LockTestTask implements Callable<String> {
    private final ResourceManager manager;
    private final String lockType;
    private final String operation;
    
    public LockTestTask(ResourceManager manager, String lockType, String operation) {
        this.manager = manager;
        this.lockType = lockType;
        this.operation = operation;
    }
    
    @Override
    public String call() {
        try {
            switch (lockType.toLowerCase()) {
                case "reentrant":
                    manager.performReentrantLockOperation(operation);
                    break;
                case "trylock":
                    manager.tryLockWithTimeout(operation, 200, TimeUnit.MILLISECONDS);
                    break;
                case "fair":
                    manager.fairLockDemo(operation);
                    break;
                case "read":
                    manager.readData(operation);
                    break;
                case "write":
                    manager.writeData(operation, "Value-" + Thread.currentThread().getId());
                    break;
            }
            return "Completed: " + operation;
        } catch (Exception e) {
            return "Failed: " + operation + " - " + e.getMessage();
        }
    }
}

public class AdvancedLocksDemo {
    public static void main(String[] args) throws InterruptedException {
        System.out.println("=== Advanced Lock Mechanisms Demo ===");
        
        ResourceManager manager = new ResourceManager();
        
        // Test ReentrantLock
        testReentrantLock(manager);
        
        // Test try lock with timeout
        testTryLockWithTimeout(manager);
        
        // Test ReadWriteLock
        testReadWriteLock(manager);
        
        // Test fair locks
        testFairLocks(manager);
        
        // Test lock contention
        testLockContention(manager);
    }
    
    private static void testReentrantLock(ResourceManager manager) throws InterruptedException {
        System.out.println("\n--- ReentrantLock Test ---");
        
        ExecutorService executor = Executors.newFixedThreadPool(3);
        List<Future<String>> futures = new ArrayList<>();
        
        for (int i = 0; i < 5; i++) {
            futures.add(executor.submit(new LockTestTask(manager, "reentrant", "Operation-" + i)));
        }
        
        // Wait for completion
        for (Future<String> future : futures) {
            try {
                System.out.println(future.get());
            } catch (ExecutionException e) {
                System.err.println("ReentrantLock task failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
    }
    
    private static void testTryLockWithTimeout(ResourceManager manager) throws InterruptedException {
        System.out.println("\n--- TryLock with Timeout Test ---");
        
        ExecutorService executor = Executors.newFixedThreadPool(5);
        List<Future<String>> futures = new ArrayList<>();
        
        for (int i = 0; i < 10; i++) {
            futures.add(executor.submit(new LockTestTask(manager, "trylock", "TimeoutOp-" + i)));
        }
        
        // Wait for completion
        for (Future<String> future : futures) {
            try {
                System.out.println(future.get());
            } catch (ExecutionException e) {
                System.err.println("TryLock task failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
    }
    
    private static void testReadWriteLock(ResourceManager manager) throws InterruptedException {
        System.out.println("\n--- ReadWriteLock Test ---");
        
        ExecutorService executor = Executors.newFixedThreadPool(10);
        List<Future<String>> futures = new ArrayList<>();
        
        // Add some initial data
        manager.writeData("key1", "value1");
        manager.writeData("key2", "value2");
        manager.writeData("key3", "value3");
        
        // Create read tasks (more concurrent reads)
        for (int i = 0; i < 8; i++) {
            futures.add(executor.submit(new LockTestTask(manager, "read", "key" + (i % 3 + 1))));
        }
        
        // Create write tasks (exclusive)
        for (int i = 0; i < 2; i++) {
            futures.add(executor.submit(new LockTestTask(manager, "write", "newkey" + i)));
        }
        
        // Wait for completion
        for (Future<String> future : futures) {
            try {
                System.out.println(future.get());
            } catch (ExecutionException e) {
                System.err.println("ReadWriteLock task failed: " + e.getMessage());
            }
        }
        
        manager.printStatistics();
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
    }
    
    private static void testFairLocks(ResourceManager manager) throws InterruptedException {
        System.out.println("\n--- Fair Locks Test ---");
        
        ExecutorService executor = Executors.newFixedThreadPool(3);
        List<Future<String>> futures = new ArrayList<>();
        
        for (int i = 0; i < 5; i++) {
            futures.add(executor.submit(new LockTestTask(manager, "fair", "FairOp-" + i)));
        }
        
        // Wait for completion
        for (Future<String> future : futures) {
            try {
                System.out.println(future.get());
            } catch (ExecutionException e) {
                System.err.println("FairLock task failed: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
    }
    
    private static void testLockContention(ResourceManager manager) throws InterruptedException {
        System.out.println("\n--- Lock Contention Test ---");
        
        ExecutorService executor = Executors.newFixedThreadPool(20);
        List<Future<String>> futures = new ArrayList<>();
        
        // Create high contention scenario
        for (int i = 0; i < 50; i++) {
            String lockType = i % 3 == 0 ? "reentrant" : (i % 3 == 1 ? "read" : "write");
            futures.add(executor.submit(new LockTestTask(manager, lockType, "Contention-" + i)));
        }
        
        // Wait for completion
        int completed = 0;
        for (Future<String> future : futures) {
            try {
                future.get();
                completed++;
            } catch (ExecutionException e) {
                System.err.println("Contention task failed: " + e.getMessage());
            }
        }
        
        System.out.printf("Completed %d out of %d tasks%n", completed, futures.size());
        manager.printStatistics();
        
        executor.shutdown();
        executor.awaitTermination(10, TimeUnit.SECONDS);
    }
}
```

## 3. Concurrent Collections

### Thread-Safe Collections in Action

```java
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.*;
import java.util.stream.*;

// Task manager using concurrent collections
class TaskManager {
    private final ConcurrentHashMap<String, Task> activeTasks = new ConcurrentHashMap<>();
    private final ConcurrentLinkedQueue<Task> taskQueue = new ConcurrentLinkedQueue<>();
    private final CopyOnWriteArrayList<TaskListener> listeners = new CopyOnWriteArrayList<>();
    private final AtomicInteger taskCounter = new AtomicInteger(0);
    
    // Task class
    static class Task {
        private final String id;
        private final String description;
        private final AtomicInteger status;
        private final long createdTime;
        
        public Task(String id, String description) {
            this.id = id;
            this.description = description;
            this.status = new AtomicInteger(0); // 0=Pending, 1=Running, 2=Completed, 3=Failed
            this.createdTime = System.currentTimeMillis();
        }
        
        public String getId() { return id; }
        public String getDescription() { return description; }
        public int getStatus() { return status.get(); }
        public void setStatus(int status) { this.status.set(status); }
        public long getCreatedTime() { return createdTime; }
        
        @Override
        public String toString() {
            return String.format("Task{id='%s', desc='%s', status=%d}", id, description, status.get());
        }
    }
    
    // Task listener interface
    interface TaskListener {
        void onTaskCreated(Task task);
        void onTaskStarted(Task task);
        void onTaskCompleted(Task task);
        void onTaskFailed(Task task);
    }
    
    // Add task to queue
    public String addTask(String description) {
        String taskId = "TASK-" + taskCounter.incrementAndGet();
        Task task = new Task(taskId, description);
        
        activeTasks.put(taskId, task);
        taskQueue.offer(task);
        
        // Notify listeners
        listeners.forEach(listener -> listener.onTaskCreated(task));
        
        System.out.printf("%s added task: %s%n", Thread.currentThread().getName(), task);
        return taskId;
    }
    
    // Process next task
    public Task processNextTask() {
        Task task = taskQueue.poll();
        if (task != null) {
            task.setStatus(1); // Running
            listeners.forEach(listener -> listener.onTaskStarted(task));
            
            // Simulate task processing
            try {
                Thread.sleep(100);
                
                // Random success/failure
                if (Math.random() > 0.2) { // 80% success rate
                    task.setStatus(2); // Completed
                    listeners.forEach(listener -> listener.onTaskCompleted(task));
                } else {
                    task.setStatus(3); // Failed
                    listeners.forEach(listener -> listener.onTaskFailed(task));
                }
                
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                task.setStatus(3); // Failed
                listeners.forEach(listener -> listener.onTaskFailed(task));
            }
        }
        return task;
    }
    
    // Get task by ID
    public Task getTask(String taskId) {
        return activeTasks.get(taskId);
    }
    
    // Get all active tasks
    public Collection<Task> getAllTasks() {
        return new ArrayList<>(activeTasks.values());
    }
    
    // Get tasks by status
    public List<Task> getTasksByStatus(int status) {
        return activeTasks.values().stream()
                .filter(task -> task.getStatus() == status)
                .collect(Collectors.toList());
    }
    
    // Add listener
    public void addListener(TaskListener listener) {
        listeners.add(listener);
    }
    
    // Get statistics
    public Map<Integer, Long> getTaskStatistics() {
        return activeTasks.values().stream()
                .collect(Collectors.groupingBy(Task::getStatus, Collectors.counting()));
    }
    
    // Clear completed tasks
    public int clearCompletedTasks() {
        int removed = 0;
        Iterator<Map.Entry<String, Task>> iterator = activeTasks.entrySet().iterator();
        while (iterator.hasNext()) {
            Map.Entry<String, Task> entry = iterator.next();
            if (entry.getValue().getStatus() == 2) { // Completed
                iterator.remove();
                removed++;
            }
        }
        return removed;
    }
}

// Cache using ConcurrentMap
class ConcurrentCache<K, V> {
    private final ConcurrentHashMap<K, CacheEntry<V>> cache = new ConcurrentHashMap<>();
    private final ScheduledExecutorService cleanupExecutor = Executors.newSingleThreadScheduledExecutor();
    private final long ttlMillis;
    
    private static class CacheEntry<V> {
        private final V value;
        private final long expiryTime;
        
        public CacheEntry(V value, long ttlMillis) {
            this.value = value;
            this.expiryTime = System.currentTimeMillis() + ttlMillis;
        }
        
        public V getValue() { return value; }
        public boolean isExpired() { return System.currentTimeMillis() > expiryTime; }
    }
    
    public ConcurrentCache(long ttlMillis) {
        this.ttlMillis = ttlMillis;
        // Schedule cleanup every minute
        cleanupExecutor.scheduleAtFixedRate(this::cleanup, 1, 1, TimeUnit.MINUTES);
    }
    
    public void put(K key, V value) {
        cache.put(key, new CacheEntry<>(value, ttlMillis));
    }
    
    public V get(K key) {
        CacheEntry<V> entry = cache.get(key);
        if (entry != null && !entry.isExpired()) {
            return entry.getValue();
        }
        return null;
    }
    
    public V getOrDefault(K key, V defaultValue) {
        V value = get(key);
        return value != null ? value : defaultValue;
    }
    
    public void remove(K key) {
        cache.remove(key);
    }
    
    public int size() {
        cleanup(); // Remove expired entries
        return cache.size();
    }
    
    public void cleanup() {
        cache.entrySet().removeIf(entry -> entry.getValue().isExpired());
    }
    
    public void shutdown() {
        cleanupExecutor.shutdown();
    }
}

// Thread-safe statistics collector
class StatisticsCollector {
    private final ConcurrentHashMap<String, LongAdder> counters = new ConcurrentHashMap<>();
    private final ConcurrentHashMap<String, AtomicLong> gauges = new ConcurrentHashMap<>();
    private final ConcurrentLinkedQueue<Event> events = new ConcurrentLinkedQueue<>();
    
    private static class Event {
        private final String type;
        private final long timestamp;
        private final Map<String, Object> data;
        
        public Event(String type, Map<String, Object> data) {
            this.type = type;
            this.timestamp = System.currentTimeMillis();
            this.data = new HashMap<>(data);
        }
        
        public String getType() { return type; }
        public long getTimestamp() { return timestamp; }
        public Map<String, Object> getData() { return data; }
    }
    
    public void increment(String counterName) {
        counters.computeIfAbsent(counterName, k -> new LongAdder()).increment();
    }
    
    public void increment(String counterName, long delta) {
        counters.computeIfAbsent(counterName, k -> new LongAdder()).add(delta);
    }
    
    public long getCounter(String counterName) {
        LongAdder adder = counters.get(counterName);
        return adder != null ? adder.sum() : 0;
    }
    
    public void setGauge(String gaugeName, long value) {
        gauges.computeIfAbsent(gaugeName, k -> new AtomicLong()).set(value);
    }
    
    public long getGauge(String gaugeName) {
        AtomicLong gauge = gauges.get(gaugeName);
        return gauge != null ? gauge.get() : 0;
    }
    
    public void recordEvent(String eventType, Map<String, Object> data) {
        Event event = new Event(eventType, data);
        events.offer(event);
        
        // Keep only last 1000 events
        while (events.size() > 1000) {
            events.poll();
        }
    }
    
    public List<Event> getRecentEvents(int maxEvents) {
        return events.stream()
                .skip(Math.max(0, events.size() - maxEvents))
                .collect(Collectors.toList());
    }
    
    public Map<String, Long> getAllCounters() {
        return counters.entrySet().stream()
                .collect(Collectors.toMap(
                    Map.Entry::getKey,
                    entry -> entry.getValue().sum()
                ));
    }
    
    public Map<String, Long> getAllGauges() {
        return gauges.entrySet().stream()
                .collect(Collectors.toMap(
                    Map.Entry::getKey,
                    entry -> entry.getValue().get()
                ));
    }
}

public class ConcurrentCollectionsDemo {
    public static void main(String[] args) throws InterruptedException {
        System.out.println("=== Concurrent Collections Demo ===");
        
        // Test TaskManager with concurrent collections
        testTaskManager();
        
        // Test ConcurrentCache
        testConcurrentCache();
        
        // Test StatisticsCollector
        testStatisticsCollector();
        
        // Test concurrent collection performance
        testConcurrentPerformance();
    }
    
    private static void testTaskManager() throws InterruptedException {
        System.out.println("\n--- TaskManager Test ---");
        
        TaskManager taskManager = new TaskManager();
        
        // Add task listener
        taskManager.addListener(new TaskManager.TaskListener() {
            @Override
            public void onTaskCreated(TaskManager.Task task) {
                System.out.printf("Listener: Task created - %s%n", task.getId());
            }
            
            @Override
            public void onTaskStarted(TaskManager.Task task) {
                System.out.printf("Listener: Task started - %s%n", task.getId());
            }
            
            @Override
            public void onTaskCompleted(TaskManager.Task task) {
                System.out.printf("Listener: Task completed - %s%n", task.getId());
            }
            
            @Override
            public void onTaskFailed(TaskManager.Task task) {
                System.out.printf("Listener: Task failed - %s%n", task.getId());
            }
        });
        
        // Create producer threads
        ExecutorService producerExecutor = Executors.newFixedThreadPool(3);
        for (int i = 0; i < 3; i++) {
            final int threadId = i;
            producerExecutor.submit(() -> {
                for (int j = 0; j < 5; j++) {
                    taskManager.addTask("Task-" + threadId + "-" + j);
                    try {
                        Thread.sleep(50);
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        // Create consumer threads
        ExecutorService consumerExecutor = Executors.newFixedThreadPool(2);
        for (int i = 0; i < 2; i++) {
            consumerExecutor.submit(() -> {
                while (true) {
                    TaskManager.Task task = taskManager.processNextTask();
                    if (task == null) {
                        break;
                    }
                    try {
                        Thread.sleep(100);
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        // Wait for producers to finish
        producerExecutor.shutdown();
        producerExecutor.awaitTermination(5, TimeUnit.SECONDS);
        
        // Give consumers time to process
        Thread.sleep(2000);
        
        // Shutdown consumers
        consumerExecutor.shutdown();
        consumerExecutor.awaitTermination(5, TimeUnit.SECONDS);
        
        // Display statistics
        System.out.println("\nTask Statistics:");
        Map<Integer, Long> stats = taskManager.getTaskStatistics();
        stats.forEach((status, count) -> {
            String statusName = status == 0 ? "Pending" : 
                               status == 1 ? "Running" : 
                               status == 2 ? "Completed" : "Failed";
            System.out.printf("%s: %d%n", statusName, count);
        });
        
        System.out.printf("Total tasks: %d%n", taskManager.getAllTasks().size());
    }
    
    private static void testConcurrentCache() throws InterruptedException {
        System.out.println("\n--- ConcurrentCache Test ---");
        
        ConcurrentCache<String, String> cache = new ConcurrentCache<>(1000); // 1 second TTL
        
        // Put some values
        cache.put("key1", "value1");
        cache.put("key2", "value2");
        cache.put("key3", "value3");
        
        System.out.println("Cache size after adding: " + cache.size());
        
        // Get values
        System.out.println("key1: " + cache.get("key1"));
        System.out.println("key2: " + cache.get("key2"));
        System.out.println("nonexistent: " + cache.get("nonexistent"));
        
        // Test concurrent access
        ExecutorService executor = Executors.newFixedThreadPool(10);
        
        for (int i = 0; i < 100; i++) {
            final int index = i;
            executor.submit(() -> {
                String key = "key" + (index % 5);
                String value = "value" + index;
                
                if (index % 3 == 0) {
                    cache.put(key, value);
                } else {
                    String retrieved = cache.get(key);
                    if (retrieved != null) {
                        // Do something with retrieved value
                    }
                }
            });
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.println("Cache size after concurrent operations: " + cache.size());
        
        // Wait for some entries to expire
        Thread.sleep(1500);
        cache.cleanup();
        System.out.println("Cache size after cleanup: " + cache.size());
        
        cache.shutdown();
    }
    
    private static void testStatisticsCollector() throws InterruptedException {
        System.out.println("\n--- StatisticsCollector Test ---");
        
        StatisticsCollector stats = new StatisticsCollector();
        
        // Increment counters
        stats.increment("requests");
        stats.increment("requests");
        stats.increment("errors");
        stats.increment("users", 5);
        
        // Set gauges
        stats.setGauge("memory_usage", 1024);
        stats.setGauge("cpu_usage", 75);
        
        // Record events
        Map<String, Object> eventData = new HashMap<>();
        eventData.put("user_id", "user123");
        eventData.put("action", "login");
        stats.recordEvent("user_action", eventData);
        
        eventData = new HashMap<>();
        eventData.put("user_id", "user456");
        eventData.put("action", "logout");
        stats.recordEvent("user_action", eventData);
        
        // Display statistics
        System.out.println("Counters:");
        stats.getAllCounters().forEach((name, value) -> 
            System.out.printf("  %s: %d%n", name, value));
        
        System.out.println("Gauges:");
        stats.getAllGauges().forEach((name, value) -> 
            System.out.printf("  %s: %d%n", name, value));
        
        System.out.println("Recent events:");
        stats.getRecentEvents(5).forEach(event -> 
            System.out.printf("  %s: %s%n", event.getType(), event.getData()));
    }
    
    private static void testConcurrentPerformance() throws InterruptedException {
        System.out.println("\n--- Concurrent Performance Test ---");
        
        int numThreads = 10;
        int numOperations = 10000;
        
        // Test ConcurrentHashMap
        long startTime = System.nanoTime();
        ConcurrentHashMap<String, Integer> concurrentMap = new ConcurrentHashMap<>();
        
        ExecutorService executor = Executors.newFixedThreadPool(numThreads);
        List<Future<?>> futures = new ArrayList<>();
        
        for (int i = 0; i < numThreads; i++) {
            final int threadId = i;
            Future<?> future = executor.submit(() -> {
                for (int j = 0; j < numOperations; j++) {
                    String key = "key-" + (threadId * numOperations + j);
                    concurrentMap.put(key, j);
                }
            });
            futures.add(future);
        }
        
        // Wait for completion
        for (Future<?> future : futures) {
            try {
                future.get();
            } catch (ExecutionException e) {
                System.err.println("Operation failed: " + e.getMessage());
            }
        }
        
        long endTime = System.nanoTime();
        long concurrentTime = endTime - startTime;
        
        System.out.printf("ConcurrentHashMap: %d operations in %d ms%n",
                         numThreads * numOperations, concurrentTime / 1_000_000);
        System.out.printf("Map size: %d%n", concurrentMap.size());
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
    }
}
```

## 4. Producer-Consumer Pattern

### Various Producer-Consumer Implementations

```java
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.*;

// Message class for producer-consumer
class Message {
    private final String content;
    private final long timestamp;
    private final int priority;
    
    public Message(String content, int priority) {
        this.content = content;
        this.priority = priority;
        this.timestamp = System.currentTimeMillis();
    }
    
    public String getContent() { return content; }
    public long getTimestamp() { return timestamp; }
    public int getPriority() { return priority; }
    
    @Override
    public String toString() {
        return String.format("Message{content='%s', priority=%d, time=%d}", 
                           content, priority, timestamp);
    }
}

// Producer-Consumer using BlockingQueue
class BlockingQueueProducerConsumer {
    private final BlockingQueue<Message> queue;
    private final AtomicInteger producedCount = new AtomicInteger(0);
    private final AtomicInteger consumedCount = new AtomicInteger(0);
    
    public BlockingQueueProducerConsumer(int capacity) {
        this.queue = new LinkedBlockingQueue<>(capacity);
    }
    
    public void start(int numProducers, int numConsumers) {
        ExecutorService executor = Executors.newFixedThreadPool(numProducers + numConsumers);
        
        // Start producers
        for (int i = 0; i < numProducers; i++) {
            executor.submit(new Producer("Producer-" + i));
        }
        
        // Start consumers
        for (int i = 0; i < numConsumers; i++) {
            executor.submit(new Consumer("Consumer-" + i));
        }
        
        // Shutdown after some time
        ScheduledExecutorService shutdownExecutor = Executors.newSingleThreadScheduledExecutor();
        shutdownExecutor.schedule(() -> {
            executor.shutdown();
            try {
                if (!executor.awaitTermination(5, TimeUnit.SECONDS)) {
                    executor.shutdownNow();
                }
            } catch (InterruptedException e) {
                executor.shutdownNow();
            }
            shutdownExecutor.shutdown();
            
            System.out.println("\nFinal Statistics:");
            System.out.println("Produced: " + producedCount.get());
            System.out.println("Consumed: " + consumedCount.get());
            System.out.println("Queue size: " + queue.size());
        }, 10, TimeUnit.SECONDS);
    }
    
    class Producer implements Runnable {
        private final String name;
        
        public Producer(String name) {
            this.name = name;
        }
        
        @Override
        public void run() {
            Random random = new Random();
            
            while (!Thread.currentThread().isInterrupted()) {
                try {
                    // Produce message
                    String content = name + "-Message-" + producedCount.incrementAndGet();
                    int priority = random.nextInt(10);
                    Message message = new Message(content, priority);
                    
                    // Put message in queue (blocks if full)
                    queue.put(message);
                    
                    System.out.printf("%s produced: %s (queue size: %d)%n",
                                     name, message, queue.size());
                    
                    // Random delay
                    Thread.sleep(random.nextInt(100));
                    
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
        }
    }
    
    class Consumer implements Runnable {
        private final String name;
        
        public Consumer(String name) {
            this.name = name;
        }
        
        @Override
        public void run() {
            Random random = new Random();
            
            while (!Thread.currentThread().isInterrupted()) {
                try {
                    // Take message from queue (blocks if empty)
                    Message message = queue.take();
                    consumedCount.incrementAndGet();
                    
                    // Process message
                    Thread.sleep(random.nextInt(50) + 10); // Processing time
                    
                    System.out.printf("%s consumed: %s (queue size: %d)%n",
                                     name, message, queue.size());
                    
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
        }
    }
}

// Producer-Consumer using wait/notify
class WaitNotifyProducerConsumer {
    private final Queue<Message> queue = new LinkedList<>();
    private final int capacity;
    private final AtomicInteger producedCount = new AtomicInteger(0);
    private final AtomicInteger consumedCount = new AtomicInteger(0);
    private final Object lock = new Object();
    
    public WaitNotifyProducerConsumer(int capacity) {
        this.capacity = capacity;
    }
    
    public void produce(Message message) throws InterruptedException {
        synchronized (lock) {
            while (queue.size() >= capacity) {
                System.out.println("Queue is full, producer waiting...");
                lock.wait();
            }
            
            queue.offer(message);
            producedCount.incrementAndGet();
            
            System.out.printf("Produced: %s (queue size: %d)%n", message, queue.size());
            
            lock.notifyAll(); // Notify consumers
        }
    }
    
    public Message consume() throws InterruptedException {
        synchronized (lock) {
            while (queue.isEmpty()) {
                System.out.println("Queue is empty, consumer waiting...");
                lock.wait();
            }
            
            Message message = queue.poll();
            consumedCount.incrementAndGet();
            
            System.out.printf("Consumed: %s (queue size: %d)%n", message, queue.size());
            
            lock.notifyAll(); // Notify producers
            
            return message;
        }
    }
    
    public void startDemo() {
        ExecutorService executor = Executors.newFixedThreadPool(4);
        
        // Producer
        executor.submit(() -> {
            for (int i = 0; i < 20; i++) {
                try {
                    Message message = new Message("Message-" + i, i % 5);
                    produce(message);
                    Thread.sleep(50);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
        });
        
        // Consumer
        executor.submit(() -> {
            for (int i = 0; i < 20; i++) {
                try {
                    consume();
                    Thread.sleep(100);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
        });
        
        executor.shutdown();
        try {
            executor.awaitTermination(10, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        
        System.out.println("\nFinal Statistics:");
        System.out.println("Produced: " + producedCount.get());
        System.out.println("Consumed: " + consumedCount.get());
    }
}

// Producer-Consumer using Semaphore
class SemaphoreProducerConsumer {
    private final Queue<Message> queue = new LinkedList<>();
    private final int capacity;
    private final Semaphore producerSemaphore;
    private final Semaphore consumerSemaphore;
    private final AtomicInteger producedCount = new AtomicInteger(0);
    private final AtomicInteger consumedCount = new AtomicInteger(0);
    
    public SemaphoreProducerConsumer(int capacity) {
        this.capacity = capacity;
        this.producerSemaphore = new Semaphore(capacity); // Producers can fill up to capacity
        this.consumerSemaphore = new Semaphore(0); // Consumers need permits to consume
    }
    
    public void produce(Message message) throws InterruptedException {
        producerSemaphore.acquire(); // Acquire permit to produce
        
        synchronized (queue) {
            queue.offer(message);
            producedCount.incrementAndGet();
            
            System.out.printf("Produced: %s (queue size: %d)%n", message, queue.size());
        }
        
        consumerSemaphore.release(); // Release permit for consumer
    }
    
    public Message consume() throws InterruptedException {
        consumerSemaphore.acquire(); // Acquire permit to consume
        
        synchronized (queue) {
            Message message = queue.poll();
            consumedCount.incrementAndGet();
            
            System.out.printf("Consumed: %s (queue size: %d)%n", message, queue.size());
            
            producerSemaphore.release(); // Release permit for producer
            return message;
        }
    }
    
    public void startDemo() {
        ExecutorService executor = Executors.newFixedThreadPool(4);
        
        // Multiple producers
        for (int i = 0; i < 2; i++) {
            final int producerId = i;
            executor.submit(() -> {
                Random random = new Random();
                for (int j = 0; j < 15; j++) {
                    try {
                        Message message = new Message("Producer-" + producerId + "-Msg-" + j, random.nextInt(10));
                        produce(message);
                        Thread.sleep(random.nextInt(100));
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        // Multiple consumers
        for (int i = 0; i < 2; i++) {
            final int consumerId = i;
            executor.submit(() -> {
                Random random = new Random();
                for (int j = 0; j < 15; j++) {
                    try {
                        consume();
                        Thread.sleep(random.nextInt(150));
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        executor.shutdown();
        try {
            executor.awaitTermination(15, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        
        System.out.println("\nFinal Statistics:");
        System.out.println("Produced: " + producedCount.get());
        System.out.println("Consumed: " + consumedCount.get());
    }
}

// Priority-based producer-consumer
class PriorityProducerConsumer {
    private final PriorityBlockingQueue<Message> queue;
    private final AtomicInteger producedCount = new AtomicInteger(0);
    private final AtomicInteger consumedCount = new AtomicInteger(0);
    
    public PriorityProducerConsumer() {
        // Custom comparator for priority (higher priority = lower number)
        this.queue = new PriorityBlockingQueue<>(11, 
            Comparator.comparingInt(Message::getPriority));
    }
    
    public void start(int numProducers, int numConsumers) {
        ExecutorService executor = Executors.newFixedThreadPool(numProducers + numConsumers);
        
        // Start producers
        for (int i = 0; i < numProducers; i++) {
            final int producerId = i;
            executor.submit(() -> {
                Random random = new Random();
                for (int j = 0; j < 20; j++) {
                    try {
                        String content = "Producer-" + producerId + "-Msg-" + j;
                        int priority = random.nextInt(10);
                        Message message = new Message(content, priority);
                        
                        queue.put(message);
                        producedCount.incrementAndGet();
                        
                        System.out.printf("Produced: %s (priority: %d)%n", message, priority);
                        
                        Thread.sleep(random.nextInt(100));
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        // Start consumers
        for (int i = 0; i < numConsumers; i++) {
            final int consumerId = i;
            executor.submit(() -> {
                Random random = new Random();
                while (!Thread.currentThread().isInterrupted()) {
                    try {
                        Message message = queue.take();
                        consumedCount.incrementAndGet();
                        
                        // Process based on priority
                        int processingTime = (10 - message.getPriority()) * 20 + 10;
                        Thread.sleep(processingTime);
                        
                        System.out.printf("Consumer-%d processed: %s (priority: %d, time: %dms)%n",
                                         consumerId, message, message.getPriority(), processingTime);
                        
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                }
            });
        }
        
        // Shutdown after processing
        ScheduledExecutorService shutdownExecutor = Executors.newSingleThreadScheduledExecutor();
        shutdownExecutor.schedule(() -> {
            executor.shutdownNow();
            shutdownExecutor.shutdown();
            
            System.out.println("\nFinal Statistics:");
            System.out.println("Produced: " + producedCount.get());
            System.out.println("Consumed: " + consumedCount.get());
            System.out.println("Queue size: " + queue.size());
        }, 15, TimeUnit.SECONDS);
    }
}

public class ProducerConsumerDemo {
    public static void main(String[] args) throws InterruptedException {
        System.out.println("=== Producer-Consumer Pattern Demo ===");
        
        // Test BlockingQueue implementation
        System.out.println("\n--- BlockingQueue Producer-Consumer ---");
        BlockingQueueProducerConsumer bqpc = new BlockingQueueProducerConsumer(10);
        bqpc.start(3, 2);
        
        Thread.sleep(12000); // Wait for completion
        
        // Test Wait/Notify implementation
        System.out.println("\n--- Wait/Notify Producer-Consumer ---");
        WaitNotifyProducerConsumer wnpc = new WaitNotifyProducerConsumer(5);
        wnpc.startDemo();
        
        Thread.sleep(2000);
        
        // Test Semaphore implementation
        System.out.println("\n--- Semaphore Producer-Consumer ---");
        SemaphoreProducerConsumer spc = new SemaphoreProducerConsumer(5);
        spc.startDemo();
        
        Thread.sleep(4000);
        
        // Test Priority-based implementation
        System.out.println("\n--- Priority-based Producer-Consumer ---");
        PriorityProducerConsumer ppc = new PriorityProducerConsumer();
        ppc.start(2, 3);
        
        Thread.sleep(17000);
    }
}
```

## Practice Exercises

1. **Synchronization**: Implement a thread-safe banking system with transfer operations
2. **Locks**: Create a resource pool using ReentrantLock with timeout
3. **Concurrent Collections**: Build a real-time analytics dashboard using concurrent collections
4. **Producer-Consumer**: Implement a logging system with multiple producers and consumers

## Interview Questions

1. What's the difference between synchronized and ReentrantLock?
2. When would you use ReadWriteLock over synchronized?
3. What are the advantages of concurrent collections over synchronized collections?
4. How does the producer-consumer pattern help in system design?
5. What's the difference between notify() and notifyAll()?
