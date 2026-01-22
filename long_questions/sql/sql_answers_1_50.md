## 🟢 Basic SQL (Select, Where, Order By, Group By)

### Question 1: What is the difference between `WHERE` and `HAVING`?

**Answer:**
Both clauses are used to filter records, but they apply at different stages of query execution.

| Feature | `WHERE` | `HAVING` |
| :--- | :--- | :--- |
| **Scope** | Filters individual rows. | Filters groups of rows (aggregated data). |
| **Usage** | Used before `GROUP BY`. | Used after `GROUP BY`. |
| **Aggregates** | Cannot use aggregate functions (e.g., `SUM`, `COUNT`). | Can use aggregate functions. |

**Example:**
```sql
-- Filter rows before grouping
SELECT Department, COUNT(*) 
FROM Employees 
WHERE Salary > 50000 
GROUP BY Department;

-- Filter groups after aggregation
SELECT Department, AVG(Salary) 
FROM Employees 
GROUP BY Department 
HAVING AVG(Salary) > 100000;
```

---

### Question 2: How do you retrieve unique records from a table?

**Answer:**
To retrieve unique records, use the `DISTINCT` keyword.

**Example:**
```sql
-- Get unique countries from Customers table
SELECT DISTINCT Country 
FROM Customers;
```

To find unique combinations of multiple columns:
```sql
SELECT DISTINCT Country, City 
FROM Customers;
```

---

### Question 3: Write a query to get the second highest salary from a table.

**Answer:**
There are multiple ways to achieve this:

**Method 1: Using `OFFSET` (Standard SQL/PostgreSQL/MySQL)**
```sql
SELECT Salary 
FROM Employees 
ORDER BY Salary DESC 
LIMIT 1 OFFSET 1;
```

**Method 2: Using Subquery**
```sql
SELECT MAX(Salary) 
FROM Employees 
WHERE Salary < (SELECT MAX(Salary) FROM Employees);
```

**Method 3: Using `DENSE_RANK()` Window Function (Handles duplicates)**
```sql
WITH RankedSalaries AS (
    SELECT Salary, DENSE_RANK() OVER (ORDER BY Salary DESC) as Rank
    FROM Employees
)
SELECT Salary 
FROM RankedSalaries 
WHERE Rank = 2;
```

---

### Question 4: How do you sort data in ascending and descending order?

**Answer:**
Use the `ORDER BY` clause. By default, it sorts in ascending (`ASC`) order. Use `DESC` for descending.

**Example:**
```sql
-- Ascending (Default)
SELECT * FROM Products 
ORDER BY Price;

-- Descending
SELECT * FROM Products 
ORDER BY Price DESC;

-- Multiple columns
SELECT * FROM Employees 
ORDER BY Department ASC, Salary DESC;
```

---

### Question 5: How do you filter records between two dates?

**Answer:**
You can use the `BETWEEN` operator or comparison operators (`>=` and `<=`).

**Example:**
```sql
-- Using BETWEEN (Inclusive)
SELECT * FROM Orders 
WHERE OrderDate BETWEEN '2023-01-01' AND '2023-12-31';

-- Using Operators
SELECT * FROM Orders 
WHERE OrderDate >= '2023-01-01' AND OrderDate <= '2023-12-31';
```

---

### Question 6: What is the difference between `COUNT(*)` and `COUNT(column)`?

**Answer:**
*   `COUNT(*)`: Counts **all rows** in the result set, including rows containing `NULL` values.
*   `COUNT(column)`: Counts only rows where the specific `column` is **NOT NULL**.

**Example:**
```sql
-- Table: Users
-- ID | Name
-- 1  | Alice
-- 2  | NULL
-- 3  | Bob

SELECT COUNT(*) FROM Users;      -- Result: 3
SELECT COUNT(Name) FROM Users;   -- Result: 2
```

---

### Question 7: How do you handle NULL values in SQL?

