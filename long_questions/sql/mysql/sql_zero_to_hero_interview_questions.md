# 🟢 **Part 1: The Basics (Zero)**

### 1. What is SQL?
"Structured Query Language (SQL) is the standard language for relational database management systems. It is declarative, meaning you describe *what* data you want, not *how* to get it."

#### Indepth
SQL is an ANSI standard, but every vendor (MySQL, PostgreSQL, Oracle) has their own dialect extensions (e.g., T-SQL for SQL Server, PL/SQL for Oracle). While the core `SELECT`, `INSERT`, `UPDATE` syntax is consistent, functions and procedural extensions vary.

---

### 2. What are the different types of SQL commands?
"SQL commands are categorized into four main types:
*   **DDL (Data Definition):** Defines structure (`CREATE`, `ALTER`, `DROP`, `TRUNCATE`).
*   **DML (Data Manipulation):** Manipulates data (`SELECT`, `INSERT`, `UPDATE`, `DELETE`).
*   **DCL (Data Control):** Manages permissions (`GRANT`, `REVOKE`).
*   **TCL (Transaction Control):** Manages transactions (`COMMIT`, `ROLLBACK`)."

#### Indepth
The distinction between DDL and DML is crucial. In many databases, DDL statements contain an implicit commit, meaning they cannot be rolled back easily. DML statements are transaction-aware.

---

### 3. What is the difference between `DELETE`, `TRUNCATE`, and `DROP`?
"**DELETE** works row-by-row (slow), allows `WHERE` clauses, and can be rolled back.
**TRUNCATE** removes all rows by deallocating pages (fast) and resets auto-increment counters.
**DROP** removes the entire table structure and data permanently."

#### Indepth
`DELETE` generates a significant amount of transaction logs because it logs every deleted row individually. `TRUNCATE` is a metadata-only operation (in SQL Server/MySQL) that deallocates data pages, making it extremely fast but dangerous.

---

### 4. Explain Primary Key, Foreign Key, and Unique Key.
"A **Primary Key** uniquely identifies a record and cannot be `NULL`.
A **Unique Key** ensures uniqueness but allows one `NULL` value (in most DBs).
A **Foreign Key** creates a link between two tables to enforce referential integrity."

#### Indepth
A table can have only **one** Primary Key but **multiple** Unique Keys. Internally, a Primary Key usually creates a Clustered Index (default behavior in SQL Server/MySQL InnoDB), while Unique Keys create Non-Clustered Indexes.

---

### 5. What are SQL Constraints?
"Constraints are rules enforced on data columns to ensure validity.
Common ones are `NOT NULL`, `UNIQUE`, `PRIMARY KEY`, `FOREIGN KEY`, `CHECK` (validates condition), and `DEFAULT` (provides default value)."

#### Indepth
`CHECK` constraints are powerful for enforcing business logic at the database level (e.g., `CHECK (age >= 18)`). This ensures that no application bug can insert invalid data.

---

# 🟡 **Part 2: Intermediate Concepts**

### 6. What is the correct Order of Execution in an SQL Query?
"The logical processing order is specific and different from the written order:
1. `FROM` / `JOIN`
2. `WHERE`
3. `GROUP BY`
4. `HAVING`
5. `SELECT`
6. `DISTINCT`
7. `ORDER BY`
8. `LIMIT` / `OFFSET`"

#### Indepth
This order explains why you cannot use a column alias defined in the `SELECT` clause inside the `WHERE` clause. The `WHERE` clause is evaluated *before* the `SELECT` clause creates the alias.

---

### 7. Explain the different types of Joins.
"**INNER JOIN:** Returns only matching records.
**LEFT JOIN:** Returns all from Left table + matches from Right (NULL if no match).
**RIGHT JOIN:** Returns all from Right table + matches from Left.
**FULL JOIN:** Returns everything (match or no match).
**CROSS JOIN:** Cartesian product (Row count = TableA * TableB)."

#### Indepth
Keep in mind performance: `INNER JOIN` is typically the most efficient. `CROSS JOIN` is dangerous on large tables. If a `LEFT JOIN` has no matching valid row on the right, the columns from the right table will be `NULL`.

