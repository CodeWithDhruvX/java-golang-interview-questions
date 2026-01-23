# NodeJS Interview Questions & Answers

## 🔹 Basics (Beginner Level)

**Q1: What is Node.js?**
Node.js is an open-source, cross-platform JavaScript runtime environment that executes JavaScript code outside a web browser. It is built on Chrome's V8 JavaScript engine and is designed to build scalable network applications.

**Q2: How is Node.js different from JavaScript in the browser?**
Node.js runs on a server and has access to the file system, OS, and network modules, but lacks DOM and window objects. Browser JavaScript runs in the client's browser, manipulates the DOM, and has restricted access to the OS for security.

**Q3: What is the V8 engine?**
V8 is an open-source high-performance JavaScript and WebAssembly engine developed by Google for Chrome. Node.js uses V8 to compile JavaScript directly into native machine code, making execution extremely fast.

**Q4: What is the role of the `package.json` file?**
`package.json` is the manifest file of a Node.js project. It holds metadata (name, version), scripts, and the list of dependencies (development and production) required to run the project.

**Q5: What is npm? How is it different from npx?**
`npm` (Node Package Manager) is used to install, manage, and share packages. `npx` (Node Package Execute) is a tool to execute packages (binaries) directly from the npm registry without installing them globally or locally first.

**Q6: What are modules in Node.js?**
Modules are reusable blocks of code (functions, objects, variables) that can be exported from one file and imported into another. Node.js uses the CommonJS module system (`require`) by default.

**Q7: How do you export and import modules in Node.js?**
To export: `module.exports = { myFunction };` or `exports.myFunction = ...`.
To import: `const { myFunction } = require('./myModule');`.

**Q8: What is the difference between `require()` and `import`?**
`require()` is CommonJS (synchronous, dynamic, used in older Node.js). `import` is ES Module (asynchronous, static analysis, standard in modern JS). Node.js supports both (ESM requires `.mjs` or `"type": "module"`).

**Q9: What is a callback function?**
A function passed as an argument to another function, which is then invoked inside the outer function to complete some kind of routine or action, typically after an asynchronous operation completes.

**Q10: What is an event loop in Node.js?**
The entity that handles asynchronous callbacks. It allows Node.js to perform non-blocking I/O operations by offloading operations to the system kernel whenever possible and processing their callbacks in phases.

**Q11: How does Node.js handle asynchronous operations?**
Using the Event Loop, Callbacks, Promises, and Async/Await. Heavy tasks are offloaded to `libuv`'s thread pool or the OS kernel, while the main thread remains free to accept new requests.

**Q12: What is non-blocking I/O?**
I/O operations (file reading, network requests) that do not block the execution of further code. Node.js initiates the operation and continues executing; the result is handled via a callback/promise when ready.

**Q13: What are some core modules in Node.js?**
`fs` (file system), `http`/`https` (server), `path` (file paths), `os` (operating system info), `events` (event emitter), `crypto` (cryptography), `util` (utilities).

**Q14: What is the use of the `fs` module?**
It provides an API for interacting with the file system. It supports both synchronous (`fs.readFileSync`) and asynchronous (`fs.readFile`) methods for reading, writing, updating, and deleting files.

**Q15: How do you read a file asynchronously in Node.js?**
Using `fs.readFile('path', 'utf8', (err, data) => { ... })` or the promise-based version `await fs.promises.readFile('path', 'utf8')`.

**Q16: How can you handle errors in Node.js?**
Using `try...catch` blocks for synchronous code/async-await, `.catch()` for Promises, and error-first callbacks `(err, result)`. Global handlers include `process.on('uncaughtException')` (use with caution).

**Q17: What is the difference between `process.nextTick()` and `setImmediate()`?**
`process.nextTick()` fires immediately after the current operation finishes, *before* the event loop continues. `setImmediate()` fires in the specific "Check" phase of the upcoming event loop iteration.

**Q18: What is a global object in Node.js?**
Objects available in all modules without importing (`global`, `process`, `console`, `Buffer`, `__dirname`, `__filename`, `setTimeout`, etc.).

**Q19: What is the difference between `__dirname` and `./`?**
`__dirname` is the absolute path to the directory of the current file. `./` is relative to the directory from which the `node` command was run (process.cwd()), except in `require()`, where it is relative to the file.

**Q20: What is the `process` object in Node.js?**
A global object providing information about and control over the current Node.js process (e.g., environment variables, arguments, memory usage, exit methods).

## 🔹 Intermediate Level

**Q21: What is the difference between synchronous and asynchronous functions?**
Synchronous functions block execution until they complete. Asynchronous functions return immediately and handle the result later via callbacks or promises, allowing other code to run in the meantime.

**Q22: What is a Promise in Node.js?**
An object representing the eventual completion (or failure) of an asynchronous operation and its resulting value. It has states: Pending, Fulfilled, or Rejected.

**Q23: How do you convert a callback to a promise?**
Manually by wrapping it in `new Promise((resolve, reject) => { ... })`, or using the `util.promisify` utility provided by Node.js.

**Q24: What are async/await keywords? How do they work?**
Syntactic sugar over Promises. `async` makes a function return a Promise. `await` pauses the function execution until the Promise resolves, making asynchronous code look and behave like synchronous code.

**Q25: What are streams in Node.js?**
Objects that let you read data from a source or write data to a destination in continuous chunks, rather than loading the entire dataset into memory at once.

**Q26: Explain the four types of streams in Node.js.**
1.  **Readable**: For reading (e.g., `fs.createReadStream`).
2.  **Writable**: For writing (e.g., `fs.createWriteStream`).
3.  **Duplex**: Both readable and writable (e.g., TCP sockets).
4.  **Transform**: Modifies data as it is written/read (e.g., `zlib.createGzip`).

**Q27: What is backpressure in streams?**
A mechanism to handle data buildup when the Readable stream produces data faster than the Writable stream can consume it. Node.js pauses the reading until the buffer clears.

**Q28: How can you create a custom readable stream?**
By extending the `Readable` class from the `stream` module and implementing the `_read()` method to push data.

**Q29: How does error handling differ between callbacks and promises?**
Callbacks require checking `if (err)` in every callback. Promises allow chaining `.catch()` to handle errors from a sequence of asynchronous steps. Async/await relies on `try...catch`.

**Q30: What is middleware in Express.js?**
Functions that have access to the request (`req`), response (`res`), and the `next` middleware function. They can execute code, modify req/res, or end the cycle.

**Q31: What is the role of the `next()` function in Express middleware?**
It passes control to the next middleware function in the stack. If not called, the request hangs.

**Q32: How can you handle 404 errors in Express?**
By adding a middleware function at the very end of your route definitions that catches any request not matched by previous routes.

**Q33: How do you configure environment variables in Node.js?**
Using the `process.env` object. Typically, a `.env` file is used in development (loaded via `dotenv`) and system variables are set in production.

**Q34: What is `dotenv` and how is it used?**
A module that loads environment variables from a `.env` file into `process.env`. usage: `require('dotenv').config()`.

**Q35: How do you handle file uploads in Node.js?**
Using middleware like `multer` (for `multipart/form-data`) or `express-fileupload`.

**Q36: How can you serve static files using Express?**
Using the built-in middleware: `app.use(express.static('public'))`.

**Q37: What is the `cluster` module?**
Allows you to create child processes (workers) that share the same server port, enabling Node.js to utilize multi-core systems.

**Q38: What is the `child_process` module used for?**
To spawn new processes to execute system commands or other scripts. Methods include `spawn`, `exec`, `execFile`, and `fork`.

**Q39: What is the difference between `spawn` and `exec`?**
`spawn` streams output (good for large data/long processes). `exec` buffers the output (simple, but can crash if output exceeds buffer size).

**Q40: How can you handle CORS in a Node.js application?**
Using the `cors` middleware package (`app.use(cors())`) or manually setting `Access-Control-Allow-Origin` headers.

## 🔹 Advanced Level

**Q41: Explain the concept of event-driven programming.**
A paradigm where flow is determined by events (mouse clicks, messages, I/O completion). Node.js uses an Event Emitter pattern to listen for and react to events asynchronously.

**Q42: What is the libuv library?**
A multi-platform support library with a focus on asynchronous I/O. It implements the Node.js event loop, thread pool, and file/network I/O handling.

**Q43: What are the phases of the event loop?**
Timers -> Pending Callbacks -> Idle, Prepare -> Poll (I/O) -> Check (`setImmediate`) -> Close Callbacks.

**Q44: What is the microtask queue vs. the macrotask queue?**
Microtasks (Promises, `process.nextTick`) are executed immediately after the current operation and before the next event loop phase. Macrotasks (`setTimeout`, I/O) are executed in their respective event loop phases.

**Q45: How does Node.js handle concurrency?**
It uses a single-threaded event loop for JavaScript execution but offloads blocking I/O tasks to the multithreaded `libuv` thread pool (C++ side).

**Q46: What are worker threads?**
A module (`worker_threads`) that enables the use of threads that execute JavaScript in parallel. Useful for CPU-intensive tasks (e.g., image processing, cryptography) which would otherwise block the main loop.

**Q47: What is the difference between `fork()` and `spawn()`?**
`spawn()` is for running any command. `fork()` is a special case of spawn specifically for Node.js modules; it establishes an IPC (Inter-Process Communication) channel between parent and child.

**Q48: What is memory leak and how do you find it in Node.js?**
When objects are no longer needed but are referenced, preventing Garbage Collection. Found using tools like Chrome DevTools (Inspection), `heapdump`, or simply observing increasing RAM usage over time.

**Q49: How can you improve the performance of a Node.js app?**
Use caching (Redis), enable Gzip compression, optimize queries, use clustering, offload CPU tasks to worker threads, and use a reverse proxy (Nginx).

**Q50: How do you scale a Node.js application?**
Vertical Scaling (more RAM/CPU), Horizontal Scaling (Cluster module, Load Balancer + multiple instances), and decomposing into Microservices.

**Q51: What is load balancing?**
Distributing incoming network traffic across multiple servers to ensure no single server becomes overwhelmed. Nginx or Cloud Load Balancers are common.

**Q52: How can you handle uncaught exceptions?**
Using `process.on('uncaughtException', (err) => { ... })`. Ideally, you should log the error and restart the process, as the state might be corrupted.

**Q53: What is the difference between `try/catch` and `process.on('uncaughtException')`?**
`try/catch` handles local, expected errors within a block. `uncaughtException` is a global failsafe for errors that bubbled up without being caught anywhere else.

**Q54: How do you debug a Node.js application?**
Using `console.log`, the `debugger` keyword, VS Code built-in debugger, or `node --inspect` with Chrome DevTools.

**Q55: How does garbage collection work in Node.js?**
Node.js (via V8) uses "Mark-and-Sweep". It periodically traverses the object graph from root; reachable objects are marked "alive", unreachable ones are swept (freed). It has New Space (young generation) and Old Space (old generation).

**Q56: What are the best practices for securing a Node.js app?**
Use Helmet (headers), validate input, implement rate limiting, sanitize data (prevent NOSQL Injection/XSS), use HTTPS, and keep dependencies updated.

**Q57: What are some common vulnerabilities in Node.js apps?**
Buffer Overflow, ReDoS (Regex Denial of Service), Prototype Pollution, Dependency vulnerabilities, Implementation flaws (auth bypass).

**Q58: What is the difference between `Buffer` and `Stream`?**
Buffer is a temporary storage spot for a chunk of data (fixed size). Stream is a sequence of data moving from one point to another over time (potentially infinite).

**Q59: What is the `zlib` module used for?**
For compression and decompression (gzip/deflate). Commonly used to compress HTTP responses to save bandwidth.

**Q60: How does HTTP/2 differ from HTTP/1.1 in Node.js?**
HTTP/2 supports multiplexing (multiple requests over single connection), header compression, and server push, making it much faster. Node.js supports it via the `http2` module.

## 🔹 API and Web Development (Express.js)

**Q61: What is Express.js?**
A minimal and flexible Node.js web application framework that provides a robust set of features for web and mobile applications, simplifying server creation and routing.

**Q62: How do you create a basic Express server?**
`const express = require('express'); const app = express(); app.listen(3000);`

**Q63: How do you define routes in Express?**
`app.get('/', (req, res) => { ... })`, `app.post('/api/user', handleUser)`, etc.

**Q64: What is a route parameter?**
Named URL segments used to capture values. E.g., `/users/:id`. Accessible via `req.params.id`.

**Q65: What is the difference between `app.use()` and `app.get()`?**
`app.use()` mounts middleware for *all* HTTP methods (GET, POST, etc.) on a path. `app.get()` only handles GET requests for that specific path.

**Q66: How do you handle request body data in Express?**
Use middleware: `app.use(express.json())` for JSON and `app.use(express.urlencoded({ extended: true }))` for form data.

**Q67: What is body-parser and is it still needed?**
It was a separate package to parse bodies. Since Express 4.16, it is built-in as `express.json()`, so the separate package is generally not needed.

**Q68: How do you implement CORS in Express?**
`const cors = require('cors'); app.use(cors());`. This enables Cross-Origin Resource Sharing headers.

**Q69: How do you structure a RESTful API in Node.js?**
Organize by resources (e.g., users, products). Use standard HTTP methods. Separation of concerns: Routes -> Controllers -> Services -> Models.

**Q70: How do you return status codes in Express?**
`res.status(404).json({ error: 'Not Found' });`. Default is 200.

**Q71: What are some status codes you commonly use?**
200 (OK), 201 (Created), 400 (Bad Request), 401 (Unauthorized), 403 (Forbidden), 404 (Not Found), 500 (Server Error).

**Q72: How do you handle authentication in a Node.js API?**
Commonly using JWT (stateless) or Sessions (stateful) with libraries like Passport.js.

**Q73: What is JWT (JSON Web Token)?**
A compact, URL-safe means of representing claims to be transferred between two parties. Used for stateless authentication; the server verifies the token signature without checking a database.

**Q74: How do you protect routes with JWT in Express?**
Create a middleware that extracts the token from the Header, verifies it using `jsonwebtoken.verify()`, and calls `next()` if valid, or returns 401 if not.

**Q75: How do you connect a Node.js app to MongoDB?**
Using the `mongodb` native driver or an ODM like `mongoose`. `mongoose.connect('mongodb://localhost/myapp')`.

## 🔹 Testing and Tools

**Q76: How do you test a Node.js application?**
Unit tests (testing functions), Integration tests (API endpoints), and E2E tests. Popular frameworks: Jest, Mocha, Chai.

**Q77: What is Mocha?**
A flexible test framework for Node.js. It runs tests but requires an assertion library (like Chai) to verify results.

**Q78: What is Chai?**
An assertion library often paired with Mocha. It allows `expect(foo).to.be.a('string')` style assertions.

**Q79: What is Jest and how is it different from Mocha?**
Jest is an "all-in-one" framework (runner + assertion + mocks + coverage) developed by Facebook. Mocha is modular and requires other libs.

**Q80: What are mocks and stubs?**
**Mocks**: Fake objects checking if specific calls were made. **Stubs**: Fake functions that return predefined data to isolate the unit being tested.

**Q81: How do you test async code?**
In Jest/Mocha, you can return the Promise, use `async/await` in the test function, or use the `done` callback.

**Q82: How do you test an Express route?**
Using `supertest`. It simulates HTTP requests to your app instance without starting the actual server port.

**Q83: What is Supertest?**
A library for testing Node.js HTTP servers. It provides a high-level abstraction for sending requests and asserting responses.

**Q84: What is the purpose of ESLint in Node.js development?**
To identify and report on patterns found in ECMAScript/JavaScript code (linting). It enforces code style and catches potential errors.

**Q85: What are some useful Node.js debugging tools?**
Chrome DevTools (`node --inspect`), VS Code Debugger, `debug` module (namespace logging), `nodemon` (auto-restart).

## 🔹 Database Integration

**Q86: How do you connect to MongoDB using Mongoose?**
`const mongoose = require('mongoose'); mongoose.connect(uri, options).then(() => console.log('Connected'));`

**Q87: What is a schema in Mongoose?**
A configuration object that defines the structure, data types, validators, and default values for documents in a MongoDB collection.

**Q88: What are Mongoose models?**
Constructors compiled from Schemas. An instance of a model represents a document. The Model handles database operations (find, save, etc.).

**Q89: How do you handle relationships in MongoDB using Mongoose?**
Using `.populate()`. You store the `ObjectId` of one document in another and reference the model name in the Schema (`ref: 'User'`).

**Q90: How do you handle transactions in MongoDB?**
Using Sessions. `const session = await mongoose.startSession(); session.startTransaction(); ... await session.commitTransaction();` (Requires Replica Set).

**Q91: What is the difference between `find()` and `findOne()`?**
`find()` returns an array of matching documents (or empty array). `findOne()` returns the *first* matching document object (or null).

**Q92: How do you implement pagination in a Node.js API?**
Using `.skip((page - 1) * limit).limit(limit)` in your database query.

## 🔹 DevOps & Deployment

**Q93: How do you deploy a Node.js app?**
Transfer code to server, install dependencies (`npm install`), set env vars, build (if needed), start with a process manager (PM2), and use a reverse proxy (Nginx).

**Q94: What is PM2 and how is it used?**
A production process manager for Node.js. It keeps apps alive (restarts on crash), enables clustering, and manages logs. `pm2 start app.js`.

**Q95: How do you manage environment-specific configurations?**
Using different `.env` files (`.env.development`, `.env.production`) or configuration libraries like `config` that load based on `NODE_ENV`.

**Q96: How do you use Docker with Node.js?**
Create a `Dockerfile` specifying the node base image, copying `package.json`, installing deps, copying code, and exposing the port. Build and run the container.

**Q97: What is the benefit of using a `.env` file?**
Security and flexibility. It keeps secrets (API keys, DB URLs) out of the codebase (git) and allows easy configuration changes between environments.

**Q98: How do you monitor a production Node.js application?**
Using APM tools (New Relic, Datadog), structured logging (Winston + ELK stack), and metric monitoring (Prometheus + Grafana).

**Q99: What cloud services can host Node.js apps?**
AWS (EC2, Lambda, Elastic Beanstalk), Google Cloud (App Engine, Cloud Run), Heroku, Vercel (for frontend/serverless), DigitalOcean.

**Q100: How do you ensure zero-downtime deployment?**
Using a load balancer with rolling updates (e.g., Kubernetes Deployment, PM2 Cluster mode reload) where new instances start before old ones stop.

## 🔹 Core Node.js (Conceptual) - Batch 2

**Q101: Why is Node.js single-threaded?**
To simplify concurrency management and avoid the complexities/overhead of thread creation and context switching. It handles concurrency via the Event Loop instead of threads.

**Q102: Can Node.js be used for CPU-intensive tasks?**
Traditionally no, as efficient processing blocks the single thread. However, with `Worker Threads` (introduced in v10.5), CPU-heavy tasks can now be parallelized effectively.

**Q103: How does Node.js differ from traditional multithreaded servers like Apache?**
Apache creates a new thread for every request (high memory usage per client). Node.js uses a single thread to handle thousands of concurrent connections using non-blocking I/O (low memory overhead).

**Q104: What are the benefits of using Node.js?**
High performance (V8), Scalability (Non-blocking I/O), Unified Stack (JS on client/server), Large Ecosystem (npm), and Active Community.

**Q105: What is event delegation in Node.js?**
While meaningful in the browser (DOM), in Node.js, it refers to the pattern of using a single Event Emitter to handle events for multiple sources or dynamically handling events based on type.

**Q106: How does the `require` cache work?**
When a module is first loaded, it is cached in `require.cache`. Subsequent `require()` calls return the cached object instead of reloading the file, boosting performance.

**Q107: What is the difference between `require()` and `import` in terms of execution time?**
`require()` is synchronous and executes code at runtime. `import` (ES Modules) is asynchronous and statically analyzed during the parsing phase, optimization/tree-shaking before execution.

**Q108: Can you overwrite a global variable in Node.js?**
Yes, but it is bad practice. `global.something = 'new value'` affects the entire application and can lead to unpredictable behavior and collisions.

**Q109: What is a REPL in Node.js?**
Read-Eval-Print Loop. It's an interactive shell (`node` in terminal) to execute JavaScript code immediately. Useful for testing and debugging.

**Q110: What are timers in Node.js and how do they work?**
Functions that schedule code execution at a future time (`setTimeout`, `setInterval`, `setImmediate`). They are handled in specific phases of the Event Loop (`Timers` and `Check` phases).

## 🔹 File System & Buffer

**Q111: How do you watch a file for changes in Node.js?**
Using `fs.watch(filename, callback)` or `fs.watchFile(filename, callback)`. `fs.watch` is more efficient as it uses OS-native events.

**Q112: What’s the difference between `fs.readFile()` and `fs.createReadStream()`?**
`readFile` loads the *entire* file into memory (bad for huge files). `createReadStream` reads the file in chunks (buffers), keeping memory usage low.

**Q113: What are buffers, and how are they different from arrays?**
Buffers are fixed-length chunks of memory allocated outside the V8 JavaScript heap, used to handle binary data (like images, TCP streams). Arrays are JS objects for storing generic data types.

**Q114: How do you convert a buffer to a string?**
`buffer.toString(encoding)`. Default encoding is 'utf8'. E.g., `buf.toString('hex')`.

**Q115: What is the default encoding for buffers?**
UTF-8. If not specified, Node.js assumes UTF-8 when converting buffers to strings.

## 🔹 Event Loop & Timers

**Q116: What are the phases of the Node.js event loop?**
1. Timers (`setTimeout`)
2. Pending Callbacks (I/O errors)
3. Idle, Prepare
4. Poll (New I/O events, main code)
5. Check (`setImmediate`)
6. Close Callbacks (`socket.on('close')`)

**Q117: In what order are `setTimeout`, `setImmediate`, and `process.nextTick` executed?**
1. `process.nextTick` (Immediately, before loop continues)
2. `setTimeout` (Timers phase)
3. `setImmediate` (Check phase)
*Note: If called within an I/O cycle, `setImmediate` always runs before `setTimeout`.*

**Q118: What is the role of the `Timers` phase in the event loop?**
It executes callbacks scheduled by `setTimeout()` and `setInterval()`.

**Q119: When should you use `setImmediate()` over `setTimeout()`?**
Use `setImmediate()` when you want to execute a callback immediately after the current I/O cycle, guaranteeing it runs before any timers scheduled for the "same" time.

**Q120: Can you starve the event loop? How?**
Yes, by running a long synchronous task (e.g., a massive `while` loop or heavy calculation) on the main thread. This blocks the loop from moving to the Poll phase/handling I/O.

## 🔹 Streams (Advanced)

**Q121: What is the difference between flowing and paused mode in streams?**
**Flowing**: Data is read automatically and provided to callbacks.
**Paused**: You must explicitly call `stream.read()` to get data chunks.

**Q122: How do you handle stream errors?**
Streams emit an `'error'` event. You must attach a listener: `stream.on('error', (err) => { ... })`. If unhandled, it crashes the app.

**Q123: What is the `pipeline()` method used for?**
A utility in `stream` module (`stream.pipeline`) to pipe streams together and automatically handle errors and cleanup. It is safer than `.pipe()`.

**Q124: How do you implement a duplex stream?**
Inherit from `Duplex` and implement both `_read()` and `_write()` methods. It effectively acts as both a Readable and Writable stream.

**Q125: What are the key events emitted by streams?**
`data` (data available), `end` (no more data to read), `finish` (writing completed), `error` (something went wrong), `drain` (buffer cleared, ready for writing).

## 🔹 Networking & HTTP

**Q126: How do you create an HTTP server without Express?**
`const http = require('http'); const server = http.createServer((req, res) => { res.end('Hello'); }); server.listen(3000);`

**Q127: What is the difference between `http.createServer` and `https.createServer`?**
`https.createServer` requires an options object with SSL private key and certificate (`key` and `cert`) to enable encryption.

**Q128: How do you handle file downloads in Node.js?**
Set header `Content-Disposition: attachment; filename="file.ext"` and pipe the file stream to `res`.

**Q129: How do you implement keep-alive in HTTP requests?**
It is enabled by default in modern Node.js. You can explicitly use the `http.Agent({ keepAlive: true })` to reuse TCP connections for multiple requests.

**Q130: What is connection pooling?**
A technique where a set of active connections (to DB or API) is reused instead of opening a new one for every request, reducing handshake overhead.

## 🔹 Express.js (Advanced)

**Q131: What is the role of the `app.all()` method?**
It matches *all* HTTP methods (GET, POST, PUT, DELETE, etc.) for a specific route. Useful for global logic like validation or logging on a path.

