# 🟡 **Intermediate (2-4 Years): Advanced Hooks, Context, & Forms - Part 1**

### 41. What is `useContext`?
"`useContext` is a React Hook that lets you directly read and subscribe to a Context from your component.

It provides a way to pass data through the component tree without having to manually pass props down step-by-step through every single nested layer (solving prop drilling). 

I use it by first creating a Context object with `React.createContext()`, wrapping my app (or a section of it) in a `<ThemeContext.Provider value={theme}>`, and then in any deeply nested child component, I just call `const theme = useContext(ThemeContext)` to instantly grab that value."

#### Indepth
Prior to Hooks, consuming context required a cumbersome "Render Props" pattern using `<Context.Consumer>`. `useContext` drastically simplifies this by returning the context value synchronously during the render. However, any component calling `useContext` will *always* re-render when the provided context value changes, fundamentally bypassing `React.memo` performance optimizations.

---

### 42. What problem does `useContext` solve?
"The primary problem `useContext` solves is **prop drilling**.

Prop drilling occurs when you have to pass data through intermediate components that don't actually need the data, just to get it to a child component that does. This tightly couples the intermediate components to data they have no business caring about, making code refactoring difficult.

For example, if the top-level `<App>` needs to pass current user details to a deeply nested `<ProfileAvatar>`, it would manually pass `user={user}` through `<Layout>`, `<Sidebar>`, and `<Profile>` first. Context skips all that."

#### Indepth
While excellent for "global-ish" data (like an authenticated user, current UI theme, or chosen language), Context is **not a replacement for global state management** (like Redux or Zustand) in highly dynamic applications. Because Context causes *all* consumers to re-render when its value changes, placing rapidly changing data (like high-frequency socket data) into a Context Provider at the root level will cause catastrophic performance issues across the entire application.

---

### 43. What is `useRef`?
"`useRef` is a hook that allows me to store a mutable value that *does not cause a re-render* when updated.

It returns an object with a single property: `current`. I use it for two main reasons:
1. **Accessing DOM Elements:** Passing it to a `ref` attribute (like `<input ref={inputRef} />`) lets me directly access the DOM node to call methods like `.focus()` or measure its dimensions.
2. **Storing Mutable "Instance" Data:** Keeping track of a timer ID from `setInterval`, or storing the Previous State to compare on the next render, without triggering another render loop when I change it."

#### Indepth
A `useRef` is functionally identical to defining an instance property inside a class component (e.g., `this.myVar`). The reference persists across the entire lifecycle of the component. Modifying `ref.current` is completely synchronous and bypasses React's render scheduling entirely, which is why it's crucial never to read or write to `ref.current` during the bare *Render Phase* (the main body of your functional component) as it makes the behavior unpredictable.

---

### 44. Difference between `useRef` and `useState`?
"They both store data across renders, but their reaction to data changes is exactly opposite.

When I update a `useState` variable using its setter function, **React schedules a re-render** of the component to update the UI with the new data.

When I update a `useRef` variable by manually changing `myRef.current = newValue`, **React completely ignores the change**. It does *not* re-render the component.

If the value governs what you see on the screen, use `useState`. If the value is strictly for internal logic (like a timer ID or tracking if a component just mounted), use `useRef`."

#### Indepth
Attempting to force UI updates by mutating a ref and expecting it to reflect on the screen is a common anti-pattern. React enforces immutable state updates to remain predictable. If a function's return value depends on mutable variables outside its scope (like reading a ref during the render block), it ceases to be a Pure Function, leading to "tearing" bugs in Concurrent Mode where two renders might display different values.

---

### 45. What is `useMemo`?
"`useMemo` is a performance optimization hook used to memoize (cache) the result of an expensive calculation.

I pass it two arguments: an expensive function that computes a value, and an array of dependencies. React will call the function on the first render, and on subsequent renders, it will only recalculate the value if one of the dependencies has changed. If the dependencies haven't changed, React skips the calculation and instantly returns the cached result.

I use this when I'm sorting massive lists of data or doing heavy mathematical processing right inside the render cycle."

