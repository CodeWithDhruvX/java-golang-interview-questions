# 📘 01 — Angular Basics & Components
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy

---

## 🔑 Must-Know Topics
- What Angular is and how it differs from AngularJS
- Component structure (`@Component`, template, class, styles)
- Lifecycle hooks: `ngOnInit`, `ngOnDestroy`, `ngOnChanges`
- Data binding: interpolation, property, event, two-way
- `@Input()` and `@Output()` for component communication
- `@ViewChild` and `ElementRef`

---

## ❓ Most Asked Questions

### Q1. What is Angular? How is it different from AngularJS?

Angular (v2+) is an **open-source, TypeScript-based frontend framework** developed by Google for building Single Page Applications. It provides a complete platform: components, routing, HTTP client, form handling, and testing utilities out of the box.

**AngularJS vs Angular:**

| Feature | AngularJS (v1) | Angular (v2+) |
|---------|---------------|----------------|
| Language | JavaScript | TypeScript |
| Architecture | MVC | Component-based |
| Change Detection | Dirty-checking digest cycle | Zone.js + tree-based |
| Mobile Support | Poor | Good |
| Performance | Slower at scale | Optimized with Ivy |
| CLI | None | Powerful Angular CLI |

> AngularJS used `$scope` and `ng-model`; Angular uses `@Component` decorators and `[(ngModel)]` from `FormsModule`.

---

### Q2. What is a Component in Angular?

A **component** is the fundamental building block of an Angular application. It controls a section of the UI.

Every component has three parts:

```typescript
@Component({
  selector: 'app-user-card',       // HTML tag used in templates
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserCardComponent implements OnInit, OnDestroy {
  @Input() userId: string;
  @Output() cardClicked = new EventEmitter<string>();

  user: User;

  constructor(private userService: UserService) {}

  ngOnInit(): void {
    this.userService.getUser(this.userId).subscribe(u => this.user = u);
  }

  ngOnDestroy(): void {
    // unsubscribe if needed
  }

  onClick(): void {
    this.cardClicked.emit(this.userId);
  }
}
```

---

### Q3. What are Angular lifecycle hooks?

| Hook | When Called |
|------|------------|
| `ngOnChanges` | Before `ngOnInit` and whenever `@Input()` values change |
| `ngOnInit` | Once, after first `ngOnChanges` |
| `ngDoCheck` | Every change detection cycle |
| `ngAfterContentInit` | After `<ng-content>` projection |
| `ngAfterContentChecked` | After projected content is checked |
| `ngAfterViewInit` | After component view + child views initialized |
| `ngAfterViewChecked` | After view is checked |
| `ngOnDestroy` | Just before component is destroyed |

**Most asked hooks:**

```typescript
export class ProductComponent implements OnInit, OnChanges, OnDestroy {
  @Input() productId: string;
  private sub: Subscription;

  ngOnChanges(changes: SimpleChanges): void {
    // Fires when @Input() productId changes from parent
    if (changes['productId'] && !changes['productId'].firstChange) {
      this.loadProduct(changes['productId'].currentValue);
    }
  }

  ngOnInit(): void {
    // Best place for initial data fetching
    this.loadProduct(this.productId);
  }

  ngOnDestroy(): void {
    // Always clean up subscriptions to prevent memory leaks
    this.sub?.unsubscribe();
  }
}
```

---

### Q4. What are the four types of data binding in Angular?

```html
<!-- 1. Interpolation — Class → View (string output) -->
<h1>Welcome, {{ username }}!</h1>
<p>Total: {{ items.length }} items</p>

<!-- 2. Property Binding — Class → View (any type) -->
<img [src]="imageUrl" [alt]="imageAlt" />
<button [disabled]="isLoading">Submit</button>

<!-- 3. Event Binding — View → Class -->
<button (click)="saveRecord()">Save</button>
<input (keyup.enter)="search($event)" />

<!-- 4. Two-way Binding — requires FormsModule -->
<input [(ngModel)]="searchQuery" placeholder="Search..." />
<p>You typed: {{ searchQuery }}</p>
```

> **Key rule:** Use `[property]` for DOM properties, `(event)` for user interactions, and `[(ngModel)]` only in template-driven forms.

---

### Q5. How does `@Input()` and `@Output()` work?

