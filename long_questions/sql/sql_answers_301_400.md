## 🟢 Tracking User Events & Analytics

### Question 301: How would you track a user’s complete clickstream path?

**Answer:**
We need to sequence events by timestamp for each session.
**Query:**
```sql
SELECT 
    UserID, 
    SessionID, 
    Timestamp, 
    EventType,
    LEAD(EventType) OVER (PARTITION BY SessionID ORDER BY Timestamp) as NextEvent
FROM Clickstream;
```

---

### Question 302: Write a query to calculate time between specific events (e.g. add\_to\_cart → purchase).

**Answer:**
Join the table to itself on UserID, where Event 1 is 'Cart' and Event 2 is 'Purchase', and Event 2 time > Event 1 time.
```sql
WITH Events AS (
    SELECT UserID, EventType, Timestamp FROM Log
)
SELECT 
    E1.UserID, 
    MIN(E1.Timestamp) as CartTime, 
    MIN(E2.Timestamp) as PurchaseTime,
    DATEDIFF(minute, MIN(E1.Timestamp), MIN(E2.Timestamp)) as MinsToPurchase
FROM Events E1
JOIN Events E2 ON E1.UserID = E2.UserID
WHERE E1.EventType = 'add_to_cart' AND E2.EventType = 'purchase' 
AND E2.Timestamp > E1.Timestamp
GROUP BY E1.UserID;
```

---

### Question 303: Identify the most common sequence of three actions by users.

**Answer:**
Concatenate 3 consecutive events using `LEAD`.
```sql
WITH Sequences AS (
    SELECT 
        UserID,
        EventType || ' > ' || 
        LEAD(EventType, 1) OVER (PARTITION BY UserID ORDER BY Time) || ' > ' ||
        LEAD(EventType, 2) OVER (PARTITION BY UserID ORDER BY Time) as Path
    FROM Events
)
SELECT Path, COUNT(*) 
FROM Sequences 
WHERE Path IS NOT NULL
GROUP BY Path 
ORDER BY COUNT(*) DESC;
```

---

### Question 304: Track first touch and last touch attribution for a user.

**Answer:**
Use `FIRST_VALUE` and `LAST_VALUE` (or `ROW_NUMBER`).
```sql
SELECT 
    UserID,
    FIRST_VALUE(Campaign) OVER (PARTITION BY UserID ORDER BY Time) as FirstTouch,
    FIRST_VALUE(Campaign) OVER (PARTITION BY UserID ORDER BY Time DESC) as LastTouch
FROM UserHits;
```

---

### Question 305: Write SQL to calculate session length using event timestamps.

**Answer:**
`MAX(Time) - MIN(Time)` group by session.

---

### Question 306: How would you identify bot-like behavior in click data?

**Answer:**
Extremely high frequency (e.g., > 100 clicks per minute) or perfect periodic intervals.

---

### Question 307: Filter event logs where users clicked the same button repeatedly within 10 seconds.

**Answer:**
Self join or LAG check where `CurrentTime - PrevTime < 10 sec` and `CurrentElement = PrevElement`.

---

### Question 308: Count the number of times a user performed an action after seeing a banner.

**Answer:**
Join Impression Log to Action Log on UserID where ActionTime > ImpressionTime.

---

### Question 309: Find users who performed a sequence A → B → C in order.

**Answer:**
Similar to Q303, filter for 'A > B > C'.

---

### Question 310: Compare engagement metrics between users who saw an experiment and those who didn’t.

**Answer:**
Left Join Users to Experiments table. Group by ExperimentGroup (Control vs Test).

---

## 🟡 Business Logic & Data Contracts

### Question 311: Write a constraint to enforce that end\_date must be after start\_date.

**Answer:**
`ALTER TABLE T ADD CONSTRAINT CK_Dates CHECK (EndDate > StartDate);`

---

### Question 312: Create a check that ensures salary is non-negative and within company limits.

**Answer:**
`CHECK (Salary >= 0 AND Salary <= 500000)`

---

### Question 313: How do you enforce allowed values in a string column?

**Answer:**
`CHECK (Status IN ('Active', 'Inactive', 'Pending'))`

---

### Question 314: Detect records violating a uniqueness constraint across multiple columns.

**Answer:**
`GROUP BY Col1, Col2 HAVING COUNT(*) > 1`.

---

### Question 315: Write SQL to identify schema mismatches in dynamic ingestion tables.

**Answer:**
Query System Catalog / Information Schema where Column Name matches but Data Type differs.

---

### Question 316: Enforce business rule: users can’t be in more than one active trial at a time.

**Answer:**
Conditional Unique Index (where `IsActive=1`).

---

### Question 317: Identify foreign keys pointing to non-existent parent records.

