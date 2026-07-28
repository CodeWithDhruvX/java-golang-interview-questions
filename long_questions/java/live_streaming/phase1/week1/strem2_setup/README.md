# Stream 2: OOP Deep Dive - Setup Files

Complete Java project files for Stream 2: Object-Oriented Programming Deep Dive

## Project Structure

```
strem2_setup/
├── src/
│   └── com/
│       └── example/
│           └── oop/
│               ├── Employee.java                    # Base class
│               ├── Manager.java                     # Subclass
│               ├── Developer.java                   # Subclass
│               ├── AbstractEmployee.java            # Abstract class
│               ├── Workable.java                    # Interface
│               ├── Billable.java                    # Interface
│               ├── FullTimeEmployee.java            # Implements abstract & interfaces
│               ├── Contractor.java                  # Implements interfaces only
│               ├── EmployeeManagementSystem.java    # Main application
│               └── demo/
│                   ├── BasicDemo.java
│                   ├── InheritanceDemo.java
│                   ├── PolymorphismDemo.java
│                   ├── AbstractionDemo.java
│                   ├── SystemDemo.java
│                   └── EdgeCaseTests.java
└── README.md
```

## Setup Instructions

1. **Open in IntelliJ IDEA**
   - File → Open → Select `strem2_setup` folder
   - Let IntelliJ detect the project structure

2. **Configure JDK**
   - File → Project Structure → Project
   - Set Project SDK to JDK 17 or 21
   - Set Project language level to match JDK

3. **Run the Project**
   - Open `EmployeeManagementSystem.java`
   - Click the green play button next to `main` method
   - Or right-click → Run 'EmployeeManagementSystem.main()'

4. **Run Demo Classes**
   - Each demo class in `demo/` folder can be run independently
   - They demonstrate specific OOP concepts

## File Descriptions

### Core Classes
- **Employee.java**: Base class with encapsulation, constructors, getters/setters
- **Manager.java**: Subclass demonstrating inheritance and method overriding
- **Developer.java**: Subclass with array fields and experience-based logic

### Abstraction Examples
- **AbstractEmployee.java**: Abstract class with concrete and abstract methods
- **Workable.java**: Interface with default and static methods
- **Billable.java**: Interface for billing functionality
- **FullTimeEmployee.java**: Extends abstract class, implements interfaces
- **Contractor.java**: Implements interfaces only (no inheritance)

### Main Application
- **EmployeeManagementSystem.java**: Complete CRUD system with interactive menu

### Demo Classes
- **BasicDemo.java**: Classes, objects, constructors
- **InheritanceDemo.java**: Inheritance, super keyword, method overriding
- **PolymorphismDemo.java**: Overloading, overriding, runtime polymorphism
- **AbstractionDemo.java**: Abstract classes vs interfaces
- **SystemDemo.java**: Complete Employee Management System
- **EdgeCaseTests.java**: Edge cases and error handling

## Learning Order

1. Start with `BasicDemo.java` - understand classes and objects
2. Run `InheritanceDemo.java` - learn inheritance
3. Run `PolymorphismDemo.java` - understand polymorphism
4. Run `AbstractionDemo.java` - abstract classes vs interfaces
5. Run `SystemDemo.java` - complete system
6. Run `EdgeCaseTests.java` - error handling

## Prerequisites

- JDK 17 or 21 installed
- IntelliJ IDEA (Community Edition works)
- Basic Java syntax knowledge

## Common Issues

**Issue**: "Package does not exist"
**Solution**: Ensure `src/com/example/oop/` directory structure matches package declaration

**Issue**: "Cannot find symbol"
**Solution**: Check imports and ensure all files are in correct package

**Issue**: Main class not found
**Solution**: Run the specific class with main method, not the package

## Next Steps

After completing this setup, proceed to Stream 3: Data Types, Operators & Control Flow
