## 🟢 Advanced Usage, Testing & Real-World (Questions 51-100)

### Question 51: How can you inspect dispatched actions and state changes?

**Answer:**
Using the **Redux DevTools Extension** in Chrome/Firefox.
You can see a timeline of actions (left panel). Clicking an action shows the `action` payload and the `diff` (what changed in state).

---

### Question 52: What is action replaying in DevTools?

**Answer:**
The ability to "Time Travel".
You can move the slider back to a previous state to see how the UI looked. You can also "Import" a state log to replay a specific sequence of user interactions that caused a bug.

---

### Question 53: How do you handle router state with NgRx?

**Answer:**
Use `@ngrx/router-store`.
It binds Angular Router events to the Store.
Whenever the route changes, an action is dispatched, and the current URL/Params are saved in the store.

---

### Question 54: What is `@ngrx/router-store`?

**Answer:**
Middleware that connects the Angular Router to the NgRx Store.
It allows you to select route params (e.g., `selectRouteParam('id')`) from the store instead of injecting `ActivatedRoute` into components.

---

### Question 55: How do you implement lazy loading with NgRx?

**Answer:**
Import the **StoreFeatureModule** in the lazy-loaded module.
`StoreModule.forFeature('featureName', reducer)`.
The state for this feature is only initialized when the user navigates to a route inside that module.

---

### Question 56: How do you handle shared state across modules?

**Answer:**
Define it in the **Root Store** (`AppModule`).
Or create a "Shared" feature module imported by others (rare).
Usually, Auth/User info is root state, accessible by all feature selectors.

---

### Question 57: How can you reset the NgRx store state?

**Answer:**
Use a **Meta-Reducer**.
Listen for a specific action `[Auth] Logout`.
If that action is seen, return `undefined` (which resets reducers to initial state) or a cleared state object.

**Code:**
```typescript
export function clearState(reducer) {
  return (state, action) => {
    if (action.type === '[Auth] Logout') {
      return reducer(undefined, action);
    }
    return reducer(state, action);
  };
}
```

---

### Question 58: How do you combine multiple reducers?

**Answer:**
`combineReducers` (internal) or standard `ActionReducerMap`.
`StoreModule.forRoot({ user: userReducer, product: productReducer })`.
This creates a state tree: `{ user: ..., product: ... }`.

---

### Question 59: Can you dynamically register reducers at runtime?

**Answer:**
Yes. The `ReducerManager` service allows adding/removing reducers manually.
`store.addReducer('key', reducer)`.
Plugins/widgets use this pattern.

---

### Question 60: How do you avoid boilerplate in NgRx?

**Answer:**
1.  **NgRx Entity:** Auto-generates CRUD.
2.  **`createActionGroup`:** Groups actions (less imports).
3.  **Schematics:** Auto-generate files.
4.  **Signal Store:** (Modern) significantly less boilerplate.

---

### Question 61: How do you test a reducer?

**Answer:**
It's a pure function.
Validates: `fn(state, action) => expectedState`.
```typescript
it('should increment', () => {
  const result = counterReducer(0, increment());
  expect(result).toBe(1);
});
```

---

### Question 62: How do you test a selector?

**Answer:**
Call the `projector` function directly.
`const result = selectCount.projector({ count: 5 });`
`expect(result).toBe(5);`
Avoid mocking the whole store; just test the logic.

---

### Question 63: How do you test an effect?

**Answer:**
Use `TestScheduler` or `jasmine-marbles`.
Model the stream of actions and the expected output stream.
`actions$ = hot('-a', { a: load() });`
`expect(effects.load$).toBeObservable(cold('-b', { b: success() }));`

---

### Question 64: How do you mock the store in unit tests?

**Answer:**
Use `provideMockStore`.
It allows you to simulate the store without setting up reducers.
You can override specific selectors to return fake data.

---

### Question 65: What is `provideMockStore()` and how is it used?

**Answer:**
A test utility from `@ngrx/store/testing`.
In `TestBed.configureTestingModule`:
`providers: [provideMockStore({ initialState })]`.

---

### Question 66: How can you test component interactions with the store?

**Answer:**
Inject `MockStore`.
1.  Spy on `dispatch` to ensure component calls actions.
2.  Override selector `store.overrideSelector(selectUsers, [...])` to test how component renders data.

