# Virtusa Java Interview – Part 1: Core Java (Q&A with Code)

---

## Q1. What is the difference between ArrayList and LinkedList?

**Spoken Answer:**
"Both implement the List interface, but they differ in internal structure.

ArrayList uses a **dynamic array**. Random access by index is O(1) because it directly jumps to the memory position. But insertions/deletions in the middle are slow — O(n) — because elements must be shifted.

LinkedList uses a **doubly linked list**. Adding/removing at the head or tail is O(1) — just a pointer change. But random access is O(n) because you traverse node by node.

In practice I use ArrayList for most cases — CPUs cache arrays better. I switch to LinkedList only for frequent head/tail insertions."

```java
import java.util.*;

public class ArrayListVsLinkedList {
    public static void main(String[] args) {

        // ArrayList - fast random access O(1)
        List<String> arrayList = new ArrayList<>();
        arrayList.add("Java"); arrayList.add("Spring"); arrayList.add("Kafka");
        System.out.println("ArrayList get(1): " + arrayList.get(1)); // O(1) = Spring

        // LinkedList - fast head/tail insert O(1)
        LinkedList<String> linkedList = new LinkedList<>();
        linkedList.addFirst("First");
        linkedList.addLast("Last");
        linkedList.add(1, "Middle");
        System.out.println("LinkedList: " + linkedList); // [First, Middle, Last]

        // Performance demo
        long start = System.nanoTime();
        List<String> al = new ArrayList<>();
        for (int i = 0; i < 10_000; i++) al.add(0, "x"); // slow - O(n) shift
        System.out.println("ArrayList insert at 0: " + (System.nanoTime() - start) + "ns");

        start = System.nanoTime();
        LinkedList<String> ll = new LinkedList<>();
        for (int i = 0; i < 10_000; i++) ll.addFirst("x"); // fast - O(1)
        System.out.println("LinkedList addFirst:   " + (System.nanoTime() - start) + "ns");
    }
}
```

---

## Q2. What is the difference between Overloading and Overriding? What should be the access specifier of the overriding method?

**Spoken Answer:**
"Overloading is **compile-time polymorphism** — same method name, different parameters, within the same class. The compiler decides which version to call.

Overriding is **runtime polymorphism** — a subclass provides its own implementation of a parent class method. The JVM decides at runtime using dynamic dispatch.

For the access specifier rule: the overriding method must be **same or wider** than the parent. If the parent has `protected`, the child can use `protected` or `public`, but NOT `private` — that would narrow access and the compiler rejects it."

```java
public class OverloadVsOverride {

    // ─── OVERLOADING: same class, different parameters ───
    static int add(int a, int b)          { return a + b; }       // ① int
    static double add(double a, double b) { return a + b; }       // ② double
    static String add(String a, String b) { return a + b; }       // ③ String

    // ─── OVERRIDING: subclass redefines parent method ───
    static class Animal {
        protected String sound() { return "Some generic sound"; }
    }

    static class Dog extends Animal {
        @Override
        public String sound() {   // wider: protected → public ✅
            return "Woof!";
        }
        // private String sound() {}  // COMPILE ERROR: narrower ❌
    }

    static class Cat extends Animal {
        @Override
        protected String sound() {  // same specifier ✅
            return "Meow!";
        }
    }

    public static void main(String[] args) {
        // Overloading — resolved at COMPILE time
        System.out.println(add(1, 2));           // 3
        System.out.println(add(1.5, 2.5));       // 4.0
        System.out.println(add("Hello", " Java")); // Hello Java

        // Overriding — resolved at RUNTIME (dynamic dispatch)
        Animal a1 = new Dog();
        Animal a2 = new Cat();
        System.out.println(a1.sound()); // Woof!  ← Dog's version
        System.out.println(a2.sound()); // Meow!  ← Cat's version
    }
}
```

---

## Q3. What are the differences between Optional.of(), Optional.ofNullable(), and Optional.empty()?

**Spoken Answer:**
"Optional was introduced in Java 8 to represent a value that may or may not be present, and to eliminate NullPointerExceptions gracefully.

- `Optional.of(value)` — wraps a **guaranteed non-null** value. If you pass null it throws NPE immediately. Use when you're 100% certain.
- `Optional.ofNullable(value)` — the safe version. If value is null it returns an empty Optional. Use when the value might be null.
- `Optional.empty()` — an explicitly empty Optional. Use as a return value from methods to signal 'nothing found'."

