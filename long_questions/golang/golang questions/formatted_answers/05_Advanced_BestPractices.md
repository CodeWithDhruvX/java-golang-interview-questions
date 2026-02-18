# 🔴 **81–100: Advanced & Best Practices**

### 82. How does Go handle memory management?
"Go abstracts logical memory management away from the developer, giving us the safety of a managed language with the performance of a compiled one.

It decides where to store variables automatically using **Escape Analysis**. If a variable is local and doesn't leave the function, it goes on the **Stack** (fast, self-cleaning). If it 'escapes' (like returning a pointer), it moves to the **Heap** (managed by the Garbage Collector).

I rarely have to think about `malloc` or `free`. However, I do keep an eye on *unintentional* heap allocations—like passing a large struct by value to an interface—because excessive heap usage triggers the GC more often, hurting latency."

#### Indepth
You can run `go build -gcflags="-m"` to see the compiler's escape analysis decisions. It outputs lines like "moved to heap: x" or "leaking param: p". This is the first step in optimizing "hot paths" to reduce GC pressure by keeping variables on the stack.

---

### 83. What is garbage collection in Go?
"Garbage collection (GC) is the automatic process of reclaiming memory that the program has allocated but is no longer using.

Go uses a **non-generational, concurrent, tri-color mark-and-sweep** collector. That’s a mouthful, but it basically means it runs alongside my code (concurrent), marking live objects and sweeping away the dead ones, with extremely short 'stop-the-world' pauses (often sub-millisecond).

In production, this low-latency focus is a huge win for web services. I almost never tune the GC (via `GOGC`) unless I have specific memory constraints or a batch job where throughput matters more than latency."

#### Indepth
The Go GC is a "write barrier" based collector. It trades off some CPU throughput (approx 25%) for ultra-low latency. `GOGC` defaults to 100, meaning "run GC when heap doubles". Setting `GOGC=off` disables it, which is useful for short-lived CLI tools that exit before running out of RAM.

---

### 84. How do you profile CPU and memory in Go?
"I use the built-in tooling ecosystem, specifically **pprof**. It’s not an external plugin; it’s part of the standard library.

For web apps, I import `net/http/pprof` to expose debug endpoints. I can then physically download a snapshot of the CPU usage or Heap allocations and visualize it with `go tool pprof -http=:8080 profile.out`.

This visualization—the 'flame graph'—is my first stop when debugging performance. If I see a wide bar for `json.Unmarshal`, I know exactly where to optimize. It removes guesswork entirely."

#### Indepth
`pprof` adds overhead, so don't leave it running on public internet endpoints (security risk). It profiles by sampling the stack traces 100 times per second. This statistical approach means it effectively finds bottlenecks without slowing down the application significantly during the profile.

---

### 85. What is the difference between compile-time and runtime errors?
"**Compile-time errors** are caught by the compiler *before* the binary is even built. These are syntax errors, type mismatches, or unused variables. I love these because they prevent me from shipping broken code.

**Runtime errors** happen *while* the program is executing. These are things like `nil pointer dereference`, `index out of range`, or division by zero. In Go, these cause a **panic** and crash the Goroutine (or the whole app) if not recovered.

My goal is always to shift errors from runtime to compile-time—for example, using strong typing instead of `interface{}`/`any` so the compiler catches mismatches immediately."

#### Indepth
Some runtime errors can be caught by **recover**, but not all. "Fatal" errors like concurrent map writes or stack overflows crash the program instantly and cannot be recovered. This fail-fast mechanism is intentional to prevent running in an undefined state.

---

### 86. How to use `go test` for unit testing?
"Go treats testing as a first-class citizen. I don't need a third-party framework like JUnit or PyTest.

I simply create a file named `_test.go` and write functions starting with `Test...` that accept a `*testing.T`. I run `go test ./...` to execute them.

