# 🗣️ Theory — Components & Props
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are the differences between functional and class components?"

> *"Functional components are just JavaScript functions that accept props and return JSX. Class components are ES6 classes that extend React.Component with a render method. Before hooks, functional components couldn't manage state or handle lifecycle — that was class component territory. Hooks changed everything. Now functional components can do everything class components can, with much less code. The React team recommends functional components for all new code. The practical differences today: class components use this to access props and state — which causes binding issues with event handlers. Functional components use closures. Class components have explicit lifecycle methods like componentDidMount and componentDidUpdate. Functional components use useEffect to handle lifecycle. Functional components are easier to test, easier to reason about, and easier to compose with custom hooks. You'll still see class components in older codebases, so you should be able to read them."*

---

## Q: "What is lifting state up and when do you do it?"

> *"Lifting state up means moving state from a child component to a parent component so that multiple children can share and synchronize around that state. You do it when two sibling components need to reflect the same changing data. The pattern: identify the common parent — the lowest component that is an ancestor of all the components that need the data. Move the state there. Pass the data down as props to the components that need to read it. Pass setter functions down as callback props to the components that need to change it. The principle is: the parent owns the state, the children are controlled. The downside is prop drilling — having to pass props through multiple layers — which is why global state solutions exist. But for nearby siblings, lifting state is the correct simple solution."*

---

## Q: "What is component composition and why is it preferred over inheritance?"

> *"React explicitly favors composition over inheritance. In class-based OOP, you'd extend a base class to add behavior. In React, you compose components by nesting them. The most common patterns are containment — where a component uses children to render arbitrary nested content, like a Card or Panel that doesn't know what's inside it — and specialization — where a specific component renders a more generic one with specific props, like an AlertDialog that uses a generic Dialog. The children prop is the gateway to containment — it lets you create layout components that wrap whatever content you pass them. More flexible components accept a render prop or a named prop that can be a component or a function. Why avoid inheritance? Components are functions, not classes in the traditional sense — composing functions is more flexible and doesn't create rigid class hierarchies."*

---

## Q: "What are Higher-Order Components and what replaced them?"

> *"A Higher-Order Component is a function that takes a component and returns a new component with enhanced behavior — think of it as a component wrapper factory. The naming convention is withSomething. You'd use them for cross-cutting concerns: authentication checks, logging, data fetching for multiple components. The problems with HOCs: they create extra layers in the component tree — visible in DevTools as nested withAuth > withLogger > withData. They can collide on props if two HOCs use the same prop name. And they're harder to type correctly in TypeScript. Custom hooks are the modern replacement. Instead of wrapping your component, you call a hook inside it. useAuth(), useLogger(), useFetch() are functional replacements that are composable, don't affect the component tree, and are much easier to type. HOCs are still valid for wrapping third-party components you don't control, but for your own logic, custom hooks are preferred."*

---

## Q: "What is the difference between controlled and uncontrolled components for forms?"

> *"A controlled component ties the input's value to React state — you pass value and onChange, so React owns the value. Every keystroke updates state, which updates the input. This gives you full control: you can validate on every keystroke, format input, or prevent invalid characters. An uncontrolled component lets the DOM manage the value — you use defaultValue for the initial value and attach a ref to read the value when needed, like on form submit. Controlled is the React way and is recommended for most cases. Uncontrolled is used specifically for file inputs — because you can't set file input values programmatically — and when integrating non-React libraries. The modern form library React Hook Form is interesting because it uses uncontrolled inputs by default for performance — avoiding re-renders on every keystroke — but gives you the control of a managed form when you need it."*

---

## Q: "How do you share data between components that aren't parent-child?"

> *"You have a few options depending on how far apart the components are. For siblings — lift state to their common parent and pass down via props. For distant relatives — use Context for stable data, or a state management library like Zustand for frequently changing data. For components that need to communicate but aren't in the same tree at all — like a sidebar and a main content area — global state is the right call. The decision tree I use: can I lift state to the nearest common ancestor without too much prop drilling? Do that. Is the data stable and rarely changes? Context. Does it change frequently and performance matters? Zustand or Redux. Is it server data — something fetched from an API? React Query with its cache handles sharing that data automatically — multiple components can use the same query key and they'll all read from the same cache without duplicate fetches."*

---

## Q: "What is the key prop and why is it special?"

> *"The key prop is a special prop that React uses internally for reconciliation — it never reaches your component as a prop you can read. When you render a list, React uses keys to match elements between renders. If an item's position changes but its key stays the same, React knows to update the existing DOM node rather than unmount and remount it. If the key changes, React treats it as a completely different element and does a full remount — this is actually a useful technique for forcing a component to reset: changing its key destroys it and creates a fresh instance. The rules: keys must be unique among siblings — not globally. Use stable, unique IDs from your data — not array indexes. Index as key breaks when the list can be reordered or filtered, because the index changes but IDs don't. The visual symptom of bad keys is form inputs retaining values when you expected them to clear."*
