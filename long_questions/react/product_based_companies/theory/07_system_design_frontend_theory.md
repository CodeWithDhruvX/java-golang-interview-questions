# 🗣️ Theory — Frontend System Design
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How would you design a large-scale component library?"

> *"A component library starts with design tokens — the raw values: colors, spacing, typography, shadows — defined as CSS variables or JavaScript constants. This is the single source of truth. On top of tokens, you build primitive components — Button, Text, Input, Box — stateless, composable, and fully accessible. These primitives have no business logic, just rendering. On top of primitives, you build composite components — Card, Modal, DatePicker, DataTable — assembled from primitives. The key architectural decisions: a variant system for components so they're configurable without prop explosion; a compound components pattern for complex multi-part components; Storybook for documentation and visual testing; automated accessibility testing with axe in CI; semantic versioning with a changelog. The hardest part is the API design — it needs to be flexible enough for diverse use cases but constrained enough to maintain consistency. You have to resist adding every feature someone requests."*

---

## Q: "What are Micro-Frontends and what problem do they solve?"

> *"Micro-Frontends apply the microservices philosophy to the frontend — instead of one large monolithic frontend application, you split it into independently deployable pieces owned by separate teams. The problem they solve: at large scale, a single frontend repo becomes a bottleneck. Multiple teams are changing the same codebase, deployments block each other, build times grow, and different teams have different tech preferences. With Micro-Frontends, the Checkout team deploys their checkout module independently without coordinating with the Catalog team. Each team can choose their tech stack. Module Federation in Webpack 5 is the modern way to implement this — a shell app loads remote modules from separate deployments at runtime. The tradeoffs: significant infrastructure overhead, shared dependency management is complex, cross-team communication for shared state is harder. It's only worth it at significant organizational scale — typically when you have 5 or more separate teams on a frontend."*

---

## Q: "How do you approach accessibility in a React application?"

> *"Accessibility — a11y — is about making your app usable by people with disabilities, which also improves usability for everyone. My approach has layers. First: semantic HTML. Using button for clickable things, nav for navigation, main for content, label for form inputs — these come with accessibility built in for free. Second: keyboard navigation — interactive elements must be focusable and operable with keyboard. Tab to focus, Enter or Space to activate. Third: ARIA attributes — when native HTML isn't enough. aria-expanded, aria-haspopup, aria-describedby for custom components. Fourth: color contrast — a minimum of 4.5:1 contrast ratio for normal text. Fifth: focus management — when a modal opens, move focus inside it. When it closes, return focus to the trigger element. Tools: eslint-jsx-a11y catches issues statically. axe DevTools extension catches issues at runtime. React can actually make accessibility easier with focus management APIs — but you have to consciously build it in from the start."*

---

## Q: "How would you implement A/B testing in a React application?"

> *"A/B testing — or feature flagging — delivers different variants to different users and measures the impact. The architecture: the assignment of which variant a user gets happens server-side — in your API or a dedicated service like LaunchDarkly — and is deterministic based on some user identifier so the same user consistently sees the same variant across sessions. The flag values come to the frontend either as part of the user session API response or through a flags API call at startup. In React, you wrap them in a provider that makes the flags available throughout the app. Components check their relevant flag with a useFeatureFlag hook. Critically: you should not determine assignments client-side based on Math.random — that creates flickering as the variant changes between renders. The flag should be stable by the time React renders. For measuring impact, you instrument events — variant shown, conversion happened — and send them to your analytics pipeline."*

---

## Q: "How do you design for performance at the architecture level, not just the component level?"

> *"Component-level optimization — memo, useMemo — is often treating symptoms rather than root causes. Architectural performance decisions have much larger impact. First: rendering strategy. CSR makes the user download and run all your JavaScript before seeing content. SSR sends HTML immediately — critical for LCP. Choose SSG for pages that don't need fresh data per request. Second: code splitting at route boundaries so users only download what they need for the current page — a 200KB initial bundle vs 1MB is a massive difference on mobile. Third: image optimization — images are often the largest bytes on a page. Use next/image or a CDN that serves WebP, correct sizes, and lazy-loads. Fourth: caching with a CDN — static assets should be cached for 1 year with immutable headers; the hash in the filename handles cache busting. Fifth: avoid render-blocking resources — fonts especially. Preload critical fonts, use font-display: swap. These architectural choices dwarf component-level optimization."*
