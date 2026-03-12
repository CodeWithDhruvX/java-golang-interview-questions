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

### Explanation
SQL injection is a security vulnerability where malicious SQL code is injected into queries through untrusted input. Parameterized queries prevent this by separating the SQL command from the data, ensuring that input values are treated as data rather than executable code. The database driver automatically handles proper escaping and quoting, eliminating injection risks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent SQL injection in Go?
**Your Response:** "I prevent SQL injection in Go by always using parameterized queries instead of string concatenation. When I write database queries, I use placeholders like '?' and pass the variables separately - the database driver handles all the escaping and quoting automatically. For example, instead of concatenating strings directly into the query, I use `db.Query('SELECT * FROM users WHERE name = ?', name)`. For complex dynamic queries with optional filters, I use query builders like squirrel instead of building SQL strings manually. This approach ensures that user input is always treated as data, never as executable code, completely eliminating SQL injection vulnerabilities."

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

### Explanation
Secure password storage requires using slow, computationally expensive hashing algorithms with built-in salts. bcrypt and argon2 are industry standards that automatically handle salt generation and provide configurable cost factors to make brute-force attacks impractical. These algorithms are designed specifically for password hashing rather than general-purpose hashing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you securely store user passwords in Go?
**Your Response:** "I securely store user passwords in Go using the bcrypt package from golang.org/x/crypto. I never store plaintext passwords - instead, I hash them using bcrypt with a cost factor of 14, which makes the hashing computationally expensive enough to prevent brute-force attacks. The bcrypt.GenerateFromPassword function automatically handles salt generation, so each password gets a unique salt. For verification, I use bcrypt.CompareHashAndPassword to safely check if the provided password matches the stored hash. This approach ensures that even if the database is compromised, attackers can't recover the original passwords due to the computational cost and salt protection."

---

### Question 503: How do you implement OAuth 2.0 in Go?

**Answer:**
Use the `golang.org/x/oauth2` library.

```go
config := &oauth2.Config{
    ClientID:     "your-client-id",
    ClientSecret: "your-secret",
    Scopes:       []string{"email", "profile"},
    Endpoint:     google.Endpoint,
}

// Redirect user to provider
url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

// Exchange code for token
token, err := config.Exchange(context.Background(), code)
```

### Explanation
OAuth 2.0 implementation in Go uses the oauth2 library which handles the complex OAuth flow. The library manages client configuration, authorization URL generation, code exchange for access tokens, and token refresh. It abstracts away the HTTP communication and token management, providing a clean interface for integrating with OAuth providers like Google, GitHub, and others.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement OAuth 2.0 in Go?
**Your Response:** "I implement OAuth 2.0 in Go using the golang.org/x/oauth2 library which handles the entire OAuth flow. I create a config with client credentials, scopes, and the provider's endpoint. To start the flow, I generate an authorization URL using config.AuthCodeURL() and redirect the user there. When the user returns with an authorization code, I exchange it for an access token using config.Exchange(). The library handles all the HTTP requests, token management, and even refresh tokens automatically. This approach provides a clean, secure way to integrate with OAuth providers like Google or GitHub without having to manually implement the complex OAuth protocol details."
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

### Explanation
CSRF attacks exploit the trust a website has in a user's browser by forcing the user to execute unwanted actions on a site where they're authenticated. Prevention involves both cookie security settings and anti-CSRF tokens. SameSite cookies restrict when cookies are sent across different sites, while CSRF tokens ensure that requests originate from the legitimate application.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is CSRF and how do you prevent it in Go web apps?
**Your Response:** "CSRF, or Cross-Site Request Forgery, is an attack that forces logged-in users to execute unwanted actions on websites they trust. I prevent CSRF in Go web apps using two main approaches. First, I set SameSite cookies using `http.SameSiteStrictMode` or `Lax` to prevent cookies from being sent on cross-site requests. Second, I implement anti-CSRF tokens using middleware like `gorilla/csrf`. The server generates a unique token, injects it into HTML forms or meta tags, and validates it on every POST request. This ensures that requests actually come from my application rather than malicious third-party sites. The combination of these protections provides robust defense against CSRF attacks."

---

### Question 505: How do you use JWT securely in a Go backend?

**Answer:**
1.  **Signing:** Use a strong algorithm (HS256 with a long secret, or RS256 with private/public keys).
2.  **Storage:** Store JWTs in `HttpOnly` `Secure` cookies (not LocalStorage) to prevent XSS theft.
3.  **Validation:** Always verify the signature and the `exp` (expiration) claim.
4.  **Library:** `github.com/golang-jwt/jwt/v5`.

