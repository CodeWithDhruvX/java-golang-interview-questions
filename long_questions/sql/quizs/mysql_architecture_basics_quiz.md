# MySQL Architecture & Basics Quiz

## 1. Questions

### Q1: Data Type Nuances
**Part A:** What is the difference between `CHAR` and `VARCHAR`? When would you strictly prefer `CHAR`?
**Part B:** How does `TIMESTAMP` differ from `DATETIME` regarding time zones?

### Q2: Data Removal Commands
**Question:** Explain the key differences between `DELETE` and `TRUNCATE`. Discuss transaction support, speed, and trigger activation.

### Q3: Set Operations
**Question:** You have two queries. Query A returns 10 rows. Query B returns 10 rows (5 of which are identical to Query A).
*   How many rows does `Query A UNION Query B` return?
*   How many rows does `Query A UNION ALL Query B` return?
*   Which is faster and why?

### Q4: Foreign Key Constraints
**Scenario:** You have a `Orders` table with a foreign key to `Customers`.
**Question:** What happens to the `Orders` if you try to delete a `Customer`? Explain the default behavior vs `ON DELETE CASCADE` vs `ON DELETE SET NULL`.

### Q5: Architecture - Replication
**Question:** In a standard MySQL Primary-Replica (Master-Slave) replication setup, is the replication **Synchronous** or **Asynchronous** by default? What impact does this have on data consistency if the Primary crashes?

### Q6: Security - SQL Injection
**Scenario:** A developer writes code like: `query = "SELECT * FROM users WHERE name = '" + userInput + "'"`
**Question:** How can an attacker exploit this? What is the standard way to fix it in modern applications?

---

## 2. Answers & Explanations

### Answer 1: Data Types
**Part A (CHAR vs VARCHAR):**
*   `CHAR(N)`: Fixed length. If you store "Hi" in `CHAR(10)`, it pads it with spaces to 10 bytes. **Use when:** Data is fixed size (e.g., Country Codes 'US', 'IN', UUIDs, Hashes) for slightly better performance and reduced fragmentation.
*   `VARCHAR(N)`: Variable length. Stores "Hi" as 2 bytes + length prefix. **Use when:** Strings vary widely (e.g., Names, Emails).

**Part B (TIMESTAMP vs DATETIME):**
*   `TIMESTAMP`: Stored as UTC. It converts *from* your connection timezone to UTC on insert, and *back* to your connection timezone on retrieval. Range ends in 2038.
*   `DATETIME`: Stored explicitly as the date/time you entered. No timezone conversion occurs. Range is 1000-9999.

### Answer 2: DELETE vs TRUNCATE
*   **DELETE**: DML command. Deletes row-by-row. Logs every deletion. **Can be rolled back** in a transaction. **Fires triggers**. Slower for large tables.
*   **TRUNCATE**: DDL command. Drops the table and re-creates it (conceptually). **Cannot be rolled back** (in many standard setups, though InnoDB allows rollback in non-committed transactions). **Does NOT fire triggers**. Extremely fast. Resets Auto-Increment counters.

### Answer 3: UNION vs UNION ALL
*   **UNION**: Removes duplicates. Returns **15 rows** (10 from A + 5 unique from B). Slower because it must sort/hash to find duplicates.
*   **UNION ALL**: Keeps duplicates. Returns **20 rows** (10 from A + 10 from B). **Faster** because it just appends results.

### Answer 4: Foreign Key Actions
*   **Default (RESTRICT/NO ACTION)**: The DB **prevents** the deletion of the Customer. It returns an error saying child records exist.
*   **ON DELETE CASCADE**: Automatically **deletes all Orders** belonging to that Customer. (Dangerous but maintains referential integrity).
*   **ON DELETE SET NULL**: Updates the `customer_id` in the `Orders` table to `NULL` (requires the column to be nullable). Keeps the order history but orphans it.

### Answer 5: Replication & Consistency
*   **Default**: **Asynchronous**. The Primary writes changes to its log and returns "Success" to the client *before* the Replica confirms it received the update.
*   **Risk**: If the Primary crashes immediately after writing but before sending to the Replica, that data is **lost** during failover (Replica Lag).

### Answer 6: SQL Injection
*   **Exploit**: Inputting `' OR '1'='1` turns the query into `SELECT * FROM users WHERE name = '' OR '1'='1'`, which is always true. This returns ALL users (dumping the database) or allows bypassing login logic.
*   **Fix**: Use **Parameterized Queries** (Prepared Statements).
    *   *Bad*: `db.query("SELECT ... " + input)`
    *   *Good*: `db.query("SELECT ... ?", input)`
    *   The database treats the input strictly as data (literal string), never as executable code.
