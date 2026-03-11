# String Programs (Complete Collection - 55 Programs)

## 📚 Beginner Level (1-10) - Fundamentals

### 1. Count Occurrence of Each Character
**Principle**: Use frequency array or HashMap.
**Question**: Print count of each character in a string.

**Brute Force Approach (O(n²))**:
```java
public class CharFrequencyBruteForce {
    public static void main(String[] args) {
        String str = "programming";
        boolean[] counted = new boolean[str.length()];
        
        for (int i = 0; i < str.length(); i++) {
            if (!counted[i]) {
                char ch = str.charAt(i);
                int count = 1;
                
                for (int j = i + 1; j < str.length(); j++) {
                    if (ch == str.charAt(j)) {
                        count++;
                        counted[j] = true;
                    }
                }
                
                System.out.println(ch + "=" + count);
                counted[i] = true;
            }
        }
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class CharFrequency {
    public static void main(String[] args) {
        String str = "programming";
        Map<Character, Integer> map = new LinkedHashMap<>();
        for (char ch : str.toCharArray()) {
            map.put(ch, map.getOrDefault(ch, 0) + 1);
        }
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + "=" + entry.getValue());
        }
    }
}
```

### 2. Find Length Without length()
**Principle**: Convert to char array and count.
**Question**: Find string length without using length() method.

**Brute Force Approach (O(n²))**:
```java
public class StringLengthBruteForce {
    public static void main(String[] args) {
        String str = "Hello";
        int count = 0;
        
        try {
            while (true) {
                str.charAt(count);
                count++;
            }
        } catch (StringIndexOutOfBoundsException e) {
            // Exception indicates end of string
        }
        
        System.out.println("Length: " + count);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class StringLength {
    public static void main(String[] args) {
        String str = "Hello";
        int count = 0;
        for(char c : str.toCharArray()) count++;
        System.out.println("Length: " + count);
    }
}
```

## 📚 Intermediate Level (11-20)

### 3. First Repeated Character
**Principle**: Use HashSet to track seen characters.
**Question**: Find first repeated character in a string.

**Brute Force Approach (O(n²))**:
```java
public class FirstRepeatedBruteForce {m
    public static void main(String[] args) {
        String str = "swiss";
        
        for (int i = 0; i < str.length(); i++) {
            for (int j = i + 1; j < str.length(); j++) {
                if (str.charAt(i) == str.charAt(j)) {
                    System.out.println("First repeated: " + str.charAt(i));
                    return;
                }
            }
        }
        System.out.println("No repeated character");
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class FirstRepeated {
    public static void main(String[] args) {
        String str = "swiss";
        Set<Character> seen = new HashSet<>();
        for (char ch : str.toCharArray()) {
            if (seen.contains(ch)) {
                System.out.println("First repeated: " + ch);
                return;
            }
            seen.add(ch);
        }
        System.out.println("No repeated character");
    }
}
```

### 4. Remove Duplicate Characters
**Principle**: Use StringBuilder with visited check.
**Question**: Remove duplicate characters from a string.

**Brute Force Approach (O(n²))**:
```java
public class RemoveDuplicatesBruteForce {
    public static void main(String[] args) {
        String str = "programming";
        String result = "";
        
        for (int i = 0; i < str.length(); i++) {
            boolean isDuplicate = false;
            
            for (int j = 0; j < result.length(); j++) {
                if (str.charAt(i) == result.charAt(j)) {
                    isDuplicate = true;
                    break;
                }
            }
            
            if (!isDuplicate) {
                result += str.charAt(i);
            }
        }
        
        System.out.println("After removing duplicates: " + result);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class RemoveDuplicates {
    public static void main(String[] args) {
        String str = "programming";
        StringBuilder result = new StringBuilder();
        boolean[] visited = new boolean[256];
        
        for (char ch : str.toCharArray()) {
            if (!visited[ch]) {
                visited[ch] = true;
                result.append(ch);
            }
        }
        System.out.println("After removing duplicates: " + result);
    }
}
```

## 📚 Advanced Level (21-30) - Very Common in Interviews

### 5. Longest Substring Without Repeating Characters
**Principle**: Sliding window with HashSet.
**Question**: Find length of longest substring without repeating characters.

**Brute Force Approach (O(n³))**:
```java
public class LongestUniqueSubstringBruteForce {
    public static void main(String[] args) {
        String str = "abcabcbb";
        int maxLen = 0;
        
        for (int i = 0; i < str.length(); i++) {
            for (int j = i; j < str.length(); j++) {
                if (hasUniqueChars(str, i, j)) {
                    maxLen = Math.max(maxLen, j - i + 1);
                }
            }
        }
        System.out.println("Longest substring length: " + maxLen);
    }
    
    static boolean hasUniqueChars(String str, int start, int end) {
        Set<Character> seen = new HashSet<>();
        for (int i = start; i <= end; i++) {
            char ch = str.charAt(i);
            if (seen.contains(ch)) return false;
            seen.add(ch);
        }
        return true;
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class LongestUniqueSubstring {
    public static void main(String[] args) {
        String str = "abcabcbb";
        int maxLen = 0, start = 0;
        Map<Character, Integer> map = new HashMap<>();
        
        for (int i = 0; i < str.length(); i++) {
            char ch = str.charAt(i);
            if (map.containsKey(ch) && map.get(ch) >= start) {
                start = map.get(ch) + 1;
            }
            map.put(ch, i);
            maxLen = Math.max(maxLen, i - start + 1);
        }
        System.out.println("Longest substring length: " + maxLen);
    }
}
```

