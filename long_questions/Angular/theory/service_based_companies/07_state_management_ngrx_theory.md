# 🗣️ Theory — State Management & NgRx
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are the different approaches to state management in Angular?"

> *"There's a spectrum from simple to complex. Component state — just properties on the component class — works fine for local UI state like whether a dropdown is open or what tab is selected. @Input and @Output handle parent-child communication. For sharing state between sibling or distant components, a service with a BehaviorSubject is the cleanest lightweight approach — you expose the state as an Observable, components subscribe, and the service is the single source of truth. For large applications with complex async flows and many teams, NgRx brings a Redux-style unidirectional data flow with Store, Actions, Reducers, Effects, and Selectors. The trade-off is boilerplate and learning curve. I always start with the simplest approach and introduce NgRx only when service-based solutions become hard to reason about."*

---

## Q: "How does a BehaviorSubject-based service work for state sharing?"

> *"BehaviorSubject is an RxJS Subject that stores the current value and immediately emits it to any new subscriber. The pattern is: create a BehaviorSubject with an initial state, expose it as a read-only Observable using .asObservable() so external code can't directly push values, and provide methods on the service for state mutations. For example, a CartService would have a private BehaviorSubject holding CartItem[], and expose cartItems$ as an Observable. Components subscribe to cartItems$ via the async pipe. When addToCart is called, you read the current value with getValue(), compute the new state, and push it with .next(). Every subscriber — the navbar showing cart count, the cart page, the mini-cart — automatically receives the update. No NgRx needed."*

---

## Q: "What is NgRx? What problem does it solve that a service can't?"

> *"NgRx is an Angular implementation of the Redux pattern — it gives you a single immutable state tree, strictly unidirectional data flow with Actions dispatched to Reducers, and Effects for side effects. The problem it solves that services struggle with at scale: in a large app with dozens of services holding BehaviorSubjects, understanding how and when state changes is hard — a mutation in one service might trigger three others. NgRx makes all state changes explicit — every mutation is an Action with a name and a payload, logged in Redux DevTools. Time-travel debugging, action replay, and state snapshots become possible. It also enforces a strict architecture that's easier to reason about in large teams. The downside is significant boilerplate. I'd say NgRx is worth it at team sizes of five or more developers sharing significant state."*

---

## Q: "Walk me through the NgRx data flow: Store, Action, Reducer, Selector, Effect."

> *"The flow always starts with a component: it dispatches an Action — a plain object with a type and optional payload, like 'Load Products'. The Action flows to the Reducer — a pure function that takes the current state and the action, and returns a new state object. The Reducer is synchronous and has no side effects. The new state is stored in the Store — a single RxJS Observable of your entire application state. Components and services read state through Selectors — pure functions that derive slices of state, memoized so they only recompute when their inputs change. For async work — like the actual HTTP call for 'Load Products' — that's an Effect. Effects listen to the Actions stream, perform the side effect, and dispatch new Actions — either 'Load Products Success' with the data or 'Load Products Failure' with an error. Reducers handle both and update state accordingly."*

---

## Q: "What is the Facade pattern with NgRx? Why is it useful?"

> *"The Facade pattern hides NgRx behind a service. Instead of components directly dispatching actions and selecting from the store, they talk to a Facade service that encapsulates all NgRx interactions. The Facade exposes Observable properties for the state slices the component needs, and regular methods for the operations the component can trigger. The benefits are significant: components don't have NgRx imports at all — they just use a service like any other. If you decide to swap NgRx for a BehaviorSubject service or Signals-based state in the future, you only change the Facade, not every component. It also makes components much simpler to test — you mock the Facade instead of setting up an entire NgRx store."*

---

## Q: "When should you use NgRx vs a service with BehaviorSubject?"

> *"I use this rule of thumb: service with BehaviorSubject when the state is simple, the team is small, and the feature is relatively self-contained. NgRx when the state is shared across many features, the team has multiple developers touching the same state, or you need strict debugging guarantees — audit logs of state changes, time-travel debugging, reproducible bug reports. If someone asks in an interview at a service-based company, I'd say: most service-company projects use BehaviorSubject services and that's perfectly fine. If I'm at a product company building a large consumer app with a big team, NgRx provides the discipline and tooling that pays off. The worst outcome is adding NgRx boilerplate to a small app where a two-line BehaviorSubject service would suffice."*
