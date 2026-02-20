# 🟢 Angular Basics, Data Binding & Components

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Infosys, Wipro, Accenture): Focus on definitions, syntax, and standard usage patterns
> - 🚀 **Product-Based** (Google, Microsoft, Flipkart, Swiggy): Focus on internals, trade-offs, performance impacts, and real-world scenarios
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr):** Core concepts, definitions, basic syntax
> - 🟡 **Mid-Level (2–4 yrs):** Practical patterns, common pitfalls, architecture basics
> - 🔴 **Senior (5+ yrs):** Internals, trade-offs, system design, performance

---

## 🔹 Basics

---

### 1. What is Angular? 🟢 | 🏭🚀

"Angular is an **open-source, TypeScript-based frontend framework** developed and maintained by **Google**. It provides a complete platform for building web applications — not just a library.

It includes everything out of the box: a component system, routing, HTTP client, form handling, animation, and testing utilities.

For me, Angular's biggest strength is its **opinionated structure**. When I join a new Angular project, the file layout, module system, and coding conventions are predictable. This makes large teams far more productive than with more flexible libraries like React."

#### In Depth
Angular is based on the **MVVM (Model-View-ViewModel)** pattern. Its compiler transforms TypeScript and HTML templates into optimized JavaScript at build time using **AOT (Ahead-of-Time)** compilation. Angular's **Ivy compiler** (introduced in Angular 9) produces smaller bundles and enables better tree-shaking compared to the older View Engine.

---

### 2. How is Angular different from AngularJS? 🟢 | 🏭🚀

"Angular (v2+) is a **complete rewrite** of AngularJS (v1.x). They share a name but are architecturally very different.

AngularJS is a JavaScript framework using `$scope` and two-way binding with `ng-model`. Angular uses **TypeScript**, a **component-based architecture**, a **proper CLI**, and **reactive programming** with RxJS.

Performance-wise, Angular uses **Ahead-of-Time compilation** and **zone-based change detection**, while AngularJS relied on a slow **dirty-checking digest cycle** that iterated over all watchers repeatedly."

#### In Depth
AngularJS uses **digest cycles**: on every event, it runs all registered `$watch` functions until the model stabilizes. This is O(n) per cycle and causes performance degradations at scale. Angular replaced this with **zone.js** which patches browser APIs and triggers targeted change detection only on the component subtree that changed.

---

### 3. What are the key features of Angular? 🟢 | 🏭

"Angular's headline features are:

- **Component-based architecture** — UI is split into reusable, self-contained components
- **TypeScript** — Static typing catches errors at compile-time
- **Angular CLI** — Powerful tooling for scaffolding, building, and testing
- **Dependency Injection** — Built-in DI system for decoupled, testable code
- **RxJS & Reactive Programming** — Async data streams with operators
- **Two-way data binding** — Sync between model and view via `ngModel`
- **Lazy Loading** — Load modules only when needed for performance
- **Ivy Compiler** — Produces smaller, faster bundles

I think the combination of TypeScript and DI makes Angular the best choice for **large enterprise apps** where maintainability matters over the long run."

#### In Depth
Angular's DI system uses a **hierarchical injector tree** — each component can have its own injector, and Angular walks up the tree to find a provider. This gives fine-grained control over service lifetimes. The Ivy compiler also enables **local compilation** of libraries, meaning each library can be compiled independently, improving compilation speed in CI/CD pipelines.

---

### 4. What is TypeScript and why is it used in Angular? 🟢 | 🏭

"TypeScript is a **statically typed superset of JavaScript** developed by Microsoft. It adds optional type annotations, interfaces, generics, and decorators to JavaScript, and compiles down to plain JS.

Angular uses TypeScript because:

1. **Type safety** catches bugs at compile time, not runtime
2. **Decorators** like `@Component`, `@Injectable` are TypeScript features that Angular's DI system relies on
3. **Better IDE support** — autocomplete, refactoring, and navigation work perfectly
4. **Metadata reflection** — TypeScript's `emitDecoratorMetadata` allows Angular to inspect constructor parameter types at runtime for DI

