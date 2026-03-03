# 🎨 06 — Styling & UI in React
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- CSS Modules — scoped styles
- Styled Components — CSS-in-JS
- Tailwind CSS basics
- Material UI (MUI) — component library
- Inline styles and dynamic styles
- CSS variables in React
- Responsive design in React

---

## ❓ Most Asked Questions

### Q1. What are CSS Modules and how do they work?

```jsx
// CSS Modules: scoped CSS — class names are auto-namespaced per file
// Prevents class name collisions across components

// Button.module.css
.button { padding: 8px 16px; border-radius: 4px; }
.primary { background: blue; color: white; }
.secondary { background: gray; color: white; }
.large { font-size: 18px; padding: 12px 24px; }

// Button.jsx
import styles from './Button.module.css';

function Button({ variant = 'primary', size, children, onClick }) {
  const classes = [
    styles.button,
    styles[variant],           // dynamic class
    size === 'large' ? styles.large : ''
  ].filter(Boolean).join(' ');

  return (
    <button className={classes} onClick={onClick}>
      {children}
    </button>
  );
}

// With clsx library (cleaner class composition)
import clsx from 'clsx';

const className = clsx(
  styles.button,
  styles[variant],
  { [styles.large]: size === 'large', [styles.disabled]: disabled }
);

// Result: <button class="Button_button__3kVQ2 Button_primary__8Bx9k">
// Classes are unique — no global conflicts!
```

---

### Q2. What is Styled Components (CSS-in-JS)?

```jsx
import styled, { css, ThemeProvider } from 'styled-components';

// Basic styled component
const Button = styled.button`
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: opacity 0.2s;

  /* Dynamic styles based on props */
  background: ${props => props.variant === 'primary' ? '#007bff' : '#6c757d'};
  color: white;

  ${props => props.size === 'large' && css`
    padding: 12px 24px;
    font-size: 18px;
  `}

  &:hover { opacity: 0.85; }
  &:disabled { opacity: 0.5; cursor: not-allowed; }
`;

// Usage
<Button variant="primary" size="large" onClick={handleClick}>
  Submit
</Button>

// Theme support
const theme = { colors: { primary: '#007bff', text: '#333' } };

const Heading = styled.h1`
  color: ${({ theme }) => theme.colors.primary};
  font-size: 2rem;
