# React Interview Questions & Answers

## 🔹 1. React Basics (Questions 1-15)

**Q1: What is React?**
A JavaScript library for building user interfaces, maintained by Meta (Facebook). It uses a component-based architecture.

**Q2: Why use React instead of other frameworks?**
Virtual DOM for performance, reusable components, unidirectional data flow, and a strong ecosystem.

**Q3: What are the main features of React?**
JSX, Virtual DOM, Components, One-way Data Binding, and Hooks.

**Q4: What is JSX?**
JavaScript XML. A syntax extension that allows writing HTML elements alongside JavaScript code in React components.

**Q5: Can we use React without JSX?**
Yes, using `React.createElement(component, props, ...children)`, but JSX makes code more readable.

**Q6: What is the virtual DOM?**
A lightweight copy of the real DOM kept in memory. React changes this first before updating the real DOM.

**Q7: How does the virtual DOM work?**
React compares the Virtual DOM to a previous version (Diffing) and updates only changed elements in the Real DOM (Reconciliation).

**Q8: Difference between Real DOM and Virtual DOM?**
Real DOM: Slow updates, HTML abstraction. Virtual DOM: Fast updates, JSON-like representation of UI.

**Q9: What is a React component?**
A reusable, self-contained building block of UI (like a function or a class) that returns HTML/JSX.

**Q10: Types of components in React?**
Functional Components (Stateless/Hook-based) and Class Components (Stateful/Lifecycle-based).

**Q11: Difference between functional and class components?**
Functional: Simple functions, use Hooks. Class: ES6 classes, use `this.state` and Lifecycle methods.

**Q12: What are props?**
Short for Properties. Read-only inputs passed from a parent component to a child to configure it.

**Q13: What is state?**
An internal data storage structure within a component that is mutable and triggers re-renders when changed.

**Q14: Difference between state and props?**
State is internal and mutable. Props are external and immutable.

**Q15: Why is React fast?**
Because of the Virtual DOM and its efficient Diffing algorithm (Reconciliation) that minimizes direct DOM manipulation.

---

## 🔹 2. Components & Props (Questions 16-25)

**Q16: How do you pass data from parent to child?**
By passing attributes (props) to the child component tag: `<Child name="John" />`.

**Q17: Can a child component modify props?**
No, props are read-only (immutable). To change data, the parent must pass a callback function.

**Q18: How do you pass data from child to parent?**
Pass a function from parent to child as a prop. The child calls this function with data as arguments.

**Q19: What is props drilling?**
The process of passing data through multiple layers of intermediate components to reach a deeply nested child.

**Q20: How to avoid props drilling?**
Use Context API, Redux, or Component Composition (Index/Slot pattern).

**Q21: What are default props?**
Default values for props if they are not passed. defined via `defaultProps` or default function parameters.

**Q22: What is `PropTypes`?**
A library (`prop-types`) used to type-check the props passed to a component during development.

**Q23: What is the `key` prop and why is it important?**
A unique identifier for elements in a list. Helps React identify which items changed, ensuring efficient updates.

**Q24: Can we use index as a key?**
Yes, but only if the list is static (no reordering/filtering), otherwise it causes rendering bugs.

**Q25: What happens if `key` is not provided?**
React will warn in the console and default to using the index, which may lead to performance issues or state bugs.

---

## 🔹 3. State & Lifecycle (Class + Hooks) (Questions 26-40)

**Q26: What is component lifecycle?**
The stages a component goes through: Mounting (created), Updating (changed), and Unmounting (removed).

**Q27: Explain React lifecycle methods.**
Class methods hooked into lifecycle phases: `componentDidMount`, `componentDidUpdate`, `componentWillUnmount`.

**Q28: What is `componentDidMount()`?**
Runs once after the initial render. Ideal for API calls, subscriptions, or DOM manipulation.

**Q29: What is `componentDidUpdate()`?**
Runs after a scheduled update (re-render) occurs (props/state change). Good for network requests based on changes.

**Q30: What is `componentWillUnmount()`?**
Runs just before the component is removed from DOM. Used for cleanup (cancelling timers/requests).

**Q31: Why lifecycle methods are not used in functional components?**
Functional components are just functions. Hooks (`useEffect`) were introduced to handle side effects and lifecycle behavior.

**Q32: What are React Hooks?**
Functions that let you "hook into" React state and lifecycle features from functional components.

**Q33: Why were hooks introduced?**
To reuse stateful logic, simplify complex components, and avoid confusing `this` keyword in classes.

**Q34: What is `useState`?**
A Hook that lets you add React state to function components. `const [count, setCount] = useState(0);`.

**Q35: What is `useEffect`?**
A Hook for performing side effects (data fetching, subscriptions) in function components.

**Q36: Difference between `useEffect` and lifecycle methods?**
`useEffect` combines `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount` into one API.

**Q37: How many times `useEffect` runs?**
By default, after every render. With `[]` dep array, implies once (Mount). With `[val]`, implies on value change.

**Q38: What is cleanup function in `useEffect`?**
A function returned by the effect. It runs before the component unmounts or before the effect re-runs.

**Q39: What is `useLayoutEffect`?**
Similar to `useEffect`, but fires synchronously after all DOM mutations. Useful for measuring DOM layout.

**Q40: Can hooks be used inside loops or conditions?**
No. Hooks must be called at the top level to ensure they are called in the same order each render.

---

## 🔹 4. Hooks (Advanced) (Questions 41-55)

**Q41: What is `useContext`?**
A Hook to consume values from a React Context without wrapping the component in a Consumer.

**Q42: What problem does `useContext` solve?**
Props drilling. It allows sharing global data (theme, auth) across component tree easily.

**Q43: What is `useRef`?**
A Hook that returns a mutable ref object. Useful for accessing DOM elements or persisting values without re-renders.

