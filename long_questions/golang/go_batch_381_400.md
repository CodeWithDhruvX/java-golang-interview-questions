## 🔴 Security and Advanced Testing (Questions 381-400)

### Question 381: How do you prevent SQL injection in Go?

**Answer:**
Always use **Parameterized Queries** (Prepared Statements). Never concatenate strings into queries.

**Vulnerable:**
```go
query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", unsafeInput)
db.Query(query) // DANGER!
```

**Secure:**
```go
// Use ? placeholders (or $1 for Postgres)
query := "SELECT * FROM users WHERE name = ?"
db.Query(query, safeInput)
```
The database driver escapes the input safely.

**ORMs:**
GORM and sqlx handle this automatically if used correctly (avoid `db.Raw` with string concatenation).

---

### Question 382: What are some common security vulnerabilities in Go apps?

**Answer:**
1. **Goroutine Leaks:** Denial of Service (DoS) by exhausting resources.
2. **Directory Traversal:** Using `filepath.Join` with user input without cleaning.
3. **Data Races:** Concurrent access causing corrupted state.
4. **Insecure Randomness:** Using `math/rand` for tokens (use `crypto/rand`).
5. **Cross-Site Scripting (XSS):** Rendering User Input in HTML templates without escaping.

---

### Question 383: How do you implement Secure Password Hashing?

**Answer:**
Use **bcrypt** or **Argon2**. Never use MD5 or SHA1/SHA256 directly.

**Library:** `golang.org/x/crypto/bcrypt`

```go
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

---

### Question 384: How do you validate input in Go APIs?

**Answer:**
1. **Manual Validation:** Checks in handlers.
2. **Validator Library:** `github.com/go-playground/validator` (standard with Gin).

```go
type User struct {
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=18"`
}

func validateUser(u User) error {
    validate := validator.New()
    return validate.Struct(u)
}
```

---

### Question 385: How do you implement JWT authentication?

**Answer:**
**Library:** `github.com/golang-jwt/jwt/v5`

**Sign (Login):**
```go
func createToken(user string) (string, error) {
    claims := jwt.MapClaims{
        "user": user,
        "exp":  time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte("my-secret-key"))
}
```

**Verify (Middleware):**
Parse the token string, provide the key callback, and check validity.

---

### Question 386: How do you handle secrets securely in Go?

**Answer:**
1. **Never commit to Git.**
2. **Environment Variables:** `os.Getenv("SECRET")`.
3. **Hashicorp Vault:** Fetch secrets at runtime.
4. **Memory Safety:** Use `[]byte` instead of `string` for sensitive data (can be wiped), though Go's GC makes this hard to guarantee.
5. **Avoid Logging:** Sanitize logs (don't log struct with "Password" field).

---

### Question 387: What is CSRF and how to mitigate it in Go?

**Answer:**
**CSRF (Cross-Site Request Forgery):** Attacker tricks user into performing action on trusted site (using cookies).

**Mitigation:**
1. **CSRF Tokens:**
   - Server sends random token in Cookie/HTML.
   - Client must send it back in Header `X-CSRF-Token`.
   - Library: `github.com/gorilla/csrf`.

2. **SameSite Cookies:**
   - Set `SameSite=Strict` or `Lax` on Auth cookies.

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session_token",
    Value:    "xyz",
    SameSite: http.SameSiteStrictMode,
    Secure:   true,
})
```

---

### Question 388: How do you test RESTful APIs in Go?

**Answer:**
Use `net/http/httptest`.

```go
func TestHealthCheck(t *testing.T) {
    // 1. Create Request
    req, _ := http.NewRequest("GET", "/health", nil)
    
    // 2. Create ResponseRecorder
    w := httptest.NewRecorder()
    
    // 3. Call Handler
    HealthHandler(w, req)
    
    // 4. Assert
    if w.Code != http.StatusOK {
        t.Errorf("Expected 200, got %d", w.Code)
    }
    if w.Body.String() != "OK" {
        t.Errorf("Unexpected body: %v", w.Body.String())
    }
}
```

---

### Question 389: What are Table-Driven Tests?

