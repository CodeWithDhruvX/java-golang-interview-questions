# ⚡ Exceptions, IO & Collections (Rapid-Fire)

> 🔑 **Master Keyword:** **"ETICS"** → Exceptions, Transient, IO-streams, Collections, Serialization

---

## ❌ Section 1: Exception Handling

### Q1: Exception Hierarchy?
🔑 **Keyword: "T-E-C-U"** → Throwable > Error/Exception > Checked/Unchecked

```
Throwable
├── Error (JVM Fatal — don't catch!)
│   ├── StackOverflowError
│   └── OutOfMemoryError
└── Exception
    ├── Checked (Compile-time — must handle)
    │   ├── IOException
    │   └── SQLException
    └── Unchecked/RuntimeException (Bugs — optional)
        ├── NullPointerException
        └── ArrayIndexOutOfBoundsException
```

---

### Q2: `throw` vs `throws`?
🔑 **Keyword: "II-MS"** → throw=Inside-method, throws=Method-Signature

```java
// throw — inside method body
throw new RuntimeException("Error occurred");

// throws — in method signature (declare potential exceptions)
public void readFile() throws IOException { }
```

---

### Q3: `try-catch-finally` combinations?
🔑 **Keyword: "TCF-Always"** → finally Always runs (unless System.exit)

- `try` alone → ❌ Compilation error
- `try-catch` → ✅
- `try-finally` → ✅
- `try-catch-finally` → ✅
- **`finally` always runs** — even if exception. Only exception: `System.exit()`

---

### Q4: Try-with-Resources (Java 7+)?
🔑 **Keyword: "ARCA"** → AutoCloseable→Resource→Close→Automatic

```java
try (BufferedReader br = new BufferedReader(new FileReader("file.txt"))) {
    String line = br.readLine();
} catch (IOException e) {
    e.printStackTrace();
} // br.close() called automatically!
```
Any class implementing `AutoCloseable` works. Eliminates `finally` for cleanup.

---

### Q5: Checked vs Unchecked — When to use?
🔑 **Keyword: "R-P"** → Recoverable=Checked, Programming-error=Unchecked

- **Checked:** Recoverable conditions outside program control (file missing, DB down) → forces caller to handle
- **Unchecked:** Programming bugs (null pointer, bad arg) → caller can't recover anyway

---

### Q6: Custom Exception?
🔑 **Keyword: "ER"** → Extend RuntimeException or Exception

```java
// Unchecked custom exception
public class InsufficientFundsException extends RuntimeException {
    public InsufficientFundsException(String msg) {
        super(msg);
    }
}

// Checked custom exception
public class OrderNotFoundException extends Exception {
    public OrderNotFoundException(String msg) { super(msg); }
}
```

---

### Q7: Exception in `finally` block?
🔑 **Keyword: "Mask"** → Finally-exception masks the original

- Exception in `finally` **consumes/masks** the original exception
- Caller only sees the `finally` exception
- **Best Practice:** Never throw from `finally`

---

### Q8: Exception Propagation?
🔑 **Keyword: "UA"** → Unchecked=Automatic, Checked=declared

- **Unchecked:** Auto-propagates up call stack until caught (or thread dies)
- **Checked:** Must explicitly propagate using `throws` at every level

---

## 📁 Section 2: IO & Serialization

### Q9: What is Serialization?
🔑 **Keyword: "OBS"** → Object→Byte-Stream (to file/network)

- Converting object to byte stream → save to file or send over network
- Implement `java.io.Serializable` (marker interface)
- `ObjectOutputStream.writeObject(obj)` to serialize
- `ObjectInputStream.readObject()` to deserialize

---

### Q10: `serialVersionUID` significance?
🔑 **Keyword: "VID-match"** → Version-ID must match

```java
private static final long serialVersionUID = 1L;
```
- Unique ID for class version
- If class changed and UID mismatches → `InvalidClassException`
- **Always define it manually** to control compatibility

---

### Q11: `transient` keyword?
🔑 **Keyword: "Skip-sensitive"** → Skip field during serialization

```java
private transient String password; // NOT serialized
// After deserialization: password = null (default value)
```
Use for: passwords, sensitive data, computed fields

---

### Q12: `Externalizable` vs `Serializable`?
🔑 **Keyword: "CM-Auto"** → Custom-Manual vs Auto-JVM

