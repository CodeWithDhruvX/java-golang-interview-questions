# 📘 05 — Security, SSR & PWA
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- XSS prevention in Angular (automatic and manual)
- `DomSanitizer` and trusted values
- CSRF protection with Angular's `HttpClient`
- Angular Universal (SSR) — setup, benefits, transfer state
- Service Workers and PWA with Angular

---

## ❓ Most Asked Questions

### Q1. How does Angular prevent XSS attacks?

**XSS (Cross-Site Scripting)** occurs when malicious JavaScript is injected into a page. Angular prevents it automatically through **output escaping**.

```typescript
// Angular AUTOMATICALLY escapes all template bindings
@Component({
  template: `
    <!-- Angular escapes this — renders as text, not HTML -->
    <p>{{ userInput }}</p>
    <!-- Even if userInput = '<script>alert("xss")</script>' — it's safe -->

    <!-- Property binding is also safe for most DOM properties -->
    <p [textContent]="userInput"></p>
  `
})
export class SafeComponent {
  userInput = '<script>alert("xss")</script>';
  // Angular renders: &lt;script&gt;alert("xss")&lt;/script&gt;
}
```

```typescript
// ⚠️ DANGEROUS — bypassing Angular's sanitization
@Component({
  template: `<div [innerHTML]="htmlContent"></div>`
})
export class DangerousComponent {
  // ❌ This is dangerous if htmlContent comes from user input!
  htmlContent = '<img src=x onerror="alert(\'xss\')">';
  // Angular DOES sanitize [innerHTML] by default — strips dangerous elements
}

// ✅ If you MUST trust HTML (e.g., from your own CMS):
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

@Component({
  template: `<div [innerHTML]="trustedHtml"></div>`
})
export class CMSComponent {
  trustedHtml: SafeHtml;

  constructor(private sanitizer: DomSanitizer) {
    // Only use this for content you explicitly trust (NOT user input!)
    this.trustedHtml = this.sanitizer.bypassSecurityTrustHtml(
      '<b>Our <em>trusted</em> CMS content</b>'
    );
  }
}
```

**Angular's DomSanitizer trust methods:**

```typescript
// Only use these when you fully trust the source!
sanitizer.bypassSecurityTrustHtml(value)    // for [innerHTML]
sanitizer.bypassSecurityTrustStyle(value)   // for [style]
sanitizer.bypassSecurityTrustScript(value)  // for script elements
sanitizer.bypassSecurityTrustUrl(value)     // for [href], [src]
sanitizer.bypassSecurityTrustResourceUrl(value)  // for iframes, workers
```

---

### Q2. How do you protect Angular apps from CSRF?

Angular's `HttpClient` includes built-in **XSRF/CSRF** protection:

```typescript
// app.module.ts — enable XSRF protection
import { HttpClientXsrfModule } from '@angular/common/http';

@NgModule({
  imports: [
    HttpClientModule,
    HttpClientXsrfModule.withOptions({
      cookieName: 'XSRF-TOKEN',    // read CSRF token from this cookie
      headerName: 'X-XSRF-TOKEN'  // send it in this header
    })
  ]
})
export class AppModule {}

// Angular automatically:
// 1. Reads the XSRF-TOKEN cookie set by the server
// 2. Adds X-XSRF-TOKEN header to all mutating requests (POST, PUT, PATCH, DELETE)
// The server validates this header against its session cookie
```

---

### Q3. What is Angular Universal (SSR)? Why is it important?

**Angular Universal** enables **Server-Side Rendering (SSR)** — the server pre-renders Angular pages into HTML before sending to the browser.

```bash
# Add SSR to an existing Angular project
ng add @nguniversal/express-engine

# Serve the SSR app
npm run dev:ssr

# Build for production
npm run build:ssr
npm run serve:ssr
```

**Why SSR matters:**

