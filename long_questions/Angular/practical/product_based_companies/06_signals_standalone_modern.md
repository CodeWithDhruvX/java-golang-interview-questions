# 📘 06 — Signals, Standalone Components & Modern Angular
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Angular Signals (Angular 16+) — reactive primitives replacing Zone.js
- `signal()`, `computed()`, `effect()`
- Standalone Components (Angular 14+)
- New control flow syntax: `@if`, `@for`, `@switch` (Angular 17+)
- `@let` template variables (Angular 18+)
- `input()`, `output()`, `model()` signal-based APIs (Angular 17.1+)

---

## ❓ Most Asked Questions

### Q1. What are Angular Signals? How do they differ from RxJS Observables?

**Signals** are a new reactive primitive in Angular 16+ that represent a value that changes over time. Unlike Observables, signals are **synchronous**, **always have a current value**, and **don't require subscription management**.

```typescript
import { signal, computed, effect } from '@angular/core';

@Component({
  standalone: true,
  template: `
    <p>Count: {{ count() }}</p>         <!-- call signal as function to read -->
    <p>Double: {{ doubled() }}</p>       <!-- computed auto-updates -->
    <button (click)="increment()">+</button>
    <button (click)="reset()">Reset</button>
  `
})
export class CounterComponent {
  // 1. signal() — writable reactive value
  count = signal(0);

  // 2. computed() — derived value, auto-updates when dependencies change
  doubled = computed(() => this.count() * 2);
  isEven = computed(() => this.count() % 2 === 0);

  // 3. effect() — side effect that runs when signals change
  logEffect = effect(() => {
    console.log(`Count changed to: ${this.count()}`);  // auto-rerun on change
  });

  increment(): void {
    this.count.update(c => c + 1);  // update based on previous value
  }

  reset(): void {
    this.count.set(0);  // set directly
  }
}
```

**Signals vs Observables:**

| Feature | Signals | Observables |
|---------|---------|------------|
| Current value | Always available (`signal()`) | Not always (need BehaviorSubject) |
| Subscription needed | No | Yes |
| Async | Synchronous | Can be async |
| Unsubscribe | Not needed | Required |
| Change detection | Fine-grained (no Zone.js) | Full tree check |
| Composability | `computed()` | `pipe()` with operators |
| When to use | Component state, inputs | HTTP, events, async streams |

---

### Q2. What are Signal-based inputs, outputs, and model?

```typescript
// Angular 17.1+ — new signal-based component APIs

import { input, output, model } from '@angular/core';

@Component({
  selector: 'app-product-form',
  standalone: true,
  template: `
    <h2>Editing: {{ product().name }}</h2>
    <input [value]="localName()" (input)="localName.set($event.target.value)" />
    <button (click)="save()">Save</button>
    <button (click)="cancelled.emit()">Cancel</button>
  `
})
export class ProductFormComponent {
  // input() — signal-based @Input()
  product = input.required<Product>();           // required — throws if not passed
  readonly = input<boolean>(false);              // optional with default

  // output() — signal-based @Output()
  saved = output<Product>();
  cancelled = output<void>();

  // model() — two-way binding signal (combines input + output)
  // Usage: <app-product-form [(name)]="productName">
  name = model<string>('');

  // Works perfectly with computed()
  localName = signal(this.product().name);
  isValid = computed(() => this.localName().trim().length >= 3);

  save(): void {
    if (this.isValid()) {
      this.saved.emit({ ...this.product(), name: this.localName() });
    }
  }
}
```

---

### Q3. What are Standalone Components? How do they replace NgModules?

```typescript
// Traditional NgModule approach
@NgModule({
  declarations: [ProductCardComponent],
  imports: [CommonModule, RouterModule],
  exports: [ProductCardComponent]
})
export class ProductsModule {}

// ✅ Standalone Component — no NgModule needed
@Component({
  selector: 'app-product-card',
  standalone: true,
  // Import dependencies directly in the component
  imports: [CommonModule, RouterModule, CurrencyPipe, DatePipe],
  templateUrl: './product-card.component.html',
})
export class ProductCardComponent {
  @Input() product: Product;
}
```

