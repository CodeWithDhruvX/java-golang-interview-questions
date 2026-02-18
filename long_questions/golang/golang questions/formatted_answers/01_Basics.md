# 🟢 **1–20: Basics**

### 1. What is a function literal (anonymous function)?
"A **function literal** is a function defined without a name. It’s a value that I can assign to a variable or pass as an argument to another function.

Technically, these are **closures**. They capture variables from their surrounding scope by reference. If the outer function returns, the variables used by the closure survive on the heap as long as the closure exists.

I use them constantly for **middleware** in web servers or for `defer` logic. For example, `defer func() { file.Close() }()` is a common pattern to ensure cleanup logic runs with the latest context."

#### Indepth
Go's specific implementation of closures relies on **escape analysis**. The compiler determines if variables captured by the closure need to survive the function call. If they do, they are moved from the stack to the heap. This ensures safety but incurs a small GC pressure.

---

### 2. How does the `net/http` package work?
"The `net/http` package is Go's battle-tested standard library for building HTTP clients and servers.

Its core abstraction is the **Handler interface**, which has a single method: `ServeHTTP`. When a request comes in, the server spawns a new **goroutine** for it and calls this method. It handles keep-alives, timeouts, and protocol parsing (HTTP/1.1 and HTTP/2) automatically.

I love it because it’s production-ready out of the box. Unlike other languages where I need a heavy framework like Express or Spring just to say 'Hello World', Go’s standard library is powerful enough for 90% of my use cases."

#### Indepth
The `DefaultServeMux` used by `http.Handle` is a global variable. In production, avoid using it because any imported package can register a handler and expose endpoints you didn't intend. Always create your own `http.ServeMux`. Also, the default `http.Server` has no timeouts, which makes it vulnerable to Slowloris attacks.

---

### 3. What is Go and who developed it?
"Go is an open-source, statically typed language developed at **Google** by Robert Griesemer, Rob Pike, and Ken Thompson (the creators of Unix and C).

They designed it to solve Google-scale problems: slow compilation times, unreadable codebases, and the difficulty of writing concurrent software. It combines the performance of **C++** with the readability of **Python**.

For me, it’s the 'boring' language in the best way possible. It doesn't have a million features, which means any developer can join a project and understand the code in a day."

#### Indepth
Go maintains the Plan 9 OS philosophy of "everything is a file" and UTF-8 handling (Ken Thompson invented UTF-8). It also strictly enforces **SemVer** and backward compatibility; Go 1 code written in 2012 still compiles today.

---

### 4. What are the key features of Go?
"Go focuses on simplicity and efficiency. Its headline features are **Goroutines** for concurrency, **Channels** for communication, and a **fast Garbage Collector**.

It is statically typed but feels dynamic because of **type inference**. It compiles to a single static binary, making deployment trivial (no 'DLL hell').

I particularly appreciate the tooling. `gofmt`, `go test`, and `go mod` are standard tools that everyone uses, eliminating the 'configuration wars' I see in the JavaScript or Java ecosystems."

#### Indepth
Go's concurrency model is based on **CSP (Communicating Sequential Processes)**, a formal language described by Tony Hoare. Unlike thread-based concurrency where memory is shared, Go encourages "Sharing memory by communicating".

---

### 5. How do you declare a variable in Go?
"I have two main ways: the `var` keyword and the **short declaration** operator `:=`.

`var x int` declares a variable with its zero value. I use this at the package level or when I want to be explicit about the type. `x := 10` declares and initializes it in one step, letting the compiler infer the type.

I use `:=` for 99% of local variables because it keeps the code concise. However, if I need to declare a variable but initialize it later (like inside an `if` block), I stick to `var`."

#### Indepth
Be careful with **variable shadowing**. A `:=` declaration inside an `if` block creates a *new* variable that masks the outer one. This is a common source of bugs `if x, err := func(); ...` where the outer `x` is not updated.

---

### 6. What are the data types in Go?
"Go uses a small, standard set of types.

We have **Basic types** (bool, string, int, float64), **Aggregate types** (arrays, structs), and **Reference types** (slices, maps, channels, pointers, functions).

There is no 'Class' type. We build complex data structures using **structs**. I find this refreshing because it forces me to separate data (structs) from behavior (methods), avoiding the deep inheritance hierarchies of OOP."

#### Indepth
Go structs have flattened memory layouts. Fields are laid out sequentially in memory, with padding added for alignment. This contrasts with Java or Python where objects are references scattered in the heap. This data locality is a major factor in Go's performance.

---

### 7. What is the zero value in Go?
"The **zero value** is the default value assigned to variables that are declared but not initialized. Go never leaves memory uninitialized.

