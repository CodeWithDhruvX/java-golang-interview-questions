# Java OOP Deep Dive — Inheritance, Polymorphism & Interface Gotchas

> **Topics:** Inheritance, Abstract Classes, Interfaces, Polymorphism, Method Resolution, Casting, Generics, Inner Classes, Enums

---

## 📋 Reading Progress

- [ ] **Section 1:** Inheritance & Method Overriding (Q1–Q18)
- [ ] **Section 2:** Abstract Classes & Interfaces (Q19–Q35)
- [ ] **Section 3:** Casting & Generics (Q36–Q50)
- [ ] **Section 4:** Inner Classes, Anonymous Classes & Enums (Q51–Q70)
- [ ] **Section 5:** Object Class Deep Dive (Q71–Q80)

> 🔖 **Last read:** <!-- e.g. Q18 · Section 1 done -->

---

## Section 1: Inheritance & Method Overriding (Q1–Q18)

### 1. Method Resolution Order
**Q: What is the output?**
```java
class A { String name() { return "A"; } }
class B extends A { @Override String name() { return "B"; } }
class C extends B {} // no override

public class Main {
    public static void main(String[] args) {
        A obj = new C();
        System.out.println(obj.name());
    }
}
```
**A:** `B`. Runtime type is `C`, which inherits `name()` from `B` (closest override). Dynamic dispatch walks up the hierarchy.

---

### 2. super.method() Access
**Q: What is the output?**
```java
class Animal { String sound() { return "..."; } }
class Dog extends Animal {
    @Override
    String sound() { return super.sound() + " Woof"; }
}

public class Main { public static void main(String[] args) { System.out.println(new Dog().sound()); } }
```
**A:** `... Woof`. `super.sound()` calls the parent implementation from within the overriding method.

---

### 3. Overriding Cannot Reduce Visibility
**Q: Does this compile?**
```java
class Parent { public void show() {} }
class Child extends Parent {
    @Override
    private void show() {} // reducing public → private
}
```
**A:** **Compile Error.** Overriding methods cannot have a more restrictive access modifier. You can only widen (protected → public).

---

### 4. Overriding Cannot Throw Broader Checked Exceptions
**Q: Does this compile?**
```java
import java.io.*;
class Parent { void read() throws IOException {} }
class Child extends Parent {
    @Override
    void read() throws Exception {} // Exception is broader than IOException
}
```
**A:** **Compile Error.** An overriding method can throw only the same or more specific checked exceptions, never broader ones.

---

### 5. Static Method Hiding vs Overriding
**Q: What is the output?**
```java
class P { static void greet() { System.out.println("P static"); } }
class C extends P { static void greet() { System.out.println("C static"); } }

public class Main {
    public static void main(String[] args) {
        P obj = new C();
        obj.greet(); // compile-time type decides — method hiding
    }
}
```
**A:** `P static`. Static methods are **hidden**, not overridden. Resolved at compile time based on declared type of `obj`.

---

### 6. Calling Overridden Method from Super Constructor
**Q: What is the output?**
```java
class Base {
    Base() { display(); }
    void display() { System.out.println("Base.display"); }
}
class Derived extends Base {
    int value = 10;
    @Override
    void display() { System.out.println("Derived.display: " + value); }
}

public class Main { public static void main(String[] args) { new Derived(); } }
```
**A:** `Derived.display: 0`. **Critical trap!** The super constructor calls `display()` which dispatches to `Derived.display()`. But `Derived`'s fields haven't been initialized yet, so `value` is `0`. Never call overridable methods from constructors.

---

### 7. Private Methods Are Not Overridden
**Q: What is the output?**
```java
class Base {
    private void secret() { System.out.println("Base secret"); }
    void callSecret() { secret(); }
}
class Child extends Base {
    void secret() { System.out.println("Child secret"); } // NOT an override
}

public class Main { public static void main(String[] args) { new Child().callSecret(); } }
```
**A:** `Base secret`. Private methods aren't part of the inheritance contract. `Child.secret()` is a brand-new method.

---

### 8. Constructor Not Inherited — Must Call super()
**Q: Does this compile?**
```java
class Animal { Animal(String name) { System.out.println("Animal: " + name); } }
class Dog extends Animal {
    Dog() {} // implicit super() — but no no-arg constructor in Animal!
}
```
**A:** **Compile Error.** Compiler inserts implicit `super()` (no-arg), but `Animal` only has `Animal(String)`. Must call `super("some name")` explicitly.

