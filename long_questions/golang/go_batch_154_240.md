### Question 154: What is Gorilla Mux and how does it compare with net/http?

**Answer:**
Gorilla Mux is a powerful router with more features than net/http:

**net/http (basic):**
```go
http.HandleFunc("/users", usersHandler)
// Limited pattern matching
```

**Gorilla Mux (advanced):**
```go
import "github.com/gorilla/mux"

r := mux.NewRouter()

// URL parameters:
r.HandleFunc("/users/{id}", getUserHandler)

// Query parameters, methods, headers:
r.HandleFunc("/api/users", getUsers).Methods("GET")
r.HandleFunc("/api/users", createUser).Methods("POST")

// Host matching:
r.Host("api.example.com").HandleFunc("/", apiHandler)

// Subrouters:
api := r.PathPrefix("/api").Subrouter()
api.HandleFunc("/users", usersHandler)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]  // Get URL parameter
    fmt.Fprintf(w, "User ID: %s", id)
}
```

**Benefits of Gorilla Mux:**
- URL path variables
- Method-based routing
- Regex in routes
- Middleware support
- Subrouters

---

### Question 155: What are Go frameworks for web APIs (Gin, Echo)?

**Answer:**
Popular Go web frameworks:

**1. Gin - Fast HTTP framework:**
```go
import "github.com/gin-gonic/gin"

r := gin.Default()  // With logger and recovery

r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
})

r.POST("/users", func(c *gin.Context) {
    var user User
    c.BindJSON(&user)
    c.JSON(201, user)
})

r.Run(":8080")
```

**2. Echo - High performance framework:**
```go
import "github.com/labstack/echo/v4"

e := echo.New()

e.GET("/users/:id", func(c echo.Context) error {
    id := c.Param("id")
    return c.JSON(200, map[string]string{"id": id})
})

e.Start(":8080")
```

**3. Fiber - Express-inspired:**
```go
import "github.com/gofiber/fiber/v2"

app := fiber.New()

app.Get("/users/:id", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"id": c.Params("id")})
})

app.Listen(":8080")
```

---

### Question 156: What are the trade-offs between using http.ServeMux and third-party routers?

**Answer:**

**http.ServeMux (Standard library):**

**Pros:**
- No external dependencies
- Simple and lightweight
- Sufficient for basic routing
- Well-tested and stable

**Cons:**
- No URL parameters
- No regex matching
- Limited pattern matching
- No middleware support built-in

**Third-party routers (Gorilla Mux, Chi, etc.):**

**Pros:**
- URL parameters: `/users/{id}`
- Method-based routing
- Middleware support
- Advanced features
- Better developer experience

**Cons:**
- External dependency
- Slightly more overhead
- Learning curve

**When to use each:**
- **Use ServeMux** for simple APIs, microservices with few routes
- **Use third-party** for complex routing needs, RESTful APIs

---

### Question 157: How would you implement authentication in a Go API?

**Answer:**
Implement JWT-based authentication:

```go
import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

// Generate JWT token:
func generateToken(userID int) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Verify token middleware:
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        
        // Remove "Bearer " prefix:
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", 401)
            return
        }
        
        // Add user ID to context:
        ctx := context.WithValue(r.Context(), "userID", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Login handler:
func login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&creds)
    
    // Verify credentials (simplified):
    if creds.Username == "admin" && creds.Password == "password" {
        token, _ := generateToken(1)
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    } else {
        http.Error(w, "Invalid credentials", 401)
    }
}
```

---

### Question 158: How do you implement file streaming in Go?

**Answer:**
Stream large files efficiently:

```go
func streamFile(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("large-video.mp4")
    if err != nil {
        http.Error(w, "File not found", 404)
        return
    }
    defer file.Close()
    
    // Get file info:
    stat, _ := file.Stat()
    
    // Set headers:
    w.Header().Set("Content-Type", "video/mp4")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
    
    // Stream file:
    io.Copy(w, file)  // Efficiently streams in chunks
}

// Stream with buffer control:
func streamWithBuffer(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("large-file.dat")
    defer file.Close()
    
    buffer := make([]byte, 4096)  // 4KB chunks
    
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        w.Write(buffer[:n])
        
        // Flush to client:
        if f, ok := w.(http.Flusher); ok {
            f.Flush()
        }
    }
}
```

---

