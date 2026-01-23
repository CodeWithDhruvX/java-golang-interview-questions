## 🟢 Ecosystem & Libraries (Questions 251-300)

### Question 251: What is the difference between Axios and `fetch()` in Node.js?

**Answer:**
*   **Axios:** External library. Auto-transforms JSON. Interceptors. Wide browser support.
*   **Fetch:** Built-in (Node 18+). Standard. Requires `.json()` call.

**Code:**
```javascript
// Axios
const res = await axios.get(url);
const data = res.data;

// Fetch
const res = await fetch(url);
const data = await res.json();
```

---

### Question 252: How does `multer` work for file uploads?

**Answer:**
Multer processes `multipart/form-data`. It adds a `body` object and a `file` (or `files`) object to the `req` object.

---

### Question 253: What are some alternatives to `express` for building APIs?

**Answer:**
1.  **Fastify:** High performance, low overhead.
2.  **Koa:** Created by Express team, smaller, uses async/await middleware.
3.  **NestJS:** Angular-style, Opinionated, TypeScript.
4.  **Hapi:** Config-centric.

---

### Question 254: How does `jsonwebtoken` verify tokens?

**Answer:**
It takes the token and the secret key. It re-hashes the header+payload with the secret and checks if it matches the signature part of the token.

**Code:**
```javascript
jwt.verify(token, 'secret', (err, decoded) => {
  if (err) throw err; // Invalid
});
```

---

### Question 255: How does `passport.js` handle different strategies?

**Answer:**
It uses the Strategy Pattern. You configure strategies (Local, Google, JWT) separately. The main middleware `passport.authenticate('strategy_name')` delegates to the chosen strategy instance.

---

### Question 256: What is `rate-limiter-flexible` and how does it work?

**Answer:**
A library to prevent DDoS/Brute force. It uses Token Bucket or Leaky Bucket algorithms. Can store state in Memory, Redis, Mongo.

---

### Question 257: What is `node-cron` used for?

**Answer:**
Task scheduling (cron jobs) in Node.js.

**Code:**
```javascript
const cron = require('node-cron');

// Run every minute
cron.schedule('* * * * *', () => {
  console.log('running a task every minute');
});
```

---

### Question 258: How does `nodemailer` work with Gmail?

**Answer:**
It acts as an SMTP client.

**Code:**
```javascript
const transporter = nodemailer.createTransport({
  service: 'gmail',
  auth: { user: 'email', pass: 'app-password' }
});

await transporter.sendMail({ to: '...', text: 'Hello' });
```

---

### Question 259: How does `bull` or `agenda` manage background jobs?

**Answer:**
*   **Bull:** Uses Redis. Robust.
*   **Agenda:** Uses MongoDB.

They persist jobs to DB, allowing them to survive server restarts and be processed by workers asynchronously.

---

### Question 260: What’s the difference between `joi` and `zod` for validation?

**Answer:**
*   **Joi:** Classic, powerful, schema-based.
*   **Zod:** Modern, TypeScript-first. Infers TS types from the schema.

**Zod:**
```javascript
const User = z.object({ name: z.string() });
type UserType = z.infer<typeof User>; // Automatic TS type!
```

---

### Question 261: How do you manage refresh tokens securely?

**Answer:**
1.  **Access Token:** Short live (15 min). InMemory/Variable.
2.  **Refresh Token:** Long lived (7 days). Store in **HttpOnly Cookie** (prevents XSS theft).

**Rotation:** When Refresh Token is used, issue a new Refresh Token and invalidate the old one (Detect reuse).

---

### Question 262: What is the OAuth2 implicit flow?

**Answer:**
Legacy flow where Access Token is returned directly in URL hash fragment. **Insecure** and deprecated. Use **Authorization Code Flow with PKCE** instead.

---

### Question 263: How would you implement RBAC in Node.js?

**Answer:**
**Role-Based Access Control**.
Middleware checks user role.

**Code:**
```javascript
const checkRole = (role) => (req, res, next) => {
  if (req.user.role !== role) return res.sendStatus(403);
  next();
};

app.delete('/users', checkRole('ADMIN'), controller);
```

---

### Question 264: How do you handle multi-factor authentication (MFA)?

**Answer:**
1.  Login (User/Pass).
2.  Generate TOTP (Time-based One Time Password) using `speakeasy`.
3.  Send to user / User checks Google Authenticator.
4.  Verify code.

---

### Question 265: How do you validate scopes or roles in a Node.js route?

**Answer:**
(See Q263). Or check OAuth scopes if using JWT.

