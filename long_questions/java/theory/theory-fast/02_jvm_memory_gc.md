# JVM, Memory & GC Interview Questions (26-35)

## JVM Architecture & Memory Management

### 26. What are the JVM memory areas?
"The JVM memory is broadly divided into the **Stack**, which handles method execution and local variables, and the **Heap**, where objects live.

Inside the Heap, you have the **Young Generation** (Eden, Survivor spaces) for new objects, and the **Old Generation** for long-lived objects.

Then there's the **Metaspace** (which replaced PermGen in Java 8). This stores class metadata, static variables, and method definitions in native memory.

It's crucial to distinguish these because StackOverflowError happens in the Stack, while OutOfMemoryError usually happens in the Heap or Metaspace."

### 27. Difference between Heap and Stack?
"Think of **Stack** as thread-local scratchpad memory. Every thread has its own stack. It stores primitive variables and references to objects. When a method exits, its stack frame is popped, and that memory is instantly reclaimed.

**Heap** is the shared memory area for the entire application. Every object created with `new` goes here. It’s managed by the Garbage Collector, meaning memory isn't freed until the GC runs.

So, Stack is fast, thread-safe, and self-cleaning. Heap is larger, slower, and requires management."

### 28. What is garbage collection?
"Garbage Collection is Java's automatic memory management process. Its job is to identify objects in the Heap that are no longer reachable from any 'GC Root' and reclaim their memory.

It essentially works in two steps: **Mark** (tagging live objects) and **Sweep** (clearing dead ones). Modern collectors also do **Compaction** to defrag memory.

The beauty is we don't have to manually `free()` memory like in C++, but the tradeoff is occasional pauses."

### 29. Difference between minor GC and major GC?
"A **Minor GC** (or Young GC) happens frequently. It collects garbage from the Young Generation (Eden space). It's usually very fast because most objects die young.

A **Major GC** (or Full GC) cleans the Old Generation. This is the expensive operation. It happens when the Old Gen fills up.

Major GCs often trigger 'Stop-The-World' events where the entire application freezes for a moment. Tuning JVM performance is mostly about minimizing the frequency and duration of these Major GCs."

### 30. What causes `OutOfMemoryError`?
"OOM (OutOfMemoryError) happens when the JVM cannot allocate an object because the Heap is full, and the Garbage Collector cannot free up any space.

Common causes are memory leaks—like adding objects to a static List and never removing them—or simply under-provisioning the heap size (`-Xmx`) for the application's load.

It can also happen in Metaspace if you're dynamically generating too many classes, or even 'GC Overhead Limit Exceeded' if the GC is spending 98% of its time freeing only 2% of memory."

### 31. Difference between `OutOfMemoryError` and `StackOverflowError`?
"`OutOfMemoryError` simply means the Heap is full. You ran out of space for *objects*.

`StackOverflowError` happens when the stack depth limit is exceeded. This almost always happens due to infinite recursion—Method A calls Method A calls Method A... until the stack blows up.

You fix OOM by fixing leaks or increasing heap size. You fix StackOverflow by fixing your recursive logic."

### 32. What are GC roots?
"GC Roots are the starting points for the Garbage Collector's reachability analysis. If an object is reachable from a GC Root, it is considered 'alive' and won't be collected.

Common GC roots include:
1.  **Local variables** currently on the Stack (in active methods).
2.  **Static variables** in loaded classes.
3.  Active **Threads**.
4.  JNI references (native code).

If an object isn't connected to any of these roots, it's eligible for garbage collection."

### 33. What is stop-the-world?
"Stop-The-World is a pause where the JVM halts *all* application threads to perform a GC cycle safely.

It’s necessary because you can't have threads moving objects around while the GC is trying to count and compact them.

In Minor GC, these pauses are negligible. In Major GC (especially with older collectors like ParallelOld), these pauses can last hundreds of milliseconds or even seconds, causing visible lag for users. Modern collectors like ZGC or Shenandoah aim for sub-millisecond pauses to solve this."

### 34. How do you analyze memory leaks?
"First, I enable GC logs to see *when* the OOM is happening.

Then, I capture a **Heap Dump** (a snapshot of memory) using tools like `jmap` or VisualVM.

I analyze that dump using a tool like **Eclipse MAT** (Memory Analyzer Tool) or IntelliJ's profiler. I look for the 'Dominator Tree' to see which objects are retaining the most memory. Usually, it's a large Collection or Cache that just keeps growing. Once I find who holds the reference, I check the code to see why we aren't clearing it."

### 35. What JVM options have you used?
"I usually set explicit Heap sizes to avoid resizing overhead: `-Xms` (min heap) and `-Xmx` (max heap). I generally set them to the same value in production containers.

I also configure the GC: `-XX:+UseG1GC` is my default for most services.

For debugging, I always add `-XX:+HeapDumpOnOutOfMemoryError` and `-XX:HeapDumpPath=/logs`. This ensures that if the app crashes at 3 AM, I have a snapshot to debug in the morning."
