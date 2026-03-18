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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output of this code and why?"

**Your Response:** "The output is `B`. Here's what happens: even though we declare the variable as type `A`, the actual object created is `new C()`. At runtime, Java uses dynamic method dispatch to find the right implementation. Since `C` doesn't override `name()`, it inherits from its parent `B`. `B` does override the method, so that's the one that gets called. This demonstrates polymorphism - the method call is resolved at runtime based on the actual object type, not the reference type."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What does this code output and how does `super` work here?"

**Your Response:** "The output is `... Woof`. The key here is understanding `super.method()` - it allows us to call the parent class implementation from within an overriding method. So when `Dog`'s `sound()` method calls `super.sound()`, it gets the `Animal`'s implementation which returns `...`, then `Dog` adds ` Woof` to it. This is useful when you want to extend the parent's behavior rather than completely replace it. It's a common pattern in inheritance hierarchies."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this code compile and why or why not?"

**Your Response:** "This doesn't compile - it's a compile error. The issue is with access modifiers during method overriding. The rule is: an overriding method cannot have a more restrictive access modifier than the method it's overriding. Here, the parent has `public void show()` but the child tries to override with `private void show()`. This violates the Liskov Substitution Principle - if you could use a `Child` object wherever a `Parent` is expected, but the `Child` makes the method less accessible, that would break the contract. You can only widen access (like `protected` to `public`), never restrict it." 

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this code compile? What's the rule with exceptions in overriding?"

**Your Response:** "This doesn't compile - it's a compile error. The rule for exceptions in method overriding is that the overriding method can only throw the same or more specific (narrower) checked exceptions. Here, the parent method throws `IOException`, but the child tries to throw `Exception`, which is broader. This makes sense from a client's perspective - if I'm calling the method through a parent reference, I only need to handle `IOException`. If the child could throw a broader exception, that would break the contract and catch clients off guard. The child could throw no exception, or a more specific one like `FileNotFoundException`, but never a broader one."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why doesn't polymorphism work here?"

**Your Response:** "The output is `P static`. This is a classic gotcha - static methods are not polymorphic! They're hidden, not overridden. The key difference is that static methods belong to the class, not to instances. When we call `obj.greet()`, Java resolves this at compile time based on the declared type of `obj`, which is `P`. It doesn't look at the actual runtime type. This is why static methods should be called using the class name like `P.greet()` or `C.greet()` - it makes the behavior explicit and avoids confusion."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why is this dangerous?"

**Your Response:** "The output is `Derived.display: 0` - and this is actually a very dangerous Java gotcha! Here's what happens: when we create a `Derived` object, the `Base` constructor runs first. The `Base` constructor calls `display()`, but because of polymorphism, this calls the `Derived` version of `display()`. However, at this point, `Derived`'s field initialization hasn't happened yet - `value` is still its default value of 0. This is why you should never call overridable methods from constructors - it can lead to subtle bugs where objects appear to be in an inconsistent state during construction. The fix would be to make `display()` final or private, or restructure the initialization logic."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why doesn't the child method get called?"

**Your Response:** "The output is `Base secret`. This demonstrates that private methods are not inherited and therefore cannot be overridden. The `Child.secret()` method is actually a completely new method, not an override. When `Base.callSecret()` calls `secret()`, it's calling its own private method. Private methods are not visible to subclasses, so they don't participate in polymorphism. If you wanted to allow overriding, you'd need to make the method at least `protected` or `public`. This is a good example of how access modifiers affect inheritance behavior."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile? What's the issue with constructors here?"

**Your Response:** "This doesn't compile - it's a compile error. The issue is with constructor chaining. In Java, if you don't explicitly call `super()` or `this()` as the first statement in a constructor, the compiler automatically inserts a call to the no-argument `super()` constructor. Here, `Dog()` doesn't call the parent constructor, so Java tries to insert `super()`, but `Animal` doesn't have a no-argument constructor - it only has `Animal(String)`. The fix is to explicitly call `super("some name")` as the first line in the `Dog` constructor. This shows that constructors are not inherited and you must explicitly chain to parent constructors."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why don't fields behave like methods?"

**Your Response:** "The output is `1` then `2`. This shows that fields in Java are not polymorphic - they don't participate in runtime polymorphism like methods do. When we access `obj.x`, Java resolves this at compile time based on the declared type of `obj`, which is `Parent`. So we get `Parent.x` which is 1. When we cast to `Child` and access `((Child)obj).x`, we get the `Child`'s field which is 2. This is different from method calls where the runtime type determines which method gets called. This is why it's generally recommended to make fields private and access them through getter methods - which do participate in polymorphism."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the lesson here?"

**Your Response:** "The output is `...`. This is a classic example of why the `@Override` annotation is so important. The `Cat` class has a method called `Speak()` with a capital S, but that's NOT an override of `Animal.speak()` - it's a completely new method because Java is case-sensitive. When we create a `Cat` and call `speak()` through an `Animal` reference, it calls the parent's method. If we had added `@Override` above `Speak()`, the compiler would have immediately caught this typo and told us that no method with that signature exists to override. This is why `@Override` should always be used - it catches these kinds of bugs at compile time rather than at runtime."

---

