# 🟡 **21–40: Intermediate**

### 21. What is the difference between synchronous and asynchronous functions?
"**Synchronous functions** block the execution of the program until they finish their task. If I call a synchronous file read, the entire Node.js server sits idle, unable to serve any other users, until that file is loaded from disk.

**Asynchronous functions** execute non-blockingly. When I call an async file read, Node.js offloads the actual reading task to the background. The server immediately moves to the next line of code and can serve thousands of other incoming HTTP requests. Once the file is ready, Node.js fires a callback to handle the data.

I stick to synchronous functions *only* during the initial startup phase of an application (like reading a config file once). For everything else that happens during the lifecycle of a request, I strictly use asynchronous functions to keep the event loop fast."

#### Indepth
In Node.js, asynchronous I/O relies on OS primitives (like `epoll` or `kqueue`) for network operations, and a libuv-managed thread pool specifically for heavy file system or cryptographic operations. This architecture ensures the main JavaScript thread is never blocked by external, slow resource access.

---

### 22. What is a Promise in Node.js?
"A **Promise** is an object representing the eventual completion (or failure) of an asynchronous operation and its resulting value.

Instead of passing callbacks into functions (which leads to deeply nested and hard-to-read code), a function returns a Promise. The Promise exists in one of three states:
1. **Pending**: The initial state, neither fulfilled nor rejected.
2. **Fulfilled**: The operation completed successfully.
3. **Rejected**: The operation failed.

I use `.then()` to handle a fulfilled promise, and `.catch()` to handle a rejection. Promises make my asynchronous code look much more linear and dramatically improve error handling."

#### Indepth
Promises resolve as **microtasks**. When a Promise resolves, its `.then()` callbacks are pushed onto the microtask queue. The event loop prioritizes the microtask queue intensely, executing it entirely before moving to the next phase (like timers or I/O), ensuring Promise callbacks run as early as possible.

---

### 23. How do you convert a callback to a promise?
"To convert an older, callback-based API into a modern Promise, I simply wrap it in a `new Promise` constructor.

Inside the constructor, I call the old function. When its callback triggers, I check for an error. If there's an error, I call the `reject()` function provided by the Promise; if it succeeds, I call the `resolve()` function with the data.

```javascript
function myPromisifiedFunction(args) {
  return new Promise((resolve, reject) => {
    oldCallbackFunction(args, (err, data) => {
      if (err) return reject(err);
      resolve(data);
    });
  });
}
```
However, in modern Node.js, I rarely do this manually anymore. I just use the built-in `util.promisify()` function to do it automatically."

#### Indepth
`util.promisify()` specifically looks for the standard Node.js error-first callback signature: `(err, value) => ...`. If a legacy library uses an unconventional callback signature, `util.promisify` will fail, and you must use the manual wrapper approach.

---

### 24. What are async/await keywords? How do they work?
"`async/await` is syntactic sugar built directly on top of Promises. 

When I declare a function with the `async` keyword, it automatically implies that the function will return a Promise. 
Inside that function, I can use the `await` keyword in front of any Promise. `await` pauses the execution of that specific function until the Promise resolves, essentially making asynchronous code look and read exactly like synchronous code.

This entirely eliminates `.then()` chains and allows me to handle asynchronous errors using standard `try...catch` blocks, which represents a massive improvement in code readability and debugging."

#### Indepth
While `await` pauses the execution of its *enclosing* `async` function, it does not block the Node.js event loop. Node.js continues to execute other code, callbacks, and events while that specific `async` function yields control back to the event loop until the awaited Task completes.

---

### 25. What are streams in Node.js?
"**Streams** are collections of data—just like arrays or strings—but the difference is that streams might not be available all at once, and they don’t have to fit into RAM.

They allow me to read or write data sequentially, piece by piece (or 'chunk by chunk'). 

If I need to serve a 2GB video file, reading it entirely into memory using `fs.readFile` would crash the server. Instead, I use a Stream to pipe the file directly to the HTTP response. It reads a small chunk, sends it, drops it from memory, and grabs the next chunk. This keeps my server's memory footprint incredibly low, regardless of the file size."

