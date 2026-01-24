# JavaScript Interview Questions & Answers (951-1000)

## 🔹 36. Advanced Modules, Imports & Bundling (Questions 951-960)

**Q951: What is the difference between `import` and `require`?**
*   **`require` (CommonJS):** Synchronous, dynamic (can call anywhere). Value is a copy (primitive) or reference (object). Node.js default.
*   **`import` (ESM):** Asynchronous (static analysis), top-level only (mostly). Live bindings to values. Browser standard.

**Q952: How does tree shaking work with ES modules?**
Bundlers (Webpack/Rollup) rely on the static structure of ESM (`import`/`export`). They literally check which exports are imported and used. Code that is not reachable from the entry point is dropped from the bundle.

**Q953: What are dynamic imports and how do they improve performance?**
`import('path')`. It allows loading code lazily (on demand, e.g., on click). It splits the bundle into smaller chunks, reducing initial load size and time-to-interactive.

**Q954: Can you `import` conditionally based on runtime values?**
Only with Dynamic Import (`import()`). Static `import ... from ...` must be at the top level and cannot be inside `if` blocks.

**Q955: What is the default module loading strategy in browsers?**
Deferred (`<script type="module">` behaves like `defer`). It downloads in parallel but executes in order after HTML parsing is complete.

**Q956: How do circular imports behave in ES Modules?**
They are supported via **Live Bindings**. If Module A imports B, and B imports A, the import in A might initially be `undefined` if accessed immediately during execution phase, but once B finishes execution, the binding in A updates to the correct value.

**Q957: What is the difference between `export default` and `named exports`?**
*   **Named:** Explicit names (`export const a`). Must import with braces (`import { a }`). Good for tooling (tree shaking).
*   **Default:** One per file (`export default val`). Import with any name (`import MyVal`). Harder to refactor.

**Q958: How does `import.meta` provide context in ES modules?**
It exposes metadata about the current module. Notably `import.meta.url` gives the absolute URL of the module file (useful for resolving relative asset paths).

**Q959: What is module hoisting in bundlers like Webpack?**
(Scope Hoisting). It concatenates the scope of all modules into a single closure (where possible) rather than warping each module in a separate function. This reduces function closure overhead and file size.

**Q960: What are the advantages of ES modules over CommonJS?**
*   Standard (Browser + Node).
*   Static analysis (Tree Shaking, faster tools).
*   Async loading support.
*   Live bindings (better for circular deps).

---

## 🔹 37. Animation, Timing, and Visuals (Questions 961-970)

**Q961: How does `requestAnimationFrame()` optimize rendering?**
It aligns execution with the browser's refresh rate (vsync). The browser can merge multiple `rAF` calls and style changes into a single reflow/repaint, avoiding frame loss.

**Q962: What is the difference between `setInterval()` and `requestAnimationFrame()`?**
*   `setInterval`: Time-based. drifts. Runs even if tabs are hidden (throttled slightly). Can cause frame drops if logic takes too long.
*   `rAF`: Frame-based. Pauses when tab is hidden. Guaranteed to run before paint.

**Q963: How can you create a smooth progress bar with JS?**
Use `requestAnimationFrame` to interpolate the width value over time, or better, use CSS transitions/animations and toggle classes/variables with JS. Using `transform: scaleX()` is more performant (Composite only) than `width` (Layout).

**Q964: How does JavaScript throttling improve animation performance?**
It limits how often computation runs (e.g., on scroll). However, for visual updates, `requestAnimationFrame` is better than throttling. Throttling is better for logic (e.g., checking positions) triggered by events.

**Q965: What are scroll-driven animations and how can you implement them?**
Animations linked to scroll position.
*   **JS:** Listen to `scroll` (passive), calculate % scrolled, set style in `rAF`.
*   **CSS (Modern):** `animation-timeline: scroll()`.

**Q966: How to synchronize CSS animations with JS logic?**
Use standard events: `animationstart`, `animationend`, `transitionend`. Or use the Web Animations API (`element.animate()`) which returns a Promise/Timeline object.

**Q967: How to pause/resume an animation using JS?**
*   **CSS:** `element.style.animationPlayState = 'paused' | 'running'`.
*   **WAAPI:** `animation.pause()`, `animation.play()`.

**Q968: How do you animate with `transform` vs `top/left` properties?**
*   `top/left`: Triggers **Layout** (calculating geometry) + Paint + Composite. Slow.
*   `transform`: Triggers only **Composite**. Handled by GPU. Fast/Smooth (60fps).

**Q969: What is the FLIP animation technique in JavaScript?**
**F**irst, **L**ast, **I**nvert, **P**lay.
1. Measure start position (First).
2. Apply change (Last).
3. Transform element backwards to match Start (Invert).
4. Remove transform with transition (Play).
Result: Smooth layout animations using only `transform`.

**Q970: How does `IntersectionObserver` help with lazy animations?**
It detects when an element enters the viewport. You can keep the animation paused/hidden (will-change) and only start playing it when the user sees it, saving CPU/GPU.

---

## 🔹 38. Mobile, Accessibility, Device APIs (Questions 971-980)

**Q971: How do you detect device orientation changes in JavaScript?**
`window.addEventListener('orientationchange')` or `matchMedia('(orientation: portrait)').addListener()`.

