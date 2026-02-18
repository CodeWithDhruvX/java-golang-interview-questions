# 🟡 **381–400: Troubleshooting and Debugging**

### 381. How do you debugging a deadlock in Go?
"I check the **Stack Traces** of all goroutines.

When a Go program deadlocks (all goroutines asleep), the runtime crashes with a full dump.
I look for:
1.  `semacquire`: Waiting for a Mutex.
2.  `chan receive`: Waiting for a message.
If Goroutine A holds Lock 1 and waits for Lock 2, while Goroutine B holds Lock 2 and waits for Lock 1, that is my deadlock. To fix it, I enforce a strict Lock Ordering."

#### Indepth
Deadlocks are often deterministic but hard to reproduce. Use `go run -race` to catch them if valid, but better yet, analyze **Lock Ordering**. If Goroutine A locks `mu1` then `mu2`, and Goroutine B locks `mu2` then `mu1`, you have a potential deadlock. Always acquire locks in a consistent global order.

---

### 382. How do you analyze a memory leak in production?
"I take a **Heap Profile** from the running application.

`go tool pprof http://localhost:8080/debug/pprof/heap`.
I look at `inuse_space`. It shows me exactly which function is holding the most memory *right now*.
If I suspect a slow leak, I take two profiles 1 hour apart and use `pprof -diff_base`. This highlights the *growth* (delta), filtering out stable memory usage."

#### Indepth
A subtle leak: **Subslices**. `original := make([]int, 1000000); small := original[:10]`. If you keep `small` in memory, the entire backing array of 1 million integers is kept in memory! Use `slices.Clone()` (Go 1.21) to copy the data and drop the large backing array.

---

### 383. How do you debugging high CPU usage?
"I take a **CPU Profile** for 30 seconds.

The **Flame Graph** is my primary tool.
Wide bars mean 'time spent on CPU'.
Common culprits:
1.  **Serialization**: `json.Marshal` taking 40% of CPU.
2.  **GC Thrashing**: Failing to allocate memory fast enough.
3.  **Busy Loops**: A `for {}` loop without a `time.Sleep` or blocking call.
4.  **Regex**: Recompiling `regexp.MustCompile` inside a hot loop."

#### Indepth
If CPU is high but pprof shows nothing (all time in `runtime`), check `GODEBUG=schedtrace=1000`. You might have **Starvation** or excessive **Context Switching**. If you spawn 1 million short-lived goroutines per second, the scheduler overhead dominates the CPU.

---

### 384. What tools do you use for distributed tracing?
"I use **OpenTelemetry** with **Jaeger** or **Tempo**.

I ensure every request has a `Trace-ID`.
When a user says 'My request failed', I search for that ID.
The waterfall view shows me:
'API took 500ms... 300ms was waiting for Redis, 180ms was waiting for User Service, 20ms was logic'.
It pinpoints the bottleneck instantly across service boundaries."

#### Indepth
Tracing must be propagated! Use `context.Context` everywhere. If you drop the context in a goroutine (`go func() { ... }()`), you break the trace. Pass the context or create a detached context with the valid parent span ID to maintain the "Trace Waterfall".

---

### 385. How do you debug a panic in a production service?
"I rely on the **Panic Stack Trace**.

I catch the panic using `recover()` in my middleware.
I log the full stack trace to my centralized logging system (ELK/Sentry).
Without the stack trace, I'm guessing. With it, I know exactly which line caused the nil pointer or index out of range.
I fix the bug by adding a nil check or validating input, never by just suppressing the panic."

#### Indepth
`recover()` only works in the **same goroutine** as the panic. It does not catch panics from child goroutines. If you spawn `go func() { panic("boom") }()`, your server crashes even if the parent has a `recover`. You must add `defer recover()` to *every* goroutine you spawn.

---

### 386. How to use `delve` for debugging?
"**Delve** (`dlv`) is the Go debugger.

I use it locally for complex logic bugs.
`dlv debug main.go`.
*   `break main.go:42`: Set breakpoint.
*   `continue`: Run to specific line.
*   `print varName`: Inspect internal state.
*   `goroutines`: specific specific status of all blocked goroutines.
It’s much faster than adding `fmt.Println` and recompiling 50 times."

#### Indepth
Delve can attach to a running process: `dlv attach <PID>`. This is a lifesaver in staging environments where you can't just restart the app to add logs. You can inject non-breaking breakpoints (tracepoints) to print variables without stopping execution.

---

### 387. How do you debug race conditions that only happen occasionally?
"I use the **Race Detector**.

`go test -race -count=1000 ./mypackage/...`.
The standard race detector instruments memory accesses at runtime. Even if the race doesn't cause a crash during the test, the detector spots the *unsynchronized access* and flags it.
I treat every race report as a P0 bug, because in production it could mean data corruption."

