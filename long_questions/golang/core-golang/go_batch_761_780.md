## 🔄 Streaming, Batching & Data Pipelines (Questions 761-780)

### Question 761: How do you process large CSV files using streaming?

**Answer:**
Read line-by-line using `encoding/csv` with `Read()` inside a loop.
Do NOT use `ReadAll()`.
Memory usage remains constant (size of one row) even for 100GB files.

### Explanation
Large CSV file processing in Go uses streaming with encoding/csv and Read() in a loop rather than ReadAll(). This maintains constant memory usage (size of one row) regardless of file size, enabling processing of very large files without memory exhaustion.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you process large CSV files using streaming?
**Your Response:** "I process large CSV files using streaming with Go's `encoding/csv` package. I read the file line-by-line using `Read()` inside a loop instead of using `ReadAll()`. This approach keeps memory usage constant - only the size of one row is in memory at any time, even for 100GB files. I create a CSV reader, then loop calling `Read()` which returns one record at a time. This is crucial for processing large files because `ReadAll()` would try to load the entire file into memory, which would cause out-of-memory errors. The streaming approach allows me to process arbitrarily large files with minimal memory footprint. I can process millions of rows efficiently while keeping memory usage predictable and low."

---

### Question 762: How do you implement backpressure in a data stream?

**Answer:**
Use unbuffered or bounded channels.
If the Consumer is slow, the Producer blocks on `ch <- data`.
This halts the reading of the source file/socket until buffer space clears.

### Explanation
Backpressure in data streams is implemented using unbuffered or bounded channels. When consumers are slow, producers block on channel sends, automatically throttling the reading of source data until buffer space becomes available, preventing memory overload.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement backpressure in a data stream?
**Your Response:** "I implement backpressure using unbuffered or bounded channels. When the consumer is slow and can't keep up, the producer blocks on the channel send operation `ch <- data`. This naturally throttles the reading of the source file or socket until the buffer space clears. If I use an unbuffered channel, the producer blocks until the consumer receives the data. With a bounded channel, the producer blocks when the buffer is full. This automatic throttling prevents the producer from reading data faster than the consumer can process it, which would otherwise cause memory to grow indefinitely. It's a simple yet effective way to handle rate mismatches between producers and consumers in streaming applications."

---

### Question 763: How do you connect Go with Apache Kafka for streaming?

**Answer:**
Use `Sarama` or `segmentio/kafka-go`.
Use a `Reader` in a loop.
Use `Consumer Groups` to allow parallel processing across multiple app instances.

### Explanation
Apache Kafka integration in Go uses libraries like Sarama or segmentio/kafka-go with Readers in loops. Consumer Groups enable parallel processing across multiple application instances, automatically balancing topic partitions among consumers for scalable stream processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you connect Go with Apache Kafka for streaming?
**Your Response:** "I connect Go with Apache Kafka using libraries like `Sarama` or `segmentio/kafka-go`. I create a `Reader` and run it in a continuous loop to consume messages from Kafka topics. For scalable processing, I use `Consumer Groups` which allow multiple instances of my application to process different partitions of the same topic in parallel. Kafka automatically balances the partitions among the consumers in the group. If one consumer fails, Kafka reassigns its partitions to the remaining consumers. This gives me both fault tolerance and horizontal scalability. The consumer group mechanism ensures that each message is processed exactly once across all instances, making it perfect for distributed stream processing applications."

---

### Question 764: How do you build an ETL pipeline in Go?

**Answer:**
**Extract:** Read from Source (SQL/API).
**Transform:** Send to Channel A. Workers read A, process, send to Channel B.
**Load:** Writer reads Channel B, batches updates, writes to Destination (Data Warehouse).

### Explanation
ETL pipelines in Go follow Extract-Transform-Load pattern: Extract reads from sources like SQL/API, Transform uses channels with workers processing data between stages, and Load batches updates to destinations like data warehouses. Channels provide natural buffering and parallelism.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build an ETL pipeline in Go?
**Your Response:** "I build ETL pipelines following the classic Extract-Transform-Load pattern using Go's concurrency features. In the Extract phase, I read data from sources like SQL databases or APIs. For Transform, I send raw data to Channel A where multiple worker goroutines read from A, process the data, and send results to Channel B. This gives me parallel processing capability. In the Load phase, a writer reads from Channel B, batches the updates for efficiency, and writes to the destination like a data warehouse. The channels provide natural buffering between stages and allow me to scale each phase independently. If the transform stage is slow, I can add more workers. If loading is the bottleneck, I can increase batch sizes. This pipeline design is both efficient and scalable."

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

