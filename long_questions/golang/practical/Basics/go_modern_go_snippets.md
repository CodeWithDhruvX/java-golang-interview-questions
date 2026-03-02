# Go Modern Snippets — Pure Code Snippet Interview Questions

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Topics covered**: `new` vs `make` · Slice tricks · Map gotchas · Struct embedding · Interface nil trap · Generics · `errors.Is`/`As` · `select` · Channel directions · `sync.Mutex` / `sync.Once` / `sync.RWMutex`

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** `new` vs `make`, Slice & Map Gotchas (Q1–Q16)
- [ ] **Section 2:** Struct Embedding & Promoted Methods (Q17–Q25)
- [ ] **Section 3:** The Interface nil Trap & errors.Is / errors.As (Q26–Q38)
- [ ] **Section 4:** Generics (Q39–Q52)
- [ ] **Section 5:** `select`, Channel Directions & Patterns (Q53–Q66)
- [ ] **Section 6:** sync.Mutex, sync.RWMutex, sync.Once, sync.Map (Q67–Q80)

> 🔖 **Last read:** <!-- e.g. Q25 · Section 2 done -->

---

## Section 1: `new` vs `make`, Slice & Map Gotchas (Q1–Q16)

### 1. `new` vs `make` — What's the Difference?
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    p := new(int)
    fmt.Println(*p)

    s := make([]int, 3)
    fmt.Println(s)
}
```
**A:**
```
0
[0 0 0]
```
`new(T)` allocates zeroed memory and returns a `*T`. `make` is only for slices, maps, and channels — it initialises internal data structures and returns the value itself (not a pointer).

---

### 2. `new` on a Struct
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func main() {
    p := new(Point)
    p.X = 10
    p.Y = 20
    fmt.Println(*p)
}
```
**A:** `{10 20}`

---

### 3. nil Map — Read vs Write
**Q: Does this compile and what is the runtime behaviour?**
```go
package main
import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["key"]) // line A
    m["key"] = 1          // line B
}
```
**A:** Line A compiles and prints `0` (reading a nil map returns the zero value). Line B causes a **runtime panic**: `assignment to entry in nil map`.

---

### 4. nil vs Empty Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s1 []int
    s2 := []int{}
    s3 := make([]int, 0)

    fmt.Println(s1 == nil, len(s1), cap(s1))
    fmt.Println(s2 == nil, len(s2), cap(s2))
    fmt.Println(s3 == nil, len(s3), cap(s3))
}
```
**A:**
```
true 0 0
false 0 0
false 0 0
```
`var s1 []int` is nil. `[]int{}` and `make([]int, 0)` are non-nil empty slices.

---

### 5. Slice Sharing Underlying Array
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3]
    b[0] = 99
    fmt.Println(a)
    fmt.Println(b)
}
```
**A:**
```
[1 99 3 4 5]
[99 3]
```
`b` is a slice of `a` sharing the same backing array. Modifying `b[0]` modifies `a[1]`.

---

### 6. Append Breaking the Shared Array
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := a[:2]
    b = append(b, 99)
    fmt.Println(a)
    fmt.Println(b)
}
```
**A:**
```
[1 2 99]
[1 2 99]
```
`b` has capacity 3 (same as `a`). `append` fits within capacity, so it writes `99` into `a[2]` — modifying the original!

---

### 7. Append Forcing a New Backing Array
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := make([]int, len(a), len(a)) // cap == len, no extra room
    copy(b, a)
    b = append(b, 99)
    b[0] = 777
    fmt.Println(a)
    fmt.Println(b)
}
```
**A:**
```
[1 2 3]
[777 2 3 99]
```
`copy` creates an independent slice. `append` exceeds capacity → new backing array allocated. `a` is untouched.

---

### 8. 2D Slice Initialisation
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    matrix := make([][]int, 3)
    for i := range matrix {
        matrix[i] = make([]int, 3)
    }
    matrix[1][1] = 42
    fmt.Println(matrix)
}
```
**A:** `[[0 0 0] [0 42 0] [0 0 0]]`. Each row must be separately initialised; they do not share memory.

---

### 9. Map Iteration Order
**Q: Is the output guaranteed?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
**A:** **No.** Map iteration order in Go is intentionally randomised. Never rely on it.

---

### 10. Deleting from a Map During Range
**Q: Does this compile and is it safe?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for k := range m {
        delete(m, k)
    }
    fmt.Println(m)
}
```
**A:** **Yes, it compiles and is safe.** Go explicitly allows deleting map entries during `range`. Output: `map[]`.

