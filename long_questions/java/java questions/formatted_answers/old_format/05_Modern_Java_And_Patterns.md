# 05. Modern Java and Design Patterns

**Q: Types of Inner Classes**
> "In Java, we have four types of inner classes:
> 1.  **Member Inner Class**: A non-static class defined inside another class. It has access to all members of the outer class, even private ones.
> 2.  **Static Nested Class**: Defined with `static`. It behaves like a normal top-level class but happens to be nested for packaging convenience. It *cannot* access instance variables of the outer class.
> 3.  **Local Inner Class**: Defined inside a method. It’s like a local variable—visible only within that method.
> 4.  **Anonymous Inner Class**: A class without a name, used for one-time instantiation (like creating an event listener on the fly)."

**Indepth:**
> **Memory Leaks**: Non-static inner classes (Member, Local, Anonymous) implicitly hold a reference to the outer class instance (`Outer.this`). This is a common source of memory leaks (e.g., in Android Activities) because the inner class prevents the outer class from being garbage collected.
>
> **Serialization**: Avoid serializing inner classes. Their synthetic fields (like the outer class reference) are compiler-dependent and can break serialization across different JVM versions.


---

**Q: Java Enums (More than just constants?)**
> "Yes! In C++, enums are just integers. In Java, **Enums are full-blown classes**.
>
> You can add methods, fields, and constructors to them. You can even implement interfaces or have abstract methods that each enum constant must implement. behavior.
>
> They are perfect for Singletons and State machines because they are thread-safe and immutable by default."

**Indepth:**
> **Under the Hood**: Enums are compiled into a class extending `java.lang.Enum`. The constants are `public static final` instances of the enum type.
>
> **Switch efficiency**: Enums in `switch` statements are optimized by the compiler using an internal jump table (`tableswitch` or `lookupswitch`), making them faster than string comparisons.


---

**Q: Java Records (Java 14+)**
> "**Records** are a game changer. They are concise, immutable data carriers.
>
> Instead of writing a class with private final fields, getters, `equals()`, `hashCode()`, and `toString()`, you just write: `public record Point(int x, int y) {}`.
>
> Java generates all that boilerplate for you. Use them whenever you just need to pass data around without behavior."

**Indepth:**
> **Canonical Constructor**: You can override the default constructor (compact constructor) to add validation logic without repeating the parameter list (`public Point { if (x < 0) throw ... }`).
>
> **Limitations**: Records are implicitly `final` and cannot extend other classes (state limitation), but they can implement interfaces.


---

**Q: Sealed Classes (Java 17+)**
> "**Sealed Classes** let you control *who* can extend your class.
>
> You use `sealed` and `permits` to strictly list the allowed subclasses. For example: `public sealed class Shape permits Circle, Square {}`.
>
> This gives you control over your inheritance hierarchy and allows the compiler to perform exhaustive checks in switch expressions (because it knows exactly which subclasses exist)."

**Indepth:**
> **Pattern Matching**: Sealed classes are the foundation for algebraic data types in Java. They enable compile-time exhaustiveness checking in switch expressions, meaning if you handle `Circle` and `Square`, the compiler knows you've handled *all* possible shapes and doesn't require a `default` case.


---

**Q: Text Blocks (Java 15+)**
> "Finally, no more ugly concatenation for multi-line strings!
>
> **Text Blocks** use triple quotes (`"""`). They preserve formatting and newlines.
>
> ```java
> String json = """
>   {
>     "name": "Java",
>     "type": "Language"
>   }
>   """;
> ```
> It’s a lifesaver for writing JSON, SQL, or HTML inside Java code."

**Indepth:**
> **Indentation**: The compiler determines the "essential whitespace" by looking at the position of the closing triple quotes (`"""`). Any indentation common to all lines (to the left of the closing quotes) is stripped automatically.
>
> **Escaping**: You don't need to escape double quotes `"` inside a text block, making HTML/JSON much cleaner.


---

**Q: Switch Expressions (Java 14+)**
> "The old switch statement was clunky and error-prone (forgetting `break;`).
>
> **Switch Expressions** can *return* a value directly. They use the arrow syntax (`->`) which doesn't need a break statement.
>
> ```java
> int days = switch (month) {
>     case JAN, MAR, MAY -> 31;
>     case FEB -> 28;
>     default -> 30;
> };
> ```
> It’s cleaner, safer, and follows functional programming styles."

