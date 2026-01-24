# Messaging Interview Questions (RabbitMQ, Kafka, ActiveMQ)

## 🔹 1. RabbitMQ (Questions 1-20)

**Q1: What is RabbitMQ?**
RabbitMQ is an open-source message broker software that implements the Advanced Message Queuing Protocol (AMQP). It accepts messages from producers and delivers them to consumers.

**Q2: What is the role of the Exchange in RabbitMQ?**
An Exchange is a message routing agent. It takes messages from the producer and routes them to queues based on routing keys and bindings.

**Q3: What are the main components of RabbitMQ?**
Producer (sends message), Exchange (routes message), Queue (stores message), Consumer (receives message), and Binding (links Exchange to Queue).

**Q4: Name the different types of Exchanges in RabbitMQ.**
Direct, Topic, Fanout, and Headers.

**Q5: Explain the Direct Exchange.**
Routes messages to a queue if the routing key exactly matches the queue binding key.

**Q6: Explain the Topic Exchange.**
Routes messages to queues based on wildcard matches between the routing key and the binding pattern (e.g., `audit.*` or `*.error`).

**Q7: Explain the Fanout Exchange.**
Broadcasts all messages it receives to all queues bound to it, ignoring routing keys.

**Q8: Explain the Headers Exchange.**
Routes messages based on header values (metadata) instead of routing keys. It can match on `any` or `all` headers.

**Q9: What is a Binding in RabbitMQ?**
A link between an association exchange and a queue. It defines the rules (like routing keys) for how messages move from exchange to queue.

**Q10: What is a Dead Letter Exchange (DLX)?**
An exchange where messages are redirected if they cannot be delivered (e.g., queue full, message rejected, TTL expired).

**Q11: How does RabbitMQ handle message durability?**
By marking queues as `durable` and messages as `persistent` (delivery mode 2), RabbitMQ writes them to disk to survive broker restarts.

**Q12: What is the difference between transient and durable queues?**
Durable queues survive a broker restart (metadata is saved). Transient queues do not. Note: Durability doesn't automatically mean messages are persistent.

**Q13: How can you ensure message delivery in RabbitMQ?**
Use Publisher Confirms (to know broker received message) and Consumer Acknowledgments (to know consumer processed message).

**Q14: What are consumer acknowledgments?**
A signal sent by the consumer to RabbitMQ indicating that a message has been successfully processed/received, allowing RabbitMQ to delete it.

**Q15: What is prefetch count in RabbitMQ?**
A setting that limits the number of unacknowledged messages a consumer can hold at once. Helps in load balancing.

**Q16: How does RabbitMQ support clustering?**
RabbitMQ nodes can be clustered to share state (exchanges, queues, bindings). Queues are typically located on one node but visible to all (unless mirrored).

**Q17: What is VHost (Virtual Host) in RabbitMQ?**
A logical grouping of connections, exchanges, queues, and permissions. It provides isolation (like Namespaces in Kubernetes).

**Q18: What is the Erlang Cookie?**
A shared secret string used for authentication between RabbitMQ nodes in a cluster. All nodes must have the same cookie.

**Q19: How do you monitor RabbitMQ?**
Using the RabbitMQ Management Plugin (Web UI), `rabbitmqctl` CLI, or external tools like Prometheus and Grafana.

**Q20: Compare RabbitMQ with Kafka.**
RabbitMQ is a smart-broker/dumb-consumer model (push-based, complex routing). Kafka is a dumb-broker/smart-consumer model (pull-based, log storage, high throughput).

---

## 🔹 2. Apache Kafka (Questions 21-40)

**Q21: What is Apache Kafka?**
Kafka is a distributed event streaming platform used for high-performance data pipelines, streaming analytics, and data integration.

**Q22: What are the main components of Kafka?**
Producer, Broker, Topic, Partition, Consumer, Consumer Group, and Zookeeper (or KRaft controller).

**Q23: What is a Topic in Kafka?**
A category or feed name to which data is sent. Topics are partitioned and replicated.

**Q24: What are Partitions in Kafka?**
A topic is split into partitions to allow parallelism. Each partition is an ordered, immutable sequence of records.

**Q25: Explain the role of a Broker.**
A single Kafka server. It receives messages from producers, assigns offsets, commits to disk, and serves data to consumers.

**Q26: What is the role of Zookeeper in Kafka?**
Used for managing cluster metadata, leader election, and health checks. (Note: Being replaced by KRaft mode in newer versions).

**Q27: What is a Producer in Kafka?**
The application that publishes (writes) data to Kafka topics.

**Q28: What is a Consumer Group?**
A set of consumers that work together to consume a topic. Each partition is consumed by only one consumer within the group.

**Q29: How does Kafka ensure message ordering?**
Ordering is guaranteed *only within a partition*. There is no global ordering across the entire topic.

**Q30: What is an Offset?**
A unique integer ID assigned to each message in a partition, identifying its position in the log.

