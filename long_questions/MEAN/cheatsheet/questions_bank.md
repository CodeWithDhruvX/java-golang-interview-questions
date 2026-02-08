# MEAN Stack Interview Cheat Sheet (Frequently Asked Topics)

---

## 1. MEAN Stack Overview

* **M**ongoDB – NoSQL database
* **E**xpress.js – Web framework for Node.js
* **A**ngular – Frontend framework
* **N**ode.js – JavaScript runtime
* Uses **JavaScript / TypeScript** end-to-end

---

## 2. MongoDB (Database)

### Core Concepts

* Document-based (JSON/BSON)
* Collection ≈ Table
* Document ≈ Row
* Schema-less (flexible schema)

### Frequently Asked Topics

* CRUD operations (`insert`, `find`, `update`, `delete`)
* `_id` and ObjectId
* Indexes (single, compound)
* Aggregation pipeline
* `find()` vs `findOne()`
* Embedded vs Referenced documents
* MongoDB vs SQL

### Mongoose (ODM)

* Schema definition
* Models
* Validation
* Middleware (pre, post hooks)
* `populate()`

---

## 3. Express.js (Backend Framework)

### Core Concepts

* Minimal & flexible Node.js framework
* Handles routing and middleware

### Frequently Asked Topics

* Routing (`GET`, `POST`, `PUT`, `DELETE`)
* Middleware (application-level, router-level)
* `req`, `res`, `next`
* Error handling middleware
* REST API structure
* Express vs Node

---

## 4. Node.js (Runtime Environment)

### Core Concepts

* Single-threaded
* Event-driven
* Non-blocking I/O

### Frequently Asked Topics

* Event loop
* Asynchronous programming
* Callbacks vs Promises vs Async/Await
* `process` object
* `require` vs `import`
* npm & package.json
* Environment variables

---

## 5. Angular (Frontend Framework)

### Core Concepts

* Component-based architecture
* TypeScript-based
* Two-way data binding

### Frequently Asked Topics

* Components
* Modules
* Services
* Dependency Injection
* Directives (structural & attribute)
* Pipes
* Lifecycle hooks
* Routing & Guards
* Reactive Forms vs Template-driven Forms
* HTTPClient

---

## 6. REST API Concepts

* Stateless architecture
* HTTP methods
* Status codes (200, 201, 400, 401, 404, 500)
* Request & Response cycle
* JSON data format

---

## 7. Authentication & Authorization

* JWT (JSON Web Token)
* Token-based authentication
* Cookies vs Local Storage
* Password hashing (bcrypt)
* Role-based access control

---

## 8. Security (Common Interview Area)

* CORS
* HTTPS
* Input validation
* SQL/NoSQL Injection
* XSS & CSRF
* Helmet & rate limiting

---

## 9. Performance & Best Practices

* Indexing in MongoDB
* Pagination
* Caching
* Modular architecture
* Error handling strategy
* Environment-based configuration

---

## 10. Common Interview Comparison Questions

* MongoDB vs MySQL
* Express vs Koa
* Angular vs React
* Node.js vs Java
* REST vs GraphQL

---

## 11. Deployment & Tools

* Environment variables
* Build & production mode
* PM2
* Docker (basic idea)
* Cloud platforms (AWS, Azure, GCP)

---

## 12. Freshers / Junior MEAN Stack – Interview Focus (Concepts & Basics)

### What Interviewers Expect

* Clear understanding of **fundamentals**, not advanced optimization
* Ability to **explain concepts simply**
* Basic coding & API flow knowledge
* Awareness of real-world usage

---

## 13. Must-Know Questions (Very Frequently Asked)

### MongoDB

* What is MongoDB?
* Difference between MongoDB and MySQL
* What is a document?
* What is `_id`?
* What are collections?
* What is indexing?

### Node.js

* What is Node.js?
* Is Node single-threaded?
* What is the event loop?
* What is asynchronous programming?
* Difference between `require` and `import`

### Express.js

* What is Express?
* What is middleware?
* Types of middleware
* Difference between `app.use()` and routes
* What is REST API?

### Angular

* What is Angular?
* What is a component?
* What is a module?
* What are services?
* What is dependency injection?
* Difference between reactive and template-driven forms

---

## 14. Very Basic Coding Expectations

### Backend

* Create a simple Express server
* Create GET and POST APIs
* Connect MongoDB using Mongoose
* Perform basic CRUD operations

### Frontend

* Create a component
* Bind data in HTML
* Call API using HttpClient
* Display API data in UI

---

## 15. Basic Concept Explanations (One-Line Style)

* **MongoDB**: NoSQL document-based database
* **Node.js**: JavaScript runtime built on Chrome V8
* **Express**: Backend framework for Node.js
* **Angular**: Frontend framework for building SPAs
* **REST API**: Communication using HTTP methods
* **Middleware**: Function between request & response

---

## 16. Common HR + Tech Mixed Questions

* Why MEAN stack?
* What project have you worked on?
* What was your role?
* Biggest challenge faced?
* Difference between frontend & backend

---

## 17. Final Interview Tip for Freshers

If you can **explain concepts + show one small project**, you are interview-ready for service-based companies.