#### Indepth
Overusing `useMemo` is a very common mistake. There is a memory and CPU overhead to checking the dependency array on every render. If you are just mapping over an array of 50 basic strings, `useMemo` is likely slower than just recreating the array. You should only use it when profiling proves a specific calculation is causing jank (usually >1ms execution time).

---

### 46. What is `useCallback`?
"`useCallback` is nearly identical to `useMemo`, but it is used to **memoize a function reference itself**, rather than the returned value.

In JavaScript, creating a new function on every render `const fn = () => {}` means the function reference changes every time. Normally, this doesn't matter. But if I pass that function down as a prop to a heavily optimized child component (like one wrapped in `React.memo`), the child will incorrectly re-render every time because it sees the 'prop' as a new function reference.

`useCallback` caches that function reference between renders unless its dependencies change."

#### Indepth
A frequent misconception is that `useCallback` makes the function run faster. It doesn't. It only caches the memory address of the function declaration. Its sole purpose is to preserve referential equality when passing functions into dependency arrays (like in a child's `useEffect`) or when passing props to memoized components to prevent useless re-renders.

---

### 47. Difference between `useMemo` and `useCallback`?
"Both hooks optimize performance by caching something based on dependency changes.

- `useMemo` calls your function and **caches the resulting value**. (e.g., You give it a function calculating taxes, it caches the number `450.50`).
- `useCallback` does not call anything. It simply **caches the function itself**. (e.g., You give it the tax function, it returns a stable reference to that exact function so you can pass it to children).

Under the hood, `useCallback(fn, deps)` is literally just syntactic sugar for `useMemo(() => fn, deps)`."

#### Indepth
When deciding which to use, always ask: "Am I trying to save CPU power on this math operation?" -> `useMemo`. "Am I trying to prevent a child component from re-rendering because it thinks its `onClick` prop changed?" -> `useCallback`.

---

### 48. What is a custom hook?
"A **custom hook** is simply a JavaScript function whose name starts with `use` (like `useWindowSize` or `useFetchData`) and that calls other React Hooks inside of it.

Custom hooks are the absolute best way to extract and share stateful logic between completely different components. For example, if both my Header and Sidebar need to know if the user is online, instead of duplicating `useState` and `window.addEventListener('online')` logic in both, I extract it all into a single `useOnlineStatus()` custom hook.

They let me separate the *logic* of a component from its *UI markup*."

#### Indepth
A custom hook doesn't share state *values* between components (like Context or Redux does); it only shares the stateful *logic*. If five components call `useOnlineStatus()`, there are five distinct, completely isolated instances of state running simultaneously. 

---

### 49. When should you create a custom hook?
"I create a custom hook whenever I notice I am copy-pasting the same `useEffect` or complex `useState` logic across multiple components.

Common examples include:
- Fetching API data (handling loading/error/success states).
- Interacting with Browser APIs (local storage, window dimensions, geolocation).
- Managing complex form validation rules.
- Handling debouncing or throttling.

If a component file is getting massive because of tangled data-fetching logic, ripping all that out into a `useEntityData()` hook instantly makes the component incredibly clean and readable."

#### Indepth
The "Rules of Hooks" apply to custom hooks just legally as they do inside components. Their name *must* start with `use`. This is a convention strictly enforced by the React linting tools to guarantee that hooks are not accidentally called dynamically inside loops or conditions by standard utility functions.

---

### 50. Rules of Hooks?
"The 'Rules of Hooks' are the two fundamental laws you absolutely must follow when writing React functional components:

1. **Only call Hooks at the Top Level:** Never call a hook inside a loop, an `if` condition, or a nested function. This guarantees that hooks are called in the exact same order every time the component renders, which is how React internally associates state values with the correct hook instances.
2. **Only call Hooks from React Functions:** You can only call them inside standard React functional components or inside your own custom hooks. You cannot call them from regular vanilla JavaScript functions.

I always rely on the `eslint-plugin-react-hooks` linter to enforce these automatically."

