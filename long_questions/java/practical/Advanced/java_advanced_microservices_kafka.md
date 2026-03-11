# Java Advanced — Microservices, Kafka & Distributed Systems

> **Topics:** Kafka Producer/Consumer, Spring Kafka (`@KafkaListener`, `KafkaTemplate`), Consumer Groups, offset management, partitions, dead letter topics, Circuit Breaker (Resilience4j), Saga, Outbox pattern, CQRS, service discovery, distributed tracing (MDC/Micrometer)

---

## 📋 Reading Progress

- [ ] **Section 1:** Kafka Fundamentals (Q1–Q15)
- [ ] **Section 2:** Spring Kafka — Producer & Consumer (Q16–Q27)
- [ ] **Section 3:** Resilience Patterns — Circuit Breaker, Retry, Bulkhead (Q28–Q38)
- [ ] **Section 4:** Distributed Patterns — Saga, Outbox, CQRS (Q39–Q47)
- [ ] **Section 5:** Observability — Tracing, Metrics, Service Discovery (Q48–Q55)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Kafka Fundamentals (Q1–Q15)

### 1. Kafka Core Concepts — Topics, Partitions, Offsets
**Q: What is the relationship?**
```
Topic: "orders"
├── Partition 0: [msg@offset0] [msg@offset1] [msg@offset2]  ← append-only log
├── Partition 1: [msg@offset0] [msg@offset1]
└── Partition 2: [msg@offset0] [msg@offset1] [msg@offset2] [msg@offset3]

Key concepts:
- Topic: named stream of records
- Partition: ordered, immutable log (append-only)
- Offset: position of a record within a partition (monotonically increasing)
- Retention: records kept for a configurable period (e.g., 7 days)
```
**A:** Kafka retains all messages in append-only logs. Consumers track their position (offset) per partition — they can replay from any offset. Partitions enable parallelism. Key-based routing sends same-key messages to the same partition (preserving order per key).

---

### 2. Producer — Key/Value and Partitioning
**Q: What determines which partition a message goes to?**
```java
import org.apache.kafka.clients.producer.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws Exception {
        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("key.serializer",   "org.apache.kafka.common.serialization.StringSerializer");
        props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");

        KafkaProducer<String, String> producer = new KafkaProducer<>(props);

        // With key: partition = hash(key) % numPartitions — same key → same partition
        ProducerRecord<String, String> record = new ProducerRecord<>("orders", "user-123", "order-data");
        RecordMetadata meta = producer.send(record).get();
        System.out.printf("topic=%s partition=%d offset=%d%n", meta.topic(), meta.partition(), meta.offset());

        // Without key: round-robin across partitions
        producer.send(new ProducerRecord<>("orders", null, "keyless-order"));
        producer.close();
    }
}
```
**A:** With key: `partition = murmur2(key) % numPartitions` — guarantees same-key messages land in the same partition (preserving order per key). Without key: round-robin (Kafka 2.4+: sticky partitioner for batching).

---

### 3. Consumer — Poll Loop
**Q: What is the correct Kafka consumer pattern?**
```java
import org.apache.kafka.clients.consumer.*;
import java.time.Duration;
import java.util.*;

public class Main {
    public static void main(String[] args) {
        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("group.id", "order-processor");
        props.put("key.deserializer",   "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("enable.auto.commit", "false"); // manual offset commit

        try (KafkaConsumer<String, String> consumer = new KafkaConsumer<>(props)) {
            consumer.subscribe(List.of("orders"));
            while (true) {
                ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
                for (ConsumerRecord<String, String> r : records) {
                    System.out.printf("partition=%d offset=%d key=%s value=%s%n",
                        r.partition(), r.offset(), r.key(), r.value());
                    processOrder(r.value()); // process before commit!
                }
                consumer.commitSync(); // commit AFTER processing
            }
        }
    }
}
```
**A:** `poll()` is the heartbeat mechanism — must be called frequently or the consumer is considered dead (and rebalance happens). Commit offsets only **after** successful processing. `enable.auto.commit=false` + manual `commitSync()` = at-least-once delivery.

---

### 4. Consumer Group — Partition Assignment
**Q: How many consumers can process in parallel?**
```
Topic "orders" with 4 partitions:

Consumer Group "order-service" with 2 consumers:
  Consumer A → Partition 0, Partition 1
  Consumer B → Partition 2, Partition 3

Consumer Group "order-service" with 4 consumers:
  Consumer A → Partition 0
  Consumer B → Partition 1
  Consumer C → Partition 2
  Consumer D → Partition 3

Consumer Group "order-service" with 5 consumers:
  Consumers A-D → one partition each
  Consumer E → IDLE (no partition assigned)
```
**A:** Maximum parallelism = number of partitions. Extra consumers beyond partition count are idle. Multiple consumer groups each get a full copy of all messages (broadcast pattern). Increase partitions to scale beyond current limit.

---

### 5. Delivery Semantics — At Most Once, At Least Once, Exactly Once
**Q: What does each mean?**
```java
// AT MOST ONCE: commit before processing → possible message loss
consumer.commitSync();
processOrder(record.value()); // if crash here, message is lost

// AT LEAST ONCE: commit after processing → possible duplicate on crash
processOrder(record.value());
consumer.commitSync(); // if crash before this, message is reprocessed

// EXACTLY ONCE: Kafka Transactions (Kafka 0.11+)
producer.beginTransaction();
producer.send(new ProducerRecord<>("processed-orders", key, value));
consumer.commitSync(offsets); // offset + produce are atomic
producer.commitTransaction();
// Transactional consumers see only committed messages
```
**A:** AT MOST ONCE: messages can be lost. AT LEAST ONCE: messages can be duplicated (most common — make consumers idempotent). EXACTLY ONCE: requires transactions + `isolation.level=read_committed` on consumer.

---

### 6. Replication Factor & Leader Election
**Q: What happens when a broker fails?**
```
Topic "orders", replication-factor=3, 3 brokers:

Partition 0:
  Leader   → Broker 1 (handles all reads/writes)
  Follower → Broker 2 (in-sync replica)
  Follower → Broker 3 (in-sync replica)

Broker 1 fails:
  Kafka controller elects new leader from ISR (in-sync replicas)
  Leader → Broker 2 (automatically)
  Follower → Broker 3

Producer/Consumer reconnect to new leader automatically.
min.insync.replicas=2 ensures at least 2 brokers acknowledged the write.
```
**A:** Replication provides fault tolerance. A topic survives up to `replication-factor - 1` broker failures. `acks=all` + `min.insync.replicas=2` ensures no message loss even if one broker fails.

---

### 7. acks Configuration — Durability Trade-offs
**Q: What do acks=0, acks=1, acks=all mean?**
```java
Properties props = new Properties();
props.put("bootstrap.servers", "localhost:9092");

// acks=0: fire and forget — fastest, possible message loss
props.put("acks", "0");

// acks=1: leader acknowledges — leader crash may cause loss
props.put("acks", "1");

// acks=all (or -1): all ISRs acknowledge — no loss (requires min.insync.replicas)
props.put("acks", "all");
props.put("min.insync.replicas", "2"); // in broker config
props.put("retries", "3");
props.put("enable.idempotence", "true"); // prevents duplicate sends on retry
```
**A:** `acks=0` → fastest, lossy. `acks=1` → medium safety. `acks=all` → strongest durability. Always use `enable.idempotence=true` with `acks=all` — prevents duplicate messages if the producer retries.

---

