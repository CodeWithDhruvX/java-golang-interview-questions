## 🟢 Custom Utilities, State Modeling & Team Workflow (Questions 201-250)

### Question 201: How can you create a custom `ActionGroup` in NgRx?

**Answer:**
Use the `createActionGroup` API.
```typescript
export const UserActions = createActionGroup({
  source: 'User API',
  events: {
    'Load': emptyProps(),
    'Load Success': props<{ users: User[] }>(),
    'Load Failure': props<{ error: string }>()
  }
});
```

---

### Question 202: How do you implement action deduplication?

**Answer:**
Use middleware (Metareducer) or Effects.
In Effect: `exhaustMap` ignores duplicates while active.
In Metareducer: Check if `action.type` matches previous action and timestamp is < threshold.

---

### Question 203: How can you define a custom logger for state transitions?

**Answer:**
Write a meta-reducer.
`console.group(action.type); console.log(prev); console.log(next); console.groupEnd();`.
Register it in `metaReducers` array in `StoreModule`.

---

### Question 204: How do you attach metadata to actions?

**Answer:**
Add generic props alongside payload.
`props<{ data: any, meta: { traceId: string } }>()`.
Or use a custom Action Creator class that enforces a `meta` property.

---

### Question 205: How can you implement conditional dispatching of actions?

**Answer:**
In an Effect:
`map(action => condition ? ActionA() : ActionB())`.
Or `filter(condition)` to stop dispatching.

---

### Question 206: How do you track time between actions using middleware?

**Answer:**
Store `timestamp` in state or a closure variable in the meta-reducer.
`const diff = Date.now() - lastTime`.
Log it if it exceeds a budget (Performance monitoring).

---

### Question 207: How do you make a reusable CRUD store using NgRx patterns?

**Answer:**
Create a Generic factory function.
`createCrudStore<T>(featureName: string)`.
It returns `{ actions, reducer, selectors, effects }` dynamically generated for type `T`.

---

### Question 208: How do you create reusable reducer logic for shared features?

**Answer:**
Higher-Order Reducers.
`function createPaginationReducer(reducer) { return (state, action) => ... }`.
Wrap your entity reducers with this to add pagination state (`page`, `limit`) to all of them automatically.

---

### Question 209: What is an “effect helper” and how can it reduce boilerplate?

**Answer:**
A custom utility function.
`createApiEffect(actions$, triggerAction, serviceCall, successAction)`.
It encapsulates the `ofType -> switchMap -> map -> catchError` pattern so you just pass the variables.

---

### Question 210: How do you combine multiple reducers into one?

**Answer:**
`combineReducers(map)`.
Used internally by `StoreModule`.
You can use it manually to compose nested states within a feature reducer.

---

### Question 211: How do you model deeply nested state structures?

**Answer:**
**Don't.** Normalize it.
If API returns `{ user: { posts: [ { comments: [] } ] } }`.
Store: `{ users: {}, posts: {}, comments: {} }`.
Use selectors to reconstruct the tree for the UI.

---

### Question 212: How do you normalize denormalized APIs using NgRx?

**Answer:**
Use the `normalizr` library in the Effect or Service *before* data hits the reducer.
Dispatch `LoadSuccess({ entities: normalizedData })`.

---

### Question 213: What is denormalization and when is it appropriate in NgRx?

**Answer:**
Recombining state (Joins) for the UI.
Done in **Selectors**.
`createSelector(selectUser, selectPosts, (u, p) => ...)`
State remains normalized; View Model is denormalized.

---

### Question 214: How do you track the state of dynamic forms?

**Answer:**
Store the form configurations (schema) and values in NgRx.
`forms: { [formId]: { fields: [], values: {}, valid: boolean } }`.
Update via generic `UpdateFormField` action.

---

### Question 215: How do you manage dependent state (e.g., city list based on selected country)?

**Answer:**
1.  State: `selectedCountryId`, `cities: []`.
2.  Action: `SelectCountry`.
3.  Effect: Listens to `SelectCountry`, fetches Cities, dispatches `LoadCitiesSuccess`.
4.  Reducer: unrelated to cities logic.

---

### Question 216: How do you handle feature flags in NgRx?

