# 🟡 Intermediate MySQL Interview Questions (Q26–60)

---

### 26. What are indexes in MySQL?

"An **index** is a data structure that improves the speed of data retrieval operations on a table. Without an index, MySQL must scan every row in the table to find matching records — a **full table scan**.

Think of an index like a book's index at the back — instead of reading every page, you jump directly to the relevant page number.

MySQL primarily uses **B+Tree** indexes. InnoDB also supports hash indexes internally for the adaptive hash index feature."

#### Indepth
Indexes speed up reads but **slow down writes** (INSERT, UPDATE, DELETE) because every write must also update the relevant indexes. The classic trade-off: a table with 10 indexes reads fast but writes slowly. Design indexes based on your query patterns, not exhaustively on every column. The `EXPLAIN` statement is your tool to verify whether indexes are actually being used.

---

### 27. What is a clustered index?

"A **clustered index** determines the **physical storage order** of rows in a table. In InnoDB, the primary key IS the clustered index. The actual row data is stored in the B+Tree leaf nodes of the primary key index.

This means lookups by primary key are extremely fast — the data lives right where the index points.

In InnoDB, every table has exactly one clustered index. If no primary key is defined, InnoDB creates a hidden 6-byte row ID as the clustered index."

#### Indepth
Because the clustered index determines row ordering, inserts with random or non-sequential primary keys (like UUIDs) cause **page splits** — existing B+Tree pages must split to accommodate out-of-order insertions. This creates fragmentation and slows writes significantly on large tables. Sequential keys (AUTO_INCREMENT integers) result in append-only leaf node growth, which is far more efficient. This is the primary argument against UUID primary keys in high-write InnoDB tables.

---

### 28. What is a non-clustered index?

"A **non-clustered index** (also called a **secondary index** in InnoDB) is a separate B+Tree structure that stores the indexed column values along with a **pointer back to the actual row**. In InnoDB, that pointer is the primary key value.

```sql
CREATE INDEX idx_email ON users(email);
```

This creates a secondary index on `email`. A lookup by email first searches the secondary B+Tree, retrieves the primary key, then fetches the full row from the clustered index (a **double lookup**)."

#### Indepth
The double lookup (secondary index → primary key → clustered index) is called a **bookmark lookup** or **key lookup**. To eliminate it, use a **covering index** — an index that includes all columns needed by the query, so MySQL never needs to visit the clustered index. `EXPLAIN` shows `Using index` when a covering index is used.

---

### 29. What is the difference between MyISAM and InnoDB?

"The two most common MySQL storage engines:

| Feature | MyISAM | InnoDB |
|---|---|---|
| Transactions | ❌ No | ✅ Yes (ACID) |
| Foreign Keys | ❌ No | ✅ Yes |
| Locking | Table-level | Row-level |
| Crash Recovery | ❌ Manual | ✅ Automatic |
| Full-text Search | ✅ (older) | ✅ Since 5.6 |

I always use **InnoDB** for production. MyISAM is fast for read-only reporting but completely unreliable for concurrent writes and has no data protection against crashes."

#### Indepth
MyISAM uses **table-level locking** — any write locks the entire table, blocking all concurrent reads. This is catastrophic for write-heavy concurrent workloads. InnoDB's **row-level locking** allows multiple writers to operate on different rows simultaneously. MyISAM also doesn't have crash recovery — a server crash can corrupt the entire table silently. InnoDB's redo/undo log ensures recovery to a consistent state.

---

### 30. What is a transaction?

"A **transaction** is a sequence of SQL operations treated as a single logical unit of work. Either **all operations succeed** (commit) or **all are rolled back** on failure (rollback), leaving the database in a consistent state.

```sql
START TRANSACTION;
UPDATE accounts SET balance = balance - 500 WHERE id = 1;
UPDATE accounts SET balance = balance + 500 WHERE id = 2;
COMMIT;
```

Without transactions, if the server crashes between the two UPDATEs, money disappears. A transaction ensures both happen atomically."

