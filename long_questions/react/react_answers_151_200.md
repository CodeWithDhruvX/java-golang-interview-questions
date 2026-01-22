## 🔹 16. Component Patterns

### Question 151: What is compound component pattern?

**Answer:**
A pattern where components work together to form a cohesive unit, usually sharing state implicitly. Think of `<select>` and `<option>` in HTML. In React, the parent component manages the state and shares it with children (often via Context or `React.Children.map`), allowing flexible UI composition.

**Example:**
```jsx
<Tabs>
  <TabList>
    <Tab>One</Tab>
    <Tab>Two</Tab>
  </TabList>
  <TabPanels>
    <Panel>Content 1</Panel>
    <Panel>Content 2</Panel>
  </TabPanels>
</Tabs>
```

---

### Question 152: What is render props pattern?

**Answer:**
A technique for sharing code between React components using a **prop whose value is a function**. The component calls this function with its internal state, allowing the parent to decide what to render based on that state.

**Example:**
```jsx
<MouseTracker render={({ x, y }) => (
  <h1>Mouse position: {x}, {y}</h1>
)} />
```

---

### Question 153: Difference between HOC and render props?

**Answer:**
*   **HOC (Higher Order Component):** A function that takes a component and returns a new component (Static composition). Can lead to "wrapper hell" and prop collision.
*   **Render Props:** A component with a function prop (Dynamic composition). More flexible, explicit data passing, but can lead to "callback hell" nesting in JSX.
*   *Note: Both are largely replaced by Custom Hooks for logic reuse.*

---

### Question 154: What is container-presentational pattern?

**Answer:**
A pattern separating concerns:
*   **Container Component:** Handles logic, fetching data, and state. Passes data down via props.
*   **Presentational Component:** Handles Look & Feel (UI). Accepts data via props and renders it.
*   *Status:* Less common with Hooks, as logic acts as the "container".

---

### Question 155: What is controlled vs uncontrolled pattern?

**Answer:**
Designing a component (like an Input or Modal) so it can work in two modes:
1.  **Controlled:** Parent manages value via props (`value`, `onChange`).
2.  **Uncontrolled:** Component manages its own internal state (`defaultValue`).
This makes components flexible for different consumer needs.

---

### Question 156: What is inversion of control in React?

**Answer:**
Instead of a component being responsible for everything (logic + UI), it hands over control of the *rendering part* to the user (the parent component) via **props** (children, render props, or slots).
*   "Don't call us, we'll call you" (rendering-wise).

---

### Question 157: What is slot pattern?

**Answer:**
Passing named components as props to fill specific "slots" in a layout, rather than dumping everything into `children`.

**Example:**
```jsx
<Layout
  header={<Header />}
  sidebar={<Sidebar />}
  content={<Feed />}
/>
```

---

### Question 158: What is provider pattern?**

**Answer:**
Using React Context's `<Provider>` to make data available to any component in the tree, regardless of depth. It solves prop drilling. Used by Redux (`<Provider store={store}>`) and React Router (`<Router>`).

---

### Question 159: What is polymorphic component?

**Answer:**
A component that can render as different HTML elements or React components, usually controlled via an `as` or `component` prop.

**Example:**
```jsx
<Text as="h1">Heading</Text>
<Text as="span">Span text</Text>
```

---

### Question 160: What is headless component?

**Answer:**
A component (often a Hook) that manages **logic, state, and accessibility** but provides **no UI markup**. The developer is fully responsible for the UI.
Examples: `react-table`, `downshift`.
"I give you the toggle logic, you build the switch button."

---

## 🔹 17. Code Splitting & Architecture

### Question 161: What is dynamic import in React?

**Answer:**
Standard JavaScript feature `import()` that loads a module asynchronously. It returns a Promise. When used with `React.lazy`, it allows splitting the code into separate chunks that are loaded only when rendered.

**Example:**
```jsx
const LazyComp = React.lazy(() => import('./LazyComp'));
```

---

### Question 162: How does webpack help React apps?

