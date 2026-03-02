# 100 Pure Code Snippet Interview Questions in Go

*This document contains strictly code-based "predict the output / spot the bug" questions, arranged from Basics to Advanced.*

---

## 🟢 Basics & Fundamentals

### 1. WaitGroup Copy Bug
**Q: What is the output or error?**
```go
package main
import "sync"

func wait(wg sync.WaitGroup) { 
    wg.Done() 
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go wait(wg)
    wg.Wait()
}
```
**A:** **Deadlock.** `sync.WaitGroup` is passed by value (copied). `wg.Wait()` blocks forever because `wg.Done()` modified a local copy in `wait`.
**Fix:** `func wait(wg *sync.WaitGroup)` and `go wait(&wg)`

### 2. Variable Shadowing
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 10
    if true {
        x := 5
        x++
    }
    fmt.Println(x)
}
```
**A:** `10`. The inner `x := 5` shadows the outer `x`.

### 3. Defer Argument Evaluation
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 1
    defer fmt.Println(i)
    i++
}
```
**A:** `1`. Arguments to a `defer` statement are evaluated immediately, not when the function returns.

### 4. Nil Map Trap
**Q: What is the output or error?**
```go
package main

func main() {
    var m map[string]int
    m["key"] = 1
}
```
**A:** **Panic.** Assignment to entry in nil map.
**Fix:** `m := make(map[string]int)`

### 5. Nil Slice Append
**Q: What is the output or error?**
```go
package main
import "fmt"

func main() {
    var s []int
    s = append(s, 1)
    fmt.Println(len(s))
}
```
**A:** `1`. Appending to a `nil` slice is perfectly valid in Go.

### 6. Value Receiver Modification
**Q: What is the output?**
```go
package main
import "fmt"

type Counter struct { val int }
func (c Counter) Increment() { c.val++ }

func main() {
    c := Counter{val: 0}
    c.Increment()
    fmt.Println(c.val)
}
```
**A:** `0`. `Increment` has a value receiver, so it modifies a copy.
**Fix:** `func (c *Counter) Increment()`

### 7. String Immutability
**Q: What is the output or error?**
```go
package main

func main() {
    s := "hello"
    s[0] = 'H'
}
```
**A:** **Compile Error.** Cannot assign to `s[0]`. Strings are perfectly immutable in Go.

### 8. Arrays vs Slices as Arguments
**Q: What is the output?**
```go
package main
import "fmt"

func modifyArray(a [3]int) { a[0] = 99 }
func modifySlice(s []int)  { s[0] = 99 }

func main() {
    arr := [3]int{1, 2, 3}
    modifyArray(arr)
    fmt.Println(arr[0])

    sl := []int{1, 2, 3}
    modifySlice(sl)
    fmt.Println(sl[0])
}
```
**A:** `1` then `99`. Arrays are copied by value. Slices pass a header that points to the same underlying array memory.

### 9. Fallthrough in Switch
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    switch 1 {
    case 1:
        fmt.Print("A")
        fallthrough
    case 2:
        fmt.Print("B")
    }
}
```
**A:** `AB`. `fallthrough` completely ignores the next case's condition and forcibly executes it.

### 10. Empty Struct Size
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println(unsafe.Sizeof(struct{}{}))
}
```
**A:** `0`. The empty struct conceptually consumes absolutely no memory in Go.

---

## 🟡 Intermediate

### 11. Loop Variable Capture (Go < 1.22)
**Q: What is the output historically (without Go 1.22+ fixes)?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for _, v := range []int{1, 2, 3} {
        wg.Add(1)
        go func() { 
            defer wg.Done()
            fmt.Print(v) 
        }()
    }
    wg.Wait()
}
```
**A:** Usually `333`. All goroutines capture the identical memory address of `v`.
**Fix:** `go func(val int) { ... }(v)`

### 12. Defer with Named Returns
**Q: What does the function return?**
```go
package main

func getResult() (res int) {
    defer func() { res += 5 }()
    return 10
}
```
**A:** `15`. The `return 10` sets `res = 10`, then the deferred function modifies the named return before finally exiting.

### 13. Map Iteration Randomness
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    for k := range m {
        fmt.Print(k)
    }
}
```
**A:** **Randomized.** (e.g., `123`, `231`, `312`). Go intentionally randomizes map iteration order natively.

