## 🟢 Analytical Thinking with SQL

### Question 201: Identify users who make purchases only on weekends.

**Answer:**
We need users where *all* their purchases fall on a Saturday (6) or Sunday (0 or 7 depending on DB).

**Query:**
```sql
SELECT UserID
FROM Sales
GROUP BY UserID
HAVING COUNT(*) = SUM(CASE WHEN DATEPART(dw, OrderDate) IN (1, 7) THEN 1 ELSE 0 END);
-- Note: Check specific DB syntax for Day of Week (1=Sunday in SQL Server, 0/7 in Postgres)
```

---

### Question 202: Write SQL to calculate the average session duration per user.

**Answer:**
Session Duration = LogoutTime - LoginTime.
```sql
SELECT 
    UserID, 
    AVG(EXTRACT(EPOCH FROM (LogoutTime - LoginTime))/60) as AvgMinutes
FROM UserSessions
GROUP BY UserID;
```

---

### Question 203: Segment users by frequency of interaction (daily, weekly, monthly).

**Answer:**
Calculate difference between consecutive logins.
```sql
WITH Freq AS (
    SELECT UserID, AVG(DATEDIFF(day, PrevDate, LoginDate)) as AvgGap
    FROM (
        SELECT UserID, LoginDate, LAG(LoginDate) OVER (PARTITION BY UserID ORDER BY LoginDate) as PrevDate
        FROM Logins
    ) T
    WHERE PrevDate IS NOT NULL
)
SELECT 
    UserID,
    CASE 
        WHEN AvgGap <= 1 THEN 'Daily'
        WHEN AvgGap <= 7 THEN 'Weekly'
        ELSE 'Monthly' 
    END as Category
FROM Freq;
```

---

### Question 204: How would you identify peak transaction hours per region?

**Answer:**
Group by Region and Hour, then rank.
```sql
WITH HourlyStats AS (
    SELECT 
        Region, 
        EXTRACT(HOUR FROM TxnTime) as Hour, 
        COUNT(*) as Volume,
        RANK() OVER (PARTITION BY Region ORDER BY COUNT(*) DESC) as rnk
    FROM Transactions
    GROUP BY 1, 2
)
SELECT Region, Hour, Volume 
FROM HourlyStats 
WHERE rnk = 1;
```

---

### Question 205: Write a query to analyze churn trend by subscription duration.

**Answer:**
```sql
SELECT 
    DATEDIFF(month, StartDate, EndDate) as DurationMonths,
    COUNT(*) as ChurnedUsers
FROM Subscriptions
WHERE Status = 'Cancelled'
GROUP BY 1
ORDER BY 1;
```

---

### Question 206: Calculate moving averages over a 7-day window.

**Answer:**
```sql
SELECT 
    Date,
    AVG(Sales) OVER (ORDER BY Date ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) as MovingAvg
FROM DailySales;
```

---

### Question 207: Find correlation between product price and return rate using SQL.

**Answer:**
Calculate Return Rate per Product, then order by Price.
```sql
SELECT 
    ProductName, 
    Price,
    (SUM(CASE WHEN Type='Return' THEN 1 ELSE 0 END) * 1.0 / COUNT(*)) as ReturnRate
FROM Sales
GROUP BY ProductName, Price
ORDER BY Price;
```

---

### Question 208: Write SQL to calculate win/loss ratio by sales rep.

**Answer:**
```sql
SELECT 
    SalesRep,
    SUM(CASE WHEN Status='Won' THEN 1 ELSE 0 END) * 1.0 / 
    NULLIF(SUM(CASE WHEN Status='Lost' THEN 1 ELSE 0 END), 0) as WinLossRatio
FROM Deals
GROUP BY SalesRep;
```

---

### Question 209: How would you analyze failed login attempts per IP?

**Answer:**
```sql
SELECT 
    IPAddress, 
    COUNT(*) as Failures,
    MIN(AttemptTime), 
    MAX(AttemptTime)
FROM LoginLogs
WHERE Status = 'Failed'
GROUP BY IPAddress
HAVING COUNT(*) > 5
ORDER BY Failures DESC;
```

---

### Question 210: Track average time between first and last touchpoint in a funnel.

**Answer:**
```sql
SELECT 
    AVG(DATEDIFF(hour, FirstTouch, LastTouch)) as AvgTimeHours
FROM (
    SELECT 
        UserID, 
        MIN(Timestamp) as FirstTouch, 
        MAX(Timestamp) as LastTouch
    FROM UserEvents
    GROUP BY UserID
) Users;
```

