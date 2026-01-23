## 🟢 Advanced & API Development (Questions 51-100)

### Question 51: What is load balancing?

**Answer:**
Load balancing is the process of distributing incoming network traffic across multiple servers to ensure no single server becomes overwhelmed.

**Context in Node.js:**
Since Node.js is single-threaded, a single instance can use only one CPU core. To handle high traffic:
1.  **Nginx/HAProxy:** Distributes traffic to multiple Node.js processes/containers.
2.  **Node Cluster Module:** Distributes traffic to forked worker processes.

**Nginx Configuration Example (Conceptual):**
```nginx
upstream node_app {
    server 127.0.0.1:3000;
    server 127.0.0.1:3001;
    server 127.0.0.1:3002;
}

server {
    location / {
        proxy_pass http://node_app;
    }
}
```

---

### Question 52: How can you handle uncaught exceptions?

**Answer:**
Uncaught exceptions can crash the Node.js process. You can listen for the `uncaughtException` event on the `process` object.

**Important Note:** It is generally recommended to restart the process (using PM2 or Docker) after logging an uncaught exception, as the process state might be corrupted.

**Code:**
```javascript
process.on('uncaughtException', (err) => {
  console.error('UNCAUGHT EXCEPTION! Shutting down...');
  console.error(err.name, err.message);
  process.exit(1); // Exit to allow restart by process manager
});

// Example triggers
throw new Error('Something bad happened');
```

---

### Question 53: What is the difference between `try/catch` and `process.on('uncaughtException')`?

**Answer:**
*   **`try/catch`**: Used for handling synchronous errors in specific blocks of code. Ideally, you should handle errors locally where they occur.
*   **`process.on('uncaughtException')`**: Global safety net for errors that were *not* caught by any `try/catch`. It's a last resort mechanism.

**Limit:** `try/catch` cannot catch errors in asynchronous callbacks (unless using async/await).

```javascript
// try/catch fails here:
try {
  setTimeout(() => {
    throw new Error('Oops'); // Crash! Catch block won't see this.
  }, 100);
} catch (e) { ... }
```

---

### Question 54: How do you debug a Node.js application?

**Answer:**
1.  **Console:** `console.log()` (Basic).
2.  **Node Inspector:** Run with `--inspect` flag and attach Chrome DevTools.
3.  **IDE Debuggers:** VS Code built-in debugger (Attach to process).
4.  **Debug Module:** Detailed logging control.

**Using Inspector:**
```bash
node --inspect-brk app.js
# Open chrome://inspect in browser
```

**Using `debug` package:**
```javascript
// run with DEBUG=app:* node app.js
const debug = require('debug')('app:startup');
debug('Starting up...');
```

---

### Question 55: How does garbage collection work in Node.js?

**Answer:**
Node.js (V8) manages memory automatically.
1.  **New Space:** New objects are allocated here. Scavenge algorithm is used (fast).
2.  **Old Space:** Long-lived objects move here. Mark-Sweep-Compact algorithm is used (slower, runs less often).

**Mark-and-Sweep:**
1.  **Mark:** V8 traverses the object graph (roots -> references) and marks "reachable" objects (active).
2.  **Sweep:** Any object not marked is considered "garbage" and memory is reclaimed.

---

### Question 56: What are the best practices for securing a Node.js app?

**Answer:**
1.  **Helmet:** Set secure HTTP headers.
2.  **Rate Limiting:** Prevent DDoS / Brute Force.
3.  **Input Validation:** Sanitize all inputs (Joi, Validator).
4.  **No SQL Injection:** Use ORMs or parametrized queries.
5.  **Environment Variables:** Hide secrets.
6.  **Update Dependencies:** `npm audit`.

**Implementing Helmet:**
```javascript
const express = require('express');
const helmet = require('helmet');
const app = express();

app.use(helmet()); // Sets X-SS-Protection, X-Frame-Options, etc.
```

---

### Question 57: What are some common vulnerabilities in Node.js apps?

