# 🟡 HTTP, Observables & RxJS

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Wipro, Cognizant): HttpClient basics, error handling, interceptors
> - 🚀 **Product-Based** (Razorpay, PhonePe, Zomato): RxJS operators, cancellation, retry strategies, real-time
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🔹 HTTP & API Integration

---

### 1. How to make HTTP requests in Angular? 🟢 | 🏭

"Angular provides the `HttpClient` service in `@angular/common/http` for all HTTP operations.

```typescript
// Import HttpClientModule
@NgModule({ imports: [HttpClientModule] })

// Service
@Injectable({ providedIn: 'root' })
export class ProductService {
  private readonly API = 'https://api.example.com/products';

  constructor(private http: HttpClient) {}

  getAll(): Observable<Product[]> {
    return this.http.get<Product[]>(this.API);
  }

  getById(id: number): Observable<Product> {
    return this.http.get<Product>(`${this.API}/${id}`);
  }

  create(product: Partial<Product>): Observable<Product> {
    return this.http.post<Product>(this.API, product);
  }

  update(id: number, product: Partial<Product>): Observable<Product> {
    return this.http.put<Product>(`${this.API}/${id}`, product);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.API}/${id}`);
  }
}
```

All methods return `Observable<T>` — the request only fires when something **subscribes** to it."

#### In Depth
`HttpClient` methods are **cold observables** — the HTTP call doesn't start until you subscribe. Each subscription triggers a new HTTP request. If a component subscribes multiple times (e.g., via `async` pipe in multiple places in the template), the same API is called multiple times. Use `shareReplay(1)` to multicaste a single request to multiple subscribers: `return this.http.get<Product[]>(API).pipe(shareReplay(1))`.

---

### 2. What is HttpClientModule? 🟢 | 🏭

"`HttpClientModule` is an Angular module that provides the `HttpClient` service along with interceptor infrastructure, response parsing, and built-in request/response transformation.

In Angular 14 and below, you import it in your `AppModule`:
```typescript
import { HttpClientModule } from '@angular/common/http';
@NgModule({ imports: [HttpClientModule] })
```

In Angular 15+, use the functional `provideHttpClient()` instead:
```typescript
// main.ts
bootstrapApplication(AppComponent, {
  providers: [provideHttpClient(withInterceptors([authInterceptor]))]
});
```

Without `HttpClientModule` (or `provideHttpClient`), injecting `HttpClient` into a service throws a DI error."

#### In Depth
Under the hood, `HttpClient` uses `XmlHttpRequest` for browser-side requests. In Angular Universal (SSR), the `@angular/platform-server` package substitutes a Node.js-based HTTP implementation automatically. This means the same `HttpClient` code works in both browser and server environments without any changes — a key benefit of Angular's dependency injection abstraction.

---

### 3. How to handle errors in HTTP calls? 🟡 | 🏭🚀

"I use RxJS `catchError` operator in the service layer to intercept errors:

```typescript
import { catchError, throwError } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';

getProduct(id: number): Observable<Product> {
  return this.http.get<Product>(`${this.API}/${id}`).pipe(
    catchError((error: HttpErrorResponse) => {
      if (error.status === 404) {
        return throwError(() => new Error('Product not found'));
      }
      if (error.status === 401) {
        this.router.navigate(['/login']);
        return throwError(() => new Error('Unauthorized'));
      }
      if (!navigator.onLine) {
        return throwError(() => new Error('No internet connection'));
      }
      return throwError(() => new Error('Something went wrong. Try again.'));
    })
  );
}
```

For **global** error handling across all HTTP calls, I use an `HttpInterceptor` rather than repeating error logic in every service."

#### In Depth
`throwError(() => new Error(...))` is the correct modern RxJS 7+ syntax. The older `throwError(new Error(...))` is deprecated because it threw the error immediately in some schedulers. By using a factory function `() => new Error(...)`, the error is only created when actually needed. For retry strategies, I use `retry({ count: 3, delay: 1000 })` to automatically re-attempt failed requests before surfacing the error.

---

### 4. What is the use of HttpInterceptor? 🟡 | 🏭🚀

"An **`HttpInterceptor`** intercepts outgoing HTTP requests and incoming responses globally — without modifying each individual service call.

Common uses:
1. **Auth tokens** — Attach JWT to every request
2. **Loading indicators** — Show/hide global spinner
3. **Error handling** — Centralized error management
4. **Logging** — Log all requests in development
5. **Caching** — Cache GET responses

```typescript
@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(private auth: AuthService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = this.auth.getToken();

    if (token) {
      const authReq = req.clone({
        setHeaders: { Authorization: `Bearer ${token}` }
      });
      return next.handle(authReq);
    }

    return next.handle(req);
  }
}

