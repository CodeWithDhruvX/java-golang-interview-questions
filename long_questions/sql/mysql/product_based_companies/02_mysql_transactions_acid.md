# 🔒 02 — MySQL Transactions & ACID
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Medium–Hard

> Transactions and locking — **critical deep-dive topics** at **Uber, Razorpay, Zepto, Meesho, PhonePe**.

---

## 🔑 Must-Know Topics
- ACID properties: Atomicity, Consistency, Isolation, Durability
- Transaction control: BEGIN, COMMIT, ROLLBACK, SAVEPOINT
- Isolation levels: READ UNCOMMITTED → SERIALIZABLE
- Dirty read, non-repeatable read, phantom read
- Row-level locking vs table locking
- Deadlocks: detection and prevention
- MVCC (Multi-Version Concurrency Control)

---

## ❓ Most Asked Questions

### Q1. What are ACID properties?

```sql
-- A = Atomicity: ALL operations succeed, or ALL are rolled back
BEGIN;
UPDATE accounts SET balance = balance - 500 WHERE id = 1;  -- debit
UPDATE accounts SET balance = balance + 500 WHERE id = 2;  -- credit
COMMIT;   -- both succeed together
-- If crash between the two UPDATEs → ROLLBACK undoes partial work

-- C = Consistency: DB moves from one valid state to another valid state
-- (constraints, foreign keys, triggers all enforced across the transaction)

-- I = Isolation: concurrent transactions don't see each other's intermediate states
-- (controlled by isolation level)

-- D = Durability: committed data survives crashes
-- InnoDB: writes to redo log (WAL) before acknowledging commit
-- Even if server crashes, data is recovered on restart via redo log replay
```

---

### Q2. How do you manage transactions in MySQL?

```sql
-- Autocommit (default = ON): every statement is its own transaction
SHOW VARIABLES LIKE 'autocommit';  -- ON by default

-- Explicit transaction
START TRANSACTION;   -- or BEGIN;
  UPDATE wallet SET amount = amount - 1000 WHERE user_id = 1;
  INSERT INTO transactions(user_id, amount, type) VALUES (1, -1000, 'debit');
COMMIT;              -- makes changes permanent

-- Rollback on error
START TRANSACTION;
  UPDATE wallet SET amount = amount - 1000 WHERE user_id = 1;
  -- simulate error:
  SELECT 1/0;       -- error!
ROLLBACK;           -- undoes the UPDATE

-- SAVEPOINT — partial rollback
START TRANSACTION;
  INSERT INTO orders(customer_id) VALUES (100);          -- step 1
  SAVEPOINT after_order;
  INSERT INTO order_items(order_id) VALUES (LAST_INSERT_ID());  -- step 2
  SAVEPOINT after_items;
  -- If step 3 fails:
  ROLLBACK TO SAVEPOINT after_items;  -- undoes only step 3, keeps 1 and 2
COMMIT;
```

---

### Q3. What are the four isolation levels? What anomalies does each prevent?

| Isolation Level       | Dirty Read | Non-Repeatable Read | Phantom Read |
|-----------------------|-----------|---------------------|-------------|
| READ UNCOMMITTED      | ✅ Yes    | ✅ Yes              | ✅ Yes      |
| READ COMMITTED        | ❌ No     | ✅ Yes              | ✅ Yes      |
| REPEATABLE READ (default) | ❌ No | ❌ No               | ⚠️ Mostly No* |
| SERIALIZABLE          | ❌ No     | ❌ No               | ❌ No       |

```sql
-- Set isolation level for current session
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

-- Set globally
SET GLOBAL TRANSACTION ISOLATION LEVEL REPEATABLE READ;

-- Check current level
SELECT @@transaction_isolation;  -- MySQL 8.0
SELECT @@tx_isolation;           -- MySQL 5.7

-- REPEATABLE READ (MySQL default): InnoDB uses MVCC to give each transaction
-- a consistent snapshot at the START of the transaction.
-- InnoDB's MVCC mostly prevents phantom reads too — one of MySQL's differentiators.

START TRANSACTION;
  SELECT COUNT(*) FROM orders WHERE status='pending';  -- sees snapshot at T=0
  -- even if another transaction inserts a pending order here, this txn still sees T=0
  SELECT COUNT(*) FROM orders WHERE status='pending';  -- same count — repeatable!
COMMIT;
```

