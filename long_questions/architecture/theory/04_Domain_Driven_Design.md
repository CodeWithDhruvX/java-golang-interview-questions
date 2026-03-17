# 🏛️ Domain-Driven Design (DDD) — Questions 1–10

> **Level:** 🔴 Senior
> **Asked at:** Amazon, Google, Uber, Flipkart, Razorpay — senior engineering and principal architect roles

---

### 1. What is Domain-Driven Design (DDD)?

"Domain-Driven Design is a software design philosophy that says: **the structure of your code should reflect the structure of your business domain**. The primary tool for doing this is deep collaboration between developers and domain experts (business stakeholders) to build a **shared model** of the domain.

The core problem DDD solves: technical code that doesn't match business language. When a developer says `UserEntitlementRecord` and the business says `Subscription`, and these map to different things — that's where bugs breed.

DDD gives us a rich vocabulary: **entities** (objects with identity), **value objects** (immutable descriptors), **aggregates** (consistency boundaries), **domain events**, and **repositories**. When I apply DDD, my code reads like the business speaks."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (Principal SDE), Google, Razorpay, CRED, Zepto

#### Indepth
DDD is divided into:
- **Strategic DDD:** How to carve up the system at the macro level — Bounded Contexts, Context Maps, Subdomains. This maps directly to microservice boundaries.
- **Tactical DDD:** Patterns for implementing the domain model — Entities, Value Objects, Aggregates, Domain Events, Repositories, Services. This is code-level design.

The **ubiquitous language** is the cornerstone of DDD: a shared vocabulary used by both technical and business teams in conversations, documentation, code, and tests. If the business says "policy", `Policy` is a class in your code — not `PolicyWrapper`, not `PolicyDto`, not `InsuranceRecord`.

Books: *Domain-Driven Design* by Eric Evans (the blue book), *Implementing Domain-Driven Design* by Vaughn Vernon.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is Domain-Driven Design (DDD)?
**Your Response:** "Domain-Driven Design (DDD) is a software development philosophy that places the **business domain at the absolute center of design**. The goal is to ensure that the code structure exactly reflects the business language and logic. 

The most important tool in DDD is the **Ubiquitous Language**—a shared vocabulary used by both developers and domain experts. When we use the same terms in our conversations, our requirements, and our code, we drastically reduce the chance of 'translation errors' that lead to bugs. DDD is generally divided into **Strategic design**, which helps us carve out microservice boundaries using Bounded Contexts, and **Tactical design**, which provides patterns like Entities and Value Objects to model complex logic within those boundaries."

---

### 2. What is a Bounded Context?

"A Bounded Context is a **clear boundary within which a domain model is defined and applicable**. The same concept (like 'Customer') can mean different things in different contexts — and that's OK.

In an e-commerce system: In the **Order context**, a Customer has a name, address, and order history. In the **Support context**, a Customer has ticket history and agent notes. In the **Marketing context**, a Customer has segments, preferences, and campaign history. These are three different models of 'Customer', and trying to make one unified model that serves all three is a nightmare.

Bounded Contexts are the boundaries that separate these models. Each context has its own codebase, its own database, its own team. In microservices, each Bounded Context is typically a service (or a small cluster of closely related services)."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Principal architect roles, senior design rounds

#### Indepth
Context Map patterns (how contexts relate to each other):
1. **Shared Kernel:** Two teams share a small model subset. High coupling, requires coordination.
2. **Customer-Supplier:** Upstream team provides APIs, downstream team consumes. Downstream doesn't control changes.
3. **Conformist:** Downstream simply adopts upstream's model as-is (no translation).
4. **Anti-Corruption Layer (ACL):** Downstream translates between the upstream model and its own domain model. Prevents pollution from a legacy or external system.
5. **Open Host Service / Published Language:** Upstream publishes a formal, versioned API (like OpenAPI or protobuf) that anyone can consume.

**Practical application:** When designing microservices, the Bounded Context analysis is the most valuable tool. Don't decompose by technical layer (DB service, API service) — decompose by business context (Checkout, Catalog, Fulfillment, Customer).

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a Bounded Context?
**Your Response:** "A Bounded Context is a clear linguistic and consistency boundary within which a specific domain model is defined and valid. The key insight of DDD is that it's impossible—and often harmful—to try and create one single 'God Model' for an entire multi-billion dollar company. 

For example, a 'User' means something very different to the **Identity Context** (passwords, MFA) than it does to the **Sales Context** (discounts, order history). By creating separate Bounded Contexts, we allow these models to evolve independently without stepping on each other's toes. In a modern architecture, a Bounded Context is the most natural boundary for a **microservice**—it ensures that each service is cohesive and decoupled from the internal details of other domains."

