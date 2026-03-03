# 📘 08 — Testing & DevOps
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Unit testing with Jasmine & Karma
- `TestBed` and `ComponentFixture`
- Mocking services with `jasmine.createSpyObj`
- Testing components, services, and pipes
- Angular CLI build commands and environments
- Basic CI/CD with Angular

---

## ❓ Most Asked Questions

### Q1. How do you write a unit test for an Angular component?

```typescript
// product-card.component.spec.ts
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { ProductCardComponent } from './product-card.component';

describe('ProductCardComponent', () => {
  let component: ProductCardComponent;
  let fixture: ComponentFixture<ProductCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ProductCardComponent],
      // imports: [CommonModule]  — add if component uses *ngIf etc.
    }).compileComponents();

    fixture = TestBed.createComponent(ProductCardComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display product name', () => {
    component.product = { id: '1', name: 'iPhone 15', price: 79999 };
    fixture.detectChanges();  // trigger change detection

    const nameEl = fixture.debugElement.query(By.css('.product-name'));
    expect(nameEl.nativeElement.textContent).toContain('iPhone 15');
  });

  it('should emit addToCart when button clicked', () => {
    component.product = { id: '1', name: 'iPhone 15', price: 79999 };
    fixture.detectChanges();

    spyOn(component.addedToCart, 'emit');

    const btn = fixture.debugElement.query(By.css('[data-testid="add-to-cart"]'));
    btn.triggerEventHandler('click', null);

    expect(component.addedToCart.emit).toHaveBeenCalledWith('1');
  });
});
```

---

### Q2. How do you test a service with HTTP calls?

```typescript
// product.service.spec.ts
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ProductService } from './product.service';

describe('ProductService', () => {
  let service: ProductService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [ProductService]
    });
    service = TestBed.inject(ProductService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();  // ensure no outstanding HTTP requests
  });

  it('should fetch products', () => {
    const mockProducts = [
      { id: '1', name: 'Laptop', price: 50000 },
      { id: '2', name: 'Phone', price: 20000 }
    ];

    service.getProducts().subscribe(products => {
      expect(products.length).toBe(2);
      expect(products[0].name).toBe('Laptop');
    });

    const req = httpMock.expectOne('https://api.example.com/products');
    expect(req.request.method).toBe('GET');
    req.flush(mockProducts);  // simulate API response
  });

  it('should handle HTTP error', () => {
    service.getProducts().subscribe({
      next: () => fail('Expected error'),
      error: (err) => expect(err.status).toBe(500)
    });

    const req = httpMock.expectOne('https://api.example.com/products');
    req.flush('Server Error', { status: 500, statusText: 'Internal Server Error' });
  });
});
```

---

### Q3. How do you mock a service in component tests?

```typescript
// cart.component.spec.ts
describe('CartComponent', () => {
  let component: CartComponent;
  let fixture: ComponentFixture<CartComponent>;
  let cartServiceSpy: jasmine.SpyObj<CartService>;

  beforeEach(async () => {
    // Create a spy object — all methods auto-mocked
    cartServiceSpy = jasmine.createSpyObj('CartService', ['addToCart', 'removeFromCart'], {
      cartItems$: of([{ productId: '1', name: 'Phone', quantity: 2, price: 20000 }]),
      cartTotal$: of(40000)
    });

    await TestBed.configureTestingModule({
      declarations: [CartComponent],
      imports: [CommonModule],
      providers: [
        { provide: CartService, useValue: cartServiceSpy }  // inject mock
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should call removeFromCart when remove button clicked', () => {
    const btn = fixture.debugElement.query(By.css('.remove-btn'));
    btn.triggerEventHandler('click', null);

    expect(cartServiceSpy.removeFromCart).toHaveBeenCalledWith('1');
  });
});
```

---

### Q4. How do you test a custom pipe?

```typescript
// truncate.pipe.spec.ts
import { TruncatePipe } from './truncate.pipe';

describe('TruncatePipe', () => {
  let pipe: TruncatePipe;

  beforeEach(() => {
    pipe = new TruncatePipe();  // pipes are plain classes — no TestBed needed
  });

  it('should create', () => expect(pipe).toBeTruthy());

  it('should return original string if shorter than limit', () => {
    expect(pipe.transform('Hello', 10)).toBe('Hello');
  });

  it('should truncate long strings', () => {
    const result = pipe.transform('Hello World this is a long string', 10);
    expect(result.length).toBeLessThanOrEqual(13);  // 10 chars + '...'
    expect(result).toContain('...');
  });

  it('should handle null/empty values', () => {
    expect(pipe.transform('', 10)).toBe('');
    expect(pipe.transform(null as any, 10)).toBe('');
  });
});
```

---

### Q5. What are Angular CLI build commands and environments?

```bash
# Development build (default)
ng serve
ng serve --port 4300 --open

# Staging build
ng build --configuration=staging

# Production build — enables AOT, minification, tree-shaking, source maps off
ng build --configuration=production

# Run unit tests
ng test                          # watch mode
ng test --no-watch --code-coverage  # single run + coverage report

# Run e2e tests
ng e2e
```

**Environment files:**

```typescript
// src/environments/environment.ts (development)
export const environment = {
  production: false,
  apiUrl: 'http://localhost:8080/api',
  enableDebugTools: true
};

// src/environments/environment.prod.ts
export const environment = {
  production: true,
  apiUrl: 'https://api.myapp.com/v1',
  enableDebugTools: false
};

// Use in service — automatically swapped by Angular CLI
import { environment } from '../environments/environment';

@Injectable()
export class ApiService {
  private baseUrl = environment.apiUrl;  // correct URL per environment
}
```

```json
// angular.json — configure environment replacement
"configurations": {
  "production": {
    "fileReplacements": [
      { "replace": "src/environments/environment.ts",
        "with": "src/environments/environment.prod.ts" }
    ]
  }
}
```

---

### Q6. What are common Angular performance best practices?

```typescript
// 1. Use OnPush change detection for heavy components
@Component({ changeDetection: ChangeDetectionStrategy.OnPush })

// 2. Use trackBy with *ngFor to avoid full list re-renders
<li *ngFor="let item of items; trackBy: trackById">
trackById = (_: number, item: Item) => item.id;

// 3. Lazy load routes
{ path: 'admin', loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule) }

// 4. Use async pipe instead of manual subscribe (auto-unsubscribes)
{{ data$ | async }}

// 5. Unsubscribe with takeUntil pattern
private destroy$ = new Subject<void>();

ngOnInit(): void {
  this.dataService.getStream().pipe(
    takeUntil(this.destroy$)
  ).subscribe(data => this.data = data);
}

ngOnDestroy(): void {
  this.destroy$.next();
  this.destroy$.complete();
}

// 6. Use pure pipes instead of method calls in templates
// ❌ Bad — called every change detection cycle
{{ formatName(user.firstName, user.lastName) }}

// ✅ Good — only called when inputs change
{{ user | fullName }}
```
