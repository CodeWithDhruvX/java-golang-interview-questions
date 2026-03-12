## 🧵 Concurrency Architecture & Design Patterns (Questions 961-980)

### Question 961: How do you architect a pub/sub system in Go?

**Answer:**
(See Q606). Channels + Map.

### Explanation
Pub/sub system architecture in Go uses channels for message passing and maps for topic management. Publishers send messages to topic channels, and subscribers receive messages from their subscribed topic channels.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you architect a pub/sub system in Go?
**Your Response:** "I architect pub/sub systems using channels and maps. I use a map to manage topics - each topic has its own channel. Publishers send messages to the appropriate topic channel, and subscribers receive from channels they're interested in. I manage subscriptions using another map that tracks which subscribers are listening to which topics. The key is using Go's channels for the actual message passing - they provide built-in buffering and synchronization. I handle subscriber management, message filtering, and cleanup when subscribers disconnect. This approach is lightweight and doesn't require external dependencies. For distributed systems, I might use Redis or NATS, but for in-process pub/sub, channels are perfect. The design is event-driven and scales well with Go's concurrency model."

---

### Question 962: How do you build a pipeline using goroutines?

**Answer:**
Stages: `Gen() -> Sq() -> Print()`.
Each function takes `<-chan In` and returns `<-chan Out`.
Connect them in main: `out = Sq(Gen())`.

### Explanation
Pipeline building with goroutines uses stages where each function takes an input channel and returns an output channel. Stages are connected by composing function calls like out = Sq(Gen()), creating a data processing pipeline.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a pipeline using goroutines?
**Your Response:** "I build pipelines using goroutines where each stage is a function that takes an input channel and returns an output channel. For example, `Gen()` generates data on a channel, `Sq()` squares the numbers, and `Print()` outputs them. Each stage runs in its own goroutine, processing data as it arrives. I connect them by composing the functions: `out = Sq(Gen())`. This creates a streaming pipeline where data flows through stages without buffering everything in memory. The key is each stage is independent and communicates only through channels. I can add more stages, parallelize work, or backpressure naturally through channel blocking. This pattern is perfect for data processing workflows and follows Go's concurrency philosophy."

---

### Question 963: What is the fan-in/fan-out pattern in Go?

**Answer:**
(See Q601/Q602).

### Explanation
Fan-in/fan-out pattern in Go distributes work across multiple goroutines (fan-out) and then combines results back into a single channel (fan-in). This pattern maximizes concurrency while maintaining coordinated result collection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the fan-in/fan-out pattern in Go?
**Your Response:** "Fan-in and fan-out are concurrency patterns for distributing and collecting work. Fan-out is when I distribute work across multiple goroutines - I send tasks to multiple worker goroutines that process them in parallel. Fan-in is when I collect results from multiple goroutines back into a single channel. The combination allows me to maximize concurrency while still coordinating results. For example, I might fan-out URL fetches to 10 goroutines, then fan-in their responses into a single results channel. The key is using channels to coordinate the work distribution and collection. This pattern scales processing by leveraging multiple CPU cores while keeping the code manageable. It's especially useful for CPU-bound or I/O-bound operations that can run in parallel."

---

### Question 964: How do you limit concurrency using semaphores?

**Answer:**
(See Q604). Buffered channel.

### Explanation
Concurrency limiting with semaphores uses buffered channels as counting semaphores. The channel capacity represents the maximum number of concurrent operations, with goroutines acquiring and releasing permits through channel operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you limit concurrency using semaphores?
**Your Response:** "I limit concurrency using buffered channels as semaphores. I create a buffered channel with capacity equal to my concurrency limit. Before starting work, a goroutine sends to the channel to acquire a permit. When done, it receives from the channel to release the permit. If the channel is full, goroutines block until a permit is available. This pattern prevents too many goroutines from running simultaneously, which is crucial for managing resources like database connections or memory. The key is the channel acts as a counting semaphore - its capacity represents the maximum concurrent operations. This approach is simple, efficient, and uses Go's built-in channel semantics. I can adjust the limit based on resource availability or system load."

