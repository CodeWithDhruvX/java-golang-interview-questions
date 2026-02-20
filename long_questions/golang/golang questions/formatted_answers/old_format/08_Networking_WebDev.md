# 🔵 **141–160: Networking, APIs, and Web Dev**

### 142. How to build a REST API in Go?
"I typically start with the standard `net/http` package.

I define handlers using `http.HandleFunc("/users", handler)`.
Inside the handler, I decode the JSON body, interact with my service layer, and encode the response to `w`.
For more complex routing (like `/users/{id}`), I'll upgrade to **Chi** or **Gin**. Chi is my favorite because it feels like a lightweight extension of the standard library, whereas Gin is a full framework."

#### Indepth
The standard library's `http.ServeMux` got a huge upgrade in Go 1.22. It now supports method-based routing (`POST /items`) and wildcards (`/items/{id}`). This negates the need for Chi or Gorilla Mux for 90% of use cases, making the "stdlib-only" approach even more viable.

---

### 143. How to parse JSON and XML in Go?
"I use `encoding/json`.

I define a struct with tags: `type User struct { Name string \`json:"name"\` }`.
Then I use `json.Unmarshal(data, &user)` to parse it.
For XML, it's identical but with `encoding/xml` and `xml:"..."` tags.
Since `encoding/json` uses reflection, for extremely high-throughput systems, I might swap it for **easyjson** or **fastjson** to generate static parsing code."

#### Indepth
`encoding/json` respects struct tags like `json:"-"` (ignore field) and `json:",omitempty"` (omit if zero-value). A common pitfall is handling `time.Time`: JSON has no date standard, so Go uses RFC3339 strings by default. You can override this by implementing `MarshalJSON` on a custom wrapper type.

---

### 144. What is the use of `http.Handler` and `http.HandlerFunc`?
"`http.Handler` is an interface with a singe method: `ServeHTTP(w, r)`.
Anything that implements this can process web requests.

`http.HandlerFunc` is a convenience adapter. It lets me take a simple function `func(w, r)` and cast it to `http.HandlerFunc(myFunc)`, which now satisfies the interface. It saves me from creating a new struct type for every single route."

#### Indepth
This adapter pattern is everywhere in Go. `http.HandlerFunc(myFunc)` works because Go allows methods on types derived from functions. It essentially says "when `ServeHTTP` is called on this function value, just execute the function itself".

---

### 145. How do you implement middleware manually in Go?
"Middleware is just a function that takes an `http.Handler` and returns a *new* `http.Handler`.

`func Logging(next http.Handler) http.Handler`.
Inside the returned handler, I do my pre-processing (logging start time), call `next.ServeHTTP(w, r)`, and then do post-processing (logging duration).
This **Chain of Responsibility** pattern is elegant because I can wrap layers infinitely: `Auth(RateLimit(Logger(Handler)))`."

#### Indepth
A critical detail in middleware is **Conditionality**. You might want to skip authentication for the `/health` endpoint. You can handle this inside the Auth middleware by checking `r.URL.Path`, or better, by wrapping only the specific sub-routers that need protection, leaving public routes outside the wrapper.

---

### 146. How do you serve static files in Go?
"I use `http.FileServer`.
`http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))`.

In production, I prefer to embed the assets into the binary using **go:embed**.
`//go:embed assets`
`var assets embed.FS`
Then I serve from that `embed.FS`. This gives me a single-file deployment—no need to copy a separate 'static' folder to the server."

