# 🔄 **761–780: Streaming, Batching & Data Pipelines**

### 761. How do you process large CSV files using streaming?
"I read line-by-line using `csv.NewReader`.
I send rows to a channel.
A worker pool consumes the channel and processes rows.
I write results to an output stream.
This allows processing a 100GB CSV with 100MB of RAM."

#### Indepth
**String Interning**. CSVs often have repetitive strings (e.g., "Country: US", "Status: Active"). Parsing repeated strings allocates new memory for each row. Using `unique.Make("US")` (Go 1.23+) or a manual `map[string]string` Interner can reduce RAM usage by 50% for high-redundancy datasets.

---

### 762. How do you implement backpressure in a data stream?
"I use unbuffered or small-buffered channels.
If the consumer is slow, the channel fills.
The sender blocks on `ch <- data`.
This halts the reader (e.g., stops reading from TCP).
The entire pipeline slows down to the speed of the slowest component, preventing OOM."

#### Indepth
**Context Cancellation**. Backpressure stops the producer, but what if the user cancels the request? You must pass `ctx` through the pipeline. `select { case ch <- data: ; case <-ctx.Done(): return }`. Without this, a producer might sit blocked on a full channel forever if the consumer crashes or quits early.

---

### 763. How do you connect Go with Apache Kafka for streaming?
"I use `segmentio/kafka-go`.
`reader := kafka.NewReader(...)`.
`for { m, _ := reader.ReadMessage(ctx); process(m) }`.
It acts as a blocking stream.
I commit offsets *after* successful processing to ensure 'At Least Once' semantics."

#### Indepth
**Consumer Groups**. Typically you run replicas of your Go service. They join the same `GroupID`. Kafka automatically rebalances partitions. If a Go pod dies, its partitions are reassigned to others. You must handle the `Rebalance` event (close DB connections, flush buffers) to avoid data duplication during the handover.

---

### 764. How do you build an ETL pipeline in Go?
"**Extract**: Read from Source. Send to Chan A.
**Transform**: Workers read Chan A, modify data, send to Chan B.
**Load**: Workers read Chan B, batch insert into Target.
I run these stages concurrently using `sync.WaitGroup` to coordinate shutdown."

#### Indepth
**Error Handling**. If `Transform` fails, should the whole pipeline crash? Usually, no. Send failed items to a `DLQ` (Dead Letter Queue) channel and continue. `select { case errCh <- err: }`. The main loop logs the error and alerts, but the pipeline keeps flowing for valid data.

---

### 765. How do you handle JSONL (JSON Lines) in real-time streams?
"I use `json.Decoder`.
`dec := json.NewDecoder(reader)`.
`for dec.More() { var v Data; dec.Decode(&v); process(v) }`.
The standard decoder handles concatenated JSON objects naturally and efficiently."

#### Indepth
**SIMD Parsing**. `encoding/json` is slow (reflection). For high-throughput (GBs/sec), use `simdjson-go`. It uses AVX2 instructions to parse JSON 10x faster than standard lib. It verifies the JSON structure and pointers without fully unmarshaling it into Go structs unless requested.

---

### 766. How do you split and parallelize stream processing?
"**Sharding**.
I hash the Key (e.g., UserID).
`shardID := hash(key) % numWorkers`.
Send to `channels[shardID]`.
This ensures all events for User 123 go to the same worker (preserving order) while utilizing multiple CPU cores."

#### Indepth
**Consistent Hashing**. If you add a worker, `hash % N` changes for almost all keys, breaking ordering (User 123 moves from Worker A to B). Worker B processes a new event while Worker A is still finishing an old one. Race! Use Consistent Hashing or ensure the pipeline is fully drained/paused before resizing the worker pool.

---

### 767. How do you deal with schema evolution in streaming data?
"I use a schema registry (Avro/Protobuf).
The message includes a Schema ID.
The consumer fetches the schema.
Go's dynamic nature is limited, so I generate structs for *all* versions and try to unmarshal into the newest."