---

### 3. What is an Entity in DDD?

"An Entity is a domain object defined by its **continuous identity** — it has a unique identifier that persists over time and across state changes.

A `User` is an Entity. If the user changes their email, name, or address, they're still the *same* user — tracked by their `userId`. Two users with the same name are different entities (different IDs).

Compare to a Value Object: a `MoneyAmount(500, INR)` is defined by its value, not an ID. Two `MoneyAmount(500, INR)` objects are identical and interchangeable.

In code: Entities have equality based on ID. Value Objects have equality based on value. This distinction matters for domain logic — don't store entity references where you should store value objects."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Companies doing DDD implementation

#### Indepth
Entity characteristics:
- **Has an identity (ID):** `userId`, `orderId`, `productId`
- **Mutable:** State changes over time — order goes from PENDING → PAID → SHIPPED
- **Tracked over time:** The lifecycle of an entity is important
- **Equality by ID:** `user1.equals(user2)` is `user1.id == user2.id`, ignoring other fields

Good Entity design in Go:
```go
type Order struct {
    id     OrderID     // identity — never changes
    status OrderStatus // mutable state
    items  []OrderItem // associated value objects
}

func (o *Order) AddItem(item OrderItem) error {
    if o.status != Pending { return ErrOrderNotEditable }
    o.items = append(o.items, item)
    return nil
}
```

**Anemic domain model anti-pattern:** When Entity classes are pure data containers (getters/setters only) and all logic lives in "Service" classes. This violates DDD's principle of rich domain objects. An `Order` should know how to `PlaceOrder()`, `Cancel()`, `ApplyDiscount()` — not delegate all logic to `OrderService`.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is an Entity in DDD?
**Your Response:** "An Entity is a domain object that has a unique and continuous **identity** throughout its entire lifecycle. Even if every single one of its attributes changes, it remains the same object as long as its ID—like a `userId` or `orderId`—is constant.

When designing Entities, I strive to avoid the **'Anemic Domain Model'** anti-pattern. Instead of creating objects that are just bags of getters and setters, I build 'rich' entities that contain their own business logic and enforce their own invariants. For instance, an `Order` entity shouldn't just have a `status` field; it should have a `Cancel()` method that knows whether the order is currently in a state that *can* be cancelled according to business rules."

---

### 4. What is a Value Object in DDD?

"A Value Object is an immutable object defined entirely by its **value**, with no identity. Two value objects with the same values are identical and interchangeable.

Classic examples: `Money(500, 'INR')`, `Address('Mumbai', 'Maharashtra', '400001')`, `DateRange(2024-01-01, 2024-12-31)`. There's no 'Money ID' or 'Address ID' — two `Address('Mumbai', 'MH', '400001')` objects are the same address.

Value Objects are immutable — instead of changing them, you create a new one. This eliminates an entire class of bugs where shared mutable state causes unexpected side effects. `money.add(100)` returns a new `Money(600, 'INR')`, it doesn't mutate the original."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** DDD-focused architecture interviews

#### Indepth
Value Object design:
```go
type Money struct {
    amount   int64
    currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
    if amount < 0 { return Money{}, ErrNegativeAmount }
    return Money{amount: amount, currency: currency}, nil
}

func (m Money) Add(other Money) (Money, error) {
    if m.currency != other.currency { return Money{}, ErrCurrencyMismatch }
    return Money{amount: m.amount + other.amount, currency: m.currency}, nil
    // Does NOT mutate m — returns new Money
}

func (m Money) Equals(other Money) bool {
    return m.amount == other.amount && m.currency == other.currency
}
```

When to use Value Object vs Entity:
- Does identity matter beyond the values? → Entity
- Is it interchangeable if values are the same? → Value Object
- Is there a lifecycle to track? → Entity

**Benefits of Value Objects:** Self-validating (constructor enforces invariants), expressive (code reads like domain language), safe to share (immutability).

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a Value Object in DDD?
**Your Response:** "A Value Object is an object that is defined entirely by its attributes, with no identity of its own. If two 'Address' objects have the exact same street, city, and zip code, they are identical and interchangeable. 

The two 'superpowers' of Value Objects are **immutability** and **self-validation**. Because they are immutable, you never have to worry about side effects when passing them between different services or threads—you treat them just like a primitive integer or string. Because they validate themselves in their constructor, you can be certain that every `Money` object in your system is mathematically valid, which makes the overall codebase much more robust and easier to reason about."

