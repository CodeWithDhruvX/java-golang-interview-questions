# 🟢 **1–24: Junior Level (0-2 Years)**

### 1. What is Apache Kafka?
"Apache Kafka is an open-source distributed event streaming platform used for high-performance data pipelines, streaming analytics, and data integration.

Unlike traditional message queues, Kafka operates as a distributed commit log. It allows me to publish events, store them durably and fault-tolerantly, and process them in real-time or retrospectively. 

I use Kafka when I need to handle massive throughput and want multiple independent consumers to read the same stream of events at their own pace without destroying the message upon consumption."

#### Indepth
Kafka is written in Scala and Java. By decoupling producers and consumers and relying heavily on the OS page cache and sequential disk I/O, it achieves millions of messages per second with sub-millisecond latency. It forms the backbone of event-driven microservices architectures.

---

### 2. Explain Kafka architecture.
"Kafka's architecture consists of a cluster of Brokers, Producers, Consumers, and historically, ZooKeeper (now transitioning to KRaft).

Producers push data to Topics, which are split into Partitions across multiple Brokers for scalability. Consumers read from these Partitions. For high availability, each Partition has one Leader replica and multiple Follower replicas. ZooKeeper or KRaft manages the cluster metadata, elects the controller, and tracks the health of the brokers."

#### Indepth
The core of Kafka's architecture is its 'dumb broker / smart consumer' design. Kafka brokers do not track which messages have been read by which consumers. Instead, consumers keep track of their own 'Offset'. This significantly reduces the overhead on the broker, enabling scaling to massive numbers of consumers and partitions.

---

### 3. What are Producer, Consumer, Broker?
"A **Broker** is a single Kafka server that receives messages, stores them on disk, and serves them to consumers. A cluster usually has multiple brokers.

A **Producer** is a client application that publishes (writes) events to Kafka topics.

A **Consumer** is a client application that subscribes to (reads) events from topics. They usually operate as part of a Consumer Group to parallelize the reading process."

#### Indepth
Producers can choose whether they want acknowledgment of their writes (acks=0, 1, or all). Brokers enforce data retention policies, deleting old data regardless of whether it has been consumed. Consumers pull data instead of having it pushed to them, which prevents them from being overwhelmed.

---

### 4. What is a Topic and Partition?
"A **Topic** is a logical category or feed name where records are published. You can think of it like a table in a database.

A **Partition** is the physical unit of a Topic. A Topic is split into one or more partitions to allow for parallel processing. Each partition is an ordered, immutable sequence of records continuously appended to a structured commit log."

#### Indepth
Partitions are the fundamental unit of scalability in Kafka. If a topic has 10 partitions, up to 10 consumers in a consumer group can read concurrently. The number of partitions determines the maximum parallelism of consumption. However, ordering is only guaranteed *within* a single partition, not across the entire topic.

---

### 5. What is Offset?
"An **Offset** is a unique, sequential ID number assigned to each message within a partition. It represents the position of a message.

Consumers use the offset to track their progress. They commit their current offset back to Kafka (usually an internal `__consumer_offsets` topic) so that if they crash, they can resume reading from exactly where they left off."

#### Indepth
Because offsets are strictly increasing integers within a partition, they simplify consumer state management. A consumer can even intentionally reset its offset to an older value to "replay" messages, which is impossible in traditional message queues where reading a message destroys it.

---

### 6. What is ZooKeeper role in Kafka?
"Historically, ZooKeeper was the central coordinator for a Kafka cluster.

It managed cluster metadata, detected broker failures, elected the cluster controller, and managed topic configurations and ACLs. Brokers communicated with ZooKeeper to know the state of the cluster.

However, I should note that Kafka is currently phasing out ZooKeeper in favor of its own built-in consensus protocol called KRaft (Kafka Raft), which simplifies deployment and improves scalability."

#### Indepth
In ZooKeeper mode, the separation of data/log storage (Kafka brokers) and metadata storage (ZooKeeper) caused bottlenecks at high partition counts (e.g., millions of partitions). KRaft moves metadata management directly into the broker quorum using an event-sourced architecture, making controller failovers near-instantaneous.

