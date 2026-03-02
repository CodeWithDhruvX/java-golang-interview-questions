# 📔 06 — Databases in Go
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `database/sql` package
- GORM ORM
- Connection pooling
- Transactions
- Prepared statements
- Common query patterns (CRUD)

---

## ❓ Most Asked Questions

### Q1. How do you connect to a MySQL/PostgreSQL database?

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"  // side-effect import
)

// MySQL connection
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
if err != nil { log.Fatal(err) }
defer db.Close()

// Test connection
if err := db.Ping(); err != nil { log.Fatal(err) }

// Configure connection pool
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)

// PostgreSQL
import _ "github.com/lib/pq"
db, _ = sql.Open("postgres", "host=localhost user=user password=pass dbname=mydb sslmode=disable")
```

---

### Q2. How do you perform CRUD operations with `database/sql`?

```go
type User struct {
    ID    int
    Name  string
    Email string
}

// CREATE
func createUser(db *sql.DB, u User) (int64, error) {
    result, err := db.Exec(
        "INSERT INTO users (name, email) VALUES (?, ?)",
        u.Name, u.Email,
    )
    if err != nil { return 0, err }
    return result.LastInsertId()
}

// READ one
func getUserByID(db *sql.DB, id int) (*User, error) {
    row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
    var u User
    err := row.Scan(&u.ID, &u.Name, &u.Email)
    if err == sql.ErrNoRows { return nil, nil }
    if err != nil { return nil, err }
    return &u, nil
}

// READ many
func getAllUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil { return nil, err }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil { return nil, err }
        users = append(users, u)
    }
    return users, rows.Err()
}

// UPDATE
func updateUser(db *sql.DB, u User) error {
    _, err := db.Exec("UPDATE users SET name=?, email=? WHERE id=?", u.Name, u.Email, u.ID)
    return err
}

// DELETE
func deleteUser(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}
```

---

### Q3. How do transactions work in Go?

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    // Begin transaction
    tx, err := db.Begin()
    if err != nil { return err }

    // Defer rollback — if anything fails, rollback
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    // Debit source account
    _, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
    if err != nil { return err }

    // Credit destination account
    _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
    if err != nil { return err }

    return nil  // triggers tx.Commit()
}
```

---

### Q4. What are prepared statements and why use them?

```go
// Prepared statement — compiled once, executed many times
// Prevents SQL injection, better performance for repeated queries
stmt, err := db.Prepare("SELECT id, name FROM users WHERE email = ?")
if err != nil { log.Fatal(err) }
defer stmt.Close()

// Execute multiple times with different params
var id int; var name string
for _, email := range emails {
    err = stmt.QueryRow(email).Scan(&id, &name)
    // process result
}
```
> **Security:** Always use parameterized queries — never concatenate user input into SQL strings.

---

### Q5. How do you use GORM for database operations?

```go
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

type Product struct {
    gorm.Model           // embeds ID, CreatedAt, UpdatedAt, DeletedAt
    Code  string         `gorm:"uniqueIndex"`
    Price float64
}

// Connect
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// Auto migrate — creates table if not exists
db.AutoMigrate(&Product{})

// Create
db.Create(&Product{Code: "D42", Price: 100})

// Read
var product Product
db.First(&product, 1)                    // by primary key
db.First(&product, "code = ?", "D42")   // with condition

// Find multiple
var products []Product
db.Where("price > ?", 50).Find(&products)

// Update
db.Model(&product).Update("Price", 200)
db.Model(&product).Updates(Product{Price: 200, Code: "F42"})

// Delete (soft delete if DeletedAt exists)
db.Delete(&product, 1)
```

---

### Q6. How do you handle connection pooling?

```go
db, _ := sql.Open("mysql", dsn)

// Pool settings — tune based on your workload
db.SetMaxOpenConns(25)       // max simultaneous connections to DB
db.SetMaxIdleConns(10)       // max idle connections in pool
db.SetConnMaxLifetime(1 * time.Hour)  // max time a connection can be reused
db.SetConnMaxIdleTime(10 * time.Minute) // max time a connection sits idle
```

> **Rule of thumb:** `MaxOpenConns` should match your DB's allowed concurrent connections divided by number of app instances.

---

### Q7. How do you prevent SQL injection in Go?

```go
// ❌ NEVER do this — SQL injection risk
userInput := "'; DROP TABLE users; --"
query := "SELECT * FROM users WHERE name = '" + userInput + "'"

// ✅ Always use parameterized queries
userInput := "'; DROP TABLE users; --"
row := db.QueryRow("SELECT * FROM users WHERE name = ?", userInput)
// The ? is parameterized — userInput treated as string, not SQL

// ✅ GORM is safe by default
db.Where("name = ?", userInput).Find(&users)
```

---

### Q8. How do you implement pagination?

```go
// SQL-level pagination
func getUsers(db *sql.DB, page, pageSize int) ([]User, error) {
    offset := (page - 1) * pageSize
    rows, err := db.Query(
        "SELECT id, name, email FROM users LIMIT ? OFFSET ?",
        pageSize, offset,
    )
    // ...scan rows
}

// GORM pagination
func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        offset := (page - 1) * pageSize
        return db.Offset(offset).Limit(pageSize)
    }
}
var users []User
db.Scopes(paginate(2, 10)).Find(&users)  // page 2, 10 per page
```

---

### Q9. How do you handle DB errors properly?

```go
import "database/sql"

row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 99)
var u User
err := row.Scan(&u.ID, &u.Name)
switch {
case err == sql.ErrNoRows:
    fmt.Println("no user found")
case err != nil:
    fmt.Println("query error:", err)
default:
    fmt.Println("found user:", u.Name)
}

// GORM
var user User
result := db.First(&user, 99)
if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    fmt.Println("record not found")
}
```