### 8. Consumer Lag — What It Means
**Q: What is consumer lag?**
```
Consumer lag = (latest offset in partition) - (committed offset of consumer group)

Example for Topic "orders" Partition 0:
  Latest offset: 1000
  Consumer "order-service" committed: 950
  Lag: 50 messages behind

Monitoring:
  kafka-consumer-groups.sh --describe --group order-service
  Or use Kafka JMX metrics: kafka.consumer:type=consumer-fetch-manager-metrics,attribute=records-lag-max

High lag means:
  - Consumer is too slow
  - Consumer crashed
  - Throughput spike
  Fix: add consumers (up to partition count) or optimize processing
```
**A:** Consumer lag is the key operational metric for Kafka consumers. Alert when lag exceeds a threshold. Monitor with Kafka Exporter + Prometheus + Grafana.

---

### 9. Compacted Topics — Keep Last Value Per Key
**Q: What is log compaction?**
```
Normal topic (retention by time/size):
  [key=A, val=1] [key=B, val=2] [key=A, val=3] [key=A, val=4]
  → all retained until retention expires

Compacted topic (cleanup.policy=compact):
  After compaction:
  [key=A, val=4] [key=B, val=2]
  → only the latest value per key is retained

Use cases:
  - Configuration/settings changelog
  - Database change capture (CDC)
  - Event sourcing snapshots

@Bean
public NewTopic settingsTopic() {
    return TopicBuilder.name("app-settings")
        .partitions(1)
        .replicas(1)
        .config(TopicConfig.CLEANUP_POLICY_CONFIG, TopicConfig.CLEANUP_POLICY_COMPACT)
        .build();
}
```
**A:** Compacted topics are key-value stores on Kafka — always have the latest value for each key, regardless of retention. Used by Kafka Streams for state stores and database CDC.

---

### 10. Schema Registry — Avro Serialization
**Q: Why use Schema Registry?**
```java
import io.confluent.kafka.serializers.KafkaAvroSerializer;
import org.apache.avro.Schema;
import org.apache.avro.generic.GenericRecord;

Properties props = new Properties();
props.put("bootstrap.servers", "localhost:9092");
props.put("schema.registry.url", "http://localhost:8081");
props.put("key.serializer",   KafkaAvroSerializer.class.getName());
props.put("value.serializer", KafkaAvroSerializer.class.getName());

// Avro schema: {"type":"record","name":"Order","fields":[{"name":"id","type":"long"},...]}
// Producer sends schema ID + binary payload (not full JSON)
// Consumer fetches schema from registry to deserialize

// Benefits:
// 1. Schema evolution with backward/forward compatibility
// 2. Binary format = smaller messages
// 3. Compatibility enforcement before deploy
```
**A:** Schema Registry stores Avro/JSON/Protobuf schemas and enforces compatibility. Producers register schemas; consumers fetch them by ID. Prevents breaking changes from being deployed without compatibility check.

---

### 11. Kafka Streams — Real-Time Processing
**Q: What does this Kafka Streams topology do?**
```java
import org.apache.kafka.streams.*;
import org.apache.kafka.streams.kstream.*;

public class WordCount {
    public static void main(String[] args) {
        StreamsBuilder builder = new StreamsBuilder();

        KStream<String, String> textLines = builder.stream("text-input");

        textLines
            .flatMapValues(line -> Arrays.asList(line.split("\\s+")))
            .groupBy((key, word) -> word)
            .count(Materialized.as("word-count-store"))
            .toStream()
            .to("word-count-output");

        KafkaStreams streams = new KafkaStreams(builder.build(), config);
        streams.start();
    }
}
```
**A:** Reads from `text-input`, splits lines into words, counts per word, writes counts to `word-count-output`. Kafka Streams maintains state in a local RocksDB store (`word-count-store`). Fault-tolerant via changelog topics.

---

### 12. Dead Letter Topic (DLT) — Handling Poison Pills
**Q: What is a poison pill message?**
```java
// A message that always fails processing — causes consumer to retry forever!
// Solution: dead letter topic

@KafkaListener(topics = "orders", errorHandler = "deadLetterPublishingRecoverer")
public void consume(OrderEvent event) {
    if (!isValid(event)) {
        throw new IllegalArgumentException("invalid order: " + event.getId());
    }
    processOrder(event);
}

@Bean
public DeadLetterPublishingRecoverer deadLetterPublishingRecoverer(KafkaTemplate<Object, Object> template) {
    return new DeadLetterPublishingRecoverer(template,
        (record, ex) -> new TopicPartition("orders.DLT", record.partition()));
}
// Failed messages → orders.DLT topic for separate investigation/replay
```
**A:** A "poison pill" causes infinite retries and blocks all subsequent messages. DLT sends failed messages to a separate topic after max retries. Operations team can inspect, fix, and replay DLT messages.

---

### 13. Idempotent Consumer — Handling Duplicates
**Q: How do you make a consumer idempotent?**
```java
@Service
class OrderConsumer {
    @Autowired OrderRepository repo;

    @KafkaListener(topics = "orders")
    @Transactional
    public void consume(ConsumerRecord<String, OrderEvent> record) {
        // Use message's idempotency key (e.g., orderId) to check for duplicates
        String idempotencyKey = record.key();         // producer sets unique key
        if (repo.existsByIdempotencyKey(idempotencyKey)) {
            log.info("Duplicate message, skipping: {}", idempotencyKey);
            return;  // already processed — safe to skip
        }
        Order order = processOrder(record.value());
        order.setIdempotencyKey(idempotencyKey);
        repo.save(order);
    }
}
```
**A:** Idempotent consumers handle duplicates by checking if the message was already processed (using a unique key stored in DB). Essential with at-least-once delivery. Alternatively use `IF NOT EXISTS` inserts or database unique constraints.

---

### 14. Kafka Headers — Metadata Propagation
**Q: How do you pass metadata with a Kafka message?**
```java
// Producer — add headers
ProducerRecord<String, String> record = new ProducerRecord<>("orders", "key", "value");
record.headers().add("correlationId", UUID.randomUUID().toString().getBytes());
record.headers().add("sourceService", "order-service".getBytes());
producer.send(record);

// Consumer — read headers
@KafkaListener(topics = "orders")
public void consume(ConsumerRecord<String, String> record) {
    Header corrId = record.headers().lastHeader("correlationId");
    String id = corrId != null ? new String(corrId.value()) : "unknown";
    MDC.put("correlationId", id); // for distributed tracing
    processOrder(record.value());
    MDC.clear();
}
```
**A:** Kafka headers propagate metadata (correlation IDs, trace IDs, source service) alongside messages without polluting the payload. Essential for distributed tracing across service boundaries.

---

### 15. Partition Assignment Strategy
**Q: What are the partition assignors?**
```java
// RangeAssignor (default): assigns consecutive partitions per consumer
// Consumer A → P0, P1, P2 | Consumer B → P3, P4, P5

// RoundRobinAssignor: distributes more evenly across topics
// Each partition assigned in turn to consumers

// StickyAssignor: minimizes partition movement during rebalances
// Keeps previous assignments as much as possible

Properties props = new Properties();
props.put("partition.assignment.strategy",
    "org.apache.kafka.clients.consumer.CooperativeStickyAssignor");
// INCREMENTAL rebalancing: only reassign necessary partitions
// Eliminates the "stop-the-world" nature of traditional rebalances
```
**A:** `CooperativeStickyAssignor` (Kafka 2.4+) is preferred — revokes only the partitions that need to move, allowing other partitions to keep consuming during rebalance. Traditional assignors block all consumers during rebalance.

