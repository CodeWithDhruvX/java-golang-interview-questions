# 🎯 Kafka — Scenario-Based Mock Interview Q&A

> **Level:** 🟡 Intermediate to 🟣 Architect
> **Asked at:** Amazon, Uber, Swiggy, Zomato, Flipkart, Netflix, LinkedIn

---

## Scenario 1: Design a Real-Time Ride Matching System (Uber-Style)

**Problem Statement:** Design Kafka's role in a ride-hailing platform where drivers broadcast their GPS location every 2 seconds and riders request rides. The system must match idle drivers to rider requests within 500ms.

---

**Candidate Answer:**

"I'd design three core Kafka topic pipelines:

**Topic Architecture:**
```
driver-location-updates  (12 partitions, keyed by driverId)
rider-requests           (12 partitions, keyed by riderId)
ride-matches             (12 partitions, keyed by riderId)
driver-status-changes    (12 partitions, keyed by driverId, compacted)
```

**Pipeline 1 — Driver Location Ingestion:**
Driver apps publish GPS coordinates at 2-second intervals. With 50,000 active drivers, this is ~25,000 msg/sec. Partition by `driverId` to preserve per-driver ordering.
```
Driver App → driver-location-updates → Location Aggregation Service
```
The location service maintains a geospatial index (H3 hexagon grid) in Redis, updated for every Kafka event.

**Pipeline 2 — Rider Request → Match:**
When a rider books, publish to `rider-requests`. A stream processor:
1. Reads the rider request + rider's location
2. Queries the Redis geospatial index for idle drivers within 1km
3. Selects the closest idle driver
4. Publishes match result to `ride-matches`
5. Publishes driver status update to `driver-status-changes` (IDLE → ASSIGNED)
```
Rider App → rider-requests → Matching Service (Consumer Pipeline)
                                    ↓
                              ride-matches → Notification Service
```

**Fault Tolerance:**
- `driver-status-changes` is a **compacted topic** — only the latest status per driver is retained. On restart, the matching service rebuilds driver-state from scratch.
- `rider-requests` uses **at-most-once** semantics — we prefer dropping rare duplicates over the complexity of deduplication (a double-match is worse than a retry UX).

