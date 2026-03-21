```java
public class StringBasics {

    // 1. Reverse a String
    // Brute Force: O(N^2), Space: O(N)
    public static String reverseStringBruteForce(String s) {
        String result = "";
        for (int i = s.length() - 1; i >= 0; i--) {
            result += s.charAt(i);
        }
        return result;
    }

    // Optimized: O(N), Space: O(N)
    public static String reverseString(String s) {
        char[] chars = s.toCharArray();
        int left = 0, right = chars.length - 1;
        while (left < right) {
            char temp = chars[left];
            chars[left] = chars[right];
            chars[right] = temp;
            left++;
            right--;
        }
        return new String(chars);
    }

    // 2. Check if String is Palindrome
    // Brute Force: O(N^2), Space: O(N)
    public static boolean isPalindromeBruteForce(String s) {
        String reversed = reverseStringBruteForce(s);
        return s.equals(reversed);
    }

    // Optimized: O(N), Space: O(1)
    public static boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }

    // 3. Count Vowels and Consonants
    // Brute Force: O(N^2), Space: O(1)
    public static int[] countVowelsConsonantsBruteForce(String s) {
        int vowels = 0;
        int consonants = 0;
        String vowelsStr = "aeiou";
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isLetter(c)) {
                char lowerChar = Character.toLowerCase(c);
                boolean isVowel = false;
                for (int j = 0; j < vowelsStr.length(); j++) {
                    if (vowelsStr.charAt(j) == lowerChar) {
                        isVowel = true;
                        break;
                    }
                }
                if (isVowel) {
                    vowels++;
                } else {
                    consonants++;
                }
            }
        }
        return new int[]{vowels, consonants};
    }

    // Optimized: O(N), Space: O(1)
    public static int[] countVowelsConsonants(String s) {
        int vowels = 0;
        int consonants = 0;
        String vowelsStr = "aeiou";
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isLetter(c)) {
                char lowerChar = Character.toLowerCase(c);
                if (vowelsStr.indexOf(lowerChar) != -1) {
                    vowels++;
                } else {
                    consonants++;
                }
            }
        }
        return new int[]{vowels, consonants};
    }

    // Test methods
    public static void main(String[] args) {
        // Test reverseString
        String test1 = "hello";
        System.out.println("Reverse of '" + test1 + "':");
        System.out.println("  Brute Force: " + reverseStringBruteForce(test1));
        System.out.println("  Optimized: " + reverseString(test1));
        
        // Test isPalindrome
        String test2 = "racecar";
        System.out.println("\nIs '" + test2 + "' palindrome?");
        System.out.println("  Brute Force: " + isPalindromeBruteForce(test2));
        System.out.println("  Optimized: " + isPalindrome(test2));
        
        // Test countVowelsConsonants
        String test3 = "Hello World";
        int[] resultBrute = countVowelsConsonantsBruteForce(test3);
        int[] resultOpt = countVowelsConsonants(test3);
        System.out.println("\nVowels and consonants in '" + test3 + "':");
        System.out.println("  Brute Force - Vowels: " + resultBrute[0] + ", Consonants: " + resultBrute[1]);
        System.out.println("  Optimized - Vowels: " + resultOpt[0] + ", Consonants: " + resultOpt[1]);
    }
}
```