```javascript
// Scope: "read:files"
if (!req.user.scopes.includes('read:files')) throw Forbidden;
```

---

### Question 266: What are race conditions in Socket.IO?

**Answer:**
When multiple events try to modify the same state simultaneously.
**Example:** Two users saving document.
**Fix:** Use versioning, locking, or atomic operations.

---

### Question 267: How do you ensure message order in WebSockets?

**Answer:**
TCP guarantees packet order.
However, app logic might fail (Async processing).
**Fix:** Append Sequence Number (`seq: 1`, `seq: 2`) to messages. Client buffers and processes in order.

---

### Question 268: How do you authenticate Socket.IO connections?

**Answer:**
Use the `auth` object in handshake.

**Server:**
```javascript
io.use((socket, next) => {
  const token = socket.handshake.auth.token;
  // Verify token...
  next();
});
```

---

### Question 269: How do you build a real-time collaborative editor using Node.js?

**Answer:**
Use **Operational Transformation (OT)** or **CRDTs (Conflict-free Replicated Data Types)** (libraries like Yjs).
Socket.IO broadcasts changes.

---

### Question 270: What is the best way to implement presence detection?

**Answer:**
**Socket.IO:**
*   `connection`: User Online (Set Redis Key).
*   `disconnect`: User Offline (Del Redis Key).
*   `heartbeat`: Update "last seen".

---

### Question 271: What is the purpose of APM in Node.js?

**Answer:**
Application Performance Monitoring.
Tools like DataDog/NewRelic hook into internal workings to track: Route latency, DB query time, External API calls, Errors.

---

### Question 272: How do you use OpenTelemetry with Node.js?

**Answer:**
Open standard for telemetry.
Install SDK (`@opentelemetry/sdk-node`). It auto-instruments http, pg, express modules to generate Traces.

---

### Question 273: What metrics should you monitor in a Node.js app?

**Answer:**
1.  **Event Loop Lag:** Is the CPU blocked?
2.  **Memory:** Heap Used vs Total.
3.  **Active Handles:** Open sockets/files.
4.  **RPS:** Requests per second.
5.  **Error Rate.**

---

### Question 274: How do you track slow queries in Node.js?

**Answer:**
1.  **DB Logs:** Enable slow query log in Mongo/Postgres.
2.  **APM:** Shows breakdown.
3.  **Middleware:** Measure duration.

```javascript
const start = Date.now();
await db.query(...);
if (Date.now() - start > 1000) log('Slow Query');
```

---

### Question 275: What is a flame graph and how is it useful?

**Answer:**
A visualization of stack traces. The x-axis is population, y-axis is stack depth.
**Wide bars** = Functions taking the most CPU time.
Used to find performance bottlenecks.

---

### Question 276: How do you prevent prototype pollution in Node.js?

**Answer:**
Attackers inject properties into `Object.prototype`.
**Fix:**
1.  Use `Object.create(null)` for plain maps.
2.  Freeze prototype: `Object.freeze(Object.prototype)`.
3.  Validate JSON input properly.

---

### Question 277: What is `eval()` and why should you avoid it?

**Answer:**
`eval()` executes a string as code.
**Risk:** Remote Code Execution (RCE). If user input ends up in eval, they own your server.
**Never use it.**

---

### Question 278: How does Node.js sanitize user input?

**Answer:**
It doesn't. You must use libraries.
1.  **XSS:** `xss` library to clean HTML.
2.  **SQL:** ORMs or Parameterized queries.
3.  **Validation:** `express-validator`.

---

### Question 279: What is SSRF and how can Node.js apps be affected?

**Answer:**
**Server-Side Request Forgery**.
If your app allows fetching generic URLs (e.g., "Image from URL"), an attacker can request `http://localhost:8080/admin` or `http://169.254.169.254` (AWS metadata) to steal secrets.

---

### Question 280: How do you audit Node.js dependencies for vulnerabilities?

**Answer:**
`npm audit`.
It checks your `package-lock.json` against a vulnerability database.
`npm audit fix` can upgrade packages automatically.

---

### Question 281: How do you implement HATEOAS in a Node.js API?

**Answer:**
Hypermedia as the Engine of Application State.
Include links in response to guide the client.

**JSON:**
```json
{
  "id": 1,
  "balance": 100,
  "links": [
    { "rel": "deposit", "method": "POST", "href": "/accounts/1/deposit" }
  ]
}
```

---

### Question 282: How do you design a bulk update endpoint?

**Answer:**
Accept array of operations.
**PATCH /items**
Body: `[ { id: 1, name: 'A' }, { id: 2, name: 'B' } ]`
Use **Transactions** to ensure all-or-nothing.

