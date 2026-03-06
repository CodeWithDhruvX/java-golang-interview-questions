# 💰 Kafka — Cost Optimization Strategies

> **Level:** 🟣 Architect
> **Asked at:** Netflix, Uber, LinkedIn, Amazon, large-scale Kafka platform teams and FinOps-aware architecture rounds

---

## Q1. What are the main cost drivers in a Kafka cluster?

"When I do a Kafka cost audit, I look at five buckets:

| Cost Driver | Typical % of Total Kafka Cost | Notes |
|---|---|---|
| **Broker disk storage** | 40–55% | SSD is expensive; grows with retention policy and replication factor |
| **Broker compute (EC2 / VMs)** | 25–35% | Over-provisioning is common; brokers often idle at 20% CPU |
| **Cross-AZ network traffic** | 10–20% | Every replication fetch across AZs costs $0.01–$0.02/GB on AWS |
| **Cross-region network traffic** | 5–15% | MirrorMaker2 or Cluster Linking replication egress is expensive |
| **Kafka Connect workers** | 5–10% | Running 10+ connectors on dedicated VMs adds up |

**The single biggest lever** is almost always disk storage. Retention policies set by developers without cost awareness result in topics retaining data for 30 days when consumers catch up in 2 hours."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Cost drivers are infrastructure-level concerns. Regardless of whether you use Spring Boot or Go, minimizing payload sizes via efficient serialization (Avro/Protobuf) directly reduces broker storage and network egress costs.

#### Indepth
Cross-AZ traffic is the hidden cost that surprises most teams. On AWS, intra-region cross-AZ data transfer costs ~$0.01/GB per direction. A Kafka cluster with replication factor 3, brokers spread across 3 AZs, processing 10TB/day means 20TB of follower-fetch replication traffic daily = $200/day = $73,000/year purely in network fees — for data that never leaves the region.

---

## Q2. How do you reduce broker storage costs?

"**1. Right-size retention policies (biggest impact):**
```bash
# Audit every topic's actual consumer lag patterns
kafka-consumer-groups.sh --bootstrap-server kafka:9092 \
  --describe --all-groups \
  | sort -k6 -rn  # sort by lag descending

# Set retention based on worst-case consumer recovery time,
# NOT developer intuition
kafka-configs.sh --bootstrap-server kafka:9092 \
  --entity-type topics --entity-name app-logs \
  --alter --add-config "retention.ms=86400000"  # 24h, not 7 days
```

**2. Enable compression (reduces storage 3–7x for text/JSON data):**
```properties
# Broker-level default compression (override per-topic too)
compression.type=zstd          # zstd gives best compression ratio
                               # lz4 is faster with slightly lower ratio
```
Impact: A 10TB/day JSON event topic compresses to ~2TB/day with zstd. That's 5x storage and 5x replication network savings together.

**3. Enable Kafka Tiered Storage (Kafka 3.6+):**
Moves cold segments to S3/GCS automatically. Typical saving: 70–85% of broker SSD cost for high-retention topics. (Covered in detail in `11_kafka_tiered_storage.md`)"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka makes compression easy: just set `spring.kafka.producer.properties.compression.type=zstd` in `application.yml`.
* **Golang:** In `kafka-go`, compression is enabled on the `Writer` via `Compression: kafka.Zstd`.

#### Indepth
**Log compaction vs retention:** For topics that represent state (user preferences, account balances), use `cleanup.policy=compact` instead of `cleanup.policy=delete`. Compaction retains only the latest value per key — a topic tracking 1M user preferences stores 1M records regardless of how many updates happened. Without compaction, the same topic grows unboundedly with every update.

---

## Q3. How do you reduce cross-AZ replication costs?

"**Fetch from follower replicas (KIP-392, available since Kafka 2.4):**
```properties
# Broker config — enable rack-aware assignment + follower fetching
broker.rack=us-east-1a          # tag each broker with its AZ

# Consumer config — fetch from same-AZ replica
client.rack=us-east-1a          # consumer fetches from local-AZ follower
replica.selector.class=org.apache.kafka.common.replica.RackAwareReplicaSelector
```

**Before (all consumers fetch from leader):**
```text
Consumer in us-east-1a → fetches from Leader in us-east-1b → cross-AZ cost incurred
```

**After (fetch from follower):**
```text
Consumer in us-east-1a → fetches from Follower in us-east-1a → SAME-AZ, zero cross-AZ cost
```

