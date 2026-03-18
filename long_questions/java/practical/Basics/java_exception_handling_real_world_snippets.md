# Java Exception Handling — Real-World Practical Code Snippets

> **Topics:** Real-world scenarios using try-catch, custom exceptions, exception chaining, resource management, and error handling patterns in business applications

---

## 📋 Reading Progress

- [ ] **Section 1:** Business Operations & Validation (Q1–Q8)
- [ ] **Section 2:** Resource Management & Cleanup (Q9–Q16)
- [ ] **Section 3:** Network & Database Operations (Q17–Q24)
- [ ] **Section 4:** Error Recovery & Resilience (Q25–Q32)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Business Operations & Validation (Q1–Q8)

### 1. Payment Processing — Transaction Validation
**Q: Handle payment processing with comprehensive validation. What is the output?**
```java
import java.util.*;

class InsufficientFundsException extends Exception {
    public InsufficientFundsException(String message) { super(message); }
}

class InvalidAmountException extends Exception {
    public InvalidAmountException(String message) { super(message); }
}

class Account {
    private String accountNumber;
    private double balance;
    
    public Account(String accountNumber, double balance) {
        this.accountNumber = accountNumber;
        this.balance = balance;
    }
    
    public void withdraw(double amount) throws InsufficientFundsException, InvalidAmountException {
        if (amount <= 0) {
            throw new InvalidAmountException("Withdrawal amount must be positive");
        }
        if (amount > balance) {
            throw new InsufficientFundsException(
                String.format("Insufficient funds: Available $%.2f, Requested $%.2f", 
                    balance, amount));
        }
        balance -= amount;
        System.out.printf("Withdrawn $%.2f. New balance: $%.2f%n", amount, balance);
    }
    
    public double getBalance() { return balance; }
}

public class Main {
    public static void main(String[] args) {
        Account account = new Account("ACC123", 1000.0);
        
        double[] withdrawalAmounts = {500.0, -50.0, 600.0, 200.0};
        
        for (double amount : withdrawalAmounts) {
            try {
                System.out.printf("Attempting to withdraw $%.2f...%n", amount);
                account.withdraw(amount);
            } catch (InvalidAmountException e) {
                System.out.println("Error: " + e.getMessage());
            } catch (InsufficientFundsException e) {
                System.out.println("Error: " + e.getMessage());
            } catch (Exception e) {
                System.out.println("Unexpected error: " + e.getMessage());
            }
        }
        
        System.out.printf("Final balance: $%.2f%n", account.getBalance());
    }
}
```
**A:** 
```
Attempting to withdraw $500.00...
Withdrawn $500.00. New balance: $500.00
Attempting to withdraw $-50.00...
Error: Withdrawal amount must be positive
Attempting to withdraw $600.00...
Error: Insufficient funds: Available $500.00, Requested $600.00
Attempting to withdraw $200.00...
Withdrawn $200.00. New balance: $300.00
Final balance: $300.00
```

---

### 2. Order Processing — Business Rule Validation
**Q: Validate orders against business rules with custom exceptions. What is the output?**
```java
import java.util.*;

class OrderValidationException extends Exception {
    public OrderValidationException(String message) { super(message); }
}

class OutOfStockException extends Exception {
    public OutOfStockException(String message) { super(message); }
}

class MinimumOrderException extends Exception {
    public MinimumOrderException(String message) { super(message); }
}

class Product {
    private String name;
    private double price;
    private int stock;
    private int minimumOrderQuantity;
    
    public Product(String name, double price, int stock, int minimumOrderQuantity) {
        this.name = name;
        this.price = price;
        this.stock = stock;
        this.minimumOrderQuantity = minimumOrderQuantity;
    }
    
    public void validateOrder(int quantity) throws OutOfStockException, MinimumOrderException {
        if (quantity < minimumOrderQuantity) {
            throw new MinimumOrderException(
                String.format("Minimum order quantity for %s is %d, requested: %d", 
                    name, minimumOrderQuantity, quantity));
        }
        if (quantity > stock) {
            throw new OutOfStockException(
                String.format("Insufficient stock for %s: Available %d, Requested: %d", 
                    name, stock, quantity));
        }
    }
    
    public double calculateTotal(int quantity) {
        return price * quantity;
    }
    
    public void reduceStock(int quantity) {
        stock -= quantity;
    }
}

class Order {
    private Map<String, Integer> items = new HashMap<>();
    private double total = 0.0;
    
    public void addItem(Product product, int quantity) throws OrderValidationException {
        try {
            product.validateOrder(quantity);
            items.put(product.name(), quantity);
            total += product.calculateTotal(quantity);
            product.reduceStock(quantity);
            System.out.printf("Added %d x %s to order. Subtotal: $%.2f%n", 
                quantity, product.name(), product.calculateTotal(quantity));
        } catch (OutOfStockException | MinimumOrderException e) {
            throw new OrderValidationException("Failed to add " + product.name() + ": " + e.getMessage());
        }
    }
    
    public double getTotal() { return total; }
}

public class Main {
    public static void main(String[] args) {
        Product laptop = new Product("Laptop", 999.99, 10, 1);
        Product mouse = new Product("Mouse", 29.99, 5, 5);
        Product keyboard = new Product("Keyboard", 79.99, 3, 2);
        
        Order order = new Order();
        
        Object[][] orderItems = {
            {laptop, 2},
            {mouse, 3},
            {keyboard, 1},
            {mouse, 5}
        };
        
        for (Object[] itemData : orderItems) {
            Product product = (Product) itemData[0];
            int quantity = (int) itemData[1];
            
            try {
                order.addItem(product, quantity);
            } catch (OrderValidationException e) {
                System.out.println("Order validation failed: " + e.getMessage());
            }
        }
        
        System.out.printf("Order total: $%.2f%n", order.getTotal());
    }
}
```
**A:** 
```
Added 2 x Laptop to order. Subtotal: $1999.98
Order validation failed: Failed to add Mouse: Minimum order quantity for Mouse is 5, requested: 3
Order validation failed: Failed to add Keyboard: Minimum order quantity for Keyboard is 2, requested: 1
Added 5 x Mouse to order. Subtotal: $149.95
Order total: $2149.93
```

