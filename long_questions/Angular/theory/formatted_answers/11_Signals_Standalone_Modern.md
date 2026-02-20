# 🔴 Angular Signals, Standalone APIs & Modern Angular (v17+)

> 🚀 **Product-Based companies** ask these questions heavily — Google, startups, Angular-first companies
> 🏭 **Service-based companies** are catching up but still mostly focus on traditional NgModule patterns
>
> 🎯 **Experience Level:**
> - 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)** — Fresher is transitional, focus on Observable first

---

## 🔹 Angular Signals (Angular 16+)

---

### 1. What are Angular Signals? 🟡 | 🚀

"**Signals** are Angular's new reactive primitives (introduced in Angular 16, stable in Angular 17). A signal is a **reactive value container** — reading it tracks the dependency, and writing it triggers updates in anything that depends on it.

```typescript
import { signal, computed, effect } from '@angular/core';

// Create a writable signal
const count = signal(0);

// Read — returns the current value
console.log(count()); // 0

// Write
count.set(5);
count.update(current => current + 1); // 6

// Computed — derived signal
const doubled = computed(() => count() * 2);
console.log(doubled()); // 12

// Effect — side effect when signal changes
effect(() => {
  console.log(`Count changed to: ${count()}`);
  // Auto-runs whenever count() changes
});
```

Signals enable **fine-grained reactivity** — Angular updates only the components and expressions that depend on a changed signal, skipping unrelated parts of the UI."

#### In Depth
Signals solve the fundamental problem of Angular's zone.js-based change detection: **too many components re-check unnecessarily**. With signals, Angular knows exactly which template expressions depend on which signals. When a signal changes, only those specific expressions re-evaluate — not the entire component tree. This enables **zoneless Angular** where zone.js is completely removed, reducing bundle size by ~35KB and eliminating thousands of unnecessary change detection cycles.

---

### 2. How do Signals differ from RxJS Observables? 🟡 | 🚀

| Aspect | Signals | Observables |
|---|---|---|
| Learning curve | Low | High |
| Subscription required | No — read by calling `signal()` | Yes — `.subscribe()` or `async` pipe |
| Multiple values | Continuous state | Stream over time |
| Synchronous | ✅ Always synchronous | ❌ Can be async |
| Cleanup required | No — automatic | Yes — `unsubscribe()` |
| Integration with CD | Native (fine-grained) | Via `async` pipe or `markForCheck()` |
| Best for | Component state | Events, HTTP, WebSockets |

"I think of Signals as **reactive variables** — replacing `@Input()` + local state patterns. Observables remain essential for **async event streams** (HTTP, WebSockets, user events). They complement each other:

```typescript
// Convert Observable to Signal using toSignal()
import { toSignal } from '@angular/core/rxjs-interop';

readonly products = toSignal(this.productService.getAll(), { initialValue: [] });

// Convert Signal to Observable using toObservable()
import { toObservable } from '@angular/core/rxjs-interop';

readonly products$ = toObservable(this.products);
```"

#### In Depth
The killer pattern combining both: `toSignal(httpCall$)` converts an Observable HTTP call into a signal automatically managing the subscription lifecycle. Angular subscribes when the signal is first read and unsubscribes when the component is destroyed — **zero manual subscription management**. Combined with the `async` pipe alternative (`products()` instead of `products$ | async`), this simplifies templates significantly.

---

### 3. What is `effect()` in Angular Signals? 🔴 | 🚀

"**`effect()`** runs a side-effect function whenever any signal it reads changes. It's the signal-world equivalent of `ngOnChanges` + `rxjs tap()`.

```typescript
@Component({...})
export class DashboardComponent {
  theme = signal<'light' | 'dark'>('light');
  userId = signal<string>('');

  constructor() {
    // Sync document theme whenever signal changes
    effect(() => {
      document.documentElement.setAttribute('data-theme', this.theme());
    });

    // Log user ID changes (for analytics)
    effect(() => {
      if (this.userId()) {
        this.analytics.trackUserLogin(this.userId());
      }
    });
  }
}
```

Effects run:
1. **Once immediately** when first set up
2. **After every change** to any signal they read"

#### In Depth
Effects have **implicit dependency tracking** — Angular automatically identifies which signals are read inside the effect body and registers them as dependencies. This is similar to Vue's `watchEffect`. However, writing to a signal that the same effect reads creates a **circular dependency** (infinite loop). Angular throws an error for this. If you must write to a signal inside an effect, use `untracked()` to break the tracking: `effect(() => { count(); untracked(() => total.update(v => v + 1)); })`.

---

### 4. How do computed signals work? 🟡 | 🚀

"**`computed()`** creates a **derived signal** that automatically recalculates when its source signals change. It's **lazy and memoized** — only recomputes when consumers read it AND a dependency changed.

```typescript
// cart.component.ts
const items = signal<CartItem[]>([]);
const discount = signal(0.1); // 10% discount

// Computed — auto-updates when items or discount changes
const subtotal = computed(() =>
  items().reduce((sum, item) => sum + item.price * item.qty, 0)
);

const total = computed(() =>
  subtotal() * (1 - discount())
);

const itemCount = computed(() => items().length);

// Template usage — no subscription needed!
// {{ total() | currency }}
// {{ itemCount() }} items
```"

#### In Depth
`computed()` is **memoized** — if none of its dependencies have changed, calling it returns the **same cached value** without recomputation. This is similar to `useMemo` in React. Computed signals are read-only (`WritableSignal` vs `Signal` TypeScript types). The computed graph is evaluated **lazily** — if nobody reads `total()`, it never recalculates, even if `items` or `discount` changed. This is more efficient than RxJS `combineLatest` which always emits on any source change.

---

## 🔹 Standalone Components (Angular 14+)

---

