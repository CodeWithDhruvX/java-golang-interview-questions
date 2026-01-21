
## 🟤 118–147: Methods, Pointers, Interfaces, Channels & Functions

### Question 118: What is the difference between passing by value and passing by pointer in Go?

**Answer:**
- **Pass by Value:** A copy of the argument is created. Modifications inside the function do not affect the original variable.
- **Pass by Pointer:** The memory address is passed. Modifications inside the function affect the original variable.
*Note:* Everything in Go is technically passed by value (pointers are values containing addresses).

---

### Question 119: How do you define a method on a struct type?

**Answer:**
Specify a receiver between the `func` keyword and the function name.
```go
type User struct { Name string }
func (u User) GetName() string { return u.Name }
```

---

### Question 120: Can you define methods on non-struct types (e.g., type MyInt int)?

**Answer:**
**Yes**, as long as the type is defined in the same package.
```go
type MyInt int
func (m MyInt) IsPositive() bool { return m > 0 }
```

---

### Question 121: What is the difference between a Value Receiver and a Pointer Receiver?

**Answer:**
- **Value Receiver `(u User)`:** Operates on a copy. Changes are discarded.
- **Pointer Receiver `(u *User)`:** Operates on the original. Changes are persisted.

---

### Question 122: When should you use a Pointer Receiver?

**Answer:**
1. If the method needs to modify the receiver.
2. If the struct is large (to avoid copying).
3. For consistency (if some methods use pointers, use them for all).

---

### Question 123: Can you call a pointer receiver method on a value variable?

**Answer:**
**Yes.** Go automatically takes the address (`&v`) if `v` is addressable.
```go
u := User{}
u.SetAge(10) // Interpreted as (&u).SetAge(10)
```

---

### Question 124: Can you call a value receiver method on a pointer variable?

**Answer:**
**Yes.** Go automatically dereferences (`*p`) the pointer.
```go
u := &User{}
u.GetName() // Interpreted as (*u).GetName()
```

---

### Question 125: What is a "Method Set" in Go?

**Answer:**
The set of methods attached to a type.
- **Method Set of `T`:** All methods with receiver `T`.
- **Method Set of `*T`:** All methods with receiver `T` **AND** `*T`.

---

### Question 126: What methods belong to the method set of type T vs *T?

**Answer:**
See previous answer. This is crucial for satisfying interfaces. If an interface requires a pointer-receiver method, only `*T` satisfies it, not `T`.

---

### Question 127: How does new() differ from make() exactly?

**Answer:**
- **`new(T)`:** allocates memory, zeros it, and returns `*T`. Used for structs, ints, etc.
- **`make(T, ...)`:** allocates and **initializes** internal structures (like slice headers or map buckets) and returns `T` (not a pointer). Only used for slices, maps, and channels.

---

### Question 128: What types can be created using make()?

**Answer:**
1. Slices
2. Maps
3. Channels

---

### Question 129: What is the return value of new(T)?

**Answer:**
A pointer `*T` to a zeroed value of type `T`.

---

### Question 130: How are interfaces represented in memory (itab and data)?

**Answer:**
An interface value is a pair of pointers: `(type, value)`.
1. **itab (Interface Table):** Points to type information and method implementations.
2. **data:** Points to the actual concrete data held by the interface.

---

### Question 131: What is a Type Assertion and how is it used?

**Answer:**
It retrieves the concrete value from an interface.
```go
val, ok := i.(string)
if ok {
    fmt.Println(val)
}
```

---

### Question 132: What is a Type Switch?

**Answer:**
A switch construct to handle multiple possible types stored in an interface.
```go
switch v := i.(type) {
case int: ...
case string: ...
default: ...
}
```

---

### Question 133: How do you check if an interface value is nil?

**Answer:**
Compare it to `nil`.
```go
if i == nil { ... }
```
*Note:* It is `nil` only if both type and value are `nil`.

---

### Question 134: Can an interface holding a nil concrete pointer be nil?

**Answer:**
**No.** This is a common "gotcha".
```go
var p *int = nil
var i interface{} = p
// i is NOT nil because it holds (type=*int, value=nil)
```

---

### Question 135: What are the methods required to implement sort.Interface?

**Answer:**
1. `Len() int`
2. `Less(i, j int) bool`
3. `Swap(i, j int)`

---

### Question 136: How do you get the capacity (cap) and length (len) of a channel?

**Answer:**
Use the built-in `cap(ch)` and `len(ch)` functions.

---

### Question 137: What happens if you send to a closed channel?

**Answer:**
It **panics**.

---

### Question 138: What happens if you receive from a closed channel?

**Answer:**
It returns the **zero value** immediately (and returns `false` for the comma-ok check). It does not panic.

---

### Question 139: How do you check if a channel is closed ensuring no panic?

**Answer:**
Use the second return value when receiving.
```go
val, ok := <-ch
if !ok {
    // Channel is closed
}
```
*Note:* You cannot check availability "before" receiving without potentially altering the state.

---

### Question 140: What is the zero value of a function type?

**Answer:**
It is `nil`.

---

### Question 141: Can functions be used as map keys?

**Answer:**
**No.** Functions are not comparable.

---

### Question 142: How do anonymous functions (closures) capture variables?

**Answer:**
They capture variables **by reference**. Modifications in the closure affect the outer scope variable.

---

### Question 143: What is a variadic function and how do you pass a slice to it?

**Answer:**
A function that accepts a variable number of arguments (`...T`).
To pass a slice `s`, append `...`:
```go
func sum(nums ...int) { ... }
s := []int{1, 2}
sum(s...)
```

---

### Question 144: How does defer work with method evaluation (arguments vs execution)?

**Answer:**
- **Arguments** are evaluated **immediately** when `defer` is executed.
- **Function execution** happens when the surrounding function returns.
```go
defer fmt.Println(i) // i is captured NOW
i++
```

---

### Question 145: What is unsafe.Pointer and when is it used?

**Answer:**
It is a pointer type that bypasses Go's type safety.
- Allows conversion between arbitrary pointer types.
- Used for low-level memory manipulation, C interop, or atomic operations on pointers.

---

### Question 146: How does uintptr differ from *int?

**Answer:**
- `*int`: A safe pointer managed by the GC.
- `uintptr`: Just an integer that happens to hold a memory address. The GC does **not** track references held in `uintptr`, so the object it points to might be collected.

---

### Question 147: How do you manually manage memory alignment (padding)?

**Answer:**
Order struct fields from largest to smallest size to minimize padding.
```go
// Bad (24 bytes)
type Bad struct {
    A bool // 1 byte (+7 padding)
    B int64 // 8 bytes
    C bool // 1 byte (+7 padding)
}
// Good (16 bytes)
type Good struct {
    B int64
    A bool
    C bool
}
```
