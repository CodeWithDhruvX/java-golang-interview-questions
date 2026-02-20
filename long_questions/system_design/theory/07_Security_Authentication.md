# 🟡 Security & Authentication — Questions 61–70

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Razorpay, Paytm, PhonePe, Zerodha, fintech companies, Amazon, Google — any company handling user data or payments

---

### 61. How to design a secure login system?
"A secure login system has several layers of defense. The core flow: user submits email + password → verify credentials → issue a session token or JWT.

For password handling, I never store plaintext passwords. I use **bcrypt** or **Argon2** with a per-user salt to hash passwords. These are intentionally slow algorithms — a 10-round bcrypt hash takes ~100ms, making brute-force attacks impractical even if the DB is compromised.

For the session token, I prefer **short-lived JWTs (15min) with refresh tokens (7 days)** stored in httpOnly cookies (not localStorage — that's vulnerable to XSS). I also add: MFA for sensitive operations, account lockout after N failed attempts, suspicious login detection (new device, new country), and rate limiting on the login endpoint."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Razorpay, Zerodha, Paytm, Amazon, any BFSI (Banking, Financial Services, Insurance) company

#### Indepth
Secure login defense in depth:
1. **Password storage:** `password_hash = bcrypt(password + per_user_salt, cost=12)`. bcrypt with cost 12 takes ~250ms on modern hardware → brute force at scale is infeasible.
2. **Timing attack prevention:** Use constant-time comparison for password hashes. `subtle.ConstantTimeCompare()` in Go. Don't early-return "password wrong" differently from "user not found" — both should take equal time.
3. **Brute force protection:** Redis-backed rate limiter per IP: `INCR attempts:{ip}; EXPIRE attempts:{ip} 3600`. After 10 failures, lockout + CAPTCHA.
4. **Credential stuffing protection:** Check submitted passwords against Have I Been Pwned (HIBP) API. Reject known breached passwords.
5. **Session fixation:** Generate a new session ID after successful login. Don't reuse the pre-login session.
6. **Secure cookie flags:** `Set-Cookie: session=...; HttpOnly; Secure; SameSite=Strict`. HttpOnly prevents JS access (XSS protection). Secure requires HTTPS. SameSite=Strict prevents CSRF.

---

### 62. What is OAuth 2.0?
"OAuth 2.0 is an **authorization framework** that lets users grant third-party applications access to their resources without sharing their credentials.

The classic example: 'Login with Google'. When you click that button, your app redirects to Google's auth server. You log in to Google directly (your app never sees your Google password). Google issues an **authorization code** back to your app. Your app exchanges the code for an **access token**. Your app uses the access token to call Google APIs on your behalf.

The four entities: **Resource Owner** (you), **Client** (the third-party app), **Authorization Server** (Google's auth server), **Resource Server** (Google's API you're accessing). The key security benefit: your credentials are only shared with the Authorization Server you trust — never with the Client app."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company implementing social login or building an API platform — Razorpay (merchant OAuth), Swiggy (Google/Facebook login), CRED

#### Indepth
OAuth 2.0 Grant Types (flows):
- **Authorization Code Flow (+ PKCE for SPAs):** Browser redirects to auth server → user logs in → auth code returned to app → backend exchanges code for access+refresh token. Safest. Used for web apps and mobile apps (PKCE prevents code interception attacks by apps that can't keep secrets).
- **Client Credentials Flow:** Service-to-service authentication, no user involved. Machine identity. Used for backend microservice auth or cron jobs.
- **Resource Owner Password Credentials (DEPRECATED):** App takes username+password directly and exchanges for token. Never use for third-party apps — defeats the purpose.
- **Implicit Flow (DEPRECATED):** Access token returned directly in redirect URL. Vulnerable to token leakage via browser history and referrer headers. Replaced by Auth Code + PKCE.

**Access Token vs Refresh Token:**
- Access token: Short-lived (15min), used for API calls
- Refresh token: Long-lived (30 days), stored securely, used to get new access tokens without re-authentication
- Refresh token rotation: Invalidate old refresh token and issue new one on each use. Stolen refresh token detected when old token is used.

---

### 63. What is JWT?
"JWT (JSON Web Token) is a **compact, self-contained token format** for transmitting user identity and claims between parties, typically signed with a secret (HMAC-SHA256) or private key (RSA/Ed25519).

A JWT has three parts — header, payload, signature — Base64-encoded and dot-separated. The payload contains claims: `{ userId: 123, email: 'user@example.com', role: 'admin', exp: 1714000000 }`. The signature is `HMAC(header + payload, secret_key)`.

The beauty: stateless authentication. My API server doesn't need a DB lookup to verify the JWT — it just verifies the signature locally using its secret key. This scales horizontally with zero state. The downside: JWTs can't be revoked before expiry (unless you maintain a token blacklist, which defeats statelessness)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Every company with an API — Razorpay, Swiggy, Meesho, Zerodha, CRED

#### Indepth
JWT security gotchas:

1. **Algorithm confusion attack:** Early JWT libraries accepted `{ "alg": "none" }` and skipped verification. Always whitelist allowed algorithms server-side.
2. **RS256 vs HS256:** HS256 uses a shared secret (must be kept on all API servers). RS256 uses public/private key pair — private key signs on auth server, public key verifies on API servers. RS256 is safer for distributed systems (public key can be freely shared).
3. **`alg: RS256 → HS256` attack:** If server accepts both, attacker takes the public key and creates HS256 token signed with public key (treating it as HMAC secret). Always fix one algorithm.
4. **Storage:** `localStorage` — vulnerable to XSS (any script can read it). `httpOnly cookie` — protected from JS, vulnerable to CSRF (use SameSite=Strict). Cookie + CSRF token is the gold standard.
5. **Sensitive data in payload:** JWT payload is Base64 encoded, not encrypted — it's readable by the client. Never put sensitive data (SSN, credit card) in JWT payload. Use JWE (JSON Web Encryption) if you must.

---

### 64. How do you store passwords securely?
"Never store plaintext passwords. Never store MD5 or SHA-256 hashes (they're fast → easy to brute force). Use a **slow, adaptive password hashing algorithm** specifically designed for passwords: **bcrypt**, **scrypt**, or **Argon2**.

My recommendation: **Argon2id** (winner of the Password Hashing Competition 2015). It's memory-hard (attackers need lots of RAM, making GPU-based cracking expensive) and time-hard. For existing systems using bcrypt, it's fine to keep using it — it's battle-tested.

The key parameters: **work factor/cost** controls how slow the hash is. As hardware gets faster, increase the work factor in your configuration — rehash passwords on next login."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Every company with user authentication — foundational security question

#### Indepth
Hashing algorithm comparison:
| Algorithm | Speed | Memory-Hard | Recommended |
|---|---|---|---|
| MD5 | Very fast (>1B/s) | No | ❌ Never |
| SHA-256 | Fast (>100M/s) | No | ❌ Never for passwords |
| bcrypt | Slow (~100-300ms) | No | ✅ Acceptable |
| scrypt | Slow + Memory-Hard | Yes | ✅ Good |
| Argon2id | Configurable | Yes | ✅ Best |

**Never roll your own crypto.** Use standard library functions: Go's `golang.org/x/crypto/bcrypt`, Node.js `bcryptjs`, Python `passlib.hash.argon2`.

**Pepper (optional):** A server-side secret appended to the password before hashing: `hash = bcrypt(password + pepper)`. The pepper is stored in the application config (not DB). If the DB is stolen but the config isn't, the pepper prevents offline cracking. Can't be per-user like salt; if compromised, all passwords need rehashing.

---

### 65. What is rate limiting?
"Rate limiting controls how many requests a client can make to an API within a time window — protecting the system from overload, abuse, and DDoS-style attacks.

I implement it at the API gateway for all routes, then sometimes also at individual service level for specific expensive operations. Common policies: 100 requests/minute per API key for free tier, 1000/minute for paid tier, 10/second burst limit to prevent sudden spikes.

When a client hits the limit, return HTTP 429 Too Many Requests with `Retry-After: 60` header so the client knows when to retry. This is better than a confusing 500 error."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay, Stripe, Swiggy, Zomato, Amazon (API Gateway pricing plans)

#### Indepth
Rate limiting at multiple scopes simultaneously:
- **Per IP:** Prevents DDoS and abusive bots. Coarser-grained.
- **Per API key/user:** Fair usage across customers. Finer-grained.
- **Per endpoint:** Expensive endpoints (ML inference, file upload) get stricter limits than cheap ones (GET product).
- **Per tenant:** In a multi-tenant SaaS, enforce per-customer quotas.

Distributed rate limiter (Redis implementation):
```lua
-- Lua script for atomic token bucket check
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local current = redis.call('INCR', key)
if current == 1 then
  redis.call('EXPIRE', key, window)
end
if current > limit then
  return 0  -- rate limited
end
return 1  -- allowed
```

Alternative: Redis module `redis-cell` implements leaky bucket with a single command `CL.THROTTLE user123 15 30 60 1` (max 30 requests per 60 seconds, burst 15). Most production systems use this.

---

### 66. How to secure APIs?
"Securing an API is a multi-layer defense: authentication, authorization, input validation, encryption, and monitoring.

**Authentication** — who is this caller? Use API keys for machine-to-machine, OAuth2 + JWT for user auth. Always verify tokens server-side.

**Authorization** — what can this caller do? After authentication, check the user's roles/permissions for the specific resource. Enforce at the business logic layer, not just the API gateway (defense in depth).

**Input validation** — validate and sanitize every input. Reject unexpected fields. Limit payload size. Don't trust anything from the client — ever.

**HTTPS everywhere** — all API traffic over TLS 1.2+. No exceptions. HTTP-only endpoints are unacceptable for any authenticated API."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Every API-first company — Razorpay, Stripe, Twilio, Amazon API Gateway team

#### Indepth
API Security checklist (OWASP API Security Top 10):
1. **Broken Object Level Authorization:** Always check that the authenticated user owns the resource they're accessing. `GET /orders/{orderId}` → verify `order.userId == authenticated.userId`. Most common API vulnerability.
2. **Authentication failures:** Use standard OAuth2/JWT. Don't roll custom auth schemes.
3. **Excessive data exposure:** Return only what the client needs. Don't return full DB row with internal fields. Use response DTOs.
4. **Rate limiting missing:** Already covered in Q65.
5. **SQL/NoSQL injection:** Use parameterized queries, never string concatenation in queries.
6. **Mass assignment:** Reject unknown fields (`strict: true` in JSON parsing). Don't blindly bind request body to DB model objects.
7. **Security misconfiguration:** Disable debug endpoints in production. Remove stack traces from error responses. Set security headers: `X-Content-Type-Options`, `X-Frame-Options`, `Strict-Transport-Security`.
8. **CORS configuration:** Only allow trusted origins. Don't use `Access-Control-Allow-Origin: *` for authenticated endpoints.

---

### 67. What is CORS?
"CORS (Cross-Origin Resource Sharing) is a browser security mechanism. By default, browsers block JavaScript from making requests to a different domain (origin) than the one that served the page. This is the **Same-Origin Policy**.

CORS lets the *server* declare which other origins are allowed to call it. The browser checks this before allowing the JavaScript call. The server responds with `Access-Control-Allow-Origin: https://myapp.com` — and the browser allows the call only if the requesting origin matches.

The confusion: CORS is a **browser enforcement mechanism**, not a server security mechanism. A curl request or Postman call ignores CORS entirely. CORS protects against malicious websites making unauthorized calls using a logged-in user's cookies — it doesn't protect against server-to-server calls."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** Frontend-heavy companies, APIs with web clients — Meesho, Swiggy, Razorpay (payment iframe CORS issues)

#### Indepth
CORS preflight flow:
1. Browser detects cross-origin request with non-simple method/headers
2. Browser sends `OPTIONS` preflight request:
   ```
   OPTIONS /api/payment HTTP/1.1
   Origin: https://myapp.com
   Access-Control-Request-Method: POST
   Access-Control-Request-Headers: Content-Type, Authorization
   ```
3. Server responds:
   ```
   Access-Control-Allow-Origin: https://myapp.com
   Access-Control-Allow-Methods: POST, GET, OPTIONS
   Access-Control-Allow-Headers: Content-Type, Authorization
   Access-Control-Max-Age: 3600  ← cache preflight for 1 hour
   ```
4. Browser sends actual POST request if preflight succeeded

**Security mistake:** `Access-Control-Allow-Origin: *` with `Access-Control-Allow-Credentials: true` is illegal — browsers reject it. Can't allow all origins AND send cookies. You must specify exact origins for credentialed requests.

**CORS in microservices:** Each service doesn't need CORS headers if traffic comes through an API gateway (same-origin for the browser). Only the API gateway layer needs CORS configured. Add CORS only at the service boundary that browsers call directly.

---

### 68. Explain SSL/TLS in web communication.
"SSL/TLS is the cryptographic protocol that secures communication between a client (browser) and server over the internet. Every HTTPS connection uses TLS.

The TLS handshake: client connects → server sends its **certificate** (signed by a trusted CA like DigiCert) → client verifies the certificate → they agree on a **cipher suite** → they exchange keys using **asymmetric encryption (RSA/ECDH)** to establish a shared **session key** → all subsequent communication uses this symmetric session key (AES) which is much faster than asymmetric.

I configure my services to require TLS 1.2+ (never SSLv3, TLS 1.0, or 1.1 — they have known vulnerabilities). I also configure **HSTS** (HTTP Strict Transport Security) so browsers always use HTTPS and never downgrade."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any backend role dealing with production security — Razorpay, PhonePe, Amazon, Zerodha

#### Indepth
TLS concepts that come up in interviews:
- **Certificate** = identity document. Contains: domain name, public key, issuer (CA), expiry, signature by CA's private key.
- **Certificate Authority (CA):** DigiCert, Let's Encrypt, Comodo. Browsers trust a hardcoded list of root CAs. If the CA signs your cert, browsers trust it. Let's Encrypt provides free automated certificates (used via Certbot).
- **mTLS (Mutual TLS):** Both client AND server present certificates and verify each other. Used for service-to-service auth in microservices (instead of API keys). Istio/Envoy implement mTLS transparently between all services.
- **Certificate Pinning:** App hardcodes the expected certificate or public key. Rejects connections with any other certificate — even a valid CA-signed one. Breaks if certificate rotates without updating the app. Used in high-security mobile apps (banking, payments).
- **TLS Termination:** Reverse proxy (Nginx, AWS ALB) decrypts TLS at the edge. Backend servers receive plain HTTP. Simplifies backend, centralizes certificate management. Traffic between LB and backend is plaintext — must be in a trusted private network.

---

### 69. What is cross-site request forgery (CSRF)?
"CSRF is an attack where a malicious website tricks a logged-in user's browser into making an unwanted request to another website where the user is authenticated.

The attack flow: you're logged into your banking site — cookies are set. You visit `evil.com`. `evil.com` has hidden HTML: `<img src='https://bank.com/transfer?to=attacker&amount=5000'>`. Your browser automatically sends the bank request *with your cookies* (same-site cookie inheritance). Bank thinks it's you and processes the transfer.

Defense: **SameSite cookies** (modern, effective — `SameSite=Strict` or `SameSite=Lax`) prevent cookies from being sent on cross-site requests. For older browsers: **CSRF token** — a unique per-session random token embedded in forms. Server validates the token; cross-site requests can't read this token (same-origin policy protects it)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Companies with web apps and user sessions — Razorpay, Paytm, any fintech/e-commerce

#### Indepth
CSRF defense options and when to use each:
- **SameSite=Strict cookie:** Browser never sends cookie on cross-site request. Breaks third-party embedding (iframes). Best for apps that don't need cross-site embedding.
- **SameSite=Lax cookie:** Browser sends cookie on top-level navigation cross-site (clicking a link) but not on form submissions or AJAX from other sites. Good balance for most apps.
- **Double Submit Cookie pattern:** Server sets a random CSRF token in both a cookie and requires it in a request header. Cross-site request can't add the correct header (cross-origin JS can't read cookies).
- **CSRF token in form hidden field:** Classic ASP.NET/Django CSRF protection. `<input type="hidden" name="csrf_token" value="...">`.
- **Custom request header:** Require a custom header (e.g., `X-Requested-With: XMLHttpRequest`) that CORS prevents cross-origin requests from adding.

**Modern APIs are naturally CSRF-resistant** if they use `Authorization: Bearer {JWT}` header instead of cookies — cross-site requests can't add custom headers (blocked by CORS browser enforcement).

---

### 70. What is cross-site scripting (XSS)?
"XSS is an attack where malicious scripts are injected into web pages viewed by other users. If my site displays user input without sanitizing it, an attacker can submit a `<script>` tag that runs in every other user's browser.

**Reflected XSS**: `https://mysite.com/search?q=<script>steal_cookies()</script>`. **Stored XSS**: attacker submits a comment containing `<script>`, and every user who views that page runs the script.

Defense: **output encoding** — encode all user-supplied data before inserting into HTML (`&lt;script&gt;` instead of `<script>`). Modern frameworks like React do this automatically — JSX values are HTML-encoded by default. The `Content-Security-Policy` header is a powerful defense-in-depth: it tells the browser to only execute scripts from trusted sources."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company with user-generated content — Meesho reviews, Swiggy comments, Flipkart Q&A

#### Indepth
XSS attack types:
- **Reflected XSS:** Script in URL parameter, server echoes it back in response. Phishing attack — tricked user clicks a malicious link.
- **Stored XSS:** Script persisted in DB (comment, profile bio). Executes for every user who views that page. Higher severity — no social engineering needed.
- **DOM-Based XSS:** Client-side JS reads from URL and writes to DOM without sanitization. No server involvement. `document.getElementById('out').innerHTML = location.hash.substring(1)` — attacker sends `#<img src=x onerror=steal()>`.

Defense layer:
1. **Input sanitization:** Reject or strip dangerous characters on input.
2. **Output encoding:** Context-aware encoding — HTML encoding for HTML context, JS encoding for JS context, URL encoding for URL context.
3. **Content-Security-Policy (CSP):** `Content-Security-Policy: default-src 'self'; script-src 'self' cdn.trusted.com`. Even if XSS payload is injected, CSP blocks it from executing (scripts must come from approved sources).
4. **HttpOnly cookies:** XSS can't read httpOnly cookies via `document.cookie` — limits damage of successful XSS.
5. **Use modern frameworks:** React, Angular, Vue automatically escape output. Avoid `innerHTML`, `dangerouslySetInnerHTML`, `eval()`, and direct DOM manipulation with user data.