---

### 11. Map with Struct Values — Cannot Address Fields
**Q: Does this compile?**
```go
package main

type Point struct{ X, Y int }

func main() {
    m := map[string]Point{"origin": {0, 0}}
    m["origin"].X = 1
}
```
**A:** **Compile Error.** `cannot assign to struct field m["origin"].X in map`. Map values are not addressable. Fix: assign the whole struct — `p := m["origin"]; p.X = 1; m["origin"] = p`.

---

### 12. Slice Passed to Function — Mutation vs Replacement
**Q: What is the output?**
```go
package main
import "fmt"

func modify(s []int) {
    s[0] = 99       // mutates original backing array
    s = append(s, 100) // new backing array; caller's slice unchanged
}

func main() {
    a := []int{1, 2, 3}
    modify(a)
    fmt.Println(a)
}
```
**A:** `[99 2 3]`. `s[0] = 99` is visible because the slice header shares the backing array. But `append` inside `modify` creates a new backing array visible only locally — the caller's `a` still has length 3.

---

### 13. `copy` Return Value
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
**A:** `3 [1 2 3]`. `copy` returns the number of elements copied — `min(len(dst), len(src))`.

---

### 14. Map of Slices — Zero Value is nil
**Q: Does this panic?**
```go
package main
import "fmt"

func main() {
    m := make(map[string][]int)
    m["a"] = append(m["a"], 1, 2, 3)
    fmt.Println(m["a"])
}
```
**A:** **No panic.** `m["a"]` returns a nil slice (zero value). `append` on a nil slice works fine. Output: `[1 2 3]`.

---

### 15. Slice of Pointers vs Slice of Values
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    nums := []int{1, 2, 3}
    ptrs := make([]*int, len(nums))
    for i := range nums {
        ptrs[i] = &nums[i]
    }
    nums[0] = 99
    fmt.Println(*ptrs[0])
}
```
**A:** `99`. `ptrs[0]` holds the address of `nums[0]`. Changing `nums[0]` is reflected through the pointer.

---

### 16. make with Length and Capacity
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Println(len(s), cap(s))
    s = append(s, 10)
    fmt.Println(len(s), cap(s), s)
}
```
**A:**
```
3 5
4 5 [0 0 0 10]
```
`make([]int, 3, 5)` creates a slice with 3 zeroed elements and capacity 5. `append` uses the preallocated capacity.

---

## Section 2: Struct Embedding & Promoted Methods (Q17–Q25)

### 17. Basic Embedding
**Q: What is the output?**
```go
package main
import "fmt"

type Animal struct{ Name string }

func (a Animal) Speak() string { return a.Name + " speaks" }

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Lab"}
    fmt.Println(d.Speak())
    fmt.Println(d.Name)
}
```
**A:**
```
Rex speaks
Rex
```
`Dog` promotes `Animal`'s fields and methods. `d.Speak()` and `d.Name` both work directly.

---

### 18. Embedding Overrides (Shadow)
**Q: What is the output?**
```go
package main
import "fmt"

type Base struct{}
func (Base) Hello() string { return "Base" }

type Child struct{ Base }
func (Child) Hello() string { return "Child" }

func main() {
    c := Child{}
    fmt.Println(c.Hello())
    fmt.Println(c.Base.Hello())
}
```
**A:**
```
Child
Base
```
`Child.Hello` shadows `Base.Hello`. The embedded method is still accessible via `c.Base.Hello()`.

---

### 19. Embedding an Interface
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

type Stringer interface{ String() string }

type Wrapper struct{ Stringer }

func main() {
    w := Wrapper{}
    fmt.Println(w.Stringer)
}
```
**A:** **Compiles.** Prints `<nil>`. Embedding an interface in a struct means `Wrapper` satisfies `Stringer`, but the embedded field is nil unless set. Calling `w.String()` would **panic** at runtime.

---

### 20. Pointer Receiver Embedded Method
**Q: Does this compile?**
```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

