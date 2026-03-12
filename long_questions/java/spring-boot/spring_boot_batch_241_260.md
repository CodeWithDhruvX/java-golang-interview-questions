## 🔹 Section 8: Reactive, Kafka & Messaging (241-260)

### Question 241: What is Spring WebFlux and how does it differ from Spring MVC?

**Answer:**
(See Q100). Non-blocking stack vs Blocking stack.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Spring WebFlux and how does it differ from Spring MVC?
**Your Response:** "Spring WebFlux is the reactive web framework in Spring Boot, using a non-blocking stack with Netty and event-loop architecture. Spring MVC uses the traditional blocking stack with servlet containers like Tomcat. WebFlux can handle thousands of concurrent connections with few threads, making it ideal for I/O-bound applications. MVC uses a thread-per-request model which is simpler but less scalable. I choose WebFlux when I need high concurrency and can work with reactive programming, and MVC for traditional applications where simplicity is preferred over maximum scalability."

---

### Question 242: What are Mono and Flux in WebFlux?

**Answer:**
Reactive Streams publishers (Project Reactor).
*   **Mono<T>:** 0 or 1 item. (Like `Optional` + `Future`).
*   **Flux<T>:** 0 to N items. (Like a Stream).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Mono and Flux in WebFlux?
**Your Response:** "Mono and Flux are the core reactive types from Project Reactor. `Mono<T>` represents a stream of 0 or 1 items - like an `Optional` combined with a `Future`. It's perfect for operations that return a single result or nothing, like finding a user by ID. `Flux<T>` represents a stream of 0 to N items - like a traditional stream but asynchronous. I use Flux for operations that return multiple results, like streaming data or querying lists. Both types support rich operators for transformation, filtering, and error handling in a non-blocking way."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement reactive REST endpoints in Spring Boot?
**Your Response:** "I implement reactive REST endpoints by returning `Mono<T>` or `Flux<T>` from my controller methods instead of traditional blocking types. For example, I can return `Flux.interval(Duration.ofSeconds(1)).map(i -> 'Msg ' + i)` to create a streaming endpoint that emits messages every second. Spring WebFlux automatically handles the reactive types and streams the response. This approach allows me to handle streaming data, server-sent events, or any asynchronous operation without blocking threads, making my endpoints highly scalable."

---

### Question 244: How to handle backpressure in WebFlux?

**Answer:**
Handled by the subscriber.
`onBackpressureBuffer()` or `onBackpressureDrop()`.
Allows the consumer to signal "I can only handle 10 items now", preventing OOM errors if producer is too fast.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle backpressure in WebFlux?
**Your Response:** "Backpressure in WebFlux is handled by the subscriber, not the publisher. The consumer can signal how many items it can handle, preventing memory issues when the producer is faster than the consumer. I use operators like `onBackpressureBuffer()` to buffer excess items, or `onBackpressureDrop()` to discard items when the consumer can't keep up. This flow control mechanism prevents out-of-memory errors and ensures stable system performance even under heavy load. It's a key advantage of reactive programming - the system gracefully handles mismatches between production and consumption rates."

---

### Question 245: What is Server-Sent Events (SSE) and how to implement it in Spring Boot?

**Answer:**
One-way streaming channel from Server -> Client.
Return `Flux<String>` with `text/event-stream` media type.
Great for notifications without WebSockets.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Server-Sent Events (SSE) and how to implement it in Spring Boot?
**Your Response:** "Server-Sent Events provide a one-way streaming channel from server to client over HTTP. I implement SSE in Spring Boot by returning a `Flux<String>` with the `text/event-stream` media type. This allows me to push real-time updates to clients without the overhead of WebSockets. SSE is perfect for notifications, live updates, or progress reporting where I only need server-to-client communication. The client automatically handles reconnections and event parsing, making it simpler than WebSockets for one-way streaming scenarios."

---

### Question 246: How to connect Spring Boot with Kafka for message publishing?

**Answer:**
Dependency: `spring-kafka`.
Inject `KafkaTemplate<String, String>`.
`template.send("topic", "key", "payload")`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to connect Spring Boot with Kafka for message publishing?
**Your Response:** "I connect Spring Boot with Kafka by adding the `spring-kafka` dependency and injecting a `KafkaTemplate<String, String>`. The template simplifies message publishing - I just call `template.send('topic', 'key', 'payload')` to send messages. Spring Boot auto-configures the producer with sensible defaults, but I can customize settings like bootstrap servers and serialization in properties. This approach makes Kafka integration straightforward while still allowing full control over the producer configuration. I can also use typed templates for specific message formats."

---

### Question 247: How to consume Kafka messages in Spring Boot?