#### Indepth
Understanding *why* the rules exist clarifies React's internals. React doesn't use "magic" to track state; it uses a hidden array inside the Fiber node. It increments an index pointer (`cursor = 0`) on every hook call (`cursor++`). If an `if` statement skips a hook on render #2, the cursor gets misaligned. The second hook call on render #2 grabs the state meant for the third hook from render #1, resulting in catastrophic UI bugs.

---

### 51. What is `useReducer`?
"`useReducer` is an alternative hook to `useState` for managing more complex state logic.

It is heavily inspired by Redux. Instead of calling generic 'set' functions to change a value, you dispatch **actions** (like a string `'INCREMENT'` or an object `{type: 'ADD_USER'}`).

You supply a **reducer function**—a pure function that takes the current `state` and the `action`, figures out what happened, and returns the entirely new state object. It's incredibly useful when the next state depends heavily on the previous state, or when the state is a complex object with nested properties."

#### Indepth
The signature is `const [state, dispatch] = useReducer(reducerFn, initialState)`. Unlike Redux, `useReducer` state is highly localized to the component calling it (it does not exist in a global store). However, you can elegantly combine `useReducer` with `useContext` to tightly couple state logic and provide it across the app, acting essentially as a lightweight, boilerplate-free alternative to Redux for smaller apps.

---

### 52. Difference between `useState` and `useReducer`?
"`useState` is simple and direct. I use it for independent primitive values like tracking a string `username` or a boolean `isModalOpen`. Calling the setter function is imperative ('Set this value to 5').

`useReducer` is built for complex state logic. I use it when multiple pieces of state change together (e.g., clicking 'Fetch' sets `isLoading: true, data: null, error: null` all at once). Dispatching actions is declarative ('Here is what the user did: FETCH_START', but the reducer decides *how* the data changes).

If I find myself writing 5 different `useStates` that always update together sequentially, I refactor to a single `useReducer`."

#### Indepth
Technically, inside the React core codebase, `useState` is actually implemented *using* `useReducer`. `useState` is just a basic wrapper around a simple reducer that says "replace the old state with exactly whatever new value was passed to the setter function."

---

### 53. What is `useImperativeHandle`?
"It is an escape hatch hook I use quite rarely. It allows a child component to explicitly customize the 'instance' (the ref) that it exposes to a parent component.

Usually, passing a `ref` from a parent to a child just gives the parent the raw underlying DOM node (like the actual HTML `<input>`). But sometimes, I don't want the parent to have full access. Using `useImperativeHandle` (combined with `forwardRef`), the child can restrict the parent to only calling specific functions, like `.focus()` or `.resetCounter()`, hiding the private internal variables.

It breaks the declarative, top-down flow of React, so I only use it for imperative actions like managing focus, text selection, or triggering complex third-party animations."

#### Indepth
The name says it all: it is for *imperative* actions. Calling methods directly on children from parents strongly resembles object-oriented jQuery patterns. It tightly couples components unnecessarily. It is the absolute last resort when lifting state up or passing a standard prop simply cannot achieve the desired timing or behavior.

---

### 54. What is `forwardRef`?
"`forwardRef` is a wrapper function used around a functional component that allows it to receive a `ref` prop from its parent and 'forward' it down to a specific DOM node inside its own return output.

