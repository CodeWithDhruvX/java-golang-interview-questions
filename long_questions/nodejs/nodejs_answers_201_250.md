## 🟢 System Design & Architecture (Questions 201-250)

### Question 201: How would you structure a large Node.js project?

**Answer:**
Separate concerns using a 3-layer architecture (Controller, Service, Data Access).

**Structure:**
```text
/src
  /api              # Controllers (Routes)
  /services         # Business Logic
  /models           # Database Schema
  /config           # Env variables
  /jobs             # Background tasks
  /loaders          # Startup script (DB connect, Express init)
  /subscribers      # Event listeners
  /types            # TypeScript types
```
*   **Controller:** Accepts request, returns response.
*   **Service:** Contains the "business" (e.g., if user exists, send email).
*   **Model:** Direct DB commands.

---

### Question 202: What is a middleware pipeline, and how do you design one?

**Answer:**
The pipeline is the chain of functions a request passes through. Design it from generic to specific.

**Ideally:**
1.  **Security:** (Helmet, CORS)
2.  **Parsers:** (Body parser, Cookie parser)
3.  **Loggers:** (Morgan)
4.  **Static Files**
5.  **Routes:** (API)
6.  **Error Handling:** (404, 500)

**Code:**
```javascript
app.use(helmet());
app.use(express.json());
app.use(logger);
app.use('/api', apiRoutes);
app.use(errorHandler);
```

---

### Question 203: How do you design a modular Express.js app?

**Answer:**
Break the app into "components" or "modules" based on domain (e.g., User, Order, Payment) rather than technical role (Controller, Model).

**Folder-by-Feature:**
```text
/components
  /users
    index.js        # Routes
    user.controller.js
    user.service.js
    user.model.js
  /orders
    index.js
    ...
```

---

### Question 204: What is a repository pattern in Node.js?

**Answer:**
It abstracts the data layer. The service layer calls the Repository, which handles the DB (SQL, Mongo, File). This makes swapping databases easier.

**Code:**
```javascript
class UserRepository {
  async findById(id) {
    return UserModel.findOne({ _id: id });
  }
}
// Service calls Repository, not Model directly.
```

---

### Question 205: How would you implement caching at the route level?

**Answer:**
Use middleware to check if data exists in Redis before executing the controller.

**Code:**
```javascript
const checkCache = async (req, res, next) => {
  const data = await redis.get(req.originalUrl);
  if (data) {
    return res.send(JSON.parse(data));
  }
  next(); // Proceed to DB
};

app.get('/data', checkCache, dataController);
```

---

### Question 206: When would you use event emitters over REST APIs?

**Answer:**
Use Event Emitters for **internal** communication within the same process (decoupling modules).
Use REST APIs for **external** communication between services.

**Scenario:**
User register -> triggers 'user_signup' event -> EmailService listens and sends email. The Signup controller doesn't need to know about the EmailService.

---

### Question 207: What is an API gateway, and how do you build one with Node.js?

**Answer:**
An API Gateway sits in front of microservices. It handles routing, auth, and rate limiting.
**Tools:** `express-gateway`, `fast-gateway`.

**Simple Gateway:**
```javascript
const gateway = require('fast-gateway');
const server = gateway({
  routes: [{
    prefix: '/users',
    target: 'http://localhost:3001'
  }]
});
server.start(8080);
```

---

### Question 208: How do you manage versioning in a Node.js API?

**Answer:**
1.  **URL Versioning:** `/api/v1/resource` (Most common).
2.  **Header Versioning:** `Accept-Version: v1`.

**Express Implementation:**
```javascript
const v1Router = require('./v1/routes');
const v2Router = require('./v2/routes');

app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);
```

---

### Question 209: What are idempotent APIs, and how would you implement them?

**Answer:**
Idempotent means making the same request multiple times has the same effect as making it once (e.g., Delete, Put).
**POST** is usually not idempotent. To make it so, clients send a unique `Idempotency-Key` header.
Server checks Redis: "Did I process this key?"
If yes -> Return saved response.
If no -> Process.

---