**Q132: How can you handle different HTTP methods in the same route?**
By chaining handlers: `app.route('/book').get(getBooks).post(addBook).put(updateBook)`.

**Q133: How do you organize routes in Express using routers?**
Use `express.Router()`. Create separate router files for different features (`users`, `products`), then mount them: `app.use('/users', userRouter)`.

**Q134: What is the order of middleware execution?**
Sequential, following the order they are defined (`app.use()`). If a middleware doesn't call `next()`, the chain stops.

**Q135: How do you create a custom middleware?**
`const logger = (req, res, next) => { console.log(req.method); next(); }; app.use(logger);`

## 🔹 Authentication & Security

**Q136: What is CSRF and how do you prevent it in Node.js?**
Cross-Site Request Forgery. Preventing it requires CSRF tokens (using libraries like `csurf`) that verify the request originated from your own trusted frontend.

**Q137: How do you hash passwords in Node.js?**
Using libraries like `bcrypt` or `argon2`. Never store plain text. `const hash = await bcrypt.hash(password, 10);`

**Q138: What is OAuth2 and how can you implement it?**
An authorization framework enabling third-party login (Google, Facebook). Implement using `Passport.js` with strategies like `passport-google-oauth20`.

**Q139: What is Helmet.js?**
A middleware package that sets various HTTP headers (Strict-Transport-Security, X-Frame-Options, etc.) to secure Express apps against common attacks.

**Q140: How can you restrict API access by IP address?**
Middleware checking `req.ip`. `if (!allowedIps.includes(req.ip)) return res.status(403).send('Forbidden');`.

## 🔹 Performance & Optimization

**Q141: What is lazy loading in Node.js?**
Deferring the loading of modules or execution of code until it is actually needed (e.g., using `require()` inside a function). Reduces startup time.

**Q142: How do you implement memoization in Node.js?**
Caching the result of expensive function calls based on arguments. Can be done manually using an Object/Map or helper libraries like `lodash.memoize`.

**Q143: What is the role of caching in Node.js apps?**
Stores frequently accessed data in memory (Redis/Memcached) to reduce database load and improve response time.

**Q144: How do you use Redis with Node.js?**
Using a client library like `ioredis` or `node-redis`. `await redis.set('key', 'value'); const val = await redis.get('key');`.

**Q145: How can you avoid blocking the event loop?**
Use asynchronous APIs (Promises/Async-Await), offload CPU-heavy tasks to Worker Threads, or break long tasks into smaller chunks using `setImmediate`.

## 🔹 Testing (Advanced)

**Q146: How do you test middleware in Express?**
Call the middleware function directly with mock `req`, `res`, and `next` objects/spies, and assert that `res` methods or `next` were called correctly.

**Q147: How do you mock database calls in tests?**
Using libraries like `sinon`, `nock` (for HTTP), or Jest mocks (`jest.mock('mongoose')`) to intercept calls and return fake data.

**Q148: What is code coverage and how do you measure it?**
A metric showing what percentage of your code is executed by tests. Jest has built-in coverage (`jest --coverage`).

**Q149: How do you perform integration testing in Node.js?**
Spin up the app connected to a test database (e.g., Docker container), run HTTP requests against it (`supertest`), and verify the database state.

**Q150: What is TDD and how does it apply to Node.js?**
Test-Driven Development. Write the test first (it fails), write the minimal code to pass it, then refactor.

## 🔹 TypeScript with Node.js

**Q151: What are the benefits of using TypeScript with Node.js?**
Static typing (catches errors at compile time), better tooling/IntelliSense, interface definitions, and modern JavaScript features.

**Q152: How do you set up a Node.js project with TypeScript?**
`npm install typescript ts-node @types/node`. Create `tsconfig.json`. Compile with `npx tsc` or run with `ts-node`.

**Q153: How do you declare types for Express request and response objects?**
Use generics or extend interfaces: `import { Request, Response } from 'express'; func(req: Request, res: Response)`.

**Q154: How do you handle module resolution issues in TypeScript?**
Configure `compilerOptions.moduleResolution: "node"` in `tsconfig.json` and ensure `@types` packages are installed.

**Q155: What is the `tsconfig.json` file used for?**
It defines the root of the project and the compiler options (target version, strict mode, output directory, included files) for TypeScript.

## 🔹 Databases & ORMs

**Q156: What is an ORM, and name a few Node.js ORMs?**
Object-Relational Mapper. Maps DB tables to Objects. Examples: Sequelize, TypeORM, Prisma, Mongoose (ODM).

**Q157: How does Sequelize differ from TypeORM?**
Sequelize uses the Active Record pattern. TypeORM supports both Active Record and Data Mapper patterns and has better TypeScript support.

**Q158: How do you define associations in Sequelize?**
`User.hasMany(Post); Post.belongsTo(User);`.

**Q159: How do you perform transactions in Sequelize?**
`await sequelize.transaction(async (t) => { ... }, { transaction: t });`.

**Q160: How can you write raw SQL queries in a Node.js app?**
`const [results, metadata] = await sequelize.query("SELECT * FROM users");`

## 🔹 Architecture & Patterns

**Q161: What is a service layer in a Node.js application?**
A layer separate from the Controller that contains the business logic. Controllers handle HTTP; Services handle logic/data.

**Q162: What is dependency injection and how is it implemented in Node.js?**
Providing dependencies to a component (class/function) rather than hardcoding them. Libraries like `InversifyJS` or `NestJS` native DI are common.

**Q163: What is the factory pattern and when would you use it?**
A function that creates and returns objects. Useful for creating similar objects without classes or when object creation logic is complex.

**Q164: How would you implement a queue system in Node.js?**
Using a message broker like `RabbitMQ` or `Redis` (with libraries like `Bull`). Producers add jobs; Consumers process them.

**Q165: What is a monorepo and how does it apply to Node.js apps?**
A single git repository containing multiple projects (backend, frontend, shared libs). Tools like `Nx`, `Lerna`, or `Turborepo` manage dependencies and builds.

## 🔹 Real-Time Applications

**Q166: What is WebSocket and how is it used in Node.js?**
A protocol providing full-duplex communication channels over a single TCP connection. The `ws` library or `socket.io` are used to implement it.

**Q167: How does Socket.IO differ from WebSocket?**
Socket.IO is a library built on top of WebSockets. It provides fallbacks (polling), auto-reconnection, and features like Rooms/Namespaces.

**Q168: How do you handle rooms and namespaces in Socket.IO?**
Namespaces (`io.of('/chat')`) isolate logic. Rooms (`socket.join('room1')`) group sockets for broadcasting messages (`io.to('room1').emit(...)`).

**Q169: What are long polling and server-sent events?**
**Long Polling**: Client requests -> Server holds until data available -> Responds -> Client requests again.
**SSE**: Server pushes updates to client over HTTP (unidirectional).

**Q170: How do you implement a pub/sub system in Node.js?**
Using Redis Pub/Sub (`subscribe`/`publish`) or Node's internal `EventEmitter` for local processes.

## 🔹 Microservices & Messaging

**Q171: What is a microservice?**
Small, independent services focusing on a specific business capability, communicating via APIs/Events, often deployed separately.

**Q172: How would you build microservices using Node.js?**
Using Express/Fastify for each service. Communication via REST or gRPC. Integration via Event Bus (RabbitMQ/Kafka). Orchestration via Kubernetes.

**Q173: What is gRPC and how is it used in Node.js?**
Google Remote Procedure Call. Uses Protocol Buffers (fast, binary). Useful for efficient inter-service communication. `@grpc/grpc-js`.

**Q174: What is RabbitMQ and how do you use it with Node.js?**
A message broker. Use the `amqplib` library to create channels, publish messages to exchanges, and consume from queues.

**Q175: What is the role of API Gateway in microservices?**
A single entry point for clients. It routes requests to backend microservices, handles cross-cutting concerns (auth, rate limiting, logging).

## 🔹 CI/CD & Tooling

**Q176: How do you write a GitHub Action to test a Node.js app?**
Create `.github/workflows/test.yml`. Use `actions/setup-node`, run `npm install`, then `npm test`.

**Q177: How do you use `nodemon` in development?**
`nodemon app.js`. It monitors file changes and automatically restarts the node process.

**Q178: How do you use `lint-staged` with Node.js?**
Run linters only on committed files. Configure in `package.json` to run ESLint/Prettier on pre-commit hooks (via `husky`).

**Q179: What is semantic versioning?**
Versioning format: `Major.Minor.Patch` (e.g., 1.0.0). Major = Breaking changes. Minor = New features (backward compatible). Patch = Bug fixes.

**Q180: How do you build and publish a Node.js package?**
`npm init`, write code, authentication (`npm login`), `npm publish`. Ensure `.npmignore` excludes unnecessary files.

## 🔹 Error Handling & Logging

**Q181: How do you implement centralized error handling in Express?**
Pass errors to `next(err)`. Create a global error middleware function `(err, req, res, next) => { ... }` at the end of the app.

**Q182: What is the difference between `throw` and `next(err)`?**
`throw` stops execution immediately (works in sync code/async-await). `next(err)` explicitly passes the error to Express error handlers (needed in callbacks).

**Q183: What is Winston and how is it used?**
A versatile logging library. Supports multiple transports (Console, File, Http). `winston.createLogger({ transports: [...] })`.

**Q184: How do you log exceptions in Node.js?**
Catch `uncaughtException` and `unhandledRejection`. Use a logger like Winston to write stack traces to a file/service before exiting.

**Q185: How can you integrate Sentry with Node.js?**
`Sentry.init({ dsn: '...' })`. Use `Sentry.Handlers.requestHandler()` and `Sentry.Handlers.errorHandler()` middleware in Express.

## 🔹 ES Modules & Modern JavaScript in Node

**Q186: How do you enable ES modules in Node.js?**
Set `"type": "module"` in `package.json` or use `.mjs` extension. Allows `import`/`export` syntax.

**Q187: What’s the difference between CommonJS and ES Modules?**
CommonJS: `require`/`module.exports`, Synchronous.
ESM: `import`/`export`, Asynchronous loading, strict mode by default.

**Q188: Can you use both ES modules and CommonJS together?**
Yes, but with caveats. Provide exports in CommonJS, import them in ESM. Directly importing ESM into CommonJS is harder (requires dynamic `import()`).

**Q189: What is tree-shaking and how does it work in Node?**
Removing unused code during the build process (by bundlers like Webpack/Rollup). Works best with static ESM `import` statements.

**Q190: How do you use top-level `await` in Node.js?**
In ES Modules (Node 14.8+), you can use `await` outside of async functions at the root level of a module.

## 🔹 Advanced Node Features

**Q191: What is async_hooks in Node.js?**
A core module to track the lifetime of asynchronous resources. Useful for creating APMs or request context tracing (CLS).

**Q192: What are V8 snapshots?**
A way to improve startup time. It allows creating a serialized heap state that can be loaded instantly instead of re-compiling standard libraries.

**Q193: What is the inspector module used for?**
To interact with the V8 inspector programmatically. Allows creating profiling tools and capturing CPU profiles/heap dumps on the fly.

**Q194: How do you use perf_hooks to measure performance?**
`performance.now()` for high-resolution timing. `PerformanceObserver` to monitor performance entries like GC, loop delay, or function execution.

**Q195: What is an AbortController in Node.js?**
A global utility (from web standards) to abort one or more Web requests or asynchronous tasks. `controller.abort()`.

## 🔹 Edge Cases & Troubleshooting

**Q196: Why might a `require()` fail even if the file exists?**
Permission issues, circular dependencies resulting in partial exports, invalid syntax in the required file, or Case Sensitivity on Linux (e.g., `File.js` vs `file.js`).

**Q197: What happens if you listen on a port already in use?**
Node.js throws an `EADDRINUSE` error. You must handle it or the process crashes.

