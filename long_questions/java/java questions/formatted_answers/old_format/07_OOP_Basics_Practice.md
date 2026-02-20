# 07. OOP Basics (Practice)

**Q: What are the 4 Pillars of OOP? Explain with real-world examples.**
> "As a recap:
> 1.  **Encapsulation**: Protecting data. Example: A bank account. You can deposit/withdraw (methods), but you can't manually set the balance (data).
> 2.  **Inheritance**: Reusability. Example: 'Mobile Phone' inherits from 'Phone'. It has all phone features + smart features.
> 3.  **Polymorphism**: Flexibility. Example: 'Speak'. A human speaks, a dog barks, a cat meows. Same action, different implementation.
> 4.  **Abstraction**: Simplicity. Example: Driving a car. You use the steering wheel (interface) without knowing how the engine combustion works (implementation)."

**Indepth:**
> **Key Differentiator**: If asked "Which is most important?", argue for **Polymorphism**. It is the core of flexible, testable, and maintainable code (enabling Mocking, Dependency Injection, and Strategy Patterns). Without it, OOP is just data structures with methods.


---

**Q: Difference between Abstract Class and Interface?**
> "*Practice Perspective*:
> Use an **Abstract Class** when you are building a *framework* and want to provide common functionality that subclasses can re-use (like a `BaseRepository`).
> Use an **Interface** when you want to define a specific *role* that different classes can play (like `Serializable` or `Comparable`)."

**Indepth:**
> **Decision Matrix**:
> *   Is-A relationship? -> Abstract Class.
> *   Can-Do relationship? -> Interface.
> *   Need state? -> Abstract Class.
> *   Need multiple inheritance? -> Interface.


---

**Q: What is Polymorphism? (Compile-time vs Runtime)**
> "In a coding interview, you'd say:
> 'Compile-time is **Overloading**—same method name, different args. The compiler picks the right one.'
> 'Runtime is **Overriding**—same method signature in Parent/Child. The JVM picks the right one based on the actual object.'
> One is static, one is dynamic."

**Indepth:**
> **Memory**:
> *   **Overloading**: Resolved at compile time. Fast. No lookup overhead.
> *   **Overriding**: Resolved at runtime via the vtable (Virtual Method Table). Slight performance cost, but negligible in modern JVMs.


---

**Q: Can you override static or private methods?**
> "No.
> **Static**: They are bound to the class. Re-declaring them hides the parent method (Method Hiding).
> **Private**: They are not visible. Re-declaring them creates a new method."

**Indepth:**
> **Wait, what about shadowing?**
> If you define a variable with the same name in the child class, it *shadows* the parent's variable. This is confusing and bad practice.
> Static methods are also "shadowed" (or hidden), whereas instance methods are "overridden".


---

**Q: What is covariant return type?**
> "It means an overriding method can return a *subclass* of the original return type.
> If `Animal.born()` returns `Animal`, then `Dog.born()` can return `Dog`. This saves you from casting the result."

**Indepth:**
> **Compatibility**: This was introduced in Java 5. Before that, you had to return the exact same type and declare variables more broadly. Covariance allows for cleaner, more specific client code.


---

**Q: Composition vs Inheritance. Which is better?**
> "Composition is usually better. It’s more flexible.
> Inheritance is 'white-box' reuse (you see the internals of the parent).
> Composition is 'black-box' reuse (you just use the public API of the component).
> Changes in a superclass propagate to subclasses (fragile base class problem), but changes in a component class rarely break the wrapper class."

**Indepth:**
> **LSP Violation**: Inheritance forces an "Is-A" relationship. If "Square extends Rectangle", you might break assumptions (setting width changes height). Composition ("Square has a Shape") avoids this semantic trap.


---

**Q: What is the super keyword?**
> "It accesses the parent class.
> Code snippet:
> ```java
> public Dog() {
>     super(); // Calls Animal() constructor
> }
> public void eat() {
>     super.eat(); // Calls Animal.eat()
>     System.out.println("Dog eating");
> }
> ```"

**Indepth:**
> **Constructor Rule**: `super()` must be the *first* statement in a constructor. Why? Because the parent must strictly be fully initialized before the child tries to access any inherited state.


---

**Q: Significance of this keyword?**
> "It refers to *this* instance.
> Mostly used to fix shadowing (when param name == field name), or to pass the current object to another helper method."

**Indepth:**
> **Fluent Interfaces**: Returning `this` at the end of a method (`return this;`) allows for method chaining: `obj.setName("A").setAge(10);`. Crucial for Builder patterns.