---

### 9. Field Hiding in Inheritance
**Q: What is the output?**
```java
class Parent { int x = 1; }
class Child extends Parent { int x = 2; }

public class Main {
    public static void main(String[] args) {
        Parent obj = new Child();
        System.out.println(obj.x);          // declared type = Parent
        System.out.println(((Child)obj).x); // cast to Child
    }
}
```
**A:**
```
1
2
```
Fields are **not polymorphic** — resolved at compile time based on the declared type of the reference.

---

### 10. @Override Catches Typos
**Q: What is the output?**
```java
class Animal { void speak() { System.out.println("..."); } }
class Cat extends Animal {
    void Speak() { System.out.println("Meow"); } // typo: capital S
}

public class Main {
    public static void main(String[] args) { new Animal() {{ }}.speak(); Animal a = new Cat(); a.speak(); }
}
```
**A:** `...`. `Speak()` is NOT an override — it's a new method. `Animal.speak()` runs. Add `@Override` and the compiler catches the typo immediately.

---

### 11. Covariant Return Type
**Q: Does this compile?**
```java
class Builder {
    Builder setName(String name) { return this; }
}
class AdvancedBuilder extends Builder {
    @Override
    AdvancedBuilder setName(String name) { return this; } // covariant return
}
```
**A:** **Yes, compiles.** An overriding method can return a subtype of the original return type (Java 5+). Enables fluent builder APIs.

---

### 12. Multiple Interface Implementation
**Q: What is the output?**
```java
interface Flyable { void fly(); }
interface Swimmable { void swim(); }
class Duck implements Flyable, Swimmable {
    public void fly() { System.out.println("flying"); }
    public void swim() { System.out.println("swimming"); }
}

public class Main { public static void main(String[] args) { Duck d = new Duck(); d.fly(); d.swim(); } }
```
**A:**
```
flying
swimming
```
A class can implement multiple interfaces.

---

### 13. Diamond Default Method Resolution
**Q: What is the output?**
```java
interface A { default String greet() { return "A"; } }
interface B extends A { default String greet() { return "B"; } }
class C implements A, B {} // which default wins?

public class Main { public static void main(String[] args) { System.out.println(new C().greet()); } }
```
**A:** `B`. Rule: more-specific interface wins. `B` extends `A`, so `B`'s default is more specific.

---

### 14. Diamond Ambiguity — Compile Error
**Q: Does this compile?**
```java
interface A { default void hello() { System.out.println("A"); } }
interface B { default void hello() { System.out.println("B"); } }
class C implements A, B {} // neither is more specific — ambiguous!
```
**A:** **Compile Error.** `C` must explicitly override `hello()` to resolve the ambiguity between two unrelated interfaces.

---

### 15. Interface Static Method Not Inherited
**Q: Does this compile?**
```java
interface MathOps { static int square(int x) { return x * x; } }
class Calc implements MathOps {}

public class Main { public static void main(String[] args) { Calc.square(3); } }
```
**A:** **Compile Error.** Static interface methods must be called via the interface name: `MathOps.square(3)`. They are not inherited by implementing classes.

---

### 16. Abstract Class Can Have Constructor
**Q: What is the output?**
```java
abstract class Vehicle {
    String brand;
    Vehicle(String brand) { this.brand = brand; System.out.println("Vehicle: " + brand); }
}
class Car extends Vehicle {
    Car(String b) { super(b); System.out.println("Car: " + b); }
}

public class Main { public static void main(String[] args) { new Car("Toyota"); } }
```
**A:**
```
Vehicle: Toyota
Car: Toyota
```

---

