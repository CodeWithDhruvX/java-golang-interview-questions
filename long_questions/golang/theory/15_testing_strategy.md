# 🟢 Go Theory Questions: 281–300 Testing Strategy, CI/CD, Observability

## 281. What are test doubles and how are they used in Go?

**Answer:**
Test Doubles are objects that mimic real dependencies in tests. The main types are Mocks, Stubs, and Fakes.

In Go, we prefer **Fakes** or **Interfaces**. If we depend on a Database, we define a `Repository` interface. In the test, we create a `FakeRepository` (an in-memory map) that satisfies the interface.

This allows us to simulate success and error scenarios without spinning up a Docker container. Unlike Java's Mockito which uses reflection magic, Go mocks are usually handwritten or generated code, making them type-safe and easier to debug.

---

## 282. How do you structure unit vs integration tests?

**Answer:**
Unit tests verify logic in isolation; Integration tests verify how components talk to each other.

For Unit tests, we use `_test.go` files right next to the source code. They often mock out external calls (DB, HTTP) to run in milliseconds.

For Integration tests, we typically have a separate `tests/` folder or use build tags `//go:build integration`. These tests spin up real dependencies (Postgres, Redis) via Docker. Because they are slow, we configure our CI pipeline to run unit tests on every commit, but integration tests only on merge requests or nightly builds.

---

## 283. What are flaky tests and how do you identify them?

**Answer:**
A flaky test passes sometimes and fails others, without code changes. They are the enemy of CI trust.

Common causes are **Race Conditions** (not waiting for a goroutine), **Order Dependency** (iterating a map), or **Shared State** (one test dirtying the DB for the next).

We identify them by running the test suite 100 times in a loop: `go test -count=100`. If it fails even once, it's flaky. We aggressively fix or delete flaky tests because a test suite that "cries wolf" eventually gets ignored by the team.

---

## 284. How do you write deterministic tests for concurrency?

**Answer:**
Concurrency is inherently non-deterministic, so testing it is hard. The trick is to synchronize using **Channels** specifically for the test.

We can inspect the internal state of a worker by injecting a "spy" channel. The test waits for a signal on that channel before making assertions.

Alternatively, we use `sync.WaitGroup` to ensure all background routines finish before the test asserts the final state. We never use `time.Sleep()` in tests (e.g., "Wait 100ms and hope it finished"). Any test relying on `Sleep` is a flaky test waiting to happen.

---

## 285. How do you test RESTful APIs in Go?

**Answer:**
We use the `httptest` package, which is part of the standard library.

It allows us to create a synthetic `http.Request` and pass it directly to our handler function. We record the response in an `httptest.ResponseRecorder` (which mimics a real `ResponseWriter`).

This tests the entire HTTP logic flow—routing, middleware, JSON decoding, status codes—without actually opening a network port. It’s extremely fast and allows us to verify edge cases (like 404s or 500s) easily.

---

## 286. How do you mock HTTP calls?

**Answer:**
When our code *calls* an external API, we don't want the test to hit the real internet.

We use `httptest.NewServer`. This spins up a real local HTTP server on a random port during the test. We point our client to this server's URL.

The test server can be programmed: "If you get a GET /users/1, return 500 Error." This proves our client handles failures gracefully. It’s better than mocking the network client object because it exercises the full network stack, including timeouts and context cancellation.

---

## 287. What is Golden Testing in Go?

**Answer:**
Golden Testing is using files to store expected complex output.

Instead of writing a giant string literal in the test code to match a JSON response, we save the expected output in a file `testdata/response.golden`.

The test reads the file and compares it to the actual output. If we intentionally change the logic, we run `go test -update`, which automatically overwrites the golden files with the new output. This is standard in Go for testing CLI output, JSON generators, or HTML templates.

---

## 288. How do you run tests in parallel?

**Answer:**
We verify `t.Parallel()` inside the test function.

This signals the test runner that this test is safe to run concurrently with other parallel tests. By default, `go test` runs tests sequentially within a package.

However, you must be careful with **closure variables**. A common bug is running parallel tests inside a loop using the loop variable `i`. You must rebind it: `i := i` before calling `t.Parallel()`. Running tests in parallel can drastically reduce CI times for IO-bound integration tests.

---

## 289. How do you mock time-dependent code?

**Answer:**
If your code says "Wait 10 seconds", you don't want your test to take 10 seconds.

We abstract time behind an interface: `type Clock interface { Now() time.Time }`.

In production, we use the System Clock. In tests, we inject a **Mock Clock**. We can freeze time, or "fast forward" time instantly. This allows us to verify logic like "Token expires in 24 hours" instantly by setting the mock clock to `Now() + 25 hours` and asserting the token is rejected.

---

## 290. How do you simulate DB failures in tests?

**Answer:**
We use a mock driver or a library like `go-sqlmock`.

We configure the mock to expect a query and return an error: `mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)`.

This ensures our application logic handles database outages correctly (e.g., retrying or returning a clean 503 error) rather than panicking. Testing the "Sad Path" is just as important as the "Happy Path" for robust systems.

