# 🗣️ Theory — Angular Basics & Components
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is Angular? How is it different from AngularJS?"

> *"Angular — that's version 2 and above — is a complete, TypeScript-based frontend framework built and maintained by Google. The key word is complete — it's not just a library. It ships with everything you need: a component system, a router, an HTTP client, form handling, and dependency injection. You don't need to assemble your own tool stack the way you would with React. AngularJS, version 1, was a completely different animal — it used JavaScript, a scope-based MVC pattern, and a digest cycle for change detection that got slow at scale. Angular rewrote everything from scratch. It uses TypeScript, a true component architecture, a real CLI, and Zone.js for reactive change detection. The performance and developer experience are incomparable."*

---

## Q: "What is a Component in Angular? What are its parts?"

> *"A component is the fundamental building block of every Angular application. It represents one self-contained part of the UI — a button, a card, a whole page. Every component has three parts. First, the TypeScript class — this is where your logic lives, your properties, methods, and lifecycle hooks. Second, the HTML template — this is the view, and it uses Angular's special template syntax for data binding and directives. Third, the metadata in the @Component decorator — this tells Angular what selector to use, where the HTML file is, and what styles to apply. When Angular sees your selector in a parent template, it instantiates your class, evaluates the template, and renders the DOM."*

---

## Q: "What are Angular lifecycle hooks? Which ones do you use most?"

> *"Lifecycle hooks are methods Angular calls at specific moments in a component's life. The key ones are: ngOnChanges — called before ngOnInit and every time an @Input() property gets a new value; ngOnInit — called once after the first ngOnChanges, this is where I put data fetching and setup logic; ngAfterViewInit — called after the component view and all child views are rendered, this is where @ViewChild references become available; and ngOnDestroy — called just before Angular destroys the component, this is where I unsubscribe from Observables and clean up. The ones I use daily are ngOnInit for initialization and ngOnDestroy for cleanup. ngOnChanges is useful when you need to react to input changes from a parent."*

---

## Q: "Explain the four types of data binding in Angular."

> *"Angular has four types. Interpolation — double curly braces around an expression — outputs text content from the class into the template. Property binding — square brackets around a DOM property — binds a component property to a DOM element property, so [disabled]='isLoading' sets the disabled property to the value of isLoading. Event binding — parentheses around an event name — listens to DOM events and calls a class method, like (click)='submit()'. And two-way binding — the banana-in-a-box [(ngModel)] — combines a property binding and an event binding so the model and view stay in sync automatically. In practice, interpolation is for text output, property binding is for setting element states, event binding is for user interactions, and two-way binding is for form inputs."*

---

## Q: "What is the difference between @Input() and @Output()?"

> *"@Input and @Output are the two mechanisms for communication between a parent and child component. @Input decorates a property on the child — the parent can pass a value down to that property from its own template using property binding. So if a child has @Input() product, the parent writes [product]='selectedProduct' to pass data in. @Output decorates an EventEmitter on the child — when something happens inside the child, it calls emit() on the EventEmitter and the parent listens for that event with (event)='handler($event)'. So data flows down via @Input and events bubble up via @Output. This one-directional flow makes component communication predictable and easy to debug."*

---

## Q: "What is the difference between the constructor and ngOnInit?"

> *"The constructor runs when the class is instantiated by Angular's dependency injection system. At that point, @Input() properties haven't been set yet — they're only available from ngOnChanges onward. So the constructor should only be used for one thing: injecting services via the parameter list. ngOnInit runs once after the first ngOnChanges, meaning @Input() values are available. This is where you put your initialization logic — HTTP calls, subscribing to route parameters, setting up state. This separation is important for testing too: constructors that do HTTP calls make components hard to test. The rule is simple — constructor for DI only, ngOnInit for everything else."*

---

## Q: "What is @ViewChild and when do you use it?"

> *"@ViewChild gives a component access to a child component, directive, or DOM element defined in its template. You declare it as a property decorated with @ViewChild and reference a template reference variable or component type. The key constraint is that the element only exists in the DOM after the view is initialized, so you must access it in ngAfterViewInit — if you try to read it in ngOnInit, it's undefined. Common use cases: programmatically focusing an input element, calling a method on a child component, or interacting with a third-party DOM library that needs a native element reference. It's a direct imperative escape hatch — prefer data binding where possible and use @ViewChild when you genuinely need to reach into the child."*

---

## Q: "What does AOT compilation mean in Angular?"

> *"AOT stands for Ahead-of-Time compilation. It means Angular compiles your TypeScript and HTML templates into JavaScript at build time, on your developer machine, before the browser ever sees the code. The alternative is JIT — Just-in-Time — where the Angular compiler runs in the browser at startup. AOT has significant advantages. The bundle is smaller because the Angular compiler itself doesn't ship to the browser. The app loads faster because there's no compilation step in the browser. And template errors — typos, references to undefined variables — are caught at build time rather than at runtime. In Angular, AOT is the default for production builds and it's powered by the Ivy compiler which was introduced in Angular 9."*
