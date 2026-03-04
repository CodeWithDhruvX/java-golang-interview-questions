# 🏗️ Kafka — Offset Management & Delivery Guarantees

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Cognizant, Tech Mahindra, Capgemini

---

## Q1. What is an offset in Kafka and what is the difference between auto-commit and manual commit?

"An **offset** is a unique sequential integer assigned to every message within a partition. It acts as the consumer's bookmark — it tells the broker how far the consumer has successfully read. When a consumer restarts, it resumes from its last committed offset, avoiding reprocessing.

**Auto-Commit (`enable.auto.commit=true`, default):**
Kafka automatically commits the consumer's current offset in the background every `auto.commit.interval.ms` (default: 5 seconds). This is the simplest setup but has a critical risk: the consumer may fetch messages, start processing them, but the auto-commit fires BEFORE processing completes. If the consumer crashes mid-processing, those messages are marked as 'done' even though they weren't — leading to **data loss**.

**Manual Commit:**
The developer explicitly calls `consumer.commitSync()` or `consumer.commitAsync()` ONLY after successfully processing the message. This guarantees **at-least-once delivery** — messages are never marked done prematurely.

```java
// Manual commit example
ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
for (ConsumerRecord<String, String> record : records) {
    process(record); // your business logic
}
consumer.commitSync(); // only commit AFTER all records in the batch are processed
```"

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
- *Config:* `enable.auto.commit=false` + manual `commitSync()` after processing.
- *Use case:* Most production systems. Handle duplicates application-side using idempotency keys.

**3. Exactly-Once (No Loss, No Duplicates):**
Achieves true exactly-once using Kafka Transactions. The producer is idempotent (`enable.idempotence=true`), and the entire produce-process-commit cycle is wrapped in a single atomic transaction.
- *Config:* `enable.idempotence=true` + `isolation.level=read_committed` on consumer.
- *Use case:* Financial systems — payment processing where duplicates cause double debits."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Tech Mahindra, Infosys, Capgemini — this conceptual question is asked in virtually every Kafka interview. Candidates must know all three, their configs, and their appropriate use cases.

#### Indepth
**Idempotent Consumers:** Even with at-least-once delivery, you can simulate exactly-once behavior at the application layer. Store a unique `messageId` (from Kafka headers or the message body) in a database table. Before processing, check if the ID already exists. If yes, skip. This is the standard pattern when true Kafka EOS (Exactly-Once Semantics) is too complex or costly to implement.

---

## Q3. How does Kafka retain messages and how do you control it with `retention.ms` and `retention.bytes`?

"Kafka is fundamentally different from traditional queues — messages are NOT deleted after consumption. They persist in the broker's log for a configurable duration, regardless of whether any consumer has read them or not.

**`retention.ms` (Time-Based Retention):**
Specifies how long Kafka keeps a message. Default is `604800000` ms (7 days). After this period, the log segment containing the message is deleted.

```bash
# Set a topic to retain messages for 1 hour
kafka-topics.sh --alter \
  --topic my-topic \
  --config retention.ms=3600000
```

**`retention.bytes` (Size-Based Retention):**
Specifies the maximum total size of the log *per partition*. Once the partition log exceeds this size, the oldest segments are deleted, regardless of time.

```bash
# Limit partition log size to 500MB per partition
kafka-topics.sh --alter \
  --topic my-topic \
  --config retention.bytes=524288000
```

If BOTH are configured, whichever limit is hit FIRST triggers the deletion.

**Retention and Consumer Lag:**
If a consumer is offline for longer than `retention.ms`, when it restarts, the messages it missed are already deleted. The consumer will get an `OffsetOutOfRangeException`. The `auto.offset.reset=earliest` setting will cause it to jump to the oldest available offset rather than crash."

#### 🏢 Company Context
**Level:** 🟢 Junior to 🟡 Intermediate | **Asked at:** TCS, Wipro, Infosys — an essential operational knowledge question to assess if the candidate understands that Kafka is a log, not a traditional queue that auto-deletes consumed messages.

#### Indepth
**`log.retention.check.interval.ms`:** Kafka doesn't continuously scan for expired data. This config (default: 5 minutes) sets how often the Log Cleaner thread wakes up and checks for segments eligible for deletion. So in practice, a message may persist slightly longer than `retention.ms` until the next cleanup cycle runs.

---
