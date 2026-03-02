# 🔬 04 — JVM Internals
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- JVM memory structure (Heap, Stack, Metaspace)
- Garbage Collection algorithms (G1, ZGC, Shenandoah)
- Class loading mechanism
- JIT (Just-In-Time) compilation
- Java memory model and escape analysis
- GC tuning flags

---

## ❓ Most Asked Questions

### Q1. What is the JVM memory structure?

```text
JVM MEMORY LAYOUT:
┌──────────────────────────────────────────────────────────┐
│ HEAP                                                      │
│  ┌─────────────────────┐  ┌───────────────────────────┐  │
│  │ Young Generation    │  │   Old Generation (Tenured)│  │
│  │  ┌────┬────┬─────┐  │  │   (long-lived objects)    │  │
│  │  │Eden│ S0 │ S1  │  │  │                           │  │
│  │  └────┴────┴─────┘  │  │                           │  │
│  └─────────────────────┘  └───────────────────────────┘  │
├──────────────────────────────────────────────────────────┤
│ METASPACE (Java 8+, replaces PermGen)                    │
│  Class metadata, static fields, method code              │
├──────────────────────────────────────────────────────────┤
│ THREAD STACKS (one per thread)                           │
│  Stack frames, local variables, operand stack            │
├──────────────────────────────────────────────────────────┤
│ CODE CACHE — compiled JIT native code                    │
│ DIRECT MEMORY — NIO ByteBuffer.allocateDirect()          │
└──────────────────────────────────────────────────────────┘
```

```java
// Heap sizing flags
// -Xms512m          — initial heap size
// -Xmx2g            — maximum heap size
// -XX:MaxMetaspaceSize=256m  — limit Metaspace
// -Xss256k          — thread stack size (reduce for more threads)

// Check memory at runtime
Runtime rt = Runtime.getRuntime();
long totalMemory = rt.totalMemory() / (1024 * 1024);  // MB
long freeMemory  = rt.freeMemory()  / (1024 * 1024);
long maxMemory   = rt.maxMemory()   / (1024 * 1024);
System.out.printf("Heap: %dMB used / %dMB total / %dMB max%n",
    (totalMemory - freeMemory), totalMemory, maxMemory);
```

---

### Q2. How does Garbage Collection work?

```text
GENERATIONAL GC THEORY:
- Most objects die young (short-lived temporaries)
- Eden: new objects allocated here
- Minor GC:  Eden + S0 → S1 (survivors age++); fast, frequent
- Major GC:  Old Gen when full; slow, infrequent (STW — Stop-The-World)
- Full GC:   Entire heap; worst-case — avoid in production!

OBJECT LIFECYCLE:
new Object() → Eden → [Minor GC] → Survivor S0 → ... → Old Gen (age threshold)
                                                    ↑
                                            (default age=15)

GC ALGORITHMS (Java 11+):
┌─────────────────┬───────────┬──────────┬─────────────────────────┐
│ Collector       │ Throughput│ Latency  │ Use Case                │
├─────────────────┼───────────┼──────────┼─────────────────────────┤
│ Serial GC       │ Low       │ High STW │ Single CPU, small apps  │
│ Parallel GC     │ High      │ High STW │ Batch processing        │
│ G1GC (default)  │ Good      │ < 200ms  │ Most apps, heap > 4GB   │
│ ZGC (Java 15+)  │ Good      │ < 10ms   │ Low-latency services    │
│ Shenandoah      │ Good      │ < 10ms   │ Ultra-low latency       │
└─────────────────┴───────────┴──────────┴─────────────────────────┘
```

```bash
# GC tuning flags
-XX:+UseG1GC                    # Use G1 (default Java 9+)
-XX:MaxGCPauseMillis=200        # target pause goal
-XX:G1HeapRegionSize=16m        # G1 region size (1-32 MB)
-XX:+UseZGC                     # Use ZGC for < 10ms pauses
-XX:+PrintGCDetails             # log GC events
-XX:+PrintGCDateStamps          # add timestamps
-Xlog:gc*:file=gc.log           # Java 11 unified logging
-XX:+HeapDumpOnOutOfMemoryError # auto-dump on OOM
-XX:HeapDumpPath=/tmp/heapdump.hprof
```

