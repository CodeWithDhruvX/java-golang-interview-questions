# 🔗 Microservices — GraphQL Federation & Apollo Federation

> **Level:** 🔴 Senior – Principal
> **Asked at:** Netflix, Swiggy, Flipkart, Myntra, Razorpay (companies with complex API composition needs)

---

## Q1. What is GraphQL Federation? How is it different from a REST API Gateway?

**"Federation solves the problem of API composition at scale. In a microservices architecture, a single product screen (like the Swiggy home screen) needs data from 6+ services: Restaurant, Menu, User, Cart, Promotions, Ratings.**

**Without Federation (REST Gateway approach):**
- Frontend makes 4–6 REST calls OR the BFF (Backend for Frontend) aggregates
- Adding a new field means modifying the BFF — a shared codebase touched by everyone
- No type safety across service boundaries

**With Apollo Federation:**
- Each service owns and publishes a GraphQL schema (subgraph)
- A `Router` (formerly Apollo Gateway) automatically stitches all subgraphs into one unified supergraph
- Frontend makes ONE GraphQL query — the router decomposes it into sub-queries, fans them out to the right services, and joins the results

```
             Mobile App
                │
                │ Single GraphQL query
                ▼
        ┌───────────────┐
        │  Apollo Router │  (Federation Gateway)
        │  (Supergraph)  │
        └───────────────┘
         │    │    │    │
         ▼    ▼    ▼    ▼
    User  Order Cart  Promotion
   subgraph  subgraph subgraph subgraph
    (Java)  (Go)   (Node)  (Python)
```

**Key difference from REST Gateway:**
| REST API Gateway | GraphQL Federation |
|---------|------------------|
| Routes by URL path | Routes by type/field in schema |
| Frontend decides which APIs to call | Router decides which subgraphs to query |
| Response shape determined by each service | Response shape determined by frontend query |
| Adding a field needs a new endpoint or BFF change | Adding a field needs only the owning subgraph to change |"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix (they helped create Federation), Swiggy, Myntra. Any company building a super-app with a complex mobile frontend will ask this.

#### Indepth
**Apollo Federation vs Schema Stitching:**
- **Schema Stitching** (old approach): Gateway manually merges schemas — brittle, requires gateway code changes for each service
- **Federation** (modern): Each subgraph declares its own type extensions. The router auto-discovers and composes. Decentralized ownership.

---

## Q2. Walk me through implementing Apollo Federation — show the code.

### ✅ Step 1: User Subgraph (Java — Spring GraphQL)

```java
// UserSubgraph — defines the User type
// pom.xml: spring-boot-starter-graphql + federation-graphql-java-support

// schema.graphqls (User service)
"""
extend schema @link(url: "https://specs.apollo.dev/federation/v2.0", import: ["@key"])

type User @key(fields: "id") {
  id: ID!
  name: String!
  email: String!
}

type Query {
  user(id: ID!): User
}
"""

// UserResolver.java
@Controller
public class UserResolver {

    @QueryMapping
    public User user(@Argument String id) {
        return userRepository.findById(id).orElseThrow();
    }

    // Required by Federation: resolves User entity by reference
    @EntityMapping
    public User resolveUserById(@Argument String id) {
        return userRepository.findById(id).orElseThrow();
    }
}
```

### ✅ Step 2: Order Subgraph (Go) — extends the User type

```go
// In Go Order service schema
// schema.graphql
"""
extend schema @link(url: "https://specs.apollo.dev/federation/v2.0", import: ["@key", "@external"])

# Order owns OrderType but EXTENDS User from another subgraph
type User @key(fields: "id") {
  id: ID! @external    # This field belongs to User subgraph
  orders: [Order!]!    # But we ADD this field here — order subgraph owns it
}

type Order {
  id: ID!
  status: String!
  total: Float!
  createdAt: String!
}
"""

// Order resolver in Go
func (r *Resolver) UserOrders(ctx context.Context, user *model.User) ([]*model.Order, error) {
    return r.orderRepo.FindByUserID(ctx, user.ID)
}
```

### ✅ Step 3: Apollo Router Configuration (router.yaml)

