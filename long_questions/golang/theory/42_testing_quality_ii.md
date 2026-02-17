# 🧪 Go Theory Questions: 821–840 Testing & Quality II

## 821. How do you structure tests for a large Go codebase?

**Answer:**
1.  **Unit Tests**: Alongside code (`foo_test.go`). White-box.
2.  **Integration Tests**: In `tests/` folder. Black-box (build binary, run against Docker).
3.  **Test Helpers**: `internal/testutil` package (reusable factories).
4.  **Packages**: Use `package foo_test` (external testing) avoids circular deps and forces testing the public API only.

---

## 822. How do you use interfaces for testability?

**Answer:**
We follow the **Consumer Defined Interface** pattern.
Service A needs Database.
Define `type DataStore interface { Save(u User) }`.
In Prod: Inject `PostgresStore`.
In Test: Inject `MockStore`.
This allows us to test Service A's logic without spinning up a real Postgres instance, keeping unit tests sub-millisecond.

---

## 823. How do you test panics in Go?

**Answer:**
We use `defer` + `recover`.
```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    TriggerPanic()
}
```
Helper libraries like `assert.Panics(t, func(){...})` make this cleaner.

---

## 824. How do you generate test data in Go?

**Answer:**
1.  **Faker Libraries**: `gofakeit` or `go-faker` to generate Names, Emails, IPs.
2.  **Factories**: Create helpers `CreateUser(t, db)` that insert a valid user and return the struct.
3.  **Golden Files**: For complex JSON/HTML outputs, store the expected string in `testdata/a.json` and compare file contents.

---

## 825. How do you test concurrent code in Go?

**Answer:**
1.  **Race Detector**: `go test -race`.
2.  **Stress Testing**: Run loop 1000 times in parallel.
```go
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() { defer wg.Done(); DoWork() }()
}
wg.Wait()
```
We assert final state (like Atomic Counter value) is consistent.

---

## 826. How do you mock database interactions in Go?

**Answer:**
1.  **sqlmock**: Mocks the `sql/driver` interface. matches SQL strings string and returns defined rows.
2.  **Repository Interface**: Mock the `GetUser` method, not the SQL. (Preferred).
3.  **Testcontainers**: Spin up a real convenient Postgres Docker container. This is an **Integration Test**, not a mock, but provides higher confidence.

---

## 827. How do you test middleware in a Go web app?

**Answer:**
Middleware is `func(next Handler) Handler`.
Test:
1.  Create a dummy `next` handler that records "I was called".
2.  Wrap it with Middleware.
3.  Call `ServeHTTP`.
4.  Assert: Middleware did its job (Set Header/Auth Check) AND called `next` (or blocked it if 403).

---

## 828. How do you use `httptest.Server`?

**Answer:**
It spins up a real HTTP server on specialized localhost port.
```go
ts := httptest.NewServer(http.HandlerFunc(func(w, r) {
    fmt.Fprintln(w, "Hello")
}))
defer ts.Close()
res, _ := http.Get(ts.URL)
```
This is perfect for testing your **HTTP Client** code to ensure it handles 200, 404, and 500 responses correctly.

---

## 829. How do you run parallel tests in Go?

**Answer:**
`t.Parallel()`.
Must be the first line in the test.
It tells the test runner: "Pause this test, run other non-parallel tests, then run all parallel tests concurrently".
**Gotcha**: Loop variables.
```go
for _, tc := range cases {
    tc := tc // Capture variable
    t.Run(tc.Name, func(t *testing.T) {
        t.Parallel()
        // usage of tc
    })
}
```

---

## 830. How do you test CLI apps in Go?

**Answer:**
We abstract `Stdin`, `Stdout`, `Stderr` in the App struct.
`type App struct { Out io.Writer }`.
In Main: `app.Out = os.Stdout`.
In Test: `app.Out = &bytes.Buffer{}`.
We execute the command and check the Buffer string.
For integration, `os/exec` to run the compiled binary: `exec.Command("./myapp", "start").Output()`.

---

## 831. How do you perform fuzz testing in Go?