### 11. Covariant Return Type
**Q: Does this compile?**
```java
class Builder {
    protected String name;

    Builder setName(String name) {
        this.name = name;
        System.out.println("Builder setting name: " + name);
        return this;
    }
}

class AdvancedBuilder extends Builder {
    // Covariant return: returning AdvancedBuilder instead of Builder
    @Override
    AdvancedBuilder setName(String name) {
        this.name = name;
        System.out.println("AdvancedBuilder setting name: " + name);
        return this;
    }

    void specialAction() {
        System.out.println(name + " is performing an advanced action!");
    }
}

public class Main {
    public static void main(String[] args) {
        // Without covariant returns, this would require a cast to call specialAction()
        new AdvancedBuilder()
            .setName("Titan")      // Returns AdvancedBuilder
            .specialAction();      // Accessible because of the covariant return
    }
}
```
**A:** **Yes, compiles.** An overriding method can return a subtype of the original return type (Java 5+). Enables fluent builder APIs.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's this feature called?"

**Your Response:** "Yes, this compiles fine. This is called covariant return types, introduced in Java 5. It allows an overriding method to return a more specific type than the method it overrides. Here, the parent `Builder.setName()` returns `Builder`, but the child `AdvancedBuilder.setName()` can return `AdvancedBuilder`, which is a subtype. This is really useful for fluent APIs and builder patterns because it allows the chaining to continue with the more specific type. Without covariant returns, you'd have to cast the result or break the chain. This maintains type safety while enabling more expressive APIs."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does multiple interface implementation work?"

**Your Response:** "The output is 'flying' then 'swimming'. This shows that a class can implement multiple interfaces, which is Java's way around the no-multiple-inheritance limitation for classes. The `Duck` class implements both `Flyable` and `Swimmable`, so it must provide implementations for all methods from both interfaces. This gives us the benefits of multiple inheritance - the ability to have different types of behavior - without the complexity of multiple inheritance of state. Interfaces define contracts, and classes can fulfill multiple contracts simultaneously. This is a key feature that makes Java's type system flexible and expressive."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Java resolve diamond problems with default methods?"

**Your Response:** "The output is `B`. This demonstrates Java's rule for resolving the diamond problem with default methods. When there are multiple default methods available, Java chooses the most specific one. Here, `B` extends `A`, so `B` is more specific than `A`. The rule is: the most specific default method wins. If `B` didn't override the method, then `A`'s default would be used. This is different from the classic diamond problem in multiple inheritance because Java interfaces don't have state, only behavior. The resolution rules are clear and deterministic, avoiding ambiguity."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the diamond problem here?"

**Your Response:** "This doesn't compile - it's a compile error. This is the classic diamond problem, but with a twist. Unlike the previous example where `B` extended `A`, here `A` and `B` are unrelated interfaces that both provide a default method with the same signature. Java doesn't know which one to choose, so it forces the implementing class to resolve the ambiguity explicitly. The fix would be for class `C` to override `hello()` and either provide its own implementation or explicitly choose one of the defaults using `A.super.hello()` or `B.super.hello()`. This ensures that the developer makes the decision consciously rather than letting the compiler pick arbitrarily."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and how do static interface methods work?"

**Your Response:** "This doesn't compile - it's a compile error. Static interface methods, introduced in Java 8, are not inherited by implementing classes. They belong to the interface itself, not to the implementing classes. You must call them using the interface name: `MathOps.square(3)`. This is different from default methods, which are inherited. Static interface methods are essentially utility methods that are related to the interface but don't participate in polymorphism. They're a way to organize helper methods that are conceptually related to the interface without polluting the implementing class's namespace."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do abstract class constructors work?"

**Your Response:** "The output shows 'Vehicle: Toyota' then 'Car: Toyota'. This demonstrates that abstract classes can have constructors, and they're called when concrete subclasses are instantiated. The `Vehicle` constructor runs first, initializing the brand and printing the message, then the `Car` constructor runs. Even though you can't instantiate an abstract class directly, its constructor is still important for initializing the state that all subclasses will share. The subclass must explicitly call the parent constructor using `super()`. This is useful for setting up common state or enforcing invariants that all subclasses need."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what are the rules for abstract methods?"

**Your Response:** "This doesn't compile - it's a compile error. Abstract methods have strict rules: they cannot be `private` because that would make them unreachable to subclasses that need to implement them. They also cannot be `static` because static methods belong to the class, not to instances, and therefore can't be overridden. Abstract methods are meant to be implemented by subclasses, so they need to be accessible. The whole point of abstract methods is to define a contract that subclasses must fulfill, so making them private or static defeats that purpose. Abstract methods must be `public` or `protected` so that subclasses can actually implement them."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the purpose of sealed classes?"

**Your Response:** "Yes, this compiles fine. This is a sealed class, introduced in Java 17 as a preview feature. The `sealed` keyword restricts which classes can extend this abstract class. Only `Circle` and `Rectangle` are permitted to extend `Shape`. If anyone tried to create another class extending `Shape`, they'd get a compile error. This gives the class author control over their inheritance hierarchy. The main benefits are: first, it enables exhaustive pattern matching in switch statements - the compiler can verify that all possible subtypes are handled. Second, it makes the inheritance hierarchy explicit and maintainable. It's a way to have controlled extensibility rather than completely open inheritance."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are the rules for interface fields?"

**Your Response:** "The output is `100`. This demonstrates that all fields in interfaces are implicitly `public static final`, even if you don't specify those modifiers. So `int MAX = 100` is automatically treated as `public static final int MAX = 100`. This means interface fields are constants that are accessible to everyone, belong to the interface itself (not instances), and cannot be changed. This is why interfaces are often used to define constants - like `MAX_CONNECTIONS` or `DEFAULT_TIMEOUT`. The compiler enforces this, so you can't have mutable instance fields in interfaces, which is a key difference from abstract classes."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the issue with method visibility?"

