# Golang Interview Questions & Answers

## 🟢 **1–20: Basics**

### **1. What is a function literal (anonymous function)?**
A function without a name, often assigned to a variable or passed as an argument.
*   **Why use it?** For short, throwaway logic (like callbacks) or closures (capturing variables).
*   **Code**:
    ```go
    func() { fmt.Println("I am anonymous") }() // Immediately invoked
    
    add := func(a, b int) int { return a + b } // Assigned to variable
    ```

### **2. How does the `net/http` package work?**
It provides HTTP client and server implementations.
*   **Server**: Listens on a port (`http.ListenAndServe`) and routes requests to handlers (`http.Handle`).
*   **Client**: Sends requests (`http.Get`, `http.Post`).
*   **Key Concept**: "Everything is a Handler" (implements `ServeHTTP`).

### **3. What is Go and who developed it?**
Go (Golang) is an open-source, statically typed, compiled language developed by **Google** (Robert Griesemer, Rob Pike, Ken Thompson).
*   **Goal**: Simplicity, high performance, and built-in concurrency.

### **4. What are the key features of Go?**
*   **Simplicity**: Minimal keywords (~25).
*   **Concurrency**: Goroutines and Channels (CSP model).
*   **Performance**: Compiles to machine code, fast GC.
*   **Safety**: Memory safe, strong typing.
*   **Tools**: Built-in formatting (`fmt`), testing (`test`), and dependency management (`mod`).

### **5. How do you declare a variable in Go?**
Three main ways:
1.  **`var` (Explicit)**: `var x int = 10`
2.  **`var` (Inferred)**: `var x = 10`
3.  **Short Declaration (Inside functions only)**: `x := 10`

### **6. What are the data types in Go?**
*   **Basic**: `bool`, `string`, `int`, `uint`, `float64`, `complex128`, `byte` (uint8), `rune` (int32).
*   **Aggregate**: `array`, `struct`.
*   **Reference**: `slice`, `map`, `channel`, `pointer`, `function`.
*   **Interface**: `interface{}`.

### **7. What is the zero value in Go?**
The default value assigned to variables if not initialized manually.
*   `int` / `float` → `0`
*   `bool` → `false`
*   `string` → `""` (empty string)
*   `pointer`, `slice`, `map`, `channel`, `interface`, `func` → `nil`

### **8. How do you define a constant in Go?**
Using the `const` keyword. Value must be known at compile time.
```go
const Pi = 3.14
const (
    StatusOk = 200
    StatusErr = 500
)
```

### **9. Explain the difference between `var`, `:=`, and `const`.**
*   **`var`**: Can be used anywhere (package or function level). Can declare without value (gets zero val).
*   **`:=`**: Short declaration. **Inside functions only**. Must provide value immediately.
*   **`const`**: Immutable value. Cannot change after definition.

### **10. What is the purpose of `init()` function in Go?**
A special function that runs **before** `main()`.
*   **Usage**: Initializing global variables, registering drivers (e.g., database drivers), verifying environment.
*   **Note**: You can have multiple `init()` functions; they run in order of file name/declaration.

### **11. How do you write a for loop in Go?**
Go only has `for`. No `while` or `do-while`.
1.  **Standard**: `for i := 0; i < 10; i++ { ... }`
2.  **While-like**: `for i < 10 { ... }`
3.  **Infinite**: `for { ... }`
4.  **Range**: `for index, value := range collection { ... }`

### **12. What is the difference between `break`, `continue`, and `goto`?**
*   **`break`**: Exits the loop *immediately*.
*   **`continue`**: Skips the *current iteration* and jumps to the next condition check.
*   **`goto`**: Jumps to a specific labeled line (rarely used, can make code "spaghetti").

### **13. What is a `defer` statement?**
Schedules a function call to run **after** the surrounding function returns.
*   **Usage**: Cleanup (close file, unlock mutex, close DB connection).
*   **Order**: LIFO (Last In, First Out) - like a stack.
    ```go
    defer fmt.Println("World")
    fmt.Println("Hello")
    // Output: Hello World
    ```

### **14. How does `defer` work with return values?**
`defer` executes **after** the return value is calculated but **before** it is actually passed back.
*   It can modify **named return values**.
    ```go
    func double() (result int) { // Named return 'result'
        defer func() { result *= 2 }()
        return 5 // result becomes 5, then defer runs (5*2), returns 10
    }
    ```

### **15. What are named return values?**
You can give names to return parameters in the function signature. They are treated as variables defined at the top of the function.
*   **Benefit**: Documentation (clarity) and allows "naked return" (just `return` without arguments).

### **16. What are variadic functions?**
Functions that accept a variable number of arguments (like `fmt.Println`).
*   **Syntax**: `...Type` as the last parameter.
*   **Under hood**: Received as a slice `[]Type`.
    ```go
    func sum(nums ...int) int { ... }
    sum(1, 2)
    sum(1, 2, 3, 4)
    ```

### **17. What is a type alias?**
Creates a new name for an existing type.
*   **Syntax**: `type NewName = ExistingType` (note the `=`).
*   **Usage**: Refactoring, compatibility (e.g., `byte` is alias for `uint8`, `rune` for `int32`).
*   **VS Type Definition**: `type MyInt int` creates a *distinct* new type; alias `type MyInt = int` is the *same* type.

### **18. What is the difference between `new()` and `make()`?**
*   **`new(T)`**: Allocates memory for T, zeroes it, returns a **pointer** (`*T`). Use for structs, primitives.
*   **`make(T)`**: Allocates and *initializes* hidden internal structure. Returns the **value** (`T`). Only for **slices, maps, and channels**.

### **19. How do you handle errors in Go?**
Explicitly. Functions return an error as the last value.
*   **Pattern**: Check `if err != nil`.
    ```go
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    ```
*   **Philosophy**: Errors are values, not exceptions.

### **20. What is panic and recover in Go?**
*   **Panic**: Stops normal execution, runs deferred functions, and crashes program (like an un-caught exception). rarely used in app logic.
*   **Recover**: Regains control of a panicking goroutine. **Must be called inside a `defer`**.
    ```go
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("boom")
    ```

### **21. What are blank identifiers in Go?**
The underscore `_` character. It is a write-only variable used to ignore values.
*   **Usage**: Ignoring return values, importing packages for side effects (init functions), or silencing "unused variable" errors.

## 🟡 **21–40: Arrays, Slices, and Maps**

### **22. What is the difference between an array and a slice?**
*   **Array**: Fixed size. Value type. `[3]int`. Size is part of the type (`[3]int` != `[4]int`).
*   **Slice**: Dynamic size. Reference type (view of an array). `[]int`.
*   **Key**: Use slices 99% of the time. Arrays are for specific low-level memory layouts.

### **23. How do you append to a slice?**
Using the built-in `append()` function.
*   It returns a **new slice value** (updates length/capacity).
    ```go
    s := []int{1, 2}
    s = append(s, 3, 4) // s is now [1, 2, 3, 4]
    ```

### **24. What happens when a slice is appended beyond its capacity?**
Go allocates a **new, larger underlying array** (usually double the size), copies data over, and points the slice to the new array.
*   **Performance Note**: Frequent resizing is expensive; pre-allocate with `make` if size is known.

### **25. How do you copy slices?**
Using the built-in `copy(dst, src)` function.
*   It copies the minimum of `len(dst)` and `len(src)` elements.
*   It does **not** resize `dst`.

### **26. What is the difference between len() and cap()?**
*   **`len()`**: Number of elements *currently* in the slice.
*   **`cap()`**: Number of elements the slice *can hold* before reallocating (from its pointer to the end of the underlying array).

### **27. How do you create a multi-dimensional slice?**
Slice of slices.
```go
matrix := [][]int{
    {1, 2},
    {3, 4},
}
// Or dynamic
rows := 3
grid := make([][]int, rows)
```

### **28. How are slices passed to functions (by value or reference)?**
Strictly **by value**, but the value is a small struct containing a pointer to the array.
*   **Effect**: Modifying elements INSIDE the function **changes** the original slice.
*   **But**: `append` inside the function won't change the original slice's length in the caller unless you return the new slice.

### **29. What are maps in Go?**
Unordered collection of key-value pairs (Hash Table).
*   **Syntax**: `map[KeyType]ValueType`.
*   **Reference Type**: Zero value is `nil`. Writing to a nil map causes a **panic**. Always use `make` or literal.

### **30. How do you check if a key exists in a map?**
Comma-ok idiom.
```go
val, ok := myMap["key"]
if ok {
    fmt.Println("Exists:", val)
} else {
    fmt.Println("Does not exist")
}
```

### **31. Can maps be compared directly?**
No. Maps can only be compared to `nil`.
*   To compare contents, you must iterate and compare keys/values manually or use `reflect.DeepEqual`.

### **32. What happens if you delete a key from a map that doesn’t exist?**
Nothing. It's a no-op. No error, no panic.

### **33. Can slices be used as map keys?**
**No**. Map keys must be **comparable** (support `==`). Slices, maps, and functions are not comparable.
*   **Workaround**: Use a string key (convert slice to string) or a struct (if fields are comparable).

### **34. How do you iterate over a map?**
Using `range`.
*   **Order**: **Random** (randomized by Go runtime to prevent relying on order).
    ```go
    for k, v := range myMap { ... }
    ```

### **35. How do you sort a map by key or value?**
Maps are unsorted.
1.  Extract keys to a slice.
2.  Sort the slice (`sort.Strings` or `slices.Sort`).
3.  Iterate the sorted slice and lookup values in the map.

### **36. What are struct types in Go?**
A collection of fields. Used to define custom data types.
```go
type Person struct {
    Name string
    Age  int
}
```

### **37. How do you define and use struct tags?**
String literals attached to fields, read via reflection.
*   **Usage**: JSON serialization, database mapping.
    ```go
    type User struct {
        ID   int    `json:"id" db:"user_id"`
        Name string `json:"name,omitempty"`
    }
    ```

### **38. How to embed one struct into another?**
Anonymous field (Composition).
*   **Effect**: Inner struct's fields/methods are promoted to the outer struct.
    ```go
    type Base struct { id int }
    type User struct {
        Base // Embedding
        Name string
    }
    u := User{}
    u.id = 1 // Access directly
    ```

### **39. How do you compare two structs?**
*   **Comparable**: If all their fields are comparable, you can use `==`.
*   **Not Comparable**: If they contain slices/maps, you must use `reflect.DeepEqual`.

### **40. What is the difference between shallow and deep copy in structs?**
*   **Assignment (`a := b`)**: Shallow copy. Copies field values. If a field is a pointer/slice, both structs point to the same data.
*   **Deep Copy**: Recursively copying all data (dereferencing pointers). Must be done manually or with libraries if the struct has reference types.

### **41. How do you convert a struct to JSON?**
Using the `encoding/json` package.
*   **Marshaling**: `json.Marshal(v)` converts Go struct -> JSON bytes.
*   **Note**: Only **exported fields** (Capitalized) are encoded.

## 🔵 **41–60: Pointers, Interfaces, and Methods**

### **42. What are pointers in Go?**
Variables that store the **memory address** of another variable.
*   **Symbol**: `*T` (type), `&` (address of), `*` (dereference).
*   **No Arithmetic**: Go does NOT support pointer arithmetic (like C++), for safety.

### **43. How do you declare and use pointers?**
```go
x := 10
p := &x         // p holds address of x
fmt.Println(*p) // Dereference: prints 10
*p = 20         // Change value at address
```

### **44. What is the difference between pointer and value receivers?**
*   **Value Receiver** (`func (s MyStruct)`): Receives a **copy**. Changes inside method **do not** affect the original.
*   **Pointer Receiver** (`func (s *MyStruct)`): Receives a **pointer**. Changes **affect** the original. Use this for modifying state or avoiding copy overhead for large structs.

### **45. What are methods in Go?**
Functions attached to a specific type (the "receiver").
```go
func (u User) Greet() string { return "Hi " + u.Name }
```

### **46. How to define an interface?**
A set of method signatures.
```go
type Shape interface {
    Area() float64
}
```

### **47. What is the empty interface in Go?**
`interface{}`. It creates a variable that can hold **values of any type** (because every type implements at least zero methods).
*   **Usage**: Generalized containers, working with unknown data (like `fmt.Println` arguments).

### **48. How do you perform type assertion?**
To extract the concrete value from an interface.
```go
val, ok := i.(string) // Assert that 'i' holds a string
if ok {
    fmt.Println(val)
}
```

### **49. How to check if a type implements an interface?**
This is usually checked at compile time.
*   **Explicit Check**: `var _ MyInterface = (*MyType)(nil)`

### **50. Can interfaces be embedded?**
Yes. An interface can include other interfaces.
```go
type ReadWriter interface {
    Reader
    Writer
}
```

### **51. What is polymorphism in Go?**
Achieved via **Interfaces**. A variable of an interface type can hold values of any concrete type that implements methods of that interface.
*   **Code**: `func PrintArea(s Shape)` can accept `Circle`, `Rectangle`, etc.

### **52. How to use interfaces to write mockable code?**
Accept interfaces in your functions instead of concrete structs.
*   **Test**: In tests, pass a `MockStruct` that implements the interface but returns stubbed data.

### **53. What is the difference between `interface{}` and `any`?**
None. `any` is an alias for `interface{}` introduced in Go 1.18.

### **54. What is duck typing?**
"If it walks like a duck and quacks like a duck, it's a duck."
*   Go interfaces are **implicit**. You don't say `implements Shape`; you just define the `Area()` method, and Go automatically knows it implements `Shape`.

