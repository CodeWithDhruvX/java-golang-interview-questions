# ⚙️ Java Fundamentals Practice
Contains runnable code examples for Questions 41-57.

## Question 41: Explain `static` keyword in Java.

### Answer
Belongs to class. Variable, Method, Block, Nested Class.

### Runnable Code
```java
package basics;

class StaticDemo {
    static int count = 0; // Shared variable
    
    // Static Block (Runs once)
    static {
        System.out.println("Class Loaded");
    }
    
    StaticDemo() { count++; }
    
    // Static Method
    static void printCount() {
        System.out.println("Count: " + count);
        // System.out.println(this); // Error: No 'this' in static context
    }
    
    // Static Nested Class
    static class Nested {
        void msg() { System.out.println("Static Nested Class"); }
    }
}

public class StaticRun {
    public static void main(String[] args) {
        new StaticDemo();
        new StaticDemo();
        StaticDemo.printCount(); // 2
        
        new StaticDemo.Nested().msg();
    }
}
```

---

## Question 42: What does `volatile` do?

### Answer
Ensures visibility (no caching in thread local memory) and prevents instruction reordering.

### Runnable Code
```java
package basics;

class VolatileFlags {
    // Ensures 'running' is read from main memory, not CPU cache
    private volatile boolean running = true;
    
    public void stop() {
        running = false;
        System.out.println("Stop signal sent");
    }
    
    public void runLoop() {
        System.out.println("Loop started");
        while (running) {
            // Busy wait (If not volatile, this loop might never exit)
        }
        System.out.println("Loop stopped");
    }
}

public class VolatileDemo {
    public static void main(String[] args) throws InterruptedException {
        VolatileFlags task = new VolatileFlags();
        
        new Thread(task::runLoop).start();
        Thread.sleep(100);
        task.stop();
    }
}
```

---

## Question 43: Comparing Objects: `==` vs `equals()`.

### Answer
`==` (Reference), `equals` (Content).

### Runnable Code
```java
package basics;

public class EqualityDemo {
    public static void main(String[] args) {
        String s1 = new String("Java");
        String s2 = new String("Java");
        String s3 = s1;
        
        System.out.println(s1 == s2);      // false (Different objects)
        System.out.println(s1.equals(s2)); // true (Same content)
        System.out.println(s1 == s3);      // true (Same reference)
    }
}
```

---

## Question 44: Common `Object` methods (`toString`, `hashCode`, `equals`).

### Answer
Methods from `java.lang.Object`.

### Runnable Code
```java
package basics;

import java.util.Objects;

class Person {
    String name;
    
    Person(String name) { this.name = name; }
    
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Person)) return false;
        Person person = (Person) o;
        return Objects.equals(name, person.name);
    }
    
    @Override
    public int hashCode() {
        return Objects.hash(name);
    }
    
    @Override
    public String toString() {
        return "Person{name='" + name + "'}";
    }
}

public class ObjectMethods {
    public static void main(String[] args) {
        Person p1 = new Person("A");
        Person p2 = new Person("A");
        
        System.out.println(p1); // toString
        System.out.println(p1.equals(p2)); // true
        System.out.println(p1.hashCode() == p2.hashCode()); // true
    }
}
```

---

## Question 45: `finalize()` method - Why is it deprecated?

### Answer
Unpredictable/Slow. Removed in Java 9.

### Runnable Code
*(See Question 24 code)*

---

## Question 46: Wrapper Classes & Autoboxing.

### Answer
Primitive <-> Object conversion.

### Runnable Code
```java
package basics;

import java.util.ArrayList;
import java.util.List;

public class WrapperDemo {
    public static void main(String[] args) {
        int prim = 10;
        
        // Autoboxing (int -> Integer)
        Integer obj = prim; 
        
        // Unboxing (Integer -> int)
        int val = obj;
        
        List<Integer> list = new ArrayList<>();
        list.add(100); // Boxes 100 to Integer
        
        // Trap: Unboxing null
        Integer nil = null;
        try {
            int crash = nil;
        } catch(NullPointerException e) {
            System.out.println("NPE on unboxing null");
        }
    }
}
```

---

## Question 47: Integer Cache (-128 to 127).

### Answer
Small Integers are cached.

### Runnable Code
```java
package basics;

public class IntegerCache {
    public static void main(String[] args) {
        Integer a = 127;
        Integer b = 127;
        System.out.println("127: " + (a == b)); // true (Cached)
        
        Integer c = 128;
        Integer d = 128;
        System.out.println("128: " + (c == d)); // false (New Objects)
    }
}
```

---

## Question 48: `BigInteger` and `BigDecimal`.

### Answer
Precision numbers (Crypto, Finance).

### Runnable Code
```java
package basics;

import java.math.BigDecimal;
import java.math.BigInteger;

public class PrecisionMath {
    public static void main(String[] args) {
        // Floating point error
        System.out.println(0.1 + 0.2); // 0.30000000000000004
        
        // BigDecimal
        BigDecimal b1 = new BigDecimal("0.1");
        BigDecimal b2 = new BigDecimal("0.2");
        System.out.println(b1.add(b2)); // 0.3
        
        // BigInteger (Large Factorial)
        BigInteger bi = new BigInteger("123456789123456789");
        System.out.println(bi.multiply(BigInteger.valueOf(2)));
    }
}
```

---

