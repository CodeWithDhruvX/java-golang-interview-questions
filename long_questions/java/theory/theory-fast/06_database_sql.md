# Database & SQL Interview Questions (81-90)

## SQL Fundamentals & Transaction Management

### 81. Difference between INNER JOIN and LEFT JOIN?
"**INNER JOIN** returns rows only when there is a match in *both* tables. If a User has no Orders, neither the User nor the Order appears.

**LEFT JOIN** (or LEFT OUTER JOIN) returns *all* rows from the left table, and the matched rows from the right table. If there’s no match, the right side will contain NULLs.

So if I want a list of 'All Customers and their Orders (if any)', I use LEFT JOIN. If I only want 'Customers who have actually placed an order', I use INNER JOIN."

### 82. What is indexing and how does it work?
"Indexing is like the index at the back of a book. Instead of flipping through every single page (scanning every row) to find 'John Doe', the database looks up 'John Doe' in the index, which tells it exactly which page (data block) to go to.

Technically, it creates a separate data structure (usually a **B-Tree**) that keeps the indexed column values sorted. This allows the database to use binary search (O(log n)) instead of linear search (O(n))."

### 83. When can indexes hurt performance?
"Indexes speed up reads, but they **slow down writes** (INSERT, UPDATE, DELETE).

Every time you write a new row, the database has to update not just the table but also every single index associated with that table.

Also, indexes take up disk space and RAM. If you index a low-cardinality column like 'Gender' (only 2-3 values), the index is huge and useless because the DB will likely just do a full table scan anyway."

### 84. What is normalization?
"Normalization is the process of organizing data to minimize redundancy and dependency.

**1NF**: Atomic values (no comma-separated lists in a column).
**2NF**: No partial dependencies (everything depends on the whole primary key).
**3NF**: No transitive dependencies (non-key columns shouldn't depend on other non-key columns).

In practice, we usually aim for 3NF to avoid data anomalies. However, in high-scale read systems, we often **denormalize** (duplicate data) to avoid expensive joins."

### 85. What is a transaction?
"A transaction is a logical unit of work that contains one or more SQL statements.

Wait, it's more than that—it treats multiple operations as a single atomic 'all-or-nothing' action. If I transfer money from Account A to Account B, two updates happen: debit A, credit B.

A transaction guarantees that if the credit to B fails, the debit to A is rolled back. You never lose money."

### 86. Explain ACID properties.
"**A - Atomicity**: All or nothing. If one part fails, the entire transaction fails.

**C - Consistency**: The database moves from one valid state to another. Constraints (Foreign Keys, Unique) are always enforced.

**I - Isolation**: Transactions shouldn't interfere with each other. If I’m reading data, I shouldn't see half-written data from another transaction.

**D - Durability**: Once committed, the data is saved permanently, even if the power goes out immediately after."

### 87. Isolation levels and problems they solve?
"This controls 'I' in ACID.

1.  **Read Uncommitted**: Lowest level. You can read uncommitted data (Dirty Read). Dangerous, rarely used.
2.  **Read Committed**: Default in Postgres/Oracle. You only read committed data. Prevents Dirty Reads.
3.  **Repeatable Read**: Default in MySQL. If I query a row twice in a transaction, I get the same result. Prevents **Non-Repeatable Reads**. But **Phantom Reads** (new rows appearing) are still possible.
4.  **Serializable**: Highest level. Transactions run sequentially. Prevents everything but kills concurrency."

### 88. Difference between `WHERE` and `HAVING`?
"**WHERE** filters rows *before* grouping or aggregation. It works on individual records.

**HAVING** filters groups *after* the aggregation has happened.

So: `SELECT city, COUNT(*) FROM users WHERE age > 18 GROUP BY city HAVING COUNT(*) > 1000`.

First, filter adults (WHERE). Then group by city. Then filter cities with more than 1000 adults (HAVING)."

### 89. How do you optimize a slow query?
"First, I run `EXPLAIN` (or `EXPLAIN ANALYZE`). This tells me if the query is using an **Index** or doing a **Full Table Scan**.

If it's a scan, I check if adding an index on the WHERE/JOIN columns helps.

Then I look at simple things:
-   Am I selecting `*` when I only need 2 columns?
-   Is there a `LIKE '%value%'` leading wildcard preventing index usage?
-   Are we doing calculations on the column side (e.g., `WHERE YEAR(date) = 2023`)? That kills indexes too.

If it's still slow, I might need to rewrite the query (CTEs vs Subqueries) or denormalize data."

### 90. What is an execution plan?
"It’s the roadmap for the database engine.

When you send a SQL query, the Optimizer decides the best way to execute it: which index to use, the order of joining tables, whether to use a nested loop join or a hash join.

The `EXPLAIN` command shows you this plan. Reading it is a skill—you look for things like 'Sequential Scan' (bad for big tables) or 'Index Scan' (good) and 'Cost' estimates."
