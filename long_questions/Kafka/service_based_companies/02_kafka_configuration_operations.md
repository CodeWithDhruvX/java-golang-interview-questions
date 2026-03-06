# 🏗️ Kafka — Service-Based Companies Configuration & Operations

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Tech Mahindra, Mindtree

---

## Q1. What is the difference between log deletion and log compaction in Kafka?

"Kafka has two fundamental retention policies to manage disk space:

**1. Delete (Default):**
When a topic is configured with `cleanup.policy=delete`, Kafka simply deletes old data based on time (e.g., `retention.ms=604800000` for 7 days) or size (e.g., `retention.bytes=1073741824` for 1GB). Once the threshold is crossed, the oldest log segments are completely purged, regardless of their content.

**2. Compact:**
When `cleanup.policy=compact` is utilized, Kafka guarantees that it will permanently retain at least the **last known value** for each specified message key. If I publish an event with the key `User123` and value `Address_A`, and later publish `User123` with value `Address_B`, the compactor thread will eventually delete the `Address_A` record. This prevents the topic from growing infinitely while ensuring consumers can always reconstruct the latest state."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Often managed declaratively using `NewTopic` bean configurations with `.config(TopicConfig.CLEANUP_POLICY_CONFIG, TopicConfig.CLEANUP_POLICY_COMPACT)`.
* **Golang:** Topic creation is generally handled via an Admin Client like `kafka.NewAdminClient()` (confluent) configuring the `cleanup.policy` in the `TopicSpecification`.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Tech Mahindra, Infosys — testing knowledge on operational data integrity.

#### Indepth
**Tombstone Messages:** To delete a key entirely in a compacted topic, a producer must send a message with the desired key but a `null` payload. This is called a tombstone. The compactor uses this signal to fully eradicate all references to that key in the next compaction pass.

---

## Q2. Why is ZooKeeper being phased out in favor of KRaft?

"Historically, Kafka relied heavily on an external distributed system called Apache ZooKeeper to manage cluster metadata, handle controller elections, and store topic configurations.

However, Kafka is actively replacing it with **KRaft (Kafka Raft Metadata mode)** for several architectural reasons:
1. **Simplified Ops:** Maintaining two separate distributed systems (Kafka and ZooKeeper) meant separate tuning, security policies, and administrative overhead. KRaft builds the raft consensus protocol directly into Kafka.
2. **Scalability Limits:** ZooKeeper struggles with metadata bloat. If a cluster had hundreds of thousands of partitions, leader election could take vital seconds, causing noticeable downtime. KRaft uses an internal metadata event log, dropping controller failover times into the milliseconds range regardless of partition count.
3. **Standalone Mode:** KRaft allows smaller teams to run Kafka on a single machine comfortably without configuring a complex quorum setup."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Both ecosystems' drivers abstract away the cluster backend. Whether the cluster runs KRaft or Zookeeper, the client library connection string just utilizes the `bootstrap.servers` list. The change is strictly transparent to application developers.

#### 🏢 Company Context
**Level:** 🟢 Junior to 🟡 Intermediate | **Asked at:** Cognizant, TCS — checking if the candidate has kept their knowledge up-to-date with modern Kafka architectures.

#### Indepth
**Metadata Quorum:** In KRaft mode, instead of all brokers connecting to ZooKeeper, a subset of the standard brokers is designated directly as 'Controller' nodes. These controllers govern the metadata amongst themselves using the Raft consensus algorithm, and the other standard 'Broker' nodes fetch metadata directly from them.

---

## Q3. How do you implement basic security and authentication in a Kafka cluster?

"Kafka provides a robust, multi-layered security model that can be implemented in production environments:

**1. Transport Encryption (SSL/TLS):**
By default, Kafka sends data in plain text, making it vulnerable to packet sniffing. We mandate SSL encryption for traffic between producers, brokers, and consumers.

**2. Client Authentication:**
There are several ways clients can authenticate to the cluster:
*   **mTLS (Mutual TLS):** The client provides a trusted certificate alongside the encryption layer.
*   **SASL/SCRAM:** Username and password-based authentication hashed natively within Kafka.
*   **SASL/OAUTHBEARER:** Integrates with corporate identity providers using OAuth2 tokens.

**3. Authorization (ACLs):**
Once authenticated, the user must have permission to act. Using Access Control Lists (ACLs), we can define rules like: *User 'App-A' is only allowed to Read from the 'orders' topic and is denied Write access completely.*"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Configured via `security.protocol=SASL_SSL` and `sasl.jaas.config` properties natively passed to the application configs. Spring makes it trivial to inject keystores for mTLS.
* **Golang:** For `confluent-kafka-go`, developers configure the config map with `security.protocol`, `sasl.mechanism`, `sasl.username`, and `sasl.password`. Setting up mTLS requires specifying `ssl.ca.location` and `ssl.certificate.location`.

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Wipro, TCS — critical for banking, healthcare, and enterprise projects where strict data compliance is enforced.
---
