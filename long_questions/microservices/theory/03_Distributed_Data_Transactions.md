# 🟢 **41–60: Distributed Data & Transactions**

### 41. What is distributed transaction?
"A distributed transaction is a database transaction in which two or more network hosts are involved. In microservices, because of the database-per-service pattern, a single business operation often spans multiple databases.

For instance, processing an e-commerce order requires deducting stock from the Inventory database and charging the card in the Payment database. Both must succeed, or both must fail. 

Managing this 'all-or-nothing' atomicity across multiple distinct services separated by a network is one of the hardest problems in distributed systems design, demanding specific architectural patterns."

#### Indepth
In traditional monoliths, an ACID-compliant relational DB handles atomicity using a local transaction manager. The transition to distributed systems requires giving up the ease of local ACID guarantees and transitioning to BASE (Basically Available, Soft state, Eventual consistency) models, drastically increasing application logic complexity.

---

### 42. Why 2PC is problematic?
"Two-Phase Commit (2PC) is a traditional protocol for distributed transactions. It operates in two phases: a 'Prepare' phase where a coordinator asks all databases if they are ready, and a 'Commit' phase where it tells them to complete the change.

I avoid 2PC in microservices because it is synchronous and heavily blocking. If the Inventory DB responds 'Prepared' but the network drops before it gets the 'Commit' signal from the coordinator, the Inventory DB locks those rows completely, preventing any other transactions from touching them.

Furthermore, it couples microservices at a deep protocol level and scales terribly under high throughput."

#### Indepth
2PC represents the "Consistency" focus of the CAP theorem. It sacrifices Availability because the entire transaction halts if a single participant fails or the coordinator crashes. Modern cloud architectures prioritize Availability, making 2PC an anti-pattern for internet-scale microservices.

---

### 43. What is Saga pattern?
"The Saga pattern is the modern alternative to 2PC for handling distributed transactions. Instead of one giant, blocking ACID transaction, a Saga breaks it down into a sequence of smaller, local ACID database transactions.

Each local transaction updates the database and publishes a message/event to trigger the next local transaction in the chain. 

If any step fails (e.g., Payment succeeds but Inventory is out of stock), the Saga executes 'compensating transactions' backward to undo the preceding steps (e.g., refund the Payment)."

#### Indepth
A Saga is not truly atomic in an ACID sense—it is Eventually Consistent. During a saga execution, external observers might read intermediate states (an order is marked 'Pending Payment' while stock is already reserved). Therefore, Sagas must fundamentally embrace semantic locking or intermediate state flags in the domain model.

---

### 44. Saga choreography vs orchestration?
"In **Choreography**, there is no central controller. Services publish events (e.g., Kafka topics) and other services listen and react. It's like dancers moving to the music. The Order service publishes 'Order Created', Payment hears it, charges the card, and publishes 'Payment Succeeded'. It's highly decoupled but hard to track the overall status.

In **Orchestration**, a central 'Orchestrator' (like Cadence, Netflix Conductor, or a simple state machine Service) commands the others. It tells Payment to charge, waits for a response, then tells Inventory to deduct stock. It's like a conductor leading an orchestra.

I prefer Orchestration for complex workflows (more than 3-4 steps) because the state machine is explicit and easy to monitor, whereas Choreography can turn into untraceable 'event spaghetti'."

#### Indepth
Choreography forces the domain logic into the messaging topology itself, which becomes an architectural bottleneck. Orchestration isolates the workflow logic into a specific service, allowing the participating microservices to remain happily ignorant of the broader business process they are participating in.

---

### 45. What is compensation transaction?
"A compensating transaction is the mechanism used in a Saga to undo a previously committed local transaction.

Because I cannot 'rollback' a database transaction that happened five minutes ago in another microservice's database, I must author a new transaction that theoretically reverses the effect. 

If Step 1 charged $100 to a credit card, the compensating transaction for Step 1 is to issue a $100 refund. If Step 2 reserved a hotel room, the compensating transaction cancels the reservation."