| Aspect | CSR (Client-Side) | SSR (Angular Universal) |
|--------|-------------------|-------------------------|
| Initial load | Browser downloads JS, then renders | Server sends pre-rendered HTML |
| SEO | Poor (crawlers see empty page) | Excellent (bots see full HTML) |
| First Contentful Paint | Slow | Fast |
| Social sharing | Empty OG tags | Populated meta tags |
| Performance | More JS on client | Less client-side computation |

```typescript
// SSR-safe code — platform detection
import { isPlatformBrowser, isPlatformServer } from '@angular/common';
import { PLATFORM_ID } from '@angular/core';

@Component({ ... })
export class AnalyticsComponent implements OnInit {
  constructor(
    @Inject(PLATFORM_ID) private platformId: Object
  ) {}

  ngOnInit(): void {
    if (isPlatformBrowser(this.platformId)) {
      // Only runs in browser
      this.initializeGoogleAnalytics();
      window.scrollTo(0, 0);
    }
    if (isPlatformServer(this.platformId)) {
      // Only runs on server
      console.log('Server-side render');
    }
  }
}
```

---

### Q4. What is Transfer State in SSR?

```typescript
// Problem: Without Transfer State, SSR pre-renders with data, then browser bootstraps
// and re-fetches the same data — DOUBLE HTTP REQUEST!

// Solution: Transfer State — server serializes data into HTML, browser reads it

// In a server-rendered component:
import { TransferState, makeStateKey } from '@angular/platform-browser';

const PRODUCTS_KEY = makeStateKey<Product[]>('products');

@Component({ ... })
export class ProductListComponent implements OnInit {
  products: Product[] = [];

  constructor(
    private transferState: TransferState,
    private productService: ProductService,
    @Inject(PLATFORM_ID) private platformId: Object
  ) {}

  ngOnInit(): void {
    if (this.transferState.hasKey(PRODUCTS_KEY)) {
      // Browser: read data from HTML — no HTTP call!
      this.products = this.transferState.get(PRODUCTS_KEY, []);
      this.transferState.remove(PRODUCTS_KEY);
    } else {
      // Server: fetch data and store in transfer state
      this.productService.getAll().subscribe(products => {
        this.products = products;
        if (isPlatformServer(this.platformId)) {
          this.transferState.set(PRODUCTS_KEY, products);
        }
      });
    }
  }
}
```

> **Angular 17+:** Angular's new SSR uses `withHttpTransferCache()` which handles Transfer State automatically for most `HttpClient` requests.

---

### Q5. How do you add PWA features to Angular?

```bash
# Add PWA support
ng add @angular/pwa

# This automatically:
# - Generates ngsw-config.json (service worker config)
# - Adds manifest.webmanifest
# - Registers the service worker in app.module.ts
```

```json
// ngsw-config.json — control what gets cached
{
  "index": "/index.html",
  "assetGroups": [
    {
      "name": "app",
      "installMode": "prefetch",  // cache on install
      "resources": {
        "files": ["/favicon.ico", "/index.html", "/*.css", "/*.js"]
      }
    },
    {
      "name": "assets",
      "installMode": "lazy",     // cache on first access
      "resources": {
        "files": ["/assets/**"],
        "urls": ["https://fonts.googleapis.com/**"]
      }
    }
  ],
  "dataGroups": [
    {
      "name": "api-freshness",
      "urls": ["/api/products/**"],
      "cacheConfig": {
        "strategy": "freshness",  // network-first, fallback to cache
        "maxSize": 100,
        "maxAge": "3d",
        "timeout": "10s"
      }
    }
  ]
}
```

```typescript
// Prompt users to update when new version is available
@Component({ ... })
export class AppComponent implements OnInit {
  constructor(private swUpdate: SwUpdate) {}

  ngOnInit(): void {
    if (this.swUpdate.isEnabled) {
      this.swUpdate.versionUpdates.pipe(
        filter(evt => evt.type === 'VERSION_READY')
      ).subscribe(() => {
        if (confirm('New version available. Reload to update?')) {
          window.location.reload();
        }
      });
    }
  }
}
```
