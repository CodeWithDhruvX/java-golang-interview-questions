# Design Patterns: Practical Programs

**Goal**: Master essential design patterns with real-world implementations and use cases.

## 1. Singleton Pattern

### Thread-Safe Singleton Implementations

```java
// 1. Eager Initialization
class EagerSingleton {
    private static final EagerSingleton INSTANCE = new EagerSingleton();
    
    private EagerSingleton() {
        // Private constructor to prevent instantiation
        System.out.println("EagerSingleton created");
    }
    
    public static EagerSingleton getInstance() {
        return INSTANCE;
    }
    
    public void showMessage() {
        System.out.println("Eager Singleton: " + this.hashCode());
    }
}

// 2. Lazy Initialization (Not Thread-Safe)
class LazySingleton {
    private static LazySingleton instance;
    
    private LazySingleton() {
        System.out.println("LazySingleton created");
    }
    
    public static LazySingleton getInstance() {
        if (instance == null) {
            instance = new LazySingleton();
        }
        return instance;
    }
    
    public void showMessage() {
        System.out.println("Lazy Singleton: " + this.hashCode());
    }
}

// 3. Thread-Safe Lazy Initialization with Synchronization
class ThreadSafeLazySingleton {
    private static ThreadSafeLazySingleton instance;
    
    private ThreadSafeLazySingleton() {
        System.out.println("ThreadSafeLazySingleton created");
    }
    
    public static synchronized ThreadSafeLazySingleton getInstance() {
        if (instance == null) {
            instance = new ThreadSafeLazySingleton();
        }
        return instance;
    }
    
    public void showMessage() {
        System.out.println("Thread-Safe Lazy Singleton: " + this.hashCode());
    }
}

// 4. Double-Checked Locking
class DoubleCheckedSingleton {
    private static volatile DoubleCheckedSingleton instance;
    
    private DoubleCheckedSingleton() {
        System.out.println("DoubleCheckedSingleton created");
    }
    
    public static DoubleCheckedSingleton getInstance() {
        if (instance == null) {
            synchronized (DoubleCheckedSingleton.class) {
                if (instance == null) {
                    instance = new DoubleCheckedSingleton();
                }
            }
        }
        return instance;
    }
    
    public void showMessage() {
        System.out.println("Double-Checked Singleton: " + this.hashCode());
    }
}

// 5. Bill Pugh Singleton (Initialization-on-demand holder idiom)
class BillPughSingleton {
    private BillPughSingleton() {
        System.out.println("BillPughSingleton created");
    }
    
    private static class SingletonHelper {
        private static final BillPughSingleton INSTANCE = new BillPughSingleton();
    }
    
    public static BillPughSingleton getInstance() {
        return SingletonHelper.INSTANCE;
    }
    
    public void showMessage() {
        System.out.println("Bill Pugh Singleton: " + this.hashCode());
    }
}

// 6. Enum Singleton (Best approach)
enum EnumSingleton {
    INSTANCE;
    
    private final DatabaseConnection connection;
    
    EnumSingleton() {
        connection = new DatabaseConnection();
        System.out.println("EnumSingleton created");
    }
    
    public DatabaseConnection getConnection() {
        return connection;
    }
    
    public void showMessage() {
        System.out.println("Enum Singleton: " + this.hashCode());
    }
}

class DatabaseConnection {
    public void connect() {
        System.out.println("Connected to database");
    }
}

public class SingletonDemo {
    public static void main(String[] args) {
        System.out.println("=== Singleton Pattern Demo ===");
        
        // Test different singleton implementations
        System.out.println("\n--- Eager Singleton ---");
        EagerSingleton eager1 = EagerSingleton.getInstance();
        EagerSingleton eager2 = EagerSingleton.getInstance();
        eager1.showMessage();
        eager2.showMessage();
        System.out.println("Same instance? " + (eager1 == eager2));
        
        System.out.println("\n--- Lazy Singleton ---");
        LazySingleton lazy1 = LazySingleton.getInstance();
        LazySingleton lazy2 = LazySingleton.getInstance();
        lazy1.showMessage();
        lazy2.showMessage();
        System.out.println("Same instance? " + (lazy1 == lazy2));
        
        System.out.println("\n--- Thread-Safe Singleton ---");
        ThreadSafeLazySingleton tsLazy1 = ThreadSafeLazySingleton.getInstance();
        ThreadSafeLazySingleton tsLazy2 = ThreadSafeLazySingleton.getInstance();
        tsLazy1.showMessage();
        tsLazy2.showMessage();
        System.out.println("Same instance? " + (tsLazy1 == tsLazy2));
        
        System.out.println("\n--- Double-Checked Singleton ---");
        DoubleCheckedSingleton dc1 = DoubleCheckedSingleton.getInstance();
        DoubleCheckedSingleton dc2 = DoubleCheckedSingleton.getInstance();
        dc1.showMessage();
        dc2.showMessage();
        System.out.println("Same instance? " + (dc1 == dc2));
        
        System.out.println("\n--- Bill Pugh Singleton ---");
        BillPughSingleton bp1 = BillPughSingleton.getInstance();
        BillPughSingleton bp2 = BillPughSingleton.getInstance();
        bp1.showMessage();
        bp2.showMessage();
        System.out.println("Same instance? " + (bp1 == bp2));
        
        System.out.println("\n--- Enum Singleton ---");
        EnumSingleton enum1 = EnumSingleton.INSTANCE;
        EnumSingleton enum2 = EnumSingleton.INSTANCE;
        enum1.showMessage();
        enum2.showMessage();
        enum1.getConnection().connect();
        System.out.println("Same instance? " + (enum1 == enum2));
        
        // Demonstrate thread safety
        System.out.println("\n--- Thread Safety Test ---");
        testThreadSafety();
    }
    
    private static void testThreadSafety() {
        Runnable task = () -> {
            DoubleCheckedSingleton instance = DoubleCheckedSingleton.getInstance();
            System.out.println("Thread " + Thread.currentThread().getName() + 
                             " got instance: " + instance.hashCode());
        };
        
        Thread t1 = new Thread(task);
        Thread t2 = new Thread(task);
        Thread t3 = new Thread(task);
        
        t1.start();
        t2.start();
        t3.start();
        
        try {
            t1.join();
            t2.join();
            t3.join();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
}
```

