# Mediator Pattern

## 🟢 What is it?
The **Mediator Pattern** defines an object that encapsulates how a set of objects interact. It promotes loose coupling by keeping objects from referring to each other explicitly, allowing you to vary their interaction independently.

Think of it like an **Air Traffic Controller (ATC)**:
*   Pilots (Planes) do not talk to each other directly ("Hey Plane B, I am landing").
*   They talk to the Tower (Mediator).
*   The Tower decides who lands, who takes off, and who waits.
*   If we didn't have the Tower, every plane would need to know the position of every other plane (Chaos).

---

## 🎯 Strategy to Implement

1.  **Mediator Interface**: Declare methods for communication with components (e.g., `notify(sender, event)`).
2.  **Concrete Mediator**: Store references to all components. Implement the logic to coordinate them.
3.  **Components**: Each component stores a reference to the Mediator. Instead of calling other components directly, they call `mediator.notify(this, "event")`.

---

## 💻 Code Example

```java
// 1. Mediator Interface
interface ChatMediator {
    void sendMessage(String msg, User user);
    void addUser(User user);
}

// 2. Concrete Mediator
import java.util.ArrayList;
import java.util.List;

class ChatRoomImpl implements ChatMediator {
    private List<User> users = new ArrayList<>();

    @Override
    public void addUser(User user) {
        this.users.add(user);
    }

    @Override
    public void sendMessage(String msg, User user) {
        for (User u : this.users) {
            // Message should not be received by the user sending it
            if (u != user) {
                u.receive(msg);
            }
        }
    }
}

// 3. Components
abstract class User {
    protected ChatMediator mediator;
    protected String name;

    public User(ChatMediator med, String name) {
        this.mediator = med;
        this.name = name;
    }

    public abstract void send(String msg);
    public abstract void receive(String msg);
}

class ConcreteUser extends User {
    public ConcreteUser(ChatMediator med, String name) {
        super(med, name);
    }

    @Override
    public void send(String msg) {
        System.out.println(this.name + " Sending Message: " + msg);
        mediator.sendMessage(msg, this);
    }

    @Override
    public void receive(String msg) {
        System.out.println(this.name + " Received Message: " + msg);
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        ChatMediator chatRoom = new ChatRoomImpl();

        User user1 = new ConcreteUser(chatRoom, "Pankaj");
        User user2 = new ConcreteUser(chatRoom, "Lisa");
        User user3 = new ConcreteUser(chatRoom, "Saurabh");
        User user4 = new ConcreteUser(chatRoom, "David");

        chatRoom.addUser(user1);
        chatRoom.addUser(user2);
        chatRoom.addUser(user3);
        chatRoom.addUser(user4);

        user1.send("Hello World!");
        
        // Output:
        // Pankaj Sending Message: Hello World!
        // Lisa Received Message: Hello World!
        // Saurabh Received Message: Hello World!
        // David Received Message: Hello World!
    }
}
```

---

## ✅ When to use?

*   **Complex Communication**: When interaction between a large number of objects is complex and unstructured (Many-to-Many turned into One-to-Many).
*   **Coupling**: When you can't reuse an object in a different program because it depends on too many other objects.
*   **Centralized Control**: When you want to centralize complex control logic in one place rather than scattering it across components.
