# Database & SQL Interview Questions (81-90)

## SQL Fundamentals & Transaction Management

### 81. Difference between INNER JOIN and LEFT JOIN?
"**INNER JOIN** returns rows only when there is a match in *both* tables. If a User has no Orders, neither the User nor the Order appears.

**LEFT JOIN** (or LEFT OUTER JOIN) returns *all* rows from the left table, and the matched rows from the right table. If there’s no match, the right side will contain NULLs.

So if I want a list of 'All Customers and their Orders (if any)', I use LEFT JOIN. If I only want 'Customers who have actually placed an order', I use INNER JOIN."

**Spoken Format:**
"Think of these as two different ways to combine information from multiple tables.

**INNER JOIN** is like having a strict club where only members who belong to BOTH clubs can enter. If you want a list of 'people who are both members AND have attended meetings', you get only those who qualify for both conditions.

**LEFT JOIN** is like having an open house party where you invite everyone from your neighborhood, even if they're not members of your club. You get all your neighbors, and for those who are also club members, you see their meeting attendance.

The key difference: INNER JOIN is more restrictive (needs match in both tables), LEFT JOIN is more inclusive (gets all from one table regardless of match in other).

Choose based on what you need: if you need only matching records, use INNER JOIN. If you need all records from one table regardless of matches in another, use LEFT JOIN."

### 82. What is indexing and how does it work?
"Indexing is like the index at the back of a book. Instead of flipping through every single page (scanning every row) to find 'John Doe', the database looks up 'John Doe' in the index, which tells it exactly which page (data block) to go to.

Technically, it creates a separate data structure (usually a **B-Tree**) that keeps the indexed column values sorted. This allows the database to use binary search (O(log n)) instead of linear search (O(n))."

**Spoken Format:**
"Indexing is like having a super-smart librarian who knows exactly where every book is located.

Without an index, if you want to find a book by author, the librarian has to check every single book in the library - that's O(n) time, very slow for large libraries.

With an index, it's like the librarian has created a separate card catalog sorted by author. When you ask for books by author, they go directly to that catalog and find the exact location instantly - that's O(log n) time, much faster.

The tradeoff is that indexes make reads much faster, but they slow down writes (inserts/updates) because every time you add a new book, the librarian has to update the author catalog too.

Indexes are perfect for columns that are frequently searched but rarely change - like last_name, email, or created_date."

### 83. When can indexes hurt performance?
"Indexes speed up reads, but they **slow down writes** (INSERT, UPDATE, DELETE).

Every time you write a new row, the database has to update not just the table but also every single index associated with that table.

Also, indexes take up disk space and RAM. If you index a low-cardinality column like 'Gender' (only 2-3 values), the index is huge and useless because DB will likely just do a full table scan anyway."

**Spoken Format:**
"Indexes are like creating shortcuts in your database, but they come with a cost.

Think about indexing a column like 'Gender' that only has 'Male', 'Female', 'Other'. Creating an index here is like creating a massive shortcut that points to only 3 possible locations.

The problem is that most of the time, the database will ignore your shortcut and just scan the entire table anyway - because it's faster to check everything than to use your tiny shortcut.

But if you index a high-cardinality column like 'Email' where every user has a unique value, the index is very effective. It's like creating a shortcut that leads directly to each person's unique email address.

The rule is: index columns that are frequently searched AND have many different values (high cardinality). Avoid indexing columns with few values like 'Gender' or boolean flags - the index overhead outweighs the benefit.

It's like choosing which roads to build shortcuts on - you build highways for major routes, not tiny paths to every house!"

### 84. What is normalization?
"Normalization is the process of organizing data to minimize redundancy and dependency.

