# Error Handling: Practical Programs

**Goal**: Master comprehensive error handling techniques including custom exceptions, try-with-resources, and exception chaining.

## 1. Custom Exception Hierarchies

### Banking System with Custom Exceptions

```java
// Base custom exception
class BankingException extends Exception {
    private String errorCode;
    private LocalDateTime timestamp;
    
    public BankingException(String message) {
        super(message);
        this.timestamp = LocalDateTime.now();
    }
    
    public BankingException(String message, Throwable cause) {
        super(message, cause);
        this.timestamp = LocalDateTime.now();
    }
    
    public BankingException(String errorCode, String message) {
        super(message);
        this.errorCode = errorCode;
        this.timestamp = LocalDateTime.now();
    }
    
    public BankingException(String errorCode, String message, Throwable cause) {
        super(message, cause);
        this.errorCode = errorCode;
        this.timestamp = LocalDateTime.now();
    }
    
    public String getErrorCode() { return errorCode; }
    public LocalDateTime getTimestamp() { return timestamp; }
    
    @Override
    public String toString() {
        return String.format("[%s] %s at %s", 
                           errorCode != null ? errorCode : "BANK_ERR", 
                           getMessage(), 
                           timestamp);
    }
}

// Specific exception types
class InsufficientFundsException extends BankingException {
    private final double availableBalance;
    private final double requestedAmount;
    
    public InsufficientFundsException(double availableBalance, double requestedAmount) {
        super("INSUFFICIENT_FUNDS", 
              String.format("Insufficient funds. Available: $%.2f, Requested: $%.2f", 
                           availableBalance, requestedAmount));
        this.availableBalance = availableBalance;
        this.requestedAmount = requestedAmount;
    }
    
    public double getAvailableBalance() { return availableBalance; }
    public double getRequestedAmount() { return requestedAmount; }
}

class AccountNotFoundException extends BankingException {
    private final String accountNumber;
    
    public AccountNotFoundException(String accountNumber) {
        super("ACCOUNT_NOT_FOUND", "Account not found: " + accountNumber);
        this.accountNumber = accountNumber;
    }
    
    public String getAccountNumber() { return accountNumber; }
}

class InvalidTransactionException extends BankingException {
    private final String transactionType;
    private final double amount;
    
    public InvalidTransactionException(String transactionType, double amount) {
        super("INVALID_TRANSACTION", 
              String.format("Invalid %s transaction with amount: $%.2f", 
                           transactionType, amount));
        this.transactionType = transactionType;
        this.amount = amount;
    }
    
    public String getTransactionType() { return transactionType; }
    public double getAmount() { return amount; }
}

class AccountLockedException extends BankingException {
    private final String accountNumber;
    private final LocalDateTime lockTime;
    
    public AccountLockedException(String accountNumber, LocalDateTime lockTime) {
        super("ACCOUNT_LOCKED", "Account is locked: " + accountNumber);
        this.accountNumber = accountNumber;
        this.lockTime = lockTime;
    }
    
    public String getAccountNumber() { return accountNumber; }
    public LocalDateTime getLockTime() { return lockTime; }
}

class DailyLimitExceededException extends BankingException {
    private final double dailyLimit;
    private final double attemptedAmount;
    private final double currentTotal;
    
    public DailyLimitExceededException(double dailyLimit, double attemptedAmount, double currentTotal) {
        super("DAILY_LIMIT_EXCEEDED", 
              String.format("Daily limit exceeded. Limit: $%.2f, Attempted: $%.2f, Current: $%.2f", 
                           dailyLimit, attemptedAmount, currentTotal));
        this.dailyLimit = dailyLimit;
        this.attemptedAmount = attemptedAmount;
        this.currentTotal = currentTotal;
    }
    
    public double getDailyLimit() { return dailyLimit; }
    public double getAttemptedAmount() { return attemptedAmount; }
    public double getCurrentTotal() { return currentTotal; }
}

// Account class
class Account {
    private String accountNumber;
    private String accountHolder;
    private double balance;
    private boolean locked;
    private LocalDateTime lockTime;
    private double dailyWithdrawalTotal;
    private final double DAILY_LIMIT = 1000.0;
    
    public Account(String accountNumber, String accountHolder, double initialBalance) {
        this.accountNumber = accountNumber;
        this.accountHolder = accountHolder;
        this.balance = initialBalance;
        this.locked = false;
        this.dailyWithdrawalTotal = 0.0;
    }
    
    public void withdraw(double amount) throws BankingException {
        validateAccount();
        validateWithdrawal(amount);
        
        balance -= amount;
        dailyWithdrawalTotal += amount;
        
        System.out.printf("Withdrawn: $%.2f. New balance: $%.2f\n", amount, balance);
    }
    
    public void deposit(double amount) throws BankingException {
        validateAccount();
        validateDeposit(amount);
        
        balance += amount;
        System.out.printf("Deposited: $%.2f. New balance: $%.2f\n", amount, balance);
    }
    
    public void transfer(Account targetAccount, double amount) throws BankingException {
        validateAccount();
        targetAccount.validateAccount();
        
        withdraw(amount);
        targetAccount.deposit(amount);
        
        System.out.printf("Transferred $%.2f to account %s\n", amount, targetAccount.getAccountNumber());
    }
    
    public void lockAccount() {
        this.locked = true;
        this.lockTime = LocalDateTime.now();
        System.out.println("Account " + accountNumber + " has been locked");
    }
    
    public void unlockAccount() {
        this.locked = false;
        this.lockTime = null;
        System.out.println("Account " + accountNumber + " has been unlocked");
    }
    
    public void resetDailyTotal() {
        this.dailyWithdrawalTotal = 0.0;
        System.out.println("Daily withdrawal total reset for account " + accountNumber);
    }
    
    private void validateAccount() throws BankingException {
        if (locked) {
            throw new AccountLockedException(accountNumber, lockTime);
        }
    }
    
    private void validateWithdrawal(double amount) throws BankingException {
        if (amount <= 0) {
            throw new InvalidTransactionException("withdrawal", amount);
        }
        
        if (amount > balance) {
            throw new InsufficientFundsException(balance, amount);
        }
        
        if (dailyWithdrawalTotal + amount > DAILY_LIMIT) {
            throw new DailyLimitExceededException(DAILY_LIMIT, amount, dailyWithdrawalTotal);
        }
    }
    
    private void validateDeposit(double amount) throws BankingException {
        if (amount <= 0) {
            throw new InvalidTransactionException("deposit", amount);
        }
        
        if (amount > 10000) {
            throw new InvalidTransactionException("large deposit", amount);
        }
    }
    
    // Getters
    public String getAccountNumber() { return accountNumber; }
    public String getAccountHolder() { return accountHolder; }
    public double getBalance() { return balance; }
    public boolean isLocked() { return locked; }
    public double getDailyWithdrawalTotal() { return dailyWithdrawalTotal; }
}

// Bank service
class BankService {
    private Map<String, Account> accounts = new HashMap<>();
    
    public void createAccount(String accountNumber, String accountHolder, double initialBalance) {
        accounts.put(accountNumber, new Account(accountNumber, accountHolder, initialBalance));
        System.out.println("Created account: " + accountNumber);
    }
    
    public Account getAccount(String accountNumber) throws AccountNotFoundException {
        Account account = accounts.get(accountNumber);
        if (account == null) {
            throw new AccountNotFoundException(accountNumber);
        }
        return account;
    }
    
    public void performTransaction(String accountNumber, String transactionType, double amount) {
        try {
            Account account = getAccount(accountNumber);
            
            switch (transactionType.toLowerCase()) {
                case "withdraw":
                    account.withdraw(amount);
                    break;
                case "deposit":
                    account.deposit(amount);
                    break;
                default:
                    throw new InvalidTransactionException(transactionType, amount);
            }
            
        } catch (BankingException e) {
            handleBankingException(e);
        }
    }
    
    public void transferFunds(String fromAccount, String toAccount, double amount) {
        try {
            Account source = getAccount(fromAccount);
            Account target = getAccount(toAccount);
            
            source.transfer(target, amount);
            
        } catch (BankingException e) {
            handleBankingException(e);
        }
    }
    
    private void handleBankingException(BankingException e) {
        System.err.println("Transaction failed: " + e);
        
        // Log to file (simplified)
        logError(e);
        
        // Additional handling based on exception type
        if (e instanceof InsufficientFundsException) {
            InsufficientFundsException ife = (InsufficientFundsException) e;
            System.err.printf("Suggestion: Deposit at least $%.2f to cover this transaction\n", 
                            ife.getRequestedAmount() - ife.getAvailableBalance());
        } else if (e instanceof AccountLockedException) {
            System.err.println("Please contact customer service to unlock your account");
        } else if (e instanceof DailyLimitExceededException) {
            DailyLimitExceededException dlee = (DailyLimitExceededException) e;
            System.err.printf("Daily limit: $%.2f. Available: $%.2f\n", 
                            dlee.getDailyLimit(), 
                            dlee.getDailyLimit() - dlee.getCurrentTotal());
        }
    }
    
    private void logError(BankingException e) {
        // In real application, this would log to a file or database
        System.err.println("[LOG] " + e.toString());
    }
    
    public void displayAccountInfo(String accountNumber) {
        try {
            Account account = getAccount(accountNumber);
            System.out.println("\n=== Account Information ===");
            System.out.println("Account Number: " + account.getAccountNumber());
            System.out.println("Account Holder: " + account.getAccountHolder());
            System.out.printf("Balance: $%.2f\n", account.getBalance());
            System.out.println("Status: " + (account.isLocked() ? "Locked" : "Active"));
            System.out.printf("Daily Withdrawal Total: $%.2f\n", account.getDailyWithdrawalTotal());
            System.out.println();
        } catch (AccountNotFoundException e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}

public class CustomExceptionDemo {
    public static void main(String[] args) {
        BankService bank = new BankService();
        
        // Create accounts
        bank.createAccount("123456", "John Doe", 500.0);
        bank.createAccount("789012", "Jane Smith", 1000.0);
        bank.createAccount("345678", "Bob Wilson", 200.0);
        
        // Display account info
        bank.displayAccountInfo("123456");
        bank.displayAccountInfo("789012");
        
        // Test normal transactions
        System.out.println("=== Normal Transactions ===");
        bank.performTransaction("123456", "deposit", 200.0);
        bank.performTransaction("123456", "withdraw", 100.0);
        
        // Test insufficient funds
        System.out.println("\n=== Testing Insufficient Funds ===");
        bank.performTransaction("345678", "withdraw", 300.0);
        
        // Test invalid transaction
        System.out.println("\n=== Testing Invalid Transaction ===");
        bank.performTransaction("123456", "deposit", -50.0);
        
        // Test account not found
        System.out.println("\n=== Testing Account Not Found ===");
        bank.performTransaction("999999", "withdraw", 100.0);
        
        // Test daily limit
        System.out.println("\n=== Testing Daily Limit ===");
        bank.performTransaction("789012", "withdraw", 600.0);
        bank.performTransaction("789012", "withdraw", 500.0);
        
        // Test locked account
        System.out.println("\n=== Testing Locked Account ===");
        Account account;
        try {
            account = bank.getAccount("123456");
            account.lockAccount();
            bank.performTransaction("123456", "withdraw", 50.0);
        } catch (AccountNotFoundException e) {
            System.err.println("Error: " + e.getMessage());
        }
        
        // Test transfer
        System.out.println("\n=== Testing Transfer ===");
        try {
            account = bank.getAccount("123456");
            account.unlockAccount();
            bank.transferFunds("123456", "789012", 150.0);
        } catch (AccountNotFoundException e) {
            System.err.println("Error: " + e.getMessage());
        }
        
        // Display final account info
        bank.displayAccountInfo("123456");
        bank.displayAccountInfo("789012");
    }
}
```

