# 🗣️ Theory — React Basics & JSX
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is React and why is it popular?"

> *"React is a JavaScript library for building user interfaces, created by Facebook — now Meta. It's popular for a few key reasons. First, the component model — you build your UI as a tree of reusable, composable components, which makes large codebases manageable. Second, the unidirectional data flow — data flows down through props, events flow up through callbacks — which makes state predictable and debugging easier. Third, the Virtual DOM — React efficiently updates only what changed rather than re-rendering the whole page, keeping UIs fast. And fourth, the massive ecosystem — React Router, Redux, Next.js, React Query, and thousands of libraries all work with it. It's not a full framework like Angular, which means you have flexibility to choose your tools, but you also need to make more architectural decisions yourself."*

---

## Q: "What is JSX and how does it work?"

> *"JSX stands for JavaScript XML — it's a syntax extension that lets you write what looks like HTML inside JavaScript. It's not valid JavaScript by itself — a compiler like Babel transforms it into function calls. When you write angle bracket syntax in a JSX file, Babel compiles it to React.createElement or, with the modern transform in React 17+, to calls from the jsx-runtime package. JSX has a few rules you need to know: every JSX expression must have a single root element — you can use a Fragment if you don't want an extra DOM node. Attribute names follow the JavaScript convention — className instead of class, htmlFor instead of for. You inject JavaScript expressions with curly braces. JSX is not required to use React — you could call createElement directly — but JSX is what everyone uses because it's far more readable."*

---

## Q: "What's the difference between a React Element and a Component?"

> *"A React Element is just a plain JavaScript object — it's the description of what should be on screen. It's what React.createElement returns: an object with a type, props, and key. It's not a DOM node — it's lightweight data. A Component, on the other hand, is a function — or class — that accepts props and returns elements. So a component is a factory for elements. When you write JSX with an uppercase tag like UserCard, React calls your function with the props as an argument and gets back an element. When you write lowercase like div, React treats it as a built-in DOM element. This distinction matters practically: you always capitalize custom components because React uses that to decide whether to call a function or create a native DOM node."*

---

## Q: "How does conditional rendering work in React?"

> *"Conditional rendering in React is just regular JavaScript conditionals applied to JSX. The most common patterns are: if-else for early returns — if a component shouldn't render, just return null or a fallback component near the top of the function. Ternary operator for inline conditions — condition ? ShowThis : ShowThat. Logical AND — condition && ShowThis — for showing something only when true. A gotcha with AND: if your condition is the number zero, React will actually render a zero character, not nothing. So be careful with falsy values that aren't false or null — use double negation or explicit comparison. For more complex conditions, extract the logic into a variable before the return statement — it keeps the JSX clean and readable."*

---

## Q: "Why do we need React Fragments?"

> *"React requires that every component returns a single root element — you can't return two sibling elements from a component. The naive solution is wrapping them in a div, but that adds an unnecessary DOM node, which can break CSS layouts — especially flexbox or grid. Fragments solve this by grouping elements without adding any real DOM node. The JSX shorthand is the empty angle bracket syntax. There are two versions: the full React.Fragment and the short syntax. The important difference is that only the full version supports the key attribute — which matters when you're mapping over a list and returning multiple elements per item. In that case, you need React.Fragment with a key prop because the short syntax doesn't accept props at all."*

---

## Q: "What is props and what's the key rule about props?"

> *"Props — short for properties — are how you pass data from a parent component to a child. In JSX, props look like HTML attributes. Inside the child component, they're received as a single plain object, and people typically destructure them in the function signature. The most fundamental rule about props is that they are read-only — the component that receives props must never modify them. This enforces unidirectional data flow. If the parent changes the prop value and re-renders, the child re-renders with the new value. If the child needs to influence the parent — like reporting a user action — it does so through a callback prop that the parent passes down. This top-down, event-up pattern is what keeps React data flow predictable."*

---

## Q: "How does React render to the DOM and what changed in React 18?"

> *"React rendering happens in two phases: reconciliation and commit. Reconciliation is React figuring out what changed — running your component functions, building a new Virtual DOM tree, and diffing against the previous one. The commit phase is actually applying those changes to the real browser DOM. In React 17 and earlier, you called ReactDOM.render to mount your app. React 18 introduces createRoot — you call ReactDOM.createRoot on your container element, then root.render. This is the gateway to all React 18 features — Automatic Batching, Concurrent Mode, Suspense improvements. Using the legacy render API still works but opts you out of everything new. The practical way to think about it: createRoot is the React 18 way, and you'd use it for all new projects."*

---

## Q: "What is Create React App vs Vite and which should I use?"

> *"Create React App was the official starter for React projects for years. It uses Webpack under the hood, has a preset configuration you can't easily change without ejecting, and is famously slow — dev server startup of 30 seconds or more on larger projects isn't unusual. Vite is the modern alternative — it uses esbuild for dependency pre-bundling and serves source files as native ES modules in development, so dev server startup is near-instant and hot module replacement is nearly instant too. For production, Vite uses Rollup which produces highly optimized bundles. In 2024, Vite is the clear recommendation for new React projects. The React team actually deprecated Create React App in the docs. For full-stack apps, Next.js is worth considering because it adds SSR, routing, and more. For pure client-side SPAs, use Vite."*
