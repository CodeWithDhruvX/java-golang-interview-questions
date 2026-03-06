# 🏗️ Generic Kafka Interview Answers (Java & Golang)

This guide provides generic answers for common Kafka interview questions that apply to **both Java and Golang** developers. It also highlights the different expectations and follow-up questions you might face at **Product-Based** vs. **Service-Based** companies.

---

## 1. What is Kafka and what are its core components?

**Generic Answer:**
"Apache Kafka is a distributed event streaming platform used for high-throughput, low-latency data pipelines. Unlike traditional message queues (like RabbitMQ) that delete messages upon consumption, Kafka is an append-only, immutable commit log. Its core components are:
*   **Producer:** Applications that publish messages to Kafka.
*   **Consumer:** Applications that read messages from Kafka.
*   **Broker:** A single Kafka server that stores data and serves client requests. A cluster contains multiple brokers.
*   **Topic:** A logical category where messages are published.
*   **Partition:** A topic is divided into partitions for scalability. Partitions allow messages to be consumed in parallel across multiple consumers.
*   **ZooKeeper / KRaft:** Used for managing cluster metadata and broker consensus."

### 💻 Language Nuances
*   **Java:** You'll typically mention using `spring-kafka` (with `@KafkaListener` and `KafkaTemplate`) or the vanilla `kafka-clients` library.
*    **Golang:** You'll mention using `confluent-kafka-go` (a CGO wrapper around `librdkafka` offering high performance) or `IBM/sarama` (a pure Go client, good when CGO is not an option).

### 🏢 Company Context
*   **Service-Based (TCS, Infosys, Cognizant):** Focus on the definitions, the difference between Kafka and RabbitMQ, and basic component architecture.
*   **Product-Based (FAANG+, Uber, Flipkart):** Focus on *why* Kafka is a commit log, how sequential disk I/O gives it high throughput, and the exact role of KRaft (removing ZooKeeper dependency).

---

## 2. Explain Consumer Groups and Partition balancing.

**Generic Answer:**
"A Consumer Group is a set of consumers cooperating to read from a topic. Kafka guarantees that each partition in a topic is consumed by exactly ONE consumer within a consumer group. This is how Kafka scales consumption:
*   If # Consumers < # Partitions: Some consumers read from multiple partitions.
*   If # Consumers == # Partitions: Each consumer reads from exactly one partition (Ideal for max parallelism).
*   If # Consumers > # Partitions: The extra consumers sit idle (used for failover).
If a consumer dies, a **Rebalance** occurs, and its partitions are reassigned to the remaining live consumers."

### 💻 Language Nuances
*   **Java:** In Spring Boot, concurrency is easily managed by setting `concurrency="3"` on `@KafkaListener` to spin up multiple consumer threads within the same JVM instance.
*   **Golang:** The standard pattern is a single consumer reading messages in a loop and fanning out processing to multiple Goroutines via channels (worker pools). However, parallel processing within a single partition can break ordering, so you usually consume per-partition and process sequentially per partition if ordering matters.

### 🏢 Company Context
*   **Service-Based:** You will be asked purely about the math (e.g., "I have 10 partitions and 5 consumers, how are they distributed?").
*   **Product-Based:** Expect deep questions on **Stop-the-world Rebalancing**. How do you prevent aggressive rebalancing? (Tuning `session.timeout.ms`, `max.poll.interval.ms`). What are Cooperative Sticky Assignors?

---

## 3. What are Offset Commits and Delivery Guarantees?

**Generic Answer:**
"An offset is a sequential ID number given to messages within a partition. It acts as the consumer's bookmark.
*   **At-Most-Once:** Consumer commits the offset *before* processing. If it crashes during processing, the message is lost. Fast, but unsafe. (Auto-commit enabled).
*   **At-Least-Once:** Consumer processes the message *first*, then commits the offset manually. If it crashes after processing but before committing, the message is re-delivered (Duplicate). This is the industry standard.
*   **Exactly-Once (EOS):** Uses Kafka Transactions and Idempotent Producers (`enable.idempotence=true`). Ensures that a read-process-write cycle is atomic."

### 💻 Language Nuances
*   **Java:** `consumer.commitSync()` vs `consumer.commitAsync()`. In Spring, `Acknowledgment.acknowledge()` is used in `MANUAL` ack modes.
*   **Golang:** With `confluent-kafka-go`, you disable `enable.auto.commit` and use `consumer.CommitMessage(msg)` explicitly after processing. Goroutine crashes handling manual commits are common topics to discuss.

### 🏢 Company Context
*   **Service-Based:** Mostly interested in At-Least-Once vs At-Most-Once. They will expect you to know that `enable.auto.commit=false` is the best practice.
*   **Product-Based:** Extremely focused on idempotency. "How do you handle the duplicates in At-Least-Once?" (Answer: Use a database with a unique constraint on the event ID or a Redis lookup table). They will also grill you on the performance overhead of Exactly-Once Semantics (EOS).

---

## 4. How does Kafka achieve high availability and fault tolerance?

**Generic Answer:**
"Kafka uses **Replication**. Every partition has one Leader and multiple Followers.
*   **Leader:** Handles all read and write requests for the partition.
*   **Followers:** Passively replicate data from the Leader.
*   **In-Sync Replicas (ISR):** The subset of followers that are fully caught up with the leader.
If a leader goes down, one of the replicas in the ISR is automatically promoted to be the new leader.
Producers can configure `acks`:
*   `acks=0`: Fire and forget (fastest, high data loss risk).
*   `acks=1`: Leader acknowledges receipt (moderate speed, some data loss risk if leader dies before replicating).
*   `acks=all` (or `-1`): Leader waits for all ISRs to acknowledge (slowest, strongest durability)."

### 💻 Language Nuances
*   Configuration parameters (`acks`, `min.insync.replicas`) are exactly the same across both Java `Properties` objects and Golang `ConfigMap` objects.

### 🏢 Company Context
*   **Service-Based:** "What is an ISR?" and "What do acks=all mean?"
*   **Product-Based:** "What happens to the cluster if `acks=all` but `min.insync.replicas` is equal to your replication factor, and one node goes down?" (Answer: Producer throws `NotEnoughReplicasException`, writing stops to maintain structural integrity).

---

## 5. How do you tune Kafka for High Throughput vs Low Latency?

**Generic Answer:**
"Kafka's performance can be heavily tuned:
*   **For High Throughput (Batching):** I tune the producer using `batch.size` (e.g., 64KB) and `linger.ms` (e.g., 10-20ms). This forces the producer to wait slightly to group more messages together into a single network payload. I also enable compression (`compression.type=snappy` or `lz4`).
*   **For Low Latency:** I keep `linger.ms=0`. The producer sends immediately. Consumers should be written in a lightweight manner."

### 💻 Language Nuances
*   **Java:** The JVM Garbage Collector can cause latency spikes if batching consumes too much memory heap space. Tuning G1GC or ZGC is part of the conversation.
*   **Golang:** Go's garbage collector is optimized for low latency. Golang applications naturally excel at high concurrency throughput via Goroutines, meaning a single Go microservice can often process higher throughput with lower CPU memory footprint compared to a typical Spring Boot app.

### 🏢 Company Context
*   **Service-Based:** Focuses on the definitions of `batch.size` and `linger.ms`.
*   **Product-Based:** Expect architectural questions like: "If you enable Snappy compression, does it compress on the producer, broker, or consumer?" (Answer: Producer compresses, broker stores compressed, consumer uncompresses. Beware of double-compression overhead if broker config differs from producer).
