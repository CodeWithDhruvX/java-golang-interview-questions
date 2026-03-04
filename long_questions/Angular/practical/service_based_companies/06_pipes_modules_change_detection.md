# 📘 06 — Pipes, Modules & Change Detection
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Built-in pipes: `date`, `currency`, `uppercase/lowercase`, `async`, `json`, `slice`, `decimal`
- Creating custom `PipeTransform` pipes
- Pure vs impure pipes
- `NgModule`: `declarations`, `imports`, `exports`, `providers`
- Change detection basics: Default vs `OnPush`

---

## ❓ Most Asked Questions

### Q1. What are Angular pipes? Show common built-in pipes.

**Pipes** transform data in templates without modifying the original data.

```html
<!-- Date pipe -->
<p>{{ today | date }}</p>                        <!-- Mar 3, 2026 -->
<p>{{ today | date: 'dd/MM/yyyy' }}</p>          <!-- 03/03/2026 -->
<p>{{ today | date: 'EEEE, MMMM d, y' }}</p>    <!-- Tuesday, March 3, 2026 -->

<!-- Currency pipe -->
<p>{{ price | currency }}</p>                    <!-- $1,234.56 -->
<p>{{ price | currency: 'INR': 'symbol' }}</p>  <!-- ₹1,234.56 -->
<p>{{ price | currency: 'EUR': 'code': '1.0-0' }}</p> <!-- EUR 1,235 -->

<!-- Number / Decimal pipe -->
<p>{{ 3.14159 | number: '1.0-2' }}</p>           <!-- 3.14 -->
<p>{{ 1234567.89 | number }}</p>                 <!-- 1,234,567.89 -->

<!-- String pipes -->
<p>{{ 'hello world' | uppercase }}</p>    <!-- HELLO WORLD -->
<p>{{ 'HELLO' | lowercase }}</p>          <!-- hello -->
<p>{{ 'angular' | titlecase }}</p>        <!-- Angular -->

<!-- Slice pipe -->
<p>{{ [1,2,3,4,5] | slice: 1:4 }}</p>    <!-- [2, 3, 4] -->
<p>{{ 'Hello World' | slice: 0:5 }}</p>  <!-- Hello -->

<!-- JSON pipe (debugging) -->
<pre>{{ user | json }}</pre>

<!-- Percent -->
<p>{{ 0.85 | percent }}</p>              <!-- 85% -->
<p>{{ 0.856 | percent: '1.0-1' }}</p>   <!-- 85.6% -->

<!-- Async pipe — subscribes to Observable or Promise -->
<p>{{ user$ | async | json }}</p>
```

---

### Q2. How do you create a custom pipe?

```typescript
// pipes/truncate.pipe.ts
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'truncate',
  pure: true  // default — only recalculates when input reference changes
})
export class TruncatePipe implements PipeTransform {
  transform(value: string, limit: number = 50, ellipsis: string = '...'): string {
    if (!value) return '';
    if (value.length <= limit) return value;
    return value.substring(0, limit).trimEnd() + ellipsis;
  }
}
```

```html
<!-- Usage -->
<p>{{ product.description | truncate: 100 }}</p>
<p>{{ article.body | truncate: 200: ' [read more]' }}</p>
```

**Another example — filter pipe:**

```typescript
@Pipe({ name: 'filterBy' })
export class FilterByPipe implements PipeTransform {
  transform(items: any[], key: string, value: any): any[] {
    if (!items || !key || value === undefined) return items;
    return items.filter(item => item[key] === value);
  }
}
```

```html
<li *ngFor="let p of products | filterBy: 'category': selectedCategory">
  {{ p.name }}
</li>
```

---

### Q3. What is the difference between pure and impure pipes?

| Feature | Pure Pipe | Impure Pipe |
|---------|-----------|------------|
| Called when | Input reference changes | Every change detection cycle |
| Performance | Fast — memoized | Slow — runs very frequently |
| Array/Object mutations | Won't detect | Detects internal changes |
| Default | Yes (`pure: true`) | `pure: false` |

```typescript
// ✅ Pure pipe — only recalculates when value/args change by reference
@Pipe({ name: 'formatDate', pure: true })
export class FormatDatePipe implements PipeTransform {
  transform(date: Date, format: string): string {
    return formatDate(date, format, 'en-US');
  }
}

// ⚠️ Impure pipe — recalculates every change detection cycle (use sparingly)
@Pipe({ name: 'filter', pure: false })
export class FilterPipe implements PipeTransform {
  transform(items: Product[], category: string): Product[] {
    return items.filter(p => p.category === category);
  }
}
```