For `int` it is `0`, `bool` is `false`, `string` is `""`. For reference types like pointers, slices, and maps, it is `nil`.

I rely on this heavily. For example, a `sync.Mutex` is ready to use immediately because its zero value is an 'unlocked mutex'. This 'make the zero value useful' philosophy is a key part of idiomatic Go API design."

#### Indepth
While standard types are safe, be careful with `nil` Maps. You can read from a `nil` map (it returns zero values), but writing to a `nil` map causes a panic. Slices are safer; you can append to a `nil` slice perfectly fine.

---

### 8. How do you define a constant in Go?
"I use the `const` keyword.

Constants are evaluated at **compile time**. They can be typed (`const X int = 10`) or untyped (`const X = 10`). Untyped constants are powerful because they have arbitrary precision and can be used in any numeric context without casting.

I often use the `iota` identifier inside a `const` block to create enums. `const ( A = iota; B )` assigns 0 to A and 1 to B automatically."

#### Indepth
Go constants are represented with at least 256 bits of precision during compilation. This allows you to perform high-precision math with constants (like `math.Pi`) and only round to a float/int at the final assignment step.

---

### 9. Explain the difference between `var`, `:=`, and `const`.
"`var` declares a variable. It works at both package and function levels.
`:=` is short variable declaration. It *only* works inside functions and infers the type.
`const` declares a value that cannot change and is computed at compile time.

I use `const` for magic numbers and config strings. I use `:=` for almost all local logic. I reserve `var` for package-level globals (which I try to avoid) or zero-value initialization."

#### Indepth
You can use `:=` to redeclare a variable *if* it appears in a multi-variable declaration where at least one variable is new. `err := fn1(); res, err := fn2()` is valid because `res` is new, and `err` is just assigned to.

---

### 10. What is the purpose of `init()` function in Go?
"`init()` is a special function that runs automatically before `main()` starts.

It is used to set up package-level state, like parsing configuration flags or registering database drivers (like `postgres`). A package can have multiple `init` functions, and they run in the order they appear.

I use it sparingly because strict ordering dependencies can be hard to debug. If possible, I prefer explicit `New()` or `Setup()` functions so the caller has control over *when* initialization happens."

#### Indepth
Package initialization order is determined by dependency graph. Imports are initialized first. Within a package, variables are initialized before `init()` functions. This deterministic order is crucial for system reliability.

---

### 11. How do you write a for loop in Go?
"Go only has one looping keyword: `for`.

It can be a C-style loop: `for i:=0; i<10; i++`.
It can be a while loop: `for x < 100`.
It can be an infinite loop: `for { }`.

My favorite is the **range** clause: `for k, v := range items`. It iterates over slices, maps, strings, and channels. It handles the index/key and value extraction cleanly, avoiding off-by-one errors."

#### Indepth
Prior to Go 1.22, the loop variable was reused across iterations, referencing the same memory address. This caused bugs when capturing the variable in a goroutine. In modern Go (1.22+), loop variables have per-iteration scope, fixing this "most common Go mistake".

---

### 12. What is the difference between `break`, `continue`, and `goto`?
"`break` exits the innermost loop or switch statement immediately.
`continue` skips the rest of the current iteration and jumps to the next one.
`goto` jumps execution to a labelled line.

I use `break` and `continue` constantly. I almost *never* use `goto`, except perhaps to break out of a deeply nested loop (`break Label` does this better) or in generated code."

#### Indepth
Labels are block-scoped and must be defined in the same function. They are most useful when you have a `switch` inside a `for` loop and want to break the *loop*, not the switch. calling `break` alone would only exit the switch.

---

### 13. What is a `defer` statement?
"`defer` schedules a function call to run immediately *after* the surrounding function returns.

The arguments are evaluated *when the defer is declared*, but the function execution happens at the end. It follows a LIFO (Last-In-First-Out) order.

This is my favorite error-handling feature. It ensures that `file.Close()` happens whether the function returns normally or panics. It puts the cleanup logic right next to the allocation logic, making leaks much less likely."

#### Indepth
Deferred calls are pushed onto a stack. When the function exits, they are popped and executed. The overhead of `defer` used to be significant (~50ns), but recent Go versions (1.14+) use "open-coded defers" which compile them inline, making them nearly zero-cost.

---

### 14. How does `defer` work with return values?
"Deferred functions run **after** the `return` statement has set the return values, but **before** the function actually exits.

This means a deferred function can inspect and even **modify** named return values.

I use this pattern to handle panics in a unified way. `defer func() { if r := recover(); r != nil { err = fmt.Errorf("panic: %v", r) } }()` allows me to turn a crash into a structured error return."