**1NF**: Atomic values (no comma-separated lists in a column).
**2NF**: No partial dependencies (everything depends on the whole primary key).
**3NF**: No transitive dependencies (non-key columns shouldn't depend on other non-key columns).

In practice, we usually aim for 3NF to avoid data anomalies. However, in high-scale read systems, we often **denormalize** (duplicate data) to avoid expensive joins."

**Spoken Format:**
"Normalization is like organizing your warehouse to eliminate redundancy and improve efficiency.

**1NF (First Normal Form)** is like giving every product its own unique shelf space. No more products sharing the same space, no confusion.

**2NF (Second Normal Form)** is like ensuring that each shelf only contains products from one category. You can't have a shelf that has both electronics and groceries mixed together.

**3NF (Third Normal Form)** is like making sure that every product location is recorded only once. If a product is stored in multiple warehouses, you don't repeat the full product details everywhere.

**Denormalization** is like intentionally duplicating some information to avoid expensive operations. Imagine you often need to show product details along with user information. Instead of joining massive tables every time, you store some duplicate information to make reads faster.

The tradeoff: Normalized data is organized and avoids anomalies, but denormalized data is faster to read. Choose based on your specific needs!"

### 85. What is a transaction?
"A transaction is a logical unit of work that contains one or more SQL statements.

Wait, it's more than that—it treats multiple operations as a single atomic 'all-or-nothing' action. If I transfer money from Account A to Account B, two updates happen: debit A, credit B.

A transaction guarantees that if the credit to B fails, the debit to A is rolled back. You never lose money."

**Spoken Format:**
"Transactions are like safety deposit boxes for your database operations - they ensure that everything happens perfectly or nothing happens at all.

**Atomicity** is the 'all or nothing' rule. Imagine transferring money between bank accounts. Either the complete transfer succeeds, or it fails completely. You never end up with money disappearing from one account but not appearing in another.

**Consistency** means the database always moves from one valid state to another. If you transfer $100 from A to B, the database will never show a state where A has $100 and B has $100 less.

**Isolation** means transactions don't interfere with each other. If two transfers happen at the same time, they don't mix up each other's data.

**Durability** means committed changes are permanent. Once the transfer is confirmed, it survives power outages and restarts.

Transactions are the database's promise that your data will always be in a valid state!"

### 86. Explain ACID properties.
"**A - Atomicity**: All or nothing. If one part fails, the entire transaction fails.

**C - Consistency**: The database moves from one valid state to another. Constraints (Foreign Keys, Unique) are always enforced.

**I - Isolation**: Transactions shouldn't interfere with each other. If I’m reading data, I shouldn’t see half-written data from another transaction.

**D - Durability**: Once committed, the data is saved permanently, even if the power goes out immediately after."

**Spoken Format:**
"ACID properties are like the four pillars of a strong database foundation.

**Atomicity** is the 'all or nothing' rule. Imagine transferring money between bank accounts. Either the complete transfer succeeds, or it fails completely. You never end up with money disappearing from one account but not appearing in another.

**Consistency** means the database always moves from one valid state to another. If you transfer $100 from A to B, the database will never show a state where A has $100 and B has $100 less.

**Isolation** means transactions don't interfere with each other. If two transfers happen at the same time, they don't mix up each other's data.

**Durability** means committed changes are permanent. Once the transfer is confirmed, it survives power outages and restarts.

ACID properties ensure that your database is always in a valid state, even in the face of failures or concurrent transactions."

### 87. Isolation levels and problems they solve?
"This controls 'I' in ACID.

1.  **Read Uncommitted**: Lowest level. You can read uncommitted data (Dirty Read). Dangerous, rarely used.
2.  **Read Committed**: Default in Postgres/Oracle. You only read committed data. Prevents Dirty Reads.
3.  **Repeatable Read**: Default in MySQL. If I query a row twice in a transaction, I get the same result. Prevents **Non-Repeatable Reads**. But **Phantom Reads** (new rows appearing) are still possible.
4.  **Serializable**: Highest level. Transactions run sequentially. Prevents everything but kills concurrency."

**Spoken Format:**
"Isolation levels are like the different levels of protection for your database transactions.

**Read Uncommitted** is like having no protection at all. You can read data that's not even committed yet - it's like reading a draft document that's not finalized.

**Read Committed** is like having a basic lock on your data. You only read data that's been committed, but you might still see changes made by other transactions.

**Repeatable Read** is like having a stronger lock on your data. You get the same result if you query a row twice in a transaction, but you might still see new rows appearing.

**Serializable** is like having the strongest lock on your data. Transactions run sequentially, so you don't see any changes made by other transactions. But this comes at the cost of concurrency - it's like having a single-lane road where only one car can pass at a time.

Choose the right isolation level based on your specific needs - do you need to prevent dirty reads, or can you tolerate some concurrency?"

### 88. Difference between `WHERE` and `HAVING`?
"**WHERE** filters rows *before* grouping or aggregation. It works on individual records.

**HAVING** filters groups *after* the aggregation has happened.

So: `SELECT city, COUNT(*) FROM users WHERE age > 18 GROUP BY city HAVING COUNT(*) > 1000`.

First, filter adults (WHERE). Then group by city. Then filter cities with more than 1000 adults (HAVING)."

**Spoken Format:**
"Think of these as two different stages of processing a large dataset.

**WHERE** is like the initial filtering - it's like sorting through all people and keeping only those who meet basic criteria. 'Find all people over 18'.

**GROUP BY** is like creating separate piles for each category. After filtering adults, you create piles for each city.

**HAVING** is like the quality control step - you look at each pile and only keep those that meet your standard. 'Only keep city piles that have more than 1000 adults'.

The magic is the order: SQL processes WHERE first, then GROUP BY, then HAVING. You can't use HAVING before GROUP BY because it filters the groups, not the individual rows.

It's like processing mail: first sort by state, then group by city, then check which groups meet criteria. Each step builds on the previous one!"

### 89. How do you optimize a slow query?
"First, I run `EXPLAIN` (or `EXPLAIN ANALYZE`). This tells me if the query is using an **Index** or doing a **Full Table Scan**.

If it's a scan, I check if adding an index on the WHERE/JOIN columns helps.

Then I look at simple things:
-   Am I selecting `*` when I only need 2 columns?
-   Is there a `LIKE '%value%'` leading wildcard preventing index usage?
-   Are we doing calculations on the column side (e.g., `WHERE YEAR(date) = 2023`)? That kills indexes too.

If it's still slow, I might need to rewrite the query (CTEs vs Subqueries) or denormalize data."

**Spoken Format:**
"Query optimization is like being a detective for finding performance problems in your database.

The first tool is **EXPLAIN** - it's like getting a detailed map of how the database plans to execute your query. It shows you if it's using an index or doing a full table scan.

Common problems I look for:

1. **Missing indexes** - the query is filtering on a column that's not indexed, causing full table scans
2. **Wildcard misuse** - using LIKE '%value' at the beginning prevents index usage
3. **Functions on columns** - using functions like `YEAR(date)` instead of direct comparisons prevents indexes
4. **Expensive operations in WHERE** - calculations on columns in the WHERE clause

If indexes don't help, I consider:
- **Query rewriting** - restructuring the query logic
- **Denormalization** - duplicating some data to avoid complex joins
- **Materialized views** - pre-computing complex results

The goal is to make the database do less work for each query. Sometimes the best optimization is changing the question you're asking!"

### 90. What is an execution plan?
"It’s the roadmap for the database engine.

When you send a SQL query, the Optimizer decides the best way to execute it: which index to use, the order of joining tables, whether to use a nested loop join or a hash join.

The `EXPLAIN` command shows you this plan. Reading it is a skill—you look for things like 'Sequential Scan' (bad for big tables) or 'Index Scan' (good) and 'Cost' estimates."
