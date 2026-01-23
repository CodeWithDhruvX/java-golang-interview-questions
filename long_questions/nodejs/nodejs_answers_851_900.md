## 🟢 IaC, Niche & Trends (Questions 851-900)

### Question 851: How do you provision a Node.js service using Terraform?

**Answer:**
Terraform defines Infrastructure.
1.  Resource `aws_instance` or `aws_lambda_function`.
2.  Provisioner `remote-exec` to install Node/PM2.
3.  Better: Use Terraform to deploy a Docker Container (ECS/EKS) containing the Node app.

---

### Question 852: What is a golden AMI and how would you build one with Node pre-installed?

**Answer:**
A pre-baked VM image.
Use **Packer**.
Config installs Node.js, updates OS, installs global tools.
Terraform uses this AMI ID to launch instances faster (no startup install time).

---

### Question 853: How do you integrate infrastructure tagging into a Node deploy pipeline?

**Answer:**
In Pipeline (GitHub Actions).
Run Terraform command with variables.
`terraform apply -var="tags={Environment='Prod', Service='NodeAPI'}"`.

---

### Question 854: How do you use Node.js to orchestrate IaC commands like Terraform or Pulumi?

**Answer:**
Spawn child process.
`exec('terraform apply')`
Or use **Pulumi's Node.js SDK**. You write infra *in* TypeScript.
`const bucket = new aws.s3.Bucket(...)`.

---

### Question 855: How do you test infrastructure changes using Node?

**Answer:**
**Terratest** (Go) is standard.
In Node: **Pulumi** unit tests.
Or deploy ephemeral env, run `axios` probes against it to verify connectivity.

---

### Question 856: How do you handle plugin-based architectures in CLI tools?

**Answer:**
Search user home dir `~/.mycli/plugins`.
Require discovered modules.
Register commands `program.command(plugin.name)`.

---

### Question 857: What’s the difference between ESM and CJS CLI bootstrapping?

**Answer:**
*   **CJS:** `index.js` `#!/usr/bin/env node`. Runs fine.
*   **ESM:** Need to ensure file extension is `.mjs` or `package.json` type module. Also `__dirname` is missing (must polyfill).

---

### Question 858: How do you dynamically generate completions for a CLI tool?

**Answer:**
(See Q583). Use `tabtab` or `yargs`.
Script responds to `COMP_LINE` env var passing back array of suggestions.

---

### Question 859: How do you ensure CLI tools work across different shell environments?

**Answer:**
Use `cross-spawn`.
Avoid shell-specific syntax (`&&`, `|`, `>`) in child processes handling.
Use Node.js APIs (`fs`, `pipe`) instead of shell commands.

---

### Question 860: How do you provide interactive multi-step workflows in Node CLIs?

**Answer:**
`Inquirer` prompt loop.
State machine pattern.
`step1()` -> result -> `step2(result)`.

---

### Question 861: How do you build a Node.js app for embedded Linux devices?

**Answer:**
1.  Compile Node.js for target Arch (ARM).
2.  Or use static builds (node-static).
3.  Minimize size (exclude devDeps/docs).
4.  Use `pkg` to ship single binary.

---

### Question 862: What is a static build of Node.js and why might you need it?

**Answer:**
A Node binary linked statically (includes libs like openssl/libc inside).
Runs on Linux distro without installing specific shared libraries (glibc versions). Portability.

---

### Question 863: How do you bundle native binaries into a Node module?

**Answer:**
Place in `bin/` or `build/Release`.
At runtime, detect `process.platform` and `process.arch`.
`require` the correct file.

---

### Question 864: How do you strip debug symbols from a compiled module?

**Answer:**
Use `strip` command (Linux/Mac) on the `.node` file.
Reduces binary size significantly.

---

### Question 865: How do you build a hybrid WebAssembly + Node.js library?

**Answer:**
Write core logic in Rust. Compile to Wasm.
Write JS wrapper to load Wasm + provide nice API.
Publish as npm package.

---

### Question 866: How do you resolve dynamic import paths securely?

**Answer:**
Avoid `import(userInput)`.
Map Input -> Whitelist of paths.
`const map = { a: './a.js' }; await import(map[input])`.

---

### Question 867: How do you sandbox CJS modules within an ESM app?

**Answer:**
`createRequire`.
`const require = createRequire(import.meta.url)`.
`const cjs = require('./legacy.js')`.
Not a sandbox (security), just compatibility. True sandbox needs `vm`.

---

### Question 868: How do you simulate a virtual filesystem in ESM?

**Answer:**
Use **Import Loaders** (Experimental).
Hook `resolve` and `load`.
If url implies virtual, return source code string generated in-memory.

---

### Question 869: What are import assertions and when are they used?

**Answer:**
`import data from './data.json' assert { type: 'json' };`
Standard way to import non-JS resources in ESM. Prevents security issues where server sends JS with Main Type "application/json".

---

### Question 870: What’s the difference between static vs. dynamic export maps?

**Answer:**
Export maps in `package.json` are **Static**.
Conditions (`import`, `require`, `node`) determine resolution.
You cannot run Logic (JS function) to determine path during resolution.

---

### Question 871: What is contract testing and how do you use it with Node APIs?

**Answer:**
Verify Provider (API) meets Consumer expectations.
**Pact**: Consumer defines "I expect { id: 1 }".
Provider runs test replay to verify it returns that.

---

### Question 872: How do you enforce schema compatibility using tools like Pact?

**Answer:**
Pact Broker.
CI checks: "Can I deploy Provider v2?"
Broker says: "No, Consumer A expects field 'name' which you deleted."

---

