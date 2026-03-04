# 🗣️ Theory — Security, SSR & PWA
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does Angular prevent XSS attacks?"

> *"Angular has XSS prevention built in at the template level. Any value you put into a template via interpolation or property binding is automatically escaped — if a user submits a value containing script tags, Angular renders them as the literal text of the script tag, not as executable HTML. The one escape hatch is [innerHTML] — you can bind raw HTML to the innerHTML property. Angular does sanitize [innerHTML] automatically by stripping dangerous elements, but you can fully bypass this using DomSanitizer.bypassSecurityTrustHtml(). This is a red-flag method — you should only call it on HTML that originates from your own server, never on user-provided content. The content security policy on your server is the second layer of defense — a CSP header that blocks inline script execution prevents XSS even if a bypass sneaks through."*

---

## Q: "What is Angular Universal? What are the benefits of SSR?"

> *"Angular Universal enables server-side rendering — the server runs Angular and pre-renders the page into HTML before sending it to the browser. The browser receives fully rendered HTML, paints it immediately, then downloads the Angular bundle and hydrates the page to make it interactive. The benefits are: first, SEO — search engine crawlers see the full HTML content rather than an empty index.html, which is critical for e-commerce, blogs, and marketing pages. Second, perceived performance — the first contentful paint is much faster because the user sees content before the JavaScript loads. Third, social sharing — Open Graph meta tags are populated server-side so link previews on Twitter and Facebook show correct images and titles. The trade-off is added infrastructure complexity — you need a Node.js server, and components must be platform-aware, checking for browser-only APIs like localStorage or window."*

---

## Q: "What is Transfer State in Angular Universal? Why is it needed?"

> *"Transfer State solves the double-fetch problem in SSR. Without it, here's what happens: the server renders the page and makes HTTP calls to populate the data. The browser receives the pre-rendered HTML. The browser downloads Angular and bootstraps — and Angular re-runs ngOnInit, which fires the HTTP calls again. The user sees briefly correct content, then a loading state, then the same content again — double fetching for no benefit. Transfer State serializes the fetched data into the HTML on the server side — it's embedded as a hidden script tag. When Angular bootstraps on the browser, it first checks Transfer State before making HTTP calls. If it finds the data there, it uses it and skips the HTTP call entirely. Angular 17's new SSR approach with httpTransferCache does this automatically for HttpClient calls — you get Transfer State benefits without manual setup."*

---

## Q: "How do you add PWA capabilities to an Angular app? What happens under the hood?"

> *"You run ng add @angular/pwa, which adds the service worker package and generates ngsw-config.json — the Angular service worker configuration. The service worker is installed in the browser when the user first visits. From that point, it serves cached assets from its cache storage rather than the network, which enables offline functionality and dramatically faster repeat loads. ngsw-config.json lets you configure two types of groups: asset groups for static resources like JS bundles and CSS — you'd prefetch these on install so the app works offline immediately — and data groups for API calls, where you choose between network-first for fresh data and cache-first for performance. Angular's service worker also handles app updates gracefully — when you deploy a new version, it notifies running clients that an update is available. You listen for this event in SwUpdate and prompt the user to reload."*

---

## Q: "How do you write platform-safe code in Angular Universal?"

> *"The fundamental challenge with SSR is that server-side, browser APIs don't exist — window, document, localStorage, navigator are all undefined in Node.js. If your component code references window, it will crash on the server. Angular provides isPlatformBrowser and isPlatformServer from @angular/common, and PLATFORM_ID from @angular/core. You inject PLATFORM_ID in your constructor and then guard browser-only code with if(isPlatformBrowser(this.platformId)). For localStorage, the standard approach is wrapping it in a service that checks the platform — the service returns null or empty values on the server. For third-party analytics scripts, you initialize them only on the browser. The angular:universal documentation recommends treating server-only components as a 'no-side-effects' zone — no DOM manipulation, no browser event listeners, just data fetching and template rendering."*