#### Indepth
The Race Detector catches races that *actually happen* during execution. It doesn't prove code correctness. If your test suite doesn't trigger the specific timing overlap, the race detector won't complain. Run tests with `-count=10` and `-race` to increase the odds of catching flakes.

---

### 388. How do you analyze goroutine leaks?
"I check the **Goroutine Count** metric.

If `runtime.NumGoroutine()` climbs steadily over days, I have a leak.
I use `pprof` to see *where* they are stuck.
`go tool pprof http://localhost:8080/debug/pprof/goroutine`.
Usually, it’s a goroutine waiting on a channel read where the sender has already exited (or vice versa). The fix is ensuring every goroutine has a `ctx.Done()` exit path."

#### Indepth
Use `runtime.NumGoroutine()` in your health check handler. If the number grows linearly over time (monitor this in Prometheus), you have a leak. Common culprit: a goroutine writing to an unbuffered channel that no one is reading from anymore.

---

### 389. How do you debug network timeouts in Go?
"I categorize the timeout.

1.  **Dial Timeout**: Firewall or DNS. The server is unreachable.
2.  **Handshake Timeout**: TLS issue.
3.  **Response Header Timeout**: Server accepted connection but is overloaded/slow.
I use `httptrace` (ClientTrace) to measure each phase. This tells me if the problem is the Network (Dial slow) or the Server App (Response Header slow)."

#### Indepth
Use `net/http/httptrace`. It gives hooks for `GotConn`, `DNSStart`, `ConnectStart`. You can see if the "Timeout" was 99% DNS resolution time (DNS infrastructure issue) vs 99% Wait time (Server overloaded). This granularity is critical for debugging "flaky networks".

---

### 390. What is core dump analysis?
"A core dump is a snapshot of process memory at crash time.

I enable it with `GOTRACEBACK=crash`.
When the app panics, it writes a `core` file to disk.
I open it with `dlv core ./app core.1234`.
I can inspect variables and stack frames just like a live session. It’s the last resort for 'impossible' bugs that happen once a month in production but never locally."

#### Indepth
**Security Warning**: Core dumps contain the entire memory of the process, including **Secrets** (DB Passwords, Private Keys). Never send core dumps to an external vendor or upload them to public storage without sanitizing or encrypting them.

---

### 391. How do you verify if the GC is the bottleneck?
"I configure the **GC Tracer**.
`GODEBUG=gctrace=1`.

I watch the `PauseNs` and `%CPU` usage.
If the GC is stealing > 25% of my CPU cycles, or if Stop-The-World pauses exceed 10ms frequently, it is a bottleneck.
My solution is always to **Reduce Allocations** (reuse buffers, avoid pointers) rather than tweaking GC knobs."

#### Indepth
If GC CPU usage is > 25%, you are creating too much garbage. Look for **Memory Ballast** (allocating a huge byte array on startup) to trick the GC into running less often (legacy trick), or better, use `GOMEMLIMIT` (Go 1.19) to set a target memory usage, allowing the GC to utilize available RAM fully before marking.

---

### 392. How do you debug a slow SQL query in Go?
"I use a **Slow Query Logger**.

In my DB driver configuration, I set a hook: 'Log any query taking > 100ms'.
When a log appears, I verify the SQL.
I run `EXPLAIN ANALYZE` on the database to check for missing indexes / full table scans.
From the Go side, I ensure I'm not fetching 10,000 rows when I only process 10 (`LIMIT` missing)."

#### Indepth
In Go `database/sql`, check `db.Stats()`. If `OpenConnections == MaxOpenConnections`, your app is waiting for a free connection from the pool. The "slow query" might actually be "fast query, but waited 500ms for a connection". Increase pool size or fix the query holding the connection too long.

---

### 393. How do you handle transient network errors?
"I assume the network is flaky.

I wrap my HTTP client with a **Retry Middleware**.
It catches `503 Service Unavailable` or `Connection Refused`.
It retries with **Exponential Backoff** (1s, 2s, 4s).
**Crucial**: I only retry *Idempotent* operations (GET, PUT). I never retry a POST (Charge Payment) blindly unless I can verify it wasn't processed."

#### Indepth
**Jitter** is mandatory. If a microservice has a hiccups and 1000 clients retry exactly 1.0s later, they will DDOS the service again (Thundering Herd). `time.Sleep(1s + rand.Intn(500ms))` smooths out the retry wave and allows the service to recover.

---

### 394. How do you debug context cancellation issues?
"I inspect the **Cancellation Cause**.

`ctx.Err()` just says 'Canceled'.
In Go 1.20+, `context.Cause(ctx)` returns the *error* that caused the cancellation.
'DeadlineExceeded'? 'Client Disconnected'? 'Explicit Cancel called'?
This tells me if the user gave up (client disconnect) or if my backend was too slow (timeout)."