In large teams, TypeScript is the difference between a maintainable codebase and a messy one."

#### In Depth
Angular uses `reflect-metadata` polyfill to read TypeScript decorator metadata at runtime. When `@Injectable()` decorates a class, TypeScript emits the constructor parameter types as metadata. Angular's DI reads this metadata to automatically resolve and inject dependencies without explicit configuration. This is why `experimentalDecorators` and `emitDecoratorMetadata` must be enabled in `tsconfig.json`.

---

### 5. What are components in Angular? 🟢 | 🏭🚀

"A **component** is the fundamental building block of an Angular application. It controls a section of the UI via an HTML template and encapsulates the associated logic in a TypeScript class.

Every component has three parts:
- **Template** (HTML) — defines the view
- **Class** (TypeScript) — defines the behavior and data
- **Metadata** (`@Component` decorator) — tells Angular how to process the class

I treat each component as a **self-contained unit**. It has its own styles (`encapsulation`), its own template, and it communicates with parents via `@Input()` and with children via `@Output()` event emitters."

#### In Depth
Under Ivy, each component compiles to an `ɵcmp` definition which includes the view factory, change detection instructions, and host bindings. Angular represents each view as a **Logical View Tree** (LView), which is a flat array structure rather than a nested object tree. This improves cache locality and reduces GC pressure compared to View Engine's ComponentRef tree.

---

### 6. What is a module in Angular? 🟢 | 🏭

"A **module** (`@NgModule`) is a container that groups related components, directives, pipes, and services together.

Every Angular app has at least one module — the **root module** (`AppModule`). Modules help organize code into **feature areas** and control the visibility of components (only exported components can be used by other modules).

With Angular 14+, **Standalone Components** allow you to skip `NgModule` entirely, importing dependencies directly in the component decorator."

#### In Depth
`NgModule` serves as a **compilation context** — the Angular compiler uses the module's `declarations`, `imports`, and `exports` to determine what templates can reference. With Ivy, modules became thinner wrappers; the actual compilation happens at the component level. This is why Standalone Components work without modules — they compile with their own local context declared via `imports: []` in `@Component`.

---

### 7. Explain Angular CLI. 🟢 | 🏭

"The **Angular CLI** (Command Line Interface) is a powerful tool for scaffolding, building, testing, and deploying Angular apps.

Key commands I use daily:
- `ng new` — bootstraps a new project with all config files
- `ng generate component` — creates a component with all boilerplate files
- `ng serve` — starts a dev server with live reload
- `ng build --configuration=production` — builds with optimization, minification, AOT
- `ng test` — runs unit tests with Karma/Jest
- `ng lint` — runs ESLint

The CLI abstracts Webpack, TypeScript compilation, and file generation. Without it, we'd spend hours configuring the build chain."

#### In Depth
Angular CLI uses **Architect** as its build orchestration layer. Each target (`build`, `test`, `serve`) is implemented as an `@angular-devkit/build-angular` builder. You can write **custom builders** to modify or extend the build pipeline — for example, to integrate a custom Webpack plugin or post-process bundles. With Angular 16+, esbuild-based builders (`@angular-devkit/build-angular:browser-esbuild`) offer dramatically faster build times.

---

### 8. What is a template in Angular? 🟢 | 🏭

"A **template** is the HTML view associated with an Angular component. It is not plain HTML — it is Angular's **template syntax** which extends HTML with additional features.

Templates support:
- **Interpolation**: `{{ title }}`
- **Property binding**: `[src]="imageUrl"`
- **Event binding**: `(click)="handleClick()"`
- **Two-way binding**: `[(ngModel)]="username"`
- **Structural directives**: `*ngIf`, `*ngFor`

The template is compiled by Angular's compiler into optimized rendering instructions — it is never interpreted at runtime in production builds."

