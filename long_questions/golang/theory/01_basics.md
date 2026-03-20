# 🟢 Go Theory Questions: 1–20 Basics

## 1. What is a Function Literal (Anonymous Function)?

**Answer:**
A function literal is essentially a function defined inline without a name, usually assigned to a variable or passed directly as an argument.

Under the hood, if this function references variables from its surrounding scope, the compiler creates a **Closure**. It allocates a struct on the heap to hold references to those variables, ensuring they survive even after the parent function returns. This allows the function to maintain state.

Practically, we use these constantly for things like **Middleware** (wrapping HTTP handlers) or spawning quick background **goroutines**. The main trade-off to watch for is Loop Variable Capture (in older Go versions), where the closure might accidentally trap the wrong value of a loop index.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Function Literal (Anonymous Function)?

**Your Response:** "A function literal is basically a function defined inline without a name - you usually assign it to a variable or pass it directly as an argument. What's really clever is that if this function references variables from its surrounding scope, the compiler creates a closure, which allocates a struct on the heap to hold those variables so they survive even after the parent function returns.

In practice, I use these constantly for things like middleware in HTTP servers or for spawning quick background goroutines. The main thing to watch out for is loop variable capture in older Go versions, where the closure might accidentally trap the wrong value of a loop index. They're essential for functional programming patterns in Go."

---

## 2. How does the `net/http` package work?

**Answer:**
`net/http` is Go’s production-grade standard library for building robust HTTP clients and servers without needing external frameworks.

For the server side, it relies on a **Handler** pattern. When `ListenAndServe` starts, it accepts TCP connections and spawns a **new goroutine** for every single request. This "one-goroutine-per-request" model is why Go servers scale so effortlessly. On the client side, it uses a `Transport` layer that manages **connection pooling**, keeping TCP connections alive to avoid the overhead of constant handshakes.

In the real world, it’s the backbone of almost every Go microservice. While it’s incredibly stable, the default `http.Client` has **no timeouts**, which is dangerous in production. If an external API hangs, your goroutine hangs forever, eventually causing a resource leak.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the `net/http` package work?

**Your Response:** "The `net/http` package is Go's production-grade standard library for building HTTP clients and servers without external frameworks. On the server side, it uses a handler pattern where `ListenAndServe` accepts TCP connections and spawns a new goroutine for every single request. This one-goroutine-per-request model is why Go servers scale so effortlessly.

On the client side, it uses a transport layer that manages connection pooling to keep TCP connections alive and avoid constant handshakes. It's the backbone of almost every Go microservice. The one thing to be careful about is that the default `http.Client` has no timeouts, which is dangerous in production - if an external API hangs, your goroutine hangs forever and causes resource leaks."

---

## 3. What is Go and who developed it?

**Answer:**
Go is an open-source, statically typed, and compiled language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

It was designed specifically to solve Google-scale engineering problems—mainly slow build times, uncontrolled dependencies, and the difficulty of writing safe concurrent code in C++. It compiles directly to machine code but includes a lightweight **Runtime** to handle garbage collection and concurrency.

Today, it’s effectively the language of the cloud. Tools like **Kubernetes, Docker, and Terraform** are all written in Go. It's perfect for networked backend services, though it does trade off some expressive features (like inheritance or method overloading) in favor of radical simplicity and readability.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Go and who developed it?

**Your Response:** "Go is an open-source statically typed compiled language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was specifically designed to solve Google-scale engineering problems like slow build times, uncontrolled dependencies, and difficulty writing safe concurrent code in C++.

Today it's effectively the language of the cloud - tools like Kubernetes, Docker, and Terraform are all written in Go. It compiles directly to machine code but includes a lightweight runtime for garbage collection and concurrency. You choose Go when you need performance close to C++ but development speed like Python, though it trades some expressive features like inheritance for simplicity and readability."

---

## 4. What are the key features of Go?