// Registration (module-based):
{ provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }
```"

#### In Depth
HTTP requests are **immutable** — you cannot modify `req` directly. `req.clone()` creates a modified copy. The `multi: true` in the provider config allows **multiple interceptors** to coexist — Angular chains them in order. The interceptor chain runs in registration order for requests (first registered = outermost), and in reverse order for responses. So a logging interceptor registered first logs both the first outgoing request and the last incoming response.

---

### 5. How to handle request headers? 🟢 | 🏭

"I add headers using `HttpHeaders` and pass them as options:

```typescript
getSecuredData(): Observable<Data[]> {
  const headers = new HttpHeaders({
    'Authorization': `Bearer ${this.auth.getToken()}`,
    'Content-Type': 'application/json',
    'X-Custom-Header': 'my-value'
  });

  return this.http.get<Data[]>(this.API, { headers });
}
```

For app-wide headers (auth tokens, correlation IDs), I use an interceptor instead of repeating the `headers` option in every call. Interceptors centralize the concern and keep services clean."

#### In Depth
`HttpHeaders` is **immutable** — methods like `.set()`, `.append()`, and `.delete()` return a new `HttpHeaders` instance. This is because HTTP requests are processed asynchronously and immutability prevents race conditions where one part of the code modifies headers that another part is using. The same applies to `HttpParams` for query string parameters.

---

## 🔹 Observables & RxJS

---

### 6. What is an Observable? 🟢 | 🏭🚀

"An **Observable** is a lazy, cancellable stream of data over time. It's the core abstraction of **RxJS** — Angular's reactive programming library.

```typescript
import { Observable } from 'rxjs';

// Creating an observable
const numbers$ = new Observable<number>(subscriber => {
  subscriber.next(1);
  subscriber.next(2);
  subscriber.next(3);
  subscriber.complete();
});

// Subscribing
numbers$.subscribe({
  next: (value) => console.log(value),    // 1, 2, 3
  error: (err) => console.error(err),
  complete: () => console.log('Done!')
});
```

Observables are **lazy** — nothing happens until you subscribe. Multiple subscribers each get their own execution (cold observables)."

#### In Depth
Angular uses observables throughout: `HttpClient`, `ActivatedRoute.params`, `FormControl.valueChanges`, `Router.events` — all are observables. Understanding observables is **non-optional** in Angular. The key insight is the **Observer pattern**: producers emit values, operators transform them, and subscribers consume them. This separation makes complex async orchestration composable and readable.

---

### 7. Difference between Observable and Promise? 🟡 | 🏭🚀

| Aspect | Observable | Promise |
|---|---|---|
| Values | Multiple values over time | Single value |
| Laziness | Lazy (starts on subscribe) | Eager (starts immediately) |
| Cancellable | ✅ Yes (unsubscribe) | ❌ No |
| Operators | Rich (map, filter, retry...) | Limited (.then, .catch) |
| Retry | `retry(3)` | Must re-create |
| Used for | Streams, events, WebSockets | One-off async calls |

"I use **Promises** for one-shot async operations (like `async/await` patterns in Node.js). In Angular, I use **Observables** because they integrate with Angular's change detection, they're cancellable (preventing memory leaks when components unmount), and they compose powerfully with RxJS operators like `switchMap`, `debounceTime`, and `catchError`."

#### In Depth
Converting between them is straightforward:
- Observable → Promise: `firstValueFrom(observable$)` or `observable$.toPromise()` (deprecated)
- Promise → Observable: `from(promise)`

The killer advantage of Observables for Angular is **cancellation**. When a component is destroyed, calling `subscription.unsubscribe()` cancels in-flight HTTP requests via `HttpClient` (which supports AbortController under the hood in modern Angular). With promises, the HTTP call completes even after the component is gone, potentially causing errors when trying to update destroyed components.

---

### 8. What is a Subject? 🟡 | 🏭🚀

"A **Subject** is both an Observable (can be subscribed to) and an Observer (can receive and emit values). It's like an event bus — multiple subscribers receive the same emissions.

```typescript
const subject = new Subject<string>();

