# Prototype Pattern

## 🟢 What is it?
The **Prototype Pattern** lets you copy existing objects without making your code dependent on their classes. It delegates the cloning process to the actual objects that are being cloned.

Think of it like **Cell Division (Mitosis)**:
*   A cell doesn't call a "Cell Factory" to make another cell.
*   It splits itself to create an exact copy.
*   Or like "Ctrl+C / Ctrl+V" in a document.

---

## 🎯 Strategy to Implement

1.  **Clone Method**: In Go, assigning a struct creates a shallow copy by default. For complex objects or deep copies, implement a `Clone()` method.
2.  **Pointer vs Value**: Be careful when cloning structs with pointers, slices, or maps. A simple dereference copy `*p` will point to the same underlying arrays/maps. You must manually copy those if you want a true independent clone.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Prototype Interface (Optional in Go, but useful)
type Cloner interface {
    Clone() Cloner
}

// 2. Concrete Prototype
type Shape struct {
    X, Y  int
    Color string
}

func (s *Shape) Clone() Cloner {
    // Shallow copy is enough for simple fields
    newS := *s
    return &newS
}

type Circle struct {
    Shape  // Embedding
    Radius int
}

func (c *Circle) Clone() Cloner {
    // Copy the parent part
    newShape := c.Shape.Clone().(*Shape)
    
    // Create new Circle
    newC := &Circle{
        Shape:  *newShape,
        Radius: c.Radius,
    }
    return newC
}

func main() {
    // Create original object
    c1 := &Circle{
        Shape:  Shape{X: 10, Y: 20, Color: "Red"},
        Radius: 15,
    }

    // Clone it
    c2 := c1.Clone().(*Circle)

    // Modify the clone
    c2.Color = "Blue"

    fmt.Printf("Circle 1: %+v\n", c1) // Red
    fmt.Printf("Circle 2: %+v\n", c2) // Blue
}
```

---

## ✅ When to use?

*   **Expensive Creation**: When creating an object from scratch is resource-intensive definition (e.g., parsing a large XML file). Cloning is often much faster.
*   **Complex Setup**: When you have objects with complex configurations (many fields set), and you need similar objects. Instead of re-configuring a new one, clone the template and tweak the differences.
