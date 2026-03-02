# Go Advanced — System Design Pattern Practical Coding

> **Top Product Company Coverage — Part 2 of 2**
> Topics: Circuit Breaker · Rate Limiter · Pub/Sub · Cache Patterns · Middleware Chain · Connection Pool · Retry with Backoff · Event-Driven · Observer · Repository Pattern

---

## Section 1: Concurrency Design Patterns (Q1–Q15)

### 1. Circuit Breaker — Full Implementation
**Q: Implement a basic circuit breaker:**
```go
package main
import (
    "errors"
    "fmt"
    "sync"
    "time"
)

type State int
const (
    Closed   State = iota // normal operation
    Open                  // failing fast
    HalfOpen              // testing recovery
)

type CircuitBreaker struct {
    mu           sync.Mutex
    state        State
    failures     int
    maxFailures  int
    timeout      time.Duration
    lastFailure  time.Time
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{maxFailures: maxFailures, timeout: timeout}
}

var ErrCircuitOpen = errors.New("circuit breaker open")

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    switch cb.state {
    case Open:
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = HalfOpen
        } else {
            cb.mu.Unlock()
            return ErrCircuitOpen
        }
    }
    cb.mu.Unlock()

    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.maxFailures {
            cb.state = Open
        }
        return err
    }
    cb.failures = 0
    cb.state = Closed
    return nil
}

func main() {
    failCount := 0
    cb := NewCircuitBreaker(3, 100*time.Millisecond)

    for i := 0; i < 6; i++ {
        err := cb.Call(func() error {
            failCount++
            if failCount <= 3 {
                return errors.New("service unavailable")
            }
            return nil
        })
        fmt.Printf("call %d: %v\n", i+1, err)
    }
}
```
**A:**
```
call 1: service unavailable
call 2: service unavailable
call 3: service unavailable
call 4: circuit breaker open     ← CB is now OPEN
call 5: circuit breaker open
call 6: circuit breaker open
```
After 3 failures, circuit opens and all calls immediately fail fast without calling the service. After timeout, moves to HalfOpen.

---

