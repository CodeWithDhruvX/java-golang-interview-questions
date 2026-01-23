## 🟢 Advanced API & Testing (Questions 351-400)

### Question 351: How do you handle partial success in batch APIs?

**Answer:**
Return `207 Multi-Status` or `200 OK` with a detailed result array.

**Response Structure:**
```json
{
  "results": [
    { "id": 1, "status": "success" },
    { "id": 2, "status": "error", "message": "fail" }
  ]
}
```

---

### Question 352: How do you enforce strong input validation across routes?

**Answer:**
Use middleware libraries like `express-validator` or `celebrate` (wraps Joi).

**Code:**
```javascript
const { body, validationResult } = require('express-validator');

app.post('/user', body('email').isEmail(), (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) return res.status(400).json({ errors: errors.array() });
  // process...
});
```

---

### Question 353: What is the role of `Accept` and `Content-Type` headers in REST APIs?

**Answer:**
*   **Content-Type:** Tells the server what the client sent (e.g., JSON).
*   **Accept:** Tells the server what the client *wants* back (e.g., XML or JSON).

---

### Question 354: How do you implement HTTP content negotiation?

**Answer:**
Use `req.accepts()`.

**Code:**
```javascript
app.get('/data', (req, res) => {
  const type = req.accepts(['json', 'html']);
  if (type === 'json') res.json({ data: 1 });
  else if (type === 'html') res.send('<p>Data: 1</p>');
  else res.status(406).end();
});
```

---

### Question 355: What’s the best way to paginate large datasets efficiently?

**Answer:**
Avoid `SKIP`/`OFFSET` (it scans previous rows). Use **Cursor-based Pagination** (Keyset).

**Query:** `WHERE id > last_seen_id LIMIT 10`.

---

### Question 356: What is structured logging?

**Answer:**
Logging in JSON format instead of plain text strings. This allows log aggregation tools (ELK, Splunk) to parse fields (`userId`, `level`) easily.

**Log:** `{"level":"info","msg":"Login","userId":123}`

---

### Question 357: How do you correlate logs across microservices?

**Answer:**
Generate a `Trace-ID` (Correlation ID) at the edge (Gateway). Pass it to every downstream service in HTTP Headers (`X-Trace-Id`). Log it in every message.

---

### Question 358: What are trace IDs and how do you generate them?

**Answer:**
A unique UUID.
Use `uuid` package. Use `AsyncLocalStorage` to store it for the request duration.

---

### Question 359: How would you implement custom log levels?

**Answer:**
In Winston:
```javascript
const myLevels = { levels: { critical: 0, error: 1, info: 2 } };
const logger = winston.createLogger({ levels: myLevels.levels });
logger.critical('Wake up!');
```

---

### Question 360: What is a sampling rate in logging/monitoring?

**Answer:**
Logging 100% of requests is expensive. Sampling means logging only 1% or 10% of requests to get a statistical view without performance penalty.

---

### Question 361: How do you prevent timing attacks in password comparisons?

**Answer:**
Do not compare strings with `===`. It fails faster on the first mismatch character (leaking info).
Use `crypto.timingSafeEqual(a, b)` which takes constant time regardless of content.

---

### Question 362: What is the difference between HTTPS and HTTP/2 security in Node?

**Answer:**
HTTP/2 effectively requires TLS (HTTPS). Browsers won't support H2 without encryption. In Node, `http2.createSecureServer` is used.

---

### Question 363: How do you protect Node.js apps behind a reverse proxy?

**Answer:**
Set `app.set('trust proxy', 1)` in Express. This allows Express to trust headers like `X-Forwarded-For` set by Nginx, so `req.ip` and `req.protocol` are correct.

---

### Question 364: What is secure cookie flag and how do you use it?

**Answer:**
Ensures cookies are sent *only* over HTTPS.
`res.cookie('auth', token, { secure: true, httpOnly: true });`

---

### Question 365: What is Subresource Integrity (SRI) and does it apply to Node?

**Answer:**
SRI is a browser feature verifying fetched scripts (CDN) match a hash.
Node.js acts as the server serving these tags:
`<script src="..." integrity="sha384-..." crossorigin="anonymous"></script>`.

---

### Question 366: How do you mock time-dependent logic in Node.js tests?

