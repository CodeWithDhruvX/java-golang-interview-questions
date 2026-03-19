# String Programs (Complete Collection - 55 Programs)

## 📚 Beginner Level (1-10) - Fundamentals

### 1. Count Occurrence of Each Character
**Principle**: Use frequency map.
**Question**: Print count of each character in a string.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "programming"
    counted := make(map[int]bool)
    
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
            
            fmt.Printf("%c=%d ", ch, count)
            counted[i] = true
        }
    }
    fmt.Println()
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
        fmt.Printf("%c=%d ", ch, count)
    }
    fmt.Println()
}
```

### 2. Find Length Without len()
**Principle**: Convert to rune slice and count.
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
**Principle**: Use Map to track seen characters.
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
**Principle**: Use strings.Builder with visited check.
**Question**: Remove duplicate characters from a string.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "programming"
    var result strings.Builder
    
    for i := 0; i < len(str); i++ {
        isDuplicate := false
        
        for j := 0; j < result.Len(); j++ {
            if str[i] == result.String()[j] {
                isDuplicate = true
                break
            }
        }
        
        if !isDuplicate {
            result.WriteByte(str[i])
        }
    }
    
    fmt.Printf("After removing duplicates: %s\n", result.String())
}
```

**Optimized Approach (O(n))**:
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "programming"
    var result strings.Builder
    visited := make(map[rune]bool)
    
    for _, ch := range str {
        if !visited[ch] {
            visited[ch] = true
            result.WriteRune(ch)
        }
    }
    
    fmt.Printf("After removing duplicates: %s\n", result.String())
}
```

## 📚 Advanced Level (21-30) - Very Common in Interviews

### 5. Longest Substring Without Repeating Characters
**Principle**: Sliding window with Map.
**Question**: Find length of longest substring without repeating characters.

**Brute Force Approach (O(n³))**:
```go
package main

import "fmt"

func hasUniqueChars(s string, start, end int) bool {
    seen := make(map[rune]bool)
    for i := start; i <= end; i++ {
        ch := rune(s[i])
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
            if hasUniqueChars(str, i, j) {
                if j-i+1 > maxLen {
                    maxLen = j - i + 1
                }
            }
        }
    }
    fmt.Printf("Longest substring length: %d\n", maxLen)
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
    indexMap := make(map[rune]int)
    
    for i, ch := range str {
        if val, exists := indexMap[ch]; exists && val >= start {
            start = val + 1
        }
        indexMap[ch] = i
        if i-start+1 > maxLen {
            maxLen = i - start + 1
        }
    }
    
    fmt.Printf("Longest substring length: %d\n", maxLen)
}
```

### 6. String Compression
**Principle**: Count consecutive characters and build compressed string.
**Question**: Compress string like aaabbc → a3b2c1.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "aaabbc"
    var compressed strings.Builder
    i := 0
    
    for i < len(str) {
        ch := str[i]
        count := 1
        
        for j := i + 1; j < len(str); j++ {
            if str[j] == ch {
                count++
            } else {
                break
            }
        }
        
        compressed.WriteByte(ch)
        compressed.WriteString(strconv.Itoa(count))
        i += count
    }
    
    fmt.Printf("Compressed: %s\n", compressed.String())
}
```

**Optimized Approach (O(n))**:
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    str := "aaabbc"
    var compressed strings.Builder
    count := 1
    
    for i := 1; i < len(str); i++ {
        if str[i] == str[i-1] {
            count++
        } else {
            compressed.WriteByte(str[i-1])
            compressed.WriteString(strconv.Itoa(count))
            count = 1
        }
    }
    compressed.WriteByte(str[len(str)-1])
    compressed.WriteString(strconv.Itoa(count))
    
    fmt.Printf("Compressed: %s\n", compressed.String())
}
```

### 7. Print All Substrings
**Principle**: Nested loops for start and end indices.
**Question**: Print all possible substrings of a string.

**Brute Force Approach (O(n³))**:
```go
package main

import "fmt"

