# 🟣 **261–280: Security and Best Practices**

### 261. How do you prevent injection attacks in Go?
"I rely on **Parameterized Queries** for SQL.

`db.Query("SELECT * FROM users WHERE name = ?", name)`.
The database driver treats `name` as data, not executable code.
For OS commands (`exec`), I strictly validate input against an allow-list (regex `^[a-zA-Z0-9]+$`). I never pass user input directly to a shell."

#### Indepth
For SQL, `?` placeholders only work for *values*, not identifiers (table/column names). If you need dynamic sorting (`ORDER BY ?`), the placeholder won't work. You must whitelist the allowed columns in Go code: `validSorts := map[string]bool{"created_at": true}` and check against that map before concatenating the string.

---

### 262. What are Go's common security vulnerabilities?
"Despite memory safety, Go apps have logic bugs.

1.  **Data Races**: Concurrent map writes crash the app (DoS).
2.  **Panics**: Uncaught panics crash the server.
3.  **Insecure Randomness**: Using `math/rand` for tokens (predictable) instead of `crypto/rand`.
4.  **Dependency Vulnerabilities**: Importing a malicious library (Supply Chain Attack)."

#### Indepth
A less obvious vulnerability in Go is **Directory Traversal** via `path/filepath.Join`. If a user supplies `../../etc/passwd`, a naive Join might resolve to a restricted file. Always clean the path and check if it starts with the expected root directory *after* resolution.

---

### 263. How do you hash passwords securely in Go?
"I use **bcrypt** (`golang.org/x/crypto/bcrypt`).

I call `bcrypt.GenerateFromPassword([]byte(password), cost)`.
It automatically handles **salting** (adding random data to prevent Rainbow Table attacks) and is slow by design (Key Stretching) to resist brute-force cracking.
I never use MD5 or SHA1 for passwords—they are broken."

#### Indepth
Argon2 (`golang.org/x/crypto/argon2`) is the newer winner of the Password Hashing Competition and is theoretically better than bcrypt because it's memory-hard (resists GPU cracking). However, bcrypt is still the industry standard for most web apps due to its maturity and ease of use in the Go ecosystem.

---

### 264. How to use `bcrypt` in Go?
"To hash (Signup):
`hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)`.

To verify (Login):
`err := bcrypt.CompareHashAndPassword(hash, pwd)`.
If `err == nil`, the password matches.
If `err == bcrypt.ErrMismatchedHashAndPassword`, it's wrong. I treat all other errors as internal server errors."

#### Indepth
Bcrypt has a **max length** of 72 bytes. If a user sends a 100-character password, bcrypt ignores the last 28 characters! To fix this, always hash the password with `sha256` first (which outputs 32 bytes) and then bcrypt the SHA256 hash. This supports passwords of any length.

---

### 265. How do you validate input in Go APIs?
"I use the **go-playground/validator** library.

I add struct tags:
`type User struct { Email string \`validate:"required,email"\` }`.
When binding JSON, I validate the struct. If it fails, I return a 400 Bad Request with a clear error message.
For critical security checks (e.g., 'is this user admin?'), I do manual checks in the service layer."

#### Indepth
Go validators are powerful but can be slow if overused (reflection). For high-performance hot paths, write custom validation logic (plain `if len(email) < 5`). Also, regex validation is vulnerable to **ReDoS** (Regular Expression Denial of Service). Ensure your regexes aren't exponential.

---

### 266. How do you implement JWT authentication?
"I use `golang-jwt/jwt/v5`.

Signing: `jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)`.
Verifying: `jwt.Parse(tokenString, func(token *jwt.Token) ...)`.
**Crucial Security Step**: Inside the parsing callback, I explicitly check `if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok`. This prevents the infamous 'Algorithm: None' attack where an attacker bypasses signature verification."

#### Indepth
JWTs are not encrypted; they are **signed**. Anyone can decode the payload (base64) and read the user's email or role. Never put sensitive data (like SSNs or passwords) inside a JWT claim. If you need privacy, use JWE (Encrypted JWT) or opaque session tokens.

---

### 267. How do you prevent race conditions in Go?
"I design for **Concurrency Isolation**.

I prefer channels to share data.
If I must share memory (e.g., a cache map), I protect it with `sync.RWMutex`.
I ensure *every* read uses `RLock()` and every write uses `Lock()`.
Most importantly, I run tests with `-race` in CI. If the race detector triggers, I fix it immediately—no exceptions."