**Answer:**
`NULL` represents missing or undefined data.Standard comparison operators (`=`, `!=`) do not work with `NULL`.

**Checking for NULL:**
```sql
SELECT * FROM Employees WHERE ManagerID IS NULL;
SELECT * FROM Employees WHERE ManagerID IS NOT NULL;
```

**Replacing NULL values (ISNULL / COALESCE / IFNULL):**
```sql
-- COALESCE (Standard SQL - returns first non-null value)
SELECT Product, COALESCE(Discount, 0) AS FinalDiscount 
FROM Sales;
```

---

### Question 8: Write a query to return the number of employees in each department.

**Answer:**
Use `GROUP BY` with the `COUNT()` aggregate function.

**Example:**
```sql
SELECT Department, COUNT(*) AS EmployeeCount 
FROM Employees 
GROUP BY Department;
```

---

### Question 9: What is the difference between `GROUP BY` and `ORDER BY`?

**Answer:**
*   **`GROUP BY`**: Aggregates data by grouping rows with the same values in specified columns. It reduces the number of rows.
*   **`ORDER BY`**: Sorts the result set. It does not change the number of rows or aggregate data on its own.

**Example:**
```sql
-- Group by Department to get counts
SELECT Department, COUNT(*) 
FROM Employees 
GROUP BY Department;

-- Order the results alphabetically
SELECT Name, Salary 
FROM Employees 
ORDER BY Name;
```

---

### Question 10: How do you rename a column in a result set?

**Answer:**
Use the `AS` keyword (aliasing).

**Example:**
```sql
SELECT FirstName AS Name, Salary * 12 AS AnnualSalary 
FROM Employees;
```

---

## 🟡 Joins (Inner, Left, Right, Full)

### Question 11: Explain INNER JOIN with an example.

**Answer:**
`INNER JOIN` returns only the rows where there is a match in **both** tables.

**Example:**
```sql
SELECT E.Name, D.DepartmentName
FROM Employees E
INNER JOIN Departments D ON E.DepartmentID = D.DepartmentID;
```
*   Employees without a department are excluded.
*   Departments with no employees are excluded.

---

### Question 12: What is the difference between LEFT JOIN and RIGHT JOIN?

**Answer:**
*   **`LEFT JOIN` (Left Outer Join)**: Returns **all** records from the left table, and the matched records from the right table. If no match is found, the right side will contain `NULL`.
*   **`RIGHT JOIN` (Right Outer Join)**: Returns **all** records from the right table, and the matched records from the left table. If no match is found, the left side will contain `NULL`.

**Example:**
```sql
-- All employees, even those without a department
SELECT E.Name, D.DepartmentName
FROM Employees E
LEFT JOIN Departments D ON E.DepartmentID = D.DepartmentID;
```

---

### Question 13: When would you use FULL OUTER JOIN?

**Answer:**
Use `FULL OUTER JOIN` when you want to retain **all** rows from both tables. It returns:
1.  Matches between both tables.
2.  Rows from the left table with no match (Right columns are NULL).
3.  Rows from the right table with no match (Left columns are NULL).

**Example:**
Finding all employees and all departments, including employees with no department and departments with no employees.
```sql
SELECT E.Name, D.DepartmentName
FROM Employees E
FULL OUTER JOIN Departments D ON E.DepartmentID = D.DepartmentID;
```

---

### Question 14: Write a query using JOIN to get employee names and their manager names.

**Answer:**
This requires a **Self Join** on the Employees table.

**Example:**
```sql
SELECT E.Name AS Employee, M.Name AS Manager
FROM Employees E
LEFT JOIN Employees M ON E.ManagerID = M.EmployeeID;
```
*   `E` represents the Employee.
*   `M` represents the Manager.

---

### Question 15: How do you get records from one table that do not exist in another?

**Answer:**
**Method 1: LEFT JOIN with NULL check** (Usually most performant)
```sql
SELECT C.Name 
FROM Customers C
LEFT JOIN Orders O ON C.CustomerID = O.CustomerID
WHERE O.OrderID IS NULL;
```

