## 🟢 Core Node.js & Concepts (Questions 101-150)

### Question 101: Why is Node.js single-threaded?

**Answer:**
Node.js uses a single-threaded event loop model architecture (primarily inspired by JavaScript in the browser) to handle concurrency.

**Reasoning:**
*   **Simplicity:** No need to manage complex thread synchronization, locks, or deadlocks.
*   **Efficiency:** Uses non-blocking I/O to handle thousands of concurrent connections with low memory overhead compared to "one thread per request" models (like Apache/PHP).

**Note:** Only the *JavaScript execution* is single-threaded. System I/O and crypto tasks run on parallel threads in the C++ `libuv` thread pool.

---

### Question 102: Can Node.js be used for CPU-intensive tasks?

**Answer:**
Historically, no, because heavy computation blocks the single event loop, stopping all other requests.

**Modern Solution:**
Yes, by using **Worker Threads** or **Child Processes** to offload the heavy calculation to a separate thread/process, keeping the main loop free.

**Blocking Example (Avoid):**
```javascript
app.get('/compute', (req, res) => {
  // Freezes the server for seconds
  let sum = 0;
  for (let i = 0; i < 1e9; i++) sum += i; 
  res.send({ sum });
});
```

**Non-Blocking Example (Worker):**
```javascript
const { Worker } = require('worker_threads');
// Spawn worker to do the loop
```

---

### Question 103: How does Node.js differ from traditional multithreaded servers like Apache?

**Answer:**
*   **Apache (Thread-based):** Creates a new thread/process for every incoming request. High memory usage per connection. If threads run out, new users wait.
*   **Node.js (Event-driven):** Single thread handles all requests. I/O operations (DB, File) are delegated. When I/O finishes, a callback fires. Low memory usage.

---

### Question 104: What are the benefits of using Node.js?

**Answer:**
1.  **Fast Processing:** V8 compiles JS to machine code.
2.  **Scalability:** Event loop handles massive concurrent I/O.
3.  **Unified Stack:** JavaScript on Frontend and Backend.
4.  **Rich Ecosystem:** NPM has millions of packages.
5.  **JSON Native:** Perfect for REST APIs and NoSQL (MongoDB).

---

### Question 105: What is event delegation in Node.js?

**Answer:**
Event delegation is a design pattern where a single listener manages events for multiple sources. While more common in DOM (Browser), in Node.js, it translates to using a central `EventEmitter` to handle events from distinct parts of the app.

---

### Question 106: How does the `require` cache work?

**Answer:**
When you `require('./module')` for the first time, Node.js executes the file and caches the `module.exports` object. Subsequent calls to `require('./module')` return the **cached object** without re-executing the file.

**Implication:**
This acts like a Singleton. If you want a new instance every time, export a function/class instead of an object.

**Code:**
```javascript
// Data loaded only once
const data = require('./large-data.json'); 
```

---

### Question 107: What is the difference between `require()` and `import` in terms of execution time?

**Answer:**
*   **`require()`**: Dynamic and synchronous. It is executed at runtime. You can conditionally require files.
*   **`import`**: Static and asynchronous setup. The structure is analyzed during the parsing phase (before execution). It allows for static analysis and tree-shaking.

---

### Question 108: Can you overwrite a global variable in Node.js?

**Answer:**
Yes, but it is considered a bad practice (globals pollution).

**Example:**
```javascript
global.console = {
  log: () => { /* Silence logs */ }
};

console.log("Hello"); // Does nothing
```

---

### Question 109: What is a REPL in Node.js?

**Answer:**
REPL stands for **Read-Eval-Print Loop**. It is the interactive shell you get when you run `node` in the terminal without a filename. Use it for testing snippets.

*   **R**ead: Reads user input.
*   **E**val: Evaluates the code.
*   **P**rint: Print result.
*   **L**oop: Loops back for next input.

---

### Question 110: What are timers in Node.js and how do they work?

