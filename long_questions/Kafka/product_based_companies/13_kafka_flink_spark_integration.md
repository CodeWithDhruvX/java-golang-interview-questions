# ⚡ Kafka — Integration with Apache Flink & Spark Streaming

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, LinkedIn, Uber, Amazon, Hotstar, PhonePe, Razorpay (data engineering rounds)

---

## Q1. When do you choose Kafka + Flink over Kafka Streams?

"Both Kafka Streams and Flink read from Kafka and write back to Kafka. The choice is architectural:

**Use Kafka Streams when:**
- You want **zero extra infrastructure** — it runs inside your Java service (no cluster to manage)
- Your processing is moderate — per-message enrichment, windowed counts, KTable joins
- Your team is Java-focused and wants one deployable JAR
- Latency requirement is p99 < 100ms and event volume is < 500K events/sec per service

**Use Kafka + Flink when:**
- You need **complex multi-stream joins** across 5+ topics with sub-second latency
- You're running **ML inference pipelines** — Flink integrates with PyFlink and TensorFlow Serving
- You need **exactly-once guarantees across heterogeneous sinks** (Kafka + Cassandra + PostgreSQL in one atomic transaction)
- Your event volume exceeds what a single Kafka Streams app can handle (10M+ events/sec requiring distributed processing across a Flink cluster)
- You need **Complex Event Processing (CEP)** — detecting patterns across events over time (e.g., 'alert if 3 failed logins within 60 seconds followed by a password reset')

**Use Kafka + Spark Streaming when:**
- You already have a Spark cluster and want to add streaming to existing batch jobs
- Your use case tolerates **micro-batch latency** (1–10 seconds is acceptable)
- You need **unified batch + streaming code** (Spark Structured Streaming handles both with the same DataFrame API)
- The team is Python/Scala-heavy and the ML team uses PySpark"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Kafka Streams is natively embedded. Flink and Spark require external clusters entirely separated from your Spring Boot microservices.
* **Golang:** Neither Kafka Streams, Spark, nor Flink are native to Go. Go teams often use Benthos for stream manipulation, or offload complex CEP to Flink SQL while Go purely acts as ingest/egress.

#### Indepth
**The operational cost difference is real.** Kafka Streams is the cheapest to operate — it runs inside your existing pods. Flink requires a dedicated cluster (JobManager + TaskManagers) or a managed service (AWS Kinesis Data Analytics for Flink, Confluent Flink, Google Dataflow). Spark requires a Spark cluster or managed EMR/Databricks. Factor this into your architecture decision.

---

## Q2. How do you integrate Kafka with Apache Flink? Show a complete example.

"**Dependencies (Maven):**
```xml
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>flink-connector-kafka</artifactId>
    <version>3.2.0-1.19</version>
</dependency>
<dependency>
    <groupId>org.apache.flink</groupId>
    <artifactId>flink-streaming-java</artifactId>
    <version>1.19.0</version>
</dependency>
```

