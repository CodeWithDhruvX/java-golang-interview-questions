# 📗 02 — OOP Concepts in Go
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

> Go uses **Composition over Inheritance**. There are no classes, but Go achieves OOP-like behavior through structs, interfaces, and embedding.

---

## 🔑 Must-Know Topics
- Structs (defining, instantiating, tags)
- Methods (value vs pointer receivers)
- Interfaces (implicit satisfaction, embedding)
- Struct embedding (composition)
- Polymorphism via interfaces
- Type assertions and type switches

---

## ❓ Most Asked Questions

### Q1. What are structs in Go? How do you define and use them?

```go
// Define a struct
type Employee struct {
    ID        int
    Name      string
    Salary    float64
    IsActive  bool
}

// Instantiate — positional
e1 := Employee{1, "Alice", 75000.0, true}

// Instantiate — named (preferred)
e2 := Employee{
    ID:       2,
    Name:     "Bob",
    Salary:   80000.0,
    IsActive: true,
}

// Access fields
fmt.Println(e2.Name)   // Bob
e2.Salary = 85000.0    // modify

// Pointer to struct
ep := &Employee{ID: 3, Name: "Charlie"}
ep.Salary = 90000.0  // auto-dereferenced
```

---

### Q2. What is the difference between value receiver and pointer receiver?

```go
type Circle struct {
    Radius float64
}

// Value receiver — works on a COPY of the struct
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Pointer receiver — works on the ORIGINAL struct (can modify it)
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor  // modifies original
}

c := Circle{Radius: 5.0}
fmt.Println(c.Area())   // 78.53...
c.Scale(2.0)
fmt.Println(c.Radius)   // 10.0
```

| | Value Receiver | Pointer Receiver |
|--|---------------|-----------------|
| Modifies original | ❌ No | ✅ Yes |
| Copying cost | Copies entire struct | Only copies pointer |
| Nil receiver safe | Yes | Must handle nil |
| Interface method set | Value + Pointer | Pointer only |

> **Rule:** Use pointer receivers when you need to modify the struct OR when the struct is large.

---

### Q3. What are interfaces in Go?

```go
// Interface defines a set of method signatures
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

// Rectangle implicitly implements Shape
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Polymorphism
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

r := Rectangle{Width: 4, Height: 3}
printShapeInfo(r)  // Area: 12.00, Perimeter: 14.00
```

> **Key:** Go interfaces are **implicitly satisfied** — no `implements` keyword needed.

---

### Q4. What is struct embedding? How is it different from inheritance?

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

// Dog EMBEDS Animal — this is composition, not inheritance
type Dog struct {
    Animal          // embedded — fields and methods promoted
    Breed string
}

d := Dog{
    Animal: Animal{Name: "Rex"},
    Breed:  "Labrador",
}

fmt.Println(d.Name)    // promoted from Animal → "Rex"
fmt.Println(d.Speak()) // promoted from Animal → "Rex makes a sound"
fmt.Println(d.Breed)   // direct field → "Labrador"

// Override (shadowing)
func (d Dog) Speak() string {
    return d.Name + " barks"
}
```

---

### Q5. What is the empty interface (`interface{}`/`any`)?

```go
// any = interface{} (Go 1.18+)
func printAnything(v interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
}

printAnything(42)
printAnything("hello")
printAnything([]int{1, 2, 3})

// Common use: generic containers
items := []interface{}{1, "two", true, 3.14}
```

---

### Q6. What is type assertion and when is it used?

```go
var i interface{} = "hello"

// Type assertion — extract concrete type
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
} else {
    fmt.Println("Not a string")
}

// Panic version (use only when sure)
s2 := i.(string)  // panics if i is not string

// Type switch — handle multiple types
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    case bool:
        fmt.Printf("bool: %t\n", v)
    default:
        fmt.Printf("unknown: %T\n", v)
    }
}
```

---

### Q7. How do you implement polymorphism in Go?

```go
type Notifier interface {
    Notify(message string) error
}

