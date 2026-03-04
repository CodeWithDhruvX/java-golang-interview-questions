# ⚡ 07 — React & Frontend Frameworks Deep Dive
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- React reconciliation and virtual DOM
- Fiber architecture and concurrent rendering
- Hooks internals (useState, useEffect, useMemo, useCallback, useRef)
- Custom hooks patterns
- React 18: Suspense, transitions, automatic batching
- State management with Context, Zustand, Redux Toolkit

---

## ❓ Most Asked Questions

### Q1. How does React's reconciliation and Fiber work?

```javascript
// Virtual DOM: JavaScript representation of the real DOM
// Reconciliation: algorithm to diff old and new VDOMs

// Fiber: React's incremental rendering architecture (React 16+)
// Each component → Fiber node (unit of work that can pause/resume)

// ❌ Before Fiber (React 15 Stack): synchronous, blocking
// Large updates → blocked main thread → janky UI

// ✅ Fiber: cooperative multitasking
// Phase 1: Render (Reconciliation) — interruptible, can pause for high-priority work
// Phase 2: Commit — synchronous, applies DOM changes

// Reconciliation rules:
// 1. Different element type → unmount old, mount new (never reuse)
// 2. Same element type with different props → update props (reuse DOM node)
// 3. Keys: help React identify which list items changed/moved

// Bad: no keys (React reorders by index — wrong!)
function BadList({ items }) {
    return items.map((item, index) => (
        <Item key={index} data={item} /> // ❌ index as key causes bugs on reorder
    ));
}

// Good: stable unique keys
function GoodList({ items }) {
    return items.map(item => (
        <Item key={item.id} data={item} /> // ✅ stable identity
    ));
}

// React 18 Concurrent Features
// Transitions: mark updates as non-urgent (can be interrupted)
import { useTransition, startTransition } from 'react';

function SearchPage() {
    const [query, setQuery] = useState('');
    const [results, setResults] = useState([]);
    const [isPending, startTransition] = useTransition();

    const handleSearch = (e) => {
        setQuery(e.target.value);          // urgent: reflects input immediately
        startTransition(() => {
            setResults(search(e.target.value)); // non-urgent: can defer
        });
    };

    return (
        <>
            <input value={query} onChange={handleSearch} />
            {isPending ? <Spinner /> : <Results items={results} />}
        </>
    );
}
```

---

### Q2. Explain useState and useReducer internals.

```javascript
// useState: stores value in Fiber's memoized state queue
// Re-renders happen when state reference changes (Object.is comparison)

// ❌ Common mistake: mutating state
function Counter() {
    const [items, setItems] = useState([]);

    const addItem = () => {
        items.push("new"); // ❌ mutates — React doesn't see change!
        setItems(items);   // same reference → no re-render
    };

    const addItemFixed = () => {
        setItems(prev => [...prev, "new"]); // ✅ new array reference
    };
}

// useState lazy initialization (for expensive initial values)
const [state, setState] = useState(() => {
    // This function only runs once — not on every render
    return parseStoredData(localStorage.getItem("state")) || defaultState;
});

// Batching in React 18: multiple setStates in one event → one re-render
function App() {
    const [count, setCount] = useState(0);
    const [text, setText] = useState("");

    const handleClick = () => {
        setCount(c => c + 1); // React 18: batched
        setText("updated");   // React 18: batched (even in async!)
        // Only ONE re-render (React 18 automatic batching)
    };
}

// useReducer: for complex state logic
const initialState = { count: 0, step: 1, history: [] };

function reducer(state, action) {
    switch (action.type) {
        case 'INCREMENT':
            return {
                ...state,
                count: state.count + state.step,
                history: [...state.history, state.count + state.step]
            };
        case 'SET_STEP':
            return { ...state, step: action.payload };
        case 'UNDO':
            const history = state.history.slice(0, -1);
            return { ...state, count: history[history.length - 1] ?? 0, history };
        default:
            return state;
    }
}

function CounterWithHistory() {
    const [state, dispatch] = useReducer(reducer, initialState);

    return (
        <>
            <span>{state.count}</span>
            <button onClick={() => dispatch({ type: 'INCREMENT' })}>+</button>
            <button onClick={() => dispatch({ type: 'UNDO' })}>Undo</button>
        </>
    );
}
```

---

### Q3. Explain useEffect — dependencies, cleanup, timing.

