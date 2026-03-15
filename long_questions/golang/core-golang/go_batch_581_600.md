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

### Explanation
The Factory pattern in Go is implemented as a simple function that returns an interface type. This allows the function to decide which concrete implementation to return based on input parameters, while clients work with the interface abstraction. The factory encapsulates object creation logic and provides a clean way to create different implementations without exposing construction details.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement the Factory pattern in Go?
**Your Response:** "I implement the Factory pattern in Go using a simple function that returns an interface type. I define an interface like `Store` with the methods I need, then create concrete implementations like `DiskStore` and `MemStore`. The factory function `NewStore()` takes a parameter to determine which implementation to return and returns the interface type. This approach encapsulates object creation logic and allows me to switch implementations easily without changing the client code. Clients work with the interface abstraction, making the code more flexible and testable. Go's implicit interfaces make this pattern particularly natural to implement."

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

### Explanation
The Strategy pattern in Go can be implemented using function types or interfaces to define a family of algorithms. Different implementations of the strategy can be swapped at runtime, allowing the client to choose which algorithm to use. Function types work well for simple strategies, while interfaces are better for more complex strategies that require state or multiple methods.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use the Strategy pattern in Go?
**Your Response:** "I use the Strategy pattern in Go by defining a family of algorithms as function types or interfaces that can be swapped at runtime. For simple cases, I use function types like `PaymentStrategy` where different payment methods implement the same function signature. For more complex strategies, I define interfaces with multiple methods. The key is that I can pass different strategy implementations to the same function, like `Checkout()`, and it will execute the appropriate algorithm. This approach makes the code more flexible and allows me to add new strategies without modifying the core logic. Go's first-class functions make this pattern particularly elegant and easy to implement."

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

### Explanation
The Singleton pattern in Go is safely implemented using sync.Once to ensure lazy initialization happens exactly once, even in concurrent scenarios. The sync.Once type guarantees that the function passed to Do() is executed exactly once, regardless of how many goroutines call GetDatabase() simultaneously. This provides thread-safe lazy initialization without the need for complex locking mechanisms.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Singleton pattern and how is it safely used in Go?
**Your Response:** "I implement the Singleton pattern safely in Go using `sync.Once` to ensure lazy initialization happens exactly once, even with concurrent access. I declare a global instance variable and a sync.Once variable. In the GetDatabase() function, I call once.Do() with the initialization logic. sync.Once guarantees that the initialization function runs exactly once, no matter how many goroutines call GetDatabase() simultaneously. This approach is thread-safe, efficient, and avoids the complexity of manual locking. It's the idiomatic way to implement singletons in Go, providing lazy initialization with guaranteed thread safety without performance overhead."

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

### Explanation
Middleware chains in Go are created by wrapping handlers in layers. Each middleware takes a handler and returns a new handler that performs additional processing before or after calling the wrapped handler. The Chain function iterates through middlewares, wrapping the handler in each one to create a processing pipeline. Libraries like Alice and Chi provide more sophisticated middleware chaining capabilities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write a middleware chain in Go?
**Your Response:** "I create middleware chains in Go by wrapping handlers in layers. Each middleware function takes a handler and returns a new handler that adds behavior before or after calling the wrapped handler. I implement this with a Chain function that iterates through the middleware slice, wrapping the handler in each middleware. This creates a processing pipeline where the request passes through each middleware in order. For more complex scenarios, I use libraries like Alice or Chi which provide built-in middleware stacking capabilities. This approach allows me to compose multiple cross-cutting concerns like authentication, logging, and rate limiting in a clean, reusable way without modifying individual handlers."

---

### Question 585: How do you use interfaces to decouple layers?

**Answer:**
The **Business Logic** layer should define the interfaces it needs (Dependency Inversion), and the **Storage** layer should implement them.
`package business` defines `UserRepository interface`.
`package postgres` implements `UserRepository`.
This allows `business` to be tested with mocks and not depend on SQL.

### Explanation
Decoupling layers using interfaces follows the Dependency Inversion Principle where the business logic layer defines the interfaces it needs, and the storage layer implements those interfaces. This creates a dependency direction pointing inward, allowing the business layer to be independent of specific storage technologies and easily testable with mock implementations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use interfaces to decouple layers?
**Your Response:** "I decouple layers using the Dependency Inversion Principle where the business logic layer defines the interfaces it needs, and the storage layer implements those interfaces. For example, the business package defines a UserRepository interface, while the postgres package implements it. This approach means the business logic doesn't depend on specific storage technologies like SQL - it only depends on the interface. This allows me to easily swap implementations, test the business layer with mocks, and maintain clean separation of concerns. The dependency direction points inward, making the architecture more flexible and testable without changing the business logic when storage needs change."

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

