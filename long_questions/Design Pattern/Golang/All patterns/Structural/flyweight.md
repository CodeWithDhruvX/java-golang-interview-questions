# Flyweight Pattern

## 🟢 What is it?
The **Flyweight Pattern** lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects.

Think of it like a **Text Editor**:
*   Instead of creating 1000 objects for lines containing 'A' with all font data, you create **one** 'A' object (Flyweight) and share it.

---

## 🎯 Strategy to Implement

1.  **Intrinsic State (Shared)**: Identify data that is constant across many objects (e.g., Color, Texture).
2.  **Extrinsic State (Unique)**: Identify data that changes per object (e.g., X/Y coordinates). Pass this to methods.
3.  **Flyweight Factory**: Create a factory that manages a map of existing Flyweights.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Flyweight Interface
type Tree interface {
    Draw(x, y int)
}

// 2. Concrete Flyweight (Shared State)
type TreeType struct {
    Name  string
    Color string
}

func (t *TreeType) Draw(x, y int) {
    // x and y are extrinsic (unique) state passed in
    fmt.Printf("Drawing %s tree (%s) at %d, %d\n", t.Name, t.Color, x, y)
}

// 3. Flyweight Factory
type TreeFactory struct {
    treeTypes map[string]*TreeType
}

var factory = &TreeFactory{treeTypes: make(map[string]*TreeType)}

func (f *TreeFactory) GetTreeType(name, color string) *TreeType {
    key := name + "-" + color
    if _, exists := f.treeTypes[key]; !exists {
        f.treeTypes[key] = &TreeType{Name: name, Color: color}
        fmt.Println("Creating new TreeType: " + name)
    }
    return f.treeTypes[key]
}

func main() {
    // We want to plant 10,000 trees but only create 2 actual objects in memory
    
    // Plant Oaks
    for i := 0; i < 3; i++ {
        t := factory.GetTreeType("Oak", "Green")
        t.Draw(i, i*2)
    }

    // Plant Pines
    for i := 0; i < 3; i++ {
        t := factory.GetTreeType("Pine", "DarkGreen")
        t.Draw(i, i*3)
    }
}
```

---

## ✅ When to use?

*   **Massive Quantity**: When your application needs to spawn a huge number of similar objects.
*   **Memory Constraints**: When storing all data for every object drains available RAM.