**Method 2: NOT EXISTS**
```sql
SELECT Name 
FROM Customers C
WHERE NOT EXISTS (
    SELECT 1 FROM Orders O WHERE O.CustomerID = C.CustomerID
);
```

**Method 3: NOT IN** (Be careful with NULLs)
```sql
SELECT Name 
FROM Customers 
WHERE CustomerID NOT IN (SELECT CustomerID FROM Orders);
```

---

### Question 16: What is a self-join? Provide an example.

**Answer:**
A **self-join** is a regular join but the table is joined with itself. It is useful for comparing rows within the same table or dealing with hierarchical data.

**Example:**
Finding pairs of employees who live in the same city.
```sql
SELECT A.Name AS Employee1, B.Name AS Employee2, A.City
FROM Employees A
JOIN Employees B ON A.City = B.City
WHERE A.EmployeeID <> B.EmployeeID; -- Avoid matching with self
```

---

### Question 17: How do you join more than two tables?

**Answer:**
You simply chain the `JOIN` clauses. The result of the first join is joined with the third table, and so on.

**Example:**
```sql
SELECT C.CustomerName, O.OrderID, P.ProductName
FROM Customers C
JOIN Orders O ON C.CustomerID = O.CustomerID
JOIN OrderDetails OD ON O.OrderID = OD.OrderID
JOIN Products P ON OD.ProductID = P.ProductID;
```

---

### Question 18: Explain the concept of Cartesian join with a query.

**Answer:**
A **Cartesian Join** (or `CROSS JOIN`) returns the Cartesian product of two tables. It combines **every row** from the first table with **every row** from the second table.
If Table A has 10 rows and Table B has 5 rows, the result matches 10 * 5 = 50 rows.

**Example:**
```sql
SELECT P.ProductName, C.ColorName
FROM Products P
CROSS JOIN Colors C;
```
*   Useful for generating all possible combinations (e.g., Size x Color variants).

---

### Question 19: Write a query to find employees who don’t have a department.

**Answer:**
This suggests checking for `NULL` in the foreign key column or using a `LEFT JOIN` exclusion if verifying against a Departments table.

**Scenario A: Foreign key allows NULLs**
```sql
SELECT * FROM Employees WHERE DepartmentID IS NULL;
```

**Scenario B: Verifying against Departments table**
```sql
SELECT E.Name
FROM Employees E
LEFT JOIN Departments D ON E.DepartmentID = D.DepartmentID
WHERE D.DepartmentID IS NULL;
```

---

### Question 20: Difference between JOIN and UNION.

**Answer:**
| Feature | JOIN | UNION |
| :--- | :--- | :--- |
| **Direction** | Horizontal combination (adds columns). | Vertical combination (adds rows). |
| **Usage** | Relates data based on a matching key. | Appends matching datasets. |
| **Requirements** | Tables can have different structures. | Number and data type of columns must match. |

**Example:**
*   `JOIN`: "Show me Order details **next to** Customer details."
*   `UNION`: "Show me a list of Customers **and** a list of Suppliers in one single column of names."

---

## 🟠 Subqueries & CTEs

### Question 21: What is a correlated subquery?

**Answer:**
A **correlated subquery** is a subquery that depends on the outer query for its values. It executes **once for each row** processed by the outer query, which can be slow.

**Example:**
Find employees who earn more than the average salary of **their own** department.
```sql
SELECT E1.Name, E1.Salary, E1.DepartmentID
FROM Employees E1
WHERE E1.Salary > (
    SELECT AVG(E2.Salary)
    FROM Employees E2
    WHERE E2.DepartmentID = E1.DepartmentID
);
```

---

### Question 22: How is a subquery different from a join?

**Answer:**
*   **Join**: Combines columns from multiple tables into a single result set. generally more efficient and optimized by the database engine.
*   **Subquery**: Returns a result set (scalar, list, or table) used by another query. Often used to filter (`WHERE`, `HAVING`) or calculate derived values.

