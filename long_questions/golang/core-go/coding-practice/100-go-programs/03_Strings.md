# String Programs (Complete Collection - 55 Programs)

## 📚 Beginner Level (1-10) - Fundamentals

### 1. Count Occurrence of Each Character
**Principle**: Use frequency array or HashMap.
**Question**: Print count of each character in a string.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "programming"
    counted := make([]bool, len(str))
    
    for i := 0; i < len(str); i++ {
        if !counted[i] {
            ch := str[i]
            count := 1
            
            for j := i + 1; j < len(str); j++ {
                if ch == str[j] {
                    count++
                    counted[j] = true
                }
            }
            
            fmt.Printf("%c=%d\n", ch, count)
            counted[i] = true
        }
    }
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "programming"
    freq := make(map[rune]int)
    
    for _, ch := range str {
        freq[ch]++
    }
    
    for ch, count := range freq {
        fmt.Printf("%c=%d\n", ch, count)
    }
}

// func main() {
//     str := "programming"
//     freq := make(map[rune]int)
//     order := []rune{}

//     for _, ch := range str {
//         if freq[ch] == 0 {
//             order = append(order, ch) // track first appearance
//         }
//         freq[ch]++
//     }

//     for _, ch := range order {
//         fmt.Printf("%c=%d\n", ch, freq[ch])
//     }
// }
```

### 2. Find Length Without len()
**Principle**: Convert to rune array and count.
**Question**: Find string length without using len() method.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "Hello"
    count := 0
    
    for {
        if count >= len(str) {
            break
        }
        count++
    }
    
    fmt.Printf("Length: %d\n", count)
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "Hello"
    count := 0
    for range str {
        count++
    }
    fmt.Printf("Length: %d\n", count)
}
```

## 📚 Intermediate Level (11-20)

### 3. First Repeated Character
**Principle**: Use HashSet to track seen characters.
**Question**: Find first repeated character in a string.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "swiss"
    
    for i := 0; i < len(str); i++ {
        for j := i + 1; j < len(str); j++ {
            if str[i] == str[j] {
                fmt.Printf("First repeated: %c\n", str[i])
                return
            }
        }
    }
    fmt.Println("No repeated character")
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "swiss"
    seen := make(map[rune]bool)
    
    for _, ch := range str {
        if seen[ch] {
            fmt.Printf("First repeated: %c\n", ch)
            return
        }
        seen[ch] = true
    }
    fmt.Println("No repeated character")
}
```

### 4. Remove Duplicate Characters
**Principle**: Use StringBuilder with visited check.
**Question**: Remove duplicate characters from a string.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "programming"
    result := ""
    
    for i := 0; i < len(str); i++ {
        isDuplicate := false
        
        for j := 0; j < len(result); j++ {
            if str[i] == result[j] {
                isDuplicate = true
                break
            }
        }
        
        if !isDuplicate {
            result += string(str[i])
        }
    }
    
    fmt.Printf("After removing duplicates: %s\n", result)
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "programming"
    var result strings.Builder
    visited := make(map[byte]bool)
    
    for i := 0; i < len(str); i++ {
        ch := str[i]
        if !visited[ch] {
            visited[ch] = true
            result.WriteByte(ch)
        }
    }
    fmt.Printf("After removing duplicates: %s\n", result.String())
}
```

## 📚 Advanced Level (21-30) - Very Common in Interviews

### 5. Longest Substring Without Repeating Characters
**Principle**: Sliding window with HashSet.
**Question**: Find length of longest substring without repeating characters.

**Brute Force Approach (O(n³))**:
```go
package main

import "fmt"

func hasUniqueChars(s string) bool {
    seen := make(map[rune]bool)
    for _, ch := range s {
        if seen[ch] {
            return false
        }
        seen[ch] = true
    }
    return true
}

func main() {
    str := "abcabcbb"
    maxLen := 0
    
    for i := 0; i < len(str); i++ {
        for j := i; j < len(str); j++ {
            substring := str[i : j+1]
            if hasUniqueChars(substring) {
                if len(substring) > maxLen {
                    maxLen = len(substring)
                }
            }
        }
    }
    
    fmt.Printf("Length of longest unique substring: %d\n", maxLen)
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "abcabcbb"
    maxLen := 0
    start := 0
    charIndex := make(map[rune]int)
    
    for i, ch := range str {
        if index, exists := charIndex[ch]; exists && index >= start {
            start = index + 1
        }
        charIndex[ch] = i
        if i-start+1 > maxLen {
            maxLen = i - start + 1
        }
    }
    
    fmt.Printf("Length of longest unique substring: %d\n", maxLen)
}
```

### 6. Check if Two Strings are Anagrams
**Principle**: Sort both strings or count frequency.
**Question**: Check if two strings are anagrams.