**Answer:**
Webpack is a **module bundler**.
1.  **Bundling:** Combines hundreds of JS files into one (or few) `bundle.js`.
2.  **Loaders:** Transforms non-JS files (Babel transforms JSX -> JS, Sass -> CSS).
3.  **HMR:** Hot Module Replacement (updates modules in browser without reload).
4.  **Optimization:** Minification, Tree Shaking.

---

### Question 163: What is tree shaking?

**Answer:**
A form of dead code elimination. During the build process, the bundler (Webpack/Rollup) statically analyzes the import/export graph and removes (shakes off) code that is **exported but never imported/used**. Requires ES6 Modules (`import/export`).

---

### Question 164: How does bundling affect performance?

**Answer:**
*   **Good:** Fewer HTTP requests (one request for `bundle.js` vs 100 requests for files).
*   **Bad:** If `bundle.js` is huge (e.g., 5MB), the browser takes a long time to download, parse, and execute it, delaying the First Interactive time.
*   **Solution:** Code Splitting (Chunks).

---

### Question 165: How do you organize components folder?

**Answer:**
Two common approaches:
1.  **Atomic Design:** atoms (Button), molecules (SearchBox), organisms (Header), templates (Page).
2.  **Feature-based (Recommended):** Group by feature (`/features/auth/Login.js`, `/features/auth/Login.css`, `/features/auth/authSlice.js`). Shared components go in `/components/ui`.

---

### Question 166: How do you manage reusable components?

**Answer:**
*   Keep them **Generic** (agnostic to business logic).
*   Use **Props** for configuration (Theme, Size, Content).
*   Document them (Storybook).
*   Prevent them from accessing global store/context directly (pass data in).

---

### Question 167: How do you manage environment variables?

**Answer:**
Store configuration in `.env` files which are not committed to Git (if secrets) or are built-in.
*   **CRA:** Prefix with `REACT_APP_` (e.g., `REACT_APP_API_URL`). Accessed via `process.env`.
*   **Vite:** Prefix with `VITE_` (e.g., `VITE_API_URL`). Accessed via `import.meta.env`.

---

### Question 168: How do you handle feature flags?

**Answer:**
Feature flags allow enabling/disabling features remotely without redeploying.
Usually managed via a Context Provider (`<FeatureFlagProvider>`) that fetches flags on load. Components check `useFeature('new-ui')` before rendering.

---

### Question 169: What is micro-frontend architecture?

**Answer:**
An architectural style where a frontend app is decomposed into individual, semi-independent "microapps" working loosely together. (e.g., Team A builds Search, Team B builds Checkout). They can be deployed independently.

---

### Question 170: How does React support micro-frontends?**

**Answer:**
React itself doesn't enforce it, but tools like **Webpack Module Federation** allow a React app to dynamically import components/pages exposed by *another* running React app at runtime.

---

## 🔹 18. Testing in React

### Question 171: What is unit testing in React?

**Answer:**
Testing the smallest units of code in isolation.
*   **Logic:** Testing utility functions `add(2,2)`.
*   **Components:** Testing a Button renders text correctly given a prop.
Doesn't call APIs or interact with other components.

---

### Question 172: What is integration testing?

**Answer:**
Testing how different units work together.
*   **Example:** Clicking a "Login" button (Component A) triggers a function that updates Context (Component B) and redirects the Router (Component C).
RTL (React Testing Library) excels here.

---

### Question 173: What is end-to-end testing?

**Answer:**
Testing the entire application flow from the user's perspective in a real browser environment.
*   **Tools:** Cypress, Playwright, Selenium.
*   **Scope:** "User visits site, logs in, adds item to cart, checks out."
*   Slowest but most confident.

---

### Question 174: What is Jest?

**Answer:**
Jest is a JavaScript **Test Runner** (maintained by Meta). It finds test files, runs them, provides assertions (`expect`), mocks functions, and generates coverage reports. It comes pre-installed with Create React App.

---

### Question 175: What is React Testing Library (RTL)?

**Answer:**
A set of helpers for testing React components.
**Philosophy:** "The more your tests resemble the way your software is used, the more confidence they can give you."
It queries the DOM by **accessibility** (Role, Label, Text) rather than implementation details (ClassName, ID, State).

