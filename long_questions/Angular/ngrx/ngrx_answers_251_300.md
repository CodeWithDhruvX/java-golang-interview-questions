## 🟢 Forms, Architecture & Monitoring (Questions 251-300)

### Question 251: How do you bind NgRx state to reactive forms?

**Answer:**
Bind state to the form using `setValue` or `patchValue`.
Listen to `valueChanges` and dispatch actions to update state.
`store.select(data).subscribe(v => form.patchValue(v, { emitEvent: false }))`.
`form.valueChanges.subscribe(v => store.dispatch(UpdateAction(v)))`.

---

### Question 252: How do you update the store on every keystroke?

**Answer:**
`form.get('input').valueChanges.subscribe(val => dispatch(Update({ val })))`.
**Performance risk:** Use `debounceTime(300)` to avoid flooding the store/Redux DevTools.

---

### Question 253: What is the pattern for form state vs API state?

**Answer:**
**API State:** The source of truth (from Database).
**Form State:** Temporary "Draft" state (User typing).
Keep them separate.
On Save: Copy Form State -> API State (via Effect).

---

### Question 254: How do you use selectors to prefill forms?

**Answer:**
In `ngOnInit`:
`this.store.select(selectItem).pipe(take(1)).subscribe(item => this.form.patchValue(item));`
Use `take(1)` if you only want to initialize. Leave open if you want real-time updates from other users.

---

### Question 255: How do you reset a form using an action?

**Answer:**
Dispatch `ResetForm`.
Reducer clears the "Draft" state.
Component listens to `selectDraft` (which becomes null) -> calls `form.reset()`.

---

### Question 256: How do you show validation errors using store state?

**Answer:**
Server-side validation errors return as Actions (`SaveFailure({ errors })`).
Store saves keys/messages.
UI selects `selectErrors` and maps field name to error message shown under input.

---

### Question 257: How do you sync tab navigation with store state?

**Answer:**
Store `selectedTabId`.
Selector drives the `[active]` input of tabs.
Event `(tabChange)` dispatches `SetTabId`.
Ensures tabs persist on refresh (if synced to LocalStorage).

---

### Question 258: How do you persist UI preferences like theme, layout using NgRx?

**Answer:**
Slice `ui: { theme: 'dark', sidebarCollapsed: true }`.
Persist this slice via `ngrx-store-localstorage`.
Apply changes via Effect (e.g., adding CSS classes).

---

### Question 259: How do you animate UI changes driven by state?

**Answer:**
Angular Animations trigger on value changes.
`<div [@fadeIn]="count$ | async">`.
When store updates, the pipe updates the binding, triggering the animation.

---

### Question 260: How do you bind NgRx data to Angular Material components?

**Answer:**
Material Table `dataSource`.
Connect `dataSource.data` to the selector Observable.
`this.store.select(selectRows).subscribe(data => this.dataSource.data = data);`

---

### Question 261: How do you plan NgRx store structure before development?

**Answer:**
Whiteboard the JSON tree.
Identify "Feature Slices" (Users, Products).
Identify "Shared State" (Auth).
Decide on Normalization strategy early.

---

### Question 262: What is the folder structure for a scalable NgRx app?

**Answer:**
Feature-based.
`/libs/products/state/src/lib/...`
Files: `products.actions.ts`, `products.reducer.ts`, `products.selectors.ts`, `products.effects.ts`, `products.facade.ts`.

---

### Question 263: How do you scale NgRx for enterprise-grade applications?

**Answer:**
1.  **Strict Rules:** No direct store access in components (Facades).
2.  **Monorepo:** Libraries enforce boundaries.
3.  **Schematics:** Auto-generate code.
4.  **Linting:** Enforce immutability.

---

### Question 264: How do you document your entire state tree?

**Answer:**
Use `interface AppState` that imports all feature states.
Visualize using DevTools "State" tab export.
TSDoc comments on the interface properties.

---

