# 🧩 Java Extra Concepts Practice
Contains runnable code examples for missing concepts (Optional, Generics, Strategy, Abstract Factory).

## Question 1: When would you use `Optional`, and when should you avoid it?

### Answer
Return type for "missing value". Avoid in fields/methods args.

### Runnable Code
```java
package extra;

import java.util.Optional;

public class OptionalUsage {
    // Bad Practice: void method(Optional<String> s)
    // Good Practice: Return Optional
    static Optional<String> findUser(int id) {
        if (id == 1) return Optional.of("Dhruv");
        return Optional.empty();
    }

    public static void main(String[] args) {
        findUser(1).ifPresent(name -> System.out.println("Found: " + name));
        
        String user = findUser(2).orElse("Guest");
        System.out.println("User: " + user);
        
        // Avoiding NullPointerException
        // findUser(2).orElseThrow(() -> new RuntimeException("User not found"));
    }
}
```

---

## Question 2: Why are generics invariant in Java?

### Answer
List<String> is NOT List<Object>.

### Runnable Code
```java
package extra;

import java.util.ArrayList;
import java.util.List;

public class GenericInvariance {
    public static void main(String[] args) {
        List<String> strings = new ArrayList<>();
        strings.add("Hello");
        
        // List<Object> objects = strings; // Compile Error!
        // If this were allowed:
        // objects.add(100); 
        // Then strings.get(1) would crash (ClassCastException)
        
        System.out.println("Generics ensure you can't insert Integers into String lists via aliases.");
    }
}
```

---

## Question 3: Strategy Pattern real-world use case?

### Answer
Payment example. Interchangeable algorithms.

### Runnable Code
```java
package extra;

interface PaymentStrategy {
    void pay(int amount);
}

class CreditCard implements PaymentStrategy {
    public void pay(int amount) { System.out.println("Paid " + amount + " via Card"); }
}

class PayPal implements PaymentStrategy {
    public void pay(int amount) { System.out.println("Paid " + amount + " via PayPal"); }
}

class ShoppingCart {
    void checkout(int amount, PaymentStrategy strategy) {
        strategy.pay(amount);
    }
}

public class StrategyDemo {
    public static void main(String[] args) {
        ShoppingCart cart = new ShoppingCart();
        cart.checkout(100, new CreditCard());
        cart.checkout(50, new PayPal());
    }
}
```

---

## Question 4: Abstract Factory vs Factory Method?

### Answer
Family of objects.

### Runnable Code
```java
package extra;

// Products
interface Button { void paint(); }
interface Checkbox { void paint(); }

// Windows Family
class WinButton implements Button { public void paint() { System.out.println("Win Button"); } }
class WinCheckbox implements Checkbox { public void paint() { System.out.println("Win Checkbox"); } }

// Mac Family
class MacButton implements Button { public void paint() { System.out.println("Mac Button"); } }
class MacCheckbox implements Checkbox { public void paint() { System.out.println("Mac Checkbox"); } }

// Abstract Factory
interface GUIFactory {
    Button createButton();
    Checkbox createCheckbox();
}

// Concrete Factories
class WinFactory implements GUIFactory {
    public Button createButton() { return new WinButton(); }
    public Checkbox createCheckbox() { return new WinCheckbox(); }
}

class MacFactory implements GUIFactory {
    public Button createButton() { return new MacButton(); }
    public Checkbox createCheckbox() { return new MacCheckbox(); }
}

public class AbstractFactoryDemo {
    public static void main(String[] args) {
        GUIFactory factory = new WinFactory(); // Swappable
        factory.createButton().paint();
        factory.createCheckbox().paint();
    }
}
```

---

## Question 5: StackOverflowError Simulation.

### Answer
Infinite recursion.

### Runnable Code
```java
package extra;

public class SOError {
    static void recursive(int i) {
        recursive(i + 1);
    }
    
    public static void main(String[] args) {
        try {
            recursive(1);
        } catch (StackOverflowError e) {
            System.out.println("Stack Overflow caught!");
        }
    }
}
```

---

## Question 6: OutOfMemoryError Simulation.

### Answer
Heap full.

### Runnable Code
```java
package extra;

import java.util.ArrayList;
import java.util.List;

public class OOMError {
    public static void main(String[] args) {
        try {
            List<byte[]> list = new ArrayList<>();
            while (true) {
                list.add(new byte[1024 * 1024]); // 1MB allocation
            }
        } catch (OutOfMemoryError e) {
            System.out.println("Heap Full! OOM Caught.");
        }
    }
}
```
