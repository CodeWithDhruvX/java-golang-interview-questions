# 🟡 Directives, Services & Dependency Injection

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Infosys, Wipro): Definitions, syntax, standard patterns
> - 🚀 **Product-Based** (Google, Amazon, Flipkart): Internals, trade-offs, design decisions
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🔹 Directives

---

### 1. What is a structural directive? 🟢 | 🏭🚀

"A **structural directive** changes the DOM's **structure** — it can add, remove, or replace elements. These are identified by the `*` prefix in templates.

Common built-in structural directives:

```html
<div *ngIf="isLoggedIn">Welcome!</div>
<li *ngFor="let item of items; trackBy: trackById">{{ item.name }}</li>
<div [ngSwitch]="status">
  <span *ngSwitchCase="'active'">Active</span>
  <span *ngSwitchDefault>Inactive</span>
</div>
```

The `*` is a **syntactic sugar** — Angular de-sugars `*ngIf="show"` into `<ng-template [ngIf]="show">..."`. The directive receives a `TemplateRef` and `ViewContainerRef` and controls when the template is inserted into the DOM."

#### In Depth
You cannot apply **two structural directives** on the same element (e.g., `*ngIf` and `*ngFor` together). This is because each desugars to a single `<ng-template>` wrapper, and nesting two would create ambiguity about which controls the host. The solution is to wrap with `<ng-container>` which doesn't render to the DOM:

```html
<ng-container *ngIf="showList">
  <li *ngFor="let item of items">{{ item.name }}</li>
</ng-container>
```

---

### 2. What is an attribute directive? 🟢 | 🏭

"An **attribute directive** changes the **appearance or behavior** of an element, component, or another directive — without modifying the DOM structure.

Built-in examples: `NgClass`, `NgStyle`.

**Custom example:**
```typescript
@Directive({ selector: '[appHighlight]' })
export class HighlightDirective {
  constructor(private el: ElementRef, private renderer: Renderer2) {}

  @HostListener('mouseenter') onMouseEnter() {
    this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', 'yellow');
  }

  @HostListener('mouseleave') onMouseLeave() {
    this.renderer.removeStyle(this.el.nativeElement, 'backgroundColor');
  }
}
```

Usage: `<p appHighlight>Hover over me!</p>`"

#### In Depth
I always use `Renderer2` instead of directly accessing `el.nativeElement.style` in attribute directives. `Renderer2` is platform-agnostic — it works in Angular Universal (server-side rendering) where there is no real DOM. Direct DOM manipulation breaks SSR because `window`, `document`, and `HTMLElement` don't exist in the Node.js environment.

---

### 3. Explain *ngIf directive. 🟢 | 🏭

"**`*ngIf`** conditionally includes or excludes a DOM element based on a boolean expression.

```html
<div *ngIf="user; else loading">
  <p>Welcome, {{ user.name }}</p>
</div>
<ng-template #loading>
  <p>Loading...</p>
</ng-template>
```

When the expression is `false`, Angular removes the element (and its component subtree) from the DOM entirely — not just hides it with CSS. This means the component's lifecycle is also terminated (`ngOnDestroy` is called).

This is different from `[hidden]="!condition"` which keeps the element in the DOM but sets `display: none`."

#### In Depth
Because `*ngIf` destroys and recreates the component subtree, it can be an **expensive operation** for complex components. For components that toggle frequently (like tabs), consider using `[hidden]` or `display: none` via `NgClass` to avoid repeated initialization costs. Conversely, `*ngIf` is beneficial for components that hold resources (subscriptions, Web Workers) because destruction automatically cleans them up.

---

### 4. Explain *ngFor directive. 🟢 | 🏭🚀

"**`*ngFor`** renders a template for each item in an iterable.

```html
<ul>
  <li *ngFor="let item of products; let i = index; trackBy: trackByProductId">
    {{ i + 1 }}. {{ item.name }} — ${{ item.price }}
  </li>
</ul>
```

The `trackBy` function is **critical for performance**. It tells Angular how to identify items uniquely, so when the array updates, Angular can reuse existing DOM nodes instead of recreating them all.

```typescript
trackByProductId(index: number, item: Product): number {
  return item.id;
}
```"

