package main

import (
	"fmt"
	"math"
)

// 459. Repeated Substring Pattern - KMP Algorithm
// Time: O(N), Space: O(N)
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	// Build LPS array
	lps := buildLPS(s)
	
	// Check if the string can be constructed by repeating a substring
	lpsValue := lps[len(s)-1]
	length := len(s)
	
	// If there's a proper prefix which is also suffix
	// and the length of the string is divisible by (length - lpsValue)
	if lpsValue > 0 && length%(length-lpsValue) == 0 {
		return true
	}
	
	return false
}

// Build LPS array
func buildLPS(s string) []int {
	lps := make([]int, len(s))
	length := 0
	
	i := 1
	for i < len(s) {
		if s[i] == s[length] {
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

// Alternative approach using string concatenation
func repeatedSubstringPatternConcat(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	// Concatenate string with itself and check if original appears in middle
	concatenated := s + s
	
	// Search for original string starting from position 1
	for i := 1; i < len(s); i++ {
		if i+len(s) <= len(concatenated) && concatenated[i:i+len(s)] == s {
			return true
		}
	}
	
	return false
}

// Optimized concatenation approach
func repeatedSubstringPatternConcatOptimized(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	concatenated := s + s
	
	// Check if original string appears in concatenated string (excluding first and last positions)
	for i := 1; i < len(s); i++ {
		if concatenated[i:i+len(s)] == s {
			return true
		}
	}
	
	return false
}

// Brute force approach
func repeatedSubstringPatternBruteForce(s string) bool {
	n := len(s)
	
	// Try all possible substring lengths
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue // Skip if length doesn't divide n
		}
		
		// Check if repeating pattern works
		pattern := s[0:length]
		valid := true
		
		for i := length; i < n; i += length {
			if s[i:i+length] != pattern {
				valid = false
				break
			}
		}
		
		if valid {
			return true
		}
	}
	
	return false
}

// Mathematical approach using divisors
func repeatedSubstringPatternMath(s string) bool {
	n := len(s)
	
	// Find all divisors of n (except n itself)
	for length := 1; length <= n/2; length++ {
		if n%length == 0 {
			// Check if this length forms a repeating pattern
			pattern := s[0:length]
			valid := true
			
			for i := length; i < n; i += length {
				if s[i:i+length] != pattern {
					valid = false
					break
				}
			}
			
			if valid {
				return true
			}
		}
	}
	
	return false
}

// Rolling hash approach
func repeatedSubstringPatternRollingHash(s string) bool {
	n := len(s)
	
	if n <= 1 {
		return false
	}
	
	// Try all possible substring lengths
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue
		}
		
		// Calculate hash of first substring
		hash := 0
		base := 26
		mod := 1000000007
		
		for i := 0; i < length; i++ {
			hash = (hash*base + int(s[i]-'a')) % mod
		}
		
		// Check if all substrings have same hash
		valid := true
		currentHash := hash
		
		for i := length; i < n; i += length {
			// Calculate hash of current substring
			subHash := 0
			for j := i; j < i+length; j++ {
				subHash = (subHash*base + int(s[j]-'a')) % mod
			}
			
			if subHash != currentHash {
				valid = false
				break
			}
		}
		
		if valid {
			// Verify to avoid hash collision
			pattern := s[0:length]
			for i := length; i < n; i += length {
				if s[i:i+length] != pattern {
					valid = false
					break
				}
			}
			
			if valid {
				return true
			}
		}
	}
	
	return false
}