**Answer:**
Method annotation:
```java
@KafkaListener(topics = "my-topic", groupId = "group_id")
public void listen(String message) { ... }
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to consume Kafka messages in Spring Boot?
**Your Response:** "I consume Kafka messages using the `@KafkaListener` annotation on a method. I specify the topics to listen to and the consumer group ID. Spring Boot automatically creates the consumer and invokes my method for each message. The framework handles all the complexity of consumer lifecycle, offset management, and deserialization. I can also add parameters for headers, acknowledgments, or consumer-aware processing. This declarative approach makes Kafka consumption clean and straightforward while still giving me access to all Kafka features."

---

### Question 248: How to handle Kafka error handling and retries?

**Answer:**
Config `DefaultErrorHandler` with `FixedBackOff`.
If retry exhausted, send to Dead Letter Topic (DLT) using `DeadLetterPublishingRecoverer`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle Kafka error handling and retries?
**Your Response:** "I handle Kafka errors and retries using `DefaultErrorHandler` with a `FixedBackOff` strategy. When a message processing fails, Spring Kafka automatically retries with the configured backoff policy. If all retries are exhausted, I use `DeadLetterPublishingRecoverer` to send the failed message to a Dead Letter Topic (DLT) for later analysis. This approach ensures that failed messages don't block the consumer and can be investigated manually. I can also implement custom error handlers for specific exception types or business logic for error recovery."

---

### Question 249: How to test Kafka consumers in Spring Boot?

**Answer:**
`@EmbeddedKafka`.
Starts an in-memory Kafka broker for the test class.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to test Kafka consumers in Spring Boot?
**Your Response:** "I test Kafka consumers using the `@EmbeddedKafka` annotation, which starts an in-memory Kafka broker specifically for the test class. This allows me to test consumer logic without requiring an external Kafka cluster. The embedded broker provides all Kafka functionality but runs locally within the test, making tests fast and isolated. I can send test messages to the embedded broker and verify that my consumer processes them correctly. This approach gives me reliable, fast integration tests for Kafka functionality without external dependencies."

---

### Question 250: How does Spring Boot support RabbitMQ integration?

**Answer:**
Dependencies `spring-boot-starter-amqp`.
Auto-configures `ConnectionFactory`.
Use `RabbitTemplate` for sending.
Use `@RabbitListener` for receiving.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Spring Boot support RabbitMQ integration?
**Your Response:** "Spring Boot supports RabbitMQ through the `spring-boot-starter-amqp` dependency. It auto-configures the `ConnectionFactory` and provides `RabbitTemplate` for sending messages and `@RabbitListener` for receiving messages. I can simply inject the template to send messages or annotate methods to consume messages from queues. Spring Boot handles connection management, message conversion, and error handling automatically. This makes RabbitMQ integration straightforward while still allowing full customization of exchanges, queues, and routing through configuration."

---

### Question 251: What are cold and hot publishers in WebFlux?

**Answer:**
*   **Cold:** Data generation starts **only** when subscribed. Each subscriber gets their own fresh stream (e.g., HTTP Call).
*   **Hot:** Data flows regardless of subscribers. Late subscribers miss old data (e.g., Mouse Events, Ticker).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are cold and hot publishers in WebFlux?
**Your Response:** "Cold publishers start generating data only when a subscriber subscribes, and each subscriber gets their own fresh stream. HTTP calls are typical cold publishers - each subscriber triggers a new request. Hot publishers emit data regardless of whether anyone is subscribed, and late subscribers miss data that was emitted before they subscribed. Mouse events or stock tickers are hot publishers. The choice affects behavior - cold publishers give each subscriber independent data, while hot publishers share the same data stream among all subscribers."

---

### Question 252: How does `WebClient` compare to `RestTemplate`?

**Answer:**
(See Q60). Functional, fluent API. Supports Streaming.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does `WebClient` compare to `RestTemplate`?
**Your Response:** "`WebClient` is the modern, reactive alternative to `RestTemplate`. It provides a functional, fluent API and supports streaming and non-blocking I/O, making it perfect for reactive applications. `RestTemplate` is blocking and simpler to use for basic HTTP calls. I prefer `WebClient` for new applications because it's more flexible, supports reactive streams, and works well with WebFlux. `RestTemplate` is still viable for simple synchronous calls, but `WebClient` is the future-proof choice that can handle both synchronous and asynchronous scenarios."

---

### Question 253: How to return streaming JSON from WebFlux endpoint?

**Answer:**
`MediaType.APPLICATION_NDJSON` (Newline Delimited JSON).
Allows client to parse array items one by one as they arrive.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to return streaming JSON from WebFlux endpoint?
**Your Response:** "I return streaming JSON from WebFlux endpoints using `MediaType.APPLICATION_NDJSON` (Newline Delimited JSON). This allows me to stream a Flux of objects where each object is on its own line. The client can parse JSON items one by one as they arrive, rather than waiting for the complete array. This is perfect for large datasets or real-time data where the client can start processing immediately instead of waiting for the full response. NDJSON provides better memory usage and faster time-to-first-byte compared to traditional JSON arrays."

---

### Question 254: How do you implement reactive backpressure with Project Reactor?

**Answer:**
(See Q244).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement reactive backpressure with Project Reactor?
**Your Response:** "I implement reactive backpressure using Project Reactor's built-in operators. The subscriber controls the flow by requesting a specific number of items from the publisher. I use `onBackpressureBuffer()` to buffer excess items when the consumer is slower, or `onBackpressureDrop()` to discard items. I can also use `onBackpressureLatest()` to keep only the latest item. This flow control prevents memory issues and ensures system stability when there's a mismatch between production and consumption rates. It's a key feature that makes reactive systems resilient under varying loads."

---

### Question 255: How to connect MongoDB reactively with Spring Boot?

**Answer:**
Dependency `spring-boot-starter-data-mongodb-reactive`.
Interface `ReactiveMongoRepository`.
Methods return `Mono` or `Flux`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to connect MongoDB reactively with Spring Boot?
**Your Response:** "I connect MongoDB reactively using the `spring-boot-starter-data-mongodb-reactive` dependency. I extend `ReactiveMongoRepository` instead of the traditional `MongoRepository`, which gives me methods that return `Mono` or `Flux` instead of blocking results. This allows me to perform database operations without blocking threads, making my entire application stack reactive. Spring Boot handles the reactive driver configuration automatically, and I can chain reactive operations from the database layer through my service layer to the web layer."

---

### Question 256: How do you test WebFlux endpoints?

**Answer:**
Use `WebTestClient`.
`.get().uri("/api").exchange().expectStatus().isOk().expectBodyList(...)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test WebFlux endpoints?
**Your Response:** "I test WebFlux endpoints using `WebTestClient`, which is designed for testing reactive web applications. I use methods like `.get().uri('/api').exchange().expectStatus().isOk().expectBodyList(...)` to test my endpoints. `WebTestClient` works with both MVC and WebFlux applications and can handle streaming responses and asynchronous operations. It provides a fluent API for testing HTTP requests and responses, including support for JSON body assertions, header validation, and status code verification. It's the modern alternative to MockMvc for reactive applications."

