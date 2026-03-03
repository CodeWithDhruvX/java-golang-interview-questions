# 🛠️ 03 — Stored Procedures, Functions & Triggers
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

> Server-side logic commonly expected in **TCS, Infosys, Wipro, HCL** technical rounds.

---

## 🔑 Must-Know Topics
- Stored procedures vs stored functions
- IN, OUT, INOUT parameters
- Triggers: BEFORE/AFTER, INSERT/UPDATE/DELETE
- Events (scheduled jobs)
- Cursors
- Error handling: DECLARE HANDLER
- DELIMITER syntax

---

## ❓ Most Asked Questions

### Q1. What is a stored procedure? Create one with IN/OUT parameters.

```sql
DELIMITER $$

CREATE PROCEDURE GetEmployeesByDept(
    IN  dept_name  VARCHAR(100),   -- input parameter
    OUT emp_count  INT             -- output parameter
)
BEGIN
    -- Return all employees in a department
    SELECT id, name, salary
    FROM employees
    WHERE department = dept_name;

    -- Set output parameter
    SELECT COUNT(*) INTO emp_count
    FROM employees
    WHERE department = dept_name;
END$$

DELIMITER ;

-- Call the procedure
CALL GetEmployeesByDept('Engineering', @count);
SELECT @count AS engineer_count;
```

---

### Q2. What is the difference between a stored procedure and a stored function?

| | Stored Procedure | Stored Function |
|---|---|---|
| **Returns** | Via OUT params or result set | Single value with RETURN |
| **Called via** | `CALL proc()` | `SELECT func()` in SQL |
| **Used in SQL** | ❌ Cannot embed in SELECT | ✅ Yes — usable in SELECT, WHERE |
| **Can modify data** | ✅ Yes | ⚠️ Avoid — side effects in queries |
| **Transactions** | Can COMMIT/ROLLBACK | Cannot |

```sql
-- Stored FUNCTION — returns a single value, usable in SQL
DELIMITER $$

CREATE FUNCTION CalculateBonus(
    base_salary DECIMAL(10,2),
    performance_rating INT
) RETURNS DECIMAL(10,2)
DETERMINISTIC                  -- same inputs always return same output
BEGIN
    DECLARE bonus DECIMAL(10,2);
    SET bonus = CASE
        WHEN performance_rating >= 9 THEN base_salary * 0.20
        WHEN performance_rating >= 7 THEN base_salary * 0.10
        ELSE base_salary * 0.05
    END;
    RETURN bonus;
END$$

DELIMITER ;

-- Use in SQL just like a built-in function
SELECT name, salary, CalculateBonus(salary, 8) AS bonus
FROM employees;
```

---

### Q3. How do you add error handling to a stored procedure?

```sql
DELIMITER $$

CREATE PROCEDURE TransferFunds(
    IN from_account INT,
    IN to_account   INT,
    IN amount       DECIMAL(10,2)
)
BEGIN
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        -- This block runs on any SQL error
        ROLLBACK;
        SELECT 'Transfer failed — rolled back' AS message;
    END;

    START TRANSACTION;

    UPDATE accounts SET balance = balance - amount WHERE id = from_account;
    IF ROW_COUNT() = 0 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Source account not found';
    END IF;

    UPDATE accounts SET balance = balance + amount WHERE id = to_account;
    IF ROW_COUNT() = 0 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Destination account not found';
    END IF;

    COMMIT;
    SELECT 'Transfer successful' AS message;
END$$

DELIMITER ;
```

---

### Q4. What is a trigger? Write BEFORE and AFTER triggers.

```sql
-- BEFORE trigger: runs BEFORE the operation, can modify NEW values
DELIMITER $$

CREATE TRIGGER before_employee_insert
BEFORE INSERT ON employees
FOR EACH ROW
BEGIN
    -- Normalize name to proper case before saving
    SET NEW.name = CONCAT(
        UPPER(SUBSTRING(NEW.name, 1, 1)),
        LOWER(SUBSTRING(NEW.name, 2))
    );
    -- Set default department if not provided
    IF NEW.department IS NULL THEN
        SET NEW.department = 'General';
    END IF;
END$$

-- AFTER trigger: runs AFTER the operation, for audit logging
CREATE TRIGGER after_salary_update
AFTER UPDATE ON employees
FOR EACH ROW
BEGIN
    IF OLD.salary != NEW.salary THEN
        INSERT INTO salary_audit_log (
            employee_id, old_salary, new_salary,
            changed_by, changed_at
        ) VALUES (
            NEW.id, OLD.salary, NEW.salary,
            USER(), NOW()
        );
    END IF;
END$$

DELIMITER ;
```

---

### Q5. Trigger types overview and key differences.

| Trigger | When it fires | Can modify NEW? | Can modify OLD? |
|---------|---------------|-----------------|-----------------|
| BEFORE INSERT | Before row inserted | ✅ Yes | N/A |
| AFTER INSERT  | After row inserted  | ❌ No  | N/A |
| BEFORE UPDATE | Before row updated  | ✅ Yes | ❌ No |
| AFTER UPDATE  | After row updated   | ❌ No  | ❌ No |
| BEFORE DELETE | Before row deleted  | N/A    | ✅ Yes |
| AFTER DELETE  | After row deleted   | N/A    | ❌ No |