| Feature | Serializable | Externalizable |
|---|---|---|
| Control | JVM does everything | You implement `writeExternal`/`readExternal` |
| Performance | Slower (all fields) | Faster (you pick what to save) |
| Usage | Simple, quick | Performance-critical serialization |

---

### Q13: Byte Stream vs Character Stream?
🔑 **Keyword: "BRC"** → Bytes=Raw-Content, Chars=Text/Unicode

| Type | Classes | Use-case |
|---|---|---|
| Byte Stream | `InputStream`, `OutputStream` | Images, audio, raw binary |
| Character Stream | `Reader`, `Writer` | Text files, auto-handles encoding |

- Bridge: `InputStreamReader` converts Bytes → Characters

---

### Q14: `Scanner` vs `BufferedReader`?
🔑 **Keyword: "SBP-Fast"** → Scanner=Parse, BufferedReader=Performance

| Feature | Scanner | BufferedReader |
|---|---|---|
| Speed | Slower (regex parsing) | Faster (line-based) |
| Parsing | `nextInt()`, `nextBoolean()` | `readLine()` only |
| Buffer | Small | 8KB by default |

---

### Q15: `File` vs `Path` (NIO.2)?
🔑 **Keyword: "LP-NIO"** → Legacy=File, Path=NIO/Modern

```java
// Legacy (blocking)
File f = new File("data.txt");

// Modern NIO (better error handling, non-blocking)
Path p = Paths.get("data.txt");
List<String> lines = Files.readAllLines(p);
```

---

## 📦 Section 3: Collections Framework

### Q16: Collections Hierarchy?
🔑 **Keyword: "LOST-MAP"** → List/Order, Set/Unique, TreeSet/Sorted, Map/Key-Value

```
Collection
├── List (ordered, duplicates OK)
│   ├── ArrayList, LinkedList, Vector, Stack
├── Set (unordered, no duplicates)
│   ├── HashSet, LinkedHashSet, TreeSet
└── Queue / Deque
    └── PriorityQueue, ArrayDeque, LinkedList

Map (Key-Value, separate hierarchy)
    ├── HashMap, LinkedHashMap, TreeMap, Hashtable
```

---

### Q17: ArrayList vs LinkedList?
🔑 **Keyword: "ARML"** → Array=Random-access, LinkedList=Modification

| Feature | ArrayList | LinkedList |
|---|---|---|
| Internal | Dynamic array | Doubly linked list |
| Get(i) | O(1) | O(n) |
| Add/Remove (middle) | O(n) | O(1) |
| Memory | Less (contiguous) | More (node pointers) |
| Best for | Read-heavy | Insert/Delete-heavy |

---

### Q18: HashMap Internals?
🔑 **Keyword: "HABT"** → Hash→Array→Bucket→Tree

1. `key.hashCode()` → hash function → **index** in array
2. Collision → **LinkedList** in that bucket
3. Java 8+: When bucket size > 8 → converts to **Red-Black Tree** (O(log n))
4. Load factor = 0.75 → when 75% full → **rehash** (resize to 2x)

```java
Map<String, Integer> map = new HashMap<>();
map.put("a", 1); // key.hashCode() → index → store
map.get("a");    // key.hashCode() → index → equals() → retrieve
```

---

### Q19: `HashMap` vs `Hashtable` vs `ConcurrentHashMap`?
🔑 **Keyword: "HSC-Thread"** → HashMap=no-sync, Hashtable=sync-whole, ConcurrentHashMap=segment-sync

| Feature | HashMap | Hashtable | ConcurrentHashMap |
|---|---|---|---|
| Thread-safe | ❌ | ✅ (whole lock) | ✅ (segment lock) |
| Null key/value | key=1 null, multi null values | ❌ | ❌ |
| Performance | Best single-thread | Slowest | Best multi-thread |

---

### Q20: HashSet vs LinkedHashSet vs TreeSet?
🔑 **Keyword: "HLT-Order"** → Hash=none, Linked=insertion, Tree=sorted

| Set | Ordering | Performance |
|---|---|---|
| `HashSet` | No order | O(1) |
| `LinkedHashSet` | Insertion order | O(1) |
| `TreeSet` | Natural sorted order | O(log n) |

---

### Q21: `Comparable` vs `Comparator`?
🔑 **Keyword: "NE-CE"** → Natural=Comparable(internal), Custom=Comparator(external)

