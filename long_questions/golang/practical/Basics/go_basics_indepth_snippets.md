# 100 In-Depth Go Basics & Fundamentals — Pure Code Snippet Questions

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Focus**: Subtle edge cases, gotchas, and deeper mechanics of Go's type system, memory model, and runtime.

---

## Section 1: Variables, Scope & Type System Deep Dives (Q1–Q18)

### 1. Short Declaration Reuse in Same Scope
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    x, y := 1, 2
    x, z := 3, 4
    fmt.Println(x, y, z)
}
```
**A:** **Yes, compiles and prints** `3 2 4`. `:=` is valid when at least one variable on the left is **new** (`z`). It reassigns `x` rather than redeclaring it.

---

### 2. Redeclaration in Inner Scope Gets Separate Variable
**Q: What exactly is printed and why?**
```go
package main
import "fmt"

func main() {
    x := 1
    {
        x, y := 2, 3
        fmt.Println(x, y)
    }
    fmt.Println(x)
}
```
**A:**
```
2 3
1
```
The inner block's `x, y := 2, 3` creates a **brand-new** `x` (because it is a new scope), shadowing the outer one. The outer `x` remains `1`.

---

### 3. Constant Expression Evaluation
**Q: What is the output?**
```go
package main
import "fmt"

const (
    a = 2
    b = 3
    c = a * b + 1
)

func main() {
    fmt.Println(c)
}
```
**A:** `7`. Constant expressions are evaluated at compile time with arbitrary-precision arithmetic.

---

### 4. Typed vs Untyped Constant Assignment
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    const x int = 5
    var y float64 = x
    fmt.Println(y)
}
```
**A:** **Compile Error.** `x` is a **typed** constant of type `int`, so it cannot be assigned to `float64` without an explicit conversion. Compare: an **untyped** `const x = 5` would work fine.

---

### 5. Type Alias Shares Method Set
**Q: Does this compile?**
```go
package main
import "fmt"

type MyString = string  // alias, not new type

func main() {
    var s MyString = "hello"
    var t string = s
    fmt.Println(t)
}
```
**A:** **Yes, compiles and prints** `hello`. A type alias (`=`) is the exact same type, so no conversion is needed.

---

### 6. Named Type Does NOT Share Method Set With Underlying Type
**Q: Does this compile?**
```go
package main

type MyInt int

func double(n int) int { return n * 2 }

func main() {
    x := MyInt(5)
    _ = double(x)
}
```
**A:** **Compile Error.** `MyInt` is a distinct type from `int`. You must convert: `double(int(x))`.

---

### 7. Multiple Return Ignore with Blank
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi("42")
    fmt.Printf("%T %v\n", n, n)
}
```
**A:** `int 42`

---

### 8. Three-Component for and Post Statement
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i += 2 {
        fmt.Print(i, " ")
    }
}
```
**A:** `0 2 `. The post-statement is `i += 2`, so it goes `0 → 2 → 4` (stops because 4 ≥ 3).

---

### 9. Zero Value of a Function Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var f func()
    fmt.Println(f == nil)
}
```
**A:** `true`. The zero value of a function type is `nil`.

---

### 10. Comparing Function Types
**Q: Does this compile?**
```go
package main

func main() {
    f := func() {}
    g := func() {}
    _ = f == g
}
```
**A:** **Compile Error.** Function values are not comparable with `==` in Go (except to `nil`).

---

### 11. Integer Literal Representations
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    dec := 42
    hex := 0x2A
    oct := 0o52
    bin := 0b101010
    fmt.Println(dec, hex, oct, bin)
}
```
**A:** `42 42 42 42`. All four literals represent the same value.

---

### 12. Floating-Point Precision Comparison
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := 0.1 + 0.2
    fmt.Println(a == 0.3)
    fmt.Printf("%.17f\n", a)
}
```
**A:**
```
false
0.30000000000000004
```
IEEE 754 floating-point arithmetic is imprecise. Never compare floats with `==`.

---

### 13. var Block Dependency Order
**Q: Does this compile?**
```go
package main
import "fmt"

var (
    x = y + 1
    y = 2
)

func main() {
    fmt.Println(x, y)
}
```
**A:** **Yes, compiles and prints** `3 2`. Package-level `var` declarations are resolved by the compiler regardless of order.

---

### 14. := in if Condition and Scope
**Q: Is `err` accessible after the if-else block?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    if n, err := strconv.Atoi("bad"); err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("value:", n)
    }
    // fmt.Println(n) // Compile error: n undefined
}
```
**A:** No. Both `n` and `err` are scoped to the `if-else` block. The print of `error: strconv...` executes. Uncommenting the last line would cause a compile error.

---

