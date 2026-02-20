# 🟡 Angular Testing

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Wipro): Unit testing basics, TestBed, Jasmine/Karma
> - 🚀 **Product-Based** (Google, Thoughtworks): Mock strategies, testing harnesses, E2E, async testing
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

### 1. What is unit testing in Angular? 🟢 | 🏭

"**Unit testing** in Angular verifies individual units of code (components, services, pipes, directives) in isolation — without the real dependencies (HTTP, router, databases).

Angular uses **Jasmine** as the testing framework (assertions and mocks) and **Karma** as the test runner (executes tests in a real browser).

```typescript
// user.service.spec.ts
describe('UserService', () => {
  let service: UserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should return formatted user name', () => {
    const user = { firstName: 'John', lastName: 'Doe' };
    expect(service.getFullName(user)).toBe('John Doe');
  });
});
```

Unit tests are fast (run in milliseconds) and run in CI on every commit."

#### In Depth
For modern Angular projects, I switch from Karma to **Jest** — Jest is 3–5x faster, runs in Node.js (no browser required), has better snapshot testing, and has superior error messages. The `@angular-builders/jest` package integrates Karma-based TestBed tests with Jest, requiring minimal changes to existing tests. I also use **pnpm --parallel** to run test suites across Nx monorepo packages concurrently.

---

### 2. How to test a component? 🟡 | 🏭🚀

"I use `TestBed` to create a testing module for the component:

```typescript
// product-card.component.spec.ts
describe('ProductCardComponent', () => {
  let component: ProductCardComponent;
  let fixture: ComponentFixture<ProductCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProductCardComponent], // Standalone component
    }).compileComponents();

    fixture = TestBed.createComponent(ProductCardComponent);
    component = fixture.componentInstance;

    // Set inputs
    component.product = { id: 1, name: 'Widget', price: 29.99 };
    fixture.detectChanges(); // Trigger initial CD and ngOnInit
  });

  it('should display product name', () => {
    const el = fixture.nativeElement.querySelector('.product-name');
    expect(el.textContent).toContain('Widget');
  });

  it('should emit addToCart event when button clicked', () => {
    spyOn(component.addToCart, 'emit');
    fixture.nativeElement.querySelector('.add-btn').click();
    expect(component.addToCart.emit).toHaveBeenCalledWith(component.product);
  });
});
```"

#### In Depth
**`fixture.detectChanges()`** is critical — it triggers Angular's change detection which runs `ngOnInit` and updates the template. Without it, the component is created but no lifecycle hooks run and the template is not rendered. I call it after setting `@Input()` values and after any data changes to ensure the DOM is updated before making assertions. For `OnPush` components, you must call `fixture.detectChanges()` explicitly after every data update.

---

### 3. What is TestBed? 🟢 | 🏭

"**`TestBed`** is Angular's primary testing utility — it creates a mini Angular application environment for testing. It compiles components, provides services, and manages the DI container for tests.

```typescript
await TestBed.configureTestingModule({
  declarations: [MyComponent],     // For module-based components
  imports: [ReactiveFormsModule, MyComponent], // For standalone
  providers: [
    { provide: UserService, useClass: MockUserService }, // Mock services
    { provide: Router, useValue: { navigate: jasmine.createSpy() } }
  ]
}).compileComponents();
```

Key methods:
- `TestBed.createComponent(X)` — Creates a component fixture
- `TestBed.inject(Service)` — Gets a service instance from the test injector
- `TestBed.overrideProvider(token, { useValue: mock })` — Override provider after setup"

#### In Depth
`TestBed.configureTestingModule` is expensive — it compiles the entire Angular module for each test file. For faster tests, use `TestBed.resetTestingModule()` sparingly and group related tests in the same `describe` block with one shared `beforeEach`. For standalone components, `TestBed.createComponent` without `compileComponents()` works directly, reducing setup time. Also, **shallow rendering** using `NO_ERRORS_SCHEMA` or `CUSTOM_ELEMENTS_SCHEMA` avoids the need to import all child components.

---

### 4. What is Jasmine? 🟢 | 🏭

"**Jasmine** is a BDD (Behavior-Driven Development) testing framework that provides:

```typescript
describe('Calculator', () => {  // Test suite
  let calculator: Calculator;

  beforeEach(() => {               // Setup before each test
    calculator = new Calculator();
  });

  afterEach(() => {                // Cleanup after each test
    // ...
  });

  it('should add two numbers', () => {  // Test case
    expect(calculator.add(2, 3)).toBe(5);     // Assertion
    expect(calculator.add(-1, 1)).toEqual(0);
    expect(calculator.add).toHaveBeenCalled(); // Spy assertion
  });

  xit('should handle division by zero', () => { // xIt = skipped test
    // ...
  });
});
```

Jasmine matchers: `toBe`, `toEqual`, `toBeTruthy`, `toContain`, `toThrow`, `toHaveBeenCalled`, `toHaveBeenCalledWith`."

#### In Depth
`toBe` uses `===` (strict equality) — good for primitives. `toEqual` performs deep structural comparison — use it for objects and arrays. Many bugs in test suites come from using `toBe` on objects (comparing references, not values). Jasmine `spyOn` creates a spy on an existing method, while `jasmine.createSpy()` creates a standalone spy. Always verify your spies with `expect(spy).toHaveBeenCalledWith(expectedArgs)` for precise assertions.

---

### 5. What is Karma? 🟢 | 🏭

"**Karma** is a **test runner** (not a framework) — it launches browsers, loads test files, runs them, and reports results. Angular CLI uses Karma with Jasmine by default.

```bash
ng test              # Runs tests in a real Chrome browser (watch mode)
ng test --no-watch   # Single run (for CI)
ng test --code-coverage  # Generate coverage report
```

