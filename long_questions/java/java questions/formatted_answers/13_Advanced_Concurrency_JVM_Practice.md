# 13. Advanced Concurrency and JVM (Practice)

**Q: How to use CompletableFuture for asynchronous programming?**
> "**CompletableFuture** (Java 8) is the modern way to write non-blocking async code. It allows you to chain tasks together like a pipeline.
>
> You start a task with `CompletableFuture.supplyAsync(() -> performTask())`.
> You can then chain what happens next: `.thenApply(result -> process(result))` or `.thenAccept(final -> print(final))`.
>
> The main thread doesn't wait. It continues execution. You can combine multiple futures, waiting for all of them to finish (`allOf`) or just the fastest one (`anyOf`). It's much more powerful than the old `Future` interface."

**Indepth:**
> **Exception Handling**: Standard `Future.get()` throws messy checked exceptions. `CompletableFuture` handles exceptions gracefully inside the pipeline using `.exceptionally()` or `.handle()`, keeping the flow clean.


---

**Q: Difference between synchronized and ReentrantLock?**
> "Both handle locking, but **ReentrantLock** has more features.
>
> **synchronized** is a keyword. It's implicit. You enter the block, you get the lock; you leave, you release it. It's cleaner but rigid. If a thread waits for a synchronized lock, it waits forever.
>
> **ReentrantLock** is a class. You manually call `lock()` and `unlock()`. This gives you power:
> 1.  **tryLock()**: 'Attempt to get the lock, but if it's busy, give up immediately' (or wait for a timeout).
> 2.  **Fairness**: You can set it to 'fair' mode, granting locks in First-Come-First-Served order (preventing starvation).
> 3.  **Interruptible**: A waiting thread can be interrupted."

**Indepth:**
> **Condition**: `ReentrantLock` allows multiple `Condition` objects (like `wait()`/`notify()` but with multiple waiting rooms). You can signal *specific* threads waiting for "Not Empty" vs "Not Full", reducing useless wake-ups.


---

**Q: What is ThreadLocal and when to use it?**
> "**ThreadLocal** allows you to create variables that can only be read and written by the *same thread*. It's like a private pocket for each thread.
>
> Imagine a web server handling 100 requests. Each request is a thread. You want a transaction ID for strictly that request/thread.
> If you save it in a static variable, all threads would overwrite each other.
> If you use `ThreadLocal`, each thread sees its own unique value.
>
> Use it sparingly though—specifically for things like user sessions or transaction contexts. And always `remove()` it when done to prevent memory leaks in thread pools."

**Indepth:**
> **Internals**: It's implemented as a `ThreadLocalMap` inside the `Thread` class itself. The key is the `ThreadLocal` object (weak reference), and the value is your data.


---

**Q: How does a ClassLoader work? (Delegation Model)**
> "When you ask Java to load a class (`MyClass`), it doesn't just look in one place. It follows the **Delegation Hierarchy**:
>
> 1.  It asks the **Application ClassLoader** (your classpath).
> 2.  That delegates to the **Extension/Platform ClassLoader** (lib/ext).
> 3.  That delegates to the **Bootstrap ClassLoader** (core Java libs like `String`, `Object`).
>
> The loading actually happens in reverse order. The Bootstrap tries first. If it can't find it, it goes back down the chain.
> This ensures security: you can't trick Java into loading your own hacked version of `java.lang.String` because the Bootstrap loader will always find the real one first."

**Indepth:**
> **Class Uniqueness**: A class is identified by its **Fully Qualified Name + ClassLoader**. You can have two classes with the precise same name `com.myapp.User` loaded by two different ClassLoaders, and JVM treats them as completely different types!


---

**Q: What are Virtual Threads (Java 21+)?**
> "This is Project Loom. It's revolutionary.
>
> Traditionally, one Java thread `==` one OS thread. OS threads are heavy (consume 1MB RAM) and limited (you can only have a few thousand).
>
> **Virtual Threads** are lightweight threads managed by the **JVM**, not the OS. You can create *millions* of them.
> When a virtual thread waits for I/O (like a database call), the JVM unmounts it and puts another virtual thread on the CPU.
> This allows high-throughput server applications written in a simple, synchronous style without needing complex async callbacks."

**Indepth:**
> **Adoption**: Virtual Threads are designed to work with existing synchronous code (blocking I/O). They do *not* improve CPU-bound tasks (number crunching), only I/O-bound tasks (web servers).


---

**Q: What is the ReadWriteLock?**
> "A standard lock is exclusive—only one thread can enter, period.
>
> A **ReadWriteLock** is smarter. It says: 'Multiple threads can read at the same time, as long as nobody is writing.'
>
> It maintains a pair of locks: a **Read Lock** (shared) and a **Write Lock** (exclusive).
> If you have a system with lots of reads but few writes (like a cache), this drastically improves performance compared to a standard synchronized block."

**Indepth:**
> **Starvation**: A risk with ReadWrite locks is "Writer Starvation". If there are constant readers, the writer might never get the lock. Modern implementations (stamped locks) try to mitigate this.


---

**Q: SOLID - Single Responsibility Principle (SRP) Example**
> "**SRP** states: 'A class should have only one reason to change.'
>
> **Violating SRP**:
> A `User` class that handles user data AND saves the user to the database AND sends a welcome email.
> If you change the database, you touch the User class. If you change the email provider, you touch the User class. Bad.
>
> **Following SRP**:
> Split it into multiple classes:
> 1.  `User`: Just holds data (POJO).
> 2.  `UserRepository`: Handles database operations.
> 3.  `EmailService`: Handles sending emails.
>
> Now, each class focuses on one job. It’s easier to test and maintain."

**Indepth:**
> **Microservices**: SRP applied at the architecture level leads to Microservices. Each service does one thing well (User Service, Email Service, Payment Service).

