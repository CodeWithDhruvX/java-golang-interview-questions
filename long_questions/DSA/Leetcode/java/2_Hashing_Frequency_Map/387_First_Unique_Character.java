import java.util.*;

public class FirstUniqueCharacterInAString {
    
    // 387. First Unique Character in a String
    // Time: O(N), Space: O(26) = O(1)
    public static int firstUniqChar(String s) {
        if (s == null || s.length() == 0) {
            return -1;
        }
        
        // Count frequency of each character
        int[] count = new int[26];
        for (char c : s.toCharArray()) {
            count[c - 'a']++;
        }
        
        // Find first unique character
        for (int i = 0; i < s.length(); i++) {
            if (count[s.charAt(i) - 'a'] == 1) {
                return i;
            }
        }
        
        return -1;
    }

    public static void main(String[] args) {
        String[] testCases = {
            "leetcode",
            "loveleetcode",
            "aabb",
            "abcabc",
            "",
            "a",
            "aa",
            "abacabad",
            "abbbcd",
            "z",
            "zz",
            "abcde",
            "aabbccddeeff",
            "abcdefga",
            "abcbad"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String s = testCases[i];
            int result = firstUniqChar(s);
            
            System.out.printf("Test Case %d: \"%s\" -> %d\n", 
                i + 1, s, result);
        }
    }
}
