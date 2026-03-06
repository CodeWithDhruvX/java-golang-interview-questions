# 🏗️ Kafka — Advanced Producer & Consumer Tuning

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Amazon, Uber, Flipkart, PhonePe, Razorpay

---

## Q1. Explain the `acks` producer config. What are the trade-offs between `acks=0`, `acks=1`, and `acks=all`?

"`acks` (acknowledgements) controls how many broker replicas must confirm receipt of a message before the producer considers it successfully sent. It is the central knob that balances **throughput vs. durability**.

**`acks=0` (Fire and Forget):**
The producer sends the message and does NOT wait for any acknowledgement from the broker. This is the fastest possible mode. If the broker crashes mid-flight, the message is silently lost. Only acceptable for non-critical metrics or telemetry data.

**`acks=1` (Leader Only):**
The partition leader acknowledges the write, but the followers haven't yet replicated it. If the leader crashes *before* the followers have synced, that message is permanently lost. Decent balance for non-critical event streams.

**`acks=all` (also written as `acks=-1`) (Majority Quorum):**
The leader waits until all replicas in the **ISR (In-Sync Replicas)** list have confirmed the write before acknowledging the producer. This is the strongest durability guarantee. Must be paired with `min.insync.replicas=2` on the broker to prevent the scenario where the ISR has shrunk to just 1 (the leader), which would make `acks=all` functionally equivalent to `acks=1`."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Controlled in `application.yml` via `spring.kafka.producer.acks=all`. Default has shifted to `all` since Kafka 3.0 ensuring high durability baseline.
* **Golang:** Passed directly into the writer configuration struct (e.g. `RequiredAcks: kafka.RequireAll`). Setting it appropriately is critical since Go defaults historically leaned towards performance (`RequireOne`).

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay, PhonePe — financial transactions must use `acks=all` + `min.insync.replicas=2` to ensure zero data loss. This is a critical design interview question.

#### Indepth
**The Unclean Leader Election Trap:** Setting `unclean.leader.election.enable=true` (default in older Kafka versions) allows an out-of-sync replica to become leader during a crisis, sacrificing data consistency for availability. Modern production systems always set this to `false` and pair it with `acks=all` for a truly durable pipeline.

---

## Q2. How do you design partition keys to avoid hot partitions?

"A **hot partition** occurs when the partition key skews — too many messages route to the same partition, creating an overloaded broker while others sit idle. This destroys horizontal scalability.

**Common Causes:**
- Using a non-uniform key (e.g., `countryCode=IN` when 90% of users are Indian)
- Using a timestamp as the key (all messages in the same second land on the same partition)
- Using `null` keys (Kafka uses round-robin, which distributes evenly, but loses ordering guarantees)

**Solutions:**

**1. Composite Key with Salt:**
```
partitionKey = userId + "_" + (random.nextInt(numPartitions))
```
This breaks the hotspot by spreading a single user's messages across partitions. Trade-off: per-key ordering is lost.

**2. Virtual Partitioning:**
Use a `region_userId` key. For global apps, include the region prefix so US and IN events use different key spaces naturally.

**3. Custom Partitioner:**
Implement the `Partitioner` interface to write domain-specific logic — e.g., VIP users (high-value orders) get routed to dedicated fast partitions while regular users use the remaining pool."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Creating a custom partitioner requires implementing `org.apache.kafka.clients.producer.Partitioner` and supplying the fully qualified class name in the `spring.kafka.producer.properties.partitioner.class` property.
* **Golang:** Both `kafka-go` and `confluent-kafka-go` lack interfaces analogous to Java's strict `Partitioner` class loading. Instead, developers utilize custom Balancer interfaces (for `kafka-go`) or simply compute the partition integer locally in Go, bypassing the default hash routine, and setting `Message.Partition` themselves.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Amazon — in ride-booking or order systems, specific driver IDs or merchant IDs can dominate traffic during a surge event, causing a single broker to fail under load.

#### Indepth
**Monitoring Hot Partitions:** Use the JMX metric `BytesInPerSec` per partition to visualize the imbalance in tools like Grafana. A 10x difference in partition throughput is the canonical signature of a hot partition.

---

## Q3. Explain `max.poll.records`, `max.poll.interval.ms`, and how they interact to prevent consumer group rebalances.

"These two configs govern the fundamental rhythm of a Kafka consumer's poll loop and must be tuned together to avoid involuntary rebalances.

**`max.poll.records` (default: 500):**
The maximum number of records returned in a single `poll()` call. If your consumer's processing logic is slow (e.g., it calls an external DB for each record), fetching 500 records at once means the processing batch takes a very long time before the next `poll()` is called.

**`max.poll.interval.ms` (default: 300,000ms = 5 minutes):**
The maximum time the consumer can go between two consecutive `poll()` calls before the broker considers it 'dead' and triggers a rebalance. If `processing_time(max.poll.records) > max.poll.interval.ms`, you get a continuous rebalance storm.

**The Fix — Balance them:**
```
# If each record takes ~100ms to process, and you have 500ms per poll interval budget:
max.poll.records = 5          # Process 5 records × 100ms = 500ms
max.poll.interval.ms = 2000   # Give ample buffer: 2 seconds
```

The rule of thumb: `max.poll.records × avg_processing_time_per_record << max.poll.interval.ms`."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring automatically pauses/resumes consumers when `max.poll.interval.ms` bounds are threatened using `@KafkaListener` threads. Spring abstracts away manual `poll()` tuning, exposing these directly via `spring.kafka.consumer.max-poll-records`.
* **Golang:** Because `kafka-go` abstracts fetching heavily, this dynamic is slightly hidden. `ReaderConfig.MaxBytes` bounds the batch size instead of pure record counts. Be highly aware of blocked goroutines; blocking `FetchMessage` or `ReadMessage` loops stalls internal background commit loops, forcing a rebalance from the coordinator side.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Flipkart, Amazon — this is the #1 most common Kafka production bug in Java microservices. The symptom is consumers constantly rebalancing with no apparent code errors.

#### Indepth
**`session.timeout.ms` vs. `max.poll.interval.ms`:** `session.timeout.ms` covers the heartbeat thread (a background thread) going silent — usually due to a JVM crash. `max.poll.interval.ms` covers the main processing thread being too slow. They are independent timeouts working in parallel.
---
