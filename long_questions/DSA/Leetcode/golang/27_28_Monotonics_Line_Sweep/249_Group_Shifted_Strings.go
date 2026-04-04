package main

import "fmt"

// 249. Group Shifted Strings - Hashing with String Processing
// Time: O(N * L), Space: O(N * L) where N is number of strings, L is average length
func groupStrings(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	// Map to store groups by key
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getShiftKey(s)
		groups[key] = append(groups[key], s)
	}
	
	// Convert map to slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

// Get shift key for a string
func getShiftKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Calculate shift from first character
	shift := s[0] - 'a'
	
	key := make([]byte, len(s))
	key[0] = 'a' // Normalize first character to 'a'
	
	for i := 1; i < len(s); i++ {
		// Calculate normalized character
		normalized := (s[i]-'a'-shift+26)%26 + 'a'
		key[i] = normalized
	}
	
	return string(key)
}

// Alternative approach using differences
func groupStringsDifferences(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getDifferenceKey(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getDifferenceKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	key := make([]byte, len(s)-1)
	
	for i := 1; i < len(s); i++ {
		diff := (s[i] - s[i-1] + 26) % 26
		key[i-1] = diff + 'a' // Convert to character
	}
	
	return string(key)
}

// Optimized version with sorting
func groupStringsOptimized(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getShiftKeyOptimized(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getShiftKeyOptimized(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	shift := s[0] - 'a'
	key := make([]byte, len(s))
	
	for i := range s {
		key[i] = (s[i]-'a'-shift+26)%26 + 'a'
	}
	
	return string(key)
}

// Version that handles all characters (not just lowercase)
func groupStringsAllChars(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getUniversalShiftKey(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getUniversalShiftKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Find the minimum character to normalize
	minChar := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
		}
	}
	
	shift := minChar - 'a'
	key := make([]byte, len(s))
	
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			key[i] = (s[i]-'a'-shift+26)%26 + 'a'
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			key[i] = (s[i]-'A'-shift+26)%26 + 'a'
		} else {
			key[i] = s[i] // Keep non-alphabetic characters as is
		}
	}
	
	return string(key)
}

// Brute force approach for comparison
func groupStringsBruteForce(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	used := make([]bool, len(strings))
	result := [][]string{}
	
	for i := 0; i < len(strings); i++ {
		if used[i] {
			continue
		}
		
		group := []string{strings[i]}
		used[i] = true
		
		// Find all strings that can be shifted to match strings[i]
		for j := i + 1; j < len(strings); j++ {
			if !used[j] && canShift(strings[i], strings[j]) {
				group = append(group, strings[j])
				used[j] = true
			}
		}
		
		result = append(result, group)
	}
	
	return result
}

func canShift(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	if len(s1) == 1 {
		return true
	}
	
	shift := (s2[0] - s1[0] + 26) % 26
	
	for i := 1; i < len(s1); i++ {
		if (s2[i]-s1[i]+26)%26 != shift {
			return false
		}
	}
	
	return true
}

// Version with detailed explanation
func groupStringsWithExplanation(strings []string) ([][]string, []string) {
	var explanation []string
	
	if len(strings) == 0 {
		explanation = append(explanation, "Empty input, returning empty result")
		return [][]string{}, explanation
	}
	
	explanation = append(explanation, fmt.Sprintf("Processing %d strings", len(strings)))
	
	groups := make(map[string][]string)
	
	for i, s := range strings {
		explanation = append(explanation, fmt.Sprintf("String %d: '%s'", i, s))
		
		key := getShiftKey(s)
		explanation = append(explanation, fmt.Sprintf("  Shift key: '%s'", key))
		
		groups[key] = append(groups[key], s)
		explanation = append(explanation, fmt.Sprintf("  Added to group with key '%s'", key))
	}
	
	explanation = append(explanation, fmt.Sprintf("Found %d groups", len(groups)))
	
	result := make([][]string, 0, len(groups))
	for key, group := range groups {
		result = append(result, group)
		explanation = append(explanation, fmt.Sprintf("Group '%s': %v", key, group))
	}
	
	return result, explanation
}

// Helper function to check if two strings are in the same group
func areInSameGroup(s1, s2 string) bool {
	return getShiftKey(s1) == getShiftKey(s2)
}

