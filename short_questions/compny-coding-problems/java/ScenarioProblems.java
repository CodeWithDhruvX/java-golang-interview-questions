package java_solutions;

import java.util.*;

public class ScenarioProblems {

    // 91. Exception Handling Logic
    public static void exceptionHandlingDemo() {
        try {
            int result = 10 / 0;
        } catch (ArithmeticException e) {
            System.out.println("Caught exception: " + e.getMessage());
        } finally {
            System.out.println("Finally block executed");
        }
    }

    // 92. Method Overloading Logic
    public static int add(int a, int b) {
        return a + b;
    }

    public static int add(int a, int b, int c) {
        return a + b + c;
    }

    public static double add(double a, double b) {
        return a + b;
    }

    // 93. Pass by Value vs Pass by Reference (Java is always Pass by Value)
    public static void modifyValue(int val, StringBuilder ref) {
        val = 100;
        ref.append(" World");
    }

    // 94. String Pool Logic
    public static void stringIdentity() {
        String s1 = "Hello";
        String s2 = "Hello";
        String s3 = new String("Hello");
        System.out.println("s1 == s2: " + (s1 == s2)); // True (String Pool)
        System.out.println("s1 == s3: " + (s1 == s3)); // False (Heap)
        System.out.println("s1.equals(s3): " + s1.equals(s3)); // True (Content)
    }

    // 95. Switch Case Logic
    public static void checkGrade(int score) {
        switch (score / 10) {
            case 10:
            case 9:
                System.out.println("A");
                break;
            case 8:
                System.out.println("B");
                break;
            case 7:
                System.out.println("C");
                break;
            default:
                System.out.println("F");
        }
    }

    // 96. Static keyword logic
    static int count = 0;

    public static void counter() {
        count++;
        System.out.println("Count: " + count);
    }

    // 97. Final keyword logic
    public static void finalDemo() {
        final double PI = 3.14159;
        // PI = 3.14; // Compile Error
        System.out.println("PI: " + PI);
    }

    // 98. Object Cloning
    static class Person implements Cloneable {
        String name;
        int age;

        Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        @Override
        protected Object clone() throws CloneNotSupportedException {
            return super.clone();
        }

        @Override
        public String toString() {
            return name + " (" + age + ")";
        }
    }

    public static void cloneDemo() {
        try {
            Person p1 = new Person("John", 30);
            Person p2 = (Person) p1.clone();
            p2.name = "Doe";
            System.out.println("p1: " + p1); // John
            System.out.println("p2: " + p2); // Doe
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
        }
    }

    // 99. Constructor Logic
    static class Employee {
        String name;
        int id;

        Employee(String name, int id) {
            this.name = name;
            this.id = id;
        }

        // Constructor Chaining
        Employee(String name) {
            this(name, 0);
        }
    }

    // 100. Immutable Class Logic
    final static class ImmutablePoint {
        private final int x;
        private final int y;

        public ImmutablePoint(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public int getX() {
            return x;
        }

        public int getY() {
            return y;
        }
    }

    public static void main(String[] args) {
        System.out.println("91. Exception Handling:");
        exceptionHandlingDemo();

        System.out.println("\n92. Overloading:");
        System.out.println("Add(1, 2): " + add(1, 2));
        System.out.println("Add(1.5, 2.5): " + add(1.5, 2.5));

        System.out.println("\n93. Pass by Value vs Ref:");
        int v = 1;
        StringBuilder r = new StringBuilder("Hello");
        modifyValue(v, r);
        System.out.println("Val: " + v + " (Unchanged), Ref: " + r + " (Changed content)");

        System.out.println("\n94. String Identity:");
        stringIdentity();

        System.out.println("\n95. Switch Case:");
        checkGrade(85);

        System.out.println("\n96. Static:");
        counter();
        counter();

        System.out.println("\n97. Final:");
        finalDemo();

        System.out.println("\n98. Cloning:");
        cloneDemo();

        System.out.println("\n99. Constructor:");
        Employee emp = new Employee("Alice");
        System.out.println("Emp: " + emp.name + ", " + emp.id);

        System.out.println("\n100. Immutable Point:");
        ImmutablePoint ip = new ImmutablePoint(10, 20);
        System.out.println("X=" + ip.getX() + ", Y=" + ip.getY());
    }
}
