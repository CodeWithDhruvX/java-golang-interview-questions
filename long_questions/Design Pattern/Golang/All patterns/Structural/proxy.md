# Proxy Pattern

## 🟢 What is it?
The **Proxy Pattern** lets you provide a substitute or placeholder for another object to control access to it.

Think of it like a **Credit Card**:
*   The Credit Card is a proxy for a **Bundle of Cash**.
*   It allows you to make payments but eventually the money comes from the Bank Account.

---

## 🎯 Strategy to Implement

1.  **Service Interface**: Define a common interface for both the Real Service and the Proxy.
2.  **Real Service**: Create the struct that does the actual work.
3.  **Proxy Struct**: Create a struct implementing the same interface. Store a reference to the Real Service.
4.  **Lazy Loading**: The proxy usually creates the Real Service only when needed.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Service Interface
type Image interface {
    Display()
}

// 2. Real Service (Heavy resource)
type RealImage struct {
    FileName string
}

func NewRealImage(fileName string) *RealImage {
    fmt.Println("Loading " + fileName + " from disk...")
    return &RealImage{FileName: fileName}
}

func (r *RealImage) Display() {
    fmt.Println("Displaying " + r.FileName)
}

// 3. Proxy Struct
type ProxyImage struct {
    Real     *RealImage
    FileName string
}

func (p *ProxyImage) Display() {
    // Lazy Loading: Create RealImage only when Display() is called
    if p.Real == nil {
        p.Real = NewRealImage(p.FileName)
    }
    p.Real.Display()
}

func main() {
    image := &ProxyImage{FileName: "test_10mb.jpg"}

    // Image is NOT loaded from disk yet.
    fmt.Println("Image object created but not loaded.")

    // Image will be loaded from disk now
    image.Display()

    // Image is already loaded, just displayed
    image.Display()
}
```

---

## ✅ When to use?

*   **Virtual Proxy (Lazy Loading)**: Heavy objects on demand.
*   **Protection Proxy**: Access control.
*   **Remote Proxy**: Local representative for remote object.
