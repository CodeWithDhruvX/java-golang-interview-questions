# Template Method Pattern

## 🟢 What is it?
The **Template Method Pattern** defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

Think of it like **Building a House**:
*   The generic "Build House" process is standard: 
    1. Build Foundation
    2. Build Walls
    3. Build Roof
    4. Install Windows.
*   This is the **Template**.
*   A "Wooden House" subclass implements "Build Walls" with Wood.
*   A "Brick House" subclass implements "Build Walls" with Brick.
*   The sequence (Foundation -> Walls -> Roof) remains exactly the same for both.

---

## 🎯 Strategy to Implement

1.  **Abstract Class**: Declare the template method as `final` so it cannot be overridden. This methods calls other steps.
2.  **Abstract Steps**: Declare the steps that vary as `abstract methods`. Subclasses *must* implement these.
3.  **Hooks (Optional)**: Declare default empty methods that subclasses *can* override if they want to hook into the process at a specific point.
4.  **Concrete Classes**: Implement the abstract steps specific to their variation.

---

## 💻 Code Example

```java
// 1. Abstract Class (The Template)
abstract class DataMiner {
    
    // The Template Method - Defines the flow
    // Final prevents subclasses from changing the sequence
    public final void mineData() {
        connect();
        extractData();
        parseData();
        analyzeData();
        disconnect();
    }

    // Concrete steps (Common to all)
    private void connect() { System.out.println("Connecting to data source..."); }
    private void disconnect() { System.out.println("Disconnecting."); }

    // Abstract steps (Subclasses must implement)
    protected abstract void extractData();
    protected abstract void parseData();

    // Hook (Optional Override)
    protected void analyzeData() {
        System.out.println("Basic analysis performed.");
    }
}

// 2. Concrete Class 1: PDF Miner
class PDFMiner extends DataMiner {
    @Override
    protected void extractData() { System.out.println("Extracting text from PDF stream..."); }

    @Override
    protected void parseData() { System.out.println("Parsing PDF structure..."); }
}

// 3. Concrete Class 2: CSV Miner
class CSVMiner extends DataMiner {
    @Override
    protected void extractData() { System.out.println("Reading CSV file lines..."); }

    @Override
    protected void parseData() { System.out.println("Splitting CSV commas..."); }
    
    // Using the Hook
    @Override
    protected void analyzeData() {
        System.out.println("Running statistical analysis on CSV data.");
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        System.out.println("--- Processing PDF ---");
        DataMiner pdf = new PDFMiner();
        pdf.mineData();

        System.out.println("\n--- Processing CSV ---");
        DataMiner csv = new CSVMiner();
        csv.mineData();
    }
}
```

---

## ✅ When to use?

*   **Standard Process**: When you have an algorithm with a fixed sequence of steps, but some steps might vary.
*   **Prevent Duplication**: When you have slightly different implementations of the same process. You can pull the common code into the superclass (the template) and leave the unique bits in subclasses.
*   **Frameworks**: Very common in frameworks (e.g., Spring, React Lifecycles) where the framework calls your code (`onInit`, `doGet`) at specific times.
