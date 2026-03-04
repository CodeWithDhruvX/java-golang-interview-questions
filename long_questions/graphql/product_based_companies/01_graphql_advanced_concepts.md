# 🟣 GraphQL Advanced Concepts — Interview Questions (Product-Based Companies)

This document covers advanced GraphQL concepts targeted at product-based company interviews (Google, Amazon, Flipkart, Swiggy, Razorpay, Atlassian, etc.) for 3–8 years experience rounds. Focus is on performance, architecture, and production readiness.

---

### Q1: What is Apollo Federation and how does it enable a distributed GraphQL architecture?

**Answer:**
**Apollo Federation** is an architecture for building a **unified GraphQL supergraph** from multiple independently deployable **subgraphs** (microservices). It's the production standard for scaling GraphQL across teams.

**Problem it solves:**
Traditional schema stitching required all schemas to merge manually and the gateway had too much business logic. Federation is declarative and allows each team to own their schema.

**Architecture:**
```
Client → Apollo Router (Supergraph Gateway)
              ↓ fan-out (query plan)
   ┌──────────┬──────────┬──────────┐
   ↓          ↓          ↓
Users      Products   Orders
Subgraph   Subgraph   Subgraph
 (Node.js)  (Java)     (Go)
```

**Entity Definition (Users Subgraph):**
```graphql
# Users Subgraph schema
type User @key(fields: "id") {     # @key declares this as a "federated entity"
  id: ID!
  name: String!
  email: String!
}

type Query {
  me: User
}
```

**Extending an Entity (Orders Subgraph):**
```graphql
# Orders Subgraph schema
extend type User @key(fields: "id") {  # Extend User from Users subgraph
  orders: [Order!]!                     # Add orders field to User
}

type Order @key(fields: "id") {
  id: ID!
  total: Float!
  status: OrderStatus!
}

enum OrderStatus { PENDING, SHIPPED, DELIVERED, CANCELLED }
```

**Reference Resolver** (Orders Subgraph must resolve User entities):
```javascript
const resolvers = {
  User: {
    // Called when gateway needs to resolve a User entity
    __resolveReference(userRef) {
      // userRef = { id: "123" } — only the key fields
      return { id: userRef.id }; // Orders subgraph only adds orders, not user data
    },
    orders: (user) => orderService.findOrdersByUser(user.id)
  }
};
```

**Key Federation directives:**
| Directive | Purpose |
|---|---|
| `@key(fields: "id")` | Marks an entity with a primary key |
| `@external` | Field defined in another subgraph |
| `@requires` | Requires fields from another subgraph before resolving |
| `@provides` | Tells gateway this subgraph can provide these fields |
| `@shareable` | Multiple subgraphs can resolve this field |
| `@override` | This subgraph takes ownership from another |

---

### Q2: How does query planning work in Apollo Router? What is the query plan used for?

**Answer:**
**Query Planning** is the process where the Apollo Router (gateway) analyses an incoming GraphQL query and determines how to split it into sub-queries that get sent to the appropriate subgraphs.

**Example:**
```graphql
# Client sends this single query
query {
  me {               # → resolved by Users subgraph
    name
    email
    orders {         # → resolved by Orders subgraph (extends User)
      id
      total
      product {      # → resolved by Products subgraph (extends Order)
        name
        price
      }
    }
  }
}
```

**Generated Query Plan (conceptual):**
```
Fetch (Users subgraph):
  query { me { id name email } }
    ↓ Returns: { me: { id: "123", name: "John", email: "j@j.com" } }

Fetch (Orders subgraph) using User.id:
  query { _entities(representations: [{ __typename: "User", id: "123" }]) {
    ... on User { orders { id total productId } }
  }}
    ↓ Returns: orders with productIds

Fetch (Products subgraph) using productIds:
  query { _entities(representations: [{ __typename: "Product", id: "p1" }]) {
    ... on Product { name price }
  }}
```

**The `_entities` query** is a special internal query used by Federation to fetch entities by their keys (reference resolver).

**Key performance insight:** The Router creates the optimal **parallel execution plan** — it only makes sequential calls when data from subgraph A is needed by subgraph B. Independent subgraphs are called in **parallel**.

---

### Q3: How do you secure a GraphQL API in production? Discuss rate limiting, query depth, and query cost analysis.

**Answer:**
GraphQL's flexibility makes it vulnerable to abuse. Unlike REST where you can rate-limit individual endpoints, a single GraphQL endpoint can receive queries of wildly varying complexity.

