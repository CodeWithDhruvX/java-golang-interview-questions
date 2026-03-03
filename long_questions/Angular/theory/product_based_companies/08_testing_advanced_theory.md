# 🗣️ Theory — Advanced Testing
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is fakeAsync and tick? How do they let you test debounced or delayed logic?"

> *"fakeAsync wraps a test in a synthetic async zone where time is simulated rather than real. Inside a fakeAsync block, async operations like setTimeout, setInterval, and Promise resolutions don't actually wait — they're queued. tick(300) advances the simulated clock by 300 milliseconds, flushing any timers or intervals that were scheduled to run within that time. This is transformative for testing debounced search inputs or polling intervals. Without fakeAsync, testing a 300ms debounce would require actually waiting 300ms in the test suite — slow and fragile. With fakeAsync and tick(300), the test runs instantly and synchronously. flushMicrotasks() is used to drain Promise chains — you call it after triggering an HTTP mock resolution or any Promise-based async to advance through the microtask queue."*

---

## Q: "How do you test NgRx reducers, selectors, and effects correctly?"

> *"Each piece of NgRx state has its ideal testing approach. Reducers are pure functions — the easiest to test. You call the reducer function directly with a state and an action, and assert on the returned state. No TestBed, no setup. Selectors with createSelector expose a projector function — you call selector.projector() directly with mock input values and assert on the output, bypassing the entire store. For effects, you configure TestBed with the effect under test, use provideMockActions to provide an Observable you control as the actions stream, mock the services the effect calls with jasmine spies, and subscribe to the effect's Observable to assert what actions it dispatches. For testing the integrated NgRx flow — the complete Store, actions, reducers, effects — use StoreModule.forRoot() in TestBed and dispatch actions, then assert on select() emissions."*

---

## Q: "What is the Angular component test harness and why is it better than raw DOM queries?"

> *"Component harnesses are an API provided by Angular CDK that gives you an abstraction layer over DOM for testing. Instead of querying '.mat-input-element' — a selector tied to Angular Material's internal implementation — you load a MatInputHarness and interact with it through a stable typed API like harness.getValue() or harness.setValue(). The fundamental advantage is that harness tests don't break when the library's internal DOM structure changes in a major version upgrade. The library owner guarantees the harness API is stable even if the underlying HTML completely changes. You get methods that match semantic intent — click, focus, selectOption — rather than DOM manipulation through native element references. Angular Material ships harnesses for all its components, and you can write harnesses for your own component library to give consumers the same resilience."*

---

## Q: "How do you test a guard that uses inject() in Angular 16+?"

> *"Modern Angular guards and resolvers are often standalone functions that call inject() internally rather than classes implementing an interface. You can't just call the function with new arguments — it needs to run inside an injection context. TestBed.runInInjectionContext() solves this: you configure TestBed with the providers you need — mocked services, the Router — and call TestBed.runInInjectionContext(() => myGuard()) . Inside the arrow function, inject() resolves from the TestBed's injector. You then assert on the return value: true for allow, a UrlTree for redirect, or false for block. This is the canonical pattern for testing functional guards, resolvers, and any function that relies on inject(). It's simpler than the old approach of instantiating a class that implements CanActivate, because there's no class to construct."*

---

## Q: "How do you test components that use Signals?"

> *"Signal-based components are generally easier to test than Observable-based ones because signals are synchronous — no async pipe timing, no fakeAsync needed for signal updates. You create the component with TestBed, read signal values directly by calling them as functions — expect(component.count()).toBe(0). For signal @Input(), which are set with ComponentFixture's setInput() method, you call fixture.componentRef.setInput('name', value) and the signal updates synchronously — call fixture.detectChanges() and the DOM reflects the new state. For computed() signals, you verify them by setting their dependencies and asserting the derived value. For effects(), since they run asynchronously in a microtask, you use fixture.detectChanges() which flushes pending effects in the testing environment. Overall the synchronous nature of signals makes assertions much more predictable than timing-sensitive Observable tests."*

---

## Q: "What is Spectator and how does it reduce Angular test boilerplate?"

> *"TestBed is powerful but verbose — configuring testing modules, creating fixtures, querying debug elements, mocking services — all of this adds a lot of boilerplate that distracts from what you're actually testing. Spectator is a library built on top of TestBed that provides a concise, expressive API. With createComponent factory, you declare your mocks and providers once in a factory config, and then get a spectator instance with helpers like spectator.query('.my-class'), spectator.click('button'), spectator.inject(MyService). Services are auto-mocked using createSpyObject and accessible via spectator.inject(). For HTTP services, createHttpFactory gives you an http spectator with expectOne and flush. The tests themselves read almost like English — spectator.click('#submit') and expect(spectator.query('.success')).toExist(). Spectator is widely used in production Angular codebases precisely because it makes test suites maintainable at scale."*
