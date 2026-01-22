## 🟢 Intermediate Queries & Logic

### Question 101: Write a query to transpose rows into columns.

**Answer:**
Transposing rows to columns (Pivoting) can be done using `CASE` statements or the `PIVOT` operator.

**Example:**
Converting:
```
Year | Quarter | Revenue
2023 | Q1      | 100
2023 | Q2      | 120
```
To:
```
Year | Q1_Rev | Q2_Rev
2023 | 100    | 120
```

**Query:**
```sql
SELECT 
    Year,
    MAX(CASE WHEN Quarter = 'Q1' THEN Revenue ELSE 0 END) AS Q1_Rev,
    MAX(CASE WHEN Quarter = 'Q2' THEN Revenue ELSE 0 END) AS Q2_Rev
FROM QuarterlySales
GROUP BY Year;
```

---

### Question 102: How would you calculate the median in SQL?

**Answer:**
The median is the middle value. `AVG()` is not the median.
**Postgres/Oracle:** Use `PERCENTILE_CONT(0.5)`.
```sql
SELECT PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY Salary) 
FROM Employees;
```
**MySQL/SQL Server (older versions):** Use `ROW_NUMBER()` and Count.
```sql
SET @row_index := -1;
SELECT AVG(val) as Median
FROM (
  SELECT @row_index:=@row_index + 1 AS row_index, val
  FROM data
  ORDER BY val
) AS subq
WHERE subq.row_index IN (FLOOR(@row_index / 2) , CEIL(@row_index / 2));
```

---

### Question 103: How do you compare records in a table against a historical version of itself?

**Answer:**
Use a **Self Join** with a date offset or `LAG()` window function.

**Example (Using LAG for Prev vs Current):**
```sql
SELECT 
    T.Date, 
    T.Price, 
    LAG(T.Price) OVER (ORDER BY T.Date) as PrevPrice,
    T.Price - LAG(T.Price) OVER (ORDER BY T.Date) as Difference
FROM StockPrices T;
```

---

### Question 104: How would you convert a comma-separated string into individual rows?

**Answer:**
**PostgreSQL:** `unnest(string_to_array(col, ','))`
**SQL Server:** `STRING_SPLIT(col, ',')`

**Example (SQL Server):**
```sql
SELECT value 
FROM STRING_SPLIT('apple,banana,cherry', ',');
```

---

### Question 105: How do you write an IF...THEN logic in SQL?

**Answer:**
Use the `CASE` statement.

**Example:**
```sql
SELECT 
    OrderID,
    CASE 
        WHEN Amount > 1000 THEN 'Large'
        WHEN Amount > 500 THEN 'Medium'
        ELSE 'Small'
    END AS OrderSize
FROM Orders;
```

---

### Question 106: How do you simulate full outer join in MySQL (which doesn't support it natively)?

**Answer:**
Combine a `LEFT JOIN` and a `RIGHT JOIN` using `UNION`.

**Example:**
```sql
SELECT * FROM A LEFT JOIN B ON A.id = B.id
UNION
SELECT * FROM A RIGHT JOIN B ON A.id = B.id;
```

---

### Question 107: How do you get the Nth row in a table without using `LIMIT`?

**Answer:**
Use `ROW_NUMBER()` or correlated subqueries.

**Example:** Get the 5th row.
```sql
WITH Ranked AS (
    SELECT *, ROW_NUMBER() OVER (ORDER BY ID) as rn
    FROM Users
)
SELECT * FROM Ranked WHERE rn = 5;
```

---

### Question 108: Write a query to remove duplicates but keep the most recent record.

**Answer:**
Using a CTE with `ROW_NUMBER()`.

**Query:**
```sql
WITH Duplicates AS (
    SELECT 
        ID, 
        ROW_NUMBER() OVER (PARTITION BY UniqueCol ORDER BY CreatedAt DESC) as rn
    FROM MyTable
)
DELETE FROM Duplicates WHERE rn > 1;
```

---

### Question 109: How do you apply a calculation conditionally across rows using window functions?

**Answer:**
You can combine `CASE` expressions *inside* the window function or windowing logic.

**Example:** Sum only positive values.
```sql
SUM(CASE WHEN Amount > 0 THEN Amount ELSE 0 END) OVER (PARTITION BY UserID)
```

---

### Question 110: How do you deal with circular references in data joins?

**Answer:**
Circular references (A refers to B, B refers to A) can cause infinite loops in recursive queries.
**Solution:**
1.  **Stop Condition:** Limit recursion depth (e.g., `WHERE Level < 10`).
2.  **Path Tracking:** Keep track of visited nodes in an array/string and stop if a node is revisited.