#### In Depth
Without `trackBy`, when an array reference changes, Angular destroys **all** existing DOM nodes and recreates them — even if only one item changed. With `trackBy`, Angular compares identities and only patches the changed nodes. In large lists (100+ items), this can mean the difference between a **100ms rerender and a 10ms patch**. For Virtual Scrolling with `@angular/cdk/scrolling`, `trackBy` is mandatory for correctness.

---

### 5. How to create a custom directive? 🟡 | 🏭🚀

"I create custom directives to encapsulate reusable DOM behaviors:

```typescript
// 1. Generate via CLI
// ng g directive auto-focus

@Directive({
  selector: '[appAutoFocus]',
  standalone: true // Angular 14+
})
export class AutoFocusDirective implements AfterViewInit {
  constructor(private el: ElementRef) {}

  ngAfterViewInit(): void {
    this.el.nativeElement.focus();
  }
}
```

Usage: `<input appAutoFocus type="text" />`

I use custom directives for:
- **Auto-focus** on rendered form fields
- **Permission gates** — show/hide elements based on user roles
- **Infinite scroll** — using IntersectionObserver
- **Debounced click** — prevent double-click submissions"

#### In Depth
Directives are more testable than mixins or base classes because they can be tested in isolation with a minimal `TestBed` setup. A directive handles a **cross-cutting concern** — multiple components can use the same `[appPermission]` directive without duplicating logic. This follows the **Single Responsibility Principle**: the component handles business logic, the directive handles UI behavior.

---

## 🔹 Services & Dependency Injection

---

### 6. What are services in Angular? 🟢 | 🏭

"A **service** is a TypeScript class that encapsulates **business logic, data access, or shared state** — any logic that doesn't belong in a component.

Services typically:
- Fetch data from REST APIs (`HttpClient`)
- Share state between unrelated components
- Encapsulate business calculations
- Wrap browser APIs (`localStorage`, `navigator`)

By separating concerns this way, components stay **thin and focused on the view**, while services handle the heavy lifting. This also makes unit testing easy — I can test a service independently without rendering any UI."

#### In Depth
The distinction between a service and a component is about **UI vs Logic**. A rule I follow: if the code doesn't reference `ElementRef`, `ViewChild`, templates, or lifecycle hooks, it belongs in a service. Services should be **stateless** where possible (using stores or BehaviorSubjects for shared state), making them easier to reuse and test in isolation.

---

### 7. How do you create a service? 🟢 | 🏭

"Using CLI:

```bash
ng generate service user
# ng g s user
```

This creates a `user.service.ts` with the `@Injectable` decorator:

```typescript
@Injectable({
  providedIn: 'root' // Singleton across the app
})
export class UserService {
  private apiUrl = '/api/users';

  constructor(private http: HttpClient) {}

  getUser(id: string): Observable<User> {
    return this.http.get<User>(`${this.apiUrl}/${id}`);
  }

  updateUser(user: User): Observable<User> {
    return this.http.put<User>(`${this.apiUrl}/${user.id}`, user);
  }
}
```

The `providedIn: 'root'` makes it a **singleton**  — one instance is shared across the entire app."

#### In Depth
`providedIn: 'root'` enables **tree-shaking** of services. If the service is never injected anywhere, the Angular build pipeline removes it from the production bundle entirely. This is not possible with the old pattern of adding the service to a `NgModule`'s `providers` array, which always includes it regardless of usage.

---

### 8. What is dependency injection? 🟢 | 🏭🚀

"**Dependency Injection (DI)** is a design pattern where a class receives its dependencies from an external source rather than creating them internally.

In Angular, the **DI framework** automatically provides instances of dependencies when a class declares them in its constructor:

```typescript
// Without DI — tightly coupled, hard to test
class UserComponent {
  private service = new UserService(new HttpClient()); // ❌
}

// With DI — loosely coupled, testable
class UserComponent {
  constructor(private userService: UserService) {} // ✅ Angular injects it
}
```

For testing, I can tell Angular to inject a **mock** instead of the real service, making unit tests isolated and fast."

#### In Depth
Angular's DI uses a **hierarchical injector tree**. When a component requests a dependency, Angular walks up the injector tree (component → parent component → … → module → root). It provides the instance from the **closest injector** that has the dependency. This allows **overriding** root services with component-level implementations — a powerful pattern for tenant-specific feature overrides in multi-tenant apps.

---

### 9. How is a service provided in Angular? 🟡 | 🏭🚀

"There are **three main ways** to provide a service:

