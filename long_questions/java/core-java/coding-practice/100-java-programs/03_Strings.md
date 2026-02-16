# String Programs (36-55)

## 36. Reverse a String
**Principle**: Use StringBuilder or iterate backward.
**Question**: Reverse a string without using built-in reverse function.
**Code**:
```java
public class ReverseString {
    public static void main(String[] args) {
        String str = "Hello", reversed = "";
        for (int i = 0; i < str.length(); i++) {
            reversed = str.charAt(i) + reversed;
        }
        System.out.println("Reversed: " + reversed);
    }
}
```

## 37. Check Palindrome String
**Principle**: String equals its reverse.
**Question**: Check if a string is a palindrome.
**Code**:
```java
public class PalindromeString {
    public static void main(String[] args) {
        String str = "madam";
        String reversed = new StringBuilder(str).reverse().toString();
        System.out.println(str.equals(reversed) ? "Palindrome" : "Not Palindrome");
    }
}
```

## 38. Count Vowels and Consonants
**Principle**: Check if char is in "aeiou".
**Question**: Count vowels and consonants in a string.
**Code**:
```java
public class CountVowels {
    public static void main(String[] args) {
        String str = "Hello World";
        int v = 0, c = 0;
        for (char ch : str.toLowerCase().toCharArray()) {
            if (ch >= 'a' && ch <= 'z') {
                if ("aeiou".indexOf(ch) != -1) v++;
                else c++;
            }
        }
        System.out.println("Vowels: " + v + ", Consonants: " + c);
    }
}
```

## 39. Anagram Check
**Principle**: Sort both strings and compare.
**Question**: Check if two strings are anagrams.
**Code**:
```java
import java.util.Arrays;

public class Anagram {
    public static void main(String[] args) {
        String s1 = "Listen", s2 = "Silent";
        char[] c1 = s1.toLowerCase().toCharArray();
        char[] c2 = s2.toLowerCase().toCharArray();
        Arrays.sort(c1);
        Arrays.sort(c2);
        System.out.println("Anagram? " + Arrays.equals(c1, c2));
    }
}
```

## 40. Duplicate Characters in String
**Principle**: Use HashMap or int[256].
**Question**: Find duplicate characters in a string.
**Code**:
```java
import java.util.*;

public class DuplicateChars {
    public static void main(String[] args) {
        String str = "programming";
        Map<Character, Integer> map = new HashMap<>();
        for (char ch : str.toCharArray()) {
            map.put(ch, map.getOrDefault(ch, 0) + 1);
        }
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue() > 1) {
                System.out.println(entry.getKey() + ": " + entry.getValue());
            }
        }
    }
}
```

## 41. Count Words in String
**Principle**: Split by space or count spaces + 1.
**Question**: Count number of words in a sentence.
**Code**:
```java
public class CountWords {
    public static void main(String[] args) {
        String str = "Java is fun";
        String[] words = str.trim().split("\\s+");
        System.out.println("Words: " + words.length);
    }
}
```

## 42. Remove White Spaces
**Principle**: Use `replaceAll`.
**Question**: Remove all white spaces from a string.
**Code**:
```java
public class RemoveSpaces {
    public static void main(String[] args) {
        String str = "J a v a";
        String noSpace = str.replaceAll("\\s", "");
        System.out.println(noSpace);
    }
}
```

## 43. Swapping Pair of Characters
**Principle**: Swap str[i] and str[i+1].
**Question**: Swap pairs of characters in a string.
**Code**:
```java
public class SwapPairs {
    public static void main(String[] args) {
        String str = "123456";
        char[] ch = str.toCharArray();
        for(int i=0; i<ch.length-1; i+=2) {
            char temp = ch[i];
            ch[i] = ch[i+1];
            ch[i+1] = temp;
        }
        System.out.println(new String(ch));
    }
}
```

## 44. String contains only Digits
**Principle**: Check each char or use regex.
**Question**: Check if a string contains only digits.
**Code**:
```java
public class OnlyDigits {
    public static void main(String[] args) {
        String str = "12345";
        System.out.println(str.matches("\\d+") ? "Only Digits" : "Mixed");
    }
}
```

## 45. Find First Non-Repeated Character
**Principle**: Use LinkedHashMap to maintain order + Count.
**Question**: Find the first non-repeated character in a string.
**Code**:
```java
import java.util.*;

public class FirstNonRepeated {
    public static void main(String[] args) {
        String str = "swiss";
        Map<Character, Integer> map = new LinkedHashMap<>();
        for (char ch : str.toCharArray()) map.put(ch, map.getOrDefault(ch, 0) + 1);
        
        for (Map.Entry<Character, Integer> entry : map.entrySet()) {
            if (entry.getValue() == 1) {
                System.out.println("First non-repeated: " + entry.getKey());
                break;
            }
        }
    }
}
```

