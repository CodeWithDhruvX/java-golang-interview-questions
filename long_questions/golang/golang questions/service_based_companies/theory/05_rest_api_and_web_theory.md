# 🗣️ Theory — REST API & Web Development in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you build REST APIs in Go? What packages are commonly used?"

> *"Go's standard library already has a solid HTTP package — `net/http` — so you can build REST APIs without any framework at all. You define handler functions with the signature `func(w http.ResponseWriter, r *http.Request)`, register them to routes using `http.HandleFunc` or `http.ServeMux`, and call `http.ListenAndServe`. For production apps though, most teams use a router like Gin or Echo. Gin is the most popular — it adds cleaner routing, path parameter extraction, automatic JSON binding, and middleware support. It's about 40 times faster than the standard library's router for complex routing scenarios."*

---

## Q: "What is the difference between `http.Handler` and `http.HandlerFunc`?"

> *"`http.Handler` is an interface with one method: `ServeHTTP(ResponseWriter, *Request)`. So any type that has that method is a handler and can be used with `http.Handle`. `http.HandlerFunc` is a function type that also implements `http.Handler`. It exists so you can use a plain function as a handler without creating a struct. When you write `http.HandleFunc('/path', myFunc)`, Go internally converts your function to an `http.HandlerFunc`. They're two different ways to create handlers — use `HandlerFunc` for simple cases, create a struct implementing `Handler` when your handler needs to carry state."*

---

## Q: "What is middleware in Go? How do you implement it?"

> *"Middleware is code that runs before or after your handler — for things like logging, authentication, compression, CORS. In Go, middleware is implemented as a function that takes an `http.Handler` and returns a new `http.Handler`. Inside, it does the middleware logic, then calls the next handler. You chain them by wrapping: `LoggingMiddleware(AuthMiddleware(myHandler))`. The inner-most handler runs last on the way in, and first on the way out. Frameworks like Gin have their own middleware system with `r.Use(middleware)`, but the concept is identical — a function that wraps the next handler."*

---

## Q: "How does JSON encoding and decoding work in Go?"

> *"Go has the `encoding/json` package built into the standard library. To convert a struct to JSON — marshaling — you call `json.Marshal(myStruct)` and get back bytes. To go the other way — unmarshaling — you call `json.Unmarshal(jsonBytes, &myStruct)`. The struct controls the output through struct tags: `json:'name'` sets the JSON key, `json:'-'` excludes a field, `json:'name,omitempty'` omits the field if it's a zero value. For HTTP request bodies specifically, use `json.NewDecoder(r.Body).Decode(&myStruct)` — it streams the body instead of reading it all into memory first."*

---

## Q: "How do you handle CORS in Go?"

> *"CORS — Cross-Origin Resource Sharing — is handled by adding response headers that tell the browser it's safe to make cross-origin requests. Your API needs to reply with `Access-Control-Allow-Origin`, `Access-Control-Allow-Methods`, and `Access-Control-Allow-Headers` headers. For preflight requests — the OPTIONS method the browser sends first — you return a 204 and stop there. Manually, you implement a middleware that sets these headers. In production, you'd use a library like `github.com/gin-contrib/cors` for Gin, which gives you a config-based approach and handles all the edge cases properly."*

---

## Q: "How do you implement graceful shutdown for a Go HTTP server?"

> *"Graceful shutdown means: stop accepting new requests, but let existing requests finish before the process exits. You do it by running the server in a goroutine, then waiting for an OS signal — usually `SIGINT` or `SIGTERM` — on a channel. When that signal arrives, you call `server.Shutdown(ctx)` with a timeout context. Shutdown stops the listener immediately so no new requests come in, then waits for existing handlers to complete — or until the context deadline. This is especially important in Kubernetes where a pod is killed with SIGTERM and you get a grace period to finish ongoing work."*

---

## Q: "What are the most common HTTP status codes you use in REST APIs?"

> *"The ones I use most are: 200 for success with a body, 201 when a resource is created, 204 for success with no body — like a delete. For client errors: 400 for a bad request where the client sent invalid data, 401 for unauthorized meaning you need to authenticate, 403 for forbidden meaning you're authenticated but don't have permission, 404 for not found, 409 for a conflict like a duplicate entry, and 422 for unprocessable entity — often used for validation failures. For server errors it's mostly 500 for unexpected internal errors and 503 for service unavailable when you're overloaded or rate limiting."*

---

## Q: "How does request validation work in Go REST APIs?"

> *"There are two levels. First, JSON parsing — if the JSON is malformed, `json.Decoder.Decode` returns an error and you send a 400. But that only validates structure, not content. For content validation — like 'is this email valid?' or 'is this age between 0 and 120?' — the de facto standard library is `github.com/go-playground/validator`. You add struct tags like `validate:'required,email'` or `validate:'gte=0,lte=120'`, then call `validate.Struct(myStruct)`. It returns a `ValidationErrors` value you can iterate to give the client specific field-level error messages. Gin also has built-in binding validation using `ShouldBindJSON`."*
