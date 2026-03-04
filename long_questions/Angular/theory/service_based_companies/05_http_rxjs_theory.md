# 🗣️ Theory — HTTP & RxJS
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does HttpClient work in Angular? What's the right way to use it?"

> *"Angular's HttpClient is a typed wrapper around the browser's HTTP APIs. You inject it into a service and call methods like get, post, put, delete, and patch — each returns an Observable, not a Promise. The Observable is cold, meaning nothing happens until something subscribes to it. That means calling http.get('/api/products') in a service doesn't actually fire the request — it fires when the component subscribes. The right pattern is to encapsulate all HTTP calls in a service layer, expose Observables, and let components subscribe via the async pipe in the template. This keeps HTTP logic out of components and makes testing straightforward — you mock the HttpClient in tests using HttpClientTestingModule."*

---

## Q: "What is an HttpInterceptor? How do you add an auth token to every request?"

> *"An HttpInterceptor sits in the HTTP pipeline and can inspect and transform requests and responses. You implement the HttpInterceptor interface, which has one method: intercept, which receives the request and a handler. You clone the request — requests are immutable — add your changes to the clone, and pass it to handler.handle(). Cloning with setHeaders lets you add an Authorization Bearer token from your AuthService. You can also use interceptors to add global error handling — catching 401 responses and redirecting to login, or catching 500 errors and showing a toast. You register interceptors in AppModule's providers array with the HTTP_INTERCEPTORS multi-token. You can chain multiple interceptors — they run in registration order."*

---

## Q: "What is an Observable? How is it different from a Promise?"

> *"A Promise represents a single future value — it's eager, meaning it starts executing immediately, it can resolve exactly once, and it can't be cancelled. An Observable represents a stream of zero or more values over time — it's lazy, nothing happens until you subscribe, it can emit multiple values, and the subscription can be cancelled by unsubscribing. For HTTP calls this distinction is subtle — an HTTP Observable also emits exactly one value and completes. But Observables shine for recurring data: form value changes, router events, WebSocket streams. They also have a composable operator pipeline — you chain map, filter, switchMap — which allows complex async logic to be expressed declaratively. In Angular, RxJS Observables are deeply integrated — HttpClient, EventEmitter, ActivatedRoute, FormControl.valueChanges — all are Observable-based."*

---

## Q: "Explain switchMap, mergeMap, and concatMap. When do you use each?"

> *"They're all higher-order mapping operators — they take each emission from a source Observable and map it to an inner Observable. The difference is in how they manage concurrent inner Observables. switchMap cancels the previous inner Observable when a new emission arrives from the source — perfect for search, where you only care about the latest result and want to discard in-flight requests from stale keystrokes. mergeMap runs all inner Observables concurrently — it subscribes to every inner Observable without cancelling old ones, so results arrive in any order — good for independent parallel requests. concatMap waits for each inner Observable to complete before subscribing to the next — it guarantees order — good for sequential operations like processing a queue of items one at a time. exhaustMap ignores new source emissions while an inner Observable is still active — good for preventing duplicate form submissions."*

---

## Q: "What is the async pipe and why is it preferred over manual subscribe?"

> *"The async pipe subscribes to an Observable or Promise, returns the latest emitted value for the template to display, and automatically unsubscribes when the component is destroyed. Compared to manually subscribing in ngOnInit, the async pipe eliminates two major problems. Memory leaks — if you manually subscribe and forget to unsubscribe in ngOnDestroy, the subscription stays alive after the component is gone. Boilerplate — you don't need a subscription variable, an ngOnDestroy, or any cleanup code. The async pipe also integrates with OnPush change detection — when the Observable emits, Angular knows to check that component even if it's in OnPush mode. The pattern 'subscribe yourself in TypeScript' is only justified when you need to do something with the data that can't be expressed in the template."*

---

## Q: "How do you handle HTTP errors globally in Angular?"

> *"The right approach is an HttpInterceptor. You create an ErrorInterceptor that wraps handler.handle(request) in a pipe with catchError. Inside catchError, you inspect the HttpErrorResponse status code — 401 means unauthorized, trigger logout and redirect to login; 403 means forbidden, show an access denied message; 500 means server error, show a generic error toast. You call throwError to rethrow the error so individual components can still handle it if they want to override the default behavior. This way you have one central place for error handling that applies to every HTTP call in the app, rather than duplicating catchError logic in every service or component."*

---

## Q: "What is debounceTime and why is it important for search inputs?"

> *"debounceTime is an RxJS operator that delays emissions from a source Observable until a specified period of silence has passed. Without it, a search feature bound to an input's valueChanges would fire an API call on every single keystroke — if the user types 'angular', you'd make seven API calls. With debounceTime(300), Angular waits 300 milliseconds after the user stops typing before emitting the value — typically reducing seven calls to one. You pair it with distinctUntilChanged, which prevents re-emitting if the value hasn't actually changed — so if the user types 'a', deletes it, and types 'a' again quickly, you still only get one emission. This combination is the standard pattern for efficient search inputs."*
