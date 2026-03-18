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


## 📋 Interview Questions

### **Design & Architecture Questions**

**Q1: Why did you choose an abstract class for `Account` instead of an interface?**
**A**: "I chose an abstract class because it allows me to provide common implementation that all accounts share - like the account number, holder name, balance, and the deposit logic. While still enforcing that every account type must implement its own withdraw rules. If I used an interface, I'd have to duplicate all that common code in every single account implementation, which violates the DRY principle. The abstract class gives me the best of both worlds - shared implementation plus contract enforcement."

**Q2: What's benefit of using `HashMap<String, Account>` instead of `ArrayList<Account>`?**
**A**: "The HashMap gives me O(1) constant time lookup by account number, which is crucial for a banking system. With an ArrayList, I'd have to iterate through every account to find the one I'm looking for - that's O(n) linear time. When you have thousands or millions of accounts, that performance difference becomes significant. A customer doesn't want to wait while the system searches through every account just to find theirs."

**Q3: Why is `balance` protected instead of private?**
**A**: "I made balance protected because subclasses like SavingsAccount need direct access to implement their specific business rules - like checking minimum balance requirements during withdrawal. If I made it private, I'd have to add getter/setter methods and the subclass couldn't efficiently validate the balance during operations. Protected still maintains encapsulation from external classes while giving subclasses the access they need for proper inheritance behavior."

### **Exception Handling Questions**

**Q4: Why create a custom `InsufficientFundsException` instead of using `IllegalArgumentException`?**
**A**: "I created a custom exception because it provides specific business context that a generic exception can't offer. When someone catches `InsufficientFundsException`, they immediately know this is a banking-specific business rule violation, not just any invalid argument. It allows for more targeted catch blocks and makes the code self-documenting. Plus, I can add specific fields or methods to the custom exception later if needed, like minimum balance requirements or suggested actions."

**Q5: Where should you handle the `InsufficientFundsException` - in `Bank` class or `Main` class?**
**A**: "The exception should be handled in the `Main` class, which represents the UI layer. The `Bank` class should focus purely on business logic and let exceptions bubble up. This maintains proper separation of concerns - the bank handles the 'what' (business rules) and the UI handles the 'how' (user communication). When the UI catches the exception, it can display user-friendly messages and decide whether to retry, log the error, or take other appropriate actions."

### **OOP Concepts Questions**

**Q6: How does this demonstrate polymorphism?**
**A**: "This demonstrates runtime polymorphism beautifully. When I have an `Account` reference, it can actually hold a `SavingsAccount` object. When I call `withdraw()` on that reference, Java determines at runtime which version to execute based on the actual object type. This means I can write code that works with the base `Account` type, but get the specific behavior of whatever subclass is actually instantiated. It's the 'program to an interface, not an implementation' principle in action."

**Q7: What OOP principle is violated if you make all fields public?**
**A**: "Making all fields public would violate the encapsulation principle - one of the four fundamental OOP pillars. Encapsulation is about hiding internal state and exposing only what's necessary through controlled methods. If fields are public, any external code can directly modify the balance to negative values or invalid amounts, bypassing all business rules. It breaks data hiding, makes the code fragile, and creates maintenance nightmares because you can't change internal representation without breaking all dependent code."

**Q8: Why use `toString()` override instead of creating a `display()` method?**
**A**: "I override `toString()` because it follows Java conventions and integrates seamlessly with the ecosystem. When I use `System.out.println(account)`, it automatically calls `toString()`. Debuggers show the `toString()` result, logging frameworks use it, and testing frameworks display it in assertions. A custom `display()` method would require explicit calls everywhere. By following the standard Java pattern, my objects work naturally with all Java tools and libraries without any extra effort."

### **Code Implementation Questions**

**Q9: What happens if two accounts have the same account number?**
**A**: "Right now, the code prevents this with the `containsKey()` check before creating a new account. But if I removed that check, the HashMap would silently overwrite the existing account with the new one. This would be dangerous - you could lose someone's entire account balance! The check provides a safeguard against accidental overwrites and gives a clear error message when someone tries to create a duplicate account."

**Q10: How would you modify this to support multiple account types (Current, Fixed Deposit)?**
**A**: "I'd create new classes like `CurrentAccount` and `FixedDepositAccount` that extend the base `Account` class, each with their own business rules. Current accounts might have overdraft protection, while fixed deposits would have lock-in periods and higher interest rates. Then I'd use the Factory pattern in the `Bank.createAccount()` method - instead of hardcoding `new SavingsAccount()`, I'd have a factory that creates the appropriate account type based on user input or parameters. This makes the system extensible without modifying existing code."

**Q11: Why use `Scanner` instead of `BufferedReader` for input?**
**A**: "I chose Scanner because it's much more convenient for this console application. Scanner gives me methods like `nextInt()`, `nextDouble()`, and `next()` that automatically handle the parsing and tokenization. With BufferedReader, I'd have to read entire lines and manually parse them using `Integer.parseInt()` or `Double.parseDouble()`, which adds complexity and error-prone code. For a simple console banking system, Scanner's built-in parsing methods save development time and reduce the chance of parsing errors."

