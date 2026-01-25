# 🎨 Design Pattern Interview Questions (Questions 1-50) - Golang Edition

---

## 🏗️ Creational Patterns (Questions 1-15)

### Question 1: What is the Singleton Pattern?

**Answer:**
The **Singleton Pattern** ensures that a type has only one instance and provides a global point of access to it. It is commonly used for logging, driver objects, caching, and thread pools.

**Key Concepts:**
- **Single Instance:** Restricts struct instantiation to one object.
- **Global Access:** Provides a package-level function to get the instance.
- **Private Access:** Lowercase struct name prevents direct instantiation from other packages.

**Code:**
```go
package singleton

import "sync"

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

**Use Case:**
Managing a connection to a database or a configuration manager where multiple instances would cause inconsistency.

---

### Question 2: How do you make Singleton thread-safe in Go?

**Answer:**
In Go, `sync.Once` is the idiomatic way to ensure an action is performed exactly once, even in the presence of concurrent calls. This handles the double-checked locking mechanisms used in other languages automatically.

**Key Concepts:**
- **sync.Once:** Ensures atomic execution of initialization.

**Code:**
```go
// See Question 1 for the sync.Once implementation.
// It is inherently thread-safe.
```

**Use Case:**
High-concurrency applications accessing a shared resource.

---

### Question 3: What is the Factory Method Pattern?

**Answer:**
The **Factory Method Pattern** defines an interface for creating an object but lets subclasses (or implementations in Go) alter the type of objects that will be created.

**Key Concepts:**
- **Decoupling:** Client code works with interfaces.
- **Polymorphism:** Returns an interface type.

**Code:**
```go
type Transport interface {
    Deliver()
}

type Truck struct{}
func (t *Truck) Deliver() { println("Delivering by land") }

type Logistics interface {
    CreateTransport() Transport
}

type RoadLogistics struct{}
func (r *RoadLogistics) CreateTransport() Transport {
    return &Truck{}
}
```

**Use Case:**
A framework needs to standardize the architectural model for a range of applications, but allow for individual applications to define their own domain objects.

---

### Question 4: What is the difference between Factory and Abstract Factory?

**Answer:**
Both create objects, but they differ in scope and complexity.

**Key Differences:**
- **Factory Method:** Creates **one** product.
- **Abstract Factory:** Creates **families** of related products (e.g., Sofa + Chair).

**Code (Concept):**
```go
// Factory Method
t := logistics.CreateTransport()

// Abstract Factory
type FurnitureFactory interface {
    CreateChair() Chair
    CreateSofa() Sofa
}

// In usage
var f FurnitureFactory = &ModernFurnitureFactory{}
c := f.CreateChair()
s := f.CreateSofa() // Guaranteed to match style
```

**Use Case:**
Use **Factory** for a single object type. Use **Abstract Factory** when you need to ensure a set of objects work together.

---

### Question 5: What is the Builder Pattern?

**Answer:**
The **Builder Pattern** separates the construction of a complex object from its representation. It allows you to create different representations using the same construction code.

**Key Concepts:**
- **Step-by-Step Construction:** methods like `SetPartA()`.
- **Fluent Interface:** Method chaining (`.Name().Age().Build()`).

**Code:**
```go
type User struct {
    Name string
    Age  int
}

type UserBuilder struct {
    u User
}

func NewUserBuilder() *UserBuilder { return &UserBuilder{} }

func (b *UserBuilder) Name(n string) *UserBuilder {
    b.u.Name = n
    return b
}

func (b *UserBuilder) Age(a int) *UserBuilder {
    b.u.Age = a
    return b
}

func (b *UserBuilder) Build() User {
    return b.u
}

// Usage
user := NewUserBuilder().Name("John").Age(30).Build()
```

**Use Case:**
Constructing complex objects with many optional parameters.

---

### Question 6: What is the Prototype Pattern?

**Answer:**
The **Prototype Pattern** is used to create a duplicate object or clone of the current object.

**Key Concepts:**
- **Cloning:** In Go, simple assignments do a shallow copy. For deep copies, you might need a dedicated method.

**Code:**
```go
type Circle struct {
    Radius int
}

