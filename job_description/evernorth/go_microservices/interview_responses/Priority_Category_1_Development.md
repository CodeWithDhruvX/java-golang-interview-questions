# Priority Category 1: Go Microservices Development

### 1. How do you implement RESTful APIs in Go for microservices?
**Your Response:**
"When I build REST APIs in Go, I prioritize simplicity and performance. I typically choose a lightweight router like `chi` because it’s 100% compatible with `net/http`. My architecture is strictly layered:
- **Handlers**: Responsible for request parsing and response formatting.
- **Service Layer**: Where the actual business logic lives, kept completely decoupled from HTTP concerns.
- **Repository**: Handles data persistence.

I use `context.Context` for every single request to handle timeouts and cancellations. For request validation, I prefer using a custom validation layer rather than generic reflection-heavy libraries to keep execution speed high."

### 2. Best practices for structuring Go microservices?
**Your Response:**
"I follow the **Standard Go Project Layout**. 
- `/cmd`: Contains the entry points for the various binaries.
- `/internal`: This is crucial for microservices. It ensures that packages inside can't be imported by other services, which prevents accidental tight coupling and 'spaghetti' cross-service dependencies.
- `/pkg`: For code that *is* intentionally shared.

I also stick to 'One service, one repo' to allow for independent scaling and deployment cycles, and I use dependency injection (manual or with `fx`) to make the components easily testable."

### 3. How do you handle inter-service communication in Go?
**Your Response:**
"It's about choosing the right protocol for the requirement.
For synchronous, high-performance communication, I use **gRPC**. The binary serialization with Protobuf is significantly faster than JSON, and the generated code ensures both services are speaking the same language.
For asynchronous, event-driven flows—like updating a search index after a user is created—I use **NATS or Kafka**. This decouples the services and ensures the system remains resilient even if one service is temporarily down."

### 4. Explain the Singleton pattern in Go for database connections.
**Your Response:**
"In Go, the cleanest way to implement a Singleton is using `sync.Once`. It guarantees that a piece of code runs exactly once, making it thread-safe for initializing a global connection pool. 
However, I always remind interviewers that `sql.DB` is already a connection pool. So the 'Singleton' here isn't about a single connection, but a single *manager* that controls the pool, ensuring we don't accidentally open hundreds of pool instances during startup."

### 5. How do you implement Circuit Breaker pattern in Go?
**Your Response:**
"I use the `gobreaker` library. I wrap any external network call in a circuit breaker. If the downstream service starts failing beyond a certain threshold (say 20% failure rate), the breaker trips to 'Open.'
This prevents my service from wasting resources on calls that are guaranteed to fail and stops the 'cascading failure' effect across the entire system. Once the service recovers, the breaker moves to 'Half-Open' to test the waters before fully closing again."

### 6. What are Go modules and how do you manage dependencies in microservices?
**Your Response:**
"Go modules are the standard for dependency management. In a microservices environment, I'm very careful with versioning. I use `go.mod` and `go.sum` to lock dependencies. 
For internal shared libraries, I use private Git repositories and set the `GOPRIVATE` environment variable. To keep builds stable, I always use a vendor directory or a proxy like `Athens` to ensure we aren't dependent on external repos staying online."

### 7. How do you implement retry mechanisms in Go services?
**Your Response:**
"I implement retries using **Exponential Backoff with Jitter**. I'll have a loop that tries the operation, and if it fails with a retriable error (like a network timeout), it sleeps for a duration that doubles each time.
The 'jitter' is the secret sauce—by adding a random delay, I prevent a 'thundering herd' where all my service instances retry at the exact same millisecond and crush the already-struggling downstream service."

### 8. Explain connection pooling in Go for databases.
**Your Response:**
"Go's `database/sql` handles the heavy lifting, but the configuration is where you prove your expertise. I always tune three parameters:
- `SetMaxOpenConns`: To limit the pressure on the DB.
- `SetMaxIdleConns`: To keep ready-to-use connections available without wasting memory.
- `SetConnMaxLifetime`: Crucial for preventing 'stale' connections that might have been closed by the DB or a load balancer. 
I base these numbers on the DB's capacity and the expected peak traffic of the specific service."
