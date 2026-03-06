# 🏗️ Kafka — Advanced Architecture Patterns: Event Sourcing, Multi-Tenancy & Ecosystem Comparison

> **Level:** 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Uber, Google, Hotstar

---

## Q1. How do you use Kafka as an Event Store for Event Sourcing? What are the limitations?

"**Event Sourcing** is an architectural pattern where the application state is derived by replaying a sequence of immutable events — rather than stored as a current-value snapshot in a relational database.

**Kafka as the Event Store:**
Kafka's append-only, immutable log maps perfectly to Event Sourcing semantics. Each command (e.g., `PlaceOrder`) is processed and results in one or more events (e.g., `OrderPlaced`, `InventoryReserved`) persisted as Kafka messages. To reconstruct entity state, a consumer replays all events for a given key from offset 0.

**Required Kafka Configuration for a True Event Store:**
```bash
# Indefinite retention — never delete events
kafka-configs.sh --alter --topic order-events \
  --add-config 'retention.ms=-1'

# Log compaction keeps only the latest per key (for snapshot behavior)
kafka-configs.sh --alter --topic order-snapshots \
  --add-config 'cleanup.policy=compact'
```

**Snapshot Pattern (for performance):**
Replaying 5 million events to reconstruct order state on every read is impractical. The solution is periodic **snapshots**: after every N events, the current aggregate state is written to a separate compacted topic. On startup, the consumer reads the latest snapshot, then replays only events since the snapshot offset.

**Limitations of using pure Kafka as an Event Store:**
1. No native query capability — you cannot query 'all orders in state PENDING' without a read-side projection (CQRS query model in Elasticsearch/Postgres).
2. No event schema evolution tooling built-in (need Schema Registry).
3. No optimistic concurrency control — two concurrent commands can produce conflicting events without application-level locking."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Data abstractions typically interface tightly with CQRS persistence layers, utilizing Axon Framework to combine domain models explicitly onto Kafka channels for Event Sourcing.
* **Golang:** The simplicity of Go channels and background consumers effortlessly pipelines Kafka records into fast, thread-safe memory projection maps, establishing highly optimized Read-Models localized beside the business logic.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Netflix, Hotstar — event sourcing is a dominant pattern in their catalog and viewing history systems where auditability and temporal state queries ('what was the user's state at T-30 days') are product requirements.

#### Indepth
**CQRS + Kafka + Kafka Streams Pipeline:**
```
Commands → OrderService → Kafka (order-events, retention=infinite)
                                ↓
                     Kafka Streams Processor
                     (builds materialized views)
                                ↓
              order-by-status topic (compacted) → Query Service (Postgres)
              order-stats topic → Analytics Dashboard (Elasticsearch)
```

---

## Q2. How would you design a multi-tenant Kafka cluster to serve 20 different product teams?

"Multi-tenancy in Kafka means sharing a single cluster infrastructure while providing isolation, resource fairness, and security boundaries between teams.

**Strategy 1 — Namespace-Based Topic Isolation:**
Enforce a naming convention: `<team>.<domain>.<event>`. Assign ACLs per-prefix:
```bash
kafka-acls.sh --add --allow-principal User:team-alpha \
  --operation All --topic alpha. --resource-pattern-type prefixed
```
This prevents Team Alpha from accidentally producing to Team Beta's topics.

**Strategy 2 — Quota Enforcement Per Client:**
Without quotas, one runaway producer can saturate broker network bandwidth for all tenants. Kafka supports client-level quotas:
```bash
# Limit team-alpha to 50MB/s produce bandwidth and 50MB/s fetch bandwidth
kafka-configs.sh --alter \
  --entity-type clients --entity-name team-alpha-producer \
  --add-config 'producer_byte_rate=52428800,consumer_byte_rate=52428800'
```

**Strategy 3 — Partition Count Governance:**
Over-partitioning multiplies broker memory usage (each partition = open file handles + memory buffers). Enforce a topic creation policy via a custom `CreateTopicPolicy` plugin that rejects topics with partitions exceeding a team-defined quota.

**Strategy 4 — Separate Clusters for Critical Workloads:**
Business-critical topics (payments, fraud detection) should NEVER share a cluster with non-critical analytics pipelines. A cluster restart or broker failure on the shared cluster would impact everything. Critical workloads get dedicated clusters with dedicated SLAs."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** As application developers, operating inside a multi-tenant cluster means injecting client credentials seamlessly based on CI/CD tokens, mapping local environments locally and utilizing strict `client.id` naming conventions in either application properties (Spring) or connection structs (Go).

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** LinkedIn, Uber — at scale, a single Kafka cluster serves hundreds of teams. Without governance, the cluster degrades into an unmanageable 'wild west' with noisy neighbors, security holes, and unpredictable performance.

#### Indepth
**Kafka Cluster Federation:** For very large organizations, a single Kafka cluster cannot scale indefinitely. The next step is cluster federation — multiple distinct Kafka clusters, each serving a business domain (payments-cluster, logistics-cluster, analytics-cluster), with MirrorMaker 2 federating events between domains as needed. This provides true blast-radius isolation.

---

## Q3. Compare Kafka, Apache Pulsar, AWS Kinesis, and RabbitMQ — when would you choose each?

"Each system is optimized for a different operational profile:

| Dimension | Kafka | Apache Pulsar | AWS Kinesis | RabbitMQ |
|---|---|---|---|---|
| **Architecture** | Log-based, broker stores data | Tiered Storage (broker + BookKeeper) | Log-based, serverless | Traditional message broker |
| **Throughput** | Exceptional (millions/sec) | Comparable to Kafka | Good but cost-limited | Moderate |
| **Geo-Replication** | MirrorMaker 2 (complex to operate) | Native, built-in | Cross-region replication available | Federation plugin |
| **Operational Cost** | High (self-managed) or Confluent ₹ | High (two systems: Pulsar + BookKeeper) | Zero Ops (fully managed) | Low-Medium |
| **Message TTL per Subscription** | No (global retention only) | Yes per subscription | No | Yes per queue |
| **Best For** | High-throughput event streaming, Event Sourcing | Multi-tenancy, geo-replication, infinite storage tier | Fully serverless AWS-native apps | Task queues, complex routing, RPC patterns |

**Decision Heuristic:**
- **Building in AWS and want zero Kafka ops overhead?** → Kinesis or MSK (Managed Kafka)
- **Need infinite storage + native multi-region at a startup?** → Pulsar (Aiven/StreamNative managed)
- **Industry-standard, large ecosystem, maximum control?** → Kafka
- **Simple task queue, RPC reply patterns, complex routing?** → RabbitMQ"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring natively abstracts almost all of these via `Spring Cloud Stream`, offering bindings where you can swap underlying brokers (Rabbit vs Kafka) strictly with yaml swaps, although native optimizations are lost in abstraction.
* **Golang:** Because Go leans into lean execution sizes rather than heavy vendor-abstractions like Spring Cloud Stream, porting between queues explicitly means rewriting data layers completely, cementing the importance of getting the infrastructure choices right initially.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Google, Netflix — this question tests architectural maturity and the ability to reason about trade-offs rather than defaulting to Kafka for every use case. Recommending Kafka when Kinesis would serve the purpose shows lack of cloud-native thinking.

#### Indepth
**Kafka Tiered Storage (KIP-405):** Kafka 3.6+ introduced native tiered storage, where older log segments are automatically offloaded to object stores (S3, GCS) while remaining accessible to consumers. This fundamentally changes the storage economics — infinite retention without infinite local disk cost — bringing Kafka closer to Pulsar's tiered model.
---
