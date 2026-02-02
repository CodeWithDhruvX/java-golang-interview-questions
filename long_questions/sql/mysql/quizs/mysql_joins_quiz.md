# MySQL Joins Interview Questions

## 1. Schema Description
Assumed simple schema for examples:
*   `Employees` (id, name, department_id, manager_id)
*   `Departments` (id, department_name)

## 2. Questions

### Q1: The "Unmatched" Left Join
**Scenario:** You need to find all **Departments** that have **NO** employees assigned to them.
**Question:** Write a query using a `LEFT JOIN` to achieve this.

### Q2: Self Join Hierarchy
**Scenario:** The `Employees` table has a `manager_id` column which refers to the `id` of another employee in the same table.
**Question:** Write a query to list every employee's name alongside their manager's name.

### Q3: Cross Join Use Case
**Scenario:** You have a table of `Colors` (Red, Blue) and a table of `Sizes` (S, M, L).
**Question:** Write a query to generate all possible combinations of Color and Size (Cartesian Product).

### Q4: Inner Join vs Implicit Join
**Question:** What is the difference between writing `SELECT * FROM A JOIN B ON A.id = B.id` versus `SELECT * FROM A, B WHERE A.id = B.id`? Which is preferred and why?

### Q5: Finding Duplicates with Self Join
**Scenario:** You have a table `Emails` (id, email).
**Question:** Write a query using a Self Join to find all email addresses that appear more than once.

---

## 3. Answers & Explanations

### Answer 1: Departments with No Employees
```sql
SELECT d.department_name
FROM Departments d
LEFT JOIN Employees e ON d.id = e.department_id
WHERE e.id IS NULL;
```
**Explanation:**
A `LEFT JOIN` returns all rows from the left table (`Departments`). If there is no match in the right table (`Employees`), the columns from the right table will be `NULL`. Filtering for `e.id IS NULL` isolates these unmatched rows.

### Answer 2: Employee-Manager Hierarchy
```sql
SELECT 
    e.name AS Employee, 
    m.name AS Manager
FROM Employees e
LEFT JOIN Employees m ON e.manager_id = m.id;
```
**Explanation:**
This is a **Self Join**. We treat the `Employees` table as two separate entities: one representing employees (`e`) and one representing managers (`m`). We join them where the employee's `manager_id` matches the manager's `id`. A `LEFT JOIN` is safer here in case an employee has no manager (e.g., the CEO).

### Answer 3: Color-Size Combinations
```sql
SELECT c.Color, s.Size
FROM Colors c
CROSS JOIN Sizes s;
```
**Explanation:**
A `CROSS JOIN` produces a Cartesian product, meaning every row from the first table is paired with every row from the second table. If `Colors` has 2 rows and `Sizes` has 3, the result will have $2 \times 3 = 6$ rows.

### Answer 4: Inner vs Implicit
**Answer:**
*   **Explicit Syntax (`JOIN ... ON`)**: `SELECT * FROM A JOIN B ON A.id = B.id`
*   **Implicit Syntax (Comma style)**: `SELECT * FROM A, B WHERE A.id = B.id`
**Preference:** The **Explicit `JOIN` syntax is largely preferred**.
**Why?**
1.  **Readability**: It clearly separates join logic (`ON`) from filtering logic (`WHERE`).
2.  **Mistake Prevention**: In implicit syntax, forgetting the `WHERE` clause accidentally results in a massive Cross Join (Cartesian Product).

### Answer 5: Finding Duplicates via Self Join
```sql
SELECT DISTINCT e1.email
FROM Emails e1
JOIN Emails e2 ON e1.email = e2.email AND e1.id <> e2.id;
```
**Explanation:**
We join the table to itself matching on the `email` column, but ensuring we are looking at different rows (`e1.id <> e2.id`). If a match is found, it means the email exists in another row. *Note: Using `GROUP BY email HAVING COUNT(*) > 1` is usually the more standard way to do this, but the Self Join approach demonstrates the concept.*
