# 🛡️ 05 — Advanced Security & Auth
> **Most Asked in Product-Based Companies** | 🛡️ Difficulty: Hard

---

## 🔑 Must-Know Topics
- OAuth 2.0 and OpenID Connect flows
- Advanced JWT strategies (Access vs Refresh tokens, Token invalidation)
- SSRF, CSRF, XSS, and SQL Injection prevention at scale
- Rate Limiting and DDoS Mitigation in Node
- Secure Cookie settings (`HttpOnly`, `Secure`, `SameSite`)

---

## ❓ Frequently Asked Questions

### Q1. Explain the concept of Access Tokens vs Refresh Tokens. Why are they used together?

**Answer:**
Using a single, long-lived JWT for authentication is bad practice. If stolen, the attacker has permanent access until the token expires. If you make it short-lived, the user has to constantly re-login. Let's solve this with two tokens.

1. **Access Token:**
   - Short lifespan (e.g., 15 minutes).
   - Sent with every API request (usually in the `Authorization: Bearer` header).
   - Contains user permissions/roles.
   - Because it's short-lived, if stolen, the damage is heavily limited.

2. **Refresh Token:**
   - Long lifespan (e.g., 7 days or 30 days).
   - Stored securely on the client (preferably in a highly secure `HttpOnly` cookie).
   - Used *only* to request a new Access Token from an auth server when the current Access Token expires.

**Flow:**
When the API returns a `401 Unauthorized` (Token Expired), the client sends the Refresh Token to a specific `/refresh` endpoint. The server validates it against a database, rotates it, and issues a new Access Token.
If a Refresh Token is stolen or compromised, the server admin simply deletes the token from the database, instantly logging the user out.

---

### Q2. How do you implement robust JWT Blacklisting/Invalidation?

**Answer:**
Because JWTs are stateless and their verification doesn't hit the database, "logging out" or invalidating a JWT before its expiration time is challenging. By design, any validly signed JWT is trusted until its `exp` time is reached.

**Strategies for Invalidation:**
1. **The Redis Denylist (Blacklist):**
   - When a user logs out, extract the `jti` (JWT ID) or the token signature.
   - Store it in Redis with a TTL (Time-To-Live) completely matching the token's remaining expiration time.
   - **Trade-off:** In your auth middleware, you must now check Redis on every single request. You lose pure statelessness, but gain instant revocation and high performance.

2. **Token Versioning (Database-side Revocation):**
   - Add a `tokenVersion` integer to the User table.
   - Include the `tokenVersion` inside the JWT payload.
   - On request, verify the JWT, query the DB for the user, and check if `jwt.tokenVersion === dbUser.tokenVersion`.
   - To revoke ALL tokens globally for a user (e.g., password change), simply increment the `tokenVersion` in the DB.

---

### Q3. What is CSRF and how do you prevent it in a Node.js/Express API?

**Answer:**
**Cross-Site Request Forgery (CSRF)** occurs when a malicious website causes a user's web browser to perform an unwanted action on a trusted site where the user is currently authenticated (via cookies).

If your Node.js API relies on HTTP Cookies for authentication, the browser will automatically include those cookies on requests to your API, regardless of which website originated the request.

**Prevention Techniques:**
1. **`SameSite` Cookie Attribute:** Setting `SameSite=Lax` or `SameSite=Strict` completely blocks the browser from sending the cookie in cross-site requests. This is the modern, primary defense.
   ```javascript
   res.cookie('token', token, { httpOnly: true, secure: true, sameSite: 'strict' });
   ```
2. **Anti-CSRF Tokens:** The server generates a random token and embeds it in the HTML page. For state-changing requests (POST/PUT), the frontend must read this token and attach it as a hidden field or custom HTTP header (e.g., `X-CSRF-Token`). Since the malicious site cannot read the token due to CORS/Same-Origin Policy, its forged requests fail.
   You can implement this in Express using the `csurf` middleware package (though it is largely deprecated in favor of `SameSite` cookies).
3. **Double Submit Cookie Pattern:** Send a random value in both a cookie and a custom request header. The server verifies they match.

---

### Q4. Prevent SSRF (Server-Side Request Forgery) in Node.js.

**Answer:**
**SSRF** is a vulnerability where an attacker forces the server to make HTTP requests to an arbitrary domain of the attacker's choosing. This is common when your API accepts URLs as input (e.g., to fetch image previews, webhooks).

The danger is the attacker can force your Node.js server to hit internal microservices, AWS Metadata APIs (`169.254.169.254`), or private databases behind the firewall.

**Prevention:**
1. **Strict Input Validation / Allowlisting:** Only allow specific target domains if the feature permits it. (e.g., checking against a safe URL list or Regex).
2. **DNS Resolution & Private IP Blocking:** Before making the HTTP request via Node.js (e.g., with `axios`), forcefully resolve the requested hostname to an IP address. Check that the IP is not a private/internal IP (like `10.x.x.x`, `127.x.x.x`, `169.x.x.x`). Reject if it is.
   - Libraries like `ssrf-req-filter` can automate this validation layer.
3. **Network Configurations:** Run the Node.js application in an isolated sub-network (VPC) where egress traffic to internal critical infrastructure is blocked via security groups.