#### Indepth
Go 1.20 added `context.WithCancelCause(parent)`. This allows you to set *why* the context was canceled (e.g., "Timeout" vs "UserAbort"). Check `context.Cause(ctx)` to log the specific reason, stopping the "Why did this request cancel?" guessing game.

---

### 395. How do you monitor thread exhaustion?
"The Go runtime spawns OS threads (M) for blocking syscalls.

If I make too many CGO calls or blocking File I/O, the runtime spawns thousands of threads.
Eventually, I hit the OS `ulimit` or `runtime: program exceeds 10000-thread limit`.
I monitor the `runtime_thread_count` metric. If it spikes, I know I'm blocking the runtime, and I need to offload that work or rate-limit it."

#### Indepth
Go runtime creates a new OS thread if a goroutine blocks in a System Call (syscall). If you have 10,000 goroutines doing blocking file IO (without async IO poller support), Go *will* spawn 10,000 threads, crashing the app. Use `debug.SetMaxThreads` to prevent this global exhaust.

---

### 396. How do you debug incorrect JSON unmarshaling?
"Common Pitfalls:
1.  **Case Sensitivity**: Struct field `Name`, JSON `name`. Go handles this, but `name` (private field) is ignored.
2.  **Type Mismatch**: JSON `id: "123"` (string), Struct `ID int`.
I use `decoder.DisallowUnknownFields()` to catch typos in the JSON payload.
I verify the error returned by `Unmarshal`—it usually points to the exact byte offset of the mismatch."

#### Indepth
JSON numbers are standardly float64, which loses precision for large IDs. Use `decoder.UseNumber()` to decode numbers as `json.Number` (string) instead of float64. This preserves the exact value of 64-bit integers or high-precision decimals during unmarshalling.

---

### 397. How do you profile lock contention?
"I use the **Mutex Profile**.

`go test -bench=. -mutexprofile=mutex.out`.
`go tool pprof mutex.out`.
It shows me exactly which `sync.Mutex` processes are fighting over.
If contention is high, I fix it by:
1.  **Sharding**: Use 10 locks (one per bucket) instead of 1 global lock.
2.  **RWMutex**: Allow concurrent reads.
3.  **Reducing Critical Section**: Do less work while holding the lock."

#### Indepth
Use `sync.Map` for read-heavy, append-only caches. Standard `sync.Mutex` acts as a bottleneck because even readers contend for the lock. `sync.Map` (or `RWMutex`) allows concurrent readers, drastically reducing contention profiling hot spots in these specific scenarios.

---

### 398. How do you investigate 502 Bad Gateway errors?
"502 means the Load Balancer (Nginx) couldn't talk to my Go app.

Causes:
1.  **Crash**: App died (Check `restart_count`).
2.  **Hang**: App is stuck (GC pause or Deadlock).
3.  **Timeout**: App accepted connection but didn't write headers in time.
4.  **Socket Exhaustion**: No ports left.
I correlate the LB logs ('upstream timed out') with my app metrics to find the root cause."

#### Indepth
502 Bad Gateway often means "Keep-Alive Mismatch". If Go closes the idle connection after 30s, but the Load Balancer (AWS ALB) thinks it's open for 60s, the ALB sends a request to a closed socket. Always set Go's `IdleTimeout` slightly *higher* than the LB's idle timeout.

---

### 399. How do you debug issues that only appear in Docker/K8s?
"I assume Environment Differences.

1.  **OOMKill**: Is the container hitting its memory limit? (Exit Code 137).
2.  **CPU Throttling**: If I set `limits.cpu=100m`, K8s throttles my app heavily, causing latency.
3.  **DNS**: K8s uses CoreDNS. `musl` (Alpine) resolves differently than `glibc`.
I use `kubectl exec -it pod -- sh` to get inside and run `curl`/`nslookup` to verify the environment."

#### Indepth
Use **Ephemeral Containers** (`kubectl debug`). This allows you to attach a generic troubleshooting container (with `curl`, `dig`, `vi`) to a crashed or distroless pod. It shares the process namespace, so you can see the files and processes of the target pod even if it has no shell.

---

### 400. How do you maintain a 'Runbook' for debugging?
"I write **Executable Runbooks**.

Instead of a Word doc saying 'Check the database', I write:
'If Alert X fires:
1.  Run `scripts/debug_db.sh`.
2.  If output > 50 conns, scale up pool.'
This saves mental energy during a 3 AM outage. The goal is to make debugging mechanical and repeatable."

#### Indepth
Treat Runbooks as Code. Store them in Markdown in the repo. Better yet, make them **Executable Runbooks** (Jupyter Notebooks for Ops). If a human has to copy-paste commands during an outage, they will make mistakes. One-click remediation scripts are safer.
