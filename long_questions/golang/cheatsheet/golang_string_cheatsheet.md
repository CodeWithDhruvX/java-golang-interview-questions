# Golang String Manipulation Cheatsheet

Comprehensive guide for string programs, mutations, and algorithms frequently asked in interviews.

---

## 🟢 Generic Helpers

### Max & Min (Integers)
```go
func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}
```

---

## 📙 Standard Library Essentials

### 1. Conversions (`strconv`)
Essential for string <-> number conversions.
```go
import "strconv"

// String to Int
i, err := strconv.Atoi("-42")
i64, err := strconv.ParseInt("42", 10, 64)

// Int to String
s := strconv.Itoa(-42)
s64 := strconv.FormatInt(-42, 10)

// String to Float
f, err := strconv.ParseFloat("3.1415", 64)
```

### 2. Common Utilities (`strings`)
```go
import "strings"

// Split & Join
parts := strings.Split("a,b,c", ",")  // ["a", "b", "c"]
joined := strings.Join([]string{"a", "b"}, "-") // "a-b"

// Contains & Index
exists := strings.Contains("seafood", "foo") // true
idx := strings.Index("chicken", "ken")       // 4

// Cleaning
clean := strings.TrimSpace("  hello  ")      // "hello"
words := strings.Fields("  foo   bar  baz")  // ["foo", "bar", "baz"]
```

### 3. Strings.Builder (Efficiency)
Use for building strings in loops to avoid excessive allocation.
```go
var sb strings.Builder
for i := 0; i < 10; i++ {
    sb.WriteString("a")
}
result := sb.String() // "aaaaaaaaaa"
```

---

## 🟡 Basic String Operations

Strings in Go are **immutable** byte slices. 
- `len(s)` gives number of **bytes**, not characters.
- iterate with `range` to get **runes** (Unicode Code Points).

### 1. Reverse a String
Handles Unicode characters correctly by converting to `[]rune`.
```go
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### 2. Check Palindrome
Checks if string reads same forward and backward.
```go
func IsPalindrome(s string) bool {
    // Optional: Pre-process to lower case and remove non-alphanumeric if needed
    // s = strings.ToLower(s) 
    
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
```

### 3. Count Vowels & Consonants
```go
func CountVowels(s string) (vowels, consonants int) {
    s = strings.ToLower(s)
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            switch r {
            case 'a', 'e', 'i', 'o', 'u':
                vowels++
            default:
                consonants++
            }
        }
    }
    return
}
```

---

## 🟣 Basic String Mutations
Strings are immutable, so we must covert to `[]rune` (for Unicode safety) or use slicing to create new strings.

### 1. Update Character at Index
```go
func UpdateChar(s string, index int, newChar rune) string {
    runes := []rune(s)
    if index >= 0 && index < len(runes) {
        runes[index] = newChar
    }
    return string(runes)
}
```

### 2. Insert Character at Index
```go
func InsertChar(s string, index int, char rune) string {
    runes := []rune(s)
    if index < 0 || index > len(runes) {
        return s
    }
    
    // Create new slice with capacity + 1
    res := make([]rune, len(runes)+1)
    copy(res[:index], runes[:index])
    res[index] = char
    copy(res[index+1:], runes[index:])
    
    return string(res)
}
```

### 3. Delete Character at Index
```go
func DeleteChar(s string, index int) string {
    runes := []rune(s)
    if index < 0 || index >= len(runes) {
        return s
    }
    return string(append(runes[:index], runes[index+1:]...))
}
```

### 4. Start & End Operations
```go
import "strings"

func StartEndOps(s string) {
    // Check Prefix/Suffix
    strings.HasPrefix(s, "Go") // true/false
    strings.HasSuffix(s, "lang")
    
    // Remove Prefix/Suffix
    res := strings.TrimPrefix(s, "Go") 
    res = strings.TrimSuffix(s, "lang")
    
    // Remove characters from start (slicing)
    if len(s) > 0 {
        s = s[1:] // Remove first char (byte-wise!)
    }
    
    // Remove from end
    if len(s) > 0 {
        s = s[:len(s)-1] // Remove last char
    }
}
```

---

## 🔵 String Algorithms & Interview Questions

### 1. Anagram Check
Uses a frequency map (or array [26]int for English only) to count characters.
```go
func AreAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    counts := make(map[rune]int)
    
    for _, r := range s1 {
        counts[r]++
    }
    
    for _, r := range s2 {
        counts[r]--
        if counts[r] < 0 {
            return false
        }
    }
    
    return true
}
```

### 2. Group Anagrams
Sort runes of each string to use as a key. O(N * K log K).
```go
func GroupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    
    for _, str := range strs {
        // Create key by sorting characters
        runes := []rune(str)
        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })
        key := string(runes)
        groups[key] = append(groups[key], str)
    }
    
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
```

### 3. First Unique Character
Two passes: Count frequencies, then find first with count 1.
```go
func FirstUniqChar(s string) int { // returns index
    freq := make(map[rune]int)
    for _, r := range s {
        freq[r]++
    }
    
    for i, r := range s {
        if freq[r] == 1 {
            return i
        }
    }
    return -1
}
```

### 4. Longest Substring Without Repeating Characters
Sliding Window pattern. O(n) time.
```go
func LengthOfLongestSubstring(s string) int {
    seen := make(map[rune]int)
    start := 0
    maxLength := 0
    
    for i, r := range s {
        if lastIdx, ok := seen[r]; ok && lastIdx >= start {
            start = lastIdx + 1
        }
        seen[r] = i
        currentLength := i - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }
    return maxLength
}
```

### 5. Longest Common Prefix
Horizontal scanning. O(S) where S is sum of all characters.
```go
func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for !strings.HasPrefix(strs[i], prefix) {
            prefix = prefix[:len(prefix)-1]
            if len(prefix) == 0 {
                return ""
            }
        }
    }
    return prefix
}
```

### 6. Run Length Encoding
Compress string (e.g., "aaabb" -> "a3b2").
```go
func RunLengthEncoding(s string) string {
    if len(s) == 0 { return "" }
    
    var sb strings.Builder
    runes := []rune(s)
    count := 1
    
    for i := 1; i < len(runes); i++ {
        if runes[i] == runes[i-1] {
            count++
        } else {
            sb.WriteRune(runes[i-1])
            sb.WriteString(strconv.Itoa(count))
            count = 1
        }
    }
    
    // Write last group
    sb.WriteRune(runes[len(runes)-1])
    sb.WriteString(strconv.Itoa(count))
    
    return sb.String()
}
```

### 7. Is Subsequence
Check if `s` is a subsequence of `t`. O(T).
```go
func IsSubsequence(s string, t string) bool {
    i, j := 0, 0
    sRunes, tRunes := []rune(s), []rune(t)
    
    for i < len(sRunes) && j < len(tRunes) {
        if sRunes[i] == tRunes[j] {
            i++
        }
        j++
    }
    return i == len(sRunes)
}
```
