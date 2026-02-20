# 06. OOP Basics

**Q: What are the 4 Pillars of OOP? Explain with real-world examples.**
> "The four pillars are essentially the rules for good object-oriented design:
>
> 1.  **Encapsulation**: 'Keep your secrets.' Bundling data and methods together and hiding the internal details.
>     *   *Real world:* A capsule. You know it's medicine, but you don't know the chemical formula inside.
>
> 2.  **Inheritance**: 'Parent and Child.' Creating new classes based on existing ones to reuse code.
>     *   *Real world:* A Child inherits traits (eye color) from their Parent.
>
> 3.  **Polymorphism**: 'Many forms.' One interface, multiple implementations.
>     *   *Real world:* A 'Person' can behave like a 'Student' in school, an 'Employee' at work, and a 'Customer' at a shop. Same person, different roles.
>
> 4.  **Abstraction**: 'Hiding complexity.' Showing only the essential features of an object.
>     *   *Real world:* Using a TV remote. You press 'Power' and it works. You don't need to know the circuit logic inside."

**Indepth:**
> **Cohesion vs Coupling**: Encapsulation increases Cohesion (classes do one thing well) and reduces Coupling (change in one class doesn't break another).
>
> **Liskov Substitution Principle**: Inheritance should follow LSP. A subclass must be substitutable for its superclass without breaking the application. If you override a method and throw a new checked exception or change its behavior drastically, you violate this principle.


---

**Q: Difference between Abstract Class and Interface?**
> "We covered this earlier, but to reinforce it:
>
> *   **Abstract Class**: Defines *identity* ('is-a'). Use it when classes share a common core but need specific implementations (e.g., `Dog` is an `Animal`). Can have state (fields) and constructors.
> *   **Interface**: Defines *capability* ('can-do'). Use it to define a contract that unrelated classes can implement (e.g., `Dog` and `Car` both implement `Moveable`). No state (until recently with static constants), multi-inheritance support."

**Indepth:**
> **Functional Interfaces**: An interface with exactly one abstract method is a Functional Interface and can be implemented using Lambda expressions (`() -> ...`). Abstract classes cannot be instantiated with Lambdas.
>
> **Constructors**: Abstract classes have constructors (called via `super()`), interfaces don't.


---

**Q: What is Polymorphism? (Compile-time vs Runtime)**
> "Polymorphism means 'many forms'.
>
> **Compile-time Polymorphism** (Static Binding) is **Method Overloading**.
> *   Same method name, different parameters. The compiler decides *which* method to call based on the arguments you pass.
>
> **Runtime Polymorphism** (Dynamic Binding) is **Method Overriding**.
> *   Same method signature in Parent and Child. The JVM decides *at runtime* which version to call based on the actual object type, not the reference type (`Animal a = new Dog(); a.sound()` calls Dog's sound)."

**Indepth:**
> **Dynamic Dispatch**: At runtime, the JVM uses the `vtable` (virtual method table) to look up the correct method to call. Overloaded methods are linked at compile time (Static Dispatch), so they are faster (no lookup).
>
> **Upcasting**: Polymorphism relies on Upcasting (`Parent p = new Child()`). You can access only the methods defined in `Parent`, but if `Child` overrides them, the `Child`'s version executes.


---

**Q: Can you override static or private methods?**
> "**No.**
>
> *   **static methods** belong to the *class*, not the instance. If you define a static method with the same signature in a subclass, it's called **Method Hiding**, not overriding. The parent's version remains if you access it via the parent class reference.
> *   **private methods** are invisible to the subclass. You can't override what you can't see. If you define a method with the same name in the subclass, it's just a completely new, unrelated method."

**Indepth:**
> **Virtual Methods**: In Java, all non-static, non-private, non-final methods are "virtual" by default, meaning they can be overridden.
>
> **Static Binding**: Static methods are bound at compile time based on the reference type. If you have `Animal a = new Dog(); a.staticMethod()`, it calls Animal's static method, not Dog's.


---

**Q: What is covariant return type?**
> "It sounds fancy, but it just means: When you override a method, the return type doesn't have to be *exactly* the same—it can be a **subclass** of the original return type.
>
> Example:
> Parent has `Animal produce()`.
> Child can override it with `Dog produce()`.
>
> This is allowed because a Dog *is* an Animal. It helps avoid type casting in client code."

**Indepth:**
> **Bridge Methods**: To support covariant return types (introduced in Java 5), the compiler generates a synthetic "bridge method" in the subclass with the original return type, which internally calls the new method. This maintains binary compatibility with older bytecode.


---

**Q: Composition vs Inheritance. Which is better?**
> "**Composition** is generally preferred over Inheritance ('Favor Composition over Inheritance').
>
> **Inheritance** (`extends`) creates a tight coupling. If the Parent class changes, the Child might break. It's an 'Is-A' relationship.
>
> **Composition** (having a private field of another class) is looser. It allows you to change behavior at runtime (by swapping the object) and doesn't expose the internal details. It's a 'Has-A' relationship."

**Indepth:**
> **Dependency Injection**: Composition is the basis of Dependency Injection (DI). Instead of hardcoding a dependency (`this.engine = new GasEngine()`), you pass it in (`this.engine = someEngine`). This makes testing easier because you can swap in a mock engine.
>
> **Fragile Base Class**: Heavy inheritance leads to the Fragile Base Class problem, where a change in a superclass causes unexpected bugs in subclasses. Composition avoids this.


---

**Q: What is the super keyword?**
> "**super** refers to the immediate parent class object.
>
> You use it to:
> 1.  Call the parent's constructor: `super(name)`.
> 2.  Call a parent's method that you have overridden: `super.printInfo()`.
> 3.  Access a parent's variable (rarely needed if encapsulated properly)."

**Indepth:**
> **Constructor Rule**: If you write a constructor in a child class, the very first line *must* be a call to `super()` or `this()`. If you don't write it, the compiler inserts `super()` (unexpectedly calling the no-arg parent constructor).


---

**Q: Significance of this keyword?**
> "**this** refers to the *current* object instance.
>
> You use it to:
> 1.  Distinguish instance variables from local variables (e.g., `this.name = name`).
> 2.  Call another constructor in the same class: `this(name, 0)`.
> 3.  Pass the current object as a parameter to another method."

**Indepth:**
> **Builder Pattern**: `this` is crucial for method chaining (Fluent Interface) where methods return `this`.
>
> **Inner Classes**: To access the outer class instance from an inner class, you use `OuterClassName.this`.


---

**Q: Can an Interface extend another Interface?**
> "**Yes.** An interface can extend multiple other interfaces.
>
> `public interface Robot extends Machine, Intelligent { ... }`
>
> This allows you to build complex contracts from smaller ones."

**Indepth:**
> **Interface Segregation Principle**: This feature supports ISP. Instead of creating one massive interface, you create small, specific interfaces. complex classes can then implement multiple of them, or a new interface can extend several of them to bundle capabilities.


---

**Q: Difference between Overloading and Overriding?**
> "**Overloading**: Same method name, *different* parameters (signature). Happens in the *same* class. Resolved at Compile-time.
>
> **Overriding**: Same method name, *same* parameters. Happens in *Parent-Child* classes. Resolved at Runtime."

**Indepth:**
> **Return Type**: You cannot overload a method *only* by changing the return type. The parameter list must change.
>
> **Exceptions**: Overridden methods cannot throw new or broader checked exceptions than the parent method (Liskov principle), but they can throw fewer or narrower exceptions. Overloaded methods have no such restrictions.


---

**Q: Can a constructor be private? Why?**
> "**Yes.**
>
> You make a constructor private to:
> 1.  Prevent instantiation from outside (Singleton Pattern).
> 2.  Prevent subclassing (you can't extend a class if you can't call its constructor).
> 3.  Force usage of static factory methods (`MyClass.create()`)."

**Indepth:**
> **Utility Classes**: `java.lang.Math` is a classic example. It has a private constructor because it only contains static methods. There is no point in creating an instance of `Math`.
>
> **Subclassing**: If a class has *only* private constructors, it cannot be subclassed (referenced in `extends`), effectively making it final.


---

**Q: Default Constructor vs No-Args Constructor?**
> "A **No-Args Constructor** is simply a constructor that takes no parameters. You can write one yourself.
>
> A **Default Constructor** is the no-args constructor that the **compiler** automatically inserts for you *only if* you haven't defined *any* other constructors.
>
> Once you write `public MyClass(int x)`, the default constructor vanishes. If you still want a no-args one, you must type it manually."

**Indepth:**
> **Serialization Issue**: If a Serializable class extends a non-Serializable parent, the parent *must* have a no-args constructor (visible to the subclass), otherwise serialization fails with `InvalidClassException`. The JVM needs it to initialize the parent's fields during deserialization.


---

**Q: What is Constructor Chaining?**
> "It's the process of calling one constructor from another.
>
> *   `this()` calls another constructor in the *same* class.
> *   `super()` calls a constructor in the *parent* class.
>
> This happens automatically (the compiler inserts `super()` if you don't), ensuring that the Object is fully initialized from the top of the hierarchy (Object) down to the specific subclass."

**Indepth:**
> **Recursive Execution**: Constructors execute top-down (Parent first, then Child).
>
> **Risk**: Avoid calling overridable methods inside a constructor! If Parent constructor calls `print()`, and Child overrides `print()`, the Child's method runs *before* the Child's constructor variables are initialized, leading to bugs (e.g., accessing null fields).


---

**Q: Use of instanceof operator?**
> "It checks if an object is an instance of a specific class (or one of its subclasses/interfaces).
>
> `if (animal instanceof Dog)` returns true if `animal` is actually a Dog.
>
> In modern Java (14+), you can use 'Pattern Matching for instanceof' to check and cast in one step:
> `if (animal instanceof Dog d) { d.bark(); }`."

**Indepth:**
> **Implementation**: `instanceof` usage is generally a code smell indicating poor polymorphism. Instead of `if (obj instanceof Dog) bark() else if (obj instanceof Cat) meow()`, you should just call `obj.makeSound()` and let polymorphism handle it.
>
> **Null**: `null instanceof Anything` always returns `false`.


---

**Q: What is an Initialization Block?**
> "It's a block of code inside a class (surrounded by `{}`). It runs every time an object is created, *before* the constructor.
>
> There is also a **Static Initialization Block** (`static {}`), which runs only *once* when the class is loaded by the ClassLoader."

**Indepth:**
> **Order of Execution**:
> 1.  Static Blocks (Class Load time)
> 2.  Parent Constructor
> 3.  Instance Initializer Blocks (in order of appearance)
> 4.  Constructor Body
>
> **Copying**: The compiler actually copies the code from instance initializer blocks into *every* constructor (right after the `super()` call).


---

**Q: Multiple Inheritance in Java?**
> "Java does **not** support multiple inheritance of *classes* (A extends B, C) to avoid the 'Diamond Problem' (ambiguity if both B and C have the same method).
>
> However, Java **does** support multiple inheritance of *interfaces* (A implements B, C). Since Java 8 (default methods), the Diamond Problem can occur with interfaces too, but the compiler forces you to override the conflicting method to resolve the ambiguity."

**Indepth:**
> **Diamond Problem Solution**: If interface B and C both define `default void run()`, and A implements both, the compiler forces A to override `run()`. Inside A, you can choose which one to call using `B.super.run()`.


---

**Q: What is a Marker Interface?**
> "It’s an empty interface with no methods (like `Serializable`, `Cloneable`, `Remote`).
>
> It serves as a 'tag' or metadata. The JVM checks for this tag to enable special behavior (like allowing serialization).
>
> In modern Java, **Annotations** are often used instead of Marker Interfaces."

**Indepth:**
> **Performance**: `instanceof` checks against marker interfaces are extremely fast.
>
> **Modern replacement**: Use Annotations implies using Reflection, which is slower. Marker interfaces are part of the type system, allowing compile-time checks (e.g., a method taking `Serializable` argument).


---

**Q: Can abstract class have constructor?**
> "**Yes.**
>
> Even though you can't create an instance of an abstract class directly (`new AbstractClass()` is illegal), the constructor is still needed to initialize the fields defined in the abstract class. It is called by the subclass constructor using `super()`."

**Indepth:**
> **Chain of Responsibility**: When you instantiate a `Child`, the runtime allocates memory for the *entire* object, including Parent fields. The `super()` call initializes those parent fields. Without it, the object would be partially uninitialized logic-wise.


---

**Q: Shallow Copy vs Deep Copy (Object Cloning)?**
> "Refers to the `clone()` method.
>
> The default `Object.clone()` performs a **Shallow Copy** (copies primitives, but copies *references* for objects).
>
> To get a **Deep Copy** (independent copy of everything), you must override `clone()` and manually clone the mutable fields, or use serialization/library tools."

**Indepth:**
> **Cloneable Interface**: To use `clone()`, your class *must* implement `Cloneable`, otherwise it throws `CloneNotSupportedException`. This is a weird design decision in Java (methods on Object, but requires interface).
>
> **Alternatives**: Copy Constructors (`public Car(Car other)`) are generally preferred over `clone()` because they are simpler to implement and don't require handling checked exceptions.


---

**Q: Immutable Class - How to create one?**
> "To make a class immutable (like String):
> 1.  Make the class `final` (so no subclasses can mess it up).
> 2.  Make all fields `private` and `final`.
> 3.  Provide only Getters, no Setters.
> 4.  If a field is a mutable object (like a Date or List), return a *copy* of it in the getter, not the original reference ('Defensive Copying')."

**Indepth:**
> **Thread Safety**: Immutable objects are inherently thread-safe. They require no synchronization.
>
> **Hash Keys**: Immutable objects make excellent HashMap keys because their `hashCode()` never changes. If a key object changes after being put in a Map, you won't be able to retrieve the value.

