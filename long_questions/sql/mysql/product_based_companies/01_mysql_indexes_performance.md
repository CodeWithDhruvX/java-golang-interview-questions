# ⚡ 01 — MySQL Indexes & Performance
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Medium–Hard

> Indexes and query optimization — the **most differentiated topic** at **Flipkart, Paytm, Swiggy, Amazon, Google**.

---

## 🔑 Must-Know Topics
- B-Tree index internals
- Composite (multi-column) indexes and prefix rule
- Covering indexes
- EXPLAIN output interpretation
- Query optimization strategies
- Slow query log
- Index types: B-Tree, Hash, Full-Text, Spatial

---

## ❓ Most Asked Questions

### Q1. What is an index? How does a B-Tree index work?

```sql
-- Create various index types
CREATE TABLE employees (
    id         INT PRIMARY KEY,          -- clustered B-Tree index
    email      VARCHAR(255),
    last_name  VARCHAR(100),
    first_name VARCHAR(100),
    department VARCHAR(50),
    salary     DECIMAL(10,2),
    created_at DATETIME
);

CREATE INDEX idx_email      ON employees(email);              -- single column
CREATE INDEX idx_name       ON employees(last_name, first_name); -- composite
CREATE INDEX idx_dept_sal   ON employees(department, salary);
CREATE UNIQUE INDEX uq_email ON employees(email);
```

> **How B-Tree works:** All leaf nodes store the actual data pointer (in InnoDB: the primary key value). Internal nodes store separator keys for navigation. B-Tree allows range scans (`>`, `<`, `BETWEEN`, `LIKE 'prefix%'`) unlike hash indexes. Height of tree is typically 3–4 levels for millions of rows.

---

### Q2. What is a clustered vs non-clustered (secondary) index?

```sql
-- In InnoDB:
-- PRIMARY KEY = clustered index (data IS the index — leaf nodes contain row data)
-- All other indexes = secondary (non-clustered) — leaf nodes contain PK value

-- When you query via secondary index:
-- 1. Walk the secondary B-Tree → find PK value
-- 2. Walk the PRIMARY key B-Tree → find actual row data
-- This double-lookup is called "secondary index lookup" or "bookmark lookup"

-- Example: query using secondary index on email
EXPLAIN SELECT * FROM employees WHERE email = 'alice@example.com';
-- key: idx_email, Extra: Using index condition
-- MySQL walks idx_email → gets id → fetches row from PK index

-- InnoDB always needs a primary key. Without it:
-- MySQL creates a hidden 6-byte rowid as the clustered key
-- This makes secondary indexes use the hidden rowid — not user-visible
```

---

### Q3. What is a covering index?

```sql
-- A covering index contains ALL columns needed by a query
-- MySQL can satisfy the query entirely from the index — no table lookup needed
-- EXPLAIN shows: Extra = "Using index"

CREATE INDEX idx_covering ON employees(department, salary, name);

-- This query is fully covered:
SELECT name, salary
FROM employees
WHERE department = 'Engineering'
ORDER BY salary DESC;
-- MySQL reads idx_covering only — never touches the main table rows

-- vs non-covering (needs row lookup for 'bio' column not in index):
SELECT name, salary, bio
FROM employees
WHERE department = 'Engineering';
-- Extra = "Using index condition" (not "Using index") — extra table lookup needed

-- Covering indexes eliminate the most expensive part of index queries: row lookups
```

---

### Q4. What is the leftmost prefix rule for composite indexes?

```sql
-- Composite index on (A, B, C) is usable for:
-- • WHERE A = ?                         ✅
-- • WHERE A = ? AND B = ?               ✅
-- • WHERE A = ? AND B = ? AND C = ?     ✅
-- • WHERE A = ? AND C = ?               ✅ (uses A, skips C)
-- • WHERE B = ?                         ❌ (skips A — cannot start mid-index)
-- • WHERE B = ? AND C = ?               ❌

CREATE INDEX idx_dept_salary_name ON employees(department, salary, name);

EXPLAIN SELECT * FROM employees WHERE department = 'Eng';           -- ✅ uses idx
EXPLAIN SELECT * FROM employees WHERE department = 'Eng' AND salary > 50000; -- ✅
EXPLAIN SELECT * FROM employees WHERE salary > 50000;               -- ❌ full scan
EXPLAIN SELECT * FROM employees WHERE department = 'Eng' ORDER BY salary; -- ✅ idx+sort
EXPLAIN SELECT * FROM employees
    WHERE department = 'Eng'
    ORDER BY name;   -- ⚠️ filesort — salary column is skipped in ORDER BY
```

