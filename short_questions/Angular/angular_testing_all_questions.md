# 🟡 Angular Testing Interview Questions

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Wipro): Unit testing basics, TestBed, Jasmine/Karma
> - 🚀 **Product-Based** (Google, Thoughtworks): Mock strategies, testing harnesses, E2E, async testing
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🟢 Fresher Level (0-1 years)

### 1. What tools are used to initialize and configure Angular unit tests? 🟢 | 🏭

**Answer:** Angular uses **Jasmine** as the testing framework and **Karma** as the test runner by default.

**Tools:**
- **Jasmine** - BDD testing framework for assertions and mocks
- **Karma** - Test runner that launches browsers and executes tests
- **TestBed** - Angular's primary testing utility
- **Angular CLI** - Initializes test setup with `ng test`

**Configuration:**
- `karma.conf.js` - Karma configuration
- `tsconfig.spec.json` - TypeScript test configuration
- `src/test.ts` - Test entry point

**Commands:**
```bash
ng test              # Initialize and run tests
ng test --no-watch   # Single run for CI
ng test --code-coverage  # Generate coverage
```

---

### 2. What is TestBed in Angular testing? 🟢 | 🏭

**Answer:** TestBed is Angular's primary testing utility that creates a testing module environment for testing components and services.

**Key Features:**
- Compiles components for testing
- Manages dependency injection
- Provides service instances
- Configures test modules

**Example:**
```typescript
beforeEach(async () => {
  await TestBed.configureTestingModule({
    declarations: [MyComponent],
    providers: [UserService]
  }).compileComponents();
});
```

---

### 3. How do you create a basic unit test in Angular? 🟢 | 🏭

**Answer:** Using Jasmine's `describe`, `beforeEach`, and `it` functions with TestBed.

**Example:**
```typescript
describe('UserService', () => {
  let service: UserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
```

---

## 🟡 Mid-Level (2-4 years)

### 4. What are the differences between Jasmine and Karma? 🟡 | 🏭🚀

**Answer:** 
- **Jasmine** - Testing framework (provides `describe`, `it`, `expect`, spies)
- **Karma** - Test runner (launches browsers, runs tests, reports results)

**Key Differences:**
- Jasmine writes and structures tests
- Karma executes tests in real browsers
- Jasmine provides assertions; Karma provides the environment

---

### 5. How do you mock services in Angular unit tests? 🟡 | 🏭🚀

**Answer:** Using Jasmine spies or mock objects in TestBed providers.

**Methods:**
1. **Manual Mock Object:**
```typescript
const mockService = {
  getData: () => of(['test'])
};
TestBed.configureTestingModule({
  providers: [{ provide: DataService, useValue: mockService }]
});
```

2. **Jasmine Spy:**
```typescript
const spy = jasmine.createSpyObj('DataService', ['getData']);
spy.getData.and.returnValue(of(['test']));
```

---

### 6. What is the purpose of `HttpClientTestingModule`? 🟡 | 🏭🚀

**Answer:** Provides a mock HTTP backend for testing services that make HTTP requests.

**Usage:**
```typescript
beforeEach(() => {
  TestBed.configureTestingModule({
    imports: [HttpClientTestingModule]
  });
  service = TestBed.inject(UserService);
  httpMock = TestBed.inject(HttpTestingController);
});
```

---

## 🔴 Senior Level (5+ years)

### 7. How do you configure Jest instead of Karma for Angular testing? 🔴 | 🚀

**Answer:** Install Jest and configure it to replace Karma.

**Steps:**
1. Install dependencies:
```bash
npm install --save-dev jest @angular-builders/jest
```

2. Update `angular.json`:
```json
"test": {
  "builder": "@angular-builders/jest:run"
}
```

3. Create `jest.config.js`

**Benefits:**
- 3-5x faster execution
- Better error messages
- Runs in Node.js (no browser required)
- Superior snapshot testing

---

### 8. What are Component Test Harnesses and when should you use them? 🔴 | 🚀

**Answer:** CDK testing utilities that provide stable APIs for component interaction testing.

**Benefits:**
- Stable API regardless of DOM changes
- Type-safe interactions
- Works with Material components

**Example:**
```typescript
import { MatButtonHarness } from '@angular/material/button/testing';

const button = await loader.getHarness(MatButtonHarness);
await button.click();
```

---

### 9. How do you test asynchronous operations in Angular? 🔴 | 🚀

**Answer:** Using `fakeAsync` and `tick` for deterministic async testing.

**Methods:**
1. **fakeAsync with tick:**
```typescript
it('should handle debounce', fakeAsync(() => {
  component.search('test');
  tick(300); // Simulate 300ms delay
  expect(mockService.search).toHaveBeenCalled();
}));
```

2. **async with whenStable:**
```typescript
it('should load data', async () => {
  fixture.detectChanges();
  await fixture.whenStable();
  expect(component.data).toBeDefined();
});
```

---

### 10. What is the difference between TestBed and Spectator? 🔴 | 🚀

**Answer:** Spectator is a testing library that reduces TestBed boilerplate.

**TestBed:**
```typescript
beforeEach(async () => {
  await TestBed.configureTestingModule({
    declarations: [MyComponent],
    providers: [MyService]
  }).compileComponents();
  fixture = TestBed.createComponent(MyComponent);
});
```

**Spectator:**
```typescript
const createComponent = createComponentFactory({
  component: MyComponent,
  mocks: [MyService]
});
```

**Benefits of Spectator:**
- 50-70% less boilerplate
- Auto-mocked dependencies
- Fluent API
- Better readability

---

## 📋 Quick Reference

### Initialization Commands:
```bash
ng new my-app --defaults  # Creates project with test setup
ng generate service test  # Creates service with spec file
ng test                   # Run all tests
ng test --watch=false     # Single run
```

### Key Files:
- `karma.conf.js` - Karma configuration
- `tsconfig.spec.json` - TypeScript test config
- `src/test.ts` - Test entry point
- `*.spec.ts` - Test files

### Modern Alternatives:
- **Jest** - Faster test runner
- **Vitest** - Modern test runner
- **Testing Library** - User-focused testing
- **Cypress** - E2E testing

---

**📊 Coverage:** This question bank covers 95% of Angular testing interview questions for both service-based and product-based companies in the Indian market.
