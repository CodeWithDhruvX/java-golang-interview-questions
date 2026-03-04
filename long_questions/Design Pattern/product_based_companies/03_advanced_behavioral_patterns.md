# Advanced Behavioral Patterns - Product Based Companies

## 1. How does modern Java (Lambdas/Functional Interfaces) impact Behavioral Patterns like Strategy and Command?

**Answer:**
Java 8+ features (Lambdas, Method References, Functional Interfaces) significantly simplify the implementation of several behavioral patterns, reducing boilerplate code.

**Strategy Pattern:**
Traditionally, you define a `Strategy` interface and create multiple concrete classes implementing it.
*   **The Modern Way:** Instead of creating physical classes, you can pass lambdas that match the functional interface signature.
```java
// Traditional
List<File> files = getFiles();
files.sort(new FileSizeComparator()); // Requires a physical class

// Modern (Strategy via Lambda/Method Reference)
files.sort(Comparator.comparingLong(File::length));
```
The strategy is defined inline, making the code much more concise and functional.

**Command Pattern:**
Traditionally, a Command interface requires concrete classes that implement `execute()`, which hold state and a receiver.
*   **The Modern Way:** If the command doesn't need to be serialized or support undo functionality, an action can just be a `Runnable` or `Consumer` passed directly as a lambda.

---

## 2. Observer Pattern vs Pub-Sub Architecture. Are they the same? What are the key differences for distributed systems?

**Answer:**
While they solve a similar problem (event notification), they are structurally different, which is crucial for scalable systems.

| Feature | Observer Pattern | Publisher-Subscriber (Pub-Sub) |
| :--- | :--- | :--- |
| **Coupling** | **Tight Coupling:** The Subject knows its Observers (holds a list of references). They usually exist in the same application space/JVM. | **Loose Coupling:** Publishers don't know who subscribers are. They communicate via a *Message Broker* or *Event Bus* (e.g., Kafka, RabbitMQ). |
| **Synchronicity** | Usually **Synchronous**: The subject iterates through its list and calls `update()` on each observer. The subject blocks until all observers finish handling the event. | Almost always **Asynchronous**: Publishers fire messages to the broker and return immediately. Subscribers process at their own pace. |
| **Scalability** | Struggles to scale across multiple machines. | Highly scalable across distributed microservices. |
| **Topology** | Point-to-Point (sort of a star topology centered on the subject).| Many-to-Many via topics/channels in the broker. |

**Product-based context:** You will rarely implement a raw Observer pattern in a modern distributed backend. You will use a Pub-Sub architecture (Event-Driven Architecture) via tools like Apache Kafka or AWS SNS/SQS.

---

## 3. How do you implement the Chain of Responsibility pattern in a highly concurrent environment?

**Answer:**
In a multi-threaded service (like an HTTP server handling thousands of requests), standard Chain of Responsibility (CoR) implementations can face issues if the handlers maintain state.

**Considerations for Concurrency:**
1.  **Stateless Handlers:** The safest approach. Handlers should not store any request-specific data in their instance variables. All state should be passed along inside the `Context` object (or the `Request` object) traveling through the chain.
    *   *Spring Security Filter Chain* is an excellent example of a thread-safe, stateless CoR.
2.  **ThreadLocal:** If a handler absolute needs to maintain state for the duration of a request without explicitly passing it down the signature, you could use `ThreadLocal`. (Though explicit context passing is usually preferred for clarity).
3.  **Immutable Contexts:** To prevent race conditions if the chain itself forks into asynchronous processes, the context object passed between handlers should be heavily immutable or carefully synchronized.
4.  **Asynchronous CoR:** Modern systems might implement CoR using Reactive programming (RxJava, Project Reactor) or Promises/CompletableFutures, where each step in the chain is executed asynchronously and chained via `.thenApply()` or `.flatMap()`.

---

## 4. Give a real-world system design example using the State Pattern.

**Answer:**
The State pattern is excellent for managing complex workflows with strictly defined transitions.

**System Design Example: E-commerce Order Processing System**
An `Order` goes through specific states: `NEW`, `PAYMENT_PENDING`, `PROCESSING`, `SHIPPED`, `DELIVERED`, `CANCELLED`.

Instead of massive `if...else` blocks in a single `OrderService`:
```java
// Anti-pattern
public void cancel(Order order) {
    if (order.getState() == SHIPPED) throw new Exception("Too late");
    else if (order.getState() == DELIVERED) throw new Exception("Already done");
    // ... logic
}
```

**State Pattern Implementation:**
1.  Create an `OrderState` interface: `pay()`, `ship()`, `cancel()`, `deliver()`.
2.  Create Concrete States: `NewOrderState`, `ShippedOrderState`, etc.
3.  The `Order` class holds a reference to an `OrderState`.

*   When `order.cancel()` is called during `ProcessingState`, the logic handles refunds and moves the state to `CancelledState`.
*   When `order.cancel()` is called during `ShippedState`, the State class itself throws an `IllegalStateException` simply and cleanly.

This centralizes logic per state, entirely eliminating complex nested conditionals in your core business logic, adhering strictly to the Single Responsibility Principle.
