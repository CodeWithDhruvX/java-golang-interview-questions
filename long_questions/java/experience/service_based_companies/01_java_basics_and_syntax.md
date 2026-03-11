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

### 🎯 How to Explain in Interview

"Well, Java has 8 primitive data types that are the building blocks of all data in Java. Think of them as the basic LEGO blocks - you have byte, short, int, and long for whole numbers, float and double for decimal numbers, char for single characters, and boolean for true/false values. Each has a specific size and range - for example, int is 4 bytes and can hold values from -2 billion to +2 billion, while long is 8 bytes for really big numbers. The key thing to remember is that primitives hold actual values directly in memory, which makes them fast and memory-efficient. We also have wrapper classes like Integer and Double because collections like ArrayList can only work with objects, not primitives."

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

### 🎯 How to Explain in Interview

"Autoboxing is Java's automatic conversion between primitives and their wrapper classes. When I write `Integer x = 42`, Java automatically converts the primitive int 42 to an Integer object - that's autoboxing. The reverse, converting wrapper back to primitive, is called unboxing. This makes code cleaner, but there's a catch - it can cause performance issues in loops because it creates lots of objects, and more importantly, if you try to unbox a null wrapper, you'll get a NullPointerException at runtime. So while convenient, we need to be careful with autoboxing in performance-critical code."

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

### 🎯 How to Explain in Interview

"Well, when it comes to strings in Java, I need to think about mutability and thread safety. String is immutable - once created, it can't be changed. When I do `s1 = s1 + " World"`, Java actually creates a completely new String object. This makes String thread-safe but inefficient for lots of modifications. StringBuilder is mutable - I can change it in place, which is much faster, but it's not thread-safe. StringBuffer is the same as StringBuilder but with synchronized methods, making it thread-safe but slower. So my rule is: use String for most cases, StringBuilder when doing lots of string manipulation in single-threaded code, and StringBuffer only when working with multiple threads."

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

### 🎯 How to Explain in Interview

"This is a classic Java interview question! The `==` operator checks if two references point to the exact same object in memory, while `.equals()` checks if the contents of two objects are the same. For strings, Java does something interesting - it maintains a string pool. So when I write `String a = "hello"` and `String b = "hello"`, both references point to the same pool object, so `a == b` is true. But if I create a new string with `new String("hello")`, it creates a separate object, so `a == c` is false, but `a.equals(c)` is true because the content is the same. That's why we should always use `.equals()` for object comparison unless we specifically want to check reference equality."

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

### 🎯 How to Explain in Interview

"Access modifiers in Java control the visibility of classes, methods, and variables. Think of them as security levels. `public` is the most open - anyone can access it from anywhere. `protected` is a bit more restrictive - only classes in the same package or subclasses can access it. The default, which we call package-private, is even more restrictive - only classes in the same package can access it. `private` is the most restrictive - only the class itself can access it. This helps with encapsulation - I can hide internal implementation details and only expose what's necessary. For example, I'd make fields private and provide public getters to control how they're accessed."

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

### 🎯 How to Explain in Interview

"The `static` keyword in Java means something belongs to the class itself, not to any particular instance. When I declare a variable as static, there's only one copy of it that's shared across all objects of that class. This is perfect for things like counters or configuration values that should be the same everywhere. Static methods are similar - they don't need an object instance to be called, which is why we can call `Math.random()` without creating a Math object. The key thing is that static members can't access instance variables directly because they don't have a `this` reference. I use static for utility methods, constants, and shared resources."

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

### 🎯 How to Explain in Interview

"Lambdas in Java 8 revolutionized how we write code! Before lambdas, if I wanted to pass behavior as a parameter, I had to write verbose anonymous classes. With lambdas, I can write concise, readable code like `(a, b) -> a.compareTo(b)`. A lambda is essentially a short, anonymous function that I can pass around. It works with functional interfaces - interfaces with just one abstract method. So when I write `names.forEach(name -> System.out.println(name))`, I'm passing a lambda that represents the Consumer interface's accept method. This makes code much more expressive and reduces boilerplate significantly."

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

### 🎯 How to Explain in Interview

"Java Streams are a game-changer for data processing! They let me process collections in a declarative, functional way. Instead of writing traditional for-loops with manual iteration, I can chain operations like `filter().map().collect()`. The beauty of streams is that they're lazy - operations only execute when I call a terminal operation like `collect()`. This allows for optimizations. I can filter a million numbers but only process the first 10, and Java won't waste time processing the rest. Streams also support parallel processing with just one method call - `parallelStream()` - which is amazing for big data tasks on multi-core systems."

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

