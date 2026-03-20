# 🟢 Go Theory Questions: 161–180 Databases (SQL, NoSQL, Patterns)

## 161. How do you use `database/sql` in Go?

**Answer:**
`database/sql` is a generic interface around SQL (or SQL-like) databases. It doesn't contain the actual database logic itself; instead, it provides a standard API that drivers (like `pq` for Postgres or `mysql` for MySQL) implement.

You open a connection with `sql.Open()`, which gives you a **DB Handle**. Crucially, this handle isn't a single connection—it’s a concurrent-safe **connection pool**.

When you query, you generally use `Query` for multiple rows, `QueryRow` for a single result, or `Exec` for inserts/updates. The key trade-off is that it’s somewhat "low-level"—you have to manually scan results into variables `rows.Scan(&name, &age)`, which can be tedious without helper libraries.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `database/sql` in Go?

**Your Response:** "database/sql is a generic interface around SQL-like databases. It doesn't contain actual database logic itself; instead, it provides a standard API that drivers implement.

You open a connection with sql.Open(), which gives you a DB Handle. When you query, you generally use Query for multiple rows, QueryRow for a single result, or Exec for inserts/updates. The key trade-off is that it's somewhat 'low-level'—you have to manually scan results into variables, which can be tedious without helper libraries."

---

## 162. How do you manage connection pools in Go?

**Answer:**
Go handles pooling automatically within the sql.DB object, but you must configure it for production. By default, the pool is unbounded, which is dangerous—a traffic spike could open 10,000 connections and crash your database server.

We control this with db.SetMaxOpenConns(N) to limit total connections, and db.SetMaxIdleConns(M) to keep some warm connections ready. We also set db.SetConnMaxLifetime(duration) to recycle connections periodically, which prevents issues with load balancers silently dropping idle connections.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage connection pools in Go?

**Your Response:** "Go handles pooling automatically within the sql.DB object, but you must configure it for production. By default, the pool is unbounded, which is dangerous—a traffic spike could open 10,000 connections and crash your database server.

We control this with db.SetMaxOpenConns(N) to limit total connections, and db.SetMaxIdleConns(M) to keep some warm connections ready. We also set db.SetConnMaxLifetime(duration) to recycle connections periodically, which prevents issues with load balancers silently dropping idle connections."

---

## 163. How do you write raw queries using `sqlx`?

**Answer:**
`sqlx` is a lightweight wrapper that extends the standard library without hiding it.

Its killer feature is **Struct Mapping**. Instead of manually scanning 20 columns line-by-line (`rows.Scan(&a, &b, ...)`), you just write `db.Select(&users, "SELECT * FROM users")`.

It uses struct tags (`db:"user_id"`) to automatically map the SQL columns to your struct fields.

We use it because it keeps the performance and control of raw SQL—you still write the query yourself—but removes the boilerplate code that usually leads to bugs.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write raw queries using `sqlx`?

**Your Response:** "`sqlx` is a lightweight wrapper that extends the standard library without hiding it. Its killer feature is Struct Mapping. Instead of manually scanning 20 columns line-by-line, you just write `db.Select(&users, 'SELECT * FROM users')`.

It uses struct tags (`db:"user_id"`) to automatically map SQL columns to your struct fields. This keeps performance and control of raw SQL—you still write the query yourself—but removes boilerplate code that usually leads to bugs."

---

## 164. How do you use GORM with PostgreSQL?

**Answer:**
GORM is a full-featured ORM (Object Relational Mapper) for Go.

You define your database schema using Go structs: `type User struct { gorm.Model; Name string }`. GORM can automatically create the table ("AutoMigrate") and gives you a fluent API like `db.Where("name = ?", "John").First(&user)`.

We use it for rapid development or CRUD-heavy applications where writing raw SQL for everything is overkill. The trade-off is performance and opacity—sometimes GORM generates inefficient queries, so for complex reporting or high-load paths, we often drop back down to raw SQL.

---

## 165. How do you handle transactions in Go?

