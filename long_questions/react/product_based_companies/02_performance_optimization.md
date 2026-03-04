# ⚡ 02 — Performance Optimization in React
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

> Meta, Google, Airbnb, Netflix demand you understand not just the APIs but *when and why* to use them.

---

## 🔑 Must-Know Topics
- React.memo and when to use it
- useMemo vs useCallback — differences and pitfalls
- Code splitting with lazy() + Suspense
- Virtualization (react-window / react-virtual)
- Web Vitals and React's impact on LCP, CLS, FID
- Profiler API and React DevTools Profiler
- Avoiding unnecessary re-renders

---

## ❓ Most Asked Questions

### Q1. When should you use `React.memo`?

```jsx
// React.memo — prevents re-render if props haven't changed (shallow equal)
const ExpensiveList = React.memo(function ExpensiveList({ items, onSelect }) {
  console.log('ExpensiveList rendered');
  return (
    <ul>
      {items.map(item => (
        <li key={item.id} onClick={() => onSelect(item.id)}>
          {item.name}
        </li>
      ))}
    </ul>
  );
});

// ⚠️ Common mistake: parent passes new object/function on every render
function Parent() {
  const [count, setCount] = useState(0);

  // ❌ New function reference every render — memo is useless!
  const handleSelect = (id) => console.log('selected', id);

  // ✅ Stable reference with useCallback
  const handleSelect = useCallback((id) => console.log('selected', id), []);

  return <ExpensiveList items={items} onSelect={handleSelect} />;
}

// Custom comparison function for complex props
const MemoComponent = React.memo(MyComponent, (prevProps, nextProps) => {
  return prevProps.userId === nextProps.userId; // return true = skip render
});
```

---

### Q2. What is the difference between `useMemo` and `useCallback`?

```jsx
// useMemo — memoizes a COMPUTED VALUE
const sortedList = useMemo(() => {
  return [...items].sort((a, b) => a.name.localeCompare(b.name));
}, [items]); // only recomputes when items changes

// useCallback — memoizes a FUNCTION REFERENCE
const handleDelete = useCallback((id) => {
  setItems(prev => prev.filter(item => item.id !== id));
}, []); // stable reference — dependency array is empty

// Key distinction:
// useMemo(fn, deps)     → returns fn()   (the result)
// useCallback(fn, deps) → returns fn     (the function itself)

// ⚠️ DON'T over-memoize — it costs memory and has overhead
// Use useMemo when: computation is genuinely expensive (sorting large arrays,
//   complex calculations, building derived data structures)
// Use useCallback when: passing callbacks to memoized children or into
//   useEffect dependencies

// Profile FIRST, optimize SECOND
```

---

### Q3. How does code splitting work in React?

```jsx
// React.lazy + Suspense for route-based code splitting
import React, { lazy, Suspense } from 'react';

// These create separate chunks — loaded only when needed
const Dashboard = lazy(() => import('./pages/Dashboard'));
const Settings  = lazy(() => import('./pages/Settings'));
const Analytics = lazy(() => import('./pages/Analytics'));

function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Routes>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/settings"  element={<Settings />} />
        <Route path="/analytics" element={<Analytics />} />
      </Routes>
    </Suspense>
  );
}

// Component-level splitting (heavy components)
const HeavyEditor = lazy(() => import('./components/RichTextEditor'));

function Page() {
  const [showEditor, setShowEditor] = useState(false);
  return (
    <>
      <button onClick={() => setShowEditor(true)}>Open Editor</button>
      {showEditor && (
        <Suspense fallback={<Skeleton />}>
          <HeavyEditor />
        </Suspense>
      )}
    </>
  );
}

// Named exports need a wrapper:
const { DataGrid } = lazy(() =>
  import('./components').then(m => ({ default: m.DataGrid }))
);
```

---

### Q4. What is list virtualization and when do you need it?

```jsx
// Problem: rendering 10,000 list items = 10,000 DOM nodes = slow!
// Solution: virtualization — only render items visible in the viewport

import { FixedSizeList } from 'react-window';

const Row = ({ index, style }) => (
  <div style={style}>       {/* style is REQUIRED — sets position/size */}
    Row #{index}: {data[index].name}
  </div>
);

function VirtualList() {
  return (
    <FixedSizeList
      height={600}           // viewport height
      width="100%"
      itemCount={10000}      // total items
      itemSize={50}          // each row height in px
    >
      {Row}
    </FixedSizeList>
  );
}

// For variable-height rows: use VariableSizeList
// For grids: use FixedSizeGrid / VariableSizeGrid
// Modern alternative: @tanstack/react-virtual (more flexible)

// When to virtualize:
// ✅ Lists > 100 items that are scrollable
// ✅ Complex row content (images, sub-components)
// ❌ Small lists (< 50 items) — overhead not worth it
```