**Key difference:** Subqueries are often easier to read for logic like "Where X IN (List)", while Joins are better for retrieving data from multiple sources simultaneously.

---

### Question 23: Write a query using a subquery to find employees earning above average salary.

**Answer:**
This uses a simple (uncorrelated) subquery.

```sql
SELECT Name, Salary 
FROM Employees 
WHERE Salary > (SELECT AVG(Salary) FROM Employees);
```

---

### Question 24: What is a CTE? How is it different from a subquery?

**Answer:**
A **CTE (Common Table Expression)** is a temporary result set defined with `WITH` that exists only during the execution of a single query.

**Differences:**
1.  **Readability**: CTEs are named and defined at the top, making complex logic easier to follow than nested subqueries.
2.  **Reusability**: A CTE can be referenced multiple times in the main query.
3.  **Recursion**: CTEs support recursion (useful for hierarchies), subqueries do not.

**Example:**
```sql
WITH DeptAvg AS (
    SELECT DepartmentID, AVG(Salary) as AvgSal 
    FROM Employees 
    GROUP BY DepartmentID
)
SELECT E.Name, E.Salary
FROM Employees E
JOIN DeptAvg D ON E.DepartmentID = D.DepartmentID
WHERE E.Salary > D.AvgSal;
```

---

### Question 25: How can you use a recursive CTE? Example: find all levels of a hierarchy.

**Answer:**
Recursive CTEs are used for hierarchical data (e.g., Org charts, category trees).

**Example:** Find a manager and all their subordinates (direct and indirect).
```sql
WITH RECURSIVE OrgChart AS (
    -- Anchor member: Start with top manager
    SELECT EmployeeID, Name, ManagerID, 1 as Level
    FROM Employees
    WHERE ManagerID IS NULL 

    UNION ALL

    -- Recursive member: Join back to the CTE
    SELECT E.EmployeeID, E.Name, E.ManagerID, O.Level + 1
    FROM Employees E
    INNER JOIN OrgChart O ON E.ManagerID = O.EmployeeID
)
SELECT * FROM OrgChart;
```

---

### Question 26: Can you use a subquery in a `SELECT` clause? Example?

**Answer:**
Yes, this is called a **scalar subquery** (must return a single value).

**Example:**
Listing each employee and the total count of employees in their department alongside.
```sql
SELECT 
    Name, 
    DepartmentID,
    (SELECT COUNT(*) FROM Employees E2 WHERE E2.DepartmentID = E1.DepartmentID) AS DeptHeadCount
FROM Employees E1;
```
*Note: A Window function (`COUNT(*) OVER(...)`) is usually more efficient for this.*

---

### Question 27: How do you write a query to get the top N rows per group?

**Answer:**
Use a Window Function like `ROW_NUMBER()`, `RANK()`, or `DENSE_RANK()`.

**Example:** Top 3 highest earners in each department.
```sql
WITH RankedEmployees AS (
    SELECT 
        Name, 
        DepartmentID, 
        Salary,
        ROW_NUMBER() OVER (PARTITION BY DepartmentID ORDER BY Salary DESC) as Rank
    FROM Employees
)
SELECT * 
FROM RankedEmployees 
WHERE Rank <= 3;
```

---

### Question 28: Explain how `WITH` clause works in SQL.

**Answer:**
The `WITH` clause defines one or more Common Table Expressions (CTEs).
1.  It creates a temporary "view" or result set.
2.  Scope is limited to the immediate `SELECT`, `INSERT`, `UPDATE`, or `DELETE` statement following it.
3.  It helps break down complex logic into modular steps.

**Syntax:**
```sql
WITH CTE_Name AS (
    SELECT ...
),
CTE_Name2 AS (
    SELECT ...
)
SELECT * FROM CTE_Name JOIN CTE_Name2 ...
```

---

