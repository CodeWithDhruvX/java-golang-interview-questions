# 🧪 07 — Testing in Go
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `testing` package basics
- Table-driven tests
- `testify` library (assert, require, mock)
- Testing HTTP handlers (`httptest`)
- Mocking interfaces
- Test coverage

---

## ❓ Most Asked Questions

### Q1. How do you write a basic unit test in Go?

```go
// File: math.go
package math

func Add(a, b int) int    { return a + b }
func Subtract(a, b int) int { return a - b }

// File: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(3, 4)
    if result != 7 {
        t.Errorf("Add(3,4) = %d; want 7", result)
    }
}

func TestSubtract(t *testing.T) {
    got  := Subtract(10, 3)
    want := 7
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

```bash
go test ./...          # run all tests
go test -v ./...       # verbose output
go test -run TestAdd   # run specific test
```

---

### Q2. What is table-driven testing?

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
        hasError bool
    }{
        {"normal division", 10, 2, 5.0, false},
        {"divide by zero", 10, 0, 0, true},
        {"decimal result", 7, 2, 3.5, false},
        {"negative numbers", -6, 2, -3.0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.a, tt.b)
            if tt.hasError {
                if err == nil { t.Errorf("expected error but got nil") }
                return
            }
            if err != nil { t.Errorf("unexpected error: %v", err) }
            if got != tt.expected {
                t.Errorf("Divide(%.1f, %.1f) = %.1f; want %.1f", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

---

### Q3. How do you use `testify` for better assertions?

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestUserService(t *testing.T) {
    user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

    assert.Equal(t, 1, user.ID)                // test continues if fails
    assert.NotNil(t, user)
    assert.Equal(t, "Alice", user.Name)
    assert.Contains(t, user.Email, "@")

    require.NotEmpty(t, user.Email)  // test STOPS if this fails (like Fatal)
    require.True(t, user.ID > 0)
}
```

---

### Q4. How do you test HTTP handlers using `httptest`?

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
)

func TestGetUserHandler(t *testing.T) {
    // Create request
    req, err := http.NewRequest(http.MethodGet, "/users/1", nil)
    if err != nil { t.Fatal(err) }

    // Create response recorder
    rr := httptest.NewRecorder()

    // Call handler
    handler := http.HandlerFunc(getUserHandler)
    handler.ServeHTTP(rr, req)

    // Assert status
    assert.Equal(t, http.StatusOK, rr.Code)

    // Assert body
    var user User
    json.NewDecoder(rr.Body).Decode(&user)
    assert.Equal(t, 1, user.ID)
    assert.Equal(t, "Alice", user.Name)
}
```

---

### Q5. How do you mock interfaces for unit testing?

```go
// Define interface
type UserRepository interface {
    GetByID(id int) (*User, error)
    Save(u *User) error
}

// Implement a mock manually
type MockUserRepo struct {
    users map[int]*User
}

func (m *MockUserRepo) GetByID(id int) (*User, error) {
    u, ok := m.users[id]
    if !ok { return nil, ErrNotFound }
    return u, nil
}

func (m *MockUserRepo) Save(u *User) error {
    m.users[u.ID] = u
    return nil
}

// Test using mock
func TestUserService_GetUser(t *testing.T) {
    repo := &MockUserRepo{users: map[int]*User{
        1: {ID: 1, Name: "Alice"},
    }}
    svc := NewUserService(repo)

    user, err := svc.GetUser(1)
    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
}

// Or use testify/mock for auto-generated mocks
import "github.com/stretchr/testify/mock"
type MockRepo struct{ mock.Mock }
func (m *MockRepo) GetByID(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}
```

---

### Q6. How do you check test coverage?

```bash
go test -cover ./...                          # show coverage %
go test -coverprofile=coverage.out ./...      # save detailed report
go tool cover -html=coverage.out              # open HTML report in browser
go tool cover -func=coverage.out              # per-function coverage
```

> **Target:** Aim for 70–80%+ coverage on core business logic. Don't chase 100% — focus on meaningful tests.

---

### Q7. How do you write benchmark tests?

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)  // b.N is determined automatically by Go
    }
}

// Run benchmarks
// go test -bench=. -benchmem ./...
// -benchmem shows memory allocations
```

---

### Q8. How do you test for expected panics?

```go
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil { panic("invalid number: " + s) }
    return n
}

func TestMustParse_Panics(t *testing.T) {
    assert.Panics(t, func() {
        mustParse("not-a-number")
    })
    assert.NotPanics(t, func() {
        mustParse("42")
    })
}
```