#### Indepth
Streams are instances of `EventEmitter`. They emit events like `'data'` (when a new chunk is available), `'end'` (when there is no more data to read), and `'error'`. Because of this event-driven architecture, streams can be easily composed and piped together using the `.pipe()` or modern `.pipeline()` methods.

---

### 26. Explain the four types of streams in Node.js.
"Node.js has four fundamental types of streams:

1. **Readable Streams**: Used for reading data (e.g., `fs.createReadStream`, or reading an incoming HTTP request).
2. **Writable Streams**: Used for writing data (e.g., `fs.createWriteStream`, or sending an HTTP response).
3. **Duplex Streams**: Streams that are both readable and writable (e.g., a TCP network socket where you can send and receive data simultaneously).
4. **Transform Streams**: A special type of Duplex stream that modifies the data as it is written and read (e.g., `zlib.createGzip()` to compress data on the fly).

Understanding these four types allows me to build highly efficient data processing pipelines without consuming excessive memory."

#### Indepth
All four stream types operate on either strings/Buffers (standard mode) or arbitrary JavaScript objects (Object Mode). In Object Mode, instead of chunking raw bytes, the stream processes discrete objects, which is extremely useful for transforming a continuous flow of database records.

---

### 27. What is backpressure in streams?
"**Backpressure** occurs when a writable stream is receiving data far faster than it can actually process or write it.

If I'm reading from a blazing fast SSD and piping it to a slow network connection, the network socket's internal buffer will quickly fill up. If Node.js kept reading from the disk indiscriminately, it would bloat the server's RAM with un-sent data.

Backpressure is the mechanism where the Writable stream says, 'Stop sending data for a moment, I need to clear my buffer.' The Readable stream pauses, waits for a `'drain'` event from the Writable stream, and only then resumes reading."

#### Indepth
When using the `.pipe()` method—or preferably, the modern `stream.pipeline()` utility—Node.js handles backpressure entirely automatically. It dynamically pauses and resumes the underlying Readable stream based on the `highWaterMark` capacity of the Writable stream, preventing memory exhaustion.

---

### 28. How can you create a custom readable stream?
"To create a custom readable stream, I inherit from the `Readable` class provided by the `stream` module and implement the `_read()` method.

The `_read(size)` method is called internally by the stream whenever it wants me to fetch more data. Inside this method, I fetch some data (from a database, an API, or generating it) and call `this.push(data)` to add it to the stream's internal buffer. 

When there is no more data left to read, I strictly call `this.push(null)` to signal an EOF (End of File) and close the stream gracefully."

#### Indepth
Instead of creating a full class structure for simple use-cases, modern Node.js allows creating streams directly from iterators or generators using `Readable.from()`. If you have an asynchronous generator function fetching database rows, `Readable.from(myGenerator())` instantly turns it into a robust, backpressure-aware Readable Stream.

---

### 29. How does error handling differ between callbacks and promises?
"In the **callback pattern**, error handling relies on the 'Error-First' convention. Every single callback must check `if (err) { /* handle */ }` at the very beginning. If you forget this check in a deeply nested callback, an error can easily go unnoticed or crash the app silently.

In the **Promises pattern**, error handling is centralized. If an error occurs anywhere in a chain of `.then()` statements, it immediately bypasses the rest of the chain and drops straight into the `.catch()` block at the end. 

With `async/await` (which uses Promises), I can wrap my asynchronous calls in a standard `try...catch` block, totally unifying synchronous and asynchronous error handling."

#### Indepth
When a callback throws a synchronous exception (e.g., a TypeError on `null.prop`) *inside* the callback function, that error goes unhandled by the callback's `if(err)` logic and crashes the application. Promises naturally capture both explicit rejections and implicit synchronous exceptions within their `.then()` blocks, routing both to `.catch()`.

---

### 30. What is middleware in Express.js?
"In Express.js, **middleware** functions are functions that have access to the request object (`req`), the response object (`res`), and the `next` middleware function in the application’s request-response cycle.

Middleware sits between the incoming request and the final route handler. I use middleware to perform tasks like:
- Parsing incoming JSON payloads.
- Verifying authentication tokens (JWTs).
- Logging request metrics.
- Handling CORS.

It allows me to write reusable blocks of code that apply to many different routes without duplicating logic."

