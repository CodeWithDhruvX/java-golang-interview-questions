# 🟢 Go Theory Questions: 181–200 Testing and Security

## 181. How do you implement Fuzz Testing in Go?

**Answer:**
Fuzzing is an automated testing technique where you throw random, malformed data at your function to see if it crashes.

Since Go 1.18, this is built-in. You write a test function starting with `FuzzXxx` that accepts a `*testing.F`. You provide a few "seed" inputs (valid examples), and then the Go runtime automatically mutates those inputs—creating millions of variations—and feeds them into your function in parallel.

It is indispensable for parsers and validators. I’ve found bugs in production JSON parsing logic that extensive unit tests never caught, simply because I didn't think to test inputs with weird unicode characters or infinite lengths.

---

## 182. How do you mock HTTP clients in Go tests?

**Answer:**
We rarely mock the `http.Client` itself because it's a struct, not an interface. Instead, we mock the **Transport**.

You can assign a custom function to `client.Transport`. This function intercepts the request and returns a precanned `http.Response` without ever touching the network.

Alternatively, for broader integration tests, we use `httptest.NewServer`. This spins up a real local HTTP server on a random port during the test. You point your API client to this localhost URL. This is often better than mocking because it proves your code actually speaks valid HTTP.

---

## 183. How do you achieve high test coverage in Go?

**Answer:**
Coverage is a metric, not a goal. But to improve it, we use Table-Driven Tests.

By defining a slice of struct cases—`{input, expected, name}`—we can easily add edge cases (empty strings, nil pointers, massive numbers) without writing new test logic.

We run `go test -coverprofile=c.out` and then `go tool cover -html=c.out`. This generates a visual HTML file showing exactly which lines of code are red (untested). We focus on turning red error-handling branches green, as that's where bugs usually hide.

---

## 184. How do you test race conditions in Go?

**Answer:**
You cannot "write" a test that deterministically proves a race condition exists, because the scheduler is unpredictable.

Instead, we enable the **Race Detector** using `go test -race`.

We write tests that intentionally spawn multiple goroutines to hammer the same function concurrently. If there is any unsynchronized access, the Race Detector (which tracks memory access at runtime) will spot it and fail the test. We run this in CI for every single build.

---

## 185. How do you test helper functions that aren't exported?

**Answer:**
In Go, test files in the same package (e.g., `package user`) have access to private identifiers.

So, for `user.go`, we write `user_test.go` utilizing `package user`. This allows us to white-box test internal logic like `validatePassword()` directly.

However, for public API testing, we often suffix the test package name: `package user_test`. This forces us to import `user` as an external dependency, ensuring we only test the public interface (Black Box testing), which leads to less brittle tests when refactoring internals.

---

## 186. How do you secure JWTs (JSON Web Tokens) in Go?

**Answer:**
Security starts with the signing algorithm. We strictly use distinct keys for signing and verification if possible (ECDSA), or a strong secret for HMAC.

In Go, we use `golang-jwt/jwt`. The detailed pitfall is the **"None" algorithm attack**. When parsing a token, you must explicitly supply the algorithm you expect.

`jwt.Parse(token, func(t *Token) (any, error) { ... })`

Inside that callback, we explicitly check `if t.Method != jwt.SigningMethodHS256 { return nil, error }`. Without this check, an attacker could strip the signature, set alg to "none", and bypass auth.

---

## 187. How do you prevent SQL Injection in Go?

**Answer:**
The `database/sql` package prevents this by default **if** you use placeholder parameters.

We never write `fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)`. That’s the vulnerability.

Instead, we write `db.Query("SELECT * FROM users WHERE name = ?", name)`. The driver sends the query template and the data separately to the database engine. The DB treats the input strictly as data, not executable code, making injection mathematically impossible.

---

## 188. How do you handle secrets (passwords/API keys) in Go?

**Answer:**
We follow the strict rule: **No secrets in code.**

We use environment variables or a secret manager (like HashiCorp Vault or AWS Secrets Manager). In Go, we read these at startup into a config struct.

For the application logic, we ensure these fields are never logged. We often use a custom type `type Secret string` and implement the `String()` method to return `"*****"` so that if a developer accidentally logs the config configuration, the implementation prevents the secret from leaking into Splunk or Datadog.

---

## 189. How do you implement CSRF protection in Go?

**Answer:**
Cross-Site Request Forgery handling depends on your architecture.

If you are building a Single Page App (React) with stateless JWTs, correct CORS configuration and `SameSite=Strict` cookies are usually sufficient.

If you are building a server-rendered Go app (HTML templates), you need a CSRF token. We use middleware (like `gorilla/csrf`) that generates a random token associated with the session. This token is injected into every `<form>` as a hidden field. The server rejects any POST request that acts on the session but lacks this valid token.

---

## 190. How do you hash passwords securely in Go?

**Answer:**
We never use generic hash functions like MD5 or SHA256 for passwords. They are too fast, making them vulnerable to brute-force attacks.

We use **Argon2** or **Bcrypt** (via `golang.org/x/crypto/bcrypt`).