#### In Depth
Angular's template compiler uses the **Template Reference AST (abstract syntax tree)** during compilation. It parses templates, resolves references to components/directives/pipes, and generates incremental DOM rendering instructions. AOT compilation catches template errors (like referencing undefined variables) at build time, which JIT compilation would only catch at runtime.

---

### 9. What are directives in Angular? 🟢 | 🏭🚀

"**Directives** are classes that add behavior to elements in Angular templates.

There are **three types**:

1. **Components** — Directives with their own template (`@Component`)
2. **Structural Directives** — Alter the DOM structure (`*ngIf`, `*ngFor`, `*ngSwitch`)
3. **Attribute Directives** — Change appearance/behavior of an element without altering DOM structure (`NgClass`, `NgStyle`, custom directives like `[highlight]`)

I've used attribute directives to implement shared behaviors like auto-focus, lazy-load image placeholders, and permission-based visibility — keeping the business logic out of templates."

#### In Depth
Structural directives work through a **de-sugaring** mechanism. `*ngIf="show"` desugars to `<ng-template [ngIf]="show">`. Angular passes a `TemplateRef` and `ViewContainerRef` to the directive. The directive then controls when to create (`viewContainer.createEmbeddedView(template)`) or destroy the view. This gives structural directives complete control over the view lifecycle.

---

### 10. What is data binding? 🟢 | 🏭

"**Data binding** is the mechanism that synchronizes data between the component class (the model) and the view (the template).

Angular supports four types:

1. **One-way (Class → View)**: Interpolation `{{ name }}` and Property binding `[value]="name"`
2. **One-way (View → Class)**: Event binding `(click)="handler()"`
3. **Two-way**: `[(ngModel)]` — combines property and event binding

Data binding eliminates the need for manual DOM manipulation. I don't need `document.getElementById` — Angular keeps the view in sync with the model automatically."

#### In Depth
Data binding is implemented via Angular's **change detection** mechanism. During each change detection run, Angular evaluates all binding expressions in the component's template. For `OnPush` components, Angular skips evaluation unless an `@Input()` reference changes, an event fires, or `markForCheck()` is called. This makes understanding data binding essential for performance tuning.

---

## 🔹 Data Binding (Deep Dive)

---

### 11. What is one-way data binding? 🟢 | 🏭

"**One-way data binding** means data flows in a single direction — either from the component class to the view, or from the view to the component class.

**Class → View**: Interpolation `{{ title }}` and Property binding `[disabled]="isLoading"`
**View → Class**: Event binding `(click)="submit()"`

One-way binding is simpler and more predictable than two-way binding. I prefer it for most cases because the data flow is explicit and easier to debug."

#### In Depth
Property binding compiles to a call like `ɵɵproperty('disabled', ctx.isLoading)`. Angular only updates the DOM if the value has changed (dirty checking). This **equality check** is a strict reference check (`===`) for objects, which is why mutating an array doesn't trigger an update — you must return a new reference.

---

### 12. What is two-way data binding? 🟢 | 🏭

"**Two-way data binding** synchronizes data in both directions — changes in the component class update the view, and changes in the view (like user input) update the component class.

It is implemented using `[(ngModel)]` which is **syntactic sugar** for combining a property binding and an event binding:

```html
<input [(ngModel)]="username" />
<!-- Is equivalent to: -->
<input [ngModel]="username" (ngModelChange)="username = $event" />
```

I use it primarily for forms where quick prototyping matters. For complex forms, I prefer **Reactive Forms** for better control and testability."

#### In Depth
The **banana-in-a-box** syntax `[()]` signals Angular that the directive supports two-way binding. Angular looks for a corresponding `<propertyName>Change` output event. Any directive can support two-way binding — not just `ngModel`. For example, `[(value)]="x"` works if the component has both `@Input() value` and `@Output() valueChange = new EventEmitter()`.

---

### 13. How do you implement two-way data binding in Angular? 🟢 | 🏭