### 15. Constant Converted to String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    const n = 65
    s := string(n)
    fmt.Println(s)
}
```
**A:** `A`. `string(65)` treats the integer as a Unicode code point (rune), producing the character `A`.

---

### 16. Shadowing Built-in Identifier
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    len := 42
    fmt.Println(len)
}
```
**A:** `42`. Go allows shadowing built-in identifiers like `len`, `make`, `new`, etc. (this is generally a bad practice but is valid).

---

### 17. Untyped Numeric Constant Default Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    const x = 3
    fmt.Printf("%T\n", x)
}
```
**A:** `int`. An untyped integer constant defaults to `int` when its type needs to be determined (e.g., as an argument to `%T`).

---

### 18. Package-Level Init Ordering
**Q: What is the output?**
```go
package main
import "fmt"

var (
    msg   = greet()
    hello = "Hello"
)

func greet() string { return hello + " World" }

func main() {
    fmt.Println(msg)
}
```
**A:** ` World`. `hello` is `""` (zero value) when `greet()` runs because `msg = greet()` is initialized first. Package-level variables are initialized in declaration order, not dependency order for cross-variable references here.

---

## Section 2: Control Flow Deep Dives (Q19–Q30)

### 19. Labeled Continue
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue outer
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
}
```
**A:** `(0,0) (1,0) (2,0) `. When `j == 1`, `continue outer` skips the rest of the inner loop AND the rest of the outer loop body, going straight to the next outer iteration.

---

### 20. Switch Case with Initialization
**Q: What is the output?**
```go
package main
import "fmt"

func val() int { return 2 }

func main() {
    switch v := val(); {
    case v == 1:
        fmt.Println("one")
    case v == 2:
        fmt.Println("two")
    default:
        fmt.Println("other")
    }
}
```
**A:** `two`. The `switch v := val();` form initializes `v` before the switch. The trailing `;` with no expression means `switch true`, making each case a boolean expression.

---

### 21. For Range Over Channel
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for v := range ch {
        fmt.Print(v, " ")
    }
}
```
**A:** `1 2 3 `. `for range` over a channel receives until the channel is closed and drained.

---

### 22. Fallthrough Cannot Be Last in Case
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    switch 1 {
    case 1:
        fallthrough
    case 2:
        fmt.Println("matched")
    case 3:
        fmt.Println("three")
        fallthrough
    }
}
```
**A:** **Compile Error.** `fallthrough` in the last case (`case 3`) is not allowed. There is no next case to fall into.

---

### 23. Break Out of Select Inside For
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 42
    for {
        select {
        case v := <-ch:
            fmt.Println(v)
            break // breaks the select, NOT the for
        }
        break // this breaks the for
    }
}
```
**A:** `42`. The first `break` exits only the `select`. The second `break` exits the `for` loop. This is a subtle trap: `break` inside `select` only breaks the `select`.

---

### 24. Blank Switch (switch true)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 25
    switch {
    case x > 100:
        fmt.Println("huge")
    case x > 10:
        fmt.Println("big")
    case x > 0:
        fmt.Println("positive")
    }
}
```
**A:** `big`. Cases are evaluated top-to-bottom; the first true one wins.

---

### 25. Range Loop Variable Reuse
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    funcs := make([]func(), 3)
    for i := 0; i < 3; i++ {
        funcs[i] = func() { fmt.Print(i) }
    }
    for _, f := range funcs {
        f()
    }
}
```
**A:** `333`. All closures capture the same `i`; after the loop, `i == 3`.  
**Fix:** `j := i; funcs[i] = func() { fmt.Print(j) }`

---

### 26. For with Post Increment Inside Body
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for i < 5 {
        i++
        if i == 3 {
            continue
        }
        fmt.Print(i, " ")
    }
}
```
**A:** `1 2 4 5 `. When `i == 3`, `continue` skips the `fmt.Print` but `i` was already incremented.

---

### 27. Nested Switch Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        switch i {
        case 1:
            fmt.Println("one")
            break // breaks switch, not for
        }
        fmt.Println("after switch:", i)
    }
}
```
**A:**
```
after switch: 0
one
after switch: 1
after switch: 2
```
`break` inside `switch` exits only the switch, not the surrounding `for`.

---

### 28. Range Index on Map
**Q: Is this valid and what does it print?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    for k := range m {
        fmt.Println(k)
    }
}
```
**A:** **Valid.** Prints `a` and `b` (order not guaranteed). When ranging a map, a single variable gets the key.

---

### 29. for init; ; post — Missing Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; ; i++ {
        if i == 3 {
            break
        }
        fmt.Print(i)
    }
}
```
**A:** `012`. Omitting the condition in a three-part `for` creates an infinite loop equivalent to `for true`.

---

### 30. Defer Runs After Return Value Set
**Q: What does `run()` return?**
```go
package main
import "fmt"

