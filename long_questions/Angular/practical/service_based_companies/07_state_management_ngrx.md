# 📘 07 — State Management & NgRx
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Component state vs shared service state
- BehaviorSubject-based service state (most common in service companies)
- NgRx basics: Store, Action, Reducer, Selector
- When to use NgRx vs simpler patterns

---

## ❓ Most Asked Questions

### Q1. What are the different ways to manage state in Angular?

| Approach | When to Use | Complexity |
|----------|------------|------------|
| Component state | Local UI state (form, toggle, counter) | Low |
| `@Input`/`@Output` | Parent-child communication | Low |
| Service with `BehaviorSubject` | Shared state across sibling/distant components | Medium |
| NgRx Store | Large apps with complex state, time-travel debugging | High |
| Signal-based state (Angular 16+) | Modern reactivity, simple sharing | Low-Medium |

---

### Q2. How do you share state between components using a service?

This is the most common pattern in service-based company projects:

```typescript
// cart.service.ts — BehaviorSubject-based state
@Injectable({ providedIn: 'root' })
export class CartService {
  // BehaviorSubject holds current state and replays it to new subscribers
  private cartItemsSubject = new BehaviorSubject<CartItem[]>([]);
  private loadingSubject = new BehaviorSubject<boolean>(false);

  // Expose as read-only Observables (components can't directly push)
  cartItems$ = this.cartItemsSubject.asObservable();
  loading$ = this.loadingSubject.asObservable();

  // Derived state using combineLatest / map
  cartCount$ = this.cartItems$.pipe(
    map(items => items.reduce((sum, i) => sum + i.quantity, 0))
  );
  cartTotal$ = this.cartItems$.pipe(
    map(items => items.reduce((sum, i) => sum + i.price * i.quantity, 0))
  );

  addToCart(product: Product, qty: number = 1): void {
    const current = this.cartItemsSubject.getValue();
    const existing = current.find(i => i.productId === product.id);

    const updated = existing
      ? current.map(i => i.productId === product.id
          ? { ...i, quantity: i.quantity + qty }
          : i)
      : [...current, { productId: product.id, name: product.name, price: product.price, quantity: qty }];

    this.cartItemsSubject.next(updated);  // emit new state
  }

  removeFromCart(productId: string): void {
    const updated = this.cartItemsSubject.getValue()
      .filter(i => i.productId !== productId);
    this.cartItemsSubject.next(updated);
  }

  clearCart(): void {
    this.cartItemsSubject.next([]);
  }
}
```

```typescript
// navbar.component.ts — reads cart count
@Component({
  template: `<span class="badge">{{ cartCount$ | async }}</span>`,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class NavbarComponent {
  cartCount$ = this.cartService.cartCount$;
  constructor(public cartService: CartService) {}
}

// product-detail.component.ts — writes to cart
@Component({ template: `<button (click)="addToCart()">Add to Cart</button>` })
export class ProductDetailComponent {
  @Input() product: Product;
  constructor(private cartService: CartService) {}

  addToCart(): void {
    this.cartService.addToCart(this.product);
  }
}
```

---

### Q3. What is NgRx and what is the Redux pattern?

**NgRx** is a state management library for Angular based on the **Redux pattern**:

```
Component → dispatches Action → Reducer → updates Store (State) → Selectors → Component (via async pipe)
                                                ↑
                                          Effects (for side effects like HTTP)
```

| Concept | Role |
|---------|------|
| **Store** | Single source of truth — immutable state object |
| **Action** | Event describing what happened (`[Cart] Add Item`) |
| **Reducer** | Pure function that calculates new state from action |
| **Effect** | Handles side effects (API calls) triggered by actions |
| **Selector** | Extracts derived/sliced state from the store |

---

### Q4. Show a complete NgRx example for a cart.

```typescript
// 1. STATE (defined alongside reducer)
export interface CartState {
  items: CartItem[];
  loading: boolean;
  error: string | null;
}
const initialState: CartState = { items: [], loading: false, error: null };

// 2. ACTIONS
export const CartActions = createActionGroup({
  source: 'Cart',
  events: {
    'Add Item': props<{ product: Product; quantity: number }>(),
    'Remove Item': props<{ productId: string }>(),
    'Load Cart': emptyProps(),
    'Load Cart Success': props<{ items: CartItem[] }>(),
    'Load Cart Failure': props<{ error: string }>(),
    'Clear Cart': emptyProps(),
  }
});

// 3. REDUCER
export const cartReducer = createReducer(
  initialState,
  on(CartActions.addItem, (state, { product, quantity }) => ({
    ...state,
    items: state.items.some(i => i.productId === product.id)
      ? state.items.map(i => i.productId === product.id
          ? { ...i, quantity: i.quantity + quantity } : i)
      : [...state.items, { productId: product.id, ...product, quantity }]
  })),
  on(CartActions.removeItem, (state, { productId }) => ({
    ...state,
    items: state.items.filter(i => i.productId !== productId)
  })),
  on(CartActions.clearCart, state => ({ ...state, items: [] })),
  on(CartActions.loadCart, state => ({ ...state, loading: true })),
  on(CartActions.loadCartSuccess, (state, { items }) => ({ ...state, items, loading: false })),
  on(CartActions.loadCartFailure, (state, { error }) => ({ ...state, error, loading: false })),
);

// 4. SELECTORS
export const selectCartState = createFeatureSelector<CartState>('cart');
export const selectCartItems = createSelector(selectCartState, s => s.items);
export const selectCartCount = createSelector(selectCartItems,
  items => items.reduce((n, i) => n + i.quantity, 0));
export const selectCartTotal = createSelector(selectCartItems,
  items => items.reduce((t, i) => t + i.price * i.quantity, 0));

// 5. EFFECTS
@Injectable()
export class CartEffects {
  loadCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.loadCart),
      switchMap(() => this.cartService.fetchCartFromServer().pipe(
        map(items => CartActions.loadCartSuccess({ items })),
        catchError(err => of(CartActions.loadCartFailure({ error: err.message })))
      ))
    )
  );

  constructor(private actions$: Actions, private cartService: CartService) {}
}
```

```typescript
// 6. COMPONENT — dispatches actions, selects state
@Component({ changeDetection: ChangeDetectionStrategy.OnPush })
export class CartComponent {
  cartItems$ = this.store.select(selectCartItems);
  cartCount$ = this.store.select(selectCartCount);
  cartTotal$ = this.store.select(selectCartTotal);

  constructor(private store: Store) {
    this.store.dispatch(CartActions.loadCart());
  }

  removeItem(productId: string): void {
    this.store.dispatch(CartActions.removeItem({ productId }));
  }
}
```

---

### Q5. When should you use NgRx vs a service with BehaviorSubject?

**Use a service with BehaviorSubject when:**
- App is small to medium sized
- State is simple (cart, auth status, theme)
- Team is not familiar with NgRx
- Quick development is priority

**Use NgRx when:**
- Large team, many features sharing state
- Complex async flows (multiple chained HTTP calls)
- Need time-travel debugging (Redux DevTools)
- State is shared across 10+ components
- Strict separation of concerns is required

> **Service company interview answer:** "For most service company projects, a BehaviorSubject-based service is sufficient. I'd introduce NgRx only when the state becomes too complex to manage in services."