---

## 🟡 Complex Data Types and Advanced Structures

### Question 211: How do you query nested JSON in a column?

**Answer:**
**Postgres:** `data->'address'->>'city'`
**SQL Server:** `JSON_VALUE(data, '$.address.city')`

---

### Question 212: Extract all keys from a JSON object stored in SQL.

**Answer:**
**Postgres:** `jsonb_object_keys(json_column)`
**SQL Server:** `OPENJSON(json_column)` returns Key/Value pairs.

---

### Question 213: Flatten an array column into individual rows.

**Answer:**
**Postgres:** `UNNEST(ArrayCol)`
```sql
SELECT ID, unnest(Tags) as Tag FROM Posts;
```
**SQL Server:** `CROSS APPLY STRING_SPLIT` (or OPENJSON for actual arrays).

---

### Question 214: Convert key-value pairs stored as string into a table.

**Answer:**
Use Parse functions or string splitting.
String: "k1:v1,k2:v2" -> Split by comma -> Split by colon.

---

### Question 215: Write SQL to parse a CSV string inside a column.

**Answer:**
Same as Q104/Q213. Use `STRING_SPLIT` or `UNNEST(STRING_TO_ARRAY)`.

---

### Question 216: Handle JSON logs with variable schemas in SQL.

**Answer:**
Use `JSONB` (Postgres) or `VARIANT` (Snowflake). Query keys dynamically using `COALESCE` if schema varies.
```sql
SELECT 
    COALESCE(log->>'error_msg', log->>'message') as info
FROM Logs;
```

---

### Question 217: Filter records based on values in a JSON array.

**Answer:**
**Postgres:** Check if JSONB Array contains value.
```sql
SELECT * FROM Products WHERE tags @> '["electronics"]';
```
**SQL Server:**
```sql
SELECT * FROM Products 
WHERE 'electronics' IN (SELECT value FROM OPENJSON(Tags));
```

---

### Question 218: Write a query that returns all columns as a single JSON object.

**Answer:**
**Postgres:** `row_to_json(table_name)`
**SQL Server:** `FOR JSON AUTO`

---

### Question 219: Convert JSON column into multiple flat columns.

**Answer:**
```sql
SELECT 
    Data->>'Name' as Name,
    Data->>'Age' as Age
FROM JsonTable;
```

---

### Question 220: Validate if a JSON field has required keys.

**Answer:**
**Postgres:** `?&` operator (Has all keys).
```sql
SELECT * FROM Table WHERE Data ?& array['id', 'name'];
```

---

## 🔵 E-Commerce SQL Problems

### Question 221: Identify abandoned cart users in the past 30 days.

**Answer:**
Items added to cart but no order placed.
```sql
SELECT DISTINCT Cart.UserID 
FROM CartItems Cart
LEFT JOIN Orders O ON Cart.CheckoutID = O.CheckoutID
WHERE Cart.Date > NOW() - INTERVAL '30 days'
AND O.OrderID IS NULL;
```

---

### Question 222: Write SQL to calculate repeat purchase rate by customer.

**Answer:**
Customers with >1 order / Total Customers.
```sql
SELECT 
    COUNT(DISTINCT CASE WHEN OrderCount > 1 THEN UserID END) * 1.0 / 
    COUNT(DISTINCT UserID)
FROM (
    SELECT UserID, COUNT(*) as OrderCount FROM Orders GROUP BY UserID
) Sub;
```

---

### Question 223: Segment customers into first-time vs returning buyers.

**Answer:**
```sql
SELECT 
    CASE WHEN COUNT(*) = 1 THEN 'First-Time' ELSE 'Returning' END as Type,
    COUNT(DISTINCT UserID)
FROM Orders
GROUP BY UserID; -- Note: Needs nested aggregation to count users per type
```
Corrected logic:
```sql
SELECT Type, COUNT(*) 
FROM (
    SELECT UserID, CASE WHEN COUNT(*) = 1 THEN 'First-Time' ELSE 'Returning' END as Type
    FROM Orders GROUP BY UserID
) Sub
GROUP BY Type;
```

---

### Question 224: Analyze conversion rate by traffic source and device.

**Answer:**
```sql
SELECT Source, Device, 
    SUM(IsConverted) * 1.0 / COUNT(*) as ConvRate
FROM Visits
GROUP BY Source, Device;
```

---

### Question 225: Track coupon usage and redemption rates.

**Answer:**
```sql
SELECT 
    CouponCode, 
    COUNT(*) as TotalUses, 
    COUNT(DISTINCT UserID) as UniqueUsers
FROM Orders
WHERE CouponCode IS NOT NULL
GROUP BY CouponCode;
```

