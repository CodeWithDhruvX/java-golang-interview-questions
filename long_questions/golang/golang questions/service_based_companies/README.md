# 🏢 Go Interview Questions — Service-Based Companies

> Companies like **TCS, Infosys, Wipro, Cognizant, Capgemini, Accenture, HCL, Tech Mahindra**, etc.

Service-based companies typically focus on **practical Go usage**, clean code, working with APIs, databases, and deployment pipelines. Interviews tend to be **moderate in depth** and emphasize day-to-day engineering skills over deep internals.

---

## 📂 Category Index

| # | File | Topic | Difficulty |
|---|------|--------|------------|
| 01 | [Basics & Syntax](./01_basics_and_syntax.md) | Variables, types, control flow, functions | 🟢 Easy |
| 02 | [OOP Concepts in Go](./02_oops_concepts.md) | Structs, interfaces, embedding, methods | 🟢 Easy–Medium |
| 03 | [Concurrency](./03_concurrency.md) | Goroutines, channels, WaitGroup, Mutex | 🟡 Medium |
| 04 | [Error Handling](./04_error_handling.md) | errors, panic/recover, custom errors | 🟢 Easy–Medium |
| 05 | [REST API & Web Dev](./05_rest_api_and_web.md) | net/http, Gin, middleware, JSON | 🟡 Medium |
| 06 | [Databases](./06_databases.md) | database/sql, GORM, transactions, migrations | 🟡 Medium |
| 07 | [Testing](./07_testing.md) | Unit tests, table-driven tests, mocks | 🟡 Medium |
| 08 | [DevOps & Deployment](./08_devops_and_deployment.md) | Docker, CI/CD, go build, environment config | 🟡 Medium |

---

## 🎯 Interview Focus Areas (Service Companies)

- ✅ Go fundamentals and syntax
- ✅ REST API building with `net/http` or Gin
- ✅ Working with relational databases (MySQL/PostgreSQL)
- ✅ Basic concurrency (goroutines + channels)
- ✅ Error handling patterns
- ✅ Writing unit tests
- ✅ Docker containerization
- ✅ Reading and writing JSON/XML

---

## 💡 Tips for Service Company Interviews

1. **Know the basics deeply** — var vs `:=`, slices vs arrays, nil handling
2. **Practice CRUD APIs** — most projects involve REST + DB
3. **Understand error handling** — Go's `error` interface, `fmt.Errorf`, panic/recover
4. **Know how to write tests** — `go test`, table-driven tests
5. **Dockerize a Go app** — multi-stage builds, env vars
