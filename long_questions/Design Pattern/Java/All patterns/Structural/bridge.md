# Bridge Pattern

## 🟢 What is it?
The **Bridge Pattern** lets you split a large class or a set of closely related classes into two separate hierarchies—**Abstraction** and **Implementation**—which can be developed independently of each other.

Think of it like **Remote Controls and TVs**:
*   **Abstraction**: The Remote Control (Features: on, off, mute).
*   **Implementation**: The TV (Features: receiveSignal, powerInternal).
*   You can have a `BasicRemote` or `AdvancedRemote`.
*   You can have a `SonyTV` or `SamsungTV`.
*   The Remote (Abstraction) holds a reference to the TV (Implementation) and bridges the user command ("Mute") to the specific TV logic.

---

## 🎯 Strategy to Implement

1.  **Implementor Interface**: Define the interface for the implementation classes (e.g., `Device`).
2.  **Concrete Implementors**: Create concrete classes (e.g., `Radio`, `TV`) implementing the interface.
3.  **Abstraction Class**: Define the high-level control class (e.g., `Remote`) that holds a reference to an object of type Implementor.
4.  **Refined Abstraction**: Extend the abstraction to include more specific features (e.g., `AdvancedRemote`).

---

## 💻 Code Example

```java
// 1. Implementor (The device itself)
interface Device {
    void turnOn();
    void turnOff();
    void setChannel(int channel);
}

// 2. Concrete Implementors
class TV implements Device {
    public void turnOn() { System.out.println("TV: ON"); }
    public void turnOff() { System.out.println("TV: OFF"); }
    public void setChannel(int channel) { System.out.println("TV: Channel " + channel); }
}

class Radio implements Device {
    public void turnOn() { System.out.println("Radio: ON"); }
    public void turnOff() { System.out.println("Radio: OFF"); }
    public void setChannel(int channel) { System.out.println("Radio: Frequency " + channel); }
}

// 3. Abstraction (The Remote)
class RemoteControl {
    protected Device device; // The "Bridge"

    public RemoteControl(Device device) {
        this.device = device;
    }

    public void togglePower() {
        System.out.println("Remote: Power button pressed.");
        device.turnOn();
    }
}

// 4. Refined Abstraction
class AdvancedRemoteControl extends RemoteControl {
    public AdvancedRemoteControl(Device device) {
        super(device);
    }

    public void mute() {
        System.out.println("Remote: Mute button pressed.");
        device.setChannel(0);
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        Device tv = new TV();
        RemoteControl remote = new RemoteControl(tv);
        remote.togglePower(); // Works on TV

        Device radio = new Radio();
        AdvancedRemoteControl radioRemote = new AdvancedRemoteControl(radio);
        radioRemote.togglePower(); // Works on Radio
        radioRemote.mute(); // Works on Radio
    }
}
```

---

## ✅ When to use?

*   **Avoid Cartwright Explosion**: When you want to divide and organize a monolithic class that has several variants of some functionality (e.g., if you had `SonyRemote`, `SonyAdvancedRemote`, `SamsungRemote`... instead of just `Remote` + `SonyDevice`).
*   **Run-time Switching**: When you need to be able to switch implementations at runtime.
*   **Independent Extensibility**: When you want to extend a class hierarchy in two independent dimensions (Platform vs. Feature).