```typescript
// Bootstrapping a fully standalone Angular app (Angular 15+)
// main.ts
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';

bootstrapApplication(AppComponent, {
  providers: [
    provideRouter([
      { path: '', component: HomeComponent },
      { path: 'products', loadComponent: () =>
          import('./product-list/product-list.component').then(c => c.ProductListComponent) }
    ]),
    provideHttpClient(),
    provideStore(),     // NgRx
    provideEffects(),   // NgRx Effects
  ]
});
```

---

### Q4. What is the new control flow syntax in Angular 17?

```html
<!-- OLD (still works): NgIf directive -->
<div *ngIf="isLoggedIn; else notLoggedIn">
  <p>Welcome!</p>
</div>
<ng-template #notLoggedIn><p>Please log in.</p></ng-template>

<!-- NEW: @if block — cleaner, built into template compiler (no directive import!) -->
@if (isLoggedIn) {
  <p>Welcome!</p>
} @else if (isPending) {
  <p>Verifying your account...</p>
} @else {
  <p>Please log in.</p>
}

<!-- OLD: *ngFor -->
<li *ngFor="let item of items; trackBy: trackById">{{ item.name }}</li>

<!-- NEW: @for with built-in track (required for performance) -->
@for (item of items; track item.id) {
  <li>{{ item.name }}</li>
} @empty {
  <li>No items found.</li>
}

<!-- OLD: *ngSwitch -->
<div [ngSwitch]="status">
  <p *ngSwitchCase="'active'">Active</p>
  <p *ngSwitchDefault>Unknown</p>
</div>

<!-- NEW: @switch -->
@switch (status) {
  @case ('active') { <p>Active</p> }
  @case ('pending') { <p>Pending</p> }
  @default { <p>Unknown</p> }
}
```

**Benefits of new control flow:**
- No need to import `CommonModule` for `*ngIf`/`*ngFor`
- `@for` **requires** `track` — enforces performance best practice
- Built into the Angular template compiler — faster change detection
- `@empty` for `@for` eliminates the need for `*ngIf="items.length === 0"` pattern

---

### Q5. What is `toSignal()` and `toObservable()`? Bridging Signals and RxJS

```typescript
import { toSignal, toObservable } from '@angular/core/rxjs-interop';

@Component({
  standalone: true,
  template: `
    <p>{{ products() | json }}</p>
    <p>Count: {{ count() }}</p>
  `
})
export class ProductListComponent {
  // Convert Observable → Signal (no subscribe/unsubscribe needed in template)
  products = toSignal(
    this.productService.getProducts(),
    { initialValue: [] }
  );

  // Computed signal based on signal from Observable
  count = computed(() => this.products().length);

  // Convert Signal → Observable (when you need RxJS operators)
  private searchTerm = signal('');
  searchResults$ = toObservable(this.searchTerm).pipe(
    debounceTime(300),
    distinctUntilChanged(),
    switchMap(term => this.productService.search(term))
  );

  constructor(private productService: ProductService) {}
}
```

---

### Q6. What is Zoneless Angular? (Angular 18+)

```typescript
// Traditional Angular: Zone.js triggers CD after every async event
// Zoneless Angular: Components opt in to fine-grained reactivity via Signals

// bootstrapApplication with zoneless (experimental in Angular 18)
bootstrapApplication(AppComponent, {
  providers: [
    // Remove Zone.js — signals handle CD
    provideExperimentalZonelessChangeDetection()
  ]
});

// Components MUST use Signals or async pipe for updates to be detected
@Component({
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush,  // required for zoneless
  template: `<p>Count: {{ count() }}</p>`
})
export class ZonelessComponent {
  count = signal(0);  // changes to this signal trigger CD automatically

  increment(): void {
    this.count.update(c => c + 1);  // Angular detects and re-renders
  }
}
```