**Q44: Difference between `useRef` and `useState`?**
Updating `useState` triggers a re-render. Updating `useRef.current` does not trigger a re-render.

**Q45: What is `useMemo`?**
A Hook that memoizes a computed value. Re-computes only when dependencies change. Optimizes expensive calculations.

**Q46: What is `useCallback`?**
A Hook that memoizes a function definition. Prevents function recreation on every render.

**Q47: Difference between `useMemo` and `useCallback`?**
`useMemo` returns a memoized *value*. `useCallback` returns a memoized *function*.

**Q48: What is custom hook?**
A JavaScript function whose name starts with "use" and calls other hooks. Used to extract reusable logic.

**Q49: When should you create a custom hook?**
When you want to share logic (e.g., fetching data, form handling) between two JavaScript functions (components).

**Q50: Rules of Hooks?**
1. Only call Hooks at the top level. 2. Only call Hooks from React function components or custom hooks.

**Q51: What is `useReducer`?**
A Hook for managing complex state logic using a reducer function `(state, action) => newState`. Similar to Redux.

**Q52: Difference between `useState` and `useReducer`?**
`useState` is for simple state. `useReducer` is for complex state transitions or when next state depends on previous.

**Q53: What is `useImperativeHandle`?**
Customizes the instance value that is exposed to parent components when using `ref`. Rare use case.

**Q54: What is `forwardRef`?**
A higher-order component that allows a parent to pass a `ref` down to a specific child DOM node.

**Q55: What is `useId`?**
A Hook for generating unique IDs that are stable across the server and client (for accessibility/forms).

---

## 🔹 5. Event Handling & Forms (Questions 56-65)

**Q56: How does event handling work in React?**
Using camelCase props (`onClick`, `onSubmit`) and passing a function handler.

**Q57: What is synthetic event?**
A cross-browser wrapper around the browser's native event. Ensures consistent behavior across browsers.

**Q58: Difference between HTML events and React events?**
HTML: `onclick="handleClick()"`, strings. React: `onClick={handleClick}`, functions, SyntheticEvent.

**Q59: How to bind events in React?**
In Class: `this.handleClick.bind(this)` or arrow functions `handleClick = () => {}`. In Functions: Just pass the function.

**Q60: What are controlled components?**
Form inputs whose value is controlled by React state. `value={state}` and `onChange={handler}`.

**Q61: What are uncontrolled components?**
Form inputs whose value is handled by the DOM itself. Accessed via Refs. default values used.

**Q62: Difference between controlled and uncontrolled components?**
Controlled: React converts input to state (Single Source of Truth). Uncontrolled: React reads from DOM when needed (Legacy/Simple).

**Q63: How to handle forms in React?**
Use `useState` for input values, `onChange` to update state, and `onSubmit` to handle submission.

**Q64: How to handle multiple inputs?**
Use a single `useState` object/dictionary and update using `e.target.name`: `setState({...state, [e.target.name]: value})`.

**Q65: How to validate forms?**
Check state values in `onChange` or `onSubmit`. Display error messages conditionally. use Formik/React Hook Form for complex cases.

---

## 🔹 6. Conditional Rendering & Lists (Questions 66-72)

**Q66: What is conditional rendering?**
Rendering different UI elements based on a condition (state/props).

**Q67: Different ways to do conditional rendering?**
`if-else`, Ternary operator `? :`, Logical AND `&&`, or Switch statements.

**Q68: How do you render lists in React?**
Using JavaScript's `map()` function to transform an array of data into an array of JSX elements.

**Q69: Why is `map()` preferred?**
It returns a new array of React elements, which React can render directly. `forEach` returns nothing.

**Q70: What is fragment?**
A wrapper to group a list of children without adding extra nodes (like `<div>`) to the DOM.

**Q71: Difference between `<div>` and `React.Fragment`?**
`<div>` adds an extra node to the DOM tree (can break layout). Fragment does not add any node.

**Q72: What is `<> </>` syntax?**
Shorthand for `<React.Fragment>`. Note: Can't accept keys (unlike `<React.Fragment key={...}>`).

---

## 🔹 7. Performance Optimization (Questions 73-80)

**Q73: What causes re-render in React?**
Change in State, Change in Props, or Parent re-rendering.

**Q74: How to prevent unnecessary re-renders?**
`React.memo`, `useMemo`, `useCallback`, ensuring stable props, and `key` in lists.

**Q75: What is `React.memo`?**
A Higher Order Component that memoizes a functional component. Renders only if props change.

**Q76: Difference between `React.memo` and `useMemo`?**
`React.memo` wraps a Component (prevents re-render). `useMemo` wraps a calculation/value (prevents re-computation).

**Q77: What is lazy loading?**
Loading components only when they are needed (e.g., when route is visited), reducing initial bundle size.

**Q78: What is `React.lazy()`?**
A function that lets you render a dynamic import as a regular component. `const Other = React.lazy(() => import('./Other'))`.

**Q79: What is `Suspense`?**
A component that lets you specify a loading state (spinner) while the children (lazy components) are loading.

**Q80: How do you optimize large React applications?**
Code splitting, Lazy loading, Memoization, Virtualization (large lists), Image optimization, tree-shaking.

---

## 🔹 8. Routing (React Router) (Questions 81-88)

**Q81: What is React Router?**
Standard library for routing in React. Emulates Multi-page app behavior in a Single Page App (SPA).

**Q82: Difference between `BrowserRouter` and `HashRouter`?**
Browser: Uses HTML5 History API (clean URLs). Hash: Uses URL hash `/#/` (legacy support).

**Q83: What is `Route`?**
Renders UI when the current URL matches the path prop.

