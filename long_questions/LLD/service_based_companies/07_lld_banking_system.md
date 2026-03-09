# Low-Level Design (LLD) - Banking System

## Problem Statement
Design a comprehensive Banking System that handles core banking operations including account management, transactions, loans, and customer services for Indian banks.

## Requirements
*   **Account Management:** Different types of accounts (Savings, Current, Fixed Deposit, Recurring Deposit)
*   **Customer Management:** Customer registration, KYC verification, profile management
*   **Transaction Processing:** Deposits, withdrawals, transfers, bill payments
*   **Loan Management:** Personal loans, home loans, car loans, EMI calculation
*   **ATM Operations:** Cash withdrawal, balance inquiry, mini statements
*   **Online Banking:** Fund transfers, bill payments, mobile recharge
*   **Card Services:** Debit cards, credit cards, card blocking, PIN management
*   **Branch Operations:** Cash management, teller operations, customer service
*   **Compliance:** RBI regulations, audit trails, reporting

## Core Entities / Classes

1.  **Bank (Singleton):** Central system managing all operations
2.  **Customer:** `customerId`, `name`, `email`, `phone`, `address`, `kycStatus`, `accounts`
3.  **Account (Abstract):** `accountNumber`, `accountType`, `balance`, `status`, `customer`, `createdAt`
    *   **SavingsAccount:** `interestRate`, `minimumBalance`, `withdrawalLimit`
    *   **CurrentAccount:** `overdraftLimit`, `transactionCharges`
    *   **FixedDepositAccount:** `depositAmount`, `tenure`, `interestRate`, `maturityDate`
    *   **RecurringDepositAccount:** `monthlyAmount`, `tenure`, `interestRate`
4.  **Transaction:** `transactionId`, `fromAccount`, `toAccount`, `amount`, `type`, `status`, `timestamp`, `description`
5.  **Loan:** `loanId`, `customer`, `loanType`, `principalAmount`, `interestRate`, `tenure`, `emiAmount`, `status`
6.  **EMI:** `emiId`, `loan`, `dueDate`, `amount`, `paidDate`, `status`
7.  **Card:** `cardNumber`, `cardType`, `account`, `expiryDate`, `cvv`, `pin`, `status`
8.  **Branch:** `branchCode`, `name`, `address`, `ifscCode`, `manager`, `atms`
9.  **ATM:** `atmId`, `branch`, `location`, `cashBalance`, `status`
10. **TransactionProcessor:** Handles different types of transactions
11. **InterestCalculator:** Calculates interest for accounts and loans
12. **KYCManager:** Handles KYC verification and document management

## Key Design Patterns Applicable
*   **Strategy Pattern:** Different interest calculation strategies for account types
*   **Factory Pattern:** Create different types of accounts and loans
*   **Observer Pattern:** Notify customers of account activities and loan dues
*   **Singleton Pattern:** `Bank` and `TransactionProcessor`
*   **State Pattern:** Account status management (ACTIVE → INACTIVE → FROZEN → CLOSED)
*   **Command Pattern:** Encapsulate banking operations for audit trails
*   **Template Method Pattern:** Standard transaction processing workflow
*   **Decorator Pattern:** Add additional services to basic accounts

## Code Snippet (Java/Go focus)

### Java Implementation
```java
// Account Hierarchy
public abstract class Account {
    protected String accountNumber;
    protected Customer customer;
    protected double balance;
    protected AccountStatus status;
    
    public abstract boolean withdraw(double amount);
    public abstract void deposit(double amount);
    public abstract double calculateInterest();
}

public class SavingsAccount extends Account {
    private double interestRate;
    private double minimumBalance;
    private double dailyWithdrawalLimit;
    
    @Override
    public boolean withdraw(double amount) {
        if (balance - amount >= minimumBalance && amount <= dailyWithdrawalLimit) {
            balance -= amount;
            return true;
        }
        return false;
    }
    
    @Override
    public double calculateInterest() {
        return balance * interestRate / 100;
    }
}

// Transaction Processing
public enum TransactionType {
    DEPOSIT, WITHDRAWAL, TRANSFER, BILL_PAYMENT
}

public class Transaction {
    private String transactionId;
    private Account fromAccount;
    private Account toAccount;
    private double amount;
    private TransactionType type;
    private TransactionStatus status;
    private LocalDateTime timestamp;
    
    public void process() {
        // Validate accounts, check balances, process transaction
        if (validateTransaction()) {
            executeTransaction();
            updateAccountBalances();
            notifyCustomer();
        }
    }
}
```

### Go Implementation
```go
// Account Hierarchy
type Account interface {
    Withdraw(amount float64) bool
    Deposit(amount float64)
    CalculateInterest() float64
    GetBalance() float64
}

type SavingsAccount struct {
    AccountNumber      string
    Customer          *Customer
    Balance           float64
    InterestRate      float64
    MinimumBalance    float64
    DailyWithdrawLimit float64
    Status            AccountStatus
}

func (sa *SavingsAccount) Withdraw(amount float64) bool {
    if sa.Balance-amount >= sa.MinimumBalance && amount <= sa.DailyWithdrawLimit {
        sa.Balance -= amount
        return true
    }
    return false
}

func (sa *SavingsAccount) Deposit(amount float64) {
    sa.Balance += amount
}

func (sa *SavingsAccount) CalculateInterest() float64 {
    return sa.Balance * sa.InterestRate / 100
}

// Transaction Processing
type TransactionType int

const (
    Deposit TransactionType = iota
    Withdrawal
    Transfer
    BillPayment
)

type Transaction struct {
    TransactionID string
    FromAccount   Account
    ToAccount     Account
    Amount        float64
    Type          TransactionType
    Status        TransactionStatus
    Timestamp     time.Time
}

func (t *Transaction) Process() error {
    if !t.validateTransaction() {
        return errors.New("transaction validation failed")
    }
    
    if err := t.executeTransaction(); err != nil {
        return err
    }
    
    t.updateAccountBalances()
    t.notifyCustomer()
    return nil
}
```

## Critical Design Considerations
*   **Transaction Atomicity:** ACID compliance for financial operations
*   **Concurrent Access:** Handle multiple simultaneous transactions
*   **Security:** Encryption, authentication, authorization
*   **Audit Trail:** Complete transaction history for compliance
*   **Performance:** Handle high-volume transactions during peak hours
*   **Data Integrity:** Prevent data corruption and ensure consistency
*   **Regulatory Compliance:** RBI guidelines and reporting requirements

## Indian Banking Specific Features
*   **IFSC Code:** Indian Financial System Code for branch identification
*   **NEFT/RTGS/IMPS:** Different fund transfer mechanisms
*   **Aadhaar Integration:** UIDAI-based KYC verification
*   **UPI Integration:** Unified Payments Interface support
*   **Pradhan Mantri Jan Dhan Yojana:** Financial inclusion accounts
*   **GST Compliance:** Goods and Services Tax integration
*   **Demonetization Handling:** Support for currency changes

## Interview Success Tips
*   Focus on transaction consistency and ACID properties
*   Discuss how to handle concurrent transactions safely
*   Address security concerns in banking systems
*   Explain database design for account and transaction management
*   Discuss edge cases: insufficient funds, account freezes, system failures
*   Talk about performance optimization for high-volume transactions
*   Explain how to ensure regulatory compliance and audit requirements