```java
import java.util.Optional;

public class OptionalDemo {

    // Repository-style: return Optional to force caller to handle absence
    static Optional<String> findById(int id) {
        if (id == 1) return Optional.of("Alice");         // definitely present
        if (id == 2) return Optional.ofNullable(null);    // might be null
        return Optional.empty();                           // not found
    }

    public static void main(String[] args) {

        // Optional.of → throws NPE on null
        try {
            Optional<String> bad = Optional.of(null);
        } catch (NullPointerException e) {
            System.out.println("Optional.of(null) → NullPointerException");
        }

        // Optional.ofNullable → returns empty, no exception
        Optional<String> safe = Optional.ofNullable(null);
        System.out.println("isPresent: " + safe.isPresent());   // false
        System.out.println("orElse:    " + safe.orElse("default")); // default

        // Optional.empty → explicit empty
        Optional<String> empty = Optional.empty();
        System.out.println("isEmpty:   " + empty.isEmpty());    // true (Java 11+)

        // Practical chaining pattern
        findById(1)
            .filter(name -> name.startsWith("A"))
            .map(String::toUpperCase)
            .ifPresent(System.out::println);  // ALICE

        String result = findById(99).orElse("Guest");
        System.out.println("Result: " + result); // Guest
    }
}
```

---

## Q4. Why is String immutable in Java? Explain String immutability with an example.

**Spoken Answer:**
"Once a String object is created its value can never be changed. Any operation like `+` or `concat()` creates a brand-new String object.

Why? Four reasons:
1. **String Pool** — JVM caches String literals. If strings were mutable, one reference changing the value would corrupt all references sharing that literal.
2. **Security** — Strings are used for class names, database URLs, file paths. Mutability would allow malicious code to change them mid-operation.
3. **Thread Safety** — Immutable objects are inherently safe to share across threads without synchronization.
4. **Hash Caching** — String caches its `hashCode`. If it were mutable, HashMap lookups would break when a key's value changes."

```java
import java.util.*;

public class StringImmutability {
    public static void main(String[] args) {

        // String Pool: same literal → same object reference
        String s1 = "Java";
        String s2 = "Java";
        String s3 = new String("Java"); // bypasses pool — heap object

        System.out.println(s1 == s2);      // true  — same pool reference
        System.out.println(s1 == s3);      // false — s3 is a new heap object
        System.out.println(s1.equals(s3)); // true  — same content

        // Immutability: concat creates a NEW object, original is unchanged
        String original = "Hello";
        String modified = original.concat(" World");
        System.out.println(original);  // Hello        ← untouched
        System.out.println(modified);  // Hello World  ← new object

        // String as HashMap key stays consistent because hashCode never changes
        Map<String, Integer> map = new HashMap<>();
        String key = "employee";
        map.put(key, 100);
        System.out.println(map.get("employee")); // 100 — always works

        // Use StringBuilder when you need mutable string building
        StringBuilder sb = new StringBuilder("Hello");
        sb.append(" World").append("!");
        System.out.println(sb); // Hello World! — same object, no new allocation
    }
}
```

---

## Q5. What is ConcurrentModificationException? What is an Iterator and how do you fix CME?

**Spoken Answer:**
"ConcurrentModificationException is thrown when you modify a collection — add or remove — while iterating over it using a for-each loop or an Iterator. Internally, collections track a `modCount`. When the iterator checks this count and finds it changed, it throws CME as a fail-fast signal.

The fix: use `iterator.remove()` which updates `modCount` properly. Or use `removeIf()` in Java 8+. Or use `CopyOnWriteArrayList` for concurrent thread scenarios."

```java
import java.util.*;
import java.util.concurrent.CopyOnWriteArrayList;

public class CMEDemo {
    public static void main(String[] args) {

        // ❌ Throws ConcurrentModificationException
        List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        try {
            for (Integer num : list) {
                if (num % 2 == 0) list.remove(num); // modifies during iteration → CME
            }
        } catch (ConcurrentModificationException e) {
            System.out.println("CME: " + e.getClass().getSimpleName());
        }

        // ✅ Fix 1: Iterator.remove() — updates modCount properly
        List<Integer> list2 = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        Iterator<Integer> it = list2.iterator();
        while (it.hasNext()) {
            if (it.next() % 2 == 0) it.remove(); // safe!
        }
        System.out.println("Iterator remove: " + list2); // [1, 3, 5]

        // ✅ Fix 2: removeIf (Java 8+) — cleanest
        List<Integer> list3 = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        list3.removeIf(n -> n % 2 == 0);
        System.out.println("removeIf:        " + list3); // [1, 3, 5]

        // ✅ Fix 3: CopyOnWriteArrayList — iterates a snapshot copy
        List<Integer> cowList = new CopyOnWriteArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        for (Integer num : cowList) {
            if (num % 2 == 0) cowList.remove(num); // no CME — iterates snapshot
        }
        System.out.println("CopyOnWrite:     " + cowList); // [1, 3, 5]
    }
}
```

---