**Answer:**
You start a transaction with `tx, err := db.Begin()`.

This gives you a generic `*sql.Tx` object. **Everything** you want to be part of that transaction must be executed on this `tx` object, not the original `db` handle.

The critical pattern here is utilizing `defer`. You typically defer a rollback function: `defer tx.Rollback()`. Then, if the function finishes successfully, you call `tx.Commit()` at the very end. If any error occurs or the function panics, the deferred Rollback ensures the database isn't left in a half-written state.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle transactions in Go?

**Your Response:** "You start a transaction with tx, err := db.Begin(). This gives you a generic *sql.Tx object. Everything you want to be part of that transaction must be executed on this tx object, not the original db handle.

The critical pattern here is utilizing defer. You typically defer a rollback function: `defer tx.Rollback()`. Then, if the function finishes successfully, you call tx.Commit() at the very end. If any error occurs or the function panics, the deferred Rollback ensures the database isn't left in a half-written state."

---

## 166. How do you create database migrations in Go?

**Answer:**
Go doesn't have a built-in migration tool, so we use libraries like Golang Migrate or Goose. The concept is to version-control your schema changes. You create SQL files named by timestamp: `20231001_create_users.up.sql` and `20231001_create_users.down.sql`.

The tool runs these files in order against your database. We typically run migrations as a separate step in our CI/CD pipeline before the main application starts, ensuring the code matches the database schema.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create database migrations in Go?

**Your Response:** "Go doesn't have a built-in migration tool, so we use libraries like Golang Migrate or Goose. The concept is to version-control your schema changes. You create SQL files named by timestamp: `20231001_create_users.up.sql` and `20231001_create_users.down.sql`.

The tool runs these files in order against your database. We typically run migrations as a separate step in our CI/CD pipeline before the main application starts, ensuring the code matches the database schema."

---

## 167. How do you use MongoDB with Go?

**Answer:**
We use the official MongoDB Go Driver.

It feels quite different from SQL. Instead of struct tags mapping to columns, we use `bson` tags (`bson:"user_id"`). Queries are constructed using `bson.M` (maps) or `bson.D` (ordered documents): `coll.Find(ctx, bson.M{"name": "Alice"})`.

Because Mongo is schema-less, unmarshalling is flexible—you can decode into a struct or just a raw `bson.M` map if the data shape is unpredictable. It integrates heavily with `context` for cancellation, which is vital for long-running aggregations.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use MongoDB with Go?

**Your Response:** "We use the official MongoDB Go Driver. It feels quite different from SQL. Instead of struct tags mapping to columns, we use bson tags (`bson:"user_id"`). Queries are constructed using bson.M (maps) or bson.D (ordered documents): `coll.Find(ctx, bson.M{"name": "Alice"})`.

Because Mongo is schema-less, unmarshalling is flexible—you can decode into a struct or just a raw bson.M map if the data shape is unpredictable. It integrates heavily with context for cancellation, which is vital for long-running aggregations."

---

## 168. How do you store JSONB in PostgreSQL using Go?

**Answer:**
Postgres generic JSONB is powerful, and Go supports it well.

If you are using `database/sql`, you typically implement the `Scanner` and `Valuer` interfaces on your struct. This tells the driver: "When you see this Go struct, marshal it to a JSON string for the DB. When reading, unmarshal the JSON string back into the struct."

Libraries like `sqlx` or `GORM` often handle this automatically via tags like `gorm:"type:jsonb"`. It allows us to have a hybrid model—strict schema for core fields (ID, Email) and flexible JSONB for dynamic attributes (User Preferences).

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you store JSONB in PostgreSQL using Go?

**Your Response:** "Postgres generic JSONB is powerful, and Go supports it well. If you are using database/sql, you typically implement Scanner and Valuer interfaces on your struct. This tells the driver: 'When you see this Go struct, marshal it to a JSON string for the DB.'

When reading, unmarshalling is automatic. The driver sees the JSONB tags and knows exactly how to map them to PostgreSQL types. Libraries like GORM often handle this automatically via struct tags like `gorm:"type:jsonb"`."