### Question 159: What is the HTTP/2 Server Push and how to use it?

**Answer:**
HTTP/2 Server Push sends resources before the client requests them:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Check if pusher is supported:
    pusher, ok := w.(http.Pusher)
    if ok {
        // Push CSS before client requests it:
        if err := pusher.Push("/style.css", nil); err != nil {
            log.Printf("Failed to push: %v", err)
        }
    }
    
    // Serve main HTML:
    http.ServeFile(w, r, "index.html")
}

func main() {
    http.HandleFunc("/", handler)
    
    // HTTP/2 requires HTTPS:
    log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
```

---

## 🟣 Databases and ORMs (Questions 160-180)

### Question 160: How do you connect to a PostgreSQL database in Go?

**Answer:**
Use database/sql with a PostgreSQL driver:

```go
import (
    "database/sql"
    _ "github.com/lib/pq"  // PostgreSQL driver
)

func main() {
    connStr := "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Test connection:
    if err := db.Ping(); err != nil {
        log.Fatal("Cannot connect:", err)
    }
    
    fmt.Println("Connected to PostgreSQL!")
}
```

---

### Question 161: What is the difference between database/sql and GORM?

**Answer:**

**database/sql (Standard library):**
- Low-level SQL operations
- Manual query writing
- Full control
- No magic/hidden queries
- More verbose

```go
rows, err := db.Query("SELECT id, name FROM users WHERE age > $1", 18)
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
}
```

**GORM (ORM):**
- High-level object mapping
- Auto migrations
- Associations/relations
- Less boilerplate
- Query builder

```go
type User struct {
    ID   uint
    Name string
    Age  int
}

db.AutoMigrate(&User{})

var users []User
db.Where("age > ?", 18).Find(&users)
```

**When to use:**
- **database/sql**: Performance-critical, complex queries, full control
- **GORM**: Rapid development, CRUD operations, less SQL knowledge needed

---

### Question 162: How do you handle SQL injections in Go?

**Answer:**
Always use parameterized queries:

**❌ BAD (Vulnerable to SQL injection):**
```go
username := r.FormValue("username")
query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
rows, _ := db.Query(query)  // DANGEROUS!
```

**✅ GOOD (Safe from SQL injection):**
```go
username := r.FormValue("username")
rows, _ := db.Query("SELECT * FROM users WHERE username = $1", username)
```

**Why it works:**
- Database driver escapes parameters
- SQL and data are sent separately
- No string concatenation

**Examples:**
```go
// INSERT:
stmt, _ := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
stmt.Exec(name, email)

// UPDATE:
db.Exec("UPDATE users SET name = $1 WHERE id = $2", newName, userID)

// DELETE:
db.Exec("DELETE FROM users WHERE id = $1", userID)

// Multiple parameters:
db.Query("SELECT * FROM orders WHERE user_id = $1 AND status = $2", userID, status)
```

---

### Question 163: How do you manage connection pools in database/sql?

**Answer:**
Configure connection pool settings:

```go
db, err := sql.Open("postgres", connStr)

// Maximum number of open connections:
db.SetMaxOpenConns(25)

// Maximum number of idle connections:
db.SetMaxIdleConns(5)

// Maximum lifetime of a connection:
db.SetConnMaxLifetime(5 * time.Minute)

// Maximum idle time:
db.SetConnMaxIdleTime(10 * time.Minute)

// Get connection stats:
stats := db.Stats()
fmt.Printf("Open connections: %d\n", stats.OpenConnections)
fmt.Printf("In use: %d\n", stats.InUse)
fmt.Printf("Idle: %d\n", stats.Idle)
```

**Best practices:**
- Set MaxOpenConns to match database server capacity
- Keep MaxIdleConns reasonable (5-10)
- Set ConnMaxLifetime to prevent stale connections
- Monitor stats in production

---

### Question 164: What are prepared statements in Go?

**Answer:**
Prepared statements compile SQL once and reuse it:

```go
// Prepare statement:
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Execute multiple times:
for _, user := range users {
    _, err := stmt.Exec(user.Name, user.Email)
    if err != nil {
        log.Printf("Error inserting %s: %v", user.Name, err)
    }
}

// SELECT with prepared statement:
stmt, _ = db.Prepare("SELECT name, email FROM users WHERE id = $1")
defer stmt.Close()