"I use the `[(ngModel)]` directive from `FormsModule`:

```typescript
// app.module.ts
import { FormsModule } from '@angular/forms';
@NgModule({ imports: [FormsModule] })

// component.html
<input [(ngModel)]="searchQuery" placeholder="Search..." />
<p>You typed: {{ searchQuery }}</p>
```

For reactive forms, two-way binding works differently — `FormControl` acts as the model and `[formControl]` binds the input."

#### In Depth
`ngModel` implements the `ControlValueAccessor` interface, which is Angular's bridge between form controls and native DOM elements. If I create a custom input component, I implement `ControlValueAccessor` to make it work with both template-driven and reactive forms using `ngModel` and `formControl`. This decouples the form API from the DOM implementation.

---

### 14. What is property binding? 🟢 | 🏭

"**Property binding** binds a component property to a DOM element property, keeping them in sync.

```html
<img [src]="imageUrl" [alt]="imageAlt" />
<button [disabled]="isSubmitting">Submit</button>
```

The difference from attribute binding (`[attr.aria-label]`) is that property binding sets the **DOM property**, not the HTML attribute. For most cases this is what you want. Use `[attr.X]` for ARIA attributes or SVG that don't have DOM property equivalents."

#### In Depth
DOM properties and HTML attributes are **not the same thing**. `<input value="hello">` sets the `defaultValue` attribute, while `inputEl.value` is the live DOM property. Angular's `[value]="x"` binds to the live DOM property. This is why `[value]` reflects real-time changes but `[attr.value]` only sets the initial attribute.

---

### 15. What is event binding? 🟢 | 🏭

"**Event binding** listens to DOM events and calls a handler method.

```html
<button (click)="saveRecord()">Save</button>
<input (keyup)="onSearch($event)" (blur)="validate()" />
```

The `$event` object captures the native DOM event. I use `$event.target.value` to read input values in template-driven patterns.

For custom components, I also use event binding with `@Output()` event emitters."

#### In Depth
Angular uses **event delegation** — it doesn't add listeners to every element. Event binding at the component level uses `zone.js`, which patches browser event APIs. When an event fires, zone.js triggers Angular's change detection. Using `runOutsideAngular()` to register some events (like scroll or `mousemove`) prevents unnecessary change detection cycles for performance-critical event handlers.

---

### 16. What is interpolation? 🟢 | 🏭

"**Interpolation** uses `{{ expression }}` to embed dynamic values into the template HTML.

```html
<h1>Welcome, {{ user.name }}!</h1>
<p>Total: {{ items.length }} items</p>
<span>{{ 2 + 2 }}</span>
```

The expression inside `{{ }}` is evaluated by Angular and the result (converted to string) is inserted into the DOM.

Interpolation is **one-way** — it reads from the component class but cannot write back. For security, Angular automatically escapes HTML to prevent XSS."

#### In Depth
Interpolation is actually **syntactic sugar for property binding** on the `textContent` DOM property. Angular sanitizes interpolated values through its **DomSanitizer** automatically. You cannot inject raw HTML via `{{ html }}` — it will be escaped. To display HTML, you must use `[innerHTML]="trustedHtml"` and explicitly mark it as safe using `DomSanitizer.bypassSecurityTrustHtml()` — which carries XSS risk.

---

### 17. Difference between interpolation and property binding? 🟡 | 🏭🚀

"They both flow data from the class to the view, but with one key difference:

- **Interpolation**: Always produces a **string** — `{{ value }}`. Good for text content.
- **Property binding**: Sets the actual **DOM property** with its native type — `[value]="count"`. Necessary for boolean (`disabled`, `checked`), numbers, objects.

```html
<!-- Wrong for boolean — sets the string "false" which is truthy! -->
<button disabled="{{ isDisabled }}">Click</button>

<!-- Correct — sets the boolean false -->
<button [disabled]="isDisabled">Click</button>
```

I use interpolation for text, property binding for everything else — especially when the type matters."

