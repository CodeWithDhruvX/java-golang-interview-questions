## 🟢 Basics (Questions 1-20)

### Question 1: What is Node.js?

**Answer:**
Node.js is an open-source, cross-platform JavaScript runtime environment that executes JavaScript code outside a web browser. It is built on Chrome's V8 JavaScript engine.

**Key Features:**
*   **Event-driven and Non-blocking I/O:** Efficiently handles concurrent operations.
*   **Single-threaded Event Loop:** Uses a single thread to handle multiple requests asynchronously.
*   **NPM:** Comes with a powerful package ecosystem.

**Basic Server Example:**
```javascript
const http = require('http');

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello, Node.js!');
});

server.listen(3000, () => {
  console.log('Server running at http://localhost:3000/');
});
```

---

### Question 2: How is Node.js different from JavaScript in the browser?

**Answer:**
While both usage JavaScript, their environments and capabilities differ significantly.

**Comparison:**
| Feature | Node.js | Browser JavaScript |
| :--- | :--- | :--- |
| **Global Object** | `global` | `window` |
| **DOM Access** | No (No `document`, `window`) | Yes (Manipulates HTML/CSS) |
| **File System** | Yes (`fs` module) | No (Security restriction) |
| **Modules** | CommonJS (`require`) & ES Modules | ES Modules (`import`) |

**Code Difference:**
```javascript
// Node.js: Accessing environment variables
console.log(process.env.NODE_ENV);

// Browser: Accessing the DOM (Will fail in Node.js)
// document.getElementById('app'); // ReferenceError: document is not defined
```

---

### Question 3: What is the V8 engine?

**Answer:**
V8 is Google's open-source high-performance JavaScript and WebAssembly engine, written in C++. It is used in Chrome and Node.js.

**Role:**
*   **JIT Compilation:** Compiles JavaScript directly to native machine code before execution, rather than interpreting it.
*   **Garbage Collection:** Manages memory allocation and reclamation.

**Optimization Example:**
V8 uses "hidden classes" to optimize property access. Always initializing objects in the same order helps V8 optimization.

```javascript
// Good for V8 optimization
class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
}

const p1 = new Point(1, 2);
const p2 = new Point(3, 4);
```

---

### Question 4: What is the role of the `package.json` file?

**Answer:**
`package.json` is the manifest file for a Node.js project. It stores metadata and configuration.

**Key Fields:**
*   **dependencies:** Packages required for the app to run.
*   **devDependencies:** Packages needed only for development (testing, linting).
*   **scripts:** Aliases for running commands.

**Example `package.json`:**
```json
{
  "name": "my-app",
  "version": "1.0.0",
  "scripts": {
    "start": "node index.js",
    "dev": "nodemon index.js"
  },
  "dependencies": {
    "express": "^4.18.2"
  }
}
```

---

### Question 5: What is npm? How is it different from npx?

**Answer:**
*   **npm (Node Package Manager):** Installs and manages packages.
*   **npx (Node Package Execute):** Executes packages without installing them globally or locally permanently.

**Comparison:**

**npm Usage:**
```bash
# Installs create-react-app globally
npm install -g create-react-app
create-react-app my-app
```

**npx Usage:**
```bash
# Downloads and runs create-react-app temporarily
npx create-react-app my-app
```

---

### Question 6: What are modules in Node.js?

**Answer:**
Modules are reusable blocks of code (functions, objects, values) that can be exported from one file and imported into another.

**Types:**
1.  **Core Modules:** Built-in (`fs`, `http`, `path`).
2.  **Local Modules:** Created by the developer.
3.  **Third-party Modules:** Installed via npm (`express`, `lodash`).

**Example:**
`math.js`:
```javascript
const add = (a, b) => a + b;
module.exports = { add };
```

`app.js`:
```javascript
const math = require('./math');
console.log(math.add(2, 3)); // 5
```

---

### Question 7: How do you export and import modules in Node.js?

**Answer:**
Node.js supports two module systems: CommonJS (default) and ES Modules (modern).

