# 🏗️ 02 — OOP Concepts in Java
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- The four pillars: Encapsulation, Abstraction, Inheritance, Polymorphism
- Abstract classes vs Interfaces
- Method overloading vs overriding
- `super` and `this` keywords
- Constructors and constructor chaining
- `instanceof` and casting

---

## ❓ Most Asked Questions

### Q1. What are the four pillars of OOP?

```java
// 1. ENCAPSULATION — bind data + methods, hide internal details
public class BankAccount {
    private double balance;   // hidden
    public void deposit(double amount) {
        if (amount > 0) balance += amount;  // controlled access
    }
    public double getBalance() { return balance; }
}

// 2. ABSTRACTION — expose only necessary details
public abstract class Shape {
    abstract double area();   // contract, no implementation
    public void print() { System.out.println("Area: " + area()); }
}

// 3. INHERITANCE — reuse parent class behaviour
public class Circle extends Shape {
    double radius;
    Circle(double r) { this.radius = r; }
    @Override double area() { return Math.PI * radius * radius; }
}

// 4. POLYMORPHISM — same interface, different behaviour
Shape s1 = new Circle(5);
Shape s2 = new Rectangle(4, 6);
s1.area();  // Circle's implementation
s2.area();  // Rectangle's implementation
```

---

### 🎯 How to Explain in Interview

"The four pillars of OOP are the foundation of object-oriented design. Encapsulation is about bundling data and methods together while hiding internal details - like a capsule that protects its contents. Abstraction means exposing only what's necessary and hiding complexity - like a car dashboard where I just see the steering wheel and pedals, not the engine. Inheritance lets me reuse code by having child classes inherit from parent classes - like how a Dog inherits from Animal. Polymorphism means 'many forms' - the same interface can have different implementations, like how both Circle and Rectangle have an area() method but calculate it differently. Together, these principles help me write clean, maintainable, and reusable code."

---

### Q2. What is the difference between abstract class and interface?

```java
// Abstract class — partial implementation, single inheritance
abstract class Vehicle {
    String brand;
    int speed;
    
    Vehicle(String brand) { this.brand = brand; }  // can have constructors
    
    abstract void start();  // must be implemented by subclass
    
    void stop() { System.out.println(brand + " stopped"); }  // concrete method
}

// Interface — pure contract (Java 8+: can have default/static methods)
interface Flyable {
    int MAX_ALTITUDE = 35000;  // implicitly public static final
    
    void fly();  // implicitly public abstract
    
    default void land() {   // Java 8 default method
        System.out.println("Landing...");
    }
    
    static void checkWeather() {   // Java 8 static method
        System.out.println("Weather OK");
    }
}

// Class can implement MULTIPLE interfaces but extend only ONE class
class Drone extends Vehicle implements Flyable, Rechargeable {
    Drone() { super("DJI"); }
    @Override public void start() { System.out.println("Drone starting"); }
    @Override public void fly()   { System.out.println("Drone flying"); }
}
```

| Feature | Abstract Class | Interface |
|---------|---------------|-----------|
| Inheritance | Single | Multiple |
| Constructor | ✅ | ❌ |
| Instance fields | ✅ | ❌ (only constants) |
| Access modifiers | Any | public only |
| When to use | "is-a" with shared code | "can-do" capability |

---

### 🎯 How to Explain in Interview

"Abstract classes and interfaces are both about abstraction, but they serve different purposes. An abstract class is like a partial template - it can have some implemented methods and some abstract ones that subclasses must complete. It's perfect when I have an 'is-a' relationship with shared code. A class can only extend one abstract class. An interface is a pure contract - it defines what a class can do but doesn't provide implementation. Since Java 8, interfaces can have default methods, but they're still primarily about defining capabilities. The key advantage is that a class can implement multiple interfaces, which gives me more flexibility. I use abstract classes for shared implementation and interfaces for defining capabilities."

---

### Q3. What is method overloading vs overriding?

```java
// OVERLOADING — same name, different parameters (compile-time polymorphism)
class MathUtils {
    int add(int a, int b) { return a + b; }
    double add(double a, double b) { return a + b; }        // different types
    int add(int a, int b, int c) { return a + b + c; }      // different count

    // Return type alone does NOT differentiate overloads
    // int add(int a, int b) { return a+b; }   ❌ — same signature
}

// OVERRIDING — subclass redefines parent method (runtime polymorphism)
class Animal {
    String speak() { return "..."; }
}

class Dog extends Animal {
    @Override                   // annotation is optional but recommended
    String speak() { return "Woof!"; }
}

class Cat extends Animal {
    @Override
    String speak() { return "Meow!"; }
}

// Runtime dispatch — actual object type determines method called
Animal a = new Dog();
System.out.println(a.speak()); // "Woof!" — Dog's version called at runtime
```

| Feature | Overloading | Overriding |
|---------|------------|------------|
| Class | Same class | Parent + Child |
| Bound | Compile-time | Runtime |
| Parameters | Must differ | Must be same |
| Return type | Can differ | Must be same (or covariant) |
| `static` methods | Can overload | Cannot override (hiding) |

