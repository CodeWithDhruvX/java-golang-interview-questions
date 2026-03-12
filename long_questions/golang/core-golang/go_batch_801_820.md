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

### Explanation
Secure password hashing in Go uses bcrypt or argon2 from golang.org/x/crypto. These algorithms are intentionally slow and memory-hard to resist brute force attacks. SHA256 should never be used for passwords as it's too fast and parallelizable, making it vulnerable to GPU-based attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you hash passwords securely in Go?
**Your Response:** "I hash passwords securely using bcrypt or argon2 from Go's crypto package. These algorithms are specifically designed for password hashing because they're intentionally slow and memory-intensive, which makes them resistant to brute force attacks. I never use SHA256 for passwords - it's designed to be fast and parallelizable, which attackers love for cracking passwords. With bcrypt, I call `GenerateFromPassword` with the user's password and a cost factor that determines how slow the hashing should be. I store the resulting hash in the database. When verifying passwords, I use `CompareHashAndPassword` which automatically handles the salt and cost factor. This approach ensures that even if my database is compromised, attackers can't easily recover the original passwords."

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

### Explanation
HMAC-based authentication in Go uses crypto/hmac with SHA256 to verify data integrity and authenticity. The client creates a signature using a secret key and message, sends both to the server, which recomputes the signature and compares to ensure the message hasn't been tampered with.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement HMAC-based authentication in Go?
**Your Response:** "I implement HMAC-based authentication using Go's `crypto/hmac` package with SHA256. HMAC provides both data integrity and authenticity verification. The client creates an HMAC by computing a hash of the message using a secret key, then sends both the message and signature to the server. The server recomputes the HMAC using the same secret key and compares the results. If they match, the server knows the message hasn't been tampered with and came from someone with the secret key. I use `hmac.New(sha256.New, []byte("secret-key"))` to create the HMAC instance, write the message data, and encode the result. This approach is great for API authentication because it's stateless and doesn't require storing session data."

---

### Question 803: How do you use JWT securely in Go APIs?

