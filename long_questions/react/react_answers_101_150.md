## 🔹 11. Core React Internals

### Question 101: What is Fiber in React?

**Answer:**
Fiber is the complete rewrite of the React core reconciliation algorithm, introduced in React 16. It is designed to enable **incremental rendering**: the ability to split rendering work into chunks and spread it out over multiple frames.

---

### Question 102: What problem does React Fiber solve?

**Answer:**
The old reconciliation algorithm (Stack) worked recursively and synchronously. If a component tree was deep, the main thread would be blocked while rendering, causing unresponsive UI ("jank"). Fiber allows React to:
1.  Pause work and come back to it later.
2.  Assign priority to different types of updates.
3.  Reuse previously completed work.
4.  Abort work if it's no longer needed.

---

### Question 103: What is concurrent rendering?

**Answer:**
Concurrent rendering (enabled in React 18) describes the new capabilities allowing React to prepare multiple versions of the UI at the same time. While it's an implementation detail using Fiber, it allows features like **interruptible rendering**, meaning React can start rendering an update, pause to handle a user click, and then resume.

---

### Question 104: What is time slicing in React?

**Answer:**
Time slicing is a technique used by Fiber to split high-priority work (like user input) and low-priority work (like network data rendering). React allocates "slices" of time to rendering tasks. If a task exceeds the frame budget (approx 5ms), it yields control back to the browser to handle events (scrolling, typing), ensuring the app remains responsive.

---

### Question 105: What is automatic batching in React 18?

**Answer:**
Prior to React 18, React only batched state updates that occurred inside React event handlers. Updates in Promises, timeouts, or native event handlers were not batched (causing multiple renders).
React 18 batches **all** state updates automatically, regardless of where they originate.

**Example:**
```jsx
setTimeout(() => {
  setCount(c => c + 1);
  setFlag(f => !f);
  // React 18: 1 Re-render
  // React 17: 2 Re-renders
}, 1000);
```

---

### Question 106: What is Strict Mode in React?

**Answer:**
`<React.StrictMode>` is a tool for highlighting potential problems in an application. It does not render any visible UI. It activates additional checks and warnings for its descendants.
It runs in **Development Mode** only.

---

### Question 107: Why does React Strict Mode render components twice?

**Answer:**
It intentionally double-invokes the following to stress-test the component for side effects and impurity:
*   Class component constructor, render, and `shouldComponentUpdate`.
*   Functional component body.
*   State updater functions.
*   `useMemo` functions.
If a function is impure (has side effects), running it twice produces different/bad results, making bugs obvious.

---

### Question 108: What is reconciliation algorithm?

**Answer:**
Reconciliation is the process through which React updates the DOM. When a component's state or props change, React decides whether an actual DOM update is necessary by comparing the newly returned element with the previously rendered one.
*   **Diffing:** Comparing trees.
*   **Commit:** Applying changes.

---

### Question 109: How does React diffing work?

**Answer:**
1.  **Different Element Types:** If the root elements have different types, React tears down the old tree and builds the new tree from scratch.
2.  **Same Element Types:** React looks at the attributes of both, keeps the underlying DOM node, and only updates the changed attributes.
3.  **Recursion on Children:** React recurses on children. Keys are used to match children efficiently.

---

### Question 110: What is the render phase vs commit phase?

**Answer:**
*   **Render Phase:** React determines what changes need to be made (comparing VDOM). This phase is **pure** and has no side effects. It can be paused, aborted, or restarted by React (Fiber).
*   **Commit Phase:** React applies the changes to the Real DOM. This phase is **impure** (DOM updates, lifecycle methods `componentDidMount/Update` run here). It cannot be interrupted.

---

## 🔹 12. Rendering Behavior

### Question 111: What triggers a re-render in React?

