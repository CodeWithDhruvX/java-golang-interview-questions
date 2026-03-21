# Java Array Problems - Complete Solutions

This document contains all basic array problems with their Java implementations, including both brute force and optimized approaches.

---

## 1. Find Largest and Smallest Element

### Problem Statement
Find the maximum and minimum values in an array in a single pass. Return both values.

### Pattern/Category
Single Pass / Two Variables Tracking

### Brute Force Solution
**Approach**: Sort the array and pick first and last elements.
**Time Complexity**: O(N log N)
**Space Complexity**: O(N) for the copy

```java
public static int[] findMinMaxBruteForce(int[] arr) {
    if (arr.length == 0) return new int[]{};
    
    // Create a copy and sort it
    int[] sorted = arr.clone();
    java.util.Arrays.sort(sorted);
    
    // First element is min, last is max
    return new int[]{sorted[0], sorted[sorted.length - 1]};
}
```

### Optimized Solution
**Approach**: Single pass maintaining two variables.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static int[] findMinMaxOptimized(int[] arr) {
    if (arr.length == 0) return new int[]{};
    
    int minVal = arr[0];
    int maxVal = arr[0];
    
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > maxVal) {
            maxVal = arr[i];
        }
        if (arr[i] < minVal) {
            minVal = arr[i];
        }
    }
    return new int[]{minVal, maxVal};
}
```

---

## 2. Reverse an Array In-Place

### Problem Statement
Reverse the given array without using a new array (modify original).
Example: `[1, 2, 3]` -> `[3, 2, 1]`

### Pattern/Category
Two Pointers (Start & End)

### Brute Force Solution
**Approach**: Create a new array and copy elements in reverse order.
**Time Complexity**: O(N)
**Space Complexity**: O(N)

```java
public static int[] reverseBruteForce(int[] arr) {
    int n = arr.length;
    int[] reversed = new int[n];
    
    for (int i = 0; i < n; i++) {
        reversed[i] = arr[n - 1 - i];
    }
    return reversed;
}
```

### Optimized Solution
**Approach**: Two pointers swapping elements.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static void reverseOptimized(int[] arr) {
    int left = 0;
    int right = arr.length - 1;
    
    while (left < right) {
        // Swap elements
        int temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
        
        left++;
        right--;
    }
}
```

---

## 3. Find Second Largest Element

### Problem Statement
Find the second highest value in the array. Must be strictly smaller than the largest.

### Pattern/Category
Two Variables / Tournament Logic

### Brute Force Solution
**Approach**: Sort array and pick second last element.
**Time Complexity**: O(N log N)
**Space Complexity**: O(N)

```java
public static int secondLargestBruteForce(int[] arr) {
    if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
    
    int[] sorted = arr.clone();
    java.util.Arrays.sort(sorted);
    
    // Handle duplicates - find the first element different from the largest
    for (int i = sorted.length - 2; i >= 0; i--) {
        if (sorted[i] != sorted[sorted.length - 1]) {
            return sorted[i];
        }
    }
    throw new IllegalArgumentException("No second largest element (all elements are equal)");
}
```

### Optimized Solution
**Approach**: Single pass maintaining largest and second largest.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static int secondLargestOptimized(int[] arr) {
    if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
    
    int largest = Integer.MIN_VALUE;
    int secondLargest = Integer.MIN_VALUE;
    
    for (int num : arr) {
        if (num > largest) {
            secondLargest = largest;
            largest = num;
        } else if (num > secondLargest && num != largest) {
            secondLargest = num;
        }
    }
    
    if (secondLargest == Integer.MIN_VALUE) {
        throw new IllegalArgumentException("No second largest element (all elements are equal)");
    }
    return secondLargest;
}
```

---

## 4. Check if Array is Sorted

### Problem Statement
Return `true` if array is sorted in non-decreasing order, `false` otherwise.

### Pattern/Category
Linear Scan / Adjacent Comparison

### Brute Force Solution
**Approach**: Sort a copy and compare with original.
**Time Complexity**: O(N log N)
**Space Complexity**: O(N)

```java
public static boolean isSortedBruteForce(int[] arr) {
    int[] sorted = arr.clone();
    java.util.Arrays.sort(sorted);
    return java.util.Arrays.equals(arr, sorted);
}
```

### Optimized Solution
**Approach**: Single pass checking adjacent elements.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static boolean isSortedOptimized(int[] arr) {
    for (int i = 0; i < arr.length - 1; i++) {
        if (arr[i] > arr[i + 1]) {
            return false;
        }
    }
    return true;
}
```

