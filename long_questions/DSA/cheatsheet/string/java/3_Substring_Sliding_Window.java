```java
import java.util.*;

public class SubstringSlidingWindow {

    // 1. Longest Substring Without Repeating Characters
    // Brute Force: O(N^3), Space: O(min(N, M))
    public static int lengthOfLongestSubstringBruteForce(String s) {
        int maxLength = 0;
        for (int i = 0; i < s.length(); i++) {
            for (int j = i; j < s.length(); j++) {
                if (hasAllUniqueChars(s, i, j)) {
                    maxLength = Math.max(maxLength, j - i + 1);
                }
            }
        }
        return maxLength;
    }
    
    private static boolean hasAllUniqueChars(String s, int start, int end) {
        Set<Character> chars = new HashSet<>();
        for (int i = start; i <= end; i++) {
            if (chars.contains(s.charAt(i))) {
                return false;
            }
            chars.add(s.charAt(i));
        }
        return true;
    }

    // Optimized: O(N), Space: O(U)
    public static int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> lastSeen = new HashMap<>();
        int start = 0;
        int maxLength = 0;
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (lastSeen.containsKey(c) && lastSeen.get(c) >= start) {
                start = lastSeen.get(c) + 1;
            }
            lastSeen.put(c, i);
            int currentLen = i - start + 1;
            if (currentLen > maxLength) {
                maxLength = currentLen;
            }
        }
        return maxLength;
    }

    // 2. Longest Palindromic Substring
    // Brute Force: O(N^3), Space: O(1)
    public static String longestPalindromeBruteForce(String s) {
        if (s == null || s.length() == 0) {
            return "";
        }
        
        String result = "";
        for (int i = 0; i < s.length(); i++) {
            for (int j = i; j < s.length(); j++) {
                String substring = s.substring(i, j + 1);
                if (isPalindrome(substring) && substring.length() > result.length()) {
                    result = substring;
                }
            }
        }
        return result;
    }
    
    private static boolean isPalindrome(String s) {
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

    // Optimized: O(N^2), Space: O(1)
    public static String longestPalindrome(String s) {
        if (s == null || s.length() == 0) {
            return "";
        }
        
        int start = 0, end = 0;
        
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int maxLen = Math.max(len1, len2);
            
            if (maxLen > end - start) {
                start = i - (maxLen - 1) / 2;
                end = i + maxLen / 2;
            }
        }
        return s.substring(start, end + 1);
    }
    
    private static int expandAroundCenter(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1;
    }

    // 3. Find All Anagrams in a String
    // Brute Force: O(N^2 * M), Space: O(M)
    public static List<Integer> findAnagramsBruteForce(String s, String p) {
        List<Integer> result = new ArrayList<>();
        if (s.length() < p.length()) {
            return result;
        }
        
        for (int i = 0; i <= s.length() - p.length(); i++) {
            String substring = s.substring(i, i + p.length());
            if (areAnagrams(substring, p)) {
                result.add(i);
            }
        }
        return result;
    }
    
    private static boolean areAnagrams(String s1, String s2) {
        if (s1.length() != s2.length()) return false;
        int[] count = new int[26];
        for (int i = 0; i < s1.length(); i++) {
            count[s1.charAt(i) - 'a']++;
            count[s2.charAt(i) - 'a']--;
        }
        for (int c : count) {
            if (c != 0) return false;
        }
        return true;
    }

    // Optimized: O(N), Space: O(1)
    public static List<Integer> findAnagrams(String s, String p) {
        List<Integer> result = new ArrayList<>();
        if (s.length() < p.length()) {
            return result;
        }
        
        int[] pFreq = new int[26];
        int[] windowFreq = new int[26];
        
        for (int i = 0; i < p.length(); i++) {
            pFreq[p.charAt(i) - 'a']++;
            windowFreq[s.charAt(i) - 'a']++;
        }
        
        if (Arrays.equals(pFreq, windowFreq)) {
            result.add(0);
        }
        
        for (int i = p.length(); i < s.length(); i++) {
            windowFreq[s.charAt(i) - 'a']++;
            windowFreq[s.charAt(i - p.length()) - 'a']--;
            
            if (Arrays.equals(pFreq, windowFreq)) {
                result.add(i - p.length() + 1);
            }
        }
        return result;
    }

    // 4. Smallest Window Containing All Characters
    // Brute Force: O(N^2 * M), Space: O(M)
    public static String minWindowBruteForce(String s, String t) {
        if (s == null || t == null || s.length() == 0 || t.length() == 0) {
            return "";
        }
        
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
    
    private static boolean containsAllChars(String window, String t) {
        int[] count = new int[256];
        for (int i = 0; i < t.length(); i++) {
            count[t.charAt(i)]++;
        }
        for (int i = 0; i < window.length(); i++) {
            count[window.charAt(i)]--;
        }
        for (int c : count) {
            if (c > 0) return false;
        }
        return true;
    }

    // Optimized: O(N+M), Space: O(1)
    public static String minWindow(String s, String t) {
        if (s == null || t == null || s.length() == 0 || t.length() == 0) {
            return "";
        }
        
        Map<Character, Integer> tFreq = new HashMap<>();
        for (int i = 0; i < t.length(); i++) {
            char c = t.charAt(i);
            tFreq.put(c, tFreq.getOrDefault(c, 0) + 1);
        }
        
        Map<Character, Integer> windowFreq = new HashMap<>();
        int left = 0, right = 0;
        int formed = 0;
        int required = tFreq.size();
        
        int[] ans = {-1, 0, 0}; // length, left, right
        
        while (right < s.length()) {
            char c = s.charAt(right);
            windowFreq.put(c, windowFreq.getOrDefault(c, 0) + 1);
            
            if (tFreq.containsKey(c) && windowFreq.get(c).intValue() == tFreq.get(c).intValue()) {
                formed++;
            }
            
            while (left <= right && formed == required) {
                c = s.charAt(left);
                
                if (ans[0] == -1 || right - left + 1 < ans[0]) {
                    ans[0] = right - left + 1;
                    ans[1] = left;
                    ans[2] = right;
                }
                
                windowFreq.put(c, windowFreq.get(c) - 1);
                if (tFreq.containsKey(c) && windowFreq.get(c) < tFreq.get(c)) {
                    formed--;
                }
                left++;
            }
            right++;
        }
        
        if (ans[0] == -1) {
            return "";
        }
        return s.substring(ans[1], ans[2] + 1);
    }

    // Test methods
    public static void main(String[] args) {
        // Test lengthOfLongestSubstring
        String test1 = "abcabcbb";
        System.out.println("Length of longest substring without repeating chars in '" + test1 + "':");
        System.out.println("  Brute Force: " + lengthOfLongestSubstringBruteForce(test1));
        System.out.println("  Optimized: " + lengthOfLongestSubstring(test1));
        
        // Test longestPalindrome
        String test2 = "babad";
        System.out.println("\nLongest palindrome in '" + test2 + "':");
        System.out.println("  Brute Force: " + longestPalindromeBruteForce(test2));
        System.out.println("  Optimized: " + longestPalindrome(test2));
        
        // Test findAnagrams
        String test3s = "cbaebabacd";
        String test3p = "abc";
        System.out.println("\nAnagrams of '" + test3p + "' in '" + test3s + "':");
        System.out.println("  Brute Force: " + findAnagramsBruteForce(test3s, test3p));
        System.out.println("  Optimized: " + findAnagrams(test3s, test3p));
        
        // Test minWindow
        String test4s = "ADOBECODEBANC";
        String test4t = "ABC";
        System.out.println("\nMinimum window containing '" + test4t + "' in '" + test4s + "':");
        System.out.println("  Brute Force: " + minWindowBruteForce(test4s, test4t));
        System.out.println("  Optimized: " + minWindow(test4s, test4t));
    }
}
```