#### Indepth
**Dynamic Protobuf**. If you don't know the schema at compile time, use `dynamicpb` (Go 1.18+). It allows inspecting a Protobuf message using a "FileDescriptor" loaded at runtime. This is how generic CLI tools (like `grpcurl`) work without needing your specific `.proto` files compiled in.

---

### 768. How do you throttle input data rate?
"I use `rate.Limiter`.
`limiter.Wait(ctx)`.
This sleeps the reader if the rate is exceeded.
This protects downstream systems from spikes (e.g., when backfilling data)."

#### Indepth
**Token Bucket**. `rate.Limiter` implements Token Bucket. It allows bursts. If you have a limit of 10/sec and silence for 5 seconds, the bucket fills. The next request might burst 50 items instantly. If you need rigid "Spacing" (pacing), use `time.Ticker`, but Token Bucket is usually better for overall throughput.

---

### 769. How do you aggregate streaming metrics?
"I use **Tumbling Windows** or **Sliding Windows**.
Keep a `map[string]int` in memory.
Every 1 minute (Ticker), flush the map to DB and reset it.
For sliding windows, I use a Ring Buffer of 60 seconds buckets."

#### Indepth
**Data Loss**. In-memory aggregation is fast but risky. If the pod crashes before the minute flush, you verify data. Solution: **Write-Ahead Log**. Write raw events to a local disk Append-Only File before aggregating in RAM. On restart, replay the AOF to restore the map state.

---

### 770. How do you implement checkpointing in Go pipelines?
"I save the state (Offset / Cursor) periodically.
Every 1000 items, I write `LatestProcessedID` to Postgres/Redis.
On restart, I read this ID and resume fetching from `ID + 1`.
I ensure this save is atomic with the data write if possible."

#### Indepth
**Atomic Commit**. Processing a Kafka message often changes DB state *and* updates the offset. If one fails, you get duplication. Pattern: Store the "Last Offset" *in the database transaction itself* (idempotency table). "Process Data + Update Offset" happens in one Postgres TX. Kafka offset commit is just an optimization then.

---

### 771. How do you persist intermediate results in streams?
"I use a local KV store (Badger) or Redis.
Stateful processing (e.g., 'Sum of last 5 events') requires memory.
If I crash, I lose memory.
So I update the KV store on every event. On startup, I load the last state."

#### Indepth
**RocksDB**. For massive local state (TB of data), `Badger` or `RocksDB` (via CGO) is standard. They use LSM trees optimized for high write throughput. This allows a Go service to act like a stateful stream processor (like Flink) without needing an external database roundtrip for every event.

---

### 772. How do you implement a rolling window average?
"I keep a slice of timestamps and values.
Add new entry.
Remove entries older than `Now - WindowSize`.
Calculate Average of remaining entries.
For high frequency, I use an **Exponential Moving Average (EMA)** which only needs to store one float."

#### Indepth
**T-Digest**. Calculating "Average" is easy. Calculating "99th Percentile" on a stream is hard (needs all values). Use probabilistic data structures like **T-Digest** or **HdrHistogram**. They estimate P99 with high accuracy using very little memory and can be merged from multiple streams.

---

### 773. How do you batch messages for optimized DB writes?
"I use the **Micro-Batching** pattern.
Channel collects items.
A loop with `Ticker` (500ms) and `Limit` (1000 items).
`select { case item := <-ch: batch = append(batch, item); if len >= Limit { flush() } case <-ticker.C: flush() }`.
This balances latency vs throughput."

#### Indepth
**Copying**. `batch = append(batch, item)` reuses the underlying array. When you flush, you must pass a *copy* to the DB worker or set `batch = nil` before re-appending. If you reuse the same slice header while the DB worker reads it, you get race conditions and data corruption.

---

### 774. How do you stream process financial transactions in Go?
"**Precision is key**. I use `big.Int`, never floats.
I stick to **ACID**.
I process sequentially per Account (Sharding).
I use a Write-Ahead Log (WAL) or Postgres Row Locking (`SELECT FOR UPDATE`) to prevent double spending."

