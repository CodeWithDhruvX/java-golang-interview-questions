## 🟢 Optimization, Testing & Entity Deep Dive (Questions 151-200)

### Question 151: How do you inject runtime configuration using meta-reducers?

**Answer:**
Inject the config service into the factory that creates the meta-reducer.
Access the config inside the reducer wrapper.
Use `APPS_INITIALIZER` to ensure config is loaded before Store boots.

---

### Question 152: How do you avoid unnecessary re-renders using selectors?

**Answer:**
Ensure selectors return the **same reference** if data hasn't effectively changed.
Use `distinctUntilChanged()` (implicit in `select()`).
Design state so frequently changing data (timestamps) doesn't pollute stable data.

---

### Question 153: What tools help you track performance issues in NgRx?

**Answer:**
1.  **Redux DevTools:** "Chart" view shows state size. "Test" view shows action frequency.
2.  **Chrome Performance Tab:** Identify "Long Tasks" triggered by `store.next`.

---

### Question 154: How do you optimize large lists in NgRx?

**Answer:**
*   **Virtual Scrolling (UI):** Render only visible items.
*   **Store:** Store as dictionary (Entity).
*   **Selector:** Select only IDs for the list (`selectIds`), then components select individual items by ID.

---

### Question 155: How do you debounce actions in NgRx?

**Answer:**
In Effects.
`ofType(InputChanged), debounceTime(300), switchMap(...)`.
The reducer sees every `InputChanged`, but the API only sees the debounced one.

---

### Question 156: What is `distinctUntilChanged` and how can it be useful with selectors?

**Answer:**
RxJS operator ignoring sequential duplicate emissions.
`store.select(getUser).pipe(distinctUntilChanged())` is redundant (built-in).
Useful if you map the selector: `select(getUser).pipe(map(u => u.name), distinctUntilChanged())`.

---

### Question 157: How do you track derived data performance?

**Answer:**
Profile the selector execution.
Wrap the projection function with `console.time()`.
If a selector takes > 2ms, it might need optimization.

---

### Question 158: How do you test selectors with memoization?

**Answer:**
Use `selector.release()` to clear cache between tests.
Test that calling it twice with same input returns same reference (toBe).

---

### Question 159: How do you mock API calls in effects tests?

**Answer:**
Spy on the Service method.
`spyOn(service, 'get').and.returnValue(of(mockData))`.
Run the effect Observable. Assert it outputs `SuccessAction(mockData)`.

---

### Question 160: How do you test meta-reducers?

**Answer:**
They are pure functions `(reducer) => newReducer`.
Create a dummy reducer.
Wrap it.
Call `wrappedReducer(state, action)`.
Verify the meta-logic (logging, reset) happened.

---

### Question 161: How do you test NgRx Entity selectors?

**Answer:**
Populate a state with entity adapter.
`const state = adapter.setAll([user1, user2], initialState)`.
Call `selectAll(state)`.
Expect array of users.

---

### Question 162: How do you simulate store state in component tests?

**Answer:**
`mockStore.setState({ ... })`.
Push a new state value at any time to verify component reaction (e.g., changes from Loading -> Loaded).

---

### Question 163: How do you test action dispatches from effects?

**Answer:**
Subscribe to the effect.
`modifiers.effects.myEffect$.subscribe(action => { expect(action.type).toBe('Success'); });`

---

### Question 164: How do you use `marble testing` in NgRx?

**Answer:**
Format: `'-a-b-|'` (time frames).
`actions$ = hot('-a', { a: load() });`
`expected = cold('-b', { b: success() });`
`expect(effect).toBeObservable(expected)`.

---

### Question 165: How can `TestScheduler` help in testing effects?

**Answer:**
Allows synchronous testing of async time-based operators (`debounceTime`, `delay`).
`scheduler.run(({ hot, cold, expectObservable }) => { ... })`.

---

