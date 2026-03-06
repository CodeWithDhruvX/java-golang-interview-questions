# 🏗️ Kafka — Spring Boot & Golang Integration

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Wipro, Cognizant, Capgemini, HCL

---

## Q1. How do you integrate Kafka with a Spring Boot or Golang application? Show the basic setup.

"Integrating Kafka depends on the target language ecosystem:

**Java Spring Boot — Step 1: Add Dependency:**
```xml
<dependency>
    <groupId>org.springframework.kafka</groupId>
    <artifactId>spring-kafka</artifactId>
</dependency>
```

**Java Spring Boot — Step 2: `application.yml` Configuration:**
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

**Java Spring Boot — Step 3: Producer & Consumer:**
```java
@Service
public class OrderService {
    @Autowired private KafkaTemplate<String, String> kafkaTemplate;

    public void send(String id, String msg) { kafkaTemplate.send("orders", id, msg); }

    @KafkaListener(topics = "orders", groupId = "my-app-group")
    public void consume(String message) { System.out.println("Received: " + message); }
}
```

**Golang — Step 1: Install Segmentio/kafka-go:**
```bash
go get github.com/segmentio/kafka-go
```

**Golang — Step 2: Producer & Consumer:**
```go
package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func produce() {
	writer := &kafka.Writer{Addr: kafka.TCP("localhost:9092"), Topic: "orders"}
	writer.WriteMessages(context.Background(), kafka.Message{Key: []byte("1"), Value: []byte("msg")})
}

func consume() {
	reader := kafka.NewReader(kafka.ReaderConfig{Brokers: []string{"localhost:9092"}, Topic: "orders", GroupID: "my-app-group"})
	m, _ := reader.ReadMessage(context.Background())
	fmt.Printf("Received: %s\n", string(m.Value))
}
```"

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys, Capgemini — the most commonly asked practical implementation question in project rounds where the candidate is expected to set up end-to-end producer/consumer in their stated primary language.

#### Indepth
**`auto-offset-reset: earliest` vs. `latest`:** `earliest` means the consumer starts reading from the very beginning of the topic if no committed offset exists (useful for new deployments). `latest` means it only reads new messages published AFTER the consumer started. Choosing the wrong one is a frequent production bug seen in freshly deployed services that miss backlogged messages.

---

## Q2. How do you consume Kafka messages as structured Objects (not plain Strings)?

"In real projects, you produce and consume JSON payloads mapped to native Structs/POJOs.

**Java Spring Boot:**
Spring Kafka supports this via `JsonSerializer` and `JsonDeserializer`.
```yaml
spring:
  kafka:
    producer:
      value-serializer: org.springframework.kafka.support.serializer.JsonSerializer
    consumer:
      value-deserializer: org.springframework.kafka.support.serializer.JsonDeserializer
      properties:
        spring.json.trusted.packages: "com.myapp.events"
```

```java
@KafkaListener(topics = "orders")
public void consume(OrderEvent order) {
    System.out.println("User: " + order.getUserId());
}
```
The `spring.json.trusted.packages` property is mandatory for security — it whitelists the packages whose classes are allowed to be deserialized from JSON.

**Golang:**
In Go, the Kafka client generally yields raw byte arrays `[]byte`. You manually unmarshal the JSON bytes into a native Go struct using `encoding/json`.
```go
type OrderEvent struct {
    OrderId string  `json:"orderId"`
    UserId  string  `json:"userId"`
    Amount  float64 `json:"amount"`
}

func consumeObject() {
    reader := kafka.NewReader(...)
    msg, _ := reader.ReadMessage(context.Background())
    
    var order OrderEvent
    json.Unmarshal(msg.Value, &order)
    fmt.Printf("User: %s\n", order.UserId)
}
```"

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Wipro, HCL, Cognizant — practical implementation widely asked in L2 interviews when the project involves Kafka with microservices.

#### Indepth
**Full Metadata access:** Instead of consuming just the mapped value, you often need the `ConsumerRecord` (Java) or `kafka.Message` (Go). This gives access to the message's full metadata: topic, partition, offset, timestamp, headers — essential for debugging and idempotency checks.

---

## Q3. How do you handle errors and retries in Kafka consumers?

"Without error handling, any exception thrown while processing will generally cause an infinite retry loop or premature offset commits.

**Java Spring Boot — `DefaultErrorHandler`:**
```java
@Bean
public DefaultErrorHandler errorHandler() {
    FixedBackOff fixedBackOff = new FixedBackOff(1000L, 3); // Retry 3 times, 1s apart
    DefaultErrorHandler errorHandler = new DefaultErrorHandler(fixedBackOff);
    errorHandler.addNotRetryableExceptions(IllegalArgumentException.class);
    return errorHandler;
}
```
Register this on the `ConcurrentKafkaListenerContainerFactory`. After exhausting retries, the message is skipped (or sent to a DLQ) preventing partition blocking.

**Golang — Custom Retry Loop:**
Golang's `kafka-go` doesn't provide automatic retry handling in the consumer loop. You must implement exponential backoff manually on the processing logic before invoking `CommitMessages()`.
```go
msg, _ := reader.FetchMessage(context.Background())
retries := 3
for i := 0; i < retries; i++ {
    err := process(msg)
    if err == nil {
        reader.CommitMessages(context.Background(), msg)
        break
    }
    time.Sleep(time.Duration(i*100) * time.Millisecond) // backoff
}
```"

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Infosys, Wipro — a very commonly asked operational question in rounds where the interviewer wants to see that the candidate understands that naive listeners break under bad data conditions.

#### Indepth
**Explicit Acknowledgement:** By default, frameworks might auto-commit. For fine-grained control, set manual commit modes (`ackMode = MANUAL` in Spring, or using `FetchMessage` + explicit `CommitMessages` in Go). You acknowledge only after successful processing. If an error occurs and retries fail, the offset is never committed over the failure. 
---
