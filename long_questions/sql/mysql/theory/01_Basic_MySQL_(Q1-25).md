# 🟢 Basic MySQL Interview Questions (Q1–25)

---

### 1. What is MySQL?

"MySQL is an open-source **relational database management system (RDBMS)** that stores data in tables and uses **SQL** (Structured Query Language) to query and manipulate that data.

It was originally developed by MySQL AB and is now maintained by Oracle. It is one of the most widely used databases in the world, especially in **LAMP stack** applications (Linux, Apache, MySQL, PHP).

I use it because it's battle-tested, has excellent community support, and integrates seamlessly with virtually every programming language and framework."

#### Indepth
MySQL uses a **client-server model**: the MySQL server manages the data while clients connect to it via TCP/IP or Unix sockets. Internally, MySQL has a pluggable storage engine architecture — this allows it to use InnoDB for transactional workloads or MyISAM for read-heavy, non-transactional scenarios.

---

### 2. What are the features of MySQL?

"MySQL offers several key features: **ACID-compliant transactions** (via InnoDB), **replication**, **full-text search**, **stored procedures**, **triggers**, **views**, and **partitioning**.

It supports multiple storage engines, making it flexible. It also offers robust **security** with user-based access control, SSL connections, and encryption.

What I appreciate most is that these features come with great tooling — MySQL Workbench, `EXPLAIN`, slow query logs — making it easy to develop and tune in production."

#### Indepth
MySQL is distinct from many RDBMSes in its **pluggable storage engine** design. You can even mix engines on a per-table basis within the same database — though this is rarely advisable in production. Since MySQL 5.5, InnoDB became the default, replacing MyISAM, which was not ACID-compliant.

---

### 3. What is a database?

"A **database** is an organized collection of structured data stored electronically. In relational databases like MySQL, data is organized into **tables** which are related to each other via keys.

Think of it like a well-organized filing cabinet. Each drawer is a table, each folder is a row, and each label on the folder is a column.

I rely on databases to provide **persistence** — so data survives application restarts — along with **concurrency** support for multiple users querying it simultaneously."

#### Indepth
In MySQL specifically, a 'database' corresponds to a **directory on disk** (under the data directory). Each table maps to one or more files within that directory. InnoDB tables, however, can optionally be stored in a shared tablespace (`ibdata1`) or in per-table `.ibd` files when `innodb_file_per_table` is ON (the default since MySQL 5.6).

---

### 4. What is a table?

"A **table** is the fundamental storage unit in a relational database. It consists of **rows** (records) and **columns** (fields). Each column has a defined data type and optional constraints.

It's like a spreadsheet, but with strict rules about what data goes in each column and relationships with other tables enforced by the database engine.

I always think carefully about table design upfront — getting the schema right saves enormous pain later when tables grow to millions of rows."

#### Indepth
In InnoDB, every table is physically a **B+Tree clustered index** ordered by the primary key. The table data IS the primary key index. This contrasts with heap-organized tables (like in PostgreSQL) where the data and index are separate structures. This design means primary key choice directly impacts storage layout and I/O patterns.

---

### 5. What are primary keys?

"A **primary key** is a column (or set of columns) that **uniquely identifies** every row in a table. It must be unique and cannot be NULL.

I always define a primary key. Without one, InnoDB internally creates a hidden 6-byte row ID, which wastes space and makes operations opaque.

I typically use `INT AUTO_INCREMENT` for simple tables or a `UUID` for distributed systems where I need globally unique IDs without coordinating with a central sequence."

#### Indepth
In InnoDB, the primary key is the **clustered index** — row data is stored alongside the index key in B+Tree leaf nodes. Secondary indexes store the primary key value as a pointer to the row. This means a large primary key (like `VARCHAR(255)`) makes all secondary indexes heavier. Prefer compact integer primary keys.

---

### 6. What are foreign keys?

"A **foreign key** is a column (or set of columns) that references the **primary key of another table**. It enforces **referential integrity** — you can't orphan a child record that points to a non-existent parent.

For example, an `orders` table's `customer_id` is a foreign key to the `customers` table's `id`.

