import java.util.*;

public class StringAlgorithms {
    
    // 28. Find the Index of the First Occurrence in a String
    // Time: O(N + M), Space: O(M) for KMP prefix table
    public static int strStr(String haystack, String needle) {
        if (haystack == null || needle == null) {
            return -1;
        }
        if (needle.length() == 0) {
            return 0;
        }
        if (haystack.length() < needle.length()) {
            return -1;
        }
        
        // KMP algorithm
        int[] lps = buildLPS(needle);
        int i = 0; // index for haystack
        int j = 0; // index for needle
        
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                
                if (j == needle.length()) {
                    return i - j; // Found match
                }
            } else {
                if (j != 0) {
                    j = lps[j - 1];
                } else {
                    i++;
                }
            }
        }
        
        return -1;
    }
    
    private static int[] buildLPS(String pattern) {
        int[] lps = new int[pattern.length()];
        int len = 0; // length of the previous longest prefix suffix
        int i = 1;
        
        while (i < pattern.length()) {
            if (pattern.charAt(i) == pattern.charAt(len)) {
                len++;
                lps[i] = len;
                i++;
            } else {
                if (len != 0) {
                    len = lps[len - 1];
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        
        return lps;
    }

    // 214. Shortest Palindrome
    // Time: O(N), Space: O(N)
    public static String shortestPalindrome(String s) {
        if (s == null || s.length() <= 1) {
            return s;
        }
        
        String rev = new StringBuilder(s).reverse().toString();
        String combined = s + "#" + rev;
        int[] lps = buildLPS(combined);
        
        int longestPalindromePrefix = lps[lps.length - 1];
        String suffix = s.substring(longestPalindromePrefix);
        
        return new StringBuilder(suffix).reverse().toString() + s;
    }

    // 459. Repeated Substring Pattern
    // Time: O(N), Space: O(N)
    public static boolean repeatedSubstringPattern(String s) {
        if (s == null || s.length() <= 1) {
            return false;
        }
        
        String doubled = s + s;
        String sub = doubled.substring(1, doubled.length() - 1);
        
        return sub.contains(s);
    }

    public static void main(String[] args) {
        // Test cases for strStr
        String[][] testCases1 = {
            {"hello", "ll"},
            {"aaaaa", "bba"},
            {"", ""},
            {"", "a"},
            {"a", ""},
            {"a", "a"},
            {"abc", "c"},
            {"abc", "d"},
            {"mississippi", "issi"},
            {"mississippi", "issip"},
            {"abababab", "abab"},
            {"aaaaaaaa", "aaa"},
            {"abcabcabc", "abc"},
            {"abcdef", "def"},
            {"abcdef", "abc"},
            {"abcdef", "xyz"}
        };
        
        // Test cases for shortestPalindrome
        String[] testCases2 = {
            "aacecaaa",
            "abcd",
            "a",
            "",
            "aa",
            "aba",
            "abac",
            "racecar",
            "abacdfgdcaba",
            "abcba",
            "ab",
            "a",
            "abc",
            "abccba",
            "abcbabc"
        };
        
        // Test cases for repeatedSubstringPattern
        String[] testCases3 = {
            "abab",
            "aba",
            "abcabcabc",
            "bbbb",
            "abac",
            "a",
            "",
            "aa",
            "aaa",
            "abcabc",
            "abcdabcd",
            "xyzxyz",
            "abc",
            "ab",
            "abcabcabcabc"
        };
        
        System.out.println("Find Index of First Occurrence:");
        for (int i = 0; i < testCases1.length; i++) {
            String haystack = testCases1[i][0];
            String needle = testCases1[i][1];
            int result = strStr(haystack, needle);
            System.out.printf("Test Case %d: \"%s\" in \"%s\" -> %d\n", 
                i + 1, needle, haystack, result);
        }
        
        System.out.println("\nShortest Palindrome:");
        for (int i = 0; i < testCases2.length; i++) {
            String s = testCases2[i];
            String result = shortestPalindrome(s);
            System.out.printf("Test Case %d: \"%s\" -> \"%s\"\n", 
                i + 1, s, result);
        }
        
        System.out.println("\nRepeated Substring Pattern:");
        for (int i = 0; i < testCases3.length; i++) {
            String s = testCases3[i];
            boolean result = repeatedSubstringPattern(s);
            System.out.printf("Test Case %d: \"%s\" -> %b\n", 
                i + 1, s, result);
        }
    }
}
