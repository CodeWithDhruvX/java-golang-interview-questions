package main

import (
	"fmt"
)

// 214. Shortest Palindrome - Z Algorithm
// Time: O(N), Space: O(N)
func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Create pattern: s + '#' + reverse(s)
	pattern := s + "#" + reverseString(s)
	
	// Compute Z-array
	zArray := computeZArray(pattern)
	
	// Find longest palindrome prefix
	maxPrefix := 0
	for i := len(s) + 1; i < len(zArray); i++ {
		if zArray[i] > maxPrefix {
			maxPrefix = zArray[i]
		}
	}
	
	// Add reverse of suffix to front
	suffix := s[maxPrefix:]
	return reverseString(suffix) + s
}

func computeZArray(s string) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	
	l, r := 0, 0
	
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	
	return z
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Z Algorithm for pattern matching
func zAlgorithmPatternSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	
	// Create pattern: pattern + '$' + text
	combined := pattern + "$" + text
	z := computeZArray(combined)
	
	// Find pattern occurrences
	occurrences := []int{}
	patternLen := len(pattern)
	
	for i := patternLen + 1; i < len(z); i++ {
		if z[i] == patternLen {
			occurrences = append(occurrences, i-(patternLen+1))
		}
	}
	
	return occurrences
}

// Z Algorithm for longest common prefix
func longestCommonPrefix(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	
	// Create pattern: s1 + '#' + s2
	combined := s1 + "#" + s2
	z := computeZArray(combined)
	
	// Find longest prefix
	maxPrefix := 0
	for i := len(s1) + 1; i < len(z); i++ {
		if z[i] > maxPrefix {
			maxPrefix = z[i]
		}
	}
	
	return maxPrefix
}

// Z Algorithm for string compression
func zAlgorithmCompression(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	z := computeZArray(s)
	
	// Find smallest pattern that can generate the string
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i && len(s)%i == 0 {
			pattern := s[:i]
			repeated := ""
			for j := 0; j < len(s)/i; j++ {
				repeated += pattern
			}
			if repeated == s {
				return fmt.Sprintf("%s{%d}", pattern, len(s)/i)
			}
		}
	}
	
	return s
}

// Z Algorithm for string rotation
func zAlgorithmRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	// Create pattern: s2 + '$' + s1 + s1
	combined := s2 + "$" + s1 + s1
	z := computeZArray(combined)
	
	// Check if s1 is rotation of s2
	for i := len(s2) + 1; i < len(z); i++ {
		if z[i] == len(s1) {
			return true
		}
	}
	
	return false
}

// Z Algorithm for palindrome detection
func zAlgorithmPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	
	// Create pattern: s + '#' + reverse(s)
	pattern := s + "#" + reverseString(s)
	z := computeZArray(pattern)
	
	// Check if entire string is palindrome
	for i := len(s) + 1; i < len(z); i++ {
		if z[i] == len(s) {
			return true
		}
	}
	
	return false
}

// Z Algorithm for substring search with wildcards
func zAlgorithmWildcardSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	
	// Replace wildcards with unique characters
	wildcardPattern := make([]byte, len(pattern))
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '?' {
			wildcardPattern[i] = 0 // Use 0 as wildcard
		} else {
			wildcardPattern[i] = pattern[i]
		}
	}
	
	// Create combined pattern
	combined := string(wildcardPattern) + "$" + text
	z := computeZArray(combined)
	
	// Find matches
	occurrences := []int{}
	patternLen := len(pattern)
	
	for i := patternLen + 1; i < len(z); i++ {
		if z[i] >= patternLen {
			occurrences = append(occurrences, i-(patternLen+1))
		}
	}
	
	return occurrences
}

// Z Algorithm for multiple pattern search
func zAlgorithmMultiplePatterns(text string, patterns []string) map[string][]int {
	results := make(map[string][]int)
	
	for _, pattern := range patterns {
		occurrences := zAlgorithmPatternSearch(text, pattern)
		results[pattern] = occurrences
	}
	
	return results
}

