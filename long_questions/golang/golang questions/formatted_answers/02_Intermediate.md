# Intermediate Level Golang Interview Questions

## From 04 Concurrency

# 🟣 **61–80: Concurrency and Goroutines**

### 62. What are goroutines?
"Goroutines are **lightweight threads** managed by the Go runtime, not the Operating System. They are the fundamental unit of concurrency in Go.

Technically, they start with a tiny 2KB stack that grows and shrinks dynamically, whereas an OS thread might take 1MB. This efficient memory usage means I can spin up tens of thousands of goroutines on a single machine without crashing it.

I treat them as 'fire-and-forget' functions. When I use the `go` keyword, the function runs independently. However, I always ensure I have a plan for how they will finish or communicate back, otherwise, I risk leaks."

#### Indepth
Goroutines are **multiplexed** onto OS threads (M:N scheduling). The Go Scheduler uses a "work-stealing" algorithm. If a processor runs out of work, it steals goroutines from another processor's queue. This ensures high CPU utilization without manual thread management.

---

### 63. What do you start a goroutine?
"I simply use the keyword `go` followed by a function invocation. It’s syntactically the easiest concurrent model I’ve ever used.

For example, `go processBytes(data)`. This immediately returns control to the main function, scheduling `processBytes` to run on the Go runtime's thread pool.

One catch is that the `main` function doesn't wait for goroutines. If `main` returns, the program exits and kills all running goroutines. So, I always use a **WaitGroup** or a **Done Channel** to ensure important background work completes before the program shuts down."

#### Indepth
Under the hood, `go func()` allocates a structure `g` on the heap and adds it to the local run queue of the current Logic Processor (`P`). It does minimal setup, usually taking nanoseconds, compared to milliseconds for an OS thread spawn.

---

### 64. What is a channel in Go?
"A **channel** is a typed conduit that allows goroutines to communicate and synchronize execution. It’s the embodiment of Go's philosophy: 'Share memory by communicating.'

Under the hood, it’s a thread-safe queue with locking built-in. When I write `ch <- value`, the value is copied into the channel. When I read `val := <-ch`, it’s copied out.

I use them not just for data, but for signaling. A receive operation blocks until data is available, which makes them perfect for coordinating workflow steps without messy implementations of mutexes or condition variables."

#### Indepth
Channels use an internal `hchan` struct protected by a `mutex`. Copying data into/out of the channel involves memory copy (`memmove`). For very large structs, passing pointers over channels is faster, but value passing is safer to prevent race conditions (shared ownership).

---

### 65. What is the difference between buffered and unbuffered channels?
"It comes down to **blocking behavior**. An **unbuffered channel** has no capacity; certain synchronicity is enforced because the sender *must* wait for a receiver to be ready, and vice versa.

A **buffered channel** has a queue size (e.g., `make(chan int, 10)`). The sender only blocks if the buffer is **full**. The receiver only blocks if the buffer is **empty**.

I generally prefer unbuffered channels for strict synchronization (like passing a 'result' or 'done' signal). I use buffered channels when I need to decouple the producer from the consumer—handling 'bursty' traffic where the producer might be temporarily faster than the worker."

#### Indepth
A buffered channel of size 1 can act as a **Mutex**. `ch <- 1` locks, `<-ch` unlocks. However, `sync.Mutex` is optimized for this specific case (using atomic CPU instructions) and is generally faster and clearer for simple critical sections.

---

### 66. How do you close a channel?
"I use the built-in `close(ch)` function. It’s a signal to all receivers that no more values will ever be sent on this channel.

Technically, this sets a flag on the channel structure. Receivers observing a closed channel will immediately return a zero-value and `false` (if checking the comma-ok idiom), rather than blocking.

Crucially, I **only** close channels from the sender side. Closing from the receiver side or closing a channel twice causes a runtime panic. If I have multiple senders, I typically use a separate synchronization mechanism (like a `WaitGroup`) to decide when to close it."

#### Indepth
Closing a channel isn't about freeing memory (GC does that). It's strictly about control flow signals. A common pattern for "N senders, 1 receiver" is to have the receiver close a separate struct `done` channel that the senders listen to, telling them to stop sending.

---

### 67. What happens when you send to a closed channel?
"It **panics** immediately. This is a fatal runtime error.

This strictness ensures integrity: you shouldn't be sending data to a stream that has declared 'I am finished.'

To avoid this, I design my systems so that the 'owner' of the channel (the one writing to it) is the only one responsible for closing it. If `send` might occur after close, strictly wrapping the send in a select with a 'done' context check is a common pattern I use."

#### Indepth
This panic is guaranteed to happen to prevent hard-to-debug data corruption. Unlike a nil pointer (which is a crash due to invalid memory), this is a logic enforcement. There is no `safeSend` function in Go; you must know the state of your system.

---

### 68. How to detect a closed channel while receiving?
"I use the **comma-ok idiom**: `value, ok := <-ch`.

If the channel is open and has data, `ok` is `true`. If the channel is closed and empty, `value` will be the zero-value (like 0 or "") and `ok` will be `false`.

I rely on this heavily inside `for range` loops. A `for msg := range ch` loop automatically breaks when the channel is closed, which is the cleanest way to drain a worker queue until shutdown."

#### Indepth
Note that a closed channel returns the *zero value* forever. It doesn't block. This is a common bug source: if you have a `select` loop and one channel closes, that case becomes non-blocking and generates an infinite loop of zero values, spiking CPU to 100%. Always set the channel variable to `nil` after seeing it close to disable that case in the select.

---

### 69. What is the `select` statement in Go?
"The `select` statement is like a `switch`, but exclusively for channel operations. It allows a goroutine to wait on multiple communication operations simultaneously.

It blocks until one of its cases can proceed. If multiple cases are ready (e.g., data arrived on two different channels), `select` picks one **at random**. This randomness prevents one active channel from starvation.

I use it in almost every long-running backend service to handle **cancellation** alongside business logic—waiting for `<-ctx.Done()` in one case and `<-jobQueue` in another."

#### Indepth
The Go runtime randomizes the polling order of cases in a `select` statement to ensure fairness. If it was deterministic (always checking top-to-bottom), the top case could starve the bottom ones if it always had high traffic.

---

### 70. How do you implement timeouts with `select`?
"I add a `case` that receives from `time.After(duration)`.

`time.After` returns a channel that sends the current time after the delay. So my structure looks like:
`select { case res := <-resultCh: handle(res); case <-time.After(2 * time.Second): return ErrTimeout }`.

Usage-wise, this is critical for production systems. I never want to block indefinitely on a network call or an external service. Hard timeouts ensure my system stays responsive even if dependencies hang."

#### Indepth
Beware that `time.After` leaks resources if the select picks another case, until the timer fires. In high-throughput loops, this accumulates millions of timers. Use `time.NewTimer` and explicitly `Stop()` it to prevent memory leaks in tight loops.

---

### 71. What is a `sync.WaitGroup`?
"It’s a synchronization primitive that waits for a collection of goroutines to finish. Think of it as a thread-safe counter.

I call `wg.Add(1)` before starting a task. The worker calls `wg.Done()` (decrement) when finished. The main thread calls `wg.Wait()`, blocking until the counter hits zero.

It’s my go-to tool for **fan-out/fan-in** patterns—like spawning 10 scrapers to fetch URLs in parallel and waiting for all of them to finish before aggregating the results."

#### Indepth
`WaitGroup` must be passed by pointer (`*sync.WaitGroup`), never by value! Copying a WaitGroup copies its internal counter state, leading to a deadlock where the worker signals the *copy*, but the main thread waits on the *original*. `go vet` catches this.

---

### 72. How does `sync.Mutex` work?
"A `Mutex` (Mutual Exclusion) creates a critical section where only *one* goroutine can execute at a time. It prevents race conditions on shared memory.

I use `mu.Lock()` to claim exclusive access and `mu.Unlock()` to release it—almost always in a `defer` statement right after locking to ensure I don't forget it if a panic occurs.

While channels are great for data flow, I prefer `Mutex` for **state**. If I just need to safely increment a counter or update a map in a struct, a Mutex is often faster and simpler than setting up a dedicated channel-coordinating goroutine."

#### Indepth
`sync.Mutex` has two modes: **Normal** and **Starvation**. In Starvation mode (triggered if a waiter waits >1ms), the mutex ownership is handed directly to the first waiter, bypassing new arriving goroutines. This prevents "tail latency" spikes where a request waits forever.

---

### 73. What is `sync.Once`?
"It’s a utility that ensures a piece of code is executed **exactly once**, regardless of how many goroutines call it simultaneously.

Internally, it uses an atomic counter and a mutex. The first caller wins the race and executes the function; subsequent callers wait until it finishes, then skip execution.

I use this for **lazy singleton initialization**—like connecting to a database or loading a heavy configuration file only when the first request arrives, rather than at application startup."

#### Indepth
`sync.Once` is implemented using an atomic variable `done`. The "fast path" just checks the atomic integer (cheap). Only if it's 0 does it acquire the slow mutex to execute the function. This makes `Do()` extremely low-overhead for frequent calls.

---

### 74. How do you avoid race conditions?
"The primary mantra is: **'Do not communicate by sharing memory; instead, share memory by communicating.'**

This means I prefer passing data copies over **channels** rather than multiple threads modifying the same pointer. If I *must* share memory (like a global cache), I protect it strictly with `sync.Mutex` or `sync.RWMutex`.

To verify my safety, I always run my tests with the `-race` flag (`go test -race`). The Go race detector is incredibly good at finding unlocked concurrent access that I missed during code review."

#### Indepth
The `-race` flag instruments code with "happens-before" annotations. It tracks memory access at runtime. It increases memory usage by 5-10x and execution time by 2-20x, so don't run it in production, but *always* run it in CI/CD pipeline.

---

### 75. What is the Go memory model?
"It’s the specification that defines **visibility**—guaranteeing when a variable written by one goroutine is visible to another.

The core rule is: without explicit synchronization (channels, mutexes, waiter groups), there is **no guarantee** that Goroutine B sees the write from Goroutine A. Compilers reorder code, and CPU caches delay writes.

This implies that 'clever' lock-free code using plain boolean flags is usually broken. I stick to the standard synchronization primitives which act as memory barriers, flushing caches and enforcing order."

#### Indepth
Go's memory model is weaker than Java's (volatile) but stronger than C++'s relaxed atomics. It guarantees that a send on a channel *happens before* the receive completes. This causality chain is what allows generic synchronization without understanding CPU cache line invalidation.

---

### 76. How do you use `context.Context` for cancellation?
"The `Context` is the standard way to carry cancellation signals across API boundaries and goroutines.

I create a context with `ctx, cancel := context.WithCancel(parent)`. When I call `cancel()`, the `ctx.Done()` channel closes. All downstream functions listening to this channel stop their work and return immediately.

I use this for every request path. If a user closes their browser tab, I cancel the context, which stops the DB query and saves resources. It’s a required pattern for any robust Go server."

#### Indepth
`context` values are immutable. `WithCancel` returns a *new* context child. This forms a tree. Cancelling the parent cancels all children. Use `context.Background()` for the root and `context.TODO()` when you're unsure (refactoring placeholders).

---

### 77. How to pass data between goroutines?
"I primarily use **channels**. They are the safest and most idiomatic way to transfer ownership of data.

If I have a Producer generating 'Jobs' and a Consumer processing them, a buffered channel acts as the perfect glue. It handles the locking and queuing logic automatically.

For simple configuration data or read-only 'context', I might pass immutable structs or use `context.Context` values, but for the actual flow of business data, channels are my default choice."

#### Indepth
Don't use channels for everything. If you have "referentially transparent" data (just values), passing them as function arguments is faster (register passing vs channel locking). Channels are for specific *coordination* or *handoff* points.

---

### 78. What is the `runtime.GOMAXPROCS()` function?
"It controls the number of **OS threads** that can execute Go code simultaneously. By default, it equals the number of CPU cores available.

Changing it limits how much actual parallelism the runtime utilizes. If I set it to 1, my program becomes strictly concurrent but single-threaded (no true parallel execution).

I rarely touch this in code. However, in **Kubernetes**, I verify it matches the CPU quota (using `automaxprocs`). If my container has 2 CPUs but the Node has 64, Go might spin up 64 threads, causing excessive context switching and throttling. Matching it to the quota fixes this."

#### Indepth
Prior to Go 1.5, GOMAXPROCS defaulted to 1. This historic artifact is why some old tutorials say "Go isn't parallel by default". Today, it defaults to `runtime.NumCPU()`. Manipulating this is rarely needed unless you are running in a containerized environment with "fractional CPUS" (like 0.5 CPU).

---

### 79. How do you detect deadlocks in Go?
"A deadlock happens when all goroutines are waiting for each other, and no one can proceed. Go has a built-in detector that panics with `fatal error: all goroutines are asleep - deadlock!` if the runtime sees zero executable goroutines.

However, this only catches complete global deadlocks. **Partial deadlocks** (where 2 threads are stuck but others run) are harder.

I detect those by using `pprof` to inspect the stack traces of stuck goroutines or by using timeouts on all channel/lock operations. If a lock takes >5 seconds, I panic or log a stack trace to debug why it was never released."

#### Indepth
You can retrieve the stack trace of all goroutines programmatically using `runtime.Stack(buf, true)`. Sending `SIGQUIT` (Ctrl+\) to a running Go program performs a "core dump" to stderr, which is invaluable for debugging stuck processes in production.

---

### 80. What are worker pools and how do you implement them?
"A worker pool is a pattern to limit concurrency. Instead of spawning `go func()` for *every* incoming request (which could exhaust RAM), I start a fixed number of workers (e.g., 5).

I create a `jobs` channel. I start 5 goroutines that `range` over the `jobs` channel. The main handler simply sends requests into the channel.

This ensures my database is never hit by more than 5 concurrent queries, acting as a natural backpressure mechanism. If the buffer fills, the callers wait, rather than the system crashing."

#### Indepth
A more advanced pattern is the **Semaphore** pattern using a buffered channel. `sem := make(chan struct{}, 5)`. Before starting a job, `sem <- struct{}{}`. When done, `<-sem`. This limits active goroutines without pre-allocating a fixed pool of workers.

---

### 81. How to write concurrent-safe data structures?
"I wrap the data structure (like a map) in a struct that has a `sync.RWMutex`.

I ensure that every read acquires a **Read Lock** (`RLock`) and every write acquires a **Write Lock** (`Lock`). Alternatively, I can use `sync.Map` for specific use cases like append-only caches, but usually, a standard map with a mutex is clearer.

