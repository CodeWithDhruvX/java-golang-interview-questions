# 🗣️ Theory — React Performance Optimization
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "When should you use React.memo and when should you NOT use it?"

> *"React.memo prevents a component from re-rendering when its props haven't changed — it does a shallow comparison. You should use it when: the component re-renders frequently with the same props, the rendering is expensive — like a large list or a chart — and the parent re-renders for unrelated reasons. The classic case is a child component that receives callback props — if the parent doesn't stabilize those with useCallback, memo becomes useless because every render creates a new function reference, which fails shallow equality. When NOT to use it: small, fast components where the comparison overhead outweighs the render cost, and when you haven't profiled and confirmed there's actually a performance problem. Premature optimization with memo is a real pitfall — it adds complexity without benefit."*

---

## Q: "Explain the difference between useMemo and useCallback."

> *"Both are about memoization across renders, but they memoize different things. useMemo memoizes the result of a function — it runs the function and returns its return value, only recomputing when dependencies change. useCallback memoizes the function itself — it returns the same function reference across renders until dependencies change. A key way to remember: useMemo(() => fn(), deps) is equivalent to useCallback(fn, deps) — you could rewrite useCallback in terms of useMemo. Use useMemo for expensive computations — filtering or sorting a large array. Use useCallback when you need a stable function reference — especially when passing callbacks to memoized child components or putting a function in a useEffect dependency array. Don't reach for either without profiling first."*

---

## Q: "How does code splitting work and what is the tradeoff?"

> *"Code splitting is the practice of splitting your JavaScript bundle into smaller chunks that are downloaded on demand rather than all upfront. React's built-in mechanism is React.lazy with dynamic import — you wrap a component in lazy, and when it first renders, React downloads that component's chunk from the server. Suspense lets you show a loading fallback while the chunk downloads. The most impactful split is route-based — each page route becomes its own chunk, so users only download the code for the page they're on. The tradeoff is the network round trip when a chunk is first needed — there's a brief loading delay. You can mitigate this with prefetching — loading the chunk in the background before the user navigates there. React Router has built-in support for this via loader functions in v6."*

---

## Q: "What is list virtualization and when is it necessary?"

> *"List virtualization — sometimes called windowing — is the technique of only rendering DOM nodes for items currently visible in the viewport, not the entire list. If you have 10,000 rows and you render all of them, you have 10,000 DOM nodes — the browser has to lay them out and paint them all, which is slow. With virtualization, you only render maybe 20-30 items visible on screen, and as the user scrolls, you swap in new items and remove off-screen ones. The DOM count stays constant regardless of dataset size. Libraries like react-window and @tanstack/react-virtual implement this. When is it necessary? Once you have more than roughly 100-200 items in a scrollable list that the user will actually scroll through. Below that, the overhead of virtualization isn't worth the complexity."*

---

## Q: "How do Web Vitals relate to React performance?"

> *"Web Vitals are Google's metrics for user experience: LCP measures how fast the largest visible content loads, INP measures how quickly the page responds to interactions, and CLS measures visual layout stability. React directly impacts all three. For LCP, client-side rendering is the biggest culprit — the browser has to download React, run it, and render before the user sees anything. Server-side rendering or Next.js SSG dramatically improves LCP because meaningful HTML arrives from the server immediately. For INP, heavy render work or long JavaScript tasks block the main thread and make the page feel unresponsive — useTransition and avoiding blocking renders helps here. For CLS, async content insertion — image loading, dynamic data — causes layout shifts. Reserving space with skeleton screens prevents that."*

---

## Q: "What is useTransition and how does it help with perceived performance?"

> *"useTransition is a React 18 API that lets you mark state updates as 'transitions' — low priority updates that can be interrupted by more urgent user interactions. The purpose is to keep the UI responsive while heavy rendering work is happening. The classic example is a search box — when the user types, you want the input to update immediately, but the search results re-render might be expensive. You put the results update inside startTransition, which tells React it can defer or interrupt that render if the user keeps typing. React works on the results render in the background, and you get isPending to show a visual indicator. Without this, every keystroke could block the UI. useTransition fundamentally changes the question from 'how do I make this render fast' to 'how do I make the user feel like the app is responsive.'"*

---

## Q: "How would you approach debugging a slow React app?"

> *"My approach is systematic — profile before optimizing. First, I open React DevTools Profiler, record an interaction that feels slow, and look for components with high actualDuration — those are the hot paths. Then I check why they're rendering using the 'why did this render' feature — it shows if it's a prop change, state change, or parent re-render. If it's a parent causing unnecessary renders, I add React.memo. If it's expensive computation inside the component, I add useMemo. If it's too many DOM nodes at once — say a long list — I add virtualization. For initial load performance, I look at bundle size with a bundle analyzer and add code splitting. Finally, I measure Web Vitals in Chrome DevTools or Lighthouse to see if my changes actually moved the needle. I never add memoization without profiling first — it can make things worse."*

---

## Q: "Explain how the React Profiler API works."

> *"The Profiler component is React's built-in tool for measuring render performance in your application, not just in DevTools. You wrap a subtree with Profiler and pass an onRender callback. On every render of that subtree, React calls your callback with timing information: actualDuration is how long the render actually took after memoization; baseDuration is the estimated time if memo didn't skip any children — the baseline. If actualDuration is much less than baseDuration, memo is doing its job. You'll also get phase — whether this was a 'mount' or 'update' — and the id you gave the Profiler, so you can identify which section of your app is slow. You'd use this in staging or with RUM monitoring to catch regressions, since React DevTools Profiler only captures during a recorded session in the browser."*
