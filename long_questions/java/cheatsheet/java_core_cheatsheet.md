# Complete Java Developer Cheat Sheet

## 1. Java Reserved Keywords (Complete List)

### Access Modifiers

#### `private`
- **Definition**: Access level only within the declaring class.
- **Syntax**: `private int secret;`
- **Example**:
```java
class Secret {
    private int code = 1234;
    public void show() { System.out.println(code); }
}
public class Main {
    public static void main(String[] args) {
        new Secret().show();
    }
}
```
- **Interview Note**: Most restrictive access modifier; essential for encapsulation.
- **Mistake**: Trying to access private members from subclass or main.

#### `protected`
- **Definition**: Access level within package and subclasses (even outside package).
- **Syntax**: `protected void legacyMethod() {}`
- **Example**:
```java
class Parent { protected void grow() { System.out.println("Growing"); } }
class Child extends Parent { 
    void doGrow() { grow(); } 
}
public class Main {
    public static void main(String[] args) { new Child().doGrow(); }
}
```
- **Interview Note**: Often misused; remember it's visible to the entire package, not just subclasses.

#### `public`
- **Definition**: Access level everywhere.
- **Syntax**: `public class Main {}`
- **Example**:
```java
public class Shared {
    public int data = 10;
}
public class Main {
    public static void main(String[] args) { System.out.println(new Shared().data); }
}
```
- **Interview Note**: Top-level classes can only be public or package-private (default).

### Class, Method, Variable Modifiers

#### `abstract`
- **Definition**: Creating incomplete classes or methods.
- **Syntax**: `abstract class Shape { abstract void draw(); }`
- **Example**:
```java
abstract class Shape { abstract void draw(); }
class Circle extends Shape {
    void draw() { System.out.println("O"); }
}
public class Main {
    public static void main(String[] args) { new Circle().draw(); }
}
```
- **Interview Note**: Abstract classes cannot be instantiated; useful for partial implementation.

#### `class`
- **Definition**: Declares a new class.
- **Syntax**: `class User {}`
- **Example**:
```java
class User { String name = "Dev"; }
public class Main {
    public static void main(String[] args) { System.out.println(new User().name); }
}
```
- **Interview Note**: Java is class-based; almost everything lives in a class.

#### `extends`
- **Definition**: Indicates inheritance from a class.
- **Syntax**: `class Son extends Father {}`
- **Example**:
```java
class A { void a() { System.out.println("A"); } }
class B extends A {}
public class Main {
    public static void main(String[] args) { new B().a(); }
}
```
- **Interview Note**: Java supports single class inheritance only.

#### `final`
- **Definition**: Prevents modification (variable cannot change, method cannot override, class cannot extend).
- **Syntax**: `final int MAX = 10;`
- **Example**:
```java
final class Constant { final int VAL = 10; }
public class Main {
    public static void main(String[] args) {
        System.out.println(new Constant().VAL);
    }
}
```
- **Interview Note**: `final` on reference means reference can't change, but object state can.

#### `implements`
- **Definition**: Indicates implementation of an interface.
- **Syntax**: `class Task implements Runnable {}`
- **Example**:
```java
interface I { void m(); }
class C implements I {
    public void m() { System.out.println("Done"); }
}
public class Main {
    public static void main(String[] args) { new C().m(); }
}
```
- **Interview Note**: A class can implement multiple interfaces.

#### `interface`
- **Definition**: Declares an interface (contract).
- **Syntax**: `interface Service { void execute(); }`
- **Example**:
```java
interface Runner { void run(); }
public class Main {
    public static void main(String[] args) {
        Runner r = () -> System.out.println("Run");
        r.run();
    }
}
```
- **Interview Note**: Implicitly abstract and public methods (before Java 8/9 features).

#### `native`
- **Definition**: Indicates method is implemented in platform-dependent code (C/C++).
- **Syntax**: `public native void fastCalc();`
- **Example**:
```java
class Loader {
    // native void load(); // Requires JNI setup to run
}
public class Main { public static void main(String[] args) { System.out.println("JNI required for native"); } }
```
- **Interview Note**: Used in JNI (Java Native Interface). Rare in standard dev.

