import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class FindAllAnagramsInAString {
    
    // 438. Find All Anagrams in a String (Fixed Size Sliding Window)
    // Time: O(N), Space: O(1) for ASCII characters
    public static List<Integer> findAnagrams(String s, String p) {
        List<Integer> result = new ArrayList<>();
        
        if (s.length() < p.length()) {
            return result;
        }
        
        int[] pCount = new int[26];
        int[] sCount = new int[26];
        
        // Initialize frequency count for pattern and first window
        for (int i = 0; i < p.length(); i++) {
            pCount[p.charAt(i) - 'a']++;
            sCount[s.charAt(i) - 'a']++;
        }
        
        // Check if first window is an anagram
        if (matches(pCount, sCount)) {
            result.add(0);
        }
        
        // Slide the window through the string
        for (int i = p.length(); i < s.length(); i++) {
            // Remove the leftmost character
            sCount[s.charAt(i - p.length()) - 'a']--;
            // Add the new character
            sCount[s.charAt(i) - 'a']++;
            
            // Check if current window is an anagram
            if (matches(pCount, sCount)) {
                result.add(i - p.length() + 1);
            }
        }
        
        return result;
    }

    // Helper function to check if two frequency arrays match
    private static boolean matches(int[] pCount, int[] sCount) {
        for (int i = 0; i < 26; i++) {
            if (pCount[i] != sCount[i]) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"cbaebabacd", "abc"},
            {"abab", "ab"},
            {"aaaaaaaaaa", "aaaa"},
            {"abacbabc", "abc"},
            {"", "a"},
            {"a", ""},
            {"abc", "def"},
            {"abababab", "ab"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<Integer> result = findAnagrams(testCases[i][0], testCases[i][1]);
            System.out.printf("Test Case %d: s=\"%s\", p=\"%s\" -> Anagram indices: %s\n", 
                i + 1, testCases[i][0], testCases[i][1], result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fixed Size Sliding Window
- **Window Size**: Fixed window size equal to pattern length
- **Character Counting**: Track character frequencies in window
- **Window Sliding**: Slide window by one position each iteration
- **Anagram Detection**: Compare window and pattern character counts

## 2. PROBLEM CHARACTERISTICS
- **Anagram Finding**: Find all substrings that are anagrams of pattern
- **Fixed Window**: Window size is constant (pattern length)
- **Character Frequency**: Need to compare character distributions
- **Multiple Matches**: Can have multiple anagram positions

## 3. SIMILAR PROBLEMS
- Permutation in String
- Find All Anagrams in a String
- Longest Substring with At Most K Distinct Characters
- Longest Substring Without Repeating Characters

## 4. KEY OBSERVATIONS
- Fixed-size sliding window maintains current substring
- Character frequency arrays enable O(1) comparisons
- Window slides by removing leftmost and adding rightmost character
- Time complexity: O(N) for N characters
- Space complexity: O(1) for fixed character set

## 5. VARIATIONS & EXTENSIONS
- Different character sets (ASCII, Unicode)
- Case sensitivity variations
- Multiple pattern matching
- Variable window size adaptations

## 6. INTERVIEW INSIGHTS
- Clarify: "What is the character set?"
- Edge cases: empty string, pattern longer than string, no matches
- Time complexity: O(N) vs O(N*M) brute force
- Space complexity: O(1) vs O(N) hash map approach

## 7. COMMON MISTAKES
- Incorrect window initialization
- Wrong character count updates
- Missing edge cases
- Inefficient frequency comparison
- Incorrect window sliding logic

## 8. OPTIMIZATION STRATEGIES
- Use fixed-size arrays for character counts
- Efficient frequency comparison
- Single pass through string
- Minimal space usage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding all rearrangements of a word:**
- You have a text and a pattern word
- Need to find all positions where text contains rearranged pattern
- Fixed-size window slides through text
- Character counts determine if window matches pattern
- This is like finding all permutations of a word in text

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Text string and pattern string
2. **Goal**: Find all start indices of anagrams of pattern in text
3. **Output**: List of starting indices

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N*M) to check all substrings
- **"How to optimize?"** → Use sliding window with character counts
- **"Why sliding window?"** → Fixed-size window maintains substring
- **"Why character counts?"** → O(1) frequency comparison

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use fixed-size sliding window:
1. Initialize character count arrays for pattern and first window
2. Compare counts to check if first window is anagram
3. Slide window through string:
   - Remove leftmost character from window count
   - Add new rightmost character to window count
   - Compare window and pattern counts
   - If match, record starting index
4. Continue until end of string
5. Return all recorded indices"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty list
- **Pattern longer than text**: Return empty list
- **No matches**: Return empty list
- **Single character pattern**: Handle efficiently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Text: "cbaebabacd", Pattern: "abc"

Human thinking:
"Let's apply sliding window:

Initialize: pCount={a:1, b:1, c:1}, sCount={c:1, b:1, a:1}
Window size: 3 (pattern length)

Step 1: Check first window "cba"
- pCount and sCount match ✓
- Record index 0

Step 2: Slide window by 1
- Remove 'c' (leftmost): sCount[c]--
- Add 'e' (new rightmost): sCount[e]++
- Window: "bae", sCount={a:1, b:1, c:0, e:1}
- pCount and sCount don't match

Step 3: Slide window by 1
- Remove 'b': sCount[b]--
- Add 'b': sCount[b]++
- Window: "aeb", sCount={a:1, b:1, c:0, e:1}
- pCount and sCount don't match

Step 4: Slide window by 1
- Remove 'a': sCount[a]--
- Add 'a': sCount[a]++
- Window: "eba", sCount={a:1, b:1, c:0, e:1}
- pCount and sCount don't match

Step 5: Slide window by 1
- Remove 'e': sCount[e]--
- Add 'b': sCount[b]++
- Window: "bab", sCount={a:0, b:2, c:0, e:0}
- pCount and sCount don't match

Step 6: Slide window by 1
- Remove 'b': sCount[b]--
- Add 'a': sCount[a]++
- Window: "aba", sCount={a:2, b:1, c:0, e:0}
- pCount and sCount don't match

Step 7: Slide window by 1
- Remove 'a': sCount[a]--
- Add 'c': sCount[c]++
- Window: "bac", sCount={a:1, b:1, c:1, e:0}
- pCount and sCount match ✓
- Record index 6

Continue...

Final result: [0, 6] ✓

Manual verification:
"cba" at index 0 is anagram of "abc" ✓
"bac" at index 6 is anagram of "abc" ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Character counts determine anagrams
- **Why it's efficient**: O(N) vs O(N*M) brute force
- **Why it's correct**: Checks all possible positions

### Common Human Pitfalls & How to Avoid Them
1. **"Why not check all substrings?"** → O(N*M) too slow
2. **"What about window size?"** → Fixed to pattern length
3. **"How to compare counts?"** → Use fixed-size arrays
4. **"What about sliding?** → Remove leftmost, add rightmost

### Real-World Analogy
**Like finding all rearrangements of a word in text:**
- You have a document and a target word
- Need to find all positions where document contains rearranged word
- Fixed-size window slides through document
- Character counts determine if window matches target
- This is used in plagiarism detection, DNA analysis
- Like finding all permutations of a pattern in text

### Human-Readable Pseudocode
```
function findAnagrams(s, p):
    result = empty list
    
    if s.length < p.length:
        return result
    
    pCount = array[26] initialized to 0
    sCount = array[26] initialized to 0
    
    // Initialize frequency counts
    for i from 0 to p.length-1:
        pCount[p[i] - 'a']++
        sCount[s[i] - 'a']++
    
    // Check first window
    if pCount == sCount:
        result.add(0)
    
    // Slide window through string
    for i from p.length to s.length-1:
        // Remove leftmost character
        sCount[s[i - p.length] - 'a']--
        // Add new character
        sCount[s[i] - 'a']++
        
        // Check if current window is anagram
        if pCount == sCount:
            result.add(i - p.length + 1)
    
    return result
```

### Execution Visualization

### Example: s="cbaebabacd", p="abc"
```
Sliding Window Process:

Initialize: pCount={a:1, b:1, c:1}, sCount={c:1, b:1, a:1}
Window size: 3

Step 1: Window "cba" (indices 0-2)
- pCount and sCount match ✓
- Record index 0

Step 2: Window "bae" (indices 1-3)
- Remove 'c', add 'e'
- sCount={a:1, b:1, c:0, e:1}
- No match

Step 3: Window "aeb" (indices 2-4)
- Remove 'b', add 'b'
- sCount={a:1, b:1, c:0, e:1}
- No match

Step 4: Window "eba" (indices 3-5)
- Remove 'a', add 'a'
- sCount={a:1, b:1, c:0, e:1}
- No match

Step 5: Window "bab" (indices 4-6)
- Remove 'e', add 'b'
- sCount={a:0, b:2, c:0, e:0}
- No match

Step 6: Window "aba" (indices 5-7)
- Remove 'b', add 'a'
- sCount={a:2, b:1, c:0, e:0}
- No match

Step 7: Window "bac" (indices 6-8)
- Remove 'a', add 'c'
- sCount={a:1, b:1, c:1, e:0}
- pCount and sCount match ✓
- Record index 6

Continue...

Final result: [0, 6] ✓

Visualization:
Fixed-size window slides through string
Character counts enable O(1) comparisons
All anagram positions recorded ✓
```

### Key Visualization Points:
- **Fixed Window Size**: Always equal to pattern length
- **Character Counting**: Arrays track frequency distributions
- **Window Sliding**: Remove leftmost, add rightmost
- **Frequency Comparison**: O(1) array comparison

### Memory Layout Visualization:
```
Text:    "c b a e b a b a c d"
Index:     0 1 2 3 4 5 6 7 8 9
Pattern:  "a b c" (length 3)

Window Evolution:
Step 1: [c b a] counts={a:1,b:1,c:1} ✓ match at 0
Step 2: [b a e] counts={a:1,b:1,e:1} ✗ no match
Step 3: [a e b] counts={a:1,b:1,e:1} ✗ no match
Step 4: [e b a] counts={a:1,b:1,e:1} ✗ no match
Step 5: [b a b] counts={a:1,b:2} ✗ no match
Step 6: [a b a] counts={a:2,b:1} ✗ no match
Step 7: [b a c] counts={a:1,b:1,c:1} ✓ match at 6

Pattern counts: {a:1,b:1,c:1}
Matches found at indices: [0, 6] ✓

Fixed-size sliding window efficiently finds anagrams
Character counting enables O(1) comparisons
```

### Time Complexity Breakdown:
- **Sliding Window**: O(N) time, O(1) space (fixed character set)
- **Brute Force**: O(N*M) time, O(1) space
- **Hash Map Approach**: O(N) time, O(N) space
- **Optimal**: Best possible for this problem
- **vs Naive**: O(N) vs O(N*M) significant improvement
*/