**Answer:**
Timers schedule code execution for the future. They are backed by `libuv`.
*   `setTimeout(cb, ms)`: Run after X ms.
*   `setInterval(cb, ms)`: Run repeats every X ms.
*   `setImmediate(cb)`: Run after I/O poll phase.
*   `process.nextTick(cb)`: Run immediately after current operation.

---

### Question 111: How do you watch a file for changes in Node.js?

**Answer:**
Use `fs.watch()` or `fs.watchFile()`.

**Code:**
```javascript
const fs = require('fs');

fs.watch('target.txt', (eventType, filename) => {
  console.log(`Event type is: ${eventType}`);
  if (filename) {
    console.log(`Filename provided: ${filename}`);
  }
});
```
*Note: `chokidar` package is often preferred for better cross-platform consistency.*

---

### Question 112: What’s the difference between `fs.readFile()` and `fs.createReadStream()`?

**Answer:**
*   **`fs.readFile()`**: Loads **entire** file into memory (Buffer) before callback. Crash risk for large files.
*   **`fs.createReadStream()`**: Reads file in **chunks**. Efficient memory usage.

**Rule:** Use streams for files > 10MB or when serving web content.

---

### Question 113: What are buffers, and how are they different from arrays?

**Answer:**
*   **Buffer:** Raw memory allocation outside the V8 heap. Stores binary data (integers 0-255). Fixed size.
*   **Array:** JS Objects on V8 heap. Dynamic size. Can hold any type.

**Creation:**
```javascript
const buf = Buffer.alloc(10); // 10 bytes of zeros
const bufStr = Buffer.from("Hello");
```

---

### Question 114: How do you convert a buffer to a string?

**Answer:**
Use the `.toString()` method on the buffer.

**Code:**
```javascript
const buf = Buffer.from('Node.js');
console.log(buf.toString()); // 'Node.js'
console.log(buf.toString('base64')); // 'Tm9kZS5qcw=='
```

---

### Question 115: What is the default encoding for buffers?

**Answer:**
Buffers do not have an encoding (they are raw bytes), but when converting to/from strings, the default is **UTF-8**.

---

### Question 116: What are the phases of the Node.js event loop?

**Answer:**
(Recap of Q43)
1.  **Timers**
2.  **Pending Callbacks**
3.  **Idle, Prepare**
4.  **Poll** (Main I/O)
5.  **Check** (`setImmediate`)
6.  **Close Callbacks**

---

### Question 117: In what order are `setTimeout`, `setImmediate`, and `process.nextTick` executed?

**Answer:**
Order in a standard I/O cycle:
1.  `process.nextTick` (Highest priority, runs before loop continues).
2.  `setTimeout` (If timer expired).
3.  `setImmediate` (Runs in Check phase, after I/O).

**Note:** Inside an I/O callback (like fs.readFile), `setImmediate` **always** runs before `setTimeout(..., 0)`.

---

### Question 118: What is the role of the `Timers` phase in the event loop?

**Answer:**
The Timers phase executes callbacks scheduled by `setTimeout()` and `setInterval()`. Node checks if the threshold time has passed.

---

### Question 119: When should you use `setImmediate()` over `setTimeout()`?

**Answer:**
Use `setImmediate()` when you want to execute a callback **after the current I/O cycle** but before timers.
*   It ensures the function runs after I/O events are polled.
*   Typical use: Spreading out computationally expensive operations to allow the loop to handle I/O in between.

---

### Question 120: Can you starve the event loop? How?

**Answer:**
Yes. If you fill the queue with high-priority tasks, or execute a blocking synchronous loop, I/O will never stop.

**Starvation via microtasks:**
```javascript
function recursive() {
  process.nextTick(recursive); // Infinite loop of microtasks
}
recursive();
// Event Loop is blocked. Timers and I/O will never run.
```

---

### Question 121: What is the difference between flowing and paused mode in streams?