---

### 5. What is an Aggregate in DDD?

"An Aggregate is a **cluster of domain objects (entities and value objects) that are treated as a single unit** for data consistency. One entity in the cluster is the **Aggregate Root** — the only entry point for external access.

Example: An `Order` aggregate contains the `Order` entity (root) plus a list of `OrderItem` value objects and a `ShippingAddress` value object. No external code can directly add an `OrderItem` to the aggregate — they must call `order.AddItem(item)`. The `Order` enforces all invariants (max 10 items, total amount must be positive) before any item can be added.

This is the DDD solution to the 'data corruption through back-door access' problem."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (Principal), senior backend architecture roles

#### Indepth
Aggregate rules:
1. **One Aggregate Root per aggregate:** External code only references the root, never internal entities
2. **Consistency boundary:** All invariants are enforced within the aggregate boundary
3. **Small aggregates:** Keep aggregates small. Large aggregates cause lock contention and performance issues.
4. **Reference by ID:** Aggregates reference other aggregates by ID only, never by object reference

Order aggregate:
```
Order (Root) {
    OrderID
    CustomerID  ← reference to Customer aggregate by ID only
    OrderItems: [
        OrderItem(productId, qty, unitPrice)  ← value object
    ]
    ShippingAddress  ← value object
    TotalAmount      ← derived value
}
```

**Aggregate size trap:** A naive designer puts `Order`, `Customer`, `Product`, and `Inventory` all in one aggregate ("because they're related"). This creates a massive, slow, lock-prone aggregate. DDD principle: **keep aggregates as small as possible**. Reference across aggregates by ID, use eventual consistency (domain events) to keep them in sync.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is an Aggregate in DDD?
**Your Response:** "An Aggregate is a cluster of related domain objects—typically one Entity and several Value Objects—that are treated as a single **consistency unit** for data changes. The **Aggregate Root** is the only member of the aggregate that external objects are allowed to hold a reference to.

This is critical because the Aggregate Root acts as the 'gatekeeper' for all business invariants. For example, if you want to add a line item to an `Order`, you *must* go through the `Order` root. It will check if the line items exceed a weight limit or if the total balance is valid before allowing the change. This pattern ensures that your database is always in a valid state and avoids the 'backdoor' corruption that happens when multiple services are tweaking internal object states directly."

---

### 6. What is a Domain Service vs an Application Service in DDD?

"Both are services, but they belong to different layers:

**Domain Service:** Contains business logic that doesn't belong to any single entity or value object. It operates on multiple domain objects. Example: `TransferService.transfer(sourceAccount, targetAccount, amount)` — neither `Account` entity should own the transfer logic because it involves two accounts.

**Application Service:** Coordinates use cases — it's the entry point for a request. It loads aggregates via repositories, calls domain services, and coordinates the flow. It doesn't contain business rules itself. It's the orchestrator, not the decision-maker.

Rule of thumb: If the logic is a business rule ('transfers require sufficient balance'), it's in a Domain Service or Entity. If it's 'load Account from DB, call transfer logic, save, publish event', it's an Application Service."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Senior engineering roles at companies using Clean/Hexagonal architecture

#### Indepth
```
Application Service (PlaceOrderUseCase):
    1. Load Order from OrderRepository (infrastructure)
    2. Load Product from ProductRepository (infrastructure)
    3. Call PricingService.calculateTotal(order, products) ← Domain Service
    4. Call order.addItem(item) ← Entity method
    5. Save Order via OrderRepository ← infrastructure
    6. Publish OrderItemAdded event ← infrastructure

Domain Service (PricingService):
    - Contains complex business pricing rules
    - Operates on Order + Product domain objects
    - No infrastructure concerns
```

Application Services are thin orchestrators. Domain Services are rich business logic. This separation keeps domain logic testable without any infrastructure (no DB, no HTTP).

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a Domain Service vs an Application Service in DDD?
**Your Response:** "The distinction is about responsibilities. A **Domain Service** contains core business logic that doesn't naturally 'fit' inside any single Entity—usually because it involves coordinating multiple different Aggregates. For example, a `PaymentGateway` or a `TransferService` is a Domain Service.

An **Application Service**, on the other hand, is the 'orchestrator' or the entry point for a user request. It doesn't contain any business rules itself; its job is to load the necessary Aggregates from a database, call the Domain Service or Entity methods to do the real work, and then save the result and publish events. Think of the Application Service as the 'Manager' who delegates the work, and the Domain logic as the 'Expert' who actually knows the rules."

