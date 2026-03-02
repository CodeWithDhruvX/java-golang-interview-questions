# 📘 01 — Go Basics & Syntax
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy

---

## 🔑 Must-Know Topics
- Variable declaration (`var`, `:=`, `const`)
- Data types and zero values
- Control flow (`if`, `for`, `switch`)
- Functions, multiple return values
- Defer, panic, recover
- Packages and imports
- Pointers fundamentals

---

## ❓ Most Asked Questions

### Q1. What is Go? What are its key features?
**Answer:**
Go (Golang) is an open-source, statically-typed, compiled language by Google. Key features:
- **Fast compilation** — compiles to native binaries
- **Garbage collected** — automatic memory management
- **Concurrency-first** — goroutines and channels built-in
- **Simplicity** — minimal syntax, no classes/inheritance
- **Strong standard library** — `net/http`, `encoding/json`, etc.
- **Cross-platform** — compile for any OS/arch with `GOOS`, `GOARCH`

---

### Q2. What is the difference between `var`, `:=`, and `const`?

```go
// var — explicit declaration, can be at package or function level
var name string = "Go"
var count int        // zero value: 0

// := — short declaration, only inside functions, type inferred
message := "Hello"

// const — compile-time constant, cannot be changed
const Pi = 3.14159
const MaxConn = 100
```

| Feature | `var` | `:=` | `const` |
|---------|-------|------|---------|
| Scope | Package + Function | Function only | Package + Function |
| Type inference | Optional | Yes | Yes |
| Mutable | Yes | Yes | ❌ No |
| Zero value | Yes | No | No |

---

### Q3. What are Go's zero values?

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
var sl []int    // nil
var m map[string]int  // nil
var fn func()   // nil
```
> Zero values ensure **no uninitialized memory** — every variable has a safe default.

---

### Q4. How does `defer` work in Go?

```go
func readFile(path string) {
    f, _ := os.Open(path)
    defer f.Close()  // runs AFTER function returns, in LIFO order
    // ... read file
}

// Multiple defers — LIFO (Last In, First Out)
func example() {
    defer fmt.Println("third")
    defer fmt.Println("second")
    defer fmt.Println("first")
}
// Output: first, second, third
```
> **Key facts:** deferred functions run even if panic occurs; they can read/modify named return values.

---

### Q5. What is the difference between `new()` and `make()`?

```go
// new() — allocates zeroed memory, returns *T (pointer)
p := new(int)   // *int pointing to 0
*p = 42

// make() — initializes slice, map, or channel with internal structure
sl := make([]int, 5, 10)        // len=5, cap=10
m  := make(map[string]int)      // ready-to-use map
ch := make(chan int, 5)         // buffered channel, cap=5
```

| | `new(T)` | `make(T, ...)` |
|--|---------|----------------|
| Returns | `*T` | `T` |
| Works with | Any type | slice, map, chan only |
| Purpose | Allocate & zero | Initialize with structure |

---

### Q6. What is a variadic function?

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)           // 6
nums := []int{4, 5, 6}
sum(nums...)           // spread slice — 15
```

---

### Q7. What are blank identifiers and when are they used?

```go
// Ignore return values you don't need
value, _ := strconv.Atoi("42")

// Ignore map existence check
m := map[string]int{"a": 1}
v, _ := m["b"]  // v=0, no error

// Force interface compliance at compile time
var _ io.Writer = (*MyWriter)(nil)

// Import for side effects only
import _ "github.com/lib/pq"
```

---

### Q8. What is the difference between `break`, `continue`, and `goto`?

```go
// break — exits the innermost for/switch/select
for i := 0; i < 10; i++ {
    if i == 5 { break }
}

// continue — skips to next iteration
for i := 0; i < 10; i++ {
    if i%2 == 0 { continue }
    fmt.Println(i)  // prints odd numbers
}

// goto — jumps to a label (use sparingly)
    goto End
End:
    fmt.Println("done")

// Labeled break — exits outer loop
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 { break outer }
    }
}
```

---

### Q9. What are named return values?

```go
// Named return values allow bare return statements
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return  // bare return — returns result=0, err=error
    }
    result = a / b
    return  // returns result and nil error
}
```
> **Use case:** Helpful for documenting what a function returns and for deferred error handling.

---

### Q10. How do you convert between types in Go?

```go
// Explicit type conversion (not implicit like C)
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversions
s := strconv.Itoa(42)           // int → string
n, _ := strconv.Atoi("42")     // string → int
b := []byte("hello")           // string → []byte
str := string(b)               // []byte → string

// Cannot assign incompatible types
type Celsius float64
type Fahrenheit float64
c := Celsius(100)
// f := Fahrenheit(c)  ✅ — same underlying type
```

---

### Q11. What is `init()` function?

```go
package main

import "fmt"

var config string

func init() {
    // Runs before main(), after package-level variables
    config = "loaded"
    fmt.Println("init called")
}

func main() {
    fmt.Println("main called, config:", config)
}
// Output: init called → main called, config: loaded
```
> - A package can have **multiple** `init()` functions
> - Called in order of source file appearance
> - Cannot be called explicitly

---

### Q12. What is a type alias vs type definition?

```go
// Type alias — just another name for the same type
type Celsius = float64  // Celsius IS float64

// Type definition — creates a NEW distinct type
type Fahrenheit float64  // Fahrenheit is NOT float64

var c Celsius = 100.0
var f Fahrenheit = Fahrenheit(c)  // explicit conversion needed
```

---

### Q13. What are Go's built-in functions?

| Function | Purpose |
|----------|---------|
| `make` | Create slice/map/chan |
| `new` | Allocate memory |
| `len` | Length of string/slice/map/array |
| `cap` | Capacity of slice/channel |
| `append` | Add elements to slice |
| `copy` | Copy slice elements |
| `delete` | Remove map key |
| `close` | Close a channel |
| `panic` | Start panic sequence |
| `recover` | Recover from panic |
| `print`, `println` | Low-level debug print |

---

### Q14. Explain the `for` loop variants in Go

```go
// Traditional (C-style)
for i := 0; i < 5; i++ { }

// While-style
n := 0
for n < 5 { n++ }

// Infinite loop
for { break }

// Range over slice
nums := []int{1, 2, 3}
for i, v := range nums { fmt.Println(i, v) }

// Range over map
m := map[string]int{"a": 1}
for k, v := range m { fmt.Println(k, v) }

// Range over string (runes)
for i, r := range "hello" { fmt.Printf("%d: %c\n", i, r) }
```

---

### Q15. What is the `switch` statement in Go?

```go
// No fallthrough by default (unlike C)
switch x := 42; x {
case 1, 2:
    fmt.Println("one or two")
case 42:
    fmt.Println("the answer")
default:
    fmt.Println("other")
}

// Expression-less switch (like if-else chain)
x := 15
switch {
case x < 10:
    fmt.Println("small")
case x < 20:
    fmt.Println("medium")
default:
    fmt.Println("large")
}

// Explicit fallthrough
switch 1 {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two also printed")
}
```