**Your Response:** "This doesn't compile - it's a compile error. The issue is that interface methods are implicitly `public`, even if you don't write it. So `void draw()` in the interface is actually `public abstract void draw()`. When implementing an interface method, the implementing method must be at least as accessible as the interface method. Here, `Canvas.draw()` is package-private, which is more restrictive than `public`. This violates the contract because someone might call the method through an interface reference, expecting it to be accessible. The fix is to make the implementing method `public`. This ensures that interface implementations don't break the Liskov Substitution Principle."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and can you use this as a counter?"

**Your Response:** "This compiles, but it won't work as a mutable counter. The issue is that interface fields are implicitly `public static final`, so `int count = 0` is actually a constant. You can't modify it later. This highlights a key difference between interfaces and abstract classes: interfaces can't have mutable instance state. If you need to maintain state that varies between instances, you should use an abstract class instead. Interfaces are for defining contracts and behavior, while abstract classes can provide both contracts and shared state. This is a fundamental design decision - if you need instance fields, go with abstract class; if you just need a contract, use an interface."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about functional interfaces?"

**Your Response:** "The output is `10`. This demonstrates functional interfaces, which are a key feature of Java 8+. The `@FunctionalInterface` annotation marks an interface with exactly one abstract method. This allows us to use lambda expressions to implement the interface concisely. Instead of writing an anonymous class, we can just write `x -> x * 2`. The lambda automatically implements the single abstract method. This is the foundation of Java's functional programming support - it enables streams, method references, and many other modern Java features. Functional interfaces make code more readable and reduce boilerplate significantly."

---

### 23. @FunctionalInterface Enforcement
**Q: Does this compile?**
```java
@FunctionalInterface
interface BadFI { void doA(); void doB(); }
```
**A:** **Compile Error.** Two abstract methods violate the functional interface contract.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the purpose of @FunctionalInterface?"

**Your Response:** "This doesn't compile - it's a compile error. The `@FunctionalInterface` annotation enforces that an interface has exactly one abstract method. Here, we have two methods `doA()` and `doB()`, so it violates the contract. The annotation is optional but very useful - it catches these kinds of mistakes at compile time. Without the annotation, the interface would still compile but wouldn't be usable with lambdas. With the annotation, the compiler ensures the interface follows the functional interface rules. This is important because lambdas only work with functional interfaces - the compiler needs to know exactly which method the lambda should implement."

---

### 24. Private Interface Methods (Java 9+)
**Q: Does this compile?**
```java
// 1. Define the Interface
interface Logger {
    // Private method: Only accessible within this interface
    private void log(String msg) { 
        System.out.println("[LOG] " + msg); 
    }

    // Default methods: Inherited by classes that implement this interface
    default void info(String msg) { 
        log("INFO: " + msg); 
    }

    default void error(String msg) { 
        log("ERROR: " + msg); 
    }
}

// 2. Implement the Interface
class ConsoleLogger implements Logger {
    // No need to override info or error unless you want custom behavior
}

// 3. Main Class to run the program
public class Main {
    public static void main(String[] args) {
        ConsoleLogger myLogger = new ConsoleLogger();

        myLogger.info("System started successfully.");
        myLogger.error("Unable to connect to database.");
        
        // Note: myLogger.log("test") would cause a compilation error 
        // because the log method is private to the interface.
    }
}
```
**A:** **Yes, compiles** (Java 9+). Private interface methods allow code reuse between default methods without exposing the helper method.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the purpose of private interface methods?"

**Your Response:** "Yes, this compiles in Java 9 and later. Private interface methods were introduced to solve a specific problem: code reuse in default methods. Before Java 9, if you had multiple default methods that needed to share common logic, you'd have to either duplicate the code or make the helper method public (which pollutes the API). With private interface methods, you can have helper methods that are used by default methods but aren't visible to implementing classes. This keeps the interface clean while allowing for better code organization. It's a refinement to the default methods feature that makes interfaces more powerful for providing reusable behavior."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what design pattern does this demonstrate?"

**Your Response:** "The output shows 'reading CSV', 'processing CSV', then 'writing'. This is the Template Method pattern, one of the Gang of Four behavioral patterns. The abstract class `DataProcessor` defines the skeleton of an algorithm in the `process()` method - it calls `readData()`, then `processData()`, then `writeData()` in a fixed order. The algorithm structure is `final` so subclasses can't change it, but they can customize specific steps by implementing the abstract methods. This lets you have different implementations (like CSV, JSON, XML processors) that all follow the same processing workflow. It's great for code reuse and ensuring consistency across different implementations."

---

### 26. Cannot Instantiate Abstract Class
**Q: Does this compile?**
```java
abstract class Base { void doSomething() { System.out.println("concrete!"); } }
public class Main { public static void main(String[] args) { new Base(); } }
```
**A:** **Compile Error.** A class declared `abstract` cannot be instantiated even if it has no abstract methods.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and when can you instantiate abstract classes?"

**Your Response:** "This doesn't compile - it's a compile error. Even though the `Base` class has no abstract methods, it's still declared as `abstract`, which means you cannot instantiate it directly. The `abstract` keyword marks a class as incomplete - it's meant to be extended by subclasses. You can only instantiate concrete subclasses that extend the abstract class. This is useful when you want to provide a base implementation but prevent direct instantiation, or when you plan to add abstract methods later. The rule is: if a class is abstract, you cannot create instances of it, regardless of whether it has abstract methods or not."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the rule for implementing abstract methods?"

