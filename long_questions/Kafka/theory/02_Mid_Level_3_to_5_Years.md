# 🟡 **25–47: Mid-Level (3-5 Years)**

### 1. How do you tune Kafka performance?
"Tuning Kafka involves tradeoffs between throughput, latency, and durability. I usually look at producer batching (`batch.size`, `linger.ms`), consumer fetching (`fetch.min.bytes`), and broker configuration (`num.io.threads`, `num.network.threads`).

For higher throughput, I increase `batch.size` and `linger.ms` inside the producer. For lower latency, I decrease `linger.ms` but at the cost of higher network overhead and CPU utilization. I also ensure topics have enough partitions to distribute the load across my consumer group."

#### Indepth
Disk I/O and network bandwidth are the primary bottlenecks in Kafka. Ensuring network interfaces are 10Gbps+ and using fast NVMe SSDs usually yields massive gains. OS-level tuning (like avoiding swapping and modifying `vm.swappiness=1` or adjusting TCP buffers) is also heavily utilized.

---

### 2. What are acks in producer?
"The `acks` setting dictates how many acknowledgments the broker must receive before considering an event 'committed'.

- `acks=0`: Producer fires and forgets. Lowest latency, highest risk of data loss.
- `acks=1`: Producer waits for the Leader to write the message locally. Medium latency, medium risk (data loss happens if Leader crashes before replicating).
- `acks=all` (or `-1`): Producer waits for the Leader and all In-Sync Replicas (ISR) to acknowledge. Highest latency, strongest durability."

#### Indepth
In modern event-driven applications, `acks=all` with `min.insync.replicas=2` (on a 3-broker cluster) is the industry standard for production. If you use `acks=1`, and the Leader acknowledges but immediately dies, the active consumers are suddenly seeing inconsistent data when a Follower becomes Leader.

---

### 3. How to secure Kafka?
"Kafka security rests on three pillars: Encryption (SSL/TLS), Authentication (SASL or mTLS), and Authorization (ACLs).

I encrypt data in transit using SSL so sniffers can't read the payloads. For Authentication, I restrict who can connect entirely using SASL (like Kerberos, SCRAM, or OAUTHBEARER). Finally, I use ACLs to enforce fine-grained access control (e.g., 'User A' can only `Read` from 'Topic X', not `Write`)."

#### Indepth
Implementing security imposes a performance hit. SSL decryption adds CPU overhead onto brokers. In zero-trust networks, you must also secure inter-broker communication and ZooKeeper/KRaft communication, not just client-to-broker connections.

---

### 4. What is SSL/SASL in Kafka?
"**SSL/TLS (Secure Sockets Layer/Transport Layer Security)** encrypts the connection between the client and the broker, preventing man-in-the-middle attacks.

**SASL (Simple Authentication and Security Layer)** is the framework Kafka uses to authenticate who the client actually is. 

Together, SASL/SSL means 'prove who you are, and keep our conversation completely private'."

#### Indepth
Kafka supports multiple SASL mechanisms. `SASL_SCRAM` integrates nicely with Kafka's native tooling. In enterprise environments like banks, I typically use `SASL_GSSAPI` to integrate directly with Active Directory (Kerberos) for centralized identity management.

---

### 5. Explain Kafka Connect.
"Kafka Connect is a framework designed to scalably and reliably stream data between Apache Kafka and other data systems.

Instead of writing custom scripts to pull from databases or push to cloud storage, I use Connect 'Source' connectors to ingest data (like CDC from PostgreSQL using Debezium) and 'Sink' connectors to export data (like pushing Kafka topics directly into Elasticsearch or S3)."

#### Indepth
Connect runs separately from the Kafka brokers as a distributed worker cluster. It manages offset tracking, fault tolerance, and workload balancing entirely on its own. It transforms the pipeline into a configuration-only operation (via JSON files) rather than custom coding.

---

### 6. What is Kafka Streams?
"Kafka Streams is a lightweight Java and Scala client library used for building real-time applications and microservices.

It allows me to consume data from Kafka topics, transform it (filter, map, join, aggregate), and write the results back to other Kafka topics. Unlike Spark or Flink, it’s not a massive cluster framework; it’s an application library I can bundle straight into a standard Spring Boot or Go app."

