import java.util.*;

/**
 * Demonstrates Comparable vs Comparator.
 * 
 * Comparable: Natural ordering (compareTo)
 * Comparator: Custom ordering (compare)
 */
class Student implements Comparable<Student> {
    int id;
    String name;

    public Student(int id, String name) {
        this.id = id;
        this.name = name;
    }

    // Natural ordering by ID
    @Override
    public int compareTo(Student other) {
        return this.id - other.id;
    }

    @Override
    public String toString() {
        return "Student{id=" + id + ", name='" + name + "'}";
    }
}

public class ComparableComparatorDemo {

    public static void main(String[] args) {
        List<Student> students = new ArrayList<>();
        students.add(new Student(3, "Alice"));
        students.add(new Student(1, "Charlie"));
        students.add(new Student(2, "Bob"));

        System.out.println("Original: " + students);

        // Sort using Comparable (Natural Order: ID)
        Collections.sort(students);
        System.out.println("Sorted by ID (Comparable): " + students);

        // Sort using Comparator (Custom Order: Name)
        // Using Lambda Expression for Comparator
        Collections.sort(students, (s1, s2) -> s1.name.compareTo(s2.name));
        System.out.println("Sorted by Name (Comparator): " + students);

        // Comparator.comparing (Java 8+)
        students.sort(Comparator.comparing(s -> s.name));
        System.out.println("Sorted by Name (Comparator.comparing): " + students);

        // Reversed Comparator
        students.sort(Comparator.comparing(Student::getName).reversed());
        // Note: Java 8 method reference Student::getName (needs getter if field
        // private, here package-private access works)
    }

    // Helper for method ref
    static String getName(Student s) {
        return s.name;
    }
}
