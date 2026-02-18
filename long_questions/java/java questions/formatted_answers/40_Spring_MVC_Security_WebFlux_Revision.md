# 40. Spring MVC, Security & WebFlux (Final Revision)

**Q: DispatcherServlet (Front Controller Pattern)**
> "Imagine a hotel reception. You don't walk directly to the chef to order food. You go to the front desk.
>
> **DispatcherServlet** is that Front Desk.
> Every single HTTP request (`/login`, `/users`, `/home`) hits this servlet first.
> It checks the URL, looks up the 'Handler Mapping' to find the right Controller, and delegates the work. You never interact with it directly, but it runs the show."

**Indepth:**
> **Context Hierarchy**: `WebApplicationContext`. The DispatcherServlet creates its own child context (containing Controllers, ViewResolvers) which inherits from the Root WebApplicationContext (containing Services, Repositories). This separation allows multiple DispatcherServlets to share common beans.


---

**Q: @PathVariable vs @RequestParam**
> "**@PathVariable**: It's part of the **Identity** of the resource.
> `/users/123` -> 123 identifies a specific user.
>
> "**@RequestParam**: It's for **Filtering** or Sorting.
> `/users?country=US&sort=age` -> You are looking at the users resource, but filtering it.
> Use PathVariable for IDs, RequestParam for options."

**Indepth:**
> **Encoding**: URL Encoding. Path variables are part of the URL structure and must be URL-encoded if they contain special characters. Request params are standard query strings. Spring decodes them automatically, but clients must send them correctly.


---

**Q: Spring Security Filter Chain**
> "Security doesn't happen in the Controller. It happens at the door.
>
> Spring Security is a chain of 10-15 filters that sit *before* the DispatcherServlet.
> *   **JwtAuthenticationFilter**: 'Do you have a token?'
> *   **UsernamePasswordFilter**: 'Are you logging in?'
> *   **AuthorizationFilter**: 'Are you allowed to see this?'
>
> If you pass all filters, you get to the Controller. If one fails, you get thrown out (401/403)."

**Indepth:**
> **Debugging**: `logging.level.org.springframework.security=DEBUG`. This is the single most useful config for debugging 403s. It prints the execution of every filter in the chain so you can see exactly which one denied the request and why (e.g., "CsrfFilter denied request").


---

**Q: Authentication vs Authorization**
> "**Authentication (Who are you?)**: 'I am John.' Checks username/password.
>
> "**Authorization (What can you do?)**: 'John is an Admin.' Checks roles and permissions.
>
> You must authenticate *before* you can be authorized."

**Indepth:**
> **HTTP Codes**: 401 Unauthorized actually means "Unauthenticated" (I don't know who you are). 403 Forbidden means "Unauthorized" (I know who you are, but you can't do this). The naming is confusing but standard.


---

**Q: OAuth2 Flow (Simple Explanation)**
> "Think of 'Login with Google'.
> 1.  You click the button.
> 2.  You are redirected to Google's server.
> 3.  You sign in there. Google asks: 'Allow this app to see your email?'
> 4.  You say Yes. Google sends a 'Code' back to your App.
> 5.  Your App talks to Google silently: 'Here is the Code, give me an Access Token.'
> 6.  Now your app has a Token to fetch your email from Google."

**Indepth:**
> **PKCE**: Proof Key for Code Exchange. In modern mobile/SPA apps, the "Code" can be intercepted. PKCE adds a cryptographic hash/verifier to simpler flows to prevent authorization code injection attacks.


---

**Q: CSRF (Cross-Site Request Forgery)**
> "If you log into your bank, and then visit `evil.com`.
> `evil.com` tries to send a hidden form POST to `bank.com/transfer`.
> Since your browser automatically sends the Bank Cookies, the Bank thinks *you* did it.
>
> **The Fix**: The Bank expects a secret `CSRF-Token` in the form. `evil.com` doesn't know this token, so the request fails.
> We disable this for REST APIs because we use Headers (Authorization), not Cookies."

**Indepth:**
> **Safe Methods**: GET, HEAD, OPTIONS are considered "Safe" (Read-only). CSRF only protects unsafe methods (POST, PUT, DELETE). Browsers execute Safe methods without restrictions, so never change state (DB writes) in a GET request.


---

**Q: Reactive Programming (Backpressure)**
> "In traditional systems, if the Producer sends data too fast, the Consumer crashes (Out of Memory).
>
> **Backpressure** is the Consumer saying: 'I am overwhelmed! Stop sending!' or 'Send me only 5 items'.
> It allows the system to handle massive load gracefully without crashing."

**Indepth:**
> **Strategies**: `onBackpressureBuffer/Drop`. What if the consumer *cannot* keep up? You can choose to Buffer the extra items (risking OOM), Drop them (data loss), or Error out. Reactive streams force you to decide this failure mode upfront.


---

**Q: WebFlux vs MVC (Threading)**
> "**MVC**: One Thread per Request. Blocking. If you have 200 threads and 201 concurrent users, the last one waits.
>
> "**WebFlux**: Event Loop. Non-Blocking. One thread handles many requests. When a request waits for DB, the thread serves someone else. It scales to 10,000+ concurrent connections with very little hardware."

**Indepth:**
> **Context Switching**: WebFlux reduces context switching because threads don't block. However, cpu-bound tasks (heavy calculation) can freeze the Event Loop, stopping *all* requests. You must offload CPU-heavy work to a separate standard thread pool.


---

**Q: R2DBC**
> "JDBC is blocking. If you use JDBC in WebFlux, you kill the performance benefits.
>
> **R2DBC** (Reactive Relational Database Connectivity) is the new driver standard for SQL databases (Postgres, MySQL) that is fully non-blocking/reactive. It's strictly required for true WebFlux apps talking to SQL."

**Indepth:**
> **Maturity**: R2DBC is newer than JDBC. It lacks some mature features like robust caching, complex mapping, or stored procedure support compared to Hibernate/JPA.


---

**Q: Mono vs Flux**
> "They are the 'futures' of Reactive Java (Project Reactor).
> *   **Mono**: A promise for 0 or 1 result. (e.g., `findById`).
> *   **Flux**: A stream of 0 to N results. (e.g., `findAll` or a live stock ticker).
> You chain operators on them (`.map`, `.filter`) to define the pipeline."

**Indepth:**
> **Multicasting**: `share()`. By default, if two people subscribe to a Flux, the source (e.g., DB query) is executed *twice*. Using `.share()` or `.publish()` multicasts the result to all subscribers, executing the source only once.

