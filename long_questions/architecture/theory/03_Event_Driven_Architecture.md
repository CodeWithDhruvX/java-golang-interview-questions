# ⚡ Event-Driven Architecture — Questions 1–12

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Flipkart, Uber, Razorpay, Swiggy — any company using Kafka or message queues

---

### 1. What is event-driven architecture (EDA)?

"Event-driven architecture is a design paradigm where **components communicate by producing and consuming events** — asynchronously — rather than calling each other directly.

An event is a record of something that happened: `OrderPlaced`, `PaymentCompleted`, `ItemShipped`. The producer doesn't know or care who's listening. Any number of consumers can react to that event independently.

This decoupling is powerful: When Swiggy processes an order, it publishes `OrderPlaced`. The Delivery service picks it up to assign a driver, the Analytics service picks it up to update dashboards, and the Notification service picks it up to send an SMS. None of them are coupled to each other — only to the event's schema."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Swiggy, Zomato, Razorpay

#### Indepth
EDA vs request-driven architecture:
| Aspect | Request-Driven (REST/gRPC) | Event-Driven |
|--------|--------------------------|--------------|
| Coupling | Tight (knows about consumer) | Loose (fire and forget) |
| Availability | Consumer must be up | Consumer can be down, catches up later |
| Scalability | Consumer must scale with producer | Consumers scale independently |
| Semantics | Synchronous, immediate response | Asynchronous, eventual |
| Trade-off | Strong consistency, simple flow | Eventual consistency, complex debugging |

**When to use EDA:** Background jobs, cross-service workflows, audit log generation, real-time analytics, notification systems. When NOT to use: When you need an immediate response in the same request (e.g., "did payment succeed? show confirmation page").

---

### 2. What is CQRS (Command Query Responsibility Segregation)?

"CQRS separates the **write model** (Commands that change state) from the **read model** (Queries that return data). Instead of one unified data model for both, you have two separate models optimized for their respective operations.

The write side handles business logic and enforces invariants. It writes to a normalized database. The read side is a denormalized, pre-computed view optimized for queries — it can be a different database entirely (e.g., ElasticSearch for full-text search, Redis for fast API responses).

Classic example: An e-commerce order system. The Write side stores orders in PostgreSQL with full normalization. The Read side maintains a pre-joined, flattened view in MongoDB that includes order + product details + user info — exactly what the frontend needs, returned in one query."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, Groww, Zepto

#### Indepth
CQRS implementation:
```
Write Side:          Read Side:
Command → Validate → Update PostgreSQL (normalized)
                  └→ Publish Event → Update MongoDB (denormalized read model)
                                   → Update ElasticSearch (search view)
Query → Read from MongoDB / ElasticSearch (not PostgreSQL)
```

Benefits:
- **Scalability:** Read and write sides scale independently (reads are usually 10x writes)
- **Performance:** Read models are optimized for query patterns — no joins required
- **Flexibility:** Multiple read models from one event stream (mobile view, web view, analytics view)

Costs:
- **Eventual consistency:** Read model lags behind write model by a few milliseconds
- **Complexity:** Two data stores to maintain, synchronization logic
- **Debugging:** Query results may not reflect very recent writes

CQRS is often paired with Event Sourcing — the events that update the write side also update the read models.

---

### 3. What is Event Sourcing?

"Event sourcing is a pattern where instead of storing the **current state** of an entity, you store the **sequence of events** that led to that state.

Normal approach: Store `order = { status: 'SHIPPED', amount: 500 }` — the current snapshot.
Event Sourcing approach: Store the event log — `[OrderCreated, PaymentReceived, ItemPacked, OrderShipped]`. The current state is derived by replaying the events.

This gives you a complete audit trail for free. Banks love it — you can replay the transaction history from day 1 and recompute any balance at any point in time. You can also project the same event stream into multiple different read models."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, CRED, Groww, fintech companies with audit requirements

#### Indepth
Event Sourcing structure:
```
Event Store (append-only log):
1. OrderCreated    { orderId: 1, userId: 42, amount: 500 }
2. PaymentReceived { orderId: 1, transactionId: "TXN123" }
3. ItemPacked      { orderId: 1, packerId: "W007" }
4. OrderShipped    { orderId: 1, trackingId: "SHIP456" }

Current state = Replay(events 1..4) → { status: SHIPPED, transactionId: TXN123, ... }
```

Key advantages:
- **Temporal queries:** "What was the order status at 3pm yesterday?" — replay events up to that timestamp
- **Audit log:** No separate audit table needed — the event log IS the source of truth
- **Debug production issues:** Replay the exact sequence of events that caused a bug
- **Multiple projections:** Same events → Order status view, Financial ledger view, Analytics view

