## 🟢 Advanced Tools & Scenarios (Questions 151-200)

### Question 151: What are the benefits of using TypeScript with Node.js?

**Answer:**
TypeScript adds static typing to JavaScript, which is beneficial for large Node.js applications.

**Benefits:**
1.  **Type Safety:** Catch errors at compile time (e.g., passing string to math function).
2.  **IntelliSense:** Better autocompletion in IDEs.
3.  **Maintainability:** Interfaces and types document the code structure.

**Example:**
```typescript
interface User {
  id: number;
  name: string;
}

function getUser(user: User): string {
  return `Hello ${user.name}`; // Properties strictly checked
}
```

---

### Question 152: How do you set up a Node.js project with TypeScript?

**Answer:**
1.  Initialize: `npm init -y`
2.  Install: `npm install typescript ts-node @types/node --save-dev`
3.  Config: `npx tsc --init` (creates `tsconfig.json`)
4.  Run: `npx ts-node src/index.ts`

---

### Question 153: How do you declare types for Express request and response objects?

**Answer:**
Install `@types/express` and import types.

**Code:**
```typescript
import { Request, Response } from 'express';

app.get('/user', (req: Request, res: Response) => {
  res.json({ name: 'John' });
});
```

---

### Question 154: How do you handle module resolution issues in TypeScript?

**Answer:**
Configure `compilerOptions` in `tsconfig.json`.
*   **`moduleResolution`**: Set to `"node"`.
*   **`baseUrl`** and **`paths`**: For aliases (e.g., `@/utils` mapping to `src/utils`).

**tsconfig.json:**
```json
{
  "compilerOptions": {
    "moduleResolution": "node",
    "baseUrl": "./",
    "paths": {
      "@utils/*": ["src/utils/*"]
    }
  }
}
```

---

### Question 155: What is the `tsconfig.json` file used for?

**Answer:**
It specifies the root files and the compiler options required to compile the project.

**Common Settings:**
*   `target`: JS version to output (e.g., "ES6").
*   `outDir`: Where to put compiled `.js` files.
*   `strict`: Enable all strict type-checking options.

---

### Question 156: What is an ORM, and name a few Node.js ORMs?

**Answer:**
**ORM (Object-Relational Mapping)** allows you to query and manipulate data from a database using object-oriented paradigms (classes/methods) instead of writing raw SQL.

**Popular ORMs:**
1.  **Sequelize:** (SQL) Feature-rich, older.
2.  **TypeORM:** (SQL) Great with TypeScript.
3.  **Prisma:** (SQL) Modern, strictly typed, auto-generated client.
4.  **Mongoose:** (NoSQL) For MongoDB.

---

### Question 157: How does Sequelize differ from TypeORM?

**Answer:**
*   **Sequelize:** Traditional callback/promise-based, defines models via functions. heavily dynamic.
*   **TypeORM:** Uses Decorators (like Java/Hibernate) and classes to define entities. Strong integration with TypeScript.

**TypeORM Entity:**
```typescript
@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;
}
```

---

### Question 158: How do you define associations in Sequelize?

**Answer:**
Using methods like `hasOne`, `belongsTo`, `hasMany`.

**Code:**
```javascript
const User = sequelize.define('User', { /* ... */ });
const Task = sequelize.define('Task', { /* ... */ });

// User has many Tasks
User.hasMany(Task);
Task.belongsTo(User);

// usage
User.findAll({ include: Task });
```

---

### Question 159: How do you perform transactions in Sequelize?

**Answer:**
Pass a `{ transaction: t }` option to query methods.

**Code:**
```javascript
const t = await sequelize.transaction();

try {
  await User.create({ name: 'Alice' }, { transaction: t });
  await User.create({ name: 'Bob' }, { transaction: t });
  
  await t.commit();
} catch (error) {
  await t.rollback();
}
```

---

### Question 160: How can you write raw SQL queries in a Node.js app?

**Answer:**
Most ORMs provide a method for raw queries, or use a driver like `pg`.

**Using `pg` (Postgres):**
```javascript
const { Client } = require('pg');
const client = new Client();
await client.connect();

const res = await client.query('SELECT $1::text as message', ['Hello world!']);
console.log(res.rows[0].message);
await client.end();
```

---

### Question 161: What is a service layer in a Node.js application?

**Answer:**
The Service Layer contains the **business logic**. It sits between the Controller (API) and the Data Access Layer (Model).

**Flow:**
Controller (`req.body`) → Service (Calculate, Validate) → Model (DB Save).

**Benefit:** Keeps controllers "skinny" and reusable.

---