func main() {
    str := "ABC"
    
    for i := 0; i < len(str); i++ {
        for j := i; j < len(str); j++ {
            var substring strings.Builder
            for k := i; k <= j; k++ {
                substring.WriteByte(str[k])
            }
            fmt.Println(substring.String())
        }
    }
}
```

**Optimized Approach (O(n²))**:
```go
package main

import "fmt"

func main() {
    str := "ABC"
    
    for i := 0; i < len(str); i++ {
        for j := i + 1; j <= len(str); j++ {
            fmt.Println(str[i:j])
        }
    }
}
```

### 8. Check Balanced Parentheses
**Principle**: Use Stack to track opening brackets.
**Question**: Check if parentheses are balanced.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strings"
)

func isBalancedBruteForce(str string) bool {
    current := str
    
    for strings.Contains(current, "()") || strings.Contains(current, "[]") || strings.Contains(current, "{}") {
        current = strings.ReplaceAll(current, "()", "")
        current = strings.ReplaceAll(current, "[]", "")
        current = strings.ReplaceAll(current, "{}", "")
    }
    
    return current == ""
}

func main() {
    str := "{[()]}"
    fmt.Printf("Balanced? %t\n", isBalancedBruteForce(str))
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func isMatching(open, close byte) bool {
    return (open == '(' && close == ')') ||
           (open == '[' && close == ']') ||
           (open == '{' && close == '}')
}

func isBalanced(str string) bool {
    stack := make([]byte, 0)
    
    for i := 0; i < len(str); i++ {
        ch := str[i]
        
        if ch == '(' || ch == '[' || ch == '{' {
            stack = append(stack, ch)
        } else if ch == ')' || ch == ']' || ch == '}' {
            if len(stack) == 0 || !isMatching(stack[len(stack)-1], ch) {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}

func main() {
    str := "{[()]}"
    if isBalanced(str) {
        fmt.Println("Balanced")
    } else {
        fmt.Println("Not Balanced")
    }
}
```

### 9. Longest Palindrome Substring
**Principle**: Expand around center for each character.
**Question**: Find longest palindromic substring.

**Brute Force Approach (O(n³))**:
```go
package main

import "fmt"

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func main() {
    str := "babad"
    longest := ""
    
    for i := 0; i < len(str); i++ {
        for j := i; j < len(str); j++ {
            substring := str[i : j+1]
            if isPalindrome(substring) && len(substring) > len(longest) {
                longest = substring
            }
        }
    }
    fmt.Printf("Longest palindrome: %s\n", longest)
}
```

**Optimized Approach (O(n²))**:
```go
package main

import "fmt"

func expand(str string, left, right int) string {
    for left >= 0 && right < len(str) && str[left] == str[right] {
        left--
        right++
    }
    return str[left+1 : right]
}

func main() {
    str := "babad"
    longest := ""
    
    for i := 0; i < len(str); i++ {
        odd := expand(str, i, i)
        even := expand(str, i, i+1)
        
        if len(odd) > len(longest) {
            longest = odd
        }
        if len(even) > len(longest) {
            longest = even
        }
    }
    fmt.Printf("Longest palindrome: %s\n", longest)
}
```

### 10. Remove Adjacent Duplicates
**Principle**: Use Stack to remove consecutive duplicates.
**Question**: Remove adjacent duplicate characters.