### 6. String Compression
**Principle**: Count consecutive characters and build compressed string.
**Question**: Compress string like aaabbc → a3b2c1.

**Brute Force Approach (O(n²))**:
```java
public class StringCompressionBruteForce {
    public static void main(String[] args) {
        String str = "aaabbc";
        String compressed = "";
        int i = 0;
        
        while (i < str.length()) {
            char ch = str.charAt(i);
            int count = 1;
            
            for (int j = i + 1; j < str.length(); j++) {
                if (str.charAt(j) == ch) {
                    count++;
                } else {
                    break;
                }
            }
            
            compressed += ch + String.valueOf(count);
            i += count;
        }
        
        System.out.println("Compressed: " + compressed);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class StringCompression {
    public static void main(String[] args) {
        String str = "aaabbc";
        StringBuilder compressed = new StringBuilder();
        int count = 1;
        
        for (int i = 1; i < str.length(); i++) {
            if (str.charAt(i) == str.charAt(i-1)) {
                count++;
            } else {
                compressed.append(str.charAt(i-1)).append(count);
                count = 1;
            }
        }
        compressed.append(str.charAt(str.length()-1)).append(count);
        System.out.println("Compressed: " + compressed);
    }
}
```

### 7. Print All Substrings
**Principle**: Nested loops for start and end indices.
**Question**: Print all possible substrings of a string.

**Brute Force Approach (O(n³))**:
```java
public class AllSubstringsBruteForce {
    public static void main(String[] args) {
        String str = "ABC";
        
        for (int i = 0; i < str.length(); i++) {
            for (int j = i; j < str.length(); j++) {
                String substring = "";
                for (int k = i; k <= j; k++) {
                    substring += str.charAt(k);
                }
                System.out.println(substring);
            }
        }
    }
}
```

**Optimized Approach (O(n²))**:
```java
public class AllSubstrings {
    public static void main(String[] args) {
        String str = "ABC";
        for (int i = 0; i < str.length(); i++) {
            for (int j = i + 1; j <= str.length(); j++) {
                System.out.println(str.substring(i, j));
            }
        }
    }
}
```

### 8. Check Balanced Parentheses
**Principle**: Use Stack to track opening brackets.
**Question**: Check if parentheses are balanced.

**Brute Force Approach (O(n²))**:
```java
public class BalancedParenthesesBruteForce {
    public static void main(String[] args) {
        String str = "{[()]}";
        System.out.println("Balanced? " + isBalancedBruteForce(str));
    }
    
    static boolean isBalancedBruteForce(String str) {
        String current = str;
        
        while (current.contains("()") || current.contains("[]") || current.contains("{}")) {
            current = current.replace("()", "");
            current = current.replace("[]", "");
            current = current.replace("{}", "");
        }
        
        return current.isEmpty();
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class BalancedParentheses {
    public static void main(String[] args) {
        String str = "{[()]}";
        Stack<Character> stack = new Stack<>();
        
        for (char ch : str.toCharArray()) {
            if (ch == '(' || ch == '[' || ch == '{') {
                stack.push(ch);
            } else if (ch == ')' || ch == ']' || ch == '}') {
                if (stack.isEmpty() || !isMatching(stack.pop(), ch)) {
                    System.out.println("Not Balanced");
                    return;
                }
            }
        }
        System.out.println(stack.isEmpty() ? "Balanced" : "Not Balanced");
    }
    
    static boolean isMatching(char open, char close) {
        return (open == '(' && close == ')') ||
               (open == '[' && close == ']') ||
               (open == '{' && close == '}');
    }
}
```

### 9. Longest Palindrome Substring
**Principle**: Expand around center for each character.
**Question**: Find longest palindromic substring.

**Brute Force Approach (O(n³))**:
```java
public class LongestPalindromeBruteForce {
    public static void main(String[] args) {
        String str = "babad";
        String longest = "";
        
        for (int i = 0; i < str.length(); i++) {
            for (int j = i; j < str.length(); j++) {
                String substring = str.substring(i, j + 1);
                if (isPalindrome(substring) && substring.length() > longest.length()) {
                    longest = substring;
                }
            }
        }
        System.out.println("Longest palindrome: " + longest);
    }
    
    static boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) return false;
            left++;
            right--;
        }
        return true;
    }
}
```

