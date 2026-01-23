## 🟢 Release, Collaboration & CLI (Questions 651-700)

### Question 651: How do you test a Node app against multiple Node versions?

**Answer:**
Use CI matrix.
**GitHub Actions:**
```yaml
strategy:
  matrix:
    node-version: [14.x, 16.x, 18.x]
steps:
  - uses: actions/setup-node@v3
    with:
      node-version: ${{ matrix.node-version }}
```

---

### Question 652: What is semantic-release and how does it automate versioning?

**Answer:**
A tool that automates the release workflow:
1.  Analyzes commits (Conventional Commits).
2.  Determines next version (Major/Minor/Patch).
3.  Generates Changelog.
4.  Publishes to npm & GitHub releases.

---

### Question 653: How do you enforce changelog updates via CI in a Node.js repo?

**Answer:**
Check if `CHANGELOG.md` was modified in the PR.
Or better, use tools like `semantic-release` or `changesets` which generate changelogs purely from commit data/metadata, removing manual need.

---

### Question 654: What is a canary release and how do you implement it in Node?

**Answer:**
Deploy new version to a small % of users.
**Impl:**
1.  Deploy v2 to new structure.
2.  Update Load Balancer to send 5% traffic to v2.
3.  Monitor error rate.
4.  Scale up.

---

### Question 655: How do you manage breaking changes in a public npm package?

**Answer:**
1.  Bump Major Version (SemVer).
2.  Provide Migration Guide.
3.  Keep old API working with Deprecation warning for one cycle.
4.  Release a "codemod" script to auto-update user code if possible.

---

### Question 656: How do you build a desktop app using Electron and Node.js?

**Answer:**
Electron combines Chromium (Frontend) and Node.js (Backend).
`main.js` (Node) creates Window.
`renderer.js` (Web) runs UI.
They communicate via IPC.

---

### Question 657: How do you create a plugin system for a design tool using Node?

**Answer:**
Run plugins in a child process or sandbox.
Expose a specific API over IPC.
Example: Figma uses WebAssembly. In Node context, you'd use `vm` or worker threads.

---

### Question 658: What are the security implications of bundling Node in desktop apps?

**Answer:**
If XSS exists in the Renderer, and `nodeIntegration: true` is set, the attacker can execute system commands (`fs.unlink`, `child_process`).
**Fix:** Disable `nodeIntegration`, enable `contextIsolation`. Use `preload` scripts.

---

### Question 659: How do you expose Node.js functionality in a VSCode extension?

**Answer:**
VSCode runs extensions in a Node.js host.
You have full access to `fs`, `child_process`, etc.
Define functionality in `activate()` function.

---

### Question 660: How do you build a test runner or build tool using Node?

**Answer:**
1.  CLI Entry point (`bin`).
2.  Glob file search pattern (find test files).
3.  Require files.
4.  Execute functions.
5.  Catch errors and pretty print.

---

### Question 661: How do you process NDJSON (newline-delimited JSON) in streams?

**Answer:**
NDJSON = One JSON object per line.
Use `split2` module to split stream by newline.
Parse each chunk.

```javascript
fs.createReadStream('data.ndjson')
  .pipe(split2(JSON.parse))
  .on('data', obj => { /* handle object */ });
```

---

### Question 662: How do you perform ETL operations in real time using Node.js?

**Answer:**
Streams/Pipeline.
Source (DB/File) -> Transform Stream (Process/Filter) -> Dest (DW/File).
Node.js is excellent for this due to backpressure support.

---

### Question 663: What are transform streams and how do they enable live data manipulation?

**Answer:**
Stream that reads input, modifies it, and produces output.
Example: Gzip, Encryption, CSV-to-JSON.
`_transform(chunk, enc, cb)` method.

---

### Question 664: How do you handle schema evolution in streamed JSON data?

**Answer:**
In Transform stream, check version field.
Map Old Schema -> New Schema object before passing down downstream.

---

### Question 665: How do you deduplicate and batch data in-flight?

