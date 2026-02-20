# 🟢 Additional Beginner Level Questions (Q101–120)

---

### 101. What is the difference between MySQL and SQL?

"**SQL** (Structured Query Language) is a standardized language for interacting with relational databases — defining schemas, querying data, controlling access. It's a language specification, not a product.

**MySQL** is a specific **database management system (DBMS)** that implements SQL. It's the software server that stores and manages data, and you use SQL to communicate with it.

Think of it like English vs a specific book written in English. SQL is the language; MySQL is one product that speaks that language."

#### Indepth
While MySQL implements the core SQL standard (ISO/IEC 9075), it also includes **proprietary extensions** — like `AUTO_INCREMENT`, `REPLACE INTO`, `SHOW` commands, and MySQL-specific functions (`GROUP_CONCAT`, `IFNULL`). These extensions make code non-portable across databases. When writing SQL that needs to work on both MySQL and PostgreSQL, stick to the standard subset and avoid MySQL-specific syntax.

---

### 102. What is the difference between SQL and NoSQL?

"**SQL** databases are relational, use structured tables with schemas, and excel at complex queries and transactions. Examples: MySQL, PostgreSQL, Oracle.

**NoSQL** databases use flexible data models (document, key-value, graph, column-family), sacrifice some query power for scalability, and are schema-flexible. Examples: MongoDB (document), Redis (key-value), Cassandra (wide-column).

I choose SQL for structured data with complex relationships and ACID requirements. I choose NoSQL for unstructured, high-volume, high-velocity data where schema flexibility and horizontal scaling matter more."

#### Indepth
The **CAP theorem** often drives the choice: SQL databases typically prioritize Consistency and Partition tolerance (CP). Many NoSQL systems prioritize Availability and Partition tolerance (AP). 'BASE' (Basically Available, Soft state, Eventually consistent) describes NoSQL systems that trade strong consistency for availability. Modern distributed SQL databases (Spanner, CockroachDB) attempt to provide ACID guarantees at global scale — blurring the SQL/NoSQL divide.

---

### 103. What are DDL, DML, DCL, and TCL commands?

"SQL commands are categorized by purpose:

- **DDL** (Data Definition Language): Define/modify schema. `CREATE`, `ALTER`, `DROP`, `TRUNCATE`, `RENAME`.
- **DML** (Data Manipulation Language): Manipulate data. `SELECT`, `INSERT`, `UPDATE`, `DELETE`.
- **DCL** (Data Control Language): Control access. `GRANT`, `REVOKE`.
- **TCL** (Transaction Control Language): Manage transactions. `COMMIT`, `ROLLBACK`, `SAVEPOINT`, `START TRANSACTION`.

I keep this taxonomy in mind because DDL statements in MySQL often cause implicit commits and cannot always be rolled back."

#### Indepth
In MySQL, **DDL statements cause implicit commit** — a `CREATE TABLE` or `DROP TABLE` inside an explicit transaction automatically commits the transaction first, making DDL irreversible. This is important when scripting migrations: accidental DDL mid-transaction doesn't roll back cleanly. Tools like Liquibase and Flyway handle migration transactions carefully for this reason. PostgreSQL supports transactional DDL (DDL can be rolled back), which makes its migration tooling safer.

---

### 104. What is the difference between PRIMARY KEY and UNIQUE KEY?

"Both enforce uniqueness, but with differences:

| Feature | PRIMARY KEY | UNIQUE KEY |
|---|---|---|
| NULLs allowed | ❌ No (NOT NULL enforced) | ✅ Yes (multiple NULLs allowed) |
| Per table | Exactly one | Multiple allowed |
| Clustered index | ✅ Yes (InnoDB) | ❌ No (secondary index) |
| Main purpose | Row identity | Enforce business uniqueness |

Example: `id` is a primary key. `email` is a unique key (a person might not have an email yet — NULL allowed)."

#### Indepth
InnoDB's clustered index behavior makes the PRIMARY KEY choice critical for performance. A UNIQUE KEY creates a separate secondary B+Tree index. When you have a natural candidate for primary key (like `email`), using a surrogate `INT AUTO_INCREMENT` PK + a UNIQUE constraint on `email` is generally better: the integer PK keeps the clustered index compact and efficient, while the unique constraint enforces business rules. String PKs increase all secondary index sizes proportionally.