Challenges:
- **Event schema evolution:** Events are immutable. If you add a field, old events don't have it. Use *upcasting* or *event versioning*.
- **Snapshots for performance:** Replaying 10,000 events to get current state is slow. Periodically save snapshots.

---

### 4. What is the difference between events, commands, and queries?

"These are the three fundamental message types in a message-driven system:

**Commands** tell a service to do something — `PlaceOrder`, `CancelPayment`. They have one intended recipient and expect a response (success/failure). They're imperative — 'do this thing'.

**Events** announce that something happened — `OrderPlaced`, `PaymentFailed`. They have no intended recipient — anyone who cares can listen. They're past tense — 'this happened'.

**Queries** ask for data — `GetOrderStatus`, `FindUserById`. They have one recipient and expect data back, making no state changes. They're pure read operations.

This distinction matters because: Commands can fail and reject, events are broadcast facts, queries should be idempotent. Mixing them up leads to confusing APIs."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, PhonePe, pattern-heavy architecture discussions

#### Indepth
From Greg Young (who coined CQRS):
- **Command:** "Please do X" — can be refused. `ReserveInventory { itemId: 5, qty: 2 }`
- **Event:** "X happened" — can't be undone (past). `InventoryReserved { itemId: 5, qty: 2 }`
- **Query:** "Tell me Y" — side-effect free. `GetInventoryLevel { itemId: 5 }`

In REST API design:
- POST/PUT/DELETE → Commands (change state)
- GET → Queries (read state)
- Event → Kafka message (asynchronous broadcast)

**Responsibility boundary:** One service owns the command. Multiple services may listen to the resulting event. This is how you achieve loose coupling — the Order service doesn't call the Notification service. It publishes `OrderPlaced`. The Notification service subscribes.

---

### 5. What is the publish-subscribe pattern?

"Pub/Sub is a messaging pattern where **publishers send messages to topics** and **subscribers receive messages from topics they're interested in** — without knowing about each other.

The key: publishers and subscribers are completely decoupled. A publisher doesn't know how many subscribers exist. A new subscriber can be added without touching the publisher code.

Kafka topics are partitioned pub/sub brokers. Swiggy publishes `restaurant.order.placed` events. The Rider Allocation service subscribes. The Merchant Notification service subscribes. The Finance Reconciliation service subscribes. All receive the same event. Adding a new Analytics subscriber tomorrow requires zero changes to the Swiggy producer code."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company using Kafka, RabbitMQ, Google Pub/Sub

#### Indepth
Pub/Sub vs Message Queue:
| Aspect | Message Queue (Point-to-Point) | Pub/Sub |
|--------|-------------------------------|---------|
| Consumers | One consumer per message | Multiple consumers per message |
| Use case | Work distribution (only one worker should process) | Event notification (all interested parties) |
| Example | Video encoding job queue (only process once) | Order events (multiple services react) |
| Tools | RabbitMQ (default queue), SQS (standard) | Kafka, RabbitMQ exchanges (fanout), SNS, Google Pub/Sub |

**Kafka consumer groups:** Kafka's elegant hybrid solution. Multiple instances of the *same* service form a consumer group — messages are split across the group (like a queue). Multiple *different* services each have their own consumer group — they all get all messages (like pub/sub).

---

### 6. What is an event schema, and how do you handle schema evolution?

"An event schema defines the **structure and meaning of an event** — its fields, types, and what each field represents. It's the contract between producers and consumers.

Schema evolution is the challenge of changing that contract over time without breaking consumers. Since events are stored durably (Kafka retains them for days/weeks), a consumer reading old events must handle old schemas even after the producer has updated.

I use **schema registry** (Confluent Schema Registry with Apache Avro) to enforce backward/forward compatibility. Before a producer can publish with a new schema version, it's validated against the registry. Consumers declare which schema version they support."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** PhonePe, Razorpay, Groww — data-platform teams

#### Indepth
Schema evolution compatibility modes:
1. **Backward compatible:** New schema can read data written with old schema. New field with a default. Consumers using new schema can read old messages.
2. **Forward compatible:** Old schema can read data written with new schema. New field is optional. Old consumers can still read new messages (ignoring new field).
3. **Full compatible:** Both backward and forward. Most restrictive but safest.
4. **Breaking change:** Renaming a required field, changing a type. Requires coordination — consumer must be updated first (backward compatibility broken).

Safe evolution practices:
- **Add new fields:** Always with default values (backward compatible)
- **Never rename fields:** Create new field, deprecate old one
- **Never change field types:** `string` → `int` breaks consumers
- **Use versioning in topic names:** `orders.v1` → `orders.v2` for breaking changes

---

