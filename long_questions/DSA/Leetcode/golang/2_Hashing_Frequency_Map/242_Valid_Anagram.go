package main

import "fmt"

// 242. Valid Anagram
// Time: O(N), Space: O(1) for ASCII characters (26 for lowercase letters)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Assuming only lowercase English letters
	count := make([]int, 26)
	
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	
	return true
}

// Alternative solution using frequency map for general characters
func isAnagramGeneral(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	count := make(map[rune]int)
	
	for _, char := range s {
		count[char]++
	}
	
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	
	return true
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Character Frequency Comparison
- **Frequency Counting**: Count characters in both strings
- **Single Pass**: Increment for one string, decrement for other
- **Zero Validation**: All counts must be zero for anagram
- **Space Optimization**: Use fixed-size array for limited character set

## 2. PROBLEM CHARACTERISTICS
- **String Comparison**: Check if two strings are anagrams
- **Character Analysis**: Need to compare character frequencies
- **Length Check**: Different lengths immediately disqualify
- **Case Sensitivity**: Typically case-sensitive comparison

## 3. SIMILAR PROBLEMS
- Group Anagrams (LeetCode 49) - Group multiple strings by anagram
- Find All Anagrams in a String (LeetCode 438) - Find all substrings that are anagrams
- Valid Palindrome (LeetCode 125) - Similar character analysis
- Permutation in String (LeetCode 567) - Check if permutation exists

## 4. KEY OBSERVATIONS
- **Length prerequisite**: Different lengths cannot be anagrams
- **Frequency equality**: Anagrams have identical character counts
- **Single array sufficient**: Can use one array with +/- operations
- **Fixed alphabet**: For lowercase letters, array of size 26 suffices

## 5. VARIATIONS & EXTENSIONS
- **Unicode support**: Handle full Unicode character set
- **Case insensitive**: Treat uppercase and lowercase as same
- **Custom alphabets**: Handle different character sets
- **Multiple comparisons**: Check one string against many others

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are strings case-sensitive? What character set?"
- Edge cases: empty strings, single characters, different lengths
- Time complexity: O(N) time, O(1) space for fixed alphabet
- Alternative: Sorting approach O(N log N) time, O(1) space

## 7. COMMON MISTAKES
- Not checking length first (wasted work)
- Using two separate frequency arrays (inefficient)
- Not handling Unicode characters properly
- Forgetting to validate all counts are zero
- Using sorting when frequency counting is better

## 8. OPTIMIZATION STRATEGIES
- **Single array**: Increment for one string, decrement for other
- **Early exit**: Check for negative counts during processing
- **Fixed-size array**: Use array instead of hash map for limited alphabet
- **Length check**: Validate lengths before processing

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like comparing two word puzzles:**
- You have two words made of letter tiles
- You want to know if they use exactly the same letters
- You count how many of each letter each word uses
- If the counts match for all letters, they're anagrams
- Any difference in letter counts means they're not anagrams

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two strings s and t
2. **Goal**: Determine if t is an anagram of s
3. **Output**: Boolean (true if anagram, false otherwise)
4. **Constraint**: Must use exactly the same characters

#### Phase 2: Key Insight Recognition
- **"Length check first"** → Different lengths cannot be anagrams
- **"Frequency counting"** → Anagrams have identical character counts
- **"Single array optimization"** → Can use +/- operations
- **"Fixed alphabet"** → For lowercase letters, array size is constant

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if two strings are anagrams.
First, I'll check if they have the same length.
If not, they can't be anagrams.
Then I'll count character frequencies.
I can use one array: increment for s, decrement for t.
If all counts end up as zero, they're anagrams.
If any count is non-zero, they're not anagrams."
```

#### Phase 4: Edge Case Handling
- **Different lengths**: Immediately return false
- **Empty strings**: Both empty → true, one empty → false
- **Single character**: Must be same character
- **Unicode characters**: Use hash map instead of fixed array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: s = "anagram", t = "nagaram"

Human thinking:
"First, check lengths: both have 7 characters ✓

Now I'll count frequencies using one array:
Initialize count array of size 26 with zeros

Process both strings simultaneously:
a: count['a'-'a']++ → count[0] = 1, count['n'-'a']-- → count[13] = -1
n: count['n'-'a']++ → count[13] = 0, count['a'-'a']-- → count[0] = 0
a: count['a'-'a']++ → count[0] = 1, count['g'-'a']-- → count[6] = -1
g: count['g'-'a']++ → count[6] = 0, count['a'-'a']-- → count[0] = 0
r: count['r'-'a']++ → count[17] = 1, count['r'-'a']-- → count[17] = 0
a: count['a'-'a']++ → count[0] = 1, count['a'-'a']-- → count[0] = 0
m: count['m'-'a']++ → count[12] = 1, count['m'-'a']-- → count[12] = 0

Final check: All counts are zero ✓
They are anagrams!"
```

#### Phase 6: Intuition Validation
- **Why length check works**: Different lengths cannot use same characters
- **Why frequency works**: Anagrams have identical character counts
- **Why single array works**: +/- operations cancel out for matching characters
- **Why O(1) space**: Fixed-size array for limited alphabet

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use sorting?"** → O(N log N) vs O(N) time
2. **"Should I use two arrays?"** → Single array is more efficient
3. **"What about Unicode?"** → Need hash map for full character set
4. **"Can I optimize further?"** → Early exit during processing

### Real-World Analogy
**Like checking if two recipes use same ingredients:**
- You have two recipes with ingredient lists
- You want to know if they use exactly the same ingredients in same quantities
- You count how much of each ingredient each recipe uses
- If all ingredient quantities match, the recipes are "anagrams"
- Any difference means they're different recipes

### Human-Readable Pseudocode
```
function isAnagram(s, t):
    if length(s) != length(t):
        return false
    
    count = array of size 26 filled with 0
    
    for i from 0 to length(s)-1:
        count[s[i]-'a']++
        count[t[i]-'a']--
    
    for c in count:
        if c != 0:
            return false
    
    return true
```

### Execution Visualization

### Example: s = "rat", t = "car"
```
Frequency Array Evolution:
Initial: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

Processing 'r' and 'c':
count['r'-'a']++ → count[17] = 1
count['c'-'a']-- → count[2] = -1
Array: [0, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0]

Processing 'a' and 'a':
count['a'-'a']++ → count[0] = 1
count['a'-'a']-- → count[0] = 0
Array: [0, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0]

Processing 't' and 'r':
count['t'-'a']++ → count[19] = 1
count['r'-'a']-- → count[17] = 0
Array: [0, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0]

Final check: count[2] = -1, count[19] = 1 → NOT all zeros
Result: false (not anagrams)
```

### Key Visualization Points:
- **Length validation**: First check for equal lengths
- **Frequency counting**: Track character differences
- **Single array optimization**: +/- operations cancel out
- **Zero validation**: All counts must be zero for anagrams

### Memory Layout Visualization:
```
Character Mapping:
a → index 0
b → index 1
c → index 2
...
r → index 17
t → index 19

Processing Flow for "rat" vs "car":
Step 1: 'r' vs 'c' → [0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0]
Step 2: 'a' vs 'a' → [0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0]
Step 3: 't' vs 'r' → [0,0,-1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]

Non-zero elements indicate mismatch
```

### Time Complexity Breakdown:
- **Frequency counting**: O(N) time, O(1) space for fixed alphabet
- **Sorting approach**: O(N log N) time, O(N) space
- **Hash map approach**: O(N) time, O(K) space where K is unique characters
- **N**: length of strings

### Alternative Approaches:

#### 1. Sorting Approach (O(N log N) time, O(N) space)
```go
func isAnagramSort(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sRunes := []rune(s)
    tRunes := []rune(t)
    
    sort.Slice(sRunes, func(i, j int) bool {
        return sRunes[i] < sRunes[j]
    })
    sort.Slice(tRunes, func(i, j int) bool {
        return tRunes[i] < tRunes[j]
    })
    
    return string(sRunes) == string(tRunes)
}
```
- **Pros**: Simple to implement
- **Cons**: O(N log N) time, requires sorting

#### 2. Hash Map for Unicode (O(N) time, O(K) space)
```go
func isAnagramUnicode(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    count := make(map[rune]int)
    
    for _, char := range s {
        count[char]++
    }
    
    for _, char := range t {
        count[char]--
        if count[char] < 0 {
            return false
        }
    }
    
    return true
}
```
- **Pros**: Handles full Unicode character set
- **Cons**: More memory than fixed array

#### 3. Early Exit Optimization (O(N) time, O(1) space)
```go
func isAnagramEarlyExit(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    count := make([]int, 26)
    
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }
    
    for _, c := range count {
        if c != 0 {
            return false
        }
    }
    
    return true
}
```
- **Pros**: Same complexity, cleaner code
- **Cons**: No early exit during processing

### Extensions for Interviews:
- **Case Insensitive**: Convert both strings to same case
- **Unicode Support**: Handle full character set with hash map
- **Multiple Comparisons**: Pre-compute signatures for many comparisons
- **Streaming**: Process characters as they arrive
- **Partial Anagrams**: Check if one is subset of another
*/
func main() {
	// Test cases
	testCases := []struct {
		s string
		t string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"a", "a"},
		{"ab", "ba"},
		{"", ""},
		{"abc", "ab"},
		{"Hello", "olleH"},
	}
	
	for i, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		resultGeneral := isAnagramGeneral(tc.s, tc.t)
		fmt.Printf("Test Case %d: \"%s\" & \"%s\" -> Anagram: %t (General: %t)\n", 
			i+1, tc.s, tc.t, result, resultGeneral)
	}
}