**Answer:**
The idiomatic way to write tests in Go. Data-driven approach.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"mixed", -1, 1, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("got %d, want %d", result, tc.expected)
            }
        })
    }
}
```

---

### Question 390: How do you mock HTTP calls in tests?

**Answer:**
1. **Interface Injection:** Define `HTTPClient` interface, mock implementation.
2. **httptest.Server:** Start a real internal server that mocks external responses.

```go
func TestExternalCall(t *testing.T) {
    // Mock Server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        w.Write([]byte(`{"status": "ok"}`))
    }))
    defer server.Close()

    // Use server.URL as API endpoint
    api := NewAPI(server.URL)
    result := api.Call()
    
    if result != "ok" { t.Fail() }
}
```

---

### Question 391: What are Flaky Tests and how to identify them?

**Answer:**
Tests that sometimes pass and sometimes fail without code changes.

**Causes:**
- Race conditions.
- Network dependencies (external API calls).
- Random number generation.
- Logic depending on time (e.g., `time.Now()`).

**Fixes:**
- Run with `-race`.
- Use Mocks for external calls.
- Run repeatedly: `go test -count=100 ./...`.

---

### Question 392: How do you generate test coverage reports?

**Answer:**
```bash
# Run tests and generate profile
go test -coverprofile=coverage.out ./...

# View detailed HTML report
go tool cover -html=coverage.out
```
This opens a browser showing exact lines covered (green) and missed (red).

---

### Question 393: What is Golden File Testing?

**Answer:**
Used for complex outputs (large JSON, HTML, Images). Instead of hardcoding the expected string, compare against a saved file.

```go
func TestJSON(t *testing.T) {
    got := generateBigJSON()
    
    if *update { // Flag to update golden file
        os.WriteFile("testdata/golden.json", got, 0644)
    }
    
    want, _ := os.ReadFile("testdata/golden.json")
    if string(got) != string(want) {
        t.Errorf("Mismatch found")
    }
}
```
Run `go test -update` to regenerate the expected output.

---

### Question 394: How do you mock database interactions?

**Answer:**
1. **Library: `go-sqlmock`:** Mock `sql/driver` behavior (rows, errors) without a real DB.
2. **Docker (Testcontainers):** Spin up a real Postgres container for integration tests (slower but more accurate).
3. **Repository Pattern:** Mock the `UserRepository` interface.

**Using sqlmock:**
```go
db, mock, _ := sqlmock.New()
mock.ExpectQuery("SELECT name FROM users").
    WithArgs(1).
    WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice"))

user, _ := GetUser(db, 1)
```

---

### Question 395: What is Fuzz Testing in Go?

**Answer:**
Native in Go 1.18+. Generates random inputs to find edge cases/crashes.

```go
func FuzzReverse(f *testing.F) {
    f.Add("hello") // Seed corpus
    
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
    })
}
```
Run: `go test -fuzz=Fuzz`

---

### Question 396: How do you benchmark Go code?

**Answer:**
Write a function starting with `Benchmark` in `_test.go`.

```go
func BenchmarkConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("hello %s", "world")
    }
}
```
Run: `go test -bench=.`
Look for `ns/op` (nanoseconds per operation) and `allocs/op`.

---

### Question 397: How do you handle Time in unit tests?

**Answer:**
Don't use `time.Now()` directly in logic.
Dependency Inject a `Clock` interface.

```go
type Clock interface {
    Now() time.Time
}

type RealClock struct{}
func (RealClock) Now() time.Time { return time.Now() }

type MockClock struct { t time.Time }
func (m MockClock) Now() time.Time { return m.t }
```
Now you can freeze time in tests!

---

### Question 398: What is `go:embed` and how does it help?

**Answer:**
Go 1.16+ feature to bundle static assets (HTML, SQL, Config) into the binary.

```go
import "embed"

//go:embed templates/*
var templateFS embed.FS

//go:embed config.json
var configFile []byte
```
Useful for:
- Single-binary deployments (Binary + Frontend).
- Database migration files.

---

### Question 399: What is the purpose of `init()` function?

**Answer:**
Runs automatically before `main()`.

**Uses:**
- Initializing global variables.
- Registering drivers (e.g., `image`, `database/sql`).

**Caveats:**
- Hard to test (side effects).
- Order of execution between files is alphabetical (tricky).
- Avoid complex logic or dependencies in `init()`.

---

### Question 400: Explain the Go memory model (briefly).

**Answer:**
Specifies conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes in another goroutine.

**Key Rule:** **"HAPPENS BEFORE"**
- A send on a channel **happens before** the corresponding receive completes.
- A lock release **happens before** the next acquire.
- `go` statement **happens before** the goroutine's execution.

If you don't use synchronization (Channels/Mutex), you have a **Data Race**, and visibility is not guaranteed!

---
