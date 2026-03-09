# 🐛 Kafka — Production Debugging Scenarios

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, Uber, Amazon, LinkedIn, Meta

---

## Q1. What causes rebalance storms and how do you prevent them?

"Rebalance storms occur when consumers repeatedly join and leave a group, causing continuous partition reassignments. This is devastating for throughput because:

**Common causes:**
1. **Too aggressive timeouts:** `session.timeout.ms` too low (default 10s)
2. **Slow message processing:** Processing time exceeds `max.poll.interval.ms` (default 5min)
3. **GC pauses:** Long garbage collection stops heartbeats
4. **Network issues:** Intermittent connectivity causes consumer timeouts
5. **Frequent deployments:** Rolling restarts without proper coordination

**Prevention strategies:**
1. **Tune timeouts appropriately:**
   - `session.timeout.ms`: 15-30s (not too aggressive)
   - `heartbeat.interval.ms`: 3-5s (1/3 of session timeout)
   - `max.poll.interval.ms`: Based on your processing time + buffer

2. **Use cooperative rebalancing:** `partition.assignment.strategy=CooperativeStickyAssignor`
   - Only reassigns necessary partitions
   - Consumers keep existing partitions during rebalance
   - Much less disruptive than eager rebalancing

3. **Optimize processing:**
   - Reduce per-message processing time
   - Use async processing with proper offset management
   - Implement backpressure mechanisms

4. **Graceful shutdown:** Call `consumer.close()` properly to trigger clean rebalance"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Set `spring.kafka.listener.partition-assignment-strategy=CooperativeStickyAssignor`. Use `@PreDestroy` to ensure graceful consumer shutdown.
* **Golang:** Implement proper context cancellation in Go consumers. Use `consumer.Close()` on shutdown signals to trigger clean rebalance.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Lyft — where high-frequency rebalances directly impact ride-matching latency and driver experience.

#### Indepth
**Rebalance Callbacks:** Implement `ConsumerRebalanceListener` to pause processing during rebalance and commit offsets before losing partitions. This prevents duplicate processing when partitions move between consumers.

---

## Q2. How do you debug consumer lag in production?

"Consumer lag is when consumers can't keep up with producers. Here's my systematic debugging approach:

**Step 1: Measure lag accurately**
```bash
# Using Kafka consumer groups command
kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --describe --group my-consumer-group

# Look for:
# - GROUP, TOPIC, PARTITION, CURRENT-OFFSET, LOG-END-OFFSET, LAG
```

**Step 2: Identify the bottleneck**
1. **Consumer processing too slow:**
   - Check CPU/memory usage on consumer machines
   - Profile application for hotspots
   - Look at processing time per message

2. **Network bottleneck:**
   - Check `fetch.min.bytes` and `max.partition.fetch.bytes`
   - Monitor network bandwidth between consumers and brokers
   - Look at `request.timeout.ms` vs actual fetch times

3. **Broker overload:**
   - Check broker CPU, disk I/O, network
   - Monitor `under-replicated-partitions`
   - Look at broker request queue sizes

**Step 3: Common fixes**
1. **Scale consumers:** Add more consumers (up to partition count)
2. **Increase partitions:** If consumers < partitions, add more partitions
3. **Optimize processing:**
   - Batch processing instead of per-message
   - Async processing with worker pools
   - Reduce database calls per message
4. **Tune fetch settings:**
   - Increase `fetch.min.bytes` to reduce network roundtrips
   - Adjust `max.poll.records` based on processing capacity"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Actuator endpoints `/actuator/kafka/consumerinfo` to monitor lag. Implement custom metrics using Micrometer.
* **Golang:** Use `confluent-kafka-go`'s `consumer.Lag()` method. Export Prometheus metrics for lag monitoring.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Amazon — where lag directly impacts user experience (video recommendations, order processing).