By default, functional components cannot take a `ref` attribute (because they don't have instances). If a parent tries `<MyCustomInput ref={inputRef} />`, nothing happens.

If I wrap `MyCustomInput` using `React.forwardRef((props, ref) => ...)`, the functional component suddenly accepts the `ref` as its second argument, and I can attach it directly to the native `<input ref={ref} />` element inside."

#### Indepth
A common oversight is omitting `displayName` when using `forwardRef`. Because it's an anonymous HOC wrapping a function, it shows up in React DevTools confusingly simply as `ForwardRef`. Always explicitly assign `MyCustomInput.displayName = 'MyCustomInput'` immediately after defining it to keep debugging tools readable.

---

### 55. What is `useId`?
"`useId` is a hook introduced in React 18 for generating unique IDs that are stable across both the server rendering and the client hydration process.

I use it primarily for accessibility features in forms. For example, explicitly connecting a `<label htmlFor={id}>` to an `<input id={id}>`.

Instead of manually generating random numbers or Math.random (which can cause mismatch errors when Server-Side Rendering output differs from Client output), I just call `const id = useId()`, and React guarantees the string returned is unique and safe to use as an HTML ID attribute."

#### Indepth
`useId` is not designed for generating keys to map over lists data. You should always use the stable inherent ID (like a database ID sequence) from your dataset for keys. Using `useId` to generate a key inside a component that renders list items gives them unique keys *locally*, but on re-render, those keys reset, destroying list performance heuristics anyway.

---

### 56. How does event handling work in React?
"Event handling in React is very similar to DOM elements but with some core syntax changes.

1. I use camelCase for event names (e.g., `onClick` not `onclick`).
2. I pass a function reference wrapped in curly braces (e.g., `onClick={handleClick}`) instead of a string.
3. React automatically prevents the default browser behavior if explicitly called via `e.preventDefault()` inside the handler; returning `false` does not work like older JS.

Behind the scenes, React intercepts these events and wraps them in its own cross-browser 'SyntheticEvent' system to guarantee identical behavior across all browsers."

#### Indepth
React implements a robust **Event Delegation** pattern. Rather than attaching a unique `addEventListener` to every single button on a screen, React attaches exactly *one* event listener to the root container where your application renders. This single listener catches all bubbling events on the page, identifies the source React element, and calls the appropriate handler prop assigned to it.

---

### 57. How to handle forms in React?
"In React, I almost exclusively use **Controlled Components** to handle form inputs.

This means I intercept everything the user types exactly when they type it, store it in React State using `useState`, and manually push that state value back into the input field. The pattern looks like this:
`<input value={name} onChange={(e) => setName(e.target.value)} />`

When the user clicks submit, I attach an `onSubmit` handler to the `<form>` tag, call `e.preventDefault()`, and gather all the current state variables to confidently send to my backend API."

#### Indepth
For massive, sprawling enterprise forms with dozens of validations and nested fields, linking 50 individual `useStates` leads to sluggish rendering because every keypress triggers a full component re-render. In those scenarios, I drop standard controlled inputs and reach for specialized open-source libraries like **React Hook Form** or Formik, which utilize uncontrolled ref-based patterns to minimize re-renders while still providing excellent validation APIs.

---

### 58. How to handle multiple form inputs?
"Managing five different `useState` variables for one form is tedious. Instead, I manage all inputs simultaneously using a single state **object**.

1. I create state: `const [formData, setFormData] = useState({ name: '', email: '' })`
2. I give every input a `name` attribute that perfectly matches a key in that object (`name="email"`).
3. I write a single, universal change handler using Computed Property Names:
```javascript
const handleChange = (e) => {
  const { name, value } = e.target;
  setFormData(prev => ({ ...prev, [name]: value }));
};
```
Now, attaching `onChange={handleChange}` to fifty different inputs just works automatically."

#### Indepth
While this dramatically reduces boilerplate, spreading objects on every keystroke (`...prev`) can still cause slight performance penalties in profoundly deep forms. One trap is forgetting the closure rules inside the handler—always use the callback form of set state (`prev => ...`) when updating objects, otherwise rapid keystrokes might overwrite each other due to stale closure variables referencing outdated state object references.

---

### 59. How to validate forms?
"I handle form validation directly within my component's logic depending on when I need user feedback:

1. **Immediate Validation:** I check criteria inside the `onChange` handler while they type. If `e.target.value.length < 5`, I instantly display a red error message.
2. **On Submit:** Inside my `onSubmit` handler, before sending data to the server, I check all my current state variables. If validation fails, I return early and display errors near the inputs.

For complex enterprise validation rules (regex testing, matching passwords, async backend checks), I always use **Yup** or **Zod** schema validators alongside a library like React Hook Form to keep my component clean."

#### Indepth
HTML5 constraint validation attributes (like `required`, `pattern`, `minLength`) are still excellent to use inside React markup because they hook into innate browser accessibility features (like screen readers audibly announcing a field is invalid). However, developers shouldn't solely rely on default browser UI modals as they cannot be fully styled via CSS; intercept the error via state and display custom UI messages instead.