**Your Response:** "This doesn't compile - it's a compile error. When a class extends an abstract class, it has two choices: either implement all the abstract methods from the parent, or declare itself as `abstract`. Here, `Car` implements `move()` but doesn't implement `stop()`, so it violates the contract. The compiler forces you to either implement `stop()` or declare `Car` as abstract, which would pass the responsibility to subclasses. This ensures that concrete classes are complete implementations. It's Java's way of enforcing that abstract contracts are fulfilled by concrete classes."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a marker interface?"

**Your Response:** "The output is `true`. This demonstrates marker interfaces, which are interfaces with no methods. The `Serializable` interface doesn't define any behavior - it just marks that a class can be serialized. It's like a tag that tells the JVM 'this class is eligible for serialization'. The compiler and runtime can then check `instanceof Serializable` to determine if an object can be serialized. Other examples of marker interfaces include `Cloneable` and `Remote`. They're a way to add metadata to classes without using annotations. Before annotations were introduced in Java 5, marker interfaces were the primary way to add this kind of metadata."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are anonymous classes used for?"

**Your Response:** "The output is 'Hello from anon!'. This shows anonymous classes, which allow you to create one-off implementations without defining a separate class. Here we're creating an implementation of `Greeter` right where we need it. Anonymous classes are useful when you need a small, single-use implementation. However, for functional interfaces like this one, modern Java would typically use a lambda expression instead: `Greeter g = () -> System.out.println("Hello from anon!");`. Lambdas are more concise and readable. Anonymous classes are still useful when you need to implement interfaces with multiple methods or when you need to add fields or additional methods."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Java choose between default methods and class methods?"

**Your Response:** "The output is 'Hello from Class'. This demonstrates the rule that class implementations always win over default interface methods. Even though `Greeter` provides a default implementation, `CustomGreeter` overrides it with its own version. This makes sense from a design perspective - the class implementation should be more specific and take precedence. This is part of Java's method resolution rules for default methods. The hierarchy is: class method beats default interface method, and more specific interface default method beats less specific one. This ensures that developers can always override default behavior when needed."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and how does interface inheritance work?"

**Your Response:** "Yes, this compiles fine. This shows that interfaces can extend multiple interfaces, which is different from classes that can only extend one class. The `Duck` interface inherits from both `Flyable` and `Swimmable`, so it effectively has three methods: `fly()`, `swim()`, and its own `quack()`. The implementing class `Mallard` must provide implementations for all three methods. This is useful for composing interfaces - you can build up more specific interfaces by combining more general ones. It's a powerful feature that allows for flexible type hierarchies without the complexity of multiple inheritance of state."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the rule for constructor calls?"

**Your Response:** "This doesn't compile - it's a compile error. The rule is that `super()` or `this()` must be the very first statement in a constructor. Here, we have a `System.out.println()` call before `super()`, which violates this rule. The reason for this rule is that the parent class needs to be properly initialized before the child class can do any work. The parent constructor sets up the inherited state, and the child class might depend on that state. By enforcing this order, Java ensures objects are constructed in a safe, predictable way. The fix would be to move the print statement after the `super()` call."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does constructor chaining work?"

**Your Response:** "The output is 'Point(0, 0)'. This demonstrates constructor chaining using `this()`. When we call the no-argument constructor `new Point()`, it immediately calls `this(0, 0)`, which chains to the two-argument constructor. This avoids code duplication - the actual initialization logic is in one place, and other constructors delegate to it. Constructor chaining is a common pattern where you have a 'master' constructor that does the real work, and other constructors just call it with default values. It's the constructor equivalent of method delegation and helps maintain DRY (Don't Repeat Yourself) principle in initialization code."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the issue here?"

**Your Response:** "This doesn't compile - it's a compile error. The problem is circular constructor invocation. The no-argument constructor calls `this(1)`, which calls the one-argument constructor, which calls `this()` back to the no-argument constructor. This creates an infinite loop. The compiler detects this pattern and prevents it because it would cause a stack overflow at runtime. This is similar to how the compiler detects circular method calls. The fix would be to break the cycle - maybe both constructors should call a common private method, or one should directly initialize the fields instead of delegating to the other."

---

### 35. Interface Cannot Be Instantiated
**Q: Does this compile?**
```java
interface Runnable {}
public class Main { public static void main(String[] args) { new Runnable(); } }
```
**A:** **Compile Error.** Interfaces cannot be instantiated. You can instantiate anonymous classes implementing the interface.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and how do you create instances of interfaces?"

**Your Response:** "This doesn't compile - it's a compile error. Interfaces cannot be directly instantiated because they're incomplete - they may have abstract methods without implementations. To create an object that implements an interface, you either need a concrete class that implements the interface, or you can use an anonymous class. For example: `Runnable r = new Runnable() { public void run() { System.out.println("running"); } };`. Or for functional interfaces, you can use a lambda: `Runnable r = () -> System.out.println("running");`. The interface defines the contract, but you need a concrete implementation to actually create an object."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's happening with the casting?"

**Your Response:** "Yes, this compiles and prints `true`. This demonstrates upcasting, which is casting from a subclass to a superclass. When we do `Animal a = d;`, we're treating a `Dog` as an `Animal`. This is always safe because a `Dog` IS-A `Animal` - it has all the capabilities of an `Animal` plus more. Upcasting is implicit in Java, so you don't need to write the cast operator. The `instanceof` check returns `true` because the object is still a `Dog` at runtime, even though the reference type is `Animal`. This is fundamental to polymorphism - you can treat specialized objects as their more general types."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What happens at runtime and what's the danger here?"