> **Best practice:** Use the `async` pipe (which is impure) for Observables. For list filtering, prefer filtering in the component class and binding to a plain array instead of using an impure pipe.

---

### Q4. What is `NgModule`? Explain `declarations`, `imports`, `exports`, `providers`.

```typescript
@NgModule({
  // Components, directives, and pipes that BELONG TO this module
  declarations: [
    ProductListComponent,
    ProductCardComponent,
    TruncatePipe,
    HighlightDirective
  ],

  // Other modules whose EXPORTED components/pipes/directives this module needs
  imports: [
    CommonModule,        // *ngIf, *ngFor, DatePipe, etc.
    ReactiveFormsModule, // Reactive forms support
    RouterModule,
    HttpClientModule,
    SharedModule         // our own shared module
  ],

  // Components/pipes/directives from this module that OTHER modules can use
  exports: [
    ProductListComponent,  // other modules can use <app-product-list>
    TruncatePipe           // other modules can use | truncate
  ],

  // Services provided at module level (prefer providedIn: 'root' for most services)
  providers: [
    ProductService,
    { provide: LOCALE_ID, useValue: 'en-IN' }
  ]
})
export class ProductsModule {}
```

---

### Q5. What is Change Detection in Angular?

**Change detection** is how Angular checks if the component's data has changed and whether the DOM needs to be updated.

**Default strategy:**
- Angular checks **every component** in the component tree after any async event (click, HTTP response, timer)
- Uses **zone.js** to intercept async operations and trigger change detection

```typescript
// Default — Angular checks this component on every event anywhere in the app
@Component({
  selector: 'app-product-list',
  changeDetection: ChangeDetectionStrategy.Default
})
export class ProductListComponent {}

// OnPush — Angular only checks when:
// 1. An @Input() reference changes
// 2. An event is triggered inside the component
// 3. An Observable subscribed with async pipe emits
// 4. markForCheck() is called manually
@Component({
  selector: 'app-product-card',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProductCardComponent {
  @Input() product: Product;  // must pass new reference to trigger update
}
```

---

### Q6. When should you use `ChangeDetectionStrategy.OnPush`?

```typescript
// Use OnPush for:
// ✅ "Presentational" / "dumb" components that only receive @Input() data
// ✅ Components in lists (many instances — big performance gain)
// ✅ Components with heavy change detection cost

@Component({
  selector: 'app-order-item',
  changeDetection: ChangeDetectionStrategy.OnPush,  // ✅
  template: `
    <div class="order-item">
      <span>{{ order.productName }}</span>
      <span>{{ order.quantity }}</span>
      <span>{{ order.total | currency: 'INR' }}</span>
      <button (click)="cancelOrder()">Cancel</button>
    </div>
  `
})
export class OrderItemComponent {
  @Input() order: Order;  // Pass new object reference from parent to trigger update
  @Output() cancelled = new EventEmitter<string>();

  constructor(private cdr: ChangeDetectorRef) {}

  cancelOrder(): void {
    this.cancelled.emit(this.order.id);
    // No need to call detectChanges here — event binding triggers it
  }

  // For manual update from external async source
  updateFromWebSocket(data: Order): void {
    this.order = data;
    this.cdr.markForCheck();  // Notify Angular to check this component
  }
}
```

---

### Q7. What is `SharedModule` and why is it used?

```typescript
// shared/shared.module.ts
@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [
    // Reusable components
    LoadingSpinnerComponent,
    ConfirmDialogComponent,
    PaginationComponent,
    // Reusable pipes
    TruncatePipe,
    TimeAgoPipe,
    // Reusable directives
    HighlightDirective,
    ClickOutsideDirective
  ],
  exports: [
    // Re-export Angular modules other feature modules always need
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    // Export all our own declarations too
    LoadingSpinnerComponent,
    ConfirmDialogComponent,
    PaginationComponent,
    TruncatePipe,
    TimeAgoPipe,
    HighlightDirective,
    ClickOutsideDirective
  ]
})
export class SharedModule {}
```

> **Rule:** Never import `SharedModule` into `AppModule`. Import it in feature modules that need the shared components.
