# JavaScript Interview Questions & Answers (851-950)

## 🔹 29. Language Internals, Quirks & Grammar (Questions 851-880)

**Q851: Why is `null` an object in JavaScript?**
It is a historical bug in the first implementation of JavaScript. The type tag for objects was `000` and `null` was represented as the NULL pointer (`0x00`), so the type check erroneously reported "object". It was never fixed to prevent breaking legacy code.

**Q852: What is the lexical environment in JavaScript?**
An internal data structure that holds identifier-variable mapping. It consists of:
1.  **Environment Record:** Storage of variables/functions.
2.  **Reference to Outer Environment:** Scope chain.

**Q853: What are environment records in execution contexts?**
The actual storage structure inside a Lexical Environment. Two types:
*   **Declarative:** Stores `let`, `const`, `function`.
*   **Object:** Stores bindings for `window` (global) or `with`.

**Q854: What is a binding identifier?**
The generic term for the identifier name (variable name) associated with a value in an environment record.

**Q855: How are function parameters stored in memory?**
They are initialized in the function's Execution Context (Environment Record) as local variables during the Creation Phase.

**Q856: How do arrow functions capture `this` differently than normal functions?**
Arrow functions do not have their own `this` binding. They resolve `this` lexically by looking up the scope chain (like a normal variable) to find the nearest non-arrow parent execution context.

**Q857: What is the purpose of `super()` and when must it be called?**
In a derived class (using `extends`), `super()` calls the parent constructor. It **must** be called before accessing `this` in the constructor, otherwise a `ReferenceError` occurs.

**Q858: Why is `typeof NaN === 'number'`?**
`NaN` (Not-a-Number) is defined in the IEEE-754 standard as a specific floating-point bit pattern representing an invalid numeric result. Technically, it is a numeric value.

**Q859: What is a completion record in JavaScript?**
Internal specification type used to explain control flow. It contains: `[[Type]]` (normal, break, throw, return), `[[Value]]`, and `[[Target]]` (label).

**Q860: How does JavaScript handle dangling commas in objects and arrays?**
It ignores them.
`[1, 2,]` is length 2. `{ a: 1, }` is valid.
(Exception: JSON does not allow dangling commas).

---

## 🔹 30. Error Handling, Custom Errors & Try/Catch Deep Dive (Questions 881-910)
*(Source list 861-870 in Header 881-910)*

**Q861: What’s the difference between a runtime and syntax error?**
*   **Syntax Error:** Detected during parsing logic (before code runs). Stops execution immediately.
*   **Runtime Error:** Occurs during execution (e.g., calling undefined function).

**Q862: How do you throw a custom error in JavaScript?**
`throw new Error('Custom Message')`. Or create a subclass `class MyError extends Error {}` and throw `new MyError()`.

**Q863: What is the `Error.captureStackTrace()` method?**
(V8 specific). It captures the current stack trace and stores it in the `.stack` property of the error object. Often used in custom error constructors to hide the constructor call itself from the trace.

**Q864: How does the stack trace behave in async functions?**
Historically, async traces were cut off at the `await` boundary. Modern engines (Zero-Cost Async Stack Traces) rebuild the async chain to point back to the function that initiated the promise.

**Q865: What happens if you throw inside a finally block?**
The exception thrown in `finally` overwrites any previous exception or return value from the `try/catch` block. The original error is lost.

**Q866: What is a rethrow pattern?**
Catching an error, logging/handling specific cases, and then `throw error` again to let higher-level handlers deal with it.

**Q867: Can `finally` override the return value?**
Yes. If `try` has `return 1` and `finally` has `return 2`, the function returns `2`.

**Q868: How do you handle promise rejections without `catch()`?**
You don't (safely). Unhandled rejections trigger `unhandledrejection` event (browser) or `unhandledRejection` (Node), which may crash the process. Always `.catch()` or use `try/catch` with async/await.

**Q869: What is `window.onerror` used for?**
A global event handler for uncaught exceptions. It receives `msg, url, line, col, error`. Useful for logging errors to analytics.

**Q870: How do you define a global error handler in Node.js?**
```javascript
process.on('uncaughtException', (err) => {
    console.error('Fatal:', err);
    process.exit(1); // Usually best to restart process
});
```

---

## 🔹 31. Object-Oriented Programming (Questions 911-940)
*(Source list 871-880)*

**Q871: What is method overriding in JavaScript classes?**
Defining a method in a child class with the same name as one in the parent class. The child's method takes precedence on instances of the child.

