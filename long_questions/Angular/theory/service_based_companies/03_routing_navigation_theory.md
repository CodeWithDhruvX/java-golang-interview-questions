# 🗣️ Theory — Routing & Navigation
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does Angular's router work? Walk me through setting it up."

> *"Angular's router maps URL paths to components. You define routes as an array of objects — each has a path and a component. The RouterModule.forRoot() call in AppModule registers these routes at the application level. In your HTML, router-outlet is the placeholder where the matched component renders. routerLink is the directive you use in templates instead of href — it integrates with Angular's router so navigation doesn't trigger a full page reload. The router reads the browser's URL on load, matches it against your route definitions top-to-bottom, and renders the matching component. The wildcard path '**' at the end catches all unmatched routes and maps to a 404 component."*

---

## Q: "How do you read route parameters in a component?"

> *"You inject ActivatedRoute into the component. ActivatedRoute gives you access to the current route context — its parameters, query parameters, data, and URL segments. For reading a route parameter like :id, you have two options. The snapshot approach — route.snapshot.paramMap.get('id') — reads the value once at the time the component loads. This is fine when the ID never changes while the component is alive. The Observable approach — route.paramMap.pipe(map(p => p.get('id'))) — gives you a stream that updates if the parameter changes, which happens when Angular reuses the component instance for different IDs. In a product detail page accessed from a list, always use the Observable approach — otherwise navigating from product 1 to product 2 won't refresh the component."*

---

## Q: "What are Route Guards? Explain CanActivate."

> *"Route guards are classes that intercept navigation and allow or block it. Angular has several guard types. CanActivate runs before a route is activated and answers 'can this user navigate to this route?'. CanDeactivate runs before leaving a route — useful for 'you have unsaved changes, are you sure you want to leave?' prompts. CanLoad prevents a lazy-loaded module from being downloaded at all if the user lacks permission. Resolve pre-fetches data before the component renders. For CanActivate, you implement the CanActivate interface, inject AuthService, and return true to allow or a UrlTree to redirect. The guard is then listed in the route definition under canActivate. Multiple guards can be listed — Angular runs them in order and requires all to return true."*

---

## Q: "What is lazy loading in Angular routing? Why does it matter?"

> *"Lazy loading means a feature module — and all its components — is not included in the initial JavaScript bundle. Instead, Angular only downloads and parses that module when the user first navigates to its route. You configure it with loadChildren using a dynamic import: loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule). The impact on initial load time is significant. If your app has an admin section, a reports section, and a user settings section, but most users only use the main section, there's no reason to download all that code on the first visit. Lazy loading splits your app into chunks — Angular CLI handles the code splitting automatically — and delivers chunks only on demand."*

---

## Q: "What is the difference between routerLink and router.navigate()?"

> *"Both navigate programmatically, but they're used in different contexts. routerLink is a directive you use in HTML templates — you bind it to a path or array of path segments like [routerLink]=['/products', product.id]. It's declarative and belongs in the template where you're rendering links and buttons. router.navigate() is a method you call from TypeScript — from a component method, after a form submission, after an API call succeeds. You inject the Router service and call this.router.navigate(['/dashboard']). There's also router.navigateByUrl() which takes a full URL string. The rule of thumb is: if the navigation is triggered by the user clicking a visible element in the template, use routerLink. If the navigation is triggered by logic, use router.navigate()."*

---

## Q: "What is a Resolve guard and when do you use it?"

> *"A Resolve guard pre-fetches data before the target component is activated. Without Resolve, you navigate to a product detail page, the component renders, then it fires an HTTP call, and the template shows loading state until the data arrives. With Resolve, Angular waits for the data to be fetched before rendering the component at all — so when the component initializes, the data is already there in the route's data snapshot. You implement the Resolve interface, fetch the data in the resolve method, and the component reads it from route.snapshot.data. The downside is the navigation itself appears delayed — the URL doesn't change until the data is ready. So Resolve is great for pages where you definitely can't render anything meaningful without the data, but use a loading state in the component when you want the URL to update immediately."*

---

## Q: "How do child routes work in Angular?"

> *"Child routes let you nest routes inside a parent route, creating hierarchical URL structures. The parent component must have its own router-outlet — this is where child components render. You define children as an array in the parent route definition. For example, a products route could have children: an empty path for the list view, ':id' for detail view, and ':id/edit' for the edit form. The parent handles any shared layout — sidebar, header, breadcrumbs — and the child route fills in the content area. This is how you build master-detail layouts and tabbed views in Angular. RouterModule.forChild() is used in feature modules, while forRoot() is only called once in AppModule."*
