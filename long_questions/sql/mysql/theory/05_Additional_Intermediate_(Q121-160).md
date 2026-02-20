# 🟡 Additional Intermediate Level Questions (Q121–160)

---

### 121. What is the difference between REPLACE and INSERT?

"**INSERT** adds a new row and fails with `ERROR 1062` if a duplicate primary/unique key already exists.

**REPLACE** first tries to INSERT. If a duplicate key conflict occurs, it **DELETEs the conflicting row** and then INSERTs the new row.

```sql
REPLACE INTO users (id, name, email) VALUES (1, 'Dhruv', 'new@email.com');
```

I use REPLACE when I want 'insert or fully overwrite' semantics. Caution: REPLACE deletes then re-inserts, so it triggers DELETE and INSERT triggers, resets AUTO_INCREMENT sequences on the row, and deletes associated foreign-key cascade children."

#### Indepth
`REPLACE` is deceptively dangerous. Because it performs a DELETE + INSERT (not an UPDATE), it: (1) Triggers `BEFORE DELETE`, `AFTER DELETE`, `BEFORE INSERT`, `AFTER INSERT` triggers — not UPDATE triggers. (2) Resets auto-incremented non-PK values. (3) Can cascade-delete child rows via foreign keys. Most of the time, `INSERT ... ON DUPLICATE KEY UPDATE` is the safer, more precise alternative since it actually updates only specified columns without deleting the row.

---

### 122. What is ON DUPLICATE KEY UPDATE?

"`ON DUPLICATE KEY UPDATE` allows an INSERT to gracefully handle unique key violations by performing an UPDATE on the existing row instead of failing.

```sql
INSERT INTO page_views (page_id, views)
VALUES (42, 1)
ON DUPLICATE KEY UPDATE views = views + 1;
```

This is an **upsert** — insert if new, update if exists.

I use this pattern constantly for counters, session tracking, and idempotent data pipelines where the same event may be received more than once."

#### Indepth
`ON DUPLICATE KEY UPDATE` affects row count differently: if a row is inserted, it returns affected rows = 1. If updated, affected rows = 2. If the UPDATE results in no change (same values), affected rows = 0. The special `VALUES(col)` function (deprecated in MySQL 8.0; use aliases instead) referenced the values from the INSERT part inside the UPDATE clause. The modern syntax uses row alias: `INSERT INTO t VALUES (1, 'foo') AS new ON DUPLICATE KEY UPDATE name = new.name`.

---

### 123. What is a covering index and when is it useful?

"A **covering index** includes all columns a query needs in its index structure, so MySQL can satisfy the entire query by reading only the index — never touching the actual row data.

It's useful when:
- A query reads only a few specific columns.
- The table is very large (index pages are much smaller than table pages).
- The query runs very frequently.

```sql
-- Query: SELECT email, status FROM users WHERE user_id = 42
-- Covering index: (user_id, email, status) -- all needed columns present
CREATE INDEX idx_covering ON users(user_id, email, status);
```"

#### Indepth
Covering indexes are identified by `Using index` in EXPLAIN's Extra column — the golden indicator of optimal index usage. The benefit scales with table width: on a 100-column table with 500-byte rows, a covering index for a 3-column query might be 30 bytes per entry — reading it is 16× more cache-efficient. For very frequently queried hot paths, designing a dedicated covering index is worth the extra write overhead. However, don't over-index: each index adds write cost and memory pressure.

---

### 124. How does MySQL use indexes in JOIN operations?

"In JOIN operations, MySQL's optimizer selects a **driving table** (the first table scanned) and uses an index on the **joining column of the other table** to look up matching rows.

```sql
SELECT o.*, c.name
FROM orders o
JOIN customers c ON o.customer_id = c.id;
```

MySQL scans `orders`, then for each row, uses the index on `customers.id` to find the matching customer. `customer_id` in `orders` and `id` in `customers` should both be indexed.

Missing indexes on JOIN columns cause nested loop joins to degrade to O(n²) full scans."

#### Indepth
MySQL uses **Nested Loop Join (NLJ)** as its primary join algorithm: a loop over the outer (driving) table, with an index lookup in the inner table for each outer row. MySQL 8.0+ also supports **Hash Join**: the smaller table is hashed into memory, then the larger table is scanned and matched against the hash — efficient for large equi-joins without a useful index. Check EXPLAIN for `hash join` in the Extra column. Hash join in MySQL avoids creating explicit temporary tables and is automatically chosen when no index exists on the join condition.

---

### 125. What is index cardinality?

"**Index cardinality** is the estimated number of **unique values** in an indexed column (or combination of columns). MySQL stores this in `INFORMATION_SCHEMA.STATISTICS.CARDINALITY`.

High cardinality = many unique values (e.g., user_id, email) → index very useful.
Low cardinality = few unique values (e.g., boolean, status with 3 values) → index may be ignored.

I check cardinality with: `SHOW INDEX FROM tablename;` or query `information_schema.STATISTICS`."

#### Indepth
Cardinality is an **estimate**, not an exact count — MySQL samples a configurable number of index pages (`innodb_stats_persistent_sample_pages`, default 20) to estimate it. For skewed data distributions, cardinality may be accurate globally but misleading for specific values. MySQL 8.0 introduced **column histograms** (`ANALYZE TABLE t UPDATE HISTOGRAM ON col`) that capture value distribution, allowing the optimizer to make better decisions when cardinality alone is insufficient (e.g., when one status value has 95% of rows).

