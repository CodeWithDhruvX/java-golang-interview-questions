package com.example.oop.demo;

import com.example.oop.Employee;
import com.example.oop.Manager;
import com.example.oop.EmployeeManagementSystem;
import java.util.List;

/**
 * Edge case tests and error handling demonstrations
 * Run this class to understand common issues and solutions
 */
public class EdgeCaseTests {
    
    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        System.out.println("=== Edge Case Tests ===\n");
        
        // Test 1: Empty System
        System.out.println("--- Test 1: Empty System ---");
        ems.displayAllEmployees();
        System.out.println("Payroll: $" + ems.calculateTotalPayroll());
        System.out.println();
        
        // Test 2: Non-existent Employee Search
        System.out.println("--- Test 2: Non-existent Employee Search ---");
        Employee notFound = ems.findEmployeeById(999);
        if (notFound == null) {
            System.out.println("Employee 999 not found (as expected)");
        }
        ems.giveRaise(999, 10);  // Should handle gracefully
        System.out.println();
        
        // Test 3: Invalid Raise Percentage
        System.out.println("--- Test 3: Invalid Raise Percentage ---");
        Employee emp = new Employee(101, "Test", 5000, "IT");
        ems.addEmployee(emp);
        
        System.out.println("Testing negative raise:");
        emp.giveRaise(-10);  // Should show invalid message
        
        System.out.println("Testing excessive raise:");
        emp.giveRaise(100);  // Should show invalid message
        System.out.println();
        
        // Test 4: Invalid Setter Values
        System.out.println("--- Test 4: Invalid Setter Values ---");
        try {
            emp.setSalary(-1000);
        } catch (IllegalArgumentException e) {
            System.out.println("Caught salary error: " + e.getMessage());
        }
        
        try {
            emp.setName("");
        } catch (IllegalArgumentException e) {
            System.out.println("Caught name error: " + e.getMessage());
        }
        
        try {
            emp.setName(null);
        } catch (IllegalArgumentException e) {
            System.out.println("Caught null name error: " + e.getMessage());
        }
        System.out.println();
        
        // Test 5: Department Filter with No Matches
        System.out.println("--- Test 5: Department Filter with No Matches ---");
        List<Employee> none = ems.filterByDepartment("NonExistent");
        System.out.println("Results for 'NonExistent': " + none.size() + " employees");
        
        List<Employee> caseSensitive = ems.filterByDepartment("it");  // Should still find "IT"
        System.out.println("Results for 'it' (case-insensitive): " + caseSensitive.size() + " employees");
        System.out.println();
        
        // Test 6: Delete Non-existent Employee
        System.out.println("--- Test 6: Delete Non-existent Employee ---");
        boolean deleted = ems.deleteEmployee(999);
        System.out.println("Delete result: " + deleted);
        System.out.println();
        
        // Test 7: Copy Constructor
        System.out.println("--- Test 7: Copy Constructor ---");
        Employee original = new Employee(102, "Original", 6000, "Sales");
        Employee copy = new Employee(original);
        
        System.out.println("Original: " + original);
        System.out.println("Copy: " + copy);
        
        // Modify original, copy should remain unchanged
        original.setName("Modified");
        original.setSalary(7000);
        
        System.out.println("\nAfter modifying original:");
        System.out.println("Original: " + original);
        System.out.println("Copy: " + copy);
        System.out.println();
        
        // Test 8: Manager-Specific Operations
        System.out.println("--- Test 8: Manager-Specific Operations ---");
        Manager mgr = new Manager(103, "Manager", 8000, "Engineering", 2000, "Team A");
        ems.addEmployee(mgr);
        
        mgr.assignProject("Project X");
        System.out.println("Manager bonus: $" + mgr.getBonus());
        System.out.println("Manager team: " + mgr.getTeam());
        System.out.println();
        
        // Test 9: Developer-Specific Operations
        System.out.println("--- Test 9: Developer-Specific Operations ---");
        String[] languages = {"Java", "Python", "JavaScript"};
        Developer dev = new Developer(104, "Developer", 7000, "Engineering", languages, 5);
        ems.addEmployee(dev);
        
        dev.writeCode("Java");
        System.out.println("Knows Java: " + dev.knowsLanguage("Java"));
        System.out.println("Knows C++: " + dev.knowsLanguage("C++"));
        System.out.println("Experience bonus applied in raise:");
        dev.giveRaise(10);
        System.out.println();
        
        // Test 10: Large Number of Employees
        System.out.println("--- Test 10: Large Number of Employees ---");
        System.out.println("Adding 50 employees...");
        for (int i = 200; i < 250; i++) {
            ems.addEmployee(new Employee(i, "Employee " + i, 4000 + (i % 2000), "Operations"));
        }
        System.out.println("Total employees: " + ems.calculateTotalPayroll());
        ems.displayAllEmployees();
        System.out.println();
        
        // Test 11: Type Casting Safety
        System.out.println("--- Test 11: Type Casting Safety ---");
        Employee emp1 = new Employee(105, "Regular", 5000, "HR");
        Employee emp2 = new Manager(106, "Manager", 8000, "Engineering", 2000, "Team");
        
        // Safe casting with instanceof
        if (emp1 instanceof Manager) {
            Manager m = (Manager) emp1;
            m.assignProject("X");
        } else {
            System.out.println("emp1 is not a Manager (correct)");
        }
        
        if (emp2 instanceof Manager) {
            Manager m = (Manager) emp2;
            m.assignProject("Y");
            System.out.println("emp2 is a Manager, project assigned (correct)");
        }
        System.out.println();
        
        // Test 12: Null Handling
        System.out.println("--- Test 12: Null Handling ---");
        Employee nullEmp = null;
        
        try {
            if (nullEmp != null) {
                nullEmp.displayDetails();
            } else {
                System.out.println("Employee is null (handled safely)");
            }
        } catch (NullPointerException e) {
            System.out.println("NullPointerException: " + e.getMessage());
        }
        
        System.out.println("\n=== All Edge Case Tests Complete ===");
    }
}