---

### Question 67: How do you migrate an existing Angular service to use NgRx?

**Answer:**
1.  Move the state (variable) from Service to Store Interface.
2.  Move methods that update state to Actions/Reducers.
3.  Move API calls to Effects.
4.  Replace Service Observables with `store.select()`.

---

### Question 68: How do you manage authentication state using NgRx?

**Answer:**
*   **State:** `user: User | null`, `token: string`.
*   **Actions:** Login, LoginSuccess, Logout.
*   **Effect:** Calls Auth API, saves token to LocalStorage.
*   **Guard:** Selects `isLoggedIn` to protect routes.

---

### Question 69: How do you implement a global notification system using NgRx?

**Answer:**
*   **State:** `notifications: Notification[]`.
*   **Action:** `AddNotification({ message, type })`.
*   **Component:** `AppComponent` selects notifications list and renders Toasts.
*   Any feature can dispatch `AddNotification`.

---

### Question 70: How do you cache API responses in the store?

**Answer:**
Check the store *before* dispatching the load action.
Or inside the Effect/Guard:
`withLatestFrom(store.select(selectData))`.
If data exists, return `EMPTY` (don't call API). If null, call API.

---

### Question 71: How do you handle paginated data using NgRx?

**Answer:**
State: `{ items: [], page: 1, total: 100 }`.
Action: `LoadPage({ page: 2 })`.
Effect fetches page 2.
Reducer appends or replaces items based on requirements (Infinite scroll vs Tables).

---

### Question 72: What are NgRx Signals?

**Answer:**
A standalone state management library (`@ngrx/signals`) designed for Angular Signals.
It is lighter, functional, and doesn't require RxJS streams for simple state.

---

### Question 73: How are signals different from selectors?

**Answer:**
*   **Selectors:** Return RxJS `Observable`. Need `async` pipe.
*   **Signals:** Return a value function `count()`. Reactive primitive. Easier to use in templates (no pipe needed).

---

### Question 74: Can signals replace observables in NgRx?

**Answer:**
For **synchronous state access**, yes (and they are preferred now).
For **asynchronous event streams** (Debounce, WebSocket), RxJS Observables are still more powerful.

---

### Question 75: How do you create signal-based selectors?

**Answer:**
Using `computed()`.
`const doubleCount = computed(() => store.count() * 2);`
It automatically tracks dependencies and updates when `store.count` changes.

---

### Question 76: How do signals help in reducing subscription complexity?

**Answer:**
No need for `subscribe()` or `async` pipe.
No memory leaks from forgetting to unsubscribe.
Finer-grained updates (only text node updates, not whole component check).

---

### Question 77: What is NgRx Data?

**Answer:**
An extension (`@ngrx/data`) built on top of Store/Entity/Effects.
It is a "Zero-Boilerplate" solution. You just define an entity name ('Hero'), and it creates the Actions, Reducers, Dispatchers, and Selectors automatically.

---

### Question 78: How is NgRx Data different from NgRx Store?

**Answer:**
**Store:** Low-level, explicit. You write everything.
**Data:** opinionated, convention-over-configuration. Good for standard REST interfaces.

---

### Question 79: What are the use cases for NgRx Data?

**Answer:**
Standard CRUD applications where you have many entities (Customer, Order, Product) that all follow the same pattern (GET, POST, PUT, DELETE). Saves time.

---

### Question 80: How do you configure NgRx Data in your app?

**Answer:**
Define `entityMetadata`.
Configure `EntityDataModule.forRoot(entityConfig)`.
API URL conventions must match expected pattern (or override `HttpUrlGenerator`).

---

### Question 81: How do you extend or override default entity services in NgRx Data?

**Answer:**
Inject `EntityDataService`.
Register a custom Data Service class for a specific entity if the backend API doesn't follow standard REST conventions.

---

### Question 82: How do you use NgRx with Angular signals?

**Answer:**
`selectSignal(selector)`.
New method on the `Store` to get a Signal instead of an Observable.
`<div *ngIf="userSignal() as user">`.

---

### Question 83: Can NgRx be integrated with forms?

**Answer:**
Yes. Manual binding (dispatch on change).
Or libraries like **ngrx-forms** (connects form state to store state). Rarely used now; Reactive Forms with simple bindings usually suffice.

---

### Question 84: How do you use NgRx with Angular services?

**Answer:**
The Service is the data provider. The Effect is the consumer.
Effect -> calls Service -> dispatches Action.
Services should be stateless (mostly).

---

### Question 85: How do you integrate NgRx with a WebSocket connection?

**Answer:**
Create one Effect that connects to WS.
When WS receives message -> Dispatch Action.
When Action dispatched -> Send WS message.
Keeps WS logic isolated from components.

---

### Question 86: Can you use NgRx with SSR (Angular Universal)?

**Answer:**
Yes. The state is built on the server.
**TransferState:** Serialize the final store state to JSON in the HTML.
Client bootstraps and rehydrates the store from that JSON to avoid double-fetching.

---

### Question 87: How do you migrate from older versions of NgRx?

**Answer:**
Use `ng update @ngrx/store`.
Modernize: Switch from Class-based Actions to `createAction`. Switch to `createReducer`.

---

### Question 88: What are common breaking changes in recent NgRx versions?

**Answer:**
Removal of decorators (`@Effect`).
Deprecation of `switchMap` in favor of functional effects (`createEffect`).
Typescript strictness improvements.

---

### Question 89: How do you upgrade NgRx without breaking the app?

**Answer:**
Incremental adoption.
Old reducers (switch-case) work fine alongside new `createReducer`.
Upgrade one feature module at a time.

---

### Question 90: How do you minimize boilerplate in NgRx?

**Answer:**
Use **Signal Store** (if new).
Use **NgRx Entity**.
Use `createActionGroup` (Autogenerates success/failure variants).

---

### Question 91: What is the role of `@ngrx/component-store`?

**Answer:**
A library for **local state management**.
Unlike the Global Store, ComponentStore is tied to a specific component's lifecycle.
Perfect for complex isolated widgets (Datepicker, Datagrid).

---

### Question 92: When should you use `ComponentStore` instead of full NgRx?

**Answer:**
When state is **not shared** globally.
Example: "Is the dropdown open?" or "Current filter of this specific table instance".
Lightweight, no global actions/reducers.

---

### Question 93: How do you manage state locally using `ComponentStore`?

**Answer:**
Extend `ComponentStore<State>`.
`readonly vm$ = this.select(...)`.
`readonly updateName = this.updater(...)`.
Provide in component `providers: []` array.

---

### Question 94: How do you lazy-load effects in large-scale applications?

**Answer:**
Provide them in the lazy module.
`EffectsModule.forFeature([LazyEffects])`.
Angular DI handles the rest.

---

### Question 95: How do you design state for a multi-feature app?

**Answer:**
**Root State:** Auth, Router, Config.
**Feature State:** Lazy slices (Dashboard, Admin).
Keep boundaries clean. No feature should depend on another feature's internal state (use shared libs if needed).

---

### Question 96: What is state normalization and why is it useful?

**Answer:**
Storing data like a database (Tables), not trees (Nested JSON).
`posts: { [id]: Post }`. `comments: { [id]: Comment }`.
Prevents duplication. deeply nested updates are easier (update 1 row, not scan recursive tree).

---

### Question 97: What is a meta-reducer?

**Answer:**
A "Higher Order Reducer". A function that wraps the main reducer.
It acts as middleware. It sees every action before the actual reducer.
Used for: Logging, Hydration (LocalStorage), Resetting state.

---

### Question 98: How do you implement undo/redo functionality using NgRx?

**Answer:**
**Meta-Reducer.**
Keep a history array of past states (snapshots).
On `UNDO` action, pop the last state and return it as current state.

---

### Question 99: How do you structure the store in a micro-frontend Angular app?

**Answer:**
Tricky.
Each MFE usually has its own Store.
Global Store (Shell) shares context via a shared library or Window Custom Events.
Often better to avoid 1 giant store across MFEs.

---

### Question 100: When should you **not** use NgRx?

**Answer:**
*   Simple Apps (Forms + Submission).
*   Static Content sites.
*   When team is small and unfamiliar with RxJS.
*   When `ComponentStore` or just Angular Services are sufficient.
**Over-engineering** is the main risk.
