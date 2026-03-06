# 🌐 Kafka — MirrorMaker 2 & Multi-Region Replication

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Uber, Amazon, Hotstar

---

## Q1. What is MirrorMaker 2 and how does it differ from MirrorMaker 1?

"**MirrorMaker 1 (MM1)** was a simple consumer-producer bridge. It consumed from source cluster topics and published to destination cluster topics. It had critical limitations:
- No consumer group offset translation (consumers couldn't failover seamlessly to DR cluster)
- No bidirectional or active-active support
- No automatic topic configuration replication
- No heartbeat / monitoring topics

**MirrorMaker 2 (MM2)** is a complete replication framework built on **Kafka Connect**, introduced in Kafka 2.4.

MM2 runs as a set of Kafka Connect connectors:
1. **MirrorSourceConnector** — replicates topic data
2. **MirrorCheckpointConnector** — translates and syncs consumer group offsets between clusters
3. **MirrorHeartbeatConnector** — publishes heartbeat events to measure replication lag

**Active cluster → Passive cluster example:**
```text
Primary (us-east-1) ──[MM2 MirrorSourceConnector]──▶ DR (us-west-2)
Primary (us-east-1) ◀─[MirrorCheckpointConnector]─── DR (us-west-2)
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** MirrorMaker 2 operates entirely independently as a Kafka Connect worker cluster. Neither Spring nor Go applications have any awareness of MM2's existence; they simply connect to the endpoints provided by whichever cluster (Active or DR) they are routed to.

#### Indepth
MM2 uses the Kafka Connect framework, meaning it benefits from Connect's distributed mode, offset management, and REST API for configuration updates without restarts. MM1 required a manual restart for any configuration change — a major operational pain in production.

---

## Q2. Walk through a complete MirrorMaker 2 configuration for Active-Passive DR.

"**Architecture:**
```text
Primary: kafka-primary:9092 (us-east-1)
DR:      kafka-dr:9092      (us-west-2)

MM2 runs on 3 dedicated nodes co-located with DR cluster.
```

**`mm2.properties` — Full Configuration:**
```properties
# Define the two clusters
clusters = primary, dr

# Cluster connection details
primary.bootstrap.servers = kafka-primary-1:9092,kafka-primary-2:9092,kafka-primary-3:9092
dr.bootstrap.servers     = kafka-dr-1:9092,kafka-dr-2:9092,kafka-dr-3:9092

# Enable replication: primary → dr
primary->dr.enabled = true

# Which topics to replicate (regex)
primary->dr.topics = .*
primary->dr.topics.exclude = __consumer_offsets,__transaction_state,mm2.*,heartbeats

# Offset syncing for consumer group failover
primary->dr.sync.group.offsets.enabled = true
primary->dr.sync.group.offsets.interval.seconds = 60

# Topic config replication (partition count, retention, etc.)
primary->dr.sync.topic.configs.enabled = true
primary->dr.sync.topic.configs.interval.seconds = 300

# Heartbeat for monitoring
primary->dr.emit.heartbeats.enabled = true
primary->dr.emit.heartbeats.interval.seconds = 5
primary->dr.emit.checkpoints.enabled = true
primary->dr.emit.checkpoints.interval.seconds = 60

# Replication factor for internal MM2 topics on each cluster
replication.factor = 3

# Number of replication tasks (higher = more throughput)
tasks.max = 8

# Topic naming on DR cluster: primary.original-topic-name
replication.policy.class = org.apache.kafka.connect.mirror.DefaultReplicationPolicy

# Consumer config for reading from primary
primary.consumer.fetch.min.bytes = 131072
primary.consumer.fetch.max.wait.ms = 500

# Producer config for writing to DR
dr.producer.acks = all
dr.producer.compression.type = lz4
dr.producer.linger.ms = 10
dr.producer.batch.size = 524288
```

**Start MM2:**
```bash
# Distributed mode (preferred for production)
connect-mirror-maker.sh mm2.properties

# Or as a Kafka Connect worker
connect-distributed.sh \
  --bootstrap-server kafka-dr-1:9092 \
  worker.properties mm2.properties
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** MM2 is configured purely via `.properties` files or REST API JSON payloads submitted to the Connect cluster. There is no language-specific SDK implementation required.

#### Indepth
**Topic Naming Convention:** By default, MM2 renames topics on the DR cluster from `orders` to `primary.orders`. This prevents topic collisions in active-active setups and clearly identifies the source cluster. The `RemoteTopicNameTranslator` in consumers automatically strips the prefix when needed.

---

## Q3. How does MirrorMaker 2 handle consumer group offset translation for seamless failover?

"This is the most critical MM2 feature for zero-data-loss DR failover.

**The Problem Without MM2:**
Consumer Group `payment-service` has committed offset 5000 on the Primary cluster.
The DR cluster is a copy of the Primary — its topic also has offset 5000.
But the consumer's offset 5000 is stored in the Primary's `__consumer_offsets` topic— the DR cluster has no record of it.
On failover, the consumer resets to `auto.offset.reset=earliest` and reprocesses millions of messages → **catastrophic duplicate processing**.

**MM2's Solution — MirrorCheckpointConnector:**
MM2 continuously reads consumer group offsets from the Primary cluster and **translates** them to equivalent offsets on the DR cluster, publishing them into the DR's internal offset topic.

```text
Primary: payment-service committed offset 5000 on primary.orders-partition-3
MM2 translates: offset 5000 on primary = offset 4998 on DR (slight lag)
MM2 writes: payment-service → DR.orders-partition-3 → offset 4998
```

**Consumer Failover Code:**
```java
// On startup, check if DR has a translated offset from MM2
RemoteClusterUtils.translateOffsets(
    adminClient,         // DR cluster admin
    "primary",           // source cluster alias
    "payment-service",   // consumer group
    Duration.ofSeconds(30)
).thenAccept(offsets -> {
    consumer.assign(offsets.keySet());
    offsets.forEach(consumer::seek);  // seek to translated offsets
    // Now consume from exactly where Primary left off
});
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka provides `ConsumerSeekAware` interfaces to intercept assignment events and perform manual offset translation seeks precisely where required before normal processing resumes.
* **Golang:** With `kafka-go`, manual offset management requires initializing the `Reader` without a group ID to manually fetch via `FetchMessage`, explicitly committing translated offsets against a new group using `CommitMessages`, and then restarting the main reader loop under the active group.

#### Indepth
**The Offset Gap:** MM2 translates offsets with a slight lag (default 60 seconds). On an unplanned Primary failure, consumers might reprocess 60 seconds of messages (at-least-once). For payment systems requiring exactly-once during DR, set `sync.group.offsets.interval.seconds=5` and accept the higher checkpoint traffic. Alternatively, use idempotency keys at the application level to deduplicate.

---

## Q4. What is Active-Active multi-region Kafka design? When is it used?

"**Active-Active** means both clusters are simultaneously handling producer and consumer traffic — no standby. MM2 runs bidirectionally, replicating all events across both regions.

```text
US Cluster (us-east-1) ◀══════════════▶ EU Cluster (eu-west-1)
                           MM2 (both ways)
US App writes to US Cluster               EU App writes to EU Cluster
US App reads from US Cluster              EU App reads from EU Cluster
All writes eventually replicated globally
```

**Configuration for bidirectional replication:**
```properties
clusters = us, eu

us.bootstrap.servers = kafka-us-1:9092,kafka-us-2:9092
eu.bootstrap.servers = kafka-eu-1:9092,kafka-eu-2:9092

# Both directions enabled
us->eu.enabled = true
eu->us.enabled = true

us->eu.topics = .*
eu->us.topics = .*

# CRITICAL: Prevent infinite replication loops
# MM2 tracks provenance — it will NOT re-replicate a message
# that originated from the target cluster
# (handled automatically by MM2's replication policy header)
```

**When to use Active-Active:**
- Global systems where users in each region must read/write with <30ms local latency
- High availability — if one region fails, the other immediately serves all traffic
- Regulatory compliance — EU user data must first land on EU servers

**Challenges:**
1. **Conflict Resolution** — Two regions write to the same key simultaneously. No built-in CRDT or conflict resolution in Kafka. Applications must be designed to be conflict-free (e.g., use globally unique IDs, never update a record from two regions simultaneously)
2. **Replication Lag** — Producers in the US region will see EU events with 50–200ms lag depending on cross-region bandwidth
3. **Cost** — Cross-region Kafka traffic incurs significant cloud egress fees ($0.08/GB on AWS)"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Active-Active conflict resolution requires applications to inject region or instance identifiers into Kafka Record Headers upon production. Consumers use these headers to deduplicate or determine canonical state mutations.

#### Indepth
**Cluster Linking (Confluent):** Confluent Platform offers **Cluster Linking** as a higher-performance MM2 alternative. Instead of consumer-producer bridge, Cluster Linking replicates at the broker-internal level using the same replica fetcher protocol as inter-broker replication. This delivers lower latency and exactly-once cross-cluster replication without the consumer group offset translation complexity.

---

## Q5. How do you monitor MirrorMaker 2 replication health?

"MM2 provides three built-in monitoring mechanisms:

**1. Heartbeat Topics:**
MM2's `MirrorHeartbeatConnector` publishes a message to `heartbeats` topic every 5 seconds. MM2 on the DR side replicates these to `primary.heartbeats`. By measuring the timestamp difference between production and reception, you get **end-to-end replication lag** in milliseconds.

**2. Checkpoint Topic:**
`MirrorCheckpointConnector` publishes translated offsets to `primary.checkpoints.internal`. Query this to verify consumer group offsets are being successfully translated.

**3. JMX Metrics to Alert On:**

```bash
# Key MirrorMaker 2 JMX metrics
kafka.connect:type=MirrorSourceConnector,target=dr
  → record-count              # Total records replicated
  → record-age-ms-max         # Max age of replicated records (= replication lag!)
  → replication-latency-ms-avg  # Average producer-to-consumer latency

kafka.connect:type=MirrorCheckpointConnector
  → checkpoint-latency-ms     # Time to translate and publish offsets
```

**Alerting Rules:**
```yaml
# Alert if replication is falling behind
- alert: KafkaMM2ReplicationLag
  expr: kafka_mirror_record_age_ms_max > 30000  # 30 seconds
  severity: critical
  message: "MirrorMaker2 replication lag exceeds 30s — DR cluster is stale"

- alert: KafkaMM2OffsetSyncDown
  expr: rate(kafka_mirror_checkpoint_count[5m]) == 0
  severity: warning
  message: "MirrorMaker2 offset checkpointing has stopped"
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Monitoring MM2 relies entirely on Prometheus pulling JMX from the Connect workers or analyzing the `heartbeats` topic natively via a dedicated Spring Boot or Go metrics consumer.

#### Indepth
**MM2 REST API for runtime management:**
```bash
# Check connector status
curl kafka-connect:8083/connectors/MirrorSourceConnector/status

# Pause replication (e.g., during planned maintenance)
curl -X PUT kafka-connect:8083/connectors/MirrorSourceConnector/pause

# Update config without restart
curl -X PUT kafka-connect:8083/connectors/MirrorSourceConnector/config \
  -H "Content-Type: application/json" \
  -d '{"topics": "orders,payments,users", "tasks.max": "12"}'
```

---

## Q6. How do you perform a planned failover with MirrorMaker 2?

"**Planned Failover (e.g., Regional Maintenance) — Step by Step:**

**Step 1 — Wait for Replication Lag to reach 0:**
```bash
# Monitor until record-age-ms-max ≈ 0
watch -n5 "kafka-consumer-groups.sh --bootstrap-server kafka-dr:9092 \
  --describe --group mm2-MirrorSourceConnector"
```

**Step 2 — Stop producers on Primary cluster:**
Deploy config flag or circuit breaker that stops all applications from writing to Primary.

**Step 3 — Verify DR is fully caught up:**
Confirm all MM2 `record-age-ms-max` metrics are near 0. The DR offset is exactly equal to Primary.

**Step 4 — Redirect consumers to DR:**
Update consumer `bootstrap.servers` to point to DR cluster. Use the translated offsets that MM2 already synced:
```java
// Consumers automatically pick up MM2-translated offsets on DR
// No manual offset manipulation needed if sync.group.offsets.enabled=true
```

**Step 5 — Redirect producers to DR:**
Update producer `bootstrap.servers`. All new writes go to DR cluster.

**Step 6 — Stop MM2 (optional):**
During maintenance, stop MM2 to avoid unnecessary replication from a stalled Primary.

**Step 7 — Reverse replication after maintenance:**
When Primary is back, start MM2 in reverse (`dr→primary.enabled=true`) to catch Primary up. Once lag is 0, failback.

**RTO/RPO:**
- Planned failover: RPO = 0 (zero data loss), RTO = 5–15 minutes (config propagation time)
- Unplanned failover: RPO = replication lag at time of failure (typically 60 seconds with default config)"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Failover is typically handled at the infrastructure layer (DNS swap) rather than application layer. If applications must failover manually, properties (`bootstap.servers`) should be dynamically injected via Spring Cloud Config or Go Viper without requiring code recompilation.