#### Indepth
MySQL uses **Multi-Version Concurrency Control (MVCC)** to implement transactions in InnoDB. Instead of locking rows for reads, each transaction sees a consistent **snapshot** of the database at its start time. Writers create new versions of rows rather than overwriting them. Old versions are cleaned up by the purge thread. This allows readers and writers to coexist without blocking each other.

---

### 31. What are ACID properties?

"**ACID** defines the guarantees of a reliable transaction:

- **Atomicity**: All operations in a transaction succeed or all fail together. No partial commits.
- **Consistency**: A transaction brings the database from one valid state to another. Constraints are never violated.
- **Isolation**: Concurrent transactions don't interfere with each other. Each sees a consistent view.
- **Durability**: Once committed, data survives crashes. Written to disk permanently.

InnoDB satisfies all four ACID properties, which is why I use it exclusively for production data."

#### Indepth
These properties are implemented via different mechanisms: **Atomicity** via undo logs (rollback); **Consistency** via constraints and triggers; **Isolation** via MVCC and locking; **Durability** via redo logs and `innodb_flush_log_at_trx_commit=1`. Setting `innodb_flush_log_at_trx_commit=2` improves performance but sacrifices Durability — committed transactions may be lost on OS crash (though not MySQL crash).

---

### 32. What is a deadlock?

"A **deadlock** occurs when two or more transactions are waiting for each other to release locks, creating a circular dependency that can never resolve on its own.

Example:
- Transaction A locks Row 1, then tries to lock Row 2.
- Transaction B locks Row 2, then tries to lock Row 1.
- Both are now waiting forever.

InnoDB detects deadlocks automatically and **kills the transaction with the fewest undo log bytes** (cheapest to roll back), allowing the other to proceed."

#### Indepth
InnoDB's deadlock detection runs a **wait-for graph** algorithm. You can view the last detected deadlock with `SHOW ENGINE INNODB STATUS`. To reduce deadlocks: always acquire locks in the **same order** across transactions; keep transactions short; use `SELECT ... FOR UPDATE` only when necessary; avoid user interaction inside transactions. If deadlocks are frequent, consider redesigning the access pattern to be less conflict-prone.

---

### 33. What is a view?

"A **view** is a stored, named SQL query that behaves like a virtual table. It doesn't store data itself — it executes the underlying query every time you SELECT from it.

```sql
CREATE VIEW active_users AS
SELECT id, name, email FROM users WHERE status = 'active';
```

I use views to simplify complex queries for application teams, enforce row-level security by limiting which rows/columns are visible, and provide a stable API layer when the underlying tables may change."

#### Indepth
MySQL views are not **materialized** — they recompute on every access. For performance-critical use cases, views don't help and can actually hide expensive operations. MySQL doesn't persist view query results. If you need a pre-computed view for performance, use a dedicated **summary table** refreshed by a scheduled event or application logic, or consider MySQL 8.0.x's window functions for in-query optimization instead.

---

### 34. What is a stored procedure?

"A **stored procedure** is a precompiled set of SQL statements stored in the database under a name, which can be called repeatedly.

```sql
CREATE PROCEDURE get_user(IN user_id INT)
BEGIN
  SELECT * FROM users WHERE id = user_id;
END;
CALL get_user(42);
```

I use stored procedures to centralize complex business logic in the database, reduce network round-trips (batch work happens server-side), and enforce consistent data access patterns across applications."

#### Indepth
MySQL stored procedures have significant **limitations** compared to PostgreSQL's: MySQL doesn't support procedure overloading, the debugging tools are poor, and the optimizer can't always optimize across procedure boundaries. In modern architectures, most teams have moved business logic back to the application layer and use stored procedures only for **batch jobs** or **scheduled maintenance operations** where reducing round-trips is critical.

---

### 35. What is a trigger?

"A **trigger** is a stored program that automatically executes in response to a specified event (`INSERT`, `UPDATE`, `DELETE`) on a table, either BEFORE or AFTER the event.

