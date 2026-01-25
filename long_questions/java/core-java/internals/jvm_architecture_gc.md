# JVM Architecture & Garbage Collection

## 1. JVM Architecture
The JVM (Java Virtual Machine) is the engine that drives the Java code. It converts Java Bytecode into machine language.

### Main Components
1.  **Class Loader Subsystem**: Loads `.class` files into memory.
    *   *Bootstrap CL*: Loads core java libs (rt.jar).
    *   *Extension CL*: Loads ext folder.
    *   *Application CL*: Loads classpath.
2.  **Runtime Data Areas (Memory)**:
    *   **Method Area**: Metadata, static variables, code.
    *   **Heap**: Objects. Shared.
    *   **Stack**: Local variables, method calls. Thread-Private.
    *   **PC Registers**: Current instruction address.
    *   **Native Method Stack**: For C/C++ native code.
3.  **Execution Engine**:
    *   **Interpreter**: Reads bytecode stream then executes the instructions.
    *   **JIT Compiler**: Compiles hot code paths to native machine code for performance.
    *   **Garbage Collector**: Reclaims memory.

## 2. Heap Structure
The Heap is divided into generations to optimize Garbage Collection (GC).

### Young Generation
*   **Eden Space**: Where new objects are created.
*   **Survivor Spaces (S0, S1)**: Objects that survive a minor GC move here.
*   *GC Type*: **Minor GC**. Very fast. "Stop the world" but short.

### Old Generation (Tenured)
*   Objects that survive multiple Minor GCs move here.
*   *GC Type*: **Major GC** (or Full GC). Identifying and cleaning old objects. Slower.

### Metaspace (Java 8+)
*   Replaced "PermGen". Stores class metadata.
*   Located in native memory (outside heap), grows automatically.

## 3. Garbage Collection Algorithms

### Mark and Sweep
1.  **Mark**: Traverse object graph from GC Roots (Stack variables, Statics) and mark reachable objects.
2.  **Sweep**: Delete everything not marked.

### G1 GC (Garbage First) - Default in Java 9+
*   Divides heap into small equal-sized regions.
*   Concurrently marks regions.
*   Collects regions with the *most garbage* first (hence the name).
*   **Pros**: Predictable pause times (you can set a target like "max 10ms pause").

### ZGC (Z Garbage Collector) - Java 15+
*   Scalable low-latency GC.
*   **Goal**: Pauses never exceed 10ms, even on Terabyte-sized heaps.
*   Uses colored pointers and load barriers.

## 4. GC Tuning Flags
*   `-Xms`: Initial Heap Size.
*   `-Xmx`: Max Heap Size.
*   `-XX:+UseG1GC`: Enable G1 GC.
*   `-XX:MaxGCPauseMillis=200`: Target pause time.

## 5. Interview Questions
1.  **What is a Memory Leak in Java?**
    *   *Ans*: Objects are no longer needed but are still referenced (e.g., inside a static Map), so GC cannot delete them. Eventually causes `OutOfMemoryError`.
2.  **Difference between Stack and Heap?**
    *   *Ans*: Stack stores primitives and references; it is thread-local and fast. Heap stores actual objects; it is shared and slower.
3.  **Why does `System.gc()` not guarantee execution?**
    *   *Ans*: It is just a hint to the JVM. The JVM decides when to run GC based on memory pressure.
