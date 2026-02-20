# 🟡 Routing & Navigation

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Infosys, Wipro): Definitions, configuration, standard patterns
> - 🚀 **Product-Based** (Google, Amazon, Uber): Guards, lazy loading, performance, edge cases
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

### 1. What is Angular Router? 🟢 | 🏭

"The **Angular Router** is a built-in routing library that maps URL paths to components, enabling **Single Page Application (SPA)** navigation without full page reloads.

It handles:
- URL-to-component mapping
- Programmatic navigation
- Route guards (authentication, authorization)
- Lazy loading of feature modules
- Route parameters and query strings
- Navigation lifecycle events

Without the router, a user refreshing the URL would always land on the root component with no way to deep-link into the app."

#### In Depth
The Angular Router uses the **browser's History API** (`pushState`) by default to maintain clean URLs. The alternative — **hash-based routing** (`useHash: true`) — uses URL fragments like `/#/products`. HashLocationStrategy is useful when deploying to static servers that don't support URL rewriting, but `PathLocationStrategy` (default) requires the server to redirect all routes to `index.html` to support deep linking.

---

### 2. How to configure routes in Angular? 🟢 | 🏭

"Routes are configured as an array of `Route` objects and registered via `RouterModule.forRoot()` in the app:

```typescript
// app-routing.module.ts
const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'products', component: ProductListComponent },
  { path: 'products/:id', component: ProductDetailComponent },
  { path: 'admin', loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule) },
  { path: '**', component: NotFoundComponent }, // Wildcard — must be last
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
```

The `**` wildcard must always be the **last route** — Angular uses **first-match wins** routing."

#### In Depth
`pathMatch: 'full'` on redirects is critical. Without it, `path: ''` matches every URL (since every URL starts with an empty string), causing all routes to redirect to `home`. `pathMatch: 'full'` ensures the path must fully match the empty string — i.e., only the root URL `/`. This is a common gotcha for developers new to Angular routing.

---

### 3. What are router-outlet and routerLink? 🟢 | 🏭

"**`<router-outlet>`** is a placeholder directive where the Router renders the matched component.

**`routerLink`** is a directive for navigation links — the Angular equivalent of `<a href>`.

```html
<!-- app.component.html -->
<nav>
  <a routerLink="/home" routerLinkActive="active">Home</a>
  <a routerLink="/products" routerLinkActive="active">Products</a>
</nav>

<router-outlet></router-outlet> <!-- Matched component renders here -->
```

`routerLinkActive="active"` automatically adds the CSS class when that route is active. I use it to highlight the current menu item without any manual logic."

#### In Depth
`routerLink` prevents the default anchor tag behavior (`href` navigation) and uses the Router's navigation instead, preserving in-memory state and avoiding page reloads. The `routerLinkActiveOptions="{ exact: true }"` is important for the root path — without `exact: true`, the `/` link would appear active on `/products` too (since both start with `/`).

---

### 4. How to implement child routes? 🟡 | 🏭

"Child routes allow **nested navigation** within a parent component. The parent must have its own `<router-outlet>` where children render.

```typescript
const routes: Routes = [
  {
    path: 'products',
    component: ProductsComponent, // Parent — has its own <router-outlet>
    children: [
      { path: '', component: ProductListComponent },    // /products
      { path: ':id', component: ProductDetailComponent }, // /products/42
      { path: ':id/edit', component: ProductEditComponent }, // /products/42/edit
    ]
  }
];
```

```html
<!-- products.component.html -->
<app-product-sidebar></app-product-sidebar>
<router-outlet></router-outlet> <!-- Children render here -->
```"

#### In Depth
Each `<router-outlet>` corresponds to a level in the route hierarchy. Angular allows **named outlets** (multiple outlets simultaneously):

```html
<router-outlet></router-outlet>           <!-- Primary outlet -->
<router-outlet name="sidebar"></router-outlet>  <!-- Auxiliary outlet -->
```

This enables **auxiliary routes** — displaying a side panel simultaneously with the main route, useful for modals, sidebars, or chat panels that overlay the main content.

---

### 5. What is lazy loading? 🟡 | 🏭🚀

"**Lazy loading** loads a feature module only when the user navigates to its route — not at initial startup. This dramatically reduces the **initial bundle size**, improving **Time to Interactive (TTI)**.

