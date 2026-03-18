# Testing Frameworks: Practical Programs

**Goal**: Master comprehensive testing with JUnit, Mockito, integration testing, and Test-Driven Development.

## Prerequisites

Add these dependencies to your `pom.xml`:

```xml
<dependencies>
    <!-- JUnit 5 -->
    <dependency>
        <groupId>org.junit.jupiter</groupId>
        <artifactId>junit-jupiter</artifactId>
        <version>5.10.0</version>
        <scope>test</scope>
    </dependency>
    
    <!-- Mockito -->
    <dependency>
        <groupId>org.mockito</groupId>
        <artifactId>mockito-core</artifactId>
        <version>5.5.0</version>
        <scope>test</scope>
    </dependency>
    <dependency>
        <groupId>org.mockito</groupId>
        <artifactId>mockito-junit-jupiter</artifactId>
        <version>5.5.0</version>
        <scope>test</scope>
    </dependency>
    
    <!-- Spring Boot Test (for integration testing) -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <version>3.2.0</version>
        <scope>test</scope>
    </dependency>
    
    <!-- AssertJ for fluent assertions -->
    <dependency>
        <groupId>org.assertj</groupId>
        <artifactId>assertj-core</artifactId>
        <version>3.24.2</version>
        <scope>test</scope>
    </dependency>
    
    <!-- TestContainers for integration testing -->
    <dependency>
        <groupId>org.testcontainers</groupId>
        <artifactId>junit-jupiter</artifactId>
        <version>1.19.0</version>
        <scope>test</scope>
    </dependency>
    <dependency>
        <groupId>org.testcontainers</groupId>
        <artifactId>mysql</artifactId>
        <version>1.19.0</version>
        <scope>test</scope>
    </dependency>
</dependencies>
```

## 1. JUnit 5 Comprehensive Testing

### Advanced JUnit Features

