# 🧠 **581–600: Design Patterns, Architecture & Real-World Scenarios**

### 581. How do you implement the Factory pattern in Go?
"I use a simple function.
`func NewStore(type string) Store`.
Inside: `switch type { case "memory": return &memStore{}; case "postgres": return &pgStore{} }`.
Since Go doesn't have classes or constructors, the 'Factory' is just the idiomatic `New...` function that returns an interface."

#### Indepth
Return Concrete Types where possible, accepting Interfaces. `func New() *Type`. However, the Factory Pattern specifically exists to return the **Interface** so the caller doesn't know the implementation. This is useful for plugins or drivers (`database/sql`), but overuse leads to code that is hard to navigate (Click to Definition -> Interface, not Code).

---

### 582. How do you use the Strategy pattern in Go?
"I define an **Interface**.
`type EvictionStrategy interface { Evict(c *Cache) }`.
I implement structs: `LRU`, `LFU`, `Random`.
My `Cache` struct has a field: `strategy EvictionStrategy`.
I can swap the strategy at runtime: `cache.SetStrategy(&LFU{})`.
It’s cleaner than a giant `if-else` block inside the cache logic."

#### Indepth
Function Types as Strategies: You don't always need a struct/interface. `type EvictFunc func(*Cache)`. The strategy can be a simple function closure. `cache.SetStrategy(func(c *Cache) { ... })`. This represents the functional programming approach to the Strategy pattern, widely used in Go middleware.

---

### 583. What is the Singleton pattern and how is it safely used in Go?
"I use `sync.Once`.
`var instance *DB; var once sync.Once`.
`func GetDB() *DB { once.Do(func() { instance = connect() }); return instance }`.
This guarantees initialization runs exactly once, even if 100 goroutines call `GetDB` simultaneously. Using a global variable without `sync.Once` is not thread-safe."

#### Indepth
`sync.Once` uses an atomic counter and a mutex under the hood. It checks `done==1` (fast path, atomic load). If 0, it locks, checks again, runs the function, sets `done=1`. This "Double-Checked Locking" optimization makes accessing a singleton cheap enough to simple calls in hot loops.

---

### 584. How do you write a middleware chain in Go?
"I use a helper function to wrap them.
`func Chain(h http.Handler, m ...Middleware) http.Handler`.
I loop backwards through the middleware slice, wrapping the handler.
`for i := len(m)-1; i >= 0; i-- { h = m[i](h) }`.
This creates an onion: Request -> M1 -> M2 -> Logic. Response Logic -> M2 -> M1 -> Client."

#### Indepth
**Decorator Pattern**. Middleware is essentially decorating the `ServeHTTP` method. Use this pattern for Cross-Cutting Concerns (Logging, Tracing, Metrics, Auth). Business logic (like Input Validation) strictly belongs in the Handler or Service layer, NOT in middleware (which should be generic).

---

### 585. How do you use interfaces to decouple layers?
"My Logic Layer depends on a `Repository` interface, not the `SQL` struct.
`type Service struct { repo UserRepository }`.
This allows me to inject a Mock Repo for testing, or swap Postgres for Mongo without changing a single line of business logic. It follows the **Dependency Inversion Principle**."

#### Indepth
**Hexagonal Architecture** (Ports and Adapters). The "Port" is the Interface (`UserRepository`). The "Adapter" is the struct (`PostgresUserRepo`). The core application logic interacts only with Ports. The `main` function wires the specific Adapters. This makes the core "Infrastructure Agnostic".

---

### 586. How do you implement the Observer pattern using channels?
"I allow subscribers to register a channel.
`type Publisher struct { subs []chan Event }`.
`func (p *Publisher) Subscribe() chan Event`.
When an event occurs, I loop and send to all channels.
Caution: I use a non-blocking send or buffered channel so that one slow subscriber doesn't block the entire publisher."

#### Indepth
Memory Leaks! If a subscriber stops reading but doesn't unsubscribe (close channel), the publisher's send will block forever (unbuffered) or fill the buffer and then block. Always include a mechanism to `Unsubscribe` or use `select { case ch <- msg: default: log.Warn("dropped") }` to handle slow consumers.

---

### 587. What is the repository pattern and when do you use it?
"It abstracts data access.
Interface: `GetByID`, `Save`, `Delete`.
It hides the details (SQL queries, Redis keys).
I use it when the domain logic is complex. It prevents SQL strings from leaking into my Controllers.
However, for simple CRUD apps, it might be overkill (an unnecessary abstraction layer)."

