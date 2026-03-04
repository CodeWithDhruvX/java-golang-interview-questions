# Low-Level Design (LLD) - ATM Machine

## Problem Statement
Design an Automated Teller Machine (ATM). The user can insert an ATM card, enter a PIN, check balance, withdraw cash, or deposit cash. Provide the object-oriented design for this system.

## Requirements
*   **Authentication:** The system determines authenticity via a Card (Card Number) and PIN.
*   **Transactions:** The system must support Balance Inquiry, Cash Withdrawal, and Deposit (Cash/Cheque).
*   **Hardware Interface:** The software must interact with hardware components: Card Reader, Keypad, Cash Dispenser, Screen, and Receipt Printer.
*   **Bank Communication:** The ATM communicates with the Bank's backend to execute transactions.

## Core Entities / Classes

1.  **ATM:** Singleton class holding the current state, Cash dispenser capacity, and references to hardware interfaces.
2.  **User / BankCustomer:** The entity attempting the transaction. Holds `Account` references.
3.  **Bank / BankNetwork:** The external system the ATM talks to for verifying PINs and balances.
4.  **Hardware Components:**
    *   `CardReader`
    *   `Keypad`
    *   `Screen`
    *   `CashDispenser`
    *   `Printer`
5.  **Transaction (Abstract):** `transactionId`, `creationDate`, `status`.
    *   `WithdrawalTransaction`
    *   `DepositTransaction`
    *   `TransferTransaction`
6.  **State / ATMState (Interface):** Defines methods like `insertCard()`, `authenticatePin()`, `selectOperation()`, `dispenseCash()`.

## Key Design Patterns Applicable
*   **State Pattern:** An ATM goes through very rigid states: `ReadyState` -> `HasCardState` -> `AuthenticatedState` -> `TransactionState` -> `ReadyState`. This is the perfect use case for the State Pattern.
*   **Command Pattern:** To encapsulate transactions (`WithdrawalCommand`, `DepositCommand`). Useful for queuing failed transactions to retry connection to the bank.
*   **Facade Pattern:** The ATM class acts as a Facade over the complex underlying hardware (Dispenser, Reader, Printer).

## Code Snippet (State Pattern transition mechanism)

```java
public interface ATMState {
    void insertCard();
    void authenticatePin(int pin);
    void requestCash(int amount);
    void ejectCard();
}

public class HasCardState implements ATMState {
    private ATM atm;
    public HasCardState(ATM atm) { this.atm = atm; }

    @Override
    public void authenticatePin(int pin) {
        boolean isCorrect = atm.getBankNetwork().verifyPin(atm.getCurrentCard(), pin);
        if (isCorrect) {
            atm.setState(atm.getAuthenticatedState());
            atm.getScreen().showMessage("PIN Verified. Please select a transaction.");
        } else {
            atm.getScreen().showMessage("Invalid PIN.");
            ejectCard();
        }
    }
    // Implement other methods...
}
```

## Follow-up Questions for Candidate
1.  What if the ATM dispenses cash but the Bank Server connection fails immediately after? How do you ensure consistency (compensating transactions)?
2.  How would you design the Cash Dispenser logic to give the minimum number of notes (e.g., greedy algorithm vs dynamic programming for coin change)?
3.  What design changes are needed if the ATM supports multiple currencies?