**Q31: What is Replication Factor?**
The number of copies of data for a topic. E.g., RF=3 means 1 leader and 2 followers for fault tolerance.

**Q32: Leader vs Follower in Kafka partitions?**
One broker is the Leader for a partition (handles all reads/writes). Others are Followers (replicate data). If Leader fails, a Follower becomes Leader.

**Q33: What is ISR (In-Sync Replicas)?**
The set of replicas that are fully caught up with the leader. Only an ISR member can be elected as a new leader.

**Q34: How does Kafka handle message retention?**
Based on time (e.g., 7 days) or size (e.g., 100GB). Old messages are deleted/compacted regardless of whether they were consumed.

**Q35: Delivery Semantics: At-least-once vs At-most-once vs Exactly-once?**
At-most-once: Message might be lost, never redelivered. At-least-once: Message never lost, might be redelivered. Exactly-once: Each message delivered exactly once (transactional).

**Q36: What is Kafka Connect?**
A framework for connecting Kafka with external systems (databases, key-value stores) via Source and Sink connectors.

**Q37: What is Kafka Streams?**
A client library for building applications and microservices, where the input and output data are stored in Kafka clusters.

**Q38: How does Kafka achieve high throughput?**
Sequential I/O (append-only logs), Zero Copy principle (sendfile), batching messages, and data compression.

**Q39: What is Log Compaction?**
A retention policy where Kafka keeps only the last known value for each message key, removing older updates for that key.

**Q40: Compare Kafka with RabbitMQ.**
Kafka: High throughput, log-based, pull-based, persistent. RabbitMQ: Low latency, queue-based, push-based, complex routing.

---

## 🔹 3. ActiveMQ (Questions 41-60)

**Q41: What is ActiveMQ?**
ActiveMQ is an open-source, multi-protocol, Java-based message broker that supports JMS (Java Message Service) widely.

**Q42: What are the key features of ActiveMQ?**
JMS compliance, multiple protocols (AMQP, MQTT, STOMP), clustering/failover, and enterprise integration patterns.

**Q43: Difference between Queue and Topic in ActiveMQ?**
Queue: Point-to-Point (one consumer gets the msg). Topic: Pub-Sub (all subscribers get a copy of the msg).

**Q44: What protocols does ActiveMQ support?**
OpenWire (native), AMQP, MQTT, STOMP, REST, and WebSocket.

**Q45: How does ActiveMQ handle persistence?**
It stores messages in a persistent store (KahaDB, JDBC, or LevelDB) to prevent data loss on failure.

**Q46: What is KahaDB?**
The default file-based persistence adapter for ActiveMQ Classic. It is optimized for high performance.

**Q47: Persistent vs Non-Persistent Delivery?**
Persistent: Written to disk (slower, safe). Non-Persistent: Held in memory (faster, lost on restart).

**Q48: What is a Dead Letter Queue (DLQ) in ActiveMQ?**
A special queue (`ActiveMQ.DLQ`) where undeliverable messages are sent after max redelivery attempts.

**Q49: How does ActiveMQ support failover?**
Using the `failover://` transport in the connection URL, allowing clients to automatically reconnect to another broker in the cluster.

**Q50: What is the Network of Brokers?**
A configuration where multiple ActiveMQ brokers are interconnected, forwarding messages to where the consumers are.

**Q51: What is an Advisory Message?**
System messages generated by ActiveMQ on a special topic to notify about events (e.g., consumer joined, queue created).

**Q52: How do you secure ActiveMQ?**
JAAS (Java Authentication and Authorization Service) for authentication, TLS/SSL for encryption, and limiting access to queues/topics.

**Q53: Difference between ActiveMQ Classic and ActiveMQ Artemis?**
Classic: Mature, stable, JMS 1.1. Artemis: Next-gen high-performance, non-blocking architecture, JMS 2.0, meant to replace Classic eventually.

**Q54: What is Consumer Priority?**
Allows high-priority consumers to receive messages before lower-priority consumers on the same queue.

**Q55: What is Slow Consumer handling?**
Strategies to handle slow consumers (e.g., discarding old messages, aborting connection) so they don't block the broker/producer.

**Q56: Can ActiveMQ be embedded?**
Yes, ActiveMQ is written in Java and can be embedded directly into a standard Java application or Spring Boot app.

**Q57: What is a Transport Connector?**
It defines how clients/brokers connect to ActiveMQ (e.g., `tcp://`, `ssl://`, `nio://`, `stomp://`).

**Q58: Compare ActiveMQ with RabbitMQ.**
ActiveMQ is JMS-centric and great for Java ecosystems. RabbitMQ is protocol-agnostic (AMQP core) and often faster for general routing.

**Q59: What are Durable Subscriptions (Topics)?**
A subscription that persists when the client disconnects. When the client reconnects, it receives messages sent while it was offline.

**Q60: How does ActiveMQ handle transactions?**
Supports JMS transactions (local) and XA transactions (distributed), ensuring atomic commit/rollback of message send/receive.
