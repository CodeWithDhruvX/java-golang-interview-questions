# 🟢 Go Theory Questions: 241–260 Architecture and Design Patterns

## 241. How do you implement the Clean Architecture in Go?

**Answer:**
Clean Architecture (often called the Onion Architecture or Hexagonal Architecture) is a software design philosophy that organizes code into concentric circles of dependency, ensuring that the core business logic remains independent of external frameworks.

The golden rule is the **Dependency Rule**: source code dependencies can only point **inward**. The center circle typically contains **Entities** (pure business structs with no tags). Surrounding that are **Use Cases** (Interactors) which orchestrate the flow of data. The outermost layer contains **Infrastructure** (Controllers, Gateways, DB implementation, HTTP Handlers). In Go, we rigidly enforce this using **Interfaces**. The Use Case layer defines a `UserRepository` interface, and the Infrastructure layer implements it.

In the real world, this makes applications testable and swappable. You can swap Postgres for MongoDB by just writing a new adapter in the outer layer without touching a single line of business logic in the center. The trade-off is boilerplate: you end up writing a lot of interface definitions and manual mapping code to convert "Database Models" (with SQL tags) into pure "Domain Models."

---

## 242. How do you break a Monolith into Microservices in Go?

**Answer:**
We typically use the **Strangler Fig Pattern**. We don't rewrite the whole system; we gradually replace specific pieces of functionality with new microservices.

