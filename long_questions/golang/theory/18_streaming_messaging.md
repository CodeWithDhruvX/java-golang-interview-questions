# 🟢 Go Theory Questions: 341–360 Streaming, Messaging, and Asynchronous Processing

## 341. How do you consume messages from Kafka in Go?

**Answer:**
We use a library like `sarama` or certain higher-level wrappers like `segmentio/kafka-go`.

The standard pattern is to create a **Consumer Group**.
`reader := kafka.NewReader(...)`
We run a `for` loop: `m, err := reader.ReadMessage(ctx)`.
Crucially, checking `err` is vital. Upon success, we process the message and then **Commit the Offset**. If we don't commit, the broker thinks we failed and will re-deliver the message to another consumer, leading to duplicates.

---

## 342. How do you publish messages to a RabbitMQ topic?

**Answer:**
We use the `amqp` (streadway) library.

Steps:
1.  Connect to RabbitMQ (`amqp.Dial`).
2.  Open a Channel (`conn.Channel()`).
3.  Declare an Exchange (Topic type).
4.  Publish: `ch.Publish(exchange, routingKey, msg)`.

Unlike Kafka (which is a log), RabbitMQ is a smart broker. We must ensure we handle **Connection Churn**. If the TCP connection drops, we need a reconnection logic loop, otherwise, the `Publish` call will panic or hang. Use a persistent connection wrapper.

---

## 343. What is the idiomatic way to implement a message handler in Go?

**Answer:**
We define an interface: `type Handler interface { Handle(ctx context.Context, msg []byte) error }`.

Ideally, the handler should be **Idempotent** and **Stateless**.
It receives a context (for timeout/cancellation) and the payload.
If it returns `nil`, we ACK the message.
If it returns an error, we NACK (or retry).

We often wrap this handler with middleware for logging, metrics, and panic recovery, similar to how we wrap HTTP handlers.

---

## 344. How would you implement a worker pool pattern?

**Answer:**
A Worker Pool limits the number of concurrent tasks to prevent resource exhaustion.

Mechanically:
1.  Create a buffered channel `jobs := make(chan Job, 100)`.
2.  Start `N` goroutines (workers) that loop over `range jobs`.
3.  The main producer sends items into `jobs`.
4.  Close `jobs` when done.

This ensures that even if 1 million requests come in, only `N` (e.g., 50) run at once. The channel buffer acts as a queue. If the queue fills up, the producer blocks (Backpressure).

---

## 345. How do you use the `context` package for cancellation in streaming apps?

**Answer:**
Context is the kill switch for streams.

When we start a stream processor, we pass it a `ctx`.
In the processing loop (e.g., consuming Kafka), we must check `ctx.Done()`:

```go
select {
case msg := <-stream:
    process(msg)
case <-ctx.Done():
    return // Stop cleanly
}
```

This allows for **Graceful Shutdown**. When we receive a SIGTERM, we cancel the parent context. All consumers notice the cancellation, finish their current message, commit offsets, and exit the loop without data loss.

---

## 346. How do you retry failed messages in Go?

**Answer:**
We use **Exponential Backoff** for transient errors (DB momentarily down).

If a message fails repeatedly, we don't block the queue forever (Head-of-Line Blocking).
We move it to a **Retry Topic** (with a delay, often using a "TTL exchange" in RabbitMQ).
After N retries, we move it to a **Dead Letter Queue (DLQ)** and alert a human.

In Go, we can use libraries like `cenkalti/backoff` to handle the math of increasing wait times (1s, 2s, 4s...) effortlessly.

---

## 347. What is dead-letter queue and how do you use it?

**Answer:**
A DLQ is a "Graveyard" for valid messages that simply could not be processed (e.g., malformed JSON, business logic violation, or persistent bug).

We never silently drop messages. If `json.Unmarshal` fails, we publish the raw message to the DLQ topic with metadata (OriginalTopic, ErrorReason).
This allows engineers to inspect the DLQ later, fix the bug in the code, and "Replay" the messages to recover the lost data.

---

## 348. How do you handle idempotency in message consumers?

**Answer:**
Idempotency means "processing the same message twice = same result."

In distributed systems, mostly you get "At Least Once" delivery (duplicates happen).
To handle this:
1.  **Natural Idempotency**: `UPDATE users SET status='active'` is safe to run twice.
2.  **Deduplication ID**: Each message has a UUID.
    We check Redis: `SETNX message_id 1`. If it returns false, we already processed it—skip.
    Or we use a DB Unique Constraint on `insert into transactions (id, ...)`.

---

## 349. How do you implement exponential backoff in Go?

**Answer:**
Naive retries (`time.Sleep(1s)`) hammer a struggling service.
Exponential backoff (`1s, 2s, 4s, 8s`) gives it breathing room.

```go
delay := baseDelay
for i := 0; i < maxRetries; i++ {
    err := tryOp()
    if err == nil { return nil }
    time.Sleep(delay)
    delay *= 2
    // Add Jitter!
}
```

Adding **Jitter** (randomness) is crucial. If 1,000 workers fail simultaneously and all retry exactly at 2.000 seconds, they create a "Thundering Herd." Jitter spreads them out (1.9s, 2.1s).

---

## 350. How do you stream logs to a file/socket in real-time?

**Answer:**
We treats logs as an `io.Writer`.

To stream to a file:
`f, _ := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY...)`.
`logger.SetOutput(f)`.

To stream to a socket (e.g., Logstash):
`conn, _ := net.Dial("tcp", "logstash:5000")`.
`logger.SetOutput(conn)`.

