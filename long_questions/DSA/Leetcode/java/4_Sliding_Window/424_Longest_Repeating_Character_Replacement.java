public class LongestRepeatingCharacterReplacement {
    
    // 424. Longest Repeating Character Replacement
    // Time: O(N), Space: O(26) = O(1)
    public static int characterReplacement(String s, int k) {
        if (s == null || s.length() == 0) {
            return 0;
        }
        
        int[] count = new int[26];
        int maxCount = 0;
        int left = 0;
        int result = 0;
        
        for (int right = 0; right < s.length(); right++) {
            char c = s.charAt(right);
            count[c - 'A']++;
            maxCount = Math.max(maxCount, count[c - 'A']);
            
            // If window size - maxCount > k, we need to shrink
            while (right - left + 1 - maxCount > k) {
                char leftChar = s.charAt(left);
                count[leftChar - 'A']--;
                left++;
            }
            
            result = Math.max(result, right - left + 1);
        }
        
        return result;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {"ABAB", 2},
            {"AABABBA", 1},
            {"AAAA", 2},
            {"ABC", 0},
            {"ABC", 1},
            {"ABC", 2},
            {"", 1},
            {"A", 0},
            {"A", 1},
            {"ABAA", 0},
            {"ABAA", 1},
            {"ABAA", 2},
            {"BAAAAB", 2},
            {"AAAAAAAAAA", 0},
            {"ABABABABAB", 3}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String s = (String) testCases[i][0];
            int k = (int) testCases[i][1];
            
            int result = characterReplacement(s, k);
            System.out.printf("Test Case %d: s=\"%s\", k=%d -> %d\n", 
                i + 1, s, k, result);
        }
    }
}