### Explanation
Secure JWT implementation requires strong signing algorithms, secure storage mechanisms, and proper validation. Use HS256 with long secrets or RS256 with key pairs for signing. Store tokens in HttpOnly Secure cookies rather than LocalStorage to prevent XSS attacks. Always verify signatures and check expiration claims to ensure token validity and prevent token abuse.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use JWT securely in a Go backend?
**Your Response:** "I use JWTs securely in Go backends by following several key practices. First, I use strong signing algorithms - either HS256 with a long, randomly generated secret or RS256 with private/public key pairs for better security. Second, I store JWTs in HttpOnly Secure cookies rather than LocalStorage to prevent theft through XSS attacks. Third, I always verify the signature and check the expiration claim on every request to ensure tokens are valid and not expired. I use the `github.com/golang-jwt/jwt/v5` library which provides robust JWT handling. This approach ensures that tokens can't be tampered with, can't be stolen through XSS, and expire appropriately, providing secure authentication for my applications."

---

### Question 506: How do you validate and sanitize user input in Go?

**Answer:**
- **Validation:** Use `go-playground/validator`. define struct tags like `validate:"required,email,gte=18"`.
- **Sanitization:** Use `microcosm-cc/bluemonday` to strip dangerous HTML tags from text inputs to prevent XSS.

```go
p := bluemonday.UGCPolicy()
clean := p.Sanitize(userInput)
```

### Explanation
Input validation and sanitization are crucial for security. The `go-playground/validator` library provides struct tag-based validation for common constraints like required fields, email format, and numeric ranges. For sanitization, `bluemonday` removes dangerous HTML tags and attributes from user input, preventing XSS attacks while allowing safe HTML content.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you validate and sanitize user input in Go?
**Your Response:** "I validate and sanitize user input in Go using specialized libraries. For validation, I use `go-playground/validator` which allows me to define constraints using struct tags like `validate:'required,email,gte=18'` for common validation rules. For sanitization, I use `microcosm-cc/bluemonday` to strip dangerous HTML tags from text inputs, preventing XSS attacks while allowing safe HTML content. I create a UGC policy with `bluemonday.UGCPolicy()` which defines what HTML is safe to keep, then use the `Sanitize()` method to clean user input. This two-step approach ensures that input is both structurally valid and safe from malicious content, protecting my application from both data integrity issues and security vulnerabilities."

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

### Explanation
Secure cookies in Go require setting specific flags to protect against various attacks. HttpOnly prevents JavaScript access, protecting against XSS. Secure ensures cookies are only sent over HTTPS. SameSite provides CSRF protection by controlling when cookies are sent cross-site. These flags work together to create a secure cookie implementation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you set secure cookies in Go?
**Your Response:** "I set secure cookies in Go by configuring the correct flags on the http.Cookie struct. I set HttpOnly to true to prevent JavaScript from accessing the cookie, which protects against XSS attacks. I set Secure to true to ensure the cookie is only sent over HTTPS connections, preventing interception. I set SameSite to LaxMode or StrictMode to provide CSRF protection by controlling when cookies are sent on cross-site requests. I also set the Path appropriately to limit where the cookie is sent. This combination of flags creates a robust security posture for session cookies, protecting against the most common web security vulnerabilities."

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

### Explanation
Path traversal vulnerabilities occur when user input is used to construct file paths without proper validation. Attackers can use `../` sequences to navigate outside the intended directory. Prevention involves cleaning paths to resolve directory traversal attempts and validating that the final path remains within the allowed directory scope.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you avoid path traversal vulnerabilities?
**Your Response:** "I prevent path traversal vulnerabilities by validating and sanitizing file paths constructed from user input. First, I use `filepath.Join()` to safely combine the base directory with user input, which handles path separator issues. Then I use `filepath.Clean()` to resolve any `..` sequences and normalize the path. Most importantly, I verify that the resulting path still starts with my intended base directory using `strings.HasPrefix()`. If the path tries to escape the base directory, I reject the request with an access denied error. This approach ensures that users can only access files within the designated directory, preventing attacks like `../../etc/passwd` that attempt to read sensitive system files."

---

### Question 509: How do you prevent XSS in Go HTML templates?

**Answer:**
Go's `html/template` package acts as a safe-buffer by default. It **Contextually Auto-Escapes** data.
- If you insert data into `<div>{{.Data}}</div>`, it HTML-escapes (`<` becomes `&lt;`).
- If you insert into `<script>var x = "{{.Data}}"</script>`, it JS-escapes.
**Warning:** Do not use `template.HTML` type unless you deliberately want to render raw, unescaped HTML.