---

### 7. What is a Repository in DDD?

"A Repository is an **abstraction over the persistence layer** that presents domain objects as if they're in an in-memory collection. The domain doesn't know MySQL or PostgreSQL exists — it just calls `orderRepo.FindById(orderId)` and gets an `Order` back.

This decoupling is the dependency inversion principle in action. The domain defines the `OrderRepository` interface. The infrastructure provides a `PostgresOrderRepository` implementation. The domain depends on the interface; the implementation is swappable.

This makes unit testing trivial: inject a mock `OrderRepository` → test business logic without hitting a real database."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Clean architecture discussions

#### Indepth
```go
// Domain layer — defines the port
type OrderRepository interface {
    FindById(ctx context.Context, id OrderID) (*Order, error)
    Save(ctx context.Context, order *Order) error
    FindByCustomerID(ctx context.Context, custID CustomerID) ([]*Order, error)
}

// Infrastructure layer — provides the adapter
type postgresOrderRepository struct {
    db *sql.DB
}

func (r *postgresOrderRepository) FindById(ctx context.Context, id OrderID) (*Order, error) {
    // SQL query, map rows to Order aggregate
}
```

#### Indepth
Repository responsibilities:
- **Reconstruct aggregates:** Map DB rows → domain objects (NOT ANEMIC DTOs — fully loaded aggregates with behavior)
- **Persist aggregates:** Map domain objects → DB rows
- **Query by specific criteria:** `FindByStatus`, `FindByCustomer`, etc.

