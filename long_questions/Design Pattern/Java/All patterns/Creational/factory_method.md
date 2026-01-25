# Factory Method Pattern

## 🟢 What is it?
The **Factory Method Pattern** provides a way to create objects **without specifying the exact class** of object that will be created.

It defines an interface (or abstract class) for creating an object but lets **subclasses decide** which class to instantiate. It pushes the "instantiation logic" (the `new` keyword) from the client code to special "Factory" classes.

Think of it like a **Logistics Company**:
*   The generic process is "Deliver Package".
*   If the delivery is by land, the "Land Logistics" department creates a **Truck**.
*   If the delivery is by sea, the "Sea Logistics" department creates a **Ship**.
*   The client (you) just says "Deliver this", not "Go buy a Truck and drive it".

---

## 🎯 Strategy to Implement

1.  **Define a Product Interface**: This is the common interface for all objects your factory will create (e.g., `Transport`).
2.  **Create Concrete Products**: Implement the interface with specific classes (e.g., `Truck`, `Ship`).
3.  **Define the Creator (Factory) Class**: Declare the factory method (e.g., `createTransport()`) that returns a `Product`.
4.  **Create Concrete Creators**: Subclass the Creator to override the factory method and return a specific instance (e.g., `RoadLogistics` returns a `Truck`).

---

## 💻 Code Example

```java
// 1. The Product Interface
interface Transport {
    void deliver();
}

// 2. Concrete Products
class Truck implements Transport {
    public void deliver() {
        System.out.println("Delivering by land in a box.");
    }
}

class Ship implements Transport {
    public void deliver() {
        System.out.println("Delivering by sea in a container.");
    }
}

// 3. The Creator (Abstract Factory)
abstract class Logistics {
    // The Factory Method - subclasses must implement this
    abstract Transport createTransport();

    // The core business logic uses the factory method
    public void planDelivery() {
        Transport t = createTransport();
        t.deliver();
    }
}

// 4. Concrete Creators
class RoadLogistics extends Logistics {
    @Override
    Transport createTransport() {
        return new Truck();
    }
}

class SeaLogistics extends Logistics {
    @Override
    Transport createTransport() {
        return new Ship();
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        // Client works with the Creator abstract class
        Logistics logistics;

        // Based on configuration or user input...
        String type = "sea"; 

        if (type.equals("road")) {
            logistics = new RoadLogistics();
        } else {
            logistics = new SeaLogistics();
        }

        // The client code doesn't know if it's using a Truck or Ship.
        // It just works.
        logistics.planDelivery(); 
    }
}
```

---

## ✅ When to use?

*   **Unknown Dependencies**: When you don't know beforehand the exact types and dependencies of the objects your code should work with.
*   **Extensibility**: When you want to provide a library or framework that users can extend. They can create new "Product" types (e.g., `Drone`) and a new "Creator" (`AirLogistics`) without breaking your existing code.
*   **Decoupling**: When you want to save system resources by reusing existing objects instead of rebuilding them each time (managing object lifecycles).
