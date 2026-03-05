# Design Patterns in Architecture (Service-Based Companies)

## 1. What is the difference between Architectural Patterns and Design Patterns?

**Expected Answer:**
*   **Architectural Patterns:** High-level strategies that define the structural organization of an entire software system. They deal with the overall big picture, components, and how they interact.
    *   *Examples:* Microservices, Monolithic, Layered (N-Tier), Event-Driven, Hexagonal.
*   **Design Patterns:** Lower-level, component-specific solutions to commonly occurring problems in object-oriented software design. They focus on how classes and objects are structured and behave.
    *   *Examples (GoF):* Singleton, Factory, Observer, Strategy, Decorator.

## 2. Explain the Layered (N-Tier) Architectural Pattern. What are its pros and cons?

**Expected Answer:**
The Layered architecture organizes code into horizontal layers, where each layer has a specific role and only depends on the layer directly below it.

*   **Common Layers (3-Tier):**
    1.  **Presentation (UI/Controller):** Handles HTTP requests, user interfaces, and JSON serialization.
    2.  **Business Logic (Service):** Contains the core business rules and computations.
    3.  **Data Access (Repository/DAO):** Communicates directly with the database or external APIs.
*   **Pros:** Easy to understand, easy to organize files, well-supported by traditional frameworks (Spring Boot MVC, Django), good separation of concerns.
*   **Cons:** Can lead to monolithic deployments, changes in the DB layer often bubble up and force changes in the Service and UI layers (tight coupling). Can result in "pass-through" layers where a service method just calls a DAO method doing nothing else.

## 3. Compare MVC (Model-View-Controller) with MVVM (Model-View-ViewModel).

**Expected Answer:**
Both are UI/Presentation architectural patterns, but they handle data flow differently.

*   **MVC (Model-View-Controller):**
    *   *Controller* intercepts user input, updates the *Model*, and selects the *View* to render.
    *   Often tightly coupled. Common in traditional web frameworks (Spring MVC, Ruby on Rails).
*   **MVVM (Model-View-ViewModel):**
    *   *ViewModel* acts as a data binder between the View and Model. It converts Model data into a format easily managed by the View.
    *   Uses **Two-Way Data Binding**: Changes in the UI automatically update the ViewModel, and changes in the ViewModel automatically update the UI.
    *   Common in modern front-end frameworks (Angular) and desktop apps (WPF).

## 4. What is the Singleton Design Pattern, and why is it often considered an anti-pattern in modern architectures?

**Expected Answer:**
*   **Singleton:** Ensures a class has only one instance globally and provides a global point of access to it. (e.g., a DB Connection Pool or File Logger).
*   **Why it's controversial (Anti-pattern):**
    *   **Hidden Dependencies:** Classes use the Singleton globally, making it hard to see what a class actually depends on.
    *   **Testing Difficulties:** Global state persists between unit tests, causing tests to interfere with each other. It's incredibly hard to mock a Singleton.
    *   **Multithreading Issues:** Creating a thread-safe Singleton requires complex locking, which can cause performance bottlenecks.
*   *Modern Alternative:* Use Dependency Injection (DI) frameworks (like Spring or Dagger) to manage a single instance (a "singleton-scoped bean") and inject it where needed, rather than writing structural Singleton code.

## 5. Explain the Observer Pattern and how it relates to modern event-driven systems.

**Expected Answer:**
*   **Observer Pattern:** A behavioral design pattern where an object (the **Subject**) maintains a list of its dependents (the **Observers**) and notifies them automatically of any state changes, usually by calling one of their methods.
*   **How it relates to Event-Driven / Pub-Sub:**
    *   The Observer pattern is the foundational concept behind Publish-Subscribe (Pub-Sub) architectures (like Kafka, RabbitMQ, or frontend EventEmitters).
    *   The Subject acts as the Publisher/Broker. The Observers act as the Subscribers.
    *   It decouples the system: The Subject doesn't need to know the specific details or classes of the Observers, it just fires the event.

## 6. What is the Factory Pattern and when would you use it at an architectural level?

**Expected Answer:**
*   **Factory Pattern:** A creational pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. It abstracts away the `new` keyword.
*   **Architectural Use Case:**
    *   When implementing a plugin architecture or interacting with external services.
    *   *Example:* A `PaymentProcessorFactory`. At runtime, based on user input or configuration, the factory decides whether to instantiate and return a `StripeService`, a `PayPalService`, or an `RazorpayService`. The calling code doesn't care which one it gets, as long as it implements the `PaymentService` interface.