---

### Question 965: How do you implement a worker pool?

**Answer:**
(See Q609).
Crucial for capping resource usage (DB connections, memory).

### Explanation
Worker pool implementation uses a fixed number of worker goroutines that process tasks from a shared queue. This caps resource usage by limiting concurrent database connections and memory consumption while providing controlled task processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a worker pool?
**Your Response:** "I implement worker pools with a fixed number of worker goroutines that process tasks from a shared queue. I create a channel for jobs and start a specific number of worker goroutines that continuously read from this channel. This caps resource usage - I can limit database connections or memory usage by controlling the number of workers. Workers process tasks independently but share the same job queue. The key is finding the right balance between throughput and resource usage. Too few workers and I'm not using resources efficiently; too many and I might overwhelm downstream systems. I can also implement graceful shutdown and task retry logic. Worker pools are perfect for processing queues, handling web requests, or any concurrent task processing where I need to control resource usage."

---

### Question 966: How do you handle retries with backoff in goroutines?

**Answer:**
(See Q617).
Combine `time.Sleep` with `select { case <-ctx.Done(): return }` to ensure retry loop stops if request is cancelled.

### Explanation
Retry with backoff in goroutines combines time.Sleep for delays with select statements that check context cancellation. This ensures retry loops stop gracefully when requests are cancelled while implementing exponential backoff strategies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle retries with backoff in goroutines?
**Your Response:** "I handle retries with backoff by combining `time.Sleep` with context cancellation checks. I use a select statement that either waits for the backoff delay or for context cancellation: `select { case <-time.After(delay): case <-ctx.Done(): return }`. This ensures my retry loop stops immediately if the request is cancelled, rather than waiting for the full delay. I implement exponential backoff by increasing the delay with each attempt. The key is respecting context cancellation while retrying operations. I also cap the maximum number of retries and maximum delay time. This pattern is essential for building resilient systems that can handle temporary failures without hanging indefinitely."

---

### Question 967: What is the circuit breaker pattern in Go?

**Answer:**
(See Q619). `sony/gobreaker`.

### Explanation
Circuit breaker pattern in Go uses sony/gobreaker library to prevent cascading failures. It monitors service failures and temporarily stops sending requests to failing services, allowing them time to recover before trying again.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the circuit breaker pattern in Go?
**Your Response:** "I implement circuit breakers using the `sony/gobreaker` library. The circuit breaker monitors service failures and trips open when failures exceed a threshold. When tripped, it immediately returns errors without making actual requests, giving the failing service time to recover. After a cooldown period, it enters a half-open state and tries a few requests. If successful, it closes again; if not, it stays open. This prevents cascading failures and system overload. The key is protecting my system from downstream failures while allowing automatic recovery. I configure thresholds based on service characteristics and monitor circuit breaker state. This pattern is essential for building resilient distributed systems."

---

### Question 968: How do you implement message deduplication?

**Answer:**
(See Q625). SHA256 of message content -> Check in Redis.

### Explanation
Message deduplication uses SHA256 hashing of message content to generate unique identifiers. These hashes are stored in Redis to detect and prevent processing of duplicate messages across distributed systems.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement message deduplication?
**Your Response:** "I implement message deduplication by computing a SHA256 hash of the message content and storing it in Redis. Before processing a message, I check if its hash already exists in Redis. If it does, I skip processing; if not, I process it and store the hash. I set an expiration on the Redis entries so old hashes eventually expire to save space. This approach works even across multiple instances since they all share Redis. The key is using a content-based hash rather than message IDs, which handles cases where the same message might arrive with different IDs. I also handle hash collisions and implement proper error handling. This pattern is essential for ensuring exactly-once processing in distributed systems."

---

### Question 969: How do you synchronize shared state across goroutines?