---

### Question 226: Find products frequently bought together (market basket analysis).

**Answer:**
Self-join OrderDetails.
```sql
SELECT A.ProductID as Prod1, B.ProductID as Prod2, COUNT(*) as Frequency
FROM OrderDetails A
JOIN OrderDetails B ON A.OrderID = B.OrderID AND A.ProductID < B.ProductID
GROUP BY A.ProductID, B.ProductID
ORDER BY Frequency DESC;
```

---

### Question 227: Calculate average revenue per user (ARPU).

**Answer:**
Total Revenue / Total Unique Users.
```sql
SELECT SUM(Amount) / COUNT(DISTINCT UserID) FROM Orders;
```

---

### Question 228: Identify dormant customers who haven’t purchased in 6 months.

**Answer:**
```sql
SELECT UserID 
FROM Orders 
GROUP BY UserID 
HAVING MAX(OrderDate) < NOW() - INTERVAL '6 months';
```

---

### Question 229: Analyze order-to-fulfillment time per region.

**Answer:**
```sql
SELECT Region, AVG(DATEDIFF(day, OrderDate, ShipDate))
FROM Orders
GROUP BY Region;
```

---

### Question 230: Track inventory turnover by product category.

**Answer:**
(Total Sales Qty / Avg Inventory).
```sql
SELECT 
    Category, 
    SUM(SalesQty) / AVG(StockLevel) as Turnover
FROM InventoryStats
GROUP BY Category;
```

---

## 🟠 Fintech / Transactions

### Question 231: Detect duplicate financial transactions within a short time window.

**Answer:**
Same User, Same Amount, within 10 minutes.
```sql
SELECT T1.ID, T1.UserID, T1.Amount
FROM Txn T1
JOIN Txn T2 ON T1.UserID = T2.UserID 
    AND T1.Amount = T2.Amount 
    AND T1.ID <> T2.ID
    AND ABS(DATEDIFF(minute, T1.Date, T2.Date)) < 10;
```

---

### Question 232: Calculate daily interest accrual for customer accounts.

**Answer:**
Balance * (Rate / 365).
```sql
SELECT AccountID, Balance * (InterestRate / 36500.0) as DailyInterest
FROM Accounts;
```

---

### Question 233: Write SQL to find customers with multiple failed payments.

**Answer:**
```sql
SELECT UserID 
FROM Payments 
WHERE Status = 'Failed' 
GROUP BY UserID 
HAVING COUNT(*) > 3;
```

---

### Question 234: Find high-value transactions flagged for manual review.

**Answer:**
```sql
SELECT * FROM Transactions 
WHERE Amount > 10000 
OR (Amount > 5000 AND Country = 'HighRisk');
```

---

### Question 235: Segment transactions by risk category.

**Answer:**
```sql
SELECT Categories, COUNT(*) 
FROM (
    SELECT CASE WHEN RiskScore > 90 THEN 'Critical' ELSE 'Normal' END as Categories
    FROM TxnRisk
) S
GROUP BY Categories;
```

---

### Question 236: Track account balance over time per customer.

**Answer:**
Running Total of transactions.
```sql
SELECT 
    UserID, Date, Amount,
    SUM(Amount) OVER (PARTITION BY UserID ORDER BY Date) as Balance
FROM Transactions;
```

---

### Question 237: Analyze deposit-to-withdrawal ratio per user.

**Answer:**
```sql
SELECT 
    UserID,
    SUM(CASE WHEN Type='Deposit' THEN Amount ELSE 0 END) / 
    NULLIF(SUM(CASE WHEN Type='Withdrawal' THEN Amount ELSE 0 END), 0)
FROM Transactions
GROUP BY UserID;
```

---

### Question 238: Find customers who cross KYC threshold but are not verified.

**Answer:**
```sql
SELECT T.UserID, SUM(T.Amount)
FROM Transactions T
JOIN Users U ON T.UserID = U.UserID
WHERE U.Verified = 0
GROUP BY T.UserID
HAVING SUM(T.Amount) > 10000;
```

---

### Question 239: Write a reconciliation report between two transaction tables.

**Answer:**
Full Outer Join Gateway vs Ledger.
```sql
SELECT G.ID, G.Amount, L.Amount, (G.Amount - L.Amount) as Diff
FROM GatewayLogs G
FULL OUTER JOIN Ledger L ON G.TxnID = L.TxnID
WHERE G.Amount <> L.Amount OR G.ID IS NULL OR L.ID IS NULL;
```

