# 🏗️ Kafka — Spring Boot Integration

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Wipro, Cognizant, Capgemini, HCL

---

## Q1. How do you integrate Kafka with a Spring Boot application? Show the basic setup.

"Integrating Kafka into Spring Boot uses the **`spring-kafka`** library which provides auto-configuration and convenient abstractions over the native Kafka client.

**Step 1 — Add Dependency:**
```xml
<dependency>
    <groupId>org.springframework.kafka</groupId>
    <artifactId>spring-kafka</artifactId>
</dependency>
```

**Step 2 — `application.yml` Configuration:**
```yaml
spring:
  kafka:
    bootstrap-servers: localhost:9092
    consumer:
      group-id: my-app-group
      auto-offset-reset: earliest
      key-deserializer: org.apache.kafka.common.serialization.StringDeserializer
      value-deserializer: org.apache.kafka.common.serialization.StringDeserializer
    producer:
      key-serializer: org.apache.kafka.common.serialization.StringSerializer
      value-serializer: org.apache.kafka.common.serialization.StringSerializer
```

**Step 3 — Producer using `KafkaTemplate`:**
```java
@Service
public class OrderProducer {
    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    public void sendOrder(String orderId, String message) {
        kafkaTemplate.send("orders-topic", orderId, message);
    }
}
```

**Step 4 — Consumer using `@KafkaListener`:**
```java
@Service
public class OrderConsumer {
    @KafkaListener(topics = "orders-topic", groupId = "my-app-group")
    public void consumeOrder(String message) {
        System.out.println("Received order: " + message);
    }
}
```"

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Capgemini — the most commonly asked practical implementation question in project rounds where the candidate is expected to set up end-to-end producer/consumer in Java.

#### Indepth
**`auto-offset-reset: earliest` vs. `latest`:** `earliest` means the consumer starts reading from the very beginning of the topic if no committed offset exists (useful for new deployments). `latest` means it only reads new messages published AFTER the consumer started. Choosing the wrong one is a frequent production bug seen in freshly deployed services that miss backlogged messages.

---

## Q2. How do you consume Kafka messages as Java Objects (not plain Strings)?

"In real projects, you produce and consume JSON payloads mapped to Java POJOs. Spring Kafka supports this via `JsonSerializer` and `JsonDeserializer`.

**Define the POJO:**
```java
public class OrderEvent {
    private String orderId;
    private String userId;
    private double amount;
    // getters, setters, no-args constructor required by Jackson
}
```

**Producer Config:**
```yaml
spring:
  kafka:
    producer:
      value-serializer: org.springframework.kafka.support.serializer.JsonSerializer
```

**Consumer Config:**
```yaml
spring:
  kafka:
    consumer:
      value-deserializer: org.springframework.kafka.support.serializer.JsonDeserializer
      properties:
        spring.json.trusted.packages: "com.myapp.events"
```

**Consumer Listener:**
```java
@KafkaListener(topics = "orders-topic", groupId = "my-app-group")
public void consume(OrderEvent order) {
    System.out.println("Order received for user: " + order.getUserId());
    // process the order...
}
```

The `spring.json.trusted.packages` property is mandatory for security — it whitelists the packages whose classes are allowed to be deserialized from JSON, preventing deserialization attacks."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Wipro, HCL, Cognizant — practical implementation widely asked in L2 interviews when the project involves Kafka with a Spring Boot microservice architecture.

#### Indepth
**`ConsumerRecord` for Full Metadata:** Instead of consuming just the value, inject `ConsumerRecord<String, OrderEvent>`. This gives access to the message's full metadata: `record.topic()`, `record.partition()`, `record.offset()`, `record.timestamp()`, `record.headers()` — essential for debugging and idempotency checks.

---

## Q3. How do you handle errors and retries in Spring Kafka consumers?

"Without error handling, any exception thrown inside a `@KafkaListener` will cause an infinite retry loop — the consumer fetches the same bad message forever, blocking the entire partition.

**Spring Kafka provides `DefaultErrorHandler` for retry logic:**

```java
@Bean
public DefaultErrorHandler errorHandler() {
    // Retry up to 3 times with 1-second fixed backoff
    FixedBackOff fixedBackOff = new FixedBackOff(1000L, 3);
    DefaultErrorHandler errorHandler = new DefaultErrorHandler(fixedBackOff);

    // Do NOT retry for these specific non-retryable exceptions
    errorHandler.addNotRetryableExceptions(
        IllegalArgumentException.class,
        NullPointerException.class
    );
    return errorHandler;
}
```

**Register it on the listener container factory:**
```java
@Bean
public ConcurrentKafkaListenerContainerFactory<String, String> kafkaListenerContainerFactory(
        ConsumerFactory<String, String> consumerFactory,
        DefaultErrorHandler errorHandler) {
    ConcurrentKafkaListenerContainerFactory<String, String> factory =
        new ConcurrentKafkaListenerContainerFactory<>();
    factory.setConsumerFactory(consumerFactory);
    factory.setCommonErrorHandler(errorHandler);
    return factory;
}
```

After exhausting retries, the message is logged and skipped (or sent to a DLQ if configured). This prevents partition blocking while maintaining visibility into failures."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Infosys, Wipro — a very commonly asked operational question in rounds where the interviewer wants to see that the candidate understands that naive listeners break under bad data conditions.

#### Indepth
**Manual Acknowledgement Mode:** For fine-grained control, set `ackMode = MANUAL` on the listener container factory. The `@KafkaListener` method then receives an `Acknowledgment` parameter. You call `acknowledgment.acknowledge()` only after successful processing. If an exception occurs, the offset is never committed, and the message will be redelivered — giving you explicit at-least-once control.

---
