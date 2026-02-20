# 🟡 **Intermediate (2-4 Years): Lists, Performance, & Routing - Part 2**

### 60. How do you render lists in React?
"The standard way to render lists in React is to use the JavaScript `Array.prototype.map()` method.

I usually take an array of data, call `.map()` on it directly inside my JSX (inside curly braces `{}`), and return a React element for each item in the array. 

For example, if I have `const users = ['Alice', 'Bob']`, I render them like this:
`<ul> {users.map(user => <li key={user}>{user}</li>)} </ul>`"

#### Indepth
React explicitly handles arrays of elements. You do not need to `.join()` the array into a big string like in old-school jQuery templating. React internally knows how to iterate over an array of React Elements and render them sequentially. The absolute golden rule here is that every element returned from a `map()` *must* possess a unique `key` prop so React can identify it during the Reconciliation diffing phase.

---

### 61. Why is `map()` preferred over `forEach()`?
"`map()` is explicitly designed to transform an array into a slightly different array, returning the new array. `forEach()` simply executes a function on each element and returns `undefined`.

Because JSX expects an array of React elements (or a single element, string, or number) to render, `map()` perfectly creates that required array of JSX notes. 

If I used `forEach()`, it would return nothing, and nothing would render on the screen. To make `forEach()` work, I would have to declare an empty array outside the block, push JSX elements into it during the loop, and then return that array, which is incredibly verbose and goes against the declarative nature of React."

#### Indepth
`map()` is functionally pure (assuming the callback provided is pure), meaning it doesn't mutate the original array. Functional programming principles strongly influence React's design, making `map()`, `filter()`, and `reduce()` the go-to tools for manipulating data structures right before rendering.

---

### 62. What is conditional rendering?
"**Conditional rendering** in React is the ability to output different UI depending on the current state of the application. It works identically to how conditions work in normal JavaScript.

If a user is logged in, I show a 'Logout' button. If they are not logged in, I show a 'Login' button.

Since React components are just functions that return JSX, I can use standard JavaScript `if/else` statements, ternary operators, or logical `&&` operators exactly where the variables control the UI."

#### Indepth
React handles 'falsy' values gracefully in JSX. If you attempt to render `false`, `null`, `undefined`, or `true`, React simply ignores them and renders absolutely nothing. (Note: rendering the number `0` will physically render a '0' string to the DOM, not an empty space, which is a common buggy gotcha perfectly solved by using `!!variableName && <Component />` instead of just `variableName &&`).

---

### 63. Different ways to do conditional rendering?
"I use three primary ways, depending on how complex the condition is:

1. **`if/else` statements:** Best used *outside* the JSX return block to completely change what a whole component renders (like rendering a full `<LoadingScreen />` instead of the `<Dashboard />`).
2. **Ternary Operator (`condition ? true : false`):** Best used *inside* the JSX for rendering one of two small things (like `{isOnline ? <GreenDot /> : <RedDot />}`).
3. **Logical AND (`condition && <Component />`):** Best used when I only want to render something if the condition is true, and render exactly nothing if it is false (like `{showModal && <Modal />}`)."

#### Indepth
For extremely complex conditions where nested ternaries become illegible (often jokingly called 'ternary hell'), you can immediately invoke a function expression (IIFE) inside the JSX to use standard `switch` statements, or better yet, extract the complex logic into a separate small functional component that handles its own internal `if` statements clearly.

---

### 64. What is a fragment?
"A **Fragment**, written as `<React.Fragment>` or just `<>`, allows you to group a list of children elements together without adding an extra, meaningless DOM node (like a generic `<div>`) to the actual HTML output.

A React component's `return` statement must return exactly **one single parent element**. If I want to return two `<p>` tags next to each other, wrapping them in a `<div>` technically solves the problem, but it clutters the DOM.

Wrapping them in a Fragment solves the React requirement of a single parent, but simply dissolves away during the actual browser rendering, keeping the DOM extremely clean."

#### Indepth
The only crucial difference between the empty tag syntax `<></>` and the explicit `<React.Fragment></React.Fragment>` is that the explicit version supports the `key` prop. If you are mapping over an array and returning multiple elements simultaneously without a parent `div`, you *must* use `<React.Fragment key={item.id}>` to satisfy React's list key requirement because `<key={item.id}></>` is invalid syntax.

---

### 65. What causes re-rendering in React?
"There are only three core things that cause a component to re-render in React:

1. **State Changes:** Whenever `useState`'s setter or `useReducer`'s dispatch function is called.
2. **Prop Changes:** If a parent component passes down new props to a child component, the child will re-render.
3. **Parent Re-renders:** By default, whenever a parent component re-renders (for any reason), React recursively re-renders **all** of its children components entirely down the tree, regardless of whether their specific props changed."

#### Indepth
Context changes also cause a re-render. Any component calling `useContext(MyContext)` will forcefully re-render if the `Provider`'s `value` changes. The concept that "all children re-render when a parent re-renders" is fundamental. React doesn't know if the child's internal calculations rely on side effects of the parent, so it assumes the safest route is to recalculate everything unless explicitly told not to via `React.memo()`.

---

### 66. How do you prevent unnecessary re-renders?
"To prevent unnecessary re-renders (specifically the cascading re-renders caused by parents), I use three main optimization techniques:

1. **`React.memo`:** I wrap a child component in it. React will now aggressively check if the props actually changed. If they didn't, it *skips* re-rendering that child even if the parent re-rendered.
2. **`useMemo`:** I wrap complex objects or heavy calculations in the parent, so I'm not accidentally passing physically new object references to the child on every single render.
3. **`useCallback`:** I wrap function definitions in the parent, so I pass a stable function reference down to the child, rather than recreating the function every render."

#### Indepth
The absolute best way to prevent re-renders is structural: **lifting state down**. If a `<TextInput>` state changes constantly in a massive `<Dashboard>` parent component, forcing everything else to re-render, don't use `React.memo`. Instead, move the `useState` into the isolated `<TextInput>` component directly. State isolation is infinitely better than memoization.

---

### 67. What is `React.memo`?
"`React.memo` is a Higher-Order Component (HOC) wrapped around functional components to optimize performance.

When wrapping a component in `React.memo(MyComponent)`, React remembers (memoizes) the previously rendered UI of that component. The next time the parent renders, React does a shallow strict equality check (`===`) on all the props. 

If none of the props changed, React bypasses the render phase for that component completely and simply re-uses the last painted result, saving valuable CPU cycles."

#### Indepth
It is crucial to understand that `React.memo` only checks **props**. It does not freeze internal state. If a component is wrapped in `React.memo` but its internal `useState` updates, or a Context it subscribes to updates, it *will* re-render. Also, because it only does a *shallow* check, passing inline objects like `style={{ color: 'red' }}` as a prop will instantly break the memoization because that object has a new memory address on every render.

---

### 68. Difference between `React.memo` and `PureComponent`?
"They serve the exact same purpose—preventing unnecessary re-renders by doing a shallow comparison of props strings—but for different types of components.

**`PureComponent`** is an ES6 Class that you extend instead of `React.Component` when writing legacy class components.

**`React.memo`** is a Higher-Order Function used entirely for modern functional components.

Since I write functional components 100% of the time now, I solely use `React.memo`."

#### Indepth
A key distinguishing feature is that `React.memo` accepts a second argument: a custom comparison function `(prevProps, nextProps) => boolean`. This allows you to write deep-equality checks or ignore specific props. `PureComponent` hardcodes a shallow comparison and does not allow for a custom override (you would have to manually implement the `shouldComponentUpdate` lifecycle method instead).

---

### 69. What is lazy loading?
"**Lazy loading** is a performance optimization technique where you delay the initialization or downloading of an object (like a heavy component, an image, or a chunk of JavaScript code) until the exact moment the user actually needs it.

Instead of forcing the user's browser to download a massive 5MB bundle containing everything in the app on the initial load, I only send the code required to render the very first screen. If they click a button to open an 'Admin Dashboard', only *then* does the app quickly fetch the code for that specific dashboard."

#### Indepth
In large React applications, bundle physics play heavily into perceived performance metrics like Time To Interactive (TTI). Webpack and Vite seamlessly support "code splitting," allowing different paths or heavily isolated components to exist as separate asynchronous `.js` files requested by the browser over the network on demand.

---

### 70. What is `React.lazy()`?
"`React.lazy()` is a built-in function that lets you dynamically import a component file only when it's first rendered. 

Instead of statically importing a component at the top of my file: `import AdminPage from './AdminPage'`, I use `const AdminPage = React.lazy(() => import('./AdminPage'))`.

React inherently knows that this code won't be available immediately, so it integrates flawlessly with Webpack's code-splitting to generate separate JS bundles for that page automatically."