**CommonJS (`require`/`module.exports`):**
```javascript
// export
module.exports = { hello: 'world' };

// import
const data = require('./data');
```

**ES Modules (`import`/`export`):**
(Requires `"type": "module"` in `package.json` or `.mjs` extension)
```javascript
// export
export const hello = 'world';

// import
import { hello } from './data.js';
```

---

### Question 8: What is the difference between `require()` and `import`?

**Answer:**
| Feature | `require()` | `import` |
| :--- | :--- | :--- |
| **Type** | CommonJS | ES Modules |
| **Loading** | Synchronous | Asynchronous |
| **Placement** | Anywhere in code | Top-level only |
| **Logic** | Can be dynamic | Static analysis |

**Code Example:**
```javascript
// Dynamic require (Correct)
if (condition) {
  const lib = require('./lib');
}

// Dynamic import (ESM supports import() function)
if (condition) {
  import('./lib').then(module => { /* ... */ });
}
```

---

### Question 9: What is a callback function?

**Answer:**
A callback is a function passed as an argument to another function, which is then invoked inside the outer function to complete some kind of routine or action.

**Sync vs Async Callback:**
```javascript
// Synchronous Callback (Array map)
[1, 2, 3].map(v => console.log(v));

// Asynchronous Callback (File Read)
const fs = require('fs');

fs.readFile('file.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  console.log('File data:', data);
});
console.log('Reading file...'); // Runs first
```

---

### Question 10: What is an event loop in Node.js?

**Answer:**
The Event Loop is the mechanism that allows Node.js to perform non-blocking I/O operations despite being single-threaded. It offloads operations to the system kernel whenever possible.

**Simplified Cycle:**
1.  **Timers:** Execute `setTimeout` and `setInterval` callbacks.
2.  **Poll:** Retrieve new I/O events; execute I/O related callbacks.
3.  **Check:** Execute `setImmediate` callbacks.
4.  **Close Callbacks:** Execute close events (e.g., `socket.on('close', ...)`).

**Visual Code Flow:**
```javascript
console.log('Start');

setTimeout(() => {
  console.log('Timeout');
}, 0);

setImmediate(() => {
  console.log('Immediate');
});

console.log('End');

// Output order:
// Start
// End
// Timeout (or Immediate, order varies depending on I/O cycle)
// Immediate (or Timeout)
```

---

### Question 11: How does Node.js handle asynchronous operations?

**Answer:**
Node.js uses the **Event Loop** and **Worker Threads** (internal C++ thread pool via `libuv`) to handle async tasks.

*   **Non-blocking I/O:** Network and file I/O callbacks are managed by the Event Loop.
*   **Heavy Work:** CPU-intensive tasks (crypto, compression) are offloaded to `libuv`'s thread pool to avoid blocking the main thread.

**Example:**
```javascript
const crypto = require('crypto');

// Offloaded to Thread Pool (Parallel execution)
crypto.pbkdf2('pass', 'salt', 100000, 64, 'sha512', () => {
  console.log('Crypto 1 done');
});

crypto.pbkdf2('pass', 'salt', 100000, 64, 'sha512', () => {
  console.log('Crypto 2 done');
});
```

---

### Question 12: What is non-blocking I/O?

**Answer:**
Non-blocking I/O means that the system does not get blocked (stuck) waiting for an I/O operation (like reading a database or file) to complete. Instead, it continues executing other code and runs a callback when the operation finishes.

**Blocking vs Non-blocking:**
```javascript
const fs = require('fs');

// Blocking (Sync)
try {
  const data = fs.readFileSync('file.txt'); // Halts execution here
  console.log(data);
} catch (err) {}

// Non-blocking (Async)
fs.readFile('file.txt', (err, data) => {
  console.log(data); // Runs later
});
console.log('Next task...'); // Runs immediately
```

---

### Question 13: What are some core modules in Node.js?

**Answer:**
Node.js comes with bundled modules.

