# Java Advanced — JVM Internals, GC & Design Patterns

> **Topics:** JVM memory model, Garbage Collection, ClassLoader, Reflection, Design Patterns (Singleton, Factory, Builder, Observer, Strategy, Proxy), Java Memory Model, happens-before

---

## 📋 Reading Progress

- [ ] **Section 1:** JVM Memory & Class Loading (Q1–Q15)
- [ ] **Section 2:** Garbage Collection (Q16–Q25)
- [ ] **Section 3:** Java Memory Model & happens-before (Q26–Q35)
- [ ] **Section 4:** Design Patterns in Java (Q36–Q55)

> 🔖 **Last read:** <!-- -->

---

## Section 1: JVM Memory & Class Loading (Q1–Q15)

### 1. JVM Memory Areas
**Q: What is stored in each JVM memory area?**
```java
public class Main {
    static int staticVar = 10;     // Method Area (static fields)
    int instanceVar = 20;          // Heap (per object)

    public static void main(String[] args) {
        int localVar = 30;         // Stack (per method frame)
        Main obj = new Main();     // reference on Stack, object on Heap
        System.out.println(staticVar + " " + obj.instanceVar + " " + localVar);
    }
}
```
**A:** `10 20 30`.
| Area | Stores |
|---|---|
| **Method Area** | Class metadata, static fields, bytecode |
| **Heap** | Objects and instance fields |
| **Stack** | Method frames, local variables, operand stack |
| **PC Register** | Address of current instruction per thread |
| **Native Stack** | JNI native method frames |

---

### 2. StackOverflowError — Infinite Recursion
**Q: What happens?**
```java
public class Main {
    static void recurse() { recurse(); } // no base case!
    public static void main(String[] args) {
        try { recurse(); }
        catch (StackOverflowError e) { System.out.println("Stack overflow!"); }
    }
}
```
**A:** `Stack overflow!`. Each method call pushes a new frame onto the stack. With no base case, the stack is exhausted. `StackOverflowError` is thrown.

---

### 3. OutOfMemoryError — Heap Exhaustion
**Q: What happens?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<byte[]> data = new ArrayList<>();
        try {
            while (true) data.add(new byte[1024 * 1024]); // add 1MB chunks
        } catch (OutOfMemoryError e) {
            System.out.println("OOM: " + e.getMessage());
        }
    }
}
```
**A:** `OOM: Java heap space`. When the heap is full and GC cannot recover enough space, `OutOfMemoryError` is thrown.

---

### 4. Static Initializer Order
**Q: What is the output?**
```java
public class Main {
    static int a = initA();
    static int b = 20;

    static int initA() {
        System.out.println("initA called, b=" + b); // b is still 0 (default) here!
        return 10;
    }

    static { System.out.println("static block: a=" + a + ", b=" + b); }

    public static void main(String[] args) { System.out.println("main: a=" + a + " b=" + b); }
}
```
**A:**
```
initA called, b=0
static block: a=10, b=20
main: a=10 b=20
```
Static fields are initialized in declaration order. During `initA()`, `b` hasn't been assigned yet (default int value is 0).

---

### 5. Class Loading — When Does It Happen?
**Q: What is the output?**
```java
public class Main {
    static class Lazy {
        static { System.out.println("Lazy loaded!"); }
        static int VALUE = 42;
    }

    public static void main(String[] args) {
        System.out.println("before access");
        System.out.println(Lazy.VALUE); // triggers class loading
        System.out.println(Lazy.VALUE); // class already loaded
    }
}
```
**A:**
```
before access
Lazy loaded!
42
42
```
A class is loaded and initialized on **first active use**. Subsequent accesses use the already-loaded class.

---

### 6. ClassLoader Hierarchy
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        ClassLoader cl = Main.class.getClassLoader();
        System.out.println(cl);                // AppClassLoader
        System.out.println(cl.getParent());    // PlatformClassLoader (Java 9+)
        System.out.println(cl.getParent().getParent()); // BootstrapClassLoader (null in Java)
    }
}
```
**A:**
```
jdk.internal.loader.ClassLoaders$AppClassLoader@...
jdk.internal.loader.ClassLoaders$PlatformClassLoader@...
null
```
Delegation model: Bootstrap → Platform → Application. A class is loaded by the highest-level loader that can find it.

---

### 7. Reflection — Access Private Field
**Q: What is the output?**
```java
import java.lang.reflect.*;
public class Main {
    static class Secret { private String value = "hidden"; }

    public static void main(String[] args) throws Exception {
        Secret s = new Secret();
        Field f = Secret.class.getDeclaredField("value");
        f.setAccessible(true); // bypass access control
        System.out.println(f.get(s));
        f.set(s, "revealed");
        System.out.println(f.get(s));
    }
}
```
**A:**
```
hidden
revealed
```
Reflection breaks encapsulation. Use carefully — it's slower than direct access and bypasses compile-time safety.