**Q84: What is `Switch` / `Routes`?**
`Routes` (v6) / `Switch` (v5) groups Routes and ensures only the first matching Route renders.

**Q85: What is `Link`?**
A component to create links. Unlike `<a>`, it updates the URL without reloading the page.

**Q86: Difference between `Link` and `NavLink`?**
`NavLink` is a special `Link` that adds a "active" class style when it matches the current URL.

**Q87: What is `useParams`?**
A Hook to access dynamic parameters from the URL (e.g., `/user/:id` -> `params.id`).

**Q88: What is `useNavigate`?**
A Hook (v6) to programmatically navigate/redirect users (e.g., `navigate('/home')`). Replaces `useHistory`.

---

## 🔹 9. State Management (Redux & Others) (Questions 89-95)

**Q89: What is Redux?**
A predictable state container for JS apps. Centralizes application state in a single Store.

**Q90: Why do we need Redux?**
For managing complex global state, predictable changes, easier debugging (DevTools), and avoiding props drilling.

**Q91: What are actions?**
Plain JS objects sending data to store. Must have `type` property. `const add = { type: 'ADD' }`.

**Q92: What are reducers?**
Pure functions `(state, action) => newState` that specify how state changes in response to actions.

**Q93: What is store?**
The object that brings Actions and Reducers together. Holds the state.

**Q94: What is Redux middleware?**
Extension point (e.g., Redux Thunk, Saga) to handle async logic or logging between Action and Reducer.

**Q95: Difference between Redux and Context API?**
Context: Built-in, good for low-frequency updates (Theme/User). Redux: External, good for high-freq updates/complex logic/DevTools.

---

## 🔹 10. Advanced & Real-World Questions (Questions 96-100)

**Q96: How does React handle reconciliation?**
Uses Diffing algorithm. Compares element types. If type changes, destroys tree. If same type, updates props. Keys help lists.

**Q97: What is hydration in React?**
Process of attaching event listeners to static HTML rendered by Server-Side Rendering (SSR) to make it interactive.

**Q98: What is server-side rendering (SSR)?**
Rendering React components to HTML on the server and sending HTML to client. Improves SEO and First Contentful Paint.

**Q99: Difference between CSR and SSR?**
CSR: Browser downloads JS -> Renders HTML. Slow init load, Fast interaction. SSR: Server sends HTML -> Browser hydrates. Fast init load.

**Q100: How do you structure a large-scale React application?**
Feature-based folders (`/features/auth`, `/features/feed`), atomic design for UI components, separation of concerns (Hooks vs UI).

---

## 🔹 11. Core React Internals (Questions 101-110)

**Q101: What is Fiber in React?**
The new reconciliation engine in React 16. It enables incremental rendering and prioritization of updates.

**Q102: What problem does React Fiber solve?**
It fixes the issue where large updates blocked the main thread (jank). Fiber allows splitting work into chunks (Time Slicing).

**Q103: What is concurrent rendering?**
React's ability to prepare multiple versions of the UI at the same time. It doesn't block the main thread.

**Q104: What is time slicing in React?**
Splitting high-priority work (user input) and low-priority work (data fetching) into small tasks to keep the UI responsive.

**Q105: What is automatic batching in React 18?**
React batches all state updates (even in promises/timeouts) into a single re-render. Previously only event handlers were batched.

**Q106: What is Strict Mode in React?**
A development-only tool (`<React.StrictMode>`) that highlights potential problems (deprecated lifecycles, unsafe effects).

**Q107: Why does React Strict Mode render components twice?**
To help detect side effects in render phase. If a component is impure, double rendering makes it obvious (bugs appear).

**Q108: What is reconciliation algorithm?**
The process React uses to verify and update the DOM. React 16+ uses Fiber, older versions used Stack Reconciliation.

**Q109: How does React diffing work?**
It assumes elements of different types produce different trees. It uses Keys to differentiate sibling elements efficiently.

**Q110: What is the render phase vs commit phase?**
Render: Pure, calculates changes (Fiber tree). Can be paused/aborted. Commit: Impure, applies changes to DOM. Cannot be interrupted.

---

## 🔹 12. Rendering Behavior (Questions 111-120)

**Q111: What triggers a re-render in React?**
State change, Props change, Parent re-render, or Context value change.

**Q112: How does React decide when to re-render a component?**
By default, whenever the parent renders. `React.memo` or `shouldComponentUpdate` can change this behavior.

**Q113: What happens when state is updated with the same value?**
React bails out (skips render). It uses `Object.is` comparison.

**Q114: Why are state updates asynchronous?**
For performance (Batching). Updating state immediately would cause too many re-renders.

**Q115: How does batching of state updates work?**
React groups multiple `setState` calls into a single update pass to avoid unnecessary renders.

**Q116: What is stale state?**
When a closure captures an old state value (e.g., inside `useEffect` or `setTimeout`) and doesn't see updates.

**Q117: How do you handle stale closures?**
Use functional state updates `setCount(prev => prev + 1)` or add the state variable to the dependency array.

**Q118: Why should state updates be immutable?**
React relies on reference equality checking (`prevProps !== nextProps`). Mutable changes don't trigger this check.

**Q119: What is referential equality?**
Checking if two variables point to the exact same object in memory (`===`).

**Q120: How does React handle immutability?**
It doesn't enforce it, but requires it for Diffing. Devs use spread `...` or libraries like Immer.

---

## 🔹 13. JSX & JavaScript Concepts (Questions 121-130)

**Q121: Why must JSX have a single parent?**
Because a function (the component) can only return one value (one root Object).

**Q122: How do expressions work inside JSX?**
Wrapped in `{}`. Any valid JS expression (variables, functions, math) can be evaluated.

**Q123: Difference between `{}` and `{{}}` in JSX?**
`{}` is for JS expression. `{{}}` is just an object literal inside an expression (often used for inline styles).