I use foreign keys to let the database enforce relationships. Without them, I'd need to write application-level checks, which are fragile and easy to bypass."

#### Indepth
MySQL (InnoDB) supports `ON DELETE` and `ON UPDATE` cascade actions: `CASCADE`, `SET NULL`, `RESTRICT`, `NO ACTION`. Careful: `CASCADE DELETE` can trigger a chain of deletes across deeply nested tables, causing unexpected data loss and performance issues. Always test cascade behavior explicitly.

---

### 7. What is a unique key?

"A **unique key** enforces that all values in a column (or combination of columns) are distinct — but unlike primary keys, it **can allow NULL values**.

For example, an `email` column in a `users` table should be unique. Multiple users shouldn't share the same email.

I use unique constraints frequently. They're backed by an index, so they also speed up lookups on those columns."

#### Indepth
MySQL allows **multiple NULL values** in a UNIQUE column, because NULL is not considered equal to NULL in SQL. This is part of the SQL standard. However, some databases (like SQL Server) only allow one NULL per unique column. Be aware of this difference when migrating schemas.

---

### 8. What is a composite key?

"A **composite key** is a primary key or unique key made up of **two or more columns** together. The combination of these columns must be unique across rows.

For example, in a `student_courses` table, neither `student_id` nor `course_id` alone uniquely identifies a row — but the combination does. So `PRIMARY KEY (student_id, course_id)` is a composite key.

I use them for **junction/mapping tables** in many-to-many relationships."

#### Indepth
In InnoDB, a composite primary key determines the physical sort order of rows. The table is a B+Tree clustered on `(col1, col2, ...)`. For secondary index lookups, both columns become part of the index pointer, increasing secondary index size. Keep composite primary keys as short as possible.

---

### 9. What is normalization?

"**Normalization** is the process of organizing a database to reduce **data redundancy** and improve **data integrity**. It involves dividing large tables into smaller, related tables and defining relationships between them.

I normalize schemas to avoid update anomalies — for example, if a customer's name is stored in 10 order rows, updating their name requires 10 updates. In a normalized schema, only one row in the `customers` table needs updating.

It follows a series of rules called **normal forms** (1NF, 2NF, 3NF, BCNF, etc.)."

#### Indepth
Normalization comes with a **JOIN cost tradeoff**. Highly normalized schemas require more JOINs at query time, which adds overhead. In analytics / OLAP workloads, **denormalization** is often preferred deliberately to reduce JOIN counts and enable columnar scans. The right level of normalization depends on the read/write ratio of the workload.

---

### 10. What are the different normal forms?