```javascript
// useEffect: synchronize component with external systems
// Runs AFTER paint (non-blocking, async to layout)
// useLayoutEffect: runs BEFORE paint (blocking, like componentDidMount)

// Cleanup function: runs before next effect + on unmount
function DataFetcher({ userId }) {
    const [data, setData] = useState(null);

    useEffect(() => {
        let cancelled = false;
        const controller = new AbortController();

        async function fetchData() {
            try {
                const res = await fetch(`/api/users/${userId}`, {
                    signal: controller.signal
                });
                const json = await res.json();
                if (!cancelled) setData(json); // avoid state update on unmounted component
            } catch (err) {
                if (err.name !== 'AbortError') console.error(err);
            }
        }

        fetchData();

        // Cleanup: cancel request when userId changes or component unmounts
        return () => {
            cancelled = true;
            controller.abort();
        };
    }, [userId]); // re-run when userId changes

    return data ? <UserProfile user={data} /> : <Spinner />;
}

// Dependency array rules:
// []          : run once after mount (componentDidMount)
// [a, b]      : run when a or b changes
// (no array)  : run after every render (usually wrong)
// ⚠️ Exhaustive deps: include ALL values used inside effect

// Custom hook: encapsulate effect logic
function useFetch(url) {
    const [state, setState] = useState({ data: null, loading: true, error: null });

    useEffect(() => {
        setState({ data: null, loading: true, error: null });
        let cancelled = false;

        fetch(url)
            .then(r => r.json())
            .then(data => !cancelled && setState({ data, loading: false, error: null }))
            .catch(error => !cancelled && setState({ data: null, loading: false, error }));

        return () => { cancelled = true; };
    }, [url]);

    return state;
}
```

---

### Q4. Performance optimization with useMemo, useCallback, memo.

```javascript
// React.memo: memoize component — skip re-render if props unchanged
const ExpensiveComponent = React.memo(({ data, onAction }) => {
    console.log("ExpensiveComponent rendered");
    return <div onClick={onAction}>{data.value}</div>;
}, (prevProps, nextProps) => {
    // Custom comparison: return true if equal (skip re-render)
    return prevProps.data.id === nextProps.data.id;
});

// useMemo: memoize expensive computed value
function DataGrid({ items, sortKey }) {
    // ❌ Without useMemo: sortedItems recomputed on EVERY render
    // const sortedItems = items.slice().sort((a, b) => a[sortKey] > b[sortKey] ? 1 : -1);

    // ✅ With useMemo: only recomputed when items or sortKey changes
    const sortedItems = useMemo(
        () => items.slice().sort((a, b) => a[sortKey] > b[sortKey] ? 1 : -1),
        [items, sortKey]
    );

    return sortedItems.map(item => <Row key={item.id} item={item} />);
}

// useCallback: stable function reference (for memo'd children or effect deps)
function Parent() {
    const [count, setCount] = useState(0);

    // ❌ Without useCallback: new function on every render → Child always re-renders
    // const handleClick = () => doSomething(count);

    // ✅ With useCallback: same ref when count hasn't changed
    const handleClick = useCallback(() => {
        doSomething(count);
    }, [count]);

    return <MemoizedChild onClick={handleClick} />;
}

// ⚠️ Don't over-optimize: useMemo/useCallback have overhead!
// Use only when: expensive computations, referential equality issues with memo'd children

// useRef: stable reference that doesn't trigger re-render
function Timer() {
    const [time, setTime] = useState(0);
    const intervalRef = useRef(null); // persists across renders, no re-render

    const start = () => {
        if (intervalRef.current) return; // already running
        intervalRef.current = setInterval(() => setTime(t => t + 1), 1000);
    };

    const stop = () => {
        clearInterval(intervalRef.current);
        intervalRef.current = null;
    };

    useEffect(() => () => clearInterval(intervalRef.current), []); // cleanup

    return <div><button onClick={start}>Start</button> {time}s</div>;
}
```

---

### Q5. React 18 — Suspense, transitions, and Concurrent Mode.

```javascript
// Suspense: show fallback while component loads (data or code)
// React 18: works with data fetching (via use() hook or libraries)

// Code splitting with lazy + Suspense
import { lazy, Suspense } from 'react';

const Dashboard = lazy(() => import('./Dashboard')); // dynamic import
const Settings  = lazy(() => import('./Settings'));

function App() {
    return (
        <Suspense fallback={<LoadingSpinner />}>
            <Router>
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/settings"  element={<Settings />} />
            </Router>
        </Suspense>
    );
}

// Data fetching with Suspense (React 18 'use' hook)
import { use } from 'react';

// Resource that suspends until data is ready
function UserProfile({ userId }) {
    // 'use' can throw a Promise (triggers Suspense) or an error (triggers ErrorBoundary)
    const user = use(fetchUser(userId));
    return <div>{user.name}</div>;
}

function UserPage({ userId }) {
    return (
        <ErrorBoundary fallback={<ErrorMessage />}>
            <Suspense fallback={<UserSkeleton />}>
                <UserProfile userId={userId} />
            </Suspense>
        </ErrorBoundary>
    );
}

// Automatic batching in React 18
// Before React 18: only batched in event handlers
// React 18: batches everywhere (setTimeout, Promises, native events)
setTimeout(() => {
    setCount(c => c + 1);   // React 18: batched!
    setFlag(f => !f);        // React 18: batched! One re-render total.
}, 1000);

// Opt out of batching: flushSync (for reading DOM synchronously)
import { flushSync } from 'react-dom';

function handleClick() {
    flushSync(() => setCount(c => c + 1)); // force synchronous re-render
    // DOM is now updated — can measure layout
    const height = element.offsetHeight;
    flushSync(() => setHeight(height));
}
```