**Q124: How do you add comments in JSX?**
`{/* This is a comment */}`.

**Q125: How do you render HTML safely in React?**
By default React escapes HTML. To render raw HTML, use `dangerouslySetInnerHTML`.

**Q126: What is `dangerouslySetInnerHTML`?**
React's replacement for `innerHTML`. Reminds developers of XSS risks. `{__html: '<p>...</p>'}`.

**Q127: How does React escape values?**
It converts values to strings before rendering, preventing XSS injection scripts.

**Q128: How do inline styles work in React?**
Pass an object with camelCased properties: `style={{ backgroundColor: 'red', fontSize: '12px' }}`.

**Q129: Why are class names written as `className`?**
Because `class` is a reserved keyword in JavaScript.

**Q130: What is spread operator usage in JSX?**
`<Component {...props} />`. Passes all properties of the object as individual props.

---

## 🔹 14. Forms & Input Handling (Advanced) (Questions 131-140)

**Q131: How do you debounce input in React?**
Use `setTimeout` inside `useEffect` or a library like `lodash.debounce`. Only trigger search after user stops typing.

**Q132: How do you throttle API calls from inputs?**
Limit execution frequency (e.g., once every 500ms). Useful for scroll events or resizing.

**Q133: How do you manage complex forms?**
Use libraries like **React Hook Form** or **Formik**. They handle validation, errors, and submission state efficiently.

**Q134: How do you handle dynamic form fields?**
Map over an array of field objects stored in state. Add/Remove items from that array to update UI.

**Q135: How do you reset a form in React?**
Set the state to initial values. In HTML forms, `form.reset()` can be used (for uncontrolled).

**Q136: How do you manage focus in forms?**
Use `useRef` -> `inputRef.current.focus()`.

**Q137: How do you auto-focus an input?**
`autoFocus` prop or `useEffect(() => ref.current.focus(), [])`.

**Q138: How do you handle keyboard events?**
`onKeyDown`, `onKeyUp`, or `onKeyPress`. Check `e.key` (e.g., 'Enter').

**Q139: How do you handle copy/paste events?**
`onCopy`, `onPaste` props. `e.clipboardData.getData('Text')`.

**Q140: How do you handle validation errors gracefully?**
Store errors in state object `{ field: 'error msg' }`. Render error text conditionally below input.

---

## 🔹 15. Context API (Advanced) (Questions 141-150)

**Q141: How does Context API work internally?**
It allows a Provider to pass data down the tree. Consumers scan up the tree to find the nearest Provider.

**Q142: What causes Context consumers to re-render?**
Whenever the value prop of the Provider changes (Reference change).

**Q143: How do you optimize Context performance?**
Split Contexts (StateContext vs DispatchContext). Memoize the value passed to Provider.

**Q144: When should you NOT use Context?**
For high-frequency updates (e.g., animations, keyboard strokes). use separate state or Redux/Zustand.

**Q145: How do you split Contexts?**
Create separate contexts for different domains (UserContext, ThemeContext) instead of one big state.

**Q146: Difference between Context and Redux Toolkit?**
Context is just a transport mechanism (Dependency Injection). RTK is a state manager with tools for fetching, caching, and debugging.

**Q147: How do you update Context state?**
Pass the `setState` function (or dispatch) along with the state value in the Context Provider.

**Q148: Can Context replace Redux completely?**
For simple apps, yes. For complex apps with middleware/dev-tools needs, No.

**Q149: How do you handle async data in Context?**
Fetch data in the Provider component, then expose the data and loading state via the Context value.

**Q150: How do you test Context-based components?**
Wrap the component in the specific `<MyContext.Provider value={mockValue}>` in the test.

---

## 🔹 16. Component Patterns (Questions 151-160)

**Q151: What is compound component pattern?**
Components that work together to share state implicitly (e.g., `<Select><Option /></Select>`). Uses Context/Children map.

**Q152: What is render props pattern?**
Passing a function as a prop (usually `children` or `render`) that returns JSX. Sharing logic (like coordinates).

**Q153: Difference between HOC and render props?**
HOC wraps component (Static composition). Render Props injects logic inside render (Dynamic).

**Q154: What is container-presentational pattern?**
Container handles logic/state/fetching. Presentational handles UI/Styles. (Less common since Hooks).

**Q155: What is controlled vs uncontrolled pattern?**
Allowing a component to be either controlled (via value value) or uncontrolled (internal state).

**Q156: What is inversion of control in React?**
Letting the user of the component control the rendering (via Render Props) rather than the component deciding everything.

**Q157: What is slot pattern?**
Passing JSX into named props instead of just `children`. `<Layout header={<Header />} />`.

**Q158: What is provider pattern?**
Using a Context Provider to inject dependencies into the tree.

**Q159: What is polymorphic component?**
A component that can render as different HTML tags (`as="a"`, `as="button"`).

**Q160: What is headless component?**
A component (usually a hook) providing only logic and accessibility, no UI. (e.g., React Table, Headless UI).

---

## 🔹 17. Code Splitting & Architecture (Questions 161-170)

**Q161: What is dynamic import in React?**
`import('./Module')`. Returns a Promise. Used with `React.lazy`.

**Q162: How does webpack help React apps?**
Bundles JS/CSS/Assets. Transpiles JSX (via Babel loader). Enables HMR (Hot Module Replacement).

**Q163: What is tree shaking?**
Removing unused code from the bundle during build process (Static Analysis).

**Q164: How does bundling affect performance?**
Reduces HTTP requests (one file). But if too large, slows initial load. Code splitting balances this.

**Q165: How do you organize components folder?**
By Feature (`/auth`, `/dashboard`) or by Type (`/atoms`, `/molecules`).