---

### 7. What is replication factor?
"The **Replication Factor** defines how many copies of the data are stored across the Kafka cluster for fault tolerance.

If I set a replication factor of 3 for a topic, it means the data is stored on one Leader broker and replicated to two Follower brokers. If the Leader goes down, one of the Followers seamlessly takes over without data loss."

#### Indepth
A replication factor of `N` allows the cluster to tolerate `N-1` broker failures if `acks=1` or `acks=all` with `min.insync.replicas` configured appropriately. In production, a replication factor of 3 is the industry standard.

---

### 8. Difference between Kafka and RabbitMQ?
"The main difference is their design philosophy. Kafka is a distributed append-only log, while RabbitMQ is a traditional message broker.

RabbitMQ pushes messages to consumers and deletes them once acknowledged. It excels in complex routing (fanout, topic, direct exchanges).

Kafka requires consumers to pull messages, retains them on disk for a configured time, and allows replay. It is designed for massive throughput, stream processing, and event sourcing."

#### Indepth
Use RabbitMQ for task queues, point-to-point messaging, or when complex routing is required. Use Kafka for massive telemetry data, real-time analytics, event sourcing, or when you need multiple independent systems to consume the same stream independently.

---

### 9. What is Consumer Group?
"A **Consumer Group** is a set of consumers working together to consume data from a topic.

Kafka ensures that each partition in the topic is read by exactly one consumer within the group. This mechanism enables horizontal scaling. If I have a topic with 4 partitions, I can spin up 4 consumers in a group, and each will process 1 partition concurrently."

#### Indepth
If there are more consumers than partitions, the extra consumers sit idle. If a consumer crashes, Kafka triggers a "rebalance", reassigning its partitions to the remaining healthy consumers. This makes consumer groups highly available and elastic.

---

### 10. What is ISR?
"**ISR** stands for **In-Sync Replicas**. It is the subset of replicas for a partition that are fully caught up with the Leader.

For example, if the replication factor is 3, the Leader and two Followers should ideally be in the ISR. If a Follower goes offline or falls too far behind the Leader (configured by `replica.lag.time.max.ms`), it is kicked out of the ISR."

#### Indepth
Only replicas in the ISR are eligible to become the new Leader if the current Leader fails. When a producer uses `acks=all`, it waits for acknowledgment from all replicas currently in the ISR. To prevent data loss if the ISR shrinks to just 1 (the leader), you configure `min.insync.replicas` (typically 2).

---

### 11. What are Kafka delivery semantics?
"Kafka supports three message delivery semantics:

1. **At-most-once**: Messages may be lost but never duplicated. The producer doesn't wait for acks, or the consumer commits offsets before processing.
2. **At-least-once**: Messages are never lost but may be duplicated. The producer retries on failure, or the consumer commits offsets *after* processing.
3. **Exactly-once (EOS)**: The holy grail. Messages are processed exactly once, even in the event of retries or failures. Kafka implements this using idempotent producers and transactional APIs."

#### Indepth
At-least-once is the default and most common. To achieve exactly-once in a Kafka-to-Kafka pipeline (like Kafka Streams), you set `processing.guarantee="exactly_once_v2"`. This ensures that reading an event, processing it, and writing the result happens atomically.

---

### 12. What is retention policy?
"The **Retention Policy** dictates how long Kafka keeps messages before deleting them to free up disk space.

It can be configured based on **time** (e.g., delete messages older than 7 days via `log.retention.hours`) or **size** (e.g., delete when a partition exceeds 100GB via `log.retention.bytes`).

Because of this, Kafka is not just a messaging system; it can act as a durable storage layer."

#### Indepth
Kafka actually deletes data at the 'segment' level, not message by message. When a log segment file is "rolled" (closed) and its contents exceed the retention limit, the entire segment file is deleted from the filesystem.

---

### 13. How to create a topic?
"I can create a topic using the Kafka CLI tools, specifically `kafka-topics.sh` (or `.bat` on Windows).

A typical command looks like this:
`kafka-topics.sh --create --topic my-logs --partitions 3 --replication-factor 3 --bootstrap-server localhost:9092`