### 2. Token Bucket Rate Limiter
**Q: Implement a token bucket rate limiter:**
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type TokenBucket struct {
    mu         sync.Mutex
    tokens     float64
    maxTokens  float64
    refillRate float64 // tokens per second
    lastRefill time.Time
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
    return &TokenBucket{
        tokens:     capacity,
        maxTokens:  capacity,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()

    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens = min(tb.maxTokens, tb.tokens + elapsed*tb.refillRate)
    tb.lastRefill = now

    if tb.tokens >= 1 {
        tb.tokens--
        return true
    }
    return false
}

func min(a, b float64) float64 {
    if a < b { return a }
    return b
}

func main() {
    limiter := NewTokenBucket(3, 10) // 3 burst, 10 tokens/sec

    for i := 0; i < 5; i++ {
        allowed := limiter.Allow()
        fmt.Printf("request %d: allowed=%v\n", i+1, allowed)
    }
}
```
**A:**
```
request 1: allowed=true
request 2: allowed=true
request 3: allowed=true
request 4: allowed=false   ← bucket empty
request 5: allowed=false
```
Token bucket allows burst up to `capacity` then refills at `refillRate`. Better than simple rate limiting because it handles bursts gracefully.

---

### 3. Pub/Sub Event Bus
**Q: Implement a concurrent pub/sub system:**
```go
package main
import (
    "fmt"
    "sync"
)

type EventBus struct {
    mu          sync.RWMutex
    subscribers map[string][]chan interface{}
}

func NewEventBus() *EventBus {
    return &EventBus{subscribers: make(map[string][]chan interface{})}
}

func (eb *EventBus) Subscribe(topic string) <-chan interface{} {
    ch := make(chan interface{}, 10)
    eb.mu.Lock()
    eb.subscribers[topic] = append(eb.subscribers[topic], ch)
    eb.mu.Unlock()
    return ch
}

func (eb *EventBus) Publish(topic string, data interface{}) {
    eb.mu.RLock()
    defer eb.mu.RUnlock()
    for _, ch := range eb.subscribers[topic] {
        select {
        case ch <- data:
        default: // drop if subscriber is slow (non-blocking)
        }
    }
}

func main() {
    bus := NewEventBus()

    ch1 := bus.Subscribe("order.created")
    ch2 := bus.Subscribe("order.created")

    var wg sync.WaitGroup
    for i, ch := range []<-chan interface{}{ch1, ch2} {
        wg.Add(1)
        go func(id int, c <-chan interface{}) {
            defer wg.Done()
            msg := <-c
            fmt.Printf("subscriber %d received: %v\n", id+1, msg)
        }(i, ch)
    }

    bus.Publish("order.created", map[string]interface{}{"id": 42, "item": "book"})
    wg.Wait()
}
```
**A:**
```
subscriber 1 received: map[id:42 item:book]
subscriber 2 received: map[id:42 item:book]
```
Both subscribers receive the event. Non-blocking publish (select+default) prevents a slow subscriber from blocking the publisher.

---

### 4. Connection Pool
**Q: Implement a generic connection pool:**
```go
package main
import (
    "errors"
    "fmt"
    "sync"
)

type Pool[T any] struct {
    mu      sync.Mutex
    conns   []T
    maxSize int
    factory func() (T, error)
    close   func(T)
}

func NewPool[T any](maxSize int, factory func() (T, error), close func(T)) *Pool[T] {
    return &Pool[T]{maxSize: maxSize, factory: factory, close: close}
}

var ErrPoolExhausted = errors.New("pool exhausted")

func (p *Pool[T]) Get() (T, error) {
    p.mu.Lock()
    defer p.mu.Unlock()
    if len(p.conns) > 0 {
        conn := p.conns[len(p.conns)-1]
        p.conns = p.conns[:len(p.conns)-1]
        return conn, nil
    }
    return p.factory()
}

func (p *Pool[T]) Put(conn T) {
    p.mu.Lock()
    defer p.mu.Unlock()
    if len(p.conns) >= p.maxSize {
        p.close(conn) // pool full: discard
        return
    }
    p.conns = append(p.conns, conn)
}

func main() {
    createConn := 0
    pool := NewPool[string](
        3,
        func() (string, error) {
            createConn++
            return fmt.Sprintf("conn-%d", createConn), nil
        },
        func(s string) { fmt.Println("closing:", s) },
    )

    c1, _ := pool.Get()
    c2, _ := pool.Get()
    fmt.Println(c1, c2)
    pool.Put(c1)
    c3, _ := pool.Get() // reuses c1
    fmt.Println(c3)
}
```
**A:**
```
conn-1 conn-2
conn-1
```
Connection pool avoids creating new connections for every request. Reuses existing connections, discards when pool is full.

---

### 5. Middleware Chain (HTTP)
**Q: Implement a composable HTTP middleware chain:**
```go
package main
import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "time"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
    })
}

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-Token") == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    base := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello!")
    })

    handler := Chain(base, Logger, Auth)

    req := httptest.NewRequest("GET", "/api", nil)
    req.Header.Set("X-Token", "secret")
    rec := httptest.NewRecorder()
    handler.ServeHTTP(rec, req)
    fmt.Print(rec.Body.String())
}
```
**A:**
```
GET /api <duration>
Hello!
```
Middleware chain: Logger wraps Auth wraps handler. Each middleware calls `next.ServeHTTP` to pass control. Applied in reverse order so they execute in declaration order.

---

### 6. Retry with Exponential Backoff + Jitter
**Q: Why is jitter important?**
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func retry(attempts int, base time.Duration, fn func() error) error {
    var err error
    for i := 0; i < attempts; i++ {
        if err = fn(); err == nil {
            return nil
        }
        // Exponential backoff with full jitter
        // Jitter prevents "thundering herd" — all retries hitting at once
        maxBackoff := float64(base) * math.Pow(2, float64(i))
        jitter := rand.Float64() * maxBackoff
        sleep := time.Duration(jitter)
        fmt.Printf("attempt %d failed: %v, retrying in %v\n", i+1, err, sleep.Round(time.Millisecond))
        time.Sleep(sleep)
    }
    return fmt.Errorf("all %d attempts failed: %w", attempts, err)
}

func main() {
    calls := 0
    err := retry(4, 100*time.Millisecond, func() error {
        calls++
        if calls < 3 { return fmt.Errorf("temporary error") }
        return nil
    })
    fmt.Println("result:", err)
}
```
**A:**
```
attempt 1 failed: temporary error, retrying in Xms
attempt 2 failed: temporary error, retrying in Xms
result: <nil>
```
Jitter spreads retries across time — critical when many clients retry simultaneously (e.g., after a service restart). Without jitter, all clients retry at the same intervals, creating spike load.

---

