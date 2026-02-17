# 💾 Go Theory Questions: 921–940 Databases (SQL, NoSQL, ORMs)

## 921. How do you use MongoDB with Go?

**Answer:**
Driver: `go.mongodb.org/mongo-driver`.
BSON (Binary JSON).
`coll := client.Database("test").Collection("users")`.
Insert: `coll.InsertOne(ctx, bson.M{"name": "John"})`.
Find: `coll.Find(ctx, bson.M{"age": bson.M{"$gt": 18}})`.
We use struct tags `bson:"name"` to map Go structs to Mongo documents.

---

## 922. How do you store JSONB in PostgreSQL using Go?

**Answer:**
In Go struct: `Metadata json.RawMessage` (or a custom struct).
In SQL: `data` column is `JSONB`.
`sqlx` or `pgx` handles it automatically if the struct implements `Valuer`/`Scanner` interfaces.
We can query inside the JSON: `WHERE data->>'role' = 'admin'`.

---

## 923. How do you index and search in Elasticsearch using Go?

**Answer:**
Client: `elastic/go-elasticsearch`.
Index: `client.Index("tweets", body)`.
Search: Send JSON query string.
`{"query": {"match": {"content": "golang"}}}`.
Parse response: `hits.hits[]._source`.
It requires manual JSON marshaling unless using a typed wrapper like `olivere/elastic`.

---

## 924. How do you use Redis with Go for caching?

**Answer:**
(See Q 662).
`rdb.Set(ctx, "key", "val", 10*time.Second)`.
Patterns:
- **Cache-Aside**: Logic in Go.
- **Pub/Sub**: `rdb.Subscribe`.
- **Pipelines**: `pipe := rdb.Pipeline()`.
We use Redis interface to mock it in unit tests (`redismock`).

---

## 925. How do you use prepared statements in Go?

**Answer:**
`stmt, _ := db.Prepare("INSERT INTO users VALUES ($1)")`.
`stmt.Exec("Alice")`.
`stmt.Exec("Bob")`.
Benefits:
1.  **Security**: Prevents SQL Injection.
2.  **Performance**: Database parses execution plan once, reuses it.
Note: `database/sql` uses prepared statements automatically under the hood for `db.Query(sql, args)`, but `db.Prepare` holds the statement open explicitly.

---

## 926. How do you prevent N+1 queries using Go ORM?

**Answer:**
**Preloading** (Eager Loading).
GORM: `db.Preload("Orders").Find(&users)`.
It executes 2 queries:
1.  `SELECT * FROM users`.
2.  `SELECT * FROM orders WHERE user_id IN (1, 2, 3...)`.
Then it stitches them in memory.
Without Preload, accessing `user.Orders` loop triggers 1 query per user (N+1).

---

## 927. How do you map complex nested objects from DB in Go?

**Answer:**
SQL returns flat rows.
We need to scan into nested structs.
Library: `scany` or manual logic.
We iterate rows.
If `UserID` changes, create new `User`.
Append `Order` to `User.Orders`.
This "Row Mapping" logic is verbose but necessary if not using an ORM.

---

## 928. How do you benchmark DB performance in Go?

**Answer:**
We verify the **Query Latency**.
Middleware allows logging SQL duration.
`start := time.Now(); db.Query(...); dur := time.Since(start)`.
If `dur > 100ms`, log warning.
We also use `EXPLAIN ANALYZE` (Postgres) by running `db.QueryRow("EXPLAIN ANALYZE " + sql)` and printing the plan.

---

## 929. How do you test DB queries with mocks?

**Answer:**
**go-sqlmock**.
`db, mock, _ := sqlmock.New()`.
`mock.ExpectQuery("SELECT name").WillReturnRows(...)`.
We call the function using `db`.
It asserts that the function executed the exact expected SQL.
This tests the *mapping logic*, not the *database result*. For real verification, use Dockerized integration tests.

---

## 930. How do you stream large query results in Go?