**Example (Postgres Recursive CTE check):**
```sql
WITH RECURSIVE PathCTE AS (
    SELECT ID, ARRAY[ID] as Path FROM Nodes WHERE Parent IS NULL
    UNION ALL
    SELECT N.ID, P.Path || N.ID
    FROM Nodes N JOIN PathCTE P ON N.Parent = P.ID
    WHERE NOT (N.ID = ANY(P.Path)) -- Stop if cyclic
)
SELECT * FROM PathCTE;
```

---

## 🟡 User Behavior & Funnel Analysis

### Question 111: Identify users who started but didn’t complete a purchase.

**Answer:**
Check specifically for the 'Start' event without a corresponding 'Purchase' event.

**Query:**
```sql
SELECT UserID
FROM Events
WHERE EventType = 'CheckoutStart'
AND UserID NOT IN (
    SELECT UserID FROM Events WHERE EventType = 'PurchaseComplete'
);
```

---

### Question 112: Calculate conversion rate from page view to checkout.

**Answer:**
Formula: (Total Checkouts / Total Page Views) * 100.

**Query:**
```sql
SELECT 
    (SUM(CASE WHEN Event = 'Checkout' THEN 1 ELSE 0 END) * 100.0) / 
    SUM(CASE WHEN Event = 'PageView' THEN 1 ELSE 0 END) as ConversionRate
FROM AnalyticsData;
```

---

### Question 113: How would you construct a funnel in SQL?

**Answer:**
Aggregate counts for each step in the funnel order.

**Query:**
```sql
SELECT 
    COUNT(DISTINCT CASE WHEN Step = 'Landing' THEN UserID END) as Step1_Landing,
    COUNT(DISTINCT CASE WHEN Step = 'SignUp' THEN UserID END) as Step2_SignUp,
    COUNT(DISTINCT CASE WHEN Step = 'Purchase' THEN UserID END) as Step3_Purchase
FROM UserJourney;
```

---

### Question 114: Identify drop-off points in a multi-step signup process.

**Answer:**
Calculate the count of users at each step and see where the biggest % drop occurred.

```sql
WITH Steps AS (
   SELECT 
       StepName, 
       COUNT(DISTINCT UserID) as UserCount,
       LEAD(COUNT(DISTINCT UserID)) OVER (ORDER BY StepOrder) as NextStepCount
   FROM SignupFlow
   GROUP BY StepName, StepOrder
)
SELECT 
    StepName, 
    UserCount, 
    NextStepCount, 
    (UserCount - NextStepCount) as DropOffs
FROM Steps;
```

---

### Question 115: How would you calculate average steps in a user journey?

**Answer:**
Count events per session per user, then average.

**Query:**
```sql
SELECT AVG(StepCount) 
FROM (
    SELECT SessionID, COUNT(*) as StepCount 
    FROM UserLogs 
    GROUP BY SessionID
) sub;
```

---

### Question 116: Find users who returned within 7 days of first visit.

**Answer:**
Join or compare first visit date with subsequent visit dates.

**Query:**
```sql
WITH FirstVisit AS (
    SELECT UserID, MIN(VisitDate) as FirstDate
    FROM Visits
    GROUP BY UserID
)
SELECT DISTINCT V.UserID
FROM Visits V
JOIN FirstVisit FV ON V.UserID = FV.UserID
WHERE V.VisitDate > FV.FirstDate 
AND V.VisitDate <= FV.FirstDate + INTERVAL '7 days';
```

---

### Question 117: How would you create a cohort analysis using SQL?

**Answer:**
1.  **Cohort Date:** Determine the user's first active month (Cohort Month).
2.  **Activity Month:** Determine month of subsequent activities.
3.  **Index:** Calculate (Activity Month - Cohort Month) as the month index (0, 1, 2...).

**Query:**
```sql
SELECT 
    DATE_TRUNC('month', FirstLogin) as CohortMonth,
    DATEDIFF(month, FirstLogin, ActivityDate) as MonthIndex,
    COUNT(DISTINCT UserID) as ActiveUsers
FROM UserActivity
GROUP BY 1, 2;
```

---

### Question 118: Build a time-to-convert distribution for marketing leads.

**Answer:**
Calculate days between 'Lead Created' and 'Converted', then group by buckets (0-7 days, 8-30 days, etc.).

**Query:**
```sql
SELECT 
    CASE 
        WHEN DATEDIFF(day, LeadDate, ConvertDate) <= 7 THEN '0-7 Days'
        WHEN DATEDIFF(day, LeadDate, ConvertDate) <= 30 THEN '8-30 Days'
        ELSE '30+ Days'
    END as TimeToConvert,
    COUNT(*) as LeadCount
FROM Leads
WHERE Status = 'Converted'
GROUP BY 1;
```

---

### Question 119: Write a query to calculate activation rate per signup source.

**Answer:**
Activation Rate = (Activated Users / Total Signups) * 100.