**Complete Flink + Kafka Pipeline — Real-Time Fraud Detection:**
```java
StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

// --- 1. Configure Kafka Source ---
KafkaSource<Transaction> kafkaSource = KafkaSource.<Transaction>builder()
    .setBootstrapServers("kafka-broker-1:9092,kafka-broker-2:9092")
    .setTopics("raw-transactions")
    .setGroupId("flink-fraud-detector")
    .setStartingOffsets(OffsetsInitializer.latest())
    .setValueOnlyDeserializer(new TransactionDeserializationSchema())
    .build();

DataStream<Transaction> transactions = env.fromSource(
    kafkaSource,
    WatermarkStrategy.<Transaction>forBoundedOutOfOrderness(Duration.ofSeconds(5))
        .withTimestampAssigner((txn, ts) -> txn.getTimestamp()),
    "Kafka Transactions Source"
);

// --- 2. Fraud Detection Logic — CEP Pattern ---
// Detect: >3 transactions from same userId within 60 seconds
Pattern<Transaction, ?> fraudPattern = Pattern
    .<Transaction>begin("first")
    .followedByAny("repeated")
    .where(new SimpleCondition<Transaction>() {
        @Override
        public boolean filter(Transaction t) { return true; }
    })
    .times(3)
    .within(Time.seconds(60));

PatternStream<Transaction> patternStream = CEP.pattern(
    transactions.keyBy(Transaction::getUserId),
    fraudPattern
);

DataStream<FraudAlert> fraudAlerts = patternStream.select(
    (PatternSelectFunction<Transaction, FraudAlert>) pattern -> {
        List<Transaction> repeated = pattern.get("repeated");
        return new FraudAlert(repeated.get(0).getUserId(), repeated.size(), System.currentTimeMillis());
    }
);

// --- 3. Configure Kafka Sink ---
KafkaSink<FraudAlert> kafkaSink = KafkaSink.<FraudAlert>builder()
    .setBootstrapServers("kafka-broker-1:9092")
    .setRecordSerializer(KafkaRecordSerializationSchema.builder()
        .setTopic("fraud-alerts")
        .setValueSerializationSchema(new FraudAlertSerializationSchema())
        .build()
    )
    .setDeliveryGuarantee(DeliveryGuarantee.EXACTLY_ONCE)  // Flink transactional producer
    .setTransactionalIdPrefix("flink-fraud-")
    .build();

fraudAlerts.sinkTo(kafkaSink);

// --- 4. Enable Checkpointing for Exactly-Once ---
env.enableCheckpointing(10000);  // checkpoint every 10 seconds
env.getCheckpointConfig().setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
env.getCheckpointConfig().setMinPauseBetweenCheckpoints(5000);

env.execute("Fraud Detection Pipeline");
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Flink applications are typically standalone Java Jobs packaged as fat JARs, operating outside the Spring ecosystem, though Spring Data repositories can be instantiated inside Flink RichFunctions if tightly coupled dependencies are required.
* **Golang:** Flink does not support Golang natively (Java/Python only). Go services produce standard JSON/Protobuf to Kafka, which Flink consumes.

#### Indepth
**Flink's Exactly-Once with Kafka** uses a two-phase commit protocol. During checkpointing, the Flink Kafka source commits offsets as part of the checkpoint. The Flink Kafka sink uses Kafka transactional producers — the transaction is committed only when the checkpoint succeeds. If the Flink job crashes and restores from checkpoint, the source replays from the last committed offset and the pending sink transaction is aborted + retried. This achieves true end-to-end exactly-once across Kafka source → Flink processing → Kafka sink.

---

## Q3. How do you integrate Kafka with Apache Spark Structured Streaming? Show a complete example.

"**Spark Structured Streaming reads Kafka topics as a continuous DataFrame. It uses a micro-batch model by default, with an optional Continuous Processing mode for lower latency.**

**Dependencies (Maven/SBT):**
```xml
<dependency>
    <groupId>org.apache.spark</groupId>
    <artifactId>spark-sql-kafka-0-10_2.12</artifactId>
    <version>3.5.0</version>
</dependency>
```

**Complete Spark + Kafka Pipeline — Real-Time Log Analytics (Scala):**
```scala
import org.apache.spark.sql.SparkSession
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming.Trigger
import org.apache.spark.sql.types._

val spark = SparkSession.builder()
  .appName("KafkaLogAnalytics")
  .config("spark.sql.streaming.checkpointLocation", "s3://my-spark-checkpoints/log-analytics")
  .getOrCreate()

// --- 1. Read from Kafka ---
val rawLogs = spark.readStream
  .format("kafka")
  .option("kafka.bootstrap.servers", "kafka-broker-1:9092,kafka-broker-2:9092")
  .option("subscribe", "app-logs")
  .option("startingOffsets", "latest")
  .option("kafka.group.id", "spark-log-analytics")
  // Security options
  .option("kafka.security.protocol", "SSL")
  .option("kafka.ssl.keystore.location", "/etc/kafka/kafka.keystore.jks")
  .load()

// --- 2. Parse Kafka Value (binary → JSON) ---
val logSchema = StructType(Seq(
  StructField("serviceId", StringType),
  StructField("level", StringType),        // INFO, WARN, ERROR
  StructField("message", StringType),
  StructField("timestamp", LongType)
))

val parsedLogs = rawLogs
  .select(from_json(col("value").cast("string"), logSchema).as("log"))
  .select("log.*")
  .withColumn("eventTime", to_timestamp(col("timestamp").divide(1000)))

// --- 3. Windowed Aggregation — Error Rate per Service ---
val errorRates = parsedLogs
  .where(col("level") === "ERROR")
  .groupBy(
    window(col("eventTime"), "1 minute", "30 seconds"),  // 1-min sliding window, 30s slide
    col("serviceId")
  )
  .count()
  .where(col("count") > 100)  // Alert if >100 errors in any 1-min window