### Question 162: What is dependency injection and how is it implemented in Node.js?

**Answer:**
Dependency Injection (DI) is a design pattern where dependencies (like DB connection, Services) are passed into a module rather than hardcoded inside it.

**Manual DI:**
```javascript
// service.js
class UserService {
  constructor(db) { this.db = db; }
  save(user) { this.db.insert(user); }
}

// app.js
const db = require('./db');
const service = new UserService(db); // Inject db
```

**Frameworks:** NestJS uses DI heavily with `@Injectable()`.

---

### Question 163: What is the factory pattern and when would you use it?

**Answer:**
Factory pattern is a function that returns an object/instance. It creates objects without exposing the instantiation logic to the client.

**Example:**
loggerFactory.js:
```javascript
const createFileLogger = () => new FileLogger();
const createConsoleLogger = () => new ConsoleLogger();

module.exports = (type) => {
  if (type === 'file') return createFileLogger();
  return createConsoleLogger();
};
```

---

### Question 164: How would you implement a queue system in Node.js?

**Answer:**
For simple in-memory queues: Arrays.
For production queues: Redis-based libraries like **Bull** or **Bee-Queue**.

**Bull Example:**
```javascript
const Queue = require('bull');
const emailQueue = new Queue('sending emails');

// Consumer
emailQueue.process(async (job) => {
  await sendEmail(job.data.email);
});

// Producer
emailQueue.add({ email: 'user@example.com' });
```

---

### Question 165: What is a monorepo and how does it apply to Node.js apps?

**Answer:**
A Monorepo is a single Git repository holding multiple projects (e.g., backend, frontend, shared-libs).

**Tools:**
*   **Workspaces (npm/yarn/pnpm):** Link packages locally.
*   **Lerna / Nx / TurboRepo:** Build tools for monorepos.

**Structure:**
```
/packages
  /api (Node backend)
  /web (React frontend)
  /common (Shared types/utils)
```

---

### Question 166: What is WebSocket and how is it used in Node.js?

**Answer:**
WebSocket is a protocol providing full-duplex communication channels over a single TCP connection.

**Usage:**
Native `ws` library.
```javascript
const WebSocket = require('ws');
const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', ws => {
  ws.on('message', message => {
    console.log('received: %s', message);
  });
  ws.send('something');
});
```

---

### Question 167: How does Socket.IO differ from WebSocket?

**Answer:**
*   **WebSocket:** The standardized protocol.
*   **Socket.IO:** A library that uses WebSockets but provides fallbacks (Long Polling) for older browsers. It adds features like **Rooms**, **Broadcast**, and **Reconnection logic**.

**Crucially:** Socket.IO client is not compatible with standard WebSocket server and vice versa.

---

### Question 168: How do you handle rooms and namespaces in Socket.IO?

**Answer:**
*   **Namespaces (`/admin`):** Separate communication channels.
*   **Rooms (`room1`):** Sub-channels within namespaces that sockets can join/leave.

**Code:**
```javascript
io.on('connection', (socket) => {
  socket.join('chat_room'); // Join
  
  // Emit to room only
  io.to('chat_room').emit('msg', 'Hello Room');
});
```

---

### Question 169: What are long polling and server-sent events?

**Answer:**
*   **Long Polling:** Client requests data. Server holds connection open until data is available. Client gets data, then immediately requests again. (Fallback for WebSockets).
*   **SSE (Server-Sent Events):** One-way channel from Server to Client over HTTP. Good for news feeds, stocks.

**SSE Header:** `Content-Type: text/event-stream`.

---

### Question 170: How do you implement a pub/sub system in Node.js?

**Answer:**
Using `EventEmitter` (internal) or Redis (external).

**Redis Pub/Sub:**
```javascript
const publisher = redis.createClient();
const subscriber = redis.createClient();

subscriber.subscribe('news');

subscriber.on('message', (channel, message) => {
  console.log(message);
});

publisher.publish('news', 'Breaking News!');
```

---

### Question 171: What is a microservice?

**Answer:**
An architectural style where an application is structured as a collection of loosely coupled services, each implementing specific business capabilities (e.g., User Service, Order Service).

---

### Question 172: How would you build microservices using Node.js?

**Answer:**
1.  **Transport:** HTTP/REST or gRPC.
2.  **Discovery:** Consul or environment variables (K8s Service).
3.  **Communication:** API calls or Message Queue (RabbitMQ).

**Example:**
Service A (`POST /order`) → RabbitMQ → Service B (Process Order).

---

### Question 173: What is gRPC and how is it used in Node.js?