### Question 210: How do you build a multi-tenant application in Node.js?

**Answer:**
1.  **Database per Tenant:** Secure, but harder to maintain.
2.  **Shared Database (discriminator column):** `tenant_id` in every table.

**Middleware:**
```javascript
app.use((req, res, next) => {
  const tenant = req.headers['x-tenant-id'];
  req.db = getDatabaseConnection(tenant); // Switch context
  next();
});
```

---

### Question 211: What is the difference between `pm2 restart` and `pm2 reload`?

**Answer:**
*   **Restart:** Kills the process and starts it again. High downtime (ms/seconds).
*   **Reload:** (Graceful Reload) Sends a signal to the process to finish current requests before stopping. If in cluster mode, it restarts workers one by one (Zero Downtime).

---

### Question 212: How do you use NGINX as a reverse proxy for Node.js?

**Answer:**
(See Q93). Client -> Nginx (Port 80) -> Node.js (Port 3000).

**Why?**
*   SSL termination.
*   Serving static files faster.
*   Load balancing.

---

### Question 213: How do you deploy a Node.js app to AWS ECS?

**Answer:**
1.  Dockerize app (`Dockerfile`).
2.  Push image to ECR (Elastic Container Registry).
3.  Create Task Definition (CPU, RAM, Image).
4.  Create Service (Load Balancer, Desired Count).

---

### Question 214: What is serverless architecture? Can Node.js be used in serverless?

**Answer:**
Serverless means you run code (functions) without provisioning servers.
Node.js is the **most popular** runtime for serverless (AWS Lambda).

**Code (Lambda):**
```javascript
exports.handler = async (event) => {
  return {
    statusCode: 200,
    body: JSON.stringify('Hello from Lambda!'),
  };
};
```

---

### Question 215: How would you set up CI/CD for a Node.js monorepo?

**Answer:**
Use tools that support smart rebuilding (Nx, Turbo).
**Pipeline:**
1.  Detect changed packages.
2.  Build/Test only changed packages.
3.  Deploy only changed services.

---

### Question 216: What are some common pitfalls in Dockerizing Node.js apps?

**Answer:**
1.  **Running as Root:** Security risk. Use `USER node`.
2.  **Copying `node_modules`:** Don't copy from host. Run `npm install` inside.
3.  **Handling Signals:** Node doesn't handle `SIGTERM` correctly as PID 1. Use `tini` (`--init` flag).

---

### Question 217: How do you implement blue/green deployment in Node.js?

**Answer:**
Requires a router/load balancer.
1.  Blue (Current) is live.
2.  Deploy Green (New).
3.  Run tests on Green.
4.  Switch Router to point to Green.
5.  If error, switch back to Blue.

---

### Question 218: What are environment secrets, and how do you manage them securely?

**Answer:**
Secrets = Passwords, Keys.
Dev: `.env`.
Prod: Use Secret Managers (AWS Secrets Manager, HashiCorp Vault), or inject as Env Vars in K8s Secrets. **Never commit them.**

---

### Question 219: What is the impact of `NODE_ENV=production`?

**Answer:**
1.  **Express:** Disables stack traces in errors, caches view templates.
2.  **Dependencies:** Many skip debug logs/checks.
3.  **Performance:** Generally 3x faster.

**Always** set this in deployment.

---

### Question 220: How would you deploy a Node.js app to Kubernetes?

**Answer:**
1.  **Deployment:** Defines Pods (Node containers) and Replicas.
2.  **Service:** Exposes Pods (Internal IP).
3.  **Ingress:** Exposes Service to outside world (Routing).

---

### Question 221: How do you use `concurrently` in a Node.js project?

**Answer:**
It allows running multiple commands in one terminal (e.g., Server + Client).

**package.json:**
```json
"scripts": {
  "server": "node server.js",
  "client": "npm start --prefix client",
  "dev": "concurrently \"npm run server\" \"npm run client\""
}
```

---

### Question 222: What does `nvm` do?

**Answer:**
**Node Version Manager**. It lets you install and switch between multiple Node.js versions on the same machine.

