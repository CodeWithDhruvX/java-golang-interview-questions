# 🔐 **501–520: Security in Golang**

### 501. How do you prevent SQL injection in Go?
"I use **Parameterized Queries** consistently.
`db.Query("SELECT * FROM users WHERE name = $1", name)`.
The driver treats `$1` as data, escaping it immediately.
I *never* use `fmt.Sprintf` to build SQL strings. If I need dynamic columns (e.g., sorting), I verify against a strict allow-list: `allowedCols := map[string]bool{"age": true}`."

#### Indepth
Prepared Statements (`stmt, _ := db.Prepare(...)`) are parsed, compiled, and optimized by the DB server once. Repeated execution with different arguments is faster than sending raw SQL strings. They also strictly separate the control plane (SQL) from the data plane, rendering injection impossible.

---

### 502. How do you securely store user passwords in Go?
"**Argon2** or **Bcrypt**.
I prefer `golang.org/x/crypto/argon2` as it acts memory-hard, resisting GPU cracking.
`hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)`.
I store the salt and the hash. I never roll my own crypto or use simple SHA256."

#### Indepth
`Argon2` comes in two variants: `Argon2i` (optimized against side-channel attacks) and `Argon2d` (optimized against GPU cracking). The recommended variant for password hashing is **Argon2id**, which is a hybrid of both. Tune the parameters so hashing takes ~500ms on your server.

---

### 503. How do you implement OAuth 2.0 in Go?
"I use `golang.org/x/oauth2`.
It handles the heavy lifting: redirecting to provider, exchanging code for token, and refreshing tokens.
I configure the `oauth2.Config` with my ClientID/Secret.
`url := conf.AuthCodeURL(state)`.
I must validate the `state` parameter on callback to prevent CSRF attacks."

#### Indepth
The `state` token isn't just for CSRF; it can store tracking info (base64-encoded JSON) like where the user should be redirected after login (`return_to=/dashboard`). Always sign this state token with `HMAC` to prevent users from tampering with the redirection logic.

---

### 504. What is CSRF and how do you prevent it in Go web apps?
"**Cross-Site Request Forgery**.
Attacker tricks a user into clicking a link that posts to my bank API.
Prevention: **Double Submit Cookie**.
I use middleware (like `gorilla/csrf`). It injects a random token into the HTML form and checks for it in the POST header `X-CSRF-Token`. If they don't match -> 403 Forbidden."

#### Indepth
If your API is purely `JSON` and uses `Authorization: Bearer` headers (no Cookies), you generally don't need CSRF protection, because browsers don't automatically attach headers like they do with Cookies. CSRF is specifically an attack against browser-to-server Session Cookie authentication.

---

### 505. How do you use JWT securely in a Go backend?
"1.  **Alg**: Enforce `HS256` (HMAC) or `RS256` (RSA). Reject `None`.
2.  **Exp**: Set short expiration (15min). Use Refresh Tokens.
3.  **Library**: Use `golang-jwt/jwt/v5`.
4.  **Claims**: I strictly verify `iss` (Issuer) and `aud` (Audience).
I never store PII in the payload since it's just base64 encoded, not encrypted."

#### Indepth
The "None" Algorithm attack is a classic JWT vulnerability. Attackers modify the header to `{"alg": "none"}` and strip the signature. If your backend doesn't explicitly check `if token.Method != jwt.SigningMethodHS256`, the library might accept the unsigned token as valid!

---

### 506. How do you validate and sanitize user input in Go?
"I use **Strict Typing** and a validator library.
`type User struct { Email string 'validate:"email"' }`.
For HTML content (XSS prevention), I use **bluemonday** to strip dangerous tags (`<script>`).
I sanitize *on output*, not input, to preserve data integrity, but I validate structure heavily on input."

#### Indepth
Sanitization is tricky. `bluemonday` uses a whitelist approach (allow `<b>`, `<i>`, remove `<script>`), which is safer than blacklisting. Be careful with "Stored XSS": if you save malicious script to the DB and then render it in a PDF report or Admin Dashboard without escaping, you are still vulnerable.

---

