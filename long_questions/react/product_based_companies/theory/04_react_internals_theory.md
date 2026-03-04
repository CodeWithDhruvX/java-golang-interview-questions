# 🗣️ Theory — React Internals & Architecture
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Can you explain the two phases of React rendering?"

> *"React rendering has two distinct phases. The Render Phase — sometimes called Reconciliation — is where React runs your component functions and builds a new Fiber tree representing the desired UI. This phase is interruptible in Concurrent Mode — React can pause it and come back later. It must be pure — no side effects allowed here. The Commit Phase is where React applies the changes to the real DOM — this is synchronous and cannot be interrupted. The commit phase itself has three sub-phases: Before Mutation — where getSnapshotBeforeUpdate runs; Mutation — where DOM insertions, updates, and deletions happen; and Layout — where useLayoutEffect and componentDidMount run synchronously after DOM mutation. After the commit is fully done, useEffect callbacks are scheduled and run asynchronously so they don't block painting."*

---

## Q: "What is Concurrent Mode and how does it change React's behavior?"

> *"Concurrent Mode — enabled by createRoot in React 18 — is a fundamental change to how React renders. In legacy mode, rendering was synchronous and blocking. Once React started rendering a component tree, it had to finish before the browser could do anything else — including responding to user input. In Concurrent Mode, React can work on rendering in the background, pause when higher-priority work comes in — like a user clicking a button — handle that immediately, then resume the background work. It's like having a task scheduler built into React. The startTransition API is how you tell React something is low priority. useDeferredValue lets React render with an old value while new work is pending. This enables UIs that stay responsive even with expensive renders happening in the background. You get this for free when you upgrade to React 18's createRoot."*

---

## Q: "How do Error Boundaries work and what are their limitations?"

> *"Error Boundaries are class components — there's no hook equivalent yet — that catch JavaScript errors anywhere in their child component tree during rendering, in lifecycle methods, and in constructors. They use two special lifecycle methods: getDerivedStateFromError to update state so the next render shows a fallback UI, and componentDidCatch to log the error to an error reporting service like Sentry. The key limitations: Error Boundaries do not catch errors in event handlers — you handle those with regular try-catch inside the handler. They don't catch async errors — an unhandled promise rejection in useEffect doesn't trigger an Error Boundary. And they don't catch errors during server-side rendering. The practical usage is to wrap sections of your UI — like a sidebar or a widget — so that one crashing component doesn't take down the whole page. The react-error-boundary library is the modern way to use them without writing class components yourself."*

---

## Q: "What are React Portals and when do you use them?"

> *"A Portal lets you render a child component's DOM output into a different DOM node than its parent in the React tree. The React tree is unchanged — prop drilling, context, events all work normally — but the rendered HTML appears somewhere else in the real DOM. The canonical use case is modals, tooltips, and dropdowns. The problem these solve: if your modal is rendered inside a div with overflow hidden or a low z-index, it's clipped or hidden by its container. By rendering it directly into document.body via a Portal, it escapes all ancestor CSS constraints. Another important detail: even though the Portal renders to document.body, React events still bubble up through the React component tree — not the DOM tree. So clicking inside a Portal still triggers React event handlers on ancestor components normally."*

---

## Q: "Explain the difference between useEffect and useLayoutEffect."

> *"Both useEffect and useLayoutEffect run after a render and accept a dependency array — the API is identical. The difference is timing relative to the browser's paint cycle. useLayoutEffect runs synchronously after React commits DOM changes and before the browser paints. useEffect runs asynchronously after the browser has painted. Why does this matter? If you need to read the DOM — measure an element's dimensions, scroll position, or computed style — and potentially update state based on that measurement, useLayoutEffect prevents the visual flicker. With useEffect, there's a brief moment where the browser paints the old state, then you update and paint again — the user sees a flash. useLayoutEffect prevents that by doing the measurement and update in the same frame, before anything was painted. The rule of thumb: use useEffect for everything by default — data fetching, subscriptions, logging. Only reach for useLayoutEffect when you're measuring the DOM and updating state in response."*

---

## Q: "What is forwardRef and why do you need it?"

> *"Normally, refs in React are handled transparently — you put a ref on a DOM element and it points to that DOM node. But when you wrap a DOM element in a custom component, the ref behavior is lost — if you put a ref on your FancyInput component, it doesn't automatically reach the underlying input element inside. forwardRef is the API to explicitly pass through: you wrap your component with React.forwardRef and it gives your function a second argument — ref — which you then attach to the inner DOM element or component you want to expose. The second API, useImperativeHandle, takes this further — instead of exposing the raw DOM node, you expose a custom object with specific methods. This is the principle of least privilege: rather than giving the parent full access to the DOM node with all its properties, you expose just play(), pause(), and seek() — the methods that make sense for a VideoPlayer API."*

---

## Q: "How does Suspense work with data fetching in React 18?"

> *"Suspense was introduced for code splitting — React.lazy — but React 18 extends it to data fetching. The concept is: a component can 'suspend' during rendering by throwing a Promise. React catches it, renders the nearest Suspense boundary's fallback, and waits for the Promise to resolve. Once resolved, React re-renders the suspended subtree. The key is that the library — React Query, Relay, or Next.js — must implement this protocol. With React Query's useSuspenseQuery, you no longer need to check isLoading or handle the undefined data case — the component simply renders as if data is always ready, because if it's not, it suspends. Error Boundaries catch any thrown errors, so you pair Suspense with ErrorBoundary. The power is in nested Suspense boundaries — different parts of your UI can independently show loading states and stream in as their data resolves, enabling progressive page reveals."*
