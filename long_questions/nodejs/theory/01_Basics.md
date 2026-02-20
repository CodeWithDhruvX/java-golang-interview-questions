# 🟢 **1–20: Basics (Beginner Level)**

### 1. What is Node.js?
"**Node.js** is an open-source, cross-platform JavaScript runtime environment built on Chrome's V8 JavaScript engine. 

It allows me to run JavaScript outside the browser, typically on a server. Node.js uses an event-driven, non-blocking I/O model, which makes it incredibly lightweight and efficient for data-intensive real-time applications.

I use it primarily for building scalable backend services, REST APIs, and microservices because it handles thousands of concurrent connections smoothly without the overhead of thread context switching."

#### Indepth
Unlike traditional server environments that spawn a new thread for every request (like Apache or Tomcat), Node.js operates on a single thread using an event loop. The underlying libuv library handles asynchronous I/O operations (like reading files, network requests) by offloading them to the system kernel or a thread pool, allowing the main thread to remain unblocked and highly performant.

---

### 2. How is Node.js different from JavaScript in the browser?
"While both use JavaScript, the contexts are entirely different. 

**Browser JavaScript** interacts with the DOM (Document Object Model) and the BOM (Browser Object Model). It handles UI updates, button clicks, and client-side logic. It doesn't have secure access to the local file system or OS.

**Node.js**, on the other hand, runs on the server. It has no DOM or `window` object. Instead, it provides access to the file system (via the `fs` module), network sockets, and OS processes (via the `process` object). 

In short: Browser JS controls what the user sees, while Node JS controls the data and services behind the scenes."

#### Indepth
Node.js injects specific C++ bindings (via the Node API) into the V8 engine to provide OS-level access. Consequently, global objects differ: the browser has `window` and `document`, whereas Node.js has `global` and `process`. Both environments implement the ECMAScript specification, but their ambient APIs are drastically different.

---

### 3. What is the V8 engine?
"The **V8 engine** is an open-source, high-performance WebAssembly and JavaScript engine developed by Google for the Chrome browser. 

It’s written in C++ and its primary job is to take raw JavaScript code and compile it directly into native machine code (rather than interpreting it in real-time or compiling it to an intermediate bytecode). 

Node.js wraps this engine to provide a server-side runtime. Whenever I run a Node.js script, it's actually the V8 engine doing the heavy lifting to execute my JavaScript at lightning speed."

#### Indepth
V8 achieves its speed through Just-In-Time (JIT) compilation. It uses two compilers: **Ignition** (a fast interpreter that generates bytecode) and **TurboFan** (an optimizing compiler that turns frequently executed "hot" paths into highly optimized machine code). It also employs a highly efficient generational garbage collector to manage memory.

---

### 4. What is the role of the `package.json` file?
"The `package.json` file is the heart of any Node.js project. It acts as the manifest for the application.

It stores critical metadata about the project, such as the name, version, and author. Most importantly, it keeps track of two things: **dependencies** (packages needed for production) and **devDependencies** (packages needed only for development, like testing frameworks). 

Whenever I clone a Node repository, simply running `npm install` looks at `package.json` and fetches the exact packages needed so the app can run on my local machine instantly."

#### Indepth
The exact version resolutions of the `package.json` ranges (like `^1.2.3`) are locked down by the `package-lock.json` file. The `package.json` also defines custom `scripts` (like `"start": "node index.js"`, `"test": "jest"`), making it standard for standardizing task execution across development teams.

---

### 5. What is npm? How is it different from npx?
"**npm** (Node Package Manager) is the default package manager for Node.js. It allows developers to install, share, and manage third-party libraries. When I type `npm install express`, npm downloads the Express library into my `node_modules` folder.

**npx** stands for Node Package Execute. It comes bundled with npm. While npm installs packages, npx *executes* them. 

I use npx when I want to run a CLI tool once without installing it globally on my machine. For example, `npx create-react-app my-app` fetches the latest version of the generator, builds the app, and leaves no global footprint."