### 14. Unbuffered Channel Deadlock
**Q: What is the output or error?**
```go
package main

func main() {
    ch := make(chan int)
    ch <- 1
    <-ch
}
```
**A:** **Fatal error: all goroutines are asleep - deadlock!** An unbuffered channel strictly requires a separate goroutine ready to read the exact moment you write to it.

### 15. The Append Overwrite Bug
**Q: What is the value of a, b, and c?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4}
    b := a[0:2]
    c := append(b, 99)
    fmt.Println(a)
}
```
**A:** `a = [1, 2, 99, 4]`. Because `b` is `[1, 2]` but still shares the capacity of `a`, appending to `b` directly overwrites index `2` inside `a`.

### 16. JSON Private Fields
**Q: What does json.Marshal produce?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name     string `json:"name"`
    password string `json:"password"`
}

func main() {
    u := User{"Alice", "secret"}
    b, _ := json.Marshal(u)
    fmt.Println(string(b))
}
```
**A:** `{"name":"Alice"}`. Unexported fields (`password`) are totally invisible to the `encoding/json` standard package.

### 17. Reading from Closed Channel
**Q: What is the output or error?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 5
    close(ch)
    
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```
**A:** `5` then `0`. You can drain pending items from closed channels. Once empty, it returns the zero value forever instead of blocking.

### 18. Writing to Closed Channel
**Q: What is the output or error?**
```go
package main

func main() {
    ch := make(chan int)
    close(ch)
    ch <- 1
}
```
**A:** **Panic.** Send on closed channel.

### 19. Nil Interface Value Trap
**Q: Does `err == nil` evaluate to true or false?**
```go
package main
import (
    "fmt"
    "os"
)

func do() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := do()
    fmt.Println(err == nil)
}
```
**A:** **False.** The interface holds a pointer value `nil`, but the interface's Type descriptor is populated (`*os.PathError`). An interface is only `nil` if BOTH type and value are strictly `nil`.

### 20. Struct Comparison
**Q: Does this compile?**
```go
package main
import "fmt"

type A struct { id int; tags []string }

func main() {
    a1 := A{1, nil}
    a2 := A{1, nil}
    fmt.Println(a1 == a2)
}
```
**A:** **Compile Error.** Structs containing fundamentally uncomparable fields (like slices or maps) cannot be compared using `==`.

### 21. Type Assertion Panic
**Q: What is the output or error?**
```go
package main
import "fmt"

func main() {
    var i interface{} = "hello"
    val := i.(int)
    fmt.Println(val)
}
```
**A:** **Panic.** Interface conversion: interface {} is string, not int.
**Fix:** `val, ok := i.(int)`

### 22. Select Default Case
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int)
    select {
    case ch <- 1:
        fmt.Println("Sent")
    default:
        fmt.Println("Dropped")
    }
}
```
**A:** `Dropped`. Unbuffered channel `ch` has no active receiver, so sending blocks. `select` instantly falls through to `default`.

### 23. Naked Returns masking
**Q: What is returned?**
```go
package main
import "fmt"

func test() (x int) {
    x = 10
    if true {
        x := 5
        _ = x
    }
    return
}
```
**A:** `10`. The `x := 5` block inner-shadows `x`. The naked return returns the outer-scoped `x` which remained `10`.

### 24. Modifying range iteration copies
**Q: What is printed?**
```go
package main
import "fmt"

func main() {
    words := []string{"a", "b"}
    for _, w := range words {
        w = "c"
    }
    fmt.Println(words)
}
```
**A:** `[a b]`. The variable `w` is strictly a local copy of the element slice. Modifying standard range loops does not inherently alter the slice.
**Fix:** `words[i] = "c"`

### 25. Nil Channel Select block
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var ch chan int // nil
    select {
    case <-ch:
        fmt.Println("Read")
    default:
        fmt.Println("Default")
    }
}
```
**A:** `Default`. Reading a `nil` channel blocks eternally, so `select` ignores that case naturally and executes `default`.

### 26. Custom Error implementation trap
**Q: Does `val` equal `nil`?**
```go
package main
import "fmt"

type MyErr struct{}
func (e *MyErr) Error() string { return "fail" }