```typescript
// PARENT COMPONENT
@Component({
  template: `
    <app-counter
      [count]="parentCount"
      (countChanged)="onCountChanged($event)">
    </app-counter>
    <p>Parent sees: {{ parentCount }}</p>
  `
})
export class ParentComponent {
  parentCount = 0;

  onCountChanged(newCount: number): void {
    this.parentCount = newCount;
  }
}

// CHILD COMPONENT
@Component({
  selector: 'app-counter',
  template: `
    <p>Count: {{ count }}</p>
    <button (click)="increment()">+</button>
  `
})
export class CounterComponent {
  @Input() count: number = 0;
  @Output() countChanged = new EventEmitter<number>();

  increment(): void {
    this.count++;
    this.countChanged.emit(this.count);  // Notify parent
  }
}
```

---

### Q6. What is `@ViewChild`? When do you use it?

```typescript
@Component({
  template: `
    <input #searchInput type="text" />
    <app-chart #chart></app-chart>
  `
})
export class SearchComponent implements AfterViewInit {
  @ViewChild('searchInput') inputRef: ElementRef;
  @ViewChild('chart') chartComponent: ChartComponent;

  ngAfterViewInit(): void {
    // ✅ DOM is ready here — safe to access ViewChild
    this.inputRef.nativeElement.focus();
    this.chartComponent.render();

    // ❌ NOT in ngOnInit — ViewChild is undefined there!
  }
}
```

> Use `@ViewChild` to access a child component, directive, or DOM element reference from the parent. Always access it in `ngAfterViewInit`, not `ngOnInit`.

---

### Q7. What is the Angular CLI and what are common commands?

**Angular CLI** is a command-line tool for scaffolding, building, testing, and deploying Angular apps.

```bash
# Create new project
ng new my-app --routing --style=scss

# Generate components / services / pipes
ng generate component dashboard
ng generate service auth
ng generate pipe format-date
# Shorthand
ng g c product-list

# Development server
ng serve --open --port 4300

# Production build
ng build --configuration=production

# Run tests
ng test
ng e2e

# Update Angular version
ng update @angular/core @angular/cli
```

---

### Q8. What is `ng-content` and content projection?

```typescript
// CARD COMPONENT (with slot)
@Component({
  selector: 'app-card',
  template: `
    <div class="card">
      <div class="card-header">
        <ng-content select="[card-header]"></ng-content>  <!-- named slot -->
      </div>
      <div class="card-body">
        <ng-content></ng-content>  <!-- default slot -->
      </div>
    </div>
  `
})
export class CardComponent {}

// USAGE
<app-card>
  <h2 card-header>Product Details</h2>   <!-- goes to named slot -->
  <p>This is the product description.</p> <!-- goes to default slot -->
</app-card>
```

> Content projection allows you to build **wrapper/layout components** that accept arbitrary content from the parent.

---

### Q9. What is the difference between `constructor` and `ngOnInit`?

| Aspect | `constructor` | `ngOnInit` |
|--------|--------------|------------|
| Purpose | Dependency injection (DI) | Component initialization logic |
| When runs | When class is instantiated | After `@Input()` values are set |
| `@Input()` available | ❌ No | ✅ Yes |
| HTTP calls | ❌ Avoid | ✅ Yes |
| Best for | Injecting services, setting defaults | Fetching data, subscribing to streams |

```typescript
// ✅ Correct pattern
constructor(private userService: UserService) {}  // Only DI here

ngOnInit(): void {
  this.userService.getProfile().subscribe(p => this.profile = p);
}
```

---

### Q10. What is AOT compilation in Angular?

**AOT (Ahead-of-Time)** compilation converts Angular HTML templates and TypeScript into JavaScript at **build time** (before the browser downloads and runs the code).

| Feature | JIT (Just-in-Time) | AOT (Ahead-of-Time) |
|---------|--------------------|----------------------|
| When compiled | In the browser (runtime) | At build time |
| Bundle size | Larger (compiler included) | Smaller |
| Error detection | Runtime | Build time |
| Initial load | Slower | Faster |
| Default in prod | ❌ | ✅ (`ng build --prod`) |

```bash
# AOT is default in production
ng build --configuration=production

# Explicitly enable AOT in development
ng serve --aot
```

> AOT catches template errors (typos, undefined variables) at **build time**, making production apps faster and more reliable.