---

### Question 257: How does error handling differ in WebFlux?

**Answer:**
No try-catch blocks.
Use operators: `onErrorReturn()`, `onErrorResume()`, `onErrorMap()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does error handling differ in WebFlux?
**Your Response:** "Error handling in WebFlux doesn't use traditional try-catch blocks. Instead, I use reactive operators like `onErrorReturn()` to provide a fallback value, `onErrorResume()` to switch to an alternative flux, or `onErrorMap()` to transform exceptions. These operators allow me to handle errors declaratively within the reactive chain. This approach fits naturally with the reactive programming model and ensures that error handling doesn't break the reactive flow. It's more expressive and composable than traditional exception handling."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `StepVerifier` in unit testing?
**Your Response:** "`StepVerifier` is a testing utility for Reactor types that allows me to test reactive streams declaratively. I use `StepVerifier.create(flux)` to start testing, then chain expectations like `expectNext('a')` to verify emitted items, and finally `verifyComplete()` to ensure the stream completed successfully. I can also test error conditions with `verifyError()`. This approach makes testing reactive streams clean and readable, allowing me to verify the exact sequence of events in my reactive pipelines."

---

### Question 259: How to handle timeouts in WebClient calls?

**Answer:**
`.timeout(Duration.ofSeconds(5))`.
Throws `TimeoutException` if no signal received.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How to handle timeouts in WebClient calls?
**Your Response:** "I handle timeouts in WebClient calls using the `.timeout(Duration.ofSeconds(5))` operator. If no signal is received within the specified duration, it throws a `TimeoutException`. I can combine this with `onErrorResume()` to provide fallback behavior when timeouts occur. This ensures that my reactive HTTP calls don't hang indefinitely and that I can gracefully handle slow or unresponsive services. Timeout handling is crucial in reactive systems to maintain responsiveness and prevent resource exhaustion."

---

### Question 260: What is the purpose of `Schedulers.parallel()` in reactive code?

**Answer:**
Specifies on which Thread Pool the operator execution should happen.
`publishOn(Schedulers.boundedElastic())` is crucial when wrapping blocking calls (JDBC) so they don't block the few Netty threads.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `Schedulers.parallel()` in reactive code?
**Your Response:** "`Schedulers.parallel()` specifies which thread pool reactive operators should execute on. In WebFlux, the default Netty server uses only a few threads, so blocking operations can starve the entire application. I use `publishOn(Schedulers.boundedElastic())` when wrapping blocking calls like JDBC to ensure they run on a separate thread pool designed for blocking operations. This prevents blocking calls from consuming the precious event-loop threads and maintains the non-blocking nature of my reactive application. Proper scheduler usage is crucial for reactive performance."

---
