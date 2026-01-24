## 🟢 Basics, Selectors & Effects (Questions 1-50)

### Question 1: What is NgRx and why is it used in Angular applications?

**Answer:**
NgRx is a reactive state management library for Angular, inspired by Redux. It provides a single source of truth (Store) for your application state.
**Why used:**
*   **Predictability:** State flows in one direction (Unidirectional Data Flow).
*   **Debuggability:** Powerful DevTools (Time-travel).
*   **Consistency:** Standardized way to handle side effects (Effects).

**Code:**
```typescript
StoreModule.forRoot({ count: counterReducer })
```

---

### Question 2: How does NgRx relate to Redux?

**Answer:**
NgRx is "Redux for Angular" powered by **RxJS**.
It uses the Redux pattern (Store, Actions, Reducers) but implements them using RxJS Observables (`select()` returns an Observable).

---

### Question 3: What is a Store in NgRx?

**Answer:**
The Store is a centralized, immutable database for your client-side application. It holds the state tree. Components `dispatch` actions to it and `select` data from it.

---

### Question 4: How do you define the state in NgRx?

**Answer:**
By defining an interface for the shape of the data.
```typescript
export interface AppState {
  users: User[];
  loading: boolean;
}
```

---

### Question 5: What are actions in NgRx?

**Answer:**
Unique events that describe *something that happened* in the application. They are plain objects with a `type` property and an optional `payload` (props).

**Code:**
```typescript
export const loadUsers = createAction('[User List] Load Users');
```

---

### Question 6: How do you dispatch an action?

**Answer:**
Inject the `Store` service and call `.dispatch()`.
```typescript
constructor(private store: Store) {}
this.store.dispatch(increment());
```

---

### Question 7: What is a reducer in NgRx?

**Answer:**
A pure function that takes the **current state** and an **action**, and returns a **new state**. It handles state transitions.

**Code:**
```typescript
export const counterReducer = createReducer(
  initialState,
  on(increment, (state) => state + 1)
);
```

---

### Question 8: How does a reducer update the state?

**Answer:**
Immutably. It *never* modifies the existing state object. It returns a **copy** of the state with the changes applied (using spread operator `...` or libraries like Immer).

---

### Question 9: What is the purpose of an initial state?

**Answer:**
It defines the default values of the state store before any actions are dispatched. Ensures the UI has data to render initially.

---

### Question 10: What is the difference between `createAction` and `Action` class?

**Answer:**
*   **`createAction` (Modern):** Functional API. Less boilerplate. Returns a function that creates the action object.
*   **`Action` Class (Legacy):** Requires defining a class with a `readonly type`. Verbose.

---

### Question 11: What is a selector in NgRx?

**Answer:**
A pure function used to obtain a slice of store state.
It wraps the state selection logic, allowing components to get data without knowing the state structure.

---

### Question 12: Why should you use selectors instead of directly accessing the store?

**Answer:**
*   **Decoupling:** Components don't need to know the state shape.
*   **Memoization:** Selectors are cached. If the input state hasn't changed, the calculation is skipped (Performance).
*   **Reusability:** Write once, use everywhere.

---

### Question 13: How do you create a selector using `createSelector`?

**Answer:**
Combine one or more input selectors and a projector function.
```typescript
export const selectCount = (state: AppState) => state.count;
export const selectDoubleCount = createSelector(
  selectCount,
  (count) => count * 2
);
```

---

### Question 14: What is the difference between `createFeatureSelector` and `createSelector`?

**Answer:**
*   **`createFeatureSelector`:** Selects a top-level feature slice (e.g., `products` from `AppState`).
*   **`createSelector`:** Derives/Computes data from other selectors.

---

### Question 15: How do you select nested state properties?

**Answer:**
Chain selectors.
1. `selectFeature` gets the feature object.
2. `createSelector(selectFeature, state => state.nested)` drills down.

---

### Question 16: Can selectors be reused across components? How?

**Answer:**
Yes. Selectors are just exported constants/functions. Import them in any component and pass to `store.select()`.

---

### Question 17: What is memoization in the context of selectors?

