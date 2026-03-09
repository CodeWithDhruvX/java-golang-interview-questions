# 🚀 Kafka — Performance & Scaling Questions

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Amazon, Netflix, Uber, Meta, Twitter

---

## Q1. How would you design Kafka to handle 1 million messages per second?

"Designing for 1M TPS requires careful planning across multiple dimensions:

**Partition strategy:**
- Target 10-20MB/s per partition (Kafka's sweet spot)
- For 1M messages/sec at 1KB each = 1GB/s throughput
- Need ~50-100 partitions per topic minimum
- Spread partitions across multiple brokers (10-20 per broker)

**Broker sizing:**
- **CPU:** 32+ cores (heavy network I/O and compression)
- **Memory:** 64GB+ (for page cache, not JVM heap)
- **Network:** 10Gbps+ NICs with multiple interfaces
- **Storage:** NVMe SSDs with 100K+ IOPS
- **Disk layout:** Separate disks for logs and OS

**Producer optimization:**
```properties
# High-throughput producer config
batch.size=65536
linger.ms=5
compression.type=lz4
acks=1
max.in.flight.requests.per.connection=5
buffer.memory=67108864
```

**Consumer scaling:**
- Match consumer count to partition count
- Use consumer groups for parallel processing
- Implement async processing pipelines
- Consider multiple consumer groups for different processing needs

**Cluster topology:**
- Multi-AZ deployment for fault tolerance
- Rack-awareness for network locality
- Separate clusters for different workload types
- Use MirrorMaker 2 for cross-region replication

**Monitoring and tuning:**
- Real-time metrics for throughput and latency
- Auto-scaling based on consumer lag
- Capacity planning with 30-40% headroom"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use `ConcurrentKafkaListenerContainerFactory` with high concurrency. Implement async processing with `@Async` and `CompletableFuture`.
* **Golang:** Leverage Go's goroutines for massive parallelism. Use worker pools with buffered channels to control memory usage while maintaining high throughput.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Twitter, Meta — platforms processing billions of events daily where infrastructure efficiency directly impacts operational costs.

#### Indepth
**Compression Trade-offs:** LZ4 offers best balance of speed and compression ratio. Snappy is faster but less efficient. Gzip provides best compression but highest CPU cost. For 1M TPS, LZ4 is usually optimal.

---

## Q2. How do you determine the optimal number of partitions for a topic?

"Choosing partition count is critical for both performance and operational complexity. Here's my systematic approach:

**Factors to consider:**

1. **Throughput requirements:**
   - Measure target messages/sec and bytes/sec
   - Each partition typically handles 10-20MB/s
   - Formula: `partitions = target_throughput / partition_capacity`

2. **Consumer parallelism:**
   - Maximum concurrent consumers = partition count
   - Plan for future scaling needs
   - Consider peak load vs average load

3. **Latency requirements:**
   - More partitions = lower per-partition load = better latency
   - But too many partitions increase overhead

4. **Storage and retention:**
   - Each partition creates separate log files
   - More partitions = more file handles = more memory usage
   - Consider disk space and retention policies

**Practical guidelines:**
- **Start small:** 6-12 partitions for most workloads
- **Scale up:** Increase partitions as load grows
- **Monitor:** Watch partition-level metrics
- **Don't over-partition:** 100+ partitions create operational complexity

**Partition calculation example:**
```
Target: 100K messages/sec, 500 bytes each = 50MB/s
Partition capacity: 15MB/s
Required partitions: 50/15 = 3.33 → 6 partitions (round up)
Add 50% headroom: 6 × 1.5 = 9 partitions
Final: 12 partitions (next power of 2 for good distribution)
```

**Important considerations:**
- Partition count cannot be decreased
- Re-partitioning requires data migration
- More partitions = longer recovery times
- Consider key distribution when choosing count"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use `@KafkaListener(topics = "my-topic", concurrency = "6")` to match partition count. Monitor consumer lag per partition.
* **Golang:** Use `kafka.Reader` with appropriate `PartitionCount`. Implement dynamic consumer scaling based on partition count changes.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Uber — where efficient resource utilization and cost optimization are critical at scale.

#### Indepth
**Partition Reassignment:** When increasing partitions, use `kafka-topics.sh --alter` to add partitions. Existing data stays in original partitions, but new messages will be distributed across all partitions using new hashing.

---

## Q3. What are the limitations of having too many partitions?

"While partitions enable parallelism, excessive partitioning creates serious problems:

**Performance impacts:**
1. **Increased latency:**
   - Each partition requires separate fetch requests
   - More network roundtrips for consumers
   - Broker memory fragmentation

2. **Resource consumption:**
   - Each partition = separate file handles
   - More memory for partition metadata
   - Increased disk I/O from concurrent log writes

3. **Replication overhead:**
   - Each partition replicated across brokers
   - More network traffic for replication
   - Slower recovery times after failures

**Operational complexity:**
1. **Monitoring challenges:**
   - Thousands of metrics to track
   - Hard to identify problematic partitions
   - Alert fatigue from too many signals

2. **Rebalancing issues:**
   - Longer rebalance times with many partitions
   - Higher risk of rebalance storms
   - More complex consumer group management

3. **Maintenance overhead:**
   - Longer cluster startup times
   - More complex backup and restore
   - Slower rolling upgrades

**Specific limits to watch:**
- **File handles:** Each partition needs multiple files (.log, .index, .timeindex)
- **Memory:** Broker memory scales with partition count
- **Controller load:** More partitions = more metadata to manage
- **ZooKeeper/KRaft:** Increased metadata storage and synchronization

**Rule of thumb:**
- **Small clusters:** <1000 partitions total
- **Medium clusters:** 1000-10,000 partitions
- **Large clusters:** 10,000+ partitions (requires careful planning)"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Monitor JVM heap usage with many partitions. Use `jstat` to track GC patterns and memory fragmentation.
* **Golang:** Go's runtime handles many goroutines well, but monitor file descriptor limits (`ulimit -n`) when dealing with many partitions.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** LinkedIn, Netflix — companies that learned partition limits the hard way and now have strict governance policies.

#### Indepth
**Partition Count per Broker:** Kafka recommends <2000 partitions per broker in production. Beyond this, you'll see increased latency and potential stability issues. Some companies push to 4000+ with careful tuning.

---

## Q4. How do you tune Kafka for maximum throughput vs minimum latency?

"Throughput and latency optimization require opposite tuning strategies. Here's how to approach each:

**Maximum throughput tuning:**

1. **Producer configuration:**
```properties
# Batch messages for efficiency
batch.size=65536
linger.ms=20
compression.type=lz4

# Allow more in-flight requests
max.in.flight.requests.per.connection=5

# Larger buffers
buffer.memory=134217728
```

2. **Broker tuning:**
```properties
# Larger network buffers
socket.send.buffer.bytes=102400
socket.receive.buffer.bytes=102400

# Better disk throughput
num.io.threads=8
num.network.threads=8

# Larger fetch sizes
fetch.message.max.bytes=1048576
```

3. **Consumer optimization:**
```properties
# Fetch more data per request
fetch.min.bytes=50000
max.partition.fetch.bytes=1048576
fetch.max.wait.ms=500
```

**Minimum latency tuning:**

1. **Producer configuration:**
```properties
# Send immediately
linger.ms=0
batch.size=16384

# No compression for speed
compression.type=none

# Wait for leader only
acks=1
```

2. **Consumer optimization:**
```properties
# Fetch immediately
fetch.min.bytes=1
fetch.max.wait.ms=0

# Smaller fetches for responsiveness
max.partition.fetch.bytes=65536
```

**Infrastructure considerations:**

**For throughput:**
- Use larger, more powerful machines
- Optimize for sequential I/O
- Use high-throughput networks
- Batch processing in consumers

**For latency:**
- Use faster CPUs (not just more cores)
- Optimize network latency (same AZ)
- Use NVMe storage
- Minimize processing time per message

**Monitoring approach:**
- **Throughput:** Monitor bytes/sec, messages/sec
- **Latency:** Track end-to-end latency percentiles
- **Balance:** Find sweet spot for your use case"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use different `KafkaTemplate` beans for throughput vs latency scenarios. Profile application to identify bottlenecks.
* **Golang:** Go's runtime is naturally low-latency. Use sync.Pool for object reuse in throughput scenarios. Use channels for pipeline parallelism.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** High-frequency trading, real-time gaming — where microsecond optimizations matter and throughput/latency trade-offs are critical.

#### Indepth
**NUMA Awareness:** For extreme throughput, consider NUMA topology. Pin Kafka processes to specific CPU cores and memory nodes to reduce cross-NUMA traffic. Use `numactl` for process binding.

---

## Q5. How would you design a Kafka cluster for 100TB/day data ingestion?

"Designing for 100TB/day requires massive scale and careful architecture:

**Daily volume breakdown:**
- 100TB/day = 1.16GB/s sustained
- Peak hours might be 3-5× average = 3-5GB/s
- With replication factor 3 = 9-15GB/s disk write throughput

**Cluster sizing:**
1. **Broker count:**
   - Each broker handles ~200MB/s write throughput
   - Need 50-75 brokers for peak load
   - Add 30% headroom = 65-100 brokers

2. **Storage per broker:**
   - Daily data per broker: 100TB/75 = 1.3TB/day
   - 7-day retention = 9.1TB per broker
   - With overhead = 12TB usable storage
   - Raw storage needed = 36TB (RAID 10)

3. **Network infrastructure:**
   - 40Gbps+ network backbone
   - Multiple network interfaces per broker
   - Dedicated storage network for replication
   - Cross-region links for DR

**Partition strategy:**
- Target 1000+ topics with 10-50 partitions each
- Total partitions: 10,000-50,000
- Spread evenly across brokers
- Use rack-awareness for fault tolerance

**Performance optimization:**
1. **Hardware:**
   - NVMe SSDs for log storage
   - 64+ cores per broker
   - 256GB+ RAM for page cache
   - Multiple 10Gbps NICs

2. **Software tuning:**
   - Large page cache (leave most RAM for OS)
   - Optimize TCP stack for high throughput
   - Tune JVM for large heap (if using Java brokers)
   - Enable compression at producer level

**Operational considerations:**
- Multi-region active-active setup
- Automated failover and recovery
- Real-time monitoring and alerting
- Capacity planning for growth

**Cost optimization:**
- Tiered storage for old data
- Compression for cold partitions
- Spot instances for non-critical workloads
- Automated scaling based on usage patterns"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use reactive programming (Spring WebFlux) for high-throughput ingestion. Implement backpressure to prevent overwhelming downstream systems.
* **Golang:** Go's efficiency makes it ideal for high-throughput ingestion services. Use connection pooling and batch processing to maximize throughput.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Netflix, Amazon AWS, Google Cloud — companies providing Kafka as a service or running massive internal data platforms.

#### Indepth
**Tiered Storage:** Kafka 3.0+ supports tiered storage, moving old data to cheaper storage (S3, HDFS). This reduces local storage costs while maintaining recent data on fast SSDs for low-latency access.
