# 🟢 Go Theory Questions: 541–560 Testing Deep Dive

## 541. How do you write table-driven tests in Go?

**Answer:**
Table-driven testing is the idiomatic Go way to avoid repetitive test code.
We define a slice of structs, where each struct represents a "Test Case" (Input + Expected Output).

```go
tests := []struct {
    name  string
    input int
    want  int
}{
    {"positive", 2, 4},
    {"negative", -1, 1},
}
```
We iterate this slice (`range tests`) and run `t.Run(tt.name, ...)` for each case. This makes adding a new edge case as simple as adding one line to the struct, without copying logic.

---

## 542. What is the difference between `t.Fatal` and `t.Errorf`?

**Answer:**
`t.Errorf` logs a failure but **continues** executing the rest of the test function. Use this to check multiple independent assertions in one test.
`t.Fatal` logs a failure and **stops** the test immediately (`runtime.Goexit`). Use this when meaningful progress is impossible (e.g., "DB setup failed" or "Context creation failed").

---

## 543. How do you use `go test -cover` to check coverage?

**Answer:**
`go test -cover` gives a simple percentage.
For details: `go test -coverprofile=c.out`.
Then visualize: `go tool cover -html=c.out`.

This opens a browser showing your code in green (covered) and red (not covered).
**Warning**: Coverage is a guide, not a goal. 100% coverage doesn't mean bug-free code; it just means every line was executed once. We aim for ~80% and focus on critical paths.

---

## 544. How do you mock a database in Go tests?

**Answer:**
We prefer **Integration Tests** with a real Dockerized DB instance (`testcontainers-go`) because SQL mocks (`go-sqlmock`) are brittle (they test the *query string*, not the result).

However, for Unit Tests:
We define a `Repository` interface.
In production: `type SQLRepo struct { *sql.DB }`.
In test: `type MockRepo struct { MockUser User }`.
The service calls `repo.GetUser()`. The mock returns the struct immediately. This isolates the service logic from the potentially slow database layer.

---

## 545. How do you unit test HTTP handlers?

**Answer:**
We use `net/http/httptest`.
It provides `NewRecorder()`, which acts as a fake `ResponseWriter`.

```go
req := httptest.NewRequest("GET", "/users", nil)
w := httptest.NewRecorder()
MyHandler(w, req)

resp := w.Result()
// Assert resp.StatusCode == 200
```
This tests the handler logic directly without spinning up a real TCP listener, making tests run in microseconds.

---

## 546. What is testable design and how does Go encourage it?

**Answer:**
Testable design means **Accepting Interfaces**.

If your function takes `func Process(f *os.File)`, you can only test it with real files (slow, messy cleanup).
If you change it to `func Process(r io.Reader)`, you can pass `strings.NewReader("data")` in the test.
Go's implicit interface satisfaction encourages small, single-method interfaces (`io.Reader`, `io.Writer`) that are trivial to mock, naturally leading to testable code.

---

## 547. How do you use interfaces to improve testability?

**Answer:**
Interfaces allow **Dependency Injection**.

Instead of:
`func Login() { db := ConnectParams() ... }` (Hardcoded dependency).
Use:
`func Login(db Database) { ... }`

In tests, we pass a `FakeDatabase` (in-memory map). This decouples the logic from infrastructure. If you can't instantiate your component without a running AWS environment, your design is not testable.

---

## 548. How do you write tests for concurrent code in Go?

**Answer:**
We use the **Race Detector** (`go test -race`).
We also use `sync.WaitGroup` or channels to ensure the test waits for goroutines to finish.

Common pattern: "Stress Test".
Launch 100 goroutines calling the function.
Assert that the internal state (e.g., a counter) matches expected (100).
Without `-race`, you might miss data races that only crash 1 in 1000 times.

---

## 549. What is the `httptest` package and how is it used?

**Answer:**
It is part of the standard library.
1.  `ResponseRecorder`: Mocks the `ResponseWriter`.
2.  `NewServer`: Starts a real HTTP server on a random local port.

`ts := httptest.NewServer(http.HandlerFunc(func(...) { ... }))`
`defer ts.Close()`
`client.Get(ts.URL)`
This is used for **Integration Testing** an HTTP Client. instead of mocking the network, you spin up a real "Server" that returns canned responses.

---

## 550. How do you mock time in tests?

**Answer:**
Testing code that says `if time.Now().After(expiry)` is hard because `time.Now()` moves.

