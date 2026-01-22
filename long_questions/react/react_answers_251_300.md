## 🔹 26. Accessibility (a11y)

### Question 251: Why is accessibility important in React apps?

**Answer:**
1.  **Inclusivity:** Ensures users with disabilities (vision, motor, cognitive) can use your app.
2.  **Legal:** Compliance with laws like ADA (US) or EAA (EU).
3.  **SEO:** Accessible sites (Semantic HTML) are better indexed by search engines.

---

### Question 252: How do you make React apps accessible?

**Answer:**
React fully supports building accessible websites.
1.  Use **Semantic HTML** (`<button>` vs `div`).
2.  Add **ARIA attributes** (`aria-label`, `aria-expanded`).
3.  Manage **Focus** (Ref focus management).
4.  Ensure **Keyboard Navigation** works.

---

### Question 253: What are ARIA roles?

**Answer:**
**Accessible Rich Internet Applications (ARIA)** roles defines the *purpose* of an element to assistive technologies (like Screen Readers).
*   `role="alert"`: Important message.
*   `role="dialog"`: Popup/Modal.
*   `role="navigation"`: Navbar.

---

### Question 254: How do you manage focus for modals?

**Answer:**
When a modal opens:
1.  Move focus *into* the modal (on the first input or close button).
2.  **Trap** focus inside the modal (Tab key should cycle inside).
3.  When modal closes, **Restore** focus to the element that opened it.

---

### Question 255: How do you make forms accessible?

**Answer:**
1.  Every input must have a label. Use `<label htmlFor="id">`.
2.  If visible label isn't possible, use `aria-label`.
3.  Group related inputs (radio buttons) with `<fieldset>` and `<legend>`.
4.  Link error messages to inputs using `aria-describedby="error-id"`.

---

### Question 256: How do you support keyboard navigation?

**Answer:**
1.  Use native interactive elements (`<button>`, `<a>`) which get keyboard support for free.
2.  If building custom widgets (`div` acting as button), add `tabIndex="0"` to make it focusable.
3.  Add `onKeyDown` listeners for `Enter` and `Space` keys to trigger actions.

---

### Question 257: How do you handle screen readers?

**Answer:**
Test using tools like **VoiceOver** (Mac) or **NVDA** (Windows).
Ensure images have `alt` text. Use `aria-hidden="true"` for decorative icons to skip them. Verify reading order matches visual order.

---

### Question 258: What is semantic HTML and why is it important?

**Answer:**
Using the correct HTML tag for the content.
*   `<header>`, `<main>`, `<footer>` instead of `<div>`.
*   `<button>` for actions, `<a>` for links.
**Importance:** Browsers and Screen Readers understand built-in behaviors (keyboard shortcuts, specific announcements) of semantic tags automatically.

---

### Question 259: How do you test accessibility?

**Answer:**
1.  **Automated:** Lighthouse Audit in Chrome.
2.  **Linting:** `eslint-plugin-jsx-a11y` (catches missing alt text, etc.).
3.  **Manual:** Navigate your site using *only* the keyboard (Tab/Shift+Tab).

---

### Question 260: What tools are used for accessibility testing?

**Answer:**
*   **Axe DevTools:** Browser extension.
*   **WAVE:** Web Accessibility Evaluation Tool.
*   **Pa11y:** CI/CD tool.
*   **Screen Readers:** JAWS, NVDA, VoiceOver.

---

## 🔹 27. Integration & Ecosystem

### Question 261: How do you integrate REST APIs in React?

**Answer:**
Use the browser's `fetch()` API or libraries like `axios` inside the `useEffect` hook.

**Example:**
```jsx
useEffect(() => {
  axios.get('/api/users').then(res => setUsers(res.data));
}, []);
```

---

### Question 262: How do you integrate GraphQL with React?

**Answer:**
Use libraries like **Apollo Client** or **Urql**.
Component sends a query string (GQL) to the endpoint (not REST). The library handles caching and loading states.
`const { loading, error, data } = useQuery(GET_USERS);`

---

### Question 263: What is Apollo Client?

**Answer:**
The most popular GraphQL client for React. It provides:
1.  **Data Fetching:** Hooks (`useQuery`, `useMutation`).
2.  **Caching:** Normalizes data to avoid re-fetching.
3.  **State Management:** Can replace Redux for server data.

---

### Question 264: What is React Query / TanStack Query?

**Answer:**
A library for fetching, caching, synchronizing, and updating **Server State**.
It handles the hard parts of server data: Caching, Deduping requests, Background updates, and Stale data.

---

### Question 265: Difference between Redux and React Query?

**Answer:**
*   **Redux:** Client State Manager. You must manually write thunks/reducers to fetch/save data.
*   **React Query:** Server State Manager. You pass a Promise; it handles the loading/error/data states and caching automatically.
*   *Trend:* Use React Query for API data + Context/Zustand for UI state.

---

