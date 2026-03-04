# 📘 01 — React Basics & JSX
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy

---

## 🔑 Must-Know Topics
- What is React and why use it
- JSX syntax and rules
- React DOM vs React
- Elements vs Components
- Rendering to the DOM
- React vs other frameworks
- CRA vs Vite project setup

---

## ❓ Most Asked Questions

### Q1. What is React? What are its key features?

**Answer:**
React is an open-source JavaScript **library** (not a framework) for building user interfaces, developed by Meta. Key features:
- **Component-Based** — UI split into reusable, independent pieces
- **Declarative** — describe what UI should look like, React handles updates
- **Virtual DOM** — efficient updates via in-memory DOM diffing
- **Unidirectional Data Flow** — data flows parent → child via props
- **Hooks** — state and side effects in functional components
- **Large Ecosystem** — React Router, Redux, Next.js, etc.

---

### Q2. What is JSX and why do we use it?

```jsx
// JSX = JavaScript XML — lets you write HTML-like code in JS
const element = <h1 className="title">Hello, React!</h1>;

// Rules of JSX:
// 1. Must return a single root element (use <> fragment if needed)
// 2. className instead of class (JS reserved word)
// 3. htmlFor instead of for
// 4. Self-close empty tags: <img />, <br />, <input />
// 5. JavaScript expressions in curly braces {}

const name = "Alice";
const greeting = <h1>Hello, {name}!</h1>;                    // ✅
const sum = <p>2 + 2 = {2 + 2}</p>;                          // ✅
const isAdmin = true;
const badge = <span>{isAdmin ? "Admin" : "User"}</span>;     // ✅

// Fragment — avoid adding extra div to DOM
const Layout = () => (
  <>
    <Header />
    <Main />
    <Footer />
  </>
);
```

---

### Q3. What is the difference between a React Element and a Component?

```jsx
// --- React Element: plain JS object describing what to render ---
const element = <h1>Hello</h1>;
// { type: 'h1', props: { children: 'Hello' }, key: null, ref: null }

// --- Component: a function (or class) that returns elements ---
function Greeting({ name }) {
  return <h1>Hello, {name}!</h1>;  // returns an element
}

// Using the component — React calls Greeting({ name: "Bob" })
const app = <Greeting name="Bob" />;

// Key difference:
// Element: lowercase tag — maps to real DOM node
// Component: PascalCase — React calls it as a function
// <div /> → DOM element
// <Greeting /> → React component call
```

---

### Q4. How do you render a React app to the DOM?

```jsx
// React 18 (createRoot API)
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// React 17 and below (legacy):
import ReactDOM from 'react-dom';
ReactDOM.render(<App />, document.getElementById('root'));

// index.html has:
// <div id="root"></div>  ← React mounts here
```

---

### Q5. What are props in React? How do you pass them?

```jsx
// Props = Properties — data passed from parent to child (read-only)
function UserCard({ name, age, isAdmin = false }) {  // with default prop
  return (
    <div>
      <h2>{name}</h2>
      <p>Age: {age}</p>
      {isAdmin && <span>Admin</span>}
    </div>
  );
}

// Passing props from parent
function App() {
  return (
    <>
      <UserCard name="Alice" age={30} isAdmin={true} />
      <UserCard name="Bob" age={25} />      {/* isAdmin defaults to false */}
    </>
  );
}

// Spread props (use carefully)
const userProps = { name: "Charlie", age: 28 };
<UserCard {...userProps} />

// Props are READ-ONLY — never mutate props inside a component
```

---

### Q6. What is the `children` prop?

```jsx
// children is a special prop — content between opening and closing tags
function Card({ children, title }) {
  return (
    <div className="card">
      <h3>{title}</h3>
      <div className="card-body">{children}</div>
    </div>
  );
}

// Usage
function App() {
  return (
    <Card title="My Card">
      <p>This is the card content</p>   {/* becomes children */}
      <button>Click me</button>
    </Card>
  );
}

// children can be: string, number, element, array, or function
// Type: React.ReactNode in TypeScript
```

---

### Q7. What are React Fragments and why use them?

```jsx
// Problem: React requires a single root element — but extra divs pollute DOM
function Table() {
  // ❌ Adds unnecessary div inside <tbody>
  return (
    <div>
      <tr><td>Row 1</td></tr>
      <tr><td>Row 2</td></tr>
    </div>
  );
}

// ✅ Fragment — groups without adding DOM node
function Table() {
  return (
    <React.Fragment>
      <tr><td>Row 1</td></tr>
      <tr><td>Row 2</td></tr>
    </React.Fragment>
  );
}

// ✅ Short syntax (<> </>) — most common
function Table() {
  return (
    <>
      <tr><td>Row 1</td></tr>
      <tr><td>Row 2</td></tr>
    </>
  );
}

// Note: <React.Fragment> supports key prop (useful in lists)
// Short <> syntax does NOT support key
```

---

### Q8. How do you conditionally render in React?

```jsx
function Dashboard({ isLoggedIn, hasData, count }) {
  // 1. if/else
  if (!isLoggedIn) return <Login />;

  // 2. Ternary operator
  const message = hasData ? <DataView /> : <EmptyState />;

  // 3. Logical AND (&&) — render if truthy
  // ⚠️ Use !! or convert to boolean to avoid rendering "0"
  const notification = count > 0 && <Badge count={count} />;
  // Safe version: {!!count && <Badge count={count} />}

  // 4. Nullish — return null to render nothing
  const badge = count > 0 ? <Badge /> : null;

  return (
    <div>
      {message}
      {notification}
      {badge}
    </div>
  );
}
```

---

### Q9. What is the difference between React and ReactDOM?

| Package | Purpose |
|---------|---------|
| `react` | Core library — components, hooks, JSX, Virtual DOM |
| `react-dom` | DOM-specific rendering — `createRoot`, `render`, `hydrate` |
| `react-native` | Native mobile rendering (instead of react-dom) |

```jsx
import React, { useState } from 'react';        // core — components, hooks
import ReactDOM from 'react-dom/client';          // DOM rendering
import { createPortal } from 'react-dom';         // DOM-specific utility

// The separation allows React to target different renderers:
// react-dom → browser DOM
// react-native → iOS/Android native
// react-three-fiber → Three.js/WebGL
// ink → terminal CLI
```

---

### Q10. What is `create-react-app` vs Vite?

| Feature | Create React App (CRA) | Vite |
|---------|------------------------|------|
| Build tool | Webpack | esbuild + Rollup |
| Dev server start | Slow (30s+) | **Instant** (<1s) |
| HMR | Slow | **Instant** |
| Bundle | Webpack bundle | ES Modules in dev |
| Maintained | ✅ | ✅ (preferred) |
| Config | Limited | Flexible |

```bash
# CRA (legacy — slower)
npx create-react-app my-app

# Vite (recommended)
npm create vite@latest my-app -- --template react
cd my-app && npm install && npm run dev

# Vite with TypeScript
npm create vite@latest my-app -- --template react-ts
```