**Answer:**
1.  **Injection Attacks (SQL/NoSQL):** Unsanitized input query.
2.  **XSS (Cross-Site Scripting):** Rendering user input directly to HTML.
3.  **Prototype Pollution:** Modifying `Object.prototype`.
4.  **ReDoS (Regex Denial of Service):** Poorly written regex crashing the thread.
5.  **Insecure Dependencies:** Using packages with known flaws.

**ReDoS Example to Avoid:**
```javascript
// Bad Regex
const regex = /(a+)+b/;
// 'aaaaaaaaaaaaa...' inputs can freeze the event loop
```

---

### Question 58: What is the difference between `Buffer` and `Stream`?

**Answer:**
*   **Buffer:** A temporary memory location for raw binary data. It has a fixed size. It holds the *entire* data snippet in memory.
*   **Stream:** A sequence of data moving from one point to another. It handles data piece by piece (chunks).

**Analogy:**
*   **Buffer:** Waiting for a video to fully download before playing.
*   **Stream:** Watching a video on YouTube (buffering chunks as you watch).

```javascript
/* Buffer */
const buf = Buffer.from('hello');

/* Stream */
fs.createReadStream('movie.mp4');
```

---

### Question 59: What is the `zlib` module used for?

**Answer:**
`zlib` provides compression and decompression functionalities (Gzip, Deflate). It is essential for optimizing HTTP responses.

**Example: Compressing a file:**
```javascript
const zlib = require('zlib');
const fs = require('fs');
const gzip = zlib.createGzip();

const inp = fs.createReadStream('input.txt');
const out = fs.createWriteStream('input.txt.gz');

inp.pipe(gzip).pipe(out);
```

---

### Question 60: How does HTTP/2 differ from HTTP/1.1 in Node.js?

**Answer:**
HTTP/2 is binary, multiplexed (multiple requests over one connection), and supports Server Push. `http2` is a core module in Node.js.

*   **HTTP/1.1:** Text-based. Head-of-line blocking (one request at a time per connection).
*   **HTTP/2:** Binary framing. Low latency. Header compression.

**Creating HTTP/2 Server:**
```javascript
const http2 = require('http2');
const fs = require('fs');

const server = http2.createSecureServer({
  key: fs.readFileSync('server-key.pem'),
  cert: fs.readFileSync('server-cert.pem')
});

server.on('stream', (stream, headers) => {
  stream.respond({ ':status': 200 });
  stream.end('Hello HTTP/2');
});

server.listen(8443);
```

---

## 🟢 API Development (Questions 61-75)

### Question 61: What is Express.js?

**Answer:**
Express.js is a minimal and flexible Node.js web application framework. It simplifies server creation, routing, and middleware management compared to the raw `http` module.

**Key Features:**
*   Robust routing systems.
*   Middleware support.
*   Template engine integration.
*   Static file serving.

---

### Question 62: How do you create a basic Express server?

**Answer:**
Install `express`, require it, initialize the app, and listen on a port.

**Code:**
```javascript
const express = require('express');
const app = express();
const PORT = 3000;

app.get('/', (req, res) => {
  res.send('Hello World!');
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
```

---

### Question 63: How do you define routes in Express?

**Answer:**
Routes define how the app responds to client requests (URI + Method).

**Syntax:** `app.METHOD(PATH, HANDLER)`

**Examples:**
```javascript
// GET
app.get('/users', (req, res) => { /* fetch users */ });

// POST
app.post('/users', (req, res) => { /* create user */ });

// PUT (Update)
app.put('/users/:id', (req, res) => { /* update user */ });

// DELETE
app.delete('/users/:id', (req, res) => { /* delete user */ });
```

---

### Question 64: What is a route parameter?

**Answer:**
Route parameters are named URL segments used to capture values specified at their position in the URL. They are accessible via `req.params`.

**Code:**
```javascript
app.get('/users/:userId/books/:bookId', (req, res) => {
    // URL: /users/34/books/8989
  const { userId, bookId } = req.params;
  
  res.json({ user: userId, book: bookId });
});
```

---

### Question 65: What is the difference between `app.use()` and `app.get()`?

**Answer:**
*   **`app.use()`**: Applies middleware to **all** requests (GET, POST, etc.) that match the path. Used for global settings like parsing, logging, static files.
*   **`app.get()`**: Specifically handles **GET** requests for a specific path.