---

### 8. Reflection — Invoke Method
**Q: What is the output?**
```java
import java.lang.reflect.*;
public class Main {
    static String greet(String name) { return "Hello, " + name + "!"; }

    public static void main(String[] args) throws Exception {
        Method m = Main.class.getDeclaredMethod("greet", String.class);
        String result = (String) m.invoke(null, "World"); // null for static method
        System.out.println(result);
    }
}
```
**A:** `Hello, World!`

---

### 9. instanceof and getClass() Difference
**Q: What is the output?**
```java
public class Main {
    static class Animal {}
    static class Dog extends Animal {}

    public static void main(String[] args) {
        Animal a = new Dog();
        System.out.println(a instanceof Animal); // true — Dog IS-A Animal
        System.out.println(a instanceof Dog);    // true — actual type is Dog
        System.out.println(a.getClass() == Animal.class); // false — exact type check
        System.out.println(a.getClass() == Dog.class);    // true
    }
}
```
**A:**
```
true
true
false
true
```
`instanceof` checks the type hierarchy. `getClass()` returns the exact runtime type.

---

### 10. Finalize and Object.finalize() Pitfall
**Q: What is the concept?**
```java
public class Main {
    // DEPRECATED since Java 9. DO NOT USE.
    @Override
    protected void finalize() throws Throwable {
        System.out.println("finalizing"); // called by GC before collection
        super.finalize();
    }
    public static void main(String[] args) throws InterruptedException {
        new Main();
        System.gc();          // suggest GC, not guaranteed
        Thread.sleep(100);
        System.out.println("done");
    }
}
```
**A:** `finalizing` may or may not print. Finalization is **not guaranteed** to run. Use `java.lang.ref.Cleaner` (Java 9+) or try-with-resources instead.

---

### 11. String in Switch — Since Java 7
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String day = "MONDAY";
        switch (day) {
            case "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY" -> System.out.println("Weekday");
            case "SATURDAY", "SUNDAY" -> System.out.println("Weekend");
            default -> System.out.println("Unknown");
        }
    }
}
```
**A:** `Weekday`. Switch with Strings uses `hashCode()` + `equals()` internally. Arrow-case switch (Java 14+) is cleaner.

---

### 12. Sealed Classes (Java 17+)
**Q: Does this compile?**
```java
public class Main {
    sealed interface Shape permits Circle, Rectangle {}
    record Circle(double radius) implements Shape {}
    record Rectangle(double w, double h) implements Shape {}
    // class Triangle implements Shape {} // compile error — not permitted!

    static double area(Shape s) {
        return switch (s) {
            case Circle c     -> Math.PI * c.radius() * c.radius();
            case Rectangle r  -> r.w() * r.h();
        }; // no default needed — exhaustive!
    }

    public static void main(String[] args) {
        System.out.printf("%.2f%n", area(new Circle(5)));
        System.out.printf("%.2f%n", area(new Rectangle(3, 4)));
    }
}
```
**A:**
```
78.54
12.00
```
Sealed classes restrict which types can implement them, enabling exhaustive pattern matching.

---

### 13. Records vs Regular Classes
**Q: What is the output?**
```java
public class Main {
    record Point(int x, int y) {
        // compact constructor
        Point { if (x < 0 || y < 0) throw new IllegalArgumentException("negative!"); }
    }

    public static void main(String[] args) {
        Point p1 = new Point(3, 4);
        Point p2 = new Point(3, 4);
        System.out.println(p1.x());           // accessor
        System.out.println(p1.equals(p2));    // value equality
        System.out.println(p1.hashCode() == p2.hashCode());
        System.out.println(p1);               // toString
    }
}
```
**A:**
```
3
true
true
Point[x=3, y=4]
```
Records auto-generate: `equals()`, `hashCode()`, `toString()`, and accessors. They are immutable by design.

---

### 14. instanceof Pattern Matching (Java 16+)
**Q: What is the output?**
```java
public class Main {
    sealed interface Shape permits Circle, Rect {}
    record Circle(double r) implements Shape {}
    record Rect(double w, double h) implements Shape {}

    static String describe(Object obj) {
        if (obj instanceof Circle c && c.r() > 5) return "big circle r=" + c.r();
        if (obj instanceof Rect r) return "rect " + r.w() + "x" + r.h();
        return "other";
    }

