# 🧵 **961–980: Concurrency Architecture & Design Patterns (Part 2)**

### 961. What is the circuit breaker pattern in Go?
"Protects my system from cascading failures.
If 'Service B' fails 10 times, the breaker **Trips (Open)**.
Calls to 'Service B' fail fast (no network call) for 30s.
Then **Half-Open** (allow 1 call).
If success, **Closed** (normal).
I use `gobreaker` or `hystrix-go`."

#### Indepth
**Bulkhead Pattern**. Circuit Breakers stop *outgoing* calls. Bulkheads stop *incoming* calls from crashing the whole system. Partition your connection pools/goroutines. "20 connections for API A, 20 for API B". If API A is slow, it consumes its 20 connections but leaves API B unaffected. Without Bulkheads, one slow dependency causes total system paralysis.

---

### 962. How do you implement message deduplication?
"Key idea: Idempotency Key in Redis.
`SETNX key 1`. If false, duplicate.
Expiration is crucial (e.g., 24 hours).
For strict dedupe, I use database unique constraints."

#### Indepth
**Bloom Filter**. Storing every RequestID in Redis gets expensive (RAM). Use a **Bloom Filter** (Probabilistic). It uses 2% of the memory. "Has this ID been seen?". Answer: "No" (Definitely not) or "Maybe" (Check DB). This filters out 99% of duplicates instantly with near-zero RAM cost, hitting the DB only for the "Maybe" cases.

---

### 963. How do you synchronize shared state across goroutines?
"1.  **Channels**: Communicate state changes (`state <- NewVal`). Monitor goroutine applies them.
2.  **Mutex**: `mu.Lock()`, Write, `mu.Unlock()`.
Channels are for passing flow/events. Mutexes are for checking/updating state flags. I use the right tool for the job."

#### Indepth
**Share Memory By Communicating**. The Go proverb. Don't use a Mutex to protect a shared `User` struct that 10 threads read/write. Instead, start *one* goroutine (Monitor) that owns the `User`. Other threads send `UpdateUserMsg` to the Monitor. The Monitor processes them sequentially. No Locks, No Deadlocks.

---

### 964. How do you detect livelocks in Go?
"Livelock: Two goroutines change state in response to each other but make no progress. 'I step left, you step left'.
Harder to detect than Deadlock (profiler shows CPU usage is high!).
Check: System is busy but throughput is 0.
Fix: Add **randomized** backoff/jitter to the retry logic."

#### Indepth
**Starvation**. A related issue. A high-priority goroutine (poller) grabs a lock, does work, releases it, and immediately grabs it again because it's in a tight loop. Low-priority goroutines trying to grab the lock never get a chance. `sync.Mutex` in Go 1.9+ has a "Starvation Mode" to prevent this, ensuring fairness after 1ms of waiting.

---

### 965. How do you timeout long-running operations?
"`ctx, cancel := context.WithTimeout(parent, 5*time.Second)`.
`defer cancel()`.
Pass `ctx` to the operation.
Inside op: `select { case <-ctx.Done(): return ctx.Err() ... }`.
Crucial: The function *must* respect the context, otherwise it keeps running in the background (goroutine leak)."

#### Indepth
**Context Cause**. Go 1.20 introduced `context.WithCancelCause(parent)`. Standard cancels just say "Canceled". `Cause` allows you to say "Canceled because Database X is down". This propagation of the *root error* through the context tree makes debugging timeout chains significantly easier.

---

### 966. How do you use the actor model in Go?
"Go channels are close to Actors.
Each Actor = 1 Goroutine + 1 Inbox (Channel).
`func Actor(inbox chan Msg) { for msg := range inbox { handle(msg) } }`.
Libraries like **Proto.Actor** provide supervision trees and remote actors, but for simple cases, raw goroutines are sufficient."

#### Indepth
**Supervision Strategy**. The core of Actor reliability (Erlang style). "One For One": If child crashes, restart child. "One For All": If child crashes, restart *all* siblings. Go goroutines have no hierarchy/supervision by default. You must implement this "Monitor" logic manually using `defer recover()` and restart loops.

---