---

### 105. What happens if you insert NULL into a PRIMARY KEY column?

"It fails with an error. Primary key columns are implicitly `NOT NULL`. Attempting to insert a NULL into a primary key column raises:

```
ERROR 1048 (23000): Column 'id' cannot be null
```

This is by design — a primary key must be able to uniquely identify every row. NULL means 'unknown', and an unknown identity is meaningless as a row identifier.

If you need to allow NULLs in an identifier-like column, use a UNIQUE key instead, which permits NULLs."

#### Indepth
Even without explicitly declaring `NOT NULL` on a primary key column, MySQL enforces it automatically. If you declare `id INT` and add `PRIMARY KEY (id)`, MySQL internally adds `NOT NULL` to `id`. This differs from UNIQUE keys where you must explicitly add `NOT NULL` if you want to disallow NULLs. Understanding this implicit constraint helps when reverse-engineering legacy schemas that may not show explicit `NOT NULL` in the DDL.

---

### 106. What is the difference between INT and BIGINT?

"Both store integers, but differ in size:

| Type | Storage | Signed Range | Unsigned Range |
|---|---|---|---|
| `INT` | 4 bytes | -2,147,483,648 to 2,147,483,647 (~2.1 billion) | 0 to ~4.3 billion |
| `BIGINT` | 8 bytes | -9.2 quintillion to +9.2 quintillion | 0 to 18.4 quintillion |

I use `INT UNSIGNED AUTO_INCREMENT` for most tables. I switch to `BIGINT` when a table may grow beyond ~2 billion rows or when storing Unix timestamps in milliseconds, epoch microseconds, or large financial values."

#### Indepth
Running out of `INT AUTO_INCREMENT` values is a real operational emergency. When an `INT UNSIGNED AUTO_INCREMENT` column hits 4,294,967,295, the next insert fails with `ERROR 1062 Duplicate entry`. **Proactively migrate** to `BIGINT` before hitting this limit. In MySQL 8.0, the AUTO_INCREMENT counter is persisted, so you can monitor it with `SHOW CREATE TABLE` or `information_schema.TABLES`. Set up alerting when `AUTO_INCREMENT` usage exceeds 75%.

---

### 107. What is ENUM data type?

"**ENUM** stores one value from a predefined list of permitted string values. Internally stored as an integer (1 byte for up to 255 values, 2 bytes for up to 65,535).

```sql
status ENUM('active', 'inactive', 'suspended') DEFAULT 'active'
```

Benefits: Natural constraint enforcement, compact storage, readable values.

I use ENUM for columns with a small fixed set of values: status fields, priority levels, categories. The major downside is that adding a new ENUM value requires an `ALTER TABLE` — avoid ENUM if the set of values changes frequently."

#### Indepth
`ALTER TABLE` to add an ENUM value was a **blocking DDL** operation before MySQL 5.7. Since 5.7, adding values at the **end** of the ENUM list is instant (no table rebuild). Adding a value in the middle or modifying existing values still requires a full table rebuild. For frequently changing value sets, use a `VARCHAR` with a foreign key reference to a lookup/reference table — it's more flexible and doesn't require schema changes when values change.

---

### 108. What is SET data type?

"**SET** is similar to ENUM but allows storing **multiple values** from the predefined list in a single column. Values are stored as a bitmask integer.

```sql
permissions SET('read', 'write', 'delete', 'admin') DEFAULT 'read'
-- Can store: 'read', 'read,write', 'read,write,delete', etc.
```

I rarely use SET in modern applications. While it's space-efficient conceptually, it's hard to query (requires `FIND_IN_SET()`), hard to index efficiently, and violates **1NF** (a single cell holds multiple values). A junction table (`user_permissions`) is almost always the better design."

#### Indepth
`SET` stores values as a **bitmask**: `read=1`, `write=2`, `delete=4`, `admin=8`. `'read,write'` stores as `3` (binary `0011`). The maximum is 64 elements. Querying is awkward: `WHERE FIND_IN_SET('write', permissions)` cannot use a B+Tree index efficiently. Bit operations (`permissions & 2`) can use function-based approaches but remain fragile. For access control systems, a proper RBAC (Role-Based Access Control) model with normalized tables is always preferred.

