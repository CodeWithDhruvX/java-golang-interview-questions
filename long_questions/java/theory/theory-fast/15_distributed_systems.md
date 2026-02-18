# Distributed Systems Interview Questions (154-160)

## Consistency & Distributed Patterns

### 154. What is CAP theorem?
"It states that in a distributed system, you can only pick two: **Consistency**, **Availability**, or **Partition Tolerance**.

But in reality, Partition Tolerance (P) is mandatory—networks *will* fail. So the real choice is CP vs AP.

**CP (Consistent & Partition Tolerant)**: If the network breaks, we refuse writes to prevent inconsistency (like a bank).
**AP (Available & Partition Tolerant)**: If the network breaks, we accept writes on all nodes, even if they drift out of sync (like a social media feed). We fix the conflict later."

### 155. Difference between strong and eventual consistency?
"**Strong Consistency**: Once a write is confirmed, *any* subsequent read will see that new value. It simplifies development but hurts latency (you have to wait for replication).

**Eventual Consistency**: If I write a value, other nodes might see the old value for a few milliseconds (or seconds). But *eventually*, if no new updates occur, all nodes will converge to the same value. This allows high availability and low latency (like DNS or Cassandra)."

### 156. What is idempotency in distributed systems?
"Idempotency guarantees that performing an operation multiple times has the same result as performing it once.

This is critical because in a distributed system, networks are unreliable. If I send a `Deduct $50` request and don't get a response, I have to retry.

If the operation isn't idempotent, I might deduct $100. If it is, the second request is recognized as a duplicate and ignored, or returns the cached success response."

### 157. How do you design idempotent APIs?
"For `GET`, `PUT`, and `DELETE`, HTTP semantics already require idempotency.

For `POST` (like logical creation or payment), I require the client to send a unique **Idempotency Key** (a UUID) in the header.

On the server, I check a Redis cache or a unique database constraint.
'Have I seen this UUID before?'
-   Yes -> Return the previous successful response immediately.
-   No -> Process the request and save the UUID."

### 158. What is saga pattern?
"Saga is a pattern for managing long-running transactions distributed across multiple services. Since we can't use ACID across microservices, we break the transaction into a sequence of smaller local transactions.

If one step fails (e.g., 'Inventory Reserved', but 'Payment Failed'), the Saga executes **Compensating Transactions** to undo the previous steps (e.g., 'Release Inventory').

This ensures data consistency without locking resources across the entire system."

### 159. Two-phase commit vs saga?
"**Two-Phase Commit (2PC)** is a strong consistency protocol. A coordinator tells all databases to 'Prepare', and if all say yes, it says 'Commit'. It guarantees ACID but is slow and holds locks on all databases during the process. It doesn't scale well.

**Saga** is an eventually consistent approach. It relies on asynchronous messaging and compensation. It’s harder to debug (because of the complex state machine) but much more scalable and resilient to partial failures."

### 160. How do you handle duplicate messages?
"In messaging systems (Kafka/RabbitMQ), 'exactly-once' delivery is extremely hard. Most systems guarantee 'at-least-once', meaning you *will* get duplicates.

So the consumer *must* be idempotent.

I usually track processed Message IDs in a database table. Before processing a message, I verify: `INSERT IGNORE INTO processed_messages (msg_id)`. If it affects 0 rows, I know it's a duplicate and I ack it without doing any work."