// Z Algorithm for string similarity
func zAlgorithmSimilarity(s1, s2 string) float64 {
	if len(s1) == 0 && len(s2) == 0 {
		return 1.0
	}
	
	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}
	
	// Find longest common prefix
	lcp := longestCommonPrefix(s1, s2)
	
	// Calculate similarity as ratio of LCP to shorter string
	shorter := len(s1)
	if len(s2) < shorter {
		shorter = len(s2)
	}
	
	return float64(lcp) / float64(shorter)
}

// Z Algorithm for string period
func zAlgorithmPeriod(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	
	z := computeZArray(s)
	
	// Find smallest period
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i && len(s)%i == 0 {
			return i
		}
	}
	
	return len(s)
}

// Z Algorithm for border detection
func zAlgorithmBorders(s string) []int {
	if len(s) <= 1 {
		return []int{}
	}
	
	z := computeZArray(s)
	borders := []int{}
	
	// Find all borders (proper prefixes that are also suffixes)
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i {
			borders = append(borders, i)
		}
	}
	
	return borders
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Z Algorithm for String Processing
- **Z-Array Computation**: Efficient substring prefix matching
- **Pattern Matching**: Find all occurrences of pattern in text
- **Palindrome Construction**: Build shortest palindrome by adding characters
- **String Analysis**: Period detection, border finding, similarity calculation

## 2. PROBLEM CHARACTERISTICS
- **String Manipulation**: Complex string operations and analysis
- **Prefix Matching**: Find longest prefix matches at each position
- **Pattern Recognition**: Detect repeating patterns and structures
- **Efficient Search**: O(N) time complexity for string operations

## 3. SIMILAR PROBLEMS
- Shortest Palindrome (LeetCode 214) - Same problem
- Find the Index of the First Occurrence in a String (LeetCode 28) - Pattern matching
- Repeated Substring Pattern (LeetCode 459) - Pattern detection
- Longest Common Prefix - Prefix matching

## 4. KEY OBSERVATIONS
- **Z-Array Power**: Z[i] gives length of longest prefix starting at position i
- **Pattern Concatenation**: Pattern + '#' + Text enables efficient matching
- **Palindrome Construction**: Use reverse string to find longest palindromic prefix
- **Linear Time**: Z algorithm computes all prefix matches in O(N) time

## 5. VARIATIONS & EXTENSIONS
- **Shortest Palindrome**: Add minimal characters to make palindrome
- **Pattern Search**: Find all pattern occurrences
- **String Compression**: Detect repeating patterns
- **Rotation Detection**: Check if strings are rotations of each other

## 6. INTERVIEW INSIGHTS
- Always clarify: "String length constraints? Character set? Multiple queries?"
- Edge cases: empty string, single character, no palindrome needed
- Time complexity: O(N) for Z array computation, O(N) for most operations
- Space complexity: O(N) for Z array storage
- Key insight: Z array reveals all prefix relationships in linear time

## 7. COMMON MISTAKES
- Wrong Z array computation with window boundaries
- Incorrect pattern concatenation format
- Missing edge cases for empty/single character strings
- Wrong palindrome construction logic
- Not handling Unicode characters properly

## 8. OPTIMIZATION STRATEGIES
- **Z Algorithm**: O(N) time, O(N) space - optimal for string matching
- **KMP Alternative**: O(N) time, O(N) space - similar complexity
- **Manacher's Algorithm**: O(N) time, O(N) space - specialized for palindromes
- **Rolling Hash**: O(N) time, O(1) space - probabilistic approach

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like analyzing a text for self-similarity:**
- The Z-array tells you how many characters match the beginning at each position
- For palindrome construction, you compare the string with its reverse
- The longest matching prefix tells you what's already a palindrome
- You only need to add the reverse of the non-palindromic suffix
- Like finding the longest mirror image in a text and completing it

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String to convert into shortest palindrome
2. **Goal**: Add minimum characters to front to make palindrome
3. **Rules**: Can only add characters to the front
4. **Output**: Shortest possible palindrome

#### Phase 2: Key Insight Recognition
- **"Z-array natural"** → Can find longest prefix matches efficiently
- **"String reversal useful"** → Compare with reverse to find palindrome prefix
- **"Pattern concatenation"** → String + '#' + reverse enables Z-array analysis
- **"Minimal addition"** → Only need to add reverse of non-palindromic part

