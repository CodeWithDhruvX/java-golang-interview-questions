# 🔴 Security & Server-Side Rendering (SSR)

> 🏢 **Company Type Guide:**
> - 🏭 **Service-Based** (TCS, Cognizant): XSS basics, DomSanitizer, CSRF concepts
> - 🚀 **Product-Based** (Razorpay, CRED, PayU): JWT handling, CSP, SSR security, advanced threat models
>
> 🎯 **Experience Level:**
> - 🟢 **Fresher (0–1 yr)** | 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

## 🔹 Security

---

### 1. What is Cross-Site Scripting (XSS)? 🟢 | 🏭🚀

"**XSS (Cross-Site Scripting)** is an attack where malicious scripts are injected into a web page and executed in the victim's browser, allowing attackers to:
- Steal cookies/localStorage (including auth tokens)
- Perform actions on behalf of the user
- Capture keystrokes and form data

**Types of XSS:**
1. **Stored XSS** — Malicious script is saved to the database (comment box attack)
2. **Reflected XSS** — Script injected via URL parameters
3. **DOM-based XSS** — Script injected via JavaScript DOM manipulation

Angular prevents XSS by **automatically escaping** all dynamic values inserted into the DOM via interpolation `{{ }}` and property binding. User input is treated as plain text, not HTML."

#### In Depth
Angular's security model is based on **trusted vs untrusted values**. By default, all values coming from external sources (APIs, user input, URL params) are **untrusted**. Angular escapes them automatically. The only way to insert HTML is via `[innerHTML]`, `[outerHTML]`, or `DomSanitizer.bypassSecurityTrust*()` — and Angular always warns in the console when you use these. I do security audits by `grep`-ing for these patterns across the codebase to review each use manually.

---

### 2. How to prevent XSS in Angular? 🟡 | 🏭🚀

"Angular's **built-in defenses** prevent most XSS automatically:

**1. Automatic escaping — Angular escapes all template values:**
```html
<!-- If user.name = '<script>alert("XSS")</script>' -->
{{ user.name }}  <!-- Renders as text: <script>alert("XSS")</script> -->
```

**2. Safe HTML rendering — DomSanitizer scrubs dangerous HTML:**
```typescript
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

constructor(private sanitizer: DomSanitizer) {}

getSafeHtml(htmlString: string): SafeHtml {
  return this.sanitizer.sanitize(SecurityContext.HTML, htmlString)!;
  // Removes <script>, event handlers (onclick), etc.
}

// Template:
// <div [innerHTML]="getSafeHtml(userContent)"></div>
```

**3. Never bypass security unless absolutely necessary:**
```typescript
// ⚠️ DANGEROUS — use only for trusted server-side HTML
this.sanitizer.bypassSecurityTrustHtml(serverContent);
```"

#### In Depth
Even with Angular's auto-escaping, **DOM-based XSS** can occur if you use `ElementRef.nativeElement.innerHTML = userInput` directly. This bypasses Angular's sanitization pipeline entirely. Always use `Renderer2.setProperty(el, 'innerHTML', value)` which routes through Angular's security layer, or use `[innerHTML]` binding which sanitizes automatically. Treat direct DOM access as a security code smell.

---

### 3. How does Angular prevent CSRF? 🔴 | 🚀

"Angular helps with **CSRF (Cross-Site Request Forgery)** through `HttpClientXsrfModule`:

```typescript
// app.module.ts
import { HttpClientXsrfModule } from '@angular/common/http';

@NgModule({
  imports: [
    HttpClientModule,
    HttpClientXsrfModule.withOptions({
      cookieName: 'XSRF-TOKEN',  // Cookie set by server
      headerName: 'X-XSRF-TOKEN' // Header Angular sends
    })
  ]
})
```

Angular reads the XSRF token from the `XSRF-TOKEN` cookie and automatically sends it as `X-XSRF-TOKEN` header in non-GET requests. The server validates this header ensures the request came from the legitimate Angular app.

CSRF tokens work because: an attacker's domain cannot read cookies from your domain (Same-Origin Policy), so they cannot forge the token."

#### In Depth
The server must cooperate by **setting the `XSRF-TOKEN` cookie** (readable by JavaScript, not HttpOnly). Spring Security, Django REST, and other backends support this automatically. On the Angular side, `HttpClientXsrfModule` only adds the header for **state-mutating requests** (POST, PUT, DELETE, PATCH) — GET requests don't need CSRF protection because they shouldn't have side effects. For APIs that use other CSRF mechanisms (custom headers, token in body), I implement a custom interceptor.

---

### 4. What is DomSanitizer? 🟡 | 🏭🚀

"**`DomSanitizer`** is Angular's built-in security service that sanitizes potentially dangerous values before they're inserted into the DOM.

```typescript
import { DomSanitizer, SafeResourceUrl, SafeHtml, SafeUrl } from '@angular/platform-browser';

@Injectable({ providedIn: 'root' })
export class MediaService {
  constructor(private sanitizer: DomSanitizer) {}

  // For embedding YouTube/PDF iframes
  getSafeUrl(url: string): SafeResourceUrl {
    return this.sanitizer.bypassSecurityTrustResourceUrl(url);
  }

  // For dynamic HTML content from CMS
  getSafeHtml(html: string): SafeHtml {
    // sanitize() removes scripts and dangerous attributes automatically
    return this.sanitizer.sanitize(SecurityContext.HTML, html) ?? '';
  }
}
```

