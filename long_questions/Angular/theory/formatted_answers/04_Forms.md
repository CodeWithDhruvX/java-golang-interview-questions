# 🟡 Angular Forms

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Infosys): Template-driven forms, basic validation
> - 🚀 **Product-Based** (Razorpay, Zepto, Meesho): Reactive forms, custom validators, dynamic forms
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

### 1. What are template-driven forms? 🟢 | 🏭

"**Template-driven forms** use Angular directives in the HTML template to build forms. The form model is created implicitly — Angular creates `FormGroup` and `FormControl` instances behind the scenes based on `ngModel` directives.

```html
<!-- Template-driven form -->
<form #userForm="ngForm" (ngSubmit)="onSubmit(userForm)">
  <input name="username" ngModel required minlength="3" />
  <input name="email" ngModel required email />
  <button type="submit" [disabled]="userForm.invalid">Submit</button>
</form>
```

```typescript
onSubmit(form: NgForm): void {
  if (form.valid) {
    console.log(form.value); // { username: '...', email: '...' }
  }
}
```

They are simple and fast to set up — good for **basic forms with simple validation**."

#### In Depth
Template-driven forms are **asynchronous** — the form model builds up during Angular's template change detection cycle. This is why accessing the form's value synchronously in the constructor or early lifecycle hooks may not work as expected. They are harder to **unit test** because the form logic lives in the HTML, not in testable TypeScript. For complex forms, I always prefer reactive forms.

---

### 2. What are reactive forms? 🟡 | 🏭🚀

"**Reactive forms** (also called Model-driven forms) define the form model **explicitly in TypeScript**. The HTML template simply binds to the pre-defined model.

```typescript
// component.ts
import { FormGroup, FormControl, Validators } from '@angular/forms';

userForm = new FormGroup({
  username: new FormControl('', [Validators.required, Validators.minLength(3)]),
  email: new FormControl('', [Validators.required, Validators.email]),
  password: new FormControl('', Validators.required),
});

onSubmit(): void {
  if (this.userForm.valid) {
    const formData = this.userForm.value;
  }
}
```

```html
<form [formGroup]="userForm" (ngSubmit)="onSubmit()">
  <input formControlName="username" />
  <span *ngIf="userForm.get('username')?.invalid">Invalid username</span>
  <input formControlName="email" />
  <button type="submit">Submit</button>
</form>
```

Reactive forms are **synchronous**, fully **type-safe**, and far **easier to unit test**."

#### In Depth
Reactive forms are built on the **Observable pattern** — `valueChanges` and `statusChanges` are RxJS observables. This unlocks powerful patterns: auto-save forms, search-as-you-type with `debounceTime()`, cross-field validation with `combineLatest()`, and conditional validation rules based on other fields. This reactive nature is why I always prefer reactive forms for anything beyond a simple contact form.

---

### 3. Difference between template-driven and reactive forms? 🟡 | 🏭🚀

| Aspect | Template-Driven | Reactive |
|---|---|---|
| Form model | Implicit (HTML) | Explicit (TypeScript) |
| Synchronicity | Asynchronous | Synchronous |
| Validation | HTML attributes | TypeScript validators |
| Unit testing | Hard (DOM-dependent) | Easy (testable in TS) |
| Dynamic forms | Difficult | Easy (`FormArray`) |
| Scalability | Small forms | Complex forms |
| Module import | `FormsModule` | `ReactiveFormsModule` |

"I choose **template-driven** for simple, quick forms with minimal validation (e.g., a newsletter signup). I use **reactive forms** for everything else — login forms, multi-step wizards, dynamic form builders. The explicit TypeScript model makes complex validation and testing much cleaner."

#### In Depth
Both approaches ultimately create the same underlying model (`FormGroup`, `FormControl`). The difference is where control is. Reactive forms give me **full programmatic control** — I can reset, patch, enable/disable controls, run validation, and subscribe to changes from TypeScript without touching the template. This is critical for **enterprise applications** where business rules change form behavior dynamically.

---

### 4. What is FormGroup? 🟢 | 🏭

"**`FormGroup`** is a collection of `FormControl` instances that together represent a logical form or a section of a form. The group's status (`valid`, `invalid`, `pristine`, `dirty`) is computed from all its children.

```typescript
loginForm = new FormGroup({
  email: new FormControl('', [Validators.required, Validators.email]),
  password: new FormControl('', [Validators.required, Validators.minLength(8)])
});

// Accessing values
const email = this.loginForm.get('email')?.value;
const allValues = this.loginForm.value; // { email: '...', password: '...' }

// Setting values
this.loginForm.setValue({ email: 'test@example.com', password: '123456' });
this.loginForm.patchValue({ email: 'new@email.com' }); // Partial update
```"

