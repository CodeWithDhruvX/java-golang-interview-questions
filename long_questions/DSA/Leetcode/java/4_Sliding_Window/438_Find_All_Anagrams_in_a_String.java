import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class FindAllAnagramsInAString {
    
    // 438. Find All Anagrams in a String (Fixed Size Sliding Window)
    // Time: O(N), Space: O(1) for ASCII characters
    public static List<Integer> findAnagrams(String s, String p) {
        List<Integer> result = new ArrayList<>();
        
        if (s.length() < p.length()) {
            return result;
        }
        
        int[] pCount = new int[26];
        int[] sCount = new int[26];
        
        // Initialize frequency count for pattern and first window
        for (int i = 0; i < p.length(); i++) {
            pCount[p.charAt(i) - 'a']++;
            sCount[s.charAt(i) - 'a']++;
        }
        
        // Check if first window is an anagram
        if (matches(pCount, sCount)) {
            result.add(0);
        }
        
        // Slide the window through the string
        for (int i = p.length(); i < s.length(); i++) {
            // Remove the leftmost character
            sCount[s.charAt(i - p.length()) - 'a']--;
            // Add the new character
            sCount[s.charAt(i) - 'a']++;
            
            // Check if current window is an anagram
            if (matches(pCount, sCount)) {
                result.add(i - p.length() + 1);
            }
        }
        
        return result;
    }

    // Helper function to check if two frequency arrays match
    private static boolean matches(int[] pCount, int[] sCount) {
        for (int i = 0; i < 26; i++) {
            if (pCount[i] != sCount[i]) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"cbaebabacd", "abc"},
            {"abab", "ab"},
            {"aaaaaaaaaa", "aaaa"},
            {"abacbabc", "abc"},
            {"", "a"},
            {"a", ""},
            {"abc", "def"},
            {"abababab", "ab"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<Integer> result = findAnagrams(testCases[i][0], testCases[i][1]);
            System.out.printf("Test Case %d: s=\"%s\", p=\"%s\" -> Anagram indices: %s\n", 
                i + 1, testCases[i][0], testCases[i][1], result);
        }
    }
}