**Answer:**
Go 1.18+ Native Fuzzing.
```go
func FuzzParser(f *testing.F) {
    f.Add("initial_seed")
    f.Fuzz(func(t *testing.T, data string) {
        Parse(data) // Should not panic
    })
}
```
The fuzzer generates random mutations of the seed to crash the parser.

---

## 832. How do you simulate network failures in tests?

**Answer:**
1.  **Custom Key Transport**: Close the connection immediately.
2.  **Context Timeout**: `ctx, _ := context.WithTimeout(ctx, 1*time.Nanosecond)`.
3.  **httptest**: Close the server loop while client is reading.
`ts.Config.ErrorLog = log.New(io.Discard, "", 0)` (to hide log noise).
We verify the app handles `io.EOF` or `context deadline exceeded` gracefully.

---

## 833. How do you write integration tests with Docker?

**Answer:**
**Testcontainers-go**.
Programmatic Docker.
1.  `redisC, _ := tests.GenericContainer(...)`
2.  `host, _ := redisC.Host(ctx)`
3.  `RunTests(host)`
4.  `defer redisC.Terminate(ctx)`
This ensures a clean environment for every test run (or suite), immune to "port 6379 already in use" errors on local dev machines.

---

## 834. How do you test gRPC services in Go?

**Answer:**
1.  **bufconn**: In-memory network listener.
    We dial the server via an in-memory pipe (no TCP overhead).
    `dialer := func() (net.Conn, error) { return lis.Dial() }`
2.  **Client Stub**: We use the generated client interface to call methods on the server instance.

---

## 835. How do you set up CI pipelines for testing Go apps?

**Answer:**
**GitHub Actions** / **GitLab CI**.
Stages:
1.  **Lint**: `golangci-lint`.
2.  **Unit**: `go test -short ./...` (Fast).
3.  **Integration**: `docker-compose up -d && go test -tags integration ./...`.
4.  **Build**: `go build`.
We cache `GOCACHE` and `GOMODCACHE` to speed up subsequent runs.

---

## 836. How do you avoid unnecessary allocations?

**Answer:**
1.  **Preallocate**: `make([]int, 0, 1000)` prevents slice growing/copying.
2.  **Reuse**: `sync.Pool` for buffers.
3.  **Strings**: Use `strings.Builder` instead of `+`.
4.  **Slices**: Use `[0:0]` to reset a slice without re-allocating the backing array.

---

## 837. How do you reduce GC pressure in Go apps?

**Answer:**
GC runs when Heap doubles (GOGC=100).
1.  **Fewer Pointers**: Large definition structs with no pointers are scanned instantly (O(1)). Pointers require traversal (O(N)).
2.  **Value Semantics**: Pass small structs by value.
3.  **Off-Heap**: Use `syscall.Mmap` for massive static datasets (GBs) that shouldn't be scanned by GC.

---

## 838. How do you profile heap allocations?

**Answer:**
`go tool pprof -alloc_space http://localhost:6060/debug/pprof/heap`.
**alloc_space**: Total allocated (even if freed). Good for creating optimization.
**inuse_space**: Current memory usage. Good for leaks.
We look for large "flat" numbers in functions that are called frequently (Hot Path).

---

## 839. How do you use escape analysis to optimize code?

**Answer:**
`go build -gcflags="-m"`.
Output: `moved to heap: x`.
If a variable escapes to heap, it causes GC pressure.
If it stays on stack, it's free (popped when function returns).
Common escapes:
- Returning pointer to local var.
- Passing to `interface{}` (`fmt.Println`).
- Slice too large for stack.
We optimize by keeping variables within the function scope or passing them down, not up.

---

## 840. How do you optimize JSON marshaling in Go?

**Answer:**
Reflection is slow.
1.  **Code Generaton**: `easyjson` / `fastjson`. They generate `MarshalJSON` methods using direct buffer writes, avoiding reflection entirely.
2.  **Streaming**: `json.NewEncoder(w).Encode(v)` is better than `json.Marshal` for large responses (writes directly to network).
3.  **Field Selection**: sending `struct { ID, Name }` is faster than `struct { ID, Name, 50_Fields }`.