```java
import org.junit.jupiter.api.*;
import org.junit.jupiter.params.*;
import org.junit.jupiter.params.provider.*;
import static org.junit.jupiter.api.*;
import static org.assertj.core.api.Assertions.*;

import java.time.*;
import java.util.*;
import java.util.stream.*;

// Class to be tested
class Calculator {
    private double result;
    private List<String> history = new ArrayList<>();
    
    public double add(double a, double b) {
        result = a + b;
        history.add(a + " + " + b + " = " + result);
        return result;
    }
    
    public double subtract(double a, double b) {
        result = a - b;
        history.add(a + " - " + b + " = " + result);
        return result;
    }
    
    public double multiply(double a, double b) {
        result = a * b;
        history.add(a + " * " + b + " = " + result);
        return result;
    }
    
    public double divide(double a, double b) {
        if (b == 0) {
            throw new ArithmeticException("Division by zero");
        }
        result = a / b;
        history.add(a + " / " + b + " = " + result);
        return result;
    }
    
    public double getResult() {
        return result;
    }
    
    public List<String> getHistory() {
        return new ArrayList<>(history);
    }
    
    public void clear() {
        result = 0;
        history.clear();
    }
    
    public double power(double base, double exponent) {
        result = Math.pow(base, exponent);
        history.add(base + " ^ " + exponent + " = " + result);
        return result;
    }
    
    public double sqrt(double number) {
        if (number < 0) {
            throw new IllegalArgumentException("Cannot calculate square root of negative number");
        }
        result = Math.sqrt(number);
        history.add("√" + number + " = " + result);
        return result;
    }
}

// Email service for testing
class EmailService {
    private List<String> sentEmails = new ArrayList<>();
    
    public void sendEmail(String to, String subject, String body) {
        if (to == null || to.trim().isEmpty()) {
            throw new IllegalArgumentException("Recipient cannot be null or empty");
        }
        if (subject == null || subject.trim().isEmpty()) {
            throw new IllegalArgumentException("Subject cannot be null or empty");
        }
        
        String email = "To: " + to + "\nSubject: " + subject + "\nBody: " + body;
        sentEmails.add(email);
        System.out.println("Email sent: " + email);
    }
    
    public List<String> getSentEmails() {
        return new ArrayList<>(sentEmails);
    }
    
    public void clearEmails() {
        sentEmails.clear();
    }
    
    public int getEmailCount() {
        return sentEmails.size();
    }
}

// User service for testing
class UserService {
    private Map<String, User> users = new HashMap<>();
    private EmailService emailService;
    
    public UserService(EmailService emailService) {
        this.emailService = emailService;
    }
    
    public User registerUser(String username, String email, int age) {
        if (username == null || username.trim().isEmpty()) {
            throw new IllegalArgumentException("Username cannot be null or empty");
        }
        if (users.containsKey(username)) {
            throw new IllegalArgumentException("Username already exists");
        }
        if (age < 18 || age > 120) {
            throw new IllegalArgumentException("Age must be between 18 and 120");
        }
        if (!email.contains("@")) {
            throw new IllegalArgumentException("Invalid email format");
        }
        
        User user = new User(username, email, age);
        users.put(username, user);
        
        // Send welcome email
        emailService.sendEmail(email, "Welcome!", "Welcome to our platform, " + username + "!");
        
        return user;
    }
    
    public User getUser(String username) {
        return users.get(username);
    }
    
    public List<User> getAllUsers() {
        return new ArrayList<>(users.values());
    }
    
    public boolean deleteUser(String username) {
        User user = users.remove(username);
        if (user != null) {
            emailService.sendEmail(user.getEmail(), "Account Deleted", 
                "Your account has been deleted successfully.");
            return true;
        }
        return false;
    }
    
    public int getUserCount() {
        return users.size();
    }
}

class User {
    private String username;
    private String email;
    private int age;
    private LocalDateTime registrationDate;
    
    public User(String username, String email, int age) {
        this.username = username;
        this.email = email;
        this.age = age;
        this.registrationDate = LocalDateTime.now();
    }
    
    // Getters
    public String getUsername() { return username; }
    public String getEmail() { return email; }
    public int getAge() { return age; }
    public LocalDateTime getRegistrationDate() { return registrationDate; }
    
    @Override
    public String toString() {
        return "User{username='" + username + "', email='" + email + "', age=" + age + "}";
    }
}

// Comprehensive JUnit test class
@DisplayName("Calculator Tests")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class CalculatorTest {
    
    private Calculator calculator;
    
    @BeforeEach
    void setUp() {
        calculator = new Calculator();
        System.out.println("Setting up calculator");
    }
    
    @AfterEach
    void tearDown() {
        calculator.clear();
        System.out.println("Tearing down calculator");
    }
    
    @Test
    @DisplayName("Should add two numbers correctly")
    @Order(1)
    void testAddition() {
        // Given
        double a = 5.0;
        double b = 3.0;
        
        // When
        double result = calculator.add(a, b);
        
        // Then
        assertEquals(8.0, result, "5 + 3 should equal 8");
        assertEquals(8.0, calculator.getResult(), "Calculator result should be updated");
    }
    
    @Test
    @DisplayName("Should subtract two numbers correctly")
    @Order(2)
    void testSubtraction() {
        double result = calculator.subtract(10.0, 4.0);
        assertEquals(6.0, result, "10 - 4 should equal 6");
    }
    
    @Test
    @DisplayName("Should multiply two numbers correctly")
    @Order(3)
    void testMultiplication() {
        double result = calculator.multiply(6.0, 7.0);
        assertEquals(42.0, result, "6 * 7 should equal 42");
    }
    
    @Test
    @DisplayName("Should divide two numbers correctly")
    @Order(4)
    void testDivision() {
        double result = calculator.divide(15.0, 3.0);
        assertEquals(5.0, result, "15 / 3 should equal 5");
    }
    
    @Test
    @DisplayName("Should throw exception when dividing by zero")
    @Order(5)
    void testDivisionByZero() {
        ArithmeticException exception = assertThrows(ArithmeticException.class, () -> {
            calculator.divide(10.0, 0.0);
        }, "Should throw ArithmeticException for division by zero");
        
        assertEquals("Division by zero", exception.getMessage());
    }
    
    @Test
    @DisplayName("Should calculate power correctly")
    @Order(6)
    void testPower() {
        double result = calculator.power(2.0, 3.0);
        assertEquals(8.0, result, 0.001, "2^3 should equal 8");
    }
    
    @Test
    @DisplayName("Should calculate square root correctly")
    @Order(7)
    void testSquareRoot() {
        double result = calculator.sqrt(16.0);
        assertEquals(4.0, result, 0.001, "√16 should equal 4");
    }
    
    @Test
    @DisplayName("Should throw exception for square root of negative number")
    @Order(8)
    void testSquareRootOfNegative() {
        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () -> {
            calculator.sqrt(-4.0);
        }, "Should throw IllegalArgumentException for negative input");
        
        assertEquals("Cannot calculate square root of negative number", exception.getMessage());
    }
    
    @Test
    @DisplayName("Should maintain calculation history")
    @Order(9)
    void testCalculationHistory() {
        // Perform multiple calculations
        calculator.add(2.0, 3.0);
        calculator.multiply(5.0, 6.0);
        calculator.divide(10.0, 2.0);
        
        List<String> history = calculator.getHistory();
        assertEquals(3, history.size(), "Should have 3 calculations in history");
        assertTrue(history.get(0).contains("2.0 + 3.0"));
        assertTrue(history.get(1).contains("5.0 * 6.0"));
        assertTrue(history.get(2).contains("10.0 / 2.0"));
    }
    
    @ParameterizedTest
    @DisplayName("Should perform addition with various inputs")
    @Order(10)
    @ValueSource(doubles = {0.0, 1.0, -1.0, 100.0, -100.0, 0.5, -0.5})
    void testAdditionWithVariousInputs(double value) {
        double result = calculator.add(value, value);
        assertEquals(value * 2, result, 0.001, 
                    value + " + " + value + " should equal " + (value * 2));
    }
    
    @ParameterizedTest
    @DisplayName("Should validate email formats")
    @Order(11)
    @CsvSource({
        "test@example.com, true",
        "user.name@domain.co.uk, true",
        "invalid-email, false",
        "@domain.com, false",
        "user@, false",
        "user@domain, false",
        "user.name@domain.com, true"
    })
    void testEmailValidation(String email, boolean expected) {
        boolean isValid = email != null && email.contains("@") && email.contains(".");
        assertEquals(expected, isValid, "Email validation failed for: " + email);
    }
    
    @ParameterizedTest
    @DisplayName("Should handle age validation")
    @Order(12)
    @MethodSource("ageProvider")
    void testAgeValidation(int age, boolean expected) {
        boolean isValid = age >= 18 && age <= 120;
        assertEquals(expected, isValid, "Age validation failed for age: " + age);
    }
    
    static Stream<Arguments> ageProvider() {
        return Stream.of(
            Arguments.of(17, false),
            Arguments.of(18, true),
            Arguments.of(25, true),
            Arguments.of(120, true),
            Arguments.of(121, false),
            Arguments.of(0, false),
            Arguments.of(-5, false)
        );
    }
    
    @Test
    @DisplayName("Should demonstrate nested tests")
    @Order(13)
    void nestedTest() {
        // Test basic operations
        assertAll("Basic Operations",
            () -> assertEquals(8.0, calculator.add(5.0, 3.0)),
            () -> assertEquals(2.0, calculator.subtract(10.0, 8.0)),
            () -> assertEquals(15.0, calculator.multiply(3.0, 5.0)),
            () -> assertEquals(4.0, calculator.divide(12.0, 3.0))
        );
        
        // Test edge cases
        assertAll("Edge Cases",
            () -> assertEquals(0.0, calculator.add(0.0, 0.0)),
            () -> assertEquals(0.0, calculator.multiply(0.0, 100.0)),
            () -> assertEquals(1.0, calculator.divide(5.0, 5.0))
        );
    }
    
    @Test
    @DisplayName("Should demonstrate timeout testing")
    @Order(14)
    void testTimeout() {
        assertTimeout(Duration.ofMillis(100), () -> {
            // Simulate some work
            Thread.sleep(50);
            calculator.add(1.0, 2.0);
        }, "Calculation should complete within 100ms");
    }
    
    @RepeatedTest(5)
    @DisplayName("Should produce consistent results across multiple runs")
    @Order(15)
    void testConsistentResults(RepetitionInfo repetitionInfo) {
        double result = calculator.add(10.0, 20.0);
        assertEquals(30.0, result, "Result should be consistent across repetitions");
        System.out.println("Repetition " + repetitionInfo.getCurrentRepetition() + 
                          " of " + repetitionInfo.getTotalRepetitions());
    }
    
    @Test
    @DisplayName("Should demonstrate custom assertions")
    @Order(16)
    void testCustomAssertions() {
        calculator.add(5.0, 3.0);
        calculator.multiply(2.0, 4.0);
        
        List<String> history = calculator.getHistory();
        
        // Custom assertions using AssertJ
        assertThat(history)
            .hasSize(2)
            .containsExactly("5.0 + 3.0 = 8.0", "2.0 * 4.0 = 8.0")
            .allSatisfy(entry -> assertThat(entry).contains("="));
        
        assertThat(calculator.getResult()).isEqualTo(8.0);
    }
}

// UserService test class
@DisplayName("User Service Tests")
class UserServiceTest {
    
    private UserService userService;
    private EmailService emailService;
    
    @BeforeEach
    void setUp() {
        emailService = new EmailService();
        userService = new UserService(emailService);
    }
    
    @Test
    @DisplayName("Should register user successfully")
    void testRegisterUser() {
        User user = userService.registerUser("john_doe", "john@example.com", 25);
        
        assertNotNull(user);
        assertEquals("john_doe", user.getUsername());
        assertEquals("john@example.com", user.getEmail());
        assertEquals(25, user.getAge());
        assertNotNull(user.getRegistrationDate());
        
        // Verify email was sent
        assertEquals(1, emailService.getEmailCount());
        List<String> emails = emailService.getSentEmails();
        assertTrue(emails.get(0).contains("john@example.com"));
        assertTrue(emails.get(0).contains("Welcome"));
    }
    
    @Test
    @DisplayName("Should reject duplicate username")
    void testDuplicateUsername() {
        userService.registerUser("john_doe", "john@example.com", 25);
        
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser("john_doe", "john2@example.com", 30);
        }, "Should throw exception for duplicate username");
    }
    
    @Test
    @DisplayName("Should validate age constraints")
    void testAgeValidation() {
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser("young_user", "young@example.com", 17);
        }, "Should reject users under 18");
        
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser("old_user", "old@example.com", 121);
        }, "Should reject users over 120");
    }
    
    @Test
    @DisplayName("Should validate email format")
    void testEmailValidation() {
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser("user", "invalid-email", 25);
        }, "Should reject invalid email format");
    }
    
    @Test
    @DisplayName("Should handle user deletion")
    void testDeleteUser() {
        User user = userService.registerUser("test_user", "test@example.com", 25);
        assertEquals(1, userService.getUserCount());
        
        boolean deleted = userService.deleteUser("test_user");
        assertTrue(deleted);
        assertEquals(0, userService.getUserCount());
        
        // Verify goodbye email was sent
        assertEquals(2, emailService.getEmailCount()); // Welcome + goodbye
    }
}

public class JUnitDemo {
    public static void main(String[] args) {
        System.out.println("This is a test class. Run tests using your IDE or Maven:");
        System.out.println("mvn test");
        System.out.println("Or run individual test classes from your IDE.");
    }
}
```

