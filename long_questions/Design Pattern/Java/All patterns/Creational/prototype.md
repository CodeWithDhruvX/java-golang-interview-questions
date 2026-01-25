# Prototype Pattern

## 🟢 What is it?
The **Prototype Pattern** lets you copy existing objects without making your code dependent on their classes. It delegates the cloning process to the actual objects that are being cloned.

Think of it like **Cell Division (Mitosis)**:
*   A cell doesn't call a "Cell Factory" to make another cell.
*   It splits itself to create an exact copy.
*   Or like "Ctrl+C / Ctrl+V" in a document. You copy the existing template and then modify the copy.

---

## 🎯 Strategy to Implement

1.  **Prototype Interface**: Declare a `clone()` method in the interface (in Java, `Cloneable` interface is often used, but a custom interface is cleaner).
2.  **Concrete Prototype**: Implement the `clone()` method. The method should create a *new* object of the current class and carry over all field values.
3.  **Registry (Optional)**: Often, a "Registry" class keeps a cache of pre-made prototypes (e.g., "Big Enemy", "Small Enemy") ready to be cloned.

---

## 💻 Code Example

```java
import java.util.ArrayList;
import java.util.List;

// 1. Abstract Prototype
abstract class Shape implements Cloneable {
    public int x, y;
    public String color;

    public Shape() {}

    // Copy Constructor
    public Shape(Shape target) {
        if (target != null) {
            this.x = target.x;
            this.y = target.y;
            this.color = target.color;
        }
    }

    // Abstract clone method
    public abstract Shape clone();

    @Override
    public boolean equals(Object object2) {
        if (!(object2 instanceof Shape)) return false;
        Shape shape2 = (Shape) object2;
        return shape2.x == x && shape2.y == y && shape2.color.equals(color);
    }
}

// 2. Concrete Prototype
class Circle extends Shape {
    public int radius;

    public Circle() {}

    public Circle(Circle target) {
        super(target); // Copy parent fields
        if (target != null) {
            this.radius = target.radius;
        }
    }

    @Override
    public Shape clone() {
        return new Circle(this);
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        // Create original object (expensive setup usually)
        Circle circle1 = new Circle();
        circle1.x = 10;
        circle1.y = 20;
        circle1.radius = 15;
        circle1.color = "Red";

        // Cone it (fast)
        Circle circle2 = (Circle) circle1.clone();
        
        // Modify the clone
        circle2.color = "Blue"; // circle1 is still "Red"

        System.out.println("Circle 1: " + circle1.color); // Red
        System.out.println("Circle 2: " + circle2.color); // Blue
    }
}
```

---

## ✅ When to use?

*   **Expensive Creation**: When creating an object from scratch is resource-intensive definition (e.g., Database query results, parsing a large XML file). Cloning is often much faster.
*   **Complex Setup**: When you have objects with complex configurations (many fields set), and you need similar objects. Instead of re-configuring a new one, clone the template and tweak the differences.
*   **Subclassing Avoidance**: When you want to avoid a hierarchy of Factory classes that parallels the hierarchy of Product classes.