### Question 29: Can you use CTEs in UPDATE or DELETE statements?

**Answer:**
Yes, usually. This is very powerful for deleting duplicates or updating based on complex joins.

**Example:** Delete duplicate rows keeping the one with the lowest ID.
```sql
WITH Duplicates AS (
    SELECT ID, 
           ROW_NUMBER() OVER (PARTITION BY Name, Email ORDER BY ID) as RowNum
    FROM Users
)
DELETE FROM Duplicates WHERE RowNum > 1;
```

---

### Question 30: What are the performance considerations between CTE and derived tables?

**Answer:**
*   **Modern Optimizers:** In most modern databases (SQL Server, Postgres, Oracle), the optimizer treats non-recursive CTEs and Derived Tables (Subqueries in FROM) almost identically. It may expand the definition or materialize it based on cost.
*   **Materialization:**
    *   **Postgres (pre-v12)**: Always materialized CTEs (calculated them once), which was a "optimization fence". Now it can inline them.
    *   **SQL Server**: Doesn't implicitly materialize; it treats CTEs like macros (inlines code), meaning if you reference a CTE twice, it might run the underlying query twice unless you force materialization (temp table).
*   **General Rule**: Use CTEs for readability. Use Temp Tables for performance if the intermediate result is large and referenced multiple times.

---

## 🔵 Data Aggregation & Analysis

### Question 31: What does `ROLLUP` do in SQL?

**Answer:**
`ROLLUP` is an extension of `GROUP BY` that creates subtotals and grand totals. It moves strictly up the hierarchy provided in the columns.

**Example:**
```sql
SELECT Region, Country, SUM(Sales)
FROM SalesData
GROUP BY ROLLUP (Region, Country);
```
**Output produces rows for:**
1.  (Region, Country) -> Specific totals
2.  (Region, NULL)    -> Subtotal per Region
3.  (NULL, NULL)      -> Grand Total

---

### Question 32: What is `CUBE` and how is it different from `ROLLUP`?

**Answer:**
*   **`ROLLUP`**: Hierarchical subtotals (e.g., Year > Month > Day). N+1 levels of aggregation.
*   **`CUBE`**: Generates subtotals for **all possible combinations** of the grouping columns. 2^N combinations.

**Example:** `GROUP BY CUBE(Year, Product)` generates totals for:
1.  Year, Product
2.  Year only
3.  Product only
4.  Grand Total

---

### Question 33: What are window functions? Give a practical example.

**Answer:**
Window functions perform calculations across a set of table rows that are related to the current row. Unlike `GROUP BY`, they **do not collapse** rows.

**Syntax:** `FUNCTION() OVER (PARTITION BY ... ORDER BY ...)`

**Example:** Calculate each employee's salary as a percentage of their department's total salary.
```sql
SELECT 
    Name, 
    Salary,
    SUM(Salary) OVER (PARTITION BY Department) as DeptTotal,
    (Salary / SUM(Salary) OVER (PARTITION BY Department)) * 100 as PctOfTotal
FROM Employees;
```

---

### Question 34: Use `RANK()` and `DENSE_RANK()` to assign rankings within groups.

**Answer:**
*   `RANK()`: Skips numbers if there are ties (e.g., 1, 2, 2, 4).
*   `DENSE_RANK()`: No gaps (e.g., 1, 2, 2, 3).

**Example:**
```sql
SELECT 
    Name, 
    Score,
    RANK() OVER (ORDER BY Score DESC) as RankVal,
    DENSE_RANK() OVER (ORDER BY Score DESC) as DenseRankVal
FROM ExamResults;
```

---

### Question 35: What does `PARTITION BY` do in window functions?

**Answer:**
`PARTITION BY` behaves like `GROUP BY` but for window functions. It divides the result set into partitions (groups) and the window function resets effectively for each partition.

**Example:** Resetting row number for each department.
```sql
ROW_NUMBER() OVER (PARTITION BY DepartmentID ORDER BY Salary DESC)
```
This counts 1, 2, 3... for Dept A, then restarts 1, 2, 3... for Dept B.

