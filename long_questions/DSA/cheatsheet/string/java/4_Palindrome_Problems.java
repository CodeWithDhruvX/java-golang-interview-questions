```java
public class PalindromeProblems {

    // 1. Count Palindromic Substrings
    // Brute Force: O(N^3), Space: O(1)
    public static int countSubstringsBruteForce(String s) {
        int count = 0;
        for (int i = 0; i < s.length(); i++) {
            for (int j = i; j < s.length(); j++) {
                if (isPalindrome(s, i, j)) {
                    count++;
                }
            }
        }
        return count;
    }
    
    private static boolean isPalindrome(String s, int left, int right) {
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
    public static int countSubstrings(String s) {
        int count = 0;
        for (int i = 0; i < s.length(); i++) {
            // Odd length palindromes
            count += countExpand(s, i, i);
            // Even length palindromes
            count += countExpand(s, i, i + 1);
        }
        return count;
    }

    private static int countExpand(String s, int left, int right) {
        int res = 0;
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            res++;
            left--;
            right++;
        }
        return res;
    }

    // 2. Valid Palindrome II (One Deletion)
    // Brute Force: O(N^2), Space: O(N)
    public static boolean validPalindromeBruteForce(String s) {
        // Try removing each character one by one
        for (int i = 0; i < s.length(); i++) {
            String modified = s.substring(0, i) + s.substring(i + 1);
            if (isPalindrome(modified, 0, modified.length() - 1)) {
                return true;
            }
        }
        return isPalindrome(s, 0, s.length() - 1);
    }

    // Optimized: O(N), Space: O(1)
    public static boolean validPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return isPalindromeRange(s, left + 1, right) || isPalindromeRange(s, left, right - 1);
            }
            left++;
            right--;
        }
        return true;
    }

    private static boolean isPalindromeRange(String s, int left, int right) {
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }

    // Test methods
    public static void main(String[] args) {
        // Test countSubstrings
        String test1 = "aaa";
        System.out.println("Count of palindromic substrings in '" + test1 + "':");
        System.out.println("  Brute Force: " + countSubstringsBruteForce(test1));
        System.out.println("  Optimized: " + countSubstrings(test1));
        
        // Test validPalindrome
        String test2 = "abca";
        System.out.println("\nIs '" + test2 + "' valid palindrome with at most one deletion?");
        System.out.println("  Brute Force: " + validPalindromeBruteForce(test2));
        System.out.println("  Optimized: " + validPalindrome(test2));
        
        String test3 = "abc";
        System.out.println("\nIs '" + test3 + "' valid palindrome with at most one deletion?");
        System.out.println("  Brute Force: " + validPalindromeBruteForce(test3));
        System.out.println("  Optimized: " + validPalindrome(test3));
    }
}
```
