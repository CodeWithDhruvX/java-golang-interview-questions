# SQL Interview Questions & Answers

## 🔹 1. Basic SQL (Questions 1-10)

**Q1: What is the difference between `WHERE` and `HAVING`?**
`WHERE` filters rows *before* grouping/aggregation. `HAVING` filters groups *after* aggregation.

**Q2: How do you retrieve unique records from a table?**
Use the `DISTINCT` keyword: `SELECT DISTINCT column FROM table;`.

**Q3: Write a query to get the second highest salary from a table.**
`SELECT MAX(salary) FROM emp WHERE salary < (SELECT MAX(salary) FROM emp);` OR `SELECT salary FROM emp ORDER BY salary DESC OFFSET 1 LIMIT 1;`.

**Q4: How do you sort data in ascending and descending order?**
Use `ORDER BY col_name ASC` (default) or `ORDER BY col_name DESC`.

**Q5: How do you filter records between two dates?**
Use the `BETWEEN` operator: `SELECT * FROM table WHERE date_col BETWEEN '2023-01-01' AND '2023-12-31';`.

**Q6: What is the difference between `COUNT(*)` and `COUNT(column)`?**
`COUNT(*)` counts total rows including NULLs. `COUNT(column)` counts only rows where the column is NOT NULL.

**Q7: How do you handle NULL values in SQL?**
Use `IS NULL` for checking. Use `COALESCE(col, default_val)` or `IFNULL(col, val)` to replace/handle them in results.

**Q8: Write a query to return the number of employees in each department.**
`SELECT dept_id, COUNT(*) FROM employees GROUP BY dept_id;`.

**Q9: What is the difference between `GROUP BY` and `ORDER BY`?**
`GROUP BY` aggregates data into summary rows by category. `ORDER BY` sorts the output result set.

**Q10: How do you rename a column in a result set?**
Use the `AS` alias keyword: `SELECT column_name AS new_name FROM table;`.

---

## 🔹 2. Joins (Questions 11-20)

**Q11: Explain INNER JOIN with an example.**
Returns records that have matching values in both tables. `SELECT * FROM Emp JOIN Dept ON Emp.dept_id = Dept.id;`.

**Q12: What is the difference between LEFT JOIN and RIGHT JOIN?**
`LEFT JOIN` returns all rows from the left table + matches from right. `RIGHT JOIN` returns all from right + matches from left.

**Q13: When would you use FULL OUTER JOIN?**
When you want all records from both tables, filling with NULLs where no match exists on either side.

**Q14: Write a query using JOIN to get employee names and their manager names.**
Self Join: `SELECT E.name AS Employee, M.name AS Manager FROM Emp E JOIN Emp M ON E.manager_id = M.id;`.

**Q15: How do you get records from one table that do not exist in another?**
`LEFT JOIN ... WHERE right.id IS NULL` or `WHERE id NOT IN (SELECT id FROM ...)`.

**Q16: What is a self-join? Provide an example.**
Joining a table to itself. Example: Hierarchy (Employees vs Managers) or finding duplicates in same table.

**Q17: How do you join more than two tables?**
Chain joins: `SELECT * FROM A JOIN B ON A.id = B.id JOIN C ON B.id = C.id;`.

**Q18: Explain the concept of Cartesian join with a query.**
Returns the Cartesian product (all combinations of rows). `SELECT * FROM TableA, TableB;` or `CROSS JOIN`.

**Q19: Write a query to find employees who don’t have a department.**
`SELECT * FROM Employees WHERE dept_id IS NULL;` OR `SELECT * FROM Employees E LEFT JOIN Dept D ON E.dept_id = D.id WHERE D.id IS NULL`.

**Q20: Difference between JOIN and UNION.**
`JOIN` combines columns from tables horizontally based on a relationship. `UNION` combines result sets (rows) vertically.

---

## 🔹 3. Subqueries & CTEs (Questions 21-30)

**Q21: What is a correlated subquery?**
A subquery that references a column from the outer query. It executes once for each row processed by the outer query.

**Q22: How is a subquery different from a join?**
Subqueries are often used in `WHERE`/`SELECT` clauses to check existence/value. Joins usually combine datasets. Joins are generally faster.

**Q23: Write a query using a subquery to find employees earning above average salary.**
`SELECT * FROM Emp WHERE salary > (SELECT AVG(salary) FROM Emp);`.

**Q24: What is a CTE? How is it different from a subquery?**
CTE (`WITH` clause) is a named temporary result set. It is more readable and reusable within the same query than nested subqueries.

**Q25: How can you use a recursive CTE? Example: find all levels of a hierarchy.**
Use `WITH RECURSIVE`. Base Select `UNION ALL` Recursive Select referencing the CTE name.

**Q26: Can you use a subquery in a `SELECT` clause? Example?**
Yes. `SELECT name, (SELECT COUNT(*) FROM orders WHERE user_id = u.id) FROM users u;`.

**Q27: How do you write a query to get the top N rows per group?**
Use Window Function: `SELECT * FROM (SELECT *, ROW_NUMBER() OVER(PARTITION BY group ORDER BY val DESC) as rn FROM table) WHERE rn <= N;`.

**Q28: Explain how `WITH` clause works in SQL.**
Defines a Common Table Expression (CTE) before the main query. `WITH CTE AS (...) SELECT * FROM CTE;`.

**Q29: Can you use CTEs in UPDATE or DELETE statements?**
Yes. `WITH OldRows AS (...) DELETE FROM Table WHERE id IN (SELECT id FROM OldRows);`.

**Q30: What are the performance considerations between CTE and derived tables?**
Optimizer often treats them similarly, but some DBs (Postgres) materialize CTEs (calc once, cache), which can be faster or slower depending on index usage.

---

## 🔹 4. Data Aggregation & Analysis (Questions 31-40)

**Q31: What does `ROLLUP` do in SQL?**
Extension of `GROUP BY`. Generates subtotals and grand totals for hierarchical groups (e.g., Year, Month, Day).

**Q32: What is `CUBE` and how is it different from `ROLLUP`?**
`CUBE` generates subtotals for *all possible combinations* of grouping columns. `ROLLUP` does hierarchical only.

**Q33: What are window functions? Give a practical example.**
Functions that perform calc across a set of table rows related to current row. `SUM(amt) OVER (ORDER BY date)` for running total.

**Q34: Use `RANK()` and `DENSE_RANK()` to assign rankings within groups.**
`RANK()` skips numbers for ties (1, 1, 3). `DENSE_RANK()` does not skip (1, 1, 2).

**Q35: What does `PARTITION BY` do in window functions?**
Divides the result set into partitions (groups) allowing the window function to calculate separately for each group.

