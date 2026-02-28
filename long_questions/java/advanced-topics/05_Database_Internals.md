# Database Internals & System Operations for Interviews

Product-based companies expect back-end and full-stack developers to know exactly what happens inside a database engine when they write a query. "I use Spring Data JPA" is not an acceptable answer when debugging a production incident involving slow queries or deadlocks.

---

## 🗄️ 1. Core Database Architecture

### B-Trees and Indexes
*   **How data is stored:** An index is just a separate data structure (usually a B+ Tree) that stores a reference (pointer) to the actual row on the disk.
*   **B-Tree vs. Hash Index:** Hash Indexes give O(1) lookups but cannot do range queries (e.g., `WHERE age > 30`). B-trees keep data sorted, allowing O(log N) lookups, range queries, and ordering (`ORDER BY`).
*   **Clustered vs. Non-Clustered Indexes:**
    *   **Clustered (Primary Key):** The actual table data is stored *inside* the leaf nodes of the B-Tree index. A table can only have ONE clustered index.
    *   **Non-Clustered (Secondary Index):** A separate structure that stores the indexed column's value and a pointer (the Row ID or Primary Key) back to the clustered index.

### query Execution & Optimization
When you run `SELECT * FROM Users WHERE email = 'x@y.com'`, the database does not just execute it immediately.
1.  **Parser/Lexer:** Checks syntax.
2.  **Query Rewriter:** Optimizes the logical query (e.g., pushing `WHERE` clauses down before a `JOIN`).
3.  **Query Planner/Optimizer:** Generates multiple execution plans and costs them based on statistics (table row count, index selectivity). It chooses the cheapest plan.
4.  **Execution Engine:** Executes the chosen plan.

**EXPLAIN / EXPLAIN ANALYZE:**
The most important command you will ever run. It tells you *how* the database plans to fetch your data.
*   **Table Scan / Seq Scan:** Bad. The DB reads every row in the table (O(N)).
*   **Index Scan:** Good. The DB uses the B-tree to find the exact row (O(log N)).
*   **Index Only Scan (Covering Index):** The best. The query only requests columns that are *already* in the index. The DB entirely skips reading the actual table data from disk.

---

## 🛡️ 2. ACID Properties & Transactions

Understanding transaction guarantees is critical for financial apps and inventory systems.

*   **Atomicity:** "All or Nothing." If a transaction has 5 `INSERT` statements and the 4th one fails, the entire transaction is rolled back. Achieved via the **Write-Ahead Log (WAL)** or UNDO logs.
*   **Consistency:** The database ensures the data constraints (Foreign Keys, Unique checks) hold true before and after the transaction ends.
*   **Isolation:** How much concurrent transactions can "see" each other's uncommitted data. (See Isolation Levels below).
*   **Durability:** Once a transaction commits, it is permanently saved to non-volatile storage (disk), even if the power fails immediately after. Achieved via REDO logs.

---

## 🚦 3. Transaction Isolation Levels (Crucial Topic)

When two transactions run concurrently, anomalies can occur. SQL defines 4 isolation levels to prevent them (at the cost of performance/locking).

### The Anomalies:
1.  **Dirty Read:** Transaction A reads data that Transaction B has modified but not yet committed. If B rolls back, A has read invalid data.
2.  **Non-Repeatable Read:** Transaction A reads a row. Transaction B updates that row and commits. Transaction A reads the same row again and gets a different value.
3.  **Phantom Read:** Transaction A runs a range query (`SELECT count(*) FROM users WHERE age > 20`). Transaction B inserts a new user (age 25) and commits. Transaction A runs the exact same query again and sees a "phantom" new row.

### The Isolation Levels (From Fastest to Safest):
1.  **Read Uncommitted (Weakest):** No locks. Prevents nothing. *Dirty Reads can happen.*
2.  **Read Committed (Postgres/Oracle Default):** A query only sees data that was committed before the query began. *Prevents Dirty Reads.*
3.  **Repeatable Read (MySQL/InnoDB Default):** Uses locks or multi-version concurrency control (MVCC) to ensure that if you read a row twice in one transaction, it will be exactly the same. *Prevents Dirty Reads and Non-Repeatable Reads.*
4.  **Serializable (Strictest):** Effectively executes transactions serially (one after another). Extensive locking. Horrible for performance. *Prevents all anomalies, including Phantom Reads.*