---

### 8. What is the difference between `WHERE` and `HAVING`?
"**WHERE** filters rows *before* they are grouped or aggregated.
**HAVING** filters groups *after* the aggregation has occurred."

#### Indepth
You cannot use aggregate functions like `COUNT()` or `SUM()` in a `WHERE` clause.
Performance Tip: Always filter as much as possible with `WHERE` before using `HAVING` to reduce the dataset size early.

---

### 9. What are Aggregate Functions?
"Functions that perform a calculation on a set of values to return a single scalar value.
Key examples: `COUNT()`, `SUM()`, `AVG()`, `MIN()`, `MAX()`."

#### Indepth
**Gotcha:** Aggregate functions (except `COUNT(*)`) ignore `NULL` values. `AVG(column)` will average only non-null values, which might skew your results if you expected `NULL` to count as 0.

---

### 10. Explain `UNION` vs `UNION ALL`.
"**UNION** combines result sets and *removes duplicates* (slower).
**UNION ALL** combines result sets and *keeps duplicates* (faster)."

#### Indepth
`UNION` performs a hidden sort/distinct operation to identify and remove duplicates. Unless you specifically need to remove duplicates, always use `UNION ALL` for better performance.

---

# 🔴 **Part 3: Advanced Concepts (Hero)**

### 11. What are Window Functions?
"Window functions perform calculations across a set of table rows that are related to the current row. Unlike `GROUP BY`, they do not collapse rows; they keep the original identity of the row.
Syntax: `function() OVER (PARTITION BY ... ORDER BY ...)`."

#### Indepth
They revolutionized SQL analytics. Use them for running totals, moving averages, and ranking. `PARTITION BY` is like `GROUP BY` but for the window scope.

---

### 12. Rank vs Dense_Rank vs Row_Number?
"**ROW_NUMBER():** Unique sequential integer (1, 2, 3, 4).
**RANK():** Rank with gaps for ties (1, 2, 2, 4).
**DENSE_RANK():** Rank without gaps for ties (1, 2, 2, 3)."

#### Indepth
If you are asked to "find the 3rd highest salary", always use `DENSE_RANK()`. If there are two people at #2, `RANK()` would skip #3 and go to #4, potentially returning no result for "3rd highest". `DENSE_RANK()` guarantees a #3 exists.

---

### 13. What is a Subquery? Types?
"A subquery is a query nested inside another.
**Scalar Subquery:** Returns a single value.
**Multi-row Subquery:** Returns a list.
**Correlated Subquery:** Depends on values from the outer query."

#### Indepth
**Correlated Subqueries** are performance killers because they execute *once for every row* of the outer query. Whenever possible, rewrite them as `JOIN`s or use `EXISTS`.

---

### 14. What is a Stored Procedure vs a Function?
"**Stored Procedures** can execute transactions, modify database state (DML), and have output parameters. They are called with `EXEC`.
**Functions** must return a value, cannot change database state, and are used inline in `SELECT` statements."

#### Indepth
User Defined Functions (UDFs) usually prevent parallel execution in query plans and can be slow if used on thousands of rows. Procedures are pre-compiled and often cached, offering better performance for complex logic.

---

### 15. What is an Index? Types?
"An index is a data structure (B-Tree) that speeds up data retrieval.
**Clustered Index:** Sorts the physical data rows (Only 1 per table).
**Non-Clustered Index:** A separate structure with pointers to the data (Multiple allowed)."

#### Indepth
Indexes improve Read speed but degrade Write speed (INSERT/UPDATE/DELETE) because the index must be updated on every change. Over-indexing is a common performance issue.

---

