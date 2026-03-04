# 🟣 GraphQL Interview Questions

A comprehensive collection of GraphQL interview questions organized for **product-based** and **service-based** company interviews. Covers fundamentals through advanced architecture and real-world production scenarios.

---

## 📁 Folder Structure

```
graphql/
├── README.md                        ← This file
│
├── service_based_companies/         ← For TCS, Infosys, Wipro, Capgemini, HCL (1–5 yrs)
│   ├── 01_graphql_basics_schema.md
│   ├── 02_graphql_resolvers_operations.md
│   ├── 03_graphql_java_springboot.md
│   └── 04_graphql_apollo_client_nodejs.md
│
├── product_based_companies/         ← For Amazon, Flipkart, Swiggy, Razorpay, Atlassian (3–8 yrs)
│   ├── 01_graphql_advanced_concepts.md
│   ├── 02_graphql_performance_schema_design.md
│   ├── 03_graphql_system_design_scenarios.md
│   ├── 04_graphql_custom_directives_monitoring.md
│   └── 05_graphql_vs_grpc_and_rest.md
│
└── (raw Q&A bank — 500 questions)
    ├── graphql_answers_1_50.md
    ├── graphql_answers_51_100.md
    ├── graphql_answers_101_150.md
    ├── graphql_answers_151_200.md
    ├── graphql_answers_201_250.md
    ├── graphql_answers_251_300.md
    ├── graphql_answers_301_350.md
    ├── graphql_answers_351_400.md
    ├── graphql_answers_401_450.md
    └── graphql_answers_451_500.md
```

---

## 🏢 Service-Based Companies

> Target: TCS, Infosys, Wipro, Capgemini, HCL, Cognizant | Experience: 1–5 years

| File | Topics Covered |
|---|---|
| [01_graphql_basics_schema.md](./service_based_companies/01_graphql_basics_schema.md) | What is GraphQL, REST vs GraphQL, schema components, resolvers, queries/mutations/subscriptions, fragments, variables & directives, introspection, context & auth |
| [02_graphql_resolvers_operations.md](./service_based_companies/02_graphql_resolvers_operations.md) | N+1 problem & DataLoader, error handling, schema-first vs code-first, cursor vs offset pagination, file uploads, `type` vs `input`, testing GraphQL APIs |
| [03_graphql_java_springboot.md](./service_based_companies/03_graphql_java_springboot.md) | Spring for GraphQL setup, `@QueryMapping` / `@MutationMapping` / `@SchemaMapping`, `@BatchMapping` for N+1, Spring Security integration, `GraphQlTester` for testing |
| [04_graphql_apollo_client_nodejs.md](./service_based_companies/04_graphql_apollo_client_nodejs.md) | Apollo Client setup in React, `useQuery` / `useMutation` / `useSubscription`, fetch policies, cache updates after mutations, Node.js + Express server setup, WebSocket subscriptions |

---

## 🚀 Product-Based Companies

> Target: Amazon, Google, Flipkart, Swiggy, Razorpay, Atlassian, ThoughtWorks, Postman | Experience: 3–8 years

