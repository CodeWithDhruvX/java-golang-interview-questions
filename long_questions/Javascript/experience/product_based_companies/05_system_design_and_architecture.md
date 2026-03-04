# ⚡ 05 — System Design & Architecture with JavaScript
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Designing scalable frontend systems
- State management patterns (Flux/Redux)
- Micro-frontends architecture
- SSR vs CSR vs ISR vs SSG trade-offs
- WebSockets and real-time communication
- Service Workers and PWAs
- Browser caching strategies

---

## ❓ Most Asked Questions

### Q1. Design a real-time notification system in JavaScript.

```javascript
// Real-time notification system: WebSocket + EventEmitter + UI layer

// --- Server-side (Node.js) ---
const WebSocket = require('ws');
const wss = new WebSocket.Server({ port: 8080 });

// Map: userId → Set<WebSocket> (user can have multiple tabs)
const connections = new Map();

wss.on('connection', (ws, req) => {
    const userId = authenticate(req); // extract from token/cookie

    if (!connections.has(userId)) connections.set(userId, new Set());
    connections.get(userId).add(ws);

    ws.on('message', (data) => {
        const msg = JSON.parse(data);
        handleMessage(userId, msg);
    });

    ws.on('close', () => {
        connections.get(userId)?.delete(ws);
        if (connections.get(userId)?.size === 0) {
            connections.delete(userId); // cleanup
        }
    });

    // Heartbeat to detect dead connections
    ws.isAlive = true;
    ws.on('pong', () => { ws.isAlive = true; });
});

// Health check interval
setInterval(() => {
    wss.clients.forEach(ws => {
        if (!ws.isAlive) return ws.terminate();
        ws.isAlive = false;
        ws.ping();
    });
}, 30000);

// Send notification to specific user (all their connections)
function sendNotification(userId, notification) {
    const userSockets = connections.get(userId);
    if (!userSockets) return;

    const payload = JSON.stringify({ type: "NOTIFICATION", ...notification });
    userSockets.forEach(ws => {
        if (ws.readyState === WebSocket.OPEN) ws.send(payload);
    });
}

// --- Client-side ---
class NotificationManager {
    #ws;
    #reconnectDelay = 1000;
    #handlers = new Map();

    connect(url, token) {
        this.#ws = new WebSocket(`${url}?token=${token}`);

        this.#ws.onmessage = ({ data }) => {
            const msg = JSON.parse(data);
            this.#handlers.get(msg.type)?.forEach(h => h(msg));
        };

        this.#ws.onclose = () => {
            // Exponential backoff reconnect
            setTimeout(() => {
                this.#reconnectDelay = Math.min(this.#reconnectDelay * 2, 30000);
                this.connect(url, token);
            }, this.#reconnectDelay);
        };

        this.#ws.onopen = () => {
            this.#reconnectDelay = 1000; // reset on success
        };
    }

    on(type, handler) {
        if (!this.#handlers.has(type)) this.#handlers.set(type, new Set());
        this.#handlers.get(type).add(handler);
        return () => this.#handlers.get(type).delete(handler); // unsubscribe
    }
}
```

---

### Q2. Design a client-side state management library (mini Redux).