### 17. Abstract Method Cannot Be private or static
**Q: Does this compile?**
```java
abstract class Broken {
    private abstract void doA();  // private abstract — ERROR
    static abstract void doB();   // static abstract — ERROR
}
```
**A:** **Compile Error.** Abstract methods cannot be `private` (unreachable to subclasses) or `static` (static methods can't be overridden).

---

### 18. Sealed Classes (Java 17+)
**Q: Does this compile?**
```java
sealed class Shape permits Circle, Rectangle {}
final class Circle extends Shape {}
final class Rectangle extends Shape {}
```
**A:** **Yes, compiles.** `sealed` restricts which classes can extend `Shape`. Non-permitted extensions cause a compile error. Enables exhaustive pattern matching.

---

## Section 2: Abstract Classes & Interfaces (Q19–Q35)

### 19. Interface Fields Are public static final
**Q: What is the output?**
```java
interface Config { int MAX = 100; }

public class Main { public static void main(String[] args) { System.out.println(Config.MAX); } }
```
**A:** `100`. All interface fields are implicitly `public static final`.

---

### 20. Interface Methods Are public abstract by Default
**Q: Does the implementing class need public?**
```java
interface Drawable { void draw(); }
class Canvas implements Drawable {
    void draw() {} // package-private — is this OK?
}
```
**A:** **Compile Error.** Interface method `draw()` is implicitly `public`. The override must be at least as accessible — `public` is required.

---

### 21. Abstract Class vs Interface — Instance State
**Q: Does this compile?**
```java
interface Counter {
    int count = 0; // implicitly static final — NOT mutable instance state!
    void increment();
}
```
**A:** **Compiles**, but `count` is a constant. You cannot use it as a mutable counter. Interfaces cannot have mutable instance fields — use abstract classes instead.

---

### 22. Functional Interface
**Q: What is the output?**
```java
@FunctionalInterface
interface Transformer { int transform(int x); }

public class Main {
    public static void main(String[] args) {
        Transformer doubler = x -> x * 2;
        System.out.println(doubler.transform(5));
    }
}
```
**A:** `10`. A `@FunctionalInterface` has exactly one abstract method and can be implemented with a lambda.

---

### 23. @FunctionalInterface Enforcement
**Q: Does this compile?**
```java
@FunctionalInterface
interface BadFI { void doA(); void doB(); }
```
**A:** **Compile Error.** Two abstract methods violate the functional interface contract.

---

### 24. Private Interface Methods (Java 9+)
**Q: Does this compile?**
```java
interface Logger {
    private void log(String msg) { System.out.println("[LOG] " + msg); }
    default void info(String msg) { log(msg); }
    default void error(String msg) { log(msg); }
}
```
**A:** **Yes, compiles** (Java 9+). Private interface methods allow code reuse between default methods without exposing the helper method.

---

### 25. Template Method Pattern
**Q: What is the output?**
```java
abstract class DataProcessor {
    final void process() { readData(); processData(); writeData(); }
    abstract void readData();
    abstract void processData();
    void writeData() { System.out.println("writing"); }
}
class CsvProcessor extends DataProcessor {
    void readData() { System.out.println("reading CSV"); }
    void processData() { System.out.println("processing CSV"); }
}

public class Main { public static void main(String[] args) { new CsvProcessor().process(); } }
```
**A:**
```
reading CSV
processing CSV
writing
```

---

### 26. Cannot Instantiate Abstract Class
**Q: Does this compile?**
```java
abstract class Base { void doSomething() { System.out.println("concrete!"); } }
public class Main { public static void main(String[] args) { new Base(); } }
```
**A:** **Compile Error.** A class declared `abstract` cannot be instantiated even if it has no abstract methods.

---

### 27. Subclass Must Implement All Abstract Methods
**Q: Does this compile?**
```java
abstract class Vehicle { abstract void move(); abstract void stop(); }
class Car extends Vehicle {
    @Override void move() { System.out.println("moving"); }
    // stop() not implemented!
}
```
**A:** **Compile Error.** `Car` must implement all abstract methods or be declared `abstract` itself.

---

### 28. Marker Interface
**Q: What is the output?**
```java
import java.io.*;
public class Main {
    static class MyData implements Serializable { int value = 42; }
    public static void main(String[] args) { System.out.println(new MyData() instanceof Serializable); }
}
```
**A:** `true`. `Serializable` is a **marker interface** — no methods, just marks eligibility for serialization.

---

### 29. Anonymous Class Implementing Interface
**Q: What is the output?**
```java
interface Greeter { void greet(); }
public class Main {
    public static void main(String[] args) {
        Greeter g = new Greeter() {
            @Override
            public void greet() { System.out.println("Hello from anon!"); }
        };
        g.greet();
    }
}
```
**A:** `Hello from anon!`. Anonymous classes create one-off implementations inline. Largely replaced by lambdas for functional interfaces.

---

### 30. Class Implementation Wins Over Default Method
**Q: What is the output?**
```java
interface Greeter { default String greet() { return "Hello from Interface"; } }
class CustomGreeter implements Greeter { @Override public String greet() { return "Hello from Class"; } }

public class Main { public static void main(String[] args) { System.out.println(new CustomGreeter().greet()); } }
```
**A:** `Hello from Class`. Class implementations always override default interface methods.

---

### 31. Interface Extending Multiple Interfaces
**Q: Does this compile?**
```java
interface Flyable { void fly(); }
interface Swimmable { void swim(); }
interface Duck extends Flyable, Swimmable { void quack(); }
class Mallard implements Duck {
    public void fly() {} public void swim() {} public void quack() {}
}
```
**A:** **Yes, compiles.** Interfaces can extend multiple interfaces. The implementing class must implement all inherited abstract methods.

---

### 32. super() Must Be First Statement in Constructor
**Q: Does this compile?**
```java
class Parent { Parent() {} }
class Child extends Parent {
    Child() {
        System.out.println("hello"); // code before super()
        super();
    }
}
```
**A:** **Compile Error.** `super()` or `this()` must be the very first statement in a constructor.

---

### 33. this() Constructor Chaining
**Q: What is the output?**
```java
class Point {
    int x, y;
    Point() { this(0, 0); }
    Point(int x, int y) { this.x = x; this.y = y; System.out.println("Point(" + x + ", " + y + ")"); }
}

public class Main { public static void main(String[] args) { new Point(); } }
```
**A:** `Point(0, 0)`. `this(0, 0)` chains to the two-arg constructor.

---

### 34. Cannot Have Circular this() Calls
**Q: Does this compile?**
```java
class Broken {
    Broken() { this(1); }
    Broken(int x) { this(); } // circular!
}
```
**A:** **Compile Error.** Circular constructor invocations are not allowed.

---

### 35. Interface Cannot Be Instantiated
**Q: Does this compile?**
```java
interface Runnable {}
public class Main { public static void main(String[] args) { new Runnable(); } }
```
**A:** **Compile Error.** Interfaces cannot be instantiated. You can instantiate anonymous classes implementing the interface.

---

## Section 3: Casting & Generics (Q36–Q50)

### 36. Upcasting — Always Safe & Implicit
**Q: Does this compile?**
```java
class Animal {}
class Dog extends Animal {}
public class Main {
    public static void main(String[] args) {
        Dog d = new Dog();
        Animal a = d; // upcasting — no cast needed
        System.out.println(a instanceof Animal);
    }
}
```
**A:** **Yes, compiles and prints** `true`. Upcasting is implicit and always safe.

---

### 37. Downcasting — ClassCastException
**Q: What happens at runtime?**
```java
class Animal {}
class Dog extends Animal {}
class Cat extends Animal {}
public class Main {
    public static void main(String[] args) {
        Animal a = new Dog();
        Cat c = (Cat) a; // wrong downcast
    }
}
```
**A:** **ClassCastException at runtime.** `a` holds a `Dog`. Casting to `Cat` fails. Use `instanceof` before downcasting.

---

### 38. Safe Downcast with Pattern Matching (Java 16+)
**Q: What is the output?**
```java
class Animal {}
class Dog extends Animal { String sound() { return "Woof"; } }
class Cat extends Animal { String sound() { return "Meow"; } }
public class Main {
    static void speak(Animal a) {
        if (a instanceof Dog d) System.out.println("Dog: " + d.sound());
        else if (a instanceof Cat c) System.out.println("Cat: " + c.sound());
    }
    public static void main(String[] args) { speak(new Dog()); speak(new Cat()); }
}
```
**A:**
```
Dog: Woof
Cat: Meow
```

---

### 39. Type Erasure — Same Class at Runtime
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> ints = new ArrayList<>();
        List<String> strs = new ArrayList<>();
        System.out.println(ints.getClass() == strs.getClass());
    }
}
```
**A:** `true`. Due to **type erasure**, generic type parameters are erased at runtime. Both are just `ArrayList` at runtime.

---

### 40. Wildcard Upper Bound — Reading
**Q: Does this compile?**
```java
import java.util.*;
public class Main {
    static double sum(List<? extends Number> list) {
        double total = 0;
        for (Number n : list) total += n.doubleValue();
        return total;
    }
    public static void main(String[] args) { System.out.println(sum(List.of(1, 2, 3))); }
}
```
**A:** **Compiles and prints** `6.0`. `? extends Number` = any subtype of Number. Read-safe but you cannot add to such a list.

---

### 41. Wildcard Lower Bound — Writing
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static void addInts(List<? super Integer> list) { list.add(1); list.add(2); }
    public static void main(String[] args) {
        List<Number> nums = new ArrayList<>();
        addInts(nums);
        System.out.println(nums);
    }
}
```
**A:** `[1, 2]`. `? super Integer` = Integer or any supertype (Number, Object). Safe to **write** Integer values (PECS: Consumer Super).