`nvm install 16`
`nvm use 16`

---

### Question 223: How do you benchmark a Node.js app?

**Answer:**
Use **Autocannon** (written in Node) or **Apache Bench (ab)**.

**Command:**
```bash
npx autocannon -c 100 -d 10 http://localhost:3000
# 100 connections for 10 seconds
```

---

### Question 224: What is the purpose of `.npmrc`?

**Answer:**
Configuration for npm.
1.  **Registry:** Point to private registry (Artifactory).
2.  **Auth:** Tokens for private packages.
3.  **Behavior:** `save-exact=true`.

---

### Question 225: How do you analyze Node.js memory usage?

**Answer:**
`process.memoryUsage()`.
*   **rss:** Resident Set Size (Total memory).
*   **heapTotal:** V8 Heap allocated.
*   **heapUsed:** V8 Heap actually used by objects.

---

### Question 226: What is the purpose of the `debug` package?

**Answer:**
A tiny library for conditional logging based on namespaces.

```bash
DEBUG=worker:*,http node app.js
```

---

### Question 227: How do you configure path aliases in a Node.js + TypeScript project?

**Answer:**
1.  **tsconfig.json:** `paths: { "@/*": ["src/*"] }`
2.  **Runtime:** Typescript compiles paths but Node doesn't understand them. Use `module-alias` package or `tsconfig-paths` in dev.

---

### Question 228: What is `husky` and how do you use it?

**Answer:**
Husky manages Git hooks. It runs scripts automatically when you commit or push.
Use it to run tests or linting (`lint-staged`) before commit.

---

### Question 229: What is the difference between `npm ci` and `npm install`?

**Answer:**
*   **`npm install`**: Reads `package.json`, installs, and **updates** `package-lock.json`. Good for Dev.
*   **`npm ci`**: (Clean Install) Reads **only** `package-lock.json`. Deletes `node_modules`. Ensures exact reproducible builds. Good for CI/CD.

---

### Question 230: How do you lock down dependency versions in Node.js?

**Answer:**
1.  Use `package-lock.json` (Commit it).
2.  Use `npm ci`.
3.  Remove `^` (caret) from `package.json` to pin exact versions.

---

### Question 231: How does the async model of Node.js compare to Go or Python?

**Answer:**
*   **Node.js:** Single-threaded Event Loop. Good for I/O.
*   **Go:** Goroutines (Green threads). Multi-threaded. Good for CPU & I/O.
*   **Python:** Threading (limited by GIL) or Asyncio (Event loop similar to Node).

---

### Question 232: What is cooperative concurrency in Node.js?

**Answer:**
The Event Loop relies on tasks yielding control back to the loop. If a task takes too long, it "hogs" the CPU. It is **cooperative** because the code must cooperate by being non-blocking.

---

### Question 233: What is the maximum number of concurrent I/O operations Node.js can handle?

**Answer:**
Dependent on OS limits (File Descriptors) and memory.
Node can handle tens of thousands of connections because they are just objects in memory, not heavy OS threads.

---

### Question 234: How would you simulate a deadlock in Node.js?

**Answer:**
Node isn't prone to standard thread deadlocks (locks). But **Promise Deadlocks** happen if you await a promise that never resolves.

```javascript
let resolve;
const p = new Promise(r => resolve = r);
await p; // Deadlock if resolve() is never called
```

---

### Question 235: How can you safely perform parallel async operations?

**Answer:**
Use `Promise.all()` (fail fast) or `Promise.allSettled()` (wait for all).

```javascript
const [user, posts] = await Promise.all([
  fetchUser(),
  fetchPosts()
]);
```

---

### Question 236: How do you throttle or debounce API calls in Node.js?

**Answer:**
*   **Throttle:** Ensure function runs at most once in X ms.
*   **Debounce:** Ensure function runs only after X ms of inactivity.
Use `lodash`.

**Example:**
Saving to DB on every keystroke? Debounce it.

---

### Question 237: How would you implement retry logic with backoff in Node.js?

**Answer:**
Use a library like `axios-retry` or write a recursive function.