**Query:**
```sql
SELECT 
    Source,
    COUNT(CASE WHEN IsActivated = 1 THEN 1 END) * 100.0 / COUNT(*) as ActivationRate
FROM Signups
GROUP BY Source;
```

---

### Question 120: Identify users who reactivated after being inactive for over 30 days.

**Answer:**
Look for a gap > 30 days between consecutive logins.

**Query:**
```sql
WITH Gaps AS (
    SELECT 
        UserID, 
        LoginDate, 
        LAG(LoginDate) OVER (PARTITION BY UserID ORDER BY LoginDate) as PrevLogin
    FROM Logins
)
SELECT DISTINCT UserID 
FROM Gaps 
WHERE DATEDIFF(day, PrevLogin, LoginDate) > 30;
```

---

## 🔵 Complex Joins & Relations

### Question 121: Join three tables where two of them have a many-to-many relationship.

**Answer:**
Usually involves a junction table (e.g., Students <-> Enrollment <-> Courses).

**Query:**
```sql
SELECT S.Name, C.CourseName
FROM Students S
JOIN Enrollments E ON S.StudentID = E.StudentID
JOIN Courses C ON E.CourseID = C.CourseID;
```

---

### Question 122: How would you handle joining a large fact table with multiple dimension tables?

**Answer:**
1.  **Star Schema Join:** Join Fact table to Dimensions on keys.
2.  **Optimization:** Ensure Foreign Keys in the Fact table are indexed. Filter the Fact table *before* joining if possible (using subquery or WHERE clauses) to reduce dataset size early.

---

### Question 123: How do you resolve data duplication caused by joins?

**Answer:**
Data duplication usually happens when joining one-to-many relationships (parent joined to multiple children replicates parent info).
**Solution:**
1.  Aggregate child data *before* joining.
2.  Use `DISTINCT` if rows are identical.
3.  Use Subqueries to fetch specific child record (like "Latest Status").

---

### Question 124: Write a query to compare two tables and list differences.

**Answer:**
Use `EXCEPT` (or `MINUS` in Oracle).

**Query (Rows in A but not in B):**
```sql
SELECT * FROM TableA
EXCEPT
SELECT * FROM TableB;
```
For bi-directional diff, UNION the result of A-B and B-A.

---

### Question 125: How do you find the latest record for each group in a joined result?

**Answer:**
Aggregate MAX(Date) in a CTE or derived table, then join back.

**Query:**
```sql
SELECT O.CustomerID, O.OrderDate, O.Total
FROM Orders O
JOIN (
    SELECT CustomerID, MAX(OrderDate) as MaxDate
    FROM Orders
    GROUP BY CustomerID
) Latest ON O.CustomerID = Latest.CustomerID AND O.OrderDate = Latest.MaxDate;
```

---

### Question 126: How would you join tables with composite keys?

**Answer:**
Join on **all** components of the key.

**Query:**
```sql
SELECT * FROM TableA A
JOIN TableB B 
    ON A.KeyPart1 = B.KeyPart1 
    AND A.KeyPart2 = B.KeyPart2;
```

---

### Question 127: How do you join a table with itself to calculate differences between rows?

**Answer:**
This is a self-join using an offset condition (usually ID-1 or Date comparison).

**Query:**
```sql
SELECT T1.Date, T1.Value - T2.Value as Diff
FROM Metrics T1
JOIN Metrics T2 ON T1.Date = T2.Date + INTERVAL '1 day';
```

---

### Question 128: How would you join on date ranges instead of exact match?

**Answer:**
Use comparison operators (`BETWEEN`, `<=`, `>=`) in the `ON` clause.

**Query:**
```sql
SELECT E.EventName, P.PromotionName
FROM Events E
JOIN Promotions P 
    ON E.EventDate BETWEEN P.StartDate AND P.EndDate;
```

---

### Question 129: Write a query to show changes over time across related entities.

**Answer:**
Usually requires joining snapshots or history tables.

**Query:**
```sql
SELECT 
    H1.EntityID, 
    H1.Status as OldStatus, 
    H2.Status as NewStatus, 
    H2.ChangeDate
FROM History H1
JOIN History H2 ON H1.EntityID = H2.EntityID 
    AND H2.ChangeDate > H1.ChangeDate;
```

---

### Question 130: What are anti-joins and how do you implement them?

**Answer:**
An anti-join returns rows from one table that have **no match** in another table.
Implemented using `LEFT JOIN ... WHERE NULL` or `NOT EXISTS`.

**Query:**
```sql
SELECT * FROM Users U
WHERE NOT EXISTS (SELECT 1 FROM Orders O WHERE O.UserID = U.UserID);
```

---

## 🟠 Audit & Data Change Tracking

### Question 131: Track when a record was inserted, updated, or deleted.