**1. Query Depth Limiting:**
Prevent deeply nested, recursive queries that cause DB explosion.
```graphql
# Malicious deep query (causes exponential DB load)
query {
  user(id: "1") {
    friends {
      friends {
        friends {
          friends { name }  # Could be millions of DB calls
        }
      }
    }
  }
}
```
```javascript
const depthLimit = require('graphql-depth-limit');

const server = new ApolloServer({
  typeDefs,
  resolvers,
  validationRules: [depthLimit(5)]   // Reject queries deeper than 5 levels
});
```

**2. Query Complexity / Cost Analysis:**
Assign a cost to each field; reject if total exceeds threshold.
```javascript
const { createComplexityLimitRule } = require('graphql-validation-complexity');

const server = new ApolloServer({
  validationRules: [
    createComplexityLimitRule(1000, {
      // Each field has a cost; list fields cost more
      onCost: (cost) => console.log(`Query cost: ${cost}`),
      formatErrorMessage: (cost) => `Query complexity ${cost} exceeds max 1000`
    })
  ]
});
```

**3. Query Timeout:**
```javascript
// Per request, use a promise race
const resolvers = {
  Query: {
    heavyQuery: async (_, args) => {
      return Promise.race([
        fetchData(args),
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error('Query timeout')), 5000)
        )
      ]);
    }
  }
};
```

**4. Persisted Queries:**
Only allow pre-approved, whitelisted query hashes in production.
```javascript
// Client sends query hash instead of full query string
// { "extensions": { "persistedQuery": { "sha256Hash": "abc123..." } } }

// Server looks up the hash, rejects unknown queries
const server = new ApolloServer({
  plugins: [persistedQueryPlugin({ cache: redisCache })]
});
```

**5. Rate Limiting by IP/Token:**
```javascript
const { RateLimiterRedis } = require('rate-limiter-flexible');
const rateLimiter = new RateLimiterRedis({ storeClient: redis, points: 100, duration: 60 });

app.use('/graphql', async (req, res, next) => {
  try {
    await rateLimiter.consume(req.ip);
    next();
  } catch {
    res.status(429).json({ error: 'Too many requests' });
  }
});
```

**6. Disable Introspection in production:**
```javascript
const server = new ApolloServer({
  introspection: process.env.NODE_ENV !== 'production'
});
```

---

### Q4: How do you implement GraphQL subscriptions at scale using WebSockets?

**Answer:**
Subscriptions require a **stateful long-lived connection** — fundamentally different from queries/mutations. This creates scaling challenges.

**Basic Setup (Apollo Server + WebSocket):**
```javascript
const { createServer } = require('http');
const { WebSocketServer } = require('ws');
const { useServer } = require('graphql-ws/lib/use/ws');
const { PubSub } = require('graphql-subscriptions');

const pubsub = new PubSub();  // In-memory — NOT for production!

const typeDefs = gql`
  type Subscription {
    commentAdded(postId: ID!): Comment!
  }
  type Mutation {
    addComment(postId: ID!, text: String!): Comment!
  }
`;

const resolvers = {
  Subscription: {
    commentAdded: {
      subscribe: (_, { postId }) =>
        pubsub.asyncIterator(`COMMENT_ADDED.${postId}`),
      resolve: (payload) => payload.comment
    }
  },
  Mutation: {
    addComment: async (_, args) => {
      const comment = await db.createComment(args);
      // Publish to all subscribers for this postId
      pubsub.publish(`COMMENT_ADDED.${args.postId}`, { comment });
      return comment;
    }
  }
};

// HTTP server for queries/mutations
const httpServer = createServer(app);

// WebSocket server for subscriptions
const wsServer = new WebSocketServer({ server: httpServer, path: '/graphql' });
useServer({ schema }, wsServer);
httpServer.listen(4000);
```

**Production Scaling Challenges:**
| Challenge | Solution |
|---|---|
| In-memory `PubSub` doesn't work across multiple pods | Use **Redis PubSub** (`graphql-redis-subscriptions`) |
| WebSocket connections are stateful | Use **dedicated WebSocket gateway** (a separate tier) |
| Thousands of concurrent connections | Use event-driven architecture (Redis + horizontal scaling of WS nodes) |
| Authorization per subscription | Check context when client subscribes |

**Redis-backed PubSub (production):**
```javascript
const { RedisPubSub } = require('graphql-redis-subscriptions');
const { createClient } = require('redis');

const pubsub = new RedisPubSub({
  publisher: createClient({ url: process.env.REDIS_URL }),
  subscriber: createClient({ url: process.env.REDIS_URL })
});
// Now works across all pods!
```

