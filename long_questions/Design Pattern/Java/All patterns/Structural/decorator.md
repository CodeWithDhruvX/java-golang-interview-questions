# Decorator Pattern

## 🟢 What is it?
The **Decorator Pattern** lets you attach new behaviors to objects by placing these objects inside special wrapper objects that experience the behaviors.

Think of it like **Wearing Clothes**:
*   You (the Object) are a person.
*   You put on a T-shirt (Decorator 1). You are now a "Clothed Person".
*   You put on a Jacket (Decorator 2) over the T-shirt. You are now a "Warm Clothed Person".
*   If it rains, you put on a Raincoat (Decorator 3).
*   You can take them off or add layers in any order dynamically. You are still the same "Person" underneath.

---

## 🎯 Strategy to Implement

1.  **Component Interface**: Define the common interface for both the object and optional layers (e.g., `Coffee`).
2.  **Concrete Component**: Create the base object class (e.g., `SimpleCoffee`).
3.  **Base Decorator**: Create a class that implements the Component Interface and has a field for a Component object. It delegates all calls to that object.
4.  **Concrete Decorators**: Extend the Base Decorator. In their methods, they call the parent method (super) and then execute their extra behavior before or after.

---

## 💻 Code Example

```java
// 1. Component Interface
interface Coffee {
    String getDescription();
    double getCost();
}

// 2. Concrete Component
class SimpleCoffee implements Coffee {
    @Override
    public String getDescription() { return "Simple Coffee"; }
    
    @Override
    public double getCost() { return 5.0; }
}

// 3. Base Decorator (Abstract)
abstract class CoffeeDecorator implements Coffee {
    protected Coffee decoratedCoffee; // Reference to the object we are wrapping

    public CoffeeDecorator(Coffee c) {
        this.decoratedCoffee = c;
    }

    public String getDescription() { return decoratedCoffee.getDescription(); }
    public double getCost() { return decoratedCoffee.getCost(); }
}

// 4. Concrete Decorators
class Milk extends CoffeeDecorator {
    public Milk(Coffee c) { super(c); }

    @Override
    public String getDescription() {
        return decoratedCoffee.getDescription() + ", Milk"; 
    }

    @Override
    public double getCost() {
        return decoratedCoffee.getCost() + 1.5; // Add cost of milk
    }
}

class Sugar extends CoffeeDecorator {
    public Sugar(Coffee c) { super(c); }

    @Override
    public String getDescription() {
        return decoratedCoffee.getDescription() + ", Sugar"; 
    }

    @Override
    public double getCost() {
        return decoratedCoffee.getCost() + 0.5; // Add cost of sugar
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        // Order: Simple Coffee
        Coffee myCoffee = new SimpleCoffee();
        System.out.println(myCoffee.getDescription() + " $" + myCoffee.getCost());

        // Add Milk
        myCoffee = new Milk(myCoffee);
        System.out.println(myCoffee.getDescription() + " $" + myCoffee.getCost());

        // Add Sugar (Wrapping the Milk-wrapped Coffee)
        myCoffee = new Sugar(myCoffee);
        System.out.println(myCoffee.getDescription() + " $" + myCoffee.getCost());
        
        // Output: Simple Coffee, Milk, Sugar $7.0
    }
}
```

---

## ✅ When to use?

*   **Dynamic Extension**: When you need to add responsibilities to individual objects dynamically and transparently, without affecting other objects.
*   **Prevent Explosion**: When extending methods by inheritance would produce an explosion of subclasses (e.g., `MilkCoffee`, `SugarCoffee`, `MilkSugarCoffee`, `MilkSugarCaramelCoffee`... impossible!).
*   **Flexible Config**: When you want to structure code so that functionality can be added or removed at runtime (like toggling UI scrollbars).
