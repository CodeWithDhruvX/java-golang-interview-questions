package com.example.oop;

/**
 * Interface defining work-related behavior
 * Any class can implement this regardless of inheritance hierarchy
 * Demonstrates interface with default and static methods (Java 8+)
 */
public interface Workable {
    
    // Constant (public static final by default)
    int STANDARD_WORK_HOURS = 40;
    
    // Abstract methods (public abstract by default)
    void work();
    
    void takeBreak();
    
    // Default method (Java 8+) - provides implementation
    default void attendMeeting(String meetingTopic) {
        System.out.println("Attending meeting: " + meetingTopic);
    }
    
    // Another default method
    default void submitTimesheet() {
        System.out.println("Timesheet submitted for " + STANDARD_WORK_HOURS + " hours");
    }
    
    // Static method (Java 8+) - utility method
    static void displayWorkPolicy() {
        System.out.println("Work Policy: " + STANDARD_WORK_HOURS + " hours/week");
        System.out.println("Breaks: 1 hour per day");
        System.out.println("Overtime: Requires manager approval");
    }
    
    // Static utility method
    static boolean isValidWorkHours(int hours) {
        return hours >= 0 && hours <= 24;
    }
}
