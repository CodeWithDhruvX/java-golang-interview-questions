# 🚀 08 — Deployment & DevOps for React
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Building React for production
- Vite vs CRA production build
- Hosting options (Vercel, Netlify, AWS S3)
- Environment variables (dev vs prod)
- Docker for React apps
- CI/CD basics for React
- Performance audit and optimization before deployment

---

## ❓ Most Asked Questions

### Q1. How do you build and optimize a React app for production?

```bash
# CRA build
npm run build
# Creates /build folder with:
# - index.html (entry point)
# - static/js/*.chunk.js (code split bundles)
# - static/css/*.chunk.css
# - static/media/ (images, fonts)
# All assets are hashed: main.abc123.js (cache busting)

# Vite build
npm run build
# Creates /dist folder (faster, smaller bundles)

# Preview production build locally
npm run preview   # Vite
npx serve build   # CRA
```

```jsx
// What React does in production build:
// ✅ Removes PropTypes validation (smaller bundle)
// ✅ Removes React.StrictMode double-renders
// ✅ Minifies JS and CSS
// ✅ Tree-shakes unused code
// ✅ Enables production-mode optimizations in V8

// Code splitting for smaller initial bundle
const Dashboard = lazy(() => import('./pages/Dashboard'));
const Reports   = lazy(() => import('./pages/Reports'));

// Bundle analysis
npx vite-bundle-analyzer  # for Vite
npx webpack-bundle-analyzer  # for CRA (eject needed or use rewired)
```

---

### Q2. How do you deploy to Vercel and Netlify?

```bash
# --- Vercel (recommended for Next.js / Vite) ---
# 1. Push code to GitHub
# 2. Connect repo at vercel.com → Import project
# 3. Set environment variables in Vercel dashboard
# 4. Every push to main = automatic deployment

# CLI deployment:
npm install -g vercel
vercel login
vercel --prod    # deploy to production

# vercel.json (optional config)
{
  "builds": [{ "src": "package.json", "use": "@vercel/static-build" }],
  "routes": [{ "src": "/(.*)", "dest": "/index.html" }]  // SPA fallback
}

# --- Netlify ---
# 1. Connect GitHub repo at netlify.com
# 2. Build command: npm run build
# 3. Publish directory: dist (Vite) or build (CRA)

# netlify.toml
[build]
  command = "npm run build"
  publish = "dist"

[[redirects]]
  from = "/*"     # SPA: redirect all routes to index.html
  to = "/index.html"
  status = 200
```

---

### Q3. How do you deploy a React app with Docker?

```dockerfile
# Dockerfile — multi-stage build (smaller image)

# Stage 1: Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci               # clean install — uses package-lock.json
COPY . .
RUN npm run build

# Stage 2: Serve with nginx
FROM nginx:alpine AS production
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

```nginx
# nginx.conf — required for React Router (SPA fallback)
server {
  listen 80;

  location / {
    root   /usr/share/nginx/html;
    index  index.html;
    try_files $uri /index.html;  # fallback for client-side routing
  }

  # Cache static assets
  location ~* \.(js|css|png|jpg|ico)$ {
    expires 1y;
    add_header Cache-Control "public, immutable";
  }
}
```

```bash
# Build and run
docker build -t my-react-app .
docker run -p 3000:80 my-react-app
# App available at http://localhost:3000
```

---

### Q4. How do you set up a basic CI/CD pipeline for React?

```yaml
# .github/workflows/deploy.yml — GitHub Actions
name: CI/CD Pipeline

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'
      - run: npm ci
      - run: npm test -- --watchAll=false --coverage
      - run: npm run lint

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with: { node-version: '20', cache: 'npm' }
      - run: npm ci
      - run: npm run build
        env:
          VITE_API_URL: ${{ secrets.PROD_API_URL }}   # from GitHub Secrets
      - uses: actions/upload-artifact@v4
        with:
          name: build-output
          path: dist/          # or build/ for CRA

  deploy:
    needs: build
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/download-artifact@v4
        with: { name: build-output }
      - uses: amondnet/vercel-action@v25
        with:
          vercel-token: ${{ secrets.VERCEL_TOKEN }}
          vercel-org-id: ${{ secrets.VERCEL_ORG_ID }}
          vercel-project-id: ${{ secrets.VERCEL_PROJECT_ID }}
          vercel-args: '--prod'
