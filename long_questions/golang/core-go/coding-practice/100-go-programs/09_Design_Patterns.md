# Bonus: Design Patterns for Interviews

**Principle**: Service-based companies often ask for "Singleton", "Factory", or "Builder" to check architectural understanding.

## 1. Factory Pattern
**Type**: Creational.
**Use Case**: Creating objects without exposing instantiation logic.
**Example**: Bank Account Types.

```go
package main

import "fmt"

type Account interface {
    View()
}

type Savings struct{}

func (s Savings) View() {
    fmt.Println("Savings Account")
}

type Current struct{}

func (c Current) View() {
    fmt.Println("Current Account")
}

type AccountFactory struct{}

func (af AccountFactory) GetAccount(accountType string) Account {
    switch accountType {
    case "SAVINGS":
        return Savings{}
    case "CURRENT":
        return Current{}
    default:
        return nil
    }
}

func main() {
    factory := AccountFactory{}
    acc := factory.GetAccount("SAVINGS")
    if acc != nil {
        acc.View()
    }
}
```

## 2. Builder Pattern
**Type**: Creational.
**Use Case**: Constructing complex objects step-by-step.
**Example**: User object with optional fields.

```go
package main

import "fmt"

type User struct {
    Name string // Required
    Age  int    // Optional
}

type UserBuilder struct {
    name string
    age  int
    ageSet bool
}

func NewUserBuilder(name string) *UserBuilder {
    return &UserBuilder{name: name}
}

func (ub *UserBuilder) Age(age int) *UserBuilder {
    ub.age = age
    ub.ageSet = true
    return ub
}

func (ub *UserBuilder) Build() User {
    user := User{Name: ub.name}
    if ub.ageSet {
        user.Age = ub.age
    }
    return user
}

func main() {
    user := NewUserBuilder("John").Age(25).Build()
    fmt.Printf("User: %+v\n", user)
}
```

## 3. Observer Pattern
**Type**: Behavioral.
**Use Case**: Notify subscribers of changes.
**Example**: Youtube Channel.

```go
package main

import "fmt"

type Observer interface {
    Update(msg string)
}

type Subscriber struct {
    name string
}

func (s Subscriber) Update(msg string) {
    fmt.Printf("%s received: %s\n", s.name, msg)
}

type Channel struct {
    subs []Observer
}

func (c *Channel) Subscribe(s Observer) {
    c.subs = append(c.subs, s)
}

func (c *Channel) Upload(title string) {
    for _, s := range c.subs {
        s.Update("New Video: " + title)
    }
}

func main() {
    ch := &Channel{}
    ch.Subscribe(Subscriber{name: "Alice"})
    ch.Subscribe(Subscriber{name: "Bob"})
    
    ch.Upload("Go Design Patterns")
}
```

## 4. Strategy Pattern
**Type**: Behavioral.
**Use Case**: Switch algorithms at runtime.
**Example**: Payment methods.

```go
package main

import "fmt"

type Payment interface {
    Pay(amount int)
}

type Card struct{}

func (c Card) Pay(amount int) {
    fmt.Printf("Paid %d via Card\n", amount)
}

type UPI struct{}

func (u UPI) Pay(amount int) {
    fmt.Printf("Paid %d via UPI\n", amount)
}

type Cart struct {
    paymentMethod Payment
}

func (c *Cart) SetPayment(p Payment) {
    c.paymentMethod = p
}

func (c *Cart) Checkout(amount int) {
    c.paymentMethod.Pay(amount)
}

func main() {
    cart := &Cart{}
    cart.SetPayment(UPI{})
    cart.Checkout(100)
}
```

## 5. Singleton Pattern
**Type**: Creational.
**Use Case**: Ensure only one instance of a class exists.
**Example**: Database connection pool.

```go
package main

import (
    "fmt"
    "sync"
)

type Database struct {
    connections int
}

var (
    instance *Database
    once     sync.Once
)

func GetDatabaseInstance() *Database {
    once.Do(func() {
        instance = &Database{connections: 10}
    })
    return instance
}

func (db *Database) GetConnectionCount() int {
    return db.connections
}

func main() {
    db1 := GetDatabaseInstance()
    db2 := GetDatabaseInstance()
    
    fmt.Printf("Same instance? %t\n", db1 == db2)
    fmt.Printf("Connections: %d\n", db1.GetConnectionCount())
}
```

## 6. Adapter Pattern
**Type**: Structural.
**Use Case**: Allow incompatible interfaces to work together.
**Example**: Legacy system integration.

```go
package main

import "fmt"

// Legacy interface
type LegacyPrinter interface {
    PrintLegacy(s string)
}

type LegacyPrinterImpl struct{}

func (lp LegacyPrinterImpl) PrintLegacy(s string) {
    fmt.Printf("LEGACY: %s\n", s)
}

// New interface
type ModernPrinter interface {
    Print(s string)
}

// Adapter
type PrinterAdapter struct {
    legacy LegacyPrinter
}

func (pa PrinterAdapter) Print(s string) {
    pa.legacy.PrintLegacy(s)
}

func main() {
    legacy := LegacyPrinterImpl{}
    modern := PrinterAdapter{legacy: legacy}
    
    modern.Print("Hello World")
}
```

