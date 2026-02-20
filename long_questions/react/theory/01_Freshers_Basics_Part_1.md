# 🟢 **Freshers (0-2 Years): React Basics and Components - Part 1**

### 1. What is React?
"React is an open-source, front-end JavaScript library developed by **Facebook** for building user interfaces or UI components.

It's used primarily for building single-page applications (SPAs) where the data changes over time without reloading the page. Instead of manipulating the DOM directly, React uses a **declarative approach**, allowing me to describe what the UI should look like for a given state, and React handles the updates efficiently.

I love React because it's component-based. I can build encapsulated components that manage their own state, then compose them to make complex UIs, making code reusable and easier to maintain."

#### Indepth
React is fundamentally a **view-layer** library (the 'V' in MVC), not a full framework like Angular. It requires additional libraries for routing (like React Router) and state management (like Redux or Zustand). Its core strength lies in its **Virtual DOM** and efficient diffing algorithm (Reconciliation), which minimizes expensive direct DOM manipulations.

---

### 2. Why use React instead of other frameworks (like Angular or Vue)?
"I prefer React because of its **flexibility and large ecosystem**.

Unlike Angular, which is a full-fledged, opinionated framework, React is just a library. This means I can choose the exact tools I want for my project—whether it's the router, state manager, or styling solution. 

Furthermore, React uses **JSX** (JavaScript XML), which allows me to write HTML structures directly within JavaScript. This colocation of logic and markup feels very intuitive to me. Also, the community support is massive; if I hit a roadblock, someone has likely already solved it."

#### Indepth
React's adoption of **unidirectional data flow** (one-way data binding) makes debugging state changes much more predictable compared to Angular's two-way data binding. Additionally, the introduction of **React Hooks** in version 16.8 revolutionized how developers write components, providing a cleaner, functional approach over cumbersome class-based lifecycle methods.

---

### 3. What are the main features of React?
"React has several key features that make it stand out:

1. **JSX (JavaScript Syntax Extension):** Allows writing HTML within JavaScript.
2. **Virtual DOM:** A lightweight copy of the real DOM used to optimize rendering.
3. **Component-Based Architecture:** UIs are built from reusable, isolated components.
4. **Unidirectional Data Flow:** Data flows down from parent to child via props, making state predictable.
5. **Declarative UI:** You describe the final state, and React updates the DOM to match it.

These features combined make React fast, scalable, and developer-friendly."

#### Indepth
Under the hood, React's rendering engine (like React DOM for web or React Native for mobile) is decoupled from its core reconciler. This architectural decision (React Fiber) allows React to be environment-agnostic. The unidirectional data flow ensures that changes in derived data don't unexpectedly mutate the source of truth, a common issue in complex UI architectures.

---

### 4. What is JSX?
"**JSX** stands for JavaScript XML. It's a syntax extension for JavaScript that looks very similar to HTML.

I use it because it makes writing React components significantly easier. Instead of writing complex `React.createElement` calls, I can write familiar HTML-like syntax to describe the UI structure.

However, browsers don't understand JSX natively. Before running in the browser, JSX is transformed into standard JavaScript objects by compilers like **Babel**."

#### Indepth
When you write `<div className="app">Hello</div>`, Babel transpiles this into `React.createElement('div', { className: 'app' }, 'Hello')`. Because JSX is ultimately just JavaScript expressions, you can embed JavaScript variables, functions, and logic directly inside the markup using curly braces `{}`.

---

### 5. Can we use React without JSX?
"Yes, it is absolutely possible to use React without JSX, though it is very rare.

To do so, I would have to use the `React.createElement(component, props, ...children)` function manually for every single element. 

For example, a simple heading would be `React.createElement('h1', null, 'Hello World')`. While it works, writing a complex component tree this way becomes incredibly verbose and hard to read, which is why almost everyone uses JSX."

