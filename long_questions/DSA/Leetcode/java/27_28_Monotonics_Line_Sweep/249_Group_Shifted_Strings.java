import java.util.*;

public class GroupShiftedStrings {
    
    // 249. Group Shifted Strings - Hashing with String Processing
    // Time: O(N * L), Space: O(N * L) where N is number of strings, L is average length
    public List<List<String>> groupStrings(String[] strings) {
        if (strings.length == 0) {
            return new ArrayList<>();
        }
        
        // Map to store groups by key
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strings) {
            String key = getShiftKey(s);
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
        }
        
        // Convert map to list
        return new ArrayList<>(groups.values());
    }
    
    // Get shift key for a string
    private String getShiftKey(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        // Calculate shift from first character
        int shift = s.charAt(0) - 'a';
        
        char[] key = new char[s.length()];
        key[0] = 'a'; // Normalize first character to 'a'
        
        for (int i = 1; i < s.length(); i++) {
            // Calculate normalized character
            char normalized = (char) ((s.charAt(i) - 'a' - shift + 26) % 26 + 'a');
            key[i] = normalized;
        }
        
        return new String(key);
    }
    
    // Alternative approach using differences
    public List<List<String>> groupStringsDifferences(String[] strings) {
        if (strings.length == 0) {
            return new ArrayList<>();
        }
        
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strings) {
            String key = getDifferenceKey(s);
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
        }
        
        return new ArrayList<>(groups.values());
    }
    
    private String getDifferenceKey(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        char[] key = new char[s.length() - 1];
        
        for (int i = 1; i < s.length(); i++) {
            int diff = (s.charAt(i) - s.charAt(i - 1) + 26) % 26;
            key[i - 1] = (char) (diff + 'a'); // Convert to character
        }
        
        return new String(key);
    }
    
    // Optimized version with sorting
    public List<List<String>> groupStringsOptimized(String[] strings) {
        if (strings.length == 0) {
            return new ArrayList<>();
        }
        
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strings) {
            String key = getShiftKeyOptimized(s);
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
        }
        
        return new ArrayList<>(groups.values());
    }
    
    private String getShiftKeyOptimized(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        int shift = s.charAt(0) - 'a';
        char[] key = new char[s.length()];
        
        for (int i = 0; i < s.length(); i++) {
            key[i] = (char) ((s.charAt(i) - 'a' - shift + 26) % 26 + 'a');
        }
        
        return new String(key);
    }
    
    // Version that handles all characters (not just lowercase)
    public List<List<String>> groupStringsAllChars(String[] strings) {
        if (strings.length == 0) {
            return new ArrayList<>();
        }
        
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strings) {
            String key = getUniversalShiftKey(s);
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
        }
        
        return new ArrayList<>(groups.values());
    }
    
    private String getUniversalShiftKey(String s) {
        if (s.length() <= 1) {
            return s;
        }
        
        // Find the minimum character to normalize
        char minChar = s.charAt(0);
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) < minChar) {
                minChar = s.charAt(i);
            }
        }
        
        int shift = minChar - 'a';
        char[] key = new char[s.length()];
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c >= 'a' && c <= 'z') {
                key[i] = (char) ((c - 'a' - shift + 26) % 26 + 'a');
            } else if (c >= 'A' && c <= 'Z') {
                key[i] = (char) ((c - 'A' - shift + 26) % 26 + 'a');
            } else {
                key[i] = c; // Keep non-alphabetic characters as is
            }
        }
        
        return new String(key);
    }
    
    // Brute force approach for comparison
    public List<List<String>> groupStringsBruteForce(String[] strings) {
        if (strings.length == 0) {
            return new ArrayList<>();
        }
        
        boolean[] used = new boolean[strings.length];
        List<List<String>> result = new ArrayList<>();
        
        for (int i = 0; i < strings.length; i++) {
            if (used[i]) {
                continue;
            }
            
            List<String> group = new ArrayList<>();
            group.add(strings[i]);
            used[i] = true;
            
            // Find all strings that can be shifted to match strings[i]
            for (int j = i + 1; j < strings.length; j++) {
                if (!used[j] && canShift(strings[i], strings[j])) {
                    group.add(strings[j]);
                    used[j] = true;
                }
            }
            
            result.add(group);
        }
        
        return result;
    }
    
    private boolean canShift(String s1, String s2) {
        if (s1.length() != s2.length()) {
            return false;
        }
        
        if (s1.length() == 1) {
            return true;
        }
        
        int shift = (s2.charAt(0) - s1.charAt(0) + 26) % 26;
        
        for (int i = 1; i < s1.length(); i++) {
            if ((s2.charAt(i) - s1.charAt(i) + 26) % 26 != shift) {
                return false;
            }
        }
        
        return true;
    }
    
    // Version with detailed explanation
    public class GroupResult {
        List<List<String>> groups;
        List<String> explanation;
        
        GroupResult(List<List<String>> groups, List<String> explanation) {
            this.groups = groups;
            this.explanation = explanation;
        }
    }
    
    public GroupResult groupStringsWithExplanation(String[] strings) {
        List<String> explanation = new ArrayList<>();
        
        if (strings.length == 0) {
            explanation.add("Empty input, returning empty result");
            return new GroupResult(new ArrayList<>(), explanation);
        }
        
        explanation.add(String.format("Processing %d strings", strings.length));
        
        Map<String, List<String>> groups = new HashMap<>();
        
        for (int i = 0; i < strings.length; i++) {
            String s = strings[i];
            explanation.add(String.format("String %d: '%s'", i, s));
            
            String key = getShiftKey(s);
            explanation.add(String.format("  Shift key: '%s'", key));
            
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
            explanation.add(String.format("  Added to group with key '%s'", key));
        }
        
        explanation.add(String.format("Found %d groups", groups.size()));
        
        List<List<String>> result = new ArrayList<>(groups.values());
        for (Map.Entry<String, List<String>> entry : groups.entrySet()) {
            explanation.add(String.format("Group '%s': %s", entry.getKey(), entry.getValue()));
        }
        
        return new GroupResult(result, explanation);
    }
    
    // Helper function to check if two strings are in the same group
    public boolean areInSameGroup(String s1, String s2) {
        return getShiftKey(s1).equals(getShiftKey(s2));
    }
    
    // Find all possible shifts of a string
    public List<String> getAllShifts(String s) {
        if (s.length() <= 1) {
            return Arrays.asList(s);
        }
        
        List<String> shifts = new ArrayList<>();
        
        for (int shift = 0; shift < 26; shift++) {
            char[] shifted = new char[s.length()];
            
            for (int i = 0; i < s.length(); i++) {
                shifted[i] = (char) ((s.charAt(i) - 'a' + shift) % 26 + 'a');
            }
            
            shifts.add(new String(shifted));
        }
        
        return shifts;
    }
    
    public static void main(String[] args) {
        GroupShiftedStrings gss = new GroupShiftedStrings();
        
        // Test cases
        System.out.println("=== Testing Group Shifted Strings ===");
        
        String[][] testCases = {
            {"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
            {"a"},
            {"abc", "def", "ghi"},
            {"a", "b", "c", "d", "e"},
            {"abc", "bcd", "cde", "def"},
            {"az", "ba", "ab", "bc"},
            {},
            {"abc", "abc", "abc"},
            {"abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"},
            {"az", "by", "cx", "dw"}
        };
        
        String[] descriptions = {
            "Standard case",
            "Single character",
            "All same length",
            "All single characters",
            "Sequential shifts",
            "Mixed shifts",
            "Empty array",
            "Duplicate strings",
            "Full alphabet",
            "Different patterns"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("  Input: %s\n", Arrays.toString(testCases[i]));
            
            List<List<String>> result1 = gss.groupStrings(testCases[i]);
            List<List<String>> result2 = gss.groupStringsDifferences(testCases[i]);
            List<List<String>> result3 = gss.groupStringsOptimized(testCases[i]);
            List<List<String>> result4 = gss.groupStringsBruteForce(testCases[i]);
            
            System.out.printf("  Standard: %s\n", result1);
            System.out.printf("  Differences: %s\n", result2);
            System.out.printf("  Optimized: %s\n", result3);
            System.out.printf("  Brute Force: %s\n\n", result4);
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        String[] testStrings = {"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"};
        GroupResult result = gss.groupStringsWithExplanation(testStrings);
        
        System.out.printf("Result: %s\n", result.groups);
        for (String step : result.explanation) {
            System.out.printf("  %s\n", step);
        }
        
        // Test helper functions
        System.out.println("\n=== Helper Functions Test ===");
        
        System.out.printf("Are 'abc' and 'bcd' in same group? %b\n", gss.areInSameGroup("abc", "bcd"));
        System.out.printf("Are 'abc' and 'xyz' in same group? %b\n", gss.areInSameGroup("abc", "xyz"));
        
        System.out.printf("All shifts of 'abc': %s\n", gss.getAllShifts("abc"));
        
        // Performance test
        System.out.println("\n=== Performance Test ===");
        
        String[] largeStrings = new String[1000];
        for (int i = 0; i < 1000; i++) {
            String base = "abc";
            int shift = i % 26;
            char[] shifted = new char[base.length()];
            for (int j = 0; j < base.length(); j++) {
                shifted[j] = (char) ((base.charAt(j) - 'a' + shift) % 26 + 'a');
            }
            largeStrings[i] = new String(shifted);
        }
        
        System.out.printf("Large test with %d strings\n", largeStrings.length);
        
        List<List<String>> largeResult = gss.groupStrings(largeStrings);
        System.out.printf("Found %d groups\n", largeResult.size());
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Mixed case strings
        String[] mixedCase = {"ABC", "BCD", "CDE", "abc", "bcd"};
        System.out.printf("Mixed case: %s\n", gss.groupStringsAllChars(mixedCase));
        
        // Strings with numbers
        String[] withNumbers = {"a1b", "b1c", "c1d"};
        System.out.printf("With numbers: %s\n", gss.groupStrings(withNumbers));
        
        // Very long strings
        String[] longStrings = {
            "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
            "bcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyza"
        };
        System.out.printf("Long strings: %s\n", gss.groupStrings(longStrings));
        
        // Test with all possible single characters
        System.out.println("\n=== All Single Characters ===");
        String[] allChars = new String[26];
        for (int i = 0; i < 26; i++) {
            allChars[i] = String.valueOf((char) ('a' + i));
        }
        
        List<List<String>> allCharsResult = gss.groupStrings(allChars);
        System.out.printf("All single characters grouped: %s\n", allCharsResult);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: String Grouping with Hashing
- **Shift Detection**: Identify character shift patterns
- **Normalization**: Convert strings to canonical form
- **Hash-based Grouping**: Group by normalized representation
- **Character Analysis**: Process string character relationships

## 2. PROBLEM CHARACTERISTICS
- **String Shifts**: Each string can be shifted by moving characters
- **Cyclic Shift**: Character 'z' shifts to 'a' (wrap-around)
- **Grouping Logic**: Strings in same group can be shifted to match
- **Hash Key**: Normalized representation for grouping

## 3. SIMILAR PROBLEMS
- Group Anagrams
- Find All Shifted Strings
- String Pattern Matching
- Cryptographic Analysis

## 4. KEY OBSERVATIONS
- Shift key can be derived from character differences
- Normalization removes shift dependency
- Hash groups strings with same shift pattern
- Multiple approaches: shift-based, difference-based, brute force
- Time complexity: O(N * L) where N is strings, L is average length

## 5. VARIATIONS & EXTENSIONS
- Different character sets (uppercase, mixed case)
- Non-alphabetic characters
- Variable shift amounts
- Unicode character support

## 6. INTERVIEW INSIGHTS
- Clarify: "What character set should I support?"
- Edge cases: empty array, single character, mixed case
- Time complexity: O(N * L) vs O(N² * L²) brute force
- Space complexity: O(N * L) for hash storage

## 7. COMMON MISTAKES
- Incorrect modulo arithmetic for cyclic shift
- Not handling single character strings
- Forgetting to normalize properly
- Case sensitivity issues
- Not handling non-alphabetic characters

## 8. OPTIMIZATION STRATEGIES
- Pre-compute character mappings
- Use efficient string building
- Handle edge cases early
- Choose appropriate normalization method

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing books by their first letter:**
- You have books (strings) that might be shifted versions
- Each book's first letter tells you the shift amount
- Normalize all books to start with 'a' to compare patterns
- Books with same normalized pattern belong to same group
- This is like organizing books by their "alphabetical fingerprint"

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of strings
2. **Goal**: Group strings that can be shifted to match
3. **Output**: Groups of shift-equivalent strings

#### Phase 2: Key Insight Recognition
- **"What defines shift equivalence?"** → Same character differences
- **"How to normalize?"** → Shift first character to 'a'
- **"What's the pattern?"** → Character relationships preserved
- **"Why hashing?"** → Efficient grouping by key

#### Phase 3: Strategy Development
```
Human thought process:
"I'll normalize strings:
1. For each string, calculate shift from first character
2. Apply inverse shift to normalize first character to 'a'
3. Apply same shift to all characters
4. Use normalized string as hash key
5. Group strings by their normalized keys"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty list
- **Single character**: Always groups with itself
- **Mixed case**: Handle separately or normalize
- **Non-alphabetic**: Keep as-is in normalization

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Strings: ["abc", "bcd", "acef", "xyz"]

Human thinking:
"Let's normalize each string:

String "abc":
- First char 'a', shift = 'a' - 'a' = 0
- Normalize: apply shift 0 → "abc"
- Key: "abc"

String "bcd":
- First char 'b', shift = 'b' - 'a' = 1
- Normalize: apply inverse shift (-1):
  - 'b' - 1 = 'a'
  - 'c' - 1 = 'b'
  - 'd' - 1 = 'c'
- Result: "abc"
- Key: "abc"

String "acef":
- First char 'a', shift = 'a' - 'a' = 0
- Normalize: apply shift 0 → "acef"
- Key: "acef"

String "xyz":
- First char 'x', shift = 'x' - 'a' = 23
- Normalize: apply inverse shift (-23):
  - 'x' - 23 = 'a'
  - 'y' - 23 = 'b'
  - 'z' - 23 = 'c'
- Result: "abc"
- Key: "abc"

Grouping:
Key "abc": ["abc", "bcd", "xyz"]
Key "acef": ["acef"]

Result: [["abc", "bcd", "xyz"], ["acef"]] ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Normalization removes shift dependency
- **Why it's efficient**: Each string processed once
- **Why it's correct**: Same normalized form means shift-equivalent

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compare all pairs?"** → O(N² * L²) brute force
2. **"What about case sensitivity?"** → Need consistent handling
3. **"How to handle wrap-around?"** → Use modulo 26 arithmetic
4. **"What about non-alphabetic?"** → Keep unchanged in normalization

### Real-World Analogy
**Like organizing encrypted messages:**
- You have encrypted messages (shifted strings)
- Each message uses Caesar cipher with different shift
- To find messages using same cipher, you need the "key"
- Normalize all messages by removing their shift
- Messages with same normalized form use same cipher
- This groups messages by their encryption pattern
- Useful for cryptanalysis and message organization

### Human-Readable Pseudocode
```
function groupStrings(strings):
    groups = empty map
    
    for string in strings:
        if string.length <= 1:
            key = string
        else:
            shift = string[0] - 'a'
            key = normalizeString(string, -shift)
        
        groups[key].add(string)
    
    return list of groups.values()
    
function normalizeString(s, shift):
    result = empty string
    for character in s:
        if character is alphabetic:
            normalized = (character - 'a' + shift + 26) % 26 + 'a'
            result += normalized
        else:
            result += character
    return result
```

### Execution Visualization

### Example: ["abc", "bcd", "acef", "xyz"]
```
Normalization Process:

"abc":
- shift = 'a' - 'a' = 0
- normalize with shift 0: "abc"
- key: "abc"

"bcd":
- shift = 'b' - 'a' = 1
- normalize with shift -1:
  - 'b' - 1 = 'a'
  - 'c' - 1 = 'b'
  - 'd' - 1 = 'c'
- result: "abc"
- key: "abc"

"acef":
- shift = 'a' - 'a' = 0
- normalize with shift 0: "acef"
- key: "acef"

"xyz":
- shift = 'x' - 'a' = 23
- normalize with shift -23:
  - 'x' - 23 = 'a'
  - 'y' - 23 = 'b'
  - 'z' - 23 = 'c'
- result: "abc"
- key: "abc"

Grouping Result:
Key "abc": ["abc", "bcd", "xyz"]
Key "acef": ["acef"]

Final: [["abc", "bcd", "xyz"], ["acef"]] ✓

Visualization:
Original:  abc  bcd  acef  xyz
Shifts:    0    1     0     23
Normalized: abc  abc  acef  abc
Groups:    [abc,bcd,xyz] [acef]
```

### Key Visualization Points:
- **Shift calculation** from first character
- **Normalization** removes shift dependency
- **Hash grouping** by normalized form
- **Cyclic arithmetic** handles wrap-around

### Memory Layout Visualization:
```
Normalization Process:
String: "bcd"
Original: b(98) c(99) d(100)
Shift: 1 (b-a)
Normalized: b-1=97(a) c-1=98(b) d-1=99(c)
Result: "abc"

Hash Map Structure:
Key "abc" → ["abc", "bcd", "xyz"]
Key "acef" → ["acef"]

Processing Flow:
For each string:
1. Calculate shift
2. Apply inverse shift to normalize
3. Use normalized string as hash key
4. Add to appropriate group
```

### Time Complexity Breakdown:
- **Each String**: O(L) where L is string length
- **All Strings**: O(N * L) where N is number of strings
- **Hash Operations**: O(1) average for grouping
- **Total**: O(N * L) time, O(N * L) space
- **Optimal**: Cannot do better than O(N * L) for this problem
- **vs Brute Force**: O(N² * L²) checking all pairs
*/
