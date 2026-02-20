# 🔴 Enterprise Architecture & Design Patterns

> 🚀 **Heavily asked by Product-Based companies** — System design, scalability, large-team patterns
> 🏭 **Service-Based companies** ask simpler architecture questions in senior-level interviews
>
> 🎯 **Experience Level:** 🔴 **Senior (5+ yrs)**

---

### 1. How do you architect a large-scale Angular application? 🔴 | 🚀

"For large-scale Angular applications, I follow a **domain-driven, layered architecture**:

```
apps/
  └── my-app/
      └── src/
          ├── app/
          │   ├── core/           # Singleton services, guards, interceptors
          │   │   ├── auth/
          │   │   ├── interceptors/
          │   │   └── services/
          │   ├── features/       # Feature modules (lazy loaded)
          │   │   ├── products/
          │   │   │   ├── data-access/   # API calls, state
          │   │   │   ├── ui/            # Presentational components
          │   │   │   └── utils/         # Feature-specific utilities
          │   │   └── orders/
          │   └── shared/         # Shared components, pipes, directives
          └── assets/
```

**Key principles:**
1. **Smart/Dumb component pattern** — Container components handle data, presentational components just render
2. **Feature encapsulation** — Features expose public APIs; internals are private
3. **Dependency direction** — Feature modules depend on `shared` and `core`, never on each other
4. **Lazy loading** — Every feature route is lazy loaded"

#### In Depth
For very large applications (50+ developers, 100+ features), I use **Nx monorepo** with workspace libraries. Each library (e.g., `@myco/products/data-access`, `@myco/products/ui`) has strict boundary rules enforced by `@nx/enforce-module-boundaries` ESLint rule. No library can import from another library in a circular manner. This prevents **dependency creep** that makes large codebases unmaintainable and creates build order problems.

---

### 2. What is the Smart/Dumb component pattern? 🟡 | 🏭🚀

"The **Smart/Dumb (Container/Presentational) pattern** separates components into two roles:

**Smart Components (Containers):**
- Know about services and state management
- Handle async operations and data fetching
- Pass data to dumb components via `@Input()`
- React to events from dumb components via `@Output()`

**Dumb Components (Presentational):**
- Only accept `@Input()` and emit `@Output()`
- No service injections
- Pure UI rendering — highly reusable
- Always use `OnPush` change detection

```typescript
// SMART — Product List Container
@Component({
  template: `
    <app-product-grid
      [products]="products$ | async"
      [loading]="loading$ | async"
      (productSelected)="onProductSelected($event)"
    />
  `
})
export class ProductListContainer {
  products$ = this.store.select(selectProducts);
  loading$ = this.store.select(selectLoading);
  constructor(private store: Store) { this.store.dispatch(loadProducts()); }
  onProductSelected(product: Product): void { /* navigate, show modal */ }
}

// DUMB — Product Grid (pure display)
@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  inputs: ['products', 'loading'],
  outputs: ['productSelected']
})
export class ProductGridComponent {}
```"

#### In Depth
This pattern enables **maximum reusability** — `ProductGridComponent` works with any data source (NgRx store, BehaviorSubject, signals) because it's completely decoupled from state. Testing is dramatically simpler: dumb components are tested by just passing `@Input()` values and asserting the DOM output — no mocking required. Smart components can be tested by mocking the store/services and verifying dispatched actions.

---

### 3. What is Module Federation in Angular? 🔴 | 🚀

"**Module Federation** (via `@angular-architects/module-federation`) enables **Micro-Frontend** architecture — multiple Angular apps running as one.

```typescript
// Shell app (host) loads remote micro-frontends
const routes: Routes = [
  {
    path: 'products',
    loadChildren: () =>
      loadRemoteModule({
        type: 'module',
        remoteEntry: 'http://localhost:4201/remoteEntry.mjs', // Products team's app
        exposedModule: './Module'
      }).then(m => m.ProductsModule)
  },
  {
    path: 'checkout',
    loadChildren: () =>
      loadRemoteModule({
        type: 'module',
        remoteEntry: 'http://localhost:4202/remoteEntry.mjs', // Checkout team's app
        exposedModule: './Module'
      }).then(m => m.CheckoutModule)
  }
];
```

Benefits:
- **Independent deployments** — Products team deploys without coordinating with Checkout team
- **Technology flexibility** — Micro-frontends can use different Angular versions
- **Team autonomy** — Each team owns their feature end-to-end"

#### In Depth
The challenge of Module Federation is **shared dependencies**. If `shell` and `products-mfe` both bundle Angular, the user downloads Angular twice. Webpack's `shared` config resolves this:

```javascript
// webpack.config.js
shared: share({
  '@angular/core': { singleton: true, strictVersion: true },
  '@angular/router': { singleton: true, strictVersion: true }
})
```

