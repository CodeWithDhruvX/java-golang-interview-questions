# 33. Messaging (Kafka) & Containerization (Docker/Kubernetes)

**Q: Server-Sent Events (SSE)**
> "SSE is a one-way communication channel from Server to Client.
> Unlike WebSockets (which are bidirectional), SSE is simpler.
>
> In Spring Boot:
> Return `Flux<ServerSentEvent<String>>`.
> The browser keeps the connection open, and you can push stock updates or notifications in real-time."

**Indepth:**
> **Reconnection**: Standard HTTP requests don't auto-reconnect. SSE has built-in reconnection logic. If the connection drops, the browser automatically tries to reconnect, sending the "Last-Event-ID" so the server can resume from where it left off.


---

**Q: Kafka Producer/Consumer (Spring Boot)**
> "It's all about `KafkaTemplate` and `@KafkaListener`.
> 1.  **Publishing**: `kafkaTemplate.send("topic_name", "message")`.
> 2.  **Consuming**:
>     ```java
>     @KafkaListener(topics = "topic_name", groupId = "my-group")
>     public void listen(String message) {
>         // Process message
>     }
>     ```"

**Indepth:**
> **Serialization**: Spring Boot uses `StringSerializer` by default for keys and values. In production, you'll likely switch the Value serializer to `JsonSerializer` (Jackson) to send complex objects easily.


---

**Q: Kafka Error Handling (Retries/DLQ)**
> "What if processing a message fails?
> 1.  **Retry**: Configure a `DefaultErrorHandler` with a `FixedBackOff`. It retries 3 times.
> 2.  **Dead Letter Queue (DLQ)**: If it still fails, send the message to a separate topic (`orders-dlt`). You can inspect these later manually."

**Indepth:**
> **Non-Blocking**: By default, retries might block the consumer thread, stopping it from processing *other* messages. **Non-Blocking Retries** (using `@RetryableTopic`) publish the failed message to a delay-queue topic, freeing up the consumer immediately.


---

**Q: WebClient vs RestTemplate**
> "**RestTemplate** is blocking. It waits for the response. Deprecated (in maintenance mode).
>
> "**WebClient** is non-blocking (Reactive).
> It uses Netty. It allows you to make parallel calls easily:
> `Mono.zip(callA(), callB())`.
> Even if you use blocking Spring MVC, you should start using WebClient for external API calls."

**Indepth:**
> **Resources**: `RestTemplate` creates a new Thread for every request. If you call 100 external APIs, you potentially block 100 threads. `WebClient` can handle 100 requests with just 1 thread using Non-Blocking IO.


---

**Q: Dockerizing Spring Boot**
> "The simplest way:
> 1.  Build the jar: `mvn clean package`.
> 2.  Write a `Dockerfile`:
>     ```dockerfile
>     FROM openjdk:17-alpine
>     COPY target/app.jar app.jar
>     ENTRYPOINT ["java", "-jar", "app.jar"]
>     ```
> 3.  `docker build -t my-app .`"

**Indepth:**
> **Multi-Stage**: Use Multi-Stage Docker builds to optimize image size. Stage 1 (Maven) builds the jar (requires 500MB of deps). Stage 2 (JRE) only copies the final jar (requires 50MB). The final image is tiny.


---

**Q: Layered JARs (Optimization)**
> "A standard Spring Boot JAR is huge (App Code + 50MB of Libraries).
> If you change one line of code, Docker has to re-push the whole 50MB layer.
>
> **Layered JARs** separate them:
> Layer 1: Dependencies (rarely change).
> Layer 2: Your Code (changes often).
> Docker reuses Layer 1 from cache and only pushes Layer 2. Faster builds, faster deployments."

**Indepth:**
> **Cache**: The `spring-boot-maven-plugin` has a `layers` configuration. When enabled, it splits `dependencies`, `spring-boot-loader`, `snapshot-dependencies`, and `application` classes into separate folders in the docker image specifically for caching.


---

**Q: Jib (Google Tool)**
> "Jib allows you to build Docker images **without** a Docker daemon and **without** a Dockerfile.
>
> You just add the `jib-maven-plugin` plugin to your pom.xml.
> Run `mvn jib:build`.
> It analyzes your project, intelligently layers it, and pushes it directly to a registry (like Docker Hub)."

**Indepth:**
> **Reproducibility**: Jib separates the application from the OS. It doesn't use a Dockerfile, so "It works on my machine" issues related to different base OS installations are minimized.


---

**Q: Spring Boot on Kubernetes (K8s)**
> "Spring Boot runs naturally on K8s.
>
> **Configuration**: Use `ConfigMaps` and `Secrets` mapped to environment variables.
> **Health**: K8s uses 'Probes' to check if your app is alive. Map them to Actuator:
> *   Liveness Probe -> `/actuator/health/liveness`
> *   Readiness Probe -> `/actuator/health/readiness`"

**Indepth:**
> **Graceful Shutdown**: Configure `server.shutdown=graceful`. When K8s kills a pod, Spring Boot will stop accepting new requests but will wait (e.g., 30s) for existing requests to finish processing before shutting down the JVM.


---

**Q: Rolling Updates**
> "K8s handles this. You don't do it in Spring.
> You tell K8s: 'Update to version 2.0'.
> K8s spins up a v2 pod. Waits for the **Readiness Probe** to pass. Then kills a v1 pod. Then repeats.
> Zero downtime."

**Indepth:**
> **Deployment Strategies**: Beyond basic rolling updates, K8s supports "Blue-Green" (spin up full v2 parallel to v1, then switch traffic) and "Canary" (send 5% of traffic to v2 to test it) deployments.


---

**Q: Idempotency in Consumers**
> "In Kafka, you might receive the same message twice (At-Least-Once delivery).
> Your consumer **must** be idempotent.
>
> Strategy:
> 1.  Use a unique `message_id`.
> 2.  Maintain a 'Processed IDs' table in your DB.
> 3.  Check: `if (repo.exists(id)) return;` before processing."

**Indepth:**
> **Transactionality**: For exactly-once semantics inside the Kafka ecosystem, you can use Kafka Transactions (`producer.send` + `consumer.commit` are atomic). But for external side-effects (DB writes), Idempotency keys are safer.


---

**Q: Avro/Protobuf (Schema Registry)**
> "Sending raw JSON is wasteful and error-prone.
> **Avro** is a binary format. It's smaller and faster.
>
> You use a **Schema Registry**. The Producer checks the schema ID, sends binary data. The Consumer downloads the schema and deserializes it. It ensures structural compatibility (Contract Testing) automatically."

**Indepth:**
> **Evolution**: Schema Registry allows Schema Evolution. You can add a nullable field to your user object, and old consumers (that don't know about the field) will simply ignore it, ensuring backward compatibility without breaking the system.

