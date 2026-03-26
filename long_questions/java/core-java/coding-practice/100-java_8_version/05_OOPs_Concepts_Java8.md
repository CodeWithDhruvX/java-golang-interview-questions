# OOPs Concepts with Java 8 Features (66-75)

## 📚 Java 8 Features Demonstrated
- **Lambda Expressions**: Concise method implementations
- **Method References**: Simplified method calls
- **Functional Interfaces**: Custom and built-in interfaces
- **Streams API**: Collection processing
- **Optional**: Null-safe operations
- **Default Methods**: Interface evolution
- **Static Methods in Interfaces**: Utility methods

---

## 66. Singleton Class
**Java 8 Approach**: Using lambda and functional interfaces

```java
import java.util.*;
import java.util.function.*;

public class SingletonJava8 {
    
    // Traditional Singleton with Java 8 enhancements
    static class Singleton {
        private static volatile Singleton instance;
        
        private Singleton() {
            System.out.println("Singleton instance created");
        }
        
        // Using Supplier for lazy initialization
        private static final Supplier<Singleton> LAZY_INSTANCE = Suppliers.memoize(() -> {
            if (instance == null) {
                synchronized (Singleton.class) {
                    if (instance == null) {
                        instance = new Singleton();
                    }
                }
            }
            return instance;
        });
        
        public static Singleton getInstance() {
            return LAZY_INSTANCE.get();
        }
        
        public void showMessage() {
            System.out.println("Hello from Singleton!");
        }
    }
    
    // Functional Singleton using Supplier
    static class FunctionalSingleton<T> {
        private final Supplier<T> supplier;
        private volatile T instance;
        
        public FunctionalSingleton(Supplier<T> supplier) {
            this.supplier = supplier;
        }
        
        public T getInstance() {
            if (instance == null) {
                synchronized (this) {
                    if (instance == null) {
                        instance = supplier.get();
                    }
                }
            }
            return instance;
        }
    }
    
    public static void main(String[] args) {
        // Traditional Singleton
        Singleton s1 = Singleton.getInstance();
        Singleton s2 = Singleton.getInstance();
        
        System.out.println("Traditional Singleton same instance? " + (s1 == s2));
        
        // Functional Singleton
        FunctionalSingleton<String> stringSingleton = new FunctionalSingleton<>(
            () -> "Functional Singleton Instance"
        );
        
        String str1 = stringSingleton.getInstance();
        String str2 = stringSingleton.getInstance();
        
        System.out.println("Functional Singleton same instance? " + (str1 == str2));
        System.out.println("Functional Singleton value: " + str1);
        
        // Using Optional for safe Singleton access
        Optional<Singleton> optionalSingleton = Optional.ofNullable(Singleton.getInstance());
        optionalSingleton.ifPresent(Singleton::showMessage);
        
        // Demonstrate thread safety with parallel streams
        List<Singleton> instances = IntStream.range(0, 100)
            .parallel()
            .mapToObj(i -> Singleton.getInstance())
            .distinct()
            .collect(Collectors.toList());
        
        System.out.println("Thread-safe Singleton instances created: " + instances.size());
    }
    
    // Utility class for memoization
    static class Suppliers {
        public static <T> Supplier<T> memoize(Supplier<T> supplier) {
            return new Supplier<T>() {
                private volatile T value;
                
                @Override
                public T get() {
                    if (value == null) {
                        synchronized (this) {
                            if (value == null) {
                                value = supplier.get();
                            }
                        }
                    }
                    return value;
                }
            };
        }
    }
}
```

## 67. Immutable Class
**Java 8 Approach**: Using streams and Optional

```java
import java.util.*;
import java.util.stream.*;

public class ImmutableJava8 {
    
    // Enhanced Immutable class with Java 8 features
    final class Person {
        private final String name;
        private final int age;
        private final List<String> hobbies;
        
        public Person(String name, int age, List<String> hobbies) {
            this.name = Objects.requireNonNull(name, "Name cannot be null");
            this.age = age;
            this.hobbies = Collections.unmodifiableList(new ArrayList<>(hobbies));
        }
        
        public String getName() { return name; }
        public int getAge() { return age; }
        public List<String> getHobbies() { return hobbies; }
        
        // Builder pattern with Java 8
        public static Builder builder() { return new Builder(); }
        
        // with methods for creating new instances
        public Person withName(String newName) {
            return new Person(newName, this.age, this.hobbies);
        }
        
        public Person withAge(int newAge) {
            return new Person(this.name, newAge, this.hobbies);
        }
        
        public Person withHobbies(List<String> newHobbies) {
            return new Person(this.name, this.age, newHobbies);
        }
        
        @Override
        public String toString() {
            return String.format("Person{name='%s', age=%d, hobbies=%s}", name, age, hobbies);
        }
        
        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Person person = (Person) o;
            return age == person.age && 
                   Objects.equals(name, person.name) && 
                   Objects.equals(hobbies, person.hobbies);
        }
        
        @Override
        public int hashCode() {
            return Objects.hash(name, age, hobbies);
        }
        
        // Static builder class
        static class Builder {
            private String name;
            private int age;
            private List<String> hobbies = new ArrayList<>();
            
            public Builder name(String name) {
                this.name = name;
                return this;
            }
            
            public Builder age(int age) {
                this.age = age;
                return this;
            }
            
            public Builder hobbies(String... hobbies) {
                this.hobbies = Arrays.asList(hobbies);
                return this;
            }
            
            public Person build() {
                return new Person(name, age, hobbies);
            }
        }
    }
    
    public static void main(String[] args) {
        // Using builder pattern
        Person person = Person.builder()
            .name("John")
            .age(30)
            .hobbies("Reading", "Swimming", "Coding")
            .build();
        
        System.out.println("Original person: " + person);
        
        // Creating new person with modifications
        Person modifiedPerson = person.withAge(31).withName("John Doe");
        System.out.println("Modified person: " + modifiedPerson);
        System.out.println("Original unchanged: " + person);
        
        // Using streams with immutable objects
        List<Person> people = Arrays.asList(
            Person.builder().name("Alice").age(25).hobbies("Music").build(),
            Person.builder().name("Bob").age(30).hobbies("Sports").build(),
            Person.builder().name("Charlie").age(35).hobbies("Reading").build()
        );
        
        // Filter and transform immutable objects
        List<String> adultNames = people.stream()
            .filter(p -> p.getAge() >= 30)
            .map(Person::getName)
            .collect(Collectors.toList());
        
        System.out.println("Adults: " + adultNames);
        
        // Group by age
        Map<Integer, List<Person>> peopleByAge = people.stream()
            .collect(Collectors.groupingBy(Person::getAge));
        
        System.out.println("People by age: " + peopleByAge);
        
        // Using Optional with immutable objects
        Optional<Person> optionalPerson = Optional.of(person);
        optionalPerson.map(Person::getName)
                    .ifPresent(name -> System.out.println("Person name: " + name));
        
        // Demonstrate immutability
        try {
            List<String> hobbies = person.getHobbies();
            hobbies.add("New Hobby"); // This should throw UnsupportedOperationException
        } catch (UnsupportedOperationException e) {
            System.out.println("Hobbies list is immutable: " + e.getMessage());
        }
        
        // Create immutable collections
        List<Person> immutablePeople = Collections.unmodifiableList(
            new ArrayList<>(people)
        );
        
        System.out.println("Immutable people list: " + immutablePeople);
        
        // Process immutable objects in parallel
        long adultCount = people.parallelStream()
            .filter(p -> p.getAge() >= 18)
            .count();
        
        System.out.println("Adult count: " + adultCount);
    }
}
```

## 68. Method Overloading
**Java 8 Approach**: Using functional interfaces and default methods

```java
import java.util.*;
import java.util.function.*;

public class MethodOverloadingJava8 {
    
    // Math utility class with overloaded methods and Java 8 features
    static class MathUtils {
        
        // Traditional overloading
        public int add(int a, int b) {
            return a + b;
        }
        
        public double add(double a, double b) {
            return a + b;
        }
        
        public int add(int a, int b, int c) {
            return a + b + c;
        }
        
        // Java 8 functional approach
        public int add(List<Integer> numbers) {
            return numbers.stream().mapToInt(Integer::intValue).sum();
        }
        
        public double add(List<Double> numbers) {
            return numbers.stream().mapToDouble(Double::doubleValue).sum();
        }
        
        // Generic addition using functional interfaces
        public <T> T add(T a, T b, BinaryOperator<T> adder) {
            return adder.apply(a, b);
        }
        
        // Varargs with streams
        public int add(int... numbers) {
            return Arrays.stream(numbers).sum();
        }
        
        public double add(double... numbers) {
            return Arrays.stream(numbers).sum();
        }
        
        // Optional-based addition
        public Optional<Integer> addOptional(Integer a, Integer b) {
            return Optional.ofNullable(a)
                .flatMap(valA -> Optional.ofNullable(b)
                    .map(valB -> valA + valB));
        }
        
        // Custom functional interface for addition
        @FunctionalInterface
        interface Adder<T> {
            T add(T a, T b);
        }
        
        public <T> T add(T a, T b, Adder<T> adder) {
            return adder.add(a, b);
        }
    }
    
    // Calculator with method overloading and lambda expressions
    static class Calculator {
        
        // Overloaded calculate methods
        public int calculate(int a, int b, String operation) {
            return switch (operation.toLowerCase()) {
                case "add" -> a + b;
                case "subtract" -> a - b;
                case "multiply" -> a * b;
                case "divide" -> b != 0 ? a / b : throw new ArithmeticException("Division by zero");
                default -> throw new IllegalArgumentException("Unknown operation: " + operation);
            };
        }
        
        public double calculate(double a, double b, String operation) {
            return switch (operation.toLowerCase()) {
                case "add" -> a + b;
                case "subtract" -> a - b;
                case "multiply" -> a * b;
                case "divide" -> b != 0 ? a / b : throw new ArithmeticException("Division by zero");
                case "power" -> Math.pow(a, b);
                default -> throw new IllegalArgumentException("Unknown operation: " + operation);
            };
        }
        
        // Functional approach with BinaryOperator
        public int calculate(int a, int b, BinaryOperator<Integer> operation) {
            return operation.apply(a, b);
        }
        
        public double calculate(double a, double b, BinaryOperator<Double> operation) {
            return operation.apply(a, b);
        }
        
        // Using Map of operations
        private static final Map<String, BinaryOperator<Integer>> INT_OPERATIONS = Map.of(
            "add", Integer::sum,
            "subtract", (a, b) -> a - b,
            "multiply", (a, b) -> a * b,
            "divide", (a, b) -> b != 0 ? a / b : throw new ArithmeticException("Division by zero")
        );
        
        private static final Map<String, BinaryOperator<Double>> DOUBLE_OPERATIONS = Map.of(
            "add", Double::sum,
            "subtract", (a, b) -> a - b,
            "multiply", (a, b) -> a * b,
            "divide", (a, b) -> b != 0 ? a / b : throw new ArithmeticException("Division by zero"),
            "power", Math::pow
        );
        
        public int calculateWithMap(int a, int b, String operation) {
            BinaryOperator<Integer> op = INT_OPERATIONS.get(operation.toLowerCase());
            if (op == null) {
                throw new IllegalArgumentException("Unknown operation: " + operation);
            }
            return op.apply(a, b);
        }
        
        public double calculateWithMap(double a, double b, String operation) {
            BinaryOperator<Double> op = DOUBLE_OPERATIONS.get(operation.toLowerCase());
            if (op == null) {
                throw new IllegalArgumentException("Unknown operation: " + operation);
            }
            return op.apply(a, b);
        }
    }
    
    public static void main(String[] args) {
        MathUtils math = new MathUtils();
        Calculator calc = new Calculator();
        
        // Traditional overloading
        System.out.println("Traditional overloading:");
        System.out.println("5 + 3 = " + math.add(5, 3));
        System.out.println("5.5 + 3.3 = " + math.add(5.5, 3.3));
        System.out.println("1 + 2 + 3 = " + math.add(1, 2, 3));
        
        // Java 8 functional approach
        System.out.println("\nFunctional approach:");
        List<Integer> intList = Arrays.asList(1, 2, 3, 4, 5);
        List<Double> doubleList = Arrays.asList(1.1, 2.2, 3.3, 4.4);
        
        System.out.println("Sum of list: " + math.add(intList));
        System.out.println("Sum of double list: " + math.add(doubleList));
        
        // Generic addition
        String result1 = math.add("Hello", " World", String::concat);
        System.out.println("String concatenation: " + result1);
        
        // Varargs
        System.out.println("Varargs sum: " + math.add(1, 2, 3, 4, 5));
        
        // Optional-based
        Optional<Integer> optionalResult = math.addOptional(5, 3);
        optionalResult.ifPresent(res -> System.out.println("Optional result: " + res));
        
        Optional<Integer> nullResult = math.addOptional(null, 3);
        System.out.println("Null result present: " + nullResult.isPresent());
        
        // Calculator examples
        System.out.println("\nCalculator examples:");
        System.out.println("5 + 3 = " + calc.calculate(5, 3, "add"));
        System.out.println("5.5^2.0 = " + calc.calculate(5.5, 2.0, "power"));
        
        // Using BinaryOperator
        int intResult = calc.calculate(10, 5, (a, b) -> a * b + b);
        System.out.println("Custom operation: " + intResult);
        
        // Using Map of operations
        System.out.println("Map-based calculation: " + calc.calculateWithMap(10, 2, "divide"));
        
        // Stream processing with overloaded methods
        List<String> operations = Arrays.asList("add", "subtract", "multiply", "divide");
        Map<String, Integer> results = operations.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                op -> calc.calculateWithMap(20, 5, op)
            ));
        
        System.out.println("Operations results: " + results);
        
        // Parallel processing
        List<Integer> numbers = IntStream.range(1, 101).boxed().collect(Collectors.toList());
        int totalSum = math.add(numbers);
        System.out.println("Total sum of 1-100: " + totalSum);
    }
}
```

