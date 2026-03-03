# 🗣️ Theory — Signals, Standalone Components & Modern Angular
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are Angular Signals? Why are they a significant change to the framework?"

> *"Signals are a new reactivity primitive Angular introduced in version 16. A signal is a wrapper around a value that notifies interested consumers when the value changes. You call signal() with an initial value, call it as a function to read the current value, and call .set() or .update() to change it. The significance is about change detection. With Zone.js, Angular has no fine-grained knowledge of what changed — it just knows 'some async thing happened, check everything'. With Signals, Angular knows exactly which component depends on which signal, because it tracks which signals are read during template evaluation. This enables fine-grained change detection — when a signal changes, only the components that read that specific signal need to re-render, not the whole tree. Long term, Signals are the path to a zoneless Angular that doesn't need Zone.js at all."*

---

## Q: "What are computed() and effect() in Angular Signals?"

> *"computed() creates a derived signal — it takes a function that reads other signals and returns a derived value. Computed signals are lazy and memoized: they don't run the function until they're first read, and they cache the result until one of the signals they depend on changes. It's similar to how selectors work in NgRx, but it's baked into the reactivity system. effect() creates a side effect that runs whenever signals it depends on change. It's Angular's equivalent of useEffect in React, but reactive — you don't specify a dependency array, Angular tracks the signals automatically. An effect runs once immediately on creation, and then again any time a read signal emits a new value. Use effects for logging, analytics, or syncing signal state to localStorage — for DOM updates, prefer template bindings which are automatically reactive to signals."*

---

## Q: "What are Standalone Components and how do they change Angular's architecture?"

> *"Standalone Components, stable since Angular 15, let you build Angular components without NgModules. Instead of declaring a component in a module's declarations array and importing needed modules into that module, you mark the component as standalone: true in the @Component decorator and import dependencies directly — component imports like CommonModule, pipes like DatePipe, or other standalone components, go directly in the component's imports array. This makes components self-contained and portable. For app bootstrapping, you call bootstrapApplication() with the root component instead of bootstrapModule(). The tree-shaking benefits are significant — the compiler can see exactly which component uses which directives and pipes, and dead code is eliminated more precisely. The Angular team's direction is strongly toward standalone components; new projects created with the CLI default to standalone."*

---

## Q: "Explain Angular 17's new control flow syntax: @if, @for, @switch."

> *"Angular 17 introduced built-in control flow blocks as an alternative to *ngIf, *ngFor, and *ngSwitch directives. @if is a block with optional @else if and @else branches — it reads almost like TypeScript if-else. @for replaces *ngFor and has a crucial difference: track is required, not optional. This enforces the performance best practice — you must provide a tracking expression for Angular to identify items. @for also has an @empty block for when the collection is empty, which eliminates the common *ngIf='items.length === 0' pattern. @switch replaces *ngSwitch with cleaner @case and @default syntax. The fundamental advantage is that these are built into the Angular template compiler — no directive import needed, not even CommonModule. They also enable better build-time analysis and are more type-aware than directives."*

---

## Q: "What are signal-based inputs with input() and model()? How do they differ from @Input()?"

> *"Angular 17.1 introduced input() as a function-based alternative to the @Input() decorator. You write name = input.required<string>() and the result is a signal — you read it with this.name() which gives you the current input value. The advantages: it composes naturally with computed() — you can write displayName = computed(() => this.name().toUpperCase()), and Angular knows to re-evaluate this whenever the input changes without any extra wiring. model() is the bidirectional equivalent — it creates a signal that the component can set internally, and the parent can also bind to with two-way binding syntax. Think of it as Signal-based @Input() plus @Output() combined. The key difference from @Input() is that input() and model() give you a signal, which integrates with Angular's reactive graph. With classic @Input(), you need ngOnChanges to react to input changes; with input(), you react with computed() or effect() declaratively."*

---

## Q: "What is Zoneless Angular and when will it become the default?"

> *"Zoneless Angular means running without Zone.js — instead of Zone.js triggering change detection after every async event, change detection is triggered selectively by Signals. It became experimentally available in Angular 18 via the provideExperimentalZonelessChangeDetection() provider. Components in a zoneless app must use Signals for reactive state — property updates that don't go through Signals won't trigger re-rendering. The benefits are: smaller bundle size since Zone.js doesn't ship, better compatibility with native browser APIs that Zone.js interferes with, and fundamentally better performance because change detection only runs for components that have signal changes. The Angular team's trajectory is clear — standalone components plus Signals plus zoneless is the future. It won't likely be the default until Angular 19 or 20, but starting new components with Signals and OnPush today positions you for a smooth migration."*
