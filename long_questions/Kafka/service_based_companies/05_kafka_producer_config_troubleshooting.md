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

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Cognizant, HCL — these are among the most common Kafka configuration questions in L2 Java developer rounds at service companies. Interviewers check if the candidate can tune Kafka beyond default settings.

#### Indepth
**Buffer Memory (`buffer.memory`):** All batches accumulate in the producer's internal memory buffer (default: 32MB). If the producer generates messages faster than the broker can accept them and the buffer fills up, `KafkaProducer.send()` starts blocking for `max.block.ms` (default: 60s) before throwing a `BufferExhaustedException`. Monitor `buffer-available-bytes` metric to detect approaching exhaustion.

---

## Q2. What are common Kafka producer errors and how do you fix them?

"Here are the most frequently encountered producer errors in projects and their root cause + fix:

**1. `TimeoutException: Failed to update metadata after [X]ms`**
- **Cause:** Producer cannot connect to any broker. Typically a network/firewall issue or wrong `bootstrap.servers` config.
- **Fix:** Verify broker hostnames are resolvable from the application host. Check security group / firewall rules for port 9092 (or 9093 for SSL).

**2. `RecordTooLargeException`**
- **Cause:** A single message exceeds the broker's `max.message.bytes` limit (default: 1MB) or the topic-level override.
- **Fix:** Either reduce message size (paginate the payload, use a reference key + object store for large blobs) or increase `max.message.bytes` on both broker and topic.

**3. `NotEnoughReplicasException`**
- **Cause:** Producer is using `acks=all` but the number of in-sync replicas has fallen below `min.insync.replicas`. This happens when brokers are down.
- **Fix:** This is by design — Kafka is refusing to accept writes to prevent data loss. Bring the brokers back online. Do NOT lower `min.insync.replicas` as a quick fix in production.

**4. `org.apache.kafka.common.errors.ProducerFencedException`**
- **Cause:** A transactional producer with the same `transactional.id` was started elsewhere (e.g., a duplicate deployment). Kafka fences the older producer instance.
- **Fix:** Ensure `transactional.id` is unique per producer instance. In Kubernetes, using `pod-name` as the transactional ID suffix ensures uniqueness.

**5. Serialization `ClassCastException`**
- **Cause:** Mismatch between the Java object type produced and the serializer config. For example, producing an `Integer` key but configured with `StringSerializer`.
- **Fix:** Ensure the key/value types in `KafkaProducer<K, V>` match the configured serializers exactly."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** TCS, Wipro, Infosys — asked in project experience rounds where the interviewer probes whether the candidate has actually worked with Kafka in production and encountered real issues.

#### Indepth
**Retries and Idempotency:** Setting `retries=Integer.MAX_VALUE` on a non-idempotent producer can cause duplicate messages during transient failures. Always pair high retry counts with `enable.idempotence=true`, which suppresses duplicates by using producer sequence numbers. In Kafka 3.0+, idempotence is enabled by default.

---

## Q3. How do you test Kafka producer and consumer code in unit and integration tests?

"Testing Kafka logic requires different strategies at different levels:

**Unit Testing — `MockProducer` and `MockConsumer`:**
Spring Kafka and the native Kafka client both provide mock implementations for testing without a real broker.

```java
// Unit test using MockProducer
@Test
void testOrderProducer() {
    MockProducer<String, String> mockProducer = new MockProducer<>(
        true, new StringSerializer(), new StringSerializer());

    OrderProducer producer = new OrderProducer(new KafkaTemplate<>(
        new DefaultKafkaProducerFactory<>(Map.of(), mockProducer)));

    producer.sendOrder("ORD-001", "{'item': 'laptop'}");

    List<ProducerRecord<String, String>> history = mockProducer.history();
    assertEquals(1, history.size());
    assertEquals("ORD-001", history.get(0).key());
    assertEquals("orders-topic", history.get(0).topic());
}
```

**Integration Testing — `@EmbeddedKafka`:**
Spring Kafka provides `@EmbeddedKafka` to spin up an in-memory Kafka broker within the test JVM — no Docker required.

```java
@SpringBootTest
@EmbeddedKafka(
    partitions = 1,
    topics = {"orders-topic"},
    brokerProperties = {"listeners=PLAINTEXT://localhost:9092"}
)
class OrderFlowIntegrationTest {

    @Autowired
    private OrderProducer producer;

    @Autowired
    private KafkaConsumer testConsumer;  // your listener bean

    @Test
    void testEndToEndOrderFlow() throws InterruptedException {
        producer.sendOrder("ORD-002", "{'item': 'phone'}");
        // Use CountDownLatch or Awaitility to wait for consumption
        assertTrue(testConsumer.getLatch().await(5, TimeUnit.SECONDS));
        assertEquals("ORD-002", testConsumer.getLastReceivedKey());
    }
}
```

**Testcontainers for Full Fidelity:**
For highest confidence, use `KafkaContainer` from Testcontainers to run a real Kafka Docker container during tests — same Kafka version as production, no embedded limitations."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Capgemini, TCS — testing Kafka integration reliably is commonly asked in rounds where the project involves Kafka since most junior developers only know how to test REST APIs via Postman and lack async messaging test strategies.

#### Indepth
**`Awaitility` for Async Assertions:** Since message consumption is asynchronous, using `Thread.sleep()` in tests is brittle. Awaitility provides polling-based assertions: `Awaitility.await().atMost(10, SECONDS).until(() -> consumer.hasReceivedMessage())` — it polls repeatedly until the condition is true or timeout is reached.

---