#### Phase 3: Strategy Development
```
Human thought process:
"I need shortest palindrome by adding characters to front.
Brute force: try all possible additions O(N²).

Z Algorithm Approach:
1. Create pattern: s + '#' + reverse(s)
2. Compute Z-array for this pattern
3. Find longest prefix that matches suffix (palindrome part)
4. Add reverse of remaining suffix to front
5. Result is shortest palindrome

This gives O(N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty (already palindrome)
- **Single character**: Return single character (already palindrome)
- **Already palindrome**: Return original string
- **No palindrome prefix**: Add entire reverse string

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: s = "aacecaaa"

Human thinking:
"Z Algorithm Approach:
Step 1: Create pattern
s = "aacecaaa"
reverse(s) = "aaacecaa"
pattern = "aacecaaa#aaacecaa"

Step 2: Compute Z-array
Z values after '#' position show prefix matches:
- At position 9: Z=7 (matches "aacecaa")
- This means "aacecaa" is both prefix and suffix

Step 3: Find longest palindrome prefix
Longest match = 7 characters ("aacecaa")
Remaining suffix = "a"

Step 4: Add reverse of suffix
reverse("a") = "a"
Result = "a" + "aacecaaa" = "aaacecaaa"

Result: "aaacecaaa" ✓"
```

#### Phase 6: Intuition Validation
- **Why Z-array works**: Reveals all prefix-suffix relationships efficiently
- **Why string reversal**: Enables palindrome detection through comparison
- **Why minimal addition**: Only add what's needed to complete the palindrome
- **Why O(N)**: Single pass Z-array computation

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just reverse and add?"** → Might not be minimal
2. **"Should I use KMP?"** → Yes, Z-array is similar but more direct here
3. **"What about Manacher's algorithm?"** → More complex, overkill for this problem
4. **"Can I use rolling hash?"** → Yes, but Z-array is deterministic
5. **"Why concatenate with '#'?"** → Prevents false matches across boundary

### Real-World Analogy
**Like completing a mirror image:**
- You have a text and want to make it symmetrical
- You compare it with its mirror image to see what's already symmetrical
- You only need to add the missing mirror parts
- The Z-array tells you exactly how much matches the mirror
- Like completing a half-drawn butterfly to make it perfectly symmetrical

### Human-Readable Pseudocode
```
function shortestPalindrome(s):
    if length(s) <= 1:
        return s
    
    # Create pattern for Z-array
    pattern = s + "#" + reverse(s)
    
    # Compute Z-array
    z = computeZArray(pattern)
    
    # Find longest palindrome prefix
    maxPrefix = 0
    for i from length(s) + 1 to length(z) - 1:
        maxPrefix = max(maxPrefix, z[i])
    
    # Add reverse of suffix to front
    suffix = s[maxPrefix:]
    return reverse(suffix) + s

function computeZArray(s):
    z = array of size length(s)
    z[0] = length(s)
    l, r = 0, 0
    
    for i from 1 to length(s) - 1:
        if i <= r:
            z[i] = min(r - i + 1, z[i - l])
        
        while i + z[i] < length(s) and s[z[i]] == s[i + z[i]]:
            z[i]++
        
        if i + z[i] - 1 > r:
            l, r = i, i + z[i] - 1
    
    return z
```

### Execution Visualization

### Example: s = "aacecaaa"
```
Z Algorithm Process:

Step 1: Create pattern
s = "aacecaaa"
reverse(s) = "aaacecaa"
pattern = "aacecaaa#aaacecaa"

Step 2: Compute Z-array (partial)
Index: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
Char:  a a c e c a a a # a  a  a  c  e  c  a  a
Z:    [16,0,0,0,0,0,0,0,0,7,0,0,0,0,0,0]

Step 3: Find palindrome prefix
At position 9 (after '#'): Z[9] = 7
This means 7 characters match: "aacecaa"

Step 4: Construct palindrome
Original: "aacecaaa"
Palindrome prefix: "aacecaa" (7 chars)
Remaining suffix: "a" (1 char)
Add reverse(suffix): "a"

Result: "a" + "aacecaaa" = "aaacecaaa" ✓
```

### Key Visualization Points:
- **Pattern Construction**: String + '#' + reverse enables comparison
- **Z-Array Computation**: Linear time prefix matching
- **Palindrome Detection**: Longest prefix-suffix match
- **Minimal Addition**: Only add what's necessary

