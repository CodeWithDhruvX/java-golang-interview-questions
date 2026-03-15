## 🔵 Streaming, Messaging, and Asynchronous Processing (Questions 341-360)

### Question 341: How do you consume messages from Kafka in Go?

**Answer:**
Use a library like `github.com/IBM/sarama` or `github.com/segmentio/kafka-go`.

**Using `kafka-go` (Simpler API):**
```go
import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)

func consume() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "my-topic",
        GroupID:  "my-group",
        MinBytes: 10e3, // 10KB
        MaxBytes: 10e6, // 10MB
    })
    defer r.Close()

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }
}
```

---

### Question 342: How do you publish messages to a RabbitMQ topic?

**Answer:**
Use `github.com/rabbitmq/amqp091-go`.

**Steps:**
1. Connect to RabbitMQ.
2. Open a channel.
3. Declare an exchange (Topic type).
4. Publish message with routing key.

**Code:**
```go
import "github.com/rabbitmq/amqp091-go"

func publish() {
    conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()

    ch, _ := conn.Channel()
    defer ch.Close()

    err := ch.ExchangeDeclare(
        "logs_topic", // name
        "topic",      // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = ch.PublishWithContext(ctx,
        "logs_topic", // exchange
        "kern.critical", // routing key
        false,        // mandatory
        false,        // immediate
        amqp091.Publishing{
            ContentType: "text/plain",
            Body:        []byte("Kernel panic!"),
        })
}
```

---

### Question 343: What is the idiomatic way to implement a message handler in Go?

**Answer:**
Use an interface-based design to decouple the handler logic from the transport (Kafka/RabbitMQ/HTTP).

**Pattern:**
```go
// 1. Define Handler Interface
type MessageHandler interface {
    Handle(ctx context.Context, msg []byte) error
}

// 2. Concrete Implementation
type OrderProcessor struct{}

func (op *OrderProcessor) Handle(ctx context.Context, msg []byte) error {
    var order Order
    json.Unmarshal(msg, &order)
    return processOrder(order)
}

// 3. Worker (Transport Layer)
func StartConsumer(handler MessageHandler, messages <-chan []byte) {
    for msg := range messages {
        go func(m []byte) {
            if err := handler.Handle(context.Background(), m); err != nil {
                log.Println("Error handling message:", err)
                // Nack/Retry logic here
            }
        }(msg)
    }
}
```
This makes unit testing the `OrderProcessor` trivial without mocking Kafka.

---

### Question 344: How would you implement a worker pool pattern?

**Answer:**
A worker pool limits concurrency to a fixed number of goroutines.

**Implementation:**
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second) // Simulate work
        results <- j * 2
    }
}