**Answer:**
1.  **Start with Channels** (Share memory by communicating).
2.  If complex/performance critical: **Mutex**.
3.  If simple counter: **Atomic**.

### Explanation
Shared state synchronization in Go follows a hierarchy: start with channels for sharing memory by communicating, use mutex for complex or performance-critical scenarios, and use atomic operations for simple counters to avoid mutex overhead.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you synchronize shared state across goroutines?
**Your Response:** "I synchronize shared state following Go's philosophy: start with channels to communicate rather than share memory. If I have complex shared data or performance-critical sections, I use mutex to protect access. For simple counters or flags, I use atomic operations which are faster than mutex. The key is choosing the right synchronization primitive for the use case. Channels are great for coordination and passing data between goroutines. Mutex is better when multiple goroutines need to access complex shared data. Atomics are perfect for simple numeric values. I avoid global variables and design my concurrency around message passing when possible. The goal is to write concurrent code that's both correct and efficient."

---

### Question 970: How do you detect livelocks in Go?

**Answer:**
Harder than Deadlocks.
Goroutines are running (burning CPU) but making no progress (e.g., constantly failing a CAS operation or retrying immediately).
Monitoring CPU usage + No Application Throughput is a sign.

### Explanation
Livelock detection in Go is harder than deadlock detection. In livelocks, goroutines are actively running and consuming CPU but making no progress, such as constantly failing CAS operations or retrying immediately. High CPU usage with no application throughput indicates a livelock.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you detect livelocks in Go?
**Your Response:** "Livelocks are harder to detect than deadlocks because the goroutines are still running and consuming CPU, but they're not making progress. I look for signs like high CPU usage but no application throughput. Common causes include goroutines constantly failing compare-and-swap operations or retrying operations immediately without backoff. I monitor system metrics and application throughput to spot these patterns. I also add logging to detect when operations are repeatedly failing. The key is understanding that in a livelock, the system is active but stuck in a loop of useless activity. I prevent livelocks by implementing proper backoff strategies and ensuring progress conditions in my concurrent algorithms."

---

### Question 971: How do you timeout long-running operations?

**Answer:**
`select`.
```go
select {
case res := <-work():
    return res
case <-time.After(5 * time.Second):
    return Timeout
}
```

### Explanation
Long-running operation timeouts use select statements with time.After. The select waits on either the operation result or a timeout channel, returning the result if available or Timeout if the time limit expires first.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you timeout long-running operations?
**Your Response:** "I timeout long-running operations using select with a time.After channel. I create a select that waits on either the operation result or a timeout. If the operation completes first, I return the result. If the timeout fires first, I return a timeout error. This pattern works for any operation that can be expressed as a channel receive. I can also use context.WithTimeout for more complex scenarios. The key is not blocking indefinitely - always having a way to cancel or timeout operations. I handle cleanup properly when timeouts occur to avoid resource leaks. This pattern is essential for building responsive systems that don't hang on slow operations."

---

### Question 972: How do you use the actor model in Go?

**Answer:**
Go doesn't have native Actors (like Erlang/Akka).
Simulate it:
Struct with a `Run()` loop and an `inbox chan Msg`.
Methods on struct just send to `inbox`.

### Explanation
Actor model in Go is simulated since Go lacks native actors. A struct with a Run() method processes messages from an inbox channel. Public methods just send messages to the inbox, encapsulating all state within the actor goroutine.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use the actor model in Go?
**Your Response:** "Go doesn't have native actors like Erlang or Akka, but I can simulate the actor model. I create a struct with an inbox channel for messages and a Run() method that runs in a goroutine, continuously processing messages from the inbox. The public methods on the struct just send messages to the inbox rather than directly manipulating state. This encapsulates all state within the actor goroutine, ensuring serialized access. The actor processes messages one at a time, avoiding race conditions. I can have multiple actors communicating through message passing. The key is treating each actor as an independent concurrent entity with its own state and message queue. This pattern works well for building concurrent systems with clear boundaries and message-based communication."

---

