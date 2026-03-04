# 📘 03 — State Management & NgRx Advanced
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- NgRx Effects and side-effect management
- `@ngrx/entity` for normalized state
- Selector memoization and `createSelector`
- NgRx ComponentStore (local state)
- Facade pattern with NgRx
- Action creators and `createActionGroup`

---

## ❓ Most Asked Questions

### Q1. How do NgRx Effects work? Explain with an example.

**Effects** handle **side effects** — actions that need async work (HTTP calls, localStorage, WebSockets). Effects listen to the `Actions` stream, perform the side effect, and dispatch new actions.

```typescript
// products.effects.ts
@Injectable()
export class ProductEffects {
  // loadProducts$ listens for LoadProducts action, calls API, dispatches result
  loadProducts$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.loadProducts),
      switchMap(({ category, page }) =>
        this.productService.getProducts({ category, page }).pipe(
          map(response => ProductActions.loadProductsSuccess({
            products: response.data,
            total: response.total
          })),
          catchError(error => of(ProductActions.loadProductsFailure({
            error: error.message
          })))
        )
      )
    )
  );

  // deleteProduct$ — shows toast on success
  deleteProduct$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.deleteProduct),
      switchMap(({ id }) =>
        this.productService.delete(id).pipe(
          map(() => ProductActions.deleteProductSuccess({ id })),
          catchError(err => of(ProductActions.deleteProductFailure({ error: err.message })))
        )
      )
    )
  );

  deleteProductSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.deleteProductSuccess),
      tap(() => this.toastService.success('Product deleted successfully'))
    ),
    { dispatch: false }  // ← this effect doesn't dispatch a new action
  );

  constructor(
    private actions$: Actions,
    private productService: ProductService,
    private toastService: ToastService
  ) {}
}
```

---

### Q2. How does `@ngrx/entity` simplify collection state?

```typescript
// Without entity: managing a list manually is painful
interface ProductState {
  ids: string[];
  entities: { [id: string]: Product };
  loading: boolean;
}

// With @ngrx/entity: EntityState + EntityAdapter do the heavy lifting
import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';

export interface ProductState extends EntityState<Product> {
  // EntityState provides: ids: string[] and entities: Dictionary<Product>
  loading: boolean;
  selectedProductId: string | null;
  error: string | null;
}

export const productAdapter: EntityAdapter<Product> = createEntityAdapter<Product>({
  selectId: (product) => product.id,
  sortComparer: (a, b) => a.name.localeCompare(b.name)  // optional sort
});

const initialState: ProductState = productAdapter.getInitialState({
  loading: false,
  selectedProductId: null,
  error: null
});

// Reducer — use adapter methods instead of manual array manipulation
export const productReducer = createReducer(
  initialState,
  on(ProductActions.loadProductsSuccess, (state, { products }) =>
    productAdapter.setAll(products, { ...state, loading: false })
  ),
  on(ProductActions.addProduct, (state, { product }) =>
    productAdapter.addOne(product, state)
  ),
  on(ProductActions.updateProduct, (state, { update }) =>
    productAdapter.updateOne(update, state)  // update = { id, changes: Partial<Product> }
  ),
  on(ProductActions.deleteProduct, (state, { id }) =>
    productAdapter.removeOne(id, state)
  ),
  on(ProductActions.loadProducts, state => ({ ...state, loading: true }))
);

// Selectors — entity adapter provides built-in selectors
const { selectAll, selectEntities, selectIds, selectTotal } =
  productAdapter.getSelectors();

export const selectProductState = createFeatureSelector<ProductState>('products');
export const selectAllProducts = createSelector(selectProductState, selectAll);
export const selectProductEntities = createSelector(selectProductState, selectEntities);
export const selectProductById = (id: string) =>
  createSelector(selectProductEntities, entities => entities[id]);
```

---

### Q3. How does NgRx selector memoization work?

