# Coding: Strings & Arrays - Interview Answers

> 🎯 **Focus:** Explain your logic *before* you write. Use built-in methods (like `StringBuilder`) unless asked to avoid them.

### 1. Reverse a String
"I’ll use a `StringBuilder` because Strings are immutable in Java. Manually concatenating in a loop would create garbage objects and be O(N²). `StringBuilder` is O(N)."

```java
public String reverse(String input) {
    if (input == null) return null;
    return new StringBuilder(input).reverse().toString();
}

// If asked to do it manually (without StringBuilder):
public String reverseManual(String input) {
    char[] chars = input.toCharArray();
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
```

---

### 2. Check for Palindrome
"A palindrome reads the same forwards and backwards. I’ll use the two-pointer approach. I start from both ends and move inwards. If I find a mismatch, it's false immediately. It’s O(N) time and O(1) space."

```java
public boolean isPalindrome(String str) {
    if (str == null) return false;
    int left = 0;
    int right = str.length() - 1;
    
    while (left < right) {
        if (str.charAt(left) != str.charAt(right)) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}
```

---

### 3. Find Duplicates in an Array
"I’ll use a `HashSet`. As I iterate through the array, I try to add each element to the set. `set.add()` returns `false` if the element is already there. That’s my duplicate. This is O(N) time complexity."

```java
public void findDuplicates(int[] nums) {
    Set<Integer> seen = new HashSet<>();
    for (int num : nums) {
        if (!seen.add(num)) {
            System.out.println("Duplicate found: " + num);
        }
    }
}
```

---

### 4. Check if two strings are Anagrams
"Anagrams have the exact same character counts.
I could sort both strings and compare them (O(N log N)), but a Frequency Array (or HashMap) is faster (O(N)).
I'll iterate the first string to increment counts, and the second to decrement. If all counts are zero at the end, it's an anagram."

```java
public boolean isAnagram(String s1, String s2) {
    if (s1.length() != s2.length()) return false;
    
    int[] counts = new int[26]; // Assuming lowercase 'a'-'z'
    
    for (int i = 0; i < s1.length(); i++) {
        counts[s1.charAt(i) - 'a']++;
        counts[s2.charAt(i) - 'a']--;
    }
    
    for (int count : counts) {
        if (count != 0) return false;
    }
    return true;
}
```

---

### 5. Two Sum (Find pair with given sum)
"The brute force way is nested loops, which is O(N²).
I can do better using a `HashMap` to store the complement.
For each number `n`, I check if `target - n` exists in the map. If yes, that's the pair. If no, I add `n` to the map. This brings it down to O(N)."

```java
public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement)) {
            return new int[] { map.get(complement), i }; // Found indices
        }
        map.put(nums[i], i);
    }
    throw new IllegalArgumentException("No solution found");
}
```

---

### 6. Count Character Occurrences
"I'll use the Stream API for this because it's concise. I convert the String to a Stream of characters, then collect them using `groupingBy` and `counting`. It returns a Map of Character to Long."

```java
public Map<String, Long> countChars(String input) {
    return Arrays.stream(input.split(""))
                 .collect(Collectors.groupingBy(
                     Function.identity(), 
                     Collectors.counting()
                 ));
}
```

---

### 7. Find Second Largest Number
"I don't need to sort the array. I can do this in a single pass O(N).
I'll maintain two variables: `largest` and `secondLargest`. I update them as I iterate through the numbers."

```java
public int findSecondLargest(int[] arr) {
    int largest = Integer.MIN_VALUE;
    int second = Integer.MIN_VALUE;
    
    for (int num : arr) {
        if (num > largest) {
            second = largest;
            largest = num;
        } else if (num > second && num != largest) {
            second = num;
        }
    }
    return second;
}
```