**Q166: How do you manage reusable components?**
Keep them in a shared `/components/ui` folder. Ensure they are pure and accept props for customization.

**Q167: How do you manage environment variables?**
`.env` files. Access via `process.env.REACT_APP_KEY` (CRA) or `import.meta.env` (Vite).

**Q168: How do you handle feature flags?**
Use a Context or Service to check flag status. `if (!flags.newFeature) return null;`.

**Q169: What is micro-frontend architecture?**
Splitting a large frontend into independent apps (e.g., Header App, Cart App) loaded together.

**Q170: How does React support micro-frontends?**
Via Module Federation (Webpack) or embedding components from different builds.

---

## 🔹 18. Testing in React (Questions 171-180)

**Q171: What is unit testing in React?**
Testing individual components or functions in isolation (e.g., "Button clicks call handler").

**Q172: What is integration testing?**
Testing how multiple components work together (e.g., "Form submits and shows success message").

**Q173: What is end-to-end testing?**
Simulating real user flows in a browser (e.g., Cypress/Playwright). "Login -> Navigate -> Logout".

**Q174: What is Jest?**
A test runner and assertion library. Standard for React apps.

**Q175: What is React Testing Library (RTL)?**
Library for testing React components. Encourages testing behavior (what user sees) not implementation (state).

**Q176: Difference between Enzyme and RTL?**
Enzyme: Tests implementation details (Checking state/props). RTL: Tests DOM/Accessibility (Clicking buttons).

**Q177: How do you test hooks?**
Use `renderHook` from `@testing-library/react-hooks`.

**Q178: How do you mock API calls?**
Use `jest.mock()` or libraries like MSW (Mock Service Worker) to intercept requests.

**Q179: How do you test async components?**
Use `await waitFor(() => expect(...))` or `findBy` queries in RTL.

**Q180: What is snapshot testing?**
Comparing the rendered JSON output of a component to a saved "snapshot" file to detect unintended changes.

---

## 🔹 19. Error Handling & Security (Questions 181-190)

**Q181: What are Error Boundaries?**
Class components that catch JS errors in their child component tree (`componentDidCatch`).

**Q182: Why error boundaries don’t catch async errors?**
They only catch errors in Render, Lifecycle, and Constructors. Async errors (Callbacks, Promises) happen outside that flow.

**Q183: How do you implement error boundaries?**
Define `static getDerivedStateFromError(error)` to render fallback UI.

**Q184: How do you handle API errors?**
`try/catch` in async functions. Store error in state. Display feedback.

**Q185: How do you handle 404 pages in React?**
Add a catch-all Route `path="*"` at the end of Routes.

**Q186: How do you prevent XSS in React?**
React auto-escapes data. Avoid `dangerouslySetInnerHTML`. Validate URLs (`javascript:` protocol).

**Q187: What is CSRF and how do you prevent it?**
Cross-Site Request Forgery. Use CSRF tokens on server. SameSite cookies.

**Q188: How do you secure React applications?**
No secrets in client code. Sanitize inputs. strict CSP headers.

**Q189: What is CORS?**
Browser security feature blocking requests to different origins. Configure Server to allow Origin or use Proxy.

**Q190: How do you handle authentication in React?**
Store JWT in HttpOnly Cookie (best) or LocalStorage. Use AuthContext to persist user session.

---

## 🔹 20. Deployment & Best Practices (Questions 191-200)

**Q191: How do you build React for production?**
`npm run build`. Creates minified, optimized static files in `/build` or `/dist`.

**Q192: What is environment-based configuration?**
Using different `.env` files (`.env.development`, `.env.production`) for API URLs, etc.

**Q193: How do you improve first contentful paint?**
SSR, critical CSS, small bundles, CDN, text compression.

**Q194: What is lighthouse score?**
Metric for Performance, Accessibility, SEO, Best Practices.

**Q195: How do you handle browser compatibility?**
Polyfills (e.g., `core-js`), Transpilation (Babel), and checking caniuse.com.

**Q196: How do you handle memory leaks?**
Cleanup listeners/intervals in `useEffect` return function. Cancel async requests if unmounted.

**Q197: How do you cancel API calls in React?**
`AbortController`. Pass `signal` to fetch. `controller.abort()` in cleanup.

**Q198: What is cleanup in async effects?**
The function returned by `useEffect` that runs when dependency changes or unmounts.

**Q199: How do you log errors in production?**
Sentry, LogRocket, or datadog RUM.

**Q200: What are React best practices?**
Functional Components, Hooks, Small Components, Prop Types/TS, Immutable State, Memoization where needed.

---

## 🔹 21. React 18 & Modern React (Questions 201-210)

**Q201: What is `createRoot` in React 18?**
The new entry point API that enables concurrent features in React 18. Replaces `ReactDOM.render`.

**Q202: Difference between `render` and `createRoot`?**
`render` (legacy) renders synchronously. `createRoot` enables concurrent rendering, automatic batching, and transitions.

**Q203: What is concurrent features opt-in?**
In React 18, concurrent rendering is only enabled when you use concurrent features like `startTransition` or `useDeferredValue`.

**Q204: What is `startTransition`?**
API to wrap state updates that are non-urgent (e.g., filtering a list). Keeps the UI responsive to user input.

**Q205: When should you use `useTransition`?**
When you have a computation-heavy UI update that blocks user interaction (like typing). It marks the update as low priority.

**Q206: What problem does `useDeferredValue` solve?**
Similar to debouncing. It defers updating a value until more urgent updates (typing) have finished.

**Q207: How does React handle priority updates?**
Using lanes. High priority: User input. Low priority: Data fetching/Transitions. High priority interrupts low priority.

