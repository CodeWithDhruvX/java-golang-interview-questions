# 📘 02 — Directives, Services & Dependency Injection
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Built-in structural directives: `*ngIf`, `*ngFor`, `*ngSwitch`
- Built-in attribute directives: `NgClass`, `NgStyle`
- Creating custom attribute directives
- What services are and why they exist
- `@Injectable` and `providedIn: 'root'`
- Constructor injection and the DI hierarchy

---

## ❓ Most Asked Questions

### Q1. What are the types of directives in Angular?

Angular has **three types** of directives:

| Type | Description | Examples |
|------|-------------|---------|
| **Component** | Directive with its own template | Every `@Component` |
| **Structural** | Adds or removes DOM elements | `*ngIf`, `*ngFor`, `*ngSwitch` |
| **Attribute** | Changes appearance/behavior of an element | `[NgClass]`, `[NgStyle]`, custom `[highlight]` |

---

### Q2. How do `*ngIf` and `*ngFor` work?

```html
<!-- *ngIf — adds/removes DOM element -->
<div *ngIf="isLoggedIn; else loginBlock">
  <h2>Welcome back, {{ username }}!</h2>
</div>
<ng-template #loginBlock>
  <p>Please log in.</p>
</ng-template>

<!-- *ngIf with async pipe -->
<div *ngIf="user$ | async as user">
  <p>{{ user.name }}</p>
</div>

<!-- *ngFor — iterates over a list -->
<ul>
  <li *ngFor="let product of products; let i = index; trackBy: trackById">
    {{ i + 1 }}. {{ product.name }} — ₹{{ product.price }}
  </li>
</ul>
```

```typescript
// trackBy improves performance — avoids full DOM re-render on list update
trackById(index: number, item: Product): number {
  return item.id;
}
```

> **Important:** `*ngIf="false"` **removes** the element from DOM. `[hidden]="true"` only hides it with CSS. Use `*ngIf` when you don't need the element; use `[hidden]` when toggling frequently.

---

### Q3. How do `NgClass` and `NgStyle` work?

```html
<!-- NgClass — conditionally apply CSS classes -->
<div [ngClass]="{
  'active': isActive,
  'disabled': isDisabled,
  'highlight': score > 90
}">Status</div>

<!-- NgClass with array -->
<div [ngClass]="['card', isLarge ? 'card-lg' : 'card-sm']">...</div>

<!-- NgStyle — conditionally apply inline styles -->
<p [ngStyle]="{
  'color': isError ? 'red' : 'green',
  'font-size.px': fontSize,
  'font-weight': isBold ? 'bold' : 'normal'
}">
  Status message
</p>
```

---

### Q4. How do you create a custom attribute directive?

```typescript
// highlight.directive.ts
import { Directive, ElementRef, HostListener, Input, OnInit } from '@angular/core';

@Directive({
  selector: '[appHighlight]'   // applied as [appHighlight] in templates
})
export class HighlightDirective implements OnInit {
  @Input('appHighlight') highlightColor: string = 'yellow';
  @Input() defaultColor: string = 'transparent';

  constructor(private el: ElementRef) {}

  ngOnInit(): void {
    this.el.nativeElement.style.backgroundColor = this.defaultColor;
  }

  @HostListener('mouseenter') onMouseEnter(): void {
    this.el.nativeElement.style.backgroundColor = this.highlightColor;
  }

  @HostListener('mouseleave') onMouseLeave(): void {
    this.el.nativeElement.style.backgroundColor = this.defaultColor;
  }
}
```

```html
<!-- Usage in template -->
<p [appHighlight]="'lightblue'" defaultColor="'lightyellow'">
  Hover over me!
</p>
```

---

### Q5. What is a Service in Angular and why is it used?

A **service** is a class that provides shared logic, data, or functionality that is not tied to a specific view. Services follow the **Single Responsibility Principle** — keep business logic out of components.

**Common service use cases:**
- HTTP API calls (`UserService`, `ProductService`)
- Shared state management
- Authentication and authorization
- Logging and error handling
- Caching

