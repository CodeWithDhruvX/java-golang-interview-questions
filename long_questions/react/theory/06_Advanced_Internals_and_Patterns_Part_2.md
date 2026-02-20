# 🔴 **Advanced (4-6 Years): Component Patterns & Testing - Part 2**

### 91. How does Context API work internally?
"Under the hood, Context API works by maintaining a stack of values.

When you render a `<Provider value={data}>`, React pushes that `data` value onto a Context Stack during the render phase. As React physically traverses down the component tree to render the children nested underneath it, those children components sit in a 'Context Scope'. 

If a deeply nested child calls `useContext(MyContext)`, React simply looks at the top of that specific Context Stack and returns the current value. When the traversal finishes rendering that branch and moves back up, it pops the value off the stack."

#### Indepth
Context essentially utilizes React's internal Fiber node structure. The `Provider` node in the Fiber tree holds the `memorizedState`. Any descendant Fiber node that calls `useContext` adds itself to an internal linked list of 'Context Dependencies'. When the Provider's `value` eventually changes, React iterates through that specific dependency list and explicitly forces those specific descendant Fibers to be marked for re-render, efficiently bypassing memoized intermediate components.

---

### 92. What causes Context consumers to re-render?
"A component subscribing to a Context will **always** re-render if the `value` prop provided by the nearest `<Provider>` above it changes.

React determines if the `value` 'changed' by running an `Object.is()` strict equality check on the old value versus the new value. 

This leads to a massive performance trap: if I pass an object dynamically created during the Provider's render (like `<Provider value={{ user, theme }}>`), the `{}` syntax creates a brand new memory allocation on *every single render* of the Provider. Even if `user` and `theme` didn't change, the object reference did. React sees a 'new' value, and it violently forces every single Consumer in the entire app to re-render simultaneously."

#### Indepth
The standard solution to this object instability is memoizing the Context value: `<Provider value={useMemo(() => ({ user, theme }), [user, theme])}>`. Now, the object reference remains identical across renders unless `user` or `theme` physically changes, successfully preventing catastrophic re-renders across the consumer tree.

---

### 93. How do you split Contexts for performance?
"The best way to prevent unnecessary re-renders when using Context is to physically separate data that changes often from data that changes rarely into entirely distinct Context Providers.

Instead of one massive `<AppContext value={{ user, posts, theme, uiState }}>`, which forces the `<UserAvatar>` component to re-render every time someone clicks a different `uiState` tab, I split them up.

I create `<UserContext>`, `<ThemeContext>`, and `<UIContext>`, nesting them. 
Then, a component only subscribes to the specific Context it cares about, remaining immune to the rendering chaos of the others. Context should be fractured logically by mutation frequency borders."

#### Indepth
For extremely complex global state where fracturing into 10 separate contexts becomes difficult to maintain, developers should completely abandon the native Context API and reach for **Zustand** or **Redux**. These atomic/selector-based libraries inherently 'split' the subscriptions under the hood, allowing a component to flawlessly subscribe to `state.foo` while ignoring changes in `state.bar` within the exact same global store.

---

### 94. What is a Higher-Order Component (HOC)?
"A **Higher-Order Component** is an advanced React pattern (primarily used before Hooks existed) used for reusing component logic.

Technically, an HOC is a **function that takes a component as an argument and returns a brand new, enhanced component.**

For example, `const withAuth = (WrappedComponent) => { ... }`. I would pass my generic `<Dashboard>` into `withAuth(Dashboard)`. The HOC function checks if the user is logged in. If they are, it renders the Dashboard. If not, it redirects them to the login page. It physically wraps the original component in a shell of reusable logic."

#### Indepth
HOCs were the absolute backbone of React logic-sharing for years (e.g., Redux's `connect()`, React Router's `withRouter()`). They often caused "Wrapper Hell" in the React DevTools, nesting a component under fifty deep `WithRouter(Connect(WithTheme(MyComponent)))` layers. Today, Custom Hooks have replaced HOCs for almost all use cases because they don't bloat the component tree, though HOCs still exist for specialized interception logic like `React.memo` or `React.forwardRef`.

---

### 95. What is the Render Props pattern?
"**Render Props** is a technique for sharing code between React components using a prop whose value is a function.

Instead of a component rendering its own hardcoded UI, it delegates the rendering responsibility back to its parent. The component simply manages the complex state logic, and then calls a `render` function passed to it as a prop, passing the state back up to be visualized.

For example, a `<MouseTracker>` component calculates the X/Y cursor position, but instead of rendering it, it calls `props.render({x, y})`. The parent decides what that data looks like: `<MouseTracker render={({x, y}) => <h1>Cursor is at {x}, {y}</h1>} />`."

#### Indepth
Like HOCs, Render Props were a crucial pattern for logic reuse before hooks (e.g., React Router's old `<Route render={() => ...}>` or Formik's `<Formik>{({ values }) => ...}</Formik>`). Custom Hooks (`const { x, y } = useMouseTracker()`) have largely obsoleted this pattern because they achieve the exact same logic separation without requiring a nested callback function pyramid of doom in the JSX.

