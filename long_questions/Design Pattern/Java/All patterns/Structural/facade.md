# Facade Pattern

## 🟢 What is it?
The **Facade Pattern** provides a simplified interface to a library, a framework, or any other complex set of classes.

Think of it like a **Home Theater Remote**:
*   To watch a movie, you have to:
    1.  Turn on the TV.
    2.  Set TV input to HDMI.
    3.  Turn on the Sound System.
    4.  Set Sound System to Movie Mode.
    5.  Turn on the Bluray Player.
    6.  Insert Disc.
    7.  Press Play.
*   A **Facade** (the "Watch Movie" button on a smart remote) does all of this for you with one click. It hides the complexity of the subsystems.

---

## 🎯 Strategy to Implement

1.  **Identify Subsystems**: Find the complex parts of your system that client code interacts with directly (e.g., DVDPlayer, Projector, Lights).
2.  **Create Facade Class**: Define a class with simple methods representing high-level actions (e.g., `watchMovie()`, `endMovie()`).
3.  **Delegate**: In the Facade methods, call the appropriate methods of the subsystems in the correct order.
4.  **Client Use**: The client code now calls the Facade methods instead of the subsystem methods.

---

## 💻 Code Example

```java
// Subsystem 1
class DVDPlayer {
    public void on() { System.out.println("DVD Player On"); }
    public void play(String movie) { System.out.println("Playing " + movie); }
}

// Subsystem 2
class Projector {
    public void on() { System.out.println("Projector On"); }
    public void setInput() { System.out.println("Projector Input set to DVD"); }
}

// Subsystem 3
class Lights {
    public void dim(int level) { System.out.println("Lights dimmed to " + level + "%"); }
}

// The Facade
class HomeTheaterFacade {
    private DVDPlayer dvd;
    private Projector projector;
    private Lights lights;

    public HomeTheaterFacade(DVDPlayer d, Projector p, Lights l) {
        this.dvd = d;
        this.projector = p;
        this.lights = l;
    }

    public void watchMovie(String movie) {
        System.out.println("Get ready to watch a movie...");
        lights.dim(10);
        projector.on();
        projector.setInput();
        dvd.on();
        dvd.play(movie);
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        // Without Facade (Complex)
        DVDPlayer dvd = new DVDPlayer();
        Projector projector = new Projector();
        Lights lights = new Lights();

        // complex setup...
        
        // With Facade (Simple)
        HomeTheaterFacade homeTheater = new HomeTheaterFacade(dvd, projector, lights);
        homeTheater.watchMovie("Inception");
    }
}
```

---

## ✅ When to use?

*   **Simplifying Complexity**: When you want to provide a simple interface to a complex subsystem.
*   **Layering**: When you want to layer your subsystems. Use a Facade to define an entry point to each subsystem level.
*   **Decoupling**: When you want to decouple the client implementation from the complex subsystem. If the subsystem changes (e.g., new type of Projector), you only modify the Facade, not the Client code.