**Answer:**
Use **Fake Timers**.
Jest: `jest.useFakeTimers()`.
Sinon: `sinon.useFakeTimers()`.

**Code:**
```javascript
jest.useFakeTimers();
test('waits 1 sec', () => {
  const spy = jest.fn();
  setTimeout(spy, 1000);
  jest.runAllTimers(); // Fast forwarded
  expect(spy).toHaveBeenCalled();
});
```

---

### Question 367: What’s the difference between spying and mocking?

**Answer:**
*   **Spy:** Wraps a real function. It tracks calls but still executes the real code.
*   **Mock:** Replaces the function. Trace interactions but **does not** execute real code.

---

### Question 368: How do you test a function that depends on `fs` without real file I/O?

**Answer:**
Use `mock-fs` (simulates file system in memory) or manual mocks.

```javascript
jest.mock('fs');
const fs = require('fs');
fs.readFileSync.mockReturnValue('fake content');
```

---

### Question 369: How do you intercept and assert HTTP requests in integration tests?

**Answer:**
Use `nock`. It intercepts outgoing HTTP requests and responds with fixtures.

**Code:**
```javascript
const nock = require('nock');
const scope = nock('http://api.com')
  .get('/users')
  .reply(200, { users: [] });
// Now axios.get('http://api.com/users') hits nock
```

---

### Question 370: How do you test error-handling middleware?

**Answer:**
Trigger an error in a route and assert the response status.

```javascript
app.get('/error', (req, res, next) => next(new Error('fail')));

request(app)
  .get('/error')
  .expect(500);
```

---

### Question 371: How do you parse CLI arguments in Node.js?

**Answer:**
Raw: `process.argv`.
Libraries: `yargs`, `commander`.

---

### Question 372: What’s the difference between `yargs` and `commander`?

**Answer:**
Both are popular CLI builders.
*   **Commander:** More opinionated, fluent API.
*   **Yargs:** Pirate-themed, very powerful parsing configurations.

---

### Question 373: How do you read and mask password input in the terminal?

**Answer:**
Use `prompts` or `inquirer` library with type `password`.

---

### Question 374: How do you implement colored terminal output?

**Answer:**
Use libraries like `chalk` or `colors`.
Conceptually: Uses ANSI escape codes (`\x1b[31m` -> Red).

---

### Question 375: How do you handle interactive prompts in a Node CLI tool?

**Answer:**
Use `inquirer`.

**Code:**
```javascript
const answers = await inquirer.prompt([
  { type: 'input', name: 'user', message: 'Username?' }
]);
```

---

### Question 376: How do you enforce commit message conventions with Node tooling?

**Answer:**
**Commitlint** + **Husky**.
Husky triggers `commit-msg` hook. Commitlint checks if message matches format (Conventional Commits: `feat: new logic`).

---

### Question 377: What is Prettier and how does it differ from ESLint?

**Answer:**
*   **Prettier:** Code Formatter (Style). It rewrites code to look good (Indentation, wrapping).
*   **ESLint:** Linter (Quality). It finds bugs (unused vars, bad logic).

---

### Question 378: How do you configure import sorting in ESLint?

**Answer:**
Use `eslint-plugin-import` or `eslint-plugin-simple-import-sort`.

---

### Question 379: What is the purpose of `.editorconfig` in a Node.js project?

**Answer:**
It helps maintain consistent coding styles (indent style, charset) between different editors and IDES. It runs before Prettier.

---

### Question 380: What are shared ESLint configurations?

**Answer:**
NPM packages that export ESLint config (e.g., `eslint-config-airbnb`). You extend them in your `.eslintrc`.

---

### Question 381: What is the difference between peerDependencies and devDependencies?

**Answer:**
*   **devDependencies:** Dependencies for building/testing (e.g., Jest). Not installed in prod.
*   **peerDependencies:** Dependencies your package expects the **host** app to have (e.g., React plugin needs React).

---

### Question 382: How does npm handle circular dependencies?

**Answer:**
NPM installs allowed circular dependencies. However, in code, circular `require()` can lead to empty objects (Q196).

---

### Question 383: How does `npm audit` work internally?

**Answer:**
It submits the dependency tree (just names/versions) to the public npm registry audit endpoint, which returns known vulnerabilities.

---

### Question 384: How do you pin exact versions for all packages?

