# 🔴 **181–200: Tools, Testing, CI/CD, Ecosystem**

### 181. What is `go vet` and what does it catch?
"`go vet` is the built-in static analysis tool.

It catches logic errors that compile but are likely bugs.
Common examples: `Printf` arguments not matching the format string, unreachable code after a return, or passing a `sync.Mutex` by value (which breaks the lock).
I run it automatically in CI, but I usually rely on `golangci-lint` which includes `vet` plus many other checkers."

#### Indepth
`go vet` uses **heuristics**, meaning it's not 100% precise but very fast. It checks for things like "CopyLock" (copying a struct that contains a Mutex). If you ever see a "copylocks" error, you are likely introducing a race condition by copying a lock state instead of sharing the pointer.

---

### 182. How does `go fmt` help maintain code quality?
"It ends all wars about code style.

It rewrites my source code to a canonical format (tabs for indentation, specific bracket placement).
Because it’s standard, I can jump into any Go project in the world and read it immediately without adjusting to a custom style guide. I configure my editor to run it on save."

#### Indepth
`gofmt` is technically a printer. `goimports` is a superset of `gofmt` that also manages your import block (adds missing, removes unused). Most developers use `goimports` as their "format on save" tool. In modern editors, `gopls` handles both formatting and imports efficiently.

---

### 183. What is `golangci-lint`?
"It’s a linter aggregator—the Swiss Army Knife of Go code quality.
It runs dozens of linters in parallel (including `vet`, `staticcheck`, `errcheck`).

I configure it to be strict: it forces me to handle every error, avoid unused variables, and limit cognitive complexity.
It’s much faster than running tools individually because it reuses the Go compilation cache."

#### Indepth
You can define a `.golangci.yml` file in your repo root to clamp down on specifics. For example, enabling `wsl` (Whitespace Linter) enforces empty lines between assignments and returns, making code more readable. Enabling `gocritic` finds subtle performance and style issues.

---

### 184. What is the difference between `go run`, `go build`, and `go install`?
"`go run` compiles and executes the binary in a temporary directory. I use it for local development and scripts.
`go build` compiles the binary and leaves it in the current directory. I use it to verify that code compiles.
`go install` compiles and moves the binary to `$GOPATH/bin`. I use it for installing CLI tools like `gopls` or `golangci-lint` so I can run them globally."

#### Indepth
Since Go 1.16, `go install package@version` is the preferred way to install tools without polluting your project's `go.mod` file. `go get` is now deprecated for installing binaries and is strictly for adding dependencies to the current module.

---

### 185. How does `go generate` work?
"It’s a tool for code generation, triggered by comments.

I add a comment `//go:generate stringer -type=Pill` in my source.
When I run `go generate ./...`, it finds these comments and executes the command.
I use it heavily for generating Mocks (`mockery`), Protobuf code (`protoc`), or Enums (`stringer`). It allows me to automate the creation of boilerplate code."

#### Indepth
`go generate` is **not** part of the build process (`go build`). You must run it manually. A common pattern is to have a `Makefile` with a `generate` target that runs `go generate ./...` before building, ensuring that all generated mocks and protobufs are up to date.

---

### 186. What is a build constraint?
"Also known as a Build Tag. It tells the compiler *when* to include a file.

`//go:build linux` or `//go:build integration`.
I use it for **OS-specific code** (e.g., using syscalls on Linux vs Windows).
I also use it for **test separation**: `//go:build integration` keeps my slow integration tests out of the standard `go test ./...` cycle unless I explicitly pass `-tags=integration`."

#### Indepth
The syntax changed in Go 1.17. The old syntax was `// +build linux`. The new syntax is `//go:build linux`. The new compiler supports boolean expressions like `//go:build linux || (darwin && amd64)`. You should prefer the new syntax.

---

### 187. How do you write tests in Go?
"I write functions starting with `Test` in `_test.go` files.

`func TestAdd(t *testing.T) { ... }`.
I prefer **Table-Driven Tests**: I define a slice of structs (inputs and expected outputs) and loop over them.
Inside the loop, I verify `got := Add(tc.a, tc.b)`. If `got != tc.want`, I call `t.Errorf`. It’s clean, extensible, and covers many edge cases easily."