---

## Section 2: Spring Kafka — Producer & Consumer (Q16–Q27)

### 16. KafkaTemplate — Sending Messages
**Q: What is the difference between send and sendDefault?**
```java
import org.springframework.kafka.core.*;

@Service
class OrderProducer {
    @Autowired KafkaTemplate<String, OrderEvent> kafkaTemplate;

    public void send(OrderEvent event) {
        // Explicit topic
        CompletableFuture<SendResult<String, OrderEvent>> future =
            kafkaTemplate.send("orders", event.getId().toString(), event);

        // Callback
        future.whenComplete((result, ex) -> {
            if (ex != null) {
                log.error("Failed to send: {}", ex.getMessage());
            } else {
                log.info("Sent to partition={} offset={}",
                    result.getRecordMetadata().partition(),
                    result.getRecordMetadata().offset());
            }
        });
    }
}
```
**A:** `kafkaTemplate.send()` is async — returns `CompletableFuture`. `future.get()` blocks until broker ACK. Always handle failures — implement retry logic or dead-letter queues.

---

### 17. @KafkaListener — Consuming Messages
**Q: What are the concurrency and batch options?**
```java
@Component
class EventConsumer {
    // Single message processing
    @KafkaListener(
        topics = "orders",
        groupId = "order-processor",
        concurrency = "3" // 3 consumer threads = 3 partitions consumed in parallel
    )
    public void process(OrderEvent event,
                        @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
                        @Header(KafkaHeaders.OFFSET) long offset) {
        System.out.printf("Processing from partition=%d offset=%d: %s%n", partition, offset, event.getId());
    }

    // Batch processing
    @KafkaListener(topics = "orders", containerFactory = "batchFactory")
    public void processBatch(List<OrderEvent> events) {
        System.out.println("Batch of " + events.size() + " events");
        // process all at once — more efficient for DB bulk inserts
    }
}
```
**A:** `concurrency` creates multiple consumer threads (each gets its own partition). `concurrency` ≤ partition count (extras sit idle). Batch mode processes multiple records per `poll()` in one invocation — great for bulk DB operations.

---

### 18. @KafkaListener — Error Handling
**Q: How do you configure retry and DLT in Spring Kafka?**
```java
@Configuration
class KafkaConfig {
    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, String> kafkaListenerContainerFactory(
        ConsumerFactory<String, String> cf,
        KafkaTemplate<String, String> template) {

        var factory = new ConcurrentKafkaListenerContainerFactory<String, String>();
        factory.setConsumerFactory(cf);

        // Retry 3 times with 2s backoff, then send to DLT
        factory.setCommonErrorHandler(new DefaultErrorHandler(
            new DeadLetterPublishingRecoverer(template),
            new FixedBackOff(2000L, 3) // 2s delay, 3 retries
        ));
        return factory;
    }
}
```
**A:** `DefaultErrorHandler` replaces `SeekToCurrentErrorHandler` (deprecated). `FixedBackOff` / `ExponentialBackOff` control retry timing. After max retries, `DeadLetterPublishingRecoverer` publishes to `topic.DLT`. Non-retryable exceptions can be classified to skip retries.

---

### 19. Transactions in Spring Kafka
**Q: How do you make produce + DB update atomic?**
```java
@Service
class OrderService {
    @Autowired KafkaTemplate<String, OrderEvent> kafkaTemplate;
    @Autowired OrderRepository repo;

    @Transactional // Kafka + DB in same transaction
    public void createOrder(OrderEvent event) {
        Order order = repo.save(new Order(event));    // DB write
        kafkaTemplate.send("orders", event.getId().toString(), event); // Kafka write
        // Both commit or both rollback together (with transactional producer)
    }
}

// application.yml:
// spring.kafka.producer.transaction-id-prefix: tx-
```
**A:** With `transaction-id-prefix`, `KafkaTemplate` uses a transactional producer. Combined with `@Transactional`, the DB commit and Kafka produce are coordinated (best-effort; true atomic requires Outbox pattern for 100% reliability).

---

### 20. KafkaListener with Manual Offset Commit
**Q: When do you need manual offset commit?**
```java
@KafkaListener(topics = "orders")
public void consume(ConsumerRecord<String, OrderEvent> record,
                    Acknowledgment ack) {
    try {
        processOrder(record.value());
        ack.acknowledge(); // commit this offset
    } catch (TransientException e) {
        // don't ack — message will be redelivered
        log.warn("Transient failure, will retry: {}", e.getMessage());
    }
}

// application.yml:
// spring.kafka.listener.ack-mode: manual
```
**A:** `AckMode.MANUAL` gives precise control — commit only after successful processing. Important when processing involves external calls that can fail transiently. Don't ack on transient failures to trigger redelivery.

---

### 21. Topic Auto-Creation — @Topic Bean
**Q: How do you create topics programmatically?**
```java
@Configuration
class TopicConfig {
    @Bean
    public NewTopic ordersTopic() {
        return TopicBuilder.name("orders")
            .partitions(6)
            .replicas(3)
            .config(TopicConfig.RETENTION_MS_CONFIG, String.valueOf(7 * 24 * 60 * 60 * 1000L)) // 7 days
            .config(TopicConfig.MIN_IN_SYNC_REPLICAS_CONFIG, "2")
            .build();
    }

    @Bean
    public NewTopic ordersDltTopic() {
        return TopicBuilder.name("orders.DLT")
            .partitions(6)
            .replicas(3)
            .build();
    }
}
```
**A:** Spring Kafka creates topics defined as `NewTopic` beans at startup using the `KafkaAdmin`. If the topic already exists, it's not recreated (use `KafkaAdmin.setAutoCreate(false)` to disable). Partitions can't be reduced once created.

---

### 22. Consumer Rebalance Listener
**Q: When do you need to handle rebalance events?**
```java
@Component
class RebalanceAwareConsumer implements ConsumerRebalanceListener {
    public void onPartitionsRevoked(Collection<TopicPartition> partitions) {
        System.out.println("Lost partitions: " + partitions);
        // Flush any in-flight batch processing
        // Commit offsets for in-progress work
    }
    public void onPartitionsAssigned(Collection<TopicPartition> partitions) {
        System.out.println("Assigned partitions: " + partitions);
        // Initialize state for new partitions
    }
}
```
**A:** Rebalance happens when a consumer joins/leaves a group or a partition count changes. During rebalance, consumption is paused (SWT is "stop-the-world" with traditional assignors). `onPartitionsRevoked` is the last chance to commit offsets.

---

### 23. Exactly-Once Semantics — Transactions
**Q: How do you implement EOS in Kafka?**
```java
Properties props = new Properties();
props.put("transactional.id", "my-transaction-id-1"); // unique per producer instance
props.put("enable.idempotence", "true");
props.put("acks", "all");

KafkaProducer<String, String> producer = new KafkaProducer<>(props);
producer.initTransactions(); // called once

producer.beginTransaction();
try {
    producer.send(new ProducerRecord<>("output", key, value));
    // Also commit consumer offsets within the transaction:
    producer.sendOffsetsToTransaction(currentOffsets, new ConsumerGroupMetadata("my-group"));
    producer.commitTransaction();
} catch (ProducerFencedException e) {
    producer.close(); // another instance took over
} catch (Exception e) {
    producer.abortTransaction();
}
```
**A:** EOS requires: transactional producer + `sendOffsetsToTransaction` (atomic consume+produce) + consumer with `isolation.level=read_committed`. Complex to implement manually — use Kafka Streams or Spring's transaction support.