**Answer:**
`npm config set save-exact=true`.
This writes versions without `^` or `~` in `package.json`.

---

### Question 385: What is a package-lock.json file and why is it important?

**Answer:**
It locks the exact version of every dependency (and sub-dependency) installed. It guarantees that `npm install` produces the exact same node_modules tree on every machine.

---

### Question 386: Can you run Node.js modules in a browser?

**Answer:**
Not directly. Browsers don't have `require`, `fs`, `process`.
You must use a bundler (Webpack, Vite) to polyfill/bundle them.

---

### Question 387: What features are exclusive to Node.js and not available in browsers?

**Answer:**
1.  OS Signal handling.
2.  File System access (`fs`).
3.  Raw TCP/UDP Sockets.
4.  Creating HTTP Servers.

---

### Question 388: How do you polyfill a Node.js module for frontend use?

**Answer:**
Webpack/Vite have plugins (`node-polyfill-webpack-plugin`).
They allow using `Buffer` or `process` in browser by injecting browser-compatible implementations.

---

### Question 389: How is `fetch()` in browsers different from `node-fetch`?

**Answer:**
*   Browser Fetch: Handles Cookies/CORS natively. Relative URLs work.
*   Node Fetch: No Cookies/CORS logic by default. Must use absolute URLs.

---

### Question 390: How do you implement server-side rendering (SSR) using Node?

**Answer:**
1.  Server receives HTTP request.
2.  Node reads React/Vue component.
3.  Renders component to HTML string (`renderToString`).
4.  Sends HTML to client.
5.  Client "hydrates" it (attaches events).

---

### Question 391: How do you expose a GraphQL API using Node.js?

**Answer:**
Use `Apollo Server` or `express-graphql`.
Define TypeDefs (Schema) and Resolvers (Logic).

**Code:**
```javascript
const { ApolloServer, gql } = require('apollo-server');

const typeDefs = gql`type Query { hello: String }`;
const resolvers = { Query: { hello: () => 'Hello world!' } };

const server = new ApolloServer({ typeDefs, resolvers });
server.listen().then(({ url }) => console.log(`🚀  ${url}`));
```

---

### Question 392: How do you integrate gRPC and REST in a Node app?

**Answer:**
You can run both servers or use gRPC for internal microservices and a REST Gateway (like `grpc-gateway` logic) to expose it to web clients.

---

### Question 393: What is a BFF (Backend-for-Frontend) and how would you implement it?

**Answer:**
A dedicated backend layer for a specific frontend (Mobile BFF, Web BFF). It aggregates calls to microservices and formats data specifically for that UI. Built using Express/Fastify / GraphQL.

---

### Question 394: How do you stream audio/video data from Node.js?

**Answer:**
Use `fs.createReadStream` with `start` and `end` options based on the `Range` header sent by the browser. Return `206 Partial Content`.

---

### Question 395: How do you serve both a web app and API from one Node server?

**Answer:**
Ordering matters. API first, then Static Files, then "Catch All" for SPA.

```javascript
app.use('/api', apiRoutes);
app.use(express.static('build'));
app.get('*', (req, res) => res.sendFile('build/index.html'));
```

---

### Question 396: What would you do if a Node.js app leaks file descriptors?

**Answer:**
Symptoms: `EMFILE: too many open files`.
Cause: Opening files/sockets without closing them.
**Fix:** Ensure `fs.close()` or `socket.destroy()` is called in `finally` blocks or error handlers. Use `graceful-fs`.

---

### Question 397: What happens if you bind to a port already in use?

**Answer:**
(Recap Q197). `EADDRINUSE`. App crashes.

---

### Question 398: How do you protect an app from being overwhelmed by too many clients?

**Answer:**
**Backpressure** (Connection limits).
**Rate Limiting**.
**Load Shedding:** If CPU > 90%, immediately return 503.

---

### Question 399: How would you safely upgrade an API in production with zero downtime?

**Answer:**
Rolling update.
V1 is running. Spin up V2. Health check V2. Route traffic to V2. Drain V1. Stop V1.

---

### Question 400: How do you rollback a deployment if your new Node.js release is buggy?

**Answer:**
Point the Load Balancer/Router back to the previous Docker Image tag / Deployment version immediately.
Ideally, automated: If Error Rate > 5%, auto-rollback.
