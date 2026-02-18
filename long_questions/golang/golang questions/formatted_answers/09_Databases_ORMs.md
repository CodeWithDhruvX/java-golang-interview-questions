# 🟣 **161–180: Databases and ORMs**

### 161. How do you connect to a PostgreSQL database in Go?
"I use the standard `database/sql` package with a driver like `lib/pq` or `pgx`.

I call `sql.Open("postgres", connStr)`.
Crucially, this **does not** create a connection. It just sets up the pool.
I always call `db.Ping()` immediately after to verify the credentials and network reachability. If `Ping` succeeds, I know the app is ready to serve traffic."

#### Indepth
`connStr` often includes complex parameters like `sslmode=disable`, `connect_timeout=10`, or `search_path=public`. Use `url.Parse` or a library helper to build this string safely, rather than manual concatenation, to avoid formatting errors with special characters in passwords.

---

### 162. What is the difference between `database/sql` and GORM?
"`database/sql` is the low-level, standard interface.
It gives me full control over the SQL. I have to manually scan rows into variables (`rows.Scan(&name)`). It’s tedious but fast and explicit.

**GORM** is an ORM. It abstracts the SQL away.
I can just call `db.Save(&user)`. It handles the `INSERT`, the ID generation, and the timestamps. I use GORM for rapid prototyping, but I often drop down to raw SQL for complex reports."

#### Indepth
ORMs in Go often rely heavily on **Reflection** (`reflect` package) to map structs to tables. This adds a CPU overhead compared to code-generated mappers (like `sqlc`) or manual scanning. For write-heavy logic (INSERTs), this overhead is negligible compared to I/O, but for reading 10k rows in a tight loop, `database/sql` is significantly faster.

---

### 163. How do you handle SQL injections in Go?
"I strictly rely on **parameterized queries**.

I never concatenate strings like `'SELECT * FROM users WHERE name = ' + input`.
Instead, I use placeholders: `$1` (Postgres) or `?` (MySQL).
`db.Query("SELECT ... WHERE name = $1", input)`.
The driver sends the query template and the data separately, making injection mathematically impossible."

#### Indepth
One common mistake is using `fmt.Sprintf` for table names or `ORDER BY` clauses, which cannot be parameterized. `db.Query("SELECT * FROM " + tableName)` is unsafe. If you must have dynamic tables, use an allow-list map to validate the input string against known safe values before concatenating.

---

### 164. How do you manage connection pools in `database/sql`?
"The `sql.DB` object **is** a connection pool.

I configure it carefully:
`db.SetMaxOpenConns(25)`: Prevents my app from overwhelming the DB.
`db.SetMaxIdleConns(25)`: Keeps connections hot so I don't pay the handshake cost on every request.
`db.SetConnMaxLifetime(5 * time.Minute)`: Rotating connections prevents stale-socket issues (like firewalls silently dropping idle connections)."

#### Indepth
If `MaxOpenConns` is reached, `db.Query` will **block** and wait for a connection to be returned. This is a hidden source of latency spikes. Monitor the `db.Stats().WaitCount` metric. If it's high, you need to either increase the pool size (if DB CPU allows) or optimize your query duration.

---

### 165. What are prepared statements in Go?
"A prepared statement is a pre-compiled SQL query.
`stmt, err := db.Prepare("INSERT INTO log (msg) VALUES ($1)")`.

I use it when I'm running the exact same query thousands of times in a loop.
It saves the database from parsing and planning the query every single time. It also reduces network bandwidth because I only send the parameters, not the full query text."

#### Indepth
Statements are prepared *per connection*. If your connection pool recycles a connection, the statement must be re-prepared (drivers handle this transparently, but it costs a round-trip). Some drivers support **Client-Side Statement Caching** to mitigate this. Always close statements with `defer stmt.Close()` to prevent leaking resources on the database server.

---

### 166. How do you map SQL rows to structs?
"With the standard library, it's painful. I have to manually `rows.Scan(&u.ID, &u.Name, ...)` in the exact order of columns.

I prefer using **sqlx**.
It allows me to do `db.Get(&user, "SELECT ...")`.
It uses the `db` struct tags to automatically map columns to fields. It saves me from writing boilerplate and mismatch errors."

#### Indepth
`sqlx` also supports `Scan` into a slice of structs (`db.Select(&users, ...)`). Be careful with `NULL` values. If a database column is NULL, scanning into a `string` (or `int`) will error. You must use `sql.NullString` or `*string` to handle potential nulls gracefully.

---

### 167. What are transactions and how are they implemented in Go?
"A transaction (`Tx`) ensures atomicity.

