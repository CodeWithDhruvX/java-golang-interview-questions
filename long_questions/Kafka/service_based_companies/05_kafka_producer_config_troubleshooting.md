# 🏗️ Kafka — Producer Configuration & Troubleshooting

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Cognizant, Infosys, Wipro, HCL, Capgemini

---

## Q1. Explain the key producer performance configs: `batch.size`, `linger.ms`, and `compression.type`.

"When a Kafka producer sends messages, it doesn't always send one message at a time. It can batch multiple messages together to improve throughput and reduce network overhead. Three configs control this batching behavior:

**`batch.size` (default: 16384 bytes = 16KB):**
The maximum number of bytes in a single batch sent to a partition. The producer accumulates messages in an internal buffer until this size is reached, then sends the batch in one network call.
- **Increase this** (e.g., to 64KB or 256KB) for high-throughput producers where latency is not critical. More data per network trip = better throughput.
- **Risk:** If messages are small and slow to arrive, the batch may never fill — causing unnecessary latency.

**`linger.ms` (default: 0ms):**
How many milliseconds the producer waits before sending a batch, even if `batch.size` hasn't been reached. Default of `0` means send immediately (lowest latency, smallest batches).
- **Increase this** (e.g., to 5–20ms) to allow more messages to accumulate in the buffer, resulting in larger, more efficient batches.
- **Common combo:** `batch.size=65536` + `linger.ms=10` — the producer sends when either 64KB is accumulated OR 10ms has elapsed.

**`compression.type` (default: `none`):**
Compresses the batch before sending. Options: `none`, `gzip`, `snappy`, `lz4`, `zstd`.
- **`lz4`/`snappy`:** Fast compression with moderate compression ratio — ideal for most production systems where CPU is not a bottleneck.
- **`gzip`/`zstd`:** Higher compression ratio at the cost of more CPU — use when network bandwidth is the bottleneck (e.g., cross-region replication)."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring's `spring.kafka.producer` yaml bindings natively map `batch-size`, `linger-ms`, and `compression-type` to the underlying Java producer client efficiently.
* **Golang:** `kafka-go` handles these via standard structs: `Writer.BatchSize` (which operates on max messages, unlike Java's byte-limit), `Writer.BatchTimeout` (analogous to `linger.ms`), and `Writer.Compression` (setting it to `kafka.Snappy` or `kafka.Lz4`).

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Cognizant, HCL — these are among the most common Kafka configuration questions in L2 developer rounds at service companies. Interviewers check if the candidate can tune Kafka beyond default settings.

#### Indepth
**Buffer Memory (`buffer.memory`):** All batches accumulate in the producer's internal memory buffer (default: 32MB). If the producer generates messages faster than the broker can accept them and the buffer fills up, `KafkaProducer.send()` starts blocking for `max.block.ms` (default: 60s) before throwing a `BufferExhaustedException`. Monitor `buffer-available-bytes` metric to detect approaching exhaustion.

---

## Q2. What are common Kafka producer errors and how do you fix them?

"Here are the most frequently encountered producer errors in projects and their root cause + fix:

**1. `TimeoutException: Failed to update metadata after [X]ms`**
- **Cause:** Producer cannot connect to any broker. Typically a network/firewall issue or wrong `bootstrap.servers` config.
- **Fix:** Verify broker hostnames are resolvable from the application host. Check security group / firewall rules for port 9092.

**2. `RecordTooLargeException`**
- **Cause:** A single message exceeds the broker's `max.message.bytes` limit (default: 1MB) or the topic-level override.
- **Fix:** Either reduce message size (paginate the payload) or increase `max.message.bytes` on both broker and topic.

**3. `NotEnoughReplicasException`**
- **Cause:** Producer is using `acks=all` but the number of in-sync replicas has fallen below `min.insync.replicas`. This happens when brokers are down.
- **Fix:** This is by design — Kafka is refusing to accept writes to prevent data loss. Bring the brokers back online.

**4. `ProducerFencedException`**
- **Cause:** A transactional producer with the same `transactional.id` was started elsewhere (e.g., a duplicate deployment). Kafka fences the older producer instance.
- **Fix:** Ensure `transactional.id` is unique per producer instance.

**5. Serialization `ClassCastException`**
- **Cause:** Mismatch between the application object type produced and the serializer config.
- **Fix:** Ensure the key/value types exactly match the configured serializers."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Exception semantics are strongly typed using Kafka's official `org.apache.kafka.common.errors` hierarchy.
* **Golang:** In Go, you inspect the returned `error` from the `WriteMessages` call, often using type asserting or `errors.Is`/`errors.As` against Kafka specific errors provided in the client library.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** TCS, Wipro, Infosys — asked in project experience rounds where the interviewer probes whether the candidate has actually worked with Kafka in production and encountered real issues.

#### Indepth
**Retries and Idempotency:** Setting `retries=Integer.MAX_VALUE` on a non-idempotent producer can cause duplicate messages during transient failures. Always pair high retry counts with `enable.idempotence=true`, which suppresses duplicates by using producer sequence numbers. In Kafka 3.0+, idempotence is enabled by default.

---

## Q3. How do you test Kafka producer and consumer code in unit and integration tests?

"Testing Kafka logic requires different strategies at different levels:

**Unit Testing:**
We generally avoid real brokers in unit tests. We mock the producer to verify that the message is sent successfully without establishing a network connection.

**Integration Testing — Embedded Brokers:**
Spinning up an in-memory Kafka broker within the test environment — no Docker required. This allows full end-to-end integration testing.

**Testcontainers for Full Fidelity:**
For highest confidence, use Testcontainers to run a real Kafka Docker container during tests — running the exact same Kafka version as production, guaranteeing there are no embedded-specific limitations."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka provides `MockProducer` and `@EmbeddedKafka` out of the box. Testcontainers' `KafkaContainer` integrates effortlessly with `@DynamicPropertySource` for substituting properties dynamically in tests.
* **Golang:** Go rarely utilizes "embedded" Kafka instances. Instead, mocking is usually done via dependency injection using Go interfaces. For integration tests, the standard approach is using `testcontainers-go` to spin up a Kafka container orchestrator natively in the testing suite.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Capgemini, TCS — testing Kafka integration reliably is commonly asked since many junior developers only know how to test REST APIs and lack async messaging test strategies.

#### Indepth
**Async Assertions:** Since message consumption is asynchronous, using `time.Sleep()` or `Thread.sleep()` in tests is brittle. In Java, `Awaitility` provides polling-based assertions, while in Golang, native `select` block mechanisms over channels with a context timeout are utilized to wait deterministically for successful test consumptions.

---
