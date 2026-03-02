# 100 Pure Code Snippet Interview Questions: Go Basics & Fundamentals

*Each question is a "predict the output / spot the bug / does it compile?" style question.*
*Topics: Variables, Types, Control Flow, Functions, Closures, Pointers, Strings, Arrays, Slices, Maps, Structs, Interfaces, Error Handling, Defer, Goroutines (basic), Type System.*

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** Variables, Constants & Types (Q1–Q15)
- [ ] **Section 2:** Control Flow (Q16–Q28)
- [ ] **Section 3:** Functions, Closures & Defer (Q29–Q44)
- [ ] **Section 4:** Pointers (Q45–Q52)
- [ ] **Section 5:** Strings & Runes (Q53–Q62)
- [ ] **Section 6:** Arrays, Slices & Maps (Q63–Q78)
- [ ] **Section 7:** Structs & Interfaces (Q79–Q91)
- [ ] **Section 8:** Error Handling (Q92–Q96)
- [ ] **Section 9:** Goroutines Basics & Misc (Q97–Q100)

> 🔖 **Last read:** <!-- e.g. Q15 · Section 1 done -->

---

## Section 1: Variables, Constants & Types (Q1–Q15)

### 1. Short Variable Declaration Outside Function
**Q: Does this compile?**
```go
package main

:= 10
fmt.Println(x)
```
**A:** **Compile Error.** Short variable declaration (`:=`) is not allowed at the package level. Use `var x = 10` or `var x int = 10`.

---

### 2. Multiple Return Values
**Q: What is the output?**
```go
package main
import "fmt"

func minMax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

func main() {
    lo, hi := minMax(7, 3)
    fmt.Println(lo, hi)
}
```
**A:** `3 7`

---

### 3. Zero Values
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```
**A:** `0 0 false ""`

---

### 4. Untyped Constants
**Q: Does this compile?**
```go
package main
import "fmt"

const x = 1e300

func main() {
    var f float32 = x
    fmt.Println(f)
}
```
**A:** **Compile Error.** Constant `1e300` overflows `float32`. Untyped constants have arbitrary precision but must fit the target type at assignment.

---

### 5. Integer Overflow (Wrapping)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    var x int8 = math.MaxInt8
    x++
    fmt.Println(x)
}
```
**A:** `-128`. Integer overflow wraps around silently in Go at runtime.

---

### 6. Multiple Assignment
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x, y := 1, 2
    x, y = y, x
    fmt.Println(x, y)
}
```
**A:** `2 1`. Go evaluates all right-hand side expressions before assigning.

---

### 7. Type Alias vs New Type
**Q: Does this compile?**
```go
package main
import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 100
    var f Fahrenheit = c
    fmt.Println(f)
}
```
**A:** **Compile Error.** `Celsius` and `Fahrenheit` are distinct named types. You need an explicit conversion: `var f Fahrenheit = Fahrenheit(c)`.

---

### 8. Iota with Bit Shift
**Q: What are the values of KB, MB, GB?**
```go
package main
import "fmt"

const (
    _  = iota
    KB = 1 << (10 * iota)
    MB
    GB
)

func main() {
    fmt.Println(KB, MB, GB)
}
```
**A:** `1024 1048576 1073741824`

---

### 9. Blank Identifier
**Q: Does this compile?**
```go
package main
import "fmt"

func twoVals() (int, string) { return 1, "a" }

func main() {
    _, s := twoVals()
    fmt.Println(s)
}
```
**A:** **Yes, compiles and prints** `a`. The blank identifier `_` discards the first return value.

---

### 10. var vs :=
**Q: What is the output?**
```go
package main
import "fmt"

var x = 10

func main() {
    x := 20
    fmt.Println(x)
}
```
**A:** `20`. The `:=` inside `main` creates a new local variable `x` that shadows the package-level `x`.

---

### 11. Typed nil
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var p *int
    fmt.Println(p == nil)
    fmt.Println(p)
}
```
**A:**
```
true
<nil>
```

---

