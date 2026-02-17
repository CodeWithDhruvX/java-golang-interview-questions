# 🟤 Go Theory Questions: 621–640 Event-Driven, Pub/Sub & Messaging

## 621. How do you publish and consume events using NATS in Go?

**Answer:**
We use `nats.go`.
**Publish**: `nc.Publish("orders.created", []byte("data"))`.
**Subscribe**: `nc.Subscribe("orders.*", func(m *nats.Msg) { ... })`.

NATS is "Fire and Forget" by default (At Most Once).
For durability, we use **JetStream**.
`js.Publish(...)`.
`js.Subscribe(..., nats.Durable("consumer-1"))`.
This ensures that if the consumer is down, the stream holds the message until it reconnects.

---

## 622. How do you use Apache Kafka in Go with `sarama`?

**Answer:**
Sarama is the low-level driver.
**Producer**:
`config := sarama.NewConfig(); config.Producer.Return.Successes = true`.
`producer.SendMessage(&sarama.ProducerMessage{Topic: "topic", Value: "val"})`.

**Consumer**:
We implement `sarama.ConsumerGroupHandler`.
`Setup`, `Cleanup`, and `ConsumeClaim(sess, claim)`.
Inside `ConsumeClaim`, we loop over `claim.Messages()`. Crucially, we must mark the message `sess.MarkMessage(msg, "")` to commit the offset back to the broker.

---

## 623. What are the trade-offs between RabbitMQ and Kafka in Go apps?

**Answer:**
**RabbitMQ (Smart Broker)**:
*   Pros: Complex routing (Topic, Fanout), Per-message Ack, Priority Queues.
*   Cons: Lower throughput. Push-based (can overwhelm consumer).
*   Go Use Case: Background jobs (Celery style), complex routing logic.

**Kafka (Dumb Broker)**:
*   Pros: Massive throughput, Replayability (Log), Persistence.
*   Cons: Strict partition ordering only.
*   Go Use Case: Event Sourcing, Analytics, high-velocity streams.

---

## 624. How do you manage message acknowledgements in Go consumers?

**Answer:**
We never Auto-Ack.
We consume the message, execute business logic (Save to DB), and *then* Ack.

```go
msg := <-messages
err := process(msg)
if err != nil {
    msg.Nack() // Requeue (RabbitMQ) or do not commit offset (Kafka)
} else {
    msg.Ack()
}
```
If the process crashes before Ack, the broker redelivers. This guarantees At-Least-Once processing.

---

## 625. How do you handle message deduplication in Go?

**Answer:**
Since we have At-Least-Once delivery, we *will* get duplicates.
Deduplication must happen at the **Consumer**.
1.  Extract `MessageID`.
2.  Check Redis: `SET message_id:processed 1 NX EX 86400`.
3.  If `SET` returns false, we recognize a duplicate and `Ack` immediately without processing.

---

## 626. How do you implement a retry queue for failed messages?

**Answer:**
We don't block the main queue.
If processing fails:
1.  Publish message to `topic-retry-1` (with a header `retry_count=1`).
2.  A separate Go consumer reads `topic-retry-1`, sleeps for 1 minute, and retries.
3.  If it fails again, publish to `topic-retry-2` (sleep 5 mins).
4.  Finally, `topic-dlq`.
This "Leveled Retry" strategy allows the main consumer to keep flying through good messages.

---

## 627. How do you batch message processing efficiently in Go?

**Answer:**
(See Q 359).
We use a **Ticker** and a **Slice**.
```go
var batch []Msg
ticker := time.NewTicker(time.Second)
for {
    select {
    case msg := <-ch:
        batch = append(batch, msg)
        if len(batch) >= 1000 { flush(batch); batch = nil }
    case <-ticker.C:
        if len(batch) > 0 { flush(batch); batch = nil }
    }
}
```
This ensures we limit DB roundtrips while respecting a max latency (1 second).

---

## 628. How do you use Google Pub/Sub with Go?

**Answer:**
It abstracts the polling loop.
`sub.Receive(ctx, func(ctx, msg) { ... })`.
**Concurrency Control**:
`sub.ReceiveSettings.MaxOutstandingMessages = 10`.
This limits how many goroutines the library spawns effectively acting as a semaphore. If we don't set this, GCP will flood our pod with 10k concurrent goroutines, potentially OOMing us.

---

## 629. How do you persist event logs for replay in Go?

**Answer:**
Ideally, the Broker (Kafka) persists them (`retention.ms = -1`).
If building manually in Go:
We append events to an **Append-Only File (AOF)** or a table `events (id, payload, created_at)`.
Replay = `SELECT * FROM events WHERE created_at > last_checkpoint ORDER BY id ASC`.
We stream these rows into the processing channel just like they were live messages.

