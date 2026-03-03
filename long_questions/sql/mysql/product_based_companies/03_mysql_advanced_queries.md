# 🪟 03 — Advanced MySQL Queries
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

> Window functions and advanced SQL — the **key differentiator** for senior roles at **Google, Meta, Atlassian, Dunzo, Cred**.

---

## 🔑 Must-Know Topics
- Window functions: ROW_NUMBER, RANK, DENSE_RANK, NTILE
- Analytic aggregates: SUM/AVG/COUNT OVER partition
- LAG, LEAD — accessing previous/next row values
- FIRST_VALUE, LAST_VALUE, NTH_VALUE
- Recursive CTEs for hierarchical data
- JSON functions (MySQL 5.7+)
- Pivot/unpivot patterns
- Dynamic SQL with PREPARE/EXECUTE

---

## ❓ Most Asked Questions

### Q1. What are window functions? Show ROW_NUMBER, RANK, DENSE_RANK.

```sql
-- Window functions: perform calculations ACROSS a set of rows related to the current row
-- Unlike GROUP BY, they do NOT collapse rows

SELECT
    name,
    department,
    salary,
    ROW_NUMBER()  OVER (PARTITION BY department ORDER BY salary DESC) AS row_num,
    RANK()        OVER (PARTITION BY department ORDER BY salary DESC) AS rnk,
    DENSE_RANK()  OVER (PARTITION BY department ORDER BY salary DESC) AS dense_rnk
FROM employees;

-- Output example (Engineering dept):
-- name     | salary | row_num | rnk | dense_rnk
-- Alice    | 120000 |    1    |  1  |     1
-- Bob      | 100000 |    2    |  2  |     2
-- Charlie  | 100000 |    3    |  2  |     2  (tie!)
-- Dave     |  90000 |    4    |  4  |     3  ← RANK skips 3, DENSE_RANK doesn't

-- Get top 1 employee per department (no DISTINCT needed)
SELECT * FROM (
    SELECT *,
           ROW_NUMBER() OVER (PARTITION BY department ORDER BY salary DESC) AS rnk
    FROM employees
) ranked
WHERE rnk = 1;
```

---

### Q2. How do you use LAG and LEAD to compare with previous/next rows?

```sql
-- LAG: access value from N rows BEFORE current row
-- LEAD: access value from N rows AFTER current row

SELECT
    sale_date,
    revenue,
    LAG(revenue, 1, 0)  OVER (ORDER BY sale_date) AS prev_day_revenue,
    LEAD(revenue, 1, 0) OVER (ORDER BY sale_date) AS next_day_revenue,
    revenue - LAG(revenue, 1, 0) OVER (ORDER BY sale_date) AS day_over_day_change,
    ROUND(
        (revenue - LAG(revenue) OVER (ORDER BY sale_date)) /
        NULLIF(LAG(revenue) OVER (ORDER BY sale_date), 0) * 100, 2
    ) AS pct_change
FROM daily_sales
ORDER BY sale_date;

-- Monthly sales growth per product
SELECT
    product_id,
    month,
    monthly_sales,
    LAG(monthly_sales) OVER (PARTITION BY product_id ORDER BY month) AS prev_month,
    ROUND(
        (monthly_sales - LAG(monthly_sales) OVER (PARTITION BY product_id ORDER BY month)) /
        NULLIF(LAG(monthly_sales) OVER (PARTITION BY product_id ORDER BY month), 0) * 100, 1
    ) AS growth_pct
FROM product_monthly_sales;
```

---

### Q3. How do you calculate running totals and moving averages?

```sql
-- Running total (cumulative sum)
SELECT
    sale_date,
    daily_revenue,
    SUM(daily_revenue) OVER (ORDER BY sale_date
        ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS running_total
FROM daily_sales;

-- 7-day moving average
SELECT
    sale_date,
    daily_revenue,
    AVG(daily_revenue) OVER (ORDER BY sale_date
        ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) AS moving_avg_7d
FROM daily_sales;

-- Running total reset per partition (per department)
SELECT
    department, sale_date, amount,
    SUM(amount) OVER (
        PARTITION BY department     -- reset per department
        ORDER BY sale_date
        ROWS UNBOUNDED PRECEDING    -- from start to current row
    ) AS dept_running_total
FROM sales;

-- RANGE vs ROWS in window frames:
-- ROWS: physical rows (row count)
-- RANGE: logical range based on ORDER BY value (handles ties differently)
```

