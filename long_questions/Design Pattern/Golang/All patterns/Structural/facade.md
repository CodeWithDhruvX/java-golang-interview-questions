# Facade Pattern

## 🟢 What is it?
The **Facade Pattern** provides a simplified interface to a library, a framework, or any other complex set of classes.

Think of it like a **Home Theater Remote**:
*   A **Facade** (the "Watch Movie" button) does all the complex steps (TV on, HDMI input, Sound on, etc.) for you with one click.

---

## 🎯 Strategy to Implement

1.  **Identify Subsystems**: Find the complex parts of your system (e.g., DVDPlayer, Projector).
2.  **Create Facade Struct**: Define a struct with simple methods (e.g., `WatchMovie()`).
3.  **Delegate**: In the Facade methods, call the appropriate methods of the subsystems.

---

## 💻 Code Example

```go
package main

import "fmt"

// Subsystem 1
type DVDPlayer struct{}
func (d *DVDPlayer) On() { fmt.Println("DVD Player On") }
func (d *DVDPlayer) Play(movie string) { fmt.Println("Playing " + movie) }

// Subsystem 2
type Projector struct{}
func (p *Projector) On() { fmt.Println("Projector On") }
func (p *Projector) SetInput() { fmt.Println("Projector Input set to DVD") }

// Subsystem 3
type Lights struct{}
func (l *Lights) Dim(level int) { fmt.Println("Lights dimmed to", level, "%") }

// The Facade
type HomeTheaterFacade struct {
    dvd       *DVDPlayer
    projector *Projector
    lights    *Lights
}

func NewHomeTheaterFacade(d *DVDPlayer, p *Projector, l *Lights) *HomeTheaterFacade {
    return &HomeTheaterFacade{dvd: d, projector: p, lights: l}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
    fmt.Println("Get ready to watch a movie...")
    h.lights.Dim(10)
    h.projector.On()
    h.projector.SetInput()
    h.dvd.On()
    h.dvd.Play(movie)
}

func main() {
    dvd := &DVDPlayer{}
    projector := &Projector{}
    lights := &Lights{}

    // With Facade (Simple)
    homeTheater := NewHomeTheaterFacade(dvd, projector, lights)
    homeTheater.WatchMovie("Inception")
}
```

---

## ✅ When to use?

*   **Simplifying Complexity**: When you want to provide a simple interface to a complex subsystem.
*   **Decoupling**: When you want to decouple the client implementation from the complex subsystem.