## 2. Factory Pattern

### Vehicle Factory Implementation

```java
// Abstract Product
interface Vehicle {
    void start();
    void stop();
    void drive();
    String getType();
}

// Concrete Products
class Car implements Vehicle {
    private String type = "Car";
    
    @Override
    public void start() {
        System.out.println("Car engine started");
    }
    
    @Override
    public void stop() {
        System.out.println("Car engine stopped");
    }
    
    @Override
    public void drive() {
        System.out.println("Driving car smoothly");
    }
    
    @Override
    public String getType() {
        return type;
    }
}

class Motorcycle implements Vehicle {
    private String type = "Motorcycle";
    
    @Override
    public void start() {
        System.out.println("Motorcycle engine started");
    }
    
    @Override
    public void stop() {
        System.out.println("Motorcycle engine stopped");
    }
    
    @Override
    public void drive() {
        System.out.println("Riding motorcycle fast");
    }
    
    @Override
    public String getType() {
        return type;
    }
}

class Truck implements Vehicle {
    private String type = "Truck";
    
    @Override
    public void start() {
        System.out.println("Truck engine started");
    }
    
    @Override
    public void stop() {
        System.out.println("Truck engine stopped");
    }
    
    @Override
    public void drive() {
        System.out.println("Driving truck with heavy load");
    }
    
    @Override
    public String getType() {
        return type;
    }
}

// Simple Factory
class SimpleVehicleFactory {
    public static Vehicle createVehicle(String type) {
        switch (type.toLowerCase()) {
            case "car":
                return new Car();
            case "motorcycle":
                return new Motorcycle();
            case "truck":
                return new Truck();
            default:
                throw new IllegalArgumentException("Unknown vehicle type: " + type);
        }
    }
}

// Factory Method Pattern
abstract class VehicleFactory {
    public final Vehicle createVehicle() {
        Vehicle vehicle = createVehicleInstance();
        configureVehicle(vehicle);
        testVehicle(vehicle);
        return vehicle;
    }
    
    protected abstract Vehicle createVehicleInstance();
    
    protected void configureVehicle(Vehicle vehicle) {
        System.out.println("Configuring " + vehicle.getType());
    }
    
    protected void testVehicle(Vehicle vehicle) {
        System.out.println("Testing " + vehicle.getType());
    }
}

class CarFactory extends VehicleFactory {
    @Override
    protected Vehicle createVehicleInstance() {
        return new Car();
    }
    
    @Override
    protected void configureVehicle(Vehicle vehicle) {
        System.out.println("Installing car seats and airbags");
        super.configureVehicle(vehicle);
    }
}

class MotorcycleFactory extends VehicleFactory {
    @Override
    protected Vehicle createVehicleInstance() {
        return new Motorcycle();
    }
    
    @Override
    protected void configureVehicle(Vehicle vehicle) {
        System.out.println("Installing helmet holder");
        super.configureVehicle(vehicle);
    }
}

class TruckFactory extends VehicleFactory {
    @Override
    protected Vehicle createVehicleInstance() {
        return new Truck();
    }
    
    @Override
    protected void configureVehicle(Vehicle vehicle) {
        System.out.println("Installing cargo bed and towing equipment");
        super.configureVehicle(vehicle);
    }
}

// Abstract Factory Pattern
interface VehicleComponentFactory {
    Engine createEngine();
    Wheel createWheel();
    Seat createSeat();
}

class CarComponentFactory implements VehicleComponentFactory {
    @Override
    public Engine createEngine() {
        return new CarEngine();
    }
    
    @Override
    public Wheel createWheel() {
        return new CarWheel();
    }
    
    @Override
    public Seat createSeat() {
        return new CarSeat();
    }
}

class MotorcycleComponentFactory implements VehicleComponentFactory {
    @Override
    public Engine createEngine() {
        return new MotorcycleEngine();
    }
    
    @Override
    public Wheel createWheel() {
        return new MotorcycleWheel();
    }
    
    @Override
    public Seat createSeat() {
        return new MotorcycleSeat();
    }
}

// Component interfaces
interface Engine {
    void start();
    void stop();
}

interface Wheel {
    void rotate();
}

interface Seat {
    void adjust();
}

// Concrete components
class CarEngine implements Engine {
    @Override
    public void start() {
        System.out.println("Car engine started");
    }
    
    @Override
    public void stop() {
        System.out.println("Car engine stopped");
    }
}

class MotorcycleEngine implements Engine {
    @Override
    public void start() {
        System.out.println("Motorcycle engine started");
    }
    
    @Override
    public void stop() {
        System.out.println("Motorcycle engine stopped");
    }
}

class CarWheel implements Wheel {
    @Override
    public void rotate() {
        System.out.println("Car wheel rotating");
    }
}

class MotorcycleWheel implements Wheel {
    @Override
    public void rotate() {
        System.out.println("Motorcycle wheel rotating");
    }
}

class CarSeat implements Seat {
    @Override
    public void adjust() {
        System.out.println("Car seat adjusted");
    }
}

class MotorcycleSeat implements Seat {
    @Override
    public void adjust() {
        System.out.println("Motorcycle seat adjusted");
    }
}

// Advanced vehicle using abstract factory
class AdvancedVehicle implements Vehicle {
    private VehicleComponentFactory componentFactory;
    private Engine engine;
    private Wheel[] wheels;
    private Seat seat;
    private String type;
    
    public AdvancedVehicle(VehicleComponentFactory factory, String type) {
        this.componentFactory = factory;
        this.type = type;
        this.engine = componentFactory.createEngine();
        this.wheels = new Wheel[4]; // Simplified
        this.seat = componentFactory.createSeat();
        
        for (int i = 0; i < wheels.length; i++) {
            wheels[i] = componentFactory.createWheel();
        }
    }
    
    @Override
    public void start() {
        engine.start();
        System.out.println(type + " started");
    }
    
    @Override
    public void stop() {
        engine.stop();
        System.out.println(type + " stopped");
    }
    
    @Override
    public void drive() {
        for (Wheel wheel : wheels) {
            wheel.rotate();
        }
        System.out.println(type + " is driving");
    }
    
    @Override
    public String getType() {
        return type;
    }
    
    public void adjustSeat() {
        seat.adjust();
    }
}

public class FactoryPatternDemo {
    public static void main(String[] args) {
        System.out.println("=== Factory Pattern Demo ===");
        
        // Simple Factory
        System.out.println("\n--- Simple Factory ---");
        Vehicle car = SimpleVehicleFactory.createVehicle("car");
        Vehicle motorcycle = SimpleVehicleFactory.createVehicle("motorcycle");
        Vehicle truck = SimpleVehicleFactory.createVehicle("truck");
        
        car.start();
        car.drive();
        car.stop();
        
        motorcycle.start();
        motorcycle.drive();
        motorcycle.stop();
        
        truck.start();
        truck.drive();
        truck.stop();
        
        // Factory Method Pattern
        System.out.println("\n--- Factory Method Pattern ---");
        VehicleFactory carFactory = new CarFactory();
        VehicleFactory motorcycleFactory = new MotorcycleFactory();
        VehicleFactory truckFactory = new TruckFactory();
        
        Vehicle factoryCar = carFactory.createVehicle();
        Vehicle factoryMotorcycle = motorcycleFactory.createVehicle();
        Vehicle factoryTruck = truckFactory.createVehicle();
        
        factoryCar.start();
        factoryCar.drive();
        factoryCar.stop();
        
        // Abstract Factory Pattern
        System.out.println("\n--- Abstract Factory Pattern ---");
        VehicleComponentFactory carComponentFactory = new CarComponentFactory();
        VehicleComponentFactory motorcycleComponentFactory = new MotorcycleComponentFactory();
        
        AdvancedVehicle advancedCar = new AdvancedVehicle(carComponentFactory, "Advanced Car");
        AdvancedVehicle advancedMotorcycle = new AdvancedVehicle(motorcycleComponentFactory, "Advanced Motorcycle");
        
        advancedCar.start();
        advancedCar.adjustSeat();
        advancedCar.drive();
        advancedCar.stop();
        
        advancedMotorcycle.start();
        advancedMotorcycle.adjustSeat();
        advancedMotorcycle.drive();
        advancedMotorcycle.stop();
    }
}
```