### 7. Observer Pattern
**Q: Implement with Go idioms:**
```go
package main
import "fmt"

type Event struct {
    Type string
    Data interface{}
}

type Handler func(Event)

type EventEmitter struct {
    handlers map[string][]Handler
}

func NewEmitter() *EventEmitter {
    return &EventEmitter{handlers: make(map[string][]Handler)}
}

func (e *EventEmitter) On(eventType string, h Handler) {
    e.handlers[eventType] = append(e.handlers[eventType], h)
}

func (e *EventEmitter) Emit(ev Event) {
    for _, h := range e.handlers[ev.Type] {
        h(ev)
    }
}

func main() {
    emitter := NewEmitter()

    emitter.On("user.created", func(ev Event) {
        fmt.Println("email service: send welcome email to", ev.Data)
    })
    emitter.On("user.created", func(ev Event) {
        fmt.Println("analytics: log user creation for", ev.Data)
    })

    emitter.Emit(Event{Type: "user.created", Data: "alice@example.com"})
}
```
**A:**
```
email service: send welcome email to alice@example.com
analytics: log user creation for alice@example.com
```

---

### 8. Repository Pattern
**Q: Implement repository with interface:**
```go
package main
import (
    "errors"
    "fmt"
)

type User struct {
    ID   int
    Name string
}

type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(u *User) error
    Delete(id int) error
}

// In-memory implementation (for testing)
type InMemoryUserRepo struct {
    store map[int]*User
    nextID int
}

func NewInMemoryRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{store: make(map[int]*User)}
}

func (r *InMemoryUserRepo) FindByID(id int) (*User, error) {
    u, ok := r.store[id]
    if !ok {
        return nil, fmt.Errorf("user %d: %w", id, errors.New("not found"))
    }
    return u, nil
}

func (r *InMemoryUserRepo) Save(u *User) error {
    if u.ID == 0 {
        r.nextID++
        u.ID = r.nextID
    }
    r.store[u.ID] = u
    return nil
}

func (r *InMemoryUserRepo) Delete(id int) error {
    delete(r.store, id)
    return nil
}

// UserService depends on interface, not implementation
type UserService struct{ repo UserRepository }

func (s *UserService) CreateUser(name string) (*User, error) {
    u := &User{Name: name}
    return u, s.repo.Save(u)
}

func main() {
    repo := NewInMemoryRepo()
    svc := &UserService{repo: repo}

    u, _ := svc.CreateUser("Alice")
    fmt.Printf("created user: id=%d name=%s\n", u.ID, u.Name)

    found, _ := repo.FindByID(u.ID)
    fmt.Println("found:", found.Name)
}
```
**A:**
```
created user: id=1 name=Alice
found: Alice
```

---

### 9. Cache-Aside Pattern with TTL
**Q: Implement TTL cache:**
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type entry[V any] struct {
    value   V
    expires time.Time
}

type TTLCache[K comparable, V any] struct {
    mu    sync.RWMutex
    store map[K]entry[V]
}

func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
    return &TTLCache[K, V]{store: make(map[K]entry[V])}
}

func (c *TTLCache[K, V]) Set(key K, val V, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.store[key] = entry[V]{value: val, expires: time.Now().Add(ttl)}
}

func (c *TTLCache[K, V]) Get(key K) (V, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    e, ok := c.store[key]
    if !ok || time.Now().After(e.expires) {
        var zero V
        return zero, false
    }
    return e.value, true
}

func main() {
    cache := NewTTLCache[string, string]()
    cache.Set("key", "value", 100*time.Millisecond)

    v, ok := cache.Get("key")
    fmt.Println(v, ok) // value true

    time.Sleep(150 * time.Millisecond)
    v2, ok2 := cache.Get("key")
    fmt.Println(v2, ok2) // "" false (expired)
}
```
**A:**
```
value true
 false
```

---

### 10. Read-Through Cache Pattern
**Q: What is the read-through pattern?**
```go
type UserCache struct {
    mu    sync.RWMutex
    cache map[int]*User
    db    UserRepository
}

