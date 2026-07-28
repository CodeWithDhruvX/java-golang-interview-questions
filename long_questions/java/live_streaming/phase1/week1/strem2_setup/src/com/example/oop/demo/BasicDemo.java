package com.example.oop.demo;

import com.example.oop.Employee;

/**
 * Basic demonstration of classes, objects, and constructors
 * Run this class to understand fundamental OOP concepts
 */
public class BasicDemo {
    
    public static void main(String[] args) {
        System.out.println("=== Basic OOP Demo ===\n");
        
        // Using default constructor
        System.out.println("--- Creating Employee with Default Constructor ---");
        Employee emp1 = new Employee();
        emp1.displayDetails();
        System.out.println();
        
        // Using parameterized constructor
        System.out.println("--- Creating Employee with Parameterized Constructor ---");
        Employee emp2 = new Employee(101, "John Doe", 75000.0, "Engineering");
        emp2.displayDetails();
        System.out.println();
        
        // Using copy constructor
        System.out.println("--- Creating Employee with Copy Constructor ---");
        Employee emp3 = new Employee(emp2);
        emp3.setName("Jane Smith");
        emp3.setId(102);
        emp3.displayDetails();
        System.out.println();
        
        // Using setters to modify
        System.out.println("--- Modifying Employee with Setters ---");
        emp1.setId(103);
        emp1.setName("Bob Johnson");
        emp1.setSalary(65000.0);
        emp1.setDepartment("Marketing");
        emp1.displayDetails();
        System.out.println();
        
        // Demonstrating getters
        System.out.println("--- Using Getters ---");
        System.out.println("Employee ID: " + emp2.getId());
        System.out.println("Employee Name: " + emp2.getName());
        System.out.println("Employee Salary: $" + emp2.getSalary());
        System.out.println("Employee Department: " + emp2.getDepartment());
        System.out.println("Annual Salary: $" + emp2.getAnnualSalary());
        System.out.println();
        
        // Demonstrating method overloading
        System.out.println("--- Method Overloading Demo ---");
        emp2.work();
        emp2.work(8);
        emp2.work("Project X");
        emp2.work(6, "Project Y");
        System.out.println();
        
        // Demonstrating toString
        System.out.println("--- toString Method ---");
        System.out.println(emp2);
        System.out.println();
        
        // Demonstrating giveRaise
        System.out.println("--- Give Raise Method ---");
        System.out.println("Before raise: $" + emp2.getSalary());
        emp2.giveRaise(10);
        System.out.println("After raise: $" + emp2.getSalary());
        System.out.println();
        
        // Demonstrating validation in setters
        System.out.println("--- Validation in Setters ---");
        try {
            emp2.setSalary(-1000);
        } catch (IllegalArgumentException e) {
            System.out.println("Caught error: " + e.getMessage());
        }
        
        try {
            emp2.setName("");
        } catch (IllegalArgumentException e) {
            System.out.println("Caught error: " + e.getMessage());
        }
    }
}