### Explanation
The Observer pattern in Go can be implemented using channels where subscribers register by providing channels, and the publisher broadcasts messages by sending to all subscriber channels. This leverages Go's built-in concurrency primitives to create a clean, type-safe observer implementation without complex callback mechanisms.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement the Observer pattern using channels?
**Your Response:** "I implement the Observer pattern in Go using channels as the communication mechanism. I create a Broker that maintains a slice of subscriber channels. When a subscriber wants to observe events, it calls Subscribe() which returns a channel for receiving messages. When an event occurs, the Publish() method iterates through all subscriber channels and sends the message to each one asynchronously using goroutines. This approach leverages Go's built-in concurrency primitives to create a clean, type-safe observer pattern without complex callback mechanisms. Channels provide natural backpressure handling and make the observer pattern very idiomatic in Go."

---

### Question 587: What is the repository pattern and when do you use it?

**Answer:**
It abstracts data access.
**Interface:** `GetUser(id int) (*User, error)`
**Impl:** `SqlUserRepository` (using GORM/SQL).
**Use it:** To separate domain logic from database details, allowing easy swapping (Postgres -> Mongo) and unit testing.

### Explanation
The Repository pattern abstracts data access by defining interfaces that hide implementation details. The interface defines data access methods like GetUser, while concrete implementations like SqlUserRepository handle the actual database operations. This separates domain logic from database details, allowing easy swapping of storage implementations and enabling unit testing with mocks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the repository pattern and when do you use it?
**Your Response:** "I use the Repository pattern to abstract data access and separate domain logic from database details. I define an interface with methods like `GetUser(id int) (*User, error)` that represents the data access contract. Then I create concrete implementations like `SqlUserRepository` using GORM or raw SQL. This approach allows me to easily swap storage implementations - for example, switching from Postgres to MongoDB - without changing the domain logic. It also makes unit testing much easier since I can mock the repository interface. The pattern provides a clean separation between business logic and data persistence, making the code more maintainable and testable."

---

### Question 588: How would you create a CQRS architecture in Go?

**Answer:**
**Command Query Responsibility Segregation.**
Split into two models:
1.  **Command (Write):** Methods that mutate state (`CreateOrder`). Often async, uses Event Sourcing.
2.  **Query (Read):** Methods that return data (`GetOrder`). optimized for reads (Materialized Views).
In Go, simple implementation: Separation of `OrderWriter` and `OrderReader` interfaces.

### Explanation
CQRS (Command Query Responsibility Segregation) separates read and write operations into different models. Commands handle write operations that mutate state, often asynchronously with event sourcing. Queries handle read operations optimized for specific use cases with materialized views. In Go, this is implemented by separating interfaces like OrderWriter and OrderReader.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you create a CQRS architecture in Go?
**Your Response:** "I implement CQRS by separating read and write operations into different models. For the write side, I create commands like `CreateOrder` that mutate state, often using event sourcing for audit trails. For the read side, I implement queries like `GetOrder` that are optimized for specific read scenarios using materialized views. In Go, I separate this into distinct interfaces like `OrderWriter` and `OrderReader`. The write model handles business logic and state changes, while the read model is optimized for fast queries. This separation allows me to scale reads and writes independently and optimize each side for its specific requirements. It's particularly useful in complex systems with different read and write patterns."

---

### Question 589: How do you design a plug-in architecture in Go?

**Answer:**
1.  **Go Plugins (`plugin` package):** Load `.so` files at runtime (Linux/Mac only, tricky versioning).
2.  **RPC/Hashicorp Plugin:** Run plugins as separate processes (binary) and communicate via gRPC/net/rpc over localhost. (Used by Terraform). Safer and more robust.

### Explanation
Plugin architecture in Go can be implemented using the built-in plugin package to load .so files at runtime, though this has limitations and versioning challenges. A more robust approach is using RPC-based plugins like Hashicorp's plugin system, where plugins run as separate processes and communicate via gRPC or net/rpc, providing better isolation and safety.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design a plug-in architecture in Go?
**Your Response:** "I design plugin architectures in Go using two main approaches. First, I can use Go's built-in plugin package to load .so files at runtime, though this has limitations - it only works on Linux/Mac and has tricky versioning requirements. For more robust solutions, I prefer the RPC-based approach like Hashicorp's plugin system, where plugins run as separate processes and communicate via gRPC or net/rpc over localhost. This approach is safer and more robust since plugins are isolated in their own processes. Terraform uses this pattern successfully. The RPC approach allows for better error handling, easier debugging, and cross-platform compatibility, making it the preferred choice for production systems."