// Find all possible shifts of a string
func getAllShifts(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	
	shifts := make([]string, 26)
	
	for shift := 0; shift < 26; shift++ {
		shifted := make([]byte, len(s))
		
		for i := range s {
			shifted[i] = (s[i]-'a'+byte(shift))%26 + 'a'
		}
		
		shifts[shift] = string(shifted)
	}
	
	return shifts
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: String Grouping with Shift Key
- **Shift Key Generation**: Create canonical representation of string shifts
- **Hashing**: Group strings by their shift key for O(1) lookup
- **Character Normalization**: Shift all characters to normalize to 'a'
- **Difference Pattern**: Track character differences instead of absolute values

## 2. PROBLEM CHARACTERISTICS
- **String Shifts**: Two strings are in same group if one can be shifted to match
- **Circular Alphabet**: Shifts wrap around from 'z' to 'a'
- **Canonical Form**: Each string has a unique normalized representation
- **Grouping Logic**: Strings with same canonical form belong to same group

## 3. SIMILAR PROBLEMS
- Valid Anagram (LeetCode 242) - Check if strings are anagrams
- Group Anagrams (LeetCode 49) - Group anagrams together
- Find All Anagrams (LeetCode 438) - Find all anagrams in string
- String Transform (LeetCode 1153) - Minimum steps to convert strings

## 4. KEY OBSERVATIONS
- **Shift Invariance**: Character differences between consecutive characters are preserved
- **Canonical Key**: Normalizing first character to 'a' creates unique key
- **Hash Grouping**: Perfect use case for hashmap grouping
- **Circular Arithmetic**: Use modulo 26 for alphabet wrap-around

## 5. VARIATIONS & EXTENSIONS
- **Different Alphabets**: Support for uppercase, extended characters
- **Custom Shift Rules**: Support for different shift definitions
- **Multi-dimensional**: Group by multiple transformation types
- **Performance Optimization**: Minimize string allocations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Alphabet size? Case sensitivity? Empty strings?"
- Edge cases: empty array, single character strings, identical strings
- Time complexity: O(N * L) where N=number of strings, L=average length
- Space complexity: O(N * L) for hashmap storage
- Key insight: character differences are invariant under shifts

## 7. COMMON MISTAKES
- Wrong modulo arithmetic causing index out of bounds
- Not handling circular alphabet correctly
- Inconsistent shift key generation
- Missing edge cases (empty strings, single characters)
- Not handling mixed case or special characters

## 8. OPTIMIZATION STRATEGIES
- **Basic Hashing**: O(N * L) time, O(N * L) space - standard
- **Difference Pattern**: O(N * L) time, O(N * L) space - more robust
- **Character Pooling**: Reuse character arrays for memory efficiency
- **Early Termination**: Skip obvious non-matches

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like recognizing shifted versions of the same word:**
- You have words that might be shifted versions of each other
- "abc" shifted by 1 becomes "bcd", shifted by 2 becomes "cde"
- But the pattern of character differences stays the same
- Normalize all words by shifting their first letter to 'a'
- Words with same normalized pattern belong to the same shift family
- Like a cipher where each letter is shifted by the same amount

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of strings
2. **Goal**: Group strings that can be shifted to match each other
3. **Rules**: Shift means adding same amount to each character (circular alphabet)
4. **Output**: Groups of strings with same shift pattern

#### Phase 2: Key Insight Recognition
- **"Character differences invariant"** → Differences between consecutive characters preserved under shifts
- **"Normalization possible"** → Can shift any string to start with 'a'
- **"Hash grouping natural"** → Same normalized strings belong to same group
- **"Circular arithmetic"** → Use modulo 26 for alphabet wrap-around

#### Phase 3: Strategy Development
```
Human thought process:
"I need to group strings that are shifts of each other.
Direct pairwise comparison would be O(N² * L²).

Shift Key Approach:
1. For each string, calculate shift key
2. Shift key: normalize first character to 'a', apply same shift to all characters
3. Use hashmap to group by shift key
4. Strings with same shift key belong to same group

This gives O(N * L) time!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single character**: All single characters in same group
- **Empty strings**: Handle as valid strings
- **Mixed case**: Clarify case handling requirements

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Strings: ["abc", "bcd", "acef", "xyz"]

Human thinking:
"Shift Key Approach:
1. String 'abc':
   - First char 'a' → shift = 'a' - 'a' = 0
   - Apply shift 0 to all chars: 'a','b','c' → key = "abc"

2. String 'bcd':
   - First char 'b' → shift = 'a' - 'b' = -1 = 25
   - Apply shift 25 to all chars: 'b'→'a', 'c'→'b', 'd'→'c' → key = "abc"

3. String 'acef':
   - First char 'a' → shift = 'a' - 'a' = 0
   - Apply shift 0 to all chars: 'a','c','e','f' → key = "acef"

4. String 'xyz':
   - First char 'x' → shift = 'a' - 'x' = -23 = 3
   - Apply shift 3 to all chars: 'x'→'a', 'y'→'b', 'z'→'c' → key = "abc"

Groups:
- "abc", "bcd", "xyz" → key "abc" → group 1
- "acef" → key "acef" → group 2

Result: [["abc", "bcd", "xyz"], ["acef"]] ✓"
```

#### Phase 6: Intuition Validation
- **Why shift key works**: Character differences preserved under uniform shifts
- **Why normalization works**: Shifting first char to 'a' creates canonical form
- **Why hashmap works**: Perfect for grouping identical keys
- **Why O(N * L)**: Each string processed once, hashmap operations O(1)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not pairwise comparison?"** → O(N² * L²) vs O(N * L)
2. **"Should I sort strings first?"** → Not needed, hashmap handles grouping
3. **"What about different shift definitions?"** → Clarify shift rules
4. **"Can I optimize further?"** → O(N * L) is optimal for this approach
5. **"What about uppercase?"** → Clarify case sensitivity

### Real-World Analogy
**Like grouping shifted versions of passwords or codes:**
- You have passwords that might be Caesar cipher versions of each other
- "password123" shifted by 1 becomes "qbttxps234"
- "hello" shifted by 2 becomes "jgnnq"
- But the pattern of character relationships stays the same
- Normalize all passwords by shifting their first letter back to 'a'
- Passwords with same normalized pattern belong to the same shift family
- Like detecting if multiple accounts use related passwords

### Human-Readable Pseudocode
```
function groupStrings(strings):
    if len(strings) == 0:
        return []
    
    groups = hashmap()
    
    for s in strings:
        if len(s) == 0:
            groups[""].append(s)
            continue
            
        # Calculate shift to normalize first character to 'a'
        shift = 'a' - s[0]
        
        # Apply shift to all characters
        key = ""
        for char in s:
            if char is alphabetic:
                normalized = (char - 'a' + shift) % 26 + 'a'
                key += normalized
            else:
                key += char  # Keep non-alphabetic as is
        
        groups[key].append(s)
    
    return list(groups.values())
```

### Execution Visualization

### Example: Strings = ["abc", "bcd", "acef", "xyz"]
```
Shift Key Calculation:
"abc":
- First char 'a' → shift = 'a' - 'a' = 0
- Apply shift 0: 'a'→'a', 'b'→'b', 'c'→'c' → key = "abc"

"bcd":
- First char 'b' → shift = 'a' - 'b' = -1 = 25
- Apply shift 25: 'b'→'a', 'c'→'b', 'd'→'c' → key = "abc"

"acef":
- First char 'a' → shift = 'a' - 'a' = 0
- Apply shift 0: 'a'→'a', 'c'→'c', 'e'→'e', 'f'→'f' → key = "acef"

"xyz":
- First char 'x' → shift = 'a' - 'x' = -23 = 3
- Apply shift 3: 'x'→'a', 'y'→'b', 'z'→'c' → key = "abc"

Hashmap Groups:
"abc": ["abc", "bcd", "xyz"]
"acef": ["acef"]

Final Result: [["abc", "bcd", "xyz"], ["acef"]] ✓
```

### Key Visualization Points:
- **Shift Calculation**: Normalize first character to 'a'
- **Character Normalization**: Apply same shift to all characters
- **Hash Grouping**: Group by normalized key
- **Circular Arithmetic**: Use modulo 26 for alphabet wrap-around

### Memory Layout Visualization:
```
Shift Key Generation:
For "bcd":
- Original: b c d
- Shift: 'a' - 'b' = -1 = 25
- Normalized: (b+25)%26='a', (c+25)%26='b', (d+25)%26='c'
- Key: "abc"

Hashmap Structure:
"abc" → ["abc", "bcd", "xyz"]
"acef" → ["acef"]

Character Processing:
Original: b c d
Shift: +25 +25 +25
Result: a b c (wrapped around alphabet)
```

### Time Complexity Breakdown:
- **Shift Key Generation**: O(N * L) time, O(L) space per string
- **Hashmap Operations**: O(1) average time per insertion/lookup
- **Total**: O(N * L) time, O(N * L) space
- **Optimal**: Each string must be processed at least once

### Alternative Approaches:

#### 1. Pairwise Comparison (O(N² * L²) time, O(1) space)
```go
func groupStringsPairwise(strings []string) [][]string {
    // Compare every pair of strings for shift relationship
    // Inefficient for large inputs
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand
- **Cons**: Quadratic time complexity

#### 2. Sorting Approach (O(N * L log N) time, O(N * L) space)
```go
func groupStringsSorted(strings []string) [][]string {
    // Sort strings, then group consecutive shifts
    // More complex than hashmap approach
    // ... implementation details omitted
}
```
- **Pros**: Deterministic order
- **Cons**: More complex implementation

#### 3. Trie-Based (O(N * L) time, O(N * L) space)
```go
func groupStringsTrie(strings []string) [][]string {
    // Build trie of all shifted versions
    // Overkill for this problem
    // ... implementation details omitted
}
```
- **Pros**: Fast lookup for many queries
- **Cons**: Complex implementation, memory overhead

### Extensions for Interviews:
- **Different Alphabets**: Support for uppercase, extended characters
- **Custom Shift Rules**: Support for different shift definitions
- **Multi-dimensional**: Group by multiple transformation types
- **Performance Optimization**: Minimize string allocations
- **Real-world Applications**: Cipher detection, password analysis
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Group Shifted Strings ===")
	
	testCases := []struct {
		strings    []string
		description string
	}{
		{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, "Standard case"},
		{[]string{"a"}, "Single character"},
		{[]string{"abc", "def", "ghi"}, "All same length"},
		{[]string{"a", "b", "c", "d", "e"}, "All single characters"},
		{[]string{"abc", "bcd", "cde", "def"}, "Sequential shifts"},
		{[]string{"az", "ba", "ab", "bc"}, "Mixed shifts"},
		{[]string{}, "Empty array"},
		{[]string{"abc", "abc", "abc"}, "Duplicate strings"},
		{[]string{"abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"}, "Full alphabet"},
		{[]string{"az", "by", "cx", "dw"}, "Different patterns"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Input: %v\n", tc.strings)
		
		result1 := groupStrings(tc.strings)
		result2 := groupStringsDifferences(tc.strings)
		result3 := groupStringsOptimized(tc.strings)
		result4 := groupStringsBruteForce(tc.strings)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Differences: %v\n", result2)
		fmt.Printf("  Optimized: %v\n", result3)
		fmt.Printf("  Brute Force: %v\n\n", result4)
	}
	
	// Detailed explanation
	fmt.Println("=== Detailed Explanation ===")
	testStrings := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}
	result, explanation := groupStringsWithExplanation(testStrings)
	
	fmt.Printf("Result: %v\n", result)
	for _, step := range explanation {
		fmt.Printf("  %s\n", step)
	}
	
	// Test helper functions
	fmt.Println("\n=== Helper Functions Test ===")
	
	fmt.Printf("Are 'abc' and 'bcd' in same group? %t\n", areInSameGroup("abc", "bcd"))
	fmt.Printf("Are 'abc' and 'xyz' in same group? %t\n", areInSameGroup("abc", "xyz"))
	
	fmt.Printf("All shifts of 'abc': %v\n", getAllShifts("abc"))
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeStrings := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		base := "abc"
		shift := i % 26
		shifted := make([]byte, len(base))
		for j := range base {
			shifted[j] = (base[j]-'a'+byte(shift))%26 + 'a'
		}
		largeStrings[i] = string(shifted)
	}
	
	fmt.Printf("Large test with %d strings\n", len(largeStrings))
	
	result = groupStrings(largeStrings)
	fmt.Printf("Found %d groups\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Mixed case strings
	mixedCase := []string{"ABC", "BCD", "CDE", "abc", "bcd"}
	fmt.Printf("Mixed case: %v\n", groupStringsAllChars(mixedCase))
	
	// Strings with numbers
	withNumbers := []string{"a1b", "b1c", "c1d"}
	fmt.Printf("With numbers: %v\n", groupStrings(withNumbers))
	
	// Very long strings
	longStrings := []string{
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		"bcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyza",
	}
	fmt.Printf("Long strings: %v\n", groupStrings(longStrings))
	
	// Test with all possible single characters
	fmt.Println("\n=== All Single Characters ===")
	var allChars []string
	for i := 0; i < 26; i++ {
		allChars = append(allChars, string('a'+i))
	}
	
	result = groupStrings(allChars)
	fmt.Printf("All single characters grouped: %v\n", result)
}