---

### Question 36: Write a query using `LEAD()` and `LAG()` functions.

**Answer:**
*   `LAG()`: Access data from the *previous* row.
*   `LEAD()`: Access data from the *next* row.

**Example:** Compare today's sales with yesterday's.
```sql
SELECT 
    SaleDate, 
    Revenue,
    LAG(Revenue, 1) OVER (ORDER BY SaleDate) as PrevDayRevenue,
    Revenue - LAG(Revenue, 1) OVER (ORDER BY SaleDate) as DailyChange
FROM DailySales;
```

---

### Question 37: How would you calculate a running total?

**Answer:**
Use `SUM()` as a window function with an `ORDER BY` clause inside.

**Example:**
```sql
SELECT 
    OrderDate, 
    Amount,
    SUM(Amount) OVER (ORDER BY OrderDate) as RunningTotal
FROM Orders;
```

---

### Question 38: Difference between `ROW_NUMBER()` and `RANK()`.

**Answer:**
*   **`ROW_NUMBER()`**: Always assigns a unique integer (1, 2, 3, 4). If there's a tie, the order is arbitrary (unless a tie-breaker is specified).
*   **`RANK()`**: Assigns the same rank to ties, but skips the next numbers (1, 2, 2, 4).

---

### Question 39: How do you find duplicate rows in a table?

**Answer:**
Group by the columns that define uniqueness and filter for count > 1.

**Example:**
```sql
SELECT Email, COUNT(*) 
FROM Users 
GROUP BY Email 
HAVING COUNT(*) > 1;
```

---

### Question 40: Write a query to pivot data in SQL.

**Answer:**
**Scenario:** Convert rows (Sales by Month) into columns (Jan, Feb, Mar).

**Method 1: CASE statements (Standard SQL)**
```sql
SELECT 
    Product,
    SUM(CASE WHEN Month = 'Jan' THEN Sales ELSE 0 END) AS Jan_Sales,
    SUM(CASE WHEN Month = 'Feb' THEN Sales ELSE 0 END) AS Feb_Sales
FROM SalesData
GROUP BY Product;
```

**Method 2: PIVOT operator (SQL Server / Oracle)**
```sql
SELECT Product, [Jan], [Feb]
FROM (SELECT Product, Month, Sales FROM SalesData) AS SourceTable
PIVOT (
    SUM(Sales) FOR Month IN ([Jan], [Feb])
) AS PivotTable;
```

---

## 🔴 Advanced SQL

### Question 41: What is normalization? Explain different normal forms.

**Answer:**
Normalization is the process of organizing data to reduce redundancy and inconsistent dependency.

*   **1NF (First Normal Form)**: Atomic values (no list/comma-separated strings in cells), unique rows.
*   **2NF**: 1NF + No partial dependency (all non-key columns depend on the *entire* primary key).
*   **3NF**: 2NF + No transitive dependency (non-key columns depend *only* on the primary key, not other non-key columns).
*   **BCNF (Boyce-Codd)**: Stricter version of 3NF.

---

### Question 42: What is denormalization and when is it used?

**Answer:**
**Denormalization** is the process of adding redundancy to an already normalized database (usually merging tables).

**Use Case:**
*   To improve read performance (avoiding complex joins).
*   Data Warehousing (OLAP) and Star Schemas.
*   When read speed is critical and write speed/storage space is less of a concern.

---

### Question 43: How would you optimize a slow-running query?

**Answer:**
1.  **Analyze Execution Plan**: Use `EXPLAIN` or `EXPLAIN ANALYZE` to see what the database is doing (Full Table Scan vs Index Scan).
2.  **Indexing**: Ensure columns in `WHERE`, `JOIN`, and `ORDER BY` clauses are indexed.
3.  **Avoid SELECT \***: Select only necessary columns to reduce I/O.
4.  **SARGable queries**: Avoid functions on columns in `WHERE` clause (e.g., `WHERE YEAR(Date) = 2023` -> `WHERE Date BETWEEN '2023-01-01' AND '2023-12-31'`).
5.  **Joins**: Check for Cartesian products; ensure statistics are up to date.

