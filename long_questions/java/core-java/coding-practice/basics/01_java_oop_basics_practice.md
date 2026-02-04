# 🟢 Java OOP Basics Practice
Contains runnable code examples for Questions 1-20.

## Question 1: What are the 4 Pillars of OOP? Explain with real-world examples.

### Answer
1.  **Encapsulation:** Bundling data (variables) and methods together and hiding details.
2.  **Abstraction:** Hiding complexity and showing only essential features.
3.  **Inheritance:** One class acquires properties of another.
4.  **Polymorphism:** One name, multiple forms.

### Runnable Code
```java
package basics;

// 1. Abstraction
abstract class Vehicle {
    abstract void start();
    
    public void stop() {
        System.out.println("Vehicle stopped.");
    }
}

// 3. Inheritance
class Car extends Vehicle {
    // 1. Encapsulation (Private fields with public accessors)
    private String model;
    
    public Car(String model) {
        this.model = model;
    }
    
    public String getModel() {
        return model;
    }

    // 4. Polymorphism (Overriding)
    @Override
    void start() {
        System.out.println(model + " car is starting... Vroom!");
    }
    
    // 4. Polymorphism (Overloading)
    void start(boolean remote) {
        if(remote) System.out.println(model + " started remotely.");
        else start();
    }
}

public class OOPPillarsDemo {
    public static void main(String[] args) {
        Car myCar = new Car("Tesla Model S");
        
        // Abstraction & Inheritance
        myCar.start(); 
        myCar.stop();
        
        // Polymorphism
        myCar.start(true);
        
        // Encapsulation
        System.out.println("Car Model: " + myCar.getModel());
    }
}
```

---

## Question 2: Difference between Abstract Class and Interface?

### Answer
*   **Abstract Class:** Can have constructors, instance variables, state. "Is-A" relationship.
*   **Interface:** Contract only (mostly). "Can-Do" relationship. Multiple inheritance support.

### Runnable Code
```java
package basics;

// Interface: Defines a capability
interface Flyable {
    void fly(); // implicitly public abstract
}

// Abstract Class: Defines a template
abstract class Bird {
    String name;
    
    Bird(String name) {
        this.name = name;
    }
    
    void eat() {
        System.out.println(name + " is eating.");
    }
    
    abstract void makeSound();
}

// Concrete class implementing both
class Eagle extends Bird implements Flyable {
    Eagle() {
        super("Eagle");
    }

    @Override
    void makeSound() {
        System.out.println("Screech!");
    }

    @Override
    public void fly() {
        System.out.println("Eagle is soaring high.");
    }
}

public class AbstractVsInterface {
    public static void main(String[] args) {
        Eagle e = new Eagle();
        e.eat();       // From Abstract Class
        e.makeSound(); // Implemented from Abstract Class
        e.fly();       // Implemented from Interface
    }
}
```

---

## Question 3: What is Polymorphism? (Compile-time vs Runtime)

### Answer
*   **Compile-time:** Method Overloading (Signature check).
*   **Runtime:** Method Overriding (Dynamic dispatch).

### Runnable Code
```java
package basics;

class Calculator {
    // Compile-time Polymorphism (Overloading)
    int add(int a, int b) {
        return a + b;
    }
    
    double add(double a, double b) {
        return a + b;
    }
}

class Animal {
    void sound() {
        System.out.println("Some generic animal sound");
    }
}

class Dog extends Animal {
    // Runtime Polymorphism (Overriding)
    @Override
    void sound() {
        System.out.println("Bark! Bark!");
    }
}

public class PolymorphismDemo {
    public static void main(String[] args) {
        // Compile-time
        Calculator calc = new Calculator();
        System.out.println("Int Sum: " + calc.add(5, 10));
        System.out.println("Double Sum: " + calc.add(5.5, 2.5));
        
        // Runtime
        Animal a = new Dog(); // Upcasting
        a.sound(); // Calls Dog's sound()
    }
}
```

---

## Question 4: Can you override `static` or `private` methods?

### Answer
*   **NO.**
*   **Static:** Class level. Hiding happens, not overriding.
*   **Private:** Not inherited.

