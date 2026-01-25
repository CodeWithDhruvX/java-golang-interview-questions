# Singleton Pattern

## 🟢 What is it?
The **Singleton Pattern** ensures that a class has **only one instance** and provides a **global point of access** to it. It is one of the simplest design patterns but arguably one of the most controversial due to its global state nature.

Think of it like a **Government**: A country can have only one official government. No matter who you are (which part of the code you are in), if you ask for "The Government", you get the same single entity.

### Key Characteristics:
1.  **Single Instance**: The class is responsible for creating and holding its one and only instance.
2.  **Global Access**: The instance is globally accessible via a static method (usually `getInstance()`).
3.  **Private Constructor**: Prevents other objects from creating new instances using `new ClassName()`.

---

## 🎯 Strategy to Implement

1.  **Make the Constructor Private**: This stops anyone from creating an instance outside the class.
2.  **Create a Static Field**: This will hold the single instance of the class.
3.  **Create a Public Static Method**: Usually named `getInstance()`. This method checks if the instance exists:
    *   If it **doesn't exist**, it creates it, saves it in the static field, and returns it.
    *   If it **does exist**, it simply returns the stored instance.
    *   *Note: In multi-threaded environments, you need to add synchronization (locking) here to prevent two threads from creating an instance at the exact same millisecond.*

---

## 💻 Code Example

Here is a Thread-Safe, Lazy-Loaded Singleton implementation.

```java
public class DatabaseConnection {

    // 1. The private static variable to hold the single instance
    // 'volatile' ensures changes are visible to all threads immediately
    private static volatile DatabaseConnection instance;

    // 2. Private constructor prevents instantiation from other classes
    private DatabaseConnection() {
        System.out.println("Database Connection Created!");
        // Initialize connection logic here...
    }

    // 3. Public static method to get the instance
    public static DatabaseConnection getInstance() {
        // First check (no locking) - Performance optimization
        if (instance == null) {
            
            // Locking specifically for the creation block
            synchronized (DatabaseConnection.class) {
                
                // Second check (Double-Checked Locking)
                // Necessary because a second thread might have waited at the lock
                if (instance == null) {
                    instance = new DatabaseConnection();
                }
            }
        }
        return instance;
    }

    public void query(String sql) {
        System.out.println("Executing query: " + sql);
    }
}
```

### Usage:

```java
// Somewhere in your code...
DatabaseConnection db1 = DatabaseConnection.getInstance();
db1.query("SELECT * FROM users");

DatabaseConnection db2 = DatabaseConnection.getInstance();

// true, because they point to the exact same object in memory
System.out.println(db1 == db2); 
```

---

## ✅ When to use?

*   **Resource Management**: When you need to manage a shared resource like a **Database Connection Pool**, **Thread Pool**, or **File System**.
*   **Configuration & Logging**: When you need a single object to hold configuration settings (properties file) for the entire app, or a **Logger** that writes to a single log file.
*   **Hardware Interface**: When controlling access to a specific piece of hardware (e.g., a Printer Spooler), where multiple simultaneous commands could cause conflicts.

### ❌ When NOT to use?
*   Do not use it just to share "global variables" across your app. This creates hidden dependencies and makes unit testing very difficult (because you can't easily mock a Singleton).
