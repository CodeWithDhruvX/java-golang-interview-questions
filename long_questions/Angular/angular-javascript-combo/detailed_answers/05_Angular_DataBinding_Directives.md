# 🔗 Angular Data Binding & Directives (Detailed Answers)

## 16. Types of data binding?

1.  **Interpolation** (`{{ value }}`): One-way, Model -> View.
2.  **Property Binding** (`[property]="value"`): One-way, Model -> View.
3.  **Event Binding** (`(event)="handler()"`): One-way, View -> Model.
4.  **Two-way Binding** (`[(ngModel)]="property"`): Two-way, Model <-> View.

## 17. Difference between property and attribute binding?
*   **Attribute**: Set in HTML initially. Does not change once DOM is loaded. (`colspan`, `aria-*`). `[attr.colspan]="colSpan"`.
*   **Property**: Current value of the DOM object. Changes dynamically. (`value`, `src`, `hidden`). `[value]="firstName"`.

## 18. What is two-way binding?
Combines Property Binding (`[ ]`) and Event Binding (`( )`).
Any change in the **UI (View)** updates the **Model (Component)** automatically, and vice-versa.
Uses `ngModel` directive (from `FormsModule`).

```html
<input [(ngModel)]="username">
<p>Hello {{ username }}!</p>
```

## 19. How does `[(ngModel)]` work internally?
It's syntactic sugar for binding to a property named `ngModel` and listening to an event named `ngModelChange`.

```html
<input [ngModel]="username" (ngModelChange)="username = $event">
```

---

## 20. Difference between structural and attribute directives?

| Type | Syntax | Purpose | Example |
| :--- | :--- | :--- | :--- |
| **Structural** | `*` prefix | Change DOM structure (Add/Remove elements). | `*ngIf`, `*ngFor` |
| **Attribute** | `[ ]` syntax | Change appearance/behavior of an element. | `ngClass`, `ngStyle`, `[disabled]` |

## 21. What is *ngIf?
Conditionally adds or removes an element from the DOM based on a boolean expression.
Prefer over `[hidden]` because `[hidden]` keeps the element in DOM (consumes resources), whereas `*ngIf` removes it completely.

```html
<div *ngIf="isLoggedIn; else loginTemplate">
  Welcome, User!
</div>
<ng-template #loginTemplate>
  Please login.
</ng-template>
```

## 22. What is *ngFor?
Repeats a portion of the HTML template once for each item in an iterable list (array).
Exports local variables: `index`, `first`, `last`, `even`, `odd`.

```html
<ul>
  <li *ngFor="let user of users; let i = index; let isOdd = odd" 
      [class.highlight]="isOdd">
    {{ i + 1 }}. {{ user.name }}
  </li>
</ul>
```

## 23. What is trackBy and why is it used?
It's a function that returns a unique identifier for each item in a list (e.g., `item.id`).
**Performance Boost**: Without `trackBy`, if the array reference changes, Angular re-renders the *entire* list (destroying DOM). With `trackBy`, Angular only updates the changed items.

```typescript
trackById(index: number, item: any): number {
  return item.id;
}
```
```html
<li *ngFor="let item of items; trackBy: trackById">...</li>
```

## 24. How do you create custom directive?
Use `@Directive` decorator. Inject `ElementRef` to access the host element.

**Example: Highlight Directive**
```typescript
@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {
  constructor(private el: ElementRef) {}

  @HostListener('mouseenter') onMouseEnter() {
    this.highlight('yellow');
  }

  @HostListener('mouseleave') onMouseLeave() {
    this.highlight(null);
  }

  private highlight(color: string | null) {
    this.el.nativeElement.style.backgroundColor = color;
  }
}
```
**Usage:**
```html
<p appHighlight>Hover me!</p>
```