#### Indepth
**Lag Monitoring Tools:** Burrow (Uber's open-source tool) provides sophisticated lag monitoring with anomaly detection. It can predict when consumers will fall behind based on historical patterns.

---

## Q3. What causes split-brain scenarios in Kafka clusters and how do you prevent them?

"Split-brain occurs when a Kafka cluster partitions into multiple groups that can't communicate, each thinking it's the legitimate cluster. This is catastrophic because:

**Split-brain causes:**
1. **Network partitions:** Network failures isolate broker subsets
2. **ZooKeeper/KRaft quorum loss:** Controller becomes unreachable
3. **Misconfigured discovery:** Brokers using different controller addresses

**Dangers of split-brain:**
- Multiple leaders for same partitions
- Data corruption and divergence
- Inconsistent consumer experiences
- Potential data loss during reconciliation

**Prevention strategies:**

1. **Proper quorum configuration:**
   - ZooKeeper: Odd number of nodes (3, 5, 7)
   - KRaft: Set `controller.quorum.voters` correctly
   - Ensure majority quorum can always be achieved

2. **Network monitoring:**
   - Monitor broker-to-broker connectivity
   - Set up alerts for partitioned clusters
   - Use health checks before failover

3. **Controller isolation:**
   - Deploy controllers on separate network segments
   - Use dedicated hardware for critical components
   - Implement network redundancy

4. **Automatic failover:**
   - Configure `unclean.leader.election.enable=false` (default)
   - This prevents data loss by waiting for proper quorum
   - Accept temporary unavailability over data corruption

**Recovery from split-brain:**
1. Identify the legitimate quorum
2. Shut down brokers in minority partition
3. Verify data consistency across partitions
4. Restart brokers with correct configuration"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Implement circuit breakers to detect and stop processing during suspected split-brain. Use Spring Cloud Config for consistent broker configuration.
* **Golang:** Use Go's context with timeouts to detect hanging operations. Implement health checks that verify cluster quorum before processing.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Amazon, Google — where data consistency across regions is critical and split-brain could cause massive financial losses.

#### Indepth
**KRaft vs ZooKeeper:** KRaft mode reduces split-brain risk by eliminating the separate ZooKeeper ensemble. The Raft protocol provides stronger consistency guarantees for controller elections and metadata management.

---

## Q4. How do you debug uneven partition load distribution?

"Uneven partition load is when some partitions have much higher throughput than others, causing hot spots. Here's my debugging approach:

**Identify the problem:**
```bash
# Check partition-level metrics
kafka-run-class.sh kafka.tools.GetOffsetShell \
  --broker-list localhost:9092 \
  --topic my-topic --time -1

# Monitor per-partition lag
kafka-consumer-groups.sh --describe --group my-group
```

**Common causes:**

1. **Poor partitioning strategy:**
   - Hash-based partitioning with uneven key distribution
   - Range partitioning with skewed data
   - Default partitioner not suitable for your key patterns

2. **Hot keys:**
   - Certain keys get much more traffic than others
   - Example: popular products in e-commerce
   - Celebrity users in social platforms

3. **Time-based patterns:**
   - Certain hours generate more traffic
   - Seasonal events create temporary skew

**Solutions:**

1. **Custom partitioner:**
```java
public class UniformHashPartitioner implements Partitioner {
    private final ConsistentHash<String> hashRing;
    
    @Override
    public int partition(String topic, Object key, byte[] keyBytes, 
                        Object value, byte[] valueBytes, Cluster cluster) {
        // Use consistent hashing for better distribution
        return hashRing.getNode(key.toString());
    }
}
```

2. **Salting technique:**
   - Add random prefix to keys: `key_1`, `key_2`, `key_3`
   - Consumers handle all salted versions
   - Spreads hot keys across partitions

3. **Dynamic partitioning:**
   - Monitor partition sizes in real-time
   - Add partitions when skew detected
   - Use Kafka's partition expansion capabilities

4. **Consumer-side load balancing:**
   - Implement smart consumer assignment
   - Move consumers to heavily loaded partitions
   - Use cooperative rebalancing for smooth transitions"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Implement custom `Partitioner` as a Spring bean. Use `@KafkaListener` with `topicPattern` to handle dynamic partition additions.
* **Golang:** Create custom partitioner using `segmentio/kafka-go`'s `Balancer` interface. Use Go's maps to track partition load and rebalance consumers.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Twitter, Instagram — where celebrity content creates massive hot spots and requires sophisticated load distribution.

#### Indepth
**Partition Reassignment:** Use `kafka-reassign-partitions.sh` tool to manually rebalance partitions. Generate assignment plans based on current load and move partitions without downtime using preferred replica election.

---

## Q5. How do you troubleshoot high end-to-end latency in Kafka pipelines?

"High end-to-end latency in Kafka pipelines requires systematic analysis across the entire chain:

**Measurement approach:**
1. **Producer latency:** Time from message creation to broker acknowledgment
2. **Broker latency:** Time from broker receive to storage
3. **Consumer latency:** Time from broker fetch to processing completion
4. **Network latency:** Time spent in transit between components

**Common latency sources:**

1. **Producer-side issues:**
   - `linger.ms` too high (default 0, but sometimes set to 20-50ms)
   - `batch.size` too large causing delays
   - Compression overhead (especially with gzip)
   - Network congestion to brokers

2. **Broker-side issues:**
   - Disk I/O bottlenecks (slow storage)
   - High CPU usage causing request queuing
   - Under-replicated partitions causing sync delays
   - Too many small requests (inefficient batching)

3. **Consumer-side issues:**
   - `fetch.min.bytes` set too high
   - Long processing time per message
   - Frequent rebalances interrupting consumption
   - Backpressure from downstream systems

**Optimization strategies:**

1. **Tune for low latency:**
```properties
# Producer
linger.ms=0
batch.size=16384
compression.type=none
acks=1

# Consumer
fetch.min.bytes=1
max.partition.fetch.bytes=1048576
fetch.max.wait.ms=500
```

2. **Monitor key metrics:**
   - `request-latency-avg` (broker)
   - `produce-throttle-time` (producer)
   - `fetch-latency-avg` (consumer)
   - `network-io-wait-time-ns-avg`

3. **Infrastructure optimizations:**
   - Use NVMe SSDs for broker storage
   - Place consumers in same AZ as brokers
   - Use dedicated network for Kafka traffic
   - Enable TCP_NODELAY for low-latency networks"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring's `Micrometer` to track latency at each stage. Implement custom `ProducerInterceptor` to add timestamps.
* **Golang:** Use Go's `time.Now()` to measure latency at each step. Leverage Go's built-in pprof to identify CPU bottlenecks.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** High-frequency trading firms, real-time gaming platforms — where sub-millisecond latency is critical for business success.

#### Indepth
**JVM Tuning:** For Java-based brokers, tune GC for low latency: use G1GC with `-XX:MaxGCPauseMillis=20`, disable biased locking, and allocate sufficient heap to avoid frequent GC cycles.