### 967. How do you architect loosely coupled goroutines?
"**Pipeline Pattern**.
Stage A -> [Chan] -> Stage B.
Stage A doesn't know Stage B exists. It just writes to a channel.
Ideally, the channels are passed in: `func StageA(out chan<- Item)`.
This allows me to swap Stage B for Stage C easily."

#### Indepth
**Backpressure**. In a pipeline `A -> B -> C`, if C is slow, B fills its buffer, then A fills its buffer. Eventually A blocks (stops reading from socket). This is **Backpressure**. It naturally propagates up the stream. Using `unbuffered channels` gives instant backpressure (tight coupling), while `buffered channels` absorb spikes (loose coupling).

---

### 968. How do you design state machines in Go?
"`type State func() State`.
Loop: `state = state()`.
Each function represents a state. It returns the next state function.
If returns `nil`, machine stops.
This is elegant and allows complex transitions without a giant `switch` statement."

#### Indepth
**Lexical Scanning**. This pattern (by Rob Pike) is famously used in `text/template`. It's better than `switch state { case A: ... }` because each state function *decides* the next state dynamically. It mimics a "Goto" but in a structured, safe, and testable way.

---

### 969. How do you throttle a job queue in Go?
"Token Bucket or Semaphore.
Workers = 5. Queue = 1000 items. Currently running = 5.
This *is* throttling. The queue absorbs the spike.
To throttle the *enqueue* rate, I use middleware that checks a Redis rate limiter before pushing to the queue."

#### Indepth
**Priority Queues**. A standard channel is FIFO. If you need "VIP User" jobs to run before "Free User" jobs, you can't use a single channel. Use two channels `high` and `low`. `select { case job := <-high: do(job); default: select { case job := <-low: do(job) } }`. Note: This naive approach might starve low priority; handle with care.

---

### 970. How do you monitor goroutine health?
"**Heartbeat**.
Goroutine sends `tick` on a channel every 5s.
Supervisor listens.
If no tick for 15s, Supervisor assumes Goroutine is dead/stuck.
It can then restart it or alert."

#### Indepth
**Watchdog Timer**. Similar concept. Hardware Watchdogs reset the CPU if software hangs. In Go, a detached goroutine monitoring `last_activity_timestamp`. `if time.Since(last) > 1min { panic("deadlock detected") }`. This is a brute-force way to recover from unknown hangs in critical loops.

---

### 971. How do you track context propagation in goroutines?
"Pass `ctx` as the first argument. ALWAYS.
If I spawn a background goroutine that should outlive the request:
`newCtx := context.WithoutCancel(ctx)` (Go 1.21+).
This ensures Trace IDs are preserved but cancellation is detached."

#### Indepth
**Context Leaks**. `WithoutCancel` is dangerous if misused. If you detach the context, the background operation *never* knows if the parent request finished. You *must* add a new Timeout to the detached context (`ctx, _ = context.WithTimeout(ctx, 30s)`), otherwise you risk orphaned goroutines running forever.

---

### 972. How do you implement saga pattern in Go services?
"Distributed Transaction.
1. Service A: `DeductMoney`.
2. Service B: `DeliverProduct`.
If B fails, I must run **Compensating Transaction**:
3. Service A: `RefundMoney`.
I use an Orchestrator (Temporal.io) to manage this rollback flow reliably."

#### Indepth
**Dual Write Problem**. "Write to DB, then Publish to Kafka". If DB succeeds but Publish fails (Crash), system is inconsistent. **Outbox Pattern**. Write `(Data, Event)` to DB in *one transaction*. A separate poller reads `Event` table and pushes to Kafka. This guarantees Atomicity without distributed transactions (2PC).

---

### 973. How do you chain async jobs with error handling?
"Promise-like structure.
In pure Go:
`res, err := Step1()`
`if err != nil { return err }`
`res2, err := Step2(res)`
It’s verbose but explicit. 'Railway Oriented Programming' libraries exist to make this flatter."

#### Indepth
**ErrGroup**. For concurrent chains, use `errgroup`. `g.Go(func() error { return Step1() }); g.Go(Step2)`. `if err := g.Wait(); err != nil`. It runs tasks in parallel (or just manages the error propagation) and returns the *first* error encountered, canceling the other tasks (if configured with Context).