// KMP with detailed explanation
func repeatedSubstringPatternKMPDetailed(s string) (bool, []string) {
	var explanation []string
	
	if len(s) <= 1 {
		explanation = append(explanation, "String length <= 1, cannot have repeated pattern")
		return false, explanation
	}
	
	explanation = append(explanation, fmt.Sprintf("Building LPS array for: %s", s))
	
	lps := buildLPSWithExplanation(s, &explanation)
	explanation = append(explanation, fmt.Sprintf("LPS array: %v", lps))
	
	lpsValue := lps[len(s)-1]
	length := len(s)
	
	explanation = append(explanation, fmt.Sprintf("LPS[-1] = %d, Length = %d", lpsValue, length))
	
	if lpsValue > 0 {
		explanation = append(explanation, "There is a proper prefix which is also suffix")
		
		if length%(length-lpsValue) == 0 {
			explanation = append(explanation, fmt.Sprintf("Length %% (Length - LPS[-1]) = %d %% %d = 0", length, length-lpsValue))
			explanation = append(explanation, "String can be constructed by repeating a substring")
			return true, explanation
		} else {
			explanation = append(explanation, fmt.Sprintf("Length %% (Length - LPS[-1]) = %d %% %d != 0", length, length-lpsValue))
			explanation = append(explanation, "Cannot construct by repeating a substring")
		}
	} else {
		explanation = append(explanation, "No proper prefix which is also suffix")
	}
	
	return false, explanation
}

func buildLPSWithExplanation(s string, explanation *[]string) []int {
	lps := make([]int, len(s))
	length := 0
	
	i := 1
	for i < len(s) {
		*explanation = append(*explanation, fmt.Sprintf("i=%d, length=%d, s[i]='%c', s[length]='%c'", 
			i, length, s[i], s[length]))
		
		if s[i] == s[length] {
			length++
			lps[i] = length
			*explanation = append(*explanation, fmt.Sprintf("Match! LPS[%d]=%d", i, length))
			i++
		} else {
			if length != 0 {
				*explanation = append(*explanation, fmt.Sprintf("Mismatch, using LPS[%d]=%d", length-1, lps[length-1]))
				length = lps[length-1]
			} else {
				lps[i] = 0
				*explanation = append(*explanation, fmt.Sprintf("No match, LPS[%d]=0", i))
				i++
			}
		}
	}
	
	return lps
}

