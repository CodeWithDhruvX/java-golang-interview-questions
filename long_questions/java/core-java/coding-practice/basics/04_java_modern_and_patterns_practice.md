# 🚀 Java Modern & Patterns Practice
Contains runnable code examples for Questions 58-74.

## Question 58: Types of Inner Classes.

### Answer
Member, Static Nested, Local, Anonymous.

### Runnable Code
```java
package basics;

class Outer {
    private String msg = "Hello";
    
    // 1. Member Inner
    class Inner {
        void show() { System.out.println(msg + " from Inner"); }
    }
    
    // 2. Static Nested
    static class Nested {
        void show() { System.out.println("Static Nested (No access to instance msg)"); }
    }
    
    void method() {
        // 3. Local Inner
        class Local {
            void run() { System.out.println("Local Inner Class"); }
        }
        new Local().run();
        
        // 4. Anonymous Inner
        Runnable r = new Runnable() {
            public void run() { System.out.println("Anonymous Inner"); }
        };
        r.run();
    }
}

public class InnerClassesDemo {
    public static void main(String[] args) {
        Outer o = new Outer();
        
        // Member
        Outer.Inner inner = o.new Inner();
        inner.show();
        
        // Static Nested
        Outer.Nested nested = new Outer.Nested();
        nested.show();
        
        o.method();
    }
}
```

---

## Question 59: Java Enums (More than just constants?).

### Answer
Enums are classes with fields/methods.

### Runnable Code
```java
package basics;

enum Status {
    START(1), RUNNING(2), STOP(0);
    
    private final int code;
    
    Status(int code) { this.code = code; }
    
    public int getCode() { return code; }
}

public class EnumDemo {
    public static void main(String[] args) {
        for (Status s : Status.values()) {
            System.out.println(s + " -> " + s.getCode());
        }
        
        Status current = Status.RUNNING;
        if (current == Status.RUNNING) {
            System.out.println("System is running...");
        }
    }
}
```

---

## Question 60: Java Records (Java 14+).

### Answer
Immutable Data Classes.

### Runnable Code
```java
package basics;

// Canonical Constructor, getters, equals/hashcode generated
record Point(int x, int y) {}

public class RecordDemo {
    public static void main(String[] args) {
        Point p1 = new Point(10, 20);
        Point p2 = new Point(10, 20);
        
        System.out.println(p1); // Point[x=10, y=20]
        System.out.println("Equal? " + p1.equals(p2)); // true
        System.out.println("X: " + p1.x()); // Accessor (no 'get' prefix)
    }
}
```

---

## Question 61: Sealed Classes (Java 17+).

### Answer
Restrict inheritance.

### Runnable Code
```java
package basics;

sealed abstract class Shape permits Circle, Square {}

final class Circle extends Shape {
    void draw() { System.out.println("Circle"); }
}

final class Square extends Shape {
    void draw() { System.out.println("Square"); }
}

// class Triangle extends Shape {} // Compile Error: Not permitted

public class SealedDemo {
    public static void main(String[] args) {
        Shape s = new Circle();
        System.out.println("Allowed shape: " + s.getClass().getSimpleName());
    }
}
```

---

## Question 62: Text Blocks (Java 15+).

### Answer
Multi-line strings.

### Runnable Code
```java
package basics;

public class TextBlockDemo {
    public static void main(String[] args) {
        String json = """
            {
                "name": "Dhruv",
                "role": "Developer"
            }
            """;
        System.out.println(json);
    }
}
```

---

## Question 63: Switch Expressions (Java 14+).

### Answer
Switch returning value.

### Runnable Code
```java
package basics;

public class SwitchExpr {
    public static void main(String[] args) {
        String day = "MONDAY";
        
        // Returns value, no break needed
        int len = switch (day) {
            case "MONDAY", "FRIDAY" -> 6;
            case "TUESDAY" -> 7;
            default -> 0;
        };
        
        System.out.println("Length: " + len);
    }
}
```

---

## Question 64: `var` keyword (Java 10+).

### Answer
Local variable type inference.

### Runnable Code
```java
package basics;

import java.util.ArrayList;

public class VarDemo {
    public static void main(String[] args) {
        var msg = "Hello"; // String inferred
        var list = new ArrayList<String>(); // ArrayList<String> inferred
        
        list.add(msg);
        System.out.println(list);
    }
}
```

---

## Question 65: Core Functional Interfaces (`Supplier`, `Consumer`, etc).

### Answer
Predicate, Consumer, Supplier, Function.

### Runnable Code
```java
package basics;

import java.util.function.*;

public class FunctionalInterfaces {
    public static void main(String[] args) {
        // 1. Predicate: Input -> Boolean
        Predicate<String> isEmpty = s -> s.isEmpty();
        System.out.println("Is Empty? " + isEmpty.test(""));
        
        // 2. Consumer: Input -> Void
        Consumer<String> print = s -> System.out.println("Consuming: " + s);
        print.accept("Java");
        
        // 3. Supplier: Void -> Output
        Supplier<Double> random = () -> Math.random();
        System.out.println("Random: " + random.get());
        
        // 4. Function: Input -> Output
        Function<Integer, String> intToStr = i -> "Number: " + i;
        System.out.println(intToStr.apply(100));
    }
}
```

---

## Question 66: What is `@FunctionalInterface`?

### Answer
Ensures interface has single abstract method.