**Q208: What are blocking vs non-blocking updates?**
Blocking: React <18, once render starts, it blocks main thread. Non-blocking: React 18+, can pause render to handle input.

**Q209: What changes were introduced in React 18?**
Concurrent Mode, Automatic Batching, Suspense on Server, `useId`, `startTransition`, `useInsertionEffect`.

**Q210: How does React 18 improve performance?**
By not blocking the main thread (Time Slicing) and batching more updates automatically.

---

## 🔹 22. State & Data Flow (Advanced) (Questions 211-220)

**Q211: What is derived state?**
State that can be calculated from existing props or state. Ideally should be variables, not separate state to avoid sync bugs.

**Q212: Why is derived state considered an anti-pattern?**
If you copy props to state, you must manually sync them. It uses extra memory and leads to "source of truth" conflicts.

**Q213: How do you lift state up?**
Move state to the nearest common ancestor of components that need it, pass it down via props.

**Q214: How do you share state between sibling components?**
Lift state to parent or use Context / Global Store (Redux/Zustand).

**Q215: What is single source of truth?**
The principle that state should live in one place, and all components strictly read from that one place.

**Q216: How do you normalize state?**
Store data like a database (Flat structure, IDs as keys) instead of deeply nested arrays. Improves update performance.

**Q217: How do you manage deeply nested state?**
Normalization, Immer (for mutable syntax), or useReducer.

**Q218: What are atomic state libraries?**
Libraries like Recoil or Jotai where state is split into small "atoms" that components subscribe to individually.

**Q219: What is state colocation?**
Keeping state as close as possible to where it is used (Local vs Global). Improves maintainability.

**Q220: When should state be local vs global?**
Local: Form inputs, Toggles specific to component. Global: User Auth, Theme, Shopping Cart.

---

## 🔹 23. Side Effects & Async Logic (Questions 221-230)

**Q221: How do you handle async logic in React?**
`useEffect` for fetching, or libraries like React Query / SWR / Redux Thunk.

**Q222: What problems can async effects cause?**
Race conditions (responses arriving out of order), Memory leaks (updating unmounted component).

**Q223: How do you avoid race conditions in `useEffect`?**
Use a boolean flag variable `let ignore = false` inside effect and check it before setting state.

**Q224: How do you handle polling in React?**
`setInterval` in `useEffect`. Or use React Query (`refetchInterval`).

**Q225: How do you handle WebSockets in React?**
Open connection in `useEffect` (empty dep), close in cleanup return. Store socket instance in `useRef` or Context.

**Q226: How do you handle long-running tasks?**
Web Workers (off main thread) or `useTransition` to keep UI responsive.

**Q227: How do you integrate background sync?**
Service Workers (PWA) or simple `window.addEventListener('online', ...)` to retry requests.

**Q228: How do you retry failed API calls?**
Recursive function with delay, or use React Query’s built-in retry mechanism.

**Q229: How do you handle optimistic UI updates?**
Update UI state immediately before API call. If API fails, rollback state to previous value.

**Q230: How do you roll back optimistic updates?**
Keep previous state copy. In `.catch()`, set state back to that copy.

---

## 🔹 24. Refs & DOM Manipulation (Questions 231-240)

**Q231: When should you access the DOM directly?**
Media playback, focus management, measuring element size/position, or integrating non-React libs (D3.js).

**Q232: How do callback refs work?**
Passing a function to `ref` attribute. React calls it with the DOM node on mount and `null` on unmount.

**Q233: Difference between object refs and callback refs?**
Object ref (`useRef`): `.current` property. Callback ref: Function gives finer control over when node attaches.

**Q234: How do you measure element size in React?**
`ref.current.getBoundingClientRect()` or `ResizeObserver`.

**Q235: How do you detect outside clicks?**
Global `click` listener on `document`. Check `!ref.current.contains(e.target)`.

**Q236: How do you implement infinite scrolling?**
Intersection Observer API monitoring a "loader" element at bottom of list.

**Q237: How do you manage scroll restoration?**
Save `window.scrollY` in sessionStorage/Context before navigation. Restore in `useLayoutEffect`.

**Q238: How do you handle animations using refs?**
Libraries like GSAP/Anime.js manipulate DOM nodes directly via refs for high performance.

**Q239: How do you integrate third-party DOM libraries?**
Initialize lib in `useEffect`, attach to a `ref` element. Destroy/Cleanup in return function.

**Q240: Why should direct DOM manipulation be minimized?**
It bypasses React's Virtual DOM, potentially causing sync issues and bugs where React overrides your changes.

---

## 🔹 25. Styling in React (Questions 241-250)

**Q241: Different ways to style React components?**
Inline styles, CSS files, CSS Modules, CSS-in-JS (Styled Components), UI Frameworks (MUI/Tailwind).

**Q242: What are CSS Modules?**
CSS files where class names are scoped locally by default. `import styles from './Button.module.css'`.

**Q243: What is styled-components?**
A library for writing actual CSS in JavaScript files using tagged template literals.

**Q244: Difference between CSS Modules and styled-components?**
CSS Modules: Separate .css files, Scoped classes. Styled Components: CSS inside JS, Dynamic styling based on props easier.

**Q245: What is inline styling limitation?**
No Pseudo-selectors (`:hover`), no Media Queries, performance cost if objects recreated every render.

**Q246: How do you apply dynamic styles?**
Conditional classes (`className={active ? 'blue' : 'red'}`) or passing props to styled-components.

**Q247: How do you handle theming in React?**
CSS Variables (Custom Properties) or `ThemeProvider` (Context) in CSS-in-JS libraries.

**Q248: How do you manage dark mode?**
Toggle a class (`.dark`) on `<body/html>`. Use CSS variables for colors. Persist preference in LocalStorage.

**Q249: How do you avoid CSS conflicts?**
Use CSS Modules, methodology like BEM, or Shadow DOM.

