package com.example.oop;

/**
 * Contractor implements interfaces but doesn't extend abstract class
 * Demonstrates that interfaces can be implemented by any class
 */
public class Contractor implements Workable, Billable {
    
    private String name;
    private String company;
    private double hourlyRate;
    private double hoursWorked;
    
    public Contractor(String name, String company, double hourlyRate) {
        this.name = name;
        this.company = company;
        this.hourlyRate = hourlyRate;
        this.hoursWorked = 0;
    }
    
    // Implement Workable interface
    @Override
    public void work() {
        System.out.println(name + " from " + company + " is working on contract");
    }
    
    @Override
    public void takeBreak() {
        System.out.println(name + " is taking a break (unpaid)");
    }
    
    // Override default method
    @Override
    public void attendMeeting(String meetingTopic) {
        System.out.println(name + " (Contractor) attending: " + meetingTopic);
    }
    
    // Implement Billable interface
    @Override
    public double calculateBillableHours() {
        return hoursWorked;
    }
    
    @Override
    public double getHourlyRate() {
        return hourlyRate;
    }
    
    // Contractor-specific method
    public void logHours(double hours) {
        this.hoursWorked += hours;
        System.out.println("Logged " + hours + " hours. Total: " + hoursWorked);
    }
    
    public void resetHours() {
        this.hoursWorked = 0;
        System.out.println("Hours reset for new billing period");
    }
    
    // Getter methods
    public String getName() {
        return name;
    }
    
    public String getCompany() {
        return company;
    }
    
    public double getHoursWorked() {
        return hoursWorked;
    }
    
    @Override
    public String toString() {
        return "Contractor[name=" + name + ", company=" + company + 
               ", rate=$" + hourlyRate + "/hour, hours=" + hoursWorked + "]";
    }
}