**Optimized Approach (O(n²))**:
```java
public class LongestPalindrome {
    public static void main(String[] args) {
        String str = "babad";
        String longest = "";
        
        for (int i = 0; i < str.length(); i++) {
            String odd = expand(str, i, i);
            String even = expand(str, i, i + 1);
            
            if (odd.length() > longest.length()) longest = odd;
            if (even.length() > longest.length()) longest = even;
        }
        System.out.println("Longest palindrome: " + longest);
    }
    
    static String expand(String str, int left, int right) {
        while (left >= 0 && right < str.length() && str.charAt(left) == str.charAt(right)) {
            left--;
            right++;
        }
        return str.substring(left + 1, right);
    }
}
```

### 10. Remove Adjacent Duplicates
**Principle**: Use Stack to remove consecutive duplicates.
**Question**: Remove adjacent duplicate characters.

**Brute Force Approach (O(n²))**:
```java
public class RemoveAdjacentDuplicatesBruteForce {
    public static void main(String[] args) {
        String str = "abbaca";
        String result = str;
        boolean changed;
        
        do {
            changed = false;
            String temp = "";
            int i = 0;
            
            while (i < result.length()) {
                if (i < result.length() - 1 && result.charAt(i) == result.charAt(i + 1)) {
                    i += 2; // Skip duplicate pair
                    changed = true;
                } else {
                    temp += result.charAt(i);
                    i++;
                }
            }
            result = temp;
        } while (changed);
        
        System.out.println("Result: " + result);
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.Stack;

public class RemoveAdjacentDuplicates {
    public static void main(String[] args) {
        String str = "abbaca";
        Stack<Character> stack = new Stack<>();
        
        for (char ch : str.toCharArray()) {
            if (!stack.isEmpty() && stack.peek() == ch) {
                stack.pop();
            } else {
                stack.push(ch);
            }
        }
        
        StringBuilder result = new StringBuilder();
        for (char ch : stack) result.append(ch);
        System.out.println("Result: " + result);
    }
}
```

### 11. Check if Strings are Isomorphic
**Principle**: One-to-one mapping between characters.
**Question**: Check if two strings are isomorphic.

**Brute Force Approach (O(n²))**:
```java
public class IsomorphicBruteForce {
    public static void main(String[] args) {
        String s1 = "egg", s2 = "add";
        System.out.println("Isomorphic? " + isIsomorphicBruteForce(s1, s2));
    }
    
    static boolean isIsomorphicBruteForce(String s1, String s2) {
        if (s1.length() != s2.length()) return false;
        
        for (int i = 0; i < s1.length(); i++) {
            char ch1 = s1.charAt(i), ch2 = s2.charAt(i);
            
            // Check if ch1 maps to consistent ch2
            for (int j = 0; j < i; j++) {
                if (s1.charAt(j) == ch1 && s2.charAt(j) != ch2) {
                    return false;
                }
                if (s2.charAt(j) == ch2 && s1.charAt(j) != ch1) {
                    return false;
                }
            }
        }
        return true;
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class Isomorphic {
    public static void main(String[] args) {
        String s1 = "egg", s2 = "add";
        if (s1.length() != s2.length()) {
            System.out.println("Not Isomorphic");
            return;
        }
        
        Map<Character, Character> map = new HashMap<>();
        Set<Character> used = new HashSet<>();
        
        for (int i = 0; i < s1.length(); i++) {
            char ch1 = s1.charAt(i), ch2 = s2.charAt(i);
            
            if (map.containsKey(ch1)) {
                if (map.get(ch1) != ch2) {
                    System.out.println("Not Isomorphic");
                    return;
                }
            } else {
                if (used.contains(ch2)) {
                    System.out.println("Not Isomorphic");
                    return;
                }
                map.put(ch1, ch2);
                used.add(ch2);
            }
        }
        System.out.println("Isomorphic");
    }
}
```

### 12. Check if Strings are One Edit Away
**Principle**: Check insert, delete, or replace scenarios.
**Question**: Check if two strings are one edit away.

**Brute Force Approach (O(n³))**:
```java
public class OneEditAwayBruteForce {
    public static void main(String[] args) {
        String s1 = "pale", s2 = "ple";
        System.out.println("One edit away? " + isOneEditAwayBruteForce(s1, s2));
    }
    
    static boolean isOneEditAwayBruteForce(String s1, String s2) {
        // Try all possible single edits on s1
        // 1. Replace each character
        for (int i = 0; i < s1.length(); i++) {
            String modified = s1.substring(0, i) + 'a' + s1.substring(i + 1);
            if (modified.equals(s2)) return true;
        }
        
        // 2. Delete each character
        for (int i = 0; i < s1.length(); i++) {
            String modified = s1.substring(0, i) + s1.substring(i + 1);
            if (modified.equals(s2)) return true;
        }
        
        // 3. Insert each possible character
        for (int i = 0; i <= s1.length(); i++) {
            for (char c = 'a'; c <= 'z'; c++) {
                String modified = s1.substring(0, i) + c + s1.substring(i);
                if (modified.equals(s2)) return true;
            }
        }
        
        return s1.equals(s2);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class OneEditAway {
    public static void main(String[] args) {
        String s1 = "pale", s2 = "ple";
        System.out.println("One edit away? " + isOneEditAway(s1, s2));
    }
    
    static boolean isOneEditAway(String s1, String s2) {
        if (s1.equals(s2)) return true;
        
        int len1 = s1.length(), len2 = s2.length();
        if (Math.abs(len1 - len2) > 1) return false;
        
        String shorter = len1 < len2 ? s1 : s2;
        String longer = len1 < len2 ? s2 : s1;
        
        boolean foundDifference = false;
        int i = 0, j = 0;
        
        while (i < shorter.length() && j < longer.length()) {
            if (shorter.charAt(i) != longer.charAt(j)) {
                if (foundDifference) return false;
                foundDifference = true;
                if (len1 == len2) i++;
            } else {
                i++;
            }
            j++;
        }
        return true;
    }
}
```

