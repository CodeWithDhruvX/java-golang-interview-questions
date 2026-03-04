# 🟣 GraphQL Performance & Schema Design — Interview Questions (Product-Based Companies)

This document covers GraphQL performance optimization and schema design patterns for product-based company interviews. Targeted at senior engineers (4–8 years experience) at high-growth companies.

---

### Q1: How do you design a production-grade GraphQL schema? Discuss naming conventions, deprecation strategy, and modularization.

**Answer:**
Schema design is a long-term architectural decision. Product companies treat GraphQL schema like a public API contract.

**Naming Conventions:**
```graphql
# Types — PascalCase
type ProductVariant { ... }
type OrderLineItem { ... }

# Fields — camelCase
type Product {
  id: ID!
  displayName: String!
  basePrice: Float!
  isAvailable: Boolean!
  variantList: [ProductVariant!]!  # Not "variants" if it's paginated
}

# Queries — camelCase verb or noun
type Query {
  product(id: ID!): Product
  products(filter: ProductFilter, pagination: PaginationInput): ProductConnection!
  searchProducts(query: String!): [Product!]!
}

# Mutations — camelCase with verb prefix
type Mutation {
  createProduct(input: CreateProductInput!): CreateProductPayload!
  updateProduct(id: ID!, input: UpdateProductInput!): UpdateProductPayload!
  deleteProduct(id: ID!): DeleteProductPayload!
}

# Input types — PascalCase + suffix "Input"
input CreateProductInput {
  name: String!
  price: Float!
  categoryId: ID!
}

# Mutation return/payload types — PascalCase + suffix "Payload"
type CreateProductPayload {
  product: Product
  errors: [UserError!]!
}
```

**Deprecation Strategy (non-breaking evolution):**
```graphql
type User {
  # Mark old field deprecated BEFORE removal
  username: String @deprecated(reason: "Use `handle` instead. Will be removed in API version 3.")
  handle: String!

  # Old scalar — switching to custom type
  profilePicUrl: String @deprecated(reason: "Use `avatar { url }` for multiple sizes.")
  avatar: Avatar

  # Old flat args — moving to Input type
  # Don't change existing mutation signature abruptly
}

type Mutation {
  # Keep old form, add new form
  createUserLegacy(name: String!, email: String!): User @deprecated(reason: "Use createUser(input:)")
  createUser(input: CreateUserInput!): CreateUserPayload!
}
```

**Schema Modularization:**
```
src/
├── graphql/
│   ├── schema.js           # Entry: merge all modules
│   ├── user/
│   │   ├── user.typedefs.js
│   │   ├── user.resolvers.js
│   │   └── user.loaders.js
│   ├── product/
│   │   ├── product.typedefs.js
│   │   ├── product.resolvers.js
│   │   └── product.loaders.js
│   └── order/
│       ├── order.typedefs.js
│       └── order.resolvers.js
```

```javascript
const { mergeTypeDefs, mergeResolvers } = require('@graphql-tools/merge');

const typeDefs = mergeTypeDefs([userTypeDefs, productTypeDefs, orderTypeDefs]);
const resolvers = mergeResolvers([userResolvers, productResolvers, orderResolvers]);
```

---

### Q2: How do you implement cursor-based connection pagination following the Relay spec (production implementation)?

**Answer:**
The **Relay Connection Specification** is the de facto standard for pagination in GraphQL. Product companies like Facebook, GitHub, and Shopify use it.

**Complete Schema:**
```graphql
# Generic reusable types
type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}

# Product-specific connection
type ProductConnection {
  edges: [ProductEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
}

type ProductEdge {
  node: Product!
  cursor: String!
}

type Query {
  # Supports both forward (first/after) and backward (last/before) pagination
  products(
    first: Int
    after: String
    last: Int
    before: String
    filter: ProductFilter
    sortBy: ProductSortInput
  ): ProductConnection!
}

input ProductFilter {
  categoryId: ID
  minPrice: Float
  maxPrice: Float
  isAvailable: Boolean
}

input ProductSortInput {
  field: ProductSortField!
  direction: SortDirection!
}

enum ProductSortField { CREATED_AT, PRICE, NAME, POPULARITY }
enum SortDirection { ASC, DESC }
```