**Common Core Modules:**
*   **`fs`**: File System operations.
*   **`http` / `https`**: Create servers and clients.
*   **`path`**: Handle file paths.
*   **`events`**: Handle events.
*   **`os`**: OS main info.
*   **`crypto`**: OpenSSL cryptographic functions.

**Usage:**
```javascript
const path = require('path');
const os = require('os');

console.log(path.join(__dirname, 'files'));
console.log(os.platform());
```

---

### Question 14: What is the use of the `fs` module?

**Answer:**
The `fs` module provides an API for interacting with the file system.

**Common Operations:**
*   Reading/Writing files.
*   Deleting files (`fs.unlink`).
*   Creating directories (`fs.mkdir`).

**Example:**
```javascript
const fs = require('fs');

// Write
fs.writeFile('log.txt', 'Hello Logs', (err) => {
  if (err) throw err;
  console.log('File saved!');
  
  // Read
  fs.readFile('log.txt', 'utf8', (err, data) => {
    console.log('Content:', data);
  });
});
```

---

### Question 15: How do you read a file asynchronously in Node.js?

**Answer:**
Using `fs.readFile()` or the promise-based `fs.promises.readFile()`.

**Modern Approach (Async/Await):**
```javascript
const fs = require('fs').promises;

async function readMyFile() {
  try {
    const data = await fs.readFile('data.txt', 'utf-8');
    console.log(data);
  } catch (err) {
    console.error('Error reading file:', err);
  }
}

readMyFile();
```

---

### Question 16: How can you handle errors in Node.js?

**Answer:**
Error handling strategies depend on the execution model.

**1. Sync Code (try/catch):**
```javascript
try {
  throw new Error("Something broke");
} catch (e) {
  console.error(e.message);
}
```

**2. Callbacks (Error-first argument):**
```javascript
// err is always the first argument
fs.readFile('file.txt', (err, data) => {
  if (err) return console.error("Read failed");
  // process data
});
```

**3. Promises/Async-Await:**
```javascript
myPromise()
  .then(data => console.log(data))
  .catch(err => console.error(err));
```

**4. Global Uncaught Exceptions (Last resort):**
```javascript
process.on('uncaughtException', (err) => {
  console.error('Crash averted:', err);
});
```

---

### Question 17: What is the difference between `process.nextTick()` and `setImmediate()`?

**Answer:**
*   **`process.nextTick()`**: Fires immediately after the current operation completes, **before** the event loop continues. It has higher priority.
*   **`setImmediate()`**: Fires on the next iteration of the event loop (Check phase).

**Order of Execution:**
```javascript
setImmediate(() => console.log('setImmediate'));
process.nextTick(() => console.log('process.nextTick'));
console.log('Main Program');

// Output:
// Main Program
// process.nextTick
// setImmediate
```

---

### Question 18: What is a global object in Node.js?

**Answer:**
Globals are objects available in all modules without importing. The top-level scope is `global` (unlike `window` in usage browsers).

**Common Globals:**
*   `process`
*   `console`
*   `module`, `exports`, `require` (in CommonJS)
*   `__dirname`, `__filename`
*   `setTimeout`, `setInterval`

**Example:**
```javascript
global.myGlobalVar = 10;
console.log(global.myGlobalVar); // 10
```

---

### Question 19: What is the difference between `__dirname` and `./`?

**Answer:**
*   **`__dirname`**: The absolute path to the directory containing the *current file*.
*   **`./`**: Relative path from where the script was *executed* (cwd), usually used in `require()`.

**Behavior:**
```javascript
const fs = require('fs');
const path = require('path');

// Safe way (uses absolute path)
fs.readFile(path.join(__dirname, 'config.json'), (err, data) => { /*...*/ });

// Unsafe way (depends on where you run 'node' command from)
fs.readFile('./config.json', (err, data) => { /*...*/ });
```

---

### Question 20: What is the `process` object in Node.js?

**Answer:**
The `process` object provides information about, and control over, the current Node.js process. It is a global object.

**Uses:**
*   Environment variables: `process.env`
*   Command line args: `process.argv`
*   Memory usage: `process.memoryUsage()`
*   Exit process: `process.exit()`