func (c *UserCache) GetUser(id int) (*User, error) {
    // 1. Check cache first
    c.mu.RLock()
    if u, ok := c.cache[id]; ok {
        c.mu.RUnlock()
        return u, nil // cache hit
    }
    c.mu.RUnlock()

    // 2. Fetch from DB (cache miss)
    u, err := c.db.FindByID(id)
    if err != nil {
        return nil, err
    }

    // 3. Populate cache
    c.mu.Lock()
    c.cache[id] = u
    c.mu.Unlock()

    return u, nil
}
```
**A:** Read-through: check cache → on miss, fetch from DB → populate cache → return. Write-through (not shown): write to DB AND cache simultaneously. Cache-aside: application manages both operations manually. Read-through hides complexity from callers.

---

### 11. Graceful Shutdown
**Q: What is the correct shutdown sequence?**
```go
package main
import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    srv := &http.Server{Addr: ":8080"}

    // Start server in goroutine
    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            fmt.Println("server error:", err)
        }
    }()
    fmt.Println("server started")

    // Wait for OS signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    fmt.Println("shutdown signal received")

    // Graceful shutdown: wait up to 30s for inflight requests
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        fmt.Println("forced shutdown:", err)
    }
    fmt.Println("server stopped cleanly")
}
```
**A:** `srv.Shutdown(ctx)` stops accepting new connections and waits for in-flight requests to complete within the deadline. This is the production mandatory pattern — prevents dropping active requests during deploys.

---

### 12. Health Check Endpoint
**Q: What is the standard health check pattern?**
```go
package main
import (
    "encoding/json"
    "net/http"
    "time"
)

type HealthStatus struct {
    Status    string            `json:"status"`
    Timestamp time.Time         `json:"timestamp"`
    Checks    map[string]string `json:"checks"`
}

func healthHandler(db DBPinger, cache CachePinger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        checks := map[string]string{}
        status := "healthy"

        if err := db.Ping(); err != nil {
            checks["database"] = "unhealthy: " + err.Error()
            status = "unhealthy"
        } else {
            checks["database"] = "ok"
        }

        if err := cache.Ping(); err != nil {
            checks["cache"] = "unhealthy: " + err.Error()
            status = "degraded"
        } else {
            checks["cache"] = "ok"
        }

        code := http.StatusOK
        if status == "unhealthy" { code = http.StatusServiceUnavailable }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(code)
        json.NewEncoder(w).Encode(HealthStatus{
            Status:    status,
            Timestamp: time.Now(),
            Checks:    checks,
        })
    }
}
```
**A:** `/health` returns 200 if healthy, 503 if unhealthy. Kubernetes liveness/readiness probes call this endpoint. Separate `/ready` (can serve traffic) from `/live` (process alive) checks.

---

### 13. Bulkhead Pattern — Resource Isolation
**Q: What problem does bulkhead solve?**
```go
package main
import (
    "errors"
    "fmt"
    "sync"
)

type Bulkhead struct {
    sem chan struct{}
}

func NewBulkhead(maxConcurrent int) *Bulkhead {
    return &Bulkhead{sem: make(chan struct{}, maxConcurrent)}
}

var ErrBulkheadFull = errors.New("bulkhead at capacity")

func (b *Bulkhead) Execute(fn func() error) error {
    select {
    case b.sem <- struct{}{}: // acquire slot
        defer func() { <-b.sem }()
        return fn()
    default:
        return ErrBulkheadFull // reject immediately
    }
}

func main() {
    // Separate bulkheads for different services — failures don't cascade
    dbBulkhead    := NewBulkhead(10)  // max 10 concurrent DB calls
    cacheBulkhead := NewBulkhead(50)  // max 50 concurrent cache calls

    var wg sync.WaitGroup
    for i := 0; i < 15; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            err := dbBulkhead.Execute(func() error {
                fmt.Printf("db call %d\n", id)
                return nil
            })
            if err != nil {
                fmt.Printf("db call %d rejected: %v\n", id, err)
            }
        }(i)
    }
    _ = cacheBulkhead
    wg.Wait()
}
```
**A:** Bulkhead isolates failure domains. If DB calls are slow (filling the DB bulkhead), cache calls are unaffected. Without bulkhead, slow DB response drains all goroutines, crashing the entire service.

---

### 14. Event Sourcing Pattern
**Q: Implement a simple event store:**
```go
package main
import (
    "fmt"
    "time"
)

type Event struct {
    ID        int
    Type      string
    Data      interface{}
    Timestamp time.Time
}

type EventStore struct {
    events []Event
    nextID int
}

func (es *EventStore) Append(eventType string, data interface{}) Event {
    es.nextID++
    ev := Event{ID: es.nextID, Type: eventType, Data: data, Timestamp: time.Now()}
    es.events = append(es.events, ev)
    return ev
}

func (es *EventStore) Replay() []Event { return es.events }

type BankAccount struct{ balance int }

func (a *BankAccount) Apply(ev Event) {
    switch ev.Type {
    case "deposited":
        a.balance += ev.Data.(int)
    case "withdrawn":
        a.balance -= ev.Data.(int)
    }
}

