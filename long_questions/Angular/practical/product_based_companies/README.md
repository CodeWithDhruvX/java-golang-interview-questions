# 🚀 Angular Interview Questions — Product-Based Companies

> Companies like **Google, Microsoft, Flipkart, Swiggy, Zomato, PhonePe, CRED, Razorpay, Meesho**, etc.

Product-based companies ask **deep, scenario-based questions** on Angular internals, performance architecture, advanced RxJS patterns, enterprise-scale state management, and system design for the frontend. Expect **whiteboard-style** discussions, architecture decision questions, and modern Angular features.

---

## 📂 Category Index

| # | File | Topic | Difficulty |
|---|------|--------|------------|
| 01 | [Performance Optimization](./01_performance_optimization.md) | OnPush, trackBy, lazy loading, bundle analysis, virtual scroll | 🟡 Medium–Hard |
| 02 | [Advanced RxJS](./02_advanced_rxjs.md) | Higher-order operators, custom operators, schedulers, backpressure | 🔴 Hard |
| 03 | [State Management & NgRx Advanced](./03_state_management_ngrx_advanced.md) | NgRx Effects, Entity adapter, ComponentStore, Facade pattern | 🔴 Hard |
| 04 | [Angular Internals & Compiler](./04_angular_internals_compiler.md) | Ivy, Zone.js, change detection tree, ahead-of-time compilation | 🔴 Hard |
| 05 | [Security, SSR & PWA](./05_security_ssr_pwa.md) | XSS, CSP, Angular Universal (SSR), Service Workers, PWA | 🟡 Medium–Hard |
| 06 | [Signals, Standalone & Modern Angular](./06_signals_standalone_modern.md) | Angular Signals, Standalone Components, new control flow syntax | 🟡 Medium–Hard |
| 07 | [System Design & Angular Architecture](./07_system_design_angular_architecture.md) | Micro frontends, design system, monorepo, module federation | 🔴 Hard |
| 08 | [Advanced Testing](./08_testing_advanced.md) | Component harness, spectator, testing effects, fakeAsync | 🔴 Hard |

---

## 🎯 Interview Focus Areas (Product Companies)

- ✅ **Performance**: OnPush, trackBy, lazy loading, virtual scrolling, bundle budgets
- ✅ **Angular Internals**: Ivy pipeline, Zone.js, LView, change detection algorithm
- ✅ **Advanced RxJS**: `switchMap` vs `mergeMap` vs `concatMap`, custom operators, `Subject` types
- ✅ **NgRx Deep Dive**: Effects, Entity, Facade, ComponentStore, selector memoization
- ✅ **Modern Angular**: Signals, Standalone Components, `@if`/`@for` control flow (Angular 17+)
- ✅ **Security**: XSS prevention, `DomSanitizer`, CSP headers, CSRF
- ✅ **SSR**: Angular Universal, hydration, SEO benefits, transfer state
- ✅ **Architecture**: Micro frontends, monorepo with Nx, module federation
- ✅ **Testing**: Component harness, `fakeAsync`/`tick`, testing NgRx effects

---

## 💡 Tips for Product Company Angular Interviews

1. **Understand Zone.js deeply** — explain what it does, when to opt-out, and `NgZone.runOutsideAngular()`
2. **OnPush + Immutability** — explain why mutating arrays/objects doesn't trigger updates with OnPush
3. **`switchMap` vs `mergeMap` vs `concatMap`** — critical distinction, always asked
4. **NgRx Selector memoization** — explain how `createSelector` uses projector caching
5. **Micro frontends** — Module Federation with Angular allows independently deployed Angular apps
6. **Angular Signals** — the new reactivity system replacing Zone.js for fine-grained updates
7. **Memory leak prevention** — `takeUntil`, `async` pipe, `DestroyRef`, unsubscribing patterns
8. **Bundle analysis** — use `webpack-bundle-analyzer` to identify and reduce chunk sizes