---

### Question 176: Difference between Enzyme and RTL?

**Answer:**
*   **Enzyme:** Focused on **implementation details**. You could check `component.state('count')` or access private methods. Tests broke easily when you refactored code.
*   **RTL:** Focused on **behavior**. You check `screen.getByText('Count: 1')`. Refactoring code doesn't break tests as long as the UI stays the same.

---

### Question 177: How do you test hooks?

**Answer:**
Since Hooks can only run inside components, you cannot test them as plain functions.
Use **`renderHook`** from `@testing-library/react-hooks` (now built into main package).
```jsx
const { result } = renderHook(() => useCounter());
act(() => result.current.increment());
expect(result.current.count).toBe(1);
```

---

### Question 178: How do you mock API calls?

**Answer:**
1.  **Jest Mock:** `jest.spyOn(global, 'fetch')` and return fake promise.
2.  **MSW (Mock Service Worker):** Intercepts network requests at the network level and returns mocked responses. (Recommended for Integration tests).

---

### Question 179: How do you test async components?

**Answer:**
Since DOM updates happen asynchronously after promise resolution, you must await the UI change.
```jsx
// Wait for "Loading" to disappear and "Data" to appear
await waitFor(() => expect(screen.getByText('Data')).toBeInTheDocument());
// Or using findBy (async query)
const item = await screen.findByText('Data');
```

---

### Question 180: What is snapshot testing?

**Answer:**
A Jest feature that renders a component, takes a serialized string Dump of the DOM tree, and saves it to a file (`__snapshots__`).
On subsequent runs, it compares the new output to the saved file. If they differ, the test fails (indicating unexpected UI change).

---

## 🔹 19. Error Handling & Security

### Question 181: What are Error Boundaries?

**Answer:**
React components that catch JavaScript errors anywhere in their child component tree, log those errors, and display a fallback UI instead of the component tree that crashed.
They must be **Class Components**.

---

### Question 182: Why error boundaries don’t catch async errors?

**Answer:**
Error Boundaries essentially wrap the `render` method and lifecycle methods in a `try/catch`. 
Async events (timeouts, promises, click handlers) happen **outside** the render loop (on the event loop stack), so React's try/catch block has already exited by the time the error occurs.

---

### Question 183: How do you implement error boundaries?

**Answer:**
Define a class component implementing:
1.  `static getDerivedStateFromError(error)`: Update state to show fallback UI.
2.  `componentDidCatch(error, info)`: Log error to reporting service.

---

### Question 184: How do you handle API errors?

**Answer:**
Since Error Boundaries don't catch them, you handle them in the API call site.
```jsx
try {
  await fetchData();
} catch (error) {
  setErrorState(true); // Manually trigger error UI state
}
```

---

### Question 185: How do you handle 404 pages in React?

**Answer:**
In React Router, define a route with `path="*"` (wildcard) at the very bottom of your route list. If no other route matches, this one renders the "Not Found" component.

---

### Question 186: How do you prevent XSS in React?

**Answer:**
Cross-Site Scripting (XSS) is mostly prevented by React automatically escaping data binding `{userArgs}`.
**Developer Responsibility:**
1.  Avoid `dangerouslySetInnerHTML`.
2.  If you must use it, sanitize the HTML content (using libraries like `dompurify`).
3.  Validate `href` attributes (an attacker can put `javascript:alert(1)` in a link).

---

### Question 187: What is CSRF and how do you prevent it?

**Answer:**
**Cross-Site Request Forgery:** An attack forcing a user to execute unwanted actions on a web app where they are authenticated.
**Prevention:** React (frontend) cannot fully prevent it; the server must implement CSRF tokens. React simply needs to send this token (usually in headers) with requests. Using `SameSite=Strict` cookies also helps.

---

### Question 188: How do you secure React applications?

**Answer:**
1.  **No Secrets:** Never store private keys/secrets in client-side code (`.env` vars in React are embedded in build).
2.  **Auth:** Use secure, HTTP-only cookies for tokens.
3.  **Dependencies:** Audit `npm` packages (`npm audit`).
4.  **CSP:** Content Security Policy headers.

