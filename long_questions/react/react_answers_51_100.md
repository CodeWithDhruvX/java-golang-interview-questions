## 🔹 4. Hooks (Advanced) - Continued

### Question 51: What is `useReducer`?

**Answer:**
`useReducer` is a hook that is usually preferable to `useState` when you have complex state logic that involves multiple sub-values or when the next state depends on the previous one. It behaves similarly to Redux.

**Syntax:**
```jsx
const [state, dispatch] = useReducer(reducer, initialState);
```

---

### Question 52: Difference between `useState` and `useReducer`?

**Answer:**
*   **useState:** Best for independent, primitive state values (e.g., toggles, inputs). Use when implementation is simple.
*   **useReducer:** Best for complex objects or when state transitions follow complex business logic. It decouples the *what* (Action) from the *how* (Reducer).

---

### Question 53: What is `useImperativeHandle`?

**Answer:**
It customizes the instance value that is exposed to parent components when using `ref`. It is rarely used, typically when you need to expose imperative methods (like `focus()` or `scroll()`) to a parent from a functional component wrapped in `forwardRef`.

---

### Question 54: What is `forwardRef`?**

**Answer:**
Ref forwarding is a technique for automatically passing a `ref` through a component to one of its children. This is required because functional components do not have instances and thus cannot accept a `ref` directly unless wrapped in `React.forwardRef()`.

**Example:**
```jsx
const Input = React.forwardRef((props, ref) => (
  <input ref={ref} {...props} />
));
```

---

### Question 55: What is `useId`?

**Answer:**
A hook introduced in React 18 for generating unique IDs that are stable across the server and client. This is critical for accessibility (linking labels to inputs) in SSR applications to avoid hydration mismatches.

**Example:**
```jsx
const id = useId();
return <label htmlFor={id}>Name</label>;
```

---

## 🔹 5. Event Handling & Forms

### Question 56: How does event handling work in React?

**Answer:**
React events are named using camelCase (`onClick` instead of `onclick`). You pass a function as the event handler, rather than a string.

**Example:**
```jsx
<button onClick={handleClick}>Click me</button>
```

---

### Question 57: What is synthetic event?

**Answer:**
A **SyntheticEvent** is a cross-browser wrapper around the browser's native event. It has the same interface as the native event (e.g., `stopPropagation()`, `preventDefault()`), but it works identically across all browsers, fixing inconsistencies.

---

### Question 58: Difference between HTML events and React events?

**Answer:**
1.  **Naming:** HTML is lowercase (`onclick`), React is camelCase (`onClick`).
2.  **Value:** HTML uses strings (`"activate()"`), React uses functions (`{activate}`).
3.  **Behavior:** In HTML `return false` prevents default behavior; in React you must call `e.preventDefault()`.

---

### Question 59: How to bind events in React?

**Answer:**
In Class Components, methods are not bound by default.
1.  **Bind in Constructor:** `this.handleClick = this.handleClick.bind(this);`
2.  **Public Class Fields (Arrow Function):** `handleClick = () => { ... }` (Recommended).
3.  **Inline Arrow Function:** `onClick={() => this.handleClick()}` (Can cause performance issues).

---

### Question 60: What are controlled components?

**Answer:**
A component where React **state** is the "single source of truth" for the input value. The input's value is controlled by React, and updates happen via `onChange` handlers.

**Example:**
```jsx
<input value={name} onChange={(e) => setName(e.target.value)} />
```

---

### Question 61: What are uncontrolled components?

**Answer:**
Components where the form data is handled by the **DOM** itself. You access the values using a **Ref** (`useRef` or `createRef`) instead of writing an event handler for every update.

---

### Question 62: Difference between controlled and uncontrolled components?

**Answer:**

| Feature | Controlled | Uncontrolled |
| :--- | :--- | :--- |
| **Data Source** | React State | DOM |
| **Value Access** | `state` variable | `ref.current.value` |
| **Use Case** | Validation, Instant UI feedback | Simple forms, integrating non-React libs |

---

### Question 63: How to handle forms in React?

**Answer:**
Typical flow:
1.  Create state variables for inputs.
2.  Bind `value` prop to state.
3.  Update state in `onChange`.
4.  Handle submission in `onSubmit` (using `e.preventDefault()`).

---

### Question 64: How to handle multiple inputs?

**Answer:**
Use a single state object and dynamic object keys.

**Example:**
```jsx
const [form, setForm] = useState({ name: '', email: '' });

const handleChange = (e) => {
  setForm({ ...form, [e.target.name]: e.target.value });
};
```

---

### Question 65: How to validate forms?

**Answer:**
1.  **Simple:** Check logic in `onSubmit` or `onChange` and set an error state.
2.  **Libraries:** For complex validation, use libraries like **Formik** or **React Hook Form** with **Yup** schemas.

---

## 🔹 6. Conditional Rendering & Lists

### Question 66: What is conditional rendering?

**Answer:**
Rendering different UI elements or components based on a condition (state, props, or logic). It works the same way conditions work in JavaScript.