// Find the actual repeating substring
func findRepeatingSubstring(s string) string {
	n := len(s)
	
	if n <= 1 {
		return ""
	}
	
	lps := buildLPS(s)
	lpsValue := lps[len(s)-1]
	
	if lpsValue > 0 && n%(n-lpsValue) == 0 {
		return s[0 : n-lpsValue]
	}
	
	return ""
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: String Pattern Detection with KMP
- **KMP LPS Array**: Longest Prefix Suffix for pattern analysis
- **String Repetition**: Check if string is composed of repeated substring
- **Divisibility Check**: Length must be divisible by pattern length
- **Concatenation Trick**: String appears in doubled string if pattern exists

## 2. PROBLEM CHARACTERISTICS
- **Pattern Recognition**: Find if string consists of repeated pattern
- **String Analysis**: Need to detect internal structure
- **Mathematical Property**: Length divisibility is key
- **Efficient Detection**: Avoid O(N²) brute force approach

## 3. SIMILAR PROBLEMS
- Find the Index of the First Occurrence in a String (LeetCode 28) - KMP string matching
- Implement strStr() (LeetCode 28) - Same problem
- Repeated String Match (LeetCode 686) - Pattern repetition counting
- Detect Squares (LeetCode 593) - Pattern detection in strings

## 4. KEY OBSERVATIONS
- **LPS Critical**: Last LPS value reveals longest proper prefix-suffix
- **Divisibility Required**: Length must be multiple of pattern length
- **Pattern Length**: Pattern length = n - LPS[n-1] if repeating exists
- **Concatenation Insight**: String appears in s+s if and only if it has repeating pattern

## 5. VARIATIONS & EXTENSIONS
- **KMP Approach**: Use LPS array for O(N) solution
- **Concatenation**: String doubling trick for O(N) solution
- **Brute Force**: Try all divisors of length
- **Mathematical**: Divisor-based pattern checking

## 6. INTERVIEW INSIGHTS
- Always clarify: "String length constraints? Character set? Performance requirements?"
- Edge cases: empty string, single character, no pattern
- Time complexity: O(N) for KMP and concatenation, O(N√N) for brute force
- Space complexity: O(N) for KMP, O(N) for concatenation
- Key insight: LPS array reveals internal string structure

## 7. COMMON MISTAKES
- Wrong LPS array construction
- Missing divisibility check
- Not handling edge cases properly
- Incorrect pattern length calculation
- Not verifying actual pattern repetition

## 8. OPTIMIZATION STRATEGIES
- **KMP LPS**: O(N) time, O(N) space - optimal approach
- **Concatenation**: O(N) time, O(N) space - elegant solution
- **Brute Force**: O(N√N) time, O(1) space - simple but slow
- **Mathematical**: O(N√N) time, O(1) space - divisor-based

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like analyzing a wallpaper pattern:**
- You want to know if the wallpaper is made by repeating a smaller pattern
- You look for the longest prefix that also appears as a suffix
- If the total length is a multiple of the remaining length, it's repeating
- Like checking if a fabric pattern is tiled from a smaller design
- The LPS array tells you how much of the pattern overlaps with itself

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String to analyze for repeating patterns
2. **Goal**: Determine if string consists of repeated substring
3. **Rules**: Must be exact repetition of substring
4. **Output**: Boolean indicating if repeating pattern exists

#### Phase 2: Key Insight Recognition
- **"LPS reveals structure"** → Longest prefix-suffix shows pattern overlap
- **"Divisibility critical"** → Total length must be multiple of pattern length
- **"Pattern length formula"** → Pattern length = n - LPS[n-1]
- **"Concatenation trick"** → String appears in doubled string if pattern exists

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if string is made of repeated pattern.
Brute force: try all possible pattern lengths O(N²).

KMP LPS Approach:
1. Build LPS array to find longest proper prefix-suffix
2. Check if LPS[n-1] > 0 (there is overlap)
3. Verify n % (n - LPS[n-1]) == 0 (divisible)
4. If both true, string has repeating pattern

This gives O(N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return false (no pattern possible)
- **Single character**: Return false (cannot repeat)
- **No overlap**: LPS[n-1] = 0, return false
- **Not divisible**: Even with overlap, return false

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: s = "abcabcabc"

Human thinking:
"KMP LPS Approach:
Step 1: Build LPS array
String: a b c a b c a b c
Index: 0 1 2 3 4 5 6 7 8
LPS:   0 0 0 1 2 3 4 5 6

Step 2: Check LPS[n-1] = LPS[8] = 6 > 0 ✓
There is overlap: "abcabc" is both prefix and suffix

Step 3: Check divisibility
n = 9, LPS[n-1] = 6
Pattern length = 9 - 6 = 3
9 % 3 = 0 ✓

Step 4: Both conditions true, pattern exists!
Pattern is "abc" (first 3 characters)

Result: true ✓"
```

#### Phase 6: Intuition Validation
- **Why LPS works**: Reveals longest self-overlapping pattern
- **Why divisibility**: Ensures pattern fits exactly into string
- **Why pattern length formula**: n - LPS[n-1] gives minimal repeating unit
- **Why O(N)**: Single pass to build LPS, constant time checks

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all lengths?"** → O(N²) vs O(N) with KMP
2. **"Should I use concatenation?"** → Yes, elegant O(N) solution
3. **"What about LPS construction?"** → Critical for pattern detection
4. **"Can I skip divisibility check?"** → No, necessary condition
5. **"Why pattern length formula?"** → Derives from LPS properties

### Real-World Analogy
**Like analyzing a musical composition:**
- You're listening to a piece of music and want to know if it's based on a repeating motif
- You find the longest phrase that repeats at the beginning and end
- If the total length is a multiple of the motif length, it's a repeating pattern
- Like identifying if a song is built from a short musical phrase
- The LPS array is like finding the overlapping musical themes

### Human-Readable Pseudocode
```
function repeatedSubstringPattern(s):
    if length(s) <= 1:
        return false
    
    # Build LPS array
    lps = buildLPS(s)
    
    # Get LPS value at last position
    lpsValue = lps[length(s) - 1]
    n = length(s)
    
    # Check conditions for repeating pattern
    if lpsValue > 0 and n % (n - lpsValue) == 0:
        return true
    
    return false

function buildLPS(s):
    lps = array of size length(s)
    length = 0
    
    i = 1
    while i < length(s):
        if s[i] == s[length]:
            length += 1
            lps[i] = length
            i += 1
        else:
            if length != 0:
                length = lps[length - 1]
            else:
                lps[i] = 0
                i += 1
    
    return lps
```

### Execution Visualization

### Example: s = "abcabcabc"
```
KMP LPS Process:

Step 1: Build LPS array
String: a b c a b c a b c
Index: 0 1 2 3 4 5 6 7 8

LPS construction:
i=1, length=0: s[1]='b' != s[0]='a' → LPS[1]=0
i=2, length=0: s[2]='c' != s[0]='a' → LPS[2]=0
i=3, length=0: s[3]='a' == s[0]='a' → length=1, LPS[3]=1
i=4, length=1: s[4]='b' == s[1]='b' → length=2, LPS[4]=2
i=5, length=2: s[5]='c' == s[2]='c' → length=3, LPS[5]=3
i=6, length=3: s[6]='a' == s[3]='a' → length=4, LPS[6]=4
i=7, length=4: s[7]='b' == s[4]='b' → length=5, LPS[7]=5
i=8, length=5: s[8]='c' == s[5]='c' → length=6, LPS[8]=6

Final LPS: [0,0,0,1,2,3,4,5,6]

Step 2: Check conditions
LPS[n-1] = LPS[8] = 6 > 0 ✓
Pattern length = 9 - 6 = 3
9 % 3 = 0 ✓

Result: true ✓ (pattern is "abc")
```

### Key Visualization Points:
- **LPS Construction**: Build array showing longest prefix-suffix at each position
- **Overlap Detection**: LPS[n-1] > 0 indicates self-overlap
- **Pattern Length**: n - LPS[n-1] gives minimal repeating unit
- **Divisibility Check**: Ensures pattern fits exactly

### Pattern Detection Visualization:
```
String: "abcabcabc"
LPS:   [0,0,0,1,2,3,4,5,6]

LPS[8] = 6 means:
- Prefix of length 6: "abcabc"
- Suffix of length 6: "abcabc"
- They are the same!

Pattern length = 9 - 6 = 3
Pattern = "abc"

Verification: "abc" + "abc" + "abc" = "abcabcabc" ✓
```

### Time Complexity Breakdown:
- **KMP LPS**: O(N) time, O(N) space - optimal approach
- **Concatenation**: O(N) time, O(N) space - elegant solution
- **Brute Force**: O(N√N) time, O(1) space - simple but slow
- **Mathematical**: O(N√N) time, O(1) space - divisor-based

### Alternative Approaches:

#### 1. Concatenation Trick (O(N) time, O(N) space)
```go
func repeatedSubstringPatternConcat(s string) bool {
    // If s appears in s+s (excluding first and last positions)
    // then s has a repeating pattern
    concatenated := s + s
    
    for i := 1; i < len(s); i++ {
        if concatenated[i:i+len(s)] == s {
            return true
        }
    }
    return false
}
```
- **Pros**: Elegant, easy to understand
- **Cons**: Requires extra space for concatenation

#### 2. Brute Force (O(N√N) time, O(1) space)
```go
func repeatedSubstringPatternBruteForce(s string) bool {
    n := len(s)
    
    // Try all possible pattern lengths (divisors of n)
    for length := 1; length <= n/2; length++ {
        if n % length != 0 {
            continue
        }
        
        // Check if pattern repeats
        pattern := s[0:length]
        valid := true
        
        for i := length; i < n; i += length {
            if s[i:i+length] != pattern {
                valid = false
                break
            }
        }
        
        if valid {
            return true
        }
    }
    
    return false
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Too slow for large strings

#### 3. Mathematical Divisors (O(N√N) time, O(1) space)
```go
func repeatedSubstringPatternMath(s string) bool {
    n := len(s)
    
    // Find all divisors of n
    for length := 1; length <= n/2; length++ {
        if n % length == 0 {
            // Check if this length forms a pattern
            // ... same as brute force
        }
    }
    
    return false
}
```
- **Pros**: Systematic approach
- **Cons**: Same complexity as brute force

### Extensions for Interviews:
- **Find Pattern**: Return the actual repeating substring
- **Count Repetitions**: How many times does the pattern repeat?
- **Pattern Variations**: Find all possible repeating patterns
- **Partial Patterns**: Handle partial repetitions
- **Real-world Applications**: Text compression, pattern recognition
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Repeated Substring Pattern ===")
	
	testCases := []struct {
		s          string
		expected   bool
		description string
	}{
		{"abab", true, "Simple repeating pattern"},
		{"aba", false, "Not a repeating pattern"},
		{"abcabcabcabc", true, "Long repeating pattern"},
		{"aaaa", true, "All same characters"},
		{"abcabc", true, "Two repetitions"},
		{"abac", false, "No pattern"},
		{"a", false, "Single character"},
		{"", false, "Empty string"},
		{"xyzxyzxyz", true, "Three repetitions"},
		{"abcd", false, "All unique"},
		{"zzzzzz", true, "Many same characters"},
		{"abcab", false, "Partial match"},
		{"abcabcabcabcabc", true, "Five repetitions"},
		{"ababa", false, "Overlapping but not repeating"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  String: '%s'\n", tc.s)
		
		result1 := repeatedSubstringPattern(tc.s)
		result2 := repeatedSubstringPatternConcat(tc.s)
		result3 := repeatedSubstringPatternBruteForce(tc.s)
		result4 := repeatedSubstringPatternMath(tc.s)
		result5 := repeatedSubstringPatternRollingHash(tc.s)
		
		fmt.Printf("  KMP: %t\n", result1)
		fmt.Printf("  Concatenation: %t\n", result2)
		fmt.Printf("  Brute Force: %t\n", result3)
		fmt.Printf("  Mathematical: %t\n", result4)
		fmt.Printf("  Rolling Hash: %t\n", result5)
		
		// Check if results match expected
		if result1 == tc.expected {
			fmt.Printf("  ✓ Expected: %t\n", tc.expected)
		} else {
			fmt.Printf("  ✗ Expected: %t, Got: %t\n", tc.expected, result1)
		}
		
		// Show repeating substring if found
		if result1 {
			pattern := findRepeatingSubstring(tc.s)
			fmt.Printf("  Repeating pattern: '%s'\n", pattern)
		}
		
		fmt.Println()
	}
	
	// KMP detailed explanation
	fmt.Println("=== KMP Detailed Explanation ===")
	_, explanation := repeatedSubstringPatternKMPDetailed("abcabcabc")
	for _, step := range explanation {
		fmt.Printf("  %s\n", step)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += "abc"
	}
	
	fmt.Printf("Long string length: %d\n", len(longString))
	fmt.Printf("KMP result: %t\n", repeatedSubstringPattern(longString))
	fmt.Printf("Concatenation result: %t\n", repeatedSubstringPatternConcat(longString))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long string with no pattern
	noPattern := ""
	for i := 0; i < 1000; i++ {
		noPattern += string(rune('a' + (i % 26)))
	}
	fmt.Printf("No pattern string: %t\n", repeatedSubstringPattern(noPattern))
	
	// All same characters
	allSame := ""
	for i := 0; i < 10000; i++ {
		allSame += "a"
	}
	fmt.Printf("All same characters: %t\n", repeatedSubstringPattern(allSame))
	
	// Complex pattern
	complex := "ababcababcababc"
	fmt.Printf("Complex pattern: %t\n", repeatedSubstringPattern(complex))
	
	// Find all repeating substrings
	fmt.Println("\n=== Finding Repeating Substrings ===")
	testStrings := []string{"abcabcabc", "ababab", "aaaaaa", "xyzxyzxyzxyz"}
	
	for _, s := range testStrings {
		pattern := findRepeatingSubstring(s)
		if pattern != "" {
			fmt.Printf("'%s' -> pattern: '%s'\n", s, pattern)
		} else {
			fmt.Printf("'%s' -> no repeating pattern\n", s)
		}
	}
}
