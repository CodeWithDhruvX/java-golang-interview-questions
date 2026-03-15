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

### Explanation
NATS is a high-performance messaging system written in Go. The nats.go library provides simple connection, subscription, and publishing APIs. Subscriptions use callback functions that receive messages, and publishing is straightforward with topic strings and byte payloads. NATS emphasizes simplicity and performance over complex features.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you publish and consume events using NATS in Go?
**Your Response:** "I use NATS in Go with the nats.go library. First, I connect to the NATS server using `nats.Connect()`. For consuming events, I subscribe to topics using `nc.Subscribe()` with a subject pattern and a callback function that receives messages. For publishing, I use `nc.Publish()` with the subject and byte payload. The beauty of NATS is its simplicity - it's lightweight, fast, and the Go client is very straightforward. The callback-based subscription model makes it easy to handle incoming messages asynchronously. NATS is perfect for high-performance scenarios where I need simple pub-sub without the complexity of message brokers like RabbitMQ or Kafka."

---

### Question 622: How do you use Apache Kafka in Go with sarama?

**Answer:**
**Sarama** is the standard low-level Go client.
- **Producer:** `sarama.NewSyncProducer`. Send `ProducerMessage`.
- **Consumer:** `sarama.NewConsumer`. Consume `PartitionConsumer`.
For Consumer Groups (handling rebalancing), it's more complex; requires implementing `ConsumerGroupHandler`.

### Explanation
Sarama is the standard low-level Go client for Apache Kafka. Producers use NewSyncProducer to send ProducerMessage objects. Consumers use NewConsumer and consume from PartitionConsumers. Consumer Groups require more complex implementation with ConsumerGroupHandler interface for handling rebalancing and partition assignment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Apache Kafka in Go with sarama?
**Your Response:** "I use Apache Kafka in Go with the Sarama library, which is the standard low-level client. For producers, I create a `NewSyncProducer` and send `ProducerMessage` objects to topics. For consumers, I create a `NewConsumer` and consume from `PartitionConsumer` instances. For more advanced scenarios with consumer groups that handle rebalancing automatically, I need to implement the `ConsumerGroupHandler` interface, which is more complex but provides automatic partition management. Sarama gives me fine-grained control over Kafka, though it requires understanding concepts like partitions, offsets, and consumer groups. For simpler use cases, I might consider higher-level libraries that abstract away some of this complexity."

---

### Question 623: What are the trade-offs between RabbitMQ and Kafka in Go apps?

**Answer:**
- **RabbitMQ:** "Smart Broker, Dumb Consumer". Great for complex routing, job queues, delivery guarantees per message. Push-based.
- **Kafka:** "Dumb Broker, Smart Consumer". Great for high throughput streaming, event sourcing, log replay. Pull-based.

### Explanation
RabbitMQ uses a smart broker approach with complex routing capabilities and per-message delivery guarantees, operating in push mode. Kafka uses a dumb broker approach with high throughput streaming capabilities, requiring smart consumers that pull data. RabbitMQ excels at complex routing and job queues, while Kafka excels at streaming and event sourcing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the trade-offs between RabbitMQ and Kafka in Go apps?
**Your Response:** "The key trade-off is 'Smart Broker, Dumb Consumer' versus 'Dumb Broker, Smart Consumer'. RabbitMQ is the smart broker - it handles complex routing, job queues, and provides per-message delivery guarantees. It pushes messages to consumers, making it simpler on the consumer side. Kafka is the dumb broker - it just stores streams of data and requires smart consumers that pull data and manage their own offsets. Kafka excels at high-throughput streaming, event sourcing, and log replay scenarios. I choose RabbitMQ when I need complex routing or reliable job queues, and Kafka when I need high-volume streaming or event sourcing. The choice really depends on whether I need broker-side intelligence or consumer-side flexibility."

---

### Question 624: How do you manage message acknowledgements in Go consumers?

**Answer:**
Always defer acknowledgement (Ack) until *after* processing succeeds.
If processing fails (panic/error), sending **Nack** (Negative Ack) tells the broker to redeliver the message.
In Kafka: `session.MarkMessage(msg, "")`.

### Explanation
Message acknowledgements in Go consumers should be deferred until after successful processing. If processing fails, sending a Nack tells the broker to redeliver the message. In Kafka, message marking is done with session.MarkMessage(). This ensures reliable message processing and proper handling of failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage message acknowledgements in Go consumers?
**Your Response:** "I manage message acknowledgements by only acknowledging messages after successful processing. I defer the acknowledgement until the very end of my processing logic to ensure it only happens if everything completes successfully. If processing fails due to a panic or error, I send a Nack (Negative Acknowledgement) which tells the broker to redeliver the message. In Kafka, I use `session.MarkMessage()` to mark the message as processed. This approach ensures reliable message processing - if my consumer crashes or fails to process a message, the broker knows to redeliver it. It's crucial for achieving the delivery guarantees that message brokers provide and prevents message loss during processing failures."

