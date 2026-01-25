# Singleton Pattern

## 🟢 What is it?
The **Singleton Pattern** ensures that a type has **only one instance** and provides a **global point of access** to it.

Think of it like a **Government**: A country can have only one official government.

### Key Characteristics:
1.  **Single Instance**: The package is responsible for holding the one instance.
2.  **Global Access**: The instance is globally accessible via a public function (usually `GetInstance()`).
3.  **Private Instantiation**: In Go, you can't strictly prevent struct creation if it's exported, but you can leave fields unexported or just follow convention.

---

## 🎯 Strategy to Implement

1.  **Private Struct**: Define an unexported struct `singleton` to discourage external creation.
2.  **Package Level Variable**: Hold the instance in a private variable.
3.  **sync.Once**: Use `sync.Once` to ensure thread-safe, lazy initialization.

---

## 💻 Code Example

Here is the Idiomatic Go Thread-Safe Singleton.

```go
package main

import (
    "fmt"
    "sync"
)

type databaseConnection struct {
    connectionString string
}

var (
    instance *databaseConnection
    once     sync.Once
)

// GetInstance ensures the instance is created only once
func GetInstance() *databaseConnection {
    // once.Do guarantees the function is called only once,
    // even if called concurrently from multiple goroutines.
    once.Do(func() {
        fmt.Println("Creating Database Connection...")
        instance = &databaseConnection{
            connectionString: "postgres://user:pass@localhost:5432/db",
        }
    })
    return instance
}

func (db *databaseConnection) Query(q string) {
    fmt.Println("Executing:", q)
}

func main() {
    // First call creates the instance
    db1 := GetInstance()
    db1.Query("SELECT * FROM users")

    // Second call returns existing instance
    db2 := GetInstance()

    if db1 == db2 {
        fmt.Println("Both are the same instance")
    }
}
```

---

## ✅ When to use?

*   **Resource Management**: Database Connection Pools, Loggers, Caches.
*   **Configuration**: App-wide settings.

### ❌ When NOT to use?
*   Do not use it just to create "globals". It makes testing harder because you can't easily swap the singleton with a mock implementation unless you put it behind an interface.