## 3. Observer Pattern

### Weather Station Implementation

```java
import java.util.*;

// Observer interface
interface Observer {
    void update(float temperature, float humidity, float pressure);
    String getObserverName();
}

// Subject interface
interface Subject {
    void registerObserver(Observer observer);
    void removeObserver(Observer observer);
    void notifyObservers();
}

// Concrete Subject
class WeatherStation implements Subject {
    private List<Observer> observers;
    private float temperature;
    private float humidity;
    private float pressure;
    
    public WeatherStation() {
        this.observers = new ArrayList<>();
    }
    
    @Override
    public void registerObserver(Observer observer) {
        if (!observers.contains(observer)) {
            observers.add(observer);
            System.out.println(observer.getObserverName() + " registered");
        }
    }
    
    @Override
    public void removeObserver(Observer observer) {
        if (observers.remove(observer)) {
            System.out.println(observer.getObserverName() + " removed");
        }
    }
    
    @Override
    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update(temperature, humidity, pressure);
        }
    }
    
    public void setMeasurements(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
        System.out.println("\n--- Weather Data Updated ---");
        System.out.printf("Temp: %.1f°C, Humidity: %.1f%%, Pressure: %.1f hPa\n", 
                         temperature, humidity, pressure);
        notifyObservers();
    }
}

// Concrete Observers
class PhoneDisplay implements Observer {
    private String name;
    private float temperature;
    private float humidity;
    private float pressure;
    
    public PhoneDisplay(String name) {
        this.name = name;
    }
    
    @Override
    public void update(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
        display();
    }
    
    private void display() {
        System.out.println("📱 " + name + " Phone Display:");
        System.out.printf("   Temperature: %.1f°C\n", temperature);
        System.out.printf("   Humidity: %.1f%%\n", humidity);
        System.out.printf("   Pressure: %.1f hPa\n", pressure);
    }
    
    @Override
    public String getObserverName() {
        return name + " Phone";
    }
}

class TVDisplay implements Observer {
    private String name;
    private float temperature;
    private float humidity;
    private float pressure;
    
    public TVDisplay(String name) {
        this.name = name;
    }
    
    @Override
    public void update(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
        display();
    }
    
    private void display() {
        System.out.println("📺 " + name + " TV Display:");
        System.out.println("   Current Weather:");
        System.out.printf("   🌡️  %.1f°C | 💧 %.1f%% | 🌊 %.1f hPa\n", 
                         temperature, humidity, pressure);
    }
    
    @Override
    public String getObserverName() {
        return name + " TV";
    }
}

class WebApp implements Observer {
    private String name;
    private List<WeatherData> history;
    
    public WebApp(String name) {
        this.name = name;
        this.history = new ArrayList<>();
    }
    
    @Override
    public void update(float temperature, float humidity, float pressure) {
        WeatherData data = new WeatherData(temperature, humidity, pressure);
        history.add(data);
        display();
    }
    
    private void display() {
        System.out.println("🌐 " + name + " Web App:");
        System.out.println("   Weather data received and stored");
        System.out.println("   Total records: " + history.size());
    }
    
    public void showHistory() {
        System.out.println("\n--- " + name + " Weather History ---");
        for (int i = 0; i < history.size(); i++) {
            WeatherData data = history.get(i);
            System.out.printf("%d. %.1f°C, %.1f%%, %.1f hPa\n", 
                            i + 1, data.temperature, data.humidity, data.pressure);
        }
    }
    
    @Override
    public String getObserverName() {
        return name + " Web App";
    }
}

class WeatherData {
    float temperature;
    float humidity;
    float pressure;
    
    public WeatherData(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
    }
}

// Advanced Observer with custom events
interface WeatherEventListener {
    void onWeatherEvent(WeatherEvent event);
}

class WeatherEvent {
    private String eventType;
    private String description;
    private long timestamp;
    
    public WeatherEvent(String eventType, String description) {
        this.eventType = eventType;
        this.description = description;
        this.timestamp = System.currentTimeMillis();
    }
    
    // Getters
    public String getEventType() { return eventType; }
    public String getDescription() { return description; }
    public long getTimestamp() { return timestamp; }
}

class AdvancedWeatherStation extends WeatherStation {
    private List<WeatherEventListener> eventListeners;
    
    public AdvancedWeatherStation() {
        super();
        this.eventListeners = new ArrayList<>();
    }
    
    public void addEventListener(WeatherEventListener listener) {
        eventListeners.add(listener);
    }
    
    public void removeEventListener(WeatherEventListener listener) {
        eventListeners.remove(listener);
    }
    
    private void fireEvent(String eventType, String description) {
        WeatherEvent event = new WeatherEvent(eventType, description);
        for (WeatherEventListener listener : eventListeners) {
            listener.onWeatherEvent(event);
        }
    }
    
    @Override
    public void setMeasurements(float temperature, float humidity, float pressure) {
        float oldTemp = temperature;
        float oldHumidity = humidity;
        float oldPressure = pressure;
        
        super.setMeasurements(temperature, humidity, pressure);
        
        // Check for significant changes and fire events
        if (Math.abs(oldTemp - temperature) > 5.0) {
            fireEvent("TEMPERATURE_CHANGE", 
                     "Temperature changed significantly: " + oldTemp + "°C → " + temperature + "°C");
        }
        
        if (humidity > 80) {
            fireEvent("HIGH_HUMIDITY", "High humidity detected: " + humidity + "%");
        }
        
        if (pressure < 1000) {
            fireEvent("LOW_PRESSURE", "Low pressure detected: " + pressure + " hPa");
        }
    }
}

class WeatherAlertSystem implements WeatherEventListener {
    private String name;
    
    public WeatherAlertSystem(String name) {
        this.name = name;
    }
    
    @Override
    public void onWeatherEvent(WeatherEvent event) {
        System.out.println("🚨 " + name + " Alert System:");
        System.out.println("   Event: " + event.getEventType());
        System.out.println("   Description: " + event.getDescription());
        System.out.println("   Time: " + new java.util.Date(event.getTimestamp()));
    }
}

public class ObserverPatternDemo {
    public static void main(String[] args) {
        System.out.println("=== Observer Pattern Demo ===");
        
        // Create weather station
        WeatherStation weatherStation = new WeatherStation();
        
        // Create observers
        PhoneDisplay phone1 = new PhoneDisplay("Alice's");
        PhoneDisplay phone2 = new PhoneDisplay("Bob's");
        TVDisplay tv1 = new TVDisplay("Living Room");
        WebApp webApp = new WebApp("WeatherHub");
        
        // Register observers
        System.out.println("\n--- Registering Observers ---");
        weatherStation.registerObserver(phone1);
        weatherStation.registerObserver(phone2);
        weatherStation.registerObserver(tv1);
        weatherStation.registerObserver(webApp);
        
        // Update weather data
        System.out.println("\n--- First Weather Update ---");
        weatherStation.setMeasurements(25.5f, 65.0f, 1013.2f);
        
        System.out.println("\n--- Second Weather Update ---");
        weatherStation.setMeasurements(28.0f, 70.5f, 1010.8f);
        
        // Remove an observer
        System.out.println("\n--- Removing Observer ---");
        weatherStation.removeObserver(phone2);
        
        System.out.println("\n--- Third Weather Update ---");
        weatherStation.setMeasurements(22.3f, 85.2f, 998.5f);
        
        // Show web app history
        webApp.showHistory();
        
        // Advanced Observer with events
        System.out.println("\n=== Advanced Observer with Events ===");
        AdvancedWeatherStation advancedStation = new AdvancedWeatherStation();
        WeatherAlertSystem alertSystem = new WeatherAlertSystem("National Weather");
        
        advancedStation.registerObserver(new PhoneDisplay("Emergency Phone"));
        advancedStation.addEventListener(alertSystem);
        
        // Trigger events
        advancedStation.setMeasurements(20.0f, 60.0f, 1015.0f);
        advancedStation.setMeasurements(30.0f, 85.5f, 995.0f); // Should trigger events
    }
}
```

