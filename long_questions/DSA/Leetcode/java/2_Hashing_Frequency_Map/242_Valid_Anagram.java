import java.util.HashMap;
import java.util.Map;

public class ValidAnagram {
    
    // 242. Valid Anagram
    // Time: O(N), Space: O(1) for ASCII characters (26 for lowercase letters)
    public static boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        
        // Assuming only lowercase English letters
        int[] count = new int[26];
        
        for (int i = 0; i < s.length(); i++) {
            count[s.charAt(i) - 'a']++;
            count[t.charAt(i) - 'a']--;
        }
        
        for (int c : count) {
            if (c != 0) {
                return false;
            }
        }
        
        return true;
    }

    // Alternative solution using frequency map for general characters
    public static boolean isAnagramGeneral(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        
        Map<Character, Integer> count = new HashMap<>();
        
        for (char c : s.toCharArray()) {
            count.put(c, count.getOrDefault(c, 0) + 1);
        }
        
        for (char c : t.toCharArray()) {
            int newCount = count.getOrDefault(c, 0) - 1;
            if (newCount < 0) {
                return false;
            }
            count.put(c, newCount);
        }
        
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"anagram", "nagaram"},
            {"rat", "car"},
            {"a", "a"},
            {"ab", "ba"},
            {"", ""},
            {"abc", "ab"},
            {"Hello", "olleH"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = isAnagram(testCases[i][0], testCases[i][1]);
            boolean resultGeneral = isAnagramGeneral(testCases[i][0], testCases[i][1]);
            System.out.printf("Test Case %d: \"%s\" & \"%s\" -> Anagram: %b (General: %b)\n", 
                i + 1, testCases[i][0], testCases[i][1], result, resultGeneral);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Character Frequency Comparison
- **Frequency Array**: Count characters in both strings
- **Single Pass**: Increment for first string, decrement for second
- **Zero Check**: All counts must be zero for anagram
- **Optimized**: O(1) space for fixed character set

## 2. PROBLEM CHARACTERISTICS
- **Two Strings**: Compare character composition
- **Anagram Check**: Same characters with same frequencies
- **Case Sensitivity**: Usually case-sensitive
- **Length Check**: Different lengths cannot be anagrams

## 3. SIMILAR PROBLEMS
- Group Anagrams
- Find All Anagrams in String
- Valid Palindrome
- String Permutation

## 4. KEY OBSERVATIONS
- Anagrams have identical character frequencies
- Length mismatch means not anagrams (early exit)
- Can use frequency array for lowercase letters
- Hash map works for general character sets

## 5. VARIATIONS & EXTENSIONS
- Handle Unicode characters
- Ignore case and non-alphabetic characters
- Return character difference details
- Optimize for specific character ranges

## 6. INTERVIEW INSIGHTS
- Clarify: "Are strings case-sensitive?"
- Clarify: "What character set?" (ASCII, Unicode)
- Edge cases: empty strings, single characters
- Space-time tradeoff: frequency array vs hash map

## 7. COMMON MISTAKES
- Not checking length first (wasted work)
- Using sorting when frequency counting is better
- Mishandling character encoding
- Not resetting frequency array properly

## 8. OPTIMIZATION STRATEGIES
- Frequency array: O(1) space for lowercase letters
- Single pass: combine counting for both strings
- Early exit: check length first
- Bit manipulation: for very limited character sets

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting inventory:**
- You have two warehouses (strings) with items (characters)
- You want to know if they have exactly the same inventory
- You count items in first warehouse, then subtract items from second
- If all counts end up at zero, inventories match!

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two strings s and t
2. **Goal**: Check if they are anagrams
3. **Output**: True if anagrams, false otherwise

#### Phase 2: Key Insight Recognition
- **"Anagrams have same letters!"** → Same character counts
- **"How to compare efficiently?"** → Count and compare
- **"Can I do it in one pass?"** → Yes! Increment and decrement
- **"What about length?"** → Must be equal (early check)

#### Phase 3: Strategy Development
```
Human thought process:
"If strings have different lengths, they can't be anagrams.
Otherwise, I'll count characters in both strings at once:
- Add 1 for each character in first string
- Subtract 1 for each character in second string
If all counts end up zero, they're perfect anagrams!"
```

#### Phase 4: Edge Case Handling
- **Different lengths**: Return false immediately
- **Empty strings**: Both empty → true, one empty → false
- **Single character**: Must be identical
- **Case sensitivity**: Depends on requirements

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: s = "anagram", t = "nagaram"

Human thinking:
"Lengths are equal (7), so continue.
I'll create an array of 26 counters (one for each letter).

For each position:
i=0: s[0]='a' → count[0]++, t[0]='n' → count[13]--
i=1: s[1]='n' → count[13]++, t[1]='a' → count[0]--
... continue for all positions

At the end, I'll check if all 26 counters are zero.
If yes, they're anagrams!"
```

#### Phase 6: Intuition Validation
- **Why it works**: Net count of zero means identical frequencies
- **Why it's efficient**: Single pass, O(1) space for lowercase
- **Why it's correct**: Every character difference is captured

### Common Human Pitfalls & How to Avoid Them
1. **"Should I sort first?"** → Frequency counting is faster
2. **"What about Unicode?"** → Use hash map for general case
3. **"Can I optimize more?"** → Single pass is optimal
4. **"What about case?"** → Clarify requirements first

### Real-World Analogy
**Like balancing a checkbook:**
- You have income (first string) and expenses (second string)
- You track each transaction type (character)
- Income adds to balance, expenses subtract from balance
- If final balance is zero for all categories, books are balanced!

### Human-Readable Pseudocode
```
function isAnagram(string1, string2):
    if lengths differ:
        return false
    
    counts = array of zeros (size depends on character set)
    
    for i from 0 to length-1:
        counts[string1[i]]++
        counts[string2[i]]--
    
    for each count in counts:
        if count != 0:
            return false
    
    return true
```

### Execution Visualization

### Example: s = "anagram", t = "nagaram"
```
Initial: s = "anagram", t = "nagaram", counts = [0,0,0,...,0]

Step 1: i=0, s[0]='a', t[0]='n'
→ counts[0]++ (for 'a'), counts[13]-- (for 'n')
→ counts[0]=1, counts[13]=-1

Step 2: i=1, s[1]='n', t[1]='a'
→ counts[13]++ (for 'n'), counts[0]-- (for 'a')
→ counts[13]=0, counts[0]=0

... continue for all positions ...

Final: All counts = [0,0,0,...,0]
→ Return true (all zeros)
```

### Key Visualization Points:
- **Frequency array** tracks net character count
- **Increment/decrement** happens in single pass
- **Zero check** confirms identical frequencies
- **Early exit** on length mismatch

### Memory Layout Visualization:
```
String s:  a n a g r a m
String t:  n a g a r a m
Counts:   [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
           a b c d e f g h i j k l m n o p q r s t u v w x y z
           
After processing: All counts = 0 ✓
```

### Time Complexity Breakdown:
- **Length Check**: O(1)
- **Counting Pass**: O(N) where N is string length
- **Verification Pass**: O(C) where C is character set size (26 for lowercase)
- **Total**: O(N + C) → O(N) since C is constant
- **Space**: O(C) → O(1) for fixed character set
*/