## Question 49: What is Type Erasure?

### Answer
Generics exist only at compile-time. Removed at runtime.

### Runnable Code
```java
package basics;

import java.util.*;

public class TypeErasure {
    public static void main(String[] args) {
        List<String> strings = new ArrayList<>();
        List<Integer> numbers = new ArrayList<>();
        
        // At runtime, both are just ArrayList.class
        System.out.println(strings.getClass() == numbers.getClass()); // true
        
        // You cannot check generic types at runtime
        // if (strings instanceof List<String>) {} // Error
    }
}
```

---

## Question 50: Wildcards in Generics (`?`, `extends`, `super`).

### Answer
PECS: Producer Extends, Consumer Super.

### Runnable Code
```java
package basics;

import java.util.*;

public class Wildcards {
    // Producer: We READ Number (or subclasses) from list
    static double sum(List<? extends Number> list) {
        double s = 0;
        for (Number n : list) s += n.doubleValue();
        // list.add(10); // Error: Cannot add
        return s;
    }
    
    // Consumer: We WRITE Integer to list
    static void addNumbers(List<? super Integer> list) {
        list.add(10); // OK
        // Number n = list.get(0); // Error: Returns Object
    }

    public static void main(String[] args) {
        List<Integer> ints = new ArrayList<>(Arrays.asList(1, 2, 3));
        List<Number> nums = new ArrayList<>();
        
        System.out.println("Sum: " + sum(ints));
        
        addNumbers(nums);
        System.out.println("Nums: " + nums);
    }
}
```

---

## Question 51: Generic Methods.

### Answer
Method declares its own type <T>.

### Runnable Code
```java
package basics;

public class GenericMethod {
    // Defined <T> before return type
    static <T> void print(T item) {
        System.out.println("Item: " + item);
    }

    public static void main(String[] args) {
        print(100);
        print("Hello");
        print(new Object());
    }
}
```

---

## Question 52: What is Reflection? Pros/Cons.

### Answer
Inspect code at runtime. Pros: Flexibility. Cons: Perf, Safety.

### Runnable Code
```java
package basics;

import java.lang.reflect.Method;

public class ReflectionSimple {
    public static void main(String[] args) throws Exception {
        Class<?> clazz = String.class;
        System.out.println("Class Name: " + clazz.getName());
        
        // List declared methods
        Method[] methods = clazz.getDeclaredMethods();
        System.out.println("Method Count: " + methods.length);
    }
}
```

---

## Question 53: How to access a Private Field using Reflection?

### Answer
`setAccessible(true)`.

### Runnable Code
```java
package basics;

import java.lang.reflect.Field;

class Secret {
    private String token = "Hidden123";
}

public class ReflectPrivate {
    public static void main(String[] args) throws Exception {
        Secret obj = new Secret();
        
        Field field = Secret.class.getDeclaredField("token");
        field.setAccessible(true); // Hack the gate open
        
        String value = (String) field.get(obj);
        System.out.println("Hacked: " + value);
    }
}
```

---

## Question 54: What is the `Class` class?

### Answer
Entry point for reflection. References the loaded class.

### Runnable Code
*(See Q52)*

---

## Question 55: Custom Annotations & Meta-Annotations.

### Answer
`@interface`. Meta: `@Retention`, `@Target`.

### Runnable Code
```java
package basics;

import java.lang.annotation.*;
import java.lang.reflect.Method;

@Retention(RetentionPolicy.RUNTIME) // Keep for reflection
@Target(ElementType.METHOD)
@interface TestInit {
    String value() default "Default";
}

class App {
    @TestInit("Startup")
    public void init() { System.out.println("Init running..."); }
}

public class AnnotationDemo {
    public static void main(String[] args) throws Exception {
        App app = new App();
        
        for (Method m : App.class.getMethods()) {
            if (m.isAnnotationPresent(TestInit.class)) {
                TestInit anno = m.getAnnotation(TestInit.class);
                System.out.println("Found @TestInit with value: " + anno.value());
                m.invoke(app);
            }
        }
    }
}
```

---

## Question 56: Breaking Singleton using Reflection.

### Answer
Use `setAccessible` on constructor.

### Runnable Code
```java
package basics;

import java.lang.reflect.Constructor;

class Single {
    private static final Single INSTANCE = new Single();
    private Single() {} 
    // Defense: if (INSTANCE != null) throw RuntimeException();
}

public class BreakSingletonReflect {
    public static void main(String[] args) throws Exception {
        Constructor<Single> cons = Single.class.getDeclaredConstructor();
        cons.setAccessible(true);
        
        Single s1 = cons.newInstance();
        Single s2 = cons.newInstance();
        
        System.out.println(s1 == s2); // false (Broken)
    }
}
```

---

## Question 57: Private vs Default vs Protected vs Public.

### Answer
Access Control Validation.

### Runnable Code
```java
package basics;

class AccessDemo {
    public int a = 1;
    protected int b = 2;
    int c = 3; // Default
    private int d = 4;
    
    void print() {
        System.out.println(d); // Private OK same class
    }
}

public class AccessTest {
    public static void main(String[] args) {
        AccessDemo obj = new AccessDemo();
        System.out.println(obj.a);
        System.out.println(obj.b); // OK (Same package)
        System.out.println(obj.c); // OK (Same package)
        // System.out.println(obj.d); // Error Private
    }
}
```