### Runnable Code
```java
package basics;

class Parent {
    static void staticMethod() {
        System.out.println("Parent Static");
    }
    
    private void privateMethod() {
        System.out.println("Parent Private");
    }
    
    public void test() {
        privateMethod();
    }
}

class Child extends Parent {
    // Method Hiding (Not overriding)
    static void staticMethod() {
        System.out.println("Child Static");
    }
    
    // New method (Not overriding parent's private method)
    public void privateMethod() {
        System.out.println("Child Private");
    }
}

public class OverrideStaticPrivate {
    public static void main(String[] args) {
        Parent p = new Child();
        
        // Static method call depends on Reference Type (Parent), not Object Type (Child)
        p.staticMethod(); // Prints "Parent Static"
        
        // Private method is not overridden
        // p.privateMethod(); // Compile Error
        p.test(); // Prints "Parent Private" (Accessed internally)
    }
}
```

---

## Question 5: What is covariant return type?

### Answer
Overriding method can return a **subtype** of the parent method's return type.

### Runnable Code
```java
package basics;

class AnimalFactory {
    Animal getAnimal() {
        return new Animal();
    }
}

class DogFactory extends AnimalFactory {
    @Override
    Dog getAnimal() { // Covariant return type (Dog is-a Animal)
        return new Dog();
    }
}

public class CovariantReturn {
    public static void main(String[] args) {
        AnimalFactory af = new DogFactory();
        Animal a = af.getAnimal(); // Returns Dog object
        System.out.println("Obtained: " + a.getClass().getSimpleName());
    }
}
```

---

## Question 6: Composition vs Inheritance. Which is better?

### Answer
*   **Inheritance:** Is-A. Tightly coupled.
*   **Composition:** Has-A. Loosely coupled. Preferred.

### Runnable Code
```java
package basics;

// Inheritance Way (Bad if strict hierarchy isn't needed)
class PcWithInheritance extends Calculator { 
    // Inherits add() methods
}

// Composition Way (Flexible)
class PcWithComposition {
    private Calculator calc = new Calculator(); // Has-A
    
    public int compute(int a, int b) {
        return calc.add(a, b); // Delegating
    }
}

public class CompositionVsInheritance {
    public static void main(String[] args) {
        PcWithComposition pc = new PcWithComposition();
        System.out.println("Result: " + pc.compute(10, 20));
    }
}
```

---

## Question 7: What is the `super` keyword?

### Answer
Refers to the immediate parent class object.

### Runnable Code
```java
package basics;

class Base {
    int num = 100;
    Base() { System.out.println("Base Constructor"); }
    void print() { System.out.println("Base Print"); }
}

class Derived extends Base {
    int num = 200;
    
    Derived() {
        super(); // 2. Call Parent Constructor
        System.out.println("Derived Constructor");
    }
    
    void display() {
        System.out.println("Local num: " + num);
        System.out.println("Parent num: " + super.num); // 1. Access Parent Variable
        super.print(); // 3. Call Parent Method
    }
}

public class SuperDemo {
    public static void main(String[] args) {
        new Derived().display();
    }
}
```

---

## Question 8: Significance of `this` keyword?

### Answer
Refers to the current object instance.

### Runnable Code
```java
package basics;

class Student {
    String name;
    
    Student() {
        this("Unknown"); // Constructor chaining
        System.out.println("Default constructor called.");
    }
    
    Student(String name) {
        this.name = name; // Disambiguate field vs param
    }
    
    Student getSelf() {
        return this; // Return current object
    }
    
    void print() {
        System.out.println("Student: " + this.name);
    }
}

public class ThisDemo {
    public static void main(String[] args) {
        Student s1 = new Student(); // Calls default -> parameterized
        s1.print();
        
        Student s2 = new Student("Alice").getSelf();
        s2.print();
    }
}
```

---

## Question 9: Can an Interface extend another Interface?

### Answer
**Yes**, multiple interfaces too.

### Runnable Code
```java
package basics;

interface A { void methodA(); }
interface B { void methodB(); }

// Inheritance in Interfaces
interface C extends A, B { 
    void methodC(); 
}

class Impl implements C {
    public void methodA() { System.out.println("A"); }
    public void methodB() { System.out.println("B"); }
    public void methodC() { System.out.println("C"); }
}

public class InterfaceInheritance {
    public static void main(String[] args) {
        C obj = new Impl();
        obj.methodA();
        obj.methodB();
        obj.methodC();
    }
}
```