### Explanation
JSONL (JSON Lines) real-time stream processing uses json.Decoder with a loop checking dec.More(). Each iteration decodes one JSON object, processes it, and continues, enabling stream processing of newline-delimited JSON without loading entire files into memory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle JSONL (JSON Lines) in real-time streams?
**Your Response:** "I handle JSONL streams using Go's `json.Decoder`. I create a decoder from my input stream and loop while `dec.More()` returns true, indicating there are more JSON objects to read. In each iteration, I declare a variable for my message type, call `dec.Decode(&m)` to read one JSON object, and then process it immediately. This approach processes one JSON object at a time without loading the entire stream into memory, making it perfect for real-time processing of large JSONL files or streaming APIs. The decoder handles the newline delimiters automatically, so I don't need to manually split lines. This pattern works great for log processing, event streams, or any newline-delimited JSON data source."

---

### Question 766: How do you split and parallelize stream processing?

**Answer:**
**Sharding.**
Hash the "Key" (UserID) % N (Workers).
Send message to `workerChans[hash]`.
Ensures order per User, but parallelism across Users.

### Explanation
Stream processing parallelization uses sharding where keys like UserID are hashed modulo number of workers. Messages are sent to worker-specific channels, ensuring order per key while achieving parallelism across different keys, balancing consistency and throughput.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you split and parallelize stream processing?
**Your Response:** "I parallelize stream processing using sharding. I hash a key like UserID modulo the number of workers, then send each message to the corresponding worker channel. This ensures that all messages for the same user go to the same worker, maintaining order per user, while different users are processed in parallel across workers. This approach gives me both consistency and scalability - users see their events in the correct order, but I get parallelism across different users. The hash function distributes the load evenly, and if I need more throughput, I can simply add more workers. This pattern is essential for high-throughput streaming systems where maintaining order per entity while scaling horizontally is crucial."

---

### Question 767: How do you deal with schema evolution in streaming data?

**Answer:**
Use **Avro** or **Protobuf** with a Schema Registry.
Producer registers schema `v1`.
Consumer checks registry.
If Producer switches to `v2` (backward compatible), Consumer can still read.

### Explanation
Schema evolution in streaming data uses Avro or Protobuf with Schema Registry. Producers register schema versions, consumers check the registry, and backward-compatible schema changes allow consumers to read new schemas without breaking existing functionality.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deal with schema evolution in streaming data?
**Your Response:** "I handle schema evolution using Avro or Protobuf with a Schema Registry. When a producer starts, it registers its schema version with the registry. Consumers check the registry to get the latest schema before processing data. If the producer updates to schema v2 with backward-compatible changes like adding optional fields, existing consumers can still read the new data. The Schema Registry maintains all schema versions and handles compatibility checks. This approach allows me to evolve my data schemas over time without breaking existing consumers. I can add fields, change data types, or restructure data as long as I maintain backward compatibility. This is crucial for long-running streaming systems where data requirements change over time but I can't afford to break existing applications."

---

### Question 768: How do you throttle input data rate?

**Answer:**
`rate.Limiter` (Token Bucket).
`limiter.Wait(ctx)` before reading next item.

### Explanation
Input data rate throttling uses rate.Limiter with token bucket algorithm. limiter.Wait(ctx) before reading the next item ensures controlled consumption rate, preventing overwhelming downstream systems or exceeding API limits.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you throttle input data rate?
**Your Response:** "I throttle input data rate using Go's `rate.Limiter` which implements a token bucket algorithm. Before reading each item from the input stream, I call `limiter.Wait(ctx)` which blocks until a token is available. This ensures I consume data at a controlled rate rather than overwhelming downstream systems or hitting API rate limits. I can configure the limiter with a specific rate like 100 requests per second. If data arrives faster than my configured rate, the limiter automatically buffers the excess and releases it at the specified rate. This approach is perfect for controlling data ingestion rates, preventing service overload, and ensuring fair resource usage when processing high-volume data streams."

---

### Question 769: How do you aggregate streaming metrics?

**Answer:**
**Tumbling Window:** `time.Ticker(1 min)`. Aggregate counts in map. On tick, flush map, reset.
**Sliding Window:** More complex. Use a Ring Buffer of 60 buckets (1 sec each). Sum all buckets for last minute.