## 📚 FAANG/Product-Based Questions (36-40)

### 13. Group Anagrams
**Principle**: Sort each string and use as key in HashMap.
**Question**: Group anagrams together.

**Brute Force Approach (O(n² * m log m))**:
```java
import java.util.*;

public class GroupAnagramsBruteForce {
    public static void main(String[] args) {
        String[] words = {"eat", "tea", "tan", "ate", "nat", "bat"};
        List<List<String>> result = groupAnagramsBruteForce(words);
        
        for (List<String> group : result) {
            System.out.println(group);
        }
    }
    
    static List<List<String>> groupAnagramsBruteForce(String[] words) {
        List<List<String>> result = new ArrayList<>();
        boolean[] used = new boolean[words.length];
        
        for (int i = 0; i < words.length; i++) {
            if (used[i]) continue;
            
            List<String> group = new ArrayList<>();
            group.add(words[i]);
            used[i] = true;
            
            for (int j = i + 1; j < words.length; j++) {
                if (!used[j] && areAnagrams(words[i], words[j])) {
                    group.add(words[j]);
                    used[j] = true;
                }
            }
            result.add(group);
        }
        return result;
    }
    
    static boolean areAnagrams(String s1, String s2) {
        if (s1.length() != s2.length()) return false;
        
        char[] c1 = s1.toCharArray();
        char[] c2 = s2.toCharArray();
        Arrays.sort(c1);
        Arrays.sort(c2);
        
        return Arrays.equals(c1, c2);
    }
}
```

**Optimized Approach (O(n * m log m))**:
```java
import java.util.*;

public class GroupAnagrams {
    public static void main(String[] args) {
        String[] words = {"eat", "tea", "tan", "ate", "nat", "bat"};
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String word : words) {
            char[] chars = word.toCharArray();
            Arrays.sort(chars);
            String key = new String(chars);
            
            groups.computeIfAbsent(key, k -> new ArrayList<>()).add(word);
        }
        
        for (List<String> group : groups.values()) {
            System.out.println(group);
        }
    }
}
```

### 14. Minimum Window Substring
**Principle**: Sliding window with character frequency.
**Question**: Find minimum window containing all characters of another string.

**Brute Force Approach (O(n²))**:
```java
import java.util.*;

public class MinWindowSubstringBruteForce {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        System.out.println("Minimum window: " + minWindowBruteForce(s, t));
    }
    
    static String minWindowBruteForce(String s, String t) {
        String result = "";
        int minLen = Integer.MAX_VALUE;
        
        for (int i = 0; i < s.length(); i++) {
            for (int j = i; j < s.length(); j++) {
                String window = s.substring(i, j + 1);
                if (containsAllChars(window, t) && window.length() < minLen) {
                    result = window;
                    minLen = window.length();
                }
            }
        }
        return result;
    }
    
    static boolean containsAllChars(String window, String t) {
        Map<Character, Integer> windowCount = new HashMap<>();
        for (char ch : window.toCharArray()) {
            windowCount.put(ch, windowCount.getOrDefault(ch, 0) + 1);
        }
        
        Map<Character, Integer> tCount = new HashMap<>();
        for (char ch : t.toCharArray()) {
            tCount.put(ch, tCount.getOrDefault(ch, 0) + 1);
        }
        
        for (char ch : tCount.keySet()) {
            if (windowCount.getOrDefault(ch, 0) < tCount.get(ch)) {
                return false;
            }
        }
        return true;
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class MinWindowSubstring {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        System.out.println("Minimum window: " + minWindow(s, t));
    }
    
    static String minWindow(String s, String t) {
        if (s.length() == 0 || t.length() == 0) return "";
        
        Map<Character, Integer> need = new HashMap<>();
        for (char ch : t.toCharArray()) need.put(ch, need.getOrDefault(ch, 0) + 1);
        
        int left = 0, formed = 0, required = need.size();
        Map<Character, Integer> window = new HashMap<>();
        int[] ans = {-1, 0, 0};
        
        for (int right = 0; right < s.length(); right++) {
            char ch = s.charAt(right);
            window.put(ch, window.getOrDefault(ch, 0) + 1);
            
            if (need.containsKey(ch) && window.get(ch).intValue() == need.get(ch).intValue()) {
                formed++;
            }
            
            while (left <= right && formed == required) {
                char c = s.charAt(left);
                if (ans[0] == -1 || right - left + 1 < ans[0]) {
                    ans[0] = right - left + 1;
                    ans[1] = left;
                    ans[2] = right;
                }
                
                window.put(c, window.get(c) - 1);
                if (need.containsKey(c) && window.get(c).intValue() < need.get(c).intValue()) {
                    formed--;
                }
                left++;
            }
        }
        return ans[0] == -1 ? "" : s.substring(ans[1], ans[2] + 1);
    }
}
```

