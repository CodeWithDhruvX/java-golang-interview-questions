package main

import (
	"fmt"
	"unicode"
)

// 125. Valid Palindrome
// Time: O(N), Space: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	
	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		
		// Compare characters (case-insensitive)
		if left < right && unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		
		left++
		right--
	}
	
	return true
}

// Helper function to check if character is alphanumeric
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers with Character Filtering
- **Left Pointer**: Starts at beginning, moves rightward
- **Right Pointer**: Starts at end, moves leftward
- **Character Filtering**: Skip non-alphanumeric characters
- **Case Insensitive**: Compare characters in lowercase

## 2. PROBLEM CHARACTERISTICS
- **Palindrome Check**: String reads same forwards and backwards
- **Character Filtering**: Ignore non-alphanumeric characters
- **Case Insensitivity**: Treat uppercase and lowercase as same
- **In-place Comparison**: No need to create cleaned string

## 3. SIMILAR PROBLEMS
- Valid Palindrome II (LeetCode 680) - can delete one character
- Reverse String (LeetCode 344)
- Reverse Vowels of a String (LeetCode 345)
- Longest Palindromic Substring (LeetCode 5)

## 4. KEY OBSERVATIONS
- **Two-pointer efficiency**: Compare characters from both ends simultaneously
- **Filtering on the fly**: No need to create separate cleaned string
- **Early termination**: Stop as soon as mismatch found
- **Character classification**: Need to identify alphanumeric characters

## 5. VARIATIONS & EXTENSIONS
- **Allow one deletion**: Can remove at most one character to make palindrome
- **Unicode support**: Handle international characters
- **Sentence palindromes**: Consider word boundaries
- **Numeric palindromes**: Only consider digits

## 6. INTERVIEW INSIGHTS
- Always clarify: "What characters should be considered? Case sensitivity?"
- Edge cases: empty string, single character, all non-alphanumeric
- Space complexity: O(1) - no extra space needed
- Time complexity: O(N) - single pass with two pointers

## 7. COMMON MISTAKES
- Not handling non-alphanumeric characters correctly
- Forgetting case insensitivity
- Creating extra string instead of in-place comparison
- Not handling empty string properly
- Using regex unnecessarily (can be slower)

## 8. OPTIMIZATION STRATEGIES
- **In-place comparison**: No extra string creation
- **Character classification**: Use ASCII ranges for speed
- **Early termination**: Return false on first mismatch
- **Unicode consideration**: Use proper character classification

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like checking if a sentence reads the same backwards:**
- You have a sentence with punctuation and mixed case
- You want to know if it reads the same forwards and backwards
- Ignore punctuation and spaces, treat 'A' and 'a' as same
- Start from both ends and move inward, comparing characters
- Stop as soon as you find a mismatch

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String with various characters
2. **Goal**: Check if it's a palindrome (ignoring non-alphanumeric, case-insensitive)
3. **Output**: True if palindrome, false otherwise
4. **Constraints**: Must be efficient, O(1) extra space

#### Phase 2: Key Insight Recognition
- **"Two-pointer approach"** → Compare from both ends simultaneously
- **"Filter on the fly"** → Skip non-alphanumeric during comparison
- **"Case normalization"** → Convert to lowercase for comparison
- **"Early exit"** → Stop at first mismatch

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use two pointers, one at the start and one at the end.
I'll move them towards each other, comparing characters.
But first, I need to skip any non-alphanumeric characters.
When I find alphanumeric characters, I'll compare them in lowercase.
If they don't match, it's not a palindrome.
If they do match, I'll move both pointers inward.
If I reach the middle without mismatches, it's a palindrome."
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return true (empty string is palindrome)
- **Single character**: Return true (single character is palindrome)
- **All non-alphanumeric**: Return true (empty after filtering)
- **Mixed case**: Handle case insensitivity properly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: "A man, a plan, a canal: Panama"

Human thinking:
"I'll compare from both ends, ignoring punctuation and case:

Left=0('A'), Right=30('a'):
   Both are letters, compare 'a' vs 'a' → match
   Move left to find next letter, right to find previous letter

Left moves to 'm', Right moves to 'm':
   Compare 'm' vs 'm' → match
   Move both pointers

Left moves to 'a', Right moves to 'a':
   Compare 'a' vs 'a' → match
   Move both pointers

