# 🧪 **821–840: Testing & Quality (Part 2)**

### 821. How do you implement contract testing in Go?
"I use **Pact**.
Consumer (Frontend) defines the contract: 'I expect GET /user to return {id, name}'.
Provider (Go API) verifies it fulfills the contract.
This prevents breaking changes between microservices without needing full End-to-End integration environments."

#### Indepth
**Pact Broker**. The real power of Pact comes from the **Broker**. It's a central registry where Consumers upload "Pacts" and Providers upload verification results. Before deploying Consumer v2 to production, the pipeline asks the Broker: "Is there a Provider version in prod that satisfies my new contract?". If no, deployment is blocked.

---

### 822. How do you run tests in parallel to speed up CI?
"`go test -p 8 ./...`.
Inside tests: `t.Parallel()`.
This allows one package's tests to run on multiple cores.
Constraint: My tests *must* be isolated (no shared global DB). If they rely on DB, I need to spin up a DB per package or use transactions that rollback."

#### Indepth
**t.Cleanup()**. When running parallel tests, `defer cleanup()` might run *after* the test function returns but *before* the parallel subtests finish (if using `t.Run`). Use `t.Cleanup(func() { ... })` instead of `defer`. `t.Cleanup` guarantees execution *after* the test and all its subtests are complete.

---

### 823. How do you manage test data in Go?
"I use **Fixtures** or **Factories**.
Factory: `NewUser(func(u) { u.Role = "admin" })`.
I avoid sharing huge JSON dumps between tests.
I create helper functions `createTestUser(t, db)` that insert minimal required data and return the object."

#### Indepth
**Go-CMP**. `reflect.DeepEqual` is brittle (it distinguishes `nil` slice from empty slice `[]int{}`). Use google's `go-cmp`. It allows ignoring unexported fields, approximating float comparisons, and sorting slices before comparing (`cmpopts.SortSlices`). It provides a readable diff of *exactly* what mismatched.

---

### 824. How do you verify log output in tests?
"I inject a custom `io.Writer` into the Logger.
In tests, I pass a `*bytes.Buffer`.
`logger := NewLogger(buf)`.
`if !strings.Contains(buf.String(), "error connecting") { t.Fail() }`.
This asserts that the expected warning/error was actually logged."

#### Indepth
**Testable Examples**. Go supports `func ExampleFoo()` in `_test.go` files. If you add a comment `// Output: hello`, `go test` runs the code and asserts stdout matches the comment. This verifies your documentation examples *are* actual running tests and never get out of date.

---

### 825. How do you test Go code that depends on `time.Now()`?
"I inject a Clock interface.
`type Clock interface { Now() time.Time }`.
Real implementation: `return time.Now()`.
Test implementation: `return fixedTime`.
This allows me to test logic like 'token expires in 1 hour' deterministically."

#### Indepth
**Time Travel**. If you use a `Clock` interface, you can simulate long durations instantly. `fakeClock.Add(24 * time.Hour)`. A test that verifies "Batch job runs after 24 hours" runs in milliseconds, not 24 hours. This is essential for testing timeouts, cache expirations, and cron jobs.

---

### 826. How do you use `httptest.ResponseRecorder` effectively?
"It implements `http.ResponseWriter`.
I pass it to my handler.
`handler.ServeHTTP(rec, req)`.
I inspect `rec.Code`, `rec.Body.String()`.
It’s an in-memory mock of a real browser connection."

#### Indepth
**Streaming Responses**. `httptest.NewRecorder` buffers the *entire* response in memory. If you are testing a streaming handler (Server-Sent Events or large file download), `Recorder` won't show the "flushing" behavior. You need to use `httptest.NewServer` and a real HTTP client to verify that data is arriving in chunks.

---

### 827. How do you test a graceful shutdown using signals?
"I send a signal to the process in my test.
`proc, _ := os.StartProcess(...)`.
`proc.Signal(syscall.SIGTERM)`.
I verify the process exits with code 0.
Internal unit test: I expose the `shutdown` channel in my `App` struct so the test can trigger it without yielding to the OS."

