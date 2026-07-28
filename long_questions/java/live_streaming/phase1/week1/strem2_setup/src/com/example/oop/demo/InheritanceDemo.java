package com.example.oop.demo;

import com.example.oop.Employee;
import com.example.oop.Manager;
import com.example.oop.Developer;

/**
 * Demonstration of inheritance, super keyword, and method overriding
 * Run this class to understand inheritance concepts
 */
public class InheritanceDemo {
    
    public static void main(String[] args) {
        System.out.println("=== Inheritance Demo ===\n");
        
        // Create Employee (parent class)
        System.out.println("--- Creating Employee (Parent Class) ---");
        Employee emp = new Employee(101, "Alice", 60000, "HR");
        emp.displayDetails();
        System.out.println("Job Title: " + emp.getJobTitle());
        System.out.println();
        
        // Create Manager (subclass)
        System.out.println("--- Creating Manager (Subclass) ---");
        Manager mgr = new Manager(102, "Bob", 90000, "Engineering", 
                                  5000, "Backend Team");
        mgr.displayDetails();  // Inherited from Employee
        System.out.println("Job Title: " + mgr.getJobTitle());  // Overridden
        System.out.println("Bonus: $" + mgr.getBonus());
        System.out.println("Team: " + mgr.getTeam());
        mgr.assignProject("New API");  // Manager-specific method
        System.out.println();
        
        // Create Developer (subclass)
        System.out.println("--- Creating Developer (Subclass) ---");
        String[] languages = {"Java", "Python", "JavaScript"};
        Developer dev = new Developer(103, "Charlie", 75000, "Engineering",
                                      languages, 5);
        dev.displayDetails();  // Inherited from Employee
        System.out.println("Job Title: " + dev.getJobTitle());  // Overridden
        System.out.println("Languages: " + String.join(", ", dev.getProgrammingLanguages()));
        System.out.println("Experience: " + dev.getYearsOfExperience() + " years");
        dev.writeCode("Java");  // Developer-specific method
        System.out.println("Knows Python: " + dev.knowsLanguage("Python"));
        System.out.println("Knows C++: " + dev.knowsLanguage("C++"));
        System.out.println();
        
        // Demonstrating method overriding
        System.out.println("--- Method Overriding Demo ---");
        System.out.println("Employee raise:");
        emp.giveRaise(10);
        System.out.println("Manager raise (includes bonus):");
        mgr.giveRaise(10);
        System.out.println("Developer raise (experience bonus):");
        dev.giveRaise(10);
        System.out.println();
        
        // Demonstrating polymorphic references
        System.out.println("--- Polymorphic References ---");
        Employee polyEmp1 = new Employee(104, "David", 50000, "Sales");
        Employee polyEmp2 = new Manager(105, "Eve", 85000, "Engineering", 
                                        3000, "Frontend Team");
        Employee polyEmp3 = new Developer(106, "Frank", 70000, "Engineering",
                                          new String[]{"Go", "Rust"}, 3);
        
        System.out.println("polyEmp1: " + polyEmp1.getJobTitle());
        System.out.println("polyEmp2: " + polyEmp2.getJobTitle());
        System.out.println("polyEmp3: " + polyEmp3.getJobTitle());
        System.out.println();
        
        // Demonstrating instanceof and type casting
        System.out.println("--- instanceof and Type Casting ---");
        if (polyEmp2 instanceof Manager) {
            Manager castedMgr = (Manager) polyEmp2;
            System.out.println("polyEmp2 is a Manager");
            castedMgr.assignProject("Mobile App");
        }
        
        if (polyEmp3 instanceof Developer) {
            Developer castedDev = (Developer) polyEmp3;
            System.out.println("polyEmp3 is a Developer");
            castedDev.writeCode("Rust");
        }
        
        if (polyEmp1 instanceof Manager) {
            System.out.println("polyEmp1 is a Manager");
        } else {
            System.out.println("polyEmp1 is NOT a Manager");
        }
    }
}