**Example:**
```javascript
if (process.env.NODE_ENV === 'production') {
  console.log('Running in production');
}

console.log('Arguments:', process.argv.slice(2));
// node app.js --user=admin -> ['--user=admin']
```

---

## 🟢 Intermediate Level (Questions 21-40)

### Question 21: What is the difference between synchronous and asynchronous functions?

**Answer:**
*   **Synchronous:** Blocks execution until completion. Uses the return statement to pass back data.
*   **Asynchronous:** Returns immediately (typically `undefined` or a Promise). Uses callbacks or promises to return data later.

**Code:**
```javascript
// Sync
const result = doBuilt(data);
console.log(result);

// Async
doBuiltAsync(data, (result) => {
  console.log(result);
});
```

---

### Question 22: What is a Promise in Node.js?

**Answer:**
A Promise represents the eventual completion (or failure) of an asynchronous operation and its resulting value.

**States:** `Pending`, `Fulfilled` (Resolved), `Rejected`.

**Creating a Promise:**
```javascript
const myPromise = new Promise((resolve, reject) => {
  const success = true;
  if (success) {
    resolve('Operation Successful');
  } else {
    reject('Operation Failed');
  }
});

myPromise
  .then(res => console.log(res))
  .catch(err => console.error(err));
```

---

### Question 23: How do you convert a callback to a promise?

**Answer:**
You can wrap the callback function in a new Promise or use `util.promisify`.

**Method 1: Wrapping**
```javascript
const fs = require('fs');

function readFilePromise(path) {
  return new Promise((resolve, reject) => {
    fs.readFile(path, 'utf8', (err, data) => {
      if (err) reject(err);
      else resolve(data);
    });
  });
}
```

**Method 2: `util.promisify`**
```javascript
const util = require('util');
const fs = require('fs');

const readFile = util.promisify(fs.readFile);

readFile('file.txt', 'utf8')
  .then(data => console.log(data))
  .catch(err => console.error(err));
```

---

### Question 24: What are async/await keywords? How do they work?

**Answer:**
`async` and `await` are syntactic sugar built on top of Promises to make asynchronous code look and behave like synchronous code.

*   `async`: Declares that a function returns a Promise.
*   `await`: Pauses execution of the `async` function until the Promise resolves.

**Example:**
```javascript
async function getData() {
  try {
    const data = await fetchApi(); // Waits here
    console.log(data);
  } catch (error) {
    console.error("Fetch failed", error);
  }
}
```

---

### Question 25: What are streams in Node.js?

**Answer:**
Streams are objects that create a continuous flow of data. They let you read data from a source or write data to a destination in chunks, rather than loading the entire dataset into memory.

**Benefits:** Memory efficiency (processing massive files) and Time efficiency (start processing before full read).

**Example:**
```javascript
const fs = require('fs');
// Reads file in chunks instead of memory dump
const stream = fs.createReadStream('huge-file.log');

stream.on('data', (chunk) => {
  console.log(`Received ${chunk.length} bytes`);
});
```

---

### Question 26: Explain the four types of streams in Node.js.

**Answer:**
1.  **Readable:** Source of data (e.g., `fs.createReadStream`, `req` in HTTP).
2.  **Writable:** Destination for data (e.g., `fs.createWriteStream`, `res` in HTTP).
3.  **Duplex:** Both Readable and Writable (e.g., TCP Sockets).
4.  **Transform:** Duplex stream that modifies data as it passes (e.g., `zlib.createGzip` for compression).

**Pipe Example (Readable to Writable):**
```javascript
const fs = require('fs');
const readable = fs.createReadStream('input.txt');
const writable = fs.createWriteStream('output.txt');

readable.pipe(writable); // Efficient copy
```

---

### Question 27: What is backpressure in streams?

**Answer:**
Backpressure occurs when the **Readable stream** produces data faster than the **Writable stream** can consume it. This can cause memory buffers to fill up.

**Handling:** `pipe` automatically manages backpressure by pausing the readable stream when the writable buffer is full.

