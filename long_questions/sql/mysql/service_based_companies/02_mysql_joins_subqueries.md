# 🔗 02 — MySQL Joins & Subqueries
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

> JOINs and Subqueries are tested at **every service company SQL round** — master these patterns.

---

## 🔑 Must-Know Topics
- INNER, LEFT, RIGHT, CROSS, SELF JOINs
- Anti-join pattern (LEFT JOIN + WHERE IS NULL)
- Subqueries: correlated vs non-correlated
- Scalar, row, and table subqueries
- EXISTS vs IN vs JOIN
- CTEs (WITH clause) — simple & recursive
- UNION, UNION ALL, INTERSECT, EXCEPT

---

## ❓ Most Asked Questions

### Q1. Write a query to get all employees and their manager names.

```sql
-- SELF JOIN — joining a table to itself
CREATE TABLE employees (
    id         INT PRIMARY KEY,
    name       VARCHAR(100),
    manager_id INT,
    department VARCHAR(50),
    salary     DECIMAL(10,2)
);

-- Get employee + their manager's name
SELECT
    e.name        AS employee,
    e.department,
    m.name        AS manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id  -- LEFT JOIN includes top-level (CEO, no manager)
ORDER BY e.department, e.name;
```

---

### Q2. Find employees in the same department earning above average salary.

```sql
-- Correlated subquery — re-executes for each row
SELECT e.name, e.department, e.salary
FROM employees e
WHERE salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department = e.department  -- correlated: references outer e.department
)
ORDER BY e.department, e.salary DESC;

-- Equivalent with JOIN + derived table (often faster)
SELECT e.name, e.department, e.salary
FROM employees e
JOIN (
    SELECT department, AVG(salary) AS avg_sal
    FROM employees
    GROUP BY department
) dept_avg ON e.department = dept_avg.department
WHERE e.salary > dept_avg.avg_sal;
```

---

### Q3. What is the difference between EXISTS, IN, and JOIN for filtering?

```sql
-- IN — good when subquery returns small, static set
SELECT name FROM customers
WHERE id IN (SELECT customer_id FROM orders WHERE total > 1000);

-- EXISTS — short-circuits (stops at first match), better for large subqueries
SELECT name FROM customers c
WHERE EXISTS (
    SELECT 1 FROM orders o
    WHERE o.customer_id = c.id AND o.total > 1000
);

-- JOIN — often fastest, optimizer can use indexes
SELECT DISTINCT c.name
FROM customers c
JOIN orders o ON c.id = o.customer_id
WHERE o.total > 1000;

-- Anti-join (NOT EXISTS) — find customers with NO orders > 1000
SELECT name FROM customers c
WHERE NOT EXISTS (
    SELECT 1 FROM orders o
    WHERE o.customer_id = c.id AND o.total > 1000
);
```

> **Rule of thumb:** JOIN > EXISTS > IN for performance, but optimizer often rewrites them anyway. Use `EXPLAIN` to verify.

---

### Q4. Write a query using CTE (WITH clause).

```sql
-- Simple CTE — improves readability
WITH high_earners AS (
    SELECT id, name, salary, department
    FROM employees
    WHERE salary > 80000
),
dept_stats AS (
    SELECT department, COUNT(*) AS cnt, AVG(salary) AS avg_sal
    FROM high_earners
    GROUP BY department
)
SELECT h.name, h.salary, d.cnt AS dept_high_earner_count
FROM high_earners h
JOIN dept_stats d ON h.department = d.department
ORDER BY h.salary DESC;
```

```sql
-- Multiple CTEs chained
WITH monthly_sales AS (
    SELECT
        DATE_FORMAT(order_date, '%Y-%m') AS month,
        SUM(total)                        AS revenue
    FROM orders
    GROUP BY month
),
monthly_ranked AS (
    SELECT *, RANK() OVER (ORDER BY revenue DESC) AS rnk
    FROM monthly_sales
)
SELECT month, revenue
FROM monthly_ranked
WHERE rnk <= 3;  -- top 3 months by revenue
```

---

### Q5. Write a recursive CTE — e.g., employee hierarchy.