var name, email string
err = stmt.QueryRow(123).Scan(&name, &email)
```

**Benefits:**
- Better performance (compiled once)
- Protection against SQL injection
- Reusable across multiple executions

---

### Question 165: How do you map SQL rows to structs?

**Answer:**
Use Scan to map rows to struct fields:

```go
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

// Single row:
func getUser(db *sql.DB, id int) (*User, error) {
    user := &User{}
    
    err := db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    
    return user, err
}

// Multiple rows:
func getUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email, age FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    
    return users, rows.Err()
}

// Using sqlx for automatic mapping:
import "github.com/jmoiron/sqlx"

var users []User
db.Select(&users, "SELECT * FROM users")  // Automatic mapping!
```

---

### Question 166: What are transactions and how are they implemented in Go?

**Answer:**
Transactions ensure atomic operations:

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    // Begin transaction:
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Defer rollback (will be skipped if committed):
    defer tx.Rollback()
    
    // Deduct from sender:
    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err  // Transaction will be rolled back
    }
    
    // Add to receiver:
    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err  // Transaction will be rolled back
    }
    
    // Commit transaction:
    return tx.Commit()
}

// Using context for timeout:
func transferWithTimeout(db *sql.DB, fromID, toID int, amount float64) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    tx, _ := db.BeginTx(ctx, nil)
    defer tx.Rollback()
    
    _, err := tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

---

### Question 167: How do you handle database migrations in Go?

**Answer:**
Use migration tools like golang-migrate:

**Using golang-migrate:**
```bash
# Install:
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration:
migrate create -ext sql -dir migrations -seq create_users_table
```

**Migration files:**
```sql
-- migrations/000001_create_users_table.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- migrations/000001_create_users_table.down.sql
DROP TABLE users;
```

**Run migrations in code:**
```go
import "github.com/golang-migrate/migrate/v4"
import _ "github.com/golang-migrate/migrate/v4/database/postgres"
import _ "github.com/golang-migrate/migrate/v4/source/file"

func runMigrations(dbURL string) error {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        return err
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    
    return nil
}
```

---

### Question 168: What is the use of sqlx in Go?

**Answer:**
sqlx extends database/sql with convenient features:

```go
import "github.com/jmoiron/sqlx"

db, _ := sqlx.Connect("postgres", connStr)

type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
}

// Get single row:
var user User
err := db.Get(&user, "SELECT * FROM users WHERE id = $1", 123)

// Get multiple rows:
var users []User
err = db.Select(&users, "SELECT * FROM users WHERE age > $1", 18)

// Named queries:
query := `INSERT INTO users (name, email) VALUES (:name, :email)`
_, err = db.NamedExec(query, map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
})

// Struct binding:
newUser := User{Name: "Jane", Email: "jane@example.com"}
_, err = db.NamedExec(`INSERT INTO users (name, email) VALUES (:name, :email)`, newUser)

// In queries:
query, args, _ := sqlx.In("SELECT * FROM users WHERE id IN (?)", []int{1, 2, 3})
query = db.Rebind(query)  // Convert ? to $1, $2 for PostgreSQL
db.Select(&users, query, args...)
```

---

### Question 169: What are the pros and cons of using an ORM in Go?

**Answer:**

**Pros:**
1. **Faster development** - Less boilerplate code
2. **Auto migrations** - Schema management
3. **Type safety** - Compile-time checks
4. **Relationship handling** - Automatic joins
5. **Database agnostic** - Easy to switch databases
6. **Query builder** - Readable, chainable queries

**Cons:**
1. **Performance overhead** - Generated queries not always optimal
2. **Learning curve** - Need to learn ORM API
3. **Hidden complexity** - Hard to debug generated SQL
4. **N+1 problem** - Can generate many queries
5. **Complex queries** - Sometimes easier in raw SQL
6. **Black box** - Less control over exact SQL

**Example:**
```go
// GORM (ORM):
var users []User
db.Where("age > ?", 18).Find(&users)

// vs database/sql (raw):
rows, _ := db.Query("SELECT * FROM users WHERE age > $1", 18)
defer rows.Close()
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name, &u.Age)
    users = append(users, u)
}
```

**Recommendation:**
- **Use ORM** for standard CRUD, rapid prototyping
- **Use raw SQL** for complex queries, performance-critical code
- Mix both as needed

---

### Question 170: How would you implement pagination in SQL queries?

**Answer:**
Use LIMIT and OFFSET:

```go
type PaginationParams struct {
    Page     int
    PageSize int
}

