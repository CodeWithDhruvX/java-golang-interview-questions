# 📘 02 — Advanced RxJS
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Higher-order mapping operators: `switchMap`, `mergeMap`, `concatMap`, `exhaustMap`
- Subject types: `Subject`, `BehaviorSubject`, `ReplaySubject`, `AsyncSubject`
- Error handling strategies: `catchError`, `retry`, `retryWhen`
- Custom RxJS operators
- Multicasting: `share`, `shareReplay`
- `combineLatest` vs `withLatestFrom` vs `zip`

---

## ❓ Most Asked Questions

### Q1. `switchMap` vs `mergeMap` vs `concatMap` vs `exhaustMap` — when to use each?

This is the **most asked RxJS question** in product company interviews.

```typescript
// switchMap — CANCELS previous inner observable, starts new one
// Use for: search, autocomplete, navigation — where latest result wins
searchControl.valueChanges.pipe(
  debounceTime(300),
  switchMap(term => this.api.search(term))
  // If user types "ang" then "angu", api.search("ang") is CANCELLED
).subscribe(results => this.results = results);

// mergeMap (flatMap) — runs ALL inner observables CONCURRENTLY
// Use for: independent parallel requests, file uploads
from(['id1', 'id2', 'id3']).pipe(
  mergeMap(id => this.api.getProduct(id))
  // All 3 requests fire simultaneously — results come in as each completes
).subscribe(product => console.log(product));

// concatMap — waits for each inner to COMPLETE before starting next
// Use for: sequential operations, ordered processing, form submissions
from(pendingOrders).pipe(
  concatMap(order => this.api.processOrder(order))
  // Processes order1, waits for it, then order2, then order3 — in order
).subscribe(result => console.log(result));

// exhaustMap — IGNORES new emissions while inner observable is active
// Use for: login button (ignore double-clicks), form submit (prevent duplicate)
loginButton.clicks.pipe(
  exhaustMap(() => this.authService.login(credentials))
  // If user clicks again while login is in progress — click is IGNORED
).subscribe(user => this.onLoginSuccess(user));
```

| Operator | Behavior | Use Case |
|----------|----------|---------|
| `switchMap` | Cancel previous, start new | Search, routing, latest fetch wins |
| `mergeMap` | Concurrent, unordered | Parallel independent requests |
| `concatMap` | Sequential, ordered | Sequential order processing |
| `exhaustMap` | Ignore while busy | Login, form submit (prevent duplicates) |

---

### Q2. What are the Subject types in RxJS?

```typescript
// 1. Subject — no state, only forwards to current subscribers
const subject = new Subject<number>();
subject.subscribe(v => console.log('A:', v));  // subscribes before next
subject.next(1);  // A: 1
subject.next(2);  // A: 2
// Late subscriber MISSES past values:
subject.subscribe(v => console.log('B:', v));
subject.next(3);  // A: 3, B: 3 (B only sees 3)

// 2. BehaviorSubject — stores CURRENT value, late subscribers get it immediately
const behavior = new BehaviorSubject<number>(0);  // must have initial value
behavior.subscribe(v => console.log('A:', v));  // A: 0 (immediately!)
behavior.next(1);  // A: 1
behavior.subscribe(v => console.log('B:', v));  // B: 1 (gets current value!)
behavior.getValue(); // 1 — synchronously get current value

// 3. ReplaySubject — buffers N last values, late subscribers receive them
const replay = new ReplaySubject<number>(3);  // buffer last 3
replay.next(1);
replay.next(2);
replay.next(3);
replay.next(4);
replay.subscribe(v => console.log(v));  // prints: 2, 3, 4 (last 3)

// 4. AsyncSubject — only emits the LAST value when complete() is called
const async$ = new AsyncSubject<number>();
async$.next(1);
async$.next(2);
async$.next(3);
async$.subscribe(v => console.log(v));  // waiting...
async$.complete();  // now prints: 3 (last value at completion)
```

| Subject Type | Stores | Late Subscriber Gets |
|-------------|--------|---------------------|
| `Subject` | Nothing | Nothing (misses past) |
| `BehaviorSubject` | Current value | Current value |
| `ReplaySubject(n)` | Last n values | Last n values |
| `AsyncSubject` | Last value | Last value on complete |

---

### Q3. How do `combineLatest`, `withLatestFrom`, and `zip` differ?

