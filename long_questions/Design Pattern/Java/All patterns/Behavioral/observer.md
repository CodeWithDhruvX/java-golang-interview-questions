# Observer Pattern

## 🟢 What is it?
The **Observer Pattern** defines a subscription mechanism to notify multiple objects about any events that happen to the object they're observing.

Think of it like **YouTube Subscription**:
*   The **Channel** is the Subject (Publisher).
*   You and millions of others are **Subscribers** (Observers).
*   The Channel does not know who you are personally, just that you are on the list.
*   When a new video is uploaded (Event), the Channel notifies everyone on the list automatically.

---

## 🎯 Strategy to Implement

1.  **Observer Interface**: Declare the update interface for objects that act as observers (e.g., `update(msg)`).
2.  **Subject (Publisher) Interface**: Declare methods for attaching and detaching observers to a subject object.
3.  **Concrete Subject**: Store the state definition and the list of subscribers. When a state change happens, notify all observers.
4.  **Concrete Observers**: Implement the Observer interface to update their state in response to notifications.

---

## 💻 Code Example

```java
import java.util.ArrayList;
import java.util.List;

// 1. Observer Interface
interface Observer {
    void update(String news);
}

// 2. Subject Interface details
interface NewsAgency {
    void subscribe(Observer o);
    void unsubscribe(Observer o);
    void notifyObservers();
}

// 3. Concrete Subject
class CNN implements NewsAgency {
    private List<Observer> channels = new ArrayList<>();
    private String breakingNews;

    @Override
    public void subscribe(Observer o) {
        channels.add(o);
    }

    @Override
    public void unsubscribe(Observer o) {
        channels.remove(o);
    }

    @Override
    public void notifyObservers() {
        for (Observer o : channels) {
            o.update(breakingNews);
        }
    }

    public void setBreakingNews(String news) {
        this.breakingNews = news;
        System.out.println("CNN reports: " + news);
        notifyObservers(); // Trigger notification
    }
}

// 4. Concrete Observers
class MobileApp implements Observer {
    private String name;
    public MobileApp(String name) { this.name = name; }
    
    @Override
    public void update(String news) {
        System.out.println("App Notification on " + name + ": " + news);
    }
}

class EmailSubscriber implements Observer {
    private String email;
    public EmailSubscriber(String email) { this.email = email; }

    @Override
    public void update(String news) {
        System.out.println("Email sent to " + email + ": " + news);
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        CNN cnn = new CNN();

        Observer phone = new MobileApp("iPhone 15");
        Observer email = new EmailSubscriber("john@example.com");

        cnn.subscribe(phone);
        cnn.subscribe(email);

        // Event happens
        cnn.setBreakingNews("Aliens landed in New York!");
        
        // Output:
        // App Notification on iPhone 15: Aliens landed in New York!
        // Email sent to john@example.com: Aliens landed in New York!

        cnn.unsubscribe(email);

        // Next event
        cnn.setBreakingNews("It was just a movie set.");
        // Only phone gets notified now.
    }
}
```

---

## ✅ When to use?

*   **Event-Driven Systems**: When changes to the state of one object may require changing other objects, and the actual set of objects is unknown or changes dynamically.
*   **MVC Architecture**: When the View (Observer) needs to be updated automatically whenever the Model (Subject) changes its data.
*   **One-to-Many Dependencies**: When creates a loose coupling between the subject and observers. The subject only knows the observers implement an interface, not their concrete class.
