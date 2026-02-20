# 🟢 **Freshers (0-2 Years): State, Lifecycle, and Basic Hooks - Part 2**

### 22. What are default props?
"**Default props** allow you to set default values for the `props` argument in a React component. 

If a parent component doesn't pass a specific prop, the component will use the default value instead of `undefined`.

In class components, or older functional components, I would use the `ComponentName.defaultProps = {}` syntax. However, in modern React using functional components, I simply use standard JavaScript ES6 default parameters during object destructuring: `function MyComponent({ title = 'Default Title' })`. This approach is cleaner and what the React team recommends."

#### Indepth
The `defaultProps` property is officially deprecated for functional components in recent React versions (React 18.3+ warns against it). While it still functionally works, the community and React core team have completely shifted toward JavaScript native default arguments because it plays much nicer with TypeScript inference.

---

### 23. What is PropTypes?
"`PropTypes` is a mechanism in React used for type-checking props passed to a component.

Before TypeScript became the standard, we used `PropTypes` to ensure a component received the correct data type (e.g., `PropTypes.string` or `PropTypes.func.isRequired`). If a parent passed the wrong type—like a number instead of a string—React would throw a warning in the browser console during development.

Today, I rarely use `PropTypes` because I use **TypeScript**, which catches these errors at compile time rather than runtime, providing a much safer and faster developer experience."

#### Indepth
PropTypes was originally built into the core React library but was extracted into its own package (`prop-types`) in React 15.5. It relies strictly on runtime validation. In production builds, PropTypes checks are stripped away for performance reasons, meaning they only help developers catch bugs during active development.

---

### 24. What is component lifecycle?
"The **component lifecycle** refers to the sequence of events that happen from the moment a React component is created to the moment it is destroyed.

It consists of three main phases:
1. **Mounting:** When the component is first inserted into the DOM.
2. **Updating:** When the component's state or props change, causing an update/re-render.
3. **Unmounting:** When the component is removed from the DOM.

Understanding this lifecycle is crucial because it dictates *when* I can fetch data, subscriptions, or manipulate the DOM securely without causing memory leaks or errors."

#### Indepth
In the era of class components, these phases mapped strictly to specific methods (like `componentDidMount`, `componentDidUpdate`, `componentWillUnmount`). In the modern functional component era, this entire lifecycle—specifically the side effects associated with these phases—is managed primarily by a single hook: `useEffect`.

---

### 25. Explain React lifecycle methods (Class components).
"In older class-based React, we had specific methods that fired at different stages of the component's life:

1. **`constructor()`**: Runs first. Used for initializing state and binding methods.
2. **`render()`**: Required method that returns JSX. Must be pure (no side effects).
3. **`componentDidMount()`**: Runs *once* immediately after the component is added to the DOM. I use this for initial API calls or setting up timers.
4. **`componentDidUpdate(prevProps, prevState)`**: Runs after every re-render (except the first). I use this to react to prop or state changes, like fetching new data based on a changed ID.
5. **`componentWillUnmount()`**: Runs *once* right before the component is destroyed. I use this to clean up timers or WebSockets to prevent memory leaks."

#### Indepth
There are also less common methods like `shouldComponentUpdate` (used for manual performance optimization before `React.memo`), `getDerivedStateFromProps`, and `getSnapshotBeforeUpdate`. The React team deprecated several older lifecycles (like `componentWillMount` and `componentWillReceiveProps`) in React 16.3 because they led to unsafe practices in concurrent rendering models.

---

### 26. Why are lifecycle methods not used in functional components?
"Lifecycle methods are exclusive to ES6 Classes because they are methods on the class instance. Functional components are exactly that—functions. There is no 'instance' object (`this`) to attach those methods to.

Instead, the React team introduced **Hooks** in version 16.8 to solve this.

With Hooks, specifically the `useEffect` hook, I can tap into the exact same lifecycle events (mounting, updating, unmounting) without needing class methods. In fact, `useEffect` is often better because it allows me to group related logic together (like subscribing and unsubscribing) rather than splitting it across different lifecycle methods."

