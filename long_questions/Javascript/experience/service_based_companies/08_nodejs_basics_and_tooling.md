# 📘 08 — JavaScript in the Real World (Node.js Basics & Tooling)
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Node.js basics (modules, `require`, event-driven)
- npm: `package.json`, scripts, dependencies
- CommonJS vs ES Modules
- Environment variables
- Build tools overview (Webpack/Vite/Babel)
- REST API basics with Express

---

## ❓ Most Asked Questions

### Q1. What is Node.js and how does it differ from browser JS?

```javascript
// Node.js: JavaScript runtime built on V8 engine
// Runs JS on the server — no browser, no DOM

// Key differences:

// Browser JS:
// - Has: window, document, localStorage, fetch, DOM APIs
// - Limited: no file system access, no OS access, sandboxed

// Node.js:
// - Has: require(), process, Buffer, __dirname, __filename
// - Has: file system (fs), path, http, crypto, streams
// - No: window, document, localStorage

// Node.js globals (no window equivalent)
console.log(process.version);    // "v20.x.x"
console.log(process.platform);   // "linux" | "win32" | "darwin"
console.log(process.env.NODE_ENV); // "development" | "production"
console.log(__dirname);           // absolute path of current file's directory
console.log(__filename);          // absolute path of current file

// Node.js uses CommonJS by default:
const fs   = require('fs');
const path = require('path');

module.exports = { myFunction };

// Or ES Modules (with "type": "module" in package.json):
import fs from 'fs/promises';
export default function myFunction() {}

// Node.js is event-driven (same event loop as browser but different APIs)
const EventEmitter = require('events');
const emitter = new EventEmitter();

emitter.on('data', (chunk) => {
    console.log("Received:", chunk.length, "bytes");
});

emitter.emit('data', Buffer.from("Hello World"));
```

---

### Q2. Explain the `package.json` file and npm.

```json
{
    "name": "my-project",
    "version": "1.0.0",
    "description": "My awesome project",
    "main": "index.js",
    "scripts": {
        "start": "node src/index.js",
        "dev": "nodemon src/index.js",
        "build": "webpack --mode production",
        "test": "jest",
        "lint": "eslint src/**/*.js"
    },
    "dependencies": {
        "express": "^4.18.2",
        "mongoose": "^8.0.0",
        "dotenv": "^16.3.1"
    },
    "devDependencies": {
        "jest": "^29.0.0",
        "eslint": "^8.0.0",
        "nodemon": "^3.0.0"
    },
    "engines": {
        "node": ">=18.0.0"
    }
}
```

```javascript
// Key npm commands:
// npm init -y                  — create package.json
// npm install express          — install and add to dependencies
// npm install jest --save-dev  — install as devDependency only
// npm install                  — install all dependencies from package.json
// npm run dev                  — run 'dev' script
// npm test                     — run test script
// npm update                   — update packages to latest compatible version

// Version ranges:
// "^4.18.2" — compatible: 4.x.x (minor and patch updates)
// "~4.18.2" — patch only: 4.18.x
// "4.18.2"  — exact version only
// "*"        — any version (dangerous!)

// package-lock.json: exact versions for reproducible installs
// .gitignore: ALWAYS ignore node_modules/

// Useful npm scripts
// package.json
{
    "scripts": {
        "start":   "node dist/index.js",
        "dev":     "nodemon src/index.js --watch src",
        "test":    "jest --coverage",
        "test:watch": "jest --watch",
        "lint":    "eslint src --ext .js,.ts",
        "format":  "prettier --write src/**/*.js"
    }
}
```

---

### Q3. How do you read environment variables in Node.js?