### Question 266: How do you cache server data?

**Answer:**
Libraries like React Query or SWR create a client-side cache key (e.g., `['user', 1]`). If you request this data again, it serves from cache immediately while fetching fresh data in the background (Stale-While-Revalidate).

---

### Question 267: How do you handle pagination?

**Answer:**
1.  Store `page` number in state.
2.  Pass `page` to the API query.
3.  When user clicks "Next", update `page` state.
4.  React Query handles pre-fetching the next page for smooth UX.

---

### Question 268: How do you handle infinite queries?

**Answer:**
Use `useInfiniteQuery` (React Query).
It handles fetching "pages" of data and appending them to the previous data arrays.
UI renders a "Load More" button or uses Intersection Observer for infinite scroll.

---

### Question 269: How do you sync server state and UI state?

**Answer:**
The UI should reflect the Server State.
When the user mutates data (e.g., "Change Name"), successful response updates the Server.
Then, you **Invalidate** the cache in the React app to trigger a refetch and update the UI.

---

### Question 270: How do you handle real-time updates?

**Answer:**
1.  **WebSockets:** Open connection, listen for events, update state store.
2.  **Server-Sent Events (SSE):** Simpler unidirectional stream.
3.  **Polling:** Fetch data every X seconds (easiest, but resource heavy).

---

## 🔹 28. Build Tools & Configuration

### Question 271: What is Vite and how is it different from CRA?

**Answer:**
*   **CRA (Create React App):** Uses **Webpack**. Bundles the *entire* app before starting the dev server. Slow on large apps.
*   **Vite:** Uses **ES Modules** (native browser support). Serves files on demand. Instantly starts dev server. Uses **Rollup** for production build.

---

### Question 272: What is Create React App (CRA)?

**Answer:**
An official CLI tool to set up React environment with zero configuration. It hides the complex Webpack/Babel config.
**Status:** Deprecated/Legacy. React team now recommends frameworks (Next.js) or Vite.

---

### Question 273: Why is CRA being deprecated?

**Answer:**
1.  **Performance:** Webpack is slower than tools like Vite/Turbopack.
2.  **Features:** Doesn't support Server Components, SSR, or modern React 18 features out of the box.
3.  **Maintenance:** Hard to customize without "ejecting".

---

### Question 274: What is Babel configuration used for?

**Answer:**
Babel is a **Transpiler**. It converts modern JavaScript (ES6+, JSX) into backward-compatible JavaScript (ES5) that older browsers can understand.
Configuration is in `.babelrc`. Presets like `@babel/preset-react` tell it how to handle JSX.

---

### Question 275: What is ESLint and why is it used?

**Answer:**
ESLint is a **Linter**. It statically analyzes code to find problems.
*   **Errors:** Syntax issues, using undefined variables.
*   **Style:** Enforcing semicolon usage, indentation.
*   **React Rules:** Hooks rules (dependency array checks).

---

### Question 276: What is Prettier?

**Answer:**
An **opinionated Code Formatter**. It rewrites your code to ensure consistent style (spacing, quotes, wrapping).
Difference: ESLint checks *code quality* (bugs); Prettier checks *formatting* (style).

---

### Question 277: How do you enforce code quality?

**Answer:**
Use **Husky** and **lint-staged**.
These are Git Hooks. Before a commit (`pre-commit`), Husky runs ESLint and Prettier on the changed files. If there are errors, the commit is blocked.

---

### Question 278: How do you manage monorepos?

**Answer:**
A Monorepo (one git repository) containing multiple packages/apps.
Tools: **Yarn Workspaces**, **Nx**, **Turborepo**, **Lerna**.
Allows sharing code (`/libs/ui`) easily between apps (`/apps/web`, `/apps/admin`).

---

### Question 279: What is Nx / Turborepo?

**Answer:**
High-performance build systems for monorepos.
They use **Computation Caching**. If you build Library A, and it hasn't changed, the tool restores the build result from cache instantly instead of rebuilding.

---

### Question 280: How do you manage multiple environments?

**Answer:**
1.  **.env files:** `.env.development`, `.env.staging`, `.env.production`.
2.  **Runtime Config:** For Docker images, inject env vars at container startup into a `window.__ENV__` object so the same build artifact works everywhere.

---

## 🔹 29. Advanced Patterns & Real-World Scenarios

### Question 281: How do you implement role-based access control?

**Answer:**
1.  Store User Roles in Auth Context.
2.  Create an `<Authorization roles={['admin']}>` wrapper.
3.  If `user.role` matches, render children. If not, render Access Denied or null.

---

### Question 282: How do you protect private routes?

**Answer:**
Create a wrapper component `<PrivateRoute>`.
```jsx
const PrivateRoute = () => {
  const { user } = useAuth();
  return user ? <Outlet /> : <Navigate to="/login" />;
};
```
Wrap protected routes inside this Route.

---

### Question 283: How do you handle feature-based rendering?

