## 🟢 DevOps, Performance & Edge Cases (Questions 451-500)

### Question 451: What are meta-reducers and how do you use them for analytics or logging?

**Answer:**
Functions that wrap the root reducer.
Use for Logging: `(reducer) => (state, action) => { log(action); return reducer(state, action); }`.
Use for Hydration: Load state from LocalStorage before passing to reducer.

---

### Question 452: How do you instrument actions for dev tooling?

**Answer:**
`StoreDevtoolsModule.instrument()`.
It monkey-patches the dispatcher to send actions to the browser extension.

---

### Question 453: How do you create a reducer that listens to external action streams?

**Answer:**
Not typical. Usually you dispatch actions *from* external streams (Effects).
But you could manually call `reducer(state, externalAction)` if you control the loop.

---

### Question 454: How do you skip reducer execution under certain state conditions?

**Answer:**
Inside the reducer case:
`on(Action, (state) => state.locked ? state : { ...state, val: 1 })`.
Return current state to change nothing.

---

### Question 455: How do you avoid race conditions in reducer/effect combinations?

**Answer:**
Reducers are synchronous (No race conditions possible).
Effects are async (Use `concatMap` to strictly order execution).

---

### Question 456: How do you rollback state when an effect fails?

**Answer:**
(Duplicate Q295).
Dispatch Failure Action. Reducer reverts.

---

### Question 457: How do you add audit logging inside reducers cleanly?

**Answer:**
Don't. Reducers must be pure.
Audit logging is a side effect -> Use an Effect or Meta-Reducer.

---

### Question 458: How do you handle action queueing or debouncing in reducers?

**Answer:**
You cannot debounce in a reducer (it is sync).
Debounce in the *Effect* or Component before dispatching.

---

### Question 459: How do you dispatch multiple actions sequentially with state dependency?

**Answer:**
Effect dispatching A.
Effect listening to A success -> checks store state -> dispatches B.
Waterfall.

---

### Question 460: What are the risks of misusing spread operator in immutable reducer logic?

**Answer:**
Shallow copies only.
`{ ...state, nested: { ...val } }`.
If you forget one level of spread, you mutate the reference shared by `state` and `newState`, breaking change detection.

---

### Question 461: How do you deal with performance issues in a large NgRx-based dashboard?

**Answer:**
1.  **OnPush** change detection everywhere.
2.  **Memoize** selectors heavily.
3.  **Virtual Scroll** for lists.
4.  **Debounce** high-frequency inputs.

---

### Question 462: How do you manage scroll position or tab state with NgRx?

**Answer:**
Sync to Store.
On Component Init, `window.scrollTop = store.selectSnapshot(pos)`.
Or use Router Store (scroll position restoration).

---

### Question 463: How do you apply route-based lazy state hydration?

**Answer:**
(Duplicate Q246).
Guard or Resolver dispatches `LoadAction`.

---

### Question 464: How do you track user journey using store updates?

**Answer:**
(Duplicate Q277).
Action sequence.

---

### Question 465: How do you integrate NgRx with Capacitor or Cordova apps?

**Answer:**
No difference. NgRx is JS.
Effect might call `Capacitor.Plugins.Camera.getPhoto()`.

---

### Question 466: How do you reduce store initialization time for large apps?

**Answer:**
Lazy Load features.
Don't load "Admin" state if user is "Guest".
Use `initialState` only for crucial config.

---

### Question 467: How do you prevent store bloat when managing a notification center?

**Answer:**
Cap the list size in the reducer.
`notifications: [new, ...state.slice(0, 49)]`.
Keep max 50 items.

---

### Question 468: How do you sync state with offline-first mobile apps?

**Answer:**
Persist Action Queue to IndexedDB when offline.
Replay Actions when online.
`ngrx-offline` libraries help.

---

### Question 469: How do you trace user session actions using store instrumentation?

**Answer:**
Sentry Breadcrumbs integration.
Every action dispatched adds a "breadcrumb".
On Error, context shows the last 10 actions.

---

### Question 470: How do you implement configurable widgets with NgRx for dashboards?

**Answer:**
`widgets: { [id]: { type: 'chart', config: {} } }`.
Component `ngSwitch` on type. Selects config.

---

### Question 471: How do you integrate a global toast notification system via NgRx?

**Answer:**
(Duplicate Q69).

---

### Question 472: How do you build a wizard flow using NgRx state transitions?

**Answer:**
(Duplicate Q190).

---

### Question 473: How do you manage unsaved draft content in the store?

**Answer:**
Separate slice `drafts`.
On "Cancel", clear drafts.
On "Save", commit drafts to entities.

---

### Question 474: How do you implement dynamic ACL (access control) with store-based logic?

**Answer:**
State: `permissions: string[]`.
Directives `*hasPermission="'admin'"` check store.

---