## 69. Method Overriding
**Java 8 Approach**: Using default methods and functional interfaces

```java
import java.util.*;
import java.util.function.*;

public class MethodOverridingJava8 {
    
    // Base class with Java 8 features
    static class Animal {
        protected String name;
        
        public Animal(String name) {
            this.name = name;
        }
        
        // Method to be overridden
        public void makeSound() {
            System.out.println(name + " makes a sound");
        }
        
        // Default method (can be overridden)
        public void describe() {
            System.out.println(name + " is an animal");
        }
        
        // Functional method
        public void performAction(String action, Consumer<String> actionPerformer) {
            System.out.print(name + " ");
            actionPerformer.accept(action);
        }
        
        // Getter with Optional
        public Optional<String> getName() {
            return Optional.ofNullable(name);
        }
        
        // Template method pattern
        public void dailyRoutine() {
            wakeUp();
            eat();
            makeSound();
            sleep();
        }
        
        protected void wakeUp() {
            System.out.println(name + " wakes up");
        }
        
        protected void eat() {
            System.out.println(name + " eats");
        }
        
        protected void sleep() {
            System.out.println(name + " sleeps");
        }
    }
    
    // Subclass overriding methods
    static class Dog extends Animal {
        private String breed;
        
        public Dog(String name, String breed) {
            super(name);
            this.breed = breed;
        }
        
        @Override
        public void makeSound() {
            System.out.println(name + " barks: Woof! Woof!");
        }
        
        @Override
        public void describe() {
            System.out.println(name + " is a " + breed + " dog");
        }
        
        @Override
        protected void eat() {
            System.out.println(name + " eats dog food");
        }
        
        // Additional method
        public void wagTail() {
            System.out.println(name + " wags tail happily");
        }
        
        // Override with functional programming
        @Override
        public void performAction(String action, Consumer<String> actionPerformer) {
            if (action.equals("bark")) {
                makeSound();
            } else {
                super.performAction(action, actionPerformer);
            }
        }
    }
    
    // Another subclass
    static class Cat extends Animal {
        public Cat(String name) {
            super(name);
        }
        
        @Override
        public void makeSound() {
            System.out.println(name + " meows: Meow! Meow!");
        }
        
        @Override
        public void describe() {
            System.out.println(name + " is a cat");
        }
        
        @Override
        protected void sleep() {
            System.out.println(name + " sleeps for 16 hours");
        }
        
        public void purr() {
            System.out.println(name + " purrs contentedly");
        }
    }
    
    // Interface with default methods
    interface Flyable {
        default void fly() {
            System.out.println("Flying through the air");
        }
        
        default void takeOff() {
            System.out.println("Taking off");
        }
        
        default void land() {
            System.out.println("Landing safely");
        }
        
        // Static method
        static void showAllFlyable(List<Flyable> flyables) {
            flyables.forEach(Flyable::fly);
        }
    }
    
    // Class implementing interface and overriding
    static class Bird extends Animal implements Flyable {
        public Bird(String name) {
            super(name);
        }
        
        @Override
        public void makeSound() {
            System.out.println(name + " chirps: Tweet! Tweet!");
        }
        
        @Override
        public void fly() {
            System.out.println(name + " flies gracefully");
        }
        
        @Override
        public void dailyRoutine() {
            super.dailyRoutine();
            fly(); // Additional behavior
        }
    }
    
    // Functional interface for animal behavior
    @FunctionalInterface
    interface AnimalBehavior {
        void behave(Animal animal);
        
        // Default method
        default void describeBehavior(Animal animal) {
            System.out.print("Behavior: ");
            behave(animal);
        }
    }
    
    public static void main(String[] args) {
        // Demonstrate method overriding
        Animal genericAnimal = new Animal("Generic");
        Dog dog = new Dog("Buddy", "Golden Retriever");
        Cat cat = new Cat("Whiskers");
        Bird bird = new Bird("Tweety");
        
        System.out.println("=== Method Overriding Demo ===");
        
        List<Animal> animals = Arrays.asList(genericAnimal, dog, cat, bird);
        
        // Polymorphic behavior
        animals.forEach(Animal::makeSound);
        
        System.out.println("\n=== Descriptions ===");
        animals.forEach(Animal::describe);
        
        System.out.println("\n=== Daily Routines ===");
        animals.forEach(Animal::dailyRoutine);
        
        System.out.println("\n=== Specific Behaviors ===");
        dog.wagTail();
        cat.purr();
        
        System.out.println("\n=== Flying Behavior ===");
        List<Flyable> flyables = Arrays.asList(bird);
        Flyable.showAllFlyable(flyables);
        
        System.out.println("\n=== Functional Programming ===");
        AnimalBehavior dogBehavior = animal -> {
            if (animal instanceof Dog) {
                ((Dog) animal).wagTail();
            }
        };
        
        dogBehavior.describeBehavior(dog);
        
        // Using streams with overridden methods
        System.out.println("\n=== Stream Processing ===");
        animals.stream()
            .map(Animal::getName)
            .filter(Optional::isPresent)
            .map(Optional::get)
            .forEach(name -> System.out.println("Animal name: " + name));
        
        // Custom behavior with lambda
        AnimalBehavior customBehavior = animal -> {
            animal.performAction("jumps", action -> System.out.println("jumps excitedly!"));
        };
        
        animals.forEach(animal -> customBehavior.describeBehavior(animal));
        
        // Demonstrate interface default methods
        System.out.println("\n=== Interface Default Methods ===");
        bird.takeOff();
        bird.fly();
        bird.land();
        
        // Parallel processing with overridden methods
        System.out.println("\n=== Parallel Processing ===");
        animals.parallelStream()
            .forEach(animal -> {
                System.out.println(Thread.currentThread().getName() + ": ");
                animal.makeSound();
            });
    }
}
```

## 70. Interface Implementation
**Java 8 Approach**: Using default methods, static methods, and functional interfaces

