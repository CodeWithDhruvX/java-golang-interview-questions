# 🗄️ Database Fundamentals — Interview Questions (Service-Based Companies)

This document covers database concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL, Cognizant. Targeted at 1–5 years of experience rounds.

---

### Q1: What is the difference between SQL and NoSQL databases? When should you use each?

**Answer:**

| Feature | SQL (Relational) | NoSQL |
|---|---|---|
| Schema | Fixed, predefined schema | Dynamic/schema-less |
| Data model | Tables with rows and columns | Document, Key-Value, Column, Graph |
| Query language | SQL (standard) | Database-specific APIs/query languages |
| Transactions | Strong ACID guarantees | Often eventual consistency (BASE) |
| Scalability | Vertical (scale-up) primarily | Horizontal (scale-out) |
| Relationships | JOINs — powerful relational queries | Denormalized, limited joins |
| Examples | PostgreSQL, MySQL, Oracle, SQL Server | MongoDB, Cassandra, Redis, DynamoDB, Neo4j |

**Use SQL when:**
- Structured data with well-defined relationships (e.g., e-commerce: orders, products, customers).
- ACID transactions are crucial (banking, payments).
- Complex queries with JOINs: reporting, analytics.

**Use NoSQL when:**
- Rapidly changing schema (startup, evolving product).
- Massive scale with simple access patterns (social media feeds, session storage).
- Specific data models: documents (MongoDB), time-series (InfluxDB), graphs (Neo4j).
- High write throughput with eventual consistency acceptable (Cassandra for log ingestion).

---

### Q2: Explain ACID properties with examples.

**Answer:**
ACID is a set of properties that guarantee database transactions are processed reliably.

**A — Atomicity:**
- A transaction is all-or-nothing. Either all operations succeed, or none take effect.
- Example: Bank transfer — debit $100 from account A AND credit $100 to account B. If credit fails, debit is rolled back.