**Q872: How does inheritance work with ES6 classes?**
Using `class Child extends Parent`. Under the hood, it sets up prototype chains: `Child.prototype.__proto__ === Parent.prototype`.

**Q873: How do static properties differ from instance properties?**
static properties are attached to the Class constructor function itself, not the `prototype`. They are accessed as `Class.prop`, not `this.prop`.

**Q874: What is a class expression in JavaScript?**
Defining a class and assigning it to a variable.
`const MyClass = class { ... };`

**Q875: Can you instantiate a class without `new`?**
No. ES6 classes have a "class constructor check" that throws `TypeError: Class constructor cannot be invoked without 'new'`.

**Q876: What’s the difference between `super.method()` and `this.method()`?**
`super.method()` looks up the method on the parent prototype chain. `this.method()` starts lookup on the instance (could be overridden).

**Q877: How can a class extend a built-in like `Array`?**
`class MyArray extends Array {}`. Since ES6, this works correctly, preserving behavior like `.map()` returning instances of `MyArray` (via `Symbol.species`).

**Q878: What happens when you `return` an object from a constructor?**
That object overrides `this` and becomes the result of the `new` expression. If you return a primitive, it is ignored (returns `this`).

**Q879: What’s the output if you call a class without `new`?**
TypeError. (See Q875).

**Q880: How do private class fields differ from closure-based privacy?**
Private fields (`#prop`) are a language feature stored in internal slots, truly hard-private. Closures rely on scope unavailability. Memory impact is usually better with private fields.

---

## 🔹 32. Functional Programming Concepts (Questions 941-970)
*(Source list 881-890)*

**Q881: What is a monad in JavaScript (simple explanation)?**
A structure (like `Promise` or `Array`) that wraps a value and provides a standard way to transform (map) or chain operations (`flat`, `then`) while handling side effects/context automatically.

**Q882: What’s the difference between `compose` and `pipe` functions?**
*   **Compose:** `f(g(x))` (Right-to-Left execution).
*   **Pipe:** `x -> g -> f` (Left-to-Right execution).

**Q883: How can you implement a curry function manually?**
```javascript
function curry(fn) {
    return function curried(...args) {
        if (args.length >= fn.length) return fn.apply(this, args);
        return function(...next) {
            return curried.apply(this, args.concat(next));
        }
    }
}
```

**Q884: What’s the difference between partial application and currying?**
*   **Currying:** Transforming into a chain of single-argument functions.
*   **Partial Application:** Fixing *some* arguments of a function to create a new one with fewer arguments.

**Q885: What is immutability and how do you enforce it in JS?**
Not modifying existing objects. Enforce via `Object.freeze()`, `const` (for assignment), or using libraries like Immutable.js / Immer.

**Q886: What are transducers?**
Composable algorithmic transformations. They allow combining multiple mappers/filters into a single reduction pass, avoiding intermediate array creation. (Efficient for large data).

**Q887: How does `reduceRight()` differ from `reduce()`?**
It iterates the array from the last element to the first (Right-to-Left).

**Q888: What is a functor in functional programming with JavaScript?**
An object/container that implements a `.map()` method to apply a function to its value(s) and return a new container (e.g., Array).

**Q889: How would you implement `compose()` in JavaScript?**
```javascript
const compose = (...fns) => x => fns.reduceRight((v, f) => f(v), x);
```

**Q890: What are point-free functions?**
Functions defined without mentioning their arguments explicitly. `const getIds = map(prop('id'))` instead of `data => data.map(x => x.id)`.

---

## 🔹 33. Internationalization (i18n) and Formatting (Questions 971-990)
*(Source list 891-900)*

**Q891: How does `Intl.NumberFormat` work?**
Formats numbers based on locale (currency, decimal separators). `new Intl.NumberFormat('de-DE', { style: 'currency' })`.

**Q892: What is the difference between `toLocaleString()` and `Intl.NumberFormat()`?**
`toLocaleString` calls `Intl` under the hood but creates a new formatter instance every time. For performance in loops, create one `Intl.NumberFormat` and reuse it.

**Q893: How to localize date/time for multiple time zones?**
`new Intl.DateTimeFormat('en-US', { timeZone: 'Asia/Tokyo' }).format(date)`.

**Q894: What is the purpose of `Intl.PluralRules`?**
Determines the plural form ("zero", "one", "two", "few", "many") for a number based on locale rules (e.g., Russian has complex rules).

**Q895: What are BCP-47 locale strings?**
Standard strings identifying languages/regions (e.g., `en-US`, `zh-Hans-CN`). Used by `Intl` API.