---

### Question 189: What is CORS?

**Answer:**
**Cross-Origin Resource Sharing.** A browser security mechanism that restricts HTTP requests initiated from scripts to a different origin (domain/port) than the one the script was loaded from.
**Fix:** The backend server must send `Access-Control-Allow-Origin` headers. React Dev Server Proxy (`setupProxy.js`) works for development.

---

### Question 190: How do you handle authentication in React?**

**Answer:**
1.  **Login:** Send credentials to API -> Receive Token (JWT).
2.  **Storage:** Store token (HttpOnly Cookie > LocalStorage).
3.  **State:** Use a Context (`AuthContext`) to store "IsAuthenticated" status and User Profile.
4.  **Protection:** Create a `ProtectedRoute` wrapper that checks auth state and redirects to Login if false.

---

## 🔹 20. Deployment & Best Practices

### Question 191: How do you build React for production?

**Answer:**
Run `npm run build` (or `yarn build`).
This uses Webpack (or Vite) to:
1.  Bundle files.
2.  Minify JS/CSS (remove whitespace/comments).
3.  Transpile to ES5 (browser compatibility).
4.  Generate Hashed filenames (cache busting).
Output is in `/build` folder.

---

### Question 192: What is environment-based configuration?

**Answer:**
Using different configuration values for Local, Staging, and Production.
React build tools inject these values at **build time**.
`REACT_APP_API_URL` might be `localhost:3000` in `.env.development` and `api.prod.com` in `.env.production`.

---

### Question 193: How do you improve first contentful paint?

**Answer:**
(FCP = Time until user sees 1st piece of content).
1.  **SSR:** Send HTML immediately.
2.  **Critical CSS:** Inline styles for above-the-fold content.
3.  **Font optimization:** Preload fonts.
4.  **Defer JS:** Don't block HTML parsing with massive JS bundles.

---

### Question 194: What is lighthouse score?

**Answer:**
An open-source tool by Google (in Chrome DevTools) that audits web pages for Performance, Accessibility, Best Practices, and SEO. It gives a score (0-100) and actionable suggestions.

---

### Question 195: How do you handle browser compatibility?

**Answer:**
1.  **Browserslist:** Config in `package.json` defining target browsers (`> 0.2%`, `last 2 versions`).
2.  **Polyfills:** Adding code to older browsers to support new features (e.g., `Promise`, `Array.flat`).
3.  **Transpilation:** Babel converts modern ES6+ syntax to widely supported ES5.

---

### Question 196: How do you handle memory leaks?

**Answer:**
React Warning: "Can't perform a React state update on an unmounted component."
**Fix:** In `useEffect`, always return a cleanup function.
*   `clearInterval(timer)`
*   `controller.abort()` (for fetch)
*   Remove event listeners.

---

### Question 197: How do you cancel API calls in React?

**Answer:**
Use the `AbortController` API.

**Example:**
```jsx
useEffect(() => {
  const controller = new AbortController();
  fetch(url, { signal: controller.signal });
  
  return () => controller.abort(); // Cancel on unmount
}, [url]);
```

---

### Question 198: What is cleanup in async effects?

**Answer:**
See Q196/Q197. It is ensuring that any ongoing side process (timer, subscription, network request) is stopped when the component unmounts or dependency changes, preventing effects from stacking up or trying to update unmounted UI.

---

### Question 199: How do you log errors in production?

**Answer:**
You cannot see the console in a user's browser. Use a logging service (Sentry, LogRocket).
Hook into `componentDidCatch` (Error Boundary) and send the error stack trace to the service.

---

### Question 200: What are React best practices?

**Answer:**
1.  **Functional Components & Hooks:** Standard since 2019.
2.  **Folder Structure:** Feature-based over type-based.
3.  **Small Components:** Single Responsibility Principle.
4.  **Prop Types / TypeScript:** Enforce data contracts.
5.  **Simplicity:** Don't reach for Redux/Context immediately; use local state until needed.
6.  **Performance:** Monitor re-renders; use Memoization only when profiled.
