# 📘 05 — Object-Oriented JavaScript & Design Patterns
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Objects: creation, methods, property descriptors
- Prototype and inheritance basics
- Classes and inheritance
- Common patterns: Singleton, Factory, Observer
- `Object.keys`, `Object.values`, `Object.entries`

---

## ❓ Most Asked Questions

### Q1. How are objects created in JavaScript?

```javascript
// Method 1: Object literal (most common)
const user = {
    id: 1,
    name: "Alice",
    email: "alice@example.com",
    fullName() {
        return this.name;
    }
};

// Method 2: Constructor function (pre-ES6)
function Person(name, age) {
    this.name = name;
    this.age  = age;
}
Person.prototype.greet = function() {
    return `Hi, I'm ${this.name}`;
};
const p = new Person("Bob", 30);

// Method 3: ES6 class
class Product {
    constructor(name, price) {
        this.name  = name;
        this.price = price;
    }
    get discountedPrice() { return this.price * 0.9; }
    toString() { return `${this.name} — ₹${this.price}`; }
}
const laptop = new Product("Laptop", 50000);

// Method 4: Object.create (specify prototype explicitly)
const animalProto = {
    speak() { return `${this.name} makes a sound`; }
};
const dog = Object.create(animalProto);
dog.name = "Rex";
dog.speak(); // "Rex makes a sound"

// Method 5: Factory function (returns object, no 'new' needed)
function createUser(name, role) {
    return {
        name,
        role,
        permissions: role === 'admin' ? ['read', 'write', 'delete'] : ['read'],
        hasPermission(action) { return this.permissions.includes(action); }
    };
}
const admin = createUser("Alice", "admin");
admin.hasPermission("delete"); // true
```

---

### Q2. How does inheritance work in JavaScript?

```javascript
// ES6 class inheritance
class Vehicle {
    constructor(make, model, year) {
        this.make  = make;
        this.model = model;
        this.year  = year;
    }

    info() {
        return `${this.year} ${this.make} ${this.model}`;
    }

    start() { return "Vroom!"; }
}

class ElectricCar extends Vehicle {
    #batteryLevel = 100; // private field

    constructor(make, model, year, range) {
        super(make, model, year); // call parent constructor
        this.range = range;
    }

    // Override parent method
    start() {
        return `${super.info()} — silently starts ⚡`;
    }

    charge(percent) {
        this.#batteryLevel = Math.min(100, this.#batteryLevel + percent);
        return this;
    }

    get batteryStatus() {
        return `Battery: ${this.#batteryLevel}%`;
    }
}

const tesla = new ElectricCar("Tesla", "Model 3", 2024, 560);
tesla.start();                     // "2024 Tesla Model 3 — silently starts ⚡"
tesla.charge(20).batteryStatus;    // "Battery: 100%"
tesla instanceof ElectricCar;      // true
tesla instanceof Vehicle;          // true

// Check prototype chain
Object.getPrototypeOf(tesla) === ElectricCar.prototype; // true
Object.getPrototypeOf(ElectricCar.prototype) === Vehicle.prototype; // true
```

---

### Q3. Implement the Singleton pattern.

```javascript
// Singleton: ensure only ONE instance of a class exists (shared state)
// Use cases: configuration, database connection, cache, logger

class Logger {
    static #instance = null;
    #logs = [];

    constructor() {
        if (Logger.#instance) {
            return Logger.#instance; // return existing instance
        }
        Logger.#instance = this;
    }

    log(level, message) {
        const entry = {
            level,
            message,
            timestamp: new Date().toISOString()
        };
        this.#logs.push(entry);
        console[level]?.(`[${entry.timestamp}] ${message}`);
    }

    info(msg)  { this.log("log", msg); }
    warn(msg)  { this.log("warn", `⚠️ ${msg}`); }
    error(msg) { this.log("error", `❌ ${msg}`); }

