## 💾 Go + Databases (Questions 921-940)

### Question 921: How do you use `database/sql` in Go?

**Answer:**
Standard library.
1.  `db, _ := sql.Open("postgres", connStr)`.
2.  `rows, _ := db.Query("SELECT ...")`.
3.  `defer rows.Close()`.
4.  `rows.Scan(&var)`.

---

### Question 922: What are connection pools and how to manage them?

**Answer:**
`sql.DB` *is* a pool.
Configure it:
- `SetMaxOpenConns(N)`: Hard limit.
- `SetMaxIdleConns(M)`: Keep warm connections.
- `SetConnMaxLifetime(T)`: Close old connections (useful for cloud LBs).

---

### Question 923: How do you write raw queries using sqlx?

**Answer:**
**sqlx** extends standard sql.
`db.Select(&users, "SELECT * FROM users")`.
Automatically maps struct fields (via `db` tag) to columns.

---

### Question 924: How do you use GORM with PostgreSQL?

**Answer:**
ORM.
`db.Find(&users)`.
Handles Migrations (`AutoMigrate`), Hooks (BeforeSave), and Associations (Preload).

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

---

### Question 926: How do you create database migrations in Go?

**Answer:**
Use **golang-migrate** (CLI/Library) or **Goose**.
SQL files: `001_init.up.sql`, `001_init.down.sql`.
Tool tracks `schema_migrations` table version.

---

### Question 927: How do you use MongoDB with Go?

**Answer:**
Official driver: `go.mongodb.org/mongo-driver`.
BSON handling: `bson.M{"name": "Alice"}`.
Context is critical for timeouts.

---

### Question 928: How do you store JSONB in PostgreSQL using Go?

**Answer:**
Implement `sql.Scanner` and `driver.Valuer` interfaces on a struct.
Inside methods, use `json.Unmarshal/Marshal`.
This allows GORM/SQL to save the struct as a JSON string automatically.

---

### Question 929: How do you index and search in Elasticsearch using Go?

**Answer:**
`olivere/elastic` or official `go-elasticsearch`.
Construct JSON query DSL map.
Send HTTP POST.

---

### Question 930: How do you use Redis with Go for caching?

**Answer:**
(See Q662). `client.Set(ctx, key, val, ttl)`.

---

### Question 931: How do you use prepared statements in Go?

**Answer:**
`stmt, _ := db.Prepare("SELECT ... ?")`.
`stmt.QueryRow(arg)` repeated.
Saves parsing time on DB side.

---

### Question 932: How do you prevent N+1 queries using Go ORM?

**Answer:**
**Preloading (Eager Loading).**
GORM: `db.Preload("Orders").Find(&users)`.
Fetches all Users (1 query), collects IDs, fetches all Orders (1 query). Total 2 queries instead of N+1.

---

### Question 933: How do you map complex nested objects from DB in Go?

**Answer:**
In raw SQL, use `JOIN`.
Scan into separate vars, then assemble struct manually.
Or use `sqlx`: `results := []struct { User User `db:"u"`; Address Address `db:"a"` }{}`.

---

### Question 934: How do you benchmark DB performance in Go?

**Answer:**
Write a Go benchmark that calls the DB (clean up after).
Measure `Allocations` (client side overhead) and `Time`.
Latency usually dominated by Network.

---

### Question 935: How do you test DB queries with mocks?

**Answer:**
(See Q544). `go-sqlmock`.

---

### Question 936: How do you stream large query results in Go?

**Answer:**
`rows.Next()`.
Process one row, write to output/stream, discard.
DB driver typically buffers 1 row, not whole result set (unless configured otherwise).

---

### Question 937: How do you use SQLite for embedded apps in Go?

**Answer:**
`mattn/go-sqlite3` (CGO).
Or `modernc.org/sqlite` (Pure Go transpilation).
Zero config. Good for prototyping or single-user desktop apps.

---

### Question 938: How do you connect Go to Amazon RDS or Aurora?

**Answer:**
It's just Postgres/MySQL protocol.
Use standard drivers.
Ensure Security Group allows access.

---

### Question 939: How do you manage read replicas in Go?

**Answer:**
Two DB pools: `MasterDB`, `ReplicaDB`.
Writes -> Master.
Reads -> Replica (Accepting Replication Lag).

---

### Question 940: How do you handle DB failovers in Go apps?

**Answer:**
The Driver usually handles reconnection logic if TCP breaks (`Bad Connection`).
Application code should retry the transaction if it sees a transient network error.

---
