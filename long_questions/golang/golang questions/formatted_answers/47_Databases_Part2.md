# 💾 **921–940: Go + Databases (SQL, NoSQL, ORMs)**

### 921. How do you use MongoDB with Go?
"I use the official `mongo-go-driver`.
`client, _ := mongo.Connect(ctx, options.Client().ApplyURI(...))`.
`coll := client.Database("db").Collection("users")`.
To query: `cursor, _ := coll.Find(ctx, bson.M{"role": "admin"})`.
I decode into structs using `cursor.All(ctx, &results)`.
I always set timeouts on the Context to avoid hanging queries."

#### Indepth
**BSON Struct Tags**. Using `bson:"field_name,omitempty"` is standard. BUT, be careful with `zerocopy` updates. If you want to unset a field in Mongo using struct, `omitempty` will just ignore it. You need to explicitly key off `msg["field"] = nil` using `bson.M` map updates for partial patches, or use pointer fields `*int` to distinguish between 0 and nil.

---

### 922. How do you store JSONB in PostgreSQL using Go?
"I implement the `sql.Scanner` and `driver.Valuer` interfaces on my struct.
`func (a *Attrs) Scan(value any) error { return json.Unmarshal(value.([]byte), a) }`.
Now I can pass the struct directly to `db.Exec("INSERT ... VALUES ($1)", myStruct)` and Postgres handles the JSON conversion."

#### Indepth
**GIN vs GIST**. When indexing JSONB in Postgres:
*   **GIN (Generalized Inverted Index)**: Best for "Contains" queries (`@>`). Faster reads, slower writes (heavy indexing).
*   **GIST**: Faster writes, slower reads. Good for geometric data.
In Go apps, 99% of the time you want GIN for searching JSON documents.

---

### 923. How do you index and search in Elasticsearch using Go?
"I use the official `go-elasticsearch` client.
Indexing: `client.Index("tweets", bodyReader)`.
Searching: `client.Search(client.Search.WithBody(queryReader))`.
Since the request body is JSON, I often use a builder pattern or struct to construct the complex Query DSL `{"query": {"match": ...}}`."

#### Indepth
**Bulk Indexing**. Sending 1 document at a time kills performance. Use `esutil.BulkIndexer`. It buffers documents in memory (e.g., up to 5MB or 1s delay) and sends them in a single HTTP request `_bulk` endpoint. This increases throughput from 100 docs/s to 10,000 docs/s.

---

### 924. How do you use Redis with Go for caching?
"I use `go-redis`.
Use cases beyond caching:
*   **Pub/Sub**: `rdb.Subscribe`.
*   **Rate Limiting**: `INCR` + `EXPIRE`.
*   **Queues**: `LPUSH` / `BRPOP`.
I treat Redis as a primary data structure server, not just a cache."

#### Indepth
**Pipelines**. Redis latency is mostly Round Trip Time (RTT). If you need to `SET` 100 keys, don't do it in a loop (100 network calls). Use `pipe := rdb.Pipeline(); pipe.Set(); ... ; pipe.Exec()`. This sends all 100 commands in 1 TCP packet. Latency drops from 100ms to 1ms.

---

### 925. How do you use prepared statements in Go?
"`stmt, _ := db.Prepare("SELECT * FROM users WHERE id = ?")`.
`defer stmt.Close()`.
`stmt.QueryRow(123)`.
It compiles the SQL on the server once.
Beneficial for repeated queries (bulk insert loop).
But `db.Query` often prepares automatically, so explicit preparation is only needed for high-perf loops."

#### Indepth
**SQL Injection**. Prepared statements are the #1 defense against SQL Injection. `SELECT * FROM users WHERE name = ' + user_input + '` is deadly. `?` acts as a placeholder that strictly treats input as *data*, never *code*. Go's `database/sql` forces this pattern naturally, making Go apps secure by default.

---

### 926. How do you prevent N+1 queries using Go ORM?
"In GORM: `Preload`.
`db.Preload("Orders").Find(&users)`.
It runs 2 queries: `SELECT * FROM users` then `SELECT * FROM orders WHERE user_id IN (...)`.
Without Preload, accessing `user.Orders` triggers a SQL query for *each* user (N+1 problem)."

#### Indepth
**Joins vs Preload**. `Preload` does 2 separate queries. `Joins` does 1 query with `LEFT JOIN`.
*   Use **Joins** when you need to *filter* by the child (Find users who bought "Apple").
*   Use **Preload** when you just need to *load* the child.
Joins transfer duplicated parent data (bandwidth heavy), Preload is cleaner but not atomic.

