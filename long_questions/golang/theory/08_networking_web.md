# 🟢 Go Theory Questions: 142–160 Networking, APIs, and Web Dev

## 142. How to build a REST API in Go?

**Answer:**
Building a REST API in Go is refreshingly simple because the standard library `net/http` is production-ready out of the box.

You define a handler function, register it to a route like `/users`, and start the server. That’s it.

For a basic microservice, that’s all you need. In the real world, as your API grows, you’ll likely need a router for path parameters (like `/users/{id}`) and middleware for auth and logging. While you can write this yourself, most teams reach for a lightweight library like **Chi** or **Gin** to save time on that boilerplate.

---

## 143. How to parse JSON and XML in Go?

**Answer:**
We use the `encoding/json` and `encoding/xml` packages.

Go relies on **Struct Tags**. You define a struct representing your data, and tag the fields: `` `json:"user_nane"` ``. Then you just call `json.Unmarshal(data, &struct)`.

It’s easy to use, but under the hood, it uses **Runtime Reflection** to inspect those tags. This is relatively slow. For extremely high-performance scenarios—like a real-time bidding server—we skip the standard library and use code-generation tools like `easyjson` to write the parsing logic at compile time.

---

## 144. What is the use of `http.Handler` and `http.HandlerFunc`?

**Answer:**
These represent the interface for everything that handles a web request.

`http.Handler` is the interface: it just needs a `ServeHTTP` method. `http.HandlerFunc` is a clever adapter type that lets you turn a regular function into an object that satisfies that interface.

This is why you can pass a simple function `func(w, r)` to the router. The adapter wraps it so it fits the interface. It’s a great example of Go utilizing interfaces to make APIs ergonomic.

---

## 145. How do you implement middleware manually in Go?

**Answer:**
Middleware in Go is just a fancy name for "wrapping a function."

You write a function that takes a Handler, and returns a new Handler. inside the returned handler, you do your pre-processing (like checking a Login Token), then call `next.ServeHTTP()`, and then do post-processing (like logging how long it took).

It forms a chain—like an onion. The request goes through layer by layer until it hits your business logic, then bubbles back up.

---

## 146. How do you serve static files in Go?

**Answer:**
You don't need Nginx for simple cases; Go can do it natively.

You use `http.FileServer`. You point it at a directory on your disk, and it returns a Handler that serves files from that folder.

It’s surprisingly robust—it handles Content-Type headers, `Range` requests (for seeking in videos), and caching headers automatically. We often use this to serve the built React/Vue assets for our frontend.

---

## 147. How do you handle CORS in Go?

**Answer:**
CORS (Cross-Origin Resource Sharing) is a browser security feature. If your frontend is on port 3000 and backend on 8080, the browser blocks requests unless you say "It's okay."

To fix it, we write a middleware that intercepts the **OPTIONS** method (the preflight request) and writes specific headers: `Access-Control-Allow-Origin: *`.

You have to be careful not to just allow `*` everywhere in production, or you open up security holes. We typically maintain a whitelist of allowed domains.

---

## 148. What are context-based timeouts in HTTP servers?

**Answer:**
This is a critical reliability pattern. Every request in Go carries a `Context`.

If a user closes their browser, or if the request takes too long, that Context is cancelled.

You must pass this `r.Context()` to your database queries. If the context dies, the database query is effectively "interrupted" instantly. This frees up server resources. Without this, your server could be processing queries for users who disconnected 10 seconds ago, eventually causing a cascading failure.

---

## 149. How do you make HTTP requests in Go?

**Answer:**
We use `http.Client`.

For a quick script, `http.Get(url)` works. But in production, you **never** use the default client because it has no timeout. It will hang forever if the server stalls.

You construct your own Client with a strict TimeOut (say, 10 seconds). And—this is the most common bug in Go—you **must** strictly close the response body (`defer resp.Body.Close()`). If you don't, the TCP connection remains open, and you will eventually leak all your file descriptors and crash the app.

---

## 150. How do you manage connection pooling in Go?

**Answer:**
The good news is: Go does it automatically.

The `http.Client` has a `Transport` layer that keeps idle TCP connections open (Keep-Alive). When you finish a request, the connection isn't closed; it's put back in a pool. The next request to the same host reuses it.

However, you have to configure it. Default settings might keep too many (wasting RAM) or too few (causing latency). Tuning `MaxIdleConnsPerHost` is a standard optimization task for high-load services.

