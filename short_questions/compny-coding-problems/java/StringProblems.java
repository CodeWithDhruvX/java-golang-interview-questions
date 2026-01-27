package java_solutions;

import java.util.*;

public class StringProblems {

    // 1. Reverse a string
    public static String reverseString(String str) {
        StringBuilder reversedStr = new StringBuilder();
        for (int i = str.length() - 1; i >= 0; i--) {
            reversedStr.append(str.charAt(i));
        }
        return reversedStr.toString();
    }

    // 2. Reverse words in a sentence
    public static String reverseWords(String sentence) {
        String[] words = sentence.split(" ");
        StringBuilder result = new StringBuilder();
        for (int i = words.length - 1; i >= 0; i--) {
            result.append(words[i]).append(" ");
        }
        return result.toString().trim();
    }

    // 3. Check if a string is palindrome
    public static boolean isPalindrome(String str) {
        int start = 0, end = str.length() - 1;
        while (start < end) {
            if (str.charAt(start) != str.charAt(end)) {
                return false;
            }
            start++;
            end--;
        }
        return true;
    }

    // 4. Check if two strings are anagrams
    public static boolean areAnagrams(String str1, String str2) {
        if (str1.length() != str2.length()) {
            return false;
        }
        char[] s1 = str1.toCharArray();
        char[] s2 = str2.toCharArray();
        Arrays.sort(s1);
        Arrays.sort(s2);
        return Arrays.equals(s1, s2);
    }

    // 5. Count vowels and consonants
    public static void countVowelsConsonants(String str) {
        int vowels = 0, consonants = 0;
        str = str.toLowerCase();
        for (char ch : str.toCharArray()) {
            if (Character.isLetter(ch)) {
                if ("aeiou".indexOf(ch) != -1) {
                    vowels++;
                } else {
                    consonants++;
                }
            }
        }
        System.out.println("Vowels: " + vowels + ", Consonants: " + consonants);
    }

    // 6. Count frequency of characters
    public static void charFrequency(String str) {
        Map<Character, Integer> freqMap = new HashMap<>();
        for (char ch : str.toCharArray()) {
            freqMap.put(ch, freqMap.getOrDefault(ch, 0) + 1);
        }
        System.out.println(freqMap);
    }

    // 7. Find first non-repeating character
    public static String firstNonRepeating(String str) {
        Map<Character, Integer> freqMap = new LinkedHashMap<>();
        for (char ch : str.toCharArray()) {
            freqMap.put(ch, freqMap.getOrDefault(ch, 0) + 1);
        }
        for (char ch : str.toCharArray()) {
            if (freqMap.get(ch) == 1) {
                return String.valueOf(ch);
            }
        }
        return "NULL";
    }

    // 8. Remove duplicate characters
    public static String removeDuplicates(String str) {
        Set<Character> seen = new HashSet<>();
        StringBuilder result = new StringBuilder();
        for (char ch : str.toCharArray()) {
            if (!seen.contains(ch)) {
                seen.add(ch);
                result.append(ch);
            }
        }
        return result.toString();
    }

    // 9. Replace spaces with special character
    public static String replaceSpaces(String str, String specialChar) {
        StringBuilder result = new StringBuilder();
        for (char ch : str.toCharArray()) {
            if (ch == ' ') {
                result.append(specialChar);
            } else {
                result.append(ch);
            }
        }
        return result.toString();
    }

    // 10. Convert lowercase to uppercase (without built-in)
    public static String toUpperCase(String str) {
        StringBuilder result = new StringBuilder();
        for (char ch : str.toCharArray()) {
            if (ch >= 'a' && ch <= 'z') {
                result.append((char) (ch - 32));
            } else {
                result.append(ch);
            }
        }
        return result.toString();
    }

    // 11. Find longest word in a string
    public static String longestWord(String sentence) {
        String[] words = sentence.split(" ");
        int maxLen = 0;
        String longest = "";
        for (String word : words) {
            if (word.length() > maxLen) {
                maxLen = word.length();
                longest = word;
            }
        }
        return longest;
    }

    // 12. Count number of words
    public static int countWords(String sentence) {
        if (sentence == null || sentence.trim().isEmpty()) {
            return 0;
        }
        String[] words = sentence.trim().split("\\s+");
        return words.length;
    }