---

### 42. Generic Method
**Q: What is the output?**
```java
public class Main {
    static <T> T identity(T val) { return val; }
    public static void main(String[] args) {
        System.out.println(identity("hello"));
        System.out.println(identity(42));
    }
}
```
**A:**
```
hello
42
```

---

### 43. Bounded Type Parameter
**Q: What is the output?**
```java
public class Main {
    static <T extends Comparable<T>> T max(T a, T b) {
        return a.compareTo(b) >= 0 ? a : b;
    }
    public static void main(String[] args) {
        System.out.println(max(3, 7));
        System.out.println(max("apple", "banana"));
    }
}
```
**A:**
```
7
banana
```

---

### 44. Raw Types — Unchecked Warning
**Q: Does this compile and what is the risk?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List list = new ArrayList();  // raw type
        list.add("hello");
        list.add(42);
        String s = (String) list.get(1); // ClassCastException at runtime!
    }
}
```
**A:** **Compiles with unchecked warnings, ClassCastException at runtime** on the cast. Raw types bypass generic safety.

---

### 45. Cannot Create Generic Array
**Q: Does this compile?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String>[] arr = new List<String>[10]; // ERROR
    }
}
```
**A:** **Compile Error.** Generic array creation is not allowed due to type erasure — the JVM can't enforce the element type at runtime.