---

### 927. How do you map complex nested objects from DB in Go?
"If not using an ORM, I use `sqlx`.
`type User struct { Address Address `db:"address"` }`.
I join the tables: `SELECT u.*, a.city AS "address.city" ...`.
`sqlx` maps the dotted column names to the nested struct fields automatically."

#### Indepth
**Flat Structures**. Alternatively, some prefer defining a "Read Model" struct that is completely flat `type UserRow struct { UserID int; City string }` and then mapping it to domain objects manually. This avoids the reflection overhead of `sqlx` and gives you compile-time safety on the mapping logic.

---

### 928. How do you benchmark DB performance in Go?
"I don't just benchmark the Go code; I benchmark the *interaction*.
I write a test that runs `db.Exec` in a loop.
I observe latency histograms.
I use tools like `pgbench` for raw DB speed, and Go benchmarks to see if my driver/serialization is the bottleneck."

#### Indepth
**Driver Overhead**. Not all drivers are equal. `pgx` (Postgres) is significantly faster than `lib/pq` (which is effectively unmaintained) because `pgx` uses binary wire protocol instead of text protocol. Switching from `pq` to `pgx` can yield 20% perf gain for free.

---

### 929. How do you test DB queries with mocks?
"I use `go-sqlmock`.
`db, mock, _ := sqlmock.New()`.
`mock.ExpectQuery("SELECT").WillReturnRows(...)`.
I inject this `db` into my repo.
However, it doesn't prove the SQL works on the real DB (syntax errors). Integration tests are better."

#### Indepth
**Dockertest**. `go-sqlmock` is for *Unit Tests*. For *Integration Tests*, use `ory/dockertest`. It spins up a real ephemeral Postgres docker container from within your `TestMain`. You run real migrations and real queries. It proves your SQL syntax is valid for that specific Postgres version (e.g., 14.2).

---

### 930. How do you stream large query results in Go?
"`rows, _ := db.Query("SELECT * FROM wide_table")`.
`for rows.Next() { scan(); process() }`.
This streams row-by-row.
I **never** use `sqlx.Select` or `cursor.All` for huge datasets because they load everything into a slice (RAM OOM)."

#### Indepth
**Cursor**. In Postgres, a simple `SELECT` might still try to buffer results on the server-side. Use a **Cursor Transaction**: `BEGIN; DECLARE mycursor CURSOR FOR SELECT...; FETCH 1000 FROM mycursor;`. This ensures the DB server also streams data lazily instead of preparing 10GB of result set in RAM.

---

### 931. How do you use SQLite for embedded apps in Go?
"I use `mattn/go-sqlite3` (CGO) or `modernc.org/sqlite` (Pure Go).
`db, _ := sql.Open("sqlite3", "file:data.db")`.
I verify to turn on WAL mode: `PRAGMA journal_mode=WAL;` to allow concurrent readers and one writer."

#### Indepth
**Litestream**. SQLite is just a file. But how do you back it up? `Litestream` (written in Go) replicates the WAL (Write Ahead Log) to S3 in real-time. This gives you "Serverless Database" capability—if your server crashes, you restore from S3 with <1s data loss. Ideal for single-node Go apps.

---

### 932. How do you connect Go to Amazon RDS or Aurora?
"Standard Postgres/MySQL driver.
But for **IAM Auth** (passwordless), I use the AWS SDK to generate a token.
`token := rdsutils.BuildAuthToken(...)`.
I use the token as the password in `sql.Open`.
I must refresh this token every 15 minutes."

#### Indepth
**RDS Proxy**. IAM Auth Tokens are computationally expensive for RDS to verify (RSA decryption). If you open 100 connections/sec, you will spike RDS CPU. Use **RDS Proxy** to reuse connections. The Proxy handles the IAM Auth, and your Go app connects to the Proxy.

---

### 933. How do you manage read replicas in Go?
"I create two DB handles.
`Primary *sql.DB`.
`Replica *sql.DB`.
In my code: `if isWrite { Primary.Exec(...) } else { Replica.Query(...) }`.
Or I use a resolver middleware if using an ORM."

#### Indepth
**Replication Lag**. Reading from Replica is "Eventually Consistent". User updates profile -> Redirect to Profile Page -> Read from Replica -> Old Profile shown (Panic!). **Sticky Sessions** or **Read-After-Write** consistency is needed. A simple fix: "If user just wrote, read from Master for 5 seconds."

