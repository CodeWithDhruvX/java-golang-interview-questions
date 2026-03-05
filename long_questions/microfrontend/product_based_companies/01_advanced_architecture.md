# Advanced Architecture & System Design for Microfrontends

This section covers high-level architectural decisions, trade-offs, and scaling strategies for Microfrontend (MFE) applications, typically asked in senior-level interviews at product-based companies.

## 1. How do you decide when to use a Microfrontend architecture versus a Monolith, and what are the primary trade-offs?

**Answer:**
The decision should be driven by organizational structure and scaling pain points, not just technology.

**When to use Microfrontends:**
*   **Large, Autonomous Teams:** When multiple independent teams (e.g., Domain-Driven Design) need to work on the same product simultaneously without blocking each other.
*   **Independent Deployments:** When different parts of the application require different release cycles (e.g., releasing a marketing page daily vs. a core banking feature monthly).
*   **Legacy Modernization:** Strangler fig pattern - slowly rewriting parts of a legacy application in a modern stark while keeping the rest intact.

**When NOT to use Microfrontends:**
*   Small teams or single teams where the overhead of MFE infrastructure outweighs the benefits.
*   Highly coupled applications where data and state are deeply intertwined across all views.

**Primary Trade-offs:**
*   **Pros:** Independent deployments, tech stack agnosticism (though not recommended to abuse), team autonomy, fault isolation.
*   **Cons:** Increased complexity in CI/CD, difficult local development environment, potential performance hits (duplicate dependencies), and complex cross-MFE state management.

## 2. Compare the different methods of composing Microfrontends. In a high-traffic e-commerce application, which would you choose and why?

**Answer:**
There are mainly three ways to compose MFEs:

1.  **Server-Side Integration (e.g., Edge Side Includes (ESI), Server-Side Rendering composition):**
    *   *How it works:* The server fetches HTML fragments from different services and stitches them together before sending the final page to the browser.
    *   *Pros:* Excellent SEO, fast Initial Load Time (TTFB), great for static or semi-static content.
    *   *Cons:* High server load, challenging to manage highly interactive, real-time state across fragments.
2.  **Build-Time Integration (e.g., NPM packages):**
    *   *How it works:* MFEs are published as NPM packages and bundled into a single application at build time.
    *   *Pros:* Great performance (optimized bundle), easy to implement initially.
    *   *Cons:* Defeats the purpose of independent deployments. A change in a child MFE requires rebuilding and deploying the entire shell.
3.  **Run-Time Integration via Client-Side (e.g., Webpack Module Federation, IFrames, Single-SPA):**
    *   *How it works:* JavaScript bundles are loaded dynamically in the browser at runtime.
    *   *Pros:* True independent deployments, highly dynamic.
    *   *Cons:* Can impact performance (multiple network requests), complex tooling required (like Module Federation), SEO can be harder to orchestrate (requires SSR hydration strategies).

**For a high-traffic e-commerce application:**
I would opt for a hybrid approach or **Client-Side Integration using Webpack Module Federation** (with a well-defined shared dependency strategy).
*   *Why:* E-commerce requires rapid, independent feature releases (e.g., checkout team vs. product catalog team). The catalog might benefit from Server-Side composition for SEO, but the highly interactive cart and checkout flows are best served via Client-Side runtime integration to ensure a seamless App-like experience without full page reloads.

## 3. How do you ensure consistent UI/UX and styling across decoupled Microfrontends built by different teams?

**Answer:**
Maintaining visual consistency is one of the hardest parts of MFE architecture.

1.  **Shared Design System / Component Library:**
    *   Create a centralized, versioned UI component library (e.g., built with Storybook) that exposes primitives (buttons, inputs, typography).
    *   *Important:* This library must be strictly a set of "dumb" components without business logic to prevent tight coupling.
2.  **CSS Isolation Strategies:**
    *   **CSS Modules / Styled Components (CSS-in-JS):** Scopes CSS locally to the component, preventing style leaks across MFEs.
    *   **Shadow DOM (Web Components):** Provides hard native encapsulation for styles and markup, ensuring zero leakage, though it can make global theming slightly more complex.
    *   **Namespacing:** (Less modern) Prefixing all CSS classes with the MFE's name (e.g., `.cart-checkout-btn`).
3.  **Design Tokens:**
    *   Abstract colors, spacing, and typography into JSON/CSS variable tokens. Even if an MFE isn't using the shared React library (e.g., a legacy Vue app), it can still consume the design tokens to look consistent.

## 4. What are the performance implications of Microfrontends, and how do you mitigate "bundle bloat"?

**Answer:**
The biggest risk of runtime MFEs is downloading the same library (like React, Lodash, or a Design System) multiple times.

**Mitigation Strategies:**
1.  **Shared Dependencies via Webpack Module Federation:**
    *   Configure the `shared` array in Module Federation. You can define libraries like `react` and `react-dom` as singletons.
    *   *How it works:* The host application loads React once. When a remote MFE loads, it checks if the required version of React is already in the browser's context. If it is, it uses the host's version; if not (or if versions are incompatible), it downloads its own.
2.  **Strict Dependency Governance:**
    *   Establish architectural guidelines limiting the use of heavy third-party libraries. If every team chooses a different charting library, performance will tank.
3.  **Performance Budgets:**
    *   Implement strict CI/CD performance budgets (e.g., using Lighthouse CI). If a remote MFE's entry bundle exceeds a certain size, the build fails.
4.  **Lazy Loading and Code Splitting:**
    *   Ensure each MFE code-splits its routes heavily. Do not load the "Settings" MFE bundle until the user actually navigates to the Settings page.

## 5. Describe a routing architecture for a large-scale MFE application. How does the Shell application intercept and route traffic to the correct MFE?

**Answer:**
Routing typically happens at two levels: Global (App Shell) and Local (Remote MFE).

1.  **The App Shell (Host):**
    *   The App Shell owns the primary layout (Header, Sidebar) and the global router.
    *   It maintains a configuration map of routes to Remote Applications. For example:
        *   `/products/*` -> loads `Catalog_MFE`
        *   `/checkout/*` -> loads `Checkout_MFE`
    *   When the URL changes, the App Shell's router intercepts it, dynamically fetches the entry file for the corresponding MFE, mounts it into a designated DOM node, and passes control.
2.  **The Remote MFE (Micro-app):**
    *   Once mounted, the remote MFE uses its own internal router (e.g., React Router) to handle sub-routes.
    *   *Critical Rule:* The Remote router must use a `MemoryRouter` or be configured with a `basename` matching the shell's mount path (e.g., `<BrowserRouter basename="/products">`).
    *   This prevents the Remote MFE from trying to control the global browser history directly, which can cause conflicts with the shell or other MFEs.

**Handling 404s:**
The Shell router attempts to match first. If no MFE matches the base path, the Shell shows a global 404. If the Shell matches `/products/*` but the `Catalog_MFE` doesn't recognize the specific ID (e.g., `/products/invalid-id`), the *Remote* MFE takes over and renders its internal 404 boundary.