---

### Q3. What is class loading in Java?

```java
// 3 phases: Loading → Linking (Verify, Prepare, Resolve) → Initialization

// CLASS LOADER HIERARCHY:
// Bootstrap ClassLoader (native) — loads java.lang.*, java.util.*
//   └► Extension/Platform ClassLoader — loads JDK extension modules
//         └► Application ClassLoader — loads app classes from classpath
//               └► Custom ClassLoader — loads from custom sources (URL, DB, etc.)

// Parent Delegation:
// Child asks parent first → Bootstrap → Platform → Application → custom
// Ensures java.lang.String is always loaded by Bootstrap (not shadowed)

// Custom ClassLoader example
public class HotReloadClassLoader extends ClassLoader {
    private final Path classDir;

    public HotReloadClassLoader(Path classDir, ClassLoader parent) {
        super(parent);
        this.classDir = classDir;
    }

    @Override
    protected Class<?> findClass(String name) throws ClassNotFoundException {
        String classFile = name.replace('.', '/') + ".class";
        Path classPath = classDir.resolve(classFile);
        try {
            byte[] bytes = Files.readAllBytes(classPath);
            return defineClass(name, bytes, 0, bytes.length);
        } catch (IOException e) {
            throw new ClassNotFoundException(name, e);
        }
    }
}
// Used for: plugin systems, hot reloading, bytecode instrumentation
```

---

### Q4. What is JIT Compilation and escape analysis?

```java
// JIT: interpreter runs bytecode first, then hot methods are compiled to native code

// Tiered compilation:
// Level 0: Interpreted
// Level 1-3: C1 (client) compilation — fast compile, less optimization
// Level 4: C2 (server) compilation — aggressive optimizations

// ESCAPE ANALYSIS — JVM detects if object "escapes" current scope
// If NOT escaping → allocate on STACK (no GC needed!) or inline fields

// This object does NOT escape — JVM may stack-allocate or eliminate it
int calculateArea(int width, int height) {
    Point p = new Point(width, height);   // may be optimized away!
    return p.x * p.y;
}

// This DOES escape — must be heap-allocated
Point createPoint(int x, int y) {
    return new Point(x, y);   // returned to caller — escapes!
}

// JVM flags
// -XX:+DoEscapeAnalysis     (default on)
// -XX:+EliminateAllocations (default on) — scalar replacement
// -XX:+PrintEscapeAnalysis  — debug what JVM decided

// Check what JIT compiled
// -XX:+PrintCompilation — see which methods were compiled
// -XX:CompileThreshold=1000 — trigger JIT after 1000 invocations (default 10000)
```

---

### Q5. How do you diagnose a memory leak?

```java
// SYMPTOMS: Heap growing continuously, frequent Full GCs, OutOfMemoryError

// STEP 1: Enable GC logging and watch for increasing Old Gen
// -Xlog:gc*:file=gc.log:time,uptime

// STEP 2: Take heap dumps
// At OOM: -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/
// Manual: jmap -dump:format=b,file=heap.hprof <pid>
// Or: jcmd <pid> GC.heap_dump /tmp/heap.hprof

// STEP 3: Analyze with Eclipse MAT or VisualVM
// Look for: "Dominator tree", "Leak suspects report"

// Common leak patterns:
// 1. Static collections that grow forever
class LeakExample {
    private static final Map<String, Object> REGISTRY = new HashMap<>();  // ❌ never cleared
    public void register(String key, Object value) { REGISTRY.put(key, value); }
    // Fix: use WeakHashMap or clear on unregister
}

// 2. Listener / callback not deregistered
eventBus.subscribe(myListener);   // must call unsubscribe when done!

// 3. ThreadLocal not cleared in thread pool threads
ThreadLocal<Connection> connLocal = new ThreadLocal<>();
// After use in pool thread:
try { /* use connLocal */ } finally {
    connLocal.remove();   // ← MUST DO THIS or connection leaks forever in pool thread
}

// 4. String.intern() overuse
String s = new String(largeBytesArray).intern();  // ❌ can fill Metaspace
```