type Widget struct{ Counter }

func main() {
    w := Widget{}
    w.Inc()
    fmt.Println(w.n)
}
```
**A:** **Compiles and prints** `1`. Promoted pointer-receiver methods are accessible on the embedding struct when it's addressable.

---

### 21. Embedding Name Conflict — Ambiguous Selector
**Q: Does this compile?**
```go
package main

type A struct{ Val int }
type B struct{ Val int }
type C struct{ A; B }

func main() {
    c := C{}
    _ = c.Val
}
```
**A:** **Compile Error.** `ambiguous selector c.Val` — both `A.Val` and `B.Val` are at the same depth. You must disambiguate: `c.A.Val` or `c.B.Val`.

---

### 22. Struct Tags
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"-"`
}

func main() {
    u := User{Name: "Alice", Email: "", Age: 30}
    b, _ := json.Marshal(u)
    fmt.Println(string(b))
}
```
**A:** `{"name":"Alice"}`. `omitempty` omits `Email` because it's empty. `"-"` causes `Age` to always be omitted.

---

### 23. Anonymous Struct
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    p := struct {
        X, Y int
    }{X: 3, Y: 4}
    fmt.Println(p.X + p.Y)
}
```
**A:** `7`. Anonymous structs are valid for one-off data groupings without declaring a named type.

---

### 24. Struct Comparison
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func main() {
    a := Point{1, 2}
    b := Point{1, 2}
    fmt.Println(a == b)
}
```
**A:** `true`. Structs are comparable if all their fields are comparable.

---

### 25. Struct with Slice Field — Not Comparable
**Q: Does this compile?**
```go
package main

type Data struct{ Items []int }

func main() {
    a := Data{Items: []int{1, 2}}
    b := Data{Items: []int{1, 2}}
    _ = a == b
}
```
**A:** **Compile Error.** `invalid operation: a == b (struct containing []int cannot be compared)`. Slices are not comparable.

---

## Section 3: The Interface nil Trap & errors.Is / errors.As (Q26–Q38)

### 26. The Classic Interface nil Trap
**Q: What is the output?**
```go
package main
import "fmt"

type MyError struct{ msg string }
func (e *MyError) Error() string { return e.msg }

func getError() error {
    var err *MyError = nil
    return err
}

func main() {
    err := getError()
    fmt.Println(err == nil)
    fmt.Println(err)
}
```
**A:**
```
false
<nil>
```
The `error` interface value is **not nil** because it holds a non-nil *type* (`*MyError`) even though the *value* pointer is nil. An interface is nil only if both its type and value are nil.

---

### 27. Correct Way to Return nil Error
**Q: What is the output?**
```go
package main
import "fmt"

type MyError struct{ msg string }
func (e *MyError) Error() string { return e.msg }

func noError() error {
    return nil // returns a nil interface — both type and value are nil
}

func main() {
    err := noError()
    fmt.Println(err == nil)
}
```
**A:** `true`. Returning `nil` untyped directly as `error` sets both the interface type and value to nil.

---

### 28. errors.New vs fmt.Errorf
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

func main() {
    e1 := errors.New("disk full")
    e2 := fmt.Errorf("write failed: %w", e1)
    fmt.Println(e2)
    fmt.Println(errors.Is(e2, e1))
}
```
**A:**
```
write failed: disk full
true
```
`%w` wraps the error. `errors.Is` unwraps the chain and finds `e1`.

---

### 29. errors.Is with Sentinel Errors
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func find(id int) error {
    if id == 0 {
        return fmt.Errorf("find(%d): %w", id, ErrNotFound)
    }
    return nil
}

func main() {
    err := find(0)
    fmt.Println(errors.Is(err, ErrNotFound))
}
```
**A:** `true`. `errors.Is` traverses the wrapped chain.

---

### 30. errors.As
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
    return fmt.Errorf("request failed: %w", &ValidationError{Field: "email"})
}

func main() {
    err := validate()
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Println("field:", ve.Field)
    }
}
```
**A:** `field: email`. `errors.As` unwraps the error chain and extracts the first matching type.

---

### 31. errors.Is Does Not Use == for Wrapped Errors
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrTimeout = errors.New("timeout")

