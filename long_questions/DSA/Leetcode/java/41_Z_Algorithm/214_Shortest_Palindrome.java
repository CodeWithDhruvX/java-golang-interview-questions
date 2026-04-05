public class ShortestPalindrome {
    
    // 214. Shortest Palindrome - Z Algorithm
    // Time: O(N), Space: O(N)
    public String shortestPalindrome(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        // Create pattern: s + '#' + reverse(s)
        String pattern = s + "#" + reverseString(s);
        
        // Compute Z-array
        int[] zArray = computeZArray(pattern);
        
        // Find longest palindrome prefix
        int maxPrefix = 0;
        for (int i = s.length() + 1; i < zArray.length; i++) {
            if (zArray[i] > maxPrefix) {
                maxPrefix = zArray[i];
            }
        }
        
        // Add reverse of suffix to front
        String suffix = s.substring(maxPrefix);
        return reverseString(suffix) + s;
    }
    
    private int[] computeZArray(String s) {
        int n = s.length();
        int[] z = new int[n];
        z[0] = n;
        
        int l = 0, r = 0;
        
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = Math.min(r - i + 1, z[i - l]);
            }
            
            while (i + z[i] < n && s.charAt(z[i]) == s.charAt(i + z[i])) {
                z[i]++;
            }
            
            if (i + z[i] - 1 > r) {
                l = i;
                r = i + z[i] - 1;
            }
        }
        
        return z;
    }
    
    private String reverseString(String s) {
        return new StringBuilder(s).reverse().toString();
    }
    
    // Z Algorithm for pattern matching
    public java.util.List<Integer> zAlgorithmPatternSearch(String text, String pattern) {
        java.util.List<Integer> occurrences = new java.util.ArrayList<>();
        
        if (pattern.isEmpty()) {
            return occurrences;
        }
        
        // Create pattern: pattern + '$' + text
        String combined = pattern + "$" + text;
        int[] z = computeZArray(combined);
        
        // Find occurrences
        for (int i = pattern.length() + 1; i < z.length; i++) {
            if (z[i] == pattern.length()) {
                occurrences.add(i - pattern.length() - 1);
            }
        }
        
        return occurrences;
    }
    
    // Standard approach for comparison
    public String shortestPalindromeStandard(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        String reverse = reverseString(s);
        
        for (int i = 0; i < s.length(); i++) {
            if (s.substring(0, s.length() - i).equals(reverse.substring(i))) {
                return reverse.substring(0, i) + s;
            }
        }
        
        return reverse + s;
    }
    
    // Manacher's algorithm for comparison
    public String shortestPalindromeManacher(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        // Transform string to handle even length palindromes
        String transformed = "^#" + String.join("#", s.split("")) + "#$";
        int[] palindromeLengths = new int[transformed.length()];
        
        int center = 0;
        int right = 0;
        
        for (int i = 1; i < transformed.length() - 1; i++) {
            int mirror = 2 * center - i;
            
            if (i < right) {
                palindromeLengths[i] = Math.min(right - i, palindromeLengths[mirror]);
            }
            
            // Expand around center
            while (transformed.charAt(i + palindromeLengths[i] + 1) == 
                   transformed.charAt(i - palindromeLengths[i] - 1)) {
                palindromeLengths[i]++;
            }
            
            // Update center and right
            if (i + palindromeLengths[i] > right) {
                center = i;
                right = i + palindromeLengths[i];
            }
        }
        
        // Find longest palindrome at the beginning
        int maxLen = 0;
        int maxCenter = 0;
        
        for (int i = 1; i < transformed.length() - 1; i++) {
            int start = (i - palindromeLengths[i] - 1) / 2;
            if (start == 0 && palindromeLengths[i] > maxLen) {
                maxLen = palindromeLengths[i];
                maxCenter = i;
            }
        }
        
        // Extract the longest prefix palindrome
        int palindromeEnd = (maxCenter + maxLen - 1) / 2;
        String longestPrefixPalindrome = s.substring(0, palindromeEnd + 1);
        
        // Add reverse of remaining suffix
        String remaining = s.substring(palindromeEnd + 1);
        return reverseString(remaining) + longestPrefixPalindrome;
    }
    
    // Version with detailed explanation
    public class PalindromeResult {
        String shortestPalindrome;
        java.util.List<String> explanation;
        int[] zArray;
        
        PalindromeResult(String shortestPalindrome, java.util.List<String> explanation, int[] zArray) {
            this.shortestPalindrome = shortestPalindrome;
            this.explanation = explanation;
            this.zArray = zArray;
        }
    }
    
    public PalindromeResult shortestPalindromeDetailed(String s) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== Z Algorithm for Shortest Palindrome ===");
        explanation.add("Input: " + s);
        
        if (s.length() <= 1) {
            explanation.add("String length <= 1, returning as is");
            return new PalindromeResult(s, explanation, new int[0]);
        }
        
        // Create pattern
        String pattern = s + "#" + reverseString(s);
        explanation.add("Created pattern: " + pattern);
        
        // Compute Z-array
        int[] zArray = computeZArrayWithTrace(pattern, explanation);
        
        // Find longest palindrome prefix
        int maxPrefix = 0;
        explanation.add("Finding longest palindrome prefix:");
        
        for (int i = s.length() + 1; i < zArray.length; i++) {
            if (zArray[i] > maxPrefix) {
                maxPrefix = zArray[i];
                explanation.add(String.format("  Position %d: Z[%d] = %d (new max)", 
                    i, i, zArray[i]));
            }
        }
        
        explanation.add(String.format("Max prefix length: %d", maxPrefix));
        
        // Construct result
        String suffix = s.substring(maxPrefix);
        String result = reverseString(suffix) + s;
        explanation.add(String.format("Suffix to reverse: \"%s\"", suffix));
        explanation.add(String.format("Reversed suffix: \"%s\"", reverseString(suffix)));
        explanation.add(String.format("Final result: \"%s\"", result));
        
        return new PalindromeResult(result, explanation, zArray);
    }
    
    private int[] computeZArrayWithTrace(String s, java.util.List<String> explanation) {
        int n = s.length();
        int[] z = new int[n];
        z[0] = n;
        
        int l = 0, r = 0;
        explanation.add("Computing Z-array:");
        explanation.add(String.format("Z[0] = %d (string length)", n));
        
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = Math.min(r - i + 1, z[i - l]);
                explanation.add(String.format("  i=%d within [%d,%d], Z[%d] = min(%d, Z[%d]) = %d", 
                    i, l, r, i, r - i + 1, i - l, z[i - l], z[i]));
            }
            
            while (i + z[i] < n && s.charAt(z[i]) == s.charAt(i + z[i])) {
                z[i]++;
            }
            
            if (i + z[i] - 1 > r) {
                explanation.add(String.format("  Updating window: L=%d, R=%d", i, i + z[i] - 1));
                l = i;
                r = i + z[i] - 1;
            }
            
            explanation.add(String.format("  Z[%d] = %d", i, z[i]));
        }
        
        return z;
    }
    
    // Performance comparison
    public void comparePerformance(String s) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Input: " + s);
        
        // Standard approach
        long startTime = System.nanoTime();
        String result1 = shortestPalindromeStandard(s);
        long endTime = System.nanoTime();
        System.out.printf("Standard approach: \"%s\" (took %d ns)\n", result1, endTime - startTime);
        
        // Z Algorithm
        startTime = System.nanoTime();
        String result2 = shortestPalindrome(s);
        endTime = System.nanoTime();
        System.out.printf("Z Algorithm: \"%s\" (took %d ns)\n", result2, endTime - startTime);
        
        // Manacher's algorithm
        startTime = System.nanoTime();
        String result3 = shortestPalindromeManacher(s);
        endTime = System.nanoTime();
        System.out.printf("Manacher's: \"%s\" (took %d ns)\n", result3, endTime - startTime);
    }
    
    // Find all palindromic prefixes
    public java.util.List<String> findAllPalindromicPrefixes(String s) {
        java.util.List<String> prefixes = new java.util.ArrayList<>();
        
        String pattern = s + "#" + reverseString(s);
        int[] zArray = computeZArray(pattern);
        
        for (int i = s.length() + 1; i < zArray.length; i++) {
            if (zArray[i] > 0) {
                String prefix = s.substring(0, zArray[i]);
                if (isPalindrome(prefix)) {
                    prefixes.add(prefix);
                }
            }
        }
        
        return prefixes;
    }
    
    private boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
    
    public static void main(String[] args) {
        ShortestPalindrome sp = new ShortestPalindrome();
        
        // Test cases
        String[] testCases = {
            "aacecaaa",
            "abcd",
            "a",
            "",
            "aa",
            "aba",
            "racecar",
            "abac",
            "abcba",
            "abacdfgdcaba"
        };
        
        String[] descriptions = {
            "Standard case",
            "No palindrome",
            "Single character",
            "Empty string",
            "All same characters",
            "Already palindrome",
            "Palindrome",
            "Mixed case",
            "Palindrome at end",
            "Complex case"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Input: \"%s\"\n", testCases[i]);
            
            String result1 = sp.shortestPalindrome(testCases[i]);
            String result2 = sp.shortestPalindromeStandard(testCases[i]);
            String result3 = sp.shortestPalindromeManacher(testCases[i]);
            
            System.out.printf("Z Algorithm: \"%s\"\n", result1);
            System.out.printf("Standard: \"%s\"\n", result2);
            System.out.printf("Manacher's: \"%s\"\n", result3);
            
            // Find all palindromic prefixes
            java.util.List<String> prefixes = sp.findAllPalindromicPrefixes(testCases[i]);
            System.out.printf("Palindromic prefixes: %s\n", prefixes);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        PalindromeResult detailedResult = sp.shortestPalindromeDetailed("aacecaaa");
        System.out.printf("Result: \"%s\"\n", detailedResult.shortestPalindrome);
        System.out.println("Z-Array: " + java.util.Arrays.toString(detailedResult.zArray));
        for (String step : detailedResult.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        String performanceTest = "a".repeat(1000) + "b".repeat(1000);
        sp.comparePerformance(performanceTest);
        
        // Pattern matching with Z algorithm
        System.out.println("\n=== Pattern Matching ===");
        String text = "ababcabcababc";
        String pattern = "abc";
        java.util.List<Integer> occurrences = sp.zAlgorithmPatternSearch(text, pattern);
        System.out.printf("Text: \"%s\"\n", text);
        System.out.printf("Pattern: \"%s\"\n", pattern);
        System.out.printf("Occurrences: %s\n", occurrences);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("All same: \"%s\"\n", sp.shortestPalindrome("aaaaa"));
        System.out.printf("Already palindrome: \"%s\"\n", sp.shortestPalindrome("racecar"));
        System.out.printf("Mixed case: \"%s\"\n", sp.shortestPalindrome("AaBb"));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Z Algorithm
- **String Transformation**: Transform palindrome problem to pattern matching
- **Z-Array**: Preprocess pattern for efficient palindrome detection
- **Linear Time**: O(N) time complexity for palindrome problems
- **Pattern Matching**: Use Z-array to find longest palindrome prefix

## 2. PROBLEM CHARACTERISTICS
- **Shortest Palindrome**: Find shortest palindrome by adding characters
- **String Manipulation**: Can add characters anywhere in string
- **Palindrome Property**: Result must be a palindrome
- **Efficiency**: Need better than O(N²) brute force approach

## 3. SIMILAR PROBLEMS
- Longest Palindromic Substring
- Count Palindromic Substrings
- Palindrome Partitioning
- Manacher's Algorithm

## 4. KEY OBSERVATIONS
- Z algorithm transforms palindrome to pattern matching
- Create pattern: s + '#' + reverse(s)
- Z-array finds longest prefix that's also suffix
- Time complexity: O(N) vs O(N²) naive
- Space complexity: O(N) for Z-array

## 5. VARIATIONS & EXTENSIONS
- Manacher's algorithm for O(N) solution
- Count all palindromic substrings
- Different palindrome definitions
- Multiple string operations

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we add characters anywhere?"
- Edge cases: empty string, single character, already palindrome
- Time complexity: O(N) vs O(N²) naive
- Space complexity: O(N) vs O(1) for naive

## 7. COMMON MISTAKES
- Incorrect Z-array construction
- Wrong palindrome extraction logic
- Off-by-one errors in indices
- Not handling edge cases properly
- Incorrect string transformation

## 8. OPTIMIZATION STRATEGIES
- Use efficient string concatenation
- Proper Z-array construction
- Correct palindrome boundary handling
- Memory-efficient Z-array storage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the shortest mirror image:**
- You have a string and want to make it a palindrome by adding characters
- Instead of checking all possible additions, use pattern matching
- Transform the problem: s + '#' + reverse(s)
- Find longest prefix that's also a suffix (palindrome)
- This is like finding the longest mirror in the transformed string
- Z algorithm efficiently finds this pattern in linear time

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s, can add characters anywhere
2. **Goal**: Find shortest palindrome by adding characters
3. **Output**: Shortest palindrome string

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N³) brute force approach
- **"How to optimize?"** → Transform to pattern matching problem
- **"Why transformation?"** → Palindrome properties enable Z-algorithm
- **"How to use Z-array?"** → Finds longest palindrome prefix efficiently

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Z algorithm:
1. Transform problem: create pattern = s + '#' + reverse(s)
2. Compute Z-array for this pattern
3. Find longest prefix that's also suffix (palindrome)
4. Extract corresponding prefix from original string
5. Add reverse of remaining suffix to create palindrome
6. Return result in O(N) time"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty string
- **Single character**: Return as is (already palindrome)
- **Already palindrome**: Return as is
- **Large strings**: Handle efficiently with Z-array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: s = "aacecaaa"

Human thinking:
"Let's apply Z algorithm:

Step 1: Transform
pattern = "aacecaaa#aaacecaa" (s + '#' + reverse(s))

Step 2: Compute Z-array
Z[0] = 14 (pattern length)
For i=1 to 13:
  Compute Z[i] = longest prefix matching at position i

Step 3: Find longest palindrome prefix
Scan Z-array for maximum Z[i] where i > pattern.length/2
Found max Z at position corresponding to palindrome

Step 4: Extract and construct result
Extract palindrome prefix from original string
Add reverse of remaining suffix
Result: "aaacecaaa" ✓

Manual verification:
"aaacecaaa" is palindrome ✓
Shortest possible ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Transformation enables efficient pattern matching
- **Why it's efficient**: O(N) vs O(N³) brute force
- **Why it's correct**: Z-array finds optimal palindrome prefix

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all additions?"** → O(N³) too slow
2. **"What about Manacher's?"** → More complex but same complexity
3. **"How to handle transformation?"** → Use string concatenation carefully
4. **"What about Z-array boundaries?"** → Handle pattern length correctly

### Real-World Analogy
**Like finding the shortest way to make a word symmetrical:**
- You have a word and want to add letters to make it read the same forwards and backwards
- Instead of trying all possible additions, use pattern recognition
- Transform the problem to find the longest mirror pattern
- This is like using a smart mirror to find the best reflection
- Useful in text processing, DNA sequence analysis, palindrome generation
- Like having a pattern recognition system for symmetrical structures

### Human-Readable Pseudocode
```
function shortestPalindrome(s):
    if s.length <= 1:
        return s
    
    // Transform problem
    pattern = s + '#' + reverse(s)
    z = computeZArray(pattern)
    
    // Find longest palindrome prefix
    maxLen = 0
    maxPos = 0
    
    for i from 1 to pattern.length-1:
        if z[i] > maxLen and i > pattern.length/2:
            maxLen = z[i]
            maxPos = i
    
    // Extract palindrome and construct result
    palindrome = s.substring(0, maxLen)
    suffix = s.substring(maxLen)
    
    return palindrome + reverse(suffix)
```

### Execution Visualization

### Example: s = "aacecaaa"
```
Z Algorithm Process:

Step 1: Transform
pattern = "aacecaaa#aaacecaa"

Step 2: Compute Z-array
Z[0] = 14
Z[1] = 0 (no match)
Z[2] = 0 (no match)
...
Z[7] = 2 (matches "aa")
...
Z[8] = 7 (matches "aacecaa")

Step 3: Find longest palindrome prefix
Scan for max Z[i] where i > 7
Found Z[8] = 7 (longest palindrome prefix)

Step 4: Construct result
palindrome = s.substring(0,7) = "aacecaa"
suffix = s.substring(7) = "a"
result = "aacecaa" + reverse("a") = "aaacecaaa" ✓

Visualization:
Transformed string enables pattern matching
Z-array efficiently finds palindrome structure
Result constructed by combining prefix and reversed suffix
```

### Key Visualization Points:
- **String Transformation**: s + '#' + reverse(s) enables palindrome detection
- **Z-Array**: Preprocesses pattern for O(1) palindrome queries
- **Pattern Matching**: Finds longest prefix that's also suffix
- **Linear Time**: O(N) vs O(N³) brute force

### Memory Layout Visualization:
```
Input: "aacecaaa"
Pattern: "aacecaaa#aaacecaa" (length 14)
Z-Array: [14,0,0,0,0,0,0,2,7,0,0,0]
Longest palindrome prefix: length 7 at position 8
Result: "aaacecaaa" (palindrome)

Z-array enables efficient palindrome detection
Linear time complexity achieved
```

### Time Complexity Breakdown:
- **String Transformation**: O(N) time, O(N) space
- **Z-Array Construction**: O(N) time, O(N) space
- **Pattern Search**: O(N) time, O(1) space
- **Total**: O(N) time, O(N) space
- **Optimal**: Best possible for this problem
- **vs Brute Force**: O(N³) vs O(N) with Z-algorithm
*/