---

### 126. What is selectivity in indexing?

"**Selectivity** is the fraction of rows an index condition eliminates: `unique_values / total_rows`. It ranges from 0 (useless) to 1 (perfectly selective).

High selectivity (0.9–1.0): `WHERE email = 'x@y.com'` matches very few rows → index is effective.
Low selectivity (0.001): `WHERE status = 'active'` may match 99% of rows → index scan slower than full table scan.

MySQL's optimizer uses selectivity to decide if using an index is worth it. It switches to a full table scan if it estimates the index would scan too many rows anyway."

#### Indepth
The optimizer's selectivity calculation is based on **statistics**, not the actual query predicate value. It doesn't know that `status = 'active'` happens to match 98% of rows unless histograms are available. Without histograms, the optimizer might consistently choose a bad index for a frequently queried value. Histograms in MySQL 8.0 provide **per-value statistics** that let the optimizer differentiate between `status = 'active'` (98% selectivity = bad index use) and `status = 'deleted'` (0.1% selectivity = excellent index use).

---

### 127. What is a slow query log and how do you enable it?

"The **slow query log** records queries that take longer than `long_query_time` seconds to execute. It's the primary tool for identifying performance bottlenecks.

Enable it:
```ini
# In my.cnf:
slow_query_log = 1
slow_query_log_file = /var/log/mysql/mysql-slow.log
long_query_time = 1        # Log queries > 1 second
log_queries_not_using_indexes = 1  # Also log index-less queries
```

Or at runtime: `SET GLOBAL slow_query_log = 1;`

I analyze the slow query log with **`pt-query-digest`** from Percona Toolkit, which aggregates and ranks queries by total time consumed."

#### Indepth
Set `long_query_time = 0` temporarily to log ALL queries during profiling sessions — useful for finding hidden micro-slow queries that add up. `log_queries_not_using_indexes = 1` is extremely valuable: it catches queries that use full table scans even if they complete quickly on small tables — these will become slow as the table grows. `log_throttle_queries_not_using_indexes` limits how many such queries are logged per minute to avoid log spam on busy servers.

---

### 128. How do you find duplicate records in a table?

"Use GROUP BY + HAVING to find duplicated values:

```sql
-- Find duplicate emails:
SELECT email, COUNT(*) AS cnt
FROM users
GROUP BY email
HAVING cnt > 1;

-- Find full duplicate rows:
SELECT email, name, COUNT(*) AS cnt
FROM users
GROUP BY email, name
HAVING cnt > 1;
```

This approach is efficient and uses indexes on the grouped columns.

In interviews, this is one of the most common SQL coding questions. I always answer with GROUP BY + HAVING first, then offer window function alternatives."

#### Indepth
An alternative using window functions (MySQL 8.0+):
```sql
SELECT * FROM (
  SELECT *, COUNT(*) OVER (PARTITION BY email) AS cnt
  FROM users
) t WHERE cnt > 1;
```
This returns ALL duplicate rows (not just a count), making it easy to review and selectively delete them. The GROUP BY approach is more efficient for just detecting duplicates; the window function approach is better for inspecting and deciding which duplicate rows to keep.

---

### 129. How do you delete duplicate rows while keeping one?

"The classic approach: delete rows that are NOT the minimum (or maximum) ID among duplicates.

```sql
DELETE FROM users
WHERE id NOT IN (
  SELECT min_id FROM (
    SELECT MIN(id) AS min_id
    FROM users
    GROUP BY email
  ) AS keep
);
```

MySQL requires the double subquery because you can't directly reference the same table being deleted in a subquery in one level.

I always run a SELECT with the same WHERE logic first to verify what will be deleted before executing the DELETE."

#### Indepth
A more efficient approach for large tables uses a **self-join DELETE**:
```sql
DELETE u1 FROM users u1
INNER JOIN users u2
ON u1.email = u2.email AND u1.id > u2.id;
```
This deletes the higher-ID duplicates (keeping the lowest ID). The self-join leverages an index on both `email` and `id`. For very large tables, delete duplicates in batches with `LIMIT` to avoid large transactions and lock contention: `DELETE u1 ... LIMIT 1000;` in a loop.

---

### 130. What is a pivot query in MySQL?

"A **pivot query** transforms rows into columns — displaying grouped data horizontally. MySQL doesn't have a native `PIVOT` keyword, so we use conditional aggregation.

```sql
SELECT department,
  SUM(CASE WHEN gender = 'M' THEN 1 ELSE 0 END) AS male_count,
  SUM(CASE WHEN gender = 'F' THEN 1 ELSE 0 END) AS female_count
FROM employees
GROUP BY department;
```

I use conditional aggregation (`CASE WHEN + SUM/COUNT`) to create pivot-style reports entirely in SQL. For dynamic pivots (unknown number of columns), I generate the SQL dynamically in the application."

#### Indepth
Dynamic pivoting in MySQL requires **dynamic SQL generation**: query the distinct pivot values first, build the SQL string, then execute it with a prepared statement. This is cumbersome and procedural. Most teams prefer to handle pivot logic in the application layer (Python pandas, or BI tools like Metabase) rather than inside MySQL. However, prepared statements for dynamic pivots are used in stored procedures for reporting dashboards that need database-side aggregation.

---

### 131. What is a recursive CTE?

"A **recursive CTE** is a CTE that references itself, enabling **hierarchical or recursive data traversal** — like traversing an org chart, category tree, or bill of materials.

