package com.example.oop.demo;

import com.example.oop.Employee;
import com.example.oop.Manager;
import com.example.oop.Developer;
import com.example.oop.EmployeeManagementSystem;
import java.util.List;

/**
 * Demonstration of the complete Employee Management System
 * Run this class to see all features in action
 */
public class SystemDemo {
    
    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        System.out.println("=== Employee Management System Demo ===\n");
        
        // Test 1: Adding Employees
        System.out.println("--- Test 1: Adding Employees ---");
        ems.addEmployee(new Employee(101, "Alice Johnson", 5000, "HR"));
        ems.addEmployee(new Manager(102, "Bob Smith", 8000, "Engineering", 2000, "Backend Team"));
        String[] langs = {"Java", "Python", "Go"};
        ems.addEmployee(new Developer(103, "Charlie Brown", 7000, "Engineering", langs, 5));
        System.out.println();
        
        // Test 2: Display All Employees
        System.out.println("--- Test 2: Display All Employees ---");
        ems.displayAllEmployees();
        
        // Test 3: Search by ID
        System.out.println("--- Test 3: Search by ID ---");
        Employee found = ems.findEmployeeById(102);
        if (found != null) {
            System.out.println("Found: " + found);
            found.displayDetails();
        }
        System.out.println();
        
        // Test 4: Give Raises
        System.out.println("--- Test 4: Give Raises ---");
        System.out.println("Before raises:");
        System.out.println("Alice: $" + ems.findEmployeeById(101).getSalary());
        System.out.println("Bob: $" + ems.findEmployeeById(102).getSalary());
        System.out.println("Charlie: $" + ems.findEmployeeById(103).getSalary());
        
        System.out.println("\nGiving 10% raises:");
        ems.giveRaise(101, 10);  // Regular employee
        ems.giveRaise(102, 10);  // Manager (gets bonus)
        ems.giveRaise(103, 10);  // Developer (experience bonus)
        
        System.out.println("\nAfter raises:");
        System.out.println("Alice: $" + ems.findEmployeeById(101).getSalary());
        System.out.println("Bob: $" + ems.findEmployeeById(102).getSalary());
        System.out.println("Charlie: $" + ems.findEmployeeById(103).getSalary());
        System.out.println();
        
        // Test 5: Calculate Total Payroll
        System.out.println("--- Test 5: Calculate Total Payroll ---");
        double totalPayroll = ems.calculateTotalPayroll();
        System.out.println("Total Monthly Payroll: $" + totalPayroll);
        System.out.println("Total Annual Payroll: $" + (totalPayroll * 12));
        System.out.println();
        
        // Test 6: Filter by Department
        System.out.println("--- Test 6: Filter by Department ---");
        List<Employee> engineering = ems.filterByDepartment("Engineering");
        System.out.println("Engineering Department (" + engineering.size() + " employees):");
        for (Employee emp : engineering) {
            System.out.println("  " + emp.getName() + " - " + emp.getJobTitle());
        }
        System.out.println();
        
        List<Employee> hr = ems.filterByDepartment("HR");
        System.out.println("HR Department (" + hr.size() + " employees):");
        for (Employee emp : hr) {
            System.out.println("  " + emp.getName() + " - " + emp.getJobTitle());
        }
        System.out.println();
        
        // Test 7: Delete Employee
        System.out.println("--- Test 7: Delete Employee ---");
        System.out.println("Before deletion:");
        ems.displayAllEmployees();
        
        ems.deleteEmployee(101);
        
        System.out.println("After deletion:");
        ems.displayAllEmployees();
        
        // Test 8: Add more employees
        System.out.println("--- Test 8: Add More Employees ---");
        ems.addEmployee(new Employee(104, "Diana Prince", 5500, "Marketing"));
        ems.addEmployee(new Manager(105, "Eve Adams", 8500, "Engineering", 
                                    3000, "Frontend Team"));
        ems.displayAllEmployees();
        
        // Test 9: Final Payroll
        System.out.println("--- Test 9: Final Payroll ---");
        System.out.println("Final Total Payroll: $" + ems.calculateTotalPayroll());
        
        System.out.println("\n=== Demo Complete ===");
        System.out.println("To run the interactive menu, run EmployeeManagementSystem.main()");
    }
}