#### `new`
- **Definition**: Creates new objects.
- **Syntax**: `Object o = new Object();`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        String s = new String("Hi");
        System.out.println(s);
    }
}
```
- **Interview Note**: Triggers memory allocation on Heap.

#### `static`
- **Definition**: Belongs to class rather than instance.
- **Syntax**: `static int count;`
- **Example**:
```java
class Counter { static int count = 0; Counter() { count++; } }
public class Main {
    public static void main(String[] args) {
        new Counter(); new Counter();
        System.out.println(Counter.count); // 2
    }
}
```
- **Interview Note**: Static methods cannot access instance members directly.

#### `strictfp`
- **Definition**: Restricts floating-point calculations to ensure portability (IEEE 754).
- **Syntax**: `strictfp class MathCalc {}`
- **Example**:
```java
strictfp class Calc {
    double sum(double a, double b) { return a + b; }
}
public class Main { public static void main(String[] args) { System.out.println(new Calc().sum(0.1, 0.2)); } }
```
- **Interview Note**: Largely redundant since Java 17 as strict semantics are now restored by default.

#### `synchronized`
- **Definition**: locks object/class for thread safety.
- **Syntax**: `synchronized void lock() {}`
- **Example**:
```java
class Counter {
    int count = 0;
    synchronized void inc() { count++; }
}
public class Main {
    public static void main(String[] args) { new Counter().inc(); }
}
```
- **Interview Note**: Critical for concurrency; incurs performance overhead.

#### `transient`
- **Definition**: Skips variable during serialization.
- **Syntax**: `transient int cache;`
- **Example**:
```java
import java.io.*;
class User implements Serializable {
    String name = "Me";
    transient String pass = "Secret";
}
public class Main { public static void main(String[] args) { System.out.println(new User().name); } }
```
- **Interview Note**: Essential security feature for serialization.

#### `volatile`
- **Definition**: Indicates variable may be modified by threads; ensures visibility.
- **Syntax**: `volatile boolean running = true;`
- **Example**:
```java
class Flag { volatile boolean active = true; }
public class Main { public static void main(String[] args) { System.out.println(new Flag().active); } }
```
- **Interview Note**: Guarantees visibility ("happens-before"), but NOT atomicity.

### Flow Control

#### `break`
- **Definition**: Exits loop or switch.
- **Syntax**: `break;`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        for(int i=0; i<5; i++) { if(i==2) break; System.out.print(i); }
    }
}
```
- **Interview Note**: Can be used with labels to break outer loops (`break label;`).

#### `case`
- **Definition**: Branch in switch statement.
- **Syntax**: `case 1: ...`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        int x = 1;
        switch(x) { case 1: System.out.println("One"); break; }
    }
}
```
- **Mistake**: Forgetting `break` causes fall-through.

#### `continue`
- **Definition**: Skips current iteration of loop.
- **Syntax**: `continue;`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        for(int i=0; i<3; i++) { if(i==1) continue; System.out.print(i); } // 02
    }
}
```
- **Interview Note**: Useful for skipping bad data in loops without nesting `if`s.

#### `default`
- **Definition**: Default branch in switch or default method in interface.
- **Syntax**: `default void method() {}`
- **Example**:
```java
interface I { default void m() { System.out.println("Def"); } }
public class Main implements I { public static void main(String[] args) { new Main().m(); } }
```
- **Interview Note**: Added in Java 8 to Interfaces to allow backward compatibility.

#### `do`
- **Definition**: Starts a do-while loop (executed at least once).
- **Syntax**: `do { ... } while(cond);`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        int i=0; do { System.out.print(i++); } while(i<2);
    }
}
```
- **Interview Note**: Check condition is at the end.

#### `else`
- **Definition**: Alternative branch in if statement.
- **Syntax**: `if (cond) {} else {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        if(false) System.out.print("T"); else System.out.print("F");
    }
}
```

#### `for`
- **Definition**: Loop control.
- **Syntax**: `for(init; cond; update) {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) { for(int i=0; i<2; i++) System.out.print(i); }
}
```
- **Interview Note**: Enhanced for-loop (`for(T t : list)`) relies on `Iterable`.

#### `if`
- **Definition**: Conditional branch.
- **Syntax**: `if (cond) {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) { if(true) System.out.println("True"); }
}
```

#### `instanceof`
- **Definition**: Tests if object is instance of class/interface.
- **Syntax**: `obj instanceof Class`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        String s = "Hi";
        System.out.println(s instanceof String); // true
    }
}
```
- **Interview Note**: Returns false for `null` (does not throw NPE).

#### `return`
- **Definition**: Returns from method.
- **Syntax**: `return value;`
- **Example**:
```java
public class Main {
    static int get() { return 5; }
    public static void main(String[] args) { System.out.println(get()); }
}
```