func main() {
    store := &EventStore{}
    store.Append("deposited", 100)
    store.Append("deposited", 50)
    store.Append("withdrawn", 30)

    // Rebuild state from events
    acc := &BankAccount{}
    for _, ev := range store.Replay() {
        acc.Apply(ev)
    }
    fmt.Println("balance:", acc.balance) // 120
}
```
**A:** `balance: 120`. Event sourcing stores state changes as immutable events. Current state = replay of all events. Enables: audit log, time-travel debugging, CQRS, and event-driven architectures.

---

### 15. Outbox Pattern (Transactional Messaging)
**Q: What problem does the outbox pattern solve?**
```go
// Problem: Writing to DB and sending event must be atomic
// BAD: two-phase approach can lose event on crash
func badCreate(db DB, bus EventBus, user User) error {
    db.Save(user)        // step 1
    bus.Publish(event)   // step 2 — crashes here = event lost!
    return nil
}

// GOOD: Outbox pattern
// 1. Write user + event to DB in ONE transaction
// 2. Background process polls outbox and publishes to message broker
// 3. Mark event as published after successful delivery

func goodCreate(tx *sql.Tx, user User) error {
    tx.Exec("INSERT INTO users ...", user)
    tx.Exec("INSERT INTO outbox (type, payload) VALUES (?, ?)",
        "user.created", marshal(user))
    return tx.Commit() // atomic: both or neither
}

// Separate publisher goroutine:
func outboxPublisher(db DB, bus EventBus) {
    for {
        events := db.Query("SELECT * FROM outbox WHERE published=false LIMIT 100")
        for _, ev := range events {
            bus.Publish(ev)
            db.Exec("UPDATE outbox SET published=true WHERE id=?", ev.ID)
        }
        time.Sleep(100 * time.Millisecond)
    }
}
```
**A:** Solves "dual write" problem — database and message broker can't be updated in a single transaction. Outbox table is in the same DB transaction; a separate process publishes durably with at-least-once delivery.

---

## Section 2: Distributed System Patterns in Go (Q16–Q25)

### 16. gRPC Server with Context Propagation
**Q: What is the correct context pattern in gRPC?**
```go
import (
    "context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// Server implementation
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    // ctx is already cancelled if client disconnects or deadline exceeded
    user, err := s.repo.FindByID(ctx, int(req.Id))
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            return nil, status.Errorf(codes.NotFound, "user %d not found", req.Id)
        }
        return nil, status.Errorf(codes.Internal, "internal error: %v", err)
    }
    return &pb.User{Id: int32(user.ID), Name: user.Name}, nil
}
```
**A:** Always pass `ctx` to all downstream calls. Use `status.Errorf` with gRPC status codes (not `fmt.Errorf`) — gRPC codes map to HTTP status codes automatically. Clients receive structured errors, not strings.

---

### 17. Distributed Tracing — OpenTelemetry Pattern
**Q: What is the standard tracing instrument pattern?**
```go
import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("user-service")

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    ctx, span := tracer.Start(ctx, "UserService.GetUser")
    defer span.End()

    span.SetAttributes(attribute.Int("user.id", id))

    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        span.RecordError(err)
        return nil, err
    }

    span.SetAttributes(attribute.String("user.name", user.Name))
    return user, nil
}
```
**A:** OpenTelemetry is the standard. `tracer.Start(ctx, "span-name")` creates a child span under the incoming trace. `defer span.End()` ensures span is always recorded. Spans flow: HTTP → service → DB as a trace tree.

---

### 18. Sync Map vs Sharded Map for High Concurrency
**Q: When does a sharded map outperform sync.Map?**
```go
package main
import "sync"

// For high write throughput: shard the map to reduce contention
const shards = 32

type ShardedMap struct {
    shards [shards]struct {
        sync.RWMutex
        m map[string]interface{}
    }
}

func NewShardedMap() *ShardedMap {
    sm := &ShardedMap{}
    for i := range sm.shards {
        sm.shards[i].m = make(map[string]interface{})
    }
    return sm
}

func (sm *ShardedMap) shard(key string) *struct {
    sync.RWMutex
    m map[string]interface{}
} {
    var hash uint32
    for _, c := range key { hash = hash*31 + uint32(c) }
    return &sm.shards[hash%shards]
}

func (sm *ShardedMap) Set(key string, val interface{}) {
    s := sm.shard(key)
    s.Lock()
    defer s.Unlock()
    s.m[key] = val
}

