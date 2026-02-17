# Bonus: Design Patterns for Interviews

**Principle**: Service-based companies often ask for "Singleton", "Factory", or "Builder" to check architectural understanding.

## 1. Factory Pattern
**Type**: Creational.
**Use Case**: Creating objects without exposing instantiation logic.
**Example**: Bank Account Types.

```java
interface Account { void view(); }

class Savings implements Account {
    public void view() { System.out.println("Savings Account"); }
}
class Current implements Account {
    public void view() { System.out.println("Current Account"); }
}

class AccountFactory {
    public Account getAccount(String type) {
        if(type.equalsIgnoreCase("SAVINGS")) return new Savings();
        else if(type.equalsIgnoreCase("CURRENT")) return new Current();
        return null;
    }
}

public class FactoryDemo {
    public static void main(String[] args) {
        AccountFactory factory = new AccountFactory();
        Account acc = factory.getAccount("SAVINGS");
        acc.view();
    }
}
```

## 2. Builder Pattern
**Type**: Creational.
**Use Case**: Constructing complex objects step-by-step.
**Example**: User object with optional fields.

```java
class User {
    private String name; // Required
    private int age;     // Optional
    
    // Private Constructor
    private User(UserBuilder builder) {
        this.name = builder.name;
        this.age = builder.age;
    }
    
    // Static Builder Class
    public static class UserBuilder {
        private String name;
        private int age;
        
        public UserBuilder(String name) { this.name = name; }
        
        public UserBuilder age(int age) {
            this.age = age;
            return this;
        }
        
        public User build() {
            return new User(this);
        }
    }
}

public class BuilderDemo {
    public static void main(String[] args) {
        User u = new User.UserBuilder("John").age(25).build();
    }
}
```

## 3. Observer Pattern
**Type**: Behavioral.
**Use Case**: Notify subscribers of changes.
**Example**: Youtube Channel.

```java
import java.util.*;

interface Observer { void update(String msg); }

class Subscriber implements Observer {
    String name;
    Subscriber(String name) { this.name = name; }
    public void update(String msg) { System.out.println(name + " received: " + msg); }
}

class Channel {
    List<Observer> subs = new ArrayList<>();
    
    void subscribe(Observer s) { subs.add(s); }
    
    void upload(String title) {
        for(Observer s : subs) s.update("New Video: " + title);
    }
}

public class ObserverDemo {
    public static void main(String[] args) {
        Channel ch = new Channel();
        ch.subscribe(new Subscriber("Alice"));
        ch.subscribe(new Subscriber("Bob"));
        
        ch.upload("Java Design Patterns");
    }
}
```

## 4. Strategy Pattern
**Type**: Behavioral.
**Use Case**: Switch algorithms at runtime.
**Example**: Payment methods.

```java
interface Payment { void pay(int amount); }

class Card implements Payment {
    public void pay(int amount) { System.out.println("Paid " + amount + " via Card"); }
}
class UPI implements Payment {
    public void pay(int amount) { System.out.println("Paid " + amount + " via UPI"); }
}

class Cart {
    Payment paymentMethod;
    void setPayment(Payment p) { this.paymentMethod = p; }
    void checkout(int amount) { paymentMethod.pay(amount); }
}

public class StrategyDemo {
    public static void main(String[] args) {
        Cart cart = new Cart();
        cart.setPayment(new UPI());
        cart.checkout(100);
    }
}
```