---

### 🎯 How to Explain in Interview

"Method overloading and overriding sound similar but are very different. Overloading is compile-time polymorphism - same method name but different parameters in the same class. It's like having multiple tools with the same name but different uses. The compiler decides which version to call based on the arguments. Overriding is runtime polymorphism - a subclass provides its own implementation of a parent class method. The actual method called is determined at runtime based on the object's type. Overriding helps me customize behavior in subclasses, while overloading gives me convenience methods with different parameter options."

---

### Q4. What does `super` do?

```java
class Employee {
    String name;
    double salary;

    Employee(String name, double salary) {
        this.name = name;
        this.salary = salary;
    }

    String details() {
        return name + " - $" + salary;
    }
}

class Manager extends Employee {
    String department;

    Manager(String name, double salary, String dept) {
        super(name, salary);        // ← call parent constructor (MUST be first line)
        this.department = dept;
    }

    @Override
    String details() {
        return super.details() + " [" + department + "]";  // ← call parent method
    }
}

Manager m = new Manager("Alice", 95000, "Engineering");
System.out.println(m.details()); // "Alice - $95000.0 [Engineering]"
```

---

### 🎯 How to Explain in Interview

"The `super` keyword is my bridge to the parent class. It has two main uses: calling the parent constructor and accessing parent methods. When I call `super()` in a constructor, it must be the first line - it's like saying 'before I do my own initialization, let me make sure my parent is properly set up'. I can also use `super.methodName()` to call the parent's version of a method, which is great when I want to extend functionality rather than completely replace it. This helps me avoid code duplication and maintain the parent-child relationship. It's especially useful in overriding scenarios where I want to add to the parent's behavior instead of completely replacing it."

---

### Q5. What is constructor chaining with `this()`?

```java
class Connection {
    String host;
    int port;
    int timeout;
    boolean ssl;

    // Most specific constructor
    Connection(String host, int port, int timeout, boolean ssl) {
        this.host = host;
        this.port = port;
        this.timeout = timeout;
        this.ssl = ssl;
    }

    // Chains to above — must be FIRST statement
    Connection(String host, int port) {
        this(host, port, 30, false);   // default timeout and no SSL
    }

    // Chains to above
    Connection(String host) {
        this(host, 5432);              // default port
    }

    // Default
    Connection() {
        this("localhost");             // default host
    }
}

Connection c = new Connection(); // host=localhost, port=5432, timeout=30, ssl=false
```

---

### 🎯 How to Explain in Interview

"Constructor chaining with `this()` is all about code reuse and maintainability. Instead of duplicating initialization logic across multiple constructors, I can have one 'master' constructor that does all the work, and other constructors chain to it using `this()`. This follows the DRY principle - Don't Repeat Yourself. The chained call must be the first statement in the constructor. This pattern is super useful for providing convenience constructors with default values. For example, I might have a detailed constructor that takes all parameters, and simpler ones that provide defaults and delegate to the detailed one. This makes my code cleaner and easier to maintain."

---

### Q6. What is the `instanceof` operator and pattern matching (Java 16+)?

```java
// Traditional instanceof + cast
Object obj = "Hello World";
if (obj instanceof String) {
    String s = (String) obj;
    System.out.println(s.length());
}

// Java 16 pattern matching — cleaner syntax
if (obj instanceof String s) {
    System.out.println(s.length());  // 's' is already cast and ready to use
}

// Use in polymorphic scenarios
void processShape(Shape shape) {
    if (shape instanceof Circle c) {
        System.out.println("Circle radius: " + c.radius);
    } else if (shape instanceof Rectangle r) {
        System.out.println("Rectangle area: " + r.width * r.height);
    }
}
```

---

### 🎯 How to Explain in Interview

"The `instanceof` operator is how I check an object's type at runtime. It returns true if the object is an instance of the specified class or any of its subclasses. This is essential for type-safe casting and polymorphic behavior. The cool thing about Java 16+ is pattern matching with instanceof - instead of the old pattern of checking and then casting, I can do both in one step: `if (obj instanceof String s)`. This not only reduces boilerplate but also eliminates the possibility of casting errors. It's particularly useful when processing collections of different types or implementing visitor patterns where I need to handle different object types differently."

---

### Q7. What is an inner class in Java?

```java
// Static nested class — does not need outer instance
class Outer {
    static class StaticNested {
        void display() { System.out.println("Static nested"); }
    }
}
Outer.StaticNested obj = new Outer.StaticNested();

// Inner class — needs outer instance, has access to outer's private members
class LinkedList {
    private int size;
    
    class Iterator {
        int index = 0;
        boolean hasNext() { return index < size; }  // accesses outer 'size'
    }
    
    Iterator iterator() { return new Iterator(); }
}

// Anonymous class — one-time use implementation
Runnable r = new Runnable() {
    @Override public void run() { System.out.println("Running"); }
};

// Lambda replaces anonymous functional interfaces (Java 8+)
Runnable r2 = () -> System.out.println("Running with lambda");
```