**Answer:**
*   **Flowing:** Data is read automatically and provided to the app via events (`data`).
*   **Paused:** You must explicitly call `stream.read()` to get data chunks.

**Switching:** Adding a `'data'` handler switches a stream to flowing mode.

---

### Question 122: How do you handle stream errors?

**Answer:**
Streams emit an `'error'` event. If you don't listen for it, the Node.js process will crash (throw uncaught error).

**Code:**
```javascript
const stream = fs.createReadStream('bad-file.txt');

stream.on('error', (err) => {
  console.error('Stream error:', err.message);
});
```

---

### Question 123: What is the `pipeline()` method used for?

**Answer:**
`stream.pipeline()` is a utility to pipe streams together and properly handle errors and cleanup (unlike simple `.pipe()`).

**Code:**
```javascript
const { pipeline } = require('stream');
const fs = require('fs');
const zlib = require('zlib');

pipeline(
  fs.createReadStream('input.txt'),
  zlib.createGzip(),
  fs.createWriteStream('input.txt.gz'),
  (err) => {
    if (err) console.error('Pipeline failed', err);
    else console.log('Pipeline succeeded');
  }
);
```

---

### Question 124: How do you implement a duplex stream?

**Answer:**
Implement both `_read` and `_write` methods.

**Example:**
A TCP socket is a Duplex stream (you can read from it and write to it).

```javascript
const { Duplex } = require('stream');

class MyDuplex extends Duplex {
  _read(size) { this.push('data'); this.push(null); }
  _write(chunk, encoding, callback) { console.log(chunk.toString()); callback(); }
}
```

---

### Question 125: What are the key events emitted by streams?

**Answer:**
*   **Readable:** `data`, `end`, `error`, `close`.
*   **Writable:** `drain`, `finish`, `error`, `close`.

---

### Question 126: How do you create an HTTP server without Express?

**Answer:**
Using the native `http` module.

**Code:**
```javascript
const http = require('http');

const server = http.createServer((req, res) => {
  if (req.url === '/' && req.method === 'GET') {
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.end('Home Page');
  } else {
    res.writeHead(404);
    res.end('Not Found');
  }
});

server.listen(3000);
```

---

### Question 127: What is the difference between `http.createServer` and `https.createServer`?

**Answer:**
`https.createServer` requires SSL/TLS certificates (private key and certificate) to encrypt traffic.

**HTTPS Code:**
```javascript
const https = require('https');
const fs = require('fs');

const options = {
  key: fs.readFileSync('key.pem'),
  cert: fs.readFileSync('cert.pem')
};

https.createServer(options, (req, res) => {
  res.writeHead(200);
  res.end('Secure Hello');
}).listen(443);
```

---

### Question 128: How do you handle file downloads in Node.js?

**Answer:**
Set appropriate headers (`Content-Disposition`) and pipe the file stream to the response.

**Code (Express):**
```javascript
app.get('/download', (req, res) => {
  const file = `${__dirname}/report.pdf`;
  // Helper method:
  res.download(file); 
  
  // Or manual stream:
  // res.setHeader('Content-Disposition', 'attachment; filename=report.pdf');
  // fs.createReadStream(file).pipe(res);
});
```

---

### Question 129: How do you implement keep-alive in HTTP requests?

**Answer:**
Node.js uses `http.Agent` to manage connection persistence (Keep-Alive).

**Client Side:**
```javascript
const http = require('http');
const agent = new http.Agent({ keepAlive: true });

http.get({ hostname: 'localhost', port: 3000, agent: agent }, (res) => {
  // connection stays open for reuse
});
```

---

### Question 130: What is connection pooling?

**Answer:**
Connection pooling is a cache of database connections maintained so that connections can be reused when future requests to the database are required. This avoids the overhead of establishing a new TCP connection for every query.

**Example (pg pool):**
```javascript
const { Pool } = require('pg');
const pool = new Pool({ max: 20 }); // pool size
// pool.query() reuses connections
```