---

**Q: Can an Interface extend another Interface?**
> "Yes, using `extends`. And it can extend *multiple* interfaces:
> `interface Hero extends Human, Flyable, Strong { ... }`"

**Indepth:**
> **Why?**: Interfaces have no state and no constructors. Extending multiple interfaces just merges their method contracts. There is no risk of the "Diamond Problem" regarding state initialization.


---

**Q: Difference between Overloading and Overriding?**
> "Overloading = New inputs, same name (Static).
> Overriding = New logic, same signature (Dynamic)."

**Indepth:**
> **Annotation**: Always use `@Override`. It forces the compiler to check your work. If you typo the name, the compiler throws an error instead of silently creating a new method.


---

**Q: Can a constructor be private? Why?**
> "Yes. To stop people from saying `new MyClass()`.
> Mandatory for:
> 1.  Singletons.
> 2.  Utility classes (like `java.util.Collections`) where everything is static."

**Indepth:**
> **Reftection Attack**: Private constructors can still be called via Reflection (`setAccessible(true)`). To be truly safe, throw an exception inside the private constructor if it's called a second time.


---

**Q: Default Constructor vs No-Args Constructor?**
> "Default is the invisible one the compiler gives you.
> No-Args is one you write explicitly (`public Foo() {}`).
> If you write *any* constructor, the Default one is gone."

**Indepth:**
> **Trap**: If you add a parameterized constructor `MyClass(String s)`, the compiler *removes* the default no-args constructor. Any code doing `new MyClass()` will suddenly break.


---

**Q: What is Constructor Chaining?**
> "Calling `this(...)` or `super(...)` as the first line of a constructor.
> It ensures code reuse between constructors and guarantees proper initialization order."

**Indepth:**
> **Output Prediction**: In an interview, if they ask for the output of a chain of constructors, remember: **Parent First**. The Object class constructor finishes first, then the Parent constant, then the Child.


---

**Q: Use of instanceof operator?**
> "Checks type safety before casting.
> ```java
> if (obj instanceof String) {
>     String s = (String) obj; // Safe
> }
> ```"

**Indepth:**
> **Pattern Matching**: Since Java 14, use `if (obj instanceof String s)`. It avoids the separate casting line and scope pollution.


---

**Q: What is an Initialization Block?**
> "Code that runs before the constructor.
> ```java
> {
>     // Instance block
>     System.out.println("Object created");
> }
> ```
> Rarely used, usually we just put this logic in the constructor."

**Indepth:**
> **Order**:
> 1. Static blocks (once, when class loads).
> 2. Instance blocks (every object creation).
> 3. Constructor (every object creation, after instance blocks).


---

**Q: Multiple Inheritance in Java?**
> "Classes: No.
> Interfaces: Yes.
> Why? To avoid the Diamond Problem with state/implementation."

**Indepth:**
> **Default Methods**: With Java 8 default methods, you *can* inherit implementation from multiple interfaces. If two interfaces define the same default method, you **must** override it in your class to resolve the conflict.


---

**Q: What is a Marker Interface?**
> "An empty interface.
> `public interface Safe {}`
> It tells the code something special about the class. Like a sticker on a box saying 'Fragile'."

**Indepth:**
> **Modern**: Annotations (`@Entity`, `@Service`) are the modern replacement for marker interfaces. They carry more metadata (values) and are more flexible.


---

**Q: Can abstract class have constructor?**
> "Yes, to initialize its own fields. It runs when a subclass is created."

**Indepth:**
> **Can you call it?**: No, `new AbstractClass()` is a compile error. But the constructor *exists* and is called via `super()` from the concrete subclass.


---

**Q: Shallow Copy vs Deep Copy?**
> "Shallow: Copies the reference (pointer). Both objects point to the same data.
> Deep: Copies the actual data. Objects are independent."

**Indepth:**
> **Clone vs Copy Constructor**: Prefer Copy Constructors (`public Car(Car c)`) over `clone()`. `clone()` is broken, throws checked exceptions, and bypasses constructors.


---

**Q: Immutable Class - How to create one?**
> "Final class. Private final fields. No setters. Defensive copies for mutable fields.
> Example: `String`, `Integer`, `LocalDate`."

**Indepth:**
> **Benefits**: Immutable objects are thread-safe (no synchronization needed), excellent Map keys (hash code never changes), and failure-atomic (state never gets inconsistent).

