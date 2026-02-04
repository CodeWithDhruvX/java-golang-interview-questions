# 🚀 Modern Java, Patterns & Advanced Features

## 1️⃣ Inner Classes & Enums

### Question 58: Types of Inner Classes.

**Answer:**
1.  **Member Inner Class:** Non-static. Has reference to Outer (`Outer.this`). `new Outer().new Inner()`.
2.  **Static Nested Class:** Like a normal static class, just nested. No accessing Outer instance. `new Outer.Inner()`.
3.  **Local Inner Class:** Defined inside a method. Can access final local variables.
4.  **Anonymous Inner Class:** No name. Used for callbacks/listeners. `new Runnable() { ... }`.

---

### Question 59: Java Enums (More than just constants?).

**Answer:**
Enums in Java are full Classes.
*   Can have **Constructors, Methods, Fields**.
*   **Singleton:** `INSTANCE;` guarantees single instance (JVM handled).
*   Can implement Interfaces (but cannot extend classes, as they extend `Enum`).
```java
enum Status {
    ACTIVE(1), INACTIVE(0);
    int code;
    Status(int code) { this.code = code; } // Constructor is private
}
```

---

## 2️⃣ Modern Java (Java 8 to 21)

### Question 60: Java Records (Java 14+).

**Answer:**
Boilerplate killer for "Data Classes".
*   Auto-generates: Constructor, `equals`, `hashCode`, `toString`, Getters (named `name()`, not `getName()`).
*   **Immutable** (All fields final).
```java
record Point(int x, int y) {}
```

---

### Question 61: Sealed Classes (Java 17+).

**Answer:**
Restrict which classes can extend them.
*   Control inheritance hierarchy.
```java
public sealed class Shape permits Circle, Square {}
public final class Circle extends Shape {} // Must be final, sealed, or non-sealed
```

---

### Question 62: Text Blocks (Java 15+).

**Answer:**
Multi-line strings without `\n` concatenation.
```java
String json = """
    {
      "name": "Java"
    }
    """;
```

---

### Question 63: Switch Expressions (Java 14+).

**Answer:**
Enhanced switch. No `break` needed. Can return value.
```java
int days = switch(month) {
    case JAN, MAR -> 31;
    case FEB -> 28;
    default -> 30;
};
```

---

### Question 64: `var` keyword (Java 10+).

**Answer:**
Local Variable Type Inference.
*   `var list = new ArrayList<String>();`
*   Compiler infers type.
*   **Usage:** Only for local variables. Not for fields/params. Must be initialized.

---

## 3️⃣ Functional Interfaces (Java 8)

### Question 65: Core Functional Interfaces (`Supplier`, `Consumer`, etc).

**Answer:**
From `java.util.function`:
1.  **`Predicate<T>`:** `T -> boolean`. Test condition. (`filter`).
2.  **`Consumer<T>`:** `T -> void`. Action. (`forEach`).
3.  **`Supplier<T>`:** `() -> T`. Factory. (`generate`).
4.  **`Function<T,R>`:** `T -> R`. Transform. (`map`).
5.  **`UnaryOperator<T>`:** `T -> T`.
6.  **`BinaryOperator<T>`:** `(T, T) -> T`. (`reduce`).

---

### Question 66: What is `@FunctionalInterface`?

**Answer:**
Annotation to ensure an interface has **exactly one abstract method**.
*   Optional but recommended.
*   Can have multiple `default` or `static` methods.

---

## 4️⃣ Core Design Patterns (Interview Logic)

### Question 67: Singleton Pattern (Strategies).

**Answer:**
1.  **Eager:** static final field. (Created at class load).
2.  **Lazy:** Create in `getInstance()`. (Not thread-safe).
3.  **Synchronized:** Perf hit.
4.  **Double-Checked Locking:** `volatile` instance + synchronized block.
5.  **Bill Pugh:** Static Inner Holder (Lazy + Thread-safe).
6.  **Enum Singleton:** Best. (Safe from Reflection/Serialization).

---

### Question 68: Factory Pattern.

**Answer:**
Creates objects without exposing creation logic.
*   Interface `Shape` -> Classes `Circle`, `Square`.
*   `ShapeFactory.getShape("CIRCLE")` returns `new Circle()`.
*   Decouples client from implementation classes.

---

### Question 69: Builder Pattern.

**Answer:**
Constructs complex objects step-by-step.
*   Solved the "Telescoping Constructor" problem (constructor with 10 args).
*   `Student.builder().setName("A").setAge(20).build();`

---

### Question 70: Observer Pattern.

**Answer:**
One-to-Many dependency.
*   Subject (Publisher) notifies Obsersvers (Subscribers) of state change.
*   Usage: Event Listeners, RxJava.

---

## 5️⃣ Date/Time & Reference Types

### Question 71: Java 8 Date/Time API (`java.time`) vs Legacy.

**Answer:**
*   **Legacy (`Date`, `Calendar`):** Mutable (Not thread-safe). Zero-indexed months (Jan=0).
*   **Modern (`LocalDate`, `Instant`):** **Immutable**. Thread-safe.
    *   `LocalDate`: Date without time (Birthday).
    *   `LocalDateTime`: Date + Time.
    *   `ZonedDateTime`: Timezone aware.
    *   `Instant`: Timestamp (UTC).

---

### Question 72: Reference Types (Strong, Soft, Weak, Phantom).

**Answer:**
1.  **Strong:** Standard `obj = new Object()`. Never collected if reachable.
2.  **Soft:** Collected only if **Memory is full**. (Good for Caches).
3.  **Weak:** Collected on **Next GC** if only weakly reachable. (Good for Metadata/WeakHashMap).
4.  **Phantom:** Collected when physically removed from memory. (Scheduling cleanup).

---

## 6️⃣ JDBC Basics

### Question 73: `Statement` vs `PreparedStatement`.

**Answer:**
*   **Statement:** Compile SQL every time. Vulnerable to **SQL Injection**.
*   **PreparedStatement:** Pre-compiled. Cached. **Prevents** SQL Injection (uses placeholders `?`). Faster for repeated execution.

---

### Question 74: Transaction Management in JDBC.

**Answer:**
By default, JDBC is `autoCommit(true)`.
To manage transaction:
1.  `conn.setAutoCommit(false);`
2.  Perform operations (SQLs).
3.  `conn.commit();` // Success
4.  `catch (e) { conn.rollback(); }` // Failure

---
