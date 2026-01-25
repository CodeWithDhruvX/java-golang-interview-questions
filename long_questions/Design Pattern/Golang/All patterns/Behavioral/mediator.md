# Mediator Pattern

## 🟢 What is it?
The **Mediator Pattern** defines an object that encapsulates how a set of objects interact. It promotes loose coupling.

Think of it like an **Air Traffic Controller (ATC)**.

---

## 🎯 Strategy to Implement

1.  **Mediator Interface**: Declare methods for communication.
2.  **Concrete Mediator**: Store references to all components.
3.  **Components**: Store a reference to the Mediator. Call `mediator.Notify()`.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Mediator Interface
type ChatMediator interface {
    SendMessage(msg string, user User)
    AddUser(user User)
}

// 2. Concrete Mediator
type ChatRoomImpl struct {
    users []User
}

func (c *ChatRoomImpl) AddUser(user User) {
    c.users = append(c.users, user)
}

func (c *ChatRoomImpl) SendMessage(msg string, user User) {
    for _, u := range c.users {
        // Message should not be received by the user sending it
        if u != user {
            u.Receive(msg)
        }
    }
}

// 3. Components
type User interface {
    Send(msg string)
    Receive(msg string)
}

type ConcreteUser struct {
    mediator ChatMediator
    Name     string
}

func NewConcreteUser(med ChatMediator, name string) *ConcreteUser {
    return &ConcreteUser{mediator: med, Name: name}
}

func (u *ConcreteUser) Send(msg string) {
    fmt.Printf("%s Sending Message: %s\n", u.Name, msg)
    u.mediator.SendMessage(msg, u)
}

func (u *ConcreteUser) Receive(msg string) {
    fmt.Printf("%s Received Message: %s\n", u.Name, msg)
}

func main() {
    chatRoom := &ChatRoomImpl{}

    user1 := NewConcreteUser(chatRoom, "Pankaj")
    user2 := NewConcreteUser(chatRoom, "Lisa")
    user3 := NewConcreteUser(chatRoom, "Saurabh")
    user4 := NewConcreteUser(chatRoom, "David")

    chatRoom.AddUser(user1)
    chatRoom.AddUser(user2)
    chatRoom.AddUser(user3)
    chatRoom.AddUser(user4)

    user1.Send("Hello World!")
}
```

---

## ✅ When to use?

*   **Complex Communication**: When many-to-many relationships exist.
*   **Coupling**: When components depend on too many others.
