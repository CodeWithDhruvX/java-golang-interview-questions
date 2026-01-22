## 🔹 31. React + JavaScript Deep Dive

### Question 301: How does JavaScript closure affect React hooks?

**Answer:**
Hooks (like `useEffect` or `useCallback`) rely on closures to capture variables from the component's scope.
**Issue:** If not managed correctly (missing dependencies), the closure captures a specific "frame" of variables (stale variables) and holds onto them even when the component re-renders with new values.

---

### Question 302: How does event loop impact React rendering?

**Answer:**
JavaScript is single-threaded. React's rendering work (diffing/reconciliation) happens on the call stack.
If rendering takes too long (>16ms), it blocks the stack, preventing the Event Loop from processing other tasks (like click events or painting), causing UI freeze.
**React 18 Solution:** Concurrent features yield to the event loop.

---

### Question 303: How does React interact with the browser paint cycle?

**Answer:**
React manages the updates in memory (Virtual DOM).
1.  **Sync:** React calculates changes.
2.  **Commit:** React modifies DOM.
3.  **Browser:** Browser sees DOM change -> Recalculate Layout -> Paint pixels.
`useLayoutEffect` blocks the paint (runs between step 2 and 3). `useEffect` runs after paint (Step 3).

---

### Question 304: How does garbage collection affect React memory usage?

**Answer:**
Closures used in event handlers or effects can keep large objects alive in memory.
If you detach a large DOM tree but keep a reference to one node in JS (e.g., via a Ref), the entire tree cannot be Garbage Collected (Memory Leak).

---

### Question 305: How does React handle microtasks vs macrotasks?

**Answer:**
*   **Macrotasks:** `setTimeout`, user events.
*   **Microtasks:** `Promise.then`, `queueMicrotask`.
State updates triggered in microtasks are usually processed immediately after the current script execution, ensuring the DOM updates before the browser paints.

---

### Question 306: How does React work with Intersection Observer?

**Answer:**
Intersection Observer API detects when an element enters the viewport.
In React:
1.  Attach a `ref` to the target element.
2.  In `useEffect`, create `new IntersectionObserver(callback)`.
3.  `observer.observe(ref.current)`.
4.  In callback, `setState` (e.g., `setIsVisible(true)`).

---

### Question 307: How does Resize Observer help React apps?

**Answer:**
It allows a component to respond to changes in **its own size**, not just the window size.
Essential for "Container Queries" logic in JavaScript before they were available in CSS.
Used in responsive charts/grids.

---

### Question 308: How does requestIdleCallback improve performance?

**Answer:**
`window.requestIdleCallback(fn)` allows you to schedule low-priority background work (e.g., sending analytics logs, pre-fetching data) to run only when the browser's main thread is **idle**, ensuring critical interactions aren't delayed.

---

### Question 309: How does React behave in low-memory devices?

**Answer:**
React creates many objects (Fiber nodes, Virtual DOM objects, Synthetic Events). On low-memory devices, this GC pressure can cause stutter.
**Optimization:** Use `production` build (strips dev warnings), virtualization, and avoid creating objects inside render loop.

---

### Question 310: How do you optimize React for slow networks?

**Answer:**
1.  **Code Splitting:** Send smaller JS bundles.
2.  **Service Workers:** Cache assets/API responses.
3.  **Optimistic UI:** Show feedback instantly.
4.  **Skeleton Screens:** Improve perceived performance.

---

## 🔹 32. Advanced Internals & Systems

### Question 311: How does React prioritize updates internally?

**Answer:**
React Fiber uses a priority / expiration time model.
1.  **Immediate:** User Input (Expires in 0ms).
2.  **User Blocking:** Animations (250ms).
3.  **Normal:** Data fetch (5000ms).
4.  **Low:** Log / Analytics (10000ms).
Scheduler processes tasks based on these priorities.

---

### Question 312: What is lane-based scheduling in React?

**Answer:**
An improvement over Expiration Time (React 17+).
Instead of a single number (time), updates are assigned to **Lanes** (Bitmask, e.g., 0b0001).
Multiple lanes can be processed together. High priority lanes (Input) are processed first.

---