#### Indepth
The shift from classes to functions is fundamentally a shift from an Object-Oriented mental model to a Functional Programming one. Class lifecycles often forced developers to split related logic—e.g., adding an event listener in `componentDidMount` and removing it in `componentWillUnmount`. `useEffect` allows that setup and teardown to exist side-by-side in the same block.

---

### 27. What are React Hooks?
"**React Hooks** are special functions that allow me to 'hook into' React's core features—like state management and lifecycle methods—from inside a Functional Component.

Before Hooks, if a functional component needed state, I had to rewrite it as a complex Class component. Hooks completely changed this.

The most common hooks I use daily are `useState` (to hold data that changes) and `useEffect` (to perform side effects like fetching data). They make components significantly shorter, easier to read, and simpler to test."

#### Indepth
Hooks revolutionized React by allowing logic reuse through "Custom Hooks". Instead of using complex patterns like Higher-Order Components (HOCs) or Render Props to share logic between classes, developers can extract stateful logic into a simple JavaScript function starting with `use-` and share it effortlessly across components.

---

### 28. Why were hooks introduced?
"Hooks were introduced to solve several major pain points in React development at the time:

1. **Complex Components:** Class components became massive, tangled messes because related logic (like fetching a user and subscribing to their status) had to be split across different lifecycle methods (`componentDidMount`, `componentDidUpdate`, `componentWillUnmount`).
2. **Logic Reuse:** Sharing stateful logic between components required "wrapper hell" (HOCs or render props), which bloated the component tree. Hooks solved this with Custom Hooks.
3. **The `this` Keyword:** JavaScript's `this` context is confusing and error-prone. Functional components avoid `this` entirely.

Hooks allow us to write cleaner, more composable logic."

#### Indepth
Another underlying, highly technical reason for Hooks was minification. Class methods do not minify well, and building tools to analyze how class instances change over time is difficult. Functional components are simply functions, which modern compilers and minifiers can aggressively optimize and tree-shake much better than object-oriented class hierarchies.

---

### 29. What is `useState`?
"`useState` is a React Hook that lets you add state to a functional component.

When I call `useState`, I pass it the initial value I want the state to have. It returns an array with two elements: the current state value, and a function that lets me update it.

```javascript
const [count, setCount] = useState(0);
```

Whenever I call the updater function (`setCount`), React notes that the state has changed and schedules a re-render of the component to update the UI."

#### Indepth
It's critical to remember that state updaters via `useState` are asynchronous. Calling `setCount(count + 1)` and then immediately checking `console.log(count)` on the next line will show the *old* value, because the state update hasn't fired the new render frame yet. If the new state depends on the previous state, you should pass a callback function: `setCount(prev => prev + 1)`.

---

### 30. What is `useEffect`?
"`useEffect` is a React Hook that lets you perform **side effects** in functional components. 

Side effects are anything that reaches outside the React component ecosystem—like fetching data from an API, directly updating the DOM document title, or setting up a timer.

It essentially replaces `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount` from class components all wrapped into one unified API. I tell React: *'After you finish rendering the UI, go run this piece of code.'*"

#### Indepth
`useEffect` takes two arguments: a callback function containing the effect logic, and an optional dependency array. If the dependency array is omitted, the effect runs after *every* render. If it's an empty array `[]`, it runs only *once* on mount. If it contains variables `[id, name]`, the effect recalculates and runs *only* if those specific variables have changed between renders.

---

### 31. Difference between `useEffect` and lifecycle methods?
"While they accomplish similar goals, they require different mental models.

Lifecycles methods like `componentDidMount` are tied to **time**—they ask *when* a component is in its life cycle.

`useEffect` is tied to **synchronization**—it asks *what data changed* to require an effect to run.

Additionally, `useEffect` allows me to group related logic logically. I don't have to split my event listener setup into `componentDidMount` and the cleanup into `componentWillUnmount`. `useEffect` handles the setup, and its return function handles the teardown right next to it."