#### Indepth
**Context Propagation**. When shutting down, you typically give a 5-10 second timeout. In tests, reduce this to 10ms. Pass a `context.WithTimeout` to your Shutdown method. If the shutdown logic (closing DB, flushing logs) respects the context, the test will fail fast if something hangs, rather than hanging forever.

---

### 828. How do you check for goroutine leaks in tests?
"I use `goleak`.
`func TestMain(m *testing.M) { goleak.VerifyTestMain(m) }`.
It checks if any extra goroutines are running after the test finishes compared to before.
It catches `go func() { ... }` that never return."

#### Indepth
**Main Thread**. `goleak` ignores the main goroutine. It focuses on *background* workers. If you start a worker in `init()` (bad practice), `goleak` might assume it's a global system routine. Explicitly ignore known globals with `goleak.IgnoreTopFunction("my/pkg.backgroundWorker")` if they are intended to live forever.

---

### 829. How do you perform load testing on Go servers?
"I use **K6** or **Vegeta**.
They are external tools.
In Go, `httptest.NewServer` is okay for micro-benchmarks, but for rigorous load testing, I deploy the app and hit it from an external agent to measure throughput (RPS) and latency (p99) properly."

#### Indepth
**Scenarios**. Latency usually degrades not linearly, but drastically after a "cliff" (e.g., when DB connection pool fills up). Testing "100 RPS" is useless. Test a *Step Pattern*: 10 RPS -> 100 RPS -> 1000 RPS. Watch for the "Knee" of the curve where latency spikes. That is your capacity limit.

---

### 830. How do you use Go's `testing/quick` package?
"It’s QuickCheck for Go.
`quick.Check(func(x, y int) bool { return Add(x, y) == Add(y, x) }, nil)`.
It generates random `int`s and runs the function.
Deprecated now in favor of Fuzzing (`go test -fuzz`), which is smarter and native."

#### Indepth
**Seed Corpus**. Fuzzing starts random. But if you have a complex format (PDF), random bytes will just fail "Invalid Header" 99% of the time, exercising nothing. Provide a **Seed Corpus** (`f.Add([]byte("%PDF-1.4..."))`). The fuzzer mutates valid inputs to find edge cases deep in the parser logic.

---

### 831. How do you assert JSON responses in Go tests?
"I verify not to string compare! Formatting differs.
`require.JSONEq(t, expected, actual)` (using `testify`).
Or Unmarshal both into structs and compare structs.
`assert.Equal(t, expectedStruct, actualStruct)`.
This ignores whitespace differences."

#### Indepth
**Golden Files**. For complex JSON responses (50+ lines), defining `expectedStruct` in Go code is tedious and unreadable. Use **Golden Files**. Save `testdata/user_response.golden.json`. In test: `actual := handler()`; `if update { os.WriteFile("...golden", actual) }`; `expected := os.ReadFile(...)`. This makes updating tests trivial (run with `-update`).

---

### 832. How do you test code that uses `os.Exit`?
"You can't test `os.Exit` directly (it kills the runner).
Refactor: `func Run() int` (returns exit code).
Main calls `os.Exit(Run())`.
Test calls `Run()` and asserts it returns 1 or 0.
If needed, launch a subprocess (`exec.Command`) to test the actual crash."

