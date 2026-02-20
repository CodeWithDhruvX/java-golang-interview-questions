# 🟣 **Senior (6+ Years): Ecosystem, Integrations & Architecture - Part 2**

### 111. What is React Query / TanStack Query?
"**React Query (now TanStack Query)** is a powerful data-fetching and server-state management library for React.

While Redux is excellent for *client* state (like a dark mode toggle or an open modal string), it is terrible at managing *server* state (data fetched from an API). Server state is inherently asynchronous, can be changed by other users remotely, and requires complex caching logic.

React Query completely replaces the need to store API data in Redux or `useState`. It provides a `useQuery` hook that handles fetching, caching, synchronizing, background updating, and garbage collecting server data with literally zero configuration required."

#### Indepth
React Query structurally alters how developers build apps. It introduces the concept of **Cache Keys** (an array like `['todos', userId]`). Any component anywhere in the app that calls a `useQuery` with that exact same key instantly receives the cached data. If the data is deemed "stale" (configurable, default is instantly), React Query silently triggers a background fetch to the server, and automatically surgically re-renders *only* the components watching that key once the fresh data arrives.

---

### 112. Difference between Redux and React Query?
"They solve fundamentally different problems:

**Redux manages Client State:** It handles synchronous data that exists purely on the user's device. Things like 'Is the sidebar open?', 'What filters are currently selected on this table?', or 'What steps has the user completed in this multi-page wizard?'

**React Query manages Server State:** It handles asynchronous data whose ultimate source of truth lives on a remote database. It manages the agonizing complexity of knowing *when* to fetch it, *where* to store it temporarily, and *how* to know when it has become outdated compared to the server.

In a modern enterprise architecture, I use React Query for 90% of my data (API calls), and Zustand (a lighter Redux alternative) for the remaining 10% (UI toggles)."

#### Indepth
Historically, developers shoved all server data into Redux because no better tool existed. This led to "Cache invalidation by Reducer," requiring thousands of lines of boilerplate to dispatch `FETCH_START`, `FETCH_SUCCESS`, and `FETCH_ERROR` actions. Ripping server data out of Redux and migrating to React Query usually deletes upwards of 40% of an application's codebase while making the app dramatically faster due to intelligent deduplication.

---

### 113. How do you integrate GraphQL with React?
"Integrating GraphQL is significantly different from REST because GraphQL only exposes a single endpoint (usually `/graphql`) where you post incredibly specific query strings.

While I *can* use the native `fetch()` API and manually construct the GraphQL string, that is tedious and error-prone. 

Instead, I use dedicated GraphQL clients, primarily **Apollo Client** or **URQL**. They provide dedicated hooks like `const { loading, error, data } = useQuery(GET_USER_QUERY)`. They operate almost identically to React Query, offering normalized caching out of the box so that if I fetch a `User` object on one page, and fetch a mutated version of that same `User` object elsewhere, Apollo intelligently merges them in its internal cache."

#### Indepth
A major advantage of GraphQL in a React environment is the ability to colocate data requirements. Utilizing fragments (`...UserAvatarFragment`), a tiny `<Avatar>` component can physically define exactly what database fields it needs (like `avatarUrl`), and the parent `<Dashboard>` component aggregates five different child fragments into one massive GraphQL network request, preventing the classic "Waterfall problem" of REST APIs firing sequentially.

---

### 114. What is Vite and how is it different from CRA?
"**Vite** is a modern, extraordinarily fast build tool created by Evan You (creator of Vue), and it has completely replaced Create React App (CRA) as the industry standard for starting new Single Page Applications.

**Create React App (Webpack):** Bundles your entire application into a massive file *before* serving it. If you have 5,000 components, spinning up the local dev server can literally take 30-60 seconds.

**Vite (esbuild + Rollup):** Leverages native ES modules built directly into modern browsers. It does not bundle your code during development. It pre-bundles your specific `node_modules` instantly using `esbuild` (written in Go, 10-100x faster than Webpack), and then serves your source files directly to the browser on demand. The dev server starts in ~200 milliseconds regardless of app size."

#### Indepth
The React Core team has officially marked Create React App as completely deprecated. If you visit the official React documentation today, they instruct you to use a meta-framework like Next.js or Remix for production systems, or Vite if you strictly need a classic isolated client-side SPA. Webpack remains highly relevant for complex legacy enterprise monoliths requiring intricate custom loader logic, but Vite dominates all new web development.

---

### 115. How do you prevent XSS (Cross-Site Scripting) in React?
"React has an incredibly powerful built-in defense against XSS out of the box: **Auto-Escaping**.

By default, any data you render using curly braces `{userData}` is physically treated as a string, not executable HTML. If a malicious user sets their username to `<script>stealCookies()</script>`, React strictly renders it harmlessly as literal text on the screen.

However, XSS vulnerabilities still occur in React if developers manually bypass this protection using the `dangerouslySetInnerHTML={{ __html: dirtyData }}` prop, or if they dynamically evaluate user input inside `href` attributes (like `<a href={userInput}>`, where `userInput` could be `javascript:maliciousCode()`)."

#### Indepth
When I absolutely must render raw HTML from a CMS or rich-text editor, I **never** pass it directly to `dangerouslySetInnerHTML`. I always pass the raw string through a dedicated sanitization library first, primarily **DOMPurify** (`DOMPurify.sanitize(dirtyHtml)`). This strips out all executable script tags and `<img onerror="...">` vector attacks while preserving safe text formatting.

---

### 116. How do you secure React applications?
"Frontend security is fundamentally an oxymoron: the client's browser is inherently a zero-trust environment. Anything sent to the React app can be read, and any 'secure' route can be manually bypassed via DevTools.

