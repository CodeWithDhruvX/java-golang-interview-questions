# Redux Interview Questions & Answers

## 🔹 1. Redux Basics & Core Concepts (Questions 1-20)

**Q1: What is Redux and why is it used?**
A predictable state container for JavaScript apps. It helps manage global state in a consistent way across different environments (client/server) and makes testing/debugging easier.

**Q2: What are the three core principles of Redux?**
1. Single Source of Truth (One Store). 2. State is Read-Only (Immutable). 3. Changes are made with Pure Functions (Reducers).

**Q3: What is the single source of truth?**
The entire application state is stored in an object tree within a single store. It makes debugging and state inspection easier.

**Q4: What is a Redux Store?**
The object that holds the application state. It provides `getState()`, `dispatch(action)`, and `subscribe(listener)` methods.

**Q5: What are Actions in Redux?**
Plain JavaScript objects that describe "what happened". Must have a `type` property and optionally a `payload` with data.

**Q6: What is a Reducer?**
A pure function `(previousState, action) => newState` that determines how the state changes in response to an action.

**Q7: What is the difference between Redux and React's Context API?**
Context is built-in and good for low-frequency updates (theme/user). Redux provides dev-tools, middleware support, and better performance for high-frequency updates.

**Q8: Why are reducers called "pure functions"?**
Because they must not have side effects, mutate arguments, or call non-pure functions (like `Date.now()`). Same input always returns same output.

**Q9: What is the significance of immutability in Redux?**
Redux compares objects by reference (shallow equality). If you mutate state directly, Redux won't detect changes and won't re-render components.

**Q10: How does data flow in a Redux application (Unidirectional Data Flow)?**
Action Dispatched -> Store calls Reducer -> Reducer returns New State -> Store updates -> UI re-renders.

**Q11: What is dispatching an action?**
Sending an action object to the Redux store using `store.dispatch(action)` to trigger a state update.

**Q12: What is the difference between state and props in the context of Redux?**
Redux state is global and stored in the Store. Props are how that state is passed down to React components (`mapStateToProps`).

**Q13: Can we have multiple stores in a Redux application?**
Technically yes, but the Redux pattern enforces a single store. For module isolation, use `combineReducers` instead of multiple stores.

**Q14: What is the Flux architecture and how does Redux relate to it?**
Flux is a pattern with Dispatcher, Store, and Views. Redux is an implementation of Flux but simplifies it by having a single Store and no Dispatcher.

**Q15: What are the main components of Redux?**
Action (Event), Reducer (Logic), Store (State holder), and View (UI).

**Q16: How do you access the store state?**
Using `store.getState()` directly, or via `useSelector` hook in React components.

**Q17: What is `subscribe()` in Redux?**
A method (`store.subscribe(listener)`) that registers a callback function to run whenever an action is dispatched and the state tree might have changed.

**Q18: What is a payload in a Redux action?**
The property of the action object (usually `action.payload`) that contains the data needed to update the state.

**Q19: What is an Action Creator?**
A function that creates and returns an action object. `const add = (text) => ({ type: 'ADD_TODO', text })`.

**Q20: Why do we need to return a new state object in reducers?**
To respect immutability. Returning a mutated state object will not trigger reference checks, causing UI to not update.

---

## 🔹 2. React-Redux Integration (Questions 21-34)

**Q21: What is the `react-redux` library?**
The official React binding for Redux. It allows React components to read data from a Redux Store and dispatch actions.

**Q22: What is the `<Provider>` component?**
A wrapper component from `react-redux` that makes the Redux store available to any nested components via React Context.

**Q23: What is the `connect()` function?**
A Higher-Order Component (HOC) that connects a React component to the Redux store. (Legacy, hooks are preferred now).

**Q24: What is `mapStateToProps`?**
A function used with `connect()` that maps parts of the Redux state to the props of a React component.

**Q25: What is `mapDispatchToProps`?**
A function used with `connect()` that maps action dispatch functions to the props of a React component.

**Q26: What is the use of `ownProps` in `mapStateToProps`?**
The second argument to `mapStateToProps`, representing the props passed to the wrapper component itself. Allows state selection based on props.

