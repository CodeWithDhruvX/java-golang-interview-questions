# 🔴 Performance Optimization & Advanced Topics

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Infosys): Lazy loading, AOT basics, basic optimization
> - 🚀 **Product-Based** (Google, Netflix, Zepto): Bundle analysis, virtual scroll, zone-less, profiling
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🔹 Performance Optimization

---

### 1. How to improve performance in Angular apps? 🟡 | 🏭🚀

"Angular performance comes from a layered approach — I address it at the **build, change detection, rendering, and network** levels:

**Build-level:**
- **AOT compilation** — `ng build --configuration=production` (AOT enabled by default in prod)
- **Tree shaking** — Remove unused code (ensure `sideEffects: false` in package.json)
- **Lazy loading** — Only load modules/components when needed
- **Budget limits** — `angular.json` budget alerts catch bundle bloat

**Change Detection:**
- `ChangeDetectionStrategy.OnPush` on all pure/presentational components
- `trackBy` on all `*ngFor` loops
- `async` pipe instead of manual subscriptions

**Rendering:**
- `@defer` blocks (Angular 17+) for deferred component loading
- Virtual scrolling (`CdkVirtualScrollViewport`) for large lists
- `ngZone.runOutsideAngular()` for high-frequency events like `mousemove`

**Network:**
- Pre-fetching data with resolvers or route preloading
- HTTP response caching with interceptors
- Image lazy loading with `loading="lazy"`"

#### In Depth
**Core Web Vitals** are my north star for Angular performance. LCP (Largest Contentful Paint) is improved by SSR/prerendering. CLS (Cumulative Layout Shift) is improved by always setting explicit `width`/`height` on images. INP (Interaction to Next Paint — replacing FID) is improved by `OnPush` + zoneless CD, because it measures how fast the UI responds to user interactions.

---

### 2. What is lazy loading? 🟢 | 🏭🚀

"**Lazy loading** defers loading a feature module (or component) until the user actually navigates to it. This reduces the **initial bundle size**, improving startup TTI (Time to Interactive).

```typescript
// app-routing.module.ts
{
  path: 'admin',
  loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule)
}

// Angular 14+ — Standalone component lazy loading
{
  path: 'dashboard',
  loadComponent: () => import('./dashboard.component').then(c => c.DashboardComponent)
}
```

Without lazy loading, `ng build` combines ALL feature code into one large `main.js`. With lazy loading, Angular creates separate `chunk-xxx.js` files — each loaded on-demand."

#### In Depth
**Preloading strategies** complement lazy loading — they background-load lazy modules after the initial app loads, so they're ready before the user navigates to them:

```typescript
RouterModule.forRoot(routes, {
  preloadingStrategy: PreloadAllModules // Load all lazy modules after boot
})
```

I use a **custom preloading strategy** that only preloads routes tagged with `data: { preload: true }`, skipping rarely-visited admin pages while preloading the top-3 most visited routes. This balances cold navigation speed with bundle efficiency.

---

### 3. What is Ahead-of-Time (AOT) compilation? 🟡 | 🏭🚀

"**AOT compilation** means Angular compiles TypeScript and HTML templates **at build time** (on the developer's machine / CI server), rather than at **run time** in the user's browser (JIT — Just-in-Time).

**Benefits:**
- **Smaller bundles** — Angular compiler is not shipped to the browser
- **Faster rendering** — Pre-compiled templates are instantly usable
- **Earlier error detection** — Template errors caught at build time
- **Better security** — No dynamic template evaluation

```bash
ng build --configuration=production  # AOT is always enabled in production
ng serve --aot                        # Enable AOT in dev (optional but recommended)
```

AOT is **enabled by default** in production builds since Angular 9 (Ivy)."

#### In Depth
AOT's compilation is powered by the **Ivy compiler (ngtsc)** which performs **partial compilation** for library authors — libraries ship a partially compiled output (`.metadata.json`) that the application compiler finalizes. This enables independent library versioning while maintaining compatibility. The alternative (full compilation) required libraries to match the exact Angular version.

---

### 4. What is tree shaking? 🟡 | 🏭🚀

"**Tree shaking** is a build optimization that removes **unused code** (dead code) from the final bundle. It works by analyzing ES module `import`/`export` statements and removing any code that's never imported anywhere.

Angular's build system (Webpack/esbuild) performs tree shaking automatically during production builds.

To enable tree shaking:
1. Use **ES modules** (`import`/`export`) — not CommonJS (`require()`)
2. Mark packages as side-effect-free: `{ "sideEffects": false }` in `package.json`
3. Use `providedIn: 'root'` for services — not `providers: [Service]` in `NgModule`

```typescript
// ✅ Tree-shakable service — removed if never injected
@Injectable({ providedIn: 'root' })
export class RarelyUsedService {}

// ❌ NOT tree-shakable — always bundled
@NgModule({ providers: [RarelyUsedService] })
```"