```javascript
// dotenv: load variables from .env file into process.env
// Install: npm install dotenv

// .env file (NEVER commit to git!)
// DB_HOST=localhost
// DB_PORT=5432
// DB_NAME=mydb
// JWT_SECRET=supersecretkey123
// NODE_ENV=development

// Load at app startup (entry point):
require('dotenv').config(); // CommonJS
// or
import 'dotenv/config'; // ESM

// Access throughout app
const dbHost   = process.env.DB_HOST;     // "localhost"
const port     = parseInt(process.env.PORT) || 3000; // default value
const isDev    = process.env.NODE_ENV === 'development';
const jwtSecret = process.env.JWT_SECRET;

if (!jwtSecret) {
    throw new Error("JWT_SECRET is required but not set");
}

// ✅ Best practice: centralize config
// config/index.js
module.exports = {
    db: {
        host:     process.env.DB_HOST     || "localhost",
        port:     parseInt(process.env.DB_PORT) || 5432,
        name:     process.env.DB_NAME     || "mydb",
        password: process.env.DB_PASSWORD // no default for sensitive values
    },
    server: {
        port: parseInt(process.env.PORT) || 3000,
        env:  process.env.NODE_ENV || "development"
    },
    auth: {
        jwtSecret:  process.env.JWT_SECRET,
        jwtExpiry:  process.env.JWT_EXPIRY || "15m"
    }
};

// Use: const config = require('./config'); config.db.host
```

---

### Q4. Build a simple Express REST API.

```javascript
// Express.js: minimal Node.js web framework
// npm install express

const express = require('express');
const app = express();

// Middleware: runs before route handlers
app.use(express.json()); // parse JSON request bodies
app.use(express.urlencoded({ extended: true })); // parse form data

// Custom middleware: logging
app.use((req, res, next) => {
    console.log(`${req.method} ${req.path} - ${new Date().toISOString()}`);
    next(); // pass to next middleware/route
});

// In-memory store (use a DB in production)
let users = [
    { id: 1, name: "Alice", email: "alice@example.com" },
    { id: 2, name: "Bob",   email: "bob@example.com" }
];
let nextId = 3;

// GET all users
app.get('/api/users', (req, res) => {
    const { name } = req.query; // ?name=alice
    const result = name
        ? users.filter(u => u.name.toLowerCase().includes(name.toLowerCase()))
        : users;
    res.json(result);
});

// GET single user
app.get('/api/users/:id', (req, res) => {
    const user = users.find(u => u.id === parseInt(req.params.id));
    if (!user) return res.status(404).json({ error: "User not found" });
    res.json(user);
});

// POST create user
app.post('/api/users', (req, res) => {
    const { name, email } = req.body;

    if (!name || !email) {
        return res.status(400).json({ error: "name and email are required" });
    }

    const newUser = { id: nextId++, name, email };
    users.push(newUser);
    res.status(201).json(newUser);
});

// PUT update user
app.put('/api/users/:id', (req, res) => {
    const index = users.findIndex(u => u.id === parseInt(req.params.id));
    if (index === -1) return res.status(404).json({ error: "User not found" });
    
    users[index] = { ...users[index], ...req.body, id: users[index].id };
    res.json(users[index]);
});

// DELETE user
app.delete('/api/users/:id', (req, res) => {
    const index = users.findIndex(u => u.id === parseInt(req.params.id));
    if (index === -1) return res.status(404).json({ error: "User not found" });
    
    users.splice(index, 1);
    res.status(204).send(); // no content
});

// Global error handler (must have 4 params)
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ error: "Internal server error" });
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
```

---

### Q5. How does CommonJS (require) differ from ES Modules (import)?