## 7. Decorator Pattern
**Type**: Structural.
**Use Case**: Add new functionality to objects dynamically.
**Example**: Coffee with extras.

```go
package main

import "fmt"

type Coffee interface {
    Cost() float64
    Description() string
}

type SimpleCoffee struct{}

func (sc SimpleCoffee) Cost() float64 {
    return 2.0
}

func (sc SimpleCoffee) Description() string {
    return "Simple Coffee"
}

type MilkDecorator struct {
    coffee Coffee
}

func (md MilkDecorator) Cost() float64 {
    return md.coffee.Cost() + 0.5
}

func (md MilkDecorator) Description() string {
    return md.coffee.Description() + ", Milk"
}

type SugarDecorator struct {
    coffee Coffee
}

func (sd SugarDecorator) Cost() float64 {
    return sd.coffee.Cost() + 0.2
}

func (sd SugarDecorator) Description() string {
    return sd.coffee.Description() + ", Sugar"
}

func main() {
    coffee := SimpleCoffee{}
    coffeeWithMilk := MilkDecorator{coffee: coffee}
    coffeeWithMilkAndSugar := SugarDecorator{coffee: coffeeWithMilk}
    
    fmt.Printf("%s: $%.2f\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())
}
```

## 8. Command Pattern
**Type**: Behavioral.
**Use Case**: Encapsulate requests as objects.
**Example**: Remote control.

```go
package main

import "fmt"

type Command interface {
    Execute()
}

type Light struct {
    isOn bool
}

func (l *Light) TurnOn() {
    l.isOn = true
    fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
    l.isOn = false
    fmt.Println("Light is OFF")
}

type LightOnCommand struct {
    light *Light
}

func (loc LightOnCommand) Execute() {
    loc.light.TurnOn()
}

type LightOffCommand struct {
    light *Light
}

func (loc LightOffCommand) Execute() {
    loc.light.TurnOff()
}

type RemoteControl struct {
    command Command
}

func (rc *RemoteControl) SetCommand(cmd Command) {
    rc.command = cmd
}

func (rc *RemoteControl) PressButton() {
    rc.command.Execute()
}

func main() {
    light := &Light{}
    onCmd := LightOnCommand{light: light}
    offCmd := LightOffCommand{light: light}
    
    remote := &RemoteControl{}
    
    remote.SetCommand(onCmd)
    remote.PressButton()
    
    remote.SetCommand(offCmd)
    remote.PressButton()
}
```

## 9. Template Method Pattern
**Type**: Behavioral.
**Use Case**: Define the skeleton of an algorithm.
**Example**: Data processing pipeline.

```go
package main

import "fmt"

type DataProcessor interface {
    ReadData()
    ProcessData()
    SaveData()
}

type BaseProcessor struct{}

func (bp BaseProcessor) Process() {
    bp.ReadData()
    bp.ProcessData()
    bp.SaveData()
}

func (bp BaseProcessor) ReadData() {
    fmt.Println("Reading data...")
}

func (bp BaseProcessor) SaveData() {
    fmt.Println("Saving data...")
}

type CSVProcessor struct {
    BaseProcessor
}

func (cp CSVProcessor) ProcessData() {
    fmt.Println("Processing CSV data...")
}

type JSONProcessor struct {
    BaseProcessor
}

func (jp JSONProcessor) ProcessData() {
    fmt.Println("Processing JSON data...")
}

func main() {
    csvProc := CSVProcessor{}
    jsonProc := JSONProcessor{}
    
    fmt.Println("CSV Processing:")
    csvProc.Process()
    
    fmt.Println("\nJSON Processing:")
    jsonProc.Process()
}
```

## 10. Iterator Pattern
**Type**: Behavioral.
**Use Case**: Provide a way to access elements of an aggregate object.
**Example**: Custom collection.

```go
package main

import "fmt"

type Iterator interface {
    HasNext() bool
    Next() interface{}
}

type Container interface {
    GetIterator() Iterator
}

type NameCollection struct {
    names []string
}

func (nc *NameCollection) AddName(name string) {
    nc.names = append(nc.names, name)
}

func (nc *NameCollection) GetIterator() Iterator {
    return &NameIterator{names: nc.names, index: 0}
}

type NameIterator struct {
    names []string
    index int
}

func (ni *NameIterator) HasNext() bool {
    return ni.index < len(ni.names)
}

func (ni *NameIterator) Next() interface{} {
    if ni.HasNext() {
        name := ni.names[ni.index]
        ni.index++
        return name
    }
    return nil
}

func main() {
    collection := &NameCollection{}
    collection.AddName("Alice")
    collection.AddName("Bob")
    collection.AddName("Charlie")
    
    iterator := collection.GetIterator()
    
    for iterator.HasNext() {
        name := iterator.Next()
        fmt.Println(name)
    }
}
```
