# 🗣️ Theory — Databases in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you connect to a database in Go?"

> *"Go's standard library has the `database/sql` package which provides a generic database interface. To use a specific database, you import a driver package as a side-effect — like `import _ 'github.com/go-sql-driver/mysql'`. The driver registers itself in `init()` and then `sql.Open('mysql', connectionString)` gives you a `*sql.DB`. That `*sql.DB` is actually a connection pool, not a single connection — it manages connections for you. You should tune it with `SetMaxOpenConns`, `SetMaxIdleConns`, and `SetConnMaxLifetime` based on your database server's limits and your traffic."*

---

## Q: "What is the difference between `database/sql` and GORM?"

> *"`database/sql` is the standard library's low-level interface — you write raw SQL, scan results into struct fields manually, and it gives you full control. It's verbose but predictable. GORM is an ORM — Object Relational Mapper — that lets you work with structs directly. `db.Create(&user)`, `db.Where('age > ?', 18).Find(&users)`. GORM generates the SQL for you, handles associations, and has auto-migration. The trade-off is that GORM adds abstraction and can generate inefficient queries or behave unexpectedly with complex queries. Many teams use GORM for most things and drop down to raw SQL for performance-sensitive queries."*

---

## Q: "What are transactions and how do you use them in Go?"

> *"A transaction groups multiple database operations into a single atomic unit — either all of them succeed, or if any one fails, all of them are rolled back. In `database/sql`, you call `db.Begin()` to get a `*sql.Tx`, then run all your queries through that `Tx` instead of through `db`. If everything succeeds, you call `tx.Commit()`. If anything goes wrong, you call `tx.Rollback()`. The idiomatic pattern is to defer a rollback immediately after Begin — `defer tx.Rollback()` — then commit at the end. If you rollback after a successful commit, it's a no-op, so the defer is safe regardless of outcome."*

---

## Q: "What are prepared statements and why do they matter?"

> *"A prepared statement is a SQL query that's compiled and cached by the database server. You send the query template once, and then execute it many times with different parameters. They give you two benefits. Performance: the query plan is computed once and reused. Security: parameters are always treated as data, never as SQL — which completely prevents SQL injection. In Go, you call `db.Prepare('SELECT * FROM users WHERE id = ?')` to get a `*sql.Stmt`, then call `stmt.QueryRow(userId)` repeatedly. Always use prepared statements or parameterized queries — never string-concatenate user input into SQL."*

---

## Q: "How does connection pooling work in Go's `database/sql`?"

> *"`sql.DB` maintains a pool of database connections automatically. When you call a query method, it checks out a connection from the pool, runs the query, and returns it. The key settings are: `SetMaxOpenConns` — the upper limit on connections to the database; `SetMaxIdleConns` — how many connections to keep open even when idle; and `SetConnMaxLifetime` — how long a connection can be reused before being replaced. If all connections are in use and you've hit the max, queries queue up and wait. Getting these numbers wrong is a common performance issue — too few connections creates a bottleneck, too many can overwhelm the database."*

---

## Q: "How do you prevent SQL injection in Go?"

> *"Always, always use parameterized queries — also called prepared statements with placeholders. Instead of building a string like `'SELECT * FROM users WHERE name = ' + userInput`, you write `db.QueryRow('SELECT * FROM users WHERE name = ?', userInput)`. The `?` placeholder means the database treats userInput as a literal value, not executable SQL, regardless of what the user put in. If you use GORM, its `Where` clause with `?` placeholders is safe. String formatting SQL — `fmt.Sprintf` into queries — is the dangerous pattern. Even if you trust your data, make parameterized queries a habit."*

---

## Q: "How do you handle the 'record not found' case differently from other database errors?"

> *"In `database/sql`, when a query returns no rows, `row.Scan()` returns `sql.ErrNoRows` — a special sentinel error that means 'nothing matched your query'. You check for this specifically: `if err == sql.ErrNoRows { return nil, nil }` — no record found isn't really an error, it's just an empty result. Any other error from Scan is a real database problem. In GORM, you check `errors.Is(result.Error, gorm.ErrRecordNotFound)`. This distinction matters for your API — a 404 Not Found is the right response when a record doesn't exist, a 500 is for actual database failures."*