### Explanation
Streaming metrics aggregation uses tumbling windows with time.Ticker for fixed-time bucket aggregation, or sliding windows with ring buffers of time buckets. Tumbling windows are simpler while sliding windows provide more granular time-based analytics.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you aggregate streaming metrics?
**Your Response:** "I aggregate streaming metrics using different windowing patterns. For tumbling windows, I use a `time.Ticker` that fires every minute. I aggregate counts in a map, and when the ticker fires, I flush the results and reset the map. This gives me metrics for fixed time periods like 'requests per minute'. For sliding windows, it's more complex - I use a ring buffer of 60 buckets representing each second, and continuously update the current bucket. To get the last minute's total, I sum all buckets. Sliding windows give me metrics like 'requests in the last 60 seconds' that update every second. The choice depends on whether I need fixed-time boundaries or rolling time periods. Tumbling windows are simpler and more efficient, while sliding windows provide more real-time visibility."

---

### Question 770: How do you implement checkpointing in Go pipelines?

**Answer:**
Track `LastProcessedID`.
Every N items or T seconds, write `LastProcessedID` to a persistent store (Redis/File).
On Restart: Read ID, Seek stream to ID + 1.

### Explanation
Checkpointing in Go pipelines tracks LastProcessedID and periodically writes it to persistent storage like Redis or files. On restart, the pipeline reads the ID and seeks to the next unprocessed item, ensuring no data loss and preventing reprocessing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement checkpointing in Go pipelines?
**Your Response:** "I implement checkpointing by tracking the `LastProcessedID` and periodically saving it to persistent storage like Redis or a file. Every N items or T seconds, I write the current ID to the checkpoint store. When the pipeline restarts, I read the last processed ID and seek the stream to ID + 1 to resume from where I left off. This ensures I don't lose any data during restarts and don't reprocess items that were already completed. The checkpoint frequency is a trade-off - more frequent checkpoints provide better recovery but add overhead. For critical systems, I might checkpoint every few seconds. For less critical pipelines, checkpointing every thousand items might be sufficient. This pattern is essential for building reliable data pipelines that can recover gracefully from failures."

---

### Question 771: How do you persist intermediate results in streams?

**Answer:**
If pipeline is `A -> B -> C`.
Use a persistent queue (Kafka/Redis List) between stages, not just Go channels.
If Stage B crashes, data is safely waiting in Queue A->B.

### Explanation
Intermediate results persistence in streams uses persistent queues like Kafka or Redis Lists between pipeline stages instead of Go channels. If a stage crashes, data remains safely in the queue, ensuring no data loss and enabling recovery without restarting the entire pipeline.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you persist intermediate results in streams?
**Your Response:** "I persist intermediate results using persistent queues between pipeline stages instead of just Go channels. If I have a pipeline A -> B -> C, I use Kafka or Redis Lists between each stage rather than in-memory channels. This way, if Stage B crashes, the data from Stage A is safely waiting in the queue A->B. When Stage B restarts, it can continue processing from where it left off. In-memory channels would lose all buffered data on crash, but persistent queues provide durability. This approach makes each stage independently restartable and resilient to failures. The trade-off is added complexity and latency compared to pure in-memory channels, but for production data pipelines where data loss is unacceptable, this durability is essential."

---

### Question 772: How do you implement a rolling window average?

**Answer:**
Keep a slice of values (size N) and a `Sum`.
On new Value `V`:
`Sum = Sum - Oldest + V`
`Oldest` = Remove head.
`Append` V to tail.
`Avg` = Sum / N.

### Explanation
Rolling window average in Go maintains a slice of values with fixed size N and a running sum. For new values, the sum is updated by subtracting the oldest value and adding the new one, the oldest is removed, the new value is appended, and average is calculated as Sum/N.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a rolling window average?
**Your Response:** "I implement rolling window averages by maintaining a fixed-size slice of values and a running sum. When a new value arrives, I update the sum by subtracting the oldest value and adding the new one: `Sum = Sum - Oldest + V`. Then I remove the oldest value from the head of the slice, append the new value to the tail, and calculate the average as Sum/N. This approach is very efficient - I don't need to recalculate the sum of all values each time, just update it incrementally. The fixed-size slice ensures the window always contains exactly N values, giving me a true rolling average. This pattern is perfect for real-time metrics like moving averages of response times, error rates, or any time-series data where I need to smooth out fluctuations while staying responsive to recent changes."

---

### Question 773: How do you batch messages for optimized DB writes?

**Answer:**
(See Question 627).
Wait for `BatchSize` or `TimeLimit` before executing `INSERT ... VALUES ...`.

