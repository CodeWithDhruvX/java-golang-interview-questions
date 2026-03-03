# 🚀 08 — Advanced React Topics
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- SSR vs CSR vs SSG vs ISR
- Next.js App Router fundamentals
- TypeScript with React — generics, utility types
- Animations (Framer Motion, CSS transitions)
- Web Workers in React
- Streaming and progressive hydration
- Performance monitoring (Sentry, DataDog RUM)

---

## ❓ Most Asked Questions

### Q1. What is the difference between SSR, CSR, SSG, and ISR?

| Rendering | When | Use Case | Tools |
|-----------|------|----------|-------|
| **CSR** (Client-Side) | In browser after JS loads | Dashboards, SPAs behind auth | CRA, Vite |
| **SSR** (Server-Side) | Per request on server | E-commerce PDPs, news sites | Next.js |
| **SSG** (Static Gen) | At build time | Blogs, docs, marketing pages | Next.js, Gatsby |
| **ISR** (Incremental Static) | Build + revalidate on demand | Product catalogs | Next.js |

```jsx
// Next.js App Router examples

// SSG — built at build time, cached at CDN edge
async function BlogPost({ params }) {
  const post = await getPost(params.slug); // runs at build time
  return <article>{post.content}</article>;
}

export async function generateStaticParams() {
  const posts = await getAllPosts();
  return posts.map(p => ({ slug: p.slug })); // pre-generate these routes
}

// SSR — run on server per request
async function Dashboard({ searchParams }) {
  const data = await db.query(`SELECT * FROM metrics`); // fresh per request
  return <MetricsView data={data} />;
}

// ISR — static but auto-revalidate every N seconds
async function ProductPage({ params }) {
  const product = await fetchProduct(params.id);
  return <ProductDetail product={product} />;
}

export const revalidate = 60; // regenerate every 60 seconds

// Client Component with CSR
'use client';
function InteractiveChart() {
  const [data, setData] = useState(null);
  useEffect(() => { fetchChartData().then(setData); }, []);
  return <Chart data={data} />;
}
```

---

### Q2. How do you use TypeScript with React effectively?

```tsx
// Typing components
interface Product {
  id: number;
  name: string;
  price: number;
  category: 'electronics' | 'clothing' | 'food';
  imageUrl?: string;
}

interface ProductCardProps {
  product: Product;
  onAddToCart: (id: number) => void;
  className?: string;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onAddToCart, className }) => {
  return (
    <div className={className}>
      <h3>{product.name}</h3>
      <p>${product.price}</p>
      <button onClick={() => onAddToCart(product.id)}>Add to Cart</button>
    </div>
  );
};

// Generic components
interface SelectProps<T extends { id: string | number; label: string }> {
  items: T[];
  value: T | null;
  onChange: (item: T) => void;
  placeholder?: string;
}

function Select<T extends { id: string | number; label: string }>({
  items, value, onChange, placeholder = 'Select...'
}: SelectProps<T>) {
  return (
    <select
      value={value?.id ?? ''}
      onChange={e => onChange(items.find(i => String(i.id) === e.target.value)!)}
    >
      {items.map(item => (
        <option key={item.id} value={item.id}>{item.label}</option>
      ))}
    </select>
  );
}

// Typing hooks
function useLocalStorage<T>(key: string, initial: T): [T, (val: T) => void] {
  const [value, setValue] = useState<T>(initial);
  const setAndStore = (val: T) => {
    setValue(val);
    localStorage.setItem(key, JSON.stringify(val));
  };
  return [value, setAndStore];
}

// useRef typing
const inputRef   = useRef<HTMLInputElement>(null);
const timerRef   = useRef<ReturnType<typeof setTimeout>>(null);
const countRef   = useRef<number>(0);
```

---

### Q3. How do you implement animations with Framer Motion?

```jsx
import { motion, AnimatePresence, useMotionValue, useTransform } from 'framer-motion';

// Basic animation
<motion.div
  initial={{ opacity: 0, y: -20 }}
  animate={{ opacity: 1, y: 0 }}
  exit={{ opacity: 0, y: 20 }}
  transition={{ duration: 0.3, ease: 'easeOut' }}
>
  Content
</motion.div>

// AnimatePresence — animates components on mount/unmount
function Toast({ message, isVisible }) {
  return (
    <AnimatePresence>
      {isVisible && (
        <motion.div
          className="toast"
          initial={{ opacity: 0, x: 100 }}
          animate={{ opacity: 1, x: 0 }}
          exit={{ opacity: 0, x: 100 }}
        >
          {message}
        </motion.div>
      )}
    </AnimatePresence>
  );
}

// Gesture-based animation
<motion.button
  whileHover={{ scale: 1.05 }}
  whileTap={{ scale: 0.95 }}
  transition={{ type: 'spring', stiffness: 300 }}
>
  Click me
</motion.button>

// Layout animation — smooth reordering
{items.map(item => (
  <motion.li key={item.id} layout>
    {item.name}
  </motion.li>
))}
```