**Answer:**
NgRx selectors remember the last calculation arguments. If the `state` argument is the same as last time, it returns the *cached* result immediately without re-running the projection function.

---

### Question 18: How do selectors help in improving performance?

**Answer:**
By preventing unnecessary recalculations (e.g., filtering a list of 10,000 items) and preventing unnecessary emissions (if result object reference is same, Angular change detection might skip).

---

### Question 19: What are effects in NgRx?

**Answer:**
Side-effects model.
They listen for dispatched actions, perform async tasks (HTTP), and dispatch **new** actions (Success/Failure) based on the result.

---

### Question 20: When should you use effects?

**Answer:**
For any logic that interacts with the "outside world":
*   API Calls (HTTP).
*   Reading/Writing to LocalStorage.
*   WebSockets.
*   Router navigation.

---

### Question 21: How do you handle side effects like API calls in NgRx?

**Answer:**
Inside an Effect:
1. Listen `ofType(loadAction)`.
2. Use `mergeMap/switchMap` to call Service.
3. Map result to `successAction`.
4. Catch error to `failureAction`.

**Code:**
```typescript
load$ = createEffect(() => this.actions$.pipe(
  ofType(PageActions.load),
  mergeMap(() => this.service.getAll().pipe(
    map(data => PageActions.loadSuccess({ data }))
  ))
));
```

---

### Question 22: What is the role of `Actions` observable in effects?

**Answer:**
`Actions` is a stream of *every* action dispatched to the store. Effects subscribe to this stream to react to specific events.

---

### Question 23: How do you cancel an ongoing effect?

**Answer:**
Use the `switchMap` operator.
If a new action arrives while the previous API call is pending, `switchMap` unsubscribes (cancels) the previous inner observable.

---

### Question 24: What is the use of `switchMap`, `mergeMap`, `exhaustMap`, and `concatMap` in effects?

**Answer:**
*   **`switchMap`:** Search (Cancel previous, run latest).
*   **`mergeMap`:** Delete (Run all in parallel).
*   **`concatMap`:** Save/Order (Run sequentially 1 by 1).
*   **`exhaustMap`:** Login (Ignore new clicks until current finishes).

---

### Question 25: How can you test an NgRx effect?

**Answer:**
Use `provideMockActions` to simulate the action stream.
Assert that the output Observable emits the expected Action.
Use `jasmine-marbles` for precise timing tests.

---

### Question 26: What is `ofType()` and how is it used?

**Answer:**
A custom NgRx operator used in Effects to filter the stream of all actions down to just the specific one(s) you care about.
`this.actions$.pipe(ofType(UserActions.login))`

---

### Question 27: What is NgRx Entity?

**Answer:**
A package (`@ngrx/entity`) that provides a standardized way to manage collections of records (Entities) in the state. Optimized for performance using IDs.

---

### Question 28: How does NgRx Entity simplify state management for collections?

**Answer:**
It automates boilerplate reducers (add, update, remove).
It stores data in a normalized format: `{ ids: [], entities: {} }` (Dictionary), making lookups O(1) instead of O(n) array scans.

---

### Question 29: What is an Entity Adapter?

**Answer:**
The core utility of NgRx Entity.
`createEntityAdapter<User>()`.
It provides the methods (`addOne`, `updateMany`, `getSelectors`) to manipulate the entity state.

---

### Question 30: How do you define and initialize entity state?

**Answer:**
```typescript
export interface State extends EntityState<User> {
  selectedId: number | null;
}
export const adapter = createEntityAdapter<User>();
export const initialState = adapter.getInitialState({ selectedId: null });
```

---

### Question 31: How do you perform CRUD operations using NgRx Entity?

**Answer:**
Use adapter methods inside the reducer:
*   `adapter.addOne(user, state)`
*   `adapter.updateOne(update, state)`
*   `adapter.removeOne(id, state)`

---

### Question 32: What are the benefits of using NgRx Entity?

**Answer:**
1.  **Normalization:** No duplicates.
2.  **Performance:** Fast lookup by ID.
3.  **Boilerplate:** Generates standard reducers/selectors for you.

---

### Question 33: How do you select all entities from the store?