#### Indepth
Kafka Streams handles complex problems out of the box like late-arriving data (using windowing), stateful processing (using local RocksDB state stores backed by internal topics), and exactly-once processing (EOS).

---

### 7. How to handle message duplication?
"Message duplication usually happens because a producer timed out waiting for an ACK and retried a send operation (`At-least-once` semantics).

I have two main approaches:
1. Enable **Idempotent Producers** (`enable.idempotence=true`), which handles deduplication at the broker level seamlessly.
2. In my consumer logic, make my database operations **idempotent**, like using UPSERT (`ON CONFLICT DO UPDATE`) in Postgres or `PUT` instead of `POST` in REST, relying on a unique message ID to prevent double-processing."

#### Indepth
Using idempotent producers is usually preferred because it simplifies the consumer and prevents duplicate data from ever taking up space on the Kafka disk. However, if duplication occurs because a consumer crashed mid-processing before committing an offset, consumer-side idempotence is mandatory.

---

### 8. How to increase partitions?
"I can increase partitions dynamically using the `kafka-topics` CLI command:
`kafka-topics.sh --alter --topic my-topic --partitions 6 --bootstrap-server localhost:9092`

However, **I can never decrease partitions** because that would require merging disparate ordered logs, which violates Kafka's core design."

#### Indepth
Increasing partitions is dangerous if your application relies on keyed messages. When partition count changes, the hash-to-partition routing formula (`hash(key) % empty_partitions`) changes. Existing keys might suddenly route to entirely new partitions, breaking strict ordering guarantees for that specific key.

---

### 9. What is log compaction?
"Normally, a log retention policy deletes old data based on time or size. **Log Compaction**, instead, retains at least the *last known value* for every individual message key within a partition.

It is incredibly useful for maintaining a snapshot of state. If I have a topic storing 'User Profiles', I don't need all 50 updates a user made over 5 years. I only need their *latest* profile to restore state if a consumer crashes."

#### Indepth
Log Compaction runs as a background thread (`LogCleaner`). It scans old log segments, identical message keys, and writes a new segment containing only the most recent offset for each key. Deletions are handled by sending a 'Tombstone' message (a message with a key and a `nil` payload).

---

### 10. How to monitor Kafka?
"I monitor Kafka by exposing its JMX metrics and scraping them using Prometheus, then visualizing them in Grafana.

The critical metrics I monitor are:
1. **Under-Replicated Partitions (URP)**: Usually indicates a dead or slow broker.
2. **Consumer Lag**: Shows how far behind a consumer is from the latest data.
3. **Bytes In/Out Rate**: To monitor network saturation.
4. **Network/IO Processor Idle Percent**: Tells me if the broker CPU is maxed out."

#### Indepth
Monitoring consumer lag isn't just a single metric; it's a combination of `LogEndOffset` (where the broker is) minus `CurrentOffset` (where the consumer is). Tools like Burrow or Kafka Exporter specialize in exporting these specific lag metrics safely without querying the brokers themselves constantly.

---

### 11. What are common Kafka exceptions?
"In my operations, I commonly encounter:
1. `OffsetOutOfRangeException`: The consumer requested an offset that no longer exists (often because data was deleted due to retention).
2. `LeaderNotAvailableException`: A broker died, and the election for the new leader hasn't finished yet.
3. `RecordTooLargeException`: The producer sent a batch larger than `max.request.size`, or the broker `message.max.bytes` rejected it.
4. `NotLeaderForPartitionException`: The client sent a request to a broker that is not currently the leader, forcing a metadata refresh."

#### Indepth
`CommitFailedException` is infamous in consumers. It happens if processing takes longer than `max.poll.interval.ms`. Kafka assumes the consumer died, kicks it out of the group, and triggers a rebalance. When the consumer finally finishes processing and tries to commit its offset, it gets rejected.

---

### 12. Design a real-time notification system using Kafka.
"I would architect this using Microservices decoupled by Kafka:
1. An upstream service (e.g., Order Service) acts as a **Producer**, writing an 'Order Shipped' event to an `orders_events` topic.
2. The Notification Service acts as a **Consumer**.
3. I would partition the topic by `user_id` so all notifications for one user are processed sequentially.
4. The Consumer reads the event, fetches user notification preferences from Redis, and fires a push notification via Firebase or SMS via Twilio."