### **55. Can you create an interface with no methods?**
Yes, that's `interface{}` (the empty interface).

### **56. Can structs implement multiple interfaces?**
Yes, indistinguishably. If `Dog` has `Walk()` and `Bark()`, it implements both `Walker` and `Barker` interfaces automatically.

### **57. What is the difference between concrete type and interface type?**
*   **Concrete**: Describes the exact memory layout (e.g., `int`, `struct`).
*   **Interface**: Describes behavior (method set). It's abstract.

### **58. How to handle nil interfaces?**
An interface value is `nil` only if **both** its type and value are `nil`.
*   **Trap**: A non-nil pointer inside an interface makes the interface itself **non-nil**.
    ```go
    var p *int = nil
    var i interface{} = p
    fmt.Println(i == nil) // False! Because i holds type *int
    ```

### **59. What are method sets?**
The set of methods attached to a type.
*   **T**: Methods with value receivers.
*   ***T**: Methods with both value AND pointer receivers.

### **60. Can a pointer implement an interface?**
Yes. In fact, if a method requires a pointer receiver, ONLY the pointer type (`*T`) implements the interface, not the value type (`T`).

### **61. What is the use of `reflect` package?**
It allows inspecting type information at runtime.
*   **Usage**: ORMs, JSON serialization/deserialization, dependency injection.
*   **Cons**: Slow, complex, no type safety. Avoid using unless necessary.

## 🟣 **61–80: Concurrency and Goroutines**

### **62. What are goroutines?**
Lightweight threads managed by the Go runtime (not OS threads).
*   **Cost**: Very cheap (~2KB stack memory).
*   **Usage**: `go funcName()`

### **63. How do you start a goroutine?**
Use the `go` keyword before a function call.
```go
go doSomething()
```

### **64. What is a channel in Go?**
A communication mechanism that allows goroutines to exchange data and synchronize execution.
*   **Mnemonic**: "Don't communicate by sharing memory; share memory by communicating."

### **65. What is the difference between buffered and unbuffered channels?**
*   **Unbuffered** (`make(chan int)`): Synchronous. Sender blocks until receiver is ready.
*   **Buffered** (`make(chan int, 5)`): Asynchronous (up to limit). Sender blocks only when buffer is full.

### **66. How do you close a channel?**
`close(ch)`
*   **Rule**: close from the **sender** side. Closing from receiver or closing a closed channel causes panic.

### **67. What happens when you send to a closed channel?**
**Panic**. Do not send to closed channels.

### **68. How to detect a closed channel while receiving?**
Comma-ok idiom.
```go
val, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}
```

### **69. What is the `select` statement in Go?**
Like a `switch` statement, but for channels. It blocks until one of its cases (send or receive) is ready.
*   If multiple are ready, it picks one **randomly**.

### **70. How do you implement timeouts with `select`?**
Using `time.After()`.
```go
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timed out")
}
```

### **71. What is a `sync.WaitGroup`?**
It waits for a collection of goroutines to finish.
*   `Add(n)`: Increment counter.
*   `Done()`: Decrement counter (defer this).
*   `Wait()`: Block until counter is zero.

### **72. How does `sync.Mutex` work?**
A mutual exclusion lock. Ensures only one goroutine accesses a critical section at a time.
*   `Lock()` and `Unlock()`.

### **73. What is `sync.Once`?**
Ensures a function runs **exactly once**, even if called from multiple goroutines.
*   **Usage**: Singleton initialization.

### **74. How do you avoid race conditions?**
1.  Use **Channels** to coordinate.
2.  Use **Mutexes** to lock shared data.
3.  Use **atomics** (`sync/atomic`) for counters.
4.  Run tests with `go test -race`.

### **75. What is the Go memory model?**
It specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes in another goroutine.
*   **Core**: Happens-before relationship (e.g., channel send happens before receive).

### **76. How do you use `context.Context` for cancellation?**
To pass cancellation signals and deadlines across API boundaries.
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Call to stop children

go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        return // Stop work
    }
}(ctx)
```

### **77. How to pass data between goroutines?**
**Channels** are the primary and idiomatic way.

### **78. What is the `runtime.GOMAXPROCS()` function?**
Controls the number of OS threads that can execute Go code simultaneously.
*   **Default**: Number of logical CPUs.

### **79. How do you detect deadlocks in Go?**
*   **Runtime**: If all goroutines are asleep, the runtime panics: `fatal error: all goroutines are asleep - deadlock!`.
*   **Tools**: `pprof` (goroutine profile) to analyze stuck routines.

### **80. What are worker pools and how do you implement them?**
A pattern to limit concurrency.
1.  Create a fixed number of goroutines (workers).
2.  All workers listen on a shared `jobs` channel.
3.  Send tasks to `jobs` channel.

### **81. How to write concurrent-safe data structures?**
Embed a `sync.Mutex` or `sync.RWMutex` in the struct and lock it during every read/write.
```go
type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}
```

## 🔴 **81–100: Advanced & Best Practices**

### **82. How does Go handle memory management?**
It uses a **Garbage Collector (GC)**.
*   **Stack**: Fast, automatic, LIFO (local variables which don't escape).
*   **Heap**: Slower, GC-managed (shared variables, pointers that escape).

### **83. What is garbage collection in Go?**
A non-generational, concurrent, tri-color mark-and-sweep collector.
*   **Goal**: Low latency (pauses < 1ms on average). Not fully real-time but very fast.

### **84. How do you profile CPU and memory in Go?**
Using `pprof` tool.
*   **Command**: `go tool pprof http://localhost:6060/debug/pprof/profile`
*   **Flag**: `go test -cpuprofile cpu.out`

### **85. What is the difference between compile-time and runtime errors?**
*   **Compile-time**: Syntax errors, type mismatches. Code won't build.
*   **Runtime**: Index out of bounds, nil pointer dereference, panic. Happens while app is running.

### **86. How to use `go test` for unit testing?**
Run `go test ./...` in the module root.
*   Files must end in `_test.go`.
*   Functions must look like `func TestName(t *testing.T)`.

### **87. What is table-driven testing in Go?**
A pattern where test inputs and expected outputs are defined in a slice of structs (table), and a loop runs the test logic for each entry.
*   **Benefit**: Easy to add new test cases without duplicating code.

### **88. How to benchmark code in Go?**
Create a function starting with `Benchmark` in a `_test.go` file.
```go
func BenchmarkMyFunc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MyFunc()
    }
}
```
*   Run with `go test -bench=.`

### **89. What is `go mod` and how does it work?**
The dependency management system.
*   `go.mod`: Defines module path and dependencies.
*   `go.sum`: Checksums for security validation.
*   **Command**: `go mod tidy` cleans up dependencies.

### **90. What is vendoring in Go modules?**
Copying all dependencies into a local `vendor` folder.
*   **Command**: `go mod vendor`.
*   **Why**: Builds without internet; ensure exact code snapshot.

### **91. How to handle versioning in modules?**
Semantic Versioning (v1.2.3).
*   **Major version bump (v2+)**: Requires changing the module path (e.g., `github.com/my/app/v2`).

### **92. How do you structure a Go project?**
Standard Go Project Layout (unofficial but common):
*   `/cmd`: Main applications (entry points).
*   `/pkg`: Library code usable by external apps.
*   `/internal`: Private library code (enforced by compiler).
*   `/api`: API definitions (protobufs, swagger).

### **93. What is the idiomatic way to name Go packages?**
Short, lowercase, singular, and descriptive.
*   Good: `http`, `json`, `user`.
*   Bad: `Users`, `httpUtil`, `go_common`.

### **94. What is the purpose of the `internal` package?**
The Go compiler forbids importing packages inside an `internal` directory unless the importer is within the same parent tree.
*   **Effect**: Enforces privacy/encapsulation across modules.

### **95. How do you handle logging in Go?**
*   **Standard**: `log` package (simple, no levels).
*   **Structured (New in 1.21)**: `log/slog`.
*   **Third-party**: `zap`, `zerolog` (for high performance).

### **96. What is the difference between `log.Fatal`, `log.Panic`, and `log.Println`?**
*   `Println`: Logs and continues.
*   `Fatal`: Logs and calls `os.Exit(1)`. **Defers do NOT run.**
*   `Panic`: Logs and calls `panic()`. Defers run.

### **97. What are build tags in Go?**
Comments that determine whether a file should be included in the package build.
*   `//go:build linux,386`
*   **Usage**: OS-specific files, separating integration tests.

### **98. What are cgo and its use cases?**
Mechanism to call C code from Go.
*   **Usage**: Interfacing with legacy C libraries (e.g., SQLite, Graphics).
*   **Cons**: Breaks cross-compilation simplicity, slower function calls, harder debugging.

### **99. What are some common Go anti-patterns?**
*   Returning interface types (return concrete types, accept interfaces).
*   Using `init()` for complex setup (hard to test).
*   Ignoring errors with `_`.
*   Unchecked type assertions.
*   Package level global variables.

### **100. What are Go code quality tools (lint, vet, staticcheck)?**
*   `go vet`: Checks for suspicious constructs (built-in).
*   `staticcheck`: Advanced linter.
*   `golangci-lint`: A runner that aggregates dozens of linters. Use this in CI.

### **101. What are the best practices for writing idiomatic Go code?**
Follow "Effective Go":
*   Use `gofmt`.
*   Handle errors explicitly.
*   Use short variable names (`i`, `err`, `ctx`).
*   Prefer composition over inheritance.
*   Keep interfaces small.

## 🟢 **101–120: Project Structure & Design Patterns**

### **102. How do you organize a large-scale Go project?**
Focus on **domain-driven design** rather than layer-driven.
*   Group by feature (e.g., `user/`, `order/`).
*   Use `internal/` to hide private code.
*   Keep `main.go` small (just bootstrapping).

### **103. What is the standard Go project layout?**
While not official, the community standard (golang-standards/project-layout) is:
*   `/cmd`: Entry points.
*   `/internal`: Private application code.
*   `/pkg`: Public library code.
*   `/api`: API specs.
*   `/configs`: Configuration files.

### **104. What is the `cmd` directory used for in Go?**
It contains the `main` packages. Each subdirectory is a binary.
*   Example: `cmd/server/main.go`, `cmd/cli/main.go`.

### **105. How do you structure code for reusable packages?**
*   Put them in the root or `/pkg`.
*   Avoid specific dependencies (e.g., specific DB drivers).
*   Accept interfaces to allow swapping implementations.

### **106. What are Go's most used design patterns?**
*   **Functional Options**: For flexible configuration.
*   **Factory**: Using functions like `NewUser()`.
*   **Singleton**: Using `sync.Once`.
*   **Adapter**: Using interfaces.

### **107. Explain the Factory Pattern in Go.**
It’s just a function that returns a struct or interface.
```go
func NewLogger(level string) Logger {
    return &SimpleLogger{level: level}
}
```

### **108. How to implement Singleton Pattern in Go?**
Use `sync.Once`.
```go
var instance *Manager
var once sync.Once

func GetManager() *Manager {
    once.Do(func() {
        instance = &Manager{}
    })
    return instance
}
```

### **109. What is Dependency Injection in Go?**
Passing dependencies (like DB, Logger) into constructors explicitly.
```go
type Service struct {
    db *sql.DB
}
func NewService(db *sql.DB) *Service {
    return &Service{db: db}
}
```

### **110. What is the difference between composition and inheritance in Go?**
Go has **no inheritance**. It uses **composition** (embedding).
*   **Inheritance**: "is-a" relationship (Dog is an Animal).
*   **Composition**: "has-a" relationship (Dog has a Tail).
*   Embedding lets you call inner methods directly, but it’s not polymorphism.

### **111. What are Go generics and how do you use them?**
Introduced in Go 1.18. Allows writing functions/types that work with *any* type.
```go
func Print[T any](s []T) {
    for _, v := range s { fmt.Println(v) }
}
```

### **112. How to implement a generic function with constraints?**
Use an interface as a constraint.
```go
func Add[T int | float64](a, b T) T {
    return a + b
}
```

### **113. What are type parameters?**
The types defined in square brackets `[T any]` that act as placeholders for valid types.

### **114. Can you implement the Strategy pattern using interfaces?**
Yes, perfectly.
```go
type Sorter interface { Sort([]int) }
type BubbleSort struct{}
type QuickSort struct{}

func SortData(s Sorter, data []int) { s.Sort(data) }
```

### **115. What is middleware in Go web apps?**
A function that wraps an `http.Handler` to perform pre/post-processing (logging, auth, cors).
```go
func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL)
        next.ServeHTTP(w, r)
    })
}
```

### **116. How do you structure code using the Clean Architecture?**
Divide into layers, dependent only on inner layers.
1.  **Entities** (Core business logic).
2.  **Use Cases** (Application logic).
3.  **Controllers/Gateways** (Interface adapters).
4.  **DB/Web** (External drivers).

### **117. What are service and repository layers?**
*   **Repository**: Handles data access (SQL, Files). Returns domain objects.
*   **Service**: Handles business logic. Calls repositories.

### **118. How would you separate concerns in a RESTful Go app?**
*   **Handler**: Parse request, validate input, send response.
*   **Service**: Business logic.
*   **Repository**: DB access.

### **119. What is the importance of interfaces in layered design?**
They decouple layers.
*   The **Service** layer defines a `Repository` interface.
*   The **Repository** implementation satisfies it.
*   Allows swapping DBs or mocking for tests easily.

