# 🪝 03 — Hooks & State
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- useState — state updates and functional updates
- useEffect — dependency array, cleanup, common patterns
- useRef — DOM access and mutable values
- useContext — consuming context
- useReducer — complex state logic
- Custom hooks — extracting reusable logic
- Rules of Hooks

---

## ❓ Most Asked Questions

### Q1. What is `useState` and how does it work?

```jsx
import { useState } from 'react';

function Counter() {
  // [currentValue, setterFunction] = useState(initialValue)
  const [count, setCount] = useState(0);

  // Direct update
  const increment = () => setCount(count + 1);

  // Functional update — use when new state depends on previous
  const incrementSafe = () => setCount(prev => prev + 1);
  // ✅ Use functional form when clicking fast (avoids stale closure issues)

  // Object state — must spread to update
  const [user, setUser] = useState({ name: '', email: '' });
  const updateName = (name) => setUser(prev => ({ ...prev, name }));

  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={incrementSafe}>+</button>
    </div>
  );
}

// Lazy initialization — for expensive initial state computation
const [data, setData] = useState(() => {
  return JSON.parse(localStorage.getItem('saved-data')) || [];
  // ↑ Function only called ONCE on mount, not every render
});
```

---

### Q2. How does `useEffect` work and what is the dependency array?

```jsx
import { useEffect } from 'react';

function Examples({ userId }) {
  // 1. Run ONCE on mount (empty deps)
  useEffect(() => {
    document.title = 'My Page';
  }, []);

  // 2. Run EVERY render (no deps array)
  useEffect(() => {
    console.log('Rendered');
  });

  // 3. Run when dependency changes
  useEffect(() => {
    fetchUser(userId);
  }, [userId]); // re-runs whenever userId changes

  // 4. Cleanup function — runs before re-run and on unmount
  useEffect(() => {
    const subscription = subscribeToUpdates(userId);

    return () => {
      subscription.unsubscribe(); // cleanup
    };
  }, [userId]);

  // 5. Async in useEffect — use IIFE or inner async function
  useEffect(() => {
    async function loadData() {
      const data = await fetchData(userId);
      setData(data);
    }
    loadData();
  }, [userId]);
  // ❌ Don't make the effect callback async directly
}
```

---

### Q3. What are common `useEffect` mistakes?

```jsx
// ❌ Mistake 1: Missing dependencies → stale closure
const [count, setCount] = useState(0);
useEffect(() => {
  const interval = setInterval(() => {
    setCount(count + 1); // stale — always uses initial count = 0
  }, 1000);
  return () => clearInterval(interval);
}, []); // ← count missing from deps!

// ✅ Fix: use functional update
useEffect(() => {
  const interval = setInterval(() => {
    setCount(prev => prev + 1); // always fresh
  }, 1000);
  return () => clearInterval(interval);
}, []);

// ❌ Mistake 2: Infinite loop — updating state in effect with that state as dep
const [data, setData] = useState([]);
useEffect(() => {
  setData([...data, newItem]); // triggers re-render → effect → re-render...
}, [data]); // ← infinite loop!

// ✅ Fix: remove data from deps, use functional update
useEffect(() => {
  setData(prev => [...prev, newItem]);
}, []); // or proper deps

// ❌ Mistake 3: Using object/array as dependency
const options = { id: userId }; // new reference every render!
useEffect(() => { fetch(options); }, [options]); // runs every render

// ✅ Fix: use primitives or useMemo
useEffect(() => { fetch({ id: userId }); }, [userId]);
```

---

### Q4. What is `useRef` and when do you use it?

```jsx
import { useRef } from 'react';

// Use case 1: Accessing DOM elements
function FocusInput() {
  const inputRef = useRef(null);

  const focusInput = () => inputRef.current.focus();

  return (
    <>
      <input ref={inputRef} placeholder="Type here" />
      <button onClick={focusInput}>Focus</button>
    </>
  );
}

// Use case 2: Storing mutable value without triggering re-render
function Timer() {
  const [seconds, setSeconds] = useState(0);
  const intervalRef = useRef(null); // persists across renders, no re-render

  const start = () => {
    intervalRef.current = setInterval(() => {
      setSeconds(s => s + 1);
    }, 1000);
  };

  const stop = () => clearInterval(intervalRef.current);

  return (
    <div>
      <p>{seconds}s</p>
      <button onClick={start}>Start</button>
      <button onClick={stop}>Stop</button>
    </div>
  );
}

// Use case 3: Tracking previous value
function usePrevious(value) {
  const prevRef = useRef(undefined);
  useEffect(() => { prevRef.current = value; });
  return prevRef.current; // previous value
}
```

