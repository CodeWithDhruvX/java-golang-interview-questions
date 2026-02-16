# 🏗 Angular Architecture & Performance (Detailed Answers)

## 46. What is Angular Routing?
Enables navigation from one view to the next as users perform application tasks.
**Key Concepts**: `Routes`, `RouterOutlet`, `RouterLink`, `ActivatedRoute`.

**Basic Setup**:
```typescript
const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'users/:id', component: UserDetailComponent },
  { path: '**', component: NotFoundComponent } // Wildcard
];
```

## 47. What is Lazy Loading?
A technique where feature modules are loaded asynchronously (on demand) when the user navigates to their route.
**Benefit**: Reduces initial bundle size -> Faster initial load time.

**Syntax (Angular 15+)**:
```typescript
{
  path: 'customers',
  loadChildren: () => import('./customers/customers.module').then(m => m.CustomersModule)
}
```

## 48. What is Route Guard?
Interfaces that can tell the router whether to allow navigation to or from a route.
**Use Case**: Authentication check (Login required).

## 49. Types of Guards?
1.  **CanActivate**: Can user visit this route?
2.  **CanActivateChild**: Can user visit child routes?
3.  **CanDeactivate**: Can user leave this route? (Unsaved changes check).
4.  **CanLoad**: Can user load this lazy module? prevents loading code if unauthorized.
5.  **Resolve**: Fetch data before route is activated.

## 50. What is Router Outlet?
`<router-outlet></router-outlet>`
A placeholder directive where the router displays the component for the active route.

---

## 51. Difference between Template-Driven and Reactive Forms?

| Feature | Template-Driven | Reactive Forms |
| :--- | :--- | :--- |
| **Philosophy** | Asynchronous, implicitly created. | Synchronous, explicitly created. |
| **Logic Location** | HTML (`[(ngModel)]`, validations). | TypeScript (`FormControl`, `FormGroup`). |
| **Data Model** | Mutable. | Immutable. |
| **Testing** | Harder (depends on DOM). | Easier (no DOM needed). |
| **Scalability** | Good for simple forms. | Essential for complex forms. |

## 52. What is FormGroup?
Tracks the value and validity state of a group of `FormControl` instances. (Usually represents the entire form).

```typescript
this.profileForm = new FormGroup({
  firstName: new FormControl(''),
  lastName: new FormControl(''),
});
```

## 53. What is FormControl?
Tracks the value and validation status of an individual form control.

## 54. How do you add custom validator?
A validator is a function that processes a `FormControl` and returns an error map or null.

```typescript
export function forbiddenNameValidator(nameRe: RegExp): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const forbidden = nameRe.test(control.value);
    return forbidden ? { forbiddenName: { value: control.value } } : null;
  };
}
```

## 55. How do you handle dynamic forms?
Using `FormArray`. It manages an array of `FormControl`, `FormGroup`, or `FormArray`.
Useful for "Add Another Item" functionality.

```typescript
get aliases() { return this.profileForm.get('aliases') as FormArray; }
addAlias() { this.aliases.push(new FormControl('')); }
```

---

## 56. What is Change Detection?
Mechanism by which Angular synchronizes the state of the application with the UI.
Angular checks for changes in data bindings and updates the DOM.

## 57. Default vs OnPush Change Detection?
*   **Default**:
    *   Checks every component in the tree on every event (click, timer, xhr).
    *   Conservative (safest but slower).
*   **OnPush**:
    *   Checks only when:
        1.  Input (`@Input`) reference changes.
        2.  Event emitted from component or child.
        3.  Async pipe emits.
        4.  Manually triggered (`markForCheck()`).
    *   **Performance**: Much faster for large apps.

## 58. What is Zone.js?
A library that monkey-patches browser asynchronous APIs (setTimeout, promises, DOM events).
It notifies Angular ("Something happened!") so Angular can run Change Detection.

## 59. How do you improve Angular performance?
1.  **OnPush Strategy**: Reduce checks.
2.  **Lazy Loading**: Load modules on demand.
3.  **AOT (Ahead-of-Time)**: Compile templates during build (smaller bundle, faster render).
4.  **TrackBy**: Optimize `*ngFor` DOM manipulations.
5.  **Pure Pipes**: Use pipes for heavy computations in templates.
6.  **Unsubscribe**: Prevent memory leaks.
7.  **Web Workers**: Offload heavy computations from main thread.

## 60. How do you prevent memory leaks?
1.  **Unsubscribe** from Observables (`Subscription.unsubscribe`, `async` pipe, `takeUntil`).
2.  **Clear Timers** (`clearTimeout`).
3.  **Detach DOM Listeners** (`removeEventListener`).
4.  **Nullify References** in `ngOnDestroy`.
