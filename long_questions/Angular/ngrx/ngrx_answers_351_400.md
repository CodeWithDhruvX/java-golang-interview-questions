## 🟢 Advanced Selectors, Tests & Migration (Questions 351-400)

### Question 351: How do you create dynamic views from an entity list (e.g., filter by tags)?

**Answer:**
Store `filterTags: string[]` in state.
Selector:
`createSelector(selectAll, selectTags, (items, tags) => items.filter(i => tags.every(t => i.tags.includes(t))))`

---

### Question 352: How do you write custom comparators for `EntityAdapter.sortComparer`?

**Answer:**
Function `(a, b) => number`.
`sortComparer: (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()`.
Sorts descending by date.

---

### Question 353: How do you bulk import data into NgRx Entity?

**Answer:**
`adapter.addMany(items, state)`.
Or `setAll` (if replacing).
Massively more efficient than looping `addOne`.

---

### Question 354: How do you sync sorted lists and entity IDs?

**Answer:**
The Entity Adapter maintains the `ids` array in sorted order (if `sortComparer` is provided).
`selectAll` returns entities in that order.

---

### Question 355: How do you filter or sort entities using component inputs?

**Answer:**
Factory Selector pattern.
`const selectSorted = (sortKey: string) => createSelector(selectAll, all => [...all].sort(...))`.
Or component filters the stream: `store.select(selectAll).pipe(map(list => sort(list)))`.

---

### Question 356: How do you derive computed fields on entity selectors?

**Answer:**
Selector projection.
`createSelector(selectUser, user => ({ ...user, fullName: user.first + ' ' + user.last }))`.
This creates a View Model object.

---

### Question 357: How do you sync user-specific flags (like starred, read, etc.) with entities?

**Answer:**
Ideally separate slice: `userMetadata: { [entityId]: { starred: boolean } }`.
Join in selector: `createSelector(selectEntities, selectMetadata, (entities, meta) => ...)`
Prevents re-fetching the main entity just to star it.

---

### Question 358: How do you track entity state transitions (new, edited, saved)?

**Answer:**
Add property `status: 'pristine' | 'dirty' | 'saving'` to the entity model (or a wrapper).
Adapter updates this field on actions `Update` vs `SaveSuccess`.

---

### Question 359: How do you enforce entity ID uniqueness across modules?

**Answer:**
Use UUIDs.
Or prefix IDs: `User_1`, `Product_1`.
Selectors can assert uniqueness if needed (Runtime check).

---

### Question 360: How do you define default values when adding entities?

**Answer:**
Middleware/Service/Reducer logic.
`on(AddUser, (state, { user }) => adapter.addOne({ role: 'Guest', ...user }, state))`.

---

### Question 361: How do you structure selectors for deeply nested state?

**Answer:**
Drill down step-by-step.
`selectFeature -> selectSubFeature -> selectProperty`.
Compose them. Don't write one giant function accessing `state.foo.bar.baz`.

---

### Question 362: How do you reuse selectors across modules?

**Answer:**
Export them from `libs/data-access`.
Other modules import `fromUser.selectUserIds`.

---

### Question 363: How do you memoize selectors with custom equality checks?

**Answer:**
`createSelectorFactory(defaultMemoize, { resultEqualityCheck: deepEqual })`.
Or `projectionFn` that returns specific references.

---

### Question 364: How do you create dynamic selectors based on component inputs?

**Answer:**
Pass props to `select`.
`store.select(selectUserById, { id: 123 })`.
Note: Deprecated in some versions; usually better to store `selectedId` in store and use `selectCurrentItem`.

---

### Question 365: What is the benefit of `props` in `createSelector`?

**Answer:**
Allows filtering/selecting specific items without putting the filter criteria in the Global Store. Good for multiple instances (e.g., multiple Item Cards).

---

### Question 366: How do you debug which selector is causing recomputation?

**Answer:**
Add `console.log` inside the projection function.
If it logs, memoization failed (inputs changed).

---

### Question 367: How do you combine entity selectors with derived state?

**Answer:**
`createSelector(selectAll, selectSettings, (items, settings) => applySettings(items, settings))`.

---

### Question 368: How do you create a filtered list with a selector based on route parameters?

**Answer:**
`createSelector(selectAll, selectRouteParam('category'), (items, cat) => items.filter(i => i.cat === cat))`.
Combines Entity Store + Router Store.

---

### Question 369: How do you implement lazy selectors for dynamic modules?

**Answer:**
Feature Selector is the key.
`createFeatureSelector<State>('lazyFeature')` works even if module loaded later. Returns undefined initially.

---

### Question 370: How do you prevent selector recomputation loops?

**Answer:**
Ensure projection functions are pure.
Don't return new object references `return { ...data }` if `data` didn't change content.

---

### Question 371: How do you unit test a selector that depends on others?

**Answer:**
Test the "Result Selector" (Projector).
You don't need to mock the inputs' internal logic, just their output values.

---

### Question 372: How do you refactor multiple similar selectors into DRY code?

**Answer:**
Higher Order Selector Creator.
`const createStatusSelector = (status) => createSelector(selectAll, list => list.filter(x => x.status === status))`.

---

### Question 373: How do you structure selectors for large UI state (tabs, modals, toggles)?

