# Service Based Companies Interview Questions (TCS, Infosys, Wipro, etc.)

These are frequently asked coding questions in Indian service-based IT companies. The focus is often on logic building without using built-in library functions where possible.

## 1. Rotational Array

**Problem:** Rotate an array of `n` elements to the right by `k` steps.

### Approach 1: Brute Force (Rotate one by one)
Rotate the array by 1 position `k` times. Time Complexity: O(n*k).

### Approach 2: Reversal Algorithm (Optimized)
**Time Complexity:** O(n)
**Space Complexity:** O(1)

**Algorithm:**
1. Reverse the whole array (0 to n-1).
2. Reverse the first `k` elements (0 to k-1).
3. Reverse the remaining `n-k` elements (k to n-1).
*(Note: For Left Rotate, the logic is slightly different: Reverse(0, k-1), Reverse(k, n-1), Reverse(0, n-1))*

**Code (Java):**
```java
import java.util.Arrays;

public class ArrayRotation {
    public static void rotateRight(int[] nums, int k) {
        if (nums == null || nums.length == 0) return;
        int n = nums.length;
        k = k % n; // Handle k > n
        
        // 1. Reverse the whole array
        reverse(nums, 0, n - 1);
        // 2. Reverse the first k elements
        reverse(nums, 0, k - 1);
        // 3. Reverse the rest
        reverse(nums, k, n - 1);
    }

    private static void reverse(int[] nums, int start, int end) {
        while (start < end) {
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int k = 3;
        rotateRight(arr, k);
        System.out.println("Rotated Array: " + Arrays.toString(arr)); 
        // Output: [5, 6, 7, 1, 2, 3, 4]
    }
}
```

---

## 2. Pattern Programs

### A. Pyramid Pattern
**Structure:**
```
    *
   ***
  *****
 *******
```

**Code (Java):**
```java
public class PyramidPattern {
    public static void main(String[] args) {
        int rows = 5;
        for (int i = 1; i <= rows; i++) {
            // Print spaces
            for (int j = 1; j <= rows - i; j++) {
                System.out.print(" ");
            }
            // Print stars
            for (int k = 1; k <= 2 * i - 1; k++) {
                System.out.print("*");
            }
            System.out.println();
        }
    }
}
```

### B. Zig-Zag Pattern (or Spiral/Snake)
Often refers to printing string in ZigZag or printing a matrix in ZigZag. A common simple "Zig-Zag" star pattern question looks like mountain peaks.
Or specific number patterns. Here is a generic Zig-Zag logic for a string (LeetCode style) or simple number printing. 

**Simple Number Zig-Zag:**
```
1 
2 3 
4 5 6 
7 8 9 10 
6 5 4 
3 2 
1
```
*(This varies heavily by specific question, but here is a standard "Hollow Inverted V" or similar is common. Below is the standard "Number Pyramid" often called Zig-Zag variation in some tests)*

**Code (Standard Zig-Zag Number Pyramid):**
```java
public class ZigZagPattern {
    public static void main(String[] args) {
        int n = 5;
        // Upper part
        for (int i = 1; i <= n; i++) {
           for (int j = 1; j <= i; j++) System.out.print(j + " ");
           System.out.println();
        }
        // Lower part
        for (int i = n - 1; i >= 1; i--) {
           for (int j = 1; j <= i; j++) System.out.print(j + " ");
           System.out.println();
        }
    }
}
```

---

## 3. Reverse a String

**Problem:** Reverse a given string without using built-in `reverse()` method.

**Code (Java):**
```java
public class StringReverse {
    public static void main(String[] args) {
        String str = "Automation";
        
        // Approach 1: Using char array (Two Pointer) - Efficient
        char[] charArray = str.toCharArray();
        int left = 0;
        int right = charArray.length - 1;
        
        while (left < right) {
            char temp = charArray[left];
            charArray[left] = charArray[right];
            charArray[right] = temp;
            left++;
            right--;
        }
        
        String reversed = new String(charArray);
        System.out.println("Reversed: " + reversed);
        
        // Approach 2: Using StringBuilder (for reference)
        // System.out.println(new StringBuilder(str).reverse().toString());
    }
}
```

---

## 4. Anagram Check

**Problem:** Check if two strings are anagrams of each other (contain same characters with same frequency).

**Approach:** Frequency Array (Best for ASCII strings).
**Time Complexity:** O(n)

**Code (Java):**
```java
import java.util.Arrays;

public class AnagramCheck {
    public static boolean isAnagram(String s1, String s2) {
        // Remove whitespace and convert to lower case for consistency
        s1 = s1.replaceAll("\\s", "").toLowerCase();
        s2 = s2.replaceAll("\\s", "").toLowerCase();

        if (s1.length() != s2.length()) return false;

        int[] count = new int[26]; // Assuming English alphabets

        for (int i = 0; i < s1.length(); i++) {
            count[s1.charAt(i) - 'a']++;
            count[s2.charAt(i) - 'a']--;
        }

        for (int c : count) {
            if (c != 0) return false;
        }
        return true;
    }

    public static void main(String[] args) {
        System.out.println(isAnagram("Listen", "Silent")); // true
        System.out.println(isAnagram("Hello", "World"));   // false
    }
}
```