    public static void main(String[] args) {
        System.out.println(describe(new Circle(10)));
        System.out.println(describe(new Circle(2)));
        System.out.println(describe(new Rect(3, 4)));
    }
}
```
**A:**
```
big circle r=10.0
other
rect 3.0x4.0
```

---

### 15. VarHandle — Low-Level Memory Access (Java 9+)
**Q: What is the output?**
```java
import java.lang.invoke.*;
public class Main {
    static class Counter { volatile int value; }
    public static void main(String[] args) throws Exception {
        VarHandle vh = MethodHandles.lookup()
                .in(Counter.class)
                .findVarHandle(Counter.class, "value", int.class);
        Counter c = new Counter();
        vh.set(c, 10);
        System.out.println(vh.get(c));
        boolean swapped = vh.compareAndSet(c, 10, 20); // CAS
        System.out.println(swapped + " " + c.value);
    }
}
```
**A:**
```
10
true 20
```
`VarHandle` provides low-level atomic access to fields — more powerful than `AtomicXxx` classes.

---

## Section 2: Garbage Collection (Q16–Q25)

### 16. GC Roots — What Keeps Objects Alive
**Q: Which objects will survive GC?**
```java
public class Main {
    static Object globalRef; // GC root — static field

    public static void main(String[] args) {
        Object a = new Object(); // reachable via local variable (GC root)
        Object b = new Object(); // will survive — referenced by 'a' chain
        a = null;                // 'a' is no longer reachable
        // b is still referenced by its local variable
        System.gc();
        System.out.println("GC roots: static fields, local vars, active threads, JNI refs");
    }
}
```
**A:** Prints the message. Objects reachable from GC roots are never collected. `a`'s object may be collected after `a = null`.

---

### 17. Object Generations in HotSpot GC
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println("HotSpot Heap Generations:");
        System.out.println("Young Gen (Eden + S0 + S1) — new objects, minor GC");
        System.out.println("Old Gen (Tenured) — long-lived objects, major GC");
        System.out.println();
        System.out.println("GC algorithms:");
        System.out.println("G1GC (Java 9+ default) — region-based, predictable pause");
        System.out.println("ZGC (Java 15+ prod) — sub-millisecond pauses");
        System.out.println("Shenandoah — concurrent compaction");
    }
}
```
**A:** Prints GC info. For interviews: G1GC is the default. Minor GC = Young gen. Major/Full GC = whole heap.

---

### 18. WeakReference — GC Eligible
**Q: What is the output?**
```java
import java.lang.ref.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        Object obj = new Object();
        WeakReference<Object> weakRef = new WeakReference<>(obj);

        System.out.println("before GC: " + (weakRef.get() != null));
        obj = null; // strong reference removed
        System.gc();
        Thread.sleep(100);
        System.out.println("after GC: " + (weakRef.get() != null)); // likely null
    }
}
```
**A:**
```
before GC: true
after GC: false
```
A `WeakReference` doesn't prevent GC. The object can be collected when only weakly reachable.

---

### 19. SoftReference — Last Resort Before OOM
**Q: What is the concept?**
```java
import java.lang.ref.*;
public class Main {
    public static void main(String[] args) {
        SoftReference<byte[]> softRef = new SoftReference<>(new byte[1024]);
        // SoftReference: collected only when JVM is about to throw OOM
        // WeakReference: collected at next GC
        // PhantomReference: collected after finalization (for cleanup)
        System.out.println("soft ref alive: " + (softRef.get() != null));
    }
}
```
**A:** `soft ref alive: true`. `SoftReference` is ideal for caches — JVM holds them as long as memory is available.

---

### 20. ReferenceQueue — Cleanup Notification
**Q: What is the output?**
```java
import java.lang.ref.*;
public class Main {
    public static void main(String[] args) throws InterruptedException {
        ReferenceQueue<Object> queue = new ReferenceQueue<>();
        WeakReference<Object> ref = new WeakReference<>(new Object(), queue);

        System.gc(); Thread.sleep(100);

        Reference<?> cleared = queue.poll();
        System.out.println("enqueued: " + (cleared != null));
    }
}
```
**A:** `enqueued: true`. After GC collects the referent, the `Reference` object is enqueued in the `ReferenceQueue` for cleanup processing.

---

### 21. Memory Leak Patterns
**Q: Which of these is a memory leak?**
```java
import java.util.*;
public class Main {
    // LEAK 1: Static collection grows unbounded
    static List<byte[]> cache = new ArrayList<>();

    // LEAK 2: Event listener never removed
    // LEAK 3: ThreadLocal not removed in thread pools
    // LEAK 4: Overriding equals/hashCode incorrectly → HashMap never finds/removes entries

    public static void main(String[] args) {
        // Demonstrates LEAK 1
        for (int i = 0; i < 100; i++) {
            cache.add(new byte[1024]); // keeps growing, never cleared
        }
        System.out.println("cache size: " + cache.size());
    }
}
```
**A:** `cache size: 100`. Classic memory leak patterns in Java to know for interviews.

---

