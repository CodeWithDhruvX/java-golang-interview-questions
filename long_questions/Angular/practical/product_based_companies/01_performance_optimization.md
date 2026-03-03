# 📘 01 — Performance Optimization
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- `ChangeDetectionStrategy.OnPush` — when and why
- `trackBy` with `*ngFor` — avoiding unnecessary DOM recreation
- Lazy loading modules and routes
- `@defer` blocks (Angular 17+) for component-level lazy loading
- Virtual scrolling for large lists
- Bundle analysis and tree-shaking
- Memory leak prevention

---

## ❓ Most Asked Questions

### Q1. How does `ChangeDetectionStrategy.OnPush` improve performance?

```typescript
// DEFAULT: Angular checks EVERY component on every async event — O(n) tree traversal
// ONPUSH: Angular skips this component unless ONE of four conditions is met

@Component({
  selector: 'app-product-card',
  changeDetection: ChangeDetectionStrategy.OnPush,  // ✅ performance optimization
  template: `
    <div class="card">
      <h3>{{ product.name }}</h3>
      <p>{{ product.price | currency: 'INR' }}</p>
      <button (click)="addToCart()">Add to Cart</button>
    </div>
  `
})
export class ProductCardComponent {
  @Input() product: Product;
  @Output() added = new EventEmitter<Product>();

  constructor(private cdr: ChangeDetectorRef) {}

  addToCart(): void {
    this.added.emit(this.product);
    // event binding triggers CD automatically — no need to call detectChanges
  }
}
```

**OnPush triggers change detection ONLY when:**
1. An `@Input()` **reference** changes (not mutated)
2. An event **inside** this component fires (click, keyup)
3. An `Observable` subscribed via `async` pipe emits
4. `markForCheck()` or `detectChanges()` is called manually

```typescript
// ❌ This WON'T trigger OnPush update — mutating in place
this.product.price = 999;  // same reference!

// ✅ This WILL trigger — new reference
this.product = { ...this.product, price: 999 };

// Manually triggering CD from outside (e.g., from a timer or WebSocket)
this.cdr.markForCheck();   // marks component + ancestors for check during next CD cycle
this.cdr.detectChanges();  // immediately runs CD on this component subtree
```

---

### Q2. Why is `trackBy` critical for `*ngFor` performance?

```typescript
// Without trackBy: Angular destroys and recreates ALL DOM nodes
// when the array reference changes (even if data is same)

// Product list component
@Component({
  template: `
    <!-- ❌ Without trackBy — re-renders ALL items on any array change -->
    <div *ngFor="let p of products">{{ p.name }}</div>

    <!-- ✅ With trackBy — Angular only re-renders items that actually changed -->
    <div *ngFor="let p of products; trackBy: trackById">
      <app-product-card [product]="p"></app-product-card>
    </div>
  `
})
export class ProductListComponent {
  products: Product[] = [];

  // trackBy function — returns a unique identifier for each item
  trackById(index: number, product: Product): string {
    return product.id;
  }

  // When you add one item:
  // WITHOUT trackBy: ALL 100 product DOM nodes destroyed and recreated
  // WITH trackBy: Only the 1 new DOM node is created
  addProduct(p: Product): void {
    this.products = [...this.products, p];  // new array reference triggers ngFor
  }
}
```

---

### Q3. What is `@defer` in Angular 17+?

```typescript
// @defer enables component-level lazy loading without route changes!
// The deferred content is split into a separate JS chunk and loaded on demand

@Component({
  template: `
    <!-- Load heavy chart component only when user scrolls near it -->
    @defer (on viewport) {
      <app-analytics-dashboard></app-analytics-dashboard>
    } @loading {
      <div class="skeleton-loader"></div>
    } @error {
      <p>Failed to load dashboard. <button (click)="retry()">Retry</button></p>
    } @placeholder {
      <div class="placeholder-box">Chart will appear here</div>
    }

    <!-- Load when user hovers over a trigger -->
    @defer (on hover(triggerRef); prefetch on idle) {
      <app-product-recommendations></app-product-recommendations>
    }
    <button #triggerRef>Show Recommendations</button>

    <!-- Load on interaction -->
    @defer (on interaction(btn)) {
      <app-comments-section></app-comments-section>
    }
    <button #btn>Load Comments</button>
  `
})
export class ProductDetailComponent {}
```

**Defer triggers:**
| Trigger | When |
|---------|------|
| `on viewport` | Element is visible in viewport |
| `on idle` | Browser is idle |
| `on interaction(ref)` | User interacts with the referenced element |
| `on hover(ref)` | Mouse hovers over element |
| `on timer(2000)` | After 2 seconds |
| `when condition` | When a boolean expression is truthy |