#### Indepth
`go:embed` can also match patterns: `//go:embed css/*.css`. The embedded filesystem is read-only and is efficient (it maps directly to the binary's data segment). This solves the "works on my machine, fails in docker" because the assets are physically inside the executable.

---

### 147. How do you handle CORS in Go?
"Cross-Origin Resource Sharing is handled via Middleware.

I write a wrapper that sets headers like `Access-Control-Allow-Origin: *`.
Crucially, it must intercept **OPTIONS** requests (preflight) and return 200 OK immediately.
I often use the `rs/cors` library because getting the headers exactly right for all edge cases (credentials, exposed headers) is tedious and error-prone manually."

#### Indepth
CORS is browser security, not server security. It prevents site A from reading data from site B via JS. If you are building a server-to-server API (like a webhook receiver), CORS is irrelevant. For public APIs, setting `Access-Control-Allow-Origin: *` is fine, but for internal apps, whitelist specific domains.

---

### 148. What are context-based timeouts in HTTP servers?
"They are my defense against slow clients (Slowloris attacks).

Use `http.TimeoutHandler` or configure `http.Server{ReadTimeout: 5s, WriteTimeout: 10s}`.
Inside the handler, I pass `r.Context()` to any blocking calls (DB, API). If the client disconnects or times out, the context is canceled, and my DB query aborts immediately. This prevents a single slow client from tying up a database connection."

#### Indepth
`http.Server` has distinct timeouts: `ReadTimeout` (time to read body), `WriteTimeout` (time to write response), and `IdleTimeout` (Keep-Alive limit). Setting `WriteTimeout` is dangerous because if it triggers, Go forcibly closes the TCP connection, potentially cutting off a valid JSON response mid-stream. Use `http.TimeoutHandler` for cleaner behavior.

---

### 149. How do you make HTTP requests in Go?
"I use `http.NewRequest` and a custom `http.Client`.

**I never use the default `http.Get()` in production.**
The default client has NO timeout. If the server hangs, my goroutine hangs forever. Eventually, I run out of file descriptors and crash.
I always instantiate: `client := &http.Client{Timeout: 10 * time.Second}`."

#### Indepth
The default client is a shared global variable. Modifying it (like `http.DefaultClient.Timeout = ...`) affects all packages using it, which is a race condition. Always create your own client instance. You can often reuse one client instance for the entire app to share the connection pool.

---

### 150. How do you manage connection pooling in Go?
"The `http.Client` handles it automatically via the `Transport`.

It keeps a pool of idle TCP connections open (Keep-Alive).
The catch: I **MUST** read the response body fully and close it (`resp.Body.Close()`).
If I fail to do this, the connection cannot be returned to the pool and remains in a `CLOSE_WAIT` state, eventually causing a connection leak."

#### Indepth
To safely drain the body, use `io.Copy(io.Discard, resp.Body)`. Just calling `Close()` isn't enough; if there are unread bytes on the wire, Go might close the TCP connection instead of reusing it. Draining allows the connection to be kept alive for the next request.

---

### 151. What is an HTTP client timeout?
"It depends on which timeout.

`http.Client.Timeout` is the **total** time limit for the interaction (Dial + TLS + Headers + Body).
If I need more granular control (e.g., '10s to connect, but 1 hour to download'), I use `net.Dialer` timeouts or `context.WithTimeout`. call."

#### Indepth
`context.WithTimeout` is usually superior because it propagates. If Service A calls Service B with a 5s timeout, and B calls C, passing the context ensures C knows it only has 4.9s left. `http.Client.Timeout` is hard/local and doesn't respect the remaining time budget of the incoming request.

---

### 152. How do you upload and download files via HTTP?
"For **Uploads**: I use `r.FormFile("file")`. Go parses the multipart form and gives me a file handle (either in memory or on disk).
For **Downloads**: I set `Content-Disposition: attachment` and stream the file to `w`.
I never read the whole file into a `[]byte`. I use `io.Copy(w, file)`. This uses a fixed 32KB buffer, so I can serve a 10TB file with 10MB of RAM."

#### Indepth
For un-trusted uploads, always wrap `io.Copy` with `io.LimitReader(r.Body, limit)` to prevent disk-filling attacks where a user claims to send 1kb but sends 100GB. Validation of file magic numbers (signatures) is also crucial, as relying on request Content-Type headers is insecure.

---

### 153. What is graceful shutdown and how do you implement it?
"It enables the server to finish existing requests before stopping.

I listen for `SIGINT` or `SIGTERM`.
When caught, I call `server.Shutdown(ctx)`.
This stops the listener immediately (so no new requests come in) but blocks until all active handlers return (or `ctx` expires). This is critical for deployments to ensure users don't see 502 Bad Gateway errors during a rollout."

#### Indepth
Kubernetes complicates this. When a Pod terminates, K8s removes it from the Service endpoints, but this propagation is asynchronous. You should sleep for ~5-10 seconds *before* calling Shutdown to allow the Load Balancer to stop sending new traffic, otherwise, you might kill requests that were in-flight during the Shutdown sequence.

---

### 154. How to work with multipart/form-data in Go?
"I use `r.ParseMultipartForm(maxMemory)`.

The `maxMemory` argument (e.g., 32MB) tells Go: 'Keep requests smaller than this in RAM; spill anything larger to temporary files on disk'.
Then I access files via `r.MultipartForm.File`. Cleaning up created temp files is usually handled automatically, but I can call `r.MultipartForm.RemoveAll()` to be sure."

#### Indepth
Multipart forms are slow to parse. If you are building a high-performance file upload service, consider using "raw" uploads (PUT binary body) instead of `multipart/form-data`. It saves the CPU cost of boundary parsing and MIME decoding.

---

### 155. How do you implement rate limiting in Go?
"I use the **Token Bucket** algorithm, usually via `golang.org/x/time/rate`.

I create a limiter per user (keyed by IP).
Middleware checks `limiter.Allow()`. If false, I return `429 Too Many Requests`.
For distributed systems, local memory isn't enough, so I implement the Token Bucket in **Redis** (using Lua scripts) to share the limit across all API instances."

#### Indepth
Rate limiting headers are important. Send `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset` to tell polite clients when to back off. Without these, clients will just blindly retry, effectively DDOS-ing your gateway with 429 errors.

---

### 156. What is Gorilla Mux and how does it compare with net/http?
"Gorilla Mux was the standard for years because `net/http` was too simple.

Mux allows Method-based routing (`.Methods("POST")`) and Regex paths (`/products/{id:[0-9]+}`).
However, as of **Go 1.22**, the standard library added these features! So now, I prefer standard `net/http` for new projects. Mux is still great, but it's in maintenance mode."

#### Indepth
One big difference: `net/http` uses exact matching or prefix matching. It prioritizes the *most specific* pattern. `/images/thumbnails/` wins over `/images/`. This behavior is robust and predictable, whereas regex-based routing orders often depend on declaration order, which is fragile.

---

### 157. What are Go frameworks for web APIs (Gin, Echo)?
"**Gin** and **Echo** are the two heavyweights.

They are faster than `net/http` (using Radix tree routers) and provide a lot of helpers: Data binding (`c.BindJSON`), Validation, Grouping routes (`v1 := r.Group("/v1")`).
I use them if I need to build a large API quickly. If I'm building a small microservice, I stick to the standard library to keep dependencies low."

#### Indepth
Gin uses a custom `Context` struct, which is **not** thread-safe. You cannot pass `c` to a goroutine because it resets after the handler returns. You must call `c.Copy()` to pass a safe snapshot to a background worker. This is a common source of panic in Gin apps.

---

### 158. What are the trade-offs between using `http.ServeMux` and third-party routers?
"**ServeMux** (Stdlib):
*   Pros: No dependencies, stable APIs, forward compatible.
*   Cons: Verbose for complex middleware chains or parameter extraction.

**Chi/Gin**:
*   Pros: Concise, powerful middleware ecosystem, fast parameter extraction.
*   Cons: External dependency risk.

I usually treat **Chi** as the sweet spot—it’s just a Router that plays nice with standard `http.Handler`."

#### Indepth
Performance differences (Radix tree vs Regex vs Map) rarely matter for typical web apps (< 10k RPS). The bottleneck is almost always the Database or Network I/O. Choose the router based on DevX (developer experience) and middleware ecosystem, not raw nanosecond benchmarks.

---

### 159. How would you implement authentication in a Go API?
"I typically use **JWTs** (JSON Web Tokens).

1.  **Login**: User sends credentials. I verify them against the DB.
2.  **Issue**: I sign a JWT containing the `user_id` and `exp` time.
3.  **Middleware**: For every protected route, I extract the token from `Authorization: Bearer <token>`, parse it, and verify the signature.
4.  **Context**: I inject the `user_id` into `r.Context()` so handlers know who is calling."

#### Indepth
Stateful sessions (Cookies + Redis) are better than JWTs if you need instant revocation (e.g., banning a user). With JWTs, you can't invalidate a token until it expires (unless you implement a blacklist, which defeats the stateless purpose). Choose the right tool for the security model.

---

### 160. How do you implement file streaming in Go?
"I rely on the universal `io.Reader` and `io.Writer` interfaces.

I don't load data. I pass the stream.
If I'm proxying a file from S3 to the user:
`s3Resp := s3.GetObject(...)`
`io.Copy(w, s3Resp.Body)`
The bytes flow from S3 -> Go -> Client in tiny chunks. This keeps my memory usage flat and minimal."

#### Indepth
When proxying, be sure to copy the implementation's `Flush()` behavior. Standard `io.Copy` buffers until the buffer fills. For real-time streaming (like server-sent events), you need to type-assert the `ResponseWriter` to `http.Flusher` and call `Flush()` after every write chunk.