func getUsers(db *sql.DB, params PaginationParams) ([]User, error) {
    // Calculate offset:
    offset := (params.Page - 1) * params.PageSize
    
    query := `
        SELECT id, name, email
        FROM users
        ORDER BY id
        LIMIT $1 OFFSET $2
    `
    
    rows, err := db.Query(query, params.PageSize, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return users, nil
}

// With total count:
type PaginatedResponse struct {
    Data       []User `json:"data"`
    Page       int    `json:"page"`
    PageSize   int    `json:"page_size"`
    TotalCount int    `json:"total_count"`
    TotalPages int    `json:"total_pages"`
}

func getUsersPaginated(db *sql.DB, params PaginationParams) (*PaginatedResponse, error) {
    // Get total count:
    var totalCount int
    db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalCount)
    
    // Get paginated data:
    offset := (params.Page - 1) * params.PageSize
    rows, _ := db.Query("SELECT id, name, email FROM users LIMIT $1 OFFSET $2", 
        params.PageSize, offset)
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return &PaginatedResponse{
        Data:       users,
        Page:       params.Page,
        PageSize:   params.PageSize,
        TotalCount: totalCount,
        TotalPages: (totalCount + params.PageSize - 1) / params.PageSize,
    }, nil
}

// Cursor-based pagination (better for large datasets):
func getUsersCursorBased(db *sql.DB, cursor int, limit int) ([]User, error) {
    query := `
        SELECT id, name, email
        FROM users
        WHERE id > $1
        ORDER BY id
        LIMIT $2
    `
    
    rows, _ := db.Query(query, cursor, limit)
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    
    return users, nil
}
```

---

### Question 171: How do you log SQL queries in Go?

**Answer:**
Implement query logging:

**Method 1: Custom logger wrapper:**
```go
type LoggedDB struct {
    *sql.DB
}

func (db *LoggedDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
    start := time.Now()
    
    log.Printf("[SQL Query] %s | Args: %v", query, args)
    
    rows, err := db.DB.Query(query, args...)
    
    duration := time.Since(start)
    if err != nil {
        log.Printf("[SQL Error] %v | Duration: %v", err, duration)
    } else {
        log.Printf("[SQL Success] Duration: %v", duration)
    }
    
    return rows, err
}

func (db *LoggedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
    start := time.Now()
    log.Printf("[SQL Exec] %s | Args: %v", query, args)
    
    result, err := db.DB.Exec(query, args...)
    
    duration := time.Since(start)
    log.Printf("[SQL] Duration: %v | Error: %v", duration, err)
    
    return result, err
}
```

**Method 2: GORM logging:**
```go
import "gorm.io/gorm/logger"

newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags),
    logger.Config{
        SlowThreshold: time.Second,   // Slow SQL threshold
        LogLevel:      logger.Info,   // Log level
        Colorful:      true,          // Color output
    },
)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: newLogger,
})
```

---

### Question 172: What is the N+1 problem in ORMs and how to avoid it?

**Answer:**
N+1 problem occurs when fetching related data in a loop:

**❌ BAD (N+1 Problem):**
```go
// 1 query to get users:
var users []User
db.Find(&users)  // SELECT * FROM users

// N queries for each user's posts:
for i := range users {
    db.Model(&users[i]).Association("Posts").Find(&users[i].Posts)
    // SELECT * FROM posts WHERE user_id = 1
    // SELECT * FROM posts WHERE user_id = 2
    // SELECT * FROM posts WHERE user_id = 3
    // ... N more queries!
}
```

**✅ GOOD (Using Preload):**
```go
// Single query with JOIN:
var users []User
db.Preload("Posts").Find(&users)
// SELECT * FROM users
// SELECT * FROM posts WHERE user_id IN (1,2,3,...)
```

**Other solutions:**
```go
// Eager loading multiple associations:
db.Preload("Posts").Preload("Comments").Find(&users)

// Nested preloading:
db.Preload("Posts.Comments").Find(&users)

// Joins (single query):
db.Joins("Posts").Find(&users)