We define an interface `Clock`:
```go
type Clock interface { Now() time.Time }
type RealClock struct{} 
```
In tests, we inject `FakeClock`. We can manually "Advance" time by 1 hour to verify that the token expiry logic triggers exactly when it should, without actually sleeping for an hour.

---

## 551. How do you perform integration testing in Go?

**Answer:**
Integration tests live in the same package (or `_test` package) but use build tags: `//go:build integration`.

They are slow. They connect to real Redis/Postgres.
Run: `go test -tags=integration ./...`.
We use **Testcontainers** to spin up ephemeral Docker containers (Redis, Postgres) for the duration of the test, ensuring a clean state every time.

---

## 552. How do you use `testify/mock` for mocking dependencies?

**Answer:**
`testify/mock` provides a mechanism to record expectations.

```go
m := new(MyMock)
m.On("Calculate", 5).Return(25)
result := service.DoThing(m)
m.AssertExpectations(t)
```
This verifies not just the output, but **Behavior**: "Did the service call Calculate exactly once with argument 5?". This is strict mocking.

---

## 553. How do you run subtests and benchmarks?

**Answer:**
Subtests: `t.Run("SubName", func(t *testing.T) { ... })`. Allows grouping tests in a hierarchical output.

Benchmarks: `func BenchmarkX(b *testing.B)`.
Run specific bench: `go test -bench=BenchmarkX`.
Compare benchmarks: **benchstat**. It compares "Before" and "After" results to tell you if the difference is statistically significant or just noise.

---

## 554. How do you test panic recovery?

**Answer:**
We verify that a function panics using `defer`.

```go
defer func() {
    if r := recover(); r == nil {
        t.Errorf("Expected panic, but code terminated normally")
    }
}()
TriggerPanic()
```
Libraries like `testify/assert` simplify this: `assert.Panics(t, func(){ ... })`.

---

## 555. How do you generate test data using faker or random data?

**Answer:**
We use `gofakeit` or `go-faker`.

Hardcoding "John Doe" 100 times leads to "Pesticide Paradox" (tests stop finding bugs).
Faker generates random functional data: emails, addresses, credit cards.
`u := User{Email: faker.Email(), Name: faker.Name()}`.
This ensures our validation logic handles variety (long names, weird characters) correctly.

---

## 556. What is golden file testing and when is it useful?

**Answer:**
Golden Files are for **Complex Output** (JSON, HTML, CSV).
Instead of asserting `s == "{...very long string...}"`.
We write the output to `testdata/output.golden`.

Current run: Compare `actual` vs `file(testdata/output.golden)`.
If they differ, fail.
If the change is intentional, run `go test -update` to overwrite the golden file with the new output. Ideally suited for CLI tools or HTML templates.

---

## 557. How do you automate test workflows with `go generate`?

**Answer:**
`go generate` runs commands embedded in comments.
`//go:generate mockgen -source=interface.go -destination=mock.go`

We run `go generate ./...` before `go test`.
This ensures that all our Mocks, Protobufs, and Stringers are up-to-date. We never edit generated files manually.

---

## 558. How do you test CLI apps built with Cobra?

**Answer:**
Cobra commands are just structs with a `RunE` function.
We manually construct the command in the test.

```go
cmd := NewRootCmd()
b := bytes.NewBufferString("")
cmd.SetOut(b)
cmd.SetArgs([]string{"create", "--dry-run"})
err := cmd.Execute()
```
We assert `err` is nil and `b.String()` contains "Dry run successful". This allows full CLI testing without spawning OS subprocesses.

---

## 559. What is fuzz testing and how do you do it in Go?

**Answer:**
(See Q 497 - Fuzzing).
Fuzzing sends random inputs to find edge cases.
`f.Fuzz(func(t *testing.T, a string, b int) { ... })`.
The tooling automatically generates the corpus.
We use it for **Parsers**, **Deserializers**, and **Crypto** code—anything that takes untrusted byte slices from the outside world.

---

## 560. How do you organize test files and test suites?

**Answer:**
1.  **Unit Tests**: In `package foo`, file `foo_test.go`. Access to private internals.
2.  **Black Box Tests**: In `package foo_test`. `import "foo"`. Only access public API. Forces you to use the library as a user would.
3.  **Testdata**: Directory named `testdata` is ignored by the compiler. Store JSON/Golden files there.
4.  **MainWrapper**: `TestMain(m *testing.M)` for global setup/teardown (like initializing a Docker container for the whole suite).
