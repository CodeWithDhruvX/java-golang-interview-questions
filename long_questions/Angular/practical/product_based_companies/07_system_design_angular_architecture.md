# 📘 07 — System Design & Angular Architecture
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Micro frontend architecture with Module Federation
- Monorepo setup with Nx
- Component library / Design system architecture
- Feature module organization strategies
- Shell architecture with lazy loading
- Angular performance at scale

---

## ❓ Most Asked Questions

### Q1. What are Micro Frontends? How does Module Federation work with Angular?

**Micro Frontends** allow large frontend applications to be split into independently developed, deployed, and maintained frontend apps.

```
                    ┌────────────────────────────────┐
                    │       Shell App (Host)          │
                    │  ┌──────────────────────────┐  │
                    │  │  Header (Shell-owned)     │  │
                    │  └──────────────────────────┘  │
                    │  ┌─────────┐  ┌─────────────┐  │
                    │  │Products │  │  Orders MFE │  │
                    │  │  MFE    │  │(Independent)│  │
                    │  │(Team A) │  │  (Team B)   │  │
                    │  └─────────┘  └─────────────┘  │
                    └────────────────────────────────┘
```

```javascript
// webpack.config.js in products-app (REMOTE)
module.exports = {
  plugins: [
    new ModuleFederationPlugin({
      name: 'productsApp',
      filename: 'remoteEntry.js',  // entry point exposed to host
      exposes: {
        // Expose Angular routes (lazy-loadable)
        './ProductsModule': './src/app/products/products.module.ts',
        './ProductCardComponent': './src/app/product-card/product-card.component.ts',
      },
      shared: {
        '@angular/core': { singleton: true, strictVersion: true },
        '@angular/common': { singleton: true },
        '@angular/router': { singleton: true },
      }
    })
  ]
};
```

```javascript
// webpack.config.js in shell-app (HOST)
module.exports = {
  plugins: [
    new ModuleFederationPlugin({
      remotes: {
        productsApp: 'productsApp@http://localhost:4201/remoteEntry.js',
        ordersApp:   'ordersApp@http://localhost:4202/remoteEntry.js',
      },
      shared: { '@angular/core': { singleton: true } }
    })
  ]
};
```

```typescript
// Shell app routing — lazy loads from remote MFE
const routes: Routes = [
  { path: '', component: HomeComponent },
  {
    path: 'products',
    loadChildren: () => import('productsApp/ProductsModule').then(m => m.ProductsModule)
  },
  {
    path: 'orders',
    loadChildren: () => import('ordersApp/OrdersModule').then(m => m.OrdersModule)
  }
];
```

---

### Q2. How do you organize a large Angular application?

**Feature-based module structure (recommended):**

```
src/
├── app/
│   ├── core/                      # Singleton services, guards, interceptors
│   │   ├── auth/
│   │   │   ├── auth.service.ts
│   │   │   └── auth.guard.ts
│   │   ├── interceptors/
│   │   │   ├── auth.interceptor.ts
│   │   │   └── error.interceptor.ts
│   │   └── core.module.ts         # imported ONLY in AppModule
│   │
│   ├── shared/                    # Reusable components, pipes, directives
│   │   ├── components/
│   │   │   ├── loading-spinner/
│   │   │   └── confirm-dialog/
│   │   ├── pipes/
│   │   ├── directives/
│   │   └── shared.module.ts       # exported for all feature modules to import
│   │
│   ├── features/                  # Feature modules (lazy loaded)
│   │   ├── products/
│   │   │   ├── components/
│   │   │   ├── services/
│   │   │   ├── store/             # NgRx state for this feature
│   │   │   │   ├── actions/
│   │   │   │   ├── reducers/
│   │   │   │   ├── effects/
│   │   │   │   └── selectors/
│   │   │   ├── models/
│   │   │   └── products.module.ts
│   │   │
│   │   ├── orders/
│   │   └── admin/
│   │
│   ├── app-routing.module.ts
│   └── app.module.ts
│
├── environments/
└── assets/
```

**Key principles:**
- `CoreModule` provides singleton services — imported once in `AppModule`
- `SharedModule` exports reusable UI — imported in each feature module
- Feature modules are lazy-loaded and self-contained
- NgRx state is feature-scoped (`StoreModule.forFeature()`)

---

### Q3. How do you set up a monorepo with Nx?

**Nx** extends the Angular CLI for monorepo management — multiple apps and shared libraries in one repository.