**Example:**
```javascript
// Runs for GET, POST, DELETE...
app.use('/api', (req, res, next) => {
  console.log('API Request received');
  next();
});

// Runs only for GET
app.get('/api/users', (req, res) => {
  res.send('Users');
});
```

---

### Question 66: How do you handle request body data in Express?

**Answer:**
Before Express 4.16+, you needed `body-parser`. Now, it uses built-in middleware for JSON and URL-encoded data.

**Setup:**
```javascript
const express = require('express');
const app = express();

// Parse JSON bodies (Content-Type: application/json)
app.use(express.json());

app.post('/user', (req, res) => {
  // Access data in req.body
  console.log(req.body.name); 
  res.send('User created');
});
```

---

### Question 67: What is body-parser and is it still needed?

**Answer:**
`body-parser` is a middleware that extracts the body portion of an incoming request stream and exposes it on `req.body`.

**Status:**
It is **no longer strictly needed** for standard JSON/UrlEncoded parsing in strict Express versions (v4.16.0+), as Express essentially bundled it.
```javascript
// Old
const bodyParser = require('body-parser');
app.use(bodyParser.json());

// New
app.use(express.json());
```
However, `body-parser` is still used for specific cases or older setups.

---

### Question 68: How do you implement CORS in Express?

**Answer:**
Use the `cors` middleware package.

**Basic Usage:**
```javascript
const cors = require('cors');
app.use(cors()); // Enable All CORS Requests
```

**Specific Origins:**
```javascript
const corsOptions = {
  origin: 'http://example.com', // Only allow this domain
  optionsSuccessStatus: 200
}
app.get('/products', cors(corsOptions), (req, res) => { ... });
```

---

### Question 69: How do you structure a RESTful API in Node.js?

**Answer:**
A common folder structure (MVC-like) promotes maintainability.

**Micro-structure:**
*   `app.js` (Entry point)
*   `routes/` (URL definitions)
*   `controllers/` (Request logic)
*   `models/` (Database Schema)
*   `middleware/` (Auth, Validation)

**Example Flow:**
1.  **Route:** `router.get('/users', userController.getAll)`
2.  **Controller:** Calls `User.find()` model.
3.  **Model:** Returns DB data.
4.  **Controller:** Sends JSON response.

---

### Question 70: How do you return status codes in Express?

**Answer:**
Use the `res.status()` method.

**Code:**
```javascript
app.post('/register', (req, res) => {
  if (!req.body.email) {
    return res.status(400).send('Email required'); // Bad Request
  }
  
  // Logic...
  
  res.status(201).json({ msg: 'Created' }); // Created
});
```

---

### Question 71: What are some status codes you commonly use?

**Answer:**
*   **200 OK:** Success.
*   **201 Created:** Resource successfully created.
*   **204 No Content:** Success but no body (e.g., Delete).
*   **400 Bad Request:** Client error (validation).
*   **401 Unauthorized:** Not logged in / invalid token.
*   **403 Forbidden:** Logged in, but permission denied.
*   **404 Not Found:** Resource doesn't exist.
*   **500 Internal Server Error:** Server crash/bug.

---

### Question 72: How do you handle authentication in a Node.js API?

**Answer:**
Common strategies:
1.  **Session-based:** Server stores session ID (Cookie).
2.  **Token-based (JWT):** Stateless. Server signs a token; Client sends it in Header.

**Passport.js** is a popular middleware for handling auth strategies (Local, Google, Facebook, JWT).

---

### Question 73: What is JWT (JSON Web Token)?

**Answer:**
JWT is an open standard (RFC 7519) for securely transmitting information between parties as a JSON object. It is compact and self-contained.

**Structure:** `Header.Payload.Signature`

1.  **Header:** Algorithm (HS256).
2.  **Payload:** Data (User ID, Role, Expiry).
3.  **Signature:** Verifies integrity using a secret key.

**Usage:** After login, server sends JWT. Client expects JWT in `Authorization: Bearer <token>` header for subsequent requests.

---

### Question 74: How do you protect routes with JWT in Express?

**Answer:**
Create a middleware function that checks the header and verifies the token.

