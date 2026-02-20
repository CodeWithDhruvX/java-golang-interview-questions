# 10. Modern Java and Patterns (Practice)

**Q: Types of Inner Classes**
> "You should know these four types cold.
>
> First, there's the **Member Inner Class**. It's just a regular class inside another. It has a special bond with the outer class—it can access all its private members. But you need an instance of the outer class to create it, like `outer.new Inner()`.
>
> Second is the **Static Nested Class**. Even though it's inside, it behaves exactly like a normal top-level class. It doesn't have access to the outer class's instance variables, only static ones. It's usually nested just to keep things organized.
>
> Third is the **Local Inner Class**. This is defined *inside a method*. Theoretically, you can do this, but practically, it's rare. Its scope is limited to that method block.
>
> Finally, the most common: **Anonymous Inner Class**. You use this when you need to extend a class or implement an interface for a one-off object, like adding an `ActionListener` to a button. You define it and instantiate it at the exact same moment."

**Indepth:**
> **Scope**: Local inner classes can access local variables of the enclosing method *only if* they are final or effectively final. This is because the inner class instance might outlive the method execution, so it captures copies of the variables.


---

**Q: Java Enums (More than just constants?)**
> "Absolutely. In many languages, enums are just glorified integers. But in Java, they are **full-fledged classes**.
>
> This means an Enum can have its own instance variables, constructors, and methods. You can even have abstract methods in the Enum that each constant must implement specifically.
>
> For example, you could have an `Operation` enum with constants `PLUS`, `MINUS`, and `MULTIPLY`, where each one implements an abstract `apply(int a, int b)` method differently. They are also singletons by design, making them super powerful."

**Indepth:**
> **Strategy Pattern**: Enums with abstract methods are a concise way to implement the Strategy Pattern. Instead of creating multiple implementation classes, you just define the behavior in the enum constants.


---

**Q: Java Records (Java 14+)**
> "Records are basically 'named tuples' for Java. They solve the problem of writing too much boilerplate for data-carrier classes.
>
> Before records, if you wanted a simple 'Person' object, you wrote fields, getters, `equals`, `hashCode`, and `toString`. That’s 50 lines of clutter.
> with a Record, you just write `public record Person(String name, int age) {}`. One line. Java generates all that other stuff for you, and it makes the class immutable by default. Use them whenever you just need to pass data around."

**Indepth:**
> **Components**: The component fields are private and final. The accessors are named `name()` and `age()`, not `getName()`. Records also provide a compact constructor format.


---

**Q: Sealed Classes (Java 17+)**
> "Sealed classes give you control over your inheritance hierarchy.
>
> Normally, if a class is public, anyone can extend it. Sometimes you don't want that. You want to say: 'This is a Shape class, and I *only* want Circle and Square to extend it, nothing else.'
>
> You allow this by adding `sealed` to the class definition and using `permits` to list the allowed subclasses. It helps modeling restricted domains and allows the compiler to be smarter about checking all possible cases."

**Indepth:**
> **Exhaustiveness**: When using sealed classes in a switch expression, you don't need a `default` case if you cover all permitted subclasses. This makes adding a new subclass "safe" because the compiler will force you to update all switches.


---

**Q: Text Blocks (Java 15+)**
> "Text Blocks are a huge quality of life improvement.
>
> Before Java 15, if you wanted to write a big chunk of SQL or JSON in your code, you had to use endless `\n` characters and `+` signs to concatenate strings. It was unreadable.
>
> Now, you just use three quotes `"""` to start and end a block. You can paste your JSON or HTML right in there, formatted exactly how you want it, and Java preserves the newlines and indentation. It makes the code much cleaner."

