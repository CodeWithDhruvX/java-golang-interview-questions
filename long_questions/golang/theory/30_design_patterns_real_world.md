# 🟢 Go Theory Questions: 581–600 Design Patterns, Architecture & Real-World Scenarios

## 581. How do you implement the Factory pattern in Go?

**Answer:**
We use a simple **Function** that returns an interface.

`func NewStore(type string) (Store, error)`
Inside:
```go
switch type {
case "postgres": return newPostgresStore(), nil
case "memory": return newMemoryStore(), nil
}
```
Go rarely uses "Factory Classes" (AbstractFactory). A simple function/constructor is the idiomatic way to encapsulate creation logic and return a concrete type behind an abstraction.

---

## 582. How do you use the Strategy pattern in Go?

**Answer:**
Strategy is about swapping algorithms at runtime.
We define a function type or interface: `type EvictionAlgo func(key string)`.

The Context struct holds this function:
`type Cache struct { algo EvictionAlgo }`

To change behavior:
`cache.algo = LOG` (Least Often Used)
`cache.algo = LRU` (Least Recently Used)
This is powerful because functions are first-class citizens in Go. We don't need heavy class hierarchies; we just pass different functions.

---

## 583. What is the Singleton pattern and how is it safely used in Go?

**Answer:**
A Singleton ensures only one instance exists.
In Go, we use `sync.Once`.

```go
var instance *Config
var once sync.Once

func GetConfig() *Config {
    once.Do(func() {
        instance = loadConfig()
    })
    return instance
}
```
`sync.Once` guarantees atomic, thread-safe execution exactly once, even if 100 goroutines call `GetConfig()` simultaneously at startup.

---

## 584. How do you write a middleware chain in Go?

**Answer:**
We often use the "Russian Doll" model.
`func(next Handler) Handler`.

To chain them:
`auth(logging(recovery(finalHandler)))`.

Libraries like **Alice** or **Chi** provide a `Use()` method to make this linear:
`r.Use(Recovery, Logging, Auth)`
Internally, it loops through the list and wraps them in reverse order, so the first one registered is the first one executed.

---

## 585. How do you use interfaces to decouple layers?

**Answer:**
We define interfaces **where they are used** (Consumer side).

The Service Layer needs to save data. It defines: `type UserRepository interface { Save(User) }`.
The Database Layer implements this.
The Service doesn't know about SQL or Mongo. It just knows `Save`. This allows testing the Service with a MockRepository independently of the DB implementation, effectively decoupling the business logic from infrastructure.

---

## 586. How do you implement the Observer pattern using channels?

**Answer:**
Observer allows 1 publisher to notify N subscribers.

Publisher: Holds a list of channels `[]chan Event`.
Subscribe: Create a channel, add to list.
Publish:
```go
for _, ch := range subscribers {
    go func(c chan Event) { c <- event }(ch)
}
```
Crucial: publish in a non-blocking way (use `select { case c <- event: default: }`) or use buffered channels, so one slow subscriber doesn't block the publisher.

---

## 587. What is the repository pattern and when do you use it?

**Answer:**
The Repository Pattern mediates between the Domain and Data Mapping layers.

It speaks "Business Language" (`FindActiveUsers`) rather than "SQL Language" (`SELECT * FROM...`).
Use it when you want to:
1.  Swap storage backends (SQL -> NoSQL).
2.  Centralize complex query logic (caching, joins).
3.  Test business logic without a real DB.
Don't use it if your app is a simple CRUD wrapper; it just adds boilerplate.

---

## 588. How would you create a CQRS architecture in Go?

**Answer:**
CQRS separates **Command** (Writes) from **Query** (Reads).

**Write Side**: `CommandHandler`. Receives `CreateOrderCmd`. Validates, Writes to Master DB, publishes `OrderCreated` event to Kafka.
**Read Side**: `QueryHandler`. Reads from a Read-Replica or a separate Materialized View (Elasticsearch).
In Go, we often use separate packages `cmd` and `query` to enforce this separation physically, ensuring that "Reads" never trigger side effects.

---

## 589. How do you design a plug-in architecture in Go?

**Answer:**
(See Q 398).
We typically use **HashiCorp/go-plugin** via gRPC.

