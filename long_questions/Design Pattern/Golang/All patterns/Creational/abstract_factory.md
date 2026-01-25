# Abstract Factory Pattern

## 🟢 What is it?
The **Abstract Factory Pattern** is a "Super Factory" that creates other factories. It provides an interface for creating **families of related or dependent objects** without specifying their concrete structs.

Think of it like **Furniture Shopping**:
*   You want to buy a **Sofa** and a **Chair**.
*   If you choose a "Victorian" style, you want a **Victorian Sofa** AND a **Victorian Chair**.
*   If you choose a "Modern" style, you want a **Modern Sofa** AND a **Modern Chair**.
*   You don't want to mix a Victorian Sofa with a Modern Chair. The Abstract Factory ensures the "family" stays consistent.

---

## 🎯 Strategy to Implement

1.  **Define Interfaces for Products**: Define interfaces for each distinct product type (e.g., `Chair`, `Sofa`).
2.  **Create Concrete Product Families**: Implement these interfaces for each variant (e.g., `VictorianChair`, `ModernChair`).
3.  **Define Abstract Factory Interface**: Declare methods to create each product type (e.g., `CreateChair()`, `CreateSofa()`).
4.  **Create Concrete Factories**: Implement the Abstract Factory for each family variant (e.g., `VictorianFurnitureFactory`, `ModernFurnitureFactory`).
5.  **Client Code**: The client works only with the Abstract Factory and Abstract Products, unaware of the specific variants.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Abstract Products
type Chair interface {
    SitOn()
}

type Sofa interface {
    LieOn()
}

// 2. Concrete Product Family 1: Modern
type ModernChair struct{}
func (c *ModernChair) SitOn() { fmt.Println("Sitting on a sleek Modern Chair.") }

type ModernSofa struct{}
func (s *ModernSofa) LieOn() { fmt.Println("Lying on a minimalist Modern Sofa.") }

// 3. Concrete Product Family 2: Victorian
type VictorianChair struct{}
func (c *VictorianChair) SitOn() { fmt.Println("Sitting on a fancy Victorian Chair.") }

type VictorianSofa struct{}
func (s *VictorianSofa) LieOn() { fmt.Println("Lying on a velvet Victorian Sofa.") }

// 4. Abstract Factory Interface
type FurnitureFactory interface {
    CreateChair() Chair
    CreateSofa() Sofa
}

// 5. Concrete Factories
type ModernFurnitureFactory struct{}
func (f *ModernFurnitureFactory) CreateChair() Chair { return &ModernChair{} }
func (f *ModernFurnitureFactory) CreateSofa() Sofa   { return &ModernSofa{} }

type VictorianFurnitureFactory struct{}
func (f *VictorianFurnitureFactory) CreateChair() Chair { return &VictorianChair{} }
func (f *VictorianFurnitureFactory) CreateSofa() Sofa   { return &VictorianSofa{} }

// Usage
type Application struct {
    chair Chair
    sofa  Sofa
}

func NewApplication(factory FurnitureFactory) *Application {
    return &Application{
        chair: factory.CreateChair(),
        sofa:  factory.CreateSofa(),
    }
}

func (a *Application) Paint() {
    a.chair.SitOn()
    a.sofa.LieOn()
}

func main() {
    // App configured with Modern Factory
    app := NewApplication(&ModernFurnitureFactory{})
    app.Paint() // Output: Modern Chair, Modern Sofa

    // App configured with Victorian Factory
    app2 := NewApplication(&VictorianFurnitureFactory{})
    app2.Paint() // Output: Victorian Chair, Victorian Sofa
}
```

---

## ✅ When to use?

*   **Product Families**: When your code needs to work with various families of related products, and you want to ensure the client doesn't mix them.
*   **Encapsulation**: When you want to reveal only the interfaces of the products, not their implementations.
*   **Context Switching**: When you want to switch the entire "theme" or "platform" of your application just by changing the factory injection.