```sql
CREATE TRIGGER before_salary_update
BEFORE UPDATE ON employees
FOR EACH ROW
BEGIN
  INSERT INTO salary_audit(emp_id, old_sal, new_sal, changed_at)
  VALUES (OLD.id, OLD.salary, NEW.salary, NOW());
END;
```

I use triggers for audit logging, enforcing complex business rules, or automatically updating derived columns."

#### Indepth
Triggers make debugging difficult because they execute **invisibly** to the application — a plain UPDATE suddenly causes other writes. They can also create cascading effects: a trigger on Table A inserting into Table B might fire a trigger on Table B. MySQL doesn't allow triggers to call stored procedures that contain transactions, and recursive triggers are disabled by default. Use triggers sparingly; prefer application-side logic for complex workflows.

---

### 36. What are cursors in MySQL?

"A **cursor** is a database construct used inside stored procedures to iterate over a result set **row by row**.

```sql
DECLARE cur CURSOR FOR SELECT id, name FROM users;
OPEN cur;
FETCH cur INTO v_id, v_name;
CLOSE cur;
```

I use cursors only when I cannot express the logic as a set-based SQL operation — for example, when each row requires complex conditional processing that can't be done with a single UPDATE or INSERT...SELECT."

#### Indepth
Cursors are fundamentally **anti-set-based thinking** — they process data sequentially rather than leveraging the database engine's ability to operate on entire sets efficiently. A well-written set-based query almost always outperforms a cursor by orders of magnitude. Treat cursors as a last resort. Most cursor-based logic can be rewritten using JOINs, window functions, or conditional aggregation.

---

### 37. What is a subquery?

"A **subquery** is a `SELECT` query nested inside another SQL statement. It can appear in the `SELECT`, `FROM`, `WHERE`, or `HAVING` clause.

```sql
SELECT name FROM employees
WHERE department_id = (SELECT id FROM departments WHERE name = 'Engineering');
```

I use subqueries when a JOIN would be more complex or when I need a computed intermediate result set. However, in many cases, a JOIN is more readable and performs better."

#### Indepth
MySQL's subquery optimizer has historically been weaker than its JOIN optimizer. In older versions (before 5.6), correlated subqueries in `SELECT` and `WHERE` could cause significant performance issues due to repeated re-evaluation. Modern MySQL (8.0+) applies **subquery materialization** and decorrelation optimizations, bringing performance much closer to equivalent JOINs. Always verify with `EXPLAIN`.

---

### 38. What is the difference between correlated and non-correlated subqueries?

"A **non-correlated subquery** executes only once and its result is used by the outer query. It has no dependency on the outer query's current row.

A **correlated subquery** references a column from the outer query, so it re-executes for **each row** of the outer query.

```sql
-- Correlated (slow, re-executes per row):
SELECT name FROM employees e
WHERE salary > (SELECT AVG(salary) FROM employees WHERE department_id = e.department_id);
```

Correlated subqueries can be very slow on large tables because they run N times for N outer rows."

#### Indepth
MySQL's optimizer sometimes automatically converts correlated subqueries to JOINs through a process called **decorrelation**. You can verify with `EXPLAIN`. If it doesn't decorrelate, rewrite it manually as a JOIN with a derived table. Correlated subqueries in `SELECT` (e.g., `SELECT name, (SELECT MAX(sal) FROM ...) AS ...`) are particularly expensive — prefer window functions in MySQL 8.0+ for such patterns.

---

### 39. What is a self join?

"A **self join** is when a table is joined to **itself**. It's useful for hierarchical or relational data within the same table.

```sql
SELECT e.name AS employee, m.name AS manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id;
```

I use self joins for organizational hierarchies (employee → manager), parent-child category trees, or any scenario where rows in a table relate to other rows in the same table."

#### Indepth
Self joins are elegant but can be tricky to read. Always use **table aliases** to distinguish the two 'copies' of the table. For deep hierarchy traversal (multiple levels), self joins require multiple levels of nesting or you need **recursive CTEs** (available in MySQL 8.0+). `WITH RECURSIVE` is a much cleaner solution for arbitrary-depth tree queries than manual multi-level self joins.

