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
}