---

### Q4. How do you handle real-time data in React?

```jsx
// WebSockets with custom hook
function useWebSocket(url) {
  const [messages, setMessages] = useState([]);
  const [status, setStatus] = useState('connecting');
  const ws = useRef(null);

  useEffect(() => {
    ws.current = new WebSocket(url);

    ws.current.onopen    = () => setStatus('connected');
    ws.current.onclose   = () => setStatus('disconnected');
    ws.current.onmessage = (e) => {
      const msg = JSON.parse(e.data);
      setMessages(prev => [...prev.slice(-99), msg]); // keep last 100
    };

    return () => ws.current?.close();
  }, [url]);

  const send = useCallback((msg) => {
    if (ws.current?.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(msg));
    }
  }, []);

  return { messages, status, send };
}

// Server-Sent Events (one-way streaming)
function useLiveUpdates(endpoint) {
  const [data, setData] = useState(null);

  useEffect(() => {
    const es = new EventSource(endpoint);
    es.onmessage = (e) => setData(JSON.parse(e.data));
    es.onerror   = () => es.close();
    return () => es.close();
  }, [endpoint]);

  return data;
}

// With React Query's refetch interval for polling
const { data } = useQuery({
  queryKey: ['live-price'],
  queryFn:  () => fetchPrice(),
  refetchInterval: 5000, // poll every 5 seconds
});
```

---

### Q5. What is streaming SSR and how does it improve performance?

```jsx
// Traditional SSR: server generates FULL HTML → send → browser renders
// Users wait for entire page before seeing anything

// Streaming SSR (React 18 + Next.js App Router):
// Server streams HTML progressively — browser renders immediately,
// fills in as more HTML arrives

// Next.js App Router + Suspense enables streaming automatically
import { Suspense } from 'react';

// This page streams — Shell renders immediately
// Heavy components stream in as data resolves
export default function ProductPage() {
  return (
    <div>
      <ProductHeader />          {/* immediate — no data needed */}
      <Suspense fallback={<ProductDetailSkeleton />}>
        <ProductDetails />       {/* streams in when DB query resolves */}
      </Suspense>
      <Suspense fallback={<ReviewsSkeleton />}>
        <ProductReviews />       {/* streams in independently */}
      </Suspense>
    </div>
  );
}

async function ProductDetails() {
  const product = await db.products.findUnique(...); // server-side query
  return <Detail product={product} />;
}

// Benefits vs traditional SSR:
// ✅ TTFB improved — first byte arrives with shell immediately
// ✅ Progressive render — user sees content sooner
// ✅ Parallel data fetching — each Suspense boundary resolves independently
// ✅ No waterfall — server components can fetch in parallel
```

---

### Q6. How do you implement infinite scroll in React?

```jsx
// Intersection Observer API for infinite scroll
import { useInfiniteQuery } from '@tanstack/react-query';

function InfiniteProductList() {
  const {
    data,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    isLoading,
  } = useInfiniteQuery({
    queryKey: ['products'],
    queryFn: ({ pageParam = 1 }) =>
      fetch(`/api/products?page=${pageParam}`).then(r => r.json()),
    getNextPageParam: (lastPage) =>
      lastPage.hasMore ? lastPage.nextPage : undefined,
  });

  // Sentinel element — triggers fetch when visible
  const sentinelRef = useRef(null);

  useEffect(() => {
    if (!sentinelRef.current || !hasNextPage) return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) fetchNextPage();
      },
      { threshold: 0.1 }
    );

    observer.observe(sentinelRef.current);
    return () => observer.disconnect();
  }, [hasNextPage, fetchNextPage]);

  const products = data?.pages.flatMap(page => page.products) ?? [];

  return (
    <div>
      {isLoading && <GridSkeleton count={8} />}
      <div className="grid">
        {products.map(p => <ProductCard key={p.id} product={p} />)}
      </div>
      {/* Sentinel element at bottom of list */}
      <div ref={sentinelRef} style={{ height: 20 }}>
        {isFetchingNextPage && <Spinner />}
      </div>
      {!hasNextPage && <p>All products loaded</p>}
    </div>
  );
}
```
