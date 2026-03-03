# 📘 05 — HTTP & RxJS
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- `HttpClient` for GET, POST, PUT, PATCH, DELETE
- Setting headers and request options
- `HttpInterceptor` for auth tokens and global error handling
- Observable basics: `subscribe`, `pipe`, `map`, `filter`
- Common RxJS operators: `switchMap`, `mergeMap`, `combineLatest`, `forkJoin`
- Error handling with `catchError`
- The `async` pipe in templates

---

## ❓ Most Asked Questions

### Q1. How do you make HTTP requests in Angular?

```typescript
// product.service.ts
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class ProductService {
  private readonly apiUrl = 'https://api.example.com/products';

  constructor(private http: HttpClient) {}

  // GET all with query params
  getProducts(category?: string, page: number = 1): Observable<ProductResponse> {
    let params = new HttpParams().set('page', page.toString());
    if (category) params = params.set('category', category);

    return this.http.get<ProductResponse>(this.apiUrl, { params });
  }

  // GET one by id
  getProduct(id: string): Observable<Product> {
    return this.http.get<Product>(`${this.apiUrl}/${id}`);
  }

  // POST
  createProduct(product: CreateProductDto): Observable<Product> {
    return this.http.post<Product>(this.apiUrl, product);
  }

  // PUT (full update)
  updateProduct(id: string, product: UpdateProductDto): Observable<Product> {
    return this.http.put<Product>(`${this.apiUrl}/${id}`, product);
  }

  // PATCH (partial update)
  patchProduct(id: string, changes: Partial<Product>): Observable<Product> {
    return this.http.patch<Product>(`${this.apiUrl}/${id}`, changes);
  }

  // DELETE
  deleteProduct(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }
}
```

```typescript
// component — subscribing
export class ProductListComponent implements OnInit, OnDestroy {
  products: Product[] = [];
  private sub: Subscription;

  constructor(private productService: ProductService) {}

  ngOnInit(): void {
    this.sub = this.productService.getProducts().subscribe({
      next: (res) => this.products = res.data,
      error: (err) => console.error('Failed to load products', err),
      complete: () => console.log('Done')
    });
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }
}
```

---

### Q2. What is `HttpInterceptor`? How do you implement one?

```typescript
// auth.interceptor.ts
@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(private authService: AuthService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = this.authService.getToken();

    // Clone the request and add the auth header
    const authReq = token
      ? req.clone({ setHeaders: { Authorization: `Bearer ${token}` } })
      : req;

    return next.handle(authReq).pipe(
      catchError((error: HttpErrorResponse) => {
        if (error.status === 401) {
          this.authService.logout();
          // redirect to login
        }
        return throwError(() => error);
      })
    );
  }
}

// Register in AppModule
@NgModule({
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: LoadingInterceptor, multi: true },
  ]
})
export class AppModule {}
```

---

### Q3. What is an Observable? How is it different from a Promise?

| Feature | Promise | Observable |
|---------|---------|-----------|
| Values emitted | One | Zero, one, or many |
| Lazy | No (executes immediately) | Yes (only when subscribed) |
| Cancellable | No | Yes (unsubscribe) |
| Operators | `.then()`, `.catch()` | Rich operators (`map`, `filter`, `switchMap`, etc.) |
| Multiple subscribers | No | Yes |

```typescript
// Promise — executes immediately, one value
const promise = fetch('/api/user');
promise.then(res => console.log(res));

// Observable — lazy, can emit multiple values, cancellable
const obs$ = this.http.get('/api/users');  // nothing happens yet
const sub = obs$.subscribe(data => console.log(data));  // starts now
sub.unsubscribe();  // cancel
```

---

### Q4. Explain common RxJS operators with examples.