func get() error {
    return MyErr{}
}

func main() {
    fmt.Println(get() == nil)
}
```
**A:** **Compile Error.** `MyErr` itself does not implement `error`, only the pointer `*MyErr` does! `get()` must return `&MyErr{}`.

### 27. Panic inside Recover
**Q: Does this save the application?**
```go
package main
import "fmt"

func main() {
    defer func() {
        func() { recover() }()
    }()
    panic("Crash")
}
```
**A:** **No, it crashes.** `recover()` MUST be called directly within the deferred function itself, not natively nested deeper inside anonymous layers.

### 28. String Byte conversion allocation
**Q: Does `for _ = range []byte(str)` allocate memory?**
```go
package main

func main() {
    s := "long string"
    for i, b := range []byte(s) {
        // ...
    }
}
```
**A:** **No.** As a native compiler optimization specifically for literal `for range []byte(str)`, Go skips memory allocation and purely reads the string's backing array natively.

### 29. Recovering from Goroutines
**Q: Does the main panics?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    defer func() { recover() }()
    go func() {
        panic("Goroutine panic")
    }()
    time.Sleep(time.Second)
    fmt.Println("App lives")
}
```
**A:** **App crashes.** `defer recover()` natively protects solely the goroutine it was called natively in. Panics in spawned goroutines instantly terminate the full application.

### 30. Empty Struct Memory Alignment
**Q: Which struct is strictly smaller?**
```go
package main
import "unsafe"

type A struct {
    a bool
    b struct{}
}

type B struct {
    b struct{}
    a bool
}
```
**A:** `B` is typically smaller. If an empty struct natively terminates a larger struct natively, the compiler artificially pads `A` to ensure pointers native don't overflow unallocated boundaries natively!

---

## 🔴 Concurrency Deep Dive & Context

### 61. Forcing Garbage Collection
**Q: Does `runtime.GC()` immediately free all unused memory?**
```go
package main
import "runtime"

func main() {
    runtime.GC()
}
```
**A:** **No.** `runtime.GC()` forces a blocking garbage collection cycle to identify and free unused objects back to the Go runtime allocator. However, the exact timing of returning that freed physical RAM back to the host Linux/OS layer is delayed and scavenged over time to avoid excessive System Calls.

### 62. Nil Map vs Empty Map initialization
**Q: Are these completely identical?**
```go
package main
import "fmt"

func main() {
    var m1 map[int]int
    m2 := map[int]int{}
    fmt.Println(m1 == nil, m2 == nil)
}
```
**A:** `true false`. `m1` is a `nil` map (writing panics). `m2` is an empty, fully initialized map ready to accept values.

### 63. Defer panic overrides
**Q: What is returned?**
```go
package main

func do() (x int) {
    defer func() {
        if recover() != nil {
            x = 2
        }
    }()
    panic("boom")
    return 1
}
```
**A:** `2`. The deferred function successfully recovers the panic and cleanly alters the dynamically named return variable before exiting.

### 64. Struct Field iteration
**Q: Does this compile?**
```go
package main

type User struct { a, b int }

func main() {
    u := User{1, 2}
    for val := range u { }
}
```
**A:** **No.** You cannot explicitly `range` loop struct fields natively. Natively, `range` loops strictly apply to Arrays, Slices, Maps, Strings, and Channels. To loop fields, you MUST use the `reflect` package dynamically.

### 65. The Select {} Spin
**Q: Will this rapidly overheat the CPU core?**
```go
package main

func main() {
    select {}
}
```
**A:** **No.** The specific idiom `select {}` acts as a perfect block. The Go Scheduler parks the Goroutine forever, consuming exactly 0% CPU. (Unlike a `for {}` loop which consumes exactly 100% of a CPU core).

### 66. Context value passing
**Q: What does this pattern commonly violate?**
```go
package main
import "context"

func connect(ctx context.Context) {
    db := ctx.Value("dbConnection")
}
```
**A:** It violates Go idioms. `context.WithValue` should strictly exclusively be used for request-scoped metadata (like Trace IDs / User IPs), never for passing required business dependencies like Database Pointers, because it bypasses static type safety entirely.