```javascript
// CommonJS (CJS): older Node.js module system
// Synchronous, loads at runtime, can be inside conditions

// math.js (CJS)
const PI = 3.14159;
function add(a, b) { return a + b; }
module.exports = { PI, add };  // exports multiple things

// app.js (CJS)
const { PI, add } = require('./math');    // destructure
const utils = require('./utils');         // whole module object

// Dynamic require
const env = process.env.NODE_ENV;
const config = require(`./config/${env}`); // runtime path — works in CJS!

// ES Modules (ESM): modern, static analysis, tree-shakeable
// package.json: "type": "module" (or use .mjs extension)

// math.mjs (ESM)
export const PI = 3.14159;
export function add(a, b) { return a + b; }
export default class Calculator { /* ... */ }

// app.mjs (ESM)
import { PI, add } from './math.mjs';      // named imports
import Calculator from './math.mjs';       // default import
import * as Math from './math.mjs';        // namespace import

// Dynamic import (also works in ESM)
const { default: logger } = await import('./logger.mjs');

// Key differences:
// | Feature          | CommonJS              | ES Modules           |
// |------------------|-----------------------|----------------------|
// | Syntax           | require/exports       | import/export        |
// | Loading          | Synchronous           | Asynchronous         |
// | File extension   | .js                   | .mjs or "type":"module" |
// | Top-level await  | ❌                   | ✅                   |
// | Tree shaking     | ❌ (runtime)         | ✅ (static)          |
// | Conditional load | ✅                   | Only dynamic import  |
// | __dirname        | ✅ available         | ❌ (use import.meta.url) |
```

---

### Q6. What are streams in Node.js?

```javascript
// Streams: handle data piece by piece instead of loading all into memory
// Essential for large files, real-time data, HTTP responses

const fs = require('fs');

// Readable stream: source of data
const readStream = fs.createReadStream('./large-file.csv', { encoding: 'utf8' });

readStream.on('data', (chunk) => {
    console.log(`Received ${chunk.length} characters`);
    process(chunk);
});

readStream.on('end', () => console.log('Done reading'));
readStream.on('error', (err) => console.error('Error:', err));

// Writable stream: destination
const writeStream = fs.createWriteStream('./output.txt');
writeStream.write('Hello ');
writeStream.write('World\n');
writeStream.end();

// Pipe: connect readable → writable (handles backpressure automatically)
const input  = fs.createReadStream('./input.txt');
const output = fs.createWriteStream('./output.txt');

input.pipe(output); // copies file efficiently — no matter how large!

// Transform stream: read, transform, write
const { Transform } = require('stream');
const zlib = require('zlib');

// Compress file with streams
const compressStream = fs.createReadStream('./file.txt')
    .pipe(zlib.createGzip())      // transform: compress
    .pipe(fs.createWriteStream('./file.txt.gz')); // write compressed

// HTTP response is a stream!
const http = require('http');
const server = http.createServer((req, res) => {
    const fileStream = fs.createReadStream('./large-video.mp4');
    res.setHeader('Content-Type', 'video/mp4');
    fileStream.pipe(res); // stream video to client without loading all in memory
});

// Why streams matter:
// Without streams: 1GB file → 1GB in memory → crash or slow
// With streams: 1GB file → process 64KB at a time → efficient
```

---

### Q7. Explain the Node.js module system and circular dependencies.

```javascript
// Node.js modules: each file is its own module with private scope
// module.exports: what is shared; everything else is private

// counter.js
let count = 0;  // private to this module

function increment() { return ++count; }
function getCount()  { return count; }

module.exports = { increment, getCount };
// count variable is NOT exported — private

// Circular dependencies: A requires B, B requires A
// Node.js handles this by returning partial (incomplete) exports

// a.js
const b = require('./b');
console.log('a loaded, b.value:', b.value);
exports.value = 'A';

// b.js  
const a = require('./a');
console.log('b loaded, a.value:', a.value); // may be {} (partial export!)
exports.value = 'B';

// Avoid circular dependencies:
// 1. Extract shared code to a third module
// 2. Lazy require (inside function, not at top level)
// 3. Refactor to break the cycle

// Module caching: modules are cached after first require
const math1 = require('./math');
const math2 = require('./math');
math1 === math2; // true — same object cached in require cache!

// Force reload (rarely needed):
delete require.cache[require.resolve('./math')];
const freshMath = require('./math'); // new instance

// __dirname and __filename
console.log(__dirname);  // "/home/user/project/src"
console.log(__filename); // "/home/user/project/src/index.js"

// Build file paths safely
const configPath = require('path').join(__dirname, 'config', 'default.json');
const config = require(configPath);
```