```sql
-- Recursive CTE: traverse manager → employee tree
WITH RECURSIVE org_chart AS (
    -- Anchor: start with CEO (no manager)
    SELECT id, name, manager_id, 0 AS level, CAST(name AS CHAR(500)) AS path
    FROM employees
    WHERE manager_id IS NULL

    UNION ALL

    -- Recursive: join employees to their managers
    SELECT e.id, e.name, e.manager_id,
           oc.level + 1,
           CONCAT(oc.path, ' → ', e.name)
    FROM employees e
    JOIN org_chart oc ON e.manager_id = oc.id
)
SELECT id, name, level, path
FROM org_chart
ORDER BY path;
```

```sql
-- Recursive CTE: generate a date series
WITH RECURSIVE dates AS (
    SELECT '2024-01-01' AS dt
    UNION ALL
    SELECT DATE_ADD(dt, INTERVAL 1 DAY)
    FROM dates
    WHERE dt < '2024-01-31'
)
SELECT dt FROM dates;
```

---

### Q6. What is CROSS JOIN and when do you use it?

```sql
-- CROSS JOIN — cartesian product: every row × every row
SELECT c.color, s.size
FROM colors c
CROSS JOIN sizes s;
-- If colors has 3 rows and sizes has 4, result: 12 rows

-- Use case: generate all combinations (e.g., product variants)
CREATE TABLE colors (color VARCHAR(20));
CREATE TABLE sizes  (size  VARCHAR(10));
INSERT INTO colors VALUES ('Red'), ('Blue'), ('Green');
INSERT INTO sizes  VALUES ('S'), ('M'), ('L'), ('XL');

SELECT CONCAT(color, '-', size) AS variant
FROM colors CROSS JOIN sizes
ORDER BY color, size;
-- Red-L, Red-M, Red-S, Red-XL, Blue-L, ...
```

---

### Q7. Find the Nth highest salary.

```sql
-- Method 1: Subquery (classic, interview standard)
-- Find 2nd highest salary
SELECT MAX(salary) AS second_highest
FROM employees
WHERE salary < (SELECT MAX(salary) FROM employees);

-- Generalizable Nth highest
SELECT DISTINCT salary
FROM employees
ORDER BY salary DESC
LIMIT 1 OFFSET 2;  -- OFFSET n-1 (0-indexed), so OFFSET 2 = 3rd highest

-- Method 2: Dense Rank (handles ties correctly)
SELECT salary
FROM (
    SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) AS rnk
    FROM employees
) ranked
WHERE rnk = 2;
```

---

### Q8. Find duplicate records in a table.

```sql
-- Find emails that appear more than once
SELECT email, COUNT(*) AS occurrences
FROM users
GROUP BY email
HAVING COUNT(*) > 1;

-- Show full rows for duplicates
SELECT u.*
FROM users u
JOIN (
    SELECT email
    FROM users
    GROUP BY email
    HAVING COUNT(*) > 1
) dups ON u.email = dups.email
ORDER BY u.email;

-- Delete duplicates, keep lowest id
DELETE u1
FROM users u1
JOIN users u2 ON u1.email = u2.email AND u1.id > u2.id;
```

---

### Q9. Get customers who placed orders in every month of 2024.

```sql
-- Customers who appear in all 12 months of 2024
SELECT customer_id
FROM orders
WHERE YEAR(order_date) = 2024
GROUP BY customer_id
HAVING COUNT(DISTINCT MONTH(order_date)) = 12;

-- More robust: divide and conquer with EXISTS
SELECT DISTINCT o.customer_id
FROM orders o
WHERE YEAR(o.order_date) = 2024
  AND NOT EXISTS (
      SELECT m.month_num
      FROM (
          SELECT 1 AS month_num UNION SELECT 2 UNION SELECT 3 UNION
          SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION
          SELECT 8 UNION SELECT 9 UNION SELECT 10 UNION SELECT 11 UNION SELECT 12
      ) m
      WHERE NOT EXISTS (
          SELECT 1 FROM orders o2
          WHERE o2.customer_id = o.customer_id
            AND YEAR(o2.order_date) = 2024
            AND MONTH(o2.order_date) = m.month_num
      )
  );
```

---

### Q10. What is UNION vs UNION ALL?