## 2. Mockito Mocking Framework

### Advanced Mocking Techniques

```java
import org.junit.jupiter.api.*;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.*;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.invocation.InvocationOnMock;
import org.mockito.stubbing.Answer;

import java.util.*;
import static org.mockito.Mockito.*;
import static org.assertj.core.api.Assertions.*;

// Data repository interface
interface UserRepository {
    Optional<User> findByUsername(String username);
    List<User> findAll();
    User save(User user);
    void delete(String username);
    boolean existsByUsername(String username);
    List<User> findByAgeGreaterThan(int age);
}

// Notification service interface
interface NotificationService {
    void sendWelcomeEmail(User user);
    void sendPasswordResetEmail(User user);
    void sendAccountDeletionEmail(User user);
    boolean sendEmail(String to, String subject, String body);
    int getSentEmailCount();
}

// Payment service interface
interface PaymentService {
    boolean processPayment(User user, double amount);
    List<Payment> getPaymentHistory(User user);
    double getTotalSpent(User user);
}

class Payment {
    private String id;
    private double amount;
    private LocalDateTime timestamp;
    
    public Payment(String id, double amount) {
        this.id = id;
        this.amount = amount;
        this.timestamp = LocalDateTime.now();
    }
    
    public String getId() { return id; }
    public double getAmount() { return amount; }
    public LocalDateTime getTimestamp() { return timestamp; }
}

// Advanced user service with dependencies
class AdvancedUserService {
    private UserRepository userRepository;
    private NotificationService notificationService;
    private PaymentService paymentService;
    
    public AdvancedUserService(UserRepository userRepository, 
                             NotificationService notificationService,
                             PaymentService paymentService) {
        this.userRepository = userRepository;
        this.notificationService = notificationService;
        this.paymentService = paymentService;
    }
    
    public User registerUser(String username, String email, int age) {
        // Validate input
        if (username == null || username.trim().isEmpty()) {
            throw new IllegalArgumentException("Username cannot be null or empty");
        }
        if (userRepository.existsByUsername(username)) {
            throw new IllegalArgumentException("Username already exists");
        }
        if (age < 18) {
            throw new IllegalArgumentException("User must be at least 18 years old");
        }
        
        // Create user
        User user = new User(username, email, age);
        
        // Save user
        User savedUser = userRepository.save(user);
        
        // Send welcome notification
        notificationService.sendWelcomeEmail(savedUser);
        
        return savedUser;
    }
    
    public boolean deleteUser(String username) {
        Optional<User> userOpt = userRepository.findByUsername(username);
        if (!userOpt.isPresent()) {
            return false;
        }
        
        User user = userOpt.get();
        
        // Check if user has outstanding payments
        double totalSpent = paymentService.getTotalSpent(user);
        if (totalSpent > 0) {
            // Process refund
            paymentService.processPayment(user, -totalSpent);
        }
        
        // Send deletion notification
        notificationService.sendAccountDeletionEmail(user);
        
        // Delete user
        userRepository.delete(username);
        
        return true;
    }
    
    public List<User> findAdultUsers() {
        return userRepository.findByAgeGreaterThan(17);
    }
    
    public String getUserSummary(String username) {
        Optional<User> userOpt = userRepository.findByUsername(username);
        if (!userOpt.isPresent()) {
            return "User not found";
        }
        
        User user = userOpt.get();
        double totalSpent = paymentService.getTotalSpent(user);
        List<Payment> payments = paymentService.getPaymentHistory(user);
        
        return String.format("User: %s, Email: %s, Age: %d, Total Spent: $%.2f, Payments: %d",
                           user.getUsername(), user.getEmail(), user.getAge(), totalSpent, payments.size());
    }
    
    public boolean sendPromotionalEmail(String username, String promotion) {
        Optional<User> userOpt = userRepository.findByUsername(username);
        if (!userOpt.isPresent()) {
            return false;
        }
        
        User user = userOpt.get();
        String subject = "Special Promotion: " + promotion;
        String body = "Dear " + user.getUsername() + ", check out our special promotion: " + promotion;
        
        return notificationService.sendEmail(user.getEmail(), subject, body);
    }
    
    public List<User> getActiveUsers() {
        List<User> allUsers = userRepository.findAll();
        List<User> activeUsers = new ArrayList<>();
        
        for (User user : allUsers) {
            double totalSpent = paymentService.getTotalSpent(user);
            if (totalSpent > 100) { // Consider users with >$100 spent as active
                activeUsers.add(user);
            }
        }
        
        return activeUsers;
    }
}

// Mockito test class
@ExtendWith(MockitoExtension.class)
@DisplayName("Advanced User Service Tests with Mockito")
class AdvancedUserServiceTest {
    
    @Mock
    private UserRepository userRepository;
    
    @Mock
    private NotificationService notificationService;
    
    @Mock
    private PaymentService paymentService;
    
    @InjectMocks
    private AdvancedUserService userService;
    
    @Test
    @DisplayName("Should register user successfully with mocked dependencies")
    void testRegisterUser() {
        // Given
        String username = "john_doe";
        String email = "john@example.com";
        int age = 25;
        User expectedUser = new User(username, email, age);
        
        // Configure mocks
        when(userRepository.existsByUsername(username)).thenReturn(false);
        when(userRepository.save(any(User.class))).thenReturn(expectedUser);
        when(notificationService.sendEmail(anyString(), anyString(), anyString())).thenReturn(true);
        
        // When
        User result = userService.registerUser(username, email, age);
        
        // Then
        assertThat(result).isNotNull();
        assertThat(result.getUsername()).isEqualTo(username);
        assertThat(result.getEmail()).isEqualTo(email);
        assertThat(result.getAge()).isEqualTo(age);
        
        // Verify interactions
        verify(userRepository).existsByUsername(username);
        verify(userRepository).save(any(User.class));
        verify(notificationService).sendEmail(eq(email), contains("Welcome"), anyString());
    }
    
    @Test
    @DisplayName("Should reject duplicate username")
    void testDuplicateUsername() {
        // Given
        String username = "existing_user";
        when(userRepository.existsByUsername(username)).thenReturn(true);
        
        // When & Then
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser(username, "test@example.com", 25);
        });
        
        // Verify that save was never called
        verify(userRepository, never()).save(any(User.class));
        verify(notificationService, never()).sendEmail(anyString(), anyString(), anyString());
    }
    
    @Test
    @DisplayName("Should delete user with refund")
    void testDeleteUser() {
        // Given
        String username = "test_user";
        User user = new User(username, "test@example.com", 30);
        double totalSpent = 150.0;
        
        when(userRepository.findByUsername(username)).thenReturn(Optional.of(user));
        when(paymentService.getTotalSpent(user)).thenReturn(totalSpent);
        when(paymentService.processPayment(user, -totalSpent)).thenReturn(true);
        when(notificationService.sendEmail(anyString(), anyString(), anyString())).thenReturn(true);
        
        // When
        boolean result = userService.deleteUser(username);
        
        // Then
        assertThat(result).isTrue();
        
        // Verify interactions in order
        InOrder inOrder = inOrder(paymentService, notificationService, userRepository);
        inOrder.verify(paymentService).getTotalSpent(user);
        inOrder.verify(paymentService).processPayment(user, -totalSpent);
        inOrder.verify(notificationService).sendEmail(eq(user.getEmail()), contains("Account Deleted"), anyString());
        inOrder.verify(userRepository).delete(username);
    }
    
    @Test
    @DisplayName("Should handle user not found during deletion")
    void testDeleteNonExistentUser() {
        // Given
        String username = "nonexistent_user";
        when(userRepository.findByUsername(username)).thenReturn(Optional.empty());
        
        // When
        boolean result = userService.deleteUser(username);
        
        // Then
        assertThat(result).isFalse();
        
        // Verify no further interactions
        verify(paymentService, never()).getTotalSpent(any());
        verify(notificationService, never()).sendEmail(anyString(), anyString(), anyString());
        verify(userRepository, never()).delete(anyString());
    }
    
    @Test
    @DisplayName("Should get user summary with payment history")
    void testGetUserSummary() {
        // Given
        String username = "test_user";
        User user = new User(username, "test@example.com", 30);
        double totalSpent = 250.75;
        List<Payment> payments = Arrays.asList(
            new Payment("p1", 100.0),
            new Payment("p2", 150.75)
        );
        
        when(userRepository.findByUsername(username)).thenReturn(Optional.of(user));
        when(paymentService.getTotalSpent(user)).thenReturn(totalSpent);
        when(paymentService.getPaymentHistory(user)).thenReturn(payments);
        
        // When
        String summary = userService.getUserSummary(username);
        
        // Then
        assertThat(summary).contains(username);
        assertThat(summary).contains("250.75");
        assertThat(summary).contains("2");
        
        verify(userRepository).findByUsername(username);
        verify(paymentService).getTotalSpent(user);
        verify(paymentService).getPaymentHistory(user);
    }
    
    @Test
    @DisplayName("Should return 'User not found' for non-existent user")
    void testGetUserSummaryNotFound() {
        // Given
        String username = "nonexistent_user";
        when(userRepository.findByUsername(username)).thenReturn(Optional.empty());
        
        // When
        String summary = userService.getUserSummary(username);
        
        // Then
        assertThat(summary).isEqualTo("User not found");
        
        verify(userRepository).findByUsername(username);
        verifyNoInteractions(paymentService);
    }
    
    @Test
    @DisplayName("Should find adult users")
    void testFindAdultUsers() {
        // Given
        List<User> adultUsers = Arrays.asList(
            new User("adult1", "adult1@example.com", 25),
            new User("adult2", "adult2@example.com", 35)
        );
        
        when(userRepository.findByAgeGreaterThan(17)).thenReturn(adultUsers);
        
        // When
        List<User> result = userService.findAdultUsers();
        
        // Then
        assertThat(result).hasSize(2);
        assertThat(result).extracting(User::getAge).allMatch(age -> age > 17);
        
        verify(userRepository).findByAgeGreaterThan(17);
    }
    
    @Test
    @DisplayName("Should send promotional email")
    void testSendPromotionalEmail() {
        // Given
        String username = "test_user";
        User user = new User(username, "test@example.com", 30);
        String promotion = "50% OFF";
        
        when(userRepository.findByUsername(username)).thenReturn(Optional.of(user));
        when(notificationService.sendEmail(anyString(), anyString(), anyString())).thenReturn(true);
        
        // When
        boolean result = userService.sendPromotionalEmail(username, promotion);
        
        // Then
        assertThat(result).isTrue();
        
        verify(notificationService).sendEmail(
            eq(user.getEmail()),
            eq("Special Promotion: " + promotion),
            contains(promotion)
        );
    }
    
    @Test
    @DisplayName("Should handle promotional email failure")
    void testSendPromotionalEmailFailure() {
        // Given
        String username = "test_user";
        User user = new User(username, "test@example.com", 30);
        
        when(userRepository.findByUsername(username)).thenReturn(Optional.of(user));
        when(notificationService.sendEmail(anyString(), anyString(), anyString())).thenReturn(false);
        
        // When
        boolean result = userService.sendPromotionalEmail(username, "Test Promotion");
        
        // Then
        assertThat(result).isFalse();
    }
    
    @Test
    @DisplayName("Should get active users based on spending")
    void testGetActiveUsers() {
        // Given
        User activeUser1 = new User("active1", "active1@example.com", 30);
        User activeUser2 = new User("active2", "active2@example.com", 25);
        User inactiveUser = new User("inactive", "inactive@example.com", 35);
        
        List<User> allUsers = Arrays.asList(activeUser1, inactiveUser, activeUser2);
        
        when(userRepository.findAll()).thenReturn(allUsers);
        when(paymentService.getTotalSpent(activeUser1)).thenReturn(150.0);
        when(paymentService.getTotalSpent(activeUser2)).thenReturn(200.0);
        when(paymentService.getTotalSpent(inactiveUser)).thenReturn(50.0);
        
        // When
        List<User> activeUsers = userService.getActiveUsers();
        
        // Then
        assertThat(activeUsers).hasSize(2);
        assertThat(activeUsers).extracting(User::getUsername).containsExactlyInAnyOrder("active1", "active2");
        
        // Verify all users were checked
        verify(paymentService).getTotalSpent(activeUser1);
        verify(paymentService).getTotalSpent(activeUser2);
        verify(paymentService).getTotalSpent(inactiveUser);
    }
    
    @Test
    @DisplayName("Should demonstrate advanced mocking with Answer")
    void testAdvancedMockingWithAnswer() {
        // Given
        when(userRepository.save(any(User.class))).thenAnswer(new Answer<User>() {
            private int idCounter = 1;
            
            @Override
            public User answer(InvocationOnMock invocation) throws Throwable {
                User user = (User) invocation.getArgument(0);
                // Simulate ID assignment
                return user; // In real scenario, would return user with ID
            }
        });
        
        when(notificationService.sendEmail(anyString(), anyString(), anyString())))
            .thenAnswer(invocation -> {
                String to = invocation.getArgument(0);
                String subject = invocation.getArgument(1);
                System.out.println("Mock email sent to: " + to + " with subject: " + subject);
                return true;
            });
        
        // When
        User user = userService.registerUser("test_user", "test@example.com", 25);
        
        // Then
        assertThat(user).isNotNull();
        verify(userRepository).save(any(User.class));
        verify(notificationService).sendEmail(anyString(), anyString(), anyString());
    }
    
    @Test
    @DisplayName("Should verify no more interactions")
    void testVerifyNoMoreInteractions() {
        // Given
        String username = "test_user";
        when(userRepository.findByUsername(username)).thenReturn(Optional.empty());
        
        // When
        String summary = userService.getUserSummary(username);
        
        // Then
        assertThat(summary).isEqualTo("User not found");
        
        // Verify only the expected interaction occurred
        verify(userRepository).findByUsername(username);
        verifyNoInteractions(paymentService);
        verifyNoMoreInteractions(userRepository);
    }
    
    @Test
    @DisplayName("Should demonstrate argument captor")
    void testArgumentCaptor() {
        // Given
        ArgumentCaptor<User> userCaptor = ArgumentCaptor.forClass(User.class);
        ArgumentCaptor<String> emailCaptor = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> subjectCaptor = ArgumentCaptor.forClass(String.class);
        
        when(userRepository.existsByUsername(anyString())).thenReturn(false);
        when(userRepository.save(any(User.class))).thenReturn(new User("test", "test@example.com", 25));
        when(notificationService.sendEmail(anyString(), anyString(), anyString())).thenReturn(true);
        
        // When
        userService.registerUser("test_user", "test@example.com", 25);
        
        // Then
        verify(userRepository).save(userCaptor.capture());
        verify(notificationService).sendEmail(emailCaptor.capture(), subjectCaptor.capture(), anyString());
        
        User capturedUser = userCaptor.getValue();
        assertThat(capturedUser.getUsername()).isEqualTo("test_user");
        
        String capturedEmail = emailCaptor.getValue();
        assertThat(capturedEmail).isEqualTo("test@example.com");
        
        String capturedSubject = subjectCaptor.getValue();
        assertThat(capturedSubject).contains("Welcome");
    }
    
    @Test
    @DisplayName("Should demonstrate spy functionality")
    void testSpyFunctionality() {
        // Given
        List<User> userList = new ArrayList<>();
        List<User> spyList = spy(userList);
        
        // Configure spy behavior
        when(spyList.size()).thenReturn(100); // Mock size method
        doReturn(true).when(spyList).add(any(User.class)); // Mock add method
        
        // When
        int size = spyList.size();
        boolean added = spyList.add(new User("test", "test@example.com", 25));
        
        // Then
        assertThat(size).isEqualTo(100); // Mocked value
        assertThat(added).isTrue(); // Mocked value
        
        // Verify real method was called for add
        verify(spyList).add(any(User.class));
    }
}

public class MockitoDemo {
    public static void main(String[] args) {
        System.out.println("This is a Mockito test class. Run tests using your IDE or Maven:");
        System.out.println("mvn test");
        System.out.println("Or run individual test classes from your IDE.");
    }
}
```