### 12. Constants Cannot Be Addressed
**Q: Does this compile?**
```go
package main

const x = 42

func main() {
    p := &x
    _ = p
}
```
**A:** **Compile Error.** You cannot take the address of a constant in Go.

---

### 13. Declaring Multiple Variables
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var (
        a = 1
        b = 2
        c = a + b
    )
    fmt.Println(c)
}
```
**A:** `3`

---

### 14. Iota Reset Per const Block
**Q: What are the values of X and Y?**
```go
package main
import "fmt"

const A = iota
const (
    X = iota
    Y
)

func main() {
    fmt.Println(A, X, Y)
}
```
**A:** `0 0 1`. `iota` resets to `0` at the start of each new `const` block.

---

### 15. Unused Variables
**Q: Does this compile?**
```go
package main

func main() {
    x := 5
}
```
**A:** **Compile Error.** `x declared and not used`. Go enforces that every declared local variable must be used.

---

## Section 2: Control Flow (Q16–Q28)

### 16. For as While Loop
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for i < 3 {
        fmt.Print(i)
        i++
    }
}
```
**A:** `012`

---

### 17. Infinite Loop with Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for {
        if i == 3 {
            break
        }
        fmt.Print(i)
        i++
    }
}
```
**A:** `012`

---

### 18. Switch Without Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 15
    switch {
    case x < 10:
        fmt.Println("small")
    case x < 20:
        fmt.Println("medium")
    default:
        fmt.Println("large")
    }
}
```
**A:** `medium`

---

### 19. Switch Type
**Q: What is the output?**
```go
package main
import "fmt"

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    default:
        fmt.Printf("other: %T\n", v)
    }
}

func main() {
    describe(42)
    describe("hi")
    describe(3.14)
}
```
**A:**
```
int: 42
string: hi
other: float64
```

---

### 20. Continue in For Loop
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Print(i)
    }
}
```
**A:** `13`

---

### 21. Labeled Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                break outer
            }
            fmt.Print(i, j, " ")
        }
    }
}
```
**A:** `00 `. The labeled `break outer` exits the outer loop immediately when `j == 1` on the first iteration.

---

### 22. Switch Fallthrough Does Not Check Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("one")
        fallthrough
    case 2:
        fmt.Println("two")
        fallthrough
    case 3:
        fmt.Println("three")
    }
}
```
**A:**
```
one
two
three
```
`fallthrough` bypasses the case condition check entirely.

---

### 23. goto Statement
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
loop:
    if i < 3 {
        fmt.Print(i)
        i++
        goto loop
    }
}
```
**A:** **Compiles and prints** `012`. `goto` is valid in Go but discouraged.

---

### 24. Range Over String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i, c := range "Go" {
        fmt.Println(i, c)
    }
}
```
**A:**
```
0 71
1 111
```
`range` over a string yields byte index and Unicode code point (`rune`). `G`=71, `o`=111.

---

### 25. Index-Only Range
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{10, 20, 30}
    for i := range s {
        fmt.Print(i)
    }
}
```
**A:** `012`. When you use a single variable in `range`, you get only the index.

---

### 26. Switch with Multiple Values Per Case
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    day := "Saturday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
```
**A:** `Weekend`

---

### 27. Defer in Loop Execution Order
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Print(i)
    }
}
```
**A:** `210`. Deferred calls execute in LIFO (last-in, first-out) order after the surrounding function returns.

---

### 28. If with Initialization Statement
**Q: What is the output?**
```go
package main
import "fmt"

func getValue() int { return 42 }

func main() {
    if v := getValue(); v > 10 {
        fmt.Println("big:", v)
    }
    // fmt.Println(v) // v is not in scope here
}
```
**A:** `big: 42`. The `v` initialized in the `if` statement is scoped only to the `if-else` block.

---

## Section 3: Functions, Closures & Defer (Q29–Q44)

### 29. Function as First-Class Value
**Q: What is the output?**
```go
package main
import "fmt"

func apply(f func(int) int, x int) int {
    return f(x)
}

func double(n int) int { return n * 2 }

func main() {
    fmt.Println(apply(double, 5))
}
```
**A:** `10`

---

