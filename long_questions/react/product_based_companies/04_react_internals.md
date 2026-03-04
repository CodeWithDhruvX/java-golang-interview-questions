# 🔬 04 — React Internals & Architecture
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

> Understanding React's internals is what separates senior engineers at companies like Meta and Google from mid-level developers.

---

## 🔑 Must-Know Topics
- React Fiber work loop and phases
- Concurrent Mode and Suspense
- Error boundaries
- Portals
- `forwardRef` and `useImperativeHandle`
- `startTransition` and priority lanes
- React 18 new APIs

---

## ❓ Most Asked Questions

### Q1. Explain the React Fiber work loop phases

```
React Fiber rendering has TWO phases:

╔══════════════════════════════════════════════════════╗
║  RENDER PHASE (Reconciliation)                       ║
║  - Can be interrupted, paused, resumed               ║
║  - Builds a new Fiber tree (work-in-progress)        ║
║  - Runs: render(), getDerivedStateFromProps(),       ║
║    shouldComponentUpdate(), hooks run (mount/update) ║
║  - Side-effect free — no DOM mutations               ║
╚══════════════════════════════════════════════════════╝
              ↓
╔══════════════════════════════════════════════════════╗
║  COMMIT PHASE                                        ║
║  - Synchronous, cannot be interrupted                ║
║  - Actually applies changes to the real DOM          ║
║  - Sub-phases:                                       ║
║    1. BeforeMutation — getSnapshotBeforeUpdate       ║
║    2. Mutation — insert/update/delete DOM nodes      ║
║    3. Layout — useLayoutEffect, componentDidMount    ║
║  - After commit: useEffect runs asynchronously       ║
╚══════════════════════════════════════════════════════╝
```

```jsx
// useLayoutEffect vs useEffect
useLayoutEffect(() => {
  // Runs SYNCHRONOUSLY after DOM mutations, before paint
  // Use for: measuring DOM, animations that prevent flicker
  const { height } = ref.current.getBoundingClientRect();
  setHeight(height);
}, []);

useEffect(() => {
  // Runs AFTER paint — doesn't block browser
  // Use for: data fetching, subscriptions, logging
}, []);
```

---

### Q2. How does Concurrent Mode work in React 18?

```jsx
// Concurrent Mode: React can work on multiple renders simultaneously
// and interrupt low-priority work to handle high-priority updates

// Enable: ReactDOM.createRoot() (React 18 default)
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);

// Priority mechanism — "lanes"
// - SyncLane (highest): discrete user events (click, keypress)
// - InputContinuousLane: drag, scroll, hover
// - DefaultLane: network responses, timeouts
// - TransitionLane (lowest): startTransition updates

// startTransition — marks state update as low priority
const [isPending, startTransition] = useTransition();

const handleSearch = (query) => {
  setInput(query);                    // ← SyncLane — immediate
  startTransition(() => {
    setSearchResults(search(query));  // ← TransitionLane — can be interrupted
  });
};

// React 18 concurrent features:
// useTransition, useDeferredValue — defer non-urgent renders
// Automatic batching — all updates batched (even in async)
// Suspense improvements — works with data fetching (via React Query etc.)
```

---

### Q3. What are Error Boundaries and how do they work?

```jsx
// Error Boundaries MUST be class components (no functional equivalent yet)
// They catch JS errors in child component tree during render, lifecycle, constructors

class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  // Called when a child throws — update state to show fallback
  static getDerivedStateFromError(error) {
    return { hasError: true, error };
  }

  // Called after error — use for logging/monitoring
  componentDidCatch(error, info) {
    console.error('ErrorBoundary caught:', error);
    logToSentry(error, info.componentStack);
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback || <h2>Something went wrong</h2>;
    }
    return this.props.children;
  }
}

// Usage
<ErrorBoundary fallback={<ErrorPage />}>
  <PaymentForm />
</ErrorBoundary>

// ⚠️ Error boundaries do NOT catch:
// - Async errors (fetch failures → handle in useEffect)
// - Event handler errors (handle with try/catch)
// - Server-side rendering errors
// - Errors in the boundary itself

// Modern alternative: react-error-boundary library
import { ErrorBoundary } from 'react-error-boundary';
<ErrorBoundary
  FallbackComponent={ErrorFallback}
  onError={(error, info) => logToSentry(error, info)}
  onReset={() => queryClient.clear()} // reset state on retry
>
  <App />
</ErrorBoundary>
```

---

### Q4. What are React Portals?