**Answer:**
*   **Inserted/Updated:** Use `CreatedAt` and `UpdatedAt` timestamp columns. Update `UpdatedAt` via triggers or app logic.
*   **Deleted:** Use Soft Delete (`IsDeleted` flag) or move to an `Archives` table via trigger.

---

### Question 132: How would you implement version control in SQL for audit logs?

**Answer:**
Use a separate **Audit Table** (mirror of main table + Version + Timestamp).
Every UPDATE/INSERT into MainTable inserts a copy into AuditTable.
Or use standard **SCD Type 4** (History table).

---

### Question 133: Write a query to find changes in user address over time.

**Answer:**
Assuming an AddressHistory table exists.

**Query:**
```sql
SELECT 
    UserID, 
    OldAddress, 
    NewAddress, 
    ChangeDate
FROM AddressAuditLog
WHERE UserID = 123
ORDER BY ChangeDate;
```

---

### Question 134: How do you detect deleted records between two data snapshots?

**Answer:**
Use `EXCEPT` / `MINUS` on Primary Keys.

**Query:**
```sql
SELECT ID FROM Snapshot_Yesterday
EXCEPT
SELECT ID FROM Snapshot_Today;
-- Result: IDs that existed yesterday but are gone today (Deleted)
```

---

### Question 135: Implement Slowly Changing Dimension Type 1 vs Type 2 in SQL.

**Answer:**
*   **Type 1 (Overwrite):** Simply `UPDATE` the column. No history.
    ```sql
    UPDATE Users SET Address = 'New St' WHERE ID=1;
    ```
*   **Type 2 (Add Row):** Mark old row inactive, insert new active row.
    ```sql
    UPDATE Users SET IsCurrent=0, EndDate=NOW() WHERE ID=1 AND IsCurrent=1;
    INSERT INTO Users (ID, Address, IsCurrent, StartDate) VALUES (1, 'New St', 1, NOW());
    ```

---

### Question 136: How would you identify hard deletes in a soft delete system?

**Answer:**
A hard delete would be a missing ID where it should be present (or marked `IsDeleted=1`).
If IDs are sequential, finding gaps works. If logs exist, compare log vs current table.

---

### Question 137: Detect price changes in products historically.

**Answer:**
Query the Product history or audit table using `LEAD/LAG`.

**Query:**
```sql
SELECT 
    ProductID,
    Price as OldPrice,
    LEAD(Price) OVER (PARTITION BY ProductID ORDER BY ChangedAt) as NewPrice,
    ChangedAt
FROM ProductPriceHistory;
```

---

### Question 138: Write a query to find duplicate IDs with different values over time.

**Answer:**
Find IDs that appear multiple times with different attributes.
```sql
SELECT ID, COUNT(DISTINCT Attribute) 
FROM HistoryTable 
GROUP BY ID 
HAVING COUNT(DISTINCT Attribute) > 1;
```

---

### Question 139: Build a changelog summary per user.

**Answer:**
Aggregate changes from audit logs.
```sql
SELECT 
    UserID,
    COUNT(*) as TotalChanges,
    MAX(ChangeDate) as LastChange,
    STRING_AGG(ChangedField, ', ') as FieldsChanged
FROM AuditLogs
GROUP BY UserID;
```

---

### Question 140: Detect reinsertions after deletions using timestamp comparisons.

**Answer:**
Check if a UserID appears with a `CreatedDate` *after* a `DeletedDate` record in archives.
```sql
SELECT L.UserID
FROM LiveUsers L
JOIN ArchivedUsers A ON L.UserID = A.UserID
WHERE L.CreatedAt > A.DeletedAt;
```

---

## 🔴 ETL, Data Pipelines & Automation

### Question 141: How would you incrementally load new data using SQL?

**Answer:**
Define a "High Watermark" (usually `Max(ID)` or `Max(UpdatedAt)`).
**Query:**
```sql
INSERT INTO TargetTable
SELECT * FROM SourceTable
WHERE UpdatedAt > (SELECT MAX(UpdatedAt) FROM TargetTable);
```

---

### Question 142: What are common issues with ETL scripts in SQL?

**Answer:**
1.  **Data Quality:** NULLs, duplicates, truncated strings.
2.  **Performance:** Locking source tables, massive transaction log growth.
3.  **Schema Drift:** Source column changes breaking destination.
4.  **Timeouts:** Long-running queries killing connection.

---

### Question 143: How would you deduplicate incoming records using SQL?

**Answer:**
Load into Staging table, then Insert into Target using `ROW_NUMBER` or `NOT EXISTS` logic to pick unique/latest.

**Query (Merge):**
```sql
MERGE INTO Target T
USINGSource S ON T.ID = S.ID
WHEN MATCHED THEN UPDATE ...
WHEN NOT MATCHED THEN INSERT ...;
```

---

### Question 144: Explain idempotency in SQL ETL workflows.