---

### Q5. How do you profile React performance?

```jsx
// 1. React DevTools Profiler — record renders, see timings
// Open DevTools → Profiler tab → Record → interact → Stop

// 2. Profiler API — programmatic profiling
import { Profiler } from 'react';

function onRender(id, phase, actualDuration, baseDuration) {
  // id: component name
  // phase: 'mount' or 'update'
  // actualDuration: time spent rendering (ms) — after memoization
  // baseDuration: estimated time without memoization
  console.log({ id, phase, actualDuration });
}

<Profiler id="Dashboard" onRender={onRender}>
  <Dashboard />
</Profiler>

// 3. React.StrictMode highlights unexpected side effects

// 4. Chrome Performance tab — record, analyze flame graph
// Look for: long tasks > 50ms, layout thrashing, forced reflows

// 5. Why Did You Render library — logs unnecessary re-renders
import whyDidYouRender from '@welldone-software/why-did-you-render';
whyDidYouRender(React, { trackAllPureComponents: true });
```

---

### Q6. How do you prevent prop drilling with performance in mind?

```jsx
// Context is convenient but causes ALL consumers to re-render on any change
// Solution: split context by concern + useMemo for value

// ❌ One big context — any state change re-renders all consumers
const AppContext = createContext();
const value = { user, theme, cart, notifications }; // everything!

// ✅ Split into separate contexts
const UserContext = createContext();
const ThemeContext = createContext();

// ✅ Memoize context value to prevent extra renders
function UserProvider({ children }) {
  const [user, setUser] = useState(null);

  const value = useMemo(() => ({ user, setUser }), [user]);
  //            ↑ only new object when user actually changes

  return <UserContext.Provider value={value}>{children}</UserContext.Provider>;
}

// ✅ Even better for high-frequency: use Zustand or Jotai
// They use subscription model — only subscribed components re-render
import { create } from 'zustand';
const useStore = create(set => ({
  count: 0,
  increment: () => set(state => ({ count: state.count + 1 }))
}));
```

---

### Q7. What are Web Vitals and how does React affect them?

| Metric | What It Measures | React Impact |
|--------|-----------------|--------------|
| **LCP** (Largest Contentful Paint) | Load time of main content | SSR/SSG reduce LCP; lazy loading hurts it |
| **CLS** (Cumulative Layout Shift) | Visual stability | Async image loading, dynamic content insertion |
| **FID/INP** (Interaction) | Responsiveness to user input | Heavy JS, long render tasks block main thread |
| **TTFB** | Server response time | SSR, caching |

```jsx
// Measure Web Vitals in React
import { onCLS, onFID, onLCP, onINP } from 'web-vitals';

function sendToAnalytics({ name, value }) {
  console.log(`${name}: ${value}`);
}

onCLS(sendToAnalytics);
onLCP(sendToAnalytics);
onINP(sendToAnalytics); // INP replaces FID in CWV 2024

// React optimizations for each:
// LCP: next/image, preload hints, SSR
// CLS: reserve space for images/ads (width + height attrs)
// INP: startTransition, defer non-critical renders, avoid long tasks
```

---

### Q8. How does `useTransition` help with performance?

```jsx
import { useState, useTransition, useDeferredValue } from 'react';

// useTransition — marks state update as non-urgent
function SearchPage() {
  const [input, setInput] = useState('');
  const [results, setResults] = useState([]);
  const [isPending, startTransition] = useTransition();

  const handleChange = (e) => {
    setInput(e.target.value);              // immediate — keeps input responsive

    startTransition(() => {
      setResults(search(e.target.value)); // deferred — can be interrupted
    });
  };

  return (
    <>
      <input value={input} onChange={handleChange} />
      {isPending ? <Spinner /> : <ResultsList results={results} />}
    </>
  );
}

// useDeferredValue — defers re-rendering with new value
function SearchResults({ query }) {
  const deferredQuery = useDeferredValue(query);
  // deferredQuery lags behind query during typing
  // component renders with old value until React has idle time

  const results = useMemo(() => expensiveSearch(deferredQuery), [deferredQuery]);
  return <List items={results} />;
}
```
