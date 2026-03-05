# Frontend Framework Basics (Angular & React)

Service-based companies usually look for practical experience in building UI components, managing state, and routing in popular frameworks like React (for MERN) or Angular (for MEAN).

## Shared Frontend Concepts

### 1. What is a Virtual DOM (React) vs. Real DOM (Angular)?
*   **Virtual DOM (React)**: A lightweight memory representation of the Real DOM. React updates the Virtual DOM, calculates the absolute minimum required changes using a 'diffing' algorithm, and updates the Real DOM in batches. This minimizes expensive DOM operations.
*   **Real DOM (Angular)**: Angular uses the Real DOM directly but employs "Change Detection" (via Zone.js historically, now moving towards signals for fine-grained reactivity) to figure out which parts of the component hierarchy need to instantly reflect data changes.

### 2. What are Single Page Applications (SPAs) and what are their benefits?
An SPA is a web application or website that interacts with the user by dynamically rewriting the current web page with new data from the web server, instead of the default method of a web browser loading entire new pages.
**Benefits**:
*   Faster transitions post-initial load (only JSON data is fetched, not HTML).
*   Smoother, more app-like user experience.
*   Easier to build responsive frontend and share backend APIs with mobile apps.

## Angular (MEAN Stack) Specifics

### 3. What is the difference between a Component and a Directive in Angular?
*   **Component**: A directive with a template. It controls a patch of the screen (the view). It has a defined HTML template and CSS styles.
*   **Directive**: A class that adds or modifies the behavior of DOM elements.
    *   *Structural Directives* (`*ngIf`, `*ngFor`): Alter DOM layout by adding/removing elements.
    *   *Attribute Directives* (`ngClass`, `ngStyle`): Change the appearance or behavior of an existing element.

### 4. Explain the Angular Component Lifecycle Hooks.
Angular calls standard lifecycle hooks in a sequence:
1.  `ngOnChanges`: When input bounds (`@Input`) change.
2.  `ngOnInit`: After first `ngOnChanges`. Ideal for component initialization and API calls.
3.  `ngDoCheck`: Custom change detection (use cautiously).
4.  `ngAfterContentInit` / `ngAfterContentChecked`: After projected content is initialized/checked (`<ng-content>`).
5.  `ngAfterViewInit` / `ngAfterViewChecked`: After component's views and child views are initialized/checked. Good for accessing DOM elements.
6.  `ngOnDestroy`: Just before Angular destroys the component. Ideal for unsubscribing from Observables to prevent memory leaks.

### 5. What are Services and Dependency Injection (DI) in Angular?
*   **Services**: Classes used to organize and share data or logic across multiple components (e.g., retrieving data from an API).
*   **Dependency Injection**: A design pattern where a class requests dependencies from external sources rather than creating them. Angular's DI framework instantiates services and injects them into components via their constructors, promoting modularity and testability.

### 6. What is RxJS and how does Angular use it?
RxJS (Reactive Extensions for JavaScript) is a library for reactive programming using Observables. Angular uses it extensively for handling asynchronous operations:
*   `HttpClient` returns Observables for HTTP requests.
*   Routing parameters and form value changes are represented as Observables.
*   *Key difference from Promises*: Observables can handle streams of multiple events over time, can be cancelled, and provide operators (like `map`, `filter`, `switchMap`) to manipulate streams.

## React (MERN Stack) Specifics

### 7. What is JSX?
JSX (JavaScript XML) is a syntax extension for React. It allows writing HTML structures within JavaScript code. React transpiles JSX into plain JavaScript (`React.createElement()`) calls under the hood.

### 8. Explain State and Props in React.
*   **Props (Properties)**: Read-only data passed downwards from parent to child components. Used for component configuration and communication.
*   **State**: Local, mutable data maintained within a component. When state changes, React re-renders the component to reflect the new data.

### 9. What are React Hooks? Name a few common ones.
Hooks are functions that let you "hook into" React state and lifecycle features from functional components (without writing classes).
*   `useState`: To add local state to functional components.
*   `useEffect`: To perform side effects (data fetching, subscriptions, DOM manipulation). Represents `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount` combined.
*   `useContext`: To access React Context directly (avoiding prop drilling).
*   `useRef`: To access a DOM element directly or hold a mutable value without causing a re-render.

### 10. Explain Prop Drilling and how to avoid it in React.
**Prop Drilling** is the process of passing props from a top-level component down through multiple layers of intermediate components just to reach a deeply nested child that actually needs the data.
**How to avoid it**:
*   React Context API.
*   State management libraries (Redux, Zustand, Recoil).
*   Component composition (passing child components as elements via props).
