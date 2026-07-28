package com.example.oop;

/**
 * Represents an employee in the company
 * This is our blueprint for creating employee objects
 * Demonstrates encapsulation, constructors, and methods
 */
public class Employee {
    
    // Instance variables (private - encapsulation)
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
    
    // Overloaded methods (compile-time polymorphism)
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
    
    // Business logic method
    public void giveRaise(double percentage) {
        if (percentage > 0 && percentage <= 50) {
            this.salary = this.salary * (1 + percentage / 100);
            System.out.println(name + " received a " + percentage + "% raise!");
        } else {
            System.out.println("Invalid raise percentage!");
        }
    }
    
    // Computed property (no backing field)
    public double getAnnualSalary() {
        return salary * 12;
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
    
    // Setter methods with validation (encapsulation)
    public void setId(int id) {
        this.id = id;
    }
    
    public void setName(String name) {
        if (name != null && !name.trim().isEmpty()) {
            this.name = name;
        } else {
            throw new IllegalArgumentException("Name cannot be empty!");
        }
    }
    
    public void setSalary(double salary) {
        if (salary > 0) {
            this.salary = salary;
        } else {
            throw new IllegalArgumentException("Salary must be positive!");
        }
    }
    
    public void setDepartment(String department) {
        this.department = department;
    }
    
    // toString method for easy printing
    @Override
    public String toString() {
        return "Employee[id=" + id + ", name=" + name + 
               ", salary=$" + salary + ", department=" + department + "]";
    }
    
    // Helper method for subclasses
    public String getJobTitle() {
        return "Employee";
    }
}