### **120. How would you implement a plugin system in Go?**
1.  **Interfaces**: Load external code that satisfies an interface.
2.  **hashicorp/go-plugin**: RPC-based plugin system (robust, process isolation).
3.  **Go Plugins (`plugin` package)**: Loads `.so` files (Linux only, tricky restrictions).

### **121. How do you avoid circular dependencies in Go packages?**
*   **Refactor**: Move shared code to a third, lower-level package.
*   **Interfaces**: Define an interface in one package and implement it in another.
*   **Design**: Circular dependencies often signal poor separation of concerns.

## 🟡 **121–140: Generics, Type System, and Advanced Types**

### **122. What is type inference in Go?**
The compiler deduces the type of a variable from the value on the right-hand side.
*   `i := 10` (infers `int`)
*   `GenericFunc(10)` (infers type arguments)

### **123. How do you use generics with struct types?**
Define the type parameter on the struct itself.
```go
type Stack[T any] struct {
    items []T
}
```

### **124. Can you restrict generic types using constraints?**
Yes. You can allow only specific types using an interface.
```go
type Number interface {
    int | float64
}
func Sum[T Number](a, b T) T { ... }
```

### **125. How to create reusable generic containers (e.g., Stack)?**
```go
type Stack[T any] struct { val []T }
func (s *Stack[T]) Push(v T) { s.val = append(s.val, v) }
func (s *Stack[T]) Pop() T   { ... }
```
Eliminates the need for casting from `interface{}`.

### **126. What is the difference between `any` and interface{}?**
Identical. `any` is just an alias for `interface{}` added for readability in generics.

### **127. Can you have multiple constraints in a generic function?**
Yes.
```go
func Map[K comparable, V any](m map[K]V) { ... }
```

### **128. Can interfaces be used in generics?**
Yes, interfaces **are** constraints.
*   **Basic Interfaces**: Have only methods. Can be used as values OR constraints.
*   **General Interfaces**: Have type terms (`int | string`). Can ONLY be used as constraints, not values.

### **129. What is type embedding and how does it differ from inheritance?**
Embedding a struct adds its fields/methods to the outer struct.
*   **Diff**: No dynamic dispatch. The method receiver is still the inner struct, not the outer one.

### **130. How does Go perform type conversion vs. type assertion?**
*   **Conversion**: Changing the type (e.g., `float64(10)`). Works if types are compatible.
*   **Assertion**: Checking/extracting the underlying type of an **interface** (e.g., `i.(int)`). Fails if `i` is not an int.

### **131. What are tagged unions and how can you simulate them in Go?**
A type that can hold one of several known types.
*   **Simulation**: An interface with a private method that only specific structs implement (Sealed Interface).

### **132. What is the use of `iota` in Go?**
A constant generator. Used for creating enums.
```go
const (
    Red = iota // 0
    Blue       // 1
    Green      // 2
)
```

### **133. How are custom types different from type aliases?**
*   **Custom Type** (`type MyInt int`): A **new** type. Cannot directly assign `MyInt` to `int` without conversion. Can attach methods.
*   **Alias** (`type MyInt = int`): Same type. Interchangable.

### **134. What are type sets in Go 1.18+?**
The set of types that can satisfy an interface constraint.
*   `interface{ int | string }` has a type set of just `int` and `string`.

### **135. Can generic types implement interfaces?**
Yes.
```go
type MyList[T any] []T
func (l MyList[T]) String() string { ... } // Implements fmt.Stringer
```

### **136. How do you handle constraints with operations like +, -, *?**
You must constrain T to types that support those operators (e.g., `int | float64`). You cannot assume `T any` supports addition.

### **137. What is structural typing?**
Type compatibility based on structure/methods rather than explicit declarations.
*   Go uses this for **Interfaces** (if it has the methods, it fits).
*   Go does NOT use this for **Structs** (named structs are distinct).

### **138. Explain the difference between concrete and abstract types. **
*   **Concrete**: Structs, primitives. Size and layout are known.
*   **Abstract**: Interfaces. Only behavior is known.

### **139. What are phantom types and are they used in Go?**
Types used only for compile-time checks, with no runtime data.
```go
type ID[T any] string // T is phantom, tags the ID to a specific type
```

### **140. How would you implement an enum pattern in Go?**
Using `const` blocks and custom types.
```go
type Status int
const (
    Pending Status = iota
    Active
)
func (s Status) String() string { ... } // Add method for string rep
```

### **141. How can you implement optional values in Go idiomatically?**
1.  **Pointers**: `*int` (nil = missing).
2.  **Ok idiom**: Return `(int, bool)`.
3.  **Generics**: Create an `Option[T]` struct (less idiomatic in Go).

## 🔵 **141–160: Networking, APIs, and Web Dev**

### **142. How to build a REST API in Go?**
Use `net/http`.
1.  Define handlers: `func getItems(w http.ResponseWriter, r *http.Request)`.
2.  Register routes: `http.HandleFunc("/items", getItems)`.
3.  Start server: `http.ListenAndServe(":8080", nil)`.

### **143. How to parse JSON and XML in Go?**
*   **JSON**: `encoding/json` (`Decoder.Decode` or `Marshal/Unmarshal`).
*   **XML**: `encoding/xml`. Similar API to JSON.

### **144. What is the use of `http.Handler` and `http.HandlerFunc`?**
*   **`Handler`**: An interface with `ServeHTTP(w, r)`.
*   **`HandlerFunc`**: A helper type that adapts a regular function `func(w, r)` to satisfy the `Handler` interface.

### **145. How do you implement middleware manually in Go?**
Wrap an `http.Handler`.
```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // ... pre-processing ...
        next.ServeHTTP(w, r)
        // ... post-processing ...
    })
}
```

### **146. How do you serve static files in Go?**
`http.FileServer`.
```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
```

### **147. How do you handle CORS in Go?**
Set headers manually in middleware or use `rs/cors` package.
```go
w.Header().Set("Access-Control-Allow-Origin", "*")
if r.Method == "OPTIONS" { return }
```

### **148. What are context-based timeouts in HTTP servers?**
Wrapping the request context with a timeout handler.
*   **Built-in**: `http.TimeoutHandler`.
*   **Manual**: `ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)`

### **149. How do you make HTTP requests in Go?**
Use `net/http` Client.
```go
resp, err := http.Get("http://example.com")
defer resp.Body.Close()
```

### **150. How do you manage connection pooling in Go?**
The default `http.Client` does this automatically via `http.Transport`.
*   **Config**: Adjust `MaxIdleConns` and `MaxIdleConnsPerHost` in `http.Transport`.

### **151. What is an HTTP client timeout?**
The limit on the entire request lifecycle (dial, handshake, headers, body).
*   **Set it**: `client := &http.Client{Timeout: 10 * time.Second}`.

### **152. How do you upload and download files via HTTP?**
*   **Upload**: Use `r.FormFile` (multipart).
*   **Download**: `io.Copy(file, resp.Body)`.

### **153. What is graceful shutdown and how do you implement it?**
Stopping the server without dropping active connections.
1.  Listen for OS signals (SIGINT).
2.  Call `server.Shutdown(ctx)` (waits for connections to finish).

### **154. How to work with multipart/form-data in Go?**
request.ParseMultipartForm(). Use `r.MultipartForm` or `r.FormFile("key")`.

### **155. How do you implement rate limiting in Go?**
*   **Token Bucket**: Use `golang.org/x/time/rate`.
*   **Middleware**: Check limiter allowed for each request.

### **156. What is Gorilla Mux and how does it compare with net/http?**
A powerful router.
*   **Vs net/http** (pre-1.22): Mux supports regex paths, variables (`/users/{id}`), and methods (`Methods("GET")`).
*   **Note**: Go 1.22 `ServeMux` now supports methods and variables too!

### **157. What are Go frameworks for web APIs (Gin, Echo)?**
*   **Gin**: Fast, minimal, uses `httprouter` (radix tree).
*   **Echo**: Feature-rich, easier binding/validation.
*   **Fiber**: Expressjs-like, based on `fasthttp` (zero alloc, but not standard `net/http` compatible).

### **158. What are the trade-offs between using `http.ServeMux` and third-party routers?**
*   **ServeMux**: Standard lib, no deps, stable. (Recommended for most now with 1.22).
*   **Third-party**: More syntactical sugar, middleware groups, parameter extraction helpers.

### **159. How would you implement authentication in a Go API?**
1.  **Middleware**: Checks `Authorization` header.
2.  **JWT**: Parse and validate token usage `golang-jwt`.
3.  **Context**: Store user info in `r.Context()` for handlers to use.

### **160. How do you implement file streaming in Go?**
Set `Transfer-Encoding: chunked` (automatic if Content-Length unknown). Writes to `ResponseWriter` are flushed to client immediately if Flusher interface is used.
```go
if f, ok := w.(http.Flusher); ok { f.Flush() }
```

## 🟣 **161–180: Databases and ORMs**

### **161. How do you connect to a PostgreSQL database in Go?**
Use `database/sql` with a driver (e.g., `lib/pq` or `pgx`).
```go
db, err := sql.Open("postgres", "connStr")
```

### **162. What is the difference between `database/sql` and GORM?**
*   **database/sql**: Low-level, raw SQL, performant, verbose scanning.
*   **GORM**: ORM, handles relationships, hooks, migrations, slower than raw SQL.

### **163. How do you handle SQL injections in Go?**
Use **Parameterized Queries** (placeholders `$1`, `?`).
*   Safe: `db.Query("SELECT * FROM users WHERE id=$1", id)`
*   Unsafe: `db.Query(fmt.Sprintf("SELECT ... %s", id))`

### **164. How do you manage connection pools in `database/sql`?**
It's built-in.
*   `db.SetMaxOpenConns(n)`
*   `db.SetMaxIdleConns(n)`
*   `db.SetConnMaxLifetime(d)`

### **165. What are prepared statements in Go?**
Pre-compiling SQL execution plan.
```go
stmt, _ := db.Prepare("INSERT INTO users VALUES($1)")
stmt.Exec("Alice")
```

### **166. How do you map SQL rows to structs?**
*   **Manual**: `rows.Scan(&dest.Name, &dest.Age)` for each field.
*   **Libs**: `sqlx` (`structscan`), or ORMs.

### **167. What are transactions and how are they implemented in Go?**
Atomic operations.
```go
tx, _ := db.Begin()
// ... operations using tx.Exec ...
if err != nil { tx.Rollback() }
tx.Commit()
```

### **168. How do you handle database migrations in Go?**
Use tools like:
*   **golang-migrate**: File-based SQL (up/down).
*   **Goose**: Go/SQL migrations.
*   **GORM**: AutoMigrate (code-first).

### **169. What is the use of `sqlx` in Go?**
An extension to `database/sql`.
*   Named parameters (`:name` instead of `$1`).
*   Struct Mapping (`db.Select(&users, "query")`).

### **170. What are the pros and cons of using an ORM in Go?**
*   **Pros**: Speed of dev, type safety, migrations.
*   **Cons**: Hidden magic, performance overhead, complex queries become hard.

### **171. How would you implement pagination in SQL queries?**
`LIMIT` and `OFFSET`.
*   Note: Large OFFSET is slow. Use **Keyset Pagination** (WHERE id > last_seen_id LIMIT n) for huge datasets.

### **172. How do you log SQL queries in Go?**
*   **GORM**: `db.Debug().Find(...)`
*   **Raw**: Wrap the driver (instrumentation) or just print before exec.

### **173. What is the N+1 problem in ORMs and how to avoid it?**
Fetching parents (1 query) and then looping to fetch children (N queries).
*   **Avoid**: Use "Eager Loading" (GORM `Preload`, JOINs in SQL).

### **174. How do you implement caching for DB queries in Go?**
Check Cache (Redis) -> If Miss -> Query DB -> Set Cache.
*   Use patterns like **Look-aside**.

### **175. How do you write custom SQL queries using GORM?**
`db.Raw("SELECT id, name FROM users WHERE name = ?", "jinzhu").Scan(&result)`

### **176. How do you handle one-to-many and many-to-many relationships in GORM?**
Struct tags and slices.
*   `HasMany`: `Posts []Post` in User struct.
*   `Many2Many`: `Languages []Language `gorm:"many2many:user_languages;"``

### **177. How would you structure your database layer in a Go project?**
Define a `Repository` interface.
```go
type UserRepository interface {
    Get(id int) (*User, error)
}
```
Implement with `PostgresUserRepo` struct.

### **178. What is context propagation in database calls?**
Pass `context.Context` to DB methods (`QueryContext`, `ExecContext`) to allow cancellation/timeouts from the upper layers to stop the DB query.

### **179. How do you handle long-running queries or timeouts?**
Use `context.WithTimeout`.
```go
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
db.QueryContext(ctx, ...)
```

### **180. How do you write unit tests for code that interacts with the DB?**
1.  **Mocking**: Use `go-sqlmock` to simulate DB responses without a real DB.
2.  **Docker**: Spin up a real Test DB container (more reliable integration test).

## 🔴 **181–200: Tools, Testing, CI/CD, Ecosystem**

### **181. What is `go vet` and what does it catch?**
A static analysis tool.
*   Catches common mistakes (e.g., `Printf` format mismatches, unreachable code, mutex copying).

### **182. How does `go fmt` help maintain code quality?**
It strictly formats code to a standard style (tabs for indent, brace placement).
*   **Result**: Eliminates bike-shedding arguments about code style.

### **183. What is `golangci-lint`?**
A fast linter runner that aggregates many linters (errcheck, staticcheck, gosec, etc.). Highly recommended for CI.

