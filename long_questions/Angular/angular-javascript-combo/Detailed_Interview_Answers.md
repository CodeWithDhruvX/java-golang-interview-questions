# Detailed Angular & JavaScript Interview Answers

## 🔥 JavaScript Question Bank (Must Prepare)

### 🧠 Core Fundamentals

1.  **What is the difference between var, let, and const?**
    *   `var`: Function scoped, hoisted (initialized as undefined), can be redeclared.
    *   `let`: Block scoped, hoisted (in Temporal Dead Zone), cannot be redeclared in same scope.
    *   `const`: Block scoped, must be assigned immediately, cannot be reassigned (but objects are mutable).

2.  **What is hoisting?**
    *   The behavior where variable and function declarations are moved to the top of their scope during the compilation phase. `var` is hoisted and initialized with `undefined`. Function declarations are fully hoisted. `let` and `const` are hoisted but remain uninitialized (TDZ).

3.  **What is temporal dead zone (TDZ)?**
    *   The period between the start of a block and the point where a variable (declared with `let` or `const`) is declared. Accessing the variable during this time throws a ReferenceError.

4.  **What is closure? Give a real use case.**
    *   A function bundled with its lexical environment, allowing it to access variables from its outer scope even after the outer function has finished executing.
    *   **Use case:** Data encapsulation (private variables), currying, event handlers maintaining state.

5.  **What is event loop?**
    *   A mechanism that handles asynchronous callbacks in JavaScript. It continuously checks the Call Stack and the Callback Queue. If the Call Stack is empty, it pushes the first task from the Queue to the Stack.

6.  **What is call stack?**
    *   A LIFO (Last In, First Out) data structure that tracks function calls. When a function is invoked, it's pushed onto the stack; when it returns, it's popped off.

7.  **What are microtasks and macrotasks?**
    *   **Microtasks**: High priority (Promises, `queueMicrotask`, MutationObserver). Executed immediately after the current script/task finishes.
    *   **Macrotasks**: Lower priority (setTimeout, setInterval, I/O). Executed after microtasks are cleared.

8.  **Difference between == and ===?**
    *   `==` (Loose equality): Performs type coercion before comparison (e.g., `5 == '5'` is true).
    *   `===` (Strict equality): Checks both value and type (e.g., `5 === '5'` is false).

9.  **What is type coercion?**
    *   The automatic conversion of values from one data type to another (e.g., in `1 + "2"`, 1 is coerced to string "1").

10. **What are primitive and non-primitive types?**
    *   **Primitive**: Immutable, stored by value (String, Number, Boolean, Null, Undefined, Symbol, BigInt).
    *   **Non-primitive**: Mutable, stored by reference (Object, Array, Function).

11. **What is undefined vs null?**
    *   `undefined`: A variable has been declared but not assigned a value.
    *   `null`: An assignment value representing "no value" or "empty object".

12. **What is NaN?**
    *   "Not-a-Number". Result of an invalid mathematical operation (e.g., `Math.sqrt(-1)`, `'abc' * 10`). `NaN === NaN` is false (use `Number.isNaN()`).

13. **What is truthy and falsy?**
    *   Values that resolve to `true` or `false` in a boolean context.
    *   **Falsy**: `false`, `0`, `""`, `null`, `undefined`, `NaN`.
    *   **Truthy**: Everything else (including `[]`, `{}`).

14. **What is a higher-order function?**
    *   A function that takes another function as an argument (callback) or returns a function. Example: `map`, `filter`, `reduce`.

15. **What is callback hell?**
    *   Deeply nested callbacks making code hard to read and debug.
    *   **Fix**: Use Promises or Async/Await.

### ⚡ Async JavaScript

16. **What is a Promise?**
    *   An object representing the eventual completion (or failure) of an asynchronous operation.

17. **States of Promise?**
    *   Pending, Fulfilled, Rejected.

18. **What is async/await?**
    *   Syntactic sugar over Promises that allows writing asynchronous code in a synchronous manner.

19. **Difference between Promise and async/await?**
    *   `async/await` is more readable and allows using `try/catch` for error handling, avoiding `.then().catch()` chains.

20. **What happens if you don't use await?**
    *   The function will return a pending Promise immediately, and execution continues to the next line without waiting for the async operation to complete.

21. **How does setTimeout work internally?**
    *   It registers a callback with the Web API (timer). When the time expires, the callback is moved to the Macrotask Queue (Callback Queue) and waits for the Event Loop to push it to the Call Stack.

22. **What is fetch API?**
    *   A modern, Promise-based interface for making HTTP requests (replacement for XMLHttpRequest).

23. **How do you handle API errors?**
    *   Using `.catch()` with Promises or `try/catch` blocks with `async/await`. Also checking `response.ok` in fetch.

