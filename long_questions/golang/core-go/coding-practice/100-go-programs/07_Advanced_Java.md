# Advanced Go Programs (86-100)

## 86. Function Types & Closures
**Principle**: Function as first-class citizen.
**Question**: Use a function expression.
**Code**:
```go
package main

import "fmt"

type Calc func(int, int) int

func main() {
    add := func(a, b int) int { return a + b }
    fmt.Println(add(5, 3))
}
```

## 87. Functional Options Pattern
**Principle**: Use variadic functions for configuration.
**Question**: Filter even numbers and square them.
**Code**:
```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4}
    
    // Filter even numbers
    var evens []int
    for _, n := range numbers {
        if n%2 == 0 {
            evens = append(evens, n)
        }
    }
    
    // Square them
    var result []int
    for _, n := range evens {
        result = append(result, n*n)
    }
    
    fmt.Println(result)
}
```

## 88. Count Elements using Function
**Principle**: Use higher-order functions.
**Question**: Count strings starting with 'A'.
**Code**:
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    strings := []string{"Apple", "Banana", "Apricot"}
    count := 0
    
    for _, s := range strings {
        if strings.HasPrefix(s, "A") {
            count++
        }
    }
    
    fmt.Println(count)
}
```

## 89. Max/Min using Built-in
**Principle**: Use sort package or manual iteration.
**Question**: Find max element.
**Code**:
```go
package main

import "fmt"

func main() {
    numbers := []int{1, 5, 3}
    if len(numbers) == 0 {
        return
    }
    
    max := numbers[0]
    for _, n := range numbers[1:] {
        if n > max {
            max = n
        }
    }
    
    fmt.Println(max)
}
```

## 90. Find Duplicate Numbers
**Principle**: Use map for tracking.
**Question**: Find duplicates.
**Code**:
```go
package main

import "fmt"

func main() {
    list := []int{1, 2, 1, 3}
    seen := make(map[int]bool)
    
    for _, n := range list {
        if seen[n] {
            fmt.Println(n)
        } else {
            seen[n] = true
        }
    }
}
```

## 91. Grouping using Map
**Principle**: Use map as grouping mechanism.
**Question**: Group strings by length.
**Code**:
```go
package main

import "fmt"

func main() {
    strings := []string{"A", "BB", "C"}
    groups := make(map[int][]string)
    
    for _, s := range strings {
        length := len(s)
        groups[length] = append(groups[length], s)
    }
    
    fmt.Println(groups)
}
```

## 92. Producer-Consumer Pattern
**Principle**: Use channels and goroutines.
**Question**: Implement Producer-Consumer.
**Code**:
```go
package main

import (
    "fmt"
    "sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("Put: %d\n", i)
    }
    close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for n := range ch {
        fmt.Printf("Got: %d\n", n)
    }
}

func main() {
    ch := make(chan int, 2)
    var wg sync.WaitGroup
    
    wg.Add(2)
    go producer(ch, &wg)
    go consumer(ch, &wg)
    
    wg.Wait()
}
```

## 93. Deadlock Scenario
**Principle**: Two goroutines holding locks waiting for each other.
**Question**: Create a deadlock.
**Code**:
```go
package main

import (
    "sync"
    "time"
)

func main() {
    var mu1, mu2 sync.Mutex
    
    // Goroutine 1 locks mu1 then mu2
    go func() {
        mu1.Lock()
        time.Sleep(100 * time.Millisecond)
        mu2.Lock()
        mu1.Unlock()
        mu2.Unlock()
    }()
    
    // Goroutine 2 locks mu2 then mu1
    go func() {
        mu2.Lock()
        time.Sleep(100 * time.Millisecond)
        mu1.Lock()
        mu2.Unlock()
        mu1.Unlock()
    }()
    
    time.Sleep(1 * time.Second)
}
```

## 94. Goroutine Creation
**Principle**: Use `go` keyword.
**Question**: Create and start a goroutine.
**Code**:
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Running")
    }()
    
    time.Sleep(100 * time.Millisecond) // Wait for goroutine
}
```

## 95. Defer for Resource Cleanup
**Principle**: Defer statement for cleanup.
**Question**: Read file using defer.
**Code**:
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer file.Close()
    
    // Use file
    fmt.Println("File opened successfully")
}
```

## 96. Pointers vs Nil Handling
**Principle**: Use pointers and nil checks.
**Question**: Use pointer to handle nil.
**Code**:
```go
package main

import "fmt"

func main() {
    var s *string
    if s == nil {
        defaultStr := "Default"
        s = &defaultStr
    }
    fmt.Println(*s)
}
```

## 97. Time Package
**Principle**: Use time package.
**Question**: Get current date.
**Code**:
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Format("2006-01-02"))
}
```

## 98. Error Handling
**Principle**: Multiple error checks.
**Question**: Handle division by zero.
**Code**:
```go
package main

import "fmt"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    _, err := divide(1, 0)
    if err != nil {
        fmt.Println("Zero Div")
    }
}
```

## 99. Worker Pool Pattern
**Principle**: Use goroutine pool.
**Question**: Create worker pool.
**Code**:
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    jobs := make(chan int, 100)
    var wg sync.WaitGroup
    
    // Start workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    wg.Wait()
}
```

## 100. Reflection Package
**Principle**: Inspect types at runtime.
**Question**: Get type name via reflection.
**Code**:
```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    str := "hello"
    t := reflect.TypeOf(str)
    fmt.Println(t.Name())
}
```