```javascript
// Redux-like store: unidirectional data flow, immutable state, pure reducers

function createStore(reducer, initialState, ...middlewares) {
    let state = initialState;
    const listeners = new Set();
    let dispatching = false;

    // Compose middlewares
    const store = {
        getState: () => state,

        dispatch(action) {
            if (dispatching) throw new Error("Reducers may not dispatch actions");
            
            try {
                dispatching = true;
                const prevState = state;
                state = reducer(state, action);

                // Only notify if state changed (reference equality for immutable state)
                if (state !== prevState) {
                    listeners.forEach(listener => listener());
                }
            } finally {
                dispatching = false;
            }
            return action;
        },

        subscribe(listener) {
            listeners.add(listener);
            return () => listeners.delete(listener); // returns unsubscribe
        },
    };

    // Apply middlewares (like redux applyMiddleware)
    const middlewareAPI = {
        getState: store.getState,
        dispatch: (action) => store.dispatch(action)
    };

    if (middlewares.length) {
        const chain = middlewares.map(m => m(middlewareAPI));
        const originalDispatch = store.dispatch;
        store.dispatch = chain.reduceRight(
            (next, middleware) => middleware(next),
            originalDispatch
        );
    }

    // Initialize state
    store.dispatch({ type: "@@INIT" });
    return store;
}

// Logger middleware
const logger = ({ getState }) => next => action => {
    console.group(action.type);
    console.log("Before:", getState());
    const result = next(action);
    console.log("After:", getState());
    console.groupEnd();
    return result;
};

// Thunk middleware (async actions)
const thunk = ({ dispatch, getState }) => next => action => {
    if (typeof action === 'function') {
        return action(dispatch, getState); // async action creator
    }
    return next(action);
};

// Usage
const store = createStore(
    (state = { count: 0 }, action) => {
        switch (action.type) {
            case 'INCREMENT': return { ...state, count: state.count + action.payload };
            case 'RESET': return { ...state, count: 0 };
            default: return state;
        }
    },
    undefined,
    thunk, logger
);

// Async thunk
const fetchUserCount = () => async (dispatch, getState) => {
    const data = await fetch('/api/count').then(r => r.json());
    dispatch({ type: 'INCREMENT', payload: data.count });
};

store.dispatch(fetchUserCount());
```

---

### Q3. Explain Service Workers and offline-first architecture.

```javascript
// Service Worker: proxy between browser and network (intercepts fetch)
// Runs in separate thread — no DOM access, persistent background

// sw.js (Service Worker)
const CACHE_NAME = 'app-v1';
const STATIC_ASSETS = [
    '/',
    '/index.html',
    '/app.js',
    '/app.css',
    '/logo.svg'
];

// Install: cache static assets
self.addEventListener('install', event => {
    event.waitUntil(
        caches.open(CACHE_NAME)
            .then(cache => cache.addAll(STATIC_ASSETS))
            .then(() => self.skipWaiting()) // activate immediately
    );
});

// Activate: clean old caches
self.addEventListener('activate', event => {
    event.waitUntil(
        caches.keys().then(keys =>
            Promise.all(
                keys.filter(k => k !== CACHE_NAME).map(k => caches.delete(k))
            )
        ).then(() => self.clients.claim()) // take control immediately
    );
});

// Fetch: intercept network requests
self.addEventListener('fetch', event => {
    const { request } = event;

    // Different strategies based on URL
    if (request.url.includes('/api/')) {
        // Network-first for API (fresh data preferred)
        event.respondWith(networkFirst(request));
    } else if (request.destination === 'image') {
        // Cache-first for images (performance)
        event.respondWith(cacheFirst(request));
    } else {
        // Stale-while-revalidate for HTML/CSS/JS
        event.respondWith(staleWhileRevalidate(request));
    }
});

async function networkFirst(request) {
    try {
        const response = await fetch(request);
        const cache = await caches.open(CACHE_NAME);
        cache.put(request, response.clone());
        return response;
    } catch {
        return caches.match(request); // fallback to cache if offline
    }
}

async function cacheFirst(request) {
    const cached = await caches.match(request);
    if (cached) return cached;
    const response = await fetch(request);
    const cache = await caches.open(CACHE_NAME);
    cache.put(request, response.clone());
    return response;
}

async function staleWhileRevalidate(request) {
    const cached = caches.match(request);
    const networkFetch = fetch(request).then(response => {
        caches.open(CACHE_NAME).then(cache => cache.put(request, response.clone()));
        return response;
    });
    return (await cached) || networkFetch; // return cached immediately, update in background
}

// Registration (main.js)
if ('serviceWorker' in navigator) {
    navigator.serviceWorker.register('/sw.js')
        .then(reg => console.log('SW registered:', reg.scope))
        .catch(err => console.error('SW failed:', err));
}
```

---

### Q4. Design a micro-frontend architecture.