### 🎯 How to Explain in Interview

"Optional is Java's elegant solution to the billion-dollar mistake - null references! Instead of returning null from a method that might not find something, I can return an Optional. This makes my code much safer and more expressive. An Optional is a container that may or may not contain a value. I can use methods like `ifPresent()` to execute code only if there's a value, `orElse()` to provide a default, or `orElseThrow()` to throw an exception if empty. The best part is that Optional forces me to think about the absence case explicitly, preventing those pesky NullPointerExceptions at runtime. It's like having built-in documentation that says 'this might be empty, handle it!'"

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

### 🎯 How to Explain in Interview

"When choosing between ArrayList and LinkedList, I think about how I'll be using the data. ArrayList is backed by an array, so it's great for random access - getting the 100th element is super fast at O(1). But inserting or deleting at the beginning is slow because all elements need to shift. LinkedList is like a chain of nodes, each pointing to the next, so adding or removing at the head is fast at O(1), but getting to the 100th element requires traversing through 99 nodes, which is O(n). In most real-world scenarios, ArrayList is the better choice because we usually access elements randomly and only add at the end."

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

### 🎯 How to Explain in Interview

"HashMap is fascinating! When I put a key-value pair, Java first calculates the hash of the key to determine which bucket it should go in. Each bucket is essentially a list of entries. If two keys have the same hash, that's a collision, and Java handles this with separate chaining - it stores them in a linked list. But here's the cool part: in Java 8, if a bucket gets too many entries (more than 8), it automatically converts from a linked list to a balanced tree for better performance. HashMap has a default capacity of 16 buckets and resizes when it's 75% full. It's not thread-safe, so in multi-threaded environments, I'd use ConcurrentHashMap which uses lock striping for better concurrency."

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

### 🎯 How to Explain in Interview

"The `final` keyword is all about immutability and preventing change. When I declare a variable as final, I can't reassign it - it's like writing in permanent marker. But here's a common confusion point: if I have a final reference to an object, like `final List<String> list`, I can't change which list it points to, but I can still modify the list's contents. Final methods can't be overridden by subclasses, which is great for security and consistency. Final classes can't be extended at all - that's why String and Integer are final, to prevent malicious subclassing. I use final for constants, method parameters that shouldn't change, and when I want to design immutable classes for thread safety."

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

### 🎯 How to Explain in Interview

"The `var` keyword in Java 10+ is all about reducing boilerplate while maintaining type safety. When I write `var message = "Hello"`, the compiler automatically figures out that message should be a String. It's not dynamic typing - the type is still determined at compile time, I just don't have to write it out. This makes code cleaner, especially with complex generic types like `Map<String, List<Integer>>`. But there are rules: I can only use var for local variables with initializers, not for fields, method parameters, or return types. Also, I try to use var when the type is obvious from the right side, but stick to explicit types when it improves readability."

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

### 🎯 How to Explain in Interview

"Try-with-resources is one of Java's best features for clean code! Before Java 7, I had to write messy finally blocks to ensure resources like database connections were closed, even if exceptions occurred. With try-with-resources, any resource that implements AutoCloseable gets automatically closed when the try block exits. It's like having a built-in cleanup crew. I can declare multiple resources, and they're closed in reverse order of creation. This eliminates resource leaks and makes code much cleaner. It handles exceptions gracefully too - if the try block throws an exception and the close method also throws, Java preserves the original exception."

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

---

### 🎯 How to Explain in Interview

"Java's exception hierarchy is split into two camps: checked and unchecked exceptions. Checked exceptions are the polite ones - the compiler forces me to either catch them or declare that my method throws them. These are for recoverable situations like file not found or network errors. Unchecked exceptions are the runtime ones - they don't need to be declared or caught, and usually indicate programming bugs like NullPointerException or ArrayIndexOutOfBoundsException. The philosophy is that checked exceptions force developers to handle exceptional cases, while unchecked exceptions represent problems that probably can't be recovered from anyway. I create custom checked exceptions for business logic failures and unchecked ones for programming errors."

---

### Q16. Explain the difference between `==` and `.equals()` in Java.

```java
// == checks for reference equality (same object)
String s1 = "hello";
String s2 = "hello";
System.out.println(s1 == s2);  // true

// .equals() checks for content equality (same value)
String s3 = new String("hello");
System.out.println(s1.equals(s3));  // true
System.out.println(s1 == s3);  // false
