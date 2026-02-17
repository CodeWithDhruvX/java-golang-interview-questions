# Mini-Project 1: Console-Based Banking System

**Goal**: Demonstrate OOPs (Encapsulation, Inheritance), Collections (HashMap), and Exception Handling.

## Class Design
1.  **Account**: Base class with `accountNumber`, `balance`, `accountHolderName`.
2.  **SavingsAccount**: Extends `Account`, adds `interestRate`.
3.  **Bank**: Manages accounts using `HashMap<String, Account>`.
4.  **Main**: Menu-driven interface.

## Code Implementation

```java
import java.util.*;

// Custom Exception
class InsufficientFundsException extends Exception {
    public InsufficientFundsException(String msg) { super(msg); }
}

// Base Account Class
abstract class Account {
    private String accountNumber;
    private String holderName;
    protected double balance;

    public Account(String accountNumber, String holderName, double balance) {
        this.accountNumber = accountNumber;
        this.holderName = holderName;
        this.balance = balance;
    }

    public String getAccountNumber() { return accountNumber; }
    public double getBalance() { return balance; }

    public void deposit(double amount) {
        if (amount > 0) {
            balance += amount;
            System.out.println("Deposited: " + amount);
        }
    }

    public abstract void withdraw(double amount) throws InsufficientFundsException;

    @Override
    public String toString() {
        return "Acc: " + accountNumber + " | Name: " + holderName + " | Bal: " + balance;
    }
}

// Savings Account
class SavingsAccount extends Account {
    private double minBalance = 500.0;

    public SavingsAccount(String accNum, String name, double bal) {
        super(accNum, name, bal);
    }

    @Override
    public void withdraw(double amount) throws InsufficientFundsException {
        if (balance - amount < minBalance) {
            throw new InsufficientFundsException("Insufficient Balance (Min Req: 500)");
        }
        balance -= amount;
        System.out.println("Withdrawn: " + amount);
    }
}

// Bank Management Class
class Bank {
    private Map<String, Account> accounts = new HashMap<>();

    public void createAccount(String accNum, String name, double bal) {
        if (accounts.containsKey(accNum)) {
            System.out.println("Account already exists!");
            return;
        }
        accounts.put(accNum, new SavingsAccount(accNum, name, bal));
        System.out.println("Account Created Successfully.");
    }

    public Account getAccount(String accNum) {
        return accounts.get(accNum);
    }

    public void displayAll() {
        for (Account acc : accounts.values()) {
            System.out.println(acc);
        }
    }
}

// Main Class
public class BankingSystem {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        Bank bank = new Bank();
        
        while (true) {
            System.out.println("\n1. Create Account\n2. Deposit\n3. Withdraw\n4. Show All\n5. Exit");
            System.out.print("Choose: ");
            int choice = sc.nextInt();

            try {
                switch (choice) {
                    case 1:
                        System.out.print("Enter Acc No: ");
                        String accNum = sc.next();
                        System.out.print("Enter Name: ");
                        String name = sc.next();
                        System.out.print("Enter Initial Balance: ");
                        double bal = sc.nextDouble();
                        bank.createAccount(accNum, name, bal);
                        break;
                    case 2:
                        System.out.print("Enter Acc No: ");
                        String dAcc = sc.next();
                        Account da = bank.getAccount(dAcc);
                        if (da != null) {
                            System.out.print("Amount: ");
                            da.deposit(sc.nextDouble());
                        } else System.out.println("Account not found.");
                        break;
                    case 3:
                        System.out.print("Enter Acc No: ");
                        String wAcc = sc.next();
                        Account wa = bank.getAccount(wAcc);
                        if (wa != null) {
                            System.out.print("Amount: ");
                            wa.withdraw(sc.nextDouble());
                        } else System.out.println("Account not found.");
                        break;
                    case 4:
                        bank.displayAll();
                        break;
                    case 5:
                        System.exit(0);
                }
            } catch (Exception e) {
                System.out.println("Error: " + e.getMessage());
            }
        }
    }
}
```

## Key Code Concepts Used
*   **Abstraction**: `Account` class defines the contract.
*   **Polymorphism**: `withdraw()` overridden in `SavingsAccount`.
*   **Encapsulation**: Private fields with public getters/methods.
*   **Collections**: `HashMap` for O(1) account retrieval.
*   **Exception Handling**: Custom `InsufficientFundsException`.