### Question 265: How do you track API call statistics using the store?

**Answer:**
Generic `RequestsState`.
`{ [reqId]: { startTime, endTime, status } }`.
Effects dispatch start/end events.
Selectors calculate `averageDuration`.

---

### Question 266: How do you handle versioning of state schema?

**Answer:**
If persisted (LocalStorage):
Check `version` key on hydration.
If `version < newVersion`, run migration function (transform object) or wipe storage.

---

### Question 267: How do you manage legacy APIs and state differences?

**Answer:**
**Adapter Pattern (in Effect):**
API returns XML/Legacy JSON.
Effect transforms it to clean modern Interface.
Store only holds clean data. UI never knows about legacy API shape.

---

### Question 268: How do you migrate a non-NgRx app to NgRx?

**Answer:**
**Strangler Pattern.**
1.  Introduce Store.
2.  Move *one* leaf feature (e.g., "Settings") to Store.
3.  Keep Services for now.
4.  Gradually move heavy shared state (User).

---

### Question 269: How do you plan refactoring large reducers safely?

**Answer:**
Unit Tests first!
Ensure input/output is locked.
Split logic into helper functions or sub-reducers.
`combineReducers` or functional composition.

---

### Question 270: What metrics indicate it's time to introduce NgRx?

**Answer:**
1.  "Prop Drilling" > 3 levels deep.
2.  Same data needed in Sidebar, Header, and Main Content.
3.  Complex optimistic UI requirements.
4.  Race conditions in bug reports.

---

### Question 271: How do you capture usage analytics with NgRx?

**Answer:**
**Meta-Reducer or Effect.**
`ofType(UserAction)`.
Send payload to Google Analytics / Mixpanel.
"Action dispatched" = "User Event".

---

### Question 272: How do you send tracking events on action dispatch?

**Answer:**
`tap(action => analytics.track(action.type, action.props))`.
Centralized place for all tracking. Clean components.

---

### Question 273: How do you detect unused actions or stale state?

**Answer:**
Static Analysis (`eslint-plugin-ngrx`).
Code search: "Is this action dispatched anywhere?".
DevTools: "Is this action ever fired in usage session?".

---

### Question 274: How do you analyze reducer performance over time?

**Answer:**
DevTools Profiler.
Or custom meta-reducer that `console.time(action.type)` around the reducer call.

---

### Question 275: How do you log time spent in each store state?

**Answer:**
Instrumentation.
Track `timestamp` of state change.
Visualize flow.

---

### Question 276: How do you track failed actions and log them centrally?

**Answer:**
(Duplicate Q185).
Effect listening to `*Failure`.
Send to Sentry.

---

### Question 277: How do you implement user journey tracking with store transitions?

**Answer:**
The sequence of Actions *is* the user journey.
`Navigate -> LoadProduct -> AddToCart -> Checkout`.
Log this sequence for funnel analysis.

---

### Question 278: How do you use effects to report metrics to external services?

**Answer:**
Effect marked `dispatch: false`.
Calls `metrics.send(...)`.
Does not disturb application flow.

---

### Question 279: How do you limit noisy actions from spamming analytics?

**Answer:**
Filter list.
`const IGNORE = [Scroll, MouseMove, InputDebounce]`.
Only log "Business Events" (Click, Submit).

---

### Question 280: How do you throttle high-frequency store updates?

**Answer:**
**Backend (Effect):** `throttleTime`.
**Frontend (Select):** `sampleTime` or `auditTime`.
If store updates 60fps, UI might only need 10fps updates.

---

### Question 281: What are the lifecycle methods in NgRx Signal Store?

**Answer:**
`withHooks({ onInit: () => ..., onDestroy: () => ... })`.
Allows setup/teardown logic when the store is provided/destroyed.

---

### Question 282: How do you convert a traditional reducer to a Signal Store model?

**Answer:**
Define `withState(initial)`.
Define `withMethods(store => ({ update: (val) => patchState(store, { val }) }))`.
Reducers become methods that call `patchState`.

