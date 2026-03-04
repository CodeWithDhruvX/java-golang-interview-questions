# 🌊 Kafka Streams — Deep Dive Interview Q&A

> **Level:** 🟡 Intermediate to 🟣 Architect
> **Asked at:** LinkedIn, Confluent, Uber, Netflix, Flipkart, Swiggy

---

## Q1. What is Kafka Streams and how does it differ from a consumer + manual processing loop?

"Kafka Streams is a **client-side library** (not a cluster) for building real-time, stateful, fault-tolerant stream processing applications directly on top of Kafka topics. It requires no additional infrastructure — no YARN, no Spark cluster, no Flink cluster.

**Key differences vs. raw consumer loop:**

| Dimension | Raw Consumer Loop | Kafka Streams |
|---|---|---|
| State Management | You build and manage state stores (Redis, DB) yourself | Built-in RocksDB state stores, automatically backed up to Kafka changelog topics |
| Fault Tolerance | You implement checkpointing/recovery | Automatic — state is restored from changelog on restart |
| Windowing | Manually track time windows | Native tumbling, hopping, sliding, session windows |
| Joins | Write your own join logic | Declarative KStream-KTable, KStream-KStream, KTable-KTable joins |
| Topology | Imperative code | Declarative DSL or Processor API |

The practical advantage: a Kafka Streams app is just a Java/Scala JAR deployed anywhere (VM, container, Lambda). It self-scales by simply running more instances — Kafka handles partition assignment automatically."

#### Indepth
Kafka Streams uses the **Streams Partition Assignment Protocol** — each partitioned task becomes a unit of parallelism. If your app processes 12 partitions and you run 3 instances, each instance gets 4 tasks. Adding a 4th instance triggers a rebalance, redistributing tasks automatically.

---

## Q2. What is the difference between KStream and KTable?

"Both represent a continuous flow of records from a Kafka topic, but they model the data differently.

**KStream — Event Log (Append-Only):**
```
Offset 0: User:A clicked BUY
Offset 1: User:A clicked BUY
Offset 2: User:B clicked VIEW
```
Every record is an **independent event**. Two records with the same key are two separate events. A KStream behaves like a standard append-only log.

**KTable — Changelog / Current State:**
```
Key: User:A → latest value: { city: Mumbai }
Key: User:B → latest value: { city: Delhi }
```
Every new record **upserts** the current value for that key. Older values are overwritten. A KTable is like a materialized database table — keyed by Kafka message key, value represents the most recent state.

**Code example:**
```java
StreamsBuilder builder = new StreamsBuilder();

// Every click event — all events matter
KStream<String, String> clickStream = builder.stream("clicks");

// User profile — only latest value matters per user
KTable<String, String> userProfiles = builder.table("user-profiles");
```"

#### Indepth
**Under the hood:** KTable is backed by a **local RocksDB state store** on the Streams instance. When a new record arrives for a key, Kafka Streams updates the RocksDB entry. To survive crashes, every KTable write is also mirrored to an internal **changelog topic** on Kafka, from which the state store is fully rebuilt on restart.

---

## Q3. What is a GlobalKTable and when do you use it over KTable?

"A **GlobalKTable** replicates the **entire topic's data to every single instance** of the Kafka Streams application, regardless of partition assignment.

**KTable (partitioned):**
- Instance 1 holds keys A–D (partitions 0, 1)
- Instance 2 holds keys E–Z (partitions 2, 3)
- If Instance 1's KStream needs to join with a key from partition 3, it **cannot** — it doesn't own that partition.

**GlobalKTable (fully replicated):**
- Every instance holds ALL keys
- Enables **non-key-based lookups** and joins across all keys on every instance

**Classic use case — enrichment with a small reference table:**
```java
GlobalKTable<String, String> productCatalog = builder.globalTable('product-catalog');

KStream<String, Order> orders = builder.stream('orders');

// Join every order with the product details — works even if keys don't match
KStream<String, EnrichedOrder> enriched = orders.join(
    productCatalog,
    (orderId, order) -> order.getProductId(),  // key extractor from stream side
    (order, product) -> new EnrichedOrder(order, product)
);
```"

