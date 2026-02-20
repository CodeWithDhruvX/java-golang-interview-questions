# 🔴 **281–300: Testing Strategy, CI/CD, Observability**

### 281. What are test doubles and how are they used in Go?
"A **Test Double** is the generic term for any simulated object used in testing.
Since Go doesn't have a built-in mocking framework like Java (Mockito) or Python (unittest.mock), we use interfaces.

*   **Dummy**: Passed around but never used.
*   **Stub**: Returns canned answers (`return true`).
*   **Mock**: Expects specific calls to verify behavior (`Expect(SendEmail).Times(1)`).
I use **mockery** to generate these doubles automatically from my interfaces."

#### Indepth
Mocks are often overused. If you mock the `database/sql` driver to test a repository, you are testing your mock, not the SQL. A **Fake** (in-memory implementation of the interface, e.g., using a map) is often superior to a Mock because it has working logic and behaves more like the real thing.

---

### 282. How do you structure unit vs integration tests?
"**Unit Tests**: Live in the same package (e.g., `user_test.go` next to `user.go`). They test individual functions in isolation. They must be fast (milliseconds).

**Integration Tests**: Live in a separate `tests/` folder or have a `_test` package name. They spin up real dependencies (Postgres, Redis) via Docker.
I use build tags (`//go:build integration`) so I can run `go test ./...` quickly without triggering the slow integration suite."

#### Indepth
Integration tests should be **Hermetic**. They should start their own dependencies (Postgres, Redis) in transient containers (using `testcontainers-go`). Sharing a "Dev DB" across developers or CI runs leads to flaky tests when two tests modify the same row simultaneously.

---

### 283. What are flaky tests and how do you identify them?
"A flaky test passes sometimes and fails others. They are the enemy of trust in CI.

Common causes: `time.Sleep()` (race conditions), map iteration order, or shared global state.
I identify them by stress-testing: `go test -count=100 -race`. If it fails even once in 100 runs, it's a bug. I treat flaky tests as P0 issues because they block deployments."

#### Indepth
When debugging flakes, use `go test -race -count=1000 -failfast`. The `-failfast` flag stops execution on the first failure, saving you time. Also, check for map iteration order dependency—this is the #1 cause of randomness in Go tests.

---

### 284. How do you write deterministic tests for concurrency?
"**Never use `time.Sleep`.** It’s non-deterministic because the OS scheduler is unpredictable.

Instead, I use synchronization primitives.
Channels are great for signaling ('Done'). `sync.WaitGroup` is great for waiting.
For time-dependent code, I inject a `Clock` interface so I can manually advance time (`clock.Add(1 * time.Hour)`) instantly in the test."

#### Indepth
For channels, a common mistake is not closing them, leading to deadlocks in tests. A useful pattern is using `context.WithTimeout` in your test: `select { case val := <-ch: assert(val); case <-ctx.Done(): t.Fatal("timeout") }`. This prevents the test suite from hanging forever if a goroutine gets stuck.

---

### 285. How do you test RESTful APIs in Go?
"I use the `net/http/httptest` package.

It allows me to spin up my handler without binding to a real TCP port.
`req := httptest.NewRequest("GET", "/users", nil)`
`w := httptest.NewRecorder()`
`handler.ServeHTTP(w, req)`
Then I assert `w.Code == 200` and check `w.Body`. It’s incredibly fast and tests the exact routing logic."

#### Indepth
`httptest` bypasses the network stack (TCP/IP). It directly calls the `ServeHTTP` method. This means it won't catch issues like "connection reset", "timeouts", or "invalid TLS certs". For those, you need a true end-to-end test that spins up a server on `localhost`.

---

### 286. How do you mock HTTP calls?
"I avoid making real network calls in unit tests.

I create an `httptest.Server`.
It gives me a local URL (e.g., `http://127.0.0.1:45321`). I configure my API client to talk to *this* URL instead of the real Stripe/AWS API.
Inside the test server, I write logic to return specific JSON responses or error codes. This makes my tests deterministic and capable of running offline."

#### Indepth
Be careful with **Global State** in `http.DefaultClient`. If your code uses the default client, and your test replaces its Transport with a mock, you might break other tests running in parallel. Always inject the `HTTPClient` interface into your service struct.

---

### 287. What is Golden Testing in Go?
"**Golden files** are saved files containing the *expected* output.

