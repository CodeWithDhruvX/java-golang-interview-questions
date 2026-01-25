# Observer Pattern

## 🟢 What is it?
The **Observer Pattern** defines a subscription mechanism to notify multiple objects about any events that happen to the object they're observing.

Think of it like **YouTube Subscription**.

---

## 🎯 Strategy to Implement

1.  **Observer Interface**: `Update(msg)`
2.  **Subject Interface**: `Subscribe(o)`, `Unsubscribe(o)`, `Notify()`
3.  **Concrete Subject**: List of subscribers.
4.  **Concrete Observers**: Implement Update.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Observer Interface
type Observer interface {
    Update(news string)
}

// 2. Subject Interface
type NewsAgency interface {
    Subscribe(o Observer)
    Unsubscribe(o Observer)
    NotifyObservers()
}

// 3. Concrete Subject
type CNN struct {
    channels     []Observer
    breakingNews string
}

func (c *CNN) Subscribe(o Observer) {
    c.channels = append(c.channels, o)
}

func (c *CNN) Unsubscribe(o Observer) {
    for i, obs := range c.channels {
        if obs == o {
            // Remove from slice
            c.channels = append(c.channels[:i], c.channels[i+1:]...)
            break
        }
    }
}

func (c *CNN) NotifyObservers() {
    for _, o := range c.channels {
        o.Update(c.breakingNews)
    }
}

func (c *CNN) SetBreakingNews(news string) {
    c.breakingNews = news
    fmt.Println("CNN reports:", news)
    c.NotifyObservers()
}

// 4. Concrete Observers
type MobileApp struct {
    name string
}

func (m *MobileApp) Update(news string) {
    fmt.Printf("App Notification on %s: %s\n", m.name, news)
}

type EmailSubscriber struct {
    email string
}

func (e *EmailSubscriber) Update(news string) {
    fmt.Printf("Email sent to %s: %s\n", e.email, news)
}

func main() {
    cnn := &CNN{}

    phone := &MobileApp{name: "iPhone 15"}
    email := &EmailSubscriber{email: "john@example.com"}

    cnn.Subscribe(phone)
    cnn.Subscribe(email)

    cnn.SetBreakingNews("Aliens landed in New York!")
    
    cnn.Unsubscribe(email)
    
    cnn.SetBreakingNews("It was just a movie set.")
}
```

---

## ✅ When to use?

*   **Event-Driven Systems**: Changes in one object affect others.
*   **One-to-Many**: Decoupled notification.
