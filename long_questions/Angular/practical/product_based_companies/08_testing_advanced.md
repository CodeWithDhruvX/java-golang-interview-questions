# 📘 08 — Advanced Testing
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- `fakeAsync` and `tick` for testing async code
- Component Test Harnesses (`@angular/cdk/testing`)
- Testing NgRx — reducers, selectors, effects
- Angular Testing Library (`@testing-library/angular`)
- `spectator` for reduced boilerplate
- Testing Signals

---

## ❓ Most Asked Questions

### Q1. What is `fakeAsync` and `tick`? How do they work?

```typescript
// fakeAsync — runs test in a fake async zone, giving you CONTROL over time
// tick(ms) — advances virtual time by ms milliseconds

import { fakeAsync, tick, discardPeriodicTasks } from '@angular/core/testing';

describe('SearchComponent', () => {
  it('should debounce search input by 300ms', fakeAsync(() => {
    const fixture = TestBed.createComponent(SearchComponent);
    const component = fixture.componentInstance;
    const mockService = TestBed.inject(SearchService) as jasmine.SpyObj<SearchService>;

    // User types
    component.searchControl.setValue('angular');
    fixture.detectChanges();

    // Before 300ms — no search yet
    expect(mockService.search).not.toHaveBeenCalled();

    tick(200);  // advance time 200ms
    expect(mockService.search).not.toHaveBeenCalled();

    tick(100);  // now 300ms total
    expect(mockService.search).toHaveBeenCalledWith('angular');
  }));

  it('should refresh automatically every 30 seconds', fakeAsync(() => {
    const fixture = TestBed.createComponent(DashboardComponent);
    const mockService = TestBed.inject(DataService) as jasmine.SpyObj<DataService>;

    tick(30000);  // advance 30 seconds
    expect(mockService.refresh).toHaveBeenCalledTimes(1);

    tick(30000);  // advance another 30 seconds
    expect(mockService.refresh).toHaveBeenCalledTimes(2);

    discardPeriodicTasks();  // clean up periodic tasks (interval/timer)
  }));
});
```

---

### Q2. How do you test NgRx reducers, selectors, and effects?

```typescript
// Testing REDUCERS — pure functions, no TestBed needed
describe('productReducer', () => {
  it('should add a product on addItem', () => {
    const initialState: ProductState = { items: [], loading: false };
    const product = { id: '1', name: 'Laptop', price: 50000 };

    const newState = productReducer(
      initialState,
      ProductActions.addProduct({ product })
    );

    expect(newState.items.length).toBe(1);
    expect(newState.items[0]).toEqual(product);
    expect(newState).not.toBe(initialState);  // immutable — new reference
  });

  it('should set loading on loadProducts', () => {
    const state = productReducer(initialState, ProductActions.loadProducts({}));
    expect(state.loading).toBe(true);
  });
});

// Testing SELECTORS
describe('product selectors', () => {
  const mockState = {
    products: {
      ids: ['1', '2'],
      entities: {
        '1': { id: '1', name: 'Laptop', price: 50000, category: 'electronics' },
        '2': { id: '2', name: 'Book', price: 500, category: 'books' }
      },
      loading: false
    }
  };

  it('should select all products', () => {
    const result = selectAllProducts.projector(mockState.products);
    expect(result.length).toBe(2);
  });

  it('should select expensive products', () => {
    const allProducts = selectAllProducts.projector(mockState.products);
    const result = selectExpensiveProducts.projector(allProducts);
    expect(result.length).toBe(1);
    expect(result[0].name).toBe('Laptop');
  });
});

// Testing EFFECTS
describe('ProductEffects', () => {
  let effects: ProductEffects;
  let actions$: Observable<Action>;
  let productServiceSpy: jasmine.SpyObj<ProductService>;

  beforeEach(() => {
    productServiceSpy = jasmine.createSpyObj('ProductService', ['getAll', 'delete']);

    TestBed.configureTestingModule({
      providers: [
        ProductEffects,
        provideMockActions(() => actions$),
        { provide: ProductService, useValue: productServiceSpy }
      ]
    });

    effects = TestBed.inject(ProductEffects);
  });

  it('should dispatch loadProductsSuccess on success', () => {
    const products = [{ id: '1', name: 'Laptop' }];
    productServiceSpy.getAll.and.returnValue(of(products));

    actions$ = of(ProductActions.loadProducts({}));

    effects.loadProducts$.subscribe(action => {
      expect(action).toEqual(ProductActions.loadProductsSuccess({ products }));
    });
  });

  it('should dispatch loadProductsFailure on error', () => {
    productServiceSpy.getAll.and.returnValue(throwError(() => new Error('Network error')));

    actions$ = of(ProductActions.loadProducts({}));

    effects.loadProducts$.subscribe(action => {
      expect(action).toEqual(ProductActions.loadProductsFailure({ error: 'Network error' }));
    });
  });
});
```