### Question 313: How does React Fiber manage interruption?

**Answer:**
It uses a `deadline` object (via `scheduler` package).
Periodically (every few milliseconds) during the specific render loop, Fiber checks `deadline.timeRemaining()`. If time is up, it saves the current state of the tree (fiber node) and returns control to the browser.

---

### Question 314: How does React batch updates across async boundaries?

**Answer:**
In React 18, `createRoot` wraps the entire application update process throughout the event loop tick. Whether the update comes from a timeout, promise, or native event, React flags it and processes all flagged updates in a single pass at the end of the tick.

---

### Question 315: How does concurrent rendering affect consistency?**

**Answer:**
It renders "Off-Screen" (in memory).
React builds the new Fiber tree in the background. The user sees the old screen. Only when the new tree is fully ready and consistent does React swap the pointers (Double Buffering technique), so the user never sees a half-rendered UI.

---

### Question 316: How would you design a design-system in React?

**Answer:**
1.  **Tokens:** Define colors, spacing, fonts (JSON/CSS Vars).
2.  **Primitives:** `<Box>`, `<Text>`, `<Flex>`.
3.  **Components:** `<Button>`, `<Input>` (consuming tokens).
4.  **Documentation:** Storybook.
5.  **Distribution:** NPM package with tree-shaking support.

---

### Question 317: How do you manage shared components across teams?

**Answer:**
1.  **Monorepo:** Source code in one repo, easy to refactor.
2.  **Versioning:** Semantic Versioning.
3.  **Bit:** Tool for isolating and sharing components from any repo.
4.  **Governance:** Owners for the Design System library to review changes.

---

### Question 318: How do you version frontend libraries?

**Answer:**
**SemVer (Semantic Versioning):** `v1.2.3` (Major.Minor.Patch).
*   **Patch:** Bug fix (Safe).
*   **Minor:** New feature (Safe).
*   **Major:** Breaking change (Requires code update).

---

### Question 319: How do you manage breaking changes?

**Answer:**
1.  **Deprecation:** Mark prop as deprecated in v1 (`console.warn`).
2.  **Support:** Support both old and new way in v1.
3.  **Removal:** Remove old way in v2.
4.  **Codemods:** Provide scripts to automate the update for consumers.

---

### Question 320: How do you enforce UI consistency at scale?

**Answer:**
1.  **Linting:** `eslint-plugin-react`.
2.  **Types:** TypeScript.
3.  **Visual Regression Testing:** Tools like Chromatic (for Storybook) take screenshots of every component and compare pixel-by-pixel with the master branch to catch unintended style changes.

---

### Question 321: How does React differ from React Native internally?

**Answer:**
*   **React (Web):** Renders to `react-dom`. Host components are `div`, `span`.
*   **React Native:** Renders to `react-native`. Host components are `View`, `Text`. Uses a **Bridge** (serialized JSON) or **JSI** (Direct C++ call) to communicate with native iOS/Android standard UI components.

---

### Question 322: How does React work with Web Components?**

**Answer:**
React can render Web Components (`<custom-element>`).
**Challenge:** React passes data via attributes (strings) and synthetic events.
**React 19:** Adds full support for Custom Elements, passing properties (objects) and handling native Custom Events correctly.

---

### Question 323: How does React integrate with WebAssembly?

**Answer:**
WebAssembly (Wasm) runs binary code (C++/Rust) in browser.
React handles the UI. For heavy calc (video processing), React calls Wasm function.
Wasm returns result -> React updates State -> UI re-renders.

---

### Question 324: What is React Server Components (RSC)?

**Answer:**
A paradigm shift (Next.js App Router).
Components run **exclusively** on the server. They have direct access to DB/Filesystem.
They send **serialized UI slots** to the client, not JS code.
**Result:** Zero bundle size for these components.

---

### Question 325: How does React compare with signals-based frameworks (Solid/Preact)?

**Answer:**
*   **React:** **Pull-based**. When state changes, React re-executes the component function (Virtual DOM Diffing).
*   **Signals (SolidJS):** **Push-based**. State (Signal) is a dependency graph. When signal changes, it updates *only* the specific DOM text node subscribed to it. No VDOM overhead.
