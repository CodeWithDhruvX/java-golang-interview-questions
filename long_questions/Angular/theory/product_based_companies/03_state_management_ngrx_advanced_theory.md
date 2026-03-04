# 🗣️ Theory — State Management & NgRx Advanced
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do NgRx Effects work? Why do reducers not handle side effects?"

> *"Reducers must be pure functions — they take state and an action and return new state. No async operations, no HTTP calls, no console.log, nothing with side effects. This is essential because reducers can be called during time-travel debugging replays, testing, and state rehydration — you need them to be deterministic and predictable. Effects exist to handle side effects triggered by actions. An Effect listens to the Actions stream using ofType to filter for specific actions, performs the async work — typically an HTTP call — and dispatches new actions based on the result. So 'Load Products' triggers an Effect that calls the API, and on success dispatches 'Load Products Success' with the data, or 'Load Products Failure' with the error message. Reducers handle both outcomes and update the store accordingly. This strict separation makes the data flow auditable."*

---

## Q: "What is ngrx/entity and what problem does it solve?"

> *"Managing a collection of items in NgRx manually is verbose and error-prone. You'd write reducers to add one item to an array by spreading the array and checking for duplicates; to remove one item by filtering; to update one item by mapping. And you'd always need to consider the case where you're loading 500 products — storing them in a plain array makes lookups O(n). ngrx/entity provides EntityState — a normalized structure with ids and entities: an array of IDs and a dictionary mapping each ID to the full entity. EntityAdapter gives you adapter methods — setAll, addOne, updateOne, removeOne — that return new state directly. Behind the scenes it manages the ids array and entities dictionary together atomically. You also get built-in selectors — selectAll, selectEntities, selectTotal — for free. It's a significant reduction in boilerplate for any feature that manages a collection."*

---

## Q: "How does NgRx selector memoization work?"

> *"createSelector builds a memoized selector from one or more input selectors and a projector function. When the selector is called, it first runs the input selectors. If their return values are the same as last time — by reference equality — it skips running the projector and returns the previously computed result from its cache. Only when inputs change does it run the projector. This matters because selectors are called inside Angular's change detection pipeline, potentially on every tick. If selectAllProducts returns the same array reference as last time, selectFilteredProducts's filter and sort logic doesn't run — the cached result is returned immediately. For expensive derived state like sorting 1000 items or computing aggregates, memoization means the cost is paid only when the underlying data actually changes."*

---

## Q: "What is ComponentStore in NgRx? When would you use it over the global Store?"

> *"ComponentStore is NgRx's solution for local state — state that belongs to one component and its subtree, not the entire application. You provide it in a component's providers array so it's scoped to that component, and Angular creates a new instance for each component instance. ComponentStore has its own state, updaters for synchronous mutations, selectors for derived state, and effects for async operations — the same concepts as the global Store but self-contained. You'd use it for a complex data table with its own sorting, filtering, and pagination state; or a multi-step form wizard; or any component where the state is only relevant to that component. It avoids polluting global NgRx state with view-level state, and the component's state is automatically cleaned up when the component is destroyed."*

---

## Q: "What is the Facade pattern in NgRx and why does large team architecture benefit from it?"

> *"The Facade is a service that abstracts all NgRx interactions — dispatching actions, selecting state — behind a clean API. Components consume the Facade instead of directly importing Store, Actions, and Selectors. This has several architectural benefits. First, components don't have NgRx knowledge — if you swap NgRx for Signals or a different state management solution in the future, you modify the Facade only, not every component. Second, testing components becomes trivial — you mock the Facade with jasmine.createSpyObj instead of setting up an entire NgRx store. Third, the Facade becomes the documented public API of a feature's state — other teams know exactly what operations are allowed without digging into action files. Fourth, the Facade can aggregate multiple store slices and derived logic, presenting a clean, simplified view of complex state to each component."*

---

## Q: "How does Redux DevTools help you debug NgRx applications?"

> *"Redux DevTools is a browser extension that connects to NgRx and gives you a real-time log of every action dispatched and the state before and after each one. This is extraordinarily useful: if a bug report says 'the cart shows the wrong count', you can replay a session and inspect exactly which action changed cart state and what the reducer returned. Time-travel debugging lets you step backward and forward through the action history — the UI rehydrates to each state. You can also dispatch actions manually, test edge cases in production state, and export/import state snapshots for reproduction. It transforms debugging from 'add console.logs everywhere and guess' to 'see exactly what happened and when'. In a large team, consistent action naming conventions — like [Cart] Add Item, [Cart] Remove Item — make the DevTools action log readable and searchable."*