```typescript
// createSelector uses memoization — result is cached until inputs change
export const selectAllProducts = createSelector(selectProductState, state => state.ids
  .map(id => state.entities[id]!));

// Derived selector — only recalculates if selectAllProducts changes
export const selectExpensiveProducts = createSelector(
  selectAllProducts,
  (products) => products.filter(p => p.price > 5000)
  // ✅ This filter only runs if selectAllProducts emits a new value
);

// Parameterized selector with props
export const selectProductById = (id: string) =>
  createSelector(selectProductEntities, (entities) => entities[id]);

// Multiple inputs
export const selectFilteredProducts = createSelector(
  selectAllProducts,
  selectCurrentCategory,
  selectSortOrder,
  (products, category, sortOrder) => {
    // This expensive computation only runs when products, category, OR sort changes
    const filtered = category ? products.filter(p => p.category === category) : products;
    return sortOrder === 'asc'
      ? filtered.sort((a, b) => a.price - b.price)
      : filtered.sort((a, b) => b.price - a.price);
  }
);
```

---

### Q4. What is `ComponentStore`? How does it differ from the global NgRx Store?

```typescript
// ComponentStore — local state management scoped to one component and its children
// Global Store — app-wide state, shared across all components

import { ComponentStore } from '@ngrx/component-store';

interface PaginationState {
  page: number;
  pageSize: number;
  sortField: string;
  sortDir: 'asc' | 'desc';
}

@Injectable()
export class ProductTableStore extends ComponentStore<PaginationState> {
  // Initial state
  constructor() {
    super({ page: 1, pageSize: 10, sortField: 'name', sortDir: 'asc' });
  }

  // Selectors
  readonly page$ = this.select(s => s.page);
  readonly pageSize$ = this.select(s => s.pageSize);
  readonly queryParams$ = this.select(
    this.page$, this.pageSize$, (page, pageSize) => ({ page, pageSize })
  );

  // Updaters (sync state mutations)
  readonly setPage = this.updater((state, page: number) => ({ ...state, page }));
  readonly setPageSize = this.updater((state, pageSize: number) =>
    ({ ...state, pageSize, page: 1 })
  );
  readonly setSort = this.updater((state, { field, dir }: { field: string, dir: 'asc'|'desc' }) =>
    ({ ...state, sortField: field, sortDir: dir, page: 1 })
  );

  // Effects (async operations)
  readonly loadProducts = this.effect((trigger$: Observable<void>) =>
    trigger$.pipe(
      switchMap(() => this.queryParams$),
      switchMap(params => this.productService.getProducts(params).pipe(
        tapResponse(
          products => this.patchState({ products }),
          err => console.error(err)
        )
      ))
    )
  );
}

// Component using ComponentStore
@Component({
  providers: [ProductTableStore],  // ← scoped to this component
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProductTableComponent implements OnInit {
  page$ = this.store.page$;
  products$ = /* ... */;

  constructor(private store: ProductTableStore) {}

  ngOnInit(): void { this.store.loadProducts(); }
  changePage(page: number): void { this.store.setPage(page); }
}
```

---

### Q5. What is the Facade pattern in NgRx?

```typescript
// The Facade abstracts NgRx internals from components
// Components don't dispatch actions or use Store directly

@Injectable({ providedIn: 'root' })
export class ProductFacade {
  // Public selectors as observables
  products$ = this.store.select(selectAllProducts);
  loading$ = this.store.select(selectProductsLoading);
  error$ = this.store.select(selectProductsError);
  total$ = this.store.select(selectProductsTotal);

  constructor(private store: Store) {}

  // Public methods instead of action dispatching
  loadProducts(category?: string): void {
    this.store.dispatch(ProductActions.loadProducts({ category }));
  }

  createProduct(product: CreateProductDto): void {
    this.store.dispatch(ProductActions.createProduct({ product }));
  }

  deleteProduct(id: string): void {
    this.store.dispatch(ProductActions.deleteProduct({ id }));
  }

  selectProduct(id: string): Observable<Product | undefined> {
    return this.store.select(selectProductById(id));
  }
}

// Component — no NgRx imports needed!
@Component({ changeDetection: ChangeDetectionStrategy.OnPush })
export class ProductListComponent implements OnInit {
  products$ = this.facade.products$;
  loading$ = this.facade.loading$;

  constructor(private facade: ProductFacade) {}

  ngOnInit(): void { this.facade.loadProducts(); }
  deleteProduct(id: string): void { this.facade.deleteProduct(id); }
}
```

> **Benefits:** Components are decoupled from NgRx; switching to a different state management library only requires updating the Facade, not every component.
