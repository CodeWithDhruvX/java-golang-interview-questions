## 🛡️ Security & Authentication (Questions 801-820)

### Question 801: How do you hash passwords securely in Go?

**Answer:**
Use `bcrypt` or `argon2` from `golang.org/x/crypto`.
Never use SHA256 (too fast/parallelizable).

```go
hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
// Store 'hash' in DB
err = bcrypt.CompareHashAndPassword(hash, []byte("password"))
```

---

### Question 802: How do you implement HMAC-based authentication in Go?

**Answer:**
HMAC (Hash-Based Message Authentication Code) verifies data integrity and authenticity.
Use `crypto/hmac` and `crypto/sha256`.

```go
h := hmac.New(sha256.New, []byte("secret-key"))
h.Write([]byte("message"))
signature := hex.EncodeToString(h.Sum(nil))
```
Client sends `message` + `signature`. Server recomputes and compares.

---

### Question 803: How do you use JWT securely in Go APIs?

**Answer:**
- **Sign:** Use HS256 (Symmetric) or RS256 (Asymmetric).
- **Validate:** Parse token, check `alg` header (ensure it's not "None"), check `exp` (expiration), and verify signature.
- **Library:** `github.com/golang-jwt/jwt/v5`.

---

### Question 804: How do you prevent SQL injection in Go?

**Answer:**
Use parameterized queries (`?` or `$1`).
Never use `fmt.Sprintf` to build SQL queries.
(See Question 381 for code).

---

### Question 805: How do you manage CSRF protection in a Go web app?

**Answer:**
Use a middleware (like `gorilla/csrf`).
It generates a token and sets it in a Cookie.
The frontend must read the Value and send it back in a Header (`X-CSRF-Token`).
Server validates the Header matches the Cookie.

---

### Question 806: How do you handle XSS prevention in Go templates?

**Answer:**
`html/template` automatically escapes content.
Be careful with `template.HTML` (bypass), `js` contexts (use `json` encoding to safely inject vars into scripts), and `href` (check protocol is not `javascript:`).

---

### Question 807: How do you implement OAuth 2.0 flows in Go?

**Answer:**
See Q503. Use `golang.org/x/oauth2`.
Standard flow: Redirect User -> User Approves -> Callback with Code -> Backend exchanges Code for Token.

---

### Question 808: How do you encrypt/decrypt sensitive data in Go?

**Answer:**
Use `crypto/aes` with **GCM** (Galois/Counter Mode) for Authenticated Encryption.
Need a Key (32 bytes for AES-256) and a unique Nonce (12 bytes) per encryption.

---

### Question 809: What’s the use of `crypto/rand` vs `math/rand`?

**Answer:**
See Q517.
`crypto/rand` reads from OS entropy source (`/dev/urandom`), suitable for security.
`math/rand` is a pseudo-random generator, predictable if you know the seed.

---

### Question 810: How do you manage TLS certs in Go servers?

**Answer:**
- **Self-Signed/Files:** `http.ListenAndServeTLS(":443", "cert.pem", "key.pem")`.
- **Auto (Let's Encrypt):** Use `golang.org/x/crypto/acme/autocert` manager. It automatically fetches/renews certs from Let's Encrypt at runtime.

---

### Question 811: How do you validate tokens in Go microservices?

**Answer:**
- **Stateful (Session):** Query Redis.
- **Stateless (JWT):** Verify Signature using Public Key (RS256). No DB call needed if key is cached.

---

### Question 812: How do you securely store API keys in Go apps?

**Answer:**
Store *Hashed* API keys in the database (SHA256 is fine here as API keys have high entropy, unlike passwords).
User sees Key once. App stores Hash.
On request, Server hashes header Key and compares with DB.

---

### Question 813: How do you create and validate secure cookies?

**Answer:**
Use `gorilla/securecookie`.
It handles serialization + Encryption (AES) + Signing (HMAC).
Prevents users from tampering with cookie data on the client side.

---

### Question 814: How do you implement role-based access control (RBAC) in Go?

**Answer:**
Middleware logic.
1.  Extract User from Header/Token.
2.  Lookup valid Permissions (DB/Memory).
3.  Check: `if !user.HasPermission("admin:write") { 403 }`.
Libraries like **Casbin** provide a standard policy engine.

---

### Question 815: How do you verify digital signatures in Go?

**Answer:**
(See Q512). Use `rsa.VerifyPKCS1v15` or `ecdsa.Verify` depending on the signing key type.

---

### Question 816: How do you generate a secure random token in Go?

**Answer:**
```go
b := make([]byte, 32)
_, err := rand.Read(b) // crypto/rand
token := base64.URLEncoding.EncodeToString(b)
```

---

### Question 817: How do you prevent replay attacks with Go?

**Answer:**
(See Q518).
Require Nonce + Timestamp in request. Cache Nonce for window duration.

---

### Question 818: How do you audit Go applications for security issues?

**Answer:**
CI/CD Pipeline steps:
1.  **govulncheck:** Check known CVEs in dependencies.
2.  **gosec:** Static analysis tool (checks for hardcoded credentials, unsafe SQL, weak crypto).

---

### Question 819: How do you apply security headers in Go HTTP servers?

**Answer:**
Middleware setting:
- `Strict-Transport-Security` (HSTS).
- `Content-Security-Policy` (CSP).
- `X-Frame-Options` (Deny).
- `X-Content-Type-Options` (nosniff).

---

### Question 820: How do you secure gRPC endpoints in Go?

**Answer:**
1.  Enable TLS.
2.  Per-RPC Credentials (like Interceptors, but strictly for Auth Data).
3.  RBAC Interceptor before handler execution.

---