We define a Protobuf contract.
The Core App is the gRPC Client.
The Plugin is a standalone binary running a gRPC Server.
The Core App spawns the Plugin process and connects over a local socket. This is robust, language-agnostic, and safe (plugin crash doesn't kill app).

---

## 590. What is a “clean architecture” in Go projects?

**Answer:**
It’s strictly organizing code by dependency rules (The Onion).

**Inner**: Domain (Structs). No dependencies.
**Middle**: Use Cases (Service Logic). Depends on Domain.
**Outer**: Adapters (HTTP, SQL, CLI). Depends on Use Cases.

Project structure:
`/internal/domain`
`/internal/service`
`/internal/handler`
`/internal/repository`
Using implicit interfaces ensures the inner layers never import outgoing layers, keeping the core logic pure.

---

## 591. How do you structure a multi-module Go project?

**Answer:**
We use a **Monorepo** with Go Workspaces (`go.work`).

Root: `go.work` (lists modules).
`/pkg/common` (go.mod) -> Utils.
`/services/auth` (go.mod) -> Auth Service.
`/services/billing` (go.mod) -> Billing Service.

Each service has its own `go.mod`. Shared code lives in `pkg/common`.
Workspaces allow us to modify `common` and `auth` simultaneously and run tests across boundaries without publishing `v0.0.1` tags constantly.

---

## 592. How do you decouple business logic from transport layers?

**Answer:**
Your "Service" struct should not know about HTTP.
**Bad**: `func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request)`.
**Good**: `func (s *Service) CreateUser(ctx context.Context, u User) (User, error)`.

The HTTP Handler is an **Adapter**. It parses JSON, calls the Service, and formats the result. This allows you to expose the exact same Service logic via gRPC, CLI, or Kafka Consumer just by writing a different adapter.

---

## 593. How would you implement retryable jobs in Go?

**Answer:**
We use a library like **River** or **Asynq** (based on Redis).

Concept:
1.  Job fails.
2.  Worker catches error.
3.  Calculates `next_try = now + 2^attempts`.
4.  Pushes job to a **ZSET** (Sorted Set) in Redis with score = timestamp.
5.  A Scheduler checks ZSET periodically: "Are any jobs due?" and moves them back to the active queue.

---

## 594. How would you design a billing system in Go?

**Answer:**
Billing requires **ACID** and **Idempotency**.

1.  **Double-Entry Ledger**: Every transaction is two rows (Debit User, Credit Revenue). Sum must be zero.
2.  **Idempotency Keys**: passed from the frontend to prevent double-charging.
3.  **State Machine**: `PENDING -> SUCCESS` or `PENDING -> FAILED`. Move formatting only happens in a DB transaction.
4.  **Math**: Use `int64` (cents) or a decimal library. NEVER use `float64` for money.

---

## 595. How would you scale a notification system written in Go?

**Answer:**
The bottleneck is usually the 3rd party (SendGrid/Twilio) rate limits.

Architecture:
1.  **Ingest API**: Pushes to Kafka.
2.  **Workers**: Read Kafka.
3.  **Rate Limiter**: Use Redis Token Bucket (global).
4.  **Sharding**: If SendGrid allows 10k/sec, we shard by UserID so that one noisy user doesn't block everyone else.
We launch 100 Go workers. If we hit rate limits, we back off and re-queue.

---

## 596. How do you build a real-time leaderboard in Go?

**Answer:**
We use **Redis Sorted Sets (`ZSET`)**.

Go code: `redis.ZAdd("leaderboard", score, userID)`.
To get Top 10: `redis.ZRevRangeWithScores("leaderboard", 0, 9)`.
Redis handles the sorting in O(log(N)).
If we need valid real-time updates to browsers, the Go server subscribes to Redis Pub/Sub updates and pushes changes via WebSockets to connected clients.

---

## 597. How would you implement transactional emails in Go?

**Answer:**
**Transactional Outbox Pattern**.

We don't send emails during the HTTP request (too slow, unreliable).
1.  DB Tx: Insert User + Insert `EmailTask` in `outbox` table. Commit.
2.  Background Go Worker: Polls `outbox`.
3.  Calls SendGrid.
4.  On success, deletes standard from `outbox`.
This guarantees that if the DB commit fails, no email is sent. If the email fails, we retry.

---

## 598. How do you model money and currencies in Go?

**Answer:**
We create a struct:
```go
type Money struct {
    Amount   int64  // in minor units (cents)
    Currency string // "USD", "EUR"
}
```
We implement methods `Add`, `Sub`, `Allocate` (split without losing cents).
We enforce that you cannot Add USD to EUR without an explicit conversion rate.
We usually store it in DB as Composite Column (`price_amount`, `price_currency`) or JSONB.

---

## 599. How do you do dependency injection in Go?

**Answer:**
We prefer **Constructor Injection**.

```go
func NewService(db DB, cache Cache, mailer Mailer) *Service {
    return &Service{db: db, cache: cache, mailer: mailer}
}
```
In `main.go`, we wire it up manually:
`svc := NewService(postgres, redis, sendgrid)`.

We avoid "Magic" DI containers (like Spring) or `reflection` based injectors (like Uber Dig) unless the app is massive, because manual wiring is explicit, compile-time checked, and easier to follow.

---

## 600. How do you create a rule engine in Go?

**Answer:**
For simple rules: **Chain of Responsibility**.
`Type Rule func(Context) bool`.
List: `[]Rule{RuleA, RuleB, RuleC}`. Loop and check.

For complex business rules (e.g., Insurance):
We might use an **Expression Evaluator** like `govaluate` or `expr`.
Store rules in DB: `price > 100 && user.age < 25`.
Go loads the string, compiles it (cached), and evaluates it against the struct parameters at runtime.
