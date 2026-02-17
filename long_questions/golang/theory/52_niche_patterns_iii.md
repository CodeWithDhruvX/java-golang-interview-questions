# 🧩 Go Theory Questions: 1021–1045 Niche Patterns, Frameworks & Tricky Snippets

## 1021. How do you implement the "Or-Done" channel pattern in Go?

**Answer:**
Wraps a channel read with a `ctx.Done()` check.
```go
func orDone(ctx context.Context, c <-chan int) <-chan int {
    valStream := make(chan int)
    go func() {
        defer close(valStream)
        for {
            select {
            case <-ctx.Done(): return
            case v, ok := <-c:
                if !ok { return }
                select {
                case valStream <- v:
                case <-ctx.Done(): return
                }
            }
        }
    }()
    return valStream
}
```
Ensures consumer loops break instantly on cancellation.

---

## 1022. What is a "Tee-Channel" and how do you implement it?

**Answer:**
Like Unix `tee`. Duplicates one stream into two.
Input `<-chan`. Returns `(<-chan, <-chan)`.
Loop input.
Send to output1 and output2 (using `select` to avoid blocking if one is slow? No, Tee implies strict duplication, usually we block or buffer).
Useful for "Process" AND "Log" simultaneously.

---

## 1023. How do you implement a "Bridge-Channel" to consume a sequence of channels?

**Answer:**
Input: `<-chan (<-chan int)`. (A channel of channels).
Output: `<-chan int`.
The bridge reads a channel from input, ranges over it until close, then reads the next channel.
Flattens a stream of streams into one single stream.

---

## 1024. What is the **Temporal.io** workflow engine and how does it use Go?

**Answer:**
Temporal Server is written in Go.
Concept: **Durable Execution**.
Code:
```go
func Workflow(ctx workflow.Context) error {
    workflow.ExecuteActivity(ctx, ActivityA).Get(&res)
    workflow.Sleep(ctx, 30*time.Hour)
    workflow.ExecuteActivity(ctx, ActivityB).Get(&res)
}
```
If the process dies during Sleep, Temporal restarts it at the exact line, restoring state. (Magic).

---

## 1025. How does **Temporal** ensure determinism in Go workflows?

**Answer:**
Strict rules:
- No `time.Now()` (Use `workflow.Now(ctx)`).
- No `go func` (Use `workflow.Go(ctx, ...)`).
- No `map` iteration (Random order).
- No external API calls (Must use Activities).
The SDK records the "History" of events. On replay, it ensures the code produces the exact same command sequence as the history log.

---

## 1026. What is the **Ent** framework and how does it differ from GORM?

**Answer:**
**Ent** (Facebook). Graph-based ORM.
Code Generation based.
Schema: Go code definition `func (User) Fields()`.
`ent generate ./schema`.
Outcome: Type-safe, fluent API. `client.User.Query().Where(user.Name("a")).All(ctx)`.
Differs from GORM (Reflection-based) by being Compile-Time safe and highly performant (no runtime overhead).

---

## 1027. How do you define graph-based schemas (Edges) in **Ent**?

**Answer:**
```go
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("pets", Pet.Type),
        edge.From("groups", Group.Type).Ref("users"),
    }
}
```
Ent creates the foreign keys or pivot tables automatically. We can query `user.QueryPets().All(ctx)` easily.

---

## 1028. **Tricky Snippet**: What is the output of `fmt.Println(s)` if `s := []int{1,2,3}; append(s[:1], 4)` is called?

**Answer:**
Output: `[1 4 3]`.
`s[:1]` is len 1, cap 3. Address points to start of `s`.
`append` writes `4` at index 1.
Since capacity exists, it mutates the underlying array of `s`.
Original `s` (len 3) sees the change.
Beware of side-effects when appending to sub-slices!

---

## 1029. **Tricky Snippet**: Why is `interface{}(*int(nil)) != nil` true?

**Answer:**
Interface = `(Type, Value)`.
`nil` interface = `(nil, nil)`.
`*int(nil)` = `(*int, nil)`.
Since Type is not nil, the interface is not nil.
Always check `if val == nil` BEFORE converting to interface, or use reflection `val.IsNil()` check.

---

## 1030. **Tricky Snippet**: What happens if you run `for k, v := range m` on the same map multiple times?

**Answer:**
**Random Order**.
Go randomizes map iteration seed to prevent users relying on hash order (which breaks between versions).
Never write tests that expect `{"a":1, "b":2}` to print "a then b".

---

## 1031. **Tricky Snippet**: What happens when you close a nil channel?

**Answer:**
**Panic**.
`panic: close of nil channel`.
Always ensure channel is initialized (`make`) before closing, or use a `Once` guard if multiple closers exist (though only sender should close).

---

## 1032. **Tricky Snippet**: Can you take the address of a map value (`&m["key"]`)? Why or why not?

**Answer:**
**Compilation Error**.
Maps grow/shrink. Moving buckets moves values in memory.
If you held a pointer `p = &m["k"]`, and map grew, `p` would dangle. Go forbids this safery.
Workaround: Store pointers in map `map[string]*Struct`.