```typescript
// Route config — lazy loading
{
  path: 'admin',
  loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule)
}

// With Angular 14+ standalone components:
{
  path: 'dashboard',
  loadComponent: () => import('./dashboard/dashboard.component').then(c => c.DashboardComponent)
}
```

The admin module is only fetched from the server when someone navigates to `/admin`, saving bandwidth and CPU for users who never visit admin."

#### In Depth
Angular supports **preloading strategies** that balance lazy loading with performance:

- `NoPreloading` (default) — Load only on navigation
- `PreloadAllModules` — Eagerly preload all lazy modules after initial boot
- **Custom strategy** — Preload based on `data.preload` route flag or user connection speed

In production, I use a custom preloading strategy that preloads high-traffic routes but keeps admin/rare routes fully lazy.

---

### 6. What is route guard? 🟡 | 🏭🚀

"A **route guard** is an interface that Angular calls during navigation to decide whether to allow, block, or redirect.

Guards are implemented as injectable services returning `boolean | UrlTree | Observable<boolean> | Promise<boolean>`.

**Common guards:**

| Guard | Purpose |
|---|---|
| `CanActivate` | Can the user access this route? |
| `CanActivateChild` | Can the user access child routes? |
| `CanDeactivate` | Can the user leave this route? (unsaved changes) |
| `CanLoad` | Can the lazy module be loaded? (blocks download) |
| `Resolve` | Pre-fetch data before the route activates |

```typescript
@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
  constructor(private auth: AuthService, private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot): boolean | UrlTree {
    return this.auth.isLoggedIn()
      ? true
      : this.router.createUrlTree(['/login']);
  }
}
```"

#### In Depth
With Angular 14+, guards can be **functional** (no class needed):

```typescript
export const authGuard: CanActivateFn = (route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);
  return auth.isLoggedIn() ? true : router.createUrlTree(['/login']);
};
```

Functional guards reduce boilerplate and are more tree-shakable. `CanLoad` is the strongest guard — it prevents the lazy module's JavaScript file from being **downloaded at all**, protecting sensitive code from bundle inspection. `CanActivate` still allows the bundle to download.

---

### 7. What is the difference between CanActivate and CanDeactivate? 🟡 | 🏭🚀

"- **`CanActivate`** — Guards **entering** a route. Prevents navigation TO the route.
- **`CanDeactivate`** — Guards **leaving** a route. Prevents navigation AWAY from the route.

`CanDeactivate` is used to warn users about **unsaved changes**:

```typescript
export interface CanDeactivateComponent {
  hasUnsavedChanges(): boolean;
}

@Injectable({ providedIn: 'root' })
export class UnsavedChangesGuard implements CanDeactivate<CanDeactivateComponent> {
  canDeactivate(component: CanDeactivateComponent): boolean {
    if (component.hasUnsavedChanges()) {
      return confirm('You have unsaved changes. Leave anyway?');
    }
    return true;
  }
}
```

This prevents accidental data loss when users click the browser back button or navigate away from a form."

#### In Depth
Using browser `confirm()` in guards is a UX anti-pattern — it's blocking and not customizable. The better pattern is to return an observable from `canDeactivate` that resolves after showing a **custom dialog modal**:

```typescript
canDeactivate(): Observable<boolean> {
  return this.dialogService.open(ConfirmDialog, {
    message: 'Discard unsaved changes?'
  }).afterClosed();
}
```

This allows a beautiful, non-blocking confirmation modal while still passing Angular's routing decision back correctly.

---

### 8. How to pass parameters in routes? 🟢 | 🏭🚀

"Angular supports **three types** of route parameters:

**1. Route params (required, part of URL):**
```typescript
// Route: { path: 'products/:id', component: ProductDetailComponent }
// URL: /products/42

// Reading:
constructor(private route: ActivatedRoute) {}
ngOnInit() {
  const id = this.route.snapshot.paramMap.get('id'); // '42'
  // Or reactive (updates without re-creating component):
  this.route.paramMap.subscribe(params => {
    this.loadProduct(params.get('id'));
  });
}
```

**2. Query params (optional, key-value pairs):**
```
// URL: /products?category=electronics&sort=price
this.route.queryParamMap.subscribe(q => {
  this.category = q.get('category');
});
```

