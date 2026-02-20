# 🟤 **621–640: Event-Driven, Pub/Sub & Messaging**

### 621. How do you publish and consume events using NATS in Go?
"I use the `nats.go` library.
Publish: `nc.Publish("orders.created", []byte("data"))`.
Subscribe: `nc.Subscribe("orders.*", func(m *nats.Msg) { ... })`.
It’s fire-and-forget. Extremely fast (millions of msgs/sec).
If I need persistence (at-least-once delivery), I use **JetStream**."

#### Indepth
**Request-Reply Pattern**. NATS excels here. `nc.Request("help", []byte("me"), 1*time.Second)`. It creates a temporary inbox (subscription), sends the request with the `Reply-To` header set to that inbox, waits for the response, and then cleans up. It makes a distributed system feel like a function call.

---

### 622. How do you use Apache Kafka in Go with `sarama`?
"I create a `sarama.ConsumerGroup`.
I implement the `ConsumerGroupHandler` interface (`Setup`, `Cleanup`, `ConsumeClaim`).
The loop reads from `claim.Messages()`.
I verify to mark offsets *after* processing (`session.MarkMessage(msg, "")`) to ensure I don't process the same message twice if the pod restarts."

#### Indepth
**Rebalance Storms**. If your processing logic takes too long (`> session.Timeout`), the broker thinks the consumer is dead and triggers a Rebalance (stop-the-world). To fix this, either optimize logic or increase `MaxProcessingTime` and use background context cancellation to abort processing when a rebalance starts.

---

### 623. What are the trade-offs between RabbitMQ and Kafka in Go apps?
"**RabbitMQ**: Smart broker. Good for complex routing (topics, fanout) and task queues. Push-based. Harder to replay old messages.
**Kafka**: Dumb broker, smart consumer. Good for high throughput stream processing. Pull-based. Retains history (log), allowing replay.
I choose Rabbit for background jobs, Kafka for data pipelines."

#### Indepth
**Ordering guarantees**. RabbitMQ guarantees order within a queue. Kafka guarantees order ONLY within a *partition*. If you need global order in Kafka, you are limited to 1 partition (no scaling). RabbitMQ is often "easier" for simpler work queues where strict ordering isn't paramount or scale is moderate.

---

### 624. How do you manage message acknowledgements in Go consumers?
"Explicit Ops.
`msg.Ack()` tells the broker 'I am done'.
If I crash before Ack, the broker redelivers.
I only Ack *after* the DB commit is successful.
If DB fails, I `Nack()` (Negative Ack) or let the ack timeout expire."

#### Indepth
Manual Acks are critical. Auto-Ack (default in many libs) is dangerous; if your app crashes after receiving the message but before processing it, the message is lost forever. Always verify `AutoAck: false` in production consumers.

---

### 625. How do you handle message deduplication in Go?
"The broker usually guarantees 'At-Least-Once'.
I make my consumer **idempotent**.
I use a table `processed_messages (message_id PRIMARY KEY)`.
Transaction:
1.  Insert message_id.
2.  If error (duplicate key) -> Return Success (already done).
3.  Else -> Process logic -> Commit."

#### Indepth
**Redis SETNX**. For a lighter check (non-critical deduplication), use Redis `SETNX key "1" EX 86400`. It's faster than a SQL insert but less durable (if Redis crashes without AOF). For financial transactions, always use the SQL Transaction approach (Idempotency Key pattern).

---

### 626. How do you implement a retry queue for failed messages?
"I use distinct topics: `main`, `retry-1m`, `dlq`.
If processing fails:
Publish to `retry-1m`. Ack original.
The `retry-1m` consumer reads, checks timestamp. If < 1m has passed, it sleeps.
After 3 tries, publish to `dlq` (Dead Letter Queue) for manual intervention."

#### Indepth
**Exponential Backoff Headers**. Don't create 10 queues (`retry-1s`, `retry-2s`...). Instead, attach a header `Next-Retry-Time: <timestamp>` and republish to the *same* retry topic. The consumer reads, checks the header, and if it's too early, does a `Nack` (re-queue) with a small sleep. This keeps the topology simple.

---

