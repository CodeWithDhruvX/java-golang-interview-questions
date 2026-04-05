import java.util.*;

public class GroupAnagrams {
    
    // 49. Group Anagrams
    // Time: O(N*K*logK), Space: O(N*K)
    public static List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> anagramMap = new HashMap<>();
        
        for (String str : strs) {
            // Sort the string to create a key
            String sortedStr = sortString(str);
            anagramMap.computeIfAbsent(sortedStr, k -> new ArrayList<>()).add(str);
        }
        
        return new ArrayList<>(anagramMap.values());
    }

    // Helper function to sort a string
    private static String sortString(String s) {
        char[] chars = s.toCharArray();
        Arrays.sort(chars);
        return new String(chars);
    }

    // Alternative solution using character count as key (more efficient)
    public static List<List<String>> groupAnagramsOptimized(String[] strs) {
        Map<String, List<String>> anagramMap = new HashMap<>();
        
        for (String str : strs) {
            // Create key based on character count
            String key = createCountKey(str);
            anagramMap.computeIfAbsent(key, k -> new ArrayList<>()).add(str);
        }
        
        return new ArrayList<>(anagramMap.values());
    }

    // Create key based on character count (26 lowercase letters)
    private static String createCountKey(String s) {
        int[] count = new int[26];
        for (char c : s.toCharArray()) {
            count[c - 'a']++;
        }
        
        return Arrays.toString(count);
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"eat", "tea", "tan", "ate", "nat", "bat"},
            {""},
            {"a"},
            {"abc", "bca", "cab", "def", "fed", "ghi"},
            {"", "", ""},
            {"a", "b", "c", "a", "b", "c"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<List<String>> result = groupAnagrams(testCases[i]);
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("Grouped Anagrams: %s\n\n", result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hash Map with Canonical Key
- **Canonical Key**: Sorted string or character count as identifier
- **Grouping**: All strings with same key belong to same group
- **Hash Map**: Key → List of strings mapping
- **Two Approaches**: Sort-based vs Count-based keys

## 2. PROBLEM CHARACTERISTICS
- **String Array**: Input of lowercase strings
- **Anagram Grouping**: Group strings with same characters
- **Order Independence**: Character order doesn't matter for grouping
- **Return Groups**: Collection of all anagram groups

## 3. SIMILAR PROBLEMS
- Valid Anagram (check if two strings are anagrams)
- Group Anagrams with different constraints
- Find all anagrams in string
- Palindrome grouping

## 4. KEY OBSERVATIONS
- Anagrams have same character frequency
- Sorted version of anagrams are identical
- Need a canonical representation for grouping
- Character count can serve as unique key

## 5. VARIATIONS & EXTENSIONS
- Handle uppercase and lowercase
- Include non-alphabetic characters
- Return groups in specific order
- Optimize for memory constraints

## 6. INTERVIEW INSIGHTS
- Clarify: "Are strings only lowercase letters?"
- Clarify: "Does output order matter?"
- Edge cases: empty strings, single characters, duplicates
- Time-space tradeoff: sorting vs counting

## 7. COMMON MISTAKES
- Using string as key without proper canonicalization
- Not handling empty strings correctly
- Inefficient key generation
- Forgetting to handle duplicates

## 8. OPTIMIZATION STRATEGIES
- Count-based approach: O(N*K) vs O(N*K*logK)
- Use fixed-size array for character counts
- Optimize hash function for count arrays
- Pre-allocate hash map size

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like sorting books by author:**
- You have a pile of books (strings)
- Books by same author (anagrams) should go together
- You create a system to identify books by author
- You sort all books and place them in author groups

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of strings
2. **Goal**: Group strings that are anagrams
3. **Output**: Collection of anagram groups

#### Phase 2: Key Insight Recognition
- **"Anagrams have same letters!"** → Just rearranged
- **"How to identify same letters?"** → Sort them!
- **"Sorted strings become identical!"** → Perfect key
- **"Can I do better?"** → Count characters instead

#### Phase 3: Strategy Development
```
Human thought process:
"I need a way to identify if two strings are anagrams.
If I sort both strings, they'll look identical if they're anagrams.
So I can use the sorted string as a key to group them.
Or even better, I can count characters - that's faster!"
```

#### Phase 4: Edge Case Handling
- **Empty strings**: All empty strings are anagrams
- **Single characters**: Group by character
- **Duplicates**: Handle correctly in groups
- **Mixed lengths**: Different lengths can't be anagrams

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: ["eat", "tea", "tan", "ate", "nat", "bat"]

Human thinking:
"Let me process each string:
'eat' → sort → 'aet' → create group ['eat']
'tea' → sort → 'aet' → add to existing group ['eat', 'tea']
'tan' → sort → 'ant' → create new group ['tan']
'ate' → sort → 'aet' → add to first group ['eat', 'tea', 'ate']
'nat' → sort → 'ant' → add to second group ['tan', 'nat']
'bat' → sort → 'abt' → create new group ['bat']

Done! I have three groups."
```

#### Phase 6: Intuition Validation
- **Why it works**: Sorted strings are identical for anagrams
- **Why it's efficient**: Hash map provides O(1) grouping
- **Why it's correct**: Every string goes to exactly one group

### Common Human Pitfalls & How to Avoid Them
1. **"Should I use sorting or counting?"** → Counting is faster, sorting is simpler
2. **"What about case sensitivity?"** → Clarify input constraints
3. **"Can I optimize more?"** → Count-based approach is optimal
4. **"What about memory?"** → Use primitive arrays for counts

### Real-World Analogy
**Like organizing a library by genre:**
- You have books with different titles but same content type
- You create a cataloging system (canonical key)
- Books with same content type go to same shelf
- Each shelf represents one group (anagram group)

### Human-Readable Pseudocode
```
function groupAnagrams(strings):
    groups = empty hash map
    
    for each string in strings:
        key = createCanonicalKey(string)
        if key not in groups:
            groups[key] = empty list
        groups[key].add(string)
    
    return all values from groups
```

### Execution Visualization

### Example: ["eat", "tea", "tan", "ate", "nat", "bat"]
```
Initial: strings = ["eat", "tea", "tan", "ate", "nat", "bat"], groups = {}

Step 1: "eat" → sort → "aet"
→ groups = {"aet": ["eat"]}

Step 2: "tea" → sort → "aet"
→ groups = {"aet": ["eat", "tea"]}

Step 3: "tan" → sort → "ant"
→ groups = {"aet": ["eat", "tea"], "ant": ["tan"]}

Step 4: "ate" → sort → "aet"
→ groups = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}

Step 5: "nat" → sort → "ant"
→ groups = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}

Step 6: "bat" → sort → "abt"
→ groups = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}

Final: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

### Key Visualization Points:
- **Canonical key** identifies anagram groups
- **Hash map** automatically groups by key
- **Sorting** creates identical keys for anagrams
- **Groups** contain all strings with same key

### Memory Layout Visualization:
```
Input:  ["eat", "tea", "tan", "ate", "nat", "bat"]
Keys:    ["aet", "aet", "ant", "aet", "ant", "abt"]
Groups:  {"aet": ["eat", "tea", "ate"], 
          "ant": ["tan", "nat"], 
          "abt": ["bat"]}
```

### Time Complexity Breakdown:
- **Sort-based**: O(N*K*logK) where K is average string length
- **Count-based**: O(N*K) - more efficient
- **Space**: O(N*K) - storing all strings plus hash map
- **Optimal**: Count-based approach is best for lowercase letters
*/