### 15. Decode String
**Principle**: Use Stack for numbers and strings.
**Question**: Decode encoded string like 3[a2[c]] → accaccacc.

**Brute Force Approach (O(n³))**:
```java
public class DecodeStringBruteForce {
    public static void main(String[] args) {
        String str = "3[a2[c]]";
        System.out.println("Decoded: " + decodeStringBruteForce(str));
    }
    
    static String decodeStringBruteForce(String s) {
        String result = s;
        boolean changed;
        
        do {
            changed = false;
            String temp = "";
            int i = 0;
            
            while (i < result.length()) {
                if (Character.isDigit(result.charAt(i))) {
                    int num = 0;
                    while (i < result.length() && Character.isDigit(result.charAt(i))) {
                        num = num * 10 + (result.charAt(i) - '0');
                        i++;
                    }
                    
                    if (i < result.length() && result.charAt(i) == '[') {
                        int j = i + 1, count = 1;
                        while (j < result.length() && count > 0) {
                            if (result.charAt(j) == '[') count++;
                            if (result.charAt(j) == ']') count--;
                            j++;
                        }
                        
                        String substring = result.substring(i + 1, j - 1);
                        String expanded = "";
                        for (int k = 0; k < num; k++) {
                            expanded += substring;
                        }
                        temp += expanded;
                        i = j;
                        changed = true;
                    }
                } else {
                    temp += result.charAt(i);
                    i++;
                }
            }
            result = temp;
        } while (changed);
        
        return result;
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.Stack;

public class DecodeString {
    public static void main(String[] args) {
        String str = "3[a2[c]]";
        System.out.println("Decoded: " + decodeString(str));
    }
    
    static String decodeString(String s) {
        Stack<Integer> countStack = new Stack<>();
        Stack<StringBuilder> stringStack = new Stack<>();
        StringBuilder current = new StringBuilder();
        int k = 0;
        
        for (char ch : s.toCharArray()) {
            if (Character.isDigit(ch)) {
                k = k * 10 + (ch - '0');
            } else if (ch == '[') {
                countStack.push(k);
                stringStack.push(current);
                current = new StringBuilder();
                k = 0;
            } else if (ch == ']') {
                StringBuilder decoded = new StringBuilder(stringStack.pop());
                int count = countStack.pop();
                for (int i = 0; i < count; i++) {
                    decoded.append(current);
                }
                current = decoded;
            } else {
                current.append(ch);
            }
        }
        return current.toString();
    }
}
```

### 16. Multiply Strings
**Principle**: Manual multiplication like grade school.
**Question**: Multiply two string numbers without converting to integer.

**Brute Force Approach (O(n³))**:
```java
public class MultiplyStringsBruteForce {
    public static void main(String[] args) {
        String num1 = "123", num2 = "456";
        System.out.println("Product: " + multiplyBruteForce(num1, num2));
    }
    
    static String multiplyBruteForce(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0")) return "0";
        
        String result = "0";
        
        for (int i = num1.length() - 1; i >= 0; i--) {
            int digit1 = num1.charAt(i) - '0';
            String temp = "";
            int carry = 0;
            
            // Multiply digit1 with entire num2
            for (int j = num2.length() - 1; j >= 0; j--) {
                int digit2 = num2.charAt(j) - '0';
                int product = digit1 * digit2 + carry;
                carry = product / 10;
                temp = (product % 10) + temp;
            }
            
            if (carry > 0) {
                temp = carry + temp;
            }
            
            // Add zeros based on position
            for (int k = 0; k < num1.length() - 1 - i; k++) {
                temp += "0";
            }
            
            result = addStrings(result, temp);
        }
        
        return result;
    }
    
    static String addStrings(String a, String b) {
        String result = "";
        int i = a.length() - 1, j = b.length() - 1, carry = 0;
        
        while (i >= 0 || j >= 0 || carry > 0) {
            int digit1 = i >= 0 ? a.charAt(i--) - '0' : 0;
            int digit2 = j >= 0 ? b.charAt(j--) - '0' : 0;
            int sum = digit1 + digit2 + carry;
            carry = sum / 10;
            result = (sum % 10) + result;
        }
        
        return result;
    }
}
```

**Optimized Approach (O(n²))**:
```java
public class MultiplyStrings {
    public static void main(String[] args) {
        String num1 = "123", num2 = "456";
        System.out.println("Product: " + multiply(num1, num2));
    }
    
    static String multiply(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0")) return "0";
        
        int m = num1.length(), n = num2.length();
        int[] result = new int[m + n];
        
        for (int i = m - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                int mul = (num1.charAt(i) - '0') * (num2.charAt(j) - '0');
                int sum = mul + result[i + j + 1];
                result[i + j + 1] = sum % 10;
                result[i + j] += sum / 10;
            }
        }
        
        StringBuilder sb = new StringBuilder();
        for (int digit : result) {
            if (!(sb.length() == 0 && digit == 0)) {
                sb.append(digit);
            }
        }
        return sb.toString();
    }
}
```

