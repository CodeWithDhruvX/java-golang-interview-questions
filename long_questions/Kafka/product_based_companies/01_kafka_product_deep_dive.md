# 🏗️ Kafka — Product-Based Companies Deep Dive

> **Level:** 🟡 Intermediate to 🔴 Senior
> **Asked at:** Amazon, Uber, LinkedIn, Netflix, Razorpay, Swiggy

---

## Q1. In a high-throughput system like Netflix or Uber, how does Kafka achieve such incredible speed and what is its underlying storage mechanism?

"Kafka achieves massive throughput primarily through two mechanisms: the **OS Page Cache** and the **Zero-Copy principle**. 

Instead of heavily relying on the JVM heap for in-memory caching, Kafka directly leverages the operating system's page cache. When messages are written, they are appended to log segments sequentially. Sequential disk I/O is extremely fast, often matching memory speeds.

Simultaneously, when a consumer requests data, Kafka bypassed the traditional path of copying data from the disk to the OS kernel, then to the application user space, back to the Kernel space, and off to network. Instead, using the **Zero-Copy** feature (the `sendfile()` system call), the OS copies data directly from the page cache directly deep into the network socket. This drastically reduces CPU overhead and memory bandwidth, enabling Kafka to serve millions of messages per second seamlessly."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, LinkedIn — companies designed around event-driven architectures with huge IO demands.

#### Indepth
**Log Segments & Indexes:** Kafka doesn't use one huge file. It rotates logs into segments based on size or time. Each segment has a `.log` file (the data), an `.index` file (mapping offsets to physical file positions), and a `.timeindex` file. This lets Kafka do binary searches extremely fast to locate messages by offset or timestamp.

---

## Q2. How exactly do you ensure exactly-once processing (EOS) in distributed systems via Kafka?

"Guaranteeing exactly-once processing is natively supported in Kafka by combining two specific features: **Idempotent Producers** and **Transactions**.

First, I configure the producer to be idempotent (`enable.idempotence=true`). When this happens, Kafka assigns a unique Producer ID (`PID`) and a sequence number to every message. If the producer retries sending a message due to a network glitch, the broker sees the same `PID` and sequence number and ignores the duplicate, guaranteeing the message is written exactly once to that specific partition.

Second, the true challenge is when a consumer reads a message, processes it, and writes a result to a new Kafka topic. To ensure this entire loop is atomic, we use the **Kafka Transactions API**. The producer wraps the offset commit (of the consumed message) and the production of the new message within a single `commitTransaction()` call. A specialized Transaction Coordinator manages this state using transaction markers, guaranteeing that either the offset commit and the output message are both visible, or neither is."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Fintechs (Razorpay, PhonePe) and E-commerce where duplicate payments/orders are catastrophic.

#### Indepth
**Transaction Markers and Isolation Levels:** A downstream consumer needs to set `isolation.level=read_committed` to avoid reading messages belonging to aborted transactions. The broker actively masks these aborted messages by using hidden control messages named **Transaction Markers** physically appended to the log.

---

## Q3. Explain the difference between eager and cooperative rebalancing. Why are rebalance storms dangerous?

"Consumer rebalancing occurs as soon as a new consumer joins a group, an existing consumer leaves (or crashes), or partitions are added.

In old-generation **Eager Rebalancing** (or stop-the-world rebalancing), all consumers in the group drop their currently assigned partitions immediately. The entire processing halts while the Group Coordinator decides how to re-distribute the partitions, and only then are operations resumed.

In high traffic deployments, this causes what we call a 'Rebalance Storm'—a severe drop in end-to-end throughput resulting in consumer latency spikes.

Thus, we prefer **Cooperative Rebalancing** (incremental). During this type of rebalance, consumers do *not* drop partitions that they are destined to keep. Only the specific partitions migrating away are temporarily suspended. This enables the unaffected consumers to remain actively processing, eliminating the stop-the-world downtime during cluster scaling or minor faults."

#### 🏢 Company Context
**Level:** 🟡 Intermediate to 🔴 Senior | **Asked at:** Flipkart, Swiggy — handling spontaneous traffic spikes safely.

#### Indepth
**Poison Pill & Health Checks:** Oftentimes, a 'Poison Pill' specific message crashes the consumer sequentially on retries, leading it to timeout and trigger a rebalance continuously. Implementing a Dead Letter Queue (DLQ) enables fast-failing bad events, maintaining stable heartbeat connections to the broker without dropping the consumer.

---

## Q4. How does Kafka deal with fault tolerance? What happens if a broker crashes right down the middle of an operation?

"Kafka implements fault tolerance through **partition replication**. Every partition has one designated 'Leader' and several 'Followers'. All read and write requests go exclusively to the Leader, whilst the Followers strictly synchronize data from it.

Kafka tracks healthy followers using the **In-Sync Replicas (ISR)** list. A follower is deemed 'in-sync' if it has been steadily keeping up with the leader's data offsets within an acceptable timeframe (`replica.lag.time.max.ms`).

If a Leader broker crashes, the Kafka Controller instantly orchestrates a **Leader Election**, choosing a new Leader directly from the existing ISR list. Since ISR members had the exact same data up to the *High Watermark*, no committed data is lost, and client operations route to the new leader transparently."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Amazon, Netflix — building highly resilient data pipelines.

#### Indepth
**What is High Watermark (HW)?** It is the offset of the last message that has been fully replicated to *all* members in the ISR. Consumers are strictly prohibited from reading messages above the HW to prevent data inconsistencies incase a consumer reads an un-replicated message and the leader dies immediately after.

---
