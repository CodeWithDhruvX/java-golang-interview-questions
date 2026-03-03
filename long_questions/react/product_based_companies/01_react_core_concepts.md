# ⚛️ 01 — React Core Concepts
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

> Companies like **Google, Meta, Atlassian** dive deep into how React works under the hood — not just how to use it.

---

## 🔑 Must-Know Topics
- Virtual DOM and reconciliation algorithm
- React Fiber architecture
- Keys and their role in diffing
- Controlled vs uncontrolled components
- Synthetic events and event delegation
- Component lifecycle (functional + class)
- Strict Mode and its behavior

---

## ❓ Most Asked Questions

### Q1. What is the Virtual DOM and how does reconciliation work?

**Answer:**
The Virtual DOM is a lightweight JavaScript representation of the real DOM. When state/props change, React:
1. Creates a new Virtual DOM tree
2. **Diffs** it against the previous tree (reconciliation)
3. Computes the minimal set of real DOM updates
4. **Commits** only those changes to the real DOM

```jsx
// React creates a virtual representation — not actual DOM nodes
const element = <h1>Hello, {name}</h1>;
// Equivalent to:
const element = React.createElement('h1', null, 'Hello, ', name);
// Returns: { type: 'h1', props: { children: ['Hello, ', 'Alice'] } }
```

> Key insight: **Reconciliation is O(n)** — React uses heuristics (same type = update, different type = unmount+mount, keys = reorder) to avoid O(n³) tree diffing.

---

### Q2. What is React Fiber and how does it differ from the old Stack Reconciler?

**Answer:**
Fiber is React's **reimplemented reconciliation engine** (introduced in React 16):

| Aspect | Old Stack Reconciler | React Fiber |
|--------|---------------------|-------------|
| Execution | Synchronous, can't interrupt | **Asynchronous, interruptible** |
| Rendering | All-or-nothing | **Incremental** (work units) |
| Priority | None | **Priority scheduling** |
| Concurrency | No | Yes (Concurrent Mode) |

```jsx
// Fiber enables priority-based rendering
// High priority: user input, animations
// Low priority: data fetching renders, background updates

// startTransition marks low-priority updates
import { startTransition } from 'react';

function handleInput(value) {
  setInputValue(value);            // high priority — immediate
  startTransition(() => {
    setSearchResults(search(value)); // low priority — can be interrupted
  });
}
```

---

### Q3. Why are keys important in React lists? What happens without them?

```jsx
// ❌ Bad — no keys, React re-creates all items on every render
const List = ({ items }) => (
  <ul>
    {items.map(item => <li>{item.name}</li>)}
  </ul>
);

// ❌ Bad — index as key (breaks when list reorders/filters)
{items.map((item, index) => <li key={index}>{item.name}</li>)}

// ✅ Good — stable, unique ID as key
{items.map(item => <li key={item.id}>{item.name}</li>)}

// Why it matters: without keys, React diffs by position.
// If you prepend an item, ALL items appear to have changed → O(n) DOM updates.
// With keys, React matches items by identity → only 1 DOM insert.
```

---

### Q4. What are controlled vs uncontrolled components?

```jsx
// --- CONTROLLED: React owns the state ---
function ControlledInput() {
  const [value, setValue] = useState('');

  return (
    <input
      value={value}                        // React controls the value
      onChange={e => setValue(e.target.value)}
    />
  );
}

// --- UNCONTROLLED: DOM owns the state ---
function UncontrolledInput() {
  const inputRef = useRef(null);

  const handleSubmit = () => {
    console.log(inputRef.current.value);   // read DOM directly
  };

  return <input ref={inputRef} defaultValue="initial" />;
}

// When to use uncontrolled: file inputs, integrating with non-React libs,
// performance-critical forms with many fields
```

---

### Q5. How does React's synthetic event system work?

```jsx
// React wraps native events in SyntheticEvent for cross-browser consistency
function Button() {
  const handleClick = (e) => {
    console.log(e.type);           // 'click'
    console.log(e.nativeEvent);    // actual browser Event
    e.preventDefault();            // works cross-browser
    e.stopPropagation();           // stop bubbling

    // ⚠️ Pre-React 17: events were pooled — access async = null
    // React 17+: event pooling removed, persists without .persist()
  };

  return <button onClick={handleClick}>Click</button>;
}

// React 17+: event delegation moved from document → root div
// This enables embedding React apps inside other React roots
```