---

## 151. What is an HTTP client timeout?

**Answer:**
It is a hard deadline you set on the Client: "If this whole operation—connecting, sending, reading—takes longer than X seconds, kill it."

This is your safety net. Without it, a flaky 3rd party API could stall your entire application by causing all your goroutines to freeze waiting for a response that never comes.

---

## 152. How do you upload and download files via HTTP?

**Answer:**
We use **Streaming**.

For a download, instead of reading the file into memory (RAM), we use `io.Copy(writer, file)`. This copies bytes from the disk directly to the network socket in small chunks.

This allows a tiny server with 512MB RAM to stream a 50GB 4K video file to thousands of users effortlessly. It’s one of Go’s greatest strengths.

---

## 153. What is graceful shutdown and how do you implement it?

**Answer:**
Graceful shutdown means when you deploy new code, you don't just kill the old process. You tell it: "Stop accepting new requests, finish the ones you have, and *then* die."

In Go, we listen for OS signals (`SIGTERM`). When we hear it, we call `server.Shutdown(ctx)`. This function blocks until all active handlers have returned, ensuring no user sees a "Connection Reset" error during a deployment.

---

## 154. How to work with multipart/form-data in Go?

**Answer:**
Multipart forms are how browsers upload files.

Go handles this with `r.ParseMultipartForm`. You give it a memory limit—say, 32MB.

If the upload is smaller than 32MB, it processes it in RAM. If it's larger, Go automatically spills the excess to temporary files on disk. This prevents a malicious user from crashing your server by uploading a 10TB file into RAM.

---

## 155. How do you implement rate limiting in Go?

**Answer:**
We use the **Token Bucket** algorithm, available in the standard library `x/time/rate`.

Imagine a bucket that gets filled with tokens at a steady rate. Every time a request comes in, it must take a token. If the bucket is empty, the request is rejected (HTTP 429).

This allows for steady traffic with occasional "bursts." We usually implement this as middleware to protect our APIs from being overwhelmed by scripts or DDoS attacks.

---

## 156. What is Gorilla Mux and how does it compare with net/http?

**Answer:**
Until Go 1.22, the standard library couldn't handle path parameters (like `/users/{id}`). You had to parse the URL string manually.

Gorilla Mux was the standard solution—it’s a powerful router that supports Regex, variables, and method-based routing.

However, modern Go (1.22+) finally added these features to the standard library. So while Gorilla Mux is great, for new projects, we often stick to the standard library to avoid dependencies.

---

## 157. What are Go frameworks for web APIs (Gin, Echo)?

**Answer:**
Go frameworks are minimalist compared to Rails or Django. They are mostly just a fast Router + a context wrapper.

**Gin** and **Echo** are the big ones. They provide convenience: they can bind JSON to structs automatically, validate inputs, and group routes easily.

We use them when we want to move fast. They reduce boilerplate. The trade-off is they introduce non-standard concepts (like Gin's `Context` object), which locks your handlers into that specific framework.

---

## 158. What are the trade-offs between using `http.ServeMux` and third-party routers?

**Answer:**
It’s the classic "Dependencies vs Convenience" debate.

`http.ServeMux` (Standard Lib) is rock solid, compatible with everything, and will never break. But it’s verbose.

Third-party routers (Chi, Gin) save you typing and offer nice features like middleware groups. But they are external code you have to maintain and audit. Generally, we start with standard lib, and only upgrade if we feel the pain.

---

## 159. How would you implement authentication in a Go API?

**Answer:**
For APIs, we almost always use **Stateless Auth**, typically JWTs (JSON Web Tokens).

When a user logs in, we give them a signed Token. For every subsequent request, they send that token in the header.

We write a middleware that parses the token, verifies the cryptographic signature, extracts the User ID, and sticks it into the `request.Context`. This way, every handler knows exactly who is making the request without us having to hit the database to check a session table.

---

## 160. How do you implement file streaming in Go?

**Answer:**
Streaming is about pipelines.

You have a `Reader` (the source) and a `Writer` (the destination). You connect them with `io.Copy`.

If you need to transform the data in the middle—say, gzipping it on the fly—you just wrap the writer: `gzip.NewWriter(httpWriter)`. Now, as you copy bytes from the file, they are compressed and sent to the client instantly. No buffering, no waiting. It’s extremely memory efficiency.