---

### 3. User Registration — Input Validation
**Q: Handle user registration with comprehensive input validation. What is the output?**
```java
import java.util.*;
import java.util.regex.*;

class ValidationException extends Exception {
    public ValidationException(String message) { super(message); }
}

class UserValidator {
    private static final Pattern EMAIL_PATTERN = Pattern.compile(
        "^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+$");
    private static final Pattern PHONE_PATTERN = Pattern.compile(
        "^[+]?[1-9]\\d{1,14}$");
    
    public static void validateEmail(String email) throws ValidationException {
        if (email == null || email.trim().isEmpty()) {
            throw new ValidationException("Email cannot be empty");
        }
        if (!EMAIL_PATTERN.matcher(email).matches()) {
            throw new ValidationException("Invalid email format: " + email);
        }
    }
    
    public static void validatePhone(String phone) throws ValidationException {
        if (phone == null || phone.trim().isEmpty()) {
            throw new ValidationException("Phone number cannot be empty");
        }
        String cleanPhone = phone.replaceAll("[^\\d+]", "");
        if (!PHONE_PATTERN.matcher(cleanPhone).matches()) {
            throw new ValidationException("Invalid phone format: " + phone);
        }
    }
    
    public static void validateAge(int age) throws ValidationException {
        if (age < 13) {
            throw new ValidationException("User must be at least 13 years old");
        }
        if (age > 120) {
            throw new ValidationException("Invalid age: " + age);
        }
    }
    
    public static void validatePassword(String password) throws ValidationException {
        if (password == null || password.length() < 8) {
            throw new ValidationException("Password must be at least 8 characters long");
        }
        if (!password.matches(".*[A-Z].*")) {
            throw new ValidationException("Password must contain at least one uppercase letter");
        }
        if (!password.matches(".*[a-z].*")) {
            throw new ValidationException("Password must contain at least one lowercase letter");
        }
        if (!password.matches(".*\\d.*")) {
            throw new ValidationException("Password must contain at least one digit");
        }
    }
}

class User {
    private String email;
    private String phone;
    private int age;
    private String password;
    
    public User(String email, String phone, int age, String password) {
        this.email = email;
        this.phone = phone;
        this.age = age;
        this.password = password;
    }
    
    public void validate() throws ValidationException {
        List<String> errors = new ArrayList<>();
        
        try {
            UserValidator.validateEmail(email);
        } catch (ValidationException e) {
            errors.add(e.getMessage());
        }
        
        try {
            UserValidator.validatePhone(phone);
        } catch (ValidationException e) {
            errors.add(e.getMessage());
        }
        
        try {
            UserValidator.validateAge(age);
        } catch (ValidationException e) {
            errors.add(e.getMessage());
        }
        
        try {
            UserValidator.validatePassword(password);
        } catch (ValidationException e) {
            errors.add(e.getMessage());
        }
        
        if (!errors.isEmpty()) {
            throw new ValidationException("Validation failed: " + String.join(", ", errors));
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Object[][] userData = {
            {"john@email.com", "+1234567890", 25, "Password123"},
            {"invalid-email", "1234567890", 16, "weak"},
            {"jane@email.com", "+1-800-555-1234", 12, "StrongPass1"},
            {"bob@email.com", "555-1234", 30, "NoDigitsHere"},
            {"alice@email.com", "+44 20 7946 0958", 28, "ValidPass123"}
        };
        
        for (Object[] data : userData) {
            String email = (String) data[0];
            String phone = (String) data[1];
            int age = (int) data[2];
            String password = (String) data[3];
            
            System.out.printf("Registering user: %s...%n", email);
            
            try {
                User user = new User(email, phone, age, password);
                user.validate();
                System.out.println("  ✓ Registration successful!");
            } catch (ValidationException e) {
                System.out.println("  ✗ Registration failed: " + e.getMessage());
            }
            System.out.println();
        }
    }
}
```
**A:** 
```
Registering user: john@email.com...
  ✓ Registration successful!

Registering user: invalid-email...
  ✗ Registration failed: Validation failed: Invalid email format: invalid-email, Invalid phone format: 1234567890

Registering user: jane@email.com...
  ✗ Registration failed: Validation failed: User must be at least 13 years old

Registering user: bob@email.com...
  ✗ Registration failed: Validation failed: Password must contain at least one digit

Registering user: alice@email.com...
  ✓ Registration successful!
```

