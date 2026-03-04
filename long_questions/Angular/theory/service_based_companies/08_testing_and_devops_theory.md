# 🗣️ Theory — Testing & DevOps
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you unit test Angular components? What does TestBed do?"

> *"TestBed is Angular's testing utility that creates a mini Angular environment for testing. You configure it with configureTestingModule — declaring the component under test, importing any modules it needs, and providing mock services. TestBed then creates the component via createComponent, which gives you a ComponentFixture. The fixture wraps the component instance and its DOM. You call fixture.detectChanges() to trigger Angular's change detection — this is equivalent to the first ngOnInit run. Then you query the DOM using fixture.debugElement.query() with By.css selectors, trigger event handlers with triggerEventHandler, and make assertions on the DOM or component state. The key principle is to mock services — never let a unit test make real HTTP calls."*

---

## Q: "How do you mock a service in an Angular component test?"

> *"The idiomatic approach is jasmine.createSpyObj — you give it the service name and an array of method names, and it creates an object where every method is a spy that records calls and lets you control return values. You provide this mock in TestBed's providers array using { provide: RealService, useValue: mockService }. Then you can use mockService.getProducts.and.returnValue(of([...])) to control what the spy returns. After performing an action, you can assert that mockService.createProduct was called with the right arguments. This pattern completely decouples the component test from any actual service implementation or HTTP calls."*

---

## Q: "How do you test services that make HTTP calls?"

> *"You import HttpClientTestingModule in your test's imports and inject HttpTestingController. After triggering the service method under test, you call httpMock.expectOne() with the expected URL — this asserts that exactly one request was made to that URL and returns a TestRequest object. You then call req.flush() with mock data to simulate the server response. After each test, you call httpMock.verify() to ensure no outstanding requests were left unhandled. This approach tests both the request shape — URL, method, headers — and the response handling — how the service transforms the response — without making real network calls."*

---

## Q: "What is fakeAsync and tick? Why do you need them?"

> *"Angular code often has asynchronous operations — debounced input, setTimeout, HTTP calls. In tests, you usually don't want to actually wait those milliseconds — tests should be fast and deterministic. fakeAsync wraps a test function in a fake asynchronous environment where you control time. tick() advances that virtual clock by however many milliseconds you specify. So if you want to test that a search fires after a 300ms debounce, you call tick(300) and then make your assertion. For completing HTTP observables or Promises, flushMicrotasks() drains pending microtasks. This lets you test time-based async logic synchronously without real delays."*

---

## Q: "What Angular CLI commands are important to know for builds and environments?"

> *"For development you run ng serve, which starts a dev server with live reload — you can add --port to change the port and --open to launch the browser automatically. For production you run ng build --configuration=production, which enables AOT compilation, minification, tree-shaking, and bundle optimization. The key concept is environment files — you have environment.ts for development and environment.prod.ts for production. Angular CLI automatically swaps them based on the build configuration. So in your services you import from environment.ts, and you get the right API URL or feature flags per environment without any runtime logic. You can also define additional configurations like staging in angular.json and run ng build --configuration=staging."*

---

## Q: "What are the main Angular performance best practices you'd mention in an interview?"

> *"I'd cover five main points. First, OnPush change detection on all presentational components — this alone cuts down Angular's change detection work dramatically. Second, trackBy with every *ngFor — prevents Angular from destroying and recreating the entire list DOM when the array changes. Third, lazy loading all feature modules — the initial bundle should only contain what's needed for the first view. Fourth, using the async pipe instead of manual subscription — it eliminates memory leaks automatically. Fifth, avoiding method calls in templates — Angular evaluates them on every change detection cycle. Use pipes or compute derived values in the component class instead. These five points demonstrate that you understand Angular's performance model, not just its API."*