**Brute Force Approach (O(n log n))**:
```go
package main

import (
    "fmt"
    "sort"
    "strings"
)

func sortString(s string) string {
    chars := strings.Split(s, "")
    sort.Strings(chars)
    return strings.Join(chars, "")
}

func main() {
    str1 := "listen"
    str2 := "silent"
    
    if sortString(str1) == sortString(str2) {
        fmt.Printf("%s and %s are anagrams\n", str1, str2)
    } else {
        fmt.Printf("%s and %s are not anagrams\n", str1, str2)
    }
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func areAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    freq := make(map[rune]int)
    for _, ch := range s1 {
        freq[ch]++
    }
    for _, ch := range s2 {
        freq[ch]--
        if freq[ch] < 0 {
            return false
        }
    }
    return true
}

func main() {
    str1 := "listen"
    str2 := "silent"
    
    if areAnagrams(str1, str2) {
        fmt.Printf("%s and %s are anagrams\n", str1, str2)
    } else {
        fmt.Printf("%s and %s are not anagrams\n", str1, str2)
    }
}
```

### 7. Reverse Each Word in String
**Principle**: Split by space, reverse each word, join back.
**Question**: Reverse each word in a string while maintaining word order.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strings"
)

func reverseWord(word string) string {
    reversed := ""
    for i := len(word) - 1; i >= 0; i-- {
        reversed += string(word[i])
    }
    return reversed
}

func main() {
    str := "Hello World"
    words := strings.Split(str, " ")
    
    for i, word := range words {
        words[i] = reverseWord(word)
    }
    
    result := strings.Join(words, " ")
    fmt.Printf("Reversed words: %s\n", result)
}
```

**Optimized Approach (O(n))**:
```go
package main

import (
    "fmt"
    "strings"
)

func reverseWord(word string) string {
    runes := []rune(word)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    str := "Hello World"
    words := strings.Split(str, " ")
    
    for i, word := range words {
        words[i] = reverseWord(word)
    }
    
    result := strings.Join(words, " ")
    fmt.Printf("Reversed words: %s\n", result)
}
```

### 8. Check if String is Palindrome
**Principle**: Compare characters from start and end.
**Question**: Check if a string is a palindrome.

**Brute Force Approach (O(n))**:
```go
package main

import "fmt"

func isPalindrome(str string) bool {
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    str := "madam"
    if isPalindrome(str) {
        fmt.Printf("%s is a palindrome\n", str)
    } else {
        fmt.Printf("%s is not a palindrome\n", str)
    }
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func isPalindrome(str string) bool {
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}

func main() {
    str := "madam"
    if isPalindrome(str) {
        fmt.Printf("%s is a palindrome\n", str)
    } else {
        fmt.Printf("%s is not a palindrome\n", str)
    }
}
```

### 9. Count Vowels and Consonants
**Principle**: Check each character against vowel set.
**Question**: Count vowels and consonants in a string.

**Brute Force Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "Hello World"
    vowels := 0
    consonants := 0
    
    for _, ch := range str {
        switch ch {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            vowels++
        default:
            if ch != ' ' {
                consonants++
            }
        }
    }
    
    fmt.Printf("Vowels: %d, Consonants: %d\n", vowels, consonants)
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "Hello World"
    vowels := 0
    consonants := 0
    vowelSet := map[rune]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
    }
    
    for _, ch := range str {
        if vowelSet[ch] {
            vowels++
        } else if ch != ' ' {
            consonants++
        }
    }
    
    fmt.Printf("Vowels: %d, Consonants: %d\n", vowels, consonants)
}
```

### 10. Convert String to Title Case
**Principle**: Capitalize first letter of each word.
**Question**: Convert string to title case.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strings"
)

func toTitleCase(str string) string {
    words := strings.Split(str, " ")
    for i, word := range words {
        if len(word) > 0 {
            words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
        }
    }
    return strings.Join(words, " ")
}

func main() {
    str := "hello world"
    result := toTitleCase(str)
    fmt.Printf("Title case: %s\n", result)
}
```

**Optimized Approach (O(n))**:
```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

func toTitleCase(str string) string {
    var result strings.Builder
    capitalizeNext := true
    
    for _, ch := range str {
        if capitalizeNext && unicode.IsLetter(ch) {
            result.WriteRune(unicode.ToUpper(ch))
            capitalizeNext = false
        } else {
            result.WriteRune(unicode.ToLower(ch))
            if ch == ' ' {
                capitalizeNext = true
            }
        }
    }
    
    return result.String()
}

func main() {
    str := "hello world"
    result := toTitleCase(str)
    fmt.Printf("Title case: %s\n", result)
}
```
