# 🟢 Go Theory Questions: 801–820 Security & Authentication II

## 801. How do you use JWT securely in Go APIs?

**Answer:**
We use `golang-jwt/jwt`.
Key security steps:
1.  **Algorithm**: Enforce `HS256` or `RS256` in the parser. `jwt.Parse(token, func(t *jwt.Token) { if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { return nil, fmt.Errorf("bad alg") } })`.
2.  **Secret Management**: Load secret from ENV, never hardcode.
3.  **Claims**: Check `exp` (Expiry) and `iss` (Issuer) strictly.
4.  **Rotate**: Support key rotation by using a `KeyFunc` that looks up the current valid key ID (kid).

---

## 802. How do you manage CSRF protection in a Go web app?

**Answer:**
(See Q 504).
Middleware `gorilla/csrf`.
It injects a randomized token into headers/cookies.
For **Single Page Apps (SPA)**:
Cookie: `XSRF-TOKEN` (readable by JS).
Header: `X-XSRF-TOKEN` (sent by JS).
The backend validates that the Header == Cookie. The browser's Same-Origin Policy forbids attacker sites from reading the cookie, so they can't construct the header.

---

## 803. How do you handle XSS prevention in Go templates?

**Answer:**
`html/template` is safer than `text/template`.
It uses **Context-Aware Escaping**.
`{{ .Var }}`.
If Var is `<script>alert(1)</script>`, Go renders it as `&lt;script&gt;...`.
**Risk**: `template.HTML`. Only use this type if the content was run through a strict sanitizer like **Bluemonday** to strip dangerous tags while keeping formatting (b, i, u).

---

## 804. How do you implement OAuth 2.0 flows in Go?

**Answer:**
Standards: `golang.org/x/oauth2`.
Flow: **Authorization Code Flow**.
1.  Redirect user to Provider (Google).
2.  User approves. Provider redirects back with `?code=xyz`.
3.  Go server swaps `code` for `Access Token` via backend channel (Client Secret).
4.  Use Access Token to fetch User Info.
We store the User Info in a session and issue our own App JWT.

---

## 805. How do you encrypt/decrypt sensitive data in Go?

**Answer:**
Standard: **AES-GCM**.
```go
block, _ := aes.NewCipher(key)
gcm, _ := cipher.NewGCM(block)
nonce := make([]byte, gcm.NonceSize())
io.ReadFull(rand.Reader, nonce)
ciphertext := gcm.Seal(nonce, nonce, data, nil)
```
The `Seal` function encrypts AND signs (Auth Tag). We prepend the Nonce to the ciphertext so we can decrypt it later (`Open`).

---

## 806. What’s the use of `crypto/rand` vs `math/rand`?

**Answer:**
`math/rand`: Seeded PRNG. Deterministic. Fast. Use for simulations, fuzzing, games.
`crypto/rand`: CSPRNG. Reads OS entropy (`/dev/urandom`). Slow. Blocking (theoretically).
**Security Rule**: ALWAYS use `crypto/rand` for anything related to Auth, Keys, Salts, or IDs. If an attacker predicts your "Random" session ID, they account takeover your users.

---

## 807. How do you manage TLS certs in Go servers?

**Answer:**
Development: `GenerateCert` (Self-signed).
Production: **Let's Encrypt** (ACME).
Go has `golang.org/x/crypto/acme/autocert`.
```go
m := &autocert.Manager{
    Cache:      autocert.DirCache("certs"),
    HostPolicy: autocert.HostWhitelist("example.com"),
}
s := &http.Server{TLSConfig: m.TLSConfig()}
```
This automatically negotiates, downloads, and renews SSL certificates from Let's Encrypt with zero manual intervention.

---

## 808. How do you validate tokens in Go microservices?

**Answer:**
Service A calls Service B.
Method 1: **Introspection**. B calls Identity Provider (IdP): "Is this token valid?". (Slow).
Method 2: **Local Validation** (JWKS).
B downloads the Public Keys (JWK Set) from IdP once (caches them).
B validates the JWT signature locally using the RSA Public Key. This is fast (0 latency) and stateless.

---

## 809. How do you securely store API keys in Go apps?

**Answer:**
We **Hash** them, just like passwords.
When we issue a key `sk_live_123`, we show it once to the user.
We store `sha256("sk_live_123")` in the DB.
On request, we hash the incoming key and compare.
This ensures that if our DB is leaked, the attacker cannot use the keys to call our API.

---

## 810. How do you create and validate secure cookies?

