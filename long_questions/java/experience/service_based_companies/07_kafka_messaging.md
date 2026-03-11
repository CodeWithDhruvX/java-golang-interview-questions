# 📨 07 — Kafka Messaging
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Kafka architecture (brokers, topics, partitions, offsets)
- Producer and Consumer in Spring Boot (spring-kafka)
- Consumer groups and offset management
- Serialization / Deserialization
- Error handling and dead-letter topics
- `@KafkaListener` configuration

---

## ❓ Most Asked Questions

### Q1. What is Apache Kafka and its core concepts?

```text
KAFKA ARCHITECTURE:

Producer ──► Topic (partitioned) ──► Consumer Group
                  |
            Partition 0: [msg1, msg2, msg3] ← offset-based
            Partition 1: [msg4, msg5]
            Partition 2: [msg6]
                  |
              Brokers (Kafka nodes, e.g. 3 for HA)
                  |
            ZooKeeper / KRaft (metadata management)
```

| Concept | Description |
|---------|-------------|
| **Topic** | Category/feed of messages |
| **Partition** | Ordered, immutable sequence within a topic; enables parallelism |
| **Offset** | Unique ID of a message within a partition |
| **Producer** | Publishes messages to topics |
| **Consumer** | Reads messages from topics |
| **Consumer Group** | N consumers sharing topic partitions (each partition → 1 consumer in group) |
| **Broker** | Kafka server that stores messages |
| **Replication Factor** | Copies of each partition across brokers (3 = standard HA) |
| **Retention** | How long messages are kept (default 7 days) |

---

### 🎯 How to Explain in Interview

"Apache Kafka is a distributed streaming platform that acts as a highly scalable message broker. At its core, Kafka has topics which are like categories for messages, and each topic is split into partitions for parallelism. Messages within a partition have a strict order identified by offsets. Producers publish messages to topics, and consumers read from topics. The magic happens with consumer groups - each partition is consumed by only one consumer in a group, enabling parallel processing. Brokers are the Kafka servers that store the data, and replication factor ensures durability by copying partitions across multiple brokers. Kafka's design makes it perfect for real-time data pipelines, event sourcing, and microservices communication."

---

### Q2. How do you configure a Kafka Producer in Spring Boot?

```yaml
# application.yml
spring:
  kafka:
    bootstrap-servers: localhost:9092
    producer:
      key-serializer: org.apache.kafka.common.serialization.StringSerializer
      value-serializer: org.springframework.kafka.support.serializer.JsonSerializer
      acks: all           # wait for all replicas to acknowledge (strongest guarantee)
      retries: 3
      batch-size: 16384   # 16 KB batch before sending
      linger-ms: 5        # wait 5ms to accumulate batch
      enable-idempotence: true  # exactly-once delivery
```

```java
// Producer service
@Service
public class OrderEventProducer {

    private final KafkaTemplate<String, OrderEvent> kafkaTemplate;
    private static final String TOPIC = "orders";

    public OrderEventProducer(KafkaTemplate<String, OrderEvent> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    // Async send — fire and forget
    public void sendOrderEvent(OrderEvent event) {
        kafkaTemplate.send(TOPIC, event.getOrderId().toString(), event);
    }

    // Send with callback
    public void sendWithCallback(OrderEvent event) {
        CompletableFuture<SendResult<String, OrderEvent>> future =
            kafkaTemplate.send(TOPIC, event.getOrderId().toString(), event);

        future.whenComplete((result, ex) -> {
            if (ex == null) {
                RecordMetadata meta = result.getRecordMetadata();
                log.info("Sent to partition {} at offset {}",
                    meta.partition(), meta.offset());
            } else {
                log.error("Failed to send event: {}", ex.getMessage());
            }
        });
    }

    // Send to specific partition
    public void sendToPartition(OrderEvent event, int partition) {
        kafkaTemplate.send(new ProducerRecord<>(TOPIC, partition,
            event.getOrderId().toString(), event));
    }
}
```

---

### 🎯 How to Explain in Interview

"Configuring a Kafka producer in Spring Boot is straightforward with the right settings. I start with bootstrap-servers to connect to Kafka, then configure serializers for keys and values - typically StringSerializer for keys and JsonSerializer for objects. For reliability, I set acks to 'all' to wait for all replicas, enable retries, and use enable-idempotence for exactly-once delivery. I also tune performance with batch-size and linger-ms to accumulate messages before sending. In code, I use KafkaTemplate which Spring provides - it's much cleaner than using the raw Kafka producer. I can send messages synchronously or asynchronously with callbacks to handle success or failure scenarios. This approach gives me reliable message publishing with good performance."

---

### Q3. How do you configure a Kafka Consumer in Spring Boot?

