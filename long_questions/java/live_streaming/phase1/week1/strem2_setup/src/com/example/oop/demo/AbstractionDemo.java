package com.example.oop.demo;

import com.example.oop.AbstractEmployee;
import com.example.oop.FullTimeEmployee;
import com.example.oop.Contractor;
import com.example.oop.Workable;
import com.example.oop.Billable;

/**
 * Demonstration of abstract classes and interfaces
 * Run this class to understand abstraction concepts
 */
public class AbstractionDemo {
    
    public static void main(String[] args) {
        System.out.println("=== Abstraction Demo ===\n");
        
        // Abstract class demonstration
        System.out.println("--- Abstract Class Demo ---");
        FullTimeEmployee ft = new FullTimeEmployee(101, "Alice", 6000, "IT");
        ft.displayBasicInfo();  // Concrete method from abstract class
        ft.generatePaycheck();  // Uses abstract methods
        System.out.println();
        
        // Interface demonstration
        System.out.println("--- Interface Demo ---");
        Contractor contractor = new Contractor("Bob", "TechCorp", 50.0);
        contractor.logHours(40);
        contractor.logHours(35);
        contractor.work();  // From Workable interface
        contractor.takeBreak();  // From Workable interface
        contractor.attendMeeting("Client Call");  // Default method from interface
        contractor.submitTimesheet();  // Default method from interface
        System.out.println();
        
        // Billable interface
        System.out.println("--- Billable Interface Demo ---");
        System.out.println("Contractor Invoice:");
        contractor.printInvoiceDetails();
        System.out.println();
        
        // FullTimeEmployee implements both abstract class and interfaces
        System.out.println("--- FullTimeEmployee (Abstract + Interfaces) ---");
        ft.work();  // From Workable
        ft.takeBreak();  // From Workable
        ft.attendMeeting("Team Standup");  // Overridden default method
        ft.generateInvoice();  // From Billable (overridden)
        System.out.println();
        
        // Interface static methods
        System.out.println("--- Interface Static Methods ---");
        Workable.displayWorkPolicy();
        System.out.println("Valid 40 hours: " + Workable.isValidWorkHours(40));
        System.out.println("Valid 25 hours: " + Workable.isValidWorkHours(25));
        System.out.println("Valid 50 hours: " + Workable.isValidWorkHours(50));
        System.out.println();
        
        // Billable static method
        System.out.println("--- Billable Static Method ---");
        double amount = 1000.0;
        double tax = Billable.calculateTax(amount, 10);
        System.out.println("Amount: $" + amount);
        System.out.println("Tax (10%): $" + tax);
        System.out.println("After tax: $" + (amount - tax));
        System.out.println();
        
        // Polymorphism with interfaces
        System.out.println("--- Polymorphism with Interfaces ---");
        Workable[] workers = {ft, contractor};
        System.out.println("All workers:");
        for (Workable worker : workers) {
            worker.work();
        }
        System.out.println();
        
        Billable[] billables = {ft, contractor};
        System.out.println("All billable entities:");
        for (Billable billable : billables) {
            System.out.println("Invoice: $" + billable.generateInvoice());
        }
        System.out.println();
        
        // Demonstrating when to use abstract class vs interface
        System.out.println("--- Abstract Class vs Interface Summary ---");
        System.out.println("AbstractEmployee (Abstract Class):");
        System.out.println("  - Has state (id, name, baseSalary)");
        System.out.println("  - Has constructor");
        System.out.println("  - Provides common implementation");
        System.out.println("  - Single inheritance");
        System.out.println();
        System.out.println("Workable (Interface):");
        System.out.println("  - No state (only constants)");
        System.out.println("  - No constructor");
        System.out.println("  - Defines contract only");
        System.out.println("  - Multiple implementation allowed");
        System.out.println();
        System.out.println("FullTimeEmployee:");
        System.out.println("  - Extends AbstractEmployee (shares state/implementation)");
        System.out.println("  - Implements Workable and Billable (multiple behaviors)");
        System.out.println();
        System.out.println("Contractor:");
        System.out.println("  - Does NOT extend AbstractEmployee (no shared state)");
        System.out.println("  - Implements Workable and Billable (behaviors only)");
    }
}