The key is strict discipline: **never** expose the underlying map directly. Only access it through methods that handle the locking."

#### Indepth
`sync.Map` is specialized for cases where keys are stable (write once, read many) or disjoint (different goroutines access different keys). For general "R/W cache" use cases, a `map` + `RWMutex` is often 2x faster due to lower interface overhead (no type assertions).


## From 08 Networking WebDev

# 🔵 **141–160: Networking, APIs, and Web Dev**

### 142. How to build a REST API in Go?
"I typically start with the standard `net/http` package.

I define handlers using `http.HandleFunc("/users", handler)`.
Inside the handler, I decode the JSON body, interact with my service layer, and encode the response to `w`.
For more complex routing (like `/users/{id}`), I'll upgrade to **Chi** or **Gin**. Chi is my favorite because it feels like a lightweight extension of the standard library, whereas Gin is a full framework."

#### Indepth
The standard library's `http.ServeMux` got a huge upgrade in Go 1.22. It now supports method-based routing (`POST /items`) and wildcards (`/items/{id}`). This negates the need for Chi or Gorilla Mux for 90% of use cases, making the "stdlib-only" approach even more viable.

---

### 143. How to parse JSON and XML in Go?
"I use `encoding/json`.

I define a struct with tags: `type User struct { Name string \`json:"name"\` }`.
Then I use `json.Unmarshal(data, &user)` to parse it.
For XML, it's identical but with `encoding/xml` and `xml:"..."` tags.
Since `encoding/json` uses reflection, for extremely high-throughput systems, I might swap it for **easyjson** or **fastjson** to generate static parsing code."

#### Indepth
`encoding/json` respects struct tags like `json:"-"` (ignore field) and `json:",omitempty"` (omit if zero-value). A common pitfall is handling `time.Time`: JSON has no date standard, so Go uses RFC3339 strings by default. You can override this by implementing `MarshalJSON` on a custom wrapper type.

---

### 144. What is the use of `http.Handler` and `http.HandlerFunc`?
"`http.Handler` is an interface with a singe method: `ServeHTTP(w, r)`.
Anything that implements this can process web requests.

`http.HandlerFunc` is a convenience adapter. It lets me take a simple function `func(w, r)` and cast it to `http.HandlerFunc(myFunc)`, which now satisfies the interface. It saves me from creating a new struct type for every single route."

#### Indepth
This adapter pattern is everywhere in Go. `http.HandlerFunc(myFunc)` works because Go allows methods on types derived from functions. It essentially says "when `ServeHTTP` is called on this function value, just execute the function itself".

---

### 145. How do you implement middleware manually in Go?
"Middleware is just a function that takes an `http.Handler` and returns a *new* `http.Handler`.

`func Logging(next http.Handler) http.Handler`.
Inside the returned handler, I do my pre-processing (logging start time), call `next.ServeHTTP(w, r)`, and then do post-processing (logging duration).
This **Chain of Responsibility** pattern is elegant because I can wrap layers infinitely: `Auth(RateLimit(Logger(Handler)))`."

#### Indepth
A critical detail in middleware is **Conditionality**. You might want to skip authentication for the `/health` endpoint. You can handle this inside the Auth middleware by checking `r.URL.Path`, or better, by wrapping only the specific sub-routers that need protection, leaving public routes outside the wrapper.

---

### 146. How do you serve static files in Go?
"I use `http.FileServer`.
`http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))`.

In production, I prefer to embed the assets into the binary using **go:embed**.
`//go:embed assets`
`var assets embed.FS`
Then I serve from that `embed.FS`. This gives me a single-file deployment—no need to copy a separate 'static' folder to the server."