**Answer:**
Left Join where Parent ID is NULL (Orphan check).

---

### Question 318: Write a query to detect duplicated timestamps for the same user.

**Answer:**
`GROUP BY UserID, Timestamp HAVING COUNT(*) > 1`.

---

### Question 319: Audit table where deleted rows must leave an audit trail.

**Answer:**
Trigger `AFTER DELETE` -> Insert into Audit.

---

### Question 320: Validate that all customer emails follow proper syntax using regex or LIKE.

**Answer:**
`WHERE Email NOT LIKE '%_@__%.__%'`

---

## 🔵 Systems Thinking with SQL

### Question 321: How do you model many-to-many relationships in SQL schema?

**Answer:**
Junction Table (Association Table) with two Foreign Keys.

---

### Question 322: What are surrogate keys and why would you use them?

**Answer:**
Artificial ID (Integers/UUIDs). Decouples DB keys from business data changes (like email changing).

---

### Question 323: Model a feature flag system using relational tables.

**Answer:**
Tables: `Features`, `Segments`, `UserAssignments`.

---

### Question 324: Design a product subscription model with upgrades and downgrades.

**Answer:**
Subscription Table with `PlanID`, `EffectiveDate`, `ExpiryDate`.

---

### Question 325: Write SQL to model and report on customer referral trees.

**Answer:**
Self-Referencing Table (`ReferrerID` points to `UserID`). Use Recursive CTE to traverse.

---

### Question 326: Implement tag-based searching across products in SQL.

**Answer:**
Table `ProductTags` (ProductID, Tag). Query: `WHERE Tag IN ('A', 'B') GROUP BY ProductID HAVING COUNT(*) = 2`.

---

### Question 327: Write a SQL schema to support multivariate A/B testing.

**Answer:**
`Experiments`, `Variants`, `UserAssignments` (UserID, VariantID).

---

### Question 328: How would you implement data versioning using SQL?

**Answer:**
Add `VersionNumber` to Primary Key, or valid time ranges.

---

### Question 329: Model and report on hierarchical department structures.

**Answer:**
Adjacency List (`ParentDeptID`).

---

### Question 330: Create a changelog table that captures inserts, updates, and deletes in detail.

**Answer:**
Columns: `TableName`, `RecordID`, `Operation` (I/U/D), `OldValue` (JSON), `NewValue` (JSON), `Timestamp`.

---

## 🟠 Anti-patterns & Tricky SQL Logic

### Question 331: What’s wrong with using SELECT \* in production?

**Answer:**
Fetching unnecessary columns wastes I/O. Schema changes (add column) can break application code anticipating fixed column count.

---

### Question 332: How does a GROUP BY without aggregation behave?

**Answer:**
Behaves like `DISTINCT` (returns unique combinations).

---

### Question 333: Why is DISTINCT often misused, and how would you fix it?

**Answer:**
Used to patch duplication caused by bad joins. Fix the join logic instead of masking it with COSTLY Distinct.

---

### Question 334: Explain a case where a LEFT JOIN becomes an INNER JOIN due to filters.

**Answer:**
`SELECT * FROM A LEFT JOIN B ON ... WHERE B.Col = 1`.
The `WHERE` clause removes NULLs generated by the Left Join, effectively turning it into an Inner Join.
**Fix:** Move condition to `ON` clause.

---

### Question 335: What are the risks of filtering inside a JOIN condition?

**Answer:**
For INNER JOIN: No risk.
For LEFT JOIN: Filtering right table in `ON` preserves left rows (returns NULLs for right). Filtering in `WHERE` removes left rows.

---

### Question 336: Write an example where window function is misused causing incorrect logic.

**Answer:**
Filtering on the result of a window function in the *same* SELECT level (not allowed). Must use subquery/CTE.

---

### Question 337: What problems arise from non-deterministic ordering without ORDER BY?

**Answer:**
Databases do not guarantee order. Results may change randomly, breaking pagination or logic relying on "first" row.

---

### Question 338: Detect use of inconsistent aggregation across metrics.

**Answer:**
Using `SUM` for Rate columns (percentages) instead of `AVG` or Weighted Avg.

---

### Question 339: Find cases where string-to-number implicit casting causes wrong outputs.

**Answer:**
'10' > '2' is False (Dictionary sort). Casting to Int makes 10 > 2 True.

---

### Question 340: Describe how NULL logic can break equality checks in joins.

**Answer:**
`NULL = NULL` is False/Unknown. Joins on nullable columns fail to match NULLs. Use `COALESCE` or `IS NOT DISTINCT FROM`.

---

## 🔴 Data Governance & Lineage

### Question 341: How do you track downstream dependencies of a SQL view?