### 30. Closure Captures Variable by Reference
**Q: What is the output?**
```go
package main
import "fmt"

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := makeCounter()
    fmt.Println(c())
    fmt.Println(c())
    fmt.Println(c())
}
```
**A:** `1 2 3`. The closure captures the variable `count` by reference — each call increments the same `count`.

---

### 31. Multiple Defers Order
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
    fmt.Println("start")
}
```
**A:**
```
start
third
second
first
```

---

### 32. Variadic Function
**Q: What is the output?**
```go
package main
import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))
    s := []int{4, 5, 6}
    fmt.Println(sum(s...))
}
```
**A:**
```
6
15
```

---

### 33. Named Return Values
**Q: What is the output?**
```go
package main
import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```
**A:** `7 10`

---

### 34. Defer and Panic Recovery
**Q: What is the output?**
```go
package main
import "fmt"

func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b
    return
}

func main() {
    r, err := safeDiv(10, 0)
    fmt.Println(r, err)
}
```
**A:** `0 recovered: runtime error: integer divide by zero`

---

### 35. Closure in Goroutine (Classic Bug)
**Q: What is the typical output (pre Go 1.22)?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Print(i)
        }()
    }
    wg.Wait()
}
```
**A:** Usually `333`. The goroutines capture the same variable `i`; by the time they run, the loop has finished and `i` is `3`.
**Fix:** `go func(n int) { ... }(i)`

---

### 36. Function Type as Struct Field
**Q: What is the output?**
```go
package main
import "fmt"

type Greeter struct {
    Greet func(name string) string
}

func main() {
    g := Greeter{
        Greet: func(name string) string {
            return "Hello, " + name
        },
    }
    fmt.Println(g.Greet("Go"))
}
```
**A:** `Hello, Go`

---

### 37. Deferred Function Sees Updated Return Value
**Q: What is returned?**
```go
package main
import "fmt"

func f() (n int) {
    defer func() {
        n++
    }()
    return 1
}

func main() {
    fmt.Println(f())
}
```
**A:** `2`. `return 1` sets named return `n = 1`, then the deferred closure increments it to `2`.

---

### 38. Anonymous Function Immediately Invoked
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    result := func(x, y int) int {
        return x + y
    }(3, 4)
    fmt.Println(result)
}
```
**A:** `7`

---

### 39. Defer with Panic (No Recover)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer fmt.Println("deferred")
    panic("crash!")
}
```
**A:**
```
deferred
panic: crash!
...
```
Deferred functions run even during a panic before the program prints the panic message and exits.

---

### 40. Recursive Closure
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    var fib func(n int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
```
**A:** **Yes, compiles and prints** `13`. You must first declare `var fib func(n int) int` so that the closure can refer to itself by name.

---

### 41. Panic in Deferred Function
**Q: What happens?**
```go
package main
import "fmt"

func main() {
    defer func() {
        fmt.Println("defer 1")
        panic("panic in defer")
    }()
    defer fmt.Println("defer 2")
    fmt.Println("main")
}
```
**A:**
```
main
defer 2
defer 1
panic: panic in defer
```
Deferred functions still run in LIFO order. A panic inside a deferred function propagates and can be recovered by another deferred function.

---

### 42. Function With No Parameters or Returns
**Q: Does this compile?**
```go
package main
import "fmt"

func sayHi() {
    fmt.Println("Hi!")
}

func main() {
    f := sayHi
    f()
}
```
**A:** **Yes, compiles and prints** `Hi!`. Functions are first-class values and can be assigned to variables.

---

### 43. Unused Function Parameter
**Q: Does this compile?**
```go
package main
import "fmt"

func greet(name string) {
    fmt.Println("Hello!")
}

func main() {
    greet("Alice")
}
```
**A:** **Yes.** Unlike local variables, unused function parameters do **not** cause a compile error.

---

### 44. defer + os.Exit
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("deferred")
    os.Exit(0)
}
```
**A:** Nothing is printed. `os.Exit` immediately terminates the process; deferred functions are **not** called.

---

## Section 4: Pointers (Q45–Q52)

### 45. Pointer Basics
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 42
    p := &x
    *p = 100
    fmt.Println(x)
}
```
**A:** `100`

---

### 46. Pointer to Pointer
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 10
    p := &x
    pp := &p
    **pp = 20
    fmt.Println(x)
}
```
**A:** `20`

