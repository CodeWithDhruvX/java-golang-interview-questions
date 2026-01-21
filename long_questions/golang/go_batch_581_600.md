## 🧠 Design Patterns, Architecture & Real-World Scenarios (Questions 581-600)

### Question 581: How do you implement the Factory pattern in Go?

**Answer:**
A simple function that returns an interface.

```go
type Store interface { Save(string) }

type DiskStore struct{}
func (d *DiskStore) Save(s string) {}

type MemStore struct{}
func (m *MemStore) Save(s string) {}

func NewStore(t string) Store {
    if t == "disk" { return &DiskStore{} }
    return &MemStore{}
}
```

---

### Question 582: How do you use the Strategy pattern in Go?

**Answer:**
Define a family of algorithms as function types or interfaces and swap them at runtime.

```go
type PaymentStrategy func(amount int) error

func PayWithCard(amount int) error { ... }
func PayWithPayPal(amount int) error { ... }

func Checkout(amount int, strategy PaymentStrategy) {
    strategy(amount)
}
```

---

### Question 583: What is the Singleton pattern and how is it safely used in Go?

**Answer:**
Use `sync.Once` to ensure lazy initialization happens exactly once, even concurrently.

```go
var instance *Database
var once sync.Once

func GetDatabase() *Database {
    once.Do(func() {
        instance = &Database{} // Connection logic
    })
    return instance
}
```

---

### Question 584: How do you write a middleware chain in Go?

**Answer:**
Chain handlers by wrapping them.

```go
func Chain(f http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}
```
Or use a library like **Alice** or **Chi**'s middleware stack.

---

### Question 585: How do you use interfaces to decouple layers?

**Answer:**
The **Business Logic** layer should define the interfaces it needs (Dependency Inversion), and the **Storage** layer should implement them.
`package business` defines `UserRepository interface`.
`package postgres` implements `UserRepository`.
This allows `business` to be tested with mocks and not depend on SQL.

---

### Question 586: How do you implement the Observer pattern using channels?

**Answer:**
Store a list of channels (subscribers). When an event occurs, iterate and send to them.

```go
type Broker struct {
    subscribers []chan string
}

func (b *Broker) Subscribe() chan string {
    ch := make(chan string)
    b.subscribers = append(b.subscribers, ch)
    return ch
}

func (b *Broker) Publish(msg string) {
    for _, ch := range b.subscribers {
        go func(c chan string) { c <- msg }(ch)
    }
}
```

---

### Question 587: What is the repository pattern and when do you use it?

**Answer:**
It abstracts data access.
**Interface:** `GetUser(id int) (*User, error)`
**Impl:** `SqlUserRepository` (using GORM/SQL).
**Use it:** To separate domain logic from database details, allowing easy swapping (Postgres -> Mongo) and unit testing.

---

### Question 588: How would you create a CQRS architecture in Go?

**Answer:**
**Command Query Responsibility Segregation.**
Split into two models:
1.  **Command (Write):** Methods that mutate state (`CreateOrder`). Often async, uses Event Sourcing.
2.  **Query (Read):** Methods that return data (`GetOrder`). optimized for reads (Materialized Views).
In Go, simple implementation: Separation of `OrderWriter` and `OrderReader` interfaces.

---

### Question 589: How do you design a plug-in architecture in Go?

**Answer:**
1.  **Go Plugins (`plugin` package):** Load `.so` files at runtime (Linux/Mac only, tricky versioning).
2.  **RPC/Hashicorp Plugin:** Run plugins as separate processes (binary) and communicate via gRPC/net/rpc over localhost. (Used by Terraform). Safer and more robust.

---

### Question 590: What is a “clean architecture” in Go projects?

**Answer:**
Standard layout (Uncle Bob):
- **Entities (Domain):** Core structs, no deps.
- **Usecases (Service):** Business rules, depends on Entities.
- **Adapters (Controller/Repo):** HTTP handlers, SQL implementations, depends on Usecases.
- **Drivers (Main):** Wires everything up (Router, DB connection).
Ensures deps point **inwards**.