func main() {
    const numJobs = 100
    const numWorkers = 5

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // 1. Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // 2. Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // Signal no more jobs

    // 3. Collect results
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

---

### Question 345: How do you use the context package for cancellation in streaming apps?

**Answer:**
Pass `context.Context` to all long-running operations. Monitor `ctx.Done()` to stop processing immediately.

**Example:**
```go
func StreamProcessor(ctx context.Context, stream <-chan Data) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Tracking stream stopped:", ctx.Err())
            return
        case data := <-stream:
            process(data)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    go StreamProcessor(ctx, dataChannel)
    
    // ... main continues ...
}
```
If the timeout hits or `cancel()` is called, the `StreamProcessor` exits immediately, preventing goroutine leaks.

---

### Question 346: How do you retry failed messages in Go?

**Answer:**
Use **Exponential Backoff** with **Jitter**.

**Library:** `github.com/cenkalti/backoff`

**Manual Implementation:**
```go
func processWithRetry(operation func() error) error {
    maxRetries := 5
    baseDelay := 100 * time.Millisecond

    for i := 0; i < maxRetries; i++ {
        err := operation()
        if err == nil {
            return nil
        }
        
        // Exponential backoff: 100ms, 200ms, 400ms...
        delay := baseDelay * time.Duration(1<<i)
        
        // Add jitter (randomness) to prevent thundering herd
        jitter := time.Duration(rand.Intn(50)) * time.Millisecond
        
        log.Printf("Retry %d after error: %v. Waiting %v", i+1, err, delay+jitter)
        time.Sleep(delay + jitter)
    }
    return fmt.Errorf("operation failed after %d retries", maxRetries)
}
```

---

### Question 347: What is dead-letter queue (DLQ) and how do you use it?

**Answer:**
A DLQ is a standard queue where "bad" messages (that failed processing after max retries) are sent for manual inspection.

**Strategy:**
1. Consumer attempts to process message.
2. Implementation fails → Retry 3 times.
3. Still fails → Publish message to `topic-dlq`.
4. Acknowledge original message to remove it from main queue.
5. **Alerting:** Monitor DLQ depth to alert developers.

**Go Code Snippet:**
```go
if err := process(msg); err != nil {
    if retries >= maxRetries {
        // Publish to DLQ
        producer.Publish("my-topic-dlq", msg.Body)
        // Ack on main topic
        msg.Ack() 
    } else {
        retries++
        // Nack to retry later
        msg.Nack()
    }
}
```

---

### Question 348: How do you handle idempotency in message consumers?

**Answer:**
Idempotency ensures processing the same message multiple times has the same effect as processing it once.

**Strategies:**
1. **Database Uniqueness:** Use the Message ID as a Primary Key/Unique Constraint in the DB.
   - If `INSERT` fails with "Duplicate Key", ignore the message.
2. **Redis Deduplication:**
   - Check if `MessageID` exists in Redis.
   - If not, process and set `MessageID` with TTL.

**Example (DB Approach):**
```go
func processOrder(db *sql.DB, startOrder StartOrderMsg) error {
    tx, _ := db.Begin()
    
    // Check if processed
    var exists bool
    tx.QueryRow("SELECT exists(SELECT 1 FROM processed_msgs WHERE id=$1)", startOrder.MsgID).Scan(&exists)
    
    if exists {
        return nil // Already processed, safe to ack
    }

    // Process logic...
    
    // Mark as processed
    tx.Exec("INSERT INTO processed_msgs (id) VALUES ($1)", startOrder.MsgID)
    
    return tx.Commit()
}
```

---

### Question 349: How do you implement exponential backoff in Go?

**Answer:**
Wait time increases exponentially with each failure ($Base \times 2^{Attempt}$).

```go
func Retry(attempts int, sleep time.Duration, f func() error) error {
    if err := f(); err != nil {
        if attempts--; attempts > 0 {
            // Jitter for robustness
            jitter := time.Duration(rand.Int63n(int64(sleep))) / 2
            sleep = sleep + jitter/2

            time.Sleep(sleep)
            return Retry(attempts, 2*sleep, f)
        }
        return err
    }
    return nil
}
```

---

### Question 350: How do you stream logs to a file/socket in real-time?

**Answer:**
Use `io.MultiWriter` to write to multiple destinations (Console + File/Socket).

```go
func main() {
    // 1. Open file
    file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    
    // 2. Open Socket (e.g., Logstash)
    conn, _ := net.Dial("tcp", "localhost:5000")

    // 3. Create MultiWriter
    logger := log.New(io.MultiWriter(os.Stdout, file, conn), "INFO: ", log.LstdFlags)

    logger.Println("This log goes to Console, File, and Socket!")
}
```

---

### Question 351: How do you work with WebSockets in Go?

**Answer:**
Use `github.com/gorilla/websocket` (standard) or `nhooyr.io/websocket` (minimal).

**Server Example (Gorilla):**
```go
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WS
    defer conn.Close()

    for {
        // Read message
        mt, message, err := conn.ReadMessage()
        if err != nil { break }
        
        log.Printf("Received: %s", message)

        // Echo back
        conn.WriteMessage(mt, message)
    }
}
```

---

### Question 352: How do you handle bi-directional streaming in gRPC?

**Answer:**
Define `stream` in both request and response in Protobuf.

**Proto:**
```protobuf
service ChatService {
  rpc Chat(stream ChatMessage) returns (stream ChatMessage);
}
```

**Go Implementation:**
```go
func (s *server) Chat(stream pb.ChatService_ChatServer) error {
    for {
        // 1. Receive
        in, err := stream.Recv()
        if err == io.EOF { return nil }
        if err != nil { return err }

        log.Printf("Got: %s", in.Message)

        // 2. Send
        err = stream.Send(&pb.ChatMessage{Message: "Reply: " + in.Message})
        if err != nil { return err }
    }
}
```

---

### Question 353: What is Server-Sent Events (SSE) and how is it done in Go?

**Answer:**
SSE sends one-way real-time updates from Server to Client over HTTP. It's simpler than WebSockets.

**Implementation:**
1. Set Headers: `Content-Type: text/event-stream`.
2. Flush writes immediately.

```go
func sseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    flusher, ok := w.(http.Flusher)
    if !ok { return }

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case t := <-ticker.C:
            // Format: "data: <payload>\n\n"
            fmt.Fprintf(w, "data: The time is %s\n\n", t.Format(time.RFC3339))
            flusher.Flush() // Send to client immediately
        case <-r.Context().Done():
            return // Client disconnected
        }
    }
}
```

---

### Question 354: How do you manage fan-in/fan-out channel patterns?

**Answer:**

**Fan-Out:** Distribute work to multiple workers.
```go
func fanOut(ch <-chan int, workers int) {
    for i := 0; i < workers; i++ {
        go worker(ch) 
    }
}
```

**Fan-In:** Multiplex multiple channels into one.
```go
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { 
        for { c <- <-input1 } 
    }()
    go func() { 
        for { c <- <-input2 } 
    }()
    return c
}
```

**Better Fan-In (using `select` + `sync.WaitGroup` to close):**
Usually requires a "merge" function that loops over all inputs and sends to output.

---

### Question 355: How would you implement throttling on async tasks?

**Answer:**
Use a **Buffered Channel** as a semaphore (Token Bucket).

```go
// Limit to 10 concurrent requests
var semaphore = make(chan struct{}, 10)

func process(req Request) {
    semaphore <- struct{}{} // Acquire token (blocks if full)
    
    go func() {
        defer func() { <-semaphore }() // Release token
        heavyOperation(req)
    }()
}
```

**Or using `golang.org/x/time/rate`:**
```go
limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10) // 10 reqs/sec

func handler() {
    if err := limiter.Wait(ctx); err != nil {
        return
    }
    // Proceed
}
```

---

### Question 356: How do you avoid data races when consuming messages?

**Answer:**
1. **Don't share memory:** Pass copies of data.
2. **Immutable Data:** If sharing read-only data, it's safe.
3. **Loop Variable Capture:** (Common Go Pitfall pre-1.22)

**BAD:**
```go
for msg := range messages {
    go func() {
        process(msg) // Race! 'msg' changes for all goroutines
    }()
}
```

**GOOD:**
```go
for msg := range messages {
    go func(m Message) {
        process(m)
    }(msg) // Pass by value
}
```
*Note: Go 1.22 fix this loop variable issue automatically.*

---

### Question 357: How would you implement a message queue from scratch in Go?

**Answer:**
For interview/simple use, use Channels + Mutex.

```go
type SimpleQueue struct {
    mu    sync.Mutex
    items []string
    cond  *sync.Cond
}

func NewQueue() *SimpleQueue {
    q := &SimpleQueue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *SimpleQueue) Enqueue(item string) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
    q.cond.Signal() // Wake up a consumer
}

func (q *SimpleQueue) Dequeue() string {
    q.mu.Lock()
    defer q.mu.Unlock()
    
    for len(q.items) == 0 {
        q.cond.Wait() // Block until data available
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item
}
```

---

### Question 358: How do you implement ordered message processing in Go?

**Answer:**
In Kafka, ordering is only guaranteed **per partition**.

**Strategy:**
1. **Partitioning:** Ensure related messages (e.g., updates for UserID: 123) always go to the same partition using Partition Keys.
2. **Single Consumer per Partition:**
   If you use a worker pool inside a consumer, you lose ordering.
   
   **To fix worker pool ordering:**
   - Hash the content (e.g., UserID) to select a specific worker channel.
   
```go
// Dispatcher
workerChans := make([]chan Msg, 10) // 10 workers

func dispatch(msg Msg) {
    // Consistent Hashing
    workerID := hash(msg.Key) % 10
    workerChans[workerID] <- msg
}
```
Now, all messages for User 123 go to Worker 7 sequentially.

---

### Question 359: How do you handle large stream ingestion (100K+ msgs/sec)?

**Answer:**
1. **Batching:** Don't write 1 by 1. Accumulate 1000 messages or wait 500ms, then write.
2. **Workers:** Use a worker pool to parallelize deserialization/validation.
3. **Zero-Allocation:** Use `sync.Pool` to reuse objects.
4. **Asynchronous Ack:** If exact durability isn't critical (fire & forget), ack immediately.

**Batching Example:**
```go
func batchWriter(ch <-chan Item) {
    batch := make([]Item, 0, 1000)
    ticker := time.NewTicker(1 * time.Second)
    
    for {
        select {
        case item := <-ch:
            batch = append(batch, item)
            if len(batch) >= 1000 {
                flush(batch)
                batch = batch[:0]
            }
        case <-ticker.C:
            if len(batch) > 0 {
                flush(batch)
                batch = batch[:0]
            }
        }
    }
}
```

---

### Question 360: How do you persist in-flight streaming data?

**Answer:**
When crashing, any data in memory (`channels`) is lost.

**Solutions:**
1. **Write-Ahead Log (WAL):** Write to disk (append-only file) before processing.
2. **Ack Last:** Don't Acknowledge (commit offset) to Kafka until *after* DB write is confirmed.
3. **Graceful Shutdown:**
   - Stop accepting new messages.
   - Wait `Waitgroup`.Wait() for workers to finish.
   - Timeout if taking too long.

```go
func shutdown() {
    close(messages) // Stop producers
    wg.Wait()       // Wait for in-flight
    // Now safe to exit
}
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you consume messages from Kafka in Go?

**Your Response:** "I use libraries like `kafka-go` or `sarama` to consume Kafka messages. With `kafka-go`, I create a reader with broker addresses, topic, and group ID configuration. I then read messages in a loop, processing each one and handling errors appropriately. I implement proper error handling with retry logic and dead letter queues for failed messages. I also use consumer groups to distribute load across multiple instances and ensure exactly-once processing. The key is to handle message acknowledgment correctly - only committing offsets after successful processing to prevent message loss during failures."

---

**Interviewer:** How do you handle message processing failures in streaming applications?

**Your Response:** "I implement a multi-layered failure handling strategy. First, I use retry logic with exponential backoff for transient failures. If retries fail, I send messages to a dead letter queue for manual inspection. I also implement circuit breakers to prevent cascading failures when downstream services are unavailable. For critical messages, I use exactly-once semantics with transactional processing. I also monitor processing lag and error rates to detect issues early. The key is to assume that failures will happen and design the system to handle them gracefully without losing data or causing downtime."

---

**Interviewer:** What's your experience with streaming patterns in Go?

**Your Response:** "I've implemented several streaming patterns including fan-out/fan-in, pipeline processing, and event-driven architectures. For fan-out, I use multiple goroutines reading from the same channel or topic. For pipeline processing, I chain channels together with each stage doing specific processing. I also use backpressure handling to prevent overwhelming consumers. Go's channels and goroutines make these patterns natural to implement. I've built real-time data processing systems that handle thousands of messages per second using these patterns. The key is to design for scalability and handle backpressure properly."

---

**Interviewer:** How do you implement backpressure in Go streaming applications?

**Your Response:** "I implement backpressure using buffered channels with limited capacity, which naturally block producers when consumers can't keep up. I also use rate limiting with token bucket algorithms to control processing rates. For more complex scenarios, I might implement adaptive backpressure that adjusts based on system metrics. I monitor queue lengths and processing times to detect when the system is overloaded. The key is to prevent memory exhaustion while maintaining good throughput. Go's blocking channel operations make implementing backpressure straightforward compared to callback-based systems."

---

**Interviewer:** How do you handle exactly-once processing in Go?

**Your Response:** "Exactly-once processing requires coordination between the message system and processing logic. I use idempotent operations so that duplicate processing doesn't cause issues. For critical operations, I implement transactional processing where I write to the database and commit the message offset in the same transaction. I also use deduplication based on message IDs or timestamps. The key is to design the system to be resilient to duplicates since true exactly-once is difficult to achieve. I prefer at-least-once processing with idempotent operations, which gives me the same effect with simpler implementation."

---

**Interviewer:** How do you optimize streaming performance in Go?

**Your Response:** "I optimize streaming performance by minimizing allocations, using object pools for frequently created structs, and batching operations where possible. I also tune channel buffer sizes based on throughput requirements and use worker pools to control concurrency. For high-throughput systems, I might use lock-free data structures and avoid unnecessary serialization. I profile with pprof to identify bottlenecks and optimize hot paths. The key is to balance throughput with latency and resource usage. I've achieved significant performance gains by reducing allocations and using more efficient data structures."

---

**Interviewer:** How do you handle graceful shutdown in streaming applications?

**Your Response:** "I implement graceful shutdown by listening for shutdown signals and then stopping new message intake while allowing in-flight processing to complete. I use WaitGroups to track active goroutines and ensure they finish before exiting. I also implement timeouts to prevent hanging during shutdown. For critical systems, I might drain all remaining messages before shutting down. The key is to ensure no messages are lost during the shutdown process. I also persist any in-flight state to disk so it can be recovered on restart. This approach ensures zero-downtime deployments and data consistency."

---

**Interviewer:** What's your experience with different message brokers in Go?

**Your Response:** "I've worked with Kafka, RabbitMQ, and NATS in Go applications. Kafka is great for high-throughput, durable event streaming with its partitioned log architecture. RabbitMQ excels at flexible routing patterns and reliable message delivery. NATS is lightweight and fast for simple pub-sub patterns. I choose based on requirements - Kafka for event sourcing and analytics, RabbitMQ for complex routing, and NATS for simple, high-performance messaging. Go has excellent client libraries for all of them. The key is to match the broker capabilities to the use case rather than using one size fits all."

---

**Interviewer:** How do you monitor streaming applications in Go?

**Your Response:** "I monitor streaming applications using Prometheus metrics for throughput, latency, error rates, and queue depths. I also track consumer lag to ensure we're not falling behind on message processing. I implement structured logging with correlation IDs to trace messages through the system. For distributed tracing, I use OpenTelemetry to track message flow across services. I also set up alerts for critical metrics like high error rates or consumer lag. The key is to have visibility into the entire message pipeline so I can quickly identify and resolve issues."

---

**Interviewer:** How do you handle schema evolution in streaming systems?

**Your Response:** "I handle schema evolution using forward and backward compatible changes. I use schema registries like Confluent Schema Registry for Avro schemas, which manages compatibility checks. I design my schemas to be tolerant of new fields and handle missing fields gracefully. For JSON messages, I use optional fields and sensible defaults. I also version my APIs and maintain multiple versions during transitions. The key is to plan for change from the beginning and design schemas that can evolve without breaking existing consumers. I've learned that being too strict with schemas leads to brittle systems."
