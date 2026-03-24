# Distributed Systems Interview Questions (154-160)

## Consistency & Distributed Patterns

### 154. What is CAP theorem?
"It states that in a distributed system, you can only pick two: **Consistency**, **Availability**, or **Partition Tolerance**.

But in reality, Partition Tolerance (P) is mandatory—networks *will* fail. So the real choice is CP vs AP.

**CP (Consistent & Partition Tolerant)**: If the network breaks, we refuse writes to prevent inconsistency (like a bank).
**AP (Available & Partition Tolerant)**: If the network breaks, we accept writes on all nodes, even if they drift out of sync (like a social media feed). We fix the conflict later."

**Spoken Format:**
"The CAP theorem is like understanding that you can't have everything in a distributed system - you have to choose what's most important.

Imagine a banking system across multiple branches:

**Consistency** is like ensuring every branch shows the same account balance - perfect accuracy but slow because every branch must coordinate.

**Availability** is like ensuring every branch can serve customers even if the network connection to head office is down - customers can still deposit money.

**Partition Tolerance** is like accepting that sometimes branches can't communicate with head office - but they can still serve customers using their local records.

**CP (Consistent)** chooses consistency over availability - if network issues occur, branches stop serving customers until communication is restored.

**AP (Available)** chooses availability over consistency - branches continue serving customers even with temporary data inconsistencies, which are fixed later.

The choice depends on your business needs: banking needs consistency, social media needs availability. You can't have both perfectly!"

### 155. Difference between strong and eventual consistency?
"**Strong Consistency**: Once a write is confirmed, *any* subsequent read will see that new value. It simplifies development but hurts latency (you have to wait for replication).

**Eventual Consistency**: If I write a value, other nodes might see the old value for a few milliseconds (or seconds). But *eventually*, if no new updates occur, all nodes will converge to the same value. This allows high availability and low latency (like DNS or Cassandra)."

**Spoken Format:**
"Strong vs eventual consistency is like choosing between perfect accuracy and perfect availability.

**Strong consistency** is like having a group project where everyone must agree before moving to the next step. It's perfectly accurate but slow because everyone must coordinate.

**Eventual consistency** is like having a group project where people can start working immediately and sync up later. It's fast and available but temporarily inconsistent.

The trade-offs:

- **Strong**: Perfect accuracy, slower performance, less available during network issues
- **Eventual**: High availability, fast performance, temporary inconsistencies

