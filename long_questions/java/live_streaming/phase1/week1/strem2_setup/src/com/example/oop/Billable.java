package com.example.oop;

/**
 * Interface for billable employees/contractors
 * Defines contract for billing-related operations
 * Demonstrates interface with default implementation
 */
public interface Billable {
    
    // Abstract methods
    double calculateBillableHours();
    
    double getHourlyRate();
    
    // Default method with implementation
    default double generateInvoice() {
        double hours = calculateBillableHours();
        double rate = getHourlyRate();
        double invoice = hours * rate;
        System.out.println("Invoice generated: $" + invoice + 
                          " (" + hours + " hours @ $" + rate + "/hour)");
        return invoice;
    }
    
    // Another default method
    default void printInvoiceDetails() {
        System.out.println("=== Invoice Details ===");
        System.out.println("Billable Hours: " + calculateBillableHours());
        System.out.println("Hourly Rate: $" + getHourlyRate());
        System.out.println("Total Amount: $" + generateInvoice());
        System.out.println("======================");
    }
    
    // Static method
    static double calculateTax(double amount, double taxRate) {
        return amount * (taxRate / 100);
    }
}
