# 🧠 Java OOP & Fundamentals

## 1️⃣ OOP Pillars & Core Concepts

### Question 1: What are the 4 Pillars of OOP? Explain with real-world examples.

**Answer:**
1.  **Encapsulation:** Bundling data (variables) and methods together and hiding details.
    *   *Example:* A `Car` class hides engine details. You just press `accelerate()`.
2.  **Abstraction:** Hiding complexity and showing only essential features.
    *   *Example:* Remote control. You know buttons work, but not the circuit logic.
3.  **Inheritance:** One class acquires properties of another.
    *   *Example:* `class Dog extends Animal`. Dog inherits `eat()` from Animal.
4.  **Polymorphism:** One name, multiple forms.
    *   *Example:* `speak()` can be `Bark` (Dog) or `Meow` (Cat).

---

### Question 2: Difference between Abstract Class and Interface?

**Answer:**
| Feature | Abstract Class | Interface |
| :--- | :--- | :--- |
| **Methods** | Can have abstract & concrete. | Abstract, default, static (Java 8+). |
| **Variables** | `static`, `final`, non-final. | Always `public static final`. |
| **Constructor** | YES. | NO. |
| **Inheritance** | Single (`extends`). | Multiple (`implements`). |
| **Usage** | "Is-A" relationship (Dog is Animal). | "Can-Do" capability (Dog implements Runnable). |

---

### Question 3: What is Polymorphism? (Compile-time vs Runtime)

**Answer:**
*   **Compile-time (Static):** Method Overloading.
    *   Methods check signature (name + args) at compile time.
    *   `add(int a, int b)` vs `add(double a, double b)`.
*   **Runtime (Dynamic):** Method Overriding.
    *   JVM decides at runtime based on the *actual object type*.
    *   `Animal a = new Dog(); a.sound();` -> Calls Dog's sound.

---

### Question 4: Can you override `static` or `private` methods?

**Answer:**
*   **NO.**
*   **Static methods:** Belong to class, not instance. If you redefine a static method in child, it's called **Method Hiding**, not overriding.
*   **Private methods:** Not visible to child class, so cannot be overridden.

---

### Question 5: What is covariant return type?

**Answer:**
Since Java 5, an overriding method can return a **subtype** of the return type declared in the parent method.
```java
class A {
    Animal get() { return new Animal(); }
}
class B extends A {
    @Override
    Dog get() { return new Dog(); } // Valid (Dog is subtype of Animal)
}
```

---

### Question 6: Composition vs Inheritance. Which is better?

**Answer:**
*   **Inheritance ("Is-A"):** Example `Car extends Vehicle`. Tightly coupled. Fragile base class problem.
*   **Composition ("Has-A"):** Example `Car has-a Engine`. loosely coupled.
*   **Verdict:** Prefer **Composition** over Inheritance. It breaks dependency chains and allows changing behavior at runtime.

---

### Question 7: What is the `super` keyword?

**Answer:**
Refers to the **immediate parent class object**.
*   `super.variable`: Access parent variable.
*   `super()`: Call parent constructor (Must be 1st line).
*   `super.method()`: Call parent method (useful when overridden).

---

### Question 8: Significance of `this` keyword?

**Answer:**
Refers to the **current object instance**.
*   `this.x = x`: Distinguish instance var from param.
*   `this()`: Call another constructor of the same class (Constructor Chaining).
*   Pass `this` as parameter to other methods.

---

### Question 9: Can an Interface extend another Interface?

**Answer:**
**Yes.** An interface can `extend` multiple interfaces.
`interface A extends B, C {}`.

---

### Question 10: Difference between `Overloading` and `Overriding`?

**Answer:**
| Feature | Overloading | Overriding |
| :--- | :--- | :--- |
| **Scope** | Same Class. | Parent vs Child Class. |
| **Signature** | Must change (args). | Must be SAME. |
| **Return** | Can be different. | Must be same or covariant. |
| **Binding** | Static (Compile-time). | Dynamic (Runtime). |

---

## 2️⃣ Object Lifecycle & Constructors

### Question 11: Can a constructor be private? Why?

**Answer:**
**Yes.**
*   Used in **Singleton Pattern** to prevent direct instantiation.
*   Used in Factory classes implies you must use a static factory method.

---

### Question 12: Default Constructor vs No-Args Constructor?

**Answer:**
*   **Default Constructor:** Compiler inserts it **only if** no other constructor is defined.
*   **No-Args Constructor:** You manually write `public ClassName() {}`.
*   If you write `public ClassName(int x)`, the default constructor disappears.

---

### Question 13: What is "Constructor Chaining"?

**Answer:**
Calling one constructor from another.
*   Within same class: `this(args)`.
*   From child to parent: `super(args)`.
*   **Rule:** `this()` or `super()` must be the **first statement**. You cannot have both.

---

### Question 14: Use of `instanceof` operator?

**Answer:**
Checks if an object is an instance of a specific class or interface.
`if (obj instanceof String) ...`
*   **In Java 14+:** Pattern Matching:
    `if (obj instanceof String s) { System.out.println(s.length()); }` (No casting needed).

---

### Question 15: What is an Initialization Block?

**Answer:**
*   **Instance Block `{ }`:** Runs before constructor, for every object creation.
*   **Static Block `static { }`:** Runs **once** when class is loaded (ClassLoader). Used to init static variables.

---

## 3️⃣ Advanced OOP

### Question 16: Multiple Inheritance in Java?

**Answer:**
*   **Classes:** NO. (Diamond Problem ambiguity).
*   **Interfaces:** YES. Use `default` methods in Java 8 to resolve conflicts if needed (`InterfaceName.super.method()`).

---

### Question 17: What is a Marker Interface?

**Answer:**
An interface with **no methods**.
*   Examples: `Serializable`, `Cloneable`, `Remote`.
*   Used to signal the JVM or frameworks to treat the object specially.
*   *Modern approach:* Use **Annotations** instead.

---

### Question 18: Can abstract class have constructor?

**Answer:**
**Yes.**
*   Even though you can't `new AbstractClass()`, the constructor is used for initializing common variables when the *concrete subclass* is instantiated via `super()`.

---

### Question 19: Shallow Copy vs Deep Copy (Object Cloning)?

**Answer:**
*   **Shallow:** Copies field values. If field is reference, copies the *reference* (both objects point to same data). `Object.clone()` is shallow.
*   **Deep:** Creates new objects for referenced fields.
*   **How to Deep Copy?**
    1.  Override `clone()` and manually clone children.
    2.  Serialization/Deserialization.
    3.  Copy Constructor.

---

### Question 20: Immutable Class - How to create one?

**Answer:**
Example: `String`.
1.  Declare class `final` (cannot extend).
2.  Make all fields `private final`.
3.  No Setters.
4.  Initialize all fields in Constructor.
5.  If field is mutable (e.g., Date/List), return a **deep copy** in the getter, not the original reference.
