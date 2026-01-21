## 🔐 Security in Golang (Questions 501-520)

### Question 501: How do you prevent SQL injection in Go?

**Answer:**
SQL Injection occurs when untrusted input is concatenated directly into a query string.
**Fix:** Always use **Parameterized Queries** (Prepared Statements). The database driver handles escaping.

```go
// BAD
db.Query("SELECT * FROM users WHERE name = '" + name + "'")

// GOOD
db.Query("SELECT * FROM users WHERE name = ?", name)
```
For dynamic queries (e.g., optional filters), use a query builder (like **Masterminds/squirrel**) rather than string concatenation.

---

### Question 502: How do you securely store user passwords in Go?

**Answer:**
Never store plaintext. Use a slow hashing algorithm with a salt.
**Standard:** `bcrypt` or `argon2`.

```go
import "golang.org/x/crypto/bcrypt"

func Hash(pwd string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14) // Cost 14
    return string(bytes), err
}

func Verify(hashedPwd, plainPwd string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
    return err == nil
}
```

---

### Question 503: How do you implement OAuth 2.0 in Go?

**Answer:**
Use the `golang.org/x/oauth2` library.
1.  **Config:** Set Client ID, Secret, Redirect URL, and Scopes.
2.  **Redirect:** Generate the URL to send the user to the provider (Google/GitHub).
3.  **Callback:** Exchange the `code` returned by the provider for an `Access Token`.

```go
conf := &oauth2.Config{...}
// Exchange code
token, err := conf.Exchange(ctx, code)
// Create client
client := conf.Client(ctx, token)
```

---

### Question 504: What is CSRF and how do you prevent it in Go web apps?

**Answer:**
**Cross-Site Request Forgery** forces a logged-in user to execute unwanted actions.
**Prevention:**
1.  **SameSite Cookies:** Set `http.SameSiteStrictMode` or `Lax` to prevent cookies from being sent on cross-site requests.
2.  **Anti-CSRF Tokens:** Using middleware (like `gorilla/csrf`). The server generates a token, injects it into the HTML form/Meta tag, and validates it on every POST request.

---

### Question 505: How do you use JWT securely in a Go backend?

**Answer:**
1.  **Signing:** Use a strong algorithm (HS256 with a long secret, or RS256 with private/public keys).
2.  **Storage:** Store JWTs in `HttpOnly` `Secure` cookies (not LocalStorage) to prevent XSS theft.
3.  **Validation:** Always verify the signature and the `exp` (expiration) claim.
4.  **Library:** `github.com/golang-jwt/jwt/v5`.

---

### Question 506: How do you validate and sanitize user input in Go?

**Answer:**
- **Validation:** Use `go-playground/validator`. define struct tags like `validate:"required,email,gte=18"`.
- **Sanitization:** Use `microcosm-cc/bluemonday` to strip dangerous HTML tags from text inputs to prevent XSS.

```go
p := bluemonday.UGCPolicy()
clean := p.Sanitize(userInput)
```

---

### Question 507: How do you set secure cookies in Go?

**Answer:**
Set the correct flags on `http.Cookie`.

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session",
    Value:    token,
    HttpOnly: true,                  // JavaScript cannot access
    Secure:   true,                  // HTTPS only
    SameSite: http.SameSiteLaxMode,  // CSRF protection
    Path:     "/",
})
```

---

### Question 508: How do you avoid path traversal vulnerabilities?

**Answer:**
Path traversal allows attackers to read files outside the target directory (e.g., `../../etc/passwd`).
**Fix:**
1.  Use `filepath.Clean()` to resolve `..`.
2.  Check if the resulting path still starts with the expected root directory.

```go
path := filepath.Join(baseDir, userInput)
if !strings.HasPrefix(path, baseDir) {
    return error("access denied")
}
```

---

### Question 509: How do you prevent XSS in Go HTML templates?

**Answer:**
Go's `html/template` package acts as a safe-buffer by default. It **Contextually Auto-Escapes** data.
- If you insert data into `<div>{{.Data}}</div>`, it HTML-escapes (`<` becomes `&lt;`).
- If you insert into `<script>var x = "{{.Data}}"</script>`, it JS-escapes.
**Warning:** Do not use `template.HTML` type unless you deliberately want to render raw, unescaped HTML.

---

### Question 510: How would you encrypt sensitive fields (PII) before storing in DB?

**Answer:**
Use **AES-GCM** (part of `crypto/cipher`).
1.  Generate a random **Nonce** (12 bytes).
2.  Use the Master Key + Nonce to encrypt the data.
3.  Store `Nonce + Ciphertext` in the DB.

```go
block, _ := aes.NewCipher(key)
gcm, _ := cipher.NewGCM(block)
nonce := make([]byte, gcm.NonceSize())
io.ReadFull(rand.Reader, nonce)
ciphertext := gcm.Seal(nonce, nonce, []byte("data"), nil)
```

---

### Question 511: How do you securely generate random strings or tokens?

**Answer:**
**NEVER** use `math/rand` (it is deterministic if seeded, or predictable).
**ALWAYS** use `crypto/rand`.

```go
import "crypto/rand"