// Raw SQL with proper joins:
db.Raw(`
    SELECT users.*, posts.*
    FROM users
    LEFT JOIN posts ON posts.user_id = users.id
`).Scan(&results)
```

---

### Question 173: How do you implement caching for DB queries in Go?

**Answer:**
Use Redis or in-memory cache:

**Using Redis:**
```go
import (
    "github.com/go-redis/redis/v8"
    "encoding/json"
)

var rdb *redis.Client

func getUser(db *sql.DB, id int) (*User, error) {
    ctx := context.Background()
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Try cache first:
    cached, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // Cache miss - query database:
    var user User
    err = db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        return nil, err
    }
    
    // Store in cache:
    userData, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, userData, 10*time.Minute)
    
    return &user, nil
}

// Invalidate cache on update:
func updateUser(db *sql.DB, user *User) error {
    _, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3",
        user.Name, user.Email, user.ID)
    
    if err == nil {
        // Invalidate cache:
        cacheKey := fmt.Sprintf("user:%d", user.ID)
        rdb.Del(context.Background(), cacheKey)
    }
    
    return err
}
```

**Using in-memory cache (sync.Map):**
```go
var cache sync.Map

type CacheItem struct {
    Data      interface{}
    ExpiresAt time.Time
}

func getCached(key string) (interface{}, bool) {
    val, ok := cache.Load(key)
    if !ok {
        return nil, false
    }
    
    item := val.(CacheItem)
    if time.Now().After(item.ExpiresAt) {
        cache.Delete(key)
        return nil, false
    }
    
    return item.Data, true
}

func setCache(key string, data interface{}, ttl time.Duration) {
    cache.Store(key, CacheItem{
        Data:      data,
        ExpiresAt: time.Now().Add(ttl),
    })
}
```

---

### Question 174: How do you write custom SQL queries using GORM?

**Answer:**
GORM provides multiple ways to execute raw SQL:

```go
// Raw SQL query:
var users []User
db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&users)

// With named parameters:
db.Raw("SELECT * FROM users WHERE name = @name", 
    sql.Named("name", "John")).Scan(&users)

// Execute SQL:
db.Exec("UPDATE users SET active = ? WHERE last_login < ?", true, time.Now())

// Complex query:
var result []map[string]interface{}
db.Raw(`
    SELECT u.name, COUNT(p.id) as post_count
    FROM users u
    LEFT JOIN posts p ON p.user_id = u.id
    GROUP BY u.name
    HAVING COUNT(p.id) > 5
`).Scan(&result)

// Mix raw SQL with GORM:
db.Where("age > ?", 18).
    Where("created_at > ?", time.Now().AddDate(0, -1, 0)).
    Find(&users)

// Combine with Scan:
type UserStats struct {
    UserID     int
    PostCount  int
    CommentCount int
}

var stats []UserStats
db.Raw(`
    SELECT 
        u.id as user_id,
        COUNT(DISTINCT p.id) as post_count,
        COUNT(DISTINCT c.id) as comment_count
    FROM users u
    LEFT JOIN posts p ON p.user_id = u.id
    LEFT JOIN comments c ON c.user_id = u.id
    GROUP BY u.id
`).Scan(&stats)
```

---

### Question 175: How do you handle one-to-many and many-to-many relationships in GORM?

**Answer:**

**One-to-Many:**
```go
type User struct {
    ID    uint
    Name  string
    Posts []Post  // Has many posts
}

type Post struct {
    ID     uint
    Title  string
    UserID uint    // Foreign key
    User   User    // Belongs to user
}

// Create with associations:
user := User{
    Name: "John",
    Posts: []Post{
        {Title: "First Post"},
        {Title: "Second Post"},
    },
}
db.Create(&user)

// Query with associations:
var user User
db.Preload("Posts").First(&user, 1)

// Add post to existing user:
post := Post{Title: "New Post", UserID: user.ID}
db.Create(&post)
```

**Many-to-Many:**
```go
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID    uint
    Name  string
    Users []User `gorm:"many2many:user_roles;"`
}

// GORM creates join table automatically:
// CREATE TABLE user_roles (user_id INT, role_id INT)

// Create with associations:
user := User{
    Name: "John",
    Roles: []Role{
        {Name: "Admin"},
        {Name: "Editor"},
    },
}
db.Create(&user)

// Query:
var user User
db.Preload("Roles").First(&user, 1)