**Q36: Write a query using `LEAD()` and `LAG()` functions.**
`SELECT salary, LAG(salary) OVER(ORDER BY date) as prev_salary FROM emp;` (Gets previous row's value).

**Q37: How would you calculate a running total?**
`SELECT date, amount, SUM(amount) OVER (ORDER BY date) as running_total FROM sales;`.

**Q38: Difference between `ROW_NUMBER()` and `RANK()`.**
`ROW_NUMBER()` gives unique ID (1, 2, 3) even for ties. `RANK()` gives same ID for ties (1, 1, 3).

**Q39: How do you find duplicate rows in a table?**
`SELECT name, COUNT(*) FROM table GROUP BY name HAVING COUNT(*) > 1;`.

**Q40: Write a query to pivot data in SQL.**
Use `CASE` inside `SUM` or `PIVOT` function. `SELECT id, SUM(CASE WHEN month='Jan' THEN val END) as Jan, ... FROM data GROUP BY id`.

---

## 🔹 5. Advanced SQL (Questions 41-50)

**Q41: What is normalization? Explain different normal forms.**
Organizing data to reduce redundancy. 1NF (Atomic), 2NF (No partial dependency), 3NF (No transitive dependency).

**Q42: What is denormalization and when is it used?**
Adding redundancy (combining tables) to improve Read performance, often used in OLAP/Reporting.

**Q43: How would you optimize a slow-running query?**
Check `EXPLAIN` plan, Add Indexes, Avoid `SELECT *`, Remove unnecessary Joins, Use proper types.

**Q44: What are indexes? Types? When to use?**
Structures to speed up retrieval. Types: B-Tree, Bitmap, Hash. Use on columns frequently used in `WHERE` or `JOIN`.

**Q45: What are stored procedures? How are they different from functions?**
Code stored in DB. Proc can execute transactions, no return value required. Function must return value, cannot change DB state usually.

**Q46: How do you implement transactions in SQL?**
`BEGIN TRANSACTION; ... SQL ... COMMIT;` (or `ROLLBACK` on error). Ensures ACID.

**Q47: What is ACID in SQL databases?**
Atomicity (All or nothing), Consistency (Valid rules), Isolation (Concurrent/Safe), Durability (Saved permanently).

**Q48: Explain isolation levels: Read Uncommitted, Committed, Repeatable Read, Serializable.**
Levels of locking. Read Uncommitted (Dirty reads). Committed (No dirty). Repeatable (No phantom updates). Serializable (Strict sequential).

**Q49: What is the use of the `EXPLAIN` or `EXPLAIN PLAN`?**
Shows how the query engine executes a statement (Indexes used, Scans type). Essential for debugging performance.

**Q50: Difference between clustered and non-clustered indexes.**
Clustered: Sorts actual data rows on disk (Only 1 per table). Non-clustered: Separate structure pointing to data rows (Many allowed).

---

## 🔹 6. Real-world Scenarios (Questions 51-60)

**Q51: Find users who made purchases on consecutive days.**
Self join on `T1.date = T2.date - 1`. Or use `LEAD(date)` window function and check diff is 1 day.

**Q52: Write a query to find repeat customers.**
`SELECT user_id FROM orders GROUP BY user_id HAVING COUNT(order_id) > 1`.

**Q53: Find top 3 selling products in each category.**
Window `RANK() <= 3`. `SELECT * FROM (SELECT *, RANK() OVER(PARTITION BY category ORDER BY sales DESC) as rank FROM prod) WHERE rank <= 3`.

**Q54: Calculate customer lifetime value (CLV).**
`SELECT user_id, SUM(order_amount) FROM orders GROUP BY user_id`.

**Q55: Track changes in customer status using SQL (SCD Type 2).**
Use `start_date` and `end_date` columns. Update old record `end_date`, insert new record with new `start_date`.

**Q56: Identify users with declining monthly usage.**
Compare `SUM(usage)` of current month vs previous month using `LAG()`.

**Q57: Write a query to calculate retention rate month over month.**
(Users active this month AND last month) / (Users active last month).

**Q58: Get average time between user signup and first purchase.**
`SELECT AVG(DATEDIFF(order_date, signup_date)) FROM users JOIN orders ON ...`.

**Q59: Segment users based on activity (e.g., high, medium, low).**
Use `CASE`: `CASE WHEN orders > 10 THEN 'High' WHEN orders > 5 THEN 'Medium' ELSE 'Low' END`.

**Q60: Detect anomalies in daily transaction volumes.**
Calculate Moving Average. Flag days where volume > (Avg + 3*StdDev).

---

## 🔹 7. Date & Time Manipulations (Questions 61-70)

**Q61: How do you extract year/month/day from a date?**
`YEAR(date)`, `MONTH(date)`, `DAY(date)` or `EXTRACT(YEAR FROM date)`.

**Q62: Write a query to get the last day of each month.**
`LAST_DAY(date)` (MySQL/Oracle) or `EOMONTH(date)` (SQL Server).

**Q63: Find the number of working days between two dates.**
Calc total days minus weekends. Usually requires a calendar utility table or complex function masking Sun/Sat.

**Q64: Calculate age from date of birth.**
`FLOOR(DATEDIFF(CURRENT_DATE, dob) / 365.25)` or `TIMESTAMPDIFF(YEAR, dob, NOW())`.

**Q65: How do you handle time zones in SQL?**
Store in UTC. Convert on display: `CONVERT_TZ(col, 'UTC', 'America/New_York')` or `AT TIME ZONE`.

**Q66: Write a query to get users who signed up in the past 7 days.**
`WHERE signup_date >= DATE_SUB(NOW(), INTERVAL 7 DAY)` / `CURRENT_DATE - 7`.

**Q67: Get the number of orders by week.**
`GROUP BY YEAR(date), WEEK(date)`.

**Q68: How do you truncate a date to the start of the month?**
`DATE_FORMAT(date, '%Y-%m-01')` or `TRUNC(date, 'MONTH')`.

**Q69: How to calculate time difference in hours between two timestamps?**
`TIMESTAMPDIFF(HOUR, start, end)` or `(end - start) * 24` (Date arithmetic).

**Q70: Write a query to get users inactive for more than 30 days.**
`SELECT user_id FROM logins GROUP BY user_id HAVING MAX(login_date) < DATE_SUB(NOW(), INTERVAL 30 DAY)`.

---

## 🔹 8. Data Cleaning & Transformation (Questions 71-80)

**Q71: How do you remove duplicates in a table?**
`DELETE FROM T WHERE id NOT IN (SELECT MIN(id) FROM T GROUP BY col1, col2)`. Or using CTE/Window `ROW_NUMBER()`.

**Q72: How do you handle NULLs in aggregate functions?**
Aggregates like `SUM/AVG` ignore NULLs automatically. Use `COALESCE` if you want to treat them as 0.

**Q73: Use `CASE` to categorize data into buckets (e.g., age groups).**
`SELECT CASE WHEN age < 18 THEN 'Child' WHEN age < 65 THEN 'Adult' ELSE 'Senior' END ...`.

**Q74: Replace NULL values with default values using `COALESCE()`.**
`SELECT COALESCE(phone, 'N/A') FROM users;`.

**Q75: How do you clean up bad data (e.g., phone numbers with text)?**
Use Regex `REGEXP_REPLACE` to strip non-numeric chars.

**Q76: Write a query to standardize country names.**
`UPDATE Users SET country = 'USA' WHERE country IN ('United States', 'US', 'U.S.');`.

**Q77: How do you extract numbers from a string?**
Regex: `REGEXP_SUBSTR(str, '[0-9]+')`.

**Q78: Normalize string formats (e.g., capitalize all names).**
`UPPER(name)`, `LOWER(name)`, or `INITCAP(name)` (DB dependent).

**Q79: Identify and fix inconsistent data (e.g., dates as strings).**
Use `CAST` or `STR_TO_DATE` to validate. Find records that return NULL on conversion.

**Q80: Write SQL to de-duplicate based on business rules.**
Use Window Function: `ROW_NUMBER() OVER (PARTITION BY id ORDER BY updated_at DESC)` -> Keep 1 (Latest).

---

## 🔹 9. Schema & DDL (Questions 81-90)

**Q81: How do you create a new table with constraints?**
`CREATE TABLE T (id INT PRIMARY KEY, email VARCHAR(50) UNIQUE, age INT CHECK (age > 0));`.

**Q82: What’s the difference between `DELETE`, `TRUNCATE`, and `DROP`?**
`DELETE`: Remove rows (loggable, slow). `TRUNCATE`: Empty table (DDL, fast, resets identities). `DROP`: Destroy table structure.

**Q83: How do you alter an existing table (add/remove columns)?**
`ALTER TABLE T ADD column_name type;` / `DROP COLUMN column_name;`.

**Q84: How to enforce referential integrity?**
Use `FOREIGN KEY (col) REFERENCES other_table(id) ON DELETE CASCADE/RESTRICT`.

**Q85: Explain primary key vs unique key.**
PK: Unique + Not Null + Only 1 per table. Unique: Unique + Nulls allowed (usually) + Multiple allowed.

**Q86: What is a composite key?**
A Primary Key consisting of two or more columns to uniquely identify a row.

**Q87: What are views? Pros and cons?**
Virtual table based on a query. Pros: Security, Simplicity. Cons: Performance (if complex), Dependence on underlying schema.

**Q88: How do you grant/revoke permissions in SQL?**
`GRANT SELECT ON table TO user;` / `REVOKE SELECT ON table FROM user;`.

**Q89: How do you copy schema and data from one table to another?**
`CREATE TABLE new_table AS SELECT * FROM old_table` (CTAS).

**Q90: What are sequences and auto-increment columns?**
Generators for unique numeric IDs. `AUTO_INCREMENT` (MySQL) or `SEQUENCE` object (Oracle/Postgres).

---

## 🔹 10. Miscellaneous / Optimization / Troubleshooting (Questions 91-100)

**Q91: What is SQL injection and how to prevent it?**
Malicious code inserted into query strings. Prevent using Parameterized Queries (PreparedStatement).

**Q92: How to handle pagination in SQL?**
`LIMIT 10 OFFSET 20` (MySQL/PG). `OFFSET 20 ROWS FETCH NEXT 10 ROWS ONLY` (SQL Server).

**Q93: What tools do you use to debug slow SQL queries?**
`EXPLAIN`, DB Profiler, Slow Query Log, `pg_stat_statements`.

**Q94: When would you use temp tables or table variables?**
Intermediate processing for complex logic/procedures. Temp Tables (`#Table`) live for session. Vars (`@Table`) live for batch.

**Q95: What is a materialized view?**
A view whose result is physically stored (cached) on disk and refreshed periodically. Good for heavy aggregation.

**Q96: Difference between OLTP and OLAP systems?**
OLTP: Transactional (Inserts/Updates/User-facing, Normalized). OLAP: Analytical (Reporting/Reads/History, Denormalized).

**Q97: How do you monitor long-running queries in production?**
`SHOW PROCESSLIST` (MySQL), `pg_stat_activity` (Postgres). Kill blocks if needed.

**Q98: Explain batch vs streaming queries.**
Batch: Process fixed bounded data (Historical). Streaming: Continuous unbounded data (Real-time).

**Q99: What is sharding and partitioning in SQL databases?**
Partitioning: Dividing table within one server (e.g. by Year). Sharding: Distributing data across multiple physical servers.

**Q100: How do you migrate large amounts of data efficiently using SQL?**
Batched inserts, Disable Indexes during load, Use `COPY`/`BULK INSERT`, ETL tools.

---

## 🔹 11. Intermediate Queries & Logic (Questions 101-110)

**Q101: Write a query to transpose rows into columns.**
Use the `PIVOT` function or conditional aggregation: `MAX(CASE WHEN type='A' THEN val END)`.

**Q102: How would you calculate the median in SQL?**
`PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY col)` (Window Function) or find middle row using count/offsets.

**Q103: How do you compare records in a table against a historical version of itself?**
Self Join on `ID` and `Version` (or Date). `T_Curr.val != T_Hist.val`.

**Q104: How would you convert a comma-separated string into individual rows?**
Use `STRING_SPLIT` (SQL Server), `UNNEST(STRING_TO_ARRAY(col, ','))` (Postgres), or recursive CTE.

**Q105: How do you write an IF...THEN logic in SQL?**
Use `CASE WHEN condition THEN result ELSE other END`.

**Q106: How do you simulate full outer join in MySQL (which doesn't support it natively)?**
`SELECT * FROM A LEFT JOIN B ... UNION SELECT * FROM A RIGHT JOIN B ...`.

**Q107: How do you get the Nth row in a table without using `LIMIT`?**
Window functions: `WHERE row_num = N`. Old way: Top N minus Top N-1.

**Q108: Write a query to remove duplicates but keep the most recent record.**
`DELETE FROM T WHERE id IN (SELECT id FROM (SELECT id, ROW_NUMBER() OVER(PARTTION BY key ORDER BY date DESC) as rn FROM T) WHERE rn > 1)`.

**Q109: How do you apply a calculation conditionally across rows using window functions?**
`SUM(CASE WHEN condition THEN val ELSE 0 END) OVER (...)`.

**Q110: How do you deal with circular references in data joins?**
Use Recursive CTEs with a specialized "visited" array check to stop cycles, or fix schema hierarchy.

---

## 🔹 12. User Behavior & Funnel Analysis (Questions 111-120)

**Q111: Identify users who started but didn’t complete a purchase.**
`SELECT user_id FROM Cart WHERE user_id NOT IN (SELECT user_id FROM Orders)`.

**Q112: Calculate conversion rate from page view to checkout.**
`(COUNT(DISTINCT checkout_users) * 1.0 / COUNT(DISTINCT view_users)) * 100`.

**Q113: How would you construct a funnel in SQL?**
Aggregated counts of Users at step 1, step 2, step 3. Often using `LEFT JOINS` on subsequent steps or `CASE` aggregation.

**Q114: Identify drop-off points in a multi-step signup process.**
Calculate count of users at each step. Substract Step N+1 from Step N to find drop-off.

**Q115: How would you calculate average steps in a user journey?**
` AVG(MAX(step_number)) GROUP BY session_id` if steps are numbered.

**Q116: Find users who returned within 7 days of first visit.**
`JOIN` visits V1 and V2 on user_id where `V2.date BETWEEN V1.date AND V1.date + 7`.

**Q117: How would you create a cohort analysis using SQL?**
Group by `DATE_TRUNC('month', signup_date)` (Cohort) and `DATEDIFF(event_date, signup_date)` (Age). Count active users.

**Q118: Build a time-to-convert distribution for marketing leads.**
Calculate `DATEDIFF` between Lead date and Sale date. Group by "Days to convert" buckets.

**Q119: Write a query to calculate activation rate per signup source.**
`(Activated Users / Total Signups)` Group By `Source`.

**Q120: Identify users who reactivated after being inactive for over 30 days.**
Compare distinct activity dates. User active today where `LAG(date) < today - 30`.

---

## 🔹 13. Complex Joins & Relations (Questions 121-130)

**Q121: Join three tables where two of them have a many-to-many relationship.**
Table A Join JunctionTable AB Join Table B.

**Q122: How would you handle joining a large fact table with multiple dimension tables?**
Star Schema joins. Ensure FK columns are indexed. Filter Fact table first if possible.

**Q123: How do you resolve data duplication caused by joins?**
Check for 1:N relationships exploding rows. Use `DISTINCT` or `GROUP BY` after join. Or aggregate the Many side *before* joining.

**Q124: Write a query to compare two tables and list differences.**
`SELECT * FROM A EXCEPT SELECT * FROM B`. (Postgres/SQL Server). Or `MINUS`.

**Q125: How do you find the latest record for each group in a joined result?**
join with `AND table.date = (SELECT MAX(date) ...)` logic or Window Function on the joined set.

**Q126: How would you join tables with composite keys?**
`ON A.key1 = B.key1 AND A.key2 = B.key2`.

**Q127: How do you join a table with itself to calculate differences between rows?**
`FROM Table T1 JOIN Table T2 ON T1.id = T2.id - 1` (Compare current row with previous row).

**Q128: How would you join on date ranges instead of exact match?**
`ON A.date BETWEEN B.start_date AND B.end_date`. (Range Join, beware performance).

**Q129: Write a query to show changes over time across related entities.**
Join Snapshots of entities on Time+ID. `SELECT time, entity, val - LAG(val) ...`

**Q130: What are anti-joins and how do you implement them?**
Join to find non-matches. `LEFT JOIN ... WHERE right.id IS NULL` or `NOT EXISTS`.

---

## 🔹 14. Audit & Data Change Tracking (Questions 131-140)

**Q131: Track when a record was inserted, updated, or deleted.**
Add columns `created_at`, `updated_at`. For deletions, use Soft Delete (`deleted_at`) or Audit Log table.

**Q132: How would you implement version control in SQL for audit logs?**
Change Data Capture (CDC). Triggers writing Old/New values to History table.

**Q133: Write a query to find changes in user address over time.**
Query Audit/History table filtering by `field = 'address'`.

**Q134: How do you detect deleted records between two data snapshots?**
`SELECT id FROM SnapshotOld WHERE id NOT IN (SELECT id FROM SnapshotNew)`.

**Q135: Implement Slowly Changing Dimension Type 1 vs Type 2 in SQL.**
Type 1: `UPDATE` (Overwrite). Type 2: `INSERT` new row with current date, close old row with end date.

**Q136: How would you identify hard deletes in a soft delete system?**
Compare the row count in backup vs live. Soft deleted rows exist but flag=true. Hard deleted are gone.

**Q137: Detect price changes in products historically.**
Window function: `lag_price = LAG(price)`. Filter `WHERE price != lag_price`.

**Q138: Write a query to find duplicate IDs with different values over time.**
`SELECT id, COUNT(DISTINCT value) FROM events GROUP BY id HAVING COUNT > 1`.

**Q139: Build a changelog summary per user.**
`SELECT user, action, date FROM AuditLog ORDER BY date`.

**Q140: Detect reinsertions after deletions using timestamp comparisons.**
Check if a user has an `INSERT` event *after* a `DELETE` event time.

---

## 🔹 15. ETL, Data Pipelines & Automation (Questions 141-150)

**Q141: How would you incrementally load new data using SQL?**
`SELECT * FROM Source WHERE updated_at > (SELECT MAX(updated_at) FROM Target)`.

**Q142: What are common issues with ETL scripts in SQL?**
Data type mismatches, NULL handling breaking logic, Running out of Temp DB space, Long locks.

**Q143: How would you deduplicate incoming records using SQL?**
Use `ROW_NUMBER() ... PARTITION BY id`. Insert only where RN=1.

**Q144: Explain idempotency in SQL ETL workflows.**
Running the script twice yields the same result. Use `MERGE` (Upsert) or Delete-Then-Insert for partition.

**Q145: Write a query to apply transformations before loading into a dimension table.**
`INSERT INTO Dim SELECT UPPER(name), COALESCE(num, 0) ... FROM Stage`.

**Q146: How do you validate an ETL result using only SQL?**
Compare `COUNT(*)` and `SUM(metric)` between Source and Target.

**Q147: Compare two datasets (source vs target) post ETL and report mismatches.**
`SELECT id, hash(row) FROM Src EXCEPT SELECT id, hash(row) FROM Tgt`.

**Q148: Build a surrogate key using SQL.**
`ROW_NUMBER() OVER(ORDER BY pk) + BaseOffset`. Or Identity column.

**Q149: How would you stage large raw files using SQL?**
`COPY` command to Staging Table (No constraints/indexes). Then process to Final Table.

**Q150: How do you implement late arriving data in SQL ETL?**
Insert a placeholder "Unknown" record in Dimension. Or insert Fact with negative/dummy FK then fix later.

---

## 🔹 16. SQL for Reporting & BI Dashboards (Questions 151-160)

**Q151: Build a daily active user (DAU) dashboard using SQL.**
`SELECT date, COUNT(DISTINCT user_id) as DAU FROM events GROUP BY date`.

**Q152: How do you prepare data for a monthly revenue report?**
`SELECT DATE_TRUNC('month', date), SUM(amount) FROM sales GROUP BY 1`.

**Q153: Generate a time series report with missing dates filled in.**
Generate a date sequence (Calendar table/Series) `LEFT JOIN` actual data. `COALESCE(val, 0)`.

**Q154: How do you format monetary values in SQL reports?**
`FORMAT(val, 'C')` (SQL Server) or `TO_CHAR(val, 'L999G99')` (Postgres). Note: Formatting usually best done in BI Tool.

**Q155: Build a year-over-year comparison of revenue.**
`SUM(revenue)` vs `LAG(SUM(revenue), 12) OVER (ORDER BY month)`.

**Q156: Write a KPI dashboard query for average resolution time of tickets.**
`AVG(resolved_at - created_at)`.

**Q157: Create a report showing top 10 vs bottom 10 performers.**
`(SELECT * ... ORDER BY val DESC LIMIT 10) UNION ALL (SELECT * ... ORDER BY val ASC LIMIT 10)`.

**Q158: Design a query for drill-down reporting by region → product → day.**
Use `GROUP BY ROLLUP(region, product, day)`.

**Q159: Build a dynamic filter for reporting using SQL variables.**
`WHERE (@region IS NULL OR region = @region)`.

**Q160: Write a SQL snippet that prepares data for pie chart visualizations.**
`SELECT category, value, (value / SUM(value) OVER()) * 100 as pct`.

---

## 🔹 17. Error Handling, Edge Cases & Data Integrity (Questions 161-170)

**Q161: How would you handle outliers in SQL?**
Filter using Z-Score (Mean +/- 3*StdDev) or Interquartile Range (IQR).

**Q162: Detect NULLs in critical columns and count them.**
`SELECT COUNT(*) FROM T WHERE crit_col IS NULL`.

**Q163: What would you do if you encounter unexpected duplicate primary keys?**
Identify causes (Bad join? Retry logic?). De-dup and Add Unique Constraint to prevent recurrence.

**Q164: Validate that all foreign keys are valid.**
`SELECT * FROM Child C LEFT JOIN Parent P ON C.pid = P.id WHERE P.id IS NULL`. (Orphan Check).

**Q165: Write a check that fails if more than 10% of data is missing.**
`SELECT CASE WHEN (cnt_null * 1.0 / total) > 0.1 THEN 'FAIL' ELSE 'OK' END ...`.

**Q166: How do you build a “data completeness” score for a dataset?**
Avg of filled columns per row. `(CASE WHEN c1 IS NOT NULL + ... ) / NumCols`.

**Q167: Write a check to ensure sales data always has a date.**
`CONSTRAINT chk_date CHECK (sale_date IS NOT NULL)`.

**Q168: Flag future-dated transactions in production data.**
`WHERE txn_date > CURRENT_TIMESTAMP`.

**Q169: How do you check for skewed distributions in SQL?**
Compare Mean vs Median. Or group buckets and count.

**Q170: Write a query to ensure no orphan records exist in child tables.**
Same as Q164.

---

## 🔹 18. Security & Access Control (Questions 171-180)

**Q171: How do you prevent unauthorized access using SQL roles?**
Create Role `ReadOnly`. Grant Select only. Assign Users to Role.

**Q172: What’s the difference between GRANT and REVOKE in SQL?**
`GRANT`: Agrees permission. `REVOKE`: Removes permission.

**Q173: Can you mask sensitive data (e.g., PII) in SQL?**
Dynamic Data Masking (SQL Server/Snowflake) `MASKED WITH (FUNCTION = 'default()')`. Or Hash it.

**Q174: How would you implement row-level security?**
Policies (Postgres RLS). `CREATE POLICY p ON T USING (user_id = current_user_id)`.

**Q175: What steps would you take to log and audit access to sensitive tables?**
Enable DB Audit Logs for SELECTs on that table.

**Q176: How do you redact customer names in query results?**
`SELECT LEFT(name, 1) + '***' FROM users`.

**Q177: How would you separate user permissions between dev and prod databases?**
Devs have DDL/Write in Dev, Read-Only in Prod. Separate credentials.

**Q178: How do you safely query encrypted fields?**
Use DB functions to decrypt if you have key/cert permissions. `DECRYPTBYKEY(col)`.

**Q179: Create a view that only shows non-sensitive data to analysts.**
`CREATE VIEW PublicUsers AS SELECT id, country (exclude email/phone) FROM Users`.

**Q180: What are SQL injection attacks and how do you avoid them in dynamic SQL?**
Appending strings to query. Avoid by using `sp_executesql` with params or ORM bindings.

---

## 🔹 19. SQL with Tools & Systems (Questions 181-190)

**Q181: How does SQL syntax differ across platforms (e.g., MySQL vs PostgreSQL vs Redshift)?**
Quote chars (`"` vs `` ` ``), Date funcs (`GETDATE` vs `NOW`), JSON handling, proprietary extensions.

**Q182: How would you write SQL that works cross-database (Snowflake, BigQuery)?**
Stick to ANSI SQL standard. Avoid specialized functions (like `DATEADD` specific syntax, use standard `INTERVAL`).

**Q183: What are best practices for writing reusable SQL in dbt?**
Use CTEs, Macros for repeated logic, `ref()` for dependencies.

**Q184: How would you handle slowly changing dimensions in Looker?**
Modeling logic in lookml view, filtering on `valid_to IS NULL`.

**Q185: What’s the role of SQL in Airflow DAGs?**
`PostgresOperator` or `BigQueryOperator` executes SQL transformation steps in pipeline.

**Q186: How do you integrate SQL queries into Power BI or Tableau?**
Direct Query or Import Mode. Use Views for cleaner integration.

**Q187: How would you write test cases for SQL models in dbt?**
Schema tests (`unique`, `not_null`), Custom SQL tests (`SELECT * FROM model WHERE invalid_condition`).

**Q188: How would you write a modular SQL model in Dataform or dbt?**
Break logic into separate files/models. Stage -> Intermediate -> Mart.

**Q189: How do you pass parameters into SQL in BI tools?**
Use Parameters features (Tableau params, Power BI Query Params).

**Q190: How would you optimize queries in a cloud warehouse (BigQuery, Snowflake)?**
Cluster/Partition keys. Avoid `SELECT *`. Filter early. Use Approx functions (`COUNT(DISTINCT)` approx).

---

## 🔹 20. Scenario-Based Business Questions (Questions 191-200)

**Q191: Find the top 3 reasons customers churn.**
`SELECT reason, COUNT(*) FROM churn_survey GROUP BY reason ORDER BY 2 DESC LIMIT 3`.

**Q192: Compare pre- and post-campaign conversion rates.**
`AVG(CASE WHEN date < launch THEN conversion END)` vs `AVG(CASE WHEN date >= launch ...)`

**Q193: Analyze sales trends during promotional periods.**
Filter sales inside `promo_start` and `ends`. Compare daily avg to non-promo periods.

**Q194: Identify high-value vs low-value customers.**
RFM Analysis (Recency, Frequency, Monetary). `NTILE(4) OVER (ORDER BY revenue)` to quartile.

**Q195: Write SQL to find product cannibalization (drop in product A after launch of product B).**
Monitor Sales(A) trend when Sales(B) starts > 0.

**Q196: Determine correlation between support ticket volume and customer churn.**
Join Churn table with Ticket counts per user. Check avg tickets for chuerned vs retained.

**Q197: Detect seasonal trends in product sales.**
Group by Month/Quarter across multiple years.

**Q198: Find which marketing channel has highest ROI using only SQL.**
`(Revenue - Cost) / Cost` Group By Channel.

**Q199: Analyze customer feedback patterns using SQL text functions.**
`CASE WHEN feedback LIKE '%slow%' THEN 'Performance' ...`.

**Q200: Predict potential out-of-stock inventory situations with lead time calculations.**
`current_stock - (daily_sales_rate * lead_time_days) < safety_stock`.

---

## 🔹 21. Analytical Thinking with SQL (Questions 201-210)

**Q201: Identify users who make purchases only on weekends.**
`SELECT user_id FROM orders GROUP BY user_id HAVING COUNT(CASE WHEN DAYOFWEEK(date) NOT IN (1,7) THEN 1 END) = 0`.

**Q202: Write SQL to calculate the average session duration per user.**
`SELECT user_id, AVG(logout_time - login_time) FROM sessions GROUP BY user_id`.

**Q203: Segment users by frequency of interaction (daily, weekly, monthly).**
Count distinct days active. If > 20 'Daily', > 4 'Weekly', else 'Monthly'.

**Q204: How would you identify peak transaction hours per region?**
`SELECT region, HOUR(time), COUNT(*) FROM txns GROUP BY 1, 2 ORDER BY 3 DESC`.

**Q205: Write a query to analyze churn trend by subscription duration.**
Group by `DATEDIFF(end_date, start_date)` buckets and calculate churn rate.

**Q206: Calculate moving averages over a 7-day window.**
`AVG(sales) OVER (ORDER BY date ROWS BETWEEN 6 PRECEDING AND CURRENT ROW)`.

**Q207: Find correlation between product price and return rate using SQL.**
Use `CORR(price, is_returned)` if supported, or manual covariance calculation.

**Q208: Write SQL to calculate win/loss ratio by sales rep.**
`SUM(CASE WHEN status='Won' THEN 1 ELSE 0 END) * 1.0 / SUM(CASE WHEN status='Lost' THEN 1 ELSE 0 END)`.

**Q209: How would you analyze failed login attempts per IP?**
`SELECT ip, COUNT(*) FROM logs WHERE status = 'Failed' GROUP BY ip HAVING COUNT(*) > Threshold`.

**Q210: Track average time between first and last touchpoint in a funnel.**
`AVG(MAX(event_time) - MIN(event_time)) GROUP BY user_id`.

---

## 🔹 22. Complex Data Types and Advanced Structures (Questions 211-220)

**Q211: How do you query nested JSON in a column?**
`col->'key'->>'subKey'` (Postgres) or `JSON_VALUE(col, '$.key.subKey')` (SQL Server/MySQL).

**Q212: Extract all keys from a JSON object stored in SQL.**
`JSON_KEYS(col)` (MySQL) or `json_object_keys(col)` (Postgres).

**Q213: Flatten an array column into individual rows.**
`UNNEST(array_col)` (Postgres) or `CROSS APPLY OPENJSON(col)` (SQL Server).

**Q214: Convert key-value pairs stored as string into a table.**
String Split then separate by Separator. Or Parse JSON.

**Q215: Write SQL to parse a CSV string inside a column.**
`STRING_SPLIT(col, ',')`.

**Q216: Handle JSON logs with variable schemas in SQL.**
Store in `JSONB` column. Query using `COALESCE(col->>'fieldA', col->>'fieldB')`.

**Q217: Filter records based on values in a JSON array.**
`WHERE 'value' = ANY(SELECT value FROM json_array_elements_text(col))`.

**Q218: Write a query that returns all columns as a single JSON object.**
`SELECT row_to_json(t) FROM table t` (Postgres).

**Q219: Convert JSON column into multiple flat columns.**
`SELECT col->>'name' as name, col->>'age' as age FROM table`.

**Q220: Validate if a JSON field has required keys.**
`col ?& array['key1', 'key2']` (Postgres - contains all keys).

---

## 🔹 23. E-Commerce SQL Problems (Questions 221-230)

**Q221: Identify abandoned cart users in the past 30 days.**
Users in Cart table but not in Orders table (Left Join/Not In) `WHERE date > NOW() - 30`.

**Q222: Write SQL to calculate repeat purchase rate by customer.**
`COUNT(CASE WHEN order_count > 1 THEN 1 END) / COUNT(*)`.

**Q223: Segment customers into first-time vs returning buyers.**
Count Orders per user. If 1 -> First-time, >1 -> Returning.

**Q224: Analyze conversion rate by traffic source and device.**
`COUNT(orders) / COUNT(visits)` Group By Source, Device.

**Q225: Track coupon usage and redemption rates.**
`(Redeemed Count / Issued Count)` Group By Coupon_ID.

**Q226: Find products frequently bought together (market basket analysis).**
Self join OrderLines on OrderID. `SELECT A.prod, B.prod, COUNT(*) ... WHERE A.prod < B.prod`.

**Q227: Calculate average revenue per user (ARPU).**
`SUM(revenue) / COUNT(DISTINCT user_id)`.

**Q228: Identify dormant customers who haven’t purchased in 6 months.**
`MAX(order_date) < DATE_SUB(NOW(), INTERVAL 6 MONTH)`.

**Q229: Analyze order-to-fulfillment time per region.**
`AVG(fulfilled_at - ordered_at)` Group By Region.

**Q230: Track inventory turnover by product category.**
`(COGS / Avg Inventory Value)` Group By Category.

---

## 🔹 24. Fintech / Transactions (Questions 231-240)

**Q231: Detect duplicate financial transactions within a short time window.**
Self join T1, T2 `ON t1.user=t2.user AND ABS(t1.time - t2.time) < 1 min AND t1.amt = t2.amt`.

**Q232: Calculate daily interest accrual for customer accounts.**
`balance * (interest_rate / 365)`. Aggregate daily.

**Q233: Write SQL to find customers with multiple failed payments.**
`COUNT(*) > 1 WHERE status='Failed' GROUP BY user_id`.

**Q234: Find high-value transactions flagged for manual review.**
`SELECT * FROM txns WHERE amount > 10000`.

**Q235: Segment transactions by risk category.**
`CASE WHEN amount > 10000 OR country = 'HighRisk' THEN 'High' ...`.

**Q236: Track account balance over time per customer.**
Running sum of transactions: `SUM(amount) OVER (PARTITION BY user ORDER BY date)`.

**Q237: Analyze deposit-to-withdrawal ratio per user.**
`SUM(CASE WHEN type='Dep' THEN amt END) / SUM(CASE WHEN type='Wd' THEN amt END)`.

**Q238: Find customers who cross KYC threshold but are not verified.**
`HAVING SUM(amt) > 10000 AND verification_status = 'Pending'`.

**Q239: Write a reconciliation report between two transaction tables.**
`FULL OUTER JOIN` Internal vs BankLeder `ON id`. Show mismatches.

**Q240: Calculate fraud detection score using SQL logic.**
Weighted sum of flags (Multi-IP + High Amt + Night Time).

---

## 🔹 25. IoT / Sensor Data Use Cases (Questions 241-250)

**Q241: Find devices that haven’t reported in the last 24 hours.**
`MAX(ping_time) < NOW() - INTERVAL 24 HOUR`.

**Q242: Calculate average temperature by room and hour.**
`AVG(temp)` Group By Room, `DATE_TRUNC('hour', time)`.

**Q243: Detect out-of-range sensor readings and alert conditions.**
`WHERE value NOT BETWEEN min_safe AND max_safe`.

**Q244: Calculate uptime percentage of devices over the last month.**
`(Count of active pings / Total expected pings) * 100`.

**Q245: Identify sensors with inconsistent reporting intervals.**
Calculate `EXTRACT(EPOCH FROM time - LAG(time))`. Variance > Threshold.

**Q246: Analyze event frequency per machine.**
`COUNT(*)` per machine per day.

**Q247: Write SQL to find overlapping signal timestamps.**
`WHERE T1.start < T2.end AND T1.end > T2.start`.

**Q248: Aggregate sensor data into 10-minute time buckets.**
`GROUP BY FLOOR(MINUTE(time) / 10)`.

**Q249: Correlate device type with error frequency.**
`(Error Count / Total Events)` Group By DeviceType.

**Q250: Rank devices by reliability score based on downtime.**
`RANK() OVER (ORDER BY downtime_minutes ASC)`.

---

## 🔹 26. Healthcare / Medical Use Cases (Questions 251-260)

**Q251: Identify patients with recurring diagnoses in last 6 months.**
`COUNT(*) > 1 WHERE diagnosis='X' AND date > NOW() - 6M`.

**Q252: Calculate average hospital stay by procedure type.**
`AVG(discharge_date - admission_date)` Group By Procedure.

**Q253: Find patients who missed follow-up appointments.**
`appointment_status = 'No Show'`.

**Q254: Write SQL to track medication adherence.**
Ratio of `Refilled Prescriptions` vs `Expected Refills`.

**Q255: Compare pre- and post-treatment metrics using window functions.**
`AVG(metric) FILTER(WHERE date < tx_date)` vs `AVG ... > tx_date`.

**Q256: Segment patients by age group and condition.**
`GROUP BY CASE WHEN age...`, Condition.

**Q257: Identify facilities with highest readmission rates.**
Count patients readmitted within 30 days of discharge / Total discharges.

**Q258: Detect abnormal lab result values.**
`WHERE result > ReferenceMax OR result < ReferenceMin`.

**Q259: Build a patient journey timeline using SQL.**
`SELECT date, event_type, details FROM Events WHERE patient_id=... ORDER BY date`.

**Q260: Aggregate diagnosis counts by ICD code and quarter.**
`GROUP BY icd_code, DATE_PART('quarter', date)`.

---

## 🔹 27. Time Series & Temporal SQL (Questions 261-270)

**Q261: Create a gap and island analysis of user activity.**
Identify consecutive sequences (islands) and missing sequences (gaps) in dates. (See "Gaps and Islands" problem).

**Q262: Detect periods of inactivity longer than 3 days.**
`DATEDIFF(date, LAG(date)) > 3`.

**Q263: Calculate duration of continuous login streaks.**
Assign GroupID to consecutive days. `COUNT(*)` per GroupID.

**Q264: How would you fill in missing daily values with zeros?**
Right Join valid data with a Calendar table (Generate Series). `COALESCE(val, 0)`.

**Q265: Write SQL to interpolate missing values linearly.**
Complex. Calculate Slope between last known and next known value. Apply slope to missing timestamps.

**Q266: Detect backward timestamp entries (data quality issue).**
`WHERE date < LAG(date) OVER (ORDER BY id)`.

**Q267: Create time buckets that shift dynamically by timezone.**
`GROUP BY DATE_TRUNC('hour', time AT TIME ZONE user_tz)`.

**Q268: Perform day-over-day and week-over-week comparisons.**
`val - LAG(val, 1)` or `val - LAG(val, 7)`.

**Q269: Calculate delta between current and previous value per device.**
`val - LAG(val) OVER (PARTITION BY device ORDER BY time)`.

**Q270: Track time spent in each state (e.g., online, offline, idle).**
`next_state_time - current_state_time` Group By State.

---

## 🔹 28. Multi-Tenant / SaaS Use Cases (Questions 271-280)

**Q271: Filter results to only show data belonging to a specific tenant ID.**
`WHERE tenant_id = current_setting('app.tenant_id')` or param.

**Q272: Write SQL to calculate average usage per customer account.**
`SUM(usage) / COUNT(users)` Group By Tenant.

**Q273: Compare active users across customers on same subscription plan.**
`COUNT(active_users)` Group By Plan, Customer.

**Q274: Identify top feature usage per tenant.**
`RANK() OVER (PARTITION BY tenant ORDER BY usage DESC) = 1`.

**Q275: Detect unusual usage patterns across tenants.**
Compare Tenant Usage vs Avg Usage of all Tenants.

**Q276: Rank customers by feature adoption rate.**
`(Users using feature / Total Users)` Order By DESC.

**Q277: Analyze churn by subscription type and usage frequency.**
Group by Plan, UsageBucket. Calc Churn Rate.

**Q278: Track support ticket volume per customer.**
`COUNT(*)` Group By CustomerID.

**Q279: Build customer health score using SQL metrics.**
Weighted Avg of (Login Freq, Usage, Support Tickets).

**Q280: Identify customers who downgraded in the last 90 days.**
`Plan_New < Plan_Old` in Subscription Log table.

---

## 🔹 29. Security / Access / Logging (Questions 281-290)

**Q281: Identify IPs with the most failed login attempts.**
`SELECT ip, COUNT(*) FROM logs WHERE result='fail' ORDER BY 2 DESC LIMIT 10`.

**Q282: Track admin actions by user and timestamp.**
`SELECT * FROM audit_log WHERE role='admin' ORDER BY time`.

**Q283: Detect brute-force login behavior using SQL.**
Count failed logins per IP/User in last minute. If > 10, flag.

**Q284: Write SQL to determine role hierarchy across users.**
Recursive query on Roles table (Role -> Parent Role).

**Q285: Audit password reset requests over time.**
`COUNT(*)` Group By Day/Hour from ResetLog.

**Q286: Detect logins from multiple locations within short timeframe.**
Self Join Login L1, L2. `L1.user=L2.user AND L1.loc != L2.loc AND time_diff < travel_time`.

**Q287: Compare access rights between two users.**
`SELECT permission FROM UserPerms WHERE user='A' EXCEPT ... 'B'`.

**Q288: Track audit log events by severity level.**
`COUNT(*)` Group By Severity.

**Q289: Identify users who changed permissions in last 7 days.**
Log event type 'PermissionChange' in last 7 days.

**Q290: Calculate average session duration per role.**
`AVG(duration)` Group By UserRole.

---

## 🔹 30. ETL Testing & Data Quality Assurance (Questions 291-300)

**Q291: Compare record counts across staging and production tables.**
`SELECT (SELECT COUNT(*) FROM stg) - (SELECT COUNT(*) FROM prod)`.

**Q292: Write SQL to detect schema drift in imported files.**
Query `INFORMATION_SCHEMA.COLUMNS` to comparing col counts/names vs expected.

**Q293: Track failed loads in ETL logs using SQL.**
`SELECT * FROM etl_logs WHERE status='Error'`.

**Q294: Validate primary key uniqueness in raw ingested data.**
`COUNT(pk) > COUNT(DISTINCT pk)`.

**Q295: Write a row-level comparison between two versions of a table.**
`Hash(Col1, Col2...)`. Join on ID. Compare Hash.

**Q296: How do you detect schema changes in SQL pipelines?**
Check if `SELECT *` fails or check metadata tables for column additions.

**Q297: Implement row hashing to compare large datasets efficiently.**
`MD5(CONCAT(c1, c2...))` or `HASHBYTES`.

**Q298: Monitor freshness of ETL data using load timestamps.**
`MAX(load_time) < NOW() - INTERVAL 1 HOUR`. Alert if true.

**Q299: Track changes in business metrics after ETL change.**
Run metric calc on old vs new data version. Diff should be 0.

**Q300: Identify partial or incomplete loads in partitioned tables.**
Count rows per partition. If dramatically lower than avg, flag.

---

## 🔹 31. Tracking User Events & Analytics (Questions 301-310)

**Q301: How would you track a user’s complete clickstream path?**
`STRING_AGG(page_url, ' -> ') WITHIN GROUP (ORDER BY timestamp)` per session.

**Q302: Write a query to calculate time between specific events (e.g. add_to_cart → purchase).**
Join Events E1 (Cart) and E2 (Purchase) on UserID `WHERE E2.time > E1.time`. `AVG(E2.time - E1.time)`.

**Q303: Identify the most common sequence of three actions by users.**
`RANK() OVER (ORDER BY Count DESC)` of `CONCAT(e1.action, '>', e2.action, '>', e3.action)`.

**Q304: Track first touch and last touch attribution for a user.**
`FIRST_VALUE(source) OVER (PARTITION BY user ORDER BY time)` and `LAST_VALUE`.

**Q305: Write SQL to calculate session length using event timestamps.**
`MAX(timestamp) - MIN(timestamp)` Group By SessionID.

**Q306: How would you identify bot-like behavior in click data?**
High frequency: `COUNT(*) > 100` events in 1 minute. Or distinct UserAgent check.

**Q307: Filter event logs where users clicked the same button repeatedly within 10 seconds.**
`LEAD(action) = action AND LEAD(time) - time < 10 sec`.

**Q308: Count the number of times a user performed an action after seeing a banner.**
`COUNT(*)` where `action_time > banner_view_time` inside session.

**Q309: Find users who performed a sequence A → B → C in order.**
Use `MATCH_RECOGNIZE` (Oracle/Snowflake) or self-joins `A.time < B.time < C.time`.

**Q310: Compare engagement metrics between users who saw an experiment and those who didn’t.**
`AVG(clicks)` Group By `ExperimentGroup` (Control vs Test).

---

## 🔹 32. Business Logic & Data Contracts (Questions 311-320)

**Q311: Write a constraint to enforce that end_date must be after start_date.**
`ALTER TABLE T ADD CONSTRAINT chk_dates CHECK (end_date > start_date)`.

**Q312: Create a check that ensures salary is non-negative and within company limits.**
`CHECK (salary >= 0 AND salary <= 500000)`.

**Q313: How do you enforce allowed values in a string column?**
`CHECK (status IN ('Active', 'Inactive', 'Pending'))` or use Enum type.

**Q314: Detect records violating a uniqueness constraint across multiple columns.**
`GROUP BY col1, col2 HAVING COUNT(*) > 1`.

**Q315: Write SQL to identify schema mismatches in dynamic ingestion tables.**
Compare column names/types from Information Schema vs Contracts table.

**Q316: Enforce business rule: users can’t be in more than one active trial at a time.**
Trigger or Partial Index: `CREATE UNIQUE INDEX ON Trials(user_id) WHERE status='Active'`.

**Q317: Identify foreign keys pointing to non-existent parent records.**
`SELECT * FROM Child WHERE pid NOT IN (SELECT id FROM Parent)`.

**Q318: Write a query to detect duplicated timestamps for the same user.**
`COUNT(timestamp) > 1` per user.

**Q319: Audit table where deleted rows must leave an audit trail.**
Use `AFTER DELETE` Trigger to insert row into `AuditTable`.

**Q320: Validate that all customer emails follow proper syntax using regex or LIKE.**
`WHERE email NOT LIKE '%_@__%.__%'`.

---

## 🔹 33. Systems Thinking with SQL (Questions 321-330)

**Q321: How do you model many-to-many relationships in SQL schema?**
Use a Junction Table (Associative Entity) containing FKs from both tables (`student_id`, `course_id`).

**Q322: What are surrogate keys and why would you use them?**
Artificial PK (Integers/UUIDs). Decouples DB design from business data (Email/SSN) which might change.

**Q323: Model a feature flag system using relational tables.**
Tables: `Features`, `Users`, `UserFeatures` (overrides), `GlobalFeatures` (defaults).

**Q324: Design a product subscription model with upgrades and downgrades.**
Table `Subscriptions` with `plan_id`, `start_date`, `end_date`. Upgrade ends current row, starts new one.

**Q325: Write SQL to model and report on customer referral trees.**
Self-referencing table `User(id, referrer_id)`. Use Recursive CTE to calculate depth/downline.

**Q326: Implement tag-based searching across products in SQL.**
`Products`, `Tags`, `ProductTags`. Query: `JOIN ProductTags WHERE tag IN (...)`.

**Q327: Write a SQL schema to support multivariate A/B testing.**
`Experiments`, `Variants`, `UserAssignments`.

**Q328: How would you implement data versioning using SQL?**
Add `version` int column to PK. Or `valid_from`, `valid_to` columns.

**Q329: Model and report on hierarchical department structures.**
Recursive CTE on `Dept(id, parent_id)`. `Path` generation: `Parent/Child`.

**Q330: Create a changelog table that captures inserts, updates, and deletes in detail.**
Cols: `TableName`, `RowID`, `Operation`, `OldValue`, `NewValue`, `User`, `Time`.

---

## 🔹 34. Anti-patterns & Tricky SQL Logic (Questions 331-340)

**Q331: What’s wrong with using SELECT * in production?**
Breaks app if schema changes. Fetches unnecessary data (Network/IO cost). Prevents Index-Only scans.

**Q332: How does a GROUP BY without aggregation behave?**
Like `DISTINCT`. Returns unique combinations of grouped columns.

**Q333: Why is DISTINCT often misused, and how would you fix it?**
Used to hide duplicate join issues. Fix the join logic instead of masking it with expensive `DISTINCT`.

**Q334: Explain a case where a LEFT JOIN becomes an INNER JOIN due to filters.**
`SELECT * FROM A LEFT JOIN B ON ... WHERE B.col = 'val'`. The WHERE clause filters out NULLs from non-matches.

**Q335: What are the risks of filtering inside a JOIN condition?**
In Inner Join: No diff. In Left Join: `ON ... AND B.col='val'` filters B *before* joining (preserves A rows). `WHERE` filters after.

**Q336: Write an example where window function is misused causing incorrect logic.**
Filtering on result of window function in same level. Must use subquery/CTE: `SELECT * FROM (SELECT RANK()...) WHERE rk=1`.

**Q337: What problems arise from non-deterministic ordering without ORDER BY?**
Result order is undefined. Pagination (`LIMIT/OFFSET`) becomes unreliable/random.

**Q338: Detect use of inconsistent aggregation across metrics.**
Mixing granularity (e.g. Summing Daily Users to get Monthly Users - wrong, should count distinct).

**Q339: Find cases where string-to-number implicit casting causes wrong outputs.**
`'10' > '2'` is False (String comparison). Ensure explicit `CAST`.

**Q340: Describe how NULL logic can break equality checks in joins.**
`NULL = NULL` is Unknown (False). Joins fail on NULL keys. Use `COALESCE` or `IS NOT DISTINCT FROM`.

---

## 🔹 35. Data Governance & Lineage (Questions 341-350)

**Q341: How do you track downstream dependencies of a SQL view?**
Query system dependency tables (`information_schema.view_table_usage` or `pg_depend`).

**Q342: Write SQL to identify columns with PII or sensitive data.**
Query `COLUMNS` table looking for names like `%ssn%`, `%email%`, `%password%`.

**Q343: How do you flag stale or unused tables in a warehouse?**
Check `Last_Accessed` or `Query_History` logs. If 0 reads in 90 days -> Stale.

**Q344: Find all views that reference a specific table or column.**
Search in `definition` text in `information_schema.views`.

**Q345: Write a lineage report showing upstream sources of a KPI metric.**
Recursively query view definitions. Hard in pure SQL, better with tools (dbt docs/Amundsen).

**Q346: Identify duplicate datasets created by different teams.**
Fuzzy match on Table Names or Column schemas.

**Q347: Enforce column naming conventions using metadata queries.**
`SELECT column_name FROM cols WHERE column_name NOT LIKE '%_snake_case%'`.

**Q348: Identify hard-coded logic in SQL scripts (e.g., fixed date ranges).**
Search script text for regex `\d{4}-\d{2}-\d{2}`.

**Q349: Track last query time for each table to flag unused ones.**
Analyze query logs metadata.

**Q350: Generate a data catalog with column descriptions from information schema.**
`SELECT table_name, column_name, comment FROM information_schema.columns`.

---

## 🔹 36. Performance Diagnostics (Questions 351-360)

**Q351: How would you detect full table scans in your warehouse?**
Check Query Plan for `Seq Scan` (Postgres) or `Table Scan`. Filter audit logs for high `rows_scanned`.

**Q352: Compare performance between indexed vs non-indexed queries.**
Measure Execution Time / Cost. Indexed: O(log N). Scan: O(N).

**Q353: What are the downsides of excessive partitioning?**
Metadata overhead. Planner takes longer to prune partitions. Small files (in Hive/Spark).

**Q354: Write SQL to find top N largest tables in storage.**
Query `pg_relation_size` (Postgres) or `sp_spaceused` (SQL Server).

**Q355: Detect queries with high CPU usage in past week.**
Order `QueryHistory` by `cpu_time` desc.

**Q356: Monitor query latency percentiles using a system table.**
`approx_percentile(duration, 0.95)` from query logs.

**Q357: Audit long-running queries by user or service.**
`WHERE duration_seconds > 60`.

**Q358: Track sudden growth in row counts over time.**
Log `COUNT(*)` daily. Alert if `(Today - Yesterday) / Yesterday > 20%`.

**Q359: Profile skew in joins by inspecting distribution of join keys.**
`SELECT key, COUNT(*) FROM T GROUP BY key ORDER BY 2 DESC LIMIT 10`. High count = Skew.

**Q360: Identify and reduce cross joins that are unintentional.**
Check for queries with multiple tables in `FROM` but missing `ON/WHERE` join conditions.

---

## 🔹 37. Time-Aware Modeling (Questions 361-370)

**Q361: Create a slowly changing dimension model with effective/expiry dates.**
Cols: `Id`, `Value`, `ValidFrom`, `ValidTo`. To query current: `WHERE ValidTo IS NULL`.

**Q362: Track changing primary addresses while keeping history.**
Address Table with `IsActive` flag or Dates.

**Q363: Write SQL to identify overlapping active date ranges for users.**
`WHERE A.Start < B.End AND A.End > B.Start`.

**Q364: Model “as of” reporting logic to look back as of any historical date.**
`WHERE ValidFrom <= @ReportDate AND (ValidTo > @ReportDate OR ValidTo IS NULL)`.

**Q365: Find the state of a subscription table at the end of every month.**
Cross Join Key Dates with Subscriptions. Filter by validity range.

**Q366: Rebuild a snapshot of data using audit logs.**
Start Base + Replay Insert/Update/Deletes in order.

**Q367: How would you handle schema drift across time in SQL?**
Use `JSON` columns for flexible fields or separate tables for versioned schemas.

**Q368: Detect partial overlaps in time periods between two records.**
See Q363.

**Q369: Write a query to compute the latest value as of a given timestamp.**
`SELECT * FROM T WHERE time <= @TargetTime ORDER BY time DESC LIMIT 1`.

**Q370: Model time-aware KPI trends with point-in-time accuracy.**
Snapshot tables daily. Query specific snapshot date.

---

## 🔹 38. Cross-Domain Logic (Questions 371-380)

**Q371: Create a pricing tier structure for SaaS usage billing.**
Table `Tiers` (Min, Max, Price). Join `Usage` on `Units BETWEEN Min AND Max`.

**Q372: Model and track shipment status across time and carriers.**
`ShipmentStatusHistory` (ShipmentID, Status, Time, Location).

**Q373: Write SQL to generate invoices with dynamic pricing.**
Sum `ItemPrice * Quantity` + Tax - Discount.

**Q374: Design a loyalty point redemption and expiration logic in SQL.**
FIFO logic. Consume points from oldest "Earn" batch first.

**Q375: Identify drivers with idle time over X minutes between deliveries.**
`PickupTime(current) - DropoffTime(previous) > X`.

**Q376: Track active incidents in a ticketing system.**
`Status NOT IN ('Resolved', 'Closed')`.

**Q377: Build a schema to handle multi-currency transactions.**
Store `Amount` and `Currency`. Reference `ExchangeRates` table (Date, Currency, RateToUSD).

**Q378: Write SQL to calculate fuel consumption per route for logistics.**
`SUM(fuel_used)` Group By RouteID.

**Q379: Detect price anomalies in product catalogs by vendor.**
Avg Price per product. Alert if VendorPrice > 1.5 * Avg.

**Q380: Monitor service usage thresholds for automatic suspension.**
If `Usage > Limit`, update State='Suspended'.

---

## 🔹 39. Uncommon or Niche SQL Features (Questions 381-390)

**Q381: How do you use `FILTER(WHERE ...)` inside aggregates?**
`COUNT(*) FILTER (WHERE type='A')`. Standard SQL (Postgres supports). Cleaner than `CASE`.

**Q382: Use `QUALIFY` clause in Snowflake to filter windowed results.**
`SELECT * ... QUALIFY ROW_NUMBER() OVER(...) = 1`. Skips need for subquery.

**Q383: What does `SEMI JOIN` or `ANTI JOIN` mean in specific dialects?**
Semi: Returns rows from A where match in B (Exists). Anti: Returns A where NO match in B (Not Exists).

**Q384: Write a lateral join (or CROSS APPLY) use case.**
Calculating logic dependent on row values. `FROM Users U CROSS APPLY (SELECT Top 1 * FROM Orders O WHERE O.uid=U.id)`.

**Q385: Use `PIVOT` or `UNPIVOT` operations for reshaping data.**
`PIVOT (SUM(val) FOR col IN ([A],[B]))` (SQL Server).

**Q386: Implement set-returning functions in PostgreSQL (e.g., `unnest()`).**
`SELECT generate_series(1, 10)`.

**Q387: Generate dynamic SQL queries using stored procedures.**
Concat string logic `SET @sql = 'SELECT * FROM ' + @tableName`. `EXEC(@sql)`.

**Q388: Use `MERGE` for upsert logic in warehouse platforms.**
`MERGE INTO Tgt USING Src ON Match MATCH THEN Update NOT MATCH THEN Insert`.

**Q389: How do you query materialized views efficiently?**
Treat like table. Ensure it's refreshed. Index it.

**Q390: Build a recursive query to compute running balance in a ledger.**
Recursive CTE carrying forward `PreviousBalance + Credit - Debit`.

---

## 🔹 40. Mock Debugging / Fix-the-Bug Questions (Questions 391-400)

**Q391: This query returns wrong revenue totals — find and fix the error.**
Check for `Fan-out` (duplication) due to 1:N joins before summing. Agg first, then join.

**Q392: The result duplicates rows unexpectedly — what’s wrong?**
Missing join condition or Joining on non-unique columns.

**Q393: This funnel step count is higher than step 1 — diagnose it.**
Left Join order wrong? Or data quality (Users skipping steps).

**Q394: This rolling average is not behaving correctly — fix the logic.**
Check `ROWS BETWEEN` frame. Is it `UNBOUNDED PRECEDING` (Running Total) or `N PRECEDING` (Moving Avg)?

**Q395: NULLs are interfering with sort order — how would you fix?**
`ORDER BY col ASC NULLS LAST`.

**Q396: A window function isn't resetting per partition — why?**
Missing `PARTITION BY` clause. It's calculating over entire dataset.

**Q397: The grouping level is too granular — what caused it?**
Included unique ID or high-cardinality column in `GROUP BY`.

**Q398: Query uses subqueries but is very slow — suggest fixes.**
Rewrite `IN (SELECT...)` to `EXISTS` or `JOIN`.

**Q399: Results look correct but are outdated — identify the cause.**
Reading from a stale snapshot/Replica or Materialized View not refreshed.

**Q400: A daily report is showing fluctuating user counts — explain possible causes.**
Timezone issues (`UTC` vs `Local`). Late arriving data (Backfill). De-duplication logic variance.