#### Indepth
Compensating transactions are fundamentally application-level code, not database-level ROLLBACK commands. They must be inherently idempotent, because a network failure might cause the orchestrator to fire the compensation request multiple times. If the compensation fails repeatedly, developers often must rely on manual intervention (dead-letter reporting).

---

### 46. What is CQRS?
"CQRS stands for Command and Query Responsibility Segregation. It separates the model used to update data (Commands) from the model used to read data (Queries).

In a microservice, writing a new user profile might involve complex validations and inserting data into a PostgreSQL DB. Reading that profile might require massive join operations that are slow.

CQRS splits these. I'll have a write API that updates PostgreSQL. I'll then asynchronously push that new data to a completely different read-optimized database (like Elasticsearch), which the read API queries directly. This scales reads and writes independently."

#### Indepth
CQRS significantly improves performance but permanently introduces eventual consistency between the Command side and Query side. It is the architectural manifestation of the CQS (Command Query Separation) principle popularized by Bertrand Meyer, scaled to a database level.

---

### 47. What is event sourcing?
"Event Sourcing stores the state of a system as an immutable sequence of state-changing events, rather than just storing the current state. 

Instead of an `Accounts` table that just shows `Balance: 500`, the database stores a log: `Account Created(0) -> Deposited(1000) -> Withdrew(500)`. 

To get the current balance, the application loads all past events and replays them. I heavily use this combined with CQRS because it provides a 100% accurate audit trail—crucial for financial subdomains—and lets me 'time-travel' to see what the system state was yesterday at 5 PM."

#### Indepth
Event Sourcing radically alters database operations. Updates and Deletes do not exist—everything is an Insert (Append-only). While this provides phenomenal write performance, reading a million events to derive a balance is impossible; therefore, "Snapshots" are periodically saved to capture state at specific checkpoints.

---

### 48. How do you handle duplicate events?
"Duplicate events happen constantly in distributed systems, usually due to a broker delivering a message multiple times ('at-least-once delivery') when a consumer crashes before acknowledging it.

I handle this strictly by building **Idempotent Consumers**. 

When a microservice receives an event, it first checks a deduplication table (often in Redis or the primary database) to see if the `eventId` was already processed. If it was, the service safely ignores it and sends back an acknowledgment to the broker."

#### Indepth
Deduplication tables in relational databases can utilize unique constraints. For example, inserting a record representing the processed `eventId`; if a duplicate event arrives, the insert will violate the unique key constraint, safely throwing a duplicate key exception which the app catches and ignores.

---

### 49. How do you ensure exactly-once delivery?
"In distributed systems, true 'exactly-once delivery' over a network is technically impossible without complex, heavy coordination. The network guarantees are either 'at-most-once' (messages can be lost) or 'at-least-once' (messages can be duplicated).

The industry standard approach is exactly what I implement: I configure the message broker for **At-Least-Once delivery** to ensure no data is lost, and I configure my receiving microservice as an **Idempotent Receiver** to filter out the duplicates.

This combination of 'Guaranteed Delivery + Idempotency' perfectly simulates 'Exactly-Once Processing' from the business's perspective."

#### Indepth
Some streaming platforms like Kafka offer an "exactly-once semantics" (EOS) feature. This utilizes transactional producers and consumer offsets tied into the same transaction scope. However, this only guarantees exactly-once processing *within* Kafka's boundaries. If the side-effect involves writing to an external database, application-level idempotency is still required.

---

### 50. What is dead-letter queue?
"A Dead-Letter Queue (DLQ) is a secondary message queue used to temporarily store messages that a consumer cannot process successfully.

If my Order service receives a corrupt JSON payload from a RabbitMQ topic, parsing it fails. The service will retry 3 times. If it still fails, it doesn't want to block the queue endlessly. It drops the corrupt message into a dedicated DLQ and moves on to the next valid message.

I monitor the DLQ for alerts. An engineer can inspect the 'dead' messages, fix the codebase bug, and replay the messages from the DLQ back into the main queue."