// Add role to user:
var admin Role
db.First(&admin, "name = ?", "Admin")
db.Model(&user).Association("Roles").Append(&admin)

// Remove role:
db.Model(&user).Association("Roles").Delete(&admin)

// Replace all roles:
db.Model(&user).Association("Roles").Replace(&newRoles)

// Clear all:
db.Model(&user).Association("Roles").Clear()

// Count:
count := db.Model(&user).Association("Roles").Count()
```

---

### Question 176: How would you structure your database layer in a Go project?

**Answer:**
Use repository pattern for clean architecture:

```
internal/
├── domain/
│   └── user.go           # Entity definitions
├── repository/
│   ├── interface.go      # Repository interfaces
│   └── postgres/
│       └── user_repo.go  # PostgreSQL implementation
└── service/
    └── user_service.go   # Business logic
```

**domain/user.go:**
```go
package domain

type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}
```

**repository/interface.go:**
```go
package repository

type UserRepository interface {
    Create(user *domain.User) error
    GetByID(id int) (*domain.User, error)
    GetAll() ([]domain.User, error)
    Update(user *domain.User) error
    Delete(id int) error
}
```

**repository/postgres/user_repo.go:**
```go
package postgres

type PostgresUserRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(user *domain.User) error {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *PostgresUserRepo) GetByID(id int) (*domain.User, error) {
    user := &domain.User{}
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    err := r.db.QueryRow(query, id).Scan(
        &user.ID, &user.Name, &user.Email, &user.CreatedAt,
    )
    return user, err
}

func (r *PostgresUserRepo) GetAll() ([]domain.User, error) {
    rows, err := r.db.Query(`SELECT id, name, email FROM users`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []domain.User
    for rows.Next() {
        var u domain.User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    return users, nil
}
```

**service/user_service.go:**
```go
package service

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
    // Business logic validation:
    if !isValidEmail(email) {
        return errors.New("invalid email")
    }
    
    user := &domain.User{Name: name, Email: email}
    return s.repo.Create(user)
}
```

---

### Question 177: What is context propagation in database calls?

**Answer:**
Pass context through database operations for timeout/cancellation:

```go
// Create context with timeout:
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use context in query:
rows, err := db.QueryContext(ctx, "SELECT * FROM large_table")
if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Query timed out!")
    }
    return err
}

// In transactions:
func transferMoney(ctx context.Context, db *sql.DB, from, to int, amount float64) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}

// With HTTP request context:
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // Get request context
    
    var users []User
    rows, err := db.QueryContext(ctx, "SELECT * FROM users")
    // Query automatically cancelled if client disconnects
}

// Repository with context:
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*User, error) {
    user := &User{}
    err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Name)
    return user, err
}
```

---

### Question 178: How do you handle long-running queries or timeouts?

**Answer:**
Use context timeouts and query optimization:

```go
// Set query timeout:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, `
    SELECT * FROM orders 
    WHERE created_at > $1
    ORDER BY id DESC
`, time.Now().AddDate(0, -1, 0))

if err != nil {
    if err == context.DeadlineExceeded {
        log.Println("Query took too long, consider optimizing or adding index")
    }
    return err
}

// For very long queries, use channels:
func longRunningQuery(db *sql.DB) <-chan QueryResult {
    results := make(chan QueryResult)
    
    go func() {
        defer close(results)
        
        rows, err := db.Query("SELECT * FROM massive_table")
        if err != nil {
            results <- QueryResult{Error: err}
            return
        }
        defer rows.Close()
        
        for rows.Next() {
            var data Data
            rows.Scan(&data.ID, &data.Value)
            results <- QueryResult{Data: data}
        }
    }()
    
    return results
}

// Use with timeout:
func processWithTimeout() {
    resultsChan := longRunningQuery(db)
    timeout := time.After(1 * time.Minute)
    
    for {
        select {
        case result, ok := <-resultsChan:
            if !ok {
                return  // Channel closed
            }
            processResult(result)
        case <-timeout:
            log.Println("Query timeout!")
            return
        }
    }
}

// Pagination for large datasets:
func processLargeDataset(db *sql.DB) error {
    pageSize := 1000
    offset := 0
    
    for {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        
        rows, err := db.QueryContext(ctx, 
            "SELECT * FROM large_table LIMIT $1 OFFSET $2",
            pageSize, offset)
        
        cancel()
        
        if err != nil {
            return err
        }
        
        count := 0
        for rows.Next() {
            // Process row
            count++
        }
        rows.Close()
        
        if count < pageSize {
            break  // Last page
        }
        
        offset += pageSize
    }
    
    return nil
}
```