---

### 4. Inventory Management — Stock Operations
**Q: Handle inventory operations with proper exception handling. What is the output?**
```java
import java.util.*;

class InventoryException extends Exception {
    public InventoryException(String message) { super(message); }
}

class ProductNotFoundException extends InventoryException {
    public ProductNotFoundException(String productId) {
        super("Product not found: " + productId);
    }
}

class InsufficientInventoryException extends InventoryException {
    public InsufficientInventoryException(String productId, int requested, int available) {
        super(String.format("Insufficient inventory for %s: Requested %d, Available %d", 
            productId, requested, available));
    }
}

class InventoryItem {
    private String productId;
    private String name;
    private int quantity;
    private double price;
    
    public InventoryItem(String productId, String name, int quantity, double price) {
        this.productId = productId;
        this.name = name;
        this.quantity = quantity;
        this.price = price;
    }
    
    public void addStock(int amount) throws InventoryException {
        if (amount <= 0) {
            throw new InventoryException("Stock addition amount must be positive: " + amount);
        }
        quantity += amount;
        System.out.printf("Added %d units of %s. New stock: %d%n", 
            amount, name, quantity);
    }
    
    public void removeStock(int amount) throws InsufficientInventoryException {
        if (amount > quantity) {
            throw new InsufficientInventoryException(productId, amount, quantity);
        }
        quantity -= amount;
        System.out.printf("Removed %d units of %s. Remaining stock: %d%n", 
            amount, name, quantity);
    }
    
    public String getProductId() { return productId; }
    public String getName() { return name; }
    public int getQuantity() { return quantity; }
    public double getPrice() { return price; }
}

class InventoryManager {
    private Map<String, InventoryItem> inventory = new HashMap<>();
    
    public void addProduct(InventoryItem item) {
        inventory.put(item.getProductId(), item);
        System.out.printf("Added product to inventory: %s (%s)%n", 
            item.getName(), item.getProductId());
    }
    
    public void processStockAdjustment(String productId, int adjustment) 
            throws InventoryException {
        InventoryItem item = inventory.get(productId);
        if (item == null) {
            throw new ProductNotFoundException(productId);
        }
        
        if (adjustment > 0) {
            item.addStock(adjustment);
        } else if (adjustment < 0) {
            try {
                item.removeStock(Math.abs(adjustment));
            } catch (InsufficientInventoryException e) {
                throw new InventoryException("Failed to adjust stock for " + 
                    item.getName() + ": " + e.getMessage());
            }
        } else {
            throw new InventoryException("Stock adjustment cannot be zero");
        }
    }
    
    public void printInventoryStatus() {
        System.out.println("\nCurrent Inventory Status:");
        inventory.values().forEach(item -> 
            System.out.printf("  %s (%s): %d units @ $%.2f%n",
                item.getName(), item.getProductId(), item.getQuantity(), item.getPrice()));
    }
}

public class Main {
    public static void main(String[] args) {
        InventoryManager manager = new InventoryManager();
        
        // Add products
        manager.addProduct(new InventoryItem("P001", "Laptop", 10, 999.99));
        manager.addProduct(new InventoryItem("P002", "Mouse", 50, 29.99));
        manager.addProduct(new InventoryItem("P003", "Keyboard", 25, 79.99));
        
        manager.printInventoryStatus();
        
        // Process stock adjustments
        Object[][] adjustments = {
            {"P001", 5},    // Add 5 laptops
            {"P002", -10},  // Remove 10 mice
            {"P003", -30},  // Try to remove 30 keyboards (should fail)
            {"P004", 5},    // Non-existent product
            {"P002", 0}     // Zero adjustment
        };
        
        for (Object[] adjustment : adjustments) {
            String productId = (String) adjustment[0];
            int amount = (int) adjustment[1];
            
            System.out.printf("\nProcessing stock adjustment for %s: %+d%n", productId, amount);
            
            try {
                manager.processStockAdjustment(productId, amount);
            } catch (InventoryException e) {
                System.out.println("Error: " + e.getMessage());
            }
        }
        
        manager.printInventoryStatus();
    }
}
```
**A:** 
```
Added product to inventory: Laptop (P001)
Added product to inventory: Mouse (P002)
Added product to inventory: Keyboard (P003)

Current Inventory Status:
  Laptop (P001): 10 units @ $999.99
  Mouse (P002): 50 units @ $29.99
  Keyboard (P003): 25 units @ $79.99

Processing stock adjustment for P001: +5
Added 5 units of Laptop. New stock: 15

Processing stock adjustment for P002: -10
Removed 10 units of Mouse. Remaining stock: 40

Processing stock adjustment for P003: -30
Error: Failed to adjust stock for Keyboard: Insufficient inventory for P003: Requested 30, Available 25

Processing stock adjustment for P004: +5
Error: Product not found: P004

Processing stock adjustment for P002: 0
Error: Stock adjustment cannot be zero

Current Inventory Status:
  Laptop (P001): 15 units @ $999.99
  Mouse (P002): 40 units @ $29.99
  Keyboard (P003): 25 units @ $79.99
```