#### Indepth
Messages can be DLQ'd for multiple reasons: unparseable payloads (Poison Pills), exceeding maximum retry thresholds, message TTL expiration, or queue size limits. DLQs are an essential safety net for asynchronous edge cases.

---

### 51. What is message ordering?
"Message ordering is the guarantee that messages published by a producer to a topic are processed by a consumer in the exact sequential order they were sent.

In Kafka, ordering is only guaranteed *within a single partition*. 

If I'm streaming stock market trades for 'AAPL' and 'TSLA', I use the stock ticker as the partitioning key. All 'AAPL' trades go to Partition 1, and all 'TSLA' trades go to Partition 2. This guarantees the consumer processes 'AAPL Buy 1' before 'AAPL Sell 2', preventing massive timeline data corruption."

#### Indepth
If strict global ordering across *all* messages is required, you must route all messages through a single partition with a single consumer instance. This introduces a massive performance bottleneck, crippling horizontal scalability, and is almost always indicative of a bad architectural design.

---

### 52. What is data partitioning?
"Data partitioning is the process of splitting a large, monolithic dataset into smaller, distinct logical segments—partitions.

In microservices, we partition data by bound context (e.g., Orders vs. Customers). Within a single microservice, if the database gets too large, we horizontally partition tables (e.g., separating active orders from historical orders).

I partition data primarily to limit query scan boundaries. Instead of a database scanning 100 million rows, it only scans the 10 million rows in the relevant partition, drastically reducing I/O latency."

#### Indepth
Partitions can be range-based (e.g., Partition 1: Jan-March, Partition 2: April-June) or hash-based (Partition = hash(userId) % N). Effective partitioning distributes read/write load evenly, avoiding "hot spots" where one partition handles 90% of the traffic.

---

### 53. What is sharding?
"Sharding is a specific type of data partitioning where the data is horizontally split across completely separate, physical database servers (shards), rather than just logical tables in one server.

If a single PostgreSQL server runs out of disk space or CPU to handle my growing user base, I shard the database. Users A-M go to Database Server 1, and Users N-Z go to Database Server 2.

I implement this at extreme scale because it allows the database infrastructure to scale horizontally infinitely, but it severely complicates cross-shard queries and transaction management."

#### Indepth
Sharding introduces the complexity of the "Routing Layer". The application must now know *which* database server holds the `userId`. Resharding—moving data when a new server is added—is mathematically complex (often relying on Consistent Hashing) and operationally terrifying without downtime.

---

### 54. What is replication?
"Replication involves copying real-time data from a primary database node to one or more secondary (replica) nodes.

If a primary database crashes, it's an immediate outage. With replication, a secondary node holds an identical copy of the data. The infrastructure detects the crash, promotes the secondary to primary, and restores service within seconds.

I treat replication as mandatory for high availability and disaster recovery in production systems. It protects against hardware failures and data center outages."

#### Indepth
Replication can be Synchronous or Asynchronous. Synchronous guarantees zero data loss but blocks the primary's write until the replica acknowledges it (horrible for latency). Asynchronous is blazingly fast but risks losing the few milliseconds of data that were flighting during a primary crash.

---

### 55. What is read replica?
"A Read Replica is a synchronized copy of the primary database that is strictly used to serve read-only queries.

If my microservice receives 1,000 requests per second, and 90% of them are `GET` requests, pushing all of that traffic to the Primary database causes CPU contention that slows down vital `POST` (write) requests.

I route all write operations to the Primary, and I load-balance all read operations across three Read Replicas. This radically improves read-latency and offloads critical strain from the Primary server."

#### Indepth
Because data is replicated asynchronously, Read Replicas exhibit **Eventual Consistency**. A user might update their profile (Primary) and refresh the page (Replica) within 10 milliseconds, yet see an old version of their profile. Modern frontend code masks this by updating local state optimistically before receiving a server response.

---

### 56. How to handle schema evolution?
"Schema evolution is the process of modifying database structures (adding/dropping columns) without causing downtime or breaking the API contracts of the microservices reading from them.