---

### Q4. What is a dirty read, non-repeatable read, and phantom read?

```sql
-- DIRTY READ: reading uncommitted changes from another transaction
-- Transaction A:
START TRANSACTION;
UPDATE accounts SET balance = 5000 WHERE id = 1;
-- (not committed yet)

-- Transaction B (READ UNCOMMITTED): reads A's uncommitted value!
SELECT balance FROM accounts WHERE id = 1;  -- returns 5000 (dirty!)
-- If A rolls back, B read data that never existed permanently.

-- NON-REPEATABLE READ: same query returns different values in same transaction
-- Transaction A reads salary=50000
-- Transaction B commits UPDATE salary=60000
-- Transaction A reads salary again → 60000  (different!! — non-repeatable)
-- Prevented by REPEATABLE READ and SERIALIZABLE

-- PHANTOM READ: new rows appear in a range query within same transaction
-- Transaction A: SELECT COUNT(*) WHERE age > 18 → returns 100
-- Transaction B: INSERT a new row with age=25, COMMIT
-- Transaction A: SELECT COUNT(*) WHERE age > 18 → returns 101 (phantom!)
-- Prevented by SERIALIZABLE (and mostly by InnoDB MVCC at REPEATABLE READ)
```

---

### Q5. What is MVCC (Multi-Version Concurrency Control)?

```sql
-- MVCC: InnoDB stores multiple "versions" of each row
-- Readers don't block writers; writers don't block readers
-- Each transaction sees a consistent snapshot of data as of when the transaction started

-- Two hidden columns on every InnoDB row:
-- DB_TRX_ID: the transaction ID that last modified this row
-- DB_ROLL_PTR: pointer to undo log (previous version of this row)

-- How a SELECT works under MVCC:
-- 1. MySQL records the current transaction ID when transaction starts
-- 2. For each row, checks DB_TRX_ID
-- 3. If the row was modified by a NEWER transaction → get older version from undo log
-- 4. If the row was modified by an OLDER committed transaction → use current version
-- 5. If the row was modified by an uncommitted transaction → get older version

-- Result: consistent reads without locking
-- Downside: long-running transactions prevent purge of old undo log versions
--           → undo log grows → "history list length" grows → performance degrades
SHOW ENGINE INNODB STATUS;  -- check "History list length"
```

---

### Q6. What is row-level locking in InnoDB?

```sql
-- InnoDB row-level locks — minimal blocking for concurrent transactions

-- Shared lock (S lock): multiple readers allowed, blocks exclusive locks
SELECT * FROM accounts WHERE id = 1 LOCK IN SHARE MODE;
-- or in MySQL 8.0+:
SELECT * FROM accounts WHERE id = 1 FOR SHARE;

-- Exclusive lock (X lock): blocks all other readers and writers
SELECT * FROM accounts WHERE id = 1 FOR UPDATE;

-- Typical use: "SELECT for update" pattern for safe balance update
START TRANSACTION;
  SELECT balance FROM accounts WHERE id = 1 FOR UPDATE;  -- locks this row
  -- now safely modify:
  UPDATE accounts SET balance = balance - 500 WHERE id = 1;
COMMIT;  -- lock released

-- Without FOR UPDATE:
-- Two transactions could both read balance=1000
-- Both subtract 500 → both write 500 → net result: 500 instead of 0 (lost update!)
```

---

### Q7. What is a deadlock? How does MySQL handle it?