### Explanation
Go's html/template package provides automatic context-aware escaping to prevent XSS attacks. It analyzes the context where data is inserted and applies appropriate escaping - HTML escaping for HTML content, JavaScript escaping for script blocks, CSS escaping for styles, etc. This ensures that user input is safely rendered without executing malicious code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent XSS in Go HTML templates?
**Your Response:** "I prevent XSS in Go HTML templates primarily by using the `html/template` package, which provides automatic context-aware escaping by default. When I insert data into HTML like `<div>{{.Data}}</div>`, it automatically HTML-escapes special characters, turning `<` into `&lt;`. If I insert into JavaScript blocks like `<script>var x = '{{.Data}}'</script>`, it applies JavaScript escaping. The template engine analyzes the context and applies the appropriate escaping type. I only use `template.HTML` when I deliberately need to render raw HTML that I trust. This built-in protection makes Go templates inherently safe against XSS attacks without requiring manual escaping in most cases."

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

### Explanation
AES-GCM provides authenticated encryption with additional data, ensuring both confidentiality and integrity. The algorithm requires a random nonce for each encryption operation. The nonce is stored alongside the ciphertext as it's needed for decryption. This approach protects sensitive PII data at rest while allowing secure retrieval when needed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you encrypt sensitive fields (PII) before storing in DB?
**Your Response:** "I encrypt sensitive PII fields before storing them in the database using AES-GCM, which provides authenticated encryption. I generate a random 12-byte nonce for each encryption operation, then use the master key combined with the nonce to encrypt the data. I store both the nonce and ciphertext together in the database since the nonce is required for decryption. The AES-GCM algorithm ensures both confidentiality and integrity of the encrypted data. This approach means that even if the database is compromised, the sensitive information remains protected without the encryption key. I use Go's crypto/cipher package which provides robust, well-vetted cryptographic primitives."

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

### Explanation
Cryptographically secure random number generation is essential for security-sensitive applications. `crypto/rand` provides unpredictable random numbers from the operating system's entropy sources, while `math/rand` produces deterministic sequences that are predictable and unsuitable for security purposes like token generation, password creation, or cryptographic operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you securely generate random strings or tokens?
**Your Response:** "I securely generate random strings and tokens in Go by always using `crypto/rand` and never `math/rand`. The `crypto/rand` package provides cryptographically secure random numbers from the operating system's entropy sources, making them unpredictable and suitable for security applications. I create a byte slice, fill it with random data using `rand.Read()`, then encode it as hex or base64 for human-readable tokens. I never use `math/rand` for security purposes because it's deterministic and predictable - even when seeded, it produces repeatable sequences. For session tokens, API keys, or any security-sensitive random data, `crypto/rand` is the only safe choice in Go."

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

### Explanation
Digital signature verification ensures the authenticity and integrity of received data. The process involves obtaining the public key, hashing the received payload, and using the appropriate cryptographic algorithm to verify that the signature matches the hash. This confirms that the data came from the expected source and hasn't been tampered with.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you verify digital signatures in Go?
**Your Response:** "I verify digital signatures in Go by following a standard cryptographic process. When receiving a webhook payload from services like Stripe, I first obtain the public key from the provider. Then I hash the received body using SHA256 to create a fixed-length digest. Finally, I use the appropriate signature verification algorithm - whether RSA, ECDSA, or Ed25519 - to verify that the signature matches the hash using the public key. For Ed25519, I simply call `ed25519.Verify(pubKey, message, signature)` which returns a boolean indicating validity. This process ensures that the webhook actually came from the expected provider and that the payload hasn't been modified in transit, providing both authentication and integrity guarantees."

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

### Explanation
TLS configuration best practices involve using modern protocol versions and secure cipher suites. Setting minimum TLS version to 1.2 or preferably 1.3 ensures deprecated protocols are rejected. Restricting cipher suites to modern, secure options prevents downgrade attacks. Preferring modern elliptic curves like X25519 and P256 provides better performance and security.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are best practices for TLS config in Go HTTP servers?
**Your Response:** "I configure TLS for high security in Go HTTP servers by following several best practices. First, I set the minimum TLS version to 1.2 or preferably 1.3 to exclude deprecated and insecure protocols. Second, I restrict cipher suites to modern, secure options like ECDHE-ECDSA-AES128-GCM-SHA256, preventing downgrade attacks to weaker encryption. Third, I set curve preferences to use modern elliptic curves like X25519 and P256 which provide better security and performance. While Go's default settings are reasonable, for high-security applications I explicitly configure these settings to ensure the strongest possible TLS configuration and protect against known vulnerabilities."

---

### Question 514: How do you implement rate limiting in Go to avoid DDoS?