**Answer:**
`information_schema.view_table_usage` or `pg_depend`.

---

### Question 342: Write SQL to identify columns with PII or sensitive data.

**Answer:**
Query catalog for columns like `%ssn%`, `%email%`, `%password%`.

---

### Question 343: How do you flag stale or unused tables in a warehouse?

**Answer:**
Check `Last_Accessed_Timestamp` from system query logs.

---

### Question 344: Find all views that reference a specific table or column.

**Answer:**
Query system dependency definitions.

---

### Question 345: Write a lineage report showing upstream sources of a KPI metric.

**Answer:**
Recursive query on DAG/Dependency graph (usually handled by tools, but conceptually SQL).

---

### Question 346: Identify duplicate datasets created by different teams.

**Answer:**
Compare schema and row counts/checksums.

---

### Question 347: Enforce column naming conventions using metadata queries.

**Answer:**
`SELECT Column_Name FROM Info_Schema.Columns WHERE Column_Name NOT LIKE '%_%'` (e.g., camelCase vs snake_case).

---

### Question 348: Identify hard-coded logic in SQL scripts (e.g., fixed date ranges).

**Answer:**
Search query texts for `'202...'` literals.

---

### Question 349: Track last query time for each table to flag unused ones.

**Answer:**
System Audit Logs aggregation.

---

### Question 350: Generate a data catalog with column descriptions from information schema.

**Answer:**
`SELECT Table_Name, Column_Name, Data_Type FROM Information_Schema.Columns`.

---

## 🟣 Performance Diagnostics

### Question 351: How would you detect full table scans in your warehouse?

**Answer:**
DB Query Plan / Profiler showing "Seq Scan" or "Table Scan" on large tables.

---

### Question 352: Compare performance between indexed vs non-indexed queries.

**Answer:**
Measure `Execution Time` and `Planning Time` using `EXPLAIN ANALYZE`.

---

### Question 353: What are the downsides of excessive partitioning?

**Answer:**
Metadata overhead. Query planning takes longer to scan 1000s of partitions.

---

### Question 354: Write SQL to find top N largest tables in storage.

**Answer:**
**Postgres:** `pg_total_relation_size`.

---

### Question 355: Detect queries with high CPU usage in past week.

**Answer:**
Query `pg_stat_statements` or `Query_History`.

---

### Question 356: Monitor query latency percentiles using a system table.

**Answer:**
Calculate P95 and P99 of execution duration.

---

### Question 357: Audit long-running queries by user or service.

**Answer:**
Filter active queries > 5 minutes.

---

### Question 358: Track sudden growth in row counts over time.

**Answer:**
Log daily row counts, check Delta > Threshold.

---

### Question 359: Profile skew in joins by inspecting distribution of join keys.

**Answer:**
`SELECT JoinKey, COUNT(*) FROM T GROUP BY 1 ORDER BY 2 DESC` (Hot key problem).

---

### Question 360: Identify and reduce cross joins that are unintentional.

**Answer:**
Look for Joins with `1=1` or missing `ON` conditions.

---

## 🟤 Time-Aware Modeling

### Question 361: Create a slowly changing dimension model with effective/expiry dates.

**Answer:**
Table with `ValidFrom`, `ValidTo`. Current row has `ValidTo = '9999-12-31'`.

---

### Question 362: Track changing primary addresses while keeping history.

**Answer:**
Address History Table linked to UserID.

---

### Question 363: Write SQL to identify overlapping active date ranges for users.

**Answer:**
Self join. `A.Start < B.End AND A.End > B.Start`.

---

### Question 364: Model “as of” reporting logic to look back as of any historical date.

**Answer:**
`WHERE ValidFrom <= @ReportDate AND ValidTo > @ReportDate`.

---

### Question 365: Find the state of a subscription table at the end of every month.

**Answer:**
Join Calendar Month End Dates to Subscription History using "As Of" logic.

---

### Question 366: Rebuild a snapshot of data using audit logs.

**Answer:**
Start with Init State + Replay Inserts/Updates up to T.

---

### Question 367: How would you handle schema drift across time in SQL?

**Answer:**
Use JSON column for flexible attributes or EAV model (Entity-Attribute-Value).

---

### Question 368: Detect partial overlaps in time periods between two records.

**Answer:**
Overlap logic (Q363).

---

### Question 369: Write a query to compute the latest value as of a given timestamp.

**Answer:**
`ORDER BY Timestamp DESC LIMIT 1` WHERE Timestamp <= Target.

---

### Question 370: Model time-aware KPI trends with point-in-time accuracy.

**Answer:**
Daily Snapshot Fact Table.

---

## ⚫ Cross-Domain Logic

### Question 371: Create a pricing tier structure for SaaS usage billing.