Alternatively, topics can be created automatically if `auto.create.topics.enable` is true in the broker config, or programmatically using the `AdminClient` API in languages like Java or Go."

#### Indepth
In production, `auto.create.topics.enable` is always set to `false` to maintain strict governance. Creation is usually handled via CI/CD pipelines using Terraform or Kafka GitOps tools to ensure correct partition counts and retention configs.

---

### 14. What is leader and follower in Kafka?
"For a given partition, the **Leader** is the replica that handles all read and write requests from clients.

The **Followers** are the other replicas that passively replicate data from the Leader. They don't serve client requests natively (though recent Kafka versions support 'follower fetching' to save cross-zone networking costs).

If the Leader fails, one of the Followers is automatically elected as the new Leader."

#### Indepth
Followers act as consumers to the Leader, issuing FetchRequests to pull new messages. By having a single Leader route all writes and reads for a partition, Kafka easily guarantees strict message ordering within that partition.

---

### 15. How does Kafka ensure fault tolerance?
"Kafka ensures fault tolerance primarily through **Replication**.

Every partition can be replicated across multiple brokers. If a broker goes down, the partitions it was leading are smoothly failed over to Follower replicas on healthy brokers.

Additionally, producers can configure `acks=all` to ensure that a message is only considered 'written' when the Leader and all In-Sync Replicas (ISR) have successfully written it to their logs."

#### Indepth
Fault tolerance also extends to rack awareness. If you configure `broker.rack`, Kafka's replica placement algorithm ensures that replicas for the same partition are distributed across different physical racks or availability zones, surviving data center-level outtages.

---

### 16. How does Kafka maintain ordering?
"Kafka guarantees strict message ordering **only within a single partition**.

If I need to guarantee order for a specific entity—like all events for 'User A'—I must ensure they map to the same partition. I do this by using the user's ID as the **Message Key** when producing. Kafka hashes the key and always routes it to the same partition, guaranteeing ordered consumption."

#### Indepth
If no key is provided, the producer uses a sticky round-robin strategy, distributing messages across all partitions evenly. This maximizes throughput but abandons ordering. Global ordering across an entire topic is only possible if the topic has exactly 1 partition, which destroys horizontal scalability.

---

### 17. Explain partitioning strategy.
"When a producer sends a message, it needs to decide which partition it goes to.

1. **Keyed Messages**: If a key is present (e.g., `user_id`), Kafka hashes the key and modulo divides it by the number of partitions. Same key always equals same partition.
2. **Null Keys**: If no key is provided, the producer uses a 'sticky partitioner'. It sticks to one partition for a batch to optimize compression and throughput, then switches to another partition.
3. **Custom Partitioner**: I can write a custom partitioner if I need domain-specific routing logic (e.g., routing high-tier customers to dedicated partitions)."

#### Indepth
A common pitfall is the 'Hot Partition' problem. If one specific key (like a celebrity user ID on a social network) generates disproportionate traffic, that single partition (and its broker) becomes a bottleneck. To solve this, you might need to append random numbers or salts to the keys of high-volume entities.

---

### 18. What happens when a broker fails?
"When a broker fails, ZooKeeper (or the KRaft controller quorum) detects that the broker's heartbeat is gone.

The Controller immediately identifies all partitions where the failed broker was the Leader. It selects a new Leader for each from the remaining In-Sync Replicas (ISR). Clients (Producers and Consumers) request a metadata update, discover the new Leaders, and seamlessly redirect their traffic."

#### Indepth
This failover takes milliseconds to seconds. The failed broker's partitions are now "under-replicated". If the broker restarts, it fetches the missed messages from the current Leaders, rebuilds its state, joins the ISR, and eventually reclaims Leadership if `auto.leader.rebalance.enable` is active.

---

### 19. How does rebalancing work?
"Rebalancing is the process where Kafka reassigns partitions among consumers within a Consumer Group.

It happens when a consumer joins the group, leaves the group, crashes, or if partitions are added to the topic. The group coordinator (a selected broker) pauses consumption, revokes partitions from existing consumers, recalculates the assignments, and distributes them fairly.