#### Indepth
The race detector (`-race`) has limitations: it can only detect races that *actually happen* during the test run. It increases memory usage by 5-10x and CPU usage by 2-20x. It is NOT safe for production use unless you have a very specific, low-traffic debugging need.

---

### 268. What is CSRF and how do you mitigate it?
"Cross-Site Request Forgery. It tricks a user's browser into executing an action on their behalf.

I use the **gorilla/csrf** middleware.
It generates a random token creates a cookie.
It requires every state-changing request (POST/PUT/DELETE) to send that token in a header (`X-CSRF-Token`).
If the token is missing or invalid, the middleware rejects the request before it reaches my handler."

#### Indepth
Cookie attributes matter. Always set `HttpOnly` (prevents XSS from stealing the cookie) and `Secure` (HTTPS only). `SameSite=Strict` is the modern defense against CSRF, effectively stopping the browser from sending cookies on cross-site requests, making the custom token redundant in some cases.

---

### 269. How to use HTTPS in Go servers?
"For local dev: `http.ListenAndServeTLS(":443", "cert.pem", "key.pem")`.

For production exposed to the internet, I use **Let's Encrypt** via `golang.org/x/crypto/acme/autocert`.
`m := &autocert.Manager{Prompt: autocert.AcceptTOS}`.
It automatically fetches and renews certificates for my domain.
However, usually, I terminate TLS at the Load Balancer (AWS ALB) and run Go in plain HTTP."

#### Indepth
If you terminate TLS in Go, check your **Cipher Suites**. Defaults in older Go versions were permissive. In modern Go, the defaults are safe (TLS 1.2+), but you should explicitly disable older versions (`MinVersion: tls.VersionTLS13`) to meet compliance standards (PCI-DSS).

---

### 270. How do you sign and verify data in Go?
"I use **HMAC-SHA256** (Symmetric).

Signing: `mac := hmac.New(sha256.New, key); mac.Write(data); sum := mac.Sum(nil)`.
Verifying: I recompute the HMAC of the received data.
Then I use `hmac.Equal(sum, expectedSum)`.
**Why `Equal`?** Because standard `==` is vulnerable to **Timing Attacks**. `hmac.Equal` takes constant time regardless of how many bytes match."

#### Indepth
Timing attacks are subtle. If you compare `input == secret`, the CPU returns `false` as soon as the first byte differs. An attacker can measure the time difference to guess the secret byte-by-byte. `subtle.ConstantTimeCompare` (which `hmac.Equal` uses) ensures it always takes the same amount of time.

---

### 271. What are best practices for handling secrets in Go?
"**Never handle them.**
I strictly pass secrets via **Environment Variables**.

`dbPass := os.Getenv("DB_PASS")`.
In Kubernetes, these come from Secrets.
I try to avoid keeping secrets in memory longer than necessary, but Go's GC makes 'wiping' memory unreliable. The best defense is to make sure your process never dumps its memory to a log or crash report."

#### Indepth
A common mistake is unmarshalling a JSON config that contains secrets into a struct, and then printing that whole struct on startup errors (`fmt.Printf("Config loaded: %+v", cfg)`). This leaks API keys to the logs. Implement `String()` method on your config struct to redact sensitive fields.

---

### 272. How do you handle OAuth2 flows in Go?
"I use `golang.org/x/oauth2`.

It abstracts the handshake.
1.  Redirect user to Provider (Google/GitHub).
2.  User approves and returns with `code`.
3.  I exchange `code` for `Access Token`.
4.  I use the token to fetch user profile.
This library handles the token refreshing automatically, which is the hardest part of OAuth."

#### Indepth
The `state` parameter is mandatory for security. It prevents **CSRF** on the callback. You generate a random string, save it in a cookie, send it to Google, and when Google redirects back, you verify the `state` param matches the cookie. If not, the flow was initiated by an attacker.

---

### 273. How do you restrict file uploads (size/type)?
"Size: `http.MaxBytesReader(w, r.Body, 10<<20)`. This caps uploads at 10MB.

Type: I ignore the `Content-Type` header (it can be spoofed).
I read the first 512 bytes (sniffing) and use `http.DetectContentType(head)`.
If it says `image/png`, I trust it. If it says `application/octet-stream`, I reject it."