---

### Q5. How do you read EXPLAIN output?

```sql
EXPLAIN SELECT e.name, d.department_name
FROM employees e
JOIN departments d ON e.department_id = d.id
WHERE e.salary > 60000
ORDER BY e.name;

-- Key EXPLAIN columns:
-- id:           Query step number (same id → parallel join)
-- select_type:  SIMPLE, PRIMARY, SUBQUERY, DERIVED, UNION
-- table:        Which table this row refers to
-- type:         Access method — from best to worst:
--               system > const > eq_ref > ref > range > index > ALL
-- possible_keys: Indexes that COULD be used
-- key:          Index actually chosen by optimizer
-- key_len:      Bytes of the index used (longer = more columns used)
-- ref:          What is being compared to the index column
-- rows:         Estimated rows examined (lower = better!)
-- filtered:     % of rows passing WHERE condition (higher = better)
-- Extra:        Additional info — critical flags:
--   Using index          → covering index (fast!)
--   Using where          → post-index filtering
--   Using filesort       → sort in memory/disk (add index on ORDER BY col)
--   Using temporary      → temp table for GROUP BY/DISTINCT (expensive!)
--   Using index condition → Index Condition Pushdown (ICP, good)
```

---

### Q6. What causes a "full table scan" and how do you fix it?

```sql
-- Causes of full scans (type=ALL in EXPLAIN):

-- 1. No index on WHERE column
SELECT * FROM orders WHERE status = 'pending';  -- no index on status → full scan
CREATE INDEX idx_status ON orders(status);

-- 2. Function on indexed column — breaks index usage
SELECT * FROM employees WHERE YEAR(created_at) = 2024;  -- ❌ full scan
-- Fix: range query avoids function on column
SELECT * FROM employees
WHERE created_at >= '2024-01-01' AND created_at < '2025-01-01';  -- ✅ uses index

-- 3. Leading wildcard LIKE
SELECT * FROM products WHERE name LIKE '%phone%';  -- ❌ full scan (leading %)
SELECT * FROM products WHERE name LIKE 'phone%';   -- ✅ uses index (trailing % only)

-- 4. Type mismatch (implicit conversion)
-- If email column is VARCHAR but you compare to INT:
SELECT * FROM users WHERE id = '123abc';   -- implicit cast → may not use index

-- 5. Low cardinality column (optimizer decides scan is cheaper than index)
-- Indexes on boolean or yes/no columns often not used
```

---

### Q7. What is the slow query log? How do you enable it?

```sql
-- Enable slow query log
SET GLOBAL slow_query_log = 'ON';
SET GLOBAL long_query_time = 1;          -- log queries > 1 second
SET GLOBAL slow_query_log_file = '/var/log/mysql/slow.log';
SET GLOBAL log_queries_not_using_indexes = 'ON';  -- also log no-index queries

-- Check current settings
SHOW VARIABLES LIKE 'slow_query%';
SHOW VARIABLES LIKE 'long_query_time';

-- Analyze slow query log with mysqldumpslow
-- mysqldumpslow -s t -t 10 /var/log/mysql/slow.log
-- -s t: sort by query time, -t 10: top 10

-- Use EXPLAIN to analyze a slow query
EXPLAIN FORMAT=JSON
SELECT * FROM orders o
JOIN customers c ON o.customer_id = c.id
WHERE o.total > 500
ORDER BY o.created_at DESC;
```

---

### Q8. What is the difference between INDEX and UNIQUE INDEX? When to use each?

```sql
-- Regular INDEX — allows duplicate values, just speeds up lookups
CREATE INDEX idx_department ON employees(department);  -- 'Engineering' can appear N times

-- UNIQUE INDEX — enforces uniqueness + speeds up lookups
CREATE UNIQUE INDEX uq_email ON users(email);  -- no two rows can have same email

-- PRIMARY KEY — special UNIQUE index, also the clustered index, no NULLs
ALTER TABLE users ADD PRIMARY KEY (id);

-- Composite UNIQUE — combination must be unique
CREATE UNIQUE INDEX uq_product_region ON prices(product_id, region_id);
-- (1, 'US') and (1, 'EU') are both allowed — combination is unique check

-- When to use: add UNIQUE when the column/combo MUST be unique by business rule
-- The UNIQUE constraint is enforced at DB level, not just app level

-- Check existing indexes
SHOW INDEX FROM employees;
SHOW CREATE TABLE employees;  -- shows all indexes
```