### Question 166: How do you define a custom sort comparer for entity adapters?

**Answer:**
`createEntityAdapter({ sortComparer: (a, b) => a.name.localeCompare(b.name) })`.
State will always maintain entities in this sorted order (in `ids` array).

---

### Question 167: How do you extend entity state with additional flags like loading?

**Answer:**
`interface State extends EntityState<User> { loading: boolean; error: string; }`.
Initial state: `adapter.getInitialState({ loading: false, error: '' })`.

---

### Question 168: How do you handle relationships between entities (e.g., one-to-many)?

**Answer:**
**Normalized:**
`users: { 1: { id: 1, name: 'John' } }`.
`posts: { 101: { id: 101, userId: 1 } }`.
Do not nest posts inside users. Selectors join them: `selectUserPosts(1)`.

---

### Question 169: How do you select filtered subsets of entities?

**Answer:**
`createSelector(selectAll, (users) => users.filter(u => u.active))`.
Store contains *all*. Selectors return *subset*.

---

### Question 170: How do you patch entity metadata efficiently?

**Answer:**
`adapter.updateOne({ id: 1, changes: { name: 'New' } }, state)`.
Only updates provided fields. Shallow merges.

---

### Question 171: How do you override default HTTP behavior in NgRx Data?

**Answer:**
Override `DefaultDataService`.
Implement `execute()`.
Or register a custom service for a specific Entity in `EntityDataService`.

---

### Question 172: What are entity metadata maps in NgRx Data?

**Answer:**
Configuration object defining entities.
`const entityMetadata: EntityMetadataMap = { Hero: {}, Villain: {} };`
Defines sort order, filter functions, etc.

---

### Question 173: How do you handle optimistic updates in NgRx Data?

**Answer:**
Use `SaveOptimistic` option in `EntityDispatcher`.
`service.add(hero, { isOptimistic: true })`.
Updates cache immediately. Reverts on error.

---

### Question 174: How do you intercept error responses in NgRx Data?

**Answer:**
`EntityAction` stream.
Subscribe to `entityActions$.pipe(ofEntityType('Hero'), ofOp(EntityOp.SAVE_ADD_ONE_ERROR))`.
Or use a global ErrorInterceptor.

---

### Question 175: How do you extend the base entity service?

**Answer:**
`export class HeroService extends EntityCollectionServiceBase<Hero>`.
Add custom methods `getTopHeroes()`.
Inject `EntityCollectionServiceElementsFactory` in super.

---

### Question 176: How do you navigate using effects?

**Answer:**
Inject `Router`.
In effect: `tap(() => this.router.navigate(['/home']))`.
Set `dispatch: false`.

---

### Question 177: How do you listen to route changes using effects?

**Answer:**
Listen to `ROUTER_NAVIGATION` action (from router-store).
`ofType(routerNavigatedAction)`.
React to URL changes (e.g., auto-select tab).

---

### Question 178: How do you store router state in NgRx?

**Answer:**
`StoreModule.forRoot({ router: routerReducer })`.
`StoreRouterConnectingModule.forRoot()`.

---

### Question 179: How do you use `routerReducer` in your application?

**Answer:**
Register it under the key `router`.
It updates automatically. You rarely touch it manually.

---

### Question 180: What are router selectors and how are they used?

**Answer:**
`getRouterSelectors()`.
`selectRouteParam('id')`. `selectQueryParam('q')`.
Use in components to get URL data reactively from store.

---

### Question 181: How do you handle global errors in effects?

**Answer:**
CatchError in the stream.
Dispatch `GlobalError({ msg })`.
Global Error Reducer/Effect shows a Snackbar/Modal.

---

### Question 182: How do you show validation errors using NgRx?

**Answer:**
Store validation errors in state: `errors: { email: 'Invalid' }`.
Selector `selectErrors`.
Form binds to this selector to display red text.

---

### Question 183: What is the best pattern for dispatching error actions?