```sql
-- UNION — combines result sets, REMOVES duplicates (slower, sorts internally)
SELECT name FROM customers
UNION
SELECT name FROM vendors;

-- UNION ALL — combines result sets, KEEPS ALL rows (faster, no dedup)
SELECT name FROM customers
UNION ALL
SELECT name FROM vendors;

-- Practical example: merge active and inactive users into one report
SELECT id, name, 'active'   AS status FROM active_users
UNION ALL
SELECT id, name, 'inactive' AS status FROM archived_users
ORDER BY name;

-- INTERSECT simulation (MySQL doesn't have INTERSECT natively)
SELECT name FROM table_a
WHERE name IN (SELECT name FROM table_b);

-- EXCEPT simulation (rows in A but not B)
SELECT name FROM table_a
WHERE name NOT IN (SELECT name FROM table_b WHERE name IS NOT NULL);
```

---

### Q11. Write a query to find departments with no employees.

```sql
-- Anti-join pattern with LEFT JOIN
SELECT d.department_name
FROM departments d
LEFT JOIN employees e ON d.id = e.department_id
WHERE e.id IS NULL;

-- Alternative: NOT IN (watch out for NULLs!)
SELECT department_name
FROM departments
WHERE id NOT IN (
    SELECT DISTINCT department_id
    FROM employees
    WHERE department_id IS NOT NULL  -- CRITICAL: NOT IN with NULL returns empty!
);

-- Alternative: NOT EXISTS (handles NULLs safely)
SELECT department_name
FROM departments d
WHERE NOT EXISTS (
    SELECT 1 FROM employees e WHERE e.department_id = d.id
);
```

---

### Q12. Write a query showing running total of sales per day.

```sql
-- Without window functions (MySQL < 8.0) — self-join approach
SELECT
    a.sale_date,
    a.daily_total,
    SUM(b.daily_total) AS running_total
FROM (
    SELECT DATE(order_date) AS sale_date, SUM(total) AS daily_total
    FROM orders GROUP BY sale_date
) a
JOIN (
    SELECT DATE(order_date) AS sale_date, SUM(total) AS daily_total
    FROM orders GROUP BY sale_date
) b ON b.sale_date <= a.sale_date
GROUP BY a.sale_date
ORDER BY a.sale_date;

-- With window functions (MySQL 8.0+) — cleaner and faster
SELECT
    DATE(order_date) AS sale_date,
    SUM(total)       AS daily_total,
    SUM(SUM(total)) OVER (ORDER BY DATE(order_date)) AS running_total
FROM orders
GROUP BY DATE(order_date)
ORDER BY sale_date;
```

---

### Q13. What is a derived table (inline view)?

```sql
-- Derived table: a subquery in the FROM clause
-- Must be aliased
SELECT dept_summary.department, dept_summary.avg_salary
FROM (
    SELECT department, AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department
) AS dept_summary                        -- ← alias is REQUIRED
WHERE dept_summary.avg_salary > 60000
ORDER BY dept_summary.avg_salary DESC;

-- Derived table vs CTE — same result, different readability:
-- CTE is preferred for complex multi-step logic
WITH dept_summary AS (
    SELECT department, AVG(salary) AS avg_salary
    FROM employees GROUP BY department
)
SELECT * FROM dept_summary WHERE avg_salary > 60000;
```

---

### Q14. How do you handle many-to-many relationships with JOINs?

```sql
-- Three tables: students, courses, enrollments (junction)
CREATE TABLE students    (id INT PRIMARY KEY, name VARCHAR(100));
CREATE TABLE courses     (id INT PRIMARY KEY, title VARCHAR(200), credits INT);
CREATE TABLE enrollments (
    student_id INT, course_id INT, grade CHAR(2),
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id)  REFERENCES courses(id)
);

-- All students with their courses
SELECT s.name, c.title, e.grade
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses     c ON c.id = e.course_id
ORDER BY s.name, c.title;

-- Students enrolled in BOTH 'Math' AND 'Physics'
SELECT s.name
FROM students s
JOIN enrollments e1 ON s.id = e1.student_id
JOIN courses c1 ON e1.course_id = c1.id AND c1.title = 'Math'
JOIN enrollments e2 ON s.id = e2.student_id
JOIN courses c2 ON e2.course_id = c2.id AND c2.title = 'Physics';
```