**Answer:**
gRPC is a high-performance RPC framework by Google. It uses **Protocol Buffers** (binary) instead of JSON.

**Benefits:** Smaller payloads, strongly typed contracts (`.proto`), bidirectional streaming.
**Library:** `@grpc/grpc-js` and `@grpc/proto-loader`.

---

### Question 174: What is RabbitMQ and how do you use it with Node.js?

**Answer:**
RabbitMQ is a message broker (AMQP) for async communication between services.

**Library:** `amqplib`.

**Producer:**
```javascript
const amqp = require('amqplib');
const conn = await amqp.connect('amqp://localhost');
const channel = await conn.createChannel();
channel.sendToQueue('tasks', Buffer.from('Work'));
```

---

### Question 175: What is the role of API Gateway in microservices?

**Answer:**
It is the single entry point for clients. It routes requests to appropriate microservices.
**Features:** Authentication, Rate Limiting, SSL Termination, Request Aggregation.

---

### Question 176: How do you write a GitHub Action to test a Node.js app?

**Answer:**
Create `.github/workflows/test.yml`.

**YAML:**
```yaml
name: Node CI
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-node@v2
      with:
        node-version: '16'
    - run: npm ci
    - run: npm test
```

---

### Question 177: How do you use `nodemon` in development?

**Answer:**
`nodemon` monitors for file changes and automatically restarts the server.

**Usage:**
1.  Install: `npm install -g nodemon`
2.  Run: `nodemon app.js`
3.  Config: `nodemon.json` (ignore files, delay, extensions).

---

### Question 178: How do you use `lint-staged` with Node.js?

**Answer:**
Run linters only on files that are staged in git, usually via a pre-commit hook (Husky).

**package.json:**
```json
"lint-staged": {
  "*.js": "eslint --fix"
}
```
This ensures bad code is never committed.

---

### Question 179: What is semantic versioning?

**Answer:**
SemVer (Major.Minor.Patch) e.g., `1.0.0`.
*   **Major**: Breaking changes.
*   **Minor**: New features (backward compatible).
*   **Patch**: Bug fixes.

**NPM Notation:**
*   `^1.2.3`: Up to next Major (1.x.x).
*   `~1.2.3`: Up to next Minor (1.2.x).

---

### Question 180: How do you build and publish a Node.js package?

**Answer:**
1.  **Code:** Write library.
2.  **Config:** Set `name`, `main`, `version` in `package.json`.
3.  **Log in:** `npm login`.
4.  **Publish:** `npm publish`.
5.  **Tag:** `npm publish --tag beta` (optional).

---

### Question 181: How do you implement centralized error handling in Express?

**Answer:**
Define an error-handling middleware (4 arguments) at the end of `app.js`.

**Code:**
```javascript
app.use((err, req, res, next) => {
  console.error(err.stack);
  const status = err.statusCode || 500;
  res.status(status).json({
    error: {
      message: err.message || 'Internal Server Error'
    }
  });
});
```

---

### Question 182: What is the difference between `throw` and `next(err)`?

**Answer:**
*   **`throw new Error()`**: Works in synchronous code. In async (Promises), it triggers `.catch`.
*   **`next(err)`**: Express specific. Passes the error to the global error middleware.

**Wait:** In Express 5 (and modern Express 4 + async handlers), `throw` inside async functions is now caught. In older Express, `throw` in async callbacks would crash the app; you *had* to use `next(err)`.

---

### Question 183: What is Winston and how is it used?

**Answer:**
Winston is a versatile logging library. It supports multiple transports (Console, File, HTTP).

**Code:**
```javascript
const winston = require('winston');

const logger = winston.createLogger({
  level: 'info',
  transports: [
    new winston.transports.File({ filename: 'error.log', level: 'error' }),
    new winston.transports.Console()
  ]
});

logger.info('Hello Winston');
```

---

### Question 184: How do you log exceptions in Node.js?

**Answer:**
Winston can handle uncaught exceptions automatically.

**Code:**
```javascript
logger.exceptions.handle(
  new winston.transports.File({ filename: 'exceptions.log' })
);
```

---

### Question 185: How can you integrate Sentry with Node.js?

**Answer:**
Sentry tracks errors and performance.

**Setup:**
```javascript
const Sentry = require("@sentry/node");

Sentry.init({ dsn: "YOUR_DSN" });

// RequestHandler creates a separate execution context
app.use(Sentry.Handlers.requestHandler());

// All controllers...

// The error handler must be before any other error middleware
app.use(Sentry.Handlers.errorHandler());
```

---

### Question 186: How do you enable ES modules in Node.js?