### **184. What is the difference between `go run`, `go build`, and `go install`?**
*   **`run`**: Compiles and executes in temp directory.
*   **`build`**: Compiles to an executable in current directory.
*   **`install`**: Compiles and moves executable to `$GOPATH/bin`.

### **185. How does `go generate` work?**
It scans for `//go:generate command` comments and runs them.
*   **Usage**: Generating mocks (`mockgen`), protobufs (`protoc`), or stringers.

### **186. What is a build constraint?**
`//go:build linux` or `// +build linux`.
Tells the compiler to only include this file if the condition is met.

### **187. How do you write tests in Go?**
`func TestXxx(t *testing.T) { ... }` in a `_test.go` file.

### **188. How do you test for expected panics?**
```go
defer func() {
    if r := recover(); r == nil { t.Errorf("Expected panic") }
}()
triggerPanic()
```

### **189. What are mocks and how do you use them in Go?**
Simulated objects that mimic behavior of real objects.
*   **Tools**: `gomock`, `testify/mock`.

### **190. How do you use the `testing` and `testify` packages?**
*   **testing**: Std lib. Minimal assertions (`if got != want { t.Error }`).
*   **testify**: Popular 3rd party. Offers `assert.Equal(t, got, want)` and suites.

### **191. How do you structure test files in Go?**
Place `foo_test.go` next to `foo.go`.
*   **Internal Test**: `package foo` (access to private members).
*   **External Test**: `package foo_test` (tests public API only, avoids circular deps).

### **192. What is a benchmark test?**
Measures performance.
```go
func BenchmarkFoo(b *testing.B) {
    for i:=0; i<b.N; i++ { Foo() }
}
```

### **193. How do you measure test coverage in Go?**
`go test -coverprofile=coverage.out`
`go tool cover -html=coverage.out` (Visualizes it).

### **194. How do you test concurrent functions?**
Use `WaitGroups` or channels in tests to ensure completion.
*   **Race Detector**: ALWAYS run with `-race`.

### **195. What is a race detector and how do you use it?**
A runtime tool that detects data races (unsynchronized access to shared memory).
*   `go test -race ./...`

### **196. What is `go.mod` and `go.sum`?**
*   **go.mod**: Dependency definitions.
*   **go.sum**: Checksum of specific versions of all direct and indirect dependencies. Ensures integrity.

### **197. How does semantic versioning work in Go modules?**
Go modules enforce semver.
*   **v0.x.x**: Unstable.
*   **v1.x.x**: Stable.
*   **v2.x.x**: Breaking changes (must change import path to `/v2`).

### **198. How to build and deploy a Go binary to production?**
1.  **Build**: `CGO_ENABLED=0 GOOS=linux go build -o app .` (Static binary).
2.  **Docker**: Copy specific binary to purely minimal `scratch` or `alpine` image.
3.  **Run**: Just the binary. No runtime/VM needed.

### **199. What tools are used for Dockerizing Go apps?**
*   **Docker**: Multi-stage builds (build in `golang` image, run in `alpine`).
*   **Ko**: Tool to build container images from Go source without Dockerfile.

### **200. How do you set up a CI/CD pipeline for a Go project?**
*   **Steps**:
    1.  Checkout code.
    2.  `go mod download`.
    3.  `golangci-lint run`.
    4.  `go test -race -cover ./...`.
    5.  `go build`.
    6.  Push binary/image to registry.

## 🟢 **201–220: Performance & Optimization**

### **201. How do you optimize memory usage in Go?**
1.  **Reduce Allocations**: Re-use objects (sync.Pool), pre-allocate slices (`make([]int, 0, 100)`).
2.  **Struct Layout**: Optimize field ordering to minimize padding.
3.  **Avoid Pointers**: Use values where possible to reduce GC pressure (if objects are small).

### **202. What is memory escape analysis in Go?**
The compiler's process to determine if a variable should live on the **Stack** or **Heap**.
*   **Stack**: Does not escape (fast, no GC).
*   **Heap**: Escapes function scope (slower, GC required).
*   **Check**: `go build -gcflags="-m"`

### **203. How to reduce allocations in tight loops?**
*   Move variable declarations **outside** the loop.
*   Use `strings.Builder` instead of `+` concatenation.
*   Reuse buffers.

### **204. How do you profile a Go application?**
Use **pprof**.
1.  Import `_ "net/http/pprof"`.
2.  Start HTTP server.
3.  Run `go tool pprof http://localhost:PORT/debug/pprof/heap`.

### **205. What is the use of `pprof` in Go?**
A tool to visualize and analyze profile data (CPU, Heap, Goroutines, Mutex contention).

### **206. How do you benchmark against memory allocations?**
Use `b.ReportAllocs()` in your benchmark test.
```go
func BenchmarkFoo(b *testing.B) {
    b.ReportAllocs()
    // ...
}
```

### **207. How can you avoid unnecessary heap allocations?**
*   Pass small structs by value.
*   Avoid using `interface{}` parameter types if concrete types work.
*   Use array buffers instead of slices if size is fixed and small.

### **208. What is inlining and how does the Go compiler handle it?**
Replacing a function call with the function body itself to save call overhead.
*   Go inlines simple, small functions (leaf functions).
*   **Check**: `go build -gcflags="-m"`.

### **209. How do you debug GC pauses?**
*   Run with `GODEBUG=gctrace=1`.
*   Look for long pause times in the output.
*   Use execution tracer (`go tool trace`).

### **210. What are some common performance bottlenecks in Go apps?**
*   **GC Pressure**: Too many small heap allocations.
*   **Contention**: Too many goroutines waiting on a single Mutex.
*   **I/O**: Blocking network/disk calls.

### **211. How to detect and fix memory leaks?**
Leaks in Go are usually logic errors (e.g., goroutines that never exit, growing maps/slices in globals).
*   **Detect**: `pprof` heap profile diffs (compare base vs current).

### **212. How do you find goroutine leaks?**
*   **Runtime**: `runtime.NumGoroutine()` monitoring.
*   **Pprof**: `go tool pprof .../goroutine`.
*   **Tests**: Use `goleak` library (uber-go/goleak).

### **213. How do you tune GC parameters in production?**
*   **GOGC**: Percent growth of heap before next GC (default 100). Increase to run GC less often (trades RAM for CPU).
*   **GOMEMLIMIT** (Go 1.19+): Soft memory limit.

### **214. How to avoid blocking operations in hot paths?**
*   Use non-blocking channel selects.
*   Do I/O in a separate goroutine.
*   Use buffered channels.

### **215. What are the trade-offs of pooling in Go?**
*   **Pros**: Reduced allocations/GC.
*   **Cons**: Complexity, risk of leaking state (forgetting to reset object), slight CPU overhead to manage pool.

### **216. How do you measure latency and throughput in Go APIs?**
*   **Middleware**: Capture `start := time.Now()` and log `time.Since(start)`.
*   **Metrics**: Populate Prometheus histograms.

### **217. What is backpressure and how do you handle it?**
Signaling to producers to slow down.
*   **Go**: Buffered channels with fixed size. If channel is full, sender blocks (natural backpressure).

### **218. When should you prefer sync.Pool?**
For objects that are **expensive to allocate** and **short-lived**, and frequently re-used (e.g., JSON encoders, buffers).

### **219. How do you manage high concurrency with low resource usage?**
*   **Worker Pools**: Limit active goroutines.
*   **Non-blocking I/O**: (Netpoller handles this automatically in Go).

### **220. How do you monitor a Go application in production?**
*   **Metrics**: Prometheus (via `/metrics` endpoint).
*   **Logs**: Structured logs (ELK/Loki).
*   **Tracing**: OpenTelemetry/Jaeger.

## 🟡 **221–240: Files, OS, and System Programming**

