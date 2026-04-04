package main

import "fmt"

// 76. Minimum Window Substring (Variable Size Sliding Window)
// Time: O(N + M), Space: O(1) for ASCII characters
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	
	// Frequency map for characters in t
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}
	
	// Number of unique characters in t that need to be present in the window
	required := len(tCount)
	
	// Left and right pointers
	left, right := 0, 0
	
	// Keep track of how many unique characters in the current window match the required count
	formed := 0
	
	// Frequency map for the current window
	windowCount := make(map[byte]int)
	
	// Result variables: (window length, left, right)
	result := []int{-1, 0, 0}
	
	for right < len(s) {
		char := s[right]
		windowCount[char]++
		
		// If the frequency of the current character matches exactly the required frequency in t
		if count, exists := tCount[char]; exists && windowCount[char] == count {
			formed++
		}
		
		// Try to contract the window till the point it ceases to be 'desirable'
		for left <= right && formed == required {
			// Save the smallest window
			windowSize := right - left + 1
			if result[0] == -1 || windowSize < result[0] {
				result[0] = windowSize
				result[1] = left
				result[2] = right
			}
			
			// The character at the position left is no longer a part of the window
			leftChar := s[left]
			windowCount[leftChar]--
			if count, exists := tCount[leftChar]; exists && windowCount[leftChar] < count {
				formed--
			}
			
			left++
		}
		
		right++
	}
	
	if result[0] == -1 {
		return ""
	}
	
	return s[result[1] : result[2]+1]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Variable Size Sliding Window with Frequency Maps
- **Target Frequency Map**: Count required characters from string t
- **Window Frequency Map**: Count characters in current window
- **Formed Counter**: Track how many required characters are satisfied
- **Window Contraction**: Shrink window when all requirements met

## 2. PROBLEM CHARACTERISTICS
- **Substring Containment**: Window must contain all characters from t
- **Frequency Matching**: Must match exact counts of each character
- **Minimum Length**: Find smallest valid window
- **Order Preservation**: Characters must maintain original order in s

## 3. SIMILAR PROBLEMS
- Find All Anagrams in a String (fixed window)
- Longest Substring with At Most K Distinct Characters
- Permutation in String (check if permutation exists)
- Smallest Subsequence of Distinct Characters

## 4. KEY OBSERVATIONS
- **Two-phase approach**: Expand window, then contract when valid
- **Frequency matching**: Need exact character counts, not just presence
- **Optimization opportunity**: Contract as much as possible when valid
- **Early termination**: Can stop if remaining characters insufficient

## 5. VARIATIONS & EXTENSIONS
- **At Most K Characters**: Allow up to K different characters
- **Permutation Check**: Check if any permutation of t exists in s
- **Multiple Queries**: Process multiple t strings efficiently
- **Unicode Support**: Handle international character sets

## 6. INTERVIEW INSIGHTS
- Always clarify: "Character set? Can t have duplicates? Case sensitivity?"
- Edge cases: empty strings, t longer than s, no valid window
- Space complexity: O(1) for ASCII, O(M) for Unicode
- Time complexity: O(N + M) where N=|s|, M=|t|

## 7. COMMON MISTAKES
- Not handling character frequencies correctly (just checking presence)
- Forgetting to update formed counter when shrinking window
- Using nested loops instead of sliding window
- Not handling case where no valid window exists
- Confusing character indices with window boundaries

## 8. OPTIMIZATION STRATEGIES
- **Array for ASCII**: Use fixed-size array instead of hash map
- **Early pruning**: Skip impossible starting positions
- **Character filtering**: Pre-filter s to only relevant characters
- **Optimized contraction**: Shrink window more aggressively

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the smallest section of a book that contains all required words:**
- You have a long book (string s) and a shopping list (string t)
- You need the smallest continuous section that contains all items from your list
- Start scanning from the beginning, collecting items in your basket
- When you have everything from your list, try to make your basket smaller
- Remove items from the left as long as you still have everything needed
- Keep track of the smallest section that worked

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s (source) and string t (target characters)
2. **Goal**: Find smallest window in s containing all characters from t
3. **Output**: The actual substring (empty if none exists)
4. **Constraint**: Must maintain order and frequency of characters