#### Indepth
To handle failures without blocking the pipeline, I would use a Dead Letter Queue (DLQ). If an SMS fails, it gets routed to a `notifications_dlq` topic. A separate, slower consumer process reads the DLQ, applies exponential backoff, and retries the external API call, ensuring the main topic consumer isn't delayed.

---

### 13. How would you design an event-driven architecture?
"I would use Kafka as the central nervous system. 

Instead of services communicating synchronously via REST (which creates tight coupling and cascading downstream failures), services emit domain events ('UserCreated', 'PaymentProcessed') to Kafka topics. Downstream services subscribe to topics they care about. This is choreographic choreography.

This decoupling means an Email Service can go down for an hour, but the main Registration flow succeeds. When Email Service boots back up, it simply consumes the backlog of 'UserCreated' events."

#### Indepth
One massive challenge in Event-Driven Architecture is dual-writes (updating a local database AND writing to Kafka simultaneously). To solve this reliably, I employ the **Transactional Outbox Pattern**—writing the event into the local database transactionally, then using Debezium (Kafka Connect) to stream that event to Kafka.

---

### 14. How do you handle consumer lag?
"Consumer lag means my application is processing messages slower than producers are publishing them.

My first step is usually scaling horizontally: I increase the number of partitions on the topic and spin up more consumer application instances.

If I cannot increase partitions, I scale vertically: I pull messages from Kafka single-threaded but dispatch the actual processing work (like an API call) to an internal Go `worker pool` or Java `ThreadPoolExecutor`. However, this sacrifices strict ordering."

#### Indepth
Sometimes the fix is tuning the consumer configs. Increasing `fetch.min.bytes` and `fetch.max.wait.ms` tells the broker to hold requests until a larger batch is ready, which vastly improves throughput. Conversely, if my processing is CPU bound, parallelizing the workload is the only way forward.

---

### 15. How to scale Kafka cluster?
"Scaling a Kafka cluster usually means adding new brokers.

However, just deploying a new broker does nothing. Kafka dynamically does not rebalance existing partitions.

To truly scale, after adding the broker, I must execute a **Partition Reassignment**. I use tools (like Cruise Control or `kafka-reassign-partitions.sh`) to generate a JSON plan that moves heavy partitions off the old brokers onto the newly added brokers, which then stream the data to replicate it."

#### Indepth
Partition reassignment is an extremely network-intensive process. Doing it during peak hours can cause a network storm and kill cluster performance. It must always be throttled using quota limits (`leader.replication.throttled.rate`) to prevent existing vital production traffic from lagging.

---

### 16. Explain transactional APIs.
"Transactional APIs allow producers to publish messages to multiple partitions atomically. Either all messages in the transaction are committed or none are.

I initialize the producer with a `transactional.id`, call `begin()`, send messages across multiple topics, optionally send the consumer offsets, and call `commitTransaction()`. If the app crashes, the coordinator aborts it. 

This enables Exactly-Once Semantics in streaming systems."

#### Indepth
This places a considerable burden on consumers. By default, consumers read uncommitted data immediately. To use transactionality properly, consumers must explicitly set `isolation.level=read_committed`. In this mode, consumers only read up to the 'Last Stable Offset' and buffer or ignore uncommitted control records.

---

### 17. How does Kafka ensure durability?
"Kafka ensures durability primarily by saving data directly to disk (into OS page cache, eventually flushed to segments) rather than holding it in a JVM heap. 

Secondly, by leveraging multi-broker replication. Even if the disk crashes, replica followers have duplicate logs. 

Thirdly, producers using `acks=all` with a minimum in-sync replica threshold ensures that the broker doesn't claim 'success' until the data is mirrored safely across fault domains."

#### Indepth
Flushing data to disk (`fsync`) blocks the main thread. To stay blazingly fast, Kafka rarely forces `fsync`. It relies on the replication protocol instead for durability. "Data is durable because it's on 3 servers' RAM at once, not because one server flushed it to an HDD platter immediately."

---

### 18. What are controller brokers?
"In an older ZooKeeper ensemble, one Kafka broker is elected as the **Active Controller**. 