### **221. How do you read a file line by line in Go?**
Use `bufio.Scanner`.
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    process(scanner.Text())
}
```

### **222. How do you write large files efficiently?**
Use `bufio.Writer` to buffer writes and flush periodically, instead of writing small chunks directly to disk.

### **223. How do you watch file system changes in Go?**
Use **fsnotify** (library) or specific OS syscalls.
*   The standard library doesn't have a cross-platform file watcher.

### **224. How to get file metadata like size, mod time?**
`os.Stat(filename)` returns a `FileInfo` struct with Size(), ModTime(), Mode().

### **225. How do you work with CSV files in Go?**
`encoding/csv`.
```go
reader := csv.NewReader(file)
record, err := reader.Read()
```

### **226. How do you compress and decompress files in Go?**
Use packages: `compress/gzip`, `compress/zlib`.
*   Wrap your reader/writer: `gzip.NewWriter(file)`.

### **227. How do you execute shell commands from Go?**
`os/exec`.
```go
cmd := exec.Command("ls", "-la")
output, err := cmd.CombinedOutput()
```

### **228. What is the `os/exec` package used for?**
Running external processes, handling stdin/stdout pipes, and environment variables for subprocesses.

### **229. How do you set environment variables in Go?**
*   **For Process**: `os.Setenv("KEY", "VALUE")`.
*   **For Subprocess**: `cmd.Env = append(os.Environ(), "KEY=VALUE")`.

### **230. How to create and manage temp files/directories?**
`os.MkdirTemp` and `os.CreateTemp`.
*   Remember to remove them: `defer os.Remove(tempFile.Name())`.

### **231. How do you handle signals like SIGINT in Go?**
`os/signal`.
```go
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
<-c // Block until signal received
```

### **232. How do you gracefully shut down a CLI app?**
Listen for SIGINT/SIGTERM (see above), then run cleanup logic (flush logs, close DB) before `os.Exit(0)`.

### **233. What are file descriptors and how does Go manage them?**
Integer handles to open files/sockets. Go uses non-blocking FDs with the Netpoller to manage thousands of concurrent I/O ops efficiently.

### **234. How to handle large file uploads and streaming?**
*   **Upload**: Use `io.Copy(dst, src)` to stream data without loading into RAM.
*   **Web**: Use `multipart.Reader` to stream parts.

### **235. How do you access OS-specific syscalls in Go?**
`syscall` package (deprecated) or `golang.org/x/sys/unix` (recommended).

### **236. How do you implement a simple CLI tool in Go?**
Read `os.Args` (slice of arguments).
*   For advanced CLIs, use **Cobra** or **Urfave/cli**.

### **237. How do you build cross-platform binaries in Go?**
Set environment variables:
`GOOS=windows GOARCH=amd64 go build`
`GOOS=linux GOARCH=arm64 go build`

### **238. What is syscall vs os vs exec package difference?**
*   **syscall**: Low-level kernel interface.
*   **os**: High-level file/process manipulation (portable).
*   **exec**: Running external commands.

### **239. How do you write to logs with rotation?**
Standard lib doesn't support rotation.
*   Use `lumberjack.v2` (nats-io/lumberjack) which implements `io.Writer` with rotation logic.

### **240. What is the use of `ioutil` and its deprecation?**
It provided helpers (`ReadFile`, `ReadAll`).
*   **Status**: **Deprecated** in Go 1.16.
*   **Replacements**:
    *   `ioutil.ReadAll` -> `io.ReadAll`
    *   `ioutil.ReadFile` -> `os.ReadFile`
    *   `ioutil.Discard` -> `io.Discard`

## 🔵 **241–260: Microservices, gRPC, and Communication**

### **241. What is gRPC and how is it used with Go?**
Google's high-performance RPC framework. Uses Protocol Buffers.
*   **Go**: Generate code from `.proto` files using `protoc-gen-go`.

### **242. How do you define Protobuf messages for Go?**
```proto
message User {
  string name = 1;
  int32 id = 2;
}
```

### **243. What are the benefits of gRPC over REST?**
*   **Binary**: Smaller, faster serialization (Protobuf vs JSON).
*   **HTTP/2**: Multiplexing, streaming.
*   **Schema**: Strongly typed contracts (.proto).

### **244. How do you implement unary and streaming RPC in Go?**
*   **Unary**: Single request, single response. `GetFeature(ctx, Point) returns (Feature)`.
*   **Streaming**: `stream` keyword in proto.
    *   Server Side: `func (s) ListFeatures(req, stream)` loops to send.

### **245. What is the difference between gRPC and HTTP/2?**
gRPC **uses** HTTP/2 as its transport layer. It adds semantics for RPC calls on top of HTTP/2 frames.

### **246. How do you add authentication in gRPC services?**
Use **Interceptors** (middleware) or Per-RPC credentials.
*   Extract token from `metadata` (headers).

### **247. How do you handle timeouts and retries in gRPC?**
*   **Timeouts**: `context.WithTimeout` on client side. Propagated via header.
*   **Retries**: Configurable in Service Config (client-side) or custom interceptor logic.

### **248. How do you secure gRPC communication?**
TLS/SSL.
*   Pass `credentials.NewServerTLSFromCert` to `grpc.NewServer`.

### **249. How do microservices communicate securely in Go?**
*   **mTLS**: Mutual TLS (both sides verify certs).
*   **Tokens**: JWT passed in Authorization header.

### **250. What are message queues and how to use them in Go?**
Async communication tools (Kafka, RabbitMQ, NATS).
*   Decouples services.
*   Use client libraries (e.g., `github.com/segmentio/kafka-go`).

### **251. How to use NATS or Kafka in Go?**
*   **NATS**: Native Go system. Lightweight. `nc.Publish("subject", data)`.
*   **Kafka**: High throughput log.

### **252. What are sagas and how would you implement them in Go?**
A pattern for distributed transactions across services.
*   **Choreography**: Events trigger next steps (Service A -> Event -> Service B).
*   **Orchestration**: A central Coordinator service calls A, then B. Rolls back if error.

### **253. How would you trace requests across services?**
**Distributed Tracing** (OpenTelemetry).
*   Pass a `TraceID` in headers (HTTP/gRPC/Messaging) across boundaries.
*   Extract and inject context.

### **254. What is service discovery and how do you handle it?**
Finding the IP/Port of a service.
*   **K8s**: DNS (service names).
*   **Consul/Etcd**: Registry lookup.

### **255. How do you implement rate limiting across services?**
*   **Local**: Token bucket in memory.
*   **Distributed**: Redis-based counter (Lua script) to share limit state.

### **256. What is the role of API gateway in microservices?**
Single entry point. Handles auth, routing, rate limiting, and aggregation.

### **257. How do you use OpenTelemetry with Go?**
Initialize a TracerProvider, create Spans for operations.
```go
ctx, span := tracer.Start(ctx, "operation")
defer span.End()
```

### **258. How do you log correlation IDs between services?**
Extract standard headers (e.g., `X-Request-ID` or TraceID) in middleware and add them to the logger context.

### **259. How would you handle distributed transactions in Go?**
Avoid Two-Phase Commit (2PC) if possible (locking issues). Use **Saga Pattern** (compensating transactions) or Event Sourcing.

### **260. How to deal with partial failures in distributed systems?**
*   **Timeouts**: Fail fast.
*   **Circuit Breakers**: Stop calling failing service.
*   **Retries**: With exponential backoff (idempotency key required).

## 🟣 **261–280: Security and Best Practices**

### **261. How do you prevent injection attacks in Go?**
*   **SQL**: Use parameterized queries (`$1`).
*   **Command**: Avoid `exec.Command` with user input. Sanitize args.

### **262. What are Go's common security vulnerabilities?**
*   Integer overflow (less common but possible).
*   Race conditions.
*   Improper error handling (leaking info).
*   Dependency vulnerabilities (use `govulncheck`).

### **263. How do you hash passwords securely in Go?**
Use **bcrypt** or **argon2**.
*   **Never** use MD5 or SHA1/SHA256 directly for passwords (too fast).

### **264. How to use `bcrypt` in Go?**
`golang.org/x/crypto/bcrypt`.
```go
hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
err := bcrypt.CompareHashAndPassword(hash, []byte("password"))
```

### **265. How do you validate input in Go APIs?**
*   **Manual**: Check regex, length.
*   **Library**: `go-playground/validator`. Struct tags (`validate:"required,email"`).

### **266. How do you implement JWT authentication?**
`golang-jwt/jwt` library.
*   **Sign**: Create token with claims, sign with secret.
*   **Verify**: Parse token, check signature method (HMAC), extract claims.

### **267. How do you prevent race conditions in Go?**
*   **Design**: Don't share memory.
*   **Sync**: Use Mutexes/Channels.
*   **Detect**: `go test -race`.

### **268. What is CSRF and how to mitigate it?**
Cross-Site Request Forgery.
*   **Fix**: Use Anti-CSRF tokens (gorilla/csrf) for form POSTs. SameSite cookie attribute (`SameSite: Strict`).

### **269. How to use HTTPS in Go servers?**
`http.ListenAndServeTLS(":443", "cert.pem", "key.pem")`.
*   **Auto**: Use `golang.org/x/crypto/acme/autocert` for Let's Encrypt.

### **270. How do you sign and verify data in Go?**
`crypto/hmac` or RSA/ECDSA signatures.
```go
mac := hmac.New(sha256.New, key)
mac.Write(data)
signature := mac.Sum(nil)
```

### **271. What are best practices for handling secrets in Go?**
*   **Env Vars**: Read from os.Getenv.
*   **Vault**: Use Hashicorp Vault.
*   **Code**: NEVER hardcode secrets in source. Use `secret` structs with redaction in logs.

### **272. How do you handle OAuth2 flows in Go?**
Use `golang.org/x/oauth2` library. Supports standard flows (Auth Code, Client Creds).

### **273. How do you restrict file uploads (size/type)?**
*   **Size**: `r.Body = http.MaxBytesReader(w, r.Body, 10<<20)` (10MB).
*   **Type**: Read first 512 bytes and use `http.DetectContentType`. Do NOT trust extension.

### **274. How do you set up CORS properly in Go?**
Don't use `*` in production. Whitelist specific origins. Support preflight `OPTIONS`.

### **275. How do you scan Go code for vulnerabilities?**
*   `govulncheck ./...` (Official Go tool).
*   `gosec ./...` (Security linter).

### **276. What is the Go ecosystem for SAST tools?**
Static Application Security Testing.
*   **gosec**: Best for code patterns.
*   **Semgrep**: Rule-based code scanning.

### **277. How to handle brute force protection in APIs?**
Implement Rate Limiting (Token Bucket) keyed by IP or Username.
*   Exponential delays after N failed attempts.

### **278. How to secure communication between microservices?**
mTLS (mutual TLS). Ensures only authorized services can talk to each other.

### **279. What is the use of `context.Context` in secure APIs?**
Carries **scoped credentials** (e.g., User Claims) and **cancellation signals** to prevent resource exhaustion attacks (timeouts).

### **280. What is certificate pinning and can it be used in Go?**
Verifying the server's cert matches a hardcoded hash.
*   Yes, in `tls.Config{VerifyPeerCertificate: func...}`. Rarely needed for server-to-server, mostly for mobile clients.

## 🔴 **281–300: Testing Strategy, CI/CD, Observability**

### **281. What are test doubles and how are they used in Go?**
Replacements for real dependencies.
*   **Mock**: Verifies behavior (method called X times).
*   **Stub**: Returns fixed data.
*   **Fake**: Simplified implementation (in-memory DB).
*   **Spy**: Records calls.

### **282. How do you structure unit vs integration tests?**
*   **Unit**: Test single function/struct in isolation. Mock deps. Fast.
*   **Integration**: Test interaction (e.g., API -> DB). Real deps (Docker). Slow.

### **283. What are flaky tests and how do you identify them?**
Tests that pass/fail randomly.
*   **Causes**: Race conditions, relying on time, relying on order.
*   **Identify**: Run same test 100 times (`go test -count=100`).

### **284. How do you write deterministic tests for concurrency?**
Avoid `time.Sleep`. Use Channels/WaitGroups or `sync.Cond` to signal exact state changes.

### **285. How do you test RESTful APIs in Go?**
Use `httptest` package.
```go
req := httptest.NewRequest("GET", "/users", nil)
w := httptest.NewRecorder()
handler(w, req)
if w.Code != 200 { t.Error(...) }
```

### **286. How do you mock HTTP calls?**
Use `httptest.NewServer`.
```go
server := httptest.NewServer(http.HandlerFunc(func(w, r) {
    w.Write([]byte("mock response"))
}))
defer server.Close()
client := server.Client() // Use this client in your code
```

### **287. What is Golden Testing in Go?**
Comparing output against a "golden" file (expected output).
*   Useful for complex JSON/HTML outputs. `go test -update` to regenerate golden files.

### **288. How do you run tests in parallel?**
Call `t.Parallel()` at start of test.
*   **Caveat**: Be careful with shared state (env vars, global maps).

### **289. How do you mock time-dependent code?**
Don't use `time.Now()` directly. Use an interface `Clock` with `Now()` method. Provide a `MockClock` in tests.

### **290. How do you simulate DB failures in tests?**
Use **go-sqlmock** to match a query and return `sql.ErrConnDone` or custom error.

### **291. How do you use GitHub Actions to test Go apps?**
YAML workflow:
```yaml
steps:
  - uses: actions/setup-go@v4
  - run: go test ./...
```

### **292. What is the structure of a Makefile for Go?**
Common targets:
*   `test`: `go test ./...`
*   `build`: `go build ...`
*   `lint`: `golangci-lint run`
*   `clean`: `rm -rf bin/`

### **293. How to build and test Go code in Docker?**
Multi-stage Dockerfile.
1.  **Build Stage**: Run tests and build binary.
2.  **Final Stage**: Copy binary. Fails build if tests fail.

### **294. What CI tools are commonly used for Go projects?**
GitHub Actions, GitLab CI, CircleCI, Jenkins.

### **295. What are the benefits of go:embed for test fixtures?**
Include test data (JSON/SQL) inside the test binary without worrying about relative paths/working directory.
```go
//go:embed testdata/*.json
var fixtures embed.FS
```

### **296. How do you generate coverage reports in HTML?**
1.  `go test -coverprofile=c.out`
2.  `go tool cover -html=c.out -o coverage.html`

### **297. How to collect logs and metrics from Go services?**
*   **Logs**: Stdout/Stderr (collected by Fluentd/Promtail).
*   **Metrics**: Prometheus client (`/metrics` HTTP handler).

### **298. What is structured logging in Go?**
Logging as JSON (key-value) instead of plain text.
*   Easier to query in ELK/Splunk.
*   Use `slog` (std lib) or `zap`.

### **299. What are common logging libraries in Go?**
*   **Zap**: Uber's fast logger.
*   **Logrus**: Popular but in maintenance mode.
*   **Slog**: New standard library (Go 1.21+).

### **300. How do you aggregate and search logs across services?**
Centralized Logging stack:
*   **ELK**: Elasticsearch, Logstash, Kibana.
*   **LGTM**: Loki, Grafana, Tempo (Tracing), Mimir.

## 🟢 **301–320: Go Internals and Runtime**

### **301. How does the Go scheduler work?**
It uses an **M:N scheduler**.
*   **M (Machine)**: OS Thread.
*   **P (Processor)**: Logical processor (held by M to execute G). Default P = NumCPU.
*   **G (Goroutine)**: Code to run.
*   **Work Stealing**: If a P runs out of Gs, it steals half from another P.

### **302. What is M:N scheduling in Golang?**
Multiplexing **M** goroutines onto **N** OS threads.
*   Allows thousands of goroutines to run on a few OS threads context-switching in user space (cheap).

### **303. How does the Go garbage collector work?**
**Concurrent Mark and Sweep**.
*   **Mark Phase**: Starts from roots, marks reachable objects (Tri-color: White, Grey, Black).
*   **Sweep Phase**: Reclaims white objects.
*   **Write Barrier**: Ensures consistency while app runs concurrently.

### **304. What are STW (stop-the-world) events in GC?**
Pauses where all application goroutines are stopped.
*   Go GC has two very short STW phases: initial stack scan and mark termination. Usually sub-millisecond.

### **305. How are goroutines implemented under the hood?**
They are structs (`runtime.g`) managing their own stack and execution state.
*   Initially **2KB** stack. Grows/shrinks dynamically.

### **306. How does stack growth work in Go?**
When a function call needs more stack than available:
1.  Runtime allocates a larger stack (2x).
2.  Copies data from old stack to new.
3.  Adjusts pointers to point to new stack.

### **307. What is the difference between blocking and non-blocking channels internally?**
*   **Blocking**: Goroutine parks itself (`gopark`) in the channel's `recvq` or `sendq` queue.
*   **Non-blocking**: If operation can't proceed, it returns immediately (used in `select` with default).

### **308. What is a GOMAXPROCS and how does it affect execution?**
Limits the number of **P** (Processors) executing User-level Go code simultaneously.
*   Does not limit **M** (OS threads) blocked in syscalls.

### **309. How does Go manage memory fragmentation?**
It uses a **TCMalloc-like** allocator.
*   Divides memory into spans of size classes (e.g., 8B, 16B, 32B).
*   Small objects fit perfectly into these classes, minimizing external fragmentation.

### **310. How are maps implemented internally in Go?**
**Hash Map** with buckets.
*   Each bucket holds up to 8 key-value pairs.
*   Overflow buckets are chained if collisions occur.

### **311. How does slice backing array reallocation work?**
`append` checks capacity. If full:
*   Allocates new array (2x size if < 256, else ~1.25x).
*   Copies elements (using `memmove`).
*   Returns new slice header.

### **312. What is the zero value concept in Go?**
Memory is initialized to binary zero.
*   Prevents uninitialized memory access bugs (common in C).

### **313. How does Go avoid data races with its memory model?**
It doesn't strictly "avoid" them; you must use synchronization.
*   **Happens-Before**: Defines ordering guarantees (e.g., Unlock happens before Lock).

### **314. What is escape analysis and how can you visualize it?**
Compiler optimization to decide stack vs heap.
*   **Visualize**: `go build -gcflags="-m -m"` (double m for verbose).

### **315. How are method sets determined in Go?**
Static rules based on receiver type.
*   `T` has methods with `(t T)` receiver.
*   `*T` has methods with `(t T)` AND `(t *T)` receivers.

### **316. What is the difference between pointer receiver and value receiver at runtime?**
*   **Value**: Field access copies the struct.
*   **Pointer**: Field access uses the address.
*   **Interface**: If assigned to interface, `T` can only call value methods; `*T` can call both.

### **317. How does Go handle panics internally?**
It walks up the stack, running deferred functions.
*   If `recover()` is found, it stops unwinding and resumes.
*   If stack empties, process crashes.

### **318. How is reflection implemented in Go?**
Interfaces store `(type, value)` pair.
*   `reflect.TypeOf` unpacks the type.
*   `reflect.ValueOf` unpacks the value.

### **319. What is type identity in Go?**
Two types are identical if they have the same literal structure (for unnamed types) or same declaration (for named types).
*   `type A int` and `int` are **different**.
*   `[]int` and `[]int` are **identical**.

### **320. How are interface values represented in memory?**
A two-word struct: `(tab, data)`.
*   `tab`: Pointer to ITab (Type info + list of function pointers for methods).
*   `data`: Pointer to the concrete value.

## 🟡 **321–340: DevOps, Docker, and Deployment**

### **321. How do you containerize a Go application?**
Write a Dockerfile.
```dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
COPY --from=builder /app/main .
CMD ["./main"]
```

### **322. What is a multi-stage Docker build and how does it help with Go?**
Using two `FROM` instructions.
*   **Stage 1**: Has Go compiler (heavy). Builds the binary.
*   **Stage 2**: Has only runtime (Alpine/Scratch). Small.
*   **Significantly** reduces image size (e.g., 800MB -> 20MB).

### **323. How do you reduce the size of a Go Docker image?**
*   Use `scratch` or `distroless` base.
*   Strip debug symbols: `go build -ldflags="-s -w"`.

### **324. How do you handle secrets in Go apps deployed via Docker?**
Inject via Environment Variables (`docker run -e SECRET=...`) or mount volumes (Kubernetes Secrets).
*   Never bake secrets into the image.

### **325. How do you use environment variables in Go?**
`os.Getenv("PORT")`.
*   Use `godotenv` for local `.env` files.

### **326. How do you compile a static Go binary for Alpine Linux?**
Alpine uses **musl** libc.
`CGO_ENABLED=0 GOOS=linux go build` (Disables CGO, makes it purely static).

### **327. What is `scratch` image in Docker and why is it used with Go?**
An empty image (0 bytes).
*   Since Go binaries are static/self-contained, they run perfectly here.
*   **Caveat**: Need to copy CA certificates manually for HTTPS.

### **328. How do you manage config files in Go across environments?**
*   **Viper**: Reads JSON/YAML/Env.
*   **Env Vars**: The 12-factor way.

### **329. How do you build Go binaries for different OS/arch?**
Cross-compilation.
`GOOS=windows GOARCH=amd64`
`GOOS=darwin GOARCH=arm64` (M1 Mac)

### **330. How do you use GoReleaser?**
Tool to automate releasing Go binaries.
*   Generates binaries for all platforms.
*   Creates GitHub Releases / Docker images.
*   Config: `.goreleaser.yaml`.

### **331. What is a Docker healthcheck for a Go app?**
A command in Dockerfile to verify app is running.
`HEALTHCHECK CMD curl -f http://localhost:8080/health || exit 1`

