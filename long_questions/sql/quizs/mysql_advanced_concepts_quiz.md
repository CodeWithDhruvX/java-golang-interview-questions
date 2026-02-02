# MySQL Advanced Concepts Interview Questions

## 1. Questions

### Q1: ACID Properties
**Question:** Explain the **Atomicity** and **Isolation** properties of ACID in the context of a banking transfer (Account A to Account B).

### Q2: Transaction Isolation Levels
**Question:** What is the difference between `READ COMMITTED` and `REPEATABLE READ`? Which one is the default in MySQL InnoDB?

### Q3: Stored Procedures vs Application Logic
**Question:** What is a major advantage and a major disadvantage of using Stored Procedures to handle business logic compared to handling it in the application code (e.g., Java/Go)?

### Q4: Window Functions
**Question:** Explain what the `ROW_NUMBER()`, `RANK()`, and `DENSE_RANK()` functions do and how they differ when handling ties (duplicate values).

### Q5: Views vs Materialized Views
**Question:** MySQL (standard) considers Views to be "virtual". What does this mean concerning performance? How does this differ from a "Materialized View" found in other DBs (like PostgreSQL or Oracle)?

---

## 2. Answers & Explanations

### Answer 1: ACID Context
*   **Atomicity**: "All or Nothing". If money is deducted from Account A but the server crashes before adding it to Account B, the entire transaction must roll back. Account A should not lose money.
*   **Isolation**: "Interference Free". If a second transaction checks the balance of Account A while the transfer is happening (but not committed), it should see the original balance (or the new one depending on isolation level), not a temporary invalid intermediate state.

### Answer 2: Isolation Levels
*   **READ COMMITTED**: A transaction sees only data that has been committed. If Row X changes during your transaction, you will see the new value if you query it again (Non-repeatable read).
*   **REPEATABLE READ (Default in MySQL)**: Ensures that if you query Row X twice within the same transaction, you get the same result. It uses "Snapshot Isolation" to present a consistent view of the database as it was at the start of the transaction.

### Answer 3: Stored Procedures
*   **Advantage**: **Performance & Reduced Network Traffic**. The logic runs directly on the database server. You send one command (`CALL proc()`) instead of multiple SQL statements back and forth.
*   **Disadvantage**: **Maintainability & Debugging**. Version controlling stored procedures is harder. Business logic gets split between app code and DB. It couples the app tightly to the specific database vendor.

### Answer 4: Window Ranking Functions
Assuming values: `[10, 20, 20, 30]`
*   **`ROW_NUMBER()`**: Unique sequential ID. Result: `1, 2, 3, 4`. (Arbitrary order for ties).
*   **`RANK()`**: Skips numbers for ties. Result: `1, 2, 2, 4`. (No '3' because '2' was tied).
*   **`DENSE_RANK()`**: Does not skip numbers. Result: `1, 2, 2, 3`.

### Answer 5: Views
*   **MySQL Virtual Views**: A view is just a saved SQL query. When you query a view, MySQL effectively runs the underlying query in real-time. It **does not store** the data. Complex views can be slow because they re-compute every time.
*   **Materialized Views**: These physically store the result of the query on disk. They are much faster to read (like a table) but need to be refreshed/updated when the underlying data changes. *Note: MySQL does not support native Materialized Views out of the box (requiring workarounds with tables+triggers), whereas Postgres/Oracle do.*