**Answer:**
Idempotency means running the same script multiple times leads to the **same result**, without duplication or errors.
**Implementation:** Always `DELETE` existing data for a specific date partition before `INSERT` ing the new data for that date.

---

### Question 145: Write a query to apply transformations before loading into a dimension table.

**Answer:**
Select explicitly with transformation functions.
```sql
INSERT INTO DimCustomer (Name, Email, Region)
SELECT 
    UPPER(Name), 
    LOWER(TRIM(Email)), 
    COALESCE(Region, 'Unknown')
FROM StagingCustomers;
```

---

### Question 146: How do you validate an ETL result using only SQL?

**Answer:**
Check Counts and Aggregates between Source and Target.
```sql
SELECT 'Source', SUM(Amount) FROM SourceTable
UNION ALL
SELECT 'Target', SUM(Amount) FROM TargetTable;
```
If sums don't match, ETL failed.

---

### Question 147: Compare two datasets (source vs target) post ETL and report mismatches.

**Answer:**
Left Join Source to Target where Target is NULL.
```sql
SELECT S.ID as MissingInTarget
FROM Source S
LEFT JOIN Target T ON S.ID = T.ID
WHERE T.ID IS NULL;
```

---

### Question 148: Build a surrogate key using SQL.

**Answer:**
A surrogate key is a system-generated unique key (like `Identity` or `UUID`) not derived from business data.
**Postgres:** `id SERIAL PRIMARY KEY`
**SQL Server:** `id INT IDENTITY(1,1) PRIMARY KEY`

---

### Question 149: How would you stage large raw files using SQL?

**Answer:**
Load raw data into a **Staging Table** (Heap, no constraints, all VARCHAR columns) first using `COPY` / `BULK INSERT`. Then perform validation/casting/transformation when moving to Production tables.

---

### Question 150: How do you implement late arriving data in SQL ETL?

**Answer:**
If Fact arrives before Dimension (e.g., Sale for unknown Product):
1.  Insert a "dummy" or "placeholder" record in the Dimension table (ID: -1, Name: 'Pending').
2.  Link the Fact to this dummy ID.
3.  When the real Dimension record arrives, update the Dimension table.

---

## 🟣 SQL for Reporting & BI Dashboards

### Question 151: Build a daily active user (DAU) dashboard using SQL.

**Answer:**
```sql
SELECT 
    ActivityDate,
    COUNT(DISTINCT UserID) as DAU
FROM UserActivity
GROUP BY ActivityDate
ORDER BY ActivityDate DESC;
```

---

### Question 152: How do you prepare data for a monthly revenue report?

**Answer:**
Aggregate by month snippet.
```sql
SELECT 
    DATE_TRUNC('month', OrderDate) as Month,
    SUM(Details.UnitPrice * Details.Quantity) as Revenue
FROM Orders O
JOIN OrderDetails Details ON O.OrderID = Details.OrderID
GROUP BY 1;
```

---

### Question 153: Generate a time series report with missing dates filled in.

**Answer:**
Generate a date sequence first, then LEFT JOIN data to it.
```sql
WITH DateSeries AS (
    SELECT generate_series('2023-01-01', '2023-01-31', interval '1 day') as Day
)
SELECT D.Day, COALESCE(COUNT(Sales.ID), 0)
FROM DateSeries D
LEFT JOIN Sales ON Sales.Date = D.Day
GROUP BY D.Day;
```

---

### Question 154: How do you format monetary values in SQL reports?

**Answer:**
**Postgres:** `TO_CHAR(Amount, 'FM$999,999,990.00')`
**SQL Server:** `FORMAT(Amount, 'C')`

---

### Question 155: Build a year-over-year comparison of revenue.

**Answer:**
Use Self Join or Lag on matching specific interval.

**Query:**
```sql
SELECT 
    Year(Current.Date) as Year,
    SUM(Current.Amount) as ThisYearSales,
    SUM(Prev.Amount) as LastYearSales
FROM Sales Current
LEFT JOIN Sales Prev 
    ON Year(Current.Date) = Year(Prev.Date) + 1 
    AND Month(Current.Date) = Month(Prev.Date)
GROUP BY 1;
```

---

### Question 156: Write a KPI dashboard query for average resolution time of tickets.

**Answer:**
```sql
SELECT 
    Date_Trunc('week', CreatedAt) as Week,
    AVG(EXTRACT(EPOCH FROM (ResolvedAt - CreatedAt))/3600) as AvgHoursToResolve
FROM Tickets
WHERE Status = 'Resolved'
GROUP BY 1;
```

---

### Question 157: Create a report showing top 10 vs bottom 10 performers.

**Answer:**
Use UNION of two limits.
```sql
(SELECT Name, Sales, 'Top' as Type FROM SalesRep ORDER BY Sales DESC LIMIT 10)
UNION ALL
(SELECT Name, Sales, 'Bottom' as Type FROM SalesRep ORDER BY Sales ASC LIMIT 10);
```