---

### Q4. How do you use NTILE for quartile/percentile bucketing?

```sql
-- NTILE(n): divide rows into n equal buckets, assign bucket number
SELECT
    name,
    salary,
    NTILE(4) OVER (ORDER BY salary) AS salary_quartile
FROM employees;
-- Q1=bottom 25%, Q2=25-50%, Q3=50-75%, Q4=top 25%

-- Percentile rank
SELECT
    name, salary,
    PERCENT_RANK() OVER (ORDER BY salary) AS pct_rank,       -- 0 to 1
    CUME_DIST()    OVER (ORDER BY salary) AS cumulative_dist  -- 0 to 1
FROM employees;

-- Find median (50th percentile)
SELECT AVG(salary) AS median_salary
FROM (
    SELECT salary,
           ROW_NUMBER() OVER (ORDER BY salary) AS rn,
           COUNT(*) OVER () AS total
    FROM employees
) t
WHERE rn IN (FLOOR((total+1)/2), CEIL((total+1)/2));
```

---

### Q5. How do you pivot data in MySQL (rows to columns)?

```sql
-- Source data: product | month | sales
-- Goal: one row per product, columns per month

-- Static pivot using conditional aggregation
SELECT
    product_name,
    SUM(CASE WHEN month = 'Jan' THEN sales ELSE 0 END) AS Jan,
    SUM(CASE WHEN month = 'Feb' THEN sales ELSE 0 END) AS Feb,
    SUM(CASE WHEN month = 'Mar' THEN sales ELSE 0 END) AS Mar,
    SUM(CASE WHEN month = 'Apr' THEN sales ELSE 0 END) AS Apr,
    SUM(sales) AS total_year
FROM monthly_product_sales
WHERE year = 2024
GROUP BY product_name
ORDER BY total_year DESC;

-- Dynamic pivot using GROUP_CONCAT + PREPARE (when columns are unknown)
SET @sql = NULL;
SELECT GROUP_CONCAT(DISTINCT
    CONCAT('SUM(CASE WHEN month=''', month, ''' THEN sales ELSE 0 END) AS `', month, '`')
    ORDER BY month
) INTO @sql
FROM monthly_product_sales WHERE year = 2024;

SET @sql = CONCAT('SELECT product_name, ', @sql, ' FROM monthly_product_sales WHERE year=2024 GROUP BY product_name');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;
```

---

### Q6. How do you work with JSON in MySQL (5.7+)?

```sql
-- Store JSON data
CREATE TABLE products (
    id         INT PRIMARY KEY,
    name       VARCHAR(200),
    attributes JSON           -- flexible schema per product
);

INSERT INTO products VALUES
(1, 'Laptop',    '{"cpu": "i7", "ram": 16, "storage": "512GB SSD", "tags": ["portable", "work"]}'),
(2, 'Headphones','{"driver": "40mm", "wireless": true, "tags": ["audio", "portable"]}');

-- JSON path extraction
SELECT
    name,
    attributes->'$.cpu'           AS cpu,         -- returns "i7" (with quotes)
    attributes->>'$.cpu'          AS cpu_unquoted, -- returns i7 (no quotes)
    JSON_EXTRACT(attributes, '$.ram') AS ram,
    JSON_UNQUOTE(attributes->'$.tags[0]') AS first_tag
FROM products;

-- JSON in WHERE clause
SELECT * FROM products WHERE attributes->>'$.wireless' = 'true';
SELECT * FROM products WHERE JSON_CONTAINS(attributes->'$.tags', '"portable"');

-- Modify JSON
UPDATE products
SET attributes = JSON_SET(attributes, '$.ram', 32, '$.ssd', true)
WHERE id = 1;

-- JSON_ARRAYAGG — aggregate into JSON array
SELECT department,
    JSON_ARRAYAGG(name ORDER BY name) AS member_names
FROM employees
GROUP BY department;
```

---

### Q7. How do you write a recursive CTE for hierarchical data?