#### Indepth
Understanding that JSX is syntactic sugar over `React.createElement` (or more recently, the `jsx-runtime` compiler functions introduced in React 17) is crucial. It explains why we used to have to `import React from 'react'` at the top of every file even if we didn't explicitly call `React` methods—the transpiled code needed the `React` object in scope.

---

### 6. What is the Virtual DOM?
"The **Virtual DOM (VDOM)** is a programming concept where an ideal, or 'virtual', representation of a UI is kept in memory.

Think of it as a lightweight JavaScript object that mirrors the actual DOM (Document Object Model) structure. The real DOM is slow to update because it triggers layout recalibrations and repaints in the browser. The Virtual DOM is extremely fast because it's just manipulating JavaScript objects without touching the screen.

React uses this Virtual DOM to figure out the most efficient way to update the real DOM."

#### Indepth
The Virtual DOM is not a feature unique to React, but React popularized it. It acts as an abstraction layer between the developer's declarative intent and the imperative DOM API. When state changes, a new Virtual DOM tree is generated. React then compares this new tree against the previous one (a process called diffing) to calculate the minimum set of mutations required.

---

### 7. How does the Virtual DOM work (Reconciliation)?
"When a component's state or props change, React follows a three-step process:

1. **Render:** React creates a new Virtual DOM tree representing the updated UI.
2. **Diffing:** React compares this new Virtual DOM tree with the previous one to find exactly what changed.
3. **Patching (Reconciliation):** Once it knows the differences, React updates **only those specific nodes** in the real DOM, leaving the rest untouched.

This batching and targeted updating is what gives React its performance edge."

#### Indepth
React's diffing algorithm operates in **O(n)** time complexity, which is remarkable given that traditional tree comparing algorithms are O(n^3). It achieves this via two heuristic assumptions: 
1. Elements of different types will produce different trees (React tears down the old tree completely).
2. Developers can hint at which child elements may be stable across renders using a `key` prop.

---

### 8. What is a React component?
"A **component** is the fundamental building block of a React application.

It's a reusable piece of code that represents a part of the user interface. Components can be as small as a single button or as large as an entire page layout. They accept inputs (called **props**) and return React elements detailing what should appear on the screen.

I treat components like JavaScript functions: they take data in, and output UI."

#### Indepth
React's component model heavily promotes composition over inheritance. You build complex UI by nesting smaller, simpler components. This separation of concerns ensures that UI logic is tightly coupled with its markup, making it easier to debug, test, and reuse across different parts of the application.

---

### 9. What are the types of components in React?
"There are two main ways to define components in React:

1. **Functional Components:** These are simple JavaScript functions that accept props and return JSX. With the introduction of Hooks, they can now manage state and side effects, making them the standard way to write React today.
2. **Class Components:** These are ES6 classes that extend `React.Component`. They require a `render()` method to return JSX and use `this.state` and lifecycle methods. 

I use functional components for almost everything now, as they are less wordy and easier to test."

#### Indepth
While Class components are considered legacy syntax, they are not deprecated. However, the React team strongly recommends Functional Components for all new code. Functional components avoid the complexities of the `this` keyword in JavaScript and allow logic to be shared more easily via Custom Hooks rather than Higher-Order Components (HOCs) or Render Props patterns typically used with classes.

---

### 10. Difference between Functional and Class components?
"The main differences are in syntax and how they handle state/lifecycles.

**Functional Components** are just plain JavaScript functions. They are simpler, use less code, and utilize **Hooks** (like `useState`, `useEffect`) to manage state and side effects. They don't have a `this` keyword context.

**Class Components** are ES6 classes. They are more verbose, require a constructor to initialize state, use the `this` keyword heavily, and manage side effects using specific lifecycle methods like `componentDidMount` and `componentDidUpdate`.

I prefer Functional Components because they lead to cleaner, more composable code."

#### Indepth
Performance-wise, functional components used to be slightly faster because they avoided the overhead of class instantiations, but modern React optimizes both well. The real advantage of functional components is the mental model: they are fundamentally closures that capture values from their render cycle, whereas classes reflect mutable instances over time, which often leads to subtle bugs in async callbacks accessing `this.state` or `this.props`.