#### In Depth
`FormGroup` propagates validity **upward** — a group is valid only if ALL its controls are valid. This means I can disable a submit button with `[disabled]="loginForm.invalid"` and it automatically reflects the state of all controls. `FormGroup` can be **nested** inside another `FormGroup` to model complex object structures like addresses within user profiles.

---

### 5. What is FormControl? 🟢 | 🏭

"**`FormControl`** is the atomic unit of a form — it tracks the value and validation state of a single form field.

```typescript
// Creating with initial value and validators
const emailControl = new FormControl('', {
  validators: [Validators.required, Validators.email],
  asyncValidators: [this.emailExistsValidator()],
  updateOn: 'blur' // Validate on blur, not on every keystroke
});

// State observation
console.log(emailControl.value);   // Current value
console.log(emailControl.valid);   // true/false
console.log(emailControl.touched); // true if user focused then left
console.log(emailControl.dirty);   // true if user changed the value

// Reacting to changes
emailControl.valueChanges.pipe(
  debounceTime(300),
  distinctUntilChanged()
).subscribe(value => this.onEmailChange(value));
```"

#### In Depth
`FormControl`'s `updateOn` option is powerful. Setting `updateOn: 'blur'` delays validation until the user leaves the field — this avoids showing error messages while the user is still typing. Setting `updateOn: 'submit'` shows all errors only when the form is submitted. These options give precise control over the **validation UX** without writing custom debounce logic.

---

### 6. What is FormBuilder? 🟢 | 🏭

"**`FormBuilder`** is a helper service that simplifies creating `FormGroup`, `FormControl`, and `FormArray` instances with less boilerplate.

```typescript
// Without FormBuilder — verbose
loginForm = new FormGroup({
  email: new FormControl('', Validators.required),
  password: new FormControl('', Validators.required)
});

// With FormBuilder — cleaner
constructor(private fb: FormBuilder) {}

loginForm = this.fb.group({
  email: ['', [Validators.required, Validators.email]],
  password: ['', Validators.required],
  address: this.fb.group({
    street: [''],
    city: ['', Validators.required]
  })
});
```

The shorthand array `['initialValue', validators, asyncValidators]` is the standard pattern in most Angular apps."

#### In Depth
Angular 14 introduced **typed reactive forms** (`FormBuilder.group<T>({...})`). The form's `.value` and `.controls` are now **fully typed** based on the generic parameter. `UntypedFormBuilder` exists for backward compatibility. Typed forms catch type mismatches at compile time — for example, setting a `number` control to a `string` is now a TypeScript error. This is a major DX improvement for large forms.

---

### 7. How to add form validation? 🟢 | 🏭🚀

"Angular provides **built-in validators** via `Validators`:

```typescript
email: ['', [
  Validators.required,        // Cannot be empty
  Validators.email,           // Must be valid email format
  Validators.maxLength(100),  // Max character limit
]],
age: ['', [
  Validators.required,
  Validators.min(18),         // Minimum value
  Validators.max(120),
  Validators.pattern(/^\d+$/) // Must be digits only
]]
```

Displaying errors in the template:

```html
<div *ngIf="email.invalid && (email.dirty || email.touched)">
  <span *ngIf="email.errors?.['required']">Email is required</span>
  <span *ngIf="email.errors?.['email']">Invalid email format</span>
  <span *ngIf="email.errors?.['maxLength']">Too long (max 100 chars)</span>
</div>
```"

#### In Depth
The pattern `email.dirty || email.touched` is important — it avoids showing validation errors on fields the user hasn't interacted with yet. `dirty` becomes `true` when the user changes the value; `touched` becomes `true` when the user focuses and then leaves the field. Showing errors before the user types (`submitted` flag is better for submit-based validation UX) degrades UX by overwhelming the user with red errors before they start.

---

### 8. How to write custom form validators? 🟡 | 🏭🚀

"A custom validator is a function that returns `null` (valid) or an error object (invalid):

```typescript
// Sync custom validator
function noSpacesValidator(control: AbstractControl): ValidationErrors | null {
  if (control.value && control.value.includes(' ')) {
    return { noSpaces: 'Username cannot contain spaces' };
  }
  return null;
}

// Group-level validator (cross-field)
function passwordMatchValidator(group: AbstractControl): ValidationErrors | null {
  const password = group.get('password')?.value;
  const confirm = group.get('confirmPassword')?.value;
  if (password !== confirm) {
    return { passwordMismatch: true };
  }
  return null;
}

// Async validator (e.g., check username availability via API)
function usernameAvailabilityValidator(userService: UserService): AsyncValidatorFn {
  return (control: AbstractControl): Observable<ValidationErrors | null> => {
    return timer(300).pipe( // Debounce API calls
      switchMap(() => userService.checkUsername(control.value)),
      map(taken => taken ? { usernameTaken: true } : null)
    );
  };
}

// Usage
username: ['', [noSpacesValidator], [usernameAvailabilityValidator(this.userService)]]
```"

