# Proxy Pattern

## 🟢 What is it?
The **Proxy Pattern** lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.

Think of it like a **Credit Card**:
*   The Credit Card is a proxy for a **Bundle of Cash**.
*   It allows you to make payments.
*   It controls access (limits, pins).
*   Eventually, the money comes from the Bank Account (Real Subject), but the merchant interacts with the Card (Proxy).

---

## 🎯 Strategy to Implement

1.  **Service Interface**: Define a common interface for both the Real Service and the Proxy.
2.  **Real Service**: Create the class that does the actual work (heavy lifting).
3.  **Proxy Class**: Create a class implementing the same interface. It should have a field to store a reference to the Real Service.
4.  **Reference Loading**: The proxy usually creates or lazy-loads the Real Service only when needed.
5.  **Delegation**: The proxy implements the interface methods by logging/checking access and then calling the same method on the Real Service object.

---

## 💻 Code Example

```java
// 1. Service Interface
interface Image {
    void display();
}

// 2. Real Service (Heavy resource)
class RealImage implements Image {
    private String fileName;

    public RealImage(String fileName) {
        this.fileName = fileName;
        loadFromDisk(fileName); // Expensive operation
    }

    private void loadFromDisk(String fileName) {
        System.out.println("Loading " + fileName);
        // Simulating 5 seconds delay...
    }

    @Override
    public void display() {
        System.out.println("Displaying " + fileName);
    }
}

// 3. Proxy Class
class ProxyImage implements Image {
    private RealImage realImage;
    private String fileName;

    public ProxyImage(String fileName) {
        this.fileName = fileName;
    }

    @Override
    public void display() {
        // Lazy Loading: Create RealImage only when display() is called
        if (realImage == null) {
            realImage = new RealImage(fileName);
        }
        realImage.display();
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        Image image = new ProxyImage("test_10mb.jpg");

        // Image is NOT loaded from disk yet.
        System.out.println("Image object created but not loaded.");

        // Image will be loaded from disk now
        image.display(); 
        
        // Image is already loaded, just displayed
        image.display(); 
    }
}
```

---

## ✅ When to use?

*   **Virtual Proxy (Lazy Loading)**: When you have a heavy object (like a large image or database connection) that wastes system resources if it's always up, but you only need it occasionally.
*   **Protection Proxy (Access Control)**: When you want only specific clients to be able to use the service object (e.g., checking admin permissions before executing a command).
*   **Remote Proxy**: When the service object is located on a remote server. The proxy serves as a local representative.
*   **Logging Proxy**: When you want to keep a history of requests to the service object.