---

### 24. KafkaTemplate — Sending with Headers (Distributed Tracing)
**Q: How do you propagate trace context via Kafka?**
```java
@Service
class TracedProducer {
    @Autowired KafkaTemplate<String, String> kafkaTemplate;

    public void send(String topic, String key, String payload) {
        // Spring Micrometer/Sleuth auto-propagates trace context via headers
        // For manual propagation:
        ProducerRecord<String, String> record = new ProducerRecord<>(topic, key, payload);
        String traceId = MDC.get("traceId"); // from Micrometer context
        if (traceId != null) {
            record.headers().add("X-B3-TraceId", traceId.getBytes());
        }
        kafkaTemplate.send(record);
    }
}
```
**A:** Distributed tracing across Kafka requires propagating trace headers. Micrometer Tracing + Spring Cloud Sleuth (or Brave) auto-instruments `KafkaTemplate` to inject/extract trace context, creating spans across producer → consumer boundaries.

---

### 25. Multi-Topic Listener
**Q: What does this consume from?**
```java
@KafkaListener(
    topics = {"orders", "returns", "cancellations"},
    groupId = "fulfillment-service"
)
public void handleAll(ConsumerRecord<String, String> record) {
    System.out.printf("topic=%s key=%s value=%s%n",
        record.topic(), record.key(), record.value());
    switch (record.topic()) {
        case "orders"       -> processOrder(record.value());
        case "returns"      -> processReturn(record.value());
        case "cancellations"-> processCancel(record.value());
    }
}
```
**A:** One listener can consume from multiple topics. All are in the same consumer group. `record.topic()` tells you which topic the message came from. Alternative: separate `@KafkaListener` per topic for cleaner code.

---

### 26. Kafka Connect — Source and Sink Connectors
**Q: How does Kafka Connect work?**
```json
// Source connector: MySQL → Kafka (Debezium CDC)
{
  "name": "mysql-source",
  "config": {
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "database.hostname": "mysql", "database.port": "3306",
    "database.user": "debezium", "database.password": "secret",
    "database.server.name": "mydb",
    "table.include.list": "mydb.orders"
  }
}
// Every INSERT/UPDATE/DELETE on orders → Kafka topic "mydb.orders"

// Sink connector: Kafka → Elasticsearch
// (ElasticsearchSinkConnector reads from topic, indexes into ES)
```
**A:** Kafka Connect streams data between Kafka and external systems without writing code. Debezium is the most popular CDC source connector — captures DB changes in real-time using binary log (binlog/WAL).

---

### 27. Consumer Configuration — Key Settings
**Q: What do these settings control?**
```java
Properties props = new Properties();
// How long to wait for messages in poll() call
props.put("fetch.min.bytes", "1024");       // wait until 1KB available
props.put("fetch.max.wait.ms", "500");      // or 500ms max wait

// How much data to fetch per partition
props.put("max.partition.fetch.bytes", "1048576"); // 1MB

// Max time between poll() calls before consumer considered dead
props.put("max.poll.interval.ms", "300000"); // 5 min — increase for slow processing
props.put("max.poll.records", "500");         // max batch size per poll

// Session timeout: heartbeat must arrive within this time
props.put("session.timeout.ms", "30000");   // 30s
props.put("heartbeat.interval.ms", "10000"); // must be < session.timeout / 3
```
**A:** `max.poll.interval.ms` is critical — if processing takes longer, the consumer is kicked out and rebalance triggered. Either increase it, reduce `max.poll.records`, or move processing async.

---

## Section 3: Resilience Patterns — Circuit Breaker, Retry, Bulkhead (Q28–Q38)

### 28. Circuit Breaker — States
**Q: What are the three states?**
```
CLOSED (normal operation):
  All requests pass through
  Track failures: if failure rate > threshold → OPEN

OPEN (protecting the service):
  All requests fail immediately (no calls to downstream)
  Wait for waitDurationInOpenState (e.g., 30s) → HALF_OPEN

HALF_OPEN (recovery probe):
  Allow limited requests (permittedNumberOfCallsInHalfOpenState)
  If they succeed → CLOSED
  If they fail → OPEN again
```
**A:** Circuit Breaker prevents cascading failures. When downstream is down, failing fast (OPEN state) frees threads and returns errors quickly instead of hanging. Introduced by Michael Nygard in "Release It!".

---

### 29. Resilience4j Circuit Breaker — Spring Boot
**Q: What does this configuration do?**
```java
import io.github.resilience4j.circuitbreaker.annotation.*;
import io.github.resilience4j.circuitbreaker.*;

@Service
class PaymentService {
    @CircuitBreaker(name = "paymentGateway", fallbackMethod = "fallbackPayment")
    public PaymentResult charge(PaymentRequest req) {
        return externalGateway.charge(req); // may fail
    }

    public PaymentResult fallbackPayment(PaymentRequest req, Throwable t) {
        log.warn("Circuit breaker open, using fallback: {}", t.getMessage());
        return PaymentResult.queued(req.getId()); // return a safe fallback
    }
}

// application.yml:
// resilience4j.circuitbreaker.instances.paymentGateway:
//   failureRateThreshold: 50       # open if 50% of calls fail
//   waitDurationInOpenState: 30s
//   slidingWindowSize: 10          # measure last 10 calls
//   permittedNumberOfCallsInHalfOpenState: 3
```
**A:** Resilience4j is the Hystrix replacement. `@CircuitBreaker` wraps the method with a circuit breaker. `fallbackMethod` provides a safe response when the circuit is open. Always provide a meaningful fallback.

---

### 30. Retry Pattern with Resilience4j
**Q: What does this retry?**
```java
@Service
class ExternalApiService {
    @Retry(name = "inventoryApi", fallbackMethod = "defaultInventory")
    public int getStock(Long productId) {
        return inventoryClient.getStock(productId); // network call
    }

    public int defaultInventory(Long productId, Exception e) {
        return -1; // unknown stock
    }
}

// application.yml:
// resilience4j.retry.instances.inventoryApi:
//   maxAttempts: 3
//   waitDuration: 1s
//   enableExponentialBackoff: true
//   exponentialBackoffMultiplier: 2
//   retryExceptions:
//     - java.io.IOException
//     - java.net.SocketTimeoutException
//   ignoreExceptions:
//     - com.example.ResourceNotFoundException    # don't retry 404
```
**A:** Retry with exponential backoff: 1st retry after 1s, 2nd after 2s, 3rd after 4s. Only retry transient errors (network issues). Never retry business errors (validation failures, 404s). After max attempts → fallback.

---

### 31. Bulkhead — Thread Pool Isolation
**Q: What problem does Bulkhead solve?**
```java
@Service
class RecommendationService {
    @Bulkhead(name = "recommendations", type = Bulkhead.Type.THREADPOOL)
    public CompletableFuture<List<Item>> getRecommendations(Long userId) {
        return CompletableFuture.supplyAsync(() -> mlModel.predict(userId));
    }
}

// application.yml:
// resilience4j.thread-pool-bulkhead.instances.recommendations:
//   maxThreadPoolSize: 4
//   coreThreadPoolSize: 2
//   queueCapacity: 10
// → If ML model is slow, max 4+10=14 requests can pile up, rest rejected immediately
```
**A:** Bulkhead isolates resource usage per service. Without it, a slow downstream (ML model) can exhaust the main thread pool — starving all other requests. Thread pool bulkhead caps concurrent calls to slow services.