```sql
WITH RECURSIVE org_chart AS (
  -- Anchor: start with top-level (no manager)
  SELECT id, name, manager_id, 0 AS depth
  FROM employees WHERE manager_id IS NULL

  UNION ALL

  -- Recursive: join each level's children
  SELECT e.id, e.name, e.manager_id, oc.depth + 1
  FROM employees e
  JOIN org_chart oc ON e.manager_id = oc.id
)
SELECT * FROM org_chart ORDER BY depth, name;
```

Available in MySQL 8.0+. Before 8.0, this required a stored procedure with loops."

#### Indepth
Recursive CTEs use a **two-part structure**: the anchor query (base case, non-recursive) and the recursive member (references the CTE itself). MySQL limits recursion depth with `cte_max_recursion_depth` (default 1000) to prevent infinite loops. For **genuine infinite recursion** (circular parent-child relationships), add a depth counter and a MAXDEPTH guard condition. Recursive CTEs in MySQL 8.0 also enable **graph traversal**, shortest path calculations, and **Fibonacci-like sequence generation** — powerful capabilities previously requiring procedural code.

---

### 132. What is the difference between RANK(), DENSE_RANK(), and ROW_NUMBER()?

"All three are window functions assigning sequence numbers, but handle ties differently:

- **`ROW_NUMBER()`**: Assigns unique sequential integers with no gaps. Ties get arbitrary ordering.
- **`RANK()`**: Assigns the same rank to ties, but **skips** subsequent numbers. (1, 2, 2, 4)
- **`DENSE_RANK()`**: Assigns the same rank to ties, with **no gaps**. (1, 2, 2, 3)

```sql
SELECT name, salary,
  ROW_NUMBER() OVER (ORDER BY salary DESC) AS row_num,
  RANK()       OVER (ORDER BY salary DESC) AS rnk,
  DENSE_RANK() OVER (ORDER BY salary DESC) AS dense_rnk
FROM employees;
```"

#### Indepth
**Use case guidance**: `ROW_NUMBER()` for pagination (you need unique sequential numbering). `RANK()` for competitive ranking where ties should reflect position gaps (like sports rankings: two 2nd places means no 3rd). `DENSE_RANK()` for category rankings where you want continuous rank numbers regardless of ties. Add `PARTITION BY` to rank within groups: `RANK() OVER (PARTITION BY department ORDER BY salary DESC)` ranks employees within each department independently.

---

### 133. What are window frames?

"A **window frame** defines the subset of rows within a window partition that a window function considers for each row's calculation.

```sql
SELECT name, salary,
  AVG(salary) OVER (
    ORDER BY hire_date
    ROWS BETWEEN 2 PRECEDING AND CURRENT ROW  -- 3-row rolling average
  ) AS rolling_avg
FROM employees;
```

Frame types:
- `ROWS`: Physical row count.
- `RANGE`: Logical value range.

Boundaries: `UNBOUNDED PRECEDING`, `N PRECEDING`, `CURRENT ROW`, `N FOLLOWING`, `UNBOUNDED FOLLOWING`."

#### Indepth
`RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW` is the default frame for most window functions and computes a **cumulative value** from the start of the partition to the current row. `ROWS BETWEEN N PRECEDING AND N FOLLOWING` creates a **sliding window** for moving averages. The distinction between `ROWS` and `RANGE` matters for ties: with `RANGE`, all rows with the same ORDER BY value as the current row are included in the frame together — potentially including many more rows than expected.

---

### 134. What is a functional index?

"A **functional index** (MySQL 8.0+) is an index built on an **expression** or **function applied to a column**, not the raw column value.

```sql
-- Index on uppercase email for case-insensitive search:
CREATE INDEX idx_upper_email ON users((UPPER(email)));

-- Query can now use the index:
SELECT * FROM users WHERE UPPER(email) = 'DHRUV@EXAMPLE.COM';
```

Without a functional index, `WHERE UPPER(col) = val` would force a full table scan (function prevents index use). Functional indexes solve this by pre-computing and indexing the function result."

#### Indepth
Functional indexes are implemented as **virtual generated column + index** under the hood. MySQL automatically creates an invisible virtual generated column with the expression, then indexes that column. You can achieve the same manually by creating an explicit generated column and indexing it. Functional indexes are particularly powerful for **JSON path indexing**: `CREATE INDEX idx_json_name ON t((data->>'$.name'))` — enables indexed lookups into JSON documents without extracting to a separate column.

---

### 135. What is an invisible index?

"An **invisible index** (MySQL 8.0+) is an index that exists and is maintained by the database but is **hidden from the optimizer** — the optimizer won't use it for query planning.

```sql
CREATE INDEX idx_test ON users(email) INVISIBLE;
-- Make existing index invisible:
ALTER TABLE users ALTER INDEX idx_test INVISIBLE;
-- Make it visible again:
ALTER TABLE users ALTER INDEX idx_test VISIBLE;
```

This is invaluable for **safely testing index removal** without actually dropping it: make the index invisible, observe query performance, then drop it if results are acceptable — or make it visible again to instantly restore performance."

#### Indepth
Invisible indexes solve a real operational pain point: **irreversible index drops**. Without invisible indexes, dropping an index and discovering it was critical results in an expensive `CREATE INDEX ... ALGORITHM=INPLACE` rebuild. With invisible indexes: the validation step is essentially free — toggle visibility, test, and either drop or restore in milliseconds. Also useful for **index canary testing**: make a new index invisible, validate with targeted queries using `USE INDEX()` hint, then make it visible when confident.