---

### 47. new() Builtin
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    p := new(int)
    fmt.Println(*p)
    *p = 7
    fmt.Println(*p)
}
```
**A:**
```
0
7
```
`new(T)` allocates a zeroed T and returns a pointer to it.

---

### 48. Nil Pointer Dereference
**Q: What is the output or error?**
```go
package main
import "fmt"

type Node struct{ val int }

func main() {
    var n *Node
    fmt.Println(n.val)
}
```
**A:** **Panic.** `runtime error: invalid memory address or nil pointer dereference`. You cannot dereference a nil pointer.

---

### 49. Passing Pointer to Function
**Q: What is the output?**
```go
package main
import "fmt"

func increment(x *int) {
    *x++
}

func main() {
    n := 5
    increment(&n)
    fmt.Println(n)
}
```
**A:** `6`

---

### 50. Pointer Comparison
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a, b := 1, 1
    p1, p2 := &a, &b
    p3 := &a
    fmt.Println(p1 == p2)
    fmt.Println(p1 == p3)
}
```
**A:**
```
false
true
```
Pointers are compared by address, not by the value they point to.

---

### 51. Returning Local Variable Pointer
**Q: Is this safe in Go?**
```go
package main
import "fmt"

func newInt() *int {
    x := 42
    return &x
}

func main() {
    p := newInt()
    fmt.Println(*p)
}
```
**A:** **Yes, it is safe.** Go's compiler detects that `x` escapes to the heap and allocates it there. This is called escape analysis. Output: `42`.

---

### 52. Modifying Struct Through Pointer
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func moveRight(p *Point) {
    p.X += 10
}

func main() {
    pt := Point{1, 2}
    moveRight(&pt)
    fmt.Println(pt)
}
```
**A:** `{11 2}`

---

## Section 5: Strings & Runes (Q53–Q62)

### 53. len() on a String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "café"
    fmt.Println(len(s))
}
```
**A:** `5`. `len()` counts bytes, not characters. The `é` character is 2 bytes in UTF-8.

---

### 54. String Concatenation in Loop (Performance)
**Q: What is the inefficiency?**
```go
package main
import "fmt"

func main() {
    s := ""
    for i := 0; i < 5; i++ {
        s += fmt.Sprintf("%d", i)
    }
    fmt.Println(s)
}
```
**A:** Output is `01234`, but each `+=` allocates a new string (strings are immutable). For large loops, use `strings.Builder`.

---

### 55. Byte Slice to String Conversion
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    b := []byte{72, 101, 108, 108, 111}
    s := string(b)
    fmt.Println(s)
}
```
**A:** `Hello`

---

### 56. String to Rune Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "café"
    r := []rune(s)
    fmt.Println(len(r))
}
```
**A:** `4`. Converting to `[]rune` counts Unicode code points, not bytes.

---

### 57. Comparing Strings
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := "hello"
    b := "hello"
    fmt.Println(a == b)
}
```
**A:** `true`. String comparison in Go is value-based (compares contents byte by byte).

---

### 58. String Indexing Returns Byte
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "Go"
    fmt.Printf("%T %v\n", s[0], s[0])
}
```
**A:** `uint8 71`. Indexing a string returns a `byte` (`uint8`), not a `rune`.

---

### 59. String Contains Check
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("seafood", "foo"))
    fmt.Println(strings.Contains("seafood", "bar"))
}
```
**A:**
```
true
false
```

---

### 60. Rune Literal
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    r := 'A'
    fmt.Printf("%T %v\n", r, r)
}
```
**A:** `int32 65`. Rune literals have type `rune`, which is an alias for `int32`.

---

### 61. String Builder
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var sb strings.Builder
    for i := 0; i < 3; i++ {
        sb.WriteString("Go")
    }
    fmt.Println(sb.String())
}
```
**A:** `GoGoGo`

---

### 62. Strings Are Not Byte Slices
**Q: Does this compile?**
```go
package main