### Z-Array Visualization:
```
String: "aacecaaa#aaacecaa"
Index:  012345678901234567
Z:     [16,0,0,0,0,0,0,0,0,7,0,0,0,0,0,0]

Z[9] = 7 means:
- Starting at position 9 ("a")
- 7 characters match the prefix "aacecaa"
- "aacecaa" is both prefix and suffix
- This is the longest palindromic prefix
```

### Time Complexity Breakdown:
- **Z Algorithm**: O(N) time, O(N) space - optimal approach
- **KMP Alternative**: O(N) time, O(N) space - similar complexity
- **Manacher's Algorithm**: O(N) time, O(N) space - specialized for palindromes
- **Rolling Hash**: O(N) time, O(1) space - probabilistic approach

### Alternative Approaches:

#### 1. KMP-based Solution (O(N) time, O(N) space)
```go
func shortestPalindromeKMP(s string) string {
    // Build LPS array for s + '#' + reverse(s)
    // Find longest prefix that's also suffix
    // Add reverse of remaining part
    // Similar to Z-array but uses LPS
}
```
- **Pros**: Same complexity, well-known algorithm
- **Cons**: More complex than Z-array for this problem

#### 2. Manacher's Algorithm (O(N) time, O(N) space)
```go
func shortestPalindromeManacher(s string) string {
    // Find longest palindromic prefix
    // Use center expansion technique
    // Optimized for palindrome problems
}
```
- **Pros**: Specialized for palindromes
- **Cons**: More complex implementation

#### 3. Rolling Hash (O(N) time, O(1) space)
```go
func shortestPalindromeHash(s string) string {
    // Use rolling hash to find palindrome prefix
    // Compare forward and backward hashes
    // Probabilistic but fast
}
```
- **Pros**: Constant space, fast
- **Cons**: Hash collision possibility

