import java.util.*;

public class GroupAnagrams {
    
    // 49. Group Anagrams
    // Time: O(N*K*logK), Space: O(N*K)
    public static List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> anagramMap = new HashMap<>();
        
        for (String str : strs) {
            // Sort the string to create a key
            String sortedStr = sortString(str);
            anagramMap.computeIfAbsent(sortedStr, k -> new ArrayList<>()).add(str);
        }
        
        return new ArrayList<>(anagramMap.values());
    }

    // Helper function to sort a string
    private static String sortString(String s) {
        char[] chars = s.toCharArray();
        Arrays.sort(chars);
        return new String(chars);
    }

    // Alternative solution using character count as key (more efficient)
    public static List<List<String>> groupAnagramsOptimized(String[] strs) {
        Map<String, List<String>> anagramMap = new HashMap<>();
        
        for (String str : strs) {
            // Create key based on character count
            String key = createCountKey(str);
            anagramMap.computeIfAbsent(key, k -> new ArrayList<>()).add(str);
        }
        
        return new ArrayList<>(anagramMap.values());
    }

    // Create key based on character count (26 lowercase letters)
    private static String createCountKey(String s) {
        int[] count = new int[26];
        for (char c : s.toCharArray()) {
            count[c - 'a']++;
        }
        
        return Arrays.toString(count);
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"eat", "tea", "tan", "ate", "nat", "bat"},
            {""},
            {"a"},
            {"abc", "bca", "cab", "def", "fed", "ghi"},
            {"", "", ""},
            {"a", "b", "c", "a", "b", "c"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<List<String>> result = groupAnagrams(testCases[i]);
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("Grouped Anagrams: %s\n\n", result);
        }
    }
}
