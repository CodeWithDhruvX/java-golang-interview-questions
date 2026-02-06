# Golang OOPs Concepts Cheatsheet

Go is not a traditional Object-Oriented language (no `class` keyword), but it supports OOP features through **Structs**, **Methods**, and **Interfaces**.

---

## 🟦 Structs (The "Class" Alternative)

In Go, a `struct` replaces the class. It defines the state (data fields).

### Definition & Instantiation
```go
// Define a struct (like a class definition)
type Employee struct {
    ID   int
    Name string
    Role string
}

// Instantiate
func main() {
    // 1. Literal style
    e1 := Employee{ID: 1, Name: "Alice", Role: "Dev"}

    // 2. Key-value style (Recommended)
    e2 := Employee{
        Name: "Bob",
        Role: "Manager", // ID is 0 (default)
    }

    // 3. New keyword (Returns pointer)
    e3 := new(Employee) // *Employee
    e3.Name = "Charlie"
}
```

### Constructor Pattern
Go doesn't have constructors. We use factory functions, typically named `New<Type>`.

```go
func NewEmployee(name string, role string) *Employee {
    return &Employee{
        ID:   rand.Int(),
        Name: name,
        Role: role,
    }
}

// Usage
emp := NewEmployee("Dave", "Lead")
```

---

## 🟧 Methods (Behavior)

Methods are functions attached to a specific type (receiver).

### Value vs. Pointer Receivers
* **Value Receiver**: Operates on a *copy*. Cannot modify original state.
* **Pointer Receiver**: Operates on the *actual* object. Can modify state.

```go
type Counter struct {
    Value int
}

// Value Receiver (Safe, Read-only logic)
func (c Counter) Get() int {
    // c.Value = 100 // This would only change the local copy
    return c.Value
}

// Pointer Receiver (Mutates state)
func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{}
    c.Increment()      // Go automatically converts &c
    fmt.Println(c.Get()) // Output: 1
}
```
**Rule of Thumb:**
- Use **Pointer Receiver** if you need to modify the struct OR if the struct is large (to avoid copying).
- Use **Value Receiver** for small, immutable types.

---

## 🔒 Encapsulation

Go uses **capitalization** to control visibility (Public/Private), not keywords like `public` or `private`.

| Naming | Visibility | Scope |
| :--- | :--- | :--- |
| **`TitleCase`** | **Public** (Exported) | Accessible from other packages |
| **`camelCase`** | **Private** (Unexported) | Accessible ONLY within the same package |

```go
package users

type Account struct {
    Owner    string // Exported (Public)
    balance  int    // Unexported (Private)
}

// Setter (Public method to access private field)
func (a *Account) Deposit(amount int) {
    if amount > 0 {
        a.balance += amount
    }
}

// Getter
func (a *Account) Balance() int {
    return a.balance
}
```

---

## 🔗 Composition (Inheritance Alternative)

Go prefers **Composition over Inheritance**. It uses **Struct Embedding**.

### Type Embedding
By embedding a struct anonymously, the inner struct's fields and methods are **promoted** to the outer struct.

```go
type Animal struct {
    Species string
}

func (a *Animal) Move() {
    fmt.Println(a.Species, "is moving")
}

type Dog struct {
    Animal // Embedded Struct (Is-a relationship)
    Breed  string
}

func main() {
    d := Dog{
        Animal: Animal{Species: "Canine"},
        Breed:  "Labrador",
    }

    // Direct access to promoted fields/methods
    d.Move()           // Works! Prints "Canine is moving"
    fmt.Println(d.Species) // Works!
}
```

### Overriding Methods (Shadowing)
If `Dog` defines its own `Move()` method, it shadows the embedded `Animal.Move()`.

```go
func (d *Dog) Move() {
    fmt.Println("Dog runs fast!")
}

// d.Move() // Calls Dog's Move
// d.Animal.Move() // Explicitly calls embedded Move
```

---

## 🟣 Polymorphism (Interfaces)

Interfaces define **behavior** (method signatures). Types implement interfaces **implicitly** just by having the methods.

### Definition
```go
type Flyer interface {
    Fly(distance int)
}
```

### Implementation (Implicit)
```go
type Bird struct{ Name string }

// Implicitly implements Flyer because signature matches
func (b Bird) Fly(dist int) {
    fmt.Println(b.Name, "flew", dist, "miles")
}

type Plane struct{ Model string }

func (p Plane) Fly(dist int) {
    fmt.Println(p.Model, "flew", dist, "miles")
}
```

### Usage (Polymorphism)
```go
func Travel(f Flyer) {
    f.Fly(100)
}

func main() {
    b := Bird{"Eagle"}
    p := Plane{"Boeing 747"}

    // Both work because they implement Flyer
    Travel(b)
    Travel(p)
}
```

### Interface Composition
Interfaces can be embedded into other interfaces.

```go
type Reader interface { Read(b []byte) (n int, err error) }
type Writer interface { Write(b []byte) (n int, err error) }

// ReadWriter embeds both
type ReadWriter interface {
    Reader
    Writer
}
```

---

## ⚡ Type Assertions & Switches

When working with interfaces (especially empty interface `interface{}`), you may need to retrieve the underlying concrete type.

### Type Assertion
```go
var i interface{} = "hello"

s, ok := i.(string) // Check if 'i' is string
if ok {
    fmt.Println("It's a string:", s)
}

r, ok := i.(float64) // False, r is 0
```

### Type Switch
```go
func checkType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case Flyer:
        fmt.Println("Implements Flyer")
    default:
        fmt.Println("Unknown type")
    }
}
```