---

### Question 625: How do you handle message deduplication in Go?

**Answer:**
Since most brokers guarantee "At Least Once", duplicates happen.
**Fix (Idempotency):**
1.  Extract `MessageID`.
2.  Check Redis `SETNX processed:messageID 1`.
3.  If set returns 0 (exists), skip processing.
4.  Set TTL on the key to avoid infinite growth.

### Explanation
Message deduplication in Go handles the at-least-once delivery guarantee of most brokers. Idempotency is achieved by extracting MessageID, using Redis SETNX to check if already processed, skipping duplicates, and setting TTL to prevent infinite growth. This ensures each message is processed only once regardless of delivery duplicates.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle message deduplication in Go?
**Your Response:** "I handle message deduplication by implementing idempotent consumers since most brokers guarantee at-least-once delivery. I extract a unique MessageID from each message and use Redis SETNX to check if I've already processed it. If SETNX returns 0, the key already exists, so I skip processing. If it returns 1, I process the message normally. I set a TTL on the Redis key to prevent infinite growth of the processed message cache. This approach ensures each message is processed exactly once even if the broker delivers it multiple times. It's essential for building reliable systems that can handle the reality of at-least-once delivery semantics without causing duplicate side effects."

---

### Question 626: How do you implement a retry queue for failed messages?

**Answer:**
**Dead Letter Exchange (DLX)** pattern (RabbitMQ) or separate topics (Kafka).
1.  Consume from `TOPIC_MAIN`.
2.  If fail -> Publish to `TOPIC_RETRY_1` (with 5m delay/TTL).
3.  If fail again -> `TOPIC_RETRY_2`.
4.  Finally -> `TOPIC_DLQ` (Manual intervention).

### Explanation
Retry queues for failed messages use the Dead Letter Exchange pattern in RabbitMQ or separate topics in Kafka. Messages that fail processing are moved through retry topics with increasing delays (5 minutes, 15 minutes, etc.) and finally to a Dead Letter Queue for manual intervention after exhausting retry attempts.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a retry queue for failed messages?
**Your Response:** "I implement retry queues using the Dead Letter Exchange pattern. When a message fails processing, I move it to a retry topic with a delay - like 5 minutes for the first retry, 15 minutes for the second, and so on. If it keeps failing, I eventually move it to a Dead Letter Queue for manual intervention. In RabbitMQ, I use the built-in DLX feature with TTL and routing. In Kafka, I create separate retry topics and have consumers that handle the retry logic. This approach gives messages multiple chances to succeed while preventing them from blocking the main queue. The exponential backoff in retry delays helps with transient issues, and the DLQ ensures problematic messages don't get lost forever."

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

### Explanation
Batch message processing in Go uses buffering until either a batch size or timeout is reached. A timer ensures partial batches are processed within time limits, while the size limit prevents excessive memory usage. This pattern improves throughput by processing multiple messages together rather than individually.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you batch message processing efficiently in Go?
**Your Response:** "I batch message processing by buffering messages until either a batch size limit or timeout is reached. I maintain a batch slice and a timer. As messages come in, I add them to the batch. If the batch reaches my target size, I flush it immediately and reset the timer. The timer ensures that even small batches get processed within a reasonable time window. This approach significantly improves throughput by processing multiple messages together rather than one at a time, reducing database calls and API requests. It's especially effective for high-volume systems where individual message processing would be inefficient. The combination of size and time limits gives me control over both throughput and latency."

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

### Explanation
Google Pub/Sub in Go uses the official cloud.google.com/go/pubsub SDK which abstracts away polling complexity. The Receive method handles message delivery with concurrent processing through multiple goroutines. Messages are acknowledged with msg.Ack() after successful processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Google Pub/Sub with Go?
**Your Response:** "I use Google Pub/Sub with the official `cloud.google.com/go/pubsub` SDK. The beauty of this library is that it abstracts away the polling complexity. I use the `sub.Receive()` method which handles all the message delivery for me. It spawns multiple goroutines to process messages concurrently, which gives me good performance out of the box. Inside the callback, I process the message data and call `msg.Ack()` to acknowledge it. The SDK handles all the complexity of pulling messages, managing subscriptions, and handling failures. This makes it very straightforward to build reliable Pub/Sub consumers without worrying about the underlying polling mechanics."

