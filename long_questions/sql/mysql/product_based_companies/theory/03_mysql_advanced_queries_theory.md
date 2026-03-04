# 🗣️ Theory — Advanced MySQL Queries
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are window functions and how do they differ from GROUP BY?"

> *"Window functions perform calculations across a set of rows related to the current row — but unlike GROUP BY, they don't collapse the result set. GROUP BY groups rows together and gives you one output row per group — you lose the individual row details. Window functions keep all rows intact and compute an additional value for each row based on a 'window' — a defined set of surrounding rows. The syntax is function() OVER (PARTITION BY ... ORDER BY ...). PARTITION BY divides rows into groups — like GROUP BY — but each row keeps its identity. ORDER BY inside OVER defines the sequence for ordered calculations like running totals or rank. You can have multiple different window specifications in the same query applying to different columns, which is impossible to express cleanly with GROUP BY. Common window functions: ROW_NUMBER for a unique sequential number, RANK and DENSE_RANK for rankings with tie handling, SUM and AVG OVER for running totals and moving averages, LAG and LEAD for accessing adjacent rows."*

---

## Q: "Explain RANK vs DENSE_RANK vs ROW_NUMBER."

> *"All three assign a number to each row within a partition ordered by some criteria — the difference is in how they handle ties. ROW_NUMBER always assigns a unique number to every row — no two rows get the same number, even if they have identical values. The assignment for tied rows is arbitrary. RANK assigns the same number to tied rows, but then skips the next numbers — so if two rows are tied for second place, they both get rank 2, and the next row gets rank 4, not 3. DENSE_RANK also assigns the same number to tied rows, but does NOT skip — so after the two rows tied at rank 2, the next row gets rank 3. In SQL interviews, the classic question is 'find the second highest salary' — use DENSE_RANK for the clean semantics: rank by salary descending, filter WHERE rnk = 2. RANK would work too but its skipping behavior can confuse. ROW_NUMBER is used when you genuinely need a unique row number — like when you need exactly one row per group, the 'top 1 per group' pattern."*

---

## Q: "How do LAG and LEAD work? What are they used for?"

> *"LAG and LEAD are window functions that let you access values from other rows relative to the current row, all within the same SELECT — without a self-join. LAG accesses a value from a previous row — N rows before the current one. LEAD accesses a value from a future row — N rows ahead. Both accept the column, the number of rows to look back or forward (default 1), and a default value if the referenced row doesn't exist — useful for the first or last rows in the partition. Classic use cases: calculating day-over-day change — subtract today's value from yesterday's using LAG. Month-over-month growth rate. Detecting consecutive events — if LAG(status) was 'inactive' and now it's 'active', that's an activation event. These used to require self-joins or correlated subqueries — which were verbose and slow. LAG and LEAD express the same logic in a single, non-joining, clean expression that the database can compute in one pass."*

---

## Q: "How do you do a pivot in MySQL? MySQL doesn't have a PIVOT keyword."

> *"MySQL doesn't have a native PIVOT statement, so you implement pivoting using conditional aggregation — wrapping CASE expressions inside aggregate functions. The logic is: for each row, the CASE expression returns the value if the row belongs to that column's category, otherwise NULL. The aggregate function — SUM or MAX — then collapses those into one value per group. So for a monthly sales pivot, you write SELECT product, SUM(CASE WHEN month = 'Jan' THEN sales END) AS Jan, SUM(CASE WHEN month = 'Feb' THEN sales END) AS Feb, and so on, then GROUP BY product. For static pivot with known columns, this is clear and fast. For dynamic pivot — where the columns are discovered at runtime from the data — you need to use GROUP_CONCAT to build a SQL string dynamically and then execute it with PREPARE and EXECUTE. Dynamic pivots are more complex and require careful handling of SQL injection via whitelisting, since column names can't be parameterized."*

---

## Q: "How does JSON support work in MySQL 5.7+?"

> *"MySQL 5.7 introduced a native JSON data type that stores JSON documents in a binary-optimized format — not as plain text — which makes reading individual values fast without parsing the entire string. The JSON path extraction operator -> is the primary way to query into JSON: table.column->'$.field' returns the value with JSON quoting, and ->> strips the quotes. You can also use JSON_EXTRACT with explicit path strings. For modification, JSON_SET adds or replaces a key, JSON_INSERT adds only if the key doesn't exist, JSON_REMOVE deletes a key. For arrays, JSON_ARRAY_APPEND and JSON_CONTAINS are useful. You can index into JSON using generated columns — you create a virtual generated column that extracts a specific JSON path, and then create a regular index on that generated column. This gives you index performance for frequently queried JSON fields without denormalizing the schema. The tradeoff of JSON vs relational: JSON gives flexibility for variable schemas but sacrifices the ability to do set-based queries, enforce referential integrity, and index arbitrary fields efficiently."*

---

## Q: "How does a recursive CTE work under the hood?"

> *"A recursive CTE has an anchor member and a recursive member joined by UNION ALL. MySQL executes it iteratively. First, it runs the anchor — the base case — and puts those results into a working table. Then it runs the recursive member, which references the CTE name — and crucially, it reads FROM the working table in this iteration. The new rows produced are added to the result set and become the new working table for the next iteration. If the recursive member produces any new rows, the cycle repeats. When the recursive member produces zero rows, recursion stops. For an org chart traversal, iteration 1 gives you the CEO, iteration 2 gives you their direct reports, iteration 3 gives you those people's direct reports, and so on until no more employees are found. There's always a max_recursion limit — in MySQL controlled by cte_max_recursion_depth, default 1000 — to prevent infinite loops if there's a cycle in the data. You can also add a WHERE depth < 10 condition in the recursive member to manually limit depth."*

---

## Q: "What is the ON DUPLICATE KEY UPDATE pattern and when do you use it?"

> *"ON DUPLICATE KEY UPDATE is MySQL's upsert mechanism — insert a row, and if it would violate a PRIMARY KEY or UNIQUE constraint, instead of failing, UPDATE the existing row. It's a single atomic operation — no separate SELECT to check existence, no race condition between check and insert. The pattern is very useful for: incrementing counters — insert a row with count=1 or increment count on duplicate. Maintaining last-seen timestamps — insert a user event or update the timestamp if the user already has one. Upserting configuration or settings. The VALUES() function inside the UPDATE clause refers to the value you were trying to INSERT — so SET count = count + VALUES(count) means 'add the amount I tried to insert to the existing count'. A common gotcha: if the table has multiple unique keys and your row matches more than one, the behavior can be surprising — it depends on which unique key MySQL detects first. So design your schema to avoid having multiple unique key conflicts possible in one INSERT."*

