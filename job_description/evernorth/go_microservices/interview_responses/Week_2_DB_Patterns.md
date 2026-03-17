# Week 2: Database Integration & Design Patterns

### 🔹 Topic: Multi-Database Integration

**Interviewer:** "How do you manage connections to multiple databases like Postgres and MongoDB in a single Go service?"

**Your Response:**
"I use a 'Repository Pattern' to abstract the data layer. I'll have an interface for each 'store' (e.g., `UserRepository`). Then I create concrete implementations—one for Postgres using `sqlx` or `gorm`, and another for MongoDB using the official driver. 

For connection management, I use the **Singleton pattern** for the database clients to ensure we aren't creating new connections on every request. I also make sure to configure the `SetMaxOpenConns` and `SetMaxIdleConns` for SQL databases to handle pooling efficiently and avoid hitting the DB's connection limit."

### 🔹 Interview Focus: Critical Questions

**1. How do you implement efficient connection pooling in Go?**
**Your Response:** "Go's `database/sql` handles pooling out of the box. The key is to initialize the `*sql.DB` object once and reuse it. I tune it by setting `SetConnMaxLifetime` to prevent using stale connections and adjusting the idle/open connection limits based on the expected load and the database capacity."

**2. Explain the 'database per service' pattern.**
**Your Response:** "It means each microservice has its own private database that no other service can access directly. If Service A needs Service B's data, it must go through Service B's API. This ensures loose coupling and allows each service to use the database best suited for its needs—like Postgres for relational data and Redis for caching."

**3. How do you handle distributed transactions?**
**Your Response:** "I try to avoid them if possible because they are complex and don't scale well. Instead, I use the **Saga Pattern**. I break the transaction into a series of local transactions. Each step has a 'compensating transaction' that rolls back the previous steps if a later one fails. This maintains eventual consistency without locking resources across services."

**4. Best practices for database connections in Go?**
**Your Response:** "Always use environment variables for credentials, implement health checks with `db.Ping()`, use context with timeouts for all queries to avoid hanging requests, and always `defer rows.Close()` to prevent memory leaks."

**5. Integrating SQL vs NoSQL?**
**Your Response:** "For SQL, I stick to `sqlx` or `gorm` for mapping. For NoSQL like Mongo, I use the BSON package and the official driver. I keep the business logic database-agnostic by using interfaces, so switching from a mock DB to a real one during testing is seamless."

### 🔹 Week 2 Practice Problems: Spoken Walkthroughs

**1. Multi-database connection pool manager:**
"I'd create a `Registry` struct that holds pointers to various DB connections. On startup, I initialize them all. I'd provide a `GetDB(name string)` method that returns the requested connection, ensuring it's thread-safe."

**2. Distributed transaction coordinator:**
"I'd implement a simple Saga Orchestrator. It would have a list of steps. For each step, it executes a function. If any step fails, it iterates backwards through the already completed steps and calls their compensation functions."

**3. Retry mechanism with exponential backoff:**
"I'd use a `for` loop with a `time.Sleep`. I'd start with a base delay (e.g., 100ms) and double it after each failure, up to a maximum limit. I'd also add 'jitter' (a random offset) to prevent all services from retrying at the exact same time and overwhelming the DB."

**4. Database abstraction layer:**
"I'd define a `Store` interface with methods like `SaveUser` or `FindProduct`. This allows the service layer to interact with data without knowing whether it's coming from MySQL, MongoDB, or an in-memory map for unit tests."
