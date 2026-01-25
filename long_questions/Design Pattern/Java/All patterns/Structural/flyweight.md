# Flyweight Pattern

## 🟢 What is it?
The **Flyweight Pattern** is a structural pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

Think of it like a **Text Editor**:
*   A document has 100,000 characters.
*   Creating a new object for every single letter 'A', 'B', 'C' with all its font data, color, and size would crash your RAM.
*   Instead, you create **one** 'A' object (Flyweight) containing the shape of 'A'.
*   All 1000 instances of 'A' in the document just point to that shared object, storing only their specific position (x, y) coordinates externally.

---

## 🎯 Strategy to Implement

1.  **Intrinsic State (Shared)**: Identify data that is constant across many objects (e.g., Font, Color, Shape). Move this to the Flyweight class.
2.  **Extrinsic State (Unique)**: Identify data that changes per object (e.g., X/Y coordinates). Pass this to the Flyweight methods as parameters.
3.  **Flyweight Factory**: Create a factory that manages a pool of existing Flyweights. When requested, it returns an existing instance or creates a new one if it doesn't exist.

---

## 💻 Code Example

```java
import java.util.HashMap;
import java.util.Map;

// 1. Flyweight Interface
interface Tree {
    void draw(int x, int y);
}

// 2. Concrete Flyweight (Shared State)
// Stores type, color, texture (Heavy data)
class TreeType implements Tree {
    private String name;
    private String color;

    public TreeType(String name, String color) {
        this.name = name;
        this.color = color;
    }

    @Override
    public void draw(int x, int y) {
        // x and y are extrinsic (unique) state passed in
        System.out.println("Drawing " + name + " tree (" + color + ") at " + x + ", " + y);
    }
}

// 3. Flyweight Factory
class TreeFactory {
    private static Map<String, TreeType> treeTypes = new HashMap<>();

    public static TreeType getTreeType(String name, String color) {
        String key = name + "-" + color;
        if (!treeTypes.containsKey(key)) {
            treeTypes.put(key, new TreeType(name, color));
            System.out.println("Creating new TreeType: " + name);
        }
        return treeTypes.get(key);
    }
}
```

### Usage:

```java
public class Forest {
    public static void main(String[] args) {
        // We want to plant 10,000 trees
        // But we only create 2 actual objects in memory (Oak-Green, Pine-DarkGreen)
        
        for (int i = 0; i < 5000; i++) {
            TreeType type = TreeFactory.getTreeType("Oak", "Green");
            type.draw(i, i * 2); // Pass unique coordinates
        }

        for (int i = 0; i < 5000; i++) {
            TreeType type = TreeFactory.getTreeType("Pine", "DarkGreen");
            type.draw(i, i * 3);
        }
    }
}
```

---

## ✅ When to use?

*   **Massive Quantity**: When your application needs to spawn a huge number of similar objects (e.g., particles in a game, characters in a text editor).
*   **Memory Constraints**: When storing all data for every object drains available RAM.
*   **Shared State**: When the objects contain duplicate states that can be extracted and shared.
