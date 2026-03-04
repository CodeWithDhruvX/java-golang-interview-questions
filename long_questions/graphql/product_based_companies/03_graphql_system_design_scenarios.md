# 🟣 GraphQL System Design & Real-World Scenarios — Interview Questions (Product-Based Companies)

This document covers system design and scenario-based GraphQL questions for senior/staff engineer interviews at product companies. Targets 5+ years experience, focusing on architecture trade-offs, debugging, and design decisions.

---

### Q1: Design a GraphQL API for a large e-commerce platform like Flipkart/Amazon. Walk through your schema design decisions.

**Answer:**
This is a system design + schema design problem. The interviewer wants to see holistic thinking.

**Requirements (clarify first):**
- Clients: Web, iOS, Android, Partner APIs
- Features: Product catalog, user accounts, cart, orders, reviews, search
- NFRs: 10M DAU, < 100ms P99 latency, 99.99% uptime

**High-Level Architecture Decision:**
Use **Apollo Federation** — multiple teams own different domains.

```
Client Layer
├── Web App (Apollo Client)
├── iOS/Android (Apollo iOS/Kotlin)
└── Partner API (REST gateway on top of GraphQL)
          ↓
Apollo Router (Supergraph Gateway)
├── Catalog Subgraph (Product/Category/Inventory)
├── User Subgraph (Auth/Profile/Addresses)
├── Cart Subgraph (Cart/Wishlist)
├── Order Subgraph (Orders/Returns/Tracking)
├── Review Subgraph (Reviews/Ratings/Q&A)
└── Search Subgraph (Elasticsearch-backed)
```

**Core Schema Design (Catalog Subgraph):**
```graphql
type Product @key(fields: "id") {
  id: ID!
  slug: String!
  title: String!
  description: String!
  basePrice: Money!
  discountedPrice: Money
  discountPercentage: Int

  category: Category!
  brand: Brand!
  variants: [ProductVariant!]!
  images: [ProductImage!]!

  # Aggregated — resolved from Review subgraph
  averageRating: Float
  reviewCount: Int

  isAvailable: Boolean!
  stockCount: Int    # Null if not shown publicly
  estimatedDelivery: DeliveryEstimate
}

type Money {
  amount: Float!
  currency: String!    # "INR", "USD"
  formatted: String!   # "₹1,499"
}

type ProductVariant {
  id: ID!
  sku: String!
  size: String
  color: String
  price: Money!
  stock: Int!
  images: [ProductImage!]
}

type ProductConnection {
  edges: [ProductEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
  appliedFilters: [AppliedFilter!]!    # Echo back applied filters
  availableFilters: [FilterOption!]!  # Server-driven filter UI
}
```

**Key Design Decisions to discuss:**
1. **Money as a scalar vs type:** Use a `Money` type (not Float) to handle currencies and formatting.
2. **Pagination:** Always cursor-based for product lists (stable with sorted data).
3. **Availability:** Return `isAvailable` boolean, not just `stock` integer (business logic stays in service).
4. **Federation:** Product is `@key(fields: "id")` so Orders/Reviews can reference it.
5. **Server-driven UI:** Return `availableFilters` from the API so client doesn't hardcode filter UI.

---

### Q2: Your GraphQL API is slow — response times have increased from 50ms to 800ms after release. How do you debug and resolve it?

**Answer:**
This is a systematic debugging question. Walk through the process methodically.

**Step 1: Instrument and Identify the slow resolver**
```javascript
// Add timing to all resolvers via a plugin
const timingPlugin = {
  requestDidStart: () => ({
    executionDidStart: () => ({
      willResolveField({ info }) {
        const start = Date.now();
        return () => {
          const elapsed = Date.now() - start;
          if (elapsed > 100) {   // Warn on resolvers > 100ms
            console.warn(`Slow resolver: ${info.parentType.name}.${info.fieldName}: ${elapsed}ms`);
          }
        };
      }
    })
  })
};
```

```graphql
# Use Apollo Studio field-level latency breakdown or enable inline traces
query {
  products {       # ← Check which field is slow
    title
    reviews { rating }   # ← Is this causing N+1?
    inventory { stock }  # ← Is this calling a slow service?
  }
}
```

**Step 2: Common Culprits and Fixes**

