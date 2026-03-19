# Array Programs (16-35)

## 16. Largest Element in Array
**Principle**: Iterate and compare.
**Question**: Find the largest element in an array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{10, 50, 20, 90, 40}
    max := arr[0]
    
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    
    fmt.Printf("Largest: %d\n", max)
}
```

## 17. Smallest Element in Array
**Principle**: Iterate and compare.
**Question**: Find the smallest element in an array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{10, 5, 20, 90, 40}
    min := arr[0]
    
    for _, num := range arr {
        if num < min {
            min = num
        }
    }
    
    fmt.Printf("Smallest: %d\n", min)
}
```

## 18. Reverse an Array
**Principle**: Swap elements from start and end moving towards center.
**Question**: Reverse the elements of an array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    start, end := 0, len(arr)-1
    
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
    
    fmt.Printf("Reversed: %v\n", arr)
}
```

## 19. Sort Array (Bubble Sort)
**Principle**: Repeatedly swap adjacent elements if custom order is wrong.
**Question**: Sort an array using Bubble Sort.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    
    fmt.Printf("Sorted: %v\n", arr)
}
```

## 20. Linear Search
**Principle**: Check every element until match found.
**Question**: Implement Linear Search.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{10, 20, 30, 40, 50}
    key := 30
    index := -1
    
    for i, val := range arr {
        if val == key {
            index = i
            break
        }
    }
    
    fmt.Printf("%d found at index: %d\n", key, index)
}
```

## 21. Binary Search
**Principle**: Divide and conquer on sorted array.
**Question**: Implement Binary Search.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{10, 20, 30, 40, 50} // Must be sorted
    key := 40
    low, high := 0, len(arr)-1
    index := -1
    
    for low <= high {
        mid := low + (high-low)/2
        
        if arr[mid] == key {
            index = mid
            break
        } else if arr[mid] < key {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    fmt.Printf("%d found at index: %d\n", key, index)
}
```

## 22. Remove Duplicates from Sorted Array
**Principle**: Use two pointers.
**Question**: Remove duplicates from a sorted array in-place.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
    j := 0
    
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] != arr[i+1] {
            arr[j] = arr[i]
            j++
        }
    }
    arr[j] = arr[len(arr)-1] // Last element
    j++
    
    // Print unique elements
    for i := 0; i < j; i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()
}
```

## 23. Second Largest Number
**Principle**: Track `largest` and `secondLargest`.
**Question**: Find the second largest number in an array.
**Code**:
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    arr := []int{12, 35, 1, 10, 34, 1}
    largest, second := math.MinInt32, math.MinInt32
    
    for _, num := range arr {
        if num > largest {
            second = largest
            largest = num
        } else if num > second && num != largest {
            second = num
        }
    }
    
    fmt.Printf("Second Largest: %d\n", second)
}
```

## 24. Missing Number in Array
**Principle**: Sum of 1 to N minus Sum of Array.
**Question**: Find missing number in range 1 to N.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 4, 5, 6}
    n := 6 // Max number
    
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    
    for _, num := range arr {
        actualSum += num
    }
    
    fmt.Printf("Missing Number: %d\n", expectedSum-actualSum)
}
```

## 25. Merge Two Arrays
**Principle**: Create new array of size `len1 + len2`.
**Question**: Merge two arrays into one.
**Code**:
```go
package main

import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := []int{4, 5, 6}
    
    res := make([]int, len(a)+len(b))
    copy(res, a)
    copy(res[len(a):], b)
    
    fmt.Printf("%v\n", res)
}
```

## 26. Check if Arrays are Equal
**Principle**: Check lengths, then check elements.
**Question**: Check if two arrays are equal.
**Code**:
```go
package main

import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := []int{1, 2, 3}
    
    equal := len(a) == len(b)
    if equal {
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                equal = false
                break
            }
        }
    }
    
    fmt.Printf("Equal? %t\n", equal)
}
```

## 27. Find Common Elements
**Principle**: Nested loops or Map.
**Question**: Find common elements between two arrays.
**Code**:
```go
package main

import "fmt"

func main() {
    arr1 := []int{1, 2, 3, 4}
    arr2 := []int{3, 4, 5, 6}
    
    set := make(map[int]bool)
    for _, val := range arr1 {
        set[val] = true
    }
    
    fmt.Print("Common elements: ")
    for _, val := range arr2 {
        if set[val] {
            fmt.Printf("%d ", val)
        }
    }
    fmt.Println()
}
```

## 28. Left Rotate Array
**Principle**: Shift elements left by `n` positions.
**Question**: Left rotate an array by 1 position.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    first := arr[0]
    
    for i := 0; i < len(arr)-1; i++ {
        arr[i] = arr[i+1]
    }
    arr[len(arr)-1] = first
    
    fmt.Printf("%v\n", arr)
}
```

## 29. Right Rotate Array
**Principle**: Shift elements right by `n` positions.
**Question**: Right rotate an array by 1 position.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    last := arr[len(arr)-1]
    
    for i := len(arr) - 1; i > 0; i-- {
        arr[i] = arr[i-1]
    }
    arr[0] = last
    
    fmt.Printf("%v\n", arr)
}
```

## 30. Move Zeros to End
**Principle**: Track index of non-zero elements.
**Question**: Move all zeros to the end of the array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{0, 1, 0, 3, 12}
    index := 0
    
    // Move non-zero elements forward
    for _, num := range arr {
        if num != 0 {
            arr[index] = num
            index++
        }
    }
    
    // Fill remaining positions with zeros
    for index < len(arr) {
        arr[index] = 0
        index++
    }
    
    fmt.Printf("%v\n", arr)
}
```

## 31. Find Duplicate Elements
**Principle**: Use Map or count frequency.
**Question**: Find duplicates in an int array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 1, 4, 2}
    set := make(map[int]bool)
    
    fmt.Print("Duplicates: ")
    for _, num := range arr {
        if set[num] {
            fmt.Printf("%d ", num)
        } else {
            set[num] = true
        }
    }
    fmt.Println()
}
```

## 32. Frequency of Each Element
**Principle**: Use Map[Element]Frequency.
**Question**: Find the frequency of each element in an array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 2, 3, 1, 4, 2}
    freq := make(map[int]int)
    
    for _, num := range arr {
        freq[num]++
    }
    
    fmt.Println("Frequency:")
    for num, count := range freq {
        fmt.Printf("%d: %d\n", num, count)
    }
}
```

## 33. Odd and Even Numbers in Array
**Principle**: Modulo operator.
**Question**: Print odd and even numbers separately.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 5, 6, 3, 2}
    
    fmt.Print("Odd: ")
    for _, val := range arr {
        if val%2 != 0 {
            fmt.Printf("%d ", val)
        }
    }
    
    fmt.Print("\nEven: ")
    for _, val := range arr {
        if val%2 == 0 {
            fmt.Printf("%d ", val)
        }
    }
    fmt.Println()
}
```

## 34. Sum of Array Elements
**Principle**: Loop and accumulate.
**Question**: Calculate sum of all elements in an array.
**Code**:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    sum := 0
    
    for _, num := range arr {
        sum += num
    }
    
    fmt.Printf("Sum: %d\n", sum)
}
```

## 35. Sort Array in Descending Order
**Principle**: Use sort package with custom comparator or manual sort.
**Question**: Sort an array in descending order.
**Code**:
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{5, 2, 9, 1, 6}
    
    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
    
    fmt.Printf("%v\n", arr)
}
```