#### Phase 2: Key Insight Recognition
- **"Frequency requirement"** → Need exact counts, not just presence
- **"Two-phase window"** → Expand to meet requirements, then contract
- **"Optimization opportunity"** → Always try to shrink when valid
- **"Character tracking"** → Need to know which requirements are satisfied

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the smallest window containing all characters from t.
First, I'll count what characters I need from t.
Then I'll expand my window in s until I have all required characters.
Once I have everything, I'll try to shrink the window from the left
while still maintaining all requirements.
I'll keep track of the smallest valid window I find."
```

#### Phase 4: Edge Case Handling
- **Empty strings**: Return empty if either string is empty
- **t longer than s**: Return empty (impossible to contain t)
- **No valid window**: Return empty if t's characters not all in s
- **Single character**: Handle simple case efficiently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: s="ADOBECODEBANC", t="ABC"

Human thinking:
"I need A, B, C in my window. Let me scan:

Position 0: 'A' - Have A, need B, C
Position 1: 'D' - Still need B, C
Position 2: 'O' - Still need B, C
Position 3: 'B' - Have A, B, need C
Position 4: 'E' - Still need C
Position 5: 'C' - Now have A, B, C! Window [0,5]="ADOBEC"
           Can I shrink? Remove 'A' (position 0) - lose A, can't shrink
           Smallest so far: "ADOBEC"

Continue expanding:
Position 6: 'O' - Still have A, B, C
Position 7: 'D' - Still have A, B, C
Position 8: 'E' - Still have A, B, C
Position 9: 'B' - Still have A, B, C, try shrinking
           Remove 'A' (position 0) - still have A at position 10? No
           Can't shrink past position 0 yet

Position 10: 'A' - Still have A, B, C, try shrinking
           Remove 'A' (position 0) - still have A at 10, good!
           Window [1,10]="DOBECODEBA", remove 'D' at 1
           Window [2,10]="OBECODEBA", remove 'O' at 2
           Window [3,10]="BECODEBA", remove 'B' at 3 - still have B at 9
           Window [4,10]="ECODEBA", remove 'E' at 4
           Window [5,10]="CODEBA", remove 'C' at 5 - still have C at 5
           Window [6,10]="ODEBA", remove 'O' at 6
           Window [7,10]="DEBA", remove 'D' at 7
           Window [8,10]="EBA", remove 'E' at 8
           Window [9,10]="BA" - lose A, can't shrink further
           New smallest: "CODEBA"

Continue to find "BANC" as final answer...
```

#### Phase 6: Intuition Validation
- **Why sliding window works**: We need contiguous substring
- **Why frequency maps work**: Need exact character counts
- **Why contraction works**: Always try to minimize when valid
- **Why O(N + M)**: Each character visited limited times

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just check presence?"** → Need exact frequencies, not just presence
2. **"Should I shrink after each addition?"** → Only when window is valid
3. **"What about multiple same characters?"** → Frequency maps handle this
4. **"Can I use two pointers only?"** → Need frequency maps for character counts

### Real-World Analogy
**Like finding the smallest section of a recipe book that contains all required ingredients:**
- You have a large recipe book (string s) and a recipe (string t)
- You need the smallest continuous section mentioning all required ingredients
- Start reading from the beginning, highlighting ingredients you find
- When you have all ingredients, try to find a smaller section
- Remove ingredients from the left as long as you still have all required ones
- Keep track of the smallest section that worked

### Human-Readable Pseudocode
```
function minWindow(source, target):
    if source empty or target empty or target longer than source:
        return ""
    
    targetCount = frequency map of target characters
    required = number of unique characters in targetCount
    
    left = 0, right = 0
    formed = 0  // how many required characters are satisfied
    windowCount = empty frequency map
    result = [infinity, 0, 0]  // [length, left, right]
    
    while right < length(source):
        char = source[right]
        windowCount[char]++
        
        // If this character's count now matches required count
        if char in targetCount and windowCount[char] == targetCount[char]:
            formed++
        
        // Try to contract window while it's still valid
        while left <= right and formed == required:
            windowSize = right - left + 1
            if windowSize < result[0]:
                result = [windowSize, left, right]
            
            # Remove character from left
            leftChar = source[left]
            windowCount[leftChar]--
            if leftChar in targetCount and windowCount[leftChar] < targetCount[leftChar]:
                formed--
            left++
        
        right++
    
    if result[0] == infinity:
        return ""
    return source[result[1] : result[2]+1]
```