The standard pattern is `bcrypt.GenerateFromPassword(password, cost)`. The "cost" factor makes the function intentionally slow (milliseconds). This slowness is the security feature—it ensures that an attacker with a stolen database can only guess a few passwords per second, not billions.

---

## 191. How do you sanitize input against XSS (Cross Site Scripting)?

**Answer:**
Go's `html/template` package is context-aware and secure by default.

If you pass a string containing `<script>alert(1)</script>` to a template, Go automatically escapes it to `&lt;script&gt;...` so it renders as text, not code.

The danger exists only if you bypass this using `template.HTML` type, which marks the string as "Safe". We audit uses of `template.HTML` rigorously. For JSON APIs, we don't sanitize on input; we rely on the frontend (React/Vue) to escape data when rendering, as the backend shouldn't assume how the data will be displayed.

---

## 192. How do you implement HMAC authentication?

**Answer:**
HMAC (Hash-based Message Authentication Code) is used to verify that a request hasn't been tampered with.

We use `crypto/hmac`. When a client sends a request, they hash the body + a secret key and send the hash in a header.

On the server, we take the received body and our copy of the secret, run the same hash, and compare. Typically, we use `hmac.Equal(sig1, sig2)` to compare them. Crucially, `Equal` is a **constant-time comparison** function. If you used `==`, it would return false faster for different first characters, allowing timing attacks.

---

## 193. How do you validate TLS certificates in Go?

**Answer:**
Go's `http.Client` validates TLS certificates automatically against the system's root CA store.

However, in internal microservices (mTLS), we often use self-signed certificates. In this case, we configure `tls.Config{RootCAs: myCertPool}`.

We broadly avoid `InsecureSkipVerify: true`. Using that flag turns off all security. If we need to talk to a local tool with a bad cert during development, we use build tags to enable that flag only in local-dev builds, ensuring it never inadvertently reaches production.

---

## 194. How do you audit Go dependencies for vulnerabilities?

**Answer:**
We use `govulncheck`, the official tool from the Go team.

It connects to the Go Vulnerability Database. Unlike other scanners that just check `go.mod` versions, `govulncheck` uses call graph analysis.

It tells you: "Yes, you use a vulnerable library, BUT you verify that you never actually call the vulnerable function." This reduces false positives massively, letting us focus on the critical security patches that actually affect our application.

---

## 195. How do you write Integration Tests with Docker?

**Answer:**
We use a library called **Testcontainers for Go**.

In your test setup (`TestMain`), you define a Postgres container. When you run `go test`, the code programmatically spins up a real Docker container, waits for port 5432 to be ready, runs migrations, and returns the connection string.

The tests run against this throwaway database, and at the end, the container is destroyed. This allows us to write "End-to-End" tests that run entirely on a laptop without needing an external staging environment.

---

## 196. How do you perform table-driven tests?

**Answer:**
This is the idiomatic Go testing style.

We define a struct `type testCase struct { name string; input int; want int; err error }`.
Then we make a slice `tests := []testCase{ ... }`.

We loop over this slice:
`for _, tc := range tests { t.Run(tc.name, func(t *testing.T) { ... }) }`.

This structure decouples the test logic from the test data. Adding a new regression test is just adding one line to the slice. It keeps specific test files readable even when covering 50 different edge cases.

---

## 197. How do you test code that depends on time?

**Answer:**
Time is the enemy of deterministic testing. `time.Now()` changes every time you run it.

We solve this by dependency injection. We define a `Clock` interface with a `Now()` method.

In production, we inject the `RealClock`. In tests, we inject a `FakeClock` that returns a fixed time. This allows us to verify logic like "Token expires in 5 minutes" by setting the fake clock to specific instants, avoiding flaky tests that sleep or wait for real time to pass.

---

## 198. How do you implement Rate Limiting security?

**Answer:**
Rate limiting is a DoS protection mechanism.

We use middleware based on the **Token Bucket** algorithm (often `golang.org/x/time/rate`). We map IP addresses to limiters.

`limiters.Get(ip).Allow()` returns true or false. If false, we immediately return status 429 Too Many Requests. For distributed systems (where multiple server pods share the limit), we move this logic out of memory and into **Redis** using a Lua script to atomically decrement counters across the cluster.

---

## 199. What are Subtests in Go?

**Answer:**
Subtests allow you to define a hierarchy of tests using `t.Run("SubName", func...)`.

They are useful for grouping. You might have `TestUserCRUD`, and inside it `t.Run("Create")`, `t.Run("Update")`.

The big advantage is Setup/Teardown. You can do expensive setup (like DB connection) once in the parent test, and all subtests share it. They also allow you to run specific subsets from the command line: `go test -run TestUserCRUD/Create`.

---

## 200. How do you verify your code is thread-safe?

**Answer:**
Static analysis can't prove thread safety fully. We rely on the **Race Detector** during execution.

But structurally, we design for it. We encapsulate mutable state. If a struct has a map, we make the map private and protect it with a Mutex. We don't export the map.

By narrowing the scope where data can be modified (only via exported methods), we make it easier to audit. If the `Update()` method has a Lock, and it's the only way to write data, we can be reasonably confident in safety.