```typescript
// map — transform each emitted value
this.productService.getProduct(id).pipe(
  map(product => product.name.toUpperCase())
).subscribe(name => console.log(name));

// filter — emit only values that pass a condition
this.products$.pipe(
  filter(p => p.price > 100)
).subscribe(expensiveProduct => console.log(expensiveProduct));

// switchMap — cancel previous inner observable, use for search (prevents stale results)
this.searchControl.valueChanges.pipe(
  debounceTime(300),
  distinctUntilChanged(),
  switchMap(term => this.productService.search(term))  // cancels previous search
).subscribe(results => this.results = results);

// mergeMap — run all inner observables concurrently
const productIds = ['1', '2', '3'];
from(productIds).pipe(
  mergeMap(id => this.productService.getProduct(id))
).subscribe(product => console.log(product));

// forkJoin — wait for ALL observables to complete, then emit combined result
forkJoin({
  user: this.userService.getCurrentUser(),
  products: this.productService.getFeatured(),
  categories: this.categoryService.getAll()
}).subscribe(({ user, products, categories }) => {
  this.setupDashboard(user, products, categories);
});

// combineLatest — emit when ANY source emits, combining latest from all
combineLatest([this.category$, this.sort$, this.page$]).pipe(
  switchMap(([category, sort, page]) =>
    this.productService.getProducts({ category, sort, page })
  )
).subscribe(products => this.products = products);

// catchError — handle errors gracefully
this.productService.getProducts().pipe(
  catchError(err => {
    console.error('Error:', err.message);
    return of([]);  // return empty array as fallback
  })
).subscribe(products => this.products = products);
```

---

### Q5. What is the `async` pipe and why is it preferred?

```typescript
// product-list.component.ts
export class ProductListComponent {
  // No manual subscribe/unsubscribe needed!
  products$ = this.productService.getProducts();
  user$ = this.userService.getCurrentUser();

  constructor(private productService: ProductService, private userService: UserService) {}
}
```

```html
<!-- The async pipe subscribes and unsubscribes automatically -->
<div *ngIf="products$ | async as products; else loading">
  <p>Found {{ products.length }} products</p>
  <div *ngFor="let product of products">{{ product.name }}</div>
</div>
<ng-template #loading><p>Loading...</p></ng-template>

<!-- Multiple async pipes with forkJoin -->
<ng-container *ngIf="{ user: user$ | async, products: products$ | async } as data">
  <h1>Welcome, {{ data.user?.name }}</h1>
  <p>{{ data.products?.length }} products</p>
</ng-container>
```

> **Always prefer `async` pipe** over manual subscription — it automatically unsubscribes when the component is destroyed, preventing memory leaks.

---

### Q6. How do you handle HTTP errors globally?

```typescript
// error.interceptor.ts
@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(private notificationService: NotificationService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      catchError((error: HttpErrorResponse) => {
        let userMessage = 'Something went wrong. Please try again.';

        switch (error.status) {
          case 400: userMessage = 'Invalid request. Please check your input.'; break;
          case 401: userMessage = 'Session expired. Please log in again.'; break;
          case 403: userMessage = 'You do not have permission to perform this action.'; break;
          case 404: userMessage = 'The requested resource was not found.'; break;
          case 500: userMessage = 'Server error. Our team has been notified.'; break;
        }

        this.notificationService.showError(userMessage);
        return throwError(() => new Error(userMessage));
      })
    );
  }
}
```

---

### Q7. What is `debounceTime` and why is it important for search?

```typescript
// Without debounce: API call on every keystroke (wasteful!)
// With debounce: API call only after user stops typing for 300ms

export class SearchComponent implements OnInit {
  searchControl = new FormControl('');

  ngOnInit(): void {
    this.searchControl.valueChanges.pipe(
      debounceTime(300),          // wait 300ms after last keystroke
      distinctUntilChanged(),     // don't emit if value hasn't changed
      filter(term => term!.length >= 2),  // minimum 2 chars
      switchMap(term =>           // cancel previous, start new
        this.searchService.search(term!).pipe(
          catchError(() => of([]))  // don't break search on error
        )
      )
    ).subscribe(results => this.results = results);
  }
}
```
