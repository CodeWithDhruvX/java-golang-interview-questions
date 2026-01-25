# Bridge Pattern

## 🟢 What is it?
The **Bridge Pattern** lets you split a large class or a set of closely related classes into two separate hierarchies—**Abstraction** and **Implementation**—which can be developed independently of each other.

Think of it like **Remote Controls and TVs**:
*   **Abstraction**: The Remote Control (Features: on, off, mute).
*   **Implementation**: The TV (Features: receiveSignal, powerInternal).
*   The Remote (Abstraction) holds a reference to the TV (Implementation) and bridges the user command ("Mute") to the specific TV logic.

---

## 🎯 Strategy to Implement

1.  **Implementor Interface**: Define the interface for the implementation (e.g., `Device`).
2.  **Concrete Implementors**: Create concrete structs (e.g., `Radio`, `TV`) implementing the interface.
3.  **Abstraction Struct**: Define the high-level control struct (e.g., `Remote`) that holds a reference to an object of type Implementor.
4.  **Refined Abstraction**: Extend the abstraction (by embedding or composition) to include more specific features (e.g., `AdvancedRemote`).

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Implementor (The device itself)
type Device interface {
    TurnOn()
    TurnOff()
    SetChannel(channel int)
}

// 2. Concrete Implementors
type TV struct{}
func (t *TV) TurnOn() { fmt.Println("TV: ON") }
func (t *TV) TurnOff() { fmt.Println("TV: OFF") }
func (t *TV) SetChannel(channel int) { fmt.Println("TV: Channel", channel) }

type Radio struct{}
func (r *Radio) TurnOn() { fmt.Println("Radio: ON") }
func (r *Radio) TurnOff() { fmt.Println("Radio: OFF") }
func (r *Radio) SetChannel(channel int) { fmt.Println("Radio: Frequency", channel) }

// 3. Abstraction (The Remote)
type RemoteControl struct {
    device Device // The "Bridge"
}

func (r *RemoteControl) TogglePower() {
    fmt.Println("Remote: Power button pressed.")
    r.device.TurnOn()
}

// 4. Refined Abstraction
// Go uses embedding for "inheritance", but for simple extension, we can just make a new struct
// or embed RemoteControl if we want to reuse its methods directly.
type AdvancedRemoteControl struct {
    RemoteControl // Embed basic remote
}

func (r *AdvancedRemoteControl) Mute() {
    fmt.Println("Remote: Mute button pressed.")
    r.device.SetChannel(0)
}

func main() {
    tv := &TV{}
    remote := &RemoteControl{device: tv}
    remote.TogglePower() // Works on TV

    radio := &Radio{}
    radioRemote := &AdvancedRemoteControl{
        RemoteControl: RemoteControl{device: radio},
    }
    radioRemote.TogglePower() // Works on Radio (inherited method)
    radioRemote.Mute()        // Works on Radio (new method)
}
```

---

## ✅ When to use?

*   **Avoid Cartwright Explosion**: When you want to divide and organize a monolithic class that has several variants of some functionality (e.g., Platform x Feature).
*   **Run-time Switching**: When you need to be able to switch implementations at runtime.