### Question 475: How do you protect feature modules with store-based permission guards?

**Answer:**
CanActivate Guard.
`store.select(isLoggedIn).pipe(filter(x => x), take(1))`.

---

### Question 476: How do you handle application-level startup data fetch with NgRx?

**Answer:**
`APP_INITIALIZER` dispatches `AppInit`.
Effect listens `AppInit` -> `LoadConfig`.
Block startup until Config loaded (Promise).

---

### Question 477: How do you implement internationalization (i18n) with NgRx?

**Answer:**
(Duplicate Q308).

---

### Question 478: How do you deal with duplicate network calls caused by eager effects?

**Answer:**
`shareReplay` in the Service.
Or check Store content in Effect before calling Service (`withLatestFrom`).

---

### Question 479: How do you implement a review/approval flow with state-based transitions?

**Answer:**
Status enum: `Draft -> Pending -> Approved`.
Actions: `Submit`, `Approve`, `Reject`.
Reducer enforces valid transitions (can't go `Draft -> Approved` directly).

---

### Question 480: How do you handle animation state flags using NgRx?

**Answer:**
`isExpanded: boolean`.
Animations trigger on boolean toggle.

---

### Question 481: How do you use Redux DevTools with large NgRx applications?

**Answer:**
Filter actions (ignore `MouseMove`).
Increase `maxAge` (history size) cautiously (memory heavy).

---

### Question 482: How do you monitor store performance using custom logging?

**Answer:**
(Duplicate Q206).
Time the reducer.

---

### Question 483: How do you set up state inspection for debugging in production?

**Answer:**
**Dangerous.** Exposes PII.
If needed, allow a hidden gesture to enable DevTools or export state to console.
Usually keep DevTools off in Prod.

---

### Question 484: How do you integrate NgRx with Sentry or Firebase Crashlytics?

**Answer:**
Global ErrorHandler dispatches to Sentry.
Sentry integration can attach Redux state snapshot to the report.

---

### Question 485: How do you minify store state size for server-side rendering (SSR)?

**Answer:**
Only transfer critical state (`TransferState`).
Don't transfer big lists that can be refetched cheaply.

---

### Question 486: How do you export/import store snapshots for testing or recovery?

**Answer:**
DevTools "Export" button (JSON).
"Import" button to restore.
Programmatically: `JSON.stringify(state)`.

---

### Question 487: How do you secure sensitive state values in client-side storage?

**Answer:**
Encrypt before saving to LocalStorage.
CryptoJS.
Decrypt on rehydration.

---

### Question 488: How do you benchmark selector recompute performance?

**Answer:**
(Duplicate Q157).
Console.time.

---

### Question 489: How do you configure state persistence encryption?

**Answer:**
Library `redux-persist-transform-encrypt` (if using redux-persist port) or manual meta-reducer logic.

---

### Question 490: How do you detect state leaks in NgRx memory graph?

**Answer:**
Chrome Memory Profiler.
Take Snapshot.
Look for detached DOM nodes retained by Store Subscriptions (Memory Leak).

---

### Question 491: How do you enforce store contract types across CI pipelines?

**Answer:**
TypeScript `strict: true`.
Ensure all Actions have payload interfaces.
Ensure all Selectors have return types.

---

### Question 492: How do you generate state diagrams from NgRx structure?

**Answer:**
Tools like `ngrx-graph`.
Parses AST to show Actions -> Effects -> Reducers flow.

---

### Question 493: How do you audit store structure for scaling issues?

**Answer:**
Review `State` interface.
Is it flat?
Are there Arrays that should be Maps?

---

### Question 494: How do you prevent state explosion in client-heavy applications?

**Answer:**
Clean up.
Action `ClearFeatureState` when leaving module.
Don't keep 10,000 logs in memory.

---

### Question 495: How do you enable strict immutability checks during builds?

**Answer:**
`StoreModule.forRoot(reducers, { runtimeChecks: { strictStateImmutability: true } })`.
Throws error if you mutate.
Disable in Prod for performance.

---

### Question 496: How do you prevent unnecessary change detection triggered by store updates?

**Answer:**
Selectors.
If selector returns same reference, `async` pipe (and OnPush) does nothing.

---

### Question 497: How do you implement a reactive profiler for store observables?

**Answer:**
(Duplicate Q274).

---

### Question 498: How do you track state usage across Angular components?

**Answer:**
Search codebase for usages of `select(mySelector)`.
Delete unused selectors.

---

### Question 499: How do you integrate NgRx with feature flag systems?

**Answer:**
(Duplicate Q216).
LaunchDarkly SDK -> Effect -> Update Store -> Selectors.

---

### Question 500: How do you measure and visualize the NgRx action flow over time?

**Answer:**
Redux DevTools "Slider".
Or custom dashboard counting Action Velocity (Actions per minute).
