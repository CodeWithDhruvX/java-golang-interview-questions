# High-Level Design (HLD): Database Internals & Core Infrastructure

Top product companies drill deep into *why* a database behaves a certain way, evaluating your understanding of distributed consensus and storage engines.

## 1. B-Trees vs. LSM-Trees (Log-Structured Merge-Trees)
**Answer:**
This is the fundamental difference between how SQL databases and modern NoSQL databases store data on disk.
*   **B-Trees (B+ Trees):**
    *   *Used in:* MySQL, PostgreSQL, Oracle.
    *   *How it works:* Data is stored in fixed-size blocks (pages), organized as a balanced tree. Updates require navigating the tree to find the right page and modifying it in place.
    *   *Pros:* Extremely fast for reads, excellent for range queries (B+ Tree leaves are linked).
    *   *Cons:* Write amplification. Small writes cause random disk I/O, which is slow.
*   **LSM-Trees:**
    *   *Used in:* Cassandra, RocksDB, LevelDB, DynamoDB.
    *   *How it works:* Writes are strictly sequentially appended to an in-memory structure (MemTable) and a commit log. When full, MemTable is flushed to disk as an immutable SSTable (Sorted String Table). Background processes "compact" and merge these files over time.
    *   *Pros:* Extremely high write throughput because all writes are sequential appends (no random disk seeks).
    *   *Cons:* Reads are slower (have to check MemTable, then multiple SSTables). Solved partially using Bloom Filters.

## 2. Explain Database Isolation Levels
**Answer:**
Isolation handles how concurrent transactions interact. From weakest (fastest) to strongest (slowest/most locked):
1.  **Read Uncommitted:** A transaction can see uncommitted data from other transactions (Dirty Reads). Almost never used.
2.  **Read Committed:** A transaction only sees committed data. Stops dirty reads, but allows *Non-Repeatable Reads* (if you select a row twice in a transaction, the data might change if another transaction commits an update in between). Default in PostgreSQL/Oracle.
3.  **Repeatable Read:** Guarantees that if you select a row twice in the same transaction, you get the same data. Stops Non-Repeatable reads, but allows *Phantom Reads* (new rows inserted by other transactions might suddenly appear in range queries). Default in MySQL (InnoDB).
4.  **Serializable:** The strictest level. Transactions execute as if they are completely sequential. Stops all anomalies. Very slow and prone to deadlock exceptions.

## 3. What is MVCC (Multi-Version Concurrency Control)?
**Answer:**
MVCC is how systems like PostgreSQL and MySQL achieve high concurrency without relying entirely on database locks.
*   Instead of locking a row when reading/writing, the DB keeps multiple versions of that row.
*   When a transaction reads data, it sees a "Snapshot" of the database as it existed when the transaction started.
*   *Advantage:* Readers don't block writers, and writers don't block readers.

## 4. How do Distributed Databases agree on state? (Consensus Algorithms: Paxos / Raft)
**Answer:**
If you have a 5-node Zookeeper or etcd cluster, how do they agree on the "Master" or the correct data if nodes or network links fail?
*   **Consensus Algorithms:** Ensure a cluster of nodes agrees on a single source of truth, as long as a *majority (quorum)* of nodes are operational. For a 5 node cluster, you need 3 nodes to agree (Quorum = N/2 + 1).
*   **Raft (easiest to understand):** Used in etcd, Consul.
    *   *Leader Election:* Nodes vote for a leader. The leader handles all client writes.
    *   *Log Replication:* The leader appends the write to its log, sends it to followers. Once a majority of followers acknowledge the append, the leader commits the entry and informs the followers to commit.
*   **Paxos:** Used in Spanner, Cassandra (lightly). Mathematically proven but notoriously complex to implement.

## 5. CAP Theorem vs. PACELC Theorem
**Answer:**
*   **CAP Theorem:** In the case of a Network Partition (P), you must choose between Availability (A) or Consistency (C).
*   **PACELC Theorem:** CAP is incomplete because network partitions are rare. PACELC expands it:
    *   If there is a **P**artition, how does the system trade off **A**vailability and **C**onsistency?
    *   **E**lse (during normal operation), how does the system trade off **L**atency and **C**onsistency?
    *   *Example (DynamoDB):* It is an PA/EL system. During a partition it chooses availability. Normally, it chooses Low Latency over Strong Consistency (unless you explicitly request synchronous strongly consistent reads).
