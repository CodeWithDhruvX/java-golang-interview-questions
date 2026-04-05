public class PermutationInString {
    
    // 567. Permutation in String
    // Time: O(N + M), Space: O(26) = O(1)
    public static boolean checkInclusion(String s1, String s2) {
        if (s1 == null || s2 == null || s1.length() > s2.length()) {
            return false;
        }
        
        int[] count = new int[26];
        int m = s1.length();
        int n = s2.length();
        
        // Initialize count with s1 characters
        for (int i = 0; i < m; i++) {
            count[s1.charAt(i) - 'a']++;
            count[s2.charAt(i) - 'a']--;
        }
        
        // Check if first window is a permutation
        if (allZeros(count)) {
            return true;
        }
        
        // Slide window through s2
        for (int i = m; i < n; i++) {
            // Add new character
            count[s2.charAt(i) - 'a']--;
            // Remove old character
            count[s2.charAt(i - m) - 'a']++;
            
            if (allZeros(count)) {
                return true;
            }
        }
        
        return false;
    }
    
    private static boolean allZeros(int[] count) {
        for (int c : count) {
            if (c != 0) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        String[][] testCases = {
            {"ab", "eidbaooo"},
            {"ab", "eidboaoo"},
            {"a", "a"},
            {"a", "b"},
            {"abc", "ccccbbbbaaaa"},
            {"abc", "bca"},
            {"abc", "cba"},
            {"abcd", "eidbaooo"},
            {"ab", "ba"},
            {"ab", "ab"},
            {"abc", "defghijklmno"},
            {"", "abc"},
            {"abc", ""},
            {"", ""},
            {"abc", "abc"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String s1 = testCases[i][0];
            String s2 = testCases[i][1];
            
            boolean result = checkInclusion(s1, s2);
            System.out.printf("Test Case %d: s1=\"%s\", s2=\"%s\" -> %b\n", 
                i + 1, s1, s2, result);
        }
    }
}
