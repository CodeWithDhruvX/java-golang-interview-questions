# SQL interview questions

**User:** Dhruv Shah (dhruvshahlinkedin@gmail.com)  
**Created:** 8/4/2025 11:32:20  
**Updated:** 1/22/2026 9:56:07  
**Exported:** 1/22/2026 9:59:17  
**Link:** [https://chatgpt.com/c/68904ce1-06a8-8331-8537-4e1e57fa74fd](https://chatgpt.com/c/68904ce1-06a8-8331-8537-4e1e57fa74fd)  

## Prompt:
sql based real word interview questions around 100

## Response:
Here’s a comprehensive list of **100 real-world, SQL-based interview questions**. These are categorized for better understanding and practice, ranging from basic to advanced and covering performance, logic, and business scenarios.

* * *

### 🟢 **Basic SQL (Select, Where, Order By, Group By)**

1.  What is the difference between `WHERE` and `HAVING`?
2.  How do you retrieve unique records from a table?
3.  Write a query to get the second highest salary from a table.
4.  How do you sort data in ascending and descending order?
5.  How do you filter records between two dates?
6.  What is the difference between `COUNT(*)` and `COUNT(column)`?
7.  How do you handle NULL values in SQL?
8.  Write a query to return the number of employees in each department.
9.  What is the difference between `GROUP BY` and `ORDER BY`?
10.  How do you rename a column in a result set?

* * *

### 🟡 **Joins (Inner, Left, Right, Full)**

11.  Explain INNER JOIN with an example.
12.  What is the difference between LEFT JOIN and RIGHT JOIN?
13.  When would you use FULL OUTER JOIN?
14.  Write a query using JOIN to get employee names and their manager names.
15.  How do you get records from one table that do not exist in another?
16.  What is a self-join? Provide an example.
17.  How do you join more than two tables?
18.  Explain the concept of Cartesian join with a query.
19.  Write a query to find employees who don’t have a department.
20.  Difference between JOIN and UNION.

* * *

### 🟠 **Subqueries & CTEs**

21.  What is a correlated subquery?
22.  How is a subquery different from a join?
23.  Write a query using a subquery to find employees earning above average salary.
24.  What is a CTE? How is it different from a subquery?
25.  How can you use a recursive CTE? Example: find all levels of a hierarchy.
26.  Can you use a subquery in a `SELECT` clause? Example?
27.  How do you write a query to get the top N rows per group?
28.  Explain how `WITH` clause works in SQL.
29.  Can you use CTEs in UPDATE or DELETE statements?
30.  What are the performance considerations between CTE and derived tables?

* * *

### 🔵 **Data Aggregation & Analysis**

31.  What does `ROLLUP` do in SQL?
32.  What is `CUBE` and how is it different from `ROLLUP`?
33.  What are window functions? Give a practical example.
34.  Use `RANK()` and `DENSE_RANK()` to assign rankings within groups.
35.  What does `PARTITION BY` do in window functions?
36.  Write a query using `LEAD()` and `LAG()` functions.
37.  How would you calculate a running total?
38.  Difference between `ROW_NUMBER()` and `RANK()`.
39.  How do you find duplicate rows in a table?
40.  Write a query to pivot data in SQL.

* * *

### 🔴 **Advanced SQL**

41.  What is normalization? Explain different normal forms.
42.  What is denormalization and when is it used?
43.  How would you optimize a slow-running query?
44.  What are indexes? Types? When to use?
45.  What are stored procedures? How are they different from functions?
46.  How do you implement transactions in SQL?
47.  What is ACID in SQL databases?
48.  Explain isolation levels: Read Uncommitted, Committed, Repeatable Read, Serializable.
49.  What is the use of the `EXPLAIN` or `EXPLAIN PLAN`?
50.  Difference between clustered and non-clustered indexes.

* * *

### 🟣 **Real-world Scenarios**

51.  Find users who made purchases on consecutive days.
52.  Write a query to find repeat customers.
53.  Find top 3 selling products in each category.
54.  Calculate customer lifetime value (CLV).
55.  Track changes in customer status using SQL (SCD Type 2).
56.  Identify users with declining monthly usage.
57.  Write a query to calculate retention rate month over month.
58.  Get average time between user signup and first purchase.
59.  Segment users based on activity (e.g., high, medium, low).
60.  Detect anomalies in daily transaction volumes.

* * *

### 🟤 **Date & Time Manipulations**

61.  How do you extract year/month/day from a date?
62.  Write a query to get the last day of each month.
63.  Find the number of working days between two dates.
64.  Calculate age from date of birth.
65.  How do you handle time zones in SQL?
66.  Write a query to get users who signed up in the past 7 days.
67.  Get the number of orders by week.
68.  How do you truncate a date to the start of the month?
69.  How to calculate time difference in hours between two timestamps?
70.  Write a query to get users inactive for more than 30 days.

* * *

### 🟧 **Data Cleaning & Transformation**

71.  How do you remove duplicates in a table?
72.  How do you handle NULLs in aggregate functions?
73.  Use `CASE` to categorize data into buckets (e.g., age groups).
74.  Replace NULL values with default values using `COALESCE()`.
75.  How do you clean up bad data (e.g., phone numbers with text)?
76.  Write a query to standardize country names.
77.  How do you extract numbers from a string?
78.  Normalize string formats (e.g., capitalize all names).
79.  Identify and fix inconsistent data (e.g., dates as strings).
80.  Write SQL to de-duplicate based on business rules.

* * *

### ⚫ **Schema & DDL**

81.  How do you create a new table with constraints?
82.  What’s the difference between `DELETE`, `TRUNCATE`, and `DROP`?
83.  How do you alter an existing table (add/remove columns)?
84.  How to enforce referential integrity?
85.  Explain primary key vs unique key.
86.  What is a composite key?
87.  What are views? Pros and cons?
88.  How do you grant/revoke permissions in SQL?
89.  How do you copy schema and data from one table to another?
90.  What are sequences and auto-increment columns?

* * *

### 🟥 **Miscellaneous / Optimization / Troubleshooting**

91.  What is SQL injection and how to prevent it?
92.  How to handle pagination in SQL?
93.  What tools do you use to debug slow SQL queries?
94.  When would you use temp tables or table variables?
95.  What is a materialized view?
96.  Difference between OLTP and OLAP systems?
97.  How do you monitor long-running queries in production?
98.  Explain batch vs streaming queries.
99.  What is sharding and partitioning in SQL databases?
100.  How do you migrate large amounts of data efficiently using SQL?

* * *

Would you like a printable version (PDF or Word)? Or should I create a quiz set or flashcards for practice?

## Prompt:
give me more without reapting above one

## Response:
Absolutely! Here's an additional set of **100 unique, real-world SQL interview questions** (no repeats from the previous list). These go deeper into **data engineering, reporting, analytics, optimization, security**, and **edge-case handling** — often asked in **product companies, data analyst, data engineer, and business intelligence interviews**.

* * *

### 🟢 **Intermediate Queries & Logic**

101.  Write a query to transpose rows into columns.
102.  How would you calculate the median in SQL?
103.  How do you compare records in a table against a historical version of itself?
104.  How would you convert a comma-separated string into individual rows?
105.  How do you write an IF...THEN logic in SQL?
106.  How do you simulate full outer join in MySQL (which doesn't support it natively)?
107.  How do you get the Nth row in a table without using `LIMIT`?
108.  Write a query to remove duplicates but keep the most recent record.
109.  How do you apply a calculation conditionally across rows using window functions?
110.  How do you deal with circular references in data joins?

* * *

### 🟡 **User Behavior & Funnel Analysis**

111.  Identify users who started but didn’t complete a purchase.
112.  Calculate conversion rate from page view to checkout.
113.  How would you construct a funnel in SQL?
114.  Identify drop-off points in a multi-step signup process.
115.  How would you calculate average steps in a user journey?
116.  Find users who returned within 7 days of first visit.
117.  How would you create a cohort analysis using SQL?
118.  Build a time-to-convert distribution for marketing leads.
119.  Write a query to calculate activation rate per signup source.
120.  Identify users who reactivated after being inactive for over 30 days.

* * *

### 🔵 **Complex Joins & Relations**

121.  Join three tables where two of them have a many-to-many relationship.
122.  How would you handle joining a large fact table with multiple dimension tables?
123.  How do you resolve data duplication caused by joins?
124.  Write a query to compare two tables and list differences.
125.  How do you find the latest record for each group in a joined result?
126.  How would you join tables with composite keys?
127.  How do you join a table with itself to calculate differences between rows?
128.  How would you join on date ranges instead of exact match?
129.  Write a query to show changes over time across related entities.
130.  What are anti-joins and how do you implement them?

* * *

### 🟠 **Audit & Data Change Tracking**

131.  Track when a record was inserted, updated, or deleted.
132.  How would you implement version control in SQL for audit logs?
133.  Write a query to find changes in user address over time.
134.  How do you detect deleted records between two data snapshots?
135.  Implement Slowly Changing Dimension Type 1 vs Type 2 in SQL.
136.  How would you identify hard deletes in a soft delete system?
137.  Detect price changes in products historically.
138.  Write a query to find duplicate IDs with different values over time.
139.  Build a changelog summary per user.
140.  Detect reinsertions after deletions using timestamp comparisons.

* * *

### 🔴 **ETL, Data Pipelines & Automation**

141.  How would you incrementally load new data using SQL?
142.  What are common issues with ETL scripts in SQL?
143.  How would you deduplicate incoming records using SQL?
144.  Explain idempotency in SQL ETL workflows.
145.  Write a query to apply transformations before loading into a dimension table.
146.  How do you validate an ETL result using only SQL?
147.  Compare two datasets (source vs target) post ETL and report mismatches.
148.  Build a surrogate key using SQL.
149.  How would you stage large raw files using SQL?
150.  How do you implement late arriving data in SQL ETL?

* * *

### 🟣 **SQL for Reporting & BI Dashboards**

151.  Build a daily active user (DAU) dashboard using SQL.
152.  How do you prepare data for a monthly revenue report?
153.  Generate a time series report with missing dates filled in.
154.  How do you format monetary values in SQL reports?
155.  Build a year-over-year comparison of revenue.
156.  Write a KPI dashboard query for average resolution time of tickets.
157.  Create a report showing top 10 vs bottom 10 performers.
158.  Design a query for drill-down reporting by region → product → day.
159.  Build a dynamic filter for reporting using SQL variables.
160.  Write a SQL snippet that prepares data for pie chart visualizations.

* * *

### 🟤 **Error Handling, Edge Cases & Data Integrity**

161.  How would you handle outliers in SQL?
162.  Detect NULLs in critical columns and count them.
163.  What would you do if you encounter unexpected duplicate primary keys?
164.  Validate that all foreign keys are valid.
165.  Write a query that fails if more than 10% of data is missing.
166.  How do you build a “data completeness” score for a dataset?
167.  Write a check to ensure sales data always has a date.
168.  Flag future-dated transactions in production data.
169.  How do you check for skewed distributions in SQL?
170.  Write a query to ensure no orphan records exist in child tables.

* * *

### ⚫ **Security & Access Control**

171.  How do you prevent unauthorized access using SQL roles?
172.  What’s the difference between GRANT and REVOKE in SQL?
173.  Can you mask sensitive data (e.g., PII) in SQL?
174.  How would you implement row-level security?
175.  What steps would you take to log and audit access to sensitive tables?
176.  How do you redact customer names in query results?
177.  How would you separate user permissions between dev and prod databases?
178.  How do you safely query encrypted fields?
179.  Create a view that only shows non-sensitive data to analysts.
180.  What are SQL injection attacks and how do you avoid them in dynamic SQL?

* * *

### 🟥 **SQL with Tools & Systems**

181.  How does SQL syntax differ across platforms (e.g., MySQL vs PostgreSQL vs Redshift)?
182.  How would you write SQL that works cross-database (Snowflake, BigQuery)?
183.  What are best practices for writing reusable SQL in dbt?
184.  How would you handle slowly changing dimensions in Looker?
185.  What’s the role of SQL in Airflow DAGs?
186.  How do you integrate SQL queries into Power BI or Tableau?
187.  How would you write test cases for SQL models in dbt?
188.  How would you write a modular SQL model in Dataform or dbt?
189.  How do you pass parameters into SQL in BI tools?
190.  How would you optimize queries in a cloud warehouse (BigQuery, Snowflake)?

* * *

### 🟧 **Scenario-Based Business Questions**

191.  Find the top 3 reasons customers churn.
192.  Compare pre- and post-campaign conversion rates.
193.  Analyze sales trends during promotional periods.
194.  Identify high-value vs low-value customers.
195.  Write SQL to find product cannibalization (drop in product A after launch of product B).
196.  Determine correlation between support ticket volume and customer churn.
197.  Detect seasonal trends in product sales.
198.  Find which marketing channel has highest ROI using only SQL.
199.  Analyze customer feedback patterns using SQL text functions.
200.  Predict potential out-of-stock inventory situations with lead time calculations.

* * *

Would you like:

*   ✅ A CSV or Excel export of all 200+ questions?
*   ✅ A study tracker / checklist version?
*   ✅ Flashcards (Anki-compatible)?  
    Let me know and I can generate it for you!

## Prompt:
give me more without reapting above one

## Response:
Absolutely — here’s **yet another fresh batch of 100 SQL interview questions** that do **not repeat** anything from the **previous 200 questions**. These focus on deeper **analytics, unusual use cases, industry-specific logic**, and **real-world data workflows** used in **e-commerce, fintech, SaaS, logistics, healthcare, media, and IoT domains**.

* * *

### 🟢 **Analytical Thinking with SQL**

201.  Identify users who make purchases only on weekends.
202.  Write SQL to calculate the average session duration per user.
203.  Segment users by frequency of interaction (daily, weekly, monthly).
204.  How would you identify peak transaction hours per region?
205.  Write a query to analyze churn trend by subscription duration.
206.  Calculate moving averages over a 7-day window.
207.  Find correlation between product price and return rate using SQL.
208.  Write SQL to calculate win/loss ratio by sales rep.
209.  How would you analyze failed login attempts per IP?
210.  Track average time between first and last touchpoint in a funnel.

* * *

### 🟡 **Complex Data Types and Advanced Structures**

211.  How do you query nested JSON in a column?
212.  Extract all keys from a JSON object stored in SQL.
213.  Flatten an array column into individual rows.
214.  Convert key-value pairs stored as string into a table.
215.  Write SQL to parse a CSV string inside a column.
216.  Handle JSON logs with variable schemas in SQL.
217.  Filter records based on values in a JSON array.
218.  Write a query that returns all columns as a single JSON object.
219.  Convert JSON column into multiple flat columns.
220.  Validate if a JSON field has required keys.

* * *

### 🔵 **E-Commerce SQL Problems**

221.  Identify abandoned cart users in the past 30 days.
222.  Write SQL to calculate repeat purchase rate by customer.
223.  Segment customers into first-time vs returning buyers.
224.  Analyze conversion rate by traffic source and device.
225.  Track coupon usage and redemption rates.
226.  Find products frequently bought together (market basket analysis).
227.  Calculate average revenue per user (ARPU).
228.  Identify dormant customers who haven’t purchased in 6 months.
229.  Analyze order-to-fulfillment time per region.
230.  Track inventory turnover by product category.

* * *

### 🟠 **Fintech / Transactions**

231.  Detect duplicate financial transactions within a short time window.
232.  Calculate daily interest accrual for customer accounts.
233.  Write SQL to find customers with multiple failed payments.
234.  Find high-value transactions flagged for manual review.
235.  Segment transactions by risk category.
236.  Track account balance over time per customer.
237.  Analyze deposit-to-withdrawal ratio per user.
238.  Find customers who cross KYC threshold but are not verified.
239.  Write a reconciliation report between two transaction tables.
240.  Calculate fraud detection score using SQL logic.

* * *

### 🔴 **IoT / Sensor Data Use Cases**

241.  Find devices that haven’t reported in the last 24 hours.
242.  Calculate average temperature by room and hour.
243.  Detect out-of-range sensor readings and alert conditions.
244.  Calculate uptime percentage of devices over the last month.
245.  Identify sensors with inconsistent reporting intervals.
246.  Analyze event frequency per machine.
247.  Write SQL to find overlapping signal timestamps.
248.  Aggregate sensor data into 10-minute time buckets.
249.  Correlate device type with error frequency.
250.  Rank devices by reliability score based on downtime.

* * *

### 🟣 **Healthcare / Medical Use Cases**

251.  Identify patients with recurring diagnoses in last 6 months.
252.  Calculate average hospital stay by procedure type.
253.  Find patients who missed follow-up appointments.
254.  Write SQL to track medication adherence.
255.  Compare pre- and post-treatment metrics using window functions.
256.  Segment patients by age group and condition.
257.  Identify facilities with highest readmission rates.
258.  Detect abnormal lab result values.
259.  Build a patient journey timeline using SQL.
260.  Aggregate diagnosis counts by ICD code and quarter.

* * *

### 🟤 **Time Series & Temporal SQL**

261.  Create a gap and island analysis of user activity.
262.  Detect periods of inactivity longer than 3 days.
263.  Calculate duration of continuous login streaks.
264.  How would you fill in missing daily values with zeros?
265.  Write SQL to interpolate missing values linearly.
266.  Detect backward timestamp entries (data quality issue).
267.  Create time buckets that shift dynamically by timezone.
268.  Perform day-over-day and week-over-week comparisons.
269.  Calculate delta between current and previous value per device.
270.  Track time spent in each state (e.g., online, offline, idle).

* * *

### ⚫ **Multi-Tenant / SaaS Use Cases**

271.  Filter results to only show data belonging to a specific tenant ID.
272.  Write SQL to calculate average usage per customer account.
273.  Compare active users across customers on same subscription plan.
274.  Identify top feature usage per tenant.
275.  Detect unusual usage patterns across tenants.
276.  Rank customers by feature adoption rate.
277.  Analyze churn by subscription type and usage frequency.
278.  Track support ticket volume per customer.
279.  Build customer health score using SQL metrics.
280.  Identify customers who downgraded in the last 90 days.

* * *

### 🟧 **Security / Access / Logging**

281.  Identify IPs with the most failed login attempts.
282.  Track admin actions by user and timestamp.
283.  Detect brute-force login behavior using SQL.
284.  Write SQL to determine role hierarchy across users.
285.  Audit password reset requests over time.
286.  Detect logins from multiple locations within short timeframe.
287.  Compare access rights between two users.
288.  Track audit log events by severity level.
289.  Identify users who changed permissions in last 7 days.
290.  Calculate average session duration per role.

* * *

### 🟥 **ETL Testing & Data Quality Assurance**

291.  Compare record counts across staging and production tables.
292.  Write SQL to detect schema drift in imported files.
293.  Track failed loads in ETL logs using SQL.
294.  Validate primary key uniqueness in raw ingested data.
295.  Write a row-level comparison between two versions of a table.
296.  How do you detect schema changes in SQL pipelines?
297.  Implement row hashing to compare large datasets efficiently.
298.  Monitor freshness of ETL data using load timestamps.
299.  Track changes in business metrics after ETL change.
300.  Identify partial or incomplete loads in partitioned tables.

* * *

Would you like:

*   📘 All 300 questions as a categorized **PDF**?
*   🧠 A self-assessment tracker or **Notion template**?
*   💡 SQL query **solutions** to any of these sets?

Let me know and I’ll create it for you.

## Prompt:
give me more without reapting above one

## Response:
Absolutely — here’s **yet another fresh batch of 100 real-world SQL interview questions**, bringing the total to **400**, **all unique** and **non-repetitive**. This set focuses on **system design with SQL**, **advanced use cases**, **event tracking**, **error-prone logic traps**, **governance**, and **data contract checks** — often explored in **senior-level interviews**, **system audits**, and **data product teams**.

* * *

### 🟢 **Tracking User Events & Analytics**

301.  How would you track a user’s complete clickstream path?
302.  Write a query to calculate time between specific events (e.g. add\_to\_cart → purchase).
303.  Identify the most common sequence of three actions by users.
304.  Track first touch and last touch attribution for a user.
305.  Write SQL to calculate session length using event timestamps.
306.  How would you identify bot-like behavior in click data?
307.  Filter event logs where users clicked the same button repeatedly within 10 seconds.
308.  Count the number of times a user performed an action after seeing a banner.
309.  Find users who performed a sequence A → B → C in order.
310.  Compare engagement metrics between users who saw an experiment and those who didn’t.

* * *

### 🟡 **Business Logic & Data Contracts**

311.  Write a constraint to enforce that end\_date must be after start\_date.
312.  Create a check that ensures salary is non-negative and within company limits.
313.  How do you enforce allowed values in a string column?
314.  Detect records violating a uniqueness constraint across multiple columns.
315.  Write SQL to identify schema mismatches in dynamic ingestion tables.
316.  Enforce business rule: users can’t be in more than one active trial at a time.
317.  Identify foreign keys pointing to non-existent parent records.
318.  Write a query to detect duplicated timestamps for the same user.
319.  Audit table where deleted rows must leave an audit trail.
320.  Validate that all customer emails follow proper syntax using regex or LIKE.

* * *

### 🔵 **Systems Thinking with SQL**

321.  How do you model many-to-many relationships in SQL schema?
322.  What are surrogate keys and why would you use them?
323.  Model a feature flag system using relational tables.
324.  Design a product subscription model with upgrades and downgrades.
325.  Write SQL to model and report on customer referral trees.
326.  Implement tag-based searching across products in SQL.
327.  Write a SQL schema to support multivariate A/B testing.
328.  How would you implement data versioning using SQL?
329.  Model and report on hierarchical department structures.
330.  Create a changelog table that captures inserts, updates, and deletes in detail.

* * *

### 🟠 **Anti-patterns & Tricky SQL Logic**

331.  What’s wrong with using SELECT \* in production?
332.  How does a GROUP BY without aggregation behave?
333.  Why is DISTINCT often misused, and how would you fix it?
334.  Explain a case where a LEFT JOIN becomes an INNER JOIN due to filters.
335.  What are the risks of filtering inside a JOIN condition?
336.  Write an example where window function is misused causing incorrect logic.
337.  What problems arise from non-deterministic ordering without ORDER BY?
338.  Detect use of inconsistent aggregation across metrics.
339.  Find cases where string-to-number implicit casting causes wrong outputs.
340.  Describe how NULL logic can break equality checks in joins.

* * *

### 🔴 **Data Governance & Lineage**

341.  How do you track downstream dependencies of a SQL view?
342.  Write SQL to identify columns with PII or sensitive data.
343.  How do you flag stale or unused tables in a warehouse?
344.  Find all views that reference a specific table or column.
345.  Write a lineage report showing upstream sources of a KPI metric.
346.  Identify duplicate datasets created by different teams.
347.  Enforce column naming conventions using metadata queries.
348.  Identify hard-coded logic in SQL scripts (e.g., fixed date ranges).
349.  Track last query time for each table to flag unused ones.
350.  Generate a data catalog with column descriptions from information schema.

* * *

### 🟣 **Performance Diagnostics**

351.  How would you detect full table scans in your warehouse?
352.  Compare performance between indexed vs non-indexed queries.
353.  What are the downsides of excessive partitioning?
354.  Write SQL to find top N largest tables in storage.
355.  Detect queries with high CPU usage in past week.
356.  Monitor query latency percentiles using a system table.
357.  Audit long-running queries by user or service.
358.  Track sudden growth in row counts over time.
359.  Profile skew in joins by inspecting distribution of join keys.
360.  Identify and reduce cross joins that are unintentional.

* * *

### 🟤 **Time-Aware Modeling**

361.  Create a slowly changing dimension model with effective/expiry dates.
362.  Track changing primary addresses while keeping history.
363.  Write SQL to identify overlapping active date ranges for users.
364.  Model “as of” reporting logic to look back as of any historical date.
365.  Find the state of a subscription table at the end of every month.
366.  Rebuild a snapshot of data using audit logs.
367.  How would you handle schema drift across time in SQL?
368.  Detect partial overlaps in time periods between two records.
369.  Write a query to compute the latest value as of a given timestamp.
370.  Model time-aware KPI trends with point-in-time accuracy.

* * *

### ⚫ **Cross-Domain Logic**

371.  Create a pricing tier structure for SaaS usage billing.
372.  Model and track shipment status across time and carriers.
373.  Write SQL to generate invoices with dynamic pricing.
374.  Design a loyalty point redemption and expiration logic in SQL.
375.  Identify drivers with idle time over X minutes between deliveries.
376.  Track active incidents in a ticketing system.
377.  Build a schema to handle multi-currency transactions.
378.  Write SQL to calculate fuel consumption per route for logistics.
379.  Detect price anomalies in product catalogs by vendor.
380.  Monitor service usage thresholds for automatic suspension.

* * *

### 🟥 **Uncommon or Niche SQL Features**

381.  How do you use `FILTER(WHERE ...)` inside aggregates?
382.  Use `QUALIFY` clause in Snowflake to filter windowed results.
383.  What does `SEMI JOIN` or `ANTI JOIN` mean in specific dialects?
384.  Write a lateral join (or CROSS APPLY) use case.
385.  Use `PIVOT` or `UNPIVOT` operations for reshaping data.
386.  Implement set-returning functions in PostgreSQL (e.g., `unnest()`).
387.  Generate dynamic SQL queries using stored procedures.
388.  Use `MERGE` for upsert logic in warehouse platforms.
389.  How do you query materialized views efficiently?
390.  Build a recursive query to compute running balance in a ledger.

* * *

### 🟧 **Mock Debugging / Fix-the-Bug Questions**

391.  This query returns wrong revenue totals — find and fix the error.
392.  The result duplicates rows unexpectedly — what’s wrong?
393.  This funnel step count is higher than step 1 — diagnose it.
394.  This rolling average is not behaving correctly — fix the logic.
395.  NULLs are interfering with sort order — how would you fix?
396.  A window function isn't resetting per partition — why?
397.  The grouping level is too granular — what caused it?
398.  Query uses subqueries but is very slow — suggest fixes.
399.  Results look correct but are outdated — identify the cause.
400.  A daily report is showing fluctuating user counts — explain possible causes.

* * *

Would you like:

*   ✅ Solutions for some or all?
*   ✅ A printable/exportable version (PDF, Excel, or Notion)?
*   ✅ Questions tailored to a specific job (e.g., data engineer, BI analyst, etc.)?

Let me know — I can generate anything from practice packs to flashcards or a quiz interface.



---
Powered by [ChatGPT Exporter](https://www.chatgptexporter.com)