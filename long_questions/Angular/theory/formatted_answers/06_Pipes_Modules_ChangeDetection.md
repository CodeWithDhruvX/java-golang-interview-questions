# 🟡 Pipes, Modules & Change Detection

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, HCL, Infosys): Built-in pipes, module types, change detection basics
> - 🚀 **Product-Based** (Google, Microsoft, Adobe): Custom pipes, OnPush strategy, zone.js internals
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🔹 Pipes

---

### 1. What are pipes in Angular? 🟢 | 🏭

"**Pipes** are template transformation functions — they take an input value, transform it, and return the formatted result for display.

```html
{{ price | currency:'INR':'symbol' }}     <!-- ₹1,200.00 -->
{{ date | date:'mediumDate' }}             <!-- Feb 20, 2026 -->
{{ username | uppercase }}                 <!-- JOHN DOE -->
{{ longText | slice:0:100 }}...            <!-- Truncated text -->
{{ jsonData | json }}                      <!-- Pretty-printed JSON -->
{{ 0.25 | percent }}                       <!-- 25% -->
```

Pipes are **composable** — I can chain them:
```html
{{ title | uppercase | slice:0:20 }}
```

They keep transformation logic out of the component class, making templates clean and testable."

#### In Depth
Built-in Angular pipes: `AsyncPipe`, `DatePipe`, `CurrencyPipe`, `DecimalPipe`, `PercentPipe`, `JsonPipe`, `SlicePipe`, `UpperCasePipe`, `LowerCasePipe`, `TitleCasePipe`, `KeyValuePipe`. The `async` pipe (`$ | async`) is the most important — it subscribes to Observables/Promises and auto-unsubscribes on component destroy, making it the cleanest way to display async data in templates.

---

### 2. What is a pure pipe? 🟡 | 🏭🚀

"A **pure pipe** only re-executes when its **input reference changes** (or primitive value changes). Angular tracks this with `===` equality.

Pure pipes are the **default** (`pure: true`). They are highly efficient because Angular skips re-running the transform if the input hasn't changed reference.

```typescript
@Pipe({ name: 'formatPrice', pure: true }) // default
export class FormatPricePipe implements PipeTransform {
  transform(value: number, currency = 'USD'): string {
    return `${currency} ${value.toFixed(2)}`;
  }
}
```

Pure pipes are **memoized** — same input always produces the same output without re-computation. They work like pure functions in functional programming."

#### In Depth
Because pure pipes use reference equality, mutating an array or object does NOT trigger the pipe to re-run — because the reference is the same. This is why `{{ items | customFilter:'active' }}` may look stale after pushing to `items`. The fix: return a **new array** from mutations (`this.items = [...this.items, newItem]`) instead of mutating in place. This is the same principle that drives **immutable state** patterns in Angular.

---

### 3. What is an impure pipe? 🟡 | 🏭🚀

"An **impure pipe** (`pure: false`) re-executes on **every change detection cycle**, regardless of input changes.

```typescript
@Pipe({ name: 'filterProducts', pure: false })
export class FilterProductsPipe implements PipeTransform {
  transform(products: Product[], status: string): Product[] {
    return products.filter(p => p.status === status);
  }
}
```

Impure pipes see mutations in arrays and objects because they run every cycle. However, they run very frequently and can cause **significant performance degradation** in large applications.

The `async` pipe (`| async`) is an impure pipe — it must check for new observable emissions on every cycle."

#### In Depth
I avoid impure pipes for computationally expensive operations. Instead, I compute the filtered/sorted result **in the component class** and store it as a property:

```typescript
// Component class
get activeProducts(): Product[] {
  return this.products.filter(p => p.status === 'active');
}
```

This runs only when explicitly changed, not on every CD cycle. For complex data needs, I use `memoization` libraries or RxJS operators (`map`, `distinctUntilChanged`) to compute derived state reactively.

---

### 4. How to create a custom pipe? 🟡 | 🏭🚀

"Custom pipes are simple: implement `PipeTransform` and use the `@Pipe` decorator:

```typescript
// CLI: ng generate pipe truncate
@Pipe({
  name: 'truncate',
  standalone: true // Angular 14+
})
export class TruncatePipe implements PipeTransform {
  transform(value: string, maxLength = 50, suffix = '...'): string {
    if (!value) return '';
    if (value.length <= maxLength) return value;
    return value.substring(0, maxLength - suffix.length) + suffix;
  }
}
```

```html
<!-- Usage -->
{{ article.content | truncate:100:'...read more' }}
```

