package com.example.oop;

/**
 * Manager is a specialized type of Employee
 * Inherits all Employee properties and behaviors
 * Adds manager-specific features
 * Demonstrates inheritance and method overriding
 */
public class Manager extends Employee {
    
    // Manager-specific properties
    private double bonus;
    private String team;
    
    // Constructor
    public Manager(int id, String name, double salary, 
                   String department, double bonus, String team) {
        // Call parent constructor using super()
        super(id, name, salary, department);
        this.bonus = bonus;
        this.team = team;
    }
    
    // Manager-specific method
    public void assignProject(String projectName) {
        System.out.println(getName() + " is assigning project: " + projectName + 
                          " to team: " + team);
    }
    
    // Override parent method (runtime polymorphism)
    @Override
    public void giveRaise(double percentage) {
        super.giveRaise(percentage);  // Call parent method
        this.bonus += 1000;  // Add manager-specific bonus
        System.out.println("Manager bonus increased by $1000!");
    }
    
    // Override getJobTitle
    @Override
    public String getJobTitle() {
        return "Manager";
    }
    
    // Override toString
    @Override
    public String toString() {
        return super.toString() + 
               ", bonus=$" + bonus + 
               ", team=" + team + "]";
    }
    
    // Getter and setter methods
    public double getBonus() {
        return bonus;
    }
    
    public void setBonus(double bonus) {
        this.bonus = bonus;
    }
    
    public String getTeam() {
        return team;
    }
    
    public void setTeam(String team) {
        this.team = team;
    }
}