func main() {
    s := "hello"
    b := []byte(s)
    s2 := string(b)
    _ = s2
    b[0] = 'H'
    _ = s
}
```
**A:** **Yes.** Mutation of `b` does not affect `s`. `string(b)` copies; strings and byte slices are independent.

---

## Section 6: Arrays, Slices & Maps (Q63–Q78)

### 63. Array Is a Value Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a
    b[0] = 99
    fmt.Println(a[0], b[0])
}
```
**A:** `1 99`. Arrays are value types; `b` is a complete copy of `a`.

---

### 64. Slice Length and Capacity
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Println(len(s), cap(s))
    s = append(s, 1, 2)
    fmt.Println(len(s), cap(s))
}
```
**A:**
```
3 5
5 5
```

---

### 65. Append Grows Capacity
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s []int
    for i := 0; i < 5; i++ {
        s = append(s, i)
        fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
    }
}
```
**A:** Len grows by 1 each time; capacity doubles when exceeded (e.g., `1 1 → 2 2 → 3 4 → 4 4 → 5 8`). Exact values may vary by runtime.

---

### 66. Nil Slice vs Empty Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s1 []int
    s2 := []int{}
    fmt.Println(s1 == nil)
    fmt.Println(s2 == nil)
    fmt.Println(len(s1), len(s2))
}
```
**A:**
```
true
false
0 0
```

---

### 67. Slice of Slice (Shared Backing Array)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3]
    b[0] = 99
    fmt.Println(a)
}
```
**A:** `[1 99 3 4 5]`. Slicing does not copy the underlying array; `b[0]` is the same memory as `a[1]`.

---

### 68. Map Key Existence
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1}
    v, ok := m["a"]
    fmt.Println(v, ok)
    v2, ok2 := m["b"]
    fmt.Println(v2, ok2)
}
```
**A:**
```
1 true
0 false
```

---

### 69. Deleting from Map
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"x": 10, "y": 20}
    delete(m, "x")
    delete(m, "z") // deleting non-existent key: no-op
    fmt.Println(len(m))
}
```
**A:** `1`

---

### 70. Map Cannot Be Compared with ==
**Q: Does this compile?**
```go
package main

func main() {
    m1 := map[string]int{"a": 1}
    m2 := map[string]int{"a": 1}
    _ = m1 == m2
}
```
**A:** **Compile Error.** Maps are not comparable with `==` in Go (only maps can be compared to `nil`). Use `reflect.DeepEqual` instead.

---

### 71. 2D Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println(matrix[1][2])
}
```
**A:** `6`

---

### 72. copy() Function
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    src := []int{1, 2, 3}
    dst := make([]int, 2)
    n := copy(dst, src)
    fmt.Println(n, dst)
}
```
**A:** `2 [1 2]`. `copy` copies `min(len(dst), len(src))` elements.

---

### 73. Range Over Map Returns Copies
**Q: Does modifying the value in range affect the map?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    for k, v := range m {
        v += 10
        _ = k
    }
    fmt.Println(m)
}
```
**A:** `map[a:1 b:2]`. The `v` in range is a copy; modifying it does not affect the original map.

---

### 74. Slice as Map Value
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string][]int{}
    m["evens"] = append(m["evens"], 2, 4, 6)
    fmt.Println(m["evens"])
}
```
**A:** `[2 4 6]`

---

### 75. Iterating Map Order
**Q: What is guaranteed about the output?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
**A:** **Nothing is guaranteed.** Map iteration order is intentionally randomized in Go.

---

### 76. Append to Nil Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s []string
    s = append(s, "hello", "world")
    fmt.Println(s, len(s))
}
```
**A:** `[hello world] 2`

---

### 77. Slice Tricks: Remove Element
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    i := 2 // remove index 2
    s = append(s[:i], s[i+1:]...)
    fmt.Println(s)
}
```
**A:** `[1 2 4 5]`

---