---

### 46. Generic Class — Pair
**Q: What is the output?**
```java
public class Main {
    static class Pair<A, B> {
        A first; B second;
        Pair(A f, B s) { first = f; second = s; }
        public String toString() { return "(" + first + ", " + second + ")"; }
    }
    public static void main(String[] args) {
        System.out.println(new Pair<>("hi", 42));
        System.out.println(new Pair<>(1, true));
    }
}
```
**A:**
```
(hi, 42)
(1, true)
```

---

### 47. Autoboxing Overload Resolution
**Q: What is the output?**
```java
public class Main {
    static void show(long x) { System.out.println("long"); }
    static void show(Integer x) { System.out.println("Integer"); }
    public static void main(String[] args) { show(5); }
}
```
**A:** `long`. Java prefers **widening** (int → long) over **autoboxing** (int → Integer) during overload resolution.

---

### 48. Overloading Resolved at Compile Time
**Q: What is the output?**
```java
public class Main {
    static void describe(Object o) { System.out.println("Object"); }
    static void describe(String s) { System.out.println("String"); }
    public static void main(String[] args) {
        Object o = "hello"; // declared as Object
        describe(o);        // resolved by declared type
    }
}
```
**A:** `Object`. Overloading is resolved at **compile time** by declared type. Contrast with overriding which is resolved at **runtime** by actual type.

---

### 49. Varargs — Always an Array
**Q: What is the output?**
```java
public class Main {
    static void inspect(Object... args) {
        System.out.println(args.getClass().getSimpleName());
        System.out.println(args.length);
    }
    public static void main(String[] args) { inspect("a", "b", "c"); }
}
```
**A:**
```
Object[]
3
```

---

### 50. Covariant Return Type — Enabler for Fluent API
**Q: Does this compile?**
```java
class Animal { Animal create() { return new Animal(); } }
class Dog extends Animal {
    @Override
    Dog create() { return new Dog(); }  // covariant return type
}
```
**A:** **Yes, compiles.** Overriding method can return a subtype of the declared return type.

---

## Section 4: Inner Classes, Anonymous Classes & Enums (Q51–Q68)

### 51. Inner Class Accessing Outer Members
**Q: What is the output?**
```java
public class Main {
    private int x = 10;
    class Inner { void display() { System.out.println("x = " + x); } }
    public static void main(String[] args) {
        Main outer = new Main();
        Inner inner = outer.new Inner();
        inner.display();
    }
}
```
**A:** `x = 10`. Non-static inner class holds an implicit reference to its outer instance and can access all its members.

---

### 52. Static Nested Class — No Outer Instance Needed
**Q: What is the output?**
```java
public class Main {
    static class Node { int val; Node(int v) { val = v; } }
    public static void main(String[] args) {
        Node n = new Node(5);
        System.out.println(n.val);
    }
}
```
**A:** `5`. Static nested classes don't need an outer class instance.

