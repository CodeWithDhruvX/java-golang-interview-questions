# 🟣 GraphQL Custom Directives & Production Monitoring (Product-Based Companies)

This document covers custom directive implementation and production-grade observability strategies — targeted at senior engineer interviews (4–8 years) at product companies like Razorpay, Postman, Atlassian, and Swiggy.

---

### Q1: How do you implement a custom GraphQL directive? Walk through an `@auth` directive.

**Answer:**
Custom directives let you declaratively inject cross-cutting logic into your schema without polluting resolvers.

**Common use cases:** `@auth`, `@rateLimit`, `@deprecated`, `@upper`, `@cacheControl`, `@log`

**Step 1 — Define in schema:**
```graphql
# Directive declaration
directive @auth(requires: Role = USER) on FIELD_DEFINITION | OBJECT

enum Role { ADMIN, MODERATOR, USER, GUEST }

type Query {
  publicPosts: [Post!]!                       # No directive — public
  myOrders: [Order!]!     @auth               # Requires USER (default)
  allUsers: [User!]!      @auth(requires: ADMIN)  # Admin only
}

type User @auth(requires: ADMIN) {            # Protects entire type
  id: ID!
  email: String!
  passwordHash: String!   # All fields on User require ADMIN
}
```

**Step 2 — Implement with `@graphql-tools/utils` (schema transformer):**
```javascript
const { mapSchema, getDirective, MapperKind } = require('@graphql-tools/utils');
const { defaultFieldResolver } = require('graphql');

function authDirectiveTransformer(schema, directiveName) {
  return mapSchema(schema, {
    // Apply to fields in Object types
    [MapperKind.OBJECT_FIELD]: (fieldConfig) => {
      const authDirective = getDirective(schema, fieldConfig, directiveName)?.[0];
      if (!authDirective) return fieldConfig;

      const { requires: requiredRole } = authDirective;
      const { resolve = defaultFieldResolver } = fieldConfig;

      // Wrap the existing resolver
      return {
        ...fieldConfig,
        resolve: async function (source, args, context, info) {
          const { user } = context;

          // Check authentication
          if (!user) {
            throw new GraphQLError('Not authenticated', {
              extensions: { code: 'UNAUTHENTICATED', http: { status: 401 } }
            });
          }

          // Check authorization (role hierarchy)
          const roleHierarchy = ['GUEST', 'USER', 'MODERATOR', 'ADMIN'];
          const userRoleLevel = roleHierarchy.indexOf(user.role);
          const requiredRoleLevel = roleHierarchy.indexOf(requiredRole);

          if (userRoleLevel < requiredRoleLevel) {
            throw new GraphQLError(`Requires ${requiredRole} role`, {
              extensions: { code: 'FORBIDDEN', http: { status: 403 } }
            });
          }

          // Proceed with original resolver
          return resolve(source, args, context, info);
        }
      };
    }
  });
}

// Apply transformer to schema
let schema = makeExecutableSchema({ typeDefs, resolvers });
schema = authDirectiveTransformer(schema, 'auth');
```

---

### Q2: Implement a `@rateLimit` directive to prevent API abuse.

**Answer:**
```graphql
directive @rateLimit(
  max: Int!            # Max requests allowed
  window: Int = 60     # Time window in seconds
  message: String = "Rate limit exceeded"
) on FIELD_DEFINITION

type Mutation {
  sendOTP(phone: String!): Boolean!
    @rateLimit(max: 3, window: 300, message: "Max 3 OTP requests per 5 minutes")

  createComment(text: String!): Comment!
    @rateLimit(max: 10, window: 60)
}
```

```javascript
const rateLimitDirectiveTransformer = (schema, directiveName, redisClient) => {
  return mapSchema(schema, {
    [MapperKind.OBJECT_FIELD]: (fieldConfig) => {
      const directive = getDirective(schema, fieldConfig, directiveName)?.[0];
      if (!directive) return fieldConfig;

      const { max, window: windowSec, message } = directive;
      const { resolve = defaultFieldResolver } = fieldConfig;

      return {
        ...fieldConfig,
        resolve: async function (source, args, context, info) {
          const key = `rateLimit:${context.user?.id || context.ip}:${info.fieldName}`;

          const current = await redisClient.incr(key);
          if (current === 1) {
            await redisClient.expire(key, windowSec);   // Set TTL on first hit
          }

          if (current > max) {
            const ttl = await redisClient.ttl(key);
            throw new GraphQLError(message, {
              extensions: {
                code: 'RATE_LIMITED',
                retryAfter: ttl
              }
            });
          }

          return resolve(source, args, context, info);
        }
      };
    }
  });
};
```

---

### Q3: Implement a `@upper` and `@trim` string transformation directive.

