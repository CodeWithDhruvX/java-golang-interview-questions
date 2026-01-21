## 🔹 Section 8: Reactive, Kafka & Messaging (241-260)

### Question 241: What is Spring WebFlux and how does it differ from Spring MVC?

**Answer:**
(See Q100). Non-blocking stack vs Blocking stack.

---

### Question 242: What are Mono and Flux in WebFlux?

**Answer:**
Reactive Streams publishers (Project Reactor).
*   **Mono<T>:** 0 or 1 item. (Like `Optional` + `Future`).
*   **Flux<T>:** 0 to N items. (Like a Stream).

---

### Question 243: How do you implement reactive REST endpoints in Spring Boot?

**Answer:**
Return `Mono<T>` or `Flux<T>` from controller.
```java
@GetMapping("/stream")
public Flux<String> stream() {
    return Flux.interval(Duration.ofSeconds(1)).map(i -> "Msg " + i);
}
```

---

### Question 244: How to handle backpressure in WebFlux?

**Answer:**
Handled by the subscriber.
`onBackpressureBuffer()` or `onBackpressureDrop()`.
Allows the consumer to signal "I can only handle 10 items now", preventing OOM errors if producer is too fast.

---

### Question 245: What is Server-Sent Events (SSE) and how to implement it in Spring Boot?

**Answer:**
One-way streaming channel from Server -> Client.
Return `Flux<String>` with `text/event-stream` media type.
Great for notifications without WebSockets.

---

### Question 246: How to connect Spring Boot with Kafka for message publishing?

**Answer:**
Dependency: `spring-kafka`.
Inject `KafkaTemplate<String, String>`.
`template.send("topic", "key", "payload")`.

---

### Question 247: How to consume Kafka messages in Spring Boot?

**Answer:**
Method annotation:
```java
@KafkaListener(topics = "my-topic", groupId = "group_id")
public void listen(String message) { ... }
```

---

### Question 248: How to handle Kafka error handling and retries?

**Answer:**
Config `DefaultErrorHandler` with `FixedBackOff`.
If retry exhausted, send to Dead Letter Topic (DLT) using `DeadLetterPublishingRecoverer`.

---

### Question 249: How to test Kafka consumers in Spring Boot?

**Answer:**
`@EmbeddedKafka`.
Starts an in-memory Kafka broker for the test class.

---

### Question 250: How does Spring Boot support RabbitMQ integration?

**Answer:**
Dependencies `spring-boot-starter-amqp`.
Auto-configures `ConnectionFactory`.
Use `RabbitTemplate` for sending.
Use `@RabbitListener` for receiving.

---

### Question 251: What are cold and hot publishers in WebFlux?

**Answer:**
*   **Cold:** Data generation starts **only** when subscribed. Each subscriber gets their own fresh stream (e.g., HTTP Call).
*   **Hot:** Data flows regardless of subscribers. Late subscribers miss old data (e.g., Mouse Events, Ticker).

---

### Question 252: How does `WebClient` compare to `RestTemplate`?

**Answer:**
(See Q60). Functional, fluent API. Supports Streaming.

---

### Question 253: How to return streaming JSON from WebFlux endpoint?

**Answer:**
`MediaType.APPLICATION_NDJSON` (Newline Delimited JSON).
Allows client to parse array items one by one as they arrive.

---

### Question 254: How do you implement reactive backpressure with Project Reactor?

**Answer:**
(See Q244).

---

### Question 255: How to connect MongoDB reactively with Spring Boot?

**Answer:**
Dependency `spring-boot-starter-data-mongodb-reactive`.
Interface `ReactiveMongoRepository`.
Methods return `Mono` or `Flux`.

---

### Question 256: How do you test WebFlux endpoints?

**Answer:**
Use `WebTestClient`.
`.get().uri("/api").exchange().expectStatus().isOk().expectBodyList(...)`.

---

### Question 257: How does error handling differ in WebFlux?

**Answer:**
No try-catch blocks.
Use operators: `onErrorReturn()`, `onErrorResume()`, `onErrorMap()`.

---

### Question 258: How do you use `StepVerifier` in unit testing?

**Answer:**
Test utility for Reactor types.
```java
StepVerifier.create(flux)
    .expectNext("a")
    .expectNext("b")
    .verifyComplete();
```

---

### Question 259: How to handle timeouts in WebClient calls?

**Answer:**
`.timeout(Duration.ofSeconds(5))`.
Throws `TimeoutException` if no signal received.

---

### Question 260: What is the purpose of `Schedulers.parallel()` in reactive code?

**Answer:**
Specifies on which Thread Pool the operator execution should happen.
`publishOn(Schedulers.boundedElastic())` is crucial when wrapping blocking calls (JDBC) so they don't block the few Netty threads.

---