**Q896: How does `Intl.ListFormat` work?**
Formats arrays into lists (e.g. "A, B, and C").
`new Intl.ListFormat('en', { style: 'long', type: 'conjunction' }).format(['A', 'B'])`.

**Q897: What is the `Intl.RelativeTimeFormat`?**
Formats time difference ("5 minutes ago", "in 2 days").

**Q898: How can you dynamically switch between languages in a JS app?**
By using an i18n library (i18next) or manually reloading `Intl` formatters with the new locale string stored in state/context.

**Q899: What is `Intl.DisplayNames` and its use case?**
Translates region, language, or script names. "US" -> "United States".

**Q900: How does `Intl.Segmenter` help with word tokenization?**
Splits text into meaningful segments (words, sentences) respecting locale rules (handling languages like Japanese without spaces).

---

## 🔹 34. Testing JavaScript (Questions 991-1020)
*(Source list 901-910)*

**Q901: What are the differences between mocking and stubbing?**
*   **Stub:** Provides canned answers to calls (fake simple behavior).
*   **Mock:** Verifies behavior (expects specific calls/params).

**Q902: What’s the purpose of `jest.fn()`?**
Creates a mock function that tracks calls, arguments, and instances, and can mock return values.

**Q903: How to test async functions using `done()` vs `async/await`?**
`async/await` is cleaner. Only use `done()` callback if testing callback-based legacy code.

**Q904: What is snapshot testing?**
Comparing the current output (UI component render tree or JSON) against a stored "golden" snapshot file. Fails if they differ.

**Q905: How does `sinon.spy()` work?**
Wraps a real function to record its execution data (call count, args) without changing its behavior.

**Q906: How to test DOM events using testing libraries?**
`fireEvent.click(element)` or `userEvent.click(element)` (React Testing Library). It synthesizes the event and dispatches it to the node.

**Q907: What are fake timers and when should you use them?**
Injects fake time control (`jest.useFakeTimers`). Allows "fast-forwarding" `setTimeout` or `setInterval` without waiting real time in tests.

**Q908: What is the difference between shallow and deep rendering?**
*   **Shallow:** Renders component one level deep (mocks children).
*   **Deep (Mount):** Renders full tree. (Preferred for integration).

**Q909: How to mock fetch in a unit test?**
`global.fetch = jest.fn(() => Promise.resolve({ json: () => ({...}) }));` or use `msw` (Mock Service Worker).

**Q910: What is mutation testing in JavaScript?**
Strategy where tool (Stryker) intentionally introduces bugs (mutations) into source code and checks if tests fail. If tests pass, the test suite is weak.

---

## 🔹 35. Security-Oriented (Questions 1021-1050)
*(Source list 911-920)*

**Q911: How do you protect against clickjacking with JavaScript?**
(See Q436). Frame-busting script.

**Q912: What is JavaScript sandboxing?**
Running untrusted code in an isolated environment (e.g., iframe with `sandbox` attribute, Web Worker, or `vm` module in Node) with restricted access to global scope/DOM.

**Q913: What is the difference between innerHTML and textContent (security)?**
`textContent` treats input as raw text (auto-escapes). `innerHTML` parses it as HTML (executes scripts).

**Q914: What are DOMPurify or similar libraries used for?**
Sanitizing HTML strings by removing dangerous tags/attributes (`<script>`, `onmouseover`) before inserting into DOM.

**Q915: How can Content Security Policy (CSP) block XSS?**
By defining allowed sources (`script-src 'self'`). It prevents browser from executing inline scripts or scripts from unauthorized domains.

**Q916: Why should you avoid eval in modern apps?**
Performance (impedes optimization) & Security (executes arbitrary code scope). "Eval is evil".

**Q917: What is the difference between escaping and encoding?**
*   **Encoding:** Changing format (URL encode space to `%20`).
*   **Escaping:** Adding markers to treat characters as data, not code (`<` to `&lt;`).

**Q918: How can JavaScript open up vulnerabilities when used with user input?**
By trusting input. DOM XSS (sinks like `innerHTML`, `location.href`).

**Q919: How do CSP nonces work with inline scripts?**
Server generates a random nonce token, puts it in header CSP, and adds it to `<script nonce="...">`. Browser executes only matching scripts.

**Q920: What is subresource integrity (SRI)?**
Attribute `integrity="sha384-..."` on `<script>`. Ensures the file fetched from CDN exactly matches the expected hash, preventing execution if CDN is compromised.