func main() {
    wrapped := fmt.Errorf("op failed: %w", ErrTimeout)
    fmt.Println(wrapped == ErrTimeout)
    fmt.Println(errors.Is(wrapped, ErrTimeout))
}
```
**A:**
```
false
true
```
`==` compares the outer error value (not equal). `errors.Is` unwraps and matches.

---

### 32. Returning Concrete Error Type vs interface — Nil Trap in Practice
**Q: Does the caller get nil?**
```go
package main
import "fmt"

type DBError struct{ Code int }
func (e *DBError) Error() string { return fmt.Sprintf("db error %d", e.Code) }

func query(ok bool) *DBError {
    if !ok { return &DBError{Code: 500} }
    return nil
}

func run() error {
    return query(true) // BUG: returns (*DBError)(nil) as error interface
}

func main() {
    err := run()
    fmt.Println(err == nil)
}
```
**A:** `false`. `query(true)` returns `(*DBError)(nil)`. Assigning that to an `error` interface wraps it in a non-nil interface. **Fix:** change `run()` to return `nil` directly when there's no error, or change `query` to return `error`.

---

### 33. Custom Is Method
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

type CodeError struct{ Code int }
func (e *CodeError) Error() string { return fmt.Sprintf("code %d", e.Code) }
func (e *CodeError) Is(target error) bool {
    t, ok := target.(*CodeError)
    return ok && t.Code == e.Code
}

func main() {
    e1 := &CodeError{Code: 404}
    e2 := &CodeError{Code: 404}
    fmt.Println(errors.Is(e1, e2))
}
```
**A:** `true`. The custom `Is` method allows `errors.Is` to compare by value instead of pointer identity.

---

### 34. Unwrap Chain Length
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrRoot = errors.New("root")

func main() {
    e1 := fmt.Errorf("layer1: %w", ErrRoot)
    e2 := fmt.Errorf("layer2: %w", e1)
    e3 := fmt.Errorf("layer3: %w", e2)
    fmt.Println(errors.Is(e3, ErrRoot))
}
```
**A:** `true`. `errors.Is` recursively unwraps through all three layers.

---

### 35. Interface Satisfaction — Compile Time Check
**Q: Does this compile?**
```go
package main

type Writer interface{ Write([]byte) (int, error) }

type NullWriter struct{}

var _ Writer = NullWriter{} // compile-time assertion
```
**A:** **Compile Error.** `NullWriter` does not implement `Writer` (missing `Write` method). The `var _ Writer = ...` pattern is a common idiom to get a compile-time interface satisfaction check.

---

### 36. Type Assertion — Comma OK Pattern
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i interface{} = "hello"

    s, ok := i.(string)
    fmt.Println(s, ok)

    n, ok := i.(int)
    fmt.Println(n, ok)
}
```
**A:**
```
hello true
0 false
```
The comma-ok form of type assertion never panics — it returns the zero value and `false` on failure.

---

### 37. Panic on Failed Type Assertion (No Comma OK)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    var i interface{} = "hello"
    _ = i.(int)
}
```
**A:** `recovered: interface conversion: interface {} is string, not int`. A type assertion without comma-ok panics on failure.

---

### 38. Empty Interface vs `any`
**Q: Do these mean the same thing?**
```go
package main
import "fmt"

func print1(v interface{}) { fmt.Println(v) }
func print2(v any)         { fmt.Println(v) }

func main() {
    print1(42)
    print2(42)
}
```
**A:**
```
42
42
```
`any` is an alias for `interface{}` introduced in Go 1.18. They are completely interchangeable.

---

## Section 4: Generics (Q39–Q52)

### 39. Basic Generic Function
**Q: What is the output?**
```go
package main
import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

func main() {
    doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
    fmt.Println(doubled)
}
```
**A:** `[2 4 6]`

---

### 40. Type Constraint — `comparable`
**Q: Does this compile?**
```go
package main
import "fmt"

func Contains[T comparable](s []T, v T) bool {
    for _, x := range s {
        if x == v { return true }
    }
    return false
}

func main() {
    fmt.Println(Contains([]string{"a", "b", "c"}, "b"))
    fmt.Println(Contains([]int{1, 2, 3}, 4))
}
```
**A:** **Yes.** Output:
```
true
false
```
`comparable` constrains `T` to types that support `==`.

---

### 41. Type Constraint Interface
**Q: Does this compile?**
```go
package main
import "fmt"

