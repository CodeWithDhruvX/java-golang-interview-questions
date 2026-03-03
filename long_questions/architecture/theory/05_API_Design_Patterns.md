# 🔌 API Design Patterns — Questions 1–12

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Flipkart, Razorpay, PhonePe, Zepto — any company building APIs for internal or external use

---

### 1. What is REST and what are its core constraints?

"REST (Representational State Transfer) is an architectural style for distributed hypermedia systems, defined by Roy Fielding in his 2000 PhD thesis. It's not a protocol — it's a set of constraints.

The 6 constraints: **Uniform interface** (standard verbs GET/POST/PUT/DELETE, resources identified by URI), **Stateless** (each request contains all information; server holds no session), **Client-Server** (UI and data storage concerns are separated), **Cacheable** (responses must declare cacheability), **Layered System** (client can't tell if talking to origin or proxy), and **Code on Demand** (optional — server can send executable code like JavaScript).

Most 'REST APIs' in the wild are actually REST-ish. True REST includes HATEOAS — responses contain links to related resources. Almost nobody implements that part."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** TCS, Infosys (as theory) | Amazon, Flipkart (as API design context)

#### Indepth
REST levels (Richardson Maturity Model):
- **Level 0:** HTTP as a transport (SOAP-style). One URI, one method.
- **Level 1:** Resources. Multiple URIs (`/orders`, `/products`), but still using POST for everything.
- **Level 2:** HTTP verbs. GET for reads, POST for creates, PUT/PATCH for updates, DELETE for deletes. Most "REST" APIs are here.
- **Level 3:** HATEOAS. Responses contain hypermedia links. `GET /orders/123` returns order + `{ "_links": { "payment": "/payments?orderId=123" }}`. Nobody in production uses this.

**REST idempotency:**
- GET → Idempotent, Safe
- PUT → Idempotent, Not Safe (modifies)
- DELETE → Idempotent (deleting twice = same state)
- POST → Not idempotent (creates a new resource each time)
- PATCH → Depends on implementation

---

### 2. What is gRPC and how does it differ from REST?

"gRPC is a high-performance, open-source RPC (Remote Procedure Call) framework developed by Google. It uses **Protocol Buffers (protobuf)** as the interface definition language and serialization format, and **HTTP/2** as the transport.

The key differences: REST is text-based (JSON over HTTP/1.1), gRPC is binary (protobuf over HTTP/2). Binary is smaller and faster to parse. HTTP/2 supports multiplexing (multiple requests on one TCP connection), header compression, and server push.

Benchmark: gRPC is typically **5-10x faster** than REST/JSON for the same RPC. This matters for internal service-to-service calls that happen thousands of times per second. I use gRPC for inter-service communication and REST for external/public APIs where broad client compatibility is needed."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon, Flipkart, Google, Zepto, any company with heavy microservices

#### Indepth
gRPC communication patterns:
1. **Unary RPC:** One request, one response. Same as REST. `GetOrder(OrderID) → Order`
2. **Server streaming RPC:** One request, stream of responses. `SubscribeToOrderUpdates(OrderID) → stream OrderEvent`
3. **Client streaming RPC:** Stream of requests, one response. `UploadData(stream DataChunk) → UploadResult`
4. **Bidirectional streaming RPC:** Both sides stream simultaneously. Chat systems, real-time collaboration.

| Aspect | REST/JSON | gRPC/Protobuf |
|--------|-----------|---------------|
| Protocol | HTTP/1.1 | HTTP/2 |
| Format | JSON (text) | Protobuf (binary) |
| Payload size | Larger | 3-10x smaller |
| Speed | Moderate | 5-10x faster |
| Schema | OpenAPI (optional) | Protobuf (required, enforced) |
| Browser support | Native | Needs grpc-web proxy |
| Streaming | SSE/WebSocket (manual) | First-class |
| Code generation | Optional | Built-in (proto → client/server stubs) |

---

### 3. What is GraphQL and when would you choose it over REST?

"GraphQL is a query language for APIs where the **client specifies exactly what data it wants** — no more, no less. Developed by Facebook in 2012 and open-sourced in 2015.

The problem it solves: With REST, you often get **over-fetching** (endpoint returns 20 fields, client needs 3) or **under-fetching** (need 5 requests to different endpoints to get all data for one page). GraphQL lets the client send one query specifying exactly which fields and related entities to include.

I'd choose GraphQL for: public APIs with diverse clients (mobile, web, third-party all need different data), rapid product iteration (frontend can get new data without waiting for backend changes), or when a single endpoint should serve complex, nested data queries."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Flipkart, Amazon (AWS AppSync), Github, Shopify

#### Indepth
GraphQL vs REST:
| Aspect | REST | GraphQL |
|--------|------|---------|
| Data fetching | Fixed response shape | Client specifies fields |
| Endpoints | Many (one per resource) | One (`/graphql`) |
| Over-fetching | Common | Eliminated |
| Under-fetching | Requires multiple requests | Single query with nested fields |
| Caching | HTTP cache (GET → cacheable) | Complex (POST, non-cacheable by default) |
| Real-time | SSE or WebSocket (separate) | Subscriptions (built-in) |
| Schema | Optional (OpenAPI) | Required (SDL - schema definition language) |
| Learning curve | Low | Higher |

**GraphQL N+1 problem:** Querying a list of users and then their posts leads to N DB queries (1 for users, N for posts). Solution: **DataLoader** (batching and caching). This is a mandatory concept in any GraphQL performance discussion.

When NOT to use GraphQL: Simple CRUD APIs, internal service-to-service (use gRPC), file upload/download scenarios.

---

### 4. How do you version APIs?

"API versioning is about **making breaking changes without breaking existing clients**. The three main strategies: URL versioning (`/v1/orders`, `/v2/orders`), header versioning (`Accept: application/vnd.myapi.v2+json`), and query param versioning (`/orders?version=2`).

I prefer **URL versioning** for public APIs — it's explicit, easy to understand, easy to cache (different URLs can have different cache policies), and easy to route in a gateway or nginx.

My versioning policy: Avoid breaking changes whenever possible (add new fields, don't remove old ones). When a breaking change is unavoidable, introduce v2, run both in parallel, give clients a deprecation deadline (typically 6-12 months), and monitor traffic to know when it's safe to retire v1."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay (public payment API), Stripe, Amazon — companies running public APIs

#### Indepth
What constitutes a breaking change:
- Removing a field from a response
- Renaming a field
- Changing a field's data type (`string` → `int`)
- Changing an endpoint's path
- Changing the semantic meaning of a field
- Making an optional field required

Non-breaking changes (safe to do in existing version):
- Adding a new optional field to a response
- Adding a new optional request parameter
- Adding a new endpoint
- Adding new enum values (careful — if client validates enum values, this can break them)

**Semantic versioning for APIs:**
- **Major version (v1 → v2):** Breaking changes
- **Minor version (v1.1 → v1.2):** New features, backward compatible
- **Patch version:** Bug fixes, no API change

---

### 5. What is API idempotency and how do you implement it?

"An idempotent API is one where **making the same request multiple times produces the same result** as making it once. GET is always idempotent. PUT (set resource to this state) is idempotent. DELETE is idempotent (deleting twice = not found, same final state). POST (create resource) is NOT idempotent by default.

The problem: A client sends a payment POST request. Network times out. Was the payment processed? If the client retries, and the server already processed it, the customer gets double-charged.

Solution: **Idempotency keys**. The client generates a unique key per operation (`Idempotency-Key: uuid-abc123`), sends it in the header. The server stores the key + result. If the same key arrives again, return the cached result. Stripe uses exactly this pattern."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay, PhonePe, Stripe, CRED, PayU — any payment API

#### Indepth
Idempotency key implementation:
```
Client → POST /payments
         Headers: Idempotency-Key: uuid-abc123

Server:
1. Check idempotency_keys table for key 'uuid-abc123'
2. If found: return cached response (200 with existing payment data)
3. If not found:
   a. Process payment
   b. Store result in idempotency_keys with 24h TTL
   c. Return result

Table: idempotency_keys { key, response_code, response_body, expires_at }
```

Key design decisions:
- **TTL:** Idempotency keys should expire (24-48 hours). Clients shouldn't retry weeks later.
- **Scope:** Keys are scoped per API key / user to prevent cross-user collisions.
- **Atomicity:** Store the key + process the request in a single transaction to prevent race conditions.
- **Error cases:** If the first request failed (network timeout), should a retry attempt at the same operation again? Yes — only successful responses should be cached for idempotency, or you implement a state machine.

---

### 6. What is rate limiting and how do you implement it?

"Rate limiting controls **how many requests a client can make in a given time window** to protect the backend from abuse, overload, or DDoS.

Common strategies: **Fixed window** (100 requests per minute, resets every minute — simple but allows bursting at boundary), **Sliding window** (100 requests in any rolling 60-second window — more accurate), **Token bucket** (fill a bucket with N tokens per second, each request consumes one token — allows controlled bursting), and **Leaky bucket** (requests enter a queue and are processed at a constant rate).

Token bucket is my go-to for API rate limiting — it's intuitive and allows legitimate burst traffic (a client sends 50 requests, then pauses — the bucket refills, they can burst again later)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay, Amazon, any API platform company

#### Indepth
Rate limit implementation with Redis:
```
Token Bucket in Redis (lua script for atomicity):
key = "rate:{userId}"
tokens = GET key
if tokens < 1: return 429 Too Many Requests
DECR key (consume token)
EXPIRE key (if new key, set TTL)
```

Rate limit dimensions:
- **Per API key:** Total requests by an API consumer (external partners)
- **Per user ID:** User-level rate limiting (login attempts: 5/hour)
- **Per IP:** IP-level rate limiting (DDoS mitigation)
- **Per endpoint:** Critical endpoints (payment create: 10/min) vs information endpoints (product get: 1000/min)

Rate limit headers to return:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 47
X-RateLimit-Reset: 1678886400 (Unix timestamp of window reset)
Retry-After: 32 (only on 429 response)
```

---

### 7. What are webhooks and how do they differ from polling?

"Webhooks are **event callbacks over HTTP** — instead of you periodically checking 'did anything change?', the service calls your URL when something changes.

Polling: `GET /payment-status/{id}` every 5 seconds until status changes. 12 requests to get one event. Wasteful.

Webhook: Register `POST https://yourapp.com/payment-callback`. When payment status changes, Stripe sends a POST to your URL with the event details. Zero wasted requests. One request when it matters.

I use webhooks for: async payment status updates (Stripe, Razorpay webhooks), CI/CD pipeline triggers (GitHub webhooks trigger Jenkins), and any integration where you need to be notified of external events."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay, PayU, Stripe — payment integration discussions

#### Indepth
Webhook reliability challenges:
1. **Consumer unavailability:** What if your server is down when the webhook fires? → Retry with exponential backoff (Stripe retries over 72 hours)
2. **Ordering:** Webhooks may arrive out of order → Include event timestamp, don't rely on order
3. **Duplicate delivery:** Webhooks may be delivered multiple times → Consumer must be idempotent (check event_id)
4. **Security:** Anyone can POST to your webhook URL → Validate signature: `HMAC-SHA256(webhookSecret, requestBody)` must match the `X-Signature` header

Webhook vs SSE vs WebSocket:
| | Webhook | SSE | WebSocket |
|---|---------|-----|-----------|
| Direction | Server → Client (push) | Server → Client | Bidirectional |
| Connection | New HTTP for each event | Persistent HTTP | Persistent TCP |
| Protocol | HTTPS | HTTP/1.1 | ws:// |
| Use case | Async external events | Live feeds, notifications | Real-time chat, live collaboration |

---

### 8. What is an API contract and how do you enforce it?

"An API contract is the **formal specification of what a service's API looks like** — endpoints, request/response schemas, authentication requirements, error codes, and rate limits. It's the agreement between producer and consumer.

Tools: **OpenAPI/Swagger** for REST APIs (defines the contract as YAML/JSON), **Protocol Buffers** for gRPC (the `.proto` file is the contract), **AsyncAPI** for event-driven APIs.

Why it matters: With a contract, frontend and backend can develop in parallel against a mock. Changes to the contract are detected at build time (for protobuf) or via contract tests, not in production. Stripe's API is famously stable because they treat the API contract as a sacred commitment."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Product companies with multiple teams

#### Indepth
Contract-first vs code-first:
- **Contract-first:** Write the OpenAPI spec or protobuf definition first. Generate server stubs and client SDKs from the contract. Forces API design before implementation. Preferred.
- **Code-first:** Write the implementation, annotate with decorators, generate the spec. Faster to start but slippery — implementation details leak into API design.

**Consumer-driven contract testing (Pact):** Each consumer writes tests specifying what they expect from the provider. The provider runs these tests to ensure they don't break consumers. Pact Broker stores and versions these contracts. This is more granular than OpenAPI spec testing — it tests the actual interactions, not just the schema.

---

### 9. What is the difference between REST and RPC style APIs?

"REST is **resource-oriented** — you interact with nouns (resources). The client says what resource to act on and uses HTTP verbs to express the action. `POST /orders` (create an order resource), `GET /orders/{id}` (read it), `DELETE /orders/{id}` (delete it).

RPC (Remote Procedure Call) is **action-oriented** — you call functions by name. `POST /createOrder`, `POST /cancelOrder`, `POST /getOrderStatus`. The body contains parameters, the response is the result.

REST is better for CRUD-heavy, resource-centric APIs (e-commerce catalog, user profiles). RPC is more natural for complex operations, workflows, and commands (`POST /processRefund`, `POST /sendVerificationEmail`). gRPC is a modern, typed, binary RPC system."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Architecture design discussions

#### Indepth
REST purists vs pragmatists:
A strict REST API for 'lock user account' would be: `PATCH /users/{id}` with body `{ "status": "locked" }` — you're updating the user resource's status property.

An RPC-style API would be: `POST /users/{id}/lock` — you're invoking the 'lock' action directly.

The RPC style is often more readable for complex domain operations. This is why many APIs are REST-ish but use action-style URLs for operations that don't map cleanly to CRUD: `POST /payments/{id}/refund`, `POST /orders/{id}/cancel`, `POST /accounts/{id}/freeze`.

Google's API Design Guide recommends REST for most cases but explicitly acknowledges **custom methods** (RPC-style) for operations that don't map to CRUD: `POST /users/{id}:ban`, `POST /emails/{id}:send`.

---

### 10. How do you design for backward compatibility in APIs?

"Backward compatibility means **existing clients continue to work without changes** when you update your API. It's about adding without breaking.

Safe changes: Add new optional fields to responses (clients ignore unknown fields), add new optional request parameters, add new endpoints, add new enum values (careful), change internal implementation without changing the contract.

Breaking changes to avoid: Remove or rename fields, change field types, change the meaning/semantics of a field, make optional fields required, change URL structure.

My rule: Treat every field in your API response as a public commitment. Once a consumer depends on it, removing it is a breaking change. Stripe has maintained backward compatibility since 2011 — their 2011 API keys still work on their 2024 API."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Public API teams at Razorpay, Stripe, PayU

#### Indepth
Backward compatibility strategies:
1. **Tolerant reader:** Consumers should ignore unknown fields (don't use strict deserialization that breaks on unknown fields)
2. **Postel's Law:** Be conservative in what you send, liberal in what you accept
3. **Semantic versioning:** v1 → v2 only for breaking changes
4. **Deprecation lifecycle:** Mark field/endpoint as deprecated in docs → give 6-12 month sunset window → remove
5. **Feature flags:** New behavior only when `X-API-Behavior: new` header is present

**Expand-contract pattern for schema migrations:**
1. **Expand:** Add the new field alongside the old one
2. **Migrate:** Update all consumers to use the new field
3. **Contract:** Remove the old field (now a non-breaking change since no consumer uses it)

This is the only safe way to rename a field in a live API.

---

### 11. What is pagination and what patterns exist?

"Pagination prevents returning millions of records in a single API response. Three main patterns: **Offset pagination**, **Cursor-based pagination**, and **Keyset pagination**.

Offset pagination: `GET /orders?page=5&limit=20`. Simple but has the 'page drift' problem — if records are added/deleted between page requests, items shift and you get duplicates or misses. Also, `OFFSET 10000 LIMIT 20` performs poorly in SQL (scans and discards 10000 rows).

Cursor-based (preferred): `GET /orders?cursor=eyJpZCI6MTAwfQ&limit=20`. The cursor encodes the position in the dataset (often a base64-encoded ID or timestamp). Stable even as data changes. Used by Twitter, Instagram, and most modern APIs."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** API design rounds at any product company

#### Indepth
Cursor pagination implementation:
```
Response:
{
  "data": [...20 orders],
  "pagination": {
    "nextCursor": "eyJpZCI6MTIwfQ",  // base64({"id": 120})
    "hasMore": true
  }
}

Next request: GET /orders?cursor=eyJpZCI6MTIwfQ&limit=20
Server decodes cursor → WHERE id > 120 ORDER BY id LIMIT 20
```

Keyset pagination for sorted data: Uses a combination of sort key + ID:
```sql
WHERE (created_at, id) < (:cursor_created_at, :cursor_id)
ORDER BY created_at DESC, id DESC
LIMIT 20
```

This is how PostgreSQL achieves O(log n) pagination instead of O(offset) — uses a B-tree index directly.

---

### 12. What is API gateway vs service mesh — when to use each?

"API Gateway is an **entry point from the outside world** — it handles north-south traffic (client → services). It does auth, rate limiting, routing, SSL termination, and often response aggregation.

Service Mesh is **infrastructure for service-to-service communication** — east-west traffic (service → service). It handles mTLS, retries, circuit breaking, distributed tracing, and load balancing between internal services.

They're not alternatives — most mature architectures use both. API Gateway handles the external boundary, Service Mesh handles the internal network. API Gateway is often Nginx/Kong/AWS API Gateway. Service Mesh is often Istio/Linkerd + Envoy sidecars."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Senior architecture roles, platform engineering

#### Indepth
Traffic direction:
- **North-South (API Gateway):** External clients → API Gateway → Internal services
- **East-West (Service Mesh):** Internal Service A → Envoy sidecar → Envoy sidecar → Internal Service B

Responsibilities:
| Concern | API Gateway | Service Mesh |
|---------|------------|--------------|
| Client authentication (JWT) | ✅ | ❌ |
| Service-to-service mTLS | ❌ | ✅ |
| Rate limiting (external clients) | ✅ | ❌ |
| Retries (internal) | ❌ | ✅ |
| Circuit breaking (internal) | ❌ | ✅ |
| API versioning | ✅ | ❌ |
| Traffic splitting (A/B, canary) | ✅ (external) | ✅ (internal) |
| Distributed tracing | Both | ✅ (automatic) |
