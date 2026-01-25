# Decorator Pattern

## 🟢 What is it?
The **Decorator Pattern** lets you attach new behaviors to objects by placing these objects inside special wrapper objects that experience the behaviors.

Think of it like **Wearing Clothes**:
*   You (the Object) are a person.
*   You put on a T-shirt (Decorator 1).
*   You put on a Jacket (Decorator 2).
*   You can take them off or add layers in any order dynamically.

---

## 🎯 Strategy to Implement

1.  **Component Interface**: Define the common interface (e.g., `Coffee`).
2.  **Concrete Component**: Create the base object struct (e.g., `SimpleCoffee`).
3.  **Decorator**: Create structs that implement the Component Interface and hold a field for a Component object.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Component Interface
type Coffee interface {
    GetDescription() string
    GetCost() float64
}

// 2. Concrete Component
type SimpleCoffee struct{}
func (s *SimpleCoffee) GetDescription() string { return "Simple Coffee" }
func (s *SimpleCoffee) GetCost() float64       { return 5.0 }

// 3. Decorator: Milk
type Milk struct {
    Coffee Coffee // The wrapped component
}

func (m *Milk) GetDescription() string {
    return m.Coffee.GetDescription() + ", Milk"
}

func (m *Milk) GetCost() float64 {
    return m.Coffee.GetCost() + 1.5
}

// 4. Decorator: Sugar
type Sugar struct {
    Coffee Coffee
}

func (s *Sugar) GetDescription() string {
    return s.Coffee.GetDescription() + ", Sugar"
}

func (s *Sugar) GetCost() float64 {
    return s.Coffee.GetCost() + 0.5
}

func main() {
    // Order: Simple Coffee
    var myCoffee Coffee = &SimpleCoffee{}
    fmt.Printf("%s $%.2f\n", myCoffee.GetDescription(), myCoffee.GetCost())

    // Add Milk
    myCoffee = &Milk{Coffee: myCoffee}
    fmt.Printf("%s $%.2f\n", myCoffee.GetDescription(), myCoffee.GetCost())

    // Add Sugar (Wrapping the Milk-wrapped Coffee)
    myCoffee = &Sugar{Coffee: myCoffee}
    fmt.Printf("%s $%.2f\n", myCoffee.GetDescription(), myCoffee.GetCost())
}
```

---

## ✅ When to use?

*   **Dynamic Extension**: When you need to add responsibilities to individual objects dynamically.
*   **Prevent Explosion**: When extending methods by inheritance would produce an explosion of subclasses.