func (c *Circle) Clone() *Circle {
    newC := *c // Copy value
    return &newC
}
```

**Use Case:**
When object creation is expensive, create one instance and clone it.

---

### Question 7: What is the difference between Shallow Copy and Deep Copy?

**Answer:**
- **Shallow Copy:** Copies the struct fields. If a field is a pointer/slice/map, it copies the *reference*, not the underlying data.
- **Deep Copy:** Recursively copies all objects referenced.

**Use Case:**
Use Deep Copy if the prototype has mutable fields (Slices, Maps) that shouldn't be shared.

---

### Question 8: Can Enum be used as Singleton?

**Answer:**
Go does not have standard `enum` types like Java. We use `const` blocks. Since Go doesn't support Enum Singletons in the Java sense, we rely on package-level variables or `sync.Once`.

**Code:**
```go
// Idiomatic Go "Enum"
type ConfigLevel int
const (
    Debug ConfigLevel = iota
    Info
)
```

---

### Question 9: What is the Telescoping Constructor anti-pattern?

**Answer:**
It occurs when you have functions like `NewUser(name)`, `NewUserWithAge(name, age)`, `NewUserFull(name, age, phone)`.

**Problem:**
Hard to read and maintain.

**Solution:**
Use the **Builder Pattern** or **Functional Options Pattern** in Go.

---

### Question 10: When would you use Static Factory Method over Constructor?

**Answer:**
Go doesn't have constructors. We use `New...` functions.

**Benefits of Factory Functions:**
- **Descriptive Names:** `NewComplexFromPolar(r, theta)` vs `NewComplex(real, imag)`.
- **Logic:** Can return errors or cached instances.

---

### Question 11: What is the Object Pool Pattern?

**Answer:**
A pattern where a set of initialized objects is kept ready to use used (e.g., `sync.Pool`).

**Key Concepts:**
- **Get:** Retrieve generic object.
- **Put:** Return object to pool.

**Use Case:**
Reducing GC pressure by reusing objects.

---

### Question 12: What is the difference between Eager and Lazy Loading?

**Answer:**
- **Eager:** `var Global = NewThing()` (Package init).
- **Lazy:** `if instance == nil { ... }` (Inside a function).

---

### Question 13: What is "Dependency Injection"?

**Answer:**
Providing dependencies to a struct rather than letting the struct create them.

**Code:**
```go
type Car struct {
    Engine Engine
}

// With DI
func NewCar(e Engine) *Car {
    return &Car{Engine: e}
}
```

---

### Question 14: What is the Monostate Pattern?

**Answer:**
Rare in Go. Using a package-level variable effectively achieves this if multiple structs access the same global state, but it is generally discouraged in favor of passing dependencies.

---

### Question 15: What is the Registry of Singletons?

**Answer:**
Using a `map[string]interface{}` to store instances.

**Code:**
```go
var registry = make(map[string]interface{})
registry["Logger"] = NewLogger()
```

---

## 🧱 Structural Patterns (Questions 16-30)

### Question 16: What is the Adapter Pattern?

**Answer:**
Allows incompatible interfaces to work together.

**Code:**
```go
type JsonReader interface { ReadJson() }
type XmlService struct{} // Incompatible
func (x *XmlService) ReadXml() {}

// Adapter
type XmlAdapter struct {
   Service *XmlService
}
func (a *XmlAdapter) ReadJson() {
   a.Service.ReadXml() // Convert call
}
```

---

### Question 17: What is the Decorator Pattern?

**Answer:**
Wraps an object to add behavior dynamically.

**Code:**
```go
type Coffee interface { Cost() int }
type SimpleCoffee struct{}
func (s *SimpleCoffee) Cost() int { return 10 }

type MilkShim struct {
    Parent Coffee
}
func (m *MilkShim) Cost() int { return m.Parent.Cost() + 5 }

// Usage
var c Coffee = &SimpleCoffee{}
c = &MilkShim{Parent: c}
```

---

### Question 18: What is the Facade Pattern?

**Answer:**
Provides a simplified interface to a library.

**Code:**
```go
func (ht *HomeTheater) WatchMovie() {
    ht.lights.Dim()
    ht.tv.On()
}
```

---

### Question 19: Adapter vs Facade?

**Answer:**
- **Adapter:** Fits square peg in round hole (interoperability).
- **Facade:** Simplifies a complex dashboard (usability).

---

### Question 20: What is the Proxy Pattern?

**Answer:**
Provides a placeholder to control access.

**Code:**
```go
type Image interface { Display() }
type RealImage struct{ file string }
func (r *RealImage) Display() { /* load and show */ }

type ProxyImage struct {
    Real *RealImage
    File string
}
func (p *ProxyImage) Display() {
    if p.Real == nil { p.Real = &RealImage{p.File} } // Lazy load
    p.Real.Display()
}
```

---

### Question 21: Proxy vs Decorator?

**Answer:**
- **Proxy:** Controls access (lazy loading, security).
- **Decorator:** Adds behavior (logging, formatting).

---

### Question 22: What is the Bridge Pattern?

**Answer:**
Decouples implementation from abstraction.

**Key Concepts:**
- **Composition:** Embed interface in struct.

**Code:**
```go
type Device interface { TurnOn() }
type Remote struct {
    dev Device
}
func (r *Remote) Toggle() { r.dev.TurnOn() }
```

---

### Question 23: What is the Composite Pattern?

**Answer:**
Tree structures where individual objects and groups are treated same.

**Code:**
```go
type Component interface { Render() }

type File struct{}
type Folder struct {
    Files []Component
}
func (f *Folder) Render() {
    for _, c := range f.Files { c.Render() }
}
```

---

### Question 24: What is the Flyweight Pattern?

**Answer:**
Sharing common state to support large numbers of objects.

**Concepts:**
- Store invariant state (color, texture) in a shared map/struct.

---

### Question 25: What is the Private Class Data Pattern?

**Answer:**
In Go, using unexported fields in a struct prevents external modification, achieving the same goal naturally.

---

### Question 26: What is the Marker Interface Pattern?

**Answer:**
An empty interface.

**Code:**
```go
type Serializable interface{} // Empty
```
Go often uses struct tags or runtime checks, but empty interfaces like `interface{}` are common for generic handling.

---

### Question 27: What is the Transfer Object Pattern (DTO)?

**Answer:**
Simple structs to transfer data.

**Code:**
```go
type UserDTO struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

