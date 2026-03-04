# 🏗️ 05 — Design Patterns in React
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Higher-Order Components (HOC)
- Render Props pattern
- Compound Components pattern
- Custom Hooks as reusable logic
- Container/Presentational pattern
- Headless components
- Provider pattern

---

## ❓ Most Asked Questions

### Q1. What is the Compound Components pattern?

```jsx
// Compound Components: components that work together sharing implicit state
// Classic example: Select/Option, Tabs/TabPanel, Accordion

// Implementation with Context
const TabsContext = createContext(null);

function Tabs({ children, defaultTab }) {
  const [activeTab, setActiveTab] = useState(defaultTab);
  return (
    <TabsContext.Provider value={{ activeTab, setActiveTab }}>
      <div className="tabs">{children}</div>
    </TabsContext.Provider>
  );
}

Tabs.List = function TabList({ children }) {
  return <div className="tab-list">{children}</div>;
};

Tabs.Tab = function Tab({ id, children }) {
  const { activeTab, setActiveTab } = useContext(TabsContext);
  return (
    <button
      className={`tab ${activeTab === id ? 'active' : ''}`}
      onClick={() => setActiveTab(id)}
    >
      {children}
    </button>
  );
};

Tabs.Panel = function TabPanel({ id, children }) {
  const { activeTab } = useContext(TabsContext);
  return activeTab === id ? <div className="panel">{children}</div> : null;
};

// Usage — beautiful, expressive API
<Tabs defaultTab="profile">
  <Tabs.List>
    <Tabs.Tab id="profile">Profile</Tabs.Tab>
    <Tabs.Tab id="settings">Settings</Tabs.Tab>
  </Tabs.List>
  <Tabs.Panel id="profile"><ProfileContent /></Tabs.Panel>
  <Tabs.Panel id="settings"><SettingsContent /></Tabs.Panel>
</Tabs>
```

---

### Q2. What is the Render Props pattern?

```jsx
// Render Props: sharing logic via a prop that is a function returning JSX
// (Largely superseded by custom hooks, but still appears in interviews)

// MouseTracker with render prop
function MouseTracker({ render }) {
  const [position, setPosition] = useState({ x: 0, y: 0 });

  const handleMouseMove = (e) => {
    setPosition({ x: e.clientX, y: e.clientY });
  };

  return (
    <div onMouseMove={handleMouseMove} style={{ height: '400px' }}>
      {render(position)}  {/* caller decides what to render with position */}
    </div>
  );
}

// Usage
<MouseTracker
  render={({ x, y }) => (
    <div>Mouse at: {x}, {y}</div>
  )}
/>

// Alternative: children as function (also render props)
<MouseTracker>
  {({ x, y }) => <div>Mouse: {x}, {y}</div>}
</MouseTracker>

// Modern equivalent — custom hook (cleaner)
function useMousePosition() {
  const [pos, setPos] = useState({ x: 0, y: 0 });
  useEffect(() => {
    const handler = (e) => setPos({ x: e.clientX, y: e.clientY });
    window.addEventListener('mousemove', handler);
    return () => window.removeEventListener('mousemove', handler);
  }, []);
  return pos;
}

// Usage of hook — much cleaner
function Tooltip() {
  const { x, y } = useMousePosition();
  return <div style={{ left: x, top: y }}>Tooltip</div>;
}
```

---

### Q3. What is the Provider pattern?

```jsx
// Provider pattern: wraps a subtree and provides data/logic via context
// Used for: auth, theme, notifications, feature flags, i18n

// Full featured provider with multiple concerns
const AuthContext = createContext(null);

function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  // Auth logic lives in provider — not scattered in components
  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged((u) => {
      setUser(u);
      setLoading(false);
    });
    return unsubscribe;
  }, []);

  const login = useCallback(async (email, password) => {
    const userCredential = await signInWithEmailAndPassword(auth, email, password);
    return userCredential.user;
  }, []);

  const logout = useCallback(() => signOut(auth), []);

  const value = useMemo(() => ({ user, loading, login, logout }), [user, loading, login, logout]);

  return (
    <AuthContext.Provider value={value}>
      {!loading && children}
    </AuthContext.Provider>
  );
}

// Custom hook for consumption
const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error('useAuth must be inside AuthProvider');
  return context;
};
```

