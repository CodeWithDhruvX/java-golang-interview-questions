package java_solutions;

import java.util.*;

public class AdditionalStrings {

    // 1. Find longest palindrome substring
    public static String longestPalindrome(String s) {
        if (s == null || s.length() < 1)
            return "";
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int len = Math.max(len1, len2);
            if (len > end - start) {
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        return s.substring(start, end + 1);
    }

    private static int expandAroundCenter(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1;
    }

    // 2. Find longest common prefix
    public static String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length == 0)
            return "";
        String prefix = strs[0];
        for (int i = 1; i < strs.length; i++) {
            while (strs[i].indexOf(prefix) != 0) {
                prefix = prefix.substring(0, prefix.length() - 1);
                if (prefix.isEmpty())
                    return "";
            }
        }
        return prefix;
    }

    // 3. Check if string contains only digits
    public static boolean isDigitsOnly(String str) {
        for (char ch : str.toCharArray()) {
            if (!Character.isDigit(ch))
                return false;
        }
        return true;
    }

    // 4. Count uppercase & lowercase letters
    public static void countCase(String str) {
        int upper = 0, lower = 0;
        for (char ch : str.toCharArray()) {
            if (Character.isUpperCase(ch))
                upper++;
            else if (Character.isLowerCase(ch))
                lower++;
        }
        System.out.println("Upper: " + upper + ", Lower: " + lower);
    }

    // 5. Remove special characters
    public static String removeSpecialChars(String str) {
        return str.replaceAll("[^a-zA-Z0-9]", "");
    }

    // 6. Find all permutations of string (basic)
    public static void permute(String str, int l, int r) {
        if (l == r)
            System.out.println(str);
        else {
            for (int i = l; i <= r; i++) {
                str = swap(str, l, i);
                permute(str, l + 1, r);
                str = swap(str, l, i);
            }
        }
    }

    private static String swap(String str, int i, int j) {
        char[] arr = str.toCharArray();
        char temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
        return new String(arr);
    }

    // 7. Check valid parentheses
    public static boolean isValidParentheses(String s) {
        Stack<Character> stack = new Stack<>();
        for (char c : s.toCharArray()) {
            if (c == '(' || c == '{' || c == '[')
                stack.push(c);
            else {
                if (stack.isEmpty())
                    return false;
                char top = stack.pop();
                if ((c == ')' && top != '(') ||
                        (c == '}' && top != '{') ||
                        (c == ']' && top != '['))
                    return false;
            }
        }
        return stack.isEmpty();
    }

    // 8. Find duplicate words in sentence
    public static void findDuplicateWords(String sentence) {
        String[] words = sentence.toLowerCase().split("\\s+");
        Map<String, Integer> map = new HashMap<>();
        for (String word : words) {
            map.put(word, map.getOrDefault(word, 0) + 1);
        }
        for (Map.Entry<String, Integer> entry : map.entrySet()) {
            if (entry.getValue() > 1) {
                System.out.print(entry.getKey() + " ");
            }
        }
        System.out.println();
    }

    // 9. Reverse each word in place
    public static String reverseEachWord(String sentence) {
        String[] words = sentence.split(" ");
        StringBuilder result = new StringBuilder();
        for (String word : words) {
            result.append(new StringBuilder(word).reverse()).append(" ");
        }
        return result.toString().trim();
    }

    // 10. Check isomorphic strings
    public static boolean isIsomorphic(String s, String t) {
        if (s.length() != t.length())
            return false;
        Map<Character, Character> mapS = new HashMap<>();
        Map<Character, Character> mapT = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char ss = s.charAt(i);
            char tt = t.charAt(i);
            if ((mapS.containsKey(ss) && mapS.get(ss) != tt) ||
                    (mapT.containsKey(tt) && mapT.get(tt) != ss)) {
                return false;
            }
            mapS.put(ss, tt);
            mapT.put(tt, ss);
        }
        return true;
    }

    // 11. Check pangram
    public static boolean isPangram(String str) {
        Set<Character> set = new HashSet<>();
        for (char ch : str.toLowerCase().toCharArray()) {
            if (Character.isLetter(ch)) {
                set.add(ch);
            }
        }
        return set.size() == 26;
    }

    // 12. Print all substrings
    public static void printSubstrings(String str) {
        for (int i = 0; i < str.length(); i++) {
            for (int j = i + 1; j <= str.length(); j++) {
                System.out.println(str.substring(i, j));
            }
        }
    }

    // 13. Remove consecutive duplicates
    public static String removeConsecutiveDeep(String str) {
        if (str.length() < 2)
            return str;
        StringBuilder sb = new StringBuilder();
        sb.append(str.charAt(0));
        for (int i = 1; i < str.length(); i++) {
            if (str.charAt(i) != str.charAt(i - 1)) {
                sb.append(str.charAt(i));
            }
        }
        return sb.toString();
    }

    // 14. Check if strings differ by one character
    public static boolean differByOne(String s1, String s2) {
        if (Math.abs(s1.length() - s2.length()) > 1)
            return false;
        int count = 0;
        int i = 0, j = 0;
        while (i < s1.length() && j < s2.length()) {
            if (s1.charAt(i) != s2.charAt(j)) {
                count++;
                if (count > 1)
                    return false;
                if (s1.length() > s2.length())
                    i++;
                else if (s2.length() > s1.length())
                    j++;
                else {
                    i++;
                    j++;
                }
            } else {
                i++;
                j++;
            }
        }
        if (i < s1.length() || j < s2.length())
            count++;
        return count == 1;
    }

    // 15. Find smallest & largest word
    public static void minMaxWords(String sentence) {
        String[] words = sentence.split("\\s+");
        if (words.length == 0)
            return;
        String min = words[0], max = words[0];
        for (String w : words) {
            if (w.length() < min.length())
                min = w;
            if (w.length() > max.length())
                max = w;
        }
        System.out.println("Min: " + min + ", Max: " + max);
    }

    public static void main(String[] args) {
        System.out.println("1. Longest Palindrome: babad -> " + longestPalindrome("babad"));
        System.out.println(
                "2. Longest Common Prefix: " + longestCommonPrefix(new String[] { "flower", "flow", "flight" }));
        System.out.println("3. Is Digits Only: 12345 -> " + isDigitsOnly("12345"));
        System.out.print("4. Count Case: Hello World -> ");
        countCase("Hello World");
        System.out.println("5. Remove Special Chars: $Hem$lo_World -> " + removeSpecialChars("$Hem$lo_World"));
        System.out.println("6. Permutations: ABC");
        permute("ABC", 0, 2);
        System.out.println("7. Valid Parentheses: ((())) -> " + isValidParentheses("((()))"));
        System.out.print("8. Duplicate Words: ");
        findDuplicateWords("Big black bug bit a big black dog");
        System.out.println("9. Reverse Each Word: Hello World -> " + reverseEachWord("Hello World"));
        System.out.println("10. IsOsomorphic: egg, add -> " + isIsomorphic("egg", "add"));
        System.out.println("11. Is Pangram: The quick brown fox jumps over the lazy dog -> "
                + isPangram("The quick brown fox jumps over the lazy dog"));
        System.out.println("12. Substrings: abc");
        printSubstrings("abc");
        System.out.println("13. Remove Consecutive Dups: aaabbc -> " + removeConsecutiveDeep("aaabbc"));
        System.out.println("14. Differ By One: abc, abd -> " + differByOne("abc", "abd"));
        System.out.print("15. Min Max Words: This is a test string -> ");
        minMaxWords("This is a test string");
    }
}