    // 13. Check substring present or not
    public static boolean isSubstring(String mainStr, String subStr) {
        return mainStr.contains(subStr);
    }

    // 14. Remove vowels from string
    public static String removeVowels(String str) {
        return str.replaceAll("[aeiouAEIOU]", "");
    }

    // 15. Sort characters in a string
    public static String sortString(String str) {
        char[] arr = str.toCharArray();
        Arrays.sort(arr);
        return new String(arr);
    }

    // 16. Find duplicate characters
    public static void findDuplicates(String str) {
        Map<Character, Integer> freqMap = new HashMap<>();
        for (char ch : str.toCharArray()) {
            freqMap.put(ch, freqMap.getOrDefault(ch, 0) + 1);
        }
        for (Map.Entry<Character, Integer> entry : freqMap.entrySet()) {
            if (entry.getValue() > 1) {
                System.out.print(entry.getKey() + " ");
            }
        }
        System.out.println();
    }

    // 17. Reverse string using recursion
    public static String reverseRecursive(String str) {
        if (str.isEmpty()) {
            return "";
        }
        return reverseRecursive(str.substring(1)) + str.charAt(0);
    }

    // 18. Print string in zig-zag format
    public static String printZigZag(String str, int k) {
        if (k == 1) return str;
        StringBuilder[] rows = new StringBuilder[k];
        for (int i = 0; i < k; i++) rows[i] = new StringBuilder();
        int row = 0;
        boolean down = true;
        for (char ch : str.toCharArray()) {
            rows[row].append(ch);
            if (row == 0) down = true;
            else if (row == k - 1) down = false;
            if (down) row++;
            else row--;
        }
        StringBuilder result = new StringBuilder();
        for (StringBuilder r : rows) result.append(r);
        return result.toString();
    }

    // 19. Check string rotation
    public static boolean isRotation(String str1, String str2) {
        if (str1.length() != str2.length()) {
            return false;
        }
        String temp = str1 + str1;
        return temp.contains(str2);
    }

    // 20. Compare two strings without using equals()
    public static boolean compareStrings(String str1, String str2) {
        if (str1.length() != str2.length()) {
            return false;
        }
        for (int i = 0; i < str1.length(); i++) {
            if (str1.charAt(i) != str2.charAt(i)) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        System.out.println("1. Reverse String: hello -> " + reverseString("hello"));
        System.out.println("2. Reverse Words: Hello World -> " + reverseWords("Hello World"));
        System.out.println("3. Is Palindrome: madam -> " + isPalindrome("madam"));
        System.out.println("4. Are Anagrams: listen, silent -> " + areAnagrams("listen", "silent"));
        System.out.print("5. Count Vowels Consonants: Hello -> "); countVowelsConsonants("Hello");
        System.out.print("6. Char Frequency: banana -> "); charFrequency("banana");
        System.out.println("7. First Non-Repeating: swiss -> " + firstNonRepeating("swiss"));
        System.out.println("8. Remove Duplicates: banana -> " + removeDuplicates("banana"));
        System.out.println("9. Replace Spaces: Hello World -> " + replaceSpaces("Hello World", "-"));
        System.out.println("10. To Upper Case: java -> " + toUpperCase("java"));
        System.out.println("11. Longest Word: I love programming -> " + longestWord("I love programming"));
        System.out.println("12. Count Words: Hello World -> " + countWords("Hello World"));
        System.out.println("13. Is Substring: Hello World, World -> " + isSubstring("Hello World", "World"));
        System.out.println("14. Remove Vowels: Hello World -> " + removeVowels("Hello World"));
        System.out.println("15. Sort String: edcba -> " + sortString("edcba"));
        System.out.print("16. Find Duplicates: banana -> "); findDuplicates("banana");
        System.out.println("17. Reverse Recursive: recursion -> " + reverseRecursive("recursion"));
        System.out.println("18. Print Zig Zag: PAYPALISHIRING, 3 -> " + printZigZag("PAYPALISHIRING", 3));
        System.out.println("19. Is Rotation: ABCD, CDAB -> " + isRotation("ABCD", "CDAB"));
        System.out.println("20. Compare Strings: abc, abc -> " + compareStrings("abc", "abc"));
    }
}