#### In Depth
Angular's Ivy compiler enables **component-level tree shaking**. In Ivy, components declare their own dependencies (directives, pipes, components they use) locally, rather than through module-level declarations. This means if a lazy-loaded module only uses 5 of 20 Material components, only those 5 are included in the lazy chunk — not the entire Material library. This is a fundamental improvement over View Engine's approach.

---

### 5. How to reduce bundle size? 🟡 | 🏭🚀

"Bundle size is a **ongoing discipline**, not a one-time fix. My checklist:

**1. Analyze the bundle:**
```bash
ng build --stats-json
npx webpack-bundle-analyzer dist/app/stats.json
# Or: source-map-explorer dist/app/main.js
```

**2. Lazy load aggressively:**
- Every feature module should be lazy loaded
- `loadComponent` for standalone components

**3. Avoid heavy libraries:**
- Replace `lodash` with native ES6 (`_.map` → `Array.map`)
- Replace `moment.js` (330KB) with `date-fns` (tree-shakable) or `Intl.DateTimeFormat`

**4. Optimize imports:**
```typescript
// ❌ Imports entire lodash
import _ from 'lodash';
// ✅ Imports only the needed function
import debounce from 'lodash/debounce';
```

**5. Set budgets in `angular.json`:**
```json
{ \"type\": \"initial\", \"maximumWarning\": \"500kb\", \"maximumError\": \"1mb\" }
```

**6. Use `@defer` (Angular 17+) for below-fold components**"

#### In Depth
The most impactful bundle reduction in my experience:

1. **Angular Material per-component imports** — Instead of importing `MatButtonModule` from the barrel, import `MatButton` directly (Angular Material 17+ supports per-component imports)
2. **Zone.js removal** — Removing zone.js saves ~35KB gzipped and eliminates CD overhead
3. **JSON bundles** — Large static JSON (translations, config) should be lazy-loaded on demand, not bundled

---

### 6. What is Angular Virtual Scrolling? 🔴 | 🚀

"**Virtual scrolling** renders only the visible items in a list, recycling DOM nodes as the user scrolls. For lists of thousands of items, it's the difference between a **5ms render** and a **5-second freeze**.

Angular CDK (`@angular/cdk/scrolling`) provides:
```typescript
import { ScrollingModule } from '@angular/cdk/scrolling';

// In template:
<cdk-virtual-scroll-viewport itemSize="72" style="height: 400px;">
  <div *cdkVirtualFor="let item of items; trackBy: trackById">
    <app-product-card [product]="item"></app-product-card>
  </div>
</cdk-virtual-scroll-viewport>
```

`itemSize` is the height in pixels of each item. For **variable-height items**, use `AutoSizeVirtualScrollStrategy`.

I use virtual scrolling for any list with more than ~200 items — below that, the overhead of virtual scrolling can outweigh the benefit."

#### In Depth
Virtual scrolling works by maintaining a **buffer** of rendered items above and below the viewport (configurable via `minBufferPx` and `maxBufferPx`). When the user scrolls, items outside the buffer range are detached from the DOM (and their components go through `ngOnDestroy`), while new items entering the buffer range are instantiated (going through `ngOnInit`). For expensive components, using `ChangeDetectorRef.detach()` on out-of-viewport items instead of destroying them can be more performant.

---

### 7. What is Angular Profiler and how to use it? 🔴 | 🚀

"The **Angular Profiler** (`ng.profiler.timeChangeDetection()`) measures the time Angular spends on change detection cycles.

```typescript
// In browser console:
ng.profiler.timeChangeDetection({ record: true })
// Runs 500 rapid change detection cycles, reports average time

// Example output:
// ran 500 change detection cycles
// 1.5 ms per check
```

**Angular DevTools** (Chrome extension) provides a visual profiler showing:
- Which components triggered CD
- How long each component took
- How many times each component was checked

I use this to find **CD performance bottlenecks** — usually components with `Default` strategy that do expensive computations in getters or template expressions."

#### In Depth
The single biggest performance win I find via profiling is **getter expressions used in templates**:

```html
<!-- ❌ Getter runs on EVERY CD cycle -->
<div>{{ processedProducts }}</div>
<!-- When processedProducts is a getter that filters 1000 items -->

<!-- ✅ Cache in property, update only when source changes -->
<div>{{ cachedProcessedProducts }}</div>
```

Template expressions should be **pure and fast** — preferably simple property reads. Push computation into `ngOnChanges` or reactive streams with `distinctUntilChanged()`.

---

### 8. What is Angular Hydration? 🔴 | 🚀

"**Hydration** in Angular (introduced in v16, stable in v17) is the process of making a **server-rendered HTML page interactive** on the client without re-rendering the entire page.

