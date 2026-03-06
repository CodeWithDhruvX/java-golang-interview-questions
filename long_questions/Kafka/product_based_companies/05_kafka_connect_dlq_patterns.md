# 🏗️ Kafka — Connect, DLQ & Production Error Patterns

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Amazon, Flipkart, Paytm, Swiggy, Razorpay

---

## Q1. What is Kafka Connect and how does it fit into a data pipeline architecture?

"**Kafka Connect** is a scalable, fault-tolerant framework for streaming data between Kafka and external systems (databases, object stores, search engines) without writing producer/consumer code from scratch.

It provides two types of connectors:
- **Source Connectors:** Pull data FROM an external system INTO Kafka. Example: a JDBC Source Connector reads rows from a MySQL table and publishes them as Kafka events.
- **Sink Connectors:** Push data FROM Kafka INTO an external system. Example: an Elasticsearch Sink Connector indexes Kafka messages into Elasticsearch in real-time.

**Architecture:**
Kafka Connect runs as a cluster of **Worker** processes. Each connector is broken into **Tasks** that run in parallel across workers. Configuration is managed via a REST API — you can deploy, pause, or restart connectors without downtime.

**Production Use Case:**
A classic data pipeline at Flipkart:
```
MySQL (Orders DB)  →  [JDBC Source Connector]  →  Kafka  →  [Elasticsearch Sink Connector]  →  Product Search Index
                                                         →  [S3 Sink Connector]              →  Data Lake (Parquet)
```
The beauty is that the source writes to ONE topic, and multiple sink connectors fan-out to different destinations independently."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Kafka Connect operates entirely outside of standard Spring Boot or Golang execution environments. It is run directly on the JVM as a standalone cluster. The only intersection is that output topics fed by Connect act as native event producers representing the DB, which your Spring Boot or Golang services freely consume.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Flipkart — data ingestion at scale without bespoke producer code. A very common architecture pattern in data-heavy product companies.

#### Indepth
**Debezium — Change Data Capture (CDC):** Debezium is a special Kafka Connect Source that reads the database's binary replication log (binlog in MySQL, WAL in Postgres) instead of polling rows. This gives true, low-latency change event streams with ZERO load on production database queries, capturing INSERT/UPDATE/DELETE events as atomic Kafka messages.

---

## Q2. How do you implement a Dead Letter Queue (DLQ) in Kafka for poison pill messages?

"A **Poison Pill** is a message that consistently causes the consumer to crash or throw an exception during processing — corrupt data, unexpected schema, null fields, etc. Without a DLQ, the consumer retries indefinitely and halts all downstream processing for the entire partition.

**DLQ Strategy:**

```
Main Topic (orders)
       ↓
  Consumer (tries to process)
       ↓ (fails deserialization / business logic exception)
  DLQ Topic (orders.dlq)
       ↓
  Monitoring / Alert / Manual Review
```

**Implementation in Spring Kafka:**
```java
@Bean
public DefaultErrorHandler errorHandler(KafkaTemplate<String, String> kafkaTemplate) {
    DeadLetterPublishingRecoverer recoverer = new DeadLetterPublishingRecoverer(kafkaTemplate,
        (record, ex) -> new TopicPartition(record.topic() + ".dlq", record.partition()));

    ExponentialBackOffWithMaxRetries backOff = new ExponentialBackOffWithMaxRetries(3);
    backOff.setInitialInterval(1000L);
    backOff.setMultiplier(2.0);

    return new DefaultErrorHandler(recoverer, backOff);
}
```

**Implementation in Golang:**
```go
// Unlike Spring, Go requires manual DLQ routing natively in your consumer worker logic.
func processMessage(msg kafka.Message, writer *kafka.Writer) {
    err := process(msg)
    if err != nil && isFatal(err) {
        dlqPayload := string(msg.Value) + " | ERR: " + err.Error()
        writer.WriteMessages(context.Background(), kafka.Message{
            Topic: "orders.dlq",
            Key: msg.Key,
            Value: []byte(dlqPayload),
        })
    }
}
```

**DLQ Message Enrichment:** Before moving to DLQ, the framework automatically adds Kafka headers: `original-topic`, `original-partition`, `original-offset`. This lets the ops team replay or diagnose without losing context."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Paytm, Swiggy, Razorpay — critical in payment and order pipelines where a bad event must NEVER block the entire queue; the system should isolate and park bad messages, not halt.

#### Indepth
**Replay from DLQ:** After fixing the root cause (e.g., deploying a schema fix), ops teams use a script to read messages from the `.dlq` topic and re-publish them to the original topic with corrected data. This is a key operational capability that distinguishes senior engineers from juniors.

---

## Q3. What are the Kafka patterns for saga orchestration in distributed transactions?

"In microservices, a distributed transaction spanning multiple services (e.g., Reserve Inventory → Charge Payment → Ship Order) cannot use traditional ACID database transactions. The **Saga Pattern** breaks the workflow into a sequence of local transactions, each published as a Kafka event.

**Choreography-Based Saga (Event-Driven):**
Each service listens for events and publishes its own completion or failure events. No central orchestrator.

```
OrderService → publishes: OrderCreated
               ↓ (consumed by)
InventoryService → publishes: InventoryReserved / InventoryFailed
                   ↓ (consumed by)
PaymentService → publishes: PaymentCharged / PaymentFailed
                  ↓ (consumed by)
ShippingService → publishes: OrderShipped
```

**Compensating Transactions (Rollback):**
If `PaymentFailed`, the saga publishes a `ReleaseInventory` compensating event that the InventoryService listens to and rolls back its reservation. Kafka's immutable log guarantees the compensating event is never lost.

**Orchestration-Based Saga:**
A central **Saga Orchestrator** service drives the workflow by commanding each step via Kafka topic and awaiting reply events — simpler to trace but introduces a coupling point."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Complex sagas are heavily managed by frameworks like Eventuate Tram or Camunda Zeebe seamlessly inside Spring workloads.
* **Golang:** Go uses tightly bound temporal worker nodes interacting directly with Kafka, or relies explicitly on `Temporal.io` (a deeply integrated Go paradigm) which functionally abstracts away the Kafka wiring for Saga coordination.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Amazon, Flipkart — tested heavily at companies building large-scale order management systems where checkout spans inventory, payments, logistics, and notifications.

#### Indepth
**Outbox Pattern:** The canonical solution for the dual-write problem in Saga: when a service commits its local DB transaction, it ALSO writes the Kafka event into an `outbox` table in the SAME database transaction. A Debezium CDC connector then reads the outbox table and publishes to Kafka. This guarantees the event is NEVER lost even if Kafka is temporarily unavailable at the time of the business transaction.
---