func run() int {
    x := 0
    defer func() { x = 100 }()
    return x
}

func main() {
    fmt.Println(run())
}
```
**A:** `0`. Unlike named returns, the deferred closure modifies a local copy `x`. The `return x` already captured `0`. Contrast with a named return where the defer CAN modify the result.

---

## Section 3: Functions & Closures Deep Dives (Q31–Q45)

### 31. Method Value vs Method Expression
**Q: What is the output?**
```go
package main
import "fmt"

type Adder struct{ base int }
func (a Adder) Add(n int) int { return a.base + n }

func main() {
    adder := Adder{base: 10}

    // Method value: bound to adder instance
    addFn := adder.Add
    fmt.Println(addFn(5))

    // Method expression: unbound, receiver passed explicitly
    addExpr := Adder.Add
    fmt.Println(addExpr(adder, 5))
}
```
**A:**
```
15
15
```

---

### 32. Passing Variadic to Variadic
**Q: What is the output?**
```go
package main
import "fmt"

func inner(nums ...int) {
    fmt.Println(nums)
}

func outer(nums ...int) {
    inner(nums...)
}

func main() {
    outer(1, 2, 3)
}
```
**A:** `[1 2 3]`. `nums...` unpacks the slice and passes it directly without creating a new slice.

---

### 33. Defer Captures Value vs Reference
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 1
    defer func() { fmt.Println("ref:", x) }()  // captures by reference
    defer fmt.Println("val:", x)                  // arg evaluated NOW (value)
    x = 99
}
```
**A:**
```
val: 1
ref: 99
```
The second `defer` line evaluates `x` immediately (value=1). The first closure sees the final value of `x` (99) when it runs.

---

### 34. Recursive Anonymous Function
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var fact func(n int) int
    fact = func(n int) int {
        if n == 0 {
            return 1
        }
        return n * fact(n-1)
    }
    fmt.Println(fact(5))
}
```
**A:** `120`

---

### 35. Functions Returning Functions (Currying)
**Q: What is the output?**
```go
package main
import "fmt"

func multiply(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    triple := multiply(3)
    fmt.Println(triple(4))
    fmt.Println(triple(10))
}
```
**A:**
```
12
30
```

---

### 36. Function Signature Mismatch
**Q: Does this compile?**
```go
package main
import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    var f func(string) int = greet
    fmt.Println(f("Go"))
}
```
**A:** **Compile Error.** `greet` returns `string` but `f` expects a function returning `int`. Function types must match exactly.

---

### 37. Multiple Defers and Named Return Interaction
**Q: What does `g()` return?**
```go
package main
import "fmt"

func g() (result int) {
    defer func() { result++ }()
    defer func() { result += 10 }()
    return 1
}

func main() {
    fmt.Println(g())
}
```
**A:** `12`. `return 1` sets `result = 1`. Defers run in LIFO: first `result += 10` → `11`, then `result++` → `12`.

---

### 38. Panicking in a Called Function
**Q: What is the output?**
```go
package main
import "fmt"

func risky() {
    panic("boom")
}

func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    risky()
    fmt.Println("after risky") // unreachable
}

func main() {
    safe()
    fmt.Println("main continues")
}
```
**A:**
```
recovered: boom
main continues
```
`recover()` stops the panic, `safe()` returns normally, and `main` continues.

---

### 39. init() Function
**Q: What is the output?**
```go
package main
import "fmt"

var x int

func init() {
    x = 42
    fmt.Println("init ran, x =", x)
}

func main() {
    fmt.Println("main ran, x =", x)
}
```
**A:**
```
init ran, x = 42
main ran, x = 42
```
`init()` always runs before `main()`.

---

### 40. Multiple init() Functions in One File
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

func init() { fmt.Println("init 1") }
func init() { fmt.Println("init 2") }

func main() { fmt.Println("main") }
```
**A:** **Yes, compiles.** Output:
```
init 1
init 2
main
```
A package can have multiple `init()` functions; they run in source order.

---

### 41. Blank Identifier to Suppress Unused Import
**Q: Does this compile?**
```go
package main
import _ "fmt"

func main() {}
```
**A:** **Yes.** The blank import `_ "fmt"` imports the package solely for its side effects (running `init()`). The package's exported names are not accessible.

---

### 42. Named Return with Early Non-Naked Return
**Q: What is printed?**
```go
package main
import "fmt"

func compute(x int) (result int) {
    if x < 0 {
        return -1 // explicit return, ignores named result
    }
    result = x * x
    return // naked return
}

func main() {
    fmt.Println(compute(-5))
    fmt.Println(compute(4))
}
```
**A:**
```
-1
16
```

---