---

### 32. Rate Limiter — Throttling
**Q: What does this limit?**
```java
@Service
class SmsService {
    @RateLimiter(name = "smsApi", fallbackMethod = "rateLimitedFallback")
    public void sendSms(String phone, String message) {
        twilioClient.send(phone, message);
    }

    public void rateLimitedFallback(String phone, String message, RequestNotPermitted e) {
        log.warn("SMS rate limit exceeded, queuing for later");
        smsQueue.add(phone, message);
    }
}

// application.yml:
// resilience4j.ratelimiter.instances.smsApi:
//   limitForPeriod: 100         # max 100 calls
//   limitRefreshPeriod: 1s      # per second
//   timeoutDuration: 0          # don't wait, fail fast
```
**A:** Rate limiter caps throughput to external APIs to avoid exceeding quotas. `limitForPeriod=100, limitRefreshPeriod=1s` = max 100 requests/second. Requests exceeding the limit throw `RequestNotPermitted`.

---

### 33. @TimeLimiter — Timeouts
**Q: What happens after the timeout?**
```java
@Service
class SlowService {
    @TimeLimiter(name = "slowOp", fallbackMethod = "timeoutFallback")
    public CompletableFuture<String> slowOperation() {
        return CompletableFuture.supplyAsync(() -> {
            Thread.sleep_unchecked(5000); // 5 second operation
            return "result";
        });
    }

    public CompletableFuture<String> timeoutFallback(TimeoutException e) {
        return CompletableFuture.completedFuture("timeout-fallback");
    }
}

// resilience4j.timelimiter.instances.slowOp:
//   timeoutDuration: 2s       # cancel after 2 seconds
//   cancelRunningFuture: true # interrupt the future
```
**A:** `@TimeLimiter` cancels operations that exceed the timeout. `cancelRunningFuture=true` interrupts the underlying thread. Always combine `@TimeLimiter` with `@CircuitBreaker` — slow operations that always timeout should trip the circuit.

---

### 34. Combining Resilience4j Annotations
**Q: What is the correct order?**
```java
@Service
class ExternalService {
    @TimeLimiter(name = "ext")         // 1st: wrap with timeout
    @CircuitBreaker(name = "ext")      // 2nd: then circuit breaker
    @Retry(name = "ext")               // 3rd: then retry
    @Bulkhead(name = "ext")            // 4th: then bulkhead (innermost)
    public CompletableFuture<String> call() {
        return CompletableFuture.supplyAsync(() -> http.get("https://api.example.com"));
    }
}
// Execution order (outermost to innermost):
// TimeLimiter → CircuitBreaker → Retry → Bulkhead → actual call
```
**A:** Resilience4j annotations are applied as AOP decorators. Recommended order: TimeLimiter (outer) > CircuitBreaker > Retry > Bulkhead (inner). This means: retry before circuit breaker counts failures, circuit opens only after retries exhausted.

---

### 35. Resilience Event Publishing
**Q: How do you monitor circuit breaker state changes?**
```java
@Component
class CircuitBreakerMonitor {
    @Autowired CircuitBreakerRegistry registry;

    @PostConstruct
    void subscribeEvents() {
        registry.circuitBreaker("paymentGateway")
            .getEventPublisher()
            .onStateTransition(event ->
                log.warn("CB state: {} → {}",
                    event.getStateTransition().getFromState(),
                    event.getStateTransition().getToState()))
            .onFailureRateExceeded(event ->
                log.error("Failure rate: {}%", event.getEventType()))
            .onCallNotPermitted(event ->
                metrics.increment("circuit_breaker.rejected"));
    }
}
```
**A:** Resilience4j publishes events for state transitions, failures, and rejections. Wire these to your metrics system (Micrometer → Prometheus → Grafana) for operational visibility. Alert on state transitions to OPEN.

---

### 36. Fallback Hierarchy
**Q: What is the fallback strategy pattern?**
```java
@Service
class ProductService {
    @CircuitBreaker(name = "productDb", fallbackMethod = "fromCache")
    public Product getProduct(Long id) {
        return db.findById(id).orElseThrow(); // primary source
    }

    private Product fromCache(Long id, Exception e) {
        log.warn("DB unavailable, trying cache");
        return cache.get("product:" + id) // secondary source
            .orElseGet(() -> defaultProduct(id, e));
    }

    private Product defaultProduct(Long id, Exception e) {
        log.error("Both DB and cache unavailable for product {}", id);
        return new Product(id, "Unknown Product", BigDecimal.ZERO); // safe default
    }
}
```
**A:** Implement a fallback hierarchy: primary (DB) → secondary (cache) → last resort (safe default). Each fallback degrades gracefully rather than failing completely. Always log which level was used.

---

### 37. Timeout Configuration — Connection vs Read
**Q: What are the different timeout types?**
```java
// Spring RestTemplate timeouts:
@Bean
public RestTemplate restTemplate() {
    SimpleClientHttpRequestFactory factory = new SimpleClientHttpRequestFactory();
    factory.setConnectTimeout(3000); // time to establish TCP connection
    factory.setReadTimeout(5000);    // time to receive first byte of response
    return new RestTemplate(factory);
}

// WebClient timeouts:
@Bean
public WebClient webClient() {
    return WebClient.builder()
        .clientConnector(new ReactorClientHttpConnector(HttpClient.create()
            .option(ChannelOption.CONNECT_TIMEOUT_MILLIS, 3000)
            .responseTimeout(Duration.ofSeconds(5))
        ))
        .build();
}
```
**A:** Connect timeout: server not reachable. Read timeout: server connected but slow to respond. Both are essential. Without timeouts, a stuck downstream service holds threads/connections indefinitely.

---

### 38. Graceful Degradation — Feature Flags + Circuit Breaker
**Q: How do you combine feature flags with resilience?**
```java
@Service
class RecommendationService {
    @Autowired FeatureFlagClient flags;

    @CircuitBreaker(name = "mlService", fallbackMethod = "popularItems")
    public List<Item> getRecommendations(Long userId) {
        if (!flags.isEnabled("ml-recommendations")) {
            return popularItems(userId, null); // feature disabled
        }
        return mlService.predict(userId);
    }

    public List<Item> popularItems(Long userId, Throwable t) {
        return itemRepo.findTopByOrderBySalesDesc(PageRequest.of(0, 10));
    }
}
```
**A:** Feature flags + circuit breaker = maximum resilience. Can disable ML recommendations without code deployment (feature flag). Circuit breaker provides automatic protection if ML service degrades.

---

## Section 4: Distributed Patterns — Saga, Outbox, CQRS (Q39–Q47)

### 39. Saga Pattern — Distributed Transactions
**Q: What problem does Saga solve?**
```
Problem: Microservices can't use database-level ACID transactions across services

Saga = sequence of local transactions, each with a compensating transaction

Order Placement Saga:
  1. Order Service:    Create Order (local TX)
                       ✓ → Publish "OrderCreated" event
  2. Payment Service: Reserve Payment (local TX)
                       ✓ → Publish "PaymentReserved"
                       ✗ → Publish "PaymentFailed"
  3. Inventory Service: Reserve Stock (local TX)
                       ✓ → Publish "StockReserved"
                       ✗ → Publish "StockUnavailable"
  4. Shipping Service: Schedule Delivery

Compensating transactions on failure:
  StockUnavailable → PaymentService: refund → OrderService: cancel order
```
**A:** Saga achieves eventual consistency across services using events + compensation. Two implementations: **Choreography** (events only, no coordinator) and **Orchestration** (central orchestrator like Temporal or Spring State Machine).

