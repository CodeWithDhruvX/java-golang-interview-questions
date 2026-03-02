# 📘 01 — Java Basics & Syntax
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy

---

## 🔑 Must-Know Topics
- Primitive types and wrapper classes
- `String`, `StringBuilder`, `StringBuffer`
- Control flow (`if`, `for`, `while`, `switch`)
- Arrays and basic Collections (`ArrayList`, `HashMap`)
- Java 8 features: Lambdas, Streams, Optional
- Static vs instance members
- Variable scoping and type casting

---

## ❓ Most Asked Questions

### Q1. What are the primitive data types in Java?

| Type | Size | Default | Range |
|------|------|---------|-------|
| `byte` | 1 byte | 0 | -128 to 127 |
| `short` | 2 bytes | 0 | -32,768 to 32,767 |
| `int` | 4 bytes | 0 | -2^31 to 2^31-1 |
| `long` | 8 bytes | 0L | -2^63 to 2^63-1 |
| `float` | 4 bytes | 0.0f | ~7 decimal digits |
| `double` | 8 bytes | 0.0d | ~15 decimal digits |
| `char` | 2 bytes | '\u0000' | 0 to 65,535 |
| `boolean` | 1 bit | false | true/false |

> Each primitive has a **Wrapper class** (`Integer`, `Double`, `Boolean`, etc.) for use in Collections.

---

### Q2. What is autoboxing and unboxing?

```java
// Autoboxing — primitive to wrapper automatically
Integer x = 42;          // int → Integer (autoboxed)
List<Integer> list = new ArrayList<>();
list.add(10);            // int 10 autoboxed to Integer

// Unboxing — wrapper to primitive automatically
int val = x;             // Integer → int (unboxed)
int sum = x + 5;         // unboxed before arithmetic

// Pitfall: NullPointerException on unboxing
Integer n = null;
int k = n;               // ❌ NullPointerException at runtime!
```

> **Beware:** Autoboxing inside hot loops causes excessive object creation. Prefer primitives for performance.

---

### Q3. What is the difference between `String`, `StringBuilder`, and `StringBuffer`?

```java
// String — immutable, thread-safe, creates new object on modification
String s1 = "Hello";
s1 = s1 + " World";   // new String object created — old one GC'd

// StringBuilder — mutable, NOT thread-safe, fast in single-threaded code
StringBuilder sb = new StringBuilder("Hello");
sb.append(" World");   // modifies in-place — O(1) amortized
sb.insert(5, ",");
sb.reverse();
String result = sb.toString();

// StringBuffer — mutable, thread-safe (synchronized methods), slower
StringBuffer sbuf = new StringBuffer("Hello");
sbuf.append(" World");  // safe to use across threads
```

| Feature | String | StringBuilder | StringBuffer |
|---------|--------|---------------|--------------|
| Mutable | ❌ | ✅ | ✅ |
| Thread-safe | ✅ (immutable) | ❌ | ✅ |
| Performance | Slow (concatenation) | Fast | Medium |
| Usage | Most cases | Single-thread loops | Multi-thread |

---

### Q4. Explain `==` vs `.equals()` in Java

```java
String a = "hello";
String b = "hello";
String c = new String("hello");

System.out.println(a == b);       // true  — same String pool reference
System.out.println(a == c);       // false — c is a new heap object
System.out.println(a.equals(c));  // true  — compares content

Integer x = 127;
Integer y = 127;
System.out.println(x == y);  // true  — cached in [-128, 127] range

Integer m = 200;
Integer n = 200;
System.out.println(m == n);  // false — outside cache range, different objects
System.out.println(m.equals(n)); // true — compares value
```

> **Rule:** Always use `.equals()` for object comparison. Use `==` only for primitives or intentional reference check.

---

### Q5. What are Java's access modifiers?