---

### Question 131: What is the role of the `app.all()` method?

**Answer:**
`app.all()` matches **all** HTTP methods (GET, POST, PUT, DELETE, etc.) for a specific path.

**Use Case:**
Global logic for a specific section of the site (e.g., Auth check for all `/api/*` routes).

```javascript
app.all('/secret/*', (req, res, next) => {
  console.log('Secret section accessed via ' + req.method);
  next();
});
```

---

### Question 132: How can you handle different HTTP methods in the same route?

**Answer:**
Using `app.route()`.

**Code:**
```javascript
app.route('/book')
  .get((req, res) => {
    res.send('Get a random book');
  })
  .post((req, res) => {
    res.send('Add a book');
  })
  .put((req, res) => {
    res.send('Update the book');
  });
```

---

### Question 133: How do you organize routes in Express using routers?

**Answer:**
Use `express.Router` to create modular, mountable route handlers.

**users.js:**
```javascript
const express = require('express');
const router = express.Router();

router.get('/', (req, res) => res.send('User List'));
router.get('/:id', (req, res) => res.send('User Detail'));

module.exports = router;
```

**app.js:**
```javascript
const usersRouter = require('./routes/users');
app.use('/users', usersRouter);
// Handles /users and /users/:id
```

---

### Question 134: What is the order of middleware execution?

**Answer:**
Middleware is executed sequentially in the order it is defined (`app.use`).

**Scenarios:**
1.  **Global Middleware:** Defined top-level.
2.  **Route Middleware:** Defined on specific routes.
3.  **Error Middleware:** Defined last (`err, req, res, next`).

---

### Question 135: How do you create a custom middleware?

**Answer:**
A function strictly with `req, res, next` signature.

**Code:**
```javascript
const requestTime = function (req, res, next) {
  req.requestTime = Date.now();
  console.log('Time:', req.requestTime);
  next(); // Essential!
};

app.use(requestTime);
```

---

### Question 136: What is CSRF and how do you prevent it in Node.js?

**Answer:**
CSRF (Cross-Site Request Forgery) attacks force a user to execute unwanted actions on an app they are logged into.

**Prevention:** Use CSRF Tokens (Synchronizer Token Pattern).
Libraries: `csurf` (deprecated but common) or `csrf-csrf`.

**Logic:** Server sends a token. Client must submit that token with form (POST). Server verifies token match.

---

### Question 137: How do you hash passwords in Node.js?

**Answer:**
Do **not** use plain crypto/md5. Use `bcrypt` or `argon2` which are slow hashing algorithms designed to resist brute force.

**Bcrypt Example:**
```javascript
const bcrypt = require('bcrypt');
const saltRounds = 10;

// Hash
const hash = await bcrypt.hash('mypassword', saltRounds);

// Verify
const match = await bcrypt.compare('mypassword', hash);
```

---

### Question 138: What is OAuth2 and how can you implement it?

**Answer:**
OAuth2 is an authorization framework that enables apps to obtain limited access to user accounts (Google, Facebook) without exposing passwords.

**Passport.js** is the standard way to implement this in Node.js.

```javascript
const GoogleStrategy = require('passport-google-oauth20').Strategy;

passport.use(new GoogleStrategy({
    clientID: GOOGLE_CLIENT_ID,
    clientSecret: GOOGLE_CLIENT_SECRET,
    callbackURL: "/auth/google/callback"
  },
  function(accessToken, refreshToken, profile, cb) {
    User.findOrCreate({ googleId: profile.id }, function (err, user) {
      return cb(err, user);
    });
  }
));
```

---

### Question 139: What is Helmet.js?

**Answer:**
Helmet is a collection of 14+ smaller middleware functions that set security-related HTTP headers to protect against well-known web vulnerabilities (XSS, Clickjacking, Sniffing).

**Usage:**
```javascript
const helmet = require('helmet');
app.use(helmet());
```