**Production Resolver:**
```javascript
const resolvers = {
  Query: {
    products: async (_, args, context) => {
      const { first = 20, after, last, before, filter, sortBy } = args;

      // Validate: can't use first+before or last+after together
      if (first && last) throw new UserInputError("Cannot use 'first' and 'last' together");

      // Decode opaque cursor → actual DB cursor value
      const decodeCursor = (cursor) =>
        cursor ? JSON.parse(Buffer.from(cursor, 'base64').toString('utf8')) : null;

      const afterCursor = decodeCursor(after);
      const beforeCursor = decodeCursor(before);
      const limit = first || last || 20;

      // Build WHERE clause from filter
      const whereClause = buildWhereClause(filter, afterCursor, beforeCursor);
      const orderClause = buildOrderClause(sortBy, !!last);

      const [products, totalCount] = await Promise.all([
        db.query(`SELECT * FROM products ${whereClause} ${orderClause} LIMIT ${limit + 1}`),
        db.queryOne(`SELECT COUNT(*) FROM products ${buildWhereClause(filter)}`)
      ]);

      const hasMore = products.length > limit;
      const items = products.slice(0, limit);
      if (last) items.reverse(); // Reverse for backward pagination

      const encodeCursor = (product) =>
        Buffer.from(JSON.stringify({ id: product.id, createdAt: product.createdAt })).toString('base64');

      const edges = items.map(product => ({
        node: product,
        cursor: encodeCursor(product)
      }));

      return {
        edges,
        totalCount: parseInt(totalCount.count),
        pageInfo: {
          hasNextPage: !!first && hasMore,
          hasPreviousPage: !!last && hasMore,
          startCursor: edges[0]?.cursor ?? null,
          endCursor: edges[edges.length - 1]?.cursor ?? null
        }
      };
    }
  }
};
```

**Client Usage:**
```javascript
const { data, fetchMore } = useQuery(GET_PRODUCTS, {
  variables: { first: 20 }
});

// Load next page
const loadMore = () => {
  fetchMore({
    variables: { after: data.products.pageInfo.endCursor },
    updateQuery: (prev, { fetchMoreResult }) => ({
      products: {
        ...fetchMoreResult.products,
        edges: [...prev.products.edges, ...fetchMoreResult.products.edges]
      }
    })
  });
};
```

---

### Q3: How do you implement a global error handling and observability strategy for a production GraphQL API?

**Answer:**
Production GraphQL APIs need structured error handling, logging, and tracing to debug issues across resolvers and subgraphs.

**1. Structured Error Taxonomy:**
```javascript
// errors/index.js — Custom error classes
const { GraphQLError } = require('graphql');

class AppError extends GraphQLError {
  constructor(message, code, httpStatus = 400, additionalProperties = {}) {
    super(message, {
      extensions: {
        code,
        httpStatus,
        timestamp: new Date().toISOString(),
        ...additionalProperties
      }
    });
  }
}

class NotFoundError extends AppError {
  constructor(resource, id) {
    super(`${resource} with id '${id}' not found`, 'NOT_FOUND', 404, { resource, id });
  }
}

class ValidationError extends AppError {
  constructor(message, fields) {
    super(message, 'VALIDATION_ERROR', 400, { invalidFields: fields });
  }
}

class RateLimitError extends AppError {
  constructor() {
    super('Too many requests', 'RATE_LIMIT_EXCEEDED', 429, { retryAfter: 60 });
  }
}

// Usage in resolver
const resolvers = {
  Query: {
    user: async (_, { id }) => {
      const user = await db.findUser(id);
      if (!user) throw new NotFoundError('User', id);
      return user;
    }
  }
};
```

**2. Global Error Formatter + Logging:**
```javascript
const server = new ApolloServer({
  formatError: (formattedError, error) => {
    // Log all errors with context
    logger.error({
      message: formattedError.message,
      code: formattedError.extensions?.code,
      path: formattedError.path,
      locations: formattedError.locations,
      stack: error.stack,
      requestId: requestContext?.requestId
    });

    // Sanitize internal errors before sending to client
    if (formattedError.extensions?.code === 'INTERNAL_SERVER_ERROR') {
      return {
        message: 'An internal error occurred',
        extensions: { code: 'INTERNAL_SERVER_ERROR' }
      };
    }

    return formattedError;
  }
});
```

**3. Apollo Studio Integration + OpenTelemetry Tracing:**
```javascript
const { ApolloServerPluginUsageReporting } = require('@apollo/server/plugin/usageReporting');
const { ApolloServerPluginInlineTrace } = require('@apollo/server/plugin/inlineTrace');
const { NodeTracerProvider } = require('@opentelemetry/sdk-trace-node');
const { GraphQLInstrumentation } = require('@opentelemetry/instrumentation-graphql');

// OpenTelemetry — traces every resolver
const provider = new NodeTracerProvider();
new GraphQLInstrumentation({ mergeItems: false }).enable(); // Auto-instruments resolvers

const server = new ApolloServer({
  plugins: [
    // Reports to Apollo Studio — field-level latency, error rates
    ApolloServerPluginUsageReporting({ sendVariableValues: { none: true } }),
    ApolloServerPluginInlineTrace()
  ]
});
```

