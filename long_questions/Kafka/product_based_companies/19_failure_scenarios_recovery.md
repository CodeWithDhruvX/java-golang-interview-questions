# 🚨 Kafka — Failure Scenarios & Recovery

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Amazon, Netflix, Uber, Google, Meta

---

## Q1. What happens when a Kafka broker crashes and how does recovery work?

"Broker crash recovery is one of Kafka's core strengths. Here's the complete recovery sequence:

**Immediate impact:**
1. **Controller detects failure:**
   - Heartbeat timeout (default 10s) triggers failure detection
   - Controller removes broker from cluster metadata
   - Partition leadership transfer initiated

2. **Leader election for affected partitions:**
   - For each partition where crashed broker was leader:
   - Controller selects new leader from remaining ISR replicas
   - Updates partition metadata in ZooKeeper/KRaft
   - Notifies all brokers and consumers of new leader

3. **Producer behavior:**
   - Producers get `NotLeaderForPartitionException`
   - Automatic retry with metadata refresh
   - Brief pause in production (typically <1 second)

4. **Consumer behavior:**
   - Consumers get `NotLeaderForPartitionException` on fetch
   - Automatic metadata refresh and retry
   - Brief pause in consumption

**Recovery timeline:**
- **Detection:** 10-30 seconds (heartbeat timeout)
- **Leader election:** 1-5 seconds
- **Metadata propagation:** 1-2 seconds
- **Total downtime:** Usually <1 minute

**Configuration impact:**
- `unclean.leader.election.enable=false` (default) prevents data loss
- `auto.leader.rebalance.enable=true` rebalances leaders after recovery
- `replica.lag.time.max.ms` controls when replicas fall out of ISR

**Data integrity:**
- No data loss if replication factor ≥ 2
- ISR ensures only fully replicated data is visible
- Consumers only read up to high watermark

**Post-recovery:**
- Failed broker rejoins cluster as follower
- Syncs with current leaders to catch up
- May become leader again after rebalancing"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring Kafka handles automatic retries and metadata refresh. Configure `spring.kafka.consumer.recovery-interval` for retry timing.
* **Golang:** `confluent-kafka-go` automatically retries on broker failures. Implement exponential backoff for custom retry logic.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Uber — where broker failures are common at scale and quick recovery is critical for service availability.

#### Indepth
**Controller Failover:** If the controller itself crashes, a new controller is elected from remaining brokers. This adds ~30 seconds to recovery time but ensures cluster continues operating.

---

## Q2. How does Kafka handle network partitions and split-brain scenarios?

"Network partitions are one of the most dangerous failure scenarios. Here's how Kafka handles them:

**Partition detection:**
1. **Broker-to-broker heartbeats:** Continuous health checks
2. **Controller monitoring:** Tracks broker availability
3. **ZooKeeper/KRaft quorum:** Maintains cluster metadata consistency

**Split-brain prevention:**
1. **Quorum-based decisions:**
   - Controller requires majority quorum
   - No decisions without quorum
   - Minority partition becomes read-only

2. **Unclean leader election disabled:**
   - `unclean.leader.election.enable=false` (default)
   - Prevents data loss during partitions
   - Prefers unavailability over inconsistency

3. **ISR management:**
   - Only in-sync replicas can become leaders
   - Lagging replicas removed from ISR
   - Prevents stale data from becoming visible

**Behavior during partition:**

**Majority side (continues operating):**
- Maintains quorum and controller
- Continues accepting writes
- Performs leader elections as needed
- Serves consumers normally

**Minority side (becomes read-only):**
- Loses quorum, no controller
- Cannot perform leader elections
- Producers get `NotEnoughReplicasException`
- Eventually stops serving requests

**Recovery process:**
1. **Network restored:**
   - Brokers reconnect to cluster
   - Metadata synchronization occurs
   - ISR membership updated

2. **Data reconciliation:**
   - Lagging brokers catch up with leaders
   - Out-of-sync data discarded
   - Cluster returns to normal operation

**Configuration best practices:**
```properties
# Prevent split-brain
unclean.leader.election.enable=false
min.insync.replicas=2

# Conservative timeouts
zookeeper.session.timeout.ms=30000
zookeeper.connection.timeout.ms=30000
```