type Number interface {
    int | int64 | float64
}

func Sum[T Number](s []T) T {
    var total T
    for _, v := range s {
        total += v
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
    fmt.Println(Sum([]float64{1.1, 2.2}))
}
```
**A:** **Yes.** Output:
```
15
3.3000000000000003
```

---

### 42. Generic with Pointer Receiver — Common Mistake
**Q: Does this compile?**
```go
package main

type Setter[T any] interface {
    Set(T)
}

type Box[T any] struct{ val T }
func (b *Box[T]) Set(v T) { b.val = v }

func Fill[T any, S Setter[T]](s S, v T) {
    s.Set(v)
}

func main() {
    b := &Box[int]{}
    Fill(b, 42)
}
```
**A:** **Compiles.** `*Box[int]` satisfies `Setter[int]` because `Set` has a pointer receiver.

---

### 43. Type Parameter Cannot Be Used in Type Switch
**Q: Does this compile?**
```go
package main
import "fmt"

func Describe[T any](v T) {
    switch v.(type) {
    case int:
        fmt.Println("int")
    }
}

func main() {
    Describe(42)
}
```
**A:** **Compile Error.** You cannot use a type switch on a generic type parameter `T` unless it is constrained to `interface{}` — and even then this is restricted. Use `any` tricks or reflection for dynamic dispatch.

---

### 44. Generic Stack
**Q: What is the output?**
```go
package main
import "fmt"

type Stack[T any] struct{ items []T }
func (s *Stack[T]) Push(v T)  { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() T    { n := len(s.items) - 1; v := s.items[n]; s.items = s.items[:n]; return v }
func (s *Stack[T]) Len() int  { return len(s.items) }

func main() {
    s := Stack[string]{}
    s.Push("a")
    s.Push("b")
    s.Push("c")
    fmt.Println(s.Pop())
    fmt.Println(s.Len())
}
```
**A:**
```
c
2
```

---

### 45. `~` Tilde in Constraints (Underlying Type)
**Q: Does this compile?**
```go
package main
import "fmt"

type MyInt int

type Numeric interface{ ~int | ~float64 }

func Double[T Numeric](v T) T { return v * 2 }

func main() {
    var x MyInt = 5
    fmt.Println(Double(x))
}
```
**A:** **Yes.** Output: `10`. `~int` means "any type whose **underlying type** is `int`", which includes `MyInt`.

---

### 46. Generic Filter
**Q: What is the output?**
```go
package main
import "fmt"

func Filter[T any](s []T, pred func(T) bool) []T {
    var result []T
    for _, v := range s {
        if pred(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens)
}
```
**A:** `[2 4 6]`

---

### 47. Generic Reduce
**Q: What is the output?**
```go
package main
import "fmt"

func Reduce[T, Acc any](s []T, init Acc, f func(Acc, T) Acc) Acc {
    acc := init
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func main() {
    product := Reduce([]int{1, 2, 3, 4}, 1, func(acc, v int) int { return acc * v })
    fmt.Println(product)
}
```
**A:** `24`

---

### 48. Type Inference in Generics
**Q: Does the call compile without explicit type arguments?**
```go
package main
import "fmt"

func Pair[T, U any](a T, b U) (T, U) { return a, b }

func main() {
    x, y := Pair(1, "hello") // no explicit [int, string]
    fmt.Println(x, y)
}
```
**A:** **Yes.** Go infers type arguments from call-site argument types. Output: `1 hello`.

---

### 49. Comparable Constraint for Map Key Generic Function
**Q: What is the output?**
```go
package main
import "fmt"

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func main() {
    m := map[string]int{"x": 1, "y": 2}
    fmt.Println(len(Keys(m)))
}
```
**A:** `2` (order not guaranteed).

---

### 50. Generic Function Cannot Have Methods
**Q: Does this compile?**
```go
package main

func Map[T any](s []T, f func(T) T) []T { return nil }

func (Map[T]) String() string { return "" }
```
**A:** **Compile Error.** Functions (as opposed to types) cannot have methods in Go. Only generic *types* can have methods.

---

### 51. `min` / `max` Builtins (Go 1.21+)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    fmt.Println(min(3, 1, 4, 1, 5))
    fmt.Println(max(3, 1, 4, 1, 5))
}
```
**A:**
```
1
5
```
`min` and `max` are built-in generics added in Go 1.21 that work on any ordered type.

---

### 52. `clear` Builtin (Go 1.21+)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    clear(m)
    fmt.Println(m)

    s := []int{1, 2, 3}
    clear(s)
    fmt.Println(s)
}
```
**A:**
```
map[]
[0 0 0]
```
`clear(map)` removes all entries. `clear(slice)` zeroes all elements (but keeps length and capacity).

---

## Section 5: `select`, Channel Directions & Patterns (Q53–Q66)

### 53. Basic select
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)
    ch1 <- "one"
    ch2 <- "two"
    select {
    case msg := <-ch1:
        fmt.Println("ch1:", msg)
    case msg := <-ch2:
        fmt.Println("ch2:", msg)
    }
}
```
**A:** Either `ch1: one` or `ch2: two` (non-deterministic). When multiple cases are ready, `select` picks one at random.

---

### 54. select with default — Non-Blocking Check
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int)
    select {
    case v := <-ch:
        fmt.Println("received:", v)
    default:
        fmt.Println("no value ready")
    }
}
```
**A:** `no value ready`. The `default` case prevents blocking — it fires immediately when no other case is ready.

---

### 55. Timeout with select and time.After
**Q: What does this pattern do?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    select {
    case v := <-ch:
        fmt.Println("got:", v)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("timeout")
    }
}
```
**A:** Prints `timeout`. `time.After` returns a channel that sends after the duration. This is the canonical Go timeout pattern.

---

### 56. Sending and Receiving in select
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 1)
    select {
    case ch <- 42:
        fmt.Println("sent")
    default:
        fmt.Println("channel full")
    }

    select {
    case v := <-ch:
        fmt.Println("received:", v)
    default:
        fmt.Println("nothing")
    }
}
```
**A:**
```
sent
received: 42
```

---

### 57. Channel Direction — Send-Only
**Q: Does this compile?**
```go
package main
import "fmt"