### **332. How do you log container stdout/stderr from Go?**
Just write to `os.Stdout` (JSON preferred). Docker daemon captures it.

### **333. How do you set up autoscaling for Go services?**
Kubernetes HPA (Horizontal Pod Autoscaler).
*   Scale based on CPU/Memory usage (Go runtime exports these).

### **334. How would you containerize a gRPC Go service?**
Same as HTTP, but expose the gRPC port.
*   Healthcheck: Use `grpc_health_probe`.

### **335. How to deploy Go microservices in Kubernetes?**
Define `Deployment` and `Service` YAMLs.
*   Use `readinessProbe` and `livenessProbe`.

### **336. How do you write Helm charts for a Go app?**
Templatize your K8s YAMLs.
*   `values.yaml` controls image tag, replicas, config.

### **337. How do you monitor a Go service in production?**
Prometheus + Grafana.
*   Expose metrics at `/metrics`.

### **338. How do you use Prometheus with a Go app?**
Library: `github.com/prometheus/client_golang`
*   Register counters/gauges.
*   Start HTTP handler.

### **339. How do you enable structured logging in production?**
Switch logger to JSON format.
*   `log.SetFormatter(&log.JSONFormatter{})` (Logrus) or `slog.NewJSONHandler`.

### **340. How do you handle log rotation in containerized Go apps?**
**DO NOT** do it in the app.
*   Log to stdout.
*   Let Docker/Kubernetes/Filebeat handle rotation of the log file on the host.

## 🔵 **341–360: Streaming, Messaging, and Asynchronous Processing**

### **341. How do you consume messages from Kafka in Go?**
Use **sarama** or **kafka-go**.
*   **kafka-go**: `reader.ReadMessage(ctx)`. Simple API.
*   **Sarama**: Low-level (ConsumerGroup).

### **342. How do you publish messages to a RabbitMQ topic?**
Use `streadway/amqp`.
1.  Connect (`amqp.Dial`).
2.  Open Channel (`conn.Channel`).
3.  Publish (`ch.Publish("exchange", "key", ...)`).

### **343. What is the idiomatic way to implement a message handler in Go?**
Define a `Handler` interface (like `http.Handler`).
```go
type Handler interface {
    Handle(ctx context.Context, msg []byte) error
}
```

### **344. How would you implement a worker pool pattern?**
1.  **Job Channel**: `jobs := make(chan Job, 100)`.
2.  **Workers**: Loop on `range jobs`.
3.  **Dispatch**: `go worker(jobs)` * N.

### **345. How do you use the `context` package for cancellation in streaming apps?**
Pass context to the stream reader.
*   If `ctx.Done()`, stop reading loop and close resources.

### **346. How do you retry failed messages in Go?**
Use **Exponential Backoff**.
*   Loop with sleep: `time.Sleep(base * 2^retries)`.
*   Or store in "Dead Letter Queue" after N tries.

### **347. What is dead-letter queue and how do you use it?**
A separate queue for messages that failed processing (poison pills).
*   Prevents blocking the main pipeline.

### **348. How do you handle idempotency in message consumers?**
Ensure processing same message twice has no side effect.
*   **DB**: Use a unique constraint on Message ID.
*   **Redis**: `SETNX msg_id "processed"`.

### **349. How do you implement exponential backoff in Go?**
```go
func retry(op func() error) {
    wait := 1 * time.Second
    for i := 0; i < 5; i++ {
        if err := op(); err == nil { return }
        time.Sleep(wait)
        wait *= 2
    }
}
```

### **350. How do you stream logs to a file/socket in real-time?**
Write to `io.MultiWriter(os.Stdout, file/socketWriter)`.
*   Wrap with a buffer (`bufio.Writer`) for performance.

### **351. How do you work with WebSockets in Go?**
Use `gorilla/websocket` or `nhooyr.io/websocket`.
*   **Upgrade**: Convert HTTP to WS.
*   **Loop**: `conn.ReadMessage` / `conn.WriteMessage`.

### **352. How do you handle bi-directional streaming in gRPC?**
`func (s) Chat(stream Chat_Stream) error`.
*   Spawn a goroutine to `stream.Recv()`.
*   Use main thread to `stream.Send()`.

### **353. What is Server-Sent Events and how is it done in Go?**
One-way stream (Server -> Client).
*   Header: `Content-Type: text/event-stream`.
*   Body: `data: message\n\n`.
*   Flush after every write.

### **354. How do you manage fan-in/fan-out channel patterns?**
*   **Fan-Out**: Multiple workers reading same channel (competing consumer) OR loop sending to multiple channels.
*   **Fan-In**: `select` reading from multiple channels OR `reflect.Select` for dynamic list.

### **355. How would you implement throttling on async tasks?**
**Token Bucket** or **Buffered Channel**.
```go
limiter := make(chan struct{}, 10) // Limit 10
for req := range reqs {
    limiter <- struct{}{}
    go func() {
        process(req)
        <-limiter
    }()
}
```

### **356. How do you avoid data races when consuming messages?**
*   Process message **locally** in the worker goroutine.
*   If shared state update needed, use Mutex.

### **357. How would you implement a message queue from scratch in Go?**
Use a buffered channel `make(chan Msg, 1000)`.
*   **Limitations**: In-memory only (lost on restart).

### **358. How do you implement ordered message processing in Go?**
*   **Kafka**: Partitioning (messages for one key go to one partition).
*   **Go**: Ensure one consumer (goroutine) per "partition" or "entity ID".

### **359. How do you handle large stream ingestion (100K+ msgs/sec)?**
*   **Batching**: Read N messages, insert to DB in one batch.
*   **Pooling**: Reuse objects to reduce GC.

### **360. How do you persist in-flight streaming data?**
Write to **Write-Ahead Log (WAL)** on disk before processing.

## 🟣 **361–380: Cloud-Native and Distributed Systems in Go**

### **361. How do you build a cloud-agnostic app in Go?**
*   **12-Factor App**: Use Env vars for config (not SDK-specific logic if possible).
*   **Adapters**: Use interfaces for storage (`BlobStorage`) and msg queues (`PubSub`) to swap AWS/GCP implementations.
*   **Go Cloud Development Kit (Go CDK)**: `gocloud.dev` provides these portable APIs.

### **362. How do you use Go SDKs with AWS (S3, Lambda)?**
Import `github.com/aws/aws-sdk-go-v2`.
*   **Config**: `config.LoadDefaultConfig(ctx)`.
*   **Call**: `s3Client.PutObject(...)`.

### **363. How do you upload a file to S3 using Go?**
```go
s3Client.PutObject(ctx, &s3.PutObjectInput{
    Bucket: aws.String("bucket"),
    Key:    aws.String("key"),
    Body:   reader,
})
```

### **364. How do you create a Pub/Sub system using Go and GCP?**
Use `cloud.google.com/go/pubsub`.
*   Create Client.
*   `topic.Publish(ctx, msg)`.
*   `sub.Receive(ctx, func(ctx, msg) { ... })`.

### **365. How would you implement cloud-native config loading?**
*   **Kubernetes**: Read `ConfigMap` mounted as volume.
*   **Hot Reload**: Monitor file change (fsnotify) and re-read.

### **366. What is the role of service meshes with Go apps?**
(Istio, Linkerd).
*   Offloads mTLS, retries, tracing from the Go code to the sidecar proxy (Envoy).
*   Keeps Go app logic simple.

### **367. How do you secure service-to-service communication in Go?**
If no Service Mesh: **mTLS** (Mutual TLS).
*   Client presents cert, Server presents cert. Both verify CA.

### **368. How do you implement service registration and discovery?**
*   **Consul/Etcd**: Service registers itself on startup (puts IP in KV store). Clients watch KV store.
*   **K8s**: Done automatically via K8s Services (DNS).

### **369. How do you manage retries and circuit breakers in Go?**
*   **Retry**: `github.com/avast/retry-go`.
*   **Circuit Breaker**: `github.com/sony/gobreaker`.
    *   State: Closed (Normal) -> Open (Fail fast) -> Half-Open (Test recovery).

### **370. How would you use etcd/Consul with Go for KV storage?**
Use client libraries (`clientv3` for etcd).
*   **Etcd**: Strong consistency (Raft). Used for distributed config/coordination.

### **371. What is leader election and how can you implement it in Go?**
Ensuring only one instance performs a task.
*   **Implementation**: Use `etcd` concurrency API (Election) or K8s `Lease` API.
*   **Logic**: Candidates try to create a key with TTL. Winner keeps refreshing valid key.

### **372. How do you build a distributed lock in Go?**
**Redis (Redlock)** or **Etcd**.
*   Acquire: `SET resource_name my_id NX PX 30000`.
*   Release: Lua script checking if value is `my_id` then delete.

### **373. How would you implement a distributed queue in Go?**
Don't implement from scratch (hard). Use **NATS JetStream** or **Redis Streams**.
*   **Key**: Durability and Acknowledgement.

### **374. How do you handle consistency in distributed Go systems?**
*   **Strong**: Distributed TX (2PC - slow).
*   **Eventual**: Saga Pattern (Compensation). Design for "At Least Once" delivery.

### **375. How do you monitor and trace distributed Go systems?**
**Observability**:
*   **Trace**: OpenTelemetry (propagates traceparent header).
*   **Metrics**: RED method (Rate, Errors, Duration).

### **376. How do you implement eventual consistency in Go?**
*   Write to local DB.
*   Publish event "EntityCreated".
*   Subscribers update their view.
*   **Reconciliation**: Background job to fix drift.

### **377. How do you replicate state in distributed Go apps?**
**Raft Consensus Algorithm** (used by Etcd/CockroachDB).
*   Library: `hashicorp/raft`.
*   All writes go to Leader -> Logs replicated to Followers -> Commit.

### **378. How do you detect and handle split-brain scenarios?**
(Network partition where two subsets think they are leaders).
*   **Quorum**: Require >50% nodes to be active to accept writes.
*   **Fencing**: STONITH (Shoot The Other Node In The Head) or revoke token.

### **379. How do you implement quorum reads/writes in Go?**
Wait for responses from `(N/2)+1` nodes.
```go
wg.Add(N)
success := 0
// Launch N goroutines...
if atomic.LoadInt32(&success) > N/2 { return OK }
```

### **380. How would you build a simple distributed cache in Go?**
**Consistent Hashing**.
*   Hash key to ring (0-360 deg).
*   Node responsible for range.
*   Library: `groupcache` (by Memcached author, written in Go).

## 🔴 **381–400: Go in Real-World Projects & Architecture**

### **381. How do you handle config versioning in Go projects?**
*   **Flags/Env**: Deprecate old flags, support both for one release, then remove.
*   **Files**: Use a `version` field in YAML/JSON. Parse into versioned structs (`ConfigV1`, `ConfigV2`).

### **382. How do you organize API versioning in Go apps?**
*   **URL Path**: `/api/v1/users`, `/api/v2/users`.
*   **Packages**: `pkg/api/v1`, `pkg/api/v2` (Clean separation).
*   **Accept Header**: Content negotiation (cleaner URLs, harder to test). Path versioning is preferred in Go community.

