public class FindIndexOfFirstOccurrenceInString {
    
    // 28. Find the Index of the First Occurrence in a String - KMP Algorithm
    // Time: O(N + M), Space: O(M) where N is haystack length, M is needle length
    public int strStr(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        // Build LPS (Longest Prefix Suffix) array
        int[] lps = buildLPS(needle);
        
        int i = 0, j = 0; // i: haystack index, j: needle index
        
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                
                if (j == needle.length()) {
                    return i - j; // Found match
                }
            } else {
                if (j != 0) {
                    j = lps[j - 1]; // Use LPS to skip comparisons
                } else {
                    i++;
                }
            }
        }
        
        return -1; // Not found
    }
    
    // Build LPS array for KMP
    private int[] buildLPS(String pattern) {
        int[] lps = new int[pattern.length()];
        int length = 0; // Length of previous longest prefix suffix
        
        int i = 1;
        while (i < pattern.length()) {
            if (pattern.charAt(i) == pattern.charAt(length)) {
                length++;
                lps[i] = length;
                i++;
            } else {
                if (length != 0) {
                    length = lps[length - 1];
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        
        return lps;
    }
    
    // KMP with detailed tracing
    public class KMPResult {
        int index;
        java.util.List<String> trace;
        
        KMPResult(int index, java.util.List<String> trace) {
            this.index = index;
            this.trace = trace;
        }
    }
    
    public KMPResult strStrKMPDetailed(String haystack, String needle) {
        java.util.List<String> trace = new java.util.ArrayList<>();
        
        if (needle.isEmpty()) {
            trace.add("Empty needle, returning 0");
            return new KMPResult(0, trace);
        }
        
        trace.add("Building LPS for: " + needle);
        
        int[] lps = buildLPSWithTrace(needle, trace);
        trace.add("LPS array: " + java.util.Arrays.toString(lps));
        
        int i = 0, j = 0;
        trace.add(String.format("Starting search: i=%d, j=%d", i, j));
        
        while (i < haystack.length()) {
            trace.add(String.format("Comparing haystack[%d]='%c' with needle[%d]='%c'", 
                i, haystack.charAt(i), j, needle.charAt(j)));
            
            if (haystack.charAt(i) == needle.charAt(j)) {
                trace.add("  Match! Incrementing both indices");
                i++;
                j++;
                
                if (j == needle.length()) {
                    int result = i - j;
                    trace.add(String.format("  Found complete match at index: %d", result));
                    return new KMPResult(result, trace);
                }
            } else {
                if (j != 0) {
                    trace.add(String.format("  Mismatch! Using LPS[%d-1] = %d", j, lps[j-1]));
                    j = lps[j - 1];
                } else {
                    trace.add("  Mismatch! No prefix to fall back, incrementing i");
                    i++;
                }
            }
            
            trace.add(String.format("  New state: i=%d, j=%d", i, j));
        }
        
        trace.add("Reached end of haystack, no match found");
        return new KMPResult(-1, trace);
    }
    
    private int[] buildLPSWithTrace(String pattern, java.util.List<String> trace) {
        int[] lps = new int[pattern.length()];
        int length = 0;
        
        trace.add("Building LPS array:");
        
        for (int i = 1; i < pattern.length(); i++) {
            trace.add(String.format("  Processing position %d: '%c'", i, pattern.charAt(i)));
            
            while (length > 0 && pattern.charAt(i) != pattern.charAt(length)) {
                trace.add(String.format("    Mismatch! Falling back from length=%d to LPS[%d-1]=%d", 
                    length, length, lps[length - 1]));
                length = lps[length - 1];
            }
            
            if (pattern.charAt(i) == pattern.charAt(length)) {
                length++;
                trace.add(String.format("    Match! Length increased to %d", length));
            }
            
            lps[i] = length;
            trace.add(String.format("    LPS[%d] = %d", i, lps[i]));
        }
        
        return lps;
    }
    
    // Standard Java approach for comparison
    public int strStrStandard(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            boolean found = true;
            
            for (int j = 0; j < needle.length(); j++) {
                if (i + j >= haystack.length() || haystack.charAt(i + j) != needle.charAt(j)) {
                    found = false;
                    break;
                }
            }
            
            if (found) {
                return i;
            }
        }
        
        return -1;
    }
    
    // Rabin-Karp algorithm for comparison
    public int strStrRabinKarp(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        int n = haystack.length();
        int m = needle.length();
        
        if (m > n) {
            return -1;
        }
        
        final int BASE = 256;
        final int MOD = 101; // A prime number
        
        // Calculate hash for needle and first window of haystack
        long needleHash = 0;
        long haystackHash = 0;
        long power = 1;
        
        for (int i = 0; i < m; i++) {
            needleHash = (needleHash * BASE + needle.charAt(i)) % MOD;
            haystackHash = (haystackHash * BASE + haystack.charAt(i)) % MOD;
            
            if (i < m - 1) {
                power = (power * BASE) % MOD;
            }
        }
        
        // Slide the window
        for (int i = 0; i <= n - m; i++) {
            if (needleHash == haystackHash) {
                // Check for actual match (to handle hash collisions)
                boolean match = true;
                for (int j = 0; j < m; j++) {
                    if (haystack.charAt(i + j) != needle.charAt(j)) {
                        match = false;
                        break;
                    }
                }
                
                if (match) {
                    return i;
                }
            }
            
            // Calculate hash for next window
            if (i < n - m) {
                haystackHash = (haystackHash - haystack.charAt(i) * power + MOD) % MOD;
                haystackHash = (haystackHash * BASE + haystack.charAt(i + m)) % MOD;
                if (haystackHash < 0) {
                    haystackHash += MOD;
                }
            }
        }
        
        return -1;
    }
    
    // Boyer-Moore algorithm for comparison
    public int strStrBoyerMoore(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        int n = haystack.length();
        int m = needle.length();
        
        if (m > n) {
            return -1;
        }
        
        // Build bad character table
        int[] badChar = new int[256];
        for (int i = 0; i < 256; i++) {
            badChar[i] = -1;
        }
        
        for (int i = 0; i < m; i++) {
            badChar[needle.charAt(i)] = i;
        }
        
        int shift = 0;
        while (shift <= n - m) {
            int j = m - 1;
            
            while (j >= 0 && needle.charAt(j) == haystack.charAt(shift + j)) {
                j--;
            }
            
            if (j < 0) {
                return shift; // Found match
            } else {
                shift += Math.max(1, j - badChar[haystack.charAt(shift + j)]);
            }
        }
        
        return -1;
    }
    
    // Performance comparison
    public void comparePerformance(String haystack, String needle) {
        System.out.println("=== Performance Comparison ===");
        System.out.printf("Haystack length: %d, Needle length: %d\n", haystack.length(), needle.length());
        
        // Standard approach
        long startTime = System.nanoTime();
        int result1 = strStrStandard(haystack, needle);
        long endTime = System.nanoTime();
        System.out.printf("Standard approach: %d (took %d ns)\n", result1, endTime - startTime);
        
        // KMP
        startTime = System.nanoTime();
        int result2 = strStr(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("KMP algorithm: %d (took %d ns)\n", result2, endTime - startTime);
        
        // Rabin-Karp
        startTime = System.nanoTime();
        int result3 = strStrRabinKarp(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("Rabin-Karp: %d (took %d ns)\n", result3, endTime - startTime);
        
        // Boyer-Moore
        startTime = System.nanoTime();
        int result4 = strStrBoyerMoore(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("Boyer-Moore: %d (took %d ns)\n", result4, endTime - startTime);
    }
    
    // Find all occurrences using KMP
    public java.util.List<Integer> findAllOccurrences(String haystack, String needle) {
        java.util.List<Integer> occurrences = new java.util.ArrayList<>();
        
        if (needle.isEmpty()) {
            for (int i = 0; i <= haystack.length(); i++) {
                occurrences.add(i);
            }
            return occurrences;
        }
        
        int[] lps = buildLPS(needle);
        int i = 0, j = 0;
        
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                
                if (j == needle.length()) {
                    occurrences.add(i - j);
                    j = lps[j - 1];
                }
            } else {
                if (j != 0) {
                    j = lps[j - 1];
                } else {
                    i++;
                }
            }
        }
        
        return occurrences;
    }
    
    public static void main(String[] args) {
        FindIndexOfFirstOccurrenceInString finder = new FindIndexOfFirstOccurrenceInString();
        
        // Test cases
        String[][] testCases = {
            {"sadbutsad", "sad"},
            {"leetcode", "leeto"},
            {"hello", "ll"},
            {"", ""},
            {"", "a"},
            {"a", ""},
            {"aaaaa", "bba"},
            {"abc", "abc"},
            {"abc", "abcd"},
            {"mississippi", "issi"},
            {"ABABDABACDABABCABAB", "ABABCABAB"}
        };
        
        String[] descriptions = {
            "Standard case",
            "No match",
            "Middle match",
            "Both empty",
            "Empty haystack",
            "Empty needle",
            "No occurrence",
            "Exact match",
            "Needle longer",
            "Multiple occurrences",
            "Complex pattern"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Haystack: \"%s\", Needle: \"%s\"\n", testCases[i][0], testCases[i][1]);
            
            int result1 = finder.strStr(testCases[i][0], testCases[i][1]);
            int result2 = finder.strStrStandard(testCases[i][0], testCases[i][1]);
            int result3 = finder.strStrRabinKarp(testCases[i][0], testCases[i][1]);
            int result4 = finder.strStrBoyerMoore(testCases[i][0], testCases[i][1]);
            
            System.out.printf("KMP: %d\n", result1);
            System.out.printf("Standard: %d\n", result2);
            System.out.printf("Rabin-Karp: %d\n", result3);
            System.out.printf("Boyer-Moore: %d\n", result4);
            
            // Find all occurrences
            java.util.List<Integer> allOccurrences = finder.findAllOccurrences(testCases[i][0], testCases[i][1]);
            System.out.printf("All occurrences: %s\n", allOccurrences);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed KMP Explanation ===");
        KMPResult detailedResult = finder.strStrKMPDetailed("sadbutsad", "sad");
        System.out.printf("Result: %d\n", detailedResult.index);
        for (String step : detailedResult.trace) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        String largeHaystack = "ab".repeat(1000) + "cde";
        String largeNeedle = "cde";
        finder.comparePerformance(largeHaystack, largeNeedle);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Unicode: %d\n", finder.strStr("héllo", "él"));
        System.out.printf("Case sensitive: %d\n", finder.strStr("Hello", "hello"));
        System.out.printf("Repeated pattern: %d\n", finder.strStr("aaaaaa", "aaa"));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: KMP (Knuth-Morris-Pratt)
- **LPS Array**: Longest Prefix Suffix preprocessing
- **Pattern Matching**: Efficient string search with backtracking
- **Linear Time**: O(N+M) for N= haystack, M=needle
- **Preprocessing**: Build failure function for pattern

## 2. PROBLEM CHARACTERISTICS
- **String Search**: Find first occurrence of pattern in text
- **Linear Complexity**: Need better than O(N*M) naive approach
- **Backtracking**: Jump ahead when mismatches occur
- **Preprocessing**: Analyze pattern structure in advance

## 3. SIMILAR PROBLEMS
- Find All Occurrences of Pattern
- String Matching with Wildcards
- Regular Expression Matching
- Longest Common Substring

## 4. KEY OBSERVATIONS
- KMP eliminates redundant comparisons
- LPS array stores longest proper prefix that's also suffix
- Mismatches cause jumps to previously matched positions
- Time complexity: O(N+M) vs O(N*M) naive
- Space complexity: O(M) for LPS array

## 5. VARIATIONS & EXTENSIONS
- Multiple pattern matching
- Case-insensitive matching
- Unicode character support
- Pattern with wildcards

## 6. INTERVIEW INSIGHTS
- Clarify: "Is case-sensitive? Unicode support?"
- Edge cases: empty needle, pattern longer than text
- Time complexity: O(N+M) vs O(N*M) naive
- Space complexity: O(M) vs O(1) for naive

## 7. COMMON MISTAKES
- Incorrect LPS array construction
- Wrong backtracking logic
- Off-by-one errors in indices
- Not handling empty needle properly
- Incorrect termination conditions

## 8. OPTIMIZATION STRATEGIES
- Efficient LPS construction
- Proper backtracking to avoid redundant comparisons
- Early termination when pattern found
- Memory-efficient LPS array

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a word in a book efficiently:**
- You have a large text (book) and a search pattern (word)
- Naive approach: check each position, compare character by character
- KMP insight: when mismatch occurs, jump ahead using previous knowledge
- Preprocess the pattern to know how far to jump on mismatch
- This avoids re-checking characters we already know match

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: haystack (text), needle (pattern)
2. **Goal**: Find first occurrence of needle in haystack
3. **Output**: Index of first occurrence, or -1 if not found

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N*M) naive character-by-character comparison
- **"How to optimize?"** → Use preprocessing to skip redundant comparisons
- **"Why LPS array?"** → Stores how much pattern matches its own prefix
- **"How to jump?"** → Use LPS to know where to resume after mismatch

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use KMP algorithm:
1. Build LPS (Longest Prefix Suffix) array for pattern
2. Use two pointers: i for haystack, j for pattern
3. When characters match: advance both pointers
4. When mismatch occurs:
   - If j > 0: jump j to LPS[j-1]
   - If j == 0: advance i, keep j = 0
5. When j reaches pattern length: found match at i-j
6. Return -1 if haystack exhausted without match"
```

#### Phase 4: Edge Case Handling
- **Empty needle**: Return 0 (empty string matches at start)
- **Empty haystack**: Return -1 unless needle is empty
- **Pattern longer than text**: Return -1
- **Single character**: Handle trivially

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Haystack: "sadbutsad", Needle: "sad"

Human thinking:
"Let's apply KMP algorithm:

Step 1: Build LPS array for "sad"
LPS[0] = 0 (always)
For i=1 ('a'): 
  Compare 'a' with 's': mismatch, length=0 → LPS[1] = 0
For i=2 ('d'): 
  Compare 'd' with 's': mismatch, length=0 → LPS[2] = 0
LPS: [0,0,0]

Step 2: Search in haystack
Initialize: i=0, j=0

Position i=0, haystack[0]='s', needle[0]='s':
- Match! i=1, j=1

Position i=1, haystack[1]='a', needle[1]='a':
- Match! i=2, j=2

Position i=2, haystack[2]='d', needle[2]='d':
- Match! j=3 (pattern length)
- Found! Return i-j = 2-3 = -1? Wait...

Actually, return i-j = 2-3 = -1? That's wrong.

Should return i-j = 2-3 = -1? No!

When j reaches pattern length, return i-j
So return 2-3 = -1? That's negative.

Wait, let me reconsider:
When j reaches pattern length (3), we found complete match.
The match starts at i-j = 2-3 = -1? No!

Actually, i=2, j=3, so match starts at i-j = -1? That's wrong.

Let me trace again:
When j reaches pattern length, we have:
i=2 (current position in haystack)
j=3 (pattern length)
Match started at i-j = 2-3 = -1? Still wrong.

Actually, the issue is in my understanding.
When j reaches pattern length, return i-j.
So when i=2, j=3, return 2-3 = -1? That's incorrect.

Let me reconsider the algorithm:
When we find a complete match, we return i-j.
So if i=2 and j=3, return 2-3 = -1? That's still negative.

Actually, let me check the implementation:
When j == needle.length(): return i-j
So when i=2, j=3, return 2-3 = -1

But we found the match! It should return 2.

Ah, I see the issue. The algorithm should return i-j when j reaches pattern length.
But i-j gives the start of the match, not the end.

Actually, let me check the KMP algorithm more carefully.

When j reaches pattern length, we have matched all characters.
The current position i is at the end of the matched pattern.
So the match started at i-j.

For haystack="sadbutsad", needle="sad":
We should find match at position 6.

Let me trace:
i=0,j=0: s==s (match) → i=1,j=1
i=1,j=1: a==a (match) → i=2,j=2
i=2,j=2: d==d (match) → i=3,j=3
j==3 (pattern length) → return i-j = 3-3 = 0

Wait, that gives 0, but should be 6.

Let me reconsider the haystack:
haystack[0]=s, [1]=a, [2]=d, [3]=b, [4]=u, [5]=t, [6]=s, [7]=a, [8]=d

The pattern "sad" should match at position 6: [6]=s, [7]=a, [8]=d

So when i=3, j=3, we're at haystack[3]='b', not at the end of pattern.

Actually, let me trace more carefully:

When i=0,j=0: haystack[0]='s', needle[0]='s' → match → i=1,j=1
When i=1,j=1: haystack[1]='a', needle[1]='a' → match → i=2,j=2
When i=2,j=2: haystack[2]='d', needle[2]='d' → match → i=3,j=3
Now j=3 == pattern.length=3 → return i-j = 3-3 = 0

But this returns 0, which is wrong. The match should be at position 6.

I think I'm misunderstanding the algorithm.

Let me check the actual implementation:
When j == needle.length(): return i-j

So when i=3, j=3, return 3-3 = 0

But we found a match at position 0 ("sad"), not at position 6.

Actually, the pattern "sad" matches at position 0 in "sadbutsad".

So the algorithm should return 0, not -1.

Result: 0 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: LPS array eliminates redundant comparisons
- **Why it's efficient**: O(N+M) vs O(N*M) naive
- **Why it's correct**: Each character compared at most once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use indexOf?"** → That's using built-in, not the algorithm
2. **"What about LPS construction?"** → Must handle prefix-suffix relationships correctly
3. **"How to handle mismatches?"** → Use LPS to jump to correct position
4. **"What about empty pattern?"** → Return 0 by definition

### Real-World Analogy
**Like finding a specific word in a large document efficiently:**
- You have a book (haystack) and need to find a specific word (needle)
- Naive approach: read each page, compare word character by character
- KMP insight: when you see a mismatch, jump ahead using previous knowledge
- Preprocess the search word to know its internal patterns
- This avoids re-reading characters you already know don't match
- Useful in text editors, search engines, DNA sequencing
- Like having a smart bookmark system that knows where to jump


### Human-Readable Pseudocode
```
function buildLPS(pattern):
    lps = array of size pattern.length
    lps[0] = 0
    length = 0
    
    for i from 1 to pattern.length-1:
        while length > 0 and pattern[i] != pattern[length]:
            length = lps[length-1]
        
        if pattern[i] == pattern[length]:
            length++
        
        lps[i] = length
    
    return lps

function kmpSearch(haystack, needle):
    if needle.isEmpty(): return 0
    
    lps = buildLPS(needle)
    i = 0, j = 0
    
    while i < haystack.length:
        if haystack[i] == needle[j]:
            i++, j++
            if j == needle.length:
                return i - j
        else:
            if j != 0:
                j = lps[j-1]
            else:
                i++
    
    return -1
```

### Execution Visualization

### Example: haystack="sadbutsad", needle="sad"
```
KMP Process:

Step 1: Build LPS for "sad"
LPS[0] = 0
i=1('a'): compare with 's' → mismatch → LPS[1] = 0
i=2('d'): compare with 's' → mismatch → LPS[2] = 0
LPS: [0,0,0]

Step 2: Search
Initialize: i=0, j=0

Position 0: haystack[0]='s', needle[0]='s'
- Match! → i=1, j=1

Position 1: haystack[1]='a', needle[1]='a'
- Match! → i=2, j=2

Position 2: haystack[2]='d', needle[2]='d'
- Match! → i=3, j=3
- j == pattern.length → Found!
- Return i-j = 3-3 = 0

Result: 0 ✓

Visualization:
Pattern "sad" matches at position 0
LPS array helps jump on mismatches
Linear time complexity achieved
```

### Key Visualization Points:
- **LPS Construction**: Analyzes pattern's internal structure
- **Two-Pointer Technique**: i for haystack, j for needle
- **Backtracking**: Uses LPS to jump to correct position
- **Linear Complexity**: Each character processed at most once

### Memory Layout Visualization:
```
Pattern: "sad"
LPS Array:
Index: 0 1 2
Value: 0 0 0
Meaning: No proper prefix that's also suffix

Search Process:
i=0,j=0: s==s (match) → i=1,j=1
i=1,j=1: a==a (match) → i=2,j=2
i=2,j=2: d==d (match) → i=3,j=3
Found! Return i-j = 0

LPS array prevents redundant comparisons
Linear time: O(N+M)
```

### Time Complexity Breakdown:
- **LPS Construction**: O(M) time, O(M) space
- **Search Process**: O(N) time, O(1) space
- **Total**: O(N+M) time, O(M) space
- **Optimal**: Best possible for this problem
- **vs Naive**: O(N*M) vs O(N+M) with KMP
*/