### 22. Premature Promotion — GC Tuning Concern
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        // Long-lived objects promotion to Old Gen happens when:
        // 1. Object survives -XX:MaxTenuringThreshold GC cycles (default 15)
        // 2. Survivor space is full (early promotion)
        // 3. Object is too large for Eden (allocated directly in Old Gen)

        // Key JVM flags:
        // -Xms/-Xmx: initial/max heap
        // -XX:NewRatio: ratio of Young:Old gen size
        // -XX:+UseG1GC: use G1 GC
        // -XX:MaxGCPauseMillis: target pause time (G1)
        System.out.println("GC tuning concepts printed");
    }
}
```
**A:** `GC tuning concepts printed`. Know these flags for senior Java interviews.

---

### 23. System.gc() vs Runtime.gc()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.gc();         // hint to JVM — not guaranteed
        Runtime.getRuntime().gc(); // same as System.gc()

        // Can be disabled with -XX:+DisableExplicitGC JVM flag
        // Calling gc() in production is an anti-pattern
        System.out.println("GC suggested (not guaranteed)");
    }
}
```
**A:** `GC suggested (not guaranteed)`. Never rely on explicit GC in production code.

---

### 24. Off-Heap Memory — ByteBuffer.allocateDirect
**Q: What is the output?**
```java
import java.nio.*;
public class Main {
    public static void main(String[] args) {
        // Heap buffer — managed by GC
        ByteBuffer heap = ByteBuffer.allocate(1024);
        System.out.println("Heap buffer: " + heap.isDirect());

        // Direct buffer — off-heap, NOT managed by GC
        ByteBuffer direct = ByteBuffer.allocateDirect(1024);
        System.out.println("Direct buffer: " + direct.isDirect());
        // Useful for I/O — avoids copying between Java heap and OS buffers
    }
}
```
**A:**
```
Heap buffer: false
Direct buffer: true
```
Direct buffers are faster for I/O but harder to manage (not GC'd directly — cleaned via `Cleaner`/`ReferenceQueue`).

---

### 25. Profiling — jmap, jstack, jcmd basics
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        long pid = ProcessHandle.current().pid();
        System.out.println("PID: " + pid);
        System.out.println();
        System.out.println("Tools for profiling:");
        System.out.println("jmap -heap <pid>    — heap summary");
        System.out.println("jmap -histo <pid>   — object histogram");
        System.out.println("jstack <pid>        — thread dump");
        System.out.println("jcmd <pid> GC.run   — trigger GC");
        System.out.println("jcmd <pid> VM.flags — JVM flags in use");
    }
}
```
**A:** Prints current PID and tool descriptions. Essential tools for production Java debugging.

---

## Section 3: Java Memory Model & happens-before (Q26–Q35)

### 26. happens-before — The Foundation
**Q: What is the output?**
```java
public class Main {
    static int x = 0;
    static boolean flag = false;

    public static void main(String[] args) throws InterruptedException {
        Thread writer = new Thread(() -> {
            x = 42;
            flag = true; // may be reordered by CPU or compiler!
        });
        Thread reader = new Thread(() -> {
            while (!flag) {} // busy-wait
            System.out.println(x); // may print 0 if reordering occurs!
        });
        writer.start(); reader.start();
        writer.join(); reader.join();
    }
}
```
**A:** Might print `0` due to instruction reordering! The write to `flag=true` might become visible before `x=42`. **Fix:** use `volatile` on both, or `synchronized`.

---

### 27. happens-before — volatile guarantees ordering
**Q: What is the output?**
```java
public class Main {
    static int x = 0;
    static volatile boolean flag = false; // volatile flag

    public static void main(String[] args) throws InterruptedException {
        Thread writer = new Thread(() -> { x = 42; flag = true; });
        Thread reader = new Thread(() -> {
            while (!flag) {} // wait for volatile flag
            System.out.println(x); // guaranteed to see x=42
        });
        writer.start(); reader.start();
        writer.join(); reader.join();
    }
}
```
**A:** `42`. Volatile write to `flag` **happens-before** any subsequent read of `flag`. All writes before the volatile write are visible after the volatile read.

---

### 28. happens-before Rules Summary
**Q: What guarantees order of operations?**
```java
public class Main {
    public static void main(String[] args) throws InterruptedException {
        // 1. Program order — within a single thread
        int a = 1; int b = a + 1; System.out.println(b);

        // 2. Monitor lock — unlock happens-before subsequent lock
        // 3. volatile write happens-before volatile read
        // 4. Thread.start() happens-before any action in the new thread
        Thread t = new Thread(() -> System.out.println("thread ran"));
        t.start();
        // 5. All actions in thread happen-before Thread.join() returns
        t.join();
        System.out.println("join complete");
    }
}
```
**A:** `2`, then `thread ran`, then `join complete`. These rules define what is safe to share across threads without explicit synchronization.

---

### 29. Instruction Reordering — CPU Effect
**Q: What is the bug?**
```java
public class Main {
    // Both fields can appear to be written out of order to other CPUs
    static int value = 0;
    static boolean initialized = false;

    // Thread A
    static void write() {
        value = 42;           // (1)
        initialized = true;   // (2) -- CPU may reorder (2) before (1)!
    }