---

### Question 28: What is the DAO Pattern?

**Answer:**
Abstracts data access.

**Code:**
```go
type UserDao interface {
    Get(id int) User
}
```

---

### Question 29: What is the Front Controller Pattern?

**Answer:**
Central request handler (e.g., `http.Handler` wrapper or middleware).

---

### Question 30: What is the Intercepting Filter Pattern?

**Answer:**
Middleware in Go (func taking `http.Handler` and returning `http.Handler`).

---

## 🧠 Behavioral Patterns (Questions 31-50)

### Question 31: What is the Strategy Pattern?

**Answer:**
Interchangeable algorithms.

**Code:**
```go
type PaymentStrategy interface { Pay(amount int) }

type Context struct {
    strategy PaymentStrategy
}
func (c *Context) Pay(amt int) { c.strategy.Pay(amt) }
```

---

### Question 32: What is the Observer Pattern?

**Answer:**
Notification system.

**Code:**
```go
type Observer interface { Update(msg string) }
type Subject struct {
    observers []Observer
}
func (s *Subject) Notify(msg string) {
    for _, o := range s.observers { o.Update(msg) }
}
```

---

### Question 33: What is the Command Pattern?

**Answer:**
Encapsulates request as object/function.

**Code:**
```go
type Command func()

func Execute(c Command) { c() }

// Usage
Execute(func() { fmt.Println("Light On") })
```

---

### Question 34: What is the Chain of Responsibility Pattern?

**Answer:**
Pass request along chain.

**Code:**
```go
type Handler interface {
    Handle(req string)
    SetNext(h Handler)
}
```

---

### Question 35: What is the Template Method Pattern?

**Answer:**
Defines skeleton. In Go, we use composition + interfaces or higher-order functions since we don't have inheritance based overrides of superclass methods in the same way.

**Code:**
```go
type Game interface {
    Initialize()
    StartPlay()
    EndPlay()
}

func PlayGame(g Game) {
    g.Initialize()
    g.StartPlay()
    g.EndPlay()
}
```

---

### Question 36: Template Method vs Strategy?

**Answer:**
Similarities. In Go, Strategy is often preferred due to lack of inheritance.

---

### Question 37: What is the State Pattern?

**Answer:**
Object alters behavior when state changes.

**Code:**
```go
type State interface { Handle(ctx *Context) }
type Context struct { Current State }
func (c *Context) Request() { c.Current.Handle(c) }
```

---

### Question 38: What is the Iterator Pattern?

**Answer:**
Traverse collection.

**Code:**
```go
// Go uses range natively for slice/map
for i, v := range items { ... }
```
Custom iterators can use channels or struct with `Next() bool`.

---

### Question 39: What is the Mediator Pattern?

**Answer:**
Centralizes communication. Channels can act as mediators in concurrent Go programs.

---

### Question 40: What is the Memento Pattern?

**Answer:**
Captures state.

**Code:**
```go
type Memento struct { state string }
func (o *Originator) Save() Memento { return Memento{o.state} }
```

---

### Question 41: What is the Visitor Pattern?

**Answer:**
Add operations to objects.

**Code:**
```go
type Visitor interface { VisitFile(f *File) }
type Element interface { Accept(v Visitor) }

func (f *File) Accept(v Visitor) { v.VisitFile(f) }
```

---

### Question 42: What is the Null Object Pattern?

**Answer:**
Default no-op implementation.

**Code:**
```go
type Logger interface { Log(msg string) }
type NullLogger struct{}
func (n *NullLogger) Log(msg string) {} // Do nothing
```

---

### Question 43: What is the Interpreter Pattern?

**Answer:**
Grammar evaluation.

---

### Question 44: What is the Service Locator Pattern?

**Answer:**
Registry for services. Anti-pattern; prefer DI.

---

### Question 45: What is the MVC Pattern?

**Answer:**
Standard web architecture.
- **Model:** Structs/DB.
- **View:** JSON/HTML Templates.
- **Controller:** Handlers.

---

### Question 46: What is MVVM?

**Answer:**
Model-View-ViewModel. Less common in backend Go, more for frontends.

---

### Question 47: What is CQRS?

**Answer:**
Separating Read/Write models. Go is great for this with simple structs for commands and queries.

---

### Question 48: What is Repository Pattern?

**Answer:**
Interface for data storage.

**Code:**
```go
type UserRepo interface {
    Save(u User) error
    Find(id int) (User, error)
}
```

---

### Question 49: What is Inversion of Control (IoC)?

**Answer:**
Framework calls you.

---

### Question 50: Cohesion vs Coupling?

**Answer:**
Same principles apply.
- **Cohesion:** Package focuses on one thing.
- **Coupling:** Dependencies between packages. keep them minimal (avoid cyclic imports).
