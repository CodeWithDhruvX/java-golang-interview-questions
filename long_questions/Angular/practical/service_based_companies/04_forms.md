# 📘 04 — Forms
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Template-driven vs Reactive forms — when to use each
- `FormControl`, `FormGroup`, `FormArray`, `FormBuilder`
- Built-in validators and custom validators
- Form state properties: `valid`, `invalid`, `touched`, `dirty`, `pristine`
- Displaying validation error messages

---

## ❓ Most Asked Questions

### Q1. What is the difference between Template-Driven and Reactive Forms?

| Feature | Template-Driven | Reactive |
|---------|----------------|---------|
| Setup | Simple — minimal TS code | More verbose — define in TS |
| Form model | Implicit (Angular creates it) | Explicit (`FormGroup`, `FormControl`) |
| Validation | HTML attributes (`required`) | Validator functions in TS |
| Testability | Hard to unit test | Easy to unit test |
| Scalability | Good for simple forms | Better for complex/dynamic forms |
| Module needed | `FormsModule` | `ReactiveFormsModule` |
| When to use | Simple login/search forms | Complex multi-step, dynamic forms |

---

### Q2. Show a complete Reactive Form example.

```typescript
// product-form.component.ts
import { FormBuilder, FormGroup, Validators, AbstractControl } from '@angular/forms';

@Component({
  selector: 'app-product-form',
  templateUrl: './product-form.component.html'
})
export class ProductFormComponent implements OnInit {
  productForm: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.productForm = this.fb.group({
      name: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(100)]],
      price: [null, [Validators.required, Validators.min(0.01)]],
      category: ['electronics', Validators.required],
      description: [''],
      stock: [0, [Validators.required, Validators.min(0), Validators.pattern(/^\d+$/)]]
    });
  }

  // Convenience getters for cleaner template access
  get name(): AbstractControl { return this.productForm.get('name')!; }
  get price(): AbstractControl { return this.productForm.get('price')!; }

  onSubmit(): void {
    if (this.productForm.valid) {
      console.log(this.productForm.value);
      // Send to API
    } else {
      this.productForm.markAllAsTouched();  // show all validation errors
    }
  }

  onReset(): void {
    this.productForm.reset({ category: 'electronics', stock: 0 });
  }
}
```

```html
<!-- product-form.component.html -->
<form [formGroup]="productForm" (ngSubmit)="onSubmit()">

  <div class="form-group">
    <label>Product Name *</label>
    <input formControlName="name" type="text" placeholder="Enter name" />
    <div *ngIf="name.invalid && name.touched" class="error">
      <span *ngIf="name.errors?.['required']">Name is required.</span>
      <span *ngIf="name.errors?.['minlength']">
        Minimum {{ name.errors?.['minlength'].requiredLength }} characters.
      </span>
    </div>
  </div>

  <div class="form-group">
    <label>Price *</label>
    <input formControlName="price" type="number" placeholder="0.00" />
    <div *ngIf="price.invalid && price.touched" class="error">
      <span *ngIf="price.errors?.['required']">Price is required.</span>
      <span *ngIf="price.errors?.['min']">Price must be greater than 0.</span>
    </div>
  </div>

  <button type="submit" [disabled]="productForm.invalid">Save Product</button>
  <button type="button" (click)="onReset()">Reset</button>
</form>
```

---

### Q3. Show a Template-Driven Form example.

```html
<!-- login.component.html -->
<form #loginForm="ngForm" (ngSubmit)="onLogin(loginForm)">

  <div class="form-group">
    <label>Email</label>
    <input
      name="email"
      [(ngModel)]="credentials.email"
      #email="ngModel"
      type="email"
      required
      email
    />
    <div *ngIf="email.invalid && email.touched" class="error">
      <span *ngIf="email.errors?.['required']">Email is required.</span>
      <span *ngIf="email.errors?.['email']">Invalid email format.</span>
    </div>
  </div>

  <div class="form-group">
    <label>Password</label>
    <input
      name="password"
      [(ngModel)]="credentials.password"
      #password="ngModel"
      type="password"
      required
      minlength="8"
    />
    <div *ngIf="password.invalid && password.touched" class="error">
      <span *ngIf="password.errors?.['required']">Password is required.</span>
      <span *ngIf="password.errors?.['minlength']">Minimum 8 characters.</span>
    </div>
  </div>

  <button type="submit" [disabled]="loginForm.invalid">Login</button>
</form>
```