---

### Question 44: What are indexes? Types? When to use?

**Answer:**
An index is a data structure (B-Tree, Hash) that improves retrieval speed.

**Types:**
*   **B-Tree**: Default, good for range queries (`<`, `>`, `=`) and sorting.
*   **Hash**: Good for equality checks (`=`), not for ranges.
*   **Bitmap**: Good for low-cardinality data (few unique values like Gender, Status).

**When to use:** On columns frequently searched, filtered, or joined.
**When NOT to use:** On small tables, or columns that are frequently updated (indexes slow down INSERT/UPDATE).

---

### Question 45: What are stored procedures? How are they different from functions?

**Answer:**
*   **Stored Procedure**: A batch of SQL code saved in the DB. Can perform modifications (`INSERT`/`UPDATE`), call other procedures, and handle transactions. Can return multiple result sets.
*   **Function (UDF)**: Designed to calculate and return a value. specific rules vary by DB, but generally, functions are used in `SELECT`/`WHERE` clauses and cannot change database state (no side effects).

---

### Question 46: How do you implement transactions in SQL?

**Answer:**
Transactions ensure a set of operations either all succeed or all fail (ACID).

**Syntax:**
```sql
BEGIN TRANSACTION; -- START TRANSACTION

UPDATE Accounts SET Balance = Balance - 100 WHERE ID = 1;
UPDATE Accounts SET Balance = Balance + 100 WHERE ID = 2;

-- Check for errors
IF @@ERROR <> 0
    ROLLBACK TRANSACTION;
ELSE
    COMMIT TRANSACTION;
```

---

### Question 47: What is ACID in SQL databases?

**Answer:**
Properties that guarantee reliable processing of transactions:
*   **A - Atomicity**: All or nothing. If one part fails, the entire transaction rolls back.
*   **C - Consistency**: Database transitions from one valid state to another (constraints enforced).
*   **I - Isolation**: Transactions happen independently (concurrency control).
*   **D - Durability**: Once committed, changes are permanent (saved to disk) even if power fails.

---

### Question 48: Explain isolation levels: Read Uncommitted, Committed, Repeatable Read, Serializable.

**Answer:**
Isolation levels control visibility of changes between concurrent transactions.

1.  **Read Uncommitted**: Can see uncommitted changes (Dirty Reads). Fastest, least safe.
2.  **Read Committed**: Default in many DBs (Postgres, SQL Server, Oracle). Can only see committed changes. (Prevents Dirty Reads).
3.  **Repeatable Read**: Ensures if you read a row twice, you get the same data. Prevents Non-Repeatable reads (updates) but valid for Phantom reads (new inserts).
4.  **Serializable**: Strict sequential execution. Slowest, essentially locks the range. Prevents all concurrency anomalies.

---

### Question 49: What is the use of the `EXPLAIN` or `EXPLAIN PLAN`?

**Answer:**
It shows the **execution plan** the query optimizer has chosen.
It reveals:
*   Whether indexes are being used.
*   Join algorithms (Nested Loop, Hash Join, Merge Join).
*   Estimated cost (CPU/IO).
*   Number of rows scanned.

Essential for performance tuning.

---

### Question 50: Difference between clustered and non-clustered indexes.

**Answer:**
*   **Clustered Index**:
    *   Determines the **physical order** of data on disk.
    *   Leaf nodes *are* the data pages.
    *   Only **one** per table (usually the Primary Key).
    *   Faster for retrieving range data.
*   **Non-Clustered Index**:
    *   Stored separately from the data; contains a pointer to the physical row (clustered index key or Row ID).
    *   Can have **multiple** per table.
    *   Slightly slower lookups (requires an extra hop).
