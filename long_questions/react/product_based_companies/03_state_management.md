# 🗃️ 03 — State Management in React
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

> Product companies expect you to justify your state management choice and understand the trade-offs between Redux, Zustand, Context, and atomic state libraries.

---

## 🔑 Must-Know Topics
- Local vs global state
- Redux Toolkit (RTK) and RTK Query
- Zustand — lightweight alternative
- Context API trade-offs
- Recoil / Jotai (atomic state)
- Derived state vs redundant state
- Server state vs client state (React Query / TanStack Query)

---

## ❓ Most Asked Questions

### Q1. What is the difference between local, shared, and global state?

```jsx
// LOCAL STATE — belongs to one component, doesn't need to be shared
function SearchBar() {
  const [query, setQuery] = useState('');  // Local: only this component needs it
  return <input value={query} onChange={e => setQuery(e.target.value)} />;
}

// SHARED STATE — needed by a few related components (lift up)
function ProductFilter() {
  const [filters, setFilters] = useState({ category: 'all', price: 100 });
  return (
    <>
      <FilterPanel filters={filters} onChange={setFilters} />
      <ProductList filters={filters} />
    </>
  );
}

// GLOBAL STATE — needs to be accessible from anywhere: auth, cart, theme
// Use Context, Redux, Zustand etc.
const useAuthStore = create(set => ({
  user: null,
  login:  (user) => set({ user }),
  logout: () => set({ user: null })
}));

// SERVER STATE — remote data: loading, caching, invalidation
// Use React Query / RTK Query — NOT useState + useEffect
const { data: products, isLoading } = useQuery({
  queryKey: ['products'],
  queryFn: () => fetch('/api/products').then(r => r.json())
});
```

---

### Q2. How does Redux Toolkit (RTK) simplify Redux?

```jsx
// Modern Redux — use Redux Toolkit, not legacy redux + react-redux

// 1. Create a slice (replaces action creators + reducer boilerplate)
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

// Async thunk for API calls
export const fetchUsers = createAsyncThunk('users/fetchAll', async () => {
  const res = await fetch('/api/users');
  return res.json(); // returned value becomes action.payload in fulfilled
});

const usersSlice = createSlice({
  name: 'users',
  initialState: { list: [], status: 'idle', error: null },

  reducers: {
    addUser: (state, action) => {
      state.list.push(action.payload);  // Immer allows "mutations" ✅
    },
    removeUser: (state, action) => {
      state.list = state.list.filter(u => u.id !== action.payload);
    },
  },

  extraReducers: (builder) => {
    builder
      .addCase(fetchUsers.pending,   (state) => { state.status = 'loading'; })
      .addCase(fetchUsers.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.list = action.payload;
      })
      .addCase(fetchUsers.rejected,  (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      });
  }
});

export const { addUser, removeUser } = usersSlice.actions;
export default usersSlice.reducer;

// 2. Configure store
import { configureStore } from '@reduxjs/toolkit';
const store = configureStore({
  reducer: { users: usersSlice.reducer }
});

// 3. Use in component
function UserList() {
  const dispatch = useDispatch();
  const { list, status } = useSelector(state => state.users);

  useEffect(() => { dispatch(fetchUsers()); }, [dispatch]);

  return status === 'loading' ? <Spinner /> : <ul>{list.map(renderUser)}</ul>;
}
```

---

### Q3. How does Zustand work and when should you choose it over Redux?

```jsx
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

// Zustand store — minimal boilerplate
const useCartStore = create(persist(
  (set, get) => ({
    items: [],

    addItem: (product) => set(state => ({
      items: [...state.items, { ...product, qty: 1 }]
    })),

    removeItem: (id) => set(state => ({
      items: state.items.filter(i => i.id !== id)
    })),

    updateQty: (id, qty) => set(state => ({
      items: state.items.map(i => i.id === id ? { ...i, qty } : i)
    })),

    // Derived — computed from state
    get total() {
      return get().items.reduce((sum, i) => sum + i.price * i.qty, 0);
    },

    clearCart: () => set({ items: [] }),
  }),
  { name: 'cart-storage' }  // persist to localStorage
));

// Usage — no Provider needed!
function Cart() {
  const { items, removeItem, total } = useCartStore();
  return (
    <div>
      {items.map(item => <CartItem key={item.id} item={item} onRemove={removeItem} />)}
      <strong>Total: ${total}</strong>
    </div>
  );
}

// Zustand vs Redux:
// ✅ Zustand: small apps, no boilerplate, no Provider, easier learning curve
// ✅ Redux:   large teams, strong conventions, DevTools time-travel, RTK Query
```