## 4. Builder Pattern

### Pizza Ordering System

```java
import java.util.*;

// Product class
class Pizza {
    private String size;
    private String dough;
    private String sauce;
    private List<String> toppings;
    private boolean cheese;
    private boolean pepperoni;
    private boolean mushrooms;
    private boolean olives;
    
    private Pizza(Builder builder) {
        this.size = builder.size;
        this.dough = builder.dough;
        this.sauce = builder.sauce;
        this.toppings = builder.toppings;
        this.cheese = builder.cheese;
        this.pepperoni = builder.pepperoni;
        this.mushrooms = builder.mushrooms;
        this.olives = builder.olives;
    }
    
    // Static nested Builder class
    public static class Builder {
        // Required parameters
        private final String size;
        private final String dough;
        
        // Optional parameters - initialized to default values
        private String sauce = "Tomato";
        private List<String> toppings = new ArrayList<>();
        private boolean cheese = true;
        private boolean pepperoni = false;
        private boolean mushrooms = false;
        private boolean olives = false;
        
        public Builder(String size, String dough) {
            this.size = size;
            this.dough = dough;
        }
        
        public Builder sauce(String sauce) {
            this.sauce = sauce;
            return this;
        }
        
        public Builder addTopping(String topping) {
            this.toppings.add(topping);
            return this;
        }
        
        public Builder cheese(boolean cheese) {
            this.cheese = cheese;
            return this;
        }
        
        public Builder pepperoni(boolean pepperoni) {
            this.pepperoni = pepperoni;
            return this;
        }
        
        public Builder mushrooms(boolean mushrooms) {
            this.mushrooms = mushrooms;
            return this;
        }
        
        public Builder olives(boolean olives) {
            this.olives = olives;
            return this;
        }
        
        public Pizza build() {
            Pizza pizza = new Pizza(this);
            validatePizza();
            return pizza;
        }
        
        private void validatePizza() {
            if (size == null || dough == null) {
                throw new IllegalStateException("Size and dough are required");
            }
        }
    }
    
    // Getters
    public String getSize() { return size; }
    public String getDough() { return dough; }
    public String getSauce() { return sauce; }
    public List<String> getToppings() { return new ArrayList<>(toppings); }
    public boolean hasCheese() { return cheese; }
    public boolean hasPepperoni() { return pepperoni; }
    public boolean hasMushrooms() { return mushrooms; }
    public boolean hasOlives() { return olives; }
    
    public void display() {
        System.out.println("=== Pizza Details ===");
        System.out.println("Size: " + size);
        System.out.println("Dough: " + dough);
        System.out.println("Sauce: " + sauce);
        System.out.println("Cheese: " + (cheese ? "Yes" : "No"));
        System.out.println("Pepperoni: " + (pepperoni ? "Yes" : "No"));
        System.out.println("Mushrooms: " + (mushrooms ? "Yes" : "No"));
        System.out.println("Olives: " + (olives ? "Yes" : "No"));
        
        if (!toppings.isEmpty()) {
            System.out.println("Additional Toppings:");
            toppings.forEach(topping -> System.out.println("  - " + topping));
        }
        System.out.println();
    }
    
    public double calculatePrice() {
        double basePrice = 0.0;
        
        // Base price by size
        switch (size.toLowerCase()) {
            case "small": basePrice = 8.99; break;
            case "medium": basePrice = 12.99; break;
            case "large": basePrice = 16.99; break;
            case "xlarge": basePrice = 20.99; break;
        }
        
        // Additional charges
        if (cheese) basePrice += 1.50;
        if (pepperoni) basePrice += 2.00;
        if (mushrooms) basePrice += 1.25;
        if (olives) basePrice += 1.00;
        
        basePrice += toppings.size() * 0.75;
        
        return basePrice;
    }
}

// Director class for common pizza configurations
class PizzaDirector {
    public Pizza createMargherita() {
        return new Pizza.Builder("Medium", "Thin Crust")
                .sauce("Tomato")
                .cheese(true)
                .addTopping("Fresh Basil")
                .addTopping("Extra Cheese")
                .build();
    }
    
    public Pizza createPepperoniLover() {
        return new Pizza.Builder("Large", "Thick Crust")
                .sauce("Tomato")
                .cheese(true)
                .pepperoni(true)
                .addTopping("Extra Pepperoni")
                .addTopping("Jalapenos")
                .build();
    }
    
    public Pizza createVeggieDelight() {
        return new Pizza.Builder("Medium", "Whole Wheat")
                .sauce("Pesto")
                .cheese(true)
                .mushrooms(true)
                .olives(true)
                .addTopping("Bell Peppers")
                .addTopping("Onions")
                .addTopping("Spinach")
                .build();
    }
    
    public Pizza createMeatFeast() {
        return new Pizza.Builder("XLarge", "Stuffed Crust")
                .sauce("BBQ")
                .cheese(true)
                .pepperoni(true)
                .addTopping("Sausage")
                .addTopping("Bacon")
                .addTopping("Ham")
                .build();
    }
}

// Complex object example: Computer Configuration
class Computer {
    private String cpu;
    private String motherboard;
    private int ram;
    private String storage;
    private String gpu;
    private String powerSupply;
    private String caseType;
    private List<String> additionalComponents;
    
    private Computer(ComputerBuilder builder) {
        this.cpu = builder.cpu;
        this.motherboard = builder.motherboard;
        this.ram = builder.ram;
        this.storage = builder.storage;
        this.gpu = builder.gpu;
        this.powerSupply = builder.powerSupply;
        this.caseType = builder.caseType;
        this.additionalComponents = builder.additionalComponents;
    }
    
    public static class ComputerBuilder {
        private String cpu;
        private String motherboard;
        private int ram = 8; // Default
        private String storage = "256GB SSD"; // Default
        private String gpu = "Integrated"; // Default
        private String powerSupply = "500W"; // Default
        private String caseType = "Mid Tower"; // Default
        private List<String> additionalComponents = new ArrayList<>();
        
        public ComputerBuilder(String cpu, String motherboard) {
            this.cpu = cpu;
            this.motherboard = motherboard;
        }
        
        public ComputerBuilder ram(int ram) {
            this.ram = ram;
            return this;
        }
        
        public ComputerBuilder storage(String storage) {
            this.storage = storage;
            return this;
        }
        
        public ComputerBuilder gpu(String gpu) {
            this.gpu = gpu;
            return this;
        }
        
        public ComputerBuilder powerSupply(String powerSupply) {
            this.powerSupply = powerSupply;
            return this;
        }
        
        public ComputerBuilder caseType(String caseType) {
            this.caseType = caseType;
            return this;
        }
        
        public ComputerBuilder addComponent(String component) {
            this.additionalComponents.add(component);
            return this;
        }
        
        public Computer build() {
            return new Computer(this);
        }
    }
    
    public void display() {
        System.out.println("=== Computer Configuration ===");
        System.out.println("CPU: " + cpu);
        System.out.println("Motherboard: " + motherboard);
        System.out.println("RAM: " + ram + "GB");
        System.out.println("Storage: " + storage);
        System.out.println("GPU: " + gpu);
        System.out.println("Power Supply: " + powerSupply);
        System.out.println("Case: " + caseType);
        
        if (!additionalComponents.isEmpty()) {
            System.out.println("Additional Components:");
            additionalComponents.forEach(comp -> System.out.println("  - " + comp));
        }
        System.out.println();
    }
}

public class BuilderPatternDemo {
    public static void main(String[] args) {
        System.out.println("=== Builder Pattern Demo ===");
        
        // Custom pizza using builder
        System.out.println("--- Custom Pizza ---");
        Pizza customPizza = new Pizza.Builder("Large", "Thin Crust")
                .sauce("Pesto")
                .cheese(true)
                .pepperoni(true)
                .mushrooms(true)
                .addTopping("Extra Cheese")
                .addTopping("Garlic")
                .addTopping("Red Pepper Flakes")
                .build();
        
        customPizza.display();
        System.out.printf("Price: $%.2f\n", customPizza.calculatePrice());
        
        // Predefined pizzas using director
        System.out.println("--- Predefined Pizzas ---");
        PizzaDirector director = new PizzaDirector();
        
        Pizza margherita = director.createMargherita();
        margherita.display();
        System.out.printf("Price: $%.2f\n", margherita.calculatePrice());
        
        Pizza pepperoniLover = director.createPepperoniLover();
        pepperoniLover.display();
        System.out.printf("Price: $%.2f\n", pepperoniLover.calculatePrice());
        
        Pizza veggieDelight = director.createVeggieDelight();
        veggieDelight.display();
        System.out.printf("Price: $%.2f\n", veggieDelight.calculatePrice());
        
        Pizza meatFeast = director.createMeatFeast();
        meatFeast.display();
        System.out.printf("Price: $%.2f\n", meatFeast.calculatePrice());
        
        // Computer configuration example
        System.out.println("--- Computer Configuration ---");
        Computer gamingPC = new Computer.ComputerBuilder("Intel i9-13900K", "ASUS ROG Strix")
                .ram(32)
                .storage("1TB NVMe SSD + 2TB HDD")
                .gpu("NVIDIA RTX 4090")
                .powerSupply("1000W Gold")
                .caseType("Full Tower")
                .addComponent("Liquid Cooling System")
                .addComponent("RGB Lighting")
                .addComponent("WiFi 6E Card")
                .build();
        
        gamingPC.display();
        
        Computer officePC = new Computer.ComputerBuilder("Intel i5-12400", "MSI B660")
                .ram(16)
                .storage("512GB SSD")
                .gpu("Integrated Intel UHD")
                .build();
        
        officePC.display();
        
        // Demonstrate builder pattern benefits
        System.out.println("--- Builder Pattern Benefits ---");
        System.out.println("1. Fluent interface for readable code");
        System.out.println("2. Immutable objects once built");
        System.out.println("3. Flexible construction with optional parameters");
        System.out.println("4. Validation at build time");
        System.out.println("5. Easy to create different configurations");
    }
}
```

## Practice Exercises

1. **Singleton**: Implement a logger class that can be used throughout the application
2. **Factory**: Create a payment processing system with different payment methods
3. **Observer**: Build a stock market monitoring system with multiple display types
4. **Builder**: Design a complex report generator with various formatting options

## Interview Questions

1. When would you use the Factory pattern over the Builder pattern?
2. What's the difference between the Observer pattern and the Publish-Subscribe pattern?
3. How does the Singleton pattern affect testability?
4. What are the potential issues with the Singleton pattern?
5. How can you make the Observer pattern more efficient for large numbers of observers?