---

### Question 179: How do you write unit tests for code that interacts with the DB?

**Answer:**
Use mocks, interfaces, and test databases:

**Method 1: Interface mocking:**
```go
// Define interface:
type UserRepository interface {
    GetUser(id int) (*User, error)
    CreateUser(user *User) error
}

// Real implementation:
type SQLUserRepo struct {
    db *sql.DB
}

func (r *SQLUserRepo) GetUser(id int) (*User, error) {
    // Real database code
}

// Mock implementation for tests:
type MockUserRepo struct {
    users map[int]*User
}

func (m *MockUserRepo) GetUser(id int) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, errors.New("not found")
    }
    return user, nil
}

func (m *MockUserRepo) CreateUser(user *User) error {
    m.users[user.ID] = user
    return nil
}

// Test:
func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepo{
        users: map[int]*User{
            1: {ID: 1, Name: "John"},
        },
    }
    
    service := NewUserService(mockRepo)
    
    user, err := service.GetUser(1)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Expected John, got %s", user.Name)
    }
}
```

**Method 2: Using sqlmock:**
```go
import "github.com/DATA-DOG/go-sqlmock"

func TestGetUser(t *testing.T) {
    // Create mock database:
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock: %v", err)
    }
    defer db.Close()
    
    // Set expectations:
    rows := sqlmock.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "John", "john@example.com")
    
    mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
        WithArgs(1).
        WillReturnRows(rows)
    
    // Test:
    repo := NewUserRepository(db)
    user, err := repo.GetUser(1)
    
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if user.Name != "John" {
        t.Errorf("Expected John, got %s", user.Name)
    }
    
    // Verify expectations:
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Unfulfilled expectations: %v", err)
    }
}
```

**Method 3: Test database:**
```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("postgres", "postgresql://localhost/test_db")
    if err != nil {
        t.Fatal(err)
    }
    
    // Run migrations:
    db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100)
    )`)
    
    t.Cleanup(func() {
        db.Exec("TRUNCATE TABLE users")
        db.Close()
    })
    
    return db
}

func TestUserRepositoryIntegration(t *testing.T) {
    db := setupTestDB(t)
    repo := NewUserRepository(db)
    
    user := &User{Name: "John", Email: "john@example.com"}
    err := repo.CreateUser(user)
    
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    }
    
    retrieved, err := repo.GetUser(user.ID)
    if retrieved.Name != "John" {
        t.Errorf("Expected John, got %s", retrieved.Name)
    }
}
```

---

### Question 180: What is connection string format for different databases?

**Answer:**

**PostgreSQL:**
```go
// Standard format:
connStr := "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"

// URL format:
connStr := "postgresql://postgres:secret@localhost:5432/mydb?sslmode=disable"

// With connection pool settings:
db, _ := sql.Open("postgres", connStr)
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
```

**MySQL:**
```go
// Format: user:password@tcp(host:port)/dbname
connStr := "user:password@tcp(localhost:3306)/mydb?parseTime=true"

// With charset:
connStr := "user:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=true"
```

**SQLite:**
```go
// File-based:
db, _ := sql.Open("sqlite3", "./mydb.db")

// In-memory:
db, _ := sql.Open("sqlite3", ":memory:")

// With options:
db, _ := sql.Open("sqlite3", "file:mydb.db?cache=shared&mode=rwc")
```

**MongoDB:**
```go
import "go.mongodb.org/mongo-driver/mongo"

// Standard:
client, _ := mongo.Connect(ctx, options.Client().
    ApplyURI("mongodb://localhost:27017"))

// With auth:
uri := "mongodb://username:password@localhost:27017/mydb?authSource=admin"
client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

// Replica set:
uri := "mongodb://host1:27017,host2:27017,host3:27017/mydb?replicaSet=rs0"
```

**Redis:**
```go
import "github.com/go-redis/redis/v8"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",  // no password
    DB:       0,   // default DB
})

// With password:
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "secret",
    DB:       0,
})

// Cluster:
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002"},
})
```

---

*[Questions 181-240 will be added in the next batch to manage file size and ensure quality]*
