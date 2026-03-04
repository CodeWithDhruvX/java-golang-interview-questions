# 🗣️ Theory — Directives, Services & Dependency Injection
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are the three types of directives in Angular?"

> *"Angular has three directive types. Components are directives with their own template — every @Component is technically a directive. Structural directives change the structure of the DOM by adding or removing elements — *ngIf, *ngFor, and *ngSwitch are the built-ins. Attribute directives change the appearance or behavior of an existing element without touching the DOM structure — NgClass, NgStyle, and any custom directives like a [highlight] or [tooltop] directive. The asterisk syntax on structural directives is syntactic sugar — *ngIf='show' desugars to a ng-template with [ngIf]='show', which lets the directive control whether to create or destroy the embedded view."*

---

## Q: "What is the difference between *ngIf and [hidden]?"

> *"Both control visibility, but they work completely differently. *ngIf removes the element from the DOM entirely when the condition is false — the component is destroyed, its subscriptions are cleaned up, and there's no memory overhead. [hidden] just adds a CSS display:none style — the element stays in the DOM, the component remains alive, all its resources are held. So *ngIf is better for initial load performance when an element is frequently not needed. [hidden] is better when you're toggling something frequently — like a dropdown — because destroying and recreating DOM nodes is expensive. A practical rule: if the element starts hidden and may never show, use *ngIf. If the user is toggling something many times per second, use [hidden]."*

---

## Q: "How do trackBy work with *ngFor and why is it important?"

> *"By default, when the array bound to *ngFor changes — even if only one item was added — Angular destroys and recreates all list DOM nodes, because it can't tell which items changed. trackBy gives Angular a unique identifier for each item. You provide a function that returns a stable ID, like the item's database ID. With trackBy, Angular can compare old and new lists item-by-item and only create, update, or destroy the nodes that actually changed. In a list of 500 products, without trackBy every API refresh destroys and recreates 500 DOM nodes. With trackBy, it only touches the nodes that actually changed. This is a significant performance win for large lists."*

---

## Q: "What is a Service in Angular and why do we use them?"

> *"A service is a plain TypeScript class decorated with @Injectable that Angular's DI system can instantiate and inject. Services exist to solve a separation-of-concerns problem — you don't want business logic, HTTP calls, and state management living inside your components. Components should focus on the view. Services extract the reusable logic. Common uses: encapsulating API calls in a ProductService, managing auth state in an AuthService, logging in a LogService, or sharing state between sibling components via a BehaviorSubject. The @Injectable decorator tells Angular's DI system that this class can be injected, and the providedin metadata controls the scope."*

---

## Q: "What does 'providedIn: root' mean? When would you scope a service differently?"

> *"'providedIn: root' means the service is registered in the root injector — there's one instance shared across the entire app. This is the right default for most services: AuthService, HttpClient wrappers, shared state services. You'd scope a service differently in two cases: providing it in a feature module means it's only available to components in that module, which is useful for module-specific services. Providing it in a component's providers array means a new instance is created for each instance of that component — useful when you want isolated state per component, like a validation service that needs fresh state for each form instance. The hierarchical injector tree means Angular walks up from child to parent to root looking for a provider."*

---

## Q: "What is Dependency Injection? Why does Angular use it?"

> *"Dependency injection is a design pattern where a class declares what it needs as constructor parameters, and something external — the injector — is responsible for creating and providing those dependencies. Angular's DI system is what makes this pattern automatic. When you write 'constructor(private productService: ProductService)', you're not creating a ProductService — Angular looks up the type token, finds the registered provider, and hands you the instance. The benefits are: loose coupling — the component doesn't know how ProductService is constructed; testability — you can easily swap UserService for a fake in tests; and singleton management — Angular manages lifetimes for you. Without DI, every component would create its own service instances and sharing state between components would require global variables."*

---

## Q: "How do you create a custom attribute directive?"

> *"You create a class decorated with @Directive and give it a selector in square brackets — like [appHighlight]. Angular uses this selector to match elements in templates. Inside the directive class, you inject ElementRef to get a reference to the host DOM element. You use @HostListener to listen to events on the host element — like mouseenter and mouseleave — and @HostBinding to bind to properties on the host element — like the backgroundColor style. @Input on the directive lets you pass configuration from the parent template. For example, [appHighlight]='yellow' passes the color. The directive pattern is ideal for reusable DOM behaviors — things like auto-resize textareas, click-outside detection, permission-based disabling, or lazy-loading images."*

---

## Q: "What is @HostListener vs @HostBinding?"

> *"@HostListener and @HostBinding are two ways a directive or component interacts with its host element. @HostListener attaches an event listener to the host — @HostListener('click') onClick() registers the onClick method to fire when the host element is clicked. @HostBinding binds a class property to a property of the host element — @HostBinding('class.active') isActive = false means when isActive is true, Angular adds the 'active' CSS class to the host element. Together they let you build fully declarative directives without touching ElementRef imperatively. Using @HostListener for events and @HostBinding for state is cleaner and more testable than calling nativeElement methods directly."*