### 67. The `int` size
**Q: Is `int` always explicitly 32 bits?**
```go
package main
import (
    "fmt"
    "unsafe"
)

func main() {
    var a int
    fmt.Println(unsafe.Sizeof(a))
}
```
**A:** **No.** `int` dynamically adapts to the architecture. On a 64-bit OS, it is 64-bits (8 bytes). On a 32-bit OS, it is 32-bits (4 bytes).

### 68. Global Variable initialization Order
**Q: What is the output?**
```go
package main
import "fmt"

var a = b + 1
var b = 1

func main() {
    fmt.Println(a, b)
}
```
**A:** `2 1`. The Go compiler is smart enough to resolve dependency graphs natively during initialization regardless of the visual top-to-bottom variable layout.

### 69. Method Expressions
**Q: Is this syntax valid natively?**
```go
package main
import "fmt"

type T int
func (t T) Print() { fmt.Println(t) }

func main() {
    f := T.Print
    f(10)
}
```
**A:** **Yes.** This is called a Method Expression natively. It essentially detaches the method from strictly bound receivers, turning the receiver specifically into explicitly the first argument smoothly.

### 70. Mutex Map concurrency natively
**Q: Will this perfectly prevent race conditions?**
```go
package main
import "sync"

type SafeMap struct {
    mu sync.Mutex
    data map[int]int
}

func (s *SafeMap) Get() map[int]int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.data
}
```
**A:** **No!** Returning exactly the internal isolated map pointer outside the Mutex lock boundaries absolutely allows the caller to mutate it without locking securely, immediately destroying all thread-safety. You must strictly copy values or hide the map structure securely.

### 71. String pointers to interfaces
**Q: Is this specifically required natively?**
```go
package main
import "io"

func print(r io.Reader) {}

func main() {
    var r *io.Reader
    print(r)
}
```
**A:** **Compile Error.** You exclusively rarely use `*io.Reader`. An interface is intrinsically distinctly effectively natively a pointer strictly containing both exactly exactly type and value pointers securely.

### 72. Recover placement
**Q: What will happen?**
```go
package main
import "fmt"

func safe() {
    recover()
    panic("Die")
}

func main() { safe() }
```
**A:** **Panic.** `recover()` must be called directly within a `defer` function. Calling it normally does nothing to stop a panic.

### 73. Slice pointer evaluation
**Q: Will this print 1 or panic?**
```go
package main
import "fmt"

func main() {
    var s *[]int
    fmt.Println(len(*s))
}
```
**A:** **Panic.** `s` is a `nil` pointer to a slice. Dereferencing a `nil` pointer (`*s`) panics before `len()` can even evaluate it.

### 74. Modifying Map Struct Values
**Q: Does this compile?**
```go
package main

type Math struct { x int }

func main() {
    m := map[string]Math{"A": {1}}
    m["A"].x = 2
}
```
**A:** **Compile Error.** Map values are unaddressable in Go. You cannot modify a struct field directly inside a map. You must reassign the whole struct.

### 75. Interface nil comparison trap
**Q: What is the output?**
```go
package main
import "fmt"

type CustomErr struct{}
func (c *CustomErr) Error() string { return "Err" }

func check() error {
    var e *CustomErr
    return e
}

func main() {
    err := check()
    fmt.Println(err == nil)
}
```
**A:** `false`. The returned `error` interface contains a `nil` pointer value, but its type is `*CustomErr`. An interface is only `nil` if both type and value are `nil`.

---

## ⚫ Expert & Internals

### 76. Map iteration pointer value
**Q: What is the risk here?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "A", 2: "B"}
    var ptrs []*string
    for _, v := range m {
        ptrs = append(ptrs, &v)
    }
    fmt.Println(*ptrs[0], *ptrs[1])
}
```
**A:** (Pre Go 1.22) It prints the same value twice (e.g., `B B`). `v` is a single variable re-used in every iteration. You are appending the exact same memory address `&v` to the slice.
**Fix:** In Go 1.22+, this is fixed explicitly natively, but broadly you should use `v := v` inside the loop or `append(ptrs, &m[k])` natively.

### 77. The Go 1.22 loop variable scoping
**Q: In Go 1.22, what does this print?**
```go
package main
import "fmt"

