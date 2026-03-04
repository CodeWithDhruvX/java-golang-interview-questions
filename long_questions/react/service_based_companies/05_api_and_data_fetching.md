# 🌐 05 — API & Data Fetching
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- fetch API and error handling
- Axios setup and interceptors
- Loading/error/success state pattern
- React Query (TanStack Query) basics
- Abort controllers for cancellation
- CORS and common API issues
- Environment variables in React

---

## ❓ Most Asked Questions

### Q1. How do you fetch data in React using `useEffect`?

```jsx
function UserList() {
  const [users, setUsers]     = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError]     = useState(null);

  useEffect(() => {
    let cancelled = false;  // prevent state update after unmount

    setLoading(true);
    setError(null);

    fetch('https://jsonplaceholder.typicode.com/users')
      .then(res => {
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
        return res.json();
      })
      .then(data => {
        if (!cancelled) {
          setUsers(data);
          setLoading(false);
        }
      })
      .catch(err => {
        if (!cancelled) {
          setError(err.message);
          setLoading(false);
        }
      });

    return () => { cancelled = true; }; // cleanup on unmount
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error)   return <p>Error: {error}</p>;
  return <ul>{users.map(u => <li key={u.id}>{u.name}</li>)}</ul>;
}
```

---

### Q2. How do you use Axios in React?

```jsx
import axios from 'axios';

// Axios instance with base config (create once in api/client.js)
const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || 'https://api.example.com',
  timeout: 10000,
  headers: { 'Content-Type': 'application/json' }
});

// Request interceptor — add auth token to every request
api.interceptors.request.use(
  config => {
    const token = localStorage.getItem('auth_token');
    if (token) config.headers.Authorization = `Bearer ${token}`;
    return config;
  },
  error => Promise.reject(error)
);

// Response interceptor — handle 401, global errors
api.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      localStorage.removeItem('auth_token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Using the instance in a component
function UserProfile({ userId }) {
  const [user, setUser]     = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const controller = new AbortController();

    api.get(`/users/${userId}`, { signal: controller.signal })
       .then(res => setUser(res.data))
       .catch(err => {
         if (!axios.isCancel(err)) console.error(err);
       })
       .finally(() => setLoading(false));

    return () => controller.abort();
  }, [userId]);

  return loading ? <Spinner /> : <div>{user?.name}</div>;
}
```

---

### Q3. How does React Query simplify data fetching?

```jsx
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

// 1. Setup — wrap app with QueryClientProvider
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      staleTime: 5 * 60 * 1000, // 5 minutes
    }
  }
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Router>...</Router>
    </QueryClientProvider>
  );
}

// 2. Fetch with useQuery — handles loading, error, caching automatically
function UserList() {
  const {
    data: users,
    isLoading,
    isError,
    error,
    refetch,
  } = useQuery({
    queryKey: ['users'],
    queryFn: () => api.get('/users').then(r => r.data),
  });

  if (isLoading) return <Spinner />;
  if (isError)   return <p>Error: {error.message} <button onClick={refetch}>Retry</button></p>;

  return <ul>{users.map(u => <li key={u.id}>{u.name}</li>)}</ul>;
}

// 3. Mutations — create/update/delete with cache invalidation
function AddUserForm() {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (newUser) => api.post('/users', newUser).then(r => r.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] }); // refetch list
    },
    onError: (err) => toast.error(err.message),
  });

  return (
    <form onSubmit={e => {
      e.preventDefault();
      mutation.mutate({ name: 'Alice', email: 'alice@example.com' });
    }}>
      <button type="submit" disabled={mutation.isPending}>
        {mutation.isPending ? 'Adding...' : 'Add User'}
      </button>
    </form>
  );
}
```

---

### Q4. How do you handle different API states in the UI?

```jsx
// Standard pattern: loading + error + empty + data states
function ProductList() {
  const { data, isLoading, isError, error } = useQuery({
    queryKey: ['products'],
    queryFn: fetchProducts,
  });

  // Loading state
  if (isLoading) {
    return (
      <div className="grid">
        {Array.from({ length: 6 }).map((_, i) => (
          <div key={i} className="skeleton-card" />  // skeleton loading
        ))}
      </div>
    );
  }

  // Error state
  if (isError) {
    return (
      <div className="error-container">
        <h3>Oops! Something went wrong</h3>
        <p>{error.message}</p>
        <button onClick={() => window.location.reload()}>Try Again</button>
      </div>
    );
  }

  // Empty state
  if (data.length === 0) {
    return (
      <div className="empty-state">
        <img src="/empty-box.svg" alt="No products" />
        <p>No products found</p>
      </div>
    );
  }

  // Data state
  return (
    <div className="grid">
      {data.map(product => <ProductCard key={product.id} product={product} />)}
    </div>
  );
}
```

---

### Q5. How do you cancel API requests in React?

```jsx
// Abort Controller — cancel fetch requests
function SearchResults({ query }) {
  const [results, setResults] = useState([]);

  useEffect(() => {
    if (!query.trim()) { setResults([]); return; }

    const controller = new AbortController();

    fetch(`/api/search?q=${encodeURIComponent(query)}`, {
      signal: controller.signal
    })
      .then(res => res.json())
      .then(data => setResults(data))
      .catch(err => {
        if (err.name !== 'AbortError') {  // ignore abort errors
          console.error('Search failed:', err);
        }
      });

    // Cancel previous request when query changes
    return () => controller.abort();
  }, [query]);

  return <ul>{results.map(r => <li key={r.id}>{r.title}</li>)}</ul>;
}

// Axios equivalent
useEffect(() => {
  const source = axios.CancelToken.source();

  axios.get('/api/search', {
    params: { q: query },
    cancelToken: source.token
  })
    .then(res => setResults(res.data))
    .catch(err => {
      if (!axios.isCancel(err)) console.error(err);
    });

  return () => source.cancel();
}, [query]);
```

---

### Q6. How do you use environment variables in React?

```bash
# .env file (at project root)
REACT_APP_API_URL=https://api.dev.example.com   # CRA
REACT_APP_API_KEY=abc123

# For Vite:
VITE_API_URL=https://api.dev.example.com
VITE_API_KEY=abc123

# .env.production
REACT_APP_API_URL=https://api.example.com
```

```jsx
// Access in code (CRA)
const API_URL = process.env.REACT_APP_API_URL;
// process.env.REACT_APP_* is replaced at BUILD TIME by webpack

// Access in code (Vite)
const API_URL = import.meta.env.VITE_API_URL;

// Usage
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

// ⚠️ Rules:
// 1. Never commit .env files to git (add to .gitignore)
// 2. NEVER put secrets/private keys here — they're visible in browser JS!
// 3. Only PUBLIC values (API URLs, public keys) go in client .env
// 4. Create .env.example with placeholder values for teammates
```

---

### Q7. What is CORS and how do you handle it in React?

```jsx
// CORS = Cross-Origin Resource Sharing
// Browser blocks requests from different origins (domain/port/protocol)

// ❌ Error: "CORS policy: No 'Access-Control-Allow-Origin' header"
// This is a SERVER-side issue, but here's how to work around it in dev:

// CRA — proxy in package.json (development only)
// package.json:
{
  "proxy": "http://localhost:5000"
}
// Now fetch('/api/users') proxies to http://localhost:5000/api/users

// Vite — vite.config.js proxy
export default {
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:5000',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
};

// Production: CORS must be configured on the SERVER
// For Express:
const cors = require('cors');
app.use(cors({ origin: 'https://myapp.com' }));

// Quick note: CORS errors visible in dev = server needs fixing
// They are NOT a frontend React issue (despite appearing in browser console)
```