"The main normal forms are:
- **1NF**: Each cell holds an atomic (single) value; no repeating groups.
- **2NF**: Must be 1NF; every non-key column depends on the *entire* primary key (no partial dependency).
- **3NF**: Must be 2NF; no transitive dependencies (non-key columns don't depend on other non-key columns).
- **BCNF**: A stricter version of 3NF for certain edge cases.
- **4NF / 5NF**: Deal with multi-valued and join dependencies (rare in practice).

I aim for 3NF in transactional (OLTP) systems."

#### Indepth
**BCNF** (Boyce-Codd Normal Form) handles an edge case that 3NF misses: when a non-key determinant determines part of a candidate key. In practice, most well-designed schemas naturally reach BCNF. 4NF and 5NF become relevant in academic settings or highly structured data models but are rarely a concern in production web applications.

---

### 11. What is denormalization?

"**Denormalization** is the intentional process of introducing **redundancy** into a database schema to improve **read performance**. We merge tables or add duplicate columns to reduce expensive JOINs.

For example, storing `product_name` directly in the `order_items` table instead of always joining with the `products` table is denormalization.

I use it for read-heavy systems — like analytics dashboards or reporting tables — where JOIN overhead is unacceptable and data consistency can be managed at the application level."

#### Indepth
Denormalization is a classic **space-vs-time tradeoff**. You trade storage space and update complexity for faster reads. Modern approaches like **materialized views**, **CQRS**, and **read replicas** achieve denormalization's performance benefits while maintaining a normalized source of truth in the write path.

---

### 12. What is the difference between DELETE, TRUNCATE, and DROP?

"These three commands remove data at different levels:
- **DELETE**: Removes specific rows using a `WHERE` clause. It's a DML operation, so it's logged row-by-row and can be rolled back inside a transaction.
- **TRUNCATE**: Removes **all** rows from a table but keeps the table structure intact. It's faster than DELETE because it deallocates data pages rather than logging each row. It cannot be rolled back in MySQL.
- **DROP**: Removes the entire table (structure + data + indexes). Irreversible.

I use DELETE when I need to remove specific rows, TRUNCATE to reset a table during testing, and DROP to remove a table entirely."

#### Indepth
In MySQL, `TRUNCATE` is technically a DDL statement (not DML). It implicitly commits any open transaction before executing. Unlike `DELETE`, `TRUNCATE` also **resets the AUTO_INCREMENT counter** back to 1. `DELETE` does not reset it. This difference matters significantly when migrating or testing with seeded data.

---

### 13. What is the difference between WHERE and HAVING?

"Both filter rows, but at different stages:
- **WHERE** filters rows **before** grouping (before `GROUP BY`). It works on individual rows.
- **HAVING** filters groups **after** grouping. It works on aggregate results.

Example: `WHERE salary > 50000` filters individual employees. `HAVING COUNT(*) > 5` after `GROUP BY department` filters departments with more than 5 employees.

I always remind myself: if there's no `GROUP BY`, `HAVING` behaves like `WHERE`, but using `WHERE` is the correct and efficient choice."

#### Indepth
`WHERE` is processed before the `GROUP BY` and can leverage indexes efficiently. `HAVING` is evaluated after grouping, so it cannot use indexes for the filter itself. For performance, always push as many conditions as possible into `WHERE` before grouping. Only use `HAVING` for conditions that genuinely depend on aggregate results.

---

### 14. What is the default port of MySQL?

"MySQL's default port is **3306**. This is the TCP port the MySQL server listens on for client connections.

When connecting via command line: `mysql -h localhost -P 3306 -u root -p`.

In production, I sometimes change this to a non-standard port as a basic security hardening step — it stops opportunistic port scanners. But it's security by obscurity and is never a substitute for proper firewall rules and strong authentication."

#### Indepth
MySQL also supports connections via **Unix domain sockets** (e.g., `/var/lib/mysql/mysql.sock`) when the client is on the same machine. Unix socket connections are slightly faster than TCP loopback connections because they bypass the network stack entirely. Most ORMs and drivers default to socket connections when the host is `localhost`.

---

### 15. What are MySQL data types?

"MySQL data types fall into four categories:
- **Numeric**: `INT`, `BIGINT`, `SMALLINT`, `TINYINT`, `DECIMAL`, `FLOAT`, `DOUBLE`
- **String**: `CHAR`, `VARCHAR`, `TEXT`, `BLOB`, `ENUM`, `SET`
- **Date/Time**: `DATE`, `DATETIME`, `TIMESTAMP`, `TIME`, `YEAR`
- **Spatial/JSON**: `JSON`, `POINT`, `GEOMETRY`

I choose data types carefully — using the **narrowest type that fits** the data to minimize storage and maximize index efficiency."

#### Indepth
Choosing the wrong data type has performance consequences. For example, storing monetary values as `FLOAT` introduces rounding errors; always use `DECIMAL(10,2)` for money. `TIMESTAMP` is stored as UTC and auto-converts to the session timezone, while `DATETIME` stores exactly what you insert — choose based on whether timezone-awareness matters.

---

### 16. What is AUTO_INCREMENT?

"**AUTO_INCREMENT** is a column attribute that automatically generates a unique, incrementing integer value for each new row. It's typically used with primary keys to avoid manually managing IDs.

```sql
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100)
);
```

Every time a row is inserted without specifying `id`, MySQL assigns the next available integer.

I use it for most transactional tables. For distributed systems, I use UUIDs or application-generated IDs instead to avoid hotspots on a single auto-increment sequence."

#### Indepth
AUTO_INCREMENT counters are **stored in memory** and rebuilt on restart by scanning the maximum existing value (in older MySQL versions). Since MySQL 8.0, the counter is persisted in the redo log, fixing the classic bug where the counter reset after restart if the max-ID row was deleted. Always use `BIGINT` for AUTO_INCREMENT on tables expected to grow large.

---

### 17. What is a NULL value?

"In SQL, **NULL** represents the absence of a value — it's not zero, not an empty string, it literally means 'unknown' or 'not applicable'.

This has important implications: `NULL = NULL` is `FALSE` in SQL. To check for NULL, you must use `IS NULL` or `IS NOT NULL`.

I'm careful with NULL in aggregate functions — most aggregates like `SUM()`, `AVG()`, and `COUNT(column)` **ignore NULLs**. Only `COUNT(*)` counts all rows including those with NULLs."

#### Indepth
NULL values add complexity to **index design**. In MySQL, NULLs ARE stored and indexed in B+Tree indexes (unlike some older databases). However, filtering `WHERE col IS NULL` can still be inefficient if the index statistics aren't accurate. Consider using a sentinel value (like `0` or `'N/A'`) when NULL semantics aren't needed, to keep queries predictable.

---

### 18. What is the difference between CHAR and VARCHAR?

"Both store string data, but differ in storage behavior:
- **CHAR(n)**: Fixed-length. Always stores exactly `n` characters, padding with spaces if shorter. Faster for fixed-length data like country codes, status flags.
- **VARCHAR(n)**: Variable-length. Stores only as many bytes as the actual string plus 1–2 length bytes. More storage-efficient for varied-length strings.

I use `CHAR` for columns where every value is the same length (e.g., `CHAR(2)` for country codes). For everything else — names, emails, descriptions — I use `VARCHAR`."

#### Indepth
In InnoDB, when a row in a CHAR column is updated with a shorter value, the row slot retains its original padded size — there's no shrinkage. `VARCHAR` can cause row overflow and off-page storage for very long values. In the `MEMORY` storage engine, `VARCHAR` is treated as `CHAR` internally for performance, so there's no storage saving for in-memory temporary tables.

---

### 19. What is a constraint?

"A **constraint** is a rule enforced at the database level to ensure data integrity. MySQL supports:
- **NOT NULL**: Column cannot store NULL.
- **UNIQUE**: All values in the column must be distinct.
- **PRIMARY KEY**: Unique + NOT NULL.
- **FOREIGN KEY**: References another table's primary/unique key.
- **CHECK**: Validates a condition (supported since MySQL 8.0.16).
- **DEFAULT**: Sets a default value if none is provided.

I prefer defining constraints in the schema rather than only in application code — the database is the last line of defense against bad data."

#### Indepth
MySQL introduced **CHECK constraint enforcement** in 8.0.16. Before that, CHECK constraints were parsed but silently ignored. This is a critical difference from PostgreSQL, which has always enforced CHECK. If you're migrating schemas written for MySQL < 8.0.16, assume CHECK constraints may have been meaningless and verify data integrity separately.

---

### 20. What is the difference between INNER JOIN and OUTER JOIN?

"**INNER JOIN** returns only rows where there is a matching record in **both** tables. Non-matching rows are excluded entirely.

**OUTER JOIN** returns all rows from one or both tables, filling in NULLs where there is no match:
- **LEFT OUTER JOIN**: All rows from the left table, NULLs for unmatched right rows.
- **RIGHT OUTER JOIN**: All rows from the right table, NULLs for unmatched left rows.
- **FULL OUTER JOIN**: All rows from both tables (MySQL doesn't natively support this — simulate with UNION of LEFT and RIGHT joins).

I use INNER JOIN by default and switch to LEFT JOIN when I need to preserve records that may not have a matching row."

#### Indepth
The MySQL query optimizer can rewrite a `LEFT JOIN` as an `INNER JOIN` when it can prove that NULL rows would be eliminated by the `WHERE` clause anyway (e.g., `WHERE right_table.id IS NOT NULL`). Understanding this helps explain why EXPLAIN output sometimes shows an INNER join for a query you wrote as an OUTER join.

---

### 21. What are aggregate functions in MySQL?

"**Aggregate functions** compute a single result from a set of rows:
- `COUNT()` — counts rows
- `SUM()` — totals a numeric column
- `AVG()` — computes the average
- `MAX()` / `MIN()` — largest / smallest value
- `GROUP_CONCAT()` — concatenates values from multiple rows into a string

They are typically used with `GROUP BY` to produce summary statistics per group.

I use them constantly in reporting queries — for example, `SELECT department, AVG(salary) FROM employees GROUP BY department`."

#### Indepth
All standard aggregate functions **ignore NULLs** except `COUNT(*)`. `GROUP_CONCAT()` has a configurable max length (`group_concat_max_len`, default 1024 bytes). Always set it explicitly in sessions doing large aggregations. MySQL 8.0 added window functions as a powerful alternative to aggregates for running totals, rankings, and moving averages without collapsing rows.

---

### 22. What is the GROUP BY clause?

"**GROUP BY** groups rows that share the same value in one or more columns and allows aggregate functions to be applied to each group independently.

```sql
SELECT department, COUNT(*), AVG(salary)
FROM employees
GROUP BY department;
```

This returns one row per department with the count and average salary.

I always pair GROUP BY with meaningful aggregates. Selecting non-aggregated columns not in GROUP BY in strict SQL mode raises an error (MySQL 5.7+ with `ONLY_FULL_GROUP_BY`)."

#### Indepth
In MySQL, `ONLY_FULL_GROUP_BY` mode (enabled by default in 5.7+) enforces SQL standard compliance: you cannot SELECT a non-aggregated column that's not in GROUP BY. Before 5.7, MySQL allowed this and returned an arbitrary value from the group — a common source of subtle bugs. Always test GROUP BY queries with `ONLY_FULL_GROUP_BY` enabled.

---

### 23. What is the ORDER BY clause?

"**ORDER BY** sorts the result set by one or more columns, either ascending (`ASC`, the default) or descending (`DESC`).

```sql
SELECT name, salary FROM employees ORDER BY salary DESC;
```

I use ORDER BY freely in ad-hoc queries. In application code, however, I only add ORDER BY when the sort order genuinely matters — unnecessary sorting on large result sets adds CPU and memory cost."

#### Indepth
MySQL can satisfy ORDER BY without a sort operation (filesort) if an appropriate **index** exists on the ORDER BY columns. This is called an **index scan in order**. The EXPLAIN output will show `Using filesort` when MySQL cannot use an index for sorting. For performance-critical queries, create indexes that cover both the WHERE and ORDER BY columns in the correct order.

---

### 24. What is the LIMIT clause?

"**LIMIT** restricts the number of rows returned by a query. Optionally, `OFFSET` skips a specified number of rows first.

```sql
SELECT * FROM products ORDER BY created_at DESC LIMIT 10 OFFSET 20;
```

This returns rows 21–30, which is the basis for **pagination**.

I use LIMIT in every user-facing paginated query. Returning unbounded result sets to an application is a common cause of out-of-memory errors and slow responses."

#### Indepth
`LIMIT` + `OFFSET` for deep pagination is a hidden performance trap. On a table with 10M rows, `LIMIT 10 OFFSET 9999990` requires the database to scan and discard ~10M rows to find the last 10. The correct pattern for deep pagination is **keyset pagination** (cursor-based): `WHERE id > last_seen_id ORDER BY id LIMIT 10`, which navigates via the index directly.

---

### 25. How do you create a database in MySQL?

"Creating a database is straightforward:

```sql
CREATE DATABASE myapp_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

I always specify `utf8mb4` (not just `utf8`!) and a collation. `utf8` in MySQL is a misleading 3-byte variant that can't store emoji or certain Unicode characters. `utf8mb4` is the true 4-byte UTF-8.

Then I select it with `USE myapp_db;` before running table creation statements."

#### Indepth
MySQL's `utf8` charset is a legacy implementation that only supports up to 3-byte Unicode code points (BMP plane), excluding emoji (which are 4-byte). `utf8mb4` was introduced to fix this and is the correct choice for any modern application. Always set `utf8mb4` at the server, database, table, and connection level for consistency — mismatched collations between client and server can cause character encoding corruption.

---
