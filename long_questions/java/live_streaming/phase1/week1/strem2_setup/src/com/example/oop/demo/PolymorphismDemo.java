package com.example.oop.demo;

import com.example.oop.Employee;
import com.example.oop.Manager;
import com.example.oop.Developer;

/**
 * Demonstration of polymorphism: method overloading and overriding
 * Run this class to understand polymorphism concepts
 */
public class PolymorphismDemo {
    
    public static void main(String[] args) {
        System.out.println("=== Polymorphism Demo ===\n");
        
        // Compile-time polymorphism (Method Overloading)
        System.out.println("--- Compile-time Polymorphism (Overloading) ---");
        Employee emp = new Employee(101, "Alice", 60000, "IT");
        emp.work();
        emp.work(8);
        emp.work("Database Migration");
        emp.work(4, "Bug Fixes");
        System.out.println();
        
        // Runtime polymorphism (Method Overriding)
        System.out.println("--- Runtime Polymorphism (Overriding) ---");
        Employee[] employees = {
            new Employee(101, "Alice", 60000, "HR"),
            new Manager(102, "Bob", 90000, "Engineering", 5000, "Backend"),
            new Developer(103, "Charlie", 75000, "Engineering", 
                          new String[]{"Java", "Python"}, 5)
        };
        
        System.out.println("Processing raises for all employees:");
        for (Employee e : employees) {
            System.out.println("\nProcessing: " + e.getName());
            e.giveRaise(10);  // Different behavior based on actual object type
        }
        System.out.println();
        
        // Polymorphic method calls
        System.out.println("--- Polymorphic Method Calls ---");
        for (Employee e : employees) {
            System.out.println(e.getName() + " - Job Title: " + e.getJobTitle());
        }
        System.out.println();
        
        // Polymorphism with arrays
        System.out.println("--- Polymorphism with Arrays ---");
        processEmployees(employees);
        System.out.println();
        
        // Demonstrating dynamic method dispatch
        System.out.println("--- Dynamic Method Dispatch ---");
        Employee emp1 = new Employee(104, "David", 55000, "Marketing");
        Employee emp2 = new Manager(105, "Eve", 85000, "Engineering", 4000, "QA");
        Employee emp3 = new Developer(106, "Frank", 70000, "Engineering",
                                      new String[]{"JavaScript"}, 4);
        
        // All references are of type Employee, but actual methods called depend on object type
        System.out.println("emp1 (Employee): " + emp1.getJobTitle());
        System.out.println("emp2 (Manager): " + emp2.getJobTitle());
        System.out.println("emp3 (Developer): " + emp3.getJobTitle());
        System.out.println();
        
        // Polymorphism in method parameters
        System.out.println("--- Polymorphism in Method Parameters ---");
        displayEmployeeInfo(emp1);
        displayEmployeeInfo(emp2);
        displayEmployeeInfo(emp3);
    }
    
    // Method that accepts Employee (polymorphic parameter)
    public static void processEmployees(Employee[] employees) {
        System.out.println("Processing " + employees.length + " employees:");
        for (Employee emp : employees) {
            System.out.println("  - " + emp.getName() + " (" + emp.getJobTitle() + ")");
        }
    }
    
    // Method that demonstrates polymorphic behavior
    public static void displayEmployeeInfo(Employee emp) {
        System.out.println("\n--- Employee Info ---");
        System.out.println("Name: " + emp.getName());
        System.out.println("Job Title: " + emp.getJobTitle());
        System.out.println("Salary: $" + emp.getSalary());
        
        // Type-specific behavior
        if (emp instanceof Manager) {
            Manager mgr = (Manager) emp;
            System.out.println("Team: " + mgr.getTeam());
            System.out.println("Bonus: $" + mgr.getBonus());
        } else if (emp instanceof Developer) {
            Developer dev = (Developer) emp;
            System.out.println("Languages: " + String.join(", ", dev.getProgrammingLanguages()));
            System.out.println("Experience: " + dev.getYearsOfExperience() + " years");
        }
    }
}
