# 🔥 High-Probability Scenario Questions (Detailed Answers)

## 1. How do you handle large data tables? (10,000+ rows)
**Scenario**: User needs to view a list of users, but there are 50,000 records.
**Problem**: Rendering 50,000 DOM nodes freezes the browser (UI blocking).

**Solution**:
1.  **Server-Side Pagination**:
    *   Fetch only 10-20 items at a time (`?page=1&limit=20`).
    *   Show "Next/Prev" buttons.
2.  **Virtual Scrolling (Windowing)**:
    *   Use `@angular/cdk/scrolling`.
    *   Only renders the visible items in the viewport (e.g., 10 items) plus a small buffer.
    *   As user scrolls, DOM nodes are recycled.

```html
<cdk-virtual-scroll-viewport itemSize="50" class="viewport">
  <div *cdkVirtualFor="let item of items" class="item">{{item}}</div>
</cdk-virtual-scroll-viewport>
```

---

## 2. How do you implement search with debounce?
**Scenario**: Search bar that filters results as user types.
**Problem**: Api call for every keystroke (`k`, `ki`, `kit`, `kitc`...) -> Overloads server.

**Solution**: Use RxJS operators.
1.  **Subject**: Bind input event to a Subject.
2.  **debounceTime(300)**: Wait for 300ms pause in typing.
3.  **distinctUntilChanged()**: Don't search if value is same (e.g., type 'a', backspace, type 'a').
4.  **switchMap()**: Cancel previous pending request if new one starts.

```typescript
search(term: string) {
  this.searchTerms.next(term);
}

ngOnInit() {
  this.users$ = this.searchTerms.pipe(
    debounceTime(300),
    distinctUntilChanged(),
    switchMap((term: string) => this.userService.searchUsers(term)),
  );
}
```

---

## 3. How do you manage state in Angular?
**Scenario**: Sharing user data, cart items across multiple unrelated components.

**Solution**:
1.  **Small/Medium App**: Shared Service with `BehaviorSubject`.
    *   Service holds state in private `_state` variable.
    *   Exposes `state$` as Observable.
    *   Methods `updateState()` to modify.
2.  **Large/Complex App**: **NgRx** (Redux pattern).
    *   **Store**: Single source of truth.
    *   **Actions**: Describe unique events (`[Auth] Login Success`).
    *   **Reducers**: Pure functions treating state immutable inputs.
    *   **Effects**: Handle side effects (API calls).
    *   **Selectors**: Slice of state for components.

---

## 4. How do you handle authentication?
**Scenario**: Secure the app.

**Flow**:
1.  **Login Component**: User enters credentials -> API call (`POST /login`).
2.  **Store Token**: Store JWT token in `localStorage` or `sessionStorage` (or HttpOnly Cookie).
3.  **Interceptor**: Create `AuthInterceptor` to attach `Authorization: Bearer <token>` to every subsequent request.
4.  **Route Guard**: Create `AuthGuard` (CanActivate) to check if token exists. If not, redirect to `/login`.
5.  **Logout**: Clear storage, navigate to `/login`.

---

## 5. How do you protect routes?
**Scenario**: Prevent specialized users (e.g., Admin) from accessing specific routes.

**Solution**:
1.  Implement `CanActivate` interface.
2.  Inject `AuthService` and `Router`.
3.  Check condition (isLoggedIn?).
4.  Return `true` or `UrlTree` (redirect).

```typescript
canActivate(): boolean | UrlTree {
  if (this.authService.isLoggedIn()) {
    return true;
  }
  return this.router.parseUrl('/login');
}
```

---

## 6. How do you implement role-based access?
**Scenario**: Admin can see "Settings", User cannot.

**Solution**:
1.  Store role in `User` object (in State/Storage).
2.  **Route Guard**: Check `user.role === 'admin'`.
3.  **Directive (*appHasRole)**: Create custom structural directive to hide elements in templates.

```html
<li *appHasRole="['admin', 'manager']"><a routerLink="/admin">Admin Panel</a></li>
```

---

## 7. What happens if API fails?
**Scenario**: Backend is down or returns 500 error.

**Solution**:
1.  **Local Handling**: `catchError` in component -> display error message variable.
2.  **Global Handling (Interceptor)**:
    *   Intercept response error.
    *   Show a Toast notification ("Something went wrong").
    *   Log error to logging service (Sentry).
3.  **Retry**: Use `retry(3)` operator for flaky networks.

---

## 8. How do you handle loader globally?
**Scenario**: Show spinner whenever *any* API call is in progress.

**Solution**:
1.  **LoaderService**: `isLoading = new BehaviorSubject(false)`.
2.  **Interceptor**:
    *   `intercept()` starts: `loaderService.show()`.
    *   `finalize()` operator: `loaderService.hide()`.
3.  **AppComponent**: Subscribes to `loaderService.isLoading` -> shows `<app-spinner>` if true.

---

## 9. How do you share data between unrelated components?
**Scenario**: Sibling components or deeply nested components need to talk.

**Solution**: **Shared Service (RxJS)**
1.  Create `DataService`.
2.  Create `private messageSource = new BehaviorSubject('default message');`
3.  Expose `currentMessage = this.messageSource.asObservable();`
4.  Method `changeMessage(message: string) { this.messageSource.next(message) }`
5.  **Sender**: Injects service, calls `changeMessage()`.
6.  **Receiver**: Injects service, subscribes to `currentMessage`.

---

## 10. How do you optimize repeated API calls?
**Scenario**: User clicks "Refresh" button repeatedly.
**Scenario**: Component re-initializes often.

**Solution**:
1.  **ShareReplay**: Cache the response.
    ```typescript
    this.data$ = this.http.get(url).pipe(shareReplay(1));
    ```
2.  **ExhaustMap**: Ignore strict clicks while previous request is pending.
3.  **SwitchMap**: Cancel previous request if new one comes (Search).