func produce(ch chan<- int) {
    ch <- 1
    ch <- 2
    close(ch)
}

func main() {
    ch := make(chan int, 3)
    produce(ch)
    for v := range ch {
        fmt.Print(v, " ")
    }
}
```
**A:** **Yes.** Output: `1 2 `. `chan<-` is a send-only channel. Attempting to receive from it inside `produce` would be a compile error.

---

### 58. Channel Direction — Receive-Only
**Q: Does this compile?**
```go
package main
import "fmt"

func consume(ch <-chan int) {
    for v := range ch {
        fmt.Print(v, " ")
    }
}

func main() {
    ch := make(chan int, 2)
    ch <- 10
    ch <- 20
    close(ch)
    consume(ch)
}
```
**A:** **Yes.** Output: `10 20 `. `<-chan` is a receive-only channel. Sending to it inside `consume` would be a compile error.

---

### 59. Closing an Already Closed Channel Panics
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    ch := make(chan int)
    close(ch)
    close(ch)
}
```
**A:** `recovered: close of closed channel`. Closing an already-closed channel panics.

---

### 60. Sending on a Closed Channel Panics
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    ch := make(chan int, 1)
    close(ch)
    ch <- 1
}
```
**A:** `recovered: send on closed channel`.

---

### 61. Receiving from Closed Channel — comma-ok
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)

    for {
        v, ok := <-ch
        if !ok {
            fmt.Println("channel closed")
            break
        }
        fmt.Println(v)
    }
}
```
**A:**
```
1
2
channel closed
```
Receiving from a closed, drained channel returns the zero value and `false`.

---

### 62. nil Channel Blocks Forever
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var ch chan int
    select {
    case v := <-ch:
        fmt.Println(v)
    default:
        fmt.Println("nil channel skipped")
    }
}
```
**A:** `nil channel skipped`. Receiving (or sending) on a nil channel blocks forever in a `select` — that case is never ready, so `default` fires.

---

### 63. Fan-Out Pattern
**Q: What does this do?**
```go
package main
import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("worker %d processed job %d\n", id, j)
    }
}