**Q972: What is the DeviceMotionEvent and when is it useful?**
Provides data about physical acceleration (accelerometer) and rotation rate (gyroscope). Used for games, gesture controls (shake to undo).

**Q973: How to detect touchscreen devices using JS?**
`'ontouchstart' in window` or `navigator.maxTouchPoints > 0`. (Note: Not 100% reliable as some laptops have both).

**Q974: How do you improve accessibility with ARIA attributes?**
Use `aria-label`, `aria-expanded`, `aria-hidden` to describe state and content to Screen Readers when semantic HTML isn't enough (e.g., custom dropdowns).

**Q975: What is the difference between `pointerdown`, `mousedown`, and `touchstart`?**
*   `pointerdown`: Unified event (Mouse, Touch, Pen). Best for modern apps.
*   `mousedown`: Mouse only.
*   `touchstart`: Touch only.

**Q976: How to detect screen reader usage with JS?**
You **cannot** reliably detect if a screen reader is active (privacy design). You must build accessible code by default.

**Q977: What is responsive font scaling using `window.devicePixelRatio`?**
It allows detecting High-DPI (Retina) screens. You typically handle scaling via CSS (`rem`, viewport units), but JS can use `devicePixelRatio` to load higher-res Canvas bitmaps.

**Q978: How do you enable keyboard-only navigation detection?**
Listen for `keydown` (Tab). Add a class (e.g., `user-is-tabbing`) to `body` to show focus outlines. Remove it on `mousedown`. (Libraries like `what-input`).

**Q979: How do you support dark mode toggling in JS?**
`window.matchMedia('(prefers-color-scheme: dark)')`. Listen for changes. Toggle a `dark-theme` class on `<html>` or update CSS Variables.

**Q980: How do you detect and respond to viewport resize on mobile?**
Use `ResizeObserver` or `window.onresize`. Handle "Virtual Keyboard" appearing which shrinks the viewport height (`visualViewport` API).

---

## 🔹 39. Real-World Scenarios & Behavioral Patterns (Questions 981-990)

**Q981: How do you detect if a user is idle?**
Listen for `mousemove`, `keydown`, `scroll`, `click`. Reset a timer on each event. If timer fires (e.g. after 5m), user is idle. Or use `IdleDetector` API.

**Q982: How do you debounce API requests on search input?**
(See Q132). Wrap the fetch call in a `debounce` function (wait 300ms after last keystroke).

**Q983: How to manage state between tabs using `localStorage`?**
Listen to the `storage` event. It fires on other tabs when localStorage changes. Sync state (Redux/Context) when event fires.

**Q984: How to handle stale cache data in a SPA?**
Versioning. Store `version` in localStorage. On load, if `store.version !== app.version`, clear storage/cache and reload.

**Q985: What happens if two tabs write to localStorage at once?**
It is synchronous and blocking operation (usually). One write wins. No race condition corruption internally, but logic race condition exists (Last Write Wins).

**Q986: How to persist scroll position across navigation in a SPA?**
Save `window.scrollY` to `sessionStorage` on unload/routeChange. Restore on load. (Or standard `history.scrollRestoration`).

**Q987: How to sync Redux or global state with URL query parameters?**
Listener on state change -> `history.replaceState`. Listener on `popstate` -> generic action to update store.

**Q988: How to manage undo/redo stack in JS?**
(See Q699 Commmand Pattern). Stack of previous states.

**Q989: How to prevent form resubmission on reload?**
Post/Redirect/Get pattern. Or `history.replaceState()` to remove query params/state after submit.

**Q990: How do you implement optimistic UI updates?**
Update UI immediately (optimistic). Send request. If fail, revert UI (rollback).

---

## 🔹 40. Web APIs & Network Strategies (Questions 991-1000)

**Q991: What is HTTP/2 Push and how can JavaScript use it?**
Server sends resources before client asks. JS doesn't "use" it directly; browser cache handles it. (Deprecated in Chrome 106+).

**Q992: What is fetch keepalive used for?**
(Duplicate Q196).

**Q993: What is a streaming response in fetch and when should you use it?**
`response.body.getReader()`. Read streams chunk by chunk. (Video, huge JSON, progress indicators).

**Q994: What is the purpose of `navigator.sendBeacon()`?**
(Duplicate Q797).

**Q995: What is the difference between offline-first and cache-first strategies?**
*   **Offline-first**: App assumes offline. Serves from cache, syncs later.
*   **Cache-first**: Try cache. If missing, fetch network. (Speed focus).

**Q996: How do you update cached assets in a PWA?**
Service Worker versioning. If `sw.js` changes, new one installs.

**Q997: What is stale-while-revalidate in service workers?**
Serve from cache immediately (fast), but update cache from network in background (fresh next time).

**Q998: How do you track network failures and retries in JavaScript?**
`window.addEventListener('online/offline')`. Catch fetch errors.

**Q999: What is background sync and how is it implemented?**
Service Worker `sync` event. Queues requests when offline, sends them when connectivity returns.

**Q1000: How to handle slow internet or flaky connections in JS UIs?**
Show skeletons/spinners. Use timeouts. Retry logic. Optimistic UI.
