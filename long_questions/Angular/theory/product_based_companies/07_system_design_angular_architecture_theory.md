# 🗣️ Theory — System Design & Angular Architecture
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are Micro Frontends? How does Module Federation enable them in Angular?"

> *"Micro Frontends apply microservices thinking to the frontend — instead of one monolithic Angular application, you have multiple independent Angular applications developed, built, and deployed separately by different teams, and composed together at runtime. Module Federation is Webpack 5's mechanism that makes this possible. A remote application exposes specific modules — an Angular feature module or even a standalone component — in its webpack config. A shell application declares what remotes it knows about and dynamically imports from them at runtime. In Angular, this looks like a lazy route that imports from a remote URL — import('productsApp/ProductsModule'). Angular router handles this exactly like lazy loading a local module. The critical constraints: each MFE must agree on sharing Angular core as a singleton — you can't have two Angular runtimes — and API contracts between MFEs must be stable."*

---

## Q: "How do you architect feature modules in a large Angular application?"

> *"The architecture I use and recommend for large Angular apps is a three-layer structure. CoreModule contains singletons: auth service, HTTP interceptors, guards, and any service that should have exactly one instance for the lifetime of the app. You import CoreModule only once in AppModule — some add a guard to throw if it's imported more than once. SharedModule contains reusable UI: presentational components, pipes, directives. It's imported by every feature module and should never provide services. Feature modules contain all the business logic for one domain — products, orders, users, admin — with their own routes, components, and services. Feature modules are lazy-loaded. Inside each feature module, I further separate into containers — smart components that own data and NgRx interactions — and presentational components that only receive @Input. This scales, because features are isolated — teams can work in their feature without affecting others."*

---

## Q: "How do you handle a global HTTP error strategy in a large Angular app?"

> *"A global HTTP error strategy belongs in an HttpInterceptor. The interceptor wraps every request's response stream with catchError. Based on the error status code, you take different actions: 401 means the session expired — you dispatch a logout action or call AuthService.logout() which clears tokens and redirects to the login page. 403 means the user is authenticated but not authorized — you navigate to a forbidden page. 409 might mean a conflict that needs user intervention. Server errors — 5xx — get logged to your error tracking service, and a toast notification tells the user something went wrong. You rethrow the error with throwError so individual components can still add context-specific error handling if needed. The key insight is that this centralized interceptor means you never duplicate error handling logic across 50 service methods."*

---

## Q: "What is the Smart/Dumb component pattern? Why does it matter for performance and testability?"

> *"The Smart and Dumb pattern — also called Container and Presentational — is about separating components by their relationship to data. A Smart or Container component owns data fetching and business logic — it connects to NgRx selectors, calls services, handles routing. A Dumb or Presentational component is purely a function of its inputs — it receives data via @Input and communicates upward via @Output, with no direct service or store dependencies. The performance implication is that Dumb components can always safely use OnPush — they only update when their @Input references change. The testing implication is that Dumb components can be tested with just their @Input values and assertions on what they render — no mocking needed. The architecture implication is that Dumb components are reusable — the same ProductCard can be used in search results, a category page, and a recommendations carousel."*

---

## Q: "How would you implement token refresh with concurrent request queueing in Angular?"

> *"This is a common interview scenario that tests both HTTP interceptor knowledge and RxJS skills. When the interceptor catches a 401, it needs to: check if a refresh is already in progress to avoid triggering multiple simultaneous refresh calls; if no refresh is in progress, call the refresh endpoint, update the stored token, and retry the original request with the new token; if a refresh is already in progress from another concurrent request, queue up and wait for the in-progress refresh to complete, then retry. The implementation uses a BehaviorSubject initialized to null. When refresh starts, you set it to null. When it completes, you set it to the new token. Queued requests pipe the token subject through filter-not-null, take(1), and switchMap to retry with the new token. This pattern handles the scenario where a user with a slow connection has multiple in-flight requests that all 401 simultaneously — only one refresh fires and all others piggyback on it."*