#### In Depth
Interpolation compiles to: `ɵɵtextInterpolate1('Welcome, ', ctx.name, '!')`. Property binding compiles to: `ɵɵproperty('disabled', ctx.isDisabled)`. These are both efficient DOM update instructions, but the property binding variant preserves the native type whereas interpolation always converts to string via the `toString()` method.

---

## 🔹 Components & Lifecycle

---

### 18. How do you create a component in Angular? 🟢 | 🏭

"The easiest way is via Angular CLI:

```bash
ng generate component user-profile
# or shorthand
ng g c user-profile
```

This creates four files: `.ts`, `.html`, `.css`, and `.spec.ts`. The component is also auto-declared in the nearest `NgModule`.

Manually, I define a class with the `@Component` decorator:

```typescript
@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss'],
})
export class UserProfileComponent {
  @Input() userId: string;
  @Output() profileUpdated = new EventEmitter<User>();
}
```"

#### In Depth
Angular CLI uses **schematics** under the hood — the `ng generate` command runs the `@schematics/angular:component` schematic. You can create custom schematics to enforce team conventions — for example, always generating components with SCSS, `OnPush` change detection, and a test harness. This reduces boilerplate setup and standardizes team patterns.

---

### 19. What is the role of the @Component decorator? 🟢 | 🏭

"The `@Component` decorator is a TypeScript decorator that tells Angular how to create and render the component. It provides the **metadata** Angular needs to compile and instantiate the component.

Key properties:
- `selector` — the CSS selector used in templates (`<app-user-profile>`)
- `template` / `templateUrl` — inline or external HTML
- `styles` / `styleUrls` — component-specific CSS
- `changeDetection` — `Default` or `OnPush`
- `encapsulation` — `Emulated`, `ShadowDom`, or `None`
- `providers` — component-scoped service providers

The decorator is processed at **compile time by the Angular compiler** (AOT), not at runtime."

#### In Depth
`@Component` is built on top of `@Directive`. Internally, it registers the class in Angular's component registry. During Ivy compilation, the decorator's metadata is transformed into a static `ɵcmp` property on the class, which contains the component definition object (template function, host bindings, dependency list, etc.). This static property is then used by Angular's runtime to create views.

---

### 20. What are lifecycle hooks in Angular? 🟢 | 🏭🚀

"**Lifecycle hooks** are methods that Angular calls at specific stages of a component's or directive's life — from creation to destruction.

The hooks in order of execution:

| Hook | When Called |
|---|---|
| `ngOnChanges` | Before `ngOnInit` and when `@Input()` values change |
| `ngOnInit` | Once, after first `ngOnChanges` |
| `ngDoCheck` | Every change detection run |
| `ngAfterContentInit` | After content (`ng-content`) is projected |
| `ngAfterContentChecked` | After projected content is checked |
| `ngAfterViewInit` | After the component's view and child views are initialized |
| `ngAfterViewChecked` | After the component's view is checked |
| `ngOnDestroy` | Just before Angular destroys the component |

I mainly use `ngOnInit` for data fetching, `ngOnDestroy` for cleanup (unsubscribing), and `ngOnChanges` to react to `@Input()` changes."

#### In Depth
`ngDoCheck` and `ngAfterContentChecked` / `ngAfterViewChecked` run on **every change detection cycle**, which can be very frequent. Putting expensive logic here is a performance anti-pattern. Always run performance-sensitive code in `ngOnInit` or triggered explicitly. `ngOnDestroy` is the correct place to call `.unsubscribe()` on RxJS subscriptions, clear intervals, or call `ChangeDetectorRef.detach()` to prevent memory leaks.

---

### 21. Explain ngOnInit(). 🟢 | 🏭

"**`ngOnInit()`** is called once by Angular after the component's first `ngOnChanges()` run — meaning `@Input()` properties have their initial values set.

This is the correct place to:
- **Fetch initial data** from a service / API
- **Subscribe to route parameters** (`ActivatedRoute`)
- **Initialize UI state** that depends on inputs