**Impact at scale (10TB/day consumer throughput, 3 AZs):**
```text
Without follower fetch: 10TB × 2/3 cross-AZ fetches = 6.7TB/day cross-AZ
                        6.7TB × $0.01/GB × 2 directions = $134/day = $49K/year

With follower fetch:    ~10% residual cross-AZ (leader elections, stragglers)
                        ~$5K/year saving of $44K/year
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Rack awareness is configured via consumer properties `client.rack` and `replica.selector.class`.
* **Golang:** In `kafka-go`, if connecting to a cluster supporting follower fetching, you must configure the `Dialer` or `Reader` transport to pass the appropriate rack ID matching the deployment availability zone.

#### Indepth
**Rack awareness during topic creation is a prerequisite.** If all 3 replicas of a partition end up on brokers in the same AZ (without rack-aware assignment), enabling follower fetch provides no cross-AZ savings because there's no local follower in other AZs. Always create topics with `replication-factor=3` on a cluster where `broker.rack` is configured — Kafka automatically spreads replicas across racks.

---

## Q4. How do you reduce cross-region replication costs (MirrorMaker 2)?

"Cross-region network egress is the most expensive Kafka cost item. On AWS, inter-region data transfer costs $0.08–$0.09/GB.

**Strategy 1 — Replicate only essential topics:**
```properties
# MM2: don't replicate everything — be explicit about what's needed for DR
primary->dr.topics = orders, payments, user-events
# NOT: primary->dr.topics = .*  (replicates internal, debug, and low-priority topics too)

primary->dr.topics.exclude = __consumer_offsets, __transaction_state, \
    app-logs, debug-events, metrics-raw
```

**Strategy 2 — Compress at MM2 producer level:**
```properties
# MM2 producer compresses before sending cross-region
dr.producer.compression.type=zstd
```
Compression reduces cross-region bytes by 60–80% for text/JSON events. At $0.08/GB egress, compressing 10TB/day to 2TB/day saves $292,800/year.

**Strategy 3 — Use PrivateLink / VPN instead of public internet:**
AWS PrivateLink between regions eliminates public internet egress fees for replication traffic. Costs ~$0.01/GB (PrivateLink fee vs $0.09/GB public egress) — 9x cheaper for high-volume replication.

**Strategy 4 — Ensure topics are compressed BEFORE replication:**
```bash
# Check source topic compression to avoid MM2 re-compressing
kafka-topics.sh --describe --topic orders --bootstrap-server kafka:9092
# If Compression: None → add compression at source producer
# MM2 can then forward pre-compressed batches without re-encoding
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** If MM2 is passing through compressed messages without re-encoding, ensure your Java producers (`KafkaTemplate`) or Go producers (`kafka.Writer`) are applying optimal compression (like Zstd) BEFORE they hit the source broker.

#### Indepth
**Confluent Cluster Linking egress optimization:** Cluster Linking transfers raw log segments (already compressed) rather than decoding and re-encoding messages. This means the compression savings from source-side compression flow through automatically to cross-region transfer costs — unlike MM2 which passes through the consumer+producer path and may re-compress.

---

## Q5. How do you right-size Kafka broker instances?

"**The most common over-provisioning pattern:** Teams provision Kafka on the same instance types as their databases (memory-optimized, i3 SSD). But Kafka's bottleneck is almost never RAM — it's disk I/O throughput, not IOPS.

**Kafka instance selection framework:**

```text
Step 1: Calculate target throughput
  Peak produce rate: 500 MB/s
  Replication overhead (RF=3): 500 MB/s × 2 followers = 1 GB/s total disk writes
  Required disk write throughput per broker (6 brokers): 1GB/6 = ~170 MB/s/broker

Step 2: Calculate storage requirement
  Daily data volume: 10 TB/day
  Retention: 2 days (hot) + tiered storage for rest
  Replication factor: 3
  Local storage needed: 10TB × 2 days × 3 = 60TB

Step 3: Select instance
  AWS d3.2xlarge: 7.5TB NVMe, 250MB/s write throughput, $0.50/hr
  Need: 60TB / 7.5TB = 8 brokers
  8 × $0.50 × 730 hrs = $2,920/month

  vs. previous over-provisioned choice:
  AWS i3en.6xlarge: 15TB NVMe, 2.5GB/s throughput, $2.70/hr
  4 × $2.70 × 730 = $7,884/month
  
  Saving: $5,000/month ($60K/year) with better-matched instance type
```

**Compute rules of thumb:**
- Kafka is NOT CPU-intensive (zero-copy means minimal CPU for data movement)
- 4–8 vCPUs per broker is usually sufficient up to 1 GB/s throughput
- RAM: allocate 6–8 GB JVM heap; leave rest for OS Page Cache (more is better, up to 32GB)
- Network: ensure instance network bandwidth ≥ peak throughput × replication factor"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** When right-sizing brokers, remember that client efficiency matters. Highly optimized Go consumers or asynchronous reactive Spring Boot consumers can process data faster, reducing consumer lag and thereby reducing the need for extended local retention windows.

