# 🟢 Go Theory Questions: 81–100 Advanced Concepts & Best Practices

## 81. How does Garbage Collection (GC) work in Go?

**Answer:**
Go uses a non-generational, concurrent, **Tri-color Mark and Sweep** garbage collector.

It assumes all memory is initially "White" (trash). It starts from the roots—like global variables and stack frames—and marks reachable objects as "Grey" then "Black." Anything left White at the end is swept away.

The key feature is that it runs **concurrently** with your program. It doesn't stop the world for seconds like older Java GCs. It creates tiny millisecond pauses using "Write Barriers" to track changes while it's running. This makes Go excellent for low-latency servers, though it might use slightly more CPU to keep up with the housekeeping.

---

## 82. What is the difference between specific OS threads and Goroutines?

**Answer:**
It's a difference of scale and management. An OS thread is heavy—it has a fixed 1MB stack and is managed by the kernel, making context switches expensive.

A Goroutine is user-space application thread. It starts with a tiny 2KB stack that grows dynamically. The Go runtime multiplexes thousands of Goroutines onto a small number of OS threads (the **M:N model**).

This means you can spin up 100,000 goroutines on a laptop without crashing memory, whereas 100,000 threads would kill the OS. This availability changes how we design software; we don't need "thread pools"—we just spawn a goroutine for every request.

---

## 83. What is escape analysis?

**Answer:**
Escape analysis is a compiler optimization phase that decides where to store variables: on the **Stack** or on the **Heap**.

Ideally, everything goes on the Stack because it's practically free—allocation is just moving a pointer, and cleanup happens automatically when the function returns. The Heap is expensive because complete cleanup requires the Garbage Collector to run.

The compiler looks at a variable; if its reference "escapes" the function (like returning a pointer or storing it in a global map), it forces it to the Heap. If it stays local, it writes it to the Stack. Understanding this helps us write high-performance code by avoiding unnecessary heap allocations.

---

## 84. How do you handle panics in a web server?

**Answer:**
You handle them with **Middleware** utilizing the `recover()` function.

A panic in a single handler (like a NIL pointer deference) will crash the entire program by default. In a web server, that's unacceptable—one bad request shouldn't kill the server for everyone else.

So, we wrap our entire request handler in a Deferred function that calls `recover()`. If a panic occurs, we catch it, log the stack trace, and return a 500 Internal Server Error to the user, allowing the main server process to continue running smoothly.

---

## 85. What are build tags?

**Answer:**
Build tags are special comments like `//go:build linux` placed at the top of a file.

They tell the compiler: "Only include this file if we are building for Linux."

This is how Go handles platform-specific code without messy if-statements everywhere. We have `file_windows.go` and `file_linux.go`. The compiler picks the right one. We also use them for integration tests—tagging files with `//go:build integration` so they don't run during fast unit tests.

---

## 86. How to optimize Go programs?

**Answer:**
Optimization in Go is data-driven, not guesswork. We use the **pprof** tool.

You run the program, capture a CPU or Memory profile, and `pprof` shows you a flame graph of exactly where resources are being spent.

Often, the bottlenecks are surprising—like spending 40% of CPU time encoding JSON or allocating memory in a tight loop. Once identified, we might switch to a faster JSON library or use `sync.Pool` to reuse objects. The rule is: Profile first, optimize second.

---

## 87. What is `cgo`?

**Answer:**
`cgo` is the bridge that allows Go code to call C code.

We use it when we need to leverage massive legacy C libraries—like SQLite, OpenGL, or OpenCV—that would be impossible to rewrite in Go from scratch.

However, we generally avoid it if possible. `cgo` breaks Go's safety guarantees, complicates the build process (you need a C compiler installed), and introduces a heavy performance penalty for every function call. It’s a necessary evil, not a default choice.

---

## 88. How do you manage dependencies in Go?

**Answer:**
We use **Go Modules**, the standard system introduced to replace the old GOPATH mess.

Your project has a `go.mod` file that lists exact versions of libraries you use. Go also maintains a `go.sum` file with cryptographic hashes to ensure that the library you download today is bit-for-bit identical to the one you downloaded yesterday.

It protects us from "Dependency Hell" and ensures reproducible builds across any machine or CI/CD pipeline.

---

## 89. What is `go vet`?

**Answer:**
`go vet` is a static analysis tool built right into the standard toolchain.

It doesn't check for style (that's `gofmt` task); it checks for logical bugs that compile but are likely wrong. For example, using a `Printf` with the wrong arguments, or unreachable code, or copying a Mutex (which breaks the lock).

It’s considered best practice to run `go vet` automatically in your CI pipeline. In fact, `go test` runs a subset of vet checks automatically because they are so valuable.

---

## 90. Explain strict typing in Go.

