## 🧪 Testing in Go (Questions 541-560)

### Question 541: How do you write table-driven tests in Go?

**Answer:**
Define a struct for test cases (input/expected), slice of that struct, and loop.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {-1, 1, 0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

---

### Question 542: What is the difference between `t.Fatal` and `t.Errorf`?

**Answer:**
- **`t.Errorf`:** Log error and **continue** executing the test. Use when testing multiple independent assertions in a single test function.
- **`t.Fatal` (or `t.Fatalf`):** Log error and **stop** the test immediately. Use when a setup step fails (e.g., db connection) making further assertions impossible.

---

### Question 543: How do you use `go test -cover` to check coverage?

**Answer:**
- **Basic:** `go test -cover` gives a % summary.
- **Profile:** `go test -coverprofile=c.out` saves detailed data.
- **View:** `go tool cover -html=c.out` opens a web browser to visually see which lines are red (uncovered) vs green.

---

### Question 544: How do you mock a database in Go tests?

**Answer:**
1.  **DATA-DOG/go-sqlmock:** Mocks the `sql/driver` layer. Tests your SQL query syntax/arguments without a real DB.
2.  **Docker (Testcontainers):** Spin up a real Postgres for Integration Tests.
3.  **Interface:** Mock the `Repository` interface (application layer) rather than the database driver.

---

### Question 545: How do you unit test HTTP handlers?

**Answer:**
Use `httptest.NewRecorder()` to capture the response.

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    
    MyHandler(w, req)
    
    resp := w.Result()
    if resp.StatusCode != 200 { t.Fail() }
}
```

---

### Question 546: What is testable design and how does Go encourage it?

**Answer:**
Testable design favors **Dependency Injection** via Interfaces.
Go encourages this by having implicit interfaces. You define `type Reader interface` where you use it, making it easy to swap a real `os.File` with an in-memory `bytes.Buffer` for testing.

---

### Question 547: How do you use interfaces to improve testability?

**Answer:**
Instead of accepting a concrete type `*Service`, accept an interface `ServiceAPI`.

```go
type Downloader interface { Download(url string) ([]byte, error) }

func Process(d Downloader) { ... }

// Test
type MockDownloader struct{}
func (m MockDownloader) Download(u string) ([]byte, error) { return []byte("fake"), nil }
// Now you can test Process() without internet.
```

---

### Question 548: How do you write tests for concurrent code in Go?

**Answer:**
1.  **Race Detector:** Always run with `go test -race`.
2.  **Stress Testing:** Run multiple iterations or use parallel subtests.
3.  **Synchronization:** Use `WaitGroup` or channels in the test to ensure all goroutines complete before checking assertions.

---

### Question 549: What is the `httptest` package and how is it used?

**Answer:**
It provides utilities for HTTP testing.
- `ResponseRecorder`: Records writes to `ResponseWriter`.
- `Server`: Starts a real (local) HTTP server on a random port for end-to-end testing of HTTP Clients.
    ```go
    ts := httptest.NewServer(http.HandlerFunc(func(w, r) { fmt.Fprintln(w, "Hello") }))
    defer ts.Close()
    client.Get(ts.URL)
    ```

---

### Question 550: How do you mock time in tests?

**Answer:**
Avoid `time.Now()`. Inject a Clock interface.

```go
type Clock interface { Now() time.Time }
type App struct { Clock Clock }

func (a *App) IsWeekEnd() bool {
    return a.Clock.Now().Weekday() == time.Sunday
}
```
In tests, pass a `MockClock` that returns a fixed static time.

---

### Question 551: How do you perform integration testing in Go?

**Answer:**
Integration tests interact with external systems (DB, Redis).
1.  Use build tags `// +build integration` to separate them from unit tests.
2.  Run with `go test -tags=integration ./...`.
3.  Use `testcontainers-go` to spin up dependencies in Docker automatically.

---

### Question 552: How do you use `testify/mock` for mocking dependencies?

**Answer:**
`testify` is a popular assertion/mocking library.

```go
type MyMock struct { mock.Mock }
func (m *MyMock) Do(n int) int {
    args := m.Called(n)
    return args.Int(0)
}

// Test
m := new(MyMock)
m.On("Do", 5).Return(10)
result := m.Do(5)
m.AssertExpectations(t)
```

---

### Question 553: How do you run subtests and benchmarks?

**Answer:**
Use `t.Run()` or `b.Run()`.

```go
func TestMath(t *testing.T) {
    t.Run("Add", func(t *testing.T) { ... })
    t.Run("Sub", func(t *testing.T) { ... })
}
```
Run specific subtest: `go test -run TestMath/Add`.

---

### Question 554: How do you test panic recovery?

**Answer:**
Use `defer` inside the test to catch the panic.

```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    CausePanic()
}
```
*(Or use `assert.Panics` from testify)*

---

### Question 555: How do you generate test data using faker or random data?

**Answer:**
Use libraries like `github.com/go-faker/faker` or `github.com/brianvoe/gofakeit`.
This generates random specific data (Emails, Names, Addresses) to test edge cases or validation logic.

```go
var u User
faker.FakeData(&u) // Fills struct with random valid data
```

---

### Question 556: What is golden file testing and when is it useful?

**Answer:**
Useful for comparing large outputs (JSON, HTML).
1.  Run test.
2.  If `update` flag is set, write output to standard file (`testdata/output.golden`).
3.  If not set, read `output.golden` and `bytes.Equal(got, expected)`.
Ensures complex output remains stable.

---

### Question 557: How do you automate test workflows with `go generate`?

**Answer:**
`go generate` runs commands embedded in comments.
Use it to generate mocks before testing.
`//go:generate mockery --name=UserService`
Run: `go generate ./...` -> `go test ./...`

---

### Question 558: How do you test CLI apps built with Cobra?

**Answer:**
Set custom input/output buffers on the Root Command.

```go
buf := new(bytes.Buffer)
rootCmd.SetOut(buf)
rootCmd.SetArgs([]string{"create", "--name", "test"})
rootCmd.Execute()

if !strings.Contains(buf.String(), "Created") { t.Fail() }
```

---

### Question 559: What is fuzz testing and how do you do it in Go?

**Answer:**
Go 1.18+ Fuzzing. Uses a genetic algorithm to mutate inputs to trigger crashes.

```go
func FuzzReverse(f *testing.F) {
    f.Fuzz(func(t *testing.T, s string) {
        // Assert property: Reverse(Reverse(s)) == s
        if Reverse(Reverse(s)) != s { t.Error("Failed") }
    })
}
```

---

### Question 560: How do you organize test files and test suites?

**Answer:**
- **Unit Tests:** Inside the same package (`package foo`), file `foo_test.go`. Access to private members.
- **Blackbox Tests:** Use `package foo_test`. Import `foo`. Tests only public API. (Preferred for integration/API tests).
- **Testdata:** Store fixtures in a `testdata/` directory (ignored by compiler).
- **Suite:** Use `testify/suite` for Setup/Teardown logic shared across tests.

---
