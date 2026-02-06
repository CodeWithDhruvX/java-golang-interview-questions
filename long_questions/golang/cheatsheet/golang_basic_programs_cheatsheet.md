# Golang Strings & Arrays Cheatsheet

A quick reference for interview-frequent basic programs involving Strings, Arrays, and Slices.

---

## 🟢 Generic Helpers

### Max & Min (Integers)
Go 1.21+ has `max()` and `min()`, but for interviews you might need to write them or use them in logic.
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

## 🟡 Strings

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

### 4. Anagram Check
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

### 5. First Unique Character
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

---

### 6. Longest Substring Without Repeating Characters
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

### 7. Group Anagrams
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

### 8. Longest Common Prefix
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

### 9. Run Length Encoding
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

### 10. Is Subsequence
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

## 🔵 Arrays & Slices

### 1. Find Max Element in Slice
```go
func FindMax(nums []int) int {
    if len(nums) == 0 {
        return 0 // or error
    }
    maxVal := nums[0]
    for _, v := range nums {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```

### 2. Reverse a Slice
In-place reversal.
```go
func ReverseSlice(nums []int) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
```

### 3. Remove Duplicates
Returns a new slice with unique elements.
```go
func RemoveDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    
    for _, num := range nums {
        if !seen[num] {
            seen[num] = true
            result = append(result, num)
        }
    }
    return result
}
```

### 4. Rotate Array (Right by K steps)
Using generic reverse helper. Step 1: Reverse whole. Step 2: Reverse first k. Step 3: Reverse rest.
```go
func Rotate(nums []int, k int) {
    n := len(nums)
    k = k % n
    
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```

### 5. Two Sum (Target Sum)
Find indices of two numbers that add up to target. O(n) using map.
```go
func TwoSum(nums []int, target int) []int {
    seen := make(map[int]int) // val -> index
    
    for i, num := range nums {
        complement := target - num
        if idx, ok := seen[complement]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}
```

### 6. Merge Two Sorted Arrays
Merge `nums2` into `nums1` (assuming nums1 has space).
```go
func MergeSorted(nums1 []int, m int, nums2 []int, n int) {
    // Start from end
    i, j, k := m-1, n-1, m+n-1
    
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
    
    // Fill remaining from nums2 (nums1 remainder is already in place)
    for j >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }
}
```

### 7. Intersection of Two Arrays
Using map for O(n+m) time.
```go
func Intersection(nums1, nums2 []int) []int {
    set := make(map[int]bool)
    for _, n := range nums1 {
        set[n] = true
    }
    
    var result []int
    for _, n := range nums2 {
        if set[n] {
            result = append(result, n)
            delete(set, n) // ensure unique in result
        }
    }
    return result
}
```
