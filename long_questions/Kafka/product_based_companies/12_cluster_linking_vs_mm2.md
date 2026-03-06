# 🔗 Kafka — Confluent Cluster Linking vs MirrorMaker 2

> **Level:** 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Confluent Platform users, Uber, Financial institutions with multi-region DR requirements

---

## Q1. What is Confluent Cluster Linking and how does it fundamentally differ from MirrorMaker 2?

"Both solve the same problem — replicating Kafka data across clusters or regions. But they operate at completely different layers of the Kafka architecture.

**MirrorMaker 2 (MM2) — Application-layer replication:**
MM2 is a Kafka Connect application. It **consumes** messages from the source cluster (acting like a regular consumer) and **produces** them to the destination cluster (acting like a regular producer). It's a bridge pattern running outside the broker.

```text
Source Cluster                    MirrorMaker 2                  Destination Cluster
Producer → [Broker] → Consumer → [MM2 Producer] → [Broker] → Consumer
```

**Cluster Linking — Broker-level replication:**
Cluster Linking replicates topics directly at the **broker internals level**, using the same internal replica-fetch protocol that Kafka brokers use for inter-broker ISR replication. No consumer or producer API is involved.

```text
Source Cluster                                    Destination Cluster
Producer → [Broker] ←── internal fetch protocol ─── [Broker] → Consumer
```

The destination broker literally acts as a 'remote follower' of the source topic, pulling data identically to how a local follower would pull from the leader."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** The architectural transparency means that neither Spring nor Go clients need any special dependencies. The destination broker looks identical to a standard Kafka broker.

#### Indepth
This architectural difference has cascading implications. MM2 always re-encodes messages through the producer path, which means it can't guarantee message offset preservation, and every message goes through producer batching, compression re-encoding, and the full produce request pipeline. Cluster Linking transfers raw log segments — offsets are preserved byte-for-byte.

---

## Q2. Compare MM2 and Cluster Linking head-to-head across all key dimensions.

"| Dimension | MirrorMaker 2 | Cluster Linking |
|---|---|---|
| **Mechanism** | Consumer + Producer bridge | Internal broker replica-fetch protocol |
| **Offset Preservation** | ❌ Offsets change on destination (requires MirrorCheckpointConnector translation) | ✅ Offsets are identical — byte-for-byte replication |
| **Replication Lag** | Higher — double the produce-consume path, plus Connect overhead | Lower — same protocol as ISR, broker-to-broker direct |
| **Exactly-Once** | Complex — requires transactional producers in MM2 config | ✅ Inherent — same protocol as ISR replication |
| **Setup Complexity** | Medium — Connect workers, connector configs, offset translation | Low — single command via Kafka admin API or Confluent Control Center |
| **Consumer Failover** | Requires `RemoteClusterUtils.translateOffsets()` in consumer code | ✅ Transparent — consumers switch bootstrap servers, same offsets work |
| **Schema Drift** | Can occur if serialization is misconfigured | N/A — raw bytes transferred, no re-serialization |
| **Open Source** | ✅ Yes — part of Apache Kafka | ❌ No — Confluent Platform (Enterprise) only |
| **Cost** | Free (compute cost of running Connect workers) | Confluent license cost |
| **Active-Active** | Supported (bidirectional MM2 config) | ✅ Supported (bidirectional cluster links) |
| **Supported Sources** | Any Kafka cluster | Confluent Platform or Confluent Cloud clusters only |"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Offset byte-for-byte preservation means you do NOT need to implement complex `seek` logic in Spring's `ConsumerSeekAware` or Go's `Reader.SetOffset`.

#### Indepth
The **offset preservation** difference is the most operationally significant. With MM2, consumer failover to DR requires a code-level integration with `RemoteClusterUtils.translateOffsets()`, and there is always a translation lag (default 60 seconds). Missed messages during unplanned failover are possible. With Cluster Linking, a failover is just a DNS/bootstrap-server update — consumers continue from the exact same committed offset on the destination cluster with zero code changes and zero replay risk.

---

## Q3. When would you choose MM2 over Cluster Linking, and vice versa?

"**Choose MirrorMaker 2 when:**

| Scenario | Reason |
|---|---|
| You are running Open Source Kafka (non-Confluent) | Cluster Linking requires Confluent Platform — MM2 is the open-source standard |
| Source and destination are different Kafka distributions | MM2 works between any Kafka clusters (OSS, MSK, Confluent, Strimzi) |
| You need cross-version replication (e.g., Kafka 2.6 → Kafka 3.6) | Cluster Linking has stricter version compatibility requirements |
| Budget is constrained — no Confluent license | MM2 is free |
| You need selective topic transformation during replication | MM2's Connect plugin framework allows SMTs (Single Message Transforms) to modify messages in-flight |

**Choose Cluster Linking when:**

| Scenario | Reason |
|---|---|
| You are on Confluent Platform or Confluent Cloud | Native feature, simpler ops |
| Consumer failover must be seamless with zero code changes | Offset identity preservation is the killer feature |
| You need the lowest possible replication lag | Broker-native protocol is faster than application-bridge |
| You are building a global active-active multi-region system with strict RPO/RTO | Superior consistency guarantees and simpler failover story |
| Your DR SLA requires < 10 second RPO | MM2's checkpoint lag (default 60s) makes this impossible; Cluster Linking achieves it natively |"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Using MM2 might necessitate custom SMTs (Single Message Transforms) written in Java, whereas Cluster Linking transfers raw segments without deserialization.