### 43. Closure Over Loop — Range Version
**Q: What is the output (pre Go 1.22)?**
```go
package main
import "fmt"

func main() {
    fns := make([]func(), 0)
    for _, v := range []string{"a", "b", "c"} {
        fns = append(fns, func() { fmt.Print(v) })
    }
    for _, f := range fns {
        f()
    }
}
```
**A:** `ccc`. All closures share the single loop variable `v`, which is `c` after the loop ends.  
**Fix (pre 1.22):** `v := v` inside the loop body.

---

### 44. Panic with Non-String Value
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer func() {
        r := recover()
        fmt.Printf("%T: %v\n", r, r)
    }()
    panic(42)
}
```
**A:** `int: 42`. `panic` can take any value, not just strings or errors.

---

### 45. init Cannot Be Called Explicitly
**Q: Does this compile?**
```go
package main
import "fmt"

func init() { fmt.Println("init") }

func main() {
    init()
}
```
**A:** **Compile Error.** `init` functions cannot be called explicitly; they are invoked automatically by the runtime.

---

## Section 4: Pointers Deep Dives (Q46–Q56)

### 46. Struct Field via Pointer Without Explicit Dereference
**Q: What is the output?**
```go
package main
import "fmt"

type Rect struct{ W, H int }

func area(r *Rect) int { return r.W * r.H }

func main() {
    r := Rect{3, 4}
    fmt.Println(area(&r))
}
```
**A:** `12`. Go auto-dereferences pointer receivers/arguments when accessing fields.

---

### 47. Pointer to Interface Is Almost Always Wrong
**Q: Does this produce expected behavior?**
```go
package main
import "fmt"

type Stringer interface{ String() string }

type T struct{ Name string }
func (t T) String() string { return t.Name }

func printIt(s *Stringer) {
    fmt.Println((*s).String())
}

func main() {
    var s Stringer = T{"Hello"}
    printIt(&s)
}
```
**A:** Prints `Hello` but passing `*Stringer` is almost always a design mistake. Interfaces should be passed by value; they already contain an internal pointer.

---

### 48. Nil Pointer in Method Call (on Pointer Receiver)
**Q: What is the output?**
```go
package main
import "fmt"

type Node struct{ val int }

func (n *Node) GetVal() int {
    if n == nil {
        return -1
    }
    return n.val
}

func main() {
    var n *Node
    fmt.Println(n.GetVal())
}
```
**A:** `-1`. Calling a method on a nil pointer is **valid** in Go as long as the method explicitly handles the nil case. The method receives `nil` as the receiver.

---

### 49. Unsafe Pointer Arithmetic (Concept)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unsafe"
)

func main() {
    arr := [3]int{10, 20, 30}
    p := unsafe.Pointer(&arr[0])
    p2 := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
    second := *(*int)(p2)
    fmt.Println(second)
}
```
**A:** `20`. Using `unsafe.Pointer` and `uintptr` to manually advance through array memory. This is fragile and unsafe — only the combination of uintptr+Pointer in a single expression is safe from GC movement.

---

### 50. Pointer Receiver on Value Type Variable
**Q: Does this compile?**
```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

func main() {
    c := Counter{}
    c.Inc()        // addressable — Go auto-takes &c
    fmt.Println(c.n)
}
```
**A:** **Yes, compiles and prints** `1`. When a variable is addressable, Go automatically takes its address for a pointer-receiver method call.

---

### 51. Map Is Already a Pointer Under the Hood
**Q: What is the output?**
```go
package main
import "fmt"

func addKey(m map[string]int) {
    m["x"] = 99
}

func main() {
    m := map[string]int{}
    addKey(m)
    fmt.Println(m["x"])
}
```
**A:** `99`. Maps are reference types (internally a pointer to a hash table). Passing a map to a function shares the same underlying data.

---

### 52. Function Modifying Slice Header vs Contents
**Q: What is the output?**
```go
package main
import "fmt"

func appendElem(s []int) {
    s = append(s, 99) // modifies local header only
}

func modifyElem(s []int) {
    s[0] = 99 // modifies shared underlying array
}

func main() {
    a := []int{1, 2, 3}
    appendElem(a)
    fmt.Println(a)

    modifyElem(a)
    fmt.Println(a)
}
```
**A:**
```
[1 2 3]
[99 2 3]
```
`append` may allocate a new array; the caller's slice header is not updated. Direct element modification affects shared memory.

---

### 53. Pointer to Slice Can Update Header
**Q: What is the output?**
```go
package main
import "fmt"

func grow(s *[]int) {
    *s = append(*s, 42)
}

func main() {
    s := []int{1, 2}
    grow(&s)
    fmt.Println(s)
}
```
**A:** `[1 2 42]`. Passing `*[]int` lets the function update the caller's slice header (len, cap, ptr).

---

