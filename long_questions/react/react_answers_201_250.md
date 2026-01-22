## 🔹 21. React 18 & Modern React

### Question 201: What is `createRoot` in React 18?

**Answer:**
`createRoot` is the new API introduced in React 18 to initialize a React application. It replaces `ReactDOM.render`. It is **required** to opt-in to new concurrent features.

**Example:**
```jsx
// Old (v17)
ReactDOM.render(<App />, document.getElementById('root'));

// New (v18)
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
```

---

### Question 202: Difference between `render` and `createRoot`?

**Answer:**
*   **Legacy `render`:** Updates are synchronous (blocking). Does not support Concurrent features.
*   **`createRoot`:** Enables **Concurrent Mode** by default. Supports features like Automatic Batching, Transitions, and Suspense on server.

---

### Question 203: What is concurrent features opt-in?

**Answer:**
In React 18, upgrading to `createRoot` doesn't automatically make rendering concurrent everywhere. Concurrent rendering is enabled **only** for updates triggered by concurrent features (like `startTransition`, `useDeferredValue`). This allows gradual migration.

---

### Question 204: What is `startTransition`?

**Answer:**
An API that lets you mark some state updates as **non-urgent**.
*   **Urgent:** Typing, Clicking, Hovering.
*   **Transition:** Updating a list, rendering a chart.
If a user types while a transition is running, React interrupts the transition to handle the keystroke.

**Example:**
```jsx
startTransition(() => {
  setFilteredList(newList);
});
```

---

### Question 205: When should you use `useTransition`?

**Answer:**
When you have a UI update that:
1.  Is computationally expensive (lags the UI).
2.  Should not block user interaction (like typing in a search input).
It allows React to deprioritize that specific update.

---

### Question 206: What problem does `useDeferredValue` solve?

**Answer:**
It is similar to debouncing/throttling but integrated with React's render cycle. It defers re-rendering a part of the UI tree based on a changing value until "more important" updates are done.

**Example:**
```jsx
const deferredQuery = useDeferredValue(query);
// List depends on deferredQuery (low priority)
// Input depends on query (high priority)
```

---

### Question 207: How does React handle priority updates?

**Answer:**
React Fiber uses a system of **Lanes** (bitmask).
1.  **Sync Lane:** User input (highest).
2.  **Transition Lane:** `startTransition` (lower).
3.  **Default Lane:** `useEffect` / `fetch`.
High priority work can interrupt low priority work.

---

### Question 208: What are blocking vs non-blocking updates?

**Answer:**
*   **Blocking (Sync):** Once React starts rendering, it cannot stop. If the tree is huge, the browser freezes until functionality completes.
*   **Non-Blocking (Concurrent):** React yields to the browser loop periodically (every 5ms) to check for input, keeping the page responsive.

---

### Question 209: What changes were introduced in React 18?

**Answer:**
1.  `createRoot` API.
2.  Automatic Batching (for Promises/Timeouts).
3.  Transitions (`useTransition`, `startTransition`).
4.  Suspense on Server (Streaming HTML).
5.  New Hooks: `useId`, `useDeferredValue`, `useSyncExternalStore`.

---

### Question 210: How does React 18 improve performance?

**Answer:**
Mainly through **Automatic Batching** (fewer renders) and **Concurrent Rendering** (better perceived performance/responsiveness). It doesn't necessarily make the computation *faster*, but it makes the user *experience* smoother by prioritizing interaction.

---

## 🔹 22. State & Data Flow (Advanced)

### Question 211: What is derived state?

**Answer:**
State that acts as a cache for data that can be calculated from props or other state variables.
**Bad:** Storing `fullName` in state when you have `firstName` and `lastName`.
**Good:** `const fullName = firstName + ' ' + lastName;` (Just a variable).

---

### Question 212: Why is derived state considered an anti-pattern?

**Answer:**
It violates the **Single Source of Truth** principle. If you duplicate data (Prop -> State), you must manually keep them in sync using `useEffect`, which is bug-prone and complex. Only store the minimal required state.

---

### Question 213: How do you lift state up?

**Answer:**
If two sibling components need access to the same state:
1.  Remove state from both children.
2.  Move it to their closest common parent.
3.  Pass state down via props to both children.

---

### Question 214: How do you share state between sibling components?

**Answer:**
See Q213 (Lifting State Up).
If components are far apart in the tree, lifting state becomes cumbersome (Props Drilling), so we use **Context** or a **Global Store**.

---

### Question 215: What is single source of truth?

**Answer:**
The philosophy that for any piece of data, there should be exactly one place code where that data lives and is modified. All other places should just read it. This prevents "state divergence" bugs.

---

### Question 216: How do you normalize state?

**Answer:**
Structuring state like a database. Instead of an array:
`[{ id: 1, name: 'A' }]`
Use an object dictionary:
`{ byId: { 1: { id: 1, name: 'A' } }, allIds: [1] }`
**Benefits:** O(1) Access, easier updates without iterating arrays.

