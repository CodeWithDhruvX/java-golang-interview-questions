# 🛡️ **801–820: Security & Authentication**

### 801. How do you implement HMAC-based authentication in Go?
"Shared Secret + Hash.
Client sends `Signature = HMAC-SHA256(Body + Timestamp, Secret)`.
Server recomputes Signature.
`mac := hmac.New(sha256.New, secret)`.
`mac.Write(body)`.
`if !hmac.Equal(mac.Sum(nil), clientSig) { return 401 }`.
This proves *integrity* and *authenticity*."

#### Indepth
**Signature Validation**. Never use `==` to compare signatures. It terminates early (at the first mismatching byte), allowing **Timing Attacks**. Always use `hmac.Equal(sig1, sig2)` (which calls `crypto/subtle.ConstantTimeCompare`). This takes the same amount of time regardless of whether the first byte matches or the last one does.

---

### 802. How do you use JWT securely in Go APIs?
"1.  **Strict Verification**: `jwt.Parse(token, func(t) { if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { return nil, "bad algo" } return secret, nil })`.
2.  **Short TTL**: 15 minutes max.
3.  **Audience Check**: Ensure `aud` claim matches my service.
4.  **HTTPS**: Never send JWT over HTTP."

#### Indepth
**Key Rotation**. What if your signing key is compromised? You must support Key Rotation. The JWT header has a `kid` (Key ID) field. Your server looks up the specific public key for that `kid`. This allows you to sign new tokens with Key B while still accepting old tokens signed with Key A until they expire.

---

### 803. How do you manage CSRF protection in a Go web app?
"I use the **Cookie-to-Header** pattern.
Middleware sets a random `_csrf` cookie (HttpOnly=False).
Frontend reads it and sends `X-CSRF-Token` header.
Middleware verifies Cookie == Header.
Since attacker.com cannot read my domain's cookies, they can't forge the header."

#### Indepth
**SameSite**. Modern browsers support `SameSite=Strict` or `Lax` cookies. This prevents the browser from sending the cookie on cross-site requests (e.g., from a link in an email). While this kills most CSRF attacks natively, the Double-Submit Cookie pattern is still recommended as a defense-in-depth strategy for older browsers.

---

### 804. How do you handle XSS prevention in Go templates?
"Go’s `html/template` does it automatically.
It contextually escapes.
`{{ .Input }}` inside `<script>var x = "{{ .Input }}"` becomes `\x22alert(1)\x22`.
I audit code for use of `template.HTML` (the 'unsafe' type) and ensure those inputs are sanitized with `bluemonday`."

#### Indepth
**CSP**. Content Security Policy (HTTP Header) is your safety net. `Content-Security-Policy: default-src 'self'; script-src 'self' https://trusted.cdn.com`. Even if an attacker injects `<script>alert(1)</script>`, the browser will refuse to execute it because inline scripts are blocked by default in strict CSP.

---

### 805. How do you implement OAuth 2.0 flows in Go?
"I use `golang.org/x/oauth2`.
Auth Code Flow:
1.  Redirect user to Provider (Google).
2.  Receive `code` in callback.
3.  Exchange `code` for `token` (Backchannel).
4.  Use `token` to fetch User Profile.
I store the `access_token` in a secure session, not LocalStorage."

#### Indepth
**PKCE**. If your client is a Mobile App or SPA (public client), you can't keep a `client_secret`. Use **PKCE** (Proof Key for Code Exchange). The client generates a random `code_verifier` and hashes it to a `code_challenge`. The IDP verifies them, ensuring the app that requested the code is the same one swapping it for a token.

---

### 806. How do you encrypt/decrypt sensitive data in Go?
"**AES-GCM**.
It provides Authenticated Encryption (Confidentiality + Integrity).
`block, _ := aes.NewCipher(key)`.
`gcm, _ := cipher.NewGCM(block)`.
`ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)`.
I prepend the `nonce` to the ciphertext so I can decrypt it later."

#### Indepth
**KMS**. Storing the encryption key (`key`) in your config file (or Env Var) is risky. Use a Key Management Service (AWS KMS / Vault). The `key` never leaves the hardware module. You send plaintext to KMS, it returns ciphertext. This separates "Permission to Encrypt" from "Possession of the Key".

---

