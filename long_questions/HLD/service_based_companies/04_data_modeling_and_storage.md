# High-Level Design (HLD): Data Modeling and Storage

For service companies doing enterprise transformations, solid foundational data modeling concepts (specifically RDBMS) are incredibly important.

## 1. Explain Database Normalization (1NF, 2NF, 3NF). Why do we do it?
**Answer:**
Normalization is the process of organizing database designs to minimize data redundancy and prevent data modification anomalies (Insertion, Deletion, Update anomalies).
*   **1NF (First Normal Form):** Every cell holds a single, atomic value. No repeating groups or arrays as columns. Every table must have a Primary Key.
*   **2NF (Second Normal Form):** Must be in 1NF. Every non-prime attribute must be fully functionally dependent on the *entire* primary key (this strictly applies when there is a composite primary key). There can be no partial dependencies.
*   **3NF (Third Normal Form):** Must be in 2NF. There can be no transitive dependencies. Non-prime attributes must depend ONLY on the primary key, and not on other non-prime attributes. (e.g., moving "ZipCode -> City, State" into a separate table from "User").
*   *Why?* To maintain massive enterprise datasets efficiently without data duplication causing inconsistencies during `UPDATE` operations.

## 2. What are ACID properties in a Database?
**Answer:**
ACID guarantees that database transactions are processed reliably, critical for things like banking software.
*   **Atomicity:** Transactions are "all or nothing." If a transaction has 5 steps, and step 4 fails, the entire transaction rolls back. No partial executions.
*   **Consistency:** A transaction must shift the database from one valid state to another, strictly respecting all predefined rules, constraints, and triggers.
*   **Isolation:** Concurrent transactions executing at the same time must not interfere with each other. The result must be as if they executed sequentially. (Managed via Lock levels and Isolation Levels like Read Committed, Serializable).
*   **Durability:** Once a transaction is committed, it will remain permanently in the database, even in the event of a power loss, crash, or error (usually achieved via transaction logs / write-ahead logging).

## 3. What is the difference between OLTP and OLAP systems?
**Answer:**
*   **OLTP (Online Transaction Processing):**
    *   *Purpose:* Running the core, day-to-day business operations (e.g., e-commerce orders, banking transactions).
    *   *Characteristics:* High volume of short, fast, atomic transactions (INSERT/UPDATE/DELETE). Highly normalized schema (3NF) to ensure data integrity and fast writes.
*   **OLAP (Online Analytical Processing):**
    *   *Purpose:* Business Intelligence, reporting, and data mining (Data Warehouses).
    *   *Characteristics:* Low volume of massively complex queries (complex `SELECT` with `JOIN` and Aggregations). Queries touch millions of rows. Highly *denormalized* schemas (Star Schema / Snowflake Schema) optimized for read speed over write efficiency.

## 4. How do you optimize slow SQL queries in a relational database?
**Answer:**
1.  **Use Indexes:** Ensure columns used in `WHERE`, `JOIN`, and `ORDER BY` clauses are indexed. 
2.  **Analyze Execution Plans (`EXPLAIN`):** Use the database's `EXPLAIN PLAN` command to see if it is doing a Full Table Scan (bad) vs an Index Scan (good).
3.  **Avoid `SELECT *`:** Only select the exact columns needed to reduce memory and I/O overhead.
4.  **Optimize Joins:** Ensure tables are joined on indexed columns.
5.  **Use Pagination (`LIMIT` / `OFFSET`):** Don't pull 10,000 records simultaneously.
6.  **Archiving:** Move old, historical data out of the live transaction tables into archive tables or a Data Warehouse to keep live table sizes small.

## 5. What are Database Indexes, what data structures are used, and what are their drawbacks?
**Answer:**
*   **What it is:** An index is an auxiliary data structure that speeds up data retrieval operations on a table at the cost of additional storage and reduced write speed.
*   **Data Structures:** Usually implemented using **B-Trees (Balanced Trees)**, which allow for efficient O(log N) searching, sequential access, insertions, and deletions. Some DBs use Hash indexes for exact-match lookups.
*   **Clustered vs Non-Clustered:** Clustered index determines the actual physical order of data on disk (only one allowed per table). Non-clustered index stores pointers to the actual data.
*   **Drawbacks:** Every time an `INSERT`, `UPDATE`, or `DELETE` occurs, the database must also update the Index tree structure. Thus, applying too many indexes drastically slow down write performance.