---

### Q5. What is `useReducer` and when should you use it over `useState`?

```jsx
import { useReducer } from 'react';

// useReducer: better when state logic is complex or state values are interdependent
type State = { count: number; history: number[] };
type Action = { type: 'increment' | 'decrement' | 'reset' };

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'increment':
      return { count: state.count + 1, history: [...state.history, state.count + 1] };
    case 'decrement':
      return { count: state.count - 1, history: [...state.history, state.count - 1] };
    case 'reset':
      return { count: 0, history: [] };
    default:
      return state;
  }
}

function Counter() {
  const [state, dispatch] = useReducer(reducer, { count: 0, history: [] });

  return (
    <div>
      <p>Count: {state.count}</p>
      <p>History: {state.history.join(', ')}</p>
      <button onClick={() => dispatch({ type: 'increment' })}>+</button>
      <button onClick={() => dispatch({ type: 'decrement' })}>-</button>
      <button onClick={() => dispatch({ type: 'reset' })}>Reset</button>
    </div>
  );
}

// Use useState when: simple independent values, 1-3 state variables
// Use useReducer when: multiple inter-related state, complex transitions, testability
```

---

### Q6. What are the Rules of Hooks?

```jsx
// Rule 1: Only call Hooks at the top level
// ❌ Don't call inside conditions, loops, or nested functions
function BadComponent({ isAdmin }) {
  if (isAdmin) {
    const [data, setData] = useState([]); // ❌ conditional hook!
  }
}

// ✅ Always call at top level, then conditionally use
function GoodComponent({ isAdmin }) {
  const [data, setData] = useState([]); // ✅ always called

  if (!isAdmin) return <NotAuthorized />;
  return <AdminPanel data={data} />;
}

// Rule 2: Only call Hooks from React functions
// ✅ Functional components
// ✅ Custom hooks (functions starting with 'use')
// ❌ Regular JS functions
// ❌ Class components

// Why these rules?
// React relies on the ORDER and COUNT of hook calls being stable
// across renders to maintain state correctly
// ESLint: eslint-plugin-react-hooks enforces these automatically
```

---

### Q7. What is a custom hook?

```jsx
// Custom hook = reusable function that uses built-in hooks
// MUST start with 'use' (convention + ESLint enforcement)

// Example: useLocalStorage
function useLocalStorage(key, initialValue) {
  const [storedValue, setStoredValue] = useState(() => {
    try {
      const item = localStorage.getItem(key);
      return item ? JSON.parse(item) : initialValue;
    } catch { return initialValue; }
  });

  const setValue = (value) => {
    try {
      const valueToStore = value instanceof Function ? value(storedValue) : value;
      setStoredValue(valueToStore);
      localStorage.setItem(key, JSON.stringify(valueToStore));
    } catch (e) { console.error(e); }
  };

  return [storedValue, setValue];
}

// Example: useFetch
function useFetch(url) {
  const [data, setData]     = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError]   = useState(null);

  useEffect(() => {
    let cancelled = false;
    setLoading(true);
    fetch(url)
      .then(r => r.json())
      .then(d => { if (!cancelled) { setData(d); setLoading(false); } })
      .catch(e => { if (!cancelled) { setError(e); setLoading(false); } });
    return () => { cancelled = true; };
  }, [url]);

  return { data, loading, error };
}

// Usage
const { data: users, loading } = useFetch('/api/users');
```

---

### Q8. What is `useContext` and how do you use it?

```jsx
import { createContext, useContext, useState } from 'react';

// 1. Create context with default value
const ThemeContext = createContext('light');

// 2. Custom hook for easier consumption
const useTheme = () => {
  const ctx = useContext(ThemeContext);
  if (!ctx) throw new Error('useTheme must be inside ThemeProvider');
  return ctx;
};

// 3. Provider component
function ThemeProvider({ children }) {
  const [theme, setTheme] = useState<'light' | 'dark'>('light');
  const toggle = () => setTheme(t => t === 'light' ? 'dark' : 'light');

  return (
    <ThemeContext.Provider value={{ theme, toggle }}>
      {children}
    </ThemeContext.Provider>
  );
}

// 4. Consumer — any component deep in the tree
function ThemeToggleButton() {
  const { theme, toggle } = useTheme();
  return <button onClick={toggle}>Current: {theme}</button>;
}

// 5. Wrap app
function App() {
  return (
    <ThemeProvider>
      <Header />   {/* ThemeToggleButton is somewhere inside */}
      <Main />
    </ThemeProvider>
  );
}
```
