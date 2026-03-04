# 📘 04 — Angular Internals & Compiler
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Ivy compiler internals and how it differs from View Engine
- Zone.js — what it does and why Angular needs it
- Change detection algorithm — how Angular traverses the component tree
- LView (Logical View) data structure
- Ahead-of-Time (AOT) vs Just-in-Time (JIT) compilation
- Metadata reflection and decorators

---

## ❓ Most Asked Questions

### Q1. What is the Ivy compiler and how does it improve Angular?

**Ivy** (introduced as default in Angular 9) is Angular's next-generation compilation and rendering engine, replacing the older **View Engine**.

| Feature | View Engine | Ivy |
|---------|------------|-----|
| Bundle size | Larger (runtime interpreter) | Smaller (truly tree-shakeable) |
| Compilation | Component needs NgModule context | Local compilation (component-level) |
| Debugging | Hard to inspect rendered output | Templates compile to readable JS |
| Lazy loading | Module-level | Component-level (Standalone) |
| Error checking | Some at runtime | More at compile time |
| Library compilation | Required at publish time | Works with pre-compiled libs |

**How Ivy works:**

```typescript
// You write this:
@Component({
  selector: 'app-greeting',
  template: '<h1>Hello, {{ name }}!</h1>'
})
export class GreetingComponent {
  name = 'Angular';
}

// Ivy compiles it to something like (simplified):
class GreetingComponent {
  name = 'Angular';

  // Generated static component definition
  static ɵcmp = defineComponent({
    type: GreetingComponent,
    selectors: [['app-greeting']],
    template: function(rf, ctx) {
      if (rf & 1) {  // CREATE phase
        ɵɵelementStart(0, 'h1');
        ɵɵtext(1);
        ɵɵelementEnd();
      }
      if (rf & 2) {  // UPDATE phase
        ɵɵadvance(1);
        ɵɵtextInterpolate1('Hello, ', ctx.name, '!');
      }
    }
  });
}
```

The template compiles to **incremental DOM instructions** — only the DOM that actually changed is updated.

---

### Q2. What is Zone.js? What does it do for Angular?

**Zone.js** is a library that **patches browser async APIs** to intercept them and notify Angular when asynchronous operations complete.

```typescript
// Zone.js intercepts ALL of these:
setTimeout, setInterval, requestAnimationFrame,
fetch, XMLHttpRequest,
Promise.then,
addEventListener (click, keyup, etc.)

// Without Zone.js:
// - User clicks button → HTTP call finishes → Angular doesn't know → DOM not updated!

// With Zone.js:
// - User clicks (zone intercepts) → Angular runs CD → HTTP call finishes → 
//   zone intercepts Promise resolution → Angular runs CD again → DOM updated ✅
```

```typescript
// Opting out of Zone.js for performance-critical code
@Component({ ... })
export class HighFrequencyComponent {
  constructor(private ngZone: NgZone) {}

  ngAfterViewInit(): void {
    // Run scroll listener OUTSIDE Angular — prevents CD on every scroll event
    this.ngZone.runOutsideAngular(() => {
      fromEvent(window, 'scroll').pipe(
        throttleTime(100),
        filter(() => this.shouldUpdate())
      ).subscribe(() => {
        this.ngZone.run(() => {
          this.updateVisibleItems();  // re-enter Angular's zone for the actual update
        });
      });
    });
  }
}
```

---

### Q3. Explain Angular's change detection algorithm.

```
Change Detection Trigger
         ↓
Angular traverses component tree TOP-DOWN
         ↓
For each component:
  1. Check @Input() bindings for reference changes
  2. Evaluate all template binding expressions
  3. If binding value changed: update DOM
  4. Recurse into children
         ↓
Repeat for ALL components (Default strategy)
OR only dirty-flagged components (OnPush strategy)
```