#### Indepth
**Coverage**. Tests that run `exec.Command("go", "run", ...)` do *not* count towards code coverage of the main test run (it's a separate process). If you need coverage for the crash logic, you must compile a test binary that includes coverage instrumentation (`go test -c -cover`) and run that.

---

### 833. How do you structure table-driven tests for readability?
"Use named fields in the struct.
`tests := []struct { name string; input int; want int }{ ... }`.
`t.Run(tt.name, ...)` ensures the output shows 'TestAdd/negative_numbers'.
I keep the test body minimal, delegating setup/teardown to helpers."

#### Indepth
**Parallel Subtests**. A common trap: `for _, tt := range tests { t.Run(tt.name, func(t *testing.T) { t.Parallel(); do(tt.input) }) }`. This creates a race condition on `tt` (loop variable). In Go < 1.22, you MUST add `tt := tt` inside the loop. In Go 1.22+, loop variables are fixed, but be aware of older codebases.

---

### 834. How do you test middlewares in isolation?
"A middleware is `func(next) -> handler`.
I create a dummy 'next' handler.
`wrapped := MyMiddleware(dummyNext)`.
`wrapped.ServeHTTP(rec, req)`.
I verify that `MyMiddleware` did its job (set a header, logged) before or after calling `next`."

#### Indepth
**Context Keys**. Middleware often sets context values (`ctx = context.WithValue(ctx, UserKey, user)`). To test this, your dummy handler should check the context: `next = http.HandlerFunc(func(w, r) { if r.Context().Value(UserKey) == nil { t.Error("User not in context") } })`.

---

### 835. How do you mock file system operations in Go?
"I use `spf13/afero`.
It provides an `Fs` interface.
`var AppFs = afero.NewOsFs()`.
In tests: `AppFs = afero.NewMemMapFs()`.
I can create files in memory without touching the disk."

#### Indepth
**Testcontainers**. Mocks (`NewMemMapFs`, `sqlmock`) simulate behavior, but sometimes they simulate *wrongly*. "It works in mock but fails on real S3". Use **Testcontainers-go**. It spins up a *real* MinIO/Postgres container in Docker for the test. It's slower but gives 100% confidence that your SQL syntax/S3 API usage is correct.

---

### 836. How do you perform end-to-end (E2E) testing in Go?
"I build the binary. I spin up docker-compose (DB + App).
My test code uses an HTTP Client to hit the real running API.
I assert the state of the *Database* changed correctly.
This proves the entire stack works together."

#### Indepth
**Ephemeral Environments**. Running E2E tests on a shared "Staging" env is flaky (data collisions). Best practice: Spin up a full *Ephemeral Environment* (namespace in K8s) per PR. Run E2E tests against that isolated URL. Tear it down after. This ensures tests never fail because "Developer Bob deleted the user I was testing with".

---

### 837. How do you test GraphQL resolvers in Go?
"I test the Resolver methods directly.
`r := &Resolver{}`.
`resp, err := r.User(ctx, args)`.
I don't need to go through the HTTP transport. I trust the library (gqlgen) to call my resolver correctly; I only verify my resolver logic."

#### Indepth
**Schema Validation**. While you test resolvers generally, you should also have *one* integration test that sends a real GraphQL query string. This catches issues where the Schema (`schema.graphql`) doesn't match the Resolver implementation (e.g., you return an `int` but schema says `String`), which compile-time checks might miss in dynamic frameworks.

---

### 838. How do you use the `testdata` directory?
"Go ignores `testdata` during compilation.
I put huge JSONs, certificates, and CSVs there.
`data, _ := os.ReadFile("testdata/payload.json")`.
It keeps my test code clean."

#### Indepth
**go:embed**. Instead of `os.ReadFile` (which depends on the CWD being correct when running the test), use `//go:embed testdata/payload.json`. `var payload []byte`. This compiles the test data *into* the test binary. You can run the test binary from anywhere (`/tmp`) and it still has access to its data.

---

### 839. How do you test complex regex patterns?
"I start with unit tests.
`matches := myRegex.FindString(input)`.
I verify edge cases (empty string, huge string, unicode).
For security, I verify that the regex doesn't have ReDoS vulnerabilities by fuzzing requests with long repeating characters."

#### Indepth
**ReDoS**. Regular Expression Denial of Service. Patterns like `(a+)+` have exponential backtracking complexity. Go's `regexp` engine (RE2) guarantees linear time O(n) execution, so it is **immune** to ReDoS (unlike Python or Java). However, it is slightly slower and doesn't support features like backreferences.

---

### 840. How do you test channels and select statements?
"To test blocking behavior, I use timeouts.
`select { case val := <-ch: assert(val); case <-time.After(100*time.Millisecond): t.Fatal("timeout") }`.
To test non-blocking:
`select { case ch <- val: ; default: t.Fatal("should not block") }`."

#### Indepth
**Closed Channels**. A common bug is reading from a closed channel (which returns zero-value immediately) thinking it's valid data. In tests, use the two-value receive: `val, ok := <-ch`. Assert `ok` is true. If `ok` is false, it means the channel was closed unexpectedly, and your test logic might be flawed.
