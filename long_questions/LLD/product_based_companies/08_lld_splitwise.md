# Low-Level Design (LLD) - Splitwise

## Problem Statement
Design an expense sharing application like Splitwise. Users can add expenses, split them in various ways (equally, exactly, percentage), and track who owes whom.

## Requirements
*   **Users:** System should support basic user registration and friends management.
*   **Expenses:** Adding an expense specifying the payer and the participants.
*   **Splits:** Support different types of splits:
    *   `Equal`: Everyone pays an equal share.
    *   `Exact`: Amounts are explicitly given (e.g., A owes 50, B owes 30).
    *   `Percentage`: Split by % (e.g., A owes 60%, B owes 40%).
*   **Balances:** Show total balances for a user across all users (who I owe, who owes me).
*   **Algorithm (Optional but common):** "Simplify Debts" to minimize total transactions between a group.

## Core Entities / Classes

1.  **ExpenseManager:** Central facade managing the expenses and balances matrix.
2.  **User:** `id`, `name`, `email`, `phone`.
3.  **Expense (Abstract):** Contains `id`, `amount`, `paidBy` (User), list of `Split`s. 
    *   `EqualExpense`
    *   `ExactExpense`
    *   `PercentExpense`
4.  **Split (Abstract):** Represents what a single user owes.
    *   `EqualSplit`
    *   `ExactSplit`
    *   `PercentSplit` (has extra `percent` attribute)

## Key Design Patterns & Algorithms Applicable
*   **Factory Pattern:** Best used to create the specific `Expense` subclasses based on user input.
*   **Strategy / Validation Pattern:** To validate the splits before saving (e.g., do PercentSplits add up to 100%? Do ExactSplits add up to the total amount?).
*   **Graph Algorithm (Simplify Debts):** If A owes B $10, and B owes C $10, simplify implies A pays C $10 directly. This is solved using a max-heap/min-heap greedy algorithm matching the maximum debtor with the maximum creditor.

## Code Snippet (Validating Exact Splits)

```java
public abstract class Split {
    private User user;
    private double amount;
    // constructors, getters, setters
}

public class ExactExpense extends Expense {
    public ExactExpense(double amount, User paidBy, List<Split> splits) {
        super(amount, paidBy, splits);
    }

    @Override
    public boolean validate() {
        double totalAmount = getAmount();
        double sumSplitAmount = 0;
        for (Split split : getSplits()) {
            if (!(split instanceof ExactSplit)) return false;
            sumSplitAmount += split.getAmount();
        }
        // Use epsilon comparison for floating-point math
        return Math.abs(sumSplitAmount - totalAmount) < 0.001;
    }
}
```

## Follow-up Questions for Candidate
1.  How and where do you store the `BalanceSheet` so it performs globally? Do you recompute O(N) every time the user opens the app, or cache the rolling balance?
2.  Explain the algorithm and time complexity for the "Simplify Debts" feature.
3.  How does your system handle concurrent updates (e.g., Alice adds an expense to Bob at the exact same time Bob adds an expense to Alice)? 
