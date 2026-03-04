# 🛣️ 04 — Routing & Forms
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- React Router DOM v6 — Routes, Route, Link, Outlet
- useNavigate, useParams, useSearchParams
- Nested routes and protected routes
- Controlled vs uncontrolled forms
- Form validation (manual + react-hook-form)
- Error handling in forms

---

## ❓ Most Asked Questions

### Q1. How do you set up routing with React Router v6?

```jsx
import { BrowserRouter, Routes, Route, Link, Outlet } from 'react-router-dom';

// App setup
function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
        <Link to="/users">Users</Link>
      </nav>

      <Routes>
        <Route path="/"         element={<Home />} />
        <Route path="/about"    element={<About />} />
        <Route path="/users"    element={<UserLayout />}>
          <Route index          element={<UserList />} />       {/* /users */}
          <Route path=":id"     element={<UserDetail />} />    {/* /users/42 */}
          <Route path="new"     element={<NewUserForm />} />   {/* /users/new */}
        </Route>
        <Route path="*"         element={<NotFound />} />      {/* catch-all */}
      </Routes>
    </BrowserRouter>
  );
}

// Layout component with nested route outlet
function UserLayout() {
  return (
    <div>
      <h1>Users Section</h1>
      <Outlet />   {/* Nested routes render here */}
    </div>
  );
}
```

---

### Q2. How do you use `useNavigate`, `useParams`, and `useSearchParams`?

```jsx
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';

// useNavigate — programmatic navigation
function LoginForm() {
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    await login(formData);
    navigate('/dashboard');               // go to route
    navigate(-1);                         // go back
    navigate('/profile', { replace: true }); // replace history (no back)
    navigate('/checkout', { state: { fromCart: true } }); // pass state
  };

  return <form onSubmit={handleLogin}>...</form>;
}

// useParams — extract URL parameters
function UserDetail() {
  const { id } = useParams();  // from route: /users/:id

  const { data: user } = useFetch(`/api/users/${id}`);
  return <div>{user?.name}</div>;
}

// useSearchParams — query string (?search=react&page=2)
function SearchPage() {
  const [searchParams, setSearchParams] = useSearchParams();

  const query = searchParams.get('search') || '';
  const page  = Number(searchParams.get('page')) || 1;

  const updateSearch = (value) => {
    setSearchParams({ search: value, page: '1' });
  };

  return <input value={query} onChange={e => updateSearch(e.target.value)} />;
}
```

---

### Q3. How do you create protected routes?

```jsx
import { Navigate, Outlet, useLocation } from 'react-router-dom';

// Protected Route component
function ProtectedRoute({ redirectTo = '/login' }) {
  const { isLoggedIn } = useAuth();
  const location = useLocation();

  if (!isLoggedIn) {
    // Redirect to login, save intended destination in state
    return <Navigate to={redirectTo} state={{ from: location }} replace />;
  }

  return <Outlet />;  // Render child routes if authenticated
}

// Usage in router
<Routes>
  <Route path="/login"  element={<Login />} />
  <Route path="/about"  element={<About />} />

  {/* All routes under here are protected */}
  <Route element={<ProtectedRoute />}>
    <Route path="/dashboard" element={<Dashboard />} />
    <Route path="/profile"   element={<Profile />} />
    <Route path="/settings"  element={<Settings />} />
  </Route>

  {/* Admin-only routes */}
  <Route element={<ProtectedRoute redirectTo="/" />}>
    <Route path="/admin" element={<AdminPanel />} />
  </Route>
</Routes>

// Redirect back after login
function Login() {
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || '/dashboard';

  const handleLogin = async () => {
    await loginUser();
    navigate(from, { replace: true }); // return to original destination
  };
}
```

---

### Q4. How do you build a controlled form?