**Answer:**
Feature Flags (Toggles).
```jsx
{flags.newModule ? <NewModule /> : <OldModule />}
```
Fetch flags from server on app load.

---

### Question 284: How do you handle large tables efficiently?

**Answer:**
Rendering 10,000 DOM rows crashes the browser.
Use **Virtualization** (e.g., `react-window`). It only renders the 20 rows visible on the screen. As you scroll, it swaps the data in those 20 rows.

---

### Question 285: How do you implement virtualized lists?

**Answer:**
You need a container with a fixed height.
The library calculates the total scroll height (ItemHeight * Count).
It positions the visible items absolutely based on `scrollTop`.

---

### Question 286: What is windowing?

**Answer:**
Synonym for **Virtualization**. "Windowing" refers to the concept of only looking at a small "window" of the data set at any given time.

---

### Question 287: How do you handle file downloads?

**Answer:**
1.  **Static:** `<a href="/file.pdf" download>`.
2.  **Dynamic:** Fetch Blob from API. Create URL `URL.createObjectURL(blob)`. Create temp `<a>`, click it, revoke URL.

---

### Question 288: How do you handle file previews?

**Answer:**
When user selects a file (`input type="file"`):
```jsx
const url = URL.createObjectURL(file);
setImageSrc(url);
```
Show this URL in an `<img>` tag immediately.

---

### Question 289: How do you handle multi-language support (i18n)?

**Answer:**
Use **react-i18next** or **react-intl**.
1.  Create JSON files for each language (`en.json`, `fr.json`).
2.  Wrap app in Provider.
3.  Use hook: `const { t } = useTranslation(); <h1>{t('title')}</h1>`.

---

### Question 290: How do you manage translations dynamically?

**Answer:**
Don't bundle all languages. Use **Lazy Loading**.
When user switches language, fetch the `fr.json` chunk from the server, verify it loaded, and then update the i18n instance language.

---

## 🔹 30. Debugging & Maintenance

### Question 291: How do you debug React applications?

**Answer:**
1.  **React DevTools:** Inspect Component tree, props, state, hooks.
2.  **Profiler:** Record user session to see which components rendered and how long they took.
3.  **Console:** `console.log` works, but `debugger` statement is better.
4.  **Network Tab:** Check API payloads.

---

### Question 292: What are common React anti-patterns?

**Answer:**
1.  **Index as Key:** Bugs in list ordering.
2.  **Props Drilling:** Hard to maintain.
3.  **Nested Component Definitions:** Defining a component *inside* another component (remounts on every render).
4.  **Mutating State:** `state.value = 5`.

---

### Question 293: How do you find memory leaks?

**Answer:**
1.  **Console Warnings:** React warns about "updates on unmounted component".
2.  **Performance Monitor:** In Chrome DevTools, check JS Heap size increasing over time without dropping (GC).
3.  **Heap Snapshot:** Compare snapshots to see objects (like Detached DOM nodes) retained in memory.

---

### Question 294: How do you use React DevTools?

**Answer:**
Install browser extension.
*   **Components Tab:** View tree, select component to see Props/State/Hooks.
*   **Profiler Tab:** Click Record -> Interact -> Stop. View flamegraph of renders. "Why did this render?"

---

### Question 295: How do you monitor performance?

**Answer:**
1.  **Core Web Vitals:** LCP (Loading), FID (Interactivity), CLS (Stability).
2.  **React Profiler:** Component-level render times.
3.  **Lighthouse:** General performance audit.

---

### Question 296: How do you log user interactions?

**Answer:**
Use a middleware (Redux) or a custom hook/utility.
Send events ("Button Clicked", "Page Viewed") to analytics services like Google Analytics, Mixpanel, or Amplitude.

---

### Question 297: How do you handle version upgrades?

**Answer:**
1.  Read the **Changelog**.
2.  Use **Codemods**: Scripts provided by the React team (e.g., `npx react-codemod rename-unsafe-lifecycles`) to automatically rewrite legacy code.
3.  Run tests.

---

### Question 298: How do you refactor legacy React code?

**Answer:**
**Strangler Fig Pattern:**
Don't rewrite everything at once.
Refactor one component at a time. Convert Class -> Function. Introduce Hooks. Replace localized state with Context where needed.

---

### Question 299: How do you migrate class components to hooks?

**Answer:**
1.  Change `class extends` to `function`.
2.  `this.props` -> `props`.
3.  `this.state` -> `useState`.
4.  `componentDidMount` -> `useEffect`.
5.  Remove `render()`, just `return`.

---

### Question 300: What makes a React application scalable?

**Answer:**
1.  **Architecture:** Solid folder structure (Feature-based).
2.  **Types:** TypeScript interfaces.
3.  **Testing:** High coverage (Unit + E2E).
4.  **Component Library:** Reusable, isolated UI components.
5.  **State Management:** Clear separation of Server vs Client state.
