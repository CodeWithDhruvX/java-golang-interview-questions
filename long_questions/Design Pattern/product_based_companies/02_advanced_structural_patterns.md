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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"The Proxy pattern is the foundation of Spring's magic, especially for AOP and transaction management. When you add annotations like `@Transactional` or `@Cacheable` to your methods, Spring doesn't actually give you the direct object - it wraps your bean in a proxy that intercepts method calls. This proxy handles all the cross-cutting concerns like opening transactions, caching results, or logging before passing the call to your actual method."

"What's really clever is how Spring chooses between JDK dynamic proxies and CGLIB proxies. If your class implements an interface, Spring uses Java's built-in reflection to create a proxy that implements the same interface. But if your class doesn't implement an interface, Spring uses CGLIB to dynamically generate a subclass at runtime. This is why Spring recommends programming to interfaces - it's more efficient and cleaner."

"The gotcha that trips up even experienced developers is the self-invocation problem. If you have a method in your service that calls another method in the same class, and that second method has `@Transactional`, it won't work! Why? Because the internal call bypasses the proxy completely - you're calling the method directly on the actual object, not the proxy wrapper. I've seen this cause subtle bugs in production where transactions weren't being committed. The solution is either to extract the method to a separate bean or use AspectJ compile-time weaving for more advanced scenarios."

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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"The Flyweight pattern is all about memory optimization when you have massive numbers of similar objects. The key insight is to separate the object's state into two parts - intrinsic state that's shared among all objects, and extrinsic state that's unique to each instance. Think about a multiplayer game with millions of trees - each tree has the same 3D model and textures (intrinsic), but different positions and orientations (extrinsic)."

"Instead of storing the heavy 3D model data for every single tree, you create one shared TreeModel object with all the expensive mesh and texture data. Then each individual tree instance just stores its position, scale, and a reference to the shared model. This reduces memory usage from gigabytes to megabytes. The rendering engine calls the shared model's draw method, passing in each tree's specific position and orientation."

"I worked on a virtual reality application at Meta where we had to render millions of particles in a particle system. Each particle had the same shader and texture but different positions, velocities, and lifetimes. Using Flyweight, we created one shared ParticleModel with the heavy GPU resources, and millions of lightweight Particle instances with just the transform data. This allowed us to render 10 million particles on a single GPU that would have otherwise crashed. The pattern is also heavily used in text rendering - instead of storing glyph data for every character, you share the glyph data and just store position and style for each character instance."

---

## 3. Compare Decorator vs Adapter vs Facade. In what architectural layers do you typically place them?

**Answer:**

| Pattern | Intent | Core Difference | Typical Architectural Layer |
| :--- | :--- | :--- | :--- |
| **Adapter** | Converts one interface to another. | Makes incompatible things work together. It wraps an existing object and provides a *different* interface. | **Integration / Infrastructure Layer:** Used when integrating 3rd party APIs, calling legacy systems, or mapping external DTOs to internal domains. |
| **Decorator** | Adds responsibilities dynamically without subclassing. | Enhances an object. It wraps an existing object and provides the *same* interface. | **Business Logic / Middleware Layer:** Used for adding chains of responsibility, like adding caching, logging, or metric tracking to a service layer without modifying the core service. |
| **Facade** | Provides a simplified interface to a complex subsystem. | Simplifies interaction. It creates a *new, higher-level* interface that sits on top of many classes. | **API Gateway / Controller Layer:** Used to provide a simple entry point to internal microservices, or wrapping a complex internal SDK to present a clean service interface to the UI. |

#### 💬 **How to Explain in Interviews (Spoken Format)**

"Adapter, Decorator, and Facade are all wrapper patterns, but they solve very different problems and live in different architectural layers. The Adapter is the diplomat - it makes incompatible interfaces work together. Think about integrating a third-party payment API that uses XML with your system that expects JSON. The Adapter translates between them. It lives in the integration layer, handling all the messy translation logic."

"The Decorator is the enhancer - it adds new capabilities to an object without changing the object itself. It provides the same interface but adds behavior. Think about adding caching, logging, or metrics to a service method. You wrap the service with a caching decorator, then wrap that with a logging decorator. Each decorator adds a specific concern. This lives in the business logic or middleware layer where you're enhancing core services."

"The Facade is the simplifier - it provides a clean, high-level interface to a complex subsystem. Think about an API gateway that hides the complexity of multiple microservices behind a single endpoint. Or a controller that orchestrates multiple service calls into one simple method for the UI. The Facade lives at the top layer - the API gateway or controller layer - providing simplicity to the outside world while hiding internal complexity."

"In a microservices architecture I designed at Netflix, we used all three patterns. Adapter for integrating with legacy systems, Decorator for adding cross-cutting concerns like caching and monitoring, and Facade for our API gateway that presented a unified interface to hundreds of internal microservices. Each pattern had its specific place and purpose in the architecture."

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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"An API Gateway is essentially a pipeline of middleware that processes every request - authentication, rate limiting, logging, routing, and so on. The Decorator pattern is perfect for this because each middleware becomes a decorator that wraps the next one in the chain. You start with the core router that actually forwards requests to microservices, then you decorate it with authentication, then decorate that with logging, then with rate limiting."

"What's beautiful about this approach is that it's completely composable. At startup, you read the configuration and dynamically assemble the decorators based on what each route needs. Some routes might need authentication and rate limiting, others might need just logging. You can even change the middleware order at runtime without touching the core routing logic. This is much more flexible than using a simple chain of responsibility where the order is hardcoded."

"I implemented this pattern for an API gateway at Uber that handled thousands of requests per second. We had decorators for authentication (JWT validation), rate limiting (Redis-based), logging (structured logs to ELK), metrics (Prometheus), and even A/B testing. The core router just knew how to forward requests, but the decorators handled all the cross-cutting concerns. When we needed to add a new security check, we just created a new decorator and updated the configuration - no changes to existing code. This made the system incredibly maintainable and extensible."

"The key insight is that decorators and chain of responsibility can solve the same problem, but decorators give you better structure and type safety. Each decorator implements the same interface, so the compiler ensures everything fits together properly. This is crucial in production systems where you need compile-time guarantees that your middleware pipeline is correctly assembled."
