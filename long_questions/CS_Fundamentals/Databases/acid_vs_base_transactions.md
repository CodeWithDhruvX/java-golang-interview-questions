# ACID vs BASE Transactions

## 1. ACID (Strong Consistency)
Standard for Relational Databases (SQL). Prioritizes data integrity.

*   **A - Atomicity**: "All or Nothing". A transaction involving multiple steps (deduct money A, add money B) must either complete fully or fail fully.
*   **C - Consistency**: The DB moves from one valid state to another valid state. Constraints (Keys, Data types) are enforced.
*   **I - Isolation**: Transactions occurring concurrently result in the same state as if they were executed sequentially.
    *   *Isolation Levels*: Read Uncommitted, Read Committed, Repeatable Read, Serializable.
*   **D - Durability**: Once committed, data is saved permanently, even if power is lost (Thanks to WAL).

## 2. BASE (Eventual Consistency)
Standard for NoSQL Databases. Prioritizes Availability and Performance over immediate Consistency.

*   **BA - Basic Availability**: The system guarantees availability. The system works even if some nodes are down. (Unlike ACID, which might reject writes).
*   **S - Soft State**: The state of the system may change over time, even without input (due to replication).
*   **E - Eventual Consistency**: If the system stops receiving input, it will eventually become consistent.
    *   *Example*: You post a comment. Your friend sees it 2 seconds later.

## 3. Comparison

| Feature | ACID (SQL) | BASE (NoSQL) |
| :--- | :--- | :--- |
| **Focus** | Consistency, Integrity | Availability, Scale |
| **Transaction** | Complex, multi-row | Simple, mostly single-row |
| **Latency** | Higher (Locks, 2PC) | Lower (No locks) |
| **Use Case** | Financial, Billing, Inventory | Social Feeds, Analytics, IoT |

## 4. Interview Questions
1.  **Can NoSQL databases support ACID?**
    *   *Ans*: Yes. Many NoSQL DBs (like MongoDB, DynamoDB) now support multi-document ACID transactions, but with performance penalties.
2.  **What is the Saga Pattern?**
    *   *Ans*: A design pattern to handle distributed transactions across microservices (where 2PC is too slow). It breaks a transaction into a sequence of local transactions. If one fails, it executes "Compensating Transactions" to undo changes.