// Subscribe
subject.subscribe(v => console.log('Sub 1:', v));
subject.subscribe(v => console.log('Sub 2:', v));

// Emit
subject.next('Hello'); // Both subs receive: 'Hello'
subject.next('World'); // Both subs receive: 'World'
```

I use subjects as:
- **Event buses** between unrelated components (via a shared service)
- **Trigger signals** for operations like `takeUntil(destroy$)`
- **Bridges** between imperative code and reactive streams"

#### In Depth
Subjects are **hot observables** — late subscribers miss previous emissions. The variants address different needs:

- `BehaviorSubject<T>(initialValue)` — Replays the LAST value to new subscribers
- `ReplaySubject<T>(bufferSize)` — Replays the last N values to new subscribers
- `AsyncSubject<T>` — Emits only the last value, and only when `complete()` is called

For **shared state**, I use `BehaviorSubject` exposed as `.asObservable()` — this allows reading the current value synchronously (`.getValue()`) while preventing external code from calling `.next()` directly.

---

### 9. What is BehaviorSubject? 🟡 | 🏭🚀

"`BehaviorSubject` is a Subject that:
1. **Requires an initial value** in its constructor
2. **Emits the current value immediately** to new subscribers
3. Allows synchronous access to the current value via `.getValue()`

```typescript
// In a shared service
private userState$ = new BehaviorSubject<User | null>(null);

// Expose as read-only observable
readonly user$ = this.userState$.asObservable();

// Update (called after login)
setUser(user: User): void {
  this.userState$.next(user);
}

// Read current value synchronously
getCurrentUser(): User | null {
  return this.userState$.getValue();
}
```

```typescript
// In a component
this.authService.user$.subscribe(user => {
  // Immediately receives current user, then updates on change
  this.currentUser = user;
});
```

This is the **most common pattern** for shared state in Angular without a state management library like NgRx."

#### In Depth
I expose `BehaviorSubject` as `.asObservable()` to prevent external code from calling `.next()` directly — this enforces **unidirectional data flow**. Only the service that owns the state can update it. Components only subscribe and read. This is the same principle as NgRx's actions/reducers: state changes are controlled and traceable. `BehaviorSubject` is lightweight NgRx — perfect for small-to-medium apps.

---

### 10. What are common RxJS operators? 🟡 | 🏭🚀

"The most important operators I use daily:

**Transformation:**
- `map(fn)` — Transform each emitted value
- `switchMap(fn)` — Map to new observable, cancel previous (network calls)
- `mergeMap(fn)` — Map to new observable, keep all concurrent
- `concatMap(fn)` — Map to new observable, queue (sequential)
- `exhaustMap(fn)` — Map to new observable, ignore new while active

**Filtering:**
- `filter(predicate)` — Only pass values that match a condition
- `debounceTime(ms)` — Wait for silence before emitting (search input)
- `distinctUntilChanged()` — Skip duplicate consecutive values
- `take(n)` — Complete after n emissions
- `takeUntil(signal$)` — Complete when another observable emits

**Error Handling:**
- `catchError(fn)` — Handle errors gracefully
- `retry(n)` — Re-subscribe on error (n times)

**Combination:**
- `forkJoin([a$, b$])` — Emit when all complete (parallel calls)
- `combineLatest([a$, b$])` — Emit when any updates (with latest of all)
- `merge(a$, b$)` — Merge multiple streams

```typescript
// Real-world example: Search with debounce
this.searchControl.valueChanges.pipe(
  debounceTime(300),         // Wait for 300ms silence
  distinctUntilChanged(),    // Only if value changed
  filter(term => term.length >= 2), // Min 2 characters
  switchMap(term => this.api.search(term)), // Cancel previous request
  catchError(() => of([]))   // Return empty on error
).subscribe(results => this.results = results);
```"

#### In Depth
The choice between `switchMap`, `mergeMap`, `concatMap`, and `exhaustMap` is critical for correctness:

- **`switchMap`**: For **search** — cancel old requests when user types new query
- **`mergeMap`**: For **parallel operations** — run all concurrently (download multiple files)
- **`concatMap`**: For **sequential operations** — process one at a time (bank transfers)
- **`exhaustMap`**: For **submit buttons** — ignore new clicks while request is in-flight

Using the wrong one causes bugs: `mergeMap` with a search field can show outdated results if an older slower request completes after a newer one.

---

### 11. How to unsubscribe from an Observable? 🟡 | 🏭🚀

"Unmanaged subscriptions cause **memory leaks** — the observable keeps the component alive in memory even after it's destroyed. Several strategies:

**1. Manual Unsubscription (verbose):**
```typescript
private sub: Subscription;
ngOnInit() { this.sub = this.data$.subscribe(...); }
ngOnDestroy() { this.sub.unsubscribe(); }
```

**2. `takeUntil` + destroy Subject (recommended for multiple subscriptions):**
```typescript
private destroy$ = new Subject<void>();

