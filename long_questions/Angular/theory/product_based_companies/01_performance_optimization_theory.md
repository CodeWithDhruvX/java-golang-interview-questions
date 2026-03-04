# 🗣️ Theory — Performance Optimization
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does OnPush change detection improve performance? What are the trade-offs?"

> *"Default change detection checks every component in the tree on every async event — a button click anywhere triggers Angular to evaluate every binding in every component. With OnPush, Angular says: 'I'll only check this component if I have a compelling reason.' That reason is one of four things: an @Input() property received a new reference, an Observable subscribed via the async pipe emitted, an event was raised inside this component, or markForCheck() was called manually. The trade-off is immutability. OnPush uses reference equality, so mutating an object or array in-place — pushing to an array, updating a property — won't trigger a re-render because the reference didn't change. You must return new objects. This constraint is actually a feature: it forces clean data flow and makes the app more predictable. In a list of a hundred product cards, each with OnPush, Angular does almost no work on events that don't touch those cards."*

---

## Q: "Explain how trackBy prevents unnecessary DOM destruction in *ngFor."

> *"Without trackBy, when the array reference changes — even if 99% of the items are identical — Angular compares old and new DOM nodes by position, not by identity. It has no idea that the product at index 0 before and after is the same product. So it destroys all old DOM nodes and creates all new ones. This is expensive: destroying components triggers ngOnDestroy, cleaning up subscriptions, garbage collection; creating new ones re-runs constructors, ngOnInit, data fetching. With trackBy, you give Angular a function that returns a stable unique ID per item — typically the database ID. Angular can now diff old and new lists using these IDs, identify which items are new, which were removed, and which just moved. It only touches the DOM nodes that actually changed. This is the single biggest DOM performance win for list-heavy UIs."*

---

## Q: "What is the @defer block in Angular 17? How does it enable component-level lazy loading?"

> *"Lazy loading in Angular was traditionally route-level — you split by module and load on navigation. @defer in Angular 17 brings lazy loading to the component level, inside a template. You wrap a component in @defer and specify a trigger: on viewport means the chunk is loaded when the element scrolls into view; on idle means load when the browser is free; on interaction means load when the user hovers or clicks a trigger element. Angular generates a separate JavaScript chunk for the deferred component — it doesn't ship in the initial bundle at all. You also get @loading, @placeholder, and @error blocks to cleanly handle the loading states. This is transformative for page performance: a heavy analytics dashboard, a comments section, or a rich editor only downloads if and when the user needs it."*

---

## Q: "What is virtual scrolling and when should you use it?"

> *"Virtual scrolling renders only the DOM nodes currently visible in the viewport, rather than all items in the list. If you have 10,000 product rows and your viewport shows 20, virtual scrolling creates only those 20 DOM nodes — as you scroll, old nodes are recycled and populated with new data. Without it, 10,000 DOM nodes would be created upfront, consuming memory and making every browser repaint slow. Angular CDK's ScrollingModule provides the cdk-virtual-scroll-viewport directive and *cdkVirtualFor which replaces *ngFor. The itemSize in pixels is required — the virtual scroller needs to know how tall each item is to calculate scroll positions. Virtual scrolling is the right choice when you have lists with more than a few hundred items and each item is non-trivial to render."*

---

## Q: "What is NgZone.runOutsideAngular() and when would you use it?"

> *"Zone.js patches all browser async APIs — setTimeout, requestAnimationFrame, addEventListener — and triggers Angular's change detection after each one completes. This is what makes Angular reactive to user events automatically. But it also means that high-frequency events like scroll, mousemove, or animation frames trigger change detection constantly, even when no Angular data changes at all. NgZone.runOutsideAngular() lets you execute code without triggering change detection. I use it for animation loops, scroll event listeners, WebSocket message handlers, and canvas rendering — anything that fires many times per second where Angular's CD adds overhead. When I do need to update the UI from outside the zone, I call ngZone.run() to re-enter the zone for just that update."*

---

## Q: "How do you analyze and reduce Angular bundle size?"

> *"The first step is understanding what's in your bundle. Run ng build --stats-json and pass the stats file to webpack-bundle-analyzer — it gives you a visual treemap of every dependency and their sizes. From there, the biggest wins: lazy load every route so each feature module is a separate chunk, use providedIn:'root' services and standalone components to enable better tree-shaking, replace heavy libraries with lighter alternatives — moment.js with date-fns, lodash/fp with tree-shakeable lodash. Set budgets in angular.json — Angular CLI will warn or error if chunks exceed your limits, which enforces discipline. In Angular 16+, the new esbuild-based builder is also significantly faster and produces slightly smaller bundles. The goal is making the critical path — the code needed for the first meaningful paint — as small as possible."*

---

## Q: "What are the main memory leak patterns in Angular and how do you prevent them?"

> *"The most common source of memory leaks in Angular is unsubscribed Observables. When a component subscribes to a BehaviorSubject or WebSocket stream in ngOnInit and doesn't unsubscribe in ngOnDestroy, the subscription keeps the component alive in memory even after Angular has destroyed it — every emission still runs the handler. The cleanest prevention strategies: use the async pipe whenever possible — it auto-unsubscribes on destroy. For manual subscriptions, use takeUntil with a destroy$ Subject — you call destroy$.next() in ngOnDestroy and all piped subscriptions complete. Angular 16 introduced DestroyRef and takeUntilDestroyed — you inject DestroyRef and pipe takeUntilDestroyed(this.destroyRef), eliminating the need for a Subject and an OnDestroy hook. Event listeners added imperatively with addEventListener also need to be removed in ngOnDestroy."*
