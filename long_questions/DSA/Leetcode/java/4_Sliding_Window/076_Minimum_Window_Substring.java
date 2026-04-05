public class MinimumWindowSubstring {
    
    // 76. Minimum Window Substring
    // Time: O(N + M), Space: O(K) where K is number of unique characters
    public static String minWindow(String s, String t) {
        if (s == null || t == null || s.length() == 0 || t.length() == 0) {
            return "";
        }
        
        // Frequency map for characters in t
        int[] need = new int[128];
        for (char c : t.toCharArray()) {
            need[c]++;
        }
        
        int left = 0, right = 0;
        int minLen = Integer.MAX_VALUE;
        int start = 0;
        int count = t.length();
        
        while (right < s.length()) {
            char c = s.charAt(right);
            
            // If this character is needed
            if (need[c] > 0) {
                count--;
            }
            
            need[c]--;
            right++;
            
            // When we have a valid window
            while (count == 0) {
                // Update minimum window
                if (right - left < minLen) {
                    minLen = right - left;
                    start = left;
                }
                
                // Try to shrink window from left
                char leftChar = s.charAt(left);
                need[leftChar]++;
                if (need[leftChar] > 0) {
                    count++;
                }
                left++;
            }
        }
        
        return minLen == Integer.MAX_VALUE ? "" : s.substring(start, start + minLen);
    }

    public static void main(String[] args) {
        String[][] testCases = {
            {"ADOBECODEBANC", "ABC"},
            {"a", "a"},
            {"a", "aa"},
            {"ab", "b"},
            {"bba", "ab"},
            {"aaaaaaaaaaaabbbbbcddddddeeeee", "abcde"},
            {"", "a"},
            {"a", ""},
            {"abc", "d"},
            {"abacbabc", "abc"},
            {"ab", "ab"},
            {"ab", "ba"},
            {"cabwefgewcwaefgcf", "cae"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String s = testCases[i][0];
            String t = testCases[i][1];
            
            String result = minWindow(s, t);
            System.out.printf("Test Case %d: s=\"%s\", t=\"%s\" -> \"%s\"\n", 
                i + 1, s, t, result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sliding Window with Frequency Counting
- **Variable Window**: Window size changes based on validity
- **Two Pointers**: left and right define window boundaries
- **Frequency Tracking**: Character counts for substring validation
- **Minimum Tracking**: Keep track of smallest valid window

## 2. PROBLEM CHARACTERISTICS
- **Substring Problem**: Find contiguous substring containing all target characters
- **Minimum Length**: Among all valid substrings, find shortest
- **Character Frequency**: Need exact counts of each character
- **Optimization**: Expand and shrink window efficiently

## 3. SIMILAR PROBLEMS
- Find All Anagrams in String
- Longest Substring with At Most K Distinct Characters
- Permutation in String
- Longest Repeating Character Replacement

## 4. KEY OBSERVATIONS
- **Window Validity**: When count == 0, current window contains all needed chars
- **Shrinking Strategy**: Try to minimize window while maintaining validity
- **Character Tracking**: Use array for O(1) character access
- **Early Termination**: Can stop when window size exceeds current minimum

## 5. VARIATIONS & EXTENSIONS
- Find all minimum windows (not just one)
- Case-insensitive matching
- Unicode character support
- Multiple pattern matching

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return empty string if no valid window?"
- Edge cases: empty strings, single character, no common characters
- Time complexity: O(N + M) where N is s.length(), M is t.length()
- Space complexity: O(K) where K is character set size

## 7. COMMON MISTAKES
- Not handling empty input cases
- Incorrect frequency counting (off-by-one errors)
- Not shrinking window properly
- Missing minimum window update

## 8. OPTIMIZATION STRATEGIES
- Use fixed-size array for character frequencies (ASCII assumed)
- Single pass through string with two pointers
- Early termination when window too large
- Minimize string operations for better performance

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a specific set of letters in a book:**
- You have a book (string s) and need to find specific letters (string t)
- You want the shortest passage containing all those letters
- Use a sliding window to scan through the book
- Track which letters you've found and adjust window size

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s (text), String t (pattern to find)
2. **Goal**: Find shortest substring of s containing all characters of t
3. **Output**: Shortest valid substring or empty string

#### Phase 2: Key Insight Recognition
- **"How to track required characters?"** → Frequency array for t
- **"How to know when window is valid?"** → Count of needed characters
- **"How to find minimum?"** → Keep shrinking while valid
- **"How to slide efficiently?"** → Two pointers with frequency updates

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use sliding window with frequency counting:
1. Build frequency map of characters in t
2. Use two pointers (left, right) to define window
3. Expand right pointer to include more characters
4. When window is valid (has all needed chars), try to shrink from left
5. Keep track of minimum window found
6. Continue until right reaches end"
```

#### Phase 4: Algorithm Walkthrough
```
Example: s="ADOBECODEBANC", t="ABC"

Human thinking:
"Build frequency: A:1, B:1, C:1, count=3

Start with left=0, right=0, window=""
Right moves to find characters:
r=0: 'A' (needed) → count=2, need[A]=0
r=1: 'D' (not needed) → need[D]=-1
r=2: 'O' (not needed) → need[O]=-1
r=3: 'B' (needed) → count=1, need[B]=0
r=4: 'E' (not needed) → need[E]=-1
r=5: 'C' (needed) → count=0, need[C]=0

Now count=0, window="ADOBEC" is valid!
Try to shrink from left:
- Remove 'A' (need[A] becomes 1) → count=1, stop shrinking
- Update min window: "ADOBEC" (length 6)

Continue expanding right...
Eventually find "BANC" (length 4) as minimum ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not generate all substrings?"** → Too slow O(N³) approach
2. **"What about character encoding?"** → Assume ASCII unless specified
3. **"How to handle duplicates?"** → Frequency counts handle duplicates naturally
4. **"When to stop shrinking?"** → Stop when window becomes invalid

### Real-World Analogy
**Like finding ingredients in a recipe book:**
- You have a recipe book (string s) and need specific ingredients (string t)
- You want the shortest recipe section that mentions all required ingredients
- Scan through the book with a highlighter window
- When you have all ingredients, try to find a shorter section
- Keep track of the shortest complete recipe section found

### Human-Readable Pseudocode
```
function minWindow(s, t):
    if s or t is empty: return ""
    
    need = frequency map of characters in t
    left = 0, right = 0
    count = length of t
    minLen = infinity, start = 0
    
    while right < length of s:
        char = s[right]
        if need[char] > 0:
            count--
        need[char]--
        right++
        
        while count == 0:  // window is valid
            if right - left < minLen:
                minLen = right - left
                start = left
            
            // try to shrink window
            leftChar = s[left]
            if need[leftChar] >= 0:
                count++
            need[leftChar]++
            left++
    
    return s[start:start+minLen] if minLen != infinity else ""
```

### Execution Visualization

### Example: s="ADOBECODEBANC", t="ABC"
```
Window Evolution:
[ADOBEC] - valid (length 6) ✓
[DOBECODEBA] - valid (length 10)
[OBECODEBAN] - valid (length 9)
[BECODEBANC] - valid (length 8)
[ECODEBANC] - valid (length 7)
[CODEBANC] - valid (length 6)
[ODEBANC] - valid (length 5)
[DEBANC] - valid (length 4) ← Minimum!

Final Answer: "BANC"
```

### Key Visualization Points:
- **Frequency Array**: Tracks character requirements in current window
- **Count Variable**: Indicates window validity (0 = valid)
- **Two Pointers**: left and right define current window
- **Minimum Tracking**: Updates when shorter valid window found

### Time Complexity Breakdown:
- **Building Frequency**: O(M) where M is length of t
- **Main Loop**: O(N) where N is length of s
- **Inner Loop**: O(N) total across all iterations (amortized)
- **Total**: O(N + M) time
- **Space**: O(K) where K is character set size (typically 128 for ASCII)
- **Optimal**: Cannot do better than O(N) for this problem
*/