---

## 169. How do you index and search in Elasticsearch using Go?

**Answer:**
We use the official `go-elasticsearch` client.

It’s effectively a low-level wrapper around the REST API. You construct JSON queries (the "Query DSL") as strings or using a helper library, send them to the cluster, and get back a massive JSON response.

Since Elasticsearch responses are deeply nested JSON, parsing them into Go structs can be tedious. We often define a "Search Result" struct that strictly matches only the fields we care about (like `Hits.Hits.Source`) to avoid mapping the entire verbose response.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you index and search in Elasticsearch using Go?

**Your Response:** "We use the official go-elasticsearch client. It's effectively a low-level wrapper around the REST API. You construct JSON queries (the 'Query DSL') as strings or using helper libraries, send them to the cluster, and get back a massive JSON response.

Since Elasticsearch responses are deeply nested JSON, parsing them into Go structs can be tedious. We often define a 'Search Result' struct that strictly matches only the fields we care about (like `hits.hits.source`) to avoid mapping the entire verbose response."

---

## 170. How do you use Redis with Go for caching?

**Answer:**
The standard library is `go-redis/redis`.

It provides a typesafe client. You mostly use `client.Set(ctx, "key", "value", expiration)` and `client.Get(ctx, "key")`.

For caching structs, Redis only stores strings or bytes. So the pattern is: Marshal your object to JSON -> Write to Redis. When reading: Read string -> Unmarshal JSON. We often wrap this in a generic "Cache Manager" interface to allow swapping Redis for Memcached or in-memory caching later if needed.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Redis with Go for caching?

**Your Response:** "The standard library is go-redis/redis. It provides a typesafe client. You mostly use client.Set(ctx, 'key', 'value', expiration) and client.Get(ctx, 'key').

For caching structs, Redis only stores strings or bytes. So the pattern is: Marshal your object to JSON -> Write to Redis. When reading: Read string -> Unmarshal JSON. We often wrap this in a generic 'Cache Manager' interface to allow swapping Redis for Memcached or in-memory caching later if needed."

---

## 171. How do you use prepared statements in Go?

**Answer:**
Prepared statements protect against SQL injection and can improve performance for repeated queries.

You use `stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")`. This sends the SQL template to the database *once* to be parsed and optimized.

Then you execute it multiple times: `stmt.QueryRow(123)`. In MySQL, this avoids re-parsing the SQL. However, in Go, standard `db.Query("SELECT ... ?", val)` actually prepares, executes, and closes the statement under the hood automatically. So strictly speaking, you only manually `Prepare` if you plan to execute the *exact same* query hundreds of times in a tight loop.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use prepared statements in Go?

**Your Response:** "Prepared statements protect against SQL injection and can improve performance for repeated queries. You use stmt, err := db.Prepare('SELECT * FROM users WHERE id = ?'). This sends the SQL template to the database *once* to be parsed and optimized.

Then you execute it multiple times: stmt.QueryRow(123). In MySQL, this avoids re-parsing the SQL. However, in Go, standard db.Query('SELECT ... ?', val) actually prepares, executes, and closes the statement under the hood automatically."

---

## 172. How do you prevent N+1 queries using Go ORM?

**Answer:**
N+1 problems happen when you fetch a list of items (1 query) and then loop over them to fetch their children (N queries).

In Go ORMs like GORM, you solve this with **Preloading** (Eager Loading). You write `db.Preload("Orders").Find(&users)`.

This tells the ORM to run two optimized queries: one for users, and one massive `SELECT * FROM orders WHERE user_id IN (...)`. It then stitches the results together in memory. It’s strictly better than the naive loop approach, reducing database round-trips from N+1 to just 2.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent N+1 queries using Go ORM?

**Your Response:** "N+1 problems happen when you fetch a list of items (1 query) and then loop over them to fetch their children (N queries). In Go ORMs like GORM, you solve this with Preloading (Eager Loading). You write db.Preload('Orders').Find(&users).