---

### Question 217: How do you manage deeply nested state?

**Answer:**
1.  **Normalization:** Flatten the structure.
2.  **Immer:** Library allowing you to write "mutable" code (`state.a.b.c = 1`) that produces immutable updates.
3.  **useReducer:** Move update logic into a reducer to keep the component clean.

---

### Question 218: What are atomic state libraries?

**Answer:**
Libraries like **Recoil** or **Jotai**.
They break global state into small units called **Atoms**. Components subscribe only to the specific atoms they need.
**Benefit:** Fine-grained re-renders (unlike Context where the whole value updates).

---

### Question 219: What is state colocation?

**Answer:**
Placing state as close as possible to where it is relevant.
Don't put a "Modal Open" boolean in Redux Global Store if only one specific component uses it. Keep it local (`useState`). This makes code easier to delete/refactor.

---

### Question 220: When should state be local vs global?

**Answer:**
*   **Local:** UI state (IsDropdownOpen), Form inputs, specific widget data.
*   **Global:** User Auth, Theme, Shopping Cart, Data used by many unrelated components (Server Cache).

---

## 🔹 23. Side Effects & Async Logic

### Question 221: How do you handle async logic in React?

**Answer:**
1.  **useEffect:** Call async function inside.
2.  **Redux Middleware:** Thunk / Saga.
3.  **Data Fetching Libraries:** React Query / SWR (Preferred for Server State).

---

### Question 222: What problems can async effects cause?

**Answer:**
1.  **Race Conditions:** Request A starts, Request B starts. B finishes. A finishes and overwrites B (wrong old data).
2.  **Memory Leaks:** Updating state on a component that has already unmounted.

---

### Question 223: How do you avoid race conditions in `useEffect`?

**Answer:**
Use a boolean flag (closure) pattern.

**Example:**
```jsx
useEffect(() => {
  let active = true;
  fetchData().then(data => {
    if (active) setData(data);
  });
  return () => { active = false; }; // Cleanup
}, [id]);
```

---

### Question 224: How do you handle polling in React?

**Answer:**
Set up a `setInterval` in `useEffect`. ensure you clear it in the return function.
Or use libraries: `useQuery(key, fn, { refetchInterval: 5000 })`.

---

### Question 225: How do you handle WebSockets in React?**

**Answer:**
Initialize the socket connection in a `useEffect` on Mount.
Add event listeners.
Close connection in cleanup.
**Note:** Often better to put the socket instance in a Context or Ref so it persists across re-renders without reconnecting.

---

### Question 226: How do you handle long-running tasks?

**Answer:**
JavaScript is single-threaded. Long tasks freeze the UI.
1.  **Web Workers:** Offload calculation to a background thread.
2.  **Time Slicing / scheduler:** Break task into small `setTimeout` chunks.
3.  **React Concurrent:** `useTransition` / `useDeferredValue`.

---

### Question 227: How do you integrate background sync?

**Answer:**
This is mostly a browser/PWA feature (Service Workers).
React can register the Service Worker (`navigator.serviceWorker.register`).
The Service Worker handles syncing when the network returns.

---

### Question 228: How do you retry failed API calls?

**Answer:**
Implement recursive retry logic with exponential backoff.
Or (Simplest): Use **React Query** (it retries 3 times by default).

---

### Question 229: How do you handle optimistic UI updates?

**Answer:**
The pattern of updating the UI *before* the server responds (assuming success).
1.  User clicks "Like".
2.  `setLikes(n => n + 1)` immediately.
3.  Send API request.
4.  If API fails, rollback `setLikes(n => n - 1)` and show error.

---

### Question 230: How do you roll back optimistic updates?

**Answer:**
You need to store the **Previous State** before the update.
In React Query context: `onMutate` (save snapshot) -> `onError` (restore snapshot).

---

## 🔹 24. Refs & DOM Manipulation

### Question 231: When should you access the DOM directly?

**Answer:**
React allows it via Refs, but it should be an "escape hatch".
Valid cases:
1.  Managing input focus/text selection.
2.  Triggering imperative animations.
3.  Integrating with third-party DOM libraries (D3, Google Maps).

---

### Question 232: How do callback refs work?

**Answer:**
Instead of passing a ref object, you pass a function: `ref={node => this.node = node}`.
React calls this function with the DOM element when it mounts, and with `null` when it unmounts. This is useful if you need to execute code *immediately* upon attachment.

---

### Question 233: Difference between object refs and callback refs?

**Answer:**
*   **Object Ref (`useRef`):** Standard. Container `{ current: null }`. React sets `current` during commit phase.
*   **Callback Ref:** Function. React calls it during commit phase. Gives you notification *when* the node is set.

---

### Question 234: How do you measure element size in React?

