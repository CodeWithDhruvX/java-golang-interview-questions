## 🔹 1. React Basics

### Question 1: What is React?

**Answer:**
React is a **JavaScript library** developed by Facebook (Meta) for building user interfaces. It follows a **component-based architecture**, allowing developers to build encapsulated components that manage their own state and compose them to make complex UIs.

**Key Characteristics:**
*   **Declarative:** You describe *what* the UI should look like for a given state, and React handles the updates.
*   **Component-Based:** Builds UIs from small, reusable pieces.
*   **Learn Once, Write Anywhere:** distinct from frameworks like Angular, React is a library and can be used for Web (ReactDOM), Mobile (React Native), or VR.

---

### Question 2: Why use React instead of other frameworks?

**Answer:**
1.  **Virtual DOM:** Ensures fast rendering by standardizing updates.
2.  **Reusable Components:** Promotes DRY (Don't Repeat Yourself) code.
3.  **Unidirectional Data Flow:** Makes data changes predictable (Parent -> Child).
4.  **Rich Ecosystem:** Huge community, many libraries (Redux, React Router).
5.  **Flexibility:** It's a library, not a full framework, so you can choose your own stack (routing, state management).

---

### Question 3: What are the main features of React?

**Answer:**
*   **JSX:** Syntax extension for writing HTML in JS.
*   **Virtual DOM:** Efficient DOM manipulation.
*   **Components:** Building blocks of the UI.
*   **One-way Data Binding:** Data flows down from parent to child.
*   **Hooks:** Functional state management (introduced in v16.8).

---

### Question 4: What is JSX?

**Answer:**
JSX stands for **JavaScript XML**. It allows us to write HTML-like syntax directly within JavaScript code. Under the hood, Babel trans-compiles JSX into standard `React.createElement()` calls.

**Example:**
```jsx
// JSX
const element = <h1>Hello, World!</h1>;

// Compiled JS
const element = React.createElement('h1', null, 'Hello, World!');
```

---

### Question 5: Can we use React without JSX?

**Answer:**
**Yes**, it is possible, but not recommended due to verbosity. You would need to use `React.createElement` manually.

**Example:**
```javascript
const e = React.createElement;
return e('div', { className: 'greeting' }, 'Hello World');
```

---

### Question 6: What is the virtual DOM?

**Answer:**
The Virtual DOM (VDOM) is a lightweight, in-memory representation of the real DOM. It is a JavaScript object that mirrors the structure of the UI.

---

### Question 7: How does the virtual DOM work?

**Answer:**
1.  **Render:** When state changes, React creates a new VDOM tree.
2.  **Diffing:** It compares the new VDOM tree with the previous one.
3.  **Reconciliation:** It calculates the minimum number of changes required.
4.  **Commit:** It updates only the changed nodes in the Real DOM.

---

### Question 8: Difference between Real DOM and Virtual DOM?

**Answer:**

| Feature | Real DOM | Virtual DOM |
| :--- | :--- | :--- |
| **Speed** | Slow to update (heavy layout recalculations) | Fast updates (just JS objects) |
| **Updates** | Updates the entire tree or requires manual targeting | Updates only changed elements (Diffing) |
| **Memory** | Heavy (browser internal structures) | Lightweight |

---

### Question 9: What is a React component?

**Answer:**
A component is a self-contained, reusable code block that divides the UI into smaller pieces. It accepts inputs (Props) and returns a React element (usually via JSX) describing what should appear on the screen.

---

### Question 10: Types of components in React?

**Answer:**
1.  **Functional Components:** Just JavaScript functions. Stateless (before Hooks). Simplified syntax.
2.  **Class Components:** ES6 classes extending `React.Component`. Have State and Lifecycle methods.

---

### Question 11: Difference between functional and class components?

**Answer:**
**Functional Component:**
```jsx
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
```
*   Uses Hooks (`useState`, `useEffect`).
*   No `this` keyword.
*   Simpler, less boilerplate.

**Class Component:**
```jsx
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```
*   Uses `Lifecycle Methods`.
*   Requires `this` binding for handlers.

---

### Question 12: What are props?

**Answer:**
**Props** (properties) are read-only inputs passed from a parent component to a child component. They allow components to be dynamic and reusable.

**Example:**
```jsx
<Greeting name="Alice" /> // 'Alice' is passed as a prop
```

---

### Question 13: What is state?

**Answer:**
**State** is a built-in object that allows components to create and manage their own data. Unlike props, state is **mutable** and fully controlled by the component. Changing state triggers a re-render.

---

### Question 14: Difference between state and props?

**Answer:**

| Feature | Props | State |
| :--- | :--- | :--- |
| **Source** | Passed from Parent | Managed internally by Component |
| **Mutability** | Immutable (Read-only) | Mutable (via `setState`) |
| **Role** | Configuration / Data passing | Internal Data storage |

---

### Question 15: Why is React fast?

**Answer:**
React is fast primarily because of the **Virtual DOM** and the **Reconciliation** algorithm. It batches updates and minimizes "Direct DOM Manipulation," which is the most expensive operation in web rendering.

---

## 🔹 2. Components & Props

### Question 16: How do you pass data from parent to child?

**Answer:**
By adding attributes to the child component tag.

**Example:**
```jsx
// Parent
<Child message="Hello form Parent" />

// Child
function Child(props) {
  return <p>{props.message}</p>;
}
```

---

### Question 17: Can a child component modify props?

**Answer:**
**No.** Props are **immutable** (read-only). If a child needs to modify the data, the parent must pass a function (callback) to the child, which the child calls to request a change in the parent's state.

---

### Question 18: How do you pass data from child to parent?

**Answer:**
Using **Callback Functions**. The parent passes a function as a prop, and the child calls that function with data.

**Example:**
```jsx
// Parent
const handleData = (data) => console.log(data);
<Child onAction={handleData} />

// Child
props.onAction("Data from Child");
```

---

### Question 19: What is props drilling?

**Answer:**
Props drilling is the process of passing data through multiple layers of components that don't need the data themselves, just to reach a deeply nested child component.
Parent -> Intermediate -> Intermediate -> Child.

---

### Question 20: How to avoid props drilling?

**Answer:**
1.  **Context API:** Defines a global state accessible by any component in the tree.
2.  **Redux / Global Store:** Centralized state management.
3.  **Component Composition:** Passing components as children or props (`<Layout content={<Widget />} />`).

---

### Question 21: What are default props?

**Answer:**
They allow you to set default values for props if one is not passed by the parent.

**Example:**
```jsx
function Button({ label = "Click Me" }) {
  return <button>{label}</button>;
}
// Or using defaultProps (Legacy)
Button.defaultProps = { label: "Click Me" };
```

---

### Question 22: What is `PropTypes`?

**Answer:**
A library used for **runtime type checking** of props.

**Example:**
```jsx
import PropTypes from 'prop-types';

function User({ age }) { ... }

User.propTypes = {
  age: PropTypes.number.isRequired
};
```

---

### Question 23: What is the `key` prop and why is it important?

**Answer:**
A `key` is a unique string attribute used when rendering lists of elements. It helps React identify which items have changed, been added, or removed. This ensures efficient DOM updates and prevents bugs in component state.

---

### Question 24: Can we use index as a key?

**Answer:**
**Yes, but only if:**
1.  The list is static (items purely display).
2.  The list is never reordered or filtered.
3.  Items have no internal state.

**Why it's bad:** If items are reordered, the index changes, leading React to incorrectly identify elements, causing state glitches and performance issues.

---

### Question 25: What happens if `key` is not provided?

**Answer:**
React will throw a warning in the console. It will default to using the array **index** as the key, which can lead to the issues mentioned in Question 24.

---

## 🔹 3. State & Lifecycle (Class + Hooks)

### Question 26: What is component lifecycle?

**Answer:**
The series of events that happen from the creation of a component to its removal.
Three main phases:
1.  **Mounting:** Inserting into the DOM.
2.  **Updating:** Re-rendering due to props/state changes.
3.  **Unmounting:** Removing from the DOM.

---

### Question 27: Explain React lifecycle methods.

**Answer:**
Methods available in Class Components:
*   `constructor()`: Init state.
*   `render()`: Render UI.
*   `componentDidMount()`: Side effects (API calls).
*   `componentDidUpdate()`: Reacting to changes.
*   `componentWillUnmount()`: Cleanup.

---

### Question 28: What is `componentDidMount()`?**

**Answer:**
Invoked immediately after a component is mounted (inserted into the tree). It is the perfect place for:
*   API Calls.
*   Setting up subscriptions.
*   Manipulating DOM elements.

---

### Question 29: What is `componentDidUpdate()`?

**Answer:**
Invoked immediately after updating occurs. Not called for the initial render.
Useful for performing DOM operations or network requests when the component updates (e.g., fetching new user data when `userID` prop changes).

---

### Question 30: What is `componentWillUnmount()`?

**Answer:**
Invoked immediately before a component is unmounted and destroyed. Used for cleanup tasks like:
*   Invalidating timers (`clearInterval`).
*   Canceling network requests.
*   Removing event listeners.

---

### Question 31: Why lifecycle methods are not used in functional components?

**Answer:**
Functional components were originally stateless. With the introduction of **Hooks** (specifically `useEffect`), we can perform side effects and lifecycle-like logic without needing classes or specific lifecycle methods.

---

### Question 32: What are React Hooks?

**Answer:**
Hooks are functions introduced in React 16.8 that allow functional components to "hook into" React features like state and lifecycle methods without writing a class.

---

### Question 33: Why were hooks introduced?**

**Answer:**
1.  **Reuse logic:** Custom hooks allow sharing stateful logic between components easily (unlike HOCs/Render Props).
2.  **Simplicity:** Functional components are less verbose than classes.
3.  **Confusion:** Removes the need to understand `this` binding in JS classes.

---

### Question 34: What is `useState`?

**Answer:**
A Hook that lets you add state to functional components. It returns an array with two elements: the current state value and a function to update it.

**Example:**
```jsx
const [count, setCount] = useState(0);
```

---

### Question 35: What is `useEffect`?

**Answer:**
A Hook that performs side effects in function components. It serves the same purpose as `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount`.

**Example:**
```jsx
useEffect(() => {
  document.title = `Count: ${count}`;
}, [count]); // Runs when 'count' changes
```

---

### Question 36: Difference between `useEffect` and lifecycle methods?

**Answer:**
*   **Lifestyle Methods:** Split logic based on *when* it runs (Mount vs Update vs Unmount).
*   **useEffect:** Groups logic based on *what it is doing* (e.g., a subscription). It handles mounting, updating, and unmounting in a single API.

---

### Question 37: How many times `useEffect` runs?**

**Answer:**
It depends on the **Dependency Array** (second argument):
1.  No dependency array: Runs after **every** render.
2.  `[]` (Empty array): Runs **only once** (on Mount).
3.  `[prop, state]`: Runs on Mount and whenever `prop` or `state` changes.

---

### Question 38: What is cleanup function in `useEffect`?

**Answer:**
It is the function returned by the effect callback. React runs it before the component unmounts (to clean up) and before re-running the effect (to clean up the previous run).

**Example:**
```jsx
useEffect(() => {
  const timer = setInterval(tick, 1000);
  return () => clearInterval(timer); // Cleanup
}, []);
```

---

### Question 39: What is `useLayoutEffect`?

**Answer:**
Identical to `useEffect`, but valid synchronously **after** all DOM mutations but **before** the browser paints. It is used for reading layout (width/height) from the DOM and synchronously re-rendering to prevent visual flickering.

---

### Question 40: Can hooks be used inside loops or conditions?

**Answer:**
**No.** Hooks must always be called at the **top level** of the function. This ensures that hooks are called in the exact same order on every render, which is critical for React to persist state correctly between renders.

---

## 🔹 4. Hooks (Advanced)

### Question 41: What is `useContext`?

**Answer:**
A hook that accepts a context object (the value returned from `React.createContext`) and returns the current context value. It simplifies consuming Context compared to the old `<Context.Consumer>` wrapper.

**Example:**
```jsx
const theme = useContext(ThemeContext);
```

---

### Question 42: What problem does `useContext` solve?

**Answer:**
It solves **Props Drilling**. It allows you to access global data (like User Auth, Theme, Language) anywhere in the component tree without passing props down manually through every level.

---

### Question 43: What is `useRef`?

**Answer:**
`useRef` returns a mutable ref object (`current` property) that persists for the full lifetime of the component.
**Common uses:**
1.  Accessing DOM elements directly.
2.  Storing mutable variables that **do not cause re-renders** when updated.

---

### Question 44: Difference between `useRef` and `useState`?

**Answer:**
*   **useState:** Updating `state` triggers a re-render. Used for data that affects the UI.
*   **useRef:** Updating `ref.current` **does not** trigger a re-render. Used for values needed behind the scenes (timers, DOM refs).

---

### Question 45: What is `useMemo`?

**Answer:**
A hook that **memoizes** a computed value. It only recomputes the value when one of its dependencies has changed. This is used to optimize performance by avoiding expensive calculations on every render.

**Example:**
```jsx
const expensiveValue = useMemo(() => computeSlowly(a, b), [a, b]);
```

---

### Question 46: What is `useCallback`?

**Answer:**
A hook that returns a **memoized callback function**. It only changes if one of the dependencies has changed. It is useful when passing callbacks to optimized child components (like those wrapped in `React.memo`) to prevent unnecessary re-renders.

---

### Question 47: Difference between `useMemo` and `useCallback`?

**Answer:**
*   `useMemo`: Returns the **result** of calling the function. (Memoizes a Value).
*   `useCallback`: Returns the **function** itself. (Memoizes a Function).
*   `useCallback(fn, deps)` is equivalent to `useMemo(() => fn, deps)`.

---

### Question 48: What is custom hook?

**Answer:**
A custom hook is a JavaScript function whose name starts with "use" and that may call other Hooks. It is a mechanism to extract and share **stateful logic** between components.

**Example:**
```jsx
function useFetch(url) {
  const [data, setData] = useState(null);
  useEffect(() => { ... }, [url]);
  return data;
}
```

---

### Question 49: When should you create a custom hook?

**Answer:**
When you find yourself duplicating logic (specifically logic involving state or effects) across multiple components. Common examples: Form handling, API fetching, Window listeners, Auth status.

---

### Question 50: Rules of Hooks?

**Answer:**
1.  **Only Call Hooks at the Top Level:** Don't call Hooks inside loops, conditions, or nested functions.
2.  **Only Call Hooks from React Functions:** Call them from React function components or custom Hooks (not regular JS functions).