## 2. Try-with-Resources

### File Processing with Automatic Resource Management

```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

// Custom AutoCloseable resources
class DatabaseConnection implements AutoCloseable {
    private String connectionName;
    private boolean connected;
    
    public DatabaseConnection(String connectionName) {
        this.connectionName = connectionName;
        this.connected = true;
        System.out.println("Database connection opened: " + connectionName);
    }
    
    public void executeQuery(String query) {
        if (!connected) {
            throw new IllegalStateException("Connection is closed");
        }
        System.out.println("Executing query: " + query);
    }
    
    public void executeUpdate(String update) {
        if (!connected) {
            throw new IllegalStateException("Connection is closed");
        }
        System.out.println("Executing update: " + update);
    }
    
    @Override
    public void close() {
        if (connected) {
            connected = false;
            System.out.println("Database connection closed: " + connectionName);
        }
    }
}

class NetworkConnection implements AutoCloseable {
    private String endpoint;
    private boolean connected;
    
    public NetworkConnection(String endpoint) {
        this.endpoint = endpoint;
        this.connected = true;
        System.out.println("Network connection opened to: " + endpoint);
    }
    
    public void sendData(String data) {
        if (!connected) {
            throw new IllegalStateException("Connection is closed");
        }
        System.out.println("Sending data to " + endpoint + ": " + data);
    }
    
    public String receiveData() {
        if (!connected) {
            throw new IllegalStateException("Connection is closed");
        }
        return "Response from " + endpoint;
    }
    
    @Override
    public void close() {
        if (connected) {
            connected = false;
            System.out.println("Network connection closed to: " + endpoint);
        }
    }
}

class FileProcessor implements AutoCloseable {
    private BufferedReader reader;
    private BufferedWriter writer;
    private String fileName;
    
    public FileProcessor(String inputFileName, String outputFileName) throws IOException {
        this.fileName = inputFileName;
        this.reader = new BufferedReader(new FileReader(inputFileName));
        this.writer = new BufferedWriter(new FileWriter(outputFileName));
        System.out.println("File processor created for: " + inputFileName);
    }
    
    public void processFile() throws IOException {
        String line;
        int lineNumber = 1;
        
        while ((line = reader.readLine()) != null) {
            // Process each line (example: convert to uppercase and add line number)
            String processedLine = lineNumber + ": " + line.toUpperCase();
            writer.write(processedLine);
            writer.newLine();
            lineNumber++;
        }
        
        System.out.println("File processing completed");
    }
    
    public void copyWithFilter(String filter) throws IOException {
        String line;
        int copiedLines = 0;
        
        while ((line = reader.readLine()) != null) {
            if (line.toLowerCase().contains(filter.toLowerCase())) {
                writer.write(line);
                writer.newLine();
                copiedLines++;
            }
        }
        
        System.out.println("Copied " + copiedLines + " lines containing filter: " + filter);
    }
    
    @Override
    public void close() throws IOException {
        if (reader != null) {
            reader.close();
            System.out.println("Input file closed");
        }
        if (writer != null) {
            writer.close();
            System.out.println("Output file closed");
        }
    }
}

public class TryWithResourcesDemo {
    
    public static void main(String[] args) {
        // 1. Basic try-with-resources with single resource
        System.out.println("=== Single Resource Example ===");
        singleResourceExample();
        
        // 2. Multiple resources
        System.out.println("\n=== Multiple Resources Example ===");
        multipleResourcesExample();
        
        // 3. Nested try-with-resources
        System.out.println("\n=== Nested Resources Example ===");
        nestedResourcesExample();
        
        // 4. Exception handling with resources
        System.out.println("\n=== Exception Handling Example ===");
        exceptionHandlingExample();
        
        // 5. File processing examples
        System.out.println("\n=== File Processing Examples ===");
        fileProcessingExamples();
        
        // 6. Custom resource management
        System.out.println("\n=== Custom Resource Management ===");
        customResourceExample();
    }
    
    private static void singleResourceExample() {
        // Traditional way (before Java 7)
        System.out.println("--- Traditional way ---");
        BufferedReader reader = null;
        try {
            reader = new BufferedReader(new StringReader("Hello World\nJava Programming"));
            String line = reader.readLine();
            System.out.println("Read: " + line);
        } catch (IOException e) {
            System.err.println("Error reading: " + e.getMessage());
        } finally {
            if (reader != null) {
                try {
                    reader.close();
                } catch (IOException e) {
                    System.err.println("Error closing reader: " + e.getMessage());
                }
            }
        }
        
        // With try-with-resources (Java 7+)
        System.out.println("--- Try-with-resources way ---");
        try (BufferedReader autoReader = new BufferedReader(new StringReader("Hello World\nJava Programming"))) {
            String line = autoReader.readLine();
            System.out.println("Read: " + line);
        } catch (IOException e) {
            System.err.println("Error reading: " + e.getMessage());
        }
        // No need for finally block - resource is automatically closed
    }
    
    private static void multipleResourcesExample() {
        try (DatabaseConnection db = new DatabaseConnection("main_db");
             NetworkConnection network = new NetworkConnection("api.example.com")) {
            
            db.executeQuery("SELECT * FROM users");
            network.sendData("GET /users");
            String response = network.receiveData();
            System.out.println("Network response: " + response);
            
        } catch (Exception e) {
            System.err.println("Error in multiple resources: " + e.getMessage());
        }
        // Both connections are automatically closed in reverse order
    }
    
    private static void nestedResourcesExample() {
        try (DatabaseConnection outerDb = new DatabaseConnection("outer_db")) {
            outerDb.executeQuery("SELECT COUNT(*) FROM orders");
            
            try (DatabaseConnection innerDb = new DatabaseConnection("inner_db")) {
                innerDb.executeQuery("SELECT * FROM customers");
                
                try (NetworkConnection network = new NetworkConnection("internal.api")) {
                    network.sendData("Process data");
                }
            }
        } catch (Exception e) {
            System.err.println("Error in nested resources: " + e.getMessage());
        }
    }
    
    private static void exceptionHandlingExample() {
        try (DatabaseConnection db = new DatabaseConnection("test_db")) {
            // Simulate an exception
            db.executeQuery("SELECT * FROM non_existent_table");
        } catch (Exception e) {
            System.err.println("Caught exception: " + e.getMessage());
            // Resource is still closed automatically
        }
        
        // Resource is closed even when exception occurs
        System.out.println("Exception handling completed - resource should be closed");
    }
    
    private static void fileProcessingExamples() {
        // Create sample input file
        String sampleContent = "Java is a programming language\n" +
                             "Python is also popular\n" +
                             "JavaScript runs in browsers\n" +
                             "C++ is for systems programming\n" +
                             "Go is a modern language";
        
        try {
            // Write sample content to file
            Files.write(Paths.get("input.txt"), sampleContent.getBytes());
            
            // Example 1: Basic file processing
            System.out.println("--- Basic File Processing ---");
            try (FileProcessor processor = new FileProcessor("input.txt", "output.txt")) {
                processor.processFile();
            }
            
            // Example 2: Filtered copying
            System.out.println("--- Filtered File Copying ---");
            try (FileProcessor processor = new FileProcessor("input.txt", "filtered.txt")) {
                processor.copyWithFilter("Java");
            }
            
            // Example 3: Stream processing with try-with-resources
            System.out.println("--- Stream Processing ---");
            try (Stream<String> lines = Files.lines(Paths.get("input.txt"))) {
                long javaCount = lines.filter(line -> line.contains("Java")).count();
                System.out.println("Lines containing 'Java': " + javaCount);
            }
            
            // Example 4: Multiple file operations
            System.out.println("--- Multiple File Operations ---");
            try (BufferedWriter writer = new BufferedWriter(new FileWriter("summary.txt"));
                 BufferedReader reader = new BufferedReader(new FileReader("input.txt"))) {
                
                String line;
                int lineCount = 0;
                int wordCount = 0;
                
                while ((line = reader.readLine()) != null) {
                    lineCount++;
                    wordCount += line.split("\\s+").length;
                }
                
                writer.write("File Summary:\n");
                writer.write("Lines: " + lineCount + "\n");
                writer.write("Words: " + wordCount + "\n");
                
                System.out.println("Summary file created");
            }
            
        } catch (IOException e) {
            System.err.println("File processing error: " + e.getMessage());
        } finally {
            // Cleanup files
            cleanupFiles();
        }
    }
    
    private static void customResourceExample() {
        // Example with custom AutoCloseable resources
        try (DatabaseConnection primaryDb = new DatabaseConnection("primary");
             DatabaseConnection replicaDb = new DatabaseConnection("replica");
             NetworkConnection apiClient = new NetworkConnection("rest.api")) {
            
            // Simulate complex business logic
            primaryDb.executeQuery("BEGIN TRANSACTION");
            replicaDb.executeQuery("SELECT * FROM users WHERE active = true");
            
            apiClient.sendData("POST /sync");
            String response = apiClient.receiveData();
            
            primaryDb.executeUpdate("UPDATE users SET last_sync = NOW()");
            
            System.out.println("Complex operation completed successfully");
            
        } catch (Exception e) {
            System.err.println("Complex operation failed: " + e.getMessage());
        }
    }
    
    private static void cleanupFiles() {
        try {
            Files.deleteIfExists(Paths.get("input.txt"));
            Files.deleteIfExists(Paths.get("output.txt"));
            Files.deleteIfExists(Paths.get("filtered.txt"));
            Files.deleteIfExists(Paths.get("summary.txt"));
        } catch (IOException e) {
            System.err.println("Error cleaning up files: " + e.getMessage());
        }
    }
}
```