### 507. How do you set secure cookies in Go?
"I set the flags on `http.Cookie`.
`HttpOnly: true` (JS can't read it).
`Secure: true` (HTTPS only).
`SameSite: http.SameSiteStrictMode` (Prevents CSRF).
For the value, I sign/encrypt it using `gorilla/securecookie` so users can't tamper with the session ID."

#### Indepth
Use `__Host-` or `__Secure-` prefixes for your cookie names. E.g., `__Host-SessionID`. Valid browsers will *reject* these cookies unless they are set with `Secure: true`, `Path=/`, and from a secure origin. This "Cookie Prefixing" is a defense-in-depth provided by the browser itself.

---

### 508. How do you avoid path traversal vulnerabilities?
"Attack: `GET /files?path=../../etc/passwd`.
Defense: `filepath.Clean(path)`.
After cleaning, I verify it starts with my expected root:
`if !strings.HasPrefix(cleanPath, rootDir) { return Error }`.
I also reject any path containing `..` explicitly before passing it to `os.Open`."

#### Indepth
Beware of **Null Byte Injection** in older systems (though Go is mostly safe). Also, on Windows, `Clean` might not catch alternate streams or UNC paths (`\\server\share`). Always resolve the final path with `filepath.EvalSymlinks` and check if it starts with the restricted root directory.

---

### 509. How do you prevent XSS in Go HTML templates?
"Go's `html/template` package is **Context-Aware**.
It automatically escapes data based on where it appears.
If I put `{{.Data}}` inside `<script>`, Go JSON-encodes it.
If inside `<div>`, it HTML-escapes it.
I avoid using `template.HTML` (which bypasses escaping) unless absolutely necessary and sanitized."

#### Indepth
Go's template engine is powerful but can be fooled if you inject into dangerous contexts. For example, injecting into `src="javascript:{{.}}"` or `onclick="{{.}}"` is risky. Content Security Policy (CSP) headers are your second line of defense if the template engine misses something.

---

### 510. How would you encrypt sensitive fields before storing in DB?
"**Envelope Encryption** (KMS).
Or locally: **AES-GCM** (`crypto/aes`, `crypto/cipher`).
`aes.NewCipher(key)`. `gcm.Seal(nonce, nonce, data, nil)`.
The `key` should not be hardcoded but loaded from a vault. I store the random `nonce` alongside the ciphertext. AES-GCM provides both confidentiality and integrity."

#### Indepth
Never reuse a **Nonce** with the same key in AES-GCM. If you do, the encryption breaks completely (XOR stream cipher). If you can't guarantee unique nonces (e.g., distributed systems), use **AES-GCM-SIV** (Synthetic IV), which is misuse-resistant.

---

### 511. How do you securely generate random strings or tokens?
"I use `crypto/rand`.
`b := make([]byte, 32); rand.Read(b)`.
`token := base64.URLEncoding.EncodeToString(b)`.
I **never** use `math/rand` for security tokens. It is deterministic (seeded). If an attacker knows the seed (often `time.Now().Unix()`), they can predict my next session ID."

#### Indepth
`crypto/rand` reads from `/dev/urandom` on Linux. In extremely early boot environments or container startups, entropy might be low, potentially blocking execution (though rare on modern kernels). `math/rand` is fine for retries (jitter) or load balancing, but never for anything security-critical.

---

### 512. How do you verify digital signatures in Go?
"Depends on the algorithm (RSA vs ECDSA).
For RSA: `rsa.VerifyPKCS1v15(pub, hashAlgo, hashed, sig)`.
I verify the hash of the data matches what was signed.
I use this for verifying webhooks (e.g., GitHub/Stripe signatures) to ensure the POST request actually came from them."

#### Indepth
When comparing signatures (e.g. `HMAC`), always use `subtle.ConstantTimeCompare(a, b)`. using `bytes.Equal` or `==` returns faster if the first byte mismatches, allowing an attacker to guess the signature byte-by-byte by measuring response time (Timing Attack).

---

### 513. What are best practices for TLS config in Go HTTP servers?
"Defaults are usually okay, but for hardening:
`MinVersion: tls.VersionTLS12`.
`CipherSuites`: explicit list of modern ciphers (ECDHE-RSA-AES256-GCM-SHA384).
`CurvePreferences`: `[]tls.CurveID{tls.X25519, tls.CurveP256}`.
I disable SSLv3 and TLS 1.0/1.1 explicitly to pass security audits."

#### Indepth
Go's default `tls.Config` is safe, but it aims for compatibility. For high security, restrict `MinVersion` to `TLS 1.3`. It removes weak ciphers, mandated perfect forward secrecy, and accelerates the handshake (1 RTT). `TLS 1.2` is the bare minimum today.

---

### 514. How do you implement rate limiting in Go to avoid DDoS?
"A layered approach.
1.  **Gateway**: Cloudflare/NGINX drops massive volumetric attacks.
2.  **App Middleware**: `tollbooth` or `golang.org/x/time/rate`.
I limit by IP. `limiter.Allow()` checks/decrements tokens. If empty -> 429.
Ideally, I use Redis for the counters so the limit applies across my whole cluster."

#### Indepth
Token Bucket vs Leaky Bucket. **Token Bucket** allows bursts (user can make 10 requests instantly, then 1/sec). **Leaky Bucket** smooths traffic (steady 1 req/sec). For APIs, Token Bucket feels snappier to users. For background processing queues, Leaky Bucket protects your database better.

---

### 515. How do you handle secrets in Go apps?
"**Environment Variables** are the standard.
But `os.Environ()` can leak in panic dumps.
For high security, I fetch from **HashiCorp Vault** or **AWS Secrets Manager** at startup and keep in memory.
I try to avoid keeping secrets in memory longer than necessary, but Go's GC makes 'wiping' memory unreliable."

#### Indepth
On Linux, you can use `mlock` (via `unix.Mlock`) to prevent sensitive memory pages from being swapped to disk. This mitigates the risk of an attacker reading your passwords from the hard drive's swap partition. HashiCorp Vault uses this trick.

---

### 516. How do you perform mutual TLS authentication in Go?
"Server side:
`tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert`.
`tlsConfig.ClientCAs = loadCA("trust-ca.pem")`.
When the handler runs, I check `r.TLS.PeerCertificates`.
If the client doesn't present a cert signed by my CA, the handshake fails at the TCP level. It’s perfect for service-to-service auth."

#### Indepth
Beware of **Certificate Revocation**. Just checking the signature isn't enough; the cert might have been stolen and revoked yesterday. Implementing `OCSP` stapling or checking a CRL (Certificate Revocation List) is required for a truly robust mTLS system.

---

### 517. What is the difference between `crypto/rand` and `math/rand`?
"**crypto/rand**: Cryptographically Secure PRNG. Reads from OS entropy (`/dev/urandom`). Slow. Use for keys, tokens, salts.
**math/rand**: Pseudo-Random. Deterministic PRNG. Fast. Use for simulations, testing, shuffling lists.
Confusion here is the #1 cause of predictable ID vulnerabilities."

#### Indepth
Go 1.22+ made `math/rand` global functions (like `rand.Intn`) automatically seeded with a random source! However, if you create a `New(NewSource(seed))`, it's still deterministic. Always audit your codebase for `math/rand` usage in token generation logic.

---

### 518. How do you prevent replay attacks using Go?
"I use a **Nonce** (Number used once) + **Timestamp**.
The client signs `{msg, nonce, timestamp}`.
Server checks:
1.  Timestamp is recent (within 5 min).
2.  Nonce hasn't been seen in Redis (set/check existence).
3.  Signature is valid.
This ensures an attacker can't just capture and re-send a valid 'transfer money' packet."

#### Indepth
Replay protection implies state (the used nonces). In a distributed system, this state must be shared (Redis). To save space, you only need to remember nonces that are within the allowed "Time Window" (e.g. 5 min). Older packets are rejected by Timestamp alone, so you can expire Redis keys after 5 min.

---

### 519. How do you build a secure authentication system in Go?
"I don't build it from scratch if I can avoid it. I use **Auth0** or **Gotrue**.
If I must:
1.  Bcrypt for storage.
2.  HttpOnly Secure cookies for sessions.
3.  CSRF tokens.
4.  Rate limiting on `/login`.
5.  Audit logging on failures.
The logic is simple, but the edge cases (orchestrating password reset safely) are where bugs hide."

#### Indepth
The "Password Reset" flow is sensitive. Don't say "Email sent" vs "User not found" (User Enumeration). Always say "If that email exists, we sent a link". Also, invalidate the existing session tokens when the password is changed to kick out any potential attackers.

---

### 520. How do you scan Go projects for vulnerabilities?
"I run **govulncheck ./...** in CI.
It parses my call graph. It alerts me only if I *call* a vulnerable function in a dependency, not just if I import it.
I also use `dependabot` to keep `go.mod` deps updated.
Static analysis with **gosec** catches code-level issues like hardcoded credentials."

#### Indepth
`govulncheck` is superior to generic dependence scanners because it uses **Call Graph Analysis**. If you import a vulnerable library version but *never call the vulnerable function*, `govulncheck` won't flag it (noise reduction). It focuses on *exploitable* vulnerabilities in your specific binary.
