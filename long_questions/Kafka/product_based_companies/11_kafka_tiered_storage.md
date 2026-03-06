# 🗄️ Kafka — Tiered Storage (KIP-405) Deep Dive

> **Level:** 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Confluent, Uber, Hotstar (large-scale data platform rounds)

---

## Q1. What is Kafka Tiered Storage and what problem does it solve?

"Traditional Kafka brokers store **all** data on local broker disks (SSDs or HDDs). This creates a fundamental cost-scaling problem.

**The core problem:** Kafka's retention policy is a business requirement — for compliance (keep 90 days), for replay (allow consumers to catch up), or for reprocessing. But local broker disk is expensive. Keeping 30TB of data on high-IOPS SSDs purely because *maybe* it'll need to be replayed is financially unjustifiable.

**Tiered Storage (KIP-405, available in Kafka 3.6+)** solves this by introducing two storage tiers:
1. **Local Tier** — Hot data on broker-local SSDs (last 1–4 hours). Serves active producers/consumers with full throughput.
2. **Remote Tier** — Cold data on cheap object storage (S3, GCS, Azure Blob). Accessible for historical reads and replay.

The split is transparent to producers and consumers — they use the exact same Kafka API regardless of which tier their data is on."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Tiered storage is 100% transparent to the client SDKs. `spring-kafka` and `kafka-go` interact with the broker exactly as they would with a purely local cluster.

#### Indepth
Before Tiered Storage, the only alternative was **Kafka with short retention + a separate archival system** (e.g., a Kafka Connect S3 Sink job). But that meant consumers couldn't use the standard `kafka-consumer-group` offset mechanism to reprocess historical data — they had to read from S3 directly via Spark or Athena. Tiered Storage unifies both into one API.

---

## Q2. How does Tiered Storage work internally?

"**Architecture:**

```text
Producer → Broker (local SSD: last 4 hours of data)
                         ↓
              [RemoteLogManager — background thread]
                         ↓
              Object Storage (S3 / GCS): full retention history

Consumer of recent data → reads from broker SSD (fast, local)
Consumer doing replay  → broker fetches from S3 transparently
```

**Step-by-step internals:**

1. **Local Write (same as always):** Producer writes to the broker's local `.log` segment files on SSD.

2. **Segment Roll:** When a local segment is closed (hits `log.segment.bytes` or `log.roll.ms`), it becomes eligible for remote upload.

3. **RemoteLogManager uploads:** A background `RemoteLogManager` thread on the broker uploads the closed segment files (`.log`, `.index`, `.timeindex`) to the remote store. It also stores remote **segment metadata** (remote offset ranges, sizes, locations) in a new internal Kafka topic: `__remote_log_metadata`.

4. **Local Deletion:** Once confirmed uploaded, the broker deletes the local segment files, freeing SSD space. Only the **local retention** window of data is kept on disk.

5. **Fetch from Remote:** When a consumer requests an offset that's in the remote tier, the broker's fetch path detects this (via the remote segment metadata), downloads the relevant segment from S3, and serves it to the consumer — completely transparent to the consumer client."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** The application merely requests a fetch at `Offset X`. The broker handles all internal logic of determining if that offset is local or remote, fetching from S3 if necessary, and serving it over the standard TCP protocol.

#### Indepth
The broker uses a **Tiered Storage Plugin API** to remain storage-agnostic. Confluent ships a production-grade S3 plugin. Netflix and LinkedIn have contributed plugins for their internal object stores. The broker code is identical regardless of the underlying remote storage backend.

---

## Q3. What is the broker configuration for Tiered Storage?

"**Enabling Tiered Storage on a broker:**

```properties
# --- Core Tiered Storage Config ---

# Enable the remote log manager
remote.log.storage.system.enable=true

# Remote storage plugin class (Confluent S3 plugin)
remote.log.storage.manager.class.name=io.confluent.kafka.tieredstorage.s3.S3RemoteStorageManager

# --- Local Retention (data kept on broker SSD) ---
# Keep only last 4 hours of data locally
log.local.retention.ms=14400000       # 4 hours
log.local.retention.bytes=-1          # no size cap on local retention

# --- Full Retention (data kept in remote storage) ---
# This is now the TOTAL retention; remote = total - local
log.retention.ms=2592000000           # 30 days
log.retention.bytes=-1

# --- S3 Plugin Config ---
remote.log.storage.manager.impl.prefix=s3.
s3.bucket.name=my-kafka-tiered-storage
s3.region=us-east-1
s3.credentials.provider.class=com.amazonaws.auth.InstanceProfileCredentialsProvider

# Number of threads for uploading segments to remote storage
remote.log.manager.thread.pool.size=10
remote.log.manager.task.interval.ms=30000
```

