# Java Multithreading — High-Speed Retrieval Cheatsheet

> **Interview-Ready Framework**: 5-Step Mental Model for Instant Recall

---

## 🔧 The 'Underlying Engine' Map

### **Category 1: Thread Management & Execution**
**Logic Pattern**: Task Creation → Thread Assignment → Execution Control
- **Thread/Runnable**: Basic task execution → Simple concurrent work
- **ExecutorService**: Thread pool management → Resource-efficient execution
- **Callable/Future**: Task with return value → Asynchronous results
- **ScheduledExecutorService**: Time-based execution → Periodic/delayed tasks

### **Category 2: Synchronization & Coordination**
**Logic Pattern**: Shared Resource → Access Control → Thread Safety
- **synchronized**: Method/block level locking → Basic mutual exclusion
- **ReentrantLock**: Flexible locking → Advanced lock patterns
- **ReadWriteLock**: Read/write separation → Optimized concurrent access
- **Atomic classes**: Lock-free operations → High-performance counters

### **Category 3: Concurrent Data Structures**
**Logic Pattern**: Data Container → Thread-Safe Operations → Performance
- **ConcurrentHashMap**: Thread-safe hashing → Concurrent key-value access
- **BlockingQueue**: Producer-consumer → Thread coordination
- **CopyOnWriteArrayList**: Read-optimized → Infrequent writes, many reads
- **AtomicReference**: Object atomicity → Thread-safe object updates

### **Category 4: Advanced Patterns**
**Logic Pattern**: Business Problem → Concurrent Solution → Real-World Application
- **Rate Limiting**: Request throttling → API protection
- **Event Bus**: Publisher-subscriber → Decoupled communication
- **Resource Pooling**: Connection management → Efficient resource usage
- **Priority Scheduling**: Task ordering → Critical job processing

---

## 🚨 The 'Red-Flag' Failure Section

### **Critical Runtime Errors**

| **Error Type** | **Trigger** | **Example** | **Fix** |
|---|---|---|---|
| **NullPointerException** | Null shared resource | `synchronized(nullObj)` | Validate before sync |
| **IllegalMonitorStateException** | Wrong monitor unlock | `lock.unlock()` without `lock()` | Use try-finally |
| **InterruptedException** | Thread interruption | `Thread.sleep()` interrupted | Check interrupt flag |
| **RejectedExecutionException** | Full thread pool | `executor.submit()` after shutdown | Use proper shutdown |
| **TimeoutException** | Resource wait timeout | `queue.poll(timeout)` | Handle timeout gracefully |

### **Logic Failure Patterns**

- **Deadlock**: Circular lock dependencies → A waits for B, B waits for A
- **Race Condition**: Unsynchronized shared state → Inconsistent results
- **Starvation**: Low-priority threads never execute → Priority inversion
- **Livelock**: Threads keep responding to each other → No progress
- **Memory Consistency**: Stale data across threads → Visibility issues
- **Resource Exhaustion**: Too many threads → OutOfMemoryError

### **Performance Killers**

- **Over-synchronization**: Unnecessary locking → Bottleneck
- **Thread creation overhead**: `new Thread()` vs pool → Resource waste
- **Context switching**: Too many threads → CPU overhead
- **False sharing**: Cache line contention → Performance degradation

---

## ⚡ The 'Performance & Complexity' Table

| **Operation** | **Time Complexity** | **Memory Usage** | **Best Use Case** |
|---|---|---|---|
| **synchronized** | O(1) (contention) | O(1) | Simple critical sections |
| **ReentrantLock** | O(1) | O(1) | Complex lock patterns |
| **AtomicInteger** | O(1) | O(1) | High-performance counters |
| **ConcurrentHashMap** | O(log n) | O(n) | Concurrent key-value store |
| **BlockingQueue** | O(1) | O(n) | Producer-consumer |
| **CopyOnWriteArrayList** | O(n) write, O(1) read | O(n) | Read-heavy workloads |
| **Thread Pool** | O(1) submit | O(pool size) | Reusable threads |
| **ReadWriteLock** | O(1) | O(1) | Read-dominant access |

### **Performance Rules**
- **Lock granularity**: Fine-grained vs coarse-grained → Trade-off complexity vs contention
- **Lock-free algorithms**: CAS operations → Better scalability under contention
- **Thread pool sizing**: CPU-bound = cores, I/O-bound = more threads
- **Queue selection**: Bounded vs unbounded → Memory vs throughput trade-off

---

## 🛡️ The 'Safe vs. Risky' Comparison

### **Standard/Safe Methods**
```java
// ✅ SAFE: Modern, recommended approaches
ExecutorService executor = Executors.newFixedThreadPool(4);
ConcurrentHashMap<String, Integer> map = new ConcurrentHashMap<>();
AtomicInteger counter = new AtomicInteger(0);
ReentrantLock lock = new ReentrantLock();
```

### **Legacy/Dangerous Methods**
```java
// ❌ RISKY: Avoid in modern code
synchronized(this) { /* on public objects */ }
new Thread().start(); // Without pool management
Thread.stop(); // Deprecated and dangerous
wait()/notify() // Complex error-prone pattern
```