I rely on `t.Error` (log and continue) or `t.Fatal` (log and stop) to report failures. The 'stdlib-only' approach encourages me to write simple, decoupled code that is easy to setup and teardown without complex mocking magic."

#### Indepth
`go test` compiles the package *and* the test files into a separate binary, executes it, and deletes it. This means `TestFunc` has access to unexported identifiers in the same package, allowing white-box testing of internal logic without exporting everything.

---

### 87. What is table-driven testing in Go?
"It’s the idiomatic pattern for testing multiple scenarios without duplicating logic. It’s cleaner and more maintainable than writing ten separate test functions.

I define a struct slice (the 'table') where each element contains the **input**, **expected output**, and a **test name**. I loop over this slice and use `t.Run(tt.name, ...)` to execute the assertion logic for each case.

I use this for almost every utility function I write. It makes it trivial to add edge cases (like 'empty string' or 'negative number') by just adding one line to the slice."

#### Indepth
Use `t.Parallel()` inside the loop to run cases concurrently. This drastically speeds up test suites involving I/O (like DB calls). However, be careful to capture the loop variable `func(tt args) { t.Parallel(); ... }(v)` or `tt := tt` to avoiding sharing scope (though Go 1.22 fixes this).

---

### 88. How to benchmark code in Go?
"I write a function starting with `Benchmark...` that takes a `*testing.B` argument.

Inside, I write a `for` loop: `for i := 0; i < b.N; i++`. The testing framework automatically adjusts `b.N` (running it 1, 100, 10,000 times) until it runs long enough to get a statistically significant measurement of time-per-operation (ns/op).

I use this constantly to make data-driven decisions—like choosing between `strings.Join` versus `fmt.Sprintf`. I never assume one is faster; I benchmark it."

#### Indepth
Be careful of compiler optimizations in micro-benchmarks. If you calculate a result but don't use it, the compiler might delete the entire function call. Always assign the result to a package-level variable (like `var result int`) or use `runtime.KeepAlive()` to ensure code execution.

---

### 89. What is `go mod` and how does it work?
"`go mod` is the standard dependency management system introduced to replace the old `GOPATH` workspace.

It uses a `go.mod` file to define the module's path and versioned dependencies (e.g., `require github.com/gin-gonic/gin v1.7.0`). Changes are tracked in `go.sum`, which contains cryptographic hashes to guarantee that the code I downloaded today is identical to what I download tomorrow.

Contextually, it made Go reproducible. I can clone a repo anywhere on my disk, run `go build`, and it just works, fetching exact versions automatically."

#### Indepth
`go mod` also supports **Minimal Version Selection (MVS)**. Unlike NPM which prefers the *latest* allowed version, Go prefers the *oldest* allowed version that satisfies requirements. This conservative approach increases stability, as you don't silently upgrade to a buggy new patch version unless you explicitly ask for it.

---

### 90. What is vendoring in Go modules?
"Vendoring is the practice of copying all third-party dependencies into a local `vendor/` directory inside my project.

I enable it by running `go mod vendor` and building with `-mod=vendor`. This essentially 'freezes' the source code of my dependencies into my git repository.

I use this for critical enterprise builds where we cannot rely on external repositories (like GitHub) being online during deployment, or to satisfy security audits that require us to own every line of code we ship."

#### Indepth
Vendoring also protects you from "LeftPad" incidents (deleted upstream repos). However, it bloats your git repo size. As a middle ground, many companies use a **Go Module Proxy** (like Athens or Artifactory) to cache dependencies without committing them to source control.

---

### 91. How to handle versioning in modules?
"Go enforces **Semantic Versioning** (SemVer).

The critical rule is the **Import Compatibility Rule**: If an old package imports `my/lib`, it should work with any new version of `my/lib` that has the same Major version. If I introduce breaking changes, I *must* increment the Major version (v2).

In Go, v2 is technically a different package path (`github.com/my/lib/v2`). This allows a program to import both v1 and v2 simultaneously, solving the 'Diamond Dependency' problem that plagues other languages."