---

### 40. Choreography Saga — Event-Driven
**Q: How is this different from Orchestration?**
```java
// Choreography: each service listens and reacts, no central coordinator

@KafkaListener(topics = "order-created")
@Transactional
public void onOrderCreated(OrderCreatedEvent event) {
    try {
        paymentService.reserve(event.getAmount(), event.getOrderId());
        eventPublisher.publish(new PaymentReservedEvent(event.getOrderId()));
    } catch (InsufficientFundsException e) {
        eventPublisher.publish(new PaymentFailedEvent(event.getOrderId(), e.getMessage()));
        // Order service listens to PaymentFailed and cancels
    }
}
```
**A:** Choreography: services react to events autonomously — no central brain. Pros: decoupled, simple. Cons: hard to trace flow, compensations are implicit. Orchestration: explicit coordinator knows the full flow — easier to monitor but adds central coupling.

---

### 41. Outbox Pattern — Guaranteed Message Delivery
**Q: What problem does the Outbox pattern solve?**
```java
@Service
class OrderService {
    @Transactional
    public Order createOrder(CreateOrderRequest req) {
        Order order = orderRepo.save(new Order(req));

        // PROBLEM: DB commit + Kafka send = two separate operations
        // If app crashes after DB commit but before Kafka send → message lost!

        // SOLUTION: Outbox pattern
        OutboxEvent event = new OutboxEvent("order-created", toJson(order));
        outboxRepo.save(event); // save event IN SAME transaction as order
        // Separate process (Debezium CDC or scheduled job) reads outbox and publishes to Kafka

        return order;
    }
}

// OutboxEvent table: id, topic, payload, published, created_at
// Debezium reads: new outbox rows → Kafka (guaranteed delivery)
```
**A:** The Outbox pattern: save event to a DB table within the same business transaction, then reliably publish with a separate process (CDC or polling). Guarantees no message is lost even if the app crashes between DB commit and Kafka send.

---

### 42. Transactional Outbox — Polling vs CDC
**Q: What are the two implementation approaches?**
```java
// Approach 1: Polling Publisher (simple, polling overhead)
@Scheduled(fixedDelay = 1000)
@Transactional
public void publishPendingEvents() {
    List<OutboxEvent> pending = outboxRepo.findByPublishedFalse();
    for (OutboxEvent event : pending) {
        kafkaTemplate.send(event.getTopic(), event.getPayload());
        event.setPublished(true);
        outboxRepo.save(event);
    }
}

// Approach 2: CDC (Change Data Capture) — zero latency, no polling
// Debezium captures INSERT on outbox table → sends to Kafka automatically
// outbox → kafka latency: milliseconds
// No polling overhead, no duplicate detection complexity
```
**A:** Polling is simple but has polling overhead and latency. CDC (Debezium) captures database change logs — near real-time with no polling. CDC is preferred in production but requires Debezium infrastructure.

---

### 43. CQRS — Command Query Responsibility Segregation
**Q: What does CQRS separate?**
```java
// Write Side (Command) — normalized, transactional
@Service class OrderCommandService {
    @Transactional
    public OrderId placeOrder(PlaceOrderCommand cmd) {
        Order order = new Order(cmd);
        orderRepo.save(order);
        eventStore.append(new OrderPlacedEvent(order));
        return order.getId();
    }
}

// Read Side (Query) — denormalized, optimized for query performance
@Service class OrderQueryService {
    // Separate read model: built from events, may be Elasticsearch/Redis
    public OrderDashboard getDashboard(Long userId) {
        return orderReadModelRepo.getDashboardForUser(userId);
        // One query, no joins, pre-computed aggregates
    }
}

// Update read model when events arrive:
@EventListener
public void on(OrderPlacedEvent event) {
    readModelUpdater.updateDashboard(event);
}
```
**A:** CQRS uses separate models for writes (normalized DB) and reads (denormalized read store). Benefits: read model optimized independently (e.g., Elasticsearch for full-text search, Redis for fast lookups). Cost: eventual consistency between write and read models.

---

### 44. Event Sourcing — State from Events
**Q: How is this different from storing current state?**
```java
// Traditional: store current state
// DB row: { id=1, status=SHIPPED, address="123 Main St", total=99.99 }

// Event Sourcing: store events
// Event 1: OrderPlaced    { orderId=1, total=99.99 }
// Event 2: AddressUpdated { orderId=1, address="123 Main St" }
// Event 3: PaymentCharged { orderId=1, amount=99.99 }
// Event 4: OrderShipped   { orderId=1, trackingId="ABC123" }

// Reconstruct state by replaying events:
@Service class OrderService {
    Order loadOrder(Long id) {
        List<Event> events = eventStore.loadEvents("Order", id);
        Order order = new Order();
        events.forEach(order::apply); // fold events into state
        return order;
    }
}
```
**A:** Event Sourcing: source of truth is the event log, not current state. Benefits: full audit trail, time travel (replay to any point), derive new projections from history. Drawbacks: complex, query performance requires projections, events are immutable.

---

### 45. Two-Phase Commit vs Eventual Consistency
**Q: When do you use each?**
```
2PC (Two-Phase Commit):
  Coordinator → "Prepare": all participants vote Yes/No
  Coordinator → "Commit" (if all Yes) or "Rollback"
  ✓ Strong consistency
  ✗ Slow, blocking, coordinator SPOF
  ✗ Not suitable for microservices across network partitions

Eventual Consistency (Saga/Outbox):
  Each service commits locally, propagates changes via events
  System is inconsistent for a short window, then converges
  ✓ High availability, partition tolerant
  ✓ Each service owns its data
  ✗ Complex compensation logic
  ✗ No guaranteed atomicity across services

Recommendation: prefer eventual consistency in microservices.
Use 2PC only within a single database engine (XA transactions).
```
**A:** CAP theorem: under network partition, choose consistency OR availability. Microservices choose availability + eventual consistency (AP). 2PC is appropriate within a single data center, single DB engine.

---

### 46. Idempotency Key — Safe Retries
**Q: How does an idempotency key work?**
```java
@RestController
class PaymentController {
    @PostMapping("/payments")
    public ResponseEntity<PaymentResult> charge(
        @RequestHeader("Idempotency-Key") String idempotencyKey,
        @RequestBody PaymentRequest req) {

        // Check if this request was already processed
        Optional<PaymentResult> existing = paymentStore.findByKey(idempotencyKey);
        if (existing.isPresent()) {
            return ResponseEntity.ok(existing.get()); // return cached result
        }

        PaymentResult result = paymentService.charge(req);
        paymentStore.save(idempotencyKey, result, Duration.ofHours(24));
        return ResponseEntity.ok(result);
    }
}
// Client generates UUID per request attempt, retries with same UUID on failure
```
**A:** Idempotency keys make POST/PUT operations safe to retry — the server returns the same result for duplicate requests. Essential for financial operations. Stripe, PayPal, and Twilio all use this pattern.

---