**Q27: How does Redux integration affect React component re-rendering?**
React-Redux components subscribe to the store. They only re-render if the specific part of the state they select changes (Performant by default).

**Q28: What are React hooks for Redux (`useSelector`, `useDispatch`)?**
`useSelector`: Extracts data from the Redux store state. `useDispatch`: Returns the dispatch function to send actions.

**Q29: How is `useSelector` different from `mapStateToProps`?**
`useSelector` uses strict reference equality (`===`) by default, whereas `connect` uses shallow equality. `useSelector` can return any value, not just an object.

**Q30: How do you avoid unnecessary re-renders when using `useSelector`?**
Select the smallest possible slice of state, or use `shallowEqual` as the second argument, or use Reselect library.

**Q31: What is the `batch` API in React-Redux?**
Allows grouping multiple dispatches into a single render pass. (Note: React 18 does automatic batching, so this is less needed).

**Q32: How do you handle component state vs Redux state?**
Use Local State (`useState`) for UI-only transient data (modals, form inputs). Use Redux for global data needed across components (Auth, Data cache).

**Q33: Can you use Redux with class components?**
Yes, primarily using the `connect()` HOC. Hooks are not supported in class components.

**Q34: What are Container vs Presentation components?**
Pattern where Container handles Redux logic/fetching (Smart), and Presentation handles rendering UI via props (Dumb).

---

## 🔹 3. Redux Middleware & Side Effects (Questions 35-47)

**Q35: What is Redux Middleware?**
A wrapper around `store.dispatch` that allows intercepting actions before they reach the reducer.

**Q36: Why do we need middleware in Redux?**
Redux is synchronous by default. Middleware enables async Logic (API calls), logging, crash reporting, and routing.

**Q37: What is Redux Thunk?**
Standard middleware for async logic. It allows action creators to return a function `(dispatch, getState) => {}` instead of an action object.

**Q38: How do you handle async operations in Redux?**
Using middleware like Redux Thunk, Redux Saga, or RTK Query. Dispatch a 'REQUEST' action, call API, then dispatch 'SUCCESS' or 'FAILURE'.

**Q39: What is Redux Saga?**
A middleware that uses JS Generators (`function*`) to handle side effects. It makes async flows easy to test and read (looks synchronous).

**Q40: What is the difference between Redux Thunk and Redux Saga?**
Thunk: Simple, functions returns functions. Good for simple promises. Saga: Complex, uses Generators/Effects. Good for complex flows (cancellation, race conditions).

**Q41: What is a Generator function (used in Sagas)?**
A function utilizing `yield` to pause and resume execution. Syntax: `function* mySaga() { yield ... }`.

**Q42: What are "Effects" in Redux Saga?**
Plain JavaScript objects containing instructions for the middleware (e.g., `call`, `put`). They are declarative.

**Q43: What is `call`, `put`, and `takeEvery` in Redux Saga?**
`call`: Run a Promise/Function. `put`: Dispatch an action. `takeEvery`: Listen for every dispatched action of a type.

**Q44: What is Redux Observable?**
Middleware based on RxJS streams (Observables). Actions are streams of events. Powerful for complex async compositions.

**Q45: How do you log actions using middleware (e.g., `redux-logger`)?**
Middleware that logs the previous state, the action dispatched, and the next state to the console for debugging.

**Q46: Can you write your own custom middleware?**
Yes. Pattern: `store => next => action => { ... next(action) }`.

**Q47: How does the control flow pass through middlewares?**
Middlewares form a chain. Each middleware processes the action and calls `next(action)` to pass it to the next one, finally reaching the reducer.

---

## 🔹 4. Redux Toolkit (Modern Redux) (Questions 48-59)

**Q48: What is Redux Toolkit (RTK) and why is it recommended?**
The official, opinionated toolset for efficient Redux development. It simplifies configuration, reduces boilerplate, and includes Immer/Thunk by default.

**Q49: What is `configureStore` in RTK?**
Replaces `createStore`. Automatically sets up the store with good defaults (DevTools, Thunk, Slice reducers).

**Q50: What is `createSlice`?**
A function that accepts an initial state, an object of reducer functions, and a "slice name". It automatically generates Action Creators and Action Types.