---

### 5. File Processing — Data Import with Error Handling
**Q: Process CSV data import with comprehensive error handling. What is the output?**
```java
import java.util.*;
import java.util.regex.*;

class DataImportException extends Exception {
    public DataImportException(String message) { super(message); }
    public DataImportException(String message, Throwable cause) { super(message, cause); }
}

class InvalidFormatException extends DataImportException {
    public InvalidFormatException(String message) { super(message); }
}

class InvalidDataException extends DataImportException {
    public InvalidDataException(String message) { super(message); }
}

class CustomerData {
    private String id;
    private String name;
    private String email;
    private int age;
    private double balance;
    
    public CustomerData(String id, String name, String email, int age, double balance) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.age = age;
        this.balance = balance;
    }
    
    public static CustomerData fromCsvLine(String line, int lineNumber) throws DataImportException {
        try {
            String[] parts = line.split(",");
            if (parts.length != 5) {
                throw new InvalidFormatException(
                    String.format("Line %d: Expected 5 fields, got %d", lineNumber, parts.length));
            }
            
            String id = parts[0].trim();
            String name = parts[1].trim();
            String email = parts[2].trim();
            int age = Integer.parseInt(parts[3].trim());
            double balance = Double.parseDouble(parts[4].trim());
            
            // Validate data
            if (id.isEmpty()) throw new InvalidDataException("Line " + lineNumber + ": Customer ID cannot be empty");
            if (name.isEmpty()) throw new InvalidDataException("Line " + lineNumber + ": Customer name cannot be empty");
            if (!email.contains("@")) throw new InvalidDataException("Line " + lineNumber + ": Invalid email format");
            if (age < 0 || age > 120) throw new InvalidDataException("Line " + lineNumber + ": Invalid age: " + age);
            if (balance < 0) throw new InvalidDataException("Line " + lineNumber + ": Balance cannot be negative");
            
            return new CustomerData(id, name, email, age, balance);
        } catch (NumberFormatException e) {
            throw new InvalidDataException("Line " + lineNumber + ": Invalid number format: " + e.getMessage(), e);
        } catch (Exception e) {
            throw new DataImportException("Line " + lineNumber + ": Unexpected error: " + e.getMessage(), e);
        }
    }
    
    @Override
    public String toString() {
        return String.format("Customer{id='%s', name='%s', email='%s', age=%d, balance=$%.2f}",
            id, name, email, age, balance);
    }
}

class DataImporter {
    private List<CustomerData> importedData = new ArrayList<>();
    private List<String> errors = new ArrayList<>();
    
    public void importData(List<String> csvLines) {
        for (int i = 0; i < csvLines.size(); i++) {
            String line = csvLines.get(i);
            int lineNumber = i + 1;
            
            try {
                CustomerData customer = CustomerData.fromCsvLine(line, lineNumber);
                importedData.add(customer);
                System.out.printf("✓ Line %d: Imported %s%n", lineNumber, customer.getName());
            } catch (DataImportException e) {
                errors.add(e.getMessage());
                System.out.printf("✗ Line %d: %s%n", lineNumber, e.getMessage());
            }
        }
    }
    
    public void printSummary() {
        System.out.printf("\nImport Summary:%n");
        System.out.printf("  Successfully imported: %d records%n", importedData.size());
        System.out.printf("  Errors encountered: %d records%n", errors.size());
        
        if (!errors.isEmpty()) {
            System.out.println("\nErrors:");
            errors.forEach(error -> System.out.println("  " + error));
        }
        
        if (!importedData.isEmpty()) {
            System.out.println("\nImported Customers:");
            importedData.forEach(customer -> System.out.println("  " + customer));
        }
    }
}

public class Main {
    public static void main(String[] args) {
        List<String> csvData = Arrays.asList(
            "C001,John Doe,john@email.com,25,1500.50",
            "C002,Jane Smith,jane@email.com,30,2500.75",
            "C003,Bob Johnson,bob@email.com,invalid_age,1800.00",
            "C004,Alice Brown,alice@invalid_email,28,1200.25",
            "C005,Charlie Wilson,charlie@email.com,35,-500.00",
            "C006,Diana Davis,diana@email.com,22,3000.00,extra_field",
            "C007,Eve Miller,eve@email.com,29,2200.50"
        );
        
        DataImporter importer = new DataImporter();
        importer.importData(csvData);
        importer.printSummary();
    }
}
```
**A:** 
```
✓ Line 1: Imported John Doe
✓ Line 2: Imported Jane Smith
✗ Line 3: Line 3: Invalid number format: For input string: "invalid_age"
✗ Line 4: Line 4: Invalid email format
✗ Line 5: Line 5: Balance cannot be negative
✗ Line 6: Line 6: Expected 5 fields, got 6
✓ Line 7: Imported Eve Miller

Import Summary:
  Successfully imported: 3 records
  Errors encountered: 4 records

Errors:
  Line 3: Invalid number format: For input string: "invalid_age"
  Line 4: Invalid email format
  Line 5: Balance cannot be negative
  Line 6: Expected 5 fields, got 6

Imported Customers:
  Customer{id='C001', name='John Doe', email='john@email.com', age=25, balance=$1500.50}
  Customer{id='C002', name='Jane Smith', email='jane@email.com', age=30, balance=$2500.75}
  Customer{id='C007', name='Eve Miller', email='eve@email.com', age=29, balance=$2200.50}
```