While it guarantees processing fairness and high availability, rebalancing temporarily pauses message consumption, causing latency."

#### Indepth
Modern Kafka versions (2.4+) use **Cooperative Rebalancing** (Incremental rebalancing). Instead of a "stop-the-world" approach where all consumers lose all partitions, consumers only give up the specific partitions that need to be migrated, significantly reducing rebalance latency spikes in production.

---

### 20. What is idempotent producer?
"An idempotent producer ensures that even if a producer sends the same message multiple times (due to network retries), it is only written to the Kafka partition **exactly once**.

I enable it by setting `enable.idempotence=true` in the producer properties. It assigns a unique Producer ID (PID) and a sequence number to every message. If the broker sees a duplicate sequence number for that PID, it ignores the retry."

#### Indepth
Before idempotence, if a producer sent a message, the broker wrote it, but the ACK was lost in the network, the producer would retry, leading to duplicate records in the log. Idempotence solves this gracefully at the broker level without requiring deduplication logic on the consumer side.

---

### 21. What is exactly-once semantics?
"Exactly-Once Semantics (EOS) guarantees that a message is processed and its effects are reflected exactly once, even in the event of producer retries, consumer crashes, or broker restarts.

In Kafka, this specifically refers to Kafka Streams or a consuming-and-producing loop. It ensures that the offset commit of the consumer, the state store updates, and the result publishing of the producer are executed as a single atomic **Transaction**."

#### Indepth
This is managed using Kafka's Transactional API. The producer initiates a transaction, sends messages, sends the consumer's offset to the `__consumer_offsets` topic within the transaction, and commits. If any step fails, the transaction aborts, and a downstream consumer configured with `isolation.level=read_committed` will never see those aborted messages.

---

### 22. Difference between at-least-once and exactly-once?
"In **At-Least-Once**, a system guarantees that no data is lost, but network retries might result in duplicate messages being processed. If I crash after processing but prior to committing my offset, I will process the message again upon restart.

In **Exactly-Once**, the system guarantees the outcome is applied only one time. It eliminates duplicates. It is significantly more complex and resource-intensive, requiring transactional coordination, whereas at-least-once just requires retries and standard offset commits."

#### Indepth
I usually default to At-Least-Once because it's highly performant. If I need Exactly-Once for things like financial transactions, I either rely on Kafka Transactions (for Kafka-to-Kafka pipelines) or design my external database logic to be idempotent (e.g., using `UPSERT` in PostgreSQL) effectively turning At-Least-Once into Exactly-Once.

---

### 23. How does Kafka achieve high throughput?
"Kafka achieves extreme throughput primarily through:

1. **Sequential I/O**: Instead of random disk access, it appends data sequentially to files, which modern HDDs and SSDs can do incredibly fast.
2. **Zero-Copy**: It bypasses user-space entirely when sending messages to consumers, moving bytes directly from the OS Page Cache to the network socket.
3. **Batching**: Producers and consumers heavily batch data, reducing network round-trips and CPU overhead.
4. **Partitioning**: Expanding topics into partitions allows massive horizontal scalability across multiple brokers and consumers."

#### Indepth
Kafka delegates memory management to the OS Page Cache rather than the JVM heap. This avoids massive Garbage Collection pauses. The data format on disk is exactly the same as over the network, which enables the `sendfile()` system call (Zero-Copy), drastically reducing context switches.

---

### 24. What is batching in Kafka?
"Batching is the practice of grouping multiple messages together before sending them over the network.

Instead of sending 100 messages individually (which triggers 100 network requests), a producer will hold messages briefly (configured by `linger.ms` and `batch.size`) and send them as a single compressed packet.

This drastically reduces network overhead, increases compression ratios, and maximizes the throughput the broker can handle."

#### Indepth
Batching slightly increases latency (e.g., waiting 5 milliseconds to form a batch) but exponentially increases throughput. This is the classic latency vs throughput tradeoff. Consumers also fetch in batches (`fetch.min.bytes`), ensuring they process chunks of data efficiently.
