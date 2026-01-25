# SOLID Principles: Practical Guide & Interview Applications

## Overview
SOLID principles are the foundation of good Object-Oriented Design. In interviews, you are often asked to "refactor this code" or "design a system" where adhering to these principles is the key to passing.

## 1. S: Single Responsibility Principle (SRP)
**Definition**: A class should have only one reason to change.

### 🔴 Bad Example (God Object)
```go
type Order struct {}

func (o *Order) CreateOrder() { /* ... */ }
func (o *Order) CalculateTax() { /* ... */ }
func (o *Order) SendEmailConfirmation() { /* ... */ }
func (o *Order) SaveToDB() { /* ... */ }
```
*Why it fails*: If email logic changes, you edit `Order`. If DB changes, you edit `Order`.

### 🟢 Good Example (Separation of Concerns)
```go
type OrderService struct {
    repo Repository
    notifier Notifier
    taxCalc TaxCalculator
}

func (s *OrderService) CreateOrder(o Order) {
    tax := s.taxCalc.Calculate(o)
    s.repo.Save(o, tax)
    s.notifier.SendConfirmation(o)
}
```

## 2. O: Open/Closed Principle (OCP)
**Definition**: Software entities should be open for extension, but closed for modification.

### 🔴 Bad Example (Modify to Extend)
```go
type Shape struct { Type string }

func Area(s Shape) double {
    if s.Type == "Circle" { return 3.14 * r * r }
    if s.Type == "Square" { return s * s }
    // Adding Triangle requires modifying this function!
}
```

### 🟢 Good Example (Interface Extension)
```go
type Shape interface {
    Area() double
}

type Circle struct { Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

type Square struct { Side float64 }
func (s Square) Area() float64 { return s.Side * s.Side }

// Adding Triangle? Just create new struct. No change to existing code.
```

## 3. L: Liskov Substitution Principle (LSP)
**Definition**: Objects of a superclass shall be replaceable with objects of its subclasses without breaking the application.

### 🔴 Bad Example (Breaking Behavior)
```go
type Bird interface { Fly() }

type Ostrich struct {}
func (o Ostrich) Fly() {
    panic("Ostrich cannot fly!") // Breaks LSP!
}
```

### 🟢 Good Example (Interface Segregation for Features)
```go
type Bird interface { Eat() }
type FlyingBird interface { Fly() }

type Sparrow struct {} // Implements both
type Ostrich struct {} // Implements only Bird
```

## 4. I: Interface Segregation Principle (ISP)
**Definition**: Clients should not be forced to depend upon interfaces that they do not use.

### 🔴 Bad Example (Fat Interface)
```go
type Worker interface {
    Work()
    Eat()
    Sleep()
}

type Robot struct {}
func (r Robot) Work() { /* ok */ }
func (r Robot) Eat() { /* useless */ } // Robots don't eat
```

### 🟢 Good Example (Small Interfaces)
```go
type Worker interface { Work() }
type Eater interface { Eat() }

type Human struct {} // Implements both
type Robot struct {} // Implements only Worker
```

## 5. D: Dependency Inversion Principle (DIP)
**Definition**: High-level modules should not depend on low-level modules. Both should depend on abstractions.

### 🔴 Bad Example (Hard Dependency)
```go
type Service struct {
    db *MySQLDatabase // Hard dependency on MySQL
}
```

### 🟢 Good Example (Abstraction)
```go
type Database interface {
    Save(data interface{}) error
}

type Service struct {
    db Database // Can be MySQL, Postgres, or Mock for testing!
}
```
*Why it matters*: This is essential for Unit Testing. You can inject a `MockDatabase` into `Service` to test it without a real DB connection.

## Summary Table

| Principle | Key Takeaway | Interview Cue |
| :--- | :--- | :--- |
| **SRP** | One reason to change | "This class is too huge." |
| **OCP** | Extend, don't modify | "How do we add a new payment method?" |
| **LSP** | Subtypes must behave | "Why is this subclass throwing errors?" |
| **ISP** | Split large interfaces | "Why do I implement methods I don't use?" |
| **DIP** | Interface over Implementation | "How do we unit test this?" |
