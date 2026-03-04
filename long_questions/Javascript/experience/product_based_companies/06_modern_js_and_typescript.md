# ⚡ 06 — Modern JavaScript (ES6–ES2024) & TypeScript
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Iterators and iterables
- Symbol and well-known symbols
- Reflect API
- Tagged template literals
- Optional chaining `?.`, nullish coalescing `??`
- Top-level `await`, logical assignment operators
- TypeScript: generics, utility types, conditional types, decorators

---

## ❓ Most Asked Questions

### Q1. Explain Iterators, Iterables, and the Iterator Protocol.

```javascript
// Iterable: has a [Symbol.iterator] method returning an Iterator
// Iterator: has a next() method returning { value, done }

// Custom Range iterable
class Range {
    constructor(start, end, step = 1) {
        this.start = start;
        this.end   = end;
        this.step  = step;
    }

    // Make Range iterable
    [Symbol.iterator]() {
        let current = this.start;
        const { end, step } = this;

        return {
            next() {
                if (current <= end) {
                    const value = current;
                    current += step;
                    return { value, done: false };
                }
                return { value: undefined, done: true };
            },
            // Make iterator itself iterable (common convention)
            [Symbol.iterator]() { return this; }
        };
    }
}

const range = new Range(1, 10, 2);
for (const n of range)   console.log(n); // 1, 3, 5, 7, 9
const arr = [...range];  // [1, 3, 5, 7, 9]
const [first, second] = range; // 1, 3 — destructuring uses iterator

// Built-in iterables: String, Array, Map, Set, NodeList, arguments
// NOT iterable by default: plain Object {} (use Object.entries/keys/values)

// Infinite iterator with early termination
function* naturals(start = 1) {
    while (true) yield start++;
}

const gen = naturals();
const first5 = Array.from({ length: 5 }, () => gen.next().value); // [1,2,3,4,5]

// for...of with break calls iterator.return() (cleanup)
function* resourceStream() {
    console.log("Opening resource");
    try {
        while (true) yield fetchNext();
    } finally {
        console.log("Closing resource"); // called when for...of breaks
    }
}
```

---

### Q2. Explain `Symbol` and well-known Symbols.

```javascript
// Symbol: unique, immutable primitive — primary use: unique property keys
const id = Symbol("id");
const name = Symbol("id"); // same description, DIFFERENT symbol
id === name; // false

// Well-known Symbols: protocol hooks into JS builtins

// Symbol.iterator — custom iteration protocol (see Q1)

// Symbol.toPrimitive — control type coercion
class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }

    [Symbol.toPrimitive](hint) {
        switch (hint) {
            case 'number': return this.amount;
            case 'string': return `${this.amount} ${this.currency}`;
            case 'default': return this.amount; // used with +, ==
        }
    }
}

const price = new Money(100, "USD");
+price;         // 100 (number hint)
`${price}`;     // "100 USD" (string hint)
price + 50;     // 150 (default hint)

// Symbol.hasInstance — customize instanceof
class EvenNumber {
    static [Symbol.hasInstance](value) {
        return Number.isInteger(value) && value % 2 === 0;
    }
}

2 instanceof EvenNumber;  // true
3 instanceof EvenNumber;  // false
4 instanceof EvenNumber;  // true

// Symbol.asyncIterator — async iteration protocol
class AsyncRange {
    constructor(start, end) { this.start = start; this.end = end; }

    async *[Symbol.asyncIterator]() {
        for (let i = this.start; i <= this.end; i++) {
            await sleep(100); // simulate async fetch
            yield i;
        }
    }
}

for await (const n of new AsyncRange(1, 5)) {
    console.log(n); // 1, 2, 3, 4, 5 with 100ms delays
}

// Symbol.species — control which constructor creates derived instances
class MyArray extends Array {
    static get [Symbol.species]() { return Array; } // map/filter return plain Array
}
const my = new MyArray(1, 2, 3);
my.map(x => x * 2) instanceof MyArray; // false — returns Array (due to species)
```

---

### Q3. What are Tagged Template Literals? Build a SQL sanitizer.