## Q6. What is the difference between HashMap and ConcurrentHashMap? How does ConcurrentHashMap work internally?

**Spoken Answer:**
"HashMap is **not thread-safe**. In multi-threaded environments, concurrent puts can cause data corruption or infinite loops.

ConcurrentHashMap is thread-safe without locking the whole map. In Java 8+, it uses **CAS (Compare-And-Swap)** for some operations and `synchronized` only on individual bucket nodes — so multiple threads can read/write to different buckets simultaneously.

Key differences:
- HashMap allows one null key and null values. ConcurrentHashMap allows **neither** — null would be ambiguous (is the key absent or mapped to null?).
- ConcurrentHashMap also provides atomic compound operations like `computeIfAbsent()` and `putIfAbsent()`."

```java
import java.util.*;
import java.util.concurrent.*;

public class HashMapVsConcurrentHashMap {
    public static void main(String[] args) throws InterruptedException {

        // HashMap: null key/value allowed
        Map<String, String> hm = new HashMap<>();
        hm.put(null, "nullKey");          // ✅
        hm.put("key", null);              // ✅

        // ConcurrentHashMap: NO null key or value
        Map<String, String> chm = new ConcurrentHashMap<>();
        try { chm.put(null, "v"); } catch (NullPointerException e) {
            System.out.println("CHM null key → NullPointerException");
        }

        // Thread-safety demo: 10 threads writing concurrently
        Map<Integer, Integer> concurrentMap = new ConcurrentHashMap<>();
        ExecutorService pool = Executors.newFixedThreadPool(10);
        for (int i = 0; i < 100; i++) {
            final int id = i;
            pool.submit(() -> concurrentMap.put(id, id * 10));
        }
        pool.shutdown();
        pool.awaitTermination(3, TimeUnit.SECONDS);
        System.out.println("CHM size after 100 concurrent puts: " + concurrentMap.size()); // 100

        // Atomic compound operations (not possible safely with HashMap)
        concurrentMap.putIfAbsent(200, 999);      // put only if key absent
        concurrentMap.compute(1, (k, v) -> v == null ? 1 : v + 1); // atomic update
        concurrentMap.merge(1, 100, Integer::sum); // merge with existing
        System.out.println("Key 1: " + concurrentMap.get(1));
    }
}
```

---

## Q7. Why are immutable objects thread-safe? If there are 1000 threads, do you create 1000 copies? Is that cost-effective?

**Spoken Answer:**
"Immutable objects are thread-safe because their state **never changes after construction**. There's nothing to lock or synchronize because no thread can ever modify the object.

No — you absolutely do NOT create 1000 copies. That would waste enormous memory and is not cost-effective at all. The whole point of immutability is that **one single instance can be safely shared** by all 1000 threads simultaneously. No copy needed."

```java
public class ImmutableThreadSafety {

    // Immutable class: all fields final, no setters, no mutable state exposed
    static final class AppConfig {
        private final String host;
        private final int port;
        private final int timeout;

        AppConfig(String host, int port, int timeout) {
            this.host = host;
            this.port = port;
            this.timeout = timeout;
        }
        public String getHost()    { return host; }
        public int getPort()       { return port; }
        public int getTimeout()    { return timeout; }
    }

    public static void main(String[] args) throws InterruptedException {

        // ONE shared instance — safe for ALL threads simultaneously
        AppConfig sharedConfig = new AppConfig("db.prod.com", 5432, 3000);

        Thread[] threads = new Thread[10]; // demo with 10 (concept applies to 1000)
        for (int i = 0; i < 10; i++) {
            threads[i] = new Thread(() -> {
                // All threads safely read the same object — no lock needed
                String info = sharedConfig.getHost() + ":" + sharedConfig.getPort();
                System.out.println(Thread.currentThread().getName() + " reads → " + info);
            }, "Thread-" + i);
            threads[i].start();
        }
        for (Thread t : threads) t.join();
        // Result: all 10 threads share ONE AppConfig — no copies, no locks ✅
    }
}
```

---

## Q8. What is a subsequence? What is the Longest Common Subsequence (LCS)?

**Coding Problem:** `String1 = "ABC"`, `String2 = "ACD"` → Output: `AC`, Length: `2`

**Spoken Answer:**
"A subsequence is a sequence derived from another by deleting some characters without changing the relative order. 'AC' is a subsequence of 'ABC'.

LCS is the longest such sequence common to both strings. I solve it with dynamic programming. `dp[i][j]` stores the LCS length for the first `i` chars of s1 and first `j` chars of s2. If characters match, I take the diagonal + 1. Otherwise, I take the max of left or top cell. Then I backtrack to recover the actual string."