---

### Question 240: Calculate fraud detection score using SQL logic.

**Answer:**
Sum of risk factors.
```sql
SELECT UserID,
  (CASE WHEN Country = 'X' THEN 50 ELSE 0 END + 
   CASE WHEN Amount > 1000 THEN 30 ELSE 0 END) as FraudScore
FROM Txn;
```

---

## 🔴 IoT / Sensor Data Use Cases

### Question 241: Find devices that haven’t reported in the last 24 hours.

**Answer:**
```sql
SELECT DeviceID 
FROM Pings 
GROUP BY DeviceID 
HAVING MAX(PingTime) < NOW() - INTERVAL '24 hours';
```

---

### Question 242: Calculate average temperature by room and hour.

**Answer:**
```sql
SELECT Room, DATE_PART('hour', Time) as Hr, AVG(Temp)
FROM Readings
GROUP BY Room, Hr;
```

---

### Question 243: Detect out-of-range sensor readings and alert conditions.

**Answer:**
```sql
SELECT * FROM Readings WHERE Value < MinThreshold OR Value > MaxThreshold;
```

---

### Question 244: Calculate uptime percentage of devices over the last month.

**Answer:**
(Pings Received / Expected Pings) * 100.

---

### Question 245: Identify sensors with inconsistent reporting intervals.

**Answer:**
StdDev of time difference between pings.

---

### Question 246: Analyze event frequency per machine.

**Answer:**
Count events group by MachineID.

---

### Question 247: Write SQL to find overlapping signal timestamps.

**Answer:**
Self join where ranges overlap `StartA <= EndB` AND `EndA >= StartB`.

---

### Question 248: Aggregate sensor data into 10-minute time buckets.

**Answer:**
Round timestamp to nearest 10 min.
**Postgres:** `DATE_TRUNC('hour', Time) + INTERVAL '10 min' * FLOOR(DATE_PART('minute', Time) / 10)`

---

### Question 249: Correlate device type with error frequency.

**Answer:**
Join Device Info, Group By Type, Count Errors.

---

### Question 250: Rank devices by reliability score based on downtime.

**Answer:**
Rank by Sum(DowntimeDuration) ASC.

---

## 🟣 Healthcare / Medical Use Cases

### Question 251: Identify patients with recurring diagnoses in last 6 months.

**Answer:**
`HAVING COUNT(*) > 1` on DiagnosisCode + PatientID.

---

### Question 252: Calculate average hospital stay by procedure type.

**Answer:**
`AVG(DischargeDate - AdmitDate)` Group By Procedure.

---

### Question 253: Find patients who missed follow-up appointments.

**Answer:**
Appointment status 'No Show'.

---

### Question 254: Write SQL to track medication adherence.

**Answer:**
Compare Calculated Refill Date vs Actual Refill Date.

---

### Question 255: Compare pre- and post-treatment metrics using window functions.

**Answer:**
Lag() metric order by Date partitioned by Patient.

---

### Question 256: Segment patients by age group and condition.

**Answer:**
Case statement on Age, Group by Condition.

---

### Question 257: Identify facilities with highest readmission rates.

**Answer:**
Readmit = Admit within 30 days of previous Discharge.

---

### Question 258: Detect abnormal lab result values.

**Answer:**
Join LabResults to ReferenceRanges Table.

---

### Question 259: Build a patient journey timeline using SQL.

**Answer:**
Union All of Admissions, Labs, Prescriptions Order By Date.

---

### Question 260: Aggregate diagnosis counts by ICD code and quarter.

**Answer:**
`GROUP BY ICD_Code, DATE_TRUNC('quarter', Date)`.

---

## 🟤 Time Series & Temporal SQL

### Question 261: Create a gap and island analysis of user activity.

**Answer:**
Group consecutive dates into "Islands". Use `Row_Number() - Date` trick.

---

### Question 262: Detect periods of inactivity longer than 3 days.

**Answer:**
`Lead(Date) - Date > 3`.

---

### Question 263: Calculate duration of continuous login streaks.

**Answer:**
Count days in current "Island" (from Q261).

---

### Question 264: How would you fill in missing daily values with zeros?

**Answer:**
Right Join data to Calendar Table. Isnull(Value, 0).

---

### Question 265: Write SQL to interpolate missing values linearly.

**Answer:**
(Advanced) Calculate Slope between known points, project value for nulls.

---

### Question 266: Detect backward timestamp entries (data quality issue).

**Answer:**
`NextTimestamp < CurrentTimestamp`.

