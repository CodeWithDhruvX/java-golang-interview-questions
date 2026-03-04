# Advanced Creational Patterns - Product Based Companies

## 1. How do you implement a truly indestructible Singleton in Java? What vulnerabilities exist in Double-Checked Locking?

**Answer:**
While Double-Checked Locking (with `volatile`) is thread-safe, the Singleton pattern can still be broken using three mechanisms:
1.  **Reflection:** An attacker can use `setAccessible(true)` on the private constructor and instantiate a new object.
2.  **Serialization/Deserialization:** Deserializing a Singleton class creates a new instance.
3.  **Cloning:** If the Singleton class implements `Cloneable`, calling `clone()` creates a new instance.

**The "Indestructible" Solution (Enum Singleton):**
The most robust way to implement a Singleton in Java, as recommended by Joshua Bloch (Effective Java), is using an `enum`.

```java
public enum SingletonEnum {
    INSTANCE;

    // Add state and behavior
    private int value;

    public void setValue(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }
}
```

**Why Enum is indestructible:**
*   **Reflection Proof:** The Java Language Specification strongly prevents instantiating an enum via reflection. If you try, it throws an `IllegalArgumentException` ("Cannot reflectively create enum objects").
*   **Serialization Proof:** Java guarantees that enums are serialized correctly. When deserialized, it returns the same `INSTANCE` without creating a new object. You don't need to implement `readResolve()`.
*   **Thread-Safe by Default:** The JVM guarantees that enum instances are created safely in a multi-threaded environment.

**Fixing the standard Singleton (If Enum is not used):**
*   **Reflection:** Throw an exception in the constructor if `instance != null`.
*   **Serialization:** Implement the `readResolve()` method to return the existing instance.
*   **Cloning:** Override the `clone()` method to throw a `CloneNotSupportedException`.

---

## 2. In what real-world framework scenarios is the Abstract Factory pattern heavily utilized?

**Answer:**
The Abstract Factory pattern is often used in building Cross-Platform UI frameworks and Database Access layers.

**Scenario 1: Cross-Platform UI Toolkits (e.g., Java AWT/Swing or React Native)**
Suppose you are building an application that needs to render native UI components on both Windows and MacOS.
*   **Abstract Factory (`GUIFactory`)**: Declares `createButton()`, `createCheckbox()`.
*   **Concrete Factories**: `WindowsFactory`, `MacFactory`.
*   **Abstract Products**: `Button`, `Checkbox`.
*   **Concrete Products**: `WindowsButton`, `MacButton`.

At runtime, the application detects the OS, instantiates the corresponding concrete factory (e.g., `new MacFactory()`), and passes it to the UI rendering engine. The engine creates buttons and checkboxes without knowing their exact classes.

**Scenario 2: Database Abstraction Layers (JDBC/ORM)**
When an application needs to support multiple databases (MySQL, PostgreSQL, Oracle).
*   **Abstract Factory (`DbConnectionFactory`)**: Declares `createConnection()`, `createCommand()`.
*   **Concrete Factories**: `MySqlFactory`, `PostgresFactory`.
This allows the application to switch databases easily by providing a different factory implementation at startup or via dependency injection.

---

## 3. How does the Prototype design pattern connect with the concept of strict Immutability and standard object copying?

**Answer:**
The **Prototype Pattern** is used to create a new object by cloning an existing one, avoiding the cost of creating a new object from scratch (especially if database calls or network requests are involved).

**Connection to Object Copying (Shallow vs Deep Copying):**
When implementing the Prototype pattern, a critical decision is whether to use a Shallow Copy or a Deep Copy.
*   **Shallow Copy (e.g., `Object.clone()` in Java default behavior):** Copies only primitive fields. References to other objects are copied as references. Modifying a mutable nested object in the clone will affect the original.
*   **Deep Copy:** Creates entirely new copies of all nested objects recursively. Modifying the clone has completely no effect on the original object.

**Connection to Immutability:**
*   If your object state consists *entirely* of primitive types or immutable objects (like `String`, `Integer`, or custom immutable types), a Shallow Copy is perfectly safe and highly efficient.
*   Therefore, designing your classes to be heavily immutable drastically simplifies the implementation of the Prototype pattern and concurrency concerns. If objects are immutable, you don't even *need* the Prototype pattern to clone them; you can just share the reference safely! The Prototype pattern is most useful specifically for complex, **mutable** objects whose construction is expensive.

---

## 4. How would you design a rate limiter using creational patterns?

**Answer:**
Designing a Rate Limiter is a classic system design question. Let's look at it through the lens of creational patterns:

**1. Singleton/Registry for the Rate Limiter Manager:**
You need a central point to access rate limiters for different users or APIs. A `RateLimiterRegistry` can be a **Singleton** (or an injected Singleton Scope bean).
```java
public class RateLimiterRegistry {
    private static RateLimiterRegistry instance = new RateLimiterRegistry();
    private ConcurrentHashMap<String, RateLimiter> limiters = new ConcurrentHashMap<>();

    public static RateLimiterRegistry getInstance() { return instance; }

    public RateLimiter getRateLimiter(String apiKey) {
        // Factory Method invocation inside
        return limiters.computeIfAbsent(apiKey, k -> RateLimiterFactory.createLimiter(k));
    }
}
```

**2. Factory Method / Abstract Factory for Strategy Selection:**
There are multiple rate-limiting algorithms (Token Bucket, Leaky Bucket, Sliding Window Log). A **Factory Method** can determine which instantiation strategy to use based on configuration.
```java
public class RateLimiterFactory {
    public static RateLimiter createLimiter(String key) {
        String algo = Config.getAlgorithmFor(key); // e.g., "TOKEN_BUCKET"
        switch(algo) {
            case "TOKEN_BUCKET": return new TokenBucketLimiter(capacity, refillRate);
            case "SLIDING_WINDOW": return new SlidingWindowLimiter(windowSize);
            default: throw new IllegalArgumentException();
        }
    }
}
```

**3. Builder Pattern for Configuration:**
A Rate Limiter instance might have multiple configuration parameters (e.g., permits per second, max burst size, warmup period). The **Builder pattern** helps create these objects cleanly.
```java
RateLimiter tokenLimiter = RateLimiterBuilder.create()
                                .withCapacity(100)
                                .withRefillTokensPerSecond(10)
                                .build();
```