#### `switch`
- **Definition**: Selects code to run based on value.
- **Syntax**: `switch(val) { ... }`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        int v = 1; switch(v) { case 1: System.out.println("1"); }
    }
}
```
- **Interview Note**: Supports int, char, String, enum. Pattern matching added in recent Java.

#### `while`
- **Definition**: Loop while condition true.
- **Syntax**: `while (cond) {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) { int i=0; while(i<2) System.out.print(i++); }
}
```

### Error Handling

#### `assert`
- **Definition**: Debugging aid; throws Error if false.
- **Syntax**: `assert cond : msg;`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        // Run with -ea flag
        int x = -1; assert x > 0 : "Must be positive";
    }
}
```
- **Interview Note**: Disabled by default. Enable with `-ea`. Do not use for production logic.

#### `catch`
- **Definition**: Catches exception in try-catch.
- **Syntax**: `catch(Exception e) {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        try { int i=1/0; } catch(Exception e) { System.out.println("Caught"); }
    }
}
```

#### `finally`
- **Definition**: Block always executed after try/catch.
- **Syntax**: `finally {}`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        try { System.out.print("Try"); } finally { System.out.print("Finally"); }
    }
}
```
- **Interview Note**: Runs even if return/throw occurs (unless `System.exit()` called).

#### `throw`
- **Definition**: Throws an exception instance.
- **Syntax**: `throw new Ex();`
- **Example**:
```java
public class Main {
    public static void main(String[] args) {
        try { throw new RuntimeException("Test"); } catch(Exception e) { System.out.println(e.getMessage()); }
    }
}
```

#### `throws`
- **Definition**: Declares exceptions method can throw.
- **Syntax**: `void m() throws IOException {}`
- **Example**:
```java
public class Main {
    static void m() throws Exception { throw new Exception(); }
    public static void main(String[] args) { try { m(); } catch(Exception e) {} }
}
```
- **Mistake**: Confusing `throw` (action) vs `throws` (declaration).

### Primitives & Void

#### `boolean`
- **Definition**: True or false values.
- **Example**: `boolean b = true;`

#### `byte`
- **Definition**: 8-bit integer (-128 to 127).
- **Example**: `byte b = 100;`

#### `char`
- **Definition**: 16-bit Unicode character.
- **Example**: `char c = 'A';`

#### `double`
- **Definition**: 64-bit floating point.
- **Example**: `double d = 3.14;`

#### `float`
- **Definition**: 32-bit floating point.
- **Example**: `float f = 3.14f;`

#### `int`
- **Definition**: 32-bit integer.
- **Example**: `int i = 100;`

#### `long`
- **Definition**: 64-bit integer.
- **Example**: `long l = 100000L;`

#### `short`
- **Definition**: 16-bit integer.
- **Example**: `short s = 1000;`

#### `void`
- **Definition**: No return value.
- **Syntax**: `void method() {}`
- **Example**: `public static void main(String[] a) {}`

### Other Reserved Words

#### `const`
- **Definition**: Reserved but unused. Use `final` instead.
- **Example**: Not valid Java.

#### `goto`
- **Definition**: Reserved but unused.
- **Example**: Not valid Java.

#### `enum`
- **Definition**: Defines a set of constants.
- **Syntax**: `enum Color { R, G, B }`
- **Example**:
```java
enum Day { MON, FRI }
public class Main {
    public static void main(String[] args) { System.out.println(Day.MON); }
}
```
- **Interview Note**: Enums are full classes in Java (maintain state, methods).

#### `import`
- **Definition**: Imports classes/packages.
- **Syntax**: `import java.util.*;`
- **Example**: `import java.util.List;`

#### `package`
- **Definition**: Defines namespace for class.
- **Syntax**: `package com.example;`
- **Example**: `package mypkg; class A {}`

#### `super`
- **Definition**: Refers to parent class instance.
- **Syntax**: `super.method()` constructor `super()`
- **Example**:
```java
class A { A() { System.out.print("A"); } }
class B extends A { B() { super(); System.out.print("B"); } }
public class Main { public static void main(String[] a) { new B(); } }
```

#### `this`
- **Definition**: Refers to current instance.
- **Syntax**: `this.field`
- **Example**:
```java
class A { int x; A(int x) { this.x = x; } }
```
- **Interview Note**: Can invoke other constructors `this()`.

---

## 2. Java Collections Framework

### Implementations

#### `ArrayList`
- **Internal**: Dynamic array. Resizes (50% growth) when full.
- **Features**: Fast access (O(1)), slow insertion/removal in middle (O(n)).
- **Syntax**: `List<String> list = new ArrayList<>();`
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("A"); list.add("B");
        System.out.println(list.get(0)); // A
    }
}
```
- **Common Methods**: `add(E)`, `get(int)`, `remove(int)`, `size()`.