Before hydration, Angular Universal would:
1. Server renders HTML → Browser shows content fast
2. Angular bootstraps → Destroys the server HTML → Re-renders from scratch → Flash/CLS!

With hydration:
1. Server renders HTML → Browser shows content fast
2. Angular bootstraps → **Attaches to existing DOM** instead of replacing it → No flash!

```typescript
// Enable in main.ts
bootstrapApplication(AppComponent, {
  providers: [
    provideClientHydration(), // Enable hydration
    provideRouter(routes)
  ]
});
```

Hydration significantly improves **LCP** and eliminates layout shift during SSR → CSR transition."

#### In Depth
Angular hydration uses **DOM reconciliation** — the server-rendered DOM is annotated with markers, and Angular's client-side bootstrap uses these markers to identify which DOM nodes correspond to which component views. If the server and client render different HTML (non-deterministic rendering), Angular detects the **hydration mismatch** and falls back to full client-side re-rendering for the mismatched component, logging a warning in development mode.

---

### 9. What is `@defer` in Angular 17+? 🔴 | 🚀

"**`@defer`** (Angular 17) is a template-level API for lazy loading components — no routing or dynamic component loading APIs needed.

```html
<!-- Component loads when visible in viewport -->
@defer (on viewport) {
  <app-product-reviews [productId]="id" />
} @loading {
  <app-skeleton-loader />
} @error {
  <p>Failed to load reviews. <button (click)="retry()">Retry</button></p>
} @placeholder {
  <div class="reviews-placeholder">Reviews loading...</div>
}
```

Trigger conditions:
- `on viewport` — Intersects the viewport (IntersectionObserver)
- `on idle` — Browser is idle
- `on interaction` — User clicks or hovers
- `on timer(5s)` — After 5 seconds
- `when condition` — Custom boolean expression"

#### In Depth
`@defer` is the most powerful lazy loading primitive Angular has introduced. Unlike route-level lazy loading (which loads an entire module), `@defer` works at the **component level within a single route**. This enables progressive rendering patterns: critical above-the-fold content renders eagerly, while non-critical below-fold content (reviews, recommendations, ads) defers until the browser is idle — dramatically improving LCP and TTI without code changes to the deferred components.

---

## 🔹 Advanced Topics

---

### 10. What is Angular Universal (SSR)? 🟡 | 🏭🚀

"**Angular Universal** enables **Server-Side Rendering (SSR)** — rendering Angular components on the server (Node.js) and sending the fully rendered HTML to the browser.

Benefits:
- **SEO** — Search engines see full content (not empty `<div id="app"></div>`)
- **Faster First Contentful Paint (FCP)** — Users see content before JS loads
- **Social sharing** — Correct OG meta tags for social media cards

Setup:
```bash
ng add @angular/ssr
```

This adds `server.ts`, configures `@angular/platform-server`, and sets up the Universal build target.

Key consideration: Server rendering means **no `window`, `document`, or `localStorage`**. Use `isPlatformBrowser()` to guard browser-specific code."

#### In Depth
The biggest challenge in SSR is **isomorphic code** — code that runs correctly in both Node.js and the browser. Angular provides `PLATFORM_ID` injection token and `isPlatformBrowser()` / `isPlatformServer()` guards. For state transfer (avoiding duplicate API calls on server and client), use `TransferState` (Angular's built-in mechanism) to serialize server-fetched data and reuse it on the client without making the same API call again.

---

### 11. What is dynamic component loading? 🔴 | 🚀

"**Dynamic component loading** creates and inserts components at runtime — without using them in templates statically.

```typescript
@Component({ template: `<ng-container #host></ng-container>` })
export class ContainerComponent implements OnInit {
  @ViewChild('host', { read: ViewContainerRef }) host!: ViewContainerRef;

  constructor(private componentFactoryResolver: ComponentFactoryResolver) {}

  ngOnInit(): void {
    this.loadComponent(AlertComponent);
  }

  loadComponent(componentClass: Type<any>): void {
    this.host.clear();
    const componentRef = this.host.createComponent(componentClass);
    // Pass inputs
    componentRef.setInput('message', 'Hello from dynamic component!');
    // or: (componentRef.instance as AlertComponent).message = '...';
  }
}
```

In Angular 14+, `createComponent` is simplified:
```typescript
createComponent(AlertComponent) // No ComponentFactoryResolver needed
```"

#### In Depth
Dynamic components power **plugin architectures** — a core app can load feature components that aren't known at compile time (loaded from a lazy chunk or even a remote URL in Module Federation). Each dynamically created component goes through Angular's full lifecycle (`ngOnInit`, `ngOnChanges`) and is properly destroyed when `componentRef.destroy()` is called. I use this pattern for dynamic **widget dashboards** where users arrange and configure panels at runtime.

---