#### Indepth
**Saga Pattern**. For transactions spanning services (Bank A -> Bank B), ACID doesn't work. Use Sagas. Go orchestrator calls "Debit A". If successful, calls "Credit B". If B fails, calls "Compensate A" (Undo). This ensures eventual consistency without distributed locks (Two-Phase Commit), which scale poorly.

---

### 775. How do you integrate with Apache Pulsar in Go?
"I use `pulsar-client-go`.
It’s similar to Kafka but supports 'Shared' subscription (Round Robin).
`consumer.Receive(ctx)`.
It handles acking individually.
I use Pulsar for Work Queues rather than Stream processing."

#### Indepth
**Key_Shared**. Pulsar's killer feature. It allows multiple consumers to read from the *same* partition, but guarantees that all messages with Key X go to Consumer A. This combines the ordering guarantees of Kafka partitions with the dynamic scaling of RabbitMQ. Kafka cannot do this (1 partition = 1 consumer max).

---

### 776. How do you compress/decompress streaming data?
"I wrap the `Reader`/`Writer`.
`gz := gzip.NewWriter(file)`.
`json.NewEncoder(gz).Encode(data)`.
This compresses on the fly.
For reading: `gz, _ := gzip.NewReader(conn)`.
This is transparent to the business logic."

#### Indepth
**Snappy/Zstd**. Gzip is CPU heavy. For internal streams (Server to Server), use **Snappy** (Google) or **Zstd** (Facebook). They offer lower compression ratios but are 10x-50x faster to encode/decode, which reduces the CPU bottleneck on high-throughput data pipelines.

---

### 777. How do you handle late data in streaming?
"I define a **Watermark**.
If an event arrives with `Timestamp < CurrentTime - 1Hour`, I drop it (or send to Side Output).
I can't wait forever for out-of-order data."

#### Indepth
**Triggers**. What if data arrives *after* the window closes? You can configure a "Late Trigger". Update the aggregation and emit a "Correction" event. The downstream system must handle updates (e.g., overwriting a previous value in the DB). This allows for "Eventual Correctness" even with delayed data.

---

### 778. How do you fan-out a stream to multiple destinations?
"I launch N goroutines.
I **Duplicate** the message.
Source -> FanOut -> Chan A -> S3.
                 -> Chan B -> ElasticSearch.
The FanOut loop sends to both. Caution: If Chan A blocks, FanOut blocks. Buffer properly!"

#### Indepth
**TeeReader**. `io.TeeReader(reader, writer)` splits a byte stream. Everything read from `reader` is automatically written to `writer`. Great for "Tap" middleware: read the HTTP Body to parse it, but also Tee it to a file logger for audit purposes without consuming it twice.

---

### 779. How do you filter events in a stream dynamically?
"I verify the filter condition (maybe a compiled Regex or CEL expression).
`if !filter.Match(event) { continue }`.
I reload the filter config dynamically so I don't need to restart the pipeline to change rules."

#### Indepth
**CEL-Go**. Google's **Common Expression Language** (CEL) is safer and faster than full scripting (Lua/JS) for filters. It evaluates innocent expressions (`event.type == "login" && event.risk > 50`) efficiently and cannot crash the runtime or loop forever. It is the industry standard for dynamic config rules.

---

### 780. How do you manage ordered processing in Kafka consumers?
"Kafka guarantees order per Partition.
I must process messages from *one* partition sequentially.
I cannot launch concurrent goroutines for the *same* partition.
However, I can launch 1 goroutine per Partition."

#### Indepth
**OutOfOrder Commits**. If you parallelize within a partition (Workers A, B, C for Offset 1, 2, 3), you can't commit Offset 3 until 1 and 2 are done. You must track a "Low Watermark". Maintain a heap of "Completed Offsets". Only commit the contiguous block. If 2 finishes but 1 fails, you can't commit > 0.