#### Indepth
**Trade-off:** GlobalKTable consumes memory proportional to the **entire topic size** on every instance. Use it only for small-to-medium sized reference data (product catalog, user roles, country lookup) — not for high-cardinality event streams. For large datasets, use a partitioned join with KTable and ensure the same partitioning key.

---

## Q4. Explain the three types of windowing in Kafka Streams.

"Windowing groups events into finite time buckets for aggregation.

**1. Tumbling Window (Fixed, Non-Overlapping):**
```
[0s–60s] → Window 1
[60s–120s] → Window 2
```
Each event belongs to exactly one window. Used for: 'count clicks per minute'.
```java
TimeWindows.ofSizeWithNoGrace(Duration.ofMinutes(1))
```

**2. Hopping Window (Fixed-Size, Overlapping):**
```
[0s–60s] → Window 1
[30s–90s] → Window 2
[60s–120s] → Window 3
```
Windows advance by a `hop` interval. An event can belong to multiple windows. Used for: 'rolling 60-second average, updated every 30 seconds'.
```java
TimeWindows.ofSizeAndGrace(Duration.ofMinutes(1), Duration.ofSeconds(5))
    .advanceBy(Duration.ofSeconds(30))
```

**3. Session Window (Dynamic, Activity-Based):**
```
User:A [10:00–10:03] → Session 1
User:A [10:20–10:22] → Session 2 (gap > inactivity timeout)
```
Windows close after a configurable period of **inactivity** per key. Merges adjacent events if gap is small. Used for: 'user session analytics'.
```java
SessionWindows.ofInactivityGapWithNoGrace(Duration.ofMinutes(5))
```"

#### Indepth
**Grace Period:** By default, Kafka Streams closes a window when the stream time (max event timestamp seen) advances past the window end. But late-arriving events (network delay) are silently dropped. The `grace` parameter holds the window open for an extra duration to absorb late data. After grace expires, late events are completely discarded.

---

## Q5. What types of stream joins does Kafka Streams support? What are their requirements?

"Kafka Streams supports three join combinations:

**1. KStream–KStream Join (Windowed):**
- Two event streams joined within a time window
- Both records must arrive within `JoinWindows.ofTimeDifferenceWithNoGrace(Duration.ofSeconds(5))`
- Use case: match payment events with order events within 5 seconds

```java
KStream<String, Payment> payments = builder.stream('payments');
KStream<String, Order> orders = builder.stream('orders');

payments.join(
    orders,
    (payment, order) -> new Receipt(payment, order),
    JoinWindows.ofTimeDifferenceWithNoGrace(Duration.ofSeconds(5))
);
```

**2. KStream–KTable Join (Non-Windowed):**
- Enriches each stream event with the *current* table value for that key
- Table side does not trigger output — only stream side triggers
- Use case: enrich order events with latest user profile

**3. KTable–KTable Join (Non-Windowed):**
- Produces an update whenever **either** table is updated
- Result is itself a KTable (current join state)
- Use case: materializing a joined view of two changelog topics

**Critical Requirement — Co-Partitioning:**
All join inputs **must** have the same number of partitions and use the same partitioning key. If `orders` has 12 partitions and `payments` has 6, the join will fail at runtime with `TopologyException`."

#### Indepth
**Left, Outer Joins:** All three join types support `leftJoin` and `outerJoin` variants. In a `leftJoin(KStream, KTable)`, if no KTable record exists for the stream key, the join still produces output with `null` for the table side — allowing you to handle missing-reference cases gracefully.

---

## Q6. What are Interactive Queries in Kafka Streams? How do you expose state to external services?

"Normally, state stores in Kafka Streams are internal. **Interactive Queries** allow external HTTP clients (or other microservices) to query the materialized state stored in RocksDB directly — without going back through Kafka.

**Use case:** A dashboard wants to show the real-time word count for a specific word — without creating another consumer.

**Step 1 — Named state store:**
```java
KTable<String, Long> wordCounts = textStream
    .flatMapValues(value -> Arrays.asList(value.toLowerCase().split(' ')))
    .groupBy((key, word) -> word)
    .count(Materialized.as('word-count-store'));  // named store
```