**Per-topic override (override the cluster default):**
```bash
# Enable tiered storage for a specific high-volume topic
kafka-topics.sh --bootstrap-server kafka:9092 \
  --alter \
  --topic app-logs \
  --config remote.storage.enable=true \
  --config local.retention.ms=3600000       # 1 hour locally
  --config retention.ms=2592000000          # 30 days total
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Configuration is strictly at the broker and topic level via `kafka-configs.sh` or infrastructure-as-code (Terraform). No client-side properties are required to support tiered storage.

#### Indepth
Tiered Storage is opt-in per topic. You would NOT enable it for low-volume operational topics like `__consumer_offsets` or `__transaction_state`. Selectively enable it only for high-volume, high-retention business event topics to maximize cost savings.

---

## Q4. What are the performance implications of Tiered Storage?

"**Active consumers (reading recent data) — ZERO impact:**
Data within the `log.local.retention.ms` window is still served from broker-local SSD and OS Page Cache. Throughput and latency are identical to a non-tiered cluster. Tiered Storage adds zero overhead to the hot path.

**Replay consumers (reading historical data) — measurable latency increase:**
Reading from S3 instead of local disk introduces:
- **S3 network fetch latency:** ~10–50ms per segment fetch (vs. ~1ms local SSD)
- **Reduced parallelism:** S3 bandwidth is shared across all replay consumers

**To mitigate replay latency, Kafka uses a Fetch Cache:**
The broker caches recently-downloaded remote segments in a local read cache (configurable size). If multiple consumers replay the same historical window simultaneously, the segments are fetched from S3 once and served to all from the cache."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** High replay latency from S3 can trigger `max.poll.interval.ms` rebalances. Ensure this value is increased if deploying massive historical replays over Tiered Storage.
* **Golang:** `kafka-go`'s `ReadBatch` blocks. Be mindful of higher latency spikes when traversing cold boundaries triggering S3 fetches behind the scenes.

#### Indepth
**Sizing the local retention correctly is critical.** If `log.local.retention.ms` is set too low (e.g., 30 minutes) and consumer lag exceeds 30 minutes (common during incidents), the consumer will start fetching from S3 even for recent data, adding latency to your operational pipeline. Always set local retention higher than your worst-case acceptable consumer lag.

---

## Q5. How much cost savings does Tiered Storage provide? Give a real-world calculation.

"**Scenario:** Netflix-style video event pipeline. 100TB/day, 7-day Kafka retention.

**Without Tiered Storage:**
```text
Total broker storage needed: 100TB × 7 days = 700TB
With replication factor 3:   700TB × 3 = 2.1 PB of SSD
AWS io1 SSD cost:            ~$0.125/GB/month
Monthly cost:                2,100,000 GB × $0.125 = $262,500/month
```

**With Tiered Storage:**
```text
Local SSD (last 4 hours):    100TB × (4/24) = ~16.7TB × 3 = 50TB SSD
SSD cost:                    50,000 GB × $0.125 = $6,250/month

Remote S3 (remaining 6.8 days): 2.1PB - 50TB = ~2.05PB S3
S3 Standard cost:            2,050,000 GB × $0.023/GB = $47,150/month

Total monthly cost:          $6,250 + $47,150 = $53,400/month
Savings:                     $262,500 - $53,400 = $209,100/month (~80% reduction)
```

**The business case is overwhelming for large-scale clusters.** At 100TB/day, Tiered Storage saves ~$2.5M/year with zero change to producer or consumer code."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Client telemetry remains exactly the same, though you may observe increased `fetch-latency` metrics in your Prometheus dashboards when reading cold data.

#### Indepth
Additional cost consideration: **S3 API call costs**. Uploading ~500 segments/day generates GET/PUT API calls (~$0.004 per 1,000 PUTs). At scale this adds a few hundred dollars monthly — negligible vs. the storage savings. However, if replay is frequent (many consumers reading cold data), S3 GET call costs can accumulate. Use **S3 Intelligent Tiering** for the remote bucket to automatically move rarely-accessed segments to Glacier-tier pricing.

---

## Q6. What are the limitations of Kafka Tiered Storage?

"**Current limitations (as of Kafka 3.6–3.9):**

| Limitation | Impact |
|---|---|
| **No compacted topics** | Tiered Storage only works with delete-retention topics. Log compacted topics (e.g., user state KTable changelogs) cannot use tiered storage. |
| **Eventual segment availability** | After a segment is closed, there's a `remote.log.manager.task.interval.ms` delay (default 30s) before it's uploaded. During this window, the data exists only locally. |
| **Remote fetch adds latency for lagging consumers** | If consumer lag > local retention, production consumer groups start hitting S3, adding 10–100ms latency to fetch responses. |
| **Plugin maturity** | The Remote Storage Plugin API is stable but third-party plugin quality varies. Confluent's S3 plugin is production-grade; community plugins for other stores need validation. |
| **No cross-region object storage support natively** | If brokers are in us-east-1 but you configure an eu-west-1 S3 bucket, cross-region egress fees will eliminate all cost savings. |"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Because compacted topics are not supported, applications relying on `KTable` (Java) or in-memory map state reconstruction (Go) must continue utilizing local SSD retention.

#### Indepth
**KRaft is required for full Tiered Storage support.** ZooKeeper-based clusters have limited Tiered Storage support in early versions. For production adoption of Tiered Storage, migrating to KRaft mode first is strongly recommended — which is aligned with Kafka's long-term direction anyway (ZooKeeper is deprecated as of Kafka 4.0).
