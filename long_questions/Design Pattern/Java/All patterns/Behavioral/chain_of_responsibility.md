# Chain of Responsibility Pattern

## 🟢 What is it?
The **Chain of Responsibility Pattern** lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Think of it like **Calling Customer Support**:
1.  **Receptionist** picks up. "Can I solve it?" No -> Transfer.
2.  **Tech Support L1** picks up. "Can I solve it?" No -> Transfer.
3.  **Senior Engineer L2** picks up. "Can I solve it?" Yes -> Solved.
4.  If nobody handles it, the request is dropped or a default error is given.

---

## 🎯 Strategy to Implement

1.  **Handler Interface**: Declare a method for processing requests (e.g., `handle()`) and a method to set the `next` handler.
2.  **Base Handler**: Implement the `next` linkage logic. If the current handler can't handle it, call `next.handle()`.
3.  **Concrete Handlers**: Implement the logic. "If I can handle this request, do it. Else, call super.handle()".
4.  **Client**: Compose the chain (A -> B -> C) and pass the request to the first link (A).

---

## 💻 Code Example

```java
// 1. Base Handler
abstract class Logger {
    public static int INFO = 1;
    public static int ERROR = 2;

    protected int level;
    protected Logger nextLogger; // The next link

    public void setNextLogger(Logger nextLogger) {
        this.nextLogger = nextLogger;
    }

    public void logMessage(int level, String message) {
        // If this logger matches the level, print it
        if (this.level <= level) {
            write(message);
        }
        // Always pass up the chain regardless ?? 
        // Or in this implementation: pass to next regardless of whether I handled it or not
        if (nextLogger != null) {
            nextLogger.logMessage(level, message);
        }
    }

    abstract protected void write(String message);
}

// 2. Concrete Handlers
class ConsoleLogger extends Logger {
    public ConsoleLogger(int level) { this.level = level; }

    @Override
    protected void write(String message) {
        System.out.println("Console (Standard Config): " + message);
    }
}

class ErrorLogger extends Logger {
    public ErrorLogger(int level) { this.level = level; }

    @Override
    protected void write(String message) {
        System.err.println("Error Log (File System): " + message);
    }
}

class EmailLogger extends Logger {
    public EmailLogger(int level) { this.level = level; }

    @Override
    protected void write(String message) {
        System.out.println("Email Sent (Admin): " + message);
    }
}
```

### Usage:

```java
public class Main {
    private static Logger getChainOfLoggers() {
        // Chain: Email(3) -> File(2) -> Console(1)
        // High priority first, often. Or low priority first depending on logic.
        
        Logger errorLogger = new ErrorLogger(Logger.ERROR); // 2
        Logger consoleLogger = new ConsoleLogger(Logger.INFO); // 1

        errorLogger.setNextLogger(consoleLogger); // Error checks then passes to Console

        return errorLogger; 
    }

    public static void main(String[] args) {
        Logger loggerChain = getChainOfLoggers();

        System.out.println("--- Sending INFO ---");
        // ErrorLogger (2) ignores it (1 < 2), passes to Console (1). Console prints.
        loggerChain.logMessage(Logger.INFO, "This is an information.");

        System.out.println("\n--- Sending ERROR ---");
        // ErrorLogger (2) matches (2 <= 2), prints. Passes to Console (1). Console prints too.
        loggerChain.logMessage(Logger.ERROR, "System Failure!");
    }
}
```

---

## ✅ When to use?

*   **Multiple Handlers**: When more than one object may handle a request, and you don't want the sender to know explicitly who handles it.
*   **Decoupling**: When you want to decouple the sender of a request from its receivers.
*   **Dynamic Handling**: When the set of objects that can handle a request should be specified dynamically (e.g., adding a new Security Filter to a web server at runtime).
