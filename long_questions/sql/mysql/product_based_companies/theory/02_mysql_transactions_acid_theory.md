# 🗣️ Theory — MySQL Transactions & ACID
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Explain ACID properties in MySQL."

> *"ACID is the set of guarantees that make a database transaction reliable. Atomicity means a transaction is all-or-nothing — if I'm transferring money from account A to account B, both the debit and the credit must succeed together, or neither happens. If the server crashes halfway through, InnoDB's redo log ensures the partial transaction is rolled back on recovery. Consistency means the database always moves from one valid state to another valid state — all constraints, foreign keys, and triggers are enforced across the entire transaction. Isolation means concurrent transactions don't interfere with each other — one transaction shouldn't see the intermediate, uncommitted state of another. And Durability means once you COMMIT, the data survives crashes — InnoDB achieves this by writing to its redo log (a write-ahead log) before confirming the commit to the client."*

---

## Q: "What are the four isolation levels? Which is MySQL's default?"

> *"MySQL's default is REPEATABLE READ. The four levels from weakest to strongest are: READ UNCOMMITTED — the weakest, lets you read uncommitted changes from other transactions — dirty reads are possible, almost never used in production. READ COMMITTED — you only see committed data, but within the same transaction you might get different values for the same query if another transaction commits in between — that's a non-repeatable read. REPEATABLE READ — MySQL's default — guarantees that within a transaction, the same SELECT gives the same result every time, even if other transactions commit changes. InnoDB achieves this with MVCC — it creates a consistent snapshot at the start of the transaction. SERIALIZABLE — the strongest — all transactions appear to run sequentially, no concurrency anomalies possible, but uses heavy locking and kills throughput. In practice, REPEATABLE READ with InnoDB's MVCC implementation gives very strong guarantees — it even prevents most phantom reads, which technically REPEATABLE READ doesn't have to prevent by the SQL standard."*

---

## Q: "What is MVCC — Multi-Version Concurrency Control?"

> *"MVCC is how InnoDB achieves high concurrency without readers blocking writers or writers blocking readers. Instead of locking a row when you read it, InnoDB keeps multiple versions of each row — the current version and historical versions in an undo log. When a transaction reads a row, InnoDB checks the transaction's snapshot timestamp and uses the version of the row that was committed at or before when the transaction started. If another transaction has modified the row but not committed, or committed after our transaction started, we see the older version from the undo log instead. This means reads never block writes and writes never block reads — which is great for concurrency. The tradeoff is that long-running transactions prevent the undo log from being purged, because older row versions are still needed for those transactions' consistent reads. You can monitor this with SHOW ENGINE INNODB STATUS and watch the 'history list length' — if it grows very large, you have long-running transactions holding back the purge."*

---

## Q: "What are dirty reads, non-repeatable reads, and phantom reads?"

> *"These are the three concurrency anomalies that isolation levels are designed to prevent. A dirty read is when you read data that another transaction has written but not yet committed — you're reading potentially temporary, invalid data that might be rolled back. Non-repeatable read is when you execute the same SELECT twice in the same transaction and get different values because another transaction committed an UPDATE in between — the same row returned different data. A phantom read is when you execute the same range query twice in the same transaction and get different ROWS — because another transaction committed an INSERT that falls within your range. Each higher isolation level prevents more of these: SERIALIZABLE prevents all three, REPEATABLE READ prevents non-repeatable reads and mostly prevents phantoms, READ COMMITTED prevents dirty reads only, and READ UNCOMMITTED prevents none of them."*

---

## Q: "How does row-level locking work in InnoDB?"

> *"InnoDB locks at the row level by default, which gives much better concurrency than table-level locking — two concurrent transactions can modify different rows in the same table without blocking each other. There are two main lock types: a shared lock — acquired with LOCK IN SHARE MODE or FOR SHARE — allows multiple readers simultaneously but blocks exclusive writers. An exclusive lock — acquired with FOR UPDATE or implicitly by INSERT/UPDATE/DELETE — blocks all other readers and writers on that row. The important pattern for correctness is: if you read a row intending to update it based on what you read, you must use SELECT ... FOR UPDATE to acquire an exclusive lock at read time. Without it, two concurrent transactions could both read the same balance, both decide to subtract the same amount, and write the same wrong result — that's a lost update anomaly."*

---

## Q: "What is a deadlock and how does MySQL handle it?"

> *"A deadlock is when two or more transactions are each waiting for a lock that the other holds, creating a circular wait with no way out. Classic example: Transaction A locks row 1 and wants row 2. Transaction B locks row 2 and wants row 1. Neither can proceed. MySQL's InnoDB detects this automatically by periodically checking the lock wait graph for cycles. When detected, MySQL picks one transaction as the 'victim' — typically the one with the least amount of undo data, meaning the cheapest to rollback — and rolls it back with error 1213: 'Deadlock found when trying to get lock; try restarting transaction'. The application should catch this error and retry. Prevention strategies: always acquire locks in the same global order across all code paths — if Transaction A always locks the lower ID row first, and Transaction B does too, circular waits can't happen. Also, keep transactions short and fast — the less time spent holding locks, the smaller the window for deadlocks."*

---

## Q: "What is the difference between optimistic and pessimistic locking?"

> *"Pessimistic locking assumes conflict WILL happen, so it locks the resource immediately on read — before you know you'll need to write. You use SELECT ... FOR UPDATE which acquires an exclusive lock right then. Nobody else can modify that row until you COMMIT. This is safe for high-contention scenarios like bank transfers — you guarantee no interference — but it reduces concurrency because locks are held for the duration of the transaction. Optimistic locking assumes conflict WON'T happen, so it does no locking on read. Instead, you add a version column to the table. You read the row and note the version number. When you write, you include a WHERE version = N condition. If affected_rows is 1, you won, no one else changed it. If it's 0, someone else changed the row and incremented the version — you retry. This is great for low-contention workloads — reads are fully concurrent, and conflicts are handled only on the rare occasion they occur."*

---

## Q: "What is a gap lock in InnoDB and why does it matter?"

> *"A gap lock is a lock on the gap between index values — not the records themselves, but the 'space' between them. InnoDB uses gap locks in REPEATABLE READ isolation level to prevent phantom reads. If I run SELECT ... FOR UPDATE WHERE id BETWEEN 10 AND 20, InnoDB locks not just the rows with IDs 10 through 20 but also the gaps — meaning no other transaction can INSERT a row with an ID that would fall in that range. This prevents a second SELECT in my transaction from returning new 'phantom' rows that didn't exist before. The practical implication is that gap locks can cause unexpected blocking: if Transaction A locks a range, Transaction B can't insert even though it's inserting a new row that doesn't exist yet. This is most noticeable with range queries and unique index lookups where the searched value doesn't exist. If you switch to READ COMMITTED isolation level, gap locks are disabled — you get only record locks, more concurrency, but you accept the possibility of phantom reads."*