```java
import java.util.*;
import java.util.function.*;

public class InterfaceImplementationJava8 {
    
    // Enhanced interface with Java 8 features
    interface Vehicle {
        // Abstract methods
        void start();
        void stop();
        String getType();
        
        // Default methods
        default void accelerate() {
            System.out.println(getType() + " is accelerating");
        }
        
        default void brake() {
            System.out.println(getType() + " is braking");
        }
        
        default void displayInfo() {
            System.out.println("Vehicle Type: " + getType());
            System.out.println("Status: " + (isRunning() ? "Running" : "Stopped"));
        }
        
        // Helper method
        default boolean isRunning() {
            return true; // Default implementation
        }
        
        // Static methods
        static void showAllVehicles(List<Vehicle> vehicles) {
            vehicles.forEach(vehicle -> System.out.println(vehicle.getType()));
        }
        
        static Vehicle createCar(String brand) {
            return new Car(brand);
        }
        
        static Vehicle createMotorcycle(String brand) {
            return new Motorcycle(brand);
        }
        
        // Functional interface method
        default void performAction(String action, Consumer<String> actionHandler) {
            System.out.print(getType() + " ");
            actionHandler.accept(action);
        }
    }
    
    // Car implementation
    static class Car implements Vehicle {
        private String brand;
        private boolean running = false;
        
        public Car(String brand) {
            this.brand = brand;
        }
        
        @Override
        public void start() {
            running = true;
            System.out.println(brand + " car started");
        }
        
        @Override
        public void stop() {
            running = false;
            System.out.println(brand + " car stopped");
        }
        
        @Override
        public String getType() {
            return "Car";
        }
        
        @Override
        public boolean isRunning() {
            return running;
        }
        
        @Override
        public void accelerate() {
            if (running) {
                System.out.println(brand + " car is speeding up");
            } else {
                System.out.println("Start the car first!");
            }
        }
        
        // Additional method
        public void honk() {
            System.out.println(brand + " car: Beep! Beep!");
        }
    }
    
    // Motorcycle implementation
    static class Motorcycle implements Vehicle {
        private String brand;
        private boolean running = false;
        
        public Motorcycle(String brand) {
            this.brand = brand;
        }
        
        @Override
        public void start() {
            running = true;
            System.out.println(brand + " motorcycle started");
        }
        
        @Override
        public void stop() {
            running = false;
            System.out.println(brand + " motorcycle stopped");
        }
        
        @Override
        public String getType() {
            return "Motorcycle";
        }
        
        @Override
        public boolean isRunning() {
            return running;
        }
        
        @Override
        public void accelerate() {
            if (running) {
                System.out.println(brand + " motorcycle is zooming");
            } else {
                System.out.println("Start the motorcycle first!");
            }
        }
        
        // Additional method
        public void wheelie() {
            if (running) {
                System.out.println(brand + " motorcycle does a wheelie!");
            } else {
                System.out.println("Start the motorcycle first!");
            }
        }
    }
    
    // Advanced interface with functional programming
    interface SmartVehicle extends Vehicle {
        // Additional abstract methods
        void connectToGPS();
        boolean isConnected();
        
        // Default methods
        default void navigate(String destination) {
            if (isConnected()) {
                System.out.println("Navigating to " + destination);
            } else {
                System.out.println("Connect to GPS first!");
            }
        }
        
        default void checkSystemStatus() {
            System.out.println("System Status:");
            System.out.println("- GPS: " + (isConnected() ? "Connected" : "Disconnected"));
            System.out.println("- Engine: " + (isRunning() ? "Running" : "Stopped"));
        }
        
        // Static factory method
        static SmartVehicle createSmartCar(String brand) {
            return new SmartCar(brand);
        }
    }
    
    // Smart car implementation
    static class SmartCar extends Car implements SmartVehicle {
        private boolean gpsConnected = false;
        
        public SmartCar(String brand) {
            super(brand);
        }
        
        @Override
        public void connectToGPS() {
            gpsConnected = true;
            System.out.println("GPS connected for smart car");
        }
        
        @Override
        public boolean isConnected() {
            return gpsConnected;
        }
        
        @Override
        public void displayInfo() {
            super.displayInfo();
            System.out.println("GPS Status: " + (gpsConnected ? "Connected" : "Disconnected"));
        }
    }
    
    // Functional interfaces for vehicle operations
    @FunctionalInterface
    interface VehicleOperation {
        void execute(Vehicle vehicle);
        
        default void describeOperation(Vehicle vehicle) {
            System.out.print("Executing operation on " + vehicle.getType() + ": ");
            execute(vehicle);
        }
    }
    
    @FunctionalInterface
    interface VehicleCondition {
        boolean check(Vehicle vehicle);
    }
    
    public static void main(String[] args) {
        // Create vehicles using static factory methods
        Vehicle car = Vehicle.createCar("Toyota");
        Vehicle motorcycle = Vehicle.createMotorcycle("Honda");
        SmartVehicle smartCar = SmartVehicle.createSmartCar("Tesla");
        
        List<Vehicle> vehicles = Arrays.asList(car, motorcycle, smartCar);
        
        System.out.println("=== Interface Implementation Demo ===");
        
        // Use static method
        Vehicle.showAllVehicles(vehicles);
        
        System.out.println("\n=== Vehicle Operations ===");
        
        // Start all vehicles
        vehicles.forEach(Vehicle::start);
        
        // Use default methods
        vehicles.forEach(Vehicle::accelerate);
        
        // Display info
        vehicles.forEach(Vehicle::displayInfo);
        
        System.out.println("\n=== Smart Vehicle Features ===");
        smartCar.connectToGPS();
        smartCar.navigate("Home");
        smartCar.checkSystemStatus();
        
        System.out.println("\n=== Functional Operations ===");
        
        // Custom operations using lambda
        VehicleOperation startAndHonk = vehicle -> {
            if (vehicle instanceof Car) {
                ((Car) vehicle).honk();
            }
        };
        
        vehicles.forEach(startAndHonk::describeOperation);
        
        // Filter and process vehicles
        VehicleCondition isRunning = Vehicle::isRunning;
        List<Vehicle> runningVehicles = vehicles.stream()
            .filter(isRunning::check)
            .collect(Collectors.toList());
        
        System.out.println("Running vehicles: " + runningVehicles.size());
        
        // Process with custom behavior
        vehicles.forEach(vehicle -> {
            vehicle.performAction("checking systems", 
                action -> System.out.println("is " + action));
        });
        
        System.out.println("\n=== Stream Processing ===");
        
        // Group vehicles by type
        Map<String, List<Vehicle>> vehiclesByType = vehicles.stream()
            .collect(Collectors.groupingBy(Vehicle::getType));
        
        vehiclesByType.forEach((type, vehicleList) -> {
            System.out.println(type + "s: " + vehicleList.size());
        });
        
        // Stop all vehicles
        System.out.println("\n=== Stopping Vehicles ===");
        vehicles.forEach(Vehicle::stop);
        
        // Parallel processing
        System.out.println("\n=== Parallel Operations ===");
        vehicles.parallelStream()
            .forEach(vehicle -> {
                System.out.println(Thread.currentThread().getName() + ": ");
                vehicle.displayInfo();
            });
        
        // Method references
        System.out.println("\n=== Method References ===");
        vehicles.forEach(System.out::println); // Will call toString()
        
        // Custom functional interface usage
        VehicleOperation complexOperation = Vehicle::start;
        complexOperation.describeOperation(car);
    }
}
```

