# 🧵 Go Theory Questions: 961–980 Concurrency Architecture & Design Patterns

## 961. How do you architect a pub/sub system in Go?

**Answer:**
Interface: `Subscribe(topic) <-chan Msg` and `Publish(topic, msg)`.
Engine: Map `map[string][]chan Msg` + RWMutex.
Publish: Lock. Iterate channels. Send (Non-blocking select to avoid hanging if one subscriber is slow).
Background: Monitor slow subscribers and close their channels if buffer full.

---

## 962. How do you build a pipeline using goroutines?

**Answer:**
Series of channels.
`Gen() -> chan A`
`Sq(chan A) -> chan B`
`Print(chan B)`
Each stage runs in its own goroutine (or pool of goroutines).
Closing the first channel propagates `nil` (EOF) down the chain, causing all stages to exit gracefully.

---

## 963. What is the fan-in/fan-out pattern in Go?

**Answer:**
**Fan-Out**: 1 Producer -> N Workers.
Loop input channel, send to workers.
**Fan-In**: N Workers -> 1 Result Channel.
`var wg sync.WaitGroup`.
For each worker, `go func() { out <- process(); wg.Done() }`.
`go func() { wg.Wait(); close(out) }`.
The Fan-In merger waits for everyone to finish before closing the output.

---

## 964. How do you limit concurrency using semaphores?

**Answer:**
Buffered Channel `sem := make(chan struct{}, 10)`.
Acquire: `sem <- struct{}{}`.
Release: `<-sem`.
Wrap the helper function:
```go
func Run() {
    sem <- struct{}{}
    defer func() { <-sem }()
    DoExpensiveWork()
}
```
This guarantees max 10 workers at once.

---

## 965. How do you implement a worker pool?

**Answer:**
1.  Job Channel: `jobs := make(chan Job, 100)`.
2.  Workers: Spawn 5 goroutines ranging over `jobs`.
3.  Dispatcher: Pushes items to `jobs`.
4.  Shutdown: Close `jobs`. Workers finish current item and exit loop.

---

## 966. How do you handle retries with backoff in goroutines?

**Answer:**
Loop with sleep.
```go
for i := 0; i < maxRetries; i++ {
    err := do()
    if err == nil { return nil }
    time.Sleep(base * (1 << i)) // Exponential 2, 4, 8s
}
return err
```
Be careful to check `ctx.Done()` in the sleep loop to ensure we can cancel the retry sequence immediately.

---

## 967. What is the circuit breaker pattern in Go?

**Answer:**
Protects failing services.
States: **Closed** (Normal), **Open** (Fail fast), **Half-Open** (Test).
Library: `gobreaker`.
If 5 errors in 10 seconds -> Switch to Open. Return "Circuit Open Error" immediately.
After 30s -> Switch Half-Open. Allow 1 request. If success -> Close. If fail -> Open.

---

## 968. How do you implement message deduplication?

**Answer:**
(See Q 625).
Map or Redis Set.
Key: `msgID`.
Check `StartProcessing`: `if redis.SetNX(msgID, 1) == false { return Duplicate }`.
Set TTL (e.g., 24h) to clean up old keys.

---

## 969. How do you synchronize shared state across goroutines?

**Answer:**
1.  **Mutex**: `mu.Lock(); state++; mu.Unlock()`.
2.  **Channels**: Monitor Goroutine. `state` is local to that goroutine. Others ask for it via `reqChan`. (Share memory by communicating).
3.  **Atomics**: `atomic.AddInt64(&state, 1)`.

---

## 970. How do you detect livelocks in Go?

**Answer:**
Harder than deadlocks.
Livelock = Processes changing state but making no progress (e.g., two people in hallway moving left/right simultaneously).
Detection: Metrics (CPU high, Throughput zero).
Fix: Add **Jitter** (Randomness) to the backoff. `Sleep(1s + rand(100ms))`. This breaks symmetry.

---

## 971. How do you timeout long-running operations?

**Answer:**
`select`.
```go
select {
case res := <-resultCh:
    return res
case <-time.After(5 * time.Second):
    return TimeoutError
}
```
Always ensure the operation being timed out checks `ctx.Done()` so it actually stops burning CPU.

---

## 972. How do you use the actor model in Go?

**Answer:**
Go doesn't have native Actors (like Erlang/Akka).
We simulate it:
AccountActor is a struct with a cmd channel.
It runs a loop `for cmd := range cmds`.
State (`balance`) is private to the loop.
Concurrency safe by design (sequential processing of mailbox).
Libraries like `protoactor-go` provide huge frameworks, but simple channel loops usually suffice.

---

## 973. How do you architect loosely coupled goroutines?

**Answer:**
**Channels** as boundaries.
Goroutine A doesn't know B exists. It just knows "I write to OutputChannel".
Wiring happens in `main`.
`out := A.Start()`
`B.Start(out)`
This allows swapping B with MockB or LoggerB without changing A.

---

## 974. How do you design state machines in Go?

**Answer:**
Struct with `State` func.
```go
type State func(*Machine) State
func Init(m *Machine) State { return Processing }
func Processing(m *Machine) State {
    if m.Done { return Finished }
    return Processing
}
```
Run loop: `for m.curr != nil { m.curr = m.curr(m) }`.
This "Function State" pattern (Rob Pike) is very idiomatic.

---

## 975. How do you throttle a job queue in Go?

**Answer:**
**Token Bucket** consumption.
Worker:
`limiter.Wait()`
`job := <-queue`
`process(job)`
This limits the rate of processing, effectively keeping the system stable under load even if the queue backlog is huge.

---

## 976. How do you monitor goroutine health?

**Answer:**
**Heartbeat Channel**.
Worker sends `heartbeat <- struct{}{}` every second.
Supervisor:
```go
select {
case <-heartbeat:
    // ok
case <-time.After(5*time.Second):
    // Worker stuck. Restart.
}
```

---

## 977. How do you track context propagation in goroutines?

**Answer:**
Pass `context.Context` as first argument to **Every** function.
When spawning a goroutine:
`go func(ctx context.Context) { ... }(ctx)`.
Wait! **Warning**: If parent function returns and cancels context, the child dies.
If child must outlive parent (Fire and Forget), create `context.WithoutCancel(ctx)` (Go 1.21) to decouple lifetime but keep values (TraceIDs).

---

## 978. How do you implement saga pattern in Go services?

**Answer:**
Orchestrator approach.
State Machine in Code.
1.  `err := ServiceA.Do()`
2.  `if err { return fail }`
3.  `err := ServiceB.Do()`
4.  `if err { ServiceA.Undo(); return fail }`
We explicitly define the Compensating Action (`Undo`) for every Forward Action.

---

## 979. How do you chain async jobs with error handling?

**Answer:**
**Pipeline with Error Channel**.
Each stage can send `Result` or `Error`.
Or **ErrGroup**.
`g, ctx := errgroup.WithContext(ctx)`.
`g.Go(func() error { ... })`.
`g.Go(func() error { ... })`.
`err := g.Wait()`.
If any job returns error, context is cancelled, all other jobs stop, and `Wait` returns the first error.

---

## 980. How do you handle non-blocking writes to full channels?

**Answer:**
`select` with `default`.
```go
select {
case ch <- msg:
    // sent
default:
    // channel full. Drop message or Log.
}
```
This is essential for Metrics or Logging systems where we prefer dropping cancel data over blocking the main application flow.