However, to secure the *system* from the frontend:
1. **Never store secrets:** API keys or database passwords must never exist in React code.
2. **JWT Storage:** I explicitly avoid storing sensitive JWT access tokens in `localStorage` (where they are susceptible to XSS). I configure the backend API to issue `HttpOnly` Secure cookies instead, tightly coupling the session to the browser itself.
3. **CSRF Protection:** If using cookies, I implement Anti-CSRF tokens to prevent other malicious websites from forcing the user's browser to make invisible requests to my API.
4. **Dependency Scanning:** I strictly run `npm audit` in CI/CD pipelines to catch vulnerabilities in the thousands of third-party packages installed."

#### Indepth
Authentication routing in React (e.g., `<ProtectedRoute>`) is purely for User Experience—preventing the user from seeing a blank screen while a failed API call bounces them. It is *not* a security boundary. The ultimate security of a React application rests 100% on the Backend API rigorously validating the authorization headers of every single incoming request, regardless of what the React UI displays.

---

### 117. How do you implement infinite scrolling?
"Infinite scrolling in React is usually achieved by merging an asynchronous data fetching hook with the browser's native **Intersection Observer API**.

1. I render a tiny, invisible `<div ref={bottomRef}>` at the absolute bottom of my list.
2. I attach an Intersection Observer to that specific ref. It listens for the exact millisecond that div becomes visible on the user's screen (meaning they scrolled to the bottom).
3. When the observer fires, I trigger my data fetching function, asking for 'Page 2'.
4. I append the new Page 2 array results to my existing Page 1 state array, causing React to seamlessly render the new items downward."

#### Indepth
React Query solves the agonizing state management of this via its `useInfiniteQuery` hook, which provides a simple `fetchNextPage()` function and `hasNextPage` boolean. However, truly massive infinite lists (like Twitter's feed) suffer from severe DOM bloat. Appending 5,000 DOM nodes will eventually crash mobile browsers. I solve this by integrating **Virtualization** (windowing) libraries like `@tanstack/react-virtual`, which aggressively unmounts elements scrolling off the top of the screen and recycles their DOM nodes for the new items scrolling up from the bottom.

---

### 118. How do you organize a large React application folder structure?
"For enterprise applications, I strictly avoid organizing files purely by their technical type (e.g., putting all components in `src/components`, all hooks in `src/hooks`, all Redux actions in `src/actions`). In a huge app, making one new feature requires opening 15 different folders.

Instead, I use a **Feature-Based (or Domain-Driven) Architecture**.

My `src/features` folder contains independent verticals: `features/auth`, `features/checkout`, `features/dashboard`. Inside `features/auth`, I colocate *everything* related solely to authentication: its specific components, its API hooks, its types, and its tests.

Only truly generic, globally reusable UI pieces (like a primary `<Button>`) go into a top-level `src/components` folder."

#### Indepth
A critical architectural pattern for scalability is enforcing **Barrel Files** (`index.ts`). A feature folder should only export the absolute minimum components required for the rest of the application to interact with it (like `export { LoginForm } from './components'`). This encapsulation prevents the rest of the app from deeply importing highly specific internal subcomponents, making massive codebase refactors significantly safer.

---

### 119. How do you structure a large-scale React application?
"Beyond folder structure, scaling a React app requires establishing rigorous boundaries and consistent patterns:

1. **State Locality:** Push state as physically low in the component tree as possible. Never use Redux/Global Context for something only two sibling components care about.
2. **Layer Separation:** Never fetch `axios` directly inside a UI button. The UI calls a Custom Hook. The Custom Hook calls an API service file. The Service file makes the Axios call.
3. **Design System:** Implement a headless or unified component library early. Do not allow developers to write raw margin/padding CSS arbitrarily.
4. **Strict Typing:** Enforce TypeScript rigorously. Refuse the `any` type.
5. **Code Splitting:** Route-based lazy loading is mandatory. Users visiting `/login` should not download the javascript bundle required for `/admin-dashboard`."

#### Indepth
The ultimate structural evolution for massive organizations (like Netflix or Spotify) is implementing **Micro-Frontends**. This involves splitting a monolithic React application into dozens of completely separate React applications that physically stitch themselves together in the user's browser via Webpack Module Federation. This allows 50 different developer squads to deploy independent parts of the page simultaneously without blocking the monolithic centralized release train.

---

### 120. How do you optimize large React applications?
"React optimization is always metrics-driven. I don't guess; I strictly use the React Profiler DevTool and Chrome Lighthouse to identify bottlenecks.

If the *Initial Load Time* is slow:
- Implement `React.lazy()` route splitting.
- Analyze Webpack/Vite bundle sizes (e.g., identifying large imports like `moment.js` and swapping to `date-fns`).
- Implement Server-Side rendering (Next.js) for instantaneous Time-To-Interactive.

If the *Runtime Rendering* is slow (animations lag, typing stutters):
- Identify cascading re-renders in the Profiler.
- Fix broken object references in prop drilling using `useMemo` and `useCallback`.
- Move rapidly animating values (like complex mouse-tracking dragging math) physically outside of React state entirely, utilizing `useRef` directly on DOM nodes to bypass the render queue loop."

#### Indepth
The most profound optimization often involves deleting code. I consistently search for massive third-party libraries that can be replaced by native browser APIs (e.g., replacing heavy intersection-observer polyfills, or using native CSS `scroll-behavior: smooth` instead of complex React animation JS libraries). The fastest JavaScript in the world is the JavaScript you don't send to the client.
