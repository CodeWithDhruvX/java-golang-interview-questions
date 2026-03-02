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