```typescript
// BEFORE CD runs in your component, this happens:
// 1. Zone.js detected an async operation completed
// 2. ApplicationRef.tick() is called
// 3. Angular starts at AppComponent and traverses down

// Default strategy: Component is checked EVERY tick
@Component({ changeDetection: ChangeDetectionStrategy.Default })

// OnPush: Component is skipped UNLESS:
// - @Input() reference changed
// - async pipe emitted
// - Internal event fired
// - markForCheck() called
@Component({ changeDetection: ChangeDetectionStrategy.OnPush })

// Manually detach from CD entirely (for static views)
@Component({ ... })
export class StaticHeaderComponent implements OnInit {
  constructor(private cdr: ChangeDetectorRef) {}

  ngOnInit(): void {
    this.cdr.detach();  // Angular never checks this component
    // Only re-attach when data actually changes
  }

  updateData(newData: Data): void {
    this.data = newData;
    this.cdr.detectChanges();  // manually trigger CD just for this component
  }
}
```

---

### Q4. What is `ExpressionChangedAfterItHasBeenCheckedError`?

```typescript
// This error only appears in DEVELOPMENT mode
// It means: Angular ran CD, the view updated, then Angular ran CD again 
// (its second verification pass) and found the value had CHANGED

// ❌ Common mistake: setting a value in ngAfterViewInit that affects the template
@Component({
  template: `<h1>{{ title }}</h1>`
})
export class BadComponent implements AfterViewInit {
  title = 'Initial';

  ngAfterViewInit(): void {
    this.title = 'Changed!';  // ❌ ExpressionChangedAfterItHasBeenCheckedError
    // View was already checked, now title changed — Angular catches this inconsistency
  }
}

// ✅ Fix 1: Use setTimeout to defer to next CD cycle
ngAfterViewInit(): void {
  setTimeout(() => this.title = 'Changed!');  // runs in next CD cycle
}

// ✅ Fix 2: Call detectChanges manually
constructor(private cdr: ChangeDetectorRef) {}

ngAfterViewInit(): void {
  this.title = 'Changed!';
  this.cdr.detectChanges();  // run CD immediately for this subtree
}
```

---

### Q5. How does AOT compilation improve security?

```typescript
// JIT — templates are compiled at RUNTIME in the browser
// Risk: the Angular compiler ships to the browser (larger bundle)
// Risk: dynamic template compilation enables potential template injection attacks

// AOT — templates are compiled at BUILD TIME
// 1. Compiler never ships to browser → smaller bundle
// 2. Template expressions are statically analyzed → template injection impossible
// 3. Errors caught at build time → safer deploys

// AOT also removes dead code (tree-shaking) more effectively:
// - If a pipe is not referenced in any template, it's removed from the bundle
// - With JIT, all declared pipes must ship (referenced dynamically at runtime)

// Enabling strict AOT (strictTemplates in tsconfig)
{
  "compilerOptions": {
    "strict": true
  },
  "angularCompilerOptions": {
    "strictTemplates": true,  // type-check all template expressions
    "strictInjectionParameters": true
  }
}
```

---

### Q6. How does Angular's Dependency Injection work under the hood?

```typescript
// DI Resolution Algorithm:
// 1. Component requests a token (type) in constructor
// 2. Angular checks component's own injector (providers in @Component)
// 3. If not found → checks parent component's injector
// 4. Keeps walking up to module injector → root injector
// 5. If not found at root → NullInjectorError

@Component({
  providers: [
    LocalCacheService,  // new instance per component instance
  ]
})
export class OrderComponent {
  constructor(
    private authService: AuthService,      // from root injector (singleton)
    private cache: LocalCacheService,      // from component's own injector
    @Optional() private logger: LogService  // optional — null if not provided
  ) {}
}

// InjectionToken for non-class dependencies
export const API_URL = new InjectionToken<string>('apiUrl');

// Provide in module
{ provide: API_URL, useValue: 'https://api.example.com' }

// Inject
constructor(@Inject(API_URL) private apiUrl: string) {}
```