#### Indepth
Middleware execution is strictly sequential based on the order it is defined via `app.use()`. If a middleware function does not abruptly end the request-response cycle by sending a response (e.g., `res.send()`), it *must* call `next()` to pass control to the next middleware function. Otherwise, the request will be left hanging indefinitely.

---

### 31. What is the role of the `next()` function in Express middleware?
"The `next()` function is the mechanism that moves the request along the Express middleware chain.

When a middleware function completes its task—say, verifying an API key—it calls `next()` to tell Express: 'I’m done, pass this request to the next middleware or route handler.' 

If I forget to call `next()`, the entire request will hang, and the client's browser will eventually time out. Conversely, if I call `next(err)` and pass an error object, Express instantly skips all normal routes and jumps straight to my specially defined Error Handling middleware."

#### Indepth
Express also provides `next('route')`, which is exclusively used within `app.get()` or similar router methods. Calling `next('route')` bypasses any remaining middleware in the current route stack and immediately jumps to the next route that matches the path, which is highly useful for conditional routing logic.

---

### 32. How can you handle 404 errors in Express?
"In Express, 404 errors don’t occur via exceptions or errors; they simply happen when a request reaches the very end of the middleware chain without any route matching the URL or sending a response.

To handle a 404 globally, I place a catch-all middleware function at the absolute bottom of my route definitions (right before the error-handling middleware):

```javascript
app.use((req, res, next) => {
  res.status(404).send("Sorry, that route doesn't exist.");
});
```
This guarantees that if no previous route responded, this final block will intercept the request and send a clean 404 page or JSON response."

#### Indepth
It is common practice to pass an error object from a 404 handler to a centralized error handler: `next(new Error('Not Found'))`. However, you must explicitly set the status: `res.status(404)`. Otherwise, the centralized error handler will default to a 500 Internal Server Error, masking a client-side navigation mistake as a server crash.

---

### 33. How do you configure environment variables in Node.js?
"Environment variables allow me to configure the application dynamically without hardcoding sensitive information (like API keys or Database passwords) into the source code.

In Node.js, these variables are exposed globally through the `process.env` object.

In a production environment (like AWS or Heroku), I set these variables directly on the server's operating system. 
In local development, constantly typing out export commands is tedious, so I use a `.env` file containing key-value pairs and load them into `process.env` using a library called `dotenv`."

#### Indepth
As of Node.js v20.6.0, native support for `.env` files was introduced via the `--env-file` flag (e.g., `node --env-file=.env app.js`). This eliminates the strict dependency on the third-party `dotenv` package for basic environment variable loading in modern Node projects.

---

### 34. What is `dotenv` and how is it used?
"`dotenv` is a zero-dependency npm module that loads environment variables from a `.env` file into Node's `process.env`.

I create a `.env` file at the root of my project (`PORT=3000\nDB_PASS=secret`) and make sure to add it to `.gitignore` so I don't accidentally push passwords to GitHub. 

Then, at the very entry point of my application (e.g., `index.js`), I call `require('dotenv').config()`. From that line onward, any other file in the application can securely access `process.env.DB_PASS`."

#### Indepth
When using ES Modules, importing `dotenv` at the top of an entry file can sometimes cause race conditions because ES imports are hoisted. To guarantee variables are loaded before any other modules execute, it is often safer to preload it via the command line: `node --preload dotenv/config index.js`.

---

### 35. How do you handle file uploads in Node.js?
"Node.js natively receives file uploads as highly chunked stream data encoded as `multipart/form-data`. Parsing this manually is complex.

In the Express ecosystem, I strictly use a middleware called **Multer**.

I configure Multer to dictate where the files should be saved (either directly to disk, or securely held in memory as Buffers). I then attach the Multer middleware to a specific route, like `app.post('/upload', upload.single('avatar'), ...`. 
Multer processes the incoming stream and neatly attaches the file metadata (filename, size, buffer) to `req.file`, making it incredibly easy to process or stream off to AWS S3."

#### Indepth
If you are streaming extremely large files directly to a cloud provider like S3, using Multer's default disk storage is inefficient because it writes the file locally first, then uploads it. Instead, use `multer.memoryStorage()` for small files, or combine Multer with streams (via `multer-s3`) to pipe the incoming network request directly to the cloud without touching the local disk.