---

### Question 158: Design a query for drill-down reporting by region → product → day.

**Answer:**
Use `GROUPING SETS` or `ROLLUP` (as seen in Q31) to get all hierarchical aggregations in one pass.
```sql
SELECT Region, Product, Day, SUM(Sales)
FROM SalesData
GROUP BY ROLLUP(Region, Product, Day);
```

---

### Question 159: Build a dynamic filter for reporting using SQL variables.

**Answer:**
```sql
-- Variable
SET @TargetRegion = 'East';

SELECT * FROM Sales
WHERE (@TargetRegion IS NULL OR Region = @TargetRegion);
```
If `@TargetRegion` is NULL, it returns ALL.

---

### Question 160: Write a SQL snippet that prepares data for pie chart visualizations.

**Answer:**
Pie charts need Category and Value (and usually %).
```sql
SELECT 
    Category, 
    COUNT(*) as Value,
    (COUNT(*) * 100.0 / (SELECT COUNT(*) FROM Data)) as Percentage
FROM Data
GROUP BY Category;
```

---

## 🟤 Error Handling, Edge Cases & Data Integrity

### Question 161: How would you handle outliers in SQL?

**Answer:**
Filter them out using standard deviation or percentiles.
```sql
-- Remove Top/Bottom 1%
FROM Sales
WHERE Amount BETWEEN PERCENTILE_CONT(0.01) AND PERCENTILE_CONT(0.99);
```

---

### Question 162: Detect NULLs in critical columns and count them.

**Answer:**
```sql
SELECT 
    SUM(CASE WHEN UserID IS NULL THEN 1 ELSE 0 END) as NullIDs,
    SUM(CASE WHEN Email IS NULL THEN 1 ELSE 0 END) as NullEmails
FROM Users;
```

---

### Question 163: What would you do if you encounter unexpected duplicate primary keys?

**Answer:**
1.  Verify constraints (`ALTER TABLE ADD PRIMARY KEY` likely missing).
2.  Identify duplicates (using `HAVING COUNT > 1`).
3.  Clean up duplicates (keep latest/oldest).
4.  Add Constraint to prevent recurrence.

---

### Question 164: Validate that all foreign keys are valid.

**Answer:**
Use `LEFT JOIN` to find orphans.
```sql
SELECT Child.ID 
FROM ChildTable Child
LEFT JOIN ParentTable Parent ON Child.ParentID = Parent.ID
WHERE Parent.ID IS NULL; 
-- Returns Child IDs with invalid ParentID
```

---

### Question 165: Write a query that fails if more than 10% of data is missing.

**Answer:**
Usually done in procedural SQL or script wrapper.
```sql
IF (SELECT COUNT(*) FROM Table WHERE Col IS NULL) * 1.0 / (SELECT COUNT(*) FROM Table) > 0.1
   THROW 50000, 'Data Quality Error: >10% Missing', 1;
```

---

### Question 166: How do you build a “data completeness” score for a dataset?

**Answer:**
Sum populated fields / Total possible fields.
```sql
SELECT 
   (COUNT(Name) + COUNT(Phone) + COUNT(Email)) * 100.0 / (COUNT(*) * 3) as CompletenessScore
FROM Users;
-- *3 because we are checking 3 columns
```

---

### Question 167: Write a check to ensure sales data always has a date.

**Answer:**
Ideally a `NOT NULL` constraint.
Query check:
```sql
SELECT COUNT(*) FROM Sales WHERE SalesDate IS NULL; -- Should be 0
```

---

### Question 168: Flag future-dated transactions in production data.

**Answer:**
```sql
SELECT * FROM Transactions WHERE TxnDate > GETDATE();
```

---

### Question 169: How do you check for skewed distributions in SQL?

**Answer:**
Check Mean vs Median, or count frequency of top keys.
```sql
SELECT Key, COUNT(*) 
FROM PartitionTable 
GROUP BY Key 
ORDER BY COUNT(*) DESC 
LIMIT 5;
```
If top key has 90% of data, it's skewed.

---

### Question 170: Write a query to ensure no orphan records exist in child tables.

**Answer:**
Same as Q164. Use `NOT EXISTS`.
```sql
SELECT * FROM Orders O
WHERE NOT EXISTS (SELECT 1 FROM Customers C WHERE C.ID = O.CustomerID);
```

---

## ⚫ Security & Access Control

### Question 171: How do you prevent unauthorized access using SQL roles?

**Answer:**
Create Roles with specific permissions, then assign Users to Roles.
```sql
CREATE ROLE ReadOnly;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO ReadOnly;
GRANT ReadOnly TO UserJohn;
```

---

### Question 172: What’s the difference between GRANT and REVOKE in SQL?

