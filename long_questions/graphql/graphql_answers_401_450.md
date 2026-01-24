## 🟢 Multi-Client, DDD & Security (Questions 401-450)

### Question 401: How do you design a GraphQL schema that supports both web and mobile apps?

**Answer:**
Avoid "BFF per client". Use a unified schema.
Mobile: Uses specific queries (smaller payload).
Web: Uses richer queries.
Use `@skip` directives or Arguments (`image(size: "mobile")`) to tailor responses.

---

### Question 402: How do you version GraphQL queries across multiple clients?

**Answer:**
Version the *Operations*, not the Schema.
Persisted Query Registry has "V1" queries.
If you need to change the schema breakingly, ensure "V1" queries map to new resolvers or fallback logic.

---

### Question 403: What are named operations and how do they help in debugging across clients?

**Answer:**
`query HomeFeed { ... }`.
Logs show "HomeFeed" instead of `query`.
Metrics show "HomeFeed" is slow.
Helps pinpoint which App Screen is affecting backend performance.

---

### Question 404: How do you reduce over-fetching in mobile GraphQL queries?

**Answer:**
Strict query auditing.
Ensure mobile devs ask only for fields they render.
Use **Persisted Queries** to lock down the query shape so it can't grow accidentally.

---

### Question 405: How do you design for a React + iOS + Android + SPA setup with GraphQL?

**Answer:**
**Codegen is mandatory.**
React: Generates Hooks.
iOS: Generates Swift structs.
Android: Generates Kotlin data classes.
Ensures if Schema changes, all client builds fail at compile time (Safety).

---

### Question 406: How do you support multiple brands (white-label) via a single GraphQL API?

**Answer:**
Header `x-brand-id`.
Schema might have `interface Product` with `BrandAProduct` and `BrandBProduct` if logic differs wildly.
Or just data-level filtering if schema is same.

---

### Question 407: What are persisted queries and how do they benefit mobile clients?

**Answer:**
Upload query map at build time.
App sends hash `sha256(query)`.
Reduces request size from 2KB -> 64 bytes.
Crucial for flaky 3G/4G networks.

---

### Question 408: How do you analyze usage of GraphQL fields across multiple teams?

**Answer:**
Headers: `apollographql-client-name: "iOS-App"`, `apollographql-client-version: "1.2"`.
Studio report: "Field `bio` used 100% by Web, 0% by iOS".

---

### Question 409: How do you handle breaking changes for multiple frontend teams?

**Answer:**
1.  Mark `@deprecated`.
2.  Send push notification to dev teams (or Slack).
3.  Monitor Client Version usage.
4.  Wait until usage by "Old Versions" drops below threshold.

---

### Question 410: How do you allow experimentation for frontend teams using GraphQL?

**Answer:**
Schema Variant "Staging".
Or fields marked `@experimental`.
Or utilize the `extensions` part of the response for loose "meta" data during dev.

---

### Question 411: How do you build GraphQL clients for low-bandwidth environments?

