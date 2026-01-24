## 🟢 Component Interactions, Store Behavior & Patterns (Questions 101-150)

### Question 101: How do you access the store in a component?

**Answer:**
Inject `Store<AppState>` in the constructor.
Use `.select(selector)` to get data.
Use `.dispatch(action)` to send data.

---

### Question 102: What is the role of `store.select()`?

**Answer:**
It returns an Observable slice of the state.
It applies memoization (if using selectors) and distinctUntilChanged (only emits when value actually differs).

---

### Question 103: What are the differences between using `async pipe` and `subscribe()` manually in NgRx?

**Answer:**
*   **`async` pipe:** Subscribes/Unsubscribes automatically. Handles `ChangeDetectionStrategy.OnPush` cleanly. Recommended.
*   **`subscribe()`:** Manual memory management required (`ngOnDestroy`). Use only when you need side effects inside component (e.g., navigating after state change).

---

### Question 104: How do you unsubscribe safely when using the store?

**Answer:**
1.  Use `async` pipe (best).
2.  `takeUntil(this.destroy$)` pattern.
3.  `Subscription.add()` and `unsubscribe()` in `ngOnDestroy`.

---

### Question 105: Can a component dispatch multiple actions at once?

**Answer:**
Yes. Just call `store.dispatch(a); store.dispatch(b);`.
Beware of race conditions if they depend on each other. Consider using an Effect to orchestrate them instead.

---

### Question 106: How do you trigger an effect without updating the state?

**Answer:**
Dispatch an action that no reducer listens to.
Or define a reducer case that returns `state` unchanged (identity).
Effects listen to *actions*, not state changes.

---

### Question 107: What is the best way to organize selectors in large applications?

**Answer:**
*   **Feature Files:** `fromUser.selectXXX`.
*   **Index Barrels:** Flatten imports.
*   **ViewModel Selectors:** Combine multiple selectors into one `selectVm` for the component to consume (`{ user, settings, isLoading }`).

---

### Question 108: How do you share selectors across multiple components?

**Answer:**
Define them in the store file (e.g., `user.selectors.ts`).
Export them.
Any component can import and use them. They are stateless pure functions.

---

### Question 109: What are the pros and cons of colocating state logic vs centralizing it?

**Answer:**
*   **Centralizing:** Easy to debug flow. Harder to scale (huge file).
*   **Colocating (Feature Stores):** Scalable. Better encapsulation. Harder to see "Global" flow.
NgRx encourages Feature Stores (Colocation).

---

### Question 110: What happens if a reducer doesn’t return a new state?

**Answer:**
If it returns `undefined`, the store breaks (Error).
If it performs mutation (`state.prop = 1` and returns same object), Angular Change Detection won't trigger UI updates because object reference is unchanged.

---

### Question 111: Can two reducers handle the same action?

**Answer:**
Yes!
Action `[Auth] Logout` can be handled by `UserReducer` (clear user) AND `ProductReducer` (clear cart).
This is the power of Event-Driven architecture.

---

### Question 112: How do you handle race conditions in effects?

**Answer:**
Use appropriate flattening operators.
`switchMap`: Cancels old requests (good for Search).
`concatMap`: Queues requests (good for Save).
`exhaustMap`: Ignores new requests while active (good for Login).

---

### Question 113: What happens if multiple effects listen to the same action?

**Answer:**
All of them trigger.
Order is not guaranteed (synchronously they run in registration order, but async completion depends on network).
They run independently.

---

### Question 114: Can you delay the execution of an effect?

**Answer:**
Yes. Use `delay()` operator.
`this.actions$.pipe(ofType(..), delay(1000), ...)`

---

### Question 115: How do you prevent state mutation in NgRx?

**Answer:**
*   **Linting:** `ngrx-no-reducer-mutation`.
*   **Runtime:** `ngrx-store-freeze` (throws error in Dev if you mutate).
*   **TS:** Use `readonly` properties.

---

### Question 116: What are the risks of mutating state directly inside reducers?