**Your Response:** "This throws a `ClassCastException` at runtime. This is downcasting - casting from a superclass to a subclass. The problem is that the variable `a` actually holds a `Dog` object, but we're trying to cast it to a `Cat`. Even though both `Dog` and `Cat` extend `Animal`, a `Dog` is not a `Cat`. Downcasting is potentially unsafe, so Java checks the actual type at runtime. If the cast is invalid, it throws `ClassCastException`. The safe approach is to use `instanceof` first: `if (a instanceof Cat) { Cat c = (Cat) a; }`. This is why downcasting requires explicit cast syntax - Java wants you to acknowledge the risk."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about this pattern?"

**Your Response:** "The output shows 'Dog: Woof' then 'Cat: Meow'. This demonstrates pattern matching with instanceof, introduced in Java 16 as a preview feature. The expression `if (a instanceof Dog d)` does two things: it checks if `a` is a `Dog`, and if so, it automatically casts `a` to `Dog` and stores it in the variable `d`. This eliminates the need for explicit casting and makes the code more concise and readable. It's called pattern matching because you're matching a pattern (type plus variable) rather than just doing a type check. This feature reduces boilerplate and makes instanceof checks more elegant."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's type erasure?"

**Your Response:** "The output is `true`. This demonstrates type erasure, which is a fundamental concept in Java generics. At compile time, `List<Integer>` and `List<String>` are different types that provide type safety. But at runtime, the generic type information is erased - both become just `ArrayList`. This is for backward compatibility with pre-Java 5 code that didn't have generics. The JVM only sees the raw type, not the generic parameters. This is why you can't use generics in certain runtime contexts like checking if something is a `List<String>` - the information just isn't there at runtime. Type erasure is also why generic arrays aren't allowed and why you sometimes get unchecked warnings."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what does the wildcard mean?"

**Your Response:** "Yes, this compiles and prints `6.0`. This demonstrates wildcard generics with an upper bound. The `? extends Number` means 'any type that extends Number' - so it could be `List<Integer>`, `List<Double>`, `List<Long>`, etc. This is great for reading because you know whatever is in the list is at least a `Number`, so you can call methods like `doubleValue()`. However, you can't add to such a list because the compiler doesn't know what specific type it accepts. This is the PECS principle: Producer Extends, Consumer Super. Since we're only reading from (producing) the list, we use extends."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this wildcard differ from the previous one?"

**Your Response:** "The output is `[1, 2]`. This shows the other side of the PECS principle - Consumer Super. The `? super Integer` means 'Integer or any supertype of Integer' - so it could be `List<Integer>`, `List<Number>`, or `List<Object>`. This is great for writing because we know the list can accept `Integer` values. Since we're writing to (consuming) the list, we use super. The key difference from extends is: with extends you can read but not write, with super you can write but the type you get back is Object. This wildcard is perfect for operations like adding elements to a collection."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a generic method?"

**Your Response:** "The output is 'hello' then '42'. This demonstrates a generic method. The `<T>` before the return type declares a type parameter for the method. This means the method can work with any type - the compiler infers the type based on the arguments. When we call `identity("hello")`, T becomes String, and when we call `identity(42)`, T becomes Integer. Generic methods are useful for creating utility methods that can work with different types while maintaining type safety. The identity function is a simple example - it just returns whatever you pass in, but with the same type."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a bounded type parameter?"

**Your Response:** "The output is '7' then 'banana'. This shows a bounded type parameter. The `<T extends Comparable<T>>` means that T can be any type that implements Comparable. This constraint allows us to call the `compareTo()` method on objects of type T. The compiler ensures that only types that satisfy this constraint can be used - so you can call `max()` with Integers, Strings, Doubles, etc., but not with types that don't implement Comparable. Bounded type parameters are powerful because they let you write generic code that works with a specific family of types while maintaining type safety."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the risk with raw types?"

**Your Response:** "This compiles with unchecked warnings and throws a ClassCastException at runtime. This shows the danger of raw types - using generics without type parameters. When we use `List list` instead of `List<String>`, we're opting out of generic type safety. The compiler lets us add both a String and an Integer to the same list, which violates type safety. The error only shows up at runtime when we try to cast the Integer to a String. Raw types exist for backward compatibility with pre-Java 5 code, but they should be avoided in modern code. Always use proper generic types to catch these errors at compile time rather than runtime."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and why can't you create generic arrays?"

**Your Response:** "This doesn't compile - it's a compile error. Generic arrays are not allowed in Java because of type erasure. At runtime, `List<String>[]` would become just `List[]`, and the JVM couldn't enforce that only `List<String>` objects are stored in the array. This would allow you to violate type safety by putting a `List<Integer>` into a `List<String>[]` without the compiler catching it. The workaround is usually to use `ArrayList<List<String>>` or create an array of the raw type and use casts, but both have trade-offs. This limitation exists to maintain type safety despite type erasure."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's a generic class?"

**Your Response:** "The output shows '(hi, 42)' then '(1, true)'. This demonstrates a generic class with multiple type parameters. The `Pair<A, B>` class can hold any two types - A and B are type parameters that are specified when you create an instance. When we create `new Pair<>("hi", 42)`, A becomes String and B becomes Integer. Generic classes are great for creating reusable data structures that maintain type safety. Instead of having separate classes for StringPair, IntPair, etc., you have one generic Pair class that works with any combination of types. This reduces code duplication while maintaining compile-time type checking."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Java choose between these methods?"