func (sm *ShardedMap) Get(key string) (interface{}, bool) {
    s := sm.shard(key)
    s.RLock()
    defer s.RUnlock()
    v, ok := s.m[key]
    return v, ok
}
```
**A:** `sync.Map` is optimized for read-heavy, write-once workloads. For balanced read/write across many keys, a sharded map (N independent maps with N mutexes) reduces lock contention by factor N.

---

### 19. gRPC Interceptors (Middleware)
**Q: What is the gRPC interceptor pattern?**
```go
import (
    "context"
    "time"
    "google.golang.org/grpc"
    "go.uber.org/zap"
)

// Unary interceptor for logging and timing
func loggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        start := time.Now()
        resp, err := handler(ctx, req)
        duration := time.Since(start)

        if err != nil {
            logger.Error("gRPC call failed",
                zap.String("method", info.FullMethod),
                zap.Duration("duration", duration),
                zap.Error(err))
        } else {
            logger.Info("gRPC call succeeded",
                zap.String("method", info.FullMethod),
                zap.Duration("duration", duration))
        }
        return resp, err
    }
}

// Register:
// grpc.NewServer(grpc.ChainUnaryInterceptor(loggingInterceptor(logger), authInterceptor()))
```
**A:** gRPC interceptors are Go's middleware pattern for RPC calls. `ChainUnaryInterceptor` composes multiple interceptors. Used for: logging, auth token validation, rate limiting, panic recovery, distributed tracing.

---

### 20. Functional Options Pattern
**Q: What is the output and why is this pattern used?**
```go
package main
import (
    "fmt"
    "time"
)

type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func WithHost(h string) Option     { return func(s *Server) { s.host = h } }
func WithPort(p int) Option        { return func(s *Server) { s.port = p } }
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }
func WithMaxConn(n int) Option     { return func(s *Server) { s.maxConn = n } }

func NewServer(opts ...Option) *Server {
    s := &Server{host: "localhost", port: 8080, timeout: 30*time.Second, maxConn: 100}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

func main() {
    s := NewServer(
        WithHost("0.0.0.0"),
        WithPort(9090),
        WithMaxConn(500),
        // timeout uses default
    )
    fmt.Printf("host=%s port=%d timeout=%v maxConn=%d\n",
        s.host, s.port, s.timeout, s.maxConn)
}
```
**A:** `host=0.0.0.0 port=9090 timeout=30s maxConn=500`. Functional options provide: (1) backward compatibility — add new options without breaking existing callers, (2) named parameters, (3) defaults, (4) self-documenting API. Used in `grpc.Dial`, `http.Server`, and many SDKs.

---

## Section 3: Memory & Performance Patterns (Q21–Q30)

### 21. Zero-Copy I/O with io.WriterTo
**Q: When does this avoid allocations?**
```go
package main
import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

// strings.Reader implements io.WriterTo — zero-copy path
func transfer(dst io.Writer, src io.WriterTo) (int64, error) {
    return src.WriteTo(dst)
}

func main() {
    src := strings.NewReader("hello, world")
    dst := &bytes.Buffer{}
    n, err := transfer(dst, src)
    fmt.Println(n, err, dst.String())
}
```
**A:** `12 <nil> hello, world`. When both src and dst implement `io.WriterTo`/`io.ReaderFrom`, `io.Copy` uses the zero-copy fast path — no intermediate buffer allocation. Crucial for high-throughput data pipelines.

---

### 22. Struct Field Ordering — Memory Alignment
**Q: Which struct uses less memory?**
```go
package main
import (
    "fmt"
    "unsafe"
)

// Poorly aligned — padding inserted between fields
type Bad struct {
    a bool    // 1 byte + 7 bytes padding
    b float64 // 8 bytes
    c bool    // 1 byte + 7 bytes padding
}

// Well aligned — fields ordered largest to smallest
type Good struct {
    b float64 // 8 bytes
    a bool    // 1 byte
    c bool    // 1 byte + 6 bytes padding
}

func main() {
    fmt.Println("Bad:", unsafe.Sizeof(Bad{}))   // 24 bytes
    fmt.Println("Good:", unsafe.Sizeof(Good{})) // 16 bytes
}
```
**A:**
```
Bad: 24
Good: 16
```
Order struct fields from largest to smallest type. `Bad` wastes 14 bytes to padding; `Good` wastes only 6. Matters significantly for large slices of structs.

---

### 23. Avoiding toString in Hot Path
**Q: What is the allocation difference?**
```go
package main
import (
    "fmt"
    "testing"
)

// BenchmarkFmtSprintf: 2 allocs/op (format string + result)
func BenchmarkSprintf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := fmt.Sprintf("user-%d", i)
        _ = s
    }
}