---

## 5. Count Even and Odd Elements

### Problem Statement
Count how many integers are even and how many are odd in an array.

### Pattern/Category
Modulo Arithmetic / Counter

### Brute Force Solution
**Approach**: Iterate and check parity (already optimal).
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static int[] countEvenOddBruteForce(int[] arr) {
    int evenCount = 0;
    int oddCount = 0;
    
    for (int num : arr) {
        if (num % 2 == 0) {
            evenCount++;
        } else {
            oddCount++;
        }
    }
    return new int[]{evenCount, oddCount};
}
```

### Optimized Solution
**Approach**: Same as brute force (already optimal).
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static int[] countEvenOddOptimized(int[] arr) {
    return countEvenOddBruteForce(arr);
}
```

---

## 6. Remove Duplicates from Sorted Array

### Problem Statement
Given a **sorted** array, remove duplicates **in-place** such that each element appears only once and return the new length.
`[1,1,2]` -> `[1,2]`, return 2.

### Pattern/Category
Two Pointers (Reader & Writer)

### Brute Force Solution
**Approach**: Use HashSet to store unique elements.
**Time Complexity**: O(N)
**Space Complexity**: O(N)

```java
public static int removeDuplicatesBruteForce(int[] arr) {
    Set<Integer> unique = new LinkedHashSet<>();
    for (int num : arr) {
        unique.add(num);
    }
    
    int index = 0;
    for (int num : unique) {
        arr[index++] = num;
    }
    return unique.size();
}
```

### Optimized Solution
**Approach**: Two pointers technique.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static int removeDuplicatesOptimized(int[] arr) {
    if (arr.length == 0) return 0;
    
    int i = 0;
    for (int j = 1; j < arr.length; j++) {
        if (arr[j] != arr[i]) {
            i++;
            arr[i] = arr[j];
        }
    }
    return i + 1;
}
```

---

## 7. Left Rotate Array by 1 Position

### Problem Statement
Shift all elements to the left by 1. The first element moves to the last position.
`[1, 2, 3, 4, 5]` -> `[2, 3, 4, 5, 1]`

### Pattern/Category
Store First & Shift

### Brute Force Solution
**Approach**: Create new array with shifted positions.
**Time Complexity**: O(N)
**Space Complexity**: O(N)

```java
public static int[] leftRotateByOneBruteForce(int[] arr) {
    int n = arr.length;
    if (n == 0) return arr;
    
    int[] rotated = new int[n];
    for (int i = 0; i < n - 1; i++) {
        rotated[i] = arr[i + 1];
    }
    rotated[n - 1] = arr[0];
    return rotated;
}
```

### Optimized Solution
**Approach**: Store first element, shift others, place first at end.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static void leftRotateByOneOptimized(int[] arr) {
    if (arr.length <= 1) return;
    
    int temp = arr[0];
    for (int i = 1; i < arr.length; i++) {
        arr[i - 1] = arr[i];
    }
    arr[arr.length - 1] = temp;
}
```

---

## 8. Left Rotate Array by K Positions

### Problem Statement
Rotate array to left by `K` steps.
`[1,2,3,4,5]`, K=2 -> `[3,4,5,1,2]`.

### Pattern/Category
Reversal Algorithm

### Brute Force Solution
**Approach**: Call rotate by 1, K times.
**Time Complexity**: O(N*K)
**Space Complexity**: O(1)

```java
public static void leftRotateByKBruteForce(int[] arr, int k) {
    int n = arr.length;
    if (n == 0) return;
    
    k = k % n;
    
    for (int i = 0; i < k; i++) {
        leftRotateByOneOptimized(arr);
    }
}
```