---

### Question 590: What is a “clean architecture” in Go projects?

**Answer:**
Standard layout (Uncle Bob):
- **Entities (Domain):** Core structs, no deps.
- **Usecases (Service):** Business rules, depends on Entities.
- **Adapters (Controller/Repo):** HTTP handlers, SQL implementations, depends on Usecases.
- **Drivers (Main):** Wires everything up (Router, DB connection).
Ensures deps point **inwards**.

### Explanation
Clean architecture in Go follows Uncle Bob's layered approach where dependencies point inward. Entities contain core business logic with no dependencies. Use cases implement business rules and depend on entities. Adapters handle external concerns like HTTP and databases and depend on use cases. Drivers wire everything together. This ensures the core business logic remains independent of external details.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a "clean architecture" in Go projects?
**Your Response:** "I implement clean architecture in Go following Uncle Bob's layered approach where dependencies point inward. The innermost layer contains Entities - core business structs with no external dependencies. The next layer has Use cases that implement business rules and depend only on entities. The Adapters layer contains HTTP handlers and database implementations that depend on use cases. The outermost Drivers layer wires everything together including routers and database connections. This dependency direction ensures the core business logic remains completely independent of external concerns like databases or web frameworks. This makes the system more testable, maintainable, and flexible to change."

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

### Explanation
Multi-module Go projects are structured using Go Workspaces with a go.work file that references multiple modules. Each module has its own go.mod file, allowing independent versioning and development. Workspaces enable developing multiple related modules simultaneously without constantly publishing to Git tags or using replace directives.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you structure a multi-module Go project?
**Your Response:** "I structure multi-module Go projects using Go Workspaces with a go.work file at the root. Each subdirectory like `/api`, `/libs`, and `/services/payment` has its own go.mod file for independent versioning. The go.work file references all the modules, allowing me to develop them simultaneously without constantly publishing to Git tags or using complex replace directives. This approach is perfect for microservices architectures or projects with shared libraries. It gives me the benefits of separate modules for deployment and versioning while maintaining a smooth development experience. Workspaces make it easy to make changes across modules and test them together before releasing."

---

### Question 592: How do you decouple business logic from transport layers?

**Answer:**
Do **not** use `gin.Context` or `http.ResponseWriter` in your Service methods.
**Bad:** `func (s *Service) Create(c *gin.Context)`
**Good:** `func (s *Service) Create(ctx context.Context, u User) error`
The HTTP Handler parses JSON -> calls Service -> returns JSON.

### Explanation
Decoupling business logic from transport layers involves avoiding framework-specific types in service methods. Instead of accepting gin.Context or http.ResponseWriter, services should accept context.Context and domain types. This keeps business logic independent of HTTP frameworks and makes it reusable across different transport layers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you decouple business logic from transport layers?
**Your Response:** "I decouple business logic from transport layers by avoiding framework-specific types in my service methods. Instead of accepting `gin.Context` or `http.ResponseWriter`, my services accept `context.Context` and domain types like `User`. The HTTP handler is responsible for parsing JSON from the request, calling the service method, and converting the response back to JSON. This approach keeps my business logic completely independent of HTTP frameworks, making it reusable across different transport layers like gRPC, CLI tools, or message queues. It also makes unit testing much simpler since I don't need to mock HTTP frameworks to test my business logic."

---

### Question 593: How would you implement retryable jobs in Go?

**Answer:**
Use a Queue (Redis/RabbitMQ) + Worker.
1.  Worker pulls job.
2.  Executes.
3.  If fail: Check `retry_count`. If < Max, publish back to queue with a delay (Exponential Backoff). If >= Max, move to **Dead Letter Queue (DLQ)**.

