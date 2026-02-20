# 🔴 State Management & NgRx

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (Infosys, TCS): BehaviorSubject-based state, basic NgRx concepts
> - 🚀 **Product-Based** (Flipkart, Meesho, PhonePe): NgRx effects, selectors, entity adapters, component store
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

### 1. What is NgRx? 🟡 | 🏭🚀

"**NgRx** is an Angular state management library inspired by **Redux**. It provides a predictable, centralized state container using:
- **Store** — A single immutable state object (the source of truth)
- **Actions** — Events that describe state changes
- **Reducers** — Pure functions that compute the new state from action + current state
- **Selectors** — Functions to query/derive data from the store
- **Effects** — Side effects (HTTP calls, localStorage) triggered by actions

```
Component → dispatch(Action) → Reducer → Store → Selector → Component
                                           ↕
                                         Effects → API → dispatch(Action)
```

NgRx is best used when multiple unrelated components share complex state that changes frequently."

#### In Depth
The core philosophy of NgRx is **unidirectional data flow** — state only flows in one direction, making state mutations traceable and debuggable. NgRx DevTools (Redux DevTools) allows **time-travel debugging** — replaying actions to reproduce bugs. However, NgRx has significant boilerplate overhead. For apps with localized state, `BehaviorSubject`-based services or Angular Signals are better choices.

---

### 2. What are actions, reducers, and effects in NgRx? 🟡 | 🏭🚀

"**Actions** — Plain objects that describe what happened:

```typescript
// actions/product.actions.ts
export const loadProducts = createAction('[Products Page] Load Products');
export const loadProductsSuccess = createAction(
  '[Products API] Load Products Success',
  props<{ products: Product[] }>()
);
export const loadProductsFailure = createAction(
  '[Products API] Load Products Failure',
  props<{ error: string }>()
);
```

**Reducers** — Pure functions that update state:

```typescript
// reducers/product.reducer.ts
interface ProductState {
  products: Product[];
  loading: boolean;
  error: string | null;
}

const initialState: ProductState = { products: [], loading: false, error: null };

export const productReducer = createReducer(
  initialState,
  on(loadProducts, state => ({ ...state, loading: true })),
  on(loadProductsSuccess, (state, { products }) => ({ ...state, loading: false, products })),
  on(loadProductsFailure, (state, { error }) => ({ ...state, loading: false, error }))
);
```

**Effects** — Handle side effects:

```typescript
@Injectable()
export class ProductEffects {
  loadProducts$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadProducts),
      switchMap(() =>
        this.productService.getAll().pipe(
          map(products => loadProductsSuccess({ products })),
          catchError(err => of(loadProductsFailure({ error: err.message })))
        )
      )
    )
  );

  constructor(private actions$: Actions, private productService: ProductService) {}
}
```"

#### In Depth
**Effects return Actions** — this is the key NgRx design principle. An effect listens for one action type and dispatches another action with the result. This creates a **reactive chain**: component dispatches → effect handles side effect → effect dispatches result → reducer updates state → selector notifies component. For fire-and-forget effects (like navigation or analytics), use `{ dispatch: false }`.

---

### 3. What are selectors in NgRx? 🟡 | 🏭🚀

"**Selectors** are pure functions that extract and derive data from the NgRx store. They are **memoized** — the same input always returns the same cached output.

```typescript
// selectors/product.selectors.ts
export const selectProductState = createFeatureSelector<ProductState>('products');

export const selectAllProducts = createSelector(
  selectProductState,
  state => state.products
);

export const selectLoading = createSelector(
  selectProductState,
  state => state.loading
);

// Derived selector (computed from other selectors)
export const selectActiveProducts = createSelector(
  selectAllProducts,
  products => products.filter(p => p.isActive)
);

export const selectProductCount = createSelector(
  selectActiveProducts,
  products => products.length
);
```

```typescript
// In component:
this.products$ = this.store.select(selectActiveProducts);
this.isLoading$ = this.store.select(selectLoading);
```"

#### In Depth
Selector memoization means that if `selectAllProducts` is called twice with the same store state, it returns the **same object reference** — no re-computation, and importantly no change detection trigger for `OnPush` components (because the reference is identical). This is a powerful optimization: complex derived data is computed **once per state change**, not once per component subscription.

---

### 4. How to persist state in local storage using NgRx? 🔴 | 🚀

"I use NgRx's `META_REDUCERS` pattern (a meta-reducer wraps each reducer) for cross-cutting concerns like state persistence:

```typescript
// State persistence meta-reducer
function localStorageMetaReducer<S extends object, A extends Action>(
  reducer: ActionReducer<S, A>
): ActionReducer<S, A> {
  return function (state: S | undefined, action: A): S {
    if (action.type === INIT) {
      // Rehydrate from localStorage on app start
      const saved = localStorage.getItem('appState');
      if (saved) {
        return { ...reducer(state, action), ...JSON.parse(saved) };
      }
    }

    const newState = reducer(state, action);
    localStorage.setItem('appState', JSON.stringify(newState));
    return newState;
  };
}

// Register in StoreModule:
providers: [
  { provide: META_REDUCERS, useValue: localStorageMetaReducer, multi: true }
]
```