**Answer:**
`selectUiState` -> `selectTabState`, `selectModalState`.
Specific selectors: `selectIsLoginModalOpen`.

---

### Question 374: How do you apply pure functions in selector pipelines?

**Answer:**
Import helper utils (lodash/ramda).
`createSelector(selectList, R.sortBy(R.prop('date')))`.

---

### Question 375: How do you use `createFeatureSelector` with nested features?

**Answer:**
`createFeatureSelector` only gets the top-level slice registered with `StoreModule`.
For nested, you must manually drill: `createSelector(selectFeature, s => s.nested)`.

---

### Question 376: How do you trace selector performance in production?

**Answer:**
Rarely done. Can wrap projection functions with timing logic in Dev/Staging.

---

### Question 377: How do you test selectors with overridden state values?

**Answer:**
Pass any state shape to the `projector`. It doesn't validate if the upstream changes were possible, just computes the result.

---

### Question 378: How do you write selectors for aggregate calculations (sum, avg)?

**Answer:**
`createSelector(selectTransactions, txs => txs.reduce((sum, tx) => sum + tx.amount, 0))`.

---

### Question 379: How do you trace recomputations using selector debug tools?

**Answer:**
Redux DevTools doesn't show selector recomputes.
Use `ngrx-monitor-selector` (community tool) or manual logging.

---

### Question 380: How do you use selectors in Angular Signals-based components?

**Answer:**
`userId = this.store.selectSignal(selectUserId)`.
`user()` evaluates the selector reactively.

---

### Question 381: How do you migrate from manual state management to NgRx?

**Answer:**
Identify Services holding state (`behaviorsubject`).
Port State interface to NgRx.
Port methods to Actions/Reducers.
Switch components to `store.select()`.

---

### Question 382: How do you migrate from legacy actions to modern `createAction` API?

**Answer:**
Replace `enum ActionTypes` with string constants (or inline).
Replace classes `implements Action` with `createAction`.
Replace switch-statement Reducer with `createReducer(on(...))`.

---

### Question 383: What are common mistakes when migrating to NgRx from services?

**Answer:**
1.  Keeping state in Services AND Store (Two sources of truth).
2.  Dispatching actions like "SetVariable" instead of "UserEvent" (RPC style).

---

### Question 384: How do you identify tight coupling during migration to NgRx?

**Answer:**
If Component A injects Service B to read data, they are coupled.
NgRx decouples them (both depend on Store).

---

### Question 385: How do you ensure store doesn’t grow beyond control after migration?

**Answer:**
Use `ComponentStore` for non-shared state.
Don't put *everything* in Redux.

---

### Question 386: How do you unit test NgRx reducer logic?

**Answer:**
(Duplicate Q61).
Test state transitions.

---

### Question 387: How do you test integration of selectors and store in components?

**Answer:**
(Duplicate Q66).
`MockStore.overrideSelector`.

---

### Question 388: How do you test deeply nested selectors?

**Answer:**
Just test the final selector.
Provide the deep state structure to the projector.

---

### Question 389: How do you isolate and test side effects using mocks?

**Answer:**
Mock the Action Stream (`provideMockActions`).
Mock the Service (`jasmine.createSpyObj`).

---

### Question 390: How do you mock store values for component testing?

**Answer:**
`provideMockStore({ initialState: { users: [] } })`.

---

### Question 391: How do you use `MockStore` in unit tests?

**Answer:**
Inject it: `store = TestBed.inject(MockStore)`.
dispatch: `spyOn(store, 'dispatch')`.
select: `store.overrideSelector(...)`.

---

### Question 392: How do you test a component that dispatches multiple actions?

**Answer:**
Trigger interactions.
Expect `store.dispatch` to have been calledWith Arg1, then Arg2.

---

### Question 393: How do you test loading spinners or UI flags from store state?

**Answer:**
`store.setState({ loading: true })`.
`fixture.detectChanges()`.
Assert spinner exists in DOM.
`store.setState({ loading: false })`.
Assert spinner removed.

---

### Question 394: How do you test error propagation from effects to UI?

**Answer:**
Mock Effect to return ErrorAction.
Or mock Store to have Error State.
Verify UI displays error message.

---

### Question 395: How do you test behavior across feature module boundaries?

**Answer:**
Integration Tests (Spectator/Cypress).
Unit tests usually mock the boundary.

---

### Question 396: How do you automate regression tests for NgRx workflows?

**Answer:**
Snapshot tests of State.
If `Action -> State` output changes unexpectedly, fail.

---

### Question 397: How do you write test utilities for NgRx patterns?

**Answer:**
Factory for Mock State.
`createMockUser()`.
Simplifies setup in `beforeEach`.

---

### Question 398: How do you simulate delays or network latency in tests?

**Answer:**
`fakeAsync`, `tick(1000)`.
Or RxJS `TestScheduler`.

---

### Question 399: How do you validate store rehydration in e2e tests?

**Answer:**
Cypress:
1.  Login.
2.  Reload Page.
3.  Assert user still logged in (Store rehydrated from Storage).

---

### Question 400: How do you mock third-party service calls inside effects for testing?

**Answer:**
Dependency Injection.
Provide `{ provide: ThirdPartyService, useValue: mockService }`.
The Effect uses the injected mock.
