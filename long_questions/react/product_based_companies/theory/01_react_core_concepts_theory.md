# 🗣️ Theory — React Core Concepts
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Can you explain how the Virtual DOM works?"

> *"The Virtual DOM is essentially a lightweight JavaScript representation of the real browser DOM — just plain objects. When your component's state or props change, React creates a brand new Virtual DOM tree representing what the UI should look like. Then it runs a 'diffing' algorithm — it compares the new tree against the previous snapshot — and figures out the minimum set of changes needed. Finally, it applies only those specific changes to the real DOM. The key insight is that DOM operations are expensive, but comparing JavaScript objects is very fast. So React batches work in memory and minimizes real DOM touches. This process is called reconciliation."*

---

## Q: "What is React Fiber and why was it introduced?"

> *"Fiber is React's completely rewritten reconciliation engine introduced in React 16. The old reconciler — called the Stack Reconciler — was synchronous and recursive. Once it started rendering a component tree, it couldn't stop until it was done. This caused jank — dropped frames — in complex UIs because the main thread was blocked. Fiber breaks work into small units called 'fibers' — one per component — and introduces a priority-based scheduler. Now React can pause rendering mid-way, yield control back to the browser for things like user input, and resume later. It also enables Concurrent Mode features like startTransition, where you can mark state updates as low priority. Fiber is what made React capable of smooth animations and responsive UIs even with complex trees."*

---

## Q: "Why are keys important in React lists and what happens without them?"

> *"Keys are how React identifies which items in a list changed, moved, or were added or removed. When you render a list without keys, React diffs by position — if you prepend an item, every existing item appears to have changed, causing all of them to unmount and remount. That's expensive. With stable unique keys — ideally a database ID — React can match elements across renders by identity rather than position. So it only updates what actually changed. The classic mistake is using the array index as a key — that seems harmless but breaks badly when you filter or reorder the list, because indexes shift and React gets confused about which item is which. Always use a stable, unique ID from your data."*

---

## Q: "Explain controlled versus uncontrolled components."

> *"A controlled component is one where React fully owns the state — you pass a value prop and an onChange handler, and every keystroke goes through React state. You have complete control: you can format input, validate on every keystroke, or prevent invalid characters. An uncontrolled component lets the DOM own the state — you just attach a ref and read the value when you need it. Controlled is the React way and is most common. Uncontrolled is useful for file inputs — because you can't programmatically set a file input's value — and when integrating third-party non-React libraries that need to own the DOM element. The tradeoff is: controlled is more powerful and predictable; uncontrolled is simpler for straightforward cases."*

---

## Q: "What is the Context API and when is it NOT the right tool?"

> *"Context is React's built-in way to pass data through the component tree without prop drilling — you create a context, wrap a subtree with a Provider, and any nested component can consume it with useContext. It's great for stable data that many components need — current user, theme, locale, auth state. But the critical performance caveat is: every consumer re-renders whenever the context value changes. If you put frequently updating state in context — like data that changes on every keypress — you'll trigger re-renders throughout your tree. For high-frequency updates, you want a subscription-based solution like Zustand or Jotai, where only the subscribed components update. Context is also not ideal for complex derived state — you'd want a selector pattern that Redux or Zustand provide."*

---

## Q: "How does React's batching of state updates work in React 18?"

> *"Batching means React groups multiple setState calls into a single re-render — so calling setCount and setFlag together triggers one render, not two. Before React 18, batching only happened inside React event handlers. If you called setState inside a setTimeout or a fetch callback, each call triggered its own render. React 18 introduced Automatic Batching — now all state updates are batched everywhere, including async code. This is a free performance improvement. If you actual need immediate DOM updates between state changes — rare, but it happens in animations or third-party integrations — you opt out with flushSync from react-dom. ReactDOM.createRoot is what enables this — the legacy ReactDOM.render API doesn't get automatic batching."*

---

## Q: "What does React StrictMode do in development?"

> *"StrictMode is a development-only wrapper that deliberately causes certain behaviors to happen twice to help you catch bugs. Specifically: it double-invokes your component render phase, your useState initializers, and your useEffect setup-and-cleanup cycle. Why? To help you find side effects in places that should be pure or properly cleaned up. If your component renders twice and shows different data each time, you have a mutation in the render phase — a bug. If your useEffect runs twice and breaks your app, your cleanup function is wrong. In production, StrictMode has zero impact — it's completely stripped out. The 'my useEffect runs twice in development' confusion almost always comes from StrictMode doing exactly what it's supposed to do."*

---

## Q: "How does JSX compilation work and what changed in React 17?"

> *"JSX is not native JavaScript — it's syntactic sugar that a compiler, usually Babel, transforms into JavaScript. Before React 17, JSX compiled to React.createElement calls, which meant you had to import React in every file that used JSX — even if you never directly referenced React. React 17 introduced a new JSX transform that compiles JSX to calls to jsx() and jsxs() from the 'react/jsx-runtime' module — which the compiler imports automatically. So you no longer need to write 'import React from react' at the top of every file. The underlying representation is still the same — a plain object with type, props, and key — but the import is automatic. This was a developer experience improvement and slightly reduces bundle size by removing those import statements."*