    // Thread B
    static void read() {
        if (initialized) System.out.println(value); // might see 0 if reordered!
    }
    public static void main(String[] args) { System.out.println("Reordering is a real hazard!"); }
}
```
**A:** `Reordering is a real hazard!`. Without synchronization or `volatile`, the CPU/JIT can reorder writes. This is a subtle but critical concurrency bug.

---

### 30. Safe Publication Patterns
**Q: Which is safely published?**
```java
public class Main {
    // UNSAFE: final field set after construction
    static class Unsafe { int value; Unsafe() {} }

    // SAFE: All fields final — values visible to all threads after construction
    static class Safe { final int value; Safe(int v) { this.value = v; } }

    // SAFE: volatile reference
    static volatile Safe sharedSafe;

    // SAFE: synchronized publication
    private static Safe lazyInit;
    static synchronized Safe get() {
        if (lazyInit == null) lazyInit = new Safe(42);
        return lazyInit;
    }

    public static void main(String[] args) {
        sharedSafe = new Safe(42); // safely published via volatile field
        System.out.println(sharedSafe.value);
    }
}
```
**A:** `42`. Safe publication methods: `final` fields, `volatile` references, `synchronized` access, or static initializers.

---

### 31. Memory Visibility — Simple Demo
**Q: What is the race?**
```java
public class Main {
    boolean stop = false;         // not volatile!

    void run() throws InterruptedException {
        Thread t = new Thread(() -> {
            int i = 0;
            while (!stop) i++;   // may loop forever — cached value of stop
            System.out.println("Stopped after " + i + " iterations");
        });
        t.start();
        Thread.sleep(100);
        stop = true;             // write may not be visible to t without volatile!
        t.join(1000);
    }
    public static void main(String[] args) throws InterruptedException { new Main().run(); }
}
```
**A:** May never stop. The thread caches `stop` in a register and never reads from main memory. Declare `volatile boolean stop` to fix.

---

### 32. AtomicStampedReference — ABA Problem
**Q: What is the ABA problem?**
```java
import java.util.concurrent.atomic.*;
public class Main {
    public static void main(String[] args) {
        // ABA problem: value changes A→B→A
        // CAS sees "A" and thinks nothing changed, but state DID change!
        // Solution: AtomicStampedReference adds a version stamp

        AtomicStampedReference<String> ref = new AtomicStampedReference<>("A", 0);
        int[] stampHolder = new int[1];

        String val = ref.get(stampHolder);
        int stamp = stampHolder[0];

        // Thread simulates A → B → A
        ref.compareAndSet("A", "B", 0, 1);
        ref.compareAndSet("B", "A", 1, 2);

        // Our CAS fails because stamp doesn't match
        boolean success = ref.compareAndSet("A", "C", stamp, stamp + 1);
        System.out.println("CAS succeeded: " + success); // false — ABA detected!
        System.out.println("value: " + ref.getReference()); // still "A"
    }
}
```
**A:**
```
CAS succeeded: false
value: A
```

---

### 33. Happens-Before: Thread.start()
**Q: What is the output?**
```java
public class Main {
    static int x = 0;
    public static void main(String[] args) throws InterruptedException {
        x = 42; // happens-before the thread's execution
        Thread t = new Thread(() -> System.out.println("x = " + x));
        t.start(); // start() establishes happens-before
        t.join();
    }
}
```
**A:** `x = 42`. Actions before `Thread.start()` happen-before any action in the newly started thread. This is a guaranteed JMM rule.

---

### 34. Happens-Before: Thread.join()
**Q: What is the output?**
```java
public class Main {
    static int result = 0;
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> result = 99);
        t.start();
        t.join(); // all of t's actions happen-before join() returns
        System.out.println(result); // safe to read
    }
}
```
**A:** `99`. `Thread.join()` guarantees all writes in the joined thread are visible after `join()` returns.

---

### 35. final Fields — Initialization Safety
**Q: What is the output?**
```java
public class Main {
    static class Immutable {
        final int x;
        final int y;
        Immutable(int x, int y) { this.x = x; this.y = y; }
    }
    static Immutable ref; // non-volatile!

    public static void main(String[] args) {
        // Final fields: guaranteed visible to all threads AFTER constructor completes
        // Even without synchronization or volatile
        ref = new Immutable(1, 2);
        System.out.println(ref.x + " " + ref.y); // always 1 2
    }
}
```
**A:** `1 2`. `final` fields initialized in constructors have a special JMM guarantee — they are safely visible to all threads after the constructor completes.

---

## Section 4: Design Patterns in Java (Q36–Q55)

### 36. Singleton — Enum Pattern (Best)
**Q: What is the output?**
```java
public class Main {
    enum Config {
        INSTANCE;
        private String dbUrl = "jdbc:mysql://localhost/db";
        String getDbUrl() { return dbUrl; }
    }