It handles all administrative actions for the entire cluster: monitoring broker failures, performing partition leader elections, and handling topic creation/deletion. The other brokers blindly follow the controller's state changes. 

In modern KRaft mode, there is a dedicated subset of brokers running entirely as 'Controllers' whose sole job is to manage cluster metadata via the Raft consensus protocol."

#### Indepth
In ZooKeeper mode, the controller election can cause latency spikes because the active controller has to fetch massive amounts of metadata from ZK and broadcast it sequentially to all brokers. Under KRaft, metadata exists natively as a Kafka topic, allowing instant, event-sourced propagation.

---

### 19. How does leader election work?
"When a broker goes down, the Active Controller detects the missing heartbeat.

For every partition where the dead broker was the Leader, the Controller looks at the current In-Sync Replicas (ISR) list. It picks the first healthy follower in the ISR and nominates it as the new Leader. 

It then updates the cluster metadata, notifying clients and remaining followers to route reads/writes to this new broker."

#### Indepth
If the ISR is empty (all replicas crashed simultaneously), Kafka halts by default to prevent data loss. However, if availability is prioritized over consistency, I can enable `unclean.leader.election.enable=true`. This permits a follower that is *out of sync* to become leader, guaranteeing the cluster stays online, but resulting in catastrophic data loss of the missed offsets.

---

### 20. How do you handle backpressure?
"Kafka inherently handles backpressure out of the box due to its **pull-based** consumer model.

In typical pub/sub like RabbitMQ, the broker pushes messages out, easily overwhelming a slow consumer. In Kafka, the consumer decides when it is ready. It issues a `Fetch` request for data. If the consumer is overwhelmed, it simply stops fetching. 

The data piles up safely on Kafka's massive broker disks instead of crashing the consumer."

#### Indepth
While Kafka handles backpressure beautifully on the consumer side, it can struggle on the producer side. If network bandwidth is saturated, `send()` requests buffer in memory until `max.block.ms` times out, potentially bringing down the producer application if the buffer memory (`buffer.memory`) gets exhausted.

---

### 21. Explain data consistency in distributed systems using Kafka.
"Kafka relies on the concept of an ordered, immutable log to act as a single source of truth across a distributed system.

When microservices need data consistency across their databases (which lack distributed transactions), I use the **Saga Pattern**. Services emit events to Kafka. If an operation spans Service A and Service B, Service A writes the first event, B processes it and updates its DB, then B emits a success event. 

If step B fails, B emits a compensating failure event, and Service A undoes its initial action, maintaining eventual consistency."

#### Indepth
This is Eventual Consistency, not immediate ACID consistency. However, because Kafka strictly orders events by key, if you configure producers securely, the eventual state across all read-models (Databases, Redis, Elasticsearch) is reliably determinable solely by replaying the Kafka commit log.

---

### 22. Compare Kafka vs Pulsar.
"Apache Pulsar was built at Yahoo directly addressing some of Kafka's operational flaws. 

The biggest difference is architecture. Kafka couples computing (brokers) and storage (disks) on the same node. Scaling means migrating terabytes of data.

Pulsar decouples them. Brokers handle routing and caching, but offload storage to **Apache BookKeeper**. Adding new storage nodes in Pulsar doesn't require massive partition reassignments, making it highly elastic."

#### Indepth
Pulsar natively multi-tenant, supports traditional message queueing seamlessly alongside streaming, and utilizes tiered storage natively. However, Kafka has the ultimate trump card: community adoption, the Kafka Streams/Connect ecosystem, and a vast pool of available talent.

---

### 23. Explain zero-copy principle in Kafka.
"The Zero-Copy principle is the main reason Kafka handles data so fast.

Normally, when a broker sends a file to a consumer network socket, the OS copies bytes from disk into OS cache, then into the JVM (User Space), then back to the OS socket buffer. That's 4 copies and 4 context switches.

Kafka avoids the JVM entirely. Because the file format on disk is fully compatible with the network protocol, Kafka uses the Linux `sendfile()` system call. Bytes route directly from the OS Page Cache into the Network Interface Card."

#### Indepth
Zero-copy drastically reduces CPU load on the brokers. It's the reason why a Kafka broker handling millions of messages can often be seen running at historically low CPU utilization, but heavily maxing out its disk I/O and network bandwidth limit.