    getLogs() { return [...this.#logs]; }

    static getInstance() {
        if (!Logger.#instance) new Logger();
        return Logger.#instance;
    }
}

// Both references point to same object
const log1 = new Logger();
const log2 = Logger.getInstance();

log1.info("Application started");
log2.warn("Low memory");

log1 === log2;              // true — same instance
log1.getLogs().length === 2; // true — shared log array

// Simpler module-level singleton
// database.js
let connection = null;
function getConnection() {
    if (!connection) {
        connection = createDatabaseConnection();
    }
    return connection;
}
export { getConnection }; // module ensures single instance
```

---

### Q4. Implement the Factory pattern.

```javascript
// Factory: create objects without specifying exact class
// Use when object creation logic is complex or type varies

// Simple factory function
function createButton(type) {
    const base = {
        render() { return `<button class="${this.className}">${this.label}</button>`; },
        onClick() { console.log(`${type} button clicked`); }
    };

    const types = {
        primary: { ...base, className: "btn btn-primary", label: "Submit" },
        danger:  { ...base, className: "btn btn-danger",  label: "Delete" },
        outline: { ...base, className: "btn btn-outline", label: "Cancel" }
    };

    if (!types[type]) throw new Error(`Unknown button type: ${type}`);
    return types[type];
}

const submitBtn = createButton("primary");
submitBtn.render(); // <button class="btn btn-primary">Submit</button>

// Class-based abstract factory
class NotificationService {
    static create(channel, config) {
        switch (channel) {
            case 'email': return new EmailNotification(config);
            case 'sms':   return new SMSNotification(config);
            case 'push':  return new PushNotification(config);
            default:      throw new Error(`Unknown channel: ${channel}`);
        }
    }
}

class EmailNotification {
    constructor({ from, to }) { this.from = from; this.to = to; }
    send(message) { return `Email from ${this.from} to ${this.to}: ${message}`; }
}

class SMSNotification {
    constructor({ phone }) { this.phone = phone; }
    send(message) { return `SMS to ${this.phone}: ${message}`; }
}

// Usage
const userPreference = "email";
const notification = NotificationService.create(userPreference, {
    from: "noreply@app.com",
    to: "user@example.com"
});
notification.send("Your order has been shipped!");
```

---

### Q5. What is the Observer pattern? Implement a simple version.

```javascript
// Observer (Pub-Sub): objects subscribe to events; publisher notifies on change
// Used in: React setState, EventEmitter, DOM events, Redux store

class EventEmitter {
    #events = {};

    on(event, listener) {
        if (!this.#events[event]) this.#events[event] = [];
        this.#events[event].push(listener);
        return () => this.off(event, listener); // returns unsubscribe fn
    }

    once(event, listener) {
        const wrapper = (...args) => {
            listener(...args);
            this.off(event, wrapper);
        };
        return this.on(event, wrapper);
    }

    emit(event, ...args) {
        (this.#events[event] || []).forEach(listener => listener(...args));
    }

    off(event, listener) {
        this.#events[event] = (this.#events[event] || []).filter(l => l !== listener);
    }
}

// ✅ Use case: shopping cart with multiple UI components
const cartEmitter = new EventEmitter();

// Header badge component subscribes
const unsubBadge = cartEmitter.on("cartUpdated", ({ itemCount }) => {
    document.getElementById("cart-badge").textContent = itemCount;
});

// Mini cart sidebar subscribes
const unsubSidebar = cartEmitter.on("cartUpdated", ({ items, total }) => {
    renderMiniCart(items, total);
});

// Cart action (trigger events)
function addToCart(product) {
    cart.items.push(product);
    cartEmitter.emit("cartUpdated", {
        items: cart.items,
        itemCount: cart.items.length,
        total: cart.items.reduce((s, p) => s + p.price, 0)
    });
}

// Cleanup
unsubBadge();   // header badge stops listening
unsubSidebar(); // sidebar stops listening
```

---

### Q6. Explain `Object.keys`, `Object.values`, `Object.entries`.

```javascript
const product = {
    id: 1,
    name: "Laptop",
    price: 50000,
    inStock: true
};

// Object.keys: array of own enumerable property NAMES
Object.keys(product);   // ["id", "name", "price", "inStock"]

// Object.values: array of own enumerable property VALUES
Object.values(product); // [1, "Laptop", 50000, true]

// Object.entries: array of [key, value] pairs
Object.entries(product);
// [["id", 1], ["name", "Laptop"], ["price", 50000], ["inStock", true]]

// ✅ Use case: filter object properties
function filterObject(obj, predicate) {
    return Object.fromEntries(
        Object.entries(obj).filter(([key, value]) => predicate(key, value))
    );
}

filterObject(product, (key, value) => typeof value !== 'boolean');
// { id: 1, name: "Laptop", price: 50000 }

// ✅ Use case: transform object values
function mapValues(obj, fn) {
    return Object.fromEntries(
        Object.entries(obj).map(([key, value]) => [key, fn(value)])
    );
}

const pricesInUSD = mapValues({ laptop: 50000, phone: 30000 }, price => price / 83);
// { laptop: 602.4, phone: 361.4 }

// ✅ Use case: object iteration (sorted)
const sortedEntries = Object.entries(product)
    .sort(([a], [b]) => a.localeCompare(b));

// Check if object is empty
const isEmpty = obj => Object.keys(obj).length === 0;
isEmpty({}); // true

// ⚠️ Only returns OWN enumerable properties (not inherited)
class Base { inheritedProp = "hello"; }
const obj = new Base();
obj.ownProp = "world";
Object.keys(obj); // ["inheritedProp", "ownProp"] — class fields are own props
```

---

### Q7. What is method chaining? Implement a query builder.

```javascript
// Method chaining: each method returns 'this' to allow chaining calls
// Makes APIs fluent and readable

class QueryBuilder {
    #table  = "";
    #conditions = [];
    #orderBy    = "";
    #limitVal   = null;
    #selectCols = ["*"];

    constructor(table) {
        this.#table = table;
    }

    select(...cols) {
        this.#selectCols = cols;
        return this; // enables chaining
    }

    where(condition) {
        this.#conditions.push(condition);
        return this;
    }

    order(col, direction = "ASC") {
        this.#orderBy = `ORDER BY ${col} ${direction}`;
        return this;
    }

    limit(n) {
        this.#limitVal = n;
        return this;
    }

    build() {
        let query = `SELECT ${this.#selectCols.join(", ")} FROM ${this.#table}`;
        if (this.#conditions.length) {
            query += ` WHERE ${this.#conditions.join(" AND ")}`;
        }
        if (this.#orderBy) query += ` ${this.#orderBy}`;
        if (this.#limitVal) query += ` LIMIT ${this.#limitVal}`;
        return query;
    }
}

// Fluent, readable API
const query = new QueryBuilder("users")
    .select("id", "name", "email")
    .where("age > 18")
    .where("status = 'active'")
    .order("name", "ASC")
    .limit(10)
    .build();

// "SELECT id, name, email FROM users WHERE age > 18 AND status = 'active' ORDER BY name ASC LIMIT 10"
```