**Brute Force Approach (O(n²))**:
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "abbaca"
    result := str
    changed := true
    
    for changed {
        changed = false
        var temp strings.Builder
        i := 0
        
        for i < len(result) {
            if i < len(result)-1 && result[i] == result[i+1] {
                i += 2 // Skip duplicate pair
                changed = true
            } else {
                temp.WriteByte(result[i])
                i++
            }
        }
        result = temp.String()
    }
    
    fmt.Printf("Result: %s\n", result)
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    str := "abbaca"
    stack := make([]byte, 0)
    
    for i := 0; i < len(str); i++ {
        ch := str[i]
        
        if len(stack) > 0 && stack[len(stack)-1] == ch {
            stack = stack[:len(stack)-1] // Pop
        } else {
            stack = append(stack, ch) // Push
        }
    }
    
    fmt.Printf("Result: %s\n", string(stack))
}
```

### 11. Check if Strings are Isomorphic
**Principle**: One-to-one mapping between characters.
**Question**: Check if two strings are isomorphic.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func isIsomorphicBruteForce(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    for i := 0; i < len(s1); i++ {
        ch1, ch2 := s1[i], s2[i]
        
        // Check if ch1 maps to consistent ch2
        for j := 0; j < i; j++ {
            if s1[j] == ch1 && s2[j] != ch2 {
                return false
            }
            if s2[j] == ch2 && s1[j] != ch1 {
                return false
            }
        }
    }
    return true
}

func main() {
    s1, s2 := "egg", "add"
    fmt.Printf("Isomorphic? %t\n", isIsomorphicBruteForce(s1, s2))
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func main() {
    s1, s2 := "egg", "add"
    
    if len(s1) != len(s2) {
        fmt.Println("Not Isomorphic")
        return
    }
    
    mapping := make(map[byte]byte)
    used := make(map[byte]bool)
    
    for i := 0; i < len(s1); i++ {
        ch1, ch2 := s1[i], s2[i]
        
        if mapped, exists := mapping[ch1]; exists {
            if mapped != ch2 {
                fmt.Println("Not Isomorphic")
                return
            }
        } else {
            if used[ch2] {
                fmt.Println("Not Isomorphic")
                return
            }
            mapping[ch1] = ch2
            used[ch2] = true
        }
    }
    fmt.Println("Isomorphic")
}
```

### 12. Check if Strings are One Edit Away
**Principle**: Check insert, delete, or replace scenarios.
**Question**: Check if two strings are one edit away.

**Brute Force Approach (O(n³))**:
```go
package main

import "fmt"

func isOneEditAwayBruteForce(s1, s2 string) bool {
    // Try all possible single edits on s1
    // 1. Replace each character
    for i := 0; i < len(s1); i++ {
        modified := s1[:i] + "a" + s1[i+1:]
        if modified == s2 {
            return true
        }
    }
    
    // 2. Delete each character
    for i := 0; i < len(s1); i++ {
        modified := s1[:i] + s1[i+1:]
        if modified == s2 {
            return true
        }
    }
    
    // 3. Insert each possible character
    for i := 0; i <= len(s1); i++ {
        for c := 'a'; c <= 'z'; c++ {
            modified := s1[:i] + string(c) + s1[i:]
            if modified == s2 {
                return true
            }
        }
    }
    
    return s1 == s2
}

func main() {
    s1, s2 := "pale", "ple"
    fmt.Printf("One edit away? %t\n", isOneEditAwayBruteForce(s1, s2))
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func isOneEditAway(s1, s2 string) bool {
    if s1 == s2 {
        return true
    }
    
    len1, len2 := len(s1), len(s2)
    if abs(len1-len2) > 1 {
        return false
    }
    
    shorter, longer := s1, s2
    if len1 > len2 {
        shorter, longer = s2, s1
    }
    
    foundDifference := false
    i, j := 0, 0
    
    for i < len(shorter) && j < len(longer) {
        if shorter[i] != longer[j] {
            if foundDifference {
                return false
            }
            foundDifference = true
            if len1 == len2 {
                i++
            }
        } else {
            i++
        }
        j++
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    s1, s2 := "pale", "ple"
    fmt.Printf("One edit away? %t\n", isOneEditAway(s1, s2))
}
```

## 📚 FAANG/Product-Based Questions (36-40)

### 13. Group Anagrams
**Principle**: Sort each string and use as key in Map.
**Question**: Group anagrams together.

**Brute Force Approach (O(n² * m log m))**:
```go
package main

import (
    "fmt"
    "sort"
    "strings"
)

func areAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    c1 := []rune(s1)
    c2 := []rune(s2)
    
    sort.Slice(c1, func(i, j int) bool { return c1[i] < c1[j] })
    sort.Slice(c2, func(i, j int) bool { return c2[i] < c2[j] })
    
    return strings.Join(string(c1), "") == strings.Join(string(c2), "")
}

func groupAnagramsBruteForce(words []string) [][]string {
    result := make([][]string, 0)
    used := make([]bool, len(words))
    
    for i := 0; i < len(words); i++ {
        if used[i] {
            continue
        }
        
        group := []string{words[i]}
        used[i] = true
        
        for j := i + 1; j < len(words); j++ {
            if !used[j] && areAnagrams(words[i], words[j]) {
                group = append(group, words[j])
                used[j] = true
            }
        }
        result = append(result, group)
    }
    return result
}

func main() {
    words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    groups := groupAnagramsBruteForce(words)
    
    for _, group := range groups {
        fmt.Println(group)
    }
}
```

