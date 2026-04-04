public class FindIndexOfFirstOccurrenceInString {
    
    // 28. Find the Index of the First Occurrence in a String - KMP Algorithm
    // Time: O(N + M), Space: O(M) where N is haystack length, M is needle length
    public int strStr(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        // Build LPS (Longest Prefix Suffix) array
        int[] lps = buildLPS(needle);
        
        int i = 0, j = 0; // i: haystack index, j: needle index
        
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                
                if (j == needle.length()) {
                    return i - j; // Found match
                }
            } else {
                if (j != 0) {
                    j = lps[j - 1]; // Use LPS to skip comparisons
                } else {
                    i++;
                }
            }
        }
        
        return -1; // Not found
    }
    
    // Build LPS array for KMP
    private int[] buildLPS(String pattern) {
        int[] lps = new int[pattern.length()];
        int length = 0; // Length of previous longest prefix suffix
        
        int i = 1;
        while (i < pattern.length()) {
            if (pattern.charAt(i) == pattern.charAt(length)) {
                length++;
                lps[i] = length;
                i++;
            } else {
                if (length != 0) {
                    length = lps[length - 1];
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        
        return lps;
    }
    
    // KMP with detailed tracing
    public class KMPResult {
        int index;
        java.util.List<String> trace;
        
        KMPResult(int index, java.util.List<String> trace) {
            this.index = index;
            this.trace = trace;
        }
    }
    
    public KMPResult strStrKMPDetailed(String haystack, String needle) {
        java.util.List<String> trace = new java.util.ArrayList<>();
        
        if (needle.isEmpty()) {
            trace.add("Empty needle, returning 0");
            return new KMPResult(0, trace);
        }
        
        trace.add("Building LPS for: " + needle);
        
        int[] lps = buildLPSWithTrace(needle, trace);
        trace.add("LPS array: " + java.util.Arrays.toString(lps));
        
        int i = 0, j = 0;
        trace.add(String.format("Starting search: i=%d, j=%d", i, j));
        
        while (i < haystack.length()) {
            trace.add(String.format("Comparing haystack[%d]='%c' with needle[%d]='%c'", 
                i, haystack.charAt(i), j, needle.charAt(j)));
            
            if (haystack.charAt(i) == needle.charAt(j)) {
                trace.add("  Match! Incrementing both indices");
                i++;
                j++;
                
                if (j == needle.length()) {
                    int result = i - j;
                    trace.add(String.format("  Found complete match at index: %d", result));
                    return new KMPResult(result, trace);
                }
            } else {
                if (j != 0) {
                    trace.add(String.format("  Mismatch! Using LPS[%d-1] = %d", j, lps[j-1]));
                    j = lps[j - 1];
                } else {
                    trace.add("  Mismatch! No prefix to fall back, incrementing i");
                    i++;
                }
            }
            
            trace.add(String.format("  New state: i=%d, j=%d", i, j));
        }
        
        trace.add("Reached end of haystack, no match found");
        return new KMPResult(-1, trace);
    }
    
    private int[] buildLPSWithTrace(String pattern, java.util.List<String> trace) {
        int[] lps = new int[pattern.length()];
        int length = 0;
        
        trace.add("Building LPS array:");
        
        for (int i = 1; i < pattern.length(); i++) {
            trace.add(String.format("  Processing position %d: '%c'", i, pattern.charAt(i)));
            
            while (length > 0 && pattern.charAt(i) != pattern.charAt(length)) {
                trace.add(String.format("    Mismatch! Falling back from length=%d to LPS[%d-1]=%d", 
                    length, length, lps[length - 1]));
                length = lps[length - 1];
            }
            
            if (pattern.charAt(i) == pattern.charAt(length)) {
                length++;
                trace.add(String.format("    Match! Length increased to %d", length));
            }
            
            lps[i] = length;
            trace.add(String.format("    LPS[%d] = %d", i, lps[i]));
        }
        
        return lps;
    }
    
    // Standard Java approach for comparison
    public int strStrStandard(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            boolean found = true;
            
            for (int j = 0; j < needle.length(); j++) {
                if (i + j >= haystack.length() || haystack.charAt(i + j) != needle.charAt(j)) {
                    found = false;
                    break;
                }
            }
            
            if (found) {
                return i;
            }
        }
        
        return -1;
    }
    
    // Rabin-Karp algorithm for comparison
    public int strStrRabinKarp(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        int n = haystack.length();
        int m = needle.length();
        
        if (m > n) {
            return -1;
        }
        
        final int BASE = 256;
        final int MOD = 101; // A prime number
        
        // Calculate hash for needle and first window of haystack
        long needleHash = 0;
        long haystackHash = 0;
        long power = 1;
        
        for (int i = 0; i < m; i++) {
            needleHash = (needleHash * BASE + needle.charAt(i)) % MOD;
            haystackHash = (haystackHash * BASE + haystack.charAt(i)) % MOD;
            
            if (i < m - 1) {
                power = (power * BASE) % MOD;
            }
        }
        
        // Slide the window
        for (int i = 0; i <= n - m; i++) {
            if (needleHash == haystackHash) {
                // Check for actual match (to handle hash collisions)
                boolean match = true;
                for (int j = 0; j < m; j++) {
                    if (haystack.charAt(i + j) != needle.charAt(j)) {
                        match = false;
                        break;
                    }
                }
                
                if (match) {
                    return i;
                }
            }
            
            // Calculate hash for next window
            if (i < n - m) {
                haystackHash = (haystackHash - haystack.charAt(i) * power + MOD) % MOD;
                haystackHash = (haystackHash * BASE + haystack.charAt(i + m)) % MOD;
                if (haystackHash < 0) {
                    haystackHash += MOD;
                }
            }
        }
        
        return -1;
    }
    
    // Boyer-Moore algorithm for comparison
    public int strStrBoyerMoore(String haystack, String needle) {
        if (needle.isEmpty()) {
            return 0;
        }
        
        int n = haystack.length();
        int m = needle.length();
        
        if (m > n) {
            return -1;
        }
        
        // Build bad character table
        int[] badChar = new int[256];
        for (int i = 0; i < 256; i++) {
            badChar[i] = -1;
        }
        
        for (int i = 0; i < m; i++) {
            badChar[needle.charAt(i)] = i;
        }
        
        int shift = 0;
        while (shift <= n - m) {
            int j = m - 1;
            
            while (j >= 0 && needle.charAt(j) == haystack.charAt(shift + j)) {
                j--;
            }
            
            if (j < 0) {
                return shift; // Found match
            } else {
                shift += Math.max(1, j - badChar[haystack.charAt(shift + j)]);
            }
        }
        
        return -1;
    }
    
    // Performance comparison
    public void comparePerformance(String haystack, String needle) {
        System.out.println("=== Performance Comparison ===");
        System.out.printf("Haystack length: %d, Needle length: %d\n", haystack.length(), needle.length());
        
        // Standard approach
        long startTime = System.nanoTime();
        int result1 = strStrStandard(haystack, needle);
        long endTime = System.nanoTime();
        System.out.printf("Standard approach: %d (took %d ns)\n", result1, endTime - startTime);
        
        // KMP
        startTime = System.nanoTime();
        int result2 = strStr(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("KMP algorithm: %d (took %d ns)\n", result2, endTime - startTime);
        
        // Rabin-Karp
        startTime = System.nanoTime();
        int result3 = strStrRabinKarp(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("Rabin-Karp: %d (took %d ns)\n", result3, endTime - startTime);
        
        // Boyer-Moore
        startTime = System.nanoTime();
        int result4 = strStrBoyerMoore(haystack, needle);
        endTime = System.nanoTime();
        System.out.printf("Boyer-Moore: %d (took %d ns)\n", result4, endTime - startTime);
    }
    
    // Find all occurrences using KMP
    public java.util.List<Integer> findAllOccurrences(String haystack, String needle) {
        java.util.List<Integer> occurrences = new java.util.ArrayList<>();
        
        if (needle.isEmpty()) {
            for (int i = 0; i <= haystack.length(); i++) {
                occurrences.add(i);
            }
            return occurrences;
        }
        
        int[] lps = buildLPS(needle);
        int i = 0, j = 0;
        
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                
                if (j == needle.length()) {
                    occurrences.add(i - j);
                    j = lps[j - 1];
                }
            } else {
                if (j != 0) {
                    j = lps[j - 1];
                } else {
                    i++;
                }
            }
        }
        
        return occurrences;
    }
    
    public static void main(String[] args) {
        FindIndexOfFirstOccurrenceInString finder = new FindIndexOfFirstOccurrenceInString();
        
        // Test cases
        String[][] testCases = {
            {"sadbutsad", "sad"},
            {"leetcode", "leeto"},
            {"hello", "ll"},
            {"", ""},
            {"", "a"},
            {"a", ""},
            {"aaaaa", "bba"},
            {"abc", "abc"},
            {"abc", "abcd"},
            {"mississippi", "issi"},
            {"ABABDABACDABABCABAB", "ABABCABAB"}
        };
        
        String[] descriptions = {
            "Standard case",
            "No match",
            "Middle match",
            "Both empty",
            "Empty haystack",
            "Empty needle",
            "No occurrence",
            "Exact match",
            "Needle longer",
            "Multiple occurrences",
            "Complex pattern"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Haystack: \"%s\", Needle: \"%s\"\n", testCases[i][0], testCases[i][1]);
            
            int result1 = finder.strStr(testCases[i][0], testCases[i][1]);
            int result2 = finder.strStrStandard(testCases[i][0], testCases[i][1]);
            int result3 = finder.strStrRabinKarp(testCases[i][0], testCases[i][1]);
            int result4 = finder.strStrBoyerMoore(testCases[i][0], testCases[i][1]);
            
            System.out.printf("KMP: %d\n", result1);
            System.out.printf("Standard: %d\n", result2);
            System.out.printf("Rabin-Karp: %d\n", result3);
            System.out.printf("Boyer-Moore: %d\n", result4);
            
            // Find all occurrences
            java.util.List<Integer> allOccurrences = finder.findAllOccurrences(testCases[i][0], testCases[i][1]);
            System.out.printf("All occurrences: %s\n", allOccurrences);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed KMP Explanation ===");
        KMPResult detailedResult = finder.strStrKMPDetailed("sadbutsad", "sad");
        System.out.printf("Result: %d\n", detailedResult.index);
        for (String step : detailedResult.trace) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        String largeHaystack = "ab".repeat(1000) + "cde";
        String largeNeedle = "cde";
        finder.comparePerformance(largeHaystack, largeNeedle);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Unicode: %d\n", finder.strStr("héllo", "él"));
        System.out.printf("Case sensitive: %d\n", finder.strStr("Hello", "hello"));
        System.out.printf("Repeated pattern: %d\n", finder.strStr("aaaaaa", "aaa"));
    }
}