**Answer:**
- **Sign:** Use HS256 (Symmetric) or RS256 (Asymmetric).
- **Validate:** Parse token, check `alg` header (ensure it's not "None"), check `exp` (expiration), and verify signature.
- **Library:** `github.com/golang-jwt/jwt/v5`.

### Explanation
Secure JWT usage in Go involves signing with HS256 (symmetric) or RS256 (asymmetric) algorithms, validation by parsing tokens and checking algorithm header, expiration, and signature verification. The golang-jwt library provides comprehensive JWT handling with security best practices.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use JWT securely in Go APIs?
**Your Response:** "I use JWT securely by following several key practices. For signing, I choose between HS256 for symmetric encryption when both services share the same secret, or RS256 for asymmetric encryption when I need public/private key separation. When validating tokens, I always parse the token first, then check that the algorithm header isn't set to 'None' to prevent algorithm downgrade attacks, verify the expiration claim to prevent token reuse, and most importantly, verify the signature to ensure the token wasn't tampered with. I use the `github.com/golang-jwt/jwt/v5` library which handles most of these security concerns automatically. I also implement proper key management, token rotation, and ensure tokens contain minimal necessary data to reduce the impact of token compromise."

---

### Question 804: How do you prevent SQL injection in Go?

**Answer:**
Use parameterized queries (`?` or `$1`).
Never use `fmt.Sprintf` to build SQL queries.
(See Question 381 for code).

### Explanation
SQL injection prevention in Go uses parameterized queries with placeholders (?) or ($1) instead of string formatting. This ensures user input is treated as data rather than executable SQL, preventing malicious SQL injection attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent SQL injection in Go?
**Your Response:** "I prevent SQL injection by always using parameterized queries with placeholders like `?` or `$1` instead of building SQL strings with `fmt.Sprintf`. When I use parameterized queries, the database driver automatically escapes and properly handles user input, treating it as data rather than executable SQL code. This completely eliminates SQL injection vulnerabilities. For example, instead of concatenating user input into SQL strings, I use `db.Query('SELECT * FROM users WHERE email = ?', email)` where the email parameter is safely bound. I never trust user input enough to put it directly into SQL strings. This approach works with all major Go database drivers and is the industry standard for preventing SQL injection attacks."

---

### Question 805: How do you manage CSRF protection in a Go web app?

**Answer:**
Use a middleware (like `gorilla/csrf`).
It generates a token and sets it in a Cookie.
The frontend must read the Value and send it back in a Header (`X-CSRF-Token`).
Server validates the Header matches the Cookie.

### Explanation
CSRF protection in Go web apps uses middleware like gorilla/csrf that generates tokens, sets them in cookies, and requires the frontend to send them back in headers. The server validates that the header token matches the cookie token, preventing cross-site request forgery attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage CSRF protection in a Go web app?
**Your Response:** "I implement CSRF protection using middleware like `gorilla/csrf`. The middleware generates a unique token for each user session and sets it in a cookie. For any state-changing requests, the frontend must read this token from the cookie and send it back in a custom header like `X-CSRF-Token`. The server then validates that the token in the header matches the token in the cookie. This prevents CSRF attacks because malicious websites can't read the cookie from my domain, so they can't include the correct token in their requests. I apply this middleware to all POST, PUT, DELETE, and PATCH endpoints. The gorilla/csrf library handles all the complexity of token generation, validation, and automatic exemption for safe HTTP methods like GET and HEAD."

---

### Question 806: How do you handle XSS prevention in Go templates?

**Answer:**
`html/template` automatically escapes content.
Be careful with `template.HTML` (bypass), `js` contexts (use `json` encoding to safely inject vars into scripts), and `href` (check protocol is not `javascript:`).

### Explanation
XSS prevention in Go templates uses html/template which automatically escapes content. Developers must be careful with template.HTML which bypasses escaping, JavaScript contexts requiring JSON encoding, and href attributes needing protocol validation to prevent javascript: URLs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle XSS prevention in Go templates?
**Your Response:** "I prevent XSS attacks primarily by using Go's `html/template` package which automatically escapes all content by default. This means if user input contains HTML or JavaScript, it gets safely encoded rather than executed. However, I need to be careful with certain edge cases. I avoid using `template.HTML` which bypasses escaping unless I absolutely trust the content. For JavaScript contexts, I use JSON encoding to safely inject variables into scripts rather than concatenating strings. For href attributes, I validate that the protocol isn't `javascript:` to prevent protocol injection attacks. The key principle is to always treat user input as untrusted and let Go's template engine handle the safe rendering. This defense-in-depth approach protects against most XSS vectors while still allowing dynamic content."

---

### Question 807: How do you implement OAuth 2.0 flows in Go?

**Answer:**
See Q503. Use `golang.org/x/oauth2`.
Standard flow: Redirect User -> User Approves -> Callback with Code -> Backend exchanges Code for Token.

### Explanation
OAuth 2.0 flows in Go use golang.org/x/oauth2 library implementing the standard authorization flow: redirect users to authorization server, get user approval, receive callback with authorization code, then exchange code for access token on the backend.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement OAuth 2.0 flows in Go?
**Your Response:** "I implement OAuth 2.0 flows using Go's `golang.org/x/oauth2` library. The standard flow starts when I redirect the user to the authorization server with my client ID and requested scopes. The user approves the request, and the authorization server redirects back to my callback URL with an authorization code. My backend then exchanges this code for an access token by making a server-to-server request with my client secret. I store the access token and refresh token securely for future API calls. The oauth2 library handles all the protocol details like building the right URLs, managing token exchange, and handling refresh flows. I configure it with the appropriate endpoints and scopes for each OAuth provider I'm integrating with, whether it's Google, GitHub, or a custom OAuth server."

---

### Question 808: How do you encrypt/decrypt sensitive data in Go?

**Answer:**
Use `crypto/aes` with **GCM** (Galois/Counter Mode) for Authenticated Encryption.
Need a Key (32 bytes for AES-256) and a unique Nonce (12 bytes) per encryption.

### Explanation
Data encryption/decryption in Go uses crypto/aes with GCM mode for authenticated encryption. Requires a 32-byte key for AES-256 and a unique 12-byte nonce for each encryption operation, providing both confidentiality and integrity protection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you encrypt/decrypt sensitive data in Go?
**Your Response:** "I encrypt sensitive data using Go's `crypto/aes` package with GCM mode, which provides authenticated encryption. GCM is important because it not only encrypts the data but also ensures it hasn't been tampered with. I use a 32-byte key for AES-256 encryption and generate a unique 12-byte nonce for each encryption operation. The nonce must never be reused with the same key. When encrypting, I create a new GCM cipher, encrypt the data, and get back both the ciphertext and an authentication tag. When decrypting, I verify the tag first - if it fails, I know the data was tampered with. This approach is perfect for protecting sensitive data like personal information or secrets that need to be stored securely. I manage the encryption keys carefully, storing them separately from the encrypted data."

---

### Question 809: What’s the use of `crypto/rand` vs `math/rand`?

**Answer:**
See Q517.
`crypto/rand` reads from OS entropy source (`/dev/urandom`), suitable for security.
`math/rand` is a pseudo-random generator, predictable if you know the seed.

### Explanation
crypto/rand vs math/rand: crypto/rand reads from OS entropy sources like /dev/urandom for cryptographically secure random numbers, while math/rand is a predictable pseudo-random generator suitable only for non-security purposes where the seed is known.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the use of `crypto/rand` vs `math/rand`?
**Your Response:** "I use `crypto/rand` for any security-related purposes like generating tokens, keys, or nonces because it reads from the operating system's entropy source like `/dev/urandom` on Linux. This provides cryptographically secure random numbers that are unpredictable. I never use `math/rand` for security because it's a pseudo-random generator that's completely predictable if you know the seed. I only use `math/rand` for non-security purposes like simulations, testing, or generating reproducible results where I want the same sequence of 'random' numbers each time. The key difference is that crypto/rand provides true randomness suitable for cryptography, while math/rand provides deterministic pseudo-randomness suitable for algorithms and testing."

---

### Question 810: How do you manage TLS certs in Go servers?

**Answer:**
- **Self-Signed/Files:** `http.ListenAndServeTLS(":443", "cert.pem", "key.pem")`.
- **Auto (Let's Encrypt):** Use `golang.org/x/crypto/acme/autocert` manager. It automatically fetches/renews certs from Let's Encrypt at runtime.

### Explanation
TLS certificate management in Go servers includes self-signed certificates with ListenAndServeTLS for static files, or automatic Let's Encrypt certificates using autocert manager that fetches and renews certificates automatically at runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage TLS certs in Go servers?
**Your Response:** "I manage TLS certificates in Go servers in two main ways. For development or internal services, I use self-signed certificates with `http.ListenAndServeTLS` where I provide cert.pem and key.pem files. For production services, I use the `golang.org/x/crypto/acme/autocert` manager which automatically handles Let's Encrypt certificates. The autocert manager fetches certificates when needed and renews them before they expire, all at runtime without requiring server restarts. I configure it with my domain, cache directory for certificates, and optionally email for account registration. This approach eliminates manual certificate management and ensures certificates are always up to date. The autocert manager handles all the ACME protocol complexity, making it easy to implement automatic HTTPS for Go web services."

---

### Question 811: How do you validate tokens in Go microservices?

**Answer:**
- **Stateful (Session):** Query Redis.
- **Stateless (JWT):** Verify Signature using Public Key (RS256). No DB call needed if key is cached.

### Explanation
Token validation in Go microservices uses stateful sessions with Redis lookups, or stateless JWT verification using public keys. JWT validation with cached public keys eliminates database calls, improving performance and scalability.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you validate tokens in Go microservices?
**Your Response:** "I validate tokens in microservices using two approaches. For stateful sessions, I store session data in Redis and query it for each request to validate the token. For stateless JWT validation, I verify the signature using a public key when using RS256 asymmetric encryption. The JWT approach is more scalable because I don't need database calls - I just cache the public key locally and verify the token signature. I choose JWT for most microservices because it's faster and doesn't require shared state between services. However, for applications that need immediate token revocation, I might use Redis-based sessions where I can delete the session data to immediately invalidate tokens. The choice depends on whether I prioritize performance and scalability (JWT) or immediate revocation capabilities (Redis sessions)."

---

### Question 812: How do you securely store API keys in Go apps?

**Answer:**
Store *Hashed* API keys in the database (SHA256 is fine here as API keys have high entropy, unlike passwords).
User sees Key once. App stores Hash.
On request, Server hashes header Key and compares with DB.

### Explanation
Secure API key storage in Go apps hashes API keys in the database using SHA256. Users see the key once, the app stores only the hash, and validates requests by hashing the header key and comparing with the stored hash. API keys have high entropy making SHA256 sufficient.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you securely store API keys in Go apps?
**Your Response:** "I store API keys securely by hashing them in the database, similar to password hashing but using SHA256 instead of bcrypt. When I generate an API key for a user, I show it to them once and immediately hash it before storing. The hash is what I keep in my database. When a request comes in with an API key in the header, I hash that key using the same algorithm and compare it with the stored hash. API keys have high entropy, so SHA256 is sufficient - I don't need the computational cost of bcrypt like I do for passwords. This approach means that even if my database is compromised, attackers can't recover the actual API keys. I also implement rate limiting and expiration policies for API keys to enhance security. The user only gets to see the full key once during generation, after which it's their responsibility to save it securely."

---

### Question 813: How do you create and validate secure cookies?

**Answer:**
Use `gorilla/securecookie`.
It handles serialization + Encryption (AES) + Signing (HMAC).
Prevents users from tampering with cookie data on the client side.

### Explanation
Secure cookies in Go use gorilla/securecookie which handles serialization, AES encryption, and HMAC signing. This prevents users from tampering with cookie data on the client side and ensures confidentiality and integrity of cookie contents.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you create and validate secure cookies?
**Your Response:** "I create and validate secure cookies using the `gorilla/securecookie` package which provides comprehensive cookie security. It handles three critical aspects: serialization to convert Go data structures into cookie-compatible formats, AES encryption to keep the data confidential, and HMAC signing to ensure integrity. When I set a cookie, the library automatically encrypts and signs the data. When reading cookies back, it verifies the signature first - if it doesn't match, it rejects the cookie as tampered. Then it decrypts and deserializes the data. This prevents users from reading or modifying cookie data on the client side. I configure it with separate encryption and authentication keys, and I rotate these keys periodically. This approach is much more secure than regular cookies where users can easily read and modify the contents."

---

### Question 814: How do you implement role-based access control (RBAC) in Go?

**Answer:**
Middleware logic.
1.  Extract User from Header/Token.
2.  Lookup valid Permissions (DB/Memory).
3.  Check: `if !user.HasPermission("admin:write") { 403 }`.
Libraries like **Casbin** provide a standard policy engine.

### Explanation
RBAC in Go uses middleware that extracts users from headers/tokens, looks up permissions from database or memory, and checks authorization. Libraries like Casbin provide standard policy engines for complex authorization rules and role management.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement role-based access control (RBAC) in Go?
**Your Response:** "I implement RBAC using middleware that handles authorization checks. The middleware first extracts the user from the request header or JWT token. Then it looks up the user's permissions from either a database or in-memory cache. Finally, it checks if the user has the required permission for the requested action, like `if !user.HasPermission('admin:write') { return 403 }`. For simple applications, I might implement this logic myself with user roles and permissions stored in a database. For more complex scenarios, I use libraries like Casbin which provide a powerful policy engine that supports roles, permissions, and resource-based access control. Casbin allows me to define policies in a configuration file and handles all the authorization logic. The middleware approach keeps authorization logic centralized and consistent across all endpoints."

---

### Question 815: How do you verify digital signatures in Go?

**Answer:**
(See Q512). Use `rsa.VerifyPKCS1v15` or `ecdsa.Verify` depending on the signing key type.

### Explanation
Digital signature verification in Go uses rsa.VerifyPKCS1v15 for RSA signatures or ecdsa.Verify for ECDSA signatures, depending on the key type used for signing. This verifies the authenticity and integrity of digitally signed data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you verify digital signatures in Go?
**Your Response:** "I verify digital signatures using Go's crypto libraries, specifically `rsa.VerifyPKCS1v15` for RSA signatures or `ecdsa.Verify` for ECDSA signatures, depending on what type of key was used for signing. The verification process involves having the original data, the signature, and the public key. I call the appropriate verification function which returns whether the signature is valid. This ensures that the data hasn't been tampered with and was indeed signed by the holder of the private key. I use this for things like verifying JWT signatures, validating webhook callbacks, or checking the integrity of software updates. The choice between RSA and ECDSA depends on the use case - ECDSA is more efficient but RSA is more widely supported. Both provide strong cryptographic guarantees when implemented correctly."

---

### Question 816: How do you generate a secure random token in Go?

**Answer:**
```go
b := make([]byte, 32)
_, err := rand.Read(b) // crypto/rand
token := base64.URLEncoding.EncodeToString(b)
```

### Explanation
Secure random token generation in Go uses crypto/rand to read random bytes, then base64 URL encoding to create URL-safe tokens. This ensures cryptographically secure randomness suitable for session tokens, API keys, or other security-sensitive identifiers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you generate a secure random token in Go?
**Your Response:** "I generate secure random tokens using Go's `crypto/rand` package. I create a byte slice of the desired length, typically 32 bytes for 256 bits of entropy, then call `rand.Read(b)` to fill it with cryptographically secure random data from the operating system's entropy source. Finally, I encode the bytes using `base64.URLEncoding.EncodeToString(b)` to create a URL-safe string token. This approach produces tokens that are unpredictable and suitable for security purposes like session identifiers, API keys, or password reset tokens. I always use `crypto/rand` instead of `math/rand` for security-sensitive tokens because it provides true randomness rather than pseudo-random numbers. The base64 URL encoding ensures the tokens are safe to use in URLs without encoding issues."

---

### Question 817: How do you prevent replay attacks with Go?

**Answer:**
(See Q518).
Require Nonce + Timestamp in request. Cache Nonce for window duration.

### Explanation
Replay attack prevention in Go requires nonces and timestamps in requests, with nonces cached for a time window. This ensures each request can only be used once and within a valid time period, preventing attackers from replaying captured legitimate requests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent replay attacks with Go?
**Your Response:** "I prevent replay attacks by requiring both a nonce and timestamp in each request. The nonce is a unique value that must never be reused, and the timestamp ensures the request is fresh. I cache recently used nonces for a specific time window, typically a few minutes. When a request comes in, I check that the timestamp is within my acceptable time window and that the nonce hasn't been used before. If either check fails, I reject the request. This prevents attackers from capturing legitimate requests and replaying them later because the nonce will have already been used or the timestamp will be too old. I implement this using Redis or an in-memory cache with automatic expiration. The combination of uniqueness and time-sensitivity makes replay attacks practically impossible while still allowing legitimate requests to proceed normally."

---

### Question 818: How do you audit Go applications for security issues?

**Answer:**
CI/CD Pipeline steps:
1.  **govulncheck:** Check known CVEs in dependencies.
2.  **gosec:** Static analysis tool (checks for hardcoded credentials, unsafe SQL, weak crypto).

### Explanation
Security auditing in Go applications uses CI/CD pipeline tools: govulncheck to check known CVEs in dependencies, and gosec for static analysis of security issues like hardcoded credentials, unsafe SQL queries, and weak cryptographic usage.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you audit Go applications for security issues?
**Your Response:** "I audit Go applications for security issues by integrating automated tools into my CI/CD pipeline. First, I run `govulncheck` which scans my dependencies for known CVEs and vulnerabilities, alerting me if I'm using any packages with known security issues. Second, I run `gosec` which performs static analysis on my source code to catch security anti-patterns like hardcoded credentials, unsafe SQL queries that could lead to injection, weak cryptographic usage, or other security vulnerabilities. These tools run automatically on every pull request, preventing security issues from reaching production. I also configure them to fail the build if they find high-severity issues. This automated approach catches common security mistakes early and ensures my application follows security best practices throughout development."

---

### Question 819: How do you apply security headers in Go HTTP servers?

**Answer:**
Middleware setting:
- `Strict-Transport-Security` (HSTS).
- `Content-Security-Policy` (CSP).
- `X-Frame-Options` (Deny).
- `X-Content-Type-Options` (nosniff).

### Explanation
Security headers in Go HTTP servers use middleware to set headers like HSTS for HTTPS enforcement, CSP for content security policy, X-Frame-Options to prevent clickjacking, and X-Content-Type-Options to prevent MIME type sniffing attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you apply security headers in Go HTTP servers?
**Your Response:** "I apply security headers using middleware that sets important HTTP security headers. I set `Strict-Transport-Security` (HSTS) to enforce HTTPS connections and prevent protocol downgrade attacks. I add `Content-Security-Policy` (CSP) to control which resources can be loaded and prevent XSS attacks. I set `X-Frame-Options` to 'Deny' to prevent clickjacking attacks by stopping my pages from being embedded in iframes. I also set `X-Content-Type-Options` to 'nosniff' to prevent browsers from MIME-sniffing content away from the declared content type. I implement this as middleware so it applies to all responses automatically. These headers work together to create multiple layers of defense against common web attacks. The combination of these headers significantly improves the security posture of my web applications without requiring changes to the application logic."

---

### Question 820: How do you secure gRPC endpoints in Go?

**Answer:**
1.  Enable TLS.
2.  Per-RPC Credentials (like Interceptors, but strictly for Auth Data).
3.  RBAC Interceptor before handler execution.

### Explanation
Secure gRPC endpoints in Go require TLS encryption, per-RPC credentials for authentication data, and RBAC interceptors that run before handler execution to enforce authorization policies. This provides comprehensive security for gRPC services.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure gRPC endpoints in Go?
**Your Response:** "I secure gRPC endpoints with a multi-layered approach. First, I enable TLS to encrypt all communication between clients and servers, preventing eavesdropping and man-in-the-middle attacks. Second, I implement per-RPC credentials to handle authentication - this is similar to HTTP interceptors but specifically designed for gRPC's streaming model. The credentials middleware extracts and validates authentication data from each RPC call. Third, I add an RBAC interceptor that runs before the actual handler execution to check if the authenticated user has permission to access the specific RPC method. This interceptor can check roles, permissions, or resource ownership. The combination of TLS for transport security, per-RPC credentials for authentication, and RBAC interceptors for authorization provides comprehensive security for gRPC services. This approach is similar to web security but adapted for gRPC's binary protocol and streaming capabilities."

---
