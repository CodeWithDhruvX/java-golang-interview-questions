## 💾 Go + Databases (Questions 921-940)

### Question 921: How do you use `database/sql` in Go?

**Answer:**
Standard library.
1.  `db, _ := sql.Open("postgres", connStr)`.
2.  `rows, _ := db.Query("SELECT ...")`.
3.  `defer rows.Close()`.
4.  `rows.Scan(&var)`.

### Explanation
database/sql in Go is the standard library for database operations. It involves opening connections with sql.Open, executing queries with db.Query, closing rows with defer, and scanning results into variables with rows.Scan.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `database/sql` in Go?
**Your Response:** "I use the `database/sql` standard library for database operations in Go. First, I open a connection using `sql.Open` with the driver name and connection string. Then I execute queries using `db.Query` for SELECT statements or `db.Exec` for INSERT/UPDATE/DELETE. I always defer `rows.Close()` to release resources, and I use `rows.Scan` to map the result columns to Go variables. For single rows, I use `QueryRow` instead of `Query`. I handle errors properly and use prepared statements for performance. The key is understanding that `sql.Open` doesn't actually connect - it just validates the parameters. The real connection happens when I execute the first query. I also manage connection pools and handle transactions properly."

---

### Question 922: What are connection pools and how to manage them?

**Answer:**
`sql.DB` *is* a pool.
Configure it:
- `SetMaxOpenConns(N)`: Hard limit.
- `SetMaxIdleConns(M)`: Keep warm connections.
- `SetConnMaxLifetime(T)`: Close old connections (useful for cloud LBs).

### Explanation
Connection pools in Go are managed by sql.DB which is itself a pool. Configuration includes SetMaxOpenConns for hard limits, SetMaxIdleConns for warm connections, and SetConnMaxLifetime to close old connections for cloud load balancers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are connection pools and how to manage them?
**Your Response:** "In Go, `sql.DB` is actually a connection pool, not a single connection. I configure it using three main methods. `SetMaxOpenConns(N)` sets the hard limit on total connections. `SetMaxIdleConns(M)` keeps some connections warm and ready to use. `SetConnMaxLifetime(T)` closes connections after they've been open too long, which is important for cloud load balancers that might terminate long-lived connections. I tune these based on my application's needs - more connections for high concurrency, fewer for resource efficiency. The pool automatically handles connection reuse and lifecycle management. The key is finding the right balance between performance and resource usage. I monitor connection metrics and adjust as needed."

---

### Question 923: How do you write raw queries using sqlx?

**Answer:**
**sqlx** extends standard sql.
`db.Select(&users, "SELECT * FROM users")`.
Automatically maps struct fields (via `db` tag) to columns.

### Explanation
sqlx extends the standard sql package with convenience methods like db.Select that automatically map struct fields to database columns using db tags, reducing boilerplate code for common database operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write raw queries using sqlx?
**Your Response:** "I use sqlx when I want the performance of raw SQL but with more convenience than standard database/sql. The main benefit is automatic struct mapping - I can use `db.Select(&users, "SELECT * FROM users")` and sqlx automatically maps the columns to struct fields using `db` tags. This eliminates the need for manual `rows.Scan` calls. I also get methods like `db.Get` for single rows and `db.Exec` for updates. Sqlx still gives me full control over SQL while reducing the boilerplate. It's a good middle ground between raw SQL and full ORM. I use it when I need complex queries but want cleaner code. The key is understanding that sqlx is a thin wrapper around standard sql - it doesn't change the fundamental database operations."

---

### Question 924: How do you use GORM with PostgreSQL?

**Answer:**
ORM.
`db.Find(&users)`.
Handles Migrations (`AutoMigrate`), Hooks (BeforeSave), and Associations (Preload).

### Explanation
GORM is an ORM for Go that provides methods like db.Find for database operations. It handles migrations with AutoMigrate, hooks like BeforeSave for lifecycle events, and associations with Preload for related data loading.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use GORM with PostgreSQL?
**Your Response:** "I use GORM as an ORM to work with PostgreSQL in a more object-oriented way. I can call `db.Find(&users)` instead of writing raw SQL. GORM handles migrations through `AutoMigrate` which creates tables based on my struct definitions. I use hooks like `BeforeSave` to add validation or logging before database operations. For related data, I use `Preload` to automatically load associations. GORM builds the SQL queries for me, which speeds up development but can make debugging harder. I use it for standard CRUD operations and fall back to raw SQL for complex queries. The key is understanding the trade-offs - GORM is faster to write but might not be as optimized as hand-tuned SQL."