**Answer:**
Transform directives modify resolver output — useful for normalizing data.
```graphql
directive @upper on FIELD_DEFINITION
directive @trim on FIELD_DEFINITION
directive @truncate(limit: Int = 100) on FIELD_DEFINITION

type Product {
  name: String! @trim @upper        # "  laptop  " → "LAPTOP"
  description: String @truncate(limit: 200)
}
```

```javascript
function stringTransformDirectiveTransformer(schema) {
  return mapSchema(schema, {
    [MapperKind.OBJECT_FIELD]: (fieldConfig, fieldName, typeName, schema) => {
      const upperDir = getDirective(schema, fieldConfig, 'upper')?.[0];
      const trimDir = getDirective(schema, fieldConfig, 'trim')?.[0];
      const truncateDir = getDirective(schema, fieldConfig, 'truncate')?.[0];

      const hasDirective = upperDir || trimDir || truncateDir;
      if (!hasDirective) return fieldConfig;

      const { resolve = defaultFieldResolver } = fieldConfig;

      return {
        ...fieldConfig,
        resolve: async function (source, args, context, info) {
          let result = await resolve(source, args, context, info);

          if (typeof result === 'string') {
            if (trimDir) result = result.trim();
            if (upperDir) result = result.toUpperCase();
            if (truncateDir) {
              const { limit } = truncateDir;
              result = result.length > limit ? result.slice(0, limit) + '...' : result;
            }
          }

          return result;
        }
      };
    }
  });
}
```

---

### Q4: How do you implement `@log` and `@deprecated` custom directives for observability and schema governance?

**Answer:**

**`@log` directive (resolver-level logging):**
```graphql
directive @log(level: LogLevel = INFO) on FIELD_DEFINITION

enum LogLevel { DEBUG, INFO, WARN, ERROR }

type Query {
  sensitiveReport: Report! @log(level: WARN)   # Log every access with WARN
  debugQuery: DebugInfo!   @log(level: DEBUG)
}
```

```javascript
const logDirectiveTransformer = (schema, logger) => mapSchema(schema, {
  [MapperKind.OBJECT_FIELD]: (fieldConfig) => {
    const directive = getDirective(schema, fieldConfig, 'log')?.[0];
    if (!directive) return fieldConfig;

    const { resolve = defaultFieldResolver } = fieldConfig;

    return {
      ...fieldConfig,
      resolve: async function (source, args, context, info) {
        const start = Date.now();
        let error = null;

        try {
          const result = await resolve(source, args, context, info);
          return result;
        } catch (err) {
          error = err;
          throw err;
        } finally {
          const level = directive.level.toLowerCase();
          logger[level]({
            field: `${info.parentType.name}.${info.fieldName}`,
            userId: context.user?.id,
            duration: Date.now() - start,
            args: JSON.stringify(args),
            error: error?.message,
            requestId: context.requestId
          });
        }
      }
    };
  }
});
```

**Custom `@visible` feature-flag directive:**
```graphql
directive @featureFlag(name: String!) on FIELD_DEFINITION

type Query {
  newDashboard: Dashboard @featureFlag(name: "new_dashboard_ui")
}
```
```javascript
// Hides field if feature flag is off for the user
resolve: async (source, args, context, info) => {
  const isEnabled = await featureFlags.isEnabled(directive.name, context.user?.id);
  if (!isEnabled) return null;
  return resolve(source, args, context, info);
}
```

---

### Q5: How do you monitor a GraphQL API in production? (Datadog, Grafana, Apollo Studio)

**Answer:**
GraphQL monitoring is more granular than REST because a single endpoint can serve thousands of different operations.

**Key Metrics to track:**

| Metric | Why it matters |
|---|---|
| Request rate per operation name | Know which operations are most used |
| Error rate per operation | Catch regressions in specific queries |
| Resolver latency (P50/P95/P99) | Find slow fields |
| Field usage count | Know which fields to safely deprecate |
| Cache hit rate | Measure DataLoader/cache effectiveness |
| Subscription connection count | WebSocket server capacity |

**1. Apollo Studio (built-in for Apollo Server):**
```javascript
const { ApolloServerPluginUsageReporting } = require('@apollo/server/plugin/usageReporting');
const { ApolloServerPluginSchemaReporting } = require('@apollo/server/plugin/schemaReporting');

const server = new ApolloServer({
  plugins: [
    ApolloServerPluginUsageReporting({
      sendVariableValues: { none: true },   // Don't send PII variables
      sendHeaders: { onlyNames: ['x-client-version'] },
      generateClientInfo: ({ request }) => ({
        clientName: request.http?.headers.get('x-client-name') || 'Unknown',
        clientVersion: request.http?.headers.get('x-client-version') || '0.0.0'
      })
    }),
    ApolloServerPluginSchemaReporting()     // Sends schema to registry
  ]
});
```
Apollo Studio shows: field-level latency heatmaps, usage by client, error traces, and safe-to-remove fields.