**Indepth:**
> **Expression vs Statement**: An expression evaluates to a value (can be assigned to a variable). A statement just performs an action.
>
> **Yield**: If a case block needs multiple lines of logic, you use the `yield` keyword to return the value (e.g., `default -> { System.out.println("Processing"); yield 0; }`).


---

**Q: var keyword (Java 10+)**
> "**var** allows for local variable type inference.
>
> Instead of `Map<String, List<Integer>> map = new HashMap<>();`, you can type `var map = new HashMap<String, List<Integer>>();`.
>
> The compiler figures out the type from the right-hand side. It reduces verbosity without losing type safety (it's still strongly typed at compile time). Note: You can only use it for local variables, not fields or method parameters."

**Indepth:**
> **Non-Denotable Types**: `var` can hold types that you cannot explicitly write down, like the type of an anonymous class or an intersection type.
>
> **Readability**: Use `var` when the type is obvious (`var stream = list.stream()`). Avoid it when the type is obscure (`var result = service.process()`), as it forces the reader to jump to the method definition to understand what `result` is.


---

**Q: Core Functional Interfaces (Supplier, Consumer, etc)**
> "Java 8 introduced these standard interfaces in `java.util.function` so we don't have to create our own for every lambda:
>
> 1.  **Supplier<T>**: Takes nothing, returns something (`() -> T`). Used for lazy generation.
> 2.  **Consumer<T>**: Takes something, returns nothing (`T -> void`). Used for printing or saving.
> 3.  **Function<T, R>**: Takes T, returns R. Used for transformation (map).
> 4.  **Predicate<T>**: Takes T, returns boolean. Used for filtering."

**Indepth:**
> **Composition**: These interfaces have default methods for composition.
> *   `Function`: `andThen()`, `compose()`
> *   `Predicate`: `and()`, `or()`, `negate()`
> *   `Consumer`: `andThen()`
>
> **Primitives**: To avoid boxing overhead (int -> Integer), use specialized versions like `IntConsumer`, `LongPredicate`, `DoubleFunction`.


---

**Q: What is @FunctionalInterface?**
> "It’s an annotation that enforces the rule: 'This interface must have exactly *one* abstract method.'
>
> If you add a second abstract method, the compiler will yell at you. This ensures the interface can be used with Lambda expressions. (Default and static methods don't count towards the limit)."

**Indepth:**
> **SAM Type**: These interfaces are often called **SAM** (Single Abstract Method) types.
>
> **Lambdas**: The compiler uses *Lambda Metafactories* (invokedynamic) to instantiate these interfaces at runtime, which is more efficient than the legacy approach of creating anonymous inner class objects.


---

**Q: Singleton Pattern (Strategies)**
> "There are a few ways to implement a Singleton:
>
> 1.  **Eager Initialization**: Create the instance `static final instance = new Singleton()` at class load. Simple and thread-safe, but resources are used even if nobody asks for the instance.
> 2.  **Lazy Initialization**: Create it inside `getInstance()` if it's null. Not thread-safe by default.
> 3.  **Double-Checked Locking**: The 'pro' way. Check if null, synchronize, check if null again. High performance and thread-safe.
> 4.  **Enum Singleton**: The 'best' way. `public enum Singleton { INSTANCE; }`. It handles serialization and reflection attacks automatically."

**Indepth:**
> **Double-Checked Locking Issue**: Without `volatile` on the instance variable, Double-Checked Locking is broken because of instruction reordering (the instance reference might be assigned before the constructor finishes executing).
>
> **ClassLoaders**: A Singleton is only unique *per ClassLoader*. In complex environments (OSGi, Java EE), you might accidentally have multiple instances.


---

**Q: Factory Pattern**
> "The **Factory Pattern** is about decoupling object creation from usage.
>
> Instead of calling `new Car()` or `new Truck()` directly in your code, you call `VehicleFactory.getVehicle("type")`.
>
> This allows you to introduce new vehicle types later without changing the client code that uses them. It follows the Open/Closed principle."

**Indepth:**
> **Factory Method vs Abstract Factory**:
> *   Method Pattern: A method (usually abstract/overridden) creates the object.
> *   Abstract Factory: An interface that creates *families* of related objects (e.g., UI themes: createButton, createScrollbar) without specifying their concrete classes.


---

**Q: Builder Pattern**
> "The **Builder Pattern** is used when you have a complex object with many optional parameters.
>
> Instead of a constructor with 10 arguments (`new User("Bob", null, null, 25, true, ...)`), which is unreadable, you use a builder chain:
> ```java
> User u = User.builder()
>     .name("Bob")
>     .age(25)
>     .active(true)
>     .build();
> ```
> It’s readable and immutable."

**Indepth:**
> **Thread Safety**: The Builder itself is usually not thread-safe, but the object it builds should be immutable and thread-safe.
>
> **Fluent API**: The method chaining (`return this;`) creates a Domain Specific Language (DSL) feel. This pattern often involves a private constructor in the target class, forcing usage of the Builder.


---

**Q: Observer Pattern**
> "The **Observer Pattern** is a subscription mechanism. You have a 'Subject' (like a YouTuber) and 'Observers' (Subscribers).
>
> When the Subject changes state (uploads a video), it automatically notifies all Observers. Use this for event-driven systems like UI buttons or stock market updates."

**Indepth:**
> **Memory Leaks**: Observers must unregister themselves when no longer needed ("Lapsed Listener Problem"). WeakReferences can help here.
>
> **Push vs Pull**:
> *   Push: Subject sends data in the update method (`update(data)`).
> *   Pull: Subject just notifies (`update()`), and observer calls getters on Subject to fetch what it needs.


---

**Q: Java 8 Date/Time API (java.time) vs Legacy**
> "The old `java.util.Date` and `Calendar` were terrible—they were mutable (not thread-safe) and had confusing offsets (months started at 0!).
>
> The new **Java 8 API** (`LocalDate`, `LocalDateTime`, `ZonedDateTime`) is:
> 1.  **Immutable**: Thread-safe.
> 2.  **Clear**: Semantic methods like `plusDays(1)`.
> 3.  **Domain-Driven**: Separates 'Human time' (ISO dates) from 'Machine time' (Instant)."

**Indepth:**
> **Joda-Time**: This API was heavily influenced by the popular Joda-Time library.
>
> **Clock**: The API includes a `Clock` class. Use it for dependency injection to make testing time-dependent code easier (you can mock a `FixedClock` instead of relying on `System.currentTimeMillis()`).


---

**Q: Reference Types (Strong, Soft, Weak, Phantom)**
> "1.  **Strong Reference**: The default (`Object o = new Object()`). GC will *never* collect it as long as it's reachable.
> 2.  **Soft Reference**: GC collects it only if memory is running low. Good for caches.
> 3.  **Weak Reference**: GC collects it as soon as it sees it (if no strong refs exist). Used in `WeakHashMap`.
> 4.  **Phantom Reference**: Used to schedule post-mortem cleanup actions. Rarely used."

**Indepth:**
> **ReferenceQueue**: Weak and Phantom references can be registered with a `ReferenceQueue`. When the GC clears the reference, it puts the reference object into this queue, allowing your program to be notified and perform cleanup actions (this is how `WeakHashMap` expunges stale entries).


---

**Q: Statement vs PreparedStatement**
> "**Statement** is used for static SQL queries. It essentially concatenates strings. It is vulnerable to **SQL Injection** attacks if you insert user input directly.
>
> **PreparedStatement** is pre-compiled by the database. It uses placeholders (`?`) for parameters. It is faster (reused execution plan) and inherently secure against SQL Injection. Always use PreparedStatement."

**Indepth:**
> **Query Plan Cache**: The database parses, compiles, and optimizes the query plan. `PreparedStatement` allows the DB to reuse this plan for subsequent executions (even with different parameters), reducing CPU load on the DB server.


---

**Q: Transaction Management in JDBC**
> "By default, JDBC is in 'Auto-Commit' mode—every SQL statement is a transaction.
>
> To manage transactions manually:
> 1.  `connection.setAutoCommit(false);`
> 2.  Run your SQL updates.
> 3.  If all good: `connection.commit();`
> 4.  If error: `connection.rollback();`
>
> This ensures ACID properties (Atomicity)—either all updates happen, or none do."

**Indepth:**
> **Isolation Levels**: ACID relies on Isolation. JDBC supports identifying transaction isolation levels (`READ_COMMITTED`, `SERIALIZABLE`, etc.) to balance performance vs consistency (dirty reads, phantom reads).
>
> **Savepoints**: JDBC also supports `Savepoints`, allowing you to roll back part of a transaction to a specific point rather than rolling back the entire thing.