---

### 40. What is a cross join?

"A **CROSS JOIN** produces a **Cartesian product** of two tables — every row from the first table paired with every row from the second table. No JOIN condition is specified.

```sql
SELECT a.color, b.size FROM colors a CROSS JOIN sizes b;
```

If `colors` has 5 rows and `sizes` has 4, the result has 20 rows.

I use CROSS JOIN intentionally when I need all combinations — like generating all color-size combinations for a product matrix. Accidentally producing a Cartesian product by omitting a JOIN condition is one of the most common and devastating performance mistakes."

#### Indepth
MySQL allows implicit cross joins by listing tables without a WHERE join condition: `SELECT * FROM a, b`. Without a filter, this returns `|a| × |b|` rows. On large tables, this can produce billions of rows and crash the query. Always use explicit JOIN syntax and always specify join conditions to avoid accidental Cartesian products.

---

### 41. What is a left join?

"A **LEFT JOIN** (LEFT OUTER JOIN) returns all rows from the **left (first) table** and the matched rows from the right table. Where there is no match, NULL values fill the right table's columns.

```sql
SELECT c.name, o.order_id
FROM customers c
LEFT JOIN orders o ON c.id = o.customer_id;
```

This returns all customers — including those with no orders (their `order_id` will be NULL).

I use LEFT JOIN whenever I want to preserve all records from one table regardless of whether a match exists in the other."

#### Indepth
A common pattern uses LEFT JOIN + `WHERE right.id IS NULL` to find rows that exist in the left table but NOT in the right — effectively an **anti-join**:

```sql
SELECT c.name FROM customers c
LEFT JOIN orders o ON c.id = o.customer_id
WHERE o.customer_id IS NULL; -- Customers with no orders
```

This is often more efficient than a `NOT IN` subquery or `NOT EXISTS` in MySQL, especially on large tables.

---

### 42. What is a right join?

"A **RIGHT JOIN** (RIGHT OUTER JOIN) is the mirror image of LEFT JOIN. It returns all rows from the **right (second) table** and matched rows from the left table. Unmatched left rows get NULLs.

```sql
SELECT e.name, d.name AS department
FROM employees e
RIGHT JOIN departments d ON e.department_id = d.id;
```

This returns all departments — including empty ones with no employees.

In practice, I almost always rewrite RIGHT JOINs as LEFT JOINs by swapping table order — LEFT JOINs are more intuitive for most people to read."

#### Indepth
The MySQL query optimizer treats LEFT JOIN and RIGHT JOIN identically after transforming them into an internal canonical form. A `RIGHT JOIN` is rewritten as a `LEFT JOIN` with swapped table operands during optimization. So choosing between them is purely a matter of readability, not performance. Most SQL style guides recommend defaulting to LEFT JOIN for consistency.

---

### 43. What is the difference between UNION and UNION ALL?

"Both combine results from multiple SELECT statements, but:
- **UNION** eliminates **duplicate rows** from the combined result. It performs a sort/dedup pass.
- **UNION ALL** includes **all rows**, including duplicates. No dedup, so it's faster.

```sql
SELECT city FROM customers
UNION ALL
SELECT city FROM suppliers; -- Includes duplicate cities
```

I always prefer **UNION ALL** unless I specifically need deduplication, because the sort/dedup of UNION adds overhead — especially costly on large result sets."

#### Indepth
`UNION`'s dedup is implemented via a **temporary table with a UNIQUE index** internally. Every row gets inserted into this temp table, and duplicates are rejected. On large datasets this is expensive in both CPU and memory. If duplicates are impossible by design (e.g., different source tables with different ID ranges), always use `UNION ALL` — it's semantically correct AND faster.

---

### 44. What is a temporary table?

"A **temporary table** exists only for the duration of the current session and is automatically dropped when the session ends. It's invisible to other sessions.

```sql
CREATE TEMPORARY TABLE temp_summary AS
SELECT department_id, SUM(salary) AS total FROM employees GROUP BY department_id;
```

