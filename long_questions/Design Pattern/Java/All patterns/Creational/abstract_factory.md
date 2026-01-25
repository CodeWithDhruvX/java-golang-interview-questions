# Abstract Factory Pattern

## 🟢 What is it?
The **Abstract Factory Pattern** is a "Super Factory" that creates other factories. It provides an interface for creating **families of related or dependent objects** without specifying their concrete classes.

Think of it like **Furniture Shopping**:
*   You want to buy a **Sofa** and a **Chair**.
*   If you choose a "Victorian" style, you want a **Victorian Sofa** AND a **Victorian Chair**.
*   If you choose a "Modern" style, you want a **Modern Sofa** AND a **Modern Chair**.
*   You don't want to mix a Victorian Sofa with a Modern Chair. The Abstract Factory ensures the "family" stays consistent.

---

## 🎯 Strategy to Implement

1.  **Define Interfaces for Products**: Define interfaces for each distinct product type (e.g., `Chair`, `Sofa`).
2.  **Create Concrete Product Families**: Implement these interfaces for each variant (e.g., `VictorianChair`, `ModernChair`).
3.  **Define Abstract Factory Interface**: Declare methods to create each product type (e.g., `createChair()`, `createSofa()`).
4.  **Create Concrete Factories**: Implement the Abstract Factory for each family variant (e.g., `VictorianFurnitureFactory`, `ModernFurnitureFactory`).
5.  **Client Code**: The client works only with the Abstract Factory and Abstract Products, unaware of the specific variants.

---

## 💻 Code Example

```java
// 1. Abstract Products
interface Chair {
    void sitOn();
}
interface Sofa {
    void lieOn();
}

// 2. Concrete Product Family 1: Modern
class ModernChair implements Chair {
    public void sitOn() { System.out.println("Sitting on a sleek Modern Chair."); }
}
class ModernSofa implements Sofa {
    public void lieOn() { System.out.println("Lying on a minimalist Modern Sofa."); }
}

// 3. Concrete Product Family 2: Victorian
class VictorianChair implements Chair {
    public void sitOn() { System.out.println("Sitting on a fancy Victorian Chair."); }
}
class VictorianSofa implements Sofa {
    public void lieOn() { System.out.println("Lying on a velvet Victorian Sofa."); }
}

// 4. Abstract Factory Interface
interface FurnitureFactory {
    Chair createChair();
    Sofa createSofa();
}

// 5. Concrete Factories
class ModernFurnitureFactory implements FurnitureFactory {
    public Chair createChair() { return new ModernChair(); }
    public Sofa createSofa() { return new ModernSofa(); }
}

class VictorianFurnitureFactory implements FurnitureFactory {
    public Chair createChair() { return new VictorianChair(); }
    public Sofa createSofa() { return new VictorianSofa(); }
}
```

### Usage:

```java
public class Application {
    private Chair chair;
    private Sofa sofa;

    // Constructor gets the factory "injected"
    public Application(FurnitureFactory factory) {
        this.chair = factory.createChair();
        this.sofa = factory.createSofa();
    }

    public void paint() {
        chair.sitOn();
        sofa.lieOn();
    }

    public static void main(String[] args) {
        // App configured with Modern Factory
        Application app = new Application(new ModernFurnitureFactory());
        app.paint(); // Output: Modern Chair, Modern Sofa
        
        // App configured with Victorian Factory
        Application app2 = new Application(new VictorianFurnitureFactory());
        app2.paint(); // Output: Victorian Chair, Victorian Sofa
    }
}
```

---

## ✅ When to use?

*   **Product Families**: When your code needs to work with various families of related products (e.g., Windows vs. Mac UI buttons/scrollbars), and you want to ensure the client doesn't mix them (e.g., Windows Button on Mac Scrollbar).
*   **Encapsulation**: When you want to reveal only the interfaces of the products, not their implementations.
*   **Context Switching**: When you want to switch the entire "theme" or "platform" of your application just by changing one line of code (the Factory initialization).