### Extensions for Interviews:
- **All Palindromes**: Find all palindromic substrings
- **Palindrome Pairs**: Find pairs that form palindromes
- **Minimum Insertions**: Find minimum insertions anywhere
- **Palindrome Partitioning**: Partition into palindromes
- **Real-world Applications**: DNA sequence analysis, text processing
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Z Algorithm ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"aacecaaa", "Standard case"},
		{"abcd", "No palindrome prefix"},
		{"a", "Single character"},
		{"", "Empty string"},
		{"racecar", "Palindrome"},
		{"abacdfgdcaba", "Multiple palindrome prefixes"},
		{"aaaaa", "All same characters"},
		{"abcba", "Odd length palindrome"},
		{"abccba", "Even length palindrome"},
		{"xyz", "No palindrome"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (\"%s\")\n", i+1, tc.description, tc.s)
		
		result := shortestPalindrome(tc.s)
		fmt.Printf("  Shortest palindrome: %s\n", result)
		
		fmt.Println()
	}
	
	// Test pattern matching
	fmt.Println("=== Pattern Matching Test ===")
	text := "ababcabcababc"
	pattern := "abc"
	occurrences := zAlgorithmPatternSearch(text, pattern)
	fmt.Printf("Text: %s, Pattern: %s\n", text, pattern)
	fmt.Printf("Occurrences: %v\n", occurrences)
	
	// Test longest common prefix
	fmt.Println("\n=== Longest Common Prefix Test ===")
	lcpTests := [][]string{
		{"abcdef", "abcxyz"},
		{"hello", "world"},
		{"same", "same"},
		{"prefix", "pre"},
	}
	
	for _, test := range lcpTests {
		lcp := longestCommonPrefix(test[0], test[1])
		fmt.Printf("LCP of \"%s\" and \"%s\": %d\n", test[0], test[1], lcp)
	}
	
	// Test string compression
	fmt.Println("\n=== String Compression Test ===")
	compressionTests := []string{
		"abababab",
		"abcabcabc",
		"aaaaaa",
		"abcdef",
		"xyzxyzxyz",
	}
	
	for _, test := range compressionTests {
		compressed := zAlgorithmCompression(test)
		fmt.Printf("Original: %s, Compressed: %s\n", test, compressed)
	}
	
	// Test string rotation
	fmt.Println("\n=== String Rotation Test ===")
	rotationTests := [][2]string{
		{"abcde", "cdeab"},
		{"hello", "llohe"},
		{"test", "sett"},
		{"abc", "def"},
	}
	
	for _, test := range rotationTests {
		isRotation := zAlgorithmRotation(test[0], test[1])
		fmt.Printf("Is \"%s\" rotation of \"%s\": %t\n", test[1], test[0], isRotation)
	}
	
	// Test palindrome detection
	fmt.Println("\n=== Palindrome Detection Test ===")
	palindromeTests := []string{
		"racecar",
		"hello",
		"madam",
		"test",
		"level",
	}
	
	for _, test := range palindromeTests {
		isPalindrome := zAlgorithmPalindrome(test)
		fmt.Printf("Is \"%s\" palindrome: %t\n", test, isPalindrome)
	}
	
	// Test wildcard search
	fmt.Println("\n=== Wildcard Search Test ===")
	wildcardTests := []struct {
		text    string
		pattern string
	}{
		{"abcde", "a?cde"},
		{"hello", "h?llo"},
		{"test", "t??t"},
		{"world", "?????"},
	}
	
	for _, test := range wildcardTests {
		wildcardOccurrences := zAlgorithmWildcardSearch(test.text, test.pattern)
		fmt.Printf("Text: %s, Pattern: %s, Occurrences: %v\n", 
			test.text, test.pattern, wildcardOccurrences)
	}
	
	// Test multiple patterns
	fmt.Println("\n=== Multiple Patterns Test ===")
	text := "ababcabcababc"
	patterns := []string{"abc", "ab", "bc"}
	multipleResults := zAlgorithmMultiplePatterns(text, patterns)
	
	fmt.Printf("Text: %s\n", text)
	for pattern, occurrences := range multipleResults {
		fmt.Printf("Pattern \"%s\": %v\n", pattern, occurrences)
	}
	
	// Test string similarity
	fmt.Println("\n=== String Similarity Test ===")
	similarityTests := [][2]string{
		{"hello", "help"},
		{"test", "testing"},
		{"abc", "abc"},
		{"xyz", "abc"},
	}
	
	for _, test := range similarityTests {
		similarity := zAlgorithmSimilarity(test[0], test[1])
		fmt.Printf("Similarity between \"%s\" and \"%s\": %.2f\n", 
			test[0], test[1], similarity)
	}
	
	// Test string period
	fmt.Println("\n=== String Period Test ===")
	periodTests := []string{
		"abababab",
		"abcabcabc",
		"aaaaaa",
		"abcdef",
		"xyzxyzxyz",
	}
	
	for _, test := range periodTests {
		period := zAlgorithmPeriod(test)
		fmt.Printf("Period of \"%s\": %d\n", test, period)
	}
	
	// Test border detection
	fmt.Println("\n=== Border Detection Test ===")
	borderTests := []string{
		"ababcab",
		"aaaaa",
		"abcab",
		"test",
		"ababa",
	}
	
	for _, test := range borderTests {
		borders := zAlgorithmBorders(test)
		fmt.Printf("Borders of \"%s\": %v\n", test, borders)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Large string
	largeString := ""
	for i := 0; i < 10000; i++ {
		largeString += string('a' + (i % 26))
	}
	
	fmt.Printf("Large string test with %d characters\n", len(largeString))
	
	start := time.Now()
	result := shortestPalindrome(largeString)
	duration := time.Since(start)
	
	fmt.Printf("Large string result length: %d\n", len(result))
	fmt.Printf("Time taken: %v\n", duration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long palindrome
	longPalindrome := ""
	for i := 0; i < 1000; i++ {
		longPalindrome += string('a' + (i % 26))
	}
	longPalindrome += reverseString(longPalindrome)
	
	fmt.Printf("Long palindrome test: length=%d\n", len(longPalindrome))
	result = shortestPalindrome(longPalindrome)
	fmt.Printf("Result length: %d\n", len(result))
	
	// Empty string
	fmt.Printf("Empty string: %s\n", shortestPalindrome(""))
	
	// Single character
	fmt.Printf("Single character: %s\n", shortestPalindrome("a"))
}