**Latency Budget:**
```
Rider submits request → Kafka produce: ~5ms
Kafka consume + Redis lookup: ~15ms
Match result to rider app: ~10ms
Total E2E: ~30ms (well within 500ms SLA)
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Pipeline 2 would elegantly be solved entirely inside a Kafka Streams embedded topology processing the `KStream`, avoiding multiple DB hops context-switches.
* **Golang:** The microsecond scaling capabilities of Goroutines fetching off Kafka partitions while executing non-blocking context-aware Redis queries makes this pipeline hyper-performant in Go.

#### Indepth
**Surge Pricing Signal:** A stream topology reads from `rider-requests` and `driver-location-updates`, counts the ratio per geohex per 30-second window, and publishes surge multipliers to a `surge-signals` topic. Price service subscribes and updates pricing dynamically — all stream-processing, zero batch jobs.

---

## Scenario 2: Design a Real-Time Order Tracking System (Swiggy/Zomato-Style)

**Problem Statement:** 500,000 food orders are placed daily. Each order transitions through: `PLACED → ACCEPTED → PREPARING → PICKED UP → DELIVERED`. Design a Kafka-based system where the rider app, restaurant app, and customer app all stay in real-time sync with minimum latency.

---

**Candidate Answer:**

"This is a classic **event-driven state machine** problem. Each state transition is an event.

**Topic Design:**
```
order-events        (24 partitions, keyed by orderId, retention=7 days)
order-state         (24 partitions, keyed by orderId, compacted — current state only)
notification-outbox (12 partitions, keyed by userId)
```

**Producer Flow (Order Service):**
```json
// Every state transition publishes an event to 'order-events' 
{ 
  "orderId": "123", 
  "status": "PREPARING", 
  "timestamp": "2024-01-01T12:00:00Z" 
}
```

**State Aggregator:**
An application continually processes `order-events` building a state machine per key. The reduced aggregate representation per order is actively published out to `order-state`.

**Notification Fan-out:**
A dedicated consumer group reads `order-events` and publishes to `notification-outbox`:
```
PREPARING event → notify customer ('Your food is being prepared!')
PICKED_UP event → notify restaurant ('Rider has collected order')
DELIVERED event → notify customer + trigger payment settlement
```

**Why compacted `order-state`?**
The rider app opens the app 3 hours after the order. Instead of replaying 5 historical events, it subscribes to `order-state` and immediately gets the latest state for that `orderId`. Sub-second initial load."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring state machine integrated directly with `@KafkaListener` ensures robust state transitions executing exactly once utilizing Spring transaction capabilities.
* **Golang:** Because State-Machines and Event processing are heavily asynchronous, utilizing Go channels effectively routes incoming Kafka consumption payloads (`msg.Value`) onto dedicated workers processing state changes natively.

#### Indepth
**Exactly-Once for Payment Settlement:** The `DELIVERED` event triggers a payment deduction. This must be exactly-once. Use Kafka transactions: the consumer, payment deduction call, and offset commit are all wrapped in a single transaction. If any step fails, the entire transaction rolls back and retries — preventing double charges.

---

## Scenario 3: Real-Time Fraud Detection System (FinTech)

**Problem Statement:** A payment platform processes 2M transactions/day. Design a Kafka-based fraud detection pipeline that flags suspicious transactions within 200ms — without blocking the payment flow.

---

**Candidate Answer:**

"The key design principle: **fraud detection must be async and non-blocking**. The payment must be allowed to proceed, with fraud flags applied post-hoc (refund if fraud confirmed).

**Topic Architecture:**
```
raw-transactions        (48 partitions, keyed by userId)
fraud-signals           (12 partitions, keyed by transactionId)
flagged-transactions    (6 partitions)
fraud-model-results     (12 partitions)
```

**Pipeline 1 — Rule-Based Real-Time Detection (<50ms):**
Analyze velocity metrics. For example, an active windowing stream processor evaluating `> 5 transactions from the same userId within 60 seconds`. Offending windows push alerts into `fraud-signals`.

**Pipeline 2 — ML Model Enrichment (async, <200ms):**
A consumer reads `raw-transactions`, enriches each transaction with:
- 30-day spend history (from Redis/Cassandra)
- Device fingerprint
- Merchant risk score

Calls ML model and publishes result:
```
raw-transactions → ML Enrichment Service → fraud-model-results
```

**Pipeline 3 — Signal Aggregation:**
A stream app joins `fraud-signals` + `fraud-model-results` (with a 200ms join window). If both signals flag as fraudulent, publish to `flagged-transactions` → triggers account freeze and manual review queue.

**Why Kafka instead of synchronous API call?**
Synchronous: Payment API calls Fraud API → Fraud API calls ML → 200ms added to P99 payment latency.
Async Kafka: Payment completes in <20ms. Fraud runs in parallel. Refund issued if fraud confirmed within 5 minutes. This is exactly how Razorpay and Stripe handle fraud at scale."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** This is a premiere Kafka Streams usage. The joining of `fraud-signals` and `fraud-model-results` over a tight sliding time-window falls automatically onto `KStream.join(KStream, ValueJoiner, JoinWindows)`.
* **Golang:** This is generally segmented across separate Go workloads pushing models sequentially, as performing a real-time sliding-window join across two Kafka topics manually in Go memory is exceptionally complex to build safely.

#### Indepth
**Velocity Checks with Multiple Windows:** A sophisticated fraud engine runs three parallel stream topologies simultaneously: 1-minute window (velocity spike), 1-hour window (card testing pattern), 24-hour window (account take-over). Each publishes independent signals. The aggregation layer scores them with weights. This multi-window approach catches different attack vectors.

---

## Scenario 4: High-Scale Log Ingestion (Netflix / Hotstar Style)

**Problem Statement:** Design a Kafka pipeline that ingests 100TB/day of application logs from 50,000 microservices, routes them to Elasticsearch for search, S3 for cold storage, and a real-time alerting pipeline.

---

**Candidate Answer:**

"This is a **fan-out with tiered storage** problem.

**Ingestion Layer:**
```
Producer: Filebeat/Fluentd agent on each service VM
Topic: app-logs (96 partitions)
Throughput: ~1.15 GB/sec sustained = ~10Gbps
Broker config: acks=1 (log ingestion tolerates rare loss), linger.ms=20, batch.size=512KB
```

**Fan-out via Consumer Groups:**

| Consumer Group | Sink | LAG SLA |
|---|---|---|
| `logs-to-elastic` | Elasticsearch (last 7 days) | <30 seconds |
| `logs-to-s3` | S3 via Kafka Connect S3 Sink | <5 minutes |
| `logs-to-alerts` | Alerting pipeline | <10 seconds |

**Kafka Connect S3 Sink Config:**
```json
{
  "connector.class": "io.confluent.connect.s3.S3SinkConnector",
  "tasks.max": "24",
  "s3.bucket.name": "netflix-logs-archive",
  "flush.size": "100000"
}
```

**Real-Time Alerting Pipeline:**
Consume logs, filter by 'ERROR' level, group by service name, create 30-second windows. If the error count exceeds a threshold (e.g. 1000), emit to an `error-rate-alerts` topic.

**Topic Retention:**
```
app-logs: retention.ms=86400000 (24 hours) — enough for all consumers to catch up
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Both applications largely only interact with the Alerting pipeline here, listening natively with Spring `@KafkaListener` or Go `kafka.Reader` to the output `error-rate-alerts` topic, connecting them directly into PagerDuty SDK calls.

