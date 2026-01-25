# Memento Pattern

## 🟢 What is it?
The **Memento Pattern** lets you save and restore the previous state of an object without revealing the details of its implementation.

Think of it like **Save Points in a Game**.

---

## 🎯 Strategy to Implement

1.  **Originator**: The object whose state needs to be saved.
2.  **Memento**: A simple struct that holds the state.
3.  **Caretaker**: The object that holds the Mementos (Stack).

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Memento (Stores State)
type Memento struct {
    state string
}

func (m *Memento) GetState() string { return m.state }

// 2. Originator
type TextEditor struct {
    content string
}

func (t *TextEditor) Type(words string) {
    if t.content == "" {
        t.content = words
    } else {
        t.content += " " + words
    }
    fmt.Println("Current Text:", t.content)
}

func (t *TextEditor) Save() *Memento {
    return &Memento{state: t.content}
}

func (t *TextEditor) Restore(m *Memento) {
    t.content = m.GetState()
    fmt.Println("Restoring Text:", t.content)
}

// 3. Caretaker
type History struct {
    history []*Memento
}

func (h *History) Save(editor *TextEditor) {
    h.history = append(h.history, editor.Save())
}

func (h *History) Undo(editor *TextEditor) {
    if len(h.history) > 0 {
        // Pop last element
        lastIndex := len(h.history) - 1
        memento := h.history[lastIndex]
        h.history = h.history[:lastIndex]
        
        editor.Restore(memento)
    } else {
        fmt.Println("Nothing to undo.")
    }
}

func main() {
    editor := &TextEditor{}
    history := &History{}

    editor.Type("This is the first sentence.")
    history.Save(editor) // Save 1

    editor.Type("This is the second.")
    history.Save(editor) // Save 2

    editor.Type("This is a mistake.")

    // Oops, undo!
    history.Undo(editor) // Restores to Save 2
    history.Undo(editor) // Restores to Save 1
}
```

---

## ✅ When to use?

*   **Undo/Redo**: Restore previous states.