`singleton: true` ensures only one version of Angular runs in the browser (the shell's version). `strictVersion: true` throws an error if version ranges don't match. This is critical for correctness — Angular cannot have two instances in the same browsing context.

---

### 4. What are effective code-splitting strategies in Angular? 🔴 | 🚀

"Code splitting reduces TTI by loading only what's needed. Strategies from coarse to fine:

**1. Route-level splitting (most common):**
```typescript
{ path: 'admin', loadChildren: () => import('./admin/admin.module') }
```

**2. Component-level splitting (`@defer`):**
```html
@defer (on viewport) { <app-data-visualization /> }
```

**3. Conditional feature splitting:**
```typescript
// Load feature only if user has specific role
if (user.role === 'admin') {
  const { AdminDashboardComponent } = await import('./admin-dashboard.component');
  viewContainerRef.createComponent(AdminDashboardComponent);
}
```

**4. Preloading strategies:**
```typescript
RouterModule.forRoot(routes, {
  preloadingStrategy: QuicklinkStrategy // Preload links visible in viewport
})
```

I analyze the bundle with `webpack-bundle-analyzer` after each change to verify splits are effective."

#### In Depth
I combine `@defer` with service workers for optimal loading:
1. `@defer (on idle)` loads non-critical components during browser idle time
2. Service worker caches these dynamic chunks after first download
3. Second visit: all chunks served from cache — near-instant rendering

The **Critical Path** principle guides what to eager-load vs defer: only components needed for the page's **Largest Contentful Paint (LCP)** element should be eager. Everything else should be `@defer`.

---

### 5. What is NGXS vs NgRx? 🔴 | 🚀

"**NGXS** and **NgRx** are both state management libraries for Angular, but with different philosophies:

| Aspect | NgRx | NGXS |
|---|---|---|
| Paradigm | Functional (Redux-style) | Object-oriented (class-based) |
| Boilerplate | High (actions + reducers + effects + selectors) | Medium (state class + actions) |
| Learning curve | Steep | Moderate |
| DevTools | Excellent (Redux DevTools) | Moderate |
| Community | Larger | Smaller |

**NgRx:**
```typescript
// Pure functions, verbose but explicit
createAction, createReducer, createEffect, createSelector
```

**NGXS:**
```typescript
@State<CartStateModel>({ name: 'cart', defaults: { items: [] } })
@Injectable()
export class CartState {
  @Action(AddItem)
  addItem(ctx: StateContext<CartStateModel>, { item }: AddItem): void {
    ctx.patchState({ items: [...ctx.getState().items, item] });
  }
}
```

I recommend **NgRx** for large teams (more community resources, stricter patterns) and **NGXS** or **signal-based stores** for smaller teams who find NgRx boilerplate overwhelming."

#### In Depth
The trend is moving toward **NgRx Signals** (`@ngrx/signals`) which combines NgRx's structured approach with the simplicity of the Signal API. It requires no actions or reducers — state is defined with `signalStore()`:

```typescript
const CartStore = signalStore(
  withState({ items: [] as CartItem[] }),
  withComputed(({ items }) => ({ total: computed(() => items().reduce(...)) })),
  withMethods((store) => ({ addItem: (item) => patchState(store, ...) }))
);
```

This is likely the future of Angular state management — structured pattern of NgRx with signal's ergonomics.

---

### 6. How to handle multi-tenancy in Angular? 🔴 | 🚀

"Multi-tenancy means **one Angular app serves multiple clients** (tenants with different branding, features, or data access).

**Approach 1: Theme-level multi-tenancy (CSS custom properties):**
```typescript
// Tenant config loaded at bootstrap
@Injectable({ providedIn: 'root' })
export class TenantService {
  applyTheme(tenantConfig: TenantConfig): void {
    const root = document.documentElement;
    root.style.setProperty('--primary-color', tenantConfig.primaryColor);
    root.style.setProperty('--font-family', tenantConfig.fontFamily);
  }
}
```

**Approach 2: Feature flags per tenant:**
```typescript
// Guard based on tenant feature access
canActivate(): boolean {
  return this.tenantService.hasFeature('advanced-analytics');
}
```

**Approach 3: White-labeling with build-time configuration:**
```bash
# Build per tenant
ng build --configuration=tenant-acme
ng build --configuration=tenant-globex
```"

#### In Depth
For true multi-tenancy at scale, I combine tenant detection at the CDN/reverse-proxy layer with Angular's feature flag system. The CDN routes `acme.myapp.com` and `globex.myapp.com` to the same Angular build, but injects a `<script>window.TENANT_CONFIG = {...}</script>` tag before the Angular bundle loads. Angular reads this at bootstrap to configure themes, features, and API endpoints. This avoids per-tenant builds while maintaining complete customization.

---