#### Indepth
`go:embed` can also match patterns: `//go:embed css/*.css`. The embedded filesystem is read-only and is efficient (it maps directly to the binary's data segment). This solves the "works on my machine, fails in docker" because the assets are physically inside the executable.

---

### 147. How do you handle CORS in Go?
"Cross-Origin Resource Sharing is handled via Middleware.

I write a wrapper that sets headers like `Access-Control-Allow-Origin: *`.
Crucially, it must intercept **OPTIONS** requests (preflight) and return 200 OK immediately.
I often use the `rs/cors` library because getting the headers exactly right for all edge cases (credentials, exposed headers) is tedious and error-prone manually."

#### Indepth
CORS is browser security, not server security. It prevents site A from reading data from site B via JS. If you are building a server-to-server API (like a webhook receiver), CORS is irrelevant. For public APIs, setting `Access-Control-Allow-Origin: *` is fine, but for internal apps, whitelist specific domains.

---

### 148. What are context-based timeouts in HTTP servers?
"They are my defense against slow clients (Slowloris attacks).

Use `http.TimeoutHandler` or configure `http.Server{ReadTimeout: 5s, WriteTimeout: 10s}`.
Inside the handler, I pass `r.Context()` to any blocking calls (DB, API). If the client disconnects or times out, the context is canceled, and my DB query aborts immediately. This prevents a single slow client from tying up a database connection."

#### Indepth
`http.Server` has distinct timeouts: `ReadTimeout` (time to read body), `WriteTimeout` (time to write response), and `IdleTimeout` (Keep-Alive limit). Setting `WriteTimeout` is dangerous because if it triggers, Go forcibly closes the TCP connection, potentially cutting off a valid JSON response mid-stream. Use `http.TimeoutHandler` for cleaner behavior.

---

### 149. How do you make HTTP requests in Go?
"I use `http.NewRequest` and a custom `http.Client`.

**I never use the default `http.Get()` in production.**
The default client has NO timeout. If the server hangs, my goroutine hangs forever. Eventually, I run out of file descriptors and crash.
I always instantiate: `client := &http.Client{Timeout: 10 * time.Second}`."

#### Indepth
The default client is a shared global variable. Modifying it (like `http.DefaultClient.Timeout = ...`) affects all packages using it, which is a race condition. Always create your own client instance. You can often reuse one client instance for the entire app to share the connection pool.

---

### 150. How do you manage connection pooling in Go?
"The `http.Client` handles it automatically via the `Transport`.

It keeps a pool of idle TCP connections open (Keep-Alive).
The catch: I **MUST** read the response body fully and close it (`resp.Body.Close()`).
If I fail to do this, the connection cannot be returned to the pool and remains in a `CLOSE_WAIT` state, eventually causing a connection leak."

#### Indepth
To safely drain the body, use `io.Copy(io.Discard, resp.Body)`. Just calling `Close()` isn't enough; if there are unread bytes on the wire, Go might close the TCP connection instead of reusing it. Draining allows the connection to be kept alive for the next request.

---

### 151. What is an HTTP client timeout?
"It depends on which timeout.

`http.Client.Timeout` is the **total** time limit for the interaction (Dial + TLS + Headers + Body).
If I need more granular control (e.g., '10s to connect, but 1 hour to download'), I use `net.Dialer` timeouts or `context.WithTimeout`. call."

#### Indepth
`context.WithTimeout` is usually superior because it propagates. If Service A calls Service B with a 5s timeout, and B calls C, passing the context ensures C knows it only has 4.9s left. `http.Client.Timeout` is hard/local and doesn't respect the remaining time budget of the incoming request.

---

### 152. How do you upload and download files via HTTP?
"For **Uploads**: I use `r.FormFile("file")`. Go parses the multipart form and gives me a file handle (either in memory or on disk).
For **Downloads**: I set `Content-Disposition: attachment` and stream the file to `w`.
I never read the whole file into a `[]byte`. I use `io.Copy(w, file)`. This uses a fixed 32KB buffer, so I can serve a 10TB file with 10MB of RAM."

#### Indepth
For un-trusted uploads, always wrap `io.Copy` with `io.LimitReader(r.Body, limit)` to prevent disk-filling attacks where a user claims to send 1kb but sends 100GB. Validation of file magic numbers (signatures) is also crucial, as relying on request Content-Type headers is insecure.

---

### 153. What is graceful shutdown and how do you implement it?
"It enables the server to finish existing requests before stopping.

I listen for `SIGINT` or `SIGTERM`.
When caught, I call `server.Shutdown(ctx)`.
This stops the listener immediately (so no new requests come in) but blocks until all active handlers return (or `ctx` expires). This is critical for deployments to ensure users don't see 502 Bad Gateway errors during a rollout."

#### Indepth
Kubernetes complicates this. When a Pod terminates, K8s removes it from the Service endpoints, but this propagation is asynchronous. You should sleep for ~5-10 seconds *before* calling Shutdown to allow the Load Balancer to stop sending new traffic, otherwise, you might kill requests that were in-flight during the Shutdown sequence.

---

### 154. How to work with multipart/form-data in Go?
"I use `r.ParseMultipartForm(maxMemory)`.

The `maxMemory` argument (e.g., 32MB) tells Go: 'Keep requests smaller than this in RAM; spill anything larger to temporary files on disk'.
Then I access files via `r.MultipartForm.File`. Cleaning up created temp files is usually handled automatically, but I can call `r.MultipartForm.RemoveAll()` to be sure."

#### Indepth
Multipart forms are slow to parse. If you are building a high-performance file upload service, consider using "raw" uploads (PUT binary body) instead of `multipart/form-data`. It saves the CPU cost of boundary parsing and MIME decoding.

---

### 155. How do you implement rate limiting in Go?
"I use the **Token Bucket** algorithm, usually via `golang.org/x/time/rate`.

I create a limiter per user (keyed by IP).
Middleware checks `limiter.Allow()`. If false, I return `429 Too Many Requests`.
For distributed systems, local memory isn't enough, so I implement the Token Bucket in **Redis** (using Lua scripts) to share the limit across all API instances."

#### Indepth
Rate limiting headers are important. Send `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset` to tell polite clients when to back off. Without these, clients will just blindly retry, effectively DDOS-ing your gateway with 429 errors.

---

### 156. What is Gorilla Mux and how does it compare with net/http?
"Gorilla Mux was the standard for years because `net/http` was too simple.

Mux allows Method-based routing (`.Methods("POST")`) and Regex paths (`/products/{id:[0-9]+}`).
However, as of **Go 1.22**, the standard library added these features! So now, I prefer standard `net/http` for new projects. Mux is still great, but it's in maintenance mode."

#### Indepth
One big difference: `net/http` uses exact matching or prefix matching. It prioritizes the *most specific* pattern. `/images/thumbnails/` wins over `/images/`. This behavior is robust and predictable, whereas regex-based routing orders often depend on declaration order, which is fragile.

---

### 157. What are Go frameworks for web APIs (Gin, Echo)?
"**Gin** and **Echo** are the two heavyweights.

They are faster than `net/http` (using Radix tree routers) and provide a lot of helpers: Data binding (`c.BindJSON`), Validation, Grouping routes (`v1 := r.Group("/v1")`).
I use them if I need to build a large API quickly. If I'm building a small microservice, I stick to the standard library to keep dependencies low."

#### Indepth
Gin uses a custom `Context` struct, which is **not** thread-safe. You cannot pass `c` to a goroutine because it resets after the handler returns. You must call `c.Copy()` to pass a safe snapshot to a background worker. This is a common source of panic in Gin apps.

---

### 158. What are the trade-offs between using `http.ServeMux` and third-party routers?
"**ServeMux** (Stdlib):
*   Pros: No dependencies, stable APIs, forward compatible.
*   Cons: Verbose for complex middleware chains or parameter extraction.

**Chi/Gin**:
*   Pros: Concise, powerful middleware ecosystem, fast parameter extraction.
*   Cons: External dependency risk.

I usually treat **Chi** as the sweet spot—it’s just a Router that plays nice with standard `http.Handler`."

#### Indepth
Performance differences (Radix tree vs Regex vs Map) rarely matter for typical web apps (< 10k RPS). The bottleneck is almost always the Database or Network I/O. Choose the router based on DevX (developer experience) and middleware ecosystem, not raw nanosecond benchmarks.

---

### 159. How would you implement authentication in a Go API?
"I typically use **JWTs** (JSON Web Tokens).

1.  **Login**: User sends credentials. I verify them against the DB.
2.  **Issue**: I sign a JWT containing the `user_id` and `exp` time.
3.  **Middleware**: For every protected route, I extract the token from `Authorization: Bearer <token>`, parse it, and verify the signature.
4.  **Context**: I inject the `user_id` into `r.Context()` so handlers know who is calling."

#### Indepth
Stateful sessions (Cookies + Redis) are better than JWTs if you need instant revocation (e.g., banning a user). With JWTs, you can't invalidate a token until it expires (unless you implement a blacklist, which defeats the stateless purpose). Choose the right tool for the security model.

---

### 160. How do you implement file streaming in Go?
"I rely on the universal `io.Reader` and `io.Writer` interfaces.

I don't load data. I pass the stream.
If I'm proxying a file from S3 to the user:
`s3Resp := s3.GetObject(...)`
`io.Copy(w, s3Resp.Body)`
The bytes flow from S3 -> Go -> Client in tiny chunks. This keeps my memory usage flat and minimal."

#### Indepth
When proxying, be sure to copy the implementation's `Flush()` behavior. Standard `io.Copy` buffers until the buffer fills. For real-time streaming (like server-sent events), you need to type-assert the `ResponseWriter` to `http.Flusher` and call `Flush()` after every write chunk.


## From 09 Databases ORMs

# 🟣 **161–180: Databases and ORMs**

### 161. How do you connect to a PostgreSQL database in Go?
"I use the standard `database/sql` package with a driver like `lib/pq` or `pgx`.

I call `sql.Open("postgres", connStr)`.
Crucially, this **does not** create a connection. It just sets up the pool.
I always call `db.Ping()` immediately after to verify the credentials and network reachability. If `Ping` succeeds, I know the app is ready to serve traffic."

#### Indepth
`connStr` often includes complex parameters like `sslmode=disable`, `connect_timeout=10`, or `search_path=public`. Use `url.Parse` or a library helper to build this string safely, rather than manual concatenation, to avoid formatting errors with special characters in passwords.

---

### 162. What is the difference between `database/sql` and GORM?
"`database/sql` is the low-level, standard interface.
It gives me full control over the SQL. I have to manually scan rows into variables (`rows.Scan(&name)`). It’s tedious but fast and explicit.

**GORM** is an ORM. It abstracts the SQL away.
I can just call `db.Save(&user)`. It handles the `INSERT`, the ID generation, and the timestamps. I use GORM for rapid prototyping, but I often drop down to raw SQL for complex reports."

#### Indepth
ORMs in Go often rely heavily on **Reflection** (`reflect` package) to map structs to tables. This adds a CPU overhead compared to code-generated mappers (like `sqlc`) or manual scanning. For write-heavy logic (INSERTs), this overhead is negligible compared to I/O, but for reading 10k rows in a tight loop, `database/sql` is significantly faster.

---

### 163. How do you handle SQL injections in Go?
"I strictly rely on **parameterized queries**.

I never concatenate strings like `'SELECT * FROM users WHERE name = ' + input`.
Instead, I use placeholders: `$1` (Postgres) or `?` (MySQL).
`db.Query("SELECT ... WHERE name = $1", input)`.
The driver sends the query template and the data separately, making injection mathematically impossible."

#### Indepth
One common mistake is using `fmt.Sprintf` for table names or `ORDER BY` clauses, which cannot be parameterized. `db.Query("SELECT * FROM " + tableName)` is unsafe. If you must have dynamic tables, use an allow-list map to validate the input string against known safe values before concatenating.

---

### 164. How do you manage connection pools in `database/sql`?
"The `sql.DB` object **is** a connection pool.

I configure it carefully:
`db.SetMaxOpenConns(25)`: Prevents my app from overwhelming the DB.
`db.SetMaxIdleConns(25)`: Keeps connections hot so I don't pay the handshake cost on every request.
`db.SetConnMaxLifetime(5 * time.Minute)`: Rotating connections prevents stale-socket issues (like firewalls silently dropping idle connections)."

#### Indepth
If `MaxOpenConns` is reached, `db.Query` will **block** and wait for a connection to be returned. This is a hidden source of latency spikes. Monitor the `db.Stats().WaitCount` metric. If it's high, you need to either increase the pool size (if DB CPU allows) or optimize your query duration.

---

### 165. What are prepared statements in Go?
"A prepared statement is a pre-compiled SQL query.
`stmt, err := db.Prepare("INSERT INTO log (msg) VALUES ($1)")`.

I use it when I'm running the exact same query thousands of times in a loop.
It saves the database from parsing and planning the query every single time. It also reduces network bandwidth because I only send the parameters, not the full query text."

#### Indepth
Statements are prepared *per connection*. If your connection pool recycles a connection, the statement must be re-prepared (drivers handle this transparently, but it costs a round-trip). Some drivers support **Client-Side Statement Caching** to mitigate this. Always close statements with `defer stmt.Close()` to prevent leaking resources on the database server.

---

### 166. How do you map SQL rows to structs?
"With the standard library, it's painful. I have to manually `rows.Scan(&u.ID, &u.Name, ...)` in the exact order of columns.

I prefer using **sqlx**.
It allows me to do `db.Get(&user, "SELECT ...")`.
It uses the `db` struct tags to automatically map columns to fields. It saves me from writing boilerplate and mismatch errors."

#### Indepth
`sqlx` also supports `Scan` into a slice of structs (`db.Select(&users, ...)`). Be careful with `NULL` values. If a database column is NULL, scanning into a `string` (or `int`) will error. You must use `sql.NullString` or `*string` to handle potential nulls gracefully.

---

### 167. What are transactions and how are they implemented in Go?
"A transaction (`Tx`) ensures atomicity.

`tx, err := db.Begin()`.
I perform multiple queries using `tx.Exec` (not `db.Exec`).
Finally, I call `tx.Commit()` to save or `tx.Rollback()` to undo.
I typically use `defer tx.Rollback()` at the start. If the function panics or returns early, it automatically rolls back. If I successfully commit, the rollback does nothing."

#### Indepth
Transactions lock rows. If your business logic inside the transaction involves a slow operation (like an HTTP call to a payment gateway), you hold those DB locks for the duration of the HTTP call. This destroys database concurrency. Always keep transactions as short as possible—logic first, then DB lock, then update, then commit immediately.

---

### 168. How do you handle database migrations in Go?
"I use a dedicated tool like **golang-migrate** or **Goose**.

My migrations are versioned SQL files: `20231001_create_users.up.sql`.
I run these as part of the deployment pipeline (e.g., `migrate up`).
I never rely on ORMs like GORM to 'AutoMigrate' in production. It’s too risky—I need to know exactly what index changes or table locks are happening."

#### Indepth
Migrations should be **idempotent** if possible, but SQL often isn't. Running a migration twice usually fails ("table already exists"). To handle this in k8s, use a `Job` with `restartPolicy: OnFailure` that runs `migrate up`. Ensure your app waits for the migration to complete (using an `initContainer`) before starting.

---

### 169. What is the use of `sqlx` in Go?
"It’s an extension of the standard library.

It gives me `StructScan`, `NamedExec` (using `:name` instead of `$1`), and `Select` (for slices).
It doesn't hide the SQL like an ORM does; it just removes the tedium of mapping results. It’s the perfect middle ground between raw `database/sql` and heavy ORMs."

#### Indepth
`sqlx` mimics the standard library interface (`Query`, `Exec`), so it's easy to drop in. It also allows using maps for named queries: `db.NamedExec("INSERT ... VALUES (:name)", map[string]interface{}{"name": "Bob"})`. This is often cleaner than struct tags for partial updates.

---

### 170. What are the pros and cons of using an ORM in Go?
"**Pros**: Velocity. I can write CRUD apps in minutes. Relationships (Preload) and basic joins are handled for me.
**Cons**: Performance penalties (reflection), hidden N+1 queries, and difficulty debugging generated SQL.

Idiomatic Go tends to prefer **explicit** over implicit, so many teams start with GORM and migrate to sqlx or sqlc as the project scales."

#### Indepth
**sqlc** is a popular alternative. You write raw SQL (`-- name: GetUser :one SELECT * ...`), and it generates type-safe Go code for you. It catches syntax errors at compile time and has zero runtime overhead (no reflection). It's effectively the "Reverse ORM".

---

### 171. How would you implement pagination in SQL queries?
"I avoid `OFFSET` for large tables because it gets slower as the page number increases (it scans and discards rows).

I use **Cursor-based Pagination** (Keyset Pagination).
`WHERE id > last_seen_id LIMIT 10`.
This uses the index on `id` to jump directly to the right row. It’s O(1) regardless of whether I'm on page 1 or page 1,000,000."

#### Indepth
The downside of Cursor Pagination is that you cannot jump to "Page 10" directly, nor can you easily implement "Previous Page" without complex reverse-query logic. You also need a unique, sortable column (often usage of time + uuid) to serve as the cursor.

---

### 172. How do you log SQL queries in Go?
"I check the driver documentation.
For GORM, it’s built-in: `gorm.Config{Logger: logger.Default.LogMode(logger.Info)}`.
For `database/sql`, I wrap the driver with a logging hook (like **sqlhooks**).
It intercepts every `Exec/Query` call, logs the SQL statement, arguments, and execution time. It’s invaluable for debugging slow queries."

#### Indepth
Be careful not to log sensitive data (PII/passwords) in the SQL arguments. Custom loggers should have a sanitization step. Also, logging every query in production will emit massive logs. Use sampling or only log queries that exceed a duration threshold (Slow Query Log).

---

### 173. What is the N+1 problem in ORMs and how to avoid it?
"It’s when I fetch N items, and then for *each* item, I accidentally trigger another query to fetch a related record.
1 query for Users + 100 queries for their Avatars.

I avoid it by **Eager Loading**.
In GORM: `db.Preload("Avatar").Find(&users)`.
This executes exactly 2 queries: one for users, and one `IN (...)` query for all avatars. Speedup is massive."

#### Indepth
Eager loading works well, but watch out for memory usage. Loading 10,000 users and their 100,000 avatars into RAM might crash the pod (`OOMKilled`). For bulk processing, disable preloading and process in batches using a cursor or `FindInBatches`.

---

### 174. How do you implement caching for DB queries in Go?
"I use the **Cache-Aside** pattern with Redis.

1.  Check Redis for the key `user:123`.
2.  If found (Hit), return it.
3.  If missing (Miss), query Postgres.
4.  Serialize the result (JSON/Protobuf) and write to Redis with a TTL (e.g., 5 mins).
5.  Return it.
This protects the primary database from read spikes."

#### Indepth
The "Thundering Herd" problem occurs when a cache key expires and 1000 requests hit the DB simultaneously. To solve this, use **singleflight** (from `golang.org/x/sync/singleflight`). It merges duplicate in-flight calls so only *one* DB query runs, and the result is shared with all 1000 waiters.

---

### 175. How do you write custom SQL queries using GORM?
"GORM has a `Raw` method for when the query builder is too restrictive.

`db.Raw("SELECT name, count(*) FROM users GROUP BY name").Scan(&result)`.
I use this for complex reports, CTEs (Common Table Expressions), or specific window functions that GORM doesn't support natively."

#### Indepth
Using `Raw` returns `*gorm.DB` but creates a potential separation between your Go structs and the result set. If you select fields that don't match the struct, they are zero-valued. Always verify that your Raw logic aliases columns correctly to match the struct field names (`SELECT count(*) as total ...`).

---

### 176. How do you handle one-to-many and many-to-many relationships in GORM?
"I define them in the struct tags.

`type User struct { Orders []Order }` (One-to-Many).
`type User struct { Groups []*Group \`gorm:"many2many:user_groups;"\` }` (Many-to-Many).
GORM handles the join table (`user_groups`) automatically.
When I save a User, GORM automatically inserts the records into the join table."

#### Indepth
Auto-save associations can be dangerous. If you update a User struct and accidentally zero out the `Groups` slice, GORM might define this as "remove all groups" depending on configuration (`FullSaveAssociations`). I often disable this feature (`db.Omit("Groups").Save(&user)`) and manage relationships explicitly to avoid data loss.

---

### 177. How would you structure your database layer in a Go project?
"I use the **Repository Pattern**.

I define an interface `UserRepository`.
`type UserRepository interface { Get(id string) (*User, error) }`.
The implementation (`PostgresUserRepository`) holds the `*sql.DB`.
This decouples the Service layer from the database. I can swap Postgres for a Mock in unit tests without changing a single line of business logic."

#### Indepth
Repository interfaces should be defined by the **consumer** (the Domain/Service layer), not the implementer. This follows the Dependency Inversion Principle. `package service` defines `type UserRepo interface {...}`, and `package postgres` implements it. This prevents the domain from depending on `package postgres`.

---

### 178. What is context propagation in database calls?
"It means threading the `context.Context` from the HTTP handler down to the DB driver.

I use `db.QueryContext(ctx, ...)` instead of `db.Query`.
If the user cancels the request (closes the browser), the `ctx` is canceled. The DB driver sees this and **terminates** the running query on the database server. It prevents 'ghost queries' from hogging resources."

#### Indepth
Most modern drivers support this, but not all operations are cancellable immediately (e.g., if the DB is stuck in a heavy CPU loop vs waiting on I/O). However, it frees the connection in the pool on the Go side immediately. Always propagate context.

---

### 179. How do you handle long-running queries or timeouts?
"I wrap the parent context with a Timeout.

`ctx, cancel := context.WithTimeout(req.Context(), 5 * time.Second)`
`defer cancel()`
I pass this `ctx` to the query. If the DB takes >5s, the query returns `context.DeadlineExceeded` immediately. It ensures my API handles failures gracefully instead of hanging indefinitely."

#### Indepth
Postgres has a server-side setting `statement_timeout`. It's good practice to set this at the session level via connection parameters as a fallback safety net, in case the application fails to cancel the context correctly.

---

### 180. How do you write unit tests for code that interacts with the DB?
"I use **go-sqlmock**.

It mocks the `database/sql` driver. I tell it:
'Expect a query matching `SELECT * FROM users` and return these fixed rows.'
This allows me to test my repository logic (e.g., proper error handling, row mapping) without needing a running Postgres instance. For integration tests, I spin up a real DB container."

#### Indepth
`go-sqlmock` is great for testing *that your code calls the library correctly* (args, order). It does **not** test if your SQL is valid syntax or matches your schema. For that, you need Integration Tests using **testcontainers-go**, which provides a real ephemeral Postgres for every test run.


## From 10 Tools Testing Ecosystem

# 🔴 **181–200: Tools, Testing, CI/CD, Ecosystem**

### 181. What is `go vet` and what does it catch?
"`go vet` is the built-in static analysis tool.

It catches logic errors that compile but are likely bugs.
Common examples: `Printf` arguments not matching the format string, unreachable code after a return, or passing a `sync.Mutex` by value (which breaks the lock).
I run it automatically in CI, but I usually rely on `golangci-lint` which includes `vet` plus many other checkers."

#### Indepth
`go vet` uses **heuristics**, meaning it's not 100% precise but very fast. It checks for things like "CopyLock" (copying a struct that contains a Mutex). If you ever see a "copylocks" error, you are likely introducing a race condition by copying a lock state instead of sharing the pointer.

---

### 182. How does `go fmt` help maintain code quality?
"It ends all wars about code style.

It rewrites my source code to a canonical format (tabs for indentation, specific bracket placement).
Because it’s standard, I can jump into any Go project in the world and read it immediately without adjusting to a custom style guide. I configure my editor to run it on save."

#### Indepth
`gofmt` is technically a printer. `goimports` is a superset of `gofmt` that also manages your import block (adds missing, removes unused). Most developers use `goimports` as their "format on save" tool. In modern editors, `gopls` handles both formatting and imports efficiently.

---

### 183. What is `golangci-lint`?
"It’s a linter aggregator—the Swiss Army Knife of Go code quality.
It runs dozens of linters in parallel (including `vet`, `staticcheck`, `errcheck`).

I configure it to be strict: it forces me to handle every error, avoid unused variables, and limit cognitive complexity.
It’s much faster than running tools individually because it reuses the Go compilation cache."

#### Indepth
You can define a `.golangci.yml` file in your repo root to clamp down on specifics. For example, enabling `wsl` (Whitespace Linter) enforces empty lines between assignments and returns, making code more readable. Enabling `gocritic` finds subtle performance and style issues.

---

### 184. What is the difference between `go run`, `go build`, and `go install`?
"`go run` compiles and executes the binary in a temporary directory. I use it for local development and scripts.
`go build` compiles the binary and leaves it in the current directory. I use it to verify that code compiles.
`go install` compiles and moves the binary to `$GOPATH/bin`. I use it for installing CLI tools like `gopls` or `golangci-lint` so I can run them globally."

#### Indepth
Since Go 1.16, `go install package@version` is the preferred way to install tools without polluting your project's `go.mod` file. `go get` is now deprecated for installing binaries and is strictly for adding dependencies to the current module.

---

### 185. How does `go generate` work?
"It’s a tool for code generation, triggered by comments.

I add a comment `//go:generate stringer -type=Pill` in my source.
When I run `go generate ./...`, it finds these comments and executes the command.
I use it heavily for generating Mocks (`mockery`), Protobuf code (`protoc`), or Enums (`stringer`). It allows me to automate the creation of boilerplate code."

#### Indepth
`go generate` is **not** part of the build process (`go build`). You must run it manually. A common pattern is to have a `Makefile` with a `generate` target that runs `go generate ./...` before building, ensuring that all generated mocks and protobufs are up to date.

---

### 186. What is a build constraint?
"Also known as a Build Tag. It tells the compiler *when* to include a file.

`//go:build linux` or `//go:build integration`.
I use it for **OS-specific code** (e.g., using syscalls on Linux vs Windows).
I also use it for **test separation**: `//go:build integration` keeps my slow integration tests out of the standard `go test ./...` cycle unless I explicitly pass `-tags=integration`."

#### Indepth
The syntax changed in Go 1.17. The old syntax was `// +build linux`. The new syntax is `//go:build linux`. The new compiler supports boolean expressions like `//go:build linux || (darwin && amd64)`. You should prefer the new syntax.

---

### 187. How do you write tests in Go?
"I write functions starting with `Test` in `_test.go` files.

`func TestAdd(t *testing.T) { ... }`.
I prefer **Table-Driven Tests**: I define a slice of structs (inputs and expected outputs) and loop over them.
Inside the loop, I verify `got := Add(tc.a, tc.b)`. If `got != tc.want`, I call `t.Errorf`. It’s clean, extensible, and covers many edge cases easily."

#### Indepth
For large test outputs (like multi-line strings or JSON), `t.Errorf` diffs can be hard to read. Use `github.com/google/go-cmp/cmp` to display a clean line-by-line diff of struct fields. It is much more readable than standard `reflect.DeepEqual` failure messages.

---

### 188. How do you test for expected panics?
"Panics are special because they crash the test runner.

I catch them using `defer`.
`defer func() { if r := recover(); r == nil { t.Errorf("expected panic") } }()`
I put this at the top of the test. Then I call the code that should crash. If it *doesn't* crash, the deferred function sees a `nil` recover and fails the test.
Libraries like `testify` have a helper `assert.Panics` that wraps this logic neatly."

#### Indepth
Don't use `assert.Panics` for normal error handling logic. Panics in Go are for *truly* exceptional, unrecoverable states (like an `init()` function with invalid config). If a function panics on bad user input, that is a bug; it should return an `error`.

---

### 189. What are mocks and how do you use them in Go?
"Mocks are fake implementations of interfaces.

I use **vektra/mockery** to generate them automatically from my interfaces.
If `Service` depends on `Database`, I pass a `MockDatabase`.
In the test: `mockDB.On("GetUser", 1).Return(&User{}, nil)`.
This isolates the Service logic. I verify that the Service calls the DB correctly, without needing a real database running."

#### Indepth
Be careful with "over-mocking". If you mock everything, you end up testing your mocks, not your code. If the logic is simple, prefer using a real in-memory implementation (e.g., a map instead of a DB) rather than a generated mock. This is called a "Fake".

---

### 190. How do you use the `testing` and `testify` packages?
"I use the standard `testing` package for the test structure (`t.Run`, `t.Parallel`).
I use **testify** for assertions.

Instead of writing `if got != want { t.Errorf(...) }`, I write `assert.Equal(t, want, got)`.
It provides better error messages (showing the diff) and makes the test code much more readable."

#### Indepth
`testify/require` is a sibling of `assert`. The difference is that `require.NoError(t, err)` calls `t.FailNow()` (stop test immediately), whereas `assert` calls `t.Fail()` (continue test). Use `require` for setup steps where continuing makes no sense (e.g., DB connection failed).

---

### 191. How do you structure test files in Go?
"I place them right next to the source code. `user.go` and `user_test.go` live together.

This allows me to test unexported functions (white-box testing).
If I strictly want black-box testing (testing only the public API), I use a different package name in the test file: `package user_test`. This forces me to import `user` and use it exactly like a consumer would."

#### Indepth
The `user_test` package pattern solves cyclic dependencies. If `user` needs to test integration with `auth`, but `auth` imports `user`, you can't test inside `user`. Moving the test to `user_test` breaks the cycle because `user_test` imports both `user` and `auth`.

---

### 192. What is a benchmark test?
"It measures the performance of a function.

`func BenchmarkHash(b *testing.B)`.
The framework calls my function `b.N` times. It automatically adjusts `N` (100, 1000, 1M) until it gets a stable timing measurement.
I use it to catch performance regressions or to compare implementation A vs Implementation B (e.g., `fmt.Sprintf` vs `strconv.Itoa`)."

#### Indepth
Run benchmarks with `go test -bench=. -benchmem`. The `-benchmem` flag shows memory allocations per operation. A function might be fast but generate 1000 allocations (GC pressure). Optimizing for **0 allocs/op** is often better for system stability than raw CPU speed.

---

### 193. How do you measure test coverage in Go?
"I use the built-in cover tool.
`go test -coverprofile=c.out ./...`.

Then I view it: `go tool cover -html=c.out`.
It opens a browser showing my code. Green lines are covered, red lines are not.
I aim for high coverage (80%+) but I don't obsess over 100%. I ensure the *critical path` and *error handling* branches are covered."

#### Indepth
100% coverage is often a vanity metric. It usually forces you to write useless tests for trivial getters/setters or error checks that "can't happen". Focus on branch coverage for complex logic. Use Codecov or Coveralls in CI to prevent coverage *regressions* in Pull Requests.

---

### 194. How do you test concurrent functions?
"Testing concurrency is tricky due to race conditions.

I use `sync.WaitGroup` or channels to synchronize the test with the goroutines.
Crucially, I **always** run with the Race Detector: `go test -race`.
Standard tests might pass even with a race, but the race detector will spot the unsynchronized memory access and fail the build."

#### Indepth
When testing concurrent code, don't use `time.Sleep()` to "wait for the goroutine". This leads to flaky tests (fails on slow CI). Always use `WaitGroup` or channels to synchronize deterministically. If you must wait for a condition, use `assert.Eventually` (polling).

---

### 195. What is a race detector and how do you use it?
"It’s a compiler feature that instruments code to track memory accesses at runtime.

If two goroutines access the same variable concurrently, and at least one is a write, it’s a race.
I enable it with `-race`.
It slows down execution by ~10x, so I don't run it in production, but it is **mandatory** in my CI pipeline. It catches bugs that are almost impossible to debug manually."

#### Indepth
The Race Detector algorithm is based on **Vector Clocks** (ThreadSanitizer). It detects *unsynchronized* access. It does not detect *deadlocks* or *logical* races (e.g., A updates before B but you wanted B before A). It only proves that memory was accessed safely.

---

### 196. What is `go.mod` and `go.sum`?
"**go.mod** is the manifest. It lists the module name, the Go version, and the direct dependencies (with versions like v1.2.3).
**go.sum** is the lockfile/checksums. It contains the cryptographic hash of every module version used.

Its purpose is security. If an attacker hacks a library I use and changes the code for v1.2.3, the hash changes. `go build` notices the mismatch in `go.sum` and refuses to build, protecting my supply chain."

#### Indepth
`go.sum` is not a lockfile in the npm/cargo sense (it doesn't resolve the dependency tree). It's strictly a checksum database. You can have multiple versions of the same library in `go.sum` if different transitive dependencies ask for them. `go mod graph` shows the full tree.

---

### 197. How does semantic versioning work in Go modules?
"Go enforces SemVer strictly.

`v1.x.x` versions are compatible. I can upgrade safely.
`v2.x.x` (Direct Major Version) is treated as a **different module**. The import path changes to `github.com/lib/foo/v2`.
This allows me to use `v1` and `v2` of the same library in the same binary (Diamond Dependency problem solved), which is unique to Go."

#### Indepth
The compiler treats `github.com/foo/v2` as a completely different string than `github.com/foo`. They cannot be cast to each other. This strict separation allows the ecosystem to move forward without "DLL Hell", but it means you must upgrade your imports manually when migrating major versions.

---

### 198. How to build and deploy a Go binary to production?
"I build a static binary: `CGO_ENABLED=0 go build -o app`.
This removes dependencies on system libraries (libc).

I package it in a **Distroless** or **Scratch** Docker image.
The resulting image is tiny (10-20MB) and secure (no shell). I simply copy the binary and the CA certificates. This is the gold standard for Go deployments."

#### Indepth
`CGO_ENABLED=0` is key. If you don't set this, `net` package might dynamically link to the host's `glibc` DNS resolver. If your Docker container (Alpine/Scratch) doesn't have `glibc`, the binary will crash with "file not found". Static builds bundle the pure-Go DNS resolver.

---

### 199. What tools are used for Dockerizing Go apps?
"I use standard **Docker**.
I write a Multi-Stage Build.

Stage 1 (`golang:alpine`): Compiles the app.
Stage 2 (`gcr.io/distroless/static`): Copies only the binary.
I often use **Ko** (`ko build`), which builds OCI images directly from Go code without a Dockerfile. It’s incredibly fast and easy for Kubernetes deployments."

#### Indepth
`ko` is powerful because it analyzes your `import` paths to minimalize the image. It doesn't use `docker daemon`. It builds the tarball and pushes the layers directly to the registry. This is safer (no root privileges needed) and faster in CI environments.

---

### 200. How do you set up a CI/CD pipeline for a Go project?
"I almost always use **GitHub Actions**.

Step 1: **Lint**. `golangci-lint run`.
Step 2: **Test**. `go test -race -cover ./...`.
Step 3: **Build**. `go build`.
If all pass, I build the Docker image and push it to the registry. I define these steps in `.github/workflows/ci.yml`. It ensures no bad code never reaches the main branch."

#### Indepth
Cache your modules! `go mod download` can be slow. In GitHub Actions, use `actions/setup-go` with `cache: true`. It automatically hashes your `go.sum` and caches `~/go/pkg/mod`. This cuts build times from minutes to seconds for large projects.


## From 23 CLI Automation

# 🟢 **441–460: CLI Tools, Automation, and Scripting**

### 441. How do you build an interactive CLI in Go?
"I use **Bubble Tea** (Charm).

It implements the **ELM Architecture** for the terminal.
1.  **Model**: My application state.
2.  **View**: Renders the state to a string.
3.  **Update**: Handles keyboard events and modifies the state.
It allows me to build rich, 60fps TUI apps (like lists, spinners, and forms) that feel like native GUI applications."

#### Indepth
Bubble Tea is based on The Elm Architecture (Model-Update-View). This makes it strictly deterministic. However, handling **Async Commands** (like HTTP requests) requires returning a `tea.Cmd` from the `Update` function. This command runs in a separate goroutine and sends a `Msg` back to the `Update` loop when finished.

---

### 442. What libraries do you use for command-line tools in Go?
"**Cobra**: The standard for structure (Flags, Subcommands). used by Kubernetes/Docker.
**Viper**: For configuration (YAML/Env/Flags).
**Bubble Tea / Lip Gloss**: For interactive UI and styling.
**Pterm**: For simple printers (tables, progress bars) if Bubble Tea is overkill."

#### Indepth
Don't use `fmt.Print` in a Bubble Tea app! It corrupts the TUI buffer. Use the `tea.Program`'s output or `log` to a file. If you need to print a final result *after* the program exits (like `jq`), return the string in the Model and print it in `main` after `p.Run()` returns.

---

### 443. How do you parse flags and config in CLI?
"I use **Cobra** binding.

`rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")`.
Then inside `RunE`, I verify arguments.
For validation (e.g., 'port must be number'), I do it explicitly.
Most importantly, Cobra auto-generates `--help`, which is a UX requirement."

#### Indepth
Cobra supports **Persistent Flags** (`rootCmd.PersistentFlags()`) which filter down to *all* subcommands. Use this for global toggles like `--verbose` or `--json`. Also, use `viper.BindPFlag` so that `viper.Get("verbose")` works whether the user set the flag OR the environment variable.

---

### 444. How do you implement bash autocompletion for Go CLI?
"Cobra generates it for free!

`rootCmd.GenBashCompletion(os.Stdout)`.
I verify by running: `source <(my-cli completion bash)`.
Users get tab-completion for subcommands and even flags. I can also add custom dynamic completion (e.g., fetching a list of Kubernetes pods) by implementing the `ValidArgsFunction`."

#### Indepth
For `zsh` users, you need to generate zsh completion scripts (`GenZshCompletion`). A common pattern is to hide a `completion` subcommand that outputs these scripts so users can simply add `source <(my-cli completion zsh)` to their `.zshrc`.

---

### 445. How would you use `cobra` to build a nested command CLI?
"It’s a command tree.

`rootCmd.AddCommand(userCmd)`.
`userCmd.AddCommand(createCmd)`.
This gives me the structure: `my-cli user create`.
Each command is a struct with `Use`, `Short`, and `Run`. This organization keeps `main.go` tiny and separates the logic for 'User' and 'Product' into different packages."

#### Indepth
Structure is key. Put commands in `cmd/root.go`, `cmd/user.go`. Avoid global variables. Pass a `Context` or a `dependencies` struct (DB connection, Logger) to the `RunE` method of your commands via a closure or a struct method receiver, so your CLI remains testable.

---

### 446. How do you manage color and styling in terminal output?
"I use **Lip Gloss**.

It’s 'CSS for the Terminal'.
`var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA"))`.
`fmt.Println(style.Render("Hello"))`.
It automatically detects if the user's terminal supports TrueColor or just ANSI 16-color, and degrades gracefully. I never hardcode ANSI escape codes (`\033[31m`) anymore."

#### Indepth
Detecting "Is Terminal" is crucial. If the user pipes your output (`my-cli | grep foo`), you should **disable** colors automatically. Check `term.IsTerminal(int(os.Stdout.Fd()))`. Lip Gloss handles this, but if you do manual coloring, use `fatih/color` which respects the `NO_COLOR` standard.

---

### 447. How would you stream CLI output like `tail -f`?
"I read and flush.

`reader := bufio.NewReader(file)`.
I loop:
`line, err := reader.ReadString('\n')`.
If `err == io.EOF`, I sleep for 100ms and try again (polling).
To display it, I write to `os.Stdout`. If I need to overwrite the current line (like a spinner), I print `\r` (Carriage Return) before the new text."

#### Indepth
Be nice to the CPU! A tight loop reading `stdin` or polling a file can hit 100% CPU. always use a `select` with a `time.Ticker` or file watcher (`fsnotify`) to wait for changes efficiently. For `tail -f`, `fsnotify` is far superior to polling.

---

### 448. How do you handle secrets securely in a CLI?
"I never pass secrets as flags (`--password=123`).
That shows up in `bash_history` and `ps aux`.

I accept them via:
1.  **Environment Variables**: `MYAPP_PASSWORD=123 my-cli`.
2.  **Stdin**: `cat pass.txt | my-cli`.
3.  **Keyring**: I use `zalando/go-keyring` to store credentials securely in the OS's native Keychain (Mac) or Credential Manager (Windows)."

#### Indepth
If you must accept a password via flag (e.g. for automation scripts), provide a `--password-stdin` flag (like Docker). This allows `cat pass.txt | my-cli --password-stdin`, which keeps the secret out of the process arguments list and shell history.

---

### 449. How do you bundle a CLI as a standalone binary?
"That's Go's killer feature.

`CGO_ENABLED=0 go build -o my-cli`.
Result: A single, static binary.
No `node_modules`, no Python venv, no shared libraries.
I distribute this file. The user hits `chmod +x` and runs it. It works on Alpine, Debian, Centos—everywhere."

#### Indepth
Embed loop! `CGO_ENABLED=0` creates a static binary. But if you rely on `net` package + DNS, on Linux it *might* still use C library versions if they exist. Use `-tags netgo -ldflags '-extldflags "-static"'` to be 100% sure it's a static binary that runs on a generic Alpine container.

---

### 450. How would you version and release CLI with GitHub Actions?
"I use **GoReleaser**.

I push a tag `v1.0.0`.
GoReleaser detects it and:
1.  Cross-compiles for Linux/Mac/Windows (amd64/arm64).
2.  Creates a GitHub Release with artifacts.
3.  Updates my Homebrew Tap (`brew install my-cli`).
It automates the entire distribution pipeline."

#### Indepth
Sign your binaries! Code Signing is becoming mandatory (Apple Gatekeeper). GoReleaser integrates with `gon` or `cosign`. If you don't sign your Mac binary, users will get a "Developer cannot be verified" warning and likely trash your app.

---

### 451. How do you schedule a Go CLI tool with cron?
"The CLI is just a process.

I add it to system crontab: `0 * * * * /usr/local/bin/my-cli sync`.
If I want the Go app *itself* to be a long-running scheduler (daemon), I use **robfig/cron** library.
`c := cron.New(); c.AddFunc("@hourly", func() { ... })`.
This is useful inside Docker where system cron is missing."

#### Indepth
Distributed Cron? If you run 3 replicas of your app, `robfig/cron` will run the job 3 times! You need a **Distributed Lock** (Redis/Postgres). `if lock.Acquire() { job() }`. Or use a dedicated scheduler like Airflow/Temporal if complexity grows.

---

### 452. How do you use Go as a scripting language?
"I use `go run main.go`.

For single-file scripts, I use a shebang: `///usr/bin/env go run`.
However, if the script grows > 100 lines, I create a proper `go.mod` project. Writing complex automation in Go is safer than Bash (Type Safety, Error Handling, Testing) but more verbose."

#### Indepth
For "Scripting", look at **Go-Script** or **Bar** (build-and-run) tools. But honestly, compile times are so fast that `go run` is usually sufficient. A neat trick: `go run .` works if you have multiple files in `package main` in the current folder.

---

### 453. How do you embed templates in your Go CLI tool?
"I use `//go:embed`.

`//go:embed templates/*`
`var content embed.FS`.
I can generate a starter project for the user:
`my-cli init`.
The CLI reads the files from its own binary and writes them to the user's disk. This makes the CLI a completely self-contained generator without external dependencies."

#### Indepth
`embed` also supports the `http.FileSystem` interface. `http.FileServer(http.FS(content))`. You can embed an entire React/Vue Single Page App into your Go binary and serve it from memory. This is how tools like Grafana or Prometheus are distributed as single binaries.

---

### 454. How would you create a system daemon in Go?
"I use `kardianos/service`.

It abstracts systemd (Linux), Launchd (Mac), and Windows Services.
I define a `Program` struct with `Start()` and `Stop()` methods.
`s, _ := service.New(prg, config); s.Install()`.
This allows my Go app to install itself as a service that starts automatically on boot."

#### Indepth
Daemonizing correctly involves handling OS signals (`SIGTERM`, `SIGINT`). Your `Stop()` method should cancel a global context, allowing all running goroutines to clean up (close DB, flush logs) before the process exits. Hard kills lead to data corruption.

---

### 455. What are good patterns for CLI testing?
"I separate Logic from `main()`.

Instead of putting code in `main`, I have `func Run(args []string, stdout io.Writer) error`.
In tests, I pass a `bytes.Buffer` as stdout.
`err := Run([]string{"--dry-run"}, buf)`.
Then I assert `buf.String()` contains the expected output. This tests the wiring without needing to spawn a child process."

#### Indepth
Golden Files! CLI output often changes (formatting, spaces). Instead of `assert.Equal(t, "expected...", got)`, write the output to `testdata/output.golden`. In the test, compare `got` vs file content. Use a flag `go test -update` to overwrite the file when you intentionally change the format.

---

### 456. How do you store and manage CLI state/config files?
"I use `os.UserConfigDir()` (e.g., `~/.config/my-cli/`).

I save `config.yaml` there.
Viper handles reading it automatically.
I ensure I verify permissions (`0600`) if storing tokens. I use `os.MkdirAll` on startup to ensure the directory exists."

#### Indepth
Check `XDG_CONFIG_HOME`. Linux standards say config goes there, not hardcoded `~/.config`. `os.UserConfigDir()` handles this mostly, but being fully XDG compliant (`CACHE_HOME`, `DATA_HOME`) makes your CLI feel more "Pro" and native to Linux power users.

---

### 457. How do you secure a CLI for local system access?
"**Principle of Least Privilege**.

If my CLI needs to edit `/etc/hosts`, I check `os.Geteuid()`. If not root, I fail gracefully: 'Please run with sudo'.
But I try to avoid needing root.
I also validate all file paths to prevent **Symlink Attacks** (writing to a user-controlled link that points to `/etc/shadow`)."

#### Indepth
If you *must* run as root (e.g., a VPN client), try to **Drop Privileges**. Start as root, open the raw socket, then switch the process effective UID to the user `syscall.Setuid(uid)`. This minimizes the window where a bug could compromise the whole system.

---

### 458. How do you test CLI tools across multiple OS in CI?
"**GitHub Actions Matrix**.

`runs-on: [ubuntu-latest, macos-latest, windows-latest]`.
I run the Go tests on all three.
**Pain Point**: File Paths. I strictly use `filepath.Join` instead of hardcoding `/` or `\` to ensure compatibility."

#### Indepth
Line Endings! Windows uses `\r\n`, Linux `\n`. If your golden files use `\n`, your tests might fail on Windows. Use `strings.ReplaceAll(got, "\r\n", "\n")` in your test helpers to normalize output before comparison.

---

### 459. How do you expose analytics and usage for a CLI?
"I use a **Privacy-Preserving Ping**.

On execution, I fire a 'fire-and-forget' UDP packet or async HTTP POST to my telemetry server: `{'cmd': 'deploy', 'os': 'linux'}`.
I **MUST** ask for opt-in permission on the first run.
I wrap it in a short timeout (500ms) so analytics *never* slow down the user's experience."

#### Indepth
Respect `DO_NOT_TRACK` env var. If `os.Getenv("DO_NOT_TRACK") != ""`, disable analytics silently. Trust is hard to gain and easy to lose. Also, explicitly state what you collect in your `--help` or `README`.

---

### 460. How would you build a CLI wrapper for REST APIs?
"I generate the client using **OpenAPI Generator**.

Then I create Cobra commands for each endpoint.
`my-cli get users` -> `client.GetUsers()`.
I focus heavily on the Output Formatting. I use `tablewriter` to render the JSON response as a pretty ASCII table, which is much nicer for humans than raw JSON."

#### Indepth
JSON output flag is mandatory. `my-cli get users --json`. Power users want to pipe your output into `jq`. If you only output ASCII tables, they have to use `awk`, which is fragile. Always provide a machine-readable bypass for your human-readable output.


## From 28 Testing

# 🧪 **541–560: Testing in Go**

### 541. How do you write table-driven tests in Go?
"It's the idiomatic standard.
I define a struct array: `tests := []struct { input string; want int }`.
Then I loop: `for _, tt := range tests { t.Run(tt.name, ...) }`.
This makes it trivial to add 20 edge cases (empty strings, huge numbers) without copy-pasting logic. The `t.Run` creates subtests, allowing me to run specific cases via `go test -run=TestMyFunc/Case1`."

#### Indepth
`t.Parallel()` is the superpower of Table-Driven tests. By adding `tt := tt; t.Run(..., func(t *testing.T) { t.Parallel(); ... })`, you can run 100 test cases across all available CPU cores. *Note*: The `tt := tt` capture is not needed in Go 1.22+, as loop variables are safe.

---

### 542. What is the difference between `t.Fatal` and `t.Errorf`?
"**`t.Error`**: Marks failure but **continues execution**. I use this when checking multiple independent fields (e.g., Status Code is wrong, but I also want to see if the Body is wrong).
**`t.Fatal`**: Marks failure and **stops immediately**. I use this when setup fails (DB didn't connect, or file not found) because proceeding would just cause a panic."

#### Indepth
Use `t.Helper()` in your utility functions (like `assertUserID`). This marks the helper function as "not the cause of the error" in the stack trace, so when the test fails, the log points to the line in `TestMyFunc` that called the helper, not the helper itself.

---

### 543. How do you use `go test -cover` to check coverage?
"`go test -coverprofile=c.out ./...`.
Then `go tool cover -html=c.out`.
This opens a visual HTML report. Green lines are covered, red are missed.
I aim for high coverage (80%) in business logic packages. I ignore boilerplate/generated code.
Crucially, I check **branch coverage** (did I test both the `if err != nil` and the `else` path?)."

#### Indepth
For critical packages, enable `atomic` mode: `go test -race -covermode=atomic`. The default coverage mode is not thread-safe and can panic if tests run in parallel. Atomic mode uses `sync/atomic` counters, which is slower but correct for concurrent code.

---

### 544. How do you mock a database in Go tests?
"Two approaches:
1.  **Interfaces**: I define a `Repository` interface. In tests, I inject a `MockRepository` (manual or generated by `vektra/mockery`).
2.  **Docker (TestContainers)**: I spin up a real ephemeral Postgres for the test.
I prefer Docker for integration tests because SQL mocking is brittle—a mock returning success doesn't prove my SQL syntax is valid."

#### Indepth
If using `pgx` (Postgres driver), consider `pgxmock` which mimics the compiled binary protocol. But nothing beats `testcontainers-go`. It spins up a fresh Postgres container in 500ms, runs migrations, tests against real DB constraints, and kills the container. It's the gold standard for integration tests.

---

### 545. How do you unit test HTTP handlers?
"I use `net/http/httptest`.
`req := httptest.NewRequest("GET", "/users", nil)`.
`w := httptest.NewRecorder()`.
`myHandler(w, req)`.
`resp := w.Result()`.
This runs the handler directly in memory without opening a network port. It’s incredibly fast and lets me inspect headers, status codes, and JSON bodies easily."

#### Indepth
Test middleware too! Often bugs hide in the middleware chain (Auth, CORS) which `myHandler(w, req)` might skip if you test the handler function in isolation. Consistently test the "assembled" router (e.g., `router.ServeHTTP(w, req)`) to ensure the full request lifecycle works.

---

### 546. What is testable design and how does Go encourage it?
"Testable design means **Dependency Injection**.
Instead of `db := sql.Open(...)` inside my handler (hard to test), I pass the DB as a dependency to the handler struct.
`func NewHandler(db DBInterface) *Handler`.
Go's implicit interfaces make this pattern natural. I define the interface where I *use* it, allowing me to mock anything easily."

#### Indepth
Beware of "Interface Pollution". Don't define a big `UserStruct` interface with 20 methods. Define tiny interfaces where you *use* them: `type UserFetcher interface { Fetch(id int) User }`. This "Consumer-Defined Interface" pattern makes mocking trivial (just mock 1 method) and keeps dependencies loose.

---

### 547. How do you use interfaces to improve testability?
"I define consumer-driven interfaces.
If my service needs to send emails, I define:
`type EmailSender interface { Send(to, body string) error }`.
In production, I inject `SMTPSender`.
In tests, I inject `MockSender` that just records the call. This decouples my business logic from the slow, external SMTP server."

#### Indepth
Mocking variadic functions or functions with complex structs can be tedious. Generated mocks (Mockery/Gomock) are fine, but "Hand-written Mocks" are often clearer. `type MockSender struct { SendFunc func(...) error }`. This lets you swap behavior inline in the test: `m.SendFunc = func(...) { return error }`.

---

### 548. How do you write tests for concurrent code in Go?
"I use the **Race Detector** (`go test -race`).
And I use channels to synchronize.
`done := make(chan bool)`.
`go func() { doWork(); done <- true }`.
`<-done`.
For complex race conditions, I run the test in a loop `go test -count=100` to increase the chance of the scheduler triggering the specific interleaving that causes the bug."

#### Indepth
The `-race` detector adds ~10x CPU overhead and ~5-10x memory usage. Don't run it in production. But *always* run it in CI. It uses a "Happens-Before" vector clock algorithm to detect unsynchronized memory access with zero false positives.

---

### 549. What is the `httptest` package and how is it used?
"It provides utilities for HTTP testing.
**`ResponseRecorder`**: Records the response of a handler (status, body).
**`Server`**: Starts a real HTTP server on a random local port.
`ts := httptest.NewServer(http.HandlerFunc(...))`.
I use `ts.URL` as the API endpoint in my client tests. It simulates a real server environment."

#### Indepth
`httptest.NewServer` picks a random open port, preventing "Address already in use" errors during parallel tests. It also supports HTTP/2 (`NewUnstartedServer` + `EnableHTTP2`). Use it when testing your *HTTP Client* code (retries, timeouts) against a real (but local) network socket.

---

### 550. How do you mock time in tests?
"I never use `time.Now()` directly in logic.
I define a `Clock` interface: `Now() time.Time`.
In prod: `RealClock`.
In tests: `MockClock`.
`mockClock.Set(time.Date(2023, 1, 1...))`.
This allows me to verify logic like 'token expires in 1 hour' deterministically, without sleeping for an hour in the test."

#### Indepth
For strict time testing, use a library like `glock` or `clockwork`. They provide a `FakeClock` that lets you `Advance(1 * time.Hour)`. This triggers all waiting `time.After/Sleep` channels instantly, making "wait for 1 hour" tests run in microseconds.

---

### 551. How do you perform integration testing in Go?
"I separate them with build tags.
`//go:build integration`.
In the file: `func TestUserSignup_Integration(t *testing.T)`.
I run unit tests fast: `go test -short ./...`.
I run integration tests slow: `go test -tags=integration ./...`.
This keeps the developer feedback loop fast (milliseconds) while ensuring correctness in CI (seconds)."

#### Indepth
Create a `Makefile` or `Taskfile` to simplify this. `make test` runs unit tests. `make test-all` runs everything. Also, integration tests often require env vars (`POSTGRES_DSN`). Use `os.Getenv` in `TestMain` to skip the test (or `t.Skip()`) if the environment isn't set up (e.g., on a laptop without Docker).

---

### 552. How do you use `testify/mock` for mocking dependencies?
"I use `mockery` to generate the code.
`mockery --name=Database`.
In tests:
`mockDB := new(mocks.Database)`.
`mockDB.On("GetUser", 123).Return(user, nil)`.
It allows strict expectations: 'GetUser must be called exactly once with arg 123'. If the code calls it twice vs zero times, the test fails."

#### Indepth
Don't verify *everything*. `mock.Anything` is your friend. Over-specifying mocks (`On("Log", "User 123 logged in at 12:00").Return(nil)`) makes tests brittle; they break if you change a timestamp format. Verify the *side effects that matter* (DB writes), not the noise (logs).

---

### 553. How do you run subtests and benchmarks?
"**Subtests**: `t.Run("case 1", func(t *testing.T) { ... })`. Allows grouping and running specific cases.
**Benchmarks**: `func BenchmarkX(b *testing.B)`.
`b.Run("small payload", ...)`
`b.Run("large payload", ...)`
This gives me a comparative performance report for different inputs in a single run."

#### Indepth
Benchmarks are not just for speed; they track allocations! `b.ReportAllocs()`. If a commit changes `0 allocs/op` to `1 allocs/op` in a hot path, that's a regression. CI tools like `gobenchdata` can graph these trends over time to catch performance degradation.

---

### 554. How do you test panic recovery?
"I use a `defer-recover` block.
`defer func() { if r := recover(); r == nil { t.Errorf("did not panic") } }()`
`triggerPanic()`.
Or use `assert.Panics(t, func(){ ... })` from testify.
Testing panics is vital for library code where I want to ensure bad inputs (like index out of bounds) don't crash the user's application."

#### Indepth
Recover returns `any`, because you can `panic("string")` or `panic(error)`. Always type assert the result: `err, ok := r.(error)`. Also, remember `recover()` only works if called directly inside a `defer` function. It does nothing if called in a regular function nested inside the defer.

---

### 555. How do you generate test data using faker or random data?
"I use libraries like `gofakeit` or `go-faker`.
`email := gofakeit.Email()`.
`name := gofakeit.Name()`.
This prevents 'testing bias' where I always test with 'John Doe'.
Random data often reveals edge cases (e.g., names with apostrophes or emails with `+` signs) that my regex validation missed."

#### Indepth
Property-Based Testing (`gopter` or `rapid`). Instead of picking 1 random email, you define a property: "For *any* valid string s, `Reverse(Reverse(s)) == s`". The library generates thousands of edge cases (empty, emojis, control chars) to try and falsify your property.

---

### 556. What is golden file testing and when is it useful?
"For complex output (huge JSON/HTML).
I don't assert field-by-field.
I assert: `got == readFile("testdata/expected.json")`.
If I change logic intentionally, I run `go test -update`.
This overwrites the golden file with the new output. It makes maintaining tests for large data structures trivial."

#### Indepth
Git attributes! Mark golden files as generated or binary in `.gitattributes` so they don't clutter PR diffs. `testdata/*.golden -diff`. Or ensure `diff` uses a custom driver. This keeps your code review focused on Logic changes, not 500 lines of JSON output changes.

---

### 557. How do you automate test workflows with `go generate`?
"I put `//go:generate mockery --all` in `main.go`.
Running `go generate ./...` rebuilds all my mocks.
I also use it to generate Wire dependency injection code (`google/wire`).
It ensures that my generated test helpers are always up-to-date with my interfaces."

#### Indepth
The tool directive: Add `//go:generate go run github.com/vektra/mockery/v2@latest` to control the version of the tool used. Better yet, create a `tools.go` file with blank imports to track tool dependencies in `go.mod`, ensuring the whole team uses the exact same version of the code generator.

---

### 558. How do you test CLI apps built with Cobra?
"I redirect `os.Stdout` and `os.Stdin`.
Better yet, I inject `io.Reader` and `io.Writer` into my `RootCmd.RunE`.
In tests:
`buf := new(bytes.Buffer)`
`cmd.SetOut(buf)`
`cmd.Execute()`
Then I assert `buf.String()` contains the expected output. This tests the CLI wiring without spawning a child process."

#### Indepth
Testing Flags: Use `SetArgs([]string{"--flag", "value"})`. Don't rely on `os.Args`. Testing interactive prompts is harder; you might need a pseudo-terminal library or restructure your code to read from a generic `Scanner` interface that you can mock with a string buffer.

---

### 559. What is fuzz testing and how do you do it in Go?
"Go 1.18+ has native fuzzing.
`func FuzzParse(f *testing.F)`.
I seed it with valid inputs: `f.Add("valid")`.
Then `f.Fuzz(func(t *testing.T, data []byte) { ... })`.
The runtime generates random mutations of `data`.
I assert that my function **does not crash** (panic) and holds invariant properties (e.g., `Decode(Encode(data)) == data`)."

#### Indepth
Seed Corpus is vital. Go fuzzing saves "interesting" inputs that cause new code paths to `testdata/fuzz`. check these into Git! They become your regression suite. If a fuzz input found a bug, that input *must* be run forever to ensure the bug stays dead.

---

### 560. How do you organize test files and test suites?
"I place unit tests next to the code: `user.go` and `user_test.go`.
I use `package mypkg` for white-box testing (access private internals).
I use `package mypkg_test` for black-box testing (public API only).
For integration tests, I put them in a separate `tests/` folder if they span multiple packages, often using a `TestMain` for global setup/teardown."

#### Indepth
`internal` packages: You can't import `internal` from an external test package (`package foo_test`). If you need to test deep internals of `internal/auth` from `cmd/api`, you can't. Move the test *into* the package (`package foo`), or expose a "Test Hook" (export a variable only in `export_test.go`) to give the test access.


## From 40 Tooling DevExp

# 🧪 **781–800: Go Tooling, CI/CD & Developer Experience**

### 781. How do you create custom `go generate` commands?
"I create a tool (e.g., `gen-assets.go`).
In my source: `//go:generate go run gen-assets.go`.
Running `go generate ./...` executes it.
I use it for generating Mocks, Protobufs, or Embedding static assets. It keeps the build reproducible without external Makefiles."

#### Indepth
**Versioning Tools**. `go:generate` depends on tools installed in your `$GOPATH/bin`. If Developer A has `mockgen v1.6` and Developer B has `v1.7`, `go generate` produces different files. Use a `tools.go` file with `_ "github.com/golang/mock/mockgen"` imports to lock tool versions in `go.mod`, and run `go run github.com/golang/mock/mockgen ...` instead of relying on global binaries.

---

### 782. How do you build a multi-binary Go project?
"I structure `cmd/` folders.
`cmd/server/main.go`.
`cmd/worker/main.go`.
Build:
`go build -o bin/server ./cmd/server`.
`go build -o bin/worker ./cmd/worker`.
They share code from `internal/` and `pkg/` but produce distinct binaries."

#### Indepth
**Magefile**. Using `go build` commands manually is error-prone. Instead of Makefiles (which require `make` installed), use **Mage** (Make-in-Go). You write build scripts in Go (`mage.go`). This is cross-platform (works on Windows without WSL) and allows you to use the full power of Go for build logic.

---

### 783. How do you configure GoReleaser for automated builds?
"`.goreleaser.yaml`.
I define `builds` (linux/amd64, darwin/arm64) and `archives`.
On `git tag v1.0.0`, GitHub Action runs `goreleaser release`.
It builds, signs, and uploads everything to GitHub Releases automatically."

#### Indepth
**Snapshots**. For local testing, running a full release is slow. Use `goreleaser release --snapshot --rm-dist`. This builds everything locally without uploading to GitHub. It allows you to verify that the Docker images and binaries are generated correctly before tagging a real release.

---

### 784. How do you sign binaries in Go before release?
"I use **cosign** or standard GPG.
GoReleaser supports `sign` hooks.
It runs `cosign sign-blob` on the generated binary.
Users can verify the signature to ensure the binary wasn't tampered with. This is crucial for supply chain security."

#### Indepth
**SBOM**. Supply Chain attacks are real. GoReleaser helps generate an **SBOM** (Software Bill of Materials) in CycloneDX or SPDX format. This lists every dependency and version in your binary. Security scanners use this to quickly find if you are affected by a new vulnerability (like Log4j) without reverse engineering your binary.

---

### 785. How do you use `go vet` to detect issues?
"It runs automatically with `go test`.
It catches:
*   `Printf` format mismatches.
*   Unreachable code.
*   Struct tag typos.
I run it explicitly in CI: `go vet ./...` to block bad merges."

#### Indepth
**Shadowing**. One common bug `go vet` catches is **Variable Shadowing**, but it's not enabled by default. `err := foo(); if err != nil { err := bar() }`. The inner `err` shadows the outer one. Use `go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest` and run `go vet -vettool=$(which shadow) ./...` to catch this.

---

### 786. How do you manage environment-specific builds in Go?
"**Build Tags**.
`//go:build pro` vs `//go:build dev`.
`func config() string { return "production" }` (in `config_prod.go`).
`func config() string { return "localhost" }` (in `config_dev.go`).
`go build -tags prod`.
This excludes debug code from the production binary completely."

#### Indepth
**Integration Tests**. You don't want to run slow integration tests on every "Save". Add `//go:build integration` to `main_test.go`. Now `go test ./...` skips them (fast). Run `go test -tags=integration ./...` in CI to run them. This keeps the inner dev loop tight.

---

### 787. How do you use `build tags` in Go?
"To support multiple OS features.
`file_windows.go`: `//go:build windows`.
`file_unix.go`: `//go:build !windows`.
Go automatically picks the right file.
I also use it for integration tests: `//go:build integration` so they don't run during normal unit tests."

#### Indepth
**Boolean Logic**. You can do complex logic: `//go:build (linux || darwin) && !cgo`. This replaced the old comment syntax `// +build linux,darwin` which was confusing (comma = OR, newline = AND). Always use the new syntax (Go 1.17+).

---

### 788. How do you profile CPU/memory usage in CI pipelines?
"I write a Benchmark.
`go test -bench=. -benchmem -cpuprofile=cpu.out`.
I can compare the result against the `main` branch (**Benchstat**).
If performance degrades > 10%, I fail the build. It prevents regressions."

#### Indepth
**Noise**. CI environments are noisy (shared CPU). A benchmark might fail just because the runner was busy. Use `benchstat` to compare *multiple* runs (count=10). It uses statistical tests (Mann-Whitney U-test) to tell you if the difference is "Statistical Noise" or "Real Regression". Only fail on "Real".

---

### 789. How do you automate `go test` and coverage in GitHub Actions?
"Step 1: `actions/setup-go`.
Step 2: `go test -race -coverprofile=coverage.txt ./...`.
Step 3: Upload `coverage.txt` to Codecov.
It tracks trend lines: 'Coverage dropped 5%'. I enforce a hard floor (e.g., 80%) to pass PRs."

#### Indepth
**Goveralls**. If you don't use Codecov, use `mattn/goveralls` to send coverage to Coveralls.io. Note that "Line Coverage" is a weak metric. You can have 100% coverage and still have bugs. Use "Mutational Testing" (`test-mutation`) to see if your tests *actually* fail when the code is broken.

---

### 790. How do you write a custom Go linter?
"I use `golang.org/x/tools/go/analysis`.
I define an `Analyzer`.
`func run(pass *analysis.Pass)`.
I traverse the AST.
If I find a pattern (e.g., 'calling Log without Context'), I report `pass.Reportf`.
I plug this into `golangci-lint` as a custom plugin."

#### Indepth
**Ruleguard**. Writing AST analyzers is hard. Use `quasilyte/go-ruleguard`. It allows you to write linter rules in a Go-like DSL. `Match("fmt.Sprintf(\"%d\", $x)").Where(m["x"].Type.Is("int")).Report("Use strconv.Itoa($x) instead")`. This makes adding custom team conventions ("Don't use `log.Print`, use `zap`") trivial.

---

### 791. How do you automate versioning and changelogs in Go projects?
"I use **Conventional Commits** (`feat:`, `fix:`).
I use a tool like `release-please`.
It analyzes commit messages since the last tag.
It bumps the SemVer (Patch for fix, Minor for feat).
It generates `CHANGELOG.md` and creates the Release Tag."

#### Indepth
**SemVer Tricks**. If you are pre-v1 (`v0.2.3`), "Breaking Changes" don't require a Major bump (v1.0.0), just a Minor bump (`v0.3.0`). Once you hit `v1.0.0`, strict Semantic Versioning applies. `go mod` treats `v1` and `v2` as completely different packages (`github.com/foo/bar/v2`). Keeping `v0` for a long time gives you flexibility.

---

### 792. How do you use `go:embed` for bundling files?
"It embeds static assets into the Go binary.
`//go:embed static/*`
`var assets embed.FS`.
This turns my simple web server into a **single binary** deployment—no need to copy HTML/CSS files alongside the executable."

#### Indepth
**HTTP FileSystem**. `embed.FS` implements `fs.FS`. To serve it over HTTP: `http.FileServer(http.FS(assets))`. Warning: `embed` preserves the directory structure. If your file is in `static/index.html`, the user must visit `/static/index.html`. You typically need to `fs.Sub(assets, "static")` to "root" the server inside the folder.

---

### 793. How do you validate Go module versions in a monorepo?
"I use a workspace or a script.
I check that all modules use the same version of shared dependencies (e.g., `grpc v1.50`).
If `module-a` uses `v1.50` and `module-b` uses `v1.40`, `go build` might panic at runtime due to symbol mismatch. I enforce consistency in CI."

#### Indepth
**Go Workspaces**. Go 1.18 introduced `go.work`. It allows you to work on multiple modules locally without messy `replace` directives in `go.mod`. `go work use ./mod-a ./mod-b`. The editor (VSCode) sees them as one big project, allowing "Go to Definition" to jump across module boundaries.

---

### 794. How do you containerize a Go application for fast startup?
"**Multi-stage build**.
Build stage: `golang:1.24`. Compile static binary (`CGO_ENABLED=0`).
Final stage: `scratch` or `distroless/static`.
Copy binary.
Result: 10MB image. Starts instantly. No OS overhead."

#### Indepth
**Ko**. `ko` is a tool (by Google) that builds Go container images *without Docker*. It compiles the binary locally and wraps it in a tarball layer directly. It's faster than `docker build` and doesn't require a Docker daemon. Great for CI/CD pipelines (Kaniko alternative).

---

### 795. How do you enable live reloading for Go dev servers?
"I use **Air**.
Config: `air.toml`.
It watches `.go` files.
On change: Kills old process, rebuilds, restarts.
It feels like Node.js development."

#### Indepth
**Proxying**. While `Air` is running, if you hit a syntax error, the server crashes. The browser sees "Connection Refused". Better setup: `Air` runs a temporary *Proxy* on port 8080. It forwards traffic to your app on 8081. If your app crashes, the Proxy holds the connection and waits for the rebuild, preventing the "Site can't be reached" error.

---

### 796. How do you run multiple Go services locally with Docker Compose?
"`docker-compose.yml`.
Service A depends on DB. Service B depends on A.
I mount the source code and use `Air` inside the container for hot-reloading *inside* Docker networking context."

#### Indepth
**Host Networking**. On Linux, you can use `network_mode: host` to let the container share the host's IP/Ports. This removes the need for port mapping. On Mac/Windows, this doesn't work (VM isolation). Use `host.docker.internal` DNS name to access the "Host localhost" from inside the container.

---

### 797. How do you handle secrets securely in Go CI pipelines?
"GitHub Secrets injected as Env Vars.
`go test` reads `os.Getenv("API_KEY")`.
I verify *never* to print these secrets to the console log."

#### Indepth
**OIDC**. Long-lived API keys (AWS_ACCESS_KEY) in GitHub Secrets are a security risk. Use **OpenID Connect (OIDC)**. GitHub Actions exchanges a temporary JWT token with AWS/GCP to get short-lived credentials for that specific job. No static keys to rotate or leak.

---

### 798. How do you cross-compile Go binaries for ARM and Linux?
"`GOOS=linux GOARCH=arm64 go build`.
That's it.
No cross-compiler toolchain needed (unlike C++).
This makes it trivial to build for Raspberry Pi or AWS Graviton instances from my MacBook."

#### Indepth
**CGO Cross-Compile**. If `CGO_ENABLED=1` (e.g., using SQLite), cross-compilation is hell. You need a C cross-compiler (`aarch64-linux-gnu-gcc`). Use **Zig**. `CC="zig cc -target aarch64-linux" CGO_ENABLED=1 go build`. Zig acts as a drop-in C compiler that supports every target out of the box.

---

### 799. How do you build Go CLIs that auto-complete in Bash and Zsh?
"Using **Cobra**.
`cmd.Root().GenBashCompletion(os.Stdout)`.
I instruct the user to `source <(my-cli completion bash)`.
Cobra handles the magic of proposing flags and subcommands."

#### Indepth
**Fig/Carapace**. Cobra's bash completion is okay. For modern, rich auto-completion (with icons and detailed descriptions), tools like **Carapace** or **Fig** integrate with Cobra. They introspect your Go binary and generate specs for Zsh/Fish/PowerShell that feel like a native GUI menu.

---

### 800. How do you keep your Go codebase idiomatic and consistent?
"**Machine Enforcement**.
1.  `gofmt` (Formatting).
2.  `golangci-lint` (Linting).
3.  `goimports` (Import sorting).
I run these on `pre-commit` hook and CI.
Code Review focuses on logic/design, not style."

#### Indepth
**Revive**. `golangci-lint` includes `revive`. Revive is a faster, configurable drop-in replacement for `golint`. It allows you to disable specific annoying rules (like "Comment the exported function") or enforce strict new ones. It gives you control over the "Idiomatic" definition.


## From 42 Testing Part2

# 🧪 **821–840: Testing & Quality (Part 2)**

### 821. How do you implement contract testing in Go?
"I use **Pact**.
Consumer (Frontend) defines the contract: 'I expect GET /user to return {id, name}'.
Provider (Go API) verifies it fulfills the contract.
This prevents breaking changes between microservices without needing full End-to-End integration environments."

#### Indepth
**Pact Broker**. The real power of Pact comes from the **Broker**. It's a central registry where Consumers upload "Pacts" and Providers upload verification results. Before deploying Consumer v2 to production, the pipeline asks the Broker: "Is there a Provider version in prod that satisfies my new contract?". If no, deployment is blocked.

---

### 822. How do you run tests in parallel to speed up CI?
"`go test -p 8 ./...`.
Inside tests: `t.Parallel()`.
This allows one package's tests to run on multiple cores.
Constraint: My tests *must* be isolated (no shared global DB). If they rely on DB, I need to spin up a DB per package or use transactions that rollback."

#### Indepth
**t.Cleanup()**. When running parallel tests, `defer cleanup()` might run *after* the test function returns but *before* the parallel subtests finish (if using `t.Run`). Use `t.Cleanup(func() { ... })` instead of `defer`. `t.Cleanup` guarantees execution *after* the test and all its subtests are complete.

---

### 823. How do you manage test data in Go?
"I use **Fixtures** or **Factories**.
Factory: `NewUser(func(u) { u.Role = "admin" })`.
I avoid sharing huge JSON dumps between tests.
I create helper functions `createTestUser(t, db)` that insert minimal required data and return the object."

#### Indepth
**Go-CMP**. `reflect.DeepEqual` is brittle (it distinguishes `nil` slice from empty slice `[]int{}`). Use google's `go-cmp`. It allows ignoring unexported fields, approximating float comparisons, and sorting slices before comparing (`cmpopts.SortSlices`). It provides a readable diff of *exactly* what mismatched.

---

### 824. How do you verify log output in tests?
"I inject a custom `io.Writer` into the Logger.
In tests, I pass a `*bytes.Buffer`.
`logger := NewLogger(buf)`.
`if !strings.Contains(buf.String(), "error connecting") { t.Fail() }`.
This asserts that the expected warning/error was actually logged."

#### Indepth
**Testable Examples**. Go supports `func ExampleFoo()` in `_test.go` files. If you add a comment `// Output: hello`, `go test` runs the code and asserts stdout matches the comment. This verifies your documentation examples *are* actual running tests and never get out of date.

---

### 825. How do you test Go code that depends on `time.Now()`?
"I inject a Clock interface.
`type Clock interface { Now() time.Time }`.
Real implementation: `return time.Now()`.
Test implementation: `return fixedTime`.
This allows me to test logic like 'token expires in 1 hour' deterministically."

#### Indepth
**Time Travel**. If you use a `Clock` interface, you can simulate long durations instantly. `fakeClock.Add(24 * time.Hour)`. A test that verifies "Batch job runs after 24 hours" runs in milliseconds, not 24 hours. This is essential for testing timeouts, cache expirations, and cron jobs.

---

### 826. How do you use `httptest.ResponseRecorder` effectively?
"It implements `http.ResponseWriter`.
I pass it to my handler.
`handler.ServeHTTP(rec, req)`.
I inspect `rec.Code`, `rec.Body.String()`.
It’s an in-memory mock of a real browser connection."

#### Indepth
**Streaming Responses**. `httptest.NewRecorder` buffers the *entire* response in memory. If you are testing a streaming handler (Server-Sent Events or large file download), `Recorder` won't show the "flushing" behavior. You need to use `httptest.NewServer` and a real HTTP client to verify that data is arriving in chunks.

---

### 827. How do you test a graceful shutdown using signals?
"I send a signal to the process in my test.
`proc, _ := os.StartProcess(...)`.
`proc.Signal(syscall.SIGTERM)`.
I verify the process exits with code 0.
Internal unit test: I expose the `shutdown` channel in my `App` struct so the test can trigger it without yielding to the OS."

#### Indepth
**Context Propagation**. When shutting down, you typically give a 5-10 second timeout. In tests, reduce this to 10ms. Pass a `context.WithTimeout` to your Shutdown method. If the shutdown logic (closing DB, flushing logs) respects the context, the test will fail fast if something hangs, rather than hanging forever.

---

### 828. How do you check for goroutine leaks in tests?
"I use `goleak`.
`func TestMain(m *testing.M) { goleak.VerifyTestMain(m) }`.
It checks if any extra goroutines are running after the test finishes compared to before.
It catches `go func() { ... }` that never return."

#### Indepth
**Main Thread**. `goleak` ignores the main goroutine. It focuses on *background* workers. If you start a worker in `init()` (bad practice), `goleak` might assume it's a global system routine. Explicitly ignore known globals with `goleak.IgnoreTopFunction("my/pkg.backgroundWorker")` if they are intended to live forever.

---

### 829. How do you perform load testing on Go servers?
"I use **K6** or **Vegeta**.
They are external tools.
In Go, `httptest.NewServer` is okay for micro-benchmarks, but for rigorous load testing, I deploy the app and hit it from an external agent to measure throughput (RPS) and latency (p99) properly."

#### Indepth
**Scenarios**. Latency usually degrades not linearly, but drastically after a "cliff" (e.g., when DB connection pool fills up). Testing "100 RPS" is useless. Test a *Step Pattern*: 10 RPS -> 100 RPS -> 1000 RPS. Watch for the "Knee" of the curve where latency spikes. That is your capacity limit.

---

### 830. How do you use Go's `testing/quick` package?
"It’s QuickCheck for Go.
`quick.Check(func(x, y int) bool { return Add(x, y) == Add(y, x) }, nil)`.
It generates random `int`s and runs the function.
Deprecated now in favor of Fuzzing (`go test -fuzz`), which is smarter and native."

#### Indepth
**Seed Corpus**. Fuzzing starts random. But if you have a complex format (PDF), random bytes will just fail "Invalid Header" 99% of the time, exercising nothing. Provide a **Seed Corpus** (`f.Add([]byte("%PDF-1.4..."))`). The fuzzer mutates valid inputs to find edge cases deep in the parser logic.

---

### 831. How do you assert JSON responses in Go tests?
"I verify not to string compare! Formatting differs.
`require.JSONEq(t, expected, actual)` (using `testify`).
Or Unmarshal both into structs and compare structs.
`assert.Equal(t, expectedStruct, actualStruct)`.
This ignores whitespace differences."

#### Indepth
**Golden Files**. For complex JSON responses (50+ lines), defining `expectedStruct` in Go code is tedious and unreadable. Use **Golden Files**. Save `testdata/user_response.golden.json`. In test: `actual := handler()`; `if update { os.WriteFile("...golden", actual) }`; `expected := os.ReadFile(...)`. This makes updating tests trivial (run with `-update`).

---

### 832. How do you test code that uses `os.Exit`?
"You can't test `os.Exit` directly (it kills the runner).
Refactor: `func Run() int` (returns exit code).
Main calls `os.Exit(Run())`.
Test calls `Run()` and asserts it returns 1 or 0.
If needed, launch a subprocess (`exec.Command`) to test the actual crash."

#### Indepth
**Coverage**. Tests that run `exec.Command("go", "run", ...)` do *not* count towards code coverage of the main test run (it's a separate process). If you need coverage for the crash logic, you must compile a test binary that includes coverage instrumentation (`go test -c -cover`) and run that.

---

### 833. How do you structure table-driven tests for readability?
"Use named fields in the struct.
`tests := []struct { name string; input int; want int }{ ... }`.
`t.Run(tt.name, ...)` ensures the output shows 'TestAdd/negative_numbers'.
I keep the test body minimal, delegating setup/teardown to helpers."

#### Indepth
**Parallel Subtests**. A common trap: `for _, tt := range tests { t.Run(tt.name, func(t *testing.T) { t.Parallel(); do(tt.input) }) }`. This creates a race condition on `tt` (loop variable). In Go < 1.22, you MUST add `tt := tt` inside the loop. In Go 1.22+, loop variables are fixed, but be aware of older codebases.

---

### 834. How do you test middlewares in isolation?
"A middleware is `func(next) -> handler`.
I create a dummy 'next' handler.
`wrapped := MyMiddleware(dummyNext)`.
`wrapped.ServeHTTP(rec, req)`.
I verify that `MyMiddleware` did its job (set a header, logged) before or after calling `next`."

#### Indepth
**Context Keys**. Middleware often sets context values (`ctx = context.WithValue(ctx, UserKey, user)`). To test this, your dummy handler should check the context: `next = http.HandlerFunc(func(w, r) { if r.Context().Value(UserKey) == nil { t.Error("User not in context") } })`.

---

### 835. How do you mock file system operations in Go?
"I use `spf13/afero`.
It provides an `Fs` interface.
`var AppFs = afero.NewOsFs()`.
In tests: `AppFs = afero.NewMemMapFs()`.
I can create files in memory without touching the disk."

#### Indepth
**Testcontainers**. Mocks (`NewMemMapFs`, `sqlmock`) simulate behavior, but sometimes they simulate *wrongly*. "It works in mock but fails on real S3". Use **Testcontainers-go**. It spins up a *real* MinIO/Postgres container in Docker for the test. It's slower but gives 100% confidence that your SQL syntax/S3 API usage is correct.

---

### 836. How do you perform end-to-end (E2E) testing in Go?
"I build the binary. I spin up docker-compose (DB + App).
My test code uses an HTTP Client to hit the real running API.
I assert the state of the *Database* changed correctly.
This proves the entire stack works together."

#### Indepth
**Ephemeral Environments**. Running E2E tests on a shared "Staging" env is flaky (data collisions). Best practice: Spin up a full *Ephemeral Environment* (namespace in K8s) per PR. Run E2E tests against that isolated URL. Tear it down after. This ensures tests never fail because "Developer Bob deleted the user I was testing with".

---

### 837. How do you test GraphQL resolvers in Go?
"I test the Resolver methods directly.
`r := &Resolver{}`.
`resp, err := r.User(ctx, args)`.
I don't need to go through the HTTP transport. I trust the library (gqlgen) to call my resolver correctly; I only verify my resolver logic."

#### Indepth
**Schema Validation**. While you test resolvers generally, you should also have *one* integration test that sends a real GraphQL query string. This catches issues where the Schema (`schema.graphql`) doesn't match the Resolver implementation (e.g., you return an `int` but schema says `String`), which compile-time checks might miss in dynamic frameworks.

---

### 838. How do you use the `testdata` directory?
"Go ignores `testdata` during compilation.
I put huge JSONs, certificates, and CSVs there.
`data, _ := os.ReadFile("testdata/payload.json")`.
It keeps my test code clean."

#### Indepth
**go:embed**. Instead of `os.ReadFile` (which depends on the CWD being correct when running the test), use `//go:embed testdata/payload.json`. `var payload []byte`. This compiles the test data *into* the test binary. You can run the test binary from anywhere (`/tmp`) and it still has access to its data.

---

### 839. How do you test complex regex patterns?
"I start with unit tests.
`matches := myRegex.FindString(input)`.
I verify edge cases (empty string, huge string, unicode).
For security, I verify that the regex doesn't have ReDoS vulnerabilities by fuzzing requests with long repeating characters."

#### Indepth
**ReDoS**. Regular Expression Denial of Service. Patterns like `(a+)+` have exponential backtracking complexity. Go's `regexp` engine (RE2) guarantees linear time O(n) execution, so it is **immune** to ReDoS (unlike Python or Java). However, it is slightly slower and doesn't support features like backreferences.

---

### 840. How do you test channels and select statements?
"To test blocking behavior, I use timeouts.
`select { case val := <-ch: assert(val); case <-time.After(100*time.Millisecond): t.Fatal("timeout") }`.
To test non-blocking:
`select { case ch <- val: ; default: t.Fatal("should not block") }`."

#### Indepth
**Closed Channels**. A common bug is reading from a closed channel (which returns zero-value immediately) thinking it's valid data. In tests, use the two-value receive: `val, ok := <-ch`. Assert `ok` is true. If `ok` is false, it means the channel was closed unexpectedly, and your test logic might be flawed.


## From 47 Databases Part2

# 💾 **921–940: Go + Databases (SQL, NoSQL, ORMs)**

### 921. How do you use MongoDB with Go?
"I use the official `mongo-go-driver`.
`client, _ := mongo.Connect(ctx, options.Client().ApplyURI(...))`.
`coll := client.Database("db").Collection("users")`.
To query: `cursor, _ := coll.Find(ctx, bson.M{"role": "admin"})`.
I decode into structs using `cursor.All(ctx, &results)`.
I always set timeouts on the Context to avoid hanging queries."

#### Indepth
**BSON Struct Tags**. Using `bson:"field_name,omitempty"` is standard. BUT, be careful with `zerocopy` updates. If you want to unset a field in Mongo using struct, `omitempty` will just ignore it. You need to explicitly key off `msg["field"] = nil` using `bson.M` map updates for partial patches, or use pointer fields `*int` to distinguish between 0 and nil.

---

### 922. How do you store JSONB in PostgreSQL using Go?
"I implement the `sql.Scanner` and `driver.Valuer` interfaces on my struct.
`func (a *Attrs) Scan(value any) error { return json.Unmarshal(value.([]byte), a) }`.
Now I can pass the struct directly to `db.Exec("INSERT ... VALUES ($1)", myStruct)` and Postgres handles the JSON conversion."

#### Indepth
**GIN vs GIST**. When indexing JSONB in Postgres:
*   **GIN (Generalized Inverted Index)**: Best for "Contains" queries (`@>`). Faster reads, slower writes (heavy indexing).
*   **GIST**: Faster writes, slower reads. Good for geometric data.
In Go apps, 99% of the time you want GIN for searching JSON documents.

---

### 923. How do you index and search in Elasticsearch using Go?
"I use the official `go-elasticsearch` client.
Indexing: `client.Index("tweets", bodyReader)`.
Searching: `client.Search(client.Search.WithBody(queryReader))`.
Since the request body is JSON, I often use a builder pattern or struct to construct the complex Query DSL `{"query": {"match": ...}}`."

#### Indepth
**Bulk Indexing**. Sending 1 document at a time kills performance. Use `esutil.BulkIndexer`. It buffers documents in memory (e.g., up to 5MB or 1s delay) and sends them in a single HTTP request `_bulk` endpoint. This increases throughput from 100 docs/s to 10,000 docs/s.

---

### 924. How do you use Redis with Go for caching?
"I use `go-redis`.
Use cases beyond caching:
*   **Pub/Sub**: `rdb.Subscribe`.
*   **Rate Limiting**: `INCR` + `EXPIRE`.
*   **Queues**: `LPUSH` / `BRPOP`.
I treat Redis as a primary data structure server, not just a cache."

#### Indepth
**Pipelines**. Redis latency is mostly Round Trip Time (RTT). If you need to `SET` 100 keys, don't do it in a loop (100 network calls). Use `pipe := rdb.Pipeline(); pipe.Set(); ... ; pipe.Exec()`. This sends all 100 commands in 1 TCP packet. Latency drops from 100ms to 1ms.

---

### 925. How do you use prepared statements in Go?
"`stmt, _ := db.Prepare("SELECT * FROM users WHERE id = ?")`.
`defer stmt.Close()`.
`stmt.QueryRow(123)`.
It compiles the SQL on the server once.
Beneficial for repeated queries (bulk insert loop).
But `db.Query` often prepares automatically, so explicit preparation is only needed for high-perf loops."

#### Indepth
**SQL Injection**. Prepared statements are the #1 defense against SQL Injection. `SELECT * FROM users WHERE name = ' + user_input + '` is deadly. `?` acts as a placeholder that strictly treats input as *data*, never *code*. Go's `database/sql` forces this pattern naturally, making Go apps secure by default.

---

### 926. How do you prevent N+1 queries using Go ORM?
"In GORM: `Preload`.
`db.Preload("Orders").Find(&users)`.
It runs 2 queries: `SELECT * FROM users` then `SELECT * FROM orders WHERE user_id IN (...)`.
Without Preload, accessing `user.Orders` triggers a SQL query for *each* user (N+1 problem)."

#### Indepth
**Joins vs Preload**. `Preload` does 2 separate queries. `Joins` does 1 query with `LEFT JOIN`.
*   Use **Joins** when you need to *filter* by the child (Find users who bought "Apple").
*   Use **Preload** when you just need to *load* the child.
Joins transfer duplicated parent data (bandwidth heavy), Preload is cleaner but not atomic.

---

### 927. How do you map complex nested objects from DB in Go?
"If not using an ORM, I use `sqlx`.
`type User struct { Address Address `db:"address"` }`.
I join the tables: `SELECT u.*, a.city AS "address.city" ...`.
`sqlx` maps the dotted column names to the nested struct fields automatically."

#### Indepth
**Flat Structures**. Alternatively, some prefer defining a "Read Model" struct that is completely flat `type UserRow struct { UserID int; City string }` and then mapping it to domain objects manually. This avoids the reflection overhead of `sqlx` and gives you compile-time safety on the mapping logic.

---

### 928. How do you benchmark DB performance in Go?
"I don't just benchmark the Go code; I benchmark the *interaction*.
I write a test that runs `db.Exec` in a loop.
I observe latency histograms.
I use tools like `pgbench` for raw DB speed, and Go benchmarks to see if my driver/serialization is the bottleneck."

#### Indepth
**Driver Overhead**. Not all drivers are equal. `pgx` (Postgres) is significantly faster than `lib/pq` (which is effectively unmaintained) because `pgx` uses binary wire protocol instead of text protocol. Switching from `pq` to `pgx` can yield 20% perf gain for free.

---

### 929. How do you test DB queries with mocks?
"I use `go-sqlmock`.
`db, mock, _ := sqlmock.New()`.
`mock.ExpectQuery("SELECT").WillReturnRows(...)`.
I inject this `db` into my repo.
However, it doesn't prove the SQL works on the real DB (syntax errors). Integration tests are better."

#### Indepth
**Dockertest**. `go-sqlmock` is for *Unit Tests*. For *Integration Tests*, use `ory/dockertest`. It spins up a real ephemeral Postgres docker container from within your `TestMain`. You run real migrations and real queries. It proves your SQL syntax is valid for that specific Postgres version (e.g., 14.2).

---

### 930. How do you stream large query results in Go?
"`rows, _ := db.Query("SELECT * FROM wide_table")`.
`for rows.Next() { scan(); process() }`.
This streams row-by-row.
I **never** use `sqlx.Select` or `cursor.All` for huge datasets because they load everything into a slice (RAM OOM)."

#### Indepth
**Cursor**. In Postgres, a simple `SELECT` might still try to buffer results on the server-side. Use a **Cursor Transaction**: `BEGIN; DECLARE mycursor CURSOR FOR SELECT...; FETCH 1000 FROM mycursor;`. This ensures the DB server also streams data lazily instead of preparing 10GB of result set in RAM.

---

### 931. How do you use SQLite for embedded apps in Go?
"I use `mattn/go-sqlite3` (CGO) or `modernc.org/sqlite` (Pure Go).
`db, _ := sql.Open("sqlite3", "file:data.db")`.
I verify to turn on WAL mode: `PRAGMA journal_mode=WAL;` to allow concurrent readers and one writer."

#### Indepth
**Litestream**. SQLite is just a file. But how do you back it up? `Litestream` (written in Go) replicates the WAL (Write Ahead Log) to S3 in real-time. This gives you "Serverless Database" capability—if your server crashes, you restore from S3 with <1s data loss. Ideal for single-node Go apps.

---

### 932. How do you connect Go to Amazon RDS or Aurora?
"Standard Postgres/MySQL driver.
But for **IAM Auth** (passwordless), I use the AWS SDK to generate a token.
`token := rdsutils.BuildAuthToken(...)`.
I use the token as the password in `sql.Open`.
I must refresh this token every 15 minutes."

#### Indepth
**RDS Proxy**. IAM Auth Tokens are computationally expensive for RDS to verify (RSA decryption). If you open 100 connections/sec, you will spike RDS CPU. Use **RDS Proxy** to reuse connections. The Proxy handles the IAM Auth, and your Go app connects to the Proxy.

---

### 933. How do you manage read replicas in Go?
"I create two DB handles.
`Primary *sql.DB`.
`Replica *sql.DB`.
In my code: `if isWrite { Primary.Exec(...) } else { Replica.Query(...) }`.
Or I use a resolver middleware if using an ORM."

#### Indepth
**Replication Lag**. Reading from Replica is "Eventually Consistent". User updates profile -> Redirect to Profile Page -> Read from Replica -> Old Profile shown (Panic!). **Sticky Sessions** or **Read-After-Write** consistency is needed. A simple fix: "If user just wrote, read from Master for 5 seconds."

---

### 934. How do you handle DB failovers in Go apps?
"I rely on the driver's reconnection logic + connection pooling.
If the connection breaks, `db.Ping()` fails.
The driver attempts to reconnect.
For logic: I use retries (exponential backoff) on `db.Query`.
The DNS switch happens automatically, but my app must close broken connections to resolve the new IP."

#### Indepth
**Cloud SQL Connector**. For Google Cloud SQL or Azure PostgreSQL, Standard DNS failover can take minutes. The "Connector" libraries (Go SDKs) are smarter. They talk to the Cloud API to find the current IP of the master. They handle certificate rotation and automatic failover much faster than DNS TTL allows.

---

### 935. How do you use Migrations in Go?
"I use **Golang Migrate** or **Goose**.
I define up/down SQL files: `001_create_users.up.sql`.
I run `migrate up` locally or in CD pipeline.
It stores the current version in a `schema_migrations` table to prevent double execution."

#### Indepth
**Embed Migrations**. Don't rely on `file://` usage in production (files might be missing in the Docker image). Use `//go:embed migrations/*.sql` to compile SQL files into the binary. `golang-migrate` supports `io/fs`. This makes your binary truly self-contained—it can migrate its own DB on startup.

---

### 936. How do you handle transaction locking in Go?
"I use `tx.Exec("SELECT ... FOR UPDATE")`.
This locks the rows until the transaction commits.
It prevents race conditions where two goroutines read the same balance and update it simultaneously."

#### Indepth
**Optimistic Locking**. `FOR UPDATE` is Pessimistic/heavy. Alternative: Add `version int` column. `UPDATE accounts SET balance=?, version=version+1 WHERE id=? AND version=current_version`. If RowsAffected is 0, someone else updated it. Retry or fail. This scales better for low-contention systems.

---

### 937. How do you implement soft deletes in Go?
"Add `DeletedAt *time.Time` to the struct.
`UPDATE users SET deleted_at = NOW() WHERE id = ?`.
Queries must filter: `WHERE deleted_at IS NULL`.
GORM handles this automatically, but manually I need to be disciplined."

#### Indepth
**Unique Indexes**. Soft Deletes break unique constraints. `UNIQUE(email)`. User A deletes `bob@gmail.com`. User B tries to register `bob@gmail.com` -> DB Error "Duplicate", even though A is deleted. Fix: `UNIQUE INDEX ... WHERE deleted_at IS NULL` (Partial Index in Postgres).

---

### 938. How do you use Listen/Notify with Postgres in Go?
"I use `pq` or `pgx` driver.
`listener := pq.NewListener(...)`.
`listener.Listen("events")`.
`case n := <-listener.Notify: handle(n.Extra)`.
This allows real-time updates without polling the DB."

#### Indepth
**CDC**. Listen/Notify has a payload limit (8000 bytes) and isn't durable (if app is down, event is lost). For robust "Data Change" pipelines, use **CDC (Change Data Capture)** like Debezium. It reads the Postgres WAL and pushes changes to Kafka. Go consumers read Kafka. This ensures 100% data fidelity.

---

### 939. How do you handle connection pooling settings?
"Crucial for stability.
`db.SetMaxOpenConns(25)`.
`db.SetMaxIdleConns(25)`.
`db.SetConnMaxLifetime(5 * time.Minute)`.
If MaxOpen is too high, I starve the DB. If too low, my app blocks waiting for connections."

#### Indepth
**Timeouts**. `ConnMaxLifetime` is critical for Load Balancers (AWS ALB). If Go keeps a connection open for 1 hour, but ALB kills it silently after 5 minutes, Go will try to use a dead connection and get "Unexpected EOF". Set Go's lifetime to be *shorter* than the infrastructure's timeout.

---

### 940. How do you use generic repositories in Go?
"With Go 1.18 Generics.
`type Repository[T any] struct { db *sql.DB }`.
`func (r *Repository[T]) Find(id int) (T, error)`.
It reduces boilerplate, but I verify not to over-abstract. Sometimes specific queries need specific SQL optimization."

#### Indepth
**Interface Segregation**. Don't make one giant `Repository` interface. Split it. `Reader`, `Writer`. Or even better, `UserFinder`, `UserSaver`. This allows you to decorate just the `Finder` with a Cache layer without implementing the `Saver` methods. Composition over Inheritance.