---

### Question 925: How do you handle transactions in Go?

**Answer:**
```go
tx, _ := db.Begin()
_, err := tx.Exec(...)
if err != nil {
    tx.Rollback()
    return
}
tx.Commit()
```

### Explanation
Transactions in Go use db.Begin to start, tx.Exec for operations, tx.Rollback on errors, and tx.Commit on success. This ensures atomicity - either all operations succeed or none do.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle transactions in Go?
**Your Response:** "I handle transactions in Go using the database/sql transaction pattern. I start with `tx, err := db.Begin()` to create a transaction context. Then I execute all my operations using `tx.Exec` or `tx.Query` instead of the regular `db` methods. If any operation fails, I immediately call `tx.Rollback()` to undo everything. If all operations succeed, I call `tx.Commit()` to make the changes permanent. I always check for errors and ensure rollback happens in case of failure. The key is using the transaction object consistently and ensuring either rollback or commit happens. I also handle transaction timeouts and use context for cancellation. This pattern ensures data consistency even when errors occur."

---

### Question 926: How do you create database migrations in Go?

**Answer:**
Use **golang-migrate** (CLI/Library) or **Goose**.
SQL files: `001_init.up.sql`, `001_init.down.sql`.
Tool tracks `schema_migrations` table version.

### Explanation
Database migrations in Go use tools like golang-migrate or Goose with SQL files for up/down migrations. These tools track versions in a schema_migrations table and apply migrations in order to manage database schema changes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create database migrations in Go?
**Your Response:** "I create database migrations using tools like golang-migrate or Goose. I write SQL files in pairs - `001_init.up.sql` for the forward migration and `001_init.down.sql` for the rollback. The migration tool tracks which migrations have been applied in a `schema_migrations` table. I can run migrations up or down as needed. This approach gives me version control for my database schema. I can run migrations as part of my deployment process or manually during development. The key is having both forward and rollback migrations so I can safely change schema. I also test migrations thoroughly before deploying to production. This pattern ensures database changes are repeatable and trackable across environments."

---

### Question 927: How do you use MongoDB with Go?

**Answer:**
Official driver: `go.mongodb.org/mongo-driver`.
BSON handling: `bson.M{"name": "Alice"}`.
Context is critical for timeouts.

### Explanation
MongoDB with Go uses the official mongo-driver which handles BSON documents using bson.M for maps. Context is critical for timeouts and cancellation in all MongoDB operations to prevent hanging requests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use MongoDB with Go?
**Your Response:** "I use MongoDB with Go using the official `go.mongodb.org/mongo-driver`. I connect to the database and work with collections using the client. For data, I use BSON types like `bson.M` for maps or define structs with BSON tags. Context is critical - I pass context to every operation to handle timeouts and cancellation properly. I can insert, find, update, and delete documents using the collection methods. The driver handles connection pooling and retry logic automatically. I also use aggregation pipelines for complex queries. The key is understanding that MongoDB operations can be slow, so proper context handling is essential. I handle errors and check for document existence properly. The driver provides a clean, idiomatic Go interface to MongoDB."

---

### Question 928: How do you store JSONB in PostgreSQL using Go?

**Answer:**
Implement `sql.Scanner` and `driver.Valuer` interfaces on a struct.
Inside methods, use `json.Unmarshal/Marshal`.
This allows GORM/SQL to save the struct as a JSON string automatically.

