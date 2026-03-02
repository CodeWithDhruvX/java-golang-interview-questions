# 📕 04 — Error Handling in Go
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- The `error` interface
- Custom error types
- `fmt.Errorf` with `%w` (error wrapping)
- `errors.Is` and `errors.As`
- Panic and recover
- Sentinel errors

---

## ❓ Most Asked Questions

### Q1. How does Go handle errors? Why no exceptions?

```go
// Go uses explicit error return values — no try/catch
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```
> **Why no exceptions?** Go's philosophy: errors are values, not control flow. This makes error handling explicit, predictable, and easy to trace.

---

### Q2. How do you create custom error types?

```go
// Method 1: errors.New (simple)
var ErrNotFound = errors.New("not found")

// Method 2: fmt.Errorf (formatted)
func findUser(id int) error {
    return fmt.Errorf("user with id %d not found", id)
}

// Method 3: Custom struct (rich error info)
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "must be non-negative"}
    }
    return nil
}
```

---

### Q3. What is error wrapping? How do `errors.Is` and `errors.As` work?

```go
var ErrNotFound = errors.New("not found")

// Wrap an error with context using %w
func getProduct(id int) error {
    return fmt.Errorf("getProduct(%d): %w", id, ErrNotFound)
}

err := getProduct(42)

// errors.Is — checks if any wrapped error in the chain matches
if errors.Is(err, ErrNotFound) {
    fmt.Println("Product not found!")  // ✅ true
}

// errors.As — extracts the error to a specific type
type DBError struct{ Code int }
func (e *DBError) Error() string { return fmt.Sprintf("db error: %d", e.Code) }

func queryDB() error {
    return fmt.Errorf("query failed: %w", &DBError{Code: 500})
}

var dbErr *DBError
if errors.As(queryDB(), &dbErr) {
    fmt.Println("DB error code:", dbErr.Code)  // 500
}
```

---

### Q4. What are sentinel errors? When should you use them?

```go
// Sentinel errors — predefined errors for comparison
var (
    ErrNotFound    = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrDuplicate   = errors.New("duplicate entry")
)

func handleRequest(id int) error {
    user, err := findUser(id)
    if errors.Is(err, ErrNotFound) {
        // handle not found
    }
    _ = user
    return nil
}
```
> **Best practice:** Use sentinel errors for known, expected conditions. Use custom struct errors when you need to carry additional context.

---

### Q5. What is panic and recover?

```go
// panic — stops normal execution, unwinds stack, runs deferred functions
func mustDivide(a, b int) int {
    if b == 0 {
        panic("cannot divide by zero")
    }
    return a / b
}

// recover — catches a panic, must be called inside a deferred function
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    result = mustDivide(a, b)
    return
}

res, err := safeDiv(10, 0)
fmt.Println(res, err)  // 0, recovered from panic: cannot divide by zero
```

---

### Q6. What are best practices for error handling in Go?

```go
// ✅ 1. Always handle errors — don't ignore
result, err := doSomething()
if err != nil {
    return fmt.Errorf("context: %w", err)  // wrap with context
}

// ✅ 2. Return early on error (avoid deep nesting)
func processRequest(r *http.Request) error {
    data, err := parseBody(r)
    if err != nil {
        return fmt.Errorf("parseBody: %w", err)
    }
    user, err := findUser(data.UserID)
    if err != nil {
        return fmt.Errorf("findUser(%d): %w", data.UserID, err)
    }
    return saveUser(user)
}

// ✅ 3. Use errors.Is / errors.As instead of string comparison
// ❌ Bad
if err.Error() == "not found" { }
// ✅ Good
if errors.Is(err, ErrNotFound) { }

// ✅ 4. Use panic only for programming errors (e.g., invalid args)
// Use error returns for expected/runtime errors
```

---

### Q7. What is `errors.Join` (Go 1.20+)?

```go
// Combine multiple errors into one
err1 := errors.New("validation failed")
err2 := errors.New("db error")
err3 := errors.New("cache miss")

combined := errors.Join(err1, err2, err3)
fmt.Println(combined)
// Output:
// validation failed
// db error
// cache miss

// errors.Is works with Joined errors
fmt.Println(errors.Is(combined, err1))  // true
```

---

### Q8. How do you handle errors in goroutines?

```go
// Use error channel pattern
errCh := make(chan error, 1)

go func() {
    if err := riskyOperation(); err != nil {
        errCh <- err
        return
    }
    errCh <- nil
}()

if err := <-errCh; err != nil {
    fmt.Println("goroutine error:", err)
}

// For multiple goroutines — use errgroup
import "golang.org/x/sync/errgroup"

g, ctx := errgroup.WithContext(context.Background())
_ = ctx

g.Go(func() error { return fetchData("url1") })
g.Go(func() error { return fetchData("url2") })

if err := g.Wait(); err != nil {
    fmt.Println("one of the goroutines failed:", err)
}
```