func main() {
    funcs := []func(){}
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() { fmt.Print(i) })
    }
    for _, f := range funcs { f() }
}
```
**A:** `012`. Go 1.22 explicitly changed `for` loops so that `i` is a distinctly new, unique variable per iteration smoothly. (Historically, it would print `333`).

### 78. RWMutex Upgrading (Deadlock)
**Q: Does this block?**
```go
package main
import "sync"

func main() {
    var rw sync.RWMutex
    rw.RLock()
    rw.Lock()
    rw.Unlock()
    rw.RUnlock()
}
```
**A:** **Yes, Deadlock.** Standard `sync.RWMutex` absolutely does not support natively "upgrading" a read-lock directly into a write-lock safely without entirely dropping it first properly.

### 79. Slice Memory Leak (String Slicing)
**Q: How much memory is realistically retained?**
```go
package main
import "fmt"

var saved string

func process(hugeFile string) {
    saved = hugeFile[:5]
}

func main() {
    process("1GB worth of string data...")
    fmt.Println(saved)
}
```
**A:** **The entire 1GB.** Taking a sub-slice of a string dynamically safely creates a tiny structural header referencing the exact original massive memory array safely organically. It never logically cleanly natively frees it natively until `saved` dies.
**Fix:** `strings.Clone(hugeFile[:5])` perfectly exclusively smartly securely allocates cleanly natively.

### 80. Mutex Unlocking
**Q: Why does this panic?**
```go
package main
import "sync"

func main() {
    var m sync.Mutex
    m.Lock()
    m.Unlock()
    m.Unlock()
}
```
**A:** **Panic.** `sync: unlock of unlocked mutex`. Unlocking a mutex twice causes a runtime panic.

### 81. RWMutex Write Starvation
**Q: If Goroutine A holds an RLock, Goroutine B requests a Lock, and Goroutine C requests an RLock, what happens to C?**
**A:** Goroutine C blocks. Go's RWMutex prevents write starvation: once a writer requests a lock, all subsequent reader requests are queued until the writer completes.

### 82. Interface Embedding
**Q: Does this compile?**
```go
package main

type Reader interface { Read() }
type Writer interface { Write() }
type ReadWriter interface {
    Reader
    Writer
}

func main() {}
```
**A:** **Yes.** Interfaces can embed other interfaces perfectly.

### 83. Struct Tag Reflection
**Q: How do you read a struct tag?**
```go
package main
import (
    "fmt"
    "reflect"
)

type User struct { Name string `json:"name"` }

func main() {
    u := User{}
    t := reflect.TypeOf(u)
    field, _ := t.FieldByName("Name")
    fmt.Println(field.Tag.Get("json"))
}
```
**A:** `name`. You use the `reflect` package to inspect the `Type` and extract the tag string.

### 84. Unsafe Pointer Conversion
**Q: Is this safe?**
```go
package main
import (
    "fmt"
    "unsafe"
)

func main() {
    i := 10
    var f *float64 = (*float64)(unsafe.Pointer(&i))
    fmt.Println(*f)
}
```
**A:** **No.** While it compiles, casting an `int` pointer to a `float64` pointer and dereferencing it yields meaningless scientific notation values because their bit-level memory formats (IEEE 754 vs Two's Complement) differ entirely.

### 85. Empty Channel Read
**Q: What happens if you read from an unbuffered channel that has no writers?**
```go
package main

func main() {
    ch := make(chan int)
    <-ch
}
```
**A:** **Deadlock.** The goroutine blocks forever waiting for a value.

### 86. Directional Channel Conversion
**Q: Does this compile?**
```go
package main

func send(ch chan<- int) { ch <- 1 }

func main() {
    ch := make(chan int)
    send(ch)
}
```
**A:** **Yes.** A bidirectional channel (`chan int`) implicitly converts to a send-only channel (`chan<- int`) when passed to a function.

### 87. Nil Context
**Q: Will passing `nil` as a Context panic?**
```go
package main
import "context"

func do(ctx context.Context) {}