```yaml
spring:
  kafka:
    consumer:
      group-id: order-service
      key-deserializer: org.apache.kafka.common.serialization.StringDeserializer
      value-deserializer: org.springframework.kafka.support.serializer.JsonDeserializer
      auto-offset-reset: earliest    # start from beginning if no offset stored
      enable-auto-commit: false       # manual commit for reliability
      max-poll-records: 100
    listener:
      ack-mode: MANUAL_IMMEDIATE     # commit after processing each record
      concurrency: 3                 # 3 listener threads (matches partition count)
```

```java
@Component
public class OrderEventConsumer {

    @KafkaListener(topics = "orders", groupId = "order-service",
                   containerFactory = "kafkaListenerContainerFactory")
    public void consumeOrderEvent(OrderEvent event,
                                  @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
                                  @Header(KafkaHeaders.OFFSET) long offset,
                                  Acknowledgment acknowledgment) {
        try {
            log.info("Received event: {} from partition {} offset {}", event, partition, offset);
            processOrder(event);
            acknowledgment.acknowledge();  // commit offset only after successful processing
        } catch (Exception e) {
            log.error("Error processing order event: {}", e.getMessage());
            // DO NOT acknowledge — message will be redelivered
            throw e;
        }
    }

    // Batch listener
    @KafkaListener(topics = "orders", groupId = "batch-service",
                   batch = "true")  // enable batch mode in config
    public void consumeBatch(List<ConsumerRecord<String, OrderEvent>> records,
                             Acknowledgment ack) {
        records.forEach(r -> processOrder(r.value()));
        ack.acknowledge();
    }
}
```

---

### 🎯 How to Explain in Interview

"For Kafka consumers in Spring Boot, I configure the consumer group ID, deserializers, and offset management strategy. I set auto-offset-reset to 'earliest' to start from the beginning if no offset exists, and disable auto-commit for manual control. The ack-mode set to MANUAL_IMMEDIATE means I commit offsets only after successfully processing each message. In the consumer code, I use @KafkaListener to subscribe to topics. I can access metadata like partition and offset through headers, and use the Acknowledgment parameter for manual commits. I also implement batch processing for better throughput - instead of processing one message at a time, I process multiple messages in a single poll. This gives me reliable message consumption with good control over processing semantics."

---

### Q4. What is a Dead Letter Topic (DLT)?

```java
// Configure error handler with DLT
@Configuration
public class KafkaErrorConfig {

    @Bean
    public DefaultErrorHandler errorHandler(KafkaTemplate<String, Object> template) {
        // Retry 3 times with 1 second backoff, then send to DLT
        DeadLetterPublishingRecoverer recoverer =
            new DeadLetterPublishingRecoverer(template,
                (record, ex) -> new TopicPartition(
                    record.topic() + ".DLT",  // e.g., "orders.DLT"
                    record.partition()
                ));

        ExponentialBackOffWithMaxRetries backoff = new ExponentialBackOffWithMaxRetries(3);
        backoff.setInitialInterval(1000);   // 1s initial
        backoff.setMultiplier(2.0);         // 1s, 2s, 4s

        return new DefaultErrorHandler(recoverer, backoff);
    }
}

// DLT Consumer — monitor and alert on dead letters
@KafkaListener(topics = "orders.DLT", groupId = "dlt-monitor")
public void consumeDeadLetter(ConsumerRecord<String, OrderEvent> record) {
    log.error("Dead letter received: key={}, reason={}",
        record.key(),
        record.headers().lastHeader("kafka_dlt-exception-message"));
    alertService.sendAlert("Order processing failed: " + record.value());
}
```

---

### 🎯 How to Explain in Interview

"Dead Letter Topics are essential for handling failed messages in Kafka. When a message can't be processed after multiple retries, instead of losing it or blocking the consumer, I route it to a special DLT topic. I configure this with a DefaultErrorHandler that retries with exponential backoff, then uses DeadLetterPublishingRecoverer to send the failed message to a topic like 'orders.DLT'. The DLT preserves the original message plus error information in headers. I then have a separate consumer monitoring the DLT to alert teams or trigger manual recovery processes. This approach ensures no messages are lost, provides visibility into processing failures, and keeps the main consumer flowing. It's a critical pattern for robust event-driven systems."

---

### Q5. How do Kafka partitions and consumer groups work?

```text
Topic: "orders" with 3 partitions

SCENARIO 1: 1 consumer group, 2 consumers, 3 partitions
Consumer A handles: Partition 0, Partition 1
Consumer B handles: Partition 2
→ Some consumers handle more partitions

SCENARIO 2: 1 consumer group, 3 consumers, 3 partitions  (IDEAL)
Consumer A handles: Partition 0
Consumer B handles: Partition 1
Consumer C handles: Partition 2
→ Parallel processing, maximum throughput

SCENARIO 3: 1 consumer group, 4 consumers, 3 partitions
Consumer A handles: Partition 0
Consumer B handles: Partition 1
Consumer C handles: Partition 2
Consumer D: IDLE — more consumers than partitions!

SCENARIO 4: 2 consumer groups reading same topic (e.g., order-service + audit-service)
Group "order-service":  ALL 3 partitions → processed for business logic
Group "audit-service":  ALL 3 partitions → processed for auditing
→ Both groups get ALL messages — independent consumption!
```