NOT a repository's job:
- Complex multi-table queries for reporting (use a dedicated Query service / read model)
- Caching (that's the application service's or an infrastructure cache layer's job)

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a Repository in DDD?
**Your Response:** "A Repository is an abstraction that allows the domain layer to treat its data as if it were in an in-memory collection. The core point is **Infrastructure Ignorance**—your business logic should NOT know if you're using MySQL, MongoDB, or an external API to store your orders.

We implement this using the **Dependency Inversion** principle. The Domain Layer defines an `OrderRepository` interface, and the Infrastructure Layer provides the implementation using a specific DB driver. This makes the system incredibly easy to test because you can swap out the heavy, slow SQL database for a fast, memory-based mock repository during unit tests, ensuring that your business rules are perfectly verifiable without any 'flaky' database dependencies."

---

### 8. What is anti-corruption layer (ACL) in DDD?

"An Anti-Corruption Layer is a translation layer that **isolates your domain model from an external or legacy system's model**. It translates between their vocabulary and yours so their complexity doesn't 'corrupt' your clean domain.

Scenario: Your new Order service needs to integrate with a legacy SAP ERP that has its own complex, messy data model for orders. Instead of mapping your `Order` entity directly to SAP's schema, you build an ACL that translates `SAPSalesDocument` → your `Order` domain object on input, and your `Order` → `SAPSalesDocument` on output.

Your domain stays clean and speaks your language. The ACL handles all the translation messiness of the external world."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Enterprises migrating from legacy systems, banking sector

#### Indepth
ACL patterns:
- **Adapter pattern:** ACL uses adapters to convert between models
- **Translator:** Maps fields: `SAPSalesDoc.VBELN` → `Order.id`
- **Facade:** Simplifies a complex external API to a clean interface your domain can use

When to use ACL:
- Integrating with a legacy system with poor model design
- Integrating with a third-party API (Stripe, Twilio) — translate their model before it enters your domain
- When an upstream bounded context doesn't align with your bounded context

Without ACL: Your `Order` entity starts having fields like `sapDocumentNumber`, `erpProcessingCode` — the legacy world bleeds into your domain.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is anti-corruption layer (ACL) in DDD?
**Your Response:** "An Anti-Corruption Layer (ACL) is a translation barrier that protects your domain model from being 'polluted' by a messy or legacy external system. When you integrate with an outside service—like a 20-year-old COBOL banking core—you don't want their archaic data structures and terminology leaking into your modern microservice.

The ACL acts as a **translator**. It takes the external system's strange inputs and converts them into the 'Ubiquitous Language' of your domain. This ensures that even if the external system is refactored or completely replaced, your core business logic stays 100% stable. It’s a defensive design pattern that keeps your architecture clean and focused on your own domain's needs."

---

### 9. What are domain events in DDD?

"Domain events are **facts that something significant happened in the domain**. They're raised by the aggregate and represent state changes that other parts of the system care about.

When `Order.place()` is called and business validation passes, the Order aggregate raises a `OrderPlaced` domain event. This event is then consumed by other bounded contexts (Inventory, Shipping, Analytics) — but the Order aggregate doesn't directly call them. It just raises the event.

Domain events are the bridge between bounded contexts — they're how you achieve loose coupling between aggregates and services while keeping strong boundaries. They complete the circle with Event Sourcing — domain events become the source of truth in an event-sourced system."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Senior architecture roles at product companies

#### Indepth
Domain events in code:
```go
type OrderPlaced struct {
    OrderID    OrderID
    CustomerID CustomerID
    TotalAmount Money
    PlacedAt   time.Time
}

// In Order aggregate:
func (o *Order) Place() error {
    if !o.isValid() { return ErrInvalidOrder }
    o.status = Placed
    o.AddDomainEvent(OrderPlaced{OrderID: o.id, ...})
    return nil
}

// Application service dispatches events after saving:
order.Place()
orderRepo.Save(order)
eventBus.PublishAll(order.DomainEvents())
```

- **Domain event:** Internal to the bounded context. Not published to external systems. Example: internal handlers react to `OrderPaymentReceived` within the Order service.
- **Integration event:** Published externally to Kafka. Example: `OrderPlaced` published for other services. Often a subset/transformation of the domain event.

#### 🗣️ How to Explain in Interview
**Interviewer:** What are domain events in DDD?
**Your Response:** "Domain Events are facts about something significant that just happened in the system—for example, `OrderPlaced` or `PaymentRefunded`. They are the primary way we achieve **loose coupling** between different parts of our architecture.

Instead of an `OrderService` directly calling the `EmailService` and the `ShippingService` (which creates a massive 'Big Ball of Mud'), the Order aggregate simply records an event. This event can then be picked up by internal handlers or transformed into **Integration Events** and sent to Kafka for other microservices. It allows the business to add new 'side effects'—like a Slack notification—without ever touching the core checkout code. It’s also the fundamental building block for **Event Sourcing** and Event-Driven Architectures."

---

### 10. How does DDD map to microservices decomposition?

"DDD Bounded Contexts provide the most principled basis for microservice decomposition. The rule: **one bounded context = one or more microservices**, but **never one microservice spanning multiple bounded contexts**.

Practical steps: (1) Run Event Storming to discover the domain — map out all events, commands, aggregates. (2) Cluster related aggregates into bounded contexts. (3) Each bounded context becomes a service candidate. (4) Evaluate for team fit (Conway's Law) and deployment independence.

The result: services that are genuinely independent because they own their own domain model, their own data, and their own vocabulary. Not services carved by database table or by technology layer."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Architecture lead, principal engineer roles

#### Indepth
Event Storming (a workshop to discover bounded contexts):
1. **Orange sticky notes:** Domain events (past tense) — `OrderPlaced`, `PaymentReceived`
2. **Blue sticky notes:** Commands that trigger events — `PlaceOrder`, `ChargePayment`
3. **Yellow sticky notes:** Actors (who issues the commands) — `Customer`, `System`
4. **Green sticky notes:** Read models (what data is needed) — `OrderSummaryView`
5. **Pink sticky notes:** External systems — Stripe, SMS Gateway

After Event Storming, draw **Context Lines** around clusters of related events/commands. These are your bounded contexts → your microservices.

E-commerce bounded contexts from Event Storming:
- **Catalog context:** ProductListed, ProductUpdated, PriceChanged
- **Order context:** CartItemAdded, OrderPlaced, OrderCancelled
- **Payment context:** PaymentInitiated, PaymentCompleted, PaymentRefunded
- **Fulfillment context:** OrderPacked, ShipmentDispatched, DeliveryCompleted
- **Notification context:** EmailSent, SMSSent, PushNotificationSent

Each of these becomes a service (or a small cluster of services). They share no databases and communicate via integration events.

#### 🗣️ How to Explain in Interview
**Interviewer:** How does DDD map to microservices decomposition?
**Your Response:** "DDD is arguably the most effective tool we have for defining **microservice boundaries**. The common trap in microservices is to split by database table—but that just leads to a 'Distributed Monolith.' DDD suggests we split by **Bounded Contexts**.

Every context owns its own domain model and its own database. By using a workshop like **'Event Storming'**, we can visualize the business flow and identify where the language shifts—for instance, when an 'Order' in the Checkout context becomes a 'Shipment' in the Logistics context. These natural 'fracture lines' are exactly where our service boundaries should be. The result is a system where teams are truly autonomous because they share no internal data or logic with anyone else."