**Answer:**
Go is defined by three core philosophies: **Simplicity, Concurrency, and Performance**.

Technically, this manifests in a few key features: **Goroutines** for cheap concurrency (M:N scheduling), a low-latency **Garbage Collector**, and a strict dependency model that ensures fast builds. It produces static binaries with no external dependencies, making deployment trivial—you just copy the file and run it.

Contextually, it’s built for modern multicore distributed systems. You choose Go when you need the performance close to C++ but the development speed of Python. The trade-off is that it's very opinionated—it forces you to handle every error explicitly, which can feel verbose but leads to highly reliable software.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the key features of Go?

**Your Response:** "Go is defined by three core philosophies: Simplicity, Concurrency, and Performance. Technically, this gives us goroutines for cheap concurrency using M:N scheduling, a low-latency garbage collector, and a strict dependency model that ensures fast builds.

One of the best things is that it produces static binaries with no external dependencies - you just copy the file and run it. It's built for modern multicore distributed systems, so you get performance close to C++ but with development speed like Python. The trade-off is that Go is very opinionated - it forces you to handle every error explicitly, which can feel verbose but leads to highly reliable software."

---

## 5. How do you declare a variable in Go?

**Answer:**
In Go, we have two main ways to declare variables, depending on scope and intent.

The first is the `var` keyword (`var x int`), which is explicit and works both globally and locally. It initializes the variable to its **zero value** automatically, so you never have uninitialized memory. The second is the `:=` short declaration (`x := 10`), which infers the type from the value but only works inside functions.

In practice, you’ll use `var` for package-level globals or strict zero-initialization, and `:=` for 90% of your local logic to keep code concise. It's a simple system, though mixing them up can sometimes lead to **variable shadowing** bugs if you aren't careful in nested scopes.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you declare a variable in Go?

**Your Response:** "In Go, we have two main ways to declare variables. The first is the `var` keyword which is explicit and works both globally and locally - it initializes variables to their zero value automatically, so we never have uninitialized memory. The second is the `:=` short declaration which infers the type from the value but only works inside functions.

In practice, I use `var` for package-level globals or when I want strict zero-initialization, and `:=` for about 90% of my local logic to keep code concise. It's a simple system, though you have to be careful about variable shadowing bugs when mixing them in nested scopes."

---

## 6. What are the data types in Go?

**Answer:**
Go’s type system is divided into **Basic** (int, string, bool), **Aggregate** (arrays, structs), **Reference** (slices, maps, channels), and **Interface** types.

The most critical distinction is between Value types and Reference types. A primitive like `int` or a struct is a value—copying it copies the data. A **slice** or **map** is a reference descriptor—copying it just copies the pointer to the underlying data.

Understanding this layout is crucial for performance. Knowing that a **slice** is just a tiny header (pointer + length + cap) means passing it to functions is cheap, whereas passing a huge **array** copies every byte. The system is strict and explicit, which prevents type-conversion errors, though the lack of implicit casting can be verbose.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the data types in Go?

**Your Response:** "Go's type system is divided into Basic types like int, string, bool; Aggregate types like arrays and structs; Reference types like slices, maps, and channels; and Interface types.

The most critical distinction is between value types and reference types. A primitive like int or a struct is a value - copying it copies the data. But a slice or map is a reference descriptor - copying it just copies the pointer to the underlying data. Understanding this is crucial for performance because a slice is just a tiny header, so passing it to functions is cheap, whereas passing a huge array copies every byte. The system is strict and explicit, which prevents type-conversion errors."

---

## 7. What is the zero value in Go?

**Answer:**
The zero value is the default value assigned to any variable that you declare without initializing. Go refuses to leave memory uninitialized or full of random garbage.

Mechanically, the compiler zeroes out the allocated memory. Integers become `0`, booleans `false`, strings `""`, and pointers/slices/maps become `nil`.