### 54. Stack vs Heap — Escape Analysis
**Q: Which variable escapes to the heap?**
```go
package main
import "fmt"

func stackVar() int {
    x := 42
    return x // x stays on stack; value is copied
}

func heapVar() *int {
    x := 42
    return &x // x escapes to heap; its address is returned
}

func main() {
    fmt.Println(stackVar())
    fmt.Println(*heapVar())
}
```
**A:** `42 42`. `x` in `heapVar` escapes to the heap because its address is returned. This is safe; Go's GC manages it.

---

### 55. Struct Copy vs Pointer Copy
**Q: What is the output?**
```go
package main
import "fmt"

type Config struct{ Debug bool }

func disable(c Config) {
    c.Debug = false
}

func disablePtr(c *Config) {
    c.Debug = false
}

func main() {
    cfg := Config{Debug: true}
    disable(cfg)
    fmt.Println(cfg.Debug)
    disablePtr(&cfg)
    fmt.Println(cfg.Debug)
}
```
**A:**
```
true
false
```

---

### 56. Nil Map Read Is Safe, Write Is Not
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["missing"]) // read from nil map: returns zero value
    m["key"] = 1              // write to nil map: PANIC
}
```
**A:** `0` then **panic: assignment to entry in nil map**. Reading from a nil map is safe; writing panics.

---

## Section 5: Strings & Runes Deep Dives (Q57–Q67)

### 57. Iterating Bytes vs Runes
**Q: What is the difference in output?**
```go
package main
import "fmt"

func main() {
    s := "Go😀"
    fmt.Println("byte len:", len(s))
    fmt.Println("rune len:", len([]rune(s)))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%d:%x ", i, s[i])
    }
}
```
**A:** `byte len: 6`, `rune len: 3`. The emoji 😀 is 4 bytes. Byte iteration: `0:47 1:6f 2:f0 3:9f 4:98 5:80`.

---

### 58. strings.Builder vs bytes.Buffer
**Q: What is printed and what is the key difference in design?**
```go
package main
import (
    "bytes"
    "fmt"
    "strings"
)

func main() {
    var sb strings.Builder
    sb.WriteString("Go")
    sb.WriteByte('!')
    fmt.Println(sb.String())

    var bb bytes.Buffer
    bb.WriteString("Go")
    bb.WriteByte('!')
    fmt.Println(bb.String())
}
```
**A:** Both print `Go!`. Key difference: `strings.Builder` is write-only and optimized for building strings (no read methods). `bytes.Buffer` supports both reading and writing.

---

### 59. Multi-line Raw String Literal
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := `line1
line2
line3`
    fmt.Println(len(s))
}
```
**A:** `17`. Raw string literals (backtick) include literal newlines (`\n`) as-is. `line1\nline2\nline3` = 5+1+5+1+5 = 17 bytes.

---

### 60. strings.Split Behavior
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    parts := strings.Split("a,,b", ",")
    fmt.Println(len(parts), parts)
}
```
**A:** `3 [a  b]`. `Split` keeps the empty string between consecutive delimiters.

---

### 61. String Formatting %v vs %s vs %q
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "Go\nLang"
    fmt.Printf("%v\n", s)
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)
}
```
**A:**
```
Go
Lang
Go
Lang
"Go\nLang"
```
`%q` quotes the string and escapes special characters.

---

### 62. Converting int to string Pitfall
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n := 65
    fmt.Println(string(n))           // rune conversion
    fmt.Println(strconv.Itoa(n))     // integer to string
}
```
**A:**
```
A
65
```
`string(n)` treats `n` as a Unicode code point. Use `strconv.Itoa` for numeric string conversion.

---

### 63. strings.TrimSpace
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "  hello world  "
    fmt.Printf("[%s]\n", strings.TrimSpace(s))
}
```
**A:** `[hello world]`

---

### 64. Byte Mutation via []byte
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    fmt.Println(string(b))
    fmt.Println(s) // original unchanged
}
```
**A:**
```
Hello
hello
```

---

### 65. strings.Replace vs strings.ReplaceAll
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "aabbaa"
    fmt.Println(strings.Replace(s, "a", "x", 2))    // replace first 2
    fmt.Println(strings.ReplaceAll(s, "a", "x"))     // replace all
}
```
**A:**
```
xxbbaa
xxbbxx
```

---

### 66. strings.Fields vs strings.Split
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "  foo   bar  baz  "
    fmt.Println(strings.Fields(s))
    fmt.Println(strings.Split(s, " "))
}
```
**A:**
```
[foo bar baz]
[ foo   bar  baz  ]
```
`Fields` splits by whitespace and ignores leading/trailing/multiple spaces. `Split` is literal.

---

### 67. Rune Arithmetic
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    r := 'A'
    for i := 0; i < 5; i++ {
        fmt.Printf("%c", r+rune(i))
    }
}
```
**A:** `ABCDE`. Rune arithmetic works on Unicode code points.