// BenchmarkStrconv: 1 alloc/op — faster
func BenchmarkStrconv(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := "user-" + strconv.Itoa(i)
        _ = s
    }
}

// BenchmarkBuilder: 1 alloc/op — reusable
func BenchmarkBuilder(b *testing.B) {
    var sb strings.Builder
    for i := 0; i < b.N; i++ {
        sb.Reset()
        sb.WriteString("user-")
        sb.WriteString(strconv.Itoa(i))
        _ = sb.String()
    }
}
```
**A:** `fmt.Sprintf` is convenient but allocates more. In hot paths (millions of calls/sec), `strings.Builder` or `strconv` functions reduce allocations by 50-75%. Profile first with `-benchmem` before optimizing.

---

### 24. Map Preallocation
**Q: Why does preallocation matter?**
```go
package main
import (
    "fmt"
    "testing"
)

// Without hint: multiple rehash + copy operations as map grows
func BenchmarkMapNoHint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        m := make(map[int]int) // no size hint
        for j := 0; j < 10000; j++ {
            m[j] = j
        }
    }
}

// With hint: allocate buckets upfront, no rehash needed
func BenchmarkMapWithHint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        m := make(map[int]int, 10000) // size hint
        for j := 0; j < 10000; j++ {
            m[j] = j
        }
    }
}

func main() {
    fmt.Println("see benchmark results with: go test -bench=. -benchmem")
}
```
**A:** With hint: ~2x faster, ~50% fewer allocations. Map hint pre-sizes the backing hash table. Same applies to slices: `make([]T, 0, cap)` avoids repeated `append` reallocations.

---

### 25. interface{} Boxing Allocation
**Q: Does this allocate?**
```go
package main
import "testing"

func noEscape(v int) int { return v + 1 }

func withInterface(v interface{}) interface{} { return v } // boxes v onto heap

func BenchmarkNoEscape(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = noEscape(42)
    }
}

func BenchmarkWithInterface(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = withInterface(42) // allocates: int 42 boxed into interface{}
    }
}
```
**A:** `withInterface(42)` allocates because the integer must be heap-allocated to be stored in an `interface{}` (boxing). `noEscape(42)` is stack-only. This is why generics (Go 1.18+) are faster than `interface{}` for type-generic code.

---

### 26. Goroutine Pool via sync.Pool (Advanced)
**Q: Why is this pattern used in production HTTP servers?**
```go
package main
import (
    "bytes"
    "encoding/json"
    "net/http"
    "sync"
)

var bufPool = sync.Pool{
    New: func() interface{} { return new(bytes.Buffer) },
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
    buf := bufPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufPool.Put(buf)

    data := map[string]string{"status": "ok"}
    if err := json.NewEncoder(buf).Encode(data); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(buf.Bytes())
}
```
**A:** At 100K req/s, allocating a `bytes.Buffer` per request = 100K GC objects/sec. `sync.Pool` reuses buffers across requests, reducing GC pressure dramatically. The Go standard library's `encoding/json` uses this pattern internally.

---

### 27. CPU Cache Line Friendly Struct
**Q: What is false sharing and how to avoid it?**
```go
package main
import (
    "sync/atomic"
    "runtime"
)

// BAD: false sharing — counters on same cache line (64 bytes)
// When goroutine on CPU0 writes counter[0], CPU1's cache of counter[1] is invalidated
type BadCounters struct {
    counters [4]int64 // all on same or adjacent cache lines
}

// GOOD: pad each counter to its own cache line
type PaddedCounter struct {
    value   int64
    _       [56]byte // padding to fill cache line (64 - 8 bytes)
}

type GoodCounters struct {
    counters [4]PaddedCounter
}