---

### 109. What is the difference between NOW() and SYSDATE()?

"Both return the current date and time, but they differ in **when** they evaluate:

- **`NOW()`**: Returns the time at the **start of the current statement** (or transaction in some contexts). Consistent within a statement.
- **`SYSDATE()`**: Returns the **actual current time** at the exact moment of evaluation. Can differ from `NOW()` within the same query if the query takes time.

I always use `NOW()` because its consistent behavior is more predictable. `SYSDATE()` is non-deterministic and can cause issues with statement-based replication."

#### Indepth
`SYSDATE()` is **non-deterministic** in replication. When replicating with `STATEMENT` format binlog, the replica re-executes the SQL statement. If the replica runs `SYSDATE()` at a slightly later time, it gets a different value than the primary had — causing **data divergence**. MySQL has a `sysdate-is-now` option to make `SYSDATE()` behave like `NOW()` for replication compatibility. Using `NOW()` from the start avoids this entirely.

---

### 110. What is the default storage engine in modern MySQL?

"Since MySQL 5.5, **InnoDB** has been the default storage engine, replacing MyISAM.

Any `CREATE TABLE` statement without specifying `ENGINE=...` uses InnoDB automatically.

InnoDB was chosen as the default because it provides ACID transactions, row-level locking, foreign key support, and crash recovery — essential features for production applications that MyISAM lacked.

I always explicitly write `ENGINE=InnoDB` in my DDL scripts for clarity and portability."

#### Indepth
To verify the default engine: `SHOW VARIABLES LIKE 'default_storage_engine';`. While InnoDB is the default, some internal MySQL system tables in very old MySQL versions still used MyISAM (e.g., the `mysql.user` table). As of MySQL 8.0, **all system tables use InnoDB**. This was a major milestone enabling atomic DDL — because InnoDB's transaction support means schema metadata changes can be committed atomically, preventing partially-applied DDL on crash.

---

### 111. What is case sensitivity in MySQL?

"Case sensitivity in MySQL depends on the context:

- **SQL keywords and functions**: Case-insensitive (`SELECT`, `select`, `Select` all work).
- **Database and table names**: Depends on the OS filesystem and `lower_case_table_names` setting. Case-sensitive on Linux (default), case-insensitive on Windows/macOS.
- **String comparisons**: Depend on the **collation**. `utf8mb4_unicode_ci` = case-insensitive. `utf8mb4_bin` = case-sensitive (binary comparison).

I set `lower_case_table_names=1` in production to ensure consistent behavior across Linux and non-Linux environments."

#### Indepth
`lower_case_table_names=1` stores all table names as lowercase. This is critical when moving databases between Linux (case-sensitive FS) and Windows/macOS (case-insensitive FS). Mismatched case sensitivity is a frequent source of bugs when developers use Windows locally and deploy to Linux production. Setting `lower_case_table_names=1` must be done **before** creating any databases — changing it after data exists can corrupt table references. Configure it at MySQL initialization time.

---

### 112. What is a constraint violation?

"A **constraint violation** occurs when an attempted data modification breaks a defined database constraint, causing the operation to fail with an error.

Common violations:
- `ERROR 1062` — UNIQUE/PRIMARY KEY duplicate value insertion.
- `ERROR 1048` — NOT NULL violation (inserting NULL into a NOT NULL column).
- `ERROR 1452` — FOREIGN KEY violation (child references non-existent parent).
- `ERROR 1265` — Data truncated (value exceeds column length or ENUM list).

Constraint violations are the database's safety mechanism. Rather than silently corrupting data, it rejects invalid operations."

#### Indepth
Application code should handle constraint violations gracefully. Specifically, `ERROR 1062` (duplicate key) during concurrent inserts is a normal occurrence that should be caught and handled (retry, return error to user, or use `INSERT ... ON DUPLICATE KEY UPDATE`). Crashing or logging as an unhandled error for expected constraint violations is poor application design. Factor constraint violations into your application's normal control flow rather than treating them as exceptional conditions.

---

### 113. What is the difference between LIKE and REGEXP?

