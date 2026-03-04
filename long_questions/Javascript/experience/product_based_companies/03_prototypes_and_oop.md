# ⚡ 03 — Prototypes, OOP & Design Patterns
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Prototype chain and inheritance
- `class` syntax vs prototype-based inheritance
- `instanceof`, `Object.create`
- Mixins
- SOLID principles in JavaScript
- Design patterns: Singleton, Observer, Factory, Module, Proxy

---

## ❓ Most Asked Questions

### Q1. Explain the prototype chain in depth.

```javascript
// Every object has [[Prototype]] (accessible via __proto__ or Object.getPrototypeOf)
// This forms a chain ending at Object.prototype.__proto__ === null

function Animal(name) {
    this.name = name;
}
Animal.prototype.speak = function() {
    return `${this.name} makes a noise.`;
};

function Dog(name) {
    Animal.call(this, name); // inherit instance properties
}
// Set up prototype chain: Dog.prototype → Animal.prototype → Object.prototype → null
Dog.prototype = Object.create(Animal.prototype);
Dog.prototype.constructor = Dog; // restore constructor reference

Dog.prototype.bark = function() {
    return `${this.name} barks.`;
};

const d = new Dog("Rex");
d.bark();   // own prototype
d.speak();  // inherited from Animal.prototype
d.toString(); // inherited from Object.prototype

// Property lookup chain:
// d → Dog.prototype → Animal.prototype → Object.prototype → null

// Verify
Object.getPrototypeOf(d) === Dog.prototype;          // true
Object.getPrototypeOf(Dog.prototype) === Animal.prototype; // true
d instanceof Dog;    // true
d instanceof Animal; // true

// hasOwnProperty: only own properties (not prototype)
d.hasOwnProperty("name");  // true (set via Animal.call)
d.hasOwnProperty("bark");  // false (on Dog.prototype)
```

---

### Q2. Implement classical inheritance without class syntax.

```javascript
// Vehicle → Car → ElectricCar inheritance chain

function Vehicle(make, model) {
    this.make = make;
    this.model = model;
    this.speed = 0;
}

Vehicle.prototype.accelerate = function(amount) {
    this.speed += amount;
    return this;
};

Vehicle.prototype.toString = function() {
    return `${this.make} ${this.model} @ ${this.speed}km/h`;
};

// --- Car extends Vehicle ---
function Car(make, model, doors) {
    Vehicle.call(this, make, model); // super()
    this.doors = doors;
}
Car.prototype = Object.create(Vehicle.prototype);
Car.prototype.constructor = Car;

Car.prototype.honk = function() { return "Beep!"; };

// --- ElectricCar extends Car ---
function ElectricCar(make, model, doors, range) {
    Car.call(this, make, model, doors); // super()
    this.range = range;
    this.battery = 100;
}
ElectricCar.prototype = Object.create(Car.prototype);
ElectricCar.prototype.constructor = ElectricCar;

ElectricCar.prototype.charge = function() {
    this.battery = 100;
    return this;
};

const tesla = new ElectricCar("Tesla", "Model 3", 4, 300);
tesla.accelerate(60).accelerate(20);
console.log(`${tesla}`); // "Tesla Model 3 @ 80km/h"
tesla.charge(); // ElectricCar method
tesla.honk();   // Car method
```

---

### Q3. How does ES6 `class` work under the hood?