```yaml
# rover.yaml — compose the supergraph
federation_version: =2.3.2
subgraphs:
  user:
    routing_url: http://user-service:8080/graphql
    schema:
      subgraph_url: http://user-service:8080/graphql
  order:
    routing_url: http://order-service:8080/graphql
    schema:
      subgraph_url: http://order-service:8080/graphql
  cart:
    routing_url: http://cart-service:8080/graphql
    schema:
      subgraph_url: http://cart-service:8080/graphql
```

```bash
# Compose the supergraph schema (run in CI/CD)
rover supergraph compose --config rover.yaml > supergraph.graphql

# Start the router
./router --config router.yaml --supergraph supergraph.graphql
```

### ✅ Step 4: Single Query From Frontend

```graphql
# One query to the Apollo Router — touches User + Order + Cart subgraphs
query GetUserDashboard($userId: ID!) {
  user(id: $userId) {
    id
    name
    email
    orders {
      id
      status
      total
    }
  }
}
```

**Router execution plan:**
1. Query `user(id)` → User subgraph
2. Pass returned `user.id` → Order subgraph for `orders`
3. Merge and return unified response

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Myntra, Swiggy. The interviewer may ask you to whiteboard the `@key` and `@external` directives — know that `@key` defines the entity's primary key for federation resolution, and `@external` marks fields owned by another subgraph.

---

## Q3. How do you handle N+1 queries in GraphQL Federation?

**"N+1 is the classic GraphQL problem: if you fetch 10 orders and each order needs user data, naively you'd make 10+1 separate calls to the User subgraph. At scale this is catastrophic.**

**DataLoader Pattern:**
```java
// DataLoader batches individual calls into one batch call
@Component
public class UserDataLoader implements BatchLoader<String, User> {

    @Autowired private UserRepository userRepository;

    @Override
    public CompletionStage<List<User>> load(List<String> userIds) {
        // Called ONCE with ALL the userIds, not N times with one id each
        return CompletableFuture.supplyAsync(() -> {
            Map<String, User> userMap = userRepository.findAllById(userIds)
                .stream()
                .collect(Collectors.toMap(User::getId, Function.identity()));

            // Return in the SAME order as userIds
            return userIds.stream()
                .map(userMap::get)
                .collect(Collectors.toList());
        });
    }
}

// In the resolver — use DataLoader instead of direct repository call
@SchemaMapping(typeName = "Order", field = "user")
public CompletableFuture<User> user(Order order, DataLoader<String, User> userDataLoader) {
    return userDataLoader.load(order.getUserId()); // batched!
}
```

**In Federation context:** The Apollo Router has built-in entity batching. When it needs to resolve the same entity type multiple times, it automatically sends a single `_entities` query to the subgraph with all IDs, not N separate queries."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked as a follow-up** at Myntra and Swiggy. N+1 is the most common performance pitfall in GraphQL and interviewers specifically probe for it.

---

## Q4. How do you version a GraphQL API? How do you deprecate fields?

**"One of GraphQL's key selling points is that you rarely need to version. Instead, you evolve the schema:**

```graphql
type Product {
  id: ID!
  name: String!
  
  price: Float! @deprecated(reason: "Use 'pricing' instead for multi-currency support")
  
  pricing: Pricing  # New field — clients can migrate at their own pace
}

type Pricing {
  currency: String!
  amount: Float!
  displayPrice: String!
}
```

**Rules for non-breaking schema changes:**
- ✅ Adding new types, fields, or arguments with defaults
- ✅ Adding `@deprecated` annotation
- ❌ Removing fields (breaking — clients using them will error)
- ❌ Changing field type (String → Int is breaking)
- ❌ Making optional fields required

**Monitoring deprecated field usage:** In Apollo Studio, you can see which clients are still using `@deprecated` fields in production, letting you identify who to migrate before removal.

**When you MUST break the API (rare):**  
Use the `@link` import version in operationName or a separate endpoint:  
`/graphql/v2` for radical schema changes — but this defeats the purpose of GraphQL."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Razorpay (external API surface that cannot break mobile clients). Schema governance is a real operational concern at scale.