#### Indepth
If a package is already installed locally in `node_modules/.bin`, npx will use that version. If it’s not installed, npx downloads a temporary copy, executes it, and deletes it. This avoids global dependency bloat and ensures you are always executing the latest version of a tool (unless a specific version is requested).

---

### 6. What are modules in Node.js?
"Modules in Node.js are encapsulated blocks of code—essentially independent files that provide specific functionality. 

By default, every file in Node.js is treated as a separate module to prevent polluting the global namespace. Functions and variables defined in one file remain private to that file unless they are explicitly exported.

This modular architecture is crucial because it allows me to split a massive backend application into small, maintainable, and highly reusable pieces (e.g., separating database controllers from routes)."

#### Indepth
Historically, Node.js utilized the **CommonJS** module system relying on synchronous `require()` and `module.exports`. In modern Node.js, standard **ES Modules (ESM)** using `import` and `export` are fully supported, providing a unified standard between the frontend and backend.

---

### 7. How do you export and import modules in Node.js?
"In the traditional CommonJS system, you export data using `module.exports` and bring it into another file using `require()`. For example, `module.exports = myFunction;` and later, `const myFunc = require('./myFunction');`.

If I have multiple things to export, I attach them to the `exports` object: `exports.foo = foo`.

In the modern ES Module (ESM) approach, I use the `export` keyword before functions or variables, and then use `import { foo } from './myFile.js'` in the receiving file. This is generally the method I prefer for new projects."

#### Indepth
In CommonJS, `require()` is fundamentally synchronous; it acts as an immediate blocking file read at runtime. ES Modules, however, are asynchronous and statically analyzed during a parsing phase before execution. This static analysis enables advanced optimizations like **Tree Shaking** (eliminating dead code).

---

### 8. What is the difference between `require()` and `import`?
"`require()` is part of the CommonJS specification (Node's original system). It is synchronous, can be called dynamically anywhere in the code (e.g., inside an `if` statement), and resolves at runtime.

`import` is part of the ES6 Module standard. It is asynchronous, statically analyzed, and must be declared at the top level of the file (though dynamic `import()` exists).

I use `import` whenever possible because it’s the standard across the entire JavaScript ecosystem (browser and server), leading to better tooling support and leaner bundles."

#### Indepth
To use `import` natively in Node.js, you must set `"type": "module"` in `package.json` or use the `.mjs` file extension. The top-level, static nature of `import` ensures that cyclic dependencies are handled differently (by exporting a binding rather than a value), resolving many classic CommonJS edge cases.

---

### 9. What is a callback function?
"A **callback function** is simply a function passed as an argument to another function, intended to be executed after the primary operation completes.

Because Node.js operations (like fetching database records) are asynchronous, the code doesn't wait for the fetch to finish. Instead, it registers a callback. Once the data is ready, Node.js fires the callback with the result.

While callbacks are the foundation of async JS, I try to avoid deeply nested callbacks—often referred to as 'Callback Hell'—by modernizing my tech stack with Promises and `async/await`."

#### Indepth
In Node.js, callbacks enforce an established convention called **Error-First Callbacks**. The first argument of the callback is always reserved for an error object (`err`), and the second argument is for the successful data (`result`). Always checking `if (err)` is critical to preventing unhandled crashes.

---

### 10. What is an event loop in Node.js?
"The **event loop** is the secret behind Node.js’s ability to handle concurrency on a single thread. 

When Node starts an async operation like reading a file, it offloads that task to the operating system or a thread pool and continues executing the rest of the script. Once the file reading is done, the system sends an event to a queue. 

The event loop constantly monitors the call stack and the message queue. If the stack is empty, it picks up the next event from the queue and executes its corresponding callback. This infinite loop is what keeps Node.js alive and listening."

#### Indepth
The event loop consists of six specific phases (Timers, Pending Callbacks, Idle/Prepare, Poll, Check, and Close Callbacks). Each phase has its own FIFO queue of callbacks to execute. Understanding these phases (especially the difference between the Poll phase and the Check phase where `setImmediate` runs) is critical for advanced performance tuning.

---