---

### Question 629: How do you persist event logs for replay in Go?

**Answer:**
**Event Sourcing.**
Store every event payload in an append-only table (Postgres/Cassandra) or Kafka.
To replay: Read events from `Timestamp=0` and feed them into the projection/handler logic to rebuild the state.

### Explanation
Event log persistence for replay uses event sourcing where every event payload is stored in append-only storage like PostgreSQL, Cassandra, or Kafka. Replay involves reading events from the beginning and feeding them through projection logic to rebuild current state, enabling state reconstruction and audit trails.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you persist event logs for replay in Go?
**Your Response:** "I persist event logs for replay using event sourcing. I store every event payload in append-only storage - either a database table in PostgreSQL or Cassandra, or in Kafka topics. The key is that I never update or delete events, only append new ones. To replay events, I read from the beginning (Timestamp=0) and feed each event through the same projection or handler logic that was used originally. This rebuilds the current state from scratch. This approach gives me a complete audit trail and the ability to reconstruct state at any point in time. It's incredibly powerful for debugging, auditing, and even fixing bugs by replaying events with corrected logic."

---

### Question 630: How do you ensure exactly-once delivery in Go message systems?

**Answer:**
True "Exactly-Once" is hard.
Typically achieved via **Idempotent Consumers** (See Q625) + **Transactional Outbox**.
Kafka Transactions API offers exactly-once semantics for stream processing (Consume -> Process -> Produce), but simpler apps rely on idempotency at the database level.

### Explanation
Exactly-once delivery in Go message systems is difficult to achieve truly. It's typically implemented through idempotent consumers combined with transactional outbox patterns. Kafka's Transactions API provides exactly-once semantics for stream processing, but most applications rely on database-level idempotency for simpler implementation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you ensure exactly-once delivery in Go message systems?
**Your Response:** "True exactly-once delivery is very difficult to achieve. I typically implement it through a combination of idempotent consumers and the transactional outbox pattern. For most applications, I rely on database-level idempotency where I ensure that processing the same message multiple times produces the same result. Kafka's Transactions API does offer exactly-once semantics for stream processing workflows, but it adds complexity. The transactional outbox pattern involves writing both the business change and the message to be sent in the same database transaction, then a separate process sends the messages. This approach gives me strong consistency guarantees without the complexity of distributed transactions."

---

### Question 631: How do you create a lightweight in-memory pub-sub system?

**Answer:**
(See Q606). Useful for decoupling components within a monolith.
Using a library like `Watermill`'s GoChannel implementation gives you a robust Pub/Sub interface that works in-memory but can start using Kafka later just by changing config.

### Explanation
Lightweight in-memory pub-sub systems are useful for decoupling components within a monolith. Libraries like Watermill provide GoChannel implementations that work in-memory but can be switched to Kafka by changing configuration, offering migration path from in-memory to distributed messaging.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create a lightweight in-memory pub-sub system?
**Your Response:** "I create lightweight in-memory pub-sub systems using libraries like Watermill's GoChannel implementation. This is perfect for decoupling components within a monolith without the overhead of external message brokers. The beauty of this approach is that I get a robust pub-sub interface that works in-memory initially, but when I need to scale out or migrate to a distributed system, I can switch to Kafka just by changing the configuration. This gives me a smooth migration path from in-memory to distributed messaging. It's ideal for applications that start small but may need to grow, allowing me to design with proper pub-sub patterns from the beginning without the complexity of external brokers until I actually need them."

---

### Question 632: How do you handle DLQs (Dead Letter Queues) in Go?

**Answer:**
A DLQ is where messages go to die (after max retries).
Write a separate Go tool/service ("The Janitor") that allows an admin to view these messages and "Replay" them (move back to Main Topic) after fixing the bug that caused the crash.

### Explanation
Dead Letter Queues (DLQs) in Go handle messages that have exhausted retry attempts. A separate admin service allows viewing failed messages and replaying them after fixing underlying issues. This provides manual intervention capability for problematic messages without losing them permanently.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle DLQs (Dead Letter Queues) in Go?
**Your Response:** "I handle Dead Letter Queues by creating a separate admin service I call 'The Janitor'. When messages exhaust all retry attempts and end up in the DLQ, this service allows administrators to view these failed messages, understand why they failed, and replay them back to the main topic after fixing the underlying bug. The key is that DLQ messages aren't lost forever - they're parked for manual review. The Janitor service provides a web interface or CLI tool where admins can inspect the message payload, error details, and retry count. Once the root cause is fixed, they can replay the message, which moves it back to the main topic for normal processing. This approach ensures no messages are permanently lost while providing visibility into processing failures."