---

### Question 591: How do you structure a multi-module Go project?

**Answer:**
Use Go Workspaces (`go.work`).
```
/project
  go.work
  /api (go.mod)
  /libs (go.mod)
  /services
    /payment (go.mod)
```
Allows developing multiple modules simultaneously without publishing to Git tags constantly.

---

### Question 592: How do you decouple business logic from transport layers?

**Answer:**
Do **not** use `gin.Context` or `http.ResponseWriter` in your Service methods.
**Bad:** `func (s *Service) Create(c *gin.Context)`
**Good:** `func (s *Service) Create(ctx context.Context, u User) error`
The HTTP Handler parses JSON -> calls Service -> returns JSON.

---

### Question 593: How would you implement retryable jobs in Go?

**Answer:**
Use a Queue (Redis/RabbitMQ) + Worker.
1.  Worker pulls job.
2.  Executes.
3.  If fail: Check `retry_count`. If < Max, publish back to queue with a delay (Exponential Backoff). If >= Max, move to **Dead Letter Queue (DLQ)**.

---

### Question 594: How would you design a billing system in Go?

**Answer:**
- **Concurrency:** Use `sync.Mutex` or DB Record Locking (`SELECT FOR UPDATE`) to prevent double-spending.
- **Precision:** NEVER use `float64`. Use `int64` (cents) or `shopspring/decimal`.
- **Idempotency:** Crucial for API.
- **Audit:** Append-only ledger table (`TransactionHistory`) for every balance change.

---

### Question 595: How would you scale a notification system written in Go?

**Answer:**
- **Fan-out:** 1 Event ("New Post") -> Produce 1000 messages ("Notify Follower X") to Kafka.
- **Workers:** 50 Go pods consuming Kafka and sending FCM/Email/SMS.
- **Rate Limiting:** Throttle sends per user/provider to avoid bans.

---

### Question 596: How do you build a real-time leaderboard in Go?

**Answer:**
Do not use SQL `ORDER BY`. Use **Redis Sorted Sets (ZSET)**.
- `ZADD leaderboard score user_id`
- `ZREVRANGE leaderboard 0 10 WITHSCORES` (Top 10).
Go service acts as a wrapper around Redis commands.

---

### Question 597: How would you implement transactional emails in Go?

**Answer:**
Reliability is key.
1.  **DB Transaction:** Write "Order Created" AND "EmailJob" (status=pending) to DB atomically (Outbox Pattern).
2.  **Worker:** Polls "EmailJob", sends via SendGrid/AWS SES.
3.  **On Success:** Delete/Update Job.
Ensures email is sent if and only if Order is committed.

---

### Question 598: How do you model money and currencies in Go?

**Answer:**
Struct approach.
```go
type Money struct {
    Amount   int64 // in minor units (cents)
    Currency string // "USD", "EUR"
}
```
Always check currency match before addition/subtraction.

---

### Question 599: How do you do dependency injection in Go?

**Answer:**
- **Manual (Idiomatic):** Pass dependencies to constructors.
    ```go
    func NewService(db *sql.DB, logger Logger) *Service { ... }
    ```
- **Libraries:** `google/wire` (Compile time code-gen) or `uber-go/dig` (Reflection based). Go prefers manual/wire (explicit) over "Magic" containers.

---

### Question 600: How do you create a rule engine in Go?

**Answer:**
1.  **Interface:** `type Rule interface { Evaluate(Context) bool }`.
2.  **Composite:** `type AndRule struct { rules []Rule }`.
3.  **DSL (Advanced):** Use `antonmedv/expr` or `google/cel-go` (Common Expression Language) to parse string rules (`user.age > 18 && user.premium`) and evaluate them safely at runtime.

---
