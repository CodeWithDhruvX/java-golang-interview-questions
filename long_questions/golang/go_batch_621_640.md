## 🟤 Event-Driven, Pub/Sub & Messaging (Questions 621-640)

### Question 621: How do you publish and consume events using NATS in Go?

**Answer:**
NATS is a high-performance messaging system written in Go.
**Library:** `github.com/nats-io/nats.go`

```go
// Connect
nc, _ := nats.Connect(nats.DefaultURL)

// Subscribe
nc.Subscribe("foo", func(m *nats.Msg) {
    fmt.Printf("Received: %s\n", string(m.Data))
})

// Publish
nc.Publish("foo", []byte("Hello"))
```

---

### Question 622: How do you use Apache Kafka in Go with sarama?

**Answer:**
**Sarama** is the standard low-level Go client.
- **Producer:** `sarama.NewSyncProducer`. Send `ProducerMessage`.
- **Consumer:** `sarama.NewConsumer`. Consume `PartitionConsumer`.
For Consumer Groups (handling rebalancing), it's more complex; requires implementing `ConsumerGroupHandler`.

---

### Question 623: What are the trade-offs between RabbitMQ and Kafka in Go apps?

**Answer:**
- **RabbitMQ:** "Smart Broker, Dumb Consumer". Great for complex routing, job queues, delivery guarantees per message. Push-based.
- **Kafka:** "Dumb Broker, Smart Consumer". Great for high throughput streaming, event sourcing, log replay. Pull-based.

---

### Question 624: How do you manage message acknowledgements in Go consumers?

**Answer:**
Always defer acknowledgement (Ack) until *after* processing succeeds.
If processing fails (panic/error), sending **Nack** (Negative Ack) tells the broker to redeliver the message.
In Kafka: `session.MarkMessage(msg, "")`.

---

### Question 625: How do you handle message deduplication in Go?

**Answer:**
Since most brokers guarantee "At Least Once", duplicates happen.
**Fix (Idempotency):**
1.  Extract `MessageID`.
2.  Check Redis `SETNX processed:messageID 1`.
3.  If set returns 0 (exists), skip processing.
4.  Set TTL on the key to avoid infinite growth.

---

### Question 626: How do you implement a retry queue for failed messages?

**Answer:**
**Dead Letter Exchange (DLX)** pattern (RabbitMQ) or separate topics (Kafka).
1.  Consume from `TOPIC_MAIN`.
2.  If fail -> Publish to `TOPIC_RETRY_1` (with 5m delay/TTL).
3.  If fail again -> `TOPIC_RETRY_2`.
4.  Finally -> `TOPIC_DLQ` (Manual intervention).

---

### Question 627: How do you batch message processing efficiently in Go?

**Answer:**
Read from channel/kafka and buffer until `BatchSize` or `Timeout` reached.

```go
var batch []Msg
timer := time.NewTimer(time.Second)
for {
    select {
    case msg := <-ch:
        batch = append(batch, msg)
        if len(batch) >= 100 { Flush(batch); batch = nil; timer.Reset(time.Second) }
    case <-timer.C:
        if len(batch) > 0 { Flush(batch); batch = nil }
        timer.Reset(time.Second)
    }
}
```

---

### Question 628: How do you use Google Pub/Sub with Go?

**Answer:**
Use the official SDK `cloud.google.com/go/pubsub`.
It abstracts polling.

```go
sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
    Process(msg.Data)
    msg.Ack()
})
```
By default, it spawns many goroutines to handle messages concurrently.

---

### Question 629: How do you persist event logs for replay in Go?

**Answer:**
**Event Sourcing.**
Store every event payload in an append-only table (Postgres/Cassandra) or Kafka.
To replay: Read events from `Timestamp=0` and feed them into the projection/handler logic to rebuild the state.

---

### Question 630: How do you ensure exactly-once delivery in Go message systems?

**Answer:**
True "Exactly-Once" is hard.
Typically achieved via **Idempotent Consumers** (See Q625) + **Transactional Outbox**.
Kafka Transactions API offers exactly-once semantics for stream processing (Consume -> Process -> Produce), but simpler apps rely on idempotency at the database level.

---

### Question 631: How do you create a lightweight in-memory pub-sub system?

**Answer:**
(See Q606). Useful for decoupling components within a monolith.
Using a library like `Watermill`'s GoChannel implementation gives you a robust Pub/Sub interface that works in-memory but can start using Kafka later just by changing config.

---

### Question 632: How do you handle DLQs (Dead Letter Queues) in Go?

**Answer:**
A DLQ is where messages go to die (after max retries).
Write a separate Go tool/service ("The Janitor") that allows an admin to view these messages and "Replay" them (move back to Main Topic) after fixing the bug that caused the crash.

---

### Question 633: How do you create idempotent message consumers in Go?

**Answer:**
Dependencies: Database Transaction.
1.  Begin Tx.
2.  Update User Balance.
3.  Insert `ProcessedMessageID` (Unique Constraint).
4.  Commit.
If step 3 fails (Duplicate), Rollback. This ensures balance is updated exactly once.

---

### Question 634: How do you enforce ordering of messages?

**Answer:**
- **Kafka:** Ordering is guaranteed only **within a partition**. Ensure related messages (e.g., all events for `UserID: 123`) go to the same partition by using `UserID` as the Partition Key.
- **Go Consumer:** Must process that partition sequentially (no concurrent goroutines per partition).

---

### Question 635: How do you use channels as message queues?

**Answer:**
A buffered channel act as a queue.
Limitation: If the app crashes, data in channel is lost (Volatile).
Use only for ephemeral work distribution (e.g., job processing) not for persistent data.

---

### Question 636: How do you handle push vs pull consumers?

**Answer:**
- **Pull (Kafka/SQS):** App runs a loop, asks for data. Better flow control (backpressure).
- **Push (Webhooks/RabbitMQ):** Broker sends data. App must handle it immediately or buffer it. Risk of overwhelming the app.

---

### Question 637: How do you deal with large payloads in a messaging system?

**Answer:**
**Claim Check Pattern.**
Don't put a 10MB PDF in RabbitMQ.
1.  Upload PDF to S3. Get URL.
2.  Put URL in Message.
3.  Consumer reads Message, downloads PDF from S3.

---

### Question 638: How do you build an event sourcing system in Go?

**Answer:**
Define `type Event interface{}`.
1.  Command `CreateUser` -> Validates -> Generates `UserCreated` event.
2.  Save `UserCreated` to EventStore.
3.  Publish `UserCreated` on Bus.
4.  Projector listens -> Updates `Users` Read-Model (SQL table).

---

### Question 639: How would you test message-driven systems?

**Answer:**
**Integration Tests.**
Spin up a dockerized broker (Redpanda/RabbitMQ).
1.  Produce Test Message.
2.  Wait 1s.
3.  Assert Consumer side effect (Row in DB).
**Contract Tests:** Ensure Producer and Consumer agree on JSON schema (Pact).

---

### Question 640: What’s the role of event schemas in Go-based systems?

**Answer:**
Use **Protobuf** or **Avro** (Schema Registry).
Strong typing in Go matches well with Protobuf.
Ensures Producer doesn't change field `age` from `int` to `string` breaking the Consumer.
`protoc` generates structs for both sides.

---