Real examples I've built: `timeAgo` (converts date to '5 minutes ago'), `fileSize` (bytes to KB/MB/GB), `highlight` (wraps search terms in `<mark>` tags)."

#### In Depth
Custom pipes can take **multiple arguments** separated by colons:

```html
{{ date | formatDate:'DD/MM/YYYY':'UTC' }}
```

These map to the subsequent parameters in `transform(value, format, timezone)`. Pipes are also fully **testable in isolation** — just create an instance and call `.transform()` directly in a Jasmine/Jest `describe` block without needing a `TestBed` setup. This is much faster than testing display logic embedded in component templates.

---

## 🔹 Angular Modules

---

### 5. What is the root module? 🟢 | 🏭

"The **root module** (`AppModule`) is the entry point that bootstraps the Angular application. Every Angular app has exactly one root module.

```typescript
@NgModule({
  declarations: [AppComponent, HeaderComponent, FooterComponent],
  imports: [
    BrowserModule,          // Required for browser apps
    AppRoutingModule,       // Root routing config
    HttpClientModule,       // Global HTTP support
    SharedModule,           // Common components/pipes/directives
  ],
  providers: [AuthService], // App-wide services
  bootstrap: [AppComponent] // Component rendered in index.html
})
export class AppModule {}
```

`BrowserModule` must be imported **only** in the root module (not in feature modules). Feature modules import `CommonModule` instead."

#### In Depth
With **Standalone Components** (Angular 15+ recommended approach), you no longer need `AppModule`. The application is bootstrapped directly:

```typescript
// main.ts
bootstrapApplication(AppComponent, {
  providers: [
    provideRouter(routes),
    provideHttpClient(),
    provideAnimations(),
  ]
});
```

This is simpler, more tree-shakable, and aligns with the future direction of Angular. `NgModule`-based apps continue to work but new projects should prefer standalone APIs.

---

### 6. What is a feature module? 🟢 | 🏭

"A **feature module** groups all components, services, directives, and pipes related to a specific **business feature** of the app.

```typescript
@NgModule({
  declarations: [ProductListComponent, ProductDetailComponent, ProductEditComponent],
  imports: [CommonModule, ReactiveFormsModule, ProductsRoutingModule],
  exports: [ProductListComponent] // Only what other modules need
})
export class ProductsModule {}
```

Benefits:
- Encapsulation — internals not visible to other modules
- **Lazy loading** — `ProductsModule` is only loaded when user navigates to `/products`
- **Team ownership** — different teams own different feature modules
- Faster builds with module-level caching"

#### In Depth
Feature modules establish **boundary enforcement** in large apps. Exporting only what's absolutely necessary from a feature module prevents **accidental coupling**. If `ProductDetailComponent` is never exported, other modules can't accidentally depend on it — any such dependency must go through the public API (exported components). This is similar to `public`/`private` access modifiers in OOP, but at the module level.

---

### 7. What is SharedModule? 🟢 | 🏭

"A **SharedModule** contains components, directives, and pipes that are used across **multiple feature modules** — things like a `LoadingSpinnerComponent`, `TruncatePipe`, `ConfirmDialogComponent`.

```typescript
@NgModule({
  declarations: [LoadingSpinnerComponent, TruncatePipe, HighlightDirective],
  imports: [CommonModule],
  exports: [
    CommonModule,          // Re-export so importing modules get it too
    LoadingSpinnerComponent,
    TruncatePipe,
    HighlightDirective,
  ]
})
export class SharedModule {}
```

**Important**: Never put **services** in SharedModule's `providers` — this creates multiple instances. Services should be declared with `providedIn: 'root'` or injected via the DI hierarchy."

#### In Depth
The **SharedModule + CoreModule pattern** is a classic Angular architecture:

- **`SharedModule`** — Stateless UI components, pipes, directives (multiply imported, safely because no providers)
- **`CoreModule`** — App-wide singleton services, global state, root guards (imported **once** in `AppModule` only)

`CoreModule` has a guard constructor to throw an error if it's imported more than once, preventing service duplication in lazy-loaded modules. With standalone components, both patterns are simplified — services use `providedIn: 'root'` and shared UI components import each other directly.

---

## 🔹 Change Detection

---

### 8. How does change detection work in Angular? 🟡 | 🏭🚀

"**Change detection** is Angular's mechanism for keeping the view synchronized with the component data. Whenever a change might have occurred (event, HTTP response, timer), Angular runs change detection — it checks all bound expressions in component templates and updates the DOM if values changed.

Angular uses **zone.js** to know when to trigger change detection. Zone.js patches browser APIs (click events, setTimeout, Promise resolution, HTTP) and notifies Angular after each async operation completes.

