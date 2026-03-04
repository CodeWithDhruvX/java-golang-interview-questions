# 🗄️ 01 — MySQL Basics
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

> Core fundamentals — heavily tested at **TCS, Infosys, Wipro, Cognizant, Capgemini**.

---

## 🔑 Must-Know Topics
- DDL vs DML vs DCL vs TCL
- Data types (INT, VARCHAR, TEXT, DECIMAL, DATETIME, JSON)
- Primary key, foreign key, unique key, composite key
- Constraints: NOT NULL, DEFAULT, CHECK, AUTO_INCREMENT
- Normalization (1NF, 2NF, 3NF, BCNF)
- Basic SELECT, INSERT, UPDATE, DELETE
- WHERE, GROUP BY, HAVING, ORDER BY, LIMIT

---

## ❓ Most Asked Questions

### Q1. What is MySQL and what are its key features?
**Answer:**
MySQL is an open-source **relational database management system (RDBMS)** that uses SQL. Key features:
- **ACID-compliant** via InnoDB storage engine
- **Pluggable storage engines** (InnoDB, MyISAM, Memory)
- **Replication** support (master–slave, group replication)
- **Stored procedures, triggers, views, events**
- **Full-text search**, **partitioning**, **JSON support** (5.7+)
- **Client-server model** — clients connect via TCP/IP (port 3306) or Unix socket

---

### Q2. What is the difference between DDL, DML, DCL, and TCL?

| Category | Full Form | Commands | Auto-commit? |
|----------|-----------|----------|--------------|
| **DDL** | Data Definition Language | CREATE, ALTER, DROP, TRUNCATE, RENAME | ✅ Yes |
| **DML** | Data Manipulation Language | SELECT, INSERT, UPDATE, DELETE | ❌ No (transactional) |
| **DCL** | Data Control Language | GRANT, REVOKE | ✅ Yes |
| **TCL** | Transaction Control Language | COMMIT, ROLLBACK, SAVEPOINT | Manual |

```sql
-- DDL
CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(100));
ALTER TABLE users ADD COLUMN email VARCHAR(255);
DROP TABLE users;

-- DML
INSERT INTO users VALUES (1, 'Alice', 'alice@example.com');
UPDATE users SET name = 'Bob' WHERE id = 1;
DELETE FROM users WHERE id = 1;
SELECT * FROM users WHERE id = 1;
```

---

### Q3. What is the difference between DELETE, TRUNCATE, and DROP?

| | DELETE | TRUNCATE | DROP |
|---|--------|----------|------|
| **Removes** | Specific rows | All rows | Entire table |
| **WHERE clause** | ✅ Yes | ❌ No | ❌ No |
| **Rollback** | ✅ Yes | ❌ No (MySQL) | ❌ No |
| **Auto-increment reset** | ❌ No | ✅ Yes | N/A |
| **Triggers fired** | ✅ Yes | ❌ No | ❌ No |
| **Category** | DML | DDL | DDL |

```sql
-- DELETE — row by row, can rollback
DELETE FROM orders WHERE status = 'cancelled';

-- TRUNCATE — removes all rows, resets AUTO_INCREMENT
TRUNCATE TABLE temp_logs;

-- DROP — removes table entirely (structure + data + indexes)
DROP TABLE IF EXISTS old_archive;
```

---

### Q4. What are primary keys, foreign keys, and unique keys?

```sql
CREATE TABLE customers (
    id         INT AUTO_INCREMENT PRIMARY KEY,   -- Primary Key: unique + NOT NULL
    email      VARCHAR(255) UNIQUE,              -- Unique: unique, can be NULL
    phone      VARCHAR(20)
);

CREATE TABLE orders (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    total       DECIMAL(10,2),
    FOREIGN KEY (customer_id)                    -- Foreign Key: referential integrity
        REFERENCES customers(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
```

| | Primary Key | Unique Key | Foreign Key |
|---|---|---|---|
| Unique | ✅ | ✅ | ✅ (references PK) |
| NULL allowed | ❌ | ✅ (one NULL) | ✅ |
| Per table | 1 | Multiple | Multiple |
| Index created | ✅ Clustered | ✅ Non-clustered | ✅ (required) |

---

### Q5. What is AUTO_INCREMENT and how does it work?

```sql
CREATE TABLE products (
    id    BIGINT AUTO_INCREMENT PRIMARY KEY,
    name  VARCHAR(200)
);

INSERT INTO products (name) VALUES ('Laptop'), ('Mouse');
-- id 1, 2 assigned automatically

-- Check last inserted ID
SELECT LAST_INSERT_ID();

-- Reset counter
ALTER TABLE products AUTO_INCREMENT = 1000;
```

> ⚠️ Use `BIGINT` for large tables. In MySQL 8.0+ the counter is persisted in redo log (fixes the restart-reset bug of older versions).

---

### Q6. What are MySQL's key data types?