## 📚 Java-Specific Theory Questions (31-35)

### 17. String Pool Demo
**Principle**: String literal pool vs heap objects.
**Question**: Demonstrate String pool concept.
**Code**:
```java
public class StringPool {
    public static void main(String[] args) {
        String s1 = "Java";
        String s2 = "Java";
        String s3 = new String("Java");
        
        System.out.println("s1 == s2: " + (s1 == s2));
        System.out.println("s1 == s3: " + (s1 == s3));
        System.out.println("s1.equals(s3): " + s1.equals(s3));
        
        String s4 = s3.intern();
        System.out.println("s1 == s4: " + (s1 == s4));
    }
}
```

### 18. equals() vs == Demo
**Principle**: Reference comparison vs content comparison.
**Question**: Show difference between equals() and ==.
**Code**:
```java
public class EqualsVsEqualsEquals {
    public static void main(String[] args) {
        String a = new String("Java");
        String b = new String("Java");
        String c = "Java";
        String d = "Java";
        
        System.out.println("a == b: " + (a == b));
        System.out.println("a.equals(b): " + a.equals(b));
        System.out.println("c == d: " + (c == d));
        System.out.println("a == c: " + (a == c));
        System.out.println("a.equals(c): " + a.equals(c));
    }
}
```

### 19. String Immutability Demo
**Principle**: String objects cannot be modified.
**Question**: Demonstrate String immutability.
**Code**:
```java
public class StringImmutability {
    public static void main(String[] args) {
        String s1 = "Hello";
        String s2 = s1.concat(" World");
        
        System.out.println("s1: " + s1);
        System.out.println("s2: " + s2);
        System.out.println("s1 == s2: " + (s1 == s2));
        
        s1 = s1 + " World";
        System.out.println("After reassignment s1: " + s1);
        System.out.println("s2: " + s2);
    }
}
```

### 20. StringBuilder vs StringBuffer Demo
**Principle**: Thread safety vs performance.
**Question**: Compare StringBuilder and StringBuffer.
**Code**:
```java
public class BuilderVsBuffer {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("Hello");
        StringBuffer sbuf = new StringBuffer("Hello");
        
        sb.append(" World");
        sbuf.append(" World");
        
        System.out.println("StringBuilder: " + sb);
        System.out.println("StringBuffer: " + sbuf);
        
        System.out.println("StringBuilder is thread-safe: " + 
            (StringBuilder.class.isSynchronized() ? "Yes" : "No"));
        System.out.println("StringBuffer is thread-safe: " + 
            (StringBuffer.class.isSynchronized() ? "Yes" : "No"));
    }
}
```

---

## 📚 Original Programs (36-55) - Already Covered

### 21. Reverse a String
**Principle**: Use StringBuilder or iterate backward.
**Question**: Reverse a string without using built-in reverse function.

**Brute Force Approach (O(n²))**:
```java
public class ReverseStringBruteForce {
    public static void main(String[] args) {
        String str = "Hello";
        String reversed = "";
        
        for (int i = 0; i < str.length(); i++) {
            reversed = str.charAt(i) + reversed;
        }
        
        System.out.println("Reversed: " + reversed);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class ReverseString {
    public static void main(String[] args) {
        String str = "Hello";
        char[] chars = str.toCharArray();
        int left = 0, right = chars.length - 1;
        
        while (left < right) {
            char temp = chars[left];
            chars[left] = chars[right];
            chars[right] = temp;
            left++;
            right--;
        }
        
        System.out.println("Reversed: " + new String(chars));
    }
}
```

## 37. Check Palindrome String
**Principle**: String equals its reverse.
**Question**: Check if a string is a palindrome.

**Brute Force Approach (O(n²))**:
```java
public class PalindromeStringBruteForce {
    public static void main(String[] args) {
        String str = "madam";
        String reversed = "";
        
        for (int i = 0; i < str.length(); i++) {
            reversed = str.charAt(i) + reversed;
        }
        
        System.out.println(str.equals(reversed) ? "Palindrome" : "Not Palindrome");
    }
}
```

**Optimized Approach (O(n))**:
```java
public class PalindromeString {
    public static void main(String[] args) {
        String str = "madam";
        int left = 0, right = str.length() - 1;
        boolean isPalindrome = true;
        
        while (left < right) {
            if (str.charAt(left) != str.charAt(right)) {
                isPalindrome = false;
                break;
            }
            left++;
            right--;
        }
        
        System.out.println(isPalindrome ? "Palindrome" : "Not Palindrome");
    }
}
```

## 38. Count Vowels and Consonants
**Principle**: Check if char is in "aeiou".
**Question**: Count vowels and consonants in a string.