---

### 934. How do you handle DB failovers in Go apps?
"I rely on the driver's reconnection logic + connection pooling.
If the connection breaks, `db.Ping()` fails.
The driver attempts to reconnect.
For logic: I use retries (exponential backoff) on `db.Query`.
The DNS switch happens automatically, but my app must close broken connections to resolve the new IP."

#### Indepth
**Cloud SQL Connector**. For Google Cloud SQL or Azure PostgreSQL, Standard DNS failover can take minutes. The "Connector" libraries (Go SDKs) are smarter. They talk to the Cloud API to find the current IP of the master. They handle certificate rotation and automatic failover much faster than DNS TTL allows.

---

### 935. How do you use Migrations in Go?
"I use **Golang Migrate** or **Goose**.
I define up/down SQL files: `001_create_users.up.sql`.
I run `migrate up` locally or in CD pipeline.
It stores the current version in a `schema_migrations` table to prevent double execution."

#### Indepth
**Embed Migrations**. Don't rely on `file://` usage in production (files might be missing in the Docker image). Use `//go:embed migrations/*.sql` to compile SQL files into the binary. `golang-migrate` supports `io/fs`. This makes your binary truly self-contained—it can migrate its own DB on startup.

---

### 936. How do you handle transaction locking in Go?
"I use `tx.Exec("SELECT ... FOR UPDATE")`.
This locks the rows until the transaction commits.
It prevents race conditions where two goroutines read the same balance and update it simultaneously."

#### Indepth
**Optimistic Locking**. `FOR UPDATE` is Pessimistic/heavy. Alternative: Add `version int` column. `UPDATE accounts SET balance=?, version=version+1 WHERE id=? AND version=current_version`. If RowsAffected is 0, someone else updated it. Retry or fail. This scales better for low-contention systems.

---

### 937. How do you implement soft deletes in Go?
"Add `DeletedAt *time.Time` to the struct.
`UPDATE users SET deleted_at = NOW() WHERE id = ?`.
Queries must filter: `WHERE deleted_at IS NULL`.
GORM handles this automatically, but manually I need to be disciplined."

#### Indepth
**Unique Indexes**. Soft Deletes break unique constraints. `UNIQUE(email)`. User A deletes `bob@gmail.com`. User B tries to register `bob@gmail.com` -> DB Error "Duplicate", even though A is deleted. Fix: `UNIQUE INDEX ... WHERE deleted_at IS NULL` (Partial Index in Postgres).

---

### 938. How do you use Listen/Notify with Postgres in Go?
"I use `pq` or `pgx` driver.
`listener := pq.NewListener(...)`.
`listener.Listen("events")`.
`case n := <-listener.Notify: handle(n.Extra)`.
This allows real-time updates without polling the DB."

#### Indepth
**CDC**. Listen/Notify has a payload limit (8000 bytes) and isn't durable (if app is down, event is lost). For robust "Data Change" pipelines, use **CDC (Change Data Capture)** like Debezium. It reads the Postgres WAL and pushes changes to Kafka. Go consumers read Kafka. This ensures 100% data fidelity.

---

### 939. How do you handle connection pooling settings?
"Crucial for stability.
`db.SetMaxOpenConns(25)`.
`db.SetMaxIdleConns(25)`.
`db.SetConnMaxLifetime(5 * time.Minute)`.
If MaxOpen is too high, I starve the DB. If too low, my app blocks waiting for connections."

#### Indepth
**Timeouts**. `ConnMaxLifetime` is critical for Load Balancers (AWS ALB). If Go keeps a connection open for 1 hour, but ALB kills it silently after 5 minutes, Go will try to use a dead connection and get "Unexpected EOF". Set Go's lifetime to be *shorter* than the infrastructure's timeout.

---

### 940. How do you use generic repositories in Go?
"With Go 1.18 Generics.
`type Repository[T any] struct { db *sql.DB }`.
`func (r *Repository[T]) Find(id int) (T, error)`.
It reduces boilerplate, but I verify not to over-abstract. Sometimes specific queries need specific SQL optimization."

#### Indepth
**Interface Segregation**. Don't make one giant `Repository` interface. Split it. `Reader`, `Writer`. Or even better, `UserFinder`, `UserSaver`. This allows you to decorate just the `Finder` with a Cache layer without implementing the `Saver` methods. Composition over Inheritance.