```

---

### Q5. How do you use environment variables in different environments?

```bash
# File names (Vite auto-loads based on environment):
.env               # loaded always
.env.local         # loaded always, gitignored
.env.development   # loaded in dev (npm run dev)
.env.production    # loaded in prod (npm run build)
.env.test          # loaded in test (npm test)

# Example files:
# .env.development
VITE_API_URL=http://localhost:5000
VITE_SENTRY_DSN=
VITE_ENABLE_MOCK_DATA=true

# .env.production
VITE_API_URL=https://api.myapp.com
VITE_SENTRY_DSN=https://xxx@sentry.io/123
VITE_ENABLE_MOCK_DATA=false
```

```jsx
// Access in code
const API_URL = import.meta.env.VITE_API_URL;
const isDev   = import.meta.env.DEV;   // Vite built-in
const isProd  = import.meta.env.PROD;  // Vite built-in
const mode    = import.meta.env.MODE;  // 'development' | 'production' | 'test'

// Type safety with TypeScript (vite-env.d.ts)
// Already generated by Vite, you can extend:
interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  readonly VITE_SENTRY_DSN: string;
}
```

---

### Q6. What is a service worker and how does it make a React app a PWA?

```javascript
// public/sw.js — service worker (runs in background, separate thread)
const CACHE_NAME = 'my-app-v1';
const URLS_TO_CACHE = ['/', '/index.html', '/static/js/main.js', '/static/css/main.css'];

// Install — cache assets
self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => cache.addAll(URLS_TO_CACHE))
  );
});

// Fetch — serve from cache, fallback to network
self.addEventListener('fetch', event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
```

```jsx
// Register service worker in React (main.jsx / index.js)
if ('serviceWorker' in navigator && process.env.NODE_ENV === 'production') {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/sw.js')
      .then(reg => console.log('SW registered:', reg.scope))
      .catch(err => console.log('SW registration failed:', err));
  });
}

// PWA requirements (manifest.json in /public):
// {
//   "name": "My React App",
//   "short_name": "MyApp",
//   "icons": [...],
//   "start_url": "/",
//   "display": "standalone",
//   "theme_color": "#007bff",
//   "background_color": "#ffffff"
// }

// Vite PWA plugin (recommended)
import { VitePWA } from 'vite-plugin-pwa';
// Auto-generates service worker and manifest
```

---

### Q7. What performance checks should you do before deployment?

```bash
# 1. Lighthouse audit (Chrome DevTools → Lighthouse tab)
# Target: Performance > 90, Accessibility > 90, Best Practices > 90, SEO > 90

# 2. Bundle size check
npx vite-bundle-visualizer        # visual bundle map
# Target: initial JS bundle < 150KB gzipped

# 3. Check for console errors/warnings
# Open browser DevTools Console with production build

# 4. Test on slower network
# Chrome DevTools → Network → Throttle to "Slow 3G"
# Check LCP, loading states, skeleton screens

# 5. Test on real mobile device or Chrome DevTools device simulation

# 6. Check Web Vitals
npx web-vitals-reporter
```

```jsx
// Code-level performance checklist:
// ✅ Lazy load routes
// ✅ Compress images (WebP format)
// ✅ Preload critical fonts
// ✅ Virtualize long lists
// ✅ Memoize expensive calculations
// ✅ Remove unused dependencies

// Check unused imports:
npx depcheck   // finds unused npm packages
```