func main() {
    good := &GoodCounters{}
    // Each goroutine writes to its own cache line — no false sharing
    for i := 0; i < runtime.NumCPU(); i++ {
        go func(idx int) {
            atomic.AddInt64(&good.counters[idx].value, 1)
        }(i % 4)
    }
}
```
**A:** False sharing = multiple goroutines on different CPUs write to different variables sharing the same 64-byte cache line. Causes cache line ping-pong between CPUs, degrading performance 10-100x. Padding forces each hot variable to its own cache line.

---

### 28. Producer-Consumer with Backpressure
**Q: What is the output and what does backpressure prevent?**
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Bounded buffer creates backpressure
    buffer := make(chan int, 5) // max 5 items in flight

    var wg sync.WaitGroup

    // Producer: fast
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(buffer)
        for i := 0; i < 10; i++ {
            buffer <- i // blocks when buffer full
            fmt.Printf("produced: %d (buffer len: %d)\n", i, len(buffer))
        }
    }()

    // Consumer: slow
    wg.Add(1)
    go func() {
        defer wg.Done()
        for v := range buffer {
            time.Sleep(20 * time.Millisecond) // slow processing
            fmt.Printf("consumed: %d\n", v)
        }
    }()

    wg.Wait()
}
```
**A:** Producer blocks when buffer is full (backpressure applied). Without bounded buffer, fast producers overwhelm slow consumers, causing unbounded memory growth. Buffered channels are Go's built-in backpressure mechanism.

---

### 29. Write-Behind Cache (Async Write)
**Q: What is the write-behind pattern?**
```go
type WriteCache struct {
    mu      sync.RWMutex
    cache   map[string]string
    writeCh chan writeOp
}

type writeOp struct{ key, value string }

func NewWriteCache(db Database) *WriteCache {
    wc := &WriteCache{
        cache:   make(map[string]string),
        writeCh: make(chan writeOp, 1000),
    }
    go wc.flushLoop(db)
    return wc
}

func (wc *WriteCache) Set(key, value string) {
    wc.mu.Lock()
    wc.cache[key] = value // write to cache synchronously (fast)
    wc.mu.Unlock()
    wc.writeCh <- writeOp{key, value} // async persist
}

func (wc *WriteCache) flushLoop(db Database) {
    for op := range wc.writeCh {
        db.Set(op.key, op.value) // slow DB write happens asynchronously
    }
}
```
**A:** Write-behind: return to caller immediately after writing to cache; persist to DB asynchronously. Improves write latency at cost of durability risk (data loss if process dies before flush). Used where throughput > durability (analytics, logging, leaderboards).

---

### 30. Singleflight for Database Stampede Prevention
**Q: What is the output and what does this prevent?**
```go
package main
import (
    "fmt"
    "sync"
    "time"
    "golang.org/x/sync/singleflight"
)

var (
    sg    singleflight.Group
    dbHits int
)

func getFromDB(key string) (string, error) {
    dbHits++
    time.Sleep(50 * time.Millisecond) // simulate slow DB
    return "value-for-" + key, nil
}

func getUser(key string) (string, error) {
    v, err, _ := sg.Do(key, func() (interface{}, error) {
        return getFromDB(key)
    })
    return v.(string), err
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            getUser("user:42")
        }()
    }
    wg.Wait()
    fmt.Println("DB hit count:", dbHits) // 1, not 10
}
```
**A:** `DB hit count: 1`. 10 goroutines requested the same key simultaneously — only ONE DB call was made. All 10 goroutines receive the same result. Prevents cache stampede under high load.

---

*End of Part 2 — System Design Pattern Practical Coding (30 questions)*

---

## Complete Advanced Coverage Summary

| File | Topics | Q# |
|---|---|---|
| `go_advanced_runtime_testing.md` | GMP model, work stealing, preemption, goroutine states, GC tri-color, GOGC, GOMEMLIMIT, escape analysis, sync.Pool, pprof, `t.Parallel`, fuzz testing, TestMain, httptest, benchstat, testcontainers | 45 |
| `go_advanced_system_design_patterns.md` | Circuit Breaker, Token Bucket Rate Limiter, Pub/Sub, Connection Pool, Middleware Chain, Retry+Jitter, Observer, Repository, TTL Cache, Read-Through Cache, Graceful Shutdown, Health Check, Bulkhead, Event Sourcing, Outbox, gRPC interceptors, Functional Options, False Sharing, Backpressure, Singleflight | 30 |
| **Total** | | **75** |

---

## Full Interview Readiness Summary

| Company Level | Coverage | Files |
|---|---|---|
| Service (TCS, Infosys) | ✅ 100% | `go_basics_fundamentals_snippets.md`, `go_basics_indepth_snippets.md`, `go_service_company_coverage.md` |
| Mid-Tier Product (Razorpay, Juspay) | ✅ 100% | + `go_intermediate_concurrency.md`, `go_intermediate_context_interfaces.md` |
| Top Product (Google, Uber) | ✅ 100% | + `go_advanced_runtime_testing.md`, `go_advanced_system_design_patterns.md` |