### 11. How does Node.js handle asynchronous operations?
"Node.js handles asynchronous operations using an event-driven, non-blocking I/O model based on the Event Loop and the Libuv library.

When I call an async file read, Node.js doesn't pause its single JavaScript thread to wait for the hard drive. Instead, it passes the task to Libuv in the background (which relies on OS kernels or a C++ worker thread pool) and continues executing my subsequent JavaScript code.

Once Libuv finishes the task, it pushes the callback I provided into the task queue, and the Event Loop eventually picks it up and executes it on the main thread."

#### Indepth
Libuv maintains a default pool of four worker threads (which can be scaled up via the `UV_THREADPOOL_SIZE` environment variable) to handle expensive tasks like file system I/O, DNS lookups, or crypto operations. True network I/O, however, relies almost entirely on efficient OS-level constructs like `epoll` (Linux) or `kqueue` (macOS), requiring no extra threads at all.

---

### 12. What is non-blocking I/O?
"**Non-blocking I/O** means that when an Input/Output operation happens—like a DB query or an HTTP request—the system does not halt the execution of the entire program while waiting for a response.

In a blocking system, a thread sits idle waiting for the database to return records. In Node.js's non-blocking system, the application fire-and-forgets the query, immediately moves on to serve the next user's request, and only goes back to the database query when a notification arrives saying the data is ready.

This single mechanism is why Node can handle huge numbers of simultaneous connections efficiently."

#### Indepth
Non-blocking I/O allows for massive scalability without the massive memory overhead of threaded servers. A thread in Java might consume 1MB of RAM just to be idle waiting for I/O; 10,000 idle connections would consume 10GB of RAM. Node.js handles those 10,000 connections with a minimal memory footprint.

---

### 13. What are some core modules in Node.js?
"Core modules are built directly into Node.js and don't require external installation via npm; you simply `require` or `import` them.

Some of the most important ones I use daily include:
- `http`: For creating basic web servers and making HTTP requests.
- `fs` (File System): For reading, writing, and manipulating files.
- `path`: For safely resolving and handling file/directory paths across different OSes.
- `crypto`: For generating hashes, encrypting data, and managing secrets.
- `events`: For creating custom event emitters."

#### Indepth
When importing core modules in newer versions of Node.js using ESM, it is highly recommended to prefix the import with `node:` (e.g., `import fs from 'node:fs'`). This prefix formally registers that the module is a builtin, circumventing require cache checks and clearly separating it from third-party NPM packages.

---

### 14. What is the use of the `fs` module?
"The `fs` (File System) module allows me to interact specifically with the local computer's file system.

It provides functionality to create, read, append to, delete, and rename files or directories. What I appreciate most is that `fs` provides both asynchronous methods (like `fs.readFile`) and synchronous methods (like `fs.readFileSync`). 

For production web servers, I strictly use the asynchronous methods to ensure I never block the event loop while reading large files from disk."

#### Indepth
The modern and idiomatic way to use the file system module is through `fs.promises` (or `import { readFile } from 'node:fs/promises'`). This provides immediate Promise-based APIs, removing the need for deeply nested callbacks or utilities like `util.promisify`.

---

### 15. How do you read a file asynchronously in Node.js?
"I typically use the Promise-based API for clean, modern code.

```javascript
import { readFile } from 'node:fs/promises';

async function getFileData() {
  try {
    const data = await readFile('./myfile.txt', 'utf8');
    console.log(data);
  } catch (error) {
    console.error("Error reading file:", error);
  }
}
```
If I were restricted to legacy code, I’d use `fs.readFile` with a callback. The critical part is specifying `'utf8'` as the encoding, otherwise, Node.js will return raw Buffer data instead of a readable string."

#### Indepth
If a file is exceptionally massive (e.g., a 2GB log file), reading it entirely into memory using `readFile` will crash the V8 heap engine. Under these circumstances, you should always use `fs.createReadStream()` to consume the file in tiny, memory-efficient chunks.

---

### 16. How can you handle errors in Node.js?
"Error handling depends entirely on whether the operation is synchronous or asynchronous.