**Brute Force Approach (O(n²))**:
```java
public class CountVowelsBruteForce {
    public static void main(String[] args) {
        String str = "Hello World";
        int v = 0, c = 0;
        
        for (char ch : str.toLowerCase().toCharArray()) {
            if (ch >= 'a' && ch <= 'z') {
                boolean isVowel = false;
                char[] vowels = {'a', 'e', 'i', 'o', 'u'};
                
                for (char vowel : vowels) {
                    if (ch == vowel) {
                        isVowel = true;
                        break;
                    }
                }
                
                if (isVowel) v++;
                else c++;
            }
        }
        
        System.out.println("Vowels: " + v + ", Consonants: " + c);
    }
}
```

**Optimized Approach (O(n))**:
```java
public class CountVowels {
    public static void main(String[] args) {
        String str = "Hello World";
        int v = 0, c = 0;
        for (char ch : str.toLowerCase().toCharArray()) {
            if (ch >= 'a' && ch <= 'z') {
                if ("aeiou".indexOf(ch) != -1) v++;
                else c++;
            }
        }
        System.out.println("Vowels: " + v + ", Consonants: " + c);
    }
}
```

## 39. Anagram Check
**Principle**: Sort both strings and compare.
**Question**: Check if two strings are anagrams.

**Brute Force Approach (O(n³))**:
```java
public class AnagramBruteForce {
    public static void main(String[] args) {
        String s1 = "Listen", s2 = "Silent";
        System.out.println("Anagram? " + areAnagramsBruteForce(s1, s2));
    }
    
    static boolean areAnagramsBruteForce(String s1, String s2) {
        if (s1.length() != s2.length()) return false;
        
        s1 = s1.toLowerCase();
        s2 = s2.toLowerCase();
        boolean[] used = new boolean[s2.length()];
        
        for (int i = 0; i < s1.length(); i++) {
            boolean found = false;
            for (int j = 0; j < s2.length(); j++) {
                if (!used[j] && s1.charAt(i) == s2.charAt(j)) {
                    used[j] = true;
                    found = true;
                    break;
                }
            }
            if (!found) return false;
        }
        
        return true;
    }
}
```

**Optimized Approach (O(n log n))**:
```java
import java.util.Arrays;

public class Anagram {
    public static void main(String[] args) {
        String s1 = "Listen", s2 = "Silent";
        char[] c1 = s1.toLowerCase().toCharArray();
        char[] c2 = s2.toLowerCase().toCharArray();
        Arrays.sort(c1);
        Arrays.sort(c2);
        System.out.println("Anagram? " + Arrays.equals(c1, c2));
    }
}
```

## 40. Duplicate Characters in String
**Principle**: Use HashMap or int[256].
**Question**: Find duplicate characters in a string.
**Code**:
```java
import java.util.*;

public class DuplicateChars {
    public static void main(String[] args) {
        String str = "programming";
        Map<Character, Integer> map = new HashMap<>();
        for (char ch : str.toCharArray()) {
            map.put(ch, map.getOrDefault(ch, 0) + 1);
        }
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue() > 1) {
                System.out.println(entry.getKey() + ": " + entry.getValue());
            }
        }
    }
}
```

## 41. Count Words in String
**Principle**: Split by space or count spaces + 1.
**Question**: Count number of words in a sentence.
**Code**:
```java
public class CountWords {
    public static void main(String[] args) {
        String str = "Java is fun";
        String[] words = str.trim().split("\\s+");
        System.out.println("Words: " + words.length);
    }
}
```

## 42. Remove White Spaces
**Principle**: Use `replaceAll`.
**Question**: Remove all white spaces from a string.
**Code**:
```java
public class RemoveSpaces {
    public static void main(String[] args) {
        String str = "J a v a";
        String noSpace = str.replaceAll("\\s", "");
        System.out.println(noSpace);
    }
}
```

## 43. Swapping Pair of Characters
**Principle**: Swap str[i] and str[i+1].
**Question**: Swap pairs of characters in a string.
**Code**:
```java
public class SwapPairs {
    public static void main(String[] args) {
        String str = "123456";
        char[] ch = str.toCharArray();
        for(int i=0; i<ch.length-1; i+=2) {
            char temp = ch[i];
            ch[i] = ch[i+1];
            ch[i+1] = temp;
        }
        System.out.println(new String(ch));
    }
}
```

## 44. String contains only Digits
**Principle**: Check each char or use regex.
**Question**: Check if a string contains only digits.
**Code**:
```java
public class OnlyDigits {
    public static void main(String[] args) {
        String str = "12345";
        System.out.println(str.matches("\\d+") ? "Only Digits" : "Mixed");
    }
}
```

## 45. Find First Non-Repeated Character
**Principle**: Use LinkedHashMap to maintain order + Count.
**Question**: Find the first non-repeated character in a string.

**Brute Force Approach (O(n²))**:
```java
public class FirstNonRepeatedBruteForce {
    public static void main(String[] args) {
        String str = "swiss";
        
        for (int i = 0; i < str.length(); i++) {
            char ch = str.charAt(i);
            boolean isRepeated = false;
            
            for (int j = 0; j < str.length(); j++) {
                if (i != j && ch == str.charAt(j)) {
                    isRepeated = true;
                    break;
                }
            }
            
            if (!isRepeated) {
                System.out.println("First non-repeated: " + ch);
                return;
            }
        }
        System.out.println("No non-repeated character");
    }
}
```

