# JVM & Performance - Interview Answers

> 🎯 **Focus:** These answers show you understand what happens "under the hood" of your application.

### 1. Structure of JVM Memory?
"The JVM memory is divided into a few key areas.

First is the **Heap**. This is where all objects live. It's the biggest area and is managed by the Garbage Collector.
Then there's the **Stack**. Each thread has its own stack for local variables and method calls. It's fast and automatically cleaned up when a method returns.
We also have the **Metaspace** (formerly PermGen), which stores class metadata and static variables.
And finally, the **Code Cache**, where the JIT compiler stores optimized machine code.

Understanding this helps when debugging exceptions—`OutOfMemoryError` usually means Heap is full, while `StackOverflowError` implies infinite recursion."

---

### 2. How does Garbage Collection (GC) work?
"GC is basically a daemon thread that looks for 'unreachable' objects—objects that no part of your live code points to anymore.

It typically works in 'generations.' New objects are born in the **Eden Space**. If they survive a few GC cycles, they are promoted to the **Old Generation**.

The idea is 'Weak Generational Hypothesis'—most objects die young. So the GC runs frequently and quickly on the Young Gen (Minor GC), and infrequently on the Old Gen (Major GC), because cleaning the Old Gen is expensive and can pause the application (Stop-the-World)."

---

### 3. Difference between G1 and Parallel GC?
"These are different algorithms for managing the heap.

**Parallel GC** is all about throughput. It's great for batch processing where you don't care if the app pauses for a second, as long as the job finishes fast overall.

**G1 (Garbage First)** is the default in modern Java. It's designed for low latency. It slices the heap into small regions and cleans them incrementally to keep pause times short and predictable.

If I'm running a user-facing REST API, I almost always stick with G1 (or ZGC in newer versions) to avoid lag spikes."

---

### 4. Stack vs Heap memory?
"**Stack** is for execution context. It stores primitives and references to objects for the currently executing method. It’s LIFO (Last In First Out) and strictly thread-local.

**Heap** is for data storage. It stores the actual objects. It’s shared globally across all threads.

So if I say `Person p = new Person();` — the reference `p` sits on the Stack, but the actual `Person` object sits on the Heap."

---

### 5. What causes a Memory Leak in Java?
"Even though we have GC, leaks happen when we unintentionally keep references to unused objects, so the GC *can't* clean them up.

The most common culprit is **static collections**. If I put objects into a `static List` or `Map` and never remove them, they stay there forever because static fields act as GC roots.

Another common one is unclosed resources—like keeping a DB connection open—or using `ThreadLocal` variables and forgetting to clear them in a thread pool."

---

### 6. Difference between `ClassNotDef` and `ClassNotFound` exceptions?
"They sound identical but happen at different times.

`ClassNotFoundException` is a checked exception. It happens when you try to load a class dynamically (like `Class.forName()`) and it’s not on the classpath. It’s usually a configuration issue.

`NoClassDefFoundError` is an Error (unchecked). It happens when the class *was* present during compilation, but is missing (or has a different initialization error) at runtime. This often happens in Maven dependency hell, where you have version conflicts."

---

### 7. What is JIT (Just-In-Time) compilation?
"Java bytecode is interpreted, which can be slow. The JIT compiler monitors the code as it runs.

If it sees a method being called frequently ('hot method'), it compiles that bytecode into native machine code directly optimized for that CPU.

This is why Java applications often run faster after a 'warm-up' period—the JIT has had time to optimize the critical paths."

---

### 8. Strong vs Weak vs Soft References?
"These define how aggressively the GC cleans up an object.

**Strong Reference** is the default (`Object o = new Object()`). GC will never touch it as long as the reference exists.
**Soft Reference** is for caches. GC will only clean it up if it's absolutely running out of memory. It tries to keep it as long as possible.
**Weak Reference** is weaker. If the GC sees it, it collects it immediately, regardless of memory pressure. This is useful for metadata mappings like `WeakHashMap`."

---

### 9. How to analyze an `OutOfMemoryError`?
"First, I look at the logs to see *which* space is full—is it Heap Space or Metaspace?

Then, I'd capture a **Heap Dump** (using `jmap` or `-XX:+HeapDumpOnOutOfMemoryError`).
I open that dump in a tool like **Eclipse MAT** or **VisualVM**.

I look for the 'Dominator Tree'—basically, which big objects are holding onto references of millions of small objects. Usually, it’s a giant List or Map that kept growing indefinitely."

---

### 10. Java 8 vs Java 11 vs Java 17/21?
"**Java 8** was the big revolution—Lambdas and Streams.
**Java 11** was the first LTS (Long Term Support) after 8. It gave us the `HTTP Client`, `var` syntax, and removed standard modules like JAXB.
**Java 17** brought Records, Sealed Classes, and text blocks—huge for readability.
**Java 21** is the new cool kid with Virtual Threads, which fundamentally changes how we handle high-concurrency throughput.

We generally upgrade to the next LTS version (8 -> 11 -> 17) to stay secure and get free performance wins from JVM improvements."