`;

<ThemeProvider theme={theme}>
  <Heading>Welcome</Heading>
</ThemeProvider>

// Extending styles
const DangerButton = styled(Button)`
  background: #dc3545;
  &:hover { background: #c82333; }
`;
```

---

### Q3. How do you use Tailwind CSS with React?

```jsx
// Tailwind: utility-first CSS — compose classes directly in JSX
// No custom CSS needed for most cases

// Basic component with Tailwind
function Card({ title, description, imageUrl }) {
  return (
    <div className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-lg transition-shadow">
      <img className="w-full h-48 object-cover" src={imageUrl} alt={title} />
      <div className="p-6">
        <h2 className="text-xl font-bold text-gray-800 mb-2">{title}</h2>
        <p className="text-gray-600 text-sm">{description}</p>
      </div>
    </div>
  );
}

// Responsive design with Tailwind breakpoints
<div className="
  grid
  grid-cols-1
  sm:grid-cols-2
  md:grid-cols-3
  lg:grid-cols-4
  gap-4
">
  {products.map(p => <ProductCard key={p.id} product={p} />)}
</div>

// Dynamic classes with clsx (don't concatenate strings with template literals)
import clsx from 'clsx';
<button
  className={clsx(
    'px-4 py-2 rounded font-medium transition-colors',
    variant === 'primary' && 'bg-blue-500 text-white hover:bg-blue-600',
    variant === 'danger'  && 'bg-red-500 text-white hover:bg-red-600',
    disabled && 'opacity-50 cursor-not-allowed'
  )}
>
  {children}
</button>
```

---

### Q4. How do you use Material UI (MUI)?

```jsx
import {
  Button, TextField, Box, Typography,
  Card, CardContent, CardActions,
  ThemeProvider, createTheme
} from '@mui/material';

// Custom theme
const theme = createTheme({
  palette: {
    primary: { main: '#1976d2' },
    secondary: { main: '#9c27b0' },
  },
  typography: {
    fontFamily: 'Inter, sans-serif',
  },
});

// Basic form with MUI components
function ContactForm() {
  const [name, setName] = useState('');

  return (
    <ThemeProvider theme={theme}>
      <Box component="form" sx={{ display: 'flex', flexDirection: 'column', gap: 2, p: 3 }}>
        <Typography variant="h5" fontWeight="bold">Contact Us</Typography>

        <TextField
          label="Name"
          variant="outlined"
          value={name}
          onChange={e => setName(e.target.value)}
          required
          fullWidth
        />
        <TextField label="Email" type="email" variant="outlined" fullWidth />
        <TextField label="Message" multiline rows={4} variant="outlined" fullWidth />

        <Button variant="contained" color="primary" size="large" type="submit">
          Send Message
        </Button>
      </Box>
    </ThemeProvider>
  );
}

// sx prop — inline styles using theme tokens
<Box sx={{
  bgcolor: 'primary.main',
  color: 'white',
  p: 2,                     // padding: theme.spacing(2) = 16px
  borderRadius: 1,
  '&:hover': { bgcolor: 'primary.dark' }
}} />
```

---

### Q5. How do you apply dynamic and conditional styles?

```jsx
// 1. Inline styles (use sparingly — no pseudo-classes/media queries)
const style = {
  color: isError ? 'red' : 'green',
  fontSize: `${size * 16}px`,
  transform: `rotate(${rotation}deg)`,
};
<div style={style}>{message}</div>

// 2. className with template literals
<div className={`card ${isActive ? 'active' : ''} ${variant}`}>

// 3. clsx library (recommended)
import clsx from 'clsx';
<button className={clsx('btn', {
  'btn-primary': variant === 'primary',
  'btn-danger':  variant === 'danger',
  'btn-loading': isLoading,
  'btn-disabled': disabled
})}>

// 4. CSS variables for dynamic values
function ProgressBar({ value }) {
  return (
    <div
      className="progress-bar"
      style={{ '--progress': `${value}%` }}   // CSS variable
    />
  );
}

// In CSS:
// .progress-bar::after { width: var(--progress); }

// 5. Styled components with dynamic props (as shown in Q2)
const Box = styled.div`
  background: ${props => props.theme.colors[props.color] || '#fff'};
`;
```

---

### Q6. How do you handle dark mode in React?

```jsx
// Using CSS variables + context
const ThemeContext = createContext(null);

function ThemeProvider({ children }) {
  const [isDark, setIsDark] = useState(() => {
    return localStorage.getItem('theme') === 'dark' ||
           window.matchMedia('(prefers-color-scheme: dark)').matches;
  });

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', isDark ? 'dark' : 'light');
    localStorage.setItem('theme', isDark ? 'dark' : 'light');
  }, [isDark]);

  return (
    <ThemeContext.Provider value={{ isDark, toggle: () => setIsDark(d => !d) }}>
      {children}
    </ThemeContext.Provider>
  );
}

// CSS variables in :root and [data-theme="dark"]
/*
:root {
  --bg-primary: #ffffff;
  --text-primary: #1a1a1a;
  --surface: #f5f5f5;
}
[data-theme="dark"] {
  --bg-primary: #1a1a1a;
  --text-primary: #f5f5f5;
  --surface: #2d2d2d;
}
body { background: var(--bg-primary); color: var(--text-primary); }
*/

// Toggle button
function DarkModeToggle() {
  const { isDark, toggle } = useContext(ThemeContext);
  return (
    <button onClick={toggle}>
      {isDark ? '☀️ Light Mode' : '🌙 Dark Mode'}
    </button>
  );
}
```

---

### Q7. How do you create a responsive layout in React?

```jsx
// Method 1: CSS Grid/Flexbox with media queries (in CSS file)
// grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));

// Method 2: useMediaQuery hook
function useMediaQuery(query) {
  const [matches, setMatches] = useState(() =>
    window.matchMedia(query).matches
  );

  useEffect(() => {
    const media = window.matchMedia(query);
    const handler = (e) => setMatches(e.matches);
    media.addEventListener('change', handler);
    return () => media.removeEventListener('change', handler);
  }, [query]);

  return matches;
}

function Sidebar() {
  const isMobile = useMediaQuery('(max-width: 768px)');
  const [isOpen, setIsOpen] = useState(false);

  if (isMobile) {
    return (
      <>
        <button onClick={() => setIsOpen(true)}>☰ Menu</button>
        {isOpen && (
          <div className="mobile-drawer">
            <button onClick={() => setIsOpen(false)}>✕</button>
            <NavLinks />
          </div>
        )}
      </>
    );
  }

  return (
    <aside className="sidebar">
      <NavLinks />
    </aside>
  );
}
```