### 627. How do you batch message processing efficiently in Go?
"I read from the channel into a buffer slice.
`batch := make([]Msg, 0, 100)`.
`timer := time.NewTimer(1 * time.Second)`.
Loop: select case msg: append; if len==100 { flush() } case timer: flush().
`flush()` sends a bulk insert to DB and then Acks all 100 messages. This reduces DB load by 100x."

#### Indepth
**Micro-batching Latency**. The downside is latency. The first message waits up to 1 second. You must tune the parameters (Batch Size vs Flush Interval). `size=100, time=50ms` is often a better sweet spot for near-real-time systems.

---

### 628. How do you use Google Pub/Sub with Go?
"I use `cloud.google.com/go/pubsub`.
It manages the pulling logic for me.
`sub.Receive(ctx, func(ctx, msg) { process(msg); msg.Ack() })`.
It spawns goroutines automatically for concurrency. I configure `ReceiveSettings` (MaxOutstandingMessages) to control memory usage."

#### Indepth
**Flow Control**. If you don't limit `MaxOutstandingMessages`, the library might pull 10,000 messages into RAM if your processing is slow. This causes OOM kills. Always set this to a reasonable number (e.g., `NumCPU * 20`) to match your worker pool capacity.

---

### 629. How do you persist event logs for replay in Go?
"I treat the **Log as the Database**.
In Kafka, I set `retention.ms = -1` (Infinite) for critical topics.
To replay: I start a new Consumer Group from `offset=0`.
My Go app re-processes every historical event. This allows me to rebuild my read-model (e.g., ElasticSearch index) from scratch."

#### Indepth
**Snapshotting**. Replaying 10 years of events takes too long. Periodically (e.g., every 10k events), create a "Snapshot" (current state) and save it to S3. To rebuild, load the latest Snapshot + replay only events *after* that snapshot. This is optimizing the "Recovery Time Objective" (RTO).

---

### 630. How do you ensure exactly-once delivery in Go message systems?
"It's theoretically impossible across boundaries without 2PC.
But Kafka **Transactional Producer** (`initTransactions`) enables it within Kafka.
For side-effects (DB writes), I rely on **Idempotency**.
Combining an Idempotent Consumer + At-Least-Once Delivery = Effectively Exactly-Once Processing."

#### Indepth
Kafka Transactions are heavy and complex (Zombie fencing, Group Coordinators). For 99% of use cases, Idempotency (handling duplicates gracefully) is the correct and more robust engineering solution. "Exactly-Once" is mostly a marketing term for "Exactly-Once *Effect*".

---

### 631. How do you create a lightweight in-memory pub-sub system?
"I use channels and a map of subscribers.
`type PubSub struct { subs map[string][]chan any }`.
`func (ps *PubSub) Pub(topic, val)`: non-blocking send to all channels.
This is great for decoupling components inside a Monolith (e.g., WebSocket handler subscribes to 'chat updates' from the Login handler)."

#### Indepth
**Blocking Sends**. If `Pub` blocks until all subscribers read, one slow subscriber kills the system. Make the subscriber channels buffered, or use `select { case ch <- msg: default: log.Warn("dropped") }` (Non-blocking send). Dropping messages is usually better than deadlocking the Publisher.

---

### 632. How do you handle DLQs (Dead Letter Queues) in Go?
"I write a specific **DLQ Consumer** tool.
It reads the bad messages.
It logs them or shows them in a UI dashboard.
It has a 'Reprocess' button which publishes the message back to the `main` topic.
I investigate *why* it failed (bad JSON? bug?) before clicking Reprocess."

#### Indepth
**Automated Redrive**. Sometimes failures are transient (3rd party API down for 1 hour). You can have a "Redrive Policy" that automatically moves moves messages from DLQ back to Main after X hours. But be careful of infinite loops if the message is "poison" (permanently unprocessable).

---

### 633. How do you create idempotent message consumers in Go?
"I use the **Outbox Pattern** or a **State Table**.
Also, I design operations to be naturally idempotent:
`UPDATE balance SET amount = 100` (Idempotent).
`UPDATE balance SET amount = amount + 10` (Not Idempotent! Needs deduplication logic)."

#### Indepth
**Database Constraints**. Let the DB do the work. `INSERT INTO processed_events (id) VALUES ($1) ON CONFLICT DO NOTHING`. Check `RowsAffected()`. If 0, it was a duplicate. This atomic check-and-insert is the bedrock of idempotency.