## 3. Integration Testing

### Spring Boot Integration Tests

```java
import org.springframework.boot.test.context.*;
import org.springframework.boot.test.web.client.*;
import org.springframework.boot.test.autoconfigure.web.servlet.*;
import org.springframework.boot.test.autoconfigure.orm.jpa.*;
import org.springframework.test.context.*;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.*;
import org.springframework.test.web.servlet.*;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

import org.junit.jupiter.api.*;
import org.testcontainers.containers.*;
import org.testcontainers.junit.jupiter.*;
import org.testcontainers.utility.*;

// Test configuration
@TestConfiguration
static class TestConfig {
    @Bean
    @Primary
    public EmailService mockEmailService() {
        return new EmailService();
    }
}

// Integration test class
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@TestPropertySource(properties = {
    "spring.datasource.url=jdbc:h2:mem:testdb",
    "spring.jpa.hibernate.ddl-auto=create-drop",
    "spring.datasource.driver-class-name=org.h2.Driver"
})
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
@DisplayName("User Service Integration Tests")
class UserServiceIntegrationTest {
    
    @Autowired
    private UserService userService;
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private EmailService emailService;
    
    @BeforeEach
    void setUp() {
        // Clean up database before each test
        userRepository.deleteAll();
        emailService.clearEmails();
    }
    
    @Test
    @Order(1)
    @DisplayName("Should register user and send welcome email")
    @Transactional
    void testUserRegistrationIntegration() {
        // When
        User user = userService.registerUser("integration_user", "integration@test.com", 25);
        
        // Then
        assertThat(user).isNotNull();
        assertThat(user.getUsername()).isEqualTo("integration_user");
        
        // Verify user is saved in database
        User savedUser = userRepository.findByUsername("integration_user").orElse(null);
        assertThat(savedUser).isNotNull();
        assertThat(savedUser.getEmail()).isEqualTo("integration@test.com");
        
        // Verify email was sent
        assertThat(emailService.getEmailCount()).isEqualTo(1);
        List<String> emails = emailService.getSentEmails();
        assertThat(emails.get(0)).contains("integration@test.com");
        assertThat(emails.get(0)).contains("Welcome");
    }
    
    @Test
    @Order(2)
    @DisplayName("Should handle user deletion with cascade")
    @Transactional
    void testUserDeletionIntegration() {
        // Given
        User user = userService.registerUser("delete_user", "delete@test.com", 30);
        assertThat(userRepository.count()).isEqualTo(1);
        
        // When
        boolean deleted = userService.deleteUser("delete_user");
        
        // Then
        assertThat(deleted).isTrue();
        assertThat(userRepository.count()).isEqualTo(0);
        
        // Verify goodbye email was sent
        assertThat(emailService.getEmailCount()).isEqualTo(2); // Welcome + goodbye
    }
    
    @Test
    @Order(3)
    @DisplayName("Should handle concurrent user registration")
    void testConcurrentUserRegistration() throws InterruptedException {
        int threadCount = 10;
        CountDownLatch latch = new CountDownLatch(threadCount);
        List<Exception> exceptions = Collections.synchronizedList(new ArrayList<>());
        
        // Create multiple threads trying to register the same user
        for (int i = 0; i < threadCount; i++) {
            new Thread(() -> {
                try {
                    userService.registerUser("concurrent_user", "concurrent@test.com", 25);
                } catch (Exception e) {
                    exceptions.add(e);
                } finally {
                    latch.countDown();
                }
            }).start();
        }
        
        latch.await(5, java.util.concurrent.TimeUnit.SECONDS);
        
        // Only one user should be registered
        assertThat(userRepository.count()).isEqualTo(1);
        
        // All but one thread should have thrown exceptions
        assertThat(exceptions).hasSize(threadCount - 1);
        exceptions.forEach(e -> assertThat(e).isInstanceOf(IllegalArgumentException.class));
    }
}

// REST Controller Integration Test
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@AutoConfigureTestDatabase(replace = AutoConfigureTestDatabase.Replace.NONE)
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
@DisplayName("User Controller Integration Tests")
class UserControllerIntegrationTest {
    
    @Autowired
    private TestRestTemplate restTemplate;
    
    @LocalServerPort
    private int port;
    
    private String baseUrl;
    
    @BeforeEach
    void setUp() {
        baseUrl = "http://localhost:" + port + "/api/users";
    }
    
    @Test
    @Order(1)
    @DisplayName("Should create user via REST API")
    void testCreateUserViaAPI() {
        // Given
        Map<String, Object> userRequest = new HashMap<>();
        userRequest.put("username", "api_user");
        userRequest.put("email", "api@test.com");
        userRequest.put("age", 25);
        
        // When
        ResponseEntity<Map> response = restTemplate.postForEntity(baseUrl, userRequest, Map.class);
        
        // Then
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.CREATED);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().get("username")).isEqualTo("api_user");
        assertThat(response.getBody().get("email")).isEqualTo("api@test.com");
    }
    
    @Test
    @Order(2)
    @DisplayName("Should get user via REST API")
    void testGetUserViaAPI() {
        // First create a user
        testCreateUserViaAPI();
        
        // When
        ResponseEntity<Map> response = restTemplate.getForEntity(baseUrl + "/api_user", Map.class);
        
        // Then
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().get("username")).isEqualTo("api_user");
    }
    
    @Test
    @Order(3)
    @DisplayName("Should handle validation errors via REST API")
    void testValidationErrorsViaAPI() {
        // Given
        Map<String, Object> invalidRequest = new HashMap<>();
        invalidRequest.put("username", ""); // Invalid
        invalidRequest.put("email", "invalid-email"); // Invalid
        invalidRequest.put("age", 15); // Invalid
        
        // When
        ResponseEntity<Map> response = restTemplate.postForEntity(baseUrl, invalidRequest, Map.class);
        
        // Then
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.BAD_REQUEST);
        assertThat(response.getBody()).isNotNull();
    }
}

// TestContainers Integration Test
@Testcontainers
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@DisplayName("Database Integration Tests with TestContainers")
class DatabaseIntegrationTest {
    
    @Container
    static MySQLContainer<?> mysql = new MySQLContainer<>("mysql:8.0")
            .withDatabaseName("testdb")
            .withUsername("test")
            .withPassword("test")
            .withReuse(true);
    
    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", mysql::getJdbcUrl);
        registry.add("spring.datasource.username", mysql::getUsername);
        registry.add("spring.datasource.password", mysql::getPassword);
        registry.add("spring.datasource.driver-class-name", () -> "com.mysql.cj.jdbc.Driver");
    }
    
    @Autowired
    private UserService userService;
    
    @Autowired
    private UserRepository userRepository;
    
    @Test
    @DisplayName("Should work with real MySQL database")
    void testRealDatabaseOperations() {
        // Given
        String username = "mysql_user";
        
        // When
        User user = userService.registerUser(username, "mysql@test.com", 30);
        
        // Then
        assertThat(user).isNotNull();
        assertThat(userRepository.findByUsername(username)).isPresent();
    }
    
    @Test
    @DisplayName("Should handle database constraints")
    void testDatabaseConstraints() {
        // Given
        userService.registerUser("constraint_user", "constraint@test.com", 25);
        
        // When & Then
        assertThrows(IllegalArgumentException.class, () -> {
            userService.registerUser("constraint_user", "constraint2@test.com", 30);
        });
    }
}

public class IntegrationTestingDemo {
    public static void main(String[] args) {
        System.out.println("This is an integration testing demo class.");
        System.out.println("Run integration tests using:");
        System.out.println("mvn test -Dtest=*IntegrationTest");
    }
}
```