| Symptom | Likely Cause | Fix |
|---|---|---|
| Latency scales with list size | N+1 problem | DataLoader / `@BatchMapping` |
| Constant high latency | DB query without index | Add DB index, check `EXPLAIN` |
| Timeout on specific mutation | External API call (email/payment) | Make async, use job queue |
| High latency for unauthenticated users | Context function hits DB every request | Cache user lookup in context |
| Slow only for deep nested queries | No DataLoader + deep nesting | DataLoader + depth limiting |
| Memory pressure | Large query responses | Pagination, field limiting |

**Step 3: Check for missing DataLoaders:**
```javascript
// Before fix — N+1
const resolvers = {
  Product: {
    category: (product) => db.findCategory(product.categoryId)  // Called 100x for 100 products!
  }
};

// After fix — BatchMapping
const resolvers = {
  Product: {
    category: (products) => {  // DataLoader batches all 100 into 1 query
      const ids = products.map(p => p.categoryId);
      return categoryLoader.loadMany(ids);
    }
  }
};
```

**Step 4: Database Query Analysis:**
```sql
-- Check for missing indexes
EXPLAIN ANALYZE SELECT * FROM products WHERE category_id = '123';
-- Add index if needed
CREATE INDEX CONCURRENTLY idx_products_category_id ON products(category_id);
```

**Step 5: Caching::**
```javascript
// Add Redis cache for frequently accessed, slowly-changing data
const resolvers = {
  Query: {
    popularProducts: async (_, __, context) => {
      const cacheKey = 'popular_products';
      const cached = await redis.get(cacheKey);
      if (cached) return JSON.parse(cached);

      const products = await db.getPopularProducts();
      await redis.setex(cacheKey, 300, JSON.stringify(products));  // TTL: 5 minutes
      return products;
    }
  }
};
```

---

### Q3: How would you implement real-time features (live scores, notifications, collaborative editing) using GraphQL Subscriptions vs Polling vs SSE?

**Answer:**
Choose the right real-time mechanism based on the use case.

**Comparison:**

| Feature | GraphQL Polling | GraphQL Subscriptions (WebSocket) | Server-Sent Events (SSE) |
|---|---|---|---|
| Protocol | HTTP | WebSocket (stateful) | HTTP (one-way stream) |
| Direction | Client initiates every N seconds | Bi-directional | Server → Client only |
| Connection overhead | High (new request each time) | Low (persistent connection) | Low (persistent) |
| Firewall friendly | ✅ Yes | ❌ Some firewalls block WS | ✅ Yes |
| Scalability | easy (stateless) | Hard (sticky sessions needed) | Medium |
| Best for | Low-frequency (refresh rates > 30s) | Real-time chat, live collaboration | Notifications, live feeds |

**Use Case Mapping:**
- **Live Cricket Score:** WebSocket Subscription (< 1s updates)
- **Order Status Update:** SSE or polling every 30s (infrequent)
- **Dashboard Analytics Refresh:** Polling every 60s
- **Collaborative Document Editing:** WebSocket Subscription + CRDT

**Subscription Implementation (live chat):**
```graphql
type Subscription {
  messageSent(channelId: ID!): Message!
  userTyping(channelId: ID!): TypingIndicator!
}
```

```javascript
// Resolver
const resolvers = {
  Subscription: {
    messageSent: {
      // Security: validate user can access channel
      subscribe: async (_, { channelId }, context) => {
        if (!context.user) throw new AuthenticationError('Not authenticated');
        const canAccess = await channelService.canUserAccess(context.user.id, channelId);
        if (!canAccess) throw new ForbiddenError('Cannot access channel');

        return pubsub.asyncIterator(`MESSAGE.${channelId}`);
      },
      resolve: (payload) => payload
    }
  },

  Mutation: {
    sendMessage: async (_, { channelId, text }, context) => {
      const message = await messageService.create({ channelId, text, userId: context.user.id });
      // Trigger subscription for all subscribers of this channel
      pubsub.publish(`MESSAGE.${channelId}`, message);
      return message;
    }
  }
};
```

**Client-Side Subscription:**
```javascript
const MESSAGES_SUBSCRIPTION = gql`
  subscription OnMessage($channelId: ID!) {
    messageSent(channelId: $channelId) {
      id
      text
      sender { name }
      createdAt
    }
  }
