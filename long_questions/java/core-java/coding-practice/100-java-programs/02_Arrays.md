# Array Programs (16-35)

## 16. Largest Element in Array
**Principle**: Iterate and compare.
**Question**: Find the largest element in an array.
**Code**:
```java
public class LargestInArray {
    public static void main(String[] args) {
        int[] arr = {10, 50, 20, 90, 40};
        int max = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > max) max = arr[i];
        }
        System.out.println("Largest: " + max);
    }
}
```

## 17. Smallest Element in Array
**Principle**: Iterate and compare.
**Question**: Find the smallest element in an array.
**Code**:
```java
public class SmallestInArray {
    public static void main(String[] args) {
        int[] arr = {10, 5, 20, 90, 40};
        int min = arr[0];
        for (int num : arr) {
            if (num < min) min = num;
        }
        System.out.println("Smallest: " + min);
    }
}
```

## 18. Reverse an Array
**Principle**: Swap elements from start and end moving towards center.
**Question**: Reverse the elements of an array.
**Code**:
```java
import java.util.Arrays;

public class ReverseArray {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int start = 0, end = arr.length - 1;
        while (start < end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++; end--;
        }
        System.out.println("Reversed: " + Arrays.toString(arr));
    }
}
```

## 19. Sort Array (Bubble Sort)
**Principle**: Repeatedly swap adjacent elements if custom order is wrong.
**Question**: Sort an array using Bubble Sort.
**Code**:
```java
import java.util.Arrays;

public class BubbleSort {
    public static void main(String[] args) {
        int[] arr = {64, 34, 25, 12, 22, 11, 90};
        for (int i = 0; i < arr.length - 1; i++) {
            for (int j = 0; j < arr.length - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
        System.out.println("Sorted: " + Arrays.toString(arr));
    }
}
```

## 20. Linear Search
**Principle**: Check every element until match found.
**Question**: Implement Linear Search.
**Code**:
```java
public class LinearSearch {
    public static void main(String[] args) {
        int[] arr = {10, 20, 30, 40, 50};
        int key = 30, index = -1;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == key) {
                index = i;
                break;
            }
        }
        System.out.println(key + " found at index: " + index);
    }
}
```

## 21. Binary Search
**Principle**: Divide and conquer on sorted array.
**Question**: Implement Binary Search.
**Code**:
```java
public class BinarySearch {
    public static void main(String[] args) {
        int[] arr = {10, 20, 30, 40, 50}; // Must be sorted
        int key = 40, low = 0, high = arr.length - 1, index = -1;
        
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] == key) {
                index = mid;
                break;
            } else if (arr[mid] < key) low = mid + 1;
            else high = mid - 1;
        }
        System.out.println(key + " found at index: " + index);
    }
}
```

## 22. Remove Duplicates from Sorted Array
**Principle**: Use two pointers.
**Question**: Remove duplicates from a sorted array in-place.
**Code**:
```java
public class RemoveDuplicates {
    public static void main(String[] args) {
        int[] arr = {1, 1, 2, 2, 3, 4, 4, 5};
        int j = 0;
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] != arr[i + 1]) {
                arr[j++] = arr[i];
            }
        }
        arr[j++] = arr[arr.length - 1]; // Last element
        
        for (int i = 0; i < j; i++) System.out.print(arr[i] + " ");
    }
}
```

## 23. Second Largest Number
**Principle**: Track `largest` and `secondLargest`.
**Question**: Find the second largest number in an array.
**Code**:
```java
public class SecondLargest {
    public static void main(String[] args) {
        int[] arr = {12, 35, 1, 10, 34, 1};
        int largest = Integer.MIN_VALUE, second = Integer.MIN_VALUE;
        
        for (int num : arr) {
            if (num > largest) {
                second = largest;
                largest = num;
            } else if (num > second && num != largest) {
                second = num;
            }
        }
        System.out.println("Second Largest: " + second);
    }
}
```

## 24. Missing Number in Array
**Principle**: Sum of 1 to N minus Sum of Array.
**Question**: Find missing number in range 1 to N.
**Code**:
```java
public class MissingNumber {
    public static void main(String[] args) {
        int[] arr = {1, 2, 4, 5, 6};
        int n = 6; // Max number
        int expectedSum = n * (n + 1) / 2;
        int actualSum = 0;
        for (int num : arr) actualSum += num;
        
        System.out.println("Missing Number: " + (expectedSum - actualSum));
    }
}
```