**Code:**
```javascript
const jwt = require('jsonwebtoken');

function authenticateToken(req, res, next) {
  const authHeader = req.headers['authorization'];
  // Form: Bearer <TOKEN>
  const token = authHeader && authHeader.split(' ')[1];

  if (!token) return res.sendStatus(401);

  jwt.verify(token, process.env.ACCESS_TOKEN_SECRET, (err, user) => {
    if (err) return res.sendStatus(403);
    req.user = user; // Attach payload to request
    next();
  });
}

app.get('/protected', authenticateToken, (req, res) => {
  res.send('This is secret data');
});
```

---

### Question 75: How do you connect a Node.js app to MongoDB?

**Answer:**
Using the native `mongodb` driver or ODM like `mongoose` (recommended).

**Using Mongoose:**
```javascript
const mongoose = require('mongoose');

const uri = "mongodb://localhost:27017/mydb";

mongoose.connect(uri)
  .then(() => console.log('MongoDB Connected'))
  .catch(err => console.log(err));
```

---

## 🟢 Testing & Tools (Questions 76-85)

### Question 76: How do you test a Node.js application?

**Answer:**
Using testing frameworks and libraries.
*   **Unit Testing:** Testing individual functions.
*   **Integration Testing:** Testing API endpoints and DB interaction.
*   **E2E Testing:** Full flow.

**Common Stack:** Jest (Runner + Assertion), Mocha (Runner) + Chai (Assertion), Supertest (HTTP assertions).

---

### Question 77: What is Mocha?

**Answer:**
Mocha is a feature-rich JavaScript test framework running on Node.js and in the browser. It provides the test structure (`describe`, `it`) but requires an assertion library (like Chai) to verify values.

**Structure:**
```javascript
const assert = require('assert');

describe('Array', function() {
  describe('#indexOf()', function() {
    it('should return -1 when not present', function() {
      assert.equal([1, 2, 3].indexOf(4), -1);
    });
  });
});
```

---

### Question 78: What is Chai?

**Answer:**
Chai is a BDD/TDD assertion library that works with Mocha. It lets you write expressive tests.

**Styles:**
1.  **Should:** `foo.should.be.a('string');`
2.  **Expect:** `expect(foo).to.be.a('string');`
3.  **Assert:** `assert.typeOf(foo, 'string');`

**Example:**
```javascript
const expect = require('chai').expect;
expect(1 + 1).to.equal(2);
```

---

### Question 79: What is Jest and how is it different from Mocha?

**Answer:**
Jest is a complete testing framework developed by Facebook.

**Differences:**
*   **All-in-one:** Jest includes runner, assertion library, and mocking tools. Mocha usually needs Chai and Sinon.
*   **Snapshot Testing:** Jest has built-in snapshots.
*   **Parallelism:** Jest runs tests in parallel by default (faster).

**Jest Code:**
```javascript
test('adds 1 + 2 to equal 3', () => {
  expect(1 + 2).toBe(3);
});
```

---

### Question 80: What are mocks and stubs?

**Answer:**
Used to isolate unit tests by faking dependencies.

*   **Stub:** A fake function with pre-programmed behavior (e.g., force a database function to return "Success" without hitting the DB).
*   **Mock:** A fake object that logs calls to verify interaction (e.g., verify that `sendEmail` was called exactly once).

**Sinon Example:**
```javascript
// Stubbing a method
sinon.stub(database, 'save').returns(Promise.resolve({ id: 1 }));
```

---

### Question 81: How do you test async code?

**Answer:**
Modern frameworks support Promises and Async/Await.

**Jest/Mocha Example:**
```javascript
// Using Async/Await
it('should fetch data', async () => {
  const data = await fetchData();
  expect(data).toBe('peanut butter');
});

// Using Done callback (Older way)
it('should fetch data', (done) => {
  fetchData(data => {
    try {
      expect(data).toBe('peanut butter');
      done();
    } catch (error) { done(error); }
  });
});
```

---

### Question 82: How do you test an Express route?

**Answer:**
Use **Supertest**. It simulates HTTP requests to your Express app without starting the actual server on a port.

