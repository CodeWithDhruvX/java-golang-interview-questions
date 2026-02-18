# 🔵 **341–360: Streaming, Messaging, and Asynchronous Processing**

### 341. How do you consume messages from Kafka in Go?
"I use the `sarama` library or `segmentio/kafka-go`.

I always implement a **Consumer Group**.
This allows me to scale horizontally. I start 10 replicas of my service, and Kafka automatically assigns a subset of partitions to each.
In the code, I loop over `claim.Messages()`. Crucially, I only mark the message (`session.MarkMessage`) *after* I have successfully processed it. This ensures I never lose data if my pod crashes."

#### Indepth
Sarama's default configuration is unsafe for high reliability. You must set `Producer.RequiredAcks = WaitForAll` to ensure the broker actually wrote the message to disk. On the consumer side, enable `Rebalance.Strategy = Sticky` to reduce the "stop-the-world" pause when a new consumer joins the group.

---

### 342. How do you publish messages to a RabbitMQ topic?
"I use the `streadway/amqp` library (or the new `rabbitmq/amqp091-go`).

1.  Connect and open a **Channel**.
2.  Declare an **Exchange** (Topic).
3.  Publish with a **Routing Key** (e.g., `order.created.us`).
RabbitMQ uses this key to route the message to the correct queues.
Since AMQP connections are stateful and fragile, I always wrap the publisher in a struct that handles **automatic reconnection** seamlessly."

#### Indepth
RabbitMQ has "Publisher Confirms". When you publish, the broker sends back an ACK. If you don't wait for this ACK, you might lose messages if the broker crashes instantly (or if the TCP packet drops). Use `channel.NotifyPublish()` to listen for these confirmations for 100% durability.

---

### 343. What is the idiomatic way to implement a message handler in Go?
"I define a simple interface.
`type Handler interface { Handle(ctx context.Context, msg []byte) error }`.

My consumer code is generic: it reads bytes and calls `handler.Handle`.
If `Handle` returns an error, the consumer decides the retry strategy (Nack, Requeue, or DLQ).
This separation means I can unit test my business logic (`Handle`) without needing to mock complex Kafka/RabbitMQ structs."

#### Indepth
Decorate your handler! `LoggingMiddleware(RetryMiddleware(Handle))`. This creates a chain of responsibility. Just like HTTP middleware, this allows you to standardise observability, panic recovery, and tracing across all your async workers without polluting the business logic.

---

### 344. How would you implement a worker pool pattern?
"I use **Buffered Channels**.

`jobs := make(chan Job, 100)`.
I start `N` goroutines (workers) that range over this channel.
`for job := range jobs { process(job) }`.
When I want to process work, I push to `jobs`. If the buffer fills up, the sender blocks. This provides natural **backpressure**, preventing memory exhaustion if the workers fall behind."

#### Indepth
Worker pools are great, but `errgroup` is often simpler for "scatter-gather" workflows. If you need to process 10 files in parallel and fail if *any* of them fail, `errgroup.WithContext` manages the goroutines and cancellation propagation for you automatically.

---

### 345. How do you use the `context` package for cancellation in streaming apps?
"Context is the **Kill Switch**.

When I start a stream consumer, I pass it a context.
In the inner loop (e.g., reading from Kafka), I always `select` on `ctx.Done()`.
`select { case msg := <-stream: process(msg); case <-ctx.Done(): return }`.
This allows me to shut down the stream instantly (stop waiting for new messages) when the application receives a `SIGTERM`."

#### Indepth
For network calls inside the loop, use `context.WithTimeout`. `req, _ := http.NewRequestWithContext(ctx...)`. If the main parent `ctx` is canceled (shutdown), the HTTP request aborts immediately. This ensures your consumers don't hang for 60s finishing a request during deployment.

---

### 346. How do you retry failed messages in Go?
"I use **Exponential Backoff**.

If a message fails, I don't retry immediately. I wait: `1s, 2s, 4s, 8s`.
I use a library like `cenkalti/backoff` for this.
However, I can't block the valid messages behind the bad one forever. After 3-5 attempts, I move the failed message to a **Dead Letter Queue (DLQ)** and move on. This keeps the pipeline flowing."

#### Indepth
Don't just log retry failures—increment a metric `job_failures_total`. Integrate your Circuit Breaker here. If the destination service is down, waiting 8 seconds (backoff) 100 times is wasteful. Trip the breaker and fail fast to the DLQ immediately.

---

### 347. What is dead-letter queue and how do you use it?
"A **DLQ** is a 'parking lot' for bad messages.

If a message causes a crash or fails validation repeatedly (Poison Pill), I assume it's unprocessable.
I publish it to a separate topic `project-dlq`.
I have alarms on this queue. Later, a human (or a script) can inspect these payloads, fix the bug, and re-inject them into the main queue."

