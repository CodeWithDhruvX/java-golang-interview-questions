package main

import "fmt"

// 3. Longest Substring Without Repeating Characters (Variable Size Sliding Window)
// Time: O(N), Space: O(min(N, M)) where M is the size of character set
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// If character is already in the window, move left pointer
		if index, exists := charMap[s[right]]; exists && index >= left {
			left = index + 1
		}
		
		// Update the character's last seen position
		charMap[s[right]] = right
		
		// Update max length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	
	return maxLength
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Variable Size Sliding Window with Hash Map
- **Left Pointer**: Window start, moves right when duplicate found
- **Right Pointer**: Window end, expands window by one each iteration
- **Hash Map**: Stores last seen position of each character
- **Window Adjustment**: Move left to max(current left, duplicate position + 1)

## 2. PROBLEM CHARACTERISTICS
- **Unique Characters**: Substring must have no repeating characters
- **Variable Window**: Window size changes based on duplicates
- **Maximum Length**: Find longest valid window
- **Character Tracking**: Need to know when characters repeat

## 3. SIMILAR PROBLEMS
- Longest Substring with At Most K Distinct Characters
- Longest Substring with At Most 2 Distinct Characters
- Find All Anagrams in a String (fixed window)
- Minimum Window Substring (variable window with constraints)

## 4. KEY OBSERVATIONS
- **Duplicate Detection**: When character repeats, must shrink window
- **Position Tracking**: Need last seen position of each character
- **Window Expansion**: Always expand right pointer
- **Window Contraction**: Only when duplicate found within current window

## 5. VARIATIONS & EXTENSIONS
- **At Most K Distinct**: Allow up to K different characters
- **Fixed Character Set**: Only allow specific characters
- **Unicode Support**: Handle international characters
- **Streaming Input**: Process characters as they arrive

## 6. INTERVIEW INSIGHTS
- Always clarify: "What character set? ASCII or Unicode?"
- Edge cases: empty string, single character, all unique, all same
- Space complexity: O(min(N, M)) where M is character set size
- Time complexity: O(N) - each character visited at most twice

## 7. COMMON MISTAKES
- Not updating left pointer correctly (should use max)
- Forgetting to check if duplicate is within current window
- Using nested loops instead of sliding window (O(N²))
- Not handling empty string properly
- Confusing character indices with window boundaries

## 8. OPTIMIZATION STRATEGIES
- **Array instead of hash map**: For ASCII, use fixed-size array
- **Early termination**: If remaining characters can't beat current max
- **Character set optimization**: Use smaller data types if possible
- **Cache-friendly access**: Sequential memory access pattern

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest unique sequence in a line of people:**
- You have people wearing colored shirts (characters)
- You want the longest sequence where no one repeats a color
- Start from the left, keep adding people to your sequence
- If someone with a repeated color appears, remove people from the left
- Keep track of the longest unique sequence you've seen
- Remember where each color last appeared to know how far to move

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String of characters
2. **Goal**: Find longest substring with all unique characters
3. **Output**: Length of longest such substring
4. **Constraint**: Substring must be contiguous

#### Phase 2: Key Insight Recognition
- **"Window concept"** → Maintain a window of unique characters
- **"Duplicate handling"** → When duplicate found, shrink from left
- **"Position memory"** → Remember last seen position of each character
- **"Greedy expansion"** → Always try to expand window rightward

#### Phase 3: Strategy Development
```
Human thought process:
"I'll maintain a window [left, right] of unique characters.
As I move right, if I encounter a character already in my window,
I need to move left to exclude the previous occurrence.
I'll use a hash map to remember where each character was last seen.
The window size is right - left + 1, and I'll track the maximum."
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return 0 (no substring)
- **Single character**: Return 1 (single character is unique)
- **All unique characters**: Return string length
- **All same character**: Return 1 (only one unique character)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: "abcabcbb"

Human thinking:
"I'll scan through the string, maintaining my unique window:

Position 0: 'a' - Window: [a], Length: 1, Max: 1
Position 1: 'b' - Window: [ab], Length: 2, Max: 2
Position 2: 'c' - Window: [abc], Length: 3, Max: 3
Position 3: 'a' - 'a' already in window at position 0
           Move left to position 1, Window: [bca], Length: 3, Max: 3
Position 4: 'b' - 'b' already in window at position 1
           Move left to position 2, Window: [cab], Length: 3, Max: 3
Position 5: 'c' - 'c' already in window at position 2
           Move left to position 3, Window: [abc], Length: 3, Max: 3
Position 6: 'b' - 'b' already in window at position 4
           Move left to position 5, Window: [cb], Length: 2, Max: 3
Position 7: 'b' - 'b' already in window at position 7
           Move left to position 8, Window: [b], Length: 1, Max: 3

Final answer: 3"
```