```bash
# Create Nx workspace
npx create-nx-workspace@latest my-org --preset=angular

# Generate a new Angular app
nx generate @nx/angular:app products-app

# Generate a shared library
nx generate @nx/angular:library ui-components --directory=libs/shared
nx generate @nx/angular:library data-access-products --directory=libs/products

# Build
nx build products-app

# Test only affected projects (saves CI time)
nx affected:test
nx affected:build

# Visual dependency graph
nx graph
```

```
my-org/
├── apps/
│   ├── shell/          # Main host app
│   ├── products-app/   # Products micro frontend
│   └── orders-app/     # Orders micro frontend
│
├── libs/
│   ├── shared/
│   │   ├── ui-components/    # Design system — buttons, cards, modals
│   │   ├── utils/            # Pure utility functions
│   │   └── data-models/      # Shared TypeScript interfaces
│   │
│   ├── products/
│   │   ├── data-access/      # ProductService, NgRx store
│   │   ├── feature/          # Product feature components
│   │   └── ui/               # Product-specific UI components
│   │
│   └── orders/
│       ├── data-access/
│       └── feature/
```

---

### Q4. How do you design a scalable component architecture?

```typescript
// SMART (Container) vs DUMB (Presentational) component pattern

// ✅ SMART/Container — knows about services, state, routing
@Component({
  selector: 'app-product-list-container',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <app-product-list
      [products]="products$ | async"
      [loading]="loading$ | async"
      [selectedCategory]="selectedCategory$ | async"
      (productSelected)="onProductSelected($event)"
      (categoryChanged)="onCategoryChanged($event)"
      (loadMore)="onLoadMore()">
    </app-product-list>
  `
})
export class ProductListContainerComponent {
  products$ = this.facade.products$;
  loading$ = this.facade.loading$;
  selectedCategory$ = this.facade.selectedCategory$;

  constructor(private facade: ProductFacade, private router: Router) {}

  onProductSelected(id: string): void {
    this.router.navigate(['/products', id]);
  }

  onCategoryChanged(category: string): void {
    this.facade.filterByCategory(category);
  }
}

// ✅ DUMB/Presentational — pure UI, only @Input/@Output
@Component({
  selector: 'app-product-list',
  changeDetection: ChangeDetectionStrategy.OnPush,  // always OnPush for dumb components
  template: `...`
})
export class ProductListComponent {
  @Input() products: Product[] | null = [];
  @Input() loading: boolean | null = false;
  @Input() selectedCategory: string | null = null;

  @Output() productSelected = new EventEmitter<string>();
  @Output() categoryChanged = new EventEmitter<string>();
  @Output() loadMore = new EventEmitter<void>();
}
```

---

### Q5. How do you handle authentication in a large Angular app?

```typescript
// Complete auth architecture:

// 1. Token storage service
@Injectable({ providedIn: 'root' })
export class TokenService {
  private readonly TOKEN_KEY = 'auth_token';
  private readonly REFRESH_KEY = 'refresh_token';

  getToken(): string | null { return localStorage.getItem(this.TOKEN_KEY); }
  setToken(token: string): void { localStorage.setItem(this.TOKEN_KEY, token); }
  clearTokens(): void {
    localStorage.removeItem(this.TOKEN_KEY);
    localStorage.removeItem(this.REFRESH_KEY);
  }
}

// 2. Auth interceptor — adds token + handles 401 with token refresh
@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  private isRefreshing = false;
  private refreshTokenSubject = new BehaviorSubject<string | null>(null);

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = this.tokenService.getToken();
    const authReq = token ? this.addToken(req, token) : req;

    return next.handle(authReq).pipe(
      catchError((err: HttpErrorResponse) => {
        if (err.status === 401) {
          return this.handle401(authReq, next);
        }
        return throwError(() => err);
      })
    );
  }

  private handle401(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (this.isRefreshing) {
      // Queue requests while refreshing
      return this.refreshTokenSubject.pipe(
        filter(t => t !== null),
        take(1),
        switchMap(token => next.handle(this.addToken(req, token!)))
      );
    }

    this.isRefreshing = true;
    this.refreshTokenSubject.next(null);

    return this.authService.refreshToken().pipe(
      switchMap(token => {
        this.isRefreshing = false;
        this.refreshTokenSubject.next(token);
        return next.handle(this.addToken(req, token));
      }),
      catchError(() => {
        this.isRefreshing = false;
        this.authService.logout();
        return EMPTY;
      })
    );
  }
}
```
