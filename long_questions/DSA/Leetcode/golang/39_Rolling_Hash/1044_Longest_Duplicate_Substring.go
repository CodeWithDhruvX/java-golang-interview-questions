package main

import (
	"fmt"
	"math"
)

// 1044. Longest Duplicate Substring - Rolling Hash with Binary Search
// Time: O(N log N), Space: O(N)
func longestDupSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	// Binary search for the length of longest duplicate
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicate(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

// Find duplicate substring of given length using rolling hash
func findDuplicate(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Rolling hash with base and mod
	base := 256
	mod := 1000000007
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < length; i++ {
		hash = (hash*base + int(s[i])) % mod
	}
	
	// Store seen hashes
	seen := make(map[int]int)
	seen[hash] = 0
	
	// Precompute base^length mod
	pow := 1
	for i := 0; i < length; i++ {
		pow = (pow * base) % mod
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash = (hash - int(s[i-1])*pow%mod + mod) % mod
		
		// Add new character
		hash = (hash*base + int(s[i+length-1])) % mod
		
		// Check for collision
		if prevIndex, exists := seen[hash]; exists {
			// Verify actual substring (handle hash collisions)
			if s[prevIndex:prevIndex+length] == s[i:i+length] {
				return s[i : i+length]
			}
		} else {
			seen[hash] = i
		}
	}
	
	return ""
}

// Optimized version using double hashing
func longestDupSubstringDoubleHash(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicateDoubleHash(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findDuplicateDoubleHash(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Two different hash functions
	base1, mod1 := 256, 1000000007
	base2, mod2 := 257, 1000000009
	
	// Calculate initial hashes
	hash1, hash2 := 0, 0
	for i := 0; i < length; i++ {
		hash1 = (hash1*base1 + int(s[i])) % mod1
		hash2 = (hash2*base2 + int(s[i])) % mod2
	}
	
	// Store seen hash pairs
	seen := make(map[[2]int]int)
	seen[[2]int{hash1, hash2}] = 0
	
	// Precompute powers
	pow1, pow2 := 1, 1
	for i := 0; i < length; i++ {
		pow1 = (pow1 * base1) % mod1
		pow2 = (pow2 * base2) % mod2
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash1 = (hash1 - int(s[i-1])*pow1%mod1 + mod1) % mod1
		hash2 = (hash2 - int(s[i-1])*pow2%mod2 + mod2) % mod2
		
		// Add new character
		hash1 = (hash1*base1 + int(s[i+length-1])) % mod1
		hash2 = (hash2*base2 + int(s[i+length-1])) % mod2
		
		hashPair := [2]int{hash1, hash2}
		
		if prevIndex, exists := seen[hashPair]; exists {
			return s[i : i+length]
		} else {
			seen[hashPair] = i
		}
	}
	
	return ""
}

// Version using Rabin-Karp with bit manipulation
func longestDupSubstringRabinKarp(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicateRabinKarp(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findDuplicateRabinKarp(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Use base 26 for lowercase letters
	base := 26
	mod := 2 << 61 // Large prime-like number
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < length; i++ {
		hash = (hash*base + int(s[i]-'a')) % mod
	}
	
	// Store seen hashes
	seen := make(map[int]int)
	seen[hash] = 0
	
	// Precompute base^length mod
	pow := 1
	for i := 0; i < length; i++ {
		pow = (pow * base) % mod
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash = (hash - int(s[i-1]-'a')*pow%mod + mod) % mod
		
		// Add new character
		hash = (hash*base + int(s[i+length-1]-'a')) % mod
		
		if prevIndex, exists := seen[hash]; exists {
			if s[prevIndex:prevIndex+length] == s[i:i+length] {
				return s[i : i+length]
			}
		} else {
			seen[hash] = i
		}
	}
	
	return ""
}

// Brute force approach for comparison (O(N^2))
func longestDupSubstringBruteForce(s string) string {
	maxLength := 0
	result := ""
	
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// Find common prefix length
			k := 0
			for i+k < len(s) && j+k < len(s) && s[i+k] == s[j+k] {
				k++
			}
			
			if k > maxLength {
				maxLength = k
				result = s[i : i+k]
			}
		}
	}
	
	return result
}

// Suffix array approach (more advanced)
func longestDupSubstringSuffixArray(s string) string {
	// Build suffix array (simplified version)
	suffixes := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		suffixes[i] = s[i:]
	}
	
	// Sort suffixes (bubble sort for demonstration)
	for i := 0; i < len(suffixes)-1; i++ {
		for j := 0; j < len(suffixes)-i-1; j++ {
			if suffixes[j] > suffixes[j+1] {
				suffixes[j], suffixes[j+1] = suffixes[j+1], suffixes[j]
			}
		}
	}
	
	// Find longest common prefix between adjacent suffixes
	maxLength := 0
	result := ""
	
	for i := 0; i < len(suffixes)-1; i++ {
		lcp := longestCommonPrefix(suffixes[i], suffixes[i+1])
		if len(lcp) > maxLength {
			maxLength = len(lcp)
			result = lcp
		}
	}
	
	return result
}

func longestCommonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	
	i := 0
	for i < minLen && a[i] == b[i] {
		i++
	}
	
	return a[:i]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Rolling Hash with Binary Search
- **Rabin-Karp Algorithm**: Efficient substring hashing
- **Binary Search**: Search for maximum length of duplicate substring
- **Hash Collision Handling**: Verify actual substrings when hash matches
- **Double Hashing**: Use two hash functions to reduce collisions

## 2. PROBLEM CHARACTERISTICS
- **String Analysis**: Find longest repeated substring
- **Substring Comparison**: Need efficient substring equality checking
- **Length Maximization**: Find maximum possible length
- **Collision Handling**: Hash collisions must be resolved

## 3. SIMILAR PROBLEMS
- Repeated DNA Sequences (LeetCode 187) - Rolling hash for patterns
- Longest Common Substring - Similar substring finding
- Substring with Concatenation of All Words - Hash-based pattern matching
- Detect Squares (LeetCode 593) - String pattern detection

## 4. KEY OBSERVATIONS
- **Binary Search Natural**: Can test if duplicate of length L exists
- **Rolling Hash Efficient**: O(1) hash updates for sliding window
- **Collision Verification**: Must verify actual substring equality
- **Hash Function Choice**: Base and modulus affect collision probability

## 5. VARIATIONS & EXTENSIONS
- **Double Hashing**: Two hash functions for collision reduction
- **Rabin-Karp**: Classic rolling hash implementation
- **Suffix Array**: Alternative approach with different complexity
- **Bit Manipulation**: Use bit operations for faster hashing

## 6. INTERVIEW INSIGHTS
- Always clarify: "String length constraints? Character set? Collision tolerance?"
- Edge cases: empty string, no duplicates, all same characters
- Time complexity: O(N log N) with binary search and rolling hash
- Space complexity: O(N) for hash storage
- Key insight: binary search + rolling hash = efficient substring comparison

## 7. COMMON MISTAKES
- Not handling hash collisions properly
- Wrong rolling hash update formula
- Incorrect binary search boundaries
- Missing modulo operations causing overflow
- Not verifying actual substring equality

## 8. OPTIMIZATION STRATEGIES
- **Single Hash**: O(N log N) time, O(N) space - basic approach
- **Double Hash**: O(N log N) time, O(N) space - collision reduction
- **Rabin-Karp**: O(N log N) time, O(N) space - optimized base
- **Suffix Array**: O(N log² N) time, O(N) space - alternative approach

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest repeated phrase in a book:**
- You want to find the longest phrase that appears multiple times
- Instead of comparing all pairs (O(N²)), you can test lengths efficiently
- For each length, you slide a window and hash each substring
- If two substrings have the same hash, they might be equal
- Like finding plagiarism by comparing sentence fragments

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String to analyze
2. **Goal**: Find longest substring that appears at least twice
3. **Rules**: Substrings can overlap, must be identical
4. **Output**: Longest duplicate substring (empty if none)

#### Phase 2: Key Insight Recognition
- **"Binary search natural"** → Can test if duplicate of length L exists
- **"Rolling hash efficient"** → O(1) hash updates for sliding window
- **"Collision handling required"** → Hash matches need verification
- **"Length maximization"** → Binary search finds maximum possible length

#### Phase 3: Strategy Development
```
Human thought process:
"I need longest duplicate substring.
Brute force: compare all pairs O(N²) substrings.

Rolling Hash + Binary Search Approach:
1. Binary search on length L (1 to N)
2. For each L, check if duplicate exists:
   - Compute rolling hash for all substrings of length L
   - Store hashes in map
   - If hash collision, verify actual substring equality
3. If duplicate exists, search longer lengths
4. If no duplicate, search shorter lengths

This gives O(N log N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty result
- **No duplicates**: Return empty result
- **All same characters**: Return string length-1
- **Single character**: Return empty result

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: s = "banana"

Human thinking:
"Rolling Hash + Binary Search Approach:
Binary search on length: left=1, right=6

Test length 3:
- Substrings: "ban", "ana", "nan", "ana"
- Hashes: h1, h2, h3, h2 (collision on "ana")
- Verify: s[1:4] == s[3:6] = "ana" ✓
- Duplicate exists, try longer

Test length 4:
- Substrings: "bana", "anan", "nana"
- Hashes: all different
- No duplicates, try shorter

Test length 2:
- Substrings: "ba", "an", "na", "an", "na"
- Hashes: h1, h2, h3, h2, h3 (collisions)
- Verify: "an" appears twice ✓
- Duplicate exists, try longer

Result: "ana" (length 3) ✓"
```

#### Phase 6: Intuition Validation
- **Why binary search**: Monotonic property - if duplicate of length L exists, duplicates of smaller lengths exist
- **Why rolling hash**: O(1) updates vs O(L) for each substring
- **Why collision verification**: Hash collisions can give false positives
- **Why O(N log N)**: Binary search (log N) × rolling hash scan (N)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compare all pairs?"** → O(N²) vs O(N log N) with rolling hash
2. **"Should I use suffix array?"** → Yes, but more complex to implement
3. **"What about hash collisions?"** → Must verify actual substring equality
4. **"Can I use single hash?"** → Yes, but double hash reduces collisions
5. **"Why binary search?"** → Monotonic property enables efficient search

### Real-World Analogy
**Like detecting plagiarism in documents:**
- You want to find the longest copied passage
- Instead of comparing all sentence pairs, you hash fragments
- You search for the longest matching fragment length
- Hash collisions require exact text comparison
- Like finding copyright infringement in large documents

### Human-Readable Pseudocode
```
function longestDupSubstring(s):
    if s is empty:
        return ""
    
    left, right = 1, length(s)
    result = ""
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if duplicate = findDuplicate(s, mid):
            result = duplicate
            left = mid + 1
        else:
            right = mid - 1
    
    return result

function findDuplicate(s, length):
    if length == 0:
        return ""
    
    # Rolling hash setup
    base = 256
    mod = 1000000007
    hash = 0
    
    # Initial hash
    for i from 0 to length-1:
        hash = (hash * base + ord(s[i])) % mod
    
    seen = {hash: 0}
    
    # Precompute base^length
    power = 1
    for i from 0 to length-1:
        power = (power * base) % mod
    
    # Rolling hash
    for i from 1 to length(s)-length:
        # Remove leftmost character
        hash = (hash - ord(s[i-1]) * power) % mod
        # Add new character
        hash = (hash * base + ord(s[i+length-1])) % mod
        
        if hash in seen:
            prevIndex = seen[hash]
            if s[prevIndex:prevIndex+length] == s[i:i+length]:
                return s[i:i+length]
        else:
            seen[hash] = i
    
    return ""
```

### Execution Visualization

### Example: s = "banana", testing length 3
```
Rolling Hash Process:

Initial hash for "ban":
hash = ((0*256 + 'b')*256 + 'a')*256 + 'n' % mod
hash = (((98*256 + 97)*256 + 110) % mod)

Roll to "ana":
hash = (hash - 'b'*power) % mod  # Remove 'b'
hash = (hash*256 + 'a') % mod      # Add 'a'

Hash values:
"ban" → h1, "ana" → h2, "nan" → h3, "ana" → h2

Collision on h2:
- Check s[1:4] == s[3:6]
- "ana" == "ana" ✓
- Found duplicate!

Result: "ana" ✓
```

### Key Visualization Points:
- **Binary Search**: Efficiently find maximum length
- **Rolling Hash**: O(1) updates for sliding window
- **Collision Handling**: Verify actual substring equality
- **Hash Storage**: Map hash to starting index

### Rolling Hash Visualization:
```
String:  b  a  n  a  n  a
Index:  0  1  2  3  4  5

Window length 3:
[0:3] "ban" → hash h1
[1:4] "ana" → hash h2  
[2:5] "nan" → hash h3
[3:6] "ana" → hash h2 (collision!)

Collision verification:
s[1:4] = "ana"
s[3:6] = "ana"
Equal! ✓
```

### Time Complexity Breakdown:
- **Single Hash**: O(N log N) time, O(N) space - basic approach
- **Double Hash**: O(N log N) time, O(N) space - collision reduction
- **Rabin-Karp**: O(N log N) time, O(N) space - optimized base
- **Suffix Array**: O(N log² N) time, O(N) space - alternative approach

### Alternative Approaches:

#### 1. Suffix Array (O(N log² N) time, O(N) space)
```go
func longestDupSubstringSuffixArray(s string) string {
    // Build suffix array
    // Sort all suffixes
    // Find longest common prefix between adjacent suffixes
    // More complex but different complexity
    // ... implementation details omitted
}
```
- **Pros**: Guaranteed correct, no hash collisions
- **Cons**: More complex implementation, different complexity

#### 2. Brute Force (O(N²) time, O(1) space)
```go
func longestDupSubstringBruteForce(s string) string {
    // Compare all pairs of positions
    // Find longest common prefix
    // Simple but slow
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Too slow for large strings

#### 3. Trie-Based (O(N²) time, O(N²) space)
```go
func longestDupSubstringTrie(s string) string {
    // Build trie of all suffixes
    // Track longest path with multiple children
    // High memory usage
    // ... implementation details omitted
}
```
- **Pros**: Intuitive suffix comparison
- **Cons**: High memory usage, still O(N²) time

### Extensions for Interviews:
- **Multiple Duplicates**: Find all longest duplicate substrings
- **Pattern Matching**: Extend to general pattern detection
- **Unicode Support**: Handle international characters
- **Streaming**: Process very large strings in chunks
- **Real-world Applications**: Plagiarism detection, DNA sequence analysis
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Longest Duplicate Substring ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"banana", "Standard case"},
		{"abcd", "No duplicates"},
		{"aaaaa", "All same characters"},
		{"abcabcabc", "Repeated pattern"},
		{"", "Empty string"},
		{"a", "Single character"},
		{"abababab", "Alternating pattern"},
		{"abcdeabcdf", "Partial match"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaa", "Long same characters"},
		{"abcdefghijklmnopqrstuvwxyza", "Long with single repeat"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  String: %s\n", tc.s)
		
		result1 := longestDupSubstring(tc.s)
		result2 := longestDupSubstringDoubleHash(tc.s)
		result3 := longestDupSubstringRabinKarp(tc.s)
		result4 := longestDupSubstringBruteForce(tc.s)
		result5 := longestDupSubstringSuffixArray(tc.s)
		
		fmt.Printf("  Standard rolling hash: %s\n", result1)
		fmt.Printf("  Double hash: %s\n", result2)
		fmt.Printf("  Rabin-Karp: %s\n", result3)
		fmt.Printf("  Brute force: %s\n", result4)
		fmt.Printf("  Suffix array: %s\n\n", result5)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += "abcdefghijklmnopqrstuvwxyz"
	}
	
	fmt.Printf("Testing with string length: %d\n", len(longString))
	
	// Test rolling hash approach
	result := longestDupSubstring(longString)
	fmt.Printf("Rolling hash result length: %d\n", len(result))
	
	// Test edge cases
	fmt.Println("\n=== Edge Case Testing ===")
	
	// Test with very long repeated pattern
	repeatPattern := ""
	for i := 0; i < 100; i++ {
		repeatPattern += "abc"
	}
	
	fmt.Printf("Long repeated pattern: %s\n", longestDupSubstring(repeatPattern))
	
	// Test with no duplicates
	noDup := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("No duplicates: %s\n", longestDupSubstring(noDup))
	
	// Test with overlapping duplicates
	overlap := "aaaaabaaaaa"
	fmt.Printf("Overlapping duplicates: %s\n", longestDupSubstring(overlap))
	
	// Test complexity analysis
	fmt.Println("\n=== Complexity Analysis ===")
	
	// Test different string sizes
	sizes := []int{10, 50, 100, 200}
	for _, size := range sizes {
		testString := ""
		for i := 0; i < size; i++ {
			testString += string(rune('a' + (i % 26)))
		}
		
		result = longestDupSubstring(testString)
		fmt.Printf("Size %d: result length %d\n", size, len(result))
	}
}
