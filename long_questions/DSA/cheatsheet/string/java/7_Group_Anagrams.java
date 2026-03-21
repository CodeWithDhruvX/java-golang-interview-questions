```java
import java.util.*;

public class GroupAnagrams {

    // 1. Group Anagrams
    // Brute Force: O(N^2 * K log K), Space: O(N * K)
    public static List<List<String>> groupAnagramsBruteForce(String[] strs) {
        List<List<String>> result = new ArrayList<>();
        boolean[] used = new boolean[strs.length];
        
        for (int i = 0; i < strs.length; i++) {
            if (used[i]) continue;
            
            List<String> group = new ArrayList<>();
            group.add(strs[i]);
            used[i] = true;
            
            for (int j = i + 1; j < strs.length; j++) {
                if (!used[j] && areAnagrams(strs[i], strs[j])) {
                    group.add(strs[j]);
                    used[j] = true;
                }
            }
            result.add(group);
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

    // Optimized: O(N * K), Space: O(N * K)
    public static List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strs) {
            int[] count = new int[26];
            for (int i = 0; i < s.length(); i++) {
                count[s.charAt(i) - 'a']++;
            }
            
            // Create a unique key for the frequency count
            StringBuilder keyBuilder = new StringBuilder();
            for (int i = 0; i < 26; i++) {
                keyBuilder.append('#');
                keyBuilder.append(count[i]);
            }
            String key = keyBuilder.toString();
            
            groups.putIfAbsent(key, new ArrayList<>());
            groups.get(key).add(s);
        }
        
        return new ArrayList<>(groups.values());
    }

    // Alternative implementation using sorted string as key
    // Time: O(N * K log K), Space: O(N * K)
    public static List<List<String>> groupAnagramsSorted(String[] strs) {
        Map<String, List<String>> groups = new HashMap<>();
        
        for (String s : strs) {
            char[] chars = s.toCharArray();
            Arrays.sort(chars);
            String sorted = new String(chars);
            
            groups.putIfAbsent(sorted, new ArrayList<>());
            groups.get(sorted).add(s);
        }
        
        return new ArrayList<>(groups.values());
    }

    // Test methods
    public static void main(String[] args) {
        // Test groupAnagrams
        String[] test1 = {"eat", "tea", "tan", "ate", "nat", "bat"};
        System.out.println("Grouping anagrams of " + Arrays.toString(test1) + ":");
        System.out.println("  Brute Force:");
        List<List<String>> result1Brute = groupAnagramsBruteForce(test1);
        for (List<String> group : result1Brute) {
            System.out.println("    " + group);
        }
        System.out.println("  Optimized:");
        List<List<String>> result1 = groupAnagrams(test1);
        for (List<String> group : result1) {
            System.out.println("    " + group);
        }
        
        // Test groupAnagramsSorted
        String[] test2 = {"abc", "bca", "cab", "def", "fed"};
        System.out.println("\nGrouping anagrams (sorted method) of " + Arrays.toString(test2) + ":");
        List<List<String>> result2 = groupAnagramsSorted(test2);
        for (List<String> group : result2) {
            System.out.println("    " + group);
        }
    }
}
```
