package main

import (
	"fmt"
	"math"
)

// 28. Find the Index of the First Occurrence in a String - KMP Algorithm
// Time: O(N + M), Space: O(M) where N is haystack length, M is needle length
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	// Build LPS (Longest Prefix Suffix) array
	lps := buildLPS(needle)
	
	i, j := 0, 0 // i: haystack index, j: needle index
	
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			
			if j == len(needle) {
				return i - j // Found match
			}
		} else {
			if j != 0 {
				j = lps[j-1] // Use LPS to skip comparisons
			} else {
				i++
			}
		}
	}
	
	return -1 // Not found
}

// Build LPS array for KMP
func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // Length of previous longest prefix suffix
	
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	
	return lps
}

// KMP with detailed tracing
func strStrKMPDetailed(haystack string, needle string) (int, []string) {
	if needle == "" {
		return 0, []string{"Empty needle, returning 0"}
	}
	
	var trace []string
	trace = append(trace, fmt.Sprintf("Building LPS for: %s", needle))
	
	lps := buildLPSWithTrace(needle, &trace)
	trace = append(trace, fmt.Sprintf("LPS array: %v", lps))
	
	i, j := 0, 0
	trace = append(trace, fmt.Sprintf("Starting search: i=%d, j=%d", i, j))
	
	for i < len(haystack) {
		trace = append(trace, fmt.Sprintf("Comparing haystack[%d]='%c' with needle[%d]='%c'", 
			i, haystack[i], j, needle[j]))
		
		if haystack[i] == needle[j] {
			i++
			j++
			trace = append(trace, fmt.Sprintf("Match! i=%d, j=%d", i, j))
			
			if j == len(needle) {
				trace = append(trace, fmt.Sprintf("Found match at index %d", i-j))
				return i - j, trace
			}
		} else {
			if j != 0 {
				trace = append(trace, fmt.Sprintf("Mismatch, using LPS[%d-1]=%d", j, lps[j-1]))
				j = lps[j-1]
			} else {
				trace = append(trace, fmt.Sprintf("Mismatch, moving i forward"))
				i++
			}
		}
	}
	
	trace = append(trace, "No match found")
	return -1, trace
}

func buildLPSWithTrace(pattern string, trace *[]string) []int {
	lps := make([]int, len(pattern))
	length := 0
	
	i := 1
	for i < len(pattern) {
		*trace = append(*trace, fmt.Sprintf("Building LPS: i=%d, j=%d, pattern[i]='%c', pattern[length]='%c'", 
			i, length, pattern[i], pattern[length]))
		
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			*trace = append(*trace, fmt.Sprintf("Match! LPS[%d]=%d", i, length))
			i++
		} else {
			if length != 0 {
				*trace = append(*trace, fmt.Sprintf("Mismatch, using LPS[%d]=%d", length-1, lps[length-1]))
				length = lps[length-1]
			} else {
				lps[i] = 0
				*trace = append(*trace, fmt.Sprintf("No match, LPS[%d]=0", i))
				i++
			}
		}
	}
	
	return lps
}

// Alternative KMP implementation
func strStrKMPAlternative(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	// Build prefix table
	prefix := buildPrefixTable(needle)
	
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		
		if j == len(needle) {
			return i
		} else if j > 0 {
			// Skip ahead using prefix table
			i += j - prefix[j-1] - 1
		}
	}
	
	return -1
}