**Your Response:** "The output is 'long'. This demonstrates Java's overload resolution rules. When we call `show(5)` with an int literal, Java has to choose between the two overloaded methods. It prefers widening (int to long) over autoboxing (int to Integer). The rule is: widening conversions are preferred over boxing conversions. This is because widening is a simpler, more efficient conversion that doesn't involve object creation. If there were a method that took an int directly, that would be chosen first. This order matters: exact match > widening > boxing > varargs. Understanding these rules helps predict which method will be called in ambiguous situations."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does overloading differ from overriding?"

**Your Response:** "The output is 'Object'. This shows that method overloading is resolved at compile time, not runtime. Even though the actual object is a String, the variable is declared as Object, so the compiler chooses the Object version of describe(). This is different from overriding, where the method choice happens at runtime based on the actual object type. With overloading, the compiler decides which method to call based on the declared types of the variables. This is a key distinction: overloading = compile-time polymorphism, overriding = runtime polymorphism. Understanding this difference is crucial for predicting method behavior in complex inheritance hierarchies."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do varargs work?"

**Your Response:** "The output shows 'Object[]' then '3'. This demonstrates varargs (variable arity parameters). When you declare a method with `Object... args`, Java automatically converts the arguments into an array. So calling `inspect("a", "b", "c")` creates an Object array containing those three strings. The varargs parameter is always treated as an array inside the method. This is syntactic sugar - it's equivalent to calling `inspect(new Object[]{"a", "b", "c"})`. Varargs must be the last parameter and a method can have only one varargs parameter. They're useful for creating flexible APIs that can accept varying numbers of arguments."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's this feature called?"

**Your Response:** "Yes, this compiles fine. This is covariant return types, which we saw earlier but in the context of generics. Here, the parent `create()` returns an `Animal`, but the child `create()` can return a `Dog`, which is a subtype of `Animal`. This maintains the Liskov Substitution Principle - anywhere you expect an `Animal`, getting a `Dog` is fine. Covariant returns are useful for factory patterns and fluent APIs where you want to return more specific types from subclasses while maintaining the contract defined by the parent class. It enables more precise type information without breaking polymorphism."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do inner classes work?"

**Your Response:** "The output is 'x = 10'. This demonstrates non-static inner classes. The key point is that inner classes have an implicit reference to their outer class instance. This allows the Inner class to directly access the private field `x` from the outer Main class. To create an inner class instance, you first need an outer class instance, then use `outer.new Inner()`. Inner classes are useful when you have a class that's only meaningful in the context of another class - like a Node class inside a LinkedList, or an Iterator inside a collection class. They maintain a tight coupling with their outer class."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this differ from the previous example?"

**Your Response:** "The output is '5'. This shows a static nested class, which is different from a regular inner class. Static nested classes don't have an implicit reference to an outer instance, so they can be created without needing an outer class object. They're essentially regular classes that are nested for namespace organization. Because they're static, they can only access static members of the outer class. Static nested classes are useful when you want to group related classes together but don't need the tight coupling of inner classes. Think of them as helper classes that belong conceptually to the outer class but don't depend on it."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the rule about local variables?"

**Your Response:** "Yes, this compiles and prints '15'. This demonstrates anonymous classes and the 'effectively final' rule. The local variable `multiplier` can be accessed inside the anonymous class because it's effectively final - meaning it's never reassigned after initialization. Before Java 8, variables had to be explicitly declared `final`. Java 8 relaxed this to 'effectively final'. This restriction exists because the anonymous class instance might outlive the method call, so Java needs to ensure the variable won't change. If you uncommented the line that changes `multiplier`, it would no longer be effectively final and the code wouldn't compile."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between lambdas and anonymous classes regarding 'this'?"

**Your Response:** "The output is 'value: 42'. This shows a key difference between lambdas and anonymous classes. In a lambda, `this` refers to the enclosing class instance (Main), not to the lambda itself. Lambdas don't create a new scope - they're more like syntactic sugar. In contrast, in an anonymous class, `this` would refer to the anonymous class instance. This means lambdas are more lightweight and behave more like they're part of the enclosing class. This is one reason why lambdas are preferred over anonymous classes for functional interfaces - they have more natural semantics and don't create the confusion of multiple 'this' references."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are the key enum methods?"

**Your Response:** "The output shows 'true', 'MON', then '0'. This demonstrates enum basics. Enums are essentially singleton constants - each enum value is a unique instance. That's why `==` works safely for comparison. The `.name()` method returns the string name of the constant, and `.ordinal()` returns its zero-based position in the declaration. Enums are more type-safe than string constants or integers, and they provide built-in methods like these. They're great for representing fixed sets of values like days of the week, directions, or states in a state machine. The ordinal is useful for some algorithms, but it's generally better to use enum values directly rather than relying on their position."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do enum constructors work?"

**Your Response:** "The output is approximately '9.80'. This shows that enums are more than just constants - they're full classes that can have fields, constructors, and methods. Each enum constant (MERCURY, EARTH) is created by calling the constructor with the provided parameters. The constructor is called once for each enum constant when the enum is first loaded. Enums can also have regular methods like `gravity()`, and even static fields and methods. This makes enums powerful for representing domain concepts with associated data and behavior, much more expressive than simple constants."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these enum methods do?"

**Your Response:** "The output shows 'RED GREEN BLUE' on one line, then 'GREEN' on the next. This demonstrates two important enum methods: `.values()` returns an array containing all enum constants in declaration order, and `.valueOf()` converts a string name back to the corresponding enum constant. These methods are automatically added by the compiler. `.values()` is great for iterating over all possible values, and `.valueOf()` is useful for parsing enum values from configuration files or user input. If the string doesn't match any enum constant, `valueOf()` throws an IllegalArgumentException."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about enum switch statements?"

