# Core Java & Basics - Interview Answers

> 🎯 **Focus:** These answers are designed for *speaking*, not writing. They focus on clarity, context, and confidence.

### 1. What is the difference between JDK, JRE, and JVM?
"Think of them as a hierarchy of tools for Java development.

The **JVM** (Java Virtual Machine) is the engine that actually runs your code—it translates bytecode into machine code. It's platform-dependent, which is why we have different downloads for Windows, Mac, and Linux, but it makes Java itself platform-independent.

The **JRE** (Java Runtime Environment) is the JVM plus the core libraries needed to *run* applications. In production, we usually just need the JRE.

The **JDK** (Java Development Kit) is the full package. It includes the JRE plus development tools like the compiler (`javac`) and debugger. So, as a developer, I always install the JDK, but on a server, we might only have the JRE."

---

### 2. Difference between Abstract Class and Interface?
"Both define structure, but the main distinction comes down to **state** and **intent**.

An **Abstract Class** is like a 'partial blueprint.' It can have member variables (state), constructors, and fully implemented methods. I use it when classes share a common identity or core logic—like a `BasePayment` class that handles logging for all payment types.

An **Interface** is more like a 'contract' or capability—like `Runnable` or `Serializable`. Traditionally, it had no state. Even though Java 8 added default methods, an interface still can't hold instance variables.

Crucially, a class can implement **multiple** interfaces but extend only **one** abstract class. So I prefer interfaces for defining capabilities across different parts of the system, and abstract classes for shared implementation details."

---

### 3. String vs StringBuilder vs StringBuffer?
"It comes down to **immutability** and **performance**.

**String** is immutable. Every time you concatenate a string, you're actually destroying the old object and creating a brand new one. This is great for caching and security, but if you do it in a loop, it kills performance.

**StringBuilder** is mutable. You modify it in place, which is much faster for string manipulation. I always use this when building large strings, like CSVs or JSON responses.

**StringBuffer** is the legacy, thread-safe version of StringBuilder. All its methods are synchronized, which makes it slower. Honestly, I rarely use it because if I need thread safety, I usually handle it at a higher level, not at the string builder level."

---

### 4. Difference between `equals()` and `==`?
"This is a classic trap. `==` checks for **reference equality**—basically, do these two variables point to the exact same memory address?

`equals()` checks for **value equality**—do these objects contain the same data?

By default, the `Object` class implementation of `equals()` acts just like `==`. That's why we override it in our DTOs and entities to compared actual fields, like IDs or usernames.

A common bug I've seen is using `==` to compare Integers. It works for small numbers due to caching, but fails for larger numbers. So I basically always use `.equals()` for objects to be safe."

---

### 5. `final` vs `finally` vs `finalize`?
"They sound similar but are completely unrelated.

**final** is a modifier. A final variable is a constant, a final method can't be overridden, and a final class can't be inherited. I use it everywhere to ensure immutability.

**finally** is a block used with try-catch. It **always** executes, whether an exception happened or not. It used to be critical for closing streams, though nowadays `try-with-resources` handles that cleaner.

**finalize** is a deprecated method from the `Object` class. It was meant for cleanup before Garbage Collection, but it's unpredictable and causes performance issues, so we don't use it anymore."

---

### 6. Checked vs Unchecked Exceptions?
"**Checked Exceptions** (like `IOException`) are enforced at compile-time. The compiler forces you to handle them with a try-catch. They represent recoverable problems where the caller *should* do something—like asking the user to re-enter a filename.

**Unchecked Exceptions** (like `NullPointerException`) extend `RuntimeException` and happen at runtime. They usually indicate programming errors or unrecoverable states.

In modern frameworks like Spring, we heavily favor **Unchecked** exceptions. If the database is down, there's nothing the controller can really do about it, so forcing every method up the stack to declare `throws SQLException` just clutters the code."

---

### 7. What is serialization and `serialVersionUID`?
"Serialization is converting an object into a byte stream to save it or send it over a network.

When you do this, Java uses a `serialVersionUID` to version the class. If you serialize an object, then change the class definition (like adding a field), and try to deserialize it back, the JVM checks this ID. If they don't match, you get an `InvalidClassException`.

Honestly, in microservices, we rarely use Java's native serialization anymore because it's brittle. We prefer **JSON** or **Protobuf**, which are language-agnostic and don't rely on `serialVersionUID`."

---

### 8. Explain `static` keyword in Java.
"`static` means 'belongs to the class,' not to any individual instance.

A **static variable** is shared across all objects. If one object changes it, everyone sees the change.
A **static method** can be called without creating an object—like `Math.abs()` or `LocalDate.now()`.

I use static mostly for utility functions or constants. However, I avoid overusing it for business logic because static methods are hard to mock in unit tests, which can lead to tight coupling."

---

### 9. What are Wrapper Classes and Autoboxing?
"Wrapper classes let us treat primitive types (like `int`) as Objects (like `Integer`).

We need them because Java Collections like `ArrayList` can't store primitives—you can't have an `ArrayList<int>`.

**Autoboxing** is just the compiler automatically converting the primitive to the wrapper for you.
The catch is that Wrapper classes can be **null**, whereas primitives cannot. I've seen `NullPointerExceptions` happen when unboxing a null Integer, so I prefer primitives for local calculations and only use Wrappers when interacting with Collections or DB entities."

---

### 10. What is type erasure in Generics?
"It means that Generics are a **compile-time** safety feature; they don't effectively exist at runtime.

When you write `List<String>`, the compiler ensures you only put Strings in it. But after compilation, the JVM essentially just sees a raw `List` of Objects. The specific type info is 'erased' to ensure backward compatibility with older Java versions.

This mainly affects me when parsing JSON. Libraries like Jackson need a `TypeReference` or `Class<T>` passed explicitly because they can't figure out the generic type just by looking at the list at runtime."

---

### 11. What is Java Reflection?
"Reflection allows code to inspect and modify itself at runtime. It's like looking in a mirror—you can see classes, methods, and fields, and invoke them by name rather than directly.

Frameworks like **Spring** and **Hibernate** run on reflection. It’s how they instantiate beans, inject dependencies, and map database columns without us writing boilerplate code.

The downside is that it breaks compile-time safety and can be slower. So while I appreciate it in frameworks, I rarely write reflection code in my own business logic."

---

### 12. Fail-fast vs Fail-safe Iterators?
"It’s about how iterators behave if you modify the collection while looping through it.

**Fail-fast** iterators (like in `ArrayList` or `HashMap`) will throw a `ConcurrentModificationException` immediately if they detect a change. They prioritize crashing over showing you potentially inconsistent data.

**Fail-safe** iterators (like in `ConcurrentHashMap`) work on a snapshot or a View. They allow you to modify the collection while iterating without throwing an exception, though you might not see the latest changes.

If I need to remove items while looping, I usually use the iterator's own `remove()` method or `removeIf()` to avoid the fail-fast exception."

---

### 13. Map() vs flatMap() in Streams?
"Both transform data, but `flatMap` is for when you need to **flatten** nested structures.

**map()** is One-to-One. It takes an element and transforms it. Like taking a `User` object and extracting just the `Username` string.

**flatMap()** is One-to-Many. It expects the function to return a Stream. It’s useful if you have a `List` of `Users`, and each user has a `List` of `PhoneNumbers`, and you want one big list of *all* phone numbers. `flatMap` will take each user's list and 'flatten' them into a single continuous stream.

I also use `flatMap` quite a bit when dealing with `Optional` inside another `Optional` to avoid nested checks."