**Step 2 — Query the store:**
```java
ReadOnlyKeyValueStore<String, Long> store = streams.store(
    StoreQueryParameters.fromNameAndType(
        'word-count-store',
        QueryableStoreTypes.keyValueStore()
    )
);
Long count = store.get('kafka');  // direct RocksDB lookup — no Kafka I/O
```

**Step 3 — Handle remote instances:**
Since state is partitioned, the key `'kafka'` might be owned by a different instance. Use `streams.metadataForKey()` to find which host owns the partition, then forward the HTTP request there:
```java
KeyQueryMetadata metadata = streams.queryMetadataForKey(
    'word-count-store', 'kafka', Serdes.String().serializer()
);
// If metadata.activeHost() != this host → proxy HTTP request
```"

#### Indepth
Interactive Queries are read-only and operate against the **local RocksDB store** — no network call to Kafka brokers involved. This is what makes them sub-millisecond. However, during a rebalance (when state is being migrated), queries may return `InvalidStateStoreException`. Production code must retry after a brief delay.

---

## Q7. How does fault tolerance and state restoration work in Kafka Streams?

"Kafka Streams achieves fault tolerance through two mechanisms:

**1. Changelog Topics:**
Every stateful operator (count, aggregate, join) automatically creates an internal Kafka topic named `<appId>-<storeName>-changelog`. Every write to the local RocksDB state store is also written as a Kafka message to this changelog. If the instance crashes, the replacement instance replays the changelog from offset 0 to restore the exact same state.

**2. Standby Replicas:**
Waiting for a full changelog replay from offset 0 on crash can take minutes for large state. Configure `num.standby.replicas=1` to maintain a **warm shadow copy** of each state store on a different instance. On failover, the standby instance is already mostly caught up — it only needs to replay the last few seconds of changelog. Failover time drops from minutes to seconds.

```java
Properties props = new Properties();
props.put(StreamsConfig.NUM_STANDBY_REPLICAS_CONFIG, 1);
```"

#### Indepth
**RocksDB checkpointing:** In addition to changelog, Kafka Streams periodically flushes the RocksDB state to a local `.checkpoint` file. On restart, it reads the checkpoint first (avoiding a full replays), then catches up from the checkpoint offset forward. This hybrid approach dramatically speeds up restore time for large state stores.

---

## Q8. Describe a real-world Kafka Streams pipeline — e.g., real-time fraud detection.

"**Problem:** Flag credit card transactions as suspicious if a user makes more than 5 transactions within a 60-second tumbling window.

**Pipeline Design:**
```java
StreamsBuilder builder = new StreamsBuilder();

// 1. Source — raw transaction events
KStream<String, Transaction> transactions = builder.stream(
    'raw-transactions',
    Consumed.with(Serdes.String(), transactionSerde)
);

// 2. Count transactions per user per 60-second window
KTable<Windowed<String>, Long> txnCountPerWindow = transactions
    .groupByKey()
    .windowedBy(TimeWindows.ofSizeWithNoGrace(Duration.ofSeconds(60)))
    .count(Materialized.as('txn-count-store'));

// 3. Filter — flag windows where count > 5
txnCountPerWindow
    .toStream()
    .filter((windowedKey, count) -> count != null && count > 5)
    .map((windowedKey, count) -> KeyValue.pair(
        windowedKey.key(),  // userId
        new FraudAlert(windowedKey.key(), count, windowedKey.window().start())
    ))
    .to('fraud-alerts', Produced.with(Serdes.String(), fraudAlertSerde));

KafkaStreams streams = new KafkaStreams(builder.build(), props);
streams.start();
```

**Flow:**
```
raw-transactions topic
      ↓
  Count per userId per 60s window (RocksDB)
      ↓
  filter(count > 5)
      ↓
fraud-alerts topic → Notification Service → Block Card
```"

#### Indepth
**Suppress operator:** Without `suppress()`, for every transaction that arrives, a partial count is emitted downstream (e.g., count=1, count=2... count=6). The fraud alert service would receive multiple updates per window. Use `.suppress(Suppressed.untilWindowCloses(...))` to emit **only once per window** after it closes and the grace period expires. This is critical for preventing duplicate fraud alerts.
