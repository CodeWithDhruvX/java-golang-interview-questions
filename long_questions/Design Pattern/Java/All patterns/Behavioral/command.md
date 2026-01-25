# Command Pattern

## 🟢 What is it?
The **Command Pattern** turns a request into a stand-alone object that contains all information about the request. This transformation lets you parameterize methods with different requests, delay or queue a request's execution, and support undoable operations.

Think of it like a **Waiter in a Restaurant**:
*   You (Client) give an **Order** (Command) to the **Waiter** (Invoker).
*   The Order contains details (Burger, no onions).
*   The Waiter doesn't cook. He takes the Order commands and sticks them on a queue.
*   The **Chef** (Receiver) sees the Order and executes it.
*   Because the order is an object, the waiter can undo it ("Cancel order"), log it ("Bill"), or queue it.

---

## 🎯 Strategy to Implement

1.  **Command Interface**: Declare the execution method (usually `execute()`).
2.  **Concrete Commands**: Implement requests. The command holds a reference to the Receiver (the object that actually does the work) and the parameters for the operation.
3.  **Receiver**: The class that contains the business logic.
4.  **Invoker**: The class that sends the request (e.g., a Button, a Remote Control, a Scheduler). It holds the Command and calls `execute()`.

---

## 💻 Code Example

```java
// 1. Receiver (Does the actual work)
class Light {
    public void turnOn() { System.out.println("Light is ON"); }
    public void turnOff() { System.out.println("Light is OFF"); }
}

// 2. Command Interface
interface Command {
    void execute();
    void undo();
}

// 3. Concrete Commands
class TurnOnLightCommand implements Command {
    private Light light;

    public TurnOnLightCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.turnOn();
    }
    
    @Override
    public void undo() {
        light.turnOff();
    }
}

class TurnOffLightCommand implements Command {
    private Light light;

    public TurnOffLightCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.turnOff();
    }
    
    @Override
    public void undo() {
        light.turnOn();
    }
}

// 4. Invoker (e.g., Remote Control / Smart Home Hub)
class RemoteControl {
    private Command command;

    public void setCommand(Command command) {
        this.command = command;
    }

    public void pressButton() {
        command.execute();
    }
    
    public void pressUndo() {
        command.undo();
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        Light livingRoomLight = new Light();
        
        Command lightsOn = new TurnOnLightCommand(livingRoomLight);
        Command lightsOff = new TurnOffLightCommand(livingRoomLight);

        RemoteControl remote = new RemoteControl();

        // Switch on
        remote.setCommand(lightsOn);
        remote.pressButton(); // Light is ON

        // Switch off
        remote.setCommand(lightsOff);
        remote.pressButton(); // Light is OFF
        
        // Undo last action (Turn it back on)
        remote.pressUndo(); // Light is ON
    }
}
```

---

## ✅ When to use?

*   **Undo/Redo**: When you want to implement reversible operations (Ctrl+Z). You store a history of executed Command objects.
*   **Queueing Tasks**: When you want to queue operations, schedule their execution, or execute them remotely (e.g., Job Queues, Thread Pools).
*   **Decoupling**: When you want to decouple the object that invokes the operation (Button) from the one that knows how to perform it (Business Logic).
