# Stream 2: Object-Oriented Programming Deep Dive - Complete Script

**Duration**: 2 hours  
**Project**: Employee Management System  
**Repository**: `long_questions/java/core-java/basics/`

---

## Stream Overview

**Target Audience**: Beginners with basic Java knowledge  
**Prerequisites**: Completed Stream 1 (Environment Setup)  
**Learning Outcomes**:
- Understand classes, objects, and constructors
- Master inheritance, polymorphism, and encapsulation
- Differentiate between abstract classes and interfaces
- Build a complete Employee Management System using OOP principles

---

## Timeline Breakdown

| Segment | Duration | Description |
|---------|----------|-------------|
| Introduction | 10 mins | Welcome, OOP concepts overview, objectives |
| Classes & Objects | 20 mins | Class structure, object creation, constructors |
| Encapsulation Deep Dive | 15 mins | Access modifiers, getters/setters, data hiding |
| Inheritance Implementation | 20 mins | Extending classes, super keyword, method overriding |
| Polymorphism in Action | 15 mins | Method overloading, runtime polymorphism |
| Abstract Classes vs Interfaces | 20 mins | When to use each, practical examples |
| Employee Management System Build | 25 mins | Complete project implementation |
| Testing & Demo | 15 mins | Run system, demonstrate features |
| Q&A | 15 mins | Viewer questions and clarifications |
| Summary | 5 mins | Recap, homework, next stream preview |

---

## Complete Script

### 1. Introduction (10 minutes)

**[0:00-0:10] Welcome & OOP Overview**

"Welcome back to Stream 2 of our Java Live Coding series! Today we're diving deep into Object-Oriented Programming - the heart and soul of Java.

**Today's Objectives:**
- Understand what OOP is and why it matters
- Master the four pillars of OOP: Encapsulation, Inheritance, Polymorphism, Abstraction
- Learn the difference between abstract classes and interfaces
- Build a complete Employee Management System using all these concepts

**Prerequisites Check:**
- JDK installed (from Stream 1)
- IntelliJ IDEA set up
- Basic Java syntax knowledge (variables, methods, main method)

**What You'll Achieve:**
By the end of this stream, you'll have built a fully functional Employee Management System that demonstrates real-world OOP usage. You'll understand how to design classes, create objects, use inheritance to share code, implement polymorphism for flexibility, and choose between abstract classes and interfaces.

**Why OOP Matters:**
OOP isn't just theory - it's how real-world Java applications are built. Companies like Google, Amazon, and Netflix use OOP principles to manage millions of lines of code. Understanding OOP is essential for:
- Writing maintainable code
- Reusing existing code
- Designing scalable systems
- Passing technical interviews

Let's jump in!"

---

### 2. Classes & Objects (20 minutes)

**[0:10-0:30] Understanding Classes and Objects**

"First, let's understand the fundamental building blocks of OOP.

**What is a Class?**
A class is a blueprint or template for creating objects. It defines:
- Properties (attributes/fields)
- Behaviors (methods)
- Constructors (special methods for initialization)

Think of a class like a cookie cutter - it defines the shape, but you need to use it to create actual cookies (objects).

**What is an Object?**
An object is an instance of a class. It has:
- State (values of properties)
- Behavior (can execute methods)
- Identity (unique in memory)

Using our cookie analogy: the cookie cutter is the class, each actual cookie is an object.

**Let's Create Our First Class**

Create new file: `Employee.java`

```java
package com.example.oop;

/**
 * Represents an employee in the company
 * This is our blueprint for creating employee objects
 */
public class Employee {
    
    // Instance variables (properties/fields)
    private int id;
    private String name;
    private double salary;
    private String department;
    
    // Default constructor
    public Employee() {
        this.id = 0;
        this.name = "Unknown";
        this.salary = 0.0;
        this.department = "Unassigned";
        System.out.println("Default constructor called");
    }
    
    // Parameterized constructor
    public Employee(int id, String name, double salary, String department) {
        this.id = id;
        this.name = name;
        this.salary = salary;
        this.department = department;
        System.out.println("Parameterized constructor called for: " + name);
    }
    
    // Copy constructor
    public Employee(Employee other) {
        this.id = other.id;
        this.name = other.name;
        this.salary = other.salary;
        this.department = other.department;
        System.out.println("Copy constructor called");
    }
    
    // Instance method (behavior)
    public void displayDetails() {
        System.out.println("Employee Details:");
        System.out.println("  ID: " + id);
        System.out.println("  Name: " + name);
        System.out.println("  Salary: $" + salary);
        System.out.println("  Department: " + department);
    }
    
    // Getter methods
    public int getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    public double getSalary() {
        return salary;
    }
    
    public String getDepartment() {
        return department;
    }
    
    // Setter methods
    public void setId(int id) {
        this.id = id;
    }
    
    public void setName(String name) {
        this.name = name;
    }
    
    public void setSalary(double salary) {
        if (salary > 0) {
            this.salary = salary;
        } else {
            System.out.println("Salary must be positive!");
        }
    }
    
    public void setDepartment(String department) {
        this.department = department;
    }
}
```

**Creating Objects**

Create `Main.java`:

```java
package com.example.oop;

public class Main {
    public static void main(String[] args) {
        System.out.println("=== Creating Employee Objects ===\n");
        
        // Using default constructor
        Employee emp1 = new Employee();
        emp1.displayDetails();
        System.out.println();
        
        // Using parameterized constructor
        Employee emp2 = new Employee(101, "John Doe", 75000.0, "Engineering");
        emp2.displayDetails();
        System.out.println();
        
        // Using copy constructor
        Employee emp3 = new Employee(emp2);
        emp3.setName("Jane Smith");
        emp3.setId(102);
        emp3.displayDetails();
        System.out.println();
        
        // Using setters to modify
        emp1.setId(103);
        emp1.setName("Bob Johnson");
        emp1.setSalary(65000.0);
        emp1.setDepartment("Marketing");
        emp1.displayDetails();
    }
}
```

**Key Concepts Explained:**

1. **`this` keyword**: Refers to the current object
   - `this.id = id` (assign parameter to instance variable)
   - Distinguishes between local variables and instance variables