---

### 634. How do you enforce ordering of messages?
"I use **Sharding / Partitioning**.
In Kafka, ordering is guaranteed *only within a partition*.
I ensure all events for `Order-123` go to Partition 5 by using `Order-123` as the message Key.
If I use a random key, `Created` might arrive after `Shipped`."

#### Indepth
**Key Skew**. If you partition by `CustomerID`, and one customer (Amazon) does 50% of your traffic, that one partition (and the 1 consumer processing it) will be overwhelmed while others verify idle. You might need "Compound Keys" or specialized sharding strategies for "Celeb" entities.

---

### 635. How do you use channels as message queues?
"Buffered channels are queues.
`q := make(chan Job, 100)`.
But they are **ephemeral**. If the app crashes, data is lost.
I use channels for *work distribution* between goroutines, not for *storage*. For reliable queuing, I always use Redis or RabbitMQ."

#### Indepth
**Persistence vs Speed**. Channels are purely RAM. Redis is RAM + Disk. Postgres is Disk. The trade-off is Durability. If you can afford to lose the job on restart (e.g. sending a "Welcome" email), channels work. If it's a "Payment", you need Disk (Postgres/Kafka).

---

### 636. How do you handle push vs pull consumers?
"**Push** (NATS): Low latency. The broker floods me. I need rate limiting to avoid OOM.
**Pull** (Kafka/SQS): I control the rate. I ask for 10 messages.
In Go, Pull is safer for avoiding backpressure issues.
I implement Pull loops: `batch := sqs.ReceiveMessage(10); process(batch)`."

#### Indepth
**Long Polling**. When pulling from SQS, use `WaitTimeSeconds=20`. This tells SQS: "If the queue is empty, hold my connection open for 20s until a message arrives." This reduces empty API calls (and cost) by 99% compared to a tight loop.

---

### 637. What is event sourcing?
"It means storing *state changes*, not state.
`OrderCreated`, `ItemAdded`, `OrderPaid`.
My Go app reads these events to calculate `CurrentBalance`.
It allows me to answer 'What was the balance last Tuesday?' by replaying events up to that timestamp. The complexity is high, so I use it sparingly."

#### Indepth
**CQRS Relationship**. Event Sourcing almost always requires CQRS. Writing to the Event Store is fast (Append Only). But *Reading* "All users with name like 'Bob'" is impossible. You need a separate process that reads events and updates a standard SQL/NoSQL table (The "Read Model") for queries.

---

### 638. How do you handle schema evolution in event-driven systems?
"I use **Schema Registry** (Avro/Protobuf).
I enable compatibility modes (Forward/Backward).
Producer cannot publish a message that breaks existing Consumers (e.g., removing a required field).
This contract enforcement prevents 'poison pill' messages."

#### Indepth
**Schema Registry**. The Producer shouldn't send the schema with every message (too big). Instead, it registers the schema (ID: 5) and sends `[ID:5][BinaryData]`. The Consumer downloads Schema 5 from the Registry to decode. Confluent Schema Registry is the standard for Kafka.

---

### 639. How do you implement a competing consumers pattern?
"I place multiple consumers (in a Group) on the same queue.
The broker delivers each message to **only one** consumer.
NATS Queue Groups or Kafka Consumer Groups handle this.
It allows horizontal scaling: if processing is slow, I add 5 more Go replicas, and the throughput increases linearly."

#### Indepth
**Partition Limit**. In Kafka, you can't have more consumers than partitions. If you have 10 partitions, the 11th consumer sits idle. This is a hard limit on horizontal scaling. Plan your partition count (e.g. 50 or 100) upfront based on expected future scale.

---

### 640. How do you monitor message lag?
"**Consumer Lag** is the metric.
`Lag = LatestOffset - CurrentOffset`.
If Lag is growing, my consumers are too slow.
I export this to Prometheus using `kafka-exporter`.
I alert if `Lag > 10000` or `Latency > 1 minute`."

#### Indepth
**Burrow**. Tools like LinkedIn's Burrow monitor lag *without* being part of the consumer group. They inspect offsets independently. This prevents "Observer Effect" where monitoring slows down the consumer. Lag is the single best metric for "Is my system healthy?".
