# Enterprise Integration - Interview Answers

> 🎯 **Focus:** These are "Senior" topics. Show you know how systems talk to each other reliably.

### 1. Spring Batch vs Spring Scheduler?
"**Spring Scheduler** is for simple, repetitive tasks. 'Run this method every 5 minutes'. It doesn't have state. If the server crashes, the job is missed.

**Spring Batch** is for processing huge volumes of data (ETL). It has built-in transaction management, retry logic, and restartability. If a batch job fails at record 50,000 of 1,000,000, I can restart it exactly from record 50,001. You can't do that with a Scheduler."

---

### 2. How does a Spring Batch Job work (Chunk-Oriented)?
"It follows a `Reader` -> `Processor` -> `Writer` flow.
You define a **Chunk Size** (say, 100).
1. The **ItemReader** reads 100 records one by one.
2. The **ItemProcessor** processes them one by one (e.g., business logic validation).
3. The **ItemWriter** writes all 100 records in a *single* transaction (batch insert).
This is much more efficient than committing every single row."

---

### 3. Kafka vs RabbitMQ?
"**RabbitMQ** is a traditional Message Queue. It's 'smart broker, dumb consumer'. It tracks who read what. Once a message is consumed, it's gone. Good for complex routing.

**Kafka** is a Distributed Streaming Platform. It's 'dumb broker, smart consumer'. It stores messages as a log on disk for retention (e.g., 7 days). Consumers track their own offset. It's designed for massive throughput and replaying history."

---

### 4. What are Kafka Partitions?
"A topic is split into **Partitions** to allow scaling.
If I have a 'UserEvents' topic with 3 partitions, I can have 3 consumers reading in parallel—one from each partition.
Ordering is guaranteed *only within a partition*, not across the whole topic. So if I need order for a specific user, I ensure all events for that User ID go to the same partition using the Partition Key."

---

### 5. What is a Dead Letter Queue (DLQ)?
"It’s a safety net for bad messages.
If a consumer fails to process a message (e.g., JSON parse error), we don't want to get stuck in an infinite retry loop blocks everything else.
After X retries, we move that message to a special queue called `DLQ`.
Developers can then manually inspect the DLQ, fix the bug or data, and re-process the message later."

---

### 6. Spring Integration vs Apache Camel?
"Both are implementations of 'Enterprise Integration Patterns' (EIP).
**Apache Camel** is a standalone giant. It has connectors for everything under the sun (Twitter, FTP, Salesforce). It has its own DSL.

**Spring Integration** is native to the Spring ecosystem. It uses concepts like `Channels`, `Gateways`, and `ServiceActivators`. It feels more like writing standard Spring Beans.
I generally prefer Spring Integration if I'm already in a Spring Boot app, as it's lighter."

---

### 7. What is an Integration Flow?
"In Spring Integration, it describes the pipeline a message travels through.
It typically looks like:
`Inbound Adapter` (e.g., File Reader) -> `Channel` -> `Transformer` (XML to JSON) -> `Filter` (discard invalid) -> `Service Activator` (Business Logic) -> `Outbound Adapter` (FTP Upload).
It allows us to decouple the 'How' we receive data from 'What' we do with it."

---

### 8. Idempotency in Messaging?
"Idempotency means processing the same message twice has the same effect as processing it once.
In distributed systems, networks fail, and brokers might deliver a message twice (At-Least-Once delivery).
To handle this, I usually track a `MessageID` in a database table. Before processing, I check: 'Have I seen this ID before?'. If yes, I acknowledge and ignore it."

---

### 9. What is `@Retryable`?
"It’s a Spring annotation to automatically retry a failed method.
I stick it on a service method: `@Retryable(value = SQLException.class, maxAttempts = 3, backoff = @Backoff(delay = 1000))`.
If the DB is momentarily down, Spring intercepts the exception, waits 1 second, and tries again. It’s a clean way to handle transient failures without writing `while` loops."

---

### 10. How to scale a Spring Batch Job?
"We can use **Partitioning**.
We split the data into partitions (e.g., Partition 1: IDs 1-10000, Partition 2: IDs 10001-20000).
Each partition is handled by a separate thread (local) or even a separate worker JVM (remote partitioning).
This allows us to process millions of records in parallel, significantly reducing the total batch window time."