```java
public class LCS {

    static String findLCS(String s1, String s2) {
        int m = s1.length(), n = s2.length();
        int[][] dp = new int[m + 1][n + 1];

        // Build DP table
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (s1.charAt(i - 1) == s2.charAt(j - 1)) {
                    dp[i][j] = 1 + dp[i - 1][j - 1]; // match: diagonal + 1
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]); // no match: best of above/left
                }
            }
        }

        // Backtrack to get the actual LCS string
        StringBuilder sb = new StringBuilder();
        int i = m, j = n;
        while (i > 0 && j > 0) {
            if (s1.charAt(i - 1) == s2.charAt(j - 1)) {
                sb.insert(0, s1.charAt(i - 1));
                i--; j--;
            } else if (dp[i - 1][j] > dp[i][j - 1]) {
                i--;
            } else {
                j--;
            }
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        String s1 = "ABC", s2 = "ACD";
        String result = findLCS(s1, s2);
        System.out.println("LCS:    " + result);           // AC
        System.out.println("Length: " + result.length());  // 2
    }
}
```

---

## Q9. What are access specifiers in Java? Name a few you are familiar with.

**Spoken Answer:**
"Java has four access specifiers that control visibility:
- `public` — accessible everywhere, any class, any package.
- `protected` — accessible within the same package AND subclasses (even across packages).
- `default` (no keyword) — accessible only within the same package. Also called package-private.
- `private` — accessible only within the same class. Most restrictive."

```java
package com.virtusa.demo;

public class AccessSpecDemo {

    public    String pub  = "anyone can see me";
    protected String prot = "same package or subclasses";
              String def  = "same package only (package-private)";
    private   String priv = "only this class";

    public void show() {
        // All accessible inside the class
        System.out.println(pub + " | " + prot + " | " + def + " | " + priv);
    }
}

class SamePackageClass {
    void test() {
        AccessSpecDemo obj = new AccessSpecDemo();
        System.out.println(obj.pub);   // ✅ public
        System.out.println(obj.prot);  // ✅ protected (same package)
        System.out.println(obj.def);   // ✅ default   (same package)
        // System.out.println(obj.priv); // ❌ private — compile error
    }
}

// In a DIFFERENT package:
// class OtherPackageClass {
//     void test() {
//         System.out.println(obj.pub);   // ✅ public only
//         System.out.println(obj.prot);  // ❌ not accessible (not subclass)
//         System.out.println(obj.def);   // ❌ not accessible (different package)
//         System.out.println(obj.priv);  // ❌ not accessible
//     }
// }
```

---

## Q10. What happens when you insert different objects with the same data into a HashSet?

**Spoken Answer:**
"HashSet uses `hashCode()` and `equals()` to determine duplicates. If you insert two objects with the same data but haven't overridden `hashCode()` and `equals()`, they will have **different hashCodes** (from Object's default identity-based implementation) — so HashSet treats them as different objects and stores both.

To make HashSet correctly deduplicate, you MUST override both `hashCode()` and `equals()`. Java 14+ Records do this automatically."

```java
import java.util.*;

public class HashSetDemo {

    // ❌ Without overriding: treated as different objects
    static class BadEmployee {
        String name; int id;
        BadEmployee(String name, int id) { this.name = name; this.id = id; }
        // NO hashCode/equals override
    }

    // ✅ With overriding: correctly deduplicated
    static class GoodEmployee {
        String name; int id;
        GoodEmployee(String name, int id) { this.name = name; this.id = id; }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (!(o instanceof GoodEmployee ge)) return false;
            return id == ge.id && Objects.equals(name, ge.name);
        }

        @Override
        public int hashCode() {
            return Objects.hash(name, id);
        }

        @Override public String toString() { return name + "(" + id + ")"; }
    }

    public static void main(String[] args) {

        // ❌ No override: both objects stored (different default hashCodes)
        Set<BadEmployee> badSet = new HashSet<>();
        badSet.add(new BadEmployee("Alice", 1));
        badSet.add(new BadEmployee("Alice", 1)); // same data, different object
        System.out.println("Bad set size: " + badSet.size()); // 2 ← WRONG!

        // ✅ With override: duplicate correctly rejected
        Set<GoodEmployee> goodSet = new HashSet<>();
        goodSet.add(new GoodEmployee("Alice", 1));
        goodSet.add(new GoodEmployee("Alice", 1)); // same data
        System.out.println("Good set size: " + goodSet.size()); // 1 ✅

        // Java Records auto-generate hashCode & equals based on components
        record Employee(String name, int id) {}
        Set<Employee> recordSet = new HashSet<>();
        recordSet.add(new Employee("Bob", 2));
        recordSet.add(new Employee("Bob", 2));
        System.out.println("Record set size: " + recordSet.size()); // 1 ✅
    }
}
```