---

### Q4. What is the Container/Presentational pattern?

```jsx
// CONTAINER: handles logic, data fetching, state — no styling
function UserListContainer() {
  const { data: users, isLoading, error } = useQuery({ queryKey: ['users'], queryFn: fetchUsers });
  const [searchQuery, setSearchQuery] = useState('');
  const [sortBy, setSortBy] = useState('name');

  const filteredUsers = useMemo(() =>
    users?.filter(u => u.name.toLowerCase().includes(searchQuery.toLowerCase()))
         .sort((a, b) => a[sortBy].localeCompare(b[sortBy]))
  , [users, searchQuery, sortBy]);

  if (isLoading) return <UserListSkeleton />;
  if (error)     return <ErrorMessage message={error.message} />;

  return (
    <UserListView
      users={filteredUsers}
      onSearch={setSearchQuery}
      onSortChange={setSortBy}
      sortBy={sortBy}
    />
  );
}

// PRESENTATIONAL: pure UI — easy to test, style, reuse, story
function UserListView({ users, onSearch, onSortChange, sortBy }) {
  return (
    <div>
      <input placeholder="Search..." onChange={e => onSearch(e.target.value)} />
      <select value={sortBy} onChange={e => onSortChange(e.target.value)}>
        <option value="name">Name</option>
        <option value="email">Email</option>
      </select>
      <ul>
        {users.map(user => (
          <li key={user.id}>{user.name} — {user.email}</li>
        ))}
      </ul>
    </div>
  );
}
// Presentational components are easy to test in Storybook
```

---

### Q5. What are Headless Components?

```jsx
// Headless: zero styling — provides BEHAVIOR only, consumer controls appearance
// Examples: Radix UI, Headless UI, Downshift, React Aria

import * as Dialog from '@radix-ui/react-dialog';

// Radix UI — accessiblity behavior built in, you bring the styles
function MyModal({ trigger, children }) {
  return (
    <Dialog.Root>
      <Dialog.Trigger asChild>{trigger}</Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="modal-overlay" />
        <Dialog.Content className="modal-content">
          {children}
          <Dialog.Close className="modal-close">✕</Dialog.Close>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}

// Why headless? Full control over styling, fully accessible (ARIA/keyboard),
// no CSS conflicts, works with any design system

// Custom headless hook example
function useToggle(initial = false) {
  const [value, setValue] = useState(initial);
  return {
    value,
    toggle: () => setValue(v => !v),
    on:     () => setValue(true),
    off:    () => setValue(false),
    // Pure logic — caller decides how to render the toggle
  };
}
```

---

### Q6. How do you build a reusable custom hook?

```jsx
// Custom hooks encapsulate stateful logic — the modern pattern over HOC/renderProps

// useDebounce — debounce any frequently changing value
function useDebounce<T>(value: T, delay: number): T {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  useEffect(() => {
    const timer = setTimeout(() => setDebouncedValue(value), delay);
    return () => clearTimeout(timer);
  }, [value, delay]);

  return debouncedValue;
}

// useOnClickOutside — close dropdown/modal when clicking outside
function useOnClickOutside(ref: RefObject<HTMLElement>, handler: () => void) {
  useEffect(() => {
    const listener = (e: MouseEvent) => {
      if (!ref.current || ref.current.contains(e.target as Node)) return;
      handler();
    };
    document.addEventListener('mousedown', listener);
    return () => document.removeEventListener('mousedown', listener);
  }, [ref, handler]);
}

// useIntersectionObserver — lazy loading, infinite scroll
function useIntersectionObserver(options = {}) {
  const [entry, setEntry] = useState<IntersectionObserverEntry | null>(null);
  const ref = useRef<HTMLElement>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([e]) => setEntry(e),
      options
    );
    if (ref.current) observer.observe(ref.current);
    return () => observer.disconnect();
  }, []);

  return { ref, isIntersecting: entry?.isIntersecting ?? false };
}

// Usage
function LazyImage({ src, alt }) {
  const { ref, isIntersecting } = useIntersectionObserver({ threshold: 0.1 });
  return (
    <div ref={ref}>
      {isIntersecting ? <img src={src} alt={alt} /> : <div className="placeholder" />}
    </div>
  );
}
```