**Answer:**
*   **GRANT**: Gives a permission (SELECT, INSERT, etc.).
*   **REVOKE**: Removes a previously granted permission.

---

### Question 173: Can you mask sensitive data (e.g., PII) in SQL?

**Answer:**
Yes, Dynamic Data Masking (SQL Server / Azure SQL).
```sql
ALTER TABLE Users 
ALTER COLUMN Email ADD MASKED WITH (FUNCTION = 'email()');
```
Users without UNMASK permission see `x***@xxxx.com`.

---

### Question 174: How would you implement row-level security?

**Answer:**
Row-Level Security (RLS) restricts rows based on user context.
**Logic:** Create a Security Policy with a Predicate Function (e.g., `WHERE ManagerID = USER_ID()`).

---

### Question 175: What steps would you take to log and audit access to sensitive tables?

**Answer:**
1.  Enable **Database Audit Specifications** (logs SELECT/INSERT to file).
2.  Use Triggers to log to an Audit Table.
3.  Check `sys.dm_exec_sessions` or logs regularly.

---

### Question 176: How do you redact customer names in query results?

**Answer:**
Use String manipulation.
```sql
SELECT CONCAT(LEFT(Name, 1), '****') as RedactedName FROM Users;
```

---

### Question 177: How would you separate user permissions between dev and prod databases?

**Answer:**
Different environments should have different Identity Providers or distinct SQL logins. Developers should have `db_owner` in Dev but only `db_datareader` (or no access) in Prod.

---

### Question 178: How do you safely query encrypted fields?

**Answer:**
If encrypted at rest (TDE), querying is transparent.
If encrypted at column level, you need to open the Master Key/Certificate in the session to decrypt.
```sql
OPEN SYMMETRIC KEY MyKey DECRYPTION BY CERTIFICATE MyCert;
SELECT Convert(varchar, DecryptByKey(EncryptedColumn)) FROM Table;
```

---

### Question 179: Create a view that only shows non-sensitive data to analysts.

**Answer:**
```sql
CREATE VIEW PublicSales AS
SELECT Date, Region, Product, Amount -- Exclude CustomerName/CreditCard
FROM Sales;
```
Grant SELECT on `PublicSales` but DENY on `Sales` table.

---

### Question 180: What are SQL injection attacks and how do you avoid them in dynamic SQL?

**Answer:**
If using Dynamic SQL (`EXEC(@sql)`), ensure parameters are passed via `sp_executesql`, NOT string concatenation.
```sql
-- Bad
EXEC('SELECT * FROM T WHERE Name = ' + @name)

-- Good
EXEC sp_executesql N'SELECT * FROM T WHERE Name = @n', N'@n varchar(50)', @name
```

---

## 🟥 SQL with Tools & Systems

### Question 181: How does SQL syntax differ across platforms (e.g., MySQL vs PostgreSQL vs Redshift)?

**Answer:**
*   **Quoting:** MySQL uses backticks \` \`, Postgres uses double quotes `" "`.
*   **Functions:** `TOP` (SQL Server) vs `LIMIT` (Postgres/MySQL). `ISNULL` (SQL Server) vs `COALESCE` (Standard) vs `IFNULL` (MySQL).
*   **Date Math:** `DATEADD` (SQL Server) vs `Date + Interval` (Postgres).

---

### Question 182: How would you write SQL that works cross-database (Snowflake, BigQuery)?

**Answer:**
Stick to **ANSI SQL Standard**.
*   Use `CASE`, `COALESCE`, `CAST`.
*   Avoid proprietary functions (`GETDATE()`, `NVL`). Use `CURRENT_TIMESTAMP`, `COALESCE`.
*   Use CTEs.

---

### Question 183: What are best practices for writing reusable SQL in dbt?

**Answer:**
1.  Use **Jinja** templating (`{{ ref('table') }}`).
2.  Break logic into small, modular models (CTEs).
3.  Use Macros for repeated logic.
4.  Test constraints (unique, not null) via YAML.

---

### Question 184: How would you handle slowly changing dimensions in Looker?

**Answer:**
Looker models usually query the current state.
If history is needed, model the Type 2 table using `sql_always_where` to filter `IsCurrent = Yes`, or expose `ValidFrom/ValidTo` as filterable dimensions for "As Of" analysis.

---

### Question 185: What’s the role of SQL in Airflow DAGs?

**Answer:**
Airflow orchestrates SQL execution using operators (`PostgresOperator`, `SnowflakeOperator`). SQL files are usually transformation steps (ELT) executed on the Data Warehouse.

---

### Question 186: How do you integrate SQL queries into Power BI or Tableau?

**Answer:**
*   **Power BI:** "Get Data" -> SQL Server -> Paste Query (Import or DirectQuery).
*   **Tableau:** Connect to Data -> Custom SQL.
*   *Best Practice:* create a View in DB and connect the BI tool to the View, rather than embedding raw SQL in the tool.

---

### Question 187: How would you write test cases for SQL models in dbt?

**Answer:**
In `schema.yml`:
```yaml
models:
  - name: my_model
    columns:
      - name: id
        tests:
          - unique
          - not_null