| File | Topics Covered |
|---|---|
| [01_graphql_advanced_concepts.md](./product_based_companies/01_graphql_advanced_concepts.md) | Apollo Federation & subgraphs, query planning & `_entities`, production security (depth/cost limiting, persisted queries), subscriptions at scale with Redis, Apollo Client normalized cache, field-level authorization & GraphQL Shield |
| [02_graphql_performance_schema_design.md](./product_based_companies/02_graphql_performance_schema_design.md) | Schema design conventions, deprecation strategy, schema modularization, full Relay cursor-based pagination, global error handling & OpenTelemetry, schema versioning & Apollo Registry, BFF pattern vs Federation |
| [03_graphql_system_design_scenarios.md](./product_based_companies/03_graphql_system_design_scenarios.md) | Designing e-commerce GraphQL API (Flipkart/Amazon scale), debugging slow API (N+1, indexes, caching), real-time features (subscriptions vs polling vs SSE), multi-tenancy in GraphQL, GraphQL in Go with `gqlgen` |
| [04_graphql_custom_directives_monitoring.md](./product_based_companies/04_graphql_custom_directives_monitoring.md) | Custom `@auth`, `@rateLimit`, `@upper`, `@log`, `@featureFlag` directives, schema transformer pattern, Prometheus/Datadog/Apollo Studio monitoring, CI/CD schema validation with Rover CLI & graphql-inspector |
| [05_graphql_vs_grpc_and_rest.md](./product_based_companies/05_graphql_vs_grpc_and_rest.md) | Full REST vs GraphQL vs gRPC comparison, when to choose gRPC (streaming, perf, polyglot), when to choose GraphQL (browser, BFF, multi-client), GraphQL+gRPC hybrid architecture, REST to GraphQL migration (strangler fig) |

---

## 🎯 Study Strategy

### For Service-Based Company Interviews (MNC)
1. Start with **01_graphql_basics_schema.md** — understand schema, types, operations
2. Practice **02_graphql_resolvers_operations.md** — coding questions on DataLoader, error handling
3. Read **03_graphql_java_springboot.md** if you work with Spring Boot (Java roles)
4. Read **04_graphql_apollo_client_nodejs.md** if you work with React/Node.js (frontend roles)

### For Product-Based Company Interviews (Startup/Unicorn/FAANG)
1. Master Federation concepts in **01_graphql_advanced_concepts.md**
2. Deep-dive on pagination, schema design in **02_graphql_performance_schema_design.md**
3. Practice system design answers from **03_graphql_system_design_scenarios.md**
4. Study custom directives and CI/CD schema validation in **04_graphql_custom_directives_monitoring.md**
5. Nail the GraphQL vs gRPC vs REST comparison in **05_graphql_vs_grpc_and_rest.md** — very commonly asked in system design rounds

---

## 🔑 Key Topics Checklist

### Fundamentals (Service-Based)
- [ ] GraphQL vs REST (key differences, trade-offs)
- [ ] Schema types: scalars, objects, enums, interfaces, unions, input types
- [ ] Resolver chain: parent → args → context → info
- [ ] Queries, mutations, subscriptions
- [ ] Fragments (named vs inline)
- [ ] Variables and directives (`@include`, `@skip`, `@deprecated`)
- [ ] Introspection and `__typename`
- [ ] N+1 problem and DataLoader
- [ ] Error handling (`errors` array, custom error codes)
- [ ] Authentication via context
- [ ] Apollo Client: `useQuery`, `useMutation`, `useSubscription`, fetch policies
- [ ] Node.js + Express GraphQL server setup

### Advanced (Product-Based)
- [ ] Apollo Federation: `@key`, `@external`, `@requires`, `@provides`
- [ ] Query planning in Apollo Router
- [ ] Depth limiting and query cost analysis
- [ ] Persisted queries for CDN caching
- [ ] Cursor-based pagination (Relay spec)
- [ ] Redis PubSub for scalable subscriptions
- [ ] Schema versioning and non-breaking evolution
- [ ] Field-level authorization (GraphQL Shield)
- [ ] BFF (Backend for Frontend) pattern
- [ ] GraphQL in Go (`gqlgen`)
- [ ] Multi-tenancy patterns
- [ ] OpenTelemetry + Prometheus/Datadog observability
- [ ] Custom directives (`@auth`, `@rateLimit`, `@log`, `@featureFlag`)
- [ ] CI/CD schema validation (Rover CLI, graphql-inspector)
- [ ] GraphQL vs gRPC vs REST trade-off decision framework
- [ ] GraphQL + gRPC hybrid architecture

---

*Questions curated to crack interviews across all major Indian IT companies and global product companies.*