**Answer:**
1.  **State Update:** Calling `setState` (or `useState` updater).
2.  **Props Update:** Parent component passes new props.
3.  **Parent Re-render:** If a parent re-renders, all children re-render (unless memoized).
4.  **Context Update:** Any component consuming a Context re-renders when the Provider's value changes.
5.  **Hooks Changes:** Changes in custom hooks that trigger state updates.

---

### Question 112: How does React decide when to re-render a component?**

**Answer:**
By default, React re-renders a component whenever its parent re-renders.
Developers can override this behavior using:
*   `React.memo()` (Function components): Compares Props.
*   `shouldComponentUpdate()` (Class components): Return `false` to skip render.
*   `PureComponent` (Class components): Shallow prop/state comparison.

---

### Question 113: What happens when state is updated with the same value?

**Answer:**
React uses `Object.is` to compare the old state and new state. If they are equal, React **bails out** of the update. It might still call the component function (render phase) but will **not** commit changes to the children or the DOM.

---

### Question 114: Why are state updates asynchronous?

**Answer:**
They are not truly "asynchronous" like Promises, but React **schedules** them. React waits until all code in event handlers has finished running before processing state updates. This allows **Batching** (grouping multiple updates into one re-render) for performance.

---

### Question 115: How does batching of state updates work?

**Answer:**
If you call `setState` 3 times in a click handler, React will not re-render 3 times. Instead, it accumulates these updates and performs a single re-render with the final state.

---

### Question 116: What is stale state?

**Answer:**
Stale state occurs when a closure (e.g., inside `useEffect`, `setTimeout`, or event listener) captures a variable from an old render cycle. Since the function is defined once (or memoized), it keeps referring to the old `count` value even after `count` has updated in subsequent renders.

---

### Question 117: How do you handle stale closures?

**Answer:**
1.  **Dependency Array:** Include the state variable in the `useEffect` or `useCallback` dependency array. `[state]`.
2.  **Functional Updates:** Use the callback form of `setState`. `setCount(prevCount => prevCount + 1)`. The callback always receives the *current* fresh state.
3.  **Refs:** Use `useRef` to store the latest value, as refs are not captured by closures (they are mutable objects).

---

### Question 118: Why should state updates be immutable?

**Answer:**
React relies on shallow comparison (`prevProps !== nextProps`) to detect changes efficiently. If you mutate an object directly (`obj.prop = 'new'`), the reference `obj` remains the same. React won't detect the change, and the component won't re-render.

---

### Question 119: What is referential equality?

**Answer:**
Referential equality checks if two variables point to the **same memory location**.
`{} === {}` is `false` (Different references).
`const a = {}; const b = a; a === b` is `true`.
React hooks (`useEffect`, `useMemo`) rely on this to determine if dependencies have changed.

---

### Question 120: How does React handle immutability?

**Answer:**
React doesn't enforce immutability; it's a developer convention.
*   **Bad:** `state.items.push(newItem); setState(state);`
*   **Good:** `setState({ ...state, items: [...state.items, newItem] });`
Developers use the Spread Operator (`...`) or libraries like **Immer** to handle immutable updates easily.

---

## 🔹 13. JSX & JavaScript Concepts

### Question 121: Why must JSX have a single parent?

**Answer:**
JSX compiles to `React.createElement()`. This function returns **one** JavaScript object (a Node). A function cannot return multiple distinct objects at once; they must be wrapped in a container (like an Array or a wrapping Object/Div).

---

### Question 122: How do expressions work inside JSX?

**Answer:**
You can embed any valid JavaScript expression inside curly braces `{}`. React evaluates the expression and renders the result.
`<p>Total: {2 + 2}</p>` becomes `<p>Total: 4</p>`.

---

### Question 123: Difference between `{}` and `{{}}` in JSX?

**Answer:**
*   `{}`: Denotes the start of a JavaScript expression.
*   `{{}}`: It's just an object literal `{ key: val }` *inside* a JS expression block `{ ... }`. Commonly seen in the `style` prop: `style={{ color: 'red' }}`.

---

### Question 124: How do you add comments in JSX?