### 78. Array vs Slice Type System
**Q: Does this compile?**
```go
package main
import "fmt"

func printSlice(s []int) { fmt.Println(s) }

func main() {
    arr := [3]int{1, 2, 3}
    printSlice(arr)
}
```
**A:** **Compile Error.** An array `[3]int` cannot be passed where a slice `[]int` is expected. Use `printSlice(arr[:])`.

---

## Section 7: Structs & Interfaces (Q79–Q91)

### 79. Struct Embedding
**Q: What is the output?**
```go
package main
import "fmt"

type Animal struct{ Name string }
func (a Animal) Speak() string { return a.Name + " speaks" }

type Dog struct{ Animal }

func main() {
    d := Dog{Animal{"Rex"}}
    fmt.Println(d.Speak())
    fmt.Println(d.Name)
}
```
**A:**
```
Rex speaks
Rex
```
Promoted methods and fields are accessible directly on the embedding struct.

---

### 80. Interface Satisfaction
**Q: Does this compile?**
```go
package main
import "fmt"

type Stringer interface {
    String() string
}

type Person struct{ Name string }
func (p Person) String() string { return "Person: " + p.Name }

func main() {
    var s Stringer = Person{"Alice"}
    fmt.Println(s.String())
}
```
**A:** **Yes, compiles and prints** `Person: Alice`. Go uses structural (implicit) interface satisfaction.

---

### 81. Interface Nil Check
**Q: What is the output?**
```go
package main
import "fmt"

type MyError struct{}
func (e *MyError) Error() string { return "error" }

func getErr() error {
    var e *MyError
    return e
}

func main() {
    err := getErr()
    fmt.Println(err == nil)
}
```
**A:** `false`. The interface value is non-nil because it holds a type descriptor (`*MyError`), even though the underlying pointer is nil.

---

### 82. Struct Tag
**Q: What does this print?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

func main() {
    u := User{Name: "Bob"}
    b, _ := json.Marshal(u)
    fmt.Println(string(b))
}
```
**A:** `{"name":"Bob"}`. `omitempty` causes the `Age` field (zero value) to be omitted.

---

### 83. Empty Interface
**Q: What is the output?**
```go
package main
import "fmt"

func printAny(v interface{}) {
    fmt.Printf("%T: %v\n", v, v)
}

func main() {
    printAny(42)
    printAny("hello")
    printAny([]int{1, 2})
}
```
**A:**
```
int: 42
string: hello
[]int: [1 2]
```

---

### 84. Struct Anonymous Fields Cannot Duplicate
**Q: Does this compile?**
```go
package main

type A struct{ ID int }
type B struct{ ID int }
type C struct {
    A
    B
}

func main() {
    c := C{}
    _ = c.ID
}
```
**A:** **Compile Error.** `c.ID` is ambiguous — both `A.ID` and `B.ID` are promoted. You must use `c.A.ID` or `c.B.ID`.

---

### 85. Value vs Pointer Method Sets
**Q: Does this compile?**
```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

type Incrementer interface{ Inc() }

func main() {
    c := Counter{}
    var i Incrementer = c // not addressable
    i.Inc()
    fmt.Println(c.n)
}
```
**A:** **Compile Error.** `Counter` (value) does not implement `Incrementer` because `Inc` has a pointer receiver. Use `var i Incrementer = &c`.

---

### 86. Struct Initialization Positional
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y, Z int }

func main() {
    p := Point{1, 2, 3}
    fmt.Println(p.Y)
}
```
**A:** `2`

---

### 87. Interfaces Can Hold nil Values
**Q: What is the output?**
```go
package main
import "fmt"

type Describer interface{ Describe() }

func process(d Describer) {
    if d != nil {
        fmt.Println("not nil interface")
    }
}

func main() {
    process(nil)
}
```
**A:** Nothing is printed. Passing a literal `nil` to an interface parameter gives a nil interface (both type and value are nil), so the check `d != nil` is false.

---

### 88. Type Assertion Safe Form
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i interface{} = "hello"
    s, ok := i.(string)
    n, ok2 := i.(int)
    fmt.Println(s, ok)
    fmt.Println(n, ok2)
}
```
**A:**
```
hello true
0 false
```

---

### 89. Struct Pointer Shorthand
**Q: What is the output?**
```go
package main
import "fmt"

