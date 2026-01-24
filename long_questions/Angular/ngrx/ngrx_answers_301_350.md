## 🟢 State Design, Effects & Async Patterns (Questions 301-350)

### Question 301: How do you design a scalable state structure for micro-frontends using NgRx?

**Answer:**
Avoid a single global store.
Use a **Shell Store** for shared context (User, Config).
Each Micro-frontend (MFE) has its own independent Feature Store (`StoreModule.forFeature`).
Communication via Window Events or a Shared Observable Bus if needed.

---

### Question 302: How do you model API pagination using NgRx?

**Answer:**
State: `articles: { ids: [], entities: {}, pagination: { page: 1, limit: 10, total: 500 } }`.
Action: `LoadPage({ page: 2 })`.
Effect: Fetches page 2.
Reducer: Updates `pagination` and `entities` (Append or Replace based on UX).

---

### Question 303: What’s the best way to persist UI filter state across sessions?

**Answer:**
Slice: `ui: { filters: { status: 'active', sort: 'date' } }`.
Meta-reducer (`localStorageSync`) to save this specific slice to localStorage.
Rehydrate on app boot.

---

### Question 304: How do you structure state to avoid redundant data?

**Answer:**
**Normalization**.
Do not store `User` inside `Post`. Store `userId`.
Select `Post` and `User` separately and join them in the Selector.

---

### Question 305: What is the "feature slice" pattern in NgRx?

**Answer:**
Dividing the state methods into vertical domains.
Each feature (Cart, User, Products) has its own Folder, Reducer, Effects, and Selectors.
They are combined only at the Root level (or lazy loaded).

---

### Question 306: When should you break state into multiple feature modules?

**Answer:**
When the bundle size gets too large.
When team ownership splits.
When data is completely unrelated (Admin vs Public Shop).

---

### Question 307: What is colocated state and when should you avoid it?

**Answer:**
Keeping state in the Component (or ComponentStore) close to usage.
**Avoid when:** The data is needed by siblings, parent, or router guards. Move to Global Store then.

---

### Question 308: How do you structure NgRx to support multi-language applications?

**Answer:**
Store `currentLang` in state.
Action `SetLanguage`.
Effect updates `TranslateService` (ngx-translate).
Selectors might use the `currentLang` to transform data triggers if needed.

---

### Question 309: How do you handle deeply nested data structures in state?

**Answer:**
**Don't.** Flatten them.
If you must, use `immer` to simplify reducer updates.
`on(update, state => produce(state, draft => draft.level1.level2.level3 = val))`.

---

### Question 310: How would you share state between Angular elements (web components)?

**Answer:**
Since they share the window/injector:
Create a `SharedStoreService` (Singleton) provided in the platform root.
Or use Custom Events (`dispatchEvent`) to sync separate stores.

---

### Question 311: What’s the tradeoff between flat state and nested state in NgRx?

**Answer:**
**Flat:** Easier updates, better performance (shallow checks). Harder to visualize relationships.
**Nested:** Easier to visualize. Nightmare to update immutably (`...state, prop: { ...state.prop, inner: ... }`).

---

### Question 312: How do you sync state with localStorage or sessionStorage?

**Answer:**
Use `ngrx-store-localstorage`.
Config: `rehydrate: true`, `keys: ['auth', 'settings']`.
Wraps root reducer.

---

### Question 313: How do you organize cross-cutting state like auth or UI config?

**Answer:**
**Core Module.**
`StoreModule.forRoot({ auth: authReducer, ui: uiReducer })`.
Other features select from this Root state.

---

### Question 314: What is the difference between app-wide state and feature-local state?

**Answer:**
**App-wide:** Session, Theme, Router, Notifications.
**Feature-local:** Form input values, Pagination cursors, Specific Entity lists.

---

### Question 315: How do you apply the façade pattern in NgRx architecture?

**Answer:**
Create a Service `UserFacade`.
Exposes `users$ = store.select(allUsers)`.
Exposes `loadUsers() { store.dispatch(Load()) }`.
Component only talks to Facade. Hides NgRx dep.

---

### Question 316: How would you restructure an overgrown store in a legacy app?

**Answer:**
Identify "Domains".
Move related actions/reducers into a `libs/domain` library.
Use `StoreModule.forFeature` to register them dynamically.
Deprecate the giant root reducer.

---

### Question 317: How do you define state contracts across teams in a monorepo?

**Answer:**
Shared Types (`libs/api-interfaces`).
Selectors are the public API.
Actions are the public API.
Internal state shape is private.

---

### Question 318: How do you manage state version upgrades across releases?

**Answer:**
(Duplicate of Q266).
Migration strategy in hydration meta-reducer.

---

### Question 319: What is a shared store and how do you manage access control around it?

**Answer:**
A Store module imported by multiple apps.
Risky. Better to make it a Library (`data-access`) that Apps import.

---

### Question 320: How do you design an NgRx store for apps with plugin architecture?

**Answer:**
Plugins dispatch "Register Plugin" actions.
Main Store holds `plugins: Dictionary<PluginConfig>`.
Dynamic Reducers for plugin-specific state.

---

### Question 321: How do you write effects that depend on the current state?

**Answer:**
Use `concatLatestFrom` (or `withLatestFrom`).
`concatLatestFrom(() => store.select(selectUser))`.
`switchMap(([action, user]) => ...)`

---

### Question 322: How do you call another effect or chain effects programmatically?

**Answer:**
Dispatch an action.
Effect A -> Returns `Action B`.
Effect B -> Listens `Action B`.
Do not call Effect class methods directly.

---

