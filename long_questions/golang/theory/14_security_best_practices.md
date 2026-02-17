# 🟢 Go Theory Questions: 261–280 Security and Best Practices

## 261. How do you prevent injection attacks in Go?

**Answer:**
We prevent SQL and logic injection by separating data from commands, primarily using **Prepared Statements**.

In `database/sql`, instead of concatenating strings like `fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", input)`, we use placeholders: `db.Query("SELECT * FROM users WHERE name = ?", input)`. The database driver treats the input strictly as a value, never as executable code.

For other injections (like Command Injection), we avoid `os/exec` with user input. If we must use it, we validate the input against a strict allowlist (e.g., only alphanumeric characters) rather than trying to sanitize bad characters, which is error-prone.

---

## 262. What are Go's common security vulnerabilities?

**Answer:**
Go is memory-safe, so it ignores buffer overflows, but it's susceptible to **Concurrency Bugs** and **Improper Error Handling**.

A common vulnerability is **Goroutine Leaks**. If a goroutine is blocked forever on a channel, it eats memory, eventually causing a Denial of Service (DoS). Another is ignoring errors: `func() { _ = doCriticalSecCheck() }`. If the security check fails but you ignore the error, the door is left open.

Additionally, developers often misuse `unsafe` or `cgo` for performance, which re-introduces the memory corruption risks that Go was designed to avoid. Use of `unsafe` is a red flag during code reviews.

---

## 263. How do you hash passwords securely in Go?

**Answer:**
We never store plain text passwords. We use computationally expensive hashing algorithms like **Bcrypt** or **Argon2**, not fast hashes like SHA-256.

Fast hashes are designed for speed, which allows hackers to test billions of passwords per second. Bcrypt is "adaptive"—it has a work factor (cost) that we can tune. We set the cost (e.g., 10 or 12) so that hashing takes about 100ms.

This 100ms delay is imperceptible to a user logging in once, but it makes brute-forcing a stolen database mathematically impossible for an attacker.

---

## 264. How to use `bcrypt` in Go?

**Answer:**
We use the official sub-repo `golang.org/x/crypto/bcrypt`.

When a user signs up, we call `bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)`. We store this resulting hash in the database.

When a user logs in, we fetch the hash from the DB and call `err := bcrypt.CompareHashAndPassword(storedHash, []byte(submittedPassword))`. If `err` is nil, the password matches. We never decouple these steps or roll our own comparison logic, as `bcrypt` handles the salt and constant-time comparison automatically to prevent timing attacks.

---

## 265. How do you validate input in Go APIs?

**Answer:**
We use a **Parse, Validate, Normalize** strategy, often with a library like `go-playground/validator`.

We define struct tags: `type RegisterReq struct { Email string \`validate:"required,email"\` }`. When the JSON is decoded, the validator checks these rules immediately.

Beyond syntax, we do semantic validation. "Is this Age > 18?" "Does this ProductID actually exist in the DB?" rejecting invalid data as early as possible prevents "Garbage In, Garbage Out" and protects the database from corruption.

---

## 266. How do you implement JWT authentication?

**Answer:**
JWT (JSON Web Token) is a stateless auth mechanism. We use `golang-jwt/jwt`.

When a login is successful, we create a token with claims (UserID, Expiry) and sign it with a secret key: `token.SignedString(secretKey)`. The client sends this token in the `Authorization` header.

On every request, middleware parses the token and verifies the signature using the same secret. The critical security detail is to explicitly check the **Signing Method** (HMAC vs RSA) to prevent the "None Algorithm" vulnerability where an attacker strips the signature to bypass auth.

---

## 267. How do you prevent race conditions in Go?

**Answer:**
Race conditions generally happen when multiple goroutines write to the same map or variable without locking.

We prevent them by "Sharing memory by communicating" (using Channels) or protecting shared state with `sync.Mutex`.

To detect them, we strictly run our tests with `go test -race`. This enables the **Race Detector**, which instruments the code to flag unsynchronized memory access at runtime. We consider any race detected by this tool as a critical blocker for deployment.

---

## 268. What is CSRF and how to mitigate it?

**Answer:**
CSRF (Cross-Site Request Forgery) is where a malicious site tricks a user's browser into sending a request to your site using their saved cookies.

If you use **Stateful Sessions** (Cookies), you must use Anti-CSRF tokens. We use middleware (like `gorilla/csrf`) that injects a random token into every HTML form. The server rejects POST requests missing this token.

If you use **Stateless Auth** (JWTs in headers, not cookies), you are generally immune to CSRF because the browser doesn't automatically attach the Auth header like it does with cookies. This is why modern SPAs often prefer Header-based auth.

---

## 269. How to use HTTPS in Go servers?

**Answer:**
Production Go servers should never speak plain HTTP. `net/http` provides `ListenAndServeTLS(certFile, keyFile)`.

In cloud environments, we often terminate TLS at the Load Balancer (AWS ALB or Nginx), so the Go app runs on HTTP inside the private VPC. This simplifies certificate management.

However, for direct exposure or zero-trust networks, we use **Go's autocert** package (Let's Encrypt). It automatically negotiates certificates with Let's Encrypt and renews them, allowing you to serve HTTPS with zero manual configuration: `http.Serve(autocert.NewListener("example.com"), handler)`.

---

## 270. How do you sign and verify data in Go?

**Answer:**
For integrity, we use **HMAC** (Hash-based Message Authentication Code).

To sign: `mac := hmac.New(sha256.New, secretKey); mac.Write(data); signature := mac.Sum(nil)`. You send the data and the signature.