## 25. Merge Two Arrays
**Principle**: Create new array of size `len1 + len2`.
**Question**: Merge two arrays into one.
**Code**:
```java
import java.util.Arrays;

public class MergeArrays {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = {4, 5, 6};
        int[] res = new int[a.length + b.length];
        
        System.arraycopy(a, 0, res, 0, a.length);
        System.arraycopy(b, 0, res, a.length, b.length);
        
        System.out.println(Arrays.toString(res));
    }
}
```

## 26. Check if Arrays are Equal
**Principle**: Check lengths, then check elements.
**Question**: Check if two arrays are equal.
**Code**:
```java
import java.util.Arrays;

public class ArrayEquality {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = {1, 2, 3};
        
        System.out.println("Equal? " + Arrays.equals(a, b));
    }
}
```

## 27. Find Common Elements
**Principle**: Nested loops or HashSet.
**Question**: Find common elements between two arrays.
**Code**:
```java
import java.util.HashSet;

public class CommonElements {
    public static void main(String[] args) {
        int[] arr1 = {1, 2, 3, 4};
        int[] arr2 = {3, 4, 5, 6};
        HashSet<Integer> set = new HashSet<>();
        
        for(int i : arr1) set.add(i);
        for(int i : arr2) {
            if(set.contains(i)) System.out.print(i + " ");
        }
    }
}
```

## 28. Left Rotate Array
**Principle**: Shift elements left by `n` positions.
**Question**: Left rotate an array by 1 position.
**Code**:
```java
import java.util.Arrays;

public class LeftRotate {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int first = arr[0];
        for (int i = 0; i < arr.length - 1; i++) {
            arr[i] = arr[i + 1];
        }
        arr[arr.length - 1] = first;
        System.out.println(Arrays.toString(arr));
    }
}
```

## 29. Right Rotate Array
**Principle**: Shift elements right by `n` positions.
**Question**: Right rotate an array by 1 position.
**Code**:
```java
import java.util.Arrays;

public class RightRotate {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int last = arr[arr.length - 1];
        for (int i = arr.length - 1; i > 0; i--) {
            arr[i] = arr[i - 1];
        }
        arr[0] = last;
        System.out.println(Arrays.toString(arr));
    }
}
```

## 30. Move Zeros to End
**Principle**: Track index of non-zero elements.
**Question**: Move all zeros to the end of the array.
**Code**:
```java
import java.util.Arrays;

public class MoveZeros {
    public static void main(String[] args) {
        int[] arr = {0, 1, 0, 3, 12};
        int index = 0;
        for (int num : arr) {
            if (num != 0) arr[index++] = num;
        }
        while (index < arr.length) arr[index++] = 0;
        System.out.println(Arrays.toString(arr));
    }
}
```

## 31. Find Duplicate Elements
**Principle**: Use HashSet or count frequency.
**Question**: Find duplicates in an int array.
**Code**:
```java
import java.util.HashSet;

public class FindDuplicates {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 1, 4, 2};
        HashSet<Integer> set = new HashSet<>();
        for (int num : arr) {
            if (!set.add(num)) System.out.println("Duplicate: " + num);
        }
    }
}
```

## 32. Frequency of Each Element
**Principle**: Use HashMap <Element, Frequency>.
**Question**: Find the frequency of each element in an array.
**Code**:
```java
import java.util.*;

public class Frequency {
    public static void main(String[] args) {
        int[] arr = {1, 2, 2, 3, 1, 4, 2};
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : arr) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }
        System.out.println(map);
    }
}
```

## 33. Odd and Even Numbers in Array
**Principle**: Modulo operator.
**Question**: Print odd and even numbers separately.
**Code**:
```java
public class OddEvenArray {
    public static void main(String[] args) {
        int[] arr = {1, 2, 5, 6, 3, 2};
        System.out.print("Odd: ");
        for(int val : arr) if(val % 2 != 0) System.out.print(val + " ");
        System.out.print("\nEven: ");
        for(int val : arr) if(val % 2 == 0) System.out.print(val + " ");
    }
}
```

## 34. Sum of Array Elements
**Principle**: Loop and accumulate.
**Question**: Calculate sum of all elements in an array.
**Code**:
```java
public class ArraySum {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int sum = 0;
        for (int num : arr) sum += num;
        System.out.println("Sum: " + sum);
    }
}
```

## 35. Sort Array in Descending Order
**Principle**: Collections.reverseOrder or manual sort.
**Question**: Sort an array in descending order.
**Code**:
```java
import java.util.*;

public class DescendingSort {
    public static void main(String[] args) {
        Integer[] arr = {5, 2, 9, 1, 6};
        Arrays.sort(arr, Collections.reverseOrder());
        System.out.println(Arrays.toString(arr));
    }
}
```
