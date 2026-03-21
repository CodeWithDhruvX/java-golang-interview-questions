```java
import java.util.*;

public class PatternMatching {

    // 1. Implement strStr() (Find Substring)
    // Brute Force: O(N*M), Space: O(1)
    public static int strStrBruteForce(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        if (needle.length() > haystack.length()) {
            return -1;
        }
        
        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            int j = 0;
            while (j < needle.length() && haystack.charAt(i + j) == needle.charAt(j)) {
                j++;
            }
            if (j == needle.length()) {
                return i;
            }
        }
        return -1;
    }

    // Optimized: O(N*M) - Using built-in which is optimized, potentially O(N+M)
    public static int strStr(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        return haystack.indexOf(needle);
    }

    // Naive implementation for educational purposes
    // Time: O(N*M), Space: O(1)
    public static int strStrNaive(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        if (needle.length() > haystack.length()) {
            return -1;
        }
        
        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            int j = 0;
            while (j < needle.length() && haystack.charAt(i + j) == needle.charAt(j)) {
                j++;
            }
            if (j == needle.length()) {
                return i;
            }
        }
        return -1;
    }

    // 2. Longest Common Prefix
    // Brute Force: O(N^2 * L), Space: O(1)
    public static String longestCommonPrefixBruteForce(String[] strs) {
        if (strs == null || strs.length == 0) {
            return "";
        }
        
        String prefix = strs[0];
        for (int i = 1; i < strs.length; i++) {
            while (strs[i].indexOf(prefix) != 0) {
                prefix = prefix.substring(0, prefix.length() - 1);
                if (prefix.isEmpty()) {
                    return "";
                }
            }
        }
        return prefix;
    }

    // Optimized: O(N*L*logN), Space: O(L)
    public static String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length == 0) {
            return "";
        }
        
        // Sort the strings
        Arrays.sort(strs);
        
        String s1 = strs[0];
        String s2 = strs[strs.length - 1]; // Most different from s1
        
        int idx = 0;
        while (idx < s1.length() && idx < s2.length()) {
            if (s1.charAt(idx) == s2.charAt(idx)) {
                idx++;
            } else {
                break;
            }
        }
        return s1.substring(0, idx);
    }

    // Test methods
    public static void main(String[] args) {
        // Test strStr
        String test1h = "hello";
        String test1n = "ll";
        System.out.println("Index of '" + test1n + "' in '" + test1h + "':");
        System.out.println("  Brute Force: " + strStrBruteForce(test1h, test1n));
        System.out.println("  Optimized: " + strStr(test1h, test1n));
        
        // Test longestCommonPrefix
        String[] test3 = {"flower", "flow", "flight"};
        System.out.println("\nLongest common prefix of " + Arrays.toString(test3) + ":");
        System.out.println("  Brute Force: '" + longestCommonPrefixBruteForce(test3) + "'");
        System.out.println("  Optimized: '" + longestCommonPrefix(test3) + "'");
        
        String[] test4 = {"dog", "racecar", "car"};
        System.out.println("\nLongest common prefix of " + Arrays.toString(test4) + ":");
        System.out.println("  Brute Force: '" + longestCommonPrefixBruteForce(test4) + "'");
        System.out.println("  Optimized: '" + longestCommonPrefix(test4) + "'");
    }
}
```