---

### 974. How do you log and trace concurrent tasks?
"Each goroutine needs a `TraceID` in its context.
When spawning `go func(ctx context.Context)`.
Log line: `log.With("trace_id", GetID(ctx)).Info(...)`.
This allows me to filter logs by TraceID in Kibana and see the interleaved logs of that specific task."

#### Indepth
**Goroutine ID**. Go deliberately hides the Goroutine ID (`goid`) to prevent Thread-Local Storage (TLS) abuse. Do NOT try to hack it using assembly. Instead, rely on explicit Context passing. If you absolutely need "Thread Local" (e.g., inside a deep 3rd party library hook), you are fighting the language design.

---

### 975. How do you create internal packages in Go?
"Put it in a directory named `internal/`.
`project/internal/mypkg`.
Go compiler enforces this: Only `project/...` can import `mypkg`.
`other-project` cannot import it.
This is the only way to enforce 'private' packages in Go."

#### Indepth
**Internal Modules**. If you have `github.com/my/lib/internal`, other repos cannot import it. But `github.com/my/lib/foo` CAN import it. It protects against *external* consumers, not intra-module dependencies. It is the strongest tool for clean API boundaries in library design.

---

### 976. How do you enforce code standards using golangci-lint?
"I put `.golangci.yml` in root.
Enable linters: `revive`, `gocritic`, `errcheck`.
Run `golangci-lint run` in CI.
If it fails, build fails.
This ends arguments about code style during PR reviews."

#### Indepth
**Nolint**. Sometimes the linter is wrong. Use `//nolint:gosec // Reason` to suppress it. ALWAYS include the reason. "It's safe because X". This documents the security exception for future auditors. Don't just blindly ignore errors to make the build pass.

---

### 977. How do you write makefiles for Go projects?
"`build:; go build -o bin/app .`
`test:; go test ./...`
I use `.PHONY` to avoid file name clashes.
Makefiles are the universal UI for build systems, standardizing commands across teams."

#### Indepth
**Taskfile**. `Make` is old and quirky (tabs vs spaces issue). **Task** (`go-task/task`) is a modern alternative written in Go. It uses YAML. `Taskfile.yml`. It supports cross-platform commands (works on Windows without WSL), parallel execution, and file fingerprinting (skip task if sources didn't change).

---

### 978. How do you manage secrets using Vault in Go?
"I use the Vault API client.
App starts (using a Kubernetes Service Account token).
Authenticates to Vault.
Reads secrets `secret/data/myapp/config`.
Values are kept in memory (not ENV).
This is more secure than K8s Secrets (base64)."

#### Indepth
**External Secrets Operator**. Instead of modifying your Go app to speak Vault API (Lock-in), use **External Secrets Operator** in K8s. It syncs Vault -> K8s Secret. Your Go app just reads standard environment variables (`valueFrom: secretKeyRef`). This decouples the App from the Secret Store implementation.

---

### 979. How do you deploy a Go app with Kubernetes?
"Dockerfile (Multi-stage).
Deployment YAML (Replicas=3).
Service YAML (ClusterIP).
Ingress YAML (Route traffic).
I use **Kustomize** or **Helm** to template these per environment (Dev/Prod)."

#### Indepth
**Ko**. If your app is pure Go, use `ko`. It builds the Go binary and creates a Docker image *without* a Dockerfile. It pushes to the registry and generates the K8s YAML with the image digest, all in one command `ko apply -f config/`. It produces tiny "Distroless" images by default.

---

### 980. How do you perform zero-downtime deployment in Go?
"**Rolling Update** (K8s default).
New Pod starts. Readiness probe passes.
Service switches traffic to New Pod.
Old Pod enters Terminating state.
Go app handles `SIGTERM`: stops accepting new requests, finishes old ones, exits.
End user sees no error."

#### Indepth
**PreStop Hook**. K8s updates are async. Service might send traffic to Terminating Pod for a few seconds. To fix: Add `preStop` hook: `sleep 10`. This ensures the Pod stays "Up" while the Load Balancer updates its routing table (draining), before the app actually receives `SIGTERM`.