**Q198: What are zombie processes and how do you prevent them?**
Child processes that terminated but are still in the process table (parent didn't wait for them). Node handles this automatically on `child_process` exit, but signal handling is key.

**Q199: How do you detect and handle memory leaks?**
Monitor memory usage (`process.memoryUsage()`). Use `heapdump` to take snapshots. Compare snapshots in Chrome DevTools to find detached objects.

**Q200: How do you troubleshoot high CPU usage in a Node.js app?**
Generate a CPU profile (`node --prof` or Inspector). Analyze flame graphs. Often caused by sync operations, heavy loops, or inefficient RegEx (ReDoS).

## 🔹 System Design & Architecture - Batch 3

**Q201: How would you structure a large Node.js project?**
Feature-based structure (grouping by domain: `components/User`, `components/Order`) rather than Technical structure (`controllers`, `models`). Use layered architecture: Route -> Controller -> Service -> Data Access/Model.

**Q202: What is a middleware pipeline, and how do you design one?**
A sequence of functions that request passes through. Design it by ordering logic broadly to specifically: Global functionality (Logging/CORS) -> Parsers -> Auth -> Business Logic -> Error Handling.

**Q203: How do you design a modular Express.js app?**
By decoupling features into separate modules with their own routes and logic, mounted on the main app. Use `express.Router()` for each module.

**Q204: What is a repository pattern in Node.js?**
An abstraction layer between the business logic and the database. It isolates data access code (`findUser`, `saveUser`) so the service layer doesn't depend directly on the database driver/ORM.

**Q205: How would you implement caching at the route level?**
Create middleware that checks if the key (e.g., URL) exists in Redis. If yes, return cached data. If no, call `next()`, and intercept the response `send()` method to cache the result before sending.

**Q206: When would you use event emitters over REST APIs?**
For internal communication between decoupled components within the same process, or when multiple parts of the app need to react to a single action (Publisher-Subscriber pattern) without tight coupling.

**Q207: What is an API gateway, and how do you build one with Node.js?**
A single entry point that aggregates/routes requests to microservices. Built using libraries like `http-proxy-middleware` or `express-gateway`, handling routing, auth, and rate limiting.

**Q208: How do you manage versioning in a Node.js API?**
URI Versioning (`/v1/users`), Header Versioning (`Accept-Version: v1`), or Parameter Versioning (`?v=1`). URI versioning is most common for public APIs.

**Q209: What are idempotent APIs, and how would you implement them?**
APIs where making the same request multiple times has the same effect as making it once. Implement using an "Idempotency-Key" header stored in a shared cache (Redis) to prevent processing duplicate requests.

**Q210: How do you build a multi-tenant application in Node.js?**
By isolating data per tenant. Either separate databases per tenant or a shared database with a `tenant_id` column. Middleware identifies the tenant (from subdomain/header) and sets the context.

## 🔹 Deployment & Infrastructure

**Q211: What is the difference between `pm2 restart` and `pm2 reload`?**
`restart`: Kills the process and starts it again (downtime). `reload`: Restarts processes one by one (zero-downtime) if running in cluster mode.

**Q212: How do you use NGINX as a reverse proxy for Node.js?**
Configure Nginx to listen on port 80/443 and inspect/forward requests to the Node.js app running on a local port (e.g., 3000) using `proxy_pass http://localhost:3000;`. Handles SSL, caching, and load balancing.

**Q213: How do you deploy a Node.js app to AWS ECS?**
Containerize the app (Docker), push image to ECR, define Task Definition (CPU/RAM/Image), and create a Service in an ECS Cluster to run the task.

**Q214: What is serverless architecture? Can Node.js be used in serverless?**
Running code without provisioning servers. Yes, Node.js is a primary runtime for AWS Lambda, Azure Functions, and Google Cloud Functions.

**Q215: How would you set up CI/CD for a Node.js monorepo?**
Use tools like `Turborepo` or `Nx` to detect changed workspaces. Configure the pipeline (GitHub Actions/Jenkins) to only build/test/deploy the affected packages.

**Q216: What are some common pitfalls in Dockerizing Node.js apps?**
Running as root (security risk), not using `.dockerignore` (`node_modules` copying), improper signal handling (PID 1 issue), and large image sizes (use `alpine`).

**Q217: How do you implement blue/green deployment in Node.js?**
Have two identical environments (Blue=Live, Green=Idle). Deploy new version to Green. Test Green. Switch load balancer to point to Green. Blue becomes idle.

**Q218: What are environment secrets, and how do you manage them securely?**
Sensitive config (API keys). Manage via cloud secret managers (AWS Secrets Manager, Vault) and inject them as Env Vars at runtime. Never commit `.env` files.

**Q219: What is the impact of `NODE_ENV=production`?**
It optimizes build performance, disables debug logging, enables caching in Express templates, and is often a flag for libraries to switch to "fast mode".

**Q220: How would you deploy a Node.js app to Kubernetes?**
Create a Docker image, push to registry. Write K8s manifests: `Deployment` (replicas, image), `Service` (networking), `Ingress` (external access), and `ConfigMap`/`Secret`. Apply with `kubectl`.

## 🔹 Tooling & Productivity

**Q221: How do you use `concurrently` in a Node.js project?**
A package to run multiple commands simultaneously. `concurrently "npm run server" "npm run client"`. Useful for dev environments with backend/frontend.

**Q222: What does `nvm` do?**
Node Version Manager. Allows installing and switching between multiple versions of Node.js on the same machine.

**Q223: How do you benchmark a Node.js app?**
Using tools like `autocannon` or `Apache Bench (ab)` to simulate high load and measure latency/throughput. `autocannon -c 100 -d 10 http://localhost:3000`.

**Q224: What is the purpose of `.npmrc`?**
Configuration file for npm. Used to set registry URLs (private repos), auth tokens, or lock dependencies behavior.

**Q225: How do you analyze Node.js memory usage?**
Using `process.memoryUsage()`, Chrome DevTools Memory tab (Heap Snapshot), or `v8.getHeapStatistics()`.

**Q226: What is the purpose of the `debug` package?**
A tiny debugging utility that allows toggling log output via the `DEBUG` environment variable. `DEBUG=app:db node app.js`.

**Q227: How do you configure path aliases in a Node.js + TypeScript project?**
Set `paths` in `tsconfig.json` (e.g., `"@/*": ["src/*"]`) and use `module-alias` or `tsconfig-paths` to resolve them at runtime.

**Q228: What is `husky` and how do you use it?**
A tool to manage Git hooks easily. Used to run scripts (lint, test) automatically before commit (`pre-commit`) or push.

**Q229: What is the difference between `npm ci` and `npm install`?**
`npm install` installs dependencies and updates `package-lock.json`. `npm ci` (Clean Install) deletes `node_modules` and installs *exactly* what is in `package-lock.json` (faster, for CI/CD).

**Q230: How do you lock down dependency versions in Node.js?**
Commit `package-lock.json`. For stricter control, remove `^` or `~` from `package.json` (pin exact versions) or use `npm shrimpwrap` (deprecated) / `overrides`.

## 🔹 Concurrency & Asynchrony

**Q231: How does the async model of Node.js compare to Go or Python?**
Node.js: Single-threaded Event Loop (Callbacks/Promises).
Go: Goroutines (Lightweight threads, multicore).
Python: Asyncio (Event loop) or Threading (GIL-limited).

**Q232: What is cooperative concurrency in Node.js?**
The idea that tasks must voluntarily yield control (finish execution) to allow the event loop to store other tasks. Long-running sync code violates this.

**Q233: What is the maximum number of concurrent I/O operations Node.js can handle?**
Theoretically limited by system resources (file descriptors, RAM), not the event loop. It can handle tens of thousands of concurrent connections.

**Q234: How would you simulate a deadlock in Node.js?**
Deadlocks (thread waiting for itself) are rare in single-threaded JS. However, "Logical Deadlock" occurs if two Promises wait for each other, or a stream waits for a 'drain' event that never happens.

**Q235: How can you safely perform parallel async operations?**
Using `Promise.all([p1, p2])` (fails fast) or `Promise.allSettled()` (waits for all).

**Q236: How do you throttle or debounce API calls in Node.js?**
**Throttle**: Ensure function runs at most once every X ms (rate limit).
**Debounce**: Ensure function runs only after X ms of silence (search input). Use `lodash`.

**Q237: How would you implement retry logic with backoff in Node.js?**
Write a recursive function or loop that catches errors, waits a delay (multiplied by factor each time), and retries the operation until max retries reached.

**Q238: Can Node.js handle millions of concurrent connections?**
Yes, with tuning (increasing `ulimit`, optimizing kernel TCP settings) and using high-performance web frameworks like Fastify or uWebSockets.js.

**Q239: What is the difference between `Promise.all` and `Promise.allSettled`?**
`Promise.all` rejects immediately if *any* promise rejects. `Promise.allSettled` waits for *all* to finish and provides status (fulfilled/rejected) for each.

**Q240: How would you build a task queue in memory?**
Using an array or linked list. Push tasks (functions/data) to it. Have a worker function `pop` and execute tasks, handling concurrency limits.

## 🔹 Advanced Error Handling

**Q241: How do you handle async/await errors globally?**
Use a wrapper function (`express-async-errors` does this) to catch rejected promises and pass them to `next(err)`, or use Node's `unhandledRejection` event.

**Q242: What are operational vs. programmer errors in Node.js?**
**Operational**: Runtime problems (validation failed, server timeout) - Handle them.
**Programmer**: Bugs (syntax, passing string to int function) - Fix the code, restart app.

**Q243: What is an unhandled rejection warning?**
A warning printed when a Promise is rejected but no `.catch()` handler is attached. In future Node versions, this will crash the process.

**Q244: What happens if you call `next()` multiple times in Express?**
It triggers an error: "Cannot set headers after they are sent to the client" if response was already started, or executes the next middleware twice (unexpected behavior).

**Q245: What’s the best way to handle 3rd-party API failures?**
Implement timeouts (do not wait forever), retries (transient errors), and Circuit Breaker pattern (stop calling if service is down) using libraries like `opossum`.

## 🔹 Node.js Internals & V8

**Q246: How do closures impact memory usage in Node.js?**
Closures retain references to their outer scope variables. If a closure lives long (e.g., in an event listener), the referenced variables cannot be garbage collected, potentially causing leaks.

**Q247: What is a "tick" in the Node.js event loop?**
Often refers to `process.nextTick`, which is a microtask execution pass that happens *between* operations, effectively "cutting the line" before the next loop phase.

**Q248: What is inline caching in the V8 engine?**
An optimization technique where V8 caches the results of property lookups on objects. If object shape doesn't change, V8 skips the lookup and goes directly to the memory address.

**Q249: How do hidden classes impact performance?**
V8 creates hidden classes (shapes) for objects. If you modify object properties dynamically (adding later), V8 creates new hidden classes provided optimization fails. Initialize all properties in constructor!

**Q250: What is the V8 heap limit and how can you increase it?**
The default memory limit (approx 1.5GB - 2GB on 64-bit). Increase it using the flag `--max-old-space-size=4096` (e.g., set to 4GB).

## 🔹 Third-Party Libraries

**Q251: What is the difference between Axios and `fetch()` in Node.js?**
`Axios`: Library, automatic JSON parsing, request interceptors, broad browser support.
`fetch()`: Native in Node 18+, standard API, requires manual `.json()` call.

**Q252: How does `multer` work for file uploads?**
It processes `multipart/form-data` streams. It parses the boundary strings, extracts file data to disk or memory, and makes it available via `req.file`.

**Q253: What are some alternatives to `express` for building APIs?**
Fastify (high performance), Koa (minimalist), Hapi (configuration-centric), NestJS (structured/modular).

**Q254: How does `jsonwebtoken` verify tokens?**
It decodes the header/payload, re-computes the signature using the provided secret key, and compares it with the token's signature. Checks expiry (`exp`) automatically.

**Q255: How does `passport.js` handle different strategies?**
It uses the Strategy Pattern. You configure a strategy (Local, Google, JWT), and Passport normalizes the auth flow, populating `req.user` on success.

**Q256: What is `rate-limiter-flexible` and how does it work?**
A powerful library for rate limiting. It uses a "Token Bucket" or "Fixed Window" algorithm, typically storing counters in Redis to limit actions by key (IP/ID).

**Q257: What is `node-cron` used for?**
A task scheduler written in pure JS for Node.js. Allows running functions on a schedule defined by cron syntax (`* * * * *`).

**Q258: How does `nodemailer` work with Gmail?**
It uses SMTP protocol. requires establishing a secure connection to `smtp.gmail.com` using OAuth2 (recommended) or App Passwords.

**Q259: How does `bull` or `agenda` manage background jobs?**
**Bull**: Uses Redis to store jobs (LIFO/FIFO/Priority).
**Agenda**: Uses MongoDB.
Both poll/listen for due jobs and execute them in processor functions.

**Q260: What’s the difference between `joi` and `zod` for validation?**
`Joi`: Older, schema description language, runtime only.
`Zod`: Newer, TypeScript-first, infers static types directly from the schema (Developer Experience win).

## 🔹 Authentication & Authorization (Advanced)

**Q261: How do you manage refresh tokens securely?**
Store refresh tokens in an `HttpOnly` cookie (prevents XSS). Access tokens (short-lived) can be in memory. On access token expiry, use refresh token to get a new one.

**Q262: What is the OAuth2 implicit flow?**
A simplified flow where the access token is returned directly in the URL fragment. Now deprecated/insecure. Use Authorization Code Flow with PKCE instead.

**Q263: How would you implement RBAC in Node.js?**
Role-Based Access Control. Assign roles (Admin, User) to users. Create middleware: `checkRole('admin')`. If `req.user.role !== 'admin'`, deny access.

**Q264: How do you handle multi-factor authentication (MFA)?**
After password login, require a second step: Verify a TOTP code (Time-based One-Time Password) generated by app (Google Authenticator) using library like `speakeasy`.

**Q265: How do you validate scopes or roles in a Node.js route?**
Middleware. `app.get('/admin', requireScope('read:admin'), handler)`. The middleware checks if the user's token scopes include the required string.

## 🔹 Real-Time Features

**Q266: What are race conditions in Socket.IO?**
When multiple events try to modify shared state simultaneously. Because Node is single-threaded, true parallel race is rare, but logic order matters (e.g., DB read-modify-write).

**Q267: How do you ensure message order in WebSockets?**
TCP ensures packet order. However, app logic must handle sequential processing. Using a queue or ensuring the client waits for ACK before sending the next message helps.

**Q268: How do you authenticate Socket.IO connections?**
Using middleware during the handshake. `io.use((socket, next) => { const token = socket.handshake.auth.token; verify(token)... })`.

**Q269: How do you build a real-time collaborative editor using Node.js?**
Using protocols like Operational Transformation (OT) or CRDTs (Conflict-free Replicated Data Types) to merge changes from multiple users consistent. `Yjs` is a popular library.

**Q270: What is the best way to implement presence detection?**
On socket connect: `redis.sadd('online_users', userId)`. On disconnect: `redis.srem(...)`. Periodically or on change event, broadcast list.

## 🔹 Monitoring & Observability

**Q271: What is the purpose of APM in Node.js?**
Application Performance Monitoring. Tools (New Relic, Dynatrace) that hook into Node.js internals to trace transactions, find slow database queries, and monitor error rates.

**Q272: How do you use OpenTelemetry with Node.js?**
Install the SDK and auto-instrumentation packages (`@opentelemetry/auto-instrumentations-node`). It automatically traces HTTP/DB calls and exports data to a backend (Jaeger/Prometheus).

**Q273: What metrics should you monitor in a Node.js app?**
Event Loop Lag (critical), CPU Usage, Memory (Heap Used), Active Handles/Requests, Garbage Collection pauses, HTTP Request Latency/Throughput.

**Q274: How do you track slow queries in Node.js?**
Use DB driver features (e.g., Mongoose `post` hooks with timer) or APM tools. Log any query taking > X ms.

**Q275: What is a flame graph and how is it useful?**
A visualization of the call stack. The width of bars represents time spent on CPU. It helps identify which functions are CPU bottlenecks (hot paths).

## 🔹 Security (Advanced)

**Q276: How do you prevent prototype pollution in Node.js?**
Freeze the prototype: `Object.freeze(Object.prototype)`. Validate JSON inputs explicitly. Avoid unsafe recursive merge functions. use distinct libraries for merging.

**Q277: What is `eval()` and why should you avoid it?**
Executes a string as code. Highly dangerous (`eval('rm -rf /')`). It prevents V8 optimizations and opens massive security holes.

**Q278: How does Node.js sanitize user input?**
It doesn't automatically. You must use libraries like `DOMPurify` (for HTML/XSS), `validator.js`, or parameterized queries (SQL injection) to sanitize.

**Q279: What is SSRF and how can Node.js apps be affected?**
Server-Side Request Forgery. If an attacker can control a URL fetched by your server (e.g., `fetch(req.query.url)`), they can access internal network services (AWS metadata, localhost).

**Q280: How do you audit Node.js dependencies for vulnerabilities?**
Run `npm audit`. Integrate `Snyk` or `GitHub Dependabot` into the CI/CD pipeline to block PRs adding vulnerable packages.

## 🔹 API Design

**Q281: How do you implement HATEOAS in a Node.js API?**
Hypermedia as the Engine of Application State. Include `links` in the JSON response: `{ "data": user, "links": { "self": "/users/1", "update": "/users/1" } }`.

**Q282: How do you design a bulk update endpoint?**
`PATCH /resources`. Body accepts an array of objects. Use a transaction to apply all or nothing. OR `PATCH /resources/:id` loop (inefficient).

**Q283: What is an idempotency key?**
A unique token sent by the client (e.g., UUID) with a POST request. The server tracks processed keys. If verified, it returns the stored response instead of re-processing payment/action.

**Q284: How do you handle file streaming over REST?**
Use `fs.createReadStream()` and pipe it to `res`. Set `Content-Type`. This transfers data in chunks without loading file to memory.

**Q285: How would you implement rate-limiting per user/IP?**
Use a sliding window algorithm backed by Redis. Incremented counter per key (`ratelimit:ip:127.0.0.1`). If > limit, return 429.

## 🔹 Data Handling & Serialization

**Q286: How do you serialize large JSON responses efficiently?**
Use streaming JSON stringifiers (like `JSONStream` or `bfj`) instead of `JSON.stringify()`, which blocks the event loop for large objects.

**Q287: What is the difference between `res.send()` and `res.json()`?**
`res.json()` forces proper JSON Content-Type header and formatting. `res.send()` infers type (String -> HTML, Object -> JSON, Buffer -> download).

**Q288: How do you handle circular JSON references?**
`JSON.stringify` throws an error. Use `flatted` or `json-stringify-safe` libraries to handle circular structures safely.

**Q289: How do you stream a large CSV file to a client?**
Read data from DB cursor -> Pipe to CSV Transformer (`fast-csv`) -> Pipe to `res`. Keeps memory footprint constant.

**Q290: What is binary protocol handling in Node.js?**
Using `Buffer` to parse headers/payloads of binary protocols (like MQTT, TCP custom protocols). Read bytes at specific offsets.

## 🔹 Date, Time & Timezones

**Q291: How do you handle timezones in a Node.js app?**
Store everything in UTC in the database. converting to local time only when displaying to the user on the client side.

**Q292: How would you handle scheduled tasks across timezones?**
Store user preferences (TimeZone). When checking if a task should run, calculate `UserLocalTime = CurrentUTC + Offset`. Or use libraries that support TZs like `cron` with timezone support.

**Q293: How do you use `luxon` or `dayjs` over `moment`?**
Moment.js is legacy/bloated (mutable). `Luxon`/`Day.js` are immutable, modern, and modular (smaller bundle). `dayjs(date).add(1, 'day')`.

**Q294: What is `process.uptime()` and when is it useful?**
Returns the number of seconds the process has been running. Useful for health checks and monitoring application stability.

**Q295: How do you benchmark time-sensitive code?**
Use `process.hrtime.bigint()` (High Resolution Time) to measure execution duration in nanoseconds accurately.

## 🔹 Miscellaneous / Scenarios

**Q296: How do you allow plugins in your Node.js application?**
Design a standard interface/API. Load plugins dynamically using `require()` from a `plugins` folder. Use an Event Emitter or Hook system to allow plugins to tap into lifecycle events.

**Q297: How do you handle large request payloads safely?**
Limit body size using body-parser: `express.json({ limit: '10kb' })`. Prevents DoS attacks from massive payloads crashing memory.

**Q298: How do you serve multiple domains from a single Node.js server?**
Check `req.hostname` in middleware. `if (req.hostname === 'a.com') handleA(); else handleB();`. Or use Nginx proxy to map domains to different internal ports.

**Q299: How do you enable graceful shutdown in a Node.js app?**
Listen for `SIGTERM`/`SIGINT`. `server.close()` to stop accepting connections. Close DB connections. Exit `process.exit(0)`.

**Q300: How would you build a CLI tool with Node.js?**
Add `#!/usr/bin/env node` at the top of the file. Add `"bin"` entry in `package.json`. Use `commander` or `yargs` for parsing arguments. `npm link` to test locally.

## 🔹 Node.js Core (Uncommon Concepts) - Batch 4

**Q301: What happens when you `require()` a JSON file?**
Node.js automatically parses the file using `JSON.parse()` and returns the resulting JavaScript object. It counts as a module, so it is cached in `require.cache`.

**Q302: How does Node.js internally handle DNS lookups?**
It uses the `c-ares` library (via libuv) for asynchronous DNS lookups (like `dns.resolve`). `dns.lookup` uses the system's synchronous `getaddrinfo` in a thread pool managed by libuv.

**Q303: What is the `vm` module and when would you use it?**
It allows compiling and running code within V8 virtual machine contexts. Useful for sandboxing (though not perfectly secure) or running code strings dynamically with a specific context.

**Q304: Can you dynamically load a module at runtime?**
Yes, using `require(variable)` (synchronous) or `import(variable)` (asynchronous). Useful for plugin systems or conditional loading.

**Q305: What is module hot-reloading and how do you implement it?**
Replacing modules in a running application without restart. In Node.js, this involves deleting entries from `require.cache` and re-requiring the file. Tools like `webpack` HMR do this, or `nodemon` restarts the process.

**Q306: How do you share a single instance of a module across files?**
Modules are cached by absolute filename. requiring the same file twice returns the exact same object reference (Singleton pattern by default).

**Q307: What is the `inspector` module used for?**
It opens a debugging session from within the code. You can programmatically connect to the V8 inspector to take heap snapshots or start a CPU profile on demand.

**Q308: How do you track asynchronous context in Node.js?**
Using `AsyncLocalStorage` (from `async_hooks` module). It allows storing data (like Request ID) that remains accessible across the entire async call chain without passing arguments.

**Q309: What is the significance of `globalThis` in Node.js?**
It is a standard global object available across environments (Node.js, Browser, Deno). In Node.js, `globalThis === global`.

**Q310: What is a symbol and how is it used in Node.js?**
A primitive data type (`Symbol()`). Used to create unique property keys that won't collide with other keys. Used in iterators (`Symbol.iterator`) or private-like object properties.

## 🔹 Process & Execution Environment

**Q311: What does `process.exitCode = 1` do compared to `process.exit(1)`?**
`process.exit(1)` forces the process to terminate *immediately*. `process.exitCode = 1` sets the exit code but allows the process to finish its current loop gracefully before exiting.

**Q312: What is the use of `process.stdin.setRawMode()`?**
It configures `stdin` to emit characters one by one (on keypress) without waiting for Enter. Essential for building interactive CLI tools (like menus).

**Q313: What are the arguments available in `process.argv`?**
An array of command-line arguments. Index 0 is the node executable path, Index 1 is the script file path, and Index 2+ are the user arguments.

**Q314: How can you access the current memory usage of a Node.js process?**
`process.memoryUsage()` returns an object with `rss` (Resident Set Size), `heapTotal`, `heapUsed`, and `external` memory metrics.

**Q315: What happens if a Node.js script exceeds the call stack size?**
A `RangeError: Maximum call stack size exceeded` is thrown. This usually happens with infinite recursion not using asynchronous calls.

## 🔹 Module Systems

**Q316: What are conditional exports in `package.json`?**
A feature allowing a package to define different entry points based on the environment or import type. e.g., using `import` (ESM) vs `require` (CJS).
`"exports": { "import": "./index.mjs", "require": "./index.js" }`.

**Q317: How does Node.js resolve module paths?**
It checks core modules -> `node_modules` in current dir -> parent `node_modules` -> up to root. Also respects `package.json` `"main"`/`"exports"`.

**Q318: What is `exports` field in package.json and how does it differ from `main`?**
`main` is legacy (defines one entry). `exports` is modern (Node 12+), allows defining subpaths (`pkg/feature`) and conditional exports (`import`/`require`), and encapsulates internal files.

**Q319: Can Node.js load `.mjs` and `.cjs` files in the same project?**
Yes. `.mjs` is treated as ESM, `.cjs` is treated as CommonJS. They can import/require each other with some limitations (ESM `import` of CJS works; CJS `require` of ESM fails).

**Q320: How does dynamic `import()` differ from static `require()`?**
`import()` returns a Promise and loads the module asynchronously. `require()` loads synchronously. `import()` can load ESM modules into a CJS file.

## 🔹 Streams (Advanced Edge Cases)

**Q321: How do you handle stream piping errors?**
Standard `.pipe()` does not propagate errors. Use `stream.pipeline(src, dest, (err) => ...)` to ensure errors are handled and streams are destroyed properly.

**Q322: What’s the difference between `highWaterMark` and backpressure?**
`highWaterMark` is the threshold (buffer size limit). Backpressure is the state/mechanism triggered when that limit is reached (writing returns `false`).

**Q323: How do you convert a stream to a buffer or string?**
Read all chunks into an array and concatenate.
`const chunks = []; for await (const chunk of stream) chunks.push(chunk); Buffer.concat(chunks);` (or use `stream/consumers` utilities).

**Q324: What is the `finished()` utility in stream handling?**
`stream.finished(stream, cb)` waits for the stream to complete (finish/end) or error out. Reliable way to detect stream cleanup.

**Q325: Can a stream emit events after it has ended?**
Yes, potentially (e.g., 'close' or 'error' if destroy is called later). But no 'data' events should occur after 'end'.

## 🔹 Cluster & Workers

**Q326: How do you share state between worker threads?**
Worker threads do not share memory by default. Use `SharedArrayBuffer` for binary data sharing, or pass messages via `parentPort.postMessage()`.

**Q327: What is a message channel in worker_threads?**
`MessageChannel` creates two entangled ports (`port1`, `port2`). You can pass one port to a worker so two workers can communicate directly without the parent.

**Q328: How do you handle graceful restart of a clustered Node.js app?**
Send a signal to the master process. The master sends a shutdown signal to workers one by one, waits for them to exit, and spawns new ones (rolling restart).

**Q329: When would you prefer clustering over load balancing with NGINX?**
Clustering is for utilizing multiple CPU cores on a *single* machine. NGINX load balancing is for distributing traffic across *multiple* machines/instances. Often used together.

**Q330: How do you handle sticky sessions with clustering?**
The native Node.js cluster does not support sticky sessions (round-robin). You need a master logic (like `sticky-session` package) routing connection based on IP, or rely on external multiple processes + NGINX ip_hash.

## 🔹 File System (Edge Behavior)

**Q331: How do you create a recursive directory structure in Node.js?**
`fs.mkdir('path/to/dir', { recursive: true }, cb)`.

**Q332: What’s the difference between `fs.exists()` and `fs.access()`?**
`fs.exists()` is deprecated. `fs.access()` checks accessibility (permissions). However, for operations, just try to open the file and handle the error (TOCTOU race condition avoidance).

**Q333: How do you prevent race conditions when reading and writing files?**
Use file locking (libraries like `proper-lockfile`) or ensure atomic writes (write to temp file, then rename/move to destination).

**Q334: How do you efficiently copy large files using streams?**
`fs.copyFile()` uses OS kernel copy (fastest). Or `fs.createReadStream().pipe(fs.createWriteStream())`.

**Q335: What is the use of `fs.promises`?**
It provides Promise-based versions of fs methods. `await fs.promises.readFile(...)`. Eliminates "callback hell".

## 🔹 Asynchronous Programming (Edge Cases)

**Q336: What happens if you `await` a non-promise value?**
It is wrapped in `Promise.resolve()`. `await 42` resolves to `42`. It still pauses execution until the next microtask tick.

**Q337: How does `async` function execution differ from regular functions?**
It returns a Promise implicitly. Exceptions inside are converted to Promise rejections. Code runs synchronously until the first `await`.

**Q338: How do you make a custom class thenable?**
Add a `then(resolve, reject)` method to the class. `await myObject` will call this method.

**Q339: How do you cancel a Promise in Node.js?**
Promises are not cancellable by native design. Use `AbortController` (passed to the async operation) or libraries like `bluebird` (cancellation feature).

**Q340: Can you make `setTimeout` cancellable?**
Yes, `clearTimeout(timerId)`. Or use `timers/promises` with an `AbortSignal`.

## 🔹 Memory & Performance

**Q341: How do you profile memory leaks in Node.js?**
Take Heap Snapshots over time. Look for objects that should have been GC'd (e.g., Request objects accumulating).

**Q342: What is the role of generational garbage collection in V8?**
Splits memory into New Space (Short-lived objects) and Old Space (Long-lived). Scavenge runs often on New Space (fast). Mark-Sweep-Compact runs strictly on Old Space (slower).

**Q343: How can you limit memory usage per request?**
Difficult to enforce strictly per request in JS. Architecture pattern: Spawn a child process per request (expensive) or use strict timeouts and monitoring to kill leaks.

**Q344: What are weak references and when are they useful?**
`WeakMap`/`WeakSet`. They hold references to objects without preventing Garbage Collection. Useful for caching metadata about DOM nodes or objects.

**Q345: What is the memory cost of closures in long-lived processes?**
High risk. If a closure holds a large scope (Variables), and that closure is stored (e.g., in a callback list), the entire scope stays in memory.

## 🔹 Middleware & Routing Logic

**Q346: How do you build dynamic route middleware?**
Use regex in routes `app.get(/api\/(.+)/, ...)` or parameterized middleware generators: `app.use(createAuthMiddleware(role))`.

**Q347: How do you attach data to `res.locals` in Express?**
`res.locals.user = user;`. This data is scoped to the request and available in views (if using a template engine) or subsequent middleware.

**Q348: How do you skip to the next route in Express?**
Call `next('route')`. This skips remaining middleware in the *current* route stack and passes control to the next matching route definition.

**Q349: How do you write conditional middleware execution logic?**
Wrap it. `const conditional = (req, res, next) => { if (condition) middleware(req, res, next); else next(); }`.

**Q350: What happens if a middleware hangs and never calls `next()`?**
The request hangs indefinitely until the client or server timeout kills the connection (usually 2 minutes default). Massive resource leak potential.

## 🔹 API Practices (Advanced)

**Q351: How do you handle partial success in batch APIs?**
Return 200 OK (or 207 Multi-Status). Response body contains an array of results: `[{ id: 1, status: 'success'}, { id: 2, status: 'error' }]`.

**Q352: How do you enforce strong input validation across routes?**
Use a validation library (`Zod`/`Joi`) middleware. Define schemas for headers, params, query, and body. Reject invalid requests before they reach controllers.

**Q353: What is the role of `Accept` and `Content-Type` headers in REST APIs?**
`Content-Type`: Tells server what format the client is sending.
`Accept`: Tells server what format the client wants back (JSON, XML). Server should inspect this for content negotiation.

**Q354: How do you implement HTTP content negotiation?**
`res.format({ 'text/plain': () => res.send('text'), 'application/json': () => res.json({}) })`. Express handles the `Accept` header.

**Q355: What’s the best way to paginate large datasets efficiently?**
Cursor-based pagination (`id > last_seen_id LIMIT 10`). Avoids `OFFSET` performance penalty (scanning previous rows) in SQL databases.

## 🔹 Logging & Observability

**Q356: What is structured logging?**
Logging in JSON format (`{ "level": "info", "message": "...", "userId": 123 }`) instead of plain text strings. Allows easy parsing and filtering by machines (ELK/Splunk).

**Q357: How do you correlate logs across microservices?**
Generate a UUID (`Correlation-ID` or `Trace-ID`) at the entry point. Pass it in headers to downstream services. Log it in every log entry.

**Q358: What are trace IDs and how do you generate them?**
Unique identifiers for a request chain. Generate via `crypto.randomUUID()` middleware or use OpenTelemetry auto-generation.

**Q359: How would you implement custom log levels?**
In `winston`: `const levels = { critical: 0, error: 1, ... };`. Configure logger with these levels.

**Q360: What is a sampling rate in logging/monitoring?**
Logging/Tracing only a percentage of requests (e.g., 5%) to reduce cost/storage overload while still getting statistically significant data.

## 🔹 Security (Advanced/Nuanced)

**Q361: How do you prevent timing attacks in password comparisons?**
Use `crypto.timingSafeEqual(buff1, buff2)`. It takes constant time regardless of how many characters match, preventing attackers from guessing the string char-by-char.

**Q362: What is the difference between HTTPS and HTTP/2 security in Node?**
HTTP/2 requires HTTPS (TLS 1.2+). Node.js `http2` module enforces secure ciphers. It is stricter than HTTP/1.1 defaults.

**Q363: How do you protect Node.js apps behind a reverse proxy?**
Trust the proxy: `app.set('trust proxy', 1)`. Use `X-Forwarded-For` to get real IP. Ensure the proxy (Nginx) strips these headers from external requests.

**Q364: What is secure cookie flag and how do you use it?**
`Secure` flag ensures cookies are only sent over HTTPS. `res.cookie('name', 'val', { secure: true })`.

**Q365: What is Subresource Integrity (SRI) and does it apply to Node?**
Applies to Frontend assets served by Node. Add `integrity="sha384-..."` to `<script>` tags to ensure CDN-served files haven't been tampered with.

## 🔹 Testing & Mocks (Advanced)

**Q366: How do you mock time-dependent logic in Node.js tests?**
Use Fake Timers (`jest.useFakeTimers()`, `sinon.useFakeTimers()`). You can advance time programmatically (`jest.advanceTimersByTime(1000)`).

**Q367: What’s the difference between spying and mocking?**
**Spy**: Wraps real function, tracks calls, but still executes original implementation (unless stubbed).
**Mock**: Replaces implementation completely.

**Q368: How do you test a function that depends on `fs` without real file I/O?**
Mock `fs` module (`mock-fs` library or `jest.mock('fs')`). It intercepts calls and operates on an in-memory virtual file system.

**Q369: How do you intercept and assert HTTP requests in integration tests?**
Use `nock`. It intercepts outgoing requests at the `http` module level. `nock('http://api.com').get('/user').reply(200, {})`.

**Q370: How do you test error-handling middleware?**
Trigger an error in a route: `app.get('/err', (req,res,next) => next(new Error('boom')))`. Assert the error middleware caught it and returned the correct status/response.

## 🔹 CLI Apps

**Q371: How do you parse CLI arguments in Node.js?**
Native: `process.argv.slice(2)`. Libraries: `commander`, `yargs`, `minimist`.

**Q372: What’s the difference between `yargs` and `commander`?**
Both are popular. `commander.js` is more declarative/fluent API. `yargs` is robust "pirate-themed", very good at complex sub-command parsing.

**Q373: How do you read and mask password input in the terminal?**
Use `read` or `prompts` packages with type `password`. It hides input or shows asterisks.

**Q374: How do you implement colored terminal output?**
Use ANSI escape codes or libraries like `chalk` (`chalk.red('Error')`) or `colors`.

**Q375: How do you handle interactive prompts in a Node CLI tool?**
Use `inquirer` or `prompts` libraries. `const answer = await inquirer.prompt([{ type: 'list', name: 'choice', choices: [...] }])`.

## 🔹 Dev Experience & Linting

**Q376: How do you enforce commit message conventions with Node tooling?**
Use `commitlint` with `husky`. It checks if commit messages match "Conventional Commits" standard (e.g., "feat: add user").

**Q377: What is Prettier and how does it differ from ESLint?**
**Prettier**: Formatting (Style).
**ESLint**: Logic/Quality (Bugs).
Use them together (`eslint-config-prettier` disables conflicting ESLint rules).

**Q378: How do you configure import sorting in ESLint?**
Use `eslint-plugin-import` or `eslint-plugin-simple-import-sort`. It enforces grouping/alphabetical order.

**Q379: What is the purpose of `.editorconfig` in a Node.js project?**
To maintain consistent coding style (indent style, char set) across different editors (VS Code, Vim) used by the team.

**Q380: What are shared ESLint configurations?**
Code style packages (like `eslint-config-airbnb` or `eslint-config-google`) that you can extend in your project to reuse standard rules.

## 🔹 Package & Dependency Management

**Q381: What is the difference between peerDependencies and devDependencies?**
**Dev**: Needed for dev/test (eslint, jest).
**Peer**: Needed at runtime, but expected to be installed by the *consumer* application (host), not your package (e.g., plugins).

**Q382: How does npm handle circular dependencies?**
It allows them. Module A gets a reference to Module B incomplete exports if B requires A. It works for functions (hoisted) but can fail for variables accessed immediately.

**Q383: How does `npm audit` work internally?**
It sends your dependency tree (`package-lock.json`) to the npm registry API, which compares it against a database of known security vulnerabilities.

**Q384: How do you pin exact versions for all packages?**
Use `save-exact=true` in `.npmrc` or commit `package-lock.json`. Or use `npm pkg set dependencies.x='1.2.3'`.

**Q385: What is a package-lock.json file and why is it important?**
It describes the exact tree that was generated, such that subsequent installs are able to generate identical trees, regardless of intermediate dependency updates.

## 🔹 Browser vs Node Differences

**Q386: Can you run Node.js modules in a browser?**
Not directly. You need a bundler (Webpack, Vite, Rollup) which can polyfill core modules (buffer, process) or resolve CJS/ESM differences.

**Q387: What features are exclusive to Node.js and not available in browsers?**
File System (`fs`), Operating System (`os`), direct Network sockets (`net`), `process` management. Browser has DOM, `window`, `localStorage`.

**Q388: How do you polyfill a Node.js module for frontend use?**
Bundlers usually handle this. e.g., mapping `crypto` to `crypto-browserify`.

**Q389: How is `fetch()` in browsers different from `node-fetch`?**
Browsers enforce CORS and Cookie policies. `node-fetch` (or Node native fetch) performs direct server-to-server requests with no CORS restrictions or automatic cookie jars.

**Q390: How do you implement server-side rendering (SSR) using Node?**
Render the React/Vue component to an HTML string (`renderToString`) on the server. Send HTML to client. Client "hydrates" it (attaches event listeners).

## 🔹 Hybrid & Modern Architectures

**Q391: How do you expose a GraphQL API using Node.js?**
Use `Apollo Server` or `express-graphql`. Define schema (Types) and Resolvers. `server.applyMiddleware({ app })`.

**Q392: How do you integrate gRPC and REST in a Node app?**
You can have an API Gateway (REST) that translates requests to gRPC calls to backend microservices. Or use libraries like `grpc-gateway` (Go) or manual translation in Node.

**Q393: What is a BFF (Backend-for-Frontend) and how would you implement it?**
A dedicated backend service for a specific UI (Mobile vs Web). It aggregates data and formats it specifically for that client.

**Q394: How do you stream audio/video data from Node.js?**
Use `fs.createReadStream` with `start` and `end` options based on `Range` header requests. Send status 206 (Partial Content).

**Q395: How do you serve both a web app and API from one Node server?**
API routes first: `app.use('/api', apiRouter)`.
Static files next: `app.use(express.static('build'))`.
Catch-all for SPA: `app.get('*', (req,res) => res.sendFile('index.html'))`.

## 🔹 Edge Cases & Production Scenarios

**Q396: What would you do if a Node.js app leaks file descriptors?**
Check `lsof -p <pid>`. Ensure all streams/files/sockets are closed/destroyed. Use `graceful-fs` to handle EMFILE errors (too many open files).

**Q397: What happens if you bind to a port already in use?**
`EADDRINUSE` error. The process crashes. Detect it, wait/retry, or fail.

**Q398: How do you protect an app from being overwhelmed by too many clients?**
Implement Rate Limiting, Backpressure (drop requests), and Load Shedding (rejecting requests when CPU > 90%).

**Q399: How would you safely upgrade an API in production with zero downtime?**
Rolling restarts. Load balancer directs traffic to running instances while one instance updates.

**Q400: How do you rollback a deployment if your new Node.js release is buggy?**
Revert the docker image tag in Kubernetes/ECS to the previous version. If automated, use Canary deployments that auto-rollback on high error rates.

## 🔹 Node.js Runtime & Internals - Batch 5

**Q401: What are native addons in Node.js and how are they written?**
Native addons are dynamically linked shared objects (written in C/C++) that can be loaded into Node.js. They provide an interface between JavaScript and C/C++ libraries.

**Q402: What is the role of `node-addon-api`?**
A module that provides a C++ wrapper for N-API (Node-API). It simplifies writing native addons by handling the underlying API complexities and providing a C++ object model.

**Q403: What are bootstrap loaders in Node.js?**
Internal scripts that run during the Node.js startup process to initialize the environment, load built-in modules, and prepare the execution context before user code runs.

**Q404: How does Node.js handle uncaught exceptions in ES modules?**
If a top-level execution throws, the process exits with a non-zero code. If a Promise in top-level await rejects, it creates an `unhandledRejection`.

**Q405: What is the purpose of the `--experimental` flag in Node.js?**
It enables features that are currently in development and not yet stable (e.g., `--experimental-fetch` in older versions). Use with caution in production.

**Q406: What’s the difference between synchronous and asynchronous garbage collection?**
GC is generally synchronous (pauses execution). However, V8 performs some GC steps (like Scavenging or marking) in background threads to minimize the "Stop-The-World" pause on the main thread.

**Q407: How does Node.js integrate with libuv?**
Node.js uses `libuv` as its platform abstraction layer. `libuv` provides the event loop, thread pool, and asynchronous I/O primitives (FS, Network) that Node.js bindings expose to JS.

**Q408: What is the difference between `queueMicrotask` and `process.nextTick()`?**
`process.nextTick` runs *before* other microtasks (Promise jobs). `queueMicrotask` queues a task in the *same* microtask queue as Promises. `nextTick` is Node-specific; `queueMicrotask` is standard Web API.

**Q409: What does the `--trace-sync-io` flag do?**
It prints a warning/stack trace whenever the application performs synchronous I/O operations (like `fs.readFileSync`) during the event loop execution (after startup). Good for performance auditing.

**Q410: How does Node.js maintain a reference to the event loop?**
Internal C++ handles (Ref/Unref). If active handles (timers, servers) exist, the loop continues. If all handles are `unref()`'d or closed, the loop exits.

## 🔹 Memory Management

**Q411: How do you force garbage collection in a development environment?**
Run node with `--expose-gc` flag. Then call `global.gc()` in your code. Useful for testing memory leaks.

**Q412: What are retained objects in heap snapshots?**
Objects that are not directly reachable from the root (Global) but are kept alive because another reachable object holds a reference to them.

**Q413: What is memory fragmentation and how does it impact long-lived apps?**
When memory is allocated and freed in random order, leaving small gaps that cannot fit large objects. This forces the OS to allocate more RAM even if "free" space exists.

**Q414: How can memory leaks occur via closures in async code?**
If an async operation (like a callback waiting on a DB) hangs or takes long, the closure context (and all variables in scope) stays in memory, preventing GC.

**Q415: What tools does Node.js offer for memory debugging?**
`node --inspect` (Chrome DevTools), `heapdump`, `process.memoryUsage()`, and `v8.getHeapStatistics()`.

## 🔹 Advanced Event Loop Analysis

**Q416: How would you trace the path of an asynchronous call from start to finish?**
Use `async_hooks` to capture `init`, `before`, `after`, and `destroy` events for the resource. Or use APM distributed tracing.

**Q417: What is async context propagation?**
Passing contextual data (User ID, Transaction ID) through the callback/promise chain automatically. Node.js uses `AsyncLocalStorage` to achieve this efficiently.

**Q418: What is a task queue vs. a job queue in Node.js?**
**Task Queue (Macrotask)**: Callbacks for `setTimeout`, I/O, `setImmediate`.
**Job Queue (Microtask)**: Promises, `queueMicrotask`, `process.nextTick`. Microtasks run immediately after the current operation.

**Q419: What types of timers are handled in the “timers” phase?**
`setTimeout` and `setInterval`. (Note: `setImmediate` is handled in the "check" phase, not timers).

**Q420: When does a Node.js process _not_ exit automatically?**
When there are active handles (Open socket, Interval running) or pending asynchronous operations in the `libuv` loop.

## 🔹 Environment & CLI Flags

**Q421: What is `NODE_OPTIONS` used for?**
An environment variable to pass command-line flags to the node executable (e.g., `NODE_OPTIONS="--max-old-space-size=4096"`). Useful in environments where you can't change the startup command.

**Q422: What does `--inspect-brk` do?**
It starts the inspector and pauses execution at the very first line of the script, waiting for a debugger to connect.

**Q423: What is the purpose of `--no-deprecation`?**
It silences deprecation warnings printed to stderr. Useful in production logs to reduce noise (though fixing the issue is better).

**Q424: How do you detect memory limits set via CLI?**
`v8.getHeapStatistics().heap_size_limit`. It reflects the default or the value set by `--max-old-space-size`.

**Q425: What’s the impact of `--max-old-space-size`?**
It sets the hard limit for the V8 Old Space (Heap). If the app exceeds this, V8 crashes with `FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed - JavaScript heap out of memory`.

## 🔹 File System (Rare Use Cases)

**Q426: How do you watch a directory tree recursively?**
`fs.watch('dir', { recursive: true }, cb)`. Note: Recursive watching is not supported on all platforms (like Linux) natively by `fs.watch` in older versions or relies on polling.

**Q427: How do symbolic links behave in `fs.readdir()`?**
`fs.readdir` just lists the name of the link. It does not traverse it. You need `fs.lstat` to identify it's a symbol link.

**Q428: What’s the difference between `fs.readSync()` and `fs.readFileSync()`?**
`fs.readFileSync` reads the whole file content at once. `fs.readSync` reads a specific chunk into a buffer (requires opening the file descriptor with `fs.openSync` first).

**Q429: How do you safely handle file descriptor limits?**
Implement a queue or pool for file operations (`graceful-fs` essentially does this). If `EMFILE` error occurs, retry after a delay.

**Q430: How do you prevent concurrent file writes from corrupting data?**
Use a lock file (like `proper-lockfile`). Ensure only one process/function can write to the file at a time.

## 🔹 Networking & Sockets

**Q431: How do you create a TCP server in Node.js?**
`const net = require('net'); const server = net.createServer(socket => { ... }); server.listen(8080);`

**Q432: What’s the difference between `net.Socket` and `tls.TLSSocket`?**
`tls.TLSSocket` wraps a `net.Socket` and handles encryption/decryption (SSL/TLS) transparently for the application.

**Q433: How would you handle slowloris attacks at the Node.js level?**
Use a reverse proxy (Nginx) with timeout settings. In Node, set `server.headersTimeout` and `server.requestTimeout` to drop connections sending headers too slowly.

**Q434: How does the `dns` module work differently in `resolve` vs. `lookup`?**
`dns.lookup` (default for `http`) uses `getaddrinfo` (OS hosts file + DNS). `dns.resolve` does a direct query to DNS servers over the network (ignores local hosts file).

**Q435: What is socket multiplexing and does Node.js support it?**
Sending multiple signals over one channel. Node.js supports it via HTTP/2 (streams over one TCP connection).

## 🔹 Experimental APIs & Web Platform Integration

**Q436: What is the Fetch API in Node.js and how is it different from browsers?**
Native `fetch` (v18+). It uses `undici` under the hood. It follows strict spec compliance but some browser features (like Service Workers) aren't present.

**Q437: What are Web Streams and how are they implemented in Node.js?**
Standard streams (`ReadableStream`, `WritableStream`). Node.js implements them (GLOBAL) to be compatible with Browsers/Deno, distinct from legacy Node.js streams (`stream` module).

**Q438: What is the File API and does Node.js support it natively?**
The `File` and `Blob` globals are part of the Web API spec. Node.js v20+ supports them, allowing compatibility with frontend code expecting `File` objects.

**Q439: How is `AbortController` integrated into Node.js APIs?**
Most modern async APIs (`fs.readFile`, `setTimeout`, `http.request`) accept an `signal` option. `abortController.abort()` triggers an `AbortError`.

**Q440: What is the `WebAssembly` module in Node.js?**
Allows checking, compiling, and instantiating Wasm modules. Node.js can load `.wasm` files (via fs) and execute high-performance binary code.

## 🔹 V8 Engine Nuances

**Q441: What are hidden classes and how do they affect performance?**
(See Q249 - Repeated concept in lists, emphasizing importance). V8 internally maps objects with same structure to same hidden class (`Map`). Dynamic property addition breaks this chain, slowing access.

**Q442: What is the cost of polymorphism in hot V8 functions?**
If a function accepts multiple object shapes (types), V8 stays "megamorphic" (cannot optimize). Monomorphic functions (same input shape) allow V8 to inline and optimize aggressively.

**Q443: How does V8 optimize frequently called functions?**
It uses "TurboFan" (optimizing compiler). It profiles the function execution (Ignition interpreter), assumes types are stable, and compiles to optimized machine code.

**Q444: What’s the difference between inline and megamorphic calls?**
**Inline**: Compiler replaces function call with function body (fast).
**Megamorphic**: Compiler cannot predict target (too many types), so performs slow lookup.

**Q445: What are deoptimization triggers in V8?**
Type changes (passing string instead of int), try/catch (in older V8), modifying arguments object, or exceeding optimization bail-out limits.

## 🔹 Debugging & Diagnostics

**Q446: How do you use `node --inspect` with Chrome DevTools?**
Start with `--inspect`. Open Chrome -> `chrome://inspect`. Click "Open dedicated DevTools for Node". You get full Breakpoints, Console, and Profiler.

**Q447: What is the role of `v8.getHeapStatistics()`?**
Returns runtime metrics about V8 memory: heap limit, used size, available size. Critical for monitoring memory pressure.

**Q448: How do you analyze a core dump from a Node.js crash?**
Use `llnode` (an LLDB plugin) to inspect the stack frames and objects in the core dump file generated by the OS on crash.

**Q449: What does `node --trace-deprecation` show you?**
Prints the stack trace whenever a deprecated API is used, helping you locate *where* in your code (or dependencies) the old method is called.

**Q450: How do you view async stack traces in modern Node?**
Enabled by default. If an error occurs in an async chain, Node prints the "Zero-cost async stack trace", showing the Promise creation point.

## 🔹 Concurrency Models & Coordination

**Q451: What is the `Atomics` module and how does it work with shared memory?**
Provides atomic operations (`add`, `sub`, `wait`, `notify`) on `SharedArrayBuffer`. Allows safe data manipulation across Worker threads without race conditions.

**Q452: What is the purpose of `SharedArrayBuffer` in Node.js?**
A memory buffer that can be read/written by multiple threads (Workers) simultaneously. zero-copy data sharing.

**Q453: How do you implement a mutex in Node.js?**
Using `Atomics.wait()` and `Atomics.notify()` on an Int32Array view of a SharedArrayBuffer. One thread waits until another notifies.

**Q454: How can you coordinate multiple async workers to share a task queue?**
Ideally, use a central broker (Redis). For purely internal workers, checking an atomic counter or using the main thread as a dispatcher (Round Robin).

**Q455: When should you use worker threads over a message broker?**
**Worker Threads**: CPU-bound tasks, shared memory needed, single-machine scale.
**Broker**: Distributed systems, decoupling services, reliability/persistence needed.

## 🔹 Job Queues & Background Work

**Q456: What’s the difference between `bull`, `kue`, and `agenda`?**
**Bull**: Redis-based, fast, robust.
**Kue**: Redis-based, older (maintenance mode).
**Agenda**: MongoDB-based, good if you already use Mongo.

**Q457: How do you persist job states across restarts?**
The queue library (Bull/Agenda) stores the job data in the database (Redis/Mongo). On restart, it queries the DB for "pending" or "stalled" jobs and processes them.

**Q458: How do you ensure at-least-once job delivery?**
Process jobs in a way that confirms completion (ACK) only after success. If worker crashes, the job remains in the queue (or moves to "failed" -> retry).

**Q459: How would you build a delayed job scheduler without external dependencies?**
For ephemeral: `setTimeout` (bad for long delays).
For persistent: Store `executionTime` in DB. Run a `setInterval` checking `DB.find({ executionTime: { $lte: now } })`.

**Q460: How do you monitor background job failure rates?**
Most queue libraries emit events (`job.on('failed')`). Listen to these, aggregate metrics (Prometheus), or hook into a UI (Bull Board).

## 🔹 Binary Data & Encoding

**Q461: How do you read a binary file byte-by-byte?**
Open file (`fs.open`). Read into a Buffer of size 1. Loop. OR iterate a ReadStream.

**Q462: How do you convert a binary buffer to a hexadecimal string?**
`buffer.toString('hex')`.

**Q463: What’s the difference between `utf8` and `utf16le` in Node?**
Encoding formats. `utf8` (1-4 bytes) is standard for web. `utf16le` (2 or 4 bytes) is often used in Windows systems. Node supports both.

**Q464: How does base64 encoding inflate data size?**
Base64 represents binary data using 64 ASCII characters. It increases size by approx 33% (4 chars per 3 bytes).

**Q465: How do you detect and handle corrupted binary input?**
Use checksums (CRC32, MD5, SHA) to verify integrity. If decoding/parsing fails (try/catch), treat as corrupt.

## 🔹 Module Publishing & Tooling

**Q466: What is a dual-module package?**
A package that supports both CommonJS (`require`) and ESM (`import`) consumers. Usually achieved via `exports` config pointing to `.cjs` and `.mjs` builds.

**Q467: How do you publish both CommonJS and ESM versions of a module?**
Compile TS/JS to two output folders (`dist/cjs`, `dist/esm`). Point `package.json` fields appropriately.

**Q468: What is the `exports` map in `package.json` used for?**
(Repeated concept - see Q318). Defines public entry points. Helps hide internal files (`require('pkg/internal')` fails if not exported).

**Q469: How do you bundle a Node library using Rollup or esbuild?**
Configure build tool to target `node`. Keep dependencies external (don't bundle `node_modules`). Output multiple formats (CJS/ESM).

**Q470: What’s the difference between a named export and a default export in ESM?**
**Named**: `export const foo`. specific import `{ foo }`. Better for tree-shaking.
**Default**: `export default`. import any name `import X`. Simpler for single-responsibility modules.

## 🔹 Cross-Platform Concerns

**Q471: How do you handle Windows file paths vs POSIX paths?**
Use `path.join()` and `path.resolve()` always. Avoid hardcoding `/` or `\` separators.

**Q472: How do you detect the current platform in Node.js?**
`process.platform` (returns `'win32'`, `'linux'`, `'darwin'`).

**Q473: How do you set environment variables cross-platform?**
Use `cross-env` package. `cross-env NODE_ENV=production node app.js`. Windows cmd doesn't support `VAR=val` syntax natively.

**Q474: How does signal handling differ on Windows vs Linux in Node?**
Windows does not support signals (SIGINT, SIGTERM) fully. Node.js emulates some, but `SIGUSR1` etc. might not work. Docker on Windows has specific signal propagation caveats.

**Q475: What are EOL issues and how do you handle them?**
Line endings (CRLF vs LF). Use `.gitattributes` to enforce LF in repo. Use `os.EOL` constant in code when writing files (if OS-specific behavior is desired).

## 🔹 Standard Library (Lesser-Known Modules)

**Q476: What is the `perf_hooks` module used for?**
(See Q194 - Repeated concept). Used for High Resolution Timing and performance measurement APIs (PerformanceObserver).

**Q477: What is the `readline/promises` API?**
A Promise-based version of the `readline` module (v17+). `const answer = await rl.question('Prompt?');`.

**Q478: How do you use the `url` module for parsing and formatting?**
`const myUrl = new URL('https://example.com');`. Native global object conforming to WHATWG URL standard. Replaces legacy `url.parse`.

**Q479: How does the `assert` module handle deep comparisons?**
`assert.deepStrictEqual(obj1, obj2)`. Recursively checks own enumerable properties. Throws AssertionError if different.

**Q480: What is the purpose of the `tty` module?**
Provides information about the terminal driver. `process.stdout.isTTY` returns true if output is a terminal (not piped to file).

## 🔹 Testing Strategies (Rare Cases)

**Q481: How do you test code that uses `process.exit()`?**
Stub `process.exit` in the test setup. `jest.spyOn(process, 'exit').mockImplementation(() => {})`. Verify it was called.

**Q482: How do you mock system time or timezone?**
System time: Fake Timers. Timezone: Set `process.env.TZ` before starting the test runner (requires restart usually) or use libraries that mock the Date constructor fully.

**Q483: How do you assert log output during tests?**
Spy on `console.log` or your logger instance. `jest.spyOn(console, 'log')`. Assert calls.

**Q484: How do you create dynamic test cases programmatically?**
Loop over data array. `testCases.forEach(input => test('should handle ' + input, () => { ... }))`.

**Q485: What’s the benefit of property-based testing in Node?**
Libraries like `fast-check` generate random inputs to find edge cases you didn't think of.

## 🔹 Ecosystem & Community Awareness

**Q486: What are some alternatives to npm (other than yarn)?**
`pnpm` (performant, disk-space efficient symlinking), `cnpm` (China mirror), `bun install` (fast).

**Q487: What is a monorepo and how do tools like Lerna or Nx help manage it?**
(Repeated concept Q165). Multiple packages in one repo. Tools handle dependency linking (`symlink` local packages) and running commands across all packages.

**Q488: What’s the role of `npx` in rapid tooling?**
Execute binaries from npm packages without installing globally. Great for one-off commands (`create-react-app`, `sequelize-cli`).

**Q489: How do you consume private npm packages securely?**
Authenticate using `.npmrc` with an Auth Token. `//registry.npmjs.org/:_authToken=${NPM_TOKEN}`.

**Q490: What are the differences between CommonJS loaders and ES module loaders?**
CJS loaders (Webpack/Node internal) can hook into `require()`. ESM loaders use `--loader` or `register()` hooks to customize resolution/loading of specifiers.

## 🔹 Modern Architectures & Interop

**Q491: How do you run Node.js in a Deno environment (if at all)?**
Deno supports running some Node.js code via `std/node` compatibility layer or `npm:` specifiers. "Denoify" tools exist.

**Q492: What are the challenges in porting a Node.js project to ESM-only?**
Updating all `require` to `import`, handling `__dirname` (missing in ESM), updating config files, and checking dependency support.

**Q493: What’s the role of edge computing in Node.js backends?**
Running logic closer to user (Vercel Edge, Cloudflare Workers). Often uses a subset of Node APIs (V8 isolates) for lower latency than centralized servers.

**Q494: How do you run Node.js in a Cloudflare Worker or similar runtime?**
You often don't run full Node.js. You run code *compatible* with the runtime's API (often based on Web Standards like Fetch/Request). Some Node polyfills exist (`node:` prefix support increasing).

**Q495: How does Node.js compare with Bun in cold start times?**
Bun (built on Zig/JavaScriptCore) claims significantly faster startup execution than Node.js (V8). Important for CLI tools/Serverless.

## 🔹 Tooling Best Practices

**Q496: How do you precompile TypeScript for Node.js deployment?**
Run `tsc` (TypeScript Compiler) to output `.js` files to `dist/` or `build/` folder. Deploy those JS files.

**Q497: What are build-time vs. runtime dependencies?**
**Build-time (devDOeps)**: Needed to compile/test (TypeScript, Jest).
**Runtime (dependencies)**: Needed to run (Express, Lodash).

**Q498: How do you automate changelogs with Node.js tools?**
Use `standard-version` or `semantic-release`. They parse commit messages (Conventional Commits) and generate CHANGELOG.md automatically.

**Q499: What is a zero-dependency package and why does it matter?**
A package with no dependencies in `package.json`. Easier to audit, smaller install size, less risk of supply chain attacks.

**Q500: What are best practices for publishing scoped packages?**
Use `@org/pkgname`. Ensure `access` is set correctly (`public` or `restricted`). Useful for internal org code or official plugin ecosystems.

## 🔹 Modern Build Tools & Dev Workflows - Batch 6

**Q501: How do you use `esbuild` in a Node.js app?**
Install `esbuild`. Run `esbuild src/app.ts --bundle --platform=node --outfile=dist/out.js`. Extreme speed for transpilation/bundling compared to TSC/Webpack.

**Q502: What’s the difference between transpiling and bundling?**
**Transpiling**: Converting source code to another version (TS -> JS, ES6 -> ES5). File-to-file.
**Bundling**: Combining multiple files/modules into a single (or few) output file(s), often resolving dependencies.

**Q503: How does Vite handle Node.js APIs?**
Vite is primarily for frontend, but for SSR (Server-Side Rendering), it builds Node-compatible code. It uses `rollup` for production builds, which can be configured to target `node` environment.

**Q504: How do you pre-bundle dependencies in a Node.js monorepo?**
Use tools like `preconstruct` or `microbundle`. Or configure `esbuild` in the root to bundle common shared packages into `dist` folders that other packages consume.

**Q505: How do you create a custom plugin for Rollup in a Node.js project?**
Export a function returning an object with hooks (`resolveId`, `load`, `transform`). `export default function myPlugin() { return { name: 'my-plugin', transform(code) { ... } } }`.

**Q506: What’s the purpose of `tsup` in Node.js development?**
A zero-config bundler powered by `esbuild`. It simplifies building TypeScript libraries for Node.js, automatically handling `.d.ts` generation and dual CJS/ESM output.

**Q507: How do you reduce cold start time for serverless Node functions?**
Minimize bundle size (tree-shaking, exclude `aws-sdk`), use `esbuild` for faster parsing, reuse DB connections outside handler, and avoid heavy sync initialization.

**Q508: How can `SWC` be used with Node.js for fast transpilation?**
Use `@swc/core` or `swc-node` (via `ts-node` or `jest`). It is a Rust-based compiler, significantly faster than `babel` or `tsc`.

**Q509: What’s the benefit of lazy importing modules?**
Reduces startup time/memory. `if (condition) { const module = require('heavy-lib'); ... }`. Only loads the module when strictly needed.

**Q510: How do you structure a hybrid TypeScript and JavaScript Node app?**
Allow `allowJs: true` in `tsconfig.json`. Migrate incrementally. Keep JS files in `src` alongside TS. TS compiler processes both.

## 🔹 Edge & Serverless Environments

**Q511: How does Node.js differ when deployed at the edge?**
Edge runtimes (Cloudflare Workers, Vercel Edge) often use V8 isolates, not full Node.js. They may lack native modules (`fs`, `child_process`) but implement Web Standards (`fetch`, `Request`).

**Q512: What are the Node.js limitations in Cloudflare Workers?**
No access to File System, no native C++ addons, specialized socket API (Connect), execution time limits (CPU time), specific memory constraints.

**Q513: How do you handle cold starts in AWS Lambda with Node.js?**
Use Provisioned Concurrency (keeps instances warm), optimize code size, and use the latest Node.js runtime version available (often optimized).

**Q514: How do you manage shared dependencies in a serverless bundle?**
Use Lambda Layers (AWS) to host common `node_modules`. Or use a bundler (Webpack/Esbuild) to inline only used code, ignoring unused parts of huge libraries.

**Q515: How do you reduce package size for edge deployment?**
Bundle code, minify, remove source maps. replace heavy libraries (like `moment`, `lodash`) with lighter alternatives (`dayjs`, `lodash-es`).

**Q516: How do you handle sessions in stateless serverless functions?**
Do not store state in memory/file. Use external stores: Redis (e.g., Upstash for serverless), DynamoDB, or encrypted Client-side Cookies/JWTs.

**Q517: What’s the difference between regional and edge Node.js deployments?**
**Regional**: Code runs in a specific data center (e.g., us-east-1). High latency for distant users.
**Edge**: Code is replicated to hundreds of locations globally. Low latency, typically strictly for stateless logic.

**Q518: How do you maintain execution context across stateless requests?**
You can't rely on global variables. Pass context (User ID, Trace ID) explicitly to functions or use `AsyncLocalStorage` restricted to the request scope/lifetime.

**Q519: What’s the impact of V8 isolates in edge runtimes?**
Faster startup (milliseconds) compared to containers/VMs. Lower overhead allows high concurrency but entails stricter sandbox limits (no system calls).

**Q520: What are best practices for debugging serverless Node.js apps?**
Use structured logging (JSON) sent to CloudWatch/Datadog. distinct Request IDs. Replicate environment locally (using `serverless-offline` or `SAM local`).

## 🔹 Machine Learning & AI Integrations

**Q521: How can Node.js consume a RESTful AI model?**
Use `axios`/`fetch` to POST data to the model server (TensorFlow Serving, OpenAI API). handle response asynchronously.

**Q522: What’s the purpose of `onnxruntime-node`?**
A library to run ONNX (Open Neural Network Exchange) models in Node.js directly. Allows running pre-trained PyTorch/TensorFlow models efficiently on CPU/GPU.

**Q523: How do you run inference from a TensorFlow.js model in Node?**
Install `@tensorflow/tfjs-node`. Load model: `tf.loadLayersModel('file://...')`. Run `model.predict(tensor)`. The `-node` binding uses C++ acceleration.

**Q524: How do you handle streaming OpenAI responses in Node?**
Set `stream: true` in API call. The response body is a stream. Iterate chunks (`for await (const chunk of stream)`) and push to client via Server-Sent Events (SSE).

**Q525: How do you fine-tune a chatbot using Node.js tooling?**
Node.js acts as the orchestrator: Preprocess dataset (JSONL), upload file to OpenAI/HuggingFace API, trigger fine-tuning job, and monitor status.

**Q526: How do you implement rate limiting when accessing AI APIs?**
Use a token bucket implementation or queues (`bull`). Track token usage (prompt + completion tokens) to stay within quotas/budgets.

**Q527: What’s the role of `worker_threads` in running AI tasks in Node.js?**
AI inference (matrix math) is CPU-heavy and blocks the event loop. Offload it to a Worker Thread so the main thread remains responsive for HTTP requests.

**Q528: How do you integrate a HuggingFace model with Node.js?**
Use the HuggingFace Inference API (`@huggingface/inference`) or download the model and run via `onnxruntime` or `transformers.js` (emerging JS-native support).

**Q529: What are the memory implications of running large models in Node?**
Models load into RAM. A 7B parameter model takes GBs of RAM. Node's V8 heap limit might be exceeded. Use native bindings (outside V8 heap) or stick to quantization/API calls.

**Q530: How do you serialize/deserialize AI model input/output in Node?**
Inputs often need to be Tensors (TypedArrays: Float32Array). Serialize as binary buffers or JSON lists. Output probabilities need mapping back to labels.

## 🔹 Advanced Plugin Systems

**Q531: How do you design a pluggable architecture in Node.js?**
Define a `Plugin` interface. Core app iterates a `plugins` directory, calling `require()`, then executes `plugin.register(appContext)`. Use hooks for extensibility.

**Q532: How would you build a plugin loader using dynamic imports?**
`const plugins = ['p1', 'p2']; for(const p of plugins) { const mod = await import(p); mod.default(app); }`.

**Q533: What is sandboxing, and how can it apply to plugins?**
Running plugin code in an isolated environment (`vm` module). Prevents plugins from accessing globals, `process`, or crashing the main app (security/stability).

**Q534: How do you isolate plugin execution to prevent memory leaks?**
Run plugins in separate Worker Threads or Child Processes. If a plugin leaks or crashes, you can terminate the container without affecting the core.

**Q535: How do you dynamically load and unload modules safely?**
Loading is easy (`require`/`import`). "Unloading" is hard (delete from `require.cache`). JS has no "unload" for code. Best way is restarting the process or workers.

## 🔹 Observability at Scale

**Q536: What are red metrics and how do you apply them to Node.js?**
**R**ate (Requests/sec), **E**rrors (Failed reqs), **D**uration (Response time). Measure these for every service endpoint using middleware (Prometheus).

**Q537: How do you track async performance bottlenecks?**
Use `perf_hooks` to wrap async functions, or an APM that patches Promises. Look for "long waits" where CPU is idle but latency is high (I/O bottleneck).

**Q538: What are the challenges in tracing requests through worker threads?**
Context (Trace ID) is lost when passing message to Worker. You must manually serialize context in the `postMessage` payload and restore it in the Worker `AsyncLocalStorage`.

**Q539: How would you log structured user actions across multiple services?**
Define a standard JSON schema event (Who, What, When). Send these events to a centralized bus (Kafka), not just log files. Service A produces "UserCreated", B consumes.

**Q540: What is distributed tracing and how do you implement it with OpenTelemetry?**
Tracing a request lifecycle across microservices. OTel SDKs inject `traceparent` headers into outgoing HTTP/gRPC requests and extract them on incoming ones.

## 🔹 Real-World Failure Scenarios

**Q541: How do you prevent resource exhaustion in a long-lived Node.js process?**
Implement strict limits: memory limits (`max-old-space-size`), connection timeouts, max payload sizes, and restart periodically/on-threshold (using PM2 or K8s liveness probes).

**Q542: What happens if you perform thousands of concurrent file reads?**
You hit `EMFILE` (Too many open files) error. The OS limits file descriptors per process. Solution: Use `graceful-fs` or restrict concurrency (pool).

**Q543: How do you protect your Node.js app against infinite request loops?**
Detect "loop" headers (via CDN/Proxy). Or implement `max-redirects`. In microservices, check call depth or duplicate request IDs.

**Q544: How would you simulate an out-of-memory crash for testing?**
Create an array and push large buffers in a loop until crash: `const arr = []; while(1) arr.push(Buffer.alloc(1e6));`

**Q545: How do you build a recovery strategy after an API rate limit block?**
Catch 429 errors. Read `Retry-After` header. Pause outgoing requests (circuit breaker/queue pause) for that duration before resuming.

## 🔹 Dynamic Runtime Control

**Q546: How do you apply feature flags at runtime in Node.js?**
Use a service (LaunchDarkly) or poll a config DB/Redis. Check flag value `if (flags.isEnabled('new-feat'))` inside the route handler.

**Q547: What is a dynamic configuration system and how can Node.js support it?**
A system where config changes without restart. Watch a shared file (`fs.watch`) or subscribe to a config server (Consul/etcd) updates and update global config object.

**Q548: How do you handle live configuration reloads without restarting the app?**
Update the configuration variable in memory. Note: Be careful with DB connection pools—you might need to recreate them gracefully if connection strings change.

**Q549: How do you implement a “kill switch” for a Node.js feature in production?**
A specific high-priority Feature Flag. If enabled (switch killed), the code path throws 503 Service Unavailable immediately.

**Q550: What’s the role of circuit breakers in fault-tolerant Node.js design?**
To detect failing dependencies (API/DB). If failures exceeds threshold, "Open" the circuit (fail fast) to prevent cascading failure and allow the dependency to recover.

## 🔹 Browser Integration via Node.js

**Q551: How do you serve pre-rendered HTML in a Node.js SSR app?**
On request, fetch data, render React/Vue component to string, inject into index.html template, and send `res.send(html)`.

**Q552: How do you hydrate browser-side components from a Node server?**
The server sends initial HTML state (often serialized in `window.__INITIAL_STATE__`). Client JS reads this state to initialize components matching the server markup.

**Q553: How do you use `puppeteer` or `playwright` for browser automation in Node.js?**
Libraries controlling headless Chrome. `const browser = await puppeteer.launch(); const page = await browser.newPage(); await page.goto(url);`. Used for scraping or PDF generation.

**Q554: How do you create a browser extension using a Node.js build system?**
Extensions are HTML/JS. Use Node (Webpack/Vite) to build the background scripts, content scripts, and popups into a `dist` folder loadable by the browser.

**Q555: How do you bridge WebRTC functionality between Node and browser clients?**
Use `wrtc` or `werift` libraries in Node to implement a WebRTC Peer. This allows Node to join a video call or data channel as a participant/recorder.

## 🔹 Security (Modern Practices)

**Q556: How do you use CSP headers with Node.js?**
`Content-Security-Policy`. Use `helmet` middleware. `helmet.contentSecurityPolicy({ directives: { "script-src": ["'self'", ...] } })`. Prevents XSS/Injection.

**Q557: How do you detect dependency confusion attacks in npm?**
Ensure your internal packages are scoped (`@myorg/pkg`). Configure npm to point the scope strictly to your private registry, never public.

**Q558: What is token binding and how could it be used in Node.js?**
Cryptographically binding a session token to the TLS connection/client certificate. Prevents stolen session cookies from being used on another machine.

**Q559: How do you implement certificate pinning in a Node HTTPS client?**
Pass `ca` or `cert` options to `https.agent`. Validate the server's certificate fingerprint matches expected hash.

**Q560: What is a secure enclave and how can Node.js interact with one?**
Hardware isolation (Intel SGX, AWS Nitro). Node apps can't easily run inside, but can offload sensitive computation (signing) to enclave services (like KMS) via API.

## 🔹 Native APIs & Binary Extensions

**Q561: What’s the difference between `ffi-napi` and `node-gyp`?**
**node-gyp**: Compiles C++ source to a native addon (.node) during install.
**ffi-napi**: Loads dynamic libraries (.dll/.so) at runtime and calls functions without writing C++ glue code.

**Q562: How do you wrap a C++ library for use in Node.js?**
Write a Node-API (N-API) binding. expose C++ functions as JS-callable functions. Parse arguments in C++, call library, return result.

**Q563: What are napi threadsafe functions?**
Special N-API functions that allow a C++ thread (background) to safely call a JavaScript function (on the main thread) without crashing V8.

**Q564: How does Node.js handle cross-compiling native modules?**
Use `node-gyp` with target architecture flags (`--arch=arm64`). Often done in CI/Docker environments matching the target OS/Arch.

**Q565: How do you debug a segmentation fault in a native Node.js module?**
Use `gdb` or `lldb`. `lldb -- node app.js`. Run, crash, view backtrace (`bt`) to see which C++ line caused the memory violation.

## 🔹 Compatibility & Portability

**Q566: What are the challenges in making a Node app portable across OSes?**
File paths (`\` vs `/`), signals (SIGINT/SIGTERM behavior), shell commands in `child_process` (cmd vs bash), and native dependencies.

**Q567: How do you handle Node.js binary compatibility across ARM and x64?**
You must install/compile native modules separately for each arch. Use multi-arch Docker images(`linux/amd64`, `linux/arm64`) to ship correct binaries.

**Q568: What is the role of Docker multi-arch builds for Node.js?**
Allows one image tag to run on MacBook M1 (ARM) and AWS EC2 (x64) transparently. `docker buildx build --platform linux/amd64,linux/arm64`.

**Q569: How do you ensure consistent `node_modules` across platforms?**
You generally can't (native modules differ). But `package-lock.json` ensures consistent *versions*. `npm ci` ensures clean install.

**Q570: What are Windows-specific pitfalls when developing with Node.js?**
Max path length (260 chars) issues in deep `node_modules` (fixed in newer npm/Windows). CRLF line endings breaking scripts.

## 🔹 API Architecture Design

**Q571: How do you implement conditional responses using ETag headers?**
Calculate hash of response body (ETag). Verify `If-None-Match` header from client. If match, return 304 Not Modified (empty body). Saves bandwidth.

**Q572: What is a schema-first approach to API development in Node?**
Define OpenAPI/GraphQL schema *before* writing code. Generate Types/Validators/Docs from the schema automatically. Ensures consistency.

**Q573: How do you dynamically generate OpenAPI docs from your Node.js app?**
Use libraries like `swagger-jsdoc` that parse JSDoc comments `@swagger` above routes and output JSON.

**Q574: How do you structure an API gateway in a Node monorepo?**
Gateway package acts as the router/proxy. It imports standard types/contracts from shared packages. It routes requests to backend services (internal IPs).

**Q575: How do you handle feature deprecation safely in APIs?**
Mark endpoints/fields `@deprecated`. Return `Warning` header in response. Monitor usage logs. Announce sunset date. Remove in major version.

## 🔹 Performance Profiling & Heatmaps

**Q576: How do you create CPU flamegraphs for Node.js?**
Generate: `node --prof app.js`. Process: `node --prof-process isolate-0x...log > processed.txt`. Visualize: Use `flamebearer` or `0x` tool.

**Q577: What is an event loop utilization metric?**
`performance.eventLoopUtilization()`. Ratio of time loop spent active vs idle. High utilization means the CPU is saturated with JS tasks (potential lag).

**Q578: How do you identify high-GC zones in your Node app?**
Run with `--trace-gc`. Analyze logs. Frequent "Scavenge" is normal. Frequent "Mark-sweep" indicates memory pressure or leaks.

**Q579: What is the role of `v8-profiler-node8` or `clinic.js`?**
Tools to diagnose performance.
**Clinic Doctor**: Diagnoses I/O, CPU, Event Loop issues.
**Clinic Flame**: Visualizes CPU.
**Clinic Bubbleprof**: Visualizes Async flow.

**Q580: How do you build a performance heatmap for your routes?**
Middleware tracks `start` time and `end` time per route. Aggregate avergage latency into a matrix (Route x TimeOfDay). Visualize where hotspots are.

## 🔹 Command Line UX Enhancements

**Q581: How do you create an animated CLI using Node.js?**
Use libraries like `ora` (spinners), `listr` (task lists), or `ink` (React for CLI). Rerender the terminal lines using ANSI escape codes.

**Q582: What is an interactive TUI (text UI) in Node.js?**
Application with layout (boxes, text input, navigation) in terminal. Libraries: `blessed`, `blessed-contrib`, `react-ink`.

**Q583: How do you support tab-completion for custom Node CLIs?**
Using `yargs` or `omelette`. Generate a completion script (bash/zsh) that the user sources in their shell profile.

**Q584: How do you manage multi-language CLI output in Node.js?**
i18n. Load JSON translation files based on OS locale (`Intl.DateTimeFormat().resolvedOptions().locale`).

**Q585: What is the purpose of chalk, ora, and inquirer combined?**
The "Holy Trinity" of Node CLIs. **Chalk** (Color), **Ora** (Spinner/Status), **Inquirer** (Prompts/Input). 

## 🔹 Emerging Trends (2024–2025)

**Q586: What is the impact of WebContainers on Node.js development?**
Running Node.js completely inside the Browser (via WebAssembly). Enables full IDEs (StackBlitz) and dev environments in the browser without backend servers.

**Q587: How do you build native apps using Node.js + Tauri?**
Node.js (sidecar) usually handles backend logic if needed, but Tauri primarily uses Rust. However, you can bundle a Node binary to establish a local server for the frontend UI.

**Q588: How do you use EdgeDB or SurrealDB with Node?**
Modern graph-relational databases. Use their official Node.js drivers. They allow advanced query composition often simpler than SQL.

**Q589: How does Bun’s compatibility layer affect Node.js libraries?**
Bun implements `node:*` APIs tailored to match Node behavior. Most libraries work, but those relying on V8 specifics or undocumented internals fail.

**Q590: How do you benchmark Node.js vs Rust in network-bound scenarios?**
Use `wrk` or `bombardier`. Ensure apples-to-apples (same machine, optimized builds). Rust usually wins on CPU/Memory, Node often competitive on pure I/O throughput.

## 🔹 Testing at Scale

**Q591: How do you run test sharding across multiple Node.js runners?**
Split test files into N groups (modulo math or time-based). Run each group on a separate CI machine. `jest --shard=1/4`.

**Q592: How do you snapshot test CLI apps?**
Capture stdout/stderr. Compare output string against saved snapshot file. Strip dynamic content (timestamps) before comparing.

**Q593: How do you write fuzz tests for JSON input handling?**
Use `fast-check`. Generate random JSON structures (deep, malformed, huge strings). Feed to API parser. Ensure it throws 400, not 500 crash.

**Q594: What are smoke tests in the context of Node deployment?**
Fast, non-destructive tests run against the Production environment immediately after deploy. (e.g., Check Health endpoint, Login with test user).

**Q595: How do you test a Node.js app behind a feature flag system?**
Mock the feature flag provider to return True/False permutations. Run test suites for both states to ensure no regression in either path.

## 🔹 Cultural & Community Ecosystem

**Q596: What is the Node.js Technical Steering Committee (TSC)?**
The group responsible for high-level technical direction and governance of the Node.js project.

**Q597: What are the benefits of LTS in Node.js versions?**
Long Term Support (Active/Maintenance for 30 months). Stability guarantees, security patches. Recommended for Enterprise Production use.

**Q598: How do Node.js release schedules impact library maintainers?**
Maintainers must support all Active LTS versions. dropping support for old versions (e.g., Node 14) usually requires a SemVer Major bump in the library.

**Q599: How do you become a Node.js module maintainer?**
Contribute heavily to an open source module. Fix bugs, improve docs, respond to issues. Ask adoption.

**Q600: What are the biggest current challenges in the Node.js ecosystem?**
ESM vs CJS fragmentation (dual package hazard), Supply Chain Security (malicious packages), and competition from newer runtimes (Bun/Deno).

## 🔹 Modern Runtimes & Compatibility - Batch 7

**Q601: How do you polyfill Node.js core modules in non-Node runtimes (like Deno)?**
Use compatibility layers (Deno's `std/node` or Bun's built-in node compatibility). Or use bundlers (Webpack/Rollup) with `node-stdlib-browser` or `rollup-plugin-polyfill-node`.

**Q602: What is the `node:` prefix in module imports, and why was it introduced?**
E.g., `import fs from 'node:fs'`. Indicates explicitly that the module is a Node.js built-in, preventing conflict with npm packages named "fs" or "path". Supported in recent Node versions.

**Q603: How does Node.js interact with WebAssembly-based modules?**
Load `.wasm` binary via `fs`, compile with `WebAssembly.compile()`, instantiate with `WebAssembly.instantiate()`. Pass imports (memory, JS functions) to the Wasm instance.

**Q604: How do you check runtime compatibility for packages targeting Node and Bun?**
Check if the package uses Node-specific APIs (V8 internals, `child_process`) that Bun might implement differently. Run the test suite in both runtimes.

**Q605: What are the limitations of using ESM-only packages in older Node versions?**
Node <12 doesn't support ESM. Node 12-13 had experimental support. You cannot `require()` an ESM-only package synchronously in CJS; you must use `await import()`.

**Q606: How do you handle dynamic import paths in hybrid ESM/CJS apps?**
Use `await import(path)`. Note that `__dirname` is missing in ESM, so you need `import.meta.url` and `fileURLToPath` to construct paths dynamically.

**Q607: What tools help ensure cross-runtime compatibility (Node, Deno, Bun)?**
`unjs` ecosystem (unbuild, unenv), `denoify`, or writing standard Web API code (Request/Response/Fetch) that works everywhere (WinterCG compliance).

**Q608: What is a runtime shim and when do you need it?**
A piece of code that intercepts API calls to adapt them to the underlying environment (e.g., implementing `window` or `fetch` in Node versions that lack it).

**Q609: How do you simulate browser APIs (like `localStorage`) in Node for SSR?**
Use libraries like `node-localstorage` or attach a mock object to `global`. `global.localStorage = { getItem: () => ... }`.

**Q610: What happens when you try to use `window` in Node.js?**
ReferenceError: window is not defined. Node has `global`. To share code, check `typeof window !== 'undefined'` or use `globalThis`.

## 🔹 Interfacing with Other Languages

**Q611: How do you call a Rust function from Node.js?**
Compile Rust to a native Node addon using `napi-rs` or `neon`. Import the resulting `.node` file in JS. Or compile to Wasm and load it.

**Q612: How can you interface Node.js with a Python script?**
Spawn a child process: `spawn('python', ['script.py'])`. Communicate via stdin/stdout (JSON strings). Or use libraries like `python-shell` or `pymport` (in-process).

**Q613: What is `grpc-node` and how does it connect to services written in Go or Java?**
It's a pure JS or Native implementation of gRPC for Node. Uses Protocol Buffers to define schema. Connects via HTTP/2. Language-agnostic contract facilitates communication.

**Q614: How do you expose a Node.js library to be used from Java?**
Expose it as a microservice (REST/gRPC). Or use Truffle/GraalVM (Polyglot runtime) to run Node.js code inside the Java VM context.

**Q615: What are foreign function interfaces (FFIs) and how are they used in Node.js?**
FFI allows calling functions in dynamic libraries (.dll/.so) written in C/C++ without writing native addon code. Libraries: `ffi-napi`.

## 🔹 Protocol Support

**Q616: How do you implement a WebSocket server in pure Node.js?**
Listen to the `upgrade` event on HTTP server. Accept the handshake (Sec-WebSocket-Accept). Parse frames directly (opcode, masking). Easier to use `ws` lib.

**Q617: How would you build an MQTT client in Node?**
Use `mqtt.js` library. Connect to broker (`mqtt://broker`). `client.subscribe('topic')`. `client.on('message', cb)`.

**Q618: How do you use Node.js to communicate over raw TCP or UDP?**
**TCP**: `net.createConnection()`.
**UDP**: `dgram.createSocket('udp4')`. Use for custom protocols or low-latency games.

**Q619: How do you build a gRPC reflection server in Node?**
Enable reflection in `@grpc/reflection`. Allows clients (like Postman/grpcurl) to query the server for available services and proto definitions at runtime.

**Q620: How do you serve and consume GraphQL over WebSockets using Node.js?**
Use `graphql-ws` or `subscriptions-transport-ws` (legacy). Define Subscription type in schema. Resolvers return `AsyncIterator`.

## 🔹 Multi-Tenant, SaaS, and Platform Architectures

**Q621: How do you isolate tenant data in a Node.js SaaS app?**
Logical isolation (Row-level `tenant_id` check in EVERY query). Or Physical isolation (Separate DB or Schema per tenant).

**Q622: What is a tenant-aware database connection, and how do you implement it?**
A wrapper around the DB pool. Middleware determines tenant. Wrapper selects the correct connection string/pool for that tenant context.

**Q623: How do you sandbox plugins per tenant in a shared Node.js runtime?**
Use `vm` context or `isolated-vm`. Set strict timeouts and memory limits. Ideally, move unsafe tenant code to separate processes/workers.

**Q624: How do you enforce per-tenant rate limits?**
Rate limiter uses a key composed of `tenantId`. `redis.incr('rate:tenant:123')`. Cap total requests across all users of that tenant.

**Q625: How would you structure a monolith-to-microservices migration in Node?**
Identify domains. Extract one module (e.g., "Billing") to a separate Node app. Route traffic via Gateway. Use "Strangler Fig" pattern.

## 🔹 Security (Beyond Basics)

**Q626: How do you validate JWTs with rotating public keys?**
Use `jwks-rsa` middleware. It automatically fetches the latest JSON Web Key Set (JWKS) from the Auth Provider (Auth0/Cognito) and caches it to verify signatures.

**Q627: What are token replay attacks, and how do you prevent them?**
Attacker captures a valid token and reuses it. Prevent via: Short expiry times, Nonce (jti claim) tracking (stateful), or limiting use to one IP (issues with mobile).

**Q628: How do you securely rotate API keys in a Node.js system?**
Generate new key. Add to valid list. Allow both keys (overlap period). Update clients. Remove old key. Store hashed keys in DB.

**Q629: How do you implement HMAC signature verification?**
Receive payload + signature header. Compute `crypto.createHmac('sha256', secret).update(payload).digest('hex')`. Compare with received signature using `timingSafeEqual`.

**Q630: How do you prevent internal SSRF attacks in internal-only Node APIs?**
Validate target URLs. Deny private IP ranges (10.x, 192.168.x, localhost). Use a DNS resolver that filters private IPs.

## 🔹 Zero-Downtime Deployments

**Q631: How do you handle WebSocket connections during a Node.js deploy?**
WebSockets are persistent. Deployment kills them. Client must have reconnection logic (exponential backoff). On server, stop accepting new conn, wait for activity to drop, then close.

**Q632: How do you implement graceful shutdown with open file handles?**
Track open handles. On SIGTERM, stop watcher/streams. `fs.close(fd)`. Wait for 'close' events. Force exit after timeout.

**Q633: What is connection draining and how is it used in load-balanced Node apps?**
LB stops sending new requests to the instance. Instance finishes existing requests. Once 0 active requests, instance terminates.

**Q634: How do you persist in-flight jobs across rolling deployments?**
Don't keep jobs in memory. Use a persistent queue (Redis). If process dies, the un-ACK'd job eventually times out in queue and is picked up by another worker.

**Q635: What’s the role of process signaling in safe Node.js shutdown?**
`SIGTERM` is the warning "Please cleanup". App closes HTTP server (stops listening). Finishes current reqs. Then exits. Platform sends `SIGKILL` if app takes too long.

## 🔹 GraphQL in Production

**Q636: How do you batch GraphQL queries on the server?**
Use `DataLoader`. It coalesces multiple requests for the same resource type (e.g., user IDs) within one tick into a single DB query (`WHERE id IN (...)`).

**Q637: What is persisted GraphQL and how does it work in Node.js?**
Client sends a query hash (SHA) instead of the full query string. Server looks up the query map. Reduces bandwidth and prevents arbitrary query attacks.

**Q638: How do you enforce query complexity limits?**
Calculate a score based on nesting depth and field cost. Reject queries with score > 1000. Use `graphql-query-complexity`.

**Q639: How do you write GraphQL subscriptions using `apollo-server`?**
Define Subscription resolver with `subscribe: () => pubsub.asyncIterator('TOPIC')`. Use WebSocket protocol.

**Q640: How would you log GraphQL resolver performance?**
Use Apollo Tracing or a plugin that measures `start` and `end` time of each resolver function. Push metrics to APM.

## 🔹 Globalization & Accessibility

**Q641: How do you implement multilingual support in a Node.js CLI app?**
Detect locale `os-locale`. Load JSON resource structure `locales/en.json`, `locales/es.json`. Replace strings with keys.

**Q642: What is i18next and how is it used in Node apps?**
Popular i18n framework. `i18next.init({ resources }).then(() => console.log(i18next.t('welcome')));`. Supports interpolation and plurals.

**Q643: How do you localize logs and error messages per user context?**
Do *not* translate internal server logs (keep English for devs). Translate only the user-facing error message sent in the API response.

**Q644: How do you detect and support right-to-left (RTL) content in SSR?**
Check locale (e.g., `ar`, `he`). Add `dir="rtl"` attribute to the `<html>` or `<body>` tag in the generated HTML template.

**Q645: What is the role of Unicode normalization in Node.js string handling?**
`str.normalize('NFC')`. Ensures consistent byte representation of characters (e.g., 'ñ' can be one char or two). crucial for string comparison/hashing.

## 🔹 Advanced Logging Infrastructure

**Q646: What is a structured log, and why is it important for observability?**
(See Q356). JSON logs. Allows querying `log.response_time > 500` in log management systems.

**Q647: How do you implement log sampling in high-throughput Node apps?**
Randomly select 1% of requests to log fully (debug level). For others, log only Errors/Warnings. Reduces noise and storage cost.

**Q648: What is the difference between log aggregation and log forwarding?**
**Forwarding**: Shipping logs from the node (Filebeat/Fluentd) to a destination.
**Aggregation**: collecting logs from all nodes into one searchable UI (ELK).

**Q649: How do you correlate logs across multiple Node.js services?**
Pass a unique `Trace-ID` header. Include this ID in every log statement associated with that request chain.

**Q650: How do you implement request tracing across async boundaries?**
Use `AsyncLocalStorage`. Store the Trace ID on entry. The storage "follows" the async execution. Retrieve it in the logger function automatically.

## 🔹 Release & Version Management

**Q651: How do you test a Node app against multiple Node versions?**
Use CI matrix. `strategy: matrix: node-version: [16.x, 18.x, 20.x]`. Run tests in parallel containers for each version.

**Q652: What is semantic-release and how does it automate versioning?**
It analyzes commits (fix, feat, breaking) to determine the next version number (patch, minor, major). It tags git, creates release, and publishes to npm.

**Q653: How do you enforce changelog updates via CI in a Node.js repo?**
Use tools like `danger-js` or check if `CHANGELOG.md` was modified in the PR diff. Or auto-generate it (better).

**Q654: What is a canary release and how do you implement it in Node?**
Deploy new version to a small % of users. Route traffic via LB. internal Node app doesn't know it's a canary, the infrastructure handles routing.

**Q655: How do you manage breaking changes in a public npm package?**
Bump Major version. Provide migration guide. Optionally provide a "codemod" script to help users upgrade automatically.

## 🔹 Node.js as a Platform (Beyond Web)

**Q656: How do you build a desktop app using Electron and Node.js?**
Electron bundles Chromium + Node.js. Main process (Node) manages windows/system. Renderer process (Chromium) shows UI. They communicate via IPC.

**Q657: How do you create a plugin system for a design tool using Node?**
(Similar to Q296). Load JS files from a user folder. Expose a secure API to manipulate the design model.

**Q658: What are the security implications of bundling Node in desktop apps?**
XSS in the renderer can become Remote Code Execution (RCE) if Node integration is enabled in the web view. Disable `nodeIntegration` in renderer; use `contextBridge`.

**Q659: How do you expose Node.js functionality in a VSCode extension?**
VSCode extensions run in a Node.js host. You have access to `fs`, `path`, etc. You can spawn child processes or run language servers.

**Q660: How do you build a test runner or build tool using Node?**
Use APIs like `glob` to find files. `fs` to read. `vm` or `require` to execute tests. `process` to report exit code. CLI styling with `chalk`.

## 🔹 Data Processing & Transformation

**Q661: How do you process NDJSON (newline-delimited JSON) in streams?**
Use `split2` module to split stream by newline. Pipe to a transformer that parses `JSON.parse`. Pipe to output.

**Q662: How do you perform ETL operations in real time using Node.js?**
Read Stream (Source) -> Transform Stream (filtering/mapping) -> Write Stream (Destination/DB). Keeps memory usage low even for GBs of data.

**Q663: What are transform streams and how do they enable live data manipulation?**
`new Transform({ transform(chunk, encoding, cb) { ... }})`. It sits in the middle of a pipe chain, modifying chunks as they pass through.

**Q664: How do you handle schema evolution in streamed JSON data?**
Include version field in data. Transform stream checks version. If old, map to new structure (migration on read).

**Q665: How do you deduplicate and batch data in-flight?**
Buffer chunks in the Transform stream. Maintain a Set of IDs. Flush the buffer only when size limit or time limit reached.

## 🔹 Project Management & Collaboration

**Q666: How do you enforce code ownership in a Node.js monorepo?**
Use `CODEOWNERS` file. Map paths (`packages/billing/*`) to GitHub teams (`@org/billing-team`).

**Q667: How do you prevent dependency drift in long-lived projects?**
Use RenovateBot or Dependabot. Automatically open PRs for updates. Merge if tests pass.

**Q668: How do you create internal reusable dev tools using Node.js?**
Create a CLI package (`@work/cli`). Publish to private registry. Devs install globally or run via `npx @work/cli`.

**Q669: What’s the difference between shared configs vs. shared code in monorepos?**
**Configs**: ESLint/TSConfig (static setup).
**Code**: Utility libraries (logging, formatting). Both reduce duplication.

**Q670: How do you handle multi-team deployment coordination in Node-based stacks?**
Micro-frontends or Microservices. Independent deployment pipelines. Contract testing (Pact) prevents breaking each other.

## 🔹 Real-Time Collaboration Features

**Q671: How do you implement CRDTs in a Node.js app?**
Use libraries like `Yjs` or `Automerge`. The Node server acts as a relay/storage peer. It merges updates from clients to ensure eventual consistency.

**Q672: What is Operational Transformation and how can it be supported in Node?**
Algorithm for syncing text edits (like Google Docs). Server acts as single source of truth, transforming incoming operations against current history. `shareb` is a library.

**Q673: How do you sync cursor position across clients using Node?**
Broadcast ephemeral messages via WebSocket. No need to store in DB. "User X is at (10, 20)". Clients render overlays.

**Q674: How do you broadcast changes only to affected clients in real time?**
Use "Rooms" or "Channels" in Socket.IO. Join clients to `room:doc:123`. Broadcast only to that room.

**Q675: How do you implement offline editing with sync in a Node backend?**
Client stores edits locally (indexedDB). When online, sends sync queue. Server merges methods (CRDT/OT). Returns new state.

## 🔹 Advanced CLI Development

**Q676: How do you parse subcommands and nested arguments?**
`yargs.command('user', 'manage users', (yargs) => { yargs.command('create', ... ) })`. Git-style subcommands.

**Q677: How do you write an interactive install wizard in a CLI tool?**
Use `inquirer`. Ask series of questions (Path? Options?). Generate config file based on answers.

**Q678: How do you auto-generate CLI docs from source code?**
Extract command descriptions from `yargs`/`commander` configuration. Output Markdown.

**Q679: What are hidden CLI flags, and when should you use them?**
Flags not listed in `--help`. Useful for debugging, experimental features, or super-user overrides.

**Q680: How do you ensure consistency across multiple CLI binaries in a monorepo?**
Share a core "CLI Framework" package. Enforce consistent branding, logging style, and error handling across tools.

## 🔹 Time-Sensitive & Scheduled Execution

**Q681: How do you implement time zone–aware job scheduling?**
Use `cron` syntax with TZ support (e.g., `cron` package). `new CronJob(time, task, null, true, 'America/New_York')`.

**Q682: What’s the difference between cron and interval-based scheduling?**
**Cron**: Wall-clock time (Run at 8am). Resilient to drifts.
**Interval**: Relative time (Run every 10m). Can drift. Stops if process stops.

**Q683: How do you queue a job for delayed execution in distributed Node apps?**
Use Redis-backed queue (`bull`). `queue.add(data, { delay: 5000 })`. Not in-memory `setTimeout`.

**Q684: How do you recover from missed schedules due to downtime?**
Persist last run time. On startup, check `if (now - lastRun > interval)`. Run immediately or skip based on policy.

**Q685: How do you simulate cron execution during test runs?**
Manually trigger the job function. Or Mock the system clock/cron library to fire immediately.

## 🔹 Unusual Bug Hunting & Debugging

**Q686: How do you detect file handle leaks in production?**
Monitor open file descriptor count (`lsof`). If growing steadily, you have a leak.

**Q687: How do you identify memory bloat from large object graphs?**
Heap Snapshot. Look for "Distance" from GC root. Deeply nested objects or large strings retained by caches.

**Q688: How do you find async functions that never resolve?**
`wtfnode` tool. It lists active handles and promises preventing exit. Or manual logging timeouts.

**Q689: What is heapdump and how is it analyzed?**
A module to write a snapshot to disk programmatically. Load `.heapsnapshot` into Chrome DevTools to analyze variables.

**Q690: How do you inspect deadlocks caused by incorrect `await` usage?**
Use `async_hooks` to track promise lifecycle. If promise is created but never resolved, debug the resolver logic. "Await cycle" (A awaits B, B awaits A) can deadlock.

## 🔹 Documentation & Developer Experience

**Q691: How do you generate interactive API documentation from JSDoc?**
Use tools that convert JSDoc comments to HTML/Markdown.

**Q692: What’s the difference between a readme and usage guides?**
**Readme**: First impression, install, quick start.
**Guide**: In-depth tutorials, concepts, advanced config.

**Q693: How do you embed live code examples in Node.js documentation?**
Use platforms like RunKit or CodeSandbox. Embed iframes in docs allowing users to run the snippet.

**Q694: What is the role of OpenAPI in dev onboarding?**
Provides a "Try it out" explorer (Swagger UI). New devs can call API without writing code to understand behavior.

**Q695: How do you document internal-only APIs in a public Node module?**
Mark with `@internal` JSDoc tag. Configure doc generator to exclude them. Or use symbol for private methods.

## 🔹 Packaging Strategies & Distribution

**Q696: How do you create a single executable binary from a Node app?**
Use `pkg` (Vercel) or `sea` (Single Executable Applications - Node 20 experimental). Bundles node runtime + script + assets.

**Q697: What’s the difference between `pkg`, `nexe`, and `vercel/pkg`?**
They are similar tools. `pkg` (by Vercel) is most popular. `nexe` is older. Node.js SEA is the future native solution.

**Q698: How do you bundle a Node app for distribution without exposing source code?**
Compile to Binary (`pkg`) or use code obfuscators (`javascript-obfuscator`). Be aware that bytecode can technically be reversed/decompiled, but it raises the bar.

**Q699: How do you embed assets (HTML, CSS, etc.) into a Node.js binary?**
`pkg` virtual filesystem. `path.join(__dirname, 'asset')` works inside the binary.

**Q700: What are the limitations of freezing a Node.js project into a binary?**
The binary is large (includes Node capabilities). Native modules (.node) need special handling (often extracted at runtime).

## 🔹 Hybrid & Polyglot Environments - Batch 8

**Q701: How do you use Node.js within a polyglot microservices architecture?**
Node.js typically handles I/O-heavy services (API Gateway, Real-time) while CPU-heavy services are in Go/Rust. They communicate via gRPC, HTTP/REST, or message queues (Kafka).

**Q702: How do you communicate between a Node.js app and a service written in Rust?**
**Network**: gRPC or REST.
**FFI**: Call Rust from Node using `napi-rs`.
**Wasm**: Compile Rust to Wasm, load in Node.

**Q703: What is the benefit of using WebAssembly modules from Node.js?**
Near-native performance for CPU tasks (image processing, crypto) without the complexity/instability of C++ addons. Sandboxed execution.

**Q704: How do you bridge native code and JavaScript via `node-addon-api`?**
Write C++ code using `Napi::Env` and `Napi::Object`. Compile with `node-gyp`. Require the `.node` binary. It exposes C++ logic as JS objects.

**Q705: How do you containerize multi-language services with Node.js entrypoints?**
Use a Docker multi-stage build. Install Node runtimes. Copy compiled Go/Rust binaries into the Node image. Exec via `child_process`.

## 🔹 Service Mesh & Node.js

**Q706: How do you integrate a Node.js app with Istio?**
Deploy app to K8s. Istio injects an Envoy sidecar container in the Pod. Node app takes to `localhost` (Envoy), which handles mTLS, routing, and tracing.

**Q707: How do sidecars affect Node.js app performance?**
They add a small latency hop (sub-millisecond) for network calls. However, they offload SSL termination and retry logic from the Node event loop, potentially *improving* total throughput.

**Q708: What headers must a Node.js app forward in a mesh for tracing?**
`x-request-id`, `x-b3-traceid`, `x-b3-spanid`, `x-b3-parentspanid`, `x-b3-sampled`, `x-ot-span-context` (Zipkin/B3 headers).

**Q709: How do you implement mutual TLS (mTLS) with Node.js in a mesh?**
Generally, you don't. You let the Envoy sidecar handle mTLS. The traffic between Sidecar and Node app is plain HTTP (or local mTLS if strict).

**Q710: How does a Node.js app behave with circuit breaking in Envoy?**
If Envoy detects the Node app failing (5xx), it cuts traffic. Node app receives fewer requests. When it recovers, Envoy slowly increases traffic.

## 🔹 Long-Running & Daemon Processes

**Q711: How do you prevent memory bloat in a long-running Node process?**
Regularly force GC (if safe) or restart periodically (PM2 `max-memory-restart`). Avoid unbounded caches. release large buffers.

**Q712: How do you write a system daemon using Node.js?**
Use a wrapper like `pm2` or create a systemd service file (`App.service`) that runs `ExecStart=/usr/bin/node /path/to/app.js`.

**Q713: How do you manage orphaned child processes from a Node app?**
By default, if parent dies, child might remain. Use `child.unref()` if intentional. Or listen to `exit` signal in parent and `child.kill()`.

**Q714: How do you daemonize a Node.js script on Linux?**
`nohup node app.js &`. Or use `daemon` npm package. Best practice: Systemd or Docker.

**Q715: How do you watch a directory and respond to file system events robustly?**
Use `chokidar` (wrapper around `fs.watch`). It normalizes events across OSes and handles debouncing of rapid write events.

## 🔹 Memory & GC Tuning

**Q716: What flags help tune V8 garbage collection for large heaps?**
`--max-old-space-size=N`, `--min-semi-space-size`, `--max-semi-space-size`. Tuning "New Space" size can reduce frequency of scavenge.

**Q717: How do you profile heap growth in production Node apps?**
Use `heapdump` triggered by signal (`SIGUSR2`). Or `v8.writeHeapSnapshot()`. Upload snapshot to S3. Analyze offline.

**Q718: What’s the difference between minor and major GC in V8?**
**Minor (Scavenge)**: Fast, moves live objects from New Space -> Old Space.
**Major (Mark-Sweep)**: Slow, pauses execution, cleans Old Space.

**Q719: How do you find and resolve detached DOM-like memory leaks?**
(In SSR contexts like JSDOM). Look for objects detached from the document tree but referenced by JS variables.

**Q720: How does object shape affect memory usage in Node?**
Consistent shapes (same keys in same order) share Hidden Classes. V8 stores them efficiently. Inconsistent shapes create separate classes and "dictionary mode" objects (more memory/slower).

## 🔹 Advanced HTTP Handling

**Q721: How do you handle `Expect: 100-continue` requests?**
Listen to `checkContinue` event on `http.Server`. If validated, call `res.writeContinue()`. Else close connection. (Default Node behavior handles this automatically).

**Q722: What is the impact of `Transfer-Encoding: chunked` on Node streams?**
Allows sending data before total size is known. Node `http` uses this by default for streams. Client receives data in pieces.

**Q723: How do you throttle uploads in a Node server?**
Pipe the request stream through a `Throttle` transform stream (pauses/resumes based on byte count/time).

**Q724: How do you respond to `OPTIONS` requests for CORS preflights?**
Return 204 No Content. Set Access-Control headers. `app.options('*', cors())`.

**Q725: How do you spoof IPs during development for geo-based testing?**
Manually set `X-Forwarded-For` header in your request (Postman). Ensure your app reads this header (trust proxy enabled).

## 🔹 CI/CD & Build Optimization

**Q726: How do you cache `node_modules` effectively in CI pipelines?**
Cache key should be hash of `package-lock.json`. Restore cache. Run `npm ci`. If lockfile changed, cache miss -> fresh install.

**Q727: What’s the difference between `npm ci` and `npm install` in CI?**
**`npm ci`**: Clean Install. Deletes `node_modules`. Installs EXACT versions from lockfile. Fails if lockfile and package.json mismatch. Faster/Deterministic.

**Q728: How do you run partial test suites using tags or filters?**
`jest -t "API Tests"`. Markers in code `describe('API Tests', ...)`

**Q729: How do you automate semantic versioning based on commit messages?**
`semantic-release`. It parses "feat:", "fix:", "BREAKING CHANGE" strings to calculate version jump.

**Q730: How do you run tests in parallel across CI runners?**
Split by filename. `circleci tests split`. Pass subset of files to Jest.

## 🔹 Authentication Patterns

**Q731: How do you implement OAuth 2.1 PKCE flow in a Node.js backend?**
PKCE is typically for public clients (Mobile/SPA). But Node can act as client. Generate `code_verifier`. Hash to `code_challenge`. Send in auth req. Exchange code + verifier for token.

**Q732: How do you validate signed cookies in stateless JWT auth?**
Cookie contains `JWT.signature`. `parser.verify(cookie_value, secret)`. Signature prevents tampering.

**Q733: What is token chaining and how do you implement it in Node?**
Service A calls Service B. A passes the User's token (Authorization header) to B. "On-Behalf-Of" flow.

**Q734: How do you handle refresh token theft in a Node app?**
Token Rotation. When refresh token is used, issue a NEW refresh token. Invalidate old one. If old one used again, invalidate ALL tokens for that family (theft detected).

**Q735: How do you federate identity from multiple providers?**
Use Passport.js strategies (`passport-google`, `passport-github`). Normalize their profiles into a standard User object in your DB.

## 🔹 Monorepo Best Practices

**Q736: What’s the difference between hoisting and non-hoisting strategies?**
**Hoisting**: Moving common deps to root `node_modules`. Saves space. Can cause "Phantom Dependencies".
**No-hoist**: Each pkg gets own `node_modules`. Safer, more disk usage.

**Q737: How do you share environment config securely across multiple packages?**
Do not commit .env. Use a shared config package `packages/config` that reads `process.env`. Load env vars at the Application Entry point (root).

**Q738: How do you manage package versioning inside a monorepo?**
Fixed mode (All packages version together) vs Independent mode (Separate versions). Lerna/Changesets handle this.

**Q739: How do you publish internal-only npm packages?**
Set `"private": true` in package.json. Or publish to private registry (Artifactory/Verdaccio).

**Q740: How do you avoid dependency hell in a monorepo with many teams?**
Use strict boundaries via `nx` or `turborepo`. Prohibit circular deps. Automate upgrades globally.

## 🔹 Modern Asset Pipelines with Node

**Q741: How do you build an image CDN using Node.js?**
Accept image URL. Fetch. Resize using `sharp` or `jimp`. Cache result (Redis/S3). Stream to response.

**Q742: How do you transcode audio/video files in a Node worker?**
Spawn `ffmpeg` process. `ffmpeg -i input.mp4 output.webm`. Monitor progress via stderr parsing.

**Q743: What are the performance implications of image manipulation via `sharp`?**
`sharp` uses `libvips` (C++). It is extremely fast and memory efficient (streaming) compared to pure JS solutions.

**Q744: How do you bundle WASM modules into Node packages?**
Include `.wasm` files in the package. Use `fs.readFile` (with `__dirname`) to load them at runtime for instantiation.

**Q745: How do you offload asset compression to a background process?**
Upload triggers job in Queue. Worker process downloads, compresses, uploads processed version. Updates DB status.

## 🔹 Deep Testing Scenarios

**Q746: How do you mock `fs` operations without breaking `realpath`?**
`mock-fs` can break `require`. Use `jest.mock('fs')` carefully, or better, structure code to accept a "FileSystem" interface (Dependency Injection) and pass a memory implementation during tests.

**Q747: How do you simulate backpressure in integration tests?**
Create a slow Writable stream. `writable.write` returns false. Verify the Readable stream pauses (stops emitting data).

**Q748: How do you test middleware order and side effects?**
Register middleware. Spy on their execution. `expect(mid1Stats).toHaveBeenCalledBefore(mid2Stats)`. Check `req` mutations.

**Q749: How do you test streaming responses in HTTP servers?**
Make request. Listen to `data` event of response. usage: `supertest`. `request(app).get('/stream').buffer(false).parse((res, cb) => { res.on('data', ...) })`.

**Q750: How do you verify telemetry output in tests?**
Mock the metrics client (StatsD/Prometheus). `expect(marketingClient.increment).toHaveBeenCalledWith('login.success')`.

## 🔹 Security Edge Cases

**Q751: What’s a prototype pollution vulnerability, and how does Node help mitigate it?**
Attacker merges `{ "__proto__": { "polluted": true } }` into an object. Affects all objects. Mitigation: Use `Object.create(null)` or `Map`. Validate JSON keys. Freeze prototypes.

**Q752: How do you defend against path traversal when serving files?**
Normalize path. Check if it starts with allowed root. `const safePath = path.normalize(req.path); if (!safePath.startsWith(baseDir)) throw Error;`

**Q753: How do you securely evaluate dynamic code (e.g., user-generated math)?**
Do NOT use `eval` or `new Function`. Use a parser (e.g., `mathjs`) that evaluates the AST safely.

**Q754: What tools detect dependency-level vulnerabilities in Node.js?**
`npm audit`, `snyk`, `dependabot`.

**Q755: What are security headers, and how do you implement them in Node?**
HSTS, X-Frame-Options, X-Content-Type-Options. Use `helmet` middleware.

## 🔹 Edge Computing & CDN Integration

**Q756: How do you cache SSR responses at the CDN layer from Node?**
Set `Cache-Control: public, max-age=3600, s-maxage=86400`. `s-maxage` tells CDN to cache for 1 day, while browser caches `max-age`.

**Q757: How do you support stale-while-revalidate in Node?**
Set `Cache-Control: s-maxage=..., stale-while-revalidate=...`. CDN serves stale content while fetching fresh content from Node in background.

**Q758: How do you differentiate bot vs. user traffic in edge Node apps?**
Check `User-Agent`. Use a list of known bots. Or check `Sec-CH-UA` (Client Hints).

**Q759: How do you run a Node function across multiple PoPs?**
Use edge platforms (Lambda@Edge, Cloudflare Workers). You deploy once; platform replicates code.

**Q760: How do you implement rate limiting at the CDN edge with Node fallback?**
CDN (Cloudflare) WAF handles gross rate limiting. Node handles fine-grained logical limiting (per user/action).

## 🔹 Platform Design & SDKs

**Q761: How do you build a Node SDK for a third-party API?**
Create a class wrapping `axios`. Methods map to endpoints. Handle auth/retries internally. usage `const client = new MySdk(apiKey)`.

**Q762: How do you version APIs exposed via an SDK?**
SDK version matches API version (mostly). or `client.v1.resource.get()`.

**Q763: How do you support plugin hooks in a Node SDK?**
Allow users to register middleware. `client.use((req, next) => { log(req); next(); })`.

**Q764: How do you enforce request schema validation in SDK usage?**
Use Typescript definitions. Runtime validation with Zod (optional) to give helpful errors "Expected string, got number".

**Q765: How do you handle breaking API changes in SDKs?**
Deprecate method in SDK log. Map new method to old logic if possible. Major ver bump.

## 🔹 Resilience Patterns

**Q766: How do you retry failed API calls with exponential backoff?**
`axios-retry`. Or loop with delay `Math.pow(2, attempt) * 100`.

**Q767: How do you implement a circuit breaker in Node?**
`opossum` library. Wrap async function. Options: `timeout`, `errorThresholdPercentage`, `resetTimeout`.

**Q768: What is a bulkhead pattern, and how does it apply to Node?**
Isolating resources. Separate connection pools for critical vs non-critical services. If non-critical pool saturates, critical still works.

**Q769: How do you isolate failed dependencies in a microservice?**
Wrap dependency calls in try/catch/CircuitBreaker. Return fallback data (default value/cache) on failure instead of crashing.

**Q770: How do you track downstream dependency health?**
Implement Deep Health Checks. `/health` endpoint calls DB, Redis, and dependent Services. Reports status of all.

## 🔹 Package Publishing

**Q771: How do you mark a package as private vs. public?**
`"private": true` (Prevents publish). OR `"publishConfig": { "access": "public" }` for scoped packages intended to be public.

**Q772: How do you publish beta tags like `@next` or `@alpha`?**
`npm publish --tag next`. Users install via `npm install pkg@next`.

**Q773: How do you deprecate a published npm version?**
`npm deprecate pkg@1.0.0 "Use v1.1.0"`. Warnings appear on install.

**Q774: What are the security risks of installing unscoped packages?**
Typosquatting (installing `react-dom` vs `react-dom`). Scoped (`@facebook/react`) provides authenticity assurance.

**Q775: How do you sign npm packages for verification?**
npm provides signature verification using `keybase` or PGP (mostly deprecated/legacy). Modern: use `npm audit signatures` (registry signatures).

## 🔹 Worker Threads & CPU-Bound Tasks

**Q776: How do you pass complex data between threads in Node?**
Clone algorithm (default). Or `SharedArrayBuffer` (memory sharing). Or `Transferable` (move ownership of ArrayBuffer, zero copy).

**Q777: How do you measure CPU load in a multithreaded Node app?**
`os.loadavg()`. Or monitor each worker's event loop lag.

**Q778: How do you handle thread crashes in a safe way?**
Listen to `worker.on('error')` and `worker.on('exit')`. Respawn worker if it died unexpectedly.

**Q779: When should you use `SharedArrayBuffer` in Node?**
When manipulating massive datasets (Video/Image/Number Arrays) in parallel where copying data between threads is too expensive (slow).

**Q780: How do you terminate a runaway thread?**
`worker.terminate()`. Forcefully stops execution.

## 🔹 Multi-Region & Geo-Awareness

**Q781: How do you direct requests to region-specific backends in Node?**
DNS Geo-routing (Route53) directs user to nearest IP. This hits Node cluster in that region.

**Q782: How do you manage latency-aware routing in a Node API?**
Middleware checks region header. If request data lives in EU but user hits US node, forward/proxy the request to EU node or error.

**Q783: How do you sync state across regions in a distributed Node app?**
Global DB (DynamoDB Global Tables, CockroachDB). Or Async replication via Kafka. "Eventual Consistency".

**Q784: What are the challenges of sticky sessions in multi-region Node deployments?**
User moves region -> Session lost. Solution: Store sessions in Global Redis or stateless JWTs.

**Q785: How do you invalidate region-specific cache entries?**
Broadcasting invalidation events to all regions (via Pub/Sub). "Fan-out".

## 🔹 Web Standards in Node

**Q786: How do you implement `fetch()` with full spec compliance in Node?**
Use built-in `fetch` (Node 18+). It supports `Request`, `Response`, `Headers` classes.

**Q787: What is the AbortController pattern and how is it used?**
Standard generic cancellation. `const ac = new AbortController(); fetch(url, { signal: ac.signal }); setTimeout(() => ac.abort(), 1000);`.

**Q788: How do you polyfill newer web APIs in older Node versions?**
`core-js` or dedicated polyfills (`whatwg-fetch`).

**Q789: How do you use Headers and FormData objects in Node streams?**
Use matching classes from `undici` or `node:http`. `new FormData().append('file', stream)`. Pass to fetch.

**Q790: How do you simulate a service worker in Node.js for SSR?**
You can't fully. But you can use libraries like `msw` (Mock Service Worker) to intercept requests at the network level for testing/prototyping.

## 🔹 Analytics & Event Tracking

**Q791: How do you implement real-time event tracking in Node?**
Use WebSockets or SSE to stream events to clients. Use a high-write DB (ClickHouse/Timescale) to ingest event stream.

**Q792: How do you debounce high-volume client events server-side?**
Group events by user ID. Process/Flush every X seconds.

**Q793: How do you batch telemetry for export to a data warehouse?**
Buffer in memory/file. Every 1 min upload a Parquet/JSON file to S3 (Data Lake).

**Q794: How do you scrub PII before logging user behavior?**
Middleware: `safeLog(obj)`. Recursive traverse. Hash or Mask keys (`email`, `password`, `ssn`).

**Q795: How do you A/B test features at the API level?**
Hash User ID -> Modulo 100. If < 50 return Feature A, else B. Log assignment.

## 🔹 Cultural / Open Source Knowledge

**Q796: What is the Node.js Release Working Group?**
Team managing release schedules, LTS promotions, and backporting commits to stable branches.

**Q797: How do you propose a new feature to core Node?**
Open an Issue/PR in `nodejs/node`. Write an RFC. Discuss in meetings.

**Q798: How do you write a good README for a Node library?**
Badges, Install, Usage, API Ref, Contributing, License.

**Q799: What are stability indices in Node.js core modules?**
0-Deprecated, 1-Experimental, 2-Stable, 3-Legacy.

**Q800: How do you contribute test coverage to the Node.js core repo?**
Clone repo. Build source. Run tests. Find gaps. Add test case in `test/parallel` directory. Open PR.

## 🔹 Node.js Internals & Engine Behavior - Batch 9

**Q801: What is the role of the `Bootstrapper` phase in Node.js startup?**
It sets up the JS environment, creates the global object, and loads internal modules (built-ins) before executing user code.

**Q802: How does Node.js initialize the event loop in C++ bindings?**
In `src/node.cc`, it instantiates `uv_default_loop()`. It associates this loop with the `Environment` instance, allowing JS callbacks to be scheduled on it.

**Q803: What is libuv and how does Node.js use it under the hood?**
A C library for cross-platform asynchronous I/O. Node.js delegates File, Network, and Timer operations to libuv, which processes them on OS threads or via epoll/kqueue.

**Q804: How does Node.js map JavaScript timers to the libuv event loop?**
`setTimeout` creates a `TimerWrap` handle in C++. This handle is inserted into libuv's timer heap map. When the loop iterates, it checks the heap for expired timers.

**Q805: How is the microtask queue managed differently from the macrotask queue?**
V8 manages the microtask queue (Promises). Node.js flushes this queue *immediately* after every individual operation in the macrotask queue (libuv phase), ensuring promises resolve eagerly.

**Q806: What is the internal representation of `Buffer` in memory?**
A V8 `Uint8Array`. For small buffers, it's on V8 heap. For large buffers, it points to external memory allocated via `malloc` (outside V8 heap) to avoid GC pressure.

**Q807: What is `process.binding()` and when should it be avoided?**
An internal bridge to C++ layers. Deprecated and unsafe for user code. Use `require('module')` instead. Most bindings are now hidden behind `internal/` loaders.

**Q808: How does Node.js bridge the V8 heap and native heap?**
Using `ArrayBuffer` with an `External` backing store. V8 knows the object exists but data relies on C++ memory. GC uses "External Memory" accounting to know when to free native pointers.

**Q809: What’s the role of `libcrypto` in Node.js native modules?**
Part of OpenSSL. Node's `crypto` module links statically to it. Native addons can also link to it to perform hashing/encryption sharing the same binary context.

**Q810: How does Node.js avoid blocking the loop when interfacing with file descriptors?**
By using non-blocking sockets (`O_NONBLOCK`). For FS operations (which are blocking on some OSes), libuv uses a Thread Pool (default 4 threads) to run the task and signal completion.

## 🔹 Advanced Event Loop Scenarios

**Q811: How do you simulate starvation in the Node event loop?**
Schedule frequent `process.nextTick` recursively. Since `nextTick` processes newly added ticks *before* moving to I/O, the loop will never advance to I/O phase.

**Q812: What happens if a promise chain throws asynchronously?**
If no `.catch()`, it bubbles to `unhandledRejection`. It does *not* crash the process by default (legacy), but modern Node (v16+) will crash with exit code 1.

**Q813: How can `setImmediate()` behave differently on different OS platforms?**
It shouldn't (it's consistent). However, `setTimeout(0)` vs `setImmediate()` order is non-deterministic *unless* inside an I/O callback (where Immediate always runs before next Timers).

**Q814: How does the event loop prioritize nextTick vs. setTimeout?**
`nextTick` is NOT part of the loop phases; it runs *between* phases (and operations). `setTimeout` runs in the Timers phase. `nextTick` always wins (runs sooner).

**Q815: What are the risks of deeply recursive `process.nextTick()` calls?**
Since the tick queue must drain completely before I/O, deep recursion blocks I/O—stopped server handling, ping timeouts, effectively a DoS.

## 🔹 Experimental Node Features

**Q816: What are Node.js snapshots and what use cases do they solve?**
`node --build-snapshot`. Serializes the Heap state after initialization. Startup restores this state instantly, skipping parsing/compilation of dependencies. (Fast Startups).

**Q817: How do you enable and test a new experimental Node.js feature?**
Pass flags like `--experimental-feature-name` (e.g., `--experimental-test-runner`). Feature might change/break in minor versions.

**Q818: How does Node.js support policies (`--experimental-policy`)?**
Security Policies (`policies.json`). Restricts which modules can be loaded, can check integrity (SRI), and limit access to globals. Like a sandbox manifest.

**Q819: What is the purpose of the `--frozen-intrinsics` flag?**
Freezes `Array.prototype`, `Object.prototype`, etc., before user code runs. Prevents "Prototype Pollution" or polyfills from modifying core behavior.

**Q820: What is the `node:test` module and how does it differ from popular testing libraries?**
Native test runner (v20 stable). No need for `jest`/`mocha`. Fast, lightweight, supports Test Anything Protocol (TAP). `import { test } from 'node:test'`.

## 🔹 Hardware Integration & System Interfaces

**Q821: How do you interact with USB devices using Node.js?**
Use `usb` library (libusb bindings). `usb.getDeviceList()`. Claim interface. Transfer data via endpoints (Interrupt/Bulk).

**Q822: What’s the role of `node-hid` and how is it used?**
Interacts with generic Human Interface Devices (Keyboard, Mouse, Gamepad). Read/Write raw byte reports.

**Q823: How do you read serial port data in Node?**
`serialport` library. `const port = new SerialPort({ path: '/dev/ttyUSB0', baudRate: 9600 })`. Listen to generic Stream events (`data`).

**Q824: How do you write a Node.js app that responds to GPIO input?**
On Linux (Raspberry Pi), use `onoff` or `rpi-gpio`. They map file descriptors at `/sys/class/gpio` to JS events (`watch((err, value) => ...)`).

**Q825: How do you integrate Node.js with camera/microphone on Linux?**
Use `ffmpeg` to capture stream and pipe to Node. Or use V4L2 (Video4Linux) bindings if low-latency frame access is needed.

## 🔹 System Monitoring and Observability

**Q826: How do you hook into OS-level metrics from Node?**
`os` built-in (CPUs, Mem). For advanced (Disk IO, Network stats), use `systeminformation` package which parses `/proc` (Linux) or WMI (Windows).

**Q827: What are the limitations of using `os` module for memory stats?**
`os.freemem()` reports system RAM, not Node's heap. It doesn't tell you if Node is close to *its* limit (V8 limit).

**Q828: How do you monitor file descriptor usage over time?**
Periodically run `fs.open` check or exec `lsof`. In Prometheus, export `process_open_fds` metric.

**Q829: What tools exist for live heap snapshotting in Node?**
`heapdump`, `v8.getHeapSnapshot()`. Dashboard: `pm2` also offers memory profiling features in paid tier.

**Q830: How do you track garbage collection frequency and duration?**
`perf_hooks`. `const obs = new PerformanceObserver((list) => ...); obs.observe({ entryTypes: ['gc'] });`. Log duration of each GC event.

## 🔹 Low-Level Network Tuning

**Q831: How do you control TCP socket options (e.g., `keepAlive`, `noDelay`) in Node?**
`socket.setKeepAlive(true, initialDelay)`. `socket.setNoDelay(true)` (Disables Nagle's algo).

**Q832: What’s the effect of `highWaterMark` in backpressure scenarios?**
(See Q322). Determines buffer size before `write()` returns false. Tuning this impacts RAM vs CPU wakeups ratio (Larger = fewer wakeups but more RAM).

**Q833: How do you detect socket timeouts and resets at runtime?**
Listen to `socket.on('timeout')`. You MUST call `socket.end()` or destroy it manually handling timeouts. `socket.on('error', err => ...)` handles resets (`ECONNRESET`).

**Q834: How do you manage half-open sockets properly?**
If `allowHalfOpen` is true, the socket stays readable after the other side sends FIN. You must eventually call `end()` to close your write side.

**Q835: What is `SO_REUSEADDR`, and how can you enable it from Node?**
Allows binding to a port in TIME_WAIT state. Node enables this by default for `net.Server` (cluster module relies on it).

## 🔹 Edge Use Cases & IoT

**Q836: How do you minimize CPU usage in event-driven IoT Node apps?**
Use streams, avoid polling loops (`setInterval`). Sleep/Suspend operations. Use interrupts (GPIO `watch`) instead of reading pins constantly.

**Q837: How do you queue sensor data reliably during network loss?**
Write data to local disk (SQLite/Log file). When network works, read file and upload. "Store and Forward".

**Q838: How do you implement OTA (over-the-air) updates in Node?**
Download new code (tarball) to temp dir. Verify signature. Replace current app files. Restart service (systemd or PM2 reload).

**Q839: How do you write an energy-efficient sensor polling system?**
Deep sleep hardware (if microcontroller). For Node (Gateway), process data in batches. Reduce heartbeat frequency.

**Q840: What serialization formats are most efficient for IoT Node apps?**
Protobuf or MessagePack. Smaller payload than JSON means less Radio/Network usage (Power savings).

## 🔹 Functional Programming in Node.js

**Q841: How do you enforce pure functions in a Node.js codebase?**
Linting (`eslint-plugin-functional`). Code reviews. Avoid `this`, avoid mutating arguments, avoid side effects (I/O) inside logic functions.

**Q842: How do you implement lazy evaluation in Node pipelines?**
Use Generators (`function*`). Process items one by one (`yield`) instead of creating full arrays.

**Q843: What is a monad and how would you emulate one in Node.js?**
A wrapper object (like Promise or Array) with a `flatMap`/`chain` method. `Maybe` or `Either` patterns (libraries like `fp-ts`) handle null checks elegantly.

**Q844: How can functional constructs help with async error handling?**
Using `Task` or `Future` monads instead of try/catch blocks. Errors are treated as values passed down the chain.

**Q845: How do you structure a fully functional Node.js service layer?**
Dependency Injection (pass dependencies as args). Currying (`const service = (db) => (req) => ...`). Compose small functions.

## 🔹 Node.js with AI/ML Workloads

**Q846: How do you run TensorFlow or ONNX models from Node.js?**
(See Q522, Q523). `onnxruntime-node` or `tfjs-node`.

**Q847: How can you offload AI tasks from Node.js to Python efficiently?**
ZeroMQ or Redis Pub/Sub. Node pushes job to queue. Python worker (PyTorch) processes and replies. Decoupled is better than spawning shells.

**Q848: What are the tradeoffs of using `@tensorflow/tfjs-node`?**
**Pro**: Easy JS API. **Con**: Binary dependency (libtensorflow), large install size, slower than optimized C++/Python for training.

**Q849: How do you use WebAssembly to run ML models in Node?**
Compile model runtime (e.g., small C inference engine) to Wasm. Load model weights into Wasm memory. Run. Universal and sandboxed.

**Q850: How do you handle large matrix computations in a streaming Node setup?**
Don't process all data in RAM. Process windowed chunks. Use `ndarray` libraries that support views/strides to avoid copying data.

## 🔹 Immutable Infrastructure & IaC

**Q851: How do you provision a Node.js service using Terraform?**
Define `resource "aws_lambda_function"` or `aws_ecs_service`. Zip Node code. Terraform uploads and configures env vars.

**Q852: What is a golden AMI and how would you build one with Node pre-installed?**
Use Packer. Base OS -> Install Node -> Install App -> Snapshot. Deploy instances from this Snapshot for fast scaling.

**Q853: How do you integrate infrastructure tagging into a Node deploy pipeline?**
In CI, read Git tags/Branch. Pass as `-var "tags={Env=Prod}"` to Terraform/CloudFormation during deploy step.

**Q854: How do you use Node.js to orchestrate IaC commands like Terraform or Pulumi?**
Pulumi *is* Node.js (you write infra in TS). For Terraform, spawn `terraform apply`. Wrapper scripts/CDK for Terraform (cdktf).

**Q855: How do you test infrastructure changes using Node?**
Use `terratest` (Go) or Pulumi's integration testing framework (Mocha/Jest) to assert that buckets/instances exist after provisioning.

## 🔹 CLI and Tooling Architectures

**Q856: How do you handle plugin-based architectures in CLI tools?**
Search user home dir (`~/.config/mycli/plugins`) for node modules. Require them. Add their exported commands to the Yargs instance.

**Q857: What’s the difference between ESM and CJS CLI bootstrapping?**
ESM needs `"type": "module"`. Extensions `.mjs`. `__dirname` unavailable. Async top-level await is allowed (great for CLI setup).

**Q858: How do you dynamically generate completions for a CLI tool?**
Expose a hidden command `mycli --get-completions "curr"`. Shell script calls this. Node app calculates suggestions based on context.

**Q859: How do you ensure CLI tools work across different shell environments?**
Use `cross-spawn` (handles Windows shebangs). Avoid shell-specific syntax (`&&`, `|`) in child_process exec strings—use logic in JS.

**Q860: How do you provide interactive multi-step workflows in Node CLIs?**
State machine. `step1 -> input -> step2 -> input`. If error, retry current step. `inquirer` loop.

## 🔹 Uncommon Build & Packaging Cases

**Q861: How do you build a Node.js app for embedded Linux devices?**
Cross-compile Native Addons. Bundle everything (including `node_modules`) into a tarball or use `pkg` to make a single binary (easier deployment).

**Q862: What is a static build of Node.js and why might you need it?**
A node binary linked with static libraries (musl libc instead of glibc). Portable (runs on Alpine and Ubuntu without dependencies).

**Q863: How do you bundle native binaries into a Node module?**
Include binaries for all platforms (`bin/linux`, `bin/win`). At runtime, detect OS (`process.platform`) and spawn the correct binary.

**Q864: How do you strip debug symbols from a compiled module?**
Run `strip -s module.node`. Reduces size significantly (important for IoT/Edge).

**Q865: How do you build a hybrid WebAssembly + Node.js library?**
Write core logic in Rust/C -> Wasm. Wrap in TS/JS for the API. `pkg` interacts with Wasm. Publish as npm package.

## 🔹 Niche Module System Challenges

**Q866: How do you resolve dynamic import paths securely?**
Whitelist allowed paths. `const allowed = ['a', 'b']; if(!allowed.includes(input)) throw; import(input)`.

**Q867: How do you sandbox CJS modules within an ESM app?**
`createRequire(import.meta.url)`. Use `require` inside the specific scope.

**Q868: How do you simulate a virtual filesystem in ESM?**
Use compilation hooks (Loaders). `verifyUrl` hook can intercept `import 'virtual:foo'` and return code dynamically generated in memory.

**Q869: What are import assertions and when are they used?**
`import json from './data.json' assert { type: 'json' };` (Syntax changing to `with`). Ensures the server sends correct MIME type/interprets as JSON, not code.

**Q870: What’s the difference between static vs. dynamic export maps?**
Export maps in package.json are static. Dynamic resolution requires a custom loader hook to change resolution at runtime.

## 🔹 Contract Testing and Protocol Compliance

**Q871: What is contract testing and how do you use it with Node APIs?**
Verifying that Provider (API) and Consumer (Frontend) agree on the format (Contract). Use Pact. Consumer generates pact file. Provider replays it to verify compliance.

**Q872: How do you enforce schema compatibility using tools like Pact?**
Pact Broker. Check "Can I deploy?". If Provider changes schema breaking the consumer's pact record, deploy is blocked.

**Q873: How do you test custom TCP protocols in Node?**
Write a "Dummy Client" in tests sending Buffer packets. Assert Server response Buffers match expected hex values.

**Q874: How do you ensure gRPC compatibility across versions?**
Use `buf` or `protolock` to detect breaking changes in `.proto` files in CI before code generation.

**Q875: How do you fuzz test a Node HTTP server?**
Send random/garbage bytes using a fuzzer (AFL connected to Node, or generic HTTP fuzzers like `radamsa`). Watch for crashes/hangs.

## 🔹 Time-Sensitive Code & Clock Drift

**Q876: How do you compensate for system clock drift in Node.js?**
Use NTP to sync OS clock. In app, allow tolerance (slight skew) for timestamps. For strict intervals, measure drift (`Date.now() - expectedTime`) and adjust next timeout.

**Q877: How do you synchronize time between Node instances?**
You don't sync the instances directly. You rely on the underlying OS NTP sync. Use logical clocks (Lamport/Vector) if ordering matters more than wall clock.

**Q878: How do you ensure consistency for time-based cache keys?**
Round time to nearest interval (flooring). `key = 'stats_' + Math.floor(Date.now() / 60000)`. All servers agree on the minute bucket.

**Q879: How do you safely use `Date.now()` in a distributed system?**
Don't use it for ordering events across machines. Use it only for duration/local timeouts. Use DB timestamps or Vector Clocks for causality.

**Q880: How do you detect and correct timestamp anomalies?**
If `timestamp_received < timestamp_previous`, clock moved back. Reject or Log warning. Monotonic clocks (`process.hrtime`) avoid this for measuring duration.

## 🔹 Accessibility (a11y) in Node-Powered Interfaces

**Q881: How do you generate accessible PDF reports with Node?**
Tag PDF structure properly. Use libraries (`pdfmake`, `puppeteer` with accessibility options) that output tagged PDFs (readable by screen readers).

**Q882: How do you check a static site for accessibility from Node?**
Run `pa11y` or `axe-core` against the URL (via puppeteer). Fail build if a11y violations found.

**Q883: What Node libraries assist in screen-reader testing?**
`@accesslint/logger`, `eslint-plugin-jsx-a11y` (linting). `selenium-webdriver` can interrogate accessibility tree.

**Q884: How do you implement ARIA-aware markup using SSR from Node?**
Ensure templates (Pug/EJS/React) render standard HTML5 semantic tags and correct `aria-*` attributes based on state.

**Q885: How can you test keyboard navigation with Node automation tools?**
Puppeteer: `await page.keyboard.press('Tab')`. `await page.evaluate(() => document.activeElement...)`. Verify focus moves correctly.

## 🔹 Knowledge of Node Ecosystem Trends

**Q886: What is the role of `undici` vs. `http` in Node’s future?**
`undici` is the new, faster HTTP/1.1 client written from scratch. It powers global `fetch`. Likely to replace or heavily influence legacy `http` client internals.

**Q887: How does the Node.js release cadence align with npm’s evolution?**
Node releases include a specific npm version. They are decoupled but synced. Major Node often brings Major npm.

**Q888: How is Node.js adapting to the rise of edge runtimes?**
Standardizing APIs via WinterCG (Web-interoperable Runtimes Community Group). Making Node subset compatible with Cloudflare/Deno (Fetch, Streams, Crypto).

**Q889: What is the Node.js Build WG, and why is it relevant?**
Team maintaining the infrastructure (Jenkins, machines) that builds and tests Node on all platforms (ARM, AIX, Windows). Critical for release stability.

**Q890: How do you track Node.js community security advisories?**
Subscribe to `nodejs-sec` mailing list. Follow Common Vulnerabilities and Exposures (CVE) database for Node.js.

## 🔹 Disaster Recovery and Fault Injection

**Q891: How do you simulate a dependency outage in Node?**
Change config to point to invalid host. Or use `toxiproxy` to cut connection mid-stream. Verify app handles timeouts gracefully.

**Q892: How do you simulate memory exhaustion in test environments?**
Run with low max_old_space_size. Trigger object creation. Ensure `catch` handles OOM or process restarts correctly (via external supervisor checking exit code).

**Q893: What is a chaos monkey and how would you write one in Node?**
A script that randomly kills child processes, closes sockets, or deletes temp files while the app runs. `setInterval(() => { randomKill() }, 10000)`.

**Q894: How do you create fault-tolerant retry queues?**
Dead Letter Queue (DLQ). If job fails N times, move to DLQ. Alert humans. Do not retry infinitely.

**Q895: How do you automate incident response simulations?**
"Game Days". Scripts trigger a specific failure (e.g., latency injection). Team follows playbook to fix. Measure MTTR (Time To Recovery).

## 🔹 Legal and Licensing Awareness

**Q896: How do you verify license compatibility in npm dependencies?**
Use `license-checker`. Scan `node_modules`. whitelist: [MIT, ISC, Apache-2.0]. Fail build if GPL (viral) found in proprietary project.

**Q897: How do you exclude GPL packages from a corporate Node project?**
Configure `license-checker` to fail on GPL. Manually review deps. Replace GPL libs with MIT alternatives.

**Q898: What tools help enforce open source license policies?**
FOSSA, Snyk, Black Duck. They scan and automate approval workflows.

**Q899: How do you generate a license report for a Node app?**
`npx license-checker --csv --out licenses.csv`. Bundle this file with your software distribution.

**Q900: How do you comply with attribution requirements in Node CLI tools?**
Include a `ThirdPartyNotices.txt` or a command `mycli --licenses` that prints the licenses of used libraries.

## 🔹 Node.js and Emerging Web Standards - Batch 10

**Q901: How do you use `fetch()` streaming with ReadableStreams in Node?**
`const res = await fetch(url); for await (const chunk of res.body) { ... }`. `res.body` is a standard Web ReadableStream.

**Q902: How does Node.js support `ReadableStream` from Web Streams Standard?**
Exposed via `blobal.ReadableStream` or `stream/web`. It differs from legacy Node streams (`require('stream')`). Use `.fromWeb()` and `.toWeb()` to convert between them.

**Q903: What is the `Blob` object in Node.js and when is it used?**
Immutable raw data object (like in Browser). Used in `fetch` bodies or MessageChannel communications. `new Blob(['text'], { type: 'text/plain' })`.

**Q904: How do you handle `FormData` server-side in Node using Web standard APIs?**
`await req.formData()` (if using a framework that supports Web Request). Or manually parse multipart body into a `FormData` object to forward it to another API.

**Q905: How does Node implement `URLPattern` for matching routes?**
`new URLPattern({ pathname: '/users/:id' }).exec(url)`. Experimental web standard for routing, available in recent Node versions or via polyfill.

## 🔹 Edge Runtime & WASM Awareness

**Q906: What are Node.js isolates in edge computing environments?**
V8 Isolates. Lightweight contexts with their own heap. They share the same underlying process/engine but are logically separate. Low startup cost.

**Q907: How do you share memory between multiple isolates safely?**
You generally don't (that's the point of isolation). You communicate via external storage (KV, Durable Objects) or serialized streams.

**Q908: What are multi-tenancy constraints when deploying Node on edge runtimes?**
Strict CPU limits (e.g., 10ms-50ms CPU time per request). Memory limits (128MB). No filesystem access. No long-running background tasks.

**Q909: How do you dynamically load WebAssembly at runtime in edge environments?**
`WebAssembly.instantiate(buffer, imports)`. Wasm is often a first-class citizen in edge runtimes (Cloudflare calls it "Wasm Workers").

**Q910: How do you detect and handle runtime limitations (e.g. memory/performance) in edge Node?**
Check `performance.memory` (if available) or catch `RangeError: Array buffer allocation failed`. Fail gracefully (return 503).

## 🔹 Secure Multi‑Process Data Sharing

**Q911: How do you share data between a parent and child process securely?**
IPC channel (`child.send()`). Validate structure on receipt. Do not `eval()` or blindly trust the payload.

**Q912: What are the pitfalls of serializing large objects via IPC?**
Serialization overhead (JSON stringify/parse) blocks the event loop. The channel pipe buffer might fill up, causing backpressure or stalls.

**Q913: How do you revoke shared buffers to prevent memory leaks?**
You can't "revoke" a SharedArrayBuffer easily unless you use `Atomics` or just let it get GC'd when all threads drop references. `MessagePort` can be closed (`port.close()`).

**Q914: How do you implement capability-based security between Node worker threads?**
Pass a specific `MessagePort` to the worker. Only allow it to communicate via that port. Do not share global handles.

**Q915: How can circular references affect structured clone in IPC?**
The HTML structured clone algorithm (used in `postMessage`) *supports* circular references (unlike `JSON.stringify`), so it works fine, but deep graphs are slow to copy.

## 🔹 Real‑Time Geo & Location Systems

**Q916: How would you build a WebSocket-based geolocation tracker in Node?**
Clients emit `{ lat, lng }`. server pushes to Redis Geospatial index (`GEOADD`). Server queries nearby (`GEORADIUS`) and emits to neighbors.

**Q917: How do you perform geospatial queries at scale from Node?**
Use PostGIS or Mongo `$near`. Index the location. Node just executes the query. Don't calculate Haversine distance in JS loop for millions of points.

**Q918: How do you handle latency-sensitive GPS data streams in Node apps?**
Use UDP (dgram) instead of TCP/WS if packet loss is acceptable but latency isn't. Or use compact binary protocols (Geobuf).

**Q919: How do you manage geofencing logic in a distributed Node system?**
Shard users by S2 cell or Geohash. The node responsible for "Cell X" checks if users inside breach the polygon bound.

**Q920: How do you simulate geo‑failover in real‑time Node clusters?**
Manually sever the connection to "Region A" DB or Cluster. Observe if Traffic Manager redirects via DNS/Anycast to "Region B".

## 🔹 Streaming Machine‑Generated Media

**Q921: How do you stream synthetized audio in real time using Node?**
Generate PCM chunks. Pipe to `ffmpeg` to encode AAC/Opus. Pipe to HTTP Response (`Content-Type: audio/mpeg`).

**Q922: How do you support live video transcoding pipelines in Node backends?**
Node acts as controller. Spawns `ffmpeg` process. Input RTMP -> FFmpeg -> HLS (.m3u8/.ts segments). Node serves the static segments.

**Q923: How do you merge multiple live media streams into one feed?**
FFmpeg complex filter `[0:v][1:v]overlay`. Node spawns the process managing input streams and piping the single output.

**Q924: How do you control media quality adaptively over unstable networks?**
(ABR). Generate multiple bitrates (Low, Mid, High). Client player switches. Node just serves the requested index file (`master.m3u8`).

**Q925: What is the role of `MediaStreamTrackProcessor` in Node?**
Part of WebRTC Insertable Streams. Allows raw manipulation of video frames/audio chunks in JS. (Mostly browser, but coming to server-side WebRTC stacks).

## 🔹 Hardware‑Accelerated Cryptography

**Q926: How do you leverage hardware-backed key stores from Node?**
Use `crypto.setEngine('openssl-engine-pkcs11')`. Configure OpenSSL to use HSM or TPM.

**Q927: How do you interface with TPM or secure enclave modules in Node?**
Use `tpm2-tss` bindings. Or execute CLI tools `tpb` via child_process.

**Q928: How do you accelerate cryptographic operations with native bindings?**
`crypto` module already uses native C++ (OpenSSL). For specific non-standard algos (e.g. BLS signatures), use C++ addon (`node-gyp`).

**Q929: How do you handle hardware‑backed signature verification in Node?**
The private key stays in Hardware. You send hash to Hardware. Hardware returns Signature. Node verifies signature with Public Key (software).

**Q930: How do you offload crypto to GPU or hardware modules securely?**
Use OpenCL/CUDA bindings for bulk hashing (e.g. password cracking context). For general SSL, use OpenSSL engine offloading (Intel QAT).

## 🔹 Typed API Contracts & Runtime Validation

**Q931: How do you integrate TypeScript runtime type validation?**
TS is compile-time. For runtime, use schema libraries: `zod`, `io-ts`, `joi`. Define schema, infer TS type from it. `schema.parse(data)`.

**Q932: How do you auto-generate validation middleware from an OpenAPI spec?**
`express-openapi-validator`. It reads `api.yaml`, matches route, validates `req.body` against schema. 400 if invalid.

**Q933: How do you enforce type-safe GraphQL schemas in Node runtime?**
`graphql-code-generator`. Generates TS types for Resolvers. `const resolver: Resolvers<Context> = ...`.

**Q934: How do you support end-to-end type safety between front-end and Node backends?**
Share `types` package (Monorepo). Or use TRPC (TypeScript Remote Procedure Call). Client imports server router type. RPC calls are typed.

**Q935: How do you validate WebSocket message payloads at runtime with type generators?**
`zod`. `ws.on('message', data => { const parsed = MessageSchema.safeParse(JSON.parse(data)); ... })`.

## 🔹 Low‑Latency Finance & Trading Systems

**Q936: How do you manage sub‑millisecond latency in WebSocket data feeds via Node?**
Disable Nagle (`setNoDelay`). Use `uws` (uWebSockets.js) C++ bindings instead of native `ws` (faster). Avoid GC (pre-allocate objects).

**Q937: How do you aggregate real-time market data streams in Node?**
Circular Buffer in memory. Ingest updates. Compute moving averages. Push snapshot to clients.

**Q938: How do you avoid GC pauses in latency‑sensitive trading systems?**
Use `off-heap` storage (Buffers/SharedArrayBuffer). Trigger GC manually (`global.gc()`) during market inactivity/idle times.

**Q939: What techniques reduce latency in Node serialization/deserialization?**
Use Schema-based binary formats (SBE, FlatBuffers). Zero-copy parsing. Avoid JSON.

**Q940: How do you recover from failed heartbeats in live trading system processes?**
Failover immediately. Have hot standby Node process subscribing to same feed. If Primary misses heartbeat, Standby takes over (VIP switch).

## 🔹 Hybrid Cloud & Multi‑Cloud Architecture

**Q941: How do you manage cross‑cloud cluster communication between Node services?**
VPN/Direct Connect between clouds. OR public internet with mTLS and strict IP whitelisting.

**Q942: How do you route data based on cloud‑region affinity?**
DNS policies. or Application Logic: "If User.region == 'aws-us-east', connect to AWS endpoint."

**Q943: How do you ensure data consistency across multi-cloud Node databases?**
Use databases designed for this (CockroachDB, Spanner). Node just targets the local endpoint. The DB handles Paxos/Raft replication.

**Q944: How do you handle failover for Node APIs between cloud providers?**
Global GSLB (Global Server Load Balancer) like Cloudflare/NS1. If AWS health check fails, update DNS to point to GCP IP.

**Q945: How do you unify logging and tracing across multi‑cloud Node deployments?**
Ship all logs to a third-party SaaS (Datadog/Splunk) or a centralized ELK stack independent of the cloud provider.

## 🔹 Edge Caching & Pre‑Rendering Strategies

**Q946: How do you aggressively pre-render API responses at build time?**
SSG (Static Site Generation). Run Node script to fetch data, generate JSON/HTML, save to disk, upload to CDN.

**Q947: How do you invalidate edge‑cached HTML generated by Node?**
Purge by URL or Cache Tag. `fetch('https://api.cloudflare.com/.../purge_cache', ...)` triggered by CMS update webhook.

**Q948: How do you handle real‑time user‑personalized content via edge rendering?**
Edge Side Includes (ESI). CDN caches the layout. A small dynamic hole `<esi:include src="/user-widget" />` is fetched from Node origin per request.

**Q949: How do you detect bot crawlers for caching logic in Node SSR?**
`is-bot` package. If bot, render full HTML (costly). If user, render skeleton + hydration (fast).

**Q950: How do you fallback to live Node render only when cache misses?**
CDN Logic: "If file exists in S3, serve it. Else, forward to Node Origin (which renders and uploads to S3 for next time)." (ISR).

## 🔹 Semantic Code Evolution & Migration

**Q951: How do you use codemods to upgrade Node codebases safely?**
`jscodeshift`. Write a transform script (Find `require`, replace with `import`). Run across all files.

**Q952: How do you automatically refactor deprecated API usages?**
AST transformation. Detect `fs.exists` (deprecated). Replace with `fs.access` or `fs.existsSync`.

**Q953: How do you enforce upgrade paths for Node major versions?**
CI check. `engines`: { "node": ">=20" }. Fail build on old Nodes. Automated PRs to bump version.

**Q954: How do you evolve data models across live running Node nodes?**
Expand-Contract pattern. Add new field. Write to both. Read from new (fallback to old). Backfill. Remove old.

**Q955: How do you guarantee schema consistency during rolling migrations?**
Code must be compatible with *both* old and new schema variants during the transition phase.

## 🔹 Simulated Time & Debugging Complex State

**Q956: How do you freeze time in Node to debug scheduled tasks?**
`sinon.useFakeTimers()`. "Ticks" the clock forward programmatically. `clock.tick(1000)`.

**Q957: How do you simulate network partitions for testing resilience?**
Use `iptables` or docker network disconnect. In Node, mock the network client to throw "ETIMEDOUT".

**Q958: How do you replay historical production traffic in test environments?**
Capture raw requests (GoReplay). Replay against Staging environment. Node app processes them as if real.

**Q959: How do you checkpoint state in-memory to rewind runtime state?**
Redux-like state container. Save state tree to array. "Time Travel" debugging.

**Q960: How do you emulate partial shearing or thread preemption manually?**
Inject random `await new Promise(r => setTimeout(r, Math.random()))` in async logic to shake out race conditions.

## 🔹 Declarative Infrastructure within Node Processes

**Q961: How do you define and manage infra via a Node app (e.g. Pulumi program)?**
`import * as aws from "@pulumi/aws"; const bucket = new aws.s3.Bucket("my-bucket");`. Run `pulumi up`.

**Q962: How do you orchestrate blue/green deployment within a Node process?**
Traffic shifting. Node Proxy (Gateway) points 100% to Blue. Spin up Green. Health check. Update Proxy to 10% Green... 100% Green.

**Q963: How do you dynamically wire services at runtime via a declarative model?**
Service Discovery. Node app asks Consul "Where is Service B?". Connects. If topology changes, app receives update event and reconnects.

**Q964: How do you trigger infrastructure changes via business‑logic code in Node?**
AWS SDK. `const ec2 = new EC2(); await ec2.runInstances(...)`. (e.g., spinning up render nodes on demand).

**Q965: How do you rollback infra as easily as application code?**
Declarative IaC (Terraform/Pulumi). `git revert`. CI applies the old state.

## 🔹 Edge‑AI Inference Pipelines

**Q966: How do you run on-device ML inference in Node within constrained environments?**
TensorFlow Lite (TFLite). Optimized for edge/mobile. `tfjs-tflite-node` bindings.

**Q967: How do you offload AI compute to specialized edge hardware?**
Coral Edge TPU. Use `coral-tpu` node bindings. Send buffer. Hardware infers. Return result.

**Q968: How do you manage model updates and versions across devices?**
IoT Shadow pattern. Device reports "Model v1". Cloud says "Desired: v2". Device downloads model blob, verifies, reloads worker.

**Q969: How do you securely fetch new model binaries at runtime?**
Signed URLs (S3). Verify Checksum/Signature of file after download before loading into heap.

**Q970: How do you certify and audit inference outputs across deployments?**
Log `(InputHash, ModelVersion, Output, Confidence)`. Analyzer checks for drift or bias.

## 🔹 Temporal & Event‑Driven Orchestration

**Q971: How do you use Temporal or Cadence with Node clients for workflows?**
(See Q683/972). `temporal-io/sdk`. Write workflows as pure functions that yield activities. Platform handles retries/history.

**Q972: How do you coordinate multi-step saga patterns in Node services?**
Orchestrator sends "Reserve Credit". If success, "Ship Item". If fail, "Refund Credit" (Compensating Transaction).

**Q973: How do you simulate event replays and compensating logic?**
Store events in Event Store. Reset "Read Position" to 0. Reprocess all events. Ensure code is Idempotent.

**Q974: How do you safely store workflow states across crashes?**
(Temporal functionality). Or manual: Write "State: Step 2" to DB *before* executing Step 2. On restart, read DB, resume Step 2.

**Q975: How do you manage versioning of workflows in Node runtime?**
`patch` method in Temporal. `if (workflowVersion < 2) { oldLogic() } else { newLogic() }`. Keep history consistent.

## 🔹 Edge‑First Metrics & Healthchecks

**Q976: How do you run edge-aware health checks for Node services?**
Probe from multiple regions. If US-East fails but EU-West works, only route traffic away from US-East.

**Q977: How do you aggregate and sample logs from edge nodes?**
Edge nodes send UDP packets to a collector (Vector). Collector aggregates and ships to Central (Elastic).

**Q978: How do you expose fine‑grained metrics (e.g. per‑CPU core) in Node?**
`os.cpus()`. Calculate usage diff over 1 sec. Expose via `/metrics` endpoint for Prometheus scraping.

**Q979: How do you dynamically adjust health TTLs based on workload?**
If load is high, reduce health check frequency (prevent thundering herd). If idle, check frequently.

**Q980: How do you handle partial network failures gracefully in health probes?**
Retry with jitter. Do not mark unhealthy immediately on one timeout. "Flapping" detection.

## 🔹 AI‑Generated Test Suites & Mutation Testing

**Q981: How do you auto‑generate tests using GPT‑like models from spec?**
Feed OpenAPI spec + prompt to LLM: "Generate Jest tests for this API". Review code. Run.

**Q982: How do you integrate mutation testing into Node pipelines?**
Stryker Mutator. `npx stryker run`. It modifies your code (reverses logic) and runs tests. If tests pass (mutant survives), your tests are weak. kill the mutant.

**Q983: What are the risks of auto‑generated code tests?**
False confidence. AI might write trivial tests (`expect(true).toBe(true)`). Always verify assertion logic.

**Q984: How do you validate coverage from AI‑generated test cases?**
Run coverage tool (`nyc` / `jest --coverage`). Ensure critical paths are hitting green.

**Q985: How do you maintain human oversight in AI test suites?**
Code Review. Treat test code as production code. Do not just auto-merge AI output.

## 🔹 Signal & Interrupt Management

**Q986: How do you capture and respond to POSIX signals (`SIGUSR1`, etc.) in Node?**
`process.on('SIGUSR1', () => { ... })`. (Often used for debugging/reloading config).

**Q987: How do you implement custom signal handlers in a multi-threaded Node app?**
Signals go to the main process usually. Main process must message workers: "Shutdown". Workers handle message and exit.

**Q988: How do you restore state safely after an interrupt-driven crash?**
Crash-only software. Maintain state in DB/Journal. Restart is same as fresh start. Recovery is automatic.

**Q989: How do you coordinate signal-handling across clustered processes?**
Cluster Master receives signal. Forwards to all workers. Waits for them to exit. Then exits itself.

**Q990: How do you trigger runtime behaviors (e.g. dump stats) via signals?**
`process.on('SIGUSR2', writeStatsToDisk)`. Admin sends `kill -SIGUSR2 <pid>`.

## 🔹 Standards Compliance & Interop Testing

**Q991: How do you test compliance with JSON-RPC protocol in Node?**
Use a compliance suite (if available). Or write tests ensuring correct Error codes (-32600, etc.) and Response structure (`jsonrpc: "2.0"`).

**Q992: How do you validate GraphQL server behavior matching the spec?**
Run `graphql-cats` (Compatibility Acceptance Tests) against your engine.

**Q993: How do you run WebSocket conformance test suites in Node?**
Autobahn Testsuite. Connects to your ecosystem, runs hundreds of fuzz/protocol tests (Frames, fragmentation, pings).

**Q994: How do you check `Multipart/form-data` correctness under large file stress?**
Upload 10GB file. Ensure checksum matches. Ensure boundary parsing doesn't buffer entire file (OOM).

**Q995: How do you ensure `HTTP/3` support works correctly in Node behind proxies?**
Test UDP connectivity. Verify `Alt-Svc` header. Ensure QUIC handshake succeeds using `curl --http3`.

## 🔹 Ethical & Privacy‑Aware Patterns

**Q996: How do you enforce PII minimization in Node logs and analytics?**
Data Masking middleware. Remove `email`, `phone` from log objects before `JSON.stringify`.

**Q997: How do you support GDPR-level data access / erasure requests?**
Implement "Forget Me" API. It orchestrates deletion across DBs, Logs (if possible, usually rotation), and backups (crypto-shredding).

**Q998: How do you store and rotate consent metadata securely?**
Store "User X consented to Cookie Policy v1 on Date". If Policy v2, prompt again. Immutable Audit log.

**Q999: How do you design APIs that respect user tracking preferences?**
Respect `DNT` (Do Not Track) or `GPC` (Global Privacy Control) headers. Disable analytics for those requests.

**Q1000: How do you build audit trails in Node for privacy compliance?**
Log: "Admin A viewed User B's profile". Store in tamper-evident log (Write-Once storage). Required for HIPAA/SOC2.