**4. Custom Plugin for Request Logging:**
```javascript
const requestLoggerPlugin = {
  async requestDidStart(requestContext) {
    const start = Date.now();
    const requestId = crypto.randomUUID();

    return {
      async willSendResponse(context) {
        logger.info({
          requestId,
          operationName: context.request.operationName,
          duration: Date.now() - start,
          hasErrors: !!context.response.body?.singleResult?.errors?.length,
          variables: context.request.variables
        });
      }
    };
  }
};
```

---

### Q4: How do you handle GraphQL schema versioning and evolution without breaking clients?

**Answer:**
GraphQL's design principle is **"versionless APIs"** — the schema should evolve without explicit versions. However, this requires discipline.

**Seven Rules for Non-Breaking Schema Changes:**

| Change Type | Breaking? | Approach |
|---|---|---|
| Add a new optional field | ✅ Non-breaking | Just add it |
| Add a new query/mutation | ✅ Non-breaking | Just add it |
| Add optional argument to a field | ✅ Non-breaking | Set a default value |
| Remove a field | ❌ Breaking | First `@deprecated`, wait, then remove |
| Rename a field | ❌ Breaking | Add new name, deprecate old |
| Change a scalar type | ❌ Breaking | Add new field, deprecate old |
| Make optional field required | ❌ Breaking | Add new mutation variant |
| Change return type | ❌ Breaking | Add new field with new type |

**Deprecation Workflow:**
```graphql
# Phase 1: Add new field, deprecate old
type User {
  username: String @deprecated(reason: "Use handle. Will be removed 2025-Q1.")
  handle: String!

  profilePic: String @deprecated(reason: "Use avatar.url for different sizes")
  avatar: Avatar
}

# Phase 2: Monitor usage in Apollo Studio (field-level metrics)
# Phase 3: When usage → 0, safely remove deprecated fields

# Never remove in < 3 months; communicate with teams beforehand
```

**Schema Registry (production):**
Use **Apollo Schema Registry** or **Hive** to:
- Validate new schemas for breaking changes before deployment
- Track schema change history
- Monitor field usage to know when it's safe to remove deprecated fields

```bash
# Apollo Rover CLI — check for breaking changes
rover schema check \
  --schema ./schema.graphql \
  --name production \
  --variant current
# Fails CI if breaking changes are detected
```

---

### Q5: How do you implement GraphQL in a microservices architecture without Federation? Discuss the BFF pattern.

**Answer:**
**BFF (Backend for Frontend)** is an alternative to Federation where each client type (web, mobile, partner API) has its own GraphQL gateway that aggregates from multiple downstream REST or gRPC services.

**Architecture:**
```
Web Browser → Web BFF (GraphQL) → User Service (REST)
                                → Order Service (gRPC)
                                → Product Service (REST)

Mobile App → Mobile BFF (GraphQL) → User Service (REST)
                                  → Notification Service (REST)
```

**BFF Implementation with Schema Stitching:**
```javascript
const { stitchSchemas } = require('@graphql-tools/stitch');
const { buildSchema, GraphQLSchema } = require('graphql');
const { wrapSchema, RenameTypes, FilterRootFields } = require('@graphql-tools/wrap');

// Remote schema executors (call microservices)
const userServiceExecutor = async ({ document, variables }) => {
  const res = await fetch('http://user-service/graphql', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ query: print(document), variables })
  });
  return res.json();
};

// Build gateway schema
const gatewaySchema = stitchSchemas({
  subschemas: [
    {
      schema: await introspectSchema(userServiceExecutor),
      executor: userServiceExecutor,
      transforms: [
        new FilterRootFields((op, field) => ['user', 'users'].includes(field))
      ]
    },
    {
      schema: await introspectSchema(orderServiceExecutor),
      executor: orderServiceExecutor
    }
  ],
  // Type merging — link User.orders across services
  typeMergingOptions: {
    User: {
      selectionSet: '{ id }',
      fieldName: 'userById',
      args: (user) => ({ id: user.id })
    }
  }
});
```

**BFF vs Federation:**

| Aspect | BFF Pattern | Apollo Federation |
|---|---|---|
| Team ownership | BFF team owns gateway | Each microservice owns its schema |
| Schema composition | Manual stitching at gateway | Declarative via `@key` directives |
| Deployment | BFF is a bottleneck | Each subgraph deploys independently |
| Suitable for | Different APIs per client type | Single unified API for all clients |
| Complexity | Simpler to start | Better for large, distributed teams |

---

*Prepared for senior-level and staff engineer interviews at product-based companies. Focuses on real production decisions and trade-offs.*