### 7. What is the difference between Kafka and RabbitMQ?

"Kafka is a **distributed event streaming platform** — it's fundamentally a durable, replayable, ordered log of events. It retains messages for days/weeks. Consumers can replay from any offset.

RabbitMQ is a **traditional message broker** — it routes messages from producers to queues to consumers, then deletes them once consumed. It's optimized for complex routing (fanout, topic routing, dead-letter queues).

My rule of thumb: Use Kafka when you need event sourcing, large throughput (millions of messages/sec), multiple consumer groups reading the same messages, or long-term retention. Use RabbitMQ for task queues with complex routing, lower throughput, and when you need the message to just 'go somewhere and be done'."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Swiggy, Zomato, Razorpay

#### Indepth
| Aspect | Kafka | RabbitMQ |
|--------|-------|----------|
| Paradigm | Event streaming log | Message broker |
| Message retention | Configurable (days/weeks) | Deleted after consumption |
| Replay | Yes (seek to offset) | No |
| Throughput | Millions/sec per partition | Thousands/sec |
| Ordering | Per partition | Per queue |
| Consumer model | Consumer groups (pull) | Push or pull |
| Routing | Topic + partition key | Exchange types (direct, fanout, topic) |
| Use cases | Event sourcing, EDA, stream processing | Task queues, RPC, complex routing |
| Horizontal scale | Add partitions, add brokers seamlessly | Requires federation/clustering setup |

**Kafka strengths:** LinkedIn processes 7 trillion messages per day on Kafka. Ideal for high-throughput event streams, real-time analytics (Kafka Streams), and audit logs.

**RabbitMQ strengths:** Complex routing patterns, short-lived messages, request-reply (simulated RPC), priority queues. Used by fintech companies for internal task distribution.

---

### 8. What is idempotency in event-driven systems?

"Idempotency means that **processing the same event multiple times produces the same result** as processing it once. It's non-negotiable in distributed systems because events are delivered **at least once** — your consumer may receive the same event twice due to retries or network issues.

Classic failure: Payment processing service receives `PaymentRequest` event. Processes it, charges the user, but crashes before acknowledging Kafka. Kafka retries — event is delivered again. The user is charged twice. That's a non-idempotent consumer.

Fix: Before processing, check a `{event_id}` lookup table. If the event was already processed, skip it. If not, process and record it. This makes the consumer idempotent — double delivery = no double charge."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** PhonePe, Razorpay, CRED, Groww — fintech is stringent about this

#### Indepth
Idempotency implementation patterns:
1. **Idempotency key table:**
```sql
CREATE TABLE processed_events (
    event_id UUID PRIMARY KEY,
    processed_at TIMESTAMP
);
-- Before processing: SELECT * FROM processed_events WHERE event_id = ?
-- If exists: skip. If not: process + INSERT. All in one transaction.
```

2. **Natural idempotency:** Some operations are naturally idempotent — setting a field to a value: `UPDATE orders SET status = 'SHIPPED'` is idempotent. Incrementing a counter: `UPDATE accounts SET balance = balance - 100` is NOT — it needs idempotency keys.

3. **Exactly-once semantics (Kafka transactions):** Kafka provides transactional producers that guarantee exactly-once delivery. More complex but eliminates the need for consumer-side deduplication. Used when state changes can't be made idempotent easily.

---

### 9. What is the inbox/outbox pattern pair?

"The outbox pattern ensures reliable event publishing from a producer. The inbox pattern ensures reliable event consumption by a consumer — together they guarantee **exactly-once processing** across services.

**Outbox (producer side):** Write your business change + the event to-publish in one DB transaction. A poller reads the outbox and publishes to Kafka. Guarantees at-least-once publishing.

**Inbox (consumer side):** When consuming an event from Kafka, write the event_id into an `inbox` table (idempotency key) in the same transaction as your business processing. If the event_id already exists, skip. Guarantees at-most-once processing.

Combined: outbox (at-least-once publish) + inbox (at-most-once process) = exactly-once end-to-end."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Fintech companies where data integrity is critical

#### Indepth
Flow diagram:
```
Producer Service:
  [Business DB + Outbox Table] → (Debezium CDC) → [Kafka Topic]
                                                         ↓
Consumer Service:                            [Consumer reads event]
                                                         ↓
                               [Check Inbox table: event_id seen?]
                                    YES ↓              NO ↓
                                   Skip          [Process + Insert inbox] (one transaction)
```

This pattern is particularly important in fintech:
- Payment service publishes `FundTransferRequested` via outbox
- Ledger service consumes it, checks inbox for deduplication
- Ledger service records debit + credit in same transaction as inbox insert
- Even if Kafka delivers the event 3 times, the ledger is only updated once

---