### Explanation
JSONB storage in PostgreSQL uses sql.Scanner and driver.Valuer interfaces on structs. These interfaces use json.Unmarshal/Marshal internally to automatically convert structs to JSON strings for database storage and back.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you store JSONB in PostgreSQL using Go?
**Your Response:** "I store JSONB data in PostgreSQL by implementing the `sql.Scanner` and `driver.Valuer` interfaces on my structs. The `Scan` method uses `json.Unmarshal` to convert JSON from the database into my struct, and the `Value` method uses `json.Marshal` to convert my struct to JSON for storage. This makes the conversion automatic - I can work with regular Go structs and the database handles the JSON conversion. GORM and other ORMs recognize these interfaces and use them automatically. The key is implementing these two methods correctly and handling errors. This approach gives me the best of both worlds - structured Go code and flexible JSON storage in the database."

---

### Question 929: How do you index and search in Elasticsearch using Go?

**Answer:**
`olivere/elastic` or official `go-elasticsearch`.
Construct JSON query DSL map.
Send HTTP POST.

### Explanation
Elasticsearch integration in Go uses libraries like olivere/elastic or go-elasticsearch. Queries are constructed as JSON DSL maps and sent via HTTP POST to Elasticsearch for indexing and searching.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you index and search in Elasticsearch using Go?
**Your Response:** "I work with Elasticsearch in Go using either the `olivere/elastic` library or the official `go-elasticsearch` client. I construct queries using the JSON query DSL - building maps or structs that represent the search criteria. Then I send these queries via HTTP POST to Elasticsearch. For indexing, I prepare documents as JSON and send them to the appropriate index. I handle the response parsing and error checking. The libraries provide helpers for common operations like aggregations, paging, and filtering. The key is understanding the Elasticsearch query DSL and structuring the queries correctly. I also handle connection pooling and retry logic for resilience. This approach gives me powerful search capabilities while working in Go."

---

### Question 930: How do you use Redis with Go for caching?

**Answer:**
(See Q662). `client.Set(ctx, key, val, ttl)`.

### Explanation
Redis caching in Go uses client.Set with context, key, value, and TTL parameters. This provides fast in-memory caching with expiration times for temporary data storage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Redis with Go for caching?
**Your Response:** "I use Redis for caching in Go by connecting with a Redis client and using `client.Set(ctx, key, val, ttl)` to store data with expiration. I pass context for timeout control, specify the key and value, and set a TTL so the data expires automatically. For retrieval, I use `client.Get` and handle the case where the key doesn't exist. I also use Redis for distributed locking, pub/sub messaging, and rate limiting. The key is understanding that Redis is an in-memory store, so it's much faster than database queries but the data can be lost. I use it for caching frequently accessed data, session storage, or real-time features. I handle connection pooling and error recovery properly."

---

### Question 931: How do you use prepared statements in Go?

**Answer:**
`stmt, _ := db.Prepare("SELECT ... ?")`.
`stmt.QueryRow(arg)` repeated.
Saves parsing time on DB side.

### Explanation
Prepared statements in Go use db.Prepare to create reusable SQL statements with placeholders. These can be executed multiple times with different arguments, saving parsing time on the database side.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use prepared statements in Go?
**Your Response:** "I use prepared statements in Go when I need to execute the same query multiple times with different parameters. I prepare the statement once using `db.Prepare` with placeholders, then execute it multiple times with `stmt.QueryRow` or `stmt.Exec`. This saves the database from parsing the same SQL repeatedly, which improves performance. Prepared statements also protect against SQL injection since parameters are sent separately from the query. I use them for high-frequency operations like inserts in loops or frequently run selects. I always close the statement when done with `defer stmt.Close()`. The key is understanding when the performance benefit outweighs the extra code complexity. For one-off queries, regular queries might be simpler."

---

### Question 932: How do you prevent N+1 queries using Go ORM?

**Answer:**
**Preloading (Eager Loading).**
GORM: `db.Preload("Orders").Find(&users)`.
Fetches all Users (1 query), collects IDs, fetches all Orders (1 query). Total 2 queries instead of N+1.

### Explanation
N+1 query prevention uses preloading/eager loading. In GORM, db.Preload("Orders").Find(&users) fetches users in one query, collects their IDs, then fetches all related orders in a second query, reducing N+1 queries to just 2.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent N+1 queries using Go ORM?
**Your Response:** "I prevent N+1 queries using preloading or eager loading. In GORM, I use `db.Preload('Orders').Find(&users)` which fetches all users in one query, then collects their IDs and fetches all related orders in a second query. Instead of N+1 queries where I'd query for orders for each user individually, I get just 2 queries total. This dramatically improves performance. The key is identifying when I'll need related data and loading it upfront rather than lazily. I can preload multiple associations by chaining Preload calls. The tradeoff is potentially loading more data than needed, but it's usually worth the performance gain. I also use Select to specify which fields I need when preloading to reduce data transfer."

