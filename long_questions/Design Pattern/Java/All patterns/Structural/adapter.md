# Adapter Pattern

## 🟢 What is it?
The **Adapter Pattern** allows objects with incompatible interfaces to collaborate. It acts as a bridge between two incompatible interfaces.

Think of it like a **Travel Adapter**:
*   You have a US laptop plug (Client).
*   You are in Europe with a European Wall Socket (Service).
*   You cannot plug the US plug into the EU socket directly.
*   You use an **Adapter**. It takes the US plug on one side and fits into the EU socket on the other.

---

## 🎯 Strategy to Implement

1.  **Consumer Interface**: Start with the interface that your Client code expects to use (e.g., `USPlug`).
2.  **Service Class**: Identify the useful class that has an incompatible interface (e.g., `EuropeanSocket`).
3.  **Adapter Class**: Create a new class that implements the Consumer Interface.
4.  **Wrap the Service**: Inside the Adapter class, store a reference to the Service class (Composition).
5.  **Translate Calls**: In the methods of the Adapter, call the methods of the Service object, translating the data if necessary (e.g., converting 110v to 220v).

---

## 💻 Code Example

```java
// 1. Target Interface (What the client expects)
interface LightningPhone {
    void recharge();
    void useLightning();
}

// 2. Adaptee (The incompatible service we want to use)
interface MicroUsbPhone {
    void recharge();
    void useMicroUsb();
}

class AndroidPhone implements MicroUsbPhone {
    public void recharge() { System.out.println("Recharging Android..."); }
    public void useMicroUsb() { System.out.println("MicroUsb connected."); }
}

class iPhone implements LightningPhone {
    public void recharge() { System.out.println("Recharging iPhone..."); }
    public void useLightning() { System.out.println("Lightning connected."); }
}

// 3. Adapter Class
// We want to use an Android Phone (MicroUsb) but the charger is for Lightning
class LightningToMicroUsbAdapter implements LightningPhone {
    private final MicroUsbPhone microUsbPhone;

    public LightningToMicroUsbAdapter(MicroUsbPhone microUsbPhone) {
        this.microUsbPhone = microUsbPhone;
    }

    @Override
    public void useLightning() {
        System.out.println("Adapter converts Lightning signal to MicroUsb...");
        microUsbPhone.useMicroUsb();
    }

    @Override
    public void recharge() {
        microUsbPhone.recharge();
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        AndroidPhone android = new AndroidPhone();
        
        // I want to charge my Android, but I only have a Lightning Cable object.
        // LightningPhone cable = new AndroidPhone(); // Compile Error!

        // Use Adapter
        LightningPhone adapter = new LightningToMicroUsbAdapter(android);
        
        adapter.useLightning();
        adapter.recharge();
    }
}
```

---

## ✅ When to use?

*   **Legacy Code integration**: When you want to use an existing class, but its interface does not match the one you need (e.g., integrating a modern Analytics library into a legacy app).
*   **Third-party libraries**: When you want to reuse several existing subclasses that lack some common functionality that can't be added to the superclass.
*   **Interface translation**: When you need to convert data formats between systems (e.g., XML to JSON) transparently.
