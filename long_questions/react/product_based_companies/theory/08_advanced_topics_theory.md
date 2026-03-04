# 🗣️ Theory — Advanced React Topics
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Explain the difference between SSR, SSG, CSR, and ISR."

> *"These are four rendering strategies for React applications. CSR — Client-Side Rendering — is the traditional SPA approach: the server sends a nearly empty HTML file, the browser downloads React, runs it, fetches data, and renders. The user sees a blank screen until all that happens. SSR — Server-Side Rendering — runs React on the server per request: the server renders full HTML and sends it immediately. The user sees content right away, which dramatically improves LCP and SEO, but the server does more work per request. SSG — Static Site Generation — renders pages at build time and serves them as static files from a CDN. Zero server compute at request time, incredibly fast, but the data is stale as soon as it's deployed. ISR — Incremental Static Regeneration — is Next.js's hybrid: pages are statically generated but can be regenerated in the background after a configurable interval. Old visitors get the cached page immediately; after the interval, Next.js regenerates in the background and subsequent visitors get the fresh version. Choose based on how dynamic your data is and your SEO requirements."*

---

## Q: "How do React Server Components differ from regular components?"

> *"React Server Components — RSC — are a new type of component that runs exclusively on the server and never in the browser. They can directly access databases, read files, use server-side secrets, and make backend function calls — with no API layer needed. Their output is not HTML — it's a serialized React tree that the browser React runtime can render. Critically, Server Components ship zero JavaScript to the client — their code never appears in the browser bundle. They cannot use state, effects, or event handlers — they're purely for data access and rendering. Regular components — Client Components, marked with 'use client' — still run in the browser with state and interactivity. In Next.js App Router, all components are Server Components by default. You add 'use client' when you need interactivity. The performance win: a page's data-fetching component runs on the server with direct DB access, adds zero JS to the bundle, and streams HTML down. Only the interactive buttons and inputs are Client Components."*

---

## Q: "What is streaming SSR and how does it improve user experience?"

> *"Traditional SSR has a limitation: the server must finish rendering the entire page — all data must be fetched — before it can send any HTML to the browser. If one data source is slow, the whole page waits. Streaming SSR changes this by sending HTML progressively, in chunks, as it becomes ready. React 18 enables this natively. The shell of the page — header, layout, skeleton — is sent immediately. Individual sections are wrapped in Suspense boundaries. As each section's data resolves on the server, React streams that HTML chunk to the browser, which fills it in. From the user's perspective: they see the page shell almost instantly — great TTFB — then watch sections fill in as data arrives — like loading skeletons being replaced — rather than staring at a blank page. The LCP dramatically improves because the browser starts rendering before all data is ready. In Next.js App Router, Suspense plus async Server Components automatically enables streaming — you get it for free with the right architecture."*

---

## Q: "How do you use TypeScript generics with React hooks?"

> *"Generics let you write hooks that are type-safe without being tied to a specific data type. The classic example is useFetch or useLocalStorage — they should work with any type of data. You declare the type parameter with angle brackets after the function name: function useFetch<T>(url: string) returns an object containing T | null for data. TypeScript infers the type when you call it: const { data } = useFetch<User[]>('/api/users') — and now TypeScript knows data is User[] or null. For custom hooks that wrap useState, typing is straightforward because useState is already generic. For hooks that wrap useRef, the type goes in the angle brackets: useRef<HTMLInputElement>(null). For custom hooks that accept callbacks, you often need to type the callback generically too. The key principle: type the inputs and outputs of the hook, not the internals. The consumer of the hook should get full type safety without knowing about the hook's implementation."*

---

## Q: "How do you optimize a React app's initial load performance?"

> *"Initial load performance is about reducing the time until the user sees and can interact with something meaningful. The biggest levers: First — bundle size. Use a bundle analyzer to identify large dependencies. Replace heavy libraries with lighter alternatives — moment.js with date-fns, lodash with native methods. Code split aggressively at route boundaries with React.lazy — this is often the highest-impact change. Second — rendering strategy. Move from CSR to SSR or SSG with Next.js. This moves the initial render cost to the server and sends HTML instead of an empty page. Third — resource loading. Preload critical resources — your main font, your logo. Use font-display: swap to avoid invisible text during font load. Compress images, use WebP format, lazy load below-the-fold images with the loading='lazy' attribute. Fourth — caching. Configure Nginx or your CDN to serve all JS and CSS with long cache headers — the hash in the filename handles busting. Fifth — critical CSS. Inline the CSS your above-the-fold content needs so it renders without a network round trip."*