#### Indepth
Breaking changes are detected by tools like `apidiff`. When releasing v2, you can use a `type Alias` in v2 to point to v1 types (or vice-versa) to allow interoperability, but typically v2 is a fresh start. The `/v2/` suffix in the `go.mod` module path is mandatory.

---

### 92. How do you structure a Go project?
"Go is unopinionated, but the community has converged on the **Standard Go Project Layout**.

I typically use:
*   `cmd/`: Main applications (entry points).
*   `pkg/`: Library code safe for other projects to import.
*   `internal/`: Private library code (compiler enforced).
*   `api/`: OpenAPI/Protobuf definitions.

This structure keeps the root directory clean and clearly separates 'runnable' code from 'reusable' logic. For microservices, I might stick to a flatter structure initially, but I move to this layout as complexity grows."

#### Indepth
This "Standard Layout" is not official Go policy (the Go team uses a flat structure), but it is the industry standard. `cmd` prevents library pollution (you can't import `main`). `pkg` is explicit "this is public". `internal` is purely for safety. It scales well for monorepos.

---

### 93. What is the idiomatic way to name Go packages?
"Short, lowercase, and singular. No underscores (`snake_case`) or mixed caps (`camelCase`).

The name should describe *what* it provides, not what it *is*. For example, use `http` instead of `http_utils` or `network`.

I also consider how it reads in code. We want `bytes.Buffer`, not `util.Buffer`. If I find myself naming a package `common` or `shared`, I usually take a step back and redesign, because those are rarely good package names."

#### Indepth
Package names occupy the "global namespace" of the importer. If you name your package `log`, every file importing it can't use a variable named `log`. This is why we have `zap` and `logrus` instead of `log`. Avoid "stuttering": `user.User` is verbose; `user.Profile` or just `user.ID` is better.

---

### 94. What is the purpose of the `internal` package?
"It’s a special compiler-enforced visibility rule.

Code inside a directory named `internal/` can only be imported by packages rooted at the same parent. If I have `github.com/my/project/internal/foo`, it cannot be imported by `github.com/other/project`.

I use this liberally to hide implementation details. It gives me the freedom to refactor public APIs or messy helper functions without worrying about breaking downstream users who might have depended on them."

#### Indepth
Use `internal` for your "Business Logic Domain" if you want to force all external access through a clear API layer (like gRPC or HTTP handlers). This prevents other teams in a monorepo from bypassing your API and calling your database logic directly.

---

### 95. How do you handle logging in Go?
"I avoid the standard `log` package in production because it lacks structure (it just prints strings).

I use **structured logging**—usually the new `slog` package or **Zap**. This allows me to output logs as JSON: `{"level":"info", "user_id":123, "msg":"login success"}`.

This is critical because I don't read logs on a server; I query them in a system like Datadog or Elasticsearch. Structure lets me filter by `user_id` or `level` instantly, whereas parsing text logs with RegEx is a nightmare."

#### Indepth
Structured logging also enables **contextual logging**. You can inject a `logger` into `context.Context` (or pass it down) with fields pre-attached (`logger.With("request_id", id)`). Every subsequent log line automatically includes that ID, enabling distributed tracing.

---

### 96. What is the difference between `log.Fatal`, `log.Panic`, and `log.Println`?
"`log.Println` handles standard informational messages. It writes to the output and execution continues normally.

`log.Fatal` writes the message and immediately calls `os.Exit(1)`. This terminates the program **hard**—deferred functions (like closing DB connections) are *not* run. I only use this in `main()` during startup if the app cannot boot.

`log.Panic` writes the message and calls `panic()`. This unwinds the stack, running deferred functions, and crashes unless `recover()` is used. It’s useful if I want to allow a middleware to potentially catch the crash."

#### Indepth
In library code, *never* call `log.Fatal` or `os.Exit`. It robs the user of the chance to handle the error or cleanup. Always return an `error`. Only `main.main` implies the right to decide "this process must die".

---

### 97. What are build tags in Go?
"Build tags are a mechanism to include or exclude files from the build based on conditions.

I place a comment `//go:build linux` at the top of a file. The compiler will completely ignore this file if I’m building for Windows.

I use this heavily for cross-platform system calls (one file for Windows syscalls, one for Unix) and for separating **Integration Tests** (`//go:build integration`) so they don’t run during my fast unit test cycle."

#### Indepth
Build tags can be combined: `//go:build linux && !cgo`. They can also be file suffixes: `file_windows.go` is automatically treated as if it had `//go:build windows`. This is how the standard library implements OS-specific logic (like file locking) cleanly.

---

### 98. What are cgo and its use cases?
"**Cgo** is a mechanism that allows Go packages to call C code.

I can write `import "C"` and call C functions directly. It’s powerful but dangerous: it complicates builds (I need a C compiler), breaks cross-compilation simplicity, and introduces performance overhead due to stack switching.

I only use it when I absolutely have to—usually to bind to a mature C library like **SQLite**, **OpenGL**, or proprietary hardware drivers that don't have a native Go equivalent."

#### Indepth
Cgo calls have a high overhead (~170ns) compared to a Go function call (~2ns). This is because the Go runtime must lock the stack and switch to the system stack (pthreads) to run C code safely. Chatty interfaces across the Cgo boundary destroy performance. Batch your calls.

---

### 99. What are some common Go anti-patterns?
"A major one is **returning interfaces**. I intuitively want to return an Interface to be 'abstract', but idiomatic Go says: 'Accept interfaces, return structs.' Returning structs lets the caller decide which interface they need.

Another is **concurrency everywhere**. Spinning up a goroutine for a tiny calculation is often slower than doing it synchronously.

Lastly, **ignoring errors**. Writing `_` to suppress an error is almost always a bug waiting to happen. I treat errors as values that must be handled or explicitly logged."

#### Indepth
Another anti-pattern is **Package-level state** (global variables). It makes testing impossible (tests pollute each other) and concurrency dangerous. Always use **Dependency Injection**: pass dependencies (DB, Config) into structs/functions rather than reading globals.

---

### 100. What are Go code quality tools (lint, vet, staticcheck)?
"**go vet** is the official static analyzer found in the Go toolchain. It catches suspicious constructs that compile but are likely wrong (like `Printf` mismatches).

**staticcheck** is a third-party tool that goes much deeper, finding unused code, infinite loops, and deprecated API usage.

I bundle these into **golangci-lint**. It runs dozens of linters in parallel. I configure it to run on every Pull Request. It removes the 'nitpicking' from code reviews because the robot acts as the bad cop."

#### Indepth
Configure your linter to be strict but practical. Enable `errcheck`, `staticcheck`, `soviet` (shadowing), and `revive`. Disable formatting rules that fight with `gofmt`. CI should block merges on lint failures. This enforces a consistent "dialect" of Go across the team.

---

### 101. What are the best practices for writing idiomatic Go code?
"I follow the philosophy of **Effective Go**.

Keep things simple. Clarity is better than cleverness.
Usage of `gofmt` is non-negotiable.
Errors are handled immediately, keeping the 'happy path' unindented on the left side (Short-Variable-Declaration-in-If pattern).
Comments should document *why*, not *what*.

If I write code that looks like Java (getters/setters, deep inheritance), I’m fighting the language. I verify to write code that looks like the standard library."

#### Indepth
Idiomatic Go often looks "dumb" to Enterprise Java developers. It repeats code (`if err != nil`) to avoid abstraction magic. It prefers **Composition over Inheritance**. It values **Readability** (local reasoning) over DRY (if DRY requires complex abstraction). "A little copying is better than a little dependency." - Rob Pike.