---

### Q4. What are selectors and why do they matter for performance?

```jsx
// Selectors compute derived data from the store
// Without memoized selectors → re-render on ANY store change

import { createSelector } from '@reduxjs/toolkit';

// ❌ Non-memoized — new array reference every time → always re-renders
const selectActiveUsers = (state) =>
  state.users.list.filter(u => u.isActive); // new [] reference each call

// ✅ Memoized selector with createSelector (reselect library)
const selectActiveUsers = createSelector(
  [(state) => state.users.list],      // input selectors
  (users) => users.filter(u => u.isActive) // result memo'd until input changes
);

// Usage
function ActiveUserList() {
  const activeUsers = useSelector(selectActiveUsers);
  // Only re-renders when the filtered result actually changes
}

// Zustand equivalent — subscribe to slice of state
const cartItems = useCartStore(state => state.items);
// Only re-renders when items changes (not when other state changes)
```

---

### Q5. What is server state and how does TanStack Query handle it?

```jsx
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

// Server state concerns: loading, caching, synchronizing, invalidating
// Problem with useState + useEffect:
// - No cache
// - Manual loading/error handling
// - Re-fetches on every mount
// - Complex dependency management

// ✅ React Query handles all of this
const queryClient = new QueryClient();

function UserProfile({ userId }) {
  // Automatic caching, background refetch, retry logic
  const { data: user, isLoading, error } = useQuery({
    queryKey: ['user', userId],  // cache key — unique per userId
    queryFn: () => fetchUser(userId),
    staleTime: 5 * 60 * 1000,   // 5 minutes before considering data stale
    gcTime: 10 * 60 * 1000,     // 10 minutes before garbage collection
  });

  if (isLoading) return <Skeleton />;
  if (error) return <Error message={error.message} />;
  return <div>{user.name}</div>;
}

// Mutations with cache invalidation
function UpdateUserForm({ userId }) {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (updates) => updateUser(userId, updates),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['user', userId] });
      // ↑ triggers refetch of user data after successful update
    },
  });

  return <form onSubmit={data => mutation.mutate(data)}>...</form>;
}
```

---

### Q6. What is Recoil / Jotai (atomic state)?

```jsx
// Atomic state: state split into small atoms — components subscribe to specific atoms
// Only components subscribed to an atom re-render when IT changes

// --- Jotai (simpler API) ---
import { atom, useAtom, useAtomValue, useSetAtom } from 'jotai';

// Atoms — minimal state units
const countAtom  = atom(0);
const nameAtom   = atom('');
const isLoadingAtom = atom(false);

// Derived atom (like a selector)
const doubleCountAtom = atom(get => get(countAtom) * 2);

function Counter() {
  const [count, setCount] = useAtom(countAtom);
  return <button onClick={() => setCount(c => c + 1)}>{count}</button>;
}

function DoubleDisplay() {
  const double = useAtomValue(doubleCountAtom);
  return <div>Double: {double}</div>; // Only re-renders when countAtom changes
}

// Jotai vs Recoil:
// Jotai: simpler, no atom keys, smaller bundle, React Suspense friendly
// Recoil: Meta-backed, selector graph, async selectors

// When to use atomic state:
// ✅ Complex derived state graphs
// ✅ Many independent state slices (spreadsheets, form builders)
// ✅ Fine-grained subscriptions needed
```

---

### Q7. How do you handle forms with complex state?

```jsx
// useReducer for complex form state (better than multiple useState)
const formReducer = (state, action) => {
  switch (action.type) {
    case 'SET_FIELD':
      return { ...state, [action.field]: action.value };
    case 'SET_ERROR':
      return { ...state, errors: { ...state.errors, [action.field]: action.message }};
    case 'RESET':
      return initialState;
    default:
      return state;
  }
};

const initialState = { name: '', email: '', password: '', errors: {} };

function SignupForm() {
  const [state, dispatch] = useReducer(formReducer, initialState);

  const handleChange = (e) => {
    dispatch({ type: 'SET_FIELD', field: e.target.name, value: e.target.value });
  };

  return (
    <form>
      <input name="name" value={state.name} onChange={handleChange} />
      <input name="email" value={state.email} onChange={handleChange} />
      {state.errors.email && <span>{state.errors.email}</span>}
    </form>
  );
}

// For production forms: React Hook Form (best performance)
import { useForm } from 'react-hook-form';
const { register, handleSubmit, formState: { errors } } = useForm();
// Only re-renders on submit — not on every keystroke
```