---

### 6. Banking Operations — Transfer Processing
**Q: Handle bank transfers with rollback on failure. What is the output?**
```java
import java.util.*;

class BankingException extends Exception {
    public BankingException(String message) { super(message); }
}

class AccountNotFoundException extends BankingException {
    public AccountNotFoundException(String accountNumber) {
        super("Account not found: " + accountNumber);
    }
}

class InsufficientFundsException extends BankingException {
    public InsufficientFundsException(String accountNumber, double available, double requested) {
        super(String.format("Insufficient funds in account %s: Available $%.2f, Requested $%.2f",
            accountNumber, available, requested));
    }
}

class TransferLimitException extends BankingException {
    public TransferLimitException(double amount, double limit) {
        super(String.format("Transfer amount $%.2f exceeds limit $%.2f", amount, limit));
    }
}

class BankAccount {
    private String accountNumber;
    private String holderName;
    private double balance;
    private double dailyTransferTotal = 0.0;
    private static final double DAILY_TRANSFER_LIMIT = 10000.0;
    
    public BankAccount(String accountNumber, String holderName, double balance) {
        this.accountNumber = accountNumber;
        this.holderName = holderName;
        this.balance = balance;
    }
    
    public void withdraw(double amount) throws InsufficientFundsException, TransferLimitException {
        if (amount > balance) {
            throw new InsufficientFundsException(accountNumber, balance, amount);
        }
        if (dailyTransferTotal + amount > DAILY_TRANSFER_LIMIT) {
            throw new TransferLimitException(amount, DAILY_TRANSFER_LIMIT - dailyTransferTotal);
        }
        balance -= amount;
        dailyTransferTotal += amount;
        System.out.printf("Withdrew $%.2f from %s. New balance: $%.2f%n", 
            amount, accountNumber, balance);
    }
    
    public void deposit(double amount) {
        balance += amount;
        System.out.printf("Deposited $%.2f to %s. New balance: $%.2f%n", 
            amount, accountNumber, balance);
    }
    
    public String getAccountNumber() { return accountNumber; }
    public double getBalance() { return balance; }
    public String getHolderName() { return holderName; }
}

class TransferService {
    private Map<String, BankAccount> accounts = new HashMap<>();
    
    public void addAccount(BankAccount account) {
        accounts.put(account.getAccountNumber(), account);
    }
    
    public void transfer(String fromAccount, String toAccount, double amount) throws BankingException {
        BankAccount source = accounts.get(fromAccount);
        BankAccount destination = accounts.get(toAccount);
        
        if (source == null) {
            throw new AccountNotFoundException(fromAccount);
        }
        if (destination == null) {
            throw new AccountNotFoundException(toAccount);
        }
        
        System.out.printf("Initiating transfer: $%.2f from %s to %s%n", 
            amount, fromAccount, toAccount);
        
        try {
            // Perform transfer
            source.withdraw(amount);
            destination.deposit(amount);
            
            System.out.printf("✓ Transfer completed successfully%n");
        } catch (InsufficientFundsException | TransferLimitException e) {
            System.out.printf("✗ Transfer failed: %s%n", e.getMessage());
            throw e; // Re-throw to let caller handle
        }
    }
    
    public void printAccountBalances() {
        System.out.println("\nAccount Balances:");
        accounts.values().forEach(account -> 
            System.out.printf("  %s (%s): $%.2f%n",
                account.getAccountNumber(), account.getHolderName(), account.getBalance()));
    }
}

public class Main {
    public static void main(String[] args) {
        TransferService transferService = new TransferService();
        
        // Add accounts
        transferService.addAccount(new BankAccount("ACC001", "John Doe", 5000.0));
        transferService.addAccount(new BankAccount("ACC002", "Jane Smith", 3000.0));
        transferService.addAccount(new BankAccount("ACC003", "Bob Wilson", 1000.0));
        
        transferService.printAccountBalances();
        
        // Process transfers
        Object[][] transfers = {
            {"ACC001", "ACC002", 1000.0},  // Valid transfer
            {"ACC002", "ACC003", 3500.0},  // Insufficient funds
            {"ACC001", "ACC004", 500.0},   // Destination account not found
            {"ACC003", "ACC001", 500.0},   // Valid transfer
            {"ACC001", "ACC002", 12000.0}  // Transfer limit exceeded
        };
        
        for (Object[] transfer : transfers) {
            String from = (String) transfer[0];
            String to = (String) transfer[1];
            double amount = (double) transfer[2];
            
            System.out.printf("\n--- Processing Transfer ---%n");
            
            try {
                transferService.transfer(from, to, amount);
            } catch (BankingException e) {
                System.out.println("Transfer failed: " + e.getMessage());
            }
        }
        
        transferService.printAccountBalances();
    }
}
```
**A:** 
```
Account Balances:
  ACC001 (John Doe): $5000.00
  ACC002 (Jane Smith): $3000.00
  ACC003 (Bob Wilson): $1000.00

--- Processing Transfer ---
Initiating transfer: $1000.00 from ACC001 to ACC002
Withdrew $1000.00 from ACC001. New balance: $4000.00
Deposited $1000.00 to ACC002. New balance: $4000.00
✓ Transfer completed successfully

--- Processing Transfer ---
Initiating transfer: $3500.00 from ACC002 to ACC003
✗ Transfer failed: Insufficient funds in account ACC002: Available $4000.00, Requested $3500.00
Transfer failed: Insufficient funds in account ACC002: Available $4000.00, Requested $3500.00

--- Processing Transfer ---
Initiating transfer: $500.00 from ACC001 to ACC004
✗ Transfer failed: Account not found: ACC004
Transfer failed: Account not found: ACC004

--- Processing Transfer ---
Initiating transfer: $500.00 from ACC003 to ACC001
Withdrew $500.00 from ACC003. New balance: $500.00
Deposited $500.00 to ACC001. New balance: $4500.00
✓ Transfer completed successfully

--- Processing Transfer ---
Initiating transfer: $12000.00 from ACC001 to ACC002
✗ Transfer failed: Transfer amount $12000.00 exceeds limit $10000.00
Transfer failed: Transfer amount $12000.00 exceeds limit $10000.00

Account Balances:
  ACC001 (John Doe): $4500.00
  ACC002 (Jane Smith): $4000.00
  ACC003 (Bob Wilson): $500.00
```