```jsx
function RegistrationForm() {
  const [formData, setFormData] = useState({
    name: '', email: '', password: '', role: 'user'
  });
  const [errors, setErrors] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  // Generic change handler for all fields
  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
    // Clear error when field is edited
    if (errors[name]) setErrors(prev => ({ ...prev, [name]: '' }));
  };

  const validate = () => {
    const newErrors = {};
    if (!formData.name.trim())              newErrors.name = 'Name is required';
    if (!/\S+@\S+\.\S+/.test(formData.email)) newErrors.email = 'Invalid email';
    if (formData.password.length < 8)      newErrors.password = 'Min 8 characters';
    return newErrors;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const validationErrors = validate();
    if (Object.keys(validationErrors).length) {
      setErrors(validationErrors);
      return;
    }
    setIsSubmitting(true);
    try {
      await registerUser(formData);
      navigate('/dashboard');
    } catch (err) {
      setErrors({ submit: err.message });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input name="name" value={formData.name} onChange={handleChange} />
      {errors.name && <span className="error">{errors.name}</span>}

      <input name="email" value={formData.email} onChange={handleChange} />
      {errors.email && <span className="error">{errors.email}</span>}

      <input name="password" type="password" value={formData.password} onChange={handleChange} />
      {errors.password && <span className="error">{errors.password}</span>}

      {errors.submit && <div className="error">{errors.submit}</div>}
      <button type="submit" disabled={isSubmitting}>
        {isSubmitting ? 'Registering...' : 'Register'}
      </button>
    </form>
  );
}
```

---

### Q5. How do you use React Hook Form?

```jsx
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';

// Define schema with Zod
const schema = z.object({
  name:     z.string().min(2, 'Name must be at least 2 characters'),
  email:    z.string().email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

type FormData = z.infer<typeof schema>;

function LoginForm() {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
  } = useForm<FormData>({
    resolver: zodResolver(schema),
    defaultValues: { name: '', email: '', password: '' }
  });

  const onSubmit = async (data: FormData) => {
    await registerUser(data);  // data is type-safe
    reset();
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register('name')} placeholder="Name" />
      {errors.name && <span>{errors.name.message}</span>}

      <input {...register('email')} placeholder="Email" />
      {errors.email && <span>{errors.email.message}</span>}

      <input {...register('password')} type="password" />
      {errors.password && <span>{errors.password.message}</span>}

      <button disabled={isSubmitting}>
        {isSubmitting ? 'Submitting...' : 'Register'}
      </button>
    </form>
  );
}
// Advantages: minimal re-renders, built-in validation, easy async submission
```

---

### Q6. How do you handle file uploads in React?

```jsx
function FileUpload() {
  const [preview, setPreview] = useState(null);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [error, setError] = useState(null);

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (!file) return;

    // Validate
    if (!file.type.startsWith('image/')) {
      setError('Only images allowed');
      return;
    }
    if (file.size > 5 * 1024 * 1024) {
      setError('File too large (max 5MB)');
      return;
    }

    // Preview
    const reader = new FileReader();
    reader.onload = () => setPreview(reader.result);
    reader.readAsDataURL(file);
  };

  const handleUpload = async (e) => {
    const file = e.target.files[0];
    const formData = new FormData();
    formData.append('file', file);
    formData.append('userId', '123');

    const res = await fetch('/api/upload', {
      method: 'POST',
      body: formData,   // No Content-Type header — browser sets it with boundary
    });
    const { url } = await res.json();
    console.log('Uploaded to:', url);
  };

  return (
    <div>
      <input type="file" accept="image/*" onChange={handleFileChange} />
      {error && <p className="error">{error}</p>}
      {preview && <img src={preview} alt="Preview" width={200} />}
    </div>
  );
}
```

---

### Q7. What is `Link` vs `NavLink` vs `a` in React Router?

```jsx
import { Link, NavLink } from 'react-router-dom';

// <a href> — causes full page reload ❌ for SPA navigation
<a href="/about">About</a>

// <Link> — client-side navigation without page reload ✅
<Link to="/about">About</Link>
<Link to="/user/42">User 42</Link>
<Link to={{ pathname: '/search', search: '?q=react' }}>Search</Link>

// <NavLink> — Link with active styling support
<NavLink
  to="/about"
  className={({ isActive }) => isActive ? 'nav-link active' : 'nav-link'}
>
  About
</NavLink>

// NavLink with style
<NavLink
  to="/dashboard"
  style={({ isActive }) => ({ fontWeight: isActive ? 'bold' : 'normal' })}
>
  Dashboard
</NavLink>

// Practical navigation bar
function Navbar() {
  const links = [
    { to: '/', label: 'Home' },
    { to: '/products', label: 'Products' },
    { to: '/about', label: 'About' },
  ];

  return (
    <nav>
      {links.map(({ to, label }) => (
        <NavLink
          key={to}
          to={to}
          className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}
          end   // exact match (important for "/" route)
        >
          {label}
        </NavLink>
      ))}
    </nav>
  );
}
```