---

### Question 283: What is an idempotency key?

**Answer:**
A header `Idempotency-Key` sent by client (e.g., a UUID).
Server stores it. If received again, server returns previous response without re-processing (e.g., Charging credit card twice).

---

### Question 284: How do you handle file streaming over REST?

**Answer:**
Serve with `Content-Type: application/octet-stream` or specific type.
Pipe stream to response.

---

### Question 285: How would you implement rate-limiting per user/IP?

**Answer:**
Store counter in Redis with Expiry.
Key: `ratelimit:${ip}`.
Increment. If > Max, return 429.

---

### Question 286: How do you serialize large JSON responses efficiently?

**Answer:**
`JSON.stringify` blocks the loop for large objects.
**Solution:** Use streaming JSON stringify libraries like `JSONStream` or `bfj`.

---

### Question 287: What is the difference between `res.send()` and `res.json()`?

**Answer:**
*   **`res.json()`**: Sets `Content-Type: application/json` and stringifies input.
*   **`res.send()`**: Determines type automatically (String -> Text, Object -> JSON, Buffer -> Bin).

---

### Question 288: How do you handle circular JSON references?

**Answer:**
`JSON.stringify` throws error.
**Solution:** Use `flatted` library or `json-stringify-safe`.

---

### Question 289: How do you stream a large CSV file to a client?

**Answer:**
Use `fast-csv` or `csv-stringify` to create a transform stream. Pipe data from DB -> CSV Stream -> Response.

---

### Question 290: What is binary protocol handling in Node.js?

**Answer:**
Handling custom binary formats (not JSON/Text).
Use **Buffers** to read/write bytes, integers, floats. Protocol Buffers (protobuf) is a common schema for this.

---

### Question 291: How do you handle timezones in a Node.js app?

**Answer:**
Always store dates as **UTC** in Database and Server.
Convert to User's Timezone only at the presentation layer (Frontend) or using libraries like `moment-timezone` / `luxon`.

---

### Question 292: How would you handle scheduled tasks across timezones?

**Answer:**
Store logic: "User wants alert at 9AM".
Store timezone: "America/New_York".
Cron: Check every hour. If `CurrentTime_UTC` translated to `New_York` is 9AM, fire.

---

### Question 293: How do you use `luxon` or `dayjs` over `moment`?

**Answer:**
Moment is mutable and large (Legacy).
Luxon/Dayjs are immutable and modern.

**Luxon:**
```javascript
const { DateTime } = require("luxon");
DateTime.now().plus({ days: 1 }).toISO();
```

---

### Question 294: What is `process.uptime()` and when is it useful?

**Answer:**
Returns number of seconds Node has been running.
Useful for Health Checks (`/health`) to see if the server restarted recently.

---

### Question 295: How do you benchmark time-sensitive code?

**Answer:**
`process.hrtime.bigint()` (High Resolution Time).
Measures nanoseconds.

---

### Question 296: How do you allow plugins in your Node.js application?

**Answer:**
Design a hook system.
1.  Plugin registers function.
2.  App iterates hooks and calls functions.
**Fastify** has a great plugin architecture.

---

### Question 297: How do you handle large request payloads safely?

**Answer:**
Limit body size.
`app.use(express.json({ limit: '10kb' }));`
Prevent DoS attacks filling memory.

---

### Question 298: How do you serve multiple domains from a single Node.js server?

**Answer:**
Check `req.hostname` or `Host` header.
```javascript
const vhost = require('vhost');
app.use(vhost('api.example.com', apiApp));
app.use(vhost('www.example.com', webApp));
```

---

### Question 299: How do you enable graceful shutdown in a Node.js app?

**Answer:**
Listen for signal `SIGTERM` / `SIGINT`.
1.  Stop new requests (close server).
2.  Finish existing requests.
3.  Close DB connections.
4.  Exit.

**Code:**
```javascript
process.on('SIGTERM', () => {
  server.close(() => {
    db.disconnect();
    process.exit(0);
  });
});
```

---

### Question 300: How would you build a CLI tool with Node.js?

**Answer:**
1.  Add `bin` in `package.json`.
2.  Add Shebang `#!/usr/bin/env node` to file.
3.  Use libraries like `commander` or `yargs`.

**Code:**
```javascript
#!/usr/bin/env node
const yargs = require('yargs');

yargs.command('greet <name>', 'says hello', {}, (argv) => {
  console.log(`Hello ${argv.name}`);
}).argv;
```