#### Indepth
Files can be **Polyglots**—valid GIF images that also contain valid JavaScript code (for XSS). If you serve user uploads, always force them to download (`Content-Disposition: attachment`) or serve them from a different domain (`user-content.com`) to prevent XSS on your main domain.

---

### 274. How do you set up CORS properly in Go?
"I use `rs/cors` library.

`c := cors.New(cors.Options{AllowedOrigins: []string{"https://myapp.com"}})`
I verify to set `AllowedMethods` (GET, POST).
If I need credentials (cookies), I must set `AllowCredentials: true` and I cannot use wildcard `*` for origins. I must list the exact domain."

#### Indepth
CORS preflight requests (OPTIONS) add latency. You can cache them in the browser by sending the `Access-Control-Max-Age` header (e.g., 24 hours). This significantly speeds up client-side apps by avoiding the preflight check on every single API call.

---

### 275. How do you scan Go code for vulnerabilities?
"I use **govulncheck**.

It’s the official Go security tool.
`govulncheck ./...`.
It analyzes my source code and dependency tree against the Go Vulnerability Database.
Unlike other tools, it only alerts if I *actually call* the vulnerable function, reducing false positives."

#### Indepth
`govulncheck` is superior to generic SBOM scanners (like Snyk or Dependabot) because of this "call graph analysis". If you import a massive library specifically for one safe function, but the library has a vulnerability in a different function you *don't* use, `govulncheck` won't spam you.

---

### 276. What is the Go ecosystem for SAST tools?
"**Static Application Security Testing**.

I use **gosec** (`securego/gosec`).
It scans my AST for:
*   Hardcoded credentials.
*   Weak crypto (MD5).
*   SQL injection risks.
*   Unsafe file permissions (0777).
I run it in CI/CD alongside the linter."

#### Indepth
Gosec can be noisy. It often flags `math/rand` (weak random) even when you are just shuffling a playlist (low risk). Use annotations `// #nosec G404` to suppress false positives, but always document *why* it's safe to ignore.

---

### 277. How to handle brute force protection in APIs?
"I implement **Rate Limiting**.

If a generic IP hits `/login` 10 times in 1 minute, I block it.
I use Redis to count attempts per IP.
For distributed attacks (botnets), application-layer limiting isn't enough; I rely on a WAF (Web Application Firewall) like Cloudflare or AWS WAF to drop traffic at the edge."

#### Indepth
For login endpoints, Rate Limiting is not enough; you need **Account Lockout** (after 5 failed attempts, lock for 15 mins). However, this allows an attacker to lock *you* out of your account by intentionally failing. The improved standard is "Exponential Backoff" or CAPTCHA after 3 failures.

---

### 278. How to secure communication between microservices?
"I use **mTLS** (Mutual TLS).

Every service has a sidecar (like Envoy in Isito) or internal logic to present a client certificate.
The server verifies the client certificate against an internal CA.
This ensures that only my 'Inventory Service' can talk to my 'Pricing Service'. A rogue binary on the network gets rejected instantly during the TLS handshake."

#### Indepth
Another benefit of mTLS is **identity**. The certificate Subject Name (CN) acts as the "User ID" of the service. You can write authorization policies: "Only certificates with `CN=billing-service` can access `POST /invoices`". This moves auth logic to the network layer.

---

### 279. What is the use of `context.Context` in secure APIs?
"It acts as the **Security Context**.

I store the authenticated User Identity in the context.
`ctx = context.WithValue(ctx, userKey, user)`.
My database layer or other services retrieve this `user` from context to enforce row-level security (e.g., 'User A can only edit their own profile')."

#### Indepth
Security context storage should be **Typed**, not string-based. Use a private struct key `type userKey struct{}` to prevent collisions. If you use string "user", a third-party library might accidentally overwrite your user value.

---

### 280. What is certificate pinning and can it be used in Go?
"Certificate Pinning protects against Compromised CAs.

In `tls.Config`, I use `VerifyPeerCertificate`.
I strictly check that the server's public key matches a hash I have hardcoded in my valid binary.
`if hash(cert.PublicKey) != pinnedHash { return error }`.
This prevents Man-in-the-Middle attacks even if the attacker has a valid certificate from a trusted authority."

#### Indepth
Certificate Pinning is brittle. If you rotate your certificate and forget to update the app binary, your app breaks for everyone. A safer middle ground is **Certificate Transparency** (CT) log monitoring or pinning the *Root CA* public key, not the leaf certificate.
