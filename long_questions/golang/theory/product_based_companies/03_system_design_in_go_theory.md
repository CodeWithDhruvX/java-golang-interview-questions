# 🗣️ Theory — System Design in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How would you design a rate limiter in Go?"

> *"The most common rate limiting algorithm for interviews is the token bucket. The idea is: imagine a bucket that fills with tokens at a constant rate. Each request costs one token. If the bucket has tokens, the request goes through and you remove a token. If empty, the request is rejected or queued. In Go, I'd implement it as a struct with a mutex — storing total tokens, max tokens, refill rate, and the last refill time. On each request, I calculate how many tokens were added since the last check, add them up to the max, and then try to consume one. The mutex ensures concurrent requests don't race. For a production system, you'd use Redis with atomic operations so it works across multiple server instances."*

---

## Q: "How would you implement an LRU cache in Go?"

> *"LRU — Least Recently Used — evicts the item that was used least recently when the cache is full. The classic data structure is a combination of a hash map and a doubly linked list. The hash map gives O(1) lookups, the linked list maintains usage order — most recently used at the head, least recently used at the tail. On every get, you move the accessed node to the front. On a put, you insert at the front and if you're over capacity, evict the tail. In Go, the `container/list` package gives you a doubly linked list. You maintain a map from key to list element. With a mutex wrapped around everything, it's thread-safe."*

---

## Q: "What is the circuit breaker pattern? Why is it important in microservices?"

> *"In a distributed system, if Service A calls Service B and B is down or slow, A's goroutines pile up waiting, consuming resources, and A itself becomes unavailable — this is called cascading failure. A circuit breaker sits between services and monitors the error rate. It has three states: Closed — requests flow normally; Open — the service is tripping, all requests fail immediately without trying; and Half-Open — after a timeout, it allows one trial request. If the trial succeeds, the breaker closes. If it fails, it reopens. This stops the cascade: instead of goroutines piling up waiting for a broken service, they fail fast and the system stays responsive."*

---

## Q: "Explain consistent hashing and why it's used in distributed systems."

> *"Consistent hashing solves the problem of distributing data across N servers when N can change. With naive modulo hashing — `key % N` — when you add or remove a server, nearly all keys would remap to different servers, causing a massive cache invalidation storm. Consistent hashing arranges all servers on a virtual ring — each server occupies multiple positions on the ring using virtual nodes. A key is hashed to a point on the ring, and then assigned to the next server clockwise. When a server is added or removed, only keys that were 'between' the changed server's position and its predecessor need to move — a fraction of total keys. This is how systems like Redis Cluster, Cassandra, and DynamoDB distribute data."*

---

## Q: "How would you implement distributed locking in Go?"

> *"A distributed lock coordinates access to a resource across multiple server processes. Redis is the classic choice. You use `SET key value NX PX timeout` — SET only if it doesn't exist, with a time-to-live. If the SET succeeds, you own the lock. To release, you use a Lua script that deletes the key only if its value matches what you set — this prevents accidentally releasing someone else's lock if yours expired. The value should be a random UUID unique to this lock acquisition. Go's Redis client wraps this cleanly. For production, Redlock — using multiple Redis nodes — gives stronger guarantees, though it has controversy in the distributed systems community."*

---

## Q: "How would you design a URL shortener system in Go?"

> *"A URL shortener maps a long URL to a short code — like `bit.ly/abc123`. The core decisions: for generating the short code, you can use a counter — increment an ID and encode it in base62 — which gives short, predictable codes. Or use a hash of the URL — which is deterministic but has collision risk. The API is simple: POST the long URL, get back the short code; GET the short code, redirect to the long URL. Storage needs to be fast for reads — Redis for the short-to-long mapping with the database as source of truth. For scale, you add a CDN in front and geographically distribute Redis replicas. In Go, the HTTP layer is straightforward with Gin — the meat of the problem is the encoding and storage strategy."*