ngOnInit() {
  this.data1$.pipe(takeUntil(this.destroy$)).subscribe(...);
  this.data2$.pipe(takeUntil(this.destroy$)).subscribe(...);
}

ngOnDestroy() {
  this.destroy$.next();
  this.destroy$.complete();
}
```

**3. `async` pipe (best for template subscriptions):**
```html
<div *ngIf="user$ | async as user">{{ user.name }}</div>
```
The `async` pipe automatically subscribes on init and unsubscribes on destroy — zero boilerplate.

**4. Angular 16+ `takeUntilDestroyed()` (cleanest):**
```typescript
constructor() {
  this.data$.pipe(takeUntilDestroyed()).subscribe(...);
}
```"

#### In Depth
The `async` pipe is the most elegant solution for template bindings — it eliminates manual subscription management entirely. However, using `async` pipe multiple times in the same template for the same observable causes multiple subscriptions, which means multiple HTTP calls. The fix is to use `*ngIf="data$ | async as data"` once to unwrap, then use `data` throughout the template. Or use `shareReplay(1)` on the observable to share the single subscription.

---

### 12. What is `switchMap` vs `mergeMap` in HTTP calls? 🔴 | 🚀

"`switchMap` and `mergeMap` both map to inner observables but handle concurrent requests differently:

```typescript
// switchMap — CANCELS previous request (perfect for search)
this.searchInput.valueChanges.pipe(
  debounceTime(300),
  switchMap(query => this.api.search(query))
  // If user types fast: 'a' → 'an' → 'ang'
  // Gets results for: 'ang' (previous requests cancelled)
).subscribe(results => this.results = results);

// mergeMap — KEEPS ALL concurrent requests
fromEvent(downloadBtns, 'click').pipe(
  mergeMap(id => this.api.downloadFile(id))
  // All downloads run in parallel
).subscribe(file => this.saveFile(file));

// exhaustMap — IGNORES new clicks while processing
fromEvent(submitBtn, 'click').pipe(
  exhaustMap(() => this.api.submitForm(this.form.value))
  // Ignores rapid double-clicks while waiting for response
).subscribe(() => this.onSuccess());
```"

#### In Depth
`switchMap` internally calls `unsubscribe()` on the previous inner observable when a new outer value arrives. For `HttpClient` observables in modern Angular, this translates to calling `AbortController.abort()` on the underlying `XmlHttpRequest`, which signals the browser to cancel the in-flight HTTP request. This not only prevents stale response handling in the component but also reduces server load by terminating unnecessary processing for cancelled requests.

---