```javascript
// Tagged templates: function called with template Parts + interpolated values

// sql sanitizer using tagged templates
function sql(strings, ...values) {
    const sanitized = values.map(val => {
        if (typeof val === 'string') {
            // Escape SQL special chars
            return "'" + val.replace(/'/g, "''") + "'";
        }
        if (typeof val === 'number') return Number(val);
        throw new TypeError("Invalid SQL value type");
    });

    return strings.reduce((query, str, i) => {
        return query + str + (sanitized[i] ?? "");
    }, "");
}

const userId = "1; DROP TABLE users;--"; // SQL injection attempt!
const minAge = 18;

const query = sql`SELECT * FROM users WHERE id = ${userId} AND age > ${minAge}`;
// "SELECT * FROM users WHERE id = '1; DROP TABLE users;--' AND age > 18"
// ✅ Input safely escaped!

// Styled-components-like tagged template
function css(strings, ...values) {
    return strings.reduce((result, str, i) => {
        const val = typeof values[i] === 'function' ? values[i]() : values[i] ?? '';
        return result + str + val;
    }, "");
}

const primaryColor = "#667eea";
const styles = css`
    background: ${primaryColor};
    color: white;
    padding: ${8}px ${16}px;
`;

// i18n with tagged templates
function i18n(strings, ...keys) {
    return (translations) => strings.reduce((result, str, i) => {
        const key = keys[i];
        return result + str + (key ? (translations[key] ?? key) : "");
    }, "");
}

const greet = i18n`Hello ${"name"}, you have ${"count"} messages`;
greet({ name: "Alice", count: 5 }); // "Hello Alice, you have 5 messages"
```

---

### Q4. TypeScript: Generics, Conditional Types, and Utility Types.

```typescript
// Generics: type-safe reusable code

// Generic function
function first<T>(arr: T[]): T | undefined {
    return arr[0];
}

first([1, 2, 3]);          // inferred as number | undefined
first(["a", "b"]);         // inferred as string | undefined

// Generic constraints
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}

const user = { id: 1, name: "Alice", age: 25 };
getProperty(user, "name");  // string
getProperty(user, "id");    // number
// getProperty(user, "xyz"); // Error: "xyz" not in user

// Generic interfaces
interface Repository<T, ID = number> {
    findById(id: ID): Promise<T | null>;
    findAll(): Promise<T[]>;
    save(entity: Omit<T, 'id'>): Promise<T>;
    delete(id: ID): Promise<void>;
}

// Conditional types
type IsArray<T> = T extends any[] ? true : false;
type IsString<T> = T extends string ? true : false;

type A = IsArray<number[]>;  // true
type B = IsArray<string>;    // false

// Infer keyword in conditional types
type UnwrapPromise<T> = T extends Promise<infer U> ? U : T;
type UnwrapArray<T>   = T extends Array<infer U>   ? U : T;

type Resolved = UnwrapPromise<Promise<string>>;  // string
type Element  = UnwrapArray<number[]>;            // number

// Utility types (built-in)
interface Product {
    id: number;
    name: string;
    price: number;
    description?: string;
}

type ProductPreview   = Pick<Product, 'id' | 'name' | 'price'>;
type ProductUpdate    = Partial<Product>;        // all optional
type RequiredProduct  = Required<Product>;       // all required
type ReadonlyProduct  = Readonly<Product>;       // no mutations
type ProductWithoutId = Omit<Product, 'id'>;

// Record type
type HttpStatus = Record<number, string>;
const statuses: HttpStatus = { 200: "OK", 404: "Not Found", 500: "Server Error" };

// Mapped types
type Nullable<T> = { [K in keyof T]: T[K] | null };
type Optional<T> = { [K in keyof T]?: T[K] };

// Template literal types (TS 4.1+)
type EventName<T extends string> = `on${Capitalize<T>}`;
type Click   = EventName<"click">;   // "onClick"
type Submit  = EventName<"submit">;  // "onSubmit"
```

---

### Q5. TypeScript decorators and reflect-metadata.

```typescript
// Decorators: compile-time metaprogramming (Stage 3 proposal)
// Enabling: "experimentalDecorators": true in tsconfig.json

// Class decorator
function Singleton<T extends new (...args: any[]) => {}>(constructor: T) {
    let instance: T;
    return class extends constructor {
        constructor(...args: any[]) {
            if (instance) return instance;
            super(...args);
            instance = this as any;
        }
    };
}

@Singleton
class Config {
    constructor(public env: string) {}
}

// Method decorator: logging
function LogExecutionTime(target: any, key: string, descriptor: PropertyDescriptor) {
    const original = descriptor.value;
    descriptor.value = async function(...args: any[]) {
        const start = performance.now();
        const result = await original.apply(this, args);
        console.log(`${key} took ${performance.now() - start}ms`);
        return result;
    };
    return descriptor;
}

// Property decorator: validation
function MinLength(min: number) {
    return function(target: any, propertyKey: string) {
        let value: string;
        Object.defineProperty(target, propertyKey, {
            get() { return value; },
            set(v: string) {
                if (v.length < min) throw new Error(`${propertyKey} must be ≥ ${min} chars`);
                value = v;
            }
        });
    };
}

class User {
    @MinLength(2)
    name!: string;

    @LogExecutionTime
    async fetchOrders() {
        return await api.get(`/users/${this.name}/orders`);
    }
}

// Parameter decorator + DI container (NestJS pattern)
const Container = new Map();

function Injectable() {
    return (constructor: Function) => {
        Container.set(constructor.name, new constructor());
    };
}

function Inject(token: string) {
    return (target: any, _: string | symbol, index: number) => {
        // Registered in reflection metadata for DI resolution
    };
}
```

