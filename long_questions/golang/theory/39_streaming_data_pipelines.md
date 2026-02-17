# 🟢 Go Theory Questions: 761–780 Streaming, Batching & Data Pipelines

## 761. How do you process large CSV files using streaming?

**Answer:**
We use `csv.NewReader` and read row-by-row.
```go
r := csv.NewReader(file)
for {
    record, err := r.Read()
    if err == io.EOF { break }
    process(record)
}
```
We never use `Wait` or `ReadAll`.
For concurrency, we push `record` to a channel and have a pool of workers process lines in parallel, then fan-in the results to a writer.

---

## 762. How do you implement backpressure in a data stream?

**Answer:**
Go Channels provide inherent backpressure if **unbuffered** or **small buffer**.
Pipeline: `Reader -> Ch1 -> Worker -> Ch2 -> Writer`.
If Writer is slow, Ch2 fills.
Worker blocks on Ch2.
Ch1 fills.
Reader blocks on Ch1.
The entire system slows down to the speed of the bottleneck (Writer), preventing OOM.

---

## 763. How do you connect Go with Apache Kafka for streaming?

**Answer:**
We typically use `sarama` or `confluent-kafka-go`.
For streaming (ETL), we use the **Consumer Group** pattern.
1.  Read Message.
2.  Transform (e.g., enrich UserID).
3.  Produce to Output Topic.
4.  Mark Offset.
We must handle **Rebalancing**: If a new consumer joins, partitions are revoked. We must flush current state/buffers before releasing the claim.

---

## 764. How do you build an ETL pipeline in Go?

**Answer:**
Structure: **Extract (Source) -> Transform (Process) -> Load (Sink)**.
Go Interface:
`type Source interface { Read() (Data, error) }`
`type Sink interface { Write(Data) error }`
We chain them with Channels.
`go Extract(ch1)`
`go Transform(ch1, ch2)`
`go Load(ch2)`
This pipeline runs in constant memory, regardless of data volume size (TB).

---

## 765. How do you handle JSONL (JSON Lines) in real-time streams?

**Answer:**
JSONL is `{"id":1}\n{"id":2}\n`.
We use `bufio.Scanner` to split by newline.
```go
scanner := bufio.NewScanner(reader)
for scanner.Scan() {
    var obj MyStruct
    json.Unmarshal(scanner.Bytes(), &obj)
    // process
}
```
Ideally, use a JSON library like `easyjson` or `fastjson` that can reuse the struct memory to reduce GC pressure during high-throughput ingestion.

---

## 766. How do you split and parallelize stream processing?

**Answer:**
**Sharding**.
If we need to process User Events in parallel but keep User actions ordered:
Calculated shard: `hash(UserID) % NumWorkers`.
We define `workers := make([]chan Event, NumWorkers)`.
The Distributor reads inputs and sends to the correct worker channel. Each worker is single-threaded, ensuring order per user, but the system processes users in parallel.

---

## 767. How do you deal with schema evolution in streaming data?

**Answer:**
We use **Protobuf** or **Avro** with a Schema Registry.
Messages carry a Schema ID.
The Go consumer fetches the schema from the registry to decode the bytes.
Forward Compatibility: New fields are ignored by old consumers.
backward Compatibility: Old fields are handled (defaults) by new consumers.
Parsing raw JSON without schema validation in a pipeline is a recipe for crashing on "Malformed Data".

---

## 768. How do you throttle input data rate?

**Answer:**
**Token Bucket** (rate limiter).
Or simple time-based:
```go
limiter := time.Tick(10 * time.Millisecond)
for line := range input {
    <-limiter
    process(line)
}
```
If we want to drop packets (Load Shedding) instead of blocking:
`select { case work <- item: ... default: dropMetric.Inc() }`.

---

## 769. How do you aggregate streaming metrics?

**Answer:**
We can't store infinity. use **Tumbling Windows** or **Sliding Windows**.
Count requests per minute.
Go: Store current minute in atomic counter.
When `time.Now()` crosses the minute boundary, push the counter to DB and reset to 0.
For approximate counts (Unique Visitors), use **HyperLogLog** (Probabilistic Data Structure) which uses constant memory (12KB) to count billions of unique items.

---

## 770. How do you implement checkpointing in Go pipelines?

**Answer:**
We can't Ack every single message (too slow).
We acknowledge periodically (Every 1000 msgs or 5 seconds).
In Go code:
`if count % 1000 == 0 { store.Checkpoint(offset) }`.
**Risk**: If crash happens between checkpoints, we will re-process up to 1000 messages on restart. The downstream logic must be Idempotent (Upsert instead of Insert) to handle this replay safely.