### 10. What is event choreography vs orchestration?

"Two patterns for coordinating multi-step workflows across microservices:

**Choreography:** No central coordinator. Each service reacts to events and publishes its own events. Order Service publishes `OrderCreated` → Inventory subscribes, reserves stock, publishes `StockReserved` → Payment subscribes, charges, publishes `PaymentDone` → Delivery subscribes, assigns driver. Each service knows what events to react to and what to emit.

**Orchestration:** A central saga orchestrator manages the workflow. It sends `Command: ReserveStock` to Inventory → on success, sends `Command: ChargePayment` to Payment → on success, sends `Command: AssignDelivery` to Delivery. The orchestrator has the full picture."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart, Swiggy — distributed transaction discussions

#### Indepth
| Aspect | Choreography | Orchestration |
|--------|-------------|---------------|
| Coordination | Decentralized, event-driven | Centralized orchestrator |
| Coupling | Low — services don't know each other | Medium — services know orchestrator |
| Visibility | Hard — must trace event chain | Easy — orchestrator logs full flow |
| Adding new step | New service subscribes to event | Orchestrator code changes |
| Failure handling | Each service compensates locally | Orchestrator centrally compensates |
| Tech | Kafka + events | Temporal.io, AWS Step Functions, Axon |

**Real world:** Choreography works well for simple 2-3 step flows. As complexity grows (conditional steps, timeouts, parallel branches, compensation), orchestration using Temporal.io becomes much easier to reason about. Uber uses Temporal.io for driver onboarding workflows with 20+ steps.

---

### 11. What is event streaming vs event messaging?

"**Event messaging** is one-time delivery — a message is sent, consumed once, and discarded. RabbitMQ is classic event messaging. Good for task queues where you just need one worker to pick up the job.

**Event streaming** treats events as a **durable, replayable log** — like a commit log for your entire system. Events are retained for days/weeks. New consumers can join and replay from the beginning. Kafka is event streaming. Good for when multiple services need to process the same events, or when you need historical replay.

The mental model: messaging is a phone call (real-time, gone when done), streaming is a podcast (recorded, anyone can listen anytime, you can rewind)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Data engineering roles, senior architecture roles

#### Indepth
Streaming enables powerful patterns impossible with messaging:
1. **Event replay:** "We have a new Analytics service — replay all orders from the past 6 months." Impossible with RabbitMQ (messages are gone). Trivial with Kafka (seek to offset 0).
2. **Multiple isolated consumers:** 5 different teams each have their own consumer group. Each gets every event independently. No coordination needed.
3. **Stream processing:** Kafka Streams or Apache Flink reads the event stream and computes running aggregates (e.g., "revenue in last 5 minutes"). Real-time dashboards at Zomato use this.
4. **Time-travel debugging:** Replay events leading up to a production incident to reproduce the bug.

Real-time pipeline at scale: Kafka → Kafka Streams (real-time aggregation) → Kafka → Consumers. This is how LinkedIn's real-time analytics works.

---

### 12. What are dead letter queues (DLQ)?

"A dead letter queue is a **holding queue for messages that failed processing** after exhausting retries. Instead of losing the message or blocking the main queue indefinitely, failed messages are moved to the DLQ for manual inspection and replay.

Scenario: The Notification service consumes `OrderPlaced` events. An event arrives with a malformed payload (e.g., null `userId`). Processing fails. After 3 retries, instead of discarding the event, it goes to the `orders.dlq` topic. An on-call engineer gets an alert, investigates, fixes the consumer bug, and replays messages from the DLQ back to the main topic.

Without a DLQ, that order notification is lost forever. With a DLQ, no data is lost — just delayed."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company operating production Kafka/SQS systems

#### Indepth
DLQ in different systems:
- **Kafka (manual):** Catch exceptions in consumer, publish failed message to `original-topic.dlq`. Include failure reason, stack trace, and retry count in the message headers.
- **AWS SQS:** Built-in DLQ support. Configure `maxReceiveCount` (e.g., 3). After 3 failed receives, message automatically moves to DLQ.
- **RabbitMQ:** Dead Letter Exchange (DLX) is configured per queue. Failed/TTL-expired messages route to DLX → DLQ.

DLQ operational practices:
1. **Alert on DLQ growth:** If DLQ message count grows, fire a PagerDuty alert immediately.
2. **Include metadata:** Add `failure_reason`, `original_topic`, `attempt_count`, `failed_at` headers to DLQ message.
3. **Replay mechanism:** Build a DLQ replay tool that reads from DLQ and re-publishes to the original topic.
4. **Poison pill detection:** If a single message fails 100 times, it's a "poison pill" — auto-archive it separately to prevent infinite retry storms.