**Answer:**
Assume 1 Million rows.
Do NOT use `db.Select` (loads all into Slice).
Use `rows.Next()` loop.
```go
rows, _ := db.Query(...)
defer rows.Close()
for rows.Next() {
    var u User
    rows.Scan(&u)
    json.NewEncoder(w).Encode(u) // Write directly to network
}
```
Memory usage remains low (1 row size).

---

## 931. How do you use SQLite for embedded apps in Go?

**Answer:**
Driver: `mattn/go-sqlite3` (CGO) or `modernc.org/sqlite` (Pure Go).
Usage is identical to `database/sql`.
Differences: No network. File locking.
We typically enable WAL mode (`PRAGMA journal_mode=WAL`) to allow concurrent readers while writing.

---

## 932. How do you connect Go to Amazon RDS or Aurora?

**Answer:**
Standard Postgres/MySQL driver.
For IAM Auth (no password):
We use AWS SDK to generate a token.
`token := rdsutils.BuildAuthToken(endpoint, region, user, creds)`.
Use `token` as the password in the DSN.
This expires every 15 minutes, so the driver must regenerate it on reconnection.

---

## 933. How do you manage read replicas in Go?

**Answer:**
We maintain **Two Connection Pools**.
1.  `writesDB`: Master.
2.  `readsDB`: Replica.
Code decision:
`func GetUser() { readsDB.Query(...) }`.
`func CreateUser() { writesDB.Exec(...) }`.
Libraries like `gorm` support "Resolver" plugins to automate this based on the operation type (SELECT vs INSERT).

---

## 934. How do you handle DB failovers in Go apps?

**Answer:**
The `database/sql` pool handles reconnections automatically.
If query fails with "Connection Refused":
The driver removes the bad connection.
Next query triggers a new Dial.
If using a DNS (AWS RDS Endpoint) that switches IP during failover, Go's default DNS cache might stale. ensure the OS or Go DNS solver respects low TTL.

---

## 935. How do you use generic repositories in Go 1.18+?

**Answer:**
```go
type Repository[T any] struct { db *sql.DB }
func (r *Repository[T]) Find(id int) (T, error) {
    var entity T
    // scan logic
    return entity, nil
}
```
This reduces boilerplate (UserRepo, ProductRepo, OrderRepo) into a single generic implementation, though handling specific table names requires reflection or passing metadata.

---

## 936. How do you use optimistic locking in Go?

**Answer:**
Add `version` column.
Read: `v := user.Version`.
Update: `UPDATE users SET name=$1, version=$2+1 WHERE id=$3 AND version=$2`.
Result: `rowsAffected, _ := res.RowsAffected()`.
If lines == 0, it means someone else updated it. Return `ErrConflict`.
The Go logic simply checks the `RowsAffected` return value.

---

## 937. How do you handle TimeZones in Go Databases?

**Answer:**
Rule: **Always store UTC**.
Go: `time.Now().UTC()`.
DB: `TIMESTAMP WITH TIME ZONE` (Postgres).
When scanning, the driver returns a `time.Time` in UTC.
Convert to Local only when displaying to the user (Frontend concern).

---

## 938. How do you bulk insert data in Go?

**Answer:**
Values syntax: `INSERT INTO x VALUES (..), (..), (..)`.
We build the query string dynamically (or use a helper).
Max params restriction (Postgres ~65k params).
We chunk the input slice (e.g., 1000 rows per batch).
Copy Protocol (`pgx.CopyFrom`) is even faster for Postgres, bypassing SQL parsing overhead.

---

## 939. How do you use distributed transactions (Saga) in Go?

**Answer:**
No ACID across microservices.
Logic:
1.  Example: Order Service creates Order (Pending).
2.  Publishes "OrderCreated".
3.  Billing Service consumes. Charges Card.
4.  If Success: Publishes "Billed". Order Service updates to "Confirmed".
5.  If Fail: Publishes "BillFailed". Order Service updates to "Failed" (Compensating Transaction).

---

## 940. How do you handle sensitive data (PII) in DBs?

**Answer:**
**Column Level Encryption**.
We encrypt the field *before* sending to SQL.
`EncryptedName := AES_Encrypt("John", Key)`.
DB stores blobs.
If DB is stolen, data is useless.
Key management (KMS) is handled by the Go application side.