---

### Q6. Advanced TypeScript — discriminated unions and exhaustiveness checking.

```typescript
// Discriminated Unions: type safe handling of different shapes

type Shape =
    | { kind: "circle";    radius: number }
    | { kind: "rectangle"; width: number; height: number }
    | { kind: "triangle";  base: number;  height: number };

function area(shape: Shape): number {
    switch (shape.kind) {
        case "circle":    return Math.PI * shape.radius ** 2;
        case "rectangle": return shape.width * shape.height;
        case "triangle":  return 0.5 * shape.base * shape.height;
        default:
            // Exhaustiveness check: if new shape added, this becomes a compile error
            const _exhaustive: never = shape;
            throw new Error(`Unknown shape: ${_exhaustive}`);
    }
}

// Result type: type-safe error handling (no exceptions)
type Result<T, E = Error> =
    | { success: true;  value: T }
    | { success: false; error: E };

function divide(a: number, b: number): Result<number, string> {
    if (b === 0) return { success: false, error: "Division by zero" };
    return { success: true, value: a / b };
}

const result = divide(10, 2);
if (result.success) {
    console.log(result.value); // TypeScript knows value exists here
} else {
    console.error(result.error); // TypeScript knows error exists here
}

// Branded types: prevent misuse of primitive types
type UserId    = string & { readonly brand: unique symbol };
type ProductId = string & { readonly brand: unique symbol };

function createUserId(id: string):    UserId    { return id as UserId; }
function createProductId(id: string): ProductId { return id as ProductId; }

function getUser(id: UserId) { /* ... */ }

const userId    = createUserId("user-123");
const productId = createProductId("prod-456");

getUser(userId);    // ✅
// getUser(productId); // ❌ TypeScript error: ProductId not assignable to UserId
// getUser("raw-id"); // ❌ TypeScript error: won't accept plain string
```

---

### Q7. What are the key ES2022–2024 features?

```javascript
// ES2022
// 1. Class Fields (public, private, static)
class Counter {
    count = 0;                    // public field
    #limit = 100;                 // private field
    static #instances = 0;        // private static
    static { Counter.#instances++; } // static initialization block

    increment() {
        if (this.count >= this.#limit) throw new Error("Limit reached");
        return ++this.count;
    }
}

// 2. Top-level await (ES modules)
const data = await fetch("/api/data").then(r => r.json());
// Can be used at module top level — no async wrapper needed!

// 3. Array.at() and Object.hasOwn()
const arr = [1, 2, 3];
arr.at(-1);           // 3 (last element, negative index)
arr.at(-2);           // 2

Object.hasOwn(obj, 'prop'); // safer than obj.hasOwnProperty (works on Object.create(null))

// ES2023
// 4. Array find from end
[1, 2, 3, 2, 1].findLast(x => x === 2);       // 2 (last match)
[1, 2, 3, 2, 1].findLastIndex(x => x === 2);  // 3 (index of last match)

// 5. Array.toSorted, toReversed, toSpliced, with (non-mutating!)
const sorted = [3, 1, 2].toSorted(); // [1, 2, 3] — original unchanged
const reversed = [1, 2, 3].toReversed(); // [3, 2, 1] — original unchanged
const updated = [1, 2, 3].with(1, 99); // [1, 99, 3] — replace at index

// ES2024
// 6. Object.groupBy
const products = [
    { name: "A", category: "fruit" },
    { name: "B", category: "veg" },
    { name: "C", category: "fruit" }
];
const grouped = Object.groupBy(products, p => p.category);
// { fruit: [{name:"A",...}, {name:"C",...}], veg: [{name:"B",...}] }

// 7. Promise.withResolvers
const { promise, resolve, reject } = Promise.withResolvers();
// Access resolve/reject outside the Promise constructor!
setTimeout(resolve, 1000, "done");
await promise; // "done"

// 8. Atomics.waitAsync (non-blocking)
// Cross-thread sync without blocking main thread
```