### Execution Visualization

### Example: s="ADOBECODEBANC", t="ABC"
```
Target counts: A:1, B:1, C:1, required=3

Initial: left=0, right=0, formed=0, windowCount={}

=== EXPANDING PHASE ===
right=0, char='A': windowCount={'A':1}, formed=1
right=1, char='D': windowCount={'A':1,'D':1}, formed=1
right=2, char='O': windowCount={'A':1,'D':1,'O':1}, formed=1
right=3, char='B': windowCount={'A':1,'D':1,'O':1,'B':1}, formed=2
right=4, char='E': windowCount={'A':1,'D':1,'O':1,'B':1,'E':1}, formed=2
right=5, char='C': windowCount={'A':1,'D':1,'O':1,'B':1,'E':1,'C':1}, formed=3

=== CONTRACTION PHASE (formed=3=required) ===
Window [0,5]="ADOBEC", size=6, result=[6,0,5]
Try shrink: remove 'A' at left=0, windowCount['A']=0, formed=2, stop contraction

Continue expanding...
Eventually find "BANC" as optimal solution
```

### Key Visualization Points:
- **Target frequency**: What characters and how many we need
- **Window frequency**: What characters we currently have
- **Formed counter**: How many requirements are satisfied
- **Contraction logic**: Shrink while maintaining validity

### Memory Layout Visualization:
```
Source: A D O B E C O D E B A N C
Index:  0 1 2 3 4 5 6 7 8 9 10 11 12
        ^                 ^
     left=0           right=5
        Window: "ADOBEC" (valid)
```

### Time Complexity Breakdown:
- **Target Processing**: O(M) to build frequency map
- **Window Expansion**: O(N) - each character added once
- **Window Contraction**: O(N) - each character removed at most once
- **Total**: O(N + M) - linear in input sizes
- **Space**: O(K) where K is number of unique characters in t

### Alternative Approaches:

#### 1. Brute Force (O(N²M))
```go
func minWindow(s string, t string) string {
    result := ""
    tCount := make(map[byte]int)
    for _, c := range t {
        tCount[c]++
    }
    
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            window := s[i:j+1]
            if containsAll(window, tCount) {
                if result == "" || len(window) < len(result) {
                    result = window
                }
                break  // smallest window starting at i
            }
        }
    }
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²M) time complexity, very slow

#### 2. Optimized with Array (O(N + M) time, O(1) space for ASCII)
```go
func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 {
        return ""
    }
    
    tCount := [128]int{}
    for _, c := range t {
        tCount[c]++
    }
    
    required := 0
    for _, count := range tCount {
        if count > 0 {
            required++
        }
    }
    
    // Rest of algorithm similar but using arrays
    // ...
}
```
- **Pros**: O(1) space for ASCII, faster operations
- **Cons**: Limited to ASCII character set

### Extensions for Interviews:
- **Permutation in String**: Check if any permutation exists
- **Find All Anagrams**: Return all starting positions
- **At Most K Distinct**: Allow up to K different characters
- **Unicode Support**: Handle international characters
*/
func main() {
	// Test cases
	testCases := []struct {
		s string
		t string
	}{
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
		{"a", "aa"},
		{"ab", "b"},
		{"ab", "c"},
		{"abc", "abc"},
		{"aa", "aa"},
		{"ab", "ab"},
		{"bba", "ab"},
		{"aaaaaaaaaaaabbbbbcdd", "abcdd"},
	}
	
	for i, tc := range testCases {
		result := minWindow(tc.s, tc.t)
		fmt.Printf("Test Case %d: s=\"%s\", t=\"%s\" -> Min window: \"%s\"\n", 
			i+1, tc.s, tc.t, result)
	}
}
