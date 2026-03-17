# Priority Category 2: System Design & Integration

### 1. Design a scalable microservice architecture using Go.
**Your Response:**
"A truly scalable Go system is **stateless**. I design my Go services to be 'disposable' pods.
1. **Entry Point**: A high-performance load balancer or API Gateway (like Nginx or NATS).
2. **Compute**: Go services running in Kubernetes, using HPA (Horizontal Pod Autoscaling) to scale based on CPU.
3. **Communication**: gRPC for internal calls, Kafka for async events.
4. **State**: Offloaded to a managed database (Postgres) and a distributed cache (Redis).
By keeping the Go logic stateless, I can scale from 1 to 100 pods in seconds without worrying about data corruption."

### 2. How do you implement database per service pattern in Go?
**Your Response:**
"This is about technical and organizational boundaries. Each Go service has its own dedicated database instance or schema.
No service is allowed to 'reach' into another service's database. If Service A needs Service B's data, it *must* ask Service B via an API or listen to its events. 
In Go, I enforce this by only giving each service the credentials for its specific database. It forces you to think about **Service Contracts**, which is the only way to maintain a clean architecture at scale."

### 3. Explain distributed transactions in Go microservices.
**Your Response:**
"In microservices, 'Distributed Transactions' (like 2PC) are generally a bad idea because they kill performance and reliability. Instead, I use the **Saga Pattern**.
It’s a sequence of local transactions. For example: Service A reserves an item, Service B charges the card. If Service B fails, the system emits a 'Compensating Event' that tells Service A to un-reserve the item. It’s all about **Eventual Consistency**."

### 4. How do you handle data consistency across services?
**Your Response:**
"I use **Eventual Consistency**. I don't try to make the whole system consistent in one millisecond. 
Instead, I use an event-driven approach (often called 'Outbox Pattern'). When a service updates its database, it also writes an event to an 'outbox' table in the same transaction. A separate process then publishes those events. This ensures that the downstream services eventually get the update, and we never lose an event if the message broker is down."

### 5. Design patterns for resilient Go services.
**Your Response:**
"Resilience isn't an accident; it's a design choice. I implement:
- **Timeouts**: Every network call has a hard deadline via `context.WithTimeout`.
- **Circuit Breakers**: To stop calling dying services.
- **Bulkheads**: Isolating resources so if one 'heavy' feature fails, it doesn't sink the whole service.
- **Retries**: For transient issues, always with backoff.
In Go, these are easy to implement using middleware and the standard `context` package."

### 6. How do you integrate multiple databases (Oracle, MongoDB, etc.) in Go?
**Your Response:**
"I use the **Repository Pattern**. I define a Go interface for each entity, like `UserRepository`. 
Then I create separate implementations: `OracleUserRepository` and `MongoUserRepository`. 
During startup, I initialize the correct one based on configuration. This keeps the 'Higher Level' business logic completely unaware of which database is actually sitting underneath. It makes swapping databases—or even just mocking them for unit tests—incredibly simple."