**Manual Handling:**
```javascript
if (!writable.write(chunk)) {
  readable.pause(); // High watermark reached
}

writable.on('drain', () => {
  readable.resume(); // Buffer drained, resume reading
});
```

---

### Question 28: How can you create a custom readable stream?

**Answer:**
Extend the `Readable` class and implement the `_read` method.

**Code:**
```javascript
const { Readable } = require('stream');

class MyStream extends Readable {
  constructor(opt) {
    super(opt);
    this.max = 5;
    this.current = 0;
  }

  _read() {
    this.current += 1;
    if (this.current > this.max) {
      this.push(null); // Signal EOF
    } else {
      this.push(String(this.current)); // Push data
    }
  }
}

const stream = new MyStream();
stream.pipe(process.stdout); // Outputs: 12345
```

---

### Question 29: How does error handling differ between callbacks and promises?

**Answer:**
*   **Callbacks:** You must manually check `if (err)` in every callback. Errors in callbacks don't automatically bubble up.
*   **Promises:** Errors propagate down the chain to the nearest `.catch()`.

**Comparison:**
```javascript
// Callback Hell Error Handling
func1((err, res) => {
  if (err) return handleError(err);
  func2(res, (err, res2) => {
    if (err) return handleError(err);
    // ...
  });
});

// Promise Chain
func1()
  .then(func2)
  .then(func3)
  .catch(handleError); // Catches error from any step
```

---

### Question 30: What is middleware in Express.js?

**Answer:**
Middleware functions are functions that have access to the request object (`req`), the response object (`res`), and the next middleware function (`next`) in the application’s request-response cycle.

**Uses:** Logging, Authentication, parsing body, error handling.

**Example:**
```javascript
const express = require('express');
const app = express();

// Middleware Function
const myLogger = function (req, res, next) {
  console.log('LOGGED');
  next(); // Pass control to next handler
};

app.use(myLogger);

app.get('/', (req, res) => {
  res.send('Hello World!');
});
```

---

### Question 31: What is the role of the `next()` function in Express middleware?

**Answer:**
`next()` passes control to the next middleware in the stack. If `next()` is not called and the request is not terminated (using `res.send`), the request will hang.

**Example:**
```javascript
app.use((req, res, next) => {
  if (req.query.auth === 'true') {
    next(); // Authorized, proceed
  } else {
    res.status(403).send('Forbidden'); // Stop here
  }
});
```

---

### Question 32: How can you handle 404 errors in Express?

**Answer:**
Add a middleware function at the very end of your middleware stack (after all routes). If execution reaches it, no route matched.

**Code:**
```javascript
// ... defined routes ...

// 404 Handler
app.use((req, res, next) => {
  res.status(404).send("Sorry caused't find that!");
});

app.listen(3000);
```

---

### Question 33: How do you configure environment variables in Node.js?

**Answer:**
Environment variables are passed via `process.env`. They are usually managed using a `.env` file and the `dotenv` library during development to keep secrets out of code.

**Accessing:**
```javascript
const port = process.env.PORT || 3000;
console.log(`Server starting on ${port}`);
```

**Running:**
```bash
PORT=4000 node app.js
```

---

### Question 34: What is `dotenv` and how is it used?

**Answer:**
`dotenv` is a zero-dependency module that loads environment variables from a `.env` file into `process.env`.

**Process:**
1.  Create `.env` file: `API_KEY=12345`
2.  Install: `npm install dotenv`
3.  Configure at app start.

**Code:**
```javascript
require('dotenv').config();

console.log(process.env.API_KEY); // 12345
```

---

### Question 35: How do you handle file uploads in Node.js?

**Answer:**
Native logic is complex (parsing multipart streams). Usually, libraries like `multer` (for Express) are used.

**Using Multer:**
```javascript
const express = require('express');
const multer  = require('multer');
const upload = multer({ dest: 'uploads/' });
const app = express();

app.post('/profile', upload.single('avatar'), (req, res, next) => {
  // req.file is the `avatar` file
  // req.body will hold the text fields, if any
  res.send('File uploaded successfully');
});
```