**Q250: What is critical CSS?**
Inlining the minimal CSS required for the above-the-fold content to speed up rendering (Core Web Vitals).

---

## 🔹 26. Accessibility (a11y) (Questions 251-260)

**Q251: Why is accessibility important in React apps?**
Inclusivity (users with disabilities), SEO benefits, and Legal requirements (ADA/WCAG).

**Q252: How do you make React apps accessible?**
Semantic HTML, `alt` text, `aria-*` attributes, Keyboard navigation support, managing focus.

**Q253: What are ARIA roles?**
Attributes (e.g., `role="dialog"`) defining the purpose of an element to assistive technologies (Screen Readers).

**Q254: How do you manage focus for modals?**
Trap focus inside modal (Tab cycles inside). Restore focus to trigger button when closed.

**Q255: How do you make forms accessible?**
Labels (`htmlFor`), proper inputs types, error messages associated via `aria-describedby`.

**Q256: How do you support keyboard navigation?**
Ensure all interactive elements are focusable (`tabindex="0"`) and triggerable via Enter/Space using `onKeyDown`.

**Q257: How do you handle screen readers?**
Use specific text for icons (`aria-label="Delete"`), hide decorative elements (`aria-hidden="true"`).

**Q258: What is semantic HTML and why is it important?**
Using `<button>`, `<nav>`, `<main>` instead of `<div>`. Provides native accessibility keyboard behavior.

**Q259: How do you test accessibility?**
Keyboard-only manual test, Lighthouse audits, Axe DevTools, Screen Readers (NVDA/VoiceOver).

**Q260: What tools are used for accessibility testing?**
`eslint-plugin-jsx-a11y`, `@axe-core/react`.

---

## 🔹 27. Integration & Ecosystem (Questions 261-270)

**Q261: How do you integrate REST APIs in React?**
`fetch()` or `axios` inside `useEffect` or event handlers.

**Q262: How do you integrate GraphQL with React?**
Use `Apollo Client` or `urql`. Queries written in template strings `gql`..

**Q263: What is Apollo Client?**
A comprehensive state management library for JavaScript that enables you to manage both local and remote data with GraphQL.

**Q264: What is React Query / TanStack Query?**
Library for fetching, caching, synchronizing and updating server state in React. "Missing data-fetching library for React".

**Q265: Difference between Redux and React Query?**
Redux: Global Client State (UI state). React Query: Server State (Caching, Background updates, Stale data).

**Q266: How do you cache server data?**
React Query / SWR do this automatically (Cache-first strategy). Or manual Key-Value store in Redux.

**Q267: How do you handle pagination?**
Pass `page` param to API. Store current page in state. On Next, update state -> triggers fetch.

**Q268: How do you handle infinite queries?**
`useInfiniteQuery` (React Query). Appends new data to existing list in the cache.

**Q269: How do you sync server state and UI state?**
Optimistic updates or invalidating queries (Refetch) after mutation to get fresh server state.

**Q270: How do you handle real-time updates?**
WebSockets (Socket.io), Server-Sent Events (SSE), or Polling. Sync with state store.

---

## 🔹 28. Build Tools & Configuration (Questions 271-280)

**Q271: What is Vite and how is it different from CRA?**
Vite: Native ES Modules based dev server (Fast start). CRA: Webpack based (Slow start, bundles everything).

**Q272: What is Create React App (CRA)?**
Opinionated, zero-config tool to bootstrap React apps. Now legacy/deprecated in favor of Vite/Next.js.

**Q273: Why is CRA being deprecated?**
Slow, hard to customize (ejecting), doesn't support modern rendering patterns (SSR/RSC) out of box.

**Q274: What is Babel configuration used for?**
Defining presets (`@babel/preset-react`) to compile JSX/ES6+ down to older JavaScript versions.

**Q275: What is ESLint and why is it used?**
Linter tool. Finds syntax errors, anti-patterns, and enforces coding style/rules.

**Q276: What is Prettier?**
Code formatter. Enforces consistent style (indentation, quotes) automatically.

**Q277: How do you enforce code quality?**
Husky (Pre-commit hooks) running Lint/Test/Prettier. CI/CD pipelines.

**Q278: How do you manage monorepos?**
NPM Workspaces, Yarn Workspaces, Lerna, Nx, or Turborepo.

**Q279: What is Nx / Turborepo?**
Build systems for Monorepos. They support caching (only rebuild changed parts) and parallel execution.

**Q280: How do you manage multiple environments?**
CI/CD variables injected at build time. Runtime configuration loading (for Docker).

---

## 🔹 29. Advanced Patterns & Real-World Scenarios (Questions 281-290)

**Q281: How do you implement role-based access control?**
Global Auth Context with User Roles. Create `<PrivateRoute roles={['admin']}>` wrapper component.

**Q282: How do you protect private routes?**
Check auth state. If not logged in, `<Navigate to="/login" />`. If logged in, render `<Outlet />` or Child.

**Q283: How do you handle feature-based rendering?**
Feature Flags. Conditional check `if (isOn('new-dashboard')) return <NewDash />;`.

**Q284: How do you handle large tables efficiently?**
Virtualization (React Virtual), Pagination, or Infinite Scroll. Render only visible rows.

**Q285: How do you implement virtualized lists?**
Libraries like `react-window` or `react-virtualized`. Calculates total height, renders only items in viewport + buffer.

**Q286: What is windowing?**
Same as virtualization. Term used for mostly rendering "window" of visible content.

**Q287: How do you handle file downloads?**
Create `<a>` tag with `href={blobUrl}` and `download` attr, click strictly, then remove. Or use `file-saver`.

**Q288: How do you handle file previews?**
`URL.createObjectURL(file)` to get a blob URL. Set as `src` of `<img>`.