"**LIKE**: Simple pattern matching with `%` (any chars) and `_` (one char). Fast, index-friendly for prefix patterns.

**REGEXP**: Full regular expression matching. Supports complex patterns like character classes, quantifiers, anchors.

```sql
-- LIKE: names starting with 'Jo'
WHERE name LIKE 'Jo%'

-- REGEXP: names starting with vowel
WHERE name REGEXP '^[AEIOUaeiou]'
```

I use LIKE for simple prefix/suffix patterns and REGEXP for complex validation patterns. REGEXP cannot use B+Tree indexes and is significantly slower than LIKE."

#### Indepth
MySQL's REGEXP uses the **ICU (International Components for Unicode)** library since MySQL 8.0, providing full Unicode-aware regex support. Functions include `REGEXP_LIKE()`, `REGEXP_REPLACE()`, `REGEXP_SUBSTR()`, `REGEXP_INSTR()`. REGEXP operations are **not index-friendly** — they always require scanning rows (sequential match). For complex text searches in production, consider FULLTEXT index or an external search engine. Never put REGEXP in a tight WHERE clause on a large unpartitioned table.

---

### 114. What is the difference between DISTINCT and GROUP BY?

"Both eliminate duplicates, but differently:

**DISTINCT** eliminates duplicate rows from SELECT output:
```sql
SELECT DISTINCT department FROM employees;
```

**GROUP BY** groups rows and allows aggregate functions per group:
```sql
SELECT department, COUNT(*) FROM employees GROUP BY department;
```

If used without aggregates, `GROUP BY department` and `SELECT DISTINCT department` produce the same output (conceptually). But GROUP BY is meant for aggregation; DISTINCT is meant for deduplication."

#### Indepth
Internally, MySQL may implement both using the same mechanism — a temporary table with a UNIQUE index for deduplication, or a filesort. However, GROUP BY can leverage an index on the grouping column to avoid a sort/temp table. DISTINCT on indexed columns may also avoid a sort. Check EXPLAIN for `Using temporary` and `Using filesort` — both indicate potentially expensive deduplication strategies. For large datasets, GROUP BY over DISTINCT is often more optimizer-friendly.

---

### 115. How do you rename a table?

"MySQL provides two ways to rename a table:

```sql
-- Method 1: RENAME TABLE
RENAME TABLE old_name TO new_name;

-- Method 2: ALTER TABLE
ALTER TABLE old_name RENAME TO new_name;

-- Rename multiple tables atomically:
RENAME TABLE table_a TO old_a, new_a TO table_a;
```

`RENAME TABLE` is **atomic** for multiple tables — this makes it perfect for a cutover pattern: build a new table, then atomically swap old and new with a single `RENAME TABLE` statement."

#### Indepth
The atomic multi-table rename is a critical technique for **zero-downtime schema changes**: (1) Create `orders_new` with updated schema, (2) Populate it, (3) `RENAME TABLE orders TO orders_old, orders_new TO orders` — this swap is atomic, so readers always see either the old or new table, never nothing. Immediately after: `DROP TABLE orders_old`. This is the foundation of tools like `pt-online-schema-change` and `gh-ost`.

---

### 116. How do you add or drop a column from a table?

"Use `ALTER TABLE`:

```sql
-- Add a column:
ALTER TABLE users ADD COLUMN phone VARCHAR(20) AFTER email;

-- Drop a column:
ALTER TABLE users DROP COLUMN phone;

-- Add with default (instant in MySQL 8.0):
ALTER TABLE orders ADD COLUMN status VARCHAR(20) DEFAULT 'pending';
```

In MySQL 8.0, `ADD COLUMN ... DEFAULT` with a constant default can be done **instantly** (no table rebuild) using `ALGORITHM=INSTANT`, regardless of table size."

#### Indepth
Before MySQL 8.0, every `ADD COLUMN` required rebuilding the entire table (O(n) operation) — blocking on large tables for minutes or hours. MySQL 8.0's **Instant ADD COLUMN** (InnoDB Instant DDL) stores the new column's metadata without rebuilding rows. Old rows are treated as having the default value at read time. This allows adding columns to billion-row tables in milliseconds. Check `ALGORITHM=INSTANT` availability with `EXPLAIN FORMAT=JSON` on the ALTER TABLE before running it in production.

