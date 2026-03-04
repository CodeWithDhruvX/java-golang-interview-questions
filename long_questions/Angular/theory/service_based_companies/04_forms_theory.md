# 🗣️ Theory — Forms
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is the difference between template-driven and reactive forms in Angular?"

> *"Template-driven forms keep the form model implicit — you use ngModel in the template and Angular creates the underlying FormControl and FormGroup objects for you. They're simple to set up for small forms like a login page, but they're hard to test because the form logic is buried in the template, and they're hard to manage when forms get complex. Reactive forms are explicit — you define the FormGroup and FormControls in the component TypeScript class, and you bind the template to that model. The form logic is in code, which makes it unit-testable, composable, and easy to add dynamic behavior like adding fields at runtime. The Angular team recommends reactive forms for anything beyond a simple login or search box."*

---

## Q: "Walk me through FormGroup, FormControl, and FormBuilder."

> *"A FormControl is the smallest unit — it represents one input field and tracks its value, validation status, and user interaction state. A FormGroup is a collection of FormControls — it represents a whole form or a subsection of a form, and its validity is the combined validity of all its controls. You could build a form by manually instantiating FormGroup and FormControl with new keywords, but FormBuilder is a helper service that makes this concise. FormBuilder.group() takes an object where keys are control names and values are arrays — the first element is the initial value, the second is validators, the third is async validators. It's just syntactic sugar but it makes form definitions far less verbose and more readable."*

---

## Q: "How do Angular form validators work? How do you write a custom one?"

> *"Angular has built-in validators in the Validators class — required, minLength, maxLength, pattern, email, min, max. You pass them when creating a FormControl. For synchronous custom validators, you write a function that takes an AbstractControl and returns either null for valid, or an object like { passwordTooWeak: true } for invalid. Angular uses the return value to populate the control's errors object. For async validators — like checking if a username is already taken via an API — you return an Observable or Promise of the same null-or-errors pattern. Angular won't proceed to async validators until all sync validators pass. You attach validators as the second argument and async validators as the third when building controls."*

---

## Q: "What are the form state properties and when do you use them for error display?"

> *"Each FormControl has several status properties. Valid and invalid reflect whether all validators pass. Touched means the user focused and then blurred the field. Dirty means the user has changed the value from its initial state. Pristine is the opposite of dirty — the value hasn't been changed yet. These are crucial for good UX in error display. You don't want to show 'name is required' the moment the form loads — that's jarring. The pattern I follow is: show errors only when the control is invalid AND either touched or dirty. That way errors only appear after the user has interacted. When the user clicks submit without touching anything, I call markAllAsTouched() on the form group — this triggers all the error messages at once."*

---

## Q: "What is FormArray? When would you use it?"

> *"FormArray is like FormGroup but for indexed collections — arrays of form controls or form groups. The classic use case is dynamic form fields: let the user add multiple phone numbers, multiple addresses, or multiple line items in an invoice. You start with an array containing one entry, and add or remove entries at runtime by calling push() or removeAt(). Each item in the array can itself be a FormGroup with its own fields. In the template you iterate over the FormArray's controls with *ngFor and use formGroupName or formControlName with the index. It's significantly more flexible than trying to fake dynamic fields with static FormGroups."*

---

## Q: "How do you implement an async validator to check if a username is already taken?"

> *"You return a function that takes an AbstractControl and returns an Observable. Inside, you typically debounce the check using timer() and switchMap — you don't want to hit the API on every keystroke. You call the API with the control's current value, and if the username is taken, return an observable of { usernameTaken: true }. If it's available, return an observable of null. You also add catchError to return null if the API fails — you don't want a network error to block the user from submitting. Angular shows a pending state while the async validator is running, so you can display a spinner while the check is in progress. You attach async validators as the third argument to FormBuilder.group or the FormControl constructor."*

---

## Q: "What is ControlValueAccessor? Why is it important?"

> *"ControlValueAccessor is an interface that Angular uses to bridge between a FormControl and a DOM element. If you create a custom input component — say, a star rating widget or a date picker — and you want it to work with both template-driven forms and reactive forms, you implement ControlValueAccessor. The interface requires four methods: writeValue, which Angular calls to push a value into your component; registerOnChange, which gives you a callback to call when your component's value changes; registerOnTouched, which gives you a callback for when your component is blurred; and setDisabledState, which Angular calls when the form control's disabled state changes. Once you implement this plus provide the NG_VALUE_ACCESSOR token, your component works seamlessly with ngModel and formControl just like a native input."*