`;

function ChatRoom({ channelId }) {
  const { data } = useSubscription(MESSAGES_SUBSCRIPTION, {
    variables: { channelId },
    onData: ({ client, data }) => {
      // Update Apollo cache to auto-refresh UI
      client.cache.writeQuery({ ... });
    }
  });
}
```

---

### Q4: How do you implement multi-tenancy in a GraphQL API?

**Answer:**
Multi-tenancy means one API serving multiple isolated customers (tenants). Common in SaaS applications.

**Strategy 1: Tenant from JWT Context:**
```javascript
// Every request carries tenant identification in the JWT
const context = async ({ req }) => {
  const token = req.headers.authorization?.replace('Bearer ', '');
  const decoded = jwt.verify(token, SECRET);
  const tenant = await tenantService.findById(decoded.tenantId);

  return {
    user: decoded,
    tenant,           // tenant is available in ALL resolvers
    db: getDbConnection(tenant.databaseUrl)  // tenant-specific DB
  };
};

// Resolver doesn't need to think about tenancy
const resolvers = {
  Query: {
    users: (_, __, context) => {
      // Uses tenant's own DB connection automatically
      return context.db.query('SELECT * FROM users');
    }
  }
};
```

**Strategy 2: Row-Level Security (shared DB):**
```javascript
// Set Postgres RLS policy for every request
const context = async ({ req }) => {
  const tenant = extractTenant(req);
  const db = await pool.connect();
  // Set current_tenant for Postgres RLS policies
  await db.query(`SET app.current_tenant = '${tenant.id}'`);

  return { db, tenant };

  // Postgres RLS policy: every SELECT/UPDATE/DELETE automatically filtered
  // CREATE POLICY tenant_isolation ON users USING (tenant_id = current_setting('app.current_tenant'));
};
```

**Strategy 3: Subdomain-based tenant identification:**
```javascript
const context = async ({ req }) => {
  const subdomain = req.hostname.split('.')[0];  // company1.app.com → "company1"
  const tenant = await tenantService.findBySubdomain(subdomain);

  if (!tenant) throw new ApolloError('Unknown tenant', 'TENANT_NOT_FOUND');

  return { tenant, db: tenantDb[tenant.id] };
};
```

**Schema isolation per tenant (using Federation):**
- Each tenant-specific subgraph can have custom fields
- Gateway routes based on tenant header

---

### Q5: How do you implement GraphQL in a Go microservice?

**Answer:**
Popular Go GraphQL libraries: `gqlgen` (code-first, type-safe), `graphql-go` (schema-first).

**Using `gqlgen` (recommended — code-first, generates resolvers):**

**1. Setup:**
```bash
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

**2. Schema** (`graph/schema.graphqls`):
```graphql
type Query {
  user(id: ID!): User
  users: [User!]!
}

type Mutation {
  createUser(input: CreateUserInput!): User!
}

type User {
  id: ID!
  name: String!
  email: String!
  posts: [Post!]!
}

input CreateUserInput {
  name: String!
  email: String!
}

type Post {
  id: ID!
  title: String!
}
```

**3. Generate resolver stubs:**
```bash
go run github.com/99designs/gqlgen generate
# Generates: graph/generated.go, graph/model/models_gen.go, graph/resolver.go
```

**4. Implement resolvers** (`graph/schema.resolvers.go`):
```go
package graph

import (
    "context"
    "github.com/yourapp/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
    user, err := r.UserService.FindByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    return user, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
    return r.UserService.Create(ctx, input.Name, input.Email)
}

// Field resolver — batching with dataloadgen
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
    return For(ctx).PostsByUserID.Load(obj.ID)  // DataLoader
}
```

**5. Main server:**
```go
package main

import (
    "net/http"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/yourapp/graph"
)

func main() {
    srv := handler.New(graph.NewExecutableSchema(graph.Config{
        Resolvers: &graph.Resolver{
            UserService: userService,
            PostService: postService,
        },
    }))

    http.Handle("/graphql", srv)
    http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
    http.ListenAndServe(":8080", nil)
}
```

**gqlgen advantage:** Fully type-safe — if your resolver returns the wrong type, it's a compile error, not a runtime panic. This is why product companies prefer it over `graphql-go`.

---

*Prepared for staff/senior-level technical interviews, system design rounds at product-based companies. Emphasizes real-world trade-offs over textbook answers.*