#### `LinkedList`
- **Internal**: Doubly linked list.
- **Features**: Fast insertion/removal at ends (O(1)), slow access (O(n)).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        LinkedList<String> list = new LinkedList<>();
        list.addFirst("Start");
        list.addLast("End");
        System.out.println(list);
    }
}
```
- **Common Methods**: `addFirst()`, `addLast()`, `removeFirst()`.

#### `HashSet`
- **Internal**: HashMap (keys are elements, value is dummy object). uses `hashCode()`.
- **Features**: Unique elements, unordered. Access O(1).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<Integer> set = new HashSet<>();
        set.add(1); set.add(1); // Ignored
        System.out.println(set.size()); // 1
    }
}
```
- **Common Methods**: `add()`, `contains()`, `remove()`.

#### `TreeSet`
- **Internal**: Red-Black Tree (TreeMap).
- **Features**: Sorted, unique and ordered elements. O(log n).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        TreeSet<Integer> set = new TreeSet<>();
        set.add(5); set.add(1);
        System.out.println(set); // [1, 5]
    }
}
```
- **Common Methods**: `first()`, `last()`, `subSet()`.

#### `HashMap`
- **Internal**: Array of Nodes (buckets). Uses hashing + equals. Since Java 8, bucket becomes Tree if collisions > 8.
- **Features**: Key-Value. Allows 1 null key. Unordered. O(1).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new HashMap<>();
        map.put("A", 1);
        System.out.println(map.get("A"));
    }
}
```
- **Common Methods**: `put()`, `get()`, `containsKey()`, `keySet()`.

#### `TreeMap`
- **Internal**: Red-Black Tree.
- **Features**: Sorted keys. O(log n).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> map = new TreeMap<>();
        map.put("B", 2); map.put("A", 1);
        System.out.println(map.keySet()); // [A, B]
    }
}
```

#### `PriorityQueue`
- **Internal**: Priority Heap.
- **Features**: Ordered by priority (natural or comparator). Head is least element. O(log n).
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Queue<Integer> pq = new PriorityQueue<>();
        pq.add(5); pq.add(1); pq.add(3);
        System.out.println(pq.poll()); // 1 (Least)
    }
}
```
- **Common Methods**: `add()`, `poll()`, `peek()`.