#### Indepth
Automate the DLQ "Replay". Write a small CLI tool that reads from the DLQ scope and publishes back to Main Topic. Often, the failure was transient (DB down), and simply replaying 2 hours later fixes everything without code changes.

---

### 348. How do you handle idempotency in message consumers?
"Idempotency means 'processing the same message twice has the same effect as once'.
Message queues often deliver duplicates (At-Least-Once).

I handle this by assigning a **Unique ID** to every message at source.
In the consumer, I check a store (Redis/Postgres) using `SETNX key_id`. If the key exists, I skip processing. This simple check makes the system robust against duplicate deliveries."

#### Indepth
For "Exactly-Once" processing without Redis, use **Bloom Filters** for a fast, probabilistic check before hitting the DB. Or better, design your DB schema to handle duplicates naturally: `INSERT ... ON CONFLICT DO NOTHING`. Relying on the DB constraint is the most robust deduplication.

---

### 349. How do you implement exponential backoff in Go?
"I typically use a loop with a `time.Sleep`.

`delay := 100 * time.Millisecond`
`for i := 0; i < retries; i++ { err := work(); if err == nil { return }; time.Sleep(delay); delay *= 2 }`.

Crucially, I add **Jitter** (randomness). Instead of exactly 200ms, I sleep `200ms + rand(50ms)`. This prevents the 'Thundering Herd' problem where 1000 failing instances all hit the database again at the exact same millisecond."

#### Indepth
Use the "Decorrelated Jitter" algorithm. Instead of just adding random noise, the formula `sleep = min(cap, random(base, sleep * 3))` allows the backoff to fluctuate wildly, which statistically spreads out the load much better than simple "Equal Jitter".

---

### 350. How do you stream logs to a file/socket in real-time?
"I implement the `io.Writer` interface.

If I'm writing to a slow socket (like Logstash), I wrap it in an **Async Writer**.
The logger writes to a channel (fast, non-blocking).
A background goroutine reads the channel and writes to the socket.
This ensures that a network glitch in the logging infrastructure effectively typically drops logs rather than freezing the main application logic."

#### Indepth
Use a **Ring Buffer** for the async writer. If the buffer is full (Logstash is down), the writer should *overwrite old logs* or drop new ones (Non-Blocking mode). Never let logging block the application. It is better to lose logs than to take down the service (`os.Stderr` is usually safe though).

---

### 351. How do you work with WebSockets in Go?
"I use `gorilla/websocket` (or `nhooyr.io/websocket`).

I upgrade the HTTP request to a WebSocket connection.
Then I enter a `for { conn.ReadMessage() }` loop.
**Critical detail**: `conn.WriteMessage` is **not thread-safe**. If I have multiple goroutines trying to send data to the same user, I must protect the write with a Mutex or standardizing on a single 'writer goroutine' fed by a channel."

#### Indepth
WebSockets need **Keep-Alives**. Intermediate load balancers (nginx) kill idle TCP connections after 60s. You must implement a "Ping/Pong" loop. The server sends a Ping every 30s. If the client doesn't respond with Pong within 10s, close the connection.

---

### 352. How do you handle bi-directional streaming in gRPC?
"I define the RPC: `rpc Chat(stream Note) returns (stream Note)`.

In the handler, I get a `stream` object.
I need concurrency here.
1.  I spawn a goroutine to `stream.Recv()` in a loop (handle incoming).
2.  I use the main thread (or another goroutine) to `stream.Send()` (push outgoing).
The connection stays open until one side calls `Close()`."

#### Indepth
Flow Control is automatic in gRPC (HTTP/2). If the receiver reads slower than the sender, the window fills up, and the sender blocks. However, you can check `ctx.Done()` in the sender loop to detect if the client disconnected, preventing "ghost" streams from leaking resources.

---

### 353. What is Server-Sent Events and how is it done in Go?
"**SSE** is a one-way channel from Server to Browser over standard HTTP.

I set the header `Content-Type: text/event-stream`.
Then I cast the `http.ResponseWriter` to `http.Flusher`.
I write data `fmt.Fprintf(w, "data: %s\n\n", msg)` and immediately call `flusher.Flush()`.
If I don't flush, Go buffers the response, and the client sees nothing until the buffer fills up."

#### Indepth
Browser limits! Browsers (Chrome) limit concurrent connections to the same domain (max 6). If you open 6 SSE tabs, the 7th will hang. HTTP/2 solves this (multiplexing), but if you are on HTTP/1.1, you must shard domains (`s1.api.com`, `s2.api.com`) to bypass this limit.

---

### 354. How do you manage fan-in/fan-out channel patterns?
"**Fan-Out**: I launch multiple worker goroutines reading from the **same** channel. The runtime distributes the tasks.