### Explanation
Batch message optimization for database writes waits for either BatchSize or TimeLimit before executing bulk INSERT statements with multiple VALUES, reducing database round trips and improving throughput for high-volume data ingestion.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you batch messages for optimized DB writes?
**Your Response:** "I batch database writes by accumulating messages and waiting for either a batch size limit or a time limit before executing bulk INSERT statements with multiple VALUES clauses. For example, I might wait until I have 1000 records or 5 seconds have passed, whichever comes first. Then I execute a single INSERT statement with all the accumulated values. This dramatically reduces database round trips and improves throughput compared to inserting one record at a time. The batch size and time limit are tunable parameters - larger batches are more efficient but increase latency. This pattern is essential for high-volume data ingestion where database performance is a bottleneck. It's a classic trade-off between throughput and latency that I tune based on the specific requirements of the application."

---

### Question 774: How do you stream process financial transactions in Go?

**Answer:**
**Precision:** `int64` (cents).
**ACID:** Transactions must be idempotent.
**Audit:** Log every step.
**Latency:** Low-latency GC tuning.

### Explanation
Financial transaction stream processing in Go requires precision using int64 for cents, ACID compliance with idempotent transactions, comprehensive audit logging, and low-latency GC tuning to meet strict financial requirements and regulatory compliance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream process financial transactions in Go?
**Your Response:** "I process financial transactions with several critical considerations. For precision, I use `int64` to represent amounts in cents to avoid floating-point errors. For ACID compliance, I ensure all transactions are idempotent - retrying the same transaction won't double-charge customers. I maintain comprehensive audit logs of every step for regulatory compliance and debugging. For performance, I tune the garbage collector for low latency since financial systems require predictable response times. I might use GOGC settings or arena allocation to minimize GC pauses. The combination of precise arithmetic, idempotent operations, thorough auditing, and performance tuning ensures the system meets the strict requirements of financial processing where accuracy, reliability, and regulatory compliance are non-negotiable."

---

### Question 775: How do you integrate with Apache Pulsar in Go?

**Answer:**
Use `apache/pulsar-client-go`.
Similar to Kafka but supports "Shared" subscription type (Round Robin) out of the box for competing consumers without partitions.

### Explanation
Apache Pulsar integration in Go uses apache/pulsar-client-go library. Pulsar supports Shared subscription type with round-robin message distribution for competing consumers without requiring partitions, offering different consumption patterns compared to Kafka.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate with Apache Pulsar in Go?
**Your Response:** "I integrate with Apache Pulsar using the `apache/pulsar-client-go` library. Pulsar is similar to Kafka but offers some unique features. One key advantage is the 'Shared' subscription type which provides round-robin message distribution among competing consumers without requiring partitioning. This makes it easier to implement load balancing when I don't need to worry about partition keys. The client library provides similar APIs to Kafka - I create consumers, subscribe to topics, and process messages in loops. Pulsar also offers built-in tiered storage, geo-replication, and multi-tenancy features. I choose Pulsar when I need these advanced features or when the shared subscription model fits my use case better than Kafka's partition-based approach. The integration patterns are similar to other messaging systems, making it straightforward to work with."

---

### Question 776: How do you compress/decompress streaming data?

**Answer:**
Wrap readers/writers.
`gzip.NewReader(socket)` -> `json.NewDecoder`.
Transparently decompresses on the fly.

### Explanation
Streaming data compression/decompression wraps readers and writers. Using gzip.NewReader with json.NewDecoder enables transparent on-the-fly decompression, allowing compressed streams to be processed without manual decompression steps or loading entire compressed data into memory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you compress/decompress streaming data?
**Your Response:** "I compress and decompress streaming data by wrapping readers and writers. For decompression, I wrap the socket or input stream with `gzip.NewReader()` and then pass that to `json.NewDecoder()`. This creates a chain where the gzip reader transparently decompresses data as it's being read, and the JSON decoder processes the decompressed data. The beauty of this approach is that it's completely transparent to the rest of my code - I don't need to manually decompress data or load entire compressed payloads into memory. For compression, I do the reverse - wrap the writer with `gzip.NewWriter()`. This streaming approach works with any compression format and is memory-efficient because only small chunks are processed at a time. It's perfect for network protocols or file streams where I want to reduce bandwidth usage without complicating the application logic."

---

### Question 777: How do you handle late data in streaming?

**Answer:**
**Watermarks.**
Accept data with timestamp `T` until `CurrentTime > T + MaxDelay`.
If arrived later, discard or send to "Late Stream" for manual correction.