**3. Route data (static config):**
```typescript
{ path: 'admin', component: AdminComponent, data: { role: 'ADMIN' } }
// Access: this.route.snapshot.data['role']
```"

#### In Depth
**`snapshot` vs observable** is a key distinction. `snapshot.paramMap` reads the value once at component creation time. If the user navigates from `/products/1` to `/products/2`, Angular **reuses the same component instance** (for performance) and `paramMap` observable emits the new value. Using `snapshot` would miss updates. Always use the observable form when the same component re-navigates with different params.

---

### 9. What is a resolver in Angular routing? 🟡 | 🚀

"A **resolver** pre-fetches data **before** the route activates. The component only mounts once the data is ready, avoiding blank loading states.

```typescript
@Injectable({ providedIn: 'root' })
export class ProductResolver implements Resolve<Product> {
  constructor(private productService: ProductService) {}

  resolve(route: ActivatedRouteSnapshot): Observable<Product> {
    const id = route.paramMap.get('id')!;
    return this.productService.getProduct(id);
  }
}

// Route config:
{ path: ':id', component: ProductDetailComponent, resolve: { product: ProductResolver } }

// Component:
ngOnInit() {
  this.product = this.route.snapshot.data['product'];
}
```"

#### In Depth
Resolvers block navigation until data is ready — this can feel slow if the API is slow. For most UX patterns, **skeleton loading states** inside the component are better than resolvers: the user sees immediate navigation feedback. I use resolvers primarily for **critical, must-have data** (like permissions or configuration) where rendering without data causes errors, not for improving perceived performance.

---

### 10. What is route reuse strategy? 🔴 | 🚀

"The **RouteReuseStrategy** controls whether Angular reuses a component instance when navigating back to a previously visited route, or destroys and recreates it.

By default, Angular destroys components on navigation. Custom reuse preserves state:

```typescript
@Injectable()
export class CustomReuseStrategy implements RouteReuseStrategy {
  private storedRoutes = new Map<string, DetachedRouteHandle>();

  shouldDetach(route: ActivatedRouteSnapshot): boolean {
    return route.data['reuse'] === true; // Only reuse routes with data.reuse flag
  }

  store(route: ActivatedRouteSnapshot, handle: DetachedRouteHandle): void {
    this.storedRoutes.set(route.routeConfig!.path!, handle);
  }

  shouldAttach(route: ActivatedRouteSnapshot): boolean {
    return this.storedRoutes.has(route.routeConfig!.path!);
  }

  retrieve(route: ActivatedRouteSnapshot): DetachedRouteHandle | null {
    return this.storedRoutes.get(route.routeConfig!.path!) ?? null;
  }

  shouldReuseRoute(future: ActivatedRouteSnapshot, curr: ActivatedRouteSnapshot): boolean {
    return future.routeConfig === curr.routeConfig;
  }
}
```

Registration: `{ provide: RouteReuseStrategy, useClass: CustomReuseStrategy }`"

#### In Depth
Route reuse is a double-edged sword. It preserves **scroll positions, filter states, and paginated data** when users hit the back button — a great UX feature. But it can cause **stale data bugs** if the stored component doesn't refresh on reattachment. I implement `ngOnInit`-equivalent logic using `router.events` to detect `NavigationEnd` and refresh data when the component is reattached, not recreated.

---

### 11. How to navigate programmatically? 🟢 | 🏭

"I use the Angular `Router` service's `.navigate()` or `.navigateByUrl()` methods:

```typescript
constructor(private router: Router) {}

// Navigate to route
goToProduct(id: number): void {
  this.router.navigate(['/products', id]);
}

// With query params
goToSearch(term: string): void {
  this.router.navigate(['/search'], { queryParams: { q: term } });
}

// From current route (relative navigation)
goToDetails(): void {
  this.router.navigate(['./details'], { relativeTo: this.route });
}

// Replace current history entry (no back button)
redirectAfterLogin(): void {
  this.router.navigate(['/dashboard'], { replaceUrl: true });
}
```"

#### In Depth
`navigateByUrl` navigates with an absolute URL string — it's simpler but loses type safety. `navigate([...])` is preferred because it constructs the URL from a typed array of segments, making it refactor-friendly. The `{ relativeTo: this.route }` option is critical in child route components — without it, navigation is always absolute from root, ignoring the current route context.

---