---

### 🎯 How to Explain in Interview

"Inner classes in Java are really interesting because they give me different ways to organize related functionality. Static nested classes don't need an outer class instance - they're like regular classes but namespaced. Inner classes do need an outer instance and can access its private members - perfect for iterators or event handlers. Anonymous classes are one-off implementations - great for listeners or callbacks. Since Java 8, many anonymous classes can be replaced with lambdas, which are much more concise. I use inner classes when the class is only meaningful in the context of its outer class - it helps with encapsulation and makes the code more readable by keeping related functionality together."

---

### Q8. What is the difference between `Comparable` and `Comparator`?

```java
// Comparable — natural ordering, implemented BY the class itself
class Student implements Comparable<Student> {
    String name;
    int grade;

    Student(String name, int grade) { this.name = name; this.grade = grade; }

    @Override
    public int compareTo(Student other) {
        return Integer.compare(this.grade, other.grade);  // sort by grade
    }
}

List<Student> students = new ArrayList<>();
Collections.sort(students);  // uses compareTo — natural order

// Comparator — external ordering, does not modify the class
Comparator<Student> byName  = Comparator.comparing(s -> s.name);
Comparator<Student> byGrade = Comparator.comparingInt(s -> s.grade);
Comparator<Student> byNameThenGrade = byName.thenComparingInt(s -> s.grade);

students.sort(byNameThenGrade);                     // by name, then grade
students.sort(Comparator.reverseOrder());            // reverse natural order
students.sort(byGrade.reversed());                   // reverse grade order
```

---

### 🎯 How to Explain in Interview

"Comparable and Comparator are both about sorting, but they serve different purposes. Comparable is for natural ordering - it's implemented by the class itself to define its default sorting behavior. When I implement Comparable, I'm saying 'this is how I should normally be sorted'. Comparator is for custom ordering - it's a separate class that lets me define different sorting strategies without modifying the original class. This is great when I need multiple sorting criteria or when I can't modify the class I want to sort. For example, I might sort Students by grade using Comparable, but use a Comparator to sort them by name for a specific report. The key difference is that Comparable is 'how I sort myself' while Comparator is 'how I want to sort others'."

---

### Q9. What is Object class and its key methods?

```java
// Every class implicitly extends Object
public class Product {
    private String name;
    private double price;

    Product(String name, double price) {
        this.name = name;
        this.price = price;
    }

    // equals — logical equality (override with hashCode together!)
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Product p)) return false;
        return Double.compare(p.price, price) == 0 && Objects.equals(name, p.name);
    }

    // hashCode — MUST override when equals is overridden (HashMap contract)
    @Override
    public int hashCode() {
        return Objects.hash(name, price);
    }

    // toString — human-readable representation
    @Override
    public String toString() {
        return "Product{name='" + name + "', price=" + price + "}";
    }

    // clone — creates a copy (implement Cloneable)
    @Override
    protected Object clone() throws CloneNotSupportedException {
        return super.clone();  // shallow copy
    }
}
```

> **Rule:** Always override `hashCode` when you override `equals` — otherwise broken HashMap behaviour.

---

### 🎯 How to Explain in Interview

"The Object class is the parent of all Java classes - every class implicitly extends it. It provides essential methods that all objects should have. The most important ones are equals(), hashCode(), and toString(). equals() defines what it means for two objects to be logically equal. hashCode() must be consistent with equals() - if two objects are equal, they must have the same hash code, otherwise HashMap and HashSet won't work correctly. toString() provides a human-readable representation, which is invaluable for debugging. There's also clone() for copying objects, though many prefer using copy constructors or serialization. Understanding these methods is crucial because they form the foundation of Java's object system and collections framework."

---

### Q10. What is covariant return type?

```java
class Animal {
    public Animal create() {
        return new Animal();
    }
}

class Dog extends Animal {
    @Override
    public Dog create() {   // ✅ covariant — Dog IS-A Animal
        return new Dog();
    }
}

// Use case: Builder pattern often relies on covariant return
class Builder {
    protected String name;
    public Builder name(String name) { this.name = name; return this; }
    public Builder build() { return this; }
}

class AdvancedBuilder extends Builder {
    private int level;
    public AdvancedBuilder level(int level) { this.level = level; return this; }
    
    @Override
    public AdvancedBuilder build() { return this; }  // covariant
}
```

---

### 🎯 How to Explain in Interview

"Covariant return types are a neat feature that lets me override a method to return a more specific type. Before this feature, if a parent method returned an Animal, the child had to return an Animal too. Now, the child can return a Dog, which is fine because a Dog IS-A Animal. This makes code much cleaner, especially in builder patterns and factory methods. It lets me maintain type safety without needing casts. For example, in a builder pattern, I can have each level return its own type, so method chaining works perfectly without the user having to cast. It's one of those small features that makes Java code more elegant and type-safe."