#### Indepth
`React.lazy` only supports default exports right now. If the module you are importing exports via named exports (e.g., `export const AdminPage = ...`), you must manually re-export it as a default export in an intermediate file or creatively use `Promise.then(module => ({ default: module.AdminPage }))` to comply with `React.lazy`'s strict API requirements.

---

### 71. What is `Suspense`?
"`Suspense` is a component (`<Suspense fallback={<Spinner />}>`) that lets you 'wait' for some asynchronous work to finish, displaying a fallback UI (like a loading spinner) in the meantime.

It works perfectly in tandem with `React.lazy()`. When React encounters a lazy component that hasn't downloaded over the network yet, it 'suspends' the rendering of that subtree and jumps straight to the nearest `Suspense` fallback boundary above it. Once the network request finishes, React silently replaces the `<Spinner />` with the actual component."

#### Indepth
Originally built just for lazy-loading JavaScript chunks, `Suspense` is the foundational architecture for React 18's Concurrent Data Fetching model. With frameworks like Next.js or libraries like React Query, components can physically 'suspend' their rendering phase while waiting on API database calls, transforming declarative loading states at the router level without complex `if (isLoading)` boolean checks scattering the codebase.

---

### 72. What is React Router?
"**React Router** is the standard routing library for React applications. It enables the creation of Single-Page Applications (SPAs) with navigation.

In a traditional website, clicking a link fetches an entirely new HTML page from a server, which causes a blank white flash. 

React Router intercepts the link click, suppresses the server request, alters the browser's URL history directly via JavaScript, and then seamlessly swaps out the active React components on the screen to simulate navigating to a new page instantly."

#### Indepth
The core paradigm is declarative routing. You define your routes as actual React components `<Route path="/about" element={<AboutPage />} />` wrapped in a `<BrowserRouter>`. This means routing is fundamentally just an extension of the UI component tree itself, rather than a separate centralized configuration file parsed during initialization.

---

### 73. Difference between `BrowserRouter` and `HashRouter`?
"**`BrowserRouter`** uses the modern HTML5 History API (`pushState`, `replaceState`) to maintain clean, beautiful URLs (like `example.com/users/123`). This is the industry standard for production apps, but it requires backend server configuration to redirect all generic requests back to `index.html`.

**`HashRouter`** uses the hash portion of the URL (like `example.com/#/users/123`) to maintain state. Older browsers strictly interpret anything after a `#` as a client-side navigation anchor, so the server never sees the route changes. It works perfectly on simple static file-hosting sites like GitHub Pages without any complex server configurations needed."

#### Indepth
The "dirty urls" of `HashRouter` are problematic for modern SEO crawlers and analytics tools. Today, standard cloud hosting providers (Netlify, Vercel, AWS CloudFront) all have simple 1-click solutions for rewriting SPAs, meaning `BrowserRouter` should be heavily favored unless targeting hyper-legacy internal corporate network constraints.

---

### 74. What is `useNavigate`?
"`useNavigate` is a hook from React Router (v6) that gives you imperative control over navigation (how to force the user to go to a new page programmatically).

Instead of relying on a user physically clicking an `<a href>` link on screen, I get a function `const navigate = useNavigate()`. 

I use it heavily regarding logic outside the UI. E.g., inside an `onSubmit` handler, I wait for the backend 'Login Success' API response, and only then do I execute `navigate('/dashboard')` to push the user into the private area. It also allows going backward `navigate(-1)`."

#### Indepth
In prior versions of React Router (v5), developers had `useHistory`, and the act of navigating was slightly more verbose (`history.push('/dashboard')`). `useNavigate` is functionally simpler, supporting both relative routing concepts (navigating sibling folders cleanly without absolute slashes) and state passing directly inside options.

---

### 75. What is `useParams`?
"`useParams` is a hook from React Router that allows me to access the dynamic variables defined identically in my route's URL path.

If I define a Route as `<Route path="/users/:userId" element={<UserProfile />} />`, the `:userId` segment means anything typed there generates a variable.

Inside the `UserProfile` component, I simply call `const { userId } = useParams()`. If the browser URL is `myapp.com/users/99`, that hook returns the string `'99'`, which I then use to fire an API fetch query for user #99's data."

#### Indepth
All parameters extracted by `useParams` inherently return as strings. If you need numerical IDs for your API lookups, you must manually parse them with `parseInt()`. The data matching uses React Router's centralized path ranking algorithm, meaning specific hardcoded routes (`/users/new`) will correctly match before generalized parameter routes (`/users/:userId`), preventing collisions.