---

### 136. What is index condition pushdown?

"**Index Condition Pushdown (ICP)** is a MySQL optimization where applicable WHERE conditions are evaluated at the **storage engine level** (InnoDB), rather than being returned to the MySQL server layer for filtering.

Without ICP: InnoDB fetches rows matching the index key, returns them all to MySQL server, MySQL server filters on remaining WHERE conditions.

With ICP: InnoDB evaluates the additional WHERE conditions on index columns **before** fetching the full row. It only returns rows that pass all conditions, reducing table-level row fetches.

EXPLAIN shows `Using index condition` when ICP is active."

#### Indepth
ICP is automatically applied when the WHERE clause contains conditions on index columns that can't drive the index range scan but can be evaluated on index data. For example, with a composite index `(a, b)` and `WHERE a = 5 AND b LIKE 'Jo%'`: the range scan is on `a = 5`, and ICP pushes `b LIKE 'Jo%'` down to InnoDB to filter before fetching full rows. This can dramatically reduce the number of full-row reads for range scans on wide tables.

---

### 137. What is the optimizer in MySQL?

"The MySQL **query optimizer** is the component that takes a parsed SQL query and determines the most efficient **execution plan** — which indexes to use, in what order to join tables, whether to use hash join or nested loop.

It's a **cost-based optimizer (CBO)**: it calculates the estimated cost (in I/O units) of multiple candidate plans and selects the cheapest one.

The optimizer uses **statistics** (cardinality, histograms), **available indexes**, table sizes, join conditions, and configuration parameters to make its decisions."

#### Indepth
MySQL's optimizer has known limitations: it evaluates join orderings exhaustively up to `optimizer_search_depth` tables (default 62), which is theoretically exhaustive but uses strict cost models that occasionally miss better plans. The optimizer cannot always look through **derived tables** (subqueries in FROM). MySQL 8.0 improved optimizer significantly with hash joins, derived table merging, and better histogram integration. For tuning: `optimizer_switch` exposes individual optimizer features as toggles — you can disable individual optimizations to diagnose whether they're causing bad plans.

---

### 138. What is a derived table?

"A **derived table** is a subquery used in the `FROM` clause that acts as a temporary virtual table.

```sql
SELECT dept, avg_salary
FROM (
  SELECT department AS dept, AVG(salary) AS avg_salary
  FROM employees
  GROUP BY department
) AS dept_averages
WHERE avg_salary > 60000;
```

The inner query is the derived table. It's computed first, then the outer query operates on its results.

I use derived tables to break complex queries into logical stages, making them easier to read and understand."

#### Indepth
MySQL 5.7+ can often **merge** derived tables back into the outer query (optimizer trace shows DERIVED_TABLE_MERGING). This is generally more efficient than materializing the derived table into a temp table. However, some derived tables cannot be merged (e.g., those using UNION, GROUP BY, LIMIT, DISTINCT, window functions) — MySQL materializes these as temp tables. In MySQL 8.0, the optimizer also considers merging lateral joins and more complex CTEs, reducing materialization overhead.

---

### 139. What is a lateral join?

"A **lateral join** (MySQL 8.0+) allows a subquery in the FROM clause to reference columns from preceding tables in the same FROM clause — creating a **correlated derived table**.

```sql
SELECT e.name, top_sale.amount
FROM employees e
JOIN LATERAL (
  SELECT amount FROM sales
  WHERE sales.employee_id = e.id
  ORDER BY amount DESC
  LIMIT 1
) AS top_sale ON TRUE;
```

Without LATERAL, a subquery in FROM cannot reference `e.id`. With LATERAL, it can — enabling per-row subquery evaluation.

This simplifies queries previously requiring correlated subqueries in SELECT or complex self-joins."

#### Indepth
LATERAL joins are especially powerful for **top-N per group** problems — finding the highest-N rows per group without window functions. The LATERAL subquery executes once per row of the driving table, similar to a correlated subquery but with the flexibility of a derived table (can use ORDER BY, LIMIT, multiple expressions). For large tables, ensure the LATERAL subquery uses indexes efficiently — it runs for every row of the outer table. Check EXPLAIN for `DEPENDENT DERIVED` to confirm lateral evaluation.

---

### 140. What is a JSON index?

"MySQL doesn't support indexing the entire JSON column, but you can index **specific JSON path expressions** using **functional indexes** (MySQL 8.0+).

```sql
-- Create functional index on a JSON field:
CREATE INDEX idx_json_city
ON users ((address->>'$.city'));

-- Query uses the index:
SELECT * FROM users WHERE address->>'$.city' = 'Mumbai';
```

Without this, any query filtering on a JSON path requires a full table scan.

I use JSON indexes to get relational-level query performance on specific frequently-queried JSON attributes."

#### Indepth
JSON indexes in MySQL are implemented as **virtual generated column + B+Tree index** pairs. The database extracts the JSON path value into a hidden virtual column and indexes that. Multiple JSON paths can be indexed independently. The type of the indexed expression matters: `->>'$.city'` extracts as a string (unquoted), while `->'$.city'` includes JSON quotes. Ensure the index expression exactly matches the query expression for the optimizer to recognize the index as applicable.

---

### 141. How do you validate JSON data?

"MySQL's `JSON_VALID()` function returns 1 if a string is valid JSON, 0 otherwise.