---

### 11. What are props?
"**Props** (short for properties) are the mechanism for passing data from a parent component to a child component in React.

They act exactly like arguments passed to a JavaScript function. They are **read-only** (immutable) inside the child component, meaning the child cannot modify the props it receives; it can only read them and render UI based on them.

I use props constantly to configure components dynamically. E.g., `<Button text="Submit" color="blue" />` passes 'text' and 'color' as props."

#### Indepth
Under the hood, all attributes passed to a component are bundled into a single JavaScript object called `props`. React strictly enforces the "Pure Rule" regarding props: a component must not modify its own props. If data needs to change in response to user input, that data must be managed as `state`, not `props`.

---

### 12. What is state?
"**State** is an object that holds data specific to a component that may change over its lifetime. 

Unlike props, which are passed down from a parent and are read-only, **state is managed entirely within the component itself**. When a component's state changes, React automatically re-renders that component to reflect the new data.

For example, I would use state to track whether a dropdown menu is open, what a user typed into an input field, or the data fetched from an API."

#### Indepth
In functional components, state is initialized and updated using the `useState` or `useReducer` hooks. State updates are generally **asynchronous** and batched by React for performance optimization. This means you cannot immediately read the new state value on the very next line after calling a state updater function.

---

### 13. Difference between state and props?
"The key difference is ownership and mutability.

**Props** are passed to a component from its parent. They are **read-only** and cannot be changed by the receiving component. They used to configure the component externally.

**State** is managed internally by the component itself. It is **mutable**, meaning the component can update its own state using updater functions (like `setState` or `setSomething`), which triggers a re-render.

I think of props as function arguments and state as a function's local variables."

#### Indepth
While props read top-down (Unidirectional Data Flow), state allows components to be interactive. A common pattern is to hold state in a parent component and pass that state down to children as props, along with a callback function (also passed as a prop) that the child can call to ask the parent to update its state.

---

### 14. Why is React fast?
"React is fast primarily due to its implementation of the **Virtual DOM**.

Instead of interacting directly with the heavy real browser DOM on every data change, React calculates the differences between the current Virtual DOM and the newly rendered Virtual DOM. It then batches these necessary changes and applies them to the real DOM in a single, optimized operation.

Additionally, component-level state ensures that only the components whose data actually changed are re-rendered, rather than refreshing the whole page."

#### Indepth
React also leverages **Synthetic Events**, a cross-browser wrapper around the browser's native event system. Instead of attaching separate event listeners to every single DOM node, React attaches a single listener to the root of the document (Event Delegation). This significantly reduces memory consumption and improves execution speed, especially in large lists or tables.

---

### 15. How do you pass data from parent to child?
"I pass data from a parent component to a child component using **props**.

In the parent component, I add attributes to the child component's JSX tag. For example, `<UserProfile name={username} age={25} />`.

Then, in the `UserProfile` child component, I receive these values via the `props` object: `function UserProfile(props) { return <div>{props.name}</div>; }`. Or, more commonly, I destructure them right in the function signature: `function UserProfile({ name, age })`."

#### Indepth
Because props flow downwards, this enforces a predictable architecture. If a child component needs to communicate data *back up* to the parent, the parent must pass down a callback function as a prop, and the child executes that callback with the new data. This is often confusing for beginners but is crucial for React's one-way data flow.

---

### 16. Can a child component modify props?
"**No, a child component cannot modify the props it receives.**

Props are strictly **read-only** (immutable) inside the receiving component. React enforces this rule strictly. A component must act like a pure function with respect to its props; it should never mutate them.

If a child component needs to change the data it received via props (e.g., a user types into an input field whose value came from a prop), it must call a callback function passed down by the parent, asking the parent to update its state."

