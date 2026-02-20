# 🗂️ JavaScript Interview Questions — Master Index

> All answers follow the format from the Go interview bible: **Spoken Answer** (interview-style) + **Indepth** (engineering depth).
> Organized by experience level and target company type.

---

## 📁 File Structure

| File | Topics | Level | Target Companies |
|---|---|---|---|
| [01_Basics.md](./01_Basics.md) | Data types, var/let/const, hoisting, coercion, DOM basics | 🟢 Fresher / Junior | TCS, Infosys, Wipro, Cognizant, Early startups |
| [02_Functions_Closures_Scope.md](./02_Functions_Closures_Scope.md) | Closures, IIFE, currying, pure functions, memoization, ES6 features | 🟡 Junior / Mid | HCL, LTI, Tech Mahindra, Mid-size products |
| [03_OOP_Prototypes_Patterns.md](./03_OOP_Prototypes_Patterns.md) | Prototype chain, classes, inheritance, design patterns, browser APIs | 🟠 Mid / Senior | Mphasis, Hexaware, Zomato, Swiggy, Paytm |
| [04_Async_EventLoop_ES6Plus.md](./04_Async_EventLoop_ES6Plus.md) | Promises, async/await, event loop, ES2020+, TypeScript | 🔴 Senior | Razorpay, PhonePe, Flipkart, Meesho |
| [05_Advanced_Internals.md](./05_Advanced_Internals.md) | TDZ, Object.freeze, strict mode, Symbols, execution context, V8, GC | 🔴 Senior / Expert | Accenture FSI, IBM, Google, Microsoft, Amazon |
| [06_Performance_Security_Testing.md](./06_Performance_Security_Testing.md) | Memory leaks, layout thrashing, XSS, CSRF, testing, mocking | 🔴 Expert | Netflix, Uber, CRED, Groww, Zepto |
| [07_Proxy_Reflect_Workers_Generators.md](./07_Proxy_Reflect_Workers_Generators.md) | Proxy, Reflect, generators, async generators, Web Workers, SharedArrayBuffer | 🟣 Staff / Principal | Atlassian, Adobe, Freshworks, Hasura |
| [08_RealWorld_Patterns_Implementations.md](./08_RealWorld_Patterns_Implementations.md) | Promise queue, concurrency, undo/redo, optimistic UI, FLIP, VDOM | 🟣 Expert / Architect | Goldman Sachs, Zerodha, Razorpay, Stripe |

---

## 🎯 Interview Strategy by Company Type

### 🏢 Service-Based Companies (TCS, Infosys, Wipro, Cognizant, Accenture)
Focus on **Files 01–04**. Expect:
- Core JS concepts (hoisting, closures, event loop)
- ES6+ syntax (arrow functions, destructuring, spread)
- Promise-based async patterns
- DOM manipulation basics
- 1–2 coding problems (debounce, deep clone, polyfills)

### 🚀 Product-Based Companies — Mid Tier (Paytm, Meesho, Swiggy, Groww)
Focus on **Files 02–06**. Expect:
- System design for frontend (state management, caching)
- Performance optimization (lazy loading, bundle size)
- Custom implementations (event emitter, Promise polyfill)
- Testing strategies (mocking, async tests)
- Security (XSS, CSRF prevention)

### 🔥 Product-Based Companies — Senior (Flipkart, Razorpay, PhonePe, CRED)
Focus on **Files 04–07**. Expect:
- Deep async patterns (race conditions, cancellation, retries)
- V8 internals and memory management
- Complex TypeScript scenarios
- Architecture-level discussions (micro-frontends, state management)
- Build tooling (Webpack, Vite, tree-shaking)

### 💎 FAANG / Staff+ (Google, Microsoft, Amazon, Atlassian, Stripe)
Focus on **All Files, especially 07–08**. Expect:
- Proxy and Reflect deep dives
- Web Workers and concurrent programming
- Implementing core abstractions from scratch
- Discussion of trade-offs in framework design
- Real-world performance bottlenecks and solutions

---

## 📊 Question Distribution

| Level | Questions | File(s) |
|---|---|---|
| Fresher (0–1 yr) | Q1–30 | 01 |
| Junior (1–2 yr) | Q31–55 | 02 |
| Mid-level (2–4 yr) | Q56–100 | 03, 04 |
| Senior (4–6 yr) | Q101–130 | 05, 06 |
| Staff/Principal (6+ yr) | Q131–160 | 07, 08 |

---

## 🔑 Must-Know Topics by Company Type

| Topic | Service | Product Mid | Product Senior | FAANG |
|---|---|---|---|---|
| Closures & Scope | ✅ | ✅ | ✅ | ✅ |
| Event Loop | ✅ | ✅ | ✅ | ✅ |
| Promises & async/await | ✅ | ✅ | ✅ | ✅ |
| Prototype & OOP | ✅ | ✅ | ✅ | ✅ |
| ES6+ Features | ✅ | ✅ | ✅ | ✅ |
| Debounce/Throttle | ⬜ | ✅ | ✅ | ✅ |
| Custom Implementations | ⬜ | ✅ | ✅ | ✅ |
| Memory & GC | ⬜ | ⬜ | ✅ | ✅ |
| V8 Internals | ⬜ | ⬜ | ✅ | ✅ |
| Proxy & Reflect | ⬜ | ⬜ | ⬜ | ✅ |
| Web Workers | ⬜ | ⬜ | ⬜ | ✅ |
| Generators | ⬜ | ⬜ | ✅ | ✅ |
| Security (XSS/CSRF) | ⬜ | ✅ | ✅ | ✅ |
| Testing (Jest/Playwright) | ⬜ | ✅ | ✅ | ✅ |
| TypeScript | ⬜ | ✅ | ✅ | ✅ |

---

*Generated from 1000 ChatGPT-curated JavaScript interview questions. Format inspired by the Golang interview answers bible.*