### 807. What’s the use of `crypto/rand` vs `math/rand`?
"**crypto/rand**: Reads from OS (`/dev/urandom`). Unpredictable. Use for Keys, Salts, Nonces.
**math/rand**: Pseudo-random. Predictable. Use for simulations.
Using `math/rand` for a session ID allows an attacker to predict the next ID and hijack sessions."

#### Indepth
**Entropy Exhaustion**. Reading from `/dev/urandom` is non-blocking and suitable for almost all crypto. Reading from `/dev/random` can *block* if the system's entropy pool is empty (on old Linux kernels). Go's `crypto/rand` uses system-specific CSPRPGs (`GetRandom` on Windows, `urandom` on *nix) which are safe and non-blocking.

---

### 808. How do you manage TLS certs in Go servers?
"For public internet: **ACME / Let's Encrypt**.
`certManager := autocert.Manager{ Prompt: autocert.AcceptTOS, Cache: autocert.DirCache("certs") }`.
`server.TLSConfig.GetCertificate = certManager.GetCertificate`.
It automatically renews certs. Zero config HTTPS."

#### Indepth
**HSTS**. Once you have HTTPS, force it. Send `Strict-Transport-Security: max-age=63072000; includeSubDomains`. This tells the browser: "For the next 2 years, NEVER talk to this domain over HTTP". It prevents SSL-Stripping attacks where a Man-in-the-Middle downgrades the user to HTTP.

---

### 809. How do you validate tokens in Go microservices?
"If using JWT: Stateless validation (CPU heavy).
If using Opaque Tokens: Call the Auth Service (Network heavy).
Optimization: **Caching**.
I cache the validation result in Redis for 1 minute."

#### Indepth
**Token Introspection**. Standard OAuth2 (RFC 7662) defines an endpoint (`POST /introspect`) to validate opaque tokens. The Resource Server sends the token to the Auth Server. This allows the Auth Server to revoke tokens instantly (by saying "Active: false"), which is impossible with stateless JWTs without complex blacklists.

---

### 810. How do you securely store API keys in Go apps?
"**Hash them!**
Treat API Keys like passwords.
Client has `sk_live_123`.
DB has `argon2(sk_live_123)`.
Start of key (`sk_live_`) is identifiable, but the secret part is hashed.
This way, if my DB is leaked, attackers can't use the keys."

