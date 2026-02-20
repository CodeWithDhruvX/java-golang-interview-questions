# 📘 SQL Theory Interview Questions

## 🟢 Basics & Database Concepts

### 1. What is the difference between DBMS and RDBMS?
"A **DBMS** is just software to store and manage data, often as files. It’s fine for single-user applications, but I wouldn't use it for a serious backend.
 
**RDBMS** (Relational DBMS) is what I use for almost every production system. It stores data in **tables** with strict relationships. It supports multiple users and enforces integrity constraints like Foreign Keys. If I'm building a web app, I'm choosing Postgres or MySQL (RDBMS), not a flat-file manager."

#### Indepth
RDBMS is based on E.F. Codd's relational model. It guarantees ACID properties, which are critical for transactional systems. A standard DBMS (like a simple XML store or early hierarchical databases) might not provide these guarantees, leading to potential data corruption in concurrent environments.

---

### 2. Explain ACID properties in a database.
"ACID is the safety net for my data. It guarantees that my transactions are processed reliably.
 
*   **Atomicity:** It’s 'all or nothing'. If I transfer money and the system crashes halfway, the money doesn't disappear; the transaction rolls back.
*   **Consistency:** The database moves from one valid state to another. Constraints are never violated.
*   **Isolation:** Use A doesn't see User B's half-finished work.
*   **Durability:** Once I get a 'success' response, that data is saved to disk forever, even if the power goes out."

#### Indepth
Implementing ACID is costly. **NoSQL** databases often sacrifice ACID (specifically Consistency and Atomicity via CAP theorem) for horizontal scalability. However, modern RDBMSs use Write-Ahead Logging (WAL) to ensure Durability and various locking mechanisms (MVCC) to handle Isolation efficiently.

---

### 3. What is Normalization?
"I use normalization to organize my database to reduce **redundancy**. I don't want to update a customer's address in ten different places; I want to update it in one `Customers` table and have everyone else reference it.
 
It prevents data anomalies. If I delete an order, I shouldn't accidentally delete the customer's info just because it was stored in the same row."

#### Indepth
*   **1NF:** Atomic values (no lists in cells), unique rows.
*   **2NF:** 1NF + Partial dependencies removed (all non-key columns depend on the *entire* Primary Key).
*   **3NF:** 2NF + Transitive dependencies removed (non-key columns depend *only* on the Primary Key).
Most production schemas aim for **3NF**. Going beyond (BCNF, 4NF) often yields diminishing returns and overly complex queries.

---

### 4. What is Denormalization?
"Denormalization is something I do *intentionally* for performance, usually in read-heavy systems or Data Warehouses.
 
If I have to join 10 tables just to show a user profile, it's too slow. So, I might redundantly store the `CityName` directly in the `Users` table instead of just the `CityID`. I trade storage space and write complexity (updating multiple places) for faster read speeds."

#### Indepth
Denormalization is NOT "bad design" if done with a purpose. In **OLAP** (Online Analytical Processing) systems like Star Schemas, denormalization is standard practice to minimize joins on massive datasets.

---

### 5. What is the difference between `DELETE`, `TRUNCATE`, and `DROP`?
"I use **`DELETE`** when I want to remove specific rows, like 'delete users who haven't logged in for a year'. It’s slower because it logs every removal.
 
I use **`TRUNCATE`** when I want to clear a table completely but keep the structure. It’s a reset button. It’s instant because it just deallocates the data pages instead of logging every row.
 
I use **`DROP`** when I want to destroy the table entirely—schema and all. Ideally, I only do this in migration scripts."

#### Indepth
`TRUNCATE` is a DDL command, not DML. This means it resets the High Water Mark and identity columns (auto-increments) back to the seed value. You generally cannot roll back a `TRUNCATE` in some databases (like MySQL) unless it's wrapped in a transaction (engine dependent), whereas `DELETE` is fully transaction-safe.

---

## 🟡 Schema & Objects

### 6. What is an Index?
"An index is exactly like the index at the back of a book. Without it, the database has to read every single page (Full Table Scan) to find 'John Doe'.
 
With an index, it jumps straight to the 'J' section and finds the storage location. I add indexes to columns that I frequently use in `WHERE`, `JOIN`, or `ORDER BY` clauses to speed up retrieval."

#### Indepth
*   **Clustered Index:** Physically sorts the table data. A table can have only one (usually the Primary Key).
*   **Non-Clustered Index:** A separate structure (B-Tree) that points to the actual data rows. A table can have many.
Indexes hurt **write performance** (INSERT/UPDATE/DELETE) because the database must update the index structure every time data changes.

---

### 7. What are Primary and Foreign Keys?
"A **Primary Key (PK)** is the unique ID for a record, like a Social Security Number. It identifies *who* you are. It cannot be NULL.
 
A **Foreign Key (FK)** is a reference to a Primary Key in another table. It creates a relationship. I use it to say 'This Order belongs to *that* Customer'. It prevents me from creating an Order for a Customer that doesn't exist."