### 5. What are standalone components? 🟡 | 🏭🚀

"**Standalone components** (introduced in Angular 14) don't need to be declared in an `NgModule`. They declare their own dependencies directly in the `@Component` decorator's `imports` array.

```typescript
// Traditional (Module-based)
@NgModule({
  declarations: [UserCardComponent],
  imports: [CommonModule, RouterModule]
})
export class UserModule {}

// Modern (Standalone)
@Component({
  selector: 'app-user-card',
  standalone: true,              // 👈 Key flag
  imports: [CommonModule, RouterLink, DatePipe, NgIf], // Dependencies here
  template: `
    <div *ngIf="user">
      <h2>{{ user.name }}</h2>
      <a [routerLink]="['/users', user.id]">View Profile</a>
    </div>
  `
})
export class UserCardComponent {
  @Input() user: User;
}
```

Bootstrap with standalone components:
```typescript
// main.ts
bootstrapApplication(AppComponent, {
  providers: [provideRouter(routes), provideHttpClient()]
});
```"

#### In Depth
Standalone components improve **tree shaking** significantly. Each component declares exactly what it uses — no module-level declarations that include everything. This makes `ng build` bundle analysis more accurate and helps the compiler eliminate unused imports. Standalone is now the **Angular team's recommended approach** for all new projects (Angular 17+ defaults to standalone in `ng new`). I migrate module-based apps incrementally by converting leaf components first.

---

### 6. How do signals impact change detection? 🔴 | 🚀

"Signals fundamentally change **how Angular knows what to update**:

**Without Signals (zone.js):**
- Any async operation triggers a full component tree check
- Angular uses `===` to compare all bound values
- CPU waste: checking components that didn't change

**With Signals (Signal-based CD):**
- Angular knows **exactly** which signals a template reads (tracked during rendering)
- When a signal changes, only templates using that signal re-render
- Zero zone.js overhead for signal-based changes

```typescript
@Component({
  template: `
    <h1>{{ title() }}</h1>   <!-- Only re-renders when title signal changes -->
    <p>{{ count() }}</p>     <!-- Only re-renders when count signal changes -->
    <app-child />            <!-- Never re-renders due to parent changes -->
  `
})
export class ParentComponent {
  title = signal('Dashboard');
  count = signal(0);
}
```

This enables **signal-only components** that work without zone.js entirely."

#### In Depth
Angular 17+ supports **mixed mode** — components can use both traditional Observable/zone-based patterns and signals. Angular 18+ will make zoneless a stable feature. The migration path is gradual: start adding signals for new state, use `toSignal()` to convert existing Observables, and eventually remove zone.js. This is the biggest architectural change in Angular since Ivy — and it's designed to be 100% backward compatible.

---

### 7. Can Signals replace NgRx in some cases? 🔴 | 🚀

"Yes — for **component and feature-scoped state**, signals are a cleaner alternative to NgRx. For **app-wide state** requiring DevTools, time-travel debugging, and complex action orchestration, NgRx remains superior.

**Signal-based store pattern (simple alternative to NgRx):**

```typescript
@Injectable({ providedIn: 'root' })
export class CartStore {
  // State
  private items = signal<CartItem[]>([]);
  private discount = signal(0);

  // Read-only public API
  readonly items$ = this.items.asReadonly();
  readonly total = computed(() =>
    this.items().reduce((s, i) => s + i.price, 0) * (1 - this.discount())
  );
  readonly count = computed(() => this.items().length);

  // Actions
  addItem(item: CartItem): void {
    this.items.update(current => [...current, item]);
  }

  removeItem(id: number): void {
    this.items.update(items => items.filter(i => i.id !== id));
  }

  applyDiscount(pct: number): void {
    this.discount.set(pct / 100);
  }
}
```

This pattern is simpler than NgRx but provides reactive state, computed values, and clear mutation points. Use it for medium-complexity apps."

#### In Depth
NgRx itself is adopting signals — **NgRx Signals** (`@ngrx/signals`) is a new package providing a Signal-based store with EntityAdapter support. It removes the action/reducer/effect boilerplate while keeping the benefits of a structured, testable state layer. I see this as the future of Angular state management: signal primitives for reactivity + NgRx's pattern for structure, without zone.js overhead.

---

### 8. What is the `@defer` block? 🔴 | 🚀

"**`@defer`** is Angular 17's template-level lazy loading syntax. It wraps components/content that should load lazily, with built-in states for loading, error, and placeholder:

```html
<!-- Load reviews only when they scroll into view -->
@defer (on viewport; prefetch on idle) {
  <app-product-reviews [productId]="id" />
} @placeholder (minimum 500ms) {
  <!-- Shown immediately while deferred content hasn't started loading -->
  <div class="reviews-skeleton" style="height: 200px;"></div>
} @loading (minimum 200ms; after 100ms) {
  <!-- Shown while the chunk is being fetched -->
  <mat-progress-bar mode="indeterminate"></mat-progress-bar>
} @error {
  <!-- Shown if loading fails -->
  <button (click)="retryReviews()">Retry loading reviews</button>
}
```

**Triggers:**
- `on viewport` — Visible in browser viewport
- `on idle` — Browser `requestIdleCallback`
- `on interaction` — User clicks/hovers the placeholder
- `on timer(2s)` — After a delay
- `when isLoggedIn()` — Custom condition (can use signals!)"

#### In Depth
`@defer` generates a **dynamic import** for the deferred component automatically — without any router configuration or explicit `loadComponent()` call. The deferred bundle is only downloaded when the trigger condition is met. The `prefetch on idle` option is a hybrid approach: the chunk is prefetched in the background during idle time, so when the user scrolls to the component, the download is already complete — combining the bundle efficiency of lazy loading with the UX of eager loading.

---