**Answer:**
Table `Tiers` (Min, Max, Price). Join usage to Tiers where `Usage BETWEEN Min AND Max`.

---

### Question 372: Model and track shipment status across time and carriers.

**Answer:**
`ShipmentEvents` (ShipmentID, Status, EventTime, Location).

---

### Question 373: Write SQL to generate invoices with dynamic pricing.

**Answer:**
Multiply usage * rate (via join).

---

### Question 374: Design a loyalty point redemption and expiration logic in SQL.

**Answer:**
FIFO logic. Deduct points from oldest available record.

---

### Question 375: Identify drivers with idle time over X minutes between deliveries.

**Answer:**
Diff `NextPickup - CurrentDropoff`.

---

### Question 376: Track active incidents in a ticketing system.

**Answer:**
Status NOT IN ('Resolved', 'Closed').

---

### Question 377: Build a schema to handle multi-currency transactions.

**Answer:**
Columns: `Amount`, `Currency`, `ExchangeRate`, `AmountUSD`.

---

### Question 378: Write SQL to calculate fuel consumption per route for logistics.

**Answer:**
Sum(FuelUsed) Group By RouteID.

---

### Question 379: Detect price anomalies in product catalogs by vendor.

**Answer:**
Z-Score of price per category per vendor.

---

### Question 380: Monitor service usage thresholds for automatic suspension.

**Answer:**
Sum(Usage) > Limit.

---

## 🟥 Uncommon or Niche SQL Features

### Question 381: How do you use `FILTER(WHERE ...)` inside aggregates?

**Answer:**
`SUM(Amount) FILTER (WHERE Status='Paid')` (Postgres).

---

### Question 382: Use `QUALIFY` clause in Snowflake to filter windowed results.

**Answer:**
`SELECT ... QUALIFY ROW_NUMBER() OVER (...) = 1`. Removes need for subquery.

---

### Question 383: What does `SEMI JOIN` or `ANTI JOIN` mean in specific dialects?

**Answer:**
Semi: Exists. Anti: Not Exists.

---

### Question 384: Write a lateral join (or CROSS APPLY) use case.

**Answer:**
Foreach row in Left, execute subquery in Right using Left's columns.

---

### Question 385: Use `PIVOT` or `UNPIVOT` operations for reshaping data.

**Answer:**
(See Q101).

---

### Question 386: Implement set-returning functions in PostgreSQL (e.g., `unnest()`).

**Answer:**
`SELECT unnest(Array)`.

---

### Question 387: Generate dynamic SQL queries using stored procedures.

**Answer:**
Concat string and Execute.

---

### Question 388: Use `MERGE` for upsert logic in warehouse platforms.

**Answer:**
(See Q143).

---

### Question 389: How do you query materialized views efficiently?

**Answer:**
Treat like tables. Ensure Refreshed.

---

### Question 390: Build a recursive query to compute running balance in a ledger.

**Answer:**
Usually Window Function is better, but Recursive CTE can traverse linked list of transactions.

---

## 🟧 Mock Debugging / Fix-the-Bug Questions

### Question 391: This query returns wrong revenue totals — find and fix the error.

**Answer:**
Check specifically for `JOIN` duplication (fan-out) or missing `WHERE` filters.

---

### Question 392: The result duplicates rows unexpectedly — what’s wrong?

**Answer:**
Join on non-unique keys.

---

### Question 393: This funnel step count is higher than step 1 — diagnose it.

**Answer:**
Data quality issue or incorrectly defined logic (Step 2 doesn't require Step 1 completion).

---

### Question 394: This rolling average is not behaving correctly — fix the logic.

**Answer:**
Check `ROWS BETWEEN` frame definition. Default is `UNBOUNDED PRECEDING`.

---

### Question 395: NULLs are interfering with sort order — how would you fix?

**Answer:**
`ORDER BY Col NULLS LAST` (or `CASE WHEN Col IS NULL THEN 1 ELSE 0 ...`).

---

### Question 396: A window function isn't resetting per partition — why?

**Answer:**
Missing `PARTITION BY` clause.

---

### Question 397: The grouping level is too granular — what caused it?

**Answer:**
Included a unique ID or Timestamp in `GROUP BY`.

---

### Question 398: Query uses subqueries but is very slow — suggest fixes.

**Answer:**
Change IN/EXISTS to JOIN, or materialize subquery (CTE).

---

### Question 399: Results look correct but are outdated — identify the cause.

**Answer:**
Querying a stale Table Snapshot or non-refreshed View.

---

### Question 400: A daily report is showing fluctuating user counts — explain possible causes.

**Answer:**
Timezone inconsistencies, late arriving data, or different definition of "Active User" in source logs.