#### Indepth
**Real-world hybrid pattern:** Many large organizations run MM2 for **cross-company or cross-cloud replication** (where Cluster Linking isn't an option) and use Cluster Linking for **within-Confluent multi-region DR**. There is no rule that you must choose one forever — pick the right tool per replication flow.

---

## Q4. How do you set up Cluster Linking? Walk through the key commands.

"**Setup requires Confluent Platform 6.0+ (or Confluent Cloud).**

**Step 1 — Create a Cluster Link from destination cluster to source cluster:**
```bash
# On the DESTINATION cluster, create a link that points to the SOURCE
kafka-cluster-links --bootstrap-server kafka-destination:9092 \
  --create \
  --link my-dr-link \
  --config bootstrap.servers=kafka-source:9092,security.protocol=SSL,ssl.keystore.location=/etc/kafka/certs/kafka.keystore.jks,ssl.keystore.password=changeit

# Verify the link is active
kafka-cluster-links --bootstrap-server kafka-destination:9092 \
  --list
```

**Step 2 — Create a Mirror Topic on the destination (linked to source topic):**
```bash
# Mirror the 'orders' topic from source to destination
kafka-mirrors --bootstrap-server kafka-destination:9092 \
  --create \
  --mirror-topic orders \
  --link my-dr-link

# Verify mirror is replicating
kafka-mirrors --bootstrap-server kafka-destination:9092 \
  --describe \
  --topics orders
```

**Step 3 — Monitor replication lag:**
```bash
# Check mirror topic metrics
kafka-mirrors --bootstrap-server kafka-destination:9092 \
  --describe \
  --topics orders

# Output includes:
# Mirror Topic: orders
# Link Name: my-dr-link
# Source Topic: orders
# State: ACTIVE
# Replication Factor: 3
# Lag: 0 messages  ← key metric
```

**Step 4 — Perform failover (planned):**
```bash
# Stop the mirror (promotes destination topic to writable)
kafka-mirrors --bootstrap-server kafka-destination:9092 \
  --failover \
  --topics orders

# Consumers now point to kafka-destination:9092/orders
# Offsets are identical — zero consumer code changes needed
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Failover is purely an infrastructure task. Applications reboot pointing to the new `bootstrap.servers` and resume processing without issue.

#### Indepth
**Automatic failover (Confluent Cloud):** In Confluent Cloud with Dedicated clusters, you can configure **Automatic Cluster Linking failover** that triggers when source cluster health checks fail. This can reduce RTO (time to recover) from 15 minutes for a manual MM2 failover to under 2 minutes for an automated Cluster Linking failover.

---

## Q5. How does Cluster Linking handle active-active multi-region?

"**Active-Active Configuration:**

In active-active, both clusters handle live traffic. Cluster Linking runs **bidirectionally** — Link A replicates source→destination, Link B replicates destination→source.

```text
US Cluster ─────Link A (us→eu)────▶ EU Cluster
US Cluster ◀────Link B (eu→us)───── EU Cluster
```

**Preventing Infinite Replication Loops:**
Cluster Linking uses a **provenance header** appended to every replicated message. Before replicating a message, the link checks this header. If the message's origin cluster matches the target cluster, replication is skipped.

```text
US writes message (origin=US) → Link A replicates to EU
EU's Link B sees message (origin=US ≠ EU) → skips it ✅
EU writes message (origin=EU) → Link B replicates to US
US's Link A sees message (origin=EU ≠ US) → skips it ✅
```

**Conflict Resolution Challenge (same for MM2 and Cluster Linking):**
Neither tool resolves write conflicts. If a US app and an EU app both write to key `user-123` simultaneously, the destination cluster ends up with both versions without a defined winner. Application architects must design **conflict-free data models** (e.g., no two regions ever own the same key, use globally unique partition ownership).

**Practical active-active pattern — Geographic Partitioning:**
```text
US Cluster:   partitions 0–11  (US user events, keyed by US userId)
EU Cluster:   partitions 12–23 (EU user events, keyed by EU userId)
Replication: Both ways — US consumers can read all 24 partitions globally
```
This eliminates write conflicts entirely: each region only writes to its own partition range."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Conflict resolution for Active-Active often involves CRDTs (Conflict-Free Replicated Data Types) implemented manually in Go structs or Java POJOs before final local DB persistence.

#### Indepth
Confluent Cloud's **Multi-Region Clusters** feature goes beyond Cluster Linking — it runs a single Kafka cluster spanning multiple regions with **observer replicas** in remote regions. Reads can be served locally from any region with zero replication lag for reads. This is architecturally closer to CockroachDB's multi-region model than MM2 or Cluster Linking, and is the gold standard for global low-latency Kafka deployments.

---

## Q6. What is the RPO and RTO comparison between MM2 and Cluster Linking?

"| Metric | MirrorMaker 2 | Cluster Linking |
|---|---|---|
| **RPO (data loss risk) — planned failover** | 0 (if lag=0 before failover) | 0 (if lag=0 before failover) |
| **RPO — unplanned failover** | Up to 60 seconds (checkpoint sync interval) | ~5 seconds (broker-level, near-real-time sync) |
| **RTO — planned failover** | 5–15 minutes (manual steps, consumer code changes) | 2–5 minutes (failover command + bootstrap-server update only, no code change) |
| **RTO — unplanned failover** | 15–45 minutes (detect, translate offsets, redeploy) | 5–10 minutes (automated failover config) |
| **Consumer code change after failover** | Required (`translateOffsets()` call) | ❌ Not required |
| **Risk of double processing after failover** | Moderate (offset translation lag) | Very Low (offset identity) |"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Cluster Linking's ~5s RPO ensures that even in catastrophic failure, transaction compensations written in Go or Spring Sagas rarely need to invoke fallback logic across regions.

#### Indepth
For financial systems (banks, payment processors) where the cost of data loss is regulatory/legal, the ~5-second RPO of Cluster Linking vs. 60-second RPO of MM2 can be a deciding factor in architecture selection — even justifying the Confluent Platform license cost.