**Example:**
```javascript
const request = require('supertest');
const app = require('./app');

describe('GET /user', function() {
  it('responds with json', function(done) {
    request(app)
      .get('/user')
      .set('Accept', 'application/json')
      .expect('Content-Type', /json/)
      .expect(200, done);
  });
});
```

---

### Question 83: What is Supertest?

**Answer:**
Supertest is a library for testing Node.js HTTP servers. It provides a high-level abstraction for testing HTTP, allowing you to fluently verify status codes, headers, and body content.

---

### Question 84: What is the purpose of ESLint in Node.js development?

**Answer:**
ESLint is a static code analysis tool.

**Purpose:**
1.  **Find Problems:** Syntax errors, unused variables, unreachable code.
2.  **Enforce Style:** Indentation, quotes, semicolons (works with Prettier).
3.  **Best Practices:** Prevent using `var`, require `===`.

**Config:** Defined in `.eslintrc.json`.

---

### Question 85: What are some useful Node.js debugging tools?

**Answer:**
1.  **`console.log` / `console.table`**: Quick checks.
2.  **`node --inspect`**: Chrome DevTools.
3.  **VS Code Debugger**: Breakpoints, Variable Watch.
4.  **`nodemon`**: Auto-restart on change.
5.  **Postman**: Testing API endpoints.
6.  **ndb**: An improved specific debugging experience for Node by Google Chrome Labs.

---

## 🟢 Database Integration (Questions 86-92)

### Question 86: How do you connect to MongoDB using Mongoose?

**Answer:**
`mongoose.connect()` returns a promise.

**Code:**
```javascript
const mongoose = require('mongoose');

mongoose.connect('mongodb://localhost/test')
  .then(() => console.log('Connected!'))
  .catch(err => console.error('Connection error', err));
```

---

### Question 87: What is a schema in Mongoose?

**Answer:**
A Schema maps into a MongoDB collection and defines the shape of the documents within that collection.

**Example:**
```javascript
const { Schema } = mongoose;

const blogSchema = new Schema({
  title:  String, // String is shorthand for {type: String}
  author: String,
  body:   String,
  comments: [{ body: String, date: Date }],
  date: { type: Date, default: Date.now },
  hidden: Boolean,
  meta: {
    votes: Number,
    favs:  Number
  }
});
```

---

### Question 88: What are Mongoose models?

**Answer:**
Models are constructors compiled from Schemas. An instance of a model is a document. Models are responsible for creating and reading documents from the underlying MongoDB database.

**Creation:**
```javascript
const Blog = mongoose.model('Blog', blogSchema);

// Usage
const myPost = new Blog({ title: 'Intro to Node' });
await myPost.save();
```

---

### Question 89: How do you handle relationships in MongoDB using Mongoose?

**Answer:**
MongoDB is NoSQL (non-relational), but Mongoose simulates relationships using `populate()`. You store the `ObjectId` of one document in another.

**Schema:**
```javascript
const userSchema = new Schema({
  name: String,
  stories: [{ type: Schema.Types.ObjectId, ref: 'Story' }]
});

const storySchema = new Schema({
  author: { type: Schema.Types.ObjectId, ref: 'User' },
  title: String
});
```

**Querying (Populate):**
```javascript
Story.findOne({ title: 'My Story' })
  .populate('author') // Replaces ID with actual User document
  .exec((err, story) => {
    console.log('The author is %s', story.author.name);
  });
```

---

### Question 90: How do you handle transactions in MongoDB?

**Answer:**
MongoDB 4.0+ supports multi-document ACID transactions. In Mongoose, you use a Session.

**Code:**
```javascript
const session = await mongoose.startSession();
session.startTransaction();

try {
  await User.create([{ name: 'John' }], { session: session });
  await Account.create([{ money: 100 }], { session: session });

  await session.commitTransaction();
} catch (error) {
  await session.abortTransaction();
} finally {
  session.endSession();
}
```

---

### Question 91: What is the difference between `find()` and `findOne()`?

**Answer:**
*   **`find(criteria)`**: Returns an **array** of all documents matching the criteria. If none match, returns empty array `[]`.
*   **`findOne(criteria)`**: Returns the **first single object** matching the criteria. If none match, returns `null`.