| Modifier | Class | Package | Subclass | World |
|----------|-------|---------|----------|-------|
| `public` | ✅ | ✅ | ✅ | ✅ |
| `protected` | ✅ | ✅ | ✅ | ❌ |
| (default) | ✅ | ✅ | ❌ | ❌ |
| `private` | ✅ | ❌ | ❌ | ❌ |

```java
public class BankAccount {
    private double balance;       // only within this class
    protected String accountId;   // subclasses can access
    String bankName;              // package-private (default)
    public String getOwner() { return "Alice"; }  // everywhere
}
```

---

### Q6. What is `static` in Java?

```java
public class Counter {
    private static int count = 0;  // shared across ALL instances
    private int id;

    public Counter() {
        count++;
        this.id = count;
    }

    public static int getCount() {  // no 'this' — no instance needed
        return count;
    }

    // Static initializer block — runs once when class is loaded
    static {
        System.out.println("Counter class loaded");
    }
}

// Usage
Counter.getCount();          // call without creating object
Counter c1 = new Counter();  // count = 1
Counter c2 = new Counter();  // count = 2
System.out.println(Counter.getCount()); // 2
```

---

### Q7. What are Java 8 Lambdas?

```java
// Before Java 8 — anonymous class
Runnable r = new Runnable() {
    @Override
    public void run() { System.out.println("Hello"); }
};

// Java 8 Lambda — concise syntax
Runnable r = () -> System.out.println("Hello");

// Lambda with parameters
Comparator<String> comp = (a, b) -> a.compareTo(b);

// Functional interface usage
List<String> names = Arrays.asList("Charlie", "Alice", "Bob");
names.sort((a, b) -> a.compareTo(b));
names.forEach(name -> System.out.println(name));
names.forEach(System.out::println);  // method reference
```

---

### Q8. What are Java Streams?

```java
List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

// Filter + Map + Collect
List<Integer> result = numbers.stream()
    .filter(n -> n % 2 == 0)    // keep even numbers
    .map(n -> n * n)             // square them
    .collect(Collectors.toList()); // [4, 16, 36, 64, 100]

// Reduce
int sum = numbers.stream()
    .reduce(0, Integer::sum);  // 55

// GroupingBy
Map<Boolean, List<Integer>> grouped = numbers.stream()
    .collect(Collectors.groupingBy(n -> n % 2 == 0));
// {true=[2,4,6,8,10], false=[1,3,5,7,9]}

// Count, min, max
long count = numbers.stream().filter(n -> n > 5).count(); // 5
Optional<Integer> max = numbers.stream().max(Integer::compareTo); // 10
```

---

### Q9. What is `Optional` and how is it used?

```java
// Avoid NullPointerException with Optional
public Optional<User> findUserById(int id) {
    return Optional.ofNullable(userRepository.get(id));
}

// Usage
Optional<User> user = findUserById(42);

// ifPresent — execute only if value exists
user.ifPresent(u -> System.out.println(u.getName()));

// orElse — provide default
String name = user.map(User::getName).orElse("Unknown");

// orElseThrow — throw if absent
User u = user.orElseThrow(() -> new RuntimeException("Not found"));

// filter + map chain
String email = findUserById(42)
    .filter(u -> u.isActive())
    .map(User::getEmail)
    .orElse("no-email@example.com");
```

---

### Q10. What is the difference between `ArrayList` and `LinkedList`?

```java
// ArrayList — backed by dynamic array
ArrayList<String> al = new ArrayList<>();
al.add("A");        // O(1) amortized
al.get(2);          // O(1) — random access
al.remove(0);       // O(n) — shifts elements

// LinkedList — doubly-linked list, also implements Deque
LinkedList<String> ll = new LinkedList<>();
ll.add("A");        // O(1)
ll.get(2);          // O(n) — must traverse
ll.addFirst("Z");   // O(1) — efficient head insertion
ll.poll();          // O(1) — dequeue from head
```