---

### Question 933: How do you map complex nested objects from DB in Go?

**Answer:**
In raw SQL, use `JOIN`.
Scan into separate vars, then assemble struct manually.
Or use `sqlx`: `results := []struct { User User `db:"u"`; Address Address `db:"a"` }{}`.

### Explanation
Complex nested object mapping from databases uses JOIN queries in raw SQL, scanning into separate variables and assembling structs manually, or using sqlx with anonymous structs that map database columns to nested fields.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you map complex nested objects from DB in Go?
**Your Response:** "I map complex nested objects from databases in two main ways. With raw SQL, I use JOIN queries to fetch related data, scan into separate variables, and then assemble the nested structs manually. This gives me full control but requires more code. With sqlx, I can define anonymous structs with nested field mappings like `User User db:'u'` and `Address Address db:'a'`, and sqlx automatically maps the JOIN results. The key is understanding how to structure the SQL query and the Go structs to match. For very complex hierarchies, I might use multiple queries and assemble the objects in Go. I choose based on the complexity and performance requirements. The goal is to get the data I need in as few queries as possible while keeping the mapping code clean."

---

### Question 934: How do you benchmark DB performance in Go?

**Answer:**
Write a Go benchmark that calls the DB (clean up after).
Measure `Allocations` (client side overhead) and `Time`.
Latency usually dominated by Network.

### Explanation
Database performance benchmarking in Go involves writing benchmarks that call database operations and clean up after. Measurements focus on allocations (client-side overhead) and time, with network latency typically being the dominant factor.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you benchmark DB performance in Go?
**Your Response:** "I benchmark database performance by writing Go benchmarks that call database operations and clean up afterward. I measure both allocations to see client-side overhead and time to see total latency. Since network latency usually dominates, I focus on optimizing round trips and query efficiency. I run benchmarks with different data sizes, query complexities, and concurrency levels. I also benchmark connection pool configurations and prepared statements vs ad-hoc queries. The key is creating realistic test scenarios and measuring what matters - actual query performance, not just Go code overhead. I use the benchmarking results to optimize queries, adjust connection pools, or choose between different approaches."

---

### Question 935: How do you test DB queries with mocks?

**Answer:**
(See Q544). `go-sqlmock`.

### Explanation
Database query testing with mocks uses go-sqlmock library to simulate database responses without requiring actual database connections, enabling unit testing of database logic.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you test DB queries with mocks?
**Your Response:** "I test database queries using the `go-sqlmock` library which lets me mock database responses without needing a real database. I set up expected queries and their results, then run my code against the mock. This lets me test my database logic in unit tests without the overhead and complexity of a real database. I can test error conditions, different result sets, and edge cases. The mock ensures my code handles the database correctly while keeping tests fast and isolated. I use this for testing repository patterns, query building logic, and error handling. The key is setting up realistic mock scenarios that match what the real database would return. This approach makes my tests reliable and fast while still validating the database interaction logic."

---

### Question 936: How do you stream large query results in Go?

**Answer:**
`rows.Next()`.
Process one row, write to output/stream, discard.
DB driver typically buffers 1 row, not whole result set (unless configured otherwise).

### Explanation
Large query result streaming in Go uses rows.Next() to iterate through results one at a time. Each row is processed and discarded, with database drivers typically buffering only one row to minimize memory usage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream large query results in Go?
**Your Response:** "I stream large query results by iterating with `rows.Next()` and processing one row at a time. I process each row, write the output to a stream or file, and then discard the row before moving to the next. The database driver typically buffers only one row at a time, so memory usage stays low even with millions of rows. I use `defer rows.Close()` to ensure cleanup and handle errors properly. This approach lets me process datasets larger than memory. I might use a worker pool to process rows in parallel if needed. The key is not loading everything into memory at once. This pattern works well for exporting data, processing large datasets, or streaming results to clients."

