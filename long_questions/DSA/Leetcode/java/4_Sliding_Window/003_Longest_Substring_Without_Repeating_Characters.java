import java.util.HashMap;
import java.util.Map;

public class LongestSubstringWithoutRepeatingCharacters {
    
    // 3. Longest Substring Without Repeating Characters (Variable Size Sliding Window)
    // Time: O(N), Space: O(min(N, M)) where M is the size of character set
    public static int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> charMap = new HashMap<>();
        int left = 0;
        int maxLength = 0;
        
        for (int right = 0; right < s.length(); right++) {
            char currentChar = s.charAt(right);
            
            // If character is already in the window, move left pointer
            if (charMap.containsKey(currentChar) && charMap.get(currentChar) >= left) {
                left = charMap.get(currentChar) + 1;
            }
            
            // Update the character's last seen position
            charMap.put(currentChar, right);
            
            // Update max length
            int currentLength = right - left + 1;
            maxLength = Math.max(maxLength, currentLength);
        }
        
        return maxLength;
    }

    public static void main(String[] args) {
        // Test cases
        String[] testCases = {
            "abcabcbb",
            "bbbbb",
            "pwwkew",
            "",
            "a",
            "au",
            "dvdf",
            "abba",
            "tmmzuxt",
            "abcdefghijklmnopqrstuvwxyz",
            "abccba"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = lengthOfLongestSubstring(testCases[i]);
            System.out.printf("Test Case %d: \"%s\" -> Length of longest substring: %d\n", 
                i + 1, testCases[i], result);
        }
    }
}