### Question 973: How do you architect loosely coupled goroutines?

**Answer:**
Do not share state variables.
Pass data via Channels.
Use `Context` for cancellation signal propagation.

### Explanation
Loosely coupled goroutines avoid sharing state variables, pass data through channels for communication, and use context for cancellation signal propagation. This design minimizes coupling and makes goroutines independently testable and maintainable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you architect loosely coupled goroutines?
**Your Response:** "I architect loosely coupled goroutines by avoiding shared state variables and using channels for all data communication. Each goroutine is self-contained and communicates only through well-defined channel interfaces. I use context for cancellation signals so goroutines can shut down cleanly when requested. This design minimizes coupling - goroutines don't need to know about each other's internal state. I can test each goroutine independently by sending test messages through its channels. The key is designing clear communication protocols and avoiding shared memory. This approach makes the system more maintainable, testable, and easier to reason about. It also follows Go's philosophy of communicating by sharing rather than sharing memory."

---

### Question 974: How do you design state machines in Go?

**Answer:**
`type State func() State`
Loop: `currentState = currentState()`.
Each state function returns the next state function. (Rob Pike's Lexer pattern).

### Explanation
State machine design in Go uses type State func() State where each state function returns the next state function. The loop continuously executes currentState = currentState(), implementing Rob Pike's lexer pattern for state transitions.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design state machines in Go?
**Your Response:** "I design state machines using Rob Pike's pattern where each state is a function that returns the next state function. I define `type State func() State` and implement each state as a function that decides what the next state should be. The main loop continuously calls `currentState = currentState()`. This approach is elegant because each state contains its own logic and determines the transition to the next state. I don't need complex switch statements or enums - the state functions themselves encode the behavior. The pattern is extensible - I can add new states without modifying existing code. It's also testable since each state function can be tested independently. This works great for parsers, protocol handlers, or any system with discrete states."

---

### Question 975: How do you throttle a job queue in Go?

**Answer:**
(See Q608). Token Bucket.

### Explanation
Job queue throttling uses token bucket algorithm where tokens are added at a fixed rate and jobs must acquire tokens before processing. This controls the rate of job processing while allowing bursts up to the bucket capacity.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you throttle a job queue in Go?
**Your Response:** "I throttle job queues using the token bucket algorithm. I create tokens at a fixed rate and jobs must acquire a token before processing. This allows bursts of work up to the bucket capacity but maintains a steady average rate. I implement this with channels - adding tokens periodically and requiring jobs to receive a token before starting. The key is balancing throughput with resource constraints - I can handle bursts while preventing overload. I adjust the token generation rate and bucket size based on system capacity and requirements. This approach is more flexible than simple rate limiting because it allows temporary bursts while maintaining long-term rate control. It's perfect for processing queues, API calls, or any workload that needs rate control."

---

### Question 976: How do you monitor goroutine health?

**Answer:**
Heartbeat channel.
Worker sends `heartbeat <- struct{}{}` every second.
Supervisor kills/restarts worker if no beat for 5s.

### Explanation
Goroutine health monitoring uses heartbeat channels where workers send periodic heartbeat signals. A supervisor monitors these signals and kills/restarts workers that miss heartbeats for a specified timeout period.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you monitor goroutine health?
**Your Response:** "I monitor goroutine health using heartbeat channels. Each worker goroutine sends a heartbeat signal every second through a channel. A supervisor goroutine monitors these heartbeats and if it doesn't receive a beat for 5 seconds, it assumes the worker is dead or stuck and kills/restarts it. I implement this with a ticker for regular heartbeats and a select statement in the supervisor that times out if no heartbeat arrives. The key is detecting when goroutines are stuck in infinite loops or blocked operations. I also track metrics like number of restarts and heartbeat intervals. This pattern ensures system resilience by automatically recovering from failed or stuck goroutines."

---

### Question 977: How do you track context propagation in goroutines?

**Answer:**
Always pass `ctx` as first argument.
Use `pprof` labels `pprof.Do(ctx, labels, func(ctx) { ... })` to attach metadata to goroutines in profiles.

### Explanation
Context propagation in goroutines always passes ctx as the first argument. pprof labels attach metadata to goroutines in profiles using pproff.Do(ctx, labels, func), enabling better tracing and profiling of concurrent operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you track context propagation in goroutines?
**Your Response:** "I track context propagation by always passing `ctx` as the first argument to functions. For better tracing, I use `pprof.Do` with labels to attach metadata to goroutines in profiles. This helps me track which goroutine belongs to which request or operation. I include request IDs, user IDs, or operation types in the context labels. When I profile the application, I can see exactly what each goroutine is doing. The key is maintaining context throughout the call chain and enriching it with relevant metadata. I also use context for cancellation signals and timeouts. This approach gives me visibility into concurrent operations and helps with debugging and performance optimization."

---

### Question 978: How do you implement saga pattern in Go services?

**Answer:**
Distributed Transactions.
Sequence of local transactions.
If Step 3 fails, execute Compensating Transactions (Undo) for Step 2 and Step 1.
Orchestrator based (Central Go code) or Choreography (Events).

### Explanation
Saga pattern in Go services manages distributed transactions through a sequence of local transactions. If any step fails, compensating transactions undo previous steps. Implementation can be orchestrator-based with central Go code or choreography-based with events.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement saga pattern in Go services?
**Your Response:** "I implement the saga pattern for distributed transactions by breaking operations into a sequence of local transactions. If any step fails, I execute compensating transactions to undo the previous steps. For example, in an order processing saga, if payment fails after reserving inventory, I release the inventory. I can implement this either with an orchestrator approach - central Go code that coordinates all steps - or with choreography where services communicate through events. The orchestrator gives me centralized control, while choreography is more decentralized. The key is having reliable compensating actions and handling failures gracefully. I also persist saga state to recover from crashes. This pattern ensures data consistency across multiple services without using distributed transactions."

---

### Question 979: How do you chain async jobs with error handling?

**Answer:**
Use `errgroup`.
It runs N jobs. If any returns error, it cancels the context for all others and returns the first error.

### Explanation
Chaining async jobs with error handling uses errgroup which runs multiple jobs concurrently. If any job returns an error, errgroup cancels the context for all other jobs and returns the first error encountered.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you chain async jobs with error handling?
**Your Response:** "I chain async jobs with error handling using `errgroup`. I create an errgroup with a context and launch multiple goroutines to run jobs concurrently. If any job returns an error, the errgroup cancels the context for all other jobs and returns the first error. This ensures that failures propagate quickly and resources aren't wasted on doomed operations. I can also use `errgroup.SetLimit` to control concurrency. The key is that errgroup handles both the concurrency coordination and error propagation for me. I wait for all jobs to complete with `g.Wait()` and handle any errors appropriately. This pattern is perfect for running multiple independent operations where I need to fail fast if any operation fails."

---

### Question 980: How do you log and trace concurrent tasks?

**Answer:**
Include `GoroutineID` (hacky) or better, a `RequestID` in the closure/logger context.

### Explanation
Logging and tracing concurrent tasks includes GoroutineID (though hacky) or preferably a RequestID in the closure/logger context. This enables tracking which goroutine or request generated each log entry for debugging.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you log and trace concurrent tasks?
**Your Response:** "I log and trace concurrent tasks by including context information in my logs. Rather than using GoroutineID which is hacky, I prefer passing a RequestID or correlation ID through the closure or logger context. When a request starts, I generate a unique ID and pass it through all goroutines involved in processing that request. Each log entry includes this ID, making it easy to trace the entire flow. I also include operation names, timestamps, and relevant metadata. For distributed tracing, I might use OpenTelemetry. The key is having enough context to reconstruct the execution flow across multiple goroutines. This approach makes debugging concurrent systems much easier by providing a clear trail of which operations belong together."

---