func main() {
    jobs := make(chan int, 5)
    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    wg.Wait()
}
```
**A:** 5 jobs distributed across 3 workers. Output order varies, but all 5 jobs are processed exactly once. This is the canonical fan-out (worker pool) pattern.

---

### 64. Done Channel — Cancellation Pattern
**Q: What does this show?**
```go
package main
import (
    "fmt"
    "time"
)

func worker(done <-chan struct{}) {
    for {
        select {
        case <-done:
            fmt.Println("worker stopped")
            return
        default:
            fmt.Println("working...")
            time.Sleep(10 * time.Millisecond)
        }
    }
}

func main() {
    done := make(chan struct{})
    go worker(done)
    time.Sleep(25 * time.Millisecond)
    close(done)
    time.Sleep(20 * time.Millisecond)
}
```
**A:** Prints `working...` a few times then `worker stopped`. Closing a `done` channel is the idiomatic Go cancellation signal — all receivers unblock simultaneously.

---

### 65. select Evaluates All Channel Expressions
**Q: What is the output?**
```go
package main
import "fmt"

func getChannel() chan int {
    fmt.Println("getChannel called")
    return nil
}

func main() {
    select {
    case <-getChannel():
    default:
        fmt.Println("default")
    }
}
```
**A:**
```
getChannel called
default
```
All channel expressions in a `select` are evaluated once before the selection logic runs.

---

### 66. Buffered Channel as Semaphore
**Q: What is the pattern?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    sem := make(chan struct{}, 3) // max 3 concurrent
    var wg sync.WaitGroup
    for i := 1; i <= 6; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sem <- struct{}{}        // acquire
            defer func() { <-sem }() // release
            fmt.Printf("running %d\n", id)
        }(i)
    }
    wg.Wait()
}
```
**A:** All 6 goroutines complete, but at most 3 run concurrently. A buffered channel of size N is a classic counting semaphore.

---

## Section 6: sync.Mutex, sync.RWMutex, sync.Once, sync.Map (Q67–Q80)

### 67. Race Condition Without Mutex
**Q: What is the bug?**
```go
package main
import (
    "fmt"
    "sync"
)

var counter int

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // DATA RACE
        }()
    }
    wg.Wait()
    fmt.Println(counter) // likely < 1000
}
```
**A:** `counter` is less than 1000 (usually). This is a data race — concurrent goroutines read and write `counter` without synchronisation. Detect with `go run -race`.

---

### 68. Fixing with sync.Mutex
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

var (
    counter int
    mu      sync.Mutex
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}
```
**A:** `1000`. Every increment is now protected by the mutex.

---

### 69. Mutex is Not Reentrant
**Q: What happens?**
```go
package main
import "sync"

var mu sync.Mutex

func a() {
    mu.Lock()
    defer mu.Unlock()
    b() // calls b() while holding the lock
}

func b() {
    mu.Lock() // tries to acquire the same mutex → deadlock
    defer mu.Unlock()
}

func main() { a() }
```
**A:** **Deadlock.** Go's `sync.Mutex` is NOT reentrant. A goroutine that tries to re-lock a mutex it already holds will deadlock forever.

---

### 70. sync.RWMutex — Multiple Readers
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

var (
    data  = 0
    rwmu  sync.RWMutex
)

func readData() int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data
}

func writeData(v int) {
    rwmu.Lock()
    defer rwmu.Unlock()
    data = v
}

func main() {
    writeData(42)
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println(readData())
        }()
    }
    wg.Wait()
}
```
**A:** Five lines of `42`. Multiple goroutines can hold `RLock` simultaneously. A `Lock` (write lock) is exclusive.

---

### 71. Defer Unlock — Always the Right Pattern
**Q: Which version is safer?**
```go
// Version A
mu.Lock()
doSomething()
mu.Unlock() // not called if doSomething panics

// Version B
mu.Lock()
defer mu.Unlock() // always runs, even on panic
doSomething()
```
**A:** **Version B.** `defer mu.Unlock()` ensures the mutex is always released even if the function panics or returns early.

---

### 72. sync.Once — Execute Exactly Once
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

var once sync.Once