func buildPrefixTable(pattern string) []int {
	prefix := make([]int, len(pattern))
	
	for i := 1; i < len(pattern); i++ {
		j := prefix[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = prefix[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		prefix[i] = j
	}
	
	return prefix
}

// Rabin-Karp approach for comparison
func strStrRabinKarp(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	
	// Calculate hash of needle
	needleHash := 0
	haystackHash := 0
	base := 256
	mod := 1000000007
	
	for i := 0; i < len(needle); i++ {
		needleHash = (needleHash*base + int(needle[i])) % mod
		haystackHash = (haystackHash*base + int(haystack[i])) % mod
	}
	
	// Precompute base^(len(needle)-1) mod
	pow := 1
	for i := 0; i < len(needle)-1; i++ {
		pow = (pow * base) % mod
	}
	
	// Slide window
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystackHash == needleHash {
			// Verify to avoid hash collision
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
		
		if i < len(haystack)-len(needle) {
			// Remove leftmost character
			haystackHash = (haystackHash - int(haystack[i])*pow%mod + mod) % mod
			// Add new character
			haystackHash = (haystackHash*base + int(haystack[i+len(needle)])) % mod
		}
	}
	
	return -1
}

// Brute force approach
func strStrBruteForce(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	
	return -1
}

// Built-in string search for comparison
func strStrBuiltIn(haystack string, needle string) int {
	return findSubstring(haystack, needle)
}

func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Knuth-Morris-Pratt (KMP) String Matching
- **LPS Array**: Longest Prefix Suffix array for pattern preprocessing
- **Efficient Matching**: Skip unnecessary comparisons using LPS
- **Linear Time**: O(N + M) where N is text length, M is pattern length
- **No Backtracking**: Never re-examine characters in text

## 2. PROBLEM CHARACTERISTICS
- **String Search**: Find first occurrence of pattern in text
- **Pattern Matching**: Exact substring matching required
- **Efficiency**: Need better than O(N × M) brute force
- **Index Return**: Return starting index or -1 if not found

## 3. SIMILAR PROBLEMS
- Repeated Substring Pattern (LeetCode 459) - KMP for pattern detection
- Implement strStr() (LeetCode 28) - Same problem
- Find the Index of the First Occurrence - Same problem
- Rabin-Karp Algorithm - Alternative string matching approach

## 4. KEY OBSERVATIONS
- **LPS Natural**: Preprocess pattern to find internal structure
- **No Backtracking**: KMP never moves text pointer backward
- **Prefix Analysis**: LPS captures pattern's self-overlapping properties
- **Linear Guarantee**: Each character examined at most twice

## 5. VARIATIONS & EXTENSIONS
- **Standard KMP**: Classic LPS-based implementation
- **Alternative Prefix**: Different LPS construction approach
- **Rabin-Karp**: Hash-based string matching
- **Built-in Functions**: Language-provided string search

## 6. INTERVIEW INSIGHTS
- Always clarify: "Empty string handling? Case sensitivity? Unicode support?"
- Edge cases: empty needle, needle longer than haystack, no match
- Time complexity: O(N + M) with KMP, O(N × M) with brute force
- Space complexity: O(M) for LPS array
- Key insight: preprocess pattern to avoid redundant comparisons

## 7. COMMON MISTAKES
- Wrong LPS array construction
- Incorrect index handling during matching
- Missing empty string edge cases
- Off-by-one errors in LPS usage
- Not handling pattern longer than text

## 8. OPTIMIZATION STRATEGIES
- **Standard KMP**: O(N + M) time, O(M) space - optimal approach
- **Alternative KMP**: O(N + M) time, O(M) space - different LPS build
- **Rabin-Karp**: O(N + M) time, O(1) space - hash-based
- **Brute Force**: O(N × M) time, O(1) space - simple but slow

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a specific word in a book efficiently:**
- Instead of re-reading from the beginning after each mismatch
- You remember how much of the pattern you've already matched
- When you mismatch, you jump to the longest possible continuation
- Like a smart reader who doesn't start over after every typo

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Haystack (text) and needle (pattern) strings
2. **Goal**: Find first occurrence index of needle in haystack
3. **Rules**: Return -1 if not found, 0 if needle is empty
4. **Output**: Starting index of first match

#### Phase 2: Key Insight Recognition
- **"Brute force wasteful"** → Re-examining characters after mismatches
- **"Pattern structure useful"** → Pattern has internal repeating structure
- **"LPS captures overlaps"** → Longest prefix that's also suffix
- **"No backtracking needed"** → Can always move forward in text

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find pattern in text efficiently.
Brute force: O(N×M) - restart from scratch after each mismatch.

KMP Algorithm Approach:
1. Build LPS array for pattern:
   - LPS[i] = length of longest proper prefix which is also suffix
   - For pattern "ABABC": LPS = [0,0,1,2,0]
2. Search using LPS:
   - When mismatch, use LPS to skip redundant comparisons
   - Never move text pointer backward
   - Each character examined at most twice

This gives O(N + M) time!"
```

#### Phase 4: Edge Case Handling
- **Empty needle**: Return 0 (by definition)
- **Empty haystack**: Return -1 unless needle is also empty
- **Needle longer than haystack**: Return -1 immediately
- **No match found**: Return -1 after complete search

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: haystack = "ABABDABACDABABCABAB", needle = "ABABCABAB"

Human thinking:
"KMP Algorithm Approach:
Step 1: Build LPS for "ABABCABAB"
- LPS[0] = 0 (first character)
- LPS[1] = 0 ("A" != "B")
- LPS[2] = 1 ("AB" prefix = "AB" suffix)
- LPS[3] = 2 ("ABA" prefix = "ABA" suffix)
- LPS[4] = 0 ("ABAB" doesn't end with "ABAB")
- Continue...
Final LPS: [0,0,1,2,0,1,2,3,4]

Step 2: Search using LPS
Compare haystack and needle character by character:
- Match: A=A, B=B, A=A, B=B, C!=D → mismatch at position 4
- Use LPS[3]=2: jump needle to position 2, continue from haystack[4]
- Continue this process until full match found at index 10

Result: 10 ✓"
```

#### Phase 6: Intuition Validation
- **Why LPS works**: Captures pattern's internal structure for smart jumps
- **Why no backtracking**: LPS ensures we never miss potential matches
- **Why O(N + M)**: Each character examined at most twice
- **Why preprocessing**: One-time cost to enable efficient matching

### Common Human Pitfalls & How to Avoid Them
1. **"Why not brute force?"** → O(N×M) vs O(N+M) with KMP
2. **"Should I use Rabin-Karp?"** → Yes, but KMP has no hash collisions
3. **"What about LPS construction?"** → Critical for KMP efficiency
4. **"Can I use built-in functions?"** → Yes, but interview expects KMP
5. **"Why LPS and not other preprocessing?"** → LPS optimal for this problem

### Real-World Analogy
**Like finding a specific passage in a book efficiently:**
- You're looking for a specific quote in a large book
- Instead of starting over from the beginning after each mismatch
- You remember how much of the quote you've already matched
- When you find a mismatch, you jump to the longest possible continuation
- Like a smart researcher who doesn't re-read pages unnecessarily

### Human-Readable Pseudocode
```
function strStr(haystack, needle):
    if needle is empty:
        return 0
    
    # Build LPS array
    lps = buildLPS(needle)
    
    i = 0  # haystack index
    j = 0  # needle index
    
    while i < length(haystack):
        if haystack[i] == needle[j]:
            i += 1
            j += 1
            
            if j == length(needle):
                return i - j  # Found match
        
        else:
            if j != 0:
                j = lps[j-1]  # Use LPS to skip
            else:
                i += 1  # Move to next character in haystack
    
    return -1  # No match found

function buildLPS(pattern):
    lps = array of size length(pattern)
    length = 0  # Length of previous longest prefix suffix
    
    i = 1
    while i < length(pattern):
        if pattern[i] == pattern[length]:
            length += 1
            lps[i] = length
            i += 1
        else:
            if length != 0:
                length = lps[length-1]
            else:
                lps[i] = 0
                i += 1
    
    return lps
```

### Execution Visualization

### Example: haystack = "ABABDABACDABABCABAB", needle = "ABABCABAB"
```
KMP Process:

Step 1: Build LPS for "ABABCABAB"
Pattern: A B A B C A B A B
Index:  0 1 2 3 4 5 6 7 8
LPS:    0 0 1 2 0 1 2 3 4

Step 2: Search using LPS
i=0,j=0: A=A ✓ → i=1,j=1
i=1,j=1: B=B ✓ → i=2,j=2  
i=2,j=2: A=A ✓ → i=3,j=3
i=3,j=3: B=B ✓ → i=4,j=4
i=4,j=4: C!=D ✗ → j=LPS[3]=2
i=4,j=2: A=A ✓ → i=5,j=3
i=5,j=3: B=B ✓ → i=6,j=4
i=6,j=4: C!=A ✗ → j=LPS[3]=2
i=6,j=2: A=A ✓ → i=7,j=3
...continue until full match at i=18,j=9

Result: 18-9 = 9 ✓
```

### Key Visualization Points:
- **LPS Construction**: Preprocess pattern to find internal structure
- **Smart Jumps**: Use LPS to skip redundant comparisons
- **No Backtracking**: Text pointer never moves backward
- **Linear Time**: Each character examined at most twice

### LPS Array Visualization:
```
Pattern: A B A B C A B A B
Index:  0 1 2 3 4 5 6 7 8
LPS:    0 0 1 2 0 1 2 3 4

Explanation:
- LPS[0]=0: No proper prefix/suffix for single char
- LPS[1]=0: "A" != "B"
- LPS[2]=1: "A" is prefix and suffix of "ABA"
- LPS[3]=2: "AB" is prefix and suffix of "ABAB"
- LPS[4]=0: No prefix/suffix match for "ABABD"
```

### Time Complexity Breakdown:
- **Standard KMP**: O(N + M) time, O(M) space - optimal approach
- **Alternative KMP**: O(N + M) time, O(M) space - different LPS build
- **Rabin-Karp**: O(N + M) time, O(1) space - hash-based
- **Brute Force**: O(N × M) time, O(1) space - simple but slow

### Alternative Approaches:

#### 1. Rabin-Karp (O(N + M) time, O(1) space)
```go
func strStrRabinKarp(haystack string, needle string) int {
    // Use rolling hash for pattern matching
    // Hash needle and sliding window in haystack
    // Handle hash collisions with verification
    // ... implementation details omitted
}
```
- **Pros**: Constant space, good average case
- **Cons**: Hash collisions require verification

#### 2. Brute Force (O(N × M) time, O(1) space)
```go
func strStrBruteForce(haystack string, needle string) int {
    // Compare needle at each possible position
    // Simple but inefficient
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Too slow for large inputs

#### 3. Built-in Functions (O(N + M) time, O(1) space)
```go
func strStrBuiltIn(haystack string, needle string) int {
    // Use language-provided string search
    // Optimized implementation
    // ... implementation details omitted
}
```
- **Pros}: Highly optimized, reliable
- **Cons**: Not suitable for interview demonstration

### Extensions for Interviews:
- **Multiple Matches**: Find all occurrences of pattern
- **Pattern Counting**: Count total number of matches
- **Pattern Replacement**: Replace all occurrences
- **Case Insensitive**: Handle case-insensitive matching
- **Real-world Applications**: Text editors, search engines, DNA sequencing
*/
func main() {
	// Test cases
	fmt.Println("=== Testing String Search Algorithms ===")
	
	testCases := []struct {
		haystack   string
		needle     string
		description string
	}{
		{"sadbutsad", "sad", "Standard case - multiple matches"},
		{"leetcode", "leeto", "No match"},
		{"", "", "Both empty"},
		{"", "a", "Empty haystack"},
		{"a", "", "Empty needle"},
		{"a", "a", "Single character match"},
		{"a", "b", "Single character no match"},
		{"hello", "ll", "Partial match"},
		{"mississippi", "issi", "Overlapping patterns"},
		{"abcabcabc", "abcabc", "Repeated pattern"},
		{"aaaaa", "aaa", "All same characters"},
		{"abcdefgh", "efg", "Middle match"},
		{"abcdefgh", "h", "End match"},
		{"abcdefgh", "a", "Start match"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Haystack: '%s', Needle: '%s'\n", tc.haystack, tc.needle)
		
		result1 := strStr(tc.haystack, tc.needle)
		result2 := strStrKMPAlternative(tc.haystack, tc.needle)
		result3 := strStrRabinKarp(tc.haystack, tc.needle)
		result4 := strStrBruteForce(tc.haystack, tc.needle)
		result5 := strStrBuiltIn(tc.haystack, tc.needle)
		
		fmt.Printf("  KMP: %d\n", result1)
		fmt.Printf("  KMP Alternative: %d\n", result2)
		fmt.Printf("  Rabin-Karp: %d\n", result3)
		fmt.Printf("  Brute Force: %d\n", result4)
		fmt.Printf("  Built-in: %d\n\n", result5)
	}
	
	// Detailed KMP tracing
	fmt.Println("=== KMP Detailed Tracing ===")
	_, trace := strStrKMPDetailed("ababcabcabababd", "ababd")
	for _, step := range trace {
		fmt.Printf("  %s\n", step)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	longHaystack := ""
	longNeedle := ""
	
	// Build long strings
	for i := 0; i < 1000; i++ {
		longHaystack += "abcdefghijklmnopqrstuvwxyz"
	}
	for i := 0; i < 100; i++ {
		longNeedle += "xyz"
	}
	
	fmt.Printf("Long haystack length: %d\n", len(longHaystack))
	fmt.Printf("Long needle length: %d\n", len(longNeedle))
	
	result := strStr(longHaystack, longNeedle)
	fmt.Printf("KMP result: %d\n", result)
	
	result = strStrRabinKarp(longHaystack, longNeedle)
	fmt.Printf("Rabin-Karp result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long needle
	veryLongNeedle := ""
	for i := 0; i < 10000; i++ {
		veryLongNeedle += "a"
	}
	
	fmt.Printf("Very long needle: %d\n", strStr("aaaaa", veryLongNeedle))
	
	// Unicode characters
	unicodeHaystack := "你好世界你好世界"
	unicodeNeedle := "世界"
	fmt.Printf("Unicode search: %d\n", strStr(unicodeHaystack, unicodeNeedle))
	
	// Case sensitivity
	caseHaystack := "HelloWorld"
	caseNeedle := "hello"
	fmt.Printf("Case sensitive: %d\n", strStr(caseHaystack, caseNeedle))
	
	// Multiple matches
	multiHaystack := "abcabcabcabc"
	multiNeedle := "abc"
	fmt.Printf("Multiple matches: %d\n", strStr(multiHaystack, multiNeedle))
}
