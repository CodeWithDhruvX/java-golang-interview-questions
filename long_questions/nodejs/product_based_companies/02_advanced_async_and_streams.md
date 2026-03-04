# 🌊 02 — Advanced Async & Streams
> **Most Asked in Product-Based Companies** | 🌊 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Deep understanding of Streams (Readable, Writable, Duplex, Transform)
- `pipe()` vs `pipeline()` vs Async Iterators
- Backpressure in Streams
- Advanced Promises (`Promise.allSettled`, `Promise.any`)
- Worker Threads for parallel CPU tasks
- `EventEmitter` anti-patterns and performance

---

## ❓ Frequently Asked Questions

### Q1. What are Streams in Node.js? Why are they highly recommended for processing large files?

**Answer:**
Streams are a set of objects that let you read data from a source or write data to a destination continuously, chunk by chunk, instead of loading the entire dataset into memory at once.

**Why they are highly recommended:**
If you try to read a 2GB file using `fs.readFile()`, Node.js will attempt to load all 2GB into RAM (V8's heap limit is normally ~1.4GB on 64-bit systems), causing a crash (`Allocation failed - JavaScript heap out of memory`).

Using Streams (`fs.createReadStream()`), the file is processed in small buffers (default 64KB chunks). The memory footprint remains incredibly low, making the application highly scalable.

---

### Q2. Describe the four types of Streams in Node.js.

**Answer:**
1. **Readable:** Streams from which data can be read (e.g., `fs.createReadStream()`, `http.IncomingMessage`).
2. **Writable:** Streams to which data can be written (e.g., `fs.createWriteStream()`, `http.ServerResponse`).
3. **Duplex:** Streams that are both Readable and Writable (e.g., a `net.Socket` for TCP connections).
4. **Transform:** A special type of Duplex stream where the output is computed based on the input (e.g., `zlib.createGzip()` to compress data on the fly).

---

### Q3. What is Backpressure in streams, and how do you handle it?

**Answer:**
**Backpressure** occurs when the *Readable* stream is reading data faster than the *Writable* stream can consume and process it. If left unhandled, the data will buffer up in RAM, potentially crashing the server (OOM).

When you write to a stream `stream.write(chunk)`, it returns a boolean.
- `true`: The internal buffer is fine. Keep writing.
- `false`: The internal buffer is full. Stop writing and wait.

**How to handle it:**
1. **The `pipe()` method:** Automatically manages backpressure under the hood. It pauses the Readable stream when the Writable stream is saturated, and resumes it when the `drain` event is emitted.
   ```javascript
   readableStream.pipe(writableStream);
   ```
2. **The `pipeline()` utility:** Even better than `pipe()`, `pipeline` correctly handles forwarding errors across multiple piped streams, preventing memory leaks if a stream fails mid-way.
   ```javascript
   const { pipeline } = require('stream');
   pipeline(readableStream, transformStream, writableStream, (err) => {
       if (err) console.error('Pipeline failed', err);
       else console.log('Pipeline succeeded');
   });
   ```

---

### Q4. How do you implement Async Iterators to consume streams?

**Answer:**
In modern Node.js, Async Iterators (`for await...of`) are the preferred way to consume readable streams over event listeners (`.on('data')`). They automatically handle backpressure and make code look synchronous.

```javascript
const fs = require('fs');

async function processFile(filePath) {
    const readable = fs.createReadStream(filePath, { encoding: 'utf8' });
    
    // Using for await...of
    for await (const chunk of readable) {
        console.log(`Received chunk of size ${chunk.length}`);
        // Backpressure is handled automatically by the JS engine here!
    }
    console.log('Finished processing.');
}
```

---

### Q5. When should you use `worker_threads`? Provide a basic example.

**Answer:**
`worker_threads` should be used when you need to perform CPU-intensive tasks (e.g., image resizing, heavy JSON parsing, cryptography) that would otherwise block the Event Loop and starve other asynchronous I/O requests.

*Workers are NOT for I/O operations (libuv handles that better).*

**Example:**
*main.js*
```javascript
const { Worker } = require('worker_threads');

function runService(workerData) {
    return new Promise((resolve, reject) => {
        const worker = new Worker('./worker.js', { workerData });
        worker.on('message', resolve); // Listen for result
        worker.on('error', reject);
        worker.on('exit', (code) => {
            if (code !== 0) reject(new Error(`Worker stopped with exit code ${code}`));
        });
    });
}

runService('Alice').then(result => console.log(result));
```

*worker.js*
```javascript
const { workerData, parentPort } = require('worker_threads');

// Heavy CPU blocking task
let sum = 0;
for (let i = 0; i < 1_000_000_000; i++) {
    sum += i;
}

parentPort.postMessage(`Hello ${workerData}, computation done: ${sum}`);
```
