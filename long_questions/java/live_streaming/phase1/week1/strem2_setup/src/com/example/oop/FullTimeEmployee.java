package com.example.oop;

/**
 * FullTimeEmployee extends abstract class and implements interfaces
 * Demonstrates using both inheritance and interface implementation
 */
public class FullTimeEmployee extends AbstractEmployee 
                              implements Workable, Billable {
    
    private String department;
    private double billableHours;
    private double hourlyRate;
    
    public FullTimeEmployee(int id, String name, double baseSalary,
                            String department) {
        super(id, name, baseSalary);
        this.department = department;
        this.hourlyRate = baseSalary / 160; // Assume 160 hours/month
        this.billableHours = 160;
    }
    
    // Implement abstract method from AbstractEmployee
    @Override
    public double calculateSalary() {
        return baseSalary;
    }
    
    @Override
    public String getJobTitle() {
        return "Full-Time Employee";
    }
    
    // Implement Workable interface
    @Override
    public void work() {
        System.out.println(name + " is working full-time in " + department);
    }
    
    @Override
    public void takeBreak() {
        System.out.println(name + " is taking a scheduled break");
    }
    
    // Override default method if needed
    @Override
    public void attendMeeting(String meetingTopic) {
        System.out.println(name + " (Full-Time) attending: " + meetingTopic);
    }
    
    // Implement Billable interface
    @Override
    public double calculateBillableHours() {
        return billableHours;
    }
    
    @Override
    public double getHourlyRate() {
        return hourlyRate;
    }
    
    @Override
    public double generateInvoice() {
        double invoice = super.calculateSalary(); // Use base salary
        System.out.println("Invoice generated for full-time employee: $" + invoice);
        return invoice;
    }
    
    // Getter and setter methods
    public String getDepartment() {
        return department;
    }
    
    public void setDepartment(String department) {
        this.department = department;
    }
    
    public void setBillableHours(double billableHours) {
        this.billableHours = billableHours;
    }
    
    @Override
    public String toString() {
        return "FullTimeEmployee[id=" + id + ", name=" + name + 
               ", salary=$" + baseSalary + ", department=" + department + "]";
    }
}