#### Indepth
A Foreign Key enforces **Referential Integrity**. If you try to delete a Customer who still has Orders (linked via FK), the database will block the delete (unless `ON DELETE CASCADE` is set). This prevents "orphaned records" in your database.

---

### 8. What is a View?
"A View is a 'virtual table'. It’s just a saved SQL query. I use it to simplify complex logic for other developers or for reporting tools.
 
Instead of asking a frontend dev to write a 20-line Join query every time, I create a view `ActiveUserOrders` and tell them: 'Just `SELECT *` from this view'. It also adds security; I can hide sensitive columns like `PasswordHash` or `Salary` inside the view."

#### Indepth
Standard Views do not store data; they run the underlying query every time you access them. **Materialized Views**, however, physically store the result of the query on disk. They are much faster for reading but need to be refreshed (recalculated) when the underlying data changes.

---

### 9. Stored Procedures vs Functions.
"I use **Functions** when I need to compute a value that I want to use *inside* a SQL statement. For example, a function to format a date or calculate tax: `SELECT Price, CalculateTax(Price) FROM Orders`.
 
I use **Stored Procedures** for executed logic or batch jobs. They can execute transactions, modify multiple tables, and handle complex business logic. They are 'called' independently, not part of a SELECT statement."

#### Indepth
Functions are generally deterministic and effectively side-effect free (in terms of schema state). Stored Procedures are compiled execution plans that can perform DDL and DML operations. Stored Procedures are also beneficial for security (preventing SQL injection) and network performance (sending one command instead of 10 separate queries).

---

### 10. What is a Trigger?
"A Trigger is code that fires automatically in response to an event, like `INSERT`, `UPDATE`, or `DELETE`.
 
I use them mostly for **auditing**. If someone changes a Salary, I have a trigger that automatically writes the old value, new value, and timestamp to an `AuditLog` table. It ensures that no matter *how* the data was changed (app, manual console, script), the audit trail is always captured."

#### Indepth
Triggers can be `BEFORE` (validating data) or `AFTER` (logging/cascading). Be careful: Triggers are invisible side effects. A simple `UPDATE` might trigger a cascade of invisible operations that slows down the system or causes unexpected locks. Debugging heavy trigger logic is a nightmare.

---

## 🟠 Queries & Advanced Logic

### 11. Explain Join types.
"**INNER JOIN** is my default. It gives me rows that match in *both* tables.
 
**LEFT JOIN** is what I use when I want 'all users, even if they haven't placed an order'. I get NULLs on the right side if there's no match.
 
**FULL JOIN** gives me everything from both sides. It’s rare, but useful for synchronization reports.
 
**CROSS JOIN** combines every row with every other row. I typically only use this for generating test data or dictionaries (e.g., every shirt size x every color)."

#### Indepth
The performance of joins depends heavily on indexing. If you join two large tables on columns that are not indexed, the database performs a "Hash Join" or "Nested Loop Join" which can be incredibly slow. Always index your foreign keys.

---

### 12. What is the difference between `WHERE` and `HAVING`?
"I use `WHERE` to filter valid **rows** *before* they are grouped.
I use `HAVING` to filter the **groups** *after* the aggregation is done.
 
For example:
`WHERE Salary > 50000` (Filter people before grouping).
`GROUP BY Department HAVING COUNT(*) > 10` (Filter departments assuming they have more than 10 people)."

#### Indepth
You strictly *cannot* use aggregate functions like `SUM()` or `COUNT()` in a `WHERE` clause because those values don't exist yet when the `WHERE` filter runs. The query order is: `FROM` -> `WHERE` -> `GROUP BY` -> `HAVING` -> `SELECT`.

---

### 13. What is a CTE (Common Table Expression)?
"A CTE, defined with `WITH`, is like a temporary variable for a query.
 
I use it to make my code readable. Instead of writing a massive nested subquery that makes my eyes bleed, I define a CTE at the top: `WITH BestCustomers AS (...)`, and then query against it below. It reads top-to-bottom like code."

#### Indepth
**Recursive CTEs** are the real superpower. They allow you to traverse hierarchical data (like Org Charts or Menu Trees) within a single query by referencing the CTE inside its own definition.

---

### 14. What are Window Functions?
"Window functions are powerful because they let me calculate values across a group of rows *without* collapsing them into a single row like `GROUP BY` does.
 
I use `ROW_NUMBER()` to deduplicate data or implement pagination. Use `RANK()` for leaderboards. Use `LEAD/LAG` to compare this month's sales to last month's in the same row."

#### Indepth
The syntax `OVER (PARTITION BY ... ORDER BY ...)` is key. `PARTITION BY` defines the bucket (e.g., per Department), and `ORDER BY` defines the sequence within that bucket. This happens *after* the initial `WHERE` filtering but *before* the final `ORDER BY` of the result set.

---

