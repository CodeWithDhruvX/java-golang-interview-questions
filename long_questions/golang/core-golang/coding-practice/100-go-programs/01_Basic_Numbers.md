# Basic Number Programs (1-15)

## 1. Fibonacci Series
**Principle**: Each number is the sum of the two preceding ones, starting from 0 and 1.
**Question**: Write a program to print the Fibonacci series up to `n` terms.
**Code**:
```go
package main

import "fmt"

func main() {
    n := 10
    first, second := 0, 1
    
    fmt.Printf("Fibonacci Series up to %d terms:\n", n)
    
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", first)
        next := first + second
        first = second
        second = next
    }
    fmt.Println()
}
```

## 2. Check Prime Number
**Principle**: A prime number is a number greater than 1 that has no positive divisors other than 1 and itself.
**Question**: Write a program to check if a given number is prime.
**Code**:
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    num := 29
    isPrime := true
    
    if num <= 1 {
        isPrime = false
    } else {
        for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
            if num%i == 0 {
                isPrime = false
                break
            }
        }
    }
    
    fmt.Printf("%d is prime? %t\n", num, isPrime)
}
```

## 3. Factorial of a Number
**Principle**: Factorial of `n` is the product of all positive integers less than or equal to `n`.
**Question**: Write a program to find the factorial of a number.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 5
    factorial := 1
    
    for i := 1; i <= num; i++ {
        factorial *= i
    }
    
    fmt.Printf("Factorial of %d = %d\n", num, factorial)
}
```

## 4. Palindrome Number
**Principle**: A number that remains the same when its digits are reversed.
**Question**: Check if a number is a Palindrome.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 121
    original := num
    reversed := 0
    
    for num != 0 {
        digit := num % 10
        reversed = reversed*10 + digit
        num /= 10
    }
    
    if original == reversed {
        fmt.Printf("%d is a palindrome.\n", original)
    } else {
        fmt.Printf("%d is not a palindrome.\n", original)
    }
}
```

## 5. Armstrong Number
**Principle**: An integer is an Armstrong number if the sum of its digits raised to the power of the number of digits equals the number itself. (e.g., 153 = 1^3 + 5^3 + 3^3).
**Question**: Check if a number is an Armstrong number.
**Code**:
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    num := 153
    original := num
    result := 0
    n := 0
    
    // Count digits
    temp := original
    for temp != 0 {
        temp /= 10
        n++
    }
    
    temp = original
    for temp != 0 {
        remainder := temp % 10
        result += int(math.Pow(float64(remainder), float64(n)))
        temp /= 10
    }
    
    if result == num {
        fmt.Printf("%d is an Armstrong number.\n", num)
    } else {
        fmt.Printf("%d is not an Armstrong number.\n", num)
    }
}
```

## 6. Reverse a Number
**Principle**: Extract last digit check, add to new number check.
**Question**: Reverse a given integer.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 1234
    reversed := 0
    
    for num != 0 {
        digit := num % 10
        reversed = reversed*10 + digit
        num /= 10
    }
    
    fmt.Printf("Reversed: %d\n", reversed)
}
```

## 7. Swap Two Numbers without Third Variable
**Principle**: Use arithmetic operations (+/- or */%) or bitwise XOR.
**Question**: Swap two numbers without using a temporary variable.
**Code**:
```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    
    // Using arithmetic operations
    a = a + b // 30
    b = a - b // 10
    a = a - b // 20
    
    fmt.Printf("a: %d, b: %d\n", a, b)
}
```

## 8. Check Even or Odd
**Principle**: Divisible by 2 (remainder 0) or use bitwise AND (`num & 1 == 0`).
**Question**: Check if a number is even or odd.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 11
    
    if num%2 == 0 {
        fmt.Printf("%d is Even\n", num)
    } else {
        fmt.Printf("%d is Odd\n", num)
    }
}
```

## 9. Check Leap Year
**Principle**: Divisible by 4 and (not divisible by 100 OR divisible by 400).
**Question**: Check if a year is a leap year.
**Code**:
```go
package main

import "fmt"

func main() {
    year := 2024
    leap := false
    
    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                leap = true
            } else {
                leap = false
            }
        } else {
            leap = true
        }
    } else {
        leap = false
    }
    
    if leap {
        fmt.Printf("%d is a leap year.\n", year)
    } else {
        fmt.Printf("%d is not a leap year.\n", year)
    }
}
```

## 10. Greatest of Three Numbers
**Principle**: Conditional checks.
**Question**: Find the largest among three numbers.
**Code**:
```go
package main

import "fmt"

func main() {
    a, b, c := 10, 20, 15
    
    var max int
    if a > b {
        if a > c {
            max = a
        } else {
            max = c
        }
    } else {
        if b > c {
            max = b
        } else {
            max = c
        }
    }
    
    fmt.Printf("Largest: %d\n", max)
}
```

## 11. Sum of Digits
**Principle**: Extract digits modulo 10 and add to sum.
**Question**: Calculate sum of digits of a number.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 1234
    sum := 0
    
    for num > 0 {
        sum += num % 10
        num /= 10
    }
    
    fmt.Printf("Sum: %d\n", sum)
}
```

## 12. GCD of Two Numbers
**Principle**: Euclidean algorithm: GCD(a, b) = GCD(b, a%b).
**Question**: Find GCD (HCF) of two numbers.
**Code**:
```go
package main

import "fmt"

func main() {
    n1, n2 := 81, 153
    
    for n2 != 0 {
        temp := n2
        n2 = n1 % n2
        n1 = temp
    }
    
    fmt.Printf("GCD: %d\n", n1)
}
```

## 13. LCM of Two Numbers
**Principle**: LCM * GCD = n1 * n2.
**Question**: Find LCM of two numbers.
**Code**:
```go
package main

import "fmt"

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    n1, n2 := 72, 120
    
    // Find GCD first
    gcdValue := gcd(n1, n2)
    
    // Calculate LCM
    lcm := (n1 * n2) / gcdValue
    
    fmt.Printf("LCM: %d\n", lcm)
}
```

## 14. Perfect Number
**Principle**: Sum of positive divisors excluding the number itself equals the number.
**Question**: Check if a number is a Perfect Number.
**Code**:
```go
package main

import "fmt"

func main() {
    num := 28
    sum := 0
    
    for i := 1; i < num; i++ {
        if num%i == 0 {
            sum += i
        }
    }
    
    if sum == num {
        fmt.Printf("%d is a Perfect Number.\n", num)
    } else {
        fmt.Printf("%d is not a Perfect Number.\n", num)
    }
}
```

## 15. Print ASCII Value
**Principle**: Cast character to rune and convert to int.
**Question**: Print ASCII value of a character.
**Code**:
```go
package main

import "fmt"

func main() {
    ch := 'a'
    fmt.Printf("ASCII of %c is %d\n", ch, ch)
}
```