## 3. Exception Chaining

### Multi-Layer Exception Handling

```java
import java.sql.*;
import java.io.*;
import java.time.LocalDateTime;

// Layer-specific exceptions
class DataAccessException extends Exception {
    private String operation;
    private LocalDateTime timestamp;
    
    public DataAccessException(String operation, String message) {
        super(message);
        this.operation = operation;
        this.timestamp = LocalDateTime.now();
    }
    
    public DataAccessException(String operation, String message, Throwable cause) {
        super(message, cause);
        this.operation = operation;
        this.timestamp = LocalDateTime.now();
    }
    
    public String getOperation() { return operation; }
    public LocalDateTime getTimestamp() { return timestamp; }
}

class ValidationException extends Exception {
    private String fieldName;
    private Object fieldValue;
    
    public ValidationException(String fieldName, Object fieldValue, String message) {
        super(message);
        this.fieldName = fieldName;
        this.fieldValue = fieldValue;
    }
    
    public ValidationException(String fieldName, Object fieldValue, String message, Throwable cause) {
        super(message, cause);
        this.fieldName = fieldName;
        this.fieldValue = fieldValue;
    }
    
    public String getFieldName() { return fieldName; }
    public Object getFieldValue() { return fieldValue; }
}

class BusinessLogicException extends Exception {
    private String businessRule;
    private Map<String, Object> context;
    
    public BusinessLogicException(String businessRule, String message) {
        super(message);
        this.businessRule = businessRule;
        this.context = new HashMap<>();
    }
    
    public BusinessLogicException(String businessRule, String message, Throwable cause) {
        super(message, cause);
        this.businessRule = businessRule;
        this.context = new HashMap<>();
    }
    
    public void addContext(String key, Object value) {
        context.put(key, value);
    }
    
    public String getBusinessRule() { return businessRule; }
    public Map<String, Object> getContext() { return context; }
}

// Service layer with exception chaining
class UserService {
    private UserRepository userRepository;
    private EmailService emailService;
    
    public UserService() {
        this.userRepository = new UserRepository();
        this.emailService = new EmailService();
    }
    
    public User createUser(String username, String email, int age) throws BusinessLogicException {
        try {
            // Validate input
            validateUserData(username, email, age);
            
            // Check if user already exists
            if (userRepository.existsByUsername(username)) {
                throw new BusinessLogicException("UNIQUE_USERNAME", 
                    "Username already exists: " + username);
            }
            
            // Create user
            User user = new User(username, email, age);
            
            // Save to database
            User savedUser = userRepository.save(user);
            
            // Send welcome email
            try {
                emailService.sendWelcomeEmail(savedUser);
            } catch (EmailException e) {
                // Email failure shouldn't prevent user creation
                System.err.println("Warning: Failed to send welcome email: " + e.getMessage());
                // Chain the exception but don't fail the operation
                throw new BusinessLogicException("EMAIL_SEND_FAILED", 
                    "User created but welcome email failed", e);
            }
            
            return savedUser;
            
        } catch (ValidationException e) {
            throw new BusinessLogicException("VALIDATION_FAILED", 
                "User validation failed", e);
        } catch (DataAccessException e) {
            throw new BusinessLogicException("DATA_ACCESS_FAILED", 
                "Failed to save user to database", e);
        }
    }
    
    public User updateUserProfile(String username, Map<String, Object> updates) throws BusinessLogicException {
        try {
            // Get existing user
            User user = userRepository.findByUsername(username);
            if (user == null) {
                throw new BusinessLogicException("USER_NOT_FOUND", 
                    "User not found: " + username);
            }
            
            // Apply updates
            applyUpdates(user, updates);
            
            // Validate updated user
            validateUserData(user.getUsername(), user.getEmail(), user.getAge());
            
            // Save changes
            User updatedUser = userRepository.update(user);
            
            // Send notification email
            try {
                emailService.sendProfileUpdateEmail(updatedUser);
            } catch (EmailException e) {
                // Log but don't fail the operation
                System.err.println("Warning: Failed to send update notification: " + e.getMessage());
            }
            
            return updatedUser;
            
        } catch (ValidationException e) {
            throw new BusinessLogicException("VALIDATION_FAILED", 
                "Updated user data is invalid", e);
        } catch (DataAccessException e) {
            throw new BusinessLogicException("DATA_ACCESS_FAILED", 
                "Failed to update user in database", e);
        }
    }
    
    private void validateUserData(String username, String email, int age) throws ValidationException {
        if (username == null || username.trim().isEmpty()) {
            throw new ValidationException("username", username, "Username cannot be empty");
        }
        if (username.length() < 3) {
            throw new ValidationException("username", username, "Username must be at least 3 characters");
        }
        
        if (email == null || !email.contains("@")) {
            throw new ValidationException("email", email, "Invalid email format");
        }
        
        if (age < 18 || age > 120) {
            throw new ValidationException("age", age, "Age must be between 18 and 120");
        }
    }
    
    private void applyUpdates(User user, Map<String, Object> updates) throws ValidationException {
        for (Map.Entry<String, Object> entry : updates.entrySet()) {
            String field = entry.getKey();
            Object value = entry.getValue();
            
            switch (field.toLowerCase()) {
                case "email":
                    if (value instanceof String) {
                        String newEmail = (String) value;
                        if (!newEmail.contains("@")) {
                            throw new ValidationException("email", newEmail, "Invalid email format");
                        }
                        user.setEmail(newEmail);
                    }
                    break;
                case "age":
                    if (value instanceof Integer) {
                        int newAge = (Integer) value;
                        if (newAge < 18 || newAge > 120) {
                            throw new ValidationException("age", newAge, "Age must be between 18 and 120");
                        }
                        user.setAge(newAge);
                    }
                    break;
                default:
                    throw new ValidationException("field", field, "Unknown field: " + field);
            }
        }
    }
}

// Repository layer
class UserRepository {
    private Map<String, User> users = new HashMap<>();
    
    public User save(User user) throws DataAccessException {
        try {
            // Simulate database operation
            if (Math.random() < 0.1) { // 10% chance of failure
                throw new SQLException("Database connection lost");
            }
            
            users.put(user.getUsername(), user);
            return user;
            
        } catch (SQLException e) {
            throw new DataAccessException("SAVE_USER", "Failed to save user", e);
        }
    }
    
    public User update(User user) throws DataAccessException {
        try {
            // Simulate database operation
            if (Math.random() < 0.1) { // 10% chance of failure
                throw new SQLException("Database timeout");
            }
            
            users.put(user.getUsername(), user);
            return user;
            
        } catch (SQLException e) {
            throw new DataAccessException("UPDATE_USER", "Failed to update user", e);
        }
    }
    
    public User findByUsername(String username) throws DataAccessException {
        try {
            // Simulate database operation
            if (Math.random() < 0.1) { // 10% chance of failure
                throw new SQLException("Database connection lost");
            }
            
            return users.get(username);
            
        } catch (SQLException e) {
            throw new DataAccessException("FIND_USER", "Failed to find user", e);
        }
    }
    
    public boolean existsByUsername(String username) throws DataAccessException {
        try {
            // Simulate database operation
            if (Math.random() < 0.1) { // 10% chance of failure
                throw new SQLException("Database connection lost");
            }
            
            return users.containsKey(username);
            
        } catch (SQLException e) {
            throw new DataAccessException("CHECK_USER_EXISTS", "Failed to check if user exists", e);
        }
    }
}

// Email service
class EmailService {
    public void sendWelcomeEmail(User user) throws EmailException {
        try {
            // Simulate email sending
            if (Math.random() < 0.2) { // 20% chance of failure
                throw new IOException("SMTP server not responding");
            }
            
            System.out.println("Welcome email sent to: " + user.getEmail());
            
        } catch (IOException e) {
            throw new EmailException("WELCOME_EMAIL", "Failed to send welcome email", e);
        }
    }
    
    public void sendProfileUpdateEmail(User user) throws EmailException {
        try {
            // Simulate email sending
            if (Math.random() < 0.2) { // 20% chance of failure
                throw new IOException("Email service unavailable");
            }
            
            System.out.println("Profile update email sent to: " + user.getEmail());
            
        } catch (IOException e) {
            throw new EmailException("UPDATE_EMAIL", "Failed to send update email", e);
        }
    }
}

class EmailException extends Exception {
    private String emailType;
    
    public EmailException(String emailType, String message) {
        super(message);
        this.emailType = emailType;
    }
    
    public EmailException(String emailType, String message, Throwable cause) {
        super(message, cause);
        this.emailType = emailType;
    }
    
    public String getEmailType() { return emailType; }
}

// User model
class User {
    private String username;
    private String email;
    private int age;
    
    public User(String username, String email, int age) {
        this.username = username;
        this.email = email;
        this.age = age;
    }
    
    // Getters and setters
    public String getUsername() { return username; }
    public void setUsername(String username) { this.username = username; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public int getAge() { return age; }
    public void setAge(int age) { this.age = age; }
}

// Exception analyzer
class ExceptionAnalyzer {
    public static void analyzeException(Exception e) {
        System.out.println("\n=== Exception Analysis ===");
        System.out.println("Root Exception: " + e.getClass().getSimpleName());
        System.out.println("Message: " + e.getMessage());
        
        // Analyze exception chain
        Throwable cause = e.getCause();
        int depth = 1;
        while (cause != null) {
            System.out.println("Cause " + depth + ": " + cause.getClass().getSimpleName());
            System.out.println("  Message: " + cause.getMessage());
            cause = cause.getCause();
            depth++;
        }
        
        // Print stack trace
        System.out.println("\nStack Trace:");
        e.printStackTrace();
    }
    
    public static void extractExceptionContext(Exception e) {
        System.out.println("\n=== Exception Context ===");
        
        if (e instanceof BusinessLogicException) {
            BusinessLogicException ble = (BusinessLogicException) e;
            System.out.println("Business Rule: " + ble.getBusinessRule());
            System.out.println("Context: " + ble.getContext());
        }
        
        if (e instanceof ValidationException) {
            ValidationException ve = (ValidationException) e;
            System.out.println("Field: " + ve.getFieldName());
            System.out.println("Value: " + ve.getFieldValue());
        }
        
        if (e instanceof DataAccessException) {
            DataAccessException dae = (DataAccessException) e;
            System.out.println("Operation: " + dae.getOperation());
            System.out.println("Timestamp: " + dae.getTimestamp());
        }
    }
}

public class ExceptionChainingDemo {
    public static void main(String[] args) {
        UserService userService = new UserService();
        
        // Test 1: Successful user creation
        System.out.println("=== Test 1: Successful User Creation ===");
        try {
            User user = userService.createUser("john_doe", "john@example.com", 25);
            System.out.println("User created successfully: " + user.getUsername());
        } catch (BusinessLogicException e) {
            ExceptionAnalyzer.analyzeException(e);
            ExceptionAnalyzer.extractExceptionContext(e);
        }
        
        // Test 2: Validation failure
        System.out.println("\n=== Test 2: Validation Failure ===");
        try {
            User user = userService.createUser("ab", "invalid-email", 15);
        } catch (BusinessLogicException e) {
            ExceptionAnalyzer.analyzeException(e);
            ExceptionAnalyzer.extractExceptionContext(e);
        }
        
        // Test 3: Duplicate username
        System.out.println("\n=== Test 3: Duplicate Username ===");
        try {
            User user = userService.createUser("john_doe", "john2@example.com", 30);
        } catch (BusinessLogicException e) {
            ExceptionAnalyzer.analyzeException(e);
            ExceptionAnalyzer.extractExceptionContext(e);
        }
        
        // Test 4: Database failure (simulate)
        System.out.println("\n=== Test 4: Database Failure Simulation ===");
        for (int i = 0; i < 20; i++) { // Try multiple times to trigger failure
            try {
                User user = userService.createUser("user" + i, "user" + i + "@example.com", 25);
                System.out.println("Created user: " + user.getUsername());
                break;
            } catch (BusinessLogicException e) {
                if (e.getCause() instanceof DataAccessException) {
                    System.out.println("Database operation failed: " + e.getMessage());
                    ExceptionAnalyzer.analyzeException(e);
                    break;
                }
            }
        }
        
        // Test 5: User update with validation error
        System.out.println("\n=== Test 5: User Update with Validation Error ===");
        try {
            Map<String, Object> updates = new HashMap<>();
            updates.put("email", "invalid-email");
            updates.put("age", 150);
            
            User updated = userService.updateUserProfile("john_doe", updates);
        } catch (BusinessLogicException e) {
            ExceptionAnalyzer.analyzeException(e);
            ExceptionAnalyzer.extractExceptionContext(e);
        }
        
        // Test 6: User not found
        System.out.println("\n=== Test 6: User Not Found ===");
        try {
            Map<String, Object> updates = new HashMap<>();
            updates.put("email", "newemail@example.com");
            
            User updated = userService.updateUserProfile("nonexistent_user", updates);
        } catch (BusinessLogicException e) {
            ExceptionAnalyzer.analyzeException(e);
            ExceptionAnalyzer.extractExceptionContext(e);
        }
    }
}
```

## Practice Exercises

1. **Custom Exceptions**: Create a library management system with comprehensive exception handling
2. **Try-with-resources**: Build a file backup utility that handles multiple resources safely
3. **Exception Chaining**: Implement a multi-layer web service with proper exception propagation

## Interview Questions

1. What's the difference between checked and unchecked exceptions?
2. When should you create custom exceptions instead of using built-in ones?
3. How does try-with-resources improve resource management?
4. What is exception chaining and when is it useful?
5. How do you decide whether to catch or propagate an exception?
