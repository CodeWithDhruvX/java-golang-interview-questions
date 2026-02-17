# 🟢 Go Theory Questions: 501–520 Security in Golang

## 501. How do you prevent SQL injection in Go?

**Answer:**
We **Parameterized Queries**.

We never concatenate strings (`"SELECT * FROM users WHERE name = '" + name + "'"`).
Instead, we use placeholders (`?` or `$1`).
`db.Query("SELECT * FROM users WHERE name = ?", name)`.
The database driver treats the input strictly as data, not executable code. This is the only defense. We also use ORMs like GORM which handle this automatically, but we audit raw SQL usages strictly.

---

## 502. How do you securely store user passwords in Go?

**Answer:**
We use **Argon2id** or **Bcrypt**.

We never store plain text or simple MD5/SHA256 hashes (which are vulnerable to Rainbow Tables).
We use `golang.org/x/crypto/bcrypt`.
`hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)`.
To verify: `bcrypt.CompareHashAndPassword(hash, []byte(password))`.
Argon2id is technically superior (memory-hard, resisting GPU cracking), but Bcrypt is the industry standard and sufficiently secure for most web apps.

---

## 503. How do you implement OAuth 2.0 in Go?

**Answer:**
We use `golang.org/x/oauth2`.

It handles the redirect dance.
1.  **Config**: ClientID, Secret, RedirectURL, Scopes.
2.  **AuthCodeURL**: Generates the link for the user to click ("Log in with Google").
3.  **Exchange**: Swaps the returned `code` for an `AuthToken` (Access + Refresh Token).
We verify the `state` parameter to prevent CSRF during the flow and typically use a strictly typed library like `markbates/goth` to normalize user profiles across providers (Google, GitHub, FB).

---

## 504. What is CSRF and how do you prevent it in Go web apps?

**Answer:**
CSRF (Cross-Site Request Forgery) tricks a user's browser into sending a request to your server (with their cookies) without their consent.

Prevention: **Synchronizer Token Pattern**.
We use `gorilla/csrf`.
It generates a random `csrf_token` and sets it in a secure cookie.
It also inserts this token into every HTML form as a hidden field (`<input name="gorilla.csrf.Token" ...>`).
On POST, the middleware verifies the cookie matches the form value. If they mismatch (or the form token is missing), it rejects the request (403).

---

## 505. How do you use JWT securely in a Go backend?

**Answer:**
We use `golang-jwt/jwt` (v5).

Security Rules:
1.  **Signing**: Use `HS256` (Symmetric) or `RS256` (Asymmetric). Never allow `None` algo.
2.  **Expiration**: Short `exp` (15 mins). Use Refresh Tokens for long sessions.
3.  **Validation**: `token.Method.(*jwt.SigningMethodHMAC)` check in the parse function is mandatory to prevent "Algorithm Confusion" attacks.
4.  **Storage**: Ideally `HttpOnly; Secure` cookies, not `localStorage` (XSS vulnerable).

---

## 506. How do you validate and sanitize user input in Go?

**Answer:**
**Validation**: Is it valid? (Email format, age > 18). Use `go-playground/validator`.
**Sanitization**: Is it safe? (Removing `<script>` tags).

For HTML output, we use `bluemonday`.
`p := bluemonday.UGCPolicy(); clean := p.Sanitize(userInput)`.
This strips dangerous tags while allowing safe ones (`<b>`, `<i>`). We sanitize *on output*, not on input, to preserve the original data fidelity in the DB.

---

## 507. How do you set secure cookies in Go?

**Answer:**
`http.Cookie` fields are critical.

```go
&http.Cookie{
    Name:     "session",
    Value:    token,
    HttpOnly: true,  // JS cannot read it (prevents XSS theft)
    Secure:   true,  // Only sent over HTTPS
    SameSite: http.SameSiteStrictMode, // Blocks CSRF
    Path:     "/",
}
```
If we don't set `HttpOnly`, a single XSS vulnerability exposes all user sessions.

---

## 508. How do you avoid path traversal vulnerabilities?

**Answer:**
Path Traversal (`../../etc/passwd`) happens when joining user input to file paths.

Fix: `filepath.Clean()`.
But `Clean` is not enough. We must verify the result starts with the intended root.
```go
fullPath := filepath.Join(baseDir, userInput)
if !strings.HasPrefix(fullPath, baseDir) {
    return error("Hacking attempt")
}
```
Better yet, use `io/fs` (Go 1.16+) which sandboxes the filesystem access to a specific root directory automatically.

---

## 509. How do you prevent XSS in Go HTML templates?

**Answer:**
Go's `html/template` package is **Context Aware** and auto-escapes by default.

If you put `{{.UserInput}}` inside `<p>`, it HTML-escapes (`<` becomes `&lt;`).
If you put it inside `<script>`, it JS-escapes.
If you put it inside `href`, it URL-escapes.
**Danger**: Using `template.HTML` type. This marks the string as "Safe" and bypasses escaping. We only cast to `template.HTML` if the content has been strictly sanitized by `bluemonday`.

---

## 510. How would you encrypt sensitive fields before storing in DB?

**Answer:**
We use **AES-GCM** (Galois/Counter Mode).