```javascript
// Micro-frontends: independently deployable frontend modules
// Approaches: iframes, Web Components, Module Federation (Webpack 5)

// Approach: Module Federation (Webpack 5) — most production-used

// shell/webpack.config.js (host)
module.exports = {
    plugins: [
        new ModuleFederationPlugin({
            name: "shell",
            remotes: {
                // Remote MFEs loaded at runtime from different deployments
                checkout: "checkout@https://checkout.example.com/remoteEntry.js",
                catalog:  "catalog@https://catalog.example.com/remoteEntry.js",
                auth:     "auth@https://auth.example.com/remoteEntry.js",
            },
            shared: {
                react:     { singleton: true, requiredVersion: "^18" },
                "react-dom": { singleton: true, requiredVersion: "^18" }
            }
        })
    ]
};

// Shell app — dynamically loads MFEs
async function loadMicroFrontend(name) {
    try {
        const module = await import(name);
        return module.default;
    } catch (err) {
        console.error(`Failed to load MFE: ${name}`, err);
        return FallbackComponent; // graceful degradation
    }
}

// Communication between MFEs: Custom Events or shared event bus
// custom-events approach (decoupled)
class MicroFrontendBus {
    emit(event, data) {
        window.dispatchEvent(new CustomEvent(`mfe:${event}`, { detail: data }));
    }

    on(event, handler) {
        window.addEventListener(`mfe:${event}`, e => handler(e.detail));
        return () => window.removeEventListener(`mfe:${event}`, e => handler(e.detail));
    }
}

const bus = new MicroFrontendBus();

// Checkout MFE: tells shell about cart changes
bus.emit('cart:updated', { count: 3, total: 1500 });

// Shell listens and updates header
bus.on('cart:updated', ({ count }) => {
    document.querySelector('.cart-badge').textContent = count;
});
```

---

### Q5. Implement a browser-side LRU Cache with storage persistence.

```javascript
// LRU Cache: Least Recently Used eviction
// Implementation: Map (preserves insertion order) + doubly linked list (O(1) ops)

class LRUCache {
    #capacity;
    #cache = new Map();   // key → value (Map preserves insertion order)

    constructor(capacity) {
        this.#capacity = capacity;
    }

    get(key) {
        if (!this.#cache.has(key)) return -1;

        // Move to end (most recently used)
        const value = this.#cache.get(key);
        this.#cache.delete(key);
        this.#cache.set(key, value);
        return value;
    }

    put(key, value) {
        if (this.#cache.has(key)) {
            this.#cache.delete(key); // remove first to re-insert at end
        } else if (this.#cache.size >= this.#capacity) {
            // Delete LRU: first key in Map
            this.#cache.delete(this.#cache.keys().next().value);
        }
        this.#cache.set(key, value);
    }

    get size() { return this.#cache.size; }
}

// Persistent LRU Cache using localStorage
class PersistentLRUCache extends LRUCache {
    #storageKey;
    #ttl; // time-to-live in ms

    constructor(capacity, storageKey, ttl = 5 * 60 * 1000) {
        super(capacity);
        this.#storageKey = storageKey;
        this.#ttl = ttl;
        this.#restore();
    }

    put(key, value) {
        super.put(key, { value, timestamp: Date.now() });
        this.#persist();
    }

    get(key) {
        const entry = super.get(key);
        if (entry === -1) return -1;

        // Check TTL
        if (Date.now() - entry.timestamp > this.#ttl) {
            // Expired — delegate to parent with actual key to delete
            return -1;
        }
        return entry.value;
    }

    #persist() {
        try {
            const data = Object.fromEntries(this._entries());
            localStorage.setItem(this.#storageKey, JSON.stringify(data));
        } catch (e) {
            if (e.name === 'QuotaExceededError') {
                // Storage full — evict half and retry
                this.#evictHalf();
                this.#persist();
            }
        }
    }

    #restore() {
        const stored = localStorage.getItem(this.#storageKey);
        if (stored) {
            const data = JSON.parse(stored);
            for (const [k, v] of Object.entries(data)) {
                if (Date.now() - v.timestamp < this.#ttl) super.put(k, v);
            }
        }
    }
}
```

---

### Q6. SSR vs CSR vs SSG vs ISR — when to use what?

