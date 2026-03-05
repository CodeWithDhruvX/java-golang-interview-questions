# Observability and Security in Microfrontends

Distributing a frontend into multiple independently deployed chunks introduces severe challenges regarding debugging, monitoring, and enforcing security policies. Senior Engineers at product companies are expected to design robust observability and security perimeters for MFEs.

## 1. How do you track and debug an error that spans across multiple Microfrontends? (Distributed Tracing for the Browser)

**Answer:**
When an error occurs in an MFE architecture, determining "whose fault is it" is the hardest part. The standard `console.error` is insufficient. You need Distributed Tracing applied to the frontend.

**The Strategy:**
1.  **Correlation IDs:**
    *   When the Host application initializes the user session, it generates a unique `SessionID`.
    *   Every time a specific action occurs (e.g., clicking 'Checkout'), the Host generates a unique `TraceID`.
2.  **Context Propagation:**
    *   These IDs must be passed to every child MFE. This can be done via CustomEvent payloads, URL parameters, or injected props at mount time.
    *   Crucially, every single HTTP network request made by *any* MFE back to the server APIs must include these IDs in the HTTP Headers (e.g., `X-B3-TraceId`, `X-Correlation-ID`).
3.  **Centralized Logging (Sentry / Datadog / New Relic):**
    *   Each MFE must initialize its own instance of the logger (or use a shared singleton provided by the Host, though independent instances are safer to avoid configuration conflicts).
    *   Each instance must tag every log and error with the current `TraceID`, `MFE_Name` (e.g., `app: cart_mfe`), and the MFE's exact `Version/GitHash`.
4.  **The Result:** If the Checkout MFE throws a JavaScript error, you can query Datadog for that `TraceID`. The waterfall graph will show:
    *   Host MFE initiated the flow.
    *   Host MFE fired `CHECKOUT_START` event.
    *   Checkout MFE received event and made API call.
    *   API returned 500 (Backend logged it with the same TraceID).
    *   Checkout MFE threw unhandled JS exception.

## 2. Implementing Content Security Policy (CSP) is notoriously tricky with Webpack Module Federation. Why is that, and how do you configure it safely?

**Answer:**
Content Security Policy (CSP) strictly dictates from which domains the browser is allowed to download and execute scripts (`script-src`). 

**The Problem:**
Webpack Module Federation dynamically injects `<script>` tags at runtime to fetch `remoteEntry.js` and Javascript chunks from remote URLs. If your Host app's CSP only allows `script-src 'self'`, the browser will block the remote scripts, and the MFE will fail to load. Furthermore, Webpack heavily relies on `eval()` in development mode, and sometimes injects inline scripts, violating strict CSPs.

**The Solution:**
1.  **Dynamic CSP via Nonces (Recommended for strict security):**
    *   Your backend server generates a cryptographically random string (a nonce) on every page request.
    *   The server injects this nonce into the CSP header: `Content-Security-Policy: script-src 'nonce-XYZ123' 'strict-dynamic'`.
    *   The server also injects this nonce into a global variable or meta tag in the HTML.
    *   You must configure Webpack to use this nonce when dynamically loading remote chunks. In Webpack 5, you configure `__webpack_nonce__` before loading federation.
2.  **Allow-listing Domains (If nonces are too complex):**
    *   Explicitly list the CDN domains of the Remote MFEs in the Host's CSP header: 
        `script-src 'self' https://cdn.cart-app.com https://cdn.catalog-app.com`.
    *   *Drawback:* Every time a new MFE domain is added, the Host's infrastructure (Nginx/Node) must be updated.
3.  **Disable `eval` in Production:** Ensure Webpack `devtool` is set to `source-map` (not `eval-source-map`) in production builds to comply with `unsafe-eval` restrictions.

## 3. How do you handle Cross-Origin Resource Sharing (CORS) when Remote Microfrontends are hosted on different domains?

**Answer:**
If the Host (e.g., `app.company.com`) tries to fetch `remoteEntry.js` or data APIs from a Remote hosted elsewhere (e.g., `checkout.company.com`), the browser will enforce CORS restrictions natively.

**Strategies:**
1.  **Under a Single Domain with Path Routing (Best Practice):**
    *   Deploy all MFEs behind a unified Reverse Proxy / API Gateway (like Nginx, AWS API Gateway, or Cloudflare).
    *   The browser only ever speaks to `https://www.company.com`.
    *   The Gateway routes `/checkout-assets/*` to the Checkout S3 bucket, and `/catalog-assets/*` to the Catalog S3 bucket.
    *   *Benefit:* CORS is entirely avoided because from the browser's perspective, everything is Same-Origin (`www.company.com`).
2.  **Configuring CORS Headers (If multi-domain is necessary):**
    *   The server hosting the Remote MFE assets (the S3 bucket/CDN) must return the liberal `Access-Control-Allow-Origin` header.
    *   If using credentials (cookies) across domains, you must set `Access-Control-Allow-Credentials: true` and specify the exact Host origin (wildcards like `*` are not allowed).
    *   You must also properly handle HTTP `OPTIONS` preflight requests on the CDN/Server.

## 4. How do you prevent a malicious Remote Microfrontend from stealing data from the Host application or other Remotes?

**Answer:**
This is the hardest security problem in Client-Side MFEs. Because Webpack Module Federation and Single-SPA run in the **same browser Window environment (same Javascript execution context)**, isolation is not guaranteed by default.

**The Risk:** 
If a malicious actor compromises the `Cart_MFE` repository, they can add code to read the `window.localStorage` (stealing JWTs) or capture global DOM keystrokes (stealing credit cards) that belong to the `Checkout_MFE`.

**Mitigations:**
1.  **Zero-Trust Client State:** 
    *   Never store sensitive data (JWTs, PII) in `localStorage`, `sessionStorage`, or accessible `cookies`. 
    *   Use `HttpOnly` cookies. If the Cart MFE tries to read `document.cookie`, it won't see the token. Browsers automatically attach the HttpOnly cookie to API requests, so the Client JS never needs to touch the raw credential.
2.  **Iframes for High-Security Boundaries:**
    *   For extremely sensitive areas (like the PCI-compliant credit card entry form), do not use Module Federation. Embed the payment form as a cross-origin `<iframe>`. 
    *   The browser enforces a hard Sandbox around Iframes. The parent window cannot read the DOM or Javascript variables inside a cross-origin Iframe, effectively neutralizing XSS data theft.
3.  **Strict Code Reviews & Supply Chain Security:**
    *   Since MFEs run in the same context, you must trust the teams building them. Implement rigorous PR reviews, automated dependency scanning (Snyk, Dependabot), and strict CI/CD access controls for every individual MFE repository to prevent malicious code injection at the source.
