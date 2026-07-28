package com.example.oop;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Complete Employee Management System
 * Demonstrates all OOP principles: encapsulation, inheritance, polymorphism, abstraction
 * Provides CRUD operations with interactive menu
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
    
    // Delete employee
    public boolean deleteEmployee(int id) {
        Employee emp = findEmployeeById(id);
        if (emp != null) {
            employees.remove(emp);
            System.out.println("Employee deleted: " + emp.getName());
            return true;
        }
        System.out.println("Employee not found with ID: " + id);
        return false;
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
            System.out.println("7. Delete Employee");
            System.out.println("8. Exit");
            System.out.print("Choose option: ");
            
            try {
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
                        deleteEmployeeMenu();
                        break;
                    case 8:
                        System.out.println("Exiting system...");
                        scanner.close();
                        return;
                    default:
                        System.out.println("Invalid option!");
                }
            } catch (Exception e) {
                System.out.println("Error: " + e.getMessage());
                scanner.nextLine(); // Clear invalid input
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
    
    private void deleteEmployeeMenu() {
        System.out.print("\nEnter Employee ID to delete: ");
        int id = scanner.nextInt();
        scanner.nextLine();
        
        deleteEmployee(id);
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