```javascript
// Client-Side Rendering (CSR) — React/Vite defaults
// ✅ Use when: highly interactive, user-specific data (dashboards, SPAs)
// ❌ Issues: poor SEO (blank HTML initially), slow First Contentful Paint
// Example: Gmail, Figma, Notion (after login)

// Server-Side Rendering (SSR) — Next.js getServerSideProps
// ✅ Use when: personalized + SEO-critical (e-commerce product pages, social feeds)
// ❌ Issues: server load per request, slower TTFB than SSG
export async function getServerSideProps({ req, params }) {
    const session = await getSession(req);
    const product = await db.products.findById(params.id);
    const price = await getPricingForUser(session.userId, params.id);
    return { props: { product, price } }; // personalized price
}

// Static Site Generation (SSG) — build-time rendering
// ✅ Use when: marketing pages, blogs, documentation (content doesn't change often)
// ❌ Issues: build time grows with content; stale until rebuild
export async function getStaticProps() {
    const posts = await fetchBlogPosts();
    return {
        props: { posts },
        revalidate: 3600 // ISR: regenerate every hour
    };
}

// Incremental Static Regeneration (ISR) — Next.js
// ✅ Best of SSG + SSR: static performance + freshness
// Pages regenerated in background after revalidation window
// Example: e-commerce catalog (not personalized but needs fresh pricing)
export async function getStaticProps({ params }) {
    const product = await fetchProduct(params.slug);
    return {
        props: { product },
        revalidate: 60 // stale-while-revalidate: regenerate after 60s
    };
}

// Streaming SSR — React 18 Suspense + Next.js App Router
// ✅ Progressive hydration: fast initial HTML, stream deferred content
// page.js (App Router)
async function ProductPage({ params }) {
    const product = await fetchProduct(params.id); // fast
    return (
        <div>
            <ProductHeader product={product} />
            <Suspense fallback={<ReviewsSkeleton />}>
                <Reviews productId={params.id} /> {/* streamed separately */}
            </Suspense>
        </div>
    );
}
```

---

### Q7. Design a front-end rate limiter and request queue.

```javascript
// Rate limiter for API calls: max N requests per M milliseconds

class APIClient {
    #baseUrl;
    #queue = [];
    #running = 0;
    #maxConcurrent;
    #minInterval; // minimum ms between requests
    #lastRequestTime = 0;

    constructor(baseUrl, { maxConcurrent = 5, rps = 10 } = {}) {
        this.#baseUrl = baseUrl;
        this.#maxConcurrent = maxConcurrent;
        this.#minInterval = 1000 / rps; // requests per second → interval
    }

    async request(endpoint, options = {}) {
        return new Promise((resolve, reject) => {
            this.#queue.push({ endpoint, options, resolve, reject });
            this.#process();
        });
    }

    async #process() {
        if (this.#running >= this.#maxConcurrent || !this.#queue.length) return;

        const now = Date.now();
        const wait = Math.max(0, this.#lastRequestTime + this.#minInterval - now);

        if (wait > 0) {
            setTimeout(() => this.#process(), wait);
            return;
        }

        const { endpoint, options, resolve, reject } = this.#queue.shift();
        this.#running++;
        this.#lastRequestTime = Date.now();

        try {
            const response = await fetch(`${this.#baseUrl}${endpoint}`, options);
            if (!response.ok) throw new Error(`HTTP ${response.status}`);
            resolve(await response.json());
        } catch (err) {
            reject(err);
        } finally {
            this.#running--;
            this.#process(); // process next in queue
        }
    }

    // Batch multiple requests into one API call (request coalescing)
    #pendingBatch = new Map();
    #batchTimer;

    batchGet(id) {
        return new Promise((resolve, reject) => {
            this.#pendingBatch.set(id, { resolve, reject });
            clearTimeout(this.#batchTimer);
            this.#batchTimer = setTimeout(() => this.#flushBatch(), 10); // 10ms window
        });
    }

    async #flushBatch() {
        const batch = new Map(this.#pendingBatch);
        this.#pendingBatch.clear();

        const ids = [...batch.keys()];
        try {
            const results = await this.request(`/batch?ids=${ids.join(",")}`);
            results.forEach(item => batch.get(item.id)?.resolve(item));
        } catch (err) {
            batch.forEach(({ reject }) => reject(err));
        }
    }
}
```