**Your Response:** "The output is 'Hot!'. This shows that enum switch statements have a special syntax - you don't need to prefix the case labels with the enum type. Instead of `case Season.SUMMER:`, you just write `case SUMMER:`. The compiler knows you're switching on a Season enum, so it infers the type. This makes enum switch statements cleaner and more readable. Also, the compiler can check that you've covered all enum constants, which gives you exhaustiveness checking - if you add a new season later, the compiler will warn you about missing cases."

---

### 59. Enum Cannot Be Extended
**Q: Does this compile?**
```java
enum Base { A, B }
enum Extended extends Base { C, D }
```
**A:** **Compile Error.** Enums implicitly extend `java.lang.Enum` and cannot be subclassed.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the rule about enum inheritance?"

**Your Response:** "This doesn't compile - it's a compile error. Enums cannot be extended because they implicitly extend `java.lang.Enum`. Since Java doesn't support multiple inheritance for classes, and enums already inherit from Enum, you can't extend them further. This is a fundamental limitation - enums are designed to be complete, self-contained sets of constants. If you need to add more values, you modify the original enum, rather than extending it. This ensures that enum types are always closed and complete sets, which is important for type safety and switch statement exhaustiveness checking."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how can enums implement interfaces?"

**Your Response:** "The output is '3.1416'. This shows that enums can implement interfaces, which is a powerful feature. Here, the `RegularShape` enum implements the `Shape` interface, so each enum constant must provide an implementation for the `area()` method. Each constant can have its own implementation - CIRCLE returns PI, SQUARE returns 1.0. This is great for creating families of related behaviors that are type-safe and have a fixed set of implementations. It's like having a strategy pattern built into the language - you get all the enum benefits (type safety, exhaustiveness) plus polymorphic behavior."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do abstract methods work in enums?"

**Your Response:** "The output shows '5' then '1'. This demonstrates that enums can have abstract methods, and each constant must provide an implementation. Here, the Operation enum declares an abstract `apply()` method, and each constant (PLUS, MINUS) provides its own implementation. This is essentially a built-in strategy pattern - each enum constant represents a different strategy. It's more type-safe than using separate classes because the compiler ensures all constants implement the required method. This pattern is great for defining families of related algorithms or behaviors."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about EnumSet?"

**Your Response:** "The output shows '5' then 'false'. This demonstrates EnumSet, which is a specialized Set implementation designed specifically for enums. It's extremely efficient because enums have a fixed, known set of values. EnumSet uses bit vectors internally, making operations very fast and memory-efficient. The `range()` method creates a set containing all enum constants from MON to FRI inclusive, so we get 5 elements. EnumSet is the preferred way to work with collections of enum values - much better than using regular HashSet with enums."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the builder pattern?"

**Your Response:** "The output is 'Alice, 30'. This shows the builder pattern, which is used to construct complex objects step by step. The static nested Builder class handles the construction, with fluent methods that return 'this' for chaining. The Person constructor is private, forcing use of the builder. This pattern is great for objects with many optional parameters or complex construction logic. It's more readable than telescoping constructors and provides better control over the object creation process. The builder pattern is widely used in modern Java libraries."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why are enums good for singletons?"

**Your Response:** "The output is '2'. This demonstrates the enum-based singleton pattern, which is considered the best way to implement singletons in Java. Enums are inherently thread-safe, serialization-safe, and reflection-safe. The JVM guarantees only one instance of each enum constant. Unlike other singleton approaches, you don't need to worry about double-checked locking, private constructors being bypassed by reflection, or serialization creating duplicates. This is why Joshua Bloch recommends enum singletons as the preferred approach."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are local classes?"

**Your Response:** "The output is 'Hello'. This shows a local class - a class defined inside a method. Local classes can access effectively final local variables from the enclosing method, just like anonymous classes. They're useful when you need a named class but only within the scope of one method. Local classes are less common than anonymous classes but can be useful when the class is complex enough to benefit from a name or when you need multiple instances of the same class within a method. They're essentially a scoped version of inner classes."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the rule about static members in inner classes?"

**Your Response:** "This doesn't compile in Java 8 through 15. Non-static inner classes cannot declare static members because they're associated with a specific instance of the outer class. Static members belong to the class itself, not to instances, so there's a conflict. The exception is `static final` constants, which are allowed. Java 16 relaxed this restriction and now allows static members in inner classes. This change was made for consistency and convenience, but in older Java versions, you'd need to use a static nested class instead."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the limitation with anonymous classes?"

**Your Response:** "Yes, this compiles, but it shows an important limitation of anonymous classes - they can only implement one interface or extend one class. Here we're only implementing interface A, so it's fine. If we needed to implement both A and B, we'd have to create a named class. This is a design trade-off - anonymous classes are great for simple, one-off implementations, but for complex cases with multiple interfaces, you need the full power of a regular class. This keeps anonymous classes simple and focused."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about this modern Java pattern?"

**Your Response:** "The output shows '78.54' then '12.00'. This demonstrates several modern Java features working together. We have a sealed interface that restricts implementations to Circle and Rectangle, records for immutable data carriers, and pattern matching in switch statements. The switch uses pattern matching to destructure the records - `case Circle c` extracts the Circle and binds it to variable c. Because the interface is sealed, the compiler can verify that all cases are covered, making the switch exhaustive. This is the future of Java - more concise, type-safe, and expressive code."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the hashCode contract?"

