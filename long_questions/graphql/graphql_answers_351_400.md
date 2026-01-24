## 🟢 Observability, DX & Strategy (Questions 351-400)

### Question 351: How do you use Apollo Rover CLI in CI pipelines?

**Answer:**
`rover subgraph check my-subgraph --schema schema.graphql --name my-subgraph`
Runs composition checks against the registry. Fails build if it breaks the Supergraph.

---

### Question 352: How do you test federated services locally?

**Answer:**
Use `rover dev`.
It starts a local Gateway (Router).
Watches local subgraph processes URLs.
Composes them in memory.

---

### Question 353: How do you generate Postman collections from a GraphQL schema?

**Answer:**
Export Schema as SDL.
In Postman: Import -> GraphQL Schema.
Postman creates a Collection with one Request per Query/Mutation type with default variables.

---

### Question 354: How do you validate that schema changes won’t break frontend apps?

**Answer:**
Use `graphql-inspector` or Apollo Studio Checks.
It compares the PR schema against the *live usage data* (traffic) from the last 7 days.
If you remove a field that was used 0 times, it passes. If 100 times, it fails.

---

### Question 355: What is a changelog generator for GraphQL schemas?

**Answer:**
Compares V1 and V2 SDL.
Returns:
*   `[+] type NewFeature`
*   `[-] field OldFeature (BREAKING)`
Useful for release notes.

---

### Question 356: How do you monitor schema drift?

**Answer:**
Drift: When code (resolvers) matches schema X, but the Registry thinks it is schema Y.
Check involves periodically introspecting the running server and diffing against Registry.

---

### Question 357: How do you automatically publish schema docs after every deploy?

**Answer:**
CI/CD Step.
`npx spectaql config.yml`.
Generates `public/index.html`.
Upload to S3 website bucket.

---

### Question 358: How do you manage secrets for GraphQL environments?

**Answer:**
Never in Schema.
Inject via `process.env`.
Context creation: `const secret = process.env.API_KEY`.
Pass to DataSources.

---

### Question 359: How do you simulate load tests for GraphQL queries?

**Answer:**
`k6` script.
Define query: `const query = 'query { users { name } }'`.
`http.post(url, JSON.stringify({ query }), { headers: ... })`.
Loop with Virtual Users.

---

### Question 360: How do you apply git tagging/versioning for GraphQL schema releases?

**Answer:**
Tag git repo.
Publish subgraph to registry with `git sha`.
Gateway config uses `refs` to know which SHA is live.

---

### Question 361: How do you log field-level resolver performance?

**Answer:**
Apollo Tracing.
Records `startOffset` and `duration` for `User.name`.
Aggregation reveals "User.name takes 50ns, but User.friends takes 200ms".

---

### Question 362: How do you track slow queries in GraphQL production logs?

**Answer:**
Middleware logger.
`if (duration > 1000ms) logger.warn("Slow Query", { query, variables, duration })`.

---

### Question 363: How do you enable tracing for Apollo Server?

**Answer:**
Install `ApolloServerPluginUsageReporting`.
Configure `sendTraces: true`.
Or use OpenTelemetry instrumentation for generic tracing.

---

### Question 364: What are Apollo Studio traces and how do you read them?

**Answer:**
Flame Graph.
Top bar = Total Request.
Nested bars = Resolvers.
Gap between bars = Node.js event loop waiting (or network overhead).

---

### Question 365: How do you send GraphQL metrics to Prometheus?

**Answer:**
`prom-client`.
Counter: `graphql_requests_total`.
Histogram: `graphql_request_duration_seconds`.
Labels: `operationName`.

---

### Question 366: How do you monitor subscriptions for dropped connections?

**Answer:**
Gauge: `websocket_active_connections`.
Start/Stop events in Subscription Server logs.
Alert if drops > 10% in 1 minute.

---

### Question 367: How do you detect schema misuse or underutilized fields?

**Answer:**
Usage Reporting.
Sort fields by "Request Count" (Ascending).
If count is 0, add `@deprecated`.

---

### Question 368: What is the role of `apollo-reporting` plugin?

**Answer:**
Offloads the heavy lifting of aggregating traces.
Buffers traces in memory. Sends to Apollo API in batches (asynchronously).
Prevents latency impact on user requests.

---

### Question 369: How do you create alerting for expensive queries?

**Answer:**
Calculate Query Cost (Complexity).
Log `cost` metric.
Alert: "P99 Query Cost > 5000".

---

### Question 370: How do you view latency for federated subgraphs?

**Answer:**
Gateway Traces.
Include "Service Time" vs "Gateway Overhead".
See if "User Service" is slow or the Gateway composition logic is slow.

---

### Question 371: What’s the best way to visualize GraphQL usage metrics?

**Answer:**
Grafana Dashboard.
Panels: "RPS", "Error Rate", "99th Percentile Latency", "Top 10 Operations".

---

### Question 372: How do you audit GraphQL requests for security threats?

**Answer:**
Log introspection queries (`__schema`).
Log queries with depth > 10.
Log validation failures (scanning for fields).

---

### Question 373: What is `apollo-server-plugin-landing-page-graphql-playground`?

**Answer:**
Configures the root URL `/`.
Renders playground HTML.
Use `ApolloServerPluginLandingPageDisabled` in production.

---

### Question 374: How do you attach correlation IDs across resolvers?

**Answer:**
Context.
`context.requestId = uuid()`.
Logger uses `context.requestId`.
Pass `x-request-id` header to downstream REST calls.

---

### Question 375: What log formats work best for GraphQL telemetry?

