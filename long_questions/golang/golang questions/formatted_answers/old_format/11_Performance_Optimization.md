# 🟢 **201–220: Performance & Optimization**

### 201. How do you optimize memory usage in Go?
"I profile first. Guessing is a waste of time.
I use **pprof** to find the heavy allocations.

Common optimizations I apply:
1.  **Pre-allocate slices**: `make([]T, 0, 100)` avoids resizing overhead.
2.  **Object Pooling**: `sync.Pool` for reusing heavy structs like JSON encoders.
3.  **Value vs Pointer**: I check escape analysis. Sometimes passing a small struct by value is cheaper than a pointer because it stays on the stack and avoids GC entirely."

#### Indepth
Pre-allocating slices (`make([]T, 0, 100)`) is the single most effective "low hanging fruit" optimization. If you don't hint the capacity, Go starts small and doubles the array capacity repeatedly as you append (1->2->4->8...), involving massive copying and GC pressure. Always estimate the size if known.

---

### 202. What is memory escape analysis in Go?
"It’s a compiler phase that decides: *'Stack or Heap?'*

**Stack**: Fast, cleaned up automatically when the function returns.
**Heap**: Slow, requires Garbage Collection.
If I return a pointer to a local variable (`return &x`), the compiler sees that reference outlives the function, so it 'moves' `x` to the heap. I verify this with `go build -gcflags='-m'`."

#### Indepth
Escape analysis is conservative. If the compiler can't *prove* a pointer is safe (e.g., passed to `fmt.Println` which uses `interface{}`), it escapes to the heap. Understanding these rules allows you to write "stack-friendly" code, like returning values instead of pointers for small structs.

---

### 203. How to reduce allocations in tight loops?
"I move variable declarations **outside** the loop.

Instead of `for ... { var b bytes.Buffer ... }`, I create `b` once and `b.Reset()` inside.
I also use `strings.Builder` with `Grow(n)` for string concatenation.
Every allocation inside a hot loop (like a message processor) creates garbage that typically triggers a GC pause later, killing throughput."

#### Indepth
`strings.Builder` is optimized to let you build a string without copying the underlying bytes when you call `String()`. It uses `unsafe` under the hood to cast `[]byte` to `string`. This is why it's strictly faster than `bytes.Buffer` for string manipulation, though `bytes.Buffer` is better for general I/O.

---

### 204. How do you profile a Go application?
"I use the standard **pprof** tool.