---

### Q6. Build a custom React hook — usePrevious, useLocalStorage, useIntersection.

```javascript
// usePrevious: capture previous value before re-render
function usePrevious(value) {
    const ref = useRef(undefined);
    useEffect(() => { ref.current = value; }); // runs AFTER render
    return ref.current; // returns value from BEFORE current render
}

function PriceTracker({ price }) {
    const prevPrice = usePrevious(price);
    const change = price - (prevPrice ?? price);
    return (
        <div className={change > 0 ? "up" : change < 0 ? "down" : "flat"}>
            ${price} ({change > 0 ? "+" : ""}{change})
        </div>
    );
}

// useLocalStorage: synchronized with localStorage
function useLocalStorage(key, initialValue) {
    const [value, setValue] = useState(() => {
        try {
            const item = window.localStorage.getItem(key);
            return item ? JSON.parse(item) : initialValue;
        } catch {
            return initialValue;
        }
    });

    const setStoredValue = useCallback((valueOrFn) => {
        setValue(prev => {
            const next = typeof valueOrFn === 'function' ? valueOrFn(prev) : valueOrFn;
            try {
                window.localStorage.setItem(key, JSON.stringify(next));
            } catch (e) {
                console.error("localStorage write failed:", e);
            }
            return next;
        });
    }, [key]);

    // Sync across tabs
    useEffect(() => {
        const handler = (e) => {
            if (e.key === key && e.newValue !== null) {
                setValue(JSON.parse(e.newValue));
            }
        };
        window.addEventListener('storage', handler);
        return () => window.removeEventListener('storage', handler);
    }, [key]);

    return [value, setStoredValue];
}

// useIntersectionObserver: lazy load, infinite scroll, analytics
function useIntersectionObserver(options = {}) {
    const [entry, setEntry] = useState(null);
    const [node, setNode] = useState(null);

    useEffect(() => {
        if (!node) return;
        const observer = new IntersectionObserver(
            ([entry]) => setEntry(entry),
            { threshold: 0.1, ...options }
        );
        observer.observe(node);
        return () => observer.disconnect();
    }, [node, options.threshold, options.rootMargin]);

    return [setNode, entry?.isIntersecting ?? false, entry];
}

// Usage
function LazySection({ children }) {
    const [ref, isVisible] = useIntersectionObserver();
    return (
        <div ref={ref}>
            {isVisible ? children : <Skeleton />}
        </div>
    );
}
```

---

### Q7. What is the Context API and when to prefer Zustand/Redux?

```javascript
// Context: built-in React state sharing (no external library)
// ⚠️ Re-renders ALL consumers when context value changes

// ✅ Context is good for: theme, locale, auth (low-frequency updates)
// ❌ Context is bad for: high-frequency updates (every keystroke → all consumers re-render)

// Split contexts to avoid unnecessary re-renders
const ThemeContext  = createContext({ theme: 'light' });
const AuthContext   = createContext({ user: null });
const CartContext   = createContext({ items: [], total: 0 });

// Context optimization: memoize the value
function CartProvider({ children }) {
    const [items, setItems] = useState([]);
    const total = useMemo(
        () => items.reduce((sum, item) => sum + item.price, 0),
        [items]
    );

    const contextValue = useMemo(
        () => ({ items, total, addItem: (item) => setItems(prev => [...prev, item]) }),
        [items, total]
    );

    return <CartContext.Provider value={contextValue}>{children}</CartContext.Provider>;
}

// Zustand: simpler external state (no boilerplate, granular subscriptions)
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const useStore = create(persist(
    (set, get) => ({
        cart: [],
        addItem: item => set(state => ({ cart: [...state.cart, item] })),
        removeItem: id => set(state => ({ cart: state.cart.filter(i => i.id !== id) })),
        total: () => get().cart.reduce((sum, item) => sum + item.price, 0),
    }),
    { name: 'cart-storage' } // persisted to localStorage automatically
));

// Component: only re-renders when subscribed slice changes
function CartBadge() {
    const count = useStore(state => state.cart.length); // granular subscription!
    return <span>{count}</span>;
}

// Choose: Context for simple/infrequent; Zustand for co-located/medium; Redux for large teams
```
