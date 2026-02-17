# Mini-Project 2: Student Management System (File I/O)

**Goal**: Demonstrate File Handling (Reading/Writing data), Serialization, and basic CRUD operations.

## Features
1.  Add Student
2.  Display All Students (Read from file)
3.  Save Data Persistence (Data survives app restart)

## Code Implementation

```java
import java.io.*;
import java.util.*;

// Student Class (Serializable to save object to file)
class Student implements Serializable {
    private static final long serialVersionUID = 1L;
    int id;
    String name;
    double marks;

    public Student(int id, String name, double marks) {
        this.id = id;
        this.name = name;
        this.marks = marks;
    }

    @Override
    public String toString() {
        return "ID: " + id + ", Name: " + name + ", Marks: " + marks;
    }
}

public class StudentManager {
    private static final String FILE_NAME = "students.dat";
    private static List<Student> students = new ArrayList<>();

    public static void main(String[] args) {
        loadData(); // Load existing data on startup
        Scanner sc = new Scanner(System.in);

        while (true) {
            System.out.println("\n1. Add Student\n2. Show All\n3. Save & Exit");
            System.out.print("Choice: ");
            int ch = sc.nextInt();

            switch (ch) {
                case 1:
                    System.out.print("ID: "); int id = sc.nextInt();
                    System.out.print("Name: "); String name = sc.next();
                    System.out.print("Marks: "); double marks = sc.nextDouble();
                    students.add(new Student(id, name, marks));
                    break;
                case 2:
                    if (students.isEmpty()) System.out.println("No records found.");
                    else students.forEach(System.out::println);
                    break;
                case 3:
                    saveData();
                    System.out.println("Data saved. Exiting...");
                    System.exit(0);
            }
        }
    }

    // Save list to file
    private static void saveData() {
        try (ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(FILE_NAME))) {
            oos.writeObject(students);
        } catch (IOException e) {
            System.out.println("Error saving data: " + e.getMessage());
        }
    }

    // Load list from file
    @SuppressWarnings("unchecked")
    private static void loadData() {
        File file = new File(FILE_NAME);
        if (!file.exists()) return;

        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream(file))) {
            students = (List<Student>) ois.readObject();
        } catch (IOException | ClassNotFoundException e) {
            System.out.println("Error loading data: " + e.getMessage());
        }
    }
}
```

## Key Code Concepts Used
*   **Serialization**: `Serializable` interface to save objects directly.
*   **File I/O**: `ObjectOutputStream` (Write) and `ObjectInputStream` (Read).
*   **Collections**: `ArrayList` to hold data in memory.
*   **Persistency**: Data remains available even after terminating the program.