**Q51: How does `createSlice` handle immutability (Immer.js)?**
It uses Immer library internally, allowing you to write "mutating" syntax (`state.value = 123`) which is safely converted to immutable updates.

**Q52: What is `createAsyncThunk`?**
A helper to generate thunks that dispatch `pending/fulfilled/rejected` action types based on a promise.

**Q53: What are the "extraReducers"?**
A field in `createSlice` to listen for actions defined *outside* the slice (e.g., thunks generated by `createAsyncThunk` or actions from other slices).

**Q54: What is `createEntityAdapter`?**
A utility to generate Reducers and Selectors for normalized data (managing collections of items with IDs).

**Q55: How do you refactor vanilla Redux to Redux Toolkit?**
Replace `createStore` with `configureStore`. Replace separate Action/Reducer files with `createSlice`. Use `createAsyncThunk` for async.

**Q56: What is RTK Query?**
A data fetching and caching tool included in Redux Toolkit. It eliminates the need for manual thunks/reducers for data fetching.

**Q57: What are the benefits of using RTK Query over `useEffect` data fetching?**
Automatic caching, deduplication, polling, invalidation, and easy loading/error state management.

**Q58: How do you handle caching in RTK Query?**
It stores query results in Redux state normalized. Cache is identified by endpoint name + arguments.

**Q59: What is "mutating" logic in Redux Toolkit reducers?**
Writing `state.todos.push(todo)` instead of `return [...state.todos, todo]`. Allowed because Immer handles the draft state.

---

## 🔹 5. Advanced & Scenarios (Questions 60-75)

**Q60: What are High-Order Reducers?**
A function that takes a reducer and returns a new reducer with added functionality (e.g., undo/redo, resetting state on logout).

**Q61: How do you structure a large Redux application?**
"Feature Folder" approach: Group Redux logic (slice, thunks) with component files by feature (`/features/auth`, `/features/posts`).

**Q62: What is Redux DevTools Extension?**
A browser extension that allows you to inspect every action, state change, and "time travel" by jumping back to previous states.

**Q63: What is Time Travel Debugging?**
The ability to move back and forth through the history of dispatched actions and see the UI state at that point in time.

**Q64: How do you persist Redux state (redux-persist)?**
Using `redux-persist` library to save the Redux store (or parts of it) to `localStorage` or `sessionStorage` and rehydrate it on app launch.

**Q65: What are Selectors and why use them?**
Functions that extract specific pieces of data from the store. They encapsulate state shape and allow memoization.

**Q66: What is the Reproxy pattern (Reselect library)?**
Creating memoized selectors. `createSelector([input], output)`. Recalculates only if input selectors change.

**Q67: How do you optimize Redux performance?**
Use memoized selectors (Reselect), batch updates, avoid storing large objects/Blobs in state, and normalize state shape.

**Q68: What are normalization and why normalize state?**
Structuring state like a database (IDs as keys). Prevents deeply nested updates and data duplication.

**Q69: How do you handle form state in Redux (Redux Form vs Formik)?**
Avoid storing every keystroke in Redux (performance). Use local state or dedicated libraries like React Hook Form or Formik.

**Q70: Is Redux dead? When should you NOT use Redux?**
No, but check if needed. Don't use for: Small apps, simple state, or when server-state tools (React Query) cover the data fetching needs.

**Q71: How do you test Redux reducers and actions?**
Test reducers as pure functions: `expect(reducer(initialState, action)).toEqual(newState)`. Dispatch actions to mock store for thunks.

**Q72: What is server-side rendering (SSR) with Redux?**
Create a new store instance on every request, preload data, pass initial state to client, and "hydrate" the client-side store.

**Q73: How to split code (lazy load) with Redux reducers?**
Use `replaceReducer` to dynamically add reducers when a chunk (route/feature) is loaded.

**Q74: What is the concept of "Duck" file structure?**
Bundling Reducers, Action Types, and Action Creators for a feature into a single file (module).

**Q75: Difference between `compose` and `applyMiddleware`?**
`applyMiddleware` adds middleware. `compose` combines multiple store enhancers (like Middleware + DevTools) into one function.