---

## Section 6: Slices & Maps Deep Dives (Q68–Q83)

### 68. Three-Index Slice (Full Slice Expression)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3:4] // len=2, cap=3
    fmt.Println(len(b), cap(b))
}
```
**A:** `2 3`. The three-index form `a[low:high:max]` sets `cap = max - low`.

---

### 69. append Does Not Modify Original If Capacity Exceeded
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := make([]int, 3, 3)
    a[0], a[1], a[2] = 1, 2, 3
    b := append(a, 4) // triggers reallocation
    b[0] = 99
    fmt.Println(a[0]) // original unaffected
    fmt.Println(b[0])
}
```
**A:**
```
1
99
```
When `append` exceeds capacity, it creates a new backing array. `a` and `b` no longer share memory.

---

### 70. Slice Tricks: Insert at Index
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{1, 2, 4, 5}
    i := 2
    s = append(s[:i+1], s[i:]...)
    s[i] = 3
    fmt.Println(s)
}
```
**A:** `[1 2 3 4 5]`

---

### 71. Map with Struct Values (Non-Addressable)
**Q: Does this compile?**
```go
package main

type Point struct{ X, Y int }

func main() {
    m := map[string]Point{"a": {1, 2}}
    m["a"].X = 10 // cannot assign to struct field in a map
}
```
**A:** **Compile Error.** Struct values stored in maps are **not addressable**. You must retrieve the whole struct, modify, and put it back:
```go
p := m["a"]; p.X = 10; m["a"] = p
```

---

### 72. Slice of Pointers vs Slice of Values
**Q: What is the output?**
```go
package main
import "fmt"

type Item struct{ V int }

func main() {
    items := []Item{{1}, {2}, {3}}
    for i := range items {
        items[i].V *= 10
    }
    for _, it := range items {
        fmt.Print(it.V, " ")
    }
}
```
**A:** `10 20 30 `. Using `items[i]` gives a pointer to the element; modifying via range value (`it`) would fail since `it` is a copy.

---

### 73. Map Concurrent Access
**Q: What is the bug?**
```go
package main
import "sync"

func main() {
    m := map[int]int{}
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            m[n] = n // DATA RACE
        }(i)
    }
    wg.Wait()
}
```
**A:** **Data race.** Concurrent writes to a map without synchronization cause undefined behavior and will be detected by `-race`. Fix: use `sync.Mutex` or `sync.Map`.

---

### 74. Delete During Iteration
**Q: Is this safe?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    for k := range m {
        if k == 2 {
            delete(m, k)
        }
    }
    fmt.Println(len(m))
}
```
**A:** **Yes, safe.** It is safe to delete map keys during a `range` loop. Output: `2`.

---

### 75. make vs literal for Map
**Q: What difference exists between these two maps?**
```go
package main
import "fmt"

func main() {
    m1 := make(map[string]int)
    m2 := map[string]int{}
    fmt.Println(m1 == nil, m2 == nil)
}
```
**A:** `false false`. Both are initialized (non-nil) maps. `make` with a capacity hint is preferred for performance when size is known in advance.

---

### 76. Slice Contains (Linear Search)
**Q: What is the output?**
```go
package main
import "fmt"

func contains(s []string, target string) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(contains([]string{"a", "b", "c"}, "b"))
    fmt.Println(contains([]string{"a", "b", "c"}, "z"))
}
```
**A:**
```
true
false
```

---

### 77. Copy Between Different Length Slices
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 3)
    n := copy(dst, src)
    fmt.Println(n, dst)
}
```
**A:** `3 [1 2 3]`. `copy` copies `min(len(dst), len(src))` elements.

---

### 78. Map of Slices: Append Pattern
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string][]int{}
    for i := 0; i < 3; i++ {
        m["nums"] = append(m["nums"], i)
    }
    fmt.Println(m["nums"])
}
```
**A:** `[0 1 2]`. The zero value for a missing map key is `nil` slice, and appending to `nil` is valid.

---

### 79. Slice of nil vs Empty Slice in JSON
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    var a []int
    b := []int{}

    ja, _ := json.Marshal(a)
    jb, _ := json.Marshal(b)
    fmt.Println(string(ja))
    fmt.Println(string(jb))
}
```
**A:**
```
null
[]
```
A nil slice marshals to JSON `null`; an empty slice marshals to `[]`.

---

### 80. Nested Map Initialization
**Q: Does this panic?**
```go
package main
import "fmt"

