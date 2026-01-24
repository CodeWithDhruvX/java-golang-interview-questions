## 🟢 Apache Kafka (Questions 21-40)

### Question 21: What is Apache Kafka?

**Answer:**
Apache Kafka is an open-source distributed event streaming platform used for high-performance data pipelines, streaming analytics, data integration, and mission-critical applications.
*   **Core Capability:** It publishes and subscribes to streams of records, similar to a message queue or enterprise messaging system, but in a fault-tolerant, durable way.

### Question 22: What are the main components of Kafka architecture?

**Answer:**
1.  **Broker:** A Kafka server. A cluster consists of multiple brokers to maintain load balance and fault tolerance.
2.  **Topic:** A category to which records are published.
3.  **Partition:** A split of a topic. This allows the topic to scale across multiple brokers.
4.  **Producer:** Applications that publish data to topics.
5.  **Consumer:** Applications that read data from topics.
6.  **Zookeeper (Legacy):** Used for metadata management and leader election (being replaced by KRaft).

### Question 23: What is a Topic in Kafka?

**Answer:**
A Topic is a logical channel to which producers send records and from which consumers read records.
*   Topics are **multi-subscriber**: A topic can have zero, one, or many consumers that subscribe to the data written to it.
*   Topics are **partitioned** and **replicated**.

### Question 24: What are Partitions in Kafka and why are they used?

**Answer:**
Partitions are the unit of parallelism in Kafka.
*   **Splitting:** A single topic is split into multiple partitions.
*   **Ordering:** Each partition is an ordered, immutable sequence of records.
*   **Parallelism:** Different partitions can be processed by different consumers in the same group simultaneously.
*   **Scalability:** Partitions allow a topic's data to be spread across multiple servers (Brokers).

### Question 25: Explain the role of a Broker in Kafka.

**Answer:**
A Kafka Broker is a server that runs the Kafka process.
*   **Responsibility:** It receives messages from producers, assigns them offsets, and commits them to storage on disk. It also services fetch requests from consumers.
*   **Leader/Follower:** A broker can be a leader for some partitions and a follower for others.

### Question 26: What is a Zookeeper in Kafka? Is it still required in newer versions (KRaft)?

**Answer:**
*   **Role:** Zookeeper tracks the status of nodes (brokers), maintains configuration information, and manages the election of partition leaders.
*   **KRaft (Kafka Raft Metadata mode):** In newer versions (Kafka 2.8+), the dependency on Zookeeper is being removed. Kafka can manage its own metadata using an internal Raft quorum, making the architecture simpler.

### Question 27: What is a Producer in Kafka?

**Answer:**
A Producer is a client application that publishes (writes) events to the Kafka system.
*   **Logic:** The producer chooses which record to assign to which partition within the topic (e.g., Round Robin, or based on the hash of a record key).
*   **Acks:** Producers can choose how many acknowledgments they need (0, 1, or all) to consider a write successful.

### Question 28: What is a Consumer and Consumer Group in Kafka?

**Answer:**
*   **Consumer:** An application that reads data from topics.
*   **Consumer Group:** A logical grouping of consumers.
    *   **Rule:** Each partition in a topic is consumed by **exactly one** consumer in the group.
    *   **Scaling:** To scale reading, you add more consumers to the group (up to the number of partitions).
    *   **Broadcasting:** If multiple groups subscribe to the same topic, each group gets a full copy of the messages.

### Question 29: How does Kafka ensure message ordering?

**Answer:**
Kafka guarantees message ordering **only within a partition**, not across the entire topic.
*   **Per Partition:** Messages are appended in order, and consumers read them in that same order (FIFO).
*   **Global Ordering:** To achieve total ordering, you would need a topic with only **one partition** (which sacrifices parallelism).

### Question 30: What is an Offset in Kafka?

**Answer:**
An offset is a unique identifier (integer) assigned to each record within a partition.
*   **Function:** It denotes the position of the record in the partition log.
*   **Tracking:** Consumers track their progress by committing the offset of the last message they processed.