**Answer:**
Use `gorilla/securecookie`.
It handles **Encryption** (AES) and **Signing** (HMAC).
`s := securecookie.New(hashKey, blockKey)`
`s.Encode("session", value)` -> produce opaque string.
`s.Decode("session", cookieVal, &value)`.
This prevents users from tampering with cookie data (e.g., changing `admin=false` to `admin=true`) and hides the contents.

---

## 811. How do you implement role-based access control in Go?

**Answer:**
Middleware + Claims.
JWT contains `roles: ["audit", "viewer"]`.
Middleware `RequireRole("admin")`.
Checks: `if !contains(claims.Roles, "admin") { return 403 }`.
For complex policies (ABAC), we use **Open Policy Agent (OPA)** or **Casbin**.
`e.Enforce(sub, obj, act)` -> Casbin checks `policy.csv` to see if User can Edit Document.

---

## 812. How do you generate a secure random token in Go?

**Answer:**
Bit of entropy.
1.  Read 32 bytes from `crypto/rand`.
2.  Encode to String.
`base64.RawURLEncoding.EncodeToString(bytes)`.
This gives a ~43 char URL-safe string.
Do not use `uuid.New()` (v4) for *secrets* (session tokens) as UUIDs are designed for uniqueness, not unpredictability (though v4 is random, a 32-byte CSPRNG string has more entropy).

---

## 813. How do you prevent replay attacks with Go?

**Answer:**
(See Q 518).
Require `timestamp` + `signature` headers.
Server checks:
1.  Signature is valid (Auth).
2.  Timestamp is fresh (< 5s old).
3.  Nonce (optional) hasn't been seen in Redis.
This prevents an attacker from grabbing a valid HTTP packet off the wire and resending it later to repeat an action (e.h. "Pay $50").

---

## 814. How do you audit Go applications for security issues?

**Answer:**
1.  **Static Analysis**: `gosec`.
    `gosec ./...` finds hardcoded credentials, weak crypto, unhandled errors.
2.  **Dependency Scan**: `govulncheck`. Finds CVEs in imported modules.
3.  **Fuzzing**: Go Fuzzing (1.18+) to crash parsers with garbage input.

---

## 815. How do you apply security headers in Go HTTP servers?

**Answer:**
Middleware.
```go
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-Frame-Options", "DENY")
w.Header().Set("Content-Security-Policy", "default-src 'self'")
```
We usually use a library like `unrolled/secure` which sets these standard hardening headers (HSTS, CSP, Referrer) automatically with safe defaults.

---

## 816. How do you secure gRPC endpoints in Go?

**Answer:**
(See Q 570 - Interceptors).
1.  **TLS**: Mandatory encryption.
2.  **Auth Interceptor**: Extract `authorization` metadata. Validate Token.
3.  **Auditing**: Log "Who did what" in the interceptor.
4.  **Rate Limiting**: Per-user limits to prevent DoS.

---

## 817. How do you mock HTTP clients in Go tests?

**Answer:**
We mock the **Transport**, not the Client.
```go
type MockTransport struct{}
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    return &http.Response{StatusCode: 200, Body: ...}, nil
}
client := &http.Client{Transport: &MockTransport{}}
```
This intercepts the network call at the lowest level, preventing any real socket connection.

---

## 818. How do you achieve high test coverage in Go?

**Answer:**
1.  **Table Driven Tests**: easy to add edge cases.
2.  **Interface Injection**: Mock dependencies.
3.  **Integration Tests**: For DB/API layers.
4.  **Gate**: Fail CI if coverage < 80%.
We focus on **Branch Coverage** (paths taken) rather than just line coverage.

---

## 819. How do you test race conditions in Go?

**Answer:**
`go test -race ./...`.
The race detector instruments memory accesses.
It crashes the test if two goroutines access the same variable concurrently and at least one is a Write.
It is **Runtime Detection** (dynamic), not static. It only finds races that actually happen during the test execution, so running the test 100 times (`-count=100`) increases confidence.

---

## 820. How do you benchmark functions in Go?

**Answer:**
(See Q 521).
`func BenchmarkX(b *testing.B)`.
Tips:
1.  **ResetTimer**: Call `b.ResetTimer()` after expensive setup.
2.  **StopTimer/StartTimer**: Pause during non-measured work.
3.  **RunParallel**: `b.RunParallel(func(pb *testing.PB) { ... })` to saturate all CPU cores and test contention.
Measurement: `ns/op` (Time) and `B/op` (Allocations).
