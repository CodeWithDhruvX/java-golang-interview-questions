# Performance & Error Handling in Microfrontends

A poorly implemented Microfrontend architecture can severely degrade performance metrics (Core Web Vitals) and result in a chaotic user experience when things fail.

## 1. You have a Host App and 5 Remote Microfrontends loading on a single dashboard page. The Largest Contentful Paint (LCP) is terrible. How do you fix it?

**Answer:**
**The Problem:**
Loading 5 Remotes simultaneously means 5 round trips to DNS, 5 TCP handshakes, and 5 separate `remoteEntry.js` files to parse, before the browser even realizes what CSS or JS chunks it actually needs to render the components. This completely blocks the main thread and delays LCP.

**Mitigation Strategies:**
1.  **Lazy Loading below the fold:**
    *   Do not mount or request Microfrontends that the user cannot immediately see (e.g., a "Related Products" MFE at the bottom of the page).
    *   Use React's `lazy` and `Suspense` wrapped with an `IntersectionObserver` to trigger the fetch only when the component scrolls into view.
2.  **Resource Hints (Prefetching / Preconnecting):**
    *   Add `<link rel="preconnect" href="https://cdn.mfe-bucket.com">` in the Host's `index.html`. This performs the DNS and TCP/TLS handshakes early while the browser is still parsing the Host's head tag.
    *   Use Webpack's magic comments to prefetch assets: `import(/* webpackPrefetch: true */ 'remoteApp/Widget')`. The browser will fetch the Remote code intelligently during idle time.
3.  **Strict Shared dependencies (Preventing library bloat):**
    *   Ensure all MFEs strictly share heavyweight libraries (React, ReactDOM, Lodash, moment.js) via the `singleton: true` Webpack Module Federation config. If 5 MFEs download 5 different versions of React, the JS parsing time will destroy the LCP.
4.  **Skeleton Screens (Optimizing CLS and perceived performance):**
    *   While the Remote MFE is asynchronously fetching its `remoteEntry.js` and rendering chunks over the network, the Host must immediately render an exact-sized Skeleton/Placeholder in the DOM. This prevents Cumulative Layout Shift (CLS) when the MFE finally mounts and pushes other UI elements down.

## 2. A crucial Microfrontend (e.g., the "Add to Cart" button) fails to load over the network, or it throws an unhandled Javascript exception upon mounting. How do you prevent the entire Host application from crashing?

**Answer:**
You must implement multiple layers of fault tolerance and graceful degradation.

**The Strategy:**
1.  **React Error Boundaries (Client-side exception handling):**
    *   When dynamically importing a Remote MFE component, always wrap it in a React Error Boundary (`<ErrorBoundary>`).
    *   *How it works:* If the Remote MFE throws a `TypeError: Cannot read properties of null (reading 'price')` during its `render` phase or inside `useEffect`, the Error Boundary catches the error.
    *   *The Result:* Instead of the entire page turning blank (a React 16+ protective mechanism), the Error Boundary catches it and renders a fallback UI specific to that component section (e.g., a "Service Temporarily Unavailable" gray box), while the rest of the application (Header, Footer, Navigation) remains perfectly functional.
2.  **Network Failure Handling (Dynamic Import Fallbacks):**
    *   If the CDN is down and the browser cannot fetch `remoteEntry.js`, the dynamic `import('remoteApp/CartButton')` promise will reject.
    *   You must `.catch()` this promise rejection (often handled transparently by React's `<Suspense>` combined with `lazy()`).
    *   *Fallback Strategy:* Render a "Retry" button, or an inactive, grayed-out version of the "Add to Cart" button that, when clicked, attempts to fetch the Remote MFE again.
3.  **Circuit Breaker Pattern (for complex state):**
    *   If an MFE repeatedly fails to load or consistently throws errors (e.g., an e-commerce "Recommendations" widget is timing out the main thread), the Host application's code should "trip" a circuit breaker. Stop attempting to render that specific MFE for the duration of the user's session to preserve battery and CPU for critical flows (like Checkout).

## 3. How do you test a Remote MFE in isolation if it heavily relies on data or context provided by the Host application? (e.g., User Authentication Token)

**Answer:**
**The Problem:**
A Remote MFE often expects the Host app to have already validated the user, set the layout boundaries, and perhaps injected a global "Theme" or "User" React Context. If a developer runs `npm start` in the Remote repository, they just get a blank screen or a crash string like `Cannot destructure property 'userId' of 'useContext(UserContext)' as it is undefined`.

**The Solution:**
You must build a "Test Shell" or "Dev Shell" within the Remote MFE's repository.

**How to implement the Dev Shell:**
1.  **Mock Host Shell Application:** Create a lightweight `/dev-app` folder inside the Remote repository.
2.  **Mocked Context Providers:** This Dev app acts exactly like the production Host. It wraps the Remote components in a mocked `<UserContext.Provider value={{ id: 999, name: "Test User", role: "admin" }}>` and a `<ThemeProvider theme="dark">`.
3.  **Entry Point Swap:** When running `npm run start` (development mode) in the remote repo, Webpack serves `index.html` pointing to the Dev App. When running `npm run build` (production mode), Webpack only compiles and outputs the remote exposed components (e.g., `remoteEntry.js`), ignoring the Dev app entirely.
4.  **Use Cases:** This Dev Shell allows developers to work perfectly locally (and run Cypress E2E tests against the remote) without needing to clone, install, and run the massive global Host Application repo.