---

### Question 36: How can you serve static files using Express?

**Answer:**
Use the built-in middleware `express.static`.

**Code:**
```javascript
const express = require('express');
const app = express();

// Serve files from the "public" directory
// accessible via http://localhost:3000/images/logo.png
app.use(express.static('public'));

app.listen(3000);
```

---

### Question 37: What is the `cluster` module?

**Answer:**
Node.js is single-threaded. The `cluster` module allows you to easily create child processes (workers) that all share server ports, utilizing multi-core systems.

**Code:**
```javascript
const cluster = require('cluster');
const http = require('http');
const numCPUs = require('os').cpus().length;

if (cluster.isMaster) {
  console.log(`Master ${process.pid} is running`);
  // Fork workers
  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }
} else {
  // Workers can share any TCP connection
  http.createServer((req, res) => {
    res.writeHead(200);
    res.end('Hello World\n');
  }).listen(8000);
  console.log(`Worker ${process.pid} started`);
}
```

---

### Question 38: What is the `child_process` module used for?

**Answer:**
It enables spinning up child processes to execute external OS commands or scripts.

**Key functions:**
*   `exec()`: Buffer output (good for short output).
*   `spawn()`: Stream output (good for large data).
*   `fork()`: Like `spawn` but for Node modules with communication channel.

**Example:**
```javascript
const { exec } = require('child_process');

exec('ls -lh', (error, stdout, stderr) => {
  if (error) return console.error(`error: ${error.message}`);
  if (stderr) return console.error(`stderr: ${stderr}`);
  console.log(`stdout: ${stdout}`);
});
```

---

### Question 39: What is the difference between `spawn` and `exec`?

**Answer:**
*   **`spawn()`**: Returns a stream. It's meant for large amounts of data transfer. It starts a new process.
*   **`exec()`**: Buffers the output. It returns the whole output at once in a callback. It spawns a shell.

**When to use:**
*   Use `exec` for small output commands (e.g., `git status`).
*   Use `spawn` for data-heavy processes (e.g., image processing, reading large binary).

---

### Question 40: How can you handle CORS in a Node.js application?

**Answer:**
CORS (Cross-Origin Resource Sharing) restricts browser requests from other domains. In Express, use the `cors` package.

**Manual Headers:**
```javascript
app.use((req, res, next) => {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Content-Type");
  next();
});
```

**Using Package:**
```javascript
const cors = require('cors');
app.use(cors()); // Allow all
```

---

## 🟢 Advanced Level (Questions 41-50)

### Question 41: Explain the concept of event-driven programming.

**Answer:**
Event-driven programming is a paradigm where the flow of the program is determined by events (user actions, sensor outputs, or messages from other programs). Node.js uses this to handle asynchronous actions.

**Mechanism:**
1.  **Emit Event:** Something happens.
2.  **Event Listener:** A function waits for that event.
3.  **Callback:** The function executes when the event occurs.

**Example:**
```javascript
const EventEmitter = require('events');
const eventEmitter = new EventEmitter();

// Listener
eventEmitter.on('scream', () => {
  console.log('I hear a scream!');
});

// Emitter
eventEmitter.emit('scream');
```

---

### Question 42: What is the libuv library?

**Answer:**
`libuv` is a multi-platform C library that provides support for asynchronous I/O based on event loops. It powers the Node.js event loop and handles filesystem, DNS, network, child processes, pipes, signal handling, and polling.

**Key Role:**
It provides the **Thread Pool** to handle "expensive" tasks (File I/O, Compression, Crypto) that would otherwise block the main V8 thread.

---

### Question 43: What are the phases of the event loop?

**Answer:**
The Node.js event loop has specific phases:
1.  **Timers:** Execute `setTimeout()` and `setInterval()` callbacks.
2.  **Pending Callbacks:** System-related callbacks (e.g., TCP errors).
3.  **Idle, Prepare:** Internal Node.js use.
4.  **Poll:** Retrieve new I/O events; execute I/O related callbacks. This is where most logic happens.
5.  **Check:** Execute `setImmediate()` callbacks.
6.  **Close Callbacks:** e.g., `socket.on('close')`.