24. **What is Promise.all?**
    *   Takes an array of promises and resolves when *all* are fulfilled. Rejects immediately if *any* promise rejects.

25. **What is Promise.race?**
    *   Takes an array of promises and settles (resolves or rejects) as soon as the *first* promise settles.

### 🔥 this & Functions

26. **What is `this` in global scope?**
    *   In browser: `window`. In Node.js: `global`. In strict mode: `undefined`.

27. **What is `this` inside object method?**
    *   It refers to the object itself (the owner of the method).

28. **What is `this` in arrow function?**
    *   Arrow functions do not have their own `this`. They inherit `this` from the surrounding (lexical) scope.

29. **What is bind, call, apply?**
    *   **Call**: Invokes function immediately with specific `this` and arguments (comma-separated).
    *   **Apply**: Same as Call, but arguments are passed as an array.
    *   **Bind**: Returns a new function with `this` bound to a specific value (does not invoke immediately).

30. **What is function currying?**
    *   Transforming a function `f(a, b, c)` into `f(a)(b)(c)`. Useful for partial application.

31. **What are IIFE?**
    *   Immediately Invoked Function Expressions. Functions that run as soon as they are defined. Used to avoid polluting global scope. `(function(){ ... })();`

32. **What is prototype?**
    *   A mechanism by which JavaScript objects inherit features from one another. Every object has a prototype property.

33. **What is prototypal inheritance?**
    *   The method by which an object accesses properties and methods of another object via the prototype chain.

### 📦 Arrays & Objects

34. **Difference between map and forEach?**
    *   `map`: Returns a new array. Chainable.
    *   `forEach`: Iterates only (returns undefined). Used for side effects.

35. **What does reduce do?**
    *   Reduces an array to a single value (accumulator) by executing a reducer function on each element.

36. **How do you remove duplicates from array?**
    *   `[...new Set(array)]`
    *   Using `filter` + `indexOf`
    *   Using `reduce`

37. **How do you deep clone an object?**
    *   `JSON.parse(JSON.stringify(obj))` (basic)
    *   `structuredClone(obj)` (modern)
    *   Recursive function or lodash `_.cloneDeep`.

38. **Difference between shallow copy and deep copy?**
    *   **Shallow**: Copies only the first level. Nested objects are referenced.
    *   **Deep**: Recursively copies all levels. Components are independent.

39. **What is destructuring?**
    *   Syntax to unpack values from arrays or properties from objects into distinct variables. `const { name } = user;`

40. **What is spread operator?**
    *   (`...`) Expands an iterable into individual elements. Used for copying arrays, merging objects/arrays.

41. **What is rest parameter?**
    *   (`...args`) Collects all remaining arguments into an array.

42. **How to merge two arrays?**
    *   `[...arr1, ...arr2]` or `arr1.concat(arr2)`.

43. **How to check if object is empty?**
    *   `Object.keys(obj).length === 0`.

44. **How to iterate object properties?**
    *   `for...in` loop.
    *   `Object.keys()`, `Object.values()`, `Object.entries()`.

### 🚀 Scenario-Based JS

45. **How do you debounce an API call?**
    *   Wrap the call in a function that clears the previous timer and sets a new one. The API executes only after the user stops triggering it for `n` milliseconds.

46. **How do you throttle a scroll event?**
    *   Ensure the function executes at most once every `n` milliseconds, regardless of how often the event fires.

47. **How do you prevent memory leaks in JS?**
    *   Remove event listeners.
    *   Clear intervals/timeouts.
    *   Avoid global variables.
    *   Nullify references to detached DOM nodes.

48. **What happens if two async calls return in different order?**
    *   Race condition. Fix using `Promise.all` (wait for both) or chaining, or strictly controlling state updates.

49. **How do you optimize performance in large list rendering?**
    *   **Virtualization/Windowing**: Render only the visible items (e.g., Virtual Scrolling).

50. **How do you debug JavaScript errors?**
    *   Browser DevTools (Console, Sources tab breakpoints, Network tab).
    *   `debugger` statement.
    *   Source maps for minified code.

---

## 🚀 Angular Question Bank

### 🏗 Angular Basics

1.  **What is Angular?**
    *   A platform and framework for building single-page client applications using HTML and TypeScript, developed by Google.

2.  **What is SPA (Single Page Application)?**
    *   An app that loads a single HTML page and dynamically updates that page as the user interacts with the app, without reloading.

3.  **What is TypeScript?**
    *   A superset of JavaScript that adds static typing. It compiles to JavaScript.

4.  **What is Angular CLI?**
    *   Command Line Interface to initialize, develop, scaffold, and maintain Angular applications (`ng new`, `ng serve`, `ng generate`).