---

## ⚡ 4. Top 10 Database Interview Questions

### 1. How does a database handle a crash exactly when a transaction is committing?
**Answer:** The **Write-Ahead Log (WAL)**. Before modifying actual data pages on disk, the database strictly writes the "intent" to modify them into an append-only WAL file. If the system crashes, upon reboot, the DB reads the WAL and replays the committed transactions (REDO) and rolls back the unfinished ones (UNDO).

### 2. What is the N+1 problem and how do you fix it at the database level?
**Answer:** While typically associated with ORMs (Hibernate), you fix it at the DB level by replacing `N` individual `SELECT` queries with a single query using a `JOIN` or fetching children with an `IN (...)` clause.

### 3. How does MVCC (Multi-Version Concurrency Control) work?
**Answer:** Instead of locking a row when someone is reading it (blocking writers), the DB creates a "snapshot" or a new version of the row.
*   *Readers don't block writers.*
*   *Writers don't block readers.*
Postgres uses XID (Transaction IDs) to check which version of a row a specific transaction is explicitly allowed to see.

### 4. Why use a Composite Index instead of two Single Indexes?
**Answer:**
If you query `WHERE first_name = 'John' AND last_name = 'Doe'`, two single indexes require the DB to fetch sets from both indexes and intersect them in memory (slower).
A composite index (`INDEX (last_name, first_name)`) stores the data pre-sorted by `last_name`, and *then* by `first_name`. A single B-Tree lookup finds the exact rows.
*(Note: Order matters! Due to the Leftmost Prefix Rule, an index on (A, B) works for queries on A, and queries on A+B, but NOT for queries on B alone).*

### 5. What is Database Sharding vs. Partitioning?
**Answer:**
*   **Partitioning:** Splitting a massive table into smaller logical tables within the *same* database instance (e.g., partitioning an `Orders` table by month). Speeds up queries and index maintenance.
*   **Sharding:** Splitting a massive table across *multiple physical database servers/instances*. Usually done using a Hash (Hash of UserID determines which shard server holds the data). Extremely complex, breaks ACID across shards, but allows infinite horizontal scaling.

### 6. Design a schema for a Hierarchical structure (like Employee-Manager or Reddit Comments).
**Answer:**
1.  **Adjacency List:** Just add a `parent_id` column. Easy to insert, but finding all descendants Requires recursive queries (`WITH RECURSIVE` CTE in SQL).
2.  **Path Enumeration (Materialized Path):** Store the whole path in a column (e.g., `1/4/9/`).
3.  **Closure Table:** A separate table storing all ancestor-descendant relationships. Very fast reads, expensive writes.

### 7. How would you paginate 10 million rows? `LIMIT/OFFSET` vs Keyset Pagination?
**Answer:**
*   **`LIMIT 10 OFFSET 5000000`:** Extremely slow. The database still has to fetch, sort, and discard the first 5,000,000 rows.
*   **Keyset Pagination (Seek Method):** Remember the last seen ID. `WHERE id > 5000000 ORDER BY id LIMIT 10`. The database uses the B-tree index to instantly jump to ID 5,000,000 and grabs the next 10. O(1) time complexity.

### 8. Explain the CAP Theorem in the context of Cassandra vs MySQL.
**Answer:**
MySQL (in a primary-replica setup with synchronous replication) emphasizes **Consistency and Partition Tolerance (CP)**. A write won't succeed unless the replica acknowledges it.
Cassandra is built for massive scale across multiple datacenters and emphasizes **Availability and Partition Tolerance (AP)**. It will always accept a write request, even if nodes are down, achieving "Eventual Consistency".

### 9. What is Connection Pooling and why is it necessary?
**Answer:**
Opening a TCP connection to a database involves network latency, 3-way handshakes, and allocating memory buffers in the DB engine. Doing this for every HTTP request is terribly slow.
A connection pool (like HikariCP) pre-opens a set of reusable connections. When your app needs one, it borrows it, runs the query, and returns it to the pool instead of closing it.

### 10. How do you find a slow query causing production issues?
**Answer:**
Enable the **Slow Query Log** in MySQL/Postgres. Use tools like `pg_stat_statements` to find the queries taking the longest cumulative time. Then take the worst offenders and run `EXPLAIN ANALYZE` on them to see where they are doing Full Table Scans. Add indexes where necessary.
