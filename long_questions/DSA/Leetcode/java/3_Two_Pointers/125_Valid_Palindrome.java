public class ValidPalindrome {
    
    // 125. Valid Palindrome
    // Time: O(N), Space: O(1)
    public static boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        
        while (left < right) {
            // Skip non-alphanumeric characters
            while (left < right && !isAlphanumeric(s.charAt(left))) {
                left++;
            }
            while (left < right && !isAlphanumeric(s.charAt(right))) {
                right--;
            }
            
            // Compare characters (case-insensitive)
            if (left < right && Character.toLowerCase(s.charAt(left)) != Character.toLowerCase(s.charAt(right))) {
                return false;
            }
            
            left++;
            right--;
        }
        
        return true;
    }

    // Helper function to check if character is alphanumeric
    private static boolean isAlphanumeric(char c) {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9');
    }

    public static void main(String[] args) {
        // Test cases
        String[] testCases = {
            "A man, a plan, a canal: Panama",
            "race a car",
            " ",
            "",
            "madam",
            "Able was I ere I saw Elba",
            "No lemon, no melon",
            "12321",
            "12345",
            ".,"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = isPalindrome(testCases[i]);
            System.out.printf("Test Case %d: \"%s\" -> Palindrome: %b\n", 
                i + 1, testCases[i], result);
        }
    }
}
