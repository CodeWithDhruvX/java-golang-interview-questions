# Arrays, Strings, and Two Pointers (Product-Based Companies)

Product-based companies like FAANG (Meta, Amazon, Apple, Netflix, Google), Uber, Microsoft, and others heavily test candidates on optimal space and time complexity. Arrays, Strings, Sliding Window, and Two Pointer techniques form the core of these assessments.

## Question 1: Trapping Rain Water
**Problem Statement:** Given `n` non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

### Answer:
The most optimal solution uses the two-pointer approach. We maintain a `left` pointer and a `right` pointer, along with `left_max` and `right_max`. We process the smaller of the two max heights to calculate the water trapped at the current position, then move the respective pointer.

**Code Implementation (Java):**
```java
public class TrappingRainWater {
    public int trap(int[] height) {
        if (height == null || height.length == 0) return 0;
        
        int left = 0, right = height.length - 1;
        int leftMax = 0, rightMax = 0;
        int result = 0;
        
        while (left < right) {
            if (height[left] < height[right]) {
                if (height[left] >= leftMax) {
                    leftMax = height[left];
                } else {
                    result += leftMax - height[left];
                }
                left++;
            } else {
                if (height[right] >= rightMax) {
                    rightMax = height[right];
                } else {
                    result += rightMax - height[right];
                }
                right--;
            }
        }
        return result;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 2: Longest Substring Without Repeating Characters
**Problem Statement:** Given a string `s`, find the length of the longest substring without repeating characters.

### Answer:
We use the Sliding Window technique with a HashSet or an Array (for ASCII characters) to keep track of characters in the current window. We expand the window by moving the `right` pointer. If a duplicate is found, we shrink the window from the `left` until the duplicate is removed.

**Code Implementation (Java):**
```java
import java.util.HashSet;

public class LongestSubstring {
    public int lengthOfLongestSubstring(String s) {
        HashSet<Character> set = new HashSet<>();
        int left = 0, right = 0, maxLen = 0;
        
        while (right < s.length()) {
            if (!set.contains(s.charAt(right))) {
                set.add(s.charAt(right));
                maxLen = Math.max(maxLen, right - left + 1);
                right++;
            } else {
                set.remove(s.charAt(left));
                left++;
            }
        }
        return maxLen;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(min(N, M)) where N is length of string and M is charset size.

---

## Question 3: 3Sum
**Problem Statement:** Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`. The solution set must not contain duplicate triplets.

### Answer:
To avoid duplicates easily and use Two Pointers, we first sort the array. We iterate through the array, fixing one number and using two pointers (`left` and `right`) to find the other two numbers that sum up to the target (which is `-nums[i]`).

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ThreeSum {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        
        for (int i = 0; i < nums.length - 2; i++) {
            // Skip duplicates for the first element
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            
            int left = i + 1;
            int right = nums.length - 1;
            int target = -nums[i];
            
            while (left < right) {
                if (nums[left] + nums[right] == target) {
                    res.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    // Skip duplicates for the second and third elements
                    while (left < right && nums[left] == nums[left + 1]) left++;
                    while (left < right && nums[right] == nums[right - 1]) right--;
                    left++;
                    right--;
                } else if (nums[left] + nums[right] < target) {
                    left++;
                } else {
                    right--;
                }
            }
        }
        return res;
    }
}
```
**Time Complexity:** O(N^2)
**Space Complexity:** O(1) or O(N) depending on sorting implementation

---

## Question 4: Minimum Window Substring
**Problem Statement:** Given two strings `s` and `t` of lengths `m` and `n` respectively, return the minimum window substring of `s` such that every character in `t` (including duplicates) is included in the window. If there is no such substring, return the empty string `""`.

### Answer:
Use a sliding window. We keep track of the character counts required (`t` mapped) and the characters currently in our window. We expand `right` until we have a valid window, then we keep shrinking `left` to minimize it while it's still valid, keeping track of the minimum length seen.

**Code Implementation (Java):**
```java
import java.util.HashMap;

public class MinimumWindowSubstring {
    public String minWindow(String s, String t) {
        if (s.length() == 0 || t.length() == 0) return "";
        
        HashMap<Character, Integer> dictT = new HashMap<>();
        for (int i = 0; i < t.length(); i++) {
            int count = dictT.getOrDefault(t.charAt(i), 0);
            dictT.put(t.charAt(i), count + 1);
        }
        
        int required = dictT.size();
        int left = 0, right = 0;
        int formed = 0;
        
        HashMap<Character, Integer> windowCounts = new HashMap<>();
        int[] ans = {-1, 0, 0}; // length, left, right
        
        while (right < s.length()) {
            char c = s.charAt(right);
            int count = windowCounts.getOrDefault(c, 0);
            windowCounts.put(c, count + 1);
            
            if (dictT.containsKey(c) && windowCounts.get(c).intValue() == dictT.get(c).intValue()) {
                formed++;
            }
            
            while (left <= right && formed == required) {
                c = s.charAt(left);
                
                if (ans[0] == -1 || right - left + 1 < ans[0]) {
                    ans[0] = right - left + 1;
                    ans[1] = left;
                    ans[2] = right;
                }
                
                windowCounts.put(c, windowCounts.get(c) - 1);
                if (dictT.containsKey(c) && windowCounts.get(c).intValue() < dictT.get(c).intValue()) {
                    formed--;
                }
                left++;
            }
            right++;
        }
        return ans[0] == -1 ? "" : s.substring(ans[1], ans[2] + 1);
    }
}
```
**Time Complexity:** O(S + T) where S and T represent lengths of strings
**Space Complexity:** O(S + T)
