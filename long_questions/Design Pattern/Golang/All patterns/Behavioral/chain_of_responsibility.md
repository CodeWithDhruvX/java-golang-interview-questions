# Chain of Responsibility Pattern

## 🟢 What is it?
The **Chain of Responsibility Pattern** lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Think of it like **Calling Customer Support**:
1.  **Receptionist** picks up.
2.  **Tech Support L1**.
3.  **Senior Engineer L2**.

---

## 🎯 Strategy to Implement

1.  **Handler Interface**: Declare a method for processing requests and a method to set the `next` handler.
2.  **Concrete Handlers**: Implement the logic. "If I can handle this request, do it. Else, call next.Handle()".

---

## 💻 Code Example

```go
package main

import "fmt"

const (
    INFO  = 1
    ERROR = 2
)

// 1. Handler Interface
type Logger interface {
    SetNext(l Logger)
    LogMessage(level int, message string)
}

// 2. Base Logic (Helper to reuse 'next' logic)
type BaseLogger struct {
    Next  Logger
    Level int
}

func (l *BaseLogger) SetNext(next Logger) {
    l.Next = next
}

func (l *BaseLogger) LogMessage(level int, message string) {
    // If this logger matches the level, print it
    if l.Level <= level {
        l.Write(message)
    }
    // Always pass to next
    if l.Next != nil {
        l.Next.LogMessage(level, message)
    }
}

// Emulate abstract method
func (l *BaseLogger) Write(message string) {}

// 3. Concrete Handlers
type ConsoleLogger struct {
    BaseLogger
}

func NewConsoleLogger(level int) *ConsoleLogger {
    return &ConsoleLogger{BaseLogger{Level: level}}
}

func (c *ConsoleLogger) Write(message string) {
    fmt.Println("Console (Standard Config): " + message)
}

// Override LogMessage to wire up the specific Write method 
// (Go embedding doesn't dynamically dispatch to the child's Write method from BaseLogger)
func (c *ConsoleLogger) LogMessage(level int, message string) {
    if c.Level <= level {
        c.Write(message)
    }
    if c.Next != nil {
        c.Next.LogMessage(level, message)
    }
}

type ErrorLogger struct {
    BaseLogger
}

func NewErrorLogger(level int) *ErrorLogger {
    return &ErrorLogger{BaseLogger{Level: level}}
}

func (e *ErrorLogger) Write(message string) {
    fmt.Println("Error Log (File System): " + message)
}

func (e *ErrorLogger) LogMessage(level int, message string) {
    if e.Level <= level {
        e.Write(message)
    }
    if e.Next != nil {
        e.Next.LogMessage(level, message)
    }
}

func main() {
    // Chain: Error(2) -> Console(1)
    
    errorLogger := NewErrorLogger(ERROR)
    consoleLogger := NewConsoleLogger(INFO)

    errorLogger.SetNext(consoleLogger)

    fmt.Println("--- Sending INFO ---")
    // ErrorLogger ignores it (1 < 2), passes to Console. Console prints.
    errorLogger.LogMessage(INFO, "This is an information.")

    fmt.Println("\n--- Sending ERROR ---")
    // ErrorLogger matches, prints. Passes to Console. Console prints.
    errorLogger.LogMessage(ERROR, "System Failure!")
}
```

---

## ✅ When to use?

*   **Multiple Handlers**: When more than one object may handle a request.
*   **Decoupling**: Sender doesn't need to know the receiver.
*   **Dynamic Handling**: Chain can be built at runtime.
