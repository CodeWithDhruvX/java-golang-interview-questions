# 🏗️ Kafka — Topic Design, Partition Strategy & Real-World Scenarios

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Cognizant, Wipro, Tech Mahindra, Hexaware

---

## Q1. How do you decide the number of partitions for a Kafka topic?

"Choosing the right partition count is one of the most consequential topic design decisions, since **you cannot decrease partitions** once set (you can only increase them, which can disrupt key-based ordering).

**The Throughput Formula:**
```
Required Partitions = max(Required Throughput / Producer Throughput per Partition,
                          Required Throughput / Consumer Throughput per Partition)
```

**Practical Rules of Thumb:**
1. **Start with a measured baseline:** Benchmark a single partition on your hardware. A standard SSD partition handles ~10–50 MB/s depending on message size and replication factor.

2. **Match consumer parallelism:** Partitions = Maximum number of consumers you'll ever want to run in parallel. If you set 3 partitions but need to scale to 10 consumers, 7 will sit idle.

3. **Avoid over-partitioning:** Each partition = open file handles + memory buffers on every broker (for replication). 10,000 partitions on a 3-broker cluster will strain broker memory significantly.

4. **A practical starting point:**
   - Low volume (\< 1k msg/sec): **3–6 partitions**
   - Medium volume (1k–50k msg/sec): **12–24 partitions**
   - High volume (50k+ msg/sec): **48–200+ partitions**

5. **Always make it a multiple of your replication factor** to ensure even distribution across brokers."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring developers often use the `TopicBuilder` or `@Bean NewTopic` in configuration classes to initialize topic partitions upfront programmatically.
* **Golang:** Topic planning is typically managed outside the core Go application (using Terraform or CLI tools), but if managed internally, `kafka-go` lacks robust topic administration. Developers often fallback to `confluent-kafka-go`'s `AdminClient` to programmatically provision `NumPartitions`.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Cognizant, Tech Mahindra — a very common design question asked when the interviewer moves from 'what is Kafka' to 'how would you use it in a real project'.

#### Indepth
**Partition Count and Ordering:** Kafka guarantees message order ONLY within a single partition, not across partitions. If you need strict global ordering (e.g., all events for a bank account processed in sequence), use a single partition per logical entity via a well-designed key. Using more partitions for a given key breaks ordering guarantees.

---

## Q2. What are Kafka topic naming conventions and design best practices?

"Good topic naming and design is critical in multi-team environments to maintain clarity, searchability, and governance.

**Naming Convention:**
```
<domain>.<entity>.<event-type>

Examples:
  payments.orders.created
  payments.orders.cancelled
  logistics.shipments.dispatched
  user-service.profiles.updated
  notifications.emails.sent
```

**Why this matters:**
- Immediately communicates the producing domain, the entity being described, and the event type.
- ACLs can be applied by prefix: `payments.*` → payments team, `logistics.*` → logistics team.
- Avoids collisions between teams (two teams cannot accidentally use the same name).

**Best Practices:**

| Practice | Reason |
|---|---|
| Use lowercase with dots or hyphens | Consistent across different Kafka clients; avoid underscores (conflicts with JMX metric names) |
| One event type per topic | Keeps consumers from filtering by event type in the consumer code; simplifies schema management |
| Separate topics for commands and events | `orders.place-order` (command) vs. `orders.order-placed` (event) |
| Prefix internal/system topics with `_` | e.g., `_dlq.payments.orders.created` makes DLQs immediately identifiable |
| Version topics for breaking changes | `user-events.v2` instead of migrating the existing topic |"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** In Spring, complex `@KafkaListener(topics = "${app.topics.orders-created}")` binding is used to centralize topic name management inside `application.yml` instead of hardcoding topic strings.
* **Golang:** Similar to Java, Go developers inject topic names via struct initialization and `Viper` configuration loaders to easily swap topic namespaces across environments (dev vs. prod).

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Wipro, Infosys, Hexaware — practical topic design questions appear when interviewers probe for project experience. Candidates who can only explain Kafka theory but not how to structure topics demonstrate a lack of real-world experience.

#### Indepth
**Compacted vs. Non-Compacted Topic Design:** Use compacted topics when the topic represents a **state table** (latest value per key matters, history doesn't): e.g., `user-profiles.current-state`. Use non-compacted (delete policy) for **event streams** where every event matters independently: e.g., `payments.transactions.completed`. Mixing these up is a common beginner mistake.

---

## Q3. Walk through a complete real-world Kafka use case: Order Management in an e-commerce system.

"This scenario covers the most commonly asked 'design a system using Kafka' question in service company interviews.

**Problem:** Design an order processing system for an e-commerce platform where placing an order triggers inventory reservation, payment charging, and email notification — all asynchronously.

**Topic Design:**
```
orders.created           → published by OrderService when order is placed
inventory.reserved       → published by InventoryService on successful reservation
inventory.failed         → published by InventoryService when stock unavailable
payments.charged         → published by PaymentService on success
payments.failed          → published by PaymentService on failure
notifications.email.send → published by any service needing to send email
```

**Event Flow (Choreography Saga):**
```
User → POST /orders → OrderService
  → validates order
  → publishes OrderCreated to `orders.created`

InventoryService (consumes `orders.created`)
  → checks stock
  → publishes InventoryReserved to `inventory.reserved`   [success]
  → OR publishes InventoryFailed to `inventory.failed`    [failure]

PaymentService (consumes `inventory.reserved`)
  → charges payment gateway
  → publishes PaymentCharged to `payments.charged`

NotificationService (consumes `payments.charged`)
  → sends order confirmation email via `notifications.email.send`

OrderService (consumes `payments.failed` or `inventory.failed`)
  → updates order status to FAILED
  → publishes compensating event to release inventory if needed
```

**Key Design Decisions:**
- **Partition key:** `orderId` — ensures all events for one order land in the same partition for ordering.
- **Consumer group per service:** Each microservice has its own consumer group so it independently processes its copy of events.
- **DLQ per topic:** `orders.created.dlq` to park bad events and prevent partition blocking."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** A Choreography Saga relies heavily on `KafkaTemplate` for reliable publication inside a distributed transaction context. Leveraging Spring Cloud Stream provides high-level functional abstractions for connecting these events to business logic.
* **Golang:** Go microservices excel in these Event-Driven architectures due to natively spinning up goroutines per event. The low memory footprint makes orchestrating a dozen small choreography services highly economical in Kubernetes.

#### 🏢 Company Context
**Level:** 🟡 Intermediate to 🔴 Senior | **Asked at:** TCS, Wipro, Cognizant, Tech Mahindra — this end-to-end scenario is the most common 'system design with Kafka' question at service companies. Being able to narrate this flow clearly with proper topic names, consumer groups, and error handling is what separates a strong candidate from an average one.

#### Indepth
**Idempotent Email Notification:** The NotificationService should store a `(orderId, eventType)` sent record in a database and check before sending. Since Kafka provides at-least-once delivery, the `PaymentCharged` event may be delivered multiple times. Without idempotency, the customer receives duplicate 'Your order is confirmed' emails — a classic Kafka integration bug in real projects.

---