**Optimized Approach (O(n * m log m))**:
```go
package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    groups := make(map[string][]string)
    
    for _, word := range words {
        chars := []rune(word)
        sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
        key := string(chars)
        
        groups[key] = append(groups[key], word)
    }
    
    for _, group := range groups {
        fmt.Println(group)
    }
}
```

### 14. Minimum Window Substring
**Principle**: Sliding window with character frequency.
**Question**: Find minimum window containing all characters of another string.

**Brute Force Approach (O(n²))**:
```go
package main

import "fmt"

func containsAllChars(window, t string) bool {
    windowCount := make(map[byte]int)
    for i := 0; i < len(window); i++ {
        windowCount[window[i]]++
    }
    
    tCount := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tCount[t[i]]++
    }
    
    for ch, count := range tCount {
        if windowCount[ch] < count {
            return false
        }
    }
    return true
}

func minWindowBruteForce(s, t string) string {
    result := ""
    minLen := 1<<31 - 1 // Max int
    
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            window := s[i : j+1]
            if containsAllChars(window, t) && len(window) < minLen {
                result = window
                minLen = len(window)
            }
        }
    }
    return result
}

func main() {
    s, t := "ADOBECODEBANC", "ABC"
    fmt.Printf("Minimum window: %s\n", minWindowBruteForce(s, t))
}
```

**Optimized Approach (O(n))**:
```go
package main

import "fmt"

func minWindow(s, t string) string {
    if len(s) == 0 || len(t) == 0 {
        return ""
    }
    
    need := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        need[t[i]]++
    }
    
    left, formed := 0, 0
    required := len(need)
    window := make(map[byte]int)
    ans := []int{-1, 0, 0} // length, left, right
    
    for right := 0; right < len(s); right++ {
        ch := s[right]
        window[ch]++
        
        if count, exists := need[ch]; exists && window[ch] == count {
            formed++
        }
        
        for left <= right && formed == required {
            c := s[left]
            if ans[0] == -1 || right-left+1 < ans[0] {
                ans[0], ans[1], ans[2] = right-left+1, left, right
            }
            
            window[c]--
            if count, exists := need[c]; exists && window[c] < count {
                formed--
            }
            left++
        }
    }
    
    if ans[0] == -1 {
        return ""
    }
    return s[ans[1] : ans[2]+1]
}

func main() {
    s, t := "ADOBECODEBANC", "ABC"
    fmt.Printf("Minimum window: %s\n", minWindow(s, t))
}
```

### 15. Decode String
**Principle**: Use Stack for numbers and strings.
**Question**: Decode encoded string like 3[a2[c]] → accaccacc.

**Brute Force Approach (O(n³))**:
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func decodeStringBruteForce(s string) string {
    result := s
    changed := true
    
    for changed {
        changed = false
        var temp strings.Builder
        i := 0
        
        for i < len(result) {
            if result[i] >= '0' && result[i] <= '9' {
                num := 0
                for i < len(result) && result[i] >= '0' && result[i] <= '9' {
                    num = num*10 + int(result[i]-'0')
                    i++
                }
                
                if i < len(result) && result[i] == '[' {
                    j := i + 1
                    count := 1
                    for j < len(result) && count > 0 {
                        if result[j] == '[' {
                            count++
                        }
                        if result[j] == ']' {
                            count--
                        }
                        j++
                    }
                    
                    substring := result[i+1 : j-1]
                    var expanded strings.Builder
                    for k := 0; k < num; k++ {
                        expanded.WriteString(substring)
                    }
                    temp.WriteString(expanded.String())
                    i = j
                    changed = true
                }
            } else {
                temp.WriteByte(result[i])
                i++
            }
        }
        result = temp.String()
    }
    
    return result
}