type Node struct{ Val int }

func main() {
    n := &Node{Val: 5}
    fmt.Println(n.Val) // auto-dereferenced
}
```
**A:** `5`. Go auto-dereferences struct pointers when accessing fields.

---

### 90. Interface Wrapping
**Q: What is the output?**
```go
package main
import "fmt"

type Inner struct{}
func (Inner) Hello() string { return "inner" }

type Outer struct{ Inner }
func (Outer) Hello() string { return "outer" }

func main() {
    o := Outer{}
    fmt.Println(o.Hello())
    fmt.Println(o.Inner.Hello())
}
```
**A:**
```
outer
inner
```
The `Outer` method `Hello` overrides the promoted `Inner.Hello`.

---

### 91. Comparing Nil Interfaces
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var a, b interface{}
    fmt.Println(a == b)
    a = 1
    fmt.Println(a == b)
}
```
**A:**
```
true
false
```

---

## Section 8: Error Handling (Q92–Q96)

### 92. Error Wrapping with fmt.Errorf
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
    return fmt.Errorf("findUser %d: %w", id, ErrNotFound)
}

func main() {
    err := findUser(42)
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrNotFound))
}
```
**A:**
```
findUser 42: not found
true
```

---

### 93. errors.As
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

type ValidationError struct{ Field string }
func (e *ValidationError) Error() string { return "invalid: " + e.Field }

func validate() error {
    return fmt.Errorf("wrap: %w", &ValidationError{Field: "email"})
}

func main() {
    err := validate()
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Println("field:", ve.Field)
    }
}
```
**A:** `field: email`

---

### 94. Ignoring Errors
**Q: What is the bug?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi("abc")
    fmt.Println(n * 2)
}
```
**A:** Output is `0`. There's no panic, but the error from `Atoi` is silently ignored. `n` defaults to `0` on failure, leading to silently wrong results.

---

### 95. Custom Error Type
**Q: What is the output?**
```go
package main
import "fmt"

type AppError struct {
    Code    int
    Message string
}

func (e AppError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func riskyOp() error {
    return AppError{Code: 404, Message: "resource not found"}
}

func main() {
    err := riskyOp()
    fmt.Println(err)
}
```
**A:** `[404] resource not found`

---

### 96. Sentinel Error Comparison
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
    "io"
)

func main() {
    err := io.EOF
    fmt.Println(err == io.EOF)
    fmt.Println(errors.Is(err, io.EOF))
}
```
**A:**
```
true
true
```

---

## Section 9: Goroutines Basics & Misc (Q97–Q100)

### 97. Goroutine Without Wait
**Q: Why might this print nothing?**
```go
package main
import "fmt"

func main() {
    go func() {
        fmt.Println("hello from goroutine")
    }()
}
```
**A:** `main` returns before the goroutine gets to execute. Since `main` exiting terminates the program, the goroutine may never run. Use `time.Sleep` or `sync.WaitGroup` to wait.

---

### 98. Goroutine Count
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(runtime.NumGoroutine())
}
```
**A:** `1`. After `wg.Wait()`, all 5 goroutines have finished, leaving only the main goroutine running.

---

### 99. GOMAXPROCS
**Q: What does this print and what does it mean?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOMAXPROCS(0))
}
```
**A:** Prints the number of logical CPUs available (e.g., `8`). Passing `0` queries the current value without changing it. By default, Go uses all available CPUs.

---

### 100. Channel Direction in Function Signature
**Q: Does this compile, and what is the benefit?**
```go
package main
import "fmt"

func producer(ch chan<- int) {
    ch <- 42
}

func consumer(ch <-chan int) {
    fmt.Println(<-ch)
}

func main() {
    ch := make(chan int, 1)
    producer(ch)
    consumer(ch)
}
```
**A:** **Yes, compiles and prints** `42`. Directional channels (`chan<-` send-only, `<-chan` receive-only) restrict how a channel is used in a function, improving type safety and self-documenting intent.

---

*End of 100 Go Basics & Fundamentals Code Snippet Questions*