---

### Question 44: What is the microtask queue vs. the macrotask queue?

**Answer:**
*   **Macrotasks:** Standard callbacks (setTimeout, setInterval, I/O). Executed one per loop tick.
*   **Microtasks:** Higher priority (Promises, `process.nextTick`). The **entire** microtask queue is processed after *every* macrotask and before the next phase of the event loop.

**Priority Order:**
1.  `process.nextTick` (technically its own queue, highest priority).
2.  `Promise.then` (Microtask).
3.  `setTimeout` / `setImmediate` (Macrotask).

---

### Question 45: How does Node.js handle concurrency?

**Answer:**
Node.js is single-threaded but handles concurrency via the **Event Loop** for I/O operations and the **Worker Pool** (libuv) for CPU-intensive tasks.

**Visual:**
1.  Client sends Request A.
2.  Node sees it's I/O, offloads to OS, goes to Request B.
3.  Client sends Request B.
4.  OS finishes Request A, triggers callback in Event Loop.
5.  Node processes Callback A.

Use `Worker Threads` or `Cluster` module for CPU-bound concurrency (parallelism).

---

### Question 46: What are worker threads?

**Answer:**
The `worker_threads` module enables the use of threads that execute JavaScript in parallel. While Node is single-threaded, Worker Threads allow CPU-heavy tasks to run efficiently without blocking the event loop.

**Code:**
```javascript
const { Worker, isMainThread, parentPort } = require('worker_threads');

if (isMainThread) {
  // Main thread spawns a worker
  const worker = new Worker(__filename);
  worker.on('message', msg => console.log('Received:', msg));
} else {
  // Worker thread code
  parentPort.postMessage('Hello from Worker!');
}
```

---

### Question 47: What is the difference between `fork()` and `spawn()`?

**Answer:**
Both create child processes, but `fork()` is specialized for Node.js modules.

*   **`spawn()`**: Launches a command (any binary). Streams data. No communication channel by default.
*   **`fork()`**: Special instance of `spawn` that runs a specific Node module. It establishes a built-in IPC (Inter-Process Communication) channel that allows sending messages between parent and child via `send()` and `on('message')`.

**Fork Example:**
```javascript
// parent.js
const { fork } = require('child_process');
const child = fork('child.js');
child.send({ hello: 'world' });
```

---

### Question 48: What is memory leak and how do you find it in Node.js?

**Answer:**
A memory leak occurs when objects are no longer needed but are still referenced, preventing the Garbage Collector from freeing them.

**Common Causes:** Global variables, uncleared timers, closures holding large scope, detached DOM elements (less relevant in Node).

**Debugging Tools:**
1.  **Node Inspector:** `node --inspect index.js` (Open chrome://inspect).
2.  **Heap Snapshots:** Compare memory usage over time.
3.  **`process.memoryUsage()`:** Monitor `heapUsed`.

---

### Question 49: How can you improve the performance of a Node.js app?

**Answer:**
*   **Use Asynchronous I/O:** Never block the main thread.
*   **Caching:** Use Redis to cache frequent DB queries.
*   **Clustering:** Use PM2 or `cluster` to utilize all CPU cores.
*   **Gzip Compression:** Compress responses.
*   **Optimize Queries:** proper DB indexing.
*   **Memory Management:** Monitor leaks, handle streams properly (pipes).

**Gzip with Express:**
```javascript
const compression = require('compression');
app.use(compression());
```

---

### Question 50: How do you scale a Node.js application?

**Answer:**
1.  **Vertical Scaling (Scale Up):** Increase the resources (RAM, CPU) of a single machine.
2.  **Horizontal Scaling (Scale Out):** Add more machines/instances.
3.  **Clustering:** Run multiple instances on one machine (one per core) using `PM2`.
4.  **Microservices:** Break app into smaller services deployed independently.

**PM2 Scaling:**
```bash
pm2 start app.js -i max
# -i max: Instances = Number of CPU cores
```