---

## 1033. How do you use `golang.org/x/sync/errgroup` for error propagation?

**Answer:**
Replaces `sync.WaitGroup` when errors matter.
`g.Go(func() error { return do() })`.
`if err := g.Wait(); err != nil`.
It returns the **First** error encountered.
Usually paired with `errgroup.WithContext(ctx)` so one failure cancels all other running tasks.

---

## 1034. What is the "Function Options" pattern for constructor configuration?

**Answer:**
(Rob Pike).
`func NewServer(opts ...Option) *Server`.
`type Option func(*Server)`.
`func WithPort(p int) Option { return func(s *Server) { s.port = p } }`.
Usage: `NewServer(WithPort(8080), WithTLS())`.
Elegant, extensible, no breaking changes when adding new config fields.

---

## 1035. How do you use `singleflight` to prevent cache stampedes?

**Answer:**
`import "golang.org/x/sync/singleflight"`.
`var g singleflight.Group`.
```go
val, err, _ := g.Do("key", func() (any, error) {
    return db.Get("key") // Executed ONCE
})
```
If 100 requests come for "key" simultaneously, 1 DB call is made. 100 requests share the result.

---

## 1036. What is `uber-go/automaxprocs` and why is it used in K8s?

**Answer:**
Go sees `runtime.NumCPU()` as Host CPUs (e.g., 64).
Container quota limit: 2 CPUs.
Go scheduler spawns 64 threads. Throttling ensues (CFS).
`import _ "go.uber.org/automaxprocs"`.
It reads `/sys/fs/cgroup`. Sets `GOMAXPROCS=2`.
Drastically improves latency in containerized envs.

---

## 1037. How do you implement "Circuit Breaker" using `sony/gobreaker` or similar?

**Answer:**
(See Q 967).
Wrapper:
```go
cb := gobreaker.NewCircuitBreaker(settings)
res, err := cb.Execute(func() (any, error) {
    return http.Get(...)
})
```
It handles the state transitions and counting automatically.

---

## 1038. How do you use build tags to separate integration tests (`//go:build integration`)?

**Answer:**
File `db_test.go`:
`//go:build integration`
Top of file.
Running `go test ./...` skips it (Fast).
Running `go test -tags integration ./...` runs it (Slow).

---

## 1039. What is the difference between `crypto/rand` and `math/rand/v2` in terms of security?

**Answer:**
- `math/rand/v2`: PCG/ChaCha8. Fast. Good statistics. **Predictable** if state is leaked. NOT for keys.
- `crypto/rand`: OS Entropy. Safe for session IDs, salts, private keys.

---

## 1040. How do you use `go-cmp` for comparing complex structs in tests?

**Answer:**
`reflect.DeepEqual` is strict (unexported fields, timestamps).
`google/go-cmp/cmp`.
`diff := cmp.Diff(want, got)`.
It prints a readable Diff (`- key: "a"`, `+ key: "b"`).
We can use options: `cmpopts.IgnoreFields(User{}, "CreatedAt")`.

---

## 1041. What is "Mutation Testing" and are there tools for it in Go?

**Answer:**
Tool: `gremlins` or `go-mutesting`.
It modifies your code (changes `a < b` to `a > b`).
Runs tests.
If tests **Pass**, the mutant "Survived" (Bad! Test coverage is weak).
If tests **Fail**, the mutant was "Killed" (Good).
Validates quality of tests beyond simple line coverage.

---

## 1042. How do you handle "Dual-Writes" (DB + Message Queue) consistency?

**Answer:**
**Outbox Pattern**.
Single Transaction:
1.  INSERT User.
2.  INSERT OutboxEvent ("UserCreated").
3.  Commit.
Background poller reads Outbox table and publishes to Kafka.
Or CDC (Debezium) reads DB log and publishes.
Never `db.Insert(); kafka.Pub()`. If pub fails, data is inconsistent.

---

## 1043. What is the "Outbox Pattern" and how to implement it in Go?

**Answer:**
(See 1042).
Implementation:
Table `outbox (id, topic, payload, processed)`.
Go Ticker:
`rows := db.Query("SELECT * FROM outbox WHERE !processed")`.
For each: `kafka.Send()`. If success: `db.Update("processed=true")`.
Be idempotent: We might send duplicate Kafka messages if we crash before updating Processed.

---

## 1044. How does Go's "semver" compatibility guarantee work for standard library?

**Answer:**
**Go 1 Promise**.
Code written for Go 1.0 (2012) compiles with Go 1.24 (2025).
Exceptions: Security fixes, Specification errors.
New features are additive.
This stability is why Go is favored for long-term enterprise software.

---

## 1045. How do you use `gdb` or `delve` to debug a running process (attach)?

**Answer:**
Find PID: `pgrep myapp`.
`dlv attach <PID>`.
Commands: `grs` (goroutines), `bt` (backtrace), `b main.go:10` (break).
Essential for debugging "Stuck" processes in production (if you have permissions) without restarting them.