**Answer:**
**Batching Transform Stream:**
Buffer chunks in an array. When array length > 100 or timeout, `push` the array and clear buffer.
**Dedup:** Keep a Set of IDs (if small) or use Redis Bloom Filter.

---

### Question 666: How do you enforce code ownership in a Node.js monorepo?

**Answer:**
`CODEOWNERS` file in GitHub/GitLab.
`/packages/api @team/backend`
`/packages/ui @team/frontend`

---

### Question 667: How do you prevent dependency drift in long-lived projects?

**Answer:**
1.  Use `package-lock.json`.
2.  Use RenovateBot / Dependabot to auto-update.
3.  Run `npm audit` in CI.

---

### Question 668: How do you create internal reusable dev tools using Node.js?

**Answer:**
Create a CLI package `@company/cli`.
Include scripts for linting, scaffolding components, deploying.
Install in every project `devDependencies`.

---

### Question 669: What’s the difference between shared configs vs. shared code in monorepos?

**Answer:**
*   **Shared Config:** Eslint, TSConfig. (Dev tools).
*   **Shared Code:** Utilities, UI Components. (Runtime dependency).

---

### Question 670: How do you handle multi-team deployment coordination in Node-based stacks?

**Answer:**
Contract Testing (Pact).
Feature Flags.
Release Trains (Deploy every Tuesday).

---

### Question 671: How do you implement CRDTs in a Node.js app?

**Answer:**
Use **Yjs** or **Automerge**.
Node.js server acts as the "connecting point" to merge updates and broadcast to clients.
CRDTs ensure eventually consistent data without conflict errors.

---

### Question 672: What is Operational Transformation and how can it be supported in Node?

**Answer:**
Older tech (Google Docs style). Relies on central server to transform operations (Insert 'A' at 0) against others.
Library: `shareb`, `ot.js`.

---

### Question 673: How do you sync cursor position across clients using Node?