---

### Question 140: How can you restrict API access by IP address?

**Answer:**
Middleware check against `req.ip`.

**Code:**
```javascript
const allowList = ['123.45.67.89', '::1'];

app.use((req, res, next) => {
  if (allowList.includes(req.ip)) {
    next();
  } else {
    res.status(403).send('Access Denied');
  }
});
```

---

### Question 141: What is lazy loading in Node.js?

**Answer:**
Lazy loading defers the initialization (requiring) of a module until it is actually needed. This reduces startup time and memory footprint.

**Code:**
```javascript
// Instead of top-level require:
// const heavyLib = require('heavy-lib'); 

app.get('/analyze', (req, res) => {
  // Require only when route is hit
  const heavyLib = require('heavy-lib'); 
  heavyLib.process();
  res.send('Done');
});
```

---

### Question 142: How do you implement memoization in Node.js?

**Answer:**
Memoization caches the result of expensive function calls based on inputs.

**Code:**
```javascript
const cache = {};

function expensiveCalc(num) {
  if (cache[num]) {
    console.log('Fetching from cache...');
    return cache[num];
  }
  
  console.log('Calculating...');
  const result = num * 2; // imagine heavy math
  cache[num] = result;
  return result;
}
```

---

### Question 143: What is the role of caching in Node.js apps?

**Answer:**
Caching stores frequently accessed data in fast memory (RAM) to avoid slow database or API queries.
**Tools:** Redis, Memcached, Node-Cache (in-memory).

---

### Question 144: How do you use Redis with Node.js?

**Answer:**
Use the `redis` client.

**Code:**
```javascript
const redis = require('redis');
const client = redis.createClient();

await client.connect();

// Set
await client.set('key', 'value');

// Get
const value = await client.get('key');
console.log(value);
```

---

### Question 145: How can you avoid blocking the event loop?

**Answer:**
1.  **Don't** use synchronous "Sync" versions of FS/Crypto APIs (`fs.readFileSync`) in hot paths.
2.  **Don't** perform heavy JSON parsing (`JSON.parse`) on huge strings.
3.  **Don't** run complex regex on user input (ReDoS).
4.  **Do** partition calculations using `setImmediate` or use Worker Threads.

---

### Question 146: How do you test middleware in Express?

**Answer:**
Middleware is just a function. You can unit test it by mocking `req`, `res`, and `next`.

**Test Code:**
```javascript
const authMiddleware = require('./auth');
const sinon = require('sinon');

it('should call next if authenticated', () => {
  const req = { user: { id: 1 } };
  const res = {};
  const next = sinon.spy();

  authMiddleware(req, res, next);

  expect(next.calledOnce).toBe(true);
});
```

---

### Question 147: How do you mock database calls in tests?

**Answer:**
Use libraries like `sinon`, `jest.mock`, or `proxyquire` to replace the Database module with a mock that returns fixed data.

```javascript
// mocking Mongoose model
jest.mock('./models/User');
const User = require('./models/User');

User.find.mockResolvedValue([{ name: 'John' }]);
// Now testing controller won't hit DB
```

---

### Question 148: What is code coverage and how do you measure it?

**Answer:**
Code coverage measures what percentage of your code is executed during tests.
**Tools:** `nyc` (Istanbul) or Jest's built-in coverage.

**Command:**
`jest --coverage`

---

### Question 149: How do you perform integration testing in Node.js?

**Answer:**
Integration tests interact with real components (DB, API) usually using a separate test database.
Use `Supertest` to hit endpoints, which then trigger controllers/models/db.

---

### Question 150: What is TDD and how does it apply to Node.js?

**Answer:**
**Test Driven Development**: Write specific tests **before** writing the code.
1.  Write a failing test (Red).
2.  Write minimum code to pass (Green).
3.  Refactor (Refactor).

This ensures better design and bug-free code in Node.js projects.