Mechanically, we identify a "Seam" or a specific "Bounded Context" (like the "Billing" module). We define a strict interface for it. Then, we build a new Go service for Billing. We stick a proxy (or modify the monolith's code) to redirect calls from the internal `BillingService` class to the new gRPC or REST client. Over time, the monolith shrinks as more features are "strangled" out into independent services.

The biggest challenge isn't the code—it’s the **Data**. You cannot easily tear apart a massive shared database. We often have to implement **Double Writes** or use **CDC (Change Data Capture)** tools like Debezium to sync data from the monolith's legacy DB to the new microservice's DB until the migration is complete.

---

## 243. What is the Functional Options pattern?

**Answer:**
The Functional Options pattern is the idiomatic Go way to create complex APIs with highly configurable but optional parameters.

Instead of a constructor taking 10 arguments (`NewServer(addr, port, timeout, logger...)`), or a config struct that forces users to define empty fields, we pass a variadic slice of functions: `NewServer(addr, opts ...Option)`. An `Option` is a closure: `type Option func(*Server)`.
Example: `WithTimeout(d time.Duration)` returns a closure that sets `s.timeout = d`.

This is superior for three reasons. First, it allows for sensible defaults (if no options are passed, use the default). Second, it’s **Future-Proof**—you can add new options later without breaking the API signature for existing users. Third, it allows for validation logic inside the option setter itself.

---

## 244. How do you implement the Circuit Breaker pattern?

**Answer:**
A Circuit Breaker is a wrapper around a network call that detects failures and prevents your application from wasting resources trying to call a dead service.

It functions like a State Machine.
1. **Closed**: Requests flow normally. If errors exceed a threshold (e.g., 50% failure rate), it trips to Open.
2. **Open**: All requests fail immediately (Fail Fast) without hitting the network. This prevents "Thread Starvation" where all your goroutines are stuck waiting for a timeout.
3. **Half-Open**: After a sleep period, it lets *one* request through. If it succeeds, the circuit closes (heals); if it fails, it stays Open.

In Go, we use libraries like `gobreaker` or `hystrix-go` to wrap our `http.Client`. It is essential in microservices to prevent **Cascading Failures**, where one down service causes the entire mesh to crash because everyone is waiting on everyone else.

---

## 245. What is the Saga Pattern in Go?

**Answer:**
The Saga Pattern is a mechanism for managing distributed transactions that span multiple microservices, where you cannot rely on a traditional database ACID lock (2-Phase Commit).

Since we can't lock two different databases at once, we break the transaction into a sequence of local steps (e.g., BookFlight, BookHotel, ChargeCard). Crucially, for every step, we define a **Compensating Action** (UndoFlight, UndoHotel, RefundCard). If step 3 fails, we execute the compensations for steps 2 and 1 strictly in reverse order to roll back the state.

In Go, we often implement this using an **Orchestrator** approach. A central "Order Service" calls the downstream services. We use a `defer` stack approach: as we execute each forward step successfully, we `defer` the push of its generic compensation function onto a stack. If an error occurs, we simply pop and execute the stack.

---

## 246. How do you architect a Pub/Sub system?

**Answer:**
Internally within a single Go app, we build Pub/Sub using **Channels**. Externally, we use a message broker like Kafka or RabbitMQ.

For an internal Event Bus, we typically have a central `Hub` struct protecting a map of `topic -> []chan Event`. When a message is `Publish`ed, the Hub locks the map, iterates through the slice of subscribers, and sends the data.

The critical architectural detail is **Isolation**. You must never let a slow subscriber block the publisher. To solve this, we strictly use **Buffered Channels** for subscribers, or we launch a dedicated goroutine for the dispatch (`go ch <- msg`). If a subscriber's buffer is full, we must decide whether to drop the message (lossy) or block (backpressure), which entails a significant trade-off between reliability and latency.

---

## 247. How do you implement the Worker Pool pattern?

**Answer:**
A Worker Pool is a concurrency pattern designed to limit the number of active goroutines processing tasks, protecting your system from resource exhaustion.

Mechanically, it involves three parts: a `Job` channel, a `Result` channel, and a fixed number of **Worker Goroutines**.
You start N workers: `for i := 0; i < NumWorkers; i++ { go worker(jobs, results) }`.
The `worker` function is a simple `range` loop over the `jobs` channel.

This creates a "Leaky Bucket" effect. No matter if you pour 1 million jobs into the channel instantly, the processing rate is strictly capped at `NumWorkers` (e.g., 5 concurrent DB connections). This is mandatory when processing webhooks or file uploads to avoid hitting `Too Many Open Files` (ULIMIT) errors or crashing the garbage collector.

---

## 248. What is the Adapter Pattern usage in Go?

**Answer:**
The Adapter Pattern allows incompatible interfaces to work together. In Go, it is ubiquitous, most notably with `http.HandlerFunc`.

Imagine you have a function `myHandler(w, r)`. The `http.Handler` interface requires a struct with a `ServeHTTP` method. You can't pass your function directly. The `http.HandlerFunc` type acts as an adapter: it’s a type definition `type HandlerFunc func(...)` that implements `ServeHTTP` by simply calling itself.

We use this extensively to retrofit third-party libraries into our Clean Architecture. If a logging library returns `Print(msg string)` but our internal interface expects `Log(msg []byte)`, we write a small struct that wraps the 3rd party logger and converts the types. This keeps our core logic decoupled from the specific library we chose.

---

## 249. How do you implement the Singleton pattern in Go?

**Answer:**
We implement Singletons using the `sync.Once` primitive.

Classic "check-lock-check" implementations for singletons are dangerous in Go due to potential race conditions and memory visibility issues. `sync.Once` is safer and cleaner.
```go
var instance *Manager
var once sync.Once
func GetInstance() *Manager {
    once.Do(func() {
        instance = &Manager{}
    })
    return instance
}
```
In modern Go architecture, however, we generally **avoid** global singletons. Instead, we perform "dependency injection" relative to the `main()` function. We create the single instance of the database or config at startup and pass it explicitly to every handler or service that needs it. This makes testing significantly easier because you can inject a mock instance without global hacks.

---

## 250. How do you handle configuration in a 12-Factor Go App?

**Answer:**
The 12-Factor App methodology states that config should be stored in the **Environment**.

In Go, we rely on the `os` package (`os.Getenv`) and often use libraries like `kelseyhightower/envconfig` or `viper`. We define a struct with tags:
`type Config struct { DBUrl string \`envconfig:"DB_URL"\` }`.
At startup, we parse the environment variables directly into this struct.

This decoupling allows the same binary to run in Dev, Staging, and Production without changes. The container orchestrator (Kubernetes) injects the different configs as environment variables. We strictly avoid baking config files (`config.json`) inside the binary image, as that requires a rebuild just to change a database password.

---

## 251. What is the Decorator Pattern in Go (Middleware)?

**Answer:**
In Go, the Decorator pattern is realized through **Middleware**. It allows you to wrap functionality around an existing function to extend its behavior without modifying the original code.

The signature is `func(http.Handler) http.Handler`. You accept a handler, and you return a *new* handler that executes some logic (like Logging or Auth) before calling `original.ServeHTTP()`.

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w, r) {
        start := time.Now()
        next.ServeHTTP(w, r) // Call the inner handler
        log.Println(time.Since(start))
    })
}
```
This nesting creates a "Chain of Responsibility." We use it for cross-cutting concerns: Authentication, Tracing (OpenTelemetry), GZIP compression, and Panic Recovery. It keeps the actual business handlers clean and focused purely on logic.

---

## 252. How do you implement the Repository Pattern?

**Answer:**
The Repository Pattern mediates between the domain and data mapping layers using a collection-like interface for accessing interactions.

In Go, we define it as an interface in the domain layer:
`type UserRepository interface { GetByID(id int) (*User, error) }`.
Then, in the infrastructure layer, we implement it: `type PostgresUserRepo struct { DB *sql.DB }`.

This abstraction allows us to decouple the business logic from the specific database technology. If we want to switch from Postgres to MongoDB, or use an In-Memory Mock for unit testing, we just provide a different struct that satisfies the interface. The domain logic remains unchanged, unaware of whether the user came from a SQL table or a JSON file.

---

## 253. What is the difference between Orchestration and Choreography in Sagas?

**Answer:**
These are the two ways to coordinate a distributed transaction (Saga).

**Orchestration** relies on a central conductor (e.g., the Order Service). It tells the Payment Service "Pay", then waits. If successful, it tells Inventory "Reserve". It tracks the state and handles errors centrally. This is easier to debug but creates a tight coupling and a collection point for logic.

**Choreography** is decentralized. The Order Service publishes an event `OrderCreated`. The Payment Service listens, processes it, and publishes `PaymentProcessed`. The Inventory Service listens to that. There is no central manager; the process emerges from the interaction. This is more decoupled but harder to monitor—you can lose track of "who has the ball" in a complex flow. In Go, we usually utilize Orchestration for critical flows to maintain control.

---

## 254. How do you implement Graceful Shutdown in a Go service?

**Answer:**
Graceful Shutdown ensures that when a server restarts, it finishes currently active requests before stopping, rather than severing connections mid-stream.

We use `os.Signal` to listen for termination signals (`SIGINT`, `SIGTERM`). When a signal is received, we call `server.Shutdown(ctx)`.
This method stops the listener from accepting new connections but keeps the main process alive until all active handlers return or the context times out.

Crucially, in a microservices architecture, this must be paired with a "Readiness Probe" in Kubernetes. When shutdown starts, the app must immediately fail the readiness probe so K8s stops sending new traffic, while the app drains the existing traffic for the next 30 seconds.

---

## 255. What is formatting vs. linting in Go project structure?

**Answer:**
**Formatting** (`gofmt`) is about code style (indentation, spacing). Go is unique because there is no debate; the tool enforces one standard style. It modifies your code automatically.

**Linting** (`golangci-lint`) is about code correctness and quality. It checks for bugs, anti-patterns, and complexity.

In a project, we enforce both in the CI pipeline. Formatting ensures the code looks uniform regardless of who wrote it. Linting prevents logical errors (like unhandled errors, unused variables, or ineffective assignments). We typically configure `golangci-lint` to run a suite of 50+ linters (govet, staticcheck, errcheck, etc.) to catch issues effectively before code review.

---

## 256. How do you design for Failure in distributed Go systems?

**Answer:**
"Everything fails all the time." We design with **Defensive Programming**.

We assume the database will timeout, the network will be flaky, and the disk will fill up.
At the architectural level, we implement **Retries** with **Exponential Backoff** (wait 1s, then 2s, then 4s) to avoid thundering herd problems. We use **Timeouts** strictly on every single I/O call.

We also implement **Idempotency**. If a client retries a "Charge Payment" call because the network dropped the response, the server must handle receiving the same request ID twice without charging the user twice. We usually achieve this by caching the request hash or ID in Redis with the result for a short window.

---

## 257. What is Domain-Driven Design (DDD) in Go?

**Answer:**
DDD is a strategic approach to software where the code structure matches the business domain (e.g., Shipping, Billing, Catalog).

In Go, we model this by organizing packages by **Domain**, not by technology. Instead of folders like `controllers/` and `models/`, we have `shipping/` and `billing/`.
Inside `shipping/`, we have everything that domain needs: the logic, the specific DB queries, and the handlers.

We use the **Ubiquitous Language**: the structs and methods should attempt to use the exact same terminology as the business experts. If the business calls it a "Consignment", the code should be `type Consignment`, not `type ShipmentPkg`. This reduces the cognitive translation gap between requirements and implementation.

---

## 258. How do you implement Event Sourcing in Go?

**Answer:**
Event Sourcing means storing the **State Changes** (Events) rather than the current state.
Instead of updating a row `Balance = 100`, we insert a new row `Event: Deposited $50`. The current balance is calculated by replaying all historical events.

In Go, we implement this by having an `Append(event)` method in our store. We often use a "Snapshot" optimization—every 100 events, we save the calculated state so we don't have to replay from the beginning of time.

This pattern provides a perfect audit trail and allows "Time Travel" debugging. However, it introduces massive complexity for simple queries (you can't just `SELECT * WHERE balance > 0`), so we often pair it with **CQRS** (Command Query Responsibility Segregation) to maintain a separate "Read Model" that is optimized for querying.

---

## 259. What is CQRS (Command Query Responsibility Segregation)?

**Answer:**
CQRS splits your application into two distinct parts: one for **Writing** (Commands) and one for **Reading** (Queries).

The **Write Side** (Command Model) focuses on business logic and validation. It might store data in a normalized Relational DB.
The **Read Side** (Query Model) focuses on fast retrieval. It might store the *same* data in a denormalized NoSQL store (like ElasticSearch or Redis) optimized for the specific UI views.

In Go, we keep these models effectively separate. A Command handler processes input and emits an event. An async worker picks up that event and updates the Read DB. This allows independent scaling—if you have 1000x more reads than writes, you can scale the Read replicas without paying for expensive Write hardware. The trade-off is **Eventual Consistency**—the user might write data and not see it immediately in the list.

---

## 260. How do you implement Hexagonal Architecture (Ports and Adapters)?

**Answer:**
Hexagonal Architecture is meant to allow an application to be driven by users, programs, automated tests, or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases.

In Go, **Ports** are Interfaces. We define a `Driving Port` (e.g., an interface for "Place Order") that the outside world uses, and a `Driven Port` (e.g., an interface for "Save Order") that our app uses to talk to the DB.

**Adapters** are the concrete implementations. A `RestHandler` is an adapter that translates HTTP to the `Driving Port`. A `PostgresRepo` is an adapter that implements the `Driven Port`.
This symmetry means the core application code doesn't care if the "Driver" is an HTTP request or a CLI command, nor if the "Driven" side is a real DB or a mock. It effectively decouples the "Inside" from the "Outside."
