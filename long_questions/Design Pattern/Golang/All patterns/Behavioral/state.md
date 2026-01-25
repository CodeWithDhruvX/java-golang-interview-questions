# State Pattern

## 🟢 What is it?
The **State Pattern** allows an object to alter its behavior when its internal state changes.

Think of it like a **Vending Machine**.

---

## 🎯 Strategy to Implement

1.  **State Interface**: Define methods for actions.
2.  **Concrete States**: Implement action methods. Transition the Context to next state.
3.  **Context**: The object that has a reference to current `State`.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Context
type VendingMachine struct {
    hasCoin State
    noCoin  State
    
    current State
}

func NewVendingMachine() *VendingMachine {
    v := &VendingMachine{}
    v.hasCoin = &HasCoinState{v}
    v.noCoin = &NoCoinState{v}
    v.current = v.noCoin
    return v
}

func (v *VendingMachine) SetState(s State) {
    v.current = s
}
func (v *VendingMachine) InsertCoin() {
    v.current.InsertCoin()
}
func (v *VendingMachine) PressButton() {
    v.current.PressButton()
}

// 2. State Interface
type State interface {
    InsertCoin()
    PressButton()
}

// 3. Concrete States
type NoCoinState struct {
    machine *VendingMachine
}
func (s *NoCoinState) InsertCoin() {
    fmt.Println("Coin inserted.")
    s.machine.SetState(s.machine.hasCoin)
}
func (s *NoCoinState) PressButton() {
    fmt.Println("No coin inserted. Cannot dispense.")
}

type HasCoinState struct {
    machine *VendingMachine
}
func (s *HasCoinState) InsertCoin() {
    fmt.Println("Coin already inserted.")
}
func (s *HasCoinState) PressButton() {
    fmt.Println("Dispensing product...")
    s.machine.SetState(s.machine.noCoin)
}

func main() {
    machine := NewVendingMachine()

    machine.PressButton() // Fail
    machine.InsertCoin()  // Success, State changes
    machine.PressButton() // Success, State changes back
}
```

---

## ✅ When to use?

*   **Behavior depends on State**: Runtime changes.
*   **Complex Switch**: Replace large conditionals.
