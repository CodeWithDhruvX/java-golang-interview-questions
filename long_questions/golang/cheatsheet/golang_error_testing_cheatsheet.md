# Golang Error Handling & Testing Cheatsheet

Modern patterns for handling errors and writing robust tests in Go.

## 🔴 Modern Error Handling (Go 1.13+)

### 1. The Basics
Errors are values. Always check them immediately.
```go
if err := validate(input); err != nil {
    return fmt.Errorf("validation failed: %w", err) // Wrap it!
}
```

### 2. Wrapping Errors (`%w`)
Use `%w` in `fmt.Errorf` to wrap an error, preserving the original chain.
```go
var ErrNotFound = errors.New("record not found")

func getUser(id int) error {
    return fmt.Errorf("db query failed: %w", ErrNotFound)
}
```

### 3. Checking Errors (`errors.Is`)
Checks if a specific error is present anywhere in the chain. Replaces `==`.
```go
err := getUser(42)
if errors.Is(err, ErrNotFound) {
    // Handle "not found" specifically
}
```

### 4. Casting Errors (`errors.As`)
Checks if an error is of a specific **type** and extracts it. Replaces type assertion.
```go
type QueryError struct {
    Query string
    Err   error
}
func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

// Usage
var qErr *QueryError
if errors.As(err, &qErr) {
    fmt.Println("Failed query:", qErr.Query)
}
``` // Corrected closing brace location from previous logic

### 5. Custom Error Types
Define custom struct errors for complex context.
```go
type APIError struct {
    Code    int
    Message string
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API Error %d: %s", e.Code, e.Message)
}
```

---

## 🟢 Testing Patterns

### 1. Table-Driven Tests (The Standard)
Use slice of structs to define test cases. Dry, readable, and easy to extend.
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -1, -2},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### 2. Subtests (`t.Run`)
Allows running specific tests via command line: `go test -run TestAdd/positive`

### 3. Test Helpers (`t.Helper`)
Marks function as a helper so logs report the *caller's* line number, not the helper's.
```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

### 4. Benchmarking
```go
func BenchmarkLogic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Logic()
    }
}
// Run: go test -bench=. -benchmem
```

### 5. Fuzzing (Go 1.18+)
Finds edge cases by generating random inputs.
```go
func FuzzReverse(f *testing.F) {
    f.Add("hello") // Seed corpus
    f.Fuzz(func(t *testing.T, original string) {
        rev := Reverse(original)
        doubleRev := Reverse(rev)
        if original != doubleRev {
            t.Errorf("Before: %q, after: %q", original, doubleRev)
        }
    })
}
```

---

## 🟡 Mocking (Interfaces)

Go has no built-in mocking framework (like Mockito). Use **Interfaces**.

### 1. Define Dependency as Interface
```go
type Database interface {
    GetUser(id int) (*User, error)
}
```

### 2. Create Mock Struct
```go
type MockDB struct {
    GetUserFunc func(int) (*User, error) // Field is a function!
}

func (m *MockDB) GetUser(id int) (*User, error) {
    return m.GetUserFunc(id)
}
```

### 3. Inject Mock in Test
```go
func TestService(t *testing.T) {
    mock := &MockDB{
        GetUserFunc: func(id int) (*User, error) {
            return &User{Name: "MockUser"}, nil
        },
    }
    service := NewService(mock)
    // ... test service ...
}
```