---

### Question 67: Different ways to do conditional rendering?

**Answer:**
1.  **If/Else:** Outside JSX.
2.  **Ternary Operator:** `condition ? <True /> : <False />`
3.  **Logical AND:** `condition && <True />` (Short-circuit).
4.  **Null:** Return `null` to hide component.

---

### Question 68: How do you render lists in React?

**Answer:**
Use the `.map()` array method to traverse the list and return a list of React elements.

**Example:**
```jsx
<ul>
  {items.map((item) => (
    <li key={item.id}>{item.text}</li>
  ))}
</ul>
```

---

### Question 69: Why is `map()` preferred?

**Answer:**
`map()` creates a **new array** of results, which is exactly what JSX needs inside `{}` to render multiple children. `forEach()` returns `undefined` and thus cannot be embedded directly in JSX.

---

### Question 70: What is fragment?

**Answer:**
A **Fragment** allows you to group a list of children without adding extra nodes to the DOM.
Syntax: `<React.Fragment>...</React.Fragment>` or `<>...</>`.

---

### Question 71: Difference between `<div>` and `React.Fragment`?

**Answer:**
*   `div`: Adds a real DOM node (`<div>`), which can break CSS layouts (like Flexbox/Grid) or produce invalid HTML (e.g., `div` inside `tr`).
*   `Fragment`: Adds **no** node to the DOM, preserving the structure.

---

### Question 72: What is `<> </>` syntax?

**Answer:**
It is the **short syntax** for Fragments.
**Limitation:** You cannot pass any props (like `key`) to short syntax. If you need a key (for a list of fragments), you must use `<React.Fragment key={id}>`.

---

## 🔹 7. Performance Optimization

### Question 73: What causes re-render in React?

**Answer:**
1.  **State Change:** `setState` is called.
2.  **Props Change:** Parent passes new props.
3.  **Parent Re-render:** If a parent re-renders, all children re-render recursively (default behavior).
4.  **Context Change:** Context value updates.

---

### Question 74: How to prevent unnecessary re-renders?

**Answer:**
1.  **React.memo:** Memoize components.
2.  **useMemo / useCallback:** Memoize values/functions to keep props stable.
3.  **Composition:** Push state down or lift content up (pass as children).
4.  **Keys:** Use proper keys in lists.

---

### Question 75: What is `React.memo`?

**Answer:**
It is a **Higher Order Component (HOC)** for functional components. It checks if the props have changed. If props are identical, it skips rendering the component and reuses the last rendered result.

**Example:**
```jsx
const MyComponent = React.memo(function MyComponent(props) {
  /* render using props */
});
```

---

### Question 76: Difference between `React.memo` and `useMemo`?

**Answer:**
*   **React.memo:** Wraps a **Component** to prevent re-rendering.
*   **useMemo:** Wraps a **function/value** *inside* a component to prevent expensive re-calculation.

---

### Question 77: What is lazy loading?

**Answer:**
Lazy loading is a design pattern where you defer the loading of non-critical resources (e.g., images, components) until they are needed (e.g., when the user scrolls to them or navigates to the route), improving initial load time.

---

### Question 78: What is `React.lazy()`?

**Answer:**
It lets you render a dynamic import as a regular component. It automatically loads the bundle containing the component when it's first rendered.

**Example:**
```jsx
const OtherComponent = React.lazy(() => import('./OtherComponent'));
```

---

### Question 79: What is `Suspense`?

**Answer:**
`Suspense` is a component that lets you specify a "loading" state (fallback UI) while the component tree below it is waiting for something (like a lazy-loaded component or data).

**Example:**
```jsx
<Suspense fallback={<div>Loading...</div>}>
  <OtherComponent />
</Suspense>
```

---

### Question 80: How do you optimize large React applications?

**Answer:**
1.  **Code Splitting:** `React.lazy` and Dynamic Imports.
2.  **Virtualization:** Use `react-window` for long lists.
3.  **Memoization:** `React.memo`, `useMemo` for expensive sub-trees.
4.  **Throttling/Debouncing:** For event handlers (scroll/resize/input).
5.  **Production Build:** Ensure minification (`npm run build`).

---

## 🔹 8. Routing (React Router)

### Question 81: What is React Router?

**Answer:**
It is the standard routing library for React. It enables navigation among views (components) in a React application, allows changing the browser URL, and keeps the UI in sync with the URL—all without full page reloads (SPA behavior).

---

### Question 82: Difference between `BrowserRouter` and `HashRouter`?

**Answer:**
*   **BrowserRouter:** Uses HTML5 History API (`pushState`). URLs look like `example.com/login`. Requires server config to handle redirects for SPA.
*   **HashRouter:** Uses URL hash (`window.location.hash`). URLs look like `example.com/#/login`. Supported by legacy browsers, no server config needed.

---

### Question 83: What is `Route`?

**Answer:**
The `Route` component is responsible for rendering some UI when its `path` matches the current URL.

---

### Question 84: What is `Switch` / `Routes`?

