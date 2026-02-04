## 🔄 Streaming, Batching & Data Pipelines (Questions 761-780)

### Question 761: How do you process large CSV files using streaming?

**Answer:**
Read line-by-line using `encoding/csv` with `Read()` inside a loop.
Do NOT use `ReadAll()`.
Memory usage remains constant (size of one row) even for 100GB files.

---

### Question 762: How do you implement backpressure in a data stream?

**Answer:**
Use unbuffered or bounded channels.
If the Consumer is slow, the Producer blocks on `ch <- data`.
This halts the reading of the source file/socket until buffer space clears.

---

### Question 763: How do you connect Go with Apache Kafka for streaming?

**Answer:**
Use `Sarama` or `segmentio/kafka-go`.
Use a `Reader` in a loop.
Use `Consumer Groups` to allow parallel processing across multiple app instances.

---

### Question 764: How do you build an ETL pipeline in Go?

**Answer:**
**Extract:** Read from Source (SQL/API).
**Transform:** Send to Channel A. Workers read A, process, send to Channel B.
**Load:** Writer reads Channel B, batches updates, writes to Destination (Data Warehouse).

---

### Question 765: How do you handle JSONL (JSON Lines) in real-time streams?

**Answer:**
Use `json.Decoder`.
```go
dec := json.NewDecoder(reader)
for dec.More() {
    var m Message
    dec.Decode(&m) // Reads one JSON object
    process(m)
}
```

---

### Question 766: How do you split and parallelize stream processing?

**Answer:**
**Sharding.**
Hash the "Key" (UserID) % N (Workers).
Send message to `workerChans[hash]`.
Ensures order per User, but parallelism across Users.

---

### Question 767: How do you deal with schema evolution in streaming data?

**Answer:**
Use **Avro** or **Protobuf** with a Schema Registry.
Producer registers schema `v1`.
Consumer checks registry.
If Producer switches to `v2` (backward compatible), Consumer can still read.

---

### Question 768: How do you throttle input data rate?

**Answer:**
`rate.Limiter` (Token Bucket).
`limiter.Wait(ctx)` before reading next item.

---

### Question 769: How do you aggregate streaming metrics?

**Answer:**
**Tumbling Window:** `time.Ticker(1 min)`. Aggregate counts in map. On tick, flush map, reset.
**Sliding Window:** More complex. Use a Ring Buffer of 60 buckets (1 sec each). Sum all buckets for last minute.

---

### Question 770: How do you implement checkpointing in Go pipelines?

**Answer:**
Track `LastProcessedID`.
Every N items or T seconds, write `LastProcessedID` to a persistent store (Redis/File).
On Restart: Read ID, Seek stream to ID + 1.

---

### Question 771: How do you persist intermediate results in streams?

**Answer:**
If pipeline is `A -> B -> C`.
Use a persistent queue (Kafka/Redis List) between stages, not just Go channels.
If Stage B crashes, data is safely waiting in Queue A->B.

---

### Question 772: How do you implement a rolling window average?

**Answer:**
Keep a slice of values (size N) and a `Sum`.
On new Value `V`:
`Sum = Sum - Oldest + V`
`Oldest` = Remove head.
`Append` V to tail.
`Avg` = Sum / N.

---

### Question 773: How do you batch messages for optimized DB writes?

**Answer:**
(See Question 627).
Wait for `BatchSize` or `TimeLimit` before executing `INSERT ... VALUES ...`.

---

### Question 774: How do you stream process financial transactions in Go?

**Answer:**
**Precision:** `int64` (cents).
**ACID:** Transactions must be idempotent.
**Audit:** Log every step.
**Latency:** Low-latency GC tuning.

---

### Question 775: How do you integrate with Apache Pulsar in Go?

**Answer:**
Use `apache/pulsar-client-go`.
Similar to Kafka but supports "Shared" subscription type (Round Robin) out of the box for competing consumers without partitions.

---

### Question 776: How do you compress/decompress streaming data?

**Answer:**
Wrap readers/writers.
`gzip.NewReader(socket)` -> `json.NewDecoder`.
Transparently decompresses on the fly.

---

### Question 777: How do you handle late data in streaming?

**Answer:**
**Watermarks.**
Accept data with timestamp `T` until `CurrentTime > T + MaxDelay`.
If arrived later, discard or send to "Late Stream" for manual correction.

---

### Question 778: How do you fan-out a stream to multiple destinations?

**Answer:**
Read once.
Iterate list of targets.
Write to each.
Handle failure: If 1 target fails, do you retry all? Or DLQ that one?

---

### Question 779: How do you filter events in a stream dynamically?

**Answer:**
Load rules (e.g., "Price > 100") into memory (Rule Engine).
Compile expressions (`antonmedv/expr`).
Run `Evaluate(event)` against rules. Drop if False.

---

### Question 780: How do you manage ordered processing in Kafka consumers?

**Answer:**
Kafka ensures order in a partition.
Your Go consumer must process that partition **serially**.
Process Msg 1 -> Ack -> Process Msg 2.
Do NOT launch goroutines per message inside a partition consumer if order matters.

---