    public static void main(String[] args) {
        System.out.println(Config.INSTANCE.getDbUrl());
        System.out.println(Config.INSTANCE == Config.INSTANCE); // true — guaranteed unique
    }
}
```
**A:**
```
jdbc:mysql://localhost/db
true
```
Enum Singleton is thread-safe, serialization-safe, and handles reflection attacks automatically. The recommended approach.

---

### 37. Factory Method Pattern
**Q: What is the output?**
```java
public class Main {
    interface Animal { String speak(); }
    static class Dog implements Animal { public String speak() { return "Woof!"; } }
    static class Cat implements Animal { public String speak() { return "Meow!"; } }

    static Animal create(String type) {
        return switch (type) {
            case "dog" -> new Dog();
            case "cat" -> new Cat();
            default    -> throw new IllegalArgumentException("Unknown: " + type);
        };
    }

    public static void main(String[] args) {
        System.out.println(create("dog").speak());
        System.out.println(create("cat").speak());
    }
}
```
**A:**
```
Woof!
Meow!
```

---

### 38. Builder Pattern
**Q: What is the output?**
```java
public class Main {
    static class Pizza {
        String size; boolean cheese, pepperoni, mushroom;
        private Pizza(Builder b) { size=b.size; cheese=b.cheese; pepperoni=b.pepperoni; mushroom=b.mushroom; }
        public String toString() { return size + " pizza: cheese=" + cheese + " pep=" + pepperoni + " mush=" + mushroom; }

        static class Builder {
            String size; boolean cheese, pepperoni, mushroom;
            Builder(String size) { this.size = size; }
            Builder cheese()    { cheese = true;    return this; }
            Builder pepperoni() { pepperoni = true;  return this; }
            Builder mushroom()  { mushroom = true;   return this; }
            Pizza build()       { return new Pizza(this); }
        }
    }

    public static void main(String[] args) {
        Pizza pizza = new Pizza.Builder("large").cheese().pepperoni().build();
        System.out.println(pizza);
    }
}
```
**A:** `large pizza: cheese=true pep=true mush=false`

---

### 39. Observer Pattern
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    interface Observer { void update(String event); }
    static class EventBus {
        List<Observer> observers = new ArrayList<>();
        void subscribe(Observer o) { observers.add(o); }
        void publish(String event) { observers.forEach(o -> o.update(event)); }
    }

    public static void main(String[] args) {
        EventBus bus = new EventBus();
        bus.subscribe(e -> System.out.println("Logger: " + e));
        bus.subscribe(e -> System.out.println("Emailer: " + e));
        bus.publish("user.signup");
        bus.publish("order.placed");
    }
}
```
**A:**
```
Logger: user.signup
Emailer: user.signup
Logger: order.placed
Emailer: order.placed
```

---

### 40. Strategy Pattern
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    @FunctionalInterface
    interface SortStrategy { void sort(int[] arr); }

    static void performSort(int[] arr, SortStrategy strategy) { strategy.sort(arr); }

    public static void main(String[] args) {
        int[] arr = {5, 2, 8, 1, 9};
        // Strategies are just lambdas — functional interface
        performSort(arr, Arrays::sort);
        System.out.println(Arrays.toString(arr));
    }
}
```
**A:** `[1, 2, 5, 8, 9]`. In modern Java, Strategy is often just a `@FunctionalInterface` — strategies are passed as lambdas.

---

### 41. Decorator Pattern
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        // Functional decorator using Function.andThen
        Function<String, String> trim    = String::trim;
        Function<String, String> upper   = String::toUpperCase;
        Function<String, String> exclaim = s -> s + "!";

        Function<String, String> pipeline = trim.andThen(upper).andThen(exclaim);
        System.out.println(pipeline.apply("  hello world  "));
    }
}
```
**A:** `HELLO WORLD!`. Decorator in functional style: compose transformations with `andThen`.

---

### 42. Proxy Pattern — Lazy Loading
**Q: What is the output?**
```java
public class Main {
    interface DataService { String fetch(); }

    static class RealService implements DataService {
        RealService() { System.out.println("RealService created (expensive!)"); }
        public String fetch() { return "real data"; }
    }

    static class LazyProxy implements DataService {
        private DataService real;
        public String fetch() {
            if (real == null) real = new RealService(); // lazy init
            return real.fetch();
        }
    }

    public static void main(String[] args) {
        DataService proxy = new LazyProxy();
        System.out.println("proxy created");
        System.out.println(proxy.fetch()); // triggers creation
        System.out.println(proxy.fetch()); // reuses existing
    }
}
```
**A:**
```
proxy created
RealService created (expensive!)
real data
real data
```

---

### 43. Template Method Pattern
**Q: What is the output?**
```java
public class Main {
    abstract static class Report {
        final void generate() { // template method — fixed skeleton
            fetchData();
            processData();
            exportData();
        }
        abstract void fetchData();
        abstract void processData();
        void exportData() { System.out.println("Exporting to default format"); }
    }

    static class CsvReport extends Report {
        void fetchData()    { System.out.println("Fetching from CSV"); }
        void processData()  { System.out.println("Processing CSV rows"); }
        @Override void exportData() { System.out.println("Saving as CSV"); }
    }

    public static void main(String[] args) { new CsvReport().generate(); }
}
```
**A:**
```
Fetching from CSV
Processing CSV rows
Saving as CSV
```