---

### Question 633: How do you create idempotent message consumers in Go?

**Answer:**
Dependencies: Database Transaction.
1.  Begin Tx.
2.  Update User Balance.
3.  Insert `ProcessedMessageID` (Unique Constraint).
4.  Commit.
If step 3 fails (Duplicate), Rollback. This ensures balance is updated exactly once.

### Explanation
Idempotent message consumers in Go use database transactions to ensure exactly-once processing. The pattern involves beginning a transaction, performing business logic, inserting a processed message ID with unique constraint, and committing. If the message ID already exists, the unique constraint violation causes a rollback, preventing duplicate processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create idempotent message consumers in Go?
**Your Response:** "I create idempotent message consumers using database transactions. I begin a transaction, perform the business logic like updating a user balance, then insert the processed message ID into a table with a unique constraint. Finally, I commit the transaction. If the same message is processed again, the unique constraint on the message ID will fail, causing the entire transaction to roll back. This ensures the business logic only succeeds once per message, even if the message is delivered multiple times. The database transaction guarantees atomicity - either the entire processing succeeds or none of it does. This pattern is essential for building reliable systems that can handle the at-least-once delivery semantics of most message brokers."

---

### Question 634: How do you enforce ordering of messages?

**Answer:**
- **Kafka:** Ordering is guaranteed only **within a partition**. Ensure related messages (e.g., all events for `UserID: 123`) go to the same partition by using `UserID` as the Partition Key.
- **Go Consumer:** Must process that partition sequentially (no concurrent goroutines per partition).

### Explanation
Message ordering in Kafka is guaranteed only within partitions. Related messages must be sent to the same partition using a consistent partition key like UserID. Go consumers must process each partition sequentially without concurrent goroutines to maintain ordering guarantees.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you enforce ordering of messages?
**Your Response:** "I enforce message ordering in Kafka by ensuring related messages go to the same partition. Kafka only guarantees ordering within a partition, not across partitions. I use a consistent partition key like UserID for all events related to the same user, which ensures they're processed in order. On the consumer side, I must process each partition sequentially - I can't use concurrent goroutines for the same partition or I'll lose the ordering guarantee. This approach is perfect for use cases like financial transactions or user activity streams where order matters. The key is understanding that ordering is a partition-level concern, not a topic-level one, and designing both the producer and consumer with this constraint in mind."

---

### Question 635: How do you use channels as message queues?

**Answer:**
A buffered channel act as a queue.
Limitation: If the app crashes, data in channel is lost (Volatile).
Use only for ephemeral work distribution (e.g., job processing) not for persistent data.

### Explanation
Channels as message queues use buffered channels to hold work items. The limitation is volatility - if the application crashes, data in the channel is lost. This approach is suitable only for ephemeral work distribution where data loss is acceptable, not for persistent message storage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use channels as message queues?
**Your Response:** "I can use buffered channels as simple message queues within a Go application. A buffered channel acts as a FIFO queue where producers can send work items and consumers can receive them. However, there's an important limitation - if the application crashes, any data in the channel is lost since it's only stored in memory. Because of this volatility, I only use channels for ephemeral work distribution where it's acceptable to lose data, like job processing within a single application instance. For persistent messaging that needs to survive crashes, I use external message brokers like Kafka or RabbitMQ. Channels are great for in-process queuing but shouldn't be confused with durable message systems."

---

### Question 636: How do you handle push vs pull consumers?

**Answer:**
- **Pull (Kafka/SQS):** App runs a loop, asks for data. Better flow control (backpressure).
- **Push (Webhooks/RabbitMQ):** Broker sends data. App must handle it immediately or buffer it. Risk of overwhelming the app.

### Explanation
Pull consumers (Kafka/SQS) run loops requesting data, providing better flow control and backpressure. Push consumers (Webhooks/RabbitMQ) receive data from the broker and must handle it immediately or buffer it, with risk of being overwhelmed by high message rates.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle push vs pull consumers?
**Your Response:** "I handle push and pull consumers differently based on their characteristics. Pull consumers like Kafka and SQS run a loop that actively asks for data, which gives me better flow control and natural backpressure - if I'm slow, I just don't pull more messages. Push consumers like RabbitMQ or webhooks receive data pushed from the broker, so I have to handle it immediately or buffer it somehow. The risk with push is that the broker can overwhelm my application if I can't keep up. I prefer pull models for better control over processing rates, though push models can be simpler for basic use cases. The choice really depends on whether I need fine-grained control over message processing or can handle the incoming message rate."

---