```sql
-- View all triggers
SHOW TRIGGERS FROM myapp_db;
SHOW TRIGGERS LIKE 'employee%';

-- Drop a trigger
DROP TRIGGER IF EXISTS after_salary_update;

-- Triggers cannot call stored procedures that do transactions
-- Triggers fire per-row (FOR EACH ROW is required in MySQL)
-- A single INSERT ... VALUES (1),(2),(3) fires trigger 3 times
```

---

### Q6. What is an Event (scheduled job) in MySQL?

```sql
-- Enable the event scheduler
SET GLOBAL event_scheduler = ON;

-- Create a recurring event (every day at 2 AM)
CREATE EVENT cleanup_old_sessions
ON SCHEDULE EVERY 1 DAY
STARTS '2024-01-01 02:00:00'
ON COMPLETION PRESERVE               -- keep event after it fires (for recurring)
COMMENT 'Delete sessions older than 30 days'
DO
BEGIN
    DELETE FROM user_sessions
    WHERE last_active < NOW() - INTERVAL 30 DAY;

    INSERT INTO maintenance_log(event_name, ran_at, rows_deleted)
    SELECT 'cleanup_old_sessions', NOW(), ROW_COUNT();
END;

-- One-time event (fires once then dropped)
CREATE EVENT send_monthly_report
ON SCHEDULE AT '2024-02-01 08:00:00'
DO
    CALL GenerateMonthlyReport();

-- Show events
SHOW EVENTS FROM myapp_db;

-- Drop event
DROP EVENT IF EXISTS cleanup_old_sessions;
```

---

### Q7. How do you use cursors in a stored procedure?

```sql
-- Cursor: iterate row by row through a result set (use sparingly — set-based is better)
DELIMITER $$

CREATE PROCEDURE ApplyYearEndBonuses()
BEGIN
    DECLARE done        INT DEFAULT FALSE;
    DECLARE emp_id      INT;
    DECLARE emp_salary  DECIMAL(10,2);
    DECLARE bonus_pct   DECIMAL(5,2);

    -- 1. Declare the cursor
    DECLARE emp_cursor CURSOR FOR
        SELECT id, salary FROM employees WHERE active = 1;

    -- 2. Declare NOT FOUND handler
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

    -- 3. Open the cursor
    OPEN emp_cursor;

    read_loop: LOOP
        -- 4. Fetch next row
        FETCH emp_cursor INTO emp_id, emp_salary;
        IF done THEN
            LEAVE read_loop;
        END IF;

        -- 5. Business logic per row
        SET bonus_pct = IF(emp_salary > 100000, 0.10, 0.15);

        UPDATE bonuses SET amount = emp_salary * bonus_pct
        WHERE employee_id = emp_id AND YEAR(bonus_year) = YEAR(NOW());

        IF ROW_COUNT() = 0 THEN
            INSERT INTO bonuses(employee_id, amount, bonus_year)
            VALUES (emp_id, emp_salary * bonus_pct, NOW());
        END IF;
    END LOOP;

    -- 6. Close the cursor
    CLOSE emp_cursor;
END$$

DELIMITER ;
```

---

### Q8. What is a view? When should you use one?

```sql
-- VIEW: a named, stored SELECT query — acts like a virtual table
CREATE VIEW vw_employee_summary AS
    SELECT
        e.id,
        e.name,
        e.department,
        e.salary,
        m.name   AS manager_name,
        d.budget AS dept_budget
    FROM employees e
    LEFT JOIN employees   m ON e.manager_id = m.id
    LEFT JOIN departments d ON e.department = d.name;

-- Use like a table
SELECT * FROM vw_employee_summary WHERE department = 'Engineering';

-- Updatable view (simple views with direct column mapping)
CREATE VIEW vw_active_users AS
    SELECT id, name, email FROM users WHERE active = 1;

UPDATE vw_active_users SET email = 'new@email.com' WHERE id = 5;  -- works!

-- Non-updatable view (aggregates, DISTINCT, UNION → read-only)
CREATE VIEW vw_dept_stats AS
    SELECT department, AVG(salary) AS avg_sal FROM employees GROUP BY department;

-- Benefits:
-- ✅ Security: expose only selected columns (hide sensitive columns like salary, SSN)
-- ✅ Simplify complex JOINs — one view, multiple places
-- ❌ No performance benefit for regular views (query runs fresh each time)
-- ✅ Materialized view (via manual table + events) = performance benefit

-- Drop view
DROP VIEW IF EXISTS vw_employee_summary;
```

---

### Q9. What is the DELIMITER command and why is it needed?

```sql
-- Default delimiter in MySQL is ';'
-- Inside stored procs/functions/triggers, we use ';' to end statements
-- This confuses the MySQL client — it thinks the proc ends at first ';'

-- SOLUTION: change the delimiter temporarily
DELIMITER $$    -- now '$$' signals end of statement to client

CREATE PROCEDURE demo()
BEGIN
    SELECT 1;   -- ';' here is just part of the procedure body
    SELECT 2;   -- ';' again — client sees '$$' as end, not ';'
END$$

DELIMITER ;     -- restore default delimiter

-- This is only needed when using command-line client or scripts
-- GUI tools (Workbench, DBeaver) handle this automatically
-- In code (Python, Go, Java): pass procedure body as a string — no DELIMITER needed

-- Check existing procedures
SHOW PROCEDURE STATUS WHERE db = 'myapp';
SHOW CREATE PROCEDURE TransferFunds;

-- Drop procedure
DROP PROCEDURE IF EXISTS demo;
```

