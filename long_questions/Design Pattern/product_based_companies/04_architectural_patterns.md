# Architectural and System Design Patterns - Product Based Companies

Product-based companies frequently ask about architectural patterns, as they bridge the gap between codebase-level design patterns and high-level system design.

## 1. What is the CQRS (Command Query Responsibility Segregation) pattern? Why use it?

**Answer:**
**CQRS** is an architectural pattern that dictates separating the data mutation operations (Commands) from the data read operations (Queries) into distinct, independent interfaces and often different logical or physical models.

*   **Command Model:** Highly transactional, focuses on validation, strong consistency, and domain logic. Optimized for fast writes. (e.g., saving to a normalized PostgreSQL DB).
*   **Query Model:** Denormalized, heavily cached, optimized for fast reads. Does not contain complex domain business logic. (e.g., reading from an Elasticsearch index or Redis cache).

**Why use it?**
1.  **Asymmetric Scaling:** In most systems, reads heavily outnumber writes (100:1 ratio). CQRS allows you to scale the read infrastructure independently of the write infrastructure.
2.  **Performance Optimization:** The read database can be denormalized (pre-joined data) specifically matching the precise UI view requirements, leading to sub-millisecond query times.
3.  **Complexity Separation:** Keeps the complex domain logic (Write side) clean and disentangled from complex reporting/search queries (Read side).

**Trade-offs:** It introduces significant architectural complexity and forces you to deal with **Eventual Consistency** (there is a delay between a write succeeding and the read model being updated).

---

## 2. Explain the Saga Pattern. What problem does it solve in Microservices?

**Answer:**
The **Saga Pattern** is a failure management pattern that helps establish distributed transactions across multiple microservices without using traditional two-phase commit (2PC) locks, which scale poorly.

**The Problem:**
In a monolithic application, you process a large workflow (e.g., E-commerce Order creation that involves Inventory, Payment, and Shipping) inside a single ACID database transaction.
In Microservices, Inventory, Payment, and Shipping have their own independent databases. A single ACID transaction cannot span them robustly.

**The Saga Solution:**
A Saga is a sequence of local transactions. Each local transaction updates its own database and publishes an event or message to trigger the next local transaction in the saga.

If a local transaction fails (e.g., Inventory confirms, but Payment fails), the Saga executes a series of **Compensating Transactions** that undo the changes made by the preceding local transactions (e.g., sending an event to Inventory to release the held stock).

**Implementation Styles:**
*   **Choreography:** Decentralized. Services listen to events published by other services and react accordingly. No central orchestrator exists. Good for simple workflows.
*   **Orchestration:** Centralized. A dedicated "Saga Orchestrator" service essentially acts as a state machine. It explicitly commands other services what to do (e.g., "Send Payment Command") and tracks the overall state of the workflow to decide to continue or compensate. Good for complex workflows.

---

## 3. What is the Circuit Breaker pattern? How does it increase system resiliency?

**Answer:**
The **Circuit Breaker** pattern protects an application from repeatedly calling a down network service, database, or API, preventing cascading failures and thread pool exhaustion in distributed systems.

**Mechanics (State Machine):**
*   **CLOSED state:** Everything is normal. Requests flow through to the external service. The circuit breaker counts failures.
*   **OPEN state:** If the failure threshold (e.g., 50% errors in the last 10 seconds) is exceeded, the circuit trips. All subsequent calls immediately fail fast (throwing an exception or returning a fallback value) *without* attempting to hit the failing service. This relieves pressure on the struggling service and prevents the caller's threads from hanging.
*   **HALF-OPEN state:** After a configured timeout period, the circuit allows a limited number of "test" requests through.
    *   If they succeed, the circuit resets completely to **CLOSED**.
    *   If they fail, the circuit breaks again and returns to **OPEN**.

**Key benefit:** Fail Fast. Instead of hanging for 30 seconds waiting for a timeout connection that will eventually fail anyway (tying up valuable tomcat/worker threads), the system receives an immediate error and gracefully handles it (perhaps showing a degraded UI to the user). Libraries like Netflix Hystrix or Resilience4j implement this.

---

## 4. Can you describe the MVC, MVP, and MVVM architectural patterns?

**Answer:**
These are presentation-layer patterns used to separate concerns between the user interface and the underlying business logic.

*   **MVC (Model-View-Controller):**
    *   **Model:** Represents the data and business rules.
    *   **View:** The UI components.
    *   **Controller:** Receives input, manipulates the Model, and determines which View to display. The View often observes the Model for changes. (Commonly used in traditional web frameworks like Spring MVC, Ruby on Rails).
*   **MVP (Model-View-Presenter):**
    *   **Model:** Data and business rules.
    *   **View:** Passive interface. Highly decoupled from the Model. Forwards inputs to Presenter.
    *   **Presenter:** Acts as a middleman. Receives user actions from the View, reads/updates the Model, and explicitly updates UI elements in the View. (Commonly used in Android development before MVVM).
*   **MVVM (Model-View-ViewModel):**
    *   **Model:** Data and business rules.
    *   **View:** Defines the structure, layout, and appearance.
    *   **ViewModel:** An abstraction of the View. Exposes state and commands. The key difference is **Data Binding**. The View binds itself directly to properties on the ViewModel. When ViewModel property changes, the View updates automatically without the ViewModel having explicit references to UI elements. (Standard in modern Android SDK, Angular, Vue, React/Redux-ish architectures).