```jsx
import { createPortal } from 'react-dom';

// Portals render children into a DIFFERENT DOM node than the parent
// Use case: modals, tooltips, dropdowns — escape CSS overflow/z-index constraints

function Modal({ isOpen, onClose, children }) {
  if (!isOpen) return null;

  // Renders into document.body even though Modal is nested deep in the tree
  return createPortal(
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <button onClick={onClose}>✕</button>
        {children}
      </div>
    </div>,
    document.body  // ← target DOM node (outside React root)
  );
}

// Key insight: React events still bubble through the React tree (not DOM tree)
// So clicking in the modal still triggers React event handlers on parent components

// Usage
function App() {
  const [showModal, setShowModal] = useState(false);
  return (
    <div style={{ overflow: 'hidden' }}>  {/* Modal escapes this! */}
      <button onClick={() => setShowModal(true)}>Open</button>
      <Modal isOpen={showModal} onClose={() => setShowModal(false)}>
        <h2>Modal Content</h2>
      </Modal>
    </div>
  );
}
```

---

### Q5. How does `forwardRef` and `useImperativeHandle` work?

```jsx
import { forwardRef, useImperativeHandle, useRef } from 'react';

// forwardRef: pass ref from parent to child DOM element or component
const FancyInput = forwardRef(function FancyInput(props, ref) {
  return <input ref={ref} className="fancy" {...props} />;
});

// Parent can now control the input
function Parent() {
  const inputRef = useRef(null);
  return (
    <>
      <FancyInput ref={inputRef} placeholder="Type here" />
      <button onClick={() => inputRef.current.focus()}>Focus Input</button>
    </>
  );
}

// useImperativeHandle: SELECTIVE exposure of methods (don't expose full DOM node)
const VideoPlayer = forwardRef(function VideoPlayer(props, ref) {
  const videoRef = useRef(null);

  useImperativeHandle(ref, () => ({
    play:  () => videoRef.current.play(),
    pause: () => videoRef.current.pause(),
    seek:  (time) => { videoRef.current.currentTime = time; },
    // NOT exposing the raw videoRef.current — only specific methods
  }));

  return <video ref={videoRef} src={props.src} />;
});

// Usage
const playerRef = useRef(null);
<VideoPlayer ref={playerRef} src="/movie.mp4" />
<button onClick={() => playerRef.current.play()}>Play</button>
```

---

### Q6. What is Suspense and how does it work with data fetching?

```jsx
// Suspense: declarative loading states — show fallback until data is ready
import { Suspense, lazy } from 'react';

// 1. Code splitting (stable API)
const Dashboard = lazy(() => import('./Dashboard'));
<Suspense fallback={<Spinner />}>
  <Dashboard />
</Suspense>

// 2. Data fetching with Suspense (React 18+ with compatible libraries)
// Libraries must implement the "Suspense protocol" (throw a Promise during render)

// React Query (TanStack Query) with Suspense
const { data } = useSuspenseQuery({
  queryKey: ['user', userId],
  queryFn: () => fetchUser(userId)
}); // No need for isLoading check — Suspense handles it!

// In your component tree:
<ErrorBoundary fallback={<ErrorFallback />}>
  <Suspense fallback={<UserSkeleton />}>
    <UserProfile userId={id} />
  </Suspense>
</ErrorBoundary>

// Nested Suspense for granular loading
<Suspense fallback={<PageShell />}>
  <Suspense fallback={<ProfileSkeleton />}>
    <Profile />
  </Suspense>
  <Suspense fallback={<PostsSkeleton />}>
    <Posts />                           {/* loads independently */}
  </Suspense>
</Suspense>
```

---

### Q7. What are React Server Components (RSC)?

```jsx
// React Server Components — run on the SERVER only (Next.js App Router)
// They can: access DB directly, read files, use secrets
// They CANNOT: use useState, useEffect, event handlers, browser APIs

// app/page.tsx (Server Component — default in Next.js App Router)
async function ProductList() {
  // Direct DB query — no API layer needed!
  const products = await db.query('SELECT * FROM products');

  return (
    <ul>
      {products.map(p => (
        <li key={p.id}>
          {p.name}
          <AddToCartButton productId={p.id} /> {/* Client Component */}
        </li>
      ))}
    </ul>
  );
}

// 'use client' — Client Component (runs in browser)
'use client';
function AddToCartButton({ productId }) {
  const [added, setAdded] = useState(false); // ✅ can use state
  return (
    <button onClick={() => setAdded(true)}>
      {added ? 'Added!' : 'Add to Cart'}
    </button>
  );
}

// Key benefits: zero JS bundle for server components, DB access, caching
// Boundary: 'use client' marks the client/server boundary
```