```sql
-- Numeric
id          TINYINT,          -- 1 byte, -128 to 127
age         SMALLINT,         -- 2 bytes
count       INT,              -- 4 bytes
population  BIGINT,           -- 8 bytes
price       DECIMAL(10, 2),   -- exact decimal (use for money!)
rating      FLOAT,            -- approximate (avoid for money)

-- String
code        CHAR(3),          -- fixed-length, padded with spaces
name        VARCHAR(255),     -- variable-length, up to 65535 bytes per row
bio         TEXT,             -- up to 65KB, stored off-page
content     LONGTEXT,         -- up to 4GB

-- Date/Time
dob         DATE,             -- 'YYYY-MM-DD'
created_at  DATETIME,         -- 'YYYY-MM-DD HH:MM:SS', no timezone
updated_at  TIMESTAMP,        -- auto UTC, converts on SELECT per session TZ
duration    TIME,             -- '-838:59:59' to '838:59:59'

-- Other
metadata    JSON,             -- structured JSON (MySQL 5.7+)
flag        ENUM('Y','N'),    -- 1 of defined values
tags        SET('red','blue') -- 0 or more defined values
```

---

### Q7. What is normalization? Explain 1NF, 2NF, 3NF.

**Normalization** organizes tables to reduce redundancy and improve integrity.

```sql
-- UNNORMALIZED: student_courses has repeating groups
-- student_id | name  | courses
--     1       | Alice | Math, Physics, CS

-- 1NF — Atomic values, no repeating groups
CREATE TABLE student_courses (
    student_id INT, name VARCHAR(100), course VARCHAR(100)
);
-- (1, Alice, Math), (1, Alice, Physics), (1, Alice, CS)

-- 2NF — No partial dependency (non-key col depends on WHOLE PK)
-- Problem: 'name' depends only on student_id, not (student_id, course)
CREATE TABLE students (student_id INT PRIMARY KEY, name VARCHAR(100));
CREATE TABLE enrollments (student_id INT, course VARCHAR(100),
    PRIMARY KEY (student_id, course));

-- 3NF — No transitive dependency (non-key col → non-key col)
-- If we had: student_id → zip_code → city, city is transitively dependent
-- Fix: Create a zip_codes table
CREATE TABLE zip_codes (zip CHAR(5) PRIMARY KEY, city VARCHAR(100));
```

---

### Q8. What is the difference between WHERE and HAVING?

```sql
-- WHERE: filters ROWS before GROUP BY — can use indexes
SELECT department, AVG(salary)
FROM employees
WHERE salary > 30000          -- ← filters individual rows first
GROUP BY department;

-- HAVING: filters GROUPS after GROUP BY — cannot use indexes for the filter
SELECT department, AVG(salary)
FROM employees
GROUP BY department
HAVING AVG(salary) > 50000;  -- ← filters groups after aggregation

-- WRONG: cannot use aggregate in WHERE
-- WHERE AVG(salary) > 50000  ← Error!

-- COMBINED
SELECT department, COUNT(*) as cnt, AVG(salary) as avg_sal
FROM employees
WHERE status = 'active'        -- filter rows first (uses index on status)
GROUP BY department
HAVING cnt > 5                 -- filter groups
ORDER BY avg_sal DESC
LIMIT 10;
```

---

### Q9. How do you use GROUP BY with aggregates?

```sql
-- Count employees per department
SELECT department, COUNT(*) AS headcount
FROM employees
GROUP BY department;

-- Multiple aggregates
SELECT
    department,
    COUNT(*)            AS headcount,
    AVG(salary)         AS avg_salary,
    MAX(salary)         AS max_salary,
    MIN(salary)         AS min_salary,
    SUM(salary)         AS total_payroll,
    GROUP_CONCAT(name ORDER BY name SEPARATOR ', ') AS members
FROM employees
GROUP BY department
ORDER BY total_payroll DESC;

-- GROUP BY with ROLLUP (subtotals + grand total)
SELECT department, job_title, SUM(salary)
FROM employees
GROUP BY department, job_title WITH ROLLUP;
```

---

### Q10. What is a composite key and when do you use it?

```sql
-- Composite Primary Key — for junction/mapping tables
CREATE TABLE student_courses (
    student_id INT        NOT NULL,
    course_id  INT        NOT NULL,
    enrolled_at DATETIME  DEFAULT CURRENT_TIMESTAMP,
    grade       CHAR(2),
    PRIMARY KEY (student_id, course_id),  -- composite PK
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id)  REFERENCES courses(id)
);

-- Composite UNIQUE key — multi-column uniqueness
CREATE TABLE product_prices (
    product_id  INT,
    region_id   INT,
    price       DECIMAL(10,2),
    UNIQUE KEY uq_product_region (product_id, region_id)  -- combination must be unique
);
```

---

### Q11. What is the CHAR vs VARCHAR difference?

```sql
-- CHAR(n) — fixed length, always pads to n characters
-- Good for: country codes, status flags, fixed-format data
CREATE TABLE countries (
    code    CHAR(2),    -- always 2 bytes (e.g., 'US', 'IN')
    status  CHAR(1)     -- 'A'=active, 'I'=inactive
);

-- VARCHAR(n) — variable length, stores actual length + 1-2 overhead bytes
-- Good for: names, emails, descriptions
CREATE TABLE users (
    email    VARCHAR(255),   -- 5 bytes for 'a@b.c', not 255
    bio      VARCHAR(1000)
);

-- CHAR is marginally faster for fixed-size equality lookups
-- VARCHAR is more storage-efficient for varied-length strings
-- In MEMORY engine, VARCHAR is stored as CHAR (no space saving)
```