---

### 44. Command Pattern
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    @FunctionalInterface interface Command { void execute(); }

    static class CommandHistory {
        Deque<Command> history = new ArrayDeque<>();
        void run(Command cmd) { cmd.execute(); history.push(cmd); }
        void size() { System.out.println("history size: " + history.size()); }
    }

    public static void main(String[] args) {
        CommandHistory ch = new CommandHistory();
        ch.run(() -> System.out.println("cmd1"));
        ch.run(() -> System.out.println("cmd2"));
        ch.size();
    }
}
```
**A:**
```
cmd1
cmd2
history size: 2
```

---

### 45. Chain of Responsibility Pattern
**Q: What is the output?**
```java
import java.util.function.*;
public class Main {
    public static void main(String[] args) {
        // Composed as a pipeline of predicates/functions
        UnaryOperator<String> authenticate = s -> {
            System.out.println("authenticating: " + s);
            return s;
        };
        UnaryOperator<String> authorize = s -> {
            System.out.println("authorizing: " + s);
            return s;
        };
        UnaryOperator<String> log = s -> {
            System.out.println("logging: " + s);
            return s;
        };

        UnaryOperator<String> chain = authenticate.andThen(authorize).andThen(log);
        chain.apply("request");
    }
}
```
**A:**
```
authenticating: request
authorizing: request
logging: request
```

---

### 46. Adapter Pattern
**Q: What is the output?**
```java
public class Main {
    interface ModernLogger { void log(String message, Level level); }
    enum Level { INFO, ERROR }

    // Legacy system with different interface
    static class LegacyLogger {
        void info(String msg) { System.out.println("INFO: " + msg); }
        void error(String msg) { System.out.println("ERROR: " + msg); }
    }

    // Adapter wraps legacy and implements modern interface
    static class LoggerAdapter implements ModernLogger {
        private final LegacyLogger legacy;
        LoggerAdapter(LegacyLogger l) { this.legacy = l; }
        public void log(String msg, Level level) {
            if (level == Level.INFO) legacy.info(msg);
            else legacy.error(msg);
        }
    }

    public static void main(String[] args) {
        ModernLogger logger = new LoggerAdapter(new LegacyLogger());
        logger.log("server started", Level.INFO);
        logger.log("disk full", Level.ERROR);
    }
}
```
**A:**
```
INFO: server started
ERROR: disk full
```

---

### 47. Flyweight Pattern
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class CharacterPool {
        private static final Map<Character, String> pool = new HashMap<>();
        static String get(char c) {
            return pool.computeIfAbsent(c, k -> "char:" + k);
        }
    }

    public static void main(String[] args) {
        String a = CharacterPool.get('A');
        String a2 = CharacterPool.get('A'); // same shared instance
        System.out.println(a == a2);  // true — same object from pool
        System.out.println(CharacterPool.get('B'));
    }
}
```
**A:**
```
true
char:B
```
Flyweight shares common state to reduce memory. Java String pool is a built-in flyweight.

---

### 48. State Pattern
**Q: What is the output?**
```java
public class Main {
    interface TrafficLight { String color(); TrafficLight next(); }

    enum Light implements TrafficLight {
        RED   { public String color() { return "RED";   } public TrafficLight next() { return GREEN;  } },
        GREEN { public String color() { return "GREEN"; } public TrafficLight next() { return YELLOW; } },
        YELLOW{ public String color() { return "YELLOW";} public TrafficLight next() { return RED;    } };
    }

    public static void main(String[] args) {
        TrafficLight light = Light.RED;
        for (int i = 0; i < 6; i++) {
            System.out.print(light.color() + " ");
            light = light.next();
        }
    }
}
```
**A:** `RED GREEN YELLOW RED GREEN YELLOW `

---

### 49. Dependency Injection — Manual DI
**Q: What is the output?**
```java
public class Main {
    interface NotificationService { void send(String msg); }
    static class EmailService implements NotificationService {
        public void send(String msg) { System.out.println("Email: " + msg); }
    }

    static class UserService {
        private final NotificationService notif; // injected dependency
        UserService(NotificationService notif) { this.notif = notif; }
        void register(String user) { notif.send("Welcome " + user); }
    }

    public static void main(String[] args) {
        // Wire dependencies manually (Spring does this automatically)
        UserService service = new UserService(new EmailService());
        service.register("Alice");
    }
}
```
**A:** `Email: Welcome Alice`. DI separates object creation from use — enables easy swapping implementations in tests.

---