### Question 637: How do you deal with large payloads in a messaging system?

**Answer:**
**Claim Check Pattern.**
Don't put a 10MB PDF in RabbitMQ.
1.  Upload PDF to S3. Get URL.
2.  Put URL in Message.
3.  Consumer reads Message, downloads PDF from S3.

### Explanation
Large payloads in messaging systems use the claim check pattern where large data is stored externally (like S3) and only a reference URL is placed in the message. Consumers download the actual payload from external storage when needed, avoiding message broker size limitations and performance issues.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you deal with large payloads in a messaging system?
**Your Response:** "I handle large payloads using the claim check pattern. Instead of putting a 10MB PDF directly in RabbitMQ, I upload the file to S3 and get a URL back. I then put only that URL in the message, which is small and efficient to transport. When the consumer receives the message, it reads the URL and downloads the actual file from S3. This approach avoids message broker size limitations and performance issues. It also reduces network traffic since the large payload doesn't flow through the messaging system. This pattern is essential for any system dealing with files, images, videos, or other large data that needs to be processed asynchronously."

---

### Question 638: How do you build an event sourcing system in Go?

**Answer:**
Define `type Event interface{}`.
1.  Command `CreateUser` -> Validates -> Generates `UserCreated` event.
2.  Save `UserCreated` to EventStore.
3.  Publish `UserCreated` on Bus.
4.  Projector listens -> Updates `Users` Read-Model (SQL table).

### Explanation
Event sourcing systems in Go define an Event interface and follow a pattern where commands validate and generate events, events are saved to an EventStore, published on a bus, and projectors update read models. This separates write operations from read operations and maintains a complete event history.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build an event sourcing system in Go?
**Your Response:** "I build event sourcing systems by defining an Event interface and following a clear pattern. When a command like CreateUser comes in, I validate it and generate a UserCreated event. I save this event to an EventStore which maintains the complete history, then publish it on a message bus. Separate projector components listen for events and update read models like SQL tables. This separation of write and read concerns gives me a complete audit trail and flexibility in how I present data to different use cases. The EventStore ensures I never lose events, and the projectors can be rebuilt or modified without affecting the core business logic. This pattern is powerful for complex domains where understanding the history of changes is important."

---

### Question 639: How would you test message-driven systems?

**Answer:**
**Integration Tests.**
Spin up a dockerized broker (Redpanda/RabbitMQ).
1.  Produce Test Message.
2.  Wait 1s.
3.  Assert Consumer side effect (Row in DB).
**Contract Tests:** Ensure Producer and Consumer agree on JSON schema (Pact).

### Explanation
Testing message-driven systems requires integration tests with dockerized brokers to test end-to-end message flow. Tests produce messages, wait for processing, and assert consumer side effects. Contract tests ensure producer and consumer agree on message schemas using tools like Pact.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you test message-driven systems?
**Your Response:** "I test message-driven systems primarily through integration tests. I spin up dockerized message brokers like Redpanda or RabbitMQ in my test environment. My tests produce a test message, wait a short time for processing, then assert the consumer side effects like a new row in the database. This validates the entire message flow end-to-end. I also use contract tests with tools like Pact to ensure the producer and consumer agree on the JSON schema - this prevents breaking changes when one side modifies the message format. The combination of integration tests for behavior validation and contract tests for interface compatibility gives me confidence that my message-driven system works correctly and won't break when components evolve independently."

---

### Question 640: What’s the role of event schemas in Go-based systems?

**Answer:**
Use **Protobuf** or **Avro** (Schema Registry).
Strong typing in Go matches well with Protobuf.
Ensures Producer doesn't change field `age` from `int` to `string` breaking the Consumer.
`protoc` generates structs for both sides.

### Explanation
Event schemas in Go-based systems use Protobuf or Avro with Schema Registry. Protobuf matches Go's strong typing and prevents breaking changes by generating structs for both producer and consumer. Schema evolution is controlled, preventing producers from changing field types in ways that would break consumers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the role of event schemas in Go-based systems?
**Your Response:** "Event schemas are crucial for maintaining compatibility in Go-based messaging systems. I use Protobuf or Avro with Schema Registry to define message contracts. Protobuf works particularly well with Go's strong typing - the protoc compiler generates structs for both producer and consumer sides. This prevents breaking changes where a producer might change a field like `age` from `int` to `string`, which would break the consumer. The schema acts as a contract between services, ensuring they can evolve independently without breaking each other. Schema Registry manages versioning and compatibility rules, making it safe to add optional fields or make other non-breaking changes. This approach gives me the confidence to evolve my services while maintaining system stability."

---