## 71. Abstract Class
**Java 8 Approach**: Using functional interfaces and streams

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class AbstractClassJava8 {
    
    // Abstract class with Java 8 features
    abstract class Shape {
        protected String name;
        protected String color;
        
        public Shape(String name, String color) {
            this.name = name;
            this.color = color;
        }
        
        // Abstract methods
        public abstract double getArea();
        public abstract double getPerimeter();
        
        // Concrete methods
        public void displayInfo() {
            System.out.println(name + " (" + color + ")");
            System.out.println("Area: " + getArea());
            System.out.println("Perimeter: " + getPerimeter());
        }
        
        // Default method with functional programming
        public void compareWith(Shape other, Comparator<Double> comparator) {
            double thisArea = this.getArea();
            double otherArea = other.getArea();
            
            int result = comparator.compare(thisArea, otherArea);
            if (result > 0) {
                System.out.println(this.name + " is larger than " + other.name);
            } else if (result < 0) {
                System.out.println(this.name + " is smaller than " + other.name);
            } else {
                System.out.println(this.name + " and " + other.name + " have equal area");
            }
        }
        
        // Template method pattern
        public void analyze() {
            System.out.println("Analyzing " + name);
            validateShape();
            calculateProperties();
            displayInfo();
        }
        
        protected void validateShape() {
            System.out.println("Validating shape properties...");
        }
        
        protected void calculateProperties() {
            System.out.println("Calculating area and perimeter...");
        }
        
        // Functional method
        public void processWith(Function<Shape, String> processor) {
            String result = processor.apply(this);
            System.out.println("Processing result: " + result);
        }
        
        // Getters
        public String getName() { return name; }
        public String getColor() { return color; }
        
        // Static method
        public static void compareAll(List<Shape> shapes) {
            shapes.forEach(shape -> 
                System.out.println(shape.name + ": " + shape.getArea())
            );
        }
    }
    
    // Concrete class: Circle
    class Circle extends Shape {
        private double radius;
        
        public Circle(String color, double radius) {
            super("Circle", color);
            this.radius = radius;
        }
        
        @Override
        public double getArea() {
            return Math.PI * radius * radius;
        }
        
        @Override
        public double getPerimeter() {
            return 2 * Math.PI * radius;
        }
        
        @Override
        protected void validateShape() {
            if (radius <= 0) {
                throw new IllegalArgumentException("Radius must be positive");
            }
            System.out.println("Circle validated: radius = " + radius);
        }
        
        // Additional method
        public double getDiameter() {
            return 2 * radius;
        }
    }
    
    // Concrete class: Rectangle
    class Rectangle extends Shape {
        private double width;
        private double height;
        
        public Rectangle(String color, double width, double height) {
            super("Rectangle", color);
            this.width = width;
            this.height = height;
        }
        
        @Override
        public double getArea() {
            return width * height;
        }
        
        @Override
        public double getPerimeter() {
            return 2 * (width + height);
        }
        
        @Override
        protected void validateShape() {
            if (width <= 0 || height <= 0) {
                throw new IllegalArgumentException("Width and height must be positive");
            }
            System.out.println("Rectangle validated: width = " + width + ", height = " + height);
        }
        
        // Additional methods
        public double getWidth() { return width; }
        public double getHeight() { return height; }
        
        public boolean isSquare() {
            return width == height;
        }
    }
    
    // Concrete class: Triangle
    class Triangle extends Shape {
        private double side1, side2, side3;
        
        public Triangle(String color, double side1, double side2, double side3) {
            super("Triangle", color);
            this.side1 = side1;
            this.side2 = side2;
            this.side3 = side3;
        }
        
        @Override
        public double getArea() {
            double s = getPerimeter() / 2;
            return Math.sqrt(s * (s - side1) * (s - side2) * (s - side3));
        }
        
        @Override
        public double getPerimeter() {
            return side1 + side2 + side3;
        }
        
        @Override
        protected void validateShape() {
            if (side1 <= 0 || side2 <= 0 || side3 <= 0) {
                throw new IllegalArgumentException("All sides must be positive");
            }
            if (side1 + side2 <= side3 || side2 + side3 <= side1 || side3 + side1 <= side2) {
                throw new IllegalArgumentException("Invalid triangle sides");
            }
            System.out.println("Triangle validated: sides = " + side1 + ", " + side2 + ", " + side3);
        }
        
        // Additional method
        public String getTriangleType() {
            if (side1 == side2 && side2 == side3) {
                return "Equilateral";
            } else if (side1 == side2 || side2 == side3 || side3 == side1) {
                return "Isosceles";
            } else {
                return "Scalene";
            }
        }
    }
    
    // Functional interface for shape operations
    @FunctionalInterface
    interface ShapeOperation {
        void apply(Shape shape);
        
        default void describe(Shape shape) {
            System.out.print("Operation on " + shape.getName() + ": ");
            apply(shape);
        }
    }
    
    // Shape factory with functional programming
    static class ShapeFactory {
        public static Shape createCircle(String color, double radius) {
            return new Circle(color, radius);
        }
        
        public static Shape createRectangle(String color, double width, double height) {
            return new Rectangle(color, width, height);
        }
        
        public static Shape createTriangle(String color, double side1, double side2, double side3) {
            return new Triangle(color, side1, side2, side3);
        }
        
        // Functional factory method
        public static Shape create(String type, String color, double... dimensions) {
            return switch (type.toLowerCase()) {
                case "circle" -> new Circle(color, dimensions[0]);
                case "rectangle" -> new Rectangle(color, dimensions[0], dimensions[1]);
                case "triangle" -> new Triangle(color, dimensions[0], dimensions[1], dimensions[2]);
                default -> throw new IllegalArgumentException("Unknown shape type: " + type);
            };
        }
    }
    
    public static void main(String[] args) {
        AbstractClassJava8 demo = new AbstractClassJava8();
        
        // Create shapes using factory
        Shape circle = demo.new Circle("Red", 5.0);
        Shape rectangle = demo.new Rectangle("Blue", 4.0, 6.0);
        Shape triangle = demo.new Triangle("Green", 3.0, 4.0, 5.0);
        
        List<Shape> shapes = Arrays.asList(circle, rectangle, triangle);
        
        System.out.println("=== Abstract Class Demo ===");
        
        // Use abstract methods
        shapes.forEach(Shape::displayInfo);
        
        System.out.println("\n=== Template Method Pattern ===");
        shapes.forEach(Shape::analyze);
        
        System.out.println("\n=== Shape Comparison ===");
        circle.compareWith(rectangle, Double::compareTo);
        rectangle.compareWith(triangle, Double::compareTo);
        
        System.out.println("\n=== Functional Operations ===");
        
        // Process shapes with functions
        ShapeOperation areaCheck = shape -> {
            if (shape.getArea() > 20) {
                System.out.println("Large shape");
            } else {
                System.out.println("Small shape");
            }
        };
        
        shapes.forEach(areaCheck::describe);
        
        // Process with custom function
        circle.processWith(shape -> "Color: " + shape.getColor() + ", Area: " + shape.getArea());
        
        // Use static method
        System.out.println("\n=== Static Method Usage ===");
        Shape.compareAll(shapes);
        
        System.out.println("\n=== Stream Processing ===");
        
        // Sort shapes by area
        List<Shape> sortedByArea = shapes.stream()
            .sorted(Comparator.comparingDouble(Shape::getArea))
            .collect(Collectors.toList());
        
        System.out.println("Shapes sorted by area:");
        sortedByArea.forEach(shape -> 
            System.out.println(shape.getName() + ": " + shape.getArea()));
        
        // Group by color
        Map<String, List<Shape>> shapesByColor = shapes.stream()
            .collect(Collectors.groupingBy(Shape::getColor));
        
        System.out.println("\nShapes by color:");
        shapesByColor.forEach((color, shapeList) -> {
            System.out.println(color + ": " + shapeList.size() + " shapes");
        });
        
        // Calculate statistics
        DoubleSummaryStatistics areaStats = shapes.stream()
            .mapToDouble(Shape::getArea)
            .summaryStatistics();
        
        System.out.println("\nArea statistics:");
        System.out.println("Average: " + areaStats.getAverage());
        System.out.println("Max: " + areaStats.getMax());
        System.out.println("Min: " + areaStats.getMin());
        
        // Find largest shape
        Optional<Shape> largestShape = shapes.stream()
            .max(Comparator.comparingDouble(Shape::getArea));
        
        largestShape.ifPresent(shape -> 
            System.out.println("Largest shape: " + shape.getName()));
        
        System.out.println("\n=== Factory Pattern ===");
        
        // Create shapes using factory
        Shape factoryCircle = ShapeFactory.create("circle", "Yellow", 3.0);
        Shape factoryRectangle = ShapeFactory.create("rectangle", "Purple", 5.0, 7.0);
        
        factoryCircle.displayInfo();
        factoryRectangle.displayInfo();
        
        // Parallel processing
        System.out.println("\n=== Parallel Processing ===");
        shapes.parallelStream()
            .forEach(shape -> {
                System.out.println(Thread.currentThread().getName() + ": ");
                shape.displayInfo();
            });
        
        // Custom operations with specific shape types
        System.out.println("\n=== Specific Shape Operations ===");
        
        shapes.stream()
            .filter(shape -> shape instanceof Rectangle)
            .map(shape -> (Rectangle) shape)
            .forEach(rect -> {
                System.out.println("Rectangle " + rect.getName() + 
                    " is square: " + rect.isSquare());
            });
        
        shapes.stream()
            .filter(shape -> shape instanceof Circle)
            .map(shape -> (Circle) shape)
            .forEach(circle -> {
                System.out.println("Circle " + circle.getName() + 
                    " diameter: " + circle.getDiameter());
            });
    }
}
```

## 72. Custom Exception
**Java 8 Approach**: Using functional interfaces and Optional

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class CustomExceptionJava8 {
    
    // Custom checked exception with Java 8 features
    static class InvalidAgeException extends Exception {
        private final int age;
        private final String message;
        
        public InvalidAgeException(int age, String message) {
            super(message);
            this.age = age;
            this.message = message;
        }
        
        public int getAge() { return age; }
        public String getDetailedMessage() { return message; }
        
        @Override
        public String toString() {
            return String.format("InvalidAgeException{age=%d, message='%s'}", age, message);
        }
    }
    
    // Custom unchecked exception
    static class InvalidDataException extends RuntimeException {
        private final String field;
        private final Object value;
        
        public InvalidDataException(String field, Object value) {
            super(String.format("Invalid value '%s' for field '%s'", value, field));
            this.field = field;
            this.value = value;
        }
        
        public String getField() { return field; }
        public Object getValue() { return value; }
    }
    
    // Functional interface for validation
    @FunctionalInterface
    interface Validator<T> {
        void validate(T value) throws InvalidAgeException;
        
        default boolean isValid(T value) {
            try {
                validate(value);
                return true;
            } catch (InvalidAgeException e) {
                return false;
            }
        }
    }
    
    // Person class with validation
    static class Person {
        private final String name;
        private final int age;
        private final String email;
        
        public Person(String name, int age, String email) throws InvalidAgeException {
            this.name = Objects.requireNonNull(name, "Name cannot be null");
            this.age = age;
            this.email = email;
            
            validateAge();
            validateEmail();
        }
        
        private void validateAge() throws InvalidAgeException {
            if (age < 0) {
                throw new InvalidAgeException(age, "Age cannot be negative");
            }
            if (age > 150) {
                throw new InvalidAgeException(age, "Age seems unrealistic");
            }
        }
        
        private void validateEmail() {
            if (email != null && !email.contains("@")) {
                throw new InvalidDataException("email", email);
            }
        }
        
        // Factory method with Optional
        public static Optional<Person> createPerson(String name, int age, String email) {
            try {
                return Optional.of(new Person(name, age, email));
            } catch (InvalidAgeException e) {
                System.err.println("Failed to create person: " + e.getMessage());
                return Optional.empty();
            }
        }
        
        // Static validation method
        public static boolean isValidAge(int age) {
            return age >= 0 && age <= 150;
        }
        
        // Getters
        public String getName() { return name; }
        public int getAge() { return age; }
        public Optional<String> getEmail() { return Optional.ofNullable(email); }
        
        @Override
        public String toString() {
            return String.format("Person{name='%s', age=%d, email='%s'}", 
                               name, age, email);
        }
    }
    
    // Service class with exception handling
    static class PersonService {
        private final List<Person> people = new ArrayList<>();
        
        // Method with traditional exception handling
        public void addPerson(String name, int age, String email) throws InvalidAgeException {
            Person person = new Person(name, age, email);
            people.add(person);
            System.out.println("Added: " + person);
        }
        
        // Method with Optional return
        public Optional<Person> addPersonSafely(String name, int age, String email) {
            return Person.createPerson(name, age, email)
                .map(person -> {
                    people.add(person);
                    System.out.println("Added safely: " + person);
                    return person;
                });
        }
        
        // Method with functional validation
        public void addPersonWithValidator(String name, int age, String email, 
                                         Validator<Integer> ageValidator) throws InvalidAgeException {
            ageValidator.validate(age);
            Person person = new Person(name, age, email);
            people.add(person);
            System.out.println("Added with validator: " + person);
        }
        
        // Batch processing with exception handling
        public void addPeople(List<Map<String, Object>> peopleData) {
            Map<Boolean, List<Map<String, Object>>> results = peopleData.stream()
                .collect(Collectors.partitioningBy(data -> {
                    try {
                        int age = (Integer) data.getOrDefault("age", 0);
                        String name = (String) data.get("name");
                        String email = (String) data.get("email");
                        Person.createPerson(name, age, email);
                        return true;
                    } catch (Exception e) {
                        System.err.println("Invalid data: " + data + " - " + e.getMessage());
                        return false;
                    }
                }));
            
            List<Map<String, Object>> validData = results.get(true);
            validData.forEach(data -> {
                try {
                    String name = (String) data.get("name");
                    int age = (Integer) data.get("age");
                    String email = (String) data.get("email");
                    addPerson(name, age, email);
                } catch (InvalidAgeException e) {
                    System.err.println("Should not happen: " + e.getMessage());
                }
            });
            
            System.out.println("Successfully added " + validData.size() + " people");
            System.out.println("Failed to add " + results.get(false).size() + " people");
        }
        
        // Find people with exception handling
        public List<Person> findAdults() {
            return people.stream()
                .filter(person -> {
                    try {
                        return person.getAge() >= 18;
                    } catch (Exception e) {
                        System.err.println("Error checking age for " + person.getName());
                        return false;
                    }
                })
                .collect(Collectors.toList());
        }
        
        // Get people statistics with safe operations
        public Map<String, Object> getStatistics() {
            return people.stream()
                .collect(Collectors.collectingAndThen(
                    Collectors.toList(),
                    list -> {
                        Map<String, Object> stats = new HashMap<>();
                        stats.put("total", list.size());
                        stats.put("averageAge", list.stream()
                            .mapToInt(Person::getAge)
                            .average()
                            .orElse(0.0));
                        stats.put("adults", (int) list.stream()
                            .filter(p -> p.getAge() >= 18)
                            .count());
                        return stats;
                    }
                ));
        }
        
        public List<Person> getPeople() { return new ArrayList<>(people); }
    }
    
    public static void main(String[] args) {
        PersonService service = new PersonService();
        
        System.out.println("=== Custom Exception Demo ===");
        
        // Traditional exception handling
        try {
            service.addPerson("John", 25, "john@example.com");
            service.addPerson("Alice", 30, "alice@example.com");
        } catch (InvalidAgeException e) {
            System.err.println("Error: " + e.getDetailedMessage());
        }
        
        // Safe addition with Optional
        System.out.println("\n=== Safe Addition ===");
        service.addPersonSafely("Bob", -5, "bob@example.com")
            .ifPresentOrElse(
                person -> System.out.println("Successfully added: " + person),
                () -> System.out.println("Failed to add person")
            );
        
        service.addPersonSafely("Charlie", 35, "charlie@example.com")
            .ifPresent(person -> System.out.println("Successfully added: " + person));
        
        // Functional validation
        System.out.println("\n=== Functional Validation ===");
        Validator<Integer> adultValidator = age -> {
            if (age < 18) {
                throw new InvalidAgeException(age, "Must be an adult (18+)");
            }
        };
        
        try {
            service.addPersonWithValidator("David", 20, "david@example.com", adultValidator);
        } catch (InvalidAgeException e) {
            System.err.println("Validation failed: " + e.getDetailedMessage());
        }
        
        try {
            service.addPersonWithValidator("Eve", 16, "eve@example.com", adultValidator);
        } catch (InvalidAgeException e) {
            System.err.println("Validation failed: " + e.getDetailedMessage());
        }
        
        // Batch processing
        System.out.println("\n=== Batch Processing ===");
        List<Map<String, Object>> peopleData = Arrays.asList(
            Map.of("name", "Frank", "age", 40, "email", "frank@example.com"),
            Map.of("name", "Grace", "age", -10, "email", "grace@example.com"),
            Map.of("name", "Henry", "age", 28, "email", "invalid-email"),
            Map.of("name", "Ivy", "age", 22, "email", "ivy@example.com")
        );
        
        service.addPeople(peopleData);
        
        // Statistics
        System.out.println("\n=== Statistics ===");
        Map<String, Object> stats = service.getStatistics();
        stats.forEach((key, value) -> System.out.println(key + ": " + value));
        
        // Find adults
        System.out.println("\n=== Adults ===");
        List<Person> adults = service.findAdults();
        adults.forEach(person -> System.out.println(person.getName() + " (" + person.getAge() + ")"));
        
        // Exception chaining
        System.out.println("\n=== Exception Chaining ===");
        try {
            try {
                service.addPerson("Invalid", 200, "invalid@example.com");
            } catch (InvalidAgeException e) {
                throw new RuntimeException("Failed to process person", e);
            }
        } catch (RuntimeException e) {
            System.err.println("Caught runtime exception: " + e.getMessage());
            if (e.getCause() instanceof InvalidAgeException) {
                InvalidAgeException cause = (InvalidAgeException) e.getCause();
                System.err.println("Root cause: " + cause.getDetailedMessage());
            }
        }
        
        // Parallel processing with exception handling
        System.out.println("\n=== Parallel Processing ===");
        List<Integer> ages = Arrays.asList(25, -5, 30, 150, 35, -10);
        
        Map<Boolean, List<Integer>> validAges = ages.parallelStream()
            .collect(Collectors.partitioningBy(Person::isValidAge));
        
        System.out.println("Valid ages: " + validAges.get(true));
        System.out.println("Invalid ages: " + validAges.get(false));
        
        // Custom exception with lambda
        System.out.println("\n=== Lambda Exception Handling ===");
        List<String> emails = Arrays.asList("valid@example.com", "invalid-email", "another@example.com");
        
        emails.forEach(email -> {
            try {
                if (!email.contains("@")) {
                    throw new InvalidDataException("email", email);
                }
                System.out.println("Valid email: " + email);
            } catch (InvalidDataException e) {
                System.err.println("Invalid email: " + e.getValue());
            }
        });
        
        // Finally block demonstration
        System.out.println("\n=== Finally Block ===");
        try {
            System.out.println("Attempting operation...");
            service.addPerson("Final", 45, "final@example.com");
        } catch (InvalidAgeException e) {
            System.err.println("Operation failed: " + e.getMessage());
        } finally {
            System.out.println("Operation completed (finally block)");
        }
    }
}
```

