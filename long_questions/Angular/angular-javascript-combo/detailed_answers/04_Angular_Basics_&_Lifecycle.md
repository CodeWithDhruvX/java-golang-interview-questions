# 🏗 Angular Basics & Lifecycle Hooks (Detailed Answers)

## 1. What is Angular?
A TypeScript-based open-source web application framework led by the Angular Team at Google. It's used for building single-page client applications using HTML and TypeScript.
**Features**: Two-way data binding, Dependency Injection (DI), Routing, Forms, Observables (RxJS).

## 2. What is SPA (Single Page Application)?
An app that loads a single HTML page (index.html) and dynamically updates that page as the user interacts with the app.
**Benefits**: No page reloads (faster), smooth UX, code reusability (components).

## 3. What is TypeScript?
A strongly typed superset of JavaScript that compiles to plain JavaScript.
**Why Angular uses it?**
*   Static typing (catches errors early).
*   Interfaces/Classes (OOP).
*   Decorators (@Component, @Injectable).
*   Better IDE support (IntelliSense).

## 4. What is Angular CLI?
A command-line interface tool to initialize, develop, scaffold, and maintain Angular applications.
**Common Commands:**
*   `ng new my-app` (Create project)
*   `ng serve -o` (Run dev server)
*   `ng generate component/service/pipe` (Create files)
*   `ng build --prod` (Build for production)

## 5. Explain Angular architecture.
1.  **Modules (NgModule)**: Confines components, directives, pipes, and services. Root module is `AppModule`.
2.  **Components**: Controls a portion of the screen (View).
3.  **Templates**: HTML view for the component.
4.  **Metadata**: Tells Angular how to process a class.
5.  **Data Binding**: Connects component and template.
6.  **Directives**: Transform the DOM.
7.  **Services**: Business logic.
8.  **Dependency Injection**: Provide services to components.

## 6. What is a component?
The fundamental building block of Angular apps.
Each component consists of:
1.  **HTML**: Template (`templateUrl`).
2.  **CSS**: Styles (`styleUrls`).
3.  **TS**: Class (properties + methods).
4.  **Metadata**: `@Component` decorator.

## 7. What is a module (NgModule)?
A container for a cohesive block of code dedicated to an application domain, a workflow, or a closely related set of capabilities.
**Important properties of @NgModule**:
*   `declarations`: Components, Directives, Pipes belonging to this module.
*   `imports`: Other modules needed (e.g., `BrowserModule`, `FormsModule`).
*   `providers`: Services available to the app.
*   `bootstrap`: Root component to load.

## 8. What is metadata?
Metadata tells Angular how to process a class.
It is added using **Decorators**.

## 9. What is a decorator?
A function that modifies JavaScript classes.
*   **Class Decorators**: `@Component`, `@NgModule`, `@Injectable`, `@Directive`, `@Pipe`.
*   **Property Decorators**: `@Input`, `@Output`, `@ContentChild`, `@ViewChild`, `@HostBinding`.
*   **Method Decorators**: `@HostListener`.
*   **Parameter Decorators**: `@Inject`, `@Optional`.

## 10. What is bootstrap array?
The `bootstrap` property in the root `NgModule` tells Angular which component to load first (usually `AppComponent`). It gets inserted into the `index.html` host file.

---

## 11. Explain Angular Lifecycle Hooks (Order of Execution)

| Hook | Purpose | Timing |
| :--- | :--- | :--- |
| **ngOnChanges** | Respond to input (`@Input`) changes. | Before ngOnInit, and whenever inputs change. |
| **ngOnInit** | Initialize the component/directive. | Called once, after the first ngOnChanges. |
| **ngDoCheck** | Detect and act upon changes that Angular can't/won't detect. | Immediately after ngOnInit/ngOnChanges. |
| **ngAfterContentInit** | Respond after Angular projects external content into component's view. | Called once after the first ngDoCheck. |
| **ngAfterContentChecked** | Respond after Angular checks the content projected into the component. | Called after ngAfterContentInit and every subsequent ngDoCheck. |
| **ngAfterViewInit** | Respond after Angular initializes the component's views and child views. | Called once after the first ngAfterContentChecked. |
| **ngAfterViewChecked** | Respond after Angular checks the component's views and child views. | Called after ngAfterViewInit and every subsequent ngAfterContentChecked. |
| **ngOnDestroy** | Cleanup just before Angular destroys the component. | Unsubscribe Observables, detach event handlers. |

## 12. Difference between ngOnInit and constructor?
*   **Constructor**: Executed by JavaScript engine. Used strictly for Dependency Injection. Inputs are **not available** yet.
*   **ngOnInit**: Executed by Angular. Inputs (`@Input`) are now **available**. Best place for API calls / logic initialization.

## 13. When does ngOnChanges trigger?
It triggers **before `ngOnInit`** and whenever a **data-bound input property** changes.
It receives a `SimpleChanges` object which holds current and previous values.

```typescript
ngOnChanges(changes: SimpleChanges) {
  if (changes['userId']) {
    this.loadUserData(changes['userId'].currentValue);
  }
}
```

## 14. What is ngOnDestroy used for? (Crucial for Memory Leaks)
Used for cleanup logic.
1.  **Unsubscribe** from observables.
2.  **Clear Timers** (`setTimeout`, `setInterval`).
3.  **Detach event listeners** (from DOM).

## 15. Real use case of ngAfterViewInit?
Used when you need access to the DOM elements or child components (`@ViewChild`) that are part of the component's view. These are **undefined** in `ngOnInit`.

```typescript
@ViewChild('myInput') inputElement: ElementRef;

ngAfterViewInit() {
  this.inputElement.nativeElement.focus();
}
```