```javascript
// ES6 class is SYNTACTIC SUGAR over prototype-based inheritance
// No new object model — JS remains prototype-based

class Animal {
    #name; // private field (truly private, not prototype-based)

    constructor(name) {
        this.#name = name;
    }

    speak() {
        return `${this.#name} makes a noise.`;
    }

    // Static method: on Animal constructor, NOT prototype
    static create(name) {
        return new Animal(name);
    }

    get name() { return this.#name; }
    set name(v) { this.#name = v.trim(); }
}

class Dog extends Animal {
    #tricks = [];

    constructor(name) {
        super(name); // must call before using 'this'
    }

    learn(trick) {
        this.#tricks.push(trick);
        return this;
    }

    // Override parent method
    speak() {
        return `${super.speak()} Specifically, ${this.name} barks.`;
    }
}

// Under the hood (equivalent ES5):
// - Animal.prototype.speak is the same as class method
// - Private fields (#name) are truly private — NOT on prototype
// - 'extends' sets up [[Prototype]] chain
// - 'static' adds to Animal (not Animal.prototype)

const d = new Dog("Rex");
typeof Animal;                 // "function"
Animal.prototype.speak;        // function speak() {...}
d.hasOwnProperty("speak");    // false — on prototype
```

---

### Q4. What are Mixins and when to use them?

```javascript
// Mixins: way to add functionality to classes without full inheritance
// Solves: "diamond problem", multiple inheritance limitations

// Mixin factories
const Serializable = (Base) => class extends Base {
    serialize() {
        return JSON.stringify(this);
    }

    static deserialize(json) {
        return Object.assign(new this(), JSON.parse(json));
    }
};

const Timestamped = (Base) => class extends Base {
    constructor(...args) {
        super(...args);
        this.createdAt = new Date();
        this.updatedAt = new Date();
    }

    touch() {
        this.updatedAt = new Date();
        return this;
    }
};

const Validatable = (Base) => class extends Base {
    validate() {
        const errors = [];
        for (const [field, rule] of Object.entries(this.constructor.rules || {})) {
            if (!rule(this[field])) errors.push(`${field} is invalid`);
        }
        return errors;
    }
};

// Compose mixins
class User extends Serializable(Timestamped(Validatable(class {}))) {
    static rules = {
        email: v => v?.includes("@"),
        name:  v => v?.length >= 2
    };

    constructor(name, email) {
        super();
        this.name = name;
        this.email = email;
    }
}

const user = new User("Alice", "alice@example.com");
user.validate();            // []
const json = user.serialize();
user.touch();               // updates updatedAt
```

---

### Q5. Implement the Proxy pattern for validation and logging.

```javascript
// Proxy: intercept and customize object operations

// ✅ Real-world: Validation proxy
function createValidatedModel(target, schema) {
    return new Proxy(target, {
        set(obj, prop, value) {
            if (schema[prop]) {
                const error = schema[prop](value);
                if (error) throw new TypeError(`${prop}: ${error}`);
            }
            obj[prop] = value;
            return true;
        },

        get(obj, prop) {
            if (prop === 'toJSON') return () => ({ ...obj });
            return obj[prop];
        }
    });
}

const userSchema = {
    age:   v => (Number.isInteger(v) && v >= 0 && v <= 150) ? null : "must be 0–150",
    email: v => v?.includes("@") ? null : "must be valid email",
    name:  v => v?.length >= 2   ? null : "must be ≥ 2 chars"
};

const user = createValidatedModel({}, userSchema);
user.name = "Alice"; // ✅
user.age = 25;       // ✅
user.age = -1;       // ❌ TypeError: age: must be 0–150

// ✅ Logging/observability proxy
function createObservable(target, onChange) {
    return new Proxy(target, {
        set(obj, prop, value) {
            const oldValue = obj[prop];
            obj[prop] = value;
            onChange({ prop, oldValue, newValue: value });
            return true;
        }
    });
}

const state = createObservable(
    { count: 0, name: "app" },
    change => console.log(`[State] ${change.prop}: ${change.oldValue} → ${change.newValue}`)
);
state.count = 1; // [State] count: 0 → 1

// ✅ Lazy loading proxy
function lazyLoad(loader) {
    let instance;
    return new Proxy({}, {
        get(_, prop) {
            if (!instance) instance = loader();
            return instance[prop];
        }
    });
}
```

---

### Q6. Implement the Singleton and Factory patterns.

```javascript
// Singleton — one instance globally
class DatabaseConnection {
    static #instance = null;
    #pool;

    constructor() {
        if (DatabaseConnection.#instance) {
            return DatabaseConnection.#instance;
        }
        this.#pool = this.#createPool();
        DatabaseConnection.#instance = this;
    }

    #createPool() {
        return { maxConnections: 10, active: 0 };
    }

    static getInstance() {
        if (!DatabaseConnection.#instance) {
            new DatabaseConnection();
        }
        return DatabaseConnection.#instance;
    }

    query(sql) {
        return `Result of: ${sql}`;
    }
}

