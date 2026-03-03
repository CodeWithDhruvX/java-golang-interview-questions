# 🗣️ Theory — Pipes, Modules & Change Detection
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are Angular pipes? Name the common built-in ones."

> *"A pipe transforms a value in an Angular template without changing the underlying data. You apply it with the pipe character — {{ date | date:'dd/MM/yyyy' }}. Angular ships with several useful built-ins: date formats a Date object, currency formats a number as money with locale-aware symbols, uppercase and lowercase and titlecase transform strings, number and percent format numeric values with decimal places, slice works on arrays and strings, async subscribes to an Observable or Promise and returns its latest value, and json serializes an object for debugging. Pipes are the right place for view-layer transformations — formatting phone numbers, truncating text, converting status codes to labels. All built-in pipes except async are pure, meaning they only recalculate when the input reference changes."*

---

## Q: "How do you create a custom pipe?"

> *"You create a class decorated with @Pipe and give it a name. The class implements PipeTransform, which requires one method: transform. The first parameter of transform is the value being piped; subsequent parameters are optional arguments passed with colons in the template. So a truncate pipe might take the max length as a second parameter — {{ description | truncate:100 }}. You return the transformed value. You declare the pipe in the NgModule's declarations array, or for standalone, you import it directly. Pure pipes — the default — only run when Angular detects a reference change in the input value. Impure pipes run on every change detection cycle, which can be expensive — only use impure when you absolutely need to react to mutations inside an object or array."*

---

## Q: "What is the difference between a pure and an impure pipe?"

> *"A pure pipe runs only when Angular detects a pure change in the input — meaning a changed primitive value, or a different object or array reference. If you pass an array and then mutate it in-place, a pure pipe won't re-execute because the reference hasn't changed. An impure pipe — you opt in with pure:false — runs on every change detection cycle, regardless of whether the input reference changed. This means it detects mutations inside objects and arrays, but it comes at a significant performance cost because Angular calls it constantly. The async pipe is the one impure built-in pipe, and it needs to be impure because it has to react to Observables emitting new values over time. For custom list-filtering pipes, most developers avoid impure pipes and prefer filtering the array in the component class instead."*

---

## Q: "What is NgModule? Explain declarations, imports, exports, and providers."

> *"NgModule is a container that organizes related Angular code into a cohesive block. The declarations array lists every component, directive, and pipe that belongs to this module — these are private by default, only usable within the module. The imports array lists other modules whose exported components, directives, and pipes you want to use — like CommonModule for *ngIf and *ngFor, or ReactiveFormsModule for reactive forms. The exports array makes some of your declarations or imported modules available to other modules that import this one — this is how SharedModule works. The providers array registers services at the module's injector level — though most services are now registered with providedIn:'root' in their @Injectable decorator."*

---

## Q: "How does Angular change detection work? What triggers it?"

> *"Angular's change detection is how it keeps the DOM in sync with the component data. Zone.js is the engine — it patches all browser async APIs so Angular gets notified whenever an event, timer, or HTTP response completes. When notified, Angular runs change detection starting from the root component and traversing the tree top-to-bottom. For each component, it evaluates every template binding expression and compares the new value to the previous one. If anything changed, Angular updates the DOM for that binding. With the Default strategy, Angular checks every component in the tree on every async event. With OnPush, Angular skips a component unless one of its @Input() references changed, an async pipe emitted, or an event inside the component fired."*

---

## Q: "What is ChangeDetectionStrategy.OnPush and when should you use it?"

> *"OnPush is a change detection strategy that tells Angular: only check this component when you have a strong reason to believe its output changed. That reason is: an @Input() property received a new reference — not a mutation — an Observable subscribed via the async pipe emitted a new value, or an event was triggered inside this component. Everything else, OnPush components are skipped. The performance win is significant for large component trees — instead of evaluating hundreds of template bindings on every click anywhere in the app, Angular skips the OnPush components. The trade-off is immutability: since OnPush only reacts to reference changes, you must return new objects and arrays rather than mutating in place. I use OnPush on all 'presentational' or 'dumb' components — those that only receive data via @Input and don't have their own async side effects."*

---

## Q: "What is SharedModule and how do you structure it?"

> *"SharedModule is a module where you put components, directives, and pipes that are reused across multiple feature modules. The pattern is: declare everything in SharedModule, and also export everything — including re-exporting commonly needed Angular modules like CommonModule and ReactiveFormsModule. Feature modules then import just SharedModule and get access to all the shared utilities. The key rule is: never import SharedModule in AppModule. SharedModule shouldn't provide singleton services — those belong in CoreModule or with providedIn:'root'. SharedModule is also where you put presentational components like loading spinners, empty state displays, and confirmation dialogs that every feature needs but which have no business logic of their own."*