**Answer:**
Standard JS comments (`//` or `/* */`) don't work directly as text. You must wrap them in braces: `{/* This is a comment */}`.

---

### Question 125: How do you render HTML safely in React?

**Answer:**
React escapes all strings by default to prevent XSS. If you receive trusted HTML from an API (e.g., a CMS content), you use the `dangerouslySetInnerHTML` prop.

---

### Question 126: What is `dangerouslySetInnerHTML`?

**Answer:**
It is React's replacement for using `innerHTML` in the browser DOM. The name is intentionally scary to warn developers that setting HTML from code is risky (XSS attacks).

**Example:**
```jsx
<div dangerouslySetInnerHTML={{ __html: '<strong>Hello</strong>' }} />
```

---

### Question 127: How does React escape values?

**Answer:**
Before rendering any value in JSX, React escapes it to a string. This helps prevent Cross-Site Scripting (XSS) attacks because it's impossible to inject arbitrary HTML or script tags via user input. (e.g., `<` becomes `&lt;`).

---

### Question 128: How do inline styles work in React?

**Answer:**
Inline styles don't use strings. They use **JavaScript Objects** with camelCased properties.

**Example:**
```jsx
<div style={{ backgroundColor: 'blue', marginTop: '10px' }}></div>
```
(Not `background-color`, `margin-top`).

---

### Question 129: Why are class names written as `className`?

**Answer:**
JSX code translates to JavaScript. In JavaScript, `class` is a reserved keyword (used for defining Classes). To avoid conflict, React uses the DOM property name `className` instead of the HTML attribute `class`.

---

### Question 130: What is spread operator usage in JSX?

**Answer:**
You can pass a whole object's properties as props using the spread syntax `...`.

**Example:**
```jsx
const props = { firstName: 'Ben', lastName: 'Hector' };
<Greeting {...props} />
// Equivalent to:
<Greeting firstName="Ben" lastName="Hector" />
```

---

## 🔹 14. Forms & Input Handling (Advanced)

### Question 131: How do you debounce input in React?

**Answer:**
Debouncing prevents a function from running too often (e.g., Search API call on every keystroke). You can use a custom hook or `useEffect` with `setTimeout` and cleanup.

**Example:**
```jsx
useEffect(() => {
  const handler = setTimeout(() => {
    fetchResults(query);
  }, 500); // Wait 500ms
  return () => clearTimeout(handler); // Cancel if typing continues
}, [query]);
```

---

### Question 132: How do you throttle API calls from inputs?

**Answer:**
Throttling ensures a function runs at most once in a specified time period (e.g., every 100ms). Useful for 'onScroll' or 'onResize' handlers. Often implemented using `lodash.throttle` or `useRef` to track last execution time.

---

### Question 133: How do you manage complex forms?

**Answer:**
For simple forms, controlled components work. For complex forms (multi-step, arrays, huge validation rules), use libraries like **React Hook Form** or **Formik**. These minimize re-renders and abstract validation logic.

---

### Question 134: How do you handle dynamic form fields?**

**Answer:**
Store the fields as an array of objects in state and map over them to render inputs.

**Example:**
```jsx
const [fields, setFields] = useState([{ value: '' }]);

const addField = () => setFields([...fields, { value: '' }]);
```

---

### Question 135: How do you reset a form in React?**

**Answer:**
For controlled components: Reset the state to initial values.
```jsx
setState(initialState);
```
For uncontrolled components (native):
```jsx
formRef.current.reset();
```

---

### Question 136: How do you manage focus in forms?

**Answer:**
Use `useRef` to get a reference to the DOM element and call `.focus()`.

**Example:**
```jsx
const inputRef = useRef(null);
// On button click
inputRef.current.focus();
```

---

### Question 137: How do you auto-focus an input?

**Answer:**
1.  **Prop:** `<input autoFocus />` (Works on mount).
2.  **Effect:**
    ```jsx
    useEffect(() => {
      inputRef.current.focus();
    }, []);
    ```

---

