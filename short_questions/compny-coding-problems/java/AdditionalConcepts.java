package java_solutions;

public class AdditionalConcepts {

    // --- OOPS CONCEPTS ---

    // 81. Method Overloading
    static class Calculator {
        int add(int a, int b) {
            return a + b;
        }

        double add(double a, double b) {
            return a + b;
        }
    }

    // 82. Method Overriding
    static class Animal {
        void sound() {
            System.out.println("Animal makes a sound");
        }
    }

    static class Dog extends Animal {
        @Override
        void sound() {
            System.out.println("Dog barks");
        }
    }

    // 83. Encapsulation
    static class Person {
        private String name;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }
    }

    // 85. Polymorphism
    interface Shape {
        void draw();
    }

    static class Circle implements Shape {
        public void draw() {
            System.out.println("Drawing Circle");
        }
    }

    static class Square implements Shape {
        public void draw() {
            System.out.println("Drawing Square");
        }
    }

    // 89. Singleton Pattern
    static class Singleton {
        private static Singleton instance;

        private Singleton() {
        }

        public static Singleton getInstance() {
            if (instance == null)
                instance = new Singleton();
            return instance;
        }
    }

    // 90. Immutable Class
    final static class Immutable {
        private final int value;

        public Immutable(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }
    }

    // --- SQL & BASIC CS (Queries as String constants) ---

    public static final String Q91_SecondHighestSalary = "SELECT MAX(Salary) FROM Employee WHERE Salary < (SELECT MAX(Salary) FROM Employee);";
    public static final String Q92_DeleteVsTruncate = "Delete is DML (can rollback), Truncate is DDL (cannot rollback, faster).";
    public static final String Q93_Joins = "Inner Join: Matches only. Left Join: All left + matches right.";
    public static final String Q94_DuplicateRows = "SELECT col, COUNT(*) FROM table GROUP BY col HAVING COUNT(*) > 1;";
    public static final String Q95_PKvsFK = "PK: Unique identifier. FK: Reference to PK of another table.";
    public static final String Q96_Normalization = "Organizing data to reduce redundancy (1NF, 2NF, 3NF).";
    public static final String Q97_Index = "Data structure to improve speed of data retrieval (B-Tree).";
    public static final String Q98_ACID = "Atomicity, Consistency, Isolation, Durability.";
    public static final String Q99_Deadlock = "Two processes waiting for each other to release resources.";
    public static final String Q100_GetVsPost = "GET requests data (idempotent), POST submits data (not idempotent).";

    public static void main(String[] args) {
        System.out.println("--- OOPS CONCEPTS ---");
        Calculator calc = new Calculator();
        System.out.println("Overloading: " + calc.add(1, 2) + " vs " + calc.add(1.5, 2.5));

        Animal myDog = new Dog();
        myDog.sound(); // Overriding

        Person p = new Person();
        p.setName("John");
        System.out.println("Encapsulation: " + p.getName());

        Shape s1 = new Circle();
        Shape s2 = new Square();
        s1.draw();
        s2.draw(); // Polymorphism

        Singleton sing1 = Singleton.getInstance();
        Singleton sing2 = Singleton.getInstance();
        System.out.println("Singleton Same Instance: " + (sing1 == sing2));

        System.out.println("\n--- SQL & CS QUESTIONS ---");
        System.out.println("91. Second Highest Salary: " + Q91_SecondHighestSalary);
        System.out.println("92. Delete vs Truncate: " + Q92_DeleteVsTruncate);
        System.out.println("93. Joins: " + Q93_Joins);
        System.out.println("94. Duplicate Rows: " + Q94_DuplicateRows);
        System.out.println("95. PK vs FK: " + Q95_PKvsFK);
        System.out.println("98. ACID: " + Q98_ACID);
        System.out.println("100. GET vs POST: " + Q100_GetVsPost);
    }
}