---

### 117. What is a default value in a column?

"A **default value** is the value MySQL automatically uses when an INSERT statement doesn't specify a value for that column.

```sql
CREATE TABLE tasks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  created_at DATETIME DEFAULT NOW(),
  status VARCHAR(20) DEFAULT 'pending',
  priority INT DEFAULT 0
);
```

If an INSERT doesn't provide `status`, MySQL inserts `'pending'` automatically.

I always define sensible defaults — especially for timestamp columns and status/flag fields. Explicit defaults make schema intent clear and prevent NULL surprises."

#### Indepth
`DEFAULT NOW()` (and `DEFAULT CURRENT_TIMESTAMP`) creates an **expression default**, available since MySQL 8.0.13 for most data types. Before 8.0.13, only `TIMESTAMP` and `DATETIME` supported dynamic defaults natively. Expression defaults can reference `NOW()`, `UUID()`, or other deterministic functions. `ON UPDATE CURRENT_TIMESTAMP` is a related feature that automatically updates a `DATETIME` column to the current time whenever the row is modified — perfect for `updated_at` tracking columns.

---

### 118. What is the difference between FLOAT and DECIMAL?

"**FLOAT** / **DOUBLE**: Approximate numeric types using IEEE 754 floating-point representation. Compact (4/8 bytes) but prone to **rounding errors**.

**DECIMAL(p, s)**: Exact numeric type. `DECIMAL(10, 2)` stores up to 10 total digits with exactly 2 decimal places. No rounding errors.

```sql
-- Wrong for money:
price FLOAT  -- 9.99 may store as 9.9899999...

-- Correct for money:
price DECIMAL(10, 2)  -- Stores exactly 9.99
```

I always use `DECIMAL` for monetary values, scientific measurements, or anywhere exact precision matters."

#### Indepth
FLOAT's imprecision has caused real financial bugs: summing FLOAT columns gradually drifts from the true value. MySQL stores DECIMAL as a sequence of 9-digit groups in 4 bytes each, preserving exact decimal representation. DECIMAL is slightly slower for arithmetic than FLOAT (software math vs hardware floating-point unit), but for correctness in financial systems, this tradeoff is non-negotiable. Even minor display rounding errors in FLOAT values can cause accounting reconciliation failures.

---

### 119. What is a composite index vs composite key?

"A **composite key** is a logical database concept — a primary key or unique key made of multiple columns to identify rows. It's about data identity and integrity.

A **composite index** is a physical performance structure — a B+Tree index built on multiple columns. It's about query optimization.

Relationship: composite keys are automatically backed by composite indexes (MySQL creates the index to enforce uniqueness/identity). But composite indexes can exist purely for performance without being a key constraint.

Example: `INDEX (last_name, first_name)` is a composite index for optimizing name searches — not a key."

#### Indepth
A composite index `(a, b, c)` is not equivalent to three separate indexes on `(a)`, `(b)`, `(c)`. The composite index enables efficient filtering on `(a)`, `(a,b)`, `(a,b,c)` but not `(b)` or `(c)` alone. Three separate indexes let the optimizer independently filter on any column but require three separate data structures. The optimizer can sometimes combine two separate indexes for a single query via **Index Merge** optimization, but this is generally less efficient than a well-designed composite index.

---

### 120. What is a database schema?

"A **database schema** defines the **structure** of a database: the tables, columns, data types, constraints, indexes, relationships (via foreign keys), views, stored procedures, and triggers.

The schema is the blueprint — it describes what data can be stored and how it's organized, without being the data itself.

In MySQL, a 'schema' and 'database' are the same thing (synonymous keywords). In other databases like PostgreSQL, schemas are sub-namespaces within a database.

I treat schema design as the most important architectural decision — a poorly designed schema causes irreversible performance, scalability, and integrity problems."

#### Indepth
Schema design should be guided by the **access patterns** of the application. An excellent approach: document every query the application will run, then design tables, indexes, and relationships to serve those queries efficiently. This "query-driven schema design" is the opposite of designing a schema in isolation and then hoping the queries work. For evolving schemas, use **migration tools** (Liquibase, Flyway) to version-control schema changes the same way you version application code.

---