## 46. Remove Special Characters
**Principle**: Regex `[^a-zA-Z0-9]`.
**Question**: Remove all special characters from a string.
**Code**:
```java
public class RemoveSpecial {
    public static void main(String[] args) {
        String str = "Ja@va#$";
        System.out.println(str.replaceAll("[^a-zA-Z0-9]", ""));
    }
}
```

## 47. String to Integer (atoi)
**Principle**: `Integer.parseInt` or manual parsing.
**Question**: Convert String to Integer.
**Code**:
```java
public class StringToInt {
    public static void main(String[] args) {
        String str = "123";
        int num = Integer.parseInt(str);
        System.out.println(num);
    }
}
```

## 48. Integer to String
**Principle**: `String.valueOf` or `Integer.toString`.
**Question**: Convert Integer to String.
**Code**:
```java
public class IntToString {
    public static void main(String[] args) {
        int i = 123;
        String s = String.valueOf(i);
        System.out.println(s);
    }
}
```

## 49. Count Occurrences of Character
**Principle**: `len - len(replace(char, ""))`.
**Question**: Count occurrences of a character in string without loop using replace.
**Code**:
```java
public class CharCount {
    public static void main(String[] args) {
        String s = "Java Programming";
        char c = 'a';
        long count = s.length() - s.replaceAll(String.valueOf(c), "").length();
        System.out.println("Count of " + c + ": " + count);
    }
}
```

## 50. Permutations of a String
**Principle**: Recursive backtracking.
**Question**: Print all permutations of a string.
**Code**:
```java
public class Permutations {
    public static void main(String[] args) {
        permute("ABC", "");
    }
    static void permute(String str, String ans) {
        if (str.length() == 0) {
            System.out.print(ans + " ");
            return;
        }
        for (int i = 0; i < str.length(); i++) {
            char ch = str.charAt(i);
            String rest = str.substring(0, i) + str.substring(i + 1);
            permute(rest, ans + ch);
        }
    }
}
```

## 51. Longest Common Prefix
**Principle**: Sort array, compare first and last strings.
**Question**: Find longest common prefix in an array of strings.
**Code**:
```java
import java.util.Arrays;

public class LongestPrefix {
    public static void main(String[] args) {
        String[] strs = {"flower", "flow", "flight"};
        Arrays.sort(strs);
        String s1 = strs[0], s2 = strs[strs.length-1];
        int i = 0;
        while(i < s1.length() && i < s2.length() && s1.charAt(i) == s2.charAt(i)) {
            i++;
        }
        System.out.println("Prefix: " + s1.substring(0, i));
    }
}
```

## 52. Reverse Words in a String
**Principle**: Split by space, reverse array, join.
**Question**: Reverse the order of words in a sentence.
**Code**:
```java
public class ReverseWords {
    public static void main(String[] args) {
        String str = "Hello World Java";
        String[] words = str.split(" ");
        String res = "";
        for(int i = words.length-1; i >= 0; i--) {
            res += words[i] + " ";
        }
        System.out.println(res.trim());
    }
}
```

## 53. Check for Subsequence
**Principle**: Two pointers.
**Question**: Check if string A is a subsequence of string B.
**Code**:
```java
public class Subsequence {
    public static void main(String[] args) {
        String s1 = "abc", s2 = "ahbgdc";
        int i=0, j=0;
        while(i < s1.length() && j < s2.length()){
            if(s1.charAt(i) == s2.charAt(j)) i++;
            j++;
        }
        System.out.println("Is Subsequence? " + (i == s1.length()));
    }
}
```

## 54. Rotation Check
**Principle**: A rotation of A is found in A+A.
**Question**: Check if one string is a rotation of another (e.g., "ABCD" and "CDAB").
**Code**:
```java
public class RotationCheck {
    public static void main(String[] args) {
        String s1 = "ABCD", s2 = "CDAB";
        if(s1.length() == s2.length() && (s1+s1).contains(s2)) {
             System.out.println("Rotation");
        } else {
             System.out.println("Not Rotation");
        }
    }
}
```

## 55. Length of String without length()
**Principle**: Convert to array and loop.
**Question**: Find length of string without using `length()` method.
**Code**:
```java
public class StringLength {
    public static void main(String[] args) {
        String str = "Hello";
        int count = 0;
        for(char c : str.toCharArray()) count++;
        System.out.println("Length: " + count);
    }
}
```
