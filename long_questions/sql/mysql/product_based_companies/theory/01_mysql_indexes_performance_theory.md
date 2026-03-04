# 🗣️ Theory — MySQL Indexes & Performance
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does a B-Tree index work in MySQL?"

> *"A B-Tree — balanced tree — index organizes values in a sorted, hierarchical structure. Internal nodes hold separator keys for navigation, and leaf nodes hold the actual indexed values plus pointers to the corresponding row data. In InnoDB, for secondary indexes, those leaf nodes hold the primary key value rather than a direct row pointer — so any secondary index lookup ultimately walks two trees: first the secondary B-Tree to find the primary key, then the primary key B-Tree to find the actual row. This is called a bookmark lookup. The key advantage of B-Tree over a hash index is that it supports range queries — BETWEEN, greater than, less than, LIKE 'prefix%', and ORDER BY — because values are stored in sorted order. Hash indexes are O(1) for exact lookups but useless for ranges."*

---

## Q: "What is a clustered vs non-clustered index?"

> *"In InnoDB, the primary key is always the clustered index. That means the actual row data IS physically organized and stored within the B-Tree leaves — the table data and the primary key index are one combined structure. When you do a primary key lookup, you jump directly to the data in one B-Tree traversal. All other indexes — email, department, whatever — are non-clustered secondary indexes. Their leaf nodes don't contain row data; they contain the primary key value. So a lookup via a secondary index requires two tree traversals: one in the secondary index to find the PK, then one in the clustered index to get the actual row. This design has an implication: keep your primary key small and compact — like an integer — because every single secondary index stores the PK value in its leaves. A large VARCHAR primary key makes every secondary index larger."*

---

## Q: "What is a covering index and why does it matter?"

> *"A covering index contains all the columns a query needs to answer — the SELECT columns, the WHERE columns, and the ORDER BY columns — so MySQL can satisfy the entire query just by reading the index, never touching the main table. In EXPLAIN output, you'll see Extra: 'Using index' — that's the signal that a covering index is being used. The performance improvement is dramatic because index reads are sequential and compact — much smaller data volume than full row reads. I design covering indexes specifically for my most critical query patterns. For example, if my most common query is SELECT name, salary FROM employees WHERE department = 'Engineering' ORDER BY salary, I'd create a composite index on (department, salary, name) — department is the filter, salary is the sort, name is the projected column — and the entire query resolves without ever touching the table rows."*

---

## Q: "Explain the leftmost prefix rule."

> *"The leftmost prefix rule says that a composite index on columns A, B, C can only be used starting from the leftmost column — you cannot skip columns. So the index is useful for queries that filter on A alone, on A and B, or on A and B and C. But if you filter only on B or only on C, MySQL cannot use that index — it would have to do a full table scan. The columns in the index must be applicable from left to right without gaps in the WHERE clause. There's a nuance though: you CAN skip columns in the WHERE clause for the trailing columns — the index will be partially used for the leading columns that match, and then MySQL applies additional filtering. For ORDER BY, the same rule applies — and if your ORDER BY sequence doesn't match the index column order, MySQL will do a filesort even if the WHERE clause used the index."*

---

## Q: "How do you read EXPLAIN output?"

> *"EXPLAIN is my first tool whenever I'm diagnosing a slow query. The most important column is 'type' — the access method. From best to worst it goes: system — single row table, const — one row via primary key, eq_ref — one row per outer row via unique index, ref — many rows via non-unique index, range — index range scan, index — full index scan, and finally ALL — full table scan. I want to see system, const, eq_ref, ref, or range. Seeing ALL or index on large tables means something is wrong. The 'key' column tells me which index MySQL actually chose. The 'rows' column is an estimate of how many rows MySQL thinks it will examine — the lower, the better. 'Extra' gives additional signals: 'Using index' is great — covering index. 'Using filesort' means MySQL needs to sort results in memory or on disk — I should add an appropriate index. 'Using temporary' is expensive — it means MySQL is creating an in-memory temp table, usually for GROUP BY or DISTINCT operations."*

---

## Q: "Why does a function on an indexed column break index usage?"

> *"When you wrap a column in a function inside a WHERE clause — like WHERE YEAR(created_at) = 2024 — MySQL cannot use the index on created_at. The index is built on the raw column values, not the function output. MySQL would have to call YEAR() for every row in the table and then check if it equals 2024 — that's effectively a full table scan. The fix is to rewrite the condition to use the raw column in a range expression: WHERE created_at >= '2024-01-01' AND created_at < '2025-01-01'. Now MySQL can do an index range scan on the raw datetime values, which is fast. This is one of the most common performance traps I see. Similar issue with: WHERE LOWER(email) = 'user@example.com' — breaks the index on email. Fix: use a case-insensitive collation or store the email pre-lowercased. MySQL 8.0 introduced functional indexes where you can actually index an expression — but that's an advanced feature."*

---

## Q: "What does 'Using filesort' in EXPLAIN mean? How do you fix it?"

> *"'Using filesort' in EXPLAIN means MySQL cannot satisfy the ORDER BY clause using an index — instead, it has to sort the result set in memory or on disk after fetching the data. Despite the word 'file', this can happen entirely in memory — it just means a sort operation is needed. It's expensive for large result sets. The fix is almost always to add or tune an index that covers both the WHERE condition and the ORDER BY columns in the right order. For example, if your query is WHERE department = 'Engineering' ORDER BY salary DESC, an index on (department, salary) would let MySQL retrieve rows already in sorted order — no separate sort step needed, no filesort. The WHERE column goes first in the index because it's used for equality filtering, and the ORDER BY column goes second. The sort direction in the index needs to match too — though MySQL can traverse a B-Tree in either direction, so ASC vs DESC is usually handled."*

---

## Q: "What is index cardinality and why does the optimizer care about it?"

> *"Cardinality is the number of distinct values in an indexed column. High cardinality means the index is selective — each value identifies a small fraction of rows — so the index is very useful for filtering. Low cardinality means many rows share the same value, so the index isn't very selective. The optimizer uses cardinality statistics to decide whether using an index is actually cheaper than a full table scan. If a column has only two values — like gender with 'M' and 'F' — and you filter for 'M', using an index might not make sense if 'M' is 50% of rows, because the optimizer has to hit the index and then fetch half the table anyway. A sequential scan could actually be cheaper. You update cardinality statistics by running ANALYZE TABLE — this is why I run ANALYZE TABLE after large bulk data operations, because stale stats can cause the optimizer to make poor index choices."*