**Answer:**
Load flags at app startup (AppGuard).
Store in `config` slice.
Selectors `selectIsFeatureEnabled('beta')`.
Guards/Components use selector to hide/show routes.

---

### Question 217: How do you model loading and error state per API request?

**Answer:**
`callState` pattern.
`state: { loading: boolean, error: string | null }`.
Or separate `LoadingState` slice tracking request IDs: `{ [reqId]: 'loading' | 'success' }`.

---

### Question 218: How do you represent dirty vs pristine flags in state?

**Answer:**
For forms/editors:
Compare `currentValue` vs `initialValue`.
Selector: `isDirty = createSelector(current, initial, (c, i) => c !== i)`.
Don't necessarily store a boolean `isDirty`. Derive it.

---

### Question 219: How do you manage temporary vs persisted state?

**Answer:**
**Temporary:** Keep in Component (or ComponentStore). (e.g., is dropdown open).
**Persisted:** Keep in Global Store. (e.g., User Profile).

---

### Question 220: How do you separate UI state from domain state?

**Answer:**
**Domain:** `users`, `products` (Mirrors DB).
**UI:** `layout`, `activeTab`, `filters`.
Keep them in separate slices. UI state often refers to Domain state by ID.

---

### Question 221: How do you split NgRx code responsibilities in a team?

**Answer:**
Vertical Slicing (Feature Modules).
Team A owns `Checkout` feature (and its store).
Team B owns `Profile` feature.
Shared kernel module for common Actions/Selectors.

---

### Question 222: How do you enforce coding standards for NgRx?

**Answer:**
Use `eslint-plugin-ngrx`.
Rules: "Force creator functions", "No reducer mutation", "Good interaction hygiene".

---

### Question 223: How can NgRx help in large-scale team development?

**Answer:**
It forces specific contracts (Actions).
Backend team and Frontend team agree on Action payloads.
Devs can work on Reducers (Data) and Components (View) independently.

---

### Question 224: How do you avoid merge conflicts in NgRx reducer files?

**Answer:**
Split large reducers into sub-reducers or files.
Don't put all cases in one giant `createReducer`.
Use `combineReducers` early.

---

### Question 225: How do you document NgRx store structure for team members?

**Answer:**
Typescript Interfaces are the best documentation.
Add JSDoc comments to Selectors and State interfaces.
Generate state diagram tools.

---

### Question 226: How do you onboard new devs into a complex NgRx codebase?

**Answer:**
Start with "View-Model" selectors (Read-only).
Then move to Actions/Reducers (State changes).
Explain the "Circle of Life" (Action -> Effect -> Reducer -> Selector).

---

### Question 227: What should be reviewed in NgRx pull requests?

**Answer:**
1.  **Immutability:** Are they mutating state?
2.  **Selectors:** Are they memoized correctly?
3.  **Effects:** Do they handle Error cases? Do they avoid race conditions?
4.  **Boilerplate:** Is it necessary?

---

### Question 228: How do you manage multiple developers working on the same feature store?

**Answer:**
Feature slicing.
One dev works on `Products List` logic.
Another on `Product Detail` logic.
Even if in same "Products" feature, they can use separate files/reducers.

---

### Question 229: How do you create reusable effect patterns for team productivity?

**Answer:**
Abstract the HTTP handling.
Create a custom operator `safeApiCall` that handles the catchError/dispatch pattern automatically.

---

### Question 230: How do you organize your NgRx code for scalability in teams?

**Answer:**
**Nx Monorepo.**
Libs: `feature-cart`, `data-access-cart`, `ui-cart`.
Store logic lives strictly in `data-access` libraries.

---

### Question 231: How do you integrate NgRx with Angular Signals API?

**Answer:**
`selectSignal` (Read).
`patchState` (if using SignalStore).
Signals consume Store data. Store is the "Write" layer (Actions).

---

### Question 232: How do you sync NgRx with third-party libraries like Firebase?

**Answer:**
Effects.
At startup, Effect subscribes to Firebase Observable.
Dispatches `UpsertAction` whenever Firebase emits.
Effect keeps Store in sync with Realtime DB.

---

### Question 233: Can NgRx work with RxJS-based WebSocket libraries?

**Answer:**
Yes, very well.
Store = Client source of truth.
WS = Async source of data (Effect).
Component never touches WS directly.

---