#### Indepth
**Secret Scanning**. Prefix your keys! Stripe uses `sk_live_...`. GitHub uses `ghp_...`. This allows regex-based "Secret Scanners" (like GitHub's own) to detect if you accidentally commit a key to a public repo and revoke it automatically within seconds. Random strings are undetectable.

---

### 811. How do you create and validate secure cookies?
"I use `gorilla/securecookie`.
It HMAC-signs the value: `s.Encode("session", value)`.
It encrypts the value: `block.Encrypt(...)`.
So the user sees gibberish. If they tamper with one bit, the signature check fails.
I always set `Secure`, `HttpOnly`, `SameSite=Lax`."

#### Indepth
**Cookie Limits**. Browsers limit cookies to 4KB. If you encrypt a large session struct, it might exceed this. `securecookie` will return an error. You must either keep the session small (IDs only) or use server-side sessions (Redis) and only store the Session ID in the cookie.

---

### 812. How do you implement role-based access control in Go?
"I use **Casbin** or a simple Middleware.
`func RequireRole(role string) Middleware`.
It checks `ctx.User.Role`.
For complex policies ('Can edit post if owner OR admin'), I express logic in OPA (Open Policy Agent) Rego policies."

#### Indepth
**ABAC**. RBAC (Roles) is essentially coarse-grained. ABAC (Attribute Based) is fine-grained. "User can edit Document if distinct(User.Dept, Doc.Dept) < 10 miles". This logic is hard to hardcode. OPA allows decoupling this Policy Logic from your Go Business Logic.

---

### 813. How do you generate a secure random token in Go?
"`b := make([]byte, 32)`.
`_, err := rand.Read(b)`.
`token := base64.URLEncoding.EncodeToString(b)`.
This gives 256 bits of entropy.
Collisions are effectively impossible."

#### Indepth
**URL Safety**. Standard Base64 uses `+` and `/`, which have special meanings in URLs. Use `base64.URLEncoding` (which uses `-` and `_`) or `hex.EncodeToString`. This prevents tokens from being mangled when passed as query parameters (`?token=a+b` might be interpreted as `a b` by some servers).

---

### 814. How do you prevent replay attacks with Go?
"Signatures must include a **Timestamp** and **Nonce**.
Server:
1.  Verify Signature.
2.  Verify `Now - Timestamp < 5 mins`.
3.  Verify `Nonce` is not in Redis (Set with 5 min TTL).
This ensures a captured request cannot be re-sent later."

#### Indepth
**JTI**. In JWTs, the `jti` (JWT ID) claim serves as a Nonce. You can blacklist a `jti` in Redis for the duration of its default validity window. If the server sees the same `jti` twice, it's a replay. This is critical for "One-Time Use" tokens like Password Reset links.

---

### 815. How do you audit Go applications for security issues?
"**Static Analysis**: `gosec ./...`. Checks for hardcoded credentials.
**Dependency Scan**: `govulncheck ./...`.
**Fuzzing**: `go test -fuzz`.
**Dynamic**: `OWASP ZAP` against the running staging API."

#### Indepth
**Govulncheck**. The new standard tool from the Go team. Unlike `dependabot` (which checks `go.mod` versions), `govulncheck` analyzes your *call graph*. If you import a vulnerable library `v1.2.0` but *never call the vulnerable function*, `govulncheck` won't flag it. This reduces false positives significantly.

---

### 816. How do you apply security headers in Go HTTP servers?
"Middleware!
`w.Header().Set("X-Content-Type-Options", "nosniff")`.
`w.Header().Set("X-Frame-Options", "DENY")`.
Libraries like `secure` do this automatically with sensible defaults."

#### Indepth
**Permissions-Policy**. Formerly "Feature-Policy". It allows you to disable browser features like Geolocation, Camera, or USB for your site. `Permissions-Policy: geolocation=(), camera=()`. This reduces the attack surface if a sub-component (like a compromised ad script) tries to access user hardware.

---

### 817. How do you secure gRPC endpoints in Go?
"1.  **TLS**: Encrypt transport.
2.  **Auth Interceptor**: Verify JWT / mTLS header.
3.  **RBAC Interceptor**: Check method options (`/UserService/DeleteUser` requires `ADMIN`)."

#### Indepth
**ALTS**. In a service mesh (like Istio), you assume the network is hostile. ALTS (Application Layer Transport Security) or mTLS ensures that Service A can only talk to Service B if both present valid certificates signed by the internal CA. Go's `credentials.NewTLS` handles the mTLS handshake seamlessly.

---

### 818. How do you handle secrets rotation in Go?
"I listen for `SIGHUP` or watch the file/vault.
When a rotation event occurs:
I acquire a Lock.
I fetch the new key.
I update the global `CurrentKey`.
I keep `OldKey` for a grace period (1 hour) to allow in-flight requests to finish."

#### Indepth
**Envelope Encryption**. For database fields, don't re-encrypt 1TB of data when the key rotates. Encrypt Data with a specific "Data Key" (DK). Encrypt the DK with a "Master Key" (MK). Store `Encrypted(DK) + EncryptedData` in DB. When rotating, you only re-encrypt the DKs (small), not the massive Data blobs.

---

### 819. How do you prevent brute force attacks in Go?
"**Rate Limiting** per IP on Login.
If `failures > 5` in 1 minute, block IP for 15 minutes.
I use Redis to track failures: `INCR login_fail:{ip}`.
I also return generic error messages ('Invalid user or password') to avoid username enumeration."

#### Indepth
**Device Factors**. IP blocking is tricky (NATs, VPNs). Better: Track "Trusted Devices". If a login comes from a new device (new User-Agent + IP geo), require 2FA even if the password is correct. This stops credential stuffing attacks where attackers use valid passwords dumped from other sites.

---

### 820. How do you mitigate Timing Attacks in Go?
"When comparing secrets, use `crypto/subtle.ConstantTimeCompare(a, b)`.
`if a == b` returns faster if the first byte differs. An attacker can measure this time to guess the secret.
Constant Time compare takes the same time regardless of content."

#### Indepth
**Double-HMAC Verification**. For even stronger protection against timing attacks when comparing a user-provided token against a stored hash: Calculate `HMAC(StoredSecret, UserInput)` and compare it to `HMAC(StoredSecret, RealSecret)`. This masks the timing of the comparison itself behind the constant-time HMAC calculation.