This makes code much safer because you don’t have to worry about random crashes from uninitialized variables. It’s particularly useful for things like **counters** or **mutexes**—you can just declare `var mu sync.Mutex` and it’s immediately ready to use (unlocked). The only downside is distinguishing between "User typed 0" and "User didn't type anything," which often requires using pointers.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the zero value in Go?

**Your Response:** "The zero value is the default value that Go assigns to any variable you declare without initializing it. Go refuses to leave memory uninitialized with random garbage. The compiler automatically zeroes out the allocated memory - so integers become 0, booleans become false, strings become empty strings, and pointers, slices, and maps become nil.

This makes code much safer because you don't have to worry about random crashes from uninitialized variables. It's particularly useful for things like counters or mutexes - you can just declare `var mu sync.Mutex` and it's immediately ready to use. The only challenge is distinguishing between when a user actually typed 0 versus when they didn't type anything, which often requires using pointers."

---

## 8. How do you define a constant in Go?

**Answer:**
A constant is a named value that is fixed at compile time and cannot be changed during execution.

Under the hood, constants in Go are unique because they can be **Untyped**. A constant like `const Pi = 3.14` isn't a `float64` yet—it’s an arbitrary-precision number that only gets a type when you use it. This allows you to use `Pi` with both `float32` and `float64` variables without continuously casting it.

We use them for configuration, magic numbers, and enums (`iota`). They are inherently efficient because they don't occupy runtime memory; the compiler essentially copy-pastes the value directly into the machine code instructions wherever it's used.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you define a constant in Go?

**Your Response:** "A constant is a named value that's fixed at compile time and cannot be changed during execution. What's unique about Go constants is that they can be untyped - a constant like `const Pi = 3.14` isn't a float64 yet, it's an arbitrary-precision number that only gets a type when you actually use it.

This allows you to use Pi with both float32 and float64 variables without continuously casting it. We use constants for configuration, magic numbers, and enums with iota. They're inherently efficient because they don't occupy runtime memory - the compiler essentially copy-pastes the value directly into the machine code wherever it's used."

---

## 9. Explain the difference between `var`, `:=`, and `const`.

**Answer:**
These are your three tools for state definition, separated by scope and mutability.

`var` is your standard tool for declaring mutable variables, especially when you want specific zero-values or package-level scope. `:=` is strictly for **local** variables where you want the compiler to infer the type to save typing. `const` creates immutable values that exist only at compile time.

In a real project, you'll see `const` for config keys, `var` for package setup, and `:=` for almost all function logic. The main friction point is usually accidental **shadowing**—re-declaring a variable with `:=` in a nested block that hides the outer one.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the difference between `var`, `:=`, and `const`.

**Your Response:** "These are our three tools for state definition, separated by scope and mutability. `var` is the standard tool for declaring mutable variables, especially when you want specific zero-values or package-level scope. `:=` is strictly for local variables where you want the compiler to infer the type to save typing. `const` creates immutable values that exist only at compile time.

In a real project, you'll see `const` for configuration keys, `var` for package setup, and `:=` for almost all function logic. The main friction point is usually accidental shadowing - re-declaring a variable with `:=` in a nested block that hides the outer one, which can be tricky to debug."

---

## 10. What is the purpose of `init()` function in Go?

**Answer:**
The `init()` function is a special hook that runs automatically before `main()` starts.

Its job is to ensure the package is fully initialized. When your program starts, Go computes the dependency graph and runs `init()` for every imported package, depth-first. This guarantees that by the time your app logic runs, all your dependencies are ready.

Real-world examples include registering database drivers (like `lib/pq`), loading environment variables from a `.env` file, or setting up complex static maps. While useful, it’s somewhat controversial because it introduces **side effects** purely by importing a package, which can make testing and debugging initialization order tricky.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of `init()` function in Go?

**Your Response:** "The `init()` function is a special hook that runs automatically before `main()` starts. Its job is to ensure the package is fully initialized. When your program starts, Go computes the dependency graph and runs `init()` for every imported package in depth-first order, guaranteeing that by the time your app logic runs, all dependencies are ready.