const db1 = DatabaseConnection.getInstance();
const db2 = DatabaseConnection.getInstance();
console.log(db1 === db2); // true

// Factory Pattern — create objects without specifying exact class
class NotificationFactory {
    static create(type, config) {
        const strategies = {
            email:  () => new EmailNotification(config),
            sms:    () => new SMSNotification(config),
            push:   () => new PushNotification(config),
            slack:  () => new SlackNotification(config)
        };

        const factory = strategies[type];
        if (!factory) throw new Error(`Unknown notification type: ${type}`);
        return factory();
    }
}

// Usage
const notification = NotificationFactory.create("email", {
    to: "user@example.com",
    template: "welcome"
});
notification.send("Hello!");
```

---

### Q7. Implement the Observer / Pub-Sub pattern.

```javascript
// Observer: objects subscribe to events and are notified on change
// Decouples producers from consumers

class Store {
    #state;
    #subscribers = new Map(); // event → Set of handlers

    constructor(initialState) {
        this.#state = { ...initialState };
    }

    // Subscribe to state changes for a specific key
    subscribe(key, handler) {
        if (!this.#subscribers.has(key)) {
            this.#subscribers.set(key, new Set());
        }
        this.#subscribers.get(key).add(handler);

        // Return unsubscribe function
        return () => this.#subscribers.get(key).delete(handler);
    }

