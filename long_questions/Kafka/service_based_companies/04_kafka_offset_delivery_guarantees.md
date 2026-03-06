# 🏗️ Kafka — Offset Management & Delivery Guarantees

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Cognizant, Tech Mahindra, Capgemini

---

## Q1. What is an offset in Kafka and what is the difference between auto-commit and manual commit?

"An **offset** is a unique sequential integer assigned to every message within a partition. It acts as the consumer's bookmark — it tells the broker how far the consumer has successfully read. When a consumer restarts, it resumes from its last committed offset, avoiding reprocessing.

**Auto-Commit (`enable.auto.commit=true`, default):**
Kafka automatically commits the consumer's current offset in the background every `auto.commit.interval.ms` (default: 5 seconds). This is the simplest setup but has a critical risk: the consumer may fetch messages, start processing them, but the auto-commit fires BEFORE processing completes. If the consumer crashes mid-processing, those messages are marked as 'done' even though they weren't — leading to **data loss**.

**Manual Commit:**
The developer explicitly calls `consumer.commitSync()` or `consumer.commitAsync()` ONLY after successfully processing the message. This guarantees **at-least-once delivery** — messages are never marked done prematurely."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** With Spring, setting `enable.auto.commit=false` isn't enough; you must also configure the container's AckMode (e.g., `MANUAL_IMMEDIATE`) and invoke `Acknowledgment.acknowledge()` or use `KafkaConsumer.commitSync()` natively.
* **Golang:** With `kafka-go`, developers use `reader.FetchMessage()` (which does NOT auto-commit, unlike `ReadMessage`) and then explicitly call `reader.CommitMessages()` after the specific goroutine or process is fully completed.

#### 🏢 Company Context
**Level:** 🟢 Junior to 🟡 Intermediate | **Asked at:** TCS, Cognizant — one of the most frequent foundational questions in any Kafka interview round, asked to test the practical understanding of data reliability.

#### Indepth
**`commitSync()` vs. `commitAsync()`:** `commitSync()` blocks the consumer thread until the broker confirms the commit — safer but slower. `commitAsync()` is non-blocking and faster but provides no guarantee on success. Production patterns often use `commitAsync()` for regular processing and fall back to `commitSync()` on shutdown/rebalance to ensure the final offset is committed reliably.

---

## Q2. Explain the three delivery guarantees in Kafka: At-Most-Once, At-Least-Once, and Exactly-Once.

"Kafka's delivery semantics describe what happens to a message during consumer failures or retries:

**1. At-Most-Once (Possible Data Loss):**
The consumer commits the offset FIRST and then processes the message. If the consumer crashes during processing, the message offset is already committed, so when it restarts it moves forward — the message is permanently skipped and lost.
- *Config:* `enable.auto.commit=true` with immediate commit, or `acks=0` on producer.
- *Use case:* Real-time analytics where occasional data loss is acceptable (e.g., website click counters).

**2. At-Least-Once (Possible Duplicates):**
The consumer processes FIRST, then commits the offset. If the consumer crashes after processing but before committing, the message is redelivered and processed again.
- *Config:* `enable.auto.commit=false` + manual commit after processing.
- *Use case:* Most production systems. Handle duplicates application-side using idempotency keys.

**3. Exactly-Once (No Loss, No Duplicates):**
Achieves true exactly-once using Kafka Transactions. The producer is idempotent (`enable.idempotence=true`), and the entire produce-process-commit cycle is wrapped in a single atomic transaction.
- *Config:* `enable.idempotence=true` + `isolation.level=read_committed` on consumer.
- *Use case:* Financial systems — payment processing where duplicates cause double debits."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring supports Exactly-Once Semantics (EOS) simply by adding `@Transactional` to the `@KafkaListener` containing producer calls, given that idempotence is enabled.
* **Golang:** True exactly-once semantics using the transaction API requires `confluent-kafka-go` (librdkafka based), as `segmentio/kafka-go` lacks full implementation of the transactional API.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Tech Mahindra, Infosys, Capgemini — this conceptual question is asked in virtually every Kafka interview. Candidates must know all three, their configs, and their appropriate use cases.

#### Indepth
**Idempotent Consumers:** Even with at-least-once delivery, you can simulate exactly-once behavior at the application layer. Store a unique `messageId` (from Kafka headers or the message body) in a database table. Before processing, check if the ID already exists. If yes, skip. This is the standard pattern when true Kafka EOS (Exactly-Once Semantics) is too complex or costly to implement.

---

## Q3. How does Kafka retain messages and how do you control it with `retention.ms` and `retention.bytes`?

"Kafka is fundamentally different from traditional queues — messages are NOT deleted after consumption. They persist in the broker's log for a configurable duration, regardless of whether any consumer has read them or not.

**`retention.ms` (Time-Based Retention):**
Specifies how long Kafka keeps a message. Default is `604800000` ms (7 days). After this period, the log segment containing the message is deleted.

**`retention.bytes` (Size-Based Retention):**
Specifies the maximum total size of the log *per partition*. Once the partition log exceeds this size, the oldest segments are deleted, regardless of time.

If BOTH are configured, whichever limit is hit FIRST triggers the deletion.

**Retention and Consumer Lag:**
If a consumer is offline for longer than `retention.ms`, when it restarts, the messages it missed are already deleted. The consumer will get an `OffsetOutOfRangeException`. The `auto.offset.reset=earliest` setting will cause it to jump to the oldest available offset rather than crash."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** This is a broker-side topic config, entirely independent of the Java consumer itself, except for reacting to `OffsetOutOfRangeException`.
* **Golang:** In Go, if a reader requests an offset out of range using `kafka-go`, it surfaces as an `OffsetOutOfRange` error, which developers must explicitly trap and handle according to the desired fault tolerance model.

#### 🏢 Company Context
**Level:** 🟢 Junior to 🟡 Intermediate | **Asked at:** TCS, Wipro, Infosys — an essential operational knowledge question to assess if the candidate understands that Kafka is a log, not a traditional queue that auto-deletes consumed messages.

#### Indepth
**`log.retention.check.interval.ms`:** Kafka doesn't continuously scan for expired data. This config (default: 5 minutes) sets how often the Log Cleaner thread wakes up and checks for segments eligible for deletion. So in practice, a message may persist slightly longer than `retention.ms` until the next cleanup cycle runs.
---