```sql
-- Deadlock: two transactions each hold a lock the other needs
-- Transaction A:
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;  -- locks row 1
-- waiting for row 2...

-- Transaction B (concurrent):
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE id = 2;  -- locks row 2
UPDATE accounts SET balance = balance + 100 WHERE id = 1;  -- waiting for row 1!
-- DEADLOCK: A waits for B's lock on row 2, B waits for A's lock on row 1

-- MySQL detects deadlocks automatically via wait-for graph
-- Victim: the transaction with LEAST undo log work is rolled back
-- Error: ERROR 1213: Deadlock found when trying to get lock; try restarting transaction

-- Prevention strategies:
-- 1. Always acquire locks in SAME ORDER across transactions
-- 2. Keep transactions SHORT — less time holding locks
-- 3. Use lower isolation levels where acceptable
-- 4. Retry logic in application code

-- View last deadlock details:
SHOW ENGINE INNODB STATUS;  -- section: LATEST DETECTED DEADLOCK
```

---

### Q8. What is the difference between optimistic and pessimistic locking?

```sql
-- PESSIMISTIC LOCKING: assume conflict WILL happen → lock immediately on read
START TRANSACTION;
SELECT * FROM products WHERE id = 1 FOR UPDATE;   -- lock acquired immediately
-- Now update safely:
UPDATE products SET stock = stock - 1 WHERE id = 1;
COMMIT;
-- Good when: high contention, critical data (bank transfers, stock deduction)
-- Bad when: low contention (locks held unnecessarily → throughput loss)

-- OPTIMISTIC LOCKING: assume conflict WON'T happen → no lock on read
-- Use a version column to detect conflicts at update time
ALTER TABLE products ADD COLUMN version INT DEFAULT 0;

-- Read phase (no lock):
SELECT stock, version FROM products WHERE id = 1;
-- Application: stock=10, version=5

-- Write phase: check version didn't change
UPDATE products
SET stock = stock - 1, version = version + 1
WHERE id = 1 AND version = 5;  -- CAS: Compare-And-Swap

-- If affected_rows = 0 → someone else modified it → retry
-- If affected_rows = 1 → success
-- Good when: low contention, read-heavy workloads
```

---

### Q9. What is a gap lock and next-key lock in InnoDB?

```sql
-- InnoDB uses gap locks to prevent phantom reads in REPEATABLE READ

-- Gap lock: locks the gap BETWEEN index values (not the record itself)
-- Prevents INSERT into the gap

-- Next-key lock: combination of record lock + gap lock
-- Locks the record AND the gap before it

-- Example:
SELECT * FROM orders WHERE id BETWEEN 10 AND 20 FOR UPDATE;
-- InnoDB locks:
-- • Records with id 10, 11, ..., 20 (record locks)
-- • Gap before 10 and after 20 (gap locks)
-- → No other transaction can INSERT id=15 into this range

-- Gap locks can cause unexpected blocking:
-- Transaction A: SELECT WHERE status = 'pending' FOR UPDATE
-- Transaction B: INSERT new row with status = 'pending'
-- Transaction B is BLOCKED even though the row doesn't exist yet!

-- Disable gap locking if using READ COMMITTED:
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
-- READ COMMITTED uses only record locks, no gap locks → more concurrency
```

---

### Q10. How do you prevent long-running transactions?

```sql
-- Check for long-running transactions
SELECT
    trx_id,
    trx_started,
    TIMESTAMPDIFF(SECOND, trx_started, NOW()) AS seconds_running,
    trx_rows_modified,
    trx_query
FROM information_schema.innodb_trx
ORDER BY seconds_running DESC;

-- Kill a long-running transaction
-- First get the connection thread id:
SELECT trx_mysql_thread_id FROM information_schema.innodb_trx WHERE trx_id = 12345;
KILL 42;  -- kills the connection (and rolls back its transaction)

-- Set lock wait timeout (default: 50 seconds)
SET SESSION innodb_lock_wait_timeout = 30;  -- fail after 30 seconds waiting for lock

-- Set max transaction time:
SET SESSION MAX_EXECUTION_TIME = 5000;  -- 5 seconds max for SELECT statements

-- Best practices:
-- - Keep transactions under 100ms ideally
-- - Never wait on user input inside a transaction
-- - Batch large updates: UPDATE ... LIMIT 1000; loop
-- - Monitor: SET GLOBAL innodb_monitor_enable = 'all';
```