**Answer:**
Use a **Callback Ref** to measure the node when it mounts, or `useLayoutEffect` to measure post-render.
`node.getBoundingClientRect()` gives width/height.

---

### Question 235: How do you detect outside clicks?

**Answer:**
Attach a `click` event listener to `document` in `useEffect`. Check if the click target is contained inside your ref.

**Example:**
```jsx
useEffect(() => {
  const listener = (e) => {
    if (ref.current && !ref.current.contains(e.target)) {
      closeModal();
    }
  };
  document.addEventListener('mousedown', listener);
  return () => document.removeEventListener('mousedown', listener);
}, []);
```

---

### Question 236: How do you implement infinite scrolling?

**Answer:**
Use the **Intersection Observer API**.
Place a hidden "sentinel" div at the bottom of the list.
When the sentinel intersects with the viewport (becomes visible), trigger the "Load More" function.

---

### Question 237: How do you manage scroll restoration?

**Answer:**
Browsers handle this for static pages, but SPAs struggle.
1.  **React Router:** `<ScrollRestoration />` component (v6.4+).
2.  **Manual:** Save `window.scrollY` in `sessionStorage` on unmount. Restore it in `useLayoutEffect` on mount.

---

### Question 238: How do you handle animations using refs?

**Answer:**
Libraries like **GSAP** (GreenSock) need direct DOM access to animate properties performantly (bypassing React render cycle).
`gsap.to(ref.current, { x: 100 })`.

---

### Question 239: How do you integrate third-party DOM libraries?

**Answer:**
Wrap the library in a React Component.
*   **Mount:** Initialize library in `useEffect` (using ref).
*   **Update:** Update library instance in `useEffect` when props change.
*   **Unmount:** Destroy/Cleanup library instance in return function.

---

### Question 240: Why should direct DOM manipulation be minimized?

**Answer:**
React maintains an internal representation (Virtual DOM). If you change the Real DOM manually (e.g., remove a child), React's Virtual DOM gets out of sync. Next time React tries to update that node, it might crash with "Node not found" errors.

---

## 🔹 25. Styling in React

### Question 241: Different ways to style React components?

**Answer:**
1.  **Inline Styles:** `style={{ color: 'red' }}`.
2.  **CSS Class:** `className="btn"`.
3.  **CSS Modules:** `import styles from './App.module.css'`.
4.  **CSS-in-JS:** Styled Components, Emotion.
5.  **Utility Classes:** Tailwind CSS.

---

### Question 242: What are CSS Modules?

**Answer:**
A build-step feature where CSS class names are scoped locally by default.
`styles.button` compiles to something like `Button_button__1a2b3`.
**Benefit:** Prevents global namespace collisions.

---

### Question 243: What is styled-components?

**Answer:**
A popular CSS-in-JS library. It allows you to write actual CSS code to style your components. It also removes the mapping between components and styles—using components as a low-level styling construct.

**Example:**
```jsx
const Button = styled.button`
  background: ${props => props.primary ? "blue" : "white"};
`;
```

---

### Question 244: Difference between CSS Modules and styled-components?

**Answer:**
*   **CSS Modules:** You write CSS in `.css` files. JS imports them. Separation of concerns.
*   **Styled Components:** You write CSS in `.js` files. CSS is tightly coupled with Logic/Props.

---

### Question 245: What is inline styling limitation?**

**Answer:**
1.  No Media Queries (`@media`).
2.  No Pseudo-classes (`:hover`, `:active`).
3.  No Keyframe animations.
4.  Performance (New object created every render).

---

### Question 246: How do you apply dynamic styles?

**Answer:**
1.  **Classes:** Template literal string. `` className={`btn ${isActive ? 'active' : ''}`} ``.
2.  **Styled Components:** Pass props to the styled component.
3.  **CSS Variables:** Update `--main-color` on the root element.

---

### Question 247: How do you handle theming in React?

**Answer:**
Use a **ThemeProvider** (Context).
It wraps the app and passes a `theme` object (colors, fonts) down. Styled Components / Material UI consume this theme in child components automatically.

---

### Question 248: How do you manage dark mode?

**Answer:**
1.  Toggle a specific class (e.g., `dark`) on the `<body>` or root `<div>`.
2.  Use **CSS Variables** (`--bg-color`, `--text-color`) and redefine them under the `.dark` class selector.

---

### Question 249: How do you avoid CSS conflicts?

**Answer:**
1.  **Scoping:** Use CSS Modules or Shadow DOM.
2.  **Naming Convention:** BEM (Block Element Modifier).
3.  **CSS-in-JS:** Generates unique hashes for classes.

---

### Question 250: What is critical CSS?

**Answer:**
A performance technique where you extract and **inline** the CSS required for the "Above the Fold" content directly in the HTML `<head>`. This prevents the "Flash of Unstyled Content" (FOUC) and improves render speed.