#### In Depth
Async validators trigger an HTTP call on every keystroke without debouncing — which floods the API. The `timer(300).pipe(switchMap(...))` pattern inside the validator acts as a debounce: if a new value arrives within 300ms, `switchMap` cancels the previous request and starts fresh. Angular also sets `control.status` to `'PENDING'` while async validation runs, which I use to show a spinner: `<span *ngIf="username.pending">Checking...</span>`.

---

### 9. How to dynamically add form controls? 🟡 | 🚀

"I use **`FormArray`** for dynamic lists of controls:

```typescript
import { FormArray, FormGroup, FormBuilder } from '@angular/forms';

form = this.fb.group({
  emails: this.fb.array([
    this.fb.control('', Validators.email) // Start with one email
  ])
});

get emails(): FormArray {
  return this.form.get('emails') as FormArray;
}

addEmail(): void {
  this.emails.push(this.fb.control('', Validators.email));
}

removeEmail(index: number): void {
  this.emails.removeAt(index);
}
```

```html
<div formArrayName="emails">
  <div *ngFor="let email of emails.controls; let i = index">
    <input [formControlName]="i" placeholder="Email {{ i + 1 }}" />
    <button (click)="removeEmail(i)">Remove</button>
  </div>
</div>
<button (click)="addEmail()">+ Add Email</button>
```"

#### In Depth
`FormArray` is powerful but has a quirk: its `formControlName` must be the **index** of the item when used with `*ngFor`. This means removing an item by index requires care — all subsequent items' indices shift down. When using `trackBy` with `*ngFor` and `FormArray`, Angular still re-renders the correct input because `formControlName` is index-based. For complex items (objects with multiple fields), each `FormArray` element is usually a `FormGroup`.

---

### 10. What is the difference between `setValue()` and `patchValue()`? 🟡 | 🏭🚀

"Both set values on a form group, but they differ in strictness:

- **`setValue()`** — Sets ALL controls and **throws an error** if you miss any or provide extra keys.
- **`patchValue()`** — Sets only the **provided controls** and ignores missing ones gracefully.

```typescript
// Assume form has: name, email, phone

// setValue — ALL fields must be provided
this.form.setValue({ name: 'John', email: 'j@e.com', phone: '123' }); // ✅
this.form.setValue({ name: 'John' }); // ❌ Throws — phone and email missing

// patchValue — partial updates OK
this.form.patchValue({ name: 'John' }); // ✅ email and phone unchanged
```

I use `patchValue()` when loading partial data from an API (e.g., the server returns only editable fields, not all form fields) and `setValue()` when I want strict validation that all fields are accounted for."

#### In Depth
A common pattern for **pre-populating edit forms** from an API:

```typescript
ngOnInit(): void {
  this.userService.getUser(this.userId).subscribe(user => {
    this.form.patchValue(user); // Safe — extra API properties are ignored
  });
}
```

`patchValue` is forgiving: if the server adds new fields that aren't in the form, they're silently ignored. `setValue` would throw, breaking the form if the API response shape changes — making `patchValue` safer for API integration patterns.

---

### 11. How to build multi-step forms with state persistence? 🔴 | 🚀

"I build a **wizard form** by maintaining one `FormGroup` at the parent level and slicing out sub-groups for each step:

```typescript
// Wizard parent component
wizardForm = this.fb.group({
  step1: this.fb.group({ name: ['', Validators.required], email: [''] }),
  step2: this.fb.group({ address: [''], city: ['', Validators.required] }),
  step3: this.fb.group({ cardNumber: [''], expiryDate: [''] }),
});

currentStep = 0;
steps = ['step1', 'step2', 'step3'];

get currentStepGroup(): FormGroup {
  return this.wizardForm.get(this.steps[this.currentStep]) as FormGroup;
}

next(): void {
  if (this.currentStepGroup.valid) {
    this.currentStep++;
  } else {
    this.currentStepGroup.markAllAsTouched(); // Show all errors
  }
}

// Persist to localStorage between sessions
saveProgress(): void {
  localStorage.setItem('wizardForm', JSON.stringify(this.wizardForm.value));
}
```

State persists in the parent's `FormGroup`, so navigating between steps preserves all entered data."

#### In Depth
For multi-session persistence (user returns after closing browser), I serialize `wizardForm.value` to `localStorage` on every `valueChanges` event with `debounceTime(500)`. On init, I call `patchValue` from stored data. The `CanDeactivate` guard prevents navigation away with unsaved changes. This pattern gives a seamless resume experience without a backend dependency during form filling.

---