**Code:**
```javascript
async function fetchWithRetry(url, retries = 3, delay = 1000) {
  try {
    return await axios.get(url);
  } catch (err) {
    if (retries === 0) throw err;
    await new Promise(r => setTimeout(r, delay));
    return fetchWithRetry(url, retries - 1, delay * 2); // Exponential backoff
  }
}
```

---

### Question 238: Can Node.js handle millions of concurrent connections?

**Answer:**
Yes, if tuned properly (ulimit, kernel settings). This is known as the **C10M problem**.
However, usually, you'd cluster or load balance multiple instances rather than pushing one instance to 1M.

---

### Question 239: What is the difference between `Promise.all` and `Promise.allSettled`?

**Answer:**
*   **`Promise.all`**: Rejects immediately if **any** promise rejects.
*   **`Promise.allSettled`**: Waits for **all** to finish, regardless of success/fail. Returns array of status (`{ status: 'fulfilled', value: ... }`).

---

### Question 240: How would you build a task queue in memory?

**Answer:**
Simple array or `async.queue`.

**Code:**
```javascript
const async = require('async');

const q = async.queue((task, callback) => {
  console.log('Processing ' + task.name);
  callback();
}, 2); // concurrency 2

q.push({ name: 'Task 1' });
```
*Note: In-memory queues vanish on restart. use Redis for persistence.*

---

### Question 241: How do you handle async/await errors globally?

**Answer:**
In Express 5, it catches them.
In Express 4, use `express-async-errors` package, or wrap routes.

**Wrapper:**
```javascript
const asyncHandler = fn => (req, res, next) =>
  Promise.resolve(fn(req, res, next)).catch(next);

app.get('/', asyncHandler(async (req, res) => {
  throw new Error('BAM');
}));
```

---

### Question 242: What are operational vs. programmer errors in Node.js?

**Answer:**
*   **Operational:** Run-time problems (Network down, File not found). Handle these.
*   **Programmer:** Bugs (Reading undefined, Syntax error). Fix the code.

---

### Question 243: What is an unhandled rejection warning?

**Answer:**
Occurs when a Promise is rejected but no `.catch()` is attached.
Node.js prints a warning. In future versions, it will crash the process (exit code 1).

**Fix:** Keep `process.on('unhandledRejection', ...)` or good code practices.

---

### Question 244: What happens if you call `next()` multiple times in Express?

**Answer:**
It usually leads to header errors (`ERR_HTTP_HEADERS_SENT`) if both handlers try to send a response. Or it might execute the next middleware twice. It is a bug.

---

### Question 245: What’s the best way to handle 3rd-party API failures?

**Answer:**
1.  **Timeout:** Don't wait forever.
2.  **Circuit Breaker:** Stop calling if it fails repeatedly (Opossum library).
3.  **Fallback:** Return cached or default data.

---

### Question 246: How do closures impact memory usage in Node.js?

**Answer:**
Closures retain access to the outer scope. If a large object is in the outer scope and the closure stays alive (e.g., in a timer), the large object acts as a leak (cannot be collected).

---

### Question 247: What is a "tick" in the Node.js event loop?

**Answer:**
A tick refers to one complete pass through the event loop phases (Timers -> Poll -> Check...).
Or specifically `process.nextTick`, which runs "buckets" of tasks between operations.

---

### Question 248: What is inline caching in the V8 engine?

**Answer:**
Optimization technique. V8 "remembers" the structure (Offset of properties) of objects at specific call sites (`obj.x`). If the object type doesn't change, V8 skips lookup and accesses memory directly.

---

### Question 249: How do hidden classes impact performance?

**Answer:**
V8 creates hidden classes for objects. If you add properties dynamically in different orders, V8 creates different hidden classes, breaking optimization.
**Fix:** Initialize all properties in `constructor`.

---

### Question 250: What is the V8 heap limit and how can you increase it?

**Answer:**
Default is ~1.5GB (64-bit).
Increase using:
`node --max-old-space-size=4096 app.js` (Sets to 4GB).