**C — Consistency:**
- A transaction brings the database from one valid state to another. All constraints, rules, and triggers are satisfied.
- Example: Account balance cannot go negative (if that's a constraint). A transaction that would violate this is rejected.

**I — Isolation:**
- Concurrent transactions execute as if they were sequential. Intermediate states are not visible to other transactions.
- Example: Two concurrent bank transfers don't interfere with each other's balance reads.
- Implemented via: locks, MVCC (Multi-Version Concurrency Control).

**D — Durability:**
- Once a transaction commits, it remains committed even if the system crashes.
- Implemented via: Write-Ahead Logging (WAL) — changes written to log on disk before being applied.
- Example: After "Transfer confirmed," a server crash won't lose the transaction.

---

### Q3: What are indexes? What types of indexes exist in SQL databases?

**Answer:**
An **index** is a data structure that speeds up data retrieval operations at the cost of additional storage and write overhead.

**Without index:** Full table scan — O(n) rows examined.
**With index:** O(log n) lookup (B-tree) or O(1) (hash index).

**Types of indexes:**

**1. B-Tree Index (default in most DBs):**
- Balanced tree structure — works for equality (`=`) AND range queries (`>`, `<`, `BETWEEN`, `LIKE 'abc%'`).
- PostgreSQL, MySQL InnoDB default.

**2. Hash Index:**
- Maps key → hash bucket. Only for equality (`=`) lookups — cannot do range queries.
- Faster for exact lookups than B-tree.

**3. Composite Index:**
- Index on multiple columns: `CREATE INDEX idx ON orders(user_id, created_at)`.
- **Left-prefix rule**: Can be used for queries on `(user_id)` or `(user_id, created_at)` but NOT on `(created_at)` alone.

**4. Unique Index:**
- Enforces uniqueness constraint. Automatically created for `PRIMARY KEY` and `UNIQUE` columns.

**5. Full-Text Index:**
- For searching text content. Tokenizes text, stores inverted index.
- Used with `MATCH … AGAINST` in MySQL, `to_tsvector` in PostgreSQL.

**6. Partial Index (PostgreSQL):**
- Index only a subset of rows: `CREATE INDEX ON orders(user_id) WHERE status = 'PENDING'`.
- Smaller, faster for specific queries.

**When to add indexes:**
- Columns used in `WHERE`, `JOIN ON`, `ORDER BY`, `GROUP BY` clauses.
- Do NOT over-index: each index slows down inserts/updates.

---

### Q4: What is database normalization? Explain 1NF, 2NF, 3NF.

**Answer:**
**Normalization** is the process of organizing a database to reduce data redundancy and improve data integrity.

**1NF (First Normal Form):**
- Each column must contain atomic (indivisible) values.
- No repeating groups or arrays.
- Each row must be unique (has a primary key).

*Violates 1NF:*
```
| Order_ID | Products           |
| 1        | "Book, Pen, Ruler" |  ← Multiple values in one column
```
*After 1NF:*
```
| Order_ID | Product |
| 1        | Book    |
| 1        | Pen     |
| 1        | Ruler   |
```

**2NF (Second Normal Form):**
- Must be in 1NF.
- No partial dependency: Every non-key column must depend on the ENTIRE primary key (applies when composite key is used).

**3NF (Third Normal Form):**
- Must be in 2NF.
- No transitive dependency: Non-key columns must not depend on other non-key columns.

*Violates 3NF:*
```
| Employee_ID | Department_ID | Department_Name |
```
`Department_Name` depends on `Department_ID` (not on `Employee_ID`) → transitive dependency.

*After 3NF:* Split into `Employee (Employee_ID, Department_ID)` and `Department (Department_ID, Department_Name)`.

**Denormalization:** Intentionally violating normalization for performance (fewer JOINs). Common in data warehouses, read-heavy applications.

---

### Q5: What is the difference between a primary key, foreign key, and unique key?

**Answer:**

| Feature | Primary Key | Foreign Key | Unique Key |
|---|---|---|---|
| Purpose | Uniquely identify each row | Link to another table's primary key | Enforce uniqueness (non-PK) |
| NULL allowed? | Never | Depends on constraint | Sometimes (NULL is not unique) |
| Multiple per table? | Only one | Multiple allowed | Multiple allowed |
| Index created | Yes (clustered in InnoDB) | Yes (usually) | Yes |
| Example | `user_id` in Users table | `user_id` in Orders table referencing Users | `email` column (only one per user) |

**Foreign Key constraints:**
- `ON DELETE CASCADE`: Delete child rows when parent deleted.
- `ON DELETE SET NULL`: Set foreign key to NULL when parent deleted.
- `ON DELETE RESTRICT`: Prevent parent deletion if children exist.

---

### Q6: What is the N+1 query problem and how do you solve it?

**Answer:**
The **N+1 problem** occurs when an application executes 1 query to fetch N records, then executes N additional queries to fetch related data for each record — instead of fetching all in 1-2 queries.

**Example (bad):**
```java
// Query 1: Fetch all orders
List<Order> orders = orderRepo.findAll();  // SELECT * FROM orders → N rows

// N queries: For each order, fetch its items
for (Order order : orders) {
    List<Item> items = itemRepo.findByOrderId(order.getId());  // N separate SELECTs
}
// Total: 1 + N queries!
```

**Solutions:**

**1. JOIN (eager loading):**
```sql
SELECT o.*, i.* FROM orders o
JOIN order_items i ON o.id = i.order_id
```
Fetches everything in one query.

**2. ORM Eager Loading:**
```java
// JPA/Hibernate: use JOIN FETCH
@Query("SELECT o FROM Order o JOIN FETCH o.items")
List<Order> findAllWithItems();
```

**3. Batch loading:**
```java
// Fetch all order IDs, then one query for all items
List<Long> orderIds = orders.stream().map(Order::getId).collect(toList());
List<Item> allItems = itemRepo.findByOrderIdIn(orderIds);  // 1 query with IN clause
// Group items by order ID in memory
```

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