## 4. Test-Driven Development (TDD)

### TDD Example: String Calculator

```java
import org.junit.jupiter.api.*;
import org.junit.jupiter.params.*;
import org.junit.jupiter.params.provider.*;
import static org.junit.jupiter.api.*;
import static org.assertj.core.api.Assertions.*;

// TDD Example: String Calculator
class StringCalculator {
    
    public int add(String numbers) {
        if (numbers == null || numbers.isEmpty()) {
            return 0;
        }
        
        // Handle custom delimiters
        String delimiter = ",|\n";
        if (numbers.startsWith("//")) {
            int delimiterIndex = numbers.indexOf('\n');
            if (delimiterIndex > 0) {
                delimiter = numbers.substring(2, delimiterIndex);
                numbers = numbers.substring(delimiterIndex + 1);
            }
        }
        
        // Split and parse numbers
        String[] numberArray = numbers.split(delimiter);
        int sum = 0;
        List<String> negatives = new ArrayList<>();
        
        for (String numStr : numberArray) {
            if (!numStr.trim().isEmpty()) {
                int num = Integer.parseInt(numStr.trim());
                if (num < 0) {
                    negatives.add(String.valueOf(num));
                } else if (num <= 1000) { // Ignore numbers > 1000
                    sum += num;
                }
            }
        }
        
        // Check for negatives
        if (!negatives.isEmpty()) {
            throw new IllegalArgumentException("Negative numbers not allowed: " + String.join(", ", negatives));
        }
        
        return sum;
    }
}

// TDD Test Class - Following Red-Green-Refactor cycle
@DisplayName("String Calculator TDD")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class StringCalculatorTest {
    
    private StringCalculator calculator;
    
    @BeforeEach
    void setUp() {
        calculator = new StringCalculator();
    }
    
    // Step 1: Empty string should return 0
    @Test
    @Order(1)
    @DisplayName("Step 1: Empty string should return 0")
    void testEmptyString() {
        // Red: Test doesn't exist yet
        // Green: Implement minimal code to pass
        assertEquals(0, calculator.add(""));
        assertEquals(0, calculator.add(null));
    }
    
    // Step 2: Single number should return that number
    @Test
    @Order(2)
    @DisplayName("Step 2: Single number should return that number")
    void testSingleNumber() {
        assertEquals(1, calculator.add("1"));
        assertEquals(5, calculator.add("5"));
        assertEquals(42, calculator.add("42"));
    }
    
    // Step 3: Two numbers should return their sum
    @Test
    @Order(3)
    @DisplayName("Step 3: Two numbers should return their sum")
    void testTwoNumbers() {
        assertEquals(3, calculator.add("1,2"));
        assertEquals(10, calculator.add("4,6"));
        assertEquals(0, calculator.add("0,0"));
    }
    
    // Step 4: Multiple numbers should return their sum
    @Test
    @Order(4)
    @DisplayName("Step 4: Multiple numbers should return their sum")
    void testMultipleNumbers() {
        assertEquals(6, calculator.add("1,2,3"));
        assertEquals(10, calculator.add("1,2,3,4"));
        assertEquals(55, calculator.add("1,2,3,4,5,6,7,8,9,10"));
    }
    
    // Step 5: Handle new lines between numbers
    @Test
    @Order(5)
    @DisplayName("Step 5: Handle new lines between numbers")
    void testNewLines() {
        assertEquals(6, calculator.add("1\n2,3"));
        assertEquals(10, calculator.add("1,2\n3,4"));
        assertEquals(3, calculator.add("1\n2"));
    }
    
    // Step 6: Support custom delimiters
    @Test
    @Order(6)
    @DisplayName("Step 6: Support custom delimiters")
    void testCustomDelimiters() {
        assertEquals(3, calculator.add("//;\n1;2"));
        assertEquals(6, calculator.add("//|\n1|2|3"));
        assertEquals(10, calculator.add("//@\n1@2@3@4"));
    }
    
    // Step 7: Handle negative numbers
    @Test
    @Order(7)
    @DisplayName("Step 7: Handle negative numbers")
    void testNegativeNumbers() {
        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () -> {
            calculator.add("1,-2,3");
        });
        assertEquals("Negative numbers not allowed: -2", exception.getMessage());
        
        exception = assertThrows(IllegalArgumentException.class, () -> {
            calculator.add("-1,-2,-3");
        });
        assertEquals("Negative numbers not allowed: -1, -2, -3", exception.getMessage());
    }
    
    // Step 8: Ignore numbers greater than 1000
    @Test
    @Order(8)
    @DisplayName("Step 8: Ignore numbers greater than 1000")
    void testLargeNumbers() {
        assertEquals(2, calculator.add("2,1001"));
        assertEquals(1002, calculator.add("1000,1001,2"));
        assertEquals(0, calculator.add("1001,1002,1003"));
    }
    
    // Step 9: Handle multiple character delimiters
    @Test
    @Order(9)
    @DisplayName("Step 9: Handle multiple character delimiters")
    void testMultipleCharacterDelimiters() {
        assertEquals(6, calculator.add("//***\n1***2***3"));
        assertEquals(10, calculator.add("//@@@\n1@@@2@@@3@@@4"));
    }
    
    // Step 10: Handle multiple delimiters
    @Test
    @Order(10)
    @DisplayName("Step 10: Handle multiple delimiters")
    void testMultipleDelimiters() {
        assertEquals(6, calculator.add("//;|\n1;2|3"));
        assertEquals(10, calculator.add("//[*][%]\n1*2%3%4"));
        assertEquals(15, calculator.add("//[**][%%]\n1**2%%3**4%%5"));
    }
    
    // Edge cases
    @Test
    @DisplayName("Should handle edge cases")
    void testEdgeCases() {
        // Empty numbers in sequence
        assertEquals(6, calculator.add("1,,2,3"));
        assertEquals(6, calculator.add("1,\n2,3"));
        
        // Only valid numbers
        assertEquals(0, calculator.add(""));
        assertEquals(0, calculator.add("   "));
    }
    
    @ParameterizedTest
    @DisplayName("Parameterized test for various inputs")
    @CsvSource({
        "'', 0",
        "'1', 1",
        "'1,2', 3",
        "'1,2,3', 6",
        "'1\n2,3', 6",
        "'//;\n1;2', 3",
        "'2,1001', 2"
    })
    void parameterizedTest(String input, int expected) {
        assertEquals(expected, calculator.add(input));
    }
}

// TDD Example: Password Validator
class PasswordValidator {
    
    public boolean validate(String password) {
        if (password == null || password.length() < 8) {
            return false;
        }
        
        boolean hasUpper = false;
        boolean hasLower = false;
        boolean hasDigit = false;
        boolean hasSpecial = false;
        
        for (char c : password.toCharArray()) {
            if (Character.isUpperCase(c)) hasUpper = true;
            else if (Character.isLowerCase(c)) hasLower = true;
            else if (Character.isDigit(c)) hasDigit = true;
            else if (!Character.isLetterOrDigit(c)) hasSpecial = true;
        }
        
        return hasUpper && hasLower && hasDigit && hasSpecial;
    }
}

@DisplayName("Password Validator TDD")
class PasswordValidatorTest {
    
    private PasswordValidator validator;
    
    @BeforeEach
    void setUp() {
        validator = new PasswordValidator();
    }
    
    @Test
    @DisplayName("Should reject null password")
    void testNullPassword() {
        assertFalse(validator.validate(null));
    }
    
    @Test
    @DisplayName("Should reject short passwords")
    void testShortPasswords() {
        assertFalse(validator.validate("abc"));
        assertFalse(validator.validate("1234567"));
    }
    
    @Test
    @DisplayName("Should require uppercase letter")
    void testRequiresUppercase() {
        assertFalse(validator.validate("password123!"));
        assertTrue(validator.validate("Password123!"));
    }
    
    @Test
    @DisplayName("Should require lowercase letter")
    void testRequiresLowercase() {
        assertFalse(validator.validate("PASSWORD123!"));
        assertTrue(validator.validate("Password123!"));
    }
    
    @Test
    @DisplayName("Should require digit")
    void testRequiresDigit() {
        assertFalse(validator.validate("Password!"));
        assertTrue(validator.validate("Password1!"));
    }
    
    @Test
    @DisplayName("Should require special character")
    void testRequiresSpecialCharacter() {
        assertFalse(validator.validate("Password123"));
        assertTrue(validator.validate("Password123!"));
    }
    
    @Test
    @DisplayName("Should accept valid passwords")
    void testValidPasswords() {
        assertTrue(validator.validate("Password123!"));
        assertTrue(validator.validate("MySecure@Pass1"));
        assertTrue(validator.validate("Complex#Password9"));
    }
    
    @ParameterizedTest
    @DisplayName("Parameterized password validation")
    @CsvSource({
        "'Password123!', true",
        "'password123!', false",
        "'PASSWORD123!', false",
        "'Password!', false",
        "'Password123', false",
        "'Pass1!', false",
        "'null', false"
    })
    void parameterizedPasswordTest(String password, boolean expected) {
        if ("null".equals(password)) {
            assertFalse(validator.validate(null));
        } else {
            assertEquals(expected, validator.validate(password));
        }
    }
}

public class TDDDemonstration {
    public static void main(String[] args) {
        System.out.println("This is a TDD demonstration class.");
        System.out.println("The tests follow the Red-Green-Refactor cycle:");
        System.out.println("1. Red: Write a failing test");
        System.out.println("2. Green: Write minimal code to make it pass");
        System.out.println("3. Refactor: Improve the code while keeping tests green");
        System.out.println("\nRun tests with: mvn test");
    }
}
```

## Practice Exercises

1. **Unit Testing**: Write comprehensive tests for a banking system
2. **Mocking**: Create tests for a service with multiple dependencies
3. **Integration Testing**: Build end-to-end tests for a REST API
4. **TDD**: Implement a new feature using test-driven development

## Interview Questions

1. What's the difference between unit tests and integration tests?
2. When would you use mocks vs. real implementations?
3. What are the benefits of Test-Driven Development?
4. How do you test private methods?
5. What's the purpose of test coverage and what's a good target percentage?