**1. Root-level (Singleton app-wide):**
```typescript
@Injectable({ providedIn: 'root' })
export class AuthService {} // One instance for the app
```

**2. Module-level (Singleton within module):**
```typescript
@NgModule({ providers: [FeatureService] })
export class FeatureModule {}
```

**3. Component-level (New instance per component):**
```typescript
@Component({
  selector: 'app-chat',
  providers: [ChatService] // New instance each time ChatComponent mounts
})
export class ChatComponent {}
```

I choose based on the required **scope**: global state uses root, feature-specific data uses module-level, and per-instance state (like a per-form validation service) uses component-level."

#### In Depth
Component-level providers create services that live and die with the component — every mount creates a new instance, every unmount destroys it. This is perfect for services that hold **component-scoped state** (like a multi-step form wizard's state). The service is also automatically garbage-collected when the component is destroyed, preventing memory leaks without manual cleanup.

---

### 10. What is the difference between `providedIn: 'root'` and `providers` array? 🟡 | 🏭🚀

"The core differences are **scope**, **singleton behavior**, and **tree-shaking**:

| Aspect | `providedIn: 'root'` | `providers: [Service]` in `@NgModule` |
|---|---|---|
| Scope | Entire application | That module (and its imports) |
| Instance | App-wide singleton | Module-wide singleton |
| Tree-shakable | ✅ Yes — removed if unused | ❌ No — always bundled |
| Lazy modules | Single instance shared | Can create separate instances per lazy module |

`providedIn: 'root'` is the modern, recommended approach for most services. I only use `providers: []` in `@NgModule` when I need module-scoped isolation or when working with legacy code."

#### In Depth
A subtle pitfall: when a **lazy-loaded module** provides a service that is also provided in root, both the root and lazy module's providers coexist in separate injectors. The lazy module component will get the lazy module's instance, while eagerly loaded components get the root instance. This can cause **state desync bugs** if you expect a shared singleton. The fix is `providedIn: 'root'` — lazy modules will share the root singleton.

---

### 11. What are advanced DI patterns? 🔴 | 🚀

"Several advanced DI patterns I use in production:

**`InjectionToken`** — For non-class dependencies:
```typescript
export const API_URL = new InjectionToken<string>('API_URL');

// In module/component providers:
{ provide: API_URL, useValue: 'https://api.example.com' }

// Injection:
constructor(@Inject(API_URL) private apiUrl: string) {}
```

**`useFactory`** — Dynamic instance creation:
```typescript
{
  provide: LoggingService,
  useFactory: (env: EnvironmentService) => {
    return env.isProd ? new NoOpLogger() : new ConsoleLogger();
  },
  deps: [EnvironmentService]
}
```

**`useExisting`** — Alias one token to another:
```typescript
{ provide: BaseService, useExisting: ConcreteService }
```

**`multi: true`** — Multiple providers for one token:
```typescript
{ provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }
```"

#### In Depth
`multi: true` providers work like a **plugin registry pattern** — multiple implementations of the same token are collected into an array. Angular uses this internally for `HTTP_INTERCEPTORS`, `APP_INITIALIZER`, and `VALIDATORS`. I use it to build extensible systems where feature modules can register plugins into a core service without the core knowing about specific implementations (Open/Closed Principle).

---

### 12. What is Self and SkipSelf in Angular DI? 🔴 | 🚀

"These decorators control **how Angular searches the injector tree** to resolve dependencies:

- **`@Self()`** — Only look in the **current component's** injector. Throws if not found.
- **`@Optional()` + `@Self()`** — Same but returns `null` if not found.
- **`@SkipSelf()`** — Skip the current injector and look in **parent injectors** instead.
- **`@Host()`** — Stop searching at the host component's injector.

```typescript
// Useful for directives that work differently depending on their context level
constructor(
  @Optional() @SkipSelf() private parentMenu: MenuComponent,
  @Self() private elementRef: ElementRef
) {}
```

I use `@SkipSelf()` in Angular Material-like component libraries where nested components need to find a parent component instance without recursive self-injection."

#### In Depth
These decorators solve the problem of **hierarchical component contexts**. For example, in a `<menu>` / `<menu-item>` pattern, `MenuItemComponent` uses `@Optional() @Host() private menu: MenuComponent` to find its parent menu. `@Host()` stops the search at the host element boundary — this is safer than `@Self()` because it respects the component boundary, not just the element boundary.

---