---

### 36. How can you serve static files using Express?
"Express has a built-in, highly optimized middleware specifically for serving static assets like images, CSS files, and frontend JavaScript files.

I use the `express.static()` middleware. For example, if I have a folder named `public` containing my images, I add this to my server:
`app.use(express.static('public'));`

Now, if a user requests `http://localhost/logo.png`, Express automatically looks inside the `public` folder, sets the correct MIME types (Content-Type: image/png), and streams the file to the browser efficiently."

#### Indepth
In high-traffic production environments, Node.js shouldn't be serving static files at all. Serving static assets blocks the event loop from doing what Node does best: building robust APIs. Static files should be offloaded to a dedicated reverse proxy (like NGINX) or a CDN (like Cloudflare or AWS CloudFront), significantly reducing the CPU load on the Node.js server.

---

### 37. What is the `cluster` module?
"Because Node.js runs on a single thread, it can only utilize one CPU core by default. If my server has an 8-core processor, 7 of those cores are sitting idle.

The `cluster` module allows me to easily spawn a network of worker processes (one for each CPU core) that all share the same server port. 

There's a main 'primary' process that listens to the port and round-robins incoming HTTP requests to the worker processes. If one worker crashes, the primary process detects it and can spawn a new one, providing both massive scalability and fault tolerance."

#### Indepth
While the `cluster` module is powerful, manually managing worker lifecycles, graceful shutdowns, and zero-downtime reloads requires complex boilerplate code. In modern production, developers almost unanimously rely on process managers like **PM2** (e.g., `pm2 start app.js -i max`), or container orchestration tools like Kubernetes, which handle clustering abstractly.

---

### 38. What is the `child_process` module used for?
"While Node.js is great for I/O, it shouldn't execute heavy CPU-bound tasks or run shell commands on the main thread.

The `child_process` module allows me to spin up external OS processes entirely outside of my Node application. I use it to run bash scripts, execute Python code, or run heavy image manipulation tools (like ImageMagick).

It allows my Node program to communicate with these external tools via `stdin`, `stdout`, and `stderr` streams, essentially acting as the central orchestrator for a multi-language backend."

#### Indepth
The module provides four primary ways to create processes: `spawn`, `exec`, `execFile`, and `fork`. Choosing the right method is critical for memory management—`exec` buffers the entire output into RAM (which can crash on large outputs), while `spawn` returns streams, making it safe for handling massive amounts of data from the child process.

---

### 39. What is the difference between `spawn` and `exec`?
"Both create child processes, but they disagree on how data is returned to Node.js.

`exec` runs a command in a shell and buffers the *entire* output into memory before passing it into a callback string. It’s perfect for small scripts (`ls -la` or getting a git hash), but will crash with a 'maxBuffer' error if the output is too large.

`spawn` does not use a shell by default and returns data asynchronously via Streams. It allows me to continually process massive amounts of data piece-by-piece as the command runs, making it the only safe choice for long-running processes or massive file conversions."

#### Indepth
Beyond memory buffering, `exec` executes commands inside a fully functional shell (like `/bin/sh`). This makes it vulnerable to **Command Injection** attacks if user input is concatenated into the execution string. `spawn` executes the binary directly, treating arguments strictly as arguments, which safely neutralizes injection attacks.

---

### 40. How can you handle CORS in a Node.js application?
"**CORS** (Cross-Origin Resource Sharing) is a browser security feature that blocks frontend code running on one domain (like `foo.com`) from making API requests to a server on a different domain (like `api.bar.com`) unless the server explicitly allows it.

In an Express app, I handle this by attaching appropriate HTTP headers to my responses (like `Access-Control-Allow-Origin`). 

Rather than writing this header logic manually, I install the `cors` npm package. By simply adding `app.use(cors())`, I securely authorize cross-origin requests. I can also configure it to only allow specific domains (whitelisting) or certain HTTP methods."

#### Indepth
For non-simple requests (like sending JSON or using custom headers), the browser will send a preliminary `OPTIONS` request called a **Preflight Request** before the actual `POST` request. The `cors` middleware automatically intercepts and heavily caches these preflight requests (`Access-Control-Max-Age`), preventing the server from processing duplicate, expensive routing logic.