## 73. Deep Copy vs Shallow Copy
**Java 8 Approach**: Using streams and functional interfaces

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class DeepCopyShallowCopyJava8 {
    
    // Original class with Java 8 features
    static class Person implements Cloneable {
        private String name;
        private int age;
        private Address address;
        private List<String> hobbies;
        
        public Person(String name, int age, Address address, List<String> hobbies) {
            this.name = name;
            this.age = age;
            this.address = address;
            this.hobbies = new ArrayList<>(hobbies);
        }
        
        // Shallow copy
        @Override
        public Person clone() {
            try {
                return (Person) super.clone();
            } catch (CloneNotSupportedException e) {
                throw new RuntimeException("Clone not supported", e);
            }
        }
        
        // Deep copy using constructor
        public Person deepCopy() {
            Address copiedAddress = new Address(address.getStreet(), address.getCity());
            List<String> copiedHobbies = new ArrayList<>(hobbies);
            return new Person(name, age, copiedAddress, copiedHobbies);
        }
        
        // Deep copy using copy constructor
        public Person(Person other) {
            this.name = other.name;
            this.age = other.age;
            this.address = new Address(other.address.getStreet(), other.address.getCity());
            this.hobbies = new ArrayList<>(other.hobbies);
        }
        
        // Functional deep copy
        public Person copyWith(Function<Person, Person> copier) {
            return copier.apply(this);
        }
        
        // Stream-based deep copy
        public static List<Person> deepCopyList(List<Person> people) {
            return people.stream()
                .map(Person::deepCopy)
                .collect(Collectors.toList());
        }
        
        // Modify methods to demonstrate copy differences
        public void modifyName(String newName) {
            this.name = newName;
        }
        
        public void modifyAge(int newAge) {
            this.age = newAge;
        }
        
        public void modifyAddress(String newStreet, String newCity) {
            this.address.setStreet(newStreet);
            this.address.setCity(newCity);
        }
        
        public void addHobby(String hobby) {
            this.hobbies.add(hobby);
        }
        
        // Getters
        public String getName() { return name; }
        public int getAge() { return age; }
        public Address getAddress() { return address; }
        public List<String> getHobbies() { return new ArrayList<>(hobbies); }
        
        @Override
        public String toString() {
            return String.format("Person{name='%s', age=%d, address=%s, hobbies=%s}", 
                               name, age, address, hobbies);
        }
        
        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Person person = (Person) o;
            return age == person.age && 
                   Objects.equals(name, person.name) && 
                   Objects.equals(address, person.address) && 
                   Objects.equals(hobbies, person.hobbies);
        }
    }
    
    // Address class (mutable)
    static class Address implements Cloneable {
        private String street;
        private String city;
        
        public Address(String street, String city) {
            this.street = street;
            this.city = city;
        }
        
        @Override
        public Address clone() {
            try {
                return (Address) super.clone();
            } catch (CloneNotSupportedException e) {
                throw new RuntimeException("Clone not supported", e);
            }
        }
        
        // Getters and setters
        public String getStreet() { return street; }
        public String getCity() { return city; }
        
        public void setStreet(String street) { this.street = street; }
        public void setCity(String city) { this.city = city; }
        
        @Override
        public String toString() {
            return String.format("Address{street='%s', city='%s'}", street, city);
        }
        
        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Address address = (Address) o;
            return Objects.equals(street, address.street) && 
                   Objects.equals(city, address.city);
        }
    }
    
    // Copy utility class with functional interfaces
    static class CopyUtils {
        
        // Functional interface for copying
        @FunctionalInterface
        interface Copier<T> {
            T copy(T original);
        }
        
        // Generic deep copy method
        public static <T> List<T> deepCopyList(List<T> original, Copier<T> copier) {
            return original.stream()
                .map(copier::copy)
                .collect(Collectors.toList());
        }
        
        // Copy with validation
        public static <T> Optional<T> safeCopy(T original, Copier<T> copier) {
            try {
                return Optional.of(copier.copy(original));
            } catch (Exception e) {
                System.err.println("Copy failed: " + e.getMessage());
                return Optional.empty();
            }
        }
        
        // Batch copy with statistics
        public static <T> CopyResult<T> batchCopy(List<T> originals, Copier<T> copier) {
            List<T> successful = new ArrayList<>();
            List<Map.Entry<T, Exception>> failed = new ArrayList<>();
            
            originals.forEach(original -> {
                try {
                    successful.add(copier.copy(original));
                } catch (Exception e) {
                    failed.add(new AbstractMap.SimpleEntry<>(original, e));
                }
            });
            
            return new CopyResult<>(successful, failed);
        }
        
        // Copy result container
        static class CopyResult<T> {
            private final List<T> successful;
            private final List<Map.Entry<T, Exception>> failed;
            
            public CopyResult(List<T> successful, List<Map.Entry<T, Exception>> failed) {
                this.successful = successful;
                this.failed = failed;
            }
            
            public List<T> getSuccessful() { return successful; }
            public List<Map.Entry<T, Exception>> getFailed() { return failed; }
            
            public boolean hasFailures() { return !failed.isEmpty(); }
            
            public double getSuccessRate() {
                return successful.size() * 100.0 / (successful.size() + failed.size());
            }
        }
    }
    
    // Person factory with copy capabilities
    static class PersonFactory {
        
        // Create person with builder pattern
        public static Person create(String name, int age, String street, String city, String... hobbies) {
            Address address = new Address(street, city);
            List<String> hobbyList = Arrays.asList(hobbies);
            return new Person(name, age, address, hobbyList);
        }
        
        // Create copy with modifications
        public static Person copyAndModify(Person original, Consumer<Person> modifier) {
            Person copy = original.deepCopy();
            modifier.accept(copy);
            return copy;
        }
        
        // Create multiple copies
        public static List<Person> createCopies(Person original, int count) {
            return IntStream.range(0, count)
                .mapToObj(i -> original.deepCopy())
                .collect(Collectors.toList());
        }
        
        // Create variations
        public static List<Person> createVariations(Person original, List<Consumer<Person>> modifiers) {
            return modifiers.stream()
                .map(modifier -> copyAndModify(original, modifier))
                .collect(Collectors.toList());
        }
    }
    
    public static void main(String[] args) {
        System.out.println("=== Deep Copy vs Shallow Copy Demo ===");
        
        // Create original person
        Person original = PersonFactory.create(
            "John Doe", 30, "123 Main St", "New York", "Reading", "Swimming"
        );
        
        System.out.println("Original: " + original);
        
        // Shallow copy
        Person shallowCopy = original.clone();
        System.out.println("Shallow copy: " + shallowCopy);
        
        // Deep copy
        Person deepCopy = original.deepCopy();
        System.out.println("Deep copy: " + deepCopy);
        
        // Modify shallow copy
        System.out.println("\n=== Modifying Shallow Copy ===");
        shallowCopy.modifyName("Jane Doe");
        shallowCopy.addHobby("Painting");
        shallowCopy.modifyAddress("456 Oak Ave", "Boston");
        
        System.out.println("After modifying shallow copy:");
        System.out.println("Original: " + original);
        System.out.println("Shallow copy: " + shallowCopy);
        System.out.println("Deep copy: " + deepCopy);
        
        // Modify deep copy
        System.out.println("\n=== Modifying Deep Copy ===");
        deepCopy.modifyAge(35);
        deepCopy.addHobby("Cooking");
        deepCopy.modifyAddress("789 Pine Rd", "Chicago");
        
        System.out.println("After modifying deep copy:");
        System.out.println("Original: " + original);
        System.out.println("Shallow copy: " + shallowCopy);
        System.out.println("Deep copy: " + deepCopy);
        
        // Functional copying
        System.out.println("\n=== Functional Copying ===");
        Person functionalCopy = original.copyWith(Person::deepCopy);
        functionalCopy.modifyName("Functional Person");
        System.out.println("Functional copy: " + functionalCopy);
        
        // Stream-based copying
        System.out.println("\n=== Stream-Based Copying ===");
        List<Person> people = Arrays.asList(
            PersonFactory.create("Alice", 25, "111 First St", "LA", "Music"),
            PersonFactory.create("Bob", 30, "222 Second St", "SF", "Sports"),
            PersonFactory.create("Charlie", 35, "333 Third St", "Seattle", "Travel")
        );
        
        List<Person> copiedPeople = Person.deepCopyList(people);
        System.out.println("Original people: " + people.size());
        System.out.println("Copied people: " + copiedPeople.size());
        
        // Modify copied list
        copiedPeople.get(0).modifyName("Modified Alice");
        System.out.println("After modification:");
        System.out.println("Original first person: " + people.get(0));
        System.out.println("Copied first person: " + copiedPeople.get(0));
        
        // Copy utilities
        System.out.println("\n=== Copy Utilities ===");
        
        CopyUtils.Copier<Person> copier = Person::deepCopy;
        List<Person> utilityCopied = CopyUtils.deepCopyList(people, copier);
        
        // Safe copy
        Optional<Person> safeCopy = CopyUtils.safeCopy(original, Person::deepCopy);
        safeCopy.ifPresent(copy -> System.out.println("Safe copy successful: " + copy));
        
        // Batch copy with statistics
        CopyUtils.CopyResult<Person> batchResult = CopyUtils.batchCopy(people, Person::deepCopy);
        System.out.println("Batch copy success rate: " + batchResult.getSuccessRate() + "%");
        
        // Factory patterns
        System.out.println("\n=== Factory Patterns ===");
        
        // Copy and modify
        Person modifiedCopy = PersonFactory.copyAndModify(original, person -> {
            person.modifyName("Modified Original");
            person.addHobby("New Hobby");
        });
        System.out.println("Modified copy: " + modifiedCopy);
        
        // Create multiple copies
        List<Person> multipleCopies = PersonFactory.createCopies(original, 3);
        System.out.println("Created " + multipleCopies.size() + " copies");
        
        // Create variations
        List<Consumer<Person>> modifiers = Arrays.asList(
            p -> p.modifyName("Variant 1"),
            p -> p.modifyAge(40),
            p -> p.addHobby("Variant Hobby")
        );
        
        List<Person> variations = PersonFactory.createVariations(original, modifiers);
        variations.forEach(variation -> System.out.println("Variation: " + variation));
        
        // Parallel copying
        System.out.println("\n=== Parallel Copying ===");
        List<Person> largeList = IntStream.range(0, 100)
            .mapToObj(i -> PersonFactory.create("Person" + i, 20 + i, 
                                             "Street" + i, "City" + i, "Hobby" + i))
            .collect(Collectors.toList());
        
        long startTime = System.currentTimeMillis();
        List<Person> parallelCopied = largeList.parallelStream()
            .map(Person::deepCopy)
            .collect(Collectors.toList());
        long parallelTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        List<Person> sequentialCopied = largeList.stream()
            .map(Person::deepCopy)
            .collect(Collectors.toList());
        long sequentialTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel copy time: " + parallelTime + "ms");
        System.out.println("Sequential copy time: " + sequentialTime + "ms");
        
        // Demonstrate immutability benefits
        System.out.println("\n=== Immutability Benefits ===");
        List<Person> immutableView = Collections.unmodifiableList(copiedPeople);
        try {
            immutableView.get(0).modifyName("Try to modify");
            System.out.println("Can modify person objects in unmodifiable list");
        } catch (Exception e) {
            System.out.println("Cannot modify: " + e.getMessage());
        }
        
        try {
            immutableView.add(new Person("New", 0, new Address("", ""), Collections.emptyList()));
            System.out.println("Should not reach here");
        } catch (Exception e) {
            System.out.println("Cannot modify unmodifiable list: " + e.getMessage());
        }
    }
}
```

## 74. Static Block vs Instance Block
**Java 8 Approach**: Using functional interfaces and streams

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class StaticInstanceBlockJava8 {
    
    // Class demonstrating static and instance blocks with Java 8 features
    static class BlockDemo {
        private static int staticCounter;
        private int instanceCounter;
        private static List<String> staticMessages = new ArrayList<>();
        private List<String> instanceMessages = new ArrayList<>();
        
        // Static block
        static {
            System.out.println("Static Block 1 - Initializing static resources");
            staticCounter = 100;
            staticMessages.add("Static initialization");
            
            // Using Java 8 features in static block
            IntStream.range(0, 5).forEach(i -> staticMessages.add("Static item " + i));
        }
        
        // Second static block
        static {
            System.out.println("Static Block 2 - Additional static initialization");
            staticCounter += 50;
            staticMessages.add("Second static block");
        }
        
        // Instance block
        {
            System.out.println("Instance Block 1 - Initializing instance resources");
            instanceCounter = staticCounter;
            instanceMessages.add("Instance initialization");
            
            // Using Java 8 features in instance block
            IntStream.range(0, 3).forEach(i -> instanceMessages.add("Instance item " + i));
        }
        
        // Second instance block
        {
            System.out.println("Instance Block 2 - Additional instance initialization");
            instanceCounter += 10;
            instanceMessages.add("Second instance block");
        }
        
        // Constructor
        public BlockDemo() {
            System.out.println("Constructor - Creating instance");
            instanceCounter++;
        }
        
        // Constructor with parameter
        public BlockDemo(String name) {
            this();
            System.out.println("Constructor with parameter: " + name);
        }
        
        // Static factory method
        public static BlockDemo create(String name) {
            System.out.println("Static factory method called");
            return new BlockDemo(name);
        }
        
        // Static method using streams
        public static void showStaticInfo() {
            System.out.println("=== Static Info ===");
            System.out.println("Static counter: " + staticCounter);
            System.out.println("Static messages: " + staticMessages.size());
            staticMessages.forEach(msg -> System.out.println("  " + msg));
        }
        
        // Instance method using streams
        public void showInstanceInfo() {
            System.out.println("=== Instance Info ===");
            System.out.println("Instance counter: " + instanceCounter);
            System.out.println("Instance messages: " + instanceMessages.size());
            instanceMessages.forEach(msg -> System.out.println("  " + msg));
        }
        
        // Getters
        public static int getStaticCounter() { return staticCounter; }
        public int getInstanceCounter() { return instanceCounter; }
        public static List<String> getStaticMessages() { return new ArrayList<>(staticMessages); }
        public List<String> getInstanceMessages() { return new ArrayList<>(instanceMessages); }
    }
    
    // Advanced class with functional initialization
    static class FunctionalBlockDemo {
        private static Map<String, Supplier<String>> staticSuppliers;
        private Map<String, Supplier<String>> instanceSuppliers;
        private static List<Consumer<String>> staticConsumers;
        private List<Consumer<String>> instanceConsumers;
        
        // Static block with functional programming
        static {
            System.out.println("Functional Static Block - Setting up functional interfaces");
            
            // Initialize static suppliers
            staticSuppliers = new HashMap<>();
            staticSuppliers.put("time", () -> new Date().toString());
            staticSuppliers.put("random", () -> String.valueOf(Math.random()));
            staticSuppliers.put("counter", () -> String.valueOf(staticSuppliers.size()));
            
            // Initialize static consumers
            staticConsumers = new ArrayList<>();
            staticConsumers.add(msg -> System.out.println("Static Consumer 1: " + msg));
            staticConsumers.add(msg -> System.out.println("Static Consumer 2: " + msg));
        }
        
        // Instance block with functional programming
        {
            System.out.println("Functional Instance Block - Setting up instance functional interfaces");
            
            // Initialize instance suppliers
            instanceSuppliers = new HashMap<>();
            instanceSuppliers.put("instance", () -> "Instance " + System.currentTimeMillis());
            instanceSuppliers.put("hash", () -> String.valueOf(this.hashCode()));
            
            // Initialize instance consumers
            instanceConsumers = new ArrayList<>();
            instanceConsumers.add(msg -> System.out.println("Instance Consumer 1: " + msg));
            instanceConsumers.add(msg -> System.out.println("Instance Consumer 2: " + msg));
        }
        
        // Constructor
        public FunctionalBlockDemo() {
            System.out.println("FunctionalBlockDemo constructor");
        }
        
        // Method to use static suppliers
        public static void demonstrateStaticSuppliers() {
            System.out.println("=== Static Suppliers ===");
            staticSuppliers.forEach((key, supplier) -> {
                String value = supplier.get();
                System.out.println(key + ": " + value);
                staticConsumers.forEach(consumer -> consumer.accept(key + " -> " + value));
            });
        }
        
        // Method to use instance suppliers
        public void demonstrateInstanceSuppliers() {
            System.out.println("=== Instance Suppliers ===");
            instanceSuppliers.forEach((key, supplier) -> {
                String value = supplier.get();
                System.out.println(key + ": " + value);
                instanceConsumers.forEach(consumer -> consumer.accept(key + " -> " + value));
            });
        }
    }
    
    // Configuration class with block initialization
    static class Configuration {
        private static Properties staticProperties;
        private Properties instanceProperties;
        private static Map<String, Object> staticConfig;
        private Map<String, Object> instanceConfig;
        
        // Static block for configuration loading
        static {
            System.out.println("Static Block - Loading static configuration");
            staticProperties = new Properties();
            staticConfig = new HashMap<>();
            
            // Simulate loading configuration
            staticProperties.setProperty("app.name", "Demo Application");
            staticProperties.setProperty("app.version", "1.0.0");
            staticProperties.setProperty("debug", "true");
            
            // Process configuration with streams
            staticProperties.stringPropertyNames()
                .stream()
                .forEach(key -> staticConfig.put(key, staticProperties.getProperty(key)));
            
            System.out.println("Static configuration loaded: " + staticConfig.size() + " properties");
        }
        
        // Instance block for instance-specific configuration
        {
            System.out.println("Instance Block - Loading instance configuration");
            instanceProperties = new Properties();
            instanceConfig = new HashMap<>();
            
            // Copy static properties to instance
            staticProperties.forEach((key, value) -> instanceProperties.setProperty(key.toString(), value.toString()));
            
            // Add instance-specific properties
            instanceProperties.setProperty("instance.id", String.valueOf(System.currentTimeMillis()));
            instanceProperties.setProperty("instance.hash", String.valueOf(this.hashCode()));
            
            // Process instance configuration
            instanceProperties.stringPropertyNames()
                .stream()
                .forEach(key -> instanceConfig.put(key, instanceProperties.getProperty(key)));
        }
        
        public Configuration() {
            System.out.println("Configuration constructor");
        }
        
        // Methods to access configuration
        public static String getStaticProperty(String key) {
            return staticProperties.getProperty(key);
        }
        
        public String getInstanceProperty(String key) {
            return instanceProperties.getProperty(key);
        }
        
        public static Map<String, Object> getStaticConfig() {
            return new HashMap<>(staticConfig);
        }
        
        public Map<String, Object> getInstanceConfig() {
            return new HashMap<>(instanceConfig);
        }
        
        // Demonstrate configuration usage
        public static void showStaticConfiguration() {
            System.out.println("=== Static Configuration ===");
            staticConfig.forEach((key, value) -> 
                System.out.println(key + " = " + value));
        }
        
        public void showInstanceConfiguration() {
            System.out.println("=== Instance Configuration ===");
            instanceConfig.forEach((key, value) -> 
                System.out.println(key + " = " + value));
        }
    }
    
    // Block execution order tracker
    static class ExecutionTracker {
        private static List<String> executionLog = new ArrayList<>();
        private List<String> instanceLog = new ArrayList<>();
        
        // Static block
        static {
            executionLog.add("Static Block 1");
            System.out.println("ExecutionTracker - Static Block 1");
        }
        
        static {
            executionLog.add("Static Block 2");
            System.out.println("ExecutionTracker - Static Block 2");
        }
        
        // Instance block
        {
            instanceLog.add("Instance Block 1");
            executionLog.add("Instance Block 1");
            System.out.println("ExecutionTracker - Instance Block 1");
        }
        
        {
            instanceLog.add("Instance Block 2");
            executionLog.add("Instance Block 2");
            System.out.println("ExecutionTracker - Instance Block 2");
        }
        
        public ExecutionTracker() {
            executionLog.add("Constructor");
            instanceLog.add("Constructor");
            System.out.println("ExecutionTracker - Constructor");
        }
        
        public static void showExecutionLog() {
            System.out.println("=== Execution Order ===");
            executionLog.forEach((index, step) -> 
                System.out.println((index + 1) + ". " + step));
        }
        
        public void showInstanceLog() {
            System.out.println("=== Instance Execution Order ===");
            instanceLog.forEach((index, step) -> 
                System.out.println((index + 1) + ". " + step));
        }
        
        public static List<String> getExecutionLog() {
            return new ArrayList<>(executionLog);
        }
    }
    
    public static void main(String[] args) {
        System.out.println("=== Static Block vs Instance Block Demo ===");
        
        // Demonstrate basic block execution
        System.out.println("\n--- Creating First Instance ---");
        BlockDemo demo1 = new BlockDemo("First");
        demo1.showInstanceInfo();
        
        System.out.println("\n--- Creating Second Instance ---");
        BlockDemo demo2 = new BlockDemo("Second");
        demo2.showInstanceInfo();
        
        // Show static info (should be the same for all instances)
        System.out.println("\n--- Static Information ---");
        BlockDemo.showStaticInfo();
        
        // Functional block demo
        System.out.println("\n=== Functional Block Demo ===");
        FunctionalBlockDemo functionalDemo = new FunctionalBlockDemo();
        FunctionalBlockDemo.demonstrateStaticSuppliers();
        functionalDemo.demonstrateInstanceSuppliers();
        
        // Configuration demo
        System.out.println("\n=== Configuration Demo ===");
        Configuration.showStaticConfiguration();
        Configuration config1 = new Configuration();
        config1.showInstanceConfiguration();
        Configuration config2 = new Configuration();
        config2.showInstanceConfiguration();
        
        // Execution order tracking
        System.out.println("\n=== Execution Order Tracking ===");
        ExecutionTracker.showExecutionLog();
        
        ExecutionTracker tracker1 = new ExecutionTracker();
        tracker1.showInstanceLog();
        
        ExecutionTracker tracker2 = new ExecutionTracker();
        tracker2.showInstanceLog();
        
        // Demonstrate with streams
        System.out.println("\n=== Stream Processing with Blocks ===");
        
        // Create multiple instances and analyze
        List<BlockDemo> demos = IntStream.range(0, 3)
            .mapToObj(i -> new BlockDemo("Demo" + i))
            .collect(Collectors.toList());
        
        // Analyze instance counters
        IntSummaryStatistics instanceStats = demos.stream()
            .mapToInt(BlockDemo::getInstanceCounter)
            .summaryStatistics();
        
        System.out.println("Instance counter statistics:");
        System.out.println("Min: " + instanceStats.getMin());
        System.out.println("Max: " + instanceStats.getMax());
        System.out.println("Average: " + instanceStats.getAverage());
        
        // Analyze static counter (should be same for all)
        System.out.println("Static counter: " + BlockDemo.getStaticCounter());
        
        // Demonstrate block execution with parallel streams
        System.out.println("\n=== Parallel Stream Demo ===");
        List<ExecutionTracker> trackers = IntStream.range(0, 5)
            .parallel()
            .mapToObj(i -> new ExecutionTracker())
            .collect(Collectors.toList());
        
        System.out.println("Created " + trackers.size() + " trackers in parallel");
        
        // Show final execution log
        System.out.println("\n--- Final Execution Log ---");
        ExecutionTracker.showExecutionLog();
        
        // Compare static vs instance initialization
        System.out.println("\n=== Static vs Instance Comparison ===");
        
        Map<String, Object> comparison = new HashMap<>();
        comparison.put("Static execution count", ExecutionTracker.getExecutionLog().stream()
            .filter(step -> step.startsWith("Static"))
            .count());
        comparison.put("Instance execution count", ExecutionTracker.getExecutionLog().stream()
            .filter(step -> step.startsWith("Instance"))
            .count());
        
        comparison.forEach((key, value) -> 
            System.out.println(key + ": " + value));
        
        // Demonstrate lazy initialization with blocks
        System.out.println("\n=== Lazy Initialization Demo ===");
        
        class LazyInit {
            private static volatile LazyInit instance;
            private static boolean staticBlockExecuted = false;
            private boolean instanceBlockExecuted = false;
            
            static {
                System.out.println("LazyInit static block");
                staticBlockExecuted = true;
            }
            
            {
                System.out.println("LazyInit instance block");
                instanceBlockExecuted = true;
            }
            
            public static LazyInit getInstance() {
                if (instance == null) {
                    synchronized (LazyInit.class) {
                        if (instance == null) {
                            instance = new LazyInit();
                        }
                    }
                }
                return instance;
            }
            
            public static boolean isStaticBlockExecuted() {
                return staticBlockExecuted;
            }
            
            public boolean isInstanceBlockExecuted() {
                return instanceBlockExecuted;
            }
        }
        
        System.out.println("Before getInstance() - Static block executed: " + LazyInit.isStaticBlockExecuted());
        LazyInit lazy1 = LazyInit.getInstance();
        System.out.println("After getInstance() - Instance block executed: " + lazy1.isInstanceBlockExecuted());
        LazyInit lazy2 = LazyInit.getInstance();
        System.out.println("Second getInstance() - Same instance: " + (lazy1 == lazy2));
    }
}
```