`tx, err := db.Begin()`.
I perform multiple queries using `tx.Exec` (not `db.Exec`).
Finally, I call `tx.Commit()` to save or `tx.Rollback()` to undo.
I typically use `defer tx.Rollback()` at the start. If the function panics or returns early, it automatically rolls back. If I successfully commit, the rollback does nothing."

#### Indepth
Transactions lock rows. If your business logic inside the transaction involves a slow operation (like an HTTP call to a payment gateway), you hold those DB locks for the duration of the HTTP call. This destroys database concurrency. Always keep transactions as short as possible—logic first, then DB lock, then update, then commit immediately.

---

### 168. How do you handle database migrations in Go?
"I use a dedicated tool like **golang-migrate** or **Goose**.

My migrations are versioned SQL files: `20231001_create_users.up.sql`.
I run these as part of the deployment pipeline (e.g., `migrate up`).
I never rely on ORMs like GORM to 'AutoMigrate' in production. It’s too risky—I need to know exactly what index changes or table locks are happening."

#### Indepth
Migrations should be **idempotent** if possible, but SQL often isn't. Running a migration twice usually fails ("table already exists"). To handle this in k8s, use a `Job` with `restartPolicy: OnFailure` that runs `migrate up`. Ensure your app waits for the migration to complete (using an `initContainer`) before starting.

---

### 169. What is the use of `sqlx` in Go?
"It’s an extension of the standard library.

It gives me `StructScan`, `NamedExec` (using `:name` instead of `$1`), and `Select` (for slices).
It doesn't hide the SQL like an ORM does; it just removes the tedium of mapping results. It’s the perfect middle ground between raw `database/sql` and heavy ORMs."

#### Indepth
`sqlx` mimics the standard library interface (`Query`, `Exec`), so it's easy to drop in. It also allows using maps for named queries: `db.NamedExec("INSERT ... VALUES (:name)", map[string]interface{}{"name": "Bob"})`. This is often cleaner than struct tags for partial updates.

---

### 170. What are the pros and cons of using an ORM in Go?
"**Pros**: Velocity. I can write CRUD apps in minutes. Relationships (Preload) and basic joins are handled for me.
**Cons**: Performance penalties (reflection), hidden N+1 queries, and difficulty debugging generated SQL.

Idiomatic Go tends to prefer **explicit** over implicit, so many teams start with GORM and migrate to sqlx or sqlc as the project scales."

#### Indepth
**sqlc** is a popular alternative. You write raw SQL (`-- name: GetUser :one SELECT * ...`), and it generates type-safe Go code for you. It catches syntax errors at compile time and has zero runtime overhead (no reflection). It's effectively the "Reverse ORM".

---

### 171. How would you implement pagination in SQL queries?
"I avoid `OFFSET` for large tables because it gets slower as the page number increases (it scans and discards rows).

I use **Cursor-based Pagination** (Keyset Pagination).
`WHERE id > last_seen_id LIMIT 10`.
This uses the index on `id` to jump directly to the right row. It’s O(1) regardless of whether I'm on page 1 or page 1,000,000."

#### Indepth
The downside of Cursor Pagination is that you cannot jump to "Page 10" directly, nor can you easily implement "Previous Page" without complex reverse-query logic. You also need a unique, sortable column (often usage of time + uuid) to serve as the cursor.

---

### 172. How do you log SQL queries in Go?
"I check the driver documentation.
For GORM, it’s built-in: `gorm.Config{Logger: logger.Default.LogMode(logger.Info)}`.
For `database/sql`, I wrap the driver with a logging hook (like **sqlhooks**).
It intercepts every `Exec/Query` call, logs the SQL statement, arguments, and execution time. It’s invaluable for debugging slow queries."

#### Indepth
Be careful not to log sensitive data (PII/passwords) in the SQL arguments. Custom loggers should have a sanitization step. Also, logging every query in production will emit massive logs. Use sampling or only log queries that exceed a duration threshold (Slow Query Log).

---

### 173. What is the N+1 problem in ORMs and how to avoid it?
"It’s when I fetch N items, and then for *each* item, I accidentally trigger another query to fetch a related record.
1 query for Users + 100 queries for their Avatars.

I avoid it by **Eager Loading**.
In GORM: `db.Preload("Avatar").Find(&users)`.
This executes exactly 2 queries: one for users, and one `IN (...)` query for all avatars. Speedup is massive."

#### Indepth
Eager loading works well, but watch out for memory usage. Loading 10,000 users and their 100,000 avatars into RAM might crash the pod (`OOMKilled`). For bulk processing, disable preloading and process in batches using a cursor or `FindInBatches`.