### Optimized Solution
**Approach**: Reversal algorithm.
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static void leftRotateByKOptimized(int[] arr, int k) {
    int n = arr.length;
    if (n == 0) return;
    
    k = k % n;
    if (k == 0) return;
    
    // Reverse first k elements
    reverse(arr, 0, k - 1);
    // Reverse remaining n-k elements
    reverse(arr, k, n - 1);
    // Reverse entire array
    reverse(arr, 0, n - 1);
}

private static void reverse(int[] arr, int start, int end) {
    while (start < end) {
        int temp = arr[start];
        arr[start] = arr[end];
        arr[end] = temp;
        start++;
        end--;
    }
}
```

---

## 9. Find Sum of All Elements

### Problem Statement
Calculate the sum of all integers in the array.

### Pattern/Category
Accumulator

### Brute Force Solution
**Approach**: Simple iteration (already optimal).
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static long sumBruteForce(int[] arr) {
    long sum = 0;
    for (int num : arr) {
        sum += num;
    }
    return sum;
}
```

### Optimized Solution
**Approach**: Same as brute force (already optimal).
**Time Complexity**: O(N)
**Space Complexity**: O(1)

```java
public static long sumOptimized(int[] arr) {
    return sumBruteForce(arr);
}
```

---

## 10. Find Frequency of Each Element

### Problem Statement
Count how many times each element appears.
`[1, 1, 2, 3]` -> `{1: 2, 2: 1, 3: 1}`

### Pattern/Category
Hash Map / Dictionary

### Brute Force Solution
**Approach**: Nested loops to count each element.
**Time Complexity**: O(N²)
**Space Complexity**: O(N)

```java
public static Map<Integer, Integer> frequencyBruteForce(int[] arr) {
    Map<Integer, Integer> freq = new HashMap<>();
    
    for (int i = 0; i < arr.length; i++) {
        int count = 0;
        for (int j = 0; j < arr.length; j++) {
            if (arr[i] == arr[j]) {
                count++;
            }
        }
        freq.put(arr[i], count);
    }
    return freq;
}
```

### Optimized Solution
**Approach**: HashMap single pass.
**Time Complexity**: O(N)
**Space Complexity**: O(N)

```java
public static Map<Integer, Integer> frequencyOptimized(int[] arr) {
    Map<Integer, Integer> freq = new HashMap<>();
    
    for (int num : arr) {
        freq.put(num, freq.getOrDefault(num, 0) + 1);
    }
    return freq;
}
```

---

## Summary of Time and Space Complexities

| Problem | Brute Force Time | Optimized Time | Brute Force Space | Optimized Space |
|---------|------------------|----------------|-------------------|-----------------|
| Min Max | O(N log N) | O(N) | O(N) | O(1) |
| Reverse | O(N) | O(N) | O(N) | O(1) |
| Second Largest | O(N log N) | O(N) | O(N) | O(1) |
| Is Sorted | O(N log N) | O(N) | O(N) | O(1) |
| Count Even/Odd | O(N) | O(N) | O(1) | O(1) |
| Remove Duplicates | O(N) | O(N) | O(N) | O(1) |
| Left Rotate by 1 | O(N) | O(N) | O(N) | O(1) |
| Left Rotate by K | O(N*K) | O(N) | O(1) | O(1) |
| Array Sum | O(N) | O(N) | O(1) | O(1) |
| Frequency Count | O(N²) | O(N) | O(N) | O(N) |

---

## Individual Java Files

Each problem is also available as a separate Java file:

1. `1_FindMinMax.java`
2. `2_ReverseArray.java`
3. `3_SecondLargest.java`
4. `4_IsSorted.java`
5. `5_CountEvenOdd.java`
6. `6_RemoveDuplicates.java`
7. `7_LeftRotateByOne.java`
8. `8_LeftRotateByK.java`
9. `9_ArraySum.java`
10. `10_FrequencyCount.java`

Each file contains both brute force and optimized solutions with test cases in the main method.