### Question 31: What is Replication Factor in Kafka?

**Answer:**
The Replication Factor (RF) determines how many copies of a topic partition are kept across the cluster.
*   **Purpose:** Fault tolerance. If a broker goes down, data is available on another broker.
*   **Example:** RF=3 means there is 1 Leader and 2 Followers. The system can survive 2 broker failures.

### Question 32: What is the Leader and Follower concept in Kafka partitions?

**Answer:**
*   **Leader:** One replica is designated as the Leader. All reads and writes for that partition go to the Leader.
*   **Follower:** Other replicas are Followers. They passively replicate (copy) data from the Leader. They take over only if the Leader fails.
*   **Note:** Clients (Producers/Consumers) talk to the Leader.

### Question 33: What is ISR (In-Sync Replicas)?

**Answer:**
ISR is the set of replicas that are fully caught up with the Leader.
*   **Relevance:** Only a member of the ISR can be elected as a new Leader if the current Leader fails.
*   If a replica falls too far behind, it is removed from the ISR until it catches up.

### Question 34: How does Kafka handle message retention?

**Answer:**
Kafka persists all messages for a set period, whether or not they have been consumed.
*   **Log Retention:** Configurable by time (e.g., 7 days) or size (e.g., 100 GB).
*   **Cleanup:** Once the limit is reached, old segments are deleted or compacted to free up space.

### Question 35: What is the difference between "at least once", "at most once", and "exactly once" delivery semantics?

**Answer:**
*   **At most once:** Message might be lost, but never redelivered. (Fire and forget).
*   **At least once:** Message is never lost, but might be redelivered. (Default, good for most use cases).
*   **Exactly once:** Message is processed exactly once, even if producers retry. (Achieved using Idempotent Producers and Transactional API).

### Question 36: What is Kafka Connect?

**Answer:**
Kafka Connect is a tool for scalable and reliable streaming of data between Apache Kafka and other systems.
*   **Source Connectors:** Import data from external systems (e.g., JDBC Source for databases) into Kafka.
*   **Sink Connectors:** Export data from Kafka to external systems (e.g., Elasticsearch Sink, S3 Sink).

### Question 37: What is Kafka Streams?

**Answer:**
Kafka Streams is a client library for building applications and microservices, where the input and output data are stored in Kafka clusters.
*   **Features:** It supports stateful and stateless processing, windowing, joins, and aggregations directly within the application code purely using Kafka topics.

### Question 38: How does Kafka achieve high throughput?

**Answer:**
1.  **Sequential I/O:** Kafka writes to disk sequentially (append-only log), which is much faster than random access.
2.  **Zero Copy:** Calls `sendfile` (OS level) to transfer data from disk to network without copying it into application memory.
3.  **Batching:** Producers send messages in batches (less network overhead).
4.  **Compression:** Supports Gzip, Snappy, LZ4 (reduces bandwidth and storage).

### Question 39: What is Log Compaction in Kafka?

**Answer:**
Log compaction is a retention policy where Kafka ensures that it retains at least the last known value for each message key within the log of data for a single topic partition.
*   **Use Case:** Restoring state (e.g., a "User Profile" topic). You only care about the latest email address for `User123`, not the 50 previous changes.

### Question 40: Compare Kafka with RabbitMQ and ActiveMQ.

**Answer:**

| Feature | Apache Kafka | RabbitMQ / ActiveMQ |
| :--- | :--- | :--- |
| **Architecture** | Log-based (Distributed Log). | Queue-based (Memory Buffer). |
| **Message Life** | Persists (retention policy). | Deleted after consumption. |
| **Consumption** | Pull (Consumer requests data). | Push (Broker sends data). |
| **Scaling** | High (Partitions across brokers). | Vertical (Cluster adds availability, not throughput). |
| **Use Case** | Streaming, Event Souring, ETL. | Complex Routing, Task Processing. |