### Explanation
Retryable jobs in Go are implemented using a queue system with workers. Workers pull jobs from the queue, execute them, and handle failures by checking retry counts. Failed jobs are requeued with exponential backoff delays up to a maximum retry limit. Jobs that exceed the retry limit are moved to a Dead Letter Queue for manual inspection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you implement retryable jobs in Go?
**Your Response:** "I implement retryable jobs using a queue system like Redis or RabbitMQ with worker processes. Workers pull jobs from the queue and execute them. If a job fails, I check the retry count - if it's below the maximum, I requeue the job with an exponential backoff delay to avoid overwhelming the system. If the job exceeds the maximum retry limit, I move it to a Dead Letter Queue for manual inspection. This approach provides resilience against transient failures while preventing infinite retries. Exponential backoff helps handle temporary issues like network problems or service unavailability. The DLQ ensures problematic jobs don't block the queue and can be analyzed later."

---

### Question 594: How would you design a billing system in Go?

**Answer:**
- **Concurrency:** Use `sync.Mutex` or DB Record Locking (`SELECT FOR UPDATE`) to prevent double-spending.
- **Precision:** NEVER use `float64`. Use `int64` (cents) or `shopspring/decimal`.
- **Idempotency:** Crucial for API.
- **Audit:** Append-only ledger table (`TransactionHistory`) for every balance change.

### Explanation
Billing system design in Go requires careful consideration of concurrency, precision, idempotency, and auditing. Concurrency is handled with mutexes or database locking to prevent double-spending. Financial precision uses int64 for cents or decimal libraries instead of float64. Idempotency prevents duplicate charges, and audit trails maintain transaction history.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you design a billing system in Go?
**Your Response:** "I design billing systems with careful attention to several critical aspects. For concurrency, I use `sync.Mutex` or database record locking with `SELECT FOR UPDATE` to prevent double-spending scenarios. For precision, I never use float64 for money - instead I use int64 representing cents or the shopspring/decimal library for exact decimal arithmetic. Idempotency is crucial for APIs to prevent duplicate charges, so I implement idempotency keys. For auditing, I maintain an append-only ledger table that records every balance change. This combination ensures financial accuracy, prevents race conditions, maintains data integrity, and provides a complete audit trail for compliance and debugging."

---

### Question 595: How would you scale a notification system written in Go?

**Answer:**
- **Fan-out:** 1 Event ("New Post") -> Produce 1000 messages ("Notify Follower X") to Kafka.
- **Workers:** 50 Go pods consuming Kafka and sending FCM/Email/SMS.
- **Rate Limiting:** Throttle sends per user/provider to avoid bans.

### Explanation
Scaling notification systems involves fan-out patterns where one event generates many messages, distributed processing with multiple worker pods, and rate limiting to prevent provider bans. The fan-out pattern transforms a single event into multiple personalized notification messages, while workers handle the actual delivery with proper throttling.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you scale a notification system written in Go?
**Your Response:** "I scale notification systems using a fan-out pattern with distributed workers. When an event like 'New Post' occurs, I fan it out into hundreds or thousands of individual messages like 'Notify Follower X' and publish them to Kafka. I then run multiple Go pods as workers that consume these messages and handle the actual delivery via FCM, email, or SMS. I implement rate limiting per user and per provider to avoid getting banned by notification services. This approach allows me to scale horizontally by adding more worker pods as needed, while the fan-out pattern ensures each follower gets a personalized notification. Rate limiting protects against overwhelming providers and maintains good delivery rates."

---

### Question 596: How do you build a real-time leaderboard in Go?

**Answer:**
Do not use SQL `ORDER BY`. Use **Redis Sorted Sets (ZSET)**.
- `ZADD leaderboard score user_id`
- `ZREVRANGE leaderboard 0 10 WITHSCORES` (Top 10).
Go service acts as a wrapper around Redis commands.

### Explanation
Real-time leaderboards in Go should use Redis Sorted Sets (ZSET) instead of SQL ORDER BY for performance. ZADD adds users with their scores, and ZREVRANGE retrieves the top players efficiently. The Go service acts as a wrapper around Redis commands, providing a clean API while leveraging Redis's optimized sorted set operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a real-time leaderboard in Go?
**Your Response:** "I build real-time leaderboards using Redis Sorted Sets instead of SQL ORDER BY for performance reasons. I use ZADD to add users with their scores to the leaderboard, and ZREVRANGE to retrieve the top players efficiently. The Go service acts as a wrapper around these Redis commands, providing a clean API while leveraging Redis's optimized sorted set operations. This approach provides O(log N) performance for updates and O(log N + M) for retrieving the top M players, which is much faster than SQL ORDER BY for real-time applications. Redis handles the sorting and ranking automatically, making the implementation simple and highly performant."

---

### Question 597: How would you implement transactional emails in Go?