---

### 96. What is the Container-Presentational pattern?
"This is an architectural pattern (popularized by Dan Abramov) separating components into two distinct categories:

1. **Container Components:** These care about *how things work*. They fetch API data, manage Redux state, handle complex logic, and contain practically zero CSS DOM markup. 
2. **Presentational Components (Dumb Components):** These care about *how things look*. They contain no state, do not fetch data, and exist purely to take props and render HTML/CSS.

The Container fetches the data and passes it down purely as props to the Presentational component to render. This makes the UI completely decoupled, reusable, and trivially easy to unit test."

#### Indepth
While still an excellent mental model for structuring large applications, the rigid separation is less strictly enforced today. Hooks mathematically uncoupled state logic from the component rendering lifecycle natively. A single functional component can now cleanly invoke `useFetchUser()` at the top and instantly render the UI below it, merging the Container and Presentational concepts cleanly without excessive file splitting.

---

### 97. What is headless component architecture?
"A **Headless Component** is a component that contains absolutely no UI whatsoever (no HTML, no CSS)—it exclusively provides logic, state management, and accessibility attributes.

You typically consume it via a custom hook or a Render Prop. The most famous examples are libraries like **React Table** or **Headless UI**.

The headless library handles the incredibly difficult math of calculating column widths, sorting algorithms, and aria-labels for a data grid. It returns the raw data and the necessary event handlers to me. I simply take those handlers and attach them to my own highly customized HTML `<table>`, giving me total 100% control over the visual styling while outsourcing the complex behavioral logic."

#### Indepth
Headless architecture is the absolute gold standard for modern generic component library design. Traditional rigid libraries (like Material UI or Ant Design) force their specific CSS implementation onto you, making override customization a nightmare. Headless libraries provide the "brain" (state machines, accessibility) but leave the "head" (the actual DOM elements) completely up to the developer's chosen design system (like TailwindCSS).

---

### 98. What is Unit Testing in React?
"**Unit testing** in React focuses on isolating and testing the smallest parts of the application independently—typically testing a single Custom Hook or a single 'dumb' Presentational Component in a vacuum.

If I have a function that formats dates, or a button component that accepts a `color` prop, a unit test verifies that given an input, the exact expected output is produced. 

I use **Jest** to run the tests and assert that `expect(result).toBe(true)`. For component rendering, I use tools like React Testing Library to mount the component in memory instantly and check if the specifically requested text actually appeared in the fake DOM."

#### Indepth
Pure functions are trivially easy to unit test. Components heavily tied directly to Redux stores, API fetch calls, or global contexts are notoriously difficult to unit test because you have to mock out the entire universe around them. This is why the Container-Presentational pattern remains relevant—it extracts complex logic (hard to unit test) away from the pure UI rendering (easy to unit test).

---

### 99. What is React Testing Library (RTL)?
"**React Testing Library** is the industry-standard library for testing React components. It completely replaced the older library 'Enzyme'.

RTL's core philosophy is: *'The more your tests resemble the way your software is used, the more confidence they can give you.'*

Instead of testing the internal implementation details of a component (like checking if a specific `useState` variable mutated to 'true' or finding an element by its internal CSS class name), RTL tests what the incredibly powerful End User actually sees. I write tests that look like: `expect(screen.getByText('Submit')).toBeVisible()`. I click the button, I don't care *how* the developer coded the state array, I just verify the new element appeared on screen."

#### Indepth
Testing implementation details (like Enzyme did by allowing you to inspect component instance variables) leads to brittle tests. If you refactor a class component to a functional component with hooks, the UI behaves identically for the user, but the Enzyme test utterly shatters because the internal variable names changed. RTL forces you to test behavioral outcomes, ensuring tests survive aggressive codebase refactors gracefully.

---

### 100. How do you test hooks?
"Testing custom hooks requires executing them, but React strictly forbids calling hooks completely outside of a component body.

To test a hook directly without building a meaningless dummy component just to host it, I use a specialized utility function called `renderHook` (provided natively by `@testing-library/react` in modern versions).

I call `const { result } = renderHook(() => useMyCustomCounter())`. I can then assert its initial state `expect(result.current.count).toBe(0)`. If I need to test an updater function, I wrap it in RTL's `act()` utility: `act(() => result.current.increment())`, and then verify the new state `expect(result.current.count).toBe(1)`."

#### Indepth
The `act()` function is critically important. In standard browser environments, React's rendering engine batches internal DOM updates seamlessly. In an artificial Jest testing environment, updating state might finish instantly while the DOM updates lag behind. Wrapping interaction code in `act()` guarantees that all React updates, effects, and DOM synchronizations are entirely resolved before the assertion engine checks the DOM on the next line of test code.