### Question 138: How do you handle keyboard events?

**Answer:**
Use synthetic events: `onKeyDown` (fires when key is pressed down), `onKeyUp` (released), or `onKeyPress` (deprecated).
Check `event.key` (e.g., `'Enter'`, `'Escape'`).

---

### Question 139: How do you handle copy/paste events?

**Answer:**
React provides `onCopy`, `onCut`, and `onPaste` events.
You can prevent them or manipulate data:
```jsx
<input onPaste={(e) => e.preventDefault()} />
```

---

### Question 140: How do you handle validation errors gracefully?

**Answer:**
Maintain an `errors` object in state mapping field names to error messages.
Determine validity:
1.  **On Change:** Validate input as user types.
2.  **On Blur:** Validate when user leaves the field.
3.  **On Submit:** Validate all fields.

---

## 🔹 15. Context API (Advanced)

### Question 141: How does Context API work internally?

**Answer:**
It uses the Publish-Subscribe pattern. The `<Provider>` holds the value and subscribes consumers. When the value prop changes, React recursively traverses down to find all Consumers (`useContext`) and triggers a re-render for them, bypassing intermediate components (avoiding props drilling).

---

### Question 142: What causes Context consumers to re-render?

**Answer:**
Any time the `value` prop passed to `<Context.Provider>` changes (specifically, referential identity change). Even if the values *inside* the object are the same, if a *new object* is created every render, all consumers re-render.

---

### Question 143: How do you optimize Context performance?

**Answer:**
1.  **Memoize Value:** Wrap the context value object in `useMemo`.
    ```jsx
    const value = useMemo(() => ({ user, login }), [user, login]);
    ```
2.  **Split Contexts:** Don't put everything in one big context. Separate `ThemeContext`, `UserContext`, etc. so updates to one don't re-render consumers of the other.

---

### Question 144: When should you NOT use Context?

**Answer:**
Context is NOT optimized for **high-frequency** state updates (e.g., drag-and-drop coordinates, animations, current time). Because every update re-renders all consumers, it can cause performance issues (lag). Use local state or specialized libraries (Recoil/Zustand) for this.

---

### Question 145: How do you split Contexts?

**Answer:**
Separate logic into domains.
Instead of `<GlobalContext>`, have:
*   `<AuthProvider>` (User data)
*   `<ThemeProvider>` (Dark mode)
*   `<NotificationProvider>` (Toasts)

---

### Question 146: Difference between Context and Redux Toolkit?

**Answer:**
Context is just a dependency injection mechanism. It doesn't "manage" state; it just transports it.
Redux Toolkit (RTK) is a full **State Management Library**. It provides:
*   Predictable state updates (Reducers).
*   Middleware (Thunk).
*   DevTools.
*   Performance optimizations (Selectors).

---

### Question 147: How do you update Context state?

**Answer:**
Since Context is read-only for consumers, you usually pass both the **state** and the **updater functions** in the Provider value.

**Example:**
```jsx
// Provider
<ThemeContext.Provider value={{ theme, toggleTheme }}>
```

---

### Question 148: Can Context replace Redux completely?

**Answer:**
For small to medium applications where you strictly need to avoid props drilling for global settings (Auth/Theme), **Yes**.
For large apps with complex data flows, frequent updates, and need for robust debugging/middleware, **Redux** is still preferred.

---

### Question 149: How do you handle async data in Context?

**Answer:**
The logic usually lives in the **Provider component**.
1.  Provider component fetches data in `useEffect`.
2.  Stores result/loading/error in local `useState`.
3.  Exposes these states via `value` prop to children.

---

### Question 150: How do you test Context-based components?

**Answer:**
You cannot test a specific component that uses `useContext` in isolation without a Provider. You must wrap the tested component in a Provider with mocked values.

**Example (RTL):**
```jsx
render(
  <AuthContext.Provider value={{ user: 'TestUser' }}>
    <UserProfile />
  </AuthContext.Provider>
);
```
