# Routing & Lifecycle in Microfrontends

Connecting different applications intuitively requires a strong grasp of browser routing and component lifecycle management. These questions explore how to make isolated applications feel like a single cohesive product.

## 1. How does routing work when you have multiple independent applications (Microfrontends) on the same page?

**Answer:**
Routing in an MFE architecture is typically divided into two distinct layers:

1.  **Global Routing (The Shell/Host Level):**
    *   The Host application is responsible for top-level routing based on the URL path prefix.
    *   It uses a router (like a custom router in Single-SPA or the top-level React Router) to determine which Remote application to load and mount into the DOM.
    *   *Example:* User visits `example.com/products`. The Global router sees `/products`, downloads the `CatalogApp` JavaScript, and mounts it.
2.  **Local Routing (The Remote/Micro-app Level):**
    *   Once the Remote application (`CatalogApp`) is mounted, it takes over routing for its specific domain.
    *   The Remote has its own internal router (e.g., React Router) to handle nested paths.
    *   *Example:* User clicks a product, URL changes to `example.com/products/tshirt`. The Remote router sees `/tshirt` and renders the specific Product Details component *without* the Host Router interfering or reloading the whole page.

**The Golden Rule:** Child routers must be aware of their base path. If `CatalogApp` is mounted at `/products`, its internal router must treat `/products` as its root (`/`), otherwise nested routes will fail to match.

## 2. In a framework like Single-SPA, what are the three main lifecycle methods a Microfrontend must implement, and what do they do?

**Answer:**
For an application to be managed by an orchestrator like Single-SPA, it must explicitly define how it is created, shown, and destroyed. It exports three asynchronous Javascript functions:

1.  **`bootstrap(props)`:**
    *   *When it runs:* Called exactly once, the very first time the MFE is requested by the Host.
    *   *What it does:* Performs one-time initialization tasks before the application is shown to the user. This might involve setting up a Redux store, initializing analytics, or pre-fetching critical initial data. It does *not* touch the DOM.
2.  **`mount(props)`:**
    *   *When it runs:* Called whenever the Global Router decides this MFE should be active (e.g., the URL matches its route).
    *   *What it does:* This is where the framework (React/Angular) physically attaches the application's components to the DOM (e.g., `ReactDOM.render()`). It starts event listeners and triggers initial component renders.
3.  **`unmount(props)`:**
    *   *When it runs:* Called when the Global Router decides to navigate away from this MFE (e.g., the user clicks a link to a different app).
    *   *What it does:* **Crucial for performance.** It must physically remove the application from the DOM (e.g., `ReactDOM.unmountComponentAtNode()`). It must also clean up any global event listeners, intervals, or open WebSocket connections to prevent massive memory leaks.

## 3. How do you prevent CSS styles written by Team A for their MFE from accidentally breaking the layout of Team B's MFE? Name two common strategies.

**Answer:**
CSS is globally scoped by default in the browser, meaning a `.button { display: none; }` rule in MFE A will hide all buttons in MFE B if they are on the screen at the same time.

Two common ways to solve this "CSS Bleeding":

1.  **CSS Modules (or CSS-in-JS like Styled Components):**
    *   *How:* Tools like Webpack or styled-components dynamically generate unique hash classes during the build process.
    *   *Effect:* A developer writes `.button { color: red; }`, but Webpack outputs `.button_X1yZ { color: red; }`. This guarantees the class name is geographically isolated to that specific component in that specific MFE.
2.  **Prefixing / BEM Naming Conventions:**
    *   *How:* Strict organizational rules mandate that every CSS class must be prefixed with the MFE's name.
    *   *Effect:* Team A writes `.cart-app-btn { color: red; }` and Team B writes `.catalog-app-btn { color: blue; }`.
    *   *Drawback:* Relies entirely on developers following naming rules; prone to human error.

*(Bonus: Web Components/Shadow DOM provides the only true, browser-native hard encapsulation for CSS, completely preventing bleed in or out).*

## 4. What happens if two different Microfrontends try to manipulate the `window.history` object (like pushing a new URL) at the same time?

**Answer:**
This creates a race condition known as "Router Thrashing" or a routing conflict, which can break the browser's back/forward buttons and cause infinite loops.

**The conflict:**
The Host app is listening to URL changes to know when to mount/unmount apps. If a Remote app aggressively pushes a new URL (using `history.pushState`), it might inadvertently tell the Host app to unmount itself, or it might conflict with another active Remote app trying to read the same URL.

**How to solve it:**
1.  **Centralized Routing Authority:** Only the Host Shell should directly manipulate the global browser history. Remote apps should dispatch custom events (e.g., `CUSTOM_NAVIGATE_EVENT`) to the Host, and the Host performs the actual `history.push`.
2.  **Memory Routing:** For internal navigation within a Remote (e.g., navigating tabs inside the User Profile MFE), the Remote should use a `MemoryRouter` (in React) instead of a `BrowserRouter`. A `MemoryRouter` keeps the URL state internally within the component and does not touch the browser's address bar or history stack, preventing conflicts entirely.