---

### 7. API Response Handling — Error Recovery
**Q: Handle API responses with various error conditions. What is the output?**
```java
import java.util.*;

class ApiException extends Exception {
    private final int statusCode;
    
    public ApiException(int statusCode, String message) {
        super(message);
        this.statusCode = statusCode;
    }
    
    public int getStatusCode() { return statusCode; }
}

class ApiResponse<T> {
    private final int statusCode;
    private final T data;
    private final String errorMessage;
    
    public ApiResponse(int statusCode, T data, String errorMessage) {
        this.statusCode = statusCode;
        this.data = data;
        this.errorMessage = errorMessage;
    }
    
    public boolean isSuccess() { return statusCode >= 200 && statusCode < 300; }
    public boolean isClientError() { return statusCode >= 400 && statusCode < 500; }
    public boolean isServerError() { return statusCode >= 500 && statusCode < 600; }
    
    public T getData() throws ApiException {
        if (!isSuccess()) {
            throw new ApiException(statusCode, errorMessage != null ? errorMessage : "Request failed");
        }
        return data;
    }
    
    public int getStatusCode() { return statusCode; }
    public String getErrorMessage() { return errorMessage; }
}

class UserService {
    public ApiResponse<String> getUserById(String userId) {
        // Simulate different API responses
        return switch (userId) {
            case "user123" -> new ApiResponse<>(200, "John Doe", null);
            case "user456" -> new ApiResponse<>(404, null, "User not found");
            case "user789" -> new ApiResponse<>(401, null, "Unauthorized access");
            case "user000" -> new ApiResponse<>(500, null, "Internal server error");
            default -> new ApiResponse<>(400, null, "Invalid user ID format");
        };
    }
    
    public ApiResponse<Map<String, Object>> getUserProfile(String userId) {
        // Simulate profile API with potential errors
        if ("user123".equals(userId)) {
            Map<String, Object> profile = new HashMap<>();
            profile.put("name", "John Doe");
            profile.put("email", "john@email.com");
            profile.put("age", 30);
            return new ApiResponse<>(200, profile, null);
        } else if ("user456".equals(userId)) {
            return new ApiResponse<>(404, null, "User profile not found");
        } else {
            return new ApiResponse<>(500, null, "Database connection failed");
        }
    }
}

class ApiClient {
    private UserService userService = new UserService();
    private int retryCount = 0;
    private static final int MAX_RETRIES = 3;
    
    public void getUserInformation(String userId) {
        System.out.printf("Fetching information for user: %s%n", userId);
        
        try {
            // Get basic user info
            ApiResponse<String> userResponse = userService.getUserById(userId);
            String userName = userResponse.getData();
            
            System.out.printf("  User name: %s%n", userName);
            
            // Get user profile with retry logic
            Map<String, Object> profile = getUserProfileWithRetry(userId);
            System.out.printf("  Profile: %s%n", profile);
            
        } catch (ApiException e) {
            handleApiError(e, userId);
        } finally {
            retryCount = 0; // Reset retry count for next user
        }
    }
    
    private Map<String, Object> getUserProfileWithRetry(String userId) throws ApiException {
        while (retryCount < MAX_RETRIES) {
            try {
                ApiResponse<Map<String, Object>> profileResponse = userService.getUserProfile(userId);
                return profileResponse.getData();
            } catch (ApiException e) {
                retryCount++;
                
                if (e.getStatusCode() == 500 && retryCount < MAX_RETRIES) {
                    System.out.printf("  Server error, retrying... (attempt %d/%d)%n", 
                        retryCount, MAX_RETRIES);
                    try {
                        Thread.sleep(100); // Brief delay before retry
                    } catch (InterruptedException ie) {
                        Thread.currentThread().interrupt();
                        throw new ApiException(500, "Retry interrupted");
                    }
                } else {
                    throw e; // Re-throw non-retryable errors or after max retries
                }
            }
        }
        throw new ApiException(500, "Max retries exceeded");
    }
    
    private void handleApiError(ApiException e, String userId) {
        if (e.getStatusCode() == 404) {
            System.out.printf("  User not found: %s%n", userId);
        } else if (e.getStatusCode() == 401) {
            System.out.printf("  Access denied for user: %s%n", userId);
        } else if (e.getStatusCode() >= 500) {
            System.out.printf("  Server error while fetching user %s: %s%n", userId, e.getMessage());
        } else {
            System.out.printf("  Client error for user %s: %s%n", userId, e.getMessage());
        }
    }
}

public class Main {
    public static void main(String[] args) {
        ApiClient client = new ApiClient();
        
        String[] userIds = {"user123", "user456", "user789", "user000", "invalid_id"};
        
        for (String userId : userIds) {
            client.getUserInformation(userId);
            System.out.println();
        }
    }
}
```
**A:** 
```
Fetching information for user: user123
  User name: John Doe
  Profile: {name=John Doe, email=john@email.com, age=30}

Fetching information for user: user456
  User not found: user456

Fetching information for user: user789
  Access denied for user: user789

Fetching information for user: user000
  User not found: user000

Fetching information for user: invalid_id
  Client error for user invalid_id: Invalid user ID format
```