#### Indepth
**Unit of Work**. The Repository handles *single* entity modifications. If you need to update a User AND a Wallet in one Transaction, you need a UoW pattern or pass the `sql.Tx` through the repository methods. `repo.WithTx(tx).Save(user)`. This keeps transaction boundaries explicit in the Service layer.

---

### 588. How would you create a CQRS architecture in Go?
"I split the app into **Commands** (Write) and **Queries** (Read).
**Commands**: `CreateUser(cmd)`. Validates and writes to DB. Returns ID. No data.
**Queries**: `GetUserView(id)`. Reads from a read-optimized view (maybe a flat JSON table).
This allows me to scale reads independently (Read Replicas) and optimize complex writes."

#### Indepth
 CQRS often pairs with **Event Sourcing**. Instead of storing current state (`Balance=100`), store events (`Deposited 50`, `Deposited 50`). To read, you replay events (or use a snapshot). This provides a perfect audit trail but adds massive complexity (Schema evolution, Snapshotting). Use with caution.

---

### 589. How do you design a plug-in architecture in Go?
"Two ways:
1.  **Go Plugins (`plugin` package)**: Loads shared libraries (`.so`). Hard to use, Linux-only, strict versioning.
2.  **RPC/HashiCorp Plugin**: The plugin is a separate binary process. My app talks to it via gRPC over localhost. This is how Terraform works. It’s robust because a crashing plugin doesn't crash the main app."

#### Indepth
Security: Loading `.so` plugins is dangerous (`init()` function runs as root/user). RPC plugins essentially sandbox the plugin in its own process. You can even run the plugin in a strict container or restricted user account to minimize the blast radius of a compromised plugin.

---

### 590. What is a “clean architecture” in Go projects?
"Concentric circles.
**Entities** (Inner): Pure Go structs. No tags, no imports.
**Use Cases**: Business logic. Depends on Entities.
**Controllers/Gateways**: HTTP handlers, SQL implementations. Depend on Use Cases.
**External**: DB, Web.
Everything points inward. I implement this using standard project layout (`internal/domain`, `internal/service`, `internal/handler`)."

#### Indepth
**The Dependency Rule**: Source code dependencies can only point *inward*. `Domain` knows nothing about `SQL`. `Services` know nothing about `HTTP`. This makes the inner circle reusable. You could wrap the same `Service` in a CLI command, a gRPC server, or a REST API without changing a line of logic.

---

### 591. How do you structure a multi-module Go project?
"I use a **Go Workspace** (`go.work`).
Root `go.work` -> `use ./module-a`, `use ./module-b`.
Each module has its own `go.mod`.
This allows me to develop them together as a monorepo while keeping their dependencies separate. Module B imports Module A via local path during dev, and git tag in prod."

#### Indepth
Before `go.work`, we used the `replace` directive in `go.mod`. `replace github.com/my/lib => ../lib`. The workspace file is cleaner because it's *local to your machine* (often gitignored) and doesn't accidentally get committed to prod code, avoiding "cannot find module ../lib" errors in CI.

---

### 592. How do you decouple business logic from transport layers?
"I never put business logic in the HTTP handler.
Handler: `Parse JSON` -> `Call Service.DoThing()` -> `Format Response`.
Service: `func DoThing(Input) (Output, error)`.
The Service knows nothing about HTTP (no `gin.Context`). It can be called by a gRPC handler, a CLI command, or a background worker equally well."

#### Indepth
**Context Pollution**. Don't pass `gin.Context` to the service. Pass `context.Context` (stdlib). The HTTP handler should extract all params (ID, JSON body) and pass them as Go types (`struct`, `int`) to the service. The service never imports `gin` or `http`.

---

### 593. How would you implement retryable jobs in Go?
"I use a queue with a `RetryCount` field.
Worker pops job.
If fails: `job.Retries++`.
If `job.Retries < Max`: Put back in queue (with exponential backoff).
If `job.Retries >= Max`: Move to **Dead Letter Queue (DLQ)** for manual inspection.
Usually, I use a library like `River` or `Asynq` (Redis-backed) to handle this reliability."

#### Indepth
**Exponential Backoff + Jitter**. Retry interval should be `Base * 2^Retry`. Add random jitter (`+/- 10%`) to prevent the "Thundering Herd" problem where 10,000 failed jobs all retry at the exact same millisecond, crashing your database again.

---