---

### Q4. How do you implement virtual scrolling for large lists?

```typescript
// Without virtual scrolling: rendering 10,000 DOM nodes → browser jank
// With virtual scrolling: only renders ~10-20 items visible in the viewport

// Install: npm install @angular/cdk

import { ScrollingModule } from '@angular/cdk/scrolling';

@Component({
  template: `
    <cdk-virtual-scroll-viewport itemSize="72" class="product-list-viewport">
      <div *cdkVirtualFor="let product of products; trackBy: trackById"
           class="product-row">
        <img [src]="product.imageUrl" />
        <span>{{ product.name }}</span>
        <span>{{ product.price | currency: 'INR' }}</span>
      </div>
    </cdk-virtual-scroll-viewport>
  `,
  styles: [`.product-list-viewport { height: 600px; }`]
})
export class ProductListComponent {
  products: Product[] = Array.from({ length: 10000 }, (_, i) => ({
    id: i.toString(), name: `Product ${i}`, price: i * 100
  }));

  trackById = (_: number, p: Product) => p.id;
}
```

---

### Q5. How do you analyze and reduce Angular bundle size?

```bash
# Step 1: Build with stats-json flag
ng build --configuration=production --stats-json

# Step 2: Analyze with webpack-bundle-analyzer
npx webpack-bundle-analyzer dist/my-app/stats.json

# Step 3: Set bundle budgets in angular.json
```

```json
// angular.json — configure budget limits (CI will fail if exceeded)
"budgets": [
  {
    "type": "initial",
    "maximumWarning": "500kb",
    "maximumError": "1mb"
  },
  {
    "type": "anyComponentStyle",
    "maximumWarning": "2kb",
    "maximumError": "4kb"
  }
]
```

**Bundle size reduction strategies:**

```typescript
// 1. Lazy load all feature modules (biggest win)
{ path: 'reports', loadChildren: () => import('./reports/reports.module') }

// 2. Use standalone components (Angular 14+) to avoid NgModule overhead
@Component({ standalone: true, imports: [CommonModule] })

// 3. Replace moment.js with date-fns (tree-shakeable)
import { format } from 'date-fns';  // only imported function is bundled

// 4. Use CDK instead of full Material for utilities

// 5. Run source-map-explorer to find large deps
npx source-map-explorer dist/my-app/*.js
```

---

### Q6. How do you prevent memory leaks in Angular?

```typescript
// Pattern 1: takeUntil with destroy$ Subject (classic approach)
@Component({ ... })
export class DataComponent implements OnInit, OnDestroy {
  private destroy$ = new Subject<void>();

  ngOnInit(): void {
    this.dataService.stream$.pipe(
      takeUntil(this.destroy$)
    ).subscribe(data => this.data = data);

    interval(1000).pipe(
      takeUntil(this.destroy$)
    ).subscribe(() => this.refresh());
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}

// Pattern 2: DestroyRef (Angular 16+) — no need for OnDestroy lifecycle
@Component({ ... })
export class ModernComponent implements OnInit {
  constructor(private destroyRef: DestroyRef) {}

  ngOnInit(): void {
    this.dataService.stream$.pipe(
      takeUntilDestroyed(this.destroyRef)  // auto-cleanup!
    ).subscribe(data => this.data = data);
  }
}

// Pattern 3: async pipe — always preferred (zero risk of leak)
// The async pipe auto-unsubscribes when the component is destroyed
{{ data$ | async }}
```

---

### Q7. What is `NgZone.runOutsideAngular()`?

```typescript
// Problem: zone.js triggers Angular's CD on EVERY async event
// including scroll, mousemove, requestAnimationFrame — causing performance issues

@Component({ ... })
export class AnimationComponent implements AfterViewInit {
  @ViewChild('canvas') canvasRef: ElementRef;

  constructor(private ngZone: NgZone) {}

  ngAfterViewInit(): void {
    const canvas = this.canvasRef.nativeElement;

    // ✅ Run animation loop OUTSIDE Angular's zone
    // Angular won't run change detection on each animation frame
    this.ngZone.runOutsideAngular(() => {
      const animate = () => {
        this.drawFrame(canvas);
        requestAnimationFrame(animate);  // won't trigger CD each frame
      };
      animate();
    });

    // When you DO need to update Angular's view from outside zone
    this.ngZone.runOutsideAngular(() => {
      fromEvent(window, 'scroll').subscribe(() => {
        if (this.shouldUpdateUI()) {
          this.ngZone.run(() => {
            this.scrollPosition = window.scrollY;  // triggers CD
          });
        }
      });
    });
  }
}
```