### 47. API Gateway Pattern
**Q: What does an API Gateway provide?**
```
Client → API Gateway → Microservices

API Gateway handles:
  ├── Routing: /api/orders → Order Service, /api/users → User Service
  ├── Authentication: validate JWT before forwarding
  ├── Rate Limiting: 100 req/s per API key
  ├── Request Aggregation: combine multiple service calls
  ├── Protocol Translation: REST → gRPC
  ├── Circuit Breaker: per upstream service
  ├── Logging & Tracing: inject correlation ID
  └── Load Balancing: across service instances

Spring Cloud Gateway:
@Bean
RouteLocator routes(RouteLocatorBuilder b) {
    return b.routes()
        .route("order-service", r -> r.path("/api/orders/**")
            .filters(f -> f.rewritePath("/api/orders/(?<segment>.*)", "/orders/${segment}")
                            .circuitBreaker(c -> c.setName("orderCB")))
            .uri("lb://order-service")) // load balanced discovery
        .build();
}
```
**A:** API Gateway is the single entry point for clients. Spring Cloud Gateway is reactive (WebFlux-based). Alternatives: Kong, Nginx, AWS API Gateway. Avoid "God Gateway" — keep routing/cross-cutting concerns only.

---

## Section 5: Observability — Tracing, Metrics, Service Discovery (Q48–Q55)

### 48. MDC — Correlation IDs for Logging
**Q: What does MDC provide?**
```java
import org.slf4j.MDC;

// In Filter or Interceptor — early in the request chain:
@Component
class CorrelationIdFilter extends OncePerRequestFilter {
    protected void doFilterInternal(HttpServletRequest req, HttpServletResponse res, FilterChain chain)
        throws ServletException, IOException {
        String correlationId = Optional.ofNullable(req.getHeader("X-Correlation-Id"))
            .orElseGet(() -> UUID.randomUUID().toString());
        MDC.put("correlationId", correlationId);
        res.addHeader("X-Correlation-Id", correlationId);
        try {
            chain.doFilter(req, res);
        } finally {
            MDC.clear(); // always clean up!
        }
    }
}

// logback.xml:
// <pattern>%d{HH:mm:ss} [%X{correlationId}] %-5level %logger{50} - %msg%n</pattern>
// Now every log line has the correlation ID automatically!
```
**A:** MDC (Mapped Diagnostic Context) attaches key-value pairs to the logging thread. All logs within a request automatically include `correlationId`. Enables grepping all logs for a single request across multiple log files.

---

### 49. Micrometer — Application Metrics
**Q: How do you record custom metrics?**
```java
import io.micrometer.core.instrument.*;

@Service
class OrderService {
    private final Counter ordersCreated;
    private final Timer  orderProcessingTime;
    private final Gauge  pendingOrders;

    OrderService(MeterRegistry registry) {
        this.ordersCreated = Counter.builder("orders.created")
            .tag("region", "us-east")
            .register(registry);

        this.orderProcessingTime = Timer.builder("orders.processing.time")
            .description("Time to process an order")
            .register(registry);

        this.pendingOrders = Gauge.builder("orders.pending",
            orderRepo, OrderRepository::countPending)
            .register(registry);
    }

    public Order create(CreateOrderRequest req) {
        return orderProcessingTime.record(() -> {
            Order o = processOrder(req);
            ordersCreated.increment();
            return o;
        });
    }
}
```
**A:** Micrometer provides a vendor-neutral metrics API. `Counter` (monotonically increasing), `Timer` (duration + count), `Gauge` (current value). Expose to Prometheus → Grafana dashboards. Spring Boot auto-configures JVM, HTTP, DB pool metrics.

---

### 50. Spring Boot Actuator — Health Checks
**Q: How do you add a custom health check?**
```java
import org.springframework.boot.actuate.health.*;

@Component
class ExternalServiceHealthIndicator implements HealthIndicator {
    @Autowired ExternalServiceClient client;

    public Health health() {
        try {
            client.ping(); // check connectivity
            return Health.up()
                .withDetail("service", "external-api")
                .withDetail("status", "reachable")
                .build();
        } catch (Exception e) {
            return Health.down()
                .withDetail("error", e.getMessage())
                .build();
        }
    }
}
// GET /actuator/health → {"status":"DOWN","components":{"externalServiceHealthIndicator":{"status":"DOWN","details":{"error":"..."}}}
// Used by Kubernetes readiness probes
```
**A:** Spring Boot aggregates all `HealthIndicator` beans. Kubernetes readiness probe calls `/actuator/health/readiness` — if DOWN, Kubernetes removes the pod from load balancer. Liveness probe calls `/actuator/health/liveness` — if DOWN, pod is restarted.

---

### 51. Distributed Tracing — Micrometer Tracing
**Q: What does a trace look like?**
```
Single HTTP request → Order Service → DB → Kafka → Inventory Service → DB

Trace: traceId=abc123
  Span 1: HTTP POST /orders (duration: 250ms) — Order Service
    Span 2: INSERT INTO orders (duration: 30ms) — DB
    Span 3: Kafka send "orders" (duration: 5ms)
  Span 4: @KafkaListener (duration: 120ms) — Inventory Service
    Span 5: SELECT FROM inventory (duration: 20ms) — DB

// Visualize in Zipkin or Jaeger
// Spring Boot auto-instruments: HTTP, Kafka, DB, RestTemplate, WebClient
```
```java
// Auto trace propagation with Spring Cloud Sleuth/Micrometer:
@Service
class OrderService {
    @Autowired Tracer tracer; // Micrometer Tracing

    public Order process(Long id) {
        Span span = tracer.nextSpan().name("business-validation").start();
        try (Tracer.SpanInScope ws = tracer.withSpan(span)) {
            validateOrder(id);
            return processOrder(id);
        } finally { span.end(); }
    }
}
```
**A:** Micrometer Tracing (replaces Spring Cloud Sleuth) auto-propagates trace context via HTTP headers (B3, W3C TraceContext) and Kafka headers. Each span is a unit of work. Traces visualize end-to-end request flows across services.

---

### 52. Service Discovery — Eureka & Spring Cloud
**Q: How does a service discover another service?**
```java
// application.yml (Service A):
// spring.application.name: order-service
// eureka.client.serviceUrl.defaultZone: http://eureka:8761/eureka

// Service B calls Service A:
@Service
class InventoryClient {
    @Autowired WebClient.Builder builder;

    public int getStock(Long productId) {
        return builder.build()
            .get()
            .uri("http://order-service/api/orders/{id}", productId) // logical name!
            .retrieve()
            .bodyToMono(Integer.class)
            .block();
    }
}
// spring.cloud.loadbalancer resolves "order-service" → actual IP:port via Eureka
```
**A:** Service Discovery removes hardcoded IPs. Services register with Eureka on startup. Clients use logical names; Spring Cloud LoadBalancer resolves to actual instances and load balances. Alternative: Kubernetes Service DNS (no Eureka needed in k8s).

---

### 53. Spring Cloud Config — Centralized Configuration
**Q: What does Config Server provide?**
```yaml
# Config Server serves configs from Git/Vault:
# http://config-server/order-service/production → application.yml from Git

# Client (Order Service):
spring:
  config:
    import: "configserver:http://config-server:8888"
  application:
    name: order-service
  profiles:
    active: production

# Config changes: POST /actuator/refresh on client → reloads @RefreshScope beans without restart
```
```java
@RefreshScope
@Configuration
class DynamicConfig {
    @Value("${feature.max-order-size:100}")
    private int maxOrderSize;
}
```
**A:** Config Server centralizes configuration management. All microservices fetch their config from one place. Git backend provides versioning + audit trail. `@RefreshScope` + `POST /actuator/refresh` enables hot-reload of config without restart.