Choose based on business needs:
- Banking: Strong consistency (can't have wrong account balances)
- Social media: Eventual consistency (likes and comments can sync later)
- DNS: Eventual consistency (users can still access domains even during updates)

The key is that 'eventually' means within milliseconds, not days!"

### 156. What is idempotency in distributed systems?
"Idempotency guarantees that performing an operation multiple times has the same result as performing it once.

This is critical because in a distributed system, networks are unreliable. If I send a `Deduct $50` request and don't get a response, I have to retry.

If the operation isn't idempotent, I might deduct $100. If it is, the second request is recognized as a duplicate and ignored, or returns the cached success response."

**Spoken Format:**
"Idempotency is like having a magic vending machine that prevents double-charging.

Imagine you insert a card to pay for parking. If the machine malfunctions and doesn't respond, you try again.

**Idempotent operation**: The second time, the machine recognizes you already paid and returns your money immediately.

**Non-idempotent operation**: The second time, the machine charges you again because it doesn't remember you already paid.

In distributed systems, this is critical because:
- Network failures happen
- Timeouts occur
- Messages get duplicated
- Retries happen

For APIs, I implement idempotency using:
- Unique request IDs
- Database constraints
- Checking if operation already completed

This ensures that retrying or network issues don't result in double-charging or duplicate operations!"

### 157. How do you design idempotent APIs?
"For `GET`, `PUT`, and `DELETE`, HTTP semantics already require idempotency.

For `POST` (like logical creation or payment), I require the client to send a unique **Idempotency Key** (a UUID) in the header.

On the server, I check a Redis cache or a unique database constraint.
'Have I seen this UUID before?'
-   Yes -> Return the previous successful response immediately.
-   No -> Process the request and save the UUID."

**Spoken Format:**
"Designing idempotent APIs is like creating a smart door that only lets in unique visitors.

Imagine a payment API that needs to ensure no duplicate payments:

**Idempotency Key**: Client sends a unique ID (UUID) with each request.

**Server-side check**: Before processing, check if the ID has been seen before using a Redis cache or database constraint.

**If duplicate**: Return the previous successful response immediately - no need to reprocess.

**If new**: Process the request and save the ID for future checks.

This approach ensures that:
- No duplicate payments
- No wasted resources
- System remains consistent even with retries
- You can handle 'at-least-once' delivery guarantees gracefully

It's like having a bouncer who remembers every visitor!"

### 158. What is saga pattern?
"Saga is a pattern for managing long-running transactions distributed across multiple services. Since we can't use ACID across microservices, we break the transaction into a sequence of smaller local transactions.

If one step fails (e.g., 'Inventory Reserved', but 'Payment Failed'), Saga executes **Compensating Transactions** to undo the previous steps (e.g., 'Release Inventory'). This ensures data consistency without locking resources across the entire system."

**Spoken Format:**
"Saga pattern is like having a smart project manager who can undo mistakes.

Imagine a complex online order involving multiple services:
1. Reserve inventory
2. Process payment
3. Update order status

If step 2 fails (payment fails), you have a problem: inventory was reserved but payment wasn't completed.

**Traditional approach**: Lock everything until payment completes - poor performance and availability.

**Saga approach**: Each service has local transactions and compensation logic:
- If payment fails, payment service triggers refund
- Inventory service releases the reserved items
- Order service updates status to failed

This ensures that even if parts fail, the system ends up in a consistent state. It's like having an automatic backup plan for complex multi-step operations!"

### 159. Two-phase commit vs saga?
"**Two-Phase Commit (2PC)** is a strong consistency protocol. A coordinator tells all databases to 'Prepare', and if all say yes, it says 'Commit'. It guarantees ACID but is slow and holds locks on all databases during the process. It doesn't scale well.

**Saga** is an eventually consistent approach. It relies on asynchronous messaging and compensation. It's harder to debug (because of complex state machine) but much more scalable and resilient to partial failures."

**Spoken Format:**
"Two-phase commit vs Saga is like choosing between a strict coordinator and a flexible project manager.

**Two-Phase Commit** is like having a strict manager who requires everyone to agree before proceeding:
1. Manager asks all departments: 'Can you complete your part?'
2. All departments must respond: 'Yes'
3. Only then does manager say: 'Everyone commit!'
- Perfect consistency but slow and doesn't scale well

**Saga** is like having a flexible project manager who allows departments to work independently:
1. Each department does its work and reports back
2. If something fails, department executes its own backup plan
3. Project manager coordinates the overall result
- Eventually consistent but much more scalable and resilient

The choice depends on your needs:
- Banking: Two-phase commit (consistency is critical)
- E-commerce: Saga (availability and scalability are more important)
- Social media: Saga (massive scale and partial failures are common)"

### 160. How do you handle duplicate messages?
"In messaging systems (Kafka/RabbitMQ), 'exactly-once' delivery is extremely hard. Most systems guarantee 'at-least-once', meaning you *will* get duplicates.

So the consumer *must* be idempotent. I usually track processed Message IDs in a database table. Before processing a message, I verify: `INSERT IGNORE INTO processed_messages (msg_id)`. If it affects 0 rows, I know it's a duplicate and I ack it without doing any work."

**Spoken Format:**
"Handling duplicate messages is like having a smart mail sorter that prevents processing the same letter twice.

In messaging systems, 'exactly-once' delivery is nearly impossible due to network issues, retries, and broker restarts.

The solution is making consumers idempotent:

**Before processing**: Check if you've seen this message before using a unique ID database table.

**If duplicate**: Acknowledge the message but skip processing - don't waste resources.

**If new**: Process normally and record the ID.

This approach ensures that:
- No duplicate processing
- No wasted resources
- System remains consistent even with message retries
- You can handle 'at-least-once' delivery guarantees gracefully

It's like having a mail clerk who remembers every letter they've processed!"