### Runnable Code
```java
package basics;

@FunctionalInterface
interface MathOp {
    int operate(int a, int b);
    // int another(); // Error if uncommented
}

public class FunctionalAnno {
    public static void main(String[] args) {
        MathOp add = (a, b) -> a + b;
        System.out.println("Sum: " + add.operate(5, 3));
    }
}
```

---

## Question 67: Singleton Pattern (Strategies).

### Answer
Enum is best.

### Runnable Code
```java
package basics;

// Enum Singleton (Thread-safe, Serialization-safe)
enum DBConnection {
    INSTANCE;
    
    public void query() {
        System.out.println("Executing Query...");
    }
}

public class SingletonPattern {
    public static void main(String[] args) {
        DBConnection.INSTANCE.query();
    }
}
```

---

## Question 68: Factory Pattern.

### Answer
Creats objects without exposing logic.

### Runnable Code
```java
package basics;

interface Notification { void notifyUser(); }

class SMS implements Notification {
    public void notifyUser() { System.out.println("Sending SMS"); }
}

class Email implements Notification {
    public void notifyUser() { System.out.println("Sending Email"); }
}

class NotifFactory {
    static Notification create(String type) {
        return switch(type) {
            case "SMS" -> new SMS();
            case "EMAIL" -> new Email();
            default -> throw new IllegalArgumentException();
        };
    }
}

public class FactoryPattern {
    public static void main(String[] args) {
        Notification n = NotifFactory.create("EMAIL");
        n.notifyUser();
    }
}
```

---

## Question 69: Builder Pattern.

### Answer
Constructs complex objects.

### Runnable Code
```java
package basics;

class Burger {
    String bun;
    String patty;
    boolean cheese;
    
    private Burger(Builder b) {
        this.bun = b.bun;
        this.patty = b.patty;
        this.cheese = b.cheese;
    }
    
    static class Builder {
        private String bun;
        private String patty;
        private boolean cheese;
        
        public Builder(String bun) { this.bun = bun; }
        
        public Builder patty(String p) { this.patty = p; return this; }
        public Builder addCheese() { this.cheese = true; return this; }
        
        public Burger build() { return new Burger(this); }
    }
    
    public String toString() { return bun + "/" + patty + (cheese ? "/Cheese" : ""); }
}

public class BuilderPattern {
    public static void main(String[] args) {
        Burger b = new Burger.Builder("Sesame").patty("Chicken").addCheese().build();
        System.out.println(b);
    }
}
```

---

## Question 70: Observer Pattern.

### Answer
Pub/Sub event listener.

### Runnable Code
```java
package basics;

import java.util.*;

// Subject
class Channel {
    List<String> subs = new ArrayList<>();
    
    void subscribe(String name) { subs.add(name); }
    
    void upload(String title) {
        for(String s : subs) System.out.println(s + " notified for " + title);
    }
}

public class ObserverDemo {
    public static void main(String[] args) {
        Channel yt = new Channel();
        yt.subscribe("User1");
        yt.subscribe("User2");
        yt.upload("New Video!");
    }
}
```

---

## Question 71: Java 8 Date/Time API (`java.time`) vs Legacy.

### Answer
Immutable, thread-safe.

### Runnable Code
```java
package basics;

import java.time.*;

public class DateTimeDemo {
    public static void main(String[] args) {
        LocalDate date = LocalDate.now();
        LocalTime time = LocalTime.now();
        LocalDateTime dt = LocalDateTime.now();
        
        System.out.println("Date: " + date);
        System.out.println("Time: " + time);
        
        // Immutable modification (returns new instance)
        LocalDate tomorrow = date.plusDays(1);
        System.out.println("Tomorrow: " + tomorrow);
    }
}
```

---

## Question 72: Reference Types (Strong, Soft, Weak, Phantom).

### Answer
GC behavior control.

### Runnable Code
```java
package basics;

import java.lang.ref.WeakReference;

public class ReferencesDemo {
    public static void main(String[] args) {
        // Strong
        Object strong = new Object();
        
        // Weak (Collected if no strong ref exists)
        WeakReference<Object> weak = new WeakReference<>(new Object());
        
        System.out.println("Weak Before GC: " + weak.get());
        
        System.gc();
        
        System.out.println("Weak After GC: " + weak.get()); // Likely null
    }
}
```

---

## Question 73: `Statement` vs `PreparedStatement`.

### Answer
Simulated JDBC.

### Runnable Code
```java
package basics;

public class JDBCSimulation {
    static void executeQuery(String sql) {
        System.out.println("Executing: " + sql);
    }
    
    public static void main(String[] args) {
        String user = "admin";
        
        // Statement (Vulnerable)
        String sql1 = "SELECT * FROM users WHERE name = '" + user + "'";
        executeQuery(sql1);
        
        // PreparedStatement (Safe)
        String sql2 = "SELECT * FROM users WHERE name = ?";
        System.out.println("Preparing: " + sql2);
        System.out.println("Binding: [1] = " + user);
        executeQuery("SELECT * FROM users WHERE name = 'admin'");
    }
}
```

---

## Question 74: Transaction Management in JDBC.

### Answer
AutoCommit(false) -> Commit/Rollback.

### Runnable Code
```java
package basics;

public class TransactionDemo {
    public static void main(String[] args) {
        System.out.println("Conn: setAutoCommit(false)");
        
        try {
            System.out.println("Debit 500 from A");
            System.out.println("Credit 500 to B");
            
            boolean error = false; 
            if (error) throw new RuntimeException();
            
            System.out.println("Conn: commit()");
        } catch (Exception e) {
            System.out.println("Conn: rollback()");
        }
    }
}
```