**Fan-In**: I have multiple result channels merging into one.
I launch a goroutine for each input channel that forwards to the output.
I use `sync.WaitGroup` to track when all inputs are closed, so I can safely close the output channel."

#### Indepth
Error handling in Fan-In is tricky. If one worker fails, do you stop everything? Use `errgroup` (again). It cancels the Context for all other workers immediately upon the first error, ensuring you fail fast and don't waste CPU on a doomed batch.

---

### 355. How would you implement throttling on async tasks?
"I use the **Semaphore** pattern with a buffered channel.

`sem := make(chan struct{}, 10)`.
Before starting a goroutine: `sem <- struct{}{}`.
Inside the goroutine (defer): `<-sem`.
This limits concurrent execution to 10. It’s much lighter than a full worker pool if the tasks are just quick computations."

#### Indepth
The `golang.org/x/sync/semaphore` package provides a Weighted Semaphore. This allows advanced limiting: "Heavy tasks take 5 slots, light tasks take 1 slot". This gives you fine-grained control over resource consumption compared to a simple channel of structs.

#### Indepth
The `golang.org/x/sync/semaphore` package provides a Weighted Semaphore. This allows advanced limiting: "Heavy tasks take 5 slots, light tasks take 1 slot". This gives you fine-grained control over resource consumption compared to a simple channel of structs.

---

### 356. How do you avoid data races when consuming messages?
"**Isolation**.
Each worker goroutine should work on its own local data.
I never share a map or slice between workers.
If they need to aggregate results (e.g., 'count total errors'), I use `atomic.AddInt64` or, better yet, send the result to a dedicated 'Aggregator' goroutine via a channel."

#### Indepth
Run the race detector (`-race`) in your Integration Tests too! Many race conditions only appear when real network latency and IO are involved. It slows down tests, but finding a data race in a payment processing pipeline is worth the CPU cycles.

---

### 357. How would you implement a message queue from scratch in Go?
"In memory: Use a buffered channel.
`queue := make(chan Job, 1000)`.

For durability: **Write-Ahead Log (WAL)**.
Before pushing to the channel, I append the job to a file on disk.
On startup, I read the file to repopulate the channel. This is basically how Kafka works fundamentally."

#### Indepth
For the "Ring Buffer" implementation, `github.com/dgraph-io/ristretto` (cache) uses a high-performance ring buffer. Study its source code. It uses `atomic` counters and power-of-two sizing (`mask = size - 1`) to avoid expensive modulo operators (`%`) during the hot path.

---

### 358. How do you implement ordered message processing in Go?
"Parallelism breaks ordering. To allow parallelism *and* ordering, I use **Sharding**.

I hash the 'Entity ID' (e.g., specific `UserID`).
`shardID := hash(userID) % numWorkers`.
Each worker has its own channel. Messages for User A always go to Worker 1. Messages for User B go to Worker 2.
This guarantees User A's events are processed in order, while allowing User A and B to run in parallel."

#### Indepth
This is the **Partition Key** strategy. Be careful of "Hot Partitions". If Justin Bieber joins your app, and all his events go to Shard 1, Shard 1 will lag while Shard 2-10 are idle. You need a "Virtual Node" consistent hashing strategy to rebalance hot keys.

---

### 359. How do you handle large stream ingestion (100K+ msgs/sec)?
"I focus on **Batching** and **Allocation reduction**.

1.  **Read Batch**: Read 100 messages at a time from Kafka.
2.  **Process Batch**: Use `sync.Pool` to reuse message objects/buffers.
3.  **Write Batch**: Insert 100 rows into Postgres in one transaction.
Processing 1 message at a time is the death of performance. Batching reduces overhead by orders of magnitude."

#### Indepth
Allocations in the hot loop kill throughput. Use `sync.Pool` to reuse the structs used for unmarshalling JSON. Also, use `json.Scanner` or `easyjson` to avoid reflection overhead. Profiling usually shows `mallocgc` as the bottleneck in stream processors.

---

### 360. How do you persist in-flight streaming data?
"I treat the Stream (Kafka) as the **Source of Truth**.

1.  Read message (offset 100).
2.  Process and write to DB.
3.  **Commit Offset 100**.
If I crash before step 3, I re-read message 100 on restart.
This is 'At-Least-Once' semantics. My DB write must be idempotent (e.g., `INSERT ON CONFLICT DO NOTHING`) to handle the duplicates."

#### Indepth
Avoid storing offset in Zookeeper/Kafka if possible. Store the offset **in the same DB transaction** as your data. `INSERT INTO users ...; UPDATE offsets SET val=101`. This effectively gives you **Exactly-Once** processing because the data and offset commit atomically.
