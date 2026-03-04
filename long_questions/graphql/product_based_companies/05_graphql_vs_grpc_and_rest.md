# 🟣 GraphQL vs gRPC vs REST — Trade-off Questions (Product-Based Companies)

This document covers the most important architectural comparison questions — "When would you pick GraphQL over gRPC or REST?" — which are a staple of system design interviews at Amazon, Flipkart, Razorpay, Swiggy, and similar companies. Targets 4–8 years experience.

---

### Q1: Compare GraphQL, REST, and gRPC. When would you choose each?

**Answer:**
This is the single most important architectural trade-off question in modern API design. There is no universal winner — each has a clear domain where it excels.

**Comprehensive Comparison:**

| Attribute | REST | GraphQL | gRPC |
|---|---|---|---|
| **Protocol** | HTTP/1.1 (JSON) | HTTP/1.1 (JSON) | HTTP/2 (Protobuf binary) |
| **Data format** | JSON | JSON | Protobuf (binary, schema-enforced) |
| **Performance** | Medium | Medium | **Highest** (binary, multiplexing) |
| **Payload size** | Larger JSON | Client-defined (smaller) | **Smallest** (binary) |
| **Type safety** | Optional (OpenAPI) | Schema-enforced | **Strict** (`.proto` file) |
| **Streaming** | SSE (one-way) | Subscriptions (WS) | **Native** bi-directional streaming |
| **Browser support** | ✅ Native | ✅ Native | ❌ Needs gRPC-Web proxy |
| **Over-fetching** | ❌ Common | ✅ Eliminated | ✅ Precise (by design) |
| **Under-fetching** | ❌ Common (multiple calls) | ✅ Single call | ✅ Single call |
| **Caching** | ✅ HTTP-native (CDN, ETags) | ❌ Hard (all POST) | ❌ No HTTP caching |
| **Learning curve** | Low | Medium | High |
| **Tooling maturity** | Very High | High | High (improving) |
| **Versioning** | URL versions (`/v1`,`/v2`) | Schema evolution | `.proto` field numbers |
| **Code generation** | Optional | Yes (Apollo Codegen) | **Required & automatic** |
| **Best for** | Public APIs, simple CRUD | Client-driven data, BFF | Internal microservices |

---

### Q2: Deep dive — When should you choose gRPC over GraphQL for microservice communication?

**Answer:**
This is a system design question where the right answer matters a lot.

**Choose gRPC for internal microservice-to-microservice calls when:**

**1. Performance is critical (high-throughput internal services):**
```
REST/GraphQL JSON payload:     ~500 bytes for a user object
gRPC Protobuf payload:         ~50 bytes for the same data (10x smaller)

At 100K RPM:
  JSON parsing overhead:       ~200ms CPU/request
  Protobuf parsing overhead:   ~20ms CPU/request
```

**2. You need bi-directional streaming:**
```proto
// gRPC supports 4 patterns REST/GraphQL can't match natively:
// 1. Unary (like REST)
rpc GetUser(UserRequest) returns (User);

// 2. Server-streaming (server sends multiple responses)
rpc WatchInventory(ProductId) returns (stream InventoryEvent);

// 3. Client-streaming (client sends multiple, server responds once)
rpc UploadBulkOrders(stream Order) returns (BulkUploadResult);

// 4. Bi-directional streaming (both sides stream simultaneously)
rpc TrackDelivery(stream Location) returns (stream DeliveryUpdate);
```

**3. You need strict contracts with code generation:**
```proto
// order.proto — Source of truth for ALL services
syntax = "proto3";

message Order {
  string id = 1;
  string user_id = 2;
  repeated OrderItem items = 3;
  double total = 5;  // Field 4 was deleted — safe, old number is reserved
  OrderStatus status = 6;
  google.protobuf.Timestamp created_at = 7;
}

enum OrderStatus {
  PENDING = 0;      // Default must be 0 in proto3
  CONFIRMED = 1;
  SHIPPED = 2;
  DELIVERED = 3;
}
```
- If you add/remove fields, the `.proto` enforces backward compatibility automatically via field numbers
- Auto-generates client code in Go, Java, Node.js, Python simultaneously

**4. Polyglot microservices:**
```bash
# Generate clients for all languages from one .proto file
protoc --go_out=./gen/go \
       --java_out=./gen/java \
       --js_out=./gen/js \
       order.proto
# All services use the same strongly-typed contract
```

---

### Q3: When should you choose GraphQL over gRPC?

**Answer:**

**Choose GraphQL when:**

**1. Client is a browser/mobile app:**
- gRPC-Web requires a proxy (Envoy/gRPC-Gateway) — adds complexity
- GraphQL is native HTTP/JSON — works everywhere

**2. Multiple client types need different data shapes:**
```graphql
# Mobile query — minimal data (bandwidth sensitive)
query MobileProductCard {
  product(id: "123") {
    name
    thumbnailUrl
    price
  }
}

# Web query — rich data (full detail page)
query WebProductDetail {
  product(id: "123") {
    name
    description
    images { url width height }
    variants { sku size color stock }
    reviews(first: 5) { rating text author { name } }
    relatedProducts(limit: 4) { name price thumbnailUrl }
  }
}
```
With gRPC, you'd need two separate RPC methods and two different proto messages. With GraphQL, one schema handles both.

**3. Rapid schema evolution with many teams:**
- Adding optional fields to GraphQL is non-breaking and instant
- gRPC field numbers are permanent — poor ergonomics for fast-moving teams

**4. You need GraphiQL / self-documenting API:**
- GraphQL introspection gives free documentation + playground
- gRPC requires separate tools (grpcurl, Postman gRPC, BloomRPC)

---

### Q4: What is the "GraphQL + gRPC" hybrid architecture? Why do major companies use it?