---

### 🎯 How to Explain in Interview

"Kafka partitions and consumer groups work together to enable both scalability and parallelism. Partitions divide a topic into parallel streams - each partition maintains order but can be consumed independently. Consumer groups enable load balancing - within a group, each partition is consumed by exactly one consumer. If I have 3 partitions and 3 consumers, each consumer handles one partition for maximum parallelism. If I have more consumers than partitions, some consumers sit idle. The beauty is that multiple consumer groups can read the same topic independently - an order service and an audit service can both process all messages without interfering with each other. This design allows me to scale horizontally by adding more partitions and consumers as needed."

---

### Q6. What is exactly-once semantics in Kafka?

```java
// At-most-once — may lose messages (auto-commit before processing)
// At-least-once — may duplicate messages (commit after processing, but can fail mid-way)
// Exactly-once — requires idempotent producer + transactional consumer

// Idempotent producer (enable-idempotence=true)
// Deduplicates retried messages per partition using sequence numbers

// Transactional producer for cross-partition exactly-once
@Bean
public ProducerFactory<String, Object> producerFactory() {
    Map<String, Object> config = new HashMap<>();
    config.put(ProducerConfig.TRANSACTIONAL_ID_CONFIG, "order-producer-1");
    config.put(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, true);
    // ...
    return new DefaultKafkaProducerFactory<>(config);
}

@Bean
public KafkaTemplate<String, Object> kafkaTemplate() {
    KafkaTemplate<String, Object> template = new KafkaTemplate<>(producerFactory());
    template.setTransactionIdPrefix("order-tx-");
    return template;
}

// Usage
@Transactional   // DB + Kafka in one transaction (Kafka + DB 2PC not supported — use Outbox pattern)
public void processAndPublish(Order order) {
    orderRepository.save(order);           // DB write
    kafkaTemplate.send("orders", order.getId().toString(), toEvent(order));  // Kafka write
}
```

---

### 🎯 How to Explain in Interview

"Exactly-once semantics in Kafka is about preventing both message loss and duplication. At-most-once might lose messages, at-least-once might duplicate them. Exactly-once requires both idempotent producers and transactional consumers. I enable idempotence on the producer to automatically deduplicate retried messages using sequence numbers. For true exactly-once across multiple operations, I use transactions with a transactional ID. However, there's a catch - Kafka and database transactions can't be natively coordinated. That's where patterns like the Outbox pattern come in. The key is understanding that exactly-once in Kafka means exactly-once delivery to Kafka, not exactly-once processing across multiple systems. For most use cases, idempotent producers plus proper error handling provide sufficient guarantees."

---

### Q7. What is the Kafka Outbox Pattern?

```text
PROBLEM: How to atomically save to DB AND publish to Kafka?
- If DB save succeeds but Kafka publish fails → inconsistency!
- If Kafka publish succeeds but DB rollback → duplicate event!

SOLUTION: Outbox Pattern
1. Save entity + outbox event in ONE DB transaction
2. Separate process (Debezium CDC or poller) reads outbox table
3. Publishes events to Kafka reliably
4. Marks events as published in DB

IMPLEMENTATION:
┌─────────────────┐     ┌────────────────┐     ┌─────────┐
│  OrderService   │──►  │  outbox_events │──►  │  Kafka  │
│  (DB transaction)│     │  (same DB)     │     │  Topic  │
└─────────────────┘     └────────────────┘     └─────────┘
                          Debezium/Poller reads
```

```java
// Outbox event entity
@Entity
@Table(name = "outbox_events")
public class OutboxEvent {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String aggregateType;   // "Order"
    private String aggregateId;     // order ID
    private String eventType;       // "OrderCreated"
    private String payload;         // JSON-serialized event
    private boolean published = false;

    @CreationTimestamp
    private LocalDateTime createdAt;
}

@Transactional
public void createOrder(CreateOrderRequest request) {
    Order order = orderRepository.save(new Order(request));

    // Same transaction — atomic!
    outboxEventRepository.save(new OutboxEvent(
        "Order", order.getId().toString(),
        "OrderCreated", objectMapper.writeValueAsString(order)
    ));
}
```

---

### 🎯 How to Explain in Interview

"The Outbox pattern solves the classic problem of atomically updating a database and publishing events to Kafka. The challenge is that if I save to the database but Kafka publishing fails, I have inconsistency. If I publish to Kafka but the database transaction rolls back, I get duplicate events. The Outbox pattern solves this by saving both the business entity and an outbox event in the same database transaction. Then a separate process - either a Debezium CDC connector or a polling application - reads the outbox table and reliably publishes events to Kafka. Once published, it marks the events as processed. This gives me true atomicity between database changes and event publishing, which is crucial for event-driven architectures. It's a bit more complex but provides the reliability guarantees needed for serious distributed systems."