func main() {
    do(nil)
}
```
**A:** It compiles and runs the function, but if `do()` attempts to call methods on the context (like `ctx.Done()`), it will panic. Always pass `context.TODO()` or `context.Background()`.

### 88. Context Cancellation Propagation
**Q: If you cancel a parent context, what happens to child contexts?**
**A:** All child contexts derived from that parent are immediately canceled as well. The cancellation cascades down the tree.

### 89. Error Wrapping (Go 1.13+)
**Q: How do you check for a specific error if it was wrapped?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
    err := fmt.Errorf("db load: %w", ErrNotFound)
    fmt.Println(errors.Is(err, ErrNotFound))
}
```
**A:** `true`. `errors.Is` unwraps the error chain to see if the target error exists anywhere within it.

### 90. HTTP Client Connection Reuse
**Q: What must you do to ensure `http.Client` reuses TCP connections?**
```go
package main
import (
    "net/http"
    "io"
)

func main() {
    resp, _ := http.Get("http://example.com")
    io.ReadAll(resp.Body)
    // Missing step?
}
```
**A:** You must call `resp.Body.Close()`. If the body is fully read and closed, the connection returns to the keep-alive pool.

### 91. Empty Interface Memory Allocation
**Q: Does assigning an integer to `interface{}` cause an allocation?**
```go
package main

func main() {
    var i interface{} = 5
    _ = i
}
```
**A:** **Yes.** Historically, boxing value types (like `int`) into an `interface{}` caused a heap allocation. Go has optimized this for small integers in recent versions, but generally, converting concrete types to interfaces allocates.

### 92. The `init()` Function
**Q: Can you call `init()` manually?**
```go
package main
import "fmt"

func init() { fmt.Println("Init") }

func main() { init() }
```
**A:** **Compile Error.** `init()` functions cannot be referenced or called manually. They are executed automatically by the Go runtime before `main()` begins.

### 93. Multiple `init()` Functions
**Q: Can a single file have multiple `init()` functions?**
```go
package main
import "fmt"

func init() { fmt.Print("1") }
func init() { fmt.Print("2") }

func main() {}
```
**A:** **Yes.** A single file can define multiple `init()` functions. They execute in the exact order they appear in the file.

### 94. Constant Typing
**Q: Is `A` a float or an int?**
```go
package main
import "fmt"

const A = 5.0

func main() {
    var b int = A
    fmt.Printf("%T\n", b)
}
```
**A:** `int`. `A` is an *untyped* floating-point constant. When assigned to an `int` variable `b`, the compiler implicitly converts it because `5.0` has no fractional part.

### 95. Range over String vs Byte Slice
**Q: What is the difference in output?**
```go
package main
import "fmt"

func main() {
    s := "ü" // 2 bytes
    for _, v := range s { fmt.Printf("%T ", v) }
    for _, v := range []byte(s) { fmt.Printf("%T ", v) }
}
```
**A:** `int32 uint8 uint8`. Ranging a string yields `rune` (`int32`). Ranging `[]byte` yields raw bytes (`uint8`).

### 96. Closure Variable Capture
**Q: What does this print?**
```go
package main
import "fmt"

func main() {
    x := 1
    f := func() { fmt.Println(x) }
    x = 2
    f()
}
```
**A:** `2`. The closure captures the variable `x` by reference, not by value at the time of creation.

### 97. Panic vs os.Exit
**Q: Does `os.Exit(1)` run deferred functions?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("Deferred")
    os.Exit(1)
}
```
**A:** **No.** `os.Exit` terminates the program immediately, bypassing all `defer` statements. A `panic` will execute `defer` statements as it unwinds the stack.

### 98. Goroutine Scheduling Point
**Q: How does Go preempt a goroutine running an infinite calculation?**
```go
package main

func main() {
    go func() { for {} }()
    // ...
}
```
**A:** Since Go 1.14, the scheduler uses asynchronous preemption via OS signals (SIGURG). It can interrupt tight CPU loops even if they contain no function calls.

### 99. The Type Assertion Second Return Value
**Q: What happens if you omit the `ok` boolean in a failed type assertion?**
```go
package main

func main() {
    var i interface{} = "hello"
    val := i.(int)
    _ = val
}
```
**A:** **Panic.** Omitting the `ok` boolean (e.g., `val, ok := i.(int)`) transforms the assertion from a safe check into a hardened demand. If the type is wrong, it panics.

### 100. Build Tags
**Q: What does this comment do at the top of a file?**
`//go:build linux`
**A:** It instructs the `go build` compiler to completely ignore the file unless the target operating system (`GOOS`) is strictly set to `linux`.
