# 🗣️ Theory — MySQL Basics
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is MySQL? Why would you choose it?"

> *"MySQL is an open-source relational database management system — it stores data in structured tables and uses SQL to query and manipulate it. I'd choose MySQL because it's battle-tested at scale — companies like Facebook, Twitter, and YouTube have run it at massive scale. It's ACID-compliant via InnoDB, supports replication out of the box, has excellent tooling like Workbench and EXPLAIN, and integrates with every major programming language and framework. It also has one of the best cost-to-performance ratios of any database option."*

---

## Q: "What's the difference between DDL and DML?"

> *"DDL is Data Definition Language — it defines the structure of the database. Commands like CREATE, ALTER, DROP, and TRUNCATE are DDL. The key thing about DDL is that it auto-commits — it cannot be rolled back in a transaction. DML is Data Manipulation Language — it works with the actual data inside tables. SELECT, INSERT, UPDATE, DELETE are DML. DML operations ARE transactional, so you can COMMIT or ROLLBACK them. There's also DCL for access control — GRANT and REVOKE — and TCL for transaction control — COMMIT, ROLLBACK, SAVEPOINT."*

---

## Q: "What is the difference between DELETE, TRUNCATE, and DROP?"

> *"All three remove data, but at very different levels. DELETE removes specific rows — you can use a WHERE clause, it fires triggers, it can be rolled back inside a transaction, and it logs every deleted row in the redo log, so it's slower. TRUNCATE removes all rows in a single fast operation by deallocating data pages — it resets the AUTO_INCREMENT counter, cannot be rolled back in MySQL, and does NOT fire triggers. DROP removes the entire table — structure, data, indexes, everything — it's completely irreversible. I use DELETE for precise row-level removal, TRUNCATE to reset a table in dev/testing, and DROP to permanently eliminate a table."*

---

## Q: "How does a primary key differ from a unique key?"

> *"Both enforce uniqueness, but they serve different purposes. A primary key is the main identifier for a row — it must be unique AND cannot be null. Every InnoDB table has exactly one primary key. Internally, InnoDB organizes the entire table's data as a B+Tree clustered on the primary key — the primary key IS the data structure. A unique key also enforces uniqueness but it allows NULL values — and MySQL even allows multiple NULL values in a unique column since NULL is not considered equal to NULL in SQL. A table can have many unique keys but only one primary key. I always define a primary key — without one, InnoDB creates a hidden 6-byte row ID that you can't work with."*

---

## Q: "What is normalization? Explain 1NF, 2NF, 3NF."

> *"Normalization is the process of organizing a database schema to reduce data redundancy and prevent update anomalies. If a customer's name appears in 10,000 order rows and they change their name, you'd need 10,000 updates — that's a normalization problem. First normal form means every column holds an atomic, single value — no comma-separated lists, no repeating groups. Second normal form means you're in 1NF and every non-key column depends on the ENTIRE primary key — if you have a composite key and some column only depends on part of it, that's a partial dependency you need to eliminate. Third normal form means you're in 2NF and there are no transitive dependencies — where one non-key column determines another non-key column. I aim for 3NF in transactional OLTP systems. For analytics and reporting, I actually denormalize deliberately to avoid expensive JOINs."*

---

## Q: "What is the difference between WHERE and HAVING?"

> *"Both filter data, but at completely different stages of query execution. WHERE runs BEFORE any grouping happens — it filters individual rows. Because it runs early, it can use indexes efficiently. HAVING runs AFTER the GROUP BY aggregation — it filters the grouped results. So if I want employees with salary over 50,000, that's a WHERE clause — it filters rows. If I want departments where the average salary is over 50,000, that's HAVING — it filters groups after the aggregation is computed. The common mistake is putting aggregate conditions in WHERE — that causes an error. A useful mental model: WHERE answers 'which rows?', HAVING answers 'which groups?'."*

---

## Q: "What is NULL in MySQL and why is it tricky?"

> *"NULL in SQL means 'unknown' or 'not applicable' — it's not zero, not an empty string, it's the absence of a value. The tricky part is how comparisons work: NULL equals nothing — not even itself. NULL = NULL is NULL, not TRUE. So you can't useEquals to check for NULL — you must use IS NULL or IS NOT NULL. Aggregate functions like SUM, AVG, COUNT(column) all silently ignore NULLs — only COUNT(*) counts every row including those with nulls. This trips people up in interview questions all the time. Another gotcha: NOT IN with a subquery that can return NULLs — if the subquery returns any NULL, the entire NOT IN check returns no rows. You have to either add WHERE column IS NOT NULL in the subquery, or switch to NOT EXISTS which handles NULLs correctly."*

---

## Q: "What is the difference between CHAR and VARCHAR?"

> *"CHAR is fixed-length: if you declare CHAR(10) and store 'Hi', MySQL pads it to 10 characters with spaces. It always takes exactly n bytes. VARCHAR is variable-length: if you declare VARCHAR(255) and store 'Hi', MySQL stores just 2 bytes plus a 1 or 2 byte length prefix. I use CHAR for columns where every value is the same length — country codes, state abbreviations, fixed-format flags — because CHAR is marginally faster for fixed-size lookups. I use VARCHAR for everything else — names, emails, descriptions — because it saves storage. One interesting edge case: in MySQL's MEMORY storage engine, VARCHAR is treated as CHAR internally, so you don't get the storage savings for in-memory temp tables."*

---

## Q: "How does AUTO_INCREMENT work in MySQL?"

> *"AUTO_INCREMENT is a column attribute — typically on the primary key — that automatically assigns a unique, ever-increasing integer to each new row. You don't specify the value on INSERT; MySQL handles it. After inserting a row, you can retrieve the assigned ID with LAST_INSERT_ID(). An important thing to know is that in MySQL versions before 8.0, the AUTO_INCREMENT counter was stored only in memory and rebuilt on restart — so if you inserted row ID 100, deleted it, and restarted, the next insert could reuse ID 100. MySQL 8.0 fixed this by persisting the counter in the redo log. For distributed systems where you need globally unique IDs without coordinating through a single sequence, I use UUIDs or application-generated IDs instead. And always use BIGINT for AUTO_INCREMENT on tables that could grow large — INT maxes out at about 2 billion."*