### 16. What is Normalization?
"The process of organizing data to reduce redundancy.
**1NF:** Atomic values (no comma-separated lists).
**2NF:** 1NF + Primary Key dependency.
**3NF:** 2NF + No transitive dependency (non-keys don't rely on other non-keys)."

#### Indepth
In high-scale systems or Data Warehouses, we often **Denormalize** (allow redundancy) to avoid expensive joins and improve read performance.

---

### 17. What are ACID properties?
"The four properties that guarantee reliable transactions:
**A**tomicity: All or nothing.
**C**onsistency: Data must be valid before and after.
**I**solation: Transactions don't interfere with each other.
**D**urability: Once committed, it's saved permanently."

#### Indepth
Isolation levels controls how strict the isolation is. Lower levels (`Read Uncommitted`) are fast but risky (Dirty Reads). Higher levels (`Serializable`) are safe but slow (Locking/Blocking).

---

# ⚔️ **Part 4: Scenario-Based Questions**

### 18. Find the Nth Highest Salary.
"The most robust way is using `DENSE_RANK()`:
```sql
SELECT salary FROM (
    SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) as rank
    FROM Employee
) as temp
WHERE rank = N;
```"

#### Indepth
Using `LIMIT N-1, 1` is MySQL specific and fails if there are duplicate salaries (ties). `DENSE_RANK` handles ties correctly and works on all modern databases (MySQL 8+, Postgres, SQL Server).

---

### 19. How to delete Duplicate Rows but keep one?
"Use a Common Table Expression (CTE) with `ROW_NUMBER()`:
```sql
WITH CTE AS (
    SELECT id, ROW_NUMBER() OVER (PARTITION BY col1, col2 ORDER BY id) as rn
    FROM table_name
)
DELETE FROM CTE WHERE rn > 1;
```"

#### Indepth
This assigns a number (1, 2, 3...) to each duplicate group. The row with `rn=1` is the "first" one (kept), and all others (`rn > 1`) are deleted.

---

### 20. How do you optimize a slow query?
"I follow a checklist:
1.  Check **EXPLAIN/Execution Plan** to see if indexes are used.
2.  Avoid `SELECT *`; fetch only needed columns.
3.  Ensure filters use **Indexed Columns**.
4.  Avoid functions on the LHS of `WHERE` (e.g., `WHERE YEAR(date) = 2020` kills the index; use range instead).
5.  Use `EXISTS` instead of `IN` for subqueries."

#### Indepth
The "Sargeable" (Search ARGument ABLE) concept is key. If you wrap a column in a function, the DB engine cannot use the index B-Tree navigation and must perform a full table scan.

---

### 21. Query to find employees hired in the last 6 months.
"Use date subtraction to filter records starting from the current date.
```sql
-- MySQL
SELECT * FROM Employee 
WHERE hire_date >= DATE_SUB(CURDATE(), INTERVAL 6 MONTH);
```"

#### Indepth
Date functions are specific to each database engine. For SQL Server, you would use `DATEADD(MONTH, -6, GETDATE())`. For PostgreSQL, you would use `NOW() - INTERVAL '6 months'`. Awareness of dialect differences is important.

---

# 🏢 **Part 5: Service vs Product Based Company Focus**

### 22. What do Service-Based Companies focus on in SQL interviews?
"Service-based companies (TCS, Wipro, Infosys) primarily focus on **Core Concepts**, **Basic Syntax**, and **Definitions**.
Expect straightforward questions like:
*   Difference between `DELETE` vs `TRUNCATE`.
*   Explain Joins with examples.
*   Basic `GROUP BY` and `HAVING` scenarios."

#### Indepth
These interviews aim to verify that you have a solid foundation and can write functional queries for typical business applications. They prioritize correct syntax and understanding of the fundamental DDL/DML operations.

---

### 23. What do Product-Based Companies focus on in SQL interviews?
"Product-based companies (Amazon, Google, Uber) focus heavily on **Query Logic**, **Performance Optimization**, and **Schema Design**.
Expect challenging questions like:
*   Handling Concurrency and Deadlocks.
*   Indexing internals (B-Trees) and Execution Plans.
*   ACID properties in distributed systems.
*   Designing a scalable schema (e.g., specifically for a high-traffic feature)."

#### Indepth
These interviews are designed to test your depth of knowledge. They want to know *how* the database works under the hood and how you would solve performance bottlenecks when scaling to millions of users. They value problem-solving over syntax memorization.