**Answer:**
Use `golang.org/x/time/rate`.
Wrap your router in a middleware.
For distributed systems (multiple server instances), local memory limiting isn't enough; use **Redis** (Generic Cell Rate Algorithm or simply Fixed Window counters) to track IP limits across the cluster.

### Explanation
Rate limiting in Go uses the `golang.org/x/time/rate` package which implements token bucket algorithms. For single-server deployments, in-memory rate limiting is sufficient. For distributed systems across multiple server instances, Redis-based rate limiting ensures consistent enforcement across the cluster using algorithms like Generic Cell Rate Algorithm or fixed window counters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement rate limiting in Go to avoid DDoS?
**Your Response:** "I implement rate limiting in Go using the `golang.org/x/time/rate` package which provides token bucket rate limiting. I wrap my router in middleware that applies rate limits to incoming requests. For single-server applications, in-memory rate limiting works well. However, for distributed systems with multiple server instances, local memory limiting isn't sufficient because each server would have separate limits. In that case, I use Redis to track IP-based request counts across the entire cluster, implementing algorithms like Generic Cell Rate Algorithm or simple fixed window counters. This ensures consistent rate limiting enforcement regardless of which server instance handles the request, providing effective protection against DDoS attacks."

---

### Question 515: How do you handle secrets in Go apps (Vault, env, etc.)?

**Answer:**
1.  **Env Vars:** Standard (`os.Getenv`). Loaded from Kubernetes Secrets.
2.  **Runtime Fetch:** Use Hashicorp Vault SDK to fetch secrets on startup (keep them in memory, never write to disk).
3.  **Cloud Secret Managers:** AWS Secrets Manager / GCP Secret Manager.
**Anti-Pattern:** Hardcoding secrets or committing `.env` files.

### Explanation
Secrets management in Go involves multiple approaches. Environment variables loaded from Kubernetes Secrets provide basic secret management. Hashicorp Vault SDK allows runtime fetching of secrets with dynamic rotation. Cloud-specific secret managers offer integrated solutions. Critical anti-patterns include hardcoding secrets in code or committing .env files to version control.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle secrets in Go apps (Vault, env, etc.)?
**Your Response:** "I handle secrets in Go applications using multiple approaches depending on the requirements. For basic cases, I use environment variables loaded from Kubernetes Secrets using `os.Getenv()`. For more sophisticated secret management, I use the Hashicorp Vault SDK to fetch secrets at runtime and keep them in memory, never writing them to disk. For cloud deployments, I use AWS Secrets Manager or GCP Secret Manager which integrate well with cloud infrastructure. Critical anti-patterns I always avoid include hardcoding secrets in source code and committing `.env` files to version control. The choice depends on the deployment environment and security requirements - from simple env vars for development to comprehensive secret management systems for production."

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

### Explanation
Mutual TLS (mTLS) requires both client and server to present certificates for authentication. The server configures TLS to require and verify client certificates, using a certificate authority pool to validate client certs. Once established, the server can access the client's certificate information through request TLS peer certificates to identify and authorize the caller.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you perform mutual TLS (mTLS) authentication in Go?
**Your Response:** "I implement mutual TLS authentication in Go by configuring both client and server certificate requirements. On the server side, I create a TLS config with `ClientAuth: tls.RequireAndVerifyClientCert` to mandate client certificates, and provide a certificate authority pool that contains the CA that signed the client certificates. In my handlers, I can access the client's certificate information through `r.TLS.PeerCertificates` to identify and authorize the caller. This two-way authentication ensures that both parties trust each other's identities, creating a secure channel where the server knows exactly which client is connecting, and the client can verify the server's identity as well. It's particularly useful for zero-trust network architectures and service-to-service communication."

---

### Question 517: What is the difference between `crypto/rand` and `math/rand`?

**Answer:**
- **`math/rand`:** Pseudo-Random Number Generator (PRNG). Fast. Deterministic (seeded). Verification: Not secure. Use for simulations, games, shuffling.
- **`crypto/rand`:** Cryptographically Secure PRNG (CSPRNG). Slower. Reads from OS entropy (`/dev/urandom`). Unpredictable. Use for passwords, keys, tokens, nonces.

### Explanation
The key difference between `math/rand` and `crypto/rand` is their intended use case. `math/rand` is a pseudo-random number generator suitable for non-security applications like simulations or games, while `crypto/rand` provides cryptographically secure random numbers from operating system entropy sources, essential for security-sensitive operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between `crypto/rand` and `math/rand`?
**Your Response:** "The key difference is their security characteristics and intended use cases. `math/rand` is a pseudo-random number generator that's fast but deterministic - if you seed it with the same value, you get the same sequence. It's perfect for simulations, games, or shuffling cards where predictability isn't a security concern. `crypto/rand` is cryptographically secure, reading from the operating system's entropy sources like `/dev/urandom`. It's slower but unpredictable, making it essential for security-sensitive operations like generating passwords, encryption keys, tokens, or nonces. I never use `math/rand` for anything security-related - only `crypto/rand` provides the unpredictability needed for cryptographic applications."