To verify: The receiver re-computes the HMAC using the shared secret. They compare their computed signature with the received one using `hmac.Equal()`. We never use `==` for signatures because it leaks timing information (returns faster if the first byte differs), allowing attackers to guess the signature byte-by-byte.

---

## 271. What are best practices for handling secrets in Go?

**Answer:**
**Rule #1: No secrets in source code.** No hardcoded API keys.

We inject secrets via **Environment Variables** (`os.Getenv`) or mount them as files in Kubernetes (`/etc/secrets/db-pass`).

At runtime, we avoid storing secrets in broad scope. If we must keep them in memory, we might use a custom type `type Password string` with a `String()` method that returns `"*****"` to prevent accidental logging. We also rely on tools like **Vault** to rotate secrets dynamically so that a leaked key is only valid for a few minutes.

---

## 272. How do you handle OAuth2 flows in Go?

**Answer:**
We use the `golang.org/x/oauth2` package. It abstracts the complexity of the handshake.

You define an `oauth2.Config` with your ClientID, Secret, and RedirectURL. You generate a "Login Link" for the user. When they return with a `code`, you exchange it: `token, err := config.Exchange(ctx, code)`.

The tricky part is state. You must generate a random `state` string, store it in a cookie, and verify it when the user returns to prevent CSRF attacks during the OAuth handshake itself.

---

## 273. How do you restrict file uploads (size/type)?

**Answer:**
We restrict size using `http.MaxBytesReader`. This wraps the request body and hard-cuts the connection if the client sends more bytes than allowed (e.g., 10MB).

`r.Body = http.MaxBytesReader(w, r.Body, 10 << 20)`

For type, we **never** trust the filename extension or the `Content-Type` header sent by the client. Instead, we read the first 512 bytes of the file and use `http.DetectContentType(headerBytes)` to sniff the actual "Magic Numbers" (file signature). If the signature says it's an EXE but the extension is .JPG, we reject it.

---

## 274. How do you set up CORS properly in Go?

**Answer:**
CORS (Cross-Origin Resource Sharing) controls which domains can access your API. We use `rs/cors` middleware.

A strictly secure setup involves an Allowlist: `AllowedOrigins: []string{"https://myapp.com"}`.

We avoid `AllowedOrigins: "*"` in production unless it's a completely public API. We also must handle the **Preflight** (OPTIONS) requests correctly by sending the `Access-Control-Allow-Methods` and headers. A common mistake is allowing `Access-Control-Allow-Credentials: true` with wildcard origins, which browsers explicitly forbid for security.

---

## 275. How do you scan Go code for vulnerabilities?

**Answer:**
We use **Govulncheck** and **Trivy**.

`govulncheck` is the official Go security tool. It analyzes your binary's symbol table. It reports a vulnerability only if your code *actually calls* the vulnerable function in the dependency, which drastically reduces false positives compared to tools that just check version numbers.

We run this in the CI/CD pipeline. `govulncheck ./...`. If it finds a critical CVE, the build fails, forcing us to upgrade the dependency before merging.

---

## 276. What is the Go ecosystem for SAST tools?

**Answer:**
SAST (Static Application Security Testing) analyzes source code for bad patterns.

The heavyweight is **Gosec** (`securego/gosec`). It scans the AST for hardcoded credentials, weak cryptography (MD5), insecure TLS settings, and SQL injection risks.

We integrate `gosec` into `golangci-lint`. It acts as an automated security auditor that reviews every Pull Request. While it sometimes flags safe code (false positives), it’s invaluable for catching obvious mistakes like `tls.Config{InsecureSkipVerify: true}` left in production code.

---

## 277. How to handle brute force protection in APIs?

**Answer:**
We implement **Rate Limiting** based on IP or UserID.

We use a "Token Bucket" or "Leaky Bucket" algorithm, typically backed by **Redis** for distributed enforcement. Even 5 failed login attempts in 1 minute trigger a 429 Too Many Requests.

For sensitive endpoints (Login), we might use exponential backoff restrictions. If you fail once, wait 1s. Fail twice, wait 2s. This keeps the API responsive for legitimate users while making brute force attacks painfully slow for attackers.

---

## 278. How to secure communication between microservices?

**Answer:**
We assume the internal network is hostile (Zero Trust).

We use **mTLS** (Mutual TLS). Not only does the server present a certificate, but the *client* must also present a valid certificate signed by our internal CA.

This authenticates the identity of the service ("Only the BillingService can call the AccountService"). Tools like **Istio** or **Linkerd** (Service Meshes) handle this automatically via sidecars, so our Go code just speaks HTTP, and the sidecar encrypts everything on the wire.

---

## 279. What is the use of `context.Context` in secure APIs?

**Answer:**
Context is the carrier of **Request-Scoped Authentication**.

When a request passes the auth middleware, we extract the UserID and Roles from the JWT and inject them into the context: `ctx = context.WithValue(ctx, "user", claims)`.

Handlers deeper in the stack retrieve this user info from the context to verify permissions. This ensures that every layer of the application (including the database layer) knows *who* is making the request, allowing for granular audit logging and Row-Level Security.

---

## 280. What is certificate pinning and can it be used in Go?

**Answer:**
Certificate Pinning forces the client to accept *only* a specific public key, not just any cert signed by a trusted CA. It prevents Man-In-The-Middle attacks even if a CA is compromised.

In Go, we implement this in `tls.Config{VerifyPeerCertificate: customVerifyFunc}`.

Inside that function, we hash the server's public key and compare it against our hardcoded "pin" hash. If they differ, we cut the connection. High-value mobile apps use this to prevent users or malware from inspecting API traffic, though it introduces a risk: if you rotate your server cert and forget to update the client app, the app breaks immediately (`Brick Risk`).