```typescript
ngOnInit(): void {
  this.userId = this.route.snapshot.paramMap.get('id');
  this.userService.getUser(this.userId).subscribe(user => {
    this.user = user;
  });
}
```

I **never** put HTTP calls in the constructor — constructors should be lightweight and only handle DI."

#### In Depth
The reason NOT to use the constructor for initialization is the **DI resolution order**. When the constructor runs, `@Input()` values have not been set yet — they are only available from `ngOnChanges()` onward, which runs before `ngOnInit()`. Also, constructors run during Angular's compilation phase in some edge cases (like in testing), making them an unreliable place for side effects.

---

### 22. What is ngOnChanges() used for? 🟡 | 🏭🚀

"**`ngOnChanges()`** is called whenever a **bound `@Input()` property changes**. It receives a `SimpleChanges` object that contains the previous and current values.

```typescript
ngOnChanges(changes: SimpleChanges): void {
  if (changes['userId']) {
    const { previousValue, currentValue, firstChange } = changes['userId'];
    if (!firstChange) {
      this.loadUser(currentValue); // Reload when userId changes
    }
  }
}
```

I use it to **react to input changes** — like fetching new data when a parent passes a different ID, or recalculating derived values.

Important: `ngOnChanges` only fires for **`@Input()` properties**, not for changes to objects' inner properties (deep changes)."

#### In Depth
`ngOnChanges` uses **reference comparison** (`===`) to detect changes in `@Input()` properties. If a parent mutates an object or array (rather than replacing the reference), `ngOnChanges` will NOT fire. This is why immutable state patterns (always returning new objects/arrays) pair well with Angular's `OnPush` strategy — they guarantee `ngOnChanges` triggers on every meaningful update.

---

### 23. What is ngAfterViewInit()? 🟡 | 🏭🚀

"**`ngAfterViewInit()`** is called once after Angular has fully initialized the component's view, including all child component views.

This is the correct lifecycle hook to access `@ViewChild()` and `@ViewChildren()` references, because child components are only rendered after the parent's view is initialized.

```typescript
@ViewChild('myCanvas') canvasRef: ElementRef;

ngAfterViewInit(): void {
  // ✅ Safe — canvas is in the DOM
  const ctx = this.canvasRef.nativeElement.getContext('2d');
  this.drawChart(ctx);
}
```

If I try to access `@ViewChild` in `ngOnInit`, it would be `undefined`."

#### In Depth
`ngAfterViewInit` fires after the **first render** of the component's view tree. Any operations that mutate the template (like setting a property that affects the template) inside `ngAfterViewInit` can trigger the **ExpressionChangedAfterItHasBeenCheckedError** in development mode, because Angular runs a second check immediately after. Use `setTimeout()` or wrap mutations in `ChangeDetectorRef.detectChanges()` to avoid this error.

---

### 24. Can we use lifecycle hooks in services? 🔴 | 🚀

"Standard lifecycle hooks like `ngOnInit` and `ngOnDestroy` are **designed for directives and components**, not services. However, there's a related pattern for services:

Since Angular 9, services can implement `OnDestroy` to get notified when they are destroyed. This works for **component-scoped services** (provided in a component's `providers` array) — they are destroyed with the component.

```typescript
@Injectable()
export class StreamService implements OnDestroy {
  private subscription = new Subscription();

  ngOnDestroy(): void {
    this.subscription.unsubscribe(); // Cleanup when service dies
  }
}
```

For root-scoped services (`providedIn: 'root'`), `ngOnDestroy` is called only when the application shutdowns, making it less useful for cleanup."

#### In Depth
Component-scoped services are destroyed when the component that provides them is destroyed. Angular calls `ngOnDestroy` on the service before the component's own `ngOnDestroy`. This allows services to clean up resources (unsubscribe, close WebSocket connections) before the component itself is torn down. This pattern is especially useful for component-scoped **store services** or **polling services**.

---