---

### Question 518: How do you prevent replay attacks using Go?

**Answer:**
A replay attack is when an attacker intercepts a valid request and resends it.
**Mitigation:**
1.  **Timestamps:** Require a timestamp in the request header/body. Reject if `Now - Timestamp > 5 mins`.
2.  **Nonce:** Require a unique Nonce. Store used nonces in Redis with a TTL (equal to the timestamp window). Reject if Nonce exists.

### Explanation
Replay attacks involve capturing legitimate requests and resending them to cause unauthorized actions. Mitigation requires ensuring each request can only be used once. Timestamps ensure requests expire after a short time window, while nonces provide unique identifiers that prevent reuse. Combining both approaches with Redis storage provides robust protection against replay attacks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent replay attacks using Go?
**Your Response:** "I prevent replay attacks by implementing two key mitigations. First, I require timestamps in request headers or bodies and reject any request where the timestamp is older than a short window, typically 5 minutes. This ensures that even if an attacker captures a request, it becomes useless after a short period. Second, I require unique nonces for each request and store used nonces in Redis with a TTL matching the timestamp window. If a nonce already exists, I reject the request as a replay attempt. This combination of timestamps and nonces ensures that each request can only be used once and within a limited time frame, effectively preventing replay attacks while allowing legitimate requests to proceed normally."

---

### Question 519: How do you build a secure authentication system in Go?

**Answer:**
Don't build it from scratch if possible. Use **Auth0**, **Firebase**, or **Keycloak**.
If you must:
1.  Use `bcrypt` for passwords.
2.  Use `Session` cookies (HttpOnly, Secure) or short-lived `JWT` + Refresh Tokens.
3.  Enforce MFA (TOTP) using `github.com/pquerna/otp`.
4.  Log security events (failed logins).

### Explanation
Building secure authentication systems is complex and error-prone. Using established services like Auth0, Firebase, or Keycloak is preferred. If building from scratch is necessary, implement bcrypt for password hashing, secure session cookies or JWT with refresh tokens, multi-factor authentication, and comprehensive security logging. Each component addresses specific security concerns in the authentication pipeline.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a secure authentication system in Go?
**Your Response:** "I prefer using established authentication services like Auth0, Firebase, or Keycloak rather than building from scratch, as authentication is complex and security-critical. If I must build it myself, I implement several key components. I use bcrypt for secure password hashing, secure session cookies with HttpOnly and Secure flags, or short-lived JWTs with refresh token rotation. I enforce multi-factor authentication using TOTP libraries like `github.com/pquerna/otp`. I also implement comprehensive security logging to track failed login attempts and suspicious activities. Each component addresses specific security concerns - from protecting passwords to preventing session hijacking and ensuring strong user verification. Building authentication correctly requires expertise across multiple security domains."

---

### Question 520: How do you scan Go projects for vulnerabilities?

**Answer:**
1.  **govulncheck:** The official Go tool.
    `go install golang.org/x/vuln/cmd/govulncheck@latest`
    `govulncheck ./...`
    It analyzes your code (call graph) to see if you actually call the vulnerable functions in your dependencies.
2.  **Trivy / Snyk:** Scans container images and go.mod files.

### Explanation
Vulnerability scanning in Go involves multiple tools. `govulncheck` is the official Go tool that analyzes your code's call graph to determine if you actually use vulnerable functions in your dependencies, reducing false positives. Third-party tools like Trivy and Snyk provide additional scanning capabilities for container images and dependency files, offering comprehensive vulnerability detection.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you scan Go projects for vulnerabilities?
**Your Response:** "I scan Go projects for vulnerabilities using multiple tools for comprehensive coverage. First, I use `govulncheck`, which is the official Go vulnerability scanner. It analyzes my code's call graph to determine if I actually call vulnerable functions in my dependencies, which reduces false positives compared to simple dependency scanning. I install it with `go install golang.org/x/vuln/cmd/govulncheck@latest` and run `govulncheck ./...` on my codebase. Second, I use third-party tools like Trivy or Snyk which scan container images and go.mod files for additional coverage. This combination provides thorough vulnerability detection - from actual usage analysis to dependency scanning and container security. Regular scanning helps me identify and address security issues before they become problems in production."

---