**Q12: What's the risk of using `sc.nextDouble()` for financial calculations?**
**A**: "Using `double` for financial calculations is dangerous because of floating-point precision errors. Doubles can't represent decimal values exactly - 0.1 might actually be stored as 0.10000000000000001. Over time, these tiny errors accumulate and can cause significant discrepancies in financial calculations. For banking systems, I should use `BigDecimal` which provides exact decimal arithmetic. It's slower but essential for accurate monetary calculations where every cent matters."

### **Scenario-Based Questions**

**Q13: How would you add transaction history to each account?**
**A**: "I'd add a `List<Transaction>` field to the `Account` class and create a `Transaction` record with fields like timestamp, type (deposit/withdrawal), amount, and running balance. Every time someone calls deposit() or withdraw(), I'd create a new Transaction object and add it to the list. This gives a complete audit trail for every account. I could also add methods like `getTransactionHistory()` or `getTransactionsInDateRange()` to make it useful for reporting and customer inquiries."

**Q14: How would you implement concurrent access to accounts?**
**A**: "For thread safety, I'd make the `Account` methods `synchronized` so only one thread can modify an account at a time. For the `Bank` class, I'd replace the `HashMap` with a `ConcurrentHashMap` which handles concurrent access safely. If I need finer control, I could use `ReentrantLock` objects instead of synchronized blocks - they give me more flexibility like tryLock with timeout. The key is ensuring that balance updates are atomic to prevent race conditions where multiple threads could read or modify the same account simultaneously."

**Q15: How would you persist account data to a file/database?**
**A**: "I'd add `saveToFile()` and `loadFromFile()` methods to the `Bank` class. For file persistence, I could use Java serialization - make all the classes implement `Serializable` and write the entire accounts HashMap to a file. For a database approach, I'd use JDBC to connect to a database and create tables for accounts and transactions. The key is handling exceptions properly - what if the file is corrupted or database is down? I'd also add a backup mechanism and ensure data consistency with proper transaction handling."

### **Code Quality Questions**

**Q16: What improvements would you make to error handling?**
**A**: "I'd implement more specific exception types instead of catching generic `Exception`. For example, `InvalidAmountException`, `AccountNotFoundException`, `DuplicateAccountException`. I'd also add input validation - check if deposit amounts are positive, account numbers follow the required format, names aren't empty. Instead of `System.out.println()`, I'd use a proper logging framework like Log4j or SLF4J with different log levels. And I'd use try-with-resources for the Scanner to ensure it's always closed properly, even if exceptions occur."

**Q17: How would you make this system more testable?**
**A**: "I'd separate the business logic from the UI code. Extract the `Bank` class to depend on interfaces rather than concrete classes - maybe an `AccountRepository` interface instead of direct HashMap usage. Then I can mock these dependencies in unit tests. I'd also move the menu logic out of the `main()` method into a separate `BankingController` class. This way, I can write JUnit tests for the core banking functionality without needing to simulate user input or read from the console. Dependency injection would make it easy to swap implementations for testing."

**Q18: What design patterns would you apply to extend this system?**
**A**: "I'd use several patterns. Factory Pattern for creating different account types - the `Bank` wouldn't know which concrete account it's creating. Observer Pattern for transaction notifications - account holders could subscribe to receive email or SMS alerts for transactions. Strategy Pattern for different interest calculation methods or fee structures. Maybe Repository Pattern for data access so I can easily switch between file-based and database storage. Each pattern solves a specific extensibility problem while keeping the code clean and maintainable."

### **Performance Questions**

**Q19: How would you optimize account lookup for millions of accounts?**
**A**: "For millions of accounts, I'd start by setting the proper initial capacity on the HashMap to avoid frequent resizing. I'd also implement a caching layer - maybe using LRU cache for frequently accessed accounts. If this becomes a distributed system, I'd move to database indexing with proper indexes on account numbers. For really high-performance scenarios, I might consider using a more specialized data structure like a Trie if account numbers have predictable patterns. The key is minimizing the number of hash collisions and ensuring good distribution of hash codes."

**Q20: What's the time complexity of creating and retrieving accounts?**
**A**: "Account creation is O(1) average case because HashMap put() is constant time when there are no hash collisions. Account retrieval is also O(1) average case for the same reason - HashMap get() is constant time. But in the worst-case scenario, if all account numbers hash to the same bucket, both operations degrade to O(n) where n is the number of accounts. That's why having a good hash function for account numbers is crucial. In practice, with a properly implemented hash code, we get near-constant time performance even with millions of accounts."

### **Security Questions**

**Q21: How would you add authentication to this banking system?**
**A**: "I'd add a password field to the Account class and implement a login method in the Bank class. For security, I'd never store plain passwords - I'd use bcrypt or Argon2 for password hashing. I'd implement session management with secure tokens that expire after inactivity. I'd also add role-based access control - maybe tellers can only view accounts while managers can modify them. And I'd implement multi-factor authentication for sensitive operations like large withdrawals or account closures."

**Q22: What security vulnerabilities exist in the current implementation?**
**A**: "The current system has several critical security issues. There's no input validation - a user could enter negative amounts or SQL injection if this were connected to a database. No authentication means anyone can access any account. Sensitive data like account numbers and balances are stored in plain text in memory. There's no audit logging - we can't track who did what. The Scanner input could be vulnerable to buffer overflow attacks. And there's no encryption for data at rest or in transit. In a production system, these would all need to be addressed with proper security controls."

---