#### Indepth
**Spot/Preemptible instances for Kafka brokers:** Unlike stateless services, Kafka brokers have state (data on disk). Using Spot instances without persistent EBS volumes means data loss on reclamation. However, Kafka clusters with RF=3 can safely run **one broker per AZ on Spot** if your Spot interruption handling is correct: use `unclean.leader.election.enable=false`, set `min.insync.replicas=2`, and ensure at least 2 brokers are on On-Demand instances. The interrupted Spot broker will re-sync from the ISR when it restarts. This can save 60–70% on broker compute cost.

---

## Q6. How do you build a Kafka FinOps dashboard? What metrics matter?

"**Key FinOps metrics to track weekly:**

| Metric | How to Get It | Target / Alert |
|---|---|---|
| **Storage per topic (GB)** | JMX `kafka.log:type=Log,name=Size,topic=X,partition=Y` | Alert if topic storage grows >20% week-over-week |
| **Cross-AZ replication bytes** | AWS CloudWatch `DataTransfer-Regional-Out` per broker ENI | Should decrease after follower-fetch enabled |
| **Consumer group lag trend** | `kafka-consumer-groups.sh --describe` daily snapshots | Lag growing → consumer is falling behind → will need more SSD soon |
| **Compression ratio per topic** | (uncompressed bytes - compressed bytes) / uncompressed bytes | < 30% compression → switch to zstd |
| **Topic utilization** | Topics with 0 consumer groups reading them | Zombie topics — delete them |

**Automated cost report script:**
```bash
#!/bin/bash
# Weekly Kafka cost report

echo '=== TOP 10 LARGEST TOPICS BY RETENTION ==='
kafka-log-dirs.sh --bootstrap-server kafka:9092 \
  --describe --topic-list $(kafka-topics.sh --list --bootstrap-server kafka:9092) \
  | python3 parse_log_dirs.py \   # parse JSON output
  | sort -k2 -rn \
  | head -10

echo '=== TOPICS WITH RETENTION > 7 DAYS ==='
for topic in $(kafka-topics.sh --list --bootstrap-server kafka:9092); do
  retention=$(kafka-configs.sh --describe --entity-type topics \
    --entity-name $topic --bootstrap-server kafka:9092 \
    | grep retention.ms | awk '{print $NF}')
  seven_days=604800000
  if [[ $retention -gt $seven_days ]]; then
    echo "WARNING: $topic has retention=$retention ms"
  fi
done

echo '=== ZOMBIE TOPICS (no active consumer groups) ==='
all_consumed=$(kafka-consumer-groups.sh --all-groups --describe \
  --bootstrap-server kafka:9092 | awk '{print $1}' | sort -u)
for topic in $(kafka-topics.sh --list --bootstrap-server kafka:9092 | grep -v '^__'); do
  if ! echo "$all_consumed" | grep -q "$topic"; then
    echo "ZOMBIE: $topic"
  fi
done
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Boot applications can easily export JMX metrics (like topic sizes, if acting as admin clients) into Prometheus using Micrometer for your FinOps dashboards.
* **Golang:** Go applications can utilize the `prometheus/client_golang` library to export custom consumer/producer efficiency metrics alongside infrastructure metrics to give a complete cost picture.

#### Indepth
**Chargeback model for shared Kafka clusters:** When multiple teams share a Kafka cluster, implement a topic-level chargeback: calculate `cost = storage_GB × $0.10/GB/month + produce_mb_per_day × $0.002`. Send weekly email reports per team. Teams with no visibility into their Kafka costs consistently set unlimited retention and high partition counts; visibility alone reduces costs by 20–30%.

---

## Q7. What is your complete Kafka cost optimization checklist?

"**Immediate wins (implement this week):**
- [ ] Audit all topics — delete zombie topics (no consumers for >30 days)
- [ ] Set `compression.type=zstd` at broker level as default
- [ ] Cap `retention.ms` to 3–7 days (not 30 days) for log/event topics
- [ ] Enable `cleanup.policy=compact` for state/changelog topics

**Short-term (1 month):**
- [ ] Enable follower-fetch (`client.rack`, `replica.selector.class`) across all consumers
- [ ] Enable Kafka Tiered Storage for high-volume,  high-retention topics (Kafka 3.6+)
- [ ] Rightsize broker instances based on actual throughput metrics, not over-provisioned defaults
- [ ] Build weekly topic storage report and assign ownership

**Long-term (3–6 months):**
- [ ] Implement topic-level chargeback reporting per team
- [ ] Use PrivateLink / VPN for cross-region MM2 replication traffic
- [ ] Evaluate Confluent Cluster Linking if on Confluent Platform (better compression pass-through)
- [ ] Implement auto-scaling for Kafka Connect workers based on connector lag metrics
- [ ] Move to Kafka on Spot instances for non-critical dev/staging clusters"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Implementing this checklist requires concerted effort across Dev (configuring clients efficiently in Java/Go) and Ops (configuring brokers and networking).
