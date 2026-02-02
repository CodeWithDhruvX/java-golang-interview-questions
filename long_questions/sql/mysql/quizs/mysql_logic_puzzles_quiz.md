# MySQL Logic Puzzles & Algorithmic Quiz

## 1. Questions

### Q1: Consecutive Logins (Gaps and Islands)
**Scenario:** Table `Logins` has columns `id` (int) and `login_date` (date).
**Question:** Write a query to find all users who have logged in for at least **3 consecutive days**.

### Q2: The N-th Highest Salary
**Scenario:** Table `Employee` has `id` and `salary`.
**Question:** Write a query to find the **N-th** highest distinct salary. (Assume N is a variable, e.g., N=3). If there is no N-th highest salary, return `NULL`.

### Q3: Swapping Values (The "Sex" Swap)
**Scenario:** Table `salary` has columns `id`, `name`, `sex` ('m' or 'f'), and `salary`.
**Question:** Write a single `UPDATE` statement to swap all 'f' and 'm' values (i.e., change all 'f' to 'm' and vice versa) with a single statement and **no intermediate temporary tables**.

### Q4: Calculating Median
**Scenario:** Table `Numbers` has a column `val` (int).
**Question:** Write a query to calculate the **median** of the `val` column. Remember that for an even number of rows, the median is the average of the two middle numbers.

### Q5: Recursive Hierarchy (Tree Traversal)
**Scenario:** Table `Employees` has `id`, `name`, and `manager_id`.
**Question:** Write a query (using a Recursive CTE) to find the entire reporting hierarchy (subordinates of subordinates) for a specific Manager ID (e.g., ID = 1).

---

## 2. Answers & Explanations

### Answer 1: Consecutive Logins
```sql
WITH LOGIN_GROUPS AS (
    SELECT 
        id, 
        login_date,
        DATE_SUB(login_date, INTERVAL ROW_NUMBER() OVER(PARTITION BY id ORDER BY login_date) DAY) as grp
    FROM Logins
)
SELECT id, MIN(login_date) as start_date, COUNT(*) as consecutive_days
FROM LOGIN_GROUPS
GROUP BY id, grp
HAVING COUNT(*) >= 3;
```
**Explanation:**
This is the classic "Gaps and Islands" problem.
*   If dates are consecutive (Jan 1, Jan 2, Jan 3), subtracting their Row Number (1, 2, 3) from them results in the **same logical date** (The "Group" or "Island").
*   We then grouping by this "Group" date allows us to count the size of the consecutive island.

### Answer 2: N-th Highest Salary
```sql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N = N - 1;
  RETURN (
      SELECT DISTINCT salary 
      FROM Employee 
      ORDER BY salary DESC 
      LIMIT 1 OFFSET N
  );
END
```
**Explanation:**
*   We use `LIMIT 1 OFFSET (N-1)`.
*   Note: `OFFSET` starts at 0, so for the 3rd highest, we skip 2 (`N-1`).
*   **Alternative (Window Function)**: `SELECT salary FROM (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) as rnk FROM Employee) t WHERE rnk = N LIMIT 1`.

### Answer 3: Swapping Values (CASE WHEN)
```sql
UPDATE salary
SET sex = CASE sex
    WHEN 'm' THEN 'f'
    ELSE 'm'
END;
```
**Explanation:**
A `CASE` statement inside an `UPDATE` allows conditional logic per row. This operation is atomic and flips the values in a single pass without needing a temporary variable.

### Answer 4: Calculate Median
```sql
SELECT AVG(val) as median_val
FROM (
    SELECT val,
           ROW_NUMBER() OVER (ORDER BY val) as row_id,
           COUNT(*) OVER () as total_rows
    FROM Numbers
) t
WHERE row_id IN (FLOOR((total_rows + 1) / 2), CEIL((total_rows + 1) / 2));
```
**Explanation:**
1.  We assign a sequential `row_id` to sorted values.
2.  We find the total count of rows `total_rows`.
3.  **Mathematical Logic**:
    *   If count is odd (e.g., 5), logic selects row 3 (`(5+1)/2 = 3` and `3`). Average of [3] is value at 3.
    *   If count is even (e.g., 4), logic selects rows 2 and 3 (`2.5` floors to 2, ceils to 3). Average of [2, 3] gives the median.

### Answer 5: Recursive CTE
```sql
WITH RECURSIVE OrgChart AS (
    -- Anchor Member: Start with the specific manager
    SELECT id, name, manager_id, 1 as level
    FROM Employees
    WHERE id = 1
    
    UNION ALL
    
    -- Recursive Member: Join output of previous step with original table
    SELECT e.id, e.name, e.manager_id, o.level + 1
    FROM Employees e
    INNER JOIN OrgChart o ON e.manager_id = o.id
)
SELECT * FROM OrgChart;
```
**Explanation:**
1.  **Anchor**: Selects the root node (Manager ID 1).
2.  **Recursive Step**: Joins the `Employees` table with the *result so far* (`OrgChart`) to find employees whose `manager_id` matches the IDs found in the previous step.
3.  This repeats until no new rows are returned.
