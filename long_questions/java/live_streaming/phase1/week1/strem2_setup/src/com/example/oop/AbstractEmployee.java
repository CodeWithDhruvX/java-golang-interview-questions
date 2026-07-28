package com.example.oop;

/**
 * Abstract class for all employees
 * Provides common implementation while requiring subclasses to implement specific methods
 * Demonstrates abstraction with both concrete and abstract methods
 */
public abstract class AbstractEmployee {
    
    // Protected fields (accessible by subclasses)
    protected int id;
    protected String name;
    protected double baseSalary;
    
    // Constructor (abstract classes can have constructors)
    public AbstractEmployee(int id, String name, double baseSalary) {
        this.id = id;
        this.name = name;
        this.baseSalary = baseSalary;
    }
    
    // Concrete method (has implementation)
    public void displayBasicInfo() {
        System.out.println("ID: " + id + ", Name: " + name);
    }
    
    // Abstract method (no implementation - must be overridden)
    public abstract double calculateSalary();
    
    public abstract String getJobTitle();
    
    // Concrete method that uses abstract methods (template method pattern)
    public void generatePaycheck() {
        System.out.println("=== PAYCHECK ===");
        displayBasicInfo();
        System.out.println("Position: " + getJobTitle());
        System.out.println("Amount: $" + calculateSalary());
        System.out.println("================");
    }
    
    // Getter methods
    public int getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    public double getBaseSalary() {
        return baseSalary;
    }
    
    // Setter methods
    public void setId(int id) {
        this.id = id;
    }
    
    public void setName(String name) {
        this.name = name;
    }
    
    public void setBaseSalary(double baseSalary) {
        this.baseSalary = baseSalary;
    }
}