// --- 4. Write Alerts back to Kafka ---
errorRates
  .select(
    col("serviceId").as("key"),
    to_json(struct("window", "serviceId", "count")).as("value")
  )
  .writeStream
  .format("kafka")
  .option("kafka.bootstrap.servers", "kafka-broker-1:9092")
  .option("topic", "error-rate-alerts")
  .option("checkpointLocation", "s3://my-spark-checkpoints/error-alerts")
  .trigger(Trigger.ProcessingTime("30 seconds"))  // micro-batch every 30s
  .outputMode("update")
  .start()
  .awaitTermination()
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spark applications are primarily written in Scala or Java, running on dedicated YARN/K8s clusters.
* **Golang:** Spark has no Go SDK. Go services act strictly as producers feeding the Kafka topics that Spark Streams analyze.

#### Indepth
**Spark Checkpointing with Kafka:** Spark stores Kafka offsets as part of the checkpoint state in S3/HDFS. On restart, Spark reads the checkpoint to find the last committed offset and resumes from there. This provides **at-least-once** guarantees by default. For **exactly-once**, use `idempotent` sink writes (e.g., `MERGE` into Delta Lake/Iceberg with upsert semantics, so replayed records don't cause duplicates).

---

## Q4. What is the difference between Spark Structured Streaming trigger modes?

"Spark gives you 3 trigger modes, and choosing correctly has major cost + latency impact:

**1. Default (micro-batch, 0ms trigger):**
```scala
.trigger(Trigger.ProcessingTime(0))  // or no trigger
```
Spark launches a new batch as soon as the previous one finishes. Lowest latency micro-batch achieves (~500ms–2s e2e). Highest CPU utilization.

**2. Fixed Interval (most common in production):**
```scala
.trigger(Trigger.ProcessingTime("30 seconds"))
```
Spark waits 30 seconds, then processes all events accumulated. Reduces cluster overhead. Latency = 30 seconds. Correct for dashboards, reporting, non-urgent aggregations.

**3. Continuous Processing (experimental, < 100ms latency):**
```scala
.trigger(Trigger.Continuous("1 second"))  // checkpoint every 1 second
```
Spark continuously processes records with ~1ms latency (comparable to Flink). Not yet production-stable for all sink types. Suitable when you need Spark's DataFrame API with near-real-time latency.

**4. One-Time (batch on demand):**
```scala
.trigger(Trigger.Once())
```
Process all available data as a single batch and stop. Used for scheduled reprocessing jobs."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Both ecosystems simply see the output of these batch jobs as discrete, rapid bursts of Kafka messages pushed into sink topics. 

#### Indepth
**Cost optimization for Spark on AWS EMR/Databricks:** Use `Trigger.ProcessingTime("5 minutes")` instead of continuous processing for non-SLA-critical pipelines. This allows Spark to accumulate larger batches, improving compression ratios (Snappy/Zstd), reducing the number of small-file writes to S3, and lowering per-query Databricks DBU costs by 40–60%.

---

## Q5. How do you implement exactly-once semantics in a Kafka → Flink → Database pipeline?

"**The challenge:** Exactly-once across Kafka (source) → Flink (processing) → PostgreSQL (sink) requires a distributed commit protocol, because Kafka commits and database commits are independent systems.

**Solution: Two-Phase Commit (2PC) via Flink's TwoPhaseCommitSinkFunction:**

```java
public class PostgresSink extends TwoPhaseCommitSinkFunction<FraudAlert, Connection, Void> {

    public PostgresSink() {
        super(new KryoSerializer<>(Connection.class, new ExecutionConfig()), VoidSerializer.INSTANCE);
    }

    @Override
    protected Connection beginTransaction() throws Exception {
        // Begin a DB transaction during checkpoint phase 1
        Connection conn = DriverManager.getConnection("jdbc:postgresql://db:5432/fraud", "user", "pass");
        conn.setAutoCommit(false);
        return conn;
    }

    @Override
    protected void invoke(Connection conn, FraudAlert alert, Context ctx) throws Exception {
        // Write to DB within the open transaction (not committed yet)
        PreparedStatement stmt = conn.prepareStatement(
            "INSERT INTO fraud_alerts(user_id, count, detected_at) VALUES (?, ?, ?) ON CONFLICT DO NOTHING"
        );
        stmt.setString(1, alert.getUserId());
        stmt.setInt(2, alert.getCount());
        stmt.setTimestamp(3, new Timestamp(alert.getTimestamp()));
        stmt.execute();
    }

    @Override
    protected void preCommit(Connection conn) throws Exception {
        // Called when checkpoint is complete — pre-commit DB transaction
        // At this point Kafka offsets are also checkpointed
    }

    @Override
    protected void commit(Connection conn) {
        // Both Kafka offset checkpoint AND DB preCommit are done — now commit DB
        try { conn.commit(); conn.close(); }
        catch (Exception e) { throw new RuntimeException(e); }
    }

    @Override
    protected void abort(Connection conn) {
        // Checkpoint failed — rollback DB transaction
        try { conn.rollback(); conn.close(); }
        catch (Exception e) { throw new RuntimeException(e); }
    }
}
```

**How 2PC works:**
1. Flink checkpoint triggers → Kafka source commits offsets to `__consumer_offsets` (phase 1)
2. Flink calls `preCommit()` on DB sink — DB transaction is prepared but not committed
3. Checkpoint succeeds and is confirmed → Flink calls `commit()` on DB sink → DB commits
4. If crash between `preCommit` and `commit` → on restart Flink replays from checkpoint offset and the uncommitted DB transaction is rolled back, preventing duplicates"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** When Go or Spring Boot services read from the output topics of Flink jobs, they should set `isolation.level=read_committed` to ensure they do not read aborted transaction data from Flink's 2PC checkpoints.

#### Indepth
**The practical reality:** True exactly-once across heterogeneous systems (Kafka + external DB) requires the external DB to support 2PC or idempotent writes. SQL databases support 2PC natively. NoSQL stores like Cassandra/DynamoDB do not support 2PC — for those, design for **at-least-once + idempotent writes** using `INSERT ... IF NOT EXISTS` or upsert patterns, which achieves the same effective result with simpler code.

---

## Q6. How do you tune Kafka consumer settings specifically for Flink and Spark batch sizes?

"**Flink Kafka Consumer Tuning:**
```java
// Flink reads Kafka in its own fetch loop internally — tune via Flink source properties
KafkaSource.<Transaction>builder()
    // Fetch 50MB max per partition per fetch request (vs default 1MB)
    .setProperty("fetch.max.bytes", "52428800")          // 50MB
    .setProperty("max.partition.fetch.bytes", "10485760") // 10MB per partition
    // Reduce polling frequency to accumulate larger batches
    .setProperty("fetch.min.bytes", "65536")              // wait for at least 64KB
    .setProperty("fetch.max.wait.ms", "500")              // wait max 500ms for min bytes
    // Adjust consumer heartbeat separately from processing timeout
    .setProperty("heartbeat.interval.ms", "3000")
    .setProperty("session.timeout.ms", "30000")
    // For high-throughput: allow more bytes in flight
    .setProperty("receive.buffer.bytes", "10485760")      // 10MB socket buffer
    .build();
```

**Spark Kafka Consumer Tuning:**
```scala
spark.readStream
  .format("kafka")
  // Max records per partition per micro-batch (throttle if needed)
  .option("maxOffsetsPerTrigger", 100000)
  // Max bytes fetched per partition per trigger
  .option("kafka.fetch.max.bytes", "52428800")
  .option("kafka.fetch.min.bytes", "65536")
  .option("kafka.max.partition.fetch.bytes", "10485760")
  // Avoid consumer rebalance during long batches
  .option("kafka.session.timeout.ms", "120000")
  .option("kafka.heartbeat.interval.ms", "10000")
  // Fetch from follower replicas (reduces broker leader pressure)
  .option("kafka.client.rack", "us-east-1a")
  .load()
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** These configuration properties map cleanly to native `spring-kafka` or `confluent-kafka-go` properties if building custom batch accumulators in-app instead of using Spark.

#### Indepth
**`maxOffsetsPerTrigger` is critical for Spark cost control.** Without it, a large backlog (e.g., consumer was down for 2 hours) causes Spark to try to process millions of records in a single micro-batch, OOMing executors. Setting `maxOffsetsPerTrigger=100000` guarantees each batch stays bounded, the cluster stays stable, and you work through the backlog gracefully over multiple batches.