#### Phase 6: Intuition Validation
- **Why sliding window works**: We only need contiguous substrings
- **Why hash map works**: Quick lookup of character positions
- **Why O(N) time**: Each character added and removed at most once
- **Why max() for left**: Ensures we never move left backward

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all substrings?"** → That's O(N²), sliding window is O(N)
2. **"Should I move left to duplicate + 1?"** → Only if duplicate is within current window
3. **"What about multiple duplicates?"** → Hash map handles any character
4. **"Can I use array instead of hash map?"** → Yes, for limited character sets

### Real-World Analogy
**Like finding the longest sequence of unique cards in a deck:**
- You're dealing cards face up in a line
- You want the longest sequence with no repeated values
- As you deal new cards, add them to your current sequence
- If a duplicate appears, remove cards from the left until duplicate is gone
- Keep track of the longest unique sequence you've seen
- Remember where each value last appeared

### Human-Readable Pseudocode
```
function longestUniqueSubstring(text):
    lastSeen = empty map  // character -> last position
    left = 0
    maxLength = 0
    
    for right from 0 to length(text)-1:
        currentChar = text[right]
        
        // If character seen before and is within current window
        if currentChar in lastSeen and lastSeen[currentChar] >= left:
            left = lastSeen[currentChar] + 1
        
        lastSeen[currentChar] = right
        currentLength = right - left + 1
        maxLength = max(maxLength, currentLength)
    
    return maxLength
```

### Execution Visualization

### Example: "abcabcbb"
```
String: a b c a b c b b
Index:  0 1 2 3 4 5 6 7

Initial: left=0, maxLength=0, charMap={}

Step 1: right=0, char='a'
→ 'a' not in charMap
→ charMap={'a':0}, window=[0,0], length=1, max=1

Step 2: right=1, char='b'
→ 'b' not in charMap
→ charMap={'a':0,'b':1}, window=[0,1], length=2, max=2

Step 3: right=2, char='c'
→ 'c' not in charMap
→ charMap={'a':0,'b':1,'c':2}, window=[0,2], length=3, max=3

Step 4: right=3, char='a'
→ 'a' in charMap at position 0 (>= left=0)
→ left = 0 + 1 = 1
→ charMap={'a':3,'b':1,'c':2}, window=[1,3], length=3, max=3

Step 5: right=4, char='b'
→ 'b' in charMap at position 1 (>= left=1)
→ left = 1 + 1 = 2
→ charMap={'a':3,'b':4,'c':2}, window=[2,4], length=3, max=3

Continue...
Final: maxLength = 3
```

### Key Visualization Points:
- **Window boundaries**: left and right pointers define current substring
- **Duplicate detection**: Check if char exists and position >= left
- **Left movement**: Only moves right, never backward
- **Max tracking**: Update after each character

### Memory Layout Visualization:
```
String: [a][b][c][a][b][c][b][b]
Index:   0  1  2  3  4  5  6  7
        ^           ^
     left=0      right=3
        'a' repeats at 3, move left to 1
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each character visited at most twice
- **Hash Map Operations**: O(1) average for insert and lookup
- **Space**: O(min(N, M)) where M is character set size
- **Window Operations**: Each character enters and leaves window once

### Alternative Approaches:

#### 1. Brute Force (O(N²))
```go
func lengthOfLongestSubstring(s string) int {
    maxLength := 0
    for i := 0; i < len(s); i++ {
        seen := make(map[byte]bool)
        for j := i; j < len(s); j++ {
            if seen[s[j]] {
                break
            }
            seen[s[j]] = true
            maxLength = max(maxLength, j-i+1)
        }
    }
    return maxLength
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time complexity

#### 2. Array for ASCII (O(N) time, O(1) space)
```go
func lengthOfLongestSubstring(s string) int {
    lastSeen := [256]int{-1} // Initialize with -1
    left := 0
    maxLength := 0
    
    for right := 0; right < len(s); right++ {
        if lastSeen[s[right]] >= left {
            left = lastSeen[s[right]] + 1
        }
        lastSeen[s[right]] = right
        maxLength = max(maxLength, right-left+1)
    }
    return maxLength
}
```
- **Pros**: O(1) space for ASCII, faster than hash map
- **Cons**: Limited to ASCII, not suitable for Unicode

### Extensions for Interviews:
- **At Most K Distinct**: Allow up to K different characters
- **Unicode Support**: Handle international characters properly
- **Streaming Input**: Process characters as they arrive
- **Return Substring**: Return actual substring instead of just length
*/
func main() {
	// Test cases
	testCases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"a",
		"au",
		"dvdf",
		"abba",
		"tmmzuxt",
		"abcdefghijklmnopqrstuvwxyz",
		"abccba",
	}
	
	for i, s := range testCases {
		result := lengthOfLongestSubstring(s)
		fmt.Printf("Test Case %d: \"%s\" -> Length of longest substring: %d\n", 
			i+1, s, result)
	}
}
