# 🌍 07 — Frontend System Design
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Design systems and component libraries
- Micro-frontend architecture
- Accessibility (WCAG) and ARIA
- Internationalization (i18n)
- Performance architecture (CDN, caching)
- A/B testing framework
- Feature flags

---

## ❓ Most Asked Questions

### Q1. How do you design a scalable component library?

```
Design System Architecture:
┌─────────────────────────────────────────────┐
│               Design Tokens                 │
│  Colors, Typography, Spacing, Shadows       │
│  (CSS variables / JS constants)             │
└─────────────────┬───────────────────────────┘
                  │
┌─────────────────▼───────────────────────────┐
│           Primitive Components              │
│  Button, Input, Text, Icon, Box, Stack      │
│  (stateless, composable, accessible)        │
└─────────────────┬───────────────────────────┘
                  │
┌─────────────────▼───────────────────────────┐
│           Composite Components              │
│  Card, Modal, Dropdown, DatePicker          │
│  (built from primitives)                   │
└─────────────────┬───────────────────────────┘
                  │
┌─────────────────▼───────────────────────────┐
│           Page Templates                    │
│  DashboardLayout, AuthLayout, etc.          │
└─────────────────────────────────────────────┘
```

```jsx
// Design tokens — single source of truth
export const tokens = {
  colors: {
    primary:   { 500: '#2563eb', 600: '#1d4ed8', 700: '#1e40af' },
    neutral:   { 100: '#f5f5f5', 900: '#171717' },
    semantic:  { success: '#16a34a', error: '#dc2626', warning: '#d97706' }
  },
  spacing: { xs: '4px', sm: '8px', md: '16px', lg: '24px', xl: '32px' },
  typography: {
    fontFamily: "'Inter', sans-serif",
    sizes: { sm: '12px', md: '14px', lg: '16px', xl: '20px', '2xl': '24px' }
  }
};

// Compound variant pattern for scalable component APIs
function Button({ variant = 'solid', color = 'primary', size = 'md', ...props }) {
  const styles = buttonVariants({ variant, color, size });
  return <button className={styles} {...props} />;
}
```

---

### Q2. What are Micro-Frontends and when should you use them?

```
Micro-Frontends: split a large frontend into independently deployable chunks
Each owned by a separate team with separate deployments

Architecture options:
1. iFrame composition — simplest, worst UX (isolation issues)
2. Web Components — good isolation, complex routing
3. Module Federation (Webpack 5) — modern, shared dependencies
4. Single-SPA — runtime orchestration framework

           ┌──────────────────────────────────┐
           │          Shell App               │
           │  (routing, auth, shared layout)  │
           └──────┬──────────┬───────────────┘
                  │          │
           ┌──────▼──┐  ┌────▼──────┐
           │Catalog  │  │ Checkout  │
           │Team     │  │ Team      │
           └─────────┘  └───────────┘
```

```javascript
// webpack.config.js — Host (Shell) consuming remotes
const { ModuleFederationPlugin } = require('webpack').container;

new ModuleFederationPlugin({
  name: 'shell',
  remotes: {
    catalog:  'catalog@https://catalog.app.com/remoteEntry.js',
    checkout: 'checkout@https://checkout.app.com/remoteEntry.js',
  },
  shared: { react: { singleton: true }, 'react-dom': { singleton: true } }
});

// Remote (Catalog) exposing components
new ModuleFederationPlugin({
  name: 'catalog',
  filename: 'remoteEntry.js',
  exposes: {
    './ProductList': './src/components/ProductList',
    './ProductCard': './src/components/ProductCard',
  },
  shared: { react: { singleton: true } }
});

// Consuming in shell
const ProductList = React.lazy(() => import('catalog/ProductList'));
<Suspense fallback={<Skeleton />}><ProductList /></Suspense>

// When to use micro-frontends:
// ✅ Large orgs with multiple teams owning separate features
// ✅ Independent deployment cycle requirements
// ❌ Small teams — overhead not worth it
```

---

### Q3. How do you implement proper accessibility in React?