---

### 53. Anonymous Class — Effectively Final Variables
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        int multiplier = 3; // effectively final
        Runnable r = new Runnable() {
            public void run() { System.out.println(5 * multiplier); }
        };
        // multiplier = 4; // uncommenting causes compile error
        r.run();
    }
}
```
**A:** **Compiles and prints** `15`. Anonymous classes can only capture effectively final local variables.

---

### 54. Lambda this vs Anonymous Class this
**Q: What does `this` refer to in a lambda?**
```java
public class Main {
    int value = 42;
    void demo() {
        Runnable lambda = () -> System.out.println("value: " + value);
        // 'this' inside lambda refers to the enclosing Main instance
        lambda.run();
    }
    public static void main(String[] args) { new Main().demo(); }
}
```
**A:** `value: 42`. In lambdas, `this` refers to the **enclosing class instance** — unlike anonymous classes where `this` refers to the anonymous object itself.

---

### 55. Enum Comparison — == is Safe
**Q: What is the output?**
```java
public class Main {
    enum Day { MON, TUE, WED, THU, FRI, SAT, SUN }
    public static void main(String[] args) {
        Day d = Day.MON;
        System.out.println(d == Day.MON);
        System.out.println(d.name());
        System.out.println(d.ordinal());
    }
}
```
**A:**
```
true
MON
0
```
Enum constants are singletons. `==` and `.equals()` both work safely. `.name()` returns the string name, `.ordinal()` the index.

---

### 56. Enum with Fields and Methods
**Q: What is the output?**
```java
public class Main {
    enum Planet {
        MERCURY(3.303e+23, 2.4397e6),
        EARTH(5.976e+24, 6.37814e6);
        final double mass, radius;
        static final double G = 6.67300E-11;
        Planet(double m, double r) { mass = m; radius = r; }
        double gravity() { return G * mass / (radius * radius); }
    }
    public static void main(String[] args) { System.out.printf("%.2f%n", Planet.EARTH.gravity()); }
}
```
**A:** `9.80` (approximately). Enums can have fields, constructors, and methods.

---

### 57. Enum values() and valueOf()
**Q: What is the output?**
```java
public class Main {
    enum Color { RED, GREEN, BLUE }
    public static void main(String[] args) {
        for (Color c : Color.values()) System.out.print(c + " ");
        System.out.println();
        System.out.println(Color.valueOf("GREEN"));
    }
}
```
**A:**
```
RED GREEN BLUE
GREEN
```

---

### 58. Enum in switch
**Q: What is the output?**
```java
public class Main {
    enum Season { SPRING, SUMMER, FALL, WINTER }
    public static void main(String[] args) {
        Season s = Season.SUMMER;
        switch (s) {
            case SUMMER: System.out.println("Hot!"); break;
            case WINTER: System.out.println("Cold!"); break;
            default: System.out.println("Mild");
        }
    }
}
```
**A:** `Hot!`. In enum switch cases, you don't prefix with `Season.`, just use the constant name.

---

### 59. Enum Cannot Be Extended
**Q: Does this compile?**
```java
enum Base { A, B }
enum Extended extends Base { C, D }
```
**A:** **Compile Error.** Enums implicitly extend `java.lang.Enum` and cannot be subclassed.

---

### 60. Enum Implementing Interface
**Q: What is the output?**
```java
public class Main {
    interface Shape { double area(); }
    enum RegularShape implements Shape {
        CIRCLE { public double area() { return Math.PI; } },
        SQUARE { public double area() { return 1.0; } };
    }
    public static void main(String[] args) {
        System.out.printf("%.4f%n", RegularShape.CIRCLE.area());
    }
}
```
**A:** `3.1416`. Each enum constant can override abstract methods defined in the enum body.

---

### 61. Enum with Abstract Method
**Q: What is the output?**
```java
public class Main {
    enum Operation {
        PLUS { @Override int apply(int x, int y) { return x + y; } },
        MINUS { @Override int apply(int x, int y) { return x - y; } };
        abstract int apply(int x, int y);
    }
    public static void main(String[] args) {
        System.out.println(Operation.PLUS.apply(3, 2));
        System.out.println(Operation.MINUS.apply(3, 2));
    }
}
```
**A:**
```
5
1
```

---

### 62. EnumSet — Efficient Set for Enums
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    enum Day { MON, TUE, WED, THU, FRI, SAT, SUN }
    public static void main(String[] args) {
        EnumSet<Day> workdays = EnumSet.range(Day.MON, Day.FRI);
        System.out.println(workdays.size());
        System.out.println(workdays.contains(Day.SAT));
    }
}
```
**A:**
```
5
false
```

