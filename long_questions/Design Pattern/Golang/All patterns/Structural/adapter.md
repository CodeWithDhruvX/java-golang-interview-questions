# Adapter Pattern

## 🟢 What is it?
The **Adapter Pattern** allows objects with incompatible interfaces to collaborate. It acts as a bridge between two incompatible interfaces.

Think of it like a **Travel Adapter**:
*   You have a US laptop plug (Client).
*   You are in Europe with a European Wall Socket (Service).
*   You cannot plug the US plug into the EU socket directly.
*   You use an **Adapter**. It takes the US plug on one side and fits into the EU socket on the other.

---

## 🎯 Strategy to Implement

1.  **Target Interface**: Start with the interface that your Client code expects to use (e.g., `LightningPhone`).
2.  **Adaptee Interface/Struct**: Identify the useful struct that has an incompatible interface (e.g., `MicroUsbPhone`).
3.  **Adapter Struct**: Create a new struct that implements the Target Interface.
4.  **Wrap the Service**: Inside the Adapter struct, store a reference to the Adaptee (Composition).
5.  **Translate Calls**: In the methods of the Adapter, call the methods of the Service object, translating the data if necessary.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Target Interface (What the client expects)
type LightningPhone interface {
    Recharge()
    UseLightning()
}

// 2. Adaptee (The incompatible service we want to use)
type MicroUsbPhone interface {
    Recharge()
    UseMicroUsb()
}

type AndroidPhone struct{}
func (a *AndroidPhone) Recharge() { fmt.Println("Recharging Android...") }
func (a *AndroidPhone) UseMicroUsb() { fmt.Println("MicroUsb connected.") }

type IPhone struct{}
func (i *IPhone) Recharge() { fmt.Println("Recharging iPhone...") }
func (i *IPhone) UseLightning() { fmt.Println("Lightning connected.") }

// 3. Adapter Struct
// We want to use an Android Phone (MicroUsb) but the charger is for Lightning
type LightningToMicroUsbAdapter struct {
    MicroUsbPhone MicroUsbPhone
}

func (a *LightningToMicroUsbAdapter) UseLightning() {
    fmt.Println("Adapter converts Lightning signal to MicroUsb...")
    a.MicroUsbPhone.UseMicroUsb()
}

func (a *LightningToMicroUsbAdapter) Recharge() {
    a.MicroUsbPhone.Recharge()
}

func main() {
    android := &AndroidPhone{}

    // I want to charge my Android, but I only have a Lightning Cable object.
    // var cable LightningPhone = android // Compile Error!

    // Use Adapter
    var adapter LightningPhone = &LightningToMicroUsbAdapter{MicroUsbPhone: android}

    adapter.UseLightning()
    adapter.Recharge()
}
```

---

## ✅ When to use?

*   **Legacy Code integration**: When you want to use an existing struct, but its interface does not match the one you need.
*   **Interface translation**: When you need to convert data formats between systems (e.g., XML to JSON) transparently.
