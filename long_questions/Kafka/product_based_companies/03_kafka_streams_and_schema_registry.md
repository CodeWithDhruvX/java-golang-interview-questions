# 🏗️ Kafka — Streams Processing & Schema Registry

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Swiggy, Razorpay, Hotstar

---

## Q1. What is Kafka Streams and how does it differ from Apache Flink or Spark Streaming?

"Kafka Streams is a **client-side Java library** — not a separate cluster — that allows you to build stateful, real-time stream processing applications. It reads from Kafka topics, processes the data, and writes results back to Kafka topics, all within your regular application process.

**Kafka Streams vs. Flink vs. Spark Streaming:**

| Dimension | Kafka Streams | Apache Flink | Spark Streaming |
|---|---|---|---|
| **Deployment** | Runs inside your app (no extra cluster) | Requires a dedicated Flink cluster | Requires a dedicated Spark cluster |
| **Latency** | True record-by-record (ms latency) | True record-by-record (ms latency) | Micro-batch (seconds latency) |
| **State Store** | RocksDB (embedded, local) | Managed state with checkpointing | In-memory RDD partitions |
| **Scaling** | Scales with Kafka partitions | Parallelism configured independently | Executors/tasks |
| **Best For** | Lightweight in-app stream processing | Complex CEP, ML pipelines, large-scale joins | Batch with streaming extension |

The key trade-off: if the team is already running a Java microservice and needs moderate stream processing with low operational overhead, Kafka Streams is the clear winner. For complex multi-stream windowed joins at petabyte scale, Flink wins."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, LinkedIn — evaluating architectural decision-making for data pipeline design.

#### Indepth
**Kafka Streams KTable:** A KTable is a changelog stream — an abstraction of an ever-updating table keyed by message key. It materializes a local RocksDB view of the latest value per key. Joins between a KStream (events) and KTable (state) are extremely powerful: for example, enriching each order event with the latest user profile data without an external database call.

---

## Q2. What is the Confluent Schema Registry and why is it essential in production Kafka systems?

"Without governance, producers and consumers evolve independently, leading to **schema drift** — a consumer crashes because a producer added/removed a field it didn't expect. The **Schema Registry** is a centralized catalog that stores and enforces versioned Avro/Protobuf/JSON schemas for Kafka topics.

**How it works:**
1. A producer registers its schema with the registry before sending messages.
2. The registry assigns a unique **schema ID** and validates it against the registered **compatibility policy**.
3. Instead of embedding the full schema in every message (wasteful), the producer only embeds the 4-byte `schema ID` in the message header.
4. The consumer fetches the schema by ID from the registry, deserializes the payload, and maps it to its own model class.

**Compatibility Modes:**
- `BACKWARD`: New consumer can read old messages (safe to add optional fields)
- `FORWARD`: Old consumers can still read new messages (safe to remove optional fields)
- `FULL`: Both backward and forward — the safest, most restrictive mode."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, Swiggy — critical in microservices where 10+ services produce/consume shared topics and independent deployment cycles can cause version mismatches.

#### Indepth
**Avro vs. Protobuf:** Avro embeds the schema inline in its binary encoding, making it self-describing but slightly bulkier. Protobuf uses field numbers to identify fields, resulting in smaller payloads but requiring both sides to have the compiled `.proto` files. In practice, Protobuf is favored in cross-language or gRPC-integrated shops, while Avro dominates in pure Java/Kafka-native environments.

---

## Q3. How do you implement windowed aggregations in Kafka Streams?

"Kafka Streams supports four types of time windows to aggregate events over a time period:

**1. Tumbling Windows:** Fixed, non-overlapping intervals. Example: count orders every 5 minutes. Each event belongs to exactly one window.

```java
KStream<String, Order> orders = builder.stream("orders-topic");
orders
  .groupByKey()
  .windowedBy(TimeWindows.ofSizeWithNoGrace(Duration.ofMinutes(5)))
  .count()
  .toStream()
  .to("order-counts-topic");
```

**2. Hopping Windows:** Fixed size but overlapping. Example: a 10-minute window that advances every 2 minutes — an event can belong to multiple windows. Used for rolling metrics.

**3. Session Windows:** Dynamic, activity-based. A window opens when a user is active and closes after a configurable inactivity gap. Perfect for user behavior analytics (e.g., 'session ended if idle for 30 minutes').

**4. Sliding Windows:** Windows defined by a specified time difference between records — used for comparing events that are close in time."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Hotstar, LinkedIn — testing real-world stream analytics capability for metrics dashboards and user engagement features.

#### Indepth
**Grace Period:** By default, Kafka Streams closes a window and emits results immediately. But late-arriving events (due to network delays) are dropped. Setting `.withGracePeriod(Duration.ofSeconds(30))` tells Kafka Streams to hold the window open for 30 extra seconds to accommodate late data before final emission.

---