b := make([]byte, 32)
rand.Read(b)
token := hex.EncodeToString(b)
```

---

### Question 512: How do you verify digital signatures in Go?

**Answer:**
If verifying a payload from a webhook (e.g., Stripe/crypto):
1.  Get the Public Key.
2.  Hash the received body (SHA256).
3.  Use the signature algorithm (RSA/ECDSA/Ed25519) to verify.

Example (Ed25519):
```go
valid := ed25519.Verify(pubKey, message, signature)
```

---

### Question 513: What are best practices for TLS config in Go HTTP servers?

**Answer:**
Default settings are okay, but for high security:
1.  **MinVersion:** `tls.VersionTLS12` or `13`.
2.  **CipherSuites:** Restrict to secure modern ciphers (ECDHE-ECDSA-AES128-GCM-SHA256, etc.).
3.  **CurvePreferences:** Prefer `X25519`, `CurveP256`.

```go
&http.Server{
    TLSConfig: &tls.Config{
        MinVersion: tls.VersionTLS13,
    },
}
```

---

### Question 514: How do you implement rate limiting in Go to avoid DDoS?

**Answer:**
Use `golang.org/x/time/rate`.
Wrap your router in a middleware.
For distributed systems (multiple server instances), local memory limiting isn't enough; use **Redis** (Generic Cell Rate Algorithm or simply Fixed Window counters) to track IP limits across the cluster.

---

### Question 515: How do you handle secrets in Go apps (Vault, env, etc.)?

**Answer:**
1.  **Env Vars:** Standard (`os.Getenv`). Loaded from Kubernetes Secrets.
2.  **Runtime Fetch:** Use Hashicorp Vault SDK to fetch secrets on startup (keep them in memory, never write to disk).
3.  **Cloud Secret Managers:** AWS Secrets Manager / GCP Secret Manager.
**Anti-Pattern:** Hardcoding secrets or committing `.env` files.

---

### Question 516: How do you perform mutual TLS (mTLS) authentication in Go?

**Answer:**
mTLS requires both Server and Client to present certificates.
**Server Side:**
```go
tlsConfig := &tls.Config{
    ClientAuth: tls.RequireAndVerifyClientCert,
    ClientCAs:  caCertPool, // Pool containing the CA that signed client certs
}
```
In `handler`, check `r.TLS.PeerCertificates` to identify the caller.

---

### Question 517: What is the difference between `crypto/rand` and `math/rand`?

**Answer:**
- **`math/rand`:** Pseudo-Random Number Generator (PRNG). Fast. Deterministic (seeded). Verification: Not secure. Use for simulations, games, shuffling.
- **`crypto/rand`:** Cryptographically Secure PRNG (CSPRNG). Slower. Reads from OS entropy (`/dev/urandom`). Unpredictable. Use for passwords, keys, tokens, nonces.

---

### Question 518: How do you prevent replay attacks using Go?

**Answer:**
A replay attack is when an attacker intercepts a valid request and resends it.
**Mitigation:**
1.  **Timestamps:** Require a timestamp in the request header/body. Reject if `Now - Timestamp > 5 mins`.
2.  **Nonce:** Require a unique Nonce. Store used nonces in Redis with a TTL (equal to the timestamp window). Reject if Nonce exists.

---

### Question 519: How do you build a secure authentication system in Go?

**Answer:**
Don't build it from scratch if possible. Use **Auth0**, **Firebase**, or **Keycloak**.
If you must:
1.  Use `bcrypt` for passwords.
2.  Use `Session` cookies (HttpOnly, Secure) or short-lived `JWT` + Refresh Tokens.
3.  Enforce MFA (TOTP) using `github.com/pquerna/otp`.
4.  Log security events (failed logins).

---

### Question 520: How do you scan Go projects for vulnerabilities?

**Answer:**
1.  **govulncheck:** The official Go tool.
    `go install golang.org/x/vuln/cmd/govulncheck@latest`
    `govulncheck ./...`
    It analyzes your code (call graph) to see if you actually call the vulnerable functions in your dependencies.
2.  **Trivy / Snyk:** Scans container images and go.mod files.

---
