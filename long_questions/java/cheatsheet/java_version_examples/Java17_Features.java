public class Java17_Features {

    // Java 17: Sealed Classes (Standard)
    // Permits specific classes to extend/implement
    public sealed interface Vehicle permits Car, Truck {}

    public static final record Car(String model) implements Vehicle {}
    public static final record Truck(int capacity) implements Vehicle {}

    public static void main(String[] args) {
        Vehicle myCar = new Car("Tesla");
        
        System.out.println("Sealed Class Hierarchy usage: " + myCar);
        
        // Pattern Matching for switch (Preview in 17)
        // Note: Standard Switch works normally
        if (myCar instanceof Car c) {
            System.out.println("It's a car: " + c.model());
        }

        System.out.println("Java 17 Features: Sealed Classes, Pattern Matching Standard, Strong Encapsulation, Hex Format API.");
    }
}