---

## Question 10: Difference between `Overloading` and `Overriding`?

### Answer
See Question 3. Code demo specifically highlighting key differences.

### Runnable Code
```java
package basics;

class Test {
    // Overloading: Same name, diff param, static binding
    void show(int a) { System.out.println("Show Int: " + a); }
    void show(String s) { System.out.println("Show String: " + s); }
}

class ParentTest {
    void display() { System.out.println("Parent Display"); }
}

class ChildTest extends ParentTest {
    // Overriding: Same signature, dynamic binding
    @Override
    void display() { System.out.println("Child Display"); }
}

public class OverloadOverride {
    public static void main(String[] args) {
        // Overloading
        Test t = new Test();
        t.show(1);
        t.show("Hi");
        
        // Overriding
        ParentTest p = new ChildTest();
        p.display(); // Child Display
    }
}
```

---

## Question 11: Can a constructor be private? Why?

### Answer
**Yes.** Used in Singleton Pattern.

### Runnable Code
```java
package basics;

class Singleton {
    private static Singleton instance;
    
    // Private Constructor prevents 'new Singleton()'
    private Singleton() { 
        System.out.println("Singleton Created");
    }
    
    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
    
    public void action() { System.out.println("Action performed"); }
}

public class PrivateConstructor {
    public static void main(String[] args) {
        // Singleton s = new Singleton(); // Compilation Error
        Singleton s = Singleton.getInstance();
        s.action();
    }
}
```

---

## Question 12: Default Constructor vs No-Args Constructor?

### Answer
*   **Default:** Provided by compiler if NO constructor exists.
*   **No-Args:** Manually defined empty constructor.

### Runnable Code
```java
package basics;

class DefaultCon {
    // No code here. Compiler creates DefaultCon() {}
}

class NoDefault {
    NoDefault(int x) { // Parameterized Only
        System.out.println("Param Constructor");
    }
}

class ManualNoArgs {
    ManualNoArgs() {
        System.out.println("Manual No-Args");
    }
}

public class ConstructorTypes {
    public static void main(String[] args) {
        new DefaultCon(); // Works (Compiler generated)
        
        // new NoDefault(); // Error: Default constructor removed!
        new NoDefault(5);
        
        new ManualNoArgs();
    }
}
```

---

## Question 13: What is "Constructor Chaining"?

### Answer
Calling one constructor from another (`this()` or `super()`).

### Runnable Code
```java
package basics;

class ChainRoot {
    ChainRoot() {
        System.out.println("1. Root Constructor");
    }
}

class ChainChild extends ChainRoot {
    ChainChild() {
        this("Default"); // Call overloaded constructor
        System.out.println("4. Child No-Args Finished");
    }
    
    ChainChild(String s) {
        super(); // Implicit, but good to show order.
        System.out.println("2. Child Param Constructor: " + s);
        System.out.println("3. Continuing chain...");
    }
}

public class ConstructorChaining {
    public static void main(String[] args) {
        new ChainChild();
    }
}
```

---

## Question 14: Use of `instanceof` operator?

### Answer
Checks type. Java 14+ has Pattern Matching.

### Runnable Code
```java
package basics;

public class InstanceOfDemo {
    public static void main(String[] args) {
        Object obj = "Hello Java";
        
        // Traditional
        if (obj instanceof String) {
            String s = (String) obj;
            System.out.println("String length: " + s.length());
        }
        
        // Java 14+ Pattern Matching
        if (obj instanceof String s) {
            System.out.println("Pattern Matching Length: " + s.length());
        }
    }
}
```

---

## Question 15: What is an Initialization Block?

### Answer
*   Instance Block `{}`: Runs before constructor.
*   Static Block `static {}`: Runs once at class load.

### Runnable Code
```java
package basics;

class InitDemo {
    static {
        System.out.println("1. Static Block (Class Loaded)");
    }
    
    {
        System.out.println("2. Instance Block (Object Creating)");
    }
    
    InitDemo() {
        System.out.println("3. Constructor");
    }
}

public class InitBlockDemo {
    public static void main(String[] args) {
        System.out.println("Main Start");
        new InitDemo();
        new InitDemo(); // Static runs only once
    }
}
```

---

## Question 16: Multiple Inheritance in Java?