```sql
-- Validate before insert:
SELECT JSON_VALID('{"name": "Dhruv", "age": 30}');  -- Returns 1
SELECT JSON_VALID('not json');                         -- Returns 0

-- Filter valid JSON from a mixed column:
SELECT * FROM events WHERE JSON_VALID(payload) = 0;
```

The `JSON` data type itself enforces validity at insert time — MySQL rejects invalid JSON with an error. But for TEXT/VARCHAR columns storing JSON, `JSON_VALID()` provides runtime validation."

#### Indepth
Beyond structural validity, **schema validation** (checking required fields, types, value ranges) isn't natively enforced by MySQL. Approaches: (1) CHECK constraints with JSON_EXTRACT expressions (MySQL 8.0.16+): `CHECK (JSON_EXTRACT(data, '$.age') > 0)`. (2) BEFORE INSERT triggers that validate JSON structure. (3) Application-layer validation before insert (most common and flexible). MySQL's JSON functions are rich: `JSON_SCHEMA_VALID()` in MySQL 8.0 validates JSON against a JSON Schema definition.

---

### 142. What is a generated (virtual) column?

"A **generated column** is a column whose value is automatically computed from an expression involving other columns — you don't manually set it.

```sql
CREATE TABLE orders (
  price DECIMAL(10,2),
  tax_rate DECIMAL(5,4),
  total AS (price * (1 + tax_rate)) VIRTUAL
);
```

**Virtual**: Computed on-the-fly at read time, not stored.
**Stored**: Computed and physically persisted on write.

I use virtual columns for derived values (full_name from first+last, total from price+tax) and especially for JSON path extraction with indexing."

#### Indepth
Generated columns enable a powerful pattern: **indexing computed expressions without application changes**. Add a generated column + index to an existing table, and queries using that expression automatically benefit from the index — no application code changes needed. For JSON, `ALTER TABLE t ADD COLUMN city VARCHAR(100) AS (address->>'$.city') VIRTUAL, ADD INDEX idx_city (city)` instantly makes `WHERE address->>'$.city' = 'X'` index-efficient.

---

### 143. What is the difference between virtual and stored generated columns?

"Both compute values from expressions, but differ in storage:

| Feature | VIRTUAL | STORED |
|---|---|---|
| Storage | Only computed at read time | Physically stored on disk |
| Storage overhead | ❌ None | ✅ Uses disk space |
| Write speed | ✅ No overhead on write | ❌ Recomputed on every write |
| Indexable | ✅ Yes | ✅ Yes |
| Use in WHERE without index | Computed per-row during scan | Available directly in the row |

I use **VIRTUAL** by default — it saves storage and write overhead. I use **STORED** only when the expression is expensive to compute repeatedly or when the column is used heavily in full table scans."

#### Indepth
**VIRTUAL columns are not stored in the clustered index** — they're computed on-the-fly. This means InnoDB reads the source columns from the row, computes the expression, and returns the result for each row. For indexed access, the virtual column's index stores the computed value — enabling index-range scans without computing the expression at scan time. **STORED columns** are in the clustered index, so they're available during full scans without recomputation — useful for expensive JSON parsing expressions queried frequently without indexes.

---

### 144. What is a full table scan?

"A **full table scan** (access type `ALL` in EXPLAIN) means MySQL reads **every row** in the table to find matching records. It's the fallback when no suitable index exists or the optimizer decides an index scan is more expensive.

For a 1,000-row table, a full scan is fine. For a 100-million-row table, it's catastrophic — potentially reading gigabytes of data and locking significant buffer pool space.

EXPLAIN's `type = ALL` is the clearest warning sign. I treat every `type = ALL` on a large table as a bug requiring immediate indexing."

#### Indepth
A full table scan isn't always wrong. MySQL's optimizer chooses a full scan over an index scan when: the table is tiny (a few hundred rows — scanning is faster than index overhead), or the index selectivity is very low (the index would return most of the table anyway). `FORCE INDEX (idx)` overrides this, but may actually be slower. Always verify with `EXPLAIN ANALYZE` and real execution times whether a forced index genuinely improves performance or the optimizer was actually correct.

---

### 145. What causes index fragmentation?

"**Index fragmentation** occurs when B+Tree index pages are incompletely filled due to random insertions, deletions, and updates.

Causes:
- **Page splits**: Inserting a row into a full B+Tree page splits it — new page is created, old page is ~50% full.
- **Deletions**: Deleted rows leave empty space in pages that isn't immediately reclaimed.
- **Random insertions**: UUID primary keys cause continuous page splits since each insert goes to a random position in the tree.

Fragmentation increases storage usage, reduces page cache efficiency, and slows range scans."

#### Indepth
View index fragmentation via `information_schema.INNODB_TABLESPACES` and `SHOW TABLE STATUS`. The `Data_free` column shows unreclaimed space in the tablespace — a rough fragmentation indicator. Run `ANALYZE TABLE` to refresh index statistics without rebuilding. To physically defragment and reclaim space: `OPTIMIZE TABLE tablename` (performs table rebuild in-place). Use `ALTER TABLE t ENGINE=InnoDB` for online rebuilding in MySQL 5.6+. For online defragmentation without downtime, use `pt-online-schema-change`.

---

### 146. How do you rebuild indexes?

"To rebuild indexes in MySQL:

```sql
-- Rebuild all indexes on a table (blocks writes, deprecated for large tables):
OPTIMIZE TABLE tablename;

-- Online rebuild (minimal locking, MySQL 5.6+):
ALTER TABLE tablename ENGINE=InnoDB;

-- Rebuild specific index online:
ALTER TABLE tablename DROP INDEX idx_name, ADD INDEX idx_name (col);
-- With minimal lock:
ALTER TABLE tablename DROP INDEX idx_name, ADD INDEX idx_name (col) ALGORITHM=INPLACE, LOCK=NONE;
```

I use `OPTIMIZE TABLE` for small tables. For large production tables, I use `pt-online-schema-change` to avoid lengthy table locks."

#### Indepth
**Why rebuild indexes?** Over time, heavy write workloads cause B+Tree page fragmentation — pages are half-empty due to splits and deletes. Rebuilding compacts pages to full capacity, improving index density and scan efficiency. Modern SSDs make this less impactful than on spinning disks (random I/O is less costly), but rebuild still improves buffer pool efficiency. Schedule rebuilds during low-traffic windows and monitor `Data_free` in `SHOW TABLE STATUS` to identify tables that genuinely need it.

---

### 147. What is ANALYZE TABLE?

"`ANALYZE TABLE` updates the **index statistics** that MySQL's optimizer uses to estimate query costs and row counts.

```sql
ANALYZE TABLE users;
ANALYZE TABLE orders, products;  -- Multiple tables
```

After bulk data changes or if EXPLAIN shows wildly inaccurate row estimates, run `ANALYZE TABLE` to refresh statistics.

InnoDB runs automatic statistics updates (`innodb_stats_auto_recalc = ON` by default) when 10% of table rows change. `ANALYZE TABLE` forces an immediate full recalculation."

#### Indepth
`ANALYZE TABLE` in InnoDB doesn't read ALL rows — it samples `innodb_stats_persistent_sample_pages` (default 20) index leaf pages and estimates cardinality from the sample. For very large tables with non-uniform data distributions, increase this value for more accurate statistics (at the cost of longer `ANALYZE TABLE` runtime). Accurate statistics are fundamental to query plan quality — the optimizer's cost estimates are only as good as the statistics it uses.

---

### 148. What is OPTIMIZE TABLE?

"`OPTIMIZE TABLE` **physically defragments and rebuilds** a table, reclaiming unused space, compacting index pages, and sorting the clustered index.

```sql
OPTIMIZE TABLE users;
```

Effects:
- Reclaims space from deleted/updated rows.
- Rebuilds B+Tree indexes to full density.
- For InnoDB: equivalent to exporting + re-importing data.
- Resets `AUTO_INCREMENT` counter if rows were deleted? No — it doesn't reset AUTO_INCREMENT.

I schedule `OPTIMIZE TABLE` for tables with heavy DELETE/UPDATE workloads after bulk cleanup operations."

#### Indepth
`OPTIMIZE TABLE` on InnoDB with `innodb_file_per_table=ON` rebuilds the table in-place and creates a new `.ibd` file with compacted data. It's NOT online — it holds a **metadata lock** for the entire duration on older MySQL. Use `pt-online-schema-change` or `ALTER TABLE ... ALGORITHM=INPLACE` for production tables that can't tolerate downtime. For MyISAM, `OPTIMIZE TABLE` is simpler and also repairs the table file. Monitor disk space: `OPTIMIZE TABLE` temporarily needs up to 2× the table's current disk space.

---

### 149. What is metadata locking?

"**Metadata Locking (MDL)** is MySQL's mechanism to prevent concurrent DDL and DML operations from conflicting on the same database object.

Shared MDL (S): Acquired by DML statements (SELECT, INSERT, etc.) — multiple readers can coexist.
Exclusive MDL (X): Acquired by DDL statements (ALTER TABLE, DROP TABLE) — no concurrent operations allowed.

DDL waits for all active DML transactions to release their shared MDL before acquiring the exclusive lock. This can cause a **MDL lock queue buildup** — subsequent DML is blocked behind the waiting DDL."