**Monitoring for partitions:**
- Track `OfflinePartitionsCount`
- Monitor `UnderReplicatedPartitions`
- Alert on controller changes
- Network latency monitoring between brokers"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Implement circuit breakers to detect partition scenarios. Use Spring Cloud Config for consistent configuration across brokers.
* **Golang:** Use Go's context with timeouts to detect hanging operations. Implement health checks that verify cluster connectivity.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Google, Amazon — companies with global infrastructure where network partitions are common and data consistency is critical.

#### Indepth
**KRaft vs ZooKeeper:** KRaft mode handles partitions more gracefully with the Raft consensus protocol, providing stronger guarantees and faster recovery from network partitions.

---

## Q3. What happens when all replicas for a partition go down?

"This is a catastrophic failure scenario with specific recovery behavior:

**Immediate impact:**
1. **Partition becomes offline:**
   - No leader available for the partition
   - Producers get `LeaderNotAvailableException`
   - Consumers cannot read from the partition

2. **Cluster state:**
   - Partition marked as offline in metadata
   - `OfflinePartitionsCount` metric increases
   - Cluster continues operating for other partitions

**Recovery scenarios:**

**Scenario 1: At least one replica comes back online:**
1. **First replica starts:**
   - Checks its log data
   - Attempts to become leader
   - If `unclean.leader.election=false`, waits for other replicas

2. **Leader election:**
   - If multiple replicas return, normal ISR-based election
   - First replica may become leader if it's the only option
   - Partition comes back online

3. **Data availability:**
   - Only data available on recovered replicas
   - Any data written to lost replicas is gone
   - Consumers can resume from available data

**Scenario 2: All replicas permanently lost:**
1. **Manual intervention required:**
   - Partition remains offline
   - Requires administrator action
   - Data in that partition is permanently lost

2. **Recovery options:**
   - Recreate topic (loses all data)
   - Use backup if available
   - Accept data loss for that partition

**Configuration impact:**
- `unclean.leader.election.enable=true` allows faster recovery but risks data loss
- `min.insync.replicas` affects when writes are rejected
- Replication factor determines resilience

**Prevention strategies:**
1. **Multi-AZ deployment:**
   - Spread replicas across availability zones
   - Prevents single AZ failure from losing all replicas
   - Requires proper rack awareness configuration

2. **Regular backups:**
   - MirrorMaker 2 to separate cluster
   - Regular snapshots of log data
   - Test recovery procedures

3. **Monitoring:**
   - Alert on under-replicated partitions
   - Monitor broker health continuously
   - Automated failover testing"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Implement dead letter queues for critical data. Use Spring Retry with exponential backoff for handling unavailable partitions.
* **Golang:** Use Go's error handling to detect and log partition unavailability. Implement fallback mechanisms for critical data paths.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Financial institutions, healthcare — where data loss is unacceptable and disaster recovery is critical.

#### Indepth
**Tiered Storage:** Modern Kafka can offload old data to object storage, providing additional protection against complete data loss even if all broker disks fail.

---

## Q4. How do you perform zero-downtime rolling upgrades of Kafka clusters?

"Zero-downtime upgrades require careful planning and execution. Here's the complete process:

**Pre-upgrade preparation:**
1. **Backup critical data:**
   - Export ZooKeeper/KRaft metadata
   - Verify MirrorMaker 2 replication
   - Document current configuration

2. **Health checks:**
   - Verify cluster is healthy
   - No under-replicated partitions
   - Consumer lag within acceptable limits

3. **Compatibility testing:**
   - Test new Kafka version in staging
   - Verify client compatibility
   - Test configuration changes

**Rolling upgrade process:**
1. **Upgrade one broker at a time:**
   - Stop broker gracefully
   - Update Kafka binaries
   - Update configuration if needed
   - Restart broker

2. **Wait for recovery:**
   - Broker rejoins cluster
   - Partitions re-replicate
   - ISR restored to full size
   - Verify no offline partitions

3. **Continue to next broker:**
   - Only proceed after previous broker is healthy
   - Monitor cluster metrics
   - Watch for unexpected behavior

**Configuration considerations:**
```properties
# Enable rolling upgrades
controlled.shutdown.enable=true
controlled.shutdown.max.retries=3

# Inter-broker protocol version
inter.broker.protocol.version=3.5
log.message.format.version=3.5
```

**Client compatibility:**
- New brokers work with old clients
- Old brokers work with new clients
- Upgrade clients after brokers complete
- Test thoroughly before production