### 50. LRU Cache — LinkedHashMap Implementation
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class LRUCache<K, V> extends LinkedHashMap<K, V> {
        private final int capacity;
        LRUCache(int capacity) {
            super(capacity, 0.75f, true); // accessOrder=true
            this.capacity = capacity;
        }
        @Override
        protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
            return size() > capacity;
        }
    }

    public static void main(String[] args) {
        LRUCache<Integer, String> cache = new LRUCache<>(3);
        cache.put(1, "one"); cache.put(2, "two"); cache.put(3, "three");
        cache.get(1); // access 1 — becomes most recent
        cache.put(4, "four"); // evicts LRU: 2
        System.out.println(cache.keySet()); // [3, 1, 4] — 2 evicted
    }
}
```
**A:** `[3, 1, 4]`. Classic LRU Cache using `LinkedHashMap` with `accessOrder=true`.

---

### 51. Circuit Breaker Pattern (Microservices)
**Q: What is the output?**
```java
import java.util.concurrent.atomic.*;
public class Main {
    enum State { CLOSED, OPEN, HALF_OPEN }

    static class CircuitBreaker {
        State state = State.CLOSED;
        AtomicInteger failures = new AtomicInteger(0);
        final int threshold = 3;

        String call(String service) {
            if (state == State.OPEN) return "Circuit OPEN — fast fail";
            try {
                if (service.equals("down")) throw new RuntimeException("service down");
                failures.set(0);
                return "success";
            } catch (Exception e) {
                if (failures.incrementAndGet() >= threshold) state = State.OPEN;
                return "error: " + e.getMessage();
            }
        }
    }

    public static void main(String[] args) {
        CircuitBreaker cb = new CircuitBreaker();
        System.out.println(cb.call("up"));
        System.out.println(cb.call("down"));
        System.out.println(cb.call("down"));
        System.out.println(cb.call("down")); // threshold reached — opens
        System.out.println(cb.call("up"));   // circuit is OPEN — fast fail
    }
}
```
**A:**
```
success
error: service down
error: service down
error: service down
Circuit OPEN — fast fail
```

---

### 52. Iterator Pattern — Custom
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class NumberRange implements Iterable<Integer> {
        private final int start, end, step;
        NumberRange(int start, int end, int step) { this.start=start; this.end=end; this.step=step; }

        public Iterator<Integer> iterator() {
            return new Iterator<>() {
                int current = start;
                public boolean hasNext() { return current <= end; }
                public Integer next() { int val = current; current += step; return val; }
            };
        }
    }

    public static void main(String[] args) {
        for (int n : new NumberRange(1, 10, 2)) System.out.print(n + " ");
    }
}
```
**A:** `1 3 5 7 9 `

---

### 53. Null Object Pattern
**Q: What is the output?**
```java
public class Main {
    interface Logger { void log(String msg); }
    static class ConsoleLogger implements Logger { public void log(String msg) { System.out.println(msg); } }
    static class NullLogger implements Logger { public void log(String msg) {} } // does nothing

    static Logger getLogger(boolean enabled) {
        return enabled ? new ConsoleLogger() : new NullLogger();
    }

    public static void main(String[] args) {
        Logger active = getLogger(true);
        Logger silent = getLogger(false);
        active.log("something happened");
        silent.log("this is silenced"); // no NPE, no null check needed
        System.out.println("done");
    }
}
```
**A:**
```
something happened
done
```

---

### 54. Event Sourcing Pattern (concept)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    sealed interface Event permits Deposited, Withdrawn {}
    record Deposited(double amount) implements Event {}
    record Withdrawn(double amount) implements Event {}

    static double replay(List<Event> events) {
        double balance = 0;
        for (Event e : events) {
            balance += switch (e) {
                case Deposited d -> d.amount();
                case Withdrawn w -> -w.amount();
            };
        }
        return balance;
    }

    public static void main(String[] args) {
        List<Event> history = List.of(
            new Deposited(1000), new Withdrawn(200), new Deposited(500), new Withdrawn(100));
        System.out.println("Balance: " + replay(history));
    }
}
```
**A:** `Balance: 1200.0`. Event Sourcing stores state as a sequence of events — replaying them gives current state.

---

### 55. Specification Pattern — Composable Business Rules
**Q: What is the output?**
```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;
public class Main {
    record Product(String name, double price, boolean inStock) {}

    public static void main(String[] args) {
        List<Product> catalog = List.of(
            new Product("Apple",    0.5,  true),
            new Product("Laptop",   999.0, false),
            new Product("Keyboard", 49.0, true),
            new Product("Monitor",  299.0, true)
        );

        Predicate<Product> affordable   = p -> p.price() < 100;
        Predicate<Product> available    = Product::inStock;
        Predicate<Product> goodDeal     = affordable.and(available);

        List<String> results = catalog.stream()
                .filter(goodDeal)
                .map(Product::name)
                .collect(Collectors.toList());
        System.out.println(results);
    }
}
```
**A:** `[Apple, Keyboard]`. Composing `Predicate`s is the functional equivalent of the Specification pattern.
