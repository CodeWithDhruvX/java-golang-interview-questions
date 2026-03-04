# Low-Level Design (LLD) - Vending Machine

## Problem Statement
Design a state machine for a fully automated Vending Machine. The machine allows users to insert coins, select products, and collect the item along with any change. Provide the low-level design.

## Requirements
*   **Inventory Management:** Machine has different products with respective prices and quantities.
*   **Currency Acceptance:** Machine should accept coins (1, 5, 10, 25 Cents/Rupees) or notes.
*   **State Management:** The machine goes through various states (NoCoinInserted, HasCoin, DispensingProduct, ReturnChange).
*   **Actions:** User can select a product, cancel a transaction at any time before dispensing, and collect their dispensed item and change.

## Core Entities / Classes

1.  **VendingMachine:** The context class holding the current `State` and `Inventory`.
2.  **Item (Enum or Class):** Represents Coke, Pepsi, Water, Soda. Properties: `name`, `price`.
3.  **Inventory:** Manages a map of `Item` and integer count.
4.  **Coin (Enum):** Value of coins accepted.
5.  **State (Interface):** Defines methods like `insertCoin()`, `selectItem()`, `dispenseItem()`, `refund()`.
6.  *Concrete States:*
    *   `IdleState`
    *   `HasMoneyState`
    *   `DispensingState`

## Key Design Patterns Applicable
*   **State Pattern:** The most critical pattern here. It enables the Vending Machine to alter its behavior when its internal state changes.
*   **Factory Pattern:** Creating the initial state or items.
*   **Singleton Pattern:** Vending machine instance (optional, depending on requirements).

## Code Snippet (State Pattern in Java)

```java
public interface State {
    public void insertCoin(VendingMachine machine, Coin coin);
    public void selectItem(VendingMachine machine, Item item);
    public void dispenseItem(VendingMachine machine);
    public void cancelAndRefundTransaction(VendingMachine machine);
}

public class IdleState implements State {
    @Override
    public void insertCoin(VendingMachine machine, Coin coin) {
        machine.addCoin(coin);
        machine.setState(machine.getHasMoneyState());
    }

    @Override
    public void selectItem(VendingMachine machine, Item item) {
        throw new IllegalStateException("Insert money first!");
    }
    // ... Implement dispense and refund
}
```

## Follow-up Questions for Candidate
1.  How will you handle multi-threading if multiple users try to buy at the exact same millisecond? (Synchronization)
2.  How will you expand the vending machine to support credit card payments?
3.  What is the time complexity of the coin-change minimum notes algorithm if we have to return exact change?