In a deployment pipeline, I never let code updates and schema updates happen in an incompatible way. 

I follow the 'Expand and Contract' pattern. 
Step 1: Release a schema migration to *add* the new column (`fullName`). 
Step 2: Update the microservice code to write to both the old (`firstName`/`lastName`) and new columns. 
Step 3: Migrate old rows. 
Step 4: Update the code to only use `fullName`. 
Step 5: Drop the old columns months later."

#### Indepth
Tight coupling between code deployment and schema migration leads to deployment failure. If the code relies on a column that hasn't been created yet, or a column is renamed while older app pods are still servicing requests, the system crashes. Utilizing tools like Flyway or Liquibase helps track and automate these backward-compatible transitions.

---

### 57. How to manage migrations?
"Database migrations must be version-controlled, automated, and deterministic.

In a Spring Boot environment, I strictly use Flyway or Liquibase. I write SQL scripts (e.g., `V1__init.sql`, `V2__add_email.sql`) and place them in the application repository.

When the microservice container boots up in CI/CD or Production, Flyway checks a metadata table in the database to see the current version, and automatically executes any pending scripts before the Spring context finishes loading. This guarantees the database schema is always correctly synced with the code version."

#### Indepth
Because multiple pods of a microservice might spin up simultaneously during a Kubernetes rollout, migration tools must utilize database locks (like a row lock on a `flyway_schema_history` table) to prevent three pods from attempting to run `V2__add_email.sql` concurrently.

---

### 58. What is optimistic locking?
"Optimistic locking prevents the 'lost update' anomaly in concurrent systems without relying on the database to physically lock rows, which kills performance.

I implement it by adding a `@Version` integer column to an entity. When User A and User B read the same record, they both get `version=1`. 

If User A updates the record and hits save, the database executes `UPDATE ... WHERE id=X AND version=1`. The save works, and `version` becomes 2. When User B hits save instantly after, their query is `UPDATE ... WHERE id=X AND version=1`. Because `version` is now 2, the DB updates 0 rows. User B threw an `OptimisticLockingFailureException` and is prompted to refresh."

#### Indepth
Optimistic locking is ideal for read-heavy systems with low probability of collision (like a user updating their personal profile). It acts as a safety net against dirty overwrites without holding long database transactions.

---

### 59. What is pessimistic locking?
"Pessimistic locking acts on the assumption that conflicting updates *will* happen frequently, so it prevents them by placing a hard, physical lock on database rows.

When User A reads a record, they issue a `SELECT ... FOR UPDATE` query. The database physically locks that row. If User B tries to read or update that same row, User B's thread completely blocks entirely until User A finishes their transaction.

I only use this in extremely high-contention, high-consequence scenarios—like calculating the final allocation of remaining seats in a ticketing system or moving money between ledgers."

#### Indepth
Pessimistic locking severely degrades system throughput. In a bad scenario, it leads to deadlocks—Transaction A locks Row 1 and waits for Row 2, while Transaction B locks Row 2 and waits for Row 1. Databases resolve deadlocks by terminating one transaction arbitrarily to free the resources.

---

### 60. What is BASE property?
"BASE is the consistency model adopted by distributed NoSQL databases and microservices, acting as the counter-philosophy to relational ACID properties.

- **B**asically **A**vailable: The system guarantees availability, responding to requests even if entire network partitions fail.
- **S**oft State: The state of the system is volatile and might change over time without explicit input, purely due to eventual consistency synchronizations.
- **E**ventually Consistent: Given time, once data stops flowing, all nodes will agree on a consistent state.

I embrace BASE in microservices architecture because web-scale traffic demands the high availability and horizontal scaling that strict ACID transactions cannot mathematically provide under the CAP theorem."

#### Indepth
ACID aims for Safety at the cost of Availability. BASE aims for Availability at the cost of transient accuracy. Most modern cloud workloads process non-critical data (like 'Likes' on a post or item recommendations), rendering BASE a perfectly acceptable trade-off for performance.