### Question 873: How do you test custom TCP protocols in Node?

**Answer:**
`net.createConnection`.
Send raw Buffer bytes.
Assert response Buffer matches spec.

---

### Question 874: How do you ensure gRPC compatibility across versions?

**Answer:**
Run `buf` (linting tool) against Protobuf files.
Detects breaking changes (renaming fields, changing types).

---

### Question 875: How do you fuzz test a Node HTTP server?

**Answer:**
Use **AFL** or **fast-fuzz**.
Send random garbage headers/body.
Monitor for process crash.

---

### Question 876: How do you compensate for system clock drift in Node.js?

**Answer:**
Use NTP (Network Time Protocol) client to get Offset.
`const now = Date.now() + ntpOffset`.
Or trust the server infrastructure (AWS NTP is usually synced).

---

### Question 877: How do you synchronize time between Node instances?

**Answer:**
You don't.
Logic should not depend on exact wall-clock sync across machines.
Use **Logical Clocks** (Lamport) or DB versions.

---

### Question 878: How do you ensure consistency for time-based cache keys?

**Answer:**
Round to interval.
`key = 'stats_' + Math.floor(Date.now() / 60000)`. (Minute bucket).
All servers agree on the minute (roughly).

---

### Question 879: How do you safely use `Date.now()` in a distributed system?

**Answer:**
Only for duration measurement or logs.
Not for ordering events (Race conditions).
Use DB Auto-increment or Snowflakes (UUIDs) for ordering.

---

### Question 880: How do you detect and correct timestamp anomalies?

**Answer:**
If `received_timestamp < last_processed_timestamp`.
Likely clock skew or late message.
Handle by adjusting window or discarding if too old.

---

### Question 881: How do you generate accessible PDF reports with Node?

**Answer:**
`puppeteer`.
Render HTML with semantic tags / ARIA.
`page.pdf()`.
The resulting PDF preserves the tagging (usually) for screen readers.

---

### Question 882: How do you check a static site for accessibility from Node?

**Answer:**
**Pa11y** or **Axe-core**.
Run against URL. Returns violations list (Contrast, Alt text).

---

### Question 883: What Node libraries assist in screen-reader testing?

**Answer:**
`@accesslint/logger`.
Integrates into jest/puppeteer to emulate assistive tech tree.

---

### Question 884: How do you implement ARIA-aware markup using SSR from Node?

**Answer:**
Standard HTML templating.
Ensure dynamic widgets (`<Dropdown>`) output correct `aria-expanded` states in the initial HTML.

---

### Question 885: How can you test keyboard navigation with Node automation tools?

**Answer:**
Puppeteer `page.keyboard.press('Tab')`.
Check `document.activeElement` equals expected button.

---

### Question 886: What is the role of `undici` vs. `http` in Node’s future?

**Answer:**
`undici` is the new, faster HTTP/1.1 client written from scratch.
It powers the global `fetch`.
It might replace parts of legacy `http` client implementation eventually.

---

### Question 887: How does the Node.js release cadence align with npm’s evolution?

**Answer:**
Node bundles a specific npm version.
Sometimes Node updates npm in minor releases.
Major Node (20) brings Major npm (9).

---

### Question 888: How is Node.js adapting to the rise of edge runtimes?

**Answer:**
Standardizing APIs (`fetch`, `Request`, `Response`, `WebStreams`).
WinterCG collaboration to ensure "Write once, run anywhere".

---

### Question 889: What is the Node.js Build WG, and why is it relevant?

**Answer:**
Maintains the build infrastructure (Jenkins, machines) for all supported platforms (ARM, AIX, Windows, Linux).
Ensures Node works on your OS.

---

### Question 890: How do you track Node.js community security advisories?

**Answer:**
Subscribe to `nodejs-sec` mailing list.
Or watch GitHub Advisories database.

---

### Question 891: How do you simulate a dependency outage in Node?

**Answer:**
Change DNS/Hosts file to point `api.google.com` to localhost.
Start a mock server that returns 500s or hangs.

---

### Question 892: How do you simulate memory exhaustion in test environments?

**Answer:**
Run with low limit: `node --max-old-space-size=50 test.js`.
Allocate generic objects until crash.

---

### Question 893: What is a chaos monkey and how would you write one in Node?

**Answer:**
Middleware that randomly injects failure.
`if (Math.random() < 0.05) throw new Error("Chaos");`
`if (Math.random() < 0.05) await delay(10000);`

---

### Question 894: How do you create fault-tolerant retry queues?

**Answer:**
Dead Letter Queue (DLQ).
Retries = 0 -> Move to DLQ.
Alert human. Fix bug. Replay DLQ.

---

### Question 895: How do you automate incident response simulations?

**Answer:**
Script that triggers Chaos (Kill DB).
Verify that "Alert PagerDuty" callback fired in the Node monitoring system.

---

### Question 896: How do you verify license compatibility in npm dependencies?

**Answer:**
`license-checker`.
Scans `node_modules`.
Fails if `GPL` found in closed-source project.

---

### Question 897: How do you exclude GPL packages from a corporate Node project?

**Answer:**
CI step: `npx license-checker --exclude "GPL"`.

---

### Question 898: What tools help enforce open source license policies?

**Answer:**
**FOSSA** or **Snyk**.
Integrate with GitHub PRs.

---

### Question 899: How do you generate a license report for a Node app?

**Answer:**
`license-checker --csv > licenses.csv`.
Include in distribution bundle.

---

### Question 900: How do you comply with attribution requirements in Node CLI tools?

**Answer:**
`show-help` or `--credits` command prints list of all used open source libraries and authors (generated from license report).