#### Indepth
The MDL queue buildup is a common production incident: a long-running transaction holds S-MDL on `orders`, an `ALTER TABLE orders ADD COLUMN ...` waits for X-MDL, and all subsequent `SELECT`s on `orders` queue behind the waiting DDL because the DDL is already in the queue ahead of them. The system appears to "freeze". Solution: `SHOW PROCESSLIST` + `KILL` the blocking transaction, or use online DDL tools (GitHub's `gh-ost`) that are designed to work around MDL by using binlog-based table copying.

---

### 150. What is implicit commit in MySQL?

"An **implicit commit** occurs when MySQL automatically commits any open transaction before executing certain statements — without an explicit `COMMIT` command.

Statements that cause implicit commit include ALL DDL statements:
```sql
START TRANSACTION;
UPDATE users SET name = 'Dhruv';   -- Not yet committed
CREATE TABLE temp_t (...);         -- Implicit commit here!
-- The UPDATE is now permanently committed
ROLLBACK;  -- Does nothing — already committed
```

This is a critical pitfall: DDL inside a transaction silently makes previous DML permanent."

#### Indepth
MySQL's documentation lists all statements that cause implicit commit: DDL (`CREATE`, `ALTER`, `DROP`, `TRUNCATE`, `RENAME`), account management (`CREATE USER`, `GRANT`, `REVOKE`), explicitly invoking `BEGIN` while inside a transaction, and others. This differs fundamentally from PostgreSQL where DDL is transactional. In MySQL, never assume DDL can be rolled back. Migration scripts must account for this — handle DDL and DML separately, and use pre-migration database snapshots as the rollback strategy.

---

### 151. What is autocommit?

"**Autocommit** is a MySQL session setting where each SQL statement is automatically treated as its own transaction and immediately committed.

By default: `autocommit = 1` (ON).

With autocommit ON: every `INSERT`, `UPDATE`, `DELETE` commits instantly. No explicit `START TRANSACTION` needed for single statements.

With autocommit OFF (`SET autocommit = 0`): statements accumulate until you `COMMIT` or `ROLLBACK`. Like having a persistent open transaction.

Most application frameworks disable autocommit and manage transactions explicitly."

#### Indepth
A common mistake: developers set `autocommit = 0` expecting it to behave like `START TRANSACTION`, but there's a subtle difference. `autocommit = 0` means MySQL opens a transaction implicitly for the first DML statement. This implicit transaction persists open until explicitly committed — connections returned to connection pools with an open transaction cause data leaks and unexpected behavior in the next borrower. Always ensure explicit `COMMIT` or `ROLLBACK` before returning connections to the pool.

---

### 152. What happens when autocommit is OFF?

"When `SET autocommit = 0`, MySQL starts an implicit transaction at the first DML statement. All subsequent DML statements are part of that transaction until you explicitly:
- `COMMIT` — permanently save all changes.
- `ROLLBACK` — discard all changes since the last commit.

Without explicit `COMMIT` or `ROLLBACK`, the transaction remains open indefinitely.

An open transaction holds row locks and MVCC undo log entries. Long-running open transactions with autocommit OFF are a major cause of **replication lag, lock contention, and InnoDB undo log bloat**."

#### Indepth
InnoDB's **undo log** stores the 'before' image of modified rows for MVCC and rollback. As long as a transaction is open, its undo log entries cannot be purged — even if all other transactions have moved on. With autocommit OFF and a forgotten `COMMIT`, the undo log grows without bound, eventually filling the tablespace. Monitor `SHOW ENGINE INNODB STATUS` for 'History list length' — when it exceeds millions, long-running uncommitted transactions are likely the cause.

---

### 153. What is SAVEPOINT?

"A **SAVEPOINT** creates a named point within a transaction that you can partially roll back to, without rolling back the entire transaction.

```sql
START TRANSACTION;
INSERT INTO orders VALUES (1, 'Item A');
SAVEPOINT sp1;
INSERT INTO orders VALUES (2, 'Item B');
SAVEPOINT sp2;
INSERT INTO orders VALUES (3, 'Item C');

ROLLBACK TO SAVEPOINT sp2;  -- Undoes 'Item C', keeps A and B
COMMIT;                      -- Commits A and B permanently
```

I use savepoints in complex batch operations where I want to retry individual sub-steps on failure without discarding the entire batch."

#### Indepth
Savepoints are particularly useful in **stored procedures** for nested error handling — rolling back individual procedure sections while preserving outer transaction progress. SAVEPOINTs are released automatically on `COMMIT` or full `ROLLBACK`. They're also released when a DDL statement causes an implicit commit (losing the savepoint's rollback capability). Savepoints don't provide isolation — concurrent transactions still see only committed data, not the partially-complete savepoint state.

---

### 154. What is XA transaction?

"An **XA transaction** is MySQL's implementation of the **Two-Phase Commit (2PC) protocol** for distributed transactions spanning multiple database systems.

Two phases:
1. **Prepare**: All participants (MySQL instances, other databases) confirm they can commit.
2. **Commit**: If ALL prepared successfully, all commit. If ANY fail, all roll back.

```sql
XA START 'txn1';
UPDATE orders SET status = 'shipped';
XA END 'txn1';
XA PREPARE 'txn1';   -- Phase 1
XA COMMIT 'txn1';    -- Phase 2
```"

#### Indepth
XA transactions in MySQL have significant **limitations**: they cannot be used with statement-based replication (requires ROW format), they have higher overhead than regular transactions, and MySQL's XA implementation is not crash-safe enough for most production distributed transaction requirements. Most modern distributed systems avoid XA in favor of application-level patterns like **Saga orchestration**, **outbox pattern**, or **compensating transactions** — which tolerate partial failures gracefully without requiring 2PC coordination.

---

### 155. What is an event scheduler?

"The MySQL **Event Scheduler** is a built-in job scheduler that runs SQL statements automatically at scheduled times — similar to cron for the database.

```sql
-- Create a daily cleanup event:
CREATE EVENT daily_cleanup
ON SCHEDULE EVERY 1 DAY STARTS '2024-01-01 02:00:00'
DO DELETE FROM sessions WHERE expires_at < NOW();

-- Start the scheduler:
SET GLOBAL event_scheduler = ON;
```

I use events for scheduled maintenance: purging old records, refreshing aggregate tables, rotating logs, sending database-triggered notifications."

#### Indepth
The event scheduler runs in a **dedicated thread** (`event_scheduler` thread visible in `SHOW PROCESSLIST`). Events run with the privileges of the `DEFINER` user. In an HA/replication setup, events run on ALL servers (primary + all replicas) by default since they read from the `mysql.event` table which is replicated. To prevent events from running on replicas: set `event_scheduler = OFF` on replicas and `ON` only on the primary. When the primary fails over, enable the event scheduler on the new primary.

---

### 156. How does MySQL handle large TEXT/BLOB data?

