# Pattern Programs (56-65)

## 56. Right Triangle Star Pattern
**Principle**: Nested loops. Outer loop for rows, inner for columns.
**Question**: Print a right triangle of stars.
**Code**:
```go
package main

import "fmt"

/*
*
**
***
****
*/
func main() {
    n := 4
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

## 57. Left Triangle Star Pattern
**Principle**: Print spaces then stars.
**Question**: Print a left-aligned triangle.
**Code**:
```go
package main

import "fmt"

/*
   *
  **
 ***
****
*/
func main() {
    n := 4
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= i; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

## 58. Pyramid Star Pattern
**Principle**: Spaces decreasing, Stars increasing (odd numbers 1, 3, 5...).
**Question**: Print a pyramid.
**Code**:
```go
package main

import "fmt"

/*
  *
 ***
*****
*/
func main() {
    n := 3
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

## 59. Diamond Pattern
**Principle**: Combine normal pyramid and inverted pyramid.
**Question**: Print a diamond shape.
**Code**:
```go
package main

import "fmt"

func main() {
    n := 3
    // Upper
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    // Lower
    for i := n - 1; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

## 60. Number Triangle Pattern
**Principle**: Print `j` (column index).
**Question**: Print number triangle.
**Code**:
```go
package main

import "fmt"

/*
1
12
123
*/
func main() {
    n := 3
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(j)
        }
        fmt.Println()
    }
}
```

## 61. Checkered/Floyd's Triangle (0-1)
**Principle**: If `(i+j)` is even print 1, else 0.
**Question**: Print 0-1 triangle.
**Code**:
```go
package main

import "fmt"

/*
1
0 1
1 0 1
*/
func main() {
    n := 3
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            if (i+j)%2 == 0 {
                fmt.Print("1 ")
            } else {
                fmt.Print("0 ")
            }
        }
        fmt.Println()
    }
}
```

## 62. Pascal's Triangle
**Principle**: `Val = Val * (i-j) / j`.
**Question**: Print Pascal's Triangle.
**Code**:
```go
package main

import "fmt"

/*
 1
 1 1
 1 2 1
*/
func main() {
    n := 3
    for i := 0; i < n; i++ {
        for s := 0; s < n-i; s++ {
            fmt.Print(" ")
        }
        val := 1
        for j := 0; j <= i; j++ {
            fmt.Printf("%d ", val)
            val = val * (i - j) / (j + 1)
        }
        fmt.Println()
    }
}
```

## 63. Rhombus Pattern
**Principle**: Shifted square.
**Question**: Print a solid rhombus.
**Code**:
```go
package main

import "fmt"

/*
  ****
 ****
****
*/
func main() {
    n := 4
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for j := 1; j <= n; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

## 64. Hollow Square
**Principle**: Print star if boundary index, else space.
**Question**: Print a hollow square pattern.
**Code**:
```go
package main

import "fmt"

func main() {
    n := 4
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i == 1 || i == n || j == 1 || j == n {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
```

## 65. Spiral Pattern (Number Grid)
**Principle**: Min distance to edge.
**Question**: Print concentric number layers.
**Code**:
```go
package main

import "fmt"

func main() {
    n := 4 // Size
    len := 2*n - 1
    
    for i := 0; i < len; i++ {
        for j := 0; j < len; j++ {
            min := i
            if j < min {
                min = j
            }
            if len-i-1 < min {
                min = len - i - 1
            }
            if len-j-1 < min {
                min = len - j - 1
            }
            fmt.Printf("%d ", n-min)
        }
        fmt.Println()
    }
}
```
