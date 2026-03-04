# 🗣️ Theory — Advanced RxJS
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is the critical difference between switchMap, mergeMap, concatMap, and exhaustMap?"

> *"This is the question that separates Angular developers who know RxJS from those who just use it. All four take each emission and map it to an inner Observable, but their subscription strategies differ fundamentally. switchMap cancels the previous inner Observable on a new emission — the previous in-flight request is abandoned and you get the result from the latest one only. Perfect for search and autocomplete where you only care about the latest result. mergeMap keeps all inner Observables running concurrently — every emission spawns a new subscription and all run in parallel. Use it for independent parallel operations like loading multiple resources simultaneously. concatMap queues emissions — it subscribes to the next inner Observable only after the current one completes, preserving order. Use this for sequential processing where order matters. exhaustMap ignores new emissions while the current inner Observable is active — perfect for form submissions where pressing submit twice should only fire once."*

---

## Q: "What are the different Subject types in RxJS? When do you use each?"

> *"There are four Subject types and knowing when to use each is crucial. A plain Subject has no memory — subscribers only get values emitted after they subscribe, nothing before. Use it as an event bus or to bridge imperative code to Observable streams. BehaviorSubject stores the current value and immediately emits it to new subscribers. This is the workhorse of Angular state management — any component that subscribes gets the current state immediately, without waiting. ReplaySubject lets you specify a buffer size — it stores the last N emissions and replays them to new subscribers. Use it for caching a sequence of recent messages or events. AsyncSubject only emits the last value, and only when it completes — it's rarely used directly but it's the model behind async/await converted to Observables. In practice, BehaviorSubject is what you reach for in 90% of state-sharing scenarios."*

---

## Q: "What is the difference between combineLatest, withLatestFrom, and zip?"

> *"They all combine multiple Observables but in fundamentally different ways. combineLatest emits whenever any source emits, combining the latest value from all sources. It only emits after all sources have emitted at least once. This is ideal for multiple filter dropdowns — every time category, sort, or page changes you get a combined emission and trigger a new search. withLatestFrom is asymmetric — it only emits when the primary Observable emits, using the secondary Observable only as a context provider. Think of it as: 'when the user clicks save, take the current form value.' The click is the trigger; the form is the context. zip pairs emissions by index — first emission from A with first emission from B, second with second. It's strict pairing, useful when you need to coordinate two sequential streams, but uncommon in practice."*

---

## Q: "What is shareReplay and why is it critical for HTTP calls in services?"

> *"HTTP Observables are cold — each subscription triggers a new HTTP request. If your component has two {{ config$ | async }} bindings, that's two HTTP calls. shareReplay solves this by multicasting — it makes the Observable hot, sharing one execution among all subscribers, and replaying the last N emitted values to new subscribers. With shareReplay(1) on an HTTP Observable, the first subscriber triggers the request; the second subscriber gets the cached response immediately without making another request. This is essential for app-wide config loading, user profile fetching, or any data that should be loaded once and shared. The key parameter is the buffer size — pass 1 to share just the most recent result. Without shareReplay, shared service Observables would innocuously multiply your HTTP traffic."*

---

## Q: "How do you create a custom RxJS operator?"

> *"A custom operator is just a function that takes an Observable and returns a new Observable — it's a function from Observable to Observable. You write it as a function that returns another function — the inner function receives the source Observable and returns source.pipe() with your logic. This allows it to be composable in a pipe chain. A practical example is retryWithBackoff — it wraps retryWhen with a scan that tracks retry count, and a delayWhen that doubles the delay on each attempt — one second, two, four, eight — up to a maximum. Another common one is withLoading — it sets a loading BehaviorSubject to true when the Observable is subscribed to, and back to false in the finalize operator which runs on both completion and error. Custom operators dramatically clean up repeated pipe patterns across your codebase."*

---

## Q: "What are the pitfalls of forkJoin? When does it fail silently?"

> *"forkJoin waits for all source Observables to complete and emits an array or object of their last values. The pitfalls: first, if any source errors, forkJoin errors immediately and other sources are cancelled — you lose all results even if only one failed. Fix this by piping catchError on each inner Observable to return a fallback. Second, if any source emits no values and completes — like an Observable that filters everything out — forkJoin emits undefined for that slot, which is surprising. Third, if any source never completes — like a BehaviorSubject or Subject without explicit completion — forkJoin hangs forever. This last one is a common bug: developers try to forkJoin two BehaviorSubjects expecting them to act like one-shot Observables, and the fork never resolves. Always use forkJoin with Observables that are guaranteed to complete, like HTTP calls."*