It's perfect for complex output like generated HTML, CLI text, or large JSON blobs.
Instead of hardcoding the expected string in the test file (which is messy), I read from `testdata/output.golden`.
If the output changes intentionally, I run `go test -update` to overwrite the file. It turns a painful string comparison into a simple file diff."

#### Indepth
Golden files are dangerous if you don't manually inspect the diffs. Developers often "blindly update" golden files to make the CI pass, accidentally baking in a bug. Treat a change in a golden file with the same scrutiny as a code change.

---

### 288. How do you run tests in parallel?
"I call `t.Parallel()` at the start of the test function.

This tells the test runner: 'Pause this test, go start the others, and run me alongside them.'
It’s crucial for speeding up I/O bound tests.
However, I have to be careful **not** to use shared global variables or tables in parallel tests, otherwise, I will get race conditions *inside* my test suite."

#### Indepth
`t.Parallel()` creates a new goroutine. If you are using Table Driven Tests, you must re-capture the loop variable: `tc := tc` inside the loop (before calling `t.Run`). Although Go 1.22 fixed the loop variable scope, older codebases (and muscle memory) still rely on this manual shadowing.

---

### 289. How do you mock time-dependent code?
"I never use `time.Now()` directly in business logic.

I define a interface:
`type Clock interface { Now() time.Time }`.
In production, I use a real implementation. In tests, I inject a `MockClock`.
This lets me test logic like 'token expires in 24 hours' by simply moving the mock clock forward, without actually waiting."

#### Indepth
Use `jonboulle/clockwork` or `benbjohnson/clock`. These libraries provide a thread-safe Mock Clock. Writing your own is easy but implementing `After()` and `Ticker()` correctly in a mock is surprisingly tricky (handling deadlock scenarios).

---

### 290. How do you simulate DB failures in tests?
"I use a mock driver like **go-sqlmock**.

I can program exact expectations:
`mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)`.
This allows me to verify that my application handles database outages gracefully (e.g., retrying or returning a 503) without needing to actually unplug a database cable."

#### Indepth
`go-sqlmock` matches SQL requires strict string matching by default (regex). This makes tests fragile—changing `SELECT a, b` to `SELECT a,b` (space removal) might break the test. Use `mock.ExpectQuery(regexp.QuoteMeta(sql))` if you want exact matching, or lenient regexes if you want robustness.

---

### 291. How do you use GitHub Actions to test Go apps?
"I define a workflow `.github/workflows/test.yml`.

It usually has three steps:
1.  **Lint**: `golangci-lint run`.
2.  **Test**: `go test -race -cover ./...`.
3.  **Build**: `go build`.
I use `actions/cache` to cache the Go module download steps. This gets my CI feedback time down to under 2 minutes."

#### Indepth
Linting should always run **before** testing. Detecting a syntax error or a shadowed variable via static analysis is instant. Waiting 5 minutes for tests to compile and run only to fail on a formatting issue is a waste of compute credits and developer focus.

---

### 292. What is the structure of a Makefile for Go?
"A Makefile standardizes the developer experience.

Common targets:
*   `test`: Runs unit tests with race detection.
*   `build`: Compiles the binary.
*   `docker`: Builds the container.
*   `lint`: Runs the linter.
This means a new developer just types `make test` and doesn't need to know the specific flags we use."

#### Indepth
Makefiles are great, but they are not cross-platform (Windows developers often lack `make`). A modern alternative is **Taskfile** (Task), which uses a YAML schema and runs natively on all OSes. It is gaining traction in the Go community.

---

### 293. How to build and test Go code in Docker?
"I use a multi-stage Dockerfile where the first stage runs the tests.

`RUN go test ./...` inside the build stage.
If the tests fail, the `docker build` command fails.
This guarantees that the artifact I'm deploying was built from code that passed tests in the *exact same environment*, eliminating 'it works on my machine' issues."

#### Indepth
Be aware of **CGO** differences. If you test on an Alpine container (musl libc) but deploy to a Debian/Distroless container (glibc), unexpected bugs can occur. Ensure your Test Stage and Release Stage use the same base OS family or use `CGO_ENABLED=0` static binaries.

---

### 294. What CI tools are commonly used for Go projects?
"**GitHub Actions** is the standard for open source and many startups.
**GitLab CI** is popular in enterprise.
**CircleCI** is known for speed.

