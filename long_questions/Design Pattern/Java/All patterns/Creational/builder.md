# Builder Pattern

## 🟢 What is it?
The **Builder Pattern** lets you construct complex objects step by step. It allows you to produce different types and representations of an object using the same construction code.

Think of it like **Subway (Sandwich Shop)**:
*   You don't just say "Give me a Sandwich" (Constructor with 10 nulls).
*   You say: "Start with Italian Bread" -> "Add Turkey" -> "Add Cheese" -> "No Onions" -> "Add Mayo".
*   The process is step-by-step, and the final result depends on your choices.

---

## 🎯 Strategy to Implement

1.  **Private Constructor**: Make the main class constructor private so it can't be instantiated directly.
2.  **Static Inner Class (Builder)**: Create a static inner class named `Builder`.
3.  **Mutable Fields**: Copy the fields from the main class to the Builder class.
4.  **Chainable Methods**: Create separate methods for setting each field in the Builder. Each method should return `this` (the Builder object itself) to allow chaining (`.setX().setY()`).
5.  **Build Method**: Create a `build()` method in the Builder that calls the private constructor of the main class, passing the Builder object to it.

---

## 💻 Code Example

```java
public class User {
    // All final fields (Immutable object)
    private final String firstName; // Required
    private final String lastName;  // Required
    private final int age;          // Optional
    private final String phone;     // Optional
    private final String address;   // Optional

    // 1. Private Constructor takes the Builder
    private User(UserBuilder builder) {
        this.firstName = builder.firstName;
        this.lastName = builder.lastName;
        this.age = builder.age;
        this.phone = builder.phone;
        this.address = builder.address;
    }

    // Getters only...

    @Override
    public String toString() {
        return "User: " + firstName + " " + lastName + ", Age: " + age + ", Phone: " + phone;
    }

    // 2. Static Inner Builder Class
    public static class UserBuilder {
        private final String firstName;
        private final String lastName;
        private int age;
        private String phone;
        private String address;

        // Constructor with Required parameters
        public UserBuilder(String firstName, String lastName) {
            this.firstName = firstName;
            this.lastName = lastName;
        }

        // 3. Chainable Setter Methods
        public UserBuilder age(int age) {
            this.age = age;
            return this; // Return builder for chaining
        }

        public UserBuilder phone(String phone) {
            this.phone = phone;
            return this;
        }

        public UserBuilder address(String address) {
            this.address = address;
            return this;
        }

        // 4. Build Method to return the final object
        public User build() {
            // Optional: Validate logic here (e.g., if age < 0 throw error)
            return new User(this);
        }
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        
        // Clean, readable, and flexible
        User user = new User.UserBuilder("John", "Doe")
                        .age(30)
                        .phone("123-456-7890")
                        .build();

        // Another variation
        User user2 = new User.UserBuilder("Jane", "Smith")
                        .address("123 Main St") // No age or phone
                        .build();

        System.out.println(user);
    }
}
```

---

## ✅ When to use?

*   **Telescoping Constructor Problem**: When your class has a constructor with 10 parameters, and you have to pass `null` or `0` for 7 of them.
*   **Complex Creation**: When creating an object requires steps that might fail or validaton logic (e.g., "If x is set, then y must be set").
*   **Immutability**: When you want "Immutable Objects" (objects that cannot be changed after creation) but they have many fields. The Builder sets them all up initially and then locks them in.