#### Indepth
For large test outputs (like multi-line strings or JSON), `t.Errorf` diffs can be hard to read. Use `github.com/google/go-cmp/cmp` to display a clean line-by-line diff of struct fields. It is much more readable than standard `reflect.DeepEqual` failure messages.

---

### 188. How do you test for expected panics?
"Panics are special because they crash the test runner.

I catch them using `defer`.
`defer func() { if r := recover(); r == nil { t.Errorf("expected panic") } }()`
I put this at the top of the test. Then I call the code that should crash. If it *doesn't* crash, the deferred function sees a `nil` recover and fails the test.
Libraries like `testify` have a helper `assert.Panics` that wraps this logic neatly."

#### Indepth
Don't use `assert.Panics` for normal error handling logic. Panics in Go are for *truly* exceptional, unrecoverable states (like an `init()` function with invalid config). If a function panics on bad user input, that is a bug; it should return an `error`.

---

### 189. What are mocks and how do you use them in Go?
"Mocks are fake implementations of interfaces.

I use **vektra/mockery** to generate them automatically from my interfaces.
If `Service` depends on `Database`, I pass a `MockDatabase`.
In the test: `mockDB.On("GetUser", 1).Return(&User{}, nil)`.
This isolates the Service logic. I verify that the Service calls the DB correctly, without needing a real database running."

#### Indepth
Be careful with "over-mocking". If you mock everything, you end up testing your mocks, not your code. If the logic is simple, prefer using a real in-memory implementation (e.g., a map instead of a DB) rather than a generated mock. This is called a "Fake".

---

### 190. How do you use the `testing` and `testify` packages?
"I use the standard `testing` package for the test structure (`t.Run`, `t.Parallel`).
I use **testify** for assertions.

Instead of writing `if got != want { t.Errorf(...) }`, I write `assert.Equal(t, want, got)`.
It provides better error messages (showing the diff) and makes the test code much more readable."

#### Indepth
`testify/require` is a sibling of `assert`. The difference is that `require.NoError(t, err)` calls `t.FailNow()` (stop test immediately), whereas `assert` calls `t.Fail()` (continue test). Use `require` for setup steps where continuing makes no sense (e.g., DB connection failed).

---

### 191. How do you structure test files in Go?
"I place them right next to the source code. `user.go` and `user_test.go` live together.

This allows me to test unexported functions (white-box testing).
If I strictly want black-box testing (testing only the public API), I use a different package name in the test file: `package user_test`. This forces me to import `user` and use it exactly like a consumer would."

#### Indepth
The `user_test` package pattern solves cyclic dependencies. If `user` needs to test integration with `auth`, but `auth` imports `user`, you can't test inside `user`. Moving the test to `user_test` breaks the cycle because `user_test` imports both `user` and `auth`.

---

### 192. What is a benchmark test?
"It measures the performance of a function.

`func BenchmarkHash(b *testing.B)`.
The framework calls my function `b.N` times. It automatically adjusts `N` (100, 1000, 1M) until it gets a stable timing measurement.
I use it to catch performance regressions or to compare implementation A vs Implementation B (e.g., `fmt.Sprintf` vs `strconv.Itoa`)."

#### Indepth
Run benchmarks with `go test -bench=. -benchmem`. The `-benchmem` flag shows memory allocations per operation. A function might be fast but generate 1000 allocations (GC pressure). Optimizing for **0 allocs/op** is often better for system stability than raw CPU speed.

---

### 193. How do you measure test coverage in Go?
"I use the built-in cover tool.
`go test -coverprofile=c.out ./...`.

Then I view it: `go tool cover -html=c.out`.
It opens a browser showing my code. Green lines are covered, red lines are not.
I aim for high coverage (80%+) but I don't obsess over 100%. I ensure the *critical path` and *error handling* branches are covered."

#### Indepth
100% coverage is often a vanity metric. It usually forces you to write useless tests for trivial getters/setters or error checks that "can't happen". Focus on branch coverage for complex logic. Use Codecov or Coveralls in CI to prevent coverage *regressions* in Pull Requests.

---

### 194. How do you test concurrent functions?
"Testing concurrency is tricky due to race conditions.

I use `sync.WaitGroup` or channels to synchronize the test with the goroutines.
Crucially, I **always** run with the Race Detector: `go test -race`.
Standard tests might pass even with a race, but the race detector will spot the unsynchronized memory access and fail the build."

