## 🧪 Testing & Quality (Questions 821-840)

### Question 821: How do you mock HTTP clients in Go tests?

**Answer:**
Inject a `Doer` interface or replace the `Transport` of the client.
Usually, use `httptest.NewServer` to stand up a real fake server and point your client's BaseURL to `ts.URL`.

---

### Question 822: How do you write table-driven tests in Go?

**Answer:**
(See Q541). Struct of cases -> Loop -> Run subtest (`t.Run`).

---

### Question 823: How do you achieve high test coverage in Go?

**Answer:**
Write tests for logic, not just success paths.
Handle errors.
Use `go test -coverprofile` to find gaps.
**Caveat:** 100% coverage doesn't mean bug-free.

---

### Question 824: How do you test race conditions in Go?

**Answer:**
Add `-race` flag. `go test -race ./...`.
It instruments memory accesses and warns if 2 goroutines touch the same var without sync.

---

### Question 825: How do you benchmark functions in Go?

**Answer:**
(See Q521). `func BenchmarkX(b *testing.B)`.

---

### Question 826: How do you structure tests for a large Go codebase?

**Answer:**
- Unit tests next to code.
- `integration/` folder for e2e tests.
- `testdata/` for fixtures.
- Shared `testutil` package for helper functions (creating users, cleaning DB).

---

### Question 827: How do you use interfaces for testability?

**Answer:**
Accept Interfaces, Return Structs.
If a function accepts `DatabaseReader` interface, you can pass `MockDB` in tests without spinning up Postgres.

---

### Question 828: How do you test panics in Go?

**Answer:**
(See Q554). `defer recover()`.

---

### Question 829: How do you generate test data in Go?

**Answer:**
Helper functions.
`func CreateUser(t *testing.T) *User`
Using `faker` libraries helps generate unique emails to avoid constraint violations.

---

### Question 830: How do you test concurrent code in Go?

**Answer:**
Use `sync.WaitGroup` to wait for goroutines.
Use Channels to verify data flow.
Run with `-race`.
Use atomics or mutexes in the test verification logic itself.

---

### Question 831: How do you mock database interactions in Go?

**Answer:**
`go-sqlmock` or `pgxmock` (for pgx driver).
Asserts that "Expected Query X with Args Y was executed".

---

### Question 832: How do you test middleware in a Go web app?

**Answer:**
Pass a dummy handler to the middleware.
Check if the recorder `w` has expected headers/status (e.g., 401 Unauthorized if no token provided).

---

### Question 833: How do you use `httptest.Server`?

**Answer:**
(See Q549). Use it to test code that *makes* HTTP requests.

---

### Question 834: How do you run parallel tests in Go?

**Answer:**
Call `t.Parallel()` at the start of the test function.
Make sure you capture loop variables properly.

---

### Question 835: How do you test CLI apps in Go?

**Answer:**
(See Q558). Capture Stdout/Stderr using `bytes.Buffer`.

---

### Question 836: How do you perform fuzz testing in Go?

**Answer:**
(See Q559). `f.Fuzz(...)`.

---

### Question 837: How do you simulate network failures in tests?

**Answer:**
1.  Close the `httptest.Server` midway.
2.  Use a custom Transport that always returns error.
3.  Use **Toxiproxy**.

---

### Question 838: How do you write integration tests with Docker?

**Answer:**
Use `testcontainers-go`.
It spins up a specialized container (e.g., Redis:7) before tests start, gives you the ephemeral Host:Port, and tears it down after tests.

---

### Question 839: How do you test gRPC services in Go?

**Answer:**
`bufconn` (Buffered Connection).
Dial the server using an in-memory listener (`net.Pipe` or equivalent) instead of TCP.
This allows testing the full gRPC stack without network overhead.

---

### Question 840: How do you set up CI pipelines for testing Go apps?

**Answer:**
GitHub Actions / GitLab CI.
Steps: Checkout -> Setup Go -> Go Mod Download -> Go Vet -> Go Test (with Race) -> Upload Coverage.

---