### **383. How do you validate struct fields with custom rules?**
Use `go-playground/validator`.
*   Register custom validation function:
    ```go
    v.RegisterValidation("is-cool", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "cool"
    })
    ```

### **384. How do you cache API responses in Go?**
*   **Headers**: Set `Cache-Control` for client caching.
*   **Server**: Middleware using Redis. Key = URL + Params.

### **385. How do you serve files over HTTP with conditional GET?**
Use `http.ServeContent(w, r, name, modtime, content)`.
*   Automatically handles `If-Modified-Since` and `Range` requests.

### **386. How do you apply SOLID principles in Go?**
*   **SRP**: Small packages/functions.
*   **OCP**: Use Interfaces.
*   **LSP**: Interfaces ensure substitutability.
*   **ISP**: Small interfaces (`Reader`, `Writer`).
*   **DIP**: Dependency Injection (pass dependencies in constructor).

### **387. How do you prevent breaking changes in shared Go modules?**
*   **APIDiff**: Use `apidiff` tool to detect changes in exported identifiers.
*   **v2**: If breaking, release as v2 (go.mod sends `.../v2`).

### **388. What is the difference between horizontal and vertical scaling in Go services?**
*   **Vertical**: Larger instance (More CPU/RAM). Go scales well with GOMAXPROCS.
*   **Horizontal**: More instances. Requires stateless design (Session in Redis, not memory).

### **389. How do you support internationalization in Go?**
Use `golang.org/x/text` packages (`message`, `catalog`).
*   `p := message.NewPrinter(language.Greek)`
*   `p.Printf("Hello %s", name)`

### **390. How do you write a Go SDK for third-party APIs?**
*   **Client Struct**: Holds HTTP client + Config.
*   **Methods**: Map to API endpoints.
*   **Context**: All methods should accept `context.Context`.

### **391. How do you manage request IDs and trace IDs?**
Middleware generates UUID.
*   Put in `context`.
*   Inject into Logger.
*   Pass in HTTP headers (`X-Request-ID`) to downstream services.

### **392. How do you implement audit logging in Go?**
Middleware or "After" hooks.
*   Asynchronously write event (Who, What, When) to immutable storage (Append-only DB/Log).

### **393. How would you version a binary CLI in Go?**
Embed version at build time using `ldflags`.
`go build -ldflags "-X main.Version=1.0.0"`.

### **394. How do you ensure backward compatibility in Go libraries?**
*   Never remove exported functions.
*   Never change function signatures.
*   Add new functions (`NewFuncV2`) if behavior must change.

### **395. How do you handle soft deletes in Go models?**
Add `DeletedAt *time.Time` field.
*   Struct tag queries to filter where `deleted_at IS NULL`.
*   GORM handles this automatically with `gorm.Model`.

### **396. How do you refactor a large legacy Go codebase?**
**Strangler Fig Pattern**.
*   Write new microservice/module for a specific feature.
*   Route % of traffic to new code.
*   Gradually replace legacy chunks.

### **397. How do you maintain a mono-repo with multiple Go modules?**
*   **Workspaces (Go 1.18+)**: `go.work` file to develop across modules locally.
*   **Tooling**: `Bazel` or `Makefiles` to build changed targets only.

### **398. How would you go about building a plugin system in Go?**
*   **RPC (Preferred)**: Plugin runs as separate process (Hashicorp plugin). Crash safe.
*   **Native**: `plugin` package. Limited (Linux only, exact dependency match).

### **399. How do you document Go APIs automatically?**
*   **Swagger**: Use comments (`// @Summary ...`) and `swag init` tool to generate OpenAPI spec.

### **400. How do you track tech debt and enforce code quality in large Go teams?**
*   **CI**: Enforce `golangci-lint` (zero warnings policy).
*   **SonarQube**: Track complexity, duplication, and coverage trends over time.

## 🔵 **401–420: Networking and Low-Level Programming**

### **401. How do you create a TCP server in Go?**
Use `net.Listen`.
```go
ln, _ := net.Listen("tcp", ":8080")
for {
    conn, _ := ln.Accept()
    go handle(conn)
}
```

### **402. How do you create a UDP client in Go?**
Use `net.Dial`.
```go
conn, _ := net.Dial("udp", "127.0.0.1:8080")
conn.Write([]byte("hello"))
```

### **403. What is the difference between `net.Listen` and `net.Dial`?**
*   **Listen**: Opens a port and waits for incoming connections (Server).
*   **Dial**: Initiates a connection to a remote address (Client).

### **404. How do you manage TCP connection pools?**
Manually implement a pool of `net.Conn` objects (using a buffered channel) OR use `net/http` which handles pooling automatically for HTTP.

### **405. How would you implement a custom HTTP transport?**
Create a struct implementing `http.RoundTripper`.
```go
type MyTransport struct{}
func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) { ... }
```

### **406. How do you read raw packets using `gopacket`?**
Library: `github.com/google/gopacket`.
*   Uses `pcap` to sniff interface.
*   `packetSource.Packets()` returns a channel of packets.

### **407. What is a connection hijack in `net/http` and how is it done?**
Taking over the underlying TCP connection from the HTTP server (e.g., for WebSockets).
```go
hj, _ := w.(http.Hijacker)
conn, buf, _ := hj.Hijack()
```

### **408. How to implement a proxy server in Go?**
Use `httputil.ReverseProxy`.
```go
target, _ := url.Parse("http://backend")
proxy := httputil.NewSingleHostReverseProxy(target)
http.ListenAndServe(":8080", proxy)
```

### **409. How would you create an HTTP2 server from scratch in Go?**
`net/http` supports HTTP/2 automaticlly over TLS.
*   To force h2c (cleartext), use `golang.org/x/net/http2/h2c`.

### **410. How does Go handle connection reuse (keep-alive)?**
By default, `http.Client` and `Server` keep TCP connections open.
*   To disable: `req.Close = true` or `Transport.DisableKeepAlives = true`.

### **411. How do you set timeouts on sockets in Go?**
`conn.SetDeadline(time.Now().Add(5 * time.Second))`.
*   Can also set read/write deadlines separately.

### **412. What is the difference between `net/http` and `fasthttp`?**
*   **net/http**: Standard, compatible, allocates per request.
*   **fasthttp**: Third-party, zero-allocation, extremely fast, but API is incompatible with standard `http.Handler`.

### **413. How do you throttle network traffic in Go?**
Use `golang.org/x/time/rate`.
*   `limiter.Wait(ctx)` before reading/writing bytes.

### **414. How would you analyze network latency in Go?**
Use `httptrace`.
*   Hooks into: `DNSStart`, `DNSDone`, `ConnectStart`, `ConnectDone`.

### **415. How would you implement WebRTC or peer-to-peer comms?**
Use `pion/webrtc` (Pure Go implementation of WebRTC).

### **416. How do you simulate a slow network in integration tests?**
*   **Proxy**: Use a toxic proxy like `Toxiproxy`.
*   **Code**: Wrap `net.Conn` and add `time.Sleep` in `Read/Write`.

### **417. What’s the difference between connection pooling and multiplexing?**
*   **Pooling**: Reusing multiple connections (HTTP/1.1).
*   **Multiplexing**: Sending multiple requests over ONE connection (HTTP/2, QUIC).

### **418. How do you verify DNS lookups in Go?**
`net.LookupHost("google.com")` or `net.Resolver`.

### **419. How do you use HTTP pipelining in Go?**
You generally strictly don't. HTTP/1.1 pipelining is problematic. Use HTTP/2 multiplexing instead.

### **420. How do you implement NAT traversal in Go?**
*   **STUN/TURN**: Use `pion/turn`.
*   **UPnP**: Use `huin/goupnp` to open ports on routers.

## 🟣 **421–440: Error Handling & Observability**

### **421. How do you create custom error types in Go?**
Create a struct and implement the `error` interface (`Error() string`).
```go
type MyError struct { Msg string }
func (e *MyError) Error() string { return e.Msg }
```

### **422. How does Go 1.20+ `errors.Join` and `errors.Is` work?**
*   **Join**: Combines multiple errors into one. `err := errors.Join(err1, err2)`.
*   **Is**: Checks if an error chain contains a specific target error. `errors.Is(err, os.ErrNotExist)`.

### **423. How do you implement error wrapping and unwrapping?**
*   **Wrap**: `fmt.Errorf("context: %w", err)`.
*   **Unwrap**: `errors.Unwrap(err)`.

### **424. What are best practices for error categorization?**
Define sentinel errors (`ErrNotFound`) or error types (`ValidationError`).
*   Check with `errors.Is` or `errors.As`.

### **425. How do you handle critical vs recoverable errors?**
*   **Critical**: `log.Fatal` (exit app). Only in `main`.
*   **Recoverable**: Return error to caller. Log and continue.

### **426. How do you recover from panics in goroutines?**
Must call `recover()` inside a deferred function **within that specific goroutine**. Panics do not propagate across goroutines.

### **427. How to capture stack traces on error?**
Standard `error` doesn't have stack.
*   Use `github.com/pkg/errors` (deprecated but common) or custom error type with `runtime.Callers`.

### **428. How do you notify Sentry/Bugsnag from Go?**
Use their SDKs.
*   `defer sentry.Recover()` in main/handlers to catch panics.
*   `sentry.CaptureException(err)`.

### **429. How do you do structured error reporting in Go?**
Return errors with context map/attributes.
*   Log as JSON: `log.Error("msg", "error", err, "user_id", 123)`.

### **430. How do you correlate logs, errors, and traces together?**
Includes `TraceID` and `SpanID` in every log message and error report.

### **431. How would you add distributed tracing to an existing Go service?**
Install **OpenTelemetry** SDK.
*   Initialize TracerProvider.
*   Add Middleware to HTTP/gRPC servers.
*   Pass `context` everywhere.

### **432. What are tags, attributes, and spans in tracing?**
*   **Span**: A unit of work (e.g., "DB Query").
*   **Attribute**: Key-value data related to span ("db.table": "users").

### **433. What is a traceparent header?**
W3C standard header (`traceparent: version-traceid-parentid-flags`) to propagate context between services.

### **434. How do you send custom metrics to Prometheus?**
Define a `Counter` or `Gauge`.
```go
requests := promauto.NewCounter(prometheus.CounterOpts{Name: "app_requests_total"})
requests.Inc()
```

### **435. What is RED metrics model and how do you apply it?**
Standard for microservices:
*   **R**ate (Requests/sec).
*   **E**rrors (Failed/sec).
*   **D**uration (Latency P99).

### **436. How do you expose application health and readiness probes?**
*   `/healthz`: Returns 200 if app is running.
*   `/ready`: Returns 200 if app can accept traffic (DB connected).

### **437. What’s the difference between logs, metrics, and traces?**
*   **Logs**: Discrete events (debugging).
*   **Metrics**: Aggregatable numbers (trends/alerts).
*   **Traces**: Request lifecycle (performance/path).

### **438. How do you benchmark error impact on performance?**
Common error paths should be cheap.
*   Benchmark `fmt.Errorf` (allocates) vs simple string return.

### **439. What’s the tradeoff between verbose and silent error handling?**
*   **Verbose**: Good debugging, possibly noisy logs/leaking info.
*   **Silent**: Hard to debug, clean UI.
*   **Balance**: Log full details internally, return sanitized generic error to user.

### **440. How would you enforce observability in a Go microservice?**
Use a **Shared Library / Chassis** that bootstraps:
*   Logger (Zap).
*   Tracer (OTEL).
*   Metrics (Prometheus).
*   Middleware.

## 🟢 **441–460: CLI Tools, Automation, and Scripting**

### **441. How do you build an interactive CLI in Go?**
Use **Cobra** (structure) and **Promptui** / **Bubbletea** (UI).
*   **Bubbletea**: The Elm Architecture for terminal apps (TUI).

### **442. What libraries do you use for command-line tools in Go?**
*   **Struct**: `spf13/cobra`, `urfave/cli`.
*   **Flags**: `spf13/pflag` (POSIX style).
*   **Colors**: `fatih/color`.

### **443. How do you parse flags and config in CLI?**
*   **Viper**: Handles config files (YAML/JSON) + Environment + Flags.
*   **Pflag**: `pflag.String("name", "default", "help")`.

### **444. How do you implement bash autocompletion for Go CLI?**
Cobra has built-in support.
`rootCmd.GenBashCompletion(os.Stdout)`

### **445. How would you use `cobra` to build a nested command CLI?**
`rootCmd.AddCommand(childCmd)`.
*   Structure: `app user create` -> `root` > `user` > `create`.

### **446. How do you manage color and styling in terminal output?**
Use ANSI escape codes or libraries like `lipgloss` (modern) or `color` (simple).

### **447. How would you stream CLI output like `tail -f`?**
Read from a channel or file continuously and print to stdout.
```go
for line := range lines { fmt.Println(line) }
```

### **448. How do you handle secrets securely in a CLI?**
*   **Input**: `term.ReadPassword(fd)` (masks input).
*   **Storage**: Use OS Keyring (`99designs/keyring`).

### **449. How do you bundle a CLI as a standalone binary?**
Go builds are static by default.
`go build -o mytool .` -> Single file, no dependencies.

### **450. How would you version and release CLI with GitHub Actions?**
Use **GoReleaser**.
*   Action triggers on tag push.
*   Builds binaries, creates release, uploads artifacts.

### **451. How do you schedule a Go CLI tool with cron?**
*   **OS Level**: Add entry to system crontab (`crontab -e`).
*   **App Level**: Use `robfig/cron` to run jobs inside a long-running daemon.