Continue this process...
Eventually pointers meet or cross without mismatches.
It's a palindrome!"
```

#### Phase 6: Intuition Validation
- **Why two pointers work**: Palindrome property is symmetric
- **Why filtering on fly works**: No need to create cleaned string
- **Why case insensitivity works**: Palindrome definition ignores case
- **Why O(1) space**: Only need two pointers, no extra storage

### Common Human Pitfalls & How to Avoid Them
1. **"Should I create a cleaned string?"** → No, can filter on the fly
2. **"What about Unicode characters?"** → Use proper character classification
3. **"How to check alphanumeric?"** → Use ASCII ranges or built-in functions
4. **"Should I use regex?"** → Possible but slower than manual filtering

### Real-World Analogy
**Like checking if a book title reads the same backwards:**
- You have a title with punctuation and capitalization
- You want to know if it's a palindrome phrase
- Ignore punctuation and spaces, treat all letters as lowercase
- Start from first and last letters, compare them
- Move inward, checking each pair of letters
- If all pairs match, it's a palindrome!

### Human-Readable Pseudocode
```
function isPalindrome(text):
    left = 0
    right = length(text) - 1
    
    while left < right:
        // Skip non-alphanumeric from left
        while left < right and not isAlphanumeric(text[left]):
            left = left + 1
            
        // Skip non-alphanumeric from right
        while left < right and not isAlphanumeric(text[right]):
            right = right - 1
            
        // Compare characters (case-insensitive)
        if left < right and lowercase(text[left]) != lowercase(text[right]):
            return false
            
        left = left + 1
        right = right - 1
    
    return true

function isAlphanumeric(char):
    return (char >= 'a' and char <= 'z') or 
           (char >= 'A' and char <= 'Z') or 
           (char >= '0' and char <= '9')
```

### Execution Visualization

### Example: "A man, a plan, a canal: Panama"
```
String: A man, a plan, a canal: Panama
Index:  0123456789012345678901234567890

Initial: left=0, right=30

Step 1: left=0('A'), right=30('a')
→ Both alphanumeric, compare 'a' vs 'a' → match
→ left=1, right=29

Step 2: left=1(' ') → skip, left=2
         right=29('m') → compare when left finds 'm'
→ left=2('m'), right=29('m')
→ Compare 'm' vs 'm' → match
→ left=3, right=28

Step 3: left=3('a'), right=28('a')
→ Compare 'a' vs 'a' → match
→ left=4, right=27

Continue this process...
Final: All pairs match, return true
```

### Key Visualization Points:
- **Pointer movement**: left moves right, right moves left
- **Character filtering**: Skip non-alphanumeric characters
- **Case normalization**: Compare in lowercase
- **Early termination**: Return false on first mismatch

### Memory Layout Visualization:
```
String: "A man, a plan, a canal: Panama"
Index:   0    5    10   15   20   25   30
         ^                              ^
       left=0                        right=30
         'A' vs 'a' (match)
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each character visited at most once
- **Constant Space**: O(1) - only two pointers
- **Character Classification**: O(1) per character using ASCII ranges
- **Early Termination**: Often returns before processing entire string

### Alternative Approaches:

#### 1. Clean String Approach (O(N) time, O(N) space)
```go
func isPalindrome(s string) bool {
    // Clean the string
    cleaned := ""
    for _, c := range s {
        if isAlphanumeric(byte(c)) {
            cleaned += string(unicode.ToLower(c))
        }
    }
    
    // Check palindrome
    left, right := 0, len(cleaned)-1
    for left < right {
        if cleaned[left] != cleaned[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```
- **Pros**: Simpler logic
- **Cons**: Uses O(N) extra space

#### 2. Regex Approach (O(N) time, O(N) space)
```go
import "regexp"

func isPalindrome(s string) bool {
    // Remove non-alphanumeric and convert to lowercase
    re := regexp.MustCompile(`[^a-zA-Z0-9]`)
    cleaned := strings.ToLower(re.ReplaceAllString(s, ""))
    
    // Check palindrome
    left, right := 0, len(cleaned)-1
    for left < right {
        if cleaned[left] != cleaned[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```
- **Pros**: Concise code
- **Cons**: Regex overhead, extra space

### Extensions for Interviews:
- **Valid Palindrome II**: Can delete at most one character
- **Unicode Support**: Handle international characters properly
- **Performance Optimization**: Use lookup table for character classification
- **Streaming Input**: Handle very long inputs without storing full string
*/
func main() {
	// Test cases
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
		"",
		"madam",
		"Able was I ere I saw Elba",
		"No lemon, no melon",
		"12321",
		"12345",
		".,",
	}
	
	for i, s := range testCases {
		result := isPalindrome(s)
		fmt.Printf("Test Case %d: \"%s\" -> Palindrome: %t\n", i+1, s, result)
	}
}