#### Indepth
This pattern is the only way to modify a return value *after* the `return` statement has executed. It is widely used for **error wrapping**, where you want to add context to an error returned by any point in the function.

---

### 15. What are named return values?
"Go allows me to give names to return parameters in the function signature: `func div(a, b int) (res int, err error)`.

These variables are initialized to their zero values. I can return them simply by typing `return` (called a 'naked return').

I use them for documentation clarity. `func (lat, long float64)` is clearer than `func (float64, float64)`. However, I avoid naked returns in long functions because they hurt readability—I prefer explicit `return res, err`."

#### Indepth
Behind the scenes, named return values are just local variables declared at the top of the function. **Naked returns** should be used sparingly; if the function is longer than a screen, explicit returns improve readability significantly.

---

### 16. What are variadic functions?
"A variadic function accepts a variable number of arguments. `fmt.Println` is the most famous example.

I define it using `...Type` as the last parameter, like `func sum(nums ...int)`. Inside the function, `nums` is just a slice `[]int`.

I use this often for 'Option' patterns in constructors, like `NewServer(opts ...Option)`. It allows the caller to provide zero, one, or ten configuration options without changing the function signature."

#### Indepth
When you pass a slice to a variadic function like `fn(slice...)`, Go does *not* create a new slice. It passes the existing slice header. This means modifying elements in the variadic function will modify the original slice's backing array.

---

### 17. What is a type alias?
"A **type alias** (`type MyInt = int`) creates an alternative name for the *same* type. They are interchangeable.

This is distinct from a **type definition** (`type MyInt int`), which creates a *new*, distinct type that requires casting.

I rarely use aliases restrictedly. Their primary use case is during **refactoring**: moving a type from Package A to Package B, but leaving an alias in Package A so I don't break existing code that imports it."

#### Indepth
Type aliases were introduced efficiently in Go 1.9 to aid strictly in **gradual code repair**. They allow large codebases to move a type between packages (e.g., `context.Context` moving from `golang.org/x/net/context` to `std`) without breaking consumers.

---

### 18. What is the difference between `new()` and `make()`?
"`new(T)` allocates zeroed memory for a type and returns a **pointer** (`*T`). It works for any type (int, struct, etc.).

`make(T)` is strictly for **slices, maps, and channels**. It initializes the internal data structure (like the hash bucket headers) and returns the **value** (`T`), not a pointer.

If I tried `new(map[string]int)`, I'd get a pointer to a nil map. Writing to it would panic. So I always use `make()` for reference types and `new()` (or `&Struct{}`) for value types."

#### Indepth
The difference exists because slices, maps, and channels are complex structures that must be initialized before use (e.g., creating the hash map buckets). `new()` only zeroes memory, which isn't enough for these types. `make()` performs that necessary setup.

---

### 19. How do you handle errors in Go?
"In Go, errors are **values**, usually returned as the last value of a function. There are no exceptions.

I check `if err != nil` immediately after a call.

This explicit handling forces me to think about failure modes at every step. I can't just 'hope' a higher-level catch block handles it. It makes Go code verbose but incredibly robust and readable—I can see exactly where the control flow goes."

#### Indepth
Go 1.13 introduced error wrapping. Use `errors.Is(err, target)` instead of `err == target` to check if an error wraps a specific sentinel value. Use `errors.As(err, &target)` to check if it wraps a specific type.

---

### 20. What is panic and recover in Go?
"**Panic** ceases normal execution. It runs any deferred functions and then crashes the program. It should only be used for unrecoverable errors (like a corrupted config at startup).

**Recover** is a built-in function that regains control of a panicking goroutine. It *only* works inside a `defer` function.

I treat `panic/recover` like an assertion mechanism, not for control flow. The only place I use `recover` is in the top-level HTTP middleware to ensure a single bad request implies a 500 error rather than crashing the entire server."

#### Indepth
Panic unwinds the stack, running deferred functions. If not recovered, it crashes the process with a non-zero exit code. `recover()` returns `nil` if called normally (not during a panic), so it effectively does nothing outside a defer.

---

### 21. What are blank identifiers in Go?
"The **blank identifier** (`_`) is a write-only variable placeholder.

Go refuses to compile unused variables. If a function returns `(val, err)` and I only care about the error, I must write `_, err := fn()`.

I also use it for **import side-effects**: `import _ "image/png"` registers the PNG decoder without exposing the package name. It’s a clean way to satisfy the compiler while ignoring data I don't need."

#### Indepth
Identified by `_`, this is a language mechanism to discard values. It is effectively a "black hole" for data. Since Go requires every declared variable to be used, `_` is necessary for ignoring return values you don't care about.