```typescript
const a$ = interval(1000).pipe(map(i => `A${i}`));
const b$ = interval(1500).pipe(map(i => `B${i}`));

// combineLatest — emits when ANY source emits, using latest from all
// Use for: multiple filter dropdowns, form fields combining
combineLatest([category$, sort$, page$]).pipe(
  switchMap(([cat, sort, page]) => this.api.getProducts(cat, sort, page))
).subscribe();
// Emits each time category, sort, OR page changes

// withLatestFrom — primary triggers, secondary provides context
// Use for: user actions that need current state as context
this.saveButton.clicks.pipe(
  withLatestFrom(this.form.valueChanges),
  // Only fires on click; form value is the CONTEXT, not a trigger
  switchMap(([_, formValue]) => this.api.save(formValue))
).subscribe();

// zip — pairs emissions index-by-index, waits for ALL to emit
// Use for: coordinating parallel requests that must be paired
zip(
  this.api.getUser(id),    // emits once
  this.api.getOrders(id),  // emits once
  this.api.getCart(id)     // emits once
).subscribe(([user, orders, cart]) => {
  this.setupProfile(user, orders, cart);
});
// Waits for ALL three to emit, then combines them
```

---

### Q4. What is `shareReplay` and why is it important?

```typescript
// Problem: Each subscriber to an HTTP Observable triggers a NEW HTTP call
const product$ = this.http.get<Product>('/api/product/1');
product$.subscribe(p => this.productA = p);  // HTTP call 1
product$.subscribe(p => this.productB = p);  // HTTP call 2 — DUPLICATE!

// ✅ shareReplay — multicasts and caches the last N emissions
const product$ = this.http.get<Product>('/api/product/1').pipe(
  shareReplay(1)  // buffer 1 emission, share with all subscribers
);
product$.subscribe(p => this.productA = p);  // HTTP call fires
product$.subscribe(p => this.productB = p);  // gets CACHED result — no 2nd call!

// Real-world: caching a config or user profile
@Injectable({ providedIn: 'root' })
export class ConfigService {
  private config$ = this.http.get<AppConfig>('/api/config').pipe(
    shareReplay(1)  // loaded once, shared across all components
  );

  getConfig(): Observable<AppConfig> {
    return this.config$;
  }
}
```

---

### Q5. How do you create a custom RxJS operator?

```typescript
// Custom operator: retry with exponential backoff
export function retryWithBackoff<T>(maxRetries: number, initialDelay: number = 1000) {
  return (source: Observable<T>): Observable<T> => {
    return source.pipe(
      retryWhen(errors =>
        errors.pipe(
          // Pair each error with its retry count
          scan((retryCount, error) => {
            if (retryCount >= maxRetries) throw error;  // give up after max
            return retryCount + 1;
          }, 0),
          // Exponential delay: 1s, 2s, 4s, 8s...
          delayWhen(retryCount => timer(initialDelay * Math.pow(2, retryCount - 1))),
          tap(retryCount => console.log(`Retry attempt ${retryCount}`))
        )
      )
    );
  };
}

// Usage
this.http.get('/api/products').pipe(
  retryWithBackoff(3, 1000)  // retry 3 times: 1s, 2s, 4s
).subscribe(data => console.log(data));

// ----------------
// Custom operator: add loading state
export function withLoading<T>(loadingSubject: BehaviorSubject<boolean>) {
  return (source: Observable<T>): Observable<T> => {
    return defer(() => {
      loadingSubject.next(true);
      return source.pipe(
        finalize(() => loadingSubject.next(false))
      );
    });
  };
}

// Usage
this.dataService.getData().pipe(
  withLoading(this.loading$)
).subscribe(data => this.data = data);
```

---

### Q6. What is `forkJoin` and when does it fail?

```typescript
// forkJoin — waits for ALL observables to COMPLETE, then emits combined result
forkJoin({
  products: this.productService.getAll(),
  categories: this.categoryService.getAll(),
  user: this.userService.getCurrent()
}).subscribe({
  next: ({ products, categories, user }) => {
    this.setupApp(products, categories, user);
  },
  error: (err) => console.error('At least one request failed', err)
});

// ⚠️ Pitfalls:
// 1. If ANY source throws, forkJoin errors immediately (others are cancelled)
// 2. If ANY source doesn't complete (like a BehaviorSubject), forkJoin never emits
// 3. If a source emits ZERO values and completes, the result is undefined

// Solution: handle errors per request
forkJoin({
  products: this.productService.getAll().pipe(catchError(() => of([]))),
  user: this.userService.getCurrent().pipe(catchError(() => of(null)))
}).subscribe(({ products, user }) => { });
// Now even if one fails, forkJoin still completes with the fallback values
```
