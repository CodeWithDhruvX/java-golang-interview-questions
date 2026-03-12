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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `0` then `[0 0 0]`. `new(int)` allocates memory for an int and returns a pointer to it, initialized to zero. `make([]int, 3)` creates a slice with 3 zeroed elements. The key difference is that `new` returns a pointer while `make` returns the initialized value itself, and `make` only works for slices, maps, and channels.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `{10 20}`. `new(Point)` allocates a Point struct and returns a pointer to it, initialized with zero values. We can then dereference the pointer with `*p` or access fields directly with `p.X` and `p.Y` since Go automatically dereferences pointers for field access. This is how you create struct instances on the heap.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what is the runtime behaviour?
**Your Response:** Line A compiles and prints `0` because reading from a nil map returns the zero value for the type. Line B panics at runtime because you can't write to a nil map - it hasn't been initialized yet. You need to use `make(map[string]int)` to create an initialized map before writing to it.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This shows the difference between nil and empty slices. `var s1 []int` creates a nil slice (true, 0, 0). `[]int{}` creates an empty slice with allocated memory (false, 0, 0). `make([]int, 0)` also creates an empty slice (false, 0, 0). All have length 0 and capacity 0, but only the first is actually nil. This distinction matters when marshaling JSON or checking if a slice was explicitly set.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `[1 99 3 4 5]` then `[99 3]`. Slice `b` is created from `a[1:3]`, so it shares the same underlying array with `a`. When we modify `b[0]` to 99, we're actually modifying `a[1]` because `b[0]` points to the same memory location as `a[1]`. This is why slices are called views into underlying arrays - they don't copy the data, just point to a portion of it.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `[1 2 99]` twice. Slice `b` is created from `a[:2]` but inherits `a`'s capacity of 3. When we append 99 to `b`, it fits within the existing capacity, so Go writes to the backing array at position 2, which is shared with `a`. This is a common gotcha - appending to a slice can modify the original slice if they share the same backing array and there's enough capacity.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `[1 2 3]` then `[777 2 3 99]`. We use `copy` to create an independent copy of `a` into `b`, with capacity exactly equal to length. When we append 99 to `b`, it exceeds capacity, so Go allocates a new backing array. Now `b` has its own array, so modifying `b[0]` doesn't affect `a`. This is the correct way to create independent slices.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints a 3x3 matrix with 42 in the middle. When creating a 2D slice, `make([][]int, 3)` creates a slice of slices, but each inner slice is nil initially. You must initialize each row separately with `make([]int, 3)`. Each row gets its own backing array, so modifying one row doesn't affect others.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is the output guaranteed?
**Your Response:** No, the output is not guaranteed. Go intentionally randomizes map iteration order to prevent developers from relying on any specific ordering. This was a design decision to avoid subtle bugs. If you need ordered iteration, you must extract the keys into a slice and sort them first.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and is it safe?
**Your Response:** Yes, this compiles and is safe. Go explicitly allows deleting entries from a map while ranging over it. The iteration continues safely, and all entries get deleted. The final output is an empty map. This is useful for filtering maps or cleaning up entries based on some condition.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. Map values are not addressable, so you can't directly modify fields of a struct stored in a map. You need to read the whole struct into a variable, modify it, then put it back in the map. This limitation exists because map entries might be relocated during operations, so direct addressing would be unsafe.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `[99 2 3]`. When we pass a slice to a function, we're passing a copy of the slice header (pointer, length, capacity). The pointer still points to the same backing array, so `s[0] = 99` modifies the original array. But `append` creates a new backing array if there's no capacity, and this change is only visible inside the function - the caller's slice header is unchanged.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `3 [1 2 3]`. `copy` copies elements from the source slice to the destination slice and returns how many elements were actually copied. It copies up to the minimum of the two slice lengths. Here, the destination has length 3, so only 3 elements are copied, and `copy` returns 3.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this panic?
**Your Response:** No, this doesn't panic. When you access a map key that doesn't exist, you get the zero value for that type - for a slice, that's `nil`. `append` works fine with nil slices - it allocates a new backing array and returns a new slice. This is a convenient feature that makes working with maps of slices easy.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `99`. We create a slice of pointers where each pointer points to an element in the `nums` slice. When we change `nums[0]` to 99, the pointer `ptrs[0]` still points to that same memory location, so dereferencing it gives us the updated value of 99. This shows the difference between storing values and storing pointers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `3 5` then `4 5 [0 0 0 10]`. `make([]int, 3, 5)` creates a slice with length 3 and capacity 5. The slice has 3 zeroed elements initially. When we append 10, it fits within the existing capacity, so no new allocation is needed - the slice becomes length 4 with the same capacity of 5.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Rex speaks` then `Rex`. This shows struct embedding where `Dog` embeds `Animal`. The embedded struct's fields and methods are promoted to the embedding struct, so we can access `d.Name` and call `d.Speak()` directly as if they were defined on `Dog`. This is Go's way of achieving composition over inheritance.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Child` then `Base`. When `Child` defines its own `Hello` method, it shadows (hides) the embedded `Base.Hello` method. But the base method is still accessible by qualifying it with the embedded type name: `c.Base.Hello()`. This is how method resolution works with embedding - the most specific definition wins.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what is the output?
**Your Response:** This compiles and prints `<nil>`. When you embed an interface in a struct, the struct satisfies that interface automatically. But the embedded field is nil until you set it. If you tried to call `w.String()`, it would panic because you're calling a method on a nil interface. This pattern is useful for optional interfaces or when you want to defer implementation.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `1`. Even though `Counter`'s `Inc` method has a pointer receiver, it's promoted to `Widget` and can be called on `Widget` when it's addressable. Go automatically takes the address of `w` to call the pointer receiver method. This is why promoted methods work with both value and pointer receivers when the embedding struct is addressable.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile due to an ambiguous selector. When `C` embeds both `A` and `B`, and both have a `Val` field, Go doesn't know which one you mean when you write `c.Val`. You must explicitly qualify it as either `c.A.Val` or `c.B.Val` to disambiguate. This prevents subtle bugs when multiple embedded types have conflicting names.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `{"name":"Alice"}`. The `json` struct tags control how the struct is marshaled to JSON. The `omitempty` option omits the field if it's empty, so `Email` is excluded. The `-` tag always omits the field, so `Age` is never included. Struct tags are a powerful Go feature for metadata that drives serialization and validation.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `7`. Anonymous structs let you define a struct type inline without giving it a name. They're useful for temporary data structures or when you don't need to reuse the type elsewhere. Here we create an anonymous struct with X and Y fields, initialize it, and immediately use it.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true`. In Go, structs can be compared with `==` if all their fields are comparable types. Since both `Point` structs have the same values for X and Y, they are equal. This is useful for testing and when you need to check if two structs represent the same data.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. Structs are only comparable if all their fields are comparable. Since this struct contains a slice field, and slices are not comparable in Go, the entire struct becomes non-comparable. To compare such structs, you'd need to write a custom comparison function or use reflection.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `false` then `<nil>`. This is the classic interface nil trap in Go. The function returns a nil pointer of type `*MyError`, but when assigned to an `error` interface, the interface is not nil because it has a concrete type. An interface is only nil if both its type and value are nil. This is why checking `err == nil` can be misleading.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true`. When you return `nil` directly as an `error` type, Go creates a truly nil interface where both the type and value are nil. This is the correct way to return no error. The previous example showed the trap where returning a nil concrete type creates a non-nil interface.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `write failed: disk full` then `true`. Using `%w` in `fmt.Errorf` wraps the error, creating an error chain. `errors.Is` traverses this chain to check if the target error is anywhere in the chain. This is the modern way to handle wrapped errors in Go, introduced in Go 1.13.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true`. We define a sentinel error `ErrNotFound` and wrap it with `fmt.Errorf` using `%w`. `errors.Is` can find the sentinel error even through multiple layers of wrapping. This pattern is great for checking specific error types while still providing context in the error message.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `field: email`. `errors.As` unwraps the error chain and extracts the first error that matches the target type. Unlike `errors.Is` which checks for equality, `errors.As` gives you the actual error value of a specific type. This is useful when you need to access the specific fields or methods of a particular error type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `false` then `true`. The direct equality check `==` compares the outer wrapped error, which is not equal to the original `ErrTimeout`. But `errors.Is` unwraps the error chain and finds the original error inside. This shows why you should use `errors.Is` instead of `==` when working with wrapped errors.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does the caller get nil?
**Your Response:** No, the caller gets `false`. This is the interface nil trap in practice. When `query` returns a nil `*DBError`, and we assign it to an `error` interface, the interface is not nil because it has a concrete type. The fix is to either return `nil` directly when there's no error, or change the function signature to return `error` instead of a concrete error type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true`. By implementing the `Is(error) bool` method on our error type, we can customize how `errors.Is` compares errors. Here we compare by the `Code` field value rather than pointer identity. This is useful when you want errors with the same semantic meaning to be considered equal even if they're different instances.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true`. We create a chain of three wrapped errors, and `errors.Is` successfully finds the original `ErrRoot` at the bottom of the chain. `errors.Is` recursively unwraps errors until it either finds a match or reaches the end of the chain. This makes it robust against multiple levels of error wrapping.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The blank identifier assignment `var _ Writer = NullWriter{}` is a compile-time check that verifies `NullWriter` implements the `Writer` interface. Since `NullWriter` doesn't have a `Write` method, the compilation fails. This is a useful pattern to catch interface compliance errors early.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `hello true` then `0 false`. The comma-ok form of type assertion (`value, ok := i.(type)`) safely checks if an interface holds a specific type. If it does, `ok` is true and you get the value. If not, `ok` is false and you get the zero value. This never panics, unlike the single-value form which panics on mismatch.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `recovered: interface conversion: interface {} is string, not int`. When you use a type assertion without the comma-ok form (`i.(int)`), it panics if the interface doesn't contain that type. Here the interface holds a string, not an int, so it panics. The recover catches the panic and prints the error message.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Do these mean the same thing?
**Your Response:** Yes, they print the same values. `any` is just a type alias for `interface{}` introduced in Go 1.18 to make code more readable when working with generic types. They are completely interchangeable - you can use either one and they mean exactly the same thing. `any` is preferred in generic code for clarity.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `true` then `false`. The `comparable` constraint restricts the generic type to only types that can be compared with `==` and `!=`. This includes basic types like int, string, and structs with comparable fields. It excludes slices, maps, and functions which cannot be compared.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `15` then `3.3000000000000003`. We define a custom type constraint `Number` that allows either int, int64, or float64. The generic `Sum` function works with any of these types. This shows how you can create type constraints that include multiple specific types rather than using `any`.
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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles. We have a generic interface `Setter` with a method `Set`. The `Box` type has a pointer receiver method `Set`, so `*Box[int]` satisfies the interface. The `Fill` function takes any type that implements `Setter[T]` and calls the Set method on it. This shows how generics work with interfaces and method receivers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. Go doesn't allow type switches on generic type parameters. The compiler needs to know the type at compile time for generics, but a type switch is runtime behavior. If you need runtime type checking with generics, you'd need to use reflection or redesign your approach to avoid the type switch.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `c` then `2`. We define a generic `Stack` type that can hold any type. The methods `Push`, `Pop`, and `Len` all use the type parameter `T`. We create a stack of strings, push three items, pop one, and check the length. Generics let us write type-safe data structures without using interfaces.
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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `10`. The tilde `~` in a type constraint means "any type with this underlying type". So `~int` includes not just `int` itself, but also any defined types whose underlying type is `int`, like `MyInt`. Without the tilde, it would only match exactly `int`, not `MyInt`.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `24`. The `Reduce` function is a generic reduce operation that takes a slice of type `T`, an accumulator of type `Acc`, and a function that combines the accumulator with each element. We use it to calculate the product of all numbers in the slice, starting with 1 as the initial value.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does the call compile without explicit type arguments?
**Your Response:** Yes, this compiles and prints `1 hello`. Go's type inference automatically deduces that `T` is `int` and `U` is `string` from the arguments we pass. You don't need to explicitly write `Pair[int, string]` - the compiler figures it out. This makes generic functions feel natural to use.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `2`. The `Keys` function extracts all keys from a map into a slice. The key type `K` is constrained to `comparable` because map keys must be comparable. The order of keys is not guaranteed since map iteration order is randomized in Go, but the count will be 2.

