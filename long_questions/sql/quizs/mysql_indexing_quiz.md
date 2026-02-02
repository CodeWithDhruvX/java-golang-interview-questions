# MySQL Indexing & Performance Quiz

## 1. Questions

### Q1: Clustered vs Secondary Index
**Question:** In InnoDB, what is the fundamental difference between a Clustered Index and a Secondary (Non-Clustered) Index? How does looking up data via a Secondary Index work?

### Q2: Leftmost Prefix Rule
**Scenario:** You have a composite index on table `Users` columns `(lastname, firstname, dob)`.
**Question:** Which of the following queries will use the index?
1.  `SELECT * FROM Users WHERE lastname = 'Smith'`
2.  `SELECT * FROM Users WHERE firstname = 'John'`
3.  `SELECT * FROM Users WHERE lastname = 'Smith' AND firstname = 'John'`
4.  `SELECT * FROM Users WHERE dob = '1990-01-01'`

### Q3: Covering Index
**Question:** What is a "Covering Index" and why is it faster? Give an example involving a query `SELECT username FROM Users WHERE user_id = 5`.

### Q4: Low Cardinality Columns
**Question:** You have a column `gender` with values 'M' and 'F' in a table of 1 million rows. Should you place an index on this column? Why or why not?

### Q5: EXPLAIN Output
**Question:** When running `EXPLAIN SELECT ...`, what does `type: ALL` mean? Is it good or bad?

---

## 2. Answers & Explanations

### Answer 1: Clustered vs Secondary
**Answer:**
*   **Clustered Index**: This **is** the table itself. The data rows are stored physically in the order of the clustered index (usually the Primary Key). Queries using the PK retrieve the data directly.
*   **Secondary Index**: Contains the indexed column values and the **Primary Key** pointing to the row.
*   **Lookup Process**: When you search via a Secondary Index, MySQL finds the Primary Key, then performs a second lookup (Traversal) in the Clustered Index to get the actual row data. This is often called "referencing" or "bookmark lookup".

### Answer 2: Leftmost Prefix Rule
**Answer:**
1.  `lastname = 'Smith'` -> **YES**. Matches the start (left) of the index.
2.  `firstname = 'John'` -> **NO**. Skips the first column (`lastname`).
3.  `lastname = ... AND firstname = ...` -> **YES**. Matches the sequence from the left.
4.  `dob = ...` -> **NO**. Skips `lastname` and `firstname`.
**Explanation:** A B-Tree composite index is sorted first by Col1, then Col2, then Col3. You cannot jump to the middle of the sort order without filtering by the previous columns.

### Answer 3: Covering Index
**Answer:**
A **Covering Index** occurs when all the fields required by a query are contained within the index itself, so the database engine **does not need to look up the actual table row** (skipping the secondary-to-clustered lookup step).
**Example:**
If you have an index on `(user_id, username)` and run `SELECT username FROM Users WHERE user_id = 5`, MySQL gets `username` directly from the index tree. This is significantly faster purely memory-based operation in many cases.

### Answer 4: Low Cardinality & Indexing
**Answer:** No, you generally should **not** index it.
**Why?**
An index works best when it significantly reduces the number of rows to scan. If a column has very low cardinality (e.g., only 'M'/'F' split 50/50), using the index means reading 50% of the table. The overhead of reading the index *plus* the table lookups is often higher than simply doing a full table scan. The optimizer will likely ignore the index anyway.

### Answer 5: EXPLAIN type: ALL
**Answer:**
`type: ALL` means a **Full Table Scan**.
*   **Verdict**: usually **BAD** for large tables.
*   **Meaning**: MySQL has to read every single row in the table to find the matching data because no suitable index was found.
