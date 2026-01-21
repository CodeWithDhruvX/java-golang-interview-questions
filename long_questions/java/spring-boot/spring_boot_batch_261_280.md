## 🔹 Section 9: Deployment, Docker & CI/CD (261-280)

### Question 261: How do you build a Spring Boot Docker image?

**Answer:**
1.  **Dockerfile:** `FROM openjdk:17 ... COPY app.jar ...`
2.  **Buildpacks:** `mvn spring-boot:build-image`. (No Dockerfile needed). Uses Cloud Native Buildpacks (Paketo).

---

### Question 262: What is the benefit of layering JARs in Spring Boot?

**Answer:**
Separates Dependencies (changing rarely) via Application Code (changing often).
Docker caches the dependency layer.
Speeds up build/push/pull times for subsequent releases.
Enabled via `<layers><enabled>true</enabled></layers>` in Maven plugin.

---

### Question 263: How do you use Jib for Spring Boot container builds?

**Answer:**
Google's plugin (`jib-maven-plugin`).
Builds optimized images directly to a registry **without Docker daemon**.
Fast and secure.

---

### Question 264: How to optimize Docker image size for Spring Boot apps?

**Answer:**
1.  Use specific JRE (not JDK).
2.  Use slim base images (`alpine` or `distroless`).
3.  Multi-stage builds (Build in one container, copy only JAR to runner container).
4.  Remove unused dependencies (`mvn dependency:analyze`).

---

### Question 265: How to deploy Spring Boot to Kubernetes?

**Answer:**
Create `Deployment.yaml` (Replicas, Image).
Create `Service.yaml` (LoadBalancer/ClusterIP).
Use Probes (Readiness/Liveness) pointing to Actuator.
`kubectl apply -f k8s/`.

---

### Question 266: How to configure Spring Boot with environment variables in Docker?

**Answer:**
Docker: `ENV SERVER_PORT=9090`.
Spring Boot automatically maps Env Vars to properties (Case insensitive, `_` maps to `.`).

---

### Question 267: What is the role of `entrypoint` in Docker for Spring Boot?

**Answer:**
Specifies the command to run.
`ENTRYPOINT ["java", "-jar", "app.jar"]`.
Allows passing additional arguments from the `docker run` command line (`docker run myimg --debug`).

---

### Question 268: How do you implement rolling updates for Spring Boot services?

**Answer:**
K8s handles this natively.
Use `Graceful Shutdown` (Q98) in Spring Boot to ensure zero downtime.
K8s waits for the new pod to be "Ready" (Readiness Probe) before killing old one.

---

### Question 269: How to set up a CI/CD pipeline for Spring Boot in GitHub Actions?

**Answer:**
1.  Checkout Code.
2.  Set up JDK.
3.  `mvn verify` (Build + Test).
4.  `mvn spring-boot:build-image`.
5.  `docker push`.
6.  Trigger K8s rollout.

---

### Question 270: How do you handle configuration secrets in containerized Spring Boot deployments?

**Answer:**
Don't bake them in the image.
Use K8s Secrets injected as Environment Variables or Mounted Volumes.
Spring Cloud Kubernetes can read ConfigMaps/Secrets directly.

---

### Question 271: How do you handle Kafka message retries in Spring Boot?

**Answer:**
(See Q248).

---

### Question 272: How to publish messages conditionally in Kafka?

**Answer:**
Logic inside Producer Service.
`if (txn.isValid()) kafkaTemplate.send(...)`.

---

### Question 273: What are Kafka partitions and how do you manage them in Spring Boot?

**Answer:**
Parallelism unit.
Spring Boot `ConcurrentKafkaListenerContainerFactory` property `concurrency=3`.
Starts 3 consumer threads, each reading from different partitions.

---

### Question 274: How do you use dead letter queues with RabbitMQ?

**Answer:**
Argument `x-dead-letter-exchange` on the Queue definition.
If consumer rejects (nack) with `requeue=false`, broker moves msg to DLQ.

---

### Question 275: How to implement transactional messaging with Kafka?

**Answer:**
`transactionIdPrefix` in ProducerFactory.
Annotate method with `@Transactional`.
Ensures "Atomic" send: Either all messages are sent, or none.

---

### Question 276: What is `KafkaTemplate` and how do you use it?

**Answer:**
(See Q246).

---

### Question 277: How to configure manual acknowledgment in Kafka consumers?

**Answer:**
`factory.setAckMode(AckMode.MANUAL)`.
Listener method receives `Acknowledgment ack`.
Call `ack.acknowledge()` only after successful processing.

---

### Question 278: How do you achieve idempotency in message consumers?

**Answer:**
1.  Use unique Message ID.
2.  Check DB/Redis if ID exists.
3.  If yes, skip. If no, process + save ID.
Atomic transaction required for process + save.

---

### Question 279: How do you configure Avro or Protobuf serialization in Spring Kafka?

**Answer:**
Use Confluent Schema Registry.
Set `value.serializer` properties to `KafkaAvroSerializer`.
Generate Java classes from `.avsc`.

---

### Question 280: What is Spring Integration and how is it used in messaging?

**Answer:**
Framework for Enterprise Integration Patterns (EIP).
Channels, Gateways, Transformers.
Abstracts the transport (can switch from File -> JMS -> Kafka without changing business logic).

---