---

## 771. How do you persist intermediate results in streams?

**Answer:**
Use a fast KV store (**RocksDB**, **BadgerDB**) embedded or **Redis**.
Example: "Count valid clicks per Session".
1.  Read Click.
2.  `clicks = redis.Incr(sessionID)`.
3.  If `clicks >= Threshold`, emit "SessionQualified" event.
State must be external (Redis) because if the Go pod restarts, in-memory maps are lost.

---

## 772. How do you implement a rolling window average?

**Answer:**
We maintain a Ring Buffer of last N samples.
Sum = Sum - Oldest + Newest.
Avg = Sum / N.
Or using Time: buckets of 1 second.
Window (1 min) = Sum of last 60 buckets.
When second 61 arrives, we drop bucket 1 and add bucket 61. This "Sliding Window" effectively keeps the moving average precise.

---

## 773. How do you batch messages for optimized DB writes?

**Answer:**
**Micro-batching**.
Channel + Ticker.
```go
func flush(batch []Msg) {
    db.Insert(batch) // Bulk Insert
}
// Loop
select {
case msg := <-ch:
    batch = append(batch, msg)
    if len(batch) > 1000 { flush() }
case <-timer.C:
    if len(batch) > 0 { flush() }
}
```
This forces a write at least every X seconds, ensuring low latency for low-volume streams while utilizing Bulk Insert speed for high volume.

---

## 774. How do you handle poison messages in a pipeline?

**Answer:**
A Poison Message crashes the consumer (Panic or infinite retry loop).
1.  **Recover**: Defer recover() in the worker.
2.  **Validation**: If unmarshal fails, don't retry.
3.  **DLQ**: Send the raw bytes + error to a "Dead Letter Queue" topic.
4.  **Ack**: Mark the poison message as "Processed" to unblock the partition.

---

## 775. How do you optimize memory in high-throughput streams?

**Answer:**
1.  **Object Reuse**: `sync.Pool` for byte buffers and structs.
2.  **Zero Allocation Parsing**: Use `index := bytes.IndexByte` instead of `strings.Split` (which allocates new strings).
3.  **Avoid Pointers**: Use arrays/slices of values to improve cache locality and reduce GC scan overhead.

---

## 776. How do you test streaming pipelines?

**Answer:**
Unit Tests: Pass closed channels.
`input := make(chan int, 2); input <- 1; input <- 2; close(input)`.
Call `Process(input)`. Assert output.

Integration Tests: Use **Redpanda** (lightweight Kafka) or MinIO (S3) in Docker.
Produce 100 messages.
Run Go pipeline.
Check Sink bucket has expected 100 results.

---

## 777. What is the fan-out/fan-in scalability pattern in streams?

**Answer:**
Stream (Kafka) has partitions (e.g., 50).
Consumer has 50 goroutines (Fan-Out).
They process via external API calls (slow).
They write results to a specific "Merge Channel" (Fan-In).
A dedicated "Committer" goroutine reads the Merge Channel and marks offsets.
This decoupling allows high CPU parallelism while maintaining the strict offset commit logic of Kafka.

---

## 778. How do you implement Change Data Capture (CDC) in Go?

**Answer:**
We listen to the DB's replication log (Binlog/WAL).
Tools: **Debezium** (Java) is standard.
In Go: `go-mysql-org/go-mysql` allows acting as a Slave replica.
We receive "RowEvent" (Insert/Update/Delete).
We marshal this to JSON and publish to Kafka/Elasticsearch.
This allows real-time sync between SQL and NoSQL without modifying the application code used to trigger the dual-write.

---

## 779. How do you handle deduplication in infinite streams?

**Answer:**
Infinite stream means infinite keys. Can't store all in Redis.
1.  **Time Window**: "Dedupe within 1 hour". Keys expire in Redis after 1h.
2.  **Bloom Filter**: Store 1 billion keys in 1GB RAM. If Bloom says "seen", drop. (Accept small false positive rate).
3.  **Content Addressable**: Store by Hash. If Write fails (Exists), it's a dupe.

---

## 780. How do you monitor lag in a Go stream processor?

**Answer:**
Lag = Broker Latest Offset - Consumer Current Offset.
We export `consumer_lag` metric to Prometheus.
If Lag is increasing, our Consumers are slower than Producers.
Action: specific Scale up (add pods) or optimize code.
We also monitor `processing_time_per_msg` to see if a specific bad deployment slowed down the logic.