```jsx
// WCAG 2.1 AA — minimum standard for accessibility

// 1. Semantic HTML — screen reader support
// ❌ <div onClick> — not keyboard focusable!
// ✅ <button> — natively accessible
<button onClick={handleSubmit}>Submit</button>

// 2. ARIA — for custom interactive elements
function Dropdown({ items, isOpen, toggle }) {
  return (
    <div>
      <button
        aria-expanded={isOpen}
        aria-haspopup="listbox"
        aria-controls="dropdown-list"
        onClick={toggle}
      >
        Choose Option
      </button>
      <ul
        id="dropdown-list"
        role="listbox"
        aria-label="Options"
        hidden={!isOpen}
      >
        {items.map(item => (
          <li key={item.id} role="option" aria-selected={item.selected}>
            {item.label}
          </li>
        ))}
      </ul>
    </div>
  );
}

// 3. Focus management — trap focus in modals
import { FocusTrap } from 'focus-trap-react';
<FocusTrap active={isModalOpen}>
  <div role="dialog" aria-modal="true" aria-labelledby="modal-title">
    <h2 id="modal-title">Modal</h2>
    ...
  </div>
</FocusTrap>

// 4. Color contrast — minimum 4.5:1 for normal text (WCAG AA)

// 5. Screen reader announcements
function LiveRegion({ message }) {
  return (
    <div
      aria-live="polite"        // announce after current action
      aria-atomic="true"
      className="sr-only"        // visually hidden
    >
      {message}
    </div>
  );
}
```

---

### Q4. How do you implement internationalization (i18n)?

```jsx
import i18n from 'i18next';
import { initReactI18next, useTranslation } from 'react-i18next';

// Setup (typically in i18n.ts)
i18n.use(initReactI18next).init({
  resources: {
    en: { translation: { 'welcome': 'Welcome, {{name}}!', 'cart': 'Cart ({{count}} items)' } },
    hi: { translation: { 'welcome': 'स्वागत है, {{name}}!', 'cart': 'कार्ट ({{count}} आइटम)' } },
    fr: { translation: { 'welcome': 'Bienvenue, {{name}} !', 'cart': 'Panier ({{count}} articles)' } },
  },
  lng: 'en',
  fallbackLng: 'en',
  interpolation: { escapeValue: false }
});

// Usage in components
function Header({ userName, cartCount }) {
  const { t, i18n } = useTranslation();

  return (
    <header>
      <h1>{t('welcome', { name: userName })}</h1>
      <span>{t('cart', { count: cartCount })}</span>
      <select
        value={i18n.language}
        onChange={e => i18n.changeLanguage(e.target.value)}
      >
        <option value="en">English</option>
        <option value="hi">हिंदी</option>
        <option value="fr">Français</option>
      </select>
    </header>
  );
}

// RTL support for Arabic/Hebrew
useEffect(() => {
  document.dir = i18n.language === 'ar' ? 'rtl' : 'ltr';
}, [i18n.language]);
```

---

### Q5. How do you implement feature flags in React?

```jsx
// Feature flags: ship code to production hidden behind flags
// Enable for specific users, regions, A/B groups

// Simple flag system
const FLAGS = {
  NEW_CHECKOUT:    process.env.REACT_APP_FLAG_NEW_CHECKOUT === 'true',
  DARK_MODE_V2:    process.env.REACT_APP_FLAG_DARK_MODE === 'true',
  AI_SEARCH:       false, // off for everyone
};

// Context-based flag system
const FeatureFlagContext = createContext({});

function FeatureFlagProvider({ children }) {
  const { data: flags } = useQuery({
    queryKey: ['feature-flags'],
    queryFn: () => fetch('/api/flags').then(r => r.json()),
    staleTime: 5 * 60 * 1000,  // refresh every 5 minutes
  });

  return (
    <FeatureFlagContext.Provider value={flags || {}}>
      {children}
    </FeatureFlagContext.Provider>
  );
}

// Custom hook for flag checks
function useFeatureFlag(flag) {
  const flags = useContext(FeatureFlagContext);
  return Boolean(flags[flag]);
}

// Usage
function CheckoutPage() {
  const hasNewCheckout = useFeatureFlag('NEW_CHECKOUT');

  return hasNewCheckout ? <NewCheckout /> : <OldCheckout />;
}

// Gate component pattern
function FeatureGate({ flag, fallback = null, children }) {
  const isEnabled = useFeatureFlag(flag);
  return isEnabled ? children : fallback;
}

// Usage
<FeatureGate flag="AI_SEARCH" fallback={<BasicSearch />}>
  <AISearch />
</FeatureGate>
```
