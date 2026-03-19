# Mini-Project 1: Console-Based Banking System

**Goal**: Demonstrate OOPs (Structs, Interfaces, Embedding), Maps, and Error Handling.

## Class Design
1.  **Account**: Base struct with `accountNumber`, `balance`, `accountHolderName`.
2.  **SavingsAccount**: Embeds `Account`, adds `interestRate`.
3.  **Bank**: Manages accounts using `map[string]AccountInterface`.
4.  **Main**: Menu-driven interface.

## Code Implementation

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Custom Error
type InsufficientFundsError struct {
	Message string
}

func (e *InsufficientFundsError) Error() string {
	return e.Message
}

// Account Interface (similar to abstract class)
type AccountInterface interface {
	GetAccountNumber() string
	GetBalance() float64
	Deposit(amount float64)
	Withdraw(amount float64) error
	String() string
}

// Base Account Struct
type Account struct {
	accountNumber string
	holderName    string
	balance       float64
}

func NewAccount(accNum, name string, bal float64) *Account {
	return &Account{
		accountNumber: accNum,
		holderName:    name,
		balance:       bal,
	}
}

func (a *Account) GetAccountNumber() string {
	return a.accountNumber
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
		fmt.Printf("Deposited: %.2f\n", amount)
	}
}

func (a *Account) Withdraw(amount float64) error {
	// Base implementation - to be overridden
	return &InsufficientFundsError{"Withdrawal not implemented for base account"}
}

func (a *Account) String() string {
	return fmt.Sprintf("Acc: %s | Name: %s | Bal: %.2f", a.accountNumber, a.holderName, a.balance)
}

// Savings Account Struct (embedding Account)
type SavingsAccount struct {
	*Account
	minBalance float64
}

func NewSavingsAccount(accNum, name string, bal float64) *SavingsAccount {
	return &SavingsAccount{
		Account:    NewAccount(accNum, name, bal),
		minBalance: 500.0,
	}
}

func (sa *SavingsAccount) Withdraw(amount float64) error {
	if sa.balance-amount < sa.minBalance {
		return &InsufficientFundsError{"Insufficient Balance (Min Req: 500)"}
	}
	sa.balance -= amount
	fmt.Printf("Withdrawn: %.2f\n", amount)
	return nil
}

// Bank Management Struct
type Bank struct {
	accounts map[string]AccountInterface
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]AccountInterface),
	}
}

func (b *Bank) CreateAccount(accNum, name string, bal float64) {
	if _, exists := b.accounts[accNum]; exists {
		fmt.Println("Account already exists!")
		return
	}
	b.accounts[accNum] = NewSavingsAccount(accNum, name, bal)
	fmt.Println("Account Created Successfully.")
}

func (b *Bank) GetAccount(accNum string) AccountInterface {
	return b.accounts[accNum]
}

func (b *Bank) DisplayAll() {
	for _, acc := range b.accounts {
		fmt.Println(acc)
	}
}