func setup() {
    fmt.Println("setup called")
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            once.Do(setup)
        }()
    }
    wg.Wait()
}
```
**A:** `setup called` — printed exactly **once**, regardless of how many goroutines call `once.Do`. This is the canonical lazy-initialisation pattern.

---

### 73. sync.Once for Singleton
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

type Config struct{ DSN string }

var (
    cfg     *Config
    cfgOnce sync.Once
)

func GetConfig() *Config {
    cfgOnce.Do(func() {
        cfg = &Config{DSN: "postgres://localhost/db"}
    })
    return cfg
}

func main() {
    c1 := GetConfig()
    c2 := GetConfig()
    fmt.Println(c1 == c2)
    fmt.Println(c1.DSN)
}
```
**A:**
```
true
postgres://localhost/db
```
Both calls return the same pointer — the singleton is initialised only once.

---

### 74. sync.Map — Concurrent Safe Map
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            m.Store(n, n*n)
        }(i)
    }
    wg.Wait()

    m.Range(func(k, v any) bool {
        fmt.Println(k, v)
        return true
    })
}
```
**A:** 5 key-value pairs (order non-deterministic), e.g.:
```
0 0
1 1
2 4
3 9
4 16
```
`sync.Map` is safe for concurrent use without a mutex — optimised for read-heavy or mostly-disjoint-key workloads.

---

### 75. sync.Map LoadOrStore
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    actual, loaded := m.LoadOrStore("key", "first")
    fmt.Println(actual, loaded)

    actual, loaded = m.LoadOrStore("key", "second")
    fmt.Println(actual, loaded)
}
```
**A:**
```
first false
first true
```
`LoadOrStore` atomically stores and returns the value if not present (`loaded=false`), or returns the existing value (`loaded=true`).

---

### 76. sync.Cond — Broadcast vs Signal
**Q: What is the typical output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    ready := false

    for i := 0; i < 3; i++ {
        go func(id int) {
            mu.Lock()
            for !ready {
                cond.Wait()
            }
            fmt.Printf("goroutine %d woke up\n", id)
            mu.Unlock()
        }(i)
    }

    // Give goroutines time to wait
    // (in real code, use a proper synchronisation mechanism)
    ready = true
    cond.Broadcast() // wakes all waiting goroutines
    // Wait for output...
}
```
**A:** All 3 goroutines print "woke up" (order varies). `Broadcast` wakes all waiters; `Signal` wakes exactly one.

---

### 77. sync.WaitGroup — Misuse: Add Inside Goroutine
**Q: What is the bug?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        go func(n int) {
            wg.Add(1)       // BUG: Add called inside goroutine
            defer wg.Done()
            fmt.Println(n)
        }(i)
    }
    wg.Wait()
}
```
**A:** **Race condition / premature return.** `wg.Wait()` may return before `wg.Add(1)` is called by any goroutine. Always call `wg.Add` **before** launching the goroutine.

---

### 78. atomic Operations vs Mutex
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}
```
**A:** `1000`. `atomic.AddInt64` is lock-free and faster than a mutex for simple integer operations.

---

### 79. sync.Pool — Object Reuse
**Q: What is the usage pattern?**
```go
package main
import (
    "fmt"
    "sync"
)

var pool = sync.Pool{
    New: func() any { return make([]byte, 1024) },
}

func main() {
    buf := pool.Get().([]byte)
    // use buf...
    buf = buf[:0]
    pool.Put(buf)
    fmt.Println("buf reused")
}
```
**A:** `buf reused`. `sync.Pool` reduces GC pressure by reusing allocated objects. The `New` function is called only when the pool is empty. Note: objects may be collected by GC between GC cycles.

---

### 80. Mutex Lock / Unlock on Different Goroutines
**Q: Is this valid?**
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

var mu sync.Mutex

func main() {
    mu.Lock()
    go func() {
        time.Sleep(10 * time.Millisecond)
        fmt.Println("unlocking from goroutine")
        mu.Unlock() // unlocking from a different goroutine
    }()
    mu.Lock() // waits for the goroutine to unlock
    fmt.Println("main acquired lock again")
    mu.Unlock()
}
```
**A:** **Valid in Go.** Unlike some languages, `sync.Mutex` in Go does NOT track ownership — any goroutine can unlock it. Output:
```
unlocking from goroutine
main acquired lock again
```
(This is unusual but legal — be careful with such patterns.)

---