### 15. What is SQL Injection and how do I prevent it?
"SQL Injection is when a user tricks the database into running malicious commands by typing them into an input field.
 
I prevent it by **always** using **Parameterized Queries** (or Prepared Statements). I never, ever concatenate strings to build a query (`'SELECT * FROM Users WHERE Name = ' + input`). Parameterization forces the database to treat the input as literal data, not executable code."

#### Indepth
It works because the database compiles the SQL statement *before* inserting the user input. If the user inputs `'; DROP TABLE Users; --`, the database just looks for a user literally named `'; DROP TABLE Users; --`. It defangs the attack completely.

---

## 🔴 Advanced Mastery & Performance

### 16. Explain Database Isolation Levels (Phantom Reads vs Dirty Reads).
"Isolation levels determine how much one transaction is shielded from another.
*   **Read Uncommitted:** The 'Wild West'. I can read data that isn't committed yet (**Dirty Read**). If that transaction rolls back, I'm holding invalid data.
*   **Read Committed:** The standard default. I only see committed data. However, if I run the same query twice, I might see new rows (**Non-repeatable Read**) if someone else committed in parallel.
*   **Repeatable Read:** Ensures that if I read a row, it stays the same for my whole transaction. But new rows might appear (**Phantom Reads**).
*   **Serializable:** The strictest level. It effectively runs transactions one by one. No side effects, but slowest performance."

#### Indepth
Most databases default to **Read Committed** (Postgres, SQL Server, Oracle). MySQL (InnoDB) defaults to **Repeatable Read**. Serializable is rarely used in high-load systems because it causes frequent deadlocks and requires retry logic.

---

### 17. How do you optimize a slow-running query?
"I have a checklist for this:
1.  **EXPLAIN Plan:** I interpret the query plan to see *what* the DB is actually doing. Is it doing a 'Seq Scan' (reading the whole table)? That's usually the culprit.
2.  **Indexes:** I check if the columns in the `WHERE` and `JOIN` clauses are indexed.
3.  **Selectivity:** I check if I'm selecting too much data (`SELECT *`).
4.  **SARGable:** I make sure I'm not doing functions on columns. `WHERE YEAR(date) = 2023` destroys the index usage. `WHERE date >= '2023-01-01'` allows the index to work."

#### Indepth
Reading an `EXPLAIN` output is a critical skill. Look for **Cost** estimates. A "Bitmap Heap Scan" or "Index Scan" is generally preferred over a "Sequential Scan" for large tables. Also, check for "Index Only Scans" (Covering Index), which are the fastest because they answer the query purely from the index tree without touching the table heap.

---

### 18. Optimistic vs Pessimistic Locking?
"**Pessimistic Locking** is 'locking the door so no one else can enter'. I count on conflict.
*   `SELECT ... FOR UPDATE` locks the rows I'm reading. No one can touch them until I commit. Good for high-conflict financial data.
 
**Optimistic Locking** is 'hoping for the best'. I read the data, and when I try to update it, I check if the version number changed.
*   I add a `Version` column. `UPDATE Tables SET Val=New, Version=2 WHERE ID=1 AND Version=1`. If 0 rows update, I know someone beat me to it, and I retry."

#### Indepth
Optimistic locking is preferred in web applications (stateless HTTP) because holding a database lock across a user's browser session is dangerous (what if they close the tab?). Pessimistic locking is better for short, critical batch jobs where collision is almost guaranteed.

---

### 19. Sharding vs Partitioning.
"**Partitioning** is slicing a table on *one* server.
*   I break my 1-billion-row `Logs` table into `Logs_Jan`, `Logs_Feb`, etc. The database engine handles this. It makes maintenance (deleting old logs) instant.
 
**Sharding** is distributing data across *multiple* servers.
*   I put Users A-M on Server 1 and N-Z on Server 2. This helps us scale write throughput beyond what a single machine can handle."

#### Indepth
Sharding introduces massive complexity. You lose ACID transactions across shards (Cross-shard joins are impossible or very slow). You usually need a specialized proxy (like Vitess) or application logic to route queries to the correct shard. Avoid sharding until you absolutely cannot scale vertically (bigger RAM/CPU) anymore.

---

### 20. CAP Theorem – Consistency vs Availability.
"The CAP Theorem states that in a distributed system, you can only pick two: **Consistency**, **Availability**, and **Partition Tolerance**.
 
Since network failures (Partitions) are inevitable, I realistically have to choose between **CP** (Consistency) or **AP** (Availability).
*   **CP (Banking):** If the network breaks, the system stops accepting writes to prevent data mismatch. Better to be down than wrong.
*   **AP (Social Media):** If the network breaks, the system keeps accepting posts, but your friend might not see your like immediately. Better to be up than perfect."

#### Indepth
RDBMS (SQL) systems traditionally lean towards **CA** (single node) or **CP** (clustered). NoSQL systems like Cassandra or DynamoDB often lean towards **AP** (Eventual Consistency), managing conflicts later to ensure 100% uptime.