---

## 291. How do you use GitHub Actions to test Go apps?

**Answer:**
We create a workflow `.yaml` file. It usually has steps: `Checkout` -> `Setup Go` -> `Test`.

A key best practice is **Caching**. We cache the `GOMODCACHE` and `GOCACHE` directories to speed up dependency downloads and build times.

We also run linters (`golangci-lint`) and security checks (`govulncheck`) in parallel jobs. If any job fails, the Pull Request is blocked. This automated gatekeeping ensures that no broken code ever reaches the main branch.

---

## 292. What is the structure of a Makefile for Go?

**Answer:**
Since `go build` is simple, Makefiles are used for the glue logic around the project.

Common targets are:
*   `make test`: Runs tests with race detection.
*   `make build`: Compiles for the current OS.
*   `make run`: Runs the app locally with hot-reload tools (like Air).
*   `make docker`: Builds the container image.

It serves as executable documentation. A new developer doesn't need to memorize long flags; they just type `make help` to see exactly how to operate the project.

---

## 293. How to build and test Go code in Docker?

**Answer:**
We use a **Multi-Stage Build**.

The first stage uses the heavy `golang:alpine` image to compile the app and run tests.
If tests pass, we copy the binary to a second, empty `scratch` image.

This guarantees that the artifact we deploy relies *only* on the code that passed tests in a clean environment, not on random files on the developer's laptop. It also results in tiny images (10MB) that are secure because they lack a shell or system tools.

---

## 294. What CI tools are commonly used for Go projects?

**Answer:**
**GitHub Actions** is the dominant player now for open source and many companies due to its tight integration.

**GitLab CI** is huge in enterprise. It works similarly but often uses Runners inside the company VPC.
**Jenkins** is legacy but still common; it requires more maintenance (plugins).

Regardless of the tool, the pipeline logic is the same: Lint -> Test (Unit) -> Test (Integration) -> Build -> Deploy.

---

## 295. What are the benefits of go:embed for test fixtures?

**Answer:**
Before Go 1.16, loading test data (like a sample.json) involved fragile relative paths `../../testdata/sample.json`, which broke depending on where you ran the test from.

With `//go:embed`, the file contents are compiled into the test binary as a string or byte slice.
`var sampleJSON []byte` (with embed tag).

This makes tests robust and portable. You can run the test binary on another machine (like a CI runner) without needing to copy the source files along with it.

---

## 296. How do you generate coverage reports in HTML?

**Answer:**
We run `go test -coverprofile=coverage.out`. This creates a raw data file mapping code blocks to execution counts.

Then we run `go tool cover -html=coverage.out`.

This opens a browser showing our source code. Green lines were executed; red lines were missed. As a strategy, we look for **Red Error Blocks**—error handling branches that were never triggered. These are dangerous blind spots. We add tests to force those errors to ensure our error handling logic is actually sound.

---

## 297. How to collect logs and metrics from Go services?

**Answer:**
We follow the **Push vs Pull** model.

For Logs, we **Push** to standard output (stdout). A collector (like Fluentd or Promtail) reads the stdout stream and pushes it to a central system (Elasticsearch or Loki).

For Metrics, we let Prometheus **Pull** (scrape) them. We expose a `/metrics` HTTP endpoint where the Go app lists its current stats (memory usage, request counts). Prometheus visits this page every 15 seconds to ingest the data.

---

## 298. What is structured logging in Go?

**Answer:**
Structured logging means emitting logs as machine-parsable JSON objects, not unstructured text strings.

Instead of `log.Printf("User %d logged in", id)`, we use `slog.Info("Login", "user_id", id)`.
This outputs `{"time":"...", "level":"INFO", "msg":"Login", "user_id":123}`.

This allows us to query logs in systems like Splunk or Datadog efficiently: `service="auth" AND user_id=123`. You can't do that easily with grep-able text logs. Go 1.21 introduced `log/slog` as the standard library solution for this.

---

## 299. What are common logging libraries in Go?

**Answer:**
**Zap** (by Uber) is the gold standard for high performance. It avoids reflection and allocation, making it zero-overhead suitable for hot paths.

**Zerolog** is another zero-allocation alternative, known for a great developer experience.
**Logrus** was the old standard; it’s widely used but now in maintenance mode (and slower).

Modern Go projects are increasingly adopting **slog** (Standard Library) because it provides a unified interface, allowing libraries to log without forcing a specific dependency like Zap on the application consumer.

---

## 300. How do you aggregate and search logs across services?

**Answer:**
We attach **Correlation IDs** (Trace IDs) to every log entry.

When a request enters the system (Gateway), we generate a UUID. We pass this UUID in HTTP headers to every microservice involved in the chain.

Every logger extracts this ID from the context and adds it to the JSON log: `{"trace_id":"abc-123", ...}`.
In our log aggregator (Elastic/Loki), we simply search for that Trace ID to see the full timeline of the request across 10 different services, allowing us to pinpoint exactly which service failed.