5.  **Explain Angular architecture.**
    *   Modules -> Components -> Templates -> Metadata -> Data Binding -> Directives -> Services -> Dependency Injection.

6.  **What is a component?**
    *   The fundamental building block of UI. Confines a view (HTML), logic (TS), and styles (CSS).

7.  **What is a module (NgModule)?**
    *   A container for a cohesive block of code dedicated to an application domain, workflow, or set of capabilities.

8.  **What is metadata?**
    *   Data that tells Angular how to process a class (e.g., `@Component({...})`).

9.  **What is a decorator?**
    *   A function that modifies a class, property, or method. Examples: `@Component`, `@Injectable`, `@Input`.

10. **What is bootstrap array?**
    *   Defines the root component that Angular creates and inserts into the `index.html` host file.

### 🔁 Lifecycle Hooks

11. **Explain Angular lifecycle hooks.**
    *   Events in the life of a component: Creation -> Change Detection -> Destruction.
    *   `ngOnChanges`, `ngOnInit`, `ngDoCheck`, `ngAfterContentInit`, `ngAfterContentChecked`, `ngAfterViewInit`, `ngAfterViewChecked`, `ngOnDestroy`.

12. **Difference between ngOnInit and constructor?**
    *   **Constructor**: JS engine feature, for dependency injection. Inputs are not available.
    *   **ngOnInit**: Angular lifecycle hook, for initialization logic. Inputs (`@Input`) are available.

13. **When does ngOnChanges trigger?**
    *   Whenever a data-bound input property (`@Input`) changes.

14. **What is ngOnDestroy used for?**
    *   Cleanup: Unsubscribing from Observables, detaching event handlers, clearing timers to prevent memory leaks.

15. **Real use case of ngAfterViewInit?**
    *   Accessing child components (`@ViewChild`) or DOM elements after the view is fully initialized.

### 🔗 Data Binding

16. **Types of data binding?**
    *   Interpolation `{{}}`
    *   Property `[ ]`
    *   Event `( )`
    *   Two-way `[( )]`