// Main Function
func main() {
	bank := NewBank()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Create Account\n2. Deposit\n3. Withdraw\n4. Show All\n5. Exit")
		fmt.Print("Choose: ")
		
		if !scanner.Scan() {
			break
		}
		
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Acc No: ")
			scanner.Scan()
			accNum := scanner.Text()
			
			fmt.Print("Enter Name: ")
			scanner.Scan()
			name := scanner.Text()
			
			fmt.Print("Enter Initial Balance: ")
			scanner.Scan()
			bal, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid balance!")
				continue
			}
			
			bank.CreateAccount(accNum, name, bal)
			
		case 2:
			fmt.Print("Enter Acc No: ")
			scanner.Scan()
			dAcc := scanner.Text()
			
			da := bank.GetAccount(dAcc)
			if da != nil {
				fmt.Print("Amount: ")
				scanner.Scan()
				amount, err := strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					fmt.Println("Invalid amount!")
					continue
				}
				da.Deposit(amount)
			} else {
				fmt.Println("Account not found.")
			}
			
		case 3:
			fmt.Print("Enter Acc No: ")
			scanner.Scan()
			wAcc := scanner.Text()
			
			wa := bank.GetAccount(wAcc)
			if wa != nil {
				fmt.Print("Amount: ")
				scanner.Scan()
				amount, err := strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					fmt.Println("Invalid amount!")
					continue
				}
				
				err = wa.Withdraw(amount)
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}
			} else {
				fmt.Println("Account not found.")
			}
			
		case 4:
			bank.DisplayAll()
			
		case 5:
			fmt.Println("Goodbye!")
			return
			
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
```

## Key Code Concepts Used
*   **Interface**: `AccountInterface` defines the contract (similar to abstract class).
*   **Embedding**: `SavingsAccount` embeds `Account` struct (similar to inheritance).
*   **Encapsulation**: Private fields (lowercase) with public methods (uppercase).
*   **Maps**: `map[string]AccountInterface` for O(1) account retrieval.
*   **Error Handling**: Custom `InsufficientFundsError` implementing `error` interface.

---

## 📋 Interview Questions

### **Design & Architecture Questions**

**Q1: Why did you choose an interface for `AccountInterface` instead of a struct?**
**A**: "I chose an interface because Go doesn't have classes or inheritance like Java. The interface provides the contract that all account types must follow, while allowing different implementations. The embedding pattern in Go (`SavingsAccount` embedding `Account`) provides similar benefits to inheritance - code reuse and composition. This approach follows Go's philosophy of composition over inheritance and gives us the flexibility to have multiple account types that all satisfy the same interface."

**Q2: What's benefit of using `map[string]AccountInterface` instead of `[]AccountInterface`?**
**A**: "The map gives me O(1) constant time lookup by account number, which is crucial for a banking system. With a slice, I'd have to iterate through every account to find the one I'm looking for - that's O(n) linear time. When you have thousands or millions of accounts, that performance difference becomes significant. A customer doesn't want to wait while the system searches through every account just to find theirs."

**Q3: Why are the Account fields lowercase instead of uppercase?**
**A**: "In Go, field visibility is controlled by capitalization. Lowercase fields are package-private, which means they can only be accessed within the same package. This provides encapsulation similar to private fields in Java. Uppercase fields would be public and accessible from any package, breaking encapsulation. The lowercase fields with uppercase getter methods follow Go's conventions and maintain proper data hiding."

### **Error Handling Questions**

**Q4: Why create a custom `InsufficientFundsError` instead of using a standard error?**
**A**: "I created a custom error type because it provides specific business context that a generic error can't offer. When someone catches an `*InsufficientFundsError`, they immediately know this is a banking-specific business rule violation, not just any error. The custom type allows for type assertions and more targeted error handling. Plus, I can add additional fields or methods to the error type later if needed, like minimum balance requirements or suggested actions."

**Q5: Where should you handle the `InsufficientFundsError` - in `Bank` struct or `main` function?**
**A**: "The error should be handled in the `main` function, which represents the UI layer. The `Bank` struct should focus purely on business logic and let errors bubble up. This maintains proper separation of concerns - the bank handles the 'what' (business rules) and the UI handles the 'how' (user communication). When the UI catches the error, it can display user-friendly messages and decide whether to retry, log the error, or take other appropriate actions."

### **Go-Specific Concepts Questions**

**Q6: How does embedding in Go compare to inheritance in Java?**
**A**: "Embedding in Go is more like composition than inheritance. When `SavingsAccount` embeds `Account`, it gets all the methods and fields of `Account`, but it's still a distinct type. Unlike Java inheritance, Go doesn't have polymorphic method calls through embedding by default - you need to use interfaces for that. The embedded struct's methods are promoted to the containing struct, which gives a similar experience to inheritance, but it's fundamentally composition. This approach avoids the diamond problem and keeps the type system simpler."

**Q7: What's the advantage of using interfaces in this banking system?**
**A**: "Interfaces give us polymorphism without inheritance. I can write code that works with the `AccountInterface` type, but at runtime it can be a `SavingsAccount`, `CurrentAccount`, or any other account type that satisfies the interface. This makes the system extensible - I can add new account types without modifying existing code. Interfaces also make testing easier since I can create mock implementations. This is Go's way of achieving the benefits of polymorphism without the complexity of inheritance hierarchies."

**Q8: Why use `bufio.Scanner` instead of `fmt.Scanln` for input?**
**A**: "I chose `bufio.Scanner` because it's more robust for reading user input. `fmt.Scanln` can have issues with whitespace and newlines, while `Scanner` provides better control over input reading. `Scanner` also handles different input scenarios more gracefully and is generally considered the preferred way to read line-by-line input in Go console applications. It gives me more reliable input handling, especially when dealing with mixed data types like strings and numbers."

### **Code Implementation Questions**

**Q9: What happens if two accounts have the same account number?**
**A**: "Right now, the code prevents this with the existence check before creating a new account. But if I removed that check, the map would silently overwrite the existing account with the new one. This would be dangerous - you could lose someone's entire account balance! The check provides a safeguard against accidental overwrites and gives a clear error message when someone tries to create a duplicate account. In a production system, I might also want to generate unique account numbers automatically."

**Q10: How would you modify this to support multiple account types (Current, Fixed Deposit)?**
**A**: "I'd create new structs like `CurrentAccount` and `FixedDepositAccount` that embed the base `Account` struct, each implementing the `AccountInterface` with their own business rules. Current accounts might have overdraft protection, while fixed deposits would have lock-in periods and higher interest rates. Then I'd use a factory pattern in the `Bank.CreateAccount()` method - instead of hardcoding the account type, I'd have a factory that creates the appropriate account type based on user input or parameters. This makes the system extensible without modifying existing code."

**Q11: Why use `float64` instead of `decimal` for financial calculations?**
**A**: "Using `float64` for financial calculations is actually problematic in Go, just like in Java. Floats can't represent decimal values exactly - 0.1 might be stored as 0.10000000000000001. Over time, these tiny errors accumulate and can cause significant discrepancies in financial calculations. For a production banking system, I should use a decimal library like `shopspring/decimal` which provides exact decimal arithmetic. The current `float64` implementation is fine for learning purposes but not suitable for real financial applications."

**Q12: What's the risk of using `strconv.ParseFloat` for user input?**
**A**: "The main risk is that user input can be malformed and cause parsing errors. If a user enters non-numeric text when we expect a number, `ParseFloat` returns an error. In the current code, I handle this by checking the error and showing an error message, but in a production system, I'd want more robust validation. I should also validate ranges (no negative amounts, reasonable limits), and potentially use more sophisticated input validation to prevent injection attacks or malformed data."

### **Scenario-Based Questions**

**Q13: How would you add transaction history to each account?**
**A**: "I'd add a `[]Transaction` field to the `Account` struct and create a `Transaction` struct with fields like Timestamp, Type (deposit/withdrawal), Amount, and RunningBalance. Every time someone calls Deposit() or Withdraw(), I'd create a new Transaction and append it to the slice. This gives a complete audit trail for every account. I could also add methods like `GetTransactionHistory()` or `GetTransactionsInDateRange()` to make it useful for reporting and customer inquiries."

**Q14: How would you implement concurrent access to accounts?**
**A**: "For thread safety, I'd use mutexes to protect account operations. I'd add a `sync.RWMutex` field to the `Account` struct and use `Lock()/Unlock()` for write operations (deposit/withdraw) and `RLock()/RUnlock()` for read operations. For the `Bank` struct, I'd replace the regular map with a `sync.Map` which handles concurrent access safely. This ensures that balance updates are atomic and prevents race conditions where multiple goroutines could read or modify the same account simultaneously."

**Q15: How would you persist account data to a file/database?**
**A**: "I'd add `SaveToFile()` and `LoadFromFile()` methods to the `Bank` struct. For file persistence, I could use JSON encoding with `json.Marshal()` and `json.Unmarshal()` since all structs are JSON-serializable. For a database approach, I'd use database/sql with an appropriate driver. The key is handling errors properly - what if the file is corrupted or database is down? I'd also add a backup mechanism and ensure data consistency with proper transaction handling. Go's standard library makes both JSON and database operations straightforward."

### **Code Quality Questions**

**Q16: What improvements would you make to error handling?**
**A**: "I'd implement more specific error types instead of generic errors. For example, `InvalidAmountError`, `AccountNotFoundError`, `DuplicateAccountError`. I'd also add input validation - check if deposit amounts are positive, account numbers follow required format, names aren't empty. Instead of `fmt.Println()`, I'd use a proper logging framework like `log` package with different log levels. And I'd add more comprehensive error checking for all operations that can fail."

**Q17: How would you make this system more testable?**
**A**: "I'd separate the business logic from the UI code. Extract the `Bank` struct to depend on interfaces rather than concrete implementations - maybe an `AccountRepository` interface instead of direct map usage. Then I can mock these dependencies in unit tests. I'd also move the menu logic out of the `main()` function into a separate `BankingController` struct. This way, I can write table-driven tests for the core banking functionality without needing to simulate user input or read from the console."

**Q18: What Go-specific design patterns would you apply to extend this system?**
**A**: "I'd use several Go-specific patterns. The Builder pattern for creating complex account configurations. The Option pattern for flexible account creation with optional parameters. The Decorator pattern for adding features like transaction logging to existing accounts. Maybe the Repository pattern for data access so I can easily switch between in-memory, file, and database storage. Each pattern leverages Go's strengths like interfaces, first-class functions, and struct composition."

### **Performance Questions**

**Q19: How would you optimize account lookup for millions of accounts?**
**A**: "For millions of accounts, I'd start by setting the proper initial capacity on the map to avoid frequent resizing. I'd also implement a caching layer - maybe using an LRU cache for frequently accessed accounts. If this becomes a distributed system, I'd move to database indexing with proper indexes on account numbers. For really high-performance scenarios, I might consider using a more specialized data structure, but the key is minimizing hash collisions and ensuring good distribution of hash codes for account numbers."

**Q20: What's the time complexity of creating and retrieving accounts?**
**A**: "Account creation is O(1) average case because map insertion is constant time when there are no hash collisions. Account retrieval is also O(1) average case for the same reason - map lookup is constant time. But in the worst-case scenario, if all account numbers hash to the same bucket, both operations degrade to O(n) where n is the number of accounts. That's why having a good hash function for account numbers is crucial. In practice, with Go's built-in map implementation, we get near-constant time performance even with millions of accounts."

### **Security Questions**

**Q21: How would you add authentication to this banking system?**
**A**: "I'd add a password field to the Account struct and implement a login method in the Bank struct. For security, I'd never store plain passwords - I'd use bcrypt or Argon2 for password hashing. I'd implement session management with secure tokens that expire after inactivity. I'd also add role-based access control - maybe tellers can only view accounts while managers can modify them. And I'd implement multi-factor authentication for sensitive operations like large withdrawals or account closures."

**Q22: What security vulnerabilities exist in the current implementation?**
**A**: "The current system has several critical security issues. There's no input validation - a user could enter negative amounts or malformed data. No authentication means anyone can access any account. Sensitive data like account numbers and balances are stored in plain text in memory. There's no audit logging - we can't track who did what. The scanner input could be vulnerable to buffer overflow attacks. And there's no encryption for data at rest or in transit. In a production system, these would all need to be addressed with proper security controls."

---