---

### Q3. How do you use Angular Component Test Harnesses?

```typescript
// Angular CDK provides typed harnesses for testing component interactions
// Benefits: stable API, works regardless of DOM structure changes

import { HarnessLoader } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatButtonHarness } from '@angular/material/button/testing';
import { MatInputHarness } from '@angular/material/input/testing';

describe('ProductFormComponent', () => {
  let loader: HarnessLoader;
  let fixture: ComponentFixture<ProductFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProductFormComponent, MatButtonModule, MatInputModule]
    }).compileComponents();

    fixture = TestBed.createComponent(ProductFormComponent);
    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should disable submit button when form is invalid', async () => {
    const submitButton = await loader.getHarness(
      MatButtonHarness.with({ text: 'Submit' })
    );
    expect(await submitButton.isDisabled()).toBe(true);
  });

  it('should enable submit after filling required fields', async () => {
    const nameInput = await loader.getHarness(
      MatInputHarness.with({ selector: '[data-testid="name-input"]' })
    );
    await nameInput.setValue('Laptop Pro');
    fixture.detectChanges();

    const submitButton = await loader.getHarness(MatButtonHarness.with({ text: 'Submit' }));
    expect(await submitButton.isDisabled()).toBe(false);
  });
});
```

---

### Q4. How do you test Angular Signals?

```typescript
// Signals are synchronous — easier to test than Observables!
describe('CounterComponent with Signals', () => {
  let component: CounterComponent;
  let fixture: ComponentFixture<CounterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CounterComponent]  // standalone component
    }).compileComponents();

    fixture = TestBed.createComponent(CounterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should have initial count of 0', () => {
    expect(component.count()).toBe(0);  // read signal value directly
  });

  it('should increment count on button click', () => {
    const btn = fixture.debugElement.query(By.css('button'));
    btn.triggerEventHandler('click', null);

    expect(component.count()).toBe(1);
  });

  it('should update computed doubled signal', () => {
    component.count.set(5);
    fixture.detectChanges();

    expect(component.doubled()).toBe(10);

    const doubledEl = fixture.debugElement.query(By.css('.doubled'));
    expect(doubledEl.nativeElement.textContent).toContain('10');
  });

  it('should accept signal input', () => {
    // For signal inputs — use TestBed.runInInjectionContext or fixture.componentRef.setInput
    fixture.componentRef.setInput('product', { id: '1', name: 'Test' });
    fixture.detectChanges();
    expect(component.product().name).toBe('Test');
  });
});
```

---

### Q5. How do you reduce test boilerplate with Spectator?

```typescript
// Without Spectator: lots of TestBed.configureTestingModule boilerplate
// With Spectator: declarative, concise test setup

import { createComponentFactory, Spectator } from '@ngneat/spectator';

describe('ProductCardComponent — with Spectator', () => {
  let spectator: Spectator<ProductCardComponent>;

  const createComponent = createComponentFactory({
    component: ProductCardComponent,
    imports: [CommonModule],
    mocks: [ProductService, CartService],  // auto-mocked — no jasmine.createSpyObj
    detectChanges: false,
  });

  beforeEach(() => {
    spectator = createComponent({
      props: { product: { id: '1', name: 'Laptop', price: 50000 } }
    });
    spectator.detectChanges();
  });

  it('should display product name', () => {
    expect(spectator.query('.product-name')).toHaveText('Laptop');
  });

  it('should call addToCart when button clicked', () => {
    const cartService = spectator.inject(CartService);
    spectator.click('[data-testid="add-to-cart"]');
    expect(cartService.addItem).toHaveBeenCalledWith({ id: '1', name: 'Laptop', price: 50000 });
  });

  it('should show price in INR', () => {
    expect(spectator.query('.price')).toHaveText('₹50,000');
  });
});
```

---

### Q6. What is `TestBed.runInInjectionContext`? (Angular 16+)

```typescript
// For testing functions that use inject() internally (guards, resolvers, etc.)

// route guard using inject()
export const authGuard = () => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return authService.isLoggedIn() || router.createUrlTree(['/login']);
};

// Testing it
describe('authGuard', () => {
  it('should allow access when logged in', () => {
    TestBed.configureTestingModule({
      providers: [
        { provide: AuthService, useValue: { isLoggedIn: () => true } }
      ]
    });

    const result = TestBed.runInInjectionContext(() => authGuard());
    expect(result).toBe(true);
  });

  it('should redirect when not logged in', () => {
    TestBed.configureTestingModule({
      providers: [
        { provide: AuthService, useValue: { isLoggedIn: () => false } }
      ]
    });

    const result = TestBed.runInInjectionContext(() => authGuard());
    expect(result).toBeInstanceOf(UrlTree);
  });
});
```