```java
// Comparable — natural ordering (inside the class)
class Student implements Comparable<Student> {
    public int compareTo(Student o) { return this.marks - o.marks; }
}

// Comparator — custom ordering (outside the class)
Comparator<Student> byName = (a, b) -> a.name.compareTo(b.name);
students.sort(byName);
```

---

### Q22: `Iterator` vs `ListIterator`?
🔑 **Keyword: "IF-LFB"** → Iterator=Forward, ListIterator=Forward+Backward

- `Iterator` → forward only, `remove()`
- `ListIterator` → both directions, `add()`, `set()`, `previousIndex()`

---

### Q23: `fail-fast` vs `fail-safe` Iterators?
🔑 **Keyword: "FFFE"** → FailFast=original, FailSafe=copy/safe

| Feature | Fail-Fast | Fail-Safe |
|---|---|---|
| Examples | `ArrayList`, `HashMap` | `CopyOnWriteArrayList`, `ConcurrentHashMap` |
| Throws | `ConcurrentModificationException` | No exception |
| Works on | Original collection | Copy of collection |

---

### Q24: `BlockingQueue` — What is it?
🔑 **Keyword: "BQ-PC"** → BlckingQueue=Producer-Consumer pattern

```java
BlockingQueue<String> queue = new LinkedBlockingQueue<>(10);
queue.put("task");   // blocks if full
queue.take();        // blocks if empty
```
Thread-safe queue for Producer-Consumer pattern.

---

### Q25: `PriorityQueue`?
🔑 **Keyword: "MinHeap"** → PriorityQueue=Min-Heap by default

```java
PriorityQueue<Integer> pq = new PriorityQueue<>(); // min-heap
pq.add(5); pq.add(1); pq.add(3);
pq.poll(); // returns 1 (smallest)
```
For max-heap: `new PriorityQueue<>(Collections.reverseOrder())`

---

## 🔀 Section 4: Streams & Functional (Java 8)

### Q26: Stream Pipeline?
🔑 **Keyword: "SIT"** → Source→Intermediate→Terminal

```java
List<String> names = List.of("Alice", "Bob", "Charlie");
long count = names.stream()              // Source
    .filter(n -> n.startsWith("A"))      // Intermediate (lazy)
    .map(String::toUpperCase)            // Intermediate (lazy)
    .count();                            // Terminal (triggers execution)
```

---

### Q27: `map()` vs `flatMap()`?
🔑 **Keyword: "MO-FF"** → Map=One-to-One, FlatMap=Flatten

```java
// map — one-to-one transformation
List<String> upper = names.stream().map(String::toUpperCase).toList();

// flatMap — flatten nested collections
List<List<String>> nested = List.of(List.of("a","b"), List.of("c","d"));
List<String> flat = nested.stream().flatMap(Collection::stream).toList();
// Result: ["a", "b", "c", "d"]
```

---

### Q28: Core Functional Interfaces?
🔑 **Keyword: "PCSFB"** → Predicate/Consumer/Supplier/Function/BiFunction

| Interface | Signature | Usage |
|---|---|---|
| `Predicate<T>` | `T → boolean` | `filter()` |
| `Consumer<T>` | `T → void` | `forEach()` |
| `Supplier<T>` | `() → T` | `generate()`, lazy load |
| `Function<T,R>` | `T → R` | `map()` |
| `UnaryOperator<T>` | `T → T` | `replaceAll()` |
| `BinaryOperator<T>` | `(T,T) → T` | `reduce()` |

---

### Q29: `Optional` — What and Why?
🔑 **Keyword: "ANP"** → Avoid-NullPointer

```java
Optional<String> name = Optional.ofNullable(getName());
name.ifPresent(n -> System.out.println(n));
String result = name.orElse("Unknown");
String val = name.orElseThrow(() -> new RuntimeException("Not found"));
```
Expresses "may or may not have value". Forces caller to handle null case explicitly.

---

### Q30: `Stream.collect()` — Common Collectors?
🔑 **Keyword: "LGMJ"** → List/Grouping/Map/Joining

```java
// Collect to list
List<String> list = stream.collect(Collectors.toList());

// Group by
Map<String, List<Employee>> deptMap = employees.stream()
    .collect(Collectors.groupingBy(Employee::getDept));

// Join strings
String csv = names.stream().collect(Collectors.joining(", "));

// Count by group
Map<String, Long> countByDept = employees.stream()
    .collect(Collectors.groupingBy(Employee::getDept, Collectors.counting()));
```

---

*End of File — Exceptions, IO & Collections*