### **Why Use X Over Y?**

| **Safe Choice** | **Risky Alternative** | **Reason** |
|---|---|---|
| `ExecutorService` | `new Thread()` | Thread reuse, resource management |
| `ConcurrentHashMap` | `HashMap + synchronized` | Better scalability, fine-grained locking |
| `AtomicInteger` | `int + synchronized` | Lock-free, better performance |
| `ReentrantLock` | `synchronized` | Try-lock, interruptible, fair option |
| `BlockingQueue` | `wait()/notify()` | Simpler, less error-prone |
| `ReadWriteLock` | `synchronized` | Read optimization |

---

## 🎯 The 'Interview Logic' Column

### **Core Concepts with Analogies & Golden Rules**

| **Concept** | **Real-World Analogy** | **Golden Rule** |
|---|---|---|
| **Thread Pool** | **Restaurant Staff**: Reusable waiters serving many tables | "Never create threads manually - use pools for resource efficiency" |
| **Synchronization** | **Single Bathroom**: Only one person at a time | "Protect shared mutable state - synchronize or make immutable" |
| **ConcurrentHashMap** | **Multi-Lane Highway**: Multiple cars traveling simultaneously | "Use concurrent collections for shared data - better than manual synchronization" |
| **Producer-Consumer** | **Assembly Line**: Workers place items, others process them | "Use BlockingQueue for thread coordination - eliminates wait/notify complexity" |
| **Atomic Operations** | **ATM Transaction**: Single, indivisible operation | "Use atomic classes for simple shared state - lock-free performance" |
| **Rate Limiting** | **Club Bouncer**: Controls entry rate to prevent overcrowding | "Implement rate limiting for API protection - use sliding window algorithm" |
| **Event Bus** | **Radio Station**: Broadcasts, multiple receivers tune in | "Use event bus for loose coupling - publisher doesn't need to know subscribers" |
| **Resource Pool** | **Library Books**: Borrow, use, return for others | "Pool expensive resources - reuse rather than recreate" |

### **Quick Interview Decision Tree**

1. **Need to run tasks concurrently?** → `ExecutorService` (never raw threads)
2. **Need to share data between threads?** → `ConcurrentHashMap` or `Atomic*`
3. **Need producer-consumer pattern?** → `BlockingQueue`
4. **Need complex locking?** → `ReentrantLock` (not `synchronized`)
5. **Need read-heavy operations?** → `ReadWriteLock` or `CopyOnWriteArrayList`
6. **Need to coordinate many threads?** → `CountDownLatch` or `CyclicBarrier`

---

## 📚 Mental Index Cards for Rapid Recall

### **Card 1: Thread Management Quick Reference**
```
ExecutorService pool = Executors.newFixedThreadPool(4);
Future<T> future = pool.submit(callable);
T result = future.get(); // Blocks for result
pool.shutdown(); // Graceful shutdown
```

### **Card 2: Synchronization Patterns**
```
// Basic
synchronized(lockObject) { /* critical section */ }

// Advanced
ReentrantLock lock = new ReentrantLock();
lock.lock();
try { /* critical section */ } finally { lock.unlock(); }
```

### **Card 3: Concurrent Collections**
```
ConcurrentHashMap<K,V> map = new ConcurrentHashMap<>();
BlockingQueue<T> queue = new LinkedBlockingQueue<>();
AtomicInteger counter = new AtomicInteger(0);
```

### **Card 4: Producer-Consumer Template**
```
BlockingQueue<T> queue = new LinkedBlockingQueue<>();
// Producer
queue.put(item);
// Consumer
T item = queue.take();
```

---

## 🔥 Top 10 Interview Patterns

1. **Thread Pool Pattern**: `ExecutorService + Callable + Future`
2. **Producer-Consumer**: `BlockingQueue + multiple threads`
3. **Shared State**: `ConcurrentHashMap + atomic operations`
4. **Resource Pooling**: `Semaphore + object reuse`
5. **Rate Limiting**: `Sliding window + thread-safe counters`
6. **Event-Driven**: `Event bus + subscriber pattern`
7. **Read-Write Lock**: `ReadWriteLock for read-heavy workloads`
8. **CountDownLatch**: `One-time coordination + multiple threads`
9. **CyclicBarrier**: `Repeated synchronization points`
10. **Atomic Updates**: `Compare-and-swap + lock-free programming`

---

## 🚨 Critical Interview Red Flags to Avoid

### **Never Say These in Interviews**
- ❌ "I use `synchronized` for everything"
- ❌ "I create new threads for each task"
- ❌ "I use `Thread.stop()` to stop threads"
- ❌ "I don't worry about deadlocks"
- ❌ "HashMap is fine for multi-threading"

### **Always Mention These**
- ✅ Thread pool sizing strategy
- ✅ Lock granularity considerations
- ✅ Memory visibility (volatile, final)
- ✅ Exception handling in concurrent code
- ✅ Graceful shutdown procedures

---

**Interview Strategy**: Start with the real-world analogy, identify the core problem (resource management, coordination, shared state), pick the appropriate pattern from this cheatsheet, and always discuss thread safety, performance implications, and shutdown procedures.