---

## 630. How do you ensure exactly-once delivery in Go message systems?

**Answer:**
True Exactly-Once is impossible in distributed systems (Two Generals Problem).
We fake it using **Idempotency** (At-Least-Once + Deduplication).

Exceptions: Kafka Transactions (EOS).
Producer: `BeginTx`. Send `Msg A`. Send `Offset Commit`. `CommitTx`.
This ensures that the output message and the input offset commit happen atomically. If crash, both roll back.

---

## 631. How do you create a lightweight in-memory pub-sub system?

**Answer:**
We use a `struct` with a `RWMutex` and a map of channels.

```go
type Bus struct {
     subs map[string][]chan interface{}
     mu   sync.RWMutex
}
```
If strict ordering isn't required, this is enough. If we need pattern matching (`topic.*`), we generally just embed NATS instead of reinventing the wheel.

---

## 632. How do you handle DLQs (Dead Letter Queues) in Go?

**Answer:**
A DLQ is just another topic/queue.
We build a **Admin Tool** (CLI or Web UI) in Go to inspect it.
`myapp dlq list` -> Shows failed JSONs and Error Strings.
`myapp dlq replay --id 123` -> Consumes from DLQ and publishes back to `main-topic`.
This manual intervention allows us to fix the code bug that caused the failure before replaying.

---

## 633. How do you create idempotent message consumers in Go?

**Answer:**
(See Q 348 and 625).
Key Strategy: **Transactional Outbox**.
If a consumer updates the DB and sends an email:
1.  Start DB Tx.
2.  Update `User`.
3.  Insert `ProcessedMessage(msg_id)`.
4.  Commit.
If step 4 crashes, the transaction rolls back. Next time we receive the message, we retry. If step 4 succeeds but we crash before Acking, next time we receive, the `INSERT ProcessedMessage` will fail (Dup Key), so we know to skip logic and just Ack.

---

## 634. How do you enforce ordering of messages?

**Answer:**
Single Partition = Total Ordering.
We ensure related messages (e.g., all updates for `Order #100`) go to the same partition.
Producer: `msg.Key = "Order-100"`.
Broker: Hashes key -> Partition 5.
Consumer: Partition 5 is owned by *one* Go consumer instance (and processing within that instance for Partition 5 is serial).

---

## 635. How do you use channels as message queues?

**Answer:**
Buffered channels *are* queues.
`q := make(chan Task, 100)`.
Limit: It's in-memory. If app restarts, data loss.
Use case: passing work between goroutines (e.g., HTTP Handler -> Background Worker).
For any data that must survive a restart, use Redis or Kafka, not channels.

---

## 636. How do you handle push vs pull consumers?

**Answer:**
**Push** (RabbitMQ/NATS): The broker forces data onto the client.
Go: We must handle flow control (QoS/Prefetch Count), otherwise the internal memory buffer explodes.
**Pull** (Kafka/Redis Stream): The client asks "Give me 10".
Go: Easier to control backpressure. We just stop asking if we are overloaded.

---

## 637. How do you deal with large payloads in a messaging system?

**Answer:**
Brokers hate large blobs (Kafka limit 1MB default).
Pattern: **Claim Check**.
1.  Upload 50MB PDF to S3.
2.  Get URL `s3://bucket/file.pdf`.
3.  Publish Message: `{ "event": "FileUploaded", "url": "..." }`.
Consumer downloads S3 URL. This keeps the broker lean and fast.

---

## 638. How do you build an event sourcing system in Go?

**Answer:**
State = Sum(Events).
`AccountBalance` is not a column; it's `Sum(Deposits) - Sum(Withdrawals)`.

Write: Append `DepositEvent` to `events` table (immutable).
Read: Can't parse 1M events every time. We use **Snapshots**.
Every 100 events, we calculate the balance and save it to a `snapshots` table.
Go: Read latest Snapshot + Replay events *after* that snapshot.

---

## 639. How would you test message-driven systems?

**Answer:**
We use **Testcontainers** to spin up real Kafka/RabbitMQ.
1.  Start Kafka container.
2.  Produce message.
3.  Wait for Go consumer to write to DB (using `Eventually` assertions).
Using generic `interface MessageQueue` mocks usually hides serialization bugs. Integration tests are mandatory here.

---

## 640. What’s the role of event schemas in Go-based systems?

**Answer:**
Distributed systems drift if schemas aren't enforced.
We use **Protobuf** or **Avro** Schema Registry.
Producer compiles `UserCreated.proto`. Consumer compiles `UserCreated.proto`.
If Producer changes a field type, the build fails (or Schema Registry rejects the update). This prevents "Poison Messages" (Sender sending JSON that Receiver cannot parse).
