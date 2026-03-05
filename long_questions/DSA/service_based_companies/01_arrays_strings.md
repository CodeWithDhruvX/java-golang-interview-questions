# Arrays and Strings (Service-Based Companies)

Service-based companies (TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, HCL, etc.) heavily focus on fundamental Array and String manipulations. These questions test your basic logical skills and algorithmic problem-solving.

## Question 1: Reverse a String
**Problem Statement:** Write a function that reverses a string. The input string is given as an array of characters or a regular string. Do this in-place with O(1) extra memory.

### Answer:
The most optimal way to reverse a string or array is by using the two-pointer approach, swapping the characters at the beginning and end until they meet in the middle.

**Code Implementation (Java):**
```java
public class ReverseString {
    public void reverseString(char[] s) {
        int left = 0;
        int right = s.length - 1;
        while (left < right) {
            char temp = s[left];
            s[left] = s[right];
            s[right] = temp;
            left++;
            right--;
        }
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 2: Two Sum
**Problem Statement:** Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

### Answer:
While a brute-force approach takes O(N^2) time, an optimized approach using a HashMap takes O(N) time. As we iterate, we store each number and its index. For each number, we check if the difference between the target and the current number exists in the HashMap.

**Code Implementation (Java):**
```java
import java.util.HashMap;

public class TwoSum {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] { map.get(complement), i };
            }
            map.put(nums[i], i);
        }
        return new int[] {}; // Should not reach here if guaranteed solution exists
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(N)

---

## Question 3: Valid Anagram
**Problem Statement:** Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

### Answer:
An anagram is a word formed by rearranging the letters of a different word, using all the original letters exactly once. We can solve this by sorting both strings or counting character frequencies. Counting frequencies is more efficient.

**Code Implementation (Java):**
```java
public class ValidAnagram {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;
        
        int[] charCounts = new int[26];
        for (int i = 0; i < s.length(); i++) {
            charCounts[s.charAt(i) - 'a']++;
            charCounts[t.charAt(i) - 'a']--;
        }
        
        for (int count : charCounts) {
            if (count != 0) return false;
        }
        return true;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1) (since the size of the array is fixed to 26)

---

## Question 4: Find the Missing Number
**Problem Statement:** Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return the only number in the range that is missing from the array.

### Answer:
The formula for the sum of the first `N` natural numbers is `N * (N + 1) / 2`. By computing this expected sum and subtracting the actual sum of the elements in the array, the result is the missing number.

**Code Implementation (Java):**
```java
public class MissingNumber {
    public int missingNumber(int[] nums) {
        int n = nums.length;
        int expectedSum = n * (n + 1) / 2;
        int actualSum = 0;
        
        for (int num : nums) {
            actualSum += num;
        }
        
        return expectedSum - actualSum;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 5: Valid Palindrome
**Problem Statement:** A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.

### Answer:
We use two pointers, one at the beginning and one at the end. We skip non-alphanumeric characters and compare the remaining characters ignoring the case.

**Code Implementation (Java):**
```java
public class ValidPalindrome {
    public boolean isPalindrome(String s) {
        int left = 0;
        int right = s.length() - 1;
        
        while (left < right) {
            while (left < right && !Character.isLetterOrDigit(s.charAt(left))) {
                left++;
            }
            while (left < right && !Character.isLetterOrDigit(s.charAt(right))) {
                right--;
            }
            if (Character.toLowerCase(s.charAt(left)) != Character.toLowerCase(s.charAt(right))) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)
