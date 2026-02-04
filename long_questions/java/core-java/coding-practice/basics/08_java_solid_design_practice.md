# 🏗️ Java SOLID Design & References Practice
Contains runnable code for SOLID principles, Comparator vs Comparable, Reference Types, and System hooks.

## Question 1: SOLID - Single Responsibility Principle (SRP) Example.

### Answer
Class should have one reason to change. Separate Logic from Persistence.

### Runnable Code
```java
package design;

// Bad: Employee class does mixed tasks
class BadEmployee {
    void calculatePay() {}
    void saveToDatabase() {}
}

// Good: Logic and Persistence separated
class Employee {
    String name;
}

class PayrollService {
    void calculatePay(Employee e) {
        System.out.println("Calculating Pay...");
    }
}

class EmployeeRepository {
    void save(Employee e) {
        System.out.println("Saving to DB...");
    }
}

public class SRPDemo {
    public static void main(String[] args) {
        new PayrollService().calculatePay(new Employee());
        new EmployeeRepository().save(new Employee());
    }
}
```

---

## Question 2: SOLID - Open/Closed Principle (OCP) Example.

### Answer
Open for extension, closed for modification. Use Polymorphism.

### Runnable Code
```java
package design;

interface Shape {
    double area();
}

class Rectangle implements Shape {
    double w, h;
    Rectangle(double w, double h) { this.w = w; this.h = h; }
    public double area() { return w * h; }
}

class Circle implements Shape {
    double r;
    Circle(double r) { this.r = r; }
    public double area() { return Math.PI * r * r; }
}

class AreaCalculator {
    // Works for ALL shapes without modifying this code
    double totalArea(Shape... shapes) {
        double sum = 0;
        for(Shape s : shapes) sum += s.area();
        return sum;
    }
}

public class OCPDemo {
    public static void main(String[] args) {
        System.out.println("Total: " + new AreaCalculator().totalArea(new Rectangle(2,3), new Circle(2)));
    }
}
```

---

## Question 3: `Comparable` vs `Comparator` implementation?

### Answer
`Comparable` (Natural order, `compareTo`). `Comparator` (Custom order, `compare`).

### Runnable Code
```java
package design;

import java.util.*;

class Student implements Comparable<Student> {
    int id;
    String name;
    
    Student(int id, String name) { this.id = id; this.name = name; }
    
    // Natural Order: by ID
    public int compareTo(Student o) { return this.id - o.id; }
    
    public String toString() { return id + ":" + name; }
}

public class ComparatorDemo {
    public static void main(String[] args) {
        List<Student> list = new ArrayList<>();
        list.add(new Student(2, "Bob"));
        list.add(new Student(1, "Alice"));
        
        // 1. Natural Sort (Comparable)
        Collections.sort(list);
        System.out.println("Natural: " + list); // 1:Alice, 2:Bob
        
        // 2. Custom Sort (Comparator) - Sort by Name
        Collections.sort(list, new Comparator<Student>() {
            public int compare(Student s1, Student s2) {
                return s1.name.compareTo(s2.name);
            }
        });
        
        // Java 8 Style
        // list.sort((s1, s2) -> s2.id - s1.id); // Reverse ID
        
        System.out.println("ByName: " + list); // 1:Alice, 2:Bob (Sorted by name logic)
    }
}
```

---

## Question 4: Reference Types: Strong, Soft, Weak, Phantom.

### Answer
Strong (Normal), Soft (Cache), Weak (Mapping), Phantom (Cleanup).

### Runnable Code
```java
package design;

import java.lang.ref.*;

public class RefTypesDemo {
    public static void main(String[] args) {
        // 1. Strong
        Object strong = new Object();
        
        // 2. Soft (Cleared only when memory is low)
        SoftReference<Object> soft = new SoftReference<>(new Object());
        
        // 3. Weak (Cleared on next GC)
        WeakReference<Object> weak = new WeakReference<>(new Object());
        System.out.println("Weak Before GC: " + weak.get());
        System.gc();
        System.out.println("Weak After GC: " + weak.get()); // null
        
        // 4. Phantom (Used with ReferenceQueue for cleanup)
        ReferenceQueue<Object> queue = new ReferenceQueue<>();
        PhantomReference<Object> phantom = new PhantomReference<>(new Object(), queue);
        System.out.println("Phantom Get is always: " + phantom.get()); // always null
    }
}
```

---

## Question 5: How to use `ShutdownHook`?

### Answer
Run code when JVM terminates.

### Runnable Code
```java
package design;

public class ShutdownHookDemo {
    public static void main(String[] args) {
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("JVM is shutting down... Cleaning up resources.");
        }));
        
        System.out.println("Main Application Running...");
        // System.exit(0); // Triggers hook
    }
}
```

---

## Question 6: Dependency Injection (Manual Implementation).

### Answer
Invert control by passing dependencies.

### Runnable Code
```java
package design;

class Database { void connect() { System.out.println("DB Connected"); } }

// Tightly Coupled
class ServiceA {
    Database db = new Database(); // Hard dependency
}

// Loose Coupled (DI)
class ServiceB {
    Database db;
    // Constructor Injection
    ServiceB(Database db) { this.db = db; }
    void run() { db.connect(); }
}

public class DIDemo {
    public static void main(String[] args) {
        ServiceB service = new ServiceB(new Database());
        service.run();
    }
}
```