It provides Confidentiality + Integrity (Authenticated Encryption).
Keys: We use a Master Key (KMS) to encrypt Data Keys (Envelope Encryption).
In Go: `cipher.NewGCM(block)`.
We store the `Nonce` (12 bytes) + `Ciphertext`. You *cannot* reuse the nonce with the same key, or the security catastrophic fails.

---

## 511. How do you securely generate random strings or tokens?

**Answer:**
**Never use `math/rand`** for security. It is deterministic.
Use `crypto/rand`.

`b := make([]byte, 32); rand.Read(b)`.
Then encode: `hex.EncodeToString(b)` or `base64.URLEncoding.EncodeToString(b)`.
This reads from `/dev/urandom` (OS entropy pool), ensuring the token is mathematically unguessable.

---

## 512. How do you verify digital signatures in Go?

**Answer:**
We mostly see **HMAC** (Webhooks) or **RSA/ECDSA** (JWT/Crypto).

**HMAC**: `hmac.Equal(generatedMAC, receivedMAC)`. Note: Use `Equal`, not `==`, to avoid Timing Attacks.
**RSA**: `rsa.VerifyPKCS1v15(pubKey, hashAlgo, hashedMsg, signature)`.
We must ensure we use the correct Public Key corresponding to the entity claiming to sign the message.

---

## 513. What are best practices for TLS config in Go HTTP servers?

**Answer:**
Defaults are okay, but for A+ rating:

1.  **MinVersion**: `tls.VersionTLS12` (Disable SSLv3, TLS 1.0/1.1).
2.  **CipherSuites**: Explicitly prefer ECDHE-RSA-AES256-GCM-SHA384 and families. (Go 1.17+ handles this well automatically).
3.  **CurvePreferences**: `curve.X25519`.
We pass this `&tls.Config{...}` to the `http.Server`.

---

## 514. How do you implement rate limiting in Go to avoid DDoS?

**Answer:**
We use the **Token Bucket** algorithm via `golang.org/x/time/rate`.

Middleware:
`limiter := rate.NewLimiter(1, 5)` (1 req/sec, burst 5).
In handler: `if !limiter.Allow() { http.Error(429) }`.
For distributed apps (multiple pods), in-memory limiters don't work (global rate could be N * limit). We must use **Redis** (Lua script) to track the bucket counts centrally.

---

## 515. How do you handle secrets in Go apps (Vault, env, etc.)?

**Answer:**
1.  **Env Vars**: Standard for 12-Factor apps. Computed by K8s Secrets.
2.  **Files**: Docker Swarm / K8s mount secrets at `/run/secrets/db_pass`.
3.  **Vault**: For dynamic secrets. We use the HashiCorp Vault SDK. The app authenticates (via K8s Service Account), asks Vault for "Database Creds", and Vault generates a generic username/password valid for 1 hour. This is the gold standard ("Ephemeral Credentials").

---

## 516. How do you perform mutual TLS authentication in Go?

**Answer:**
mTLS means the Server verifies the Client's certificate too.

In `tls.Config`:
`ClientAuth: tls.RequireAndVerifyClientCert`.
`ClientCAs: caCertPool`.

The server will reject the handshake if the client doesn't present a certificate signed by the trusted CA in `ClientCAs`. This is standard for internal microservice-to-microservice zero-trust networks.

---

## 517. What is the difference between `crypto/rand` and `math/rand`?

**Answer:**
`math/rand` is **Pseudo-Random** (PRNG).
It uses a seed (usually time). If you know the seed (and the algorithm), you can predict every future number. Fast, but insecure.

`crypto/rand` is **Cryptographically Secure** (CSPRNG).
It asks the OS kernel for entropy (thermal noise, interrupts). It blocks if entropy is low (rarely on modern OS). It is slower but unpredictable. Always use `crypto/rand` for IDs, tokens, salts, and keys.

---

## 518. How do you prevent replay attacks using Go?

**Answer:**
A Replay Attack: Attacker intercepts a valid `POST /transfer` request and sends it again 10 times.

Fix: **Nonces** and **Timestamps**.
The request must include a `timestamp` and a unique `nonce`.
The server:
1.  Checks `abs(now - timestamp) < 5 min`.
2.  Checks Redis: `SETNX nonce 1 EX 300` (Exists? Reject).
This ensures that even if intercepted, the request is valid only once and only for a short window.

---

## 519. How do you build a secure authentication system in Go?

**Answer:**
We don't "build" it; we wire it.

1.  **Identity**: OAuth2 (Google/GitHub) or `bcrypt` for local users.
2.  **Session**: HttpOnly Secure Cookies (stateful) or JWT (stateless).
3.  **MFA**: TOTP (`pquerna/otp`) integration (QR Code scanning).
4.  **Middleware**: `RequireAuth` handler that checks the session before every protected route.
We audit every piece. Rolling your own crypto or auth logic is the #1 source of CVEs.

---

## 520. How do you scan Go projects for vulnerabilities?

**Answer:**
We use **Govulncheck** (Go Vulnerability Database).
`govulncheck ./...`

It analyzes your `go.mod` AND your source code to see if you actually *call* the vulnerable functions.
We also use **Snyk** or **Trivy** in CI/CD pipeline.
These tools block the build if we import a library with a known Critical CVE (Common Vulnerabilities and Exposures).