func main() {
    m := map[string]map[string]int{}
    m["a"]["x"] = 1
    fmt.Println(m)
}
```
**A:** **Panic.** `m["a"]` returns `nil` (the zero value for `map[string]int`). You must initialize the inner map: `m["a"] = map[string]int{}; m["a"]["x"] = 1`.

---

### 81. Slice as Set (Using map[T]struct{})
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    set := map[string]struct{}{}
    words := []string{"go", "is", "fun", "go"}
    for _, w := range words {
        set[w] = struct{}{}
    }
    fmt.Println(len(set))
}
```
**A:** `3`. Duplicate `"go"` is overwritten. The empty struct uses zero memory as the value type.

---

### 82. Modifying Slice Backing Array via Two Slices
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    base := make([]int, 5)
    a := base[0:3]
    b := base[2:5]
    a[2] = 99
    fmt.Println(b[0])
}
```
**A:** `99`. `a[2]` and `b[0]` both point to `base[2]` — the same underlying array element.

---

### 83. Map Key Types Must Be Comparable
**Q: Does this compile?**
```go
package main

func main() {
    m := map[[]int]string{}
    _ = m
}
```
**A:** **Compile Error.** Slices are not comparable and therefore cannot be used as map keys. Valid key types include: bool, numeric, string, pointer, channel, interface, array, struct (if all fields are comparable).

---

## Section 7: Structs & Interfaces Deep Dives (Q84–Q95)

### 84. Struct Promoted Method Conflict Resolution
**Q: What is the output?**
```go
package main
import "fmt"

type Base struct{}
func (Base) Describe() string { return "Base" }

type Child struct{ Base }
func (Child) Describe() string { return "Child" }

func main() {
    c := Child{}
    fmt.Println(c.Describe())
    fmt.Println(c.Base.Describe())
}
```
**A:**
```
Child
Base
```
`Child.Describe()` overrides the promoted method. You can still access `Base.Describe` explicitly.

---

### 85. Interface with Multiple Types Implementing
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

type Shape interface{ Area() float64 }

type Circle struct{ R float64 }
type Rect struct{ W, H float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
func (r Rect) Area() float64   { return r.W * r.H }

func printArea(s Shape) { fmt.Printf("%.2f\n", s.Area()) }

func main() {
    printArea(Circle{R: 1})
    printArea(Rect{W: 3, H: 4})
}
```
**A:**
```
3.14
12.00
```

---

### 86. Embedding Interface in Struct
**Q: Does this compile and does it satisfy the interface?**
```go
package main
import "fmt"

type Speaker interface{ Speak() string }

type Robot struct{ Speaker }

func main() {
    r := Robot{}
    fmt.Println(r.Speak())
}
```
**A:** **Compiles but panics at runtime.** `Robot.Speaker` is `nil` (zero value). Calling `r.Speak()` is a nil pointer dereference on the embedded interface.

---

### 87. Empty Interface as Function Parameter
**Q: Can an `interface{}` argument modify the original value?**
```go
package main
import "fmt"

func tryChange(v interface{}) {
    v = 999
}

func main() {
    x := 42
    tryChange(x)
    fmt.Println(x)
}
```
**A:** `42`. `interface{}` wraps the value; rebinding `v` inside the function does not affect the original.

---

### 88. Stringer Interface Auto-Used by fmt
**Q: What is the output?**
```go
package main
import "fmt"

type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) String() string {
    return []string{"Red", "Green", "Blue"}[c]
}

func main() {
    fmt.Println(Red)
    fmt.Println(Green)
    fmt.Println(Blue)
}
```
**A:**
```
Red
Green
Blue
```
`fmt` package automatically calls `String()` if a type implements the `fmt.Stringer` interface.

---

### 89. Struct Tags Are Metadata Only
**Q: Does modifying a tag affect runtime behavior of non-json usage?**
```go
package main
import (
    "fmt"
    "reflect"
)

type T struct {
    Name string `mytag:"custom"`
}

func main() {
    t := T{Name: "Go"}
    field, _ := reflect.TypeOf(t).FieldByName("Name")
    fmt.Println(field.Tag.Get("mytag"))
    fmt.Println(t.Name)
}
```
**A:**
```
custom
Go
```
Struct tags are only accessible via reflection; they don't affect normal field access.

---

### 90. Nil Receiver of Concrete Type
**Q: What is the output?**
```go
package main
import "fmt"

type T struct{ val int }

func (t *T) Print() {
    if t == nil {
        fmt.Println("nil receiver")
        return
    }
    fmt.Println(t.val)
}

func main() {
    var t *T
    t.Print()
}
```
**A:** `nil receiver`. Calling a method on a nil pointer is valid if the method handles it.

---

### 91. Interface Comparison — Different Dynamic Types
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var a interface{} = 1
    var b interface{} = 1
    var c interface{} = "1"
    fmt.Println(a == b)
    fmt.Println(a == c)
}
```
**A:**
```
true
false
```
Interface values are equal only if both their dynamic type AND dynamic value are equal.

---

### 92. Type Switch vs Type Assertion
**Q: What is the output?**
```go
package main
import "fmt"