---

### Question 937: How do you use SQLite for embedded apps in Go?

**Answer:**
`mattn/go-sqlite3` (CGO).
Or `modernc.org/sqlite` (Pure Go transpilation).
Zero config. Good for prototyping or single-user desktop apps.

### Explanation
SQLite for embedded Go apps uses mattn/go-sqlite3 with CGO or modernc.org/sqlite with pure Go transpilation. Both provide zero-configuration databases ideal for prototyping and single-user desktop applications.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use SQLite for embedded apps in Go?
**Your Response:** "I use SQLite for embedded applications using either `mattn/go-sqlite3` which uses CGO, or `modernc.org/sqlite` which is pure Go. The pure Go version is easier to deploy since it doesn't require CGO, though the CGO version might be slightly faster. SQLite is great because it requires zero configuration - the database is just a file. I use it for prototyping, desktop applications, or small web apps where I don't need a full database server. It supports standard SQL so I can use the same database/sql interface. The key is understanding that SQLite is for single-user scenarios - it doesn't handle concurrent writes well. For embedded apps or tools, it's perfect because there's no external database to manage."

---

### Question 938: How do you connect Go to Amazon RDS or Aurora?

**Answer:**
It's just Postgres/MySQL protocol.
Use standard drivers.
Ensure Security Group allows access.

### Explanation
Amazon RDS/Aurora connection in Go uses standard Postgres/MySQL drivers since they speak standard protocols. The key is ensuring security groups allow access from the application's EC2 instances or IP ranges.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you connect Go to Amazon RDS or Aurora?
**Your Response:** "Connecting Go to Amazon RDS or Aurora is straightforward because they use standard database protocols. I use the same standard drivers I'd use for local PostgreSQL or MySQL - `lib/pq` for Postgres or `go-sql-driver/mysql` for MySQL. The connection string contains the RDS endpoint, credentials, and database name. The key configuration is in AWS - I need to ensure the security group allows connections from my application's EC2 instances or IP range. I also handle the connection properly for the cloud environment - setting appropriate timeouts and handling network issues. The beauty is that my Go code doesn't need to know it's talking to RDS - it's just a standard database connection."

---

### Question 939: How do you manage read replicas in Go?

**Answer:**
Two DB pools: `MasterDB`, `ReplicaDB`.
Writes -> Master.
Reads -> Replica (Accepting Replication Lag).

### Explanation
Read replica management in Go uses two database pools - MasterDB for writes and ReplicaDB for reads. This accepts replication lag where reads might be slightly behind writes but improves read performance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage read replicas in Go?
**Your Response:** "I manage read replicas by maintaining two database connection pools - one for the master database and one for the replica. All write operations go to the master, while read operations go to the replica. This distributes the read load and improves performance. The key consideration is replication lag - reads from the replica might be slightly behind writes to the master. For operations that need to read data immediately after writing, I might read from the master instead. I implement this pattern by having two `sql.DB` instances and routing queries appropriately. I also handle failover - if the replica is unavailable, I can fall back to reading from the master. This approach scales read-heavy applications effectively."

---

### Question 940: How do you handle DB failovers in Go apps?

**Answer:**
The Driver usually handles reconnection logic if TCP breaks (`Bad Connection`).
Application code should retry the transaction if it sees a transient network error.

### Explanation
Database failover handling in Go relies on drivers that automatically handle reconnection for TCP breaks. Application code should retry transactions when encountering transient network errors to ensure resilience during failover events.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle DB failovers in Go apps?
**Your Response:** "I handle database failovers by relying on the database driver's built-in reconnection logic for TCP breaks. When a connection fails, the driver typically tries to reconnect automatically. In my application code, I retry transactions when I encounter transient network errors - these are temporary issues that resolve themselves. I implement retry logic with exponential backoff for these cases. I also use context with timeouts so operations don't hang forever. The key is distinguishing between transient errors that should be retried and permanent errors that should fail immediately. I monitor connection health and might implement circuit breakers for severe failures. Most importantly, I design my application to handle temporary database unavailability gracefully."

---