| Operation | ArrayList | LinkedList |
|-----------|-----------|------------|
| Random access | O(1) | O(n) |
| Add at end | O(1) amortized | O(1) |
| Add at beginning | O(n) | O(1) |
| Remove at beginning | O(n) | O(1) |
| Memory | Less (array) | More (node pointers) |

> **Use:** `ArrayList` for most cases; `LinkedList` only when frequent head/tail insertions are needed.

---

### Q11. How does `HashMap` work internally?

```java
Map<String, Integer> map = new HashMap<>();
map.put("apple", 1);   // hash("apple") → bucket index → store
map.get("apple");       // hash("apple") → same bucket → return

// Collision resolution — separate chaining (linked list, then tree at 8+)
// Java 8+: bucket converts to Red-Black Tree when chain > 8 entries

// Key facts:
// - Default initial capacity: 16
// - Load factor: 0.75 (rehash when 75% full)
// - NOT thread-safe — use ConcurrentHashMap in multi-threaded code
// - Allows ONE null key, multiple null values

// Thread-safe alternatives
Map<String, Integer> safe = new ConcurrentHashMap<>();
Map<String, Integer> synced = Collections.synchronizedMap(new HashMap<>());
```

---

### Q12. What is the `final` keyword?

```java
// final variable — cannot be reassigned
final int MAX = 100;
// MAX = 200;  ❌ compilation error

// final reference — reference cannot change, but object can be mutated
final List<String> list = new ArrayList<>();
list.add("A");   // ✅ — list contents can change
// list = new ArrayList<>();  ❌ — reference cannot change

// final method — cannot be overridden
class Base {
    final void display() { System.out.println("Base"); }
}

// final class — cannot be extended (e.g., String, Integer)
final class ImmutableConfig {
    private final String host;
    public ImmutableConfig(String host) { this.host = host; }
    public String getHost() { return host; }
}
```

---

### Q13. What is `var` in Java (Java 10+)?

```java
// Local variable type inference — compiler infers type
var message = "Hello, Java 10!";     // inferred as String
var numbers = new ArrayList<Integer>(); // inferred as ArrayList<Integer>
var map = new HashMap<String, List<Integer>>();

// In loops
for (var entry : map.entrySet()) {
    System.out.println(entry.getKey() + ": " + entry.getValue());
}

// NOT allowed:
// - class/instance fields
// - method parameters
// - return types
// var x;  ❌ — cannot infer without initializer
```

---

### Q14. Explain `try-with-resources`

```java
// Before Java 7 — manual close
Connection conn = null;
try {
    conn = getConnection();
    // use conn
} finally {
    if (conn != null) conn.close();  // must manually close
}

// Java 7+ try-with-resources — auto-closes AutoCloseable
try (Connection conn = getConnection();
     PreparedStatement ps = conn.prepareStatement("SELECT * FROM users")) {
    ResultSet rs = ps.executeQuery();
    while (rs.next()) {
        System.out.println(rs.getString("name"));
    }
}  // conn and ps automatically closed here, even if exception thrown
```

---

### Q15. What are checked vs unchecked exceptions?

```java
// Checked exceptions — must be caught or declared in method signature
public void readFile(String path) throws IOException {   // must declare
    FileReader fr = new FileReader(path);   // throws FileNotFoundException (checked)
}

// Unchecked exceptions (RuntimeException subclasses) — no forced handling
int[] arr = new int[5];
arr[10] = 1;           // ArrayIndexOutOfBoundsException (unchecked)

String s = null;
s.length();            // NullPointerException (unchecked)

int result = 10 / 0;  // ArithmeticException (unchecked)

// Custom exceptions
public class UserNotFoundException extends RuntimeException {  // unchecked
    public UserNotFoundException(String message) {
        super(message);
    }
}
// vs
public class InsufficientFundsException extends Exception {   // checked
    private double amount;
    public InsufficientFundsException(double amount) {
        super("Insufficient funds: need " + amount);
        this.amount = amount;
    }
}
```
