# Concurrency & Multithreading - Interview Answers

> 🎯 **Focus:** Concurrency is tricky. These answers focus on safety, modern APIs, and avoiding common pitfalls like deadlocks.

### 1. Difference between `Thread` and `Runnable`?
"It’s the difference between the **worker** and the **work**.

`Thread` is a class that represents the actual OS thread—the worker.
`Runnable` is an interface that represents the task to be done—the work.

You should almost always implement `Runnable` (or `Callable`) rather than extending `Thread`. Extending `Thread` limits you because Java only supports single inheritance, so you can't extend anything else. Plus, by using `Runnable`, you decouple the task from the execution policy, allowing you to pass the same task to a Thread Pool easily."

---

### 2. Difference between `Callable` and `Runnable`?
"`Runnable` is the older interface. Its `run()` method returns `void` and cannot throw checked exceptions. It’s strictly 'fire and forget'.

`Callable` is the modern upgrade introduced with the Executor framework. Its `call()` method returns a result (via Generics) and can throw exceptions.

When I submit a `Callable` to a thread pool, I get back a `Future` object, which is a handle I can use to retrieve the result once the task is finished."

---

### 3. What is `ExecutorService`?
"It’s a framework that manages thread lifecycles for you, so you don't have to manually do `new Thread().start()`.

Creating threads is expensive for the OS. `ExecutorService` maintains a **Pool** of reusable threads. You just throw tasks at it (`submit()`), and it assigns them to available threads.

In production, I typically use `ThreadPoolExecutor` with a fixed size or a cached pool to prevent creating too many threads and crashing the application under load."

---

### 4. What is a Deadlock? How to prevent it?
"A deadlock is when two threads are stuck waiting for each other forever. Thread A holds Lock 1 and wants Lock 2. Thread B holds Lock 2 and wants Lock 1. Neither can proceed.

To prevent it, the golden rule is **Lock Ordering**. If all threads acquire locks in the exact same order (e.g., always Lock A, then Lock B), a deadlock is mathematically impossible.

Another safeguard is using `tryLock()` with a timeout. If a thread can't get a lock within 2 seconds, it gives up and backs off, rather than waiting forever."

---

### 5. `volatile` vs `synchronized`?
"`synchronized` provides both **atomicity** (mutual exclusion) and **visibility** (memory changes are seen by others). It essentially says, 'Only one thread can do this at a time.'

`volatile` is lighter. It only provides **visibility**. It tells the compiler and CPU, 'Do not cache this variable in a register; always read it from main memory.'

However, `volatile` is **not** atomic. If you do `count++` on a volatile variable, you can still have race conditions. So I use `volatile` mostly for simple flags like `stopping = true`, but for counters, I use `AtomicInteger`."

---

### 6. What is `CompletableFuture`?
"It’s Java's way of doing asynchronous programming properly, similar to Promises in JavaScript.

Before Java 8, we had `Future`, but it was blocking—you had to call `.get()` and wait. `CompletableFuture` lets you build non-blocking pipelines. You can say, 'Do task A, **then** do task B with the result, **then** handle any errors,' all without blocking the main thread.

I use it heavily for microservices—like calling 3 external APIs in parallel and then combining their results."

---

### 7. What involves the `ThreadLocal` class?
"It allows you to create variables that can only be read and written by the **same thread**. It’s like a private instance variable for a thread.

This is super useful for things that aren't thread-safe, like `SimpleDateFormat` (in older Java) or for storing context that needs to travel with the request, like a User ID or Transaction ID.

However, you have to be careful to `remove()` the value when the thread finishes work (especially in thread pools), otherwise you get memory leaks because the thread stays alive and keeps holding onto that data."

---

### 8. `synchronized` block vs method?
"Functionally they do the same thing, but a **block** gives you more control.

With a synchronized **method**, you lock the entire object (`this`). This can be a performance bottleneck if the method is long.

With a **block**, you can lock a smaller section of code (critical section) and use a specific lock object. This reduces the scope of the lock, meaning threads spend less time waiting. I generally prefer blocks or explicit `ReentrantLocks` for this reason."

---

### 9. What is `AtomicInteger`?
"It’s a thread-safe integer that uses non-blocking CPU instructions (CAS - Compare And Swap) instead of locks.

Because it doesn't use `synchronized`, it’s much faster for simple things like counters or sequence generators. If I need to count active users in a concurrent environment, `AtomicInteger.incrementAndGet()` is the standard way to do it."

---

### 10. `wait()` vs `sleep()`?
"The key difference is ownership of the lock.

`sleep()` is a static method on `Thread`. It pauses the execution but **keeps** any locks the thread is holding. It’s like taking a nap at your desk—you're not working, but you're still occupying the seat.

`wait()` is a method on `Object`. It releases the lock and waits to be notified by another thread. It’s like leaving the desk and waiting in the hall until someone calls you back. It’s used for inter-thread communication."