func check(i interface{}) {
    switch v := i.(type) {
    case int, float64:
        fmt.Printf("numeric: %v\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    }
}

func main() {
    check(3.14)
    check("hello")
    check(true)
}
```
**A:**
```
numeric: 3.14
string: hello
```
`true` matches no case and falls through silently (no default).

---

### 93. Embedding Promotes Fields, Not Methods When Overridden
**Q: What is the output?**
```go
package main
import "fmt"

type Animal struct{ Name string }
func (a Animal) Sound() string { return "..." }

type Cat struct {
    Animal
    Sound string // field named Sound shadows promoted method
}

func main() {
    c := Cat{Animal{"Whiskers"}, "Meow"}
    fmt.Println(c.Sound)       // field
    fmt.Println(c.Animal.Sound()) // method
}
```
**A:**
```
Meow
...
```
The `Sound` field in `Cat` shadows the promoted `Sound()` method from `Animal`.

---

### 94. Interface Slice Type Assertion
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i interface{} = []int{1, 2, 3}
    s, ok := i.([]int)
    fmt.Println(s, ok)
}
```
**A:** `[1 2 3] true`

---

### 95. Structural Typing — Duck Typing
**Q: Does this compile?**
```go
package main
import "fmt"

type Walker interface{ Walk() }

type Person struct{}
func (p Person) Walk() { fmt.Println("walking") }

type Robot struct{}
func (r Robot) Walk() { fmt.Println("rolling") }

func makeWalk(w Walker) { w.Walk() }

func main() {
    makeWalk(Person{})
    makeWalk(Robot{})
}
```
**A:** **Yes.** Go uses structural typing — any type with the `Walk()` method satisfies `Walker` without explicitly declaring it.
```
walking
rolling
```

---

## Section 8: Error Handling Deep Dives (Q96–Q100)

### 96. errors.New vs fmt.Errorf
**Q: What is the difference?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrA = errors.New("error A")

func main() {
    err1 := errors.New("error A") // new sentinel each time
    err2 := fmt.Errorf("context: %w", ErrA) // wraps ErrA

    fmt.Println(err1 == ErrA)           // false: different pointers
    fmt.Println(errors.Is(err2, ErrA))  // true: unwraps to find ErrA
}
```
**A:**
```
false
true
```

---

### 97. Wrapping Errors Multiple Levels
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrRoot = errors.New("root cause")

func level2() error { return fmt.Errorf("L2: %w", ErrRoot) }
func level1() error { return fmt.Errorf("L1: %w", level2()) }

func main() {
    err := level1()
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrRoot))
}
```
**A:**
```
L1: L2: root cause
true
```
`errors.Is` unwraps through the chain to find `ErrRoot`.

---

### 98. Panic vs Error Return Convention
**Q: What is the idiomatic Go approach?**
```go
package main
import (
    "errors"
    "fmt"
)

// Idiomatic: return error for expected failures
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}
```
**A:** `Error: division by zero`. Idiomatic Go: return errors for expected/recoverable failures; use `panic` only for truly unrecoverable programming errors.

---

### 99. Error Type Switch
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("/nonexistent/file")
    if err != nil {
        switch e := err.(type) {
        case *os.PathError:
            fmt.Println("PathError:", e.Op, e.Path)
        default:
            fmt.Println("other:", e)
        }
    }
}
```
**A:** `PathError: open /nonexistent/file`

---

### 100. Sentinel Error Pattern
**Q: What is the output and why is `var` (not `const`) used?**
```go
package main
import (
    "errors"
    "fmt"
)

var (
    ErrNotFound   = errors.New("not found")
    ErrPermission = errors.New("permission denied")
)

func fetch(id int) error {
    if id < 0 {
        return ErrPermission
    }
    if id == 0 {
        return ErrNotFound
    }
    return nil
}

func main() {
    for _, id := range []int{0, -1, 1} {
        err := fetch(id)
        fmt.Printf("id=%d: errors.Is(ErrNotFound)=%v errors.Is(ErrPermission)=%v\n",
            id,
            errors.Is(err, ErrNotFound),
            errors.Is(err, ErrPermission))
    }
}
```
**A:**
```
id=0: errors.Is(ErrNotFound)=true errors.Is(ErrPermission)=false
id=-1: errors.Is(ErrNotFound)=false errors.Is(ErrPermission)=true
id=1: errors.Is(ErrNotFound)=false errors.Is(ErrPermission)=false
```
`var` is used because `errors.New` returns a pointer and **identity** is what uniquely identifies sentinel errors. Constants cannot hold pointer values.

---

*End of 100 In-Depth Go Basics & Fundamentals Code Snippet Questions*