**Your Response:** "The output is `null`. This demonstrates a critical Java contract: if you override `equals()`, you must also override `hashCode()`. Here, we overrode `equals()` but not `hashCode()`, so two Item objects with the same name are considered equal, but they have different hash codes. When we put the first Item in the HashMap, it goes to one bucket. When we try to look up with a second equal Item, it goes to a different bucket and isn't found. The rule is: equal objects must have equal hash codes. This is a common bug that can cause subtle issues in collections."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about Objects.equals()?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates `Objects.equals()`, which is a null-safe utility method. Unlike calling `equals()` directly on a potentially null object (which would throw NPE), `Objects.equals()` handles null values gracefully - it returns true if both arguments are null, false if one is null, or delegates to the regular `equals()` method if both are non-null. This is extremely useful in real code where you need to compare objects that might be null. It's part of Java's defensive programming utilities."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the Comparable interface work?"

**Your Response:** "The output shows the sorted temperatures: '36.1 37.5 38.9'. This demonstrates the `Comparable` interface, which defines the natural ordering of objects. The `compareTo()` method returns a negative value if this object is less than the other, zero if equal, or positive if greater. Here, the Temp record implements `Comparable<Temp>`, so `Collections.sort()` can sort the list automatically. This is different from `Comparator`, which defines external ordering. `Comparable` is for when an object has a natural ordering, like numbers, strings, or dates."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Comparator differ from Comparable?"

**Your Response:** "The output is 'Banana Apple Cherry'. This demonstrates `Comparator`, which defines external ordering separate from the class itself. Unlike `Comparable` where the class defines its own natural ordering, `Comparator` lets you define multiple different orderings. Here, we're sorting by price using `Comparator.comparingDouble()`. The Product class doesn't need to implement any interface - we just provide a comparator when we need it. This is more flexible than `Comparable` because you can have many different comparators for the same class (sort by name, price, etc.)."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does string interning work?"

**Your Response:** "The output is `true`. This demonstrates string interning and the string pool. When we call `intern()`, Java puts the string in the string pool (a special memory area for string literals) and returns a reference to the pooled version. String literals like `"hello"` are automatically interned. So both `s1` and `s2` end up referring to the same object in the pool. This is why string equality with `==` works for literals. String interning saves memory and is used internally by Java for optimization."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the purpose of Optional?"

**Your Response:** "The output shows 'result' then 'default'. This demonstrates `Optional`, which is Java's way to handle potentially absent values without using null. `Optional.of()` wraps a value, while `Optional.empty()` represents absence. The `orElse()` method returns the value if present, or a default if empty. This is much safer than returning null from methods because it forces the caller to handle the absence case. Optional is part of Java's move toward null-safe programming and helps avoid NullPointerExceptions."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Optional.map work?"

**Your Response:** "The output is 'ALICE'. This demonstrates the `map()` method on Optional, which allows you to transform values if they're present. The `map()` method applies a function to the contained value and returns a new Optional with the result. If the Optional is empty, the function isn't called and the result is empty. This creates a fluent chain of operations: trim the string, convert to uppercase, and provide a default if any step fails. This is much cleaner than nested if-null checks and is a key part of functional programming in Java."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the purpose of requireNonNull?"

**Your Response:** "The output is 'Caught: owner cannot be null'. This demonstrates `Objects.requireNonNull()`, which is a defensive programming utility. It checks if the argument is null and throws a NullPointerException with a custom message if it is. This is useful for validating method parameters early, failing fast rather than letting null values propagate through your code and cause problems later. The custom message makes debugging easier. It's cleaner than writing `if (param == null) throw new NullPointerException()` manually and is widely used in modern Java APIs."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the issue with clone()?"

**Your Response:** "The output is `[1, 2, 3, 4]`. This demonstrates the danger of shallow copying. The default `clone()` method does a shallow copy - it copies the fields, but if a field is a reference to an object, both the original and clone reference the same object. When we modify the clone's items list, we're actually modifying the same list that the original references. For a deep copy, you'd need to clone the mutable fields as well. This is why many developers prefer copy constructors or serialization over the clone() method."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between instanceof and getClass()?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates the difference between `instanceof` and `getClass()`. The `instanceof` operator checks if an object IS-A certain type (including inheritance), so `new Dog() instanceof Animal` is true. But `getClass()` returns the exact runtime class, so `a.getClass() == Animal.class` is false because the actual object is a Dog. `getClass() == Dog.class` is true because it matches the exact type. Use `instanceof` for type checking with inheritance, and `getClass()` when you need exact type matching."

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

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's special about records?"

**Your Response:** "The output shows the point coordinates, the default toString representation, and the caught exception. This demonstrates records, which are immutable data carriers introduced in Java 16. Records automatically generate constructors, getters, equals(), hashCode(), and toString(). The compact constructor `Point { ... }` allows validation logic. Records are great for DTOs, value objects, and any data-holding classes. They reduce boilerplate code significantly while ensuring immutability. The validation in the compact constructor shows that records can still have business logic."

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

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Why is finalize() unreliable and what's the modern approach?"

**Your Response:** "The output is unpredictable because `finalize()` is unreliable. The garbage collector calls `finalize()` at an indeterminate time, or never at all. This makes `finalize()` unsuitable for resource cleanup. It was deprecated in Java 9 because it can cause performance issues, resurrect objects, and hide resource leaks. The modern approach is `try-with-resources` with the `AutoCloseable` interface, which provides deterministic cleanup when the try block exits. This is more predictable, efficient, and follows the RAII (Resource Acquisition Is Initialization) pattern."