2. **Constructors**: Special methods to initialize objects
   - Same name as class
   - No return type (not even void)
   - Can be overloaded (multiple constructors)
   - Called automatically when using `new`

3. **Instance variables**: Properties specific to each object
   - Each object has its own copy
   - Different from static variables (shared across all objects)

**Run the program and observe:**
- Default constructor creates employee with default values
- Parameterized constructor creates employee with specific values
- Copy constructor duplicates an existing employee
- Setters allow modification after creation

**Common Mistakes:**
- Forgetting `this` keyword when parameter names match field names
- Not initializing all fields in constructor
- Creating constructor with return type (it won't be a constructor!)"

---

### 3. Encapsulation Deep Dive (15 minutes)

**[0:30-0:45] Data Hiding and Access Control**

"Encapsulation is about hiding internal details and providing controlled access. It's one of the most important OOP principles.

**What is Encapsulation?**
- Bundling data (variables) and methods together
- Hiding internal state (making variables private)
- Providing access through public methods (getters/setters)

**Why Encapsulation Matters:**
1. **Data Protection**: Prevent invalid data
2. **Flexibility**: Can change internal implementation without affecting code that uses the class
3. **Maintainability**: Easier to debug and modify
4. **Control**: Validate data before setting

**Let's Enhance Our Employee Class with Better Encapsulation**

Add to `Employee.java`:

```java
public class Employee {
    // All fields are private (encapsulation)
    private int id;
    private String name;
    private double salary;
    private String department;
    
    // ... existing constructors ...
    
    // Enhanced setter with validation
    public void setSalary(double salary) {
        if (salary > 0) {
            this.salary = salary;
        } else {
            throw new IllegalArgumentException("Salary must be positive!");
        }
    }
    
    public void setName(String name) {
        if (name != null && !name.trim().isEmpty()) {
            this.name = name;
        } else {
            throw new IllegalArgumentException("Name cannot be empty!");
        }
    }
    
    // Read-only field (no setter)
    public int getId() {
        return id;
    }
    
    // Computed property (no backing field)
    public double getAnnualSalary() {
        return salary * 12;
    }
    
    // Business logic method
    public void giveRaise(double percentage) {
        if (percentage > 0 && percentage <= 50) {
            this.salary = this.salary * (1 + percentage / 100);
            System.out.println(name + " received a " + percentage + "% raise!");
        } else {
            System.out.println("Invalid raise percentage!");
        }
    }
    
    // toString method for easy printing
    @Override
    public String toString() {
        return "Employee[id=" + id + ", name=" + name + 
               ", salary=$" + salary + ", department=" + department + "]";
    }
}
```

**Access Modifiers Explained:**

```java
public class AccessDemo {
    public int publicVar = 10;        // Accessible everywhere
    protected int protectedVar = 20;   // Accessible in same package and subclasses
    int packageVar = 30;              // Default: accessible only in same package
    private int privateVar = 40;      // Accessible only in this class
    
    public void demoAccess() {
        System.out.println(publicVar);      // ✓ OK
        System.out.println(protectedVar);   // ✓ OK
        System.out.println(packageVar);     // ✓ OK
        System.out.println(privateVar);     // ✓ OK
    }
}
```

**Best Practices for Encapsulation:**
1. Make fields `private` by default
2. Provide `public` getters and setters
3. Add validation in setters
4. Make read-only fields (getter only)
5. Use `protected` for subclass access
6. Keep implementation details hidden

**Test Encapsulation:**

```java
public class TestEncapsulation {
    public static void main(String[] args) {
        Employee emp = new Employee(101, "John", 50000, "IT");
        
        // Valid operations
        emp.setSalary(60000);
        emp.giveRaise(10);
        
        // Invalid operations (will throw exceptions)
        try {
            emp.setSalary(-1000);  // Throws IllegalArgumentException
        } catch (IllegalArgumentException e) {
            System.out.println("Error: " + e.getMessage());
        }
        
        try {
            emp.setName("");  // Throws IllegalArgumentException
        } catch (IllegalArgumentException e) {
            System.out.println("Error: " + e.getMessage());
        }
        
        // Using computed property
        System.out.println("Annual salary: $" + emp.getAnnualSalary());
        
        // Using toString
        System.out.println(emp);
    }
}
```

**Interview Question Alert:**
'What's the difference between encapsulation and data hiding?' 
- Encapsulation: Bundling data and methods together
- Data hiding: Making data private (part of encapsulation)"

---

### 4. Inheritance Implementation (20 minutes)

**[0:45-1:05] Code Reuse Through Inheritance**

"Inheritance allows us to create new classes based on existing classes, promoting code reuse.

**What is Inheritance?**
- A subclass (child) inherits from a superclass (parent)
- Child gets all parent's fields and methods
- Child can add new fields and methods
- Child can override parent methods

**Why Use Inheritance?**
- Code reuse (don't repeat yourself)
- Hierarchical relationships (natural modeling)
- Polymorphism (treat objects uniformly)
- Easy maintenance (changes in parent affect all children)

**Let's Create an Inheritance Hierarchy**

Create `Manager.java` (subclass of Employee):

```java
package com.example.oop;

/**
 * Manager is a specialized type of Employee
 * Inherits all Employee properties and behaviors
 * Adds manager-specific features
 */
public class Manager extends Employee {
    
    // Manager-specific property
    private double bonus;
    private String team;
    
    // Constructor
    public Manager(int id, String name, double salary, 
                   String department, double bonus, String team) {
        // Call parent constructor using super()
        super(id, name, salary, department);
        this.bonus = bonus;
        this.team = team;
    }
    
    // Manager-specific method
    public void assignProject(String projectName) {
        System.out.println(getName() + " is assigning project: " + projectName + 
                          " to team: " + team);
    }
    
    // Override parent method
    @Override
    public void giveRaise(double percentage) {
        super.giveRaise(percentage);  // Call parent method
        this.bonus += 1000;  // Add manager-specific bonus
        System.out.println("Manager bonus increased by $1000!");
    }
    
    // Override toString
    @Override
    public String toString() {
        return super.toString() + 
               ", bonus=$" + bonus + 
               ", team=" + team + "]";
    }
    
    // Getters and setters
    public double getBonus() {
        return bonus;
    }
    
    public void setBonus(double bonus) {
        this.bonus = bonus;
    }
    
    public String getTeam() {
        return team;
    }
    
    public void setTeam(String team) {
        this.team = team;
    }
}
```

Create `Developer.java` (another subclass):

```java
package com.example.oop;

/**
 * Developer is a specialized type of Employee
 */
public class Developer extends Employee {
    
    private String[] programmingLanguages;
    private int yearsOfExperience;
    
    public Developer(int id, String name, double salary, 
                      String department, String[] languages, int experience) {
        super(id, name, salary, department);
        this.programmingLanguages = languages;
        this.yearsOfExperience = experience;
    }
    
    public void writeCode(String language) {
        System.out.println(getName() + " is writing code in " + language);
    }
    
    @Override
    public void giveRaise(double percentage) {
        // Developers get higher raises based on experience
        double adjustedPercentage = percentage + (yearsOfExperience * 0.5);
        super.giveRaise(adjustedPercentage);
    }
    
    @Override
    public String toString() {
        return super.toString() + 
               ", languages=" + String.join(", ", programmingLanguages) +
               ", experience=" + yearsOfExperience + " years]";
    }
}
```

**Understanding `super` keyword:**

```java
// super refers to the parent class
super();                    // Call parent constructor
super.methodName();         // Call parent method
super.fieldName;            // Access parent field
```

**Test Inheritance:**

```java
public class TestInheritance {
    public static void main(String[] args) {
        System.out.println("=== Testing Inheritance ===\n");
        
        // Create Employee (parent)
        Employee emp = new Employee(101, "Alice", 60000, "HR");
        emp.displayDetails();
        System.out.println();
        
        // Create Manager (child)
        Manager mgr = new Manager(102, "Bob", 90000, "Engineering", 
                                  5000, "Backend Team");
        mgr.displayDetails();  // Inherited from Employee
        mgr.assignProject("New API");  // Manager-specific
        System.out.println(mgr);
        System.out.println();
        
        // Create Developer (child)
        String[] languages = {"Java", "Python", "JavaScript"};
        Developer dev = new Developer(103, "Charlie", 75000, "Engineering",
                                      languages, 5);
        dev.displayDetails();  // Inherited from Employee
        dev.writeCode("Java");  // Developer-specific
        System.out.println(dev);
        System.out.println();
        
        // Test method overriding
        System.out.println("=== Testing Method Overriding ===");
        mgr.giveRaise(10);  // Manager's version
        dev.giveRaise(10);   // Developer's version
    }
}
```

**Key Inheritance Concepts:**

1. **`extends` keyword**: Defines inheritance relationship
2. **`super` keyword**: Access parent class members
3. **`@Override` annotation**: Indicates method override (best practice)
4. **Single inheritance**: Java classes can only extend one class
5. **Method overriding**: Child provides specific implementation

**Common Mistakes:**
- Forgetting to call `super()` in constructor
- Trying to extend multiple classes (not allowed in Java)
- Overriding without `@Override` (misspelling errors)
- Accessing private parent members directly (use protected or getters)"

---

### 5. Polymorphism in Action (15 minutes)

**[1:05-1:20] Flexibility Through Polymorphism**

"Polymorphism means 'many forms' - it allows objects to be treated as instances of their parent class rather than their actual class.

**Types of Polymorphism:**

1. **Compile-time Polymorphism (Method Overloading)**
   - Same method name, different parameters
   - Resolved at compile time

2. **Runtime Polymorphism (Method Overriding)**
   - Subclass provides specific implementation
   - Resolved at runtime based on actual object type

**Method Overloading Example:**

Add to `Employee.java`:

```java
// Overloaded methods - same name, different parameters
public void work() {
    System.out.println(name + " is working.");
}

public void work(int hours) {
    System.out.println(name + " is working for " + hours + " hours.");
}

public void work(String project) {
    System.out.println(name + " is working on project: " + project);
}

public void work(int hours, String project) {
    System.out.println(name + " is working on " + project + 
                       " for " + hours + " hours.");
}
```

**Runtime Polymorphism Example:**

```java
public class TestPolymorphism {
    public static void main(String[] args) {
        System.out.println("=== Runtime Polymorphism ===\n");
        
        // Polymorphic references
        Employee emp1 = new Employee(101, "Alice", 60000, "HR");
        Employee emp2 = new Manager(102, "Bob", 90000, "Engineering", 
                                    5000, "Backend Team");
        Employee emp3 = new Developer(103, "Charlie", 75000, "Engineering",
                                       new String[]{"Java", "Python"}, 5);
        
        // All treated as Employee, but behave differently
        Employee[] employees = {emp1, emp2, emp3};
        
        for (Employee emp : employees) {
            System.out.println(emp);
            emp.giveRaise(10);  // Calls appropriate override
            System.out.println();
        }
        
        // Type casting to access subclass-specific methods
        if (emp2 instanceof Manager) {
            Manager mgr = (Manager) emp2;
            mgr.assignProject("New Feature");
        }
        
        if (emp3 instanceof Developer) {
            Developer dev = (Developer) emp3;
            dev.writeCode("Java");
        }
    }
}
```

**Understanding `instanceof` operator:**
- Checks if an object is an instance of a class
- Returns true if object is of specified type or subclass
- Used for safe type casting

**Polymorphic Method Calls:**

```java
public class PayrollSystem {
    public static void processPayroll(Employee[] employees) {
        for (Employee emp : employees) {
            // Polymorphic call - actual method depends on object type
            double monthlyPay = emp.getSalary();
            System.out.println("Processing payroll for: " + emp.getName());
            System.out.println("Monthly pay: $" + monthlyPay);
            
            // Polymorphic behavior
            if (emp instanceof Manager) {
                System.out.println("  + Manager bonus included");
            }
            if (emp instanceof Developer) {
                System.out.println("  + Tech allowance included");
            }
            System.out.println();
        }
    }
    
    public static void main(String[] args) {
        Employee[] staff = {
            new Employee(101, "Alice", 5000, "HR"),
            new Manager(102, "Bob", 7000, "Engineering", 1000, "Backend"),
            new Developer(103, "Charlie", 6000, "Engineering", 
                          new String[]{"Java"}, 3)
        };
        
        processPayroll(staff);
    }
}
```

**Why Polymorphism is Powerful:**
- Write flexible code that works with multiple types
- Add new types without changing existing code
- Follow Open/Closed Principle (open for extension, closed for modification)
- Simplify code with common interfaces

**Interview Question Alert:**
'What's the difference between overloading and overriding?'
- Overloading: Same method name, different parameters (compile-time)
- Overriding: Subclass provides new implementation (runtime)"

---

### 6. Abstract Classes vs Interfaces (20 minutes)

**[1:20-1:40] Choosing the Right Abstraction**

"Abstraction is about hiding complexity and showing only essential features. Java provides two ways to achieve this: abstract classes and interfaces.

**Abstract Classes:**
- Can have both abstract and concrete methods
- Can have instance variables
- Can have constructors
- Single inheritance (can extend only one)
- Used when classes share implementation

**Interfaces:**
- All methods are abstract (before Java 8)
- Can have default and static methods (Java 8+)
- Cannot have instance variables (only constants)
- No constructors
- Multiple inheritance (can implement multiple)
- Used to define contract/behavior

**When to Use Abstract Class:**
- When you want to share code among related classes
- When you need non-static, non-final fields
- When you want to declare non-public members
- When you need to use constructors

**When to Use Interface:**
- When you want to define a contract
- When unrelated classes can implement same interface
- When you want multiple inheritance
- When you want to specify behavior without implementation

**Let's Create Both:**

Abstract class `AbstractEmployee.java`:

```java
package com.example.oop;

/**
 * Abstract class for all employees
 * Provides common implementation
 */
public abstract class AbstractEmployee {
    
    protected int id;
    protected String name;
    protected double baseSalary;
    
    public AbstractEmployee(int id, String name, double baseSalary) {
        this.id = id;
        this.name = name;
        this.baseSalary = baseSalary;
    }
    
    // Concrete method (has implementation)
    public void displayBasicInfo() {
        System.out.println("ID: " + id + ", Name: " + name);
    }
    
    // Abstract method (no implementation - must be overridden)
    public abstract double calculateSalary();
    
    public abstract String getJobTitle();
    
    // Concrete method with abstract behavior
    public void generatePaycheck() {
        System.out.println("=== PAYCHECK ===");
        displayBasicInfo();
        System.out.println("Position: " + getJobTitle());
        System.out.println("Amount: $" + calculateSalary());
        System.out.println("================");
    }
}
```

Interface `Workable.java`:

```java
package com.example.oop;

/**
 * Interface defining work-related behavior
 * Any class can implement this
 */
public interface Workable {
    
    // Constant (public static final by default)
    int STANDARD_WORK_HOURS = 40;
    
    // Abstract method (public abstract by default)
    void work();
    
    void takeBreak();
    
    // Default method (Java 8+) - provides implementation
    default void attendMeeting(String meetingTopic) {
        System.out.println("Attending meeting: " + meetingTopic);
    }
    
    // Static method (Java 8+) - utility method
    static void displayWorkPolicy() {
        System.out.println("Work Policy: " + STANDARD_WORK_HOURS + " hours/week");
    }
}
```

Interface `Billable.java`:

```java
package com.example.oop;

/**
 * Interface for billable employees
 */
public interface Billable {
    
    double calculateBillableHours();
    
    double getHourlyRate();
    
    default double generateInvoice() {
        return calculateBillableHours() * getHourlyRate();
    }
}
```

**Implementing Abstract Class and Interface:**

```java
package com.example.oop;

/**
 * FullTimeEmployee extends abstract class and implements interfaces
 */
public class FullTimeEmployee extends AbstractEmployee 
                              implements Workable, Billable {
    
    private String department;
    private double billableHours;
    private double hourlyRate;
    
    public FullTimeEmployee(int id, String name, double baseSalary,
                            String department) {
        super(id, name, baseSalary);
        this.department = department;
        this.hourlyRate = baseSalary / 160; // Assume 160 hours/month
        this.billableHours = 160;
    }
    
    @Override
    public double calculateSalary() {
        return baseSalary;
    }
    
    @Override
    public String getJobTitle() {
        return "Full-Time Employee";
    }
    
    // Implement Workable interface
    @Override
    public void work() {
        System.out.println(name + " is working full-time in " + department);
    }
    
    @Override
    public void takeBreak() {
        System.out.println(name + " is taking a scheduled break");
    }
    
    // Override default method if needed
    @Override
    public void attendMeeting(String meetingTopic) {
        System.out.println(name + " (Full-Time) attending: " + meetingTopic);
    }
    
    // Implement Billable interface
    @Override
    public double calculateBillableHours() {
        return billableHours;
    }
    
    @Override
    public double getHourlyRate() {
        return hourlyRate;
    }
    
    @Override
    public double generateInvoice() {
        double invoice = super.calculateSalary(); // Use base salary
        System.out.println("Invoice generated for full-time employee: $" + invoice);
        return invoice;
    }
}
```

**Contractor implementing only interfaces:**

```java
package com.example.oop;

/**
 * Contractor implements interfaces but doesn't extend abstract class
 */
public class Contractor implements Workable, Billable {
    
    private String name;
    private String company;
    private double hourlyRate;
    private double hoursWorked;
    
    public Contractor(String name, String company, double hourlyRate) {
        this.name = name;
        this.company = company;
        this.hourlyRate = hourlyRate;
        this.hoursWorked = 0;
    }
    
    @Override
    public void work() {
        System.out.println(name + " from " + company + " is working on contract");
    }
    
    @Override
    public void takeBreak() {
        System.out.println(name + " is taking a break (unpaid)");
    }
    
    @Override
    public double calculateBillableHours() {
        return hoursWorked;
    }
    
    @Override
    public double getHourlyRate() {
        return hourlyRate;
    }
    
    public void logHours(double hours) {
        this.hoursWorked += hours;
        System.out.println("Logged " + hours + " hours. Total: " + hoursWorked);
    }
    
    public String getName() {
        return name;
    }
}
```

**Test Abstract Classes and Interfaces:**

```java
public class TestAbstraction {
    public static void main(String[] args) {
        System.out.println("=== Testing Abstract Classes & Interfaces ===\n");
        
        // Full-time employee (extends abstract class, implements interfaces)
        FullTimeEmployee ft = new FullTimeEmployee(101, "Alice", 6000, "IT");
        ft.generatePaycheck();
        ft.work();
        ft.attendMeeting("Project Planning");
        System.out.println("Invoice: $" + ft.generateInvoice());
        System.out.println();
        
        // Contractor (implements interfaces only)
        Contractor contractor = new Contractor("Bob", "TechCorp", 50.0);
        contractor.logHours(40);
        contractor.logHours(35);
        contractor.work();
        contractor.attendMeeting("Client Call");
        System.out.println("Invoice: $" + contractor.generateInvoice());
        System.out.println();
        
        // Interface static method
        Workable.displayWorkPolicy();
        
        // Polymorphism with interfaces
        Workable[] workers = {ft, contractor};
        System.out.println("\n=== All Workers ===");
        for (Workable worker : workers) {
            worker.work();
        }
    }
}
```

**Key Differences Summary:**

| Feature | Abstract Class | Interface |
|---------|---------------|-----------|
| Methods | Can have concrete & abstract | All abstract (before Java 8) |
| Variables | Can have instance variables | Only constants |
| Constructors | Yes | No |
| Inheritance | Single | Multiple |
| Usage | Share implementation | Define contract |

**Best Practice:**
- Use abstract class when you have shared implementation
- Use interface when you want to define behavior contract
- Use default methods in interfaces to avoid breaking existing implementations
- Prefer interfaces over abstract classes for flexibility"

---

### 7. Employee Management System Build (25 minutes)

**[1:40-2:05] Complete Project Implementation**

"Now let's build a complete Employee Management System that brings together all the OOP concepts we've learned.

**System Features:**
- Add employees (different types)
- Display all employees
- Search employees by ID
- Give raises to employees
- Calculate total payroll
- Filter employees by department

**Complete System Implementation:**

Create `EmployeeManagementSystem.java`:

```java
package com.example.oop;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Complete Employee Management System
 * Demonstrates all OOP principles
 */
public class EmployeeManagementSystem {
    
    private List<Employee> employees;
    private Scanner scanner;
    
    public EmployeeManagementSystem() {
        this.employees = new ArrayList<>();
        this.scanner = new Scanner(System.in);
    }
    
    // Add employee
    public void addEmployee(Employee emp) {
        employees.add(emp);
        System.out.println("Employee added: " + emp.getName());
    }
    
    // Display all employees
    public void displayAllEmployees() {
        System.out.println("\n=== All Employees ===");
        if (employees.isEmpty()) {
            System.out.println("No employees in system.");
            return;
        }
        
        for (Employee emp : employees) {
            System.out.println(emp);
        }
        System.out.println("Total: " + employees.size() + " employees\n");
    }
    
    // Search employee by ID
    public Employee findEmployeeById(int id) {
        for (Employee emp : employees) {
            if (emp.getId() == id) {
                return emp;
            }
        }
        return null;
    }
    
    // Give raise to employee
    public void giveRaise(int employeeId, double percentage) {
        Employee emp = findEmployeeById(employeeId);
        if (emp != null) {
            emp.giveRaise(percentage);
            System.out.println("Raise given to " + emp.getName());
        } else {
            System.out.println("Employee not found with ID: " + employeeId);
        }
    }
    
    // Calculate total payroll
    public double calculateTotalPayroll() {
        double total = 0;
        for (Employee emp : employees) {
            total += emp.getSalary();
        }
        return total;
    }
    
    // Filter by department
    public List<Employee> filterByDepartment(String department) {
        List<Employee> filtered = new ArrayList<>();
        for (Employee emp : employees) {
            if (emp.getDepartment().equalsIgnoreCase(department)) {
                filtered.add(emp);
            }
        }
        return filtered;
    }
    
    // Interactive menu
    public void runMenu() {
        while (true) {
            System.out.println("\n=== Employee Management System ===");
            System.out.println("1. Add Employee");
            System.out.println("2. Display All Employees");
            System.out.println("3. Search Employee by ID");
            System.out.println("4. Give Raise");
            System.out.println("5. Calculate Total Payroll");
            System.out.println("6. Filter by Department");
            System.out.println("7. Exit");
            System.out.print("Choose option: ");
            
            int choice = scanner.nextInt();
            scanner.nextLine(); // Consume newline
            
            switch (choice) {
                case 1:
                    addEmployeeMenu();
                    break;
                case 2:
                    displayAllEmployees();
                    break;
                case 3:
                    searchEmployeeMenu();
                    break;
                case 4:
                    giveRaiseMenu();
                    break;
                case 5:
                    System.out.println("Total Payroll: $" + calculateTotalPayroll());
                    break;
                case 6:
                    filterByDepartmentMenu();
                    break;
                case 7:
                    System.out.println("Exiting system...");
                    scanner.close();
                    return;
                default:
                    System.out.println("Invalid option!");
            }
        }
    }
    
    private void addEmployeeMenu() {
        System.out.println("\n--- Add Employee ---");
        System.out.println("1. Regular Employee");
        System.out.println("2. Manager");
        System.out.println("3. Developer");
        System.out.print("Choose type: ");
        
        int type = scanner.nextInt();
        scanner.nextLine();
        
        System.out.print("Enter ID: ");
        int id = scanner.nextInt();
        scanner.nextLine();
        
        System.out.print("Enter Name: ");
        String name = scanner.nextLine();
        
        System.out.print("Enter Salary: ");
        double salary = scanner.nextDouble();
        scanner.nextLine();
        
        System.out.print("Enter Department: ");
        String department = scanner.nextLine();
        
        Employee emp;
        
        switch (type) {
            case 1:
                emp = new Employee(id, name, salary, department);
                break;
            case 2:
                System.out.print("Enter Bonus: ");
                double bonus = scanner.nextDouble();
                scanner.nextLine();
                System.out.print("Enter Team: ");
                String team = scanner.nextLine();
                emp = new Manager(id, name, salary, department, bonus, team);
                break;
            case 3:
                System.out.print("Enter number of programming languages: ");
                int numLangs = scanner.nextInt();
                scanner.nextLine();
                String[] languages = new String[numLangs];
                for (int i = 0; i < numLangs; i++) {
                    System.out.print("Enter language " + (i+1) + ": ");
                    languages[i] = scanner.nextLine();
                }
                System.out.print("Enter years of experience: ");
                int experience = scanner.nextInt();
                emp = new Developer(id, name, salary, department, languages, experience);
                break;
            default:
                System.out.println("Invalid type!");
                return;
        }
        
        addEmployee(emp);
    }
    
    private void searchEmployeeMenu() {
        System.out.print("\nEnter Employee ID: ");
        int id = scanner.nextInt();
        scanner.nextLine();
        
        Employee emp = findEmployeeById(id);
        if (emp != null) {
            System.out.println("Found: " + emp);
            emp.displayDetails();
        } else {
            System.out.println("Employee not found.");
        }
    }
    
    private void giveRaiseMenu() {
        System.out.print("\nEnter Employee ID: ");
        int id = scanner.nextInt();
        scanner.nextLine();
        
        System.out.print("Enter raise percentage: ");
        double percentage = scanner.nextDouble();
        scanner.nextLine();
        
        giveRaise(id, percentage);
    }
    
    private void filterByDepartmentMenu() {
        System.out.print("\nEnter Department: ");
        String department = scanner.nextLine();
        
        List<Employee> filtered = filterByDepartment(department);
        System.out.println("\n--- Employees in " + department + " ---");
        for (Employee emp : filtered) {
            System.out.println(emp);
        }
        System.out.println("Total: " + filtered.size() + " employees");
    }
    
    public static void main(String[] args) {
        EmployeeManagementSystem system = new EmployeeManagementSystem();
        
        // Add sample data
        system.addEmployee(new Employee(101, "Alice Johnson", 5000, "HR"));
        system.addEmployee(new Manager(102, "Bob Smith", 8000, "Engineering", 
                                        2000, "Backend Team"));
        String[] langs = {"Java", "Python", "JavaScript"};
        system.addEmployee(new Developer(103, "Charlie Brown", 7000, "Engineering",
                                          langs, 5));
        
        // Run interactive menu
        system.runMenu();
    }
}
```

**OOP Principles Demonstrated:**

1. **Encapsulation**: All fields are private, accessed via getters/setters
2. **Inheritance**: Manager and Developer extend Employee
3. **Polymorphism**: Employee list contains different types, each behaves differently
4. **Abstraction**: Could use abstract class/interface for common employee behavior

**Run the System:**

```java
public class Main {
    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        // Add sample employees
        ems.addEmployee(new Employee(101, "Alice", 5000, "HR"));
        ems.addEmployee(new Manager(102, "Bob", 8000, "Engineering", 2000, "Backend"));
        String[] langs = {"Java", "Python"};
        ems.addEmployee(new Developer(103, "Charlie", 7000, "Engineering", langs, 3));
        
        // Display all
        ems.displayAllEmployees();
        
        // Search
        Employee found = ems.findEmployeeById(102);
        if (found != null) {
            System.out.println("Found: " + found);
        }
        
        // Give raise
        ems.giveRaise(101, 10);
        
        // Calculate payroll
        System.out.println("Total Payroll: $" + ems.calculateTotalPayroll());
        
        // Filter
        List<Employee> engineering = ems.filterByDepartment("Engineering");
        System.out.println("\nEngineering Department:");
        for (Employee emp : engineering) {
            System.out.println(emp);
        }
    }
}
```

**System Features Explained:**
- **ArrayList**: Dynamic array to store employees
- **Scanner**: User input handling
- **Switch-case**: Menu navigation
- **Polymorphism**: List of Employee holds different types
- **Encapsulation**: All operations go through methods
- **Inheritance**: Different employee types with shared behavior"

---

### 8. Testing & Demo (15 minutes)

**[2:05-2:20] Running and Demonstrating**

"Let's run our complete Employee Management System and test all features.

**Test 1: Basic Operations**

```java
public class SystemDemo {
    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        System.out.println("=== Test 1: Adding Employees ===");
        ems.addEmployee(new Employee(101, "Alice", 5000, "HR"));
        ems.addEmployee(new Manager(102, "Bob", 8000, "Engineering", 2000, "Backend"));
        String[] langs = {"Java", "Python", "Go"};
        ems.addEmployee(new Developer(103, "Charlie", 7000, "Engineering", langs, 5));
        
        System.out.println("\n=== Test 2: Display All ===");
        ems.displayAllEmployees();
        
        System.out.println("\n=== Test 3: Search by ID ===");
        Employee found = ems.findEmployeeById(102);
        if (found != null) {
            System.out.println("Found: " + found);
            found.displayDetails();
        }
        
        System.out.println("\n=== Test 4: Give Raises ===");
        ems.giveRaise(101, 10);  // Regular employee
        ems.giveRaise(102, 10);  // Manager (gets bonus)
        ems.giveRaise(103, 10);  // Developer (experience bonus)
        
        System.out.println("\n=== Test 5: Total Payroll ===");
        System.out.println("Total Payroll: $" + ems.calculateTotalPayroll());
        
        System.out.println("\n=== Test 6: Filter by Department ===");
        List<Employee> engineering = ems.filterByDepartment("Engineering");
        System.out.println("Engineering employees: " + engineering.size());
        for (Employee emp : engineering) {
            System.out.println("  " + emp.getName() + " - " + emp.getJobTitle());
        }
    }
}
```

**Test 2: Edge Cases**

```java
public class EdgeCaseTests {
    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        // Test empty system
        System.out.println("=== Test: Empty System ===");
        ems.displayAllEmployees();
        System.out.println("Payroll: $" + ems.calculateTotalPayroll());
        
        // Test non-existent employee
        System.out.println("\n=== Test: Non-existent Employee ===");
        ems.giveRaise(999, 10);
        Employee notFound = ems.findEmployeeById(999);
        System.out.println("Result: " + (notFound == null ? "Not found" : "Found"));
        
        // Test invalid raise
        System.out.println("\n=== Test: Invalid Raise ===");
        Employee emp = new Employee(101, "Test", 5000, "IT");
        ems.addEmployee(emp);
        ems.giveRaise(101, -5);  // Should handle gracefully
        
        // Test department filter with no matches
        System.out.println("\n=== Test: No Department Matches ===");
        List<Employee> none = ems.filterByDepartment("NonExistent");
        System.out.println("Results: " + none.size());
    }
}
```

**Test 3: Interactive Mode**

Run the main method with interactive menu:
```java
public static void main(String[] args) {
    EmployeeManagementSystem ems = new EmployeeManagementSystem();
    ems.runMenu();
}
```

**Demonstrate:**
1. Add different employee types
2. Display all employees
3. Search for specific employee
4. Give raises and observe different behaviors
5. Calculate payroll
6. Filter by department
7. Exit system

**Common Issues and Solutions:**

**Issue 1: Scanner Input Skipping**
```java
// Problem: nextInt() leaves newline in buffer
int id = scanner.nextInt();
String name = scanner.nextLine();  // Skips!

// Solution: Consume newline after nextInt()
int id = scanner.nextInt();
scanner.nextLine();  // Consume newline
String name = scanner.nextLine();
```

**Issue 2: NullPointerException**
```java
// Problem: Employee not found
Employee emp = findEmployeeById(999);
emp.displayDetails();  // Crashes!

// Solution: Check for null
Employee emp = findEmployeeById(999);
if (emp != null) {
    emp.displayDetails();
} else {
    System.out.println("Employee not found");
}
```

**Issue 3: Type Casting Issues**
```java
// Problem: Can't access subclass methods
Employee emp = new Manager(...);
emp.assignProject("X");  // Compile error!

// Solution: Check type and cast
if (emp instanceof Manager) {
    Manager mgr = (Manager) emp;
    mgr.assignProject("X");
}
```

**Debugging Tips:**
- Use IntelliJ's debugger to step through menu operations
- Add print statements to track program flow
- Test each feature independently before combining
- Use try-catch for user input errors"

---

### 9. Q&A Session (15 minutes)

**[2:20-2:35] Viewer Questions**

"Let's address common questions about OOP in Java.

**Q1: When should I use inheritance vs composition?**
A: Use inheritance when there's an 'is-a' relationship (Manager IS-A Employee). Use composition when there's a 'has-a' relationship (Car HAS-A Engine). Favor composition over inheritance for flexibility.

**Q2: Can a class extend multiple classes in Java?**
A: No, Java supports only single inheritance for classes. But a class can implement multiple interfaces. This avoids the 'diamond problem' from multiple inheritance.

**Q3: What's the difference between abstract class and interface in Java 8+?**
A: The line is blurring with default methods in interfaces. Key difference: abstract class can have state (instance variables), interfaces cannot. Use abstract class for shared state/implementation, interfaces for contracts.

**Q4: Why use getters and setters instead of public variables?**
A: Encapsulation! Getters/setters allow validation, logging, and can change implementation without breaking code. Example: `setSalary()` can validate positive values.

**Q5: What is method hiding vs method overriding?**
A: Overriding: Instance methods in subclass replace parent's version (runtime polymorphism). Hiding: Static methods in subclass hide parent's version (compile-time).

**Q6: Can constructors be overridden?**
A: No, constructors aren't inherited, so they can't be overridden. But they can be overloaded (multiple constructors with different parameters).

**Q7: What is the purpose of the `final` keyword in OOP?**
A: `final` class: Cannot be extended (e.g., String). `final` method: Cannot be overridden. `final` variable: Cannot be reassigned. Used for immutability and design decisions.

**Q8: How does polymorphism improve code design?**
A: Allows writing flexible code that works with multiple types. Example: `processPayroll(Employee[])` works with Employee, Manager, Developer without knowing specific types.

**Q9: What is a marker interface?**
A: Interface with no methods (e.g., Serializable, Cloneable). Used to mark a class for special treatment by JVM. Can be replaced with annotations in modern Java.

**Q10: Should I always use OOP in Java?**
A: Not always! For simple scripts, procedural style is fine. OOP shines in larger systems with complex relationships. Use the right tool for the job.

**Q11: What's the difference between `==` and `.equals()` in OOP?**
A: `==` compares references (memory addresses). `.equals()` compares content (can be overridden). Always override `.equals()` and `hashCode()` together for custom classes.

**Q12: How do I decide between abstract class and interface?**
A: Ask: Do classes share implementation? → Abstract class. Do unrelated classes need same behavior? → Interface. Need multiple inheritance? → Interface."

---

### 10. Summary & Next Steps (5 minutes)

**[2:35-2:40] Recap and Homework**

"Excellent work completing Stream 2! Let's recap what we accomplished:

**Today's Achievements:**
✅ Mastered classes and objects with constructors  
✅ Implemented encapsulation with access modifiers  
✅ Created inheritance hierarchy with Manager and Developer  
✅ Demonstrated polymorphism with method overloading and overriding  
✅ Differentiated between abstract classes and interfaces  
✅ Built complete Employee Management System  
✅ Tested all features with edge cases  

**Homework for Next Stream:**
1. **Practice**: Add a new employee type `Intern` that extends Employee
   - Add field: `universityName`
   - Override `giveRaise()` to limit to 5%
   - Add method: `completeTraining()`

2. **Explore**: Read about Java's `Object` class and its methods (`equals()`, `hashCode()`, `toString()`)

3. **Challenge**: Add serialization to Employee Management System
   - Save employees to file
   - Load employees from file
   - Use `Serializable` interface

**Next Stream Preview:**
Stream 3: Data Types, Operators & Control Flow
- Primitive vs reference types deep dive
- Operator precedence and type casting
- Loops and conditional statements
- Project: Calculator Application

**Repository Update:**
Today's code will be pushed to: `long_questions/java/core-java/basics/stream2_oop_deep_dive/`

**Key Takeaways:**
- OOP is about modeling real-world relationships in code
- Encapsulation protects data and enables validation
- Inheritance promotes code reuse
- Polymorphism enables flexible, extensible code
- Abstract classes and interfaces serve different purposes

**Community:**
- Join our Discord for OOP discussions
- GitHub repo: [link]
- Twitter: [handle] for updates

**Thank you for joining! See you in Stream 3!**"

---

## Additional Resources

### OOP Principles Summary

**Encapsulation**: Hide internal state, require interaction through methods
```java
private int id;
public int getId() { return id; }
public void setId(int id) { this.id = id; }
```

**Inheritance**: Create new classes from existing classes
```java
public class Manager extends Employee { }
```

**Polymorphism**: Objects of different types treated uniformly
```java
Employee emp = new Manager();  // Upcasting
```

**Abstraction**: Hide complexity, show essential features
```java
public abstract class Employee { }
public interface Workable { }
```

### Common Design Patterns

1. **Singleton**: Ensure only one instance exists
2. **Factory**: Create objects without specifying exact class
3. **Strategy**: Define family of algorithms, make them interchangeable
4. **Observer**: One-to-many dependency (when object changes, notify dependents)

### Quick Reference

| Concept | Keyword | Purpose |
|---------|---------|---------|
| Inheritance | `extends` | Inherit from class |
| Implementation | `implements` | Implement interface |
| Override | `@Override` | Indicate method override |
| Parent access | `super` | Access parent members |
| Current object | `this` | Reference current object |
| Final | `final` | Prevent modification |
| Abstract | `abstract` | Cannot be instantiated |
| Static | `static` | Class-level member |

### File Structure for Project

```
EmployeeManagementSystem/
├── src/
│   └── com/
│       └── example/
│           └── oop/
│               ├── Employee.java
│               ├── Manager.java
│               ├── Developer.java
│               ├── EmployeeManagementSystem.java
│               └── Main.java
└── out/
    └── production/
        └── com/example/oop/
            ├── Employee.class
            ├── Manager.class
            ├── Developer.class
            └── EmployeeManagementSystem.class
```

---

## Interview Questions from This Stream

1. **What are the four pillars of OOP?**
   - Encapsulation, Inheritance, Polymorphism, Abstraction

2. **What is the difference between a class and an object?**
   - Class is blueprint/template, object is instance of class

3. **Explain the `this` keyword in Java.**
   - Refers to current object, distinguishes instance variables from parameters

4. **What is method overloading vs method overriding?**
   - Overloading: Same name, different parameters (compile-time)
   - Overriding: Subclass provides new implementation (runtime)

5. **When should you use abstract class vs interface?**
   - Abstract class: Share implementation, need state
   - Interface: Define contract, multiple inheritance

6. **What is encapsulation and why is it important?**
   - Bundling data and methods, hiding internal state, enabling validation

7. **Can a class extend multiple classes in Java? Why/why not?**
   - No, single inheritance to avoid diamond problem

8. **What is polymorphism? Give an example.**
   - Objects treated as parent type, actual behavior based on actual type
   - Example: `Employee emp = new Manager(); emp.giveRaise();`

9. **What is the purpose of constructors?**
   - Initialize object state, called when using `new`

10. **What is the difference between `==` and `.equals()`?**
    - `==` compares references, `.equals()` compares content

11. **What is method hiding in Java?**
    - Static method in subclass hides parent's static method

12. **Explain the `super` keyword.**
    - Refers to parent class, used to call parent constructor/methods

13. **Can constructors be overridden?**
    - No, constructors aren't inherited

14. **What is a marker interface?**
    - Interface with no methods, used to mark class (e.g., Serializable)

15. **What is the difference between composition and inheritance?**
    - Inheritance: is-a relationship (Manager IS-A Employee)
    - Composition: has-a relationship (Car HAS-A Engine)

---

## Code Files Reference

### Employee.java (Base Class)
```java
package com.example.oop;

public class Employee {
    private int id;
    private String name;
    private double salary;
    private String department;
    
    public Employee() { }
    
    public Employee(int id, String name, double salary, String department) {
        this.id = id;
        this.name = name;
        this.salary = salary;
        this.department = department;
    }
    
    public void displayDetails() { /* ... */ }
    public void giveRaise(double percentage) { /* ... */ }
    
    // Getters and setters...
}
```

### Manager.java (Subclass)
```java
package com.example.oop;

public class Manager extends Employee {
    private double bonus;
    private String team;
    
    public Manager(int id, String name, double salary, 
                   String department, double bonus, String team) {
        super(id, name, salary, department);
        this.bonus = bonus;
        this.team = team;
    }
    
    @Override
    public void giveRaise(double percentage) {
        super.giveRaise(percentage);
        this.bonus += 1000;
    }
    
    public void assignProject(String projectName) { /* ... */ }
}
```

### Developer.java (Subclass)
```java
package com.example.oop;

public class Developer extends Employee {
    private String[] programmingLanguages;
    private int yearsOfExperience;
    
    // Constructor and methods...
}
```

### EmployeeManagementSystem.java
See complete implementation in Section 7.

---

## Notes for Streamer

### Preparation Checklist
- [ ] Test all code examples before stream
- [ ] Prepare Employee, Manager, Developer classes
- [ ] Have EmployeeManagementSystem ready
- [ ] Test interactive menu functionality
- [ ] Prepare edge case scenarios
- [ ] Have backup code snippets for common errors

### During Stream
- Use IntelliJ's diagram view to show class hierarchy
- Demonstrate inheritance with visual diagrams
- Show polymorphism behavior with live examples
- Use debugger to step through method calls
- Keep terminal visible for compilation/execution

### Common Viewer Issues to Anticipate
- "Can't find symbol" → Package/import issues
- "Constructor not found" → Parameter mismatch
- "Method not visible" → Access modifier issues
- Scanner input problems → Newline consumption
- Type casting errors → instanceof checks needed

### Backup Plans
- If IntelliJ fails: Show command-line compilation
- If code has bugs: Have pre-compiled versions ready
- If time runs short: Skip advanced features, focus on basics
- If interactive menu fails: Demonstrate with hardcoded test cases

### Engagement Tips
- Ask viewers to predict output before running code
- Poll: "Abstract class or interface for this scenario?"
- Challenge: "Add a new employee type on the fly"
- Q&A: Address specific OOP interview questions

---

**Stream Duration**: 2 hours  
**Difficulty Level**: Beginner to Intermediate  
**Prerequisites**: Stream 1 (Environment Setup)  
**Next Stream**: Data Types, Operators & Control Flow