---

### 174. How do you implement caching for DB queries in Go?
"I use the **Cache-Aside** pattern with Redis.

1.  Check Redis for the key `user:123`.
2.  If found (Hit), return it.
3.  If missing (Miss), query Postgres.
4.  Serialize the result (JSON/Protobuf) and write to Redis with a TTL (e.g., 5 mins).
5.  Return it.
This protects the primary database from read spikes."

#### Indepth
The "Thundering Herd" problem occurs when a cache key expires and 1000 requests hit the DB simultaneously. To solve this, use **singleflight** (from `golang.org/x/sync/singleflight`). It merges duplicate in-flight calls so only *one* DB query runs, and the result is shared with all 1000 waiters.

---

### 175. How do you write custom SQL queries using GORM?
"GORM has a `Raw` method for when the query builder is too restrictive.

`db.Raw("SELECT name, count(*) FROM users GROUP BY name").Scan(&result)`.
I use this for complex reports, CTEs (Common Table Expressions), or specific window functions that GORM doesn't support natively."

#### Indepth
Using `Raw` returns `*gorm.DB` but creates a potential separation between your Go structs and the result set. If you select fields that don't match the struct, they are zero-valued. Always verify that your Raw logic aliases columns correctly to match the struct field names (`SELECT count(*) as total ...`).

---

### 176. How do you handle one-to-many and many-to-many relationships in GORM?
"I define them in the struct tags.

`type User struct { Orders []Order }` (One-to-Many).
`type User struct { Groups []*Group \`gorm:"many2many:user_groups;"\` }` (Many-to-Many).
GORM handles the join table (`user_groups`) automatically.
When I save a User, GORM automatically inserts the records into the join table."

#### Indepth
Auto-save associations can be dangerous. If you update a User struct and accidentally zero out the `Groups` slice, GORM might define this as "remove all groups" depending on configuration (`FullSaveAssociations`). I often disable this feature (`db.Omit("Groups").Save(&user)`) and manage relationships explicitly to avoid data loss.

---

### 177. How would you structure your database layer in a Go project?
"I use the **Repository Pattern**.

I define an interface `UserRepository`.
`type UserRepository interface { Get(id string) (*User, error) }`.
The implementation (`PostgresUserRepository`) holds the `*sql.DB`.
This decouples the Service layer from the database. I can swap Postgres for a Mock in unit tests without changing a single line of business logic."

#### Indepth
Repository interfaces should be defined by the **consumer** (the Domain/Service layer), not the implementer. This follows the Dependency Inversion Principle. `package service` defines `type UserRepo interface {...}`, and `package postgres` implements it. This prevents the domain from depending on `package postgres`.

---

### 178. What is context propagation in database calls?
"It means threading the `context.Context` from the HTTP handler down to the DB driver.

I use `db.QueryContext(ctx, ...)` instead of `db.Query`.
If the user cancels the request (closes the browser), the `ctx` is canceled. The DB driver sees this and **terminates** the running query on the database server. It prevents 'ghost queries' from hogging resources."

#### Indepth
Most modern drivers support this, but not all operations are cancellable immediately (e.g., if the DB is stuck in a heavy CPU loop vs waiting on I/O). However, it frees the connection in the pool on the Go side immediately. Always propagate context.

---

### 179. How do you handle long-running queries or timeouts?
"I wrap the parent context with a Timeout.

`ctx, cancel := context.WithTimeout(req.Context(), 5 * time.Second)`
`defer cancel()`
I pass this `ctx` to the query. If the DB takes >5s, the query returns `context.DeadlineExceeded` immediately. It ensures my API handles failures gracefully instead of hanging indefinitely."

#### Indepth
Postgres has a server-side setting `statement_timeout`. It's good practice to set this at the session level via connection parameters as a fallback safety net, in case the application fails to cancel the context correctly.

---

### 180. How do you write unit tests for code that interacts with the DB?
"I use **go-sqlmock**.

It mocks the `database/sql` driver. I tell it:
'Expect a query matching `SELECT * FROM users` and return these fixed rows.'
This allows me to test my repository logic (e.g., proper error handling, row mapping) without needing a running Postgres instance. For integration tests, I spin up a real DB container."

#### Indepth
`go-sqlmock` is great for testing *that your code calls the library correctly* (args, order). It does **not** test if your SQL is valid syntax or matches your schema. For that, you need Integration Tests using **testcontainers-go**, which provides a real ephemeral Postgres for every test run.