---

### Q9. How do you use FORCE INDEX / USE INDEX hints?

```sql
-- Optimizer doesn't always pick the best index — you can hint it
-- USE INDEX: suggest an index (optimizer still may ignore)
SELECT * FROM employees
USE INDEX (idx_dept_salary)
WHERE department = 'Engineering' AND salary > 70000;

-- FORCE INDEX: force the optimizer to use this index
SELECT * FROM orders
FORCE INDEX (idx_created_at)
WHERE customer_id = 100
ORDER BY created_at DESC
LIMIT 20;

-- IGNORE INDEX: exclude a specific index from consideration
SELECT * FROM products
IGNORE INDEX (idx_category)
WHERE category = 'Electronics';

-- When useful:
-- 1. Optimizer picks wrong index due to stale statistics → run ANALYZE TABLE first
-- 2. You KNOW from testing that a specific index performs best

ANALYZE TABLE employees;  -- update index statistics so optimizer makes better choices
```

---

### Q10. What is index cardinality and why does it matter?

```sql
-- Cardinality = number of unique values in an indexed column
-- High cardinality = index is selective (useful)
-- Low cardinality = index is not selective (optimizer may ignore)

-- Check cardinality
SHOW INDEX FROM employees;
-- Cardinality column shows estimated unique values

-- Example:
-- gender column: 2 unique values (M/F) → very low cardinality → index rarely useful
-- email column: millions of unique values → very high cardinality → index very useful
-- department: ~10 unique values → medium cardinality → may help for rare departments

-- For low-cardinality columns, combine them into composite index:
CREATE INDEX idx_dept_status ON employees(department, employment_status);
-- (Engineering, active): more selective than either alone

-- Update cardinality stats:
ANALYZE TABLE employees;  -- collects fresh statistics
-- Or use:
OPTIMIZE TABLE employees; -- also defragments table (causes table rebuild)
```

---

### Q11. How do you manage index bloat and maintenance?

```sql
-- Check index sizes
SELECT
    table_name,
    index_name,
    ROUND(stat_value * @@innodb_page_size / 1024 / 1024, 2) AS size_mb
FROM mysql.innodb_index_stats
WHERE database_name = 'myapp'
  AND stat_name = 'size'
ORDER BY size_mb DESC;

-- Find unused indexes (MySQL 8.0 performance_schema)
SELECT object_schema, object_name, index_name, count_star
FROM performance_schema.table_io_waits_summary_by_index_usage
WHERE index_name IS NOT NULL
  AND count_star = 0
  AND object_schema NOT IN ('mysql', 'information_schema', 'performance_schema')
ORDER BY object_schema, object_name;

-- Drop unused index
DROP INDEX idx_old_column ON employees;

-- Rebuild fragmented index
ALTER TABLE employees ENGINE=InnoDB;  -- rebuilds table (pt-online-schema-change for large tables)
OPTIMIZE TABLE employees;             -- reclaims space from deleted rows
```

---

### Q12. What are full-text indexes in MySQL?

```sql
-- Full-text index — for natural language search within TEXT/VARCHAR columns
CREATE TABLE articles (
    id      INT PRIMARY KEY,
    title   VARCHAR(255),
    body    TEXT,
    FULLTEXT INDEX ft_content (title, body)  -- full-text index
);

-- MATCH ... AGAINST — natural language mode (default)
SELECT id, title,
    MATCH(title, body) AGAINST ('mysql performance') AS relevance_score
FROM articles
WHERE MATCH(title, body) AGAINST ('mysql performance')
ORDER BY relevance_score DESC;

-- Boolean mode — supports +, -, *, "phrase", ~
SELECT * FROM articles
WHERE MATCH(title, body) AGAINST ('+mysql -oracle "query optimization"' IN BOOLEAN MODE);

-- Full-text vs LIKE '%keyword%':
-- LIKE '%keyword%' → always full scan
-- FULLTEXT → uses specialized inverted index, much faster
-- Full-text NOT suitable for: exact match lookups, range queries, short strings
```