### Explanation
Late data handling in streams uses watermarks to define acceptable delay windows. Data with timestamp T is accepted until CurrentTime exceeds T + MaxDelay. Data arriving after this window is either discarded or sent to a late stream for manual correction.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle late data in streaming?
**Your Response:** "I handle late data using watermarks, which define the maximum acceptable delay for data. I accept data with timestamp T as long as CurrentTime is less than T plus a configured MaxDelay. For example, if MaxDelay is 5 minutes, I'll accept events that are up to 5 minutes late. If data arrives after this window, I either discard it or send it to a separate 'late stream' for manual processing. This approach balances completeness with timeliness - I wait a reasonable time for late-arriving data but eventually move forward to prevent the stream from getting stuck waiting for stragglers. The MaxDelay value is a business decision based on how late data is still considered useful. For some applications, a few seconds is acceptable; for others, hours might be reasonable. The late stream allows me to manually review and potentially reprocess truly important data that arrived too late."

---

### Question 778: How do you fan-out a stream to multiple destinations?

**Answer:**
Read once.
Iterate list of targets.
Write to each.
Handle failure: If 1 target fails, do you retry all? Or DLQ that one?

### Explanation
Stream fan-out to multiple destinations reads data once, iterates through target destinations, and writes to each. Failure handling strategies include retrying all targets or sending failed items to dead-letter queues for individual retry logic.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you fan-out a stream to multiple destinations?
**Your Response:** "I fan-out streams to multiple destinations by reading each message once and then writing it to each target destination in sequence. I iterate through my list of targets - databases, APIs, files, or other systems - and write the same data to each one. The key challenge is handling failures. If one target fails, I need to decide whether to retry the entire operation for all targets or just send the failed message to a dead-letter queue for that specific target. The retry-all approach ensures all targets have consistent data but might be inefficient. The per-target DLQ approach is more efficient but requires more complex monitoring and recovery logic. I choose based on the consistency requirements - for critical data, I might retry all; for high-throughput systems, per-target DLQs work better. The fan-out pattern is essential for scenarios like data replication, event broadcasting, or maintaining multiple data stores."

---

### Question 779: How do you filter events in a stream dynamically?

**Answer:**
Load rules (e.g., "Price > 100") into memory (Rule Engine).
Compile expressions (`antonmedv/expr`).
Run `Evaluate(event)` against rules. Drop if False.

### Explanation
Dynamic event filtering in streams loads rules into memory as a rule engine, compiles expressions using libraries like antonmedv/expr, and evaluates events against rules. Events not matching rules are dropped, enabling real-time filtering without code deployment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you filter events in a stream dynamically?
**Your Response:** "I filter events dynamically by loading rules into memory as a rule engine. I store filter expressions like 'Price > 100' and 'Category == Electronics' in a rules table. At runtime, I compile these expressions using a library like `antonmedv/expr` which creates fast evaluators. For each event in the stream, I run `Evaluate(event)` against all the rules. If an event doesn't match the required rules, I drop it from the stream. This approach allows me to change filtering logic without redeploying code - I just update the rules in the database and reload them. The compiled expressions are very fast, so they don't become a bottleneck. This pattern is perfect for scenarios where business rules change frequently or where different customers need different filtering criteria. It gives me the flexibility to adapt filtering logic in real-time."

---

### Question 780: How do you manage ordered processing in Kafka consumers?

**Answer:**
Kafka ensures order in a partition.
Your Go consumer must process that partition **serially**.
Process Msg 1 -> Ack -> Process Msg 2.
Do NOT launch goroutines per message inside a partition consumer if order matters.

### Explanation
Ordered processing in Kafka consumers requires serial processing within each partition since Kafka guarantees order only at the partition level. Messages must be processed sequentially (Process Msg 1 -> Ack -> Process Msg 2) without parallel goroutines per message to maintain order.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage ordered processing in Kafka consumers?
**Your Response:** "I manage ordered processing in Kafka by understanding that Kafka only guarantees order within a single partition. My Go consumer must process each partition serially to maintain this order. I process Message 1, acknowledge it, then process Message 2, and so on. The key mistake to avoid is launching goroutines per message within a partition consumer - that would break the ordering guarantee. If I need parallelism, I scale by consuming from multiple partitions simultaneously, each with its own serial processing. This ensures that all messages for a given key (which Kafka routes to the same partition) are processed in order, while I still get overall parallelism across different keys. This pattern is essential when order matters, like financial transactions or event streams where sequence is critical for correctness."

---
