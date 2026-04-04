package main

import "fmt"

// 438. Find All Anagrams in a String (Fixed Size Sliding Window)
// Time: O(N), Space: O(1) for ASCII characters
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	
	result := []int{}
	pCount := make([]int, 26)
	sCount := make([]int, 26)
	
	// Initialize frequency count for pattern and first window
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}
	
	// Check if first window is an anagram
	if matches(pCount, sCount) {
		result = append(result, 0)
	}
	
	// Slide the window through the string
	for i := len(p); i < len(s); i++ {
		// Remove the leftmost character
		sCount[s[i-len(p)]-'a']--
		// Add the new character
		sCount[s[i]-'a']++
		
		// Check if current window is an anagram
		if matches(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}
	
	return result
}

// Helper function to check if two frequency arrays match
func matches(pCount, sCount []int) bool {
	for i := 0; i < 26; i++ {
		if pCount[i] != sCount[i] {
			return false
		}
	}
	return true
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fixed Size Sliding Window with Frequency Arrays
- **Fixed Window**: Window size equals length of pattern string
- **Frequency Arrays**: Track character counts for pattern and current window
- **Sliding Update**: Remove leftmost character, add new rightmost character
- **Efficient Comparison**: Compare frequency arrays in O(26) time

## 2. PROBLEM CHARACTERISTICS
- **Anagram Detection**: Find all substrings that are permutations of pattern
- **Fixed Window Size**: Window size is constant (len(pattern))
- **Character Frequency**: Need exact character count matching
- **All Occurrences**: Return all starting positions of valid anagrams

## 3. SIMILAR PROBLEMS
- Permutation in String (LeetCode 567) - check if any permutation exists
- Find All Anagrams in a String (current problem)
- Minimum Window Substring (variable window)
- Longest Substring with At Most K Distinct Characters

## 4. KEY OBSERVATIONS
- **Fixed window size**: Always len(pattern) characters
- **Character set limited**: Only lowercase letters (26 possible)
- **Frequency matching**: Exact character counts required
- **Sliding efficiency**: O(1) update per slide

## 5. VARIATIONS & EXTENSIONS
- **Unicode Support**: Handle larger character sets
- **Case Insensitive**: Treat uppercase and lowercase as same
- **Multiple Patterns**: Process multiple patterns efficiently
- **Return Substrings**: Return actual substrings instead of indices

## 6. INTERVIEW INSIGHTS
- Always clarify: "Character set? Case sensitivity? Empty strings?"
- Edge cases: pattern longer than string, empty strings, no matches
- Space complexity: O(1) for fixed character set (26)
- Time complexity: O(N) - each character visited once

## 7. COMMON MISTAKES
- Using hash map instead of array (slower for fixed character set)
- Not handling case where pattern longer than string
- Recomputing frequency arrays from scratch (O(N²))
- Forgetting to check first window before sliding
- Not handling empty strings properly

## 8. OPTIMIZATION STRATEGIES
- **Array instead of hash map**: Use fixed-size array for 26 letters
- **Early termination**: Stop if remaining characters insufficient
- **Optimized comparison**: Track mismatch count instead of full comparison
- **Character filtering**: Pre-filter if pattern has limited characters

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding all arrangements of letters that match a pattern:**
- You have a long word (string) and a shorter pattern word
- You want to find all sections of the long word that use exactly the same letters
- The sections must be the same length as the pattern
- Order doesn't matter, just the collection of letters
- Slide a window of pattern length through the long word
- At each position, check if the letters match the pattern

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s (source) and string p (pattern)
2. **Goal**: Find all starting indices of anagrams of p in s
3. **Output**: Array of starting indices
4. **Constraint**: Anagrams are permutations (same letters, different order)

#### Phase 2: Key Insight Recognition
- **"Fixed window size"** → Window always has length len(p)
- **"Character frequency"** → Need exact letter counts, not just presence
- **"Sliding efficiency"** → Can update counts in O(1) when sliding
- **"Array optimization"** → Fixed 26-size array for lowercase letters

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use a sliding window of size len(p).
I'll count the frequency of each letter in p.
Then I'll slide this window through s, maintaining letter counts.
For each position, I'll compare the window counts with pattern counts.
If they match, I found an anagram!
I can update counts efficiently by removing the leftmost letter
and adding the new rightmost letter as I slide."
```

#### Phase 4: Edge Case Handling
- **Pattern longer than string**: Return empty array
- **Empty strings**: Return empty array
- **Single character patterns**: Handle efficiently
- **No matches**: Return empty array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: s="cbaebabacd", p="abc"

Human thinking:
"Pattern 'abc' has: a:1, b:1, c:1
Window size is 3. Let me slide through s:

Position 0-2: "cba"
   Window has: c:1, b:1, a:1
   Matches pattern! Add index 0 to result

Slide to position 1-3: "bae" 
   Window has: b:1, a:1, e:1
   Doesn't match (e instead of c)

Slide to position 2-4: "aeb"
   Window has: a:1, e:1, b:1
   Doesn't match (e instead of c)

Slide to position 3-5: "eba"
   Window has: e:1, b:1, a:1
   Doesn't match (e instead of c)

Slide to position 4-6: "bab"
   Window has: b:2, a:1
   Doesn't match (b:2 instead of b:1, no c)

Slide to position 5-7: "aba"
   Window has: a:2, b:1
   Doesn't match (a:2 instead of a:1, no c)

Slide to position 6-8: "bac"
   Window has: b:1, a:1, c:1
   Matches pattern! Add index 6 to result

Final result: [0, 6]"
```

#### Phase 6: Intuition Validation
- **Why fixed window works**: Anagrams must have same length as pattern
- **Why frequency arrays work**: Need exact character counts
- **Why sliding update works**: Only two characters change per slide
- **Why O(N) time**: Each character processed once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use hash map?"** → Array is faster for fixed 26-character set
2. **"Should I recompute counts each time?"** → No, update incrementally
3. **"What about uppercase letters?"** → Need to handle or clarify constraints
4. **"Can I use string comparison?"** → No, order doesn't matter in anagrams

### Real-World Analogy
**Like finding all scrambled versions of a word in a long text:**
- You have a book (string) and a target word (pattern)
- You want to find all sections where the letters match your word exactly
- The letters can be in any order, but must be the same collection
- Slide a window the size of your word through the book
- At each position, check if the letters match (regardless of order)
- Mark all positions where you find a match

### Human-Readable Pseudocode
```
function findAnagrams(text, pattern):
    if length(pattern) > length(text):
        return []
    
    patternCount = array[26] initialized to 0
    windowCount = array[26] initialized to 0
    result = []
    
    # Count characters in pattern and first window
    for i from 0 to length(pattern)-1:
        patternCount[pattern[i]-'a']++
        windowCount[text[i]-'a']++
    
    # Check first window
    if patternCount == windowCount:
        result.append(0)
    
    # Slide window through text
    for i from length(pattern) to length(text)-1:
        # Remove leftmost character
        windowCount[text[i-length(pattern)]-'a']--
        # Add new character
        windowCount[text[i]-'a']++
        
        # Check if current window matches
        if patternCount == windowCount:
            result.append(i - length(pattern) + 1)
    
    return result
```

### Execution Visualization

### Example: s="cbaebabacd", p="abc"
```
Pattern counts: a:1, b:1, c:1

Initial: window="cba", counts={c:1,b:1,a:1}
→ Match! result=[0]

Slide 1: remove 'c', add 'e'
→ window="bae", counts={b:1,a:1,e:1}
→ No match

Slide 2: remove 'b', add 'b'
→ window="aeb", counts={a:1,e:1,b:1}
→ No match

Slide 3: remove 'a', add 'a'
→ window="eba", counts={e:1,b:1,a:1}
→ No match

Slide 4: remove 'e', add 'b'
→ window="bab", counts={b:2,a:1}
→ No match

Slide 5: remove 'b', add 'a'
→ window="aba", counts={a:2,b:1}
→ No match

Slide 6: remove 'a', add 'c'
→ window="bac", counts={b:1,a:1,c:1}
→ Match! result=[0,6]

Final result: [0,6]
```

### Key Visualization Points:
- **Fixed window size**: Always len(pattern) characters
- **Frequency arrays**: Track counts for 26 lowercase letters
- **Sliding update**: Remove left, add right in O(1)
- **Comparison**: Check if arrays are identical

### Memory Layout Visualization:
```
String: c b a e b a b a c d
Index:  0 1 2 3 4 5 6 7 8 9
        [window size=3]
        ^ ^ ^
        0 1 2  -> "cba" (match)
          ^ ^ ^
          1 2 3  -> "bae" (no match)
```

### Time Complexity Breakdown:
- **Initial setup**: O(K) where K = len(pattern)
- **Sliding phase**: O(N-K) slides, each O(1) update + O(26) comparison
- **Total**: O(N) where N = len(s)
- **Space**: O(1) - fixed 26-size arrays

### Alternative Approaches:

#### 1. Hash Map Approach (O(N) time, O(K) space)
```go
func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {
        return []int{}
    }
    
    pCount := make(map[byte]int)
    for _, c := range p {
        pCount[c]++
    }
    
    result := []int{}
    for i := 0; i <= len(s)-len(p); i++ {
        window := s[i:i+len(p)]
        wCount := make(map[byte]int)
        for _, c := range window {
            wCount[c]++
        }
        
        if mapsEqual(pCount, wCount) {
            result = append(result, i)
        }
    }
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(N*K) time, creates new map each window

#### 2. Optimized with Mismatch Counter (O(N) time, O(1) space)
```go
func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {
        return []int{}
    }
    
    pCount := [26]int{}
    sCount := [26]int{}
    result := []int{}
    
    // Initialize counts
    for i := 0; i < len(p); i++ {
        pCount[p[i]-'a']++
        sCount[s[i]-'a']++
    }
    
    // Count mismatches
    mismatches := 0
    for i := 0; i < 26; i++ {
        if pCount[i] != sCount[i] {
            mismatches++
        }
    }
    
    if mismatches == 0 {
        result = append(result, 0)
    }
    
    // Slide window
    for i := len(p); i < len(s); i++ {
        left := i - len(p)
        right := i
        
        // Update mismatches for removed character
        if sCount[s[left]-'a'] == pCount[s[left]-'a'] {
            mismatches++
        }
        sCount[s[left]-'a']--
        if sCount[s[left]-'a'] == pCount[s[left]-'a'] {
            mismatches--
        }
        
        // Update mismatches for added character
        if sCount[s[right]-'a'] == pCount[s[right]-'a'] {
            mismatches++
        }
        sCount[s[right]-'a']++
        if sCount[s[right]-'a'] == pCount[s[right]-'a'] {
            mismatches--
        }
        
        if mismatches == 0 {
            result = append(result, left+1)
        }
    }
    
    return result
}
```
- **Pros**: Faster comparison, avoids full array comparison
- **Cons**: More complex logic

### Extensions for Interviews:
- **Unicode Support**: Handle larger character sets
- **Case Insensitive**: Convert to lowercase first
- **Multiple Patterns**: Process multiple patterns efficiently
- **Return Substrings**: Return actual anagrams instead of indices
*/
func main() {
	// Test cases
	testCases := []struct {
		s string
		p string
	}{
		{"cbaebabacd", "abc"},
		{"abab", "ab"},
		{"aaaaaaaaaa", "aaaa"},
		{"abacbabc", "abc"},
		{"", "a"},
		{"a", ""},
		{"abc", "def"},
		{"abababab", "ab"},
	}
	
	for i, tc := range testCases {
		result := findAnagrams(tc.s, tc.p)
		fmt.Printf("Test Case %d: s=\"%s\", p=\"%s\" -> Anagram indices: %v\n", 
			i+1, tc.s, tc.p, result)
	}
}