#### Indepth
`useEffect` fundamentally shifts the paradigm from imperative "do this when the component mounts" to declarative "synchronize this system with the current React state." This forces developers to define their dependencies explicitly, which helps React optimize exactly when an effect needs to run and prevents stale closure bugs that were common in class methods accessing mutable `this` states.

---

### 32. How many times does `useEffect` run?
"It entirely depends on the **dependency array** (the second argument).

1. **No array passed (`useEffect(() => {...})`):** It runs after *every single render*.
2. **Empty array passed (`useEffect(() => {...}, [])`):** It runs exactly *once*, immediately after the component mounts.
3. **Array with variables (`useEffect(() => {...}, [x, y])`):** It runs on mount, and then runs again *only* if the values of `x` or `y` have changed since the last render.

Knowing how to use this array is critical to preventing infinite loops caused by state updates inside an effect."

#### Indepth
In React 18 Strict Mode during development, components are intentionally mounted, unmounted, and instantly re-mounted again to flush out bugs related to missing cleanup functions. Because of this, developers often incorrectly assume their `useEffect(..., [])` is running twice in production. It only runs twice in development Strict Mode.

---

### 33. What is the cleanup function in `useEffect`?
"The cleanup function is what you optionally `return` from inside the `useEffect` callback.

```javascript
useEffect(() => {
  const timer = setInterval(() => console.log('tick'), 1000);
  return () => clearInterval(timer); // Cleanup function!
}, []);
```

React runs this cleanup function right before the component unmounts. However, if the effect has dependencies and runs multiple times, React also runs the cleanup from the *previous* render before applying the effect for the *next* render. This ensures that subscriptions or event listeners aren't duplicated."

#### Indepth
Failing to return a cleanup function for things like WebSockets, DOM event listeners, or Intervals results in egregious memory leaks because the background task continues running long after the component UI has been destroyed. This is a very common source of performance degradation in poorly written single-page applications.

---

### 34. What is `useLayoutEffect`?
"It is similar to `useEffect`, but the timing of when it fires is different.

`useEffect` is asynchronous. It fires *after* React has updated the DOM and the browser has painted the screen. This is best for 99% of use cases (like API calls) because it doesn't block the visual update.

`useLayoutEffect` is synchronous. It fires *after* React updates the DOM, but *before* the browser actually paints the pixels on the screen. I use this exclusively for measuring DOM nodes (like getting the scroll position or the exact width of a `div`) so I can make quick calculations before the user ever sees a flicker."

#### Indepth
Because `useLayoutEffect` blocks the browser's painting process, using it heavily will make your app feel sluggish and unresponsive. The official React documentation strongly advises defaulting to `useEffect` and only reaching for `useLayoutEffect` when you notice a visual jump or layout shift caused by `useEffect` firing too late.

---

### 35. Can hooks be used inside loops or conditions?
"**Absolutely not. This violates the 'Rules of Hooks.'**

Hooks must *always* be called at the top level of your functional component. You cannot place them inside `if` statements, `for` loops, or nested functions.

The reason is internal to React. React relies on the exact **call order** of hooks to associate a specific state or effect with a specific `useState` or `useEffect` call. If a hook is put inside a condition that evaluates to false on the second render, the hook order shifts, and React throws a catastrophic error because it loses track of which state belongs where."

#### Indepth
The internal implementation of hooks is analogous to a singly linked list (or an array). React maintains a cursor that increments with each hook call. On re-render, it walks this same list in the identical order. If the number of hooks changes mid-render due to an `if` statement, the cursor misaligns, pointing the third hook's data to the fourth hook's logic. Using the `eslint-plugin-react-hooks` linter is mandatory to enforce these rules.

---

### 36. What is a synthetic event?
"A **Synthetic Event** is React's cross-browser wrapper around the native browser events (like `click`, `change`, `submit`).

Different browsers (Chrome, Firefox, Safari) used to handle native events slightly differently, which caused bugs. React normalizes these events into a single API that works identically across all platforms. 

