# SQL & Database - Interview Answers

> 🎯 **Focus:** Data is everything. Focus on integrity (ACID) and performance (indexing).

### 1. ACID Properties?
"It represents the guarantee of a Transaction.
**Atomicity**: All or nothing. If one part fails, the whole transaction rolls back.
**Consistency**: The database moves from one valid state to another (constraints are respected).
**Isolation**: Transactions don't interfere with each other (concurrency control).
**Durability**: Once committed, it stays committed even if the power goes out (saved to disk)."

---

### 2. Normalization (1NF, 2NF, 3NF)?
"It’s the process of organizing data to reduce redundancy.
**1NF**: Atomic values. No lists in a single cell.
**2NF**: No partial dependencies. All non-key columns depend on the *whole* primary key.
**3NF**: No transitive dependencies. Columns shouldn't depend on other non-key columns.

In practice, I usually aim for 3NF to keep data clean, but sometimes I 'Denormalize' (intentionally add redundancy) for read performance in reporting tables."

---

### 3. Inner vs Outer Joins?
"**Inner Join**: Returns only rows where there is a match in *both* tables. The intersection.
**Left (Outer) Join**: Returns everything from the Left table, and matches from the Right (or NULL if no match).
**Right Join**: The opposite.
**Full Join**: Everything from both tables (Union).

I use Left Joins the most—like 'Get all Users and their Orders, even if they have zero orders'."

---

### 4. Clustered vs Non-Clustered Index?
"An index makes searching fast.
**Clustered Index**: Determines the physical order of data on the disk. A table can have only *one* (usually the Primary Key). It's like the page numbers in a book.
**Non-Clustered Index**: A separate structure that points to the data. You can have many. It's like the index at the back of the book.

Indexes speed up reads but slow down writes (inserts/updates), so I only add them on columns we actually filter or sort by."

---

### 5. `HAVING` vs `WHERE`?
"Both filter data, but at different times.
**WHERE** filters rows *before* grouping.
**HAVING** filters groups *after* aggregation.

Example: `WHERE salary > 5000` (Individual filter).
`GROUP BY dept HAVING AVG(salary) > 5000` (Group filter)."

---

### 6. Primary Key vs Unique Key?
"Both ensure uniqueness.
**Primary Key**: Uniquely identifies the row. Cannot be NULL. Only one per table.
**Unique Key**: Ensures values are unique. Can usually accept one NULL value. You can have multiple Unique keys per table.

I use PK for ID and Unique Key for logic like 'Email Address'."

---

### 7. Stored Procedure vs View?
"**View**: A virtual table based on a query. It doesn't store data (unless it's a Materialized View). I use it to simplify complex joins for read-only access.
**Stored Procedure**: A script stored in the DB. It can have logic, loops, loops, and parameters. I try to avoid them in modern apps because business logic should live in the Java Service layer, not hidden in the database."

---

### 8. `UNION` vs `UNION ALL`?
"Both combine results from two queries.
**UNION**: Removes duplicates. It acts like a Set. Slower because it has to sort/distinct.
**UNION ALL**: Keeps duplicates. It just appends the lists. Faster.

I default to `UNION ALL` unless I specifically need to filter duplicates."

---

### 9. What is Database Sharding?
"It’s a horizontal scaling strategy.
Instead of one giant database, you split data across multiple databases (shards) based on a key (e.g., User ID).
Users 1-1000 go to DB Server A. Users 1001-2000 go to DB Server B.
It allows theoretically infinite writes, but it makes queries strictly harder (joining across shards is a nightmare)."

---

### 10. `DELETE` vs `TRUNCATE`?
"**DELETE**: A DML command. Deletes row by row. Can be rolled back. Triggers fire.
**TRUNCATE**: A DDL command. Drops the whole table content and recreates it. Instant. Cannot be rolled back (usually). No triggers.

If I need to clear a test table cleanly, I use TRUNCATE."