---

### Q6. Explain the React component lifecycle (functional with hooks)

```jsx
function LifecycleDemoComponent({ userId }) {
  const [data, setData] = useState(null);

  // componentDidMount + componentDidUpdate
  useEffect(() => {
    let cancelled = false;

    async function fetchUser() {
      const res = await fetch(`/api/users/${userId}`);
      const json = await res.json();
      if (!cancelled) setData(json); // guard against stale closures
    }

    fetchUser();

    // componentWillUnmount (cleanup)
    return () => { cancelled = true; };
  }, [userId]); // re-runs when userId changes

  // Render phase — must be pure, no side effects here
  if (!data) return <p>Loading...</p>;
  return <div>{data.name}</div>;
}

// Lifecycle mapping:
// useState initializer        → constructor
// useEffect(() => {}, [])     → componentDidMount
// useEffect(() => {}, [dep])  → componentDidUpdate (when dep changes)
// useEffect cleanup           → componentWillUnmount
```

---

### Q7. What does React StrictMode do?

```jsx
// Wrap app in Strict Mode — development only, no production impact
<React.StrictMode>
  <App />
</React.StrictMode>

// StrictMode deliberately:
// 1. Double-invokes render functions/component bodies to catch side effects
// 2. Double-invokes useState initializers and useReducer functions
// 3. Double-invokes useEffect setup+cleanup to find missing cleanups
// 4. Warns about deprecated APIs (findDOMNode, legacy context, etc.)

// This is WHY your useEffect runs twice in development — it's intentional!
// If cleanup is correct, running twice should have no visible effect.

// Example issue it catches:
function BadComponent() {
  // This breaks under StrictMode double-invoke:
  let count = 0;
  return <div onClick={() => count++}>{count}</div>; // mutation in render!
}
```

---

### Q8. What is the difference between `React.createElement` and JSX?

```jsx
// JSX is syntactic sugar — Babel transforms it to createElement calls

// JSX:
const element = (
  <div className="container">
    <h1>Hello</h1>
    <p>World</p>
  </div>
);

// After Babel transform (React 17+ new JSX transform):
import { jsx as _jsx, jsxs as _jsxs } from 'react/jsx-runtime';
const element = _jsxs('div', {
  className: 'container',
  children: [
    _jsx('h1', { children: 'Hello' }),
    _jsx('p', { children: 'World' })
  ]
});

// Older transform (React 16 and below) needed: import React from 'react'
// React 17+ new transform — no import needed, automatic
```

---

### Q9. How does React handle batching of state updates?

```jsx
// React 18+: ALL updates are batched automatically (Automatic Batching)
// React 17 and below: only batched inside React event handlers

function Example() {
  const [count, setCount] = useState(0);
  const [flag, setFlag] = useState(false);

  // React 18: these are batched → 1 re-render
  const handleClick = () => {
    setCount(c => c + 1);
    setFlag(f => !f);
    // Triggers only ONE re-render
  };

  // React 18: even async batching works
  const handleAsync = async () => {
    await fetch('/api');
    setCount(c => c + 1);  // React 18: BATCHED
    setFlag(f => !f);       // React 17: NOT batched (2 renders)
  };

  // Opt out of batching:
  import { flushSync } from 'react-dom';
  flushSync(() => setCount(c => c + 1)); // forces immediate re-render
  flushSync(() => setFlag(f => !f));     // another immediate re-render
}
```

---

### Q10. What is the Context API and when should you NOT use it?

```jsx
// Create typed context
const ThemeContext = createContext<'light' | 'dark'>('light');

// Provider — wraps the component tree
function App() {
  const [theme, setTheme] = useState<'light' | 'dark'>('light');

  return (
    <ThemeContext.Provider value={theme}>
      <Toolbar />
    </ThemeContext.Provider>
  );
}

// Consumer — any nested component
function ThemedButton() {
  const theme = useContext(ThemeContext);
  return <button className={theme}>Click</button>;
}

// ⚠️ When NOT to use Context:
// 1. High-frequency updates (every keypress) — causes all consumers to re-render
// 2. Complex derived state — use Zustand/Recoil instead
// 3. Cross-cutting concerns in large apps — performance degrades

// ✅ Good uses: theme, locale/i18n, current user, auth state
```
