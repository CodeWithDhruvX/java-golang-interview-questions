# Frontend System Design (Product-Based Companies)

Evaluating frontend system design involves discussing how you would architect a large-scale web application from scratch. This goes beyond framework specifics and touches on state management, performance, network requests, and overall architecture.

## High-Level Architecture

### 1. How would you design a micro-frontend architecture? What are the pros and cons?
**Micro-frontends** extend the microservices concept to the frontend. The web app is broken down into features, each owned by an independent team right up to the database.
**Implementation Approaches:**
*   **Build-time integration**: Installing separate NPM packages containing components. (Requires releasing the host app for every update - often avoids the true purpose of micro-frontends).
*   **Run-time integration via Iframes**: Easiest isolation, but terrible for UX, routing, and shared state.
*   **Run-time integration via JavaScript (Module Federation in Webpack 5)**: The current standard. Allows dynamically loading code from other builds at runtime without server-side routing.
**Pros**: Independent deployments, autonomous teams, technology agnosticism (e.g., treating a Vue app and React app seamlessly).
**Cons**: Complex CI/CD pipelines, payload size redundancy (loading React multiple times if not shared properly), complex state sharing, inconsistent styling if not governed.

### 2. Compare SSR (Server-Side Rendering), CSR (Client-Side Rendering), SSG (Static Site Generation), and ISR (Incremental Static Regeneration).
*   **CSR (e.g., plain Create React App, standard Angular)**: HTML sent to the browser is barebones. JS downloads and renders the UI in the browser. *Pros*: Great highly interactive app feel post-load. *Cons*: Poor SEO, slow Time to Interactive (TTI) on slow devices.
*   **SSR (e.g., Next.js `getServerSideProps`, Angular Universal)**: Server fetches data and generates full HTML string per request. *Pros*: Excellent SEO, fast First Contentful Paint (FCP). *Cons*: Slower Time to First Byte (TTFB), higher server load.
*   **SSG (e.g., Next.js `getStaticProps` without revalidate)**: HTML is generated at *build time*. *Pros*: Blazing fast (served from CDN), highly secure, great SEO. *Cons*: Requires full rebuild when data changes. Not suitable for highly dynamic data.
*   **ISR (e.g., Next.js `getStaticProps` with revalidate)**: SSG, but the static pages automatically rebuild in the background after a specified timeout when traffic hits them. *Pros*: Best of both SSG and SSR—fast initial load but data stays relatively fresh.

## Web Performance Optimization

### 3. What are Core Web Vitals, and how do you optimize for them?
Core Web Vitals are Google's metrics for measuring user experience.
*   **Largest Contentful Paint (LCP)**: Measures loading performance. The time it takes for the largest image/text block to render.
    *   *Optimization*: Use CDNs, optimize images (WebP, lazy load off-screen images), preload critical assets, implement SSR/SSG.
*   **Interaction to Next Paint (INP) / First Input Delay (FID)**: Measures responsiveness. The time from when a user first interacts (clicks) to when the browser responds.
    *   *Optimization*: Break up long tasks in JS, minimize main-thread work, defer non-critical JS, use web workers.
*   **Cumulative Layout Shift (CLS)**: Measures visual stability. How much the UI shifts unexpectedly during load.
    *   *Optimization*: Always include `width` and `height` attributes on images/videos, statically reserve space for ad slots, never insert content *above* existing content dynamically without user interaction.

### 4. How do you implement lazy loading and code splitting in modern JS applications?
Code splitting allows you to split your bundle into multiple smaller chunks that are loaded on demand, rather than loading one massive `bundle.js`.
*   **Route-based splitting**: Only load the JS required for the current page/route. (e.g., using `React.lazy()` or Angular's `loadChildren` in the router).
*   **Component-based splitting**: Loading heavy components (like a complex chart library or a modal) only when they are needed.
*   Implemented using dynamic `import()` statements, which bundlers like Webpack/Vite automatically recognize and split.

## State Management and Networking

### 5. In a large React application, when would you choose Context API vs. Redux vs. Zustand/Jotai?
*   **Context API**: Good for low-frequency updates affecting many components (Theme, Locale, Auth status). Terrible for high-frequency state due to unnecessary re-renders of all consumers when a single property changes.
*   **Redux**: Excellent for massive applications with highly complex state logic requiring strict traceability (Redux DevTools), middleware (Redux Thunk/Saga), and predictable unidirectional data flow. Boilerplate-heavy.
*   **Zustand / Recoil / Jotai**: Modern alternatives. Zustand provides a Redux-like centralized store with practically zero boilerplate. Recoil/Jotai use an "atomic" state model, ideal for highly scattered, independently updating pieces of state (like hundreds of nodes in a canvas editor).

### 6. How would you handle network resilience in the frontend?
*   **Caching**: Implement HTTP caching strategies (ETags, Cache-Control). Use a Service Worker (PWA) to cache API responses for offline capability.
*   **Optimistic UI Updates**: Update the UI immediately assuming the API request will succeed. Revert the UI if the request fails. Makes the app feel incredibly fast.
*   **Retries and Backoff**: If an API request fails due to a network glitch, automatically retry with exponential backoff. Wait 1s, then 2s, then 4s, etc., to avoid hammering a struggling server.
*   **Debouncing / Throttling**: Restrict the frequency of API calls (e.g., hitting the search API only 300ms *after* the user stops typing a search query using debounce).