#### Indepth
**Tiered Storage for Cost Reduction:** At 100TB/day, keeping 7 days of Kafka retention would require 700TB of broker disk — prohibitively expensive. Enable Kafka Tiered Storage (KIP-405, Kafka 3.6+): hot data (last 4 hours) stays on broker SSDs, cold data is automatically offloaded to S3. Consumers transparently fetch from either tier. This cuts broker storage costs by **90%** while preserving full replay capability.

---

## Scenario 5: When NOT to Use Kafka?

**Problem Statement:** Your tech lead says 'use Kafka for everything'. When would you push back?

---

**Candidate Answer:**

"I'd push back in these cases:

| Use Case | Better Choice | Reason |
|---|---|---|
| Simple task queue (send email on signup) | RabbitMQ / SQS | Kafka is overengineered; RabbitMQ has better TTL, priority queues, and dead-lettering for simple task flows |
| Request-Reply RPC pattern | gRPC / REST | Kafka has no native correlation ID mechanism; building RPC on Kafka adds unnecessary complexity |
| < 1000 msg/day event pipeline | Database outbox pattern | Full Kafka cluster for low-volume events is massive operational overhead |
| Real-time query (< 5ms) | Redis | Kafka is a log, not a cache. Interactive queries add latency; Redis is the right tool |
| Schema-heavy OLTP events with complex joins | PostgreSQL CDC → Debezium → Kafka | Pushing raw OLTP events into Kafka without CDC means duplicating transactional logic; Debezium preserves change semantics |

**The core question:** Does this problem need a **durable, ordered, replayable event log with multiple independent consumers**? If yes → Kafka. If it just needs 'send a message from A to B' → there are simpler tools."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Cloud integration simplifies dropping in RabbitMQ binders rather than forcefully tying simple email jobs to Kafka out of dogma.
* **Golang:** The Go community's ethos favors gRPC for synchronous RPC responses directly—bypassing heavy middleware brokers almost entirely when true Point-to-Point architectures satisfy the requirement.
---
