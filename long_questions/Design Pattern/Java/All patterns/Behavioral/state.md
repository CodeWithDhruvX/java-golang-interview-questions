# State Pattern

## 🟢 What is it?
The **State Pattern** allows an object to alter its behavior when its internal state changes. It appears as if the object changed its class.

Think of it like a **Vending Machine**:
*   States: **Has Selection**, **Has Money**, **Sold Out**.
*   Action "Press Button":
    *   If "Has Selection": Dispense Item.
    *   If "Sold Out": Show Error.
    *   If "Idle": Ask for Money first.
*   The "Press Button" method does different things depending on the current state variable.

---

## 🎯 Strategy to Implement

1.  **Context Class**: The main class (VendingMachine) that has a reference to a `State` object.
2.  **State Interface**: Define methods for all possible actions (e.g., `insertCoin()`, `pressButton()`, `dispense()`).
3.  **Concrete States**: Create classes for each state (e.g., `IdleState`, `HasCoinState`). Implement the action methods. In the methods, change the Context's state to the next appropriate state (Transitioning).

---

## 💻 Code Example

```java
// 1. State Interface
interface State {
    void insertCoin();
    void pressButton();
}

// 2. Concrete States
class NoCoinState implements State {
    VendingMachine machine;

    public NoCoinState(VendingMachine m) { this.machine = m; }

    public void insertCoin() {
        System.out.println("Coin inserted.");
        machine.setState(machine.getHasCoinState()); // Transition
    }

    public void pressButton() {    
        System.out.println("No coin inserted. Cannot dispense.");
    }
}

class HasCoinState implements State {
    VendingMachine machine;

    public HasCoinState(VendingMachine m) { this.machine = m; }

    public void insertCoin() {
        System.out.println("Coin already inserted.");
    }

    public void pressButton() {
        System.out.println("Dispensing product...");
        machine.setState(machine.getNoCoinState()); // Transition back
    }
}

// 3. Context
class VendingMachine {
    State hasCoinState;
    State noCoinState;
    
    State currentState; // Current State

    public VendingMachine() {
        hasCoinState = new HasCoinState(this);
        noCoinState = new NoCoinState(this);
        currentState = noCoinState; // Initial State
    }

    public void setState(State state) {
        this.currentState = state;
    }

    public void insertCoin() {
        currentState.insertCoin();
    }

    public void pressButton() {
        currentState.pressButton();
    }

    // Getters for states
    public State getHasCoinState() { return hasCoinState; }
    public State getNoCoinState() { return noCoinState; }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        VendingMachine machine = new VendingMachine();

        // 1. Try to press button without coin
        machine.pressButton(); // "No coin inserted..."

        // 2. Insert coin
        machine.insertCoin();  // "Coin inserted." (State changes to HasCoin)

        // 3. Press button
        machine.pressButton(); // "Dispensing..." (State changes to NoCoin)
    }
}
```

---

## ✅ When to use?

*   **Behavior depends on State**: When an object's behavior depends on its state, and it must change its behavior at runtime depending on that state.
*   **Giant Switch Statements**: When you have large, multipart conditional statements that depend on the object's state fields.
*   **Explicit Transitions**: When appropriate transitions between states need to be enforced (e.g., a Document cannot go from "Published" back to "Draft" directly).
