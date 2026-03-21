```java
import java.util.*;

public class StringTransformation {

    // 1. Reverse Words in a Sentence
    // Brute Force: O(N^2), Space: O(N)
    public static String reverseWordsBruteForce(String s) {
        String[] words = s.trim().split("\\s+");
        String result = "";
        for (int i = words.length - 1; i >= 0; i--) {
            result += words[i];
            if (i > 0) {
                result += " ";
            }
        }
        return result;
    }

    // Optimized: O(N), Space: O(N)
    public static String reverseWords(String s) {
        String[] parts = s.trim().split("\\s+");
        // Reverse parts
        int left = 0, right = parts.length - 1;
        while (left < right) {
            String temp = parts[left];
            parts[left] = parts[right];
            parts[right] = temp;
            left++;
            right--;
        }
        return String.join(" ", parts);
    }

    // 2. String Compression
    // Brute Force: O(N^2), Space: O(N)
    public static int compressBruteForce(char[] chars) {
        StringBuilder result = new StringBuilder();
        int i = 0;
        while (i < chars.length) {
            char current = chars[i];
            int count = 0;
            while (i < chars.length && chars[i] == current) {
                count++;
                i++;
            }
            result.append(current);
            if (count > 1) {
                result.append(count);
            }
        }
        String compressed = result.toString();
        for (int j = 0; j < compressed.length(); j++) {
            chars[j] = compressed.charAt(j);
        }
        return compressed.length();
    }

    // Optimized: O(N), Space: O(N) (StringBuilder)
    public static int compress(char[] chars) {
        StringBuilder sb = new StringBuilder();
        int write = 0;
        int anchor = 0;
        
        for (int read = 0; read < chars.length; read++) {
            if (read + 1 == chars.length || chars[read + 1] != chars[read]) {
                chars[write] = chars[anchor];
                write++;
                
                if (read > anchor) {
                    int count = read - anchor + 1;
                    String countStr = String.valueOf(count);
                    for (int i = 0; i < countStr.length(); i++) {
                        chars[write] = countStr.charAt(i);
                        write++;
                    }
                }
                anchor = read + 1;
            }
        }
        return write;
    }

    // 3. Rotate String
    // Brute Force: O(N^2), Space: O(1)
    public static boolean rotateStringBruteForce(String s, String goal) {
        if (s.length() != goal.length()) {
            return false;
        }
        
        // Try all possible rotations
        for (int i = 0; i < s.length(); i++) {
            String rotated = s.substring(i) + s.substring(0, i);
            if (rotated.equals(goal)) {
                return true;
            }
        }
        return false;
    }

    // Optimized: O(N), Space: O(N)
    public static boolean rotateString(String s, String goal) {
        if (s.length() != goal.length()) {
            return false;
        }
        String doubled = s + s;
        return doubled.contains(goal);
    }

    // Test methods
    public static void main(String[] args) {
        // Test reverseWords
        String test1 = "the sky is blue";
        System.out.println("Reverse words in '" + test1 + "':");
        System.out.println("  Brute Force: '" + reverseWordsBruteForce(test1) + "'");
        System.out.println("  Optimized: '" + reverseWords(test1) + "'");
        
        // Test compress
        char[] test2 = {'a', 'a', 'b', 'b', 'c', 'c', 'c'};
        char[] test2Brute = test2.clone();
        int compressedLengthBrute = compressBruteForce(test2Brute);
        int compressedLength = compress(test2);
        System.out.println("\nCompressed array:");
        System.out.println("  Brute Force length: " + compressedLengthBrute);
        System.out.println("  Brute Force result: '" + new String(test2Brute, 0, compressedLengthBrute) + "'");
        System.out.println("  Optimized length: " + compressedLength);
        System.out.println("  Optimized result: '" + new String(test2, 0, compressedLength) + "'");
        
        // Test rotateString
        String test3s = "abcde";
        String test3goal = "cdeab";
        System.out.println("\nCan rotate '" + test3s + "' to get '" + test3goal + "'?");
        System.out.println("  Brute Force: " + rotateStringBruteForce(test3s, test3goal));
        System.out.println("  Optimized: " + rotateString(test3s, test3goal));
        
        String test4s = "abcde";
        String test4goal = "abced";
        System.out.println("\nCan rotate '" + test4s + "' to get '" + test4goal + "'?");
        System.out.println("  Brute Force: " + rotateStringBruteForce(test4s, test4goal));
        System.out.println("  Optimized: " + rotateString(test4s, test4goal));
    }
}
```