**Optimized Approach (O(n))**:
```java
import java.util.*;

public class FirstNonRepeated {
    public static void main(String[] args) {
        String str = "swiss";
        Map<Character, Integer> map = new LinkedHashMap<>();
        for (char ch : str.toCharArray()) map.put(ch, map.getOrDefault(ch, 0) + 1);
        
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue() == 1) {
                System.out.println("First non-repeated: " + entry.getKey());
                break;
            }
        }
    }
}
```

## 46. Remove Special Characters
**Principle**: Regex `[^a-zA-Z0-9]`.
**Question**: Remove all special characters from a string.
**Code**:
```java
public class RemoveSpecial {
    public static void main(String[] args) {
        String str = "Ja@va#$";
        System.out.println(str.replaceAll("[^a-zA-Z0-9]", ""));
    }
}
```

## 47. String to Integer (atoi)
**Principle**: `Integer.parseInt` or manual parsing.
**Question**: Convert String to Integer.
**Code**:
```java
public class StringToInt {
    public static void main(String[] args) {
        String str = "123";
        int num = Integer.parseInt(str);
        System.out.println(num);
    }
}
```

## 48. Integer to String
**Principle**: `String.valueOf` or `Integer.toString`.
**Question**: Convert Integer to String.
**Code**:
```java
public class IntToString {
    public static void main(String[] args) {
        int i = 123;
        String s = String.valueOf(i);
        System.out.println(s);
    }
}
```

## 49. Count Occurrences of Character
**Principle**: `len - len(replace(char, ""))`.
**Question**: Count occurrences of a character in string without loop using replace.
**Code**:
```java
public class CharCount {
    public static void main(String[] args) {
        String s = "Java Programming";
        char c = 'a';
        long count = s.length() - s.replaceAll(String.valueOf(c), "").length();
        System.out.println("Count of " + c + ": " + count);
    }
}
```

## 50. Permutations of a String
**Principle**: Recursive backtracking.
**Question**: Print all permutations of a string.
**Code**:
```java
public class Permutations {
    public static void main(String[] args) {
        permute("ABC", "");
    }
    static void permute(String str, String ans) {
        if (str.length() == 0) {
            System.out.print(ans + " ");
            return;
        }
        for (int i = 0; i < str.length(); i++) {
            char ch = str.charAt(i);
            String rest = str.substring(0, i) + str.substring(i + 1);
            permute(rest, ans + ch);
        }
    }
}
```

## 51. Longest Common Prefix
**Principle**: Sort array, compare first and last strings.
**Question**: Find longest common prefix in an array of strings.
**Code**:
```java
import java.util.Arrays;

public class LongestPrefix {
    public static void main(String[] args) {
        String[] strs = {"flower", "flow", "flight"};
        Arrays.sort(strs);
        String s1 = strs[0], s2 = strs[strs.length-1];
        int i = 0;
        while(i < s1.length() && i < s2.length() && s1.charAt(i) == s2.charAt(i)) {
            i++;
        }
        System.out.println("Prefix: " + s1.substring(0, i));
    }
}
```

## 52. Reverse Words in a String
**Principle**: Split by space, reverse array, join.
**Question**: Reverse the order of words in a sentence.
**Code**:
```java
public class ReverseWords {
    public static void main(String[] args) {
        String str = "Hello World Java";
        String[] words = str.split(" ");
        String res = "";
        for(int i = words.length-1; i >= 0; i--) {
            res += words[i] + " ";
        }
        System.out.println(res.trim());
    }
}
```

## 53. Check for Subsequence
**Principle**: Two pointers.
**Question**: Check if string A is a subsequence of string B.
**Code**:
```java
public class Subsequence {
    public static void main(String[] args) {
        String s1 = "abc", s2 = "ahbgdc";
        int i=0, j=0;
        while(i < s1.length() && j < s2.length()){
            if(s1.charAt(i) == s2.charAt(j)) i++;
            j++;
        }
        System.out.println("Is Subsequence? " + (i == s1.length()));
    }
}
```

## 54. Rotation Check
**Principle**: A rotation of A is found in A+A.
**Question**: Check if one string is a rotation of another (e.g., "ABCD" and "CDAB").
**Code**:
```java
public class RotationCheck {
    public static void main(String[] args) {
        String s1 = "ABCD", s2 = "CDAB";
        if(s1.length() == s2.length() && (s1+s1).contains(s2)) {
             System.out.println("Rotation");
        } else {
             System.out.println("Not Rotation");
        }
    }
}
```

## 55. Length of String without length()
**Principle**: Convert to array and loop.
**Question**: Find length of string without using `length()` method.
**Code**:
```java
public class StringLength {
    public static void main(String[] args) {
        String str = "Hello";
        int count = 0;
        for(char c : str.toCharArray()) count++;
        System.out.println("Length: " + count);
    }
}
```