**Answer:**
*   **v5 (Switch):** Renders the *first* child `<Route>` that matches the location.
*   **v6 (Routes):** Replaces Switch. It intelligently picks the *best* match (ranking) instead of just the first match.

---

### Question 85: What is `Link`?

**Answer:**
A component allowing declarative navigation. Unlike the HTML `<a>` tag, it changes the URL without reloading the page, preserving the React state.

---

### Question 86: Difference between `Link` and `NavLink`?

**Answer:**
`NavLink` is a special version of `Link` that will add styling attributes (like a class `active`) to the rendered element when it matches the current URL. Perfect for navigation menus.

---

### Question 87: What is `useParams`?

**Answer:**
A hook in React Router that returns an object of key/value pairs of URL parameters.
**URL:** `/users/:id` -> **Path:** `/users/123` -> `useParams()` returns `{ id: '123' }`.

---

### Question 88: What is `useNavigate`?

**Answer:**
A hook introduced in React Router v6 used for programmatic navigation (redirecting user after an action).
**Example:** `const navigate = useNavigate(); navigate('/home');`
(Replaced `useHistory` from v5).

---

## 🔹 9. State Management (Redux & Others)

### Question 89: What is Redux?

**Answer:**
Redux is a predictable state container for JavaScript apps. It helps you write applications that behave consistently, run in different environments (client/server/native), and are easy to test. It centralizes the application's state and logic.

---

### Question 90: Why do we need Redux?

**Answer:**
1.  **Global State:** Share data between components at any nesting level.
2.  **Predictability:** State is read-only; changes are made via pure functions (reducers).
3.  **Debugging:** Redux DevTools allow time-travel debugging.

---

### Question 91: What are actions?

**Answer:**
Actions are payloads of information that send data from your application to your store. They are plain JavaScript objects and **must** have a `type` property.

**Example:**
```js
{ type: 'ADD_TODO', text: 'Learn React' }
```

---

### Question 92: What are reducers?

**Answer:**
Reducers are **pure functions** that take the current `state` and an `action` as arguments and return a new `state`. They specify *how* the application's state changes in response to actions.
`(previousState, action) => newState`

---

### Question 93: What is store?

**Answer:**
The object that holds the application state. There is only **one** store in a Redux app.
Responsibilities:
*   Holds application state.
*   Allows access via `getState()`.
*   Allows state update via `dispatch(action)`.
*   Registers listeners via `subscribe(listener)`.

---

### Question 94: What is Redux middleware?

**Answer:**
Middleware provides a third-party extension point between dispatching an action and the moment it reaches the reducer. It is used for logging, crash reporting, and primarily **asynchronous tasks** (API calls) using libraries like **Redux Thunk** or **Redux Saga**.

---

### Question 95: Difference between Redux and Context API?

**Answer:**

| Feature | Context API | Redux |
| :--- | :--- | :--- |
| **Setup** | Built-in (Easy) | Requires external lib (Boilerplate) |
| **Usage** | Static/Low-freq updates (Theme, User) | High-freq updates, Complex logic |
| **DevTools** | Basic | Advanced (Time Travel) |
| **Middleware** | No built-in support | Built-in support (Thunk/Saga) |

---

## 🔹 10. Advanced & Real-World Questions

### Question 96: How does React handle reconciliation?

**Answer:**
React uses a heuristic O(n) algorithm.
1.  **Diffing:** Compares the two root elements.
2.  **Different Types:** If root elements have different types (div -> span), it tears down the old tree and builds a new one.
3.  **Same Types:** If same type, it keeps the DOM node and just updates changed attributes.
4.  **Keys:** Used to match children in lists to avoid recreating elements.

---

### Question 97: What is hydration in React?

**Answer:**
Hydration is the process where React attaches event listeners to the HTML markup that was rendered on the server (SSR). React "hydrates" the static HTML to make it interactive/dynamic on the client side.

---

### Question 98: What is server-side rendering (SSR)?

**Answer:**
SSR creates the HTML page on the server for each request and sends the fully rendered page to the browser.
**Benefits:** Better **SEO** (crawlers see content), Faster **First Contentful Paint** (user sees content immediately).
**Framework:** Next.js is the standard for React SSR.

---

### Question 99: Difference between CSR and SSR?

**Answer:**
*   **CSR (Client-Side Rendering):** Browser downloads minimal HTML -> downloads JS -> JS runs and renders content. (Slow initial load, fast navigation).
*   **SSR (Server-Side Rendering):** Server sends full HTML -> Browser downloads JS -> JS hydrates HTML. (Fast initial load, better SEO).

---

### Question 100: How do you structure a large-scale React application?

**Answer:**
Standard approach (Feature-based):
```
/src
  /assets
  /components (Shared UI: Button, Input)
  /features (Domain logic)
    /auth (API, Hooks, Slice, Components)
    /dashboard
  /hooks (Shared hooks)
  /utils
  /store
```
Key principles: Separation of Concerns, Atomic Design, and keeping Logic separate from UI (Custom Hooks).