---

### 54. Reactive Microservices — WebFlux
**Q: What is the output of this reactive chain?**
```java
import reactor.core.publisher.*;
import org.springframework.web.reactive.function.client.*;

@Service
class ProductService {
    @Autowired WebClient webClient;

    public Mono<ProductDetails> getDetails(Long id) {
        Mono<Product> productMono = webClient.get()
            .uri("/products/{id}", id)
            .retrieve()
            .bodyToMono(Product.class);

        Mono<Review[]> reviewsMono = webClient.get()
            .uri("/products/{id}/reviews", id)
            .retrieve()
            .bodyToMono(Review[].class);

        // Parallel execution:
        return Mono.zip(productMono, reviewsMono)
            .map(tuple -> new ProductDetails(tuple.getT1(), tuple.getT2()));
    }
}
```
**A:** `Mono.zip()` subscribes to both `productMono` and `reviewsMono` in parallel — waits for both. Total latency = max(product latency, reviews latency) instead of sum. WebFlux handles I/O non-blocking with one thread per core (vs one thread per request in Tomcat).

---

### 55. Structured Logging — JSON Logs for ELK
**Q: How do you configure structured logging?**
```java
// logback-spring.xml with logstash-logback-encoder:
// <encoder class="net.logstash.logback.encoder.LogstashEncoder"/>

// All logs output as JSON:
// {
//   "timestamp": "2024-01-15T10:30:00.000Z",
//   "level": "INFO",
//   "traceId": "abc123",
//   "spanId": "def456",
//   "service": "order-service",
//   "message": "Order created",
//   "orderId": 42,
//   "userId": 100
// }

@Service
class OrderService {
    @Autowired KafkaTemplate<String, OrderEvent> kafkaTemplate;

    public void create(OrderRequest req) {
        Order order = repo.save(new Order(req));
        log.info("Order created", // SLF4J structured logging (Logback 1.3+)
            StructuredArguments.keyValue("orderId", order.getId()),
            StructuredArguments.keyValue("userId", req.getUserId()));
    }
}
```
**A:** JSON logs are machine-parseable — Logstash/Filebeat ships to Elasticsearch, Kibana queries by field. MDC values (`traceId`, `spanId`) appear in every log line automatically when using `LogstashEncoder`. Always use structured logging in production microservices.

---

## Section 6: How to Explain in Interview (Spoken Style Format)

### General Interview Tips for Microservices & Kafka

**Interviewer:** "Explain Kafka's core concepts like topics, partitions, and offsets."

**Your Response:** "Certainly! In Kafka, a **topic** is like a table name - it's a named stream of records. Each topic is split into **partitions**, which are ordered, immutable logs. Partitions enable parallelism - more partitions mean more consumers can process in parallel. Each message within a partition has an **offset** - a unique, monotonically increasing number. Consumers track their position by storing the last offset they've processed. This design allows Kafka to scale horizontally and provide replayability - consumers can rewind to any offset and reprocess messages."

**Interviewer:** "How does Kafka guarantee message ordering?"

**Your Response:** "Kafka guarantees ordering **within a partition** only. All messages with the same key go to the same partition using `hash(key) % numPartitions`. So if you send orders with key `orderId`, all updates for that order are processed in order. However, there's no guarantee across different partitions. For global ordering, you'd need a single-partition topic, but that kills throughput. In practice, we design around this - we care about ordering per entity, not globally."

**Interviewer:** "What's the difference between at-least-once and exactly-once delivery?"

**Your Response:** "**At-least-once** means messages might be duplicated but never lost. We commit offsets only after successful processing. If the consumer crashes after processing but before committing, the message gets redelivered. **Exactly-once** uses Kafka transactions - the producer sends messages and commits consumer offsets atomically. Consumers set `isolation.level=read_committed` to see only committed messages. Exactly-once is complex and has performance overhead, so many systems use at-least-once with idempotent consumers."

**Interviewer:** "How do you handle consumer lag in production?"

**Your Response:** "Consumer lag is when consumers can't keep up with producers. I monitor it using `kafka-consumer-groups.sh --describe` or Prometheus metrics. High lag means consumers are too slow or there's a throughput spike. Solutions include: adding more consumers (up to partition count), optimizing processing logic, or increasing partitions. For critical systems, I set up alerts when lag exceeds a threshold. Sometimes we use batch processing or async processing to speed up consumption."

**Interviewer:** "Explain the Circuit Breaker pattern and when to use it."

**Your Response:** "Circuit Breaker prevents cascading failures. It wraps calls to downstream services and tracks failures. When failure rate exceeds a threshold, it 'trips' and fails fast - no more calls to the struggling service. After a timeout, it enters half-open state and allows a few test calls. If those succeed, it closes again; if not, it stays open. I use Resilience4j in Spring Boot with `@CircuitBreaker`. It's essential for microservices where one service failure shouldn't bring down the entire system."

**Interviewer:** "What's the Saga pattern for distributed transactions?"

**Your Response:** "Saga breaks a distributed transaction into a series of local transactions with compensating actions. Instead of one atomic ACID transaction across services, each service performs its local transaction and publishes an event. The next service executes its transaction based on that event. If any step fails, we run compensating transactions in reverse order to undo previous work. For example, in an order flow: create order → reserve inventory → process payment. If payment fails, we release inventory and cancel the order. It's eventual consistency but avoids distributed deadlocks."

**Interviewer:** "How do you implement the Outbox pattern?"

**Your Response:** "The Outbox pattern solves the dual-write problem when updating a database and publishing to Kafka. Instead of writing to both, we write to an 'outbox' table in the same database transaction. A separate connector process reads the outbox table and publishes to Kafka. This guarantees atomicity - either both the DB update and message happen, or neither. Debezium CDC connector is perfect for this - it captures outbox changes as Kafka messages automatically."

**Interviewer:** "What are dead letter topics and when do you use them?"

**Your Response:** "Dead Letter Topics (DLT) handle 'poison pill' messages that repeatedly fail processing. After max retries, instead of losing the message or blocking the consumer, we publish it to a separate DLT topic. The operations team can inspect, fix, and replay these messages. In Spring Kafka, I use `DeadLetterPublishingRecoverer` with `DefaultErrorHandler`. DLTs are crucial for production reliability - they prevent one bad message from stopping the entire consumer group."

**Interviewer:** "How do you ensure idempotent consumers?"

**Your Response:** "Idempotent consumers handle duplicates safely. The key is using a unique identifier from the message - like `orderId` or a dedicated `idempotencyKey`. Before processing, we check if we've already processed this key (using a database table or Redis). If yes, we skip processing. Alternatively, we use database unique constraints or `INSERT ... ON CONFLICT` for upserts. This is essential with at-least-once delivery where duplicates are guaranteed to happen."

**Interviewer:** "What's the role of Schema Registry in Kafka?"

**Your Response:** "Schema Registry manages Avro/Protobuf schemas and enforces compatibility. Producers register schemas; consumers fetch them by ID. This prevents breaking changes - you can't deploy a producer with incompatible schema. It also reduces message size since we send schema ID + binary data, not full JSON. For microservices, Schema Registry is crucial for governance - it ensures all services agree on data contracts and enables safe schema evolution."

---

*This interview format helps you articulate complex microservices and Kafka concepts clearly and concisely. Practice explaining these concepts out loud to build confidence for your actual interview.*

---

> 🔖 **Last read:** <!-- update here -->