**Indepth:**
> **Formatting**: You can use `\` at the end of a line to suppress the newline, allowing you to format code nicely in the editor but keep it as a single long string in the variable.


---

**Q: Switch Expressions (Java 14+)**
> "This is the modern upgrade to the old `switch` statement.
>
> The key difference is that a switch *expression* returns a value. You can assign the result of the switch directly to a variable: `var result = switch(day) { ... };`.
>
> It also uses the new arrow syntax `->`. The best part? No fall-through! You don't need to write `break;` at the end of every case anymore, so mistakes are much rarer."

**Indepth:**
> **Scope**: Variables defined inside a switch case block are now scoped correctly if you use `{}` blocks, avoiding name collisions between cases.


---

**Q: var keyword (Java 10+)**
> "This is for local variable type inference.
>
> Instead of typing `ArrayList<String> list = new ArrayList<String>();`, which repeats information, you can just type `var list = new ArrayList<String>();`.
>
> The compiler looks at the right side and figures out that `list` must be an ArrayList. It cuts down on verbosity. Just remember: Java is still strongly typed. `list` is still an ArrayList, and you can't put an integer into it later."

**Indepth:**
> **Polymorphism**: `var` infers the specific type, not the interface. `var list = new ArrayList<>()` infers `ArrayList`, not `List`. Pass `var` variables to methods expecting the interface type works fine due to polymorphism.


---

**Q: Core Functional Interfaces (Supplier, Consumer, etc)**
> "These are the standard interfaces in `java.util.function` that make Lambda expressions work. You don't need to invent your own interfaces anymore.
>
> *   **Supplier**: 'I provide something.' It takes no arguments but returns a result. Like a factory.
> *   **Consumer**: 'I eat something.' It takes an argument and does something with it, returning nothing. Like printing to the console.
> *   **Function**: 'I transform something.' It takes an input, processes it, and returns an output. Like converting String to Integer.
> *   **Predicate**: 'I test something.' It takes an input and returns true or false. Like checking if a list is empty."

**Indepth:**
> **Streams**: These interfaces are the building blocks of the Stream API. `filter` takes a Predicate. `map` takes a Function. `forEach` takes a Consumer.


---

**Q: What is @FunctionalInterface?**
> "This annotation is a safeguard. You put it on an interface to tell the compiler: 'This interface is intended to have exactly ONE abstract method.'
>
> Having one abstract method is the requirement for using that interface with Lambda expressions. If you accidentally add a second abstract method, the compiler will error out, saving you from breaking your lambdas."

**Indepth:**
> **Object Methods**: Methods from `java.lang.Object` (toString, equals) do not count towards the abstract method limit.


---

**Q: Singleton Pattern (Strategies)**
> "The Singleton pattern ensures a class has only one instance. There are a few ways to do it.
>
> The simplest is **Eager Initialization**: you create the `static final` instance right when the class loads. Thread-safe and easy.
>
> The performance-conscious way is **Lazy Initialization**: you create it only when someone calls `getInstance()`. But you have to be careful with threads.
>
> The robust way is **Double-Checked Locking**: you check if it's null, lock the class, and check again.
>
> But honestly, the **best** way in Java is using an **enum**. `public enum Singleton { INSTANCE; }`. It handles serialization and thread-safety for you automatically."

**Indepth:**
> **Serialization**: For non-Enum singletons, you must implement `readResolve()` to return the existing instance, otherwise deserialization creates a new copy. Enums do this automatically.


---

**Q: Factory Pattern**
> "Imagine you have a `Car` class and a `Truck` class. Instead of using `new Car()` directly in your code, you ask a `VehicleFactory` to give you a vehicle.
>
> You say `VehicleFactory.getVehicle("car")`.
>
> This is useful because your code doesn't need to know the complex logic of *how* a car is created. Plus, if you invent a `FlyingCar` later, you just update the Factory, and your original code works without changes."

**Indepth:**
> **Static Factory Method**: `Calendar.getInstance()` is a classic example. It returns a `GregorianCalendar` (usually), but the client code just treats it as a `Calendar`.


---

**Q: Builder Pattern**
> "This pattern is a lifesaver when you have objects with lots of parameters, especially optional ones.
>
> Instead of a constructor call like `new Pizza(true, false, true, false, true)`, which is confusing (is that extra cheese or pepperoni?), you use a Builder.
>
> It looks like: `Pizza.builder().cheese(true).pepperoni(false).build();`. It reads like a sentence, making your code much more readable and maintainable."

**Indepth:**
> **Immutability**: Builders are excellent for creating immutable objects. The Builder collects the parameters, checks validity, and then the private constructor creates the final object.


---

**Q: Observer Pattern**
> "This is the backbone of event-driven programming. Think of it like a YouTube subscription.
>
> You have a 'Subject' (the channel) and 'Observers' (the subscribers). When the Subject does something interesting (uploads a video), it notifies all the Observers automatically.
>
> You see this everywhere in UI programming: Button clicks are essentially the Observer pattern."

**Indepth:**
> **Listeners**: In Swing or Android, `OnClickListener` is the Observer pattern. You register a callback (Behavior) to be executed when the Event (State change) happens.


---

**Q: Java 8 Date/Time API (java.time) vs Legacy**
> "The old `java.util.Date` was a mess. It was mutable (meaning not thread-safe) and had weird design choices, like months starting from 0 (January was 0).
>
> The new Java 8 usage of `java.time` fixes all that. It introduces **LocalDate**, **LocalTime**, and **ZonedDateTime**.
>
> They are **immutable**, which makes them thread-safe. They are semantic and easy to read. And months start from 1, as they should!"

**Indepth:**
> **Instant**: Represents a specific point on the timeline (GMT). `LocalDateTime` is "wall clock" time without a time zone. `ZonedDateTime` combines both.


---

**Q: Reference Types (Strong, Soft, Weak, Phantom)**
> "Java has different strengths of references that interact with the Garbage Collector differently.
>
> *   **Strong**: The default. `Object o = new Object()`. The GC will never touch it as long as you hold this reference.
> *   **Soft**: The GC will only wipe this if memory is running critically low. Great for caches.
> *   **Weak**: The GC will wipe this as soon as it runs, provided no strong references exist. Used for `WeakHashMap`.
> *   **Phantom**: The weakest. Used for tracking when an object is about to be collected, mostly for cleanup scheduling."

**Indepth:**
> **WeakHashMap**: Excellent for metadata caching. If the key (e.g., a Class object) is unloaded, the entry is automatically removed from the map.


---

**Q: Statement vs PreparedStatement**
> "Always, always prefer **PreparedStatement**.
>
> A regular **Statement** takes your query string and sends it to the DB. If you concatenate user input into that string, you are wide open to SQL Injection attacks.
>
> **PreparedStatement** pre-compiles the SQL query structure first, and then treats user input strictly as data values, not executable code. It’s safer (no injection) and generally faster because the database can reuse the compiled query plan."

**Indepth:**
> **Bind Variables**: `PreparedStatement` sends the query structure *once*, then just sends the parameters. This saves parsing time on the DB side for repeated queries.


---

**Q: Transaction Management in JDBC**
> "A transaction is a group of operations that must succeed or fail as a unit.
>
> In JDBC, you manage this by first turning off the default behavior: `connection.setAutoCommit(false)`.
>
> Now, you run your multiple SQL updates. If everything looks good, you call `connection.commit()` to save it. If *anything* goes wrong (exception), you call `connection.rollback()` to undo everything. This ensures data integrity."

**Indepth:**
> **Spring**: Spring's `@Transactional` manages this boilerplate for you (opening connection, disabling auto-commit, committing/rolling back).