---

### Q5: Explain GraphQL caching strategies — client-side (Apollo), CDN, and response caching.

**Answer:**
GraphQL's single endpoint makes HTTP-level caching harder (everything is a POST). Different caching layers are needed.

**1. Apollo Client Normalized Cache (client-side):**
Apollo Client stores every entity by `id + __typename` in a flat cache, automatically deduplicating.
```javascript
const client = new ApolloClient({
  cache: new InMemoryCache({
    typePolicies: {
      User: {
        keyFields: ['id'],      // Cache key for User entities
        fields: {
          posts: {
            // Custom merge strategy for pagination
            keyArgs: ['status'],
            merge(existing = [], incoming) {
              return [...existing, ...incoming];
            }
          }
        }
      }
    }
  })
});

// Fetch policies
const { data } = useQuery(GET_USER, {
  fetchPolicy: 'cache-first',       // Use cache, fallback to network
  // 'cache-and-network'            // Return cache immediately, update from network
  // 'network-only'                 // Always hit network
  // 'no-cache'                     // Always network, don't store result
});
```

**2. Server-Side Response Caching (Apollo Server):**
```javascript
const { ApolloServerPluginCacheControl } = require('@apollo/server/plugin/cacheControl');

const server = new ApolloServer({
  plugins: [ApolloServerPluginCacheControl({ defaultMaxAge: 30 })]
});
```

```graphql
# In schema — per-field cache hints
type User {
  id: ID!
  name: String! @cacheControl(maxAge: 300, scope: PUBLIC)
  email: String! @cacheControl(maxAge: 0, scope: PRIVATE)   # Don't cache private data!
  posts: [Post] @cacheControl(maxAge: 60)
}
```

**3. Persisted Query CDN Caching:**
Send a GET request with the query hash. CDN can cache the response!
```bash
# GET request — CDN-cacheable!
GET /graphql?operationName=GetHomePage&variables={}&extensions={"persistedQuery":{"sha256Hash":"abc..."}}
```

**4. DataLoader (per-request batching — not long-term caching):**
Caches within a single request lifecycle to prevent duplicate fetches.

**Caching hierarchy:**
```
CDN (persisted GET queries) → Server response cache → DataLoader (request-level) → Apollo Client cache
```

---

### Q6: How do you implement real-world authorization patterns in GraphQL (field-level, type-level, rule-based)?

**Answer:**
Authorization in GraphQL should be done in the **business logic layer** (resolvers), not the schema layer.

**Pattern 1: Check in each resolver (basic):**
```javascript
const resolvers = {
  Query: {
    adminDashboard: (_, __, context) => {
      if (!context.user?.roles.includes('ADMIN')) throw new ForbiddenError('Admin only');
      return dashboardService.getSummary();
    }
  }
};
```

**Pattern 2: GraphQL Shield (rule-based middleware):**
```javascript
const { shield, rule, and, or, not } = require('graphql-shield');

const isAuthenticated = rule({ cache: 'contextual' })(
  async (parent, args, ctx) => ctx.user !== null
);

const isAdmin = rule({ cache: 'contextual' })(
  async (parent, args, ctx) => ctx.user?.role === 'ADMIN'
);

const isOwner = rule({ cache: 'strict' })(
  async (parent, args, ctx) => parent.userId === ctx.user.id
);

const permissions = shield({
  Query: {
    users: isAdmin,
    me: isAuthenticated
  },
  Mutation: {
    createUser: isAuthenticated,
    deleteUser: and(isAuthenticated, isAdmin)
  },
  User: {
    email: or(isAdmin, isOwner),     // Field-level: only admin or self can see email
    '*': isAuthenticated             // Default: all User fields require auth
  }
});
```

**Pattern 3: Attribute-Based Access Control (ABAC):**
```javascript
// Resolver checks specific attributes
const resolvers = {
  Query: {
    document: async (_, { id }, ctx) => {
      const doc = await documentService.findById(id);
      const canView = await permissionService.check({
        subject: ctx.user,
        action: 'view',
        resource: doc,
        conditions: {
          department: ctx.user.department === doc.department,
          classification: ctx.user.clearanceLevel >= doc.classificationLevel
        }
      });
      if (!canView) throw new ForbiddenError('Insufficient clearance');
      return doc;
    }
  }
};
```

---

*Prepared for senior-level technical interviews at product-based companies (Amazon, Flipkart, Swiggy, Razorpay, Atlassian, ThoughtWorks, Postman).*