func main() {
    str := "3[a2[c]]"
    fmt.Printf("Decoded: %s\n", decodeStringBruteForce(str))
}
```

**Optimized Approach (O(n))**:
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func decodeString(s string) string {
    countStack := make([]int, 0)
    stringStack := make([]string, 0)
    var current strings.Builder
    k := 0
    
    for i := 0; i < len(s); i++ {
        ch := s[i]
        
        if ch >= '0' && ch <= '9' {
            k = k*10 + int(ch-'0')
        } else if ch == '[' {
            countStack = append(countStack, k)
            stringStack = append(stringStack, current.String())
            current.Reset()
            k = 0
        } else if ch == ']' {
            var decoded strings.Builder
            decoded.WriteString(stringStack[len(stringStack)-1])
            stringStack = stringStack[:len(stringStack)-1]
            
            count := countStack[len(countStack)-1]
            countStack = countStack[:len(countStack)-1]
            
            for j := 0; j < count; j++ {
                decoded.WriteString(current.String())
            }
            current = decoded
        } else {
            current.WriteByte(ch)
        }
    }
    return current.String()
}

func main() {
    str := "3[a2[c]]"
    fmt.Printf("Decoded: %s\n", decodeString(str))
}
```

### 16. Multiply Strings
**Principle**: Manual multiplication like grade school.
**Question**: Multiply two string numbers without converting to integer.

**Brute Force Approach (O(n³))**:
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func addStrings(a, b string) string {
    var result strings.Builder
    i, j, carry := len(a)-1, len(b)-1, 0
    
    for i >= 0 || j >= 0 || carry > 0 {
        digit1 := 0
        if i >= 0 {
            digit1 = int(a[i] - '0')
            i--
        }
        
        digit2 := 0
        if j >= 0 {
            digit2 = int(b[j] - '0')
            j--
        }
        
        sum := digit1 + digit2 + carry
        carry = sum / 10
        result.WriteByte(byte(sum%10 + '0'))
    }
    
    // Reverse the result
    res := result.String()
    var reversed strings.Builder
    for i := len(res) - 1; i >= 0; i-- {
        reversed.WriteByte(res[i])
    }
    
    return reversed.String()
}

func multiplyBruteForce(num1, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    result := "0"
    
    for i := len(num1) - 1; i >= 0; i-- {
        digit1 := int(num1[i] - '0')
        var temp strings.Builder
        carry := 0
        
        // Multiply digit1 with entire num2
        for j := len(num2) - 1; j >= 0; j-- {
            digit2 := int(num2[j] - '0')
            product := digit1*digit2 + carry
            carry = product / 10
            temp.WriteByte(byte(product%10 + '0'))
        }
        
        if carry > 0 {
            temp.WriteString(strconv.Itoa(carry))
        }
        
        // Add zeros based on position
        for k := 0; k < len(num1)-1-i; k++ {
            temp.WriteString("0")
        }
        
        // Reverse temp
        tempStr := temp.String()
        var reversed strings.Builder
        for k := len(tempStr) - 1; k >= 0; k-- {
            reversed.WriteByte(tempStr[k])
        }
        
        result = addStrings(result, reversed.String())
    }
    
    return result
}

func main() {
    num1, num2 := "123", "456"
    fmt.Printf("Product: %s\n", multiplyBruteForce(num1, num2))
}
```

**Optimized Approach (O(n²))**:
```go
package main

import "fmt"

func multiply(num1, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    m, n := len(num1), len(num2)
    result := make([]int, m+n)
    
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            mul := int(num1[i]-'0') * int(num2[j]-'0')
            sum := mul + result[i+j+1]
            
            result[i+j+1] = sum % 10
            result[i+j] += sum / 10
        }
    }
    
    // Skip leading zeros
    start := 0
    for start < len(result) && result[start] == 0 {
        start++
    }
    
    // Convert to string
    var resultStr strings.Builder
    for i := start; i < len(result); i++ {
        resultStr.WriteByte(byte(result[i] + '0'))
    }
    
    return resultStr.String()
}

func main() {
    num1, num2 := "123", "456"
    fmt.Printf("Product: %s\n", multiply(num1, num2))
}
```
