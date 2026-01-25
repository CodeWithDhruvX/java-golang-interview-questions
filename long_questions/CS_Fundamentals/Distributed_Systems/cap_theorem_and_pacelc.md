# CAP Theorem & PACELC

## 1. CAP Theorem
Proposed by Eric Brewer, it states that a distributed data store can only provide **two** of the following three guarantees:

### C - Consistency (Linearizability)
*   **Definition**: Every read receives the most recent write or an error.
*   **Meaning**: All nodes see the same data at the same time.

### A - Availability
*   **Definition**: Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
*   **Meaning**: The system stays up even if some nodes are down.

### P - Partition Tolerance
*   **Definition**: The system continues to operate despite an arbitrary number of messages being dropped or delayed by the network between nodes.
*   **Reality**: In distributed systems, network partitions (P) are inevitable. You *must* choose P.
*   **Theorem Implication**: You essentially choose between **CP** (Consistency + Partition Tolerance) or **AP** (Availability + Partition Tolerance).

### Examples
*   **CP (Consistency over Availability)**:
    *   *Example*: Banking usage. If the ATM cannot reach the main bank server, it should refuse to withdraw money rather than risk an overdraft (inconsistent balance).
    *   *Databases*: MongoDB (default), HBase, Redis (if configured for strong consistency).
    *   *Behavior*: Returns Error/Timeout during partition.
*   **AP (Availability over Consistency)**:
    *   *Example*: Facebook Likes. It's okay if I see 100 likes and you see 99 for a few seconds, as long as the page loads.
    *   *Databases*: Cassandra, DynamoDB, CouchDB.
    *   *Behavior*: Returns potentially stale data during partition, reconciles later (Eventual Consistency).

## 2. PACELC Theorem
CAP is too simplistic because it only describes behavior *during a partition*. What about when the network is running normally?

**PACELC** extends CAP:
*   If there is a Partition (**P**), how does the system trade off availability and consistency (**A** vs **C**)?
*   **E**lse (**E**), when the system is running normally in the absence of a partition, how does the system trade off latency (**L**) and consistency (**C**)?

### The "ELC" Part (Else Latency vs Consistency)
*   If you want strong consistency (C), you must replicate data to all nodes before confirming the write. This increases **Latency (L)**.
*   If you want low latency (L), you return success after writing to one node and replicate asynchronously. This sacrifices **Consistency (C)**.

### Examples
*   **DynamoDB / Cassandra**: PA / EL.
    *   Partition? Choose Availability.
    *   Normal? Choose Low Latency (Eventual Consistency).
*   **MongoDB / HBase**: PC / EC.
    *   Partition? Choose Consistency (Stop writes).
    *   Normal? Choose Consistency (Wait for replica ack).

## 3. Interview Questions
1.  **Can a system be CA?**
    *   *Ans*: Only if it's not distributed (Single Node). If you are distributed, you cannot guarantee no partitions. So for distributed systems, CA is impossible.
2.  **How does Master-Slave replication fit into CAP?**
    *   *Ans*:
        *   **Async Replication**: AP. If master acts, slave might be stale.
        *   **Sync Replication**: CP. If slave is down, master cannot confirm write, so availability is lost.