### Question 323: How do you cancel a previous API call when a new one starts?

**Answer:**
`switchMap`.
It automatically unsubscribes the inner observable (HTTP) if a new action flows in.

---

### Question 324: How do you retry a failed effect with exponential backoff?

**Answer:**
`retryWhen(errors => errors.pipe(delay(1000), take(3)))`.
Or RxJS `retry({ count: 3, delay: 1000 })`.

---

### Question 325: What are the risks of calling services directly from effects?

**Answer:**
Tight coupling.
Harder to mock in tests (if not using DI).
Generally, it's the *correct* way (Service handles HTTP, Effect handles orchestration).

---

### Question 326: How do you implement file upload progress tracking in NgRx?

**Answer:**
Effect calls Service with `reportProgress: true`.
Service emits events (UploadProgress, Response).
Effect maps events to `UploadProgressAction({ % })`.
Reducer updates specific file status.

---

### Question 327: How do you manage effect dependencies across multiple modules?

**Answer:**
Dispatch Actions.
Module A Effect -> `SuccessAction`.
Module B Effect -> `ofType(ModuleA_SuccessAction)`.
Loose coupling via Event Bus (Actions).

---

### Question 328: How can effects react to router state or URL query params?

**Answer:**
`ofType(routerNavigatedAction)`.
`concatLatestFrom(() => store.select(selectQueryParams))`.

---

### Question 329: How do you inject runtime config into effects dynamically?

**Answer:**
Inject `AppConfigService` into Effect constructor.
Use it in the stream.

---

### Question 330: How do you debounce an effect based on user input?

**Answer:**
`debounceTime(300)`.
Good for Search Inputs.

---

### Question 331: How do you combine multiple streams inside a single effect?

**Answer:**
`combineLatest`, `forkJoin`, `zip`.
Usually inside the `switchMap`.

---

### Question 332: How do you isolate side-effects for multiple async tasks in parallel?

**Answer:**
`mergeMap`.
It subscribes to each inner observable concurrently.
(File upload 1, File upload 2 happen together).

---

### Question 333: How do you handle WebSocket or SSE updates using effects?

**Answer:**
Create one long-living Effect (no `switchMap` that completes).
`mergeMap(() => wsSubject$)`.
Dispatch actions for incoming messages.

---

### Question 334: How can you cancel an effect when a user logs out?

**Answer:**
`takeUntil(this.actions$.pipe(ofType(LogoutAction)))`.
Apply this to long-lived streams (Wait, polling).

---

### Question 335: How do you listen for completion of multiple effects before proceeding?

**Answer:**
Wait for the resulting Actions.
Effect C listens to `combineLatest([actionA$, actionB$])` (Complex RxJS).
Better: Store flags `loadedA`, `loadedB`. Effect listens to State Change `(loadedA && loadedB)`.

---

### Question 336: How do you persist the last triggered effect’s result across reloads?

**Answer:**
Save to LocalStorage (Meta-Reducer).
Effect reads from LocalStorage on Init.

---

### Question 337: How do you wrap API logic in reusable effect factories?

**Answer:**
Function that returns `createEffect`.
Pass `actions$`, `service`, `featureName`.
Returns standard CRUD effect.

---

### Question 338: How do you test long-running effects or polling logic?

**Answer:**
`fakeAsync` and `tick()`.
`TestScheduler` with virtual time.
Advance time by 5000ms and assert action dispatched.

---

### Question 339: How do you log the start and end of every effect execution?

**Answer:**
`tap(() => log('Start'))`.
`finalize(() => log('End'))`. (Inside the inner pipe).

---

### Question 340: How do you separate API orchestration logic from core effects?

**Answer:**
Move complex orchestration to a Service (`OrchestrationService`).
Effect just calls `service.doComplexThing()`.

---

### Question 341: How do you manage related entities in normalized NgRx Entity state?

**Answer:**
Foreign Keys.
`Post { id, authorId }`.
`User { id }`.
They live in separate `entity` slices.

---

### Question 342: How do you deal with polymorphic entities in NgRx?

**Answer:**
Union Types in TS.
`type Item = Video | Article`.
Store: `{ [id]: Item }`.
Runtime: Switch on `item.type` in Component.

---

### Question 343: How do you remove all entities of a type with one action?

**Answer:**
`adapter.removeAll(state)`.

---

### Question 344: How do you update multiple entities at once in NgRx?

**Answer:**
`adapter.updateMany([{ id: 1, changes: {} }, { ... }], state)`.

---

### Question 345: How do you persist selection state in entity lists?

**Answer:**
Add `selectedIds: number[]` to the State interface (alongside `ids` and `entities`).
Action `ToggleSelection(id)`.

---

### Question 346: What’s the best pattern for optimistic entity deletion?

**Answer:**
1.  `removeOne(id)` immediately.
2.  Call API.
3.  On Fail: `addOne(savedEntity)` (Restore).

---

### Question 347: How do you rehydrate entity state from server after login?

**Answer:**
`adapter.setAll(users, state)`.
Replaces current list with server list.

---

### Question 348: How do you sync the store after partial update responses?

**Answer:**
`adapter.upsertOne(user, state)`.
Updates if exists, adds if new.

---

### Question 349: How do you handle duplicates or merge conflicts in entities?

**Answer:**
`upsertMany`.
Last writer wins (usually).
Or custom logic in reducer to merge fields.

---

### Question 350: How do you create a paginated view from entity selectors?

**Answer:**
Select All.
Selector: `(all, page, size) => all.slice(start, end)`.
(Note: Better to store pages in API if dataset is huge, but this works for small-medium lists).