**Answer:**
The most production-proven architecture at scale is: **GraphQL at the gateway layer, gRPC for internal service-to-service communication.**

```
Internet
    ↓
[GraphQL Gateway / Apollo Router]
    ↓                ↓                 ↓
[User Service]  [Product Service]  [Order Service]
  (gRPC)           (gRPC)             (gRPC)
    ↓                ↓                 ↓
[User DB]       [Product DB]       [Order DB]
```

**Why this is ideal:**
- **GraphQL Gateway** handles: client flexibility, aggregation, authentication, field-level auth, caching
- **gRPC internally** handles: high-speed service calls, streaming, strict typed contracts, polyglot teams

**Implementation — Gateway calling gRPC services:**
```javascript
// Apollo Subgraph → calls internal gRPC services
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const packageDef = protoLoader.loadSync('./proto/user.proto');
const proto = grpc.loadPackageDefinition(packageDef);
const userClient = new proto.UserService('user-service:50051', grpc.credentials.createInsecure());

// GraphQL resolver that calls gRPC
const resolvers = {
  Query: {
    user: (_, { id }) =>
      new Promise((resolve, reject) => {
        userClient.GetUser({ id }, (err, user) => {
          if (err) reject(err);
          else resolve(user);
        });
      })
  }
};
```

**Real-world examples:**
- **Netflix**: GraphQL BFF layer + gRPC for backend services
- **Uber**: GraphQL gateway (Marketplace) + gRPC for microservices
- **Shopify**: GraphQL Storefront API + internal gRPC services

---

### Q5: How does REST compare to GraphQL for building partner/public APIs?

**Answer:**
For public/partner APIs (third-party developers consume your API), **REST still has advantages over GraphQL**.

**REST advantages for public APIs:**

**1. HTTP caching works out of the box:**
```bash
# REST response — CDN can cache this GET request
GET /api/v1/products/123
→ Cache-Control: public, max-age=3600
→ ETag: "abc123"

# GraphQL — POST requests are NOT cached by CDN
POST /graphql
{ "query": "{ product(id: \"123\") { name price } }" }
→ Cache-Control: no-store (unsafe for POST)
```

**2. Simpler for simple CRUD — lower developer experience barrier:**
```bash
# REST — anyone using curl knows this immediately
curl https://api.stripe.com/v1/charges \
     -H "Authorization: Bearer sk_test_..." \
     -d amount=2000 \
     -d currency=usd

# GraphQL — requires knowing schema, GraphQL syntax
```

**3. HTTP verbs have semantic meaning:**
```
GET    /products       → Read (safe, idempotent) → CDN cacheable
POST   /products       → Create
PUT    /products/123   → Full Update
PATCH  /products/123   → Partial Update
DELETE /products/123   → Delete
```
GraphQL merges all writes into `Mutation` — loses HTTP verb semantics.

**4. Webhooks and event notification are REST-native:**
```
Stripe, Razorpay, GitHub all use REST webhooks:
POST https://your-server.com/webhooks/stripe
←  { "event": "payment.succeeded", "data": {...} }
```

**Decision Summary:**

| Use Case | Winner |
|---|---|
| Internal microservice calls (high throughput) | **gRPC** |
| Real-time streaming (telemetry, IoT, chat) | **gRPC** |
| Mobile/web client data fetching (BFF) | **GraphQL** |
| Multiple clients needing different data shapes | **GraphQL** |
| Public API for 3rd-party developers | **REST** (or REST + GraphQL) |
| Simple CRUD API, startup with small team | **REST** |
| API with CDN caching requirements | **REST** |
| Strict schema contracts, polyglot teams | **gRPC** |

---

### Q6: How would you migrate a large REST API to GraphQL without breaking existing clients?

**Answer:**
This is a **strangler fig pattern** migration — incrementally introduce GraphQL alongside REST.

**Phase 1 — GraphQL over REST (wrapper layer):**
```javascript
// GraphQL resolvers call existing REST endpoints
const resolvers = {
  Query: {
    user: async (_, { id }) => {
      const response = await fetch(`http://rest-api.internal/users/${id}`);
      return response.json();
    },
    products: async (_, { limit = 20 }) => {
      const response = await fetch(`http://rest-api.internal/products?limit=${limit}`);
      const data = await response.json();
      return data.items;  // Adapt REST response shape to GraphQL
    }
  }
};
// Risk: adds a latency hop. Benefit: zero changes to existing REST clients.
```

**Phase 2 — Migrate resolvers to access services directly:**
```javascript
// Remove REST hop — resolver now queries DB directly or calls gRPC
const resolvers = {
  Query: {
    user: async (_, { id }, context) => context.db.users.findById(id),
    products: (_, { limit }, context) => context.db.products.findAll({ limit })
  }
};
// REST endpoints still work for legacy clients via old routes
```

**Phase 3 — Deprecate REST endpoints:**
```yaml
# Sunset header (RFC 8594) — inform clients of upcoming REST deprecation
HTTP/1.1 200 OK
Sunset: Sat, 01 Jan 2026 00:00:00 GMT
Deprecation: Mon, 01 Jun 2025 00:00:00 GMT
Link: <https://graphql.api.com/graphql>; rel="alternate"
```

**Key migration checklist:**
- [ ] Set up GraphQL gateway alongside REST (running in parallel)
- [ ] Implement all critical REST endpoints as GraphQL queries/mutations
- [ ] Use HTTP response headers to communicate REST deprecation dates
- [ ] Monitor which REST endpoints are still receiving traffic
- [ ] Provide migration guides + GraphQL Playground for developers
- [ ] Remove a REST endpoint only when traffic to it drops to zero

---

*Prepared for system design rounds, API design discussions, and architectural decision questions at product companies.*