Real-world examples include registering database drivers like `lib/pq`, loading environment variables from a .env file, or setting up complex static maps. While useful, it's somewhat controversial because it introduces side effects purely by importing a package, which can make testing and debugging initialization order tricky."

---

## 11. How do you write a for loop in Go?

**Answer:**
Go effectively removed `while` and `do-while` loops and unified everything under a single `for` keyword.

You have the standard C-style loop for counters (`i++`), the condition-only loop (which acts like a `while`), and the `range` loop for iterating over collections like maps and slices.

This simplifies the language significantly—you don’t have to memorize different loop syntaxes. The `range` loop specifically is the workhorse of Go code, handling everything from iterating database rows to draining channels. The only trade-off is that complex patterns, like a `do-while`, require a slightly verbose `for { ... if break }` structure.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you write a for loop in Go?

**Your Response:** "Go effectively removed while and do-while loops and unified everything under a single `for` keyword. You have the standard C-style loop for counters, the condition-only loop which acts like a while, and the range loop for iterating over collections like maps and slices.

This simplifies the language significantly - you don't have to memorize different loop syntaxes. The range loop is the workhorse of Go code, handling everything from iterating database rows to draining channels. The only trade-off is that complex patterns like a do-while require a slightly verbose `for { ... if break }` structure, but that's a small price for the simplicity."

---

## 12. What is the difference between `break`, `continue`, and `goto`?

**Answer:**
These are your control flow jumps inside loops.

`break` is the emergency exit—it snaps you out of the innermost loop immediately. `continue` is the skip button—it jumps straight to the next iteration. `goto` is the teleporter—it jumps unconditionally to a specific labeled line in your code.

In daily coding, you use `break` for search algorithms and `continue` for filtering data. `goto` is rare and largely discouraged because it leads to spaghetti code, but strictly speaking, it is sometimes used in highly optimized code to handle error cleanup blocks or break out of deeply nested loops where a simple `break` isn't enough.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `break`, `continue`, and `goto`?

**Your Response:** "These are our control flow jumps inside loops. `break` is the emergency exit - it snaps you out of the innermost loop immediately. `continue` is the skip button - it jumps straight to the next iteration. `goto` is the teleporter - it jumps unconditionally to a specific labeled line in your code.

In daily coding, I use `break` for search algorithms and `continue` for filtering data. `goto` is rare and largely discouraged because it leads to spaghetti code, but it's sometimes used in highly optimized code to handle error cleanup blocks or break out of deeply nested loops where a simple `break` isn't enough."

---

## 13. What is a `defer` statement?

**Answer:**
`defer` is a keyword that schedules a function call to run immediately after the surrounding function returns.

When you call `defer`, the arguments are evaluated *right now*, but the function execution is pushed onto a **LIFO stack**. When your function exits—whether it returns normally or panics—that stack unwinds and cleans everything up.

This is critical for resource management. Instead of remembering to close a file at every possible return point (and likely forgetting one), you just open it and immediately write `defer f.Close()`. It keeps your cleanup logic right next to your allocation logic, drastically reducing resource leaks with negligible performance cost.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a `defer` statement?

**Your Response:** "`defer` is a keyword that schedules a function call to run immediately after the surrounding function returns. When you call `defer`, the arguments are evaluated right now, but the function execution is pushed onto a LIFO stack. When your function exits - whether it returns normally or panics - that stack unwinds and cleans everything up.

This is critical for resource management. Instead of remembering to close a file at every possible return point and likely forgetting one, you just open it and immediately write `defer f.Close()`. It keeps your cleanup logic right next to your allocation logic, drastically reducing resource leaks with negligible performance cost."

---

## 14. How does `defer` work with return values?

**Answer:**
This is a nuanced feature: `defer` executes *after* the return value is prepared (assigned) but *before* control is actually passed back to the caller.

