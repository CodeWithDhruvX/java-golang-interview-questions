# React Supplementary Questions (Practical & Redux Toolkit)

## 🔹 33. Redux Toolkit (Modern Redux) (Questions 326-335)

**Q326: What is Redux Toolkit (RTK) and why is it recommended?**
Redux Toolkit is the official, opinionated, batteries-included toolset for efficient Redux development. It simplifies store setup, reduces boilerplate code, and includes popular tools like Immer (for immutable state) and Thunk (for async logic) by default.

**Q327: What is `configureStore` in RTK?**
A wrapper around the standard `createStore` that automatically sets up the store with good defaults, including Redux Thunk and DevTools Extension.
```javascript
import { configureStore } from '@reduxjs/toolkit';
import rootReducer from './reducers';
const store = configureStore({ reducer: rootReducer });
```

**Q328: What is `createSlice`?**
A function that accepts an initial state, an object of reducer functions, and a "slice name", and automatically generates action creators and action types that correspond to the reducers and state.
```javascript
const counterSlice = createSlice({
  name: 'counter',
  initialState: 0,
  reducers: {
    increment: state => state + 1, // Immer handles immutability
  }
});
export const { increment } = counterSlice.actions;
export default counterSlice.reducer;
```

**Q329: How does `createAsyncThunk` work?**
A utility to define async logic. It generates three Redux action creators: `pending`, `fulfilled`, and `rejected`, which can be handled in the `extraReducers` field of a slice to update state based on the promise status.

**Q330: What is RTK Query?**
A powerful data fetching and caching tool built into Redux Toolkit. It simplifies loading data, tracking loading state, caching results, and avoiding duplicate requests.

**Q331: How do you access the store in a component using Hooks?**
Use `useSelector` to read data from the store and `useDispatch` to dispatch actions.
```javascript
const count = useSelector(state => state.counter.value);
const dispatch = useDispatch();
```

**Q332: What is the `extraReducers` field?**
It allows a slice to respond to other actions defined elsewhere (e.g., thunks or other slices). It is commonly used to handle the lifecycle actions of `createAsyncThunk`.

**Q333: What is the `createEntityAdapter` API?**
A helper to generate a set of pre-built reducers and selectors for performing CRUD operations on a normalized state structure (ids array and entities dictionary).

**Q334: How do you handle errors in RTK Query?**
RTK Query hooks return an `error` property and an `isError` boolean.
```javascript
const { data, error, isError } = useGetPokemonQuery('pikachu');
if (isError) return <div>Error: {error.message}</div>;
```

**Q335: Can Redux Toolkit be used with TypeScript?**
Yes, RTK is written in TypeScript and provides excellent type inference, making it much easier to type Redux stores and reducers compared to vanilla Redux.

---

## 🔹 34. Practical Coding Questions (Questions 336-345)

**Q336: Implement a simple Counter component.**
```jsx
import { useState } from 'react';
function Counter() {
  const [count, setCount] = useState(0);
  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={() => setCount(c => c + 1)}>+</button>
      <button onClick={() => setCount(c => c - 1)}>-</button>
    </div>
  );
}
```

**Q337: Create a custom hook `useFetch` to fetch data.**
```javascript
import { useState, useEffect } from 'react';
function useFetch(url) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  useEffect(() => {
    fetch(url)
      .then(res => res.json())
      .then(data => { setData(data); setLoading(false); });
  }, [url]);
  return { data, loading };
}
```

**Q338: Implement a "Theme Toggler" using Context.**
```jsx
const ThemeContext = React.createContext();
function ThemeProvider({ children }) {
  const [theme, setTheme] = useState("light");
  const toggle = () => setTheme(prev => prev === "light" ? "dark" : "light");
  return <ThemeContext.Provider value={{ theme, toggle }}>{children}</ThemeContext.Provider>;
}
// Usage: const { theme, toggle } = useContext(ThemeContext);
```

**Q339: Implement an Infinite Scroll component.**
Use `IntersectionObserver` to detect when the user reaches the bottom.
```jsx
const observer = useRef();
const lastElementRef = useCallback(node => {
  if (loading) return;
  if (observer.current) observer.current.disconnect();
  observer.current = new IntersectionObserver(entries => {
    if (entries[0].isIntersecting && hasMore) {
      setPageNumber(prev => prev + 1);
    }
  });
  if (node) observer.current.observe(node);
}, [loading, hasMore]);
```

**Q340: Implement a Debounce hook.**
```javascript
function useDebounce(value, delay) {
  const [debouncedValue, setDebouncedValue] = useState(value);
  useEffect(() => {
    const handler = setTimeout(() => setDebouncedValue(value), delay);
    return () => clearTimeout(handler);
  }, [value, delay]);
  return debouncedValue;
}
```

**Q341: Create a Modal using React Portals.**
```jsx
import ReactDOM from 'react-dom';
function Modal({ children, onClose }) {
  return ReactDOM.createPortal(
    <div className="overlay">
      <div className="modal">
        <button onClick={onClose}>X</button>
        {children}
      </div>
    </div>,
    document.getElementById('portal-root')
  );
}
```

**Q342: Implement a controlled Input with validation.**
```jsx
function Form() {
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const handleChange = (e) => {
    const val = e.target.value;
    setEmail(val);
    setError(val.includes("@") ? "" : "Invalid email");
  };
  return <><input value={email} onChange={handleChange} /><span>{error}</span></>;
}
```

**Q343: How do you prevent a specific child from re-rendering?**
Wrap it in `React.memo`.
```jsx
const Child = React.memo(({ name }) => {
  console.log("Rendered");
  return <div>{name}</div>;
});
```

**Q344: Implement a detailed "Star Rating" component.**
```jsx
function StarRating({ rating, setRating }) {
  return (
    <div>
      {[...Array(5)].map((_, i) => (
        <span key={i} onClick={() => setRating(i + 1)} 
              style={{ color: i < rating ? "gold" : "gray", cursor: "pointer" }}>
          ★
        </span>
      ))}
    </div>
  );
}
```

**Q345: Implement a basic "Protected Route" wrapper.**
```jsx
import { Navigate, Outlet } from 'react-router-dom';
const PrivateRoutes = () => {
  const isAuth = useAuth(); // Custom hook logic
  return isAuth ? <Outlet /> : <Navigate to="/login" />;
}
```