### Answer
Classes NO. Interfaces YES (via default methods for logic).

### Runnable Code
```java
package basics;

interface I1 {
    default void show() { System.out.println("I1 Show"); }
}

interface I2 {
    default void show() { System.out.println("I2 Show"); }
}

class MultiInh implements I1, I2 {
    // Diamond Problem solved: Must override if conflict
    @Override
    public void show() {
        I1.super.show(); // Call specific parent
        I2.super.show(); 
        System.out.println("Child Show");
    }
}

public class MultipleInheritance {
    public static void main(String[] args) {
        new MultiInh().show();
    }
}
```

---

## Question 17: What is a Marker Interface?

### Answer
Empty interface (e.g., `Serializable`).

### Runnable Code
```java
package basics;

// Custom Marker
interface SafeToDelete {}

class TempFile implements SafeToDelete {
    String name = "temp.txt";
}

class CriticalFile {
    String name = "passwords.txt";
}

public class MarkerDemo {
    static void delete(Object obj) {
        if (obj instanceof SafeToDelete) {
            System.out.println("Deleting: " + obj.getClass().getSimpleName());
        } else {
            System.out.println("Cannot Delete: " + obj.getClass().getSimpleName());
        }
    }

    public static void main(String[] args) {
        delete(new TempFile());
        delete(new CriticalFile());
    }
}
```

---

## Question 18: Can abstract class have constructor?

### Answer
**Yes.** Called by subclass.

### Runnable Code
```java
package basics;

abstract class Template {
    String type;
    
    Template(String type) {
        this.type = type;
        System.out.println("Template initialized: " + type);
    }
}

class RealImpl extends Template {
    RealImpl() {
        super("Real"); // Must call parent constructor
        System.out.println("RealImpl Created");
    }
}

public class AbstractConstructor {
    public static void main(String[] args) {
        new RealImpl();
    }
}
```

---

## Question 19: Shallow Copy vs Deep Copy (Object Cloning)?

### Answer
*   **Shallow:** Copies references.
*   **Deep:** Creates new specific copies of objects.

### Runnable Code
```java
package basics;

class Address implements Cloneable {
    String city;
    Address(String c) { city = c; }
    
    @Override
    protected Object clone() throws CloneNotSupportedException {
        return super.clone();
    }
}

class Customer implements Cloneable {
    String name;
    Address addr;
    
    Customer(String n, String c) {
        name = n;
        addr = new Address(c);
    }
    
    // Default clone() is Shallow
    @Override
    protected Object clone() throws CloneNotSupportedException {
        Customer cloned = (Customer) super.clone();
        // Deep Copy Manual Step:
        cloned.addr = (Address) this.addr.clone(); 
        return cloned;
    }
}

public class DeepCopyDemo {
    public static void main(String[] args) throws CloneNotSupportedException {
        Customer c1 = new Customer("John", "NY");
        Customer c2 = (Customer) c1.clone();
        
        c2.addr.city = "LA";
        
        // If Deep Copy worked, c1 should remains NY
        System.out.println("Original City: " + c1.addr.city); // Prints NY
        System.out.println("Cloned City: " + c2.addr.city);   // Prints LA
    }
}
```

---

## Question 20: Immutable Class - How to create one?

### Answer
Final class, private final fields, no setters, deep copy in getters.

### Runnable Code
```java
package basics;

import java.util.*;

final class ImmutableStudent {
    private final String name;
    private final List<String> subjects; // Mutable type
    
    public ImmutableStudent(String name, List<String> subjects) {
        this.name = name;
        // Deep Copy in Constructor
        this.subjects = new ArrayList<>(subjects);
    }
    
    public String getName() { return name; }
    
    public List<String> getSubjects() {
        // Deep Copy in Getter (Return unmodifiable view or copy)
        return new ArrayList<>(subjects);
    }
}

public class ImmutableDemo {
    public static void main(String[] args) {
        List<String> subs = new ArrayList<>();
        subs.add("Math");
        
        ImmutableStudent student = new ImmutableStudent("Dhruv", subs);
        
        // Try to modify original list
        subs.add("Science");
        
        // Try to modify getter list
        student.getSubjects().add("History");
        
        System.out.println("Student Subjects: " + student.getSubjects()); // Only [Math]
    }
}
```
