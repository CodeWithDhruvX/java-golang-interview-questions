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
}
