# OOPs Concepts (66-75)

## 66. Singleton Class
**Principle**: Private constructor, static instance.
**Question**: Determine how to create a Singleton class.
**Code**:
```java
class Singleton {
    private static Singleton instance;
    private Singleton() {} // Private Constructor
    
    public static Singleton getInstance() {
        if (instance == null) instance = new Singleton();
        return instance;
    }
}
public class Main {
    public static void main(String[] args) {
        Singleton s1 = Singleton.getInstance();
        Singleton s2 = Singleton.getInstance();
        System.out.println(s1 == s2); // true
    }
}
```

## 67. Immutable Class
**Principle**: Final class, private final fields, no setters.
**Question**: Create an immutable class having one field.
**Code**:
```java
final class Immutable {
    private final String name;
    public Immutable(String name) { this.name = name; }
    public String getName() { return name; }
}
```

## 68. Method Overloading
**Principle**: Same name, different parameters.
**Question**: Demonstrate method overloading.
**Code**:
```java
class MathUtil {
    int add(int a, int b) { return a + b; }
    double add(double a, double b) { return a + b; }
}
```

## 69. Method Overriding
**Principle**: Same signature in subclass with `@Override`.
**Question**: Demonstrate method overriding.
**Code**:
```java
class Parent { void show() { System.out.println("Parent"); } }
class Child extends Parent { 
    @Override 
    void show() { System.out.println("Child"); } 
}
public class Main {
    public static void main(String[] args) {
        Parent p = new Child();
        p.show(); // Child
    }
}
```

## 70. Interface Implementation
**Principle**: Define contract, implement in class.
**Question**: Implement an interface.
**Code**:
```java
interface Animal { void sound(); }
class Dog implements Animal {
    public void sound() { System.out.println("Bark"); }
}
```

## 71. Abstract Class
**Principle**: Partial implementation, cannot be instantiated.
**Question**: Use an abstract class.
**Code**:
```java
abstract class Vehicle { abstract void drive(); }
class Car extends Vehicle {
    void drive() { System.out.println("Drive Car"); }
}
```

## 72. Custom Exception
**Principle**: Extend `Exception` or `RuntimeException`.
**Question**: Create a custom checked exception.
**Code**:
```java
class MyEx extends Exception {
    public MyEx(String msg) { super(msg); }
}
public class Main {
    public static void main(String[] args) {
        try { throw new MyEx("Error"); }
        catch(MyEx e) { System.out.println(e.getMessage()); }
    }
}
```

## 73. Deep Copy vs Shallow Copy
**Principle**: Shallow copies reference; Deep copies object graph.
**Question**: Demonstrate cloning (Shallow Copy).
**Code**:
```java
class Node implements Cloneable {
    int val;
    public Object clone() throws CloneNotSupportedException { return super.clone(); }
}
```

## 74. Static Block vs Instance Block
**Principle**: Static runs on class load (once), Instance on creation (every time).
**Question**: Show execution order of static and instance blocks.
**Code**:
```java
class Test {
    static { System.out.println("Static Block"); }
    { System.out.println("Instance Block"); }
    Test() { System.out.println("Constructor"); }
}
public class Main {
    public static void main(String[] args) { new Test(); new Test(); }
}
```

## 75. Comparator vs Comparable
**Principle**: `Comparable` (natural order, `compareTo`), `Comparator` (custom order, `compare`).
**Question**: Sort list using Comparator.
**Code**:
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<String> list = Arrays.asList("B", "A");
        Collections.sort(list, (s1, s2) -> s1.compareTo(s2));
        System.out.println(list);
    }
}
```