"InnoDB stores TEXT/BLOB data **inline** if it fits within the page (16KB), or **off-page** in separate overflow pages if it exceeds the inline threshold.

With `DYNAMIC` row format (default in MySQL 5.7+): all TEXT/BLOB data beyond a few hundred bytes is stored off-page. The main row contains only a 20-byte pointer.

Large BLOB/TEXT columns should be stored in separate tables if they're not always needed — avoids loading them into memory on every row fetch for queries that don't need them."

#### Indepth
Off-page storage creates an additional I/O for every row that has a BLOB/TEXT value — even `SELECT *` triggers an off-page read for each non-NULL BLOB column. This is why `SELECT *` on tables with BLOB columns is especially costly. Best practice: store BLOB/TEXT data in a dedicated `attachments` or `content` table joined by ID. Only JOIN when the content is actually needed. For very large files (>16MB), consider storing them in object storage (S3) and keeping only the file path/URL in MySQL.

---

### 157. What is the maximum row size in MySQL?

"The maximum row size in InnoDB (with `DYNAMIC` or `COMPRESSED` row format) is technically unlimited for TEXT/BLOB data (stored off-page), but the **inline row size is limited to 65,535 bytes** for VARCHAR, CHAR, and similar inline types.

For practical purposes: the maximum **in-page** row size is approximately 8,000 bytes (half of a 16KB page), beyond which InnoDB starts storing columns off-page even for shorter variable-length types.

Creating a table with too many wide `VARCHAR` columns exceeds row size limits and raises `ERROR 1118: Row size too large`."

#### Indepth
The 65,535 byte limit applies to the **entire row's inline storage** across all columns. A table with 100 × `VARCHAR(655)` columns would theoretically hit this limit. Using `TEXT` or `BLOB` instead of `VARCHAR` for large data bypasses the inline limit since they're always stored off-page beyond the first few bytes. When you hit `ERROR 1118`, convert wide VARCHAR columns to TEXT (if data can exceed ~16KB) or restructure by splitting rarely-used large columns into a secondary table.

---

### 158. What are temporary tables stored in memory vs disk?

"MySQL's internal temporary tables (created for GROUP BY, DISTINCT, UNION, filesort) start in **memory** using the `TempTable` storage engine (MySQL 8.0+), then **spill to disk** if they exceed the memory limit.

Memory limit: `tmp_table_size` AND `max_heap_table_size` (whichever is smaller, default 16MB each).
```sql
SHOW VARIABLES LIKE 'tmp_table_size';
SHOW STATUS LIKE 'Created_tmp_disk_tables';  -- How many spilled to disk
```

Memory temp tables are fast. Disk temp tables (tmpdir on disk) are significantly slower."

#### Indepth
MySQL 8.0 replaced the old MEMORY engine for internal temp tables with `TempTable` (configurable with `internal_tmp_mem_storage_engine`). TempTable uses `temptable_max_ram` (default 1GB) and can now spill to memory-mapped files (`temptable_use_mmap`) before writing to disk — a middle ground faster than actual disk I/O. Monitor `Created_tmp_disk_tables` / `Created_tmp_tables` ratio — if disk tables exceed 10% of total temp tables, increase `tmp_table_size` or optimize queries to reduce temp table size.

---

### 159. What is a hash index?

"A **hash index** stores a hash of the indexed key value in a hash table, enabling O(1) exact-match lookups — faster than B+Tree's O(log n) for point queries.

MySQL's **MEMORY** storage engine supports explicit hash indexes.

InnoDB does NOT support user-created hash indexes, but it has an internal **Adaptive Hash Index (AHI)**: automatically builds a hash index in memory for frequently accessed B+Tree pages, transparently accelerating hot access patterns."

#### Indepth
Hash indexes have critical limitations: they only support **equality comparisons** (`=`, `IN`). They cannot support range queries (`>`, `<`, `BETWEEN`), sorting (`ORDER BY`), or prefix matching (`LIKE 'Jo%'`). This makes them unsuitable for most MySQL workloads beyond pure key-value lookups. InnoDB's Adaptive Hash Index (AHI) is entirely automatic — you can monitor its effectiveness with `SHOW ENGINE INNODB STATUS` under the "INSERT BUFFER AND ADAPTIVE HASH INDEX" section. On workloads with poor AHI hit rates, disable it to reclaim buffer pool memory.

---

### 160. What is the difference between BTREE and HASH index?

"| Feature | B+Tree Index | Hash Index |
|---|---|---|
| Lookup type | Equality + Range | Equality only |
| ORDER BY support | ✅ Yes | ❌ No |
| LIKE prefix support | ✅ Yes | ❌ No |
| Complexity | O(log n) | O(1) average |
| Storage engine | InnoDB, MyISAM, Memory | Memory (explicit), InnoDB (internal AHI only) |
| Main use | General purpose | In-memory key-value caching |

B+Tree is the universal choice. Hash is only appropriate for tiny lookup tables held entirely in memory (MEMORY engine), or for Redis-style caching, not MySQL production tables."

#### Indepth
When creating an index on a MEMORY table: `CREATE INDEX idx ON t(col) USING HASH` creates a hash index; `USING BTREE` creates a B+Tree. The MEMORY engine actually supports both. In production MySQL (InnoDB), all user-created indexes use B+Tree. If you need hash index semantics (O(1) key-value lookups), use **Redis** or **Memcached** as a dedicated caching layer rather than the MEMORY engine, which is volatile (data lost on restart) and can't scale beyond a single server's RAM.

---