For synchronous code, or when using modern `async/await`, I wrap the risky code in a standard `try...catch` block. 
For older callback-based APIs, I check the first parameter of the callback using the 'error-first' convention: `if (err) return console.log(err);`.
For Promise chains, I append a `.catch(err => ...)` clause.

Additionally, at the application level, I listen for the `process.on('unhandledRejection')` and `process.on('uncaughtException')` events to log critical crashes before they shut the server down."

#### Indepth
By default, Node.js treats Unhandled Promise Rejections very seriously. Historically, it only printed a warning, but modern Node.js versions will gracefully exit the process with a non-zero status code if left unhandled. Strong centralized error handling (like an Express error middleware pipeline) is paramount to avoiding random crashes.

---

### 17. What is the difference between `process.nextTick()` and `setImmediate()`?
"This is a classic Node.js trick question. Despite their names, they operate somewhat oppositely.

`process.nextTick()` tells Node.js to fire the callback *immediately* after the current operation finishes, pushing it to the very front of the microtask queue. It executes *before* the event loop proceeds to the next phase.
`setImmediate()` schedules a callback to run on the upcoming loop cycle in the **Check phase**, typically after I/O events have fired.

If I want to defer something to avoid blocking but allow I/O to run first, I use `setImmediate()`. If I critically need to run a small piece of code logically 'now' before any async stuff, I use `nextTick()`."

#### Indepth
A recursive or infinite loop using `process.nextTick()` will permanently freeze a Node.js process because the event loop will never leave the microtask queue to process I/O or network requests. A recursive `setImmediate()`, however, yields to the event loop each tick, allowing HTTP requests and file reading to continue without freezing.

---

### 18. What is a global object in Node.js?
"In Node.js, `global` is the top-level object analogous to the `window` object in a web browser. 

Any variable or function that you attach directly to `global.myVar` becomes accessible across every single module in your application without needing to `require` or `import` it.

There are also built-in global objects automatically provided by Node.js, such as `console`, `process`, `setTimeout()`, and `Buffer`. However, as a best practice, I heavily avoid modifying the user-defined `global` object to prevent unpredictable state and namespace collisions."

#### Indepth
With the introduction of ES2020, `globalThis` was introduced as the standardized way to reference the top-level global object regardless of the environment (it evaluates to `window` in the browser, `global` in Node.js, and `self` in Web Workers), enabling true cross-platform JavaScript module authoring.

---

### 19. What is the difference between `__dirname` and `./`?
"`__dirname` is an absolute path. It points strictly to the directory path of the exact file where it is currently written. 

`./` is a relative path. In core operations like `fs.readFile('./file.txt')`, `./` represents the directory from which the Node process was initially *launched* (the Current Working Directory, or CWD), not the directory of the script.

I strictly use `path.join(__dirname, 'file.txt')` for file operations to guarantee the app behaves predictably, regardless of which folder my terminal is in when I type `node start`."

#### Indepth
Because `__dirname` and `__filename` are variables automatically injected by CommonJS wrapper functions, they do technically **not exist** in modern ES Modules. To replicate `__dirname` in an ESM file, you must construct it using `import.meta.url` like this: `dirname(fileURLToPath(import.meta.url))`.

---

### 20. What is the `process` object in Node.js?
"The `process` object is a global object that provides incredibly useful information and control over the current Node.js execution environment.

I use it heavily for a few key workflows:
1. Accessing environment variables via `process.env` (for API keys, passwords, database URLs).
2. Controlling graceful application shutdowns with `process.exit(0)`.
3. Reading command-line arguments using `process.argv`.
4. Catching application-wide crashes via `process.on('uncaughtException')`.

It essentially acts as the bridge connecting my JavaScript code to the host operating system."

#### Indepth
`process.env` does not reflect the exact live environment block of the OS process securely; reading and writing variables inside it forces expensive C++ bridging. For extremely high-performance scenarios, avoid constantly querying `process.env` inside loops—instead, cache the target variables early at the top of your scripts during initialization.
