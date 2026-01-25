# Factory Method Pattern

## 🟢 What is it?
The **Factory Method Pattern** provides a way to create objects **without specifying the exact struct** of object that will be created.

It defines an interface for creating an object but lets **implementations decide** which struct to instantiate. It pushes the "instantiation logic" from the client code to special "Factory" methods.

Think of it like a **Logistics Company**:
*   The generic process is "Deliver Package".
*   If the delivery is by land, the "Land Logistics" department creates a **Truck**.
*   If the delivery is by sea, the "Sea Logistics" department creates a **Ship**.
*   The client (you) just says "Deliver this", not "Go buy a Truck and drive it".

---

## 🎯 Strategy to Implement

1.  **Define a Product Interface**: This is the common interface for all objects your factory will create (e.g., `Transport`).
2.  **Create Concrete Products**: Implement the interface with specific structs (e.g., `Truck`, `Ship`).
3.  **Define the Creator Interface**: Declare the factory method (e.g., `CreateTransport()`) that returns a `Transport`.
4.  **Create Concrete Creators**: Implement the Creator interface to return a specific instance.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. The Product Interface
type Transport interface {
    Deliver()
}

// 2. Concrete Products
type Truck struct{}
func (t *Truck) Deliver() {
    fmt.Println("Delivering by land in a box.")
}

type Ship struct{}
func (s *Ship) Deliver() {
    fmt.Println("Delivering by sea in a container.")
}

// 3. The Creator Interface
type Logistics interface {
    CreateTransport() Transport
}

// 4. Concrete Creators
type RoadLogistics struct{}
func (r *RoadLogistics) CreateTransport() Transport {
    return &Truck{}
}

type SeaLogistics struct{}
func (s *SeaLogistics) CreateTransport() Transport {
    return &Ship{}
}

// Core business logic (Client)
func PlanDelivery(l Logistics) {
    t := l.CreateTransport()
    t.Deliver()
}

func main() {
    // Based on configuration or user input...
    deliveryType := "sea"

    var logistics Logistics

    if deliveryType == "road" {
        logistics = &RoadLogistics{}
    } else {
        logistics = &SeaLogistics{}
    }

    // The client code doesn't know if it's using a Truck or Ship.
    PlanDelivery(logistics)
}
```

---

## ✅ When to use?

*   **Unknown Dependencies**: When you don't know beforehand the exact types and dependencies of the objects your code should work with.
*   **Extensibility**: When you want to provide a library or framework that users can extend. They can create new "Product" types and a new "Creator" without breaking your existing code.
*   **Decoupling**: When you want to save system resources by reusing existing objects instead of rebuilding them each time (managing object lifecycles).
