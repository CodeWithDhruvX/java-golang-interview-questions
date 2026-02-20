# 11. Concurrency (Practice)

**Q: Difference between Thread and Runnable?**
> "In Java, you have two main ways to create a thread.
>
> One is to **extend the Thread class**. You write `class MyThread extends Thread`. The problem is, Java only allows single inheritance. If you extend `Thread`, you can't extend anything else.
>
> The better way is to **implement the Runnable interface**. You write `class MyTask implements Runnable`. This leaves your class free to extend another class if needed. You then pass your `Runnable` instance to a `Thread` object to run it.
>
> So, rule of thumb: Always implement `Runnable` (or `Callable`) unless you are actually modifying the behavior of the Thread class itself."

**Indepth:**
> **Virtual Threads**: In Java 21+, `Thread.ofVirtual().start(runnable)` creates lightweight threads mapped to OS threads by the JVM. Implementing `Runnable` makes upgrading to Virtual Threads trivial.


---

**Q: What is a deadlock? How do you create one?**
> "A deadlock is a standoff. It happens when two threads are stuck forever, each waiting for the other to release a lock.
>
> Imagine Thread A holds 'Key 1' and wants 'Key 2'.
> At the same time, Thread B holds 'Key 2' and wants 'Key 1'.
> neither will give up what they have, so they wait forever.
>
> Creating one is easy: Just have two threads and two locks. Make Thread A lock Lock1 then wait for Lock2. Make Thread B lock Lock2 then wait for Lock1. Run them together, and your app freezes."

**Indepth:**
> **Prevention**: Avoid nested locks. If you must use multiple locks, always acquire them in the *same order* (e.g., always lock Resource A before Resource B) to prevent cycles.


---

**Q: What is ExecutorService?**
> "Managing raw threads (creating them, starting them) is expensive and error-prone.
>
> **ExecutorService** is a higher-level framework introduced in Java 5 to handle this. It manages a **pool of threads** for you.
>
> Instead of saying `new Thread(task).start()`, you say `executor.submit(task)`.
> It recycles threads (saving memory), handles scheduling, and gives you `Future` objects to track progress. You should almost always use this instead of raw Threads."

**Indepth:**
> **Shutdown**: Don't forget to shut it down! `executor.shutdown()` stops accepting new tasks. If you don't call it, your app might never exit because the pool threads are still alive.


---

**Q: Difference between Callable and Runnable?**
> "They are both tasks meant for threads, but they have key differences.
>
> **Runnable** is the old one (Java 1.0). Its method is `public void run()`. It doesn't return anything, and it *cannot* throw a checked exception to the caller.
>
> **Callable** is the new one (Java 5). Its method is `public T call()`. It *returns a result* (T), and it *can* throw an exception.
>
> If you need a value back from your thread (like a calculation result), use `Callable`."

**Indepth:**
> **Future**: `submit(callable)` returns a `Future<T>`. Calling `future.get()` blocks the current thread until the result is ready (or throws an exception).


---

**Q: What are atomic classes (AtomicInteger)?**
> "In multi-threading, simple operations like `count++` are not safe. It actually involves three steps: read, increment, write. If two threads do it at the same time, you lose data.
>
> **Atomic Classes** (like `AtomicInteger`) provide a way to perform these operations safely *without* using heavy locks (`synchronized`).
>
> They use low-level CPU instructions (CAS - Compare-And-Swap) to ensure that `incrementAndGet()` happens atomically. It's much faster than using synchronization for simple counters."

**Indepth:**
> **Non-blocking**: Synchronized blocks put threads to sleep (blocking). Atomics spin or use hardware instructions (non-blocking), which is much more scalable under high contention.


---

**Q: Thread-safe Singleton (Double-Checked Locking)**
> "If you are creating a Singleton lazily (creating it only when asked), you have to be careful.
>
> If two threads call `getInstance()` at the same time, and the instance is null, both might create a new object!
>
> To fix this efficiently, we use **Double-Checked Locking**:
> 1.  Check if instance is null (no locking, fast).
> 2.  If null, enter a `synchronized` block.
> 3.  Check if instance is null *again* (just in case another thread beat us to it while we were waiting to enter the block).
> 4.  Create the instance.
>
> Also, the instance variable must be marked `volatile` to prevent instruction reordering issues."

**Indepth:**
> **Volatile**: Without `volatile`, the "Create instance" step (allocate memory, init variables, assign reference) can be reordered. Another thread might see a non-null but *partially initialized* object and crash.