17. **Difference between property and attribute binding?**
    *   **Attribute**: Initial value defined in HTML (doesn't change).
    *   **Property**: Current value of the DOM object (changes dynamically).

18. **What is two-way binding?**
    *   Syncs data between the model and the view. Updates in one reflect in the other immediately.

19. **How does `[(ngModel)]` work internally?**
    *   It binds a property `[ngModel]` and an event `(ngModelChange)`.

### 🧩 Directives

20. **Difference between structural and attribute directives?**
    *   **Structural**: Change DOM layout (add/remove elements). Start with `*` (e.g., `*ngIf`).
    *   **Attribute**: Change appearance/behavior of element. (e.g., `ngClass`, `ngStyle`).

21. **What is *ngIf?**
    *   Conditionally includes or excludes a template based on the value of an expression.

22. **What is *ngFor?**
    *   Repeats a node for each item in a collection.

23. **What is trackBy and why is it used?**
    *   A function used with `*ngFor` to define how to identify items. Improves performance by preventing re-rendering of unchanged items.

24. **How do you create custom directive?**
    *   Create a class decorated with `@Directive`. Inject `ElementRef` to manipulate the host element.

### 💉 Services & DI

25. **What is dependency injection (DI)?**
    *   Design pattern where dependencies (services) are provided to a class/component externally rather than created inside it.

26. **What is @Injectable?**
    *   Decorator that marks a class as a service that can be injected.

27. **What is providedIn: 'root'?**
    *   Registers the service as a singleton in the root injector. Makes it tree-shakable (unused services are removed from build).

28. **How do you share data between components?**
    *   Parent-Child: `@Input`/`@Output`.
    *   Unrelated: Shared Service (using Observables/Subjects).
    *   State Management: NgRx/Akita.

29. **Singleton service in Angular?**
    *   Services provided in `root` or a module are singletons within that scope (usually app-wide).

### 🌊 Observables & RxJS

30. **What is Observable?**
    *   A stream of data that can be synchronous or asynchronous. It can emit multiple values over time.

31. **Difference between Promise and Observable?**
    *   **Promise**: Single value, not cancellable, eager.
    *   **Observable**: Multiple values, cancellable, lazy (runs only on subscribe).

32. **What is subscription?**
    *   The execution of an Observable. Created by calling `.subscribe()`.

33. **What happens if you don't unsubscribe?**
    *   Memory leaks, especially if the Observable is long-lived (e.g., interval, DOM events).

34. **What is switchMap?**
    *   Maps to a new Observable and *cancels* the previous inner Observable. Good for typeahead/search.

35. **Difference between mergeMap and switchMap?**
    *   **mergeMap**: runs all inner Observables concurrently.
    *   **switchMap**: key is *cancellation* of previous.

36. **What is pipe?**
    *   A method to chain RxJS operators (`map`, `filter`, etc.) to transform the data stream.

37. **What is Subject?**
    *   A special Observable that is also an Observer. It can multicast to multiple observers.

38. **What is BehaviorSubject?**
    *   A Subject that requires an initial value and emits the current value to new subscribers immediately.

39. **What is ReplaySubject?**
    *   A Subject that replays a specified number of old values to new subscribers.

### 🌐 HTTP & API

40. **How do you call API in Angular?**
    *   Using `HttpClient` module.

41. **What is HttpClient?**
    *   Simplified API for HTTP requests. Returns Observables. Supports interceptors, typed responses.

42. **How do you handle errors?**
    *   Using `catchError` operator in pipe.

43. **What is interceptor?**
    *   Middleware that intercepts HTTP requests/responses (e.g., to add auth tokens or log errors) globally.

44. **How do you add token in header?**
    *   In an Interceptor, clone the request: `req.clone({ setHeaders: { Authorization: token } })`.

45. **How do you cancel previous API calls?**
    *   Using `switchMap` operator.

### 🛣 Routing

46. **What is Angular routing?**
    *   Features enabling navigation from one view to another.

47. **What is lazy loading?**
    *   Loading feature modules only when requested (navigated to). improves startup time.

48. **What is route guard?**
    *   Interfaces to control whether a user can navigate to or away from a route (authentication/authorization).

49. **Types of guards?**
    *   `CanActivate`, `CanDeactivate`, `CanLoad`, `Resolve`.

50. **What is router outlet?**
    *   `<router-outlet>`: Placeholder directive where the router displays the routed component.

### 📝 Forms

51. **Difference between template-driven and reactive forms?**
    *   **Template**: Logic in HTML, easy for simple forms, uses `ngModel`.
    *   **Reactive**: Logic in TS, explicit, scalable, testable, uses `FormControl/Group`.

52. **What is FormGroup?**
    *   Tracks the value and validity state of a group of FormControl instances.

53. **What is FormControl?**
    *   Tracks the value and validity state of an individual form control.

54. **How do you add custom validator?**
    *   Create a function that returns `ValidationErrors | null`.

55. **How do you handle dynamic forms?**
    *   Using `FormArray` to add controls dynamically.

### ⚡ Performance

56. **What is change detection?**
    *   Mechanism to sync the application state with the UI.

57. **Default vs OnPush?**
    *   **Default**: Checks every component on every event.
    *   **OnPush**: Checks only when inputs reference changes or Async pipe emits. Better performance.

58. **What is zone.js?**
    *   Library that monkeys-patches async tasks (setTimeout, promises, DOM events) to notify Angular when to run change detection.

59. **How do you improve Angular performance?**
    *   Use OnPush strategy.
    *   Lazy loading.
    *   AOT (Ahead of Time) compilation.
    *   `trackBy` in `*ngFor`.
    *   Unsubscribe observables.

60. **How do you prevent memory leaks?**
    *   Unsubscribe in `ngOnDestroy` (or use `async` pipe).
    *   Detach event listeners.

### 🔥 High-Probability Scenario Questions

1.  **How do you handle large data table?**
    *   Pagination (Server-side).
    *   Virtual Scrolling (cdk-virtual-scroll).

2.  **How do you implement search with debounce?**
    *   Use `Subject` for input. Pipe: `debounceTime(300)`, `distinctUntilChanged()`, `switchMap(apiCall)`.

3.  **How do you manage state in Angular?**
    *   Shared Services with RxJS (BehaviorSubject) for simple apps.
    *   NgRx, Akita, or NgXS for complex apps.

4.  **How do you handle authentication?**
    *   Login API -> Store Token (LocalStorage/Cookie) -> Interceptor (Add Token) -> Auth Guard (Protect Routes).

5.  **How do you protect routes?**
    *   Implement `CanActivate` guard: check if token exists/valid. If not, redirect to login.

6.  **How do you implement role-based access?**
    *   In Auth Guard, check user role from token/state. Allow/Deny access or load specific modules.

7.  **What happens if API fails?**
    *   `catchError` block. Show user friendly message (Toastr/Snackbar). Retry logic (`retry(3)`).

8.  **How do you handle loader globally?**
    *   Interceptor: `onRequest` -> Show Loader. `onResponse/Error` -> Hide Loader. Use a LoaderService to toggle state.

9.  **How do you share data between unrelated components?**
    *   Create a SharedService with a `BehaviorSubject`. Comp A calls `next(data)`. Comp B `subscribes()` to it.

10. **How do you optimize repeated API calls?**
    *   `switchMap` (cancels previous).
    *   `shareReplay(1)` (caches response).
    *   Debouncing.
