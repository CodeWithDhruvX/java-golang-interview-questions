package com.example.oop;

/**
 * Developer is a specialized type of Employee
 * Inherits all Employee properties and behaviors
 * Adds developer-specific features
 * Demonstrates inheritance with array fields and custom logic
 */
public class Developer extends Employee {
    
    // Developer-specific properties
    private String[] programmingLanguages;
    private int yearsOfExperience;
    
    // Constructor
    public Developer(int id, String name, double salary, 
                      String department, String[] languages, int experience) {
        super(id, name, salary, department);
        this.programmingLanguages = languages;
        this.yearsOfExperience = experience;
    }
    
    // Developer-specific method
    public void writeCode(String language) {
        System.out.println(getName() + " is writing code in " + language);
    }
    
    // Override giveRaise with experience-based logic
    @Override
    public void giveRaise(double percentage) {
        // Developers get higher raises based on experience
        double adjustedPercentage = percentage + (yearsOfExperience * 0.5);
        super.giveRaise(adjustedPercentage);
        System.out.println("Experience bonus applied: +" + (yearsOfExperience * 0.5) + "%");
    }
    
    // Override getJobTitle
    @Override
    public String getJobTitle() {
        return "Developer";
    }
    
    // Override toString
    @Override
    public String toString() {
        return super.toString() + 
               ", languages=" + String.join(", ", programmingLanguages) +
               ", experience=" + yearsOfExperience + " years]";
    }
    
    // Getter and setter methods
    public String[] getProgrammingLanguages() {
        return programmingLanguages;
    }
    
    public void setProgrammingLanguages(String[] programmingLanguages) {
        this.programmingLanguages = programmingLanguages;
    }
    
    public int getYearsOfExperience() {
        return yearsOfExperience;
    }
    
    public void setYearsOfExperience(int yearsOfExperience) {
        this.yearsOfExperience = yearsOfExperience;
    }
    
    // Helper method to check if developer knows a language
    public boolean knowsLanguage(String language) {
        for (String lang : programmingLanguages) {
            if (lang.equalsIgnoreCase(language)) {
                return true;
            }
        }
        return false;
    }
}