This means if you are using **Named Return Values**, a deferred function can actually inspect and **modify** them.

We use this pattern frequently for **Error Wrapping**. We can defer a function that checks for a panic (`recover()`) or inspects the returned error, and wraps it with extra context. It allows you to have a single "exit handler" for your function that standardizes the output, even if the function has 20 different return points.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does `defer` work with return values?

**Your Response:** "This is a nuanced feature: `defer` executes after the return value is prepared but before control is actually passed back to the caller. This means if you're using named return values, a deferred function can actually inspect and modify them.

We use this pattern frequently for error wrapping. We can defer a function that checks for a panic using `recover()` or inspects the returned error, and wraps it with extra context. It allows you to have a single exit handler for your function that standardizes the output, even if the function has 20 different return points. It's a powerful pattern for clean error handling."

---

## 15. What are named return values?

**Answer:**
In a function signature, you can give names to your return parameters, like `func Div(a, b int) (result int, err error)`.

These variables are initialized to their zero values at the start of the function. This allows for **Naked Returns**—you can just write `return` and Go returns whatever values are currently in those variables. It also enables the `defer` modification trick mentioned previously.

I mostly use them for **Documentation**. Seeing `(lat, long float64)` is much clearer than `(float64, float64)`. However, I avoid "Naked Returns" in long functions because they force the reader to scroll up to find where the variable was defined, which hurts readability.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are named return values?

**Your Response:** "In a function signature, you can give names to your return parameters, like `func Div(a, b int) (result int, err error)`. These variables are initialized to their zero values at the start of the function, which allows for naked returns - you can just write `return` and Go returns whatever values are currently in those variables.

It also enables the defer modification trick where deferred functions can modify the return values. I mostly use them for documentation - seeing `(lat, long float64)` is much clearer than `(float64, float64)`. However, I avoid naked returns in long functions because they force the reader to scroll up to find where the variable was defined, which hurts readability."

---

## 16. What are variadic functions?

**Answer:**
A variadic function accepts a variable number of arguments, defined with the `...Type` syntax (like `...int`).

Inside the function, these arguments are bundled into a **slice**. You can pass individual values (`Sum(1, 2, 3)`) or unfold an existing slice (`Sum(nums...)`) to pass it in.

The most famous example is `fmt.Println`. In design patterns, we use this for **Functional Options**: `NewServer(WithPort(80), WithTimeout(5s))`. This allows us to create flexible APIs that accept zero, one, or ten configuration options without forcing the user to pass `nil`s or empty structs for the parameters they don't care about.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are variadic functions?

**Your Response:** "A variadic function accepts a variable number of arguments, defined with the `...Type` syntax like `...int`. Inside the function, these arguments are bundled into a slice. You can pass individual values like `Sum(1, 2, 3)` or unfold an existing slice like `Sum(nums...)` to pass it in.

The most famous example is `fmt.Println`. In design patterns, we use this for functional options like `NewServer(WithPort(80), WithTimeout(5s))`. This allows us to create flexible APIs that accept zero, one, or ten configuration options without forcing the user to pass nils or empty structs for parameters they don't care about."

---

## 17. What is a type alias?

**Answer:**
A type alias, written as `type Alias = Name`, provides a new name for an existing type that is **identical** to the original.

This is different from a Type Definition (`type MyInt int`), which creates a distinct type. With an alias, the compiler treats `Alias` and `Name` as completely interchangeable—you can assign one to the other without casting.

We primarily use this for **Code Migration**. If I move a `User` struct from package `old` to `new`, I can put `type User = new.User` in the old package. This prevents breaking changes for existing users—they can still import the old package, but they are secretly using the new type. Once the migration is done, we delete the alias.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a type alias?