**Answer:**
Go is very rigid about types. It refuses to implicitly convert anything.

An `int` is not an `int64`. A `type UserID int` is not an `int`. You cannot add them together without manually casting one to the other: `int(myID) + x`.

While this creates some boilerplate "casting noise" in the code, it prevents an enormous class of bugs—like accidentally mixing meters and feet, or mixing IDs with counts. It forces you to be explicit about exactly what your data represented.

---

## 91. What is the difference between `os.Exit` and `panic`?

**Answer:**
`os.Exit` is a hard "pull the plug." It tells the Operating System to kill the process immediately with a specific status code. Critically, **defer statements do not run**.

`panic` is a "crash with style." It unwinds the stack, executing all deferred cleanup functions (like flushing logs) before crashing.

We use `os.Exit` in CLI tools—like when the user types invalid arguments. We use `panic` (rarely) for internal logical errors. Generally, `os.Exit` is for the user; `panic` is for the developer.

---

## 92. What are closure functions?

**Answer:**
A closure is an anonymous function that "closes over" or remembers variables from its surrounding scope.

Even if the outer function returns, the closure keeps access to those variables. The Go compiler is smart enough to move those variables to the Heap so they survive.

We use them constantly for callbacks, middleware, or sorting. For example, `sort.Slice` takes a closure where you define the sorting logic using variables from the surrounding function context.

---

## 93. What is the blank identifier `_`?

**Answer:**
The underscore `_` is a "black hole" for values.

Go forces you to use every variable you declare. If a function returns `(result, error)` but you don't care about the error, you can't just ignore it. You must assign it to `_`.

It’s also used in imports: `import _ "image/png"`. This imports the package solely for its side-effects (like registering the PNG decoder) without actually naming the package for use in your code.

---

## 94. How to run unit tests in Go?

**Answer:**
It’s built-in. You don't need JUnit or PyTest. You just write a file ending in `_test.go` and function starting with `TestXxx`.

Then you run `go test ./...`.

The testing style is also unique: Go doesn't really use "Assertions" like `assert.Equals`. Idiomatic Go tests just use plain `if` statements: `if got != expected { t.Errorf(...) }`. This keeps the tests simple, readable, and free of "magic."

---

## 95. What is the `race` detector?

**Answer:**
The Race Detector is a powerful compiler feature enabled with `go run -race`.

It instruments your code to track every memory read and write. If two goroutines access the same memory without a lock, it prints a loud warning and crash report.

It slows down execution by about 10x, so you don't run it in production. But strictly running it during CI/CD tests is arguably the single most important step for maintaining a stable Go codebase.

---

## 96. How to write a benchmark in Go?

**Answer:**
Benchmark functions live in test files but start with `Benchmark`.

They accept a `*testing.B` parameter. You write a standard `for` loop: `for i := 0; i < b.N; i++`. The testing tool automatically adjusts `N`—running your code 1 time, then 100, then 10,000—until it gets a statistically significant measurement of how fast it is.

It’s an incredible tool for optimizing code because it stops you from guessing. You measure, change, measure again.

---

## 97. What is `go generate`?

**Answer:**
`go generate` is a command that scans your source code for specific comments starting with `//go:generate`.

It executes whatever command follows. We use it to automate the creation of boilerplate code—like generating Mocks for testing, or generating Protocol Buffer definitions.

It ensures that the instructions for building the code live *inside* the code itself, rather than in a disconnected README or Makefile.

---

## 98. How does `init()` ordering work with multiple files?

**Answer:**
It follows a strict dependency tree.

If Package A imports Package B, B’s `init()` runs first. Guaranteed.

Within a single package, the order is alphabetical by filename (though the spec says strictly "lexical order"). This ambiguity is risky, so the best practice is to **never rely on init order** within a single package. If file A needs file B initialized, do it explicitly, don't just hope A comes after B in the alphabet.

---

## 99. What are common functional options patterns?

**Answer:**
This is a design pattern used to construct complex objects with optional configuration.

Instead of a constructor like `NewServer(addr, port, timeout, logger, ...)` which is messy and rigid, we write `NewServer(WithPort(80), WithTimeout(5s))`.

Each option is a function that modifies the config. It makes APIs clean, readable, and infinitely extensible without breaking backward compatibility when you add new features later.

---

## 100. How does Go handle stack overflow?

**Answer:**
Go handles stacks dynamically.

In languages like C, you have a fixed stack size. If you recurse too deep, you crash. In Go, stacks start small (2KB). If a function needs more space, the runtime pauses, allocates a bigger stack (usually double), copies your data over, and continues.

This allows safe recursion and massive concurrency. However, there is a hard limit (usually 1GB on 64-bit systems) to prevent infinite loops from eating all RAM, at which point the program will finally panic.