type EmailNotifier struct{ Address string }
type SMSNotifier struct{ Phone string }

func (e EmailNotifier) Notify(msg string) error {
    fmt.Printf("Email to %s: %s\n", e.Address, msg)
    return nil
}

func (s SMSNotifier) Notify(msg string) error {
    fmt.Printf("SMS to %s: %s\n", s.Phone, msg)
    return nil
}

// Polymorphic function — works with any Notifier
func sendAlert(n Notifier, msg string) {
    if err := n.Notify(msg); err != nil {
        fmt.Println("failed:", err)
    }
}

// Usage
sendAlert(EmailNotifier{Address: "a@b.com"}, "Server down!")
sendAlert(SMSNotifier{Phone: "+1234567890"}, "Server down!")
```

---

### Q8. What are struct tags and how are they used?

```go
import "encoding/json"

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Password string `json:"-"`           // excluded from JSON
    Email    string `json:"email,omitempty"` // omit if empty
    Age      int    `json:"age,string"`  // marshal as string
}

u := User{ID: 1, Name: "Alice", Password: "secret"}
data, _ := json.Marshal(u)
// Output: {"id":1,"name":"Alice"} — password excluded, email omitted

// Unmarshal
jsonStr := `{"id":2,"name":"Bob","email":"bob@example.com"}`
var u2 User
json.Unmarshal([]byte(jsonStr), &u2)
```

Common tag packages:
| Tag | Package |
|-----|---------|
| `json:"..."` | `encoding/json` |
| `db:"..."` | `sqlx` |
| `validate:"..."` | `go-playground/validator` |
| `yaml:"..."` | `gopkg.in/yaml.v3` |

---

### Q9. Can a struct implement multiple interfaces?

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(s string)
}

type ReadWriter interface {  // embedded interface
    Reader
    Writer
}

type File struct {
    content string
}

func (f *File) Read() string      { return f.content }
func (f *File) Write(s string)    { f.content = s }

// File implements Reader, Writer, AND ReadWriter
var r Reader   = &File{}
var w Writer   = &File{}
var rw ReadWriter = &File{}
```

---

### Q10. How do you handle nil interfaces?

```go
type Animal interface {
    Sound() string
}

type Dog struct{}
func (d *Dog) Sound() string { return "woof" }

var a Animal          // nil interface — both type and value are nil
fmt.Println(a == nil) // true

var d *Dog = nil
a = d                 // interface has type *Dog but nil value
fmt.Println(a == nil) // FALSE! — interface is not nil (has type info)

// Safe pattern
func makeSound(a Animal) {
    if a == nil {
        fmt.Println("no animal")
        return
    }
    fmt.Println(a.Sound())
}
```

---

### Q11. What is the difference between struct comparison and deep equality?

```go
type Point struct{ X, Y int }

p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2)  // true — structs comparable if all fields are comparable

// Structs with slices/maps are NOT directly comparable
type Data struct {
    Values []int
}
// d1 == d2  // ❌ compile error

// Use reflect.DeepEqual for deep comparison
import "reflect"
d1 := Data{Values: []int{1, 2, 3}}
d2 := Data{Values: []int{1, 2, 3}}
fmt.Println(reflect.DeepEqual(d1, d2))  // true
```

---

### Q12. What is method sets? What methods can be called on a value vs pointer?

```go
type T struct{}

func (t T)  ValueMethod()   {}  // part of T and *T method set
func (t *T) PointerMethod() {}  // part of *T method set only

var v T  = T{}
var p *T = &T{}

v.ValueMethod()    // ✅
v.PointerMethod()  // ✅ — Go auto-takes address: (&v).PointerMethod()
p.ValueMethod()    // ✅ — Go auto-dereferences: (*p).ValueMethod()
p.PointerMethod()  // ✅

// Interface assignment rules
type I interface { PointerMethod() }
var i I = p   // ✅ *T implements I
// var i I = v  // ❌ T does NOT implement I (PointerMethod not in T's method set)
```