---

### 50. Generic Function Cannot Have Methods
**Q: Does this compile?**
```go
package main

func Map[T any](s []T, f func(T) T) []T { return nil }

func (Map[T]) String() string { return "" }
```
**A:** **Compile Error.** Functions (as opposed to types) cannot have methods in Go. Only generic *types* can have methods.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. In Go, only named types can have methods, not functions. Even though `Map` is generic, it's still a function, not a type. If you want methods, you need to define a generic type like `type Mapper[T any] struct{...}` and then add methods to that type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1` then `5`. Go 1.21 introduced built-in generic functions `min` and `max` that work on any ordered type (ints, floats, strings). They can take multiple arguments and return the minimum or maximum value. These are the first generic built-ins added to Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `map[]` then `[0 0 0]`. The `clear` builtin, added in Go 1.21, works on both maps and slices. For maps, it removes all entries, leaving an empty map. For slices, it sets all elements to their zero value but keeps the same length and capacity. This is useful for resetting data structures without reallocating.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints either `ch1: one` or `ch2: two` - it's non-deterministic. Both channels have values ready, so `select` randomly picks one. This randomness is intentional to prevent bias towards any particular case. If you need deterministic behavior, you should use a different pattern like prioritized selects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `no value ready`. The `default` case in a `select` statement fires immediately when no other channel operations are ready. This creates a non-blocking operation - instead of waiting for a value, we immediately fall through to the default case. This is useful for checking if a channel has data without blocking.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern do?
**Your Response:** This prints `timeout`. `time.After(100ms)` returns a channel that will send a value after 100 milliseconds. The `select` waits for either a value from `ch` or the timeout. Since no value is ever sent on `ch`, the timeout fires first. This is the standard pattern for implementing timeouts in Go, preventing operations from blocking forever.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `sent` then `received: 42`. The first `select` tries to send 42 to the buffered channel. Since the channel has capacity 1 and is empty, the send succeeds immediately. The second `select` receives the value back. Both operations use the default case to make them non-blocking, but since the channel operations are ready, they execute instead of the default.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `1 2`. The `chan<- int` type creates a send-only channel direction. Inside the `produce` function, we can only send to the channel, not receive from it. This type safety prevents accidental reads in a producer function. The channel can still be received from in `main` where it's not restricted.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `10 20`. The `<-chan int` type creates a receive-only channel direction. Inside the `consume` function, we can only receive from the channel, not send to it. This provides type safety for consumer functions that should never write to the channel. The restriction is enforced at compile time.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `recovered: close of closed channel`. In Go, closing an already-closed channel causes a panic. This is why you need to be careful when closing channels - either ensure you only close once, or use a sync.Once to guarantee single execution. The recover catches the panic and prints the error message.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `recovered: send on closed channel`. Sending a value to a closed channel causes a panic in Go. This is a runtime check to prevent undefined behavior. Once a channel is closed, you can't send more values to it, though you can still receive any remaining values. Always make sure the sender knows when the channel is closed.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1`, `2`, then `channel closed`. When receiving from a channel with the comma-ok pattern, the second value `ok` indicates whether the channel is still open. After the channel is closed and all values are drained, further receives return the zero value (0 for int) and `false`. This is the standard way to detect when a channel is closed.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `nil channel skipped`. When a channel is `nil`, any operation on it in a `select` statement is never considered ready. The case is effectively disabled. This is useful for dynamically enabling/disabling channels in a select - setting a channel to nil removes it from consideration without changing the select structure.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this do?
**Your Response:** This demonstrates the fan-out pattern with a worker pool. We create 3 worker goroutines that all read from the same jobs channel. When we send 5 jobs and close the channel, the workers distribute the work among themselves. Each job is processed exactly once, but the order varies based on which goroutine picks up each job. This is the standard pattern for parallelizing work in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this show?
**Your Response:** This shows the cancellation pattern using a done channel. The worker loops forever, checking the done channel in a select. When main closes the done channel, all goroutines waiting on it unblock immediately. This is the idiomatic way to signal cancellation in Go - closing a channel broadcasts to all waiting goroutines at once.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `getChannel called` then `default`. All channel expressions in a select are evaluated once before checking which case is ready. Even though `getChannel()` returns a nil channel (which would block), it's still called. This is important to know - any function calls in channel operations will execute regardless of which case is selected.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This shows a counting semaphore pattern using a buffered channel. The channel has capacity 3, so only 3 goroutines can acquire the semaphore at once by sending to it. When a goroutine finishes, it receives from the channel, releasing the slot for another goroutine. This limits concurrency to 3 while allowing all 6 goroutines to complete eventually.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This has a race condition. Multiple goroutines increment the same `counter` variable without synchronization. The increment operation is not atomic - it involves reading the value, incrementing it, and writing it back. When multiple goroutines do this simultaneously, some increments get lost. You can detect this with `go run -race` which will report the data race.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1000`. We fix the race condition by using a `sync.Mutex`. The `Lock()` and `Unlock()` ensure that only one goroutine can access `counter` at a time. The increment operation is now atomic - the entire read-modify-write sequence happens while holding the lock, preventing any lost updates.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens?
**Your Response:** This deadlocks. Go's `sync.Mutex` is not reentrant - the same goroutine cannot lock it twice. When `a()` locks the mutex and calls `b()`, `b()` tries to lock the same mutex again and blocks forever, waiting for itself to release the lock. In Go, you need to design your code to avoid reentrancy or use other patterns like recursive locks.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `42` five times. `sync.RWMutex` allows multiple readers simultaneously. All 5 goroutines can call `readData()` at the same time because they use `RLock()`. A write lock (`Lock()`) would be exclusive and block all readers. RWMutex is great for read-heavy workloads where writes are infrequent.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which version is safer?
**Your Response:** Version B is safer. Using `defer mu.Unlock()` ensures the mutex is always released, even if the function panics or returns early. In Version A, if `doSomething()` panics, `Unlock()` never gets called, leaving the mutex permanently locked and causing all other goroutines to deadlock.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `setup called` exactly once. `sync.Once` ensures that a function is executed exactly one time, no matter how many goroutines call `once.Do()`. Internally it uses atomic operations to track whether the function has been called. This is perfect for expensive one-time initialization like setting up database connections or loading configuration.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `true` then `postgres://localhost/db`. Both calls to `GetConfig()` return the same pointer, proving it's a singleton. The `sync.Once` ensures the configuration is initialized only once, even if multiple goroutines call `GetConfig()` simultaneously. This is the thread-safe way to implement singletons in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 5 key-value pairs where each value is the square of its key. `sync.Map` is a concurrent-safe map optimized for scenarios with many reads and fewer writes, or when different goroutines access different keys. It uses atomic operations internally and avoids the overhead of a mutex for these specific use cases.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `first false` then `first true`. `LoadOrStore` is an atomic operation that either loads an existing value or stores a new one if it doesn't exist. The first call stores "first" and reports `loaded=false`. The second call finds the existing value and returns it with `loaded=true`. This is great for cache-like patterns.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the typical output?
**Your Response:** All 3 goroutines print "woke up". `sync.Cond` is a condition variable that lets goroutines wait for a condition. `Broadcast()` wakes up all waiting goroutines, while `Signal()` would wake only one. The goroutines wait using `cond.Wait()` which atomically releases the mutex and blocks, then re-acquires it when woken.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** This has a race condition. `wg.Add(1)` is called inside the goroutine, but `wg.Wait()` might execute before any goroutine has had a chance to add itself to the WaitGroup. This can cause `Wait()` to return immediately while goroutines are still running. Always call `Add()` before starting the goroutine to avoid this race.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1000`. `atomic.AddInt64` performs atomic integer operations without locks. For simple operations like incrementing a counter, atomic operations are faster than mutexes because they use special CPU instructions. They're part of Go's `sync/atomic` package and are perfect for high-performance counters and flags.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the usage pattern?
**Your Response:** This prints `buf reused`. `sync.Pool` is a pool of reusable objects that reduces garbage collection pressure. When you need an object, you `Get()` from the pool; when done, you `Put()` it back. Objects in the pool may be garbage collected if unused for a while. This is great for frequently allocated short-lived objects like buffers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is this valid?
**Your Response:** Yes, this is valid in Go. Unlike some languages, Go's `sync.Mutex` doesn't track which goroutine locked it - any goroutine can unlock it. This shows a goroutine unlocking a mutex that was locked by main. While legal, this pattern is unusual and can be confusing. It's generally better to have the same goroutine lock and unlock the mutex.

---
