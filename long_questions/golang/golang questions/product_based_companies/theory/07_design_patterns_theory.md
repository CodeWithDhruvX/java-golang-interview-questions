# 🗣️ Theory — Design Patterns in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is the Factory Pattern and how do you implement it in Go?"

> *"The Factory Pattern abstracts the creation of objects so the caller doesn't need to know which concrete type is being created. In Go, this is just a function — `NewNotification(type string) Notification` — that returns an interface. Depending on the type string, it creates and returns an email, SMS, or push notifier. The caller just works with the `Notification` interface without knowing what's behind it. It's valuable when you want to add new types without changing the calling code. In Go we don't need a `Factory` class — a plain function or even a map of constructor functions does the job more cleanly."*

---

## Q: "What is the Functional Options pattern? Why is it preferred in Go?"

> *"Functional options is probably the most Go-specific pattern you'll encounter. The problem it solves: you have a type with many optional configuration values. In Java you'd use a Builder class. In Go, the idiomatic approach is to define an `Option` type as a function — `type Option func(*Server)` — and expose factory functions for each option — `func WithPort(p int) Option`. Then your constructor takes `...Option` and applies each option to the struct. The advantages: options have names so they're self-documenting, adding new options is backward-compatible, you can have complex option logic, and the zero value defaults are always sensible. You see this in library APIs like the `grpc.Dial` options."*

---

## Q: "What is the Strategy Pattern and how is it naturally expressed in Go?"

> *"The Strategy Pattern says: define a family of algorithms, encapsulate each one, and make them interchangeable. In Go, this is trivially expressed using interfaces. You define an interface — `type SortStrategy interface { Sort([]int) []int }` — and create multiple implementations. Your context type holds a `SortStrategy` field which you can swap at runtime. What's beautiful is that in Go you don't even need to name the pattern — it just emerges naturally from interface design. Any time you have an interface field on a struct that can be swapped for different behaviors, you have the Strategy pattern."*

---

## Q: "What is the Repository Pattern and why do teams use it in Go?"

> *"The Repository Pattern isolates your data layer behind an interface. Instead of your business logic calling `db.Query(...)` directly, it calls `userRepo.GetByID(ctx, id)`. The interface defines what operations are available — Get, Save, Delete, List — without specifying how they're implemented. Benefits: first, you can swap the underlying database without changing business logic. Second and more importantly for day-to-day development, you can inject a mock repository in tests — no real database needed. This makes unit tests fast and reliable. It's a core part of Clean Architecture in Go: domain logic doesn't depend on infrastructure, infrastructure implements domain interfaces."*

---

## Q: "What is the Observer Pattern and how is it implemented in Go using channels?"

> *"The Observer Pattern has subjects that emit events and observers that subscribe to receive them. In Go, channels are a natural fit. You create an event bus with a map from event type to a slice of subscriber channels. Subscribing means adding a new channel to a topic's list. Publishing means iterating through all channels for a topic and sending the event — using a non-blocking select with a default case so a slow subscriber doesn't block the publisher. This is the message-bus pattern you see in real-time systems. The advantage over a traditional callback-based observer is that the publisher and subscriber are naturally decoupled — they don't share any code, just a channel."*

---

## Q: "What is Clean Architecture in Go?"

> *"Clean Architecture — also called Hexagonal or Onion Architecture — organizes code into concentric layers where inner layers have no knowledge of outer layers. In Go, I think of it as: the innermost layer is domain — your core business logic, structs, and interfaces. It has zero dependencies. The next layer is use-cases or services — they orchestrate domain logic. They depend on domain interfaces, not on concrete implementations. The outermost layer is infrastructure — database adapters, HTTP handlers, gRPC servers. They implement the domain interfaces. The rule is: dependencies point inward. HTTP handlers call services; services call repository interfaces; repository implementations call the database. This makes the business logic testable without any infrastructure."*

---

## Q: "What is the Builder Pattern in Go and when is it useful?"

> *"The Builder Pattern constructs complex objects step by step, separating the construction logic from the final object. In Go, it's usually implemented as a method-chaining struct where each method modifies the builder and returns it — `query.Where('age > ?', 18).OrderBy('name').Limit(10).Build()`. It's most useful for constructing SQL queries, HTTP request objects, or test fixtures. The key difference from functional options: a Builder accumulates state across multiple calls and produces a new object at the end via `Build()`, while functional options directly configure an existing object in its constructor."*
