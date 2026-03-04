# Advanced Structural Patterns - Product Based Companies

## 1. How is the Proxy design pattern used in Spring Framework (specifically AOP and Transaction Management)?

**Answer:**
The **Proxy pattern** is the foundational building block for Spring's Aspect-Oriented Programming (AOP).

When you annotate a method with `@Transactional`, `@Cacheable`, or your custom AOP annotations, Spring does not give the client a direct reference to the target object. Instead, it dynamically creates a **Proxy** object that wraps the target bean.

**How it works (The Interceptor mechanism):**
1.  **JDK Dynamic Proxies:** If the target class implements an interface, Spring uses Java's built-in reflection to create a dynamic proxy.
2.  **CGLIB Proxies:** If the target class does not implement an interface, Spring uses the CGLIB library to dynamically generate a subclass of the target class at runtime to act as the proxy.

**Execution Flow (e.g., `@Transactional`):**
1.  The client calls `userService.createUser()`. They are actually calling a method on the Proxy.
2.  The Proxy "intercepts" the call. It asks the `TransactionManager` to open a database transaction.
3.  The Proxy forwards the call to the *actual* target object's `createUser()` method.
4.  If the target method completes successfully, the Proxy intercepts the return and commits the transaction.
5.  If an exception is thrown, the Proxy catches it, rolls back the transaction, and re-throws the exception to the client.

**Interview Gotcha:** The "Self-Invocation" problem. If a method inside the target class calls *another* method inside the same class that has `@Transactional`, the transaction aspect will *not* work. This is because the internal call bypasses the Spring Proxy completely!

---

## 2. Explain the Flyweight design pattern. How would you use it to optimize a massive multi-player online game (MMOG)?

**Answer:**
The **Flyweight pattern** is a structural pattern used to minimize memory usage or computational expenses by sharing as much as possible with similar objects. It is crucial when dealing with an enormous number of objects that would crash the system if instantiated individually.

**Key Concept: Intrinsic vs Extrinsic State**
*   **Intrinsic State:** Data that is context-independent and identical across many objects (e.g., the 3D mesh model, base textures, max health stats of a character class). This state is *immutable* and stored *inside* the Flyweight object.
*   **Extrinsic State:** Data that depends on the specific context and changes at runtime (e.g., the specific X/Y coordinate, current health, current equipped weapon). This state is stored *outside* the object and passed to the Flyweight's methods when needed.

**MMOG Use Case (Rendering a Forest):**
Imagine a game with 1 million trees. Each tree consists of a complex 3D mesh and textures (requiring 5MB of RAM). 1 million trees * 5MB = 5 Terabytes of RAM!

**Flyweight Solution:**
1.  **The Flyweight (`TreeModel`):** Create a single `TreeModel` object containing the heavy intrinsic data (the 5MB mesh/textures).
2.  **The Context (`Tree`):** Create a lightweight `Tree` class that only stores the extrinsic state: `x`, `y`, `z` coordinates, `heightScale`, `colorTint`, and a *reference* to the shared `TreeModel`. (This class takes maybe 32 bytes).

When rendering, the game iterates over the millions of lightweight `Tree` contexts and passes their specific coordinates to the shared `TreeModel.draw(x, y, scale)` method.

---

## 3. Compare Decorator vs Adapter vs Facade. In what architectural layers do you typically place them?

**Answer:**

| Pattern | Intent | Core Difference | Typical Architectural Layer |
| :--- | :--- | :--- | :--- |
| **Adapter** | Converts one interface to another. | Makes incompatible things work together. It wraps an existing object and provides a *different* interface. | **Integration / Infrastructure Layer:** Used when integrating 3rd party APIs, calling legacy systems, or mapping external DTOs to internal domains. |
| **Decorator** | Adds responsibilities dynamically without subclassing. | Enhances an object. It wraps an existing object and provides the *same* interface. | **Business Logic / Middleware Layer:** Used for adding chains of responsibility, like adding caching, logging, or metric tracking to a service layer without modifying the core service. |
| **Facade** | Provides a simplified interface to a complex subsystem. | Simplifies interaction. It creates a *new, higher-level* interface that sits on top of many classes. | **API Gateway / Controller Layer:** Used to provide a simple entry point to internal microservices, or wrapping a complex internal SDK to present a clean service interface to the UI. |

---

## 4. How would you design an API Gateway middleware chain using structural patterns?

**Answer:**
An API Gateway essentially processes an incoming HTTP request through a series of filters/middlewares (Authentication, Rate Limiting, Logging, Routing).

**Pattern Choice: Chain of Responsibility (Behavioral) OR Decorator (Structural)**

Using the **Decorator Pattern** for Middleware:
You define a core interface `RequestHandler`.
```java
public interface RequestHandler {
    Response handle(Request request);
}
```

The core implementation is actual routing to the microservice:
```java
public class CoreRouter implements RequestHandler {
    public Response handle(Request request) { ...route to upstream... }
}
```

Then you create structural Decorators:
```java
public class AuthDecorator implements RequestHandler {
    private RequestHandler nextInChain;
    public AuthDecorator(RequestHandler next) { this.nextInChain = next; }

    public Response handle(Request request) {
        if (!validateToken(request)) throw new UnauthorizedException();
        return nextInChain.handle(request);
    }
}

public class LoggingDecorator implements RequestHandler { ... }
```

At gateway startup, you dynamically weave the decorators together based on route configuration:
```java
// Assembly
RequestHandler pipeline = new LoggingDecorator(
                            new AuthDecorator(
                                new CoreRouter()));
// Execution
pipeline.handle(incomingRequest);
```
This structural composition allows infinite flexibility in adding or removing middleware layers at runtime without altering the core routing logic.