**Q289: How do you handle multi-language support (i18n)?**
`react-i18next` or `react-intl`. Store translations in JSON. Hook `useTranslation` returns correct string.

**Q290: How do you manage translations dynamically?**
Load translation JSON files asynchronously (Lazy load) based on user locale selection.

---

## 🔹 30. Debugging & Maintenance (Questions 291-300)

**Q291: How do you debug React applications?**
React DevTools (Components/Profiler), Console logs, Breakpoints in Source tab, Network tab.

**Q292: What are common React anti-patterns?**
Props drilling, Nested components definition, Mutating state directly, Index as Key, Huge UseEffects.

**Q293: How do you find memory leaks?**
Chrome Performance Monitor (Heap snapshots). Look for detached DOM nodes or uncleared intervals.

**Q294: How do you use React DevTools?**
Inspect Component hierarchy, view current Props/State, see who rendered whom (Profiler).

**Q295: How do you monitor performance?**
Web Vitals (LCP, FID, CLS). React Profiler (Render duration).

**Q296: How do you log user interactions?**
Analytics events (Click tracking). Error boundaries/Loggers for exceptions.

**Q297: How do you handle version upgrades?**
Read detailed changelog. Use codemods (Facebook provided scripts) to auto-update syntax.

**Q298: How do you refactor legacy React code?**
Incremental adoption. Convert Class to Functional one by one. Add tests before changing.

**Q299: How do you migrate class components to hooks?**
`state` -> `useState`. `componentDidMount` -> `useEffect`. `class methods` -> `const functions`.

**Q300: What makes a React application scalable?**
Folder structure, strict linting/typing (TypeScript), Testing, Component Reusability, State normalization.

---

## 🔹 31. React + JavaScript Deep Dive (Questions 301-310)

**Q301: How does JavaScript closure affect React hooks?**
Hooks rely on closures to capture state. Stale closures occur when an inner function captures an old variable value despite re-renders.

**Q302: How does event loop impact React rendering?**
React updates (re-renders) are synchronous computations (Stack) but can schedule tasks. Heavy rendering blocks the Event Loop, causing lag.

**Q303: How does React interact with the browser paint cycle?**
React batches DOM updates and applies them (Commit phase) typically before the browser paints the next frame.

**Q304: How does garbage collection affect React memory usage?**
Uncleared event listeners, detached DOM nodes, and large state objects held in closures prevent GC, causing leaks.

**Q305: How does React handle microtasks vs macrotasks?**
State updates inside Promises (microtasks) are processed immediately after the current script, before the next event loop tick/render.

**Q306: How does React work with Intersection Observer?**
Used for Lazy Loading images or Infinite Scroll. A `ref` is observed. When it enters viewport, state is updated to render content.

**Q307: How does Resize Observer help React apps?**
Allows components to respond to their own size changes (container queries) rather than just window resize events.

**Q308: How does requestIdleCallback improve performance?**
Allows scheduling non-essential work (e.g., analytics logs) when the browser is idle/not busy rendering.

**Q309: How does React behave in low-memory devices?**
Large component trees and frequent allocations can cause crashes. Virtualization and memory-efficient data structures usage is critical.

**Q310: How do you optimize React for slow networks?**
Code splitting, Service Workers (Caching), Compression (Brotli), and optimistic UI (instant feedback before server response).

---

## 🔹 32. Advanced Internals & Systems (Questions 311-325)

**Q311: How does React prioritize updates internally?**
Using a priority queue (Scheduler). User Blocking (Input) > Normal (Data fetch) > Idle (Logs).

**Q312: What is lane-based scheduling in React?**
A bitmask system where updates are assigned to specific "lanes" (bits). React processes the highest priority lane first.

**Q313: How does React Fiber manage interruption?**
It checks `shouldYield()` periodically. If the frame budget (5ms) is exceeded, it pauses work and yields back to main thread.

**Q314: How does React batch updates across async boundaries?**
Before React 18: No auto-batching in async. React 18+: Auto-batching everywhere using `createRoot`.

**Q315: How does concurrent rendering affect consistency?**
It produces consistent UI by working on a separate tree in memory. Only complete/valid trees are committed to the DOM.

**Q316: How would you design a design-system in React?**
Define tokens (Colors, Typography), primitives (Box, Text), and composite components (Button, Modal). use Storybook for documentation.

**Q317: How do you manage shared components across teams?**
Publish as NPM packages. Use Bit or Monorepo tools (Nx/Turborepo) to share code.

**Q318: How do you version frontend libraries?**
Semantic Versioning (SemVer). Major (Breaking), Minor (Feature), Patch (Fix).

**Q319: How do you manage breaking changes?**
Deprecation warnings in console. Codemods to automate migration. Major version bump.

**Q320: How do you enforce UI consistency at scale?**
Design tokens, Linters (Stylelint), strict TypeScript interfaces, and Visual Regression Testing (Storybook/Chromatic).

**Q321: How does React differ from React Native internally?**
React renders to HTML DOM. React Native uses the Bridge/Fabric to render to Native Views (UIView/AndroidView).

**Q322: How does React work with Web Components?**
React can render Web Components (custom elements). But passing data *out* (events) required manual listeners until React 19 improved support.

**Q323: How does React integrate with WebAssembly?**
Wasm modules can be imported and called from React components for heavy computation (e.g., Image processing).

**Q324: What is React Server Components (RSC)?**
Components that run *only* on the server. They send zero JS bundles to client, only serialization. (e.g., Next.js App Router).

**Q325: How does React compare with signals-based frameworks (Solid/Preact)?**
React uses Virtual DOM + Comparison. Signals use fine-grained reactivity (updates precise DOM nodes on value change without tree diffing).