I use temporary tables to break complex queries into steps — precompute intermediate results, then JOIN against them. This sometimes helps the optimizer better understand cardinality compared to nested subqueries."

#### Indepth
MySQL creates internal **implicit temporary tables** for operations like UNION, GROUP BY without an index, and filesorts. These start in memory (controlled by `tmp_table_size` and `max_heap_table_size`) and spill to disk when they overflow. You can monitor temp table creation with `SHOW STATUS LIKE 'Created_tmp%'`. Frequent disk-based temp tables indicate a query optimization opportunity.

---

### 45. What are prepared statements?

"A **prepared statement** is a SQL template compiled once and executed multiple times with different parameters. The query is parsed and optimized on first `PREPARE`, then efficiently re-executed with `EXECUTE`.

```sql
PREPARE get_user FROM 'SELECT * FROM users WHERE id = ?';
SET @user_id = 5;
EXECUTE get_user USING @user_id;
DEALLOCATE PREPARE get_user;
```

I primarily value them for **SQL injection prevention** — parameters are sent separately from the query, so user input can never alter the query structure."

#### Indepth
Prepared statements provide two key benefits: **security** (parameterized queries prevent SQL injection) and **performance** (parse/plan is amortized across executions). Most application drivers (JDBC, Go's `database/sql`) use prepared statements automatically. MySQL server-side prepared statements also bypass the query cache (deprecated in 8.0), relying on the re-use of the execution plan instead.

---

### 46. What is a full-text index?

"A **full-text index** enables efficient text search using natural language processing — much more powerful than `LIKE '%keyword%'` which can't use a regular B+Tree index.

```sql
CREATE FULLTEXT INDEX ft_content ON articles(title, body);
SELECT * FROM articles WHERE MATCH(title, body) AGAINST('MySQL tutorial' IN NATURAL LANGUAGE MODE);
```

I use full-text indexes for search features within MySQL. For production search at scale, I use dedicated search engines like **Elasticsearch** — but FULLTEXT is excellent for small to medium search requirements."

#### Indepth
MySQL's FULLTEXT index uses an **inverted index** structure: a mapping from each word to the rows containing it. It supports `NATURAL LANGUAGE MODE` (relevance ranked) and `BOOLEAN MODE` (operators like `+`, `-`, `*`). Words shorter than `ft_min_word_len` (default 4) are excluded. It doesn't support CJK languages well natively. InnoDB supports FULLTEXT since MySQL 5.6+.

---

### 47. What is a composite index?

"A **composite index** (or multi-column index) indexes multiple columns together in a single index structure.

```sql
CREATE INDEX idx_name_dept ON employees(last_name, department_id);
```

MySQL can use this index for queries that filter or sort by `last_name` alone or by `(last_name, department_id)` together — but **not** by `department_id` alone (the leftmost prefix rule).

I design composite indexes carefully based on the query's WHERE and ORDER BY clauses, putting the most selective column first."

#### Indepth
The **leftmost prefix rule** is critical: a composite index `(a, b, c)` can satisfy queries on `(a)`, `(a, b)`, or `(a, b, c)` — but not `(b)` or `(b, c)` alone. However, if all values of `a` are scanned (e.g., `a` has one distinct value), MySQL can use `(a, b)` to filter on `b` even though the filter skips `a`. Understanding this rule helps design indexes that serve multiple query patterns with a single index.

---

### 48. What is the EXPLAIN statement?

"`EXPLAIN` shows the **execution plan** MySQL will use for a query — which indexes are used, join types, estimated row counts, and whether filesorts or temporary tables are involved.

```sql
EXPLAIN SELECT * FROM orders WHERE customer_id = 5;
```

Key columns to check:
- `type`: join type (best `system` → `const` → `ref` → `range` → `ALL` = full scan)
- `key`: which index is used
- `rows`: estimated rows scanned
- `Extra`: `Using filesort`, `Using temporary` are bad signals

I run `EXPLAIN` on every slow query before attempting to optimize it."

#### Indepth
MySQL 8.0 introduced `EXPLAIN ANALYZE`, which actually **executes** the query and returns real runtime statistics alongside estimates. This reveals when the optimizer's row estimates are wildly off (common with skewed data distributions). `EXPLAIN FORMAT=JSON` provides the most detailed output including cost estimates. Always use the JSON or ANALYZE form for serious query tuning rather than the default tabular form.

---

### 49. What is query optimization?

"**Query optimization** is the process of improving a query's performance by reducing the amount of data it needs to process. This involves:
- Adding or refining **indexes**
- Rewriting queries (avoiding correlated subqueries, using JOINs, covering indexes)
- Using `EXPLAIN` / `EXPLAIN ANALYZE` to understand the current plan
- Analyzing table statistics with `ANALYZE TABLE`
- Avoiding `SELECT *` — only fetch needed columns

I treat optimization as a loop: measure → analyze → change → measure again."

#### Indepth
The MySQL **query optimizer** is cost-based. It estimates the cost of various execution plans using table statistics and chooses the cheapest. It can sometimes make poor choices when statistics are stale or data is skewed. Use `ANALYZE TABLE` to refresh statistics. Use `USE INDEX` / `FORCE INDEX` hints as a last resort when the optimizer consistently picks the wrong plan — but document why, as these hints can become stale after schema changes.

---

### 50. What is a stored function?

"A **stored function** is like a stored procedure but it **returns a single scalar value** and can be used inside SQL expressions.

```sql
CREATE FUNCTION get_full_name(first VARCHAR(50), last VARCHAR(50))
RETURNS VARCHAR(100)
DETERMINISTIC
RETURN CONCAT(first, ' ', last);

-- Usage:
SELECT get_full_name(first_name, last_name) FROM employees;
```

I use stored functions for reusable computation logic that needs to fit inside a `SELECT`, `WHERE`, or `ORDER BY`."

#### Indepth
Stored functions called inside queries can cause **serious performance issues** if called row-by-row on large tables — the function executes once per matching row. They also prevent index usage in WHERE clauses: `WHERE my_func(col) = value` cannot use an index on `col`. Prefer set-based SQL expressions. Use stored functions only when the computation is lightweight and not applied to every row in a full table scan.

---

### 51. What is the difference between a procedure and a function?

"| Feature | Stored Procedure | Stored Function |
|---|---|---|
| Return value | Multiple output params or none | Single scalar return value |
| Called with | `CALL proc()` | Inside SQL expressions |
| Can use in SQL? | No | Yes (in SELECT, WHERE) |
| Transactions | Can contain COMMIT/ROLLBACK | Cannot |

I use **procedures** for multi-step operations, batch jobs, and complex workflows. I use **functions** for reusable computations that need to embed in SQL statements."

#### Indepth
A stored function must have a `RETURNS` clause and `RETURN` statement. It must be marked as `DETERMINISTIC` (same inputs → same output) or `NOT DETERMINISTIC` for the binary log to handle it correctly in replication. Non-deterministic functions in replication with statement-based binary logging can cause data divergence between master and replica — always declare determinism accurately.

---

### 52. What are MySQL storage engines?

"A **storage engine** handles how MySQL physically stores and retrieves data. MySQL uses a pluggable architecture, so different tables can use different engines.

Key engines:
- **InnoDB**: Default. ACID transactions, row-level locking, foreign keys. Use for everything.
- **MyISAM**: Older. Fast reads, no transactions, table locks. Legacy only.
- **Memory**: Data in RAM. Extremely fast, lost on restart. Good for temp data.
- **Archive**: Compressed, insert-only. Good for log data.
- **CSV**: Stores data as CSV files. For data exchange.

InnoDB for nearly everything in production."

#### Indepth
The storage engine is specified at the table level: `CREATE TABLE t (...) ENGINE=InnoDB;`. MySQL treats engine selection as part of schema design. The optimizer generates execution plans that are engine-aware. Mixing engines on related tables can cause issues: for example, foreign key constraints only work between InnoDB tables — a foreign key from an InnoDB table to a MyISAM table silently fails to enforce referential integrity.

---

### 53. What is a partitioned table?

"**Table partitioning** divides a large table into smaller, manageable physical segments based on a partition key, while the table appears as a single logical unit to queries.

Types: **RANGE**, **LIST**, **HASH**, **KEY**.

```sql
CREATE TABLE orders (
  id INT, order_date DATE, amount DECIMAL(10,2)
) PARTITION BY RANGE (YEAR(order_date)) (
  PARTITION p2022 VALUES LESS THAN (2023),
  PARTITION p2023 VALUES LESS THAN (2024)
);
```

I use partitioning for time-series data where I can prune old partitions and for queries that always filter by the partition key."

#### Indepth
**Partition pruning** is the key performance benefit: queries with a `WHERE` clause on the partition key only scan the relevant partitions, not the entire table. However, if the query doesn't filter on the partition key, MySQL scans ALL partitions — potentially worse than a regular indexed table. Partitioning doesn't replace proper indexing. Also, MySQL's partitioning doesn't cross storage nodes — for true distributed scale, use sharding at the application level.

---

### 54. What is replication in MySQL?

"**Replication** is the process of copying data from one MySQL server (the **primary/master**) to one or more other servers (**replicas/slaves**) automatically and continuously.

The primary writes changes to the **binary log (binlog)**. Replicas read the binlog and apply the changes using I/O and SQL threads.

I use replication for: **high availability** (promote a replica on primary failure), **read scaling** (distribute read queries across replicas), and **zero-downtime backups** (take backup from replica)."

#### Indepth
MySQL replication is **asynchronous** by default — the primary commits and returns to the client before confirming replicas received the change. This means a primary failure in the window before the replica applied the change results in **data loss**. To prevent this, use **semi-synchronous replication** (at least one replica ACKs before the primary commits) or **Group Replication** which uses Paxos-based consensus for true synchronous commit.

---

### 55. What is master-slave replication?

"**Master-slave replication** (now called primary-replica in modern MySQL) is a replication topology where one server (master/primary) handles all writes, and one or more servers (slaves/replicas) receive and apply those writes for read queries.

The master logs all changes to the **binary log**. Each slave has an I/O thread that downloads binlog events and a SQL thread that applies them.

This is the foundational topology for MySQL high availability — combined with a proxy like ProxySQL, reads are distributed and the master is protected from read load."

#### Indepth
**Replication delay** (replica lag) is a critical operational concern. Replicas can fall behind if write volume exceeds the single-threaded SQL thread's apply capacity. MySQL 5.7+ introduced **multi-threaded replication** (parallel apply by schema or commit group), dramatically reducing lag. Monitor `Seconds_Behind_Master` via `SHOW REPLICA STATUS`. If a read query requires up-to-date data, route it to the master or use `WAIT_FOR_EXECUTED_GTID_SET()` to ensure the replica has caught up.

---

### 56. What is a trigger event?

"A **trigger event** is the DML operation that fires a trigger. MySQL supports three trigger events:
- `INSERT`: Fired when a new row is inserted.
- `UPDATE`: Fired when an existing row is modified (access OLD and NEW values).
- `DELETE`: Fired when a row is removed (access OLD values).

Combined with timing (`BEFORE` or `AFTER`), you get 6 trigger types: BEFORE INSERT, AFTER INSERT, BEFORE UPDATE, AFTER UPDATE, BEFORE DELETE, AFTER DELETE.

`OLD` references the row before modification; `NEW` references the row after."

#### Indepth
`BEFORE` triggers can **modify NEW values** before the actual insert/update — useful for auto-populating derived columns or validating data. `AFTER` triggers cannot modify the triggering row's data but are appropriate for audit logging or cascading to other tables. One important limitation: triggers can see their own table via SELECT but **cannot modify it** directly (infinite loop prevention). They can modify other tables.

---

### 57. What are the types of joins in MySQL?

"MySQL supports these join types:
- **INNER JOIN**: Rows matching in both tables.
- **LEFT (OUTER) JOIN**: All left rows + matched right rows (NULLs for non-matches).
- **RIGHT (OUTER) JOIN**: All right rows + matched left rows (NULLs for non-matches).
- **CROSS JOIN**: Cartesian product (every combination).
- **SELF JOIN**: A table joined to itself.
- **FULL OUTER JOIN**: Not natively supported — simulated with `LEFT JOIN UNION RIGHT JOIN`.

Each join type serves a different purpose. Choosing the wrong type produces incorrect results silently — a LEFT JOIN when INNER was needed includes NULL-padded garbage rows."

#### Indepth
Internally, MySQL's optimizer may use different **join algorithms**: **Nested Loop Join** (NLJ), **Block Nested Loop** (BNL), or **Hash Join** (introduced in MySQL 8.0.18). Hash joins are especially effective for equi-joins on large unsorted tables where neither side has a useful index. Monitor `EXPLAIN` output for `hash join` in 8.0+ — it signals the optimizer chose a significantly more efficient algorithm.

---

### 58. What is the difference between BETWEEN and IN?

"`BETWEEN` tests if a value falls within a **continuous range** (inclusive on both ends):
```sql
WHERE age BETWEEN 18 AND 65
```

`IN` tests if a value matches any item in a **discrete list**:
```sql
WHERE status IN ('active', 'pending', 'suspended')
```

I use `BETWEEN` for date ranges, numeric ranges, and any continuous domain. I use `IN` for categorical filters with a fixed list of values. `IN` with a subquery is also common: `WHERE id IN (SELECT id FROM premium_users)`."

#### Indepth
`BETWEEN` is inclusive on both ends: `BETWEEN 1 AND 10` includes 1 and 10. For dates with `DATETIME`, use `< '2024-01-01'` instead of `BETWEEN '2023-12-01' AND '2023-12-31'` to avoid missing rows with times on Dec 31. `IN` with a large list (e.g., 10,000 IDs) can be slow and may exceed the `max_allowed_packet` size. For large ID lists, use a temporary table + JOIN instead.

---

### 59. What is a wildcard character?

"MySQL supports two wildcards in LIKE patterns:
- `%` (percent): Matches **any sequence** of characters (including zero).
- `_` (underscore): Matches exactly **one** character.

```sql
WHERE name LIKE 'Jo%'      -- Starts with 'Jo'
WHERE code LIKE 'A_C'     -- 'A' + any one char + 'C'
WHERE email LIKE '%@gmail.com'  -- Ends with @gmail.com
```

Leading wildcards (`'%text'`) prevent index use — MySQL can't narrow the B+Tree search from the left side without knowing the prefix."

#### Indepth
The leading wildcard problem (`LIKE '%keyword'`) forces a **full index scan** or full table scan because B+Tree indexes are built left-to-right and can't satisfy an open-ended prefix search. Solutions: use FULLTEXT index for keyword search, store a reversed version of the column and search reversed keywords with a trailing `%`, or migrate to an external search engine like Elasticsearch.

---

### 60. What is the difference between COUNT(*) and COUNT(column)?

"`COUNT(*)` counts **all rows** in the result set, including those with NULLs in any column.

`COUNT(column)` counts only rows where that specific column's value is **NOT NULL**.

```sql
SELECT COUNT(*), COUNT(phone_number) FROM users;
-- May return 1000, 750 if 250 users have no phone number
```

I use `COUNT(*)` for total row counts and `COUNT(column)` when I specifically want to know how many non-NULL values exist in a column. Using them interchangeably is a frequent bug."

#### Indepth
In InnoDB, `COUNT(*)` is **not stored as metadata** (unlike MyISAM which caches the row count). InnoDB must scan rows to count them, because MVCC means different transactions see different row counts simultaneously. For very large tables, `SELECT COUNT(*)` can be slow. Use `SHOW TABLE STATUS` for an approximate count, or maintain a separate counter table updated by triggers/application code if exact counts are needed with low latency.

---