---

### Question 92: How do you implement pagination in a Node.js API?

**Answer:**
Use `limit()` (items per page) and `skip()` (offset).

**Formula:**
`skip = (pageNumber - 1) * pageSize`

**Code:**
```javascript
// GET /products?page=2&limit=10
app.get('/products', async (req, res) => {
  const page = parseInt(req.query.page) || 1;
  const limit = parseInt(req.query.limit) || 10;
  const skipIndex = (page - 1) * limit;

  const products = await Product.find()
    .sort({ _id: 1 })
    .limit(limit)
    .skip(skipIndex)
    .exec();

  res.json(products);
});
```

---

## 🟢 DevOps & Deployment (Questions 93-100)

### Question 93: How do you deploy a Node.js app?

**Answer:**
1.  **Prepare:** Ensure `npm start` works. Set `PORT` via environment.
2.  **Platform:**
    *   **PaaS:** Heroku, Vercel, Render (git push deploy).
    *   **VPS:** AWS EC2, DigitalOcean.
3.  **Process Manager:** Use PM2 on VPS to keep app alive.
4.  **Reverse Proxy:** Set up Nginx to forward port 80/443 to Node app.

---

### Question 94: What is PM2 and how is it used?

**Answer:**
PM2 (Process Manager 2) is a production process manager for Node.js. It keeps applications alive forever, reloads them without downtime, and helps with logging and clustering.

**Commands:**
```bash
npm install pm2 -g
pm2 start app.js --name "my-app"
pm2 list
pm2 logs
pm2 restart my-app
```

---

### Question 95: How do you manage environment-specific configurations?

**Answer:**
Avoid hardcoding. Use:
1.  **Environment Variables (`.env`)**: Secrets and ports.
2.  **Config Module (`config` package)**: Loads files based on `NODE_ENV`.

**Structure:**
*   `config/default.json`
*   `config/production.json`
*   `config/development.json`

**Code:**
```javascript
const config = require('config');
const dbHost = config.get('customer.dbConfig.host');
```

---

### Question 96: How do you use Docker with Node.js?

**Answer:**
Create a `Dockerfile` to containerize the app.

**Simple Dockerfile:**
```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
EXPOSE 3000
CMD ["node", "app.js"]
```

**Build & Run:**
```bash
docker build -t my-node-app .
docker run -p 3000:3000 my-node-app
```

---

### Question 97: What is the benefit of using a `.env` file?

**Answer:**
*   **Security:** Keeps sensitive keys (API Keys, DB Passwords) out of version control (git).
*   **Flexibility:** Allows changing configuration (Port, DB Host) without changing code.
*   **Consistency:** Standard way to manage config across Dev, Stage, and Prod.

**Note:** Always add `.env` to `.gitignore`.

---

### Question 98: How do you monitor a production Node.js application?

**Answer:**
1.  **Logs:** Centralized logging (ELK Stack, CloudWatch).
2.  **Uptime:** Pingdom / UptimeRobot.
3.  **Performance (APM):** New Relic, Datadog.
4.  **PM2 Monitor:** `pm2 monit` gives basic CPU/RAM stats.

---

### Question 99: What cloud services can host Node.js apps?

**Answer:**
*   **PaaS (Easiest):** Heroku, Vercel (Next.js), Render, Railway, Google App Engine.
*   **IaaS (Control):** AWS EC2, DigitalOcean Droplets, Google Compute Engine.
*   **Serverless:** AWS Lambda, Azure Functions, Google Cloud Functions (pay per execution).
*   **Containers:** AWS ECS, Kubernetes (EKS/GKE).

---

### Question 100: How do you ensure zero-downtime deployment?

**Answer:**
Zero-downtime deployment means the service stays available to users while the new version is being deployed.

**Techniques:**
1.  **PM2 Cluster Reload:** `pm2 reload all` restarts workers one by one.
2.  **Blue-Green Deployment:** Run two identical environments. Route traffic to Green only after it's ready. Switch router.
3.  **Rolling Updates (Kubernetes):** Replaces pods gradually.
4.  **Load Balancer:** Remove instance from pool, update, add back.