For me as a developer, I use it exactly like a normal DOM event—I can call `e.preventDefault()` or read `e.target.value`—but under the hood, React is managing the complexities of browser compatibility."

#### Indepth
Before React 17, Synthetic Events were "pooled". This meant React reused the event objects for performance, nullifying their properties immediately after the callback finished. This meant you couldn't access `e.target.value` asynchronously inside a `setTimeout` without calling `e.persist()`. In Modern React (v17+), event pooling has been removed entirely, and events behave identically to native asynchronous closures.

---

### 37. Difference between HTML events and React events?
"There are three main differences between standard HTML events and React Synthetic events:

1. **Naming:** HTML events are all lowercase (e.g., `onclick`, `onchange`). React events use camelCase (e.g., `onClick`, `onChange`).
2. **Value Binding:** In HTML, you pass a string to the event attribute (`onclick="handleClick()"`). In React, you pass a function reference inside curly braces (`onClick={handleClick}`).
3. **Preventing Default:** In HTML, you can return `false` to prevent the default behavior of an element. In React, you *must* explicitly call `e.preventDefault()`.

These changes make React's event system more aligned with JavaScript conventions."

#### Indepth
Another major architectural difference is **Event Delegation**. If you attach an `onChange` to 50 input fields in standard HTML/JS, the browser registers 50 separate event listeners. In React, it registers exactly *one* global event listener at the root of the document. When an input is clicked, the event bubbles up to the root, and React's internal system dispatches it to the correct synthetic event handler, saving significant memory.

---

### 38. How to bind events in React?
"In modern functional components, you don't really have to 'bind' events the way you did in classes. You simply pass the function reference directly:

`<button onClick={handleSave}>Save</button>`

In older class components, the `this` keyword was notoriously problematic. If you passed a class method as an event handler, it lost its binding to the class instance (making `this` undefined). To fix it, you either had to `this.handleSave = this.handleSave.bind(this)` in the constructor or use Arrow Functions for the class methods, as arrow functions automatically inherit `this` from their surrounding lexical scope."

#### Indepth
While arrow functions (`onClick={() => handleSave(id)}`) are extremely convenient in functional components when you need to pass an argument to your handler, they technically create a *new* function instance on every single render. Usually, React is fast enough that this doesn't matter. But if you are passing this arrow function down as a prop to a heavily optimized child component (like one wrapped in `React.memo`), the child will incorrectly re-render every time because the function reference changed. In those rare scenarios, you wrap the handler in `useCallback`.

---

### 39. What are controlled components?
"A **controlled component** is an input element (like an `<input>`, `<textarea>`, or `<select>`) whose value is entirely controlled by React state.

React becomes the 'single source of truth' for that input.
1. The component's `state` determines the `value` attribute of the input.
2. The `onChange` event listener triggers a state update (`setState`) whenever the user types.

I use controlled components 95% of the time because it allows me to perform immediate validation (like disabling a submit button until a valid email is typed) and instantly format input (like forcing uppercase text)."

#### Indepth
The code pattern looks like this: `<input value={name} onChange={e => SetName(e.target.value)} />`. Because the UI strictly reflects the data model state, the input is physically incapable of showing data that React doesn't know about. This synchronization ensures bugs where the DOM and the app state disagree are impossible.

---

### 40. What are uncontrolled components?
"An **uncontrolled component** is an input element that manages its own internal state using the native DOM API, without using React state.

Instead of binding a `value` and an `onChange` handler, I use a **`ref`** (`useRef`) to grab the value directly from the DOM node only exactly when I need it (like when a user finally clicks 'Submit').

I use them very rarely. They are primarily useful when I'm integrating React with a non-React library that expects direct DOM access, or if I have a massive form where linking 50 distinct `useStates` to 50 inputs is causing performance lag."

#### Indepth
Uncontrolled components are created by assigning a `ref` attribute to the input element. If you need to give the input an initial default value, you use the `defaultValue={initialData}` prop instead of `value={state}`, allowing the DOM to handle all subsequent keystrokes on its own. While slightly faster rendering-wise, they completely break the declarative nature of React forms.
