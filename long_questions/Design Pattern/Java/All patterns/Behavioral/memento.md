# Memento Pattern

## 🟢 What is it?
The **Memento Pattern** lets you save and restore the previous state of an object without revealing the details of its implementation (encapsulation).

Think of it like **Save Points in a Game**:
*   You reach a boss fight. You hit "Quick Save".
*   The game freezes the current state (Health, Ammo, Position) into a **Memento**.
*   You fight and die.
*   The game loads the Memento. Your state is restored exactly as it was.

---

## 🎯 Strategy to Implement

1.  **Originator**: The object whose state needs to be saved. It creates a Memento (`save()`) and restores its state from a Memento (`restore()`).
2.  **Memento**: A simple value object (POJO) that holds the state. It should be immutable and provide data only to the Originator.
3.  **Caretaker**: The object that holds the Mementos (e.g., a History stack). It doesn't modify the Memento, just stores it.

---

## 💻 Code Example

```java
// 1. Memento (Stores State)
class Memento {
    private final String state;

    public Memento(String state) {
        this.state = state;
    }

    public String getState() {
        return state;
    }
}

// 2. Originator (The object we want to save)
class TextEditor {
    private String content;

    public void type(String words) {
        this.content = this.content == null ? words : this.content + " " + words;
        System.out.println("Current Text: " + content);
    }

    // Save state
    public Memento save() {
        return new Memento(content);
    }

    // Restore state
    public void restore(Memento memento) {
        this.content = memento.getState();
        System.out.println("Restoring Text: " + content);
    }
}

// 3. Caretaker (History Keeper)
import java.util.Stack;

class History {
    private Stack<Memento> history = new Stack<>();

    public void save(TextEditor editor) {
        history.push(editor.save());
    }

    public void undo(TextEditor editor) {
        if (!history.isEmpty()) {
            editor.restore(history.pop());
        } else {
            System.out.println("Nothing to undo.");
        }
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        TextEditor editor = new TextEditor();
        History history = new History();

        editor.type("This is the first sentence.");
        history.save(editor); // Save 1

        editor.type("This is the second.");
        history.save(editor); // Save 2

        editor.type("This is a mistake.");
        
        // Oops, undo!
        history.undo(editor); // Restores to Save 2 ("This is the second.")
        
        history.undo(editor); // Restores to Save 1 ("This is the first sentence.")
    }
}
```

---

## ✅ When to use?

*   **Undo/Redo**: When you want to produce snapshots of the object's state to be able to restore a previous state of the object.
*   **Direct Access Violation**: When direct access to the object's fields/getters/setters would violate its encapsulation (forcing you to expose private data just to save it).