```typescript
// login.component.ts
export class LoginComponent {
  credentials = { email: '', password: '' };

  onLogin(form: NgForm): void {
    if (form.valid) {
      console.log(form.value);
    }
  }
}
```

---

### Q4. What are the form state properties?

```typescript
const ctrl = this.productForm.get('name')!;

// Validity states
ctrl.valid;      // true if all validators pass
ctrl.invalid;    // true if any validator fails
ctrl.errors;     // { required: true } or { minlength: { requiredLength: 3, actualLength: 1 } }

// Interaction states
ctrl.touched;   // true after the user focused and then blurred
ctrl.untouched; // initial state — user hasn't interacted yet
ctrl.dirty;     // true after the user changed the value
ctrl.pristine;  // true until the user types something

// Disabled state
ctrl.disabled;
ctrl.enabled;
```

```html
<!-- Show errors only after user has interacted -->
<div *ngIf="name.invalid && (name.dirty || name.touched)">
  {{ name.errors | json }}
</div>
```

---

### Q5. How do you create a custom validator?

```typescript
// validators/password-strength.validator.ts

// Function-based validator (synchronous)
export function passwordStrengthValidator(control: AbstractControl): ValidationErrors | null {
  const value: string = control.value || '';
  const hasUpperCase = /[A-Z]/.test(value);
  const hasLowerCase = /[a-z]/.test(value);
  const hasDigit = /\d/.test(value);
  const hasSpecial = /[!@#$%^&*]/.test(value);

  const valid = hasUpperCase && hasLowerCase && hasDigit && hasSpecial;
  return valid ? null : {
    passwordStrength: {
      hasUpperCase,
      hasLowerCase,
      hasDigit,
      hasSpecial
    }
  };
}

// Cross-field validator — passwords must match
export function passwordMatchValidator(group: AbstractControl): ValidationErrors | null {
  const password = group.get('password')?.value;
  const confirm = group.get('confirmPassword')?.value;
  return password === confirm ? null : { passwordMismatch: true };
}
```

```typescript
// Usage in form
this.registerForm = this.fb.group({
  password: ['', [Validators.required, passwordStrengthValidator]],
  confirmPassword: ['', Validators.required]
}, { validators: passwordMatchValidator });
```

---

### Q6. What is `FormArray`? When do you use it?

```typescript
// Dynamic list of fields (e.g., add multiple phone numbers)
@Component({ templateUrl: './profile.component.html' })
export class ProfileComponent implements OnInit {
  profileForm: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.profileForm = this.fb.group({
      name: ['', Validators.required],
      phones: this.fb.array([
        this.createPhoneGroup()  // at least one phone
      ])
    });
  }

  get phones(): FormArray {
    return this.profileForm.get('phones') as FormArray;
  }

  createPhoneGroup(): FormGroup {
    return this.fb.group({
      type: ['mobile', Validators.required],
      number: ['', [Validators.required, Validators.pattern(/^\d{10}$/)]]
    });
  }

  addPhone(): void {
    this.phones.push(this.createPhoneGroup());
  }

  removePhone(index: number): void {
    this.phones.removeAt(index);
  }
}
```

```html
<div formArrayName="phones">
  <div *ngFor="let phone of phones.controls; let i = index" [formGroupName]="i">
    <select formControlName="type">
      <option value="mobile">Mobile</option>
      <option value="home">Home</option>
    </select>
    <input formControlName="number" placeholder="10-digit number" />
    <button type="button" (click)="removePhone(i)">Remove</button>
  </div>
</div>
<button type="button" (click)="addPhone()">+ Add Phone</button>
```

---

### Q7. How do you implement async validators?

```typescript
// Check if username is already taken (API call)
export function usernameAvailableValidator(userService: UserService): AsyncValidatorFn {
  return (control: AbstractControl): Observable<ValidationErrors | null> => {
    if (!control.value) return of(null);

    return timer(400).pipe(  // debounce 400ms to avoid API spam
      switchMap(() => userService.checkUsername(control.value)),
      map(isTaken => isTaken ? { usernameTaken: true } : null),
      catchError(() => of(null))  // treat errors as validation pass
    );
  };
}

// Apply async validator
this.registerForm = this.fb.group({
  username: [
    '',
    [Validators.required, Validators.minLength(3)],       // sync validators
    [usernameAvailableValidator(this.userService)]         // async validators
  ]
});
```

```html
<input formControlName="username" />
<span *ngIf="username.pending">Checking availability...</span>
<span *ngIf="username.errors?.['usernameTaken']">Username already taken!</span>
```