**Answer:**
1.  Persisted Queries.
2.  Response Compression (Brotli).
3.  Code splitting queries (don't fetch tab 2 data until tab 2 clicked).
4.  Normalized Caching (minimize re-fetch).

---

### Question 412: How do you manage fragment co-location across multiple clients?

**Answer:**
Each Component file exports a Fragment.
Parent component interpolates it.
`AccountDetails` fragment used by both Web and Mobile logic (if shared codebase like React Native).

---

### Question 413: How do you allow dynamic UI configuration using GraphQL?

**Answer:**
**Server-Driven UI.**
Query returns a tree of components:
`layout { type: "HERO_BANNER", props: JSON }`.
Client just maps Type -> Component.

---

### Question 414: How do you group operations by product team in schema analytics?**

**Answer:**
Naming Convention or Annotations.
`query CheckoutTeam_SubmitOrder`.
Analytics regex matches `CheckoutTeam_*`.

---

### Question 415: How do you build feature flags using GraphQL query conditions?

**Answer:**
Pass flag state as Variable.
`query ($isBeta: Boolean!) { newFeature @include(if: $isBeta) }`.
Or return `null` from server if flag disabled.

---

### Question 416: How do you scale query delivery across multiple environments?

**Answer:**
Edge CDNs.
Route `/graphql` to nearest origin.
Cache public queries at the edge.

---

### Question 417: How do you provide access-controlled documentation for client teams?

**Answer:**
Introspection filtering.
"Partner A" gets Schema A (filtered).
"Internal Team" gets Schema Full.
Host docs behind Auth.

---

### Question 418: How do you support legacy clients that rely on old schema patterns?

**Answer:**
Snapshot the old schema logic.
Maybe run a "Legacy Subgraph" or handle it via "Legacy Resolvers" that map old fields to new DB structure.

---

### Question 419: What are GraphQL Operation Collections in Apollo?

**Answer:**
Registry of "Safe" queries.
You can "publish" queries from your codebase to the registry.
Server only runs queries found in the registry.

---

### Question 420: How do you detect unused operations across all consuming apps?

**Answer:**
Compare "Registered Operations" (from code) vs "Executed Operations" (from logs).
If registered but never executed, the code is dead.

---

### Question 421: How do you design a schema for a multi-tenant SaaS product?

**Answer:**
Global Context `tenantId`.
Type `Tenant { settings, users }`.
Root query `currentTenant { ... }`.

---

### Question 422: How do you model roles, permissions, and feature flags in GraphQL?

**Answer:**
`type User { permissions: [String] }`.
`type Query { featureFlags: JSON }`.
Frontend checks these to toggle UI.

---

### Question 423: How do you expose product configuration via GraphQL?

**Answer:**
Singleton types.
`type Configuration { currency, language, theme }`.
Fetched once on App Boot.

---

### Question 424: How do you model workflow states (like “Draft”, “Pending”, “Live”)?

**Answer:**
`enum Status`.
Valid transitions enforced by Mutations (`submitDraft`, `publish`).
Types might change: `DraftPost` (editable) vs `PublishedPost` (immutable).

---

### Question 425: How do you structure schema for a digital marketplace?

**Answer:**
Interfaces.
`interface Product`.
`DigitalProduct`, `PhysicalProduct`.
`search(query: String): [Product]`.

---

### Question 426: How do you model financial transactions with GraphQL?

**Answer:**
Precision is key.
`type Money { amount: Int, currency: String }` (Amount in cents).
Never use `Float` for money.

---

### Question 427: How do you design an audit log system with GraphQL queries?

**Answer:**
`type AuditEvent { id, user, action, targetType, targetId, changes }`.
Pagination is heavy here. Use strict cursor.

---

### Question 428: How do you model e-commerce order states in GraphQL?

**Answer:**
`type Order { status: OrderStatus, history: [OrderStatusHistory] }`.
Subscriptions for status changes are critical.

---

### Question 429: How do you expose shipping or logistics data via GraphQL?

**Answer:**
`type Shipment`.
`trackingUpdates` (List of events).
External API integration (FedEx) often needed in resolver.

---

### Question 430: How do you handle schema evolution for compliance (e.g. GDPR)?

**Answer:**
`@pii` directive.
If user requests "Forget Me", mutations scrub DB.
Query returns `null` or "Anonymized User".

---

### Question 431: How do you expose A/B test variants via GraphQL?

**Answer:**
`type Experiment { id, variant }`.
Resolvers interface with LaunchDarkly/Split.io SDKs.

---

### Question 432: How do you model timezones, locales, and formatting in schema?

**Answer:**
Return raw UTC `DateTime` (ISO).
Let client handle formatting.
Or fields `formattedDate(format: String): String`.

---

### Question 433: How do you model document uploads and previews in GraphQL?

**Answer:**
`type Document { url: String, previewUrl: String, mime: String }`.
Generate URLs on the fly (SAS/Signed).

---

### Question 434: How do you expose configuration-driven business rules via schema?

**Answer:**
`type Constraint { min: Int, max: Int, regex: String }`.
`type FieldConfig { name: String, constraints: Constraint }`.
Frontend uses this to validate forms dynamically.

---

### Question 435: How do you represent user-to-user relationships in GraphQL?

**Answer:**
Graph!
`User.friends` -> `[User]`.
`User.friendsOfFriends` -> `[User]`.
Watch out for complexity cost.

---

### Question 436: How do you structure content scheduling (like posts) in GraphQL?

**Answer:**
`publishAt` field.
Resolver filters out posts where `publishAt > Now` unless user is Admin.

---

### Question 437: How do you model event-driven workflows (e.g., triggers)?

**Answer:**
`type Workflow { trigger: EventType, steps: [Step] }`.
Polymorphic `Step` (EmailStep, WebhookStep).

---

### Question 438: How do you expose email, SMS, or push templates via GraphQL?

**Answer:**
`type Template { id, bodyTemplate, variables }`.
Mutation `previewTemplate(id: ID, vars: JSON): String` (Returns rendered HTML).

---

### Question 439: How do you represent pricing, billing tiers, and discounts?

**Answer:**
`type PricingBucket`.
`discounts: [Discount]`.
Calculated field `finalPrice`.

---

### Question 440: How do you handle batch updates to related entities in GraphQL?

**Answer:**
Transaction script in mutation.
`updateProject(input: { id: 1, name: "New", tasks: [ { id: 2, status: DONE } ] })`.
Deep Mutation.

---

### Question 441: How do you implement row-level access control in GraphQL?

**Answer:**
Context -> Service -> DB.
`db.posts.find({ userId: context.user.id })`.
The filter is implicit in the data access layer.

---

### Question 442: How do you manage access tokens in GraphQL operations?

**Answer:**
HTTPS Headers.
`Authorization: Bearer <token>`.
Context middleware validates it. Returns 401 if invalid.

---

### Question 443: What are security implications of introspection in production?

**Answer:**
Attacker learns all types, fields, and arguments.
Can find hidden Admin mutations or deprecated fields with known vulnerability.

---

### Question 444: How do you expose RBAC in GraphQL fields?

**Answer:**
Directives: `@authorized(roles: [ADMIN])`.
Library `graphql-auth-directives`.

---

### Question 445: How do you ensure a mutation is only allowed by certain roles?

**Answer:**
Guard logic.
```javascript
if (!user.hasRole('EDITOR')) throw new ForbiddenError();
```

---

### Question 446: How do you prevent IDOR (Insecure Direct Object Reference) in GraphQL?**

**Answer:**
Resolver validation.
`const doc = db.get(args.id);`
`if (doc.ownerId !== context.user.id) throw Error;`

---

### Question 447: What are common GraphQL-specific attack vectors?

**Answer:**
1.  **Batching:** Sending [Query, Query, Query...] to exhaust server.
2.  **Aliasing:** `query { a: user(id:1), b: user(id:1)... }` (1000 times).
3.  **Depth:** Cyclic queries.

---

### Question 448: How do you prevent enumeration attacks in GraphQL?**

**Answer:**
Standardize Errors.
Don't say "User not found". Say "Invalid credentials".
Rate limit.

---

### Question 449: How do you safely expose file URLs in a GraphQL schema?

**Answer:**
Signed URLs with short expiry (e.g. 15 mins).
`url` field resolver calls S3 `getSignedUrl`.

---

### Question 450: How do you integrate JWT authentication with resolvers?

**Answer:**
Decoded JWT payload is available in `context.user`.
Resolvers trust this context.
