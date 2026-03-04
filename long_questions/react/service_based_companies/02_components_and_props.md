# 🧩 02 — Components & Props
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Functional vs class components
- Props and PropTypes validation
- Default props
- Component composition
- Lifting state up
- Pure components
- `children` prop patterns

---

## ❓ Most Asked Questions

### Q1. What is the difference between functional and class components?

```jsx
// --- Class Component (legacy — still maintained) ---
class Welcome extends React.Component {
  constructor(props) {
    super(props);
    this.state = { count: 0 };
  }

  handleClick = () => this.setState({ count: this.state.count + 1 });

  render() {
    return (
      <div>
        <h1>Hello, {this.props.name}</h1>
        <p>Count: {this.state.count}</p>
        <button onClick={this.handleClick}>Increment</button>
      </div>
    );
  }
}

// --- Functional Component (modern — preferred) ---
function Welcome({ name }) {
  const [count, setCount] = useState(0);
  return (
    <div>
      <h1>Hello, {name}</h1>
      <p>Count: {count}</p>
      <button onClick={() => setCount(c => c + 1)}>Increment</button>
    </div>
  );
}
```

| Feature | Class | Functional |
|---------|-------|------------|
| State | `this.state` | `useState` |
| Lifecycle | Methods | `useEffect` |
| `this` binding | Required | Not needed |
| Reusability | Mixins (deprecated) | Custom Hooks |
| Code size | Verbose | Concise |

---

### Q2. What is PropTypes and how do you use it?

```jsx
import PropTypes from 'prop-types';

function UserProfile({ name, age, email, role, address, onClick }) {
  return <div>{name}</div>;
}

UserProfile.propTypes = {
  name:    PropTypes.string.isRequired,
  age:     PropTypes.number,
  email:   PropTypes.string.isRequired,
  role:    PropTypes.oneOf(['admin', 'user', 'guest']),
  address: PropTypes.shape({
    city:    PropTypes.string,
    country: PropTypes.string,
  }),
  onClick: PropTypes.func,
};

UserProfile.defaultProps = {
  age:     18,
  role:    'user',
  onClick: () => {},
};

// PropTypes validates in development only (not production)
// For production-grade type safety, use TypeScript instead:
interface UserProfileProps {
  name:  string;
  age?:  number;
  email: string;
}
```

---

### Q3. How do you lift state up?

```jsx
// Problem: sibling components need to share state
// Solution: lift state to their common parent

// ❌ Both components manage their own state — can't sync
function TemperatureInCelsius() {
  const [temp, setTemp] = useState('');
  return <input value={temp} onChange={e => setTemp(e.target.value)} />;
}

// ✅ Lifted state — parent owns state, passes down as props
function TemperatureConverter() {
  const [celsius, setCelsius] = useState('');

  const fahrenheit = celsius ? (celsius * 9/5 + 32).toFixed(1) : '';

  return (
    <div>
      <TemperatureInput
        label="Celsius"
        value={celsius}
        onChange={setCelsius}
      />
      <TemperatureInput
        label="Fahrenheit"
        value={fahrenheit}
        onChange={f => setCelsius(((f - 32) * 5/9).toFixed(1))}
      />
    </div>
  );
}

function TemperatureInput({ label, value, onChange }) {
  return (
    <label>
      {label}: <input value={value} onChange={e => onChange(e.target.value)} />
    </label>
  );
}
```

---

### Q4. What is component composition?

```jsx
// Composition over inheritance — React's core pattern
// Build complex UIs by combining smaller components

// Specialization via composition
function Dialog({ title, message, onClose }) {
  return (
    <div className="dialog">
      <h2>{title}</h2>
      <p>{message}</p>
      <button onClick={onClose}>Close</button>
    </div>
  );
}

// Specialized version
function ConfirmDialog({ onConfirm, onCancel }) {
  return (
    <div className="dialog">
      <p>Are you sure?</p>
      <button onClick={onConfirm}>Yes</button>
      <button onClick={onCancel}>No</button>
    </div>
  );
}

// Containment pattern with children
function Panel({ title, children, footer }) {
  return (
    <div className="panel">
      <div className="panel-header">{title}</div>
      <div className="panel-body">{children}</div>
      {footer && <div className="panel-footer">{footer}</div>}
    </div>
  );
}

// Usage
<Panel
  title="User Info"
  footer={<button>Save</button>}
>
  <UserForm />
</Panel>
```

---

### Q5. What are pure components?

```jsx
// A pure component only re-renders when props or state change
// (shallow comparison)

// Class Pure Component
class PureCounter extends React.PureComponent {
  render() {
    return <div>{this.props.count}</div>;
    // Only re-renders when count prop actually changes
  }
}

// Functional equivalent: React.memo
const PureCounter = React.memo(function Counter({ count }) {
  return <div>{count}</div>;
});

// ⚠️ Shallow comparison only:
// primitives: number, string, boolean — works fine
// objects/arrays: compares REFERENCE, not content!

const obj = { a: 1 };
<PureCounter data={obj} />  // Same reference → skip render ✅
<PureCounter data={{ a: 1 }} /> // New object each render → always renders ❌
```

---

### Q6. How do you handle events in React?

```jsx
function EventExamples() {
  // Synthetic event — React wraps native events
  const handleClick = (e) => {
    e.preventDefault();
    e.stopPropagation();
    console.log('Clicked:', e.target.value);
  };

  // Passing arguments
  const handleItemClick = (id, e) => {
    console.log('Item:', id, 'Event:', e);
  };

  // Keyboard events
  const handleKeyDown = (e) => {
    if (e.key === 'Enter') submitForm();
    if (e.key === 'Escape') closeModal();
  };

  // Form events
  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setForm(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  return (
    <div>
      <button onClick={handleClick}>Click</button>
      <button onClick={(e) => handleItemClick(42, e)}>Item</button>
      <input onKeyDown={handleKeyDown} onChange={handleChange} />
    </div>
  );
}
```

---

### Q7. What are Higher-Order Components (HOC)?

```jsx
// HOC = function that takes a component and returns enhanced component
// Naming convention: withXxx

function withAuth(WrappedComponent) {
  return function AuthenticatedComponent(props) {
    const { isLoggedIn, user } = useAuth();

    if (!isLoggedIn) {
      return <Navigate to="/login" />;
    }

    return <WrappedComponent {...props} currentUser={user} />;
  };
}

// Usage
const ProtectedDashboard = withAuth(Dashboard);

// Another HOC — withLoading
function withLoading(WrappedComponent) {
  return function({ isLoading, ...props }) {
    if (isLoading) return <Spinner />;
    return <WrappedComponent {...props} />;
  };
}

// Modern alternative: Custom Hooks (preferred for logic reuse)
// HOCs are still valid for wrapping third-party components
```

---

### Q8. What is the difference between `defaultProps` and default parameters?

```jsx
// defaultProps (class/functional — legacy approach)
function Button({ label, variant }) {
  return <button className={variant}>{label}</button>;
}
Button.defaultProps = {
  label:   'Submit',
  variant: 'primary',
};

// Default parameters (modern — preferred)
function Button({ label = 'Submit', variant = 'primary' }) {
  return <button className={variant}>{label}</button>;
}

// ⚠️ defaultProps is deprecated for function components in React 18.3+
// Always use default parameter syntax for functional components

// TypeScript with defaults
interface ButtonProps {
  label?:   string;
  variant?: 'primary' | 'secondary';
}
function Button({ label = 'Submit', variant = 'primary' }: ButtonProps) {
  return <button className={variant}>{label}</button>;
}
```