**2. Custom Prometheus/Grafana metrics:**
```javascript
const client = require('prom-client');

const graphqlRequestDuration = new client.Histogram({
  name: 'graphql_request_duration_seconds',
  help: 'GraphQL request duration',
  labelNames: ['operationName', 'operationType', 'status'],
  buckets: [0.01, 0.05, 0.1, 0.3, 0.5, 1, 2, 5]
});

const graphqlErrors = new client.Counter({
  name: 'graphql_errors_total',
  help: 'Total GraphQL errors',
  labelNames: ['operationName', 'errorCode']
});

// Plugin to collect metrics
const metricsPlugin = {
  requestDidStart: () => {
    const start = Date.now();
    return {
      didEncounterErrors: ({ request, errors }) => {
        errors.forEach(err => {
          graphqlErrors.inc({
            operationName: request.operationName || 'unknown',
            errorCode: err.extensions?.code || 'UNKNOWN'
          });
        });
      },
      willSendResponse: ({ request, response }) => {
        const duration = (Date.now() - start) / 1000;
        const hasErrors = !!response.body?.singleResult?.errors?.length;
        graphqlRequestDuration.observe(
          {
            operationName: request.operationName || 'anonymous',
            operationType: getOperationType(request.query),
            status: hasErrors ? 'error' : 'success'
          },
          duration
        );
      }
    };
  }
};

// Expose /metrics for Prometheus scraping
app.get('/metrics', async (req, res) => {
  res.set('Content-Type', client.register.contentType);
  res.end(await client.register.metrics());
});
```

**3. Datadog APM with distributed tracing:**
```javascript
const tracer = require('dd-trace').init({
  service: 'graphql-api',
  env: process.env.NODE_ENV,
  analytics: true
});

// Datadog auto-instruments Apollo Server
// Add custom operation name span tag for better visibility
const datadogPlugin = {
  requestDidStart: () => ({
    executionDidStart: ({ document, operationName }) => {
      const span = tracer.scope().active();
      if (span) {
        span.setTag('graphql.operation.name', operationName || 'anonymous');
        span.setTag('graphql.operation.type', document?.definitions?.[0]?.operation);
      }
    }
  })
};
```

**Grafana Dashboard panels to set up:**
- Requests/sec by operation name (time series)
- Error rate % by operation (gauge)
- P95 latency by operation (bar chart)
- Slow resolver heat map (table sorted by P95)
- Subscription active connections (gauge)

---

### Q6: How do you set up CI/CD schema validation to prevent breaking changes from reaching production?

**Answer:**
Schema changes should be validated automatically in CI before they can be merged/deployed.

**Using Rover CLI (Apollo) in GitHub Actions:**
```yaml
# .github/workflows/graphql-checks.yml
name: GraphQL Schema Check
on: [push, pull_request]

jobs:
  schema-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Install Rover CLI
        run: curl -sSL https://rover.apollo.dev/nix/latest | sh
        
      - name: Validate Schema (no syntax errors)
        run: rover graph introspect http://localhost:4000/graphql

      - name: Check for Breaking Changes vs Production
        env:
          APOLLO_KEY: ${{ secrets.APOLLO_KEY }}
        run: |
          rover graph check my-graph@current \
            --schema ./schema.graphql
          # ❌ Fails if: removed field, changed type, removed enum value
          # ✅ Passes if: added field, added optional arg, added type

      - name: Publish Schema if Merged to Main
        if: github.ref == 'refs/heads/main'
        run: |
          rover graph publish my-graph@current \
            --schema ./schema.graphql
```

**Using Hive (open-source alternative to Apollo Studio):**
```bash
# Hive CLI
npm install -g @graphql-hive/cli

# Check for breaking changes
hive schema:check --service my-service ./schema.graphql

# Publish schema
hive schema:publish --service my-service ./schema.graphql
```

**Custom breaking-change detection with `@graphql-inspector`:**
```bash
npm install -g @graphql-inspector/cli

# Compare current schema against production
graphql-inspector diff \
  "schema.graphql"   \
  "https://api.prod.com/graphql"

# Output:
# ❌ Field 'User.username' was removed (BREAKING)  
# ✅ Field 'User.handle' was added (NON-BREAKING)
# ⚠️  Field 'Post.body' was deprecated (NON-BREAKING)
```

This ensures **zero accidental breaking changes** ever reach production, which is a key expectation at companies like Razorpay and Postman where external partners consume the API.

---

*Prepared for staff/senior engineer interviews. Covers real production tooling expected at product companies with mature GraphQL stacks.*
