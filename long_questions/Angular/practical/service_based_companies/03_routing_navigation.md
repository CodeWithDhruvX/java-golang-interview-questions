# 📘 03 — Routing & Navigation
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Setting up `RouterModule` with routes
- `routerLink` and `routerLinkActive`
- Route parameters (`ActivatedRoute`)
- Route guards (`CanActivate`, `CanDeactivate`)
- Child routes and nested routing
- Lazy loading feature modules

---

## ❓ Most Asked Questions

### Q1. How do you set up routing in Angular?

```typescript
// app-routing.module.ts
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'products', component: ProductListComponent },
  { path: 'products/:id', component: ProductDetailComponent },
  { path: 'admin', loadChildren: () =>
      import('./admin/admin.module').then(m => m.AdminModule) },  // lazy load
  { path: '**', component: PageNotFoundComponent }  // wildcard — must be last
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
```

```html
<!-- app.component.html — required outlet -->
<nav>
  <a routerLink="/home" routerLinkActive="active">Home</a>
  <a routerLink="/products" routerLinkActive="active">Products</a>
</nav>
<router-outlet></router-outlet>
```

---

### Q2. How do you read route parameters?

```typescript
// Route: { path: 'products/:id', component: ProductDetailComponent }

@Component({ template: '<p>Product: {{ product?.name }}</p>' })
export class ProductDetailComponent implements OnInit, OnDestroy {
  product: Product;
  private sub: Subscription;

  constructor(
    private route: ActivatedRoute,
    private productService: ProductService
  ) {}

  ngOnInit(): void {
    // Option 1: snapshot (for route params that never change while component is alive)
    const id = this.route.snapshot.paramMap.get('id');

    // Option 2: Observable (recommended — reacts to param changes without re-rendering)
    this.sub = this.route.paramMap.pipe(
      map(params => params.get('id')),
      switchMap(id => this.productService.getProduct(id!))
    ).subscribe(p => this.product = p);
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }
}
```

---

### Q3. How do you pass query parameters?

```typescript
// Programmatic navigation with query params
constructor(private router: Router) {}

search(term: string): void {
  this.router.navigate(['/products'], {
    queryParams: { q: term, page: 1 },
    queryParamsHandling: 'merge'  // merge with existing params
  });
}
```

```html
<!-- Template link with query params -->
<a [routerLink]="['/products']" [queryParams]="{ category: 'electronics', sort: 'price' }">
  Electronics
</a>
```

```typescript
// Reading query params
export class ProductListComponent implements OnInit {
  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.route.queryParamMap.subscribe(params => {
      const category = params.get('category');
      const sort = params.get('sort');
      this.loadProducts(category, sort);
    });
  }
}
```

---

### Q4. What are Route Guards? Explain `CanActivate`.

Route guards control **navigation access** to routes.

```typescript
// auth.guard.ts
@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree {
    if (this.authService.isLoggedIn()) {
      return true;  // allow navigation
    }
    // Redirect to login with returnUrl
    return this.router.createUrlTree(['/login'], {
      queryParams: { returnUrl: state.url }
    });
  }
}

// Apply guard to routes
const routes: Routes = [
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: 'admin', component: AdminComponent, canActivate: [AuthGuard, AdminGuard] },
];
```

**Available route guards:**

| Guard | Purpose |
|-------|---------|
| `CanActivate` | Can the user navigate TO a route? |
| `CanActivateChild` | Can the user navigate to a child route? |
| `CanDeactivate` | Can the user navigate AWAY from a route? (e.g., unsaved changes) |
| `CanLoad` | Can a lazy-loaded module be loaded? |
| `Resolve` | Pre-fetch data before activating the route |

---

### Q5. What is lazy loading and how does it improve performance?

**Lazy loading** defers loading of feature modules until the user navigates to a route associated with them. This reduces the **initial bundle size** and speeds up initial page load.

```typescript
// app-routing.module.ts
const routes: Routes = [
  // Eager loading (default) — loaded immediately on app start
  { path: 'home', component: HomeComponent },

  // Lazy loading — module is fetched only when user visits /admin
  {
    path: 'admin',
    loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule)
  },
  {
    path: 'reports',
    loadChildren: () => import('./reports/reports.module').then(m => m.ReportsModule)
  },
];
```

```typescript
// admin/admin-routing.module.ts
const adminRoutes: Routes = [
  { path: '', component: AdminDashboardComponent },
  { path: 'users', component: UserManagementComponent },
  { path: 'settings', component: SettingsComponent },
];

@NgModule({
  imports: [RouterModule.forChild(adminRoutes)],  // forChild for feature modules
  exports: [RouterModule]
})
export class AdminRoutingModule {}
```

> **Interview tip:** Always mention lazy loading when asked about Angular performance optimization.

---

### Q6. What is the difference between `routerLink` and `router.navigate()`?

| Aspect | `routerLink` | `router.navigate()` |
|--------|-------------|---------------------|
| Usage | In HTML templates | In TypeScript code |
| Type | Directive | Method on `Router` service |
| When to use | Static links in nav bars | Dynamic navigation (after form submit, API call) |

```html
<!-- Template — use routerLink -->
<a [routerLink]="['/products', product.id]">View Product</a>
<a routerLink="/home">Home</a>
```

```typescript
// Code — use router.navigate()
onLoginSuccess(): void {
  const returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/dashboard';
  this.router.navigate([returnUrl]);
}

onProductClick(id: string): void {
  this.router.navigate(['/products', id], { queryParams: { ref: 'list' } });
}
```

---

### Q7. What is a `Resolve` guard?

```typescript
// product.resolver.ts — pre-fetch data before rendering the component
@Injectable({ providedIn: 'root' })
export class ProductResolver implements Resolve<Product> {
  constructor(private productService: ProductService) {}

  resolve(route: ActivatedRouteSnapshot): Observable<Product> {
    const id = route.paramMap.get('id')!;
    return this.productService.getProduct(id).pipe(
      catchError(() => EMPTY)  // EMPTY completes without emitting = navigation cancelled
    );
  }
}

// Attach to route
const routes: Routes = [
  {
    path: 'products/:id',
    component: ProductDetailComponent,
    resolve: { product: ProductResolver }
  }
];

// In the component — no loading state needed!
export class ProductDetailComponent implements OnInit {
  product: Product;

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.product = this.route.snapshot.data['product'];
  }
}
```

---

### Q8. What are child routes?

```typescript
const routes: Routes = [
  {
    path: 'products',
    component: ProductsLayoutComponent,  // parent with <router-outlet>
    children: [
      { path: '', component: ProductListComponent },
      { path: 'new', component: ProductFormComponent },
      { path: ':id', component: ProductDetailComponent },
      { path: ':id/edit', component: ProductFormComponent },
    ]
  }
];
```

```html
<!-- products-layout.component.html -->
<div class="products-page">
  <aside class="sidebar">
    <a routerLink="new">Add Product</a>
  </aside>
  <main>
    <router-outlet></router-outlet>  <!-- child components render here -->
  </main>
</div>
```