**Security contexts:**
- `SecurityContext.HTML` — For `[innerHTML]`
- `SecurityContext.URL` — For `[href]`, `[src]`
- `SecurityContext.RESOURCE_URL` — For `[src]` in `<iframe>`, `<script>`
- `SecurityContext.STYLE` — For `[style]`"

#### In Depth
`sanitize()` is the **safe** API — it removes dangerous content and returns a safe string. `bypassSecurityTrust*()` methods are for cases where you've externally verified the content is safe (e.g., it comes from your own backend, is already sanitized server-side). I document every usage of `bypassSecurityTrust*` with a code comment explaining why it's safe, making security audits faster. An audit that finds an undocumented bypass is automatically suspicious.

---

### 5. How to manage JWT tokens securely in Angular? 🔴 | 🚀

"JWT storage strategy is a **security trade-off**:

**Option 1: `localStorage` / `sessionStorage`** (simple but XSS-vulnerable):
- Accessible via JS → vulnerable to XSS attacks
- `localStorage` persists across tabs and browser restarts
- `sessionStorage` cleared when tab closes

**Option 2: HttpOnly Cookies** (more secure):
- Not accessible via JS → XSS-resistant
- Server sets: `Set-Cookie: token=xxx; HttpOnly; Secure; SameSite=Strict`
- Angular HttpClient automatically sends cookies with `withCredentials: true`

```typescript
// For cookie-based auth
this.http.post('/api/login', credentials, { withCredentials: true }).subscribe();
// All subsequent requests:
this.http.get('/api/data', { withCredentials: true }).subscribe();
```

**My recommendation**: Use **HttpOnly cookies** for auth tokens. Use an interceptor to handle 401s (token expiry) and trigger refresh flow automatically."

#### In Depth
For production-grade JWT handling, I implement a **token refresh interceptor**: when a 401 is received, intercept it, call the refresh endpoint, and retry the original request with the new token — all transparently without the user losing their session:

```typescript
catchError((error: HttpErrorResponse) => {
  if (error.status === 401 && !request.url.includes('/refresh')) {
    return this.authService.refreshToken().pipe(
      switchMap(() => next.handle(request)) // Retry with new token
    );
  }
  return throwError(() => error);
})
```

---

## 🔹 Server-Side Rendering (SSR)

---

### 6. How to set up Angular Universal? 🟡 | 🏭🚀

"**Angular Universal** adds SSR to an existing Angular app:

```bash
ng add @angular/ssr
```

This generates:
- `server.ts` — Express.js server that renders Angular
- `app.config.server.ts` — Server-side providers (no browser APIs)
- Updated `angular.json` with server build target

```bash
# Build for SSR
ng build && ng run app:server

# Or development:
ng serve --ssr
```

Key isomorphic code patterns:
```typescript
// Never use browser APIs directly in services
// ❌ window.localStorage.getItem('key')
// ✅ Use injection token to abstract platform

export const BROWSER_STORAGE = new InjectionToken<Storage>('Browser Storage', {
  providedIn: 'root',
  factory: () => {
    return isPlatformBrowser(inject(PLATFORM_ID)) ? localStorage : {} as Storage;
  }
});
```"

#### In Depth
SSR fundamentally changes Angular's execution context. Issues I consistently encounter and solve:

1. **`window is undefined`** — Use `isPlatformBrowser()` guards
2. **Double HTTP calls** — Use `TransferState` to pass server-fetched data to client
3. **Absolute URLs required** — `HttpClient` calls from server need full URLs (no relative `/api`)
4. **Session state** — Server has no cookies by default; pass request context using `REQUEST` injection token from `@nguniversal/express-engine`

---

### 7. What is TransferState in Angular SSR? 🔴 | 🚀

"**`TransferState`** is a key-value store that transfers server-computed state to the browser during hydration, preventing duplicate API calls.

```typescript
// In a server-side service
import { TransferState, makeStateKey } from '@angular/core';

const PRODUCTS_KEY = makeStateKey<Product[]>('products');

@Injectable({ providedIn: 'root' })
export class ProductService {
  constructor(
    private http: HttpClient,
    private transferState: TransferState,
    @Inject(PLATFORM_ID) private platformId: Object
  ) {}

  getProducts(): Observable<Product[]> {
    // Server: fetch from API, store in TransferState
    if (isPlatformServer(this.platformId)) {
      return this.http.get<Product[]>('/api/products').pipe(
        tap(products => this.transferState.set(PRODUCTS_KEY, products))
      );
    }

    // Client: use transferred state, avoiding duplicate API call
    if (this.transferState.hasKey(PRODUCTS_KEY)) {
      const products = this.transferState.get(PRODUCTS_KEY, []);
      this.transferState.remove(PRODUCTS_KEY); // Clear after use
      return of(products);
    }

    return this.http.get<Product[]>('/api/products');
  }
}
```"

#### In Depth
`TransferState` serializes the transferred data into the HTML as a JSON script tag (`<script id="serverApp-state" type="application/json">{...}</script>`). The Angular browser bundle reads and deserializes this during bootstrap. Without `TransferState`, the browser-side Angular would make the same API calls the server already made — doubling backend load and causing a **flash of empty content** before data arrives. With hydration + `TransferState`, the page is fully rendered and interactive from the first byte.

---