For production, I use the **ngrx-store-localstorage** package which handles this with serialization options, key whitelisting, and encryption."

#### In Depth
Persisting the ENTIRE store to localStorage is rarely advisable — it can persist sensitive data (user tokens, PII) or stale data that causes bugs after schema changes. I persist **selectively**: only UI preferences, cart contents, or wizard progress. I also version the stored state and clear it if the schema version changes — preventing crashes from stale state structures after deployments.

---

### 5. NgRx vs BehaviorSubject — When to use which? 🔴 | 🚀

"This is a judgment call based on **complexity and team size**:

**Use `BehaviorSubject` + Service when:**
- Small to medium app (< 10 shared state slices)
- State changes are localized to a feature
- Team is not familiar with Redux concepts
- You want minimal boilerplate

```typescript
@Injectable({ providedIn: 'root' })
export class CartService {
  private items$ = new BehaviorSubject<CartItem[]>([]);
  readonly cart$ = this.items$.asObservable();

  addItem(item: CartItem): void {
    this.items$.next([...this.items$.getValue(), item]);
  }
}
```

**Use NgRx when:**
- Large app with 5+ teams sharing global state
- You need time-travel debugging
- Complex action orchestration (optimistic updates, undo/redo)
- State shape is complex with many derived views

I've seen apps where developers defaulted to NgRx for every feature and ended up with 500 files of boilerplate for what could be 50 files with `BehaviorSubject`. Choose the **complexity that fits the problem**."

#### In Depth
The modern answer is increasingly **Angular Signals** for local and shared state. Signals provide reactive state without the Subject API complexity, and with `signal()`, `computed()`, and `effect()`, you get reactive state management that integrates natively with Angular's change detection. For truly global state that benefits from DevTools debugging, NgRx remains the best option. For everything else, Signals are the future.

---

### 6. What are NgRx Entity Adapters? 🔴 | 🚀

"**`@ngrx/entity`** provides utility functions to manage collections of entities (records with an ID) in the NgRx store. It eliminates boilerplate CRUD reducer logic.

```typescript
import { createEntityAdapter, EntityState } from '@ngrx/entity';

export interface Product { id: number; name: string; price: number; }

// 1. Create adapter
const adapter = createEntityAdapter<Product>();

// 2. Define state
interface ProductState extends EntityState<Product> {
  loading: boolean;
}

const initialState = adapter.getInitialState({ loading: false });

// 3. Use adapter operations in reducers
const productReducer = createReducer(
  initialState,
  on(loadProductsSuccess, (state, { products }) =>
    adapter.setAll(products, { ...state, loading: false })
  ),
  on(updateProduct, (state, { product }) =>
    adapter.updateOne({ id: product.id, changes: product }, state)
  ),
  on(deleteProduct, (state, { id }) =>
    adapter.removeOne(id, state)
  )
);

// 4. Built-in selectors
const { selectAll, selectEntities, selectIds, selectTotal } = adapter.getSelectors();
```"

#### In Depth
Entity adapters store data in a **normalized form**: `{ ids: [1, 2, 3], entities: { 1: {...}, 2: {...} } }`. This **O(1) lookup by ID** (instead of O(n) array search) is crucial for large collections. Operations like "update product #42" are `O(1)` instead of `O(n)`. The built-in selectors (`selectAll`, `selectEntities`) provide ergonomic access to both the array form and the dictionary form of the data.

---

### 7. What is `@ngrx/component-store`? 🔴 | 🚀

"**`ComponentStore`** is NgRx's solution for **component-scoped state** — state that's specific to a single component or feature section, and doesn't need to be global.

It's simpler than the global store (no actions, no effects boilerplate) but still reactive:

```typescript
interface MovieState { movies: Movie[]; loading: boolean; }

@Injectable()
export class MovieStore extends ComponentStore<MovieState> {
  constructor(private movieService: MovieService) {
    super({ movies: [], loading: false }); // Initial state
  }

  // Selectors
  readonly movies$ = this.select(state => state.movies);
  readonly loading$ = this.select(state => state.loading);

  // Updaters (synchronous state changes)
  readonly setMovies = this.updater((state, movies: Movie[]) => ({
    ...state, movies, loading: false
  }));

  // Effects (async operations)
  readonly loadMovies = this.effect((trigger$: Observable<void>) =>
    trigger$.pipe(
      tap(() => this.patchState({ loading: true })),
      switchMap(() => this.movieService.getAll().pipe(
        tapResponse(movies => this.setMovies(movies), () => this.patchState({ loading: false }))
      ))
    )
  );
}
```"

#### In Depth
`ComponentStore` is provided at the **component level** (`providers: [MovieStore]`), so it lives and dies with the component — automatic cleanup without manual subscriptions. It bridges the gap between global NgRx store (too heavy for local state) and `BehaviorSubject` services (too low-level for complex local state). I use it for page-level components that manage their own complex async state without sharing it globally.

---