```typescript
@Injectable({
  providedIn: 'root'  // singleton — one instance for the whole app
})
export class CartService {
  private cartItems: CartItem[] = [];

  addItem(product: Product, quantity: number): void {
    const existing = this.cartItems.find(i => i.productId === product.id);
    if (existing) {
      existing.quantity += quantity;
    } else {
      this.cartItems.push({ productId: product.id, product, quantity });
    }
  }

  getItems(): CartItem[] {
    return [...this.cartItems];  // return copy to prevent external mutation
  }

  getTotal(): number {
    return this.cartItems.reduce((sum, i) => sum + i.product.price * i.quantity, 0);
  }

  clearCart(): void {
    this.cartItems = [];
  }
}
```

---

### Q6. What is Dependency Injection (DI) in Angular?

**Dependency Injection** is a design pattern where a class receives its dependencies from external sources rather than creating them itself. Angular has a built-in DI system.

```typescript
// WITHOUT DI — tightly coupled, hard to test
class OrderComponent {
  private service = new OrderService();  // creates its own dependency ❌
}

// WITH DI — loosely coupled, fully testable
@Component({ selector: 'app-order', template: '...' })
export class OrderComponent implements OnInit {
  orders: Order[] = [];

  // Angular's DI injects OrderService automatically ✅
  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    this.orderService.getOrders().subscribe(o => this.orders = o);
  }
}
```

---

### Q7. What is `providedIn: 'root'` vs providing in a module?

```typescript
// 1. providedIn: 'root' — Singleton across the entire app (recommended)
@Injectable({ providedIn: 'root' })
export class AuthService { }

// 2. Provided in NgModule — Singleton within that module's scope
@NgModule({
  providers: [ReportService]  // available to all components in this module
})
export class ReportsModule { }

// 3. Provided in a Component — New instance per component
@Component({
  selector: 'app-product-form',
  providers: [ValidationService]  // fresh instance for each ProductFormComponent
})
export class ProductFormComponent { }
```

| Scope | How | Effect |
|-------|-----|--------|
| App-wide singleton | `providedIn: 'root'` | One instance everywhere |
| Module-scoped | `providers` in `@NgModule` | Shared within module |
| Component-scoped | `providers` in `@Component` | New instance per component |

---

### Q8. What is `*ngSwitch`?

```html
<!-- *ngSwitch — cleaner than multiple *ngIf for exclusive conditions -->
<div [ngSwitch]="orderStatus">
  <p *ngSwitchCase="'pending'">⏳ Your order is being processed</p>
  <p *ngSwitchCase="'shipped'">🚚 Your order is on the way!</p>
  <p *ngSwitchCase="'delivered'">✅ Order delivered successfully</p>
  <p *ngSwitchCase="'cancelled'">❌ Order was cancelled</p>
  <p *ngSwitchDefault>ℹ️ Unknown status</p>
</div>
```

> Prefer `*ngSwitch` over multiple `*ngIf` when switching between mutually exclusive views based on a single variable.

---

### Q9. How does Angular's hierarchical DI tree work?

Angular's injectors form a **tree** — child injectors inherit from parent injectors.

```
AppModule Injector (root)
  └── FeatureModule Injector
        └── ComponentA Injector
              └── ComponentB Injector
```

- Angular walks **up the tree** to find the first provider that matches the requested token
- If found at root → same singleton instance returned everywhere
- If component provides its own → that component (and its children) get a local instance
- If no provider found → `NullInjectorError` at runtime

```typescript
// Force a new local instance instead of the root one
@Component({
  providers: [{ provide: DataService, useClass: LocalDataService }]
})
export class ChildComponent { }
```

---

### Q10. What is `@HostListener` and `@HostBinding`?

```typescript
@Directive({ selector: '[appAutoGrow]' })
export class AutoGrowDirective {
  // @HostBinding — binds a property to the host element
  @HostBinding('style.height') height: string = 'auto';

  // @HostListener — listens to events on the host element
  @HostListener('input', ['$event.target'])
  onInput(target: HTMLTextAreaElement): void {
    target.style.height = 'auto';
    target.style.height = target.scrollHeight + 'px';  // grow with content
  }

  @HostBinding('class.focused') isFocused = false;

  @HostListener('focus') onFocus() { this.isFocused = true; }
  @HostListener('blur') onBlur() { this.isFocused = false; }
}
```

```html
<!-- Usage -->
<textarea appAutoGrow rows="3">Type here...</textarea>
```
