public class demo {
    public static void main(String[] args) {
        int age = -5;

        // We assume age will always be non-negative
        assert age >= 0 : "Age cannot be negative! Current value: " + age;

        System.out.println("Age is: " + age);
    }
}