---

### 63. Builder Pattern with Static Nested Class
**Q: What is the output?**
```java
public class Main {
    static class Person {
        String name; int age;
        private Person() {}
        static class Builder {
            private Person p = new Person();
            Builder name(String n) { p.name = n; return this; }
            Builder age(int a) { p.age = a; return this; }
            Person build() { return p; }
        }
    }
    public static void main(String[] args) {
        Person p = new Person.Builder().name("Alice").age(30).build();
        System.out.println(p.name + ", " + p.age);
    }
}
```
**A:** `Alice, 30`. Classic builder pattern using a static nested class.

---

### 64. Singleton with Enum
**Q: What is the output?**
```java
public class Main {
    enum Singleton { INSTANCE; int value = 0; void inc() { value++; } }
    public static void main(String[] args) {
        Singleton.INSTANCE.inc();
        Singleton.INSTANCE.inc();
        System.out.println(Singleton.INSTANCE.value);
    }
}
```
**A:** `2`. Enum-based singleton is thread-safe, serialization-safe, and reflection-safe.

---

### 65. Local Class Captures Effectively Final
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String msg = "Hello";
        class Printer { void print() { System.out.println(msg); } }
        new Printer().print();
    }
}
```
**A:** `Hello`. Local classes (defined inside methods) can access effectively final local variables.

---

### 66. Inner Class Cannot Have Static Members (Pre-Java 16)
**Q: Does this compile in Java 8?**
```java
public class Main {
    class Inner {
        static int count = 0; // error in Java 8
    }
}
```
**A:** **Compile Error in Java 8–15.** Non-static inner classes cannot declare static members (except `static final` constants). Java 16+ allows this.

---

### 67. Anonymous Class — Multiple Interface Not Possible
**Q: Does this compile?**
```java
interface A { void doA(); }
interface B { void doB(); }
public class Main {
    public static void main(String[] args) {
        Object o = new A() { // only one interface!
            public void doA() {}
        };
    }
}
```
**A:** **Compiles.** But an anonymous class can only implement **one** interface or extend **one** class. For multiple interfaces, use a named class.

---

### 68. Sealed + Record + Switch (Java 17+)
**Q: What is the output?**
```java
public class Main {
    sealed interface Shape permits Circle, Rectangle {}
    record Circle(double r) implements Shape {}
    record Rectangle(double w, double h) implements Shape {}

    static double area(Shape s) {
        return switch (s) {
            case Circle c    -> Math.PI * c.r() * c.r();
            case Rectangle r -> r.w() * r.h();
        };
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

---

## Section 5: Object Class Deep Dive (Q69–Q80)

### 69. hashCode Contract — Must Override Both
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class Item {
        String name;
        Item(String n) { name = n; }
        @Override public boolean equals(Object o) {
            return o instanceof Item && name.equals(((Item)o).name);
        }
        // hashCode NOT overridden!
    }
    public static void main(String[] args) {
        Map<Item, Integer> map = new HashMap<>();
        map.put(new Item("apple"), 1);
        System.out.println(map.get(new Item("apple")));
    }
}
```
**A:** `null`. Without `hashCode()`, two logically equal `Item` objects land in different hash buckets. Always override `hashCode()` when overriding `equals()`.

---

### 70. Objects.equals() — Null Safe
**Q: What is the output?**
```java
import java.util.Objects;
public class Main {
    public static void main(String[] args) {
        String s = null;
        System.out.println(Objects.equals(s, null));
        System.out.println(Objects.equals(s, "hello"));
        System.out.println(Objects.equals("hi", "hi"));
    }
}
```
**A:**
```
true
false
true
```

---

### 71. compareTo Contract — Comparable
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    record Temp(double degrees) implements Comparable<Temp> {
        public int compareTo(Temp other) { return Double.compare(degrees, other.degrees); }
    }
    public static void main(String[] args) {
        List<Temp> list = new ArrayList<>(List.of(new Temp(37.5), new Temp(36.1), new Temp(38.9)));
        Collections.sort(list);
        list.forEach(t -> System.out.print(t.degrees() + " "));
    }
}
```
**A:** `36.1 37.5 38.9 `

---

### 72. Comparator vs Comparable
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    record Product(String name, double price) {}
    public static void main(String[] args) {
        List<Product> products = new ArrayList<>(List.of(
            new Product("Apple", 1.5), new Product("Banana", 0.5), new Product("Cherry", 2.5)));
        products.sort(Comparator.comparingDouble(Product::price));
        products.forEach(p -> System.out.print(p.name() + " "));
    }
}
```
**A:** `Banana Apple Cherry `. `Comparator` defines custom external ordering without modifying the class.

---

### 73. String Pool and intern()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = new String("hello").intern();
        String s2 = "hello";
        System.out.println(s1 == s2);
    }
}
```
**A:** `true`. `intern()` places the string in the pool and returns the pooled reference.