## 75. Comparator vs Comparable
**Java 8 Approach**: Using lambda expressions and method references

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class ComparatorComparableJava8 {
    
    // Person class implementing Comparable
    static class Person implements Comparable<Person> {
        private String name;
        private int age;
        private double salary;
        private String department;
        
        public Person(String name, int age, double salary, String department) {
            this.name = name;
            this.age = age;
            this.salary = salary;
            this.department = department;
        }
        
        // Natural ordering by name
        @Override
        public int compareTo(Person other) {
            return this.name.compareTo(other.name);
        }
        
        // Getters
        public String getName() { return name; }
        public int getAge() { return age; }
        public double getSalary() { return salary; }
        public String getDepartment() { return department; }
        
        // Setters for modification
        public void setSalary(double salary) { this.salary = salary; }
        
        @Override
        public String toString() {
            return String.format("Person{name='%s', age=%d, salary=%.2f, department='%s'}", 
                               name, age, salary, department);
        }
        
        // Static comparators
        public static Comparator<Person> byAge() {
            return Comparator.comparingInt(Person::getAge);
        }
        
        public static Comparator<Person> bySalary() {
            return Comparator.comparingDouble(Person::getSalary);
        }
        
        public static Comparator<Person> byDepartment() {
            return Comparator.comparing(Person::getDepartment);
        }
        
        // Multi-field comparators
        public static Comparator<Person> byAgeThenName() {
            return Comparator.comparingInt(Person::getAge)
                     .thenComparing(Person::getName);
        }
        
        public static Comparator<Person> byDepartmentThenSalary() {
            return Comparator.comparing(Person::getDepartment)
                     .thenComparingDouble(Person::getSalary);
        }
        
        // Custom comparator with lambda
        public static Comparator<Person> byNameLength() {
            return Comparator.comparingInt(person -> person.getName().length());
        }
        
        // Reversed comparators
        public static Comparator<Person> byAgeReversed() {
            return byAge().reversed();
        }
        
        public static Comparator<Person> bySalaryReversed() {
            return bySalary().reversed();
        }
        
        // Complex comparator
        public static Comparator<Person> byComplexCriteria() {
            return Comparator.comparing(Person::getDepartment)
                     .thenComparing(Person::getAge)
                     .thenComparingDouble(Person::getSalary).reversed()
                     .thenComparing(Person::getName);
        }
    }
    
    // Employee class extending Person with additional comparators
    static class Employee extends Person {
        private String employeeId;
        private int yearsOfService;
        
        public Employee(String name, int age, double salary, String department, 
                       String employeeId, int yearsOfService) {
            super(name, age, salary, department);
            this.employeeId = employeeId;
            this.yearsOfService = yearsOfService;
        }
        
        public String getEmployeeId() { return employeeId; }
        public int getYearsOfService() { return yearsOfService; }
        
        // Employee-specific comparators
        public static Comparator<Employee> byEmployeeId() {
            return Comparator.comparing(Employee::getEmployeeId);
        }
        
        public static Comparator<Employee> byYearsOfService() {
            return Comparator.comparingInt(Employee::getYearsOfService);
        }
        
        public static Comparator<Employee> byPerformance() {
            return Comparator.comparingDouble(Employee::getSalary)
                     .thenComparingInt(Employee::getYearsOfService);
        }
        
        @Override
        public String toString() {
            return String.format("Employee{id='%s', name='%s', age=%d, salary=%.2f, dept='%s', service=%d}", 
                               employeeId, getName(), getAge(), getSalary(), getDepartment(), yearsOfService);
        }
    }
    
    // Utility class for sorting operations
    static class SortingUtils {
        
        // Generic sorting method with comparator
        public static <T> List<T> sortBy(List<T> list, Comparator<T> comparator) {
            return list.stream()
                .sorted(comparator)
                .collect(Collectors.toList());
        }
        
        // Multi-criteria sorting
        public static <T> List<T> sortByMultiple(List<T> list, Comparator<T>... comparators) {
            Comparator<T> combined = Arrays.stream(comparators)
                .reduce(Comparator::thenComparing)
                .orElse(Comparator.naturalOrder());
            
            return list.stream()
                .sorted(combined)
                .collect(Collectors.toList());
        }
        
        // Sort and group
        public static <T, K> Map<K, List<T>> sortAndGroupBy(List<T> list, 
                                                           Function<T, K> classifier, 
                                                           Comparator<T> comparator) {
            return list.stream()
                .sorted(comparator)
                .collect(Collectors.groupingBy(classifier));
        }
        
        // Find extremes
        public static <T> Optional<T> findMax(List<T> list, Comparator<T> comparator) {
            return list.stream()
                .max(comparator);
        }
        
        public static <T> Optional<T> findMin(List<T> list, Comparator<T> comparator) {
            return list.stream()
                .min(comparator);
        }
        
        // Partition by comparator
        public static <T> Map<Boolean, List<T>> partitionBy(List<T> list, 
                                                           Predicate<T> predicate, 
                                                           Comparator<T> comparator) {
            return list.stream()
                .sorted(comparator)
                .collect(Collectors.partitioningBy(predicate));
        }
        
        // Custom sorting with lambda
        public static List<Person> sortByCustomCriteria(List<Person> people, 
                                                       Function<Person, String> criteria) {
            return people.stream()
                .sorted(Comparator.comparing(criteria))
                .collect(Collectors.toList());
        }
        
        // Parallel sorting
        public static <T> List<T> parallelSort(List<T> list, Comparator<T> comparator) {
            return list.parallelStream()
                .sorted(comparator)
                .collect(Collectors.toList());
        }
    }
    
    // Comparator factory
    static class ComparatorFactory {
        
        // Create comparator for Person
        public static Comparator<Person> createPersonComparator(String... fields) {
            return Arrays.stream(fields)
                .map(field -> switch (field.toLowerCase()) {
                    case "name" -> Comparator.comparing(Person::getName);
                    case "age" -> Comparator.comparingInt(Person::getAge);
                    case "salary" -> Comparator.comparingDouble(Person::getSalary);
                    case "department" -> Comparator.comparing(Person::getDepartment);
                    default -> throw new IllegalArgumentException("Unknown field: " + field);
                })
                .reduce(Comparator::thenComparing)
                .orElse(Comparator.naturalOrder());
        }
        
        // Create reversed comparator
        public static <T> Comparator<T> reversed(Comparator<T> comparator) {
            return comparator.reversed();
        }
        
        // Create null-safe comparator
        public static <T> Comparator<T> nullSafe(Comparator<T> comparator) {
            return Comparator.nullsFirst(comparator);
        }
        
        // Create compound comparator
        public static <T> Comparator<T> compound(List<Comparator<T>> comparators) {
            return comparators.stream()
                .reduce(Comparator::thenComparing)
                .orElse(Comparator.naturalOrder());
        }
    }
    
    public static void main(String[] args) {
        // Create test data
        List<Person> people = Arrays.asList(
            new Person("Alice", 30, 75000.0, "IT"),
            new Person("Bob", 25, 60000.0, "HR"),
            new Person("Charlie", 35, 85000.0, "IT"),
            new Person("Diana", 28, 70000.0, "Finance"),
            new Person("Eve", 32, 80000.0, "HR"),
            new Person("Frank", 40, 90000.0, "IT")
        );
        
        List<Employee> employees = Arrays.asList(
            new Employee("Alice", 30, 75000.0, "IT", "E001", 5),
            new Employee("Bob", 25, 60000.0, "HR", "E002", 3),
            new Employee("Charlie", 35, 85000.0, "IT", "E003", 8),
            new Employee("Diana", 28, 70000.0, "Finance", "E004", 4),
            new Employee("Eve", 32, 80000.0, "HR", "E005", 6)
        );
        
        System.out.println("=== Comparator vs Comparable Demo ===");
        
        // Natural ordering (Comparable)
        System.out.println("\n--- Natural Ordering (by name) ---");
        List<Person> sortedByName = people.stream()
            .sorted()
            .collect(Collectors.toList());
        
        sortedByName.forEach(person -> System.out.println(person.getName()));
        
        // Age comparison
        System.out.println("\n--- Sorted by Age ---");
        List<Person> sortedByAge = SortingUtils.sortBy(people, Person.byAge());
        sortedByAge.forEach(person -> 
            System.out.println(person.getName() + ": " + person.getAge()));
        
        // Salary comparison (reversed)
        System.out.println("\n--- Sorted by Salary (descending) ---");
        List<Person> sortedBySalaryDesc = SortingUtils.sortBy(people, Person.bySalaryReversed());
        sortedBySalaryDesc.forEach(person -> 
            System.out.println(person.getName() + ": $" + person.getSalary()));
        
        // Multi-criteria sorting
        System.out.println("\n--- Sorted by Department then Salary ---");
        List<Person> sortedByDeptSalary = SortingUtils.sortBy(people, Person.byDepartmentThenSalary());
        sortedByDeptSalary.forEach(person -> 
            System.out.println(person.getName() + ": " + person.getDepartment() + ", $" + person.getSalary()));
        
        // Complex criteria
        System.out.println("\n--- Sorted by Complex Criteria ---");
        List<Person> complexSorted = SortingUtils.sortBy(people, Person.byComplexCriteria());
        complexSorted.forEach(System.out::println);
        
        // Custom lambda comparator
        System.out.println("\n--- Sorted by Name Length ---");
        List<Person> sortedByNameLength = SortingUtils.sortBy(people, Person.byNameLength());
        sortedByNameLength.forEach(person -> 
            System.out.println(person.getName() + " (" + person.getName().length() + " chars)"));
        
        // Employee-specific sorting
        System.out.println("\n--- Employees sorted by Performance ---");
        List<Employee> sortedByPerformance = SortingUtils.sortBy(employees, Employee.byPerformance());
        sortedByPerformance.forEach(System.out::println);
        
        // Find extremes
        System.out.println("\n--- Finding Extremes ---");
        
        Optional<Person> oldest = SortingUtils.findMax(people, Person.byAge());
        Optional<Person> highestPaid = SortingUtils.findMax(people, Person.bySalary());
        Optional<Person> youngest = SortingUtils.findMin(people, Person.byAge());
        
        oldest.ifPresent(person -> System.out.println("Oldest: " + person.getName() + " (" + person.getAge() + ")"));
        highestPaid.ifPresent(person -> System.out.println("Highest paid: " + person.getName() + " ($" + person.getSalary() + ")"));
        youngest.ifPresent(person -> System.out.println("Youngest: " + person.getName() + " (" + person.getAge() + ")"));
        
        // Sort and group
        System.out.println("\n--- Sort and Group by Department ---");
        Map<String, List<Person>> groupedByDept = SortingUtils.sortAndGroupBy(
            people, Person::getDepartment, Person.bySalary()
        );
        
        groupedByDept.forEach((dept, deptPeople) -> {
            System.out.println(department + ":");
            deptPeople.forEach(person -> 
                System.out.println("  " + person.getName() + ": $" + person.getSalary()));
        });
        
        // Partition by age
        System.out.println("\n--- Partition by Age (30+) ---");
        Map<Boolean, List<Person>> partitionedByAge = SortingUtils.partitionBy(
            people, person -> person.getAge() >= 30, Person.byAge()
        );
        
        partitionedByAge.get(true).forEach(person -> 
            System.out.println("30+: " + person.getName() + " (" + person.getAge() + ")"));
        partitionedByAge.get(false).forEach(person -> 
            System.out.println("<30: " + person.getName() + " (" + person.getAge() + ")"));
        
        // Comparator factory
        System.out.println("\n--- Comparator Factory ---");
        Comparator<Person> multiFieldComparator = ComparatorFactory.createPersonComparator("department", "age", "salary");
        List<Person> factorySorted = SortingUtils.sortBy(people, multiFieldComparator);
        factorySorted.forEach(System.out::println);
        
        // Null-safe comparator
        System.out.println("\n--- Null-safe Comparator ---");
        List<Person> peopleWithNulls = new ArrayList<>(people);
        peopleWithNulls.add(null);
        
        List<Person> nullSafeSorted = peopleWithNulls.stream()
            .sorted(ComparatorFactory.nullSafe(Person.byAge()))
            .collect(Collectors.toList());
        
        nullSafeSorted.forEach(person -> {
            if (person == null) {
                System.out.println("null");
            } else {
                System.out.println(person.getName());
            }
        });
        
        // Parallel sorting
        System.out.println("\n--- Parallel Sorting ---");
        List<Person> largeList = IntStream.range(0, 1000)
            .mapToObj(i -> new Person("Person" + i, 20 + (i % 50), 50000 + (i % 100) * 1000, 
                                     Arrays.asList("IT", "HR", "Finance").get(i % 3)))
            .collect(Collectors.toList());
        
        long startTime = System.currentTimeMillis();
        List<Person> parallelSorted = SortingUtils.parallelSort(largeList, Person.bySalary());
        long parallelTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        List<Person> sequentialSorted = SortingUtils.sortBy(largeList, Person.bySalary());
        long sequentialTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel sort time: " + parallelTime + "ms");
        System.out.println("Sequential sort time: " + sequentialTime + "ms");
        
        // Method references
        System.out.println("\n--- Method References ---");
        List<String> names = people.stream()
            .sorted(Person.byAge())
            .map(Person::getName)
            .collect(Collectors.toList());
        
        System.out.println("Names sorted by age: " + names);
        
        // Custom sorting with lambda
        System.out.println("\n--- Custom Lambda Sorting ---");
        List<Person> customSorted = SortingUtils.sortByCustomCriteria(people, Person::getDepartment);
        customSorted.forEach(person -> 
            System.out.println(person.getName() + ": " + person.getDepartment()));
        
        // Chaining comparators
        System.out.println("\n--- Chaining Comparators ---");
        Comparator<Person> chainedComparator = Person.byDepartment()
            .thenComparing(Person::getAge)
            .thenComparing(Person::getSalary).reversed();
        
        List<Person> chainedSorted = SortingUtils.sortBy(people, chainedComparator);
        chainedSorted.forEach(System.out::println);
        
        // Statistics after sorting
        System.out.println("\n--- Statistics ---");
        Map<String, Long> deptCount = people.stream()
            .sorted(Person.byDepartment())
            .collect(Collectors.groupingBy(Person::getDepartment, Collectors.counting()));
        
        System.out.println("Department count: " + deptCount);
        
        DoubleSummaryStatistics salaryStats = people.stream()
            .sorted(Person.bySalary())
            .mapToDouble(Person::getSalary)
            .summaryStatistics();
        
        System.out.println("Salary statistics: " + salaryStats);
    }
}
```

---

## 🎯 Key Java 8 Benefits for OOPs Concepts

1. **Lambda Expressions**: Concise method implementations
2. **Functional Interfaces**: Enhanced abstraction capabilities
3. **Default Methods**: Interface evolution without breaking changes
4. **Method References**: Simplified method calls
5. **Streams API**: Powerful collection processing
6. **Optional**: Null-safe operations
7. **Parallel Processing**: Easy multi-threading

## 📝 Best Practices

1. **Use functional interfaces** for behavior parameterization
2. **Leverage default methods** for interface evolution
3. **Use Optional** for null-safe operations
4. **Prefer streams** for collection processing
5. **Use method references** for simple lambdas
6. **Consider parallel streams** for large datasets

---

*This collection demonstrates how Java 8 features enhance traditional OOPs concepts, making code more expressive, maintainable, and efficient.*