**Answer:**
Use the default selector `selectAll`.
```typescript
const { selectAll } = adapter.getSelectors();
export const selectAllUsers = createSelector(selectUserState, selectAll);
```

---

### Question 34: How do you select a specific entity by ID?

**Answer:**
Since state is a Dictionary, just access properties.
`entities[id]`.
Or create a selector with props:
`createSelector(selectEntities, (entities, props) => entities[props.id])`.

---

### Question 35: How do you set up StoreModule in the root of your Angular app?

**Answer:**
Import it in `AppModule`.
```typescript
imports: [
  StoreModule.forRoot({ check: checkReducer }, { metaReducers })
]
```

---

### Question 36: How do you configure a feature store?

**Answer:**
In a lazy-loaded module:
`StoreModule.forFeature('products', productsReducer)`
This adds the `products` slice to the global state only when the module loads.

---

### Question 37: What is the role of `StoreModule.forFeature()`?

**Answer:**
It dynamically injects a reducer into the global store at runtime. Essential for lazy loading to keep the main bundle small.

---

### Question 38: How do you configure EffectsModule in your app?

**Answer:**
`AppModule`: `EffectsModule.forRoot([AppEffects])`
`FeatureModule`: `EffectsModule.forFeature([ProductEffects])`

---

### Question 39: Can you lazy-load feature states and effects?

**Answer:**
Yes. Using `forFeature` in the lazy module ensuring state/logic is downloaded only when the user visits that route.

---

### Question 40: What are some best practices for structuring an NgRx store?

**Answer:**
*   **Feature-Sliced:** Group by domain (User, Order) not type (Reducers, Actions).
*   **Action Hygiene:** specific action names `[Product Page] Load` vs generic `Load`.
*   **Normalization:** Keep state flat.

---

### Question 41: How do you organize files in an NgRx project?

**Answer:**
Common pattern:
`store/actions`, `store/reducers`, `store/effects`, `store/selectors`.
Or co-located with the feature: `features/users/state/{actions,reducer}.ts`.

---

### Question 42: How do you handle error states in NgRx?

**Answer:**
Store the error in the state.
`state = { data: null, error: '404 Not Found' }`.
Effect catches error -> dispatches `Failure({ error })`. Reducer saves it. Component selects it to show toast.

---

### Question 43: How do you handle loading indicators using NgRx?

**Answer:**
Add a `loading: boolean` flag to the state.
*   Action `Load`: `loading = true`.
*   Action `Success/Failure`: `loading = false`.
Ui selects `pageLoading$` to show spinner.

---

### Question 44: How can you persist NgRx store data across sessions?

**Answer:**
Use **Meta-Reducers**.
`ngrx-store-localstorage`.
Syncs specific slices of state to `localStorage` on every change and rehydrates on app start.

---

### Question 45: How do you handle optimistic updates in NgRx?

**Answer:**
1.  Dispatch Action (e.g., `Upvote`).
2.  Reducer updates state *immediately* (UI updates).
3.  Effect calls API.
4.  If API fails -> Dispatch `UndoUpvote` (Reducer reverts state).

---

### Question 46: What are good naming conventions for actions and state slices?

**Answer:**
**Actions:** `[Source] Event`.
`[Auth API] Login Success`.
`[Login Page] Login Clicked`.
**Slice:** CamelCase noun (`userConfig`, `shoppingCart`).

---

### Question 47: Should every feature have its own state module?

**Answer:**
Ideally yes, if it has data.
If a feature only presents data from another parent feature, maybe not.
But `forFeature` isolation keeps the app scalable.

---

### Question 48: What is the NgRx Store DevTools?

**Answer:**
A browser extension bridge. It allows you to inspect the state tree, see the log of actions, and "Time Travel" (replay/rollback actions).

---

### Question 49: How do you integrate DevTools with your app?

**Answer:**
`npm install @ngrx/store-devtools`.
Import `StoreDevtoolsModule.instrument({ maxAge: 25 })` in `AppModule`.

---

### Question 50: What are the benefits of using the NgRx DevTools extension?

**Answer:**
*   Visualizing state changes.
*   Debugging race conditions (order of actions).
*   Exporting state to a file to reproduce bugs on another machine.