---

### 74. Optional — Avoid Null Returns
**Q: What is the output?**
```java
import java.util.Optional;
public class Main {
    static Optional<String> find(boolean found) {
        return found ? Optional.of("result") : Optional.empty();
    }
    public static void main(String[] args) {
        System.out.println(find(true).orElse("default"));
        System.out.println(find(false).orElse("default"));
    }
}
```
**A:**
```
result
default
```

---

### 75. Optional.map Chaining
**Q: What is the output?**
```java
import java.util.Optional;
public class Main {
    public static void main(String[] args) {
        String result = Optional.of("  alice  ")
                .map(String::trim)
                .map(String::toUpperCase)
                .orElse("UNKNOWN");
        System.out.println(result);
    }
}
```
**A:** `ALICE`

---

### 76. Objects.requireNonNull()
**Q: What is the output?**
```java
import java.util.Objects;
public class Main {
    static class Account {
        String owner;
        Account(String owner) { this.owner = Objects.requireNonNull(owner, "owner cannot be null"); }
    }
    public static void main(String[] args) {
        try { new Account(null); }
        catch (NullPointerException e) { System.out.println("Caught: " + e.getMessage()); }
    }
}
```
**A:** `Caught: owner cannot be null`

---

### 77. clone() — Shallow Copy Warning
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class Container implements Cloneable {
        List<Integer> items;
        Container(List<Integer> items) { this.items = items; }
        @Override protected Object clone() throws CloneNotSupportedException { return super.clone(); }
    }
    public static void main(String[] args) throws CloneNotSupportedException {
        Container orig = new Container(new ArrayList<>(List.of(1, 2, 3)));
        Container copy = (Container) orig.clone();
        copy.items.add(4);
        System.out.println(orig.items);
    }
}
```
**A:** `[1, 2, 3, 4]`. `super.clone()` is shallow — the `items` list is shared between original and clone.

---

### 78. getClass() vs instanceof
**Q: What is the output?**
```java
public class Main {
    static class Animal {}
    static class Dog extends Animal {}
    public static void main(String[] args) {
        Animal a = new Dog();
        System.out.println(a instanceof Animal);          // IS-A check
        System.out.println(a.getClass() == Animal.class); // exact class match
        System.out.println(a.getClass() == Dog.class);    // runtime class
    }
}
```
**A:**
```
true
false
true
```

---

### 79. Record — Immutable Data Carrier (Java 16+)
**Q: What is the output?**
```java
public class Main {
    record Point(int x, int y) {
        Point { if (x < 0 || y < 0) throw new IllegalArgumentException("negative"); }
    }
    public static void main(String[] args) {
        Point p = new Point(3, 4);
        System.out.println(p.x() + ", " + p.y());
        System.out.println(p);
        try { new Point(-1, 0); } catch (IllegalArgumentException e) { System.out.println("Caught: " + e.getMessage()); }
    }
}
```
**A:**
```
3, 4
Point[x=3, y=4]
Caught: negative
```

---

### 80. finalize() — Deprecated, Avoid
**Q: Why is finalize() unreliable?**
```java
public class Main {
    static class Resource {
        @Override protected void finalize() { System.out.println("cleaning up"); }
    }
    public static void main(String[] args) throws Exception {
        new Resource();
        System.gc(); // not guaranteed to run
        Thread.sleep(100);
    }
}
```
**A:** **Unpredictable** — may or may not print. `finalize()` is deprecated (Java 9+). Use `try-with-resources` and `AutoCloseable` for deterministic resource cleanup.
