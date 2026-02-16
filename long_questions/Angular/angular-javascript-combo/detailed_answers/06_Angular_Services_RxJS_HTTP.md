# 💉 Angular Services, RxJS & HTTP (Detailed Answers)

## 25. What is Dependency Injection (DI)?
A design pattern where a class requests dependencies from external sources rather than creating them itself.
Angular has its own **Dependency Injection (DI)** framework.

**Benefits**: Reusability, Testability (mocking deps), Maintainability.

## 26. What is @Injectable?
A decorator that marks a class as available to the injector for creation.
Without `@Injectable`, Angular's DI system throws an error when trying to inject it.

## 27. What is providedIn: 'root'?
It registers the service as a **singleton** in the root injector.
**Benefit**: Tree-shaking. If the service is not used, it is removed from the final bundle.

```typescript
@Injectable({
  providedIn: 'root'
})
export class DataService { ... }
```

## 28. How do you share data between components?
1.  **Input/Output**: Parent -> Child (`@Input`), Child -> Parent (`@Output`).
2.  **Service (Singleton)**: Unrelated components subscribe to the same `Subject`/`BehaviorSubject`.
3.  **Route Params**: For data passing during navigation (`ActivatedRoute`).
4.  **State Management**: NgRx/Akita/NgXS (complex/large applications).

---

## 29. What is Observable? (RxJS)
An Observable represents a stream of data that can arrive synchronously or asynchronously over time.
**Metaphor**: Like a YouTube channel. You subscribe to get notifications about new videos.

**Creation**:
```typescript
const obs$ = new Observable(subscriber => {
  subscriber.next(1);
  subscriber.next(2);
  subscriber.complete();
});
obs$.subscribe(val => console.log(val)); // 1, 2
```

## 30. Difference between Promise and Observable?

| Feature | Promise | Observable |
| :--- | :--- | :--- |
| **Values** | Single value (resolve/reject) | Multiple values over time. |
| **Execution** | Eager (executes immediately) | Lazy (executes only on subscribe). |
| **Cancellation** | Not cancellable (standard). | Cancellable (`unsubscribe`). |
| **Operators** | `.then`, `.catch` | powerful operators (`map`, `filter`, `retry`, `debounce`). |

## 31. What is Subscription?
Represents the execution of an Observable.
Created by calling `.subscribe()`.
**Important**: Must be stored in a variable to call `.unsubscribe()` later to prevent memory leaks.

```typescript
this.sub = this.api.getData().subscribe();
ngOnDestroy() {
  this.sub.unsubscribe();
}
```

## 32. What is pipe?
A method on Observables used to chain multiple operators together.
```typescript
obs$.pipe(
  filter(x => x > 10),
  map(x => x * 2)
).subscribe();
```

## 33. What is Subject?
A special type of Observable that allows values to be **multicasted** to many Observers.
While plain Observables are unicast (each subscribed Observer owns an independent execution of the Observable), Subjects are multicast.

## 34. What is BehaviorSubject?
A variant of Subject that requires an **initial value** and emits the **current value** to new subscribers.
Perfect for storing state (e.g., `currentUser`, `isLoggedIn`).

```typescript
const subject = new BehaviorSubject(0); // 0 is initial value
subject.subscribe(v => console.log('A: ' + v)); // A: 0
subject.next(1); // A: 1
subject.subscribe(v => console.log('B: ' + v)); // B: 1 (gets last value immediately)
```

## 35. What is ReplaySubject?
Similar to BehaviorSubject, but it can record multiple values from the Observable execution and replay them to new subscribers.

---

## 36. Operators: mergeMap vs switchMap?

| Operator | Behavior | Use Case |
| :--- | :--- | :--- |
| **mergeMap** | Maps to inner Observable, subscribes to it, and emits its values. Handles concurrency (parallel). | Independent requests (e.g., Delete multiple items). |
| **switchMap** | Maps to inner Observable, cancels previous inner Observable if new value arrives. | Scenarios involving user input (Search, filtering), where only latest result matters. |
| **concatMap** | Maps to inner Observable, waits for previous to complete before subscribing to next. (Sequential). | Order-sensitive operations (e.g., Update sequential steps). |
| **exhaustMap** | Ignores new values while current inner Observable is running. | Preventing double-clicks (e.g., Login button). |

---

## 37. What is HttpClient?
A modern, Promise-based implementation for HTTP requests (GET, POST, PUT, DELETE).
Returns **Observables**.

## 38. How do you handle errors in HttpClient?
Use the `catchError` operator in the pipe.

```typescript
this.http.get(url).pipe(
  catchError((error) => {
    console.error(error);
    return throwError(error); // Re-throw or return default value
  })
).subscribe();
```

## 39. What is an Interceptor?
Middleware that inspects and transforms HTTP requests/responses globally.
**Use Case**:
1.  Add Auth Token to headers.
2.  Global Error Handling (Toastr notifications).
3.  Loader Spinner (Show/Hide).
4.  Logging.

## 40. How do you add token in header using Interceptor?
Implement `HttpInterceptor` interface.

```typescript
@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = localStorage.getItem('token');
    if (token) {
      const cloned = req.clone({
        headers: req.headers.set("Authorization", "Bearer " + token)
      });
      return next.handle(cloned);
    }
    return next.handle(req);
  }
}
```

## 41. How do you cancel previous API calls?
Use `switchMap` operator.
Ideally used in typeahead search. If user types 'a', api called. typing 'b' cancels 'a' call and starts 'ab' call.

```typescript
inputChange$.pipe(
  debounceTime(300),
  distinctUntilChanged(),
  switchMap(term => this.searchService.search(term))
).subscribe();
```