**Answer:**
Always return an Action in `catchError` (don't throw).
`catchError(err => of(ActionFailure({ payload: err })))`.
Keep the stream alive!

---

### Question 184: How do you retry failed actions in NgRx?

**Answer:**
`retry(3)` operator in Effect.
Or `retryWhen` with delay logic.
Retries the *inner* observable (API call), not the whole effect stream.

---

### Question 185: How do you integrate centralized logging for failed effects?

**Answer:**
Create a `ErrorLoggingEffect`.
`ofType(AllFailureActions)`.
`tap(action => logToSentry(action.error))`.

---

### Question 186: How do you implement a dynamic breadcrumb system using NgRx?

**Answer:**
Router Store -> Selector gets URL.
Selector maps URL to Breadcrumb Labels.
Component renders list.

---

### Question 187: How do you track user activity using NgRx actions?

**Answer:**
Meta-Reducer or overarching Effect.
`ofType(UserInteractions)`.
Update `lastActive` timestamp in state.
Auto-logout if `now - lastActive > 15min`.

---

### Question 188: How do you implement dark mode state management with NgRx?

**Answer:**
State: `theme: 'dark' | 'light'`.
Action: `ToggleTheme`.
Effect: `tap(theme => document.body.classList.add(theme))`. And save to LocalStorage.

---

### Question 189: How do you manage WebSocket real-time updates with NgRx?

**Answer:**
Effect creates `webSocket()`.
It `switchMaps` to the WS stream.
Incoming msg -> Dispatch `UpdateEntity`.
Store updates. UI updates live.

---

### Question 190: How do you manage state for a wizard-style form with multiple steps?

**Answer:**
State: `step: 1`, `formData: {}`.
Actions: `NextStep`, `PrevStep`, `UpdateData`.
Selectors: `selectCurrentStep`, `selectIsFormValid`.
Components subscribe to `selectCurrentStep` to show correct UI.

---

### Question 191: What is ComponentStore and how does it differ from Store?

**Answer:**
(Duplicate Q91).
It is an @Injectable service holding local state. No global dispatch.

---

### Question 192: When should you prefer ComponentStore over NgRx Store?

**Answer:**
(Duplicate Q92).
For isolating complex UI logic (Dropdowns, Grids) that shouldn't pollute global AppState.

---

### Question 193: How do you manage side effects in ComponentStore?

**Answer:**
`readonly load = this.effect<string>(params$ => ... )`.
It encapsulates the RxJS subscription management.

---

### Question 194: How do you compose selectors in ComponentStore?

**Answer:**
`readonly vm$ = this.select(this.users$, this.loading$, (users, loading) => ({ users, loading }))`.
Combines multiple internal selectors.

---

### Question 195: How do you combine local and global state with ComponentStore?

**Answer:**
Inject `GlobalStore`.
`this.select(this.state$, this.globalStore.select(selectUser), ...)`
Combine them in the `vm$` selector.

---

### Question 196: What are the benefits of NgRx Signal Store?

**Answer:**
Modern API.
Works seamlessly with Angular Signals.
No RxJS cognitive load for simple state.
Type-safe and extensible.

---

### Question 197: How do you create a computed signal from the store?

**Answer:**
`const double = computed(() => store.count() * 2)`.

---

### Question 198: How do signals improve developer ergonomics?

**Answer:**
Imperative reading `signal()` vs Declarative `observable$ | async`.
Reading value is easier in logic (no `.pipe(take(1))`).

---

### Question 199: How do you convert an observable selector to a signal selector?

**Answer:**
`toSignal(store.select(selector))`.
Note: It needs an injection context or `Injector`.

---

### Question 200: What is the future of NgRx in Angular with Signals?

**Answer:**
**SignalStore** will likely become the default for feature state.
Global Store might remain for complex event sourcing.
Signals reduce the need for heavy RxJS in Views.
