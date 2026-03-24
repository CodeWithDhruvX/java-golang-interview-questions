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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"CQRS is all about recognizing that reads and writes have very different requirements in most systems. Reads typically outnumber writes by 100 to 1, and they need different optimizations. The write side needs strong consistency, validation, and business logic enforcement - usually in a normalized relational database. The read side needs to be fast, denormalized, and optimized for specific UI queries - often in Elasticsearch or Redis."

"The beauty of CQRS is that you can scale each side independently. If your read traffic is high, you can add more read replicas or cache layers without touching the write infrastructure. The read model can be pre-joined and structured exactly how your UI needs it, leading to sub-millisecond query times. I implemented CQRS in an e-commerce system at Amazon where product searches needed to be incredibly fast but product updates needed strong consistency. The write side used PostgreSQL for ACID compliance, while the read side used Elasticsearch for lightning-fast search and filtering."

"The big trade-off is eventual consistency. When you update a product, there's a small delay before the search index reflects the change. This is acceptable for most applications but not for all. The architecture is also more complex - you need to handle event publishing, model synchronization, and consistency checking. But for high-scale systems where read performance is critical, CQRS is often the best solution. The key is knowing when the benefits outweigh the complexity."

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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"The Saga pattern solves the distributed transaction problem in microservices. In a monolith, you can wrap everything in a single ACID transaction, but in microservices, each service has its own database. When you need to coordinate across services - like an e-commerce order that needs to update inventory, process payment, and schedule shipping - you can't use traditional two-phase commit because it doesn't scale well."

"A Saga breaks the workflow into a series of local transactions, each in its own service. If one step fails, the pattern executes compensating transactions to undo the previous steps. For example, if payment fails after inventory was reserved, the Saga triggers a compensating transaction to release the inventory. This maintains data consistency without distributed locks."

"There are two implementation styles. Choreography is decentralized - services listen to events and react independently. It's simpler but harder to track complex workflows. Orchestration uses a central coordinator that explicitly commands each service. It's more complex but gives you better visibility and control. I used orchestration in a banking system at JP Morgan because we needed strict compliance and audit trails - the orchestrator gave us a clear view of every transaction's state. The choice depends on your workflow complexity and operational requirements."

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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"The Circuit Breaker pattern protects your system from cascading failures when external services go down. Imagine you're calling a payment service that's struggling - if you keep sending requests, they'll all hang and eventually exhaust your thread pool, bringing down your entire application. The Circuit Breaker monitors failures and 'trips' when the failure rate exceeds a threshold."

"When the circuit is open, all calls to that service fail immediately - no waiting, no hanging threads. After a timeout, it enters half-open state and allows a few test requests. If those succeed, it closes again and normal operation resumes. This fail-fast approach is crucial for system resilience."

"I implemented Circuit Breakers throughout a microservices architecture at Netflix. We had them around database connections, external APIs, and even internal service calls. When the recommendation service went down, the circuit breaker opened and our main application continued serving content with default recommendations instead of crashing. The pattern buys you time to fix the underlying issue while keeping the user experience functional."

"The key insight is that Circuit Breakers aren't just about handling failures - they're about maintaining system availability during partial outages. In distributed systems, failures are inevitable, but cascading failures are preventable. Libraries like Hystrix and Resilience4j make this easy to implement, and they're essential for any production microservices architecture."

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

#### 💬 **How to Explain in Interviews (Spoken Format)**

"MVC, MVP, and MVVM are all about separating UI from business logic, but they do it in different ways. MVC is the classic pattern - the Controller receives input, manipulates the Model, and tells the View what to display. The View often observes the Model directly for changes. This works well for traditional web applications where the server renders HTML."

"MVP evolved to solve some MVC problems. The Presenter acts as a complete middleman - the View is passive and just forwards user input to the Presenter. The Presenter updates the View directly. This gives you better testability because the Presenter can be tested independently of the UI. It was popular in Android development before MVVM came along."

"MVVM is the modern approach, especially for mobile and SPA frameworks. The key innovation is data binding - the View automatically updates when ViewModel properties change, and vice versa. The ViewModel has no reference to UI elements, making it highly testable. This is the standard in modern Android, Angular, and React applications."

"In my experience, the choice depends on your platform and requirements. For server-side web apps, MVC still works great. For Android or complex UI applications, MVVM is usually the best choice because of data binding and testability. I've used all three patterns, and MVVM generally provides the cleanest separation of concerns for modern applications. The key is understanding that they all solve the same fundamental problem - keeping business logic out of your UI - but with different levels of sophistication and platform support."