Regardless of the tool, the pipeline is the same: Checkout -> Lint -> Test -> Build -> Push Docker Image.
Go’s tooling (fmt, vet, test) is so standard that it works seamlessly in all of them."

#### Indepth
The only "CI-specific" config needed is usually **Module Caching**. Go modules are immutable. Once downloaded, `ver 1.0.0` never changes. Aggressively cache the `$GOPATH/pkg/mod` directory. If the `go.sum` hasn't changed, the restore is free.

---

### 295. What are the benefits of go:embed for test fixtures?
"Before `go:embed`, reading `testdata/sample.json` was painful because relative paths broke depending on where I ran the test from.

Now, I use `//go:embed testdata/*.json` and capture it in a `embed.FS`.
It’s a virtual file system baked into the test binary. It makes tests completely portable and self-contained, which is great for remote execution."

#### Indepth
`embed` is read-only. You cannot write back to the embedded files during the test. If your test needs to modify a fixture (e.g., "load config, change port, save it"), you must `io.Copy` the embedded file to a temp directory (`os.MkdirTemp`) first.

---

### 296. How do you generate coverage reports in HTML?
"It’s a standard two-step command.

1.  `go test -coverprofile=c.out ./...` (Collect data).
2.  `go tool cover -html=c.out` (Visualize).
This opens a browser where I can see exactly which lines are green (covered) and red (missed). I use this during code review to spot gaps in testing."

#### Indepth
Watch out for "Table Driven Test" traps in coverage. If you have one test function that loops over 50 cases, the coverage tool highlights the function body as "covered" if *at least one* case ran. It doesn't prove that *all* edge cases were hit. Logic coverage != Data coverage.

---

### 297. How to collect logs and metrics from Go services?
"I follow the **Push vs Pull** model.

**Logs (Push)**: The app writes JSON to `stdout`. Using a log shipper (Fluentd), we push these to Elasticsearch.
**Metrics (Pull)**: The app exposes a `/metrics` HTTP endpoint. Prometheus scrapes (pulls) this endpoint every 15s.
This separation ensures that if my logging backend dies, it doesn't crash my application."

#### Indepth
Prometheus is **Cumulative**. Counter metrics (like `requests_total`) only go up. If your app crashes and restarts, the counter resets to 0. Prometheus handles this "counter reset" automatically in `rate()` calculations, but be aware that the raw number is meaningless without the rate.

---

### 298. What is structured logging in Go?
"Structured logging is treating logs as data (JSON), not text.

Instead of `log.Printf("Failed: %v", err)`, I use:
`logger.Error("failed", "error", err, "user_id", u.ID)`.
This produces `{"msg":"failed", "error":"timeout", "user_id":123}`.
This allows me to query logs like a database: `show logs where user_id = 123`."

#### Indepth
Structued logs are huge. A single log entry can be 1KB JSON. If you log inside a tight loop, you will saturate I/O. Use **Log Levels** dynamically. In Prod, set level to `INFO`. In Dev, set to `DEBUG`. Some advanced frameworks allow changing the log level at runtime via an HTTP call without restarting the app.

---

### 299. What are common logging libraries in Go?
"Go 1.21 introduced **slog**, which is now the standard built-in structured logger. I use it for all new projects.

Historically, **Zap** (by Uber) was the performance king (zero allocation). **Logrus** was the usability king but is now slow.
I generally recommend `slog` because it has zero external dependencies."

#### Indepth
`slog` is an interface. You can swap the "Handler". You can verify one handler that writes JSON to stdout, and another that sends data to a channel for processing. This pluggability makes it compatible with legacy logging backends seamlessly.

---

### 300. How do you aggregate and search logs across services?
"I use a centralized logging stack like **ELK** (Elasticsearch) or **Loki** (Grafana).

My Go services simply write to `stdout`.
The infrastructure (Docker/Kubernetes) captures these streams and forwards them to the central server.
Using a correlation ID (Trace ID) in the logs allows me to trace a single request across 10 different microservices."

#### Indepth
Don't use `grep` on a server for logs. It doesn't scale. If you can't afford ELK/Splunk, use **Loki**. It's designed like Prometheus (index only labels, not the content), making it extremely cheap to operate while still allowing `grep`-like queries over terabytes of logs.