However, writing to a network socket can block the app if the receiver is slow. So we typically use an **Async Wrapper** (a channel buffer) or delegate this entirely to the OS/Docker (write to stdout) to avoid freezing the app logic.

---

## 351. How do you work with WebSockets in Go?

**Answer:**
Standard `net/http` doesn't support WebSockets well. We use `gorilla/websocket` or `nhooyr.io/websocket`.

The flow:
1.  Receive HTTP GET.
2.  **Upgrade** connection: `upgrader.Upgrade(w, r, nil)`. This hijacks the TCP connection.
3.  Enter a read/write loop.

Since one goroutine is needed per connection, 100k users = 100k goroutines. This is fine in Go (unlike Node or Java threads) but requires careful memory management per goroutine (e.g., reusing buffers).

---

## 352. How do you handle bi-directional streaming in gRPC?

**Answer:**
Bi-directional streaming allows client and server to send messages independently over a single TCP connection.

In Go, the handler receives a `stream` object.
We spawn two goroutines:
1.  `go func() { for { stream.Recv() ... } }` (Read Loop)
2.  `go func() { for { stream.Send() ... } }` (Write Loop)

We coordinate them using a `done` channel. If either side closes the stream (EOF) or errors out, we close the channel to stop the other goroutine, ensuring no leaks.

---

## 353. What is Server-Sent Events and how is it done in Go?

**Answer:**
SSE is a one-way channel from Server to Browser (simpler than WebSockets).

We set headers:
`w.Header().Set("Content-Type", "text/event-stream")`
`w.Header().Set("Cache-Control", "no-cache")`

Then we loop:
`fmt.Fprintf(w, "data: %s\n\n", message)`
`w.(http.Flusher).Flush()`

The flush is critical—it forces the data out of the buffer immediately so the client sees it in real-time. It’s perfect for "Live Tickers" or "Notification Feeds."

---

## 354. How do you manage fan-in/fan-out channel patterns?

**Answer:**
**Fan-Out**: Multiple workers reading from one channel. Used to parallelize CPU work.
`for i:=0; i<3; i++ { go worker(jobs) }`.

**Fan-In**: Multiplexing multiple channels into one.
We verify `sync.WaitGroup`. We spawn a goroutine for each input channel that reads and forwards to a single output channel. Once all input routines finish, we close the output channel. This consolidates data streams (e.g., fetching data from Twitter, FB, and Reddit simultaneously).

---

## 355. How would you implement throttling on async tasks?

**Answer:**
We use a **Token Bucket** or simply a **Buffered Channel** as a semaphore.

If we want max 10 concurrent DB writes:
`sem := make(chan struct{}, 10)`

Before launching a task: `sem <- struct{}{}` (Acquire token).
After task: `<-sem` (Release token).

If the channel is full, the sender blocks. This creates natural backpressure, preventing us from overwhelming the database, without complex logic.

---

## 356. How do you avoid data races when consuming messages?

**Answer:**
The most common race is processing messages in parallel but updating a shared map or counter.

Solution 1: **Mutex**. `mu.Lock(); data[id]++; mu.Unlock()`.
Solution 2: **Sharding**. Launch 10 workers. Worker `Hash(ID) % 10` always handles user `ID`. This ensures that User A's data is always processed sequentially by Worker 3, eliminating the need for locks while still allowing parallelism across *different* users.

---

## 357. How would you implement a message queue from scratch in Go?

**Answer:**
At its core, a naive MQ is just a Slice + Mutex + Cond (for notifying waiters).

Struct: `Queue { items []T, mu sync.Mutex, cond *sync.Cond }`.
**Enqueue**: Lock, append, Signal(), Unlock.
**Dequeue**: Lock, while empty { Wait() }, remove first, Unlock.

For a production-grade one (like NATS), you need disk persistence (WAL - Write Ahead Log) and offset tracking, but the in-memory part is essentially a synchronized buffer.

---

## 358. How do you implement ordered message processing in Go?

**Answer:**
Kafka guarantees order *per partition*, not globally.

In Go, if we just spin up 100 concurrent workers for a topic, we lose that ordering (Worker 2 might finish offset 5 before Worker 1 finishes offset 4).

To preserve order: **Key-Based Routing**.
We read the message. We hash the `Key` (e.g., UserID).
We send it to a specific buffered channel: `workerChans[hash(key) % numWorkers] <- msg`.
This ensures User 123's events are processed serially by the same goroutine, preserving order, while User 456 is processed in parallel.

---

## 359. How do you handle large stream ingestion (100K+ msgs/sec)?

**Answer:**
1.  **Batching**: Don't write to DB 1 by 1. Accumulate 1,000 msgs or wait 500ms, then `INSERT INTO ... VALUES (...), (...)`.
2.  **Zero Allocation**: Reuse buffers with `sync.Pool`. Don't creating new byte slices for every message.
3.  **Async Commit**: In Kafka, committing every offset is slow. Commit asynchronously every N seconds or N messages.
4.  **Parallelism**: Ensure GOMAXPROCS is utilized. Use the Fan-Out pattern to decode/validate messages in parallel before batching for persistence.

---

## 360. How do you persist in-flight streaming data?

**Answer:**
If the app crashes, memory is lost. We need **Checkpoints**.

We maintain a **Write-Ahead Log (WAL)** on disk (like SQLite or just an append-only file).
1.  Receive Msg.
2.  Write to WAL.
3.  ACK to sender.
4.  Process asynchronously.

On startup, we replay the WAL.
For simpler apps, we rely on the Broker (Kafka/RabbitMQ) retention. We only ACK the message *after* full processing is done. If we crash mid-process, we never ACKed, so the broker redelivers it to the new instance.