#### `ArrayDeque`
- **Internal**: Resizable array (circular buffer).
- **Features**: Double-ended queue. Faster than LinkedList for queue/stack.
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Deque<String> stack = new ArrayDeque<>();
        stack.push("A"); stack.push("B");
        System.out.println(stack.pop()); // B
    }
}
```

---

## 3. Functional Interfaces
- **Concept**: Interface with exactly one abstract method.

#### `Function<T, R>`
- **Def**: Takes T, returns R. `R apply(T t)`
- **Example**:
```java
import java.util.function.Function;
public class Main {
    public static void main(String[] args) {
        Function<Integer, String> stringify = i -> "Val: " + i;
        System.out.println(stringify.apply(10));
    }
}
```

#### `Predicate<T>`
- **Def**: Takes T, returns boolean. `boolean test(T t)`
- **Example**:
```java
import java.util.function.Predicate;
public class Main {
    public static void main(String[] args) {
        Predicate<String> isEmpty = s -> s.isEmpty();
        System.out.println(isEmpty.test(""));
    }
}
```

#### `Consumer<T>`
- **Def**: Takes T, returns void. `void accept(T t)`
- **Example**:
```java
import java.util.function.Consumer;
public class Main {
    public static void main(String[] args) {
        Consumer<String> print = s -> System.out.println(s);
        print.accept("Hello");
    }
}
```

#### `Supplier<T>`
- **Def**: Takes nothing, returns T. `T get()`
- **Example**:
```java
import java.util.function.Supplier;
public class Main {
    public static void main(String[] args) {
        Supplier<Double> random = () -> Math.random();
        System.out.println(random.get());
    }
}
```

#### `BiFunction<T, U, R>`
- **Def**: Takes T and U, returns R.
```java
BiFunction<Integer, Integer, Integer> add = (a, b) -> a + b;
```

---

## 4. Stream API

#### `filter`
- **Def**: Keeps elements keeping predicate true.
- **Example**:
```java
import java.util.*;
public class Main {
    public static void main(String[] a) {
        Arrays.asList(1, 2, 3).stream().filter(x -> x > 1).forEach(System.out::print); // 23
    }
}
```

#### `map`
- **Def**: Transforms elements.
- **Example**:
```java
Arrays.asList("a", "b").stream().map(s -> s.toUpperCase()).forEach(System.out::print); // AB
```

#### `flatMap`
- **Def**: Flattens nested structures.
- **Example**:
```java
List<List<Integer>> nums = Arrays.asList(Arrays.asList(1), Arrays.asList(2));
nums.stream().flatMap(List::stream).forEach(System.out::print); // 12
```

#### `reduce`
- **Def**: Aggregates elements.
- **Example**:
```java
int sum = Arrays.asList(1, 2, 3).stream().reduce(0, Integer::sum);
System.out.println(sum); // 6
```

#### `sorted`, `distinct`, `limit`, `skip`
```java
list.stream().sorted().distinct().limit(5).skip(1)...
```

#### `anyMatch`, `allMatch`, `noneMatch`
- **Def**: Short-circuiting boolean checks.
- **Example**:
```java
boolean hasEven = list.stream().anyMatch(n -> n % 2 == 0);
```

#### `findFirst`, `findAny`
- **Def**: Returns Optional describing an element.
- **Example**:
```java
Optional<String> val = list.stream().findFirst();
```

#### `count`, `min`, `max`
- **Def**: Aggregation terminal operations.
- **Example**:
```java
long count = list.stream().count();
int min = list.stream().min(Integer::compare).orElse(0);
```

#### `forEach`
- **Def**: Iterates over elements.
- **Example**: `list.stream().forEach(System.out::println);`

---

## 5. Collectors

#### `groupingBy`
- **Def**: Groups elements by classifier.
- **Example**:
```java
import java.util.stream.*; import java.util.*;
public class Main {
    public static void main(String[] a) {
        List<String> items = Arrays.asList("apple", "apricot", "banana");
        Map<Character, List<String>> grp = items.stream()
            .collect(Collectors.groupingBy(s -> s.charAt(0)));
        System.out.println(grp); // {a=[apple, apricot], b=[banana]}
    }
}
```

#### `partitioningBy`
- **Def**: Groups into true/false lists.
- **Example**: `collect(Collectors.partitioningBy(s -> s.length() > 3))`

#### `joining`
- **Def**: Joins strings.
- **Example**: `collect(Collectors.joining(", "))` -> "a, b, c"

#### `toList`, `toSet`, `toMap`
- **Def**: Accumulates into collections.
- **Example**:
```java
List<String> list = stream.collect(Collectors.toList());
Set<String> set = stream.collect(Collectors.toSet());
Map<String, Integer> map = stream.collect(Collectors.toMap(s -> s, String::length));
```

#### `counting`, `summarizingInt`
- **Def**: Statistical collectors.
- **Example**:
```java
Long count = stream.collect(Collectors.counting());
IntSummaryStatistics stats = stream.collect(Collectors.summarizingInt(String::length));
```

#### `mapping`
- **Def**: Adapts the collector.
- **Example**:
```java
collect(Collectors.groupingBy(s -> s.charAt(0), Collectors.mapping(String::toUpperCase, Collectors.toList())));
```

---

## 6. Generics

#### Class Example
```java
class Box<T> {
    private T t;
    public void set(T t) { this.t = t; }
    public T get() { return t; }
}
```

#### Wildcards & Bounds
- **`<?>`**: Any type.
- **`<? extends Number>`**: Upper bound (Number or subclass). READ only (PECS - Producer Extends).
- **`<? super Integer>`**: Lower bound (Integer or superclass). WRITE allowed (Consumer Super).

#### Type Erasure
- **Concept**: Generics meant for compile-time safety. Runtime byte code has no generic info (`Box<String>` becomes `Box`, T becomes Object).

---

## 7. Concurrency

#### `synchronized` Block
```java
public void add(int value) {
    synchronized(this) {
        this.count += value;
    }
}
```

#### `ConcurrentHashMap`
- **Def**: Thread-safe hash map, allows concurrent reads/writes without full lock. Uses partial locking (buckets).
- **Example**:
```java
import java.util.concurrent.*;
public class Main {
    public static void main(String[] args) {
        ConcurrentMap<String, String> map = new ConcurrentHashMap<>();
        map.put("key", "value");
        map.putIfAbsent("key", "new"); // Atomic check-then-act
    }
}
```