```sql
-- Organizational hierarchy: find all reports under a given manager
WITH RECURSIVE subordinates AS (
    -- Anchor: start with the target manager
    SELECT id, name, manager_id, 1 AS depth
    FROM employees
    WHERE id = 10  -- starting manager

    UNION ALL

    -- Recursive: find direct reports of each person already in the CTE
    SELECT e.id, e.name, e.manager_id, s.depth + 1
    FROM employees e
    JOIN subordinates s ON e.manager_id = s.id
)
SELECT id, name, depth
FROM subordinates
ORDER BY depth, name;

-- Bill of Materials (BOM): find all components of a product
WITH RECURSIVE bom AS (
    SELECT component_id, parent_id, quantity, 1 AS level,
           CAST(component_id AS CHAR(100)) AS path
    FROM bill_of_materials
    WHERE parent_id IS NULL  -- top-level products

    UNION ALL

    SELECT b.component_id, b.parent_id, b.quantity * bom.quantity,
           bom.level + 1,
           CONCAT(bom.path, ' > ', b.component_id)
    FROM bill_of_materials b
    JOIN bom ON b.parent_id = bom.component_id
)
SELECT * FROM bom ORDER BY path;
```

---

### Q8. How do you use FIRST_VALUE, LAST_VALUE, NTH_VALUE?

```sql
-- FIRST_VALUE: value from the first row in the window partition
-- LAST_VALUE: value from the last row in the window partition
-- NTH_VALUE: value from the Nth row in the window partition

SELECT
    department,
    name,
    salary,
    FIRST_VALUE(name)  OVER w AS highest_earner_in_dept,
    LAST_VALUE(name)   OVER (PARTITION BY department ORDER BY salary DESC
                             ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
                            ) AS lowest_earner_in_dept,
    NTH_VALUE(name, 2) OVER w AS second_highest_in_dept
FROM employees
WINDOW w AS (PARTITION BY department ORDER BY salary DESC)  -- named window
ORDER BY department, salary DESC;

-- ⚠️ LAST_VALUE gotcha: default frame is ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
-- This means LAST_VALUE = current row (not the real last!)
-- Always specify: ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
-- to get the true last row value
```

---

### Q9. How do you use dynamic SQL with PREPARE and EXECUTE?

```sql
-- Use case: build query string dynamically based on parameters
DELIMITER $$

CREATE PROCEDURE SearchEmployees(
    IN search_field VARCHAR(50),
    IN search_value VARCHAR(200),
    IN sort_column  VARCHAR(50)
)
BEGIN
    SET @query = CONCAT(
        'SELECT id, name, department, salary FROM employees ',
        'WHERE ', search_field, ' LIKE ? ',
        'ORDER BY ', sort_column, ' ',
        'LIMIT 100'
    );

    SET @search = CONCAT('%', search_value, '%');

    PREPARE stmt FROM @query;
    EXECUTE stmt USING @search;
    DEALLOCATE PREPARE stmt;
END$$

DELIMITER ;

-- Call it
CALL SearchEmployees('department', 'Eng', 'salary DESC');
CALL SearchEmployees('name', 'john', 'name');

-- ⚠️ SECURITY: NEVER concatenate user input directly — use ? placeholders
-- Column names cannot be parameterized (use whitelist validation in application)
```

---

### Q10. Common advanced MySQL query patterns.

```sql
-- 1. Upsert: insert or update on duplicate key
INSERT INTO product_inventory (product_id, quantity, last_updated)
VALUES (100, 50, NOW())
ON DUPLICATE KEY UPDATE
    quantity    = quantity + VALUES(quantity),
    last_updated = NOW();

-- 2. INSERT IGNORE: silently ignore duplicate key errors
INSERT IGNORE INTO user_events (user_id, event_type, event_date)
VALUES (1, 'login', NOW());

-- 3. REPLACE INTO: delete existing row + insert new (careful with foreign keys!)
REPLACE INTO settings (key_name, value) VALUES ('max_connections', '200');

-- 4. UPDATE with JOIN
UPDATE employees e
JOIN departments d ON e.department = d.name
SET e.salary = e.salary * 1.10
WHERE d.budget > 1000000;  -- 10% raise for employees in high-budget departments

-- 5. Conditional UPDATE (no need for application-level if-else)
UPDATE orders
SET status = CASE
    WHEN payment_status = 'paid' AND shipped = 1 THEN 'complete'
    WHEN payment_status = 'paid' AND shipped = 0 THEN 'processing'
    WHEN payment_status = 'pending'              THEN 'awaiting_payment'
    ELSE 'unknown'
END;
```