This tells the ORM to run two optimized queries: one for users, and one massive SELECT * FROM orders WHERE user_id IN (...). It then stitches the results together in memory."

---

## 173. How do you map complex nested objects from DB in Go?

**Answer:**
If you aren't using an ORM, this is tricky because SQL returns flat rows.

If you join `Users` and `Address`, you get columns like `user_id, user_name, addr_city, addr_zip`.

To map this to specific nested Go structs, we usually use `sqlx`. It supports dot-notation in tags: `db:"addr.city"`. This tells the scanner to put that column into the `City` field of the nested `Address` struct inside `User`. Without that, you have to verify columns manually and verify nulls, which is error-prone.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you map complex nested objects from DB in Go?

**Your Response:** "If you aren't using an ORM, this is tricky because SQL returns flat rows. If you join Users and Address, you get columns like user_id, user_name, addr_city, addr_zip.

To map this to specific nested Go structs, we usually use sqlx. It supports dot-notation in tags: db:'addr.city'. This tells the scanner to put that column into the City field of the nested Address struct inside User."

---

## 174. How do you benchmark DB performance in Go?

**Answer:**
You don't just benchmark the Go code; you have to treat the DB as an I/O system.

We write Go benchmarks (`BenchmarkQuery`) that run specific queries. Crucially, strictly separate "Driver Overhead" from "DB Execution Time."

To test the driver/allocations, we might mock the DB or run `SELECT 1`. To test the actual query performance, we run against a real Dockerized database. We look at **allocations per operation** in Go—if scanning a row takes 100 allocations, we are probably using too much `reflection` and should optimize the scanning logic.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you benchmark DB performance in Go?

**Your Response:** "We write Go benchmarks (`BenchmarkQuery`) that run specific queries. Crucially, we strictly separate 'Driver Overhead' from 'DB Execution Time.' To test the actual query performance, we run against a real Dockerized database. We look at 'allocations per operation' in Go—if scanning a row takes 100 allocations, we know the driver is doing too much work."

---

## 175. How do you test DB queries with mocks?

**Answer:**
We use `go-sqlmock`.

It allows you to create a fake `sql.DB` connection. You tell the mock: "Expect a generic SELECT query matching this regex, and if you see it, return these 3 fake rows."

This allows you to test your **data access logic** (checking if you handle NULLs correctly, or if you return the right error when no rows are found) without needing a running Postgres instance. It makes unit tests extremely fast (milliseconds).

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test DB queries with mocks?

**Your Response:** "We use go-sqlmock. It allows you to create a fake sql.DB connection. You tell the mock: 'Expect a generic SELECT query matching this regex, and if you see it, return these 3 fake rows.'

This allows you to test your data access logic (checking if you handle NULLs correctly, or if you return the right error when no rows are found) without needing a running Postgres instance. It makes unit tests extremely fast (milliseconds) and completely deterministic."

---

## 176. How do you stream large query results in Go?

**Answer:**
You **never** use `sqlx.Select` or `gorm.Find` for millions of rows, because they load everything into a slice in RAM.

Instead, you use the standard iterative pattern: `rows, _ := db.Query(...)` followed by `for rows.Next()`.

Inside the loop, you `Scan` one row, process it (e.g., write it to a CSV file or send it to a channel), and then discard it. This keeps your memory usage constant (flat), effectively equal to the size of a single row, regardless of whether you process 1,000 or 10 billion records.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream large query results in Go?

**Your Response:** "You never use sqlx.Select or gorm.Find for millions of rows. Instead, you use the standard iterative pattern: rows, _ := db.Query(...) followed by `for rows.Next()`.

Inside the loop, you `Scan` one row, process it (e.g., write it to a CSV file or send it to a channel), and then discard it. This keeps your memory usage constant (flat), effectively equal to the size of a single row, regardless of whether you process 1,000 or 10 billion records."

---

## 177. How do you use SQLite for embedded apps in Go?