By default, Angular runs change detection on the **entire component tree** (Default strategy) — starting from the root and checking every component."

#### In Depth
Angular's change detection is built on a **linear view tree** (not a recursive object graph). Each component gets a **LView** (Logical View). During a CD run, Angular iterates through all LViews in **top-down order**, evaluating each binding expression. If all component inputs use primitives or immutable objects, Angular can skip entire subtrees with `OnPush`. This is why CD is actually very fast — it's simple array iteration, not deep object comparison.

---

### 9. What is ChangeDetectionStrategy? 🟡 | 🏭🚀

"**`ChangeDetectionStrategy`** is an enum that controls how angular decides when to run change detection for a component.

There are two strategies:

**`Default` (CheckAlways):**
- Angular checks this component on **every CD cycle**
- Triggered by any event, HTTP call, or timer anywhere in the app
- Safe but potentially slow for complex UIs

**`OnPush`:**
- Angular only runs CD for this component when:
  1. An `@Input()` reference changes
  2. A DOM event occurs in this component
  3. An Observable used with `async` pipe emits
  4. `markForCheck()` or `detectChanges()` is called explicitly
- Much more performant for leaf/pure components

```typescript
@Component({
  selector: 'app-product-card',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `<h2>{{ product.name }}</h2>`
})
export class ProductCardComponent {
  @Input() product: Product; // OnPush checks reference equality
}
```"

#### In Depth
`OnPush` is not a silver bullet — it can cause bugs if not used carefully. If you mutate an `@Input()` object/array instead of replacing it with a new reference, `OnPush` won't detect the change. This forces an **immutable data flow** discipline: always return new references from operations that change data. The combination of `OnPush` + immutable state + `async` pipe is the highest-performance Angular pattern I use in production for data-heavy dashboards.

---

### 10. Difference between Default and OnPush strategy? 🟡 | 🏭🚀

| Aspect | Default | OnPush |
|---|---|---|
| CD triggers | Every cycle | Input ref change / Event / Observable |
| Performance | O(n) all components | O(k) only dirty components |
| Data mutation | Detects all changes | Only detects reference changes |
| Complexity | Low | Medium (requires immutable patterns) |
| Use case | Quick prototyping | Performance-critical components |

"In a large list component with `*ngFor` rendering 1000 product cards:

- **Default**: All 1000 cards check their bindings every click anywhere in the app
- **OnPush**: Only cards whose input references changed re-check

This can mean the difference between a 50ms render cycle (Default) and a 5ms render (OnPush) for data-heavy pages. I apply `OnPush` to all **pure presentational components** (those that only receive `@Input()` and emit `@Output()`)."

#### In Depth
`ChangeDetectorRef.markForCheck()` marks the component and all its ancestors as dirty, scheduling them for the next CD cycle. `ChangeDetectorRef.detectChanges()` runs CD immediately for the component and its descendants (synchronously). I use `detectChanges()` when I need to update the DOM immediately within a `setTimeout` or other non-Angular async operation running outside the zone.

---

### 11. What is Zone.js and how does it affect change detection? 🔴 | 🚀

"**Zone.js** is a library that creates an **execution context** (zone) that tracks asynchronous operations. Angular uses it to automatically trigger change detection after any async task completes.

Zone.js patches:
- Browser events (`click`, `input`, `keydown`, etc.)
- `setTimeout` / `setInterval`
- Promise resolution
- HTTP requests (via XMLHttpRequest)

When any of these complete, zone.js notifies Angular, which then runs change detection.

```typescript
// Running outside Angular zone (no CD triggered)
constructor(private ngZone: NgZone) {}

setupHighFrequencyListener(): void {
  this.ngZone.runOutsideAngular(() => {
    document.addEventListener('mousemove', (e) => {
      // This won't trigger CD — performance win for mousemove events
      this.updateCursorPosition(e.clientX, e.clientY);
    });
  });
}

// Bring back into zone when you need CD
updateSomethingThatAffectsUI(): void {
  this.ngZone.run(() => {
    this.position = this.latestPosition; // Triggers CD
  });
}
```"

#### In Depth
The future of Angular is **zoneless** (Angular 17+ experimental). Without zone.js, Angular won't automatically detect async changes — components must use `Signals` or explicit `markForCheck()` calls. Zoneless Angular is significantly faster because removing zone.js eliminates all the monkey-patching overhead and reduces the number of spurious CD cycles triggered by unrelated async operations. Signals are designed specifically to enable this zoneless future.

---