---

### Question 267: Create time buckets that shift dynamically by timezone.

**Answer:**
`AT TIME ZONE` conversion before bucketing.

---

### Question 268: Perform day-over-day and week-over-week comparisons.

**Answer:**
`Lag(Val, 1)` and `Lag(Val, 7)`.

---

### Question 269: Calculate delta between current and previous value per device.

**Answer:**
`Value - Lag(Value)`.

---

### Question 270: Track time spent in each state (e.g., online, offline, idle).

**Answer:**
`NextTimestamp - CurrentTimestamp`.

---

## ⚫ Multi-Tenant / SaaS Use Cases

### Question 271: Filter results to only show data belonging to a specific tenant ID.

**Answer:**
Always include `WHERE TenantID = ?`.

---

### Question 272: Write SQL to calculate average usage per customer account.

**Answer:**
Sum(Usage) / Count(Customers).

---

### Question 273: Compare active users across customers on same subscription plan.

**Answer:**
Group by Plan, Customer.

---

### Question 274: Identify top feature usage per tenant.

**Answer:**
Rank() usage count desc partition by Tenant.

---

### Question 275: Detect unusual usage patterns across tenants.

**Answer:**
Avg usage + 2*StdDev (Z-Score).

---

### Question 276: Rank customers by feature adoption rate.

**Answer:**
Distinct Features Used / Total Features Available.

---

### Question 277: Analyze churn by subscription type and usage frequency.

**Answer:**
Group by Plan, UsageBucket.

---

### Question 278: Track support ticket volume per customer.

**Answer:**
Count(*) Group By CustomerID.

---

### Question 279: Build customer health score using SQL metrics.

**Answer:**
Weighted sum of logic (LoginFreq * 0.5 + Usage * 0.3 + Tickets * 0.2).

---

### Question 280: Identify customers who downgraded in the last 90 days.

**Answer:**
Self join subscription history where NewPlan < OldPlan.

---

## 🟧 Security / Access / Logging

### Question 281: Identify IPs with the most failed login attempts.

**Answer:**
`COUNT(*) WHERE Success=0 GROUP BY IP ORDER BY DESC`.

---

### Question 282: Track admin actions by user and timestamp.

**Answer:**
`SELECT * FROM Audit WHERE Role='Admin'`.

---

### Question 283: Detect brute-force login behavior using SQL.

**Answer:**
High count of failures in short window (Q209).

---

### Question 284: Write SQL to determine role hierarchy across users.

**Answer:**
Recursive CTE on Roles table.

---

### Question 285: Audit password reset requests over time.

**Answer:**
Count by Day where Action='PasswordReset'.

---

### Question 286: Detect logins from multiple locations within short timeframe.

**Answer:**
Self join login log, diff IP/Location, diff Time < 1 hour. "Impossible Travel".

---

### Question 287: Compare access rights between two users.

**Answer:**
Intersect / Except on UserPermissions table.

---

### Question 288: Track audit log events by severity level.

**Answer:**
Group by Severity.

---

### Question 289: Identify users who changed permissions in last 7 days.

**Answer:**
`Action='PermissionChange' AND Date > Now-7`.

---

### Question 290: Calculate average session duration per role.

**Answer:**
Join Session to User to Role, then Avg.

---

## 🟥 ETL Testing & Data Quality Assurance

### Question 291: Compare record counts across staging and production tables.

**Answer:**
`Select Count` from both.

---

### Question 292: Write SQL to detect schema drift in imported files.

**Answer:**
Query Information_Schema.Columns.

---

### Question 293: Track failed loads in ETL logs using SQL.

**Answer:**
`Status='Failed'`.

---

### Question 294: Validate primary key uniqueness in raw ingested data.

**Answer:**
`Count distinct ID = Count ID`.

---

### Question 295: Write a row-level comparison between two versions of a table.

**Answer:**
Full Outer Join (Hash of col values).

---

### Question 296: How do you detect schema changes in SQL pipelines?

**Answer:**
Exception handling or Metadata check.

---

### Question 297: Implement row hashing to compare large datasets efficiently.

**Answer:**
`MD5(Concat(Col1, Col2...))`.

---

### Question 298: Monitor freshness of ETL data using load timestamps.

**Answer:**
`Max(LoadDate)`.

---

### Question 299: Track changes in business metrics after ETL change.

**Answer:**
Compare metric before vs after deployment.

---

### Question 300: Identify partial or incomplete loads in partitioned tables.

**Answer:**
Check row counts per partition against expectation.