**Answer:**
**Structured JSON.**
Easier to query in Splunk/ELK.
`{ "level": "info", "gql": { "op": "Foo", "hash": "abc" } }`.

---

### Question 376: How do you build a query performance dashboard?

**Answer:**
Group by `OperationName`.
Avg Latency.
Request Count.
Error %.

---

### Question 377: How do you detect breaking changes via observability?

**Answer:**
Watch for `GRAPHQL_VALIDATION_FAILED` error code spikes.
Means clients are sending queries the server no longer understands.

---

### Question 378: How can observability inform schema refactoring?**

**Answer:**
Data-driven decisions.
"We want to split User.address. But 99% of queries ask for them together. Keep them together."

---

### Question 379: How do you set up distributed tracing in a GraphQL app?

**Answer:**
OpenTelemetry auto-instrumentation for `graphql`.
Propagate `traceparent` header.
View full trace in Jaeger/Zipkin.

---

### Question 380: How do you identify and fix redundant resolvers?

**Answer:**
Review trace.
If a resolver takes 0ms (sync return) and is called 1000 times, it might be blocking the event loop. Optimize shape.

---

### Question 381: How would you introduce GraphQL to a legacy backend team?

**Answer:**
Don't rewrite.
Build a "GraphQL Facade" (Gateway).
Wrap existing REST APIs.
Let Frontend enjoy GraphQL while Backend stays on REST temporarily.

---

### Question 382: What’s the ROI of migrating from REST to GraphQL?

**Answer:**
Front-end velocity increases (no blocked on backend changes).
Reduced bandwidth bills (Mobile).
Self-documenting API reduces communication overhead.

---

### Question 383: When should you avoid using GraphQL?

**Answer:**
1.  **Simple CRUD:** Overkill.
2.  **Binary Streams:** Video/File server (Use HTTP).
3.  **Server-to-Server:** gRPC is more efficient.

---

### Question 384: How do you pitch GraphQL to a non-technical stakeholder?

**Answer:**
"It makes the app faster on phones because we download less data. It makes developers faster because they don't have to write custom endpoints for every new button."

---

### Question 385: What types of projects benefit the most from GraphQL?

**Answer:**
Rich Clients (Dashboards, Social Feeds).
Multi-platform (Web + iOS + Android).
Graph-like data (Social networks).

---

### Question 386: What’s the difference in client usage pattern between REST and GraphQL?

**Answer:**
REST: Many parallel small requests on page load.
GraphQL: One or two large requests on page load.

---

### Question 387: How do you manage schema growth in a large enterprise?

**Answer:**
**Federation.**
Autonomous teams own their subgraphs.
Central "Platform Team" manages the Gateway and Governance rules.

---

### Question 388: What’s the best onboarding path for a new dev joining a GraphQL team?

**Answer:**
Explorer (GraphiQL).
"Run `query { me { name } }`".
"Read the docs tab".
"Add a field to User".

---

### Question 389: How do you balance performance vs flexibility in schema design?

**Answer:**
Prioritize Flexibility (Correct Semantics).
Fix Performance with Caching/Dataloaders.
Don't create `userWithPostsAndComments` root field just for speed (that's RPC).

---

### Question 390: What business risks come with exposing too much schema?

**Answer:**
Competitors scraping pricing/inventory algorithms.
Leakage of upcoming features (hidden flags in schema).

---

### Question 391: How do you sunset a deprecated GraphQL field safely?

**Answer:**
1.  Add `@deprecated`.
2.  Wait.
3.  Check Usage Logs (Studio).
4.  If Usage > 0, email consumers.
5.  If Usage == 0, delete.

---

### Question 392: What are common pitfalls in GraphQL project kickoffs?

**Answer:**
Leaking Database Schema 1:1.
Not using DataLoader early enough (Performance panic later).
Not implementing Authorization logic early.

---

### Question 393: How do you align frontend/backend teams on GraphQL practices?**

**Answer:**
Treat Schema changes like API contracts.
Frontend reviews the Schema PR.
Backend generates mocks for Frontend to start immediately.

---

### Question 394: How do you write GraphQL contracts in cross-team setups?

**Answer:**
Use IDL (SDL) as the source of truth.
Store in a shared repo or registry.

---

### Question 395: What are the signs of a well-maintained GraphQL API?

**Answer:**
Clear descriptions.
Consistent naming conventions.
Few deprecated fields (cleanup happens).
No "Any" or "JSON" types.

---

### Question 396: What KPIs should a GraphQL platform team measure?

**Answer:**
*   **Latency:** p95/p99.
*   **Availability:** Error rate.
*   **Adoption:** % of traffic using GraphQL vs REST.
*   **Breaking Changes:** Frequency (Lower is better).

---

### Question 397: What is the long-term maintenance cost of a GraphQL API?

**Answer:**
Higher Tooling cost (Caching, Registry, Monitoring setup).
Lower Iteration cost (Adding features is cheap).

---

### Question 398: How do you support 100+ frontend apps consuming the same GraphQL?

**Answer:**
**Client Registries.** Requires App ID header.
Persisted Queries (Strict mode).
Usage reporting per App ID.

---

### Question 399: How do you document and expose schema examples for 3rd-party devs?

**Answer:**
Public Developer Hub.
Embedded GraphiQL with *mocked* endpoint (so they can try without auth).
Tutorials showing common query patterns.

---

### Question 400: How do you evaluate GraphQL vendor solutions?

**Answer:**
Look for:
1.  **Federation Support:** (Apollo / Cosmo).
2.  **Security:** (WAF capabilities).
3.  **Observability:** (Traces/Metrics).
4.  **Lock-in:** Open Source core vs Proprietary.