For a web server, I import `net/http/pprof`.
I hit `curl localhost:6060/debug/pprof/profile?seconds=30`.
Then I analyze the file: `go tool pprof -http=:8080 cpu.prof`.
The **Flame Graph** view instantly shows me which function is hogging the CPU (e.g., usually it's JSON serialization or strict memory allocation)."

#### Indepth
Profiling in production is safe in Go because it has low overhead (~5%). However, don't leave it exposed to the public internet! The endpoints reveal sensitive system information. Bind the pprof server to `localhost` or protect it with an auth middleware.

---

### 205. What is the use of `pprof` in Go?
"It’s the built-in observability tool for the Go runtime.

It answers two specific questions:
1.  **CPU Profile**: Where is the app spending time? (e.g., `sha256.Sum256`)
2.  **Heap Profile**: Who is allocating memory? (e.g., `json.Unmarshal`)
It can also trace Goroutine blocking (mutex contention) and Thread creation. It’s indispensable for debugging 'why is my app slow?'"

#### Indepth
Don't forget the **Mutex Profile** (`go tool pprof --mutex ...`). It shows how much time goroutines spend *waiting* for locks, which CPU profiling misses (since verify rarely consumes CPU while waiting). This is critical for debugging contention issues in highly concurrent apps.

---

### 206. How do you benchmark against memory allocations?
"I use `go test -bench=. -benchmem`.

The `-benchmem` flag is key. It adds two columns to the output: `B/op` (bytes per op) and `allocs/op`.
If detailed optimization is needed, I use `b.ReportAllocs()` inside the benchmark function.
My goal is usually **Zero Allocs** for hot paths (0 allocs/op), which means the function runs entirely on the stack."

#### Indepth
While "Zero Alloc" is a noble goal, don't optimize prematurely. Converting `[]byte` to `string` (and back) usually allocates, but optimizing it away using `unsafe` requires careful maintenance. Only target the functions that show up in the top 10% of your CPU profile.

---

### 207. How can you avoid unnecessary heap allocations?
"I keep variables on the stack.

1.  **Avoid Pointers**: Pointers are harder for the compiler to prove 'safe', so they often escape.
2.  **Avoid Interfaces**: Assigning a concrete value to an `interface{}` always allocates (to store the type info).
3.  **Use Arrays**: `[32]byte` is passed on the stack; `[]byte` (slice) usually hits the heap if it grows."

#### Indepth
Interfaces are a common source of hidden allocations. When you assign a concrete value (like `int`) to an `interface{}`, Go must allocate a small structure on the heap to hold the type information and the value. If you do this in a tight loop, it generates substantial garbage.

---

### 208. What is inlining and how does the Go compiler handle it?
"Inlining replaces a function call with the actual body of the function.

It removes the call overhead (jumping, creating a stack frame).
The Go compiler automatically inlines small, leaf functions (e.g., 'getter' methods).
I can force it to tell me what it did with `go build -gcflags='-m'`. If logic is too complex (contains `defer` or `select`), the compiler won't inline it."

#### Indepth
Inlining is key because it enables further optimizations like **Dead Code Elimination**. If a function is inlined, the compiler can see that `if false { ... }` inside it is unreachable and delete the code entirely. This reduces binary size and instruction cache pressure.

---

### 209. How do you debug GC pauses?
"I run the application with `GODEBUG=gctrace=1`.

This prints a single line to `stderr` for every GC cycle.
`gc 1 @0.1s 1%: 0.5+1.0+0.5 ms ...`
I look at the 'wall clock' time. If I see frequent pauses >10ms, I know the GC is thrashing. algorithm is always: **Allocate Less**."

#### Indepth
Go's GC is **concurrent-mark-sweep**. It runs *alongside* your code. However, it still has brief "Stop The World" (STW) phases to turn on write barriers. In modern Go (1.8+), these pauses are sub-millisecond, but high allocation rates force the GC to run more often, stealing CPU cycles from your app.

---

### 210. What are some common performance bottlenecks in Go apps?
"1. **Serialization**: `encoding/json` relies on reflection and is slow.
2.  **GC Pressure**: Allocating millions of short-lived objects.
3.  **Lock Contention**: Too many goroutines fighting for a `sync.Mutex`.
4.  **Database Drivers**: Not using prepared statements or connection pooling correctly."

#### Indepth
Reflection is a performance killer. `encoding/json` scans struct tags at runtime. For high-throughput endpoints, switch to code-generation libraries like **easyjson**, or use **Protocol Buffers** which generate efficient marshalling code at compile time.

---

### 211. How to detect and fix memory leaks?
"Go is garbage collected, so leaks are rare, but they happen.
Usually, it’s a **Goroutine Leak**.

I start a goroutine that waits on a channel, but 10 hours later, the channel is never closed. The goroutine stays in RAM forever.
I detect this by looking at `pprof/goroutine`. If the count linearly increases over time, I have a leak. I fix it by ensuring every blocking receive has a timeout or a `ctx.Done()` case."

#### Indepth
A subtle leak happens with **Time Tickers**. `time.Tick` returns a channel that *never closes*. If you use it inside a short-lived loop or function, the ticker stays active forever. Always use `time.NewTicker()` and explicitly call `ticker.Stop()` when done.

---

### 212. How do you find goroutine leaks?
"I use **goleak** (by Uber) in my unit tests.

It scans the active goroutine stack at the start and end of a test.
If test A spawns a worker but forgets to stop it, `goleak` fails the test.
In production, I monitor the `go_goroutines` metric in Prometheus. A steady upward trend is a smoking gun."

#### Indepth
`goleak` works by capturing the stack trace of all running goroutines. It filters out standard runtime routines (GC, signal handling) and alerts you if any *user* goroutines are still running after `TestMain` finishes. It's a must-have for library authors.

---

### 213. How do you tune GC parameters in production?
"Traditionally, I set `GOGC`.
Default is 100 (run GC when heap grows 100%).
For memory-hungry batch jobs, I set `GOGC=200` (less frequent GC, more RAM usage).

In Kubernetes, I use the new **GOMEMLIMIT** (Go 1.19+).
`GOMEMLIMIT=350MiB` (for a 512MB pod).
This tells the GC: 'Be aggressive only when we get close to this limit'. It prevents OOM kills much better than tweaking `GOGC`."

#### Indepth
`GOMEMLIMIT` is a game changer for containerized workloads. Previously, Go had no idea it was running in a 512MB Docker container and would happily grow the heap until the OS killed it. Now it acts like Java's `-Xmx`, triggering GC aggressively as it nears the limit to stay alive.

---

### 214. How to avoid blocking operations in hot paths?
"I move the blocking work **Out of Band**.

If a user request requires sending an email (slow), I don't do it in the HTTP handler.
I push the job to a buffered channel or a Redis queue.
A background worker picks it up.
The user gets a `202 Accepted` response in 10ms, rather than waiting 2 seconds for the SMTP handshake."

#### Indepth
This is often called the **Outbox Pattern** or **Async Job Queue**. For reliability, don't just use an in-memory channel (which dies if the app crashes). Store the job in Redis, Postgres, or Kafka so it survives a restart.

---

### 215. What are the trade-offs of pooling in Go?
"**Pros**: Massive performance gain. Reusing a `[]byte` buffer avoids allocation and GC work.
**Cons**: Dangerous bugs.
If I forget to `buffer.Reset()`, the next user sees old data (Data Bleed).
I only use `sync.Pool` for objects that are allocated frequently and are expensive (like 64KB buffers or complicated structs)."

#### Indepth
`sync.Pool` is local to each P (Processor). When a goroutine on P1 puts an item in the pool, it stays on P1. This optimizes for locality (L1/L2 cache hits) and minimizes lock contention between threads. However, the pool is emptied during every Garbage Collection cycle, so it's only useful for *frequently allocated* objects.

---

### 216. How do you measure latency and throughput in Go APIs?
"I implement **Middleware**.

`start := time.Now()`
`next.ServeHTTP(w, r)`
`duration := time.Since(start)`
I record this duration in a Prometheus Histogram (`http_request_duration_seconds`).
This allows me to query `p99` latency. Measuring averages is useless because it hides the slow outliers that annoy users."

#### Indepth
Be careful with **High Cardinality** in metrics. If you record `http_request_duration_seconds` with a label `user_id`, and you have 1 million users, you will crash your Prometheus server. Only use bounded labels like `status_code` (200, 404, 500) or `method` (GET, POST).

---

### 217. What is backpressure and how do you handle it?
"Backpressure is saying 'No'.

When my system is overloaded, accepting more work will cause it to crash (OOM).
I implement backpressure using **Buffered Channels**.
If the buffer is full (`len(ch) == cap(ch)`), the sender blocks.
For APIs, I return **HTTP 429 Too Many Requests**. It’s better to fail 5% of requests fast than to crash the server and fail 100%."

#### Indepth
Another form of backpressure is **Load Shedding**. If the queue latency exceeds a threshold (e.g., 500ms), the service can proactively reject new requests *before* processing them, allowing it to catch up on the backlog. This prevents a "death spiral."

---

### 218. When should you prefer `sync.Pool`?
"Only when the Garbage Collector detects as the bottleneck.

If my profile shows 20% of CPU time in `runtime.mallocgc`, I reach for `sync.Pool`.
Typical targets: `bytes.Buffer`, `gzip.Writer`, or custom Request Context objects.
I never use it for database connections (use a driver pool) or simple things like `int` pointers."

#### Indepth
Don't use `sync.Pool` just to "be fast". It adds complexity (`Get`, type assertion, `Put`, `Reset`). If the object is small and short-lived, the stack allocator is faster and safer. Use `sync.Pool` only when escape analysis shows the object is hitting the heap and causing GC churn.

---

### 219. How do you manage high concurrency with low resource usage?
"I rely on Go's **Non-Blocking I/O**.

One goroutine uses 2KB of stack.
I can handle 10,000 concurrent WebSocket connections with ~500MB of RAM.
The key is to **not block** OS threads. I stick to standard Go networking (`net/http`), which uses the **Netpoller** to handle thousands of connections on just a few OS threads."

#### Indepth
This is the **Reactor Pattern**. The Go Runtime (Netpoller) uses `epoll` (Linux) or `kqueue` (macOS) to watch network sockets. When a socket is readable, the runtime wakes up the specific goroutine responsible for it. This is why Go servers scale better than "one thread per request" servers (like Apache or older Java).

---

### 220. How do you monitor a Go application in production?
"I use the **Observability Triad**.

1.  **Metrics** (Prometheus): Counters (`requests_total`) and Gauges (`memory_usage`). "Is it healthy?"
2.  **Logs** (Zap/Slog): Structured JSON. "What happened?"
3.  **Traces** (OpenTelemetry): "Where did it slow down?"
I expose `/metrics` and run a sidecar/agent to scrape it."

#### Indepth
Logging is the most expensive part of observability. Writing to `stdout` involves syscalls and mutexes. Use sampling (log only 1% of success requests) and buffering to keep performance high. `slog` (standard in Go 1.21) is highly optimized for this.