#### Indepth
Attempting to mutate props directly (e.g., `props.title = "New Title"`) violates React's core principles. In strict mode or specialized environments, the `props` object is usually frozen via `Object.freeze()`, meaning direct mutations will throw a JavaScript TypeError in strict mode. State must be lifted to the lowest common ancestor if siblings need to share or modify data.

---

### 17. How do you pass data from child to parent?
"Since data in React flows downwards (parent to child), passing data upwards requires a specific pattern using **callbacks**.

1. The parent component defines a function that updates its own state.
2. The parent passes this function down to the child component as a **prop**.
3. When an event occurs in the child (like a button click), the child calls that prop function, passing the necessary data as arguments.

This triggers the parent's function, updating the parent's state, and re-rendering the app with the new data."

#### Indepth
This pattern is sometimes called "Inverse Data Flow." It highlights that while *data* always flows down as props, *actions* (events) flow up via callbacks. If this pattern becomes too deeply nested across many component layers, it leads to "prop drilling," at which point Context API or global state management (Redux) becomes necessary.

---

### 18. What is props drilling?
"**Props drilling** is a situation where data is passed down from a top-level component to a deeply nested component through many intermediary components that do not actually need the data themselves.

For example, if Component A wants to pass user data to Component F, it might have to pass it through Components B, C, D, and E just to reach F.

While it's a normal part of React, excessive prop drilling makes the code hard to maintain, as every component in the chain must declare and pass the prop along."

#### Indepth
Prop drilling isn't inherently bad; it explicitly shows data dependencies. However, it tightly couples intermediate components to data they don't care about, making refactoring difficult. The primary solutions to avoid prop drilling are component composition (passing components as `children`), React Context API, or global state libraries like Redux or Zustand.

---

### 19. How to avoid props drilling?
"There are primarily three ways to avoid prop drilling in React:

1. **Context API:** I create a Context at the top level and 'Provide' the data. Any deeply nested component can 'Consume' that context directly using the `useContext` hook, bypassing the intermediary components entirely.
2. **Component Composition:** Instead of passing data down, I pass the *component itself* down using the `children` prop. The parent provides the data directly to the nested component before passing it down.
3. **State Management Libraries:** For complex applications, I use tools like Redux, Zustand, or Jotai to store data globally outside the component tree. Components simply connect to the store to get what they need."

#### Indepth
Component Composition is often overlooked but extremely powerful. If Component A renders Component B which renders Component C. Instead of A passing data to B, to pass to C. A can render `<B><C data={mydata} /></B>`. B then simply renders `{props.children}`. This keeps B completely oblivious to C's data requirements.

---

### 20. What is the `key` prop and why is it important?
"The `key` prop is a special string attribute you must include when creating arrays of elements in React, typically inside a `.map()` function.

Keys help React identify which items have changed, been added, or been removed. They give the elements a stable identity across renders.

Without unique keys, React doesn't know exactly which element changed in a list, so it re-renders the entire list from scratch, which destroys performance and can cause bugs with component state (like losing input focus)."

#### Indepth
When Diffing, React uses the `key` to match children in the original tree with children in the subsequent tree. If a key is stable, predictable, and unique (like a database ID), React can simply move the existing DOM node to the new position instead of destroying and recreating it. 

---

### 21. Can we use the array index as a key?
"Yes, you *can* use the array index as a key, but **you should generally avoid it**, especially if the list can be reordered, filtered, or prepended to.

If I use the index and then delete the first item in an array, the item at index 1 shifts to index 0. React sees the key '0' is still there, so it thinks the *element* hasn't changed, but its *contents* have. This forces React to mutate the existing DOM node instead of moving it, which is slower and can lead to incorrect state behaving unexpectedly (e.g., an input field retaining the value of a deleted row).

I only use index as a key if the list is completely static and will never change order."

#### Indepth
If no key is provided, React defaults to using the index as the key, and it will give you a warning in the console. The best practice is to always use a unique, deterministic identifier from your data object, such as a database ID or a UUID generated when the item was created.