### Question 234: How do you use NgRx with GraphQL?

**Answer:**
Effect calls `apollo.watchQuery()`.
Stream results to Actions.
Apollo has its own cache, so you must decide: Use Apollo Cache directly (bypass NgRx) OR sync to NgRx (Redundant but consistent).

---

### Question 235: How do you integrate NgRx with ngrx-query or ngrx-rtk-query style libraries?

**Answer:**
These libraries (like `Angular Query`) manage server state (caching/refetching) automatically.
Use NgRx Store only for *Client State* (UI toggles, filters) effectively hybrid approach.

---

### Question 236: How do you use NgRx alongside Akita or Apollo?

**Answer:**
Separate concerns.
Apollo = Data.
NgRx = UI State / Complex inter-component communication.
Don't duplicate data.

---

### Question 237: How do you combine Angular Service Workers with NgRx for caching?

**Answer:**
SW caches Network requests.
NgRx caches Memory state.
SW works at HTTP layer (offline support). NgRx works at App layer.
They complement each other.

---

### Question 238: How do you use NgRx in a project using Nx monorepo?

**Answer:**
Generate Access libs: `nx g @nrwl/angular:lib data-access-users`.
Add NgRx feature state there.
Import this lib into any App that needs User state.

---

### Question 239: Can you combine NgRx and Zustand or other frontend state tools?

**Answer:**
Technically yes, but why?
Inconsistent patterns confuse teams. Stick to one global state manager.
Maybe use lighter tools (Signals) for local state and NgRx for global.

---

### Question 240: How do you handle notifications with RxJS toast libraries?

**Answer:**
Effect listens to `[App] Failure`.
Calls `toastService.error()`.
Don't store the Toast object in NgRx state (execution, not data).

---

### Question 241: How do you dynamically register reducers for lazy-loaded routes?

**Answer:**
`StoreModule.forFeature`.
Angular Router loads the module.
Module constructor (or import) executes `forFeature`, registering the reducer.

---

### Question 242: How do you conditionally load effects for A/B testing?

**Answer:**
Provide different Effect classes based on Environment or Flags.
`providers: [ { provide: USER_EFFECTS, useClass: IsBeta ? BetaEffects : V1Effects } ]`.

---

### Question 243: What are challenges with lazy loading state and how do you handle them?

**Answer:**
**Challenge:** Selectors might fail if state isn't initialized yet.
**Fix:** Feature Detectors. Ensure selectors handle `undefined` state gracefully or use Route Guards to ensure module loaded.

---

### Question 244: How do you clean up feature store state after navigating away?

**Answer:**
Dispatch `ResetFeature` action in `ngOnDestroy`.
Or use a Meta-Reducer that cleans up keys when route changes.

---

### Question 245: Can you defer loading of state logic until the user interacts?

**Answer:**
Yes. Don't use `forFeature` in Module imports.
Inject `Store` and `ReducerManager` manually in the Component's `click` handler and add reducer then (Advanced/Hack).

---

### Question 246: How do you preload feature states for fast navigation?

**Answer:**
Dispatch a `LoadData` action *before* navigation (e.g., on hover of the link).
The Effect starts fetching.
When user clicks, data is already arriving.

---

### Question 247: What happens to effects when the feature module is destroyed?

**Answer:**
They generally stay alive (unless carefully managed).
Effects are usually Singletons registered in Root.
Lazy Modules aren't really "destroyed" completely by Angular Router (they stay in memory).
To stop effects, you must manually unsubscribe/kill stream, typically using `takeUntil` logic matched to navigation.

---

### Question 248: How do you separate shared state vs lazy state?

**Answer:**
**Shared:** `CoreModule` imports `StoreModule.forFeature('auth')`.
**Lazy:** `LazyModule` imports `StoreModule.forFeature('lazy')`.

---

### Question 249: How do you handle dynamic actions from plugin modules?

**Answer:**
Action stream is global.
Plugins can dispatch generic actions `PluginAction({ pluginId, type })`.
Main reducer handles generic action.

---

### Question 250: How do you structure routes and state together?

**Answer:**
Sync them.
Url is source of truth for `id`, `page`, `filter`.
Effect listens to Route Change -> Updates State.
State -> UI.
Don't duplicate url params in state unless necessary for caching.