---

### 8. Configuration Management — Validation and Defaults
**Q: Handle configuration loading with validation and fallback values. What is the output?**
```java
import java.util.*;
import java.util.regex.*;

class ConfigurationException extends Exception {
    public ConfigurationException(String message) { super(message); }
}

class InvalidConfigValueException extends ConfigurationException {
    public InvalidConfigValueException(String key, String value, String reason) {
        super(String.format("Invalid value for '%s': '%s' - %s", key, value, reason));
    }
}

class MissingRequiredConfigException extends ConfigurationException {
    public MissingRequiredConfigException(String key) {
        super("Missing required configuration: " + key);
    }
}

class ConfigurationManager {
    private Map<String, String> config = new HashMap<>();
    private Map<String, String> defaults = new HashMap<>();
    
    public ConfigurationManager() {
        // Set default values
        defaults.put("server.port", "8080");
        defaults.put("server.timeout", "30000");
        defaults.put("database.pool.size", "10");
        defaults.put("log.level", "INFO");
        defaults.put("app.name", "MyApplication");
    }
    
    public void loadConfiguration(Map<String, String> properties) throws ConfigurationException {
        List<String> errors = new ArrayList<>();
        
        for (Map.Entry<String, String> entry : properties.entrySet()) {
            String key = entry.getKey();
            String value = entry.getValue();
            
            try {
                validateConfigValue(key, value);
                config.put(key, value);
            } catch (InvalidConfigValueException e) {
                errors.add(e.getMessage());
            }
        }
        
        if (!errors.isEmpty()) {
            throw new ConfigurationException("Configuration validation failed: " + String.join(", ", errors));
        }
    }
    
    private void validateConfigValue(String key, String value) throws InvalidConfigValueException {
        switch (key) {
            case "server.port":
                try {
                    int port = Integer.parseInt(value);
                    if (port < 1 || port > 65535) {
                        throw new InvalidConfigValueException(key, value, "Port must be between 1 and 65535");
                    }
                } catch (NumberFormatException e) {
                    throw new InvalidConfigValueException(key, value, "Must be a valid integer");
                }
                break;
                
            case "server.timeout":
                try {
                    int timeout = Integer.parseInt(value);
                    if (timeout < 1000 || timeout > 300000) {
                        throw new InvalidConfigValueException(key, value, "Timeout must be between 1000 and 300000 ms");
                    }
                } catch (NumberFormatException e) {
                    throw new InvalidConfigValueException(key, value, "Must be a valid integer");
                }
                break;
                
            case "database.pool.size":
                try {
                    int poolSize = Integer.parseInt(value);
                    if (poolSize < 1 || poolSize > 100) {
                        throw new InvalidConfigValueException(key, value, "Pool size must be between 1 and 100");
                    }
                } catch (NumberFormatException e) {
                    throw new InvalidConfigValueException(key, value, "Must be a valid integer");
                }
                break;
                
            case "log.level":
                String[] validLevels = {"DEBUG", "INFO", "WARN", "ERROR"};
                if (!Arrays.asList(validLevels).contains(value.toUpperCase())) {
                    throw new InvalidConfigValueException(key, value, "Must be one of: " + String.join(", ", validLevels));
                }
                break;
                
            case "app.name":
                if (value == null || value.trim().isEmpty()) {
                    throw new InvalidConfigValueException(key, value, "Application name cannot be empty");
                }
                if (!value.matches("^[a-zA-Z0-9_-]+$")) {
                    throw new InvalidConfigValueException(key, value, "Must contain only alphanumeric characters, underscores, and hyphens");
                }
                break;
        }
    }
    
    public String getConfigValue(String key) throws MissingRequiredConfigException {
        String value = config.get(key);
        if (value == null) {
            value = defaults.get(key);
        }
        if (value == null) {
            throw new MissingRequiredConfigException(key);
        }
        return value;
    }
    
    public int getIntConfigValue(String key) throws MissingRequiredConfigException {
        String value = getConfigValue(key);
        try {
            return Integer.parseInt(value);
        } catch (NumberFormatException e) {
            throw new MissingRequiredConfigException(key + " (invalid integer format)");
        }
    }
    
    public void printConfiguration() {
        System.out.println("Configuration Summary:");
        
        Set<String> allKeys = new HashSet<>();
        allKeys.addAll(config.keySet());
        allKeys.addAll(defaults.keySet());
        
        allKeys.stream().sorted()
            .forEach(key -> {
                String value = config.containsKey(key) ? config.get(key) : defaults.get(key);
                String source = config.containsKey(key) ? "user" : "default";
                System.out.printf("  %s = %s (%s)%n", key, value, source);
            });
    }
}

public class Main {
    public static void main(String[] args) {
        ConfigurationManager configManager = new ConfigurationManager();
        
        // Test configuration sets
        Map<String, String>[] configSets = new Map[]{
            // Valid configuration
            Map.of(
                "server.port", "9090",
                "database.pool.size", "20",
                "log.level", "DEBUG"
            ),
            // Invalid configuration
            Map.of(
                "server.port", "99999",  // Invalid port
                "server.timeout", "500",  // Too low
                "database.pool.size", "150",  // Too high
                "log.level", "VERBOSE",  // Invalid level
                "app.name", "App With Spaces"  // Invalid format
            ),
            // Mixed valid/invalid
            Map.of(
                "server.port", "8080",
                "database.pool.size", "invalid_number",  // Invalid
                "log.level", "INFO"
            )
        };
        
        for (int i = 0; i < configSets.length; i++) {
            System.out.printf("\n--- Configuration Set %d ---%n", i + 1);
            Map<String, String> configSet = configSets[i];
            
            try {
                configManager.loadConfiguration(configSet);
                configManager.printConfiguration();
                
                // Test getting specific values
                System.out.println("\nSpecific Configuration Values:");
                System.out.printf("  Server Port: %d%n", configManager.getIntConfigValue("server.port"));
                System.out.printf("  Log Level: %s%n", configManager.getConfigValue("log.level"));
                System.out.printf("  App Name: %s%n", configManager.getConfigValue("app.name"));
                
            } catch (ConfigurationException e) {
                System.out.println("Configuration Error: " + e.getMessage());
                
                // Show partial configuration that was successfully loaded
                System.out.println("\nPartially Loaded Configuration:");
                configManager.printConfiguration();
            }
            
            // Reset for next test
            configManager = new ConfigurationManager();
        }
    }
}
```
**A:** 
```
--- Configuration Set 1 ---
Configuration Summary:
  app.name = MyApplication (default)
  database.pool.size = 20 (user)
  log.level = DEBUG (user)
  server.port = 9090 (user)
  server.timeout = 30000 (default)

Specific Configuration Values:
  Server Port: 9090
  Log Level: DEBUG
  App Name: MyApplication

--- Configuration Set 2 ---
Configuration Error: Configuration validation failed: Invalid value for 'server.port': '99999' - Port must be between 1 and 65535, Invalid value for 'server.timeout': '500' - Timeout must be between 1000 and 300000 ms, Invalid value for 'database.pool.size': '150' - Pool size must be between 1 and 100, Invalid value for 'log.level': 'VERBOSE' - Must be one of: DEBUG, INFO, WARN, ERROR, Invalid value for 'app.name': 'App With Spaces' - Must contain only alphanumeric characters, underscores, and hyphens

Partially Loaded Configuration:
  app.name = MyApplication (default)
  database.pool.size = 10 (default)
  log.level = INFO (default)
  server.port = 8080 (default)
  server.timeout = 30000 (default)

--- Configuration Set 3 ---
Configuration Error: Configuration validation failed: Invalid value for 'database.pool.size': 'invalid_number' - Must be a valid integer

Partially Loaded Configuration:
  app.name = MyApplication (default)
  database.pool.size = 10 (default)
  log.level = INFO (user)
  server.port = 8080 (user)
  server.timeout = 30000 (default)
```