---

### Q12. How does NULL work in MySQL?

```sql
-- NULL means "unknown" — not zero, not empty string
SELECT NULL = NULL;     -- NULL (not TRUE!)
SELECT NULL IS NULL;    -- 1 (TRUE)
SELECT NULL IS NOT NULL;-- 0 (FALSE)

-- Aggregate functions IGNORE NULLs (except COUNT(*))
SELECT COUNT(*) FROM employees;          -- counts ALL rows
SELECT COUNT(bonus) FROM employees;      -- counts only non-NULL bonus rows

-- COALESCE — returns first non-NULL value
SELECT COALESCE(bonus, 0) AS payout FROM employees;

-- IFNULL — MySQL-specific shorthand
SELECT IFNULL(commission, 0) AS commission FROM sales;

-- NULLIF — returns NULL if two values are equal (avoid division by zero)
SELECT total / NULLIF(count, 0) AS average FROM stats;
```

---

### Q13. What are MySQL constraints?

```sql
CREATE TABLE employees (
    id          INT AUTO_INCREMENT PRIMARY KEY,   -- PK: unique + NOT NULL
    name        VARCHAR(100) NOT NULL,            -- NOT NULL: must have value
    email       VARCHAR(255) UNIQUE,              -- UNIQUE: no duplicates
    department  VARCHAR(50)  DEFAULT 'General',   -- DEFAULT: fallback value
    salary      DECIMAL(10,2) CHECK (salary > 0), -- CHECK: validated (MySQL 8.0.16+)
    manager_id  INT,
    FOREIGN KEY (manager_id) REFERENCES employees(id) ON DELETE SET NULL
);

-- Add constraint after creation
ALTER TABLE employees ADD CONSTRAINT chk_salary CHECK (salary <= 1000000);
ALTER TABLE employees DROP CHECK chk_salary;

-- Temporarily disable foreign key checks (for bulk import)
SET FOREIGN_KEY_CHECKS = 0;
-- ... bulk insert ...
SET FOREIGN_KEY_CHECKS = 1;
```

---

### Q14. What is the ORDER BY + LIMIT + OFFSET pattern?

```sql
-- Basic ordering + limiting
SELECT * FROM products
ORDER BY price DESC
LIMIT 10;

-- Pagination with OFFSET (page 3, 10 per page)
SELECT * FROM products
ORDER BY created_at DESC
LIMIT 10 OFFSET 20;     -- skip first 20, return rows 21-30

-- Shorthand: LIMIT offset, count
SELECT * FROM products ORDER BY id LIMIT 20, 10;

-- ⚠️ OFFSET pagination is SLOW for large tables (scans and discards rows)
-- Better: Keyset/cursor pagination
SELECT * FROM products
WHERE id > 1500          -- last seen id from previous page
ORDER BY id
LIMIT 10;
```

---

### Q15. How do you use CASE expressions in MySQL?

```sql
-- Simple CASE
SELECT name,
    CASE department
        WHEN 'Engineering' THEN 'Tech'
        WHEN 'Sales'       THEN 'Revenue'
        ELSE 'Other'
    END AS department_group
FROM employees;

-- Searched CASE (more powerful)
SELECT name, salary,
    CASE
        WHEN salary < 40000  THEN 'Junior'
        WHEN salary < 80000  THEN 'Mid'
        WHEN salary < 120000 THEN 'Senior'
        ELSE 'Principal'
    END AS level
FROM employees;

-- CASE in aggregate (conditional counting)
SELECT
    COUNT(CASE WHEN gender = 'M' THEN 1 END) AS male_count,
    COUNT(CASE WHEN gender = 'F' THEN 1 END) AS female_count
FROM employees;
```

---

### Q16. What is the difference between INNER JOIN, LEFT JOIN, and RIGHT JOIN?

```sql
-- Sample tables
-- customers: id, name       orders: id, customer_id, total

-- INNER JOIN — only matching rows in BOTH tables
SELECT c.name, o.total
FROM customers c
INNER JOIN orders o ON c.id = o.customer_id;

-- LEFT JOIN — all customers, NULL if no orders
SELECT c.name, o.total
FROM customers c
LEFT JOIN orders o ON c.id = o.customer_id;

-- RIGHT JOIN — all orders, NULL if no customer (rare, prefer LEFT JOIN)
SELECT c.name, o.total
FROM customers c
RIGHT JOIN orders o ON c.id = o.customer_id;

-- Find customers with NO orders (anti-join pattern)
SELECT c.name
FROM customers c
LEFT JOIN orders o ON c.id = o.customer_id
WHERE o.id IS NULL;

-- FULL OUTER JOIN simulation (MySQL doesn't support it natively)
SELECT c.name, o.total FROM customers c LEFT  JOIN orders o ON c.id = o.customer_id
UNION
SELECT c.name, o.total FROM customers c RIGHT JOIN orders o ON c.id = o.customer_id;
```