**Answer:**
Reliability is key.
1.  **DB Transaction:** Write "Order Created" AND "EmailJob" (status=pending) to DB atomically (Outbox Pattern).
2.  **Worker:** Polls "EmailJob", sends via SendGrid/AWS SES.
3.  **On Success:** Delete/Update Job.
Ensures email is sent if and only if Order is committed.

### Explanation
Transactional emails in Go require reliability through the outbox pattern. A database transaction writes both the order and email job atomically. A separate worker polls for pending jobs and sends emails via services like SendGrid or AWS SES. Success results in job deletion/completion. This ensures emails are sent only when orders are successfully committed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you implement transactional emails in Go?
**Your Response:** "I implement transactional emails using the outbox pattern for reliability. First, I write both the order data and an email job record to the database in a single atomic transaction. The email job has a pending status initially. A separate worker polls for pending email jobs and sends them through services like SendGrid or AWS SES. On successful delivery, the worker updates or deletes the job record. This approach ensures that emails are sent if and only if the order is successfully committed to the database, preventing scenarios where customers receive emails for failed orders. The atomic transaction guarantees consistency, while the worker handles the unreliable network communication to email providers."

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

### Explanation
Money and currencies in Go are modeled using a struct with amount as int64 (representing minor units like cents) and currency as a string. This avoids floating-point precision issues. Operations must validate currency matching before performing arithmetic to prevent mixing different currencies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you model money and currencies in Go?
**Your Response:** "I model money and currencies using a struct with the amount as int64 representing minor units like cents, and the currency as a string. This approach avoids the precision issues that come with using float64 for financial calculations. The int64 amount ensures exact arithmetic without rounding errors. Before performing any addition or subtraction operations, I always check that the currencies match to prevent mixing different currencies. This pattern provides type safety, precision, and clear domain modeling for financial operations. It's a simple but effective approach that handles the complexities of monetary calculations in a reliable way."

---

### Question 599: How do you do dependency injection in Go?

**Answer:**
- **Manual (Idiomatic):** Pass dependencies to constructors.
    ```go
    func NewService(db *sql.DB, logger Logger) *Service { ... }
    ```
- **Libraries:** `google/wire` (Compile time code-gen) or `uber-go/dig` (Reflection based). Go prefers manual/wire (explicit) over "Magic" containers.

### Explanation
Dependency injection in Go can be done manually by passing dependencies to constructors, which is the idiomatic approach. Libraries like google/wire provide compile-time code generation, while uber-go/dig uses reflection. The Go community generally prefers manual or wire-based approaches for their explicitness over magic containers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you do dependency injection in Go?
**Your Response:** "I do dependency injection in Go primarily using the manual approach by passing dependencies to constructors like `NewService(db *sql.DB, logger Logger)`. This is the idiomatic Go way - explicit and easy to understand. For more complex scenarios, I might use `google/wire` which provides compile-time code generation, making dependency relationships clear at build time. I generally avoid reflection-based containers like `uber-go/dig` because the Go community prefers explicit approaches over 'magic' containers. The manual approach works well for most applications and keeps the dependency graph visible and maintainable. Wire is great for larger projects where manual wiring becomes cumbersome, but still maintains compile-time safety."

---

### Question 600: How do you create a rule engine in Go?

**Answer:**
1.  **Interface:** `type Rule interface { Evaluate(Context) bool }`.
2.  **Composite:** `type AndRule struct { rules []Rule }`.
3.  **DSL (Advanced):** Use `antonmedv/expr` or `google/cel-go` (Common Expression Language) to parse string rules (`user.age > 18 && user.premium`) and evaluate them safely at runtime.

### Explanation
Rule engines in Go can be implemented with a Rule interface for evaluation, composite patterns for complex rules, and DSL libraries for string-based rules. The interface approach allows type-safe rule evaluation, while DSL libraries like expr or cel-go enable parsing and evaluating string expressions safely at runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create a rule engine in Go?
**Your Response:** "I create rule engines using several approaches. First, I define a Rule interface with an Evaluate method that takes a Context and returns a boolean. For complex rules, I implement composite patterns like AndRule that combines multiple rules. For more advanced scenarios, I use DSL libraries like `antonmedv/expr` or `google/cel-go` which allow me to parse string expressions like 'user.age > 18 && user.premium' and evaluate them safely at runtime. This combination provides flexibility - simple rules use the interface approach, complex business logic uses composites, and dynamic rules use DSL parsing. The DSL approach is particularly powerful when rules need to be stored in databases or configured by non-developers, while maintaining type safety and security in evaluation."

---
