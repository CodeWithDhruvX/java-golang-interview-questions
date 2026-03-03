# 🗣️ Theory — MySQL Joins & Subqueries
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Explain the different types of JOINs in MySQL."

> *"JOINs combine rows from two or more tables based on a related column. The most common ones: INNER JOIN returns only rows that have a matching record in both tables — non-matching rows are excluded entirely. LEFT JOIN returns all rows from the left table and fills in NULLs for columns from the right table where there's no match — useful when I want to preserve all records regardless of whether a match exists. RIGHT JOIN is the mirror image — all rows from the right table, NULLs for unmatched left rows. MySQL doesn't have FULL OUTER JOIN natively — I simulate it by UNION-ing a LEFT and RIGHT JOIN. CROSS JOIN produces a cartesian product — every row from table A paired with every row from table B. And SELF JOIN is when a table joins to itself — classic example is an employee table where each employee has a manager_id that references the same employees table."*

---

## Q: "What's the difference between a correlated subquery and a non-correlated subquery?"

> *"A non-correlated subquery executes once and its result is used by the outer query — it's independent. For example: SELECT name FROM employees WHERE salary > (SELECT AVG(salary) FROM employees). That inner SELECT runs once, returns one number, and the outer query uses it. A correlated subquery references columns from the outer query, so it re-executes once for EVERY row the outer query processes. For example: finding employees whose salary is above the average of their OWN department — the inner query references e.department from the outer query, so it runs once per row. Correlated subqueries can be slow on large tables because they re-execute N times. Often, I can rewrite them as a JOIN against a derived table or CTE that pre-aggregates — that's usually faster because the database can optimize set-based operations better than row-by-row loops."*

---

## Q: "When would you use EXISTS vs IN vs JOIN?"

> *"The honest answer is the MySQL query optimizer often rewrites these to the same execution plan, so the performance difference is smaller than it used to be. But conceptually: IN is clean for a small, static subquery result — WHERE id IN (1,2,3,4). EXISTS short-circuits — it stops as soon as it finds one match, which makes it better for large subqueries because it doesn't need to collect the full result set. JOIN tends to be the fastest when properly indexed because the optimizer can leverage join algorithms and index lookups efficiently. For the anti-join pattern — finding rows that DON'T match — I prefer NOT EXISTS over NOT IN because NOT IN has a serious NULL trap: if the subquery returns even one NULL value, the entire NOT IN returns no results at all. NOT EXISTS handles NULLs safely."*

---

## Q: "What is a CTE and how is it different from a subquery?"

> *"A CTE — Common Table Expression — is a named temporary result set defined with the WITH keyword at the start of a query. It behaves like a named subquery, but there are meaningful differences. Readability is the biggest one — CTEs let you break complex queries into named, logical steps that read top to bottom like a story, whereas nested subqueries become hard to follow quickly. You can define multiple CTEs and reference earlier ones in later ones. CTEs also enable recursion — a recursive CTE references itself, which is how you traverse hierarchical data like org charts or bill-of-materials trees. Regular subqueries can't do that. Performance-wise, in MySQL CTEs are typically inlined — materialized on first reference — so they run once rather than every time they're referenced. Though the MySQL optimizer can sometimes materialize a CTE differently, so always EXPLAIN complex CTE queries to verify."*

---

## Q: "How do you find the Nth highest salary?"

> *"There are a few ways. The classic approach uses DISTINCT and ORDER BY with LIMIT OFFSET — ORDER BY salary DESC, then LIMIT 1 OFFSET N-1. So the second highest would be OFFSET 1, third would be OFFSET 2. A common gotcha: if there are ties — two people with the same salary — you need to think about whether you want the Nth distinct salary value or the Nth row. For the Nth distinct value, use SELECT DISTINCT or DENSE_RANK. DENSE_RANK is the cleanest solution: wrap the employees in a subquery that adds a DENSE_RANK column ordered by salary descending, then filter the outer query WHERE rnk = N. DENSE_RANK handles ties correctly — it assigns the same rank to tied values and doesn't skip numbers. RANK skips numbers after ties, ROW_NUMBER doesn't handle ties at all — it gives a unique number to every row regardless."*

---

## Q: "What is the issue with NOT IN when the subquery can return NULLs?"

> *"This is a classic SQL trap. If your subquery returns any NULL at all, NOT IN behaves as if no rows satisfy the condition — you get zero results. The reason comes down to SQL's three-valued logic. When you say WHERE id NOT IN (1, 2, NULL), SQL evaluates each id against every value in the list. For the NULL check: is my_id NOT EQUAL to NULL? In SQL, any comparison with NULL returns UNKNOWN — not TRUE and not FALSE. And for NOT IN, if any comparison in the list is UNKNOWN, the whole expression is UNKNOWN, which means the row is filtered out. So if there's even one NULL in the subquery, nothing passes the NOT IN filter. The fix is to always add WHERE column IS NOT NULL inside the NOT IN subquery. Or better yet, switch to NOT EXISTS — it handles NULLs correctly because it's checking for row existence, not value equality."*

---

## Q: "How does a recursive CTE work for hierarchical data?"

> *"A recursive CTE has two parts separated by UNION ALL. The first part is the anchor — the base case that returns the starting rows with no recursion. The second part is the recursive member — it references the CTE itself by name and joins it to get the next level of data. MySQL executes this by first running the anchor, putting those results in a working table, then repeatedly running the recursive member using the working table as input until no new rows are produced. For an org chart, the anchor selects the CEO — the employee with no manager. The recursive step joins the employees table to the CTE results to find all direct reports of whoever is currently in the result set. With each iteration, you go one level deeper in the hierarchy. You should always have a MAXRECURSION limit to prevent infinite loops in case of circular references in the data."*

---

## Q: "What is UNION vs UNION ALL and when do you use each?"

> *"Both UNION and UNION ALL combine the result sets of two or more SELECT statements — the columns and data types must be compatible. The difference: UNION deduplicates the results — it performs an implicit DISTINCT operation which requires sorting or hashing the entire result set to find and remove duplicates. UNION ALL keeps all rows including duplicates — it just concatenates the result sets, no deduplication step. UNION ALL is faster because there's no deduplication overhead — I use it whenever I know the sets are non-overlapping or when duplicates are acceptable. A practical example: merging active and archived user records for a report — I know they're separate tables with no overlap, so UNION ALL is appropriate. I only use UNION when I genuinely need deduplication and can't do it more efficiently in another way."*