### **452. How do you use Go as a scripting language?**
Use `gorun` (shebang support) or just `go run main.go`.
*   Go 1.22+ makes it easier to use `go run .`

### **453. How do you embed templates in your Go CLI tool?**
`//go:embed templates/*`.
*   Access via `embed.FS` at runtime. No external files needed.

### **454. How would you create a system daemon in Go?**
Use `kardianos/service`.
*   Handles install/uninstall/start/stop for Systemd (Linux), Launchd (Mac), SCM (Windows).

### **455. What are good patterns for CLI testing?**
*   Abstract `stdin/stdout` in your struct (`io.Reader`, `io.Writer`).
*   Test by passing `bytes.Buffer`.

### **456. How do you store and manage CLI state/config files?**
Store in `os.UserConfigDir()`.
*   Linux: `~/.config/myapp/`.
*   Mac: `~/Library/Application Support/myapp/`.

### **457. How do you secure a CLI for local system access?**
Run with least privilege.
*   Don't require `sudo` unless necessary.

### **458. How do you test CLI tools across multiple OS in CI?**
GitHub Actions matrix strategy.
```yaml
strategy:
  matrix:
    os: [ubuntu-latest, windows-latest, macos-latest]
```

### **459. How do you expose analytics and usage for a CLI?**
Send anonymous ping to a collection endpoint (Google Analytics / Custom) on execution. **Must** allow opt-out.

### **460. How would you build a CLI wrapper for REST APIs?**
*   Generate client from OpenAPI.
*   Map Cobra commands to API calls.
*   Format JSON output (support `--json` flag or `jq` friendly).

## 🟢 **461–480: AI, ML, and Data Science in Go**

### **461. Is Go suitable for machine learning?**
Yes, but less mature than Python.
*   **Pros**: Speed (inference), deployment (static binary), concurrency.
*   **Cons**: Smaller ecosystem (pandas/numpy equivalents are younger).

### **462. What libraries are used for ML in Go?**
*   **Gorgonia**: Deep learning (like PyTorch/TF).
*   **Gonum**: Numerical computing (matrices, linear algebra like NumPy).
*   **GoLearn**: Scikit-learn equivalent.

### **463. How do you integrate Python ML models with Go?**
*   **Inference Server**: Python serves HTTP/gRPC, Go calls it.
*   **ONNX**: Convert model to ONNX, run in Go using `onnx-go`.
*   **TensorFlow C API**: Bindings to run TF models directly in Go.

### **464. How does `gonum` compare to NumPy?**
Gonum provides matrix/vector math (`mat` package).
*   Performant (BLAS integration).
*   More verbose syntax than Python (`c.Mul(a, b)` vs `a * b`).

### **465. How do you implement a simple neural network in Go?**
Using `Gorgonia`.
Define graph (ExprGraph), Nodes, Activations (Sigmoid/ReLU), and Solver (SGD).

### **466. What is the role of Go in AI infrastructure?**
Deployment and Orchestration.
*   Kubernetes (written in Go) manages ML pods.
*   Inference APIs (low latency) often written in Go.

### **467. How do you perform data framing in Go (like Pandas)?**
Use `kotaoue/go-dataframe` or `go-gota/gota`.
*   Filtering, selecting columns, sorting.

### **468. How do you handle large datasets in Go for ML?**
Streaming.
*   Go's Reader/Writer interfaces allow processing TBs of data without loading all into RAM.

### **469. How do you run ONNX models in Go?**
Library: `github.com/owulveryck/onnx-go` or wrappers around C runtime.
*   Load `.onnx` file -> Create Backend -> Predict.

### **470. How do you use OpenAI API with Go?**
Official or community client.
Iterate over streams (`stream=True`) using Go channels for real-time tokens.

### **471. How do you process images for Computer Vision in Go?**
`image` standard library + `gocv` (OpenCV bindings).
*   `gocv.IMRead`, `gocv.GaussianBlu`, etc.

### **472. How do you implement NLP tokenization in Go?**
*   **Simple**: `strings.Split` or regex.
*   **Advanced**: `bleve` (search engine) has tokenizers. `jdkato/prose` (NLP library).

### **473. How do you visualize data in Go?**
`go-echarts/go-echarts` or `gonum/plot`.
Generates HTML/Images.

### **474. What is the advantage of using Go for ML inference?**
Latency and Footprint.
*   No Python interpreter overhead.
*   High concurrency (handle thousands of requests/sec).

### **475. How do you perform vector similarity search in Go?**
*   Calculuate Cosine Similarity (dot product of vectors).
*   Use libraries/DBs like **Milvus** (Go SDK) or simple in-memory search for small sets.

### **476. How do you implement genetic algorithms in Go?**
Define libraries for:
1.  Population (structs).
2.  Selection/Crossover/Mutation (methods).
3.  Concurrent evaluation (Goroutines for fitness calc).

### **477. How is specific hardware (GPU) accessed in Go?**
Through CGO (CUDA bindings).
*   **Gorgonia/Cu** provides CUDA support.

### **478. What is `bleve` and how is it used?**
A text indexing/search library (like Lucene).
*   Used for RAG (Retrieval Augmented Generation) context retrieval.

### **479. How do you handle time-series data in Go?**
Use specific DBs (Prometheus/InfluxDB) or simple ring buffers for windowed aggregation.

### **480. How would you build a recommender system service in Go?**
*   **Offline**: Python trains model (Matrix Factorization).
*   **Online**: Go loads user vectors + item vectors, computes top-k dot products fast.

## 🟢 **481–500: WebAssembly, Blockchain, and Advanced Topics**

### **481. Can you run Go in the browser?**
Yes, via **WebAssembly (WASM)**.
*   Go compiles to `.wasm` binary that runs in JS runtime.

### **482. How do you compile Go to WASM?**
`GOOS=js GOARCH=wasm go build -o main.wasm`.
*   Requires `wasm_exec.js` (found in GOROOT) to glue JS and Go.

### **483. How do you interact with DOM from Go (WASM)?**
Use `syscall/js`.
*   `js.Global().Get("document").Call("getElementById", "myBtn")`.

### **484. What is the difference between TinyGo and Go?**
*   **Go**: Standard Runtime, heavy binary for WASM (2MB+).
*   **TinyGo**: LLVM-based compiler for embedded/WASM. Very small binaries (10KB+). Best for WASM.

### **485. How do you use Go for Blockchain development?**
*   **Ethereum**: `go-ethereum` (geth) is the official Go client.
*   **Hyperledger Fabric**: Written in Go. Chaincode is written in Go.

### **486. How dows `geth` handle concurrency?**
Heavily relies on Goroutines for P2P networking, transaction propogation, and mining.

### **487. How do you implement a Smart Contract in Go (Hyperledger)?**
Implement `Chaincode` interface.
```go
func (s *SmartContract) Invoke(ctx contractapi.TransactionContextInterface) error { ... }
```

### **488. What is IPFS and how does Go relate to it?**
InterPlanetary File System.
*   The reference implementation (`kubo`) is written in Go.

### **489. How do you optimize Go binary size for embedded systems?**
*   Use `TinyGo`.
*   Strip symbols (`-s -w`).
*   Disable reflection if possible (removes metadata).

### **490. How do you manage hardware interrupts in Go (Embedded)?**
TinyGo supports `machine` package to handle GPIO interrupts.
```go
btn.Configure(machine.PinConfig{Mode: machine.PinInput})
btn.SetInterrupt(machine.PinRising, func(p machine.Pin) { ... })
```

### **491. How do you handle Hot Reloading in Go web apps?**
Use **Air**.
*   Watches files, rebuilds binary, restarts process on save.

### **492. What are Go plugins and their limitations?**
Load shared libraries (`.so`) at runtime.
*   **Limits**: Linux/Mac only. Host and Plugin must share exact dependencies versions.

### **493. How do you implement feature toggles/flags?**
*   **Simple**: Env vars / Config.
*   **Advanced**: LaunchDarkly SDK or `go-feature-flag`.

### **494. How do you implement search (Full-Text) in Go app?**
*   **Simple**: `strings.Contains` (for tiny data).
*   **Medium**: `Bleve` (embedded index).
*   **Large**: Elasticsearch / Meilisearch via client.

### **495. How do you generate PDFs in Go?**
Use `gofpdf` or `maroto`.
*   Programmatic layout definition.

### **496. How do you process Excel files in Go?**
Use `excelize`.
*   Read/Write `.xlsx` files efficiently.

### **497. How do you handle caching invalidation in Go?**
*   **TTL**: Expire after time.
*   **Events**: Listen to update events to clear keys.

### **498. What is the future of Go (Generics, Iterators)?**
*   **Generics**: Added in 1.18.
*   **Iterators (Range over func)**: Experimental in 1.22, likely standard in 1.23. allowing `for v := range mySeq`.

### **499. How do you stay updated with Go changes?**
*   Go Blog.
*   Release Notes (twice a year: Feb/Aug).
*   Proposals (GitHub issues).

### **500. Why choose Go over Rust or Java?**
*   **Vs Java**: Faster startup, lower memory, no JVM warmup, simple syntax.
*   **Vs Rust**: Faster compilation, easier learning curve (GC), better productivity for web services.

## 🔐 **501–520: Security in Golang**

### **501. How do you prevent SQL injection in Go?**
Use **Parameterized Queries** (`$1`, `?`).
*   `db.Query("SELECT * FROM users WHERE id=$1", id)`.
*   Avoid `fmt.Sprintf` for SQL construction.

### **502. How do you securely store user passwords in Go?**
Use **Argon2** or **Bcrypt**.
*   `bcrypt.GenerateFromPassword(pwd, cost)` stores hash + salt.
*   Never use MD5/SHA1.

### **503. How do you implement OAuth 2.0 in Go?**
Use `golang.org/x/oauth2` library.
*   Handles redirection to provider (Google/GitHub), token exchange, and refreshes.

### **504. What is CSRF and how do you prevent it in Go web apps?**
Cross-Site Request Forgery.
*   Use **CSRF Tokens** (middleware injects a unique token in forms/headers).
*   Set `SameSite=Strict` on cookies.

### **505. How do you use JWT securely in a Go backend?**
*   **Sign**: Use HS256 (HMAC) or RS256 (RSA).
*   **Validate**: Always check `alg` header matches expected method.
*   **Store**: HttpOnly Secure Cookies (mostly) or LocalStorage (if careful with XSS).

### **506. How do you validate and sanitize user input in Go?**
*   **HTML**: `bluemonday` library to sanitize HTML.
*   **Structs**: `go-playground/validator`.
*   **SQL**: Parameterization.

### **507. How do you set secure cookies in Go?**
```go
http.Cookie{
    Secure:   true, // HTTPS only
    HttpOnly: true, // No JS access
    SameSite: http.SameSiteStrictMode,
}
```

### **508. How do you avoid path traversal vulnerabilities?**
Use `filepath.Clean(path)` and verify it starts with absolute root directory.
*   Avoid `os.Open(userInputWithoutChecks)`.

### **509. How do you prevent XSS in Go HTML templates?**
`html/template` automatically escapes output by default.
*   Do NOT use `template.HTML(str)` on untrusted input.

### **510. How would you encrypt sensitive fields before storing in DB?**
Use **AES-GCM**.
*   Encrypt data in app layer before `INSERT`.
*   Manage keys (DEK/KEK) using a Key Management System (Vault).

### **511. How do you securely generate random strings or tokens?**
Use `crypto/rand`, NOT `math/rand`.
```go
b := make([]byte, 32)
rand.Read(b) // from crypto/rand
```

### **512. How do you verify digital signatures in Go?**
Use `crypto/rsa` or `crypto/ed25519`.
*   `rsa.VerifyPKCS1v15(pubKey, hash, signedData, signature)`.

### **513. What are best practices for TLS config in Go HTTP servers?**
*   Disable old versions (SSLv3, TLS 1.0, 1.1). Set `MinVersion: tls.VersionTLS12`.
*   Prefer modern cipher suites (`TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`).

### **514. How do you implement rate limiting in Go to avoid DDoS?**
**Token Bucket** (x/time/rate) per IP.
*   Middleware checks if IP bucket has tokens. Use Redis for distributed limiting.

### **515. How do you handle secrets in Go apps (Vault, env, etc.)?**
*   **Dev**: `.env` files.
*   **Prod**: Env vars injected by platform (K8s Secrets).
*   **Advanced**: Hashicorp Vault agents fetch secrets dynamically.

### **516. How do you perform mutual TLS authentication in Go?**
Configure `ClientAuth` in `tls.Config`.
```go
tls.Config{
    ClientAuth: tls.RequireAndVerifyClientCert,
    ClientCAs:  caPool,
}
```

### **517. What is the difference between `crypto/rand` and `math/rand`?**
*   **math/rand**: Pseudo-random (deterministic if seeded). Fast. Unsafe for secrets.
*   **crypto/rand**: Cryptographically Secure (OS entropy). Slow. Safe for secrets.

### **518. How do you prevent replay attacks using Go?**
Use **Nonces** (random numbers used once) and **Timestamps** in requests.
*   Verify signature covers timestamp. Reject messages older than N seconds.

### **519. How do you build a secure authentication system in Go?**
Combine:
1.  Bcrypt for passwords.
2.  MFA (TOTP) using `pquerna/otp`.
3.  JWT/PASETO for sessions.
4.  Secure Cookies.

### **520. How do you scan Go projects for vulnerabilities?**
*   `govulncheck ./...`
*   `trivy fs .` (Containers/Deps).
*   `gosec ./...` (Code analysis).