Karma's `karma.conf.js` controls which browsers to use, coverage thresholds, and reporter formats.

For CI, I use **headless Chrome** (`ChromeHeadless`) to run tests without a display server."

#### In Depth
Karma's main limitation is **speed** — starting a real browser for every test run is slow. For large projects, I migrate from Karma+Jasmine to **Jest** or **Vitest** which run tests in Node.js (jsdom) without browser overhead. Jest runs 5–10x faster on large test suites and has better parallelization. The Angular team has been improving Jest integration and it's now the recommended choice for new projects.

---

### 6. How to test services? 🟡 | 🏭🚀

"Services are the easiest to test because they're plain TypeScript classes. I mock their dependencies:

```typescript
describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify(); // Ensure no outstanding HTTP requests
  });

  it('should fetch user by ID', () => {
    const mockUser = { id: '1', name: 'Alice' };

    service.getUser('1').subscribe(user => {
      expect(user).toEqual(mockUser);
    });

    // Mock the HTTP response
    const req = httpMock.expectOne('/api/users/1');
    expect(req.request.method).toBe('GET');
    req.flush(mockUser); // Respond with mock data
  });
});
```"

#### In Depth
`HttpTestingController.verify()` in `afterEach` ensures every expected HTTP call was made and every intercepted call was handled. Without it, a test that sets up an `expectOne` but never triggers the observable would silently pass — `verify()` catches this. For services with complex dependencies, I use `jasmine.createSpyObj('ServiceName', ['method1', 'method2'])` to create fully mocked dependencies without any real implementation.

---

### 7. What is end-to-end (e2e) testing? 🟡 | 🏭🚀

"**E2E testing** tests the entire application flow from the user's perspective — clicking buttons, filling forms, and verifying real UI outcomes in a real browser.

**Cypress** is the modern choice for Angular E2E tests:

```typescript
// cypress/e2e/login.cy.ts
describe('Login Flow', () => {
  beforeEach(() => {
    cy.visit('/login');
  });

  it('should log in successfully', () => {
    cy.get('[data-testid=\"email\"]').type('user@example.com');
    cy.get('[data-testid=\"password\"]').type('securePass123');
    cy.get('[data-testid=\"submit\"]').click();

    cy.url().should('include', '/dashboard');
    cy.get('.welcome-message').should('contain', 'Welcome, User');
  });

  it('should show error for invalid credentials', () => {
    cy.get('[data-testid=\"email\"]').type('wrong@email.com');
    cy.get('[data-testid=\"password\"]').type('wrongPass');
    cy.get('[data-testid=\"submit\"]').click();

    cy.get('.error-message').should('be.visible');
  });
});
```

E2E tests are **slower and flakier** than unit tests — I run them in CI on PRs, not on every commit."

#### In Depth
**`data-testid`** attributes are the best locator strategy for E2E tests — they're independent from CSS classes (which designers change), text content (which copywriters change), and DOM structure (which developers change). I make it a **team convention** to add `data-testid` to interactive elements when they're built, not retroactively when E2E tests break. This investment pays off massively in test stability across design iterations.

---

### 8. How to test asynchronous behavior? 🟡 | 🏭🚀

"Angular's `TestBed` provides two helpers for async testing:

**`fakeAsync` + `tick()`** — Simulates time passing synchronously:

```typescript
it('should show result after debounce', fakeAsync(() => {
  component.searchTerm = 'angular';
  fixture.detectChanges();

  tick(300); // Simulate 300ms passing (debounceTime)
  fixture.detectChanges();

  expect(component.results.length).toBeGreaterThan(0);
}));
```

**`async` + `whenStable()`** — Waits for real promises to resolve:

```typescript
it('should load data', async () => {
  fixture.detectChanges();
  await fixture.whenStable(); // Wait for async code to complete
  fixture.detectChanges();

  expect(component.isLoading).toBeFalse();
  expect(component.items.length).toBe(5);
});
```"

#### In Depth
`fakeAsync` is more deterministic — I prefer it because I control time explicitly. `tick(ms)` advances the virtual clock by ms, `flush()` runs all pending microtasks and macrotasks. `flushMicrotasks()` drains only microtasks (Promise chains). For RxJS-specific timing tests, the **`TestScheduler`** provides marble testing — a visual way to test observable sequences with precise timing: `'--a--b--c|'`.

---

### 9. How to mock services in unit tests? 🟡 | 🏭🚀

"Several mocking strategies:

**1. Manual mock object:**
```typescript
const mockAuthService = {
  isLoggedIn: () => true,
  getUser: () => of({ id: 1, name: 'Test User' }),
  logout: jasmine.createSpy('logout')
};

TestBed.configureTestingModule({
  providers: [{ provide: AuthService, useValue: mockAuthService }]
});
```

**2. Jasmine spyObj:**
```typescript
const authSpy = jasmine.createSpyObj('AuthService', ['isLoggedIn', 'logout']);
authSpy.isLoggedIn.and.returnValue(true);
authSpy.logout.and.returnValue(of(null));

providers: [{ provide: AuthService, useValue: authSpy }]
```

**3. useClass with MockClass:**
```typescript
class MockAuthService extends AuthService {
  override isLoggedIn() { return true; }
}

providers: [{ provide: AuthService, useClass: MockAuthService }]
```"

#### In Depth
The `jasmine.createSpyObj` approach is my preference for most cases — it creates a mock with **trackable methods** (I can verify `expect(authSpy.logout).toHaveBeenCalled()`). The key insight: **test the behavior, not the implementation**. I don't care which internal methods are called — I care that when `isLoggedIn` returns `false`, the component redirects to `/login`. Mocking lets me control the dependency's behavior to test all branches.

---
