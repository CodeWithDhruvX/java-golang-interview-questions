# 🗣️ Theory — Angular Internals & Compiler
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is the Ivy compiler and how does it improve on View Engine?"

> *"Ivy is Angular's current compilation and rendering engine, the default since Angular 9. The previous engine was View Engine, and Ivy is a ground-up rewrite. The key improvement is locality: Ivy can compile each component independently without needing to know about the whole app or its NgModule graph. View Engine required the compiler to see the entire module context to compile a component. This locality unlocks genuine tree-shaking — if a component is never used, its code is never included in the bundle, which wasn't fully true with View Engine. Ivy also produces human-readable compiled output — the component's template compiles to JavaScript instructions like ɵɵelementStart, ɵɵtext, ɵɵproperty — you can actually read and understand what Angular is doing. Bundle sizes shrunk, build times improved, and component-level lazy loading became possible."*

---

## Q: "What is Zone.js? What does it actually do under the hood?"

> *"Zone.js patches the browser's async APIs — setTimeout, setInterval, XMLHttpRequest, fetch, Promise.prototype.then, addEventListener — all of them. When Zone.js patches these, it wraps them with code that runs Angular's change detection after the callback completes. So when you click a button — Zone.js intercepts the event handler — Angular's change detection runs — the DOM updates. When an HTTP call resolves — Zone.js intercepts the Promise resolution — Angular's change detection runs. This is how Angular knows when to re-render without you explicitly telling it. The downside is that Zone.js has no way to know which async event actually changed data, so it re-runs change detection for every async event, even ones irrelevant to your UI. This is why NgZone.runOutsideAngular() is useful for high-frequency events — it temporarily steps outside the Zone so change detection isn't triggered."*

---

## Q: "Explain Angular's change detection algorithm. How does it traverse the component tree?"

> *"Angular's change detection runs top-down from the root component. For every component in the tree, it evaluates each template binding expression — interpolations, property bindings, and pipe calls — and compares the new value to the previously stored value. If they differ, Angular updates the DOM. The direction is strictly top-to-bottom, child components are checked after their parents, and the algorithm runs in a single pass. With Default strategy, every component in the entire tree is checked on every run. With OnPush, Angular skips a component subtree entirely unless one of four triggers is detected. A component with OnPush is like telling Angular 'skip my subtree unless I wave at you.' The massive performance gain comes from Angular not evaluating template expressions in components that haven't received new inputs or fired internal events."*

---

## Q: "What causes ExpressionChangedAfterItHasBeenCheckedError and how do you fix it?"

> *"In development mode, Angular runs change detection twice — once to update the DOM, and once to verify that the first pass didn't cause any further changes. If a value in the template changed between the first and second pass, Angular throws this error because it means your template has an instability — updating it caused it to need updating again. The most common cause is setting a component property inside ngAfterViewInit that the component's own template binds to. ngAfterViewInit runs after the first CD pass, you set the value, Angular's verification pass finds it different from what it just rendered. The fix is to defer the assignment: wrap it in setTimeout(), which defers to the next CD cycle. Alternatively, call ChangeDetectorRef.detectChanges() immediately after setting the value, which forces a fresh CD pass that Angular is then fine with."*

---

## Q: "How does Angular's Dependency Injection resolve tokens? Walk through the injector tree."

> *"When Angular needs to resolve a dependency, it follows a hierarchical injector tree. The hierarchy has roughly three levels: the element injector level for each component and directive, the module injector level for each NgModule, and the root injector which is the app-wide singleton level. Angular walks up from the requesting component's injector through parent component injectors, then module injectors, and finally the root injector. The first provider found wins. If no provider is found anywhere up the tree, NullInjectorError is thrown at runtime. The practical implications: a service with providedIn:'root' is always found at the top and is a singleton. A service provided in a component's providers array creates a new instance for that component and shadows any root-level provider of the same type — child components of this component also get the local instance."*

---

## Q: "What is metadata reflection (reflect-metadata) and why does Angular need it?"

> *"TypeScript decorators like @Component and @Injectable add metadata to classes at declaration time. The reflect-metadata polyfill extends this so Angular can read that metadata at runtime using the Reflect API. Most critically: when Angular's DI sees a class constructor, it needs to know what types to inject for each parameter. With emitDecoratorMetadata enabled in tsconfig, TypeScript emits the types of constructor parameters as metadata — Reflect.metadata('design:paramtypes', [ProductService, Router]) on the class. Angular's injector reads this at runtime to know which tokens to resolve without you having to manually annotate each parameter. This is why the Angular compiler options require both experimentalDecorators and emitDecoratorMetadata. In modern Angular with the Ivy compiler and standalone components, some of this dependency on runtime reflection is being reduced in favor of static analysis."*