**Your Response:** "A type alias, written as `type Alias = Name`, provides a new name for an existing type that is identical to the original. This is different from a type definition like `type MyInt int`, which creates a distinct type. With an alias, the compiler treats `Alias` and `Name` as completely interchangeable - you can assign one to the other without casting.

We primarily use this for code migration. If I move a User struct from package old to new, I can put `type User = new.User` in the old package. This prevents breaking changes for existing users - they can still import the old package, but they're secretly using the new type. Once the migration is done, we delete the alias."

---

## 18. What is the difference between `new()` and `make()`?

**Answer:**
`new(T)` allocates memory, zeros it, and returns a **pointer** (`*T`). It works for any type (int, struct, etc.).

`make(T)` is strictly for **Maps, Slices, and Channels**. It allocates memory, initializes the internal complex structure (like hash buckets or ring buffers), and returns the **value** (`T`), not a pointer.

In the real world, I use `make` 99% of the time because slices and maps are everywhere. I rarely use `new`—I prefer taking the address of a struct literal (`&User{}`) because it allows me to initialize fields at the same time, whereas `new` leaves everything zeroed.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `new()` and `make()`?

**Your Response:** "`new(T)` allocates memory, zeros it, and returns a pointer to it. It works for any type like int or struct. `make(T)` is strictly for maps, slices, and channels. It allocates memory, initializes the internal complex structure like hash buckets or ring buffers, and returns the value itself, not a pointer.

In the real world, I use `make` 99% of the time because slices and maps are everywhere. I rarely use `new` - I prefer taking the address of a struct literal like `&User{}` because it allows me to initialize fields at the same time, whereas `new` leaves everything zeroed."

---

## 19. How do you handle errors in Go?

**Answer:**
Errors in Go are **Values**, not Exceptions.

Functions typically return a tuple `(result, error)`. The caller must explicitly check `if err != nil`. This makes error handling visible and explicit in the control flow, preventing "hidden" crashes where an exception bubbles up from deep in the stack.

While validly criticized for being verbose (writing `if err != nil` 50 times a day), it leads to highly robust software. You cannot ignore an error by accident; you have to consciously type `_` to ignore it. Modern Go uses `%w` to **wrap** errors, allowing us to preserve the root cause while adding context as the error moves up the stack.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle errors in Go?

**Your Response:** "Errors in Go are values, not exceptions. Functions typically return a tuple of result and error. The caller must explicitly check `if err != nil`. This makes error handling visible and explicit in the control flow, preventing hidden crashes where an exception bubbles up from deep in the stack.

While it's sometimes criticized for being verbose since you write `if err != nil` 50 times a day, it leads to highly robust software. You cannot ignore an error by accident - you have to consciously type `_` to ignore it. Modern Go uses `%w` to wrap errors, allowing us to preserve the root cause while adding context as the error moves up the stack."

---

## 20. What is panic and recover in Go?

**Answer:**
`panic` is Go's version of an exception, but it is reserved strictly for **fatal, unrecoverable** errors—like an index out of bounds or a nil pointer dereference.

`recover` is a built-in function used inside a `defer` block to catch a panic and regain control of the goroutine, stopping it from crashing the entire program.

We basically only use this pair in **Middleware**. If a single HTTP request handler panics, we don't want the whole server to exit. The middleware `recover`s the panic, logs the stack trace, and returns a 500 error to the user. Using panic for regular control flow (like "File Not Found") is considered a major anti-pattern in Go.

---

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is panic and recover in Go?

**Your Response:** "`panic` is Go's version of an exception, but it's reserved strictly for fatal, unrecoverable errors like index out of bounds or nil pointer dereference. `recover` is a built-in function used inside a defer block to catch a panic and regain control of the goroutine, stopping it from crashing the entire program.

We basically only use this pair in middleware. If a single HTTP request handler panics, we don't want the whole server to exit. The middleware recovers the panic, logs the stack trace, and returns a 500 error to the user. Using panic for regular control flow like 'File Not Found' is considered a major anti-pattern in Go."
