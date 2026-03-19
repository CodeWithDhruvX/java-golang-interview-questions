# Golang Readability & Clean Design Guide

> **Core Philosophy:** Simple is better than complex. Explicit is better than implicit.

---

## 1. The "Happy Path" Principle
The most important rule in Go readability is to keep the "successful" execution path aligned to the left margin.

**Bad (Deeply Nested):**
```go
func (c *Client) Process() error {
    if c.active {
        data, err := c.Fetch()
        if err == nil {
            return c.Save(data)
        } else {
            return err
        }
    }
    return errors.New("client inactive")
}
```

**Good (Guard Clauses):**
```go
func (c *Client) Process() error {
    if !c.active {
        return errors.New("client inactive")
    }

    data, err := c.Fetch()
    if err != nil {
        return err
    }

    return c.Save(data)
}
```

---

## 2. Idiomatic Naming (The "Go Way")
Go favors brevity in small scopes and precision in large scopes.

*   **Method Receivers:** Use 1–2 letters (e.g., `func (r *Repo) ...`). Never use `self` or `this`.
*   **Variable Scoping:** Short variable names are okay in small blocks (`i` for index, `v` for value).
*   **Exported Names:** Choose descriptive names like `BufferedScanner` instead of just `Scanner`.
*   **Acronyms:** Acronyms should be consistent in case (e.g., `ServeHTTP`, `APIClient`, not `ServeHttp`, `ApiClient`).

---

## 3. Explicit Error Handling
Go treats errors as values, not exceptions. This makes failure points clearly visible to the reader.

*   **Actionable Errors:** Don't just return `err`. Wrap it to give context: `return fmt.Errorf("failed to sync user %d: %w", userID, err)`.
*   **Sentinel Errors:** Use `errors.New` or `var ErrNotFound = errors.New("not found")` for logical failures that callers should check.

---

## 4. Interfaces and Abstraction
*   **Consumer Defines Interface:** In Go, the package that *uses* the dependency should define the interface, not the package that *provides* the implementation.
*   **Keep Interfaces Small:** Aim for "Single Method Interfaces" (e.g., `Writer`, `Reader`, `Closer`).
*   **Abstract ONLY what you need:** Don't create an interface with 20 methods if you only use 2.

---

## 5. Avoiding "Clever" Code
If your code requires a comment to explain *how* it works (instead of *why* it exists), it's probably too "clever."

*   **Avoid Reflection:** It breaks type safety and makes reading stack traces a nightmare.
*   **Avoid Large Init Functions:** They hide initialization logic and make unit tests nearly impossible to isolate.
*   **Avoid Naked Returns:** Except in very tiny (sub 5-line) functions, always return values explicitly.

---

## 6. Documentation (Godoc)
Readable code is documented code. Every exported type and function should have a doc comment.
```go
// UserStore provides thread-safe access to user data.
type UserStore struct { ... }

// FindByID retrieves a user from the cache.
func (s *UserStore) FindByID(id int) (*User, error) { ... }
```

---

## 7. Struct Alignment (Advanced)
Order fields from largest to smallest to minimize memory padding.
```go
type Optimized struct {
    BigField   [256]byte // 256 bytes
    MediumField int64    // 8 bytes
    BoolField  bool      // 1 byte
}
```

---

## Summary Checklist for Peer Review
1. [ ] Is the "Happy Path" on the left?
2. [ ] Are method receivers short (1–2 letters)?
3. [ ] Are errors wrapped with context?
4. [ ] Are interfaces as small as possible?
5. [ ] Is the code free of unnecessary `init()` or `reflect` usage?