---

### Question 283: What’s the difference between a computed signal and memoized selector?

**Answer:**
Mechanically similar.
**Computed:** Angular primitive. Fine-grained.
**Selector:** NgRx primitive. Memoized reference check.

---

### Question 284: How do you trigger effects from signal updates?

**Answer:**
`effect(() => { const x = store.val(); api.call(x); })`.
Use Angular's `effect` primitive inside the injection context.

---

### Question 285: How do signals affect app performance?

**Answer:**
Positive.
Zoneless change detection capability.
Updates pinpoint specific DOM nodes rather than checking whole component tree.

---

### Question 286: How does signal-based state management improve ergonomics?

**Answer:**
No Observables. No `$` suffix. No `subscription`.
`store.users()` returns current value. Simple.

---

### Question 287: How do you debug signal stores?

**Answer:**
Redux DevTools support is coming/plugin-based.
Or plain checking `store.state()` in console.

---

### Question 288: Can signals and observables coexist in a component?

**Answer:**
Yes.
`count$` (Observable) and `count()` (Signal).
Use `toSignal` and `toObservable` to bridge them.

---

### Question 289: How do you hydrate signal stores with SSR data?

**Answer:**
Pass data in `withState` initial value (read from TransferState).
Unlike NgRx Store (global), Signal Stores are often local, but can be global (provided in root).

---

### Question 290: What are the caveats of using Signal Store in production?

**Answer:**
Newer API (evolving patterns).
Ecosystem (plugins) smaller than Redux.
Learning curve for RxJS devs shifting mindset to Signals.

---

### Question 291: How do you deal with conflicting actions from multiple tabs?

**Answer:**
(Duplicate Q202/BroadcastChannel).
Use storage events to sync tabs.
Or optimistic locking on backend (Version ID).

---

### Question 292: How do you merge multiple API responses into one reducer?

**Answer:**
Action: `LoadAllSuccess({ users, products })`.
Reducer: Updates both slices.
Or separate actions dispatched via `concatMap`: `[LoadUsersSuccess, LoadProductsSuccess]`.

---

### Question 293: How do you cancel previously triggered actions?

**Answer:**
You can't "cancel" an action once dispatched (it's synchronous).
You can cancel the **Effect** (async process) triggered by it using `switchMap`.

---

### Question 294: What happens if an effect throws an unhandled exception?

**Answer:**
The Effect Stream **dies** (unsubscribes).
It won't listen to future actions.
**CRITICAL:** Always `.pipe(catchError(...))` *inside* the exhaust/switchMap.

---

### Question 295: How do you rollback the state if an effect fails?

**Answer:**
Dispatch `FailureAction`.
Reducer listens to Failure -> Reverts state (e.g., sets `loading: false`, restores `previousValue`).

---

### Question 296: What is the retry pattern in case of network flakiness?

**Answer:**
`retry({ count: 3, delay: 1000 })` inside the Effect's inner pipe.

---

### Question 297: How do you debug infinite loops in effects?

**Answer:**
Check if Action A triggers Effect -> dispatches Action A (cycle).
DevTools shows 1000s of actions firing instantly.
Fix logic or use `dispatch: false`.

---

### Question 298: How do you manage state for infinite scroll UIs?

**Answer:**
Action: `LoadMore`.
Reducer: `items: [...state.items, ...newItems]`.
Selector: Returns flat list.

---

### Question 299: How do you ensure store consistency after app crashes or reloads?

**Answer:**
Persist to LocalStorage.
Rehydrate on boot.
Validate schema version to ensure structure matches code.

---

### Question 300: How do you implement soft delete and restore in NgRx state?

**Answer:**
Action `SoftDelete`.
Reducer: sets `deleted: true`.
Selector filters `deleted: false`.
Action `Restore` sets `deleted: false`.
Effect calls API to flag logic.