### 594. How would you design a billing system in Go?
"**ACID** is king.
Use a relational DB (Postgres).
Use transactions for everything.
Double-entry bookkeeping (Credit one account, Debit another).
Idempotency keys on every transaction.
In Go: `tx, _ := db.Begin()`, pass `tx` to all repository methods, and `defer tx.Rollback()`, `tx.Commit()` at the end."

#### Indepth
Floating Point Math is forbidden. Use `int64` (micros/nanos) or `pgtype.Numeric`. JavaScript clients struggle with `int64` (max safe integer is 2^53). Serialize amounts as **Strings** in JSON (`"amount": "10.00"`) to be safe, or splitting them (`dollars: 10, cents: 0`).

---

### 595. How would you scale a notification system written in Go?
"Decouple Ingestion from Delivery.
API -> Kafka Topic (`notifications`).
Go Workers (Consumers) read Kafka.
Workers invoke 3rd party APIs (Twilio/SendGrid).
To scale, I just add more Worker Pods.
Since 3rd parties have rate limits, I implement a **Rate Limiter** per worker or robust backoff logic."

#### Indepth
**At-Least-Once Delivery**. Kafka guarantees the message is delivered, but your worker might crash *after* sending the email but *before* committing the offset. The user gets 2 emails. Design APIs to be Idempotent (`msg_id` deduplication in the worker) to handle this gracefully.

---

### 596. How do you build a real-time leaderboard in Go?
"I don't use the SQL DB for sorting.
I use **Redis Sorted Sets**.
Go App -> `Redis.ZAdd("leaderboard", score, user)`.
Read -> `Redis.ZRevRange`.
It’s O(log N). Even with 10M players, retrieving the Top 100 is instant. Storing this in SQL (`ORDER BY score DESC`) would kill the DB."

#### Indepth
**Skip List**. Redis Sorted Sets (`ZSET`) are implemented using Skip Lists. This probabilistic data structure allows fast insertion and ranking (finding the rank of a user, e.g., "You are #4521"). Updating high-frequency scores in SQL locks rows; in Redis, it's a lock-free memory update.

---

### 597. How would you implement transactional emails in Go?
"I listen for domain events.
`UserCreated` event -> Event Bus -> `EmailHandler`.
The handler renders the template and calls the email provider.
Crucially, if the email fails, I don't rollback the `UserCreated` transaction. I retry the email independently (Eventual Consistency)."

#### Indepth
**Transactional Outbox Pattern**. Save the email task to a `outbox` table in the *same transaction* as the user creation. Then, a background poller picks up the `outbox` row and sends it. This guarantees atomicity: "If User is created, Email task is created". No orphan users.

---

### 598. How do you model money and currencies in Go?
"**Never use float64!** Calculate in cents (integers).
$10.00 = `1000`.
Or use `shopspring/decimal`.
I always store the currency code (`USD`) alongside the amount.
Struct: `type Money struct { Amount int64; Currency string }`."

#### Indepth
**Rounding Issues**. When splitting money (3 people split $10), you get $3.333... You must decide where the extra penny goes. The "Allocation" algorithm creates `[334, 333, 333]`. Never rely on default float rounding; explicitly handle the remainder.

---

### 599. How do you do dependency injection in Go?
"I prefer **Constructor Injection**.
`func NewService(db *DB, logger *Logger) *Service`.
I wire everything up in `main.go`.
`db := ...`
`svc := NewService(db, log)`
`handler := NewHandler(svc)`.
I avoid DI frameworks (like Uber Dig) unless the app is massive, because they hide the dependency graph and make code harder to follow."

#### Indepth
**Wire** (by Google) is a Code-Generation DI tool. It's safer than Reflection-based DI (Dig/Fx) because it generates standard Go code at compile time. If a dependency is missing, your code won't compile. This provides the convenience of auto-wiring with the safety of explicit composition.

---

### 600. How do you create a rule engine in Go?
"I define an Interface `Rule { Evaluate(Context) bool }`.
I create a chain of Rules: `[]Rule{RuleA{}, RuleB{}}`.
Run: `for r := range rules { if !r.Evaluate(ctx) { return Fail } }`.
For dynamic rules (defined by users), I use an expression language like `expr` to parse string rules safely."

#### Indepth
**AST Traversal**. A rule engine basically evaluates an Abstract Syntax Tree. `expr` compiles the string `user.Age > 18` into a bytecode VM. It's safe/sandboxed (no infinite loops, no file access). For hardcoded rules, use the **Specification Pattern** (Interface `IsSatisfiedBy(candidate)`).