**Monitoring during upgrade:**
- `UnderReplicatedPartitions` should be 0
- `OfflinePartitionsCount` should be 0
- Consumer lag should not increase
- Request latency should remain stable

**Rollback plan:**
1. **Stop upgrade process**
2. **Roll back to previous version**
3. **Verify cluster health**
4. **Investigate issues**

**Common pitfalls:**
- Upgrading multiple brokers simultaneously
- Ignoring client compatibility
- Not waiting for full recovery
- Insufficient monitoring

**Upgrade timeline:**
- Small cluster (3-5 brokers): 1-2 hours
- Medium cluster (10-20 brokers): 4-6 hours
- Large cluster (50+ brokers): 1-2 days"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Use Spring Boot Actuator health endpoints to monitor application health during upgrades. Implement graceful shutdown hooks.
* **Golang:** Use Go's signal handling for graceful shutdown. Implement health check endpoints that verify Kafka connectivity.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Amazon — companies with 24/7 services where downtime is unacceptable and upgrades must be seamless.

#### Indepth
**KRaft Migration:** When upgrading from ZooKeeper to KRaft, perform a two-phase migration: first run in dual mode (both ZK and KRaft), then switch to KRaft-only after verification.

---

## Q5. How do you handle disk corruption and data recovery in Kafka?

"Disk corruption is a serious issue that requires immediate attention. Here's the recovery process:

**Detection:**
1. **Kafka logs corruption errors:**
   - `LogSegment` corruption detected
   - `IllegalStateException` in broker logs
   - Broker fails to read from corrupted segment

2. **Symptoms:**
   - Broker crashes or becomes unresponsive
   - Consumers get `CorruptRecordException`
   - Under-replicated partitions increase

**Immediate response:**
1. **Isolate affected broker:**
   - Stop broker to prevent further damage
   - Identify corrupted partitions/segments
   - Check system logs for hardware issues

2. **Assess damage:**
   - Run `kafka-run-class.sh kafka.tools.DumpLogSegments`
   - Identify corrupted segments
   - Determine if recovery is possible

**Recovery options:**

**Option 1: Delete corrupted segments:**
```bash
# Stop broker
systemctl stop kafka

# Find corrupted segments
find /kafka-logs -name "*.log" -exec kafka-run-class.sh kafka.tools.DumpLogSegments --files {} \;

# Remove corrupted segments
rm /kafka-logs/topic-name/partition-number/00000000000000000000.corrupted

# Restart broker
systemctl start kafka
```

**Option 2: Restore from replicas:**
1. **Broker automatically recovers:**
   - Other replicas have clean data
   - Broker syncs from leader
   - Data consistency maintained

2. **Manual intervention:**
   - If all replicas corrupted, restore from backup
   - Use MirrorMaker 2 if available
   - Recreate topic as last resort

**Option 3: File system recovery:**
```bash
# Check file system
fsck -f /dev/sda1

# Attempt repair
fsck -y /dev/sda1

# Check disk health
smartctl -a /dev/sda
```

**Prevention strategies:**
1. **RAID configuration:**
   - Use RAID 10 for performance and redundancy
   - Monitor RAID health continuously
   - Replace failed disks immediately

2. **Regular backups:**
   - MirrorMaker 2 to separate cluster
   - Regular snapshots to cloud storage
   - Test restore procedures

3. **Hardware monitoring:**
   - SMART disk monitoring
   - Temperature and vibration sensors
   - Predictive failure analysis

**Configuration for resilience:**
```properties
# Enable log cleaning for corrupted data
log.cleaner.enable=true

# Configure retention to prevent disk full
log.retention.hours=168
log.segment.bytes=1073741824
```

**Post-recovery verification:**
- Verify all partitions are online
- Check consumer lag is normal
- Run data consistency checks
- Monitor broker health for 24-48 hours"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Implement data validation in consumers to detect corruption. Use Spring Retry for handling transient corruption errors.
* **Golang:** Use Go's checksum validation for critical data. Implement circuit breakers to skip corrupted messages while continuing processing.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Financial services, healthcare — where data integrity is critical and corruption must be handled immediately.

#### Indepth
**Tiered Storage:** With tiered storage enabled, corrupted local data can be recovered from cloud storage, providing an additional layer of protection against disk corruption.