**Answer:**
Ephemeral data (don't save to DB usually, maybe Redis).
WebSocket broadcasts `{ userId, x, y }` to all other connected clients in the room.

---

### Question 674: How do you broadcast changes only to affected clients in real time?

**Answer:**
Rooms/Channels.
User edits "Doc A".
Server sends update to `room:doc_A`.
Users viewing Doc B (different room) don't get traffic.

---

### Question 675: How do you implement offline editing with sync in a Node backend?

**Answer:**
Client (PWA) stores edits in IndexedDB.
When online, sends queue of ops to Node.
Node merges strategies (Last Write Wins or CRDT).
Sends back latest state.

---

### Question 676: How do you parse subcommands and nested arguments?

**Answer:**
Libraries like `yargs` support `.command()`.
`cli users create --name=John`.
`cli users delete --id=1`.

---

### Question 677: How do you write an interactive install wizard in a CLI tool?

**Answer:**
`inquirer.js`.
Ask series of questions.
Use answers to generate config files or run install scripts.

---

### Question 678: How do you auto-generate CLI docs from source code?

**Answer:**
If using `commander` or `yargs`, they auto-generate help text (`--help`).
For Markdown docs, plug the help output into a generator script.

---

### Question 679: What are hidden CLI flags, and when should you use them?

**Answer:**
Flags not shown in `--help`.
`--debug-mode` or `--danger-reset`.
Used for dev debugging or advanced operations.

---

### Question 680: How do you ensure consistency across multiple CLI binaries in a monorepo?

**Answer:**
Share the core logic library.
Binaries are just thin wrappers.
Share standard versioning and help text generation logic.

---

### Question 681: How do you implement time zone–aware job scheduling?

**Answer:**
Store schedule as `{ cron: "0 9 * * *", tz: "Asia/Tokyo" }`.
Use `cron` library's timezone support.
`new CronJob(time, fn, null, true, 'Asia/Tokyo')`.

---

### Question 682: What’s the difference between cron and interval-based scheduling?

**Answer:**
*   **Cron:** Wall-clock time ("Every day at 8am"). Handles Daylight Savings.
*   **Interval:** Relative time ("Every 24 hours"). Drifts, ignores Daylight Savings.

---

### Question 683: How do you queue a job for delayed execution in distributed Node apps?

**Answer:**
Redis Sorted Sets (ZSET).
Score = Timestamp to run.
Poller checks ZRANGEBYSCORE.
Or utilize Delayed Exchanges in RabbitMQ.

---

### Question 684: How do you recover from missed schedules due to downtime?

**Answer:**
Store "Last Run" time in DB.
On Startup, check if Current Time > Last Run + Interval.
If yes, run immediately (Catch up).

---

### Question 685: How do you simulate cron execution during test runs?

**Answer:**
Manually trigger the job function.
Or mock the clock to jump to the configured time.

---

### Question 686: How do you detect file handle leaks in production?

**Answer:**
`lsof -p PID` on Linux.
Monitor Open File Descriptors metric.
If growing linearly, you have a leak.

---

### Question 687: How do you identify memory bloat from large object graphs?

**Answer:**
Heap Snapshot.
Look at **Dominators** view.
Find root object holding reference to huge tree (e.g., Global Cache Map).

---

### Question 688: How do you find async functions that never resolve?

**Answer:**
Very hard.
Some tools (`wtfnode`) list active handles/promises on exit.
Add timeouts to all async ops.

---

### Question 689: What is heapdump and how is it analyzed?

**Answer:**
Module to write V8 snapshot to disk programmatically.
Analyze by loading `.heapsnapshot` into Chrome DevTools.

---

### Question 690: How do you inspect deadlocks caused by incorrect `await` usage?

**Answer:**
Node doesn't have "threads" to deadlock, but logic can stall.
Review code for `await` loops or dependencies (A awaits B, B awaits A - usually happens with Events/Promises).
Use logging/timeouts.

---

### Question 691: How do you generate interactive API documentation from JSDoc?

**Answer:**
Tools like `better-docs` or `jsdoc-to-markdown`.
Or convert JSDoc comments to OpenAPI spec.

---

### Question 692: What’s the difference between a readme and usage guides?

**Answer:**
*   **Readme:** "What is this? How to install? Quickstart."
*   **Guide:** "How to do X? Tutorials. Deep dive."

---

### Question 693: How do you embed live code examples in Node.js documentation?

**Answer:**
Runkit (for npm).
Or verify code snippets in docs by running them as tests (`dtslint` or custom parsers).

---

### Question 694: What is the role of OpenAPI in dev onboarding?

**Answer:**
Single source of truth.
New dev can try endpoints via Swagger UI without writing code.
Clients (SDKs) can be auto-generated.

---

### Question 695: How do you document internal-only APIs in a public Node module?

**Answer:**
Mark JSDoc with `@private` or `@internal`.
Doc generators usually exclude these tags.

---

### Question 696: How do you create a single executable binary from a Node app?

**Answer:**
Use `pkg` or `vercel/pkg` (fork).
It bundles Node.js runtime + Your Code + Assets into one `.exe` / linux binary.

---

### Question 697: What’s the difference between `pkg`, `nexe`, and `vercel/pkg`?

**Answer:**
*   **nexe:** Older, compiles Node from source (slow).
*   **pkg:** Fetches pre-built binaries, appends code. Fast.
*   **vercel/pkg:** Maintained fork of pkg.

---

### Question 698: How do you bundle a Node app for distribution without exposing source code?

**Answer:**
Binary packaging (`pkg`) hides the source somewhat (it's inside the binary).
Use `v8-compile-cache` or `bytenode` to distribute **Bytecode** instead of Source. (Harder to reverse engineer).

---

### Question 699: How do you embed assets (HTML, CSS, etc.) into a Node.js binary?

**Answer:**
`pkg` detects `path.join(__dirname, 'file')` and bundles it manifest.
At runtime, it patches `fs` to read from the binary's virtual filesystem.

---

### Question 700: What are the limitations of freezing a Node.js project into a binary?

**Answer:**
1.  **Native Modules:** `.node` files must be physically present next to exe (cannot be bundled easily).
2.  **Dynamic Require:** `require(variable)` fails because pkg cannot predict what to bundle.