```

---

### Question 188: How would you write a modular SQL model in Dataform or dbt?

**Answer:**
Reference upstream tables dynamically.
**dbt:** `FROM {{ ref('upstream_model') }}`
**Dataform:** `FROM ${ref("upstream_table")}`
This builds a dependency graph automatically.

---

### Question 189: How do you pass parameters into SQL in BI tools?

**Answer:**
*   **Tableau:** Parameters map to WHERE clauses/Calculated fields. `WHERE Region = <Parameters.Region>`.
*   **Power BI:** Query Parameters (M query) -> mapped to SQL string.

---

### Question 190: How would you optimize queries in a cloud warehouse (BigQuery, Snowflake)?

**Answer:**
1.  **Partitioning:** Filter by Partition Key (usually Date) is mandatory.
2.  **Clustering:** Sort data by frequently filtered columns.
3.  **Selectivity:** Select only needed columns (Columnar storage implies cost is per-column).

---

## 🟧 Scenario-Based Business Questions

### Question 191: Find the top 3 reasons customers churn.

**Answer:**
Analyze `ChurnReason` column.
```sql
SELECT Reason, COUNT(*) 
FROM ChurnedUsers 
GROUP BY Reason 
ORDER BY COUNT(*) DESC 
LIMIT 3;
```

---

### Question 192: Compare pre- and post-campaign conversion rates.

**Answer:**
Compare rates before StartDate and after StartDate.
```sql
SELECT 
    CASE WHEN Date < '2023-01-01' THEN 'Pre' ELSE 'Post' END as Period,
    AVG(ConversionRate)
FROM DailyMetrics
GROUP BY 1;
```

---

### Question 193: Analyze sales trends during promotional periods.

**Answer:**
Join Sales to Promotions and compare Avg Sales.
```sql
SELECT P.PromoName, AVG(S.Amount)
FROM Sales S
JOIN Promotions P ON S.Date BETWEEN P.Start AND P.End
GROUP BY P.PromoName;
```

---

### Question 194: Identify high-value vs low-value customers.

**Answer:**
Use RFM (Recency, Frequency, Monetary) or simple quantile bucketing.
```sql
NTILE(3) OVER (ORDER BY TotalSpend DESC)
-- 1 = High, 2 = Mid, 3 = Low
```

---

### Question 195: Write SQL to find product cannibalization (drop in product A after launch of product B).

**Answer:**
Compare Product A sales before B's LaunchDate vs After.
```sql
SELECT 
    Period, 
    AVG(DailySales) 
FROM (
    SELECT 
        CASE WHEN Date < 'LaunchDateB' THEN 'Pre' ELSE 'Post' END as Period,
        Amount as DailySales
    FROM Sales
    WHERE Product = 'A'
)
GROUP BY Period;
```

---

### Question 196: Determine correlation between support ticket volume and customer churn.

**Answer:**
(Advanced) Calculate bucketed ticket counts and churn rate per bucket.
```sql
SELECT 
    TicketsBucket, -- e.g., 0, 1-3, 4+
    AVG(IsChurned) as ChurnRate
FROM CustomerStats
GROUP BY TicketsBucket;
```

---

### Question 197: Detect seasonal trends in product sales.

**Answer:**
Group by Month of Year.
```sql
SELECT MONTH(Date), AVG(Sales) 
FROM Sales 
GROUP BY MONTH(Date) 
ORDER BY 1;
```

---

### Question 198: Find which marketing channel has highest ROI using only SQL.

**Answer:**
ROI = (Revenue - Cost) / Cost.
```sql
SELECT 
    Channel, 
    (SUM(Revenue) - SUM(Cost)) / SUM(Cost) as ROI
FROM MarketingData
GROUP BY Channel
ORDER BY ROI DESC 
LIMIT 1;
```

---

### Question 199: Analyze customer feedback patterns using SQL text functions.

**Answer:**
Look for keywords in 'Feedback' column.
```sql
SELECT 
    SUM(CASE WHEN Feedback LIKE '%slow%' THEN 1 ELSE 0 END) as Complaints_Slow,
    SUM(CASE WHEN Feedback LIKE '%expensive%' THEN 1 ELSE 0 END) as Complaints_Price
FROM Reviews;
```

---

### Question 200: Predict potential out-of-stock inventory situations with lead time calculations.

**Answer:**
Stock < (DailyUsage * LeadTimeDays).
```sql
SELECT Product, CurrentStock
FROM Inventory
WHERE CurrentStock < (AvgDailySales * LeadTimeDays);
```