**Answer:**
It breaks:
1.  **Time Travel Debugging** (History is corrupted).
2.  **Selectors** (Memoization relies on usage of new references).
3.  **UI Updates** (`OnPush` components won't re-render).

---

### Question 117: How do you debug incorrect state updates?

**Answer:**
Redux DevTools.
Look at the `diff` tab for the specific action.
Verify if the change matches expectation. Check the reducer logic corresponding to that action type.

---

### Question 118: How does immutability affect performance in NgRx?

**Answer:**
**Positive:** Faster change detection (Reference check `===` is O(1) vs Deep Compare O(N)).
**Negative:** Object creation overhead (GC pressure). Usually negligible compared to DOM rendering gains.

---

### Question 119: What is an action union type and how do you use it?

**Answer:**
(Legacy Pattern).
`type UserActions = LoadUser | LoadUserSuccess;`
Used in Reducer: `function reducer(state, action: UserActions)`.
Not needed with modern `createReducer`.

---

### Question 120: How can you group related actions for cleaner code?

**Answer:**
use `createActionGroup`.
```typescript
export const PageActions = createActionGroup({
  source: 'Product Page',
  events: {
    'Enter': emptyProps(),
    'Load Data': props<{ id: string }>()
  }
});
```

---

### Question 121: What are success/failure actions and why are they useful?

**Answer:**
`Load`, `LoadSuccess`, `LoadFailure`.
Standard Async Pattern.
Allows the UI to transition states: Loading -> Data (or Error).
Keeps side effects (API Error) separate from trigger (User Click).

---

### Question 122: What is the Command–Query Separation pattern in NgRx?

**Answer:**
*   **Command (Action):** "Do This" (Trigger).
*   **Query (Selector):** "Get This" (Read).
We never "Get and Do" in one step.

---

### Question 123: What is an action creator factory pattern?

**Answer:**
Functions that return Actions.
`createAction` generates these factory functions.
`const action = loadUser({ id: 1 });`

---

### Question 124: How do you enforce strict typing in actions and state?

**Answer:**
TypeScript generics.
`createAction('Type', props<{ id: number }>)` enforces that `dispatch(loadUser({ id: "string" }))` is a compile error.

---

### Question 125: How can you define reusable actions?

**Answer:**
Generic Actions using Generics (harder).
Better: Factory functions that generate a set of actions for a given "Feature Key".
`createCrudActions('User')` -> returns `LoadUser`, `UpdateUser`, etc.

---

### Question 126: What is the DRY way to define action constants?

**Answer:**
Use `createActionGroup` (Angular 14+).
Define source once, define events object. It generates string types `[Source] Event Name` automatically.

---

### Question 127: How do you isolate state for a lazy-loaded module?

**Answer:**
`StoreModule.forFeature('lazyFeature', lazyReducer)`.
This key `'lazyFeature'` only exists in the global object when module loads.

---

### Question 128: What is the benefit of `StoreFeatureModule`?

**Answer:**
It handles the injection of the reducer and effects automatically.
Allows code splitting (reducer code is not in main bundle).

---

### Question 129: How do you reuse logic between feature stores?

**Answer:**
Shared Reducers (Higher Order Reducers).
Or Utility functions.
If multiple features need "Pagination", write a `paginationReducer` and compose it in `featureAReducer` and `featureBReducer`.

---

### Question 130: What’s the difference between root store and feature store?

**Answer:**
*   **Root:** `StoreModule.forRoot`. Always present (Auth, Config).
*   **Feature:** `StoreModule.forFeature`. Route-dependent (Products, Admin).

---

### Question 131: How do you dynamically add feature states at runtime?

**Answer:**
The `Store` service doesn't expose `addReducer`.
You must use `ReducerManager.addReducer(key, reducer)`.
This is what `forFeature` does under the hood.

---

### Question 132: How can you clear feature state when leaving a module?

**Answer:**
In `ngOnDestroy` of the module's root component (or Route Guard), dispatch a `ClearState` action.
The reducer listens to it and resets to initial state.
Or use `ReducerManager.removeReducer` (advanced).

---

### Question 133: How do you combine multiple selectors?

**Answer:**
`createSelector` accepts up to 8 input selectors.
`createSelector(selectUser, selectOrders, (user, orders) => ...)`

---

### Question 134: How do you derive computed state using selectors?

**Answer:**
Selectors IS the way to derive state.
`selectTotalPrice = createSelector(selectItems, items => items.reduce(...))`
Never store `totalPrice` in the store if it can be calculated from `items`.

---

### Question 135: How do you filter entities using selectors?

**Answer:**
`createSelector(selectAll, selectFilter, (items, filter) => items.filter(i => i.name.includes(filter)))`.
Memoization ensures this only runs when items or filter changes.

---

### Question 136: Can you create parameterized selectors?

**Answer:**
Yes, using factory functions (props).
`const selectById = (id: string) => createSelector(..., (entities) => entities[id])`.
**Warning:** This breaks standard memoization if not used carefully (creates new selector instance every time).

---

### Question 137: What is `createSelectorFactory`?

**Answer:**
Allows customizing the memoization strategy.
Default is `defaultMemoize` (reference check).
You can swap it for `deepEqual` check if you have mutable data (rare).

---

### Question 138: How do you memoize selectors manually?

**Answer:**
You usually don't.
But `defaultMemoize` function is exported if you want to use it in your own utility functions.

---

### Question 139: When should you avoid recomputing selectors?

**Answer:**
Always. That's why we use selectors.
If a calculation is heavy (Sorting 5k rows), ensure the inputs are stable references so it rarely runs.

---

### Question 140: How do you handle side effects with multiple API calls?

**Answer:**
In Effect:
`switchMap(action => forkJoin([ callA(action.id), callB(action.id) ]))`.
Returns `[resA, resB]`. Map to `Success({ resA, resB })`.

---

### Question 141: How do you chain effects together?

**Answer:**
Effect A dispatches Action A_Success.
Effect B listens to Action A_Success -> performs task -> dispatches B_Success.
Waterfall pattern.

---

### Question 142: Can effects dispatch multiple actions?

**Answer:**
Yes. Return an array of actions and use `switchMap`.
`switchMap(() => [ ActionA(), ActionB() ])`.
(Requires `concatMaps` or similar if order strictness matters, but usually array emit works).

---

### Question 143: How do you handle long-running background processes with effects?

**Answer:**
Dispatch `StartProcess`.
Effect: `switchMap(() => task$.pipe(takeUntil(stopAction$)))`.
Dispatch `StopProcess` to kill it.

---

### Question 144: How do you perform optimistic UI updates with rollback?

**Answer:**
(See Q45).
Key is the `catchError` in Effect:
`catchError(error => of(UndoAction(), NotificationAction(error)))`.

---

### Question 145: What is the difference between `dispatch: true` and `false` in effects?

**Answer:**
*   `true` (Default): The Effect must return an Action, which gets dispatched back to Store.
*   `false`: The Effect does not return an Action (e.g., just `tap(() => log())` or `router.navigate`).

---

### Question 146: How do you deal with API polling using NgRx effects?

**Answer:**
`switchMap(() => timer(0, 5000).pipe( switchMap(() => apiCall()) ))`.
Starts polling on action. Stops on unsubscribe (component destroy or cancel action).

---

### Question 147: How do you pass additional state into effects?

**Answer:**
`concatLatestFrom(() => this.store.select(selectFilter))`.
Combines the Action with current State value.

---

### Question 148: What is a meta-reducer?

**Answer:**
(Duplicate Q97).
Middleware for reducers. `(reducer) => (state, action) => newState`.

---

### Question 149: How do you log all actions and state transitions with a meta-reducer?

**Answer:**
```typescript
export function debug(reducer) {
  return (state, action) => {
    console.log('Action', action);
    console.log('State', state);
    return reducer(state, action);
  };
}
```

---

### Question 150: How can you implement undo/redo using meta-reducers?

**Answer:**
(Duplicate Q98).
Maintain `past`, `present`, `future` arrays in a wrapper state.
Intercept `UNDO` action to manipulate these pointers.
