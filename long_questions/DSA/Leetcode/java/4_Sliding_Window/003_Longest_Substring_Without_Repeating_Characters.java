import java.util.HashMap;
import java.util.Map;

public class LongestSubstringWithoutRepeatingCharacters {
    
    // 3. Longest Substring Without Repeating Characters (Variable Size Sliding Window)
    // Time: O(N), Space: O(min(N, M)) where M is the size of character set
    public static int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> charMap = new HashMap<>();
        int left = 0;
        int maxLength = 0;
        
        for (int right = 0; right < s.length(); right++) {
            char currentChar = s.charAt(right);
            
            // If character is already in the window, move left pointer
            if (charMap.containsKey(currentChar) && charMap.get(currentChar) >= left) {
                left = charMap.get(currentChar) + 1;
            }
            
            // Update the character's last seen position
            charMap.put(currentChar, right);
            
            // Update max length
            int currentLength = right - left + 1;
            maxLength = Math.max(maxLength, currentLength);
        }
        
        return maxLength;
    }

    public static void main(String[] args) {
        // Test cases
        String[] testCases = {
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
            "abccba"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = lengthOfLongestSubstring(testCases[i]);
            System.out.printf("Test Case %d: \"%s\" -> Length of longest substring: %d\n", 
                i + 1, testCases[i], result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Variable Size Sliding Window
- **Window Expansion**: Expand right pointer to include new elements
- **Window Contraction**: Contract left pointer when condition violated
- **Hash Map Tracking**: Track character positions efficiently
- **Longest Valid Window**: Maintain maximum window size

## 2. PROBLEM CHARACTERISTICS
- **Longest Substring**: Find longest substring without repeating characters
- **Unique Characters**: All characters in window must be distinct
- **Variable Window Size**: Window size changes based on conditions
- **Character Tracking**: Need to track last seen positions

## 3. SIMILAR PROBLEMS
- Longest Substring with At Most K Distinct Characters
- Longest Substring with At Most K Repeating Characters
- Find All Anagrams in a String
- Permutation in String

## 4. KEY OBSERVATIONS
- Sliding window maintains current valid substring
- Hash map stores last seen position of each character
- When duplicate found, move left pointer past previous occurrence
- Time complexity: O(N) for N characters
- Space complexity: O(min(N, M)) where M is character set size

## 5. VARIATIONS & EXTENSIONS
- Different character sets (ASCII, Unicode)
- Case sensitivity variations
- Maximum window constraints
- Multiple substring problems

## 6. INTERVIEW INSIGHTS
- Clarify: "What is the character set?"
- Edge cases: empty string, single character, all unique, all same
- Time complexity: O(N) vs O(N²) brute force
- Space complexity: O(min(N, M)) vs O(N) naive

## 7. COMMON MISTAKES
- Incorrect left pointer movement
- Not updating character positions properly
- Wrong window size calculation
- Missing edge cases
- Inefficient character tracking

## 8. OPTIMIZATION STRATEGIES
- Use HashMap for O(1) character lookups
- Efficient left pointer updates
- Single pass through string
- Minimal space usage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest unique sequence:**
- You have a string of characters
- Need to find longest sequence with no repeats
- Sliding window represents current valid sequence
- When duplicate appears, shrink window to remove previous occurrence
- This is like finding the longest stretch of unique items

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String of characters
2. **Goal**: Find longest substring without repeating characters
3. **Output**: Length of longest valid substring

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N²) to check all substrings
- **"How to optimize?"** → Use sliding window with hash map
- **"Why sliding window?"** → Maintains current valid substring
- **"Why hash map?"** → Efficient character position tracking

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use sliding window:
1. Initialize left pointer at start, empty hash map
2. Expand window by moving right pointer
3. For each character:
   - If already in window, move left pointer past previous occurrence
   - Update character's last seen position
   - Update maximum window length
4. Continue until end of string
5. Return maximum length found"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return 0
- **Single character**: Return 1
- **All unique**: Return string length
- **All same**: Return 1

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
String: "abcabcbb"

Human thinking:
"Let's apply sliding window:

Initialize: left=0, maxLength=0, charMap={}

Step 1: right=0, char='a'
- 'a' not in charMap
- charMap={'a':0}
- window=[0,0], length=1, maxLength=1

Step 2: right=1, char='b'
- 'b' not in charMap
- charMap={'a':0, 'b':1}
- window=[0,1], length=2, maxLength=2

Step 3: right=2, char='c'
- 'c' not in charMap
- charMap={'a':0, 'b':1, 'c':2}
- window=[0,2], length=3, maxLength=3

Step 4: right=3, char='a'
- 'a' in charMap and charMap['a']=0 >= left=0
- Move left to charMap['a']+1 = 1
- charMap={'a':3, 'b':1, 'c':2}
- window=[1,3], length=3, maxLength=3

Step 5: right=4, char='b'
- 'b' in charMap and charMap['b']=1 >= left=1
- Move left to charMap['b']+1 = 2
- charMap={'a':3, 'b':4, 'c':2}
- window=[2,4], length=3, maxLength=3

Continue...

Final result: maxLength=3 ✓

Manual verification:
Longest substrings: "abc", "bca", "cab" (all length 3) ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Sliding window always contains unique characters
- **Why it's efficient**: O(N) vs O(N²) brute force
- **Why it's correct**: Considers all possible substrings

### Common Human Pitfalls & How to Avoid Them
1. **"Why not check all substrings?"** → O(N²) too slow
2. **"What about left pointer?"** → Move past previous occurrence
3. **"How to track characters?"** → Use hash map for positions
4. **"What about window size?"** → right-left+1

### Real-World Analogy
**Like finding the longest sequence of unique items:**
- You have a conveyor belt of items (string)
- Need to find longest stretch with no duplicates
- Sliding window represents current unique sequence
- When duplicate appears, remove items until duplicate is gone
- This is used in data deduplication, pattern matching
- Like finding the longest unique sequence in a stream

### Human-Readable Pseudocode
```
function lengthOfLongestSubstring(s):
    charMap = empty hash map
    left = 0
    maxLength = 0
    
    for right from 0 to s.length-1:
        currentChar = s[right]
        
        // If character is already in window, move left pointer
        if currentChar in charMap and charMap[currentChar] >= left:
            left = charMap[currentChar] + 1
        
        // Update character's last seen position
        charMap[currentChar] = right
        
        // Update maximum window length
        currentLength = right - left + 1
        maxLength = max(maxLength, currentLength)
    
    return maxLength
```

### Execution Visualization

### Example: s="abcabcbb"
```
Sliding Window Process:

Initialize: left=0, maxLength=0, charMap={}

Step 1: right=0, char='a'
- charMap={'a':0}
- window=[0,0], length=1, maxLength=1

Step 2: right=1, char='b'
- charMap={'a':0, 'b':1}
- window=[0,1], length=2, maxLength=2

Step 3: right=2, char='c'
- charMap={'a':0, 'b':1, 'c':2}
- window=[0,2], length=3, maxLength=3

Step 4: right=3, char='a'
- 'a' already in window at position 0
- Move left to 1
- charMap={'a':3, 'b':1, 'c':2}
- window=[1,3], length=3, maxLength=3

Step 5: right=4, char='b'
- 'b' already in window at position 1
- Move left to 2
- charMap={'a':3, 'b':4, 'c':2}
- window=[2,4], length=3, maxLength=3

Continue...

Final result: maxLength=3 ✓

Visualization:
Sliding window maintains unique characters
Left pointer moves past duplicates
Maximum length tracked throughout ✓
```

### Key Visualization Points:
- **Window Expansion**: Right pointer moves forward
- **Window Contraction**: Left pointer moves past duplicates
- **Character Tracking**: Hash map stores last positions
- **Maximum Length**: Updated when window expands

### Memory Layout Visualization:
```
String: "a b c a b c b b"
Index:  0 1 2 3 4 5 6 7

Window Evolution:
Step 1: [a] (0,0) length=1
Step 2: [a b] (0,1) length=2
Step 3: [a b c] (0,2) length=3
Step 4: [b c a] (1,3) length=3
Step 5: [c a b] (2,4) length=3
Step 6: [a b c] (3,5) length=3
Step 7: [b c] (4,6) length=2
Step 8: [c] (5,7) length=1

Maximum length: 3 ✓

Sliding window efficiently tracks unique characters
Hash map enables O(1) duplicate detection
```

### Time Complexity Breakdown:
- **Sliding Window**: O(N) time, O(min(N, M)) space
- **Brute Force**: O(N²) time, O(1) space
- **Optimal**: Best possible for this problem
- **vs Naive**: O(N) vs O(N²) significant improvement
*/