    setState(updates) {
        const prev = { ...this.#state };
        this.#state = { ...this.#state, ...updates };

        // Notify only affected subscribers
        Object.keys(updates).forEach(key => {
            if (this.#subscribers.has(key)) {
                this.#subscribers.get(key).forEach(handler =>
                    handler(this.#state[key], prev[key])
                );
            }
        });
    }

    getState() { return { ...this.#state }; }
}

// Usage — like a mini Redux
const store = new Store({ count: 0, user: null });

const unsubscribe = store.subscribe("count", (newVal, oldVal) => {
    console.log(`count changed: ${oldVal} → ${newVal}`);
});

store.setState({ count: 1 }); // "count changed: 0 → 1"
store.setState({ count: 2 }); // "count changed: 1 → 2"
unsubscribe();                 // stop listening
store.setState({ count: 3 }); // no log
```

---

### Q8. What is the Module pattern and how does ES Modules differ?

```javascript
// Module pattern (pre-ES6): IIFE to create private scope
const Cart = (function() {
    // private
    let items = [];
    let total = 0;

    function calcTotal() {
        return items.reduce((sum, item) => sum + item.price * item.qty, 0);
    }

    // public API
    return {
        addItem(item) {
            items.push(item);
            total = calcTotal();
            return this;
        },
        removeItem(id) {
            items = items.filter(i => i.id !== id);
            total = calcTotal();
            return this;
        },
        getTotal() { return total; },
        getItems() { return [...items]; } // return copy, not reference
    };
})();

Cart.addItem({ id: 1, name: "Book", price: 20, qty: 2 });
Cart.getTotal(); // 40

// ES Modules (ESM) — native, static, tree-shakeable
// cart.js
export class Cart {
    #items = [];
    addItem(item) { this.#items.push(item); }
    get total() { return this.#items.reduce((s, i) => s + i.price, 0); }
}
export const TAX_RATE = 0.18;
export default Cart; // default export

// main.js
import Cart, { TAX_RATE } from './cart.js';
import * as CartModule from './cart.js'; // namespace import

// ESM advantages over CommonJS/IIFE:
// ✅ Static analysis (tree-shaking possible)
// ✅ Async imports (code splitting): const { feature } = await import('./feature.js')
// ✅ Live bindings (exports reflect current value)
// ✅ Strict mode by default
// ✅ Top-level await support
```

---

### Q9. Implement the Strategy and Command patterns.

```javascript
// Strategy: define family of algorithms, encapsulate each, make them interchangeable

class SortContext {
    #strategy;

    setStrategy(strategy) { this.#strategy = strategy; return this; }
    sort(data) { return this.#strategy.sort([...data]); }
}

const BubbleSort = {
    sort(arr) {
        // O(n²) — small arrays
        for (let i = 0; i < arr.length; i++)
            for (let j = 0; j < arr.length - i - 1; j++)
                if (arr[j] > arr[j + 1]) [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        return arr;
    }
};

const QuickSort = {
    sort(arr) {
        if (arr.length <= 1) return arr;
        const pivot = arr[Math.floor(arr.length / 2)];
        const left  = arr.filter(x => x < pivot);
        const mid   = arr.filter(x => x === pivot);
        const right = arr.filter(x => x > pivot);
        return [...this.sort(left), ...mid, ...this.sort(right)];
    }
};

const ctx = new SortContext();
ctx.setStrategy(arr.length < 20 ? BubbleSort : QuickSort)
   .sort([3, 1, 4, 1, 5, 9]);

// Command: encapsulate requests as objects (supports undo/redo)
class TextEditor {
    #text = "";
    #history = [];
    #redoStack = [];

    execute(command) {
        this.#text = command.execute(this.#text);
        this.#history.push(command);
        this.#redoStack = []; // clear redo on new action
        return this;
    }

    undo() {
        const cmd = this.#history.pop();
        if (cmd) {
            this.#text = cmd.undo(this.#text);
            this.#redoStack.push(cmd);
        }
        return this;
    }

    get text() { return this.#text; }
}

const InsertCommand = (str, pos) => ({
    execute: text => text.slice(0, pos) + str + text.slice(pos),
    undo:    text => text.slice(0, pos) + text.slice(pos + str.length)
});

const editor = new TextEditor();
editor.execute(InsertCommand("Hello", 0));
editor.execute(InsertCommand(" World", 5));
console.log(editor.text); // "Hello World"
editor.undo();
console.log(editor.text); // "Hello"
```

---

### Q10. SOLID Principles in JavaScript — with examples.

```javascript
// S — Single Responsibility Principle
// ❌ Bad: one class doing too much
class UserManager {
    createUser(data) { /* DB logic */ }
    sendWelcomeEmail(user) { /* Email logic */ }
    generateReport(users) { /* Report logic */ }
}

// ✅ Good: separate concerns
class UserRepository  { save(user) { /* DB */ } }
class EmailService    { sendWelcome(user) { /* Email */ } }
class ReportGenerator { generate(users) { /* Report */ } }

// O — Open/Closed Principle (open for extension, closed for modification)
// ❌ Bad: must modify class to add new discount type
class Discount {
    calculate(type, price) {
        if (type === "seasonal") return price * 0.9;
        if (type === "loyalty") return price * 0.85;
        // Adding new type requires modifying here
    }
}

// ✅ Good: extend via new strategies
const discounts = {
    seasonal: price => price * 0.9,
    loyalty:  price => price * 0.85,
    vip:      price => price * 0.7,    // just add new entry
};
const applyDiscount = (type, price) => (discounts[type] ?? (p => p))(price);

// L — Liskov Substitution: subclass must be usable anywhere parent is used
// I — Interface Segregation: don't force implementing unused methods
// D — Dependency Inversion: depend on abstractions, not concretions

// DI example
class OrderService {
    constructor(
        private paymentGateway,  // abstract — could be Stripe, Razorpay, etc.
        private emailService,
        private logger
    ) {}

    async processOrder(order) {
        await this.paymentGateway.charge(order.amount);
        await this.emailService.sendConfirmation(order);
        this.logger.log(`Order ${order.id} processed`);
    }
}
```