#### Indepth
When testing concurrent code, don't use `time.Sleep()` to "wait for the goroutine". This leads to flaky tests (fails on slow CI). Always use `WaitGroup` or channels to synchronize deterministically. If you must wait for a condition, use `assert.Eventually` (polling).

---

### 195. What is a race detector and how do you use it?
"It’s a compiler feature that instruments code to track memory accesses at runtime.

If two goroutines access the same variable concurrently, and at least one is a write, it’s a race.
I enable it with `-race`.
It slows down execution by ~10x, so I don't run it in production, but it is **mandatory** in my CI pipeline. It catches bugs that are almost impossible to debug manually."

#### Indepth
The Race Detector algorithm is based on **Vector Clocks** (ThreadSanitizer). It detects *unsynchronized* access. It does not detect *deadlocks* or *logical* races (e.g., A updates before B but you wanted B before A). It only proves that memory was accessed safely.

---

### 196. What is `go.mod` and `go.sum`?
"**go.mod** is the manifest. It lists the module name, the Go version, and the direct dependencies (with versions like v1.2.3).
**go.sum** is the lockfile/checksums. It contains the cryptographic hash of every module version used.

Its purpose is security. If an attacker hacks a library I use and changes the code for v1.2.3, the hash changes. `go build` notices the mismatch in `go.sum` and refuses to build, protecting my supply chain."

#### Indepth
`go.sum` is not a lockfile in the npm/cargo sense (it doesn't resolve the dependency tree). It's strictly a checksum database. You can have multiple versions of the same library in `go.sum` if different transitive dependencies ask for them. `go mod graph` shows the full tree.

---

### 197. How does semantic versioning work in Go modules?
"Go enforces SemVer strictly.

`v1.x.x` versions are compatible. I can upgrade safely.
`v2.x.x` (Direct Major Version) is treated as a **different module**. The import path changes to `github.com/lib/foo/v2`.
This allows me to use `v1` and `v2` of the same library in the same binary (Diamond Dependency problem solved), which is unique to Go."

#### Indepth
The compiler treats `github.com/foo/v2` as a completely different string than `github.com/foo`. They cannot be cast to each other. This strict separation allows the ecosystem to move forward without "DLL Hell", but it means you must upgrade your imports manually when migrating major versions.

---

### 198. How to build and deploy a Go binary to production?
"I build a static binary: `CGO_ENABLED=0 go build -o app`.
This removes dependencies on system libraries (libc).

I package it in a **Distroless** or **Scratch** Docker image.
The resulting image is tiny (10-20MB) and secure (no shell). I simply copy the binary and the CA certificates. This is the gold standard for Go deployments."

#### Indepth
`CGO_ENABLED=0` is key. If you don't set this, `net` package might dynamically link to the host's `glibc` DNS resolver. If your Docker container (Alpine/Scratch) doesn't have `glibc`, the binary will crash with "file not found". Static builds bundle the pure-Go DNS resolver.

---

### 199. What tools are used for Dockerizing Go apps?
"I use standard **Docker**.
I write a Multi-Stage Build.

Stage 1 (`golang:alpine`): Compiles the app.
Stage 2 (`gcr.io/distroless/static`): Copies only the binary.
I often use **Ko** (`ko build`), which builds OCI images directly from Go code without a Dockerfile. It’s incredibly fast and easy for Kubernetes deployments."

#### Indepth
`ko` is powerful because it analyzes your `import` paths to minimalize the image. It doesn't use `docker daemon`. It builds the tarball and pushes the layers directly to the registry. This is safer (no root privileges needed) and faster in CI environments.

---

### 200. How do you set up a CI/CD pipeline for a Go project?
"I almost always use **GitHub Actions**.

Step 1: **Lint**. `golangci-lint run`.
Step 2: **Test**. `go test -race -cover ./...`.
Step 3: **Build**. `go build`.
If all pass, I build the Docker image and push it to the registry. I define these steps in `.github/workflows/ci.yml`. It ensures no bad code never reaches the main branch."

#### Indepth
Cache your modules! `go mod download` can be slow. In GitHub Actions, use `actions/setup-go` with `cache: true`. It automatically hashes your `go.sum` and caches `~/go/pkg/mod`. This cuts build times from minutes to seconds for large projects.