**Answer:**
Since Go 1.20ish, we often use generic CGO-free drivers like `modernc.org/sqlite` or the classic `mattn/go-sqlite3` (which requires CGO).

SQLite is literally just a file. You `sql.Open("sqlite3", "./data.db")`.

It is perfect for single-binary tools or edge devices. The main "gotcha" in Go is concurrency. SQLite (by default) only likes one writer at a time. So connection pooling usually needs to be limited: `db.SetMaxOpenConns(1)` is essentially mandatory if you want to avoid "database is locked" errors during concurrent writes.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use SQLite for embedded apps in Go?

**Your Response:** "SQLite is literally just a file. You sql.Open('sqlite3', './data.db').

It is perfect for single-binary tools or edge devices. The main 'gotcha' in Go is concurrency. SQLite (by default) only likes one writer at a time. So connection pooling usually needs to be limited: db.SetMaxOpenConns(1) is essentially mandatory."

---

## 178. How do you connect Go to Amazon RDS or Aurora?

**Answer:**
Connection-wise, it’s just a standard Postgres or MySQL drive. Go doesn't know it's AWS.

However, auth is different. Instead of a hardcoded password, we often use **IAM Authentication**.

We use the AWS SDK to generate an authentication token (which looks like a massive signed URL) and use that as the password in the connection string. This token expires every 15 minutes, so the Go application needs a mechanism to regenerate the token when reconnecting, ensuring we never check long-lived database passwords into source control.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you connect Go to Amazon RDS or Aurora?

**Your Response:** "Connection-wise, it's just a standard Postgres or MySQL drive. Go doesn't know it's AWS.

However, auth is different. Instead of a hardcoded password, we often use IAM Authentication. We use the AWS SDK to generate an authentication token (which looks like a massive signed URL). This token expires every 15 minutes, so the Go application needs a mechanism to regenerate tokens when reconnecting."

---

## 179. How do you manage read replicas in Go?

**Answer:**
Standard database/sql handles usually point to a single writer. To use read replicas, we typically maintain two DB handles: PrimaryDB and ReplicaDB.

In our code, we act intentionally. Commands (`Create`, `Update`) go to PrimaryDB. Queries (`Get`, `List`) go to ReplicaDB. Some ORMs have plugins to do this automatically (read/write splitting), but doing it manually is often safer to avoid 'Replica Lag' bugs where a user saves a post and immediately refreshes fails to see it.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage read replicas in Go?

**Your Response:** "Standard database/sql handles usually point to a single writer. To use read replicas, we typically maintain two DB handles: PrimaryDB and ReplicaDB.

In our code, we act intentionally. Commands (`Create`, `Update`) go to PrimaryDB. Queries (`Get`, `List`) go to ReplicaDB. Some ORMs have plugins to do this automatically (read/write splitting), but doing it manually is often safer to avoid 'Replica Lag' bugs where a user saves a post and immediately refreshes fails to see it."

---

## 180. How do you handle DB failovers in Go apps?

**Answer:**
Database failover (like an AWS RDS writer switching) appears to the app as a broken TCP connection.

Go's built-in connection pool is resilient. If a connection dies, `db.Ping()` or the next query will fail. The driver will discard that bad connection.

The key strategy is **Retries**. We wrap our meaningful database operations in a retry loop using libraries like `avast/retry-go`. If the DB is failing over, the first few queries fail, we wait exponentially (backoff), and by the 3rd retry, the new writer is promoted and the app reconnects automatically.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle DB failovers in Go apps?

**Your Response:** "Database failover (like an AWS RDS writer switching) appears to the app as a broken TCP connection. Go's built-in connection pool is resilient. If a connection dies, db.Ping() or db.Exec() will fail, and the driver will discard that bad connection.

The key strategy is Retries. We wrap our meaningful database operations in a retry loop using libraries like avast/retry-go. If the DB is failing, we wait exponentially (backoff), and by the 3rd retry, we promote the new writer to primary. This makes our app self-healing without manual intervention."
