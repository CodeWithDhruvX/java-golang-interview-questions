# Command Pattern

## 🟢 What is it?
The **Command Pattern** turns a request into a stand-alone object. This lets you parameterize methods, queue requests, and support undoable operations.

Think of it like a **Waiter in a Restaurant**:
*   You (Client) give an **Order** (Command) to the **Waiter**.
*   The **Chef** (Receiver) executes it.

---

## 🎯 Strategy to Implement

1.  **Command Interface**: Declare the `Execute()` method.
2.  **Concrete Commands**: Store the Receiver and params.
3.  **Receiver**: Functionality logic.
4.  **Invoker**: Triggers the command.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Receiver
type Light struct {
    IsOn bool
}

func (l *Light) TurnOn()  { fmt.Println("Light is ON"); l.IsOn = true }
func (l *Light) TurnOff() { fmt.Println("Light is OFF"); l.IsOn = false }

// 2. Command Interface
type Command interface {
    Execute()
    Undo()
}

// 3. Concrete Commands
type TurnOnLightCommand struct {
    light *Light
}

func (c *TurnOnLightCommand) Execute() { c.light.TurnOn() }
func (c *TurnOnLightCommand) Undo()    { c.light.TurnOff() }

type TurnOffLightCommand struct {
    light *Light
}

func (c *TurnOffLightCommand) Execute() { c.light.TurnOff() }
func (c *TurnOffLightCommand) Undo()    { c.light.TurnOn() }

// 4. Invoker
type RemoteControl struct {
    command Command
}

func (r *RemoteControl) SetCommand(c Command) {
    r.command = c
}

func (r *RemoteControl) PressButton() {
    r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
    if r.command != nil {
        r.command.Undo()
    }
}

func main() {
    livingRoomLight := &Light{}

    lightsOn := &TurnOnLightCommand{light: livingRoomLight}
    lightsOff := &TurnOffLightCommand{light: livingRoomLight}

    remote := &RemoteControl{}

    // Switch on
    remote.SetCommand(lightsOn)
    remote.PressButton() // Light is ON

    // Switch off
    remote.SetCommand(lightsOff)
    remote.PressButton() // Light is OFF

    // Undo last action (Turn it back on)
    remote.PressUndo() // Light is ON
}
```

---

## ✅ When to use?

*   **Undo/Redo**: Reversible operations.
*   **Queueing Tasks**: Job Queues.
*   **Decoupling**: Invoker doesn't know details of execution.