**Answer:**
Two ways:
1.  Extension: Use `.mjs` file extension.
2.  Package: Add `"type": "module"` in `package.json`.

Then use `import` / `export`.

---

### Question 187: What’s the difference between CommonJS and ES Modules?

**Answer:**
*   **CommonJS (`require`):** Dynamic, synchronous, `module.exports`, default in Node.
*   **ESM (`import`):** Static, asynchronous, `export default`, standard in JS spec.

**Key:** ESM module code runs in strict mode by default. `__dirname` is not available in ESM (use `import.meta.url`).

---

### Question 188: Can you use both ES modules and CommonJS together?

**Answer:**
Yes, but limited.
*   ESM can import CJS (`import foo from 'foo.cjs'`).
*   CJS cannot easily `require()` ESM (because ESM is async). You must use `await import('foo.mjs')`.

---

### Question 189: What is tree-shaking and how does it work in Node?

**Answer:**
Tree-shaking removes unused code. It relies on the static structure of ES Modules.
Node.js runtime doesn't tree-shake (it loads files). Tree-shaking is done by bundlers (Webpack, Rollup) during a build step before deploying.

---

### Question 190: How do you use top-level `await` in Node.js?

**Answer:**
Available in ES Modules (Node 14.8+). You can use `await` outside of async functions.

**Code (`script.mjs`):**
```javascript
const response = await fetch('https://api.github.com');
const data = await response.json();
console.log(data);
```

---

### Question 191: What is async_hooks in Node.js?

**Answer:**
The `async_hooks` module provides an API to track asynchronous resources (like promises, sockets). It is used to build APM tools or request context tracking (like Correlation IDs).

**Events:** `init`, `before`, `after`, `destroy`.

---

### Question 192: What are V8 snapshots?

**Answer:**
A startup performance feature. You can take a snapshot of the heap after initialization and save it. When starting the app again, V8 loads the snapshot instead of re-executing initialization code, significantly speeding up bootstrap time.

---

### Question 193: What is the inspector module used for?

**Answer:**
The `inspector` module provides an API for interacting with the V8 inspector (debugging protocol) programmatically.
You can take heap profiles or set breakpoints from within the code.

---

### Question 194: How do you use perf_hooks to measure performance?

**Answer:**
Using `performance.now()` for high-precision timing.

**Code:**
```javascript
const { performance } = require('perf_hooks');

const start = performance.now();
doWork();
const end = performance.now();

console.log(`Duration: ${end - start}ms`);
```

---

### Question 195: What is an AbortController in Node.js?

**Answer:**
A standard API (from Web) to cancel asynchronous tasks (fetch, timers, streams).

**Code:**
```javascript
const controller = new AbortController();
const { signal } = controller;

setTimeout(() => controller.abort(), 100); // Timeout after 100ms

try {
  await fetch(url, { signal });
} catch (err) {
  if (err.name === 'AbortError') console.log('Request aborted');
}
```

---

### Question 196: Why might a `require()` fail even if the file exists?

**Answer:**
1.  **Circular Dependency:** If Module A requires B, and B requires A, and one uses the exported value immediately, it might be undefined.
2.  **Case Sensitivity:** Windows is case-insensitive, Linux is case-sensitive. `require('./File')` works on Windows but fails on Linux if file is `file.js`.
3.  **Permissions:** Read access denied.

---

### Question 197: What happens if you listen on a port already in use?

**Answer:**
Node.js throws an `EADDRINUSE` error.

**Handling:**
```javascript
server.on('error', (e) => {
  if (e.code === 'EADDRINUSE') {
    console.log('Port busy, retrying or exiting...');
  }
});
```

---

### Question 198: What are zombie processes and how do you prevent them?

**Answer:**
A zombie is a child process that has completed but its parent hasn't read its exit status (`wait()`).
In Node.js, `child_process` usually handles this. If using Docker, use `tini` as init process (`pid 1`) to reap zombies.

---

### Question 199: How do you detect and handle memory leaks?

**Answer:**
(See Q48).
Use `heapdump` to create snapshots and inspect in Chrome DevTools. Look for objects that grow indefinitely (e.g., Arrays used as queues without clearing).

---

### Question 200: How do you troubleshoot high CPU usage in a Node.js app?

**Answer:**
1.  **Profiler:** Use `node --prof` to generate a tick log. Use `node --prof-process` to analyze it.
2.  **Flame Graphs:** Visualize where CPU time is spent (stack depth).
3.  **Inspector:** Chrome DevTools > Profiler tab > Record CPU Profile.
4.  **Look for:** Long sync loops, heavy crypto, huge JSON.parse.
