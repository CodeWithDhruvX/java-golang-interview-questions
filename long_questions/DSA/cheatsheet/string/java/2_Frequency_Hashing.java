```java
import java.util.*;

public class FrequencyHashing {

    // 1. First Non-Repeating Character
    // Brute Force: O(N^2), Space: O(1)
    public static char firstNonRepeatingCharBruteForce(String s) {
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            boolean isRepeating = false;
            for (int j = 0; j < s.length(); j++) {
                if (i != j && s.charAt(j) == c) {
                    isRepeating = true;
                    break;
                }
            }
            if (!isRepeating) {
                return c;
            }
        }
        return '\0';
    }

    // Optimized: O(N), Space: O(1) (fixed alphabet)
    public static char firstNonRepeatingChar(String s) {
        Map<Character, Integer> freq = new HashMap<>();
        
        // Pass 1: Count frequencies
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            freq.put(c, freq.getOrDefault(c, 0) + 1);
        }
        
        // Pass 2: Find first with count 1
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (freq.get(c) == 1) {
                return c;
            }
        }
        return '\0'; // or throw exception
    }

    // 2. Check if Two Strings are Anagrams
    // Brute Force: O(N^2), Space: O(1)
    public static boolean areAnagramsBruteForce(String s1, String s2) {
        if (s1.length() != s2.length()) {
            return false;
        }
        
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
            if (!found) {
                return false;
            }
        }
        return true;
    }

    // Optimized: O(N), Space: O(1)
    public static boolean areAnagrams(String s1, String s2) {
        if (s1.length() != s2.length()) {
            return false;
        }
        
        int[] freq = new int[26];
        for (int i = 0; i < s1.length(); i++) {
            freq[s1.charAt(i) - 'a']++;
        }
        
        for (int i = 0; i < s2.length(); i++) {
            freq[s2.charAt(i) - 'a']--;
            if (freq[s2.charAt(i) - 'a'] < 0) {
                return false;
            }
        }
        return true;
    }

    // 3. Find All Duplicate Characters
    // Brute Force: O(N^2), Space: O(1)
    public static List<Character> findDuplicateCharsBruteForce(String s) {
        List<Character> duplicates = new ArrayList<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            boolean isDuplicate = false;
            boolean alreadyAdded = false;
            
            // Check if already added to avoid duplicates in result
            for (char added : duplicates) {
                if (added == c) {
                    alreadyAdded = true;
                    break;
                }
            }
            if (alreadyAdded) continue;
            
            // Check if current character appears more than once
            for (int j = i + 1; j < s.length(); j++) {
                if (s.charAt(j) == c) {
                    isDuplicate = true;
                    break;
                }
            }
            if (isDuplicate) {
                duplicates.add(c);
            }
        }
        return duplicates;
    }

    // Optimized: O(N), Space: O(N)
    public static List<Character> findDuplicateChars(String s) {
        Map<Character, Integer> freq = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            freq.put(c, freq.getOrDefault(c, 0) + 1);
        }
        
        List<Character> duplicates = new ArrayList<>();
        for (Map.Entry<Character, Integer> entry : freq.entrySet()) {
            if (entry.getValue() > 1) {
                duplicates.add(entry.getKey());
            }
        }
        return duplicates;
    }

    // 4. Most Frequent Character
    // Brute Force: O(N^2), Space: O(1)
    public static char mostFrequentCharBruteForce(String s) {
        if (s.isEmpty()) return '\0';
        
        char maxChar = s.charAt(0);
        int maxCount = 0;
        
        for (int i = 0; i < s.length(); i++) {
            char current = s.charAt(i);
            int count = 0;
            for (int j = 0; j < s.length(); j++) {
                if (s.charAt(j) == current) {
                    count++;
                }
            }
            if (count > maxCount) {
                maxCount = count;
                maxChar = current;
            }
        }
        return maxChar;
    }

    // Optimized: O(N), Space: O(1)
    public static char mostFrequentChar(String s) {
        if (s.isEmpty()) return '\0';
        
        Map<Character, Integer> freq = new HashMap<>();
        char maxChar = s.charAt(0);
        int maxCount = 0;
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            int count = freq.getOrDefault(c, 0) + 1;
            freq.put(c, count);
            
            if (count > maxCount) {
                maxCount = count;
                maxChar = c;
            }
        }
        return maxChar;
    }

    // Test methods
    public static void main(String[] args) {
        // Test firstNonRepeatingChar
        String test1 = "leetcode";
        System.out.println("First non-repeating char in '" + test1 + "':");
        System.out.println("  Brute Force: " + firstNonRepeatingCharBruteForce(test1));
        System.out.println("  Optimized: " + firstNonRepeatingChar(test1));
        
        // Test areAnagrams
        String test2a = "anagram";
        String test2b = "nagaram";
        System.out.println("\nAre '" + test2a + "' and '" + test2b + "' anagrams?");
        System.out.println("  Brute Force: " + areAnagramsBruteForce(test2a, test2b));
        System.out.println("  Optimized: " + areAnagrams(test2a, test2b));
        
        // Test findDuplicateChars
        String test3 = "abca";
        System.out.println("\nDuplicate chars in '" + test3 + "':");
        System.out.println("  Brute Force: " + findDuplicateCharsBruteForce(test3));
        System.out.println("  Optimized: " + findDuplicateChars(test3));
        
        // Test mostFrequentChar
        String test4 = "hello";
        System.out.println("\nMost frequent char in '" + test4 + "':");
        System.out.println("  Brute Force: " + mostFrequentCharBruteForce(test4));
        System.out.println("  Optimized: " + mostFrequentChar(test4));
    }
}
